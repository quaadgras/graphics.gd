// Meta Quest (Horizon OS / OpenXR) build target.
//
// GOOS=metaquest is a pseudo-platform: it builds the Go shared library
// for android/arm64 (same as GOOS=android, GOARCH=arm64), then post-
// processes the Godot-exported APK to bake in the Khronos OpenXR
// Android loader and the GodotVR Meta vendor plugin pulled directly
// from Maven Central. The result is a standalone Quest-ready APK that
// requires no gradle build and no Android SDK on the user's machine
// beyond what graphics.gd already manages.
//
// Dependencies are all Apache-2.0 / BSD-3-Clause and fetched at build
// time from Maven Central:
//
//   - org.khronos.openxr:openxr_loader_for_android   — the OpenXR loader .so
//   - org.godotengine:godot-openxr-vendors-meta      — the Meta runtime adapter
//   - com.android.tools.r8:r8                        — used as `D8` to compile
//     the vendor plugin's
//     classes.jar → dex
//
// `java` must be on PATH at build time to run the R8 jar; this is the
// only non-graphics.gd-managed prerequisite. We do not auto-install a
// JDK (distributions vary too much) — users see an install hint instead.
package builder

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

// MetaQuest extends the Android builder with an injection step that
// bakes the OpenXR loader + Meta vendor plugin into the produced APK.
type MetaQuest struct {
	Android
}

func (mq MetaQuest) Build(args ...string) error {
	// Force arm64 — Quest has no other targets — and delegate to the
	// regular Android compile path. The post-processing is only done
	// in BuildMain / Run after Godot has produced the APK.
	if os.Getenv("GOARCH") == "" {
		os.Setenv("GOARCH", "arm64")
	}
	os.Setenv("GOOS", "android")
	return mq.Android.Build(args...)
}

func (mq MetaQuest) Test(args ...string) error {
	return fmt.Errorf("gd test: metaquest not supported")
}

func (mq MetaQuest) BuildMain(args ...string) error {
	if err := mq.Build(args...); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-release", "Android"); err != nil {
		return xray.New(err)
	}
	apk := filepath.Join(project.ReleasesDirectory, "android", "arm64", project.Name+".apk")
	if err := injectMetaQuest(apk); err != nil {
		return xray.New(err)
	}
	return signAPK(apk, releaseKeystore())
}

func (mq MetaQuest) Run(args ...string) error {
	if err := mq.Build(args...); err != nil {
		return xray.New(err)
	}
	adb, err := tooling.AndroidDebugBridge.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if _, err := tooling.AndroidPackageSigner.Lookup(); err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll(filepath.Join(project.ReleasesDirectory, "android", "arm64"), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-debug", "Android"); err != nil {
		return xray.New(err)
	}
	apk := filepath.Join(project.ReleasesDirectory, "android", "arm64", project.Name+".apk")
	if err := injectMetaQuest(apk); err != nil {
		return xray.New(err)
	}
	if err := signAPK(apk, debugKeystore()); err != nil {
		return xray.New(err)
	}

	install := exec.Command(adb, "install", "-r", apk)
	install.Stdout, install.Stderr = os.Stdout, os.Stderr
	if err := install.Run(); err != nil {
		fmt.Println("Quest not recognized? Enable Developer Mode in the Meta Horizon app and accept USB debugging on the headset.")
		return xray.New(err)
	}
	pkg := "com.example." + project.AndroidSafePackageName(path.Base(project.Directory))
	_ = exec.Command(adb, "logcat", "-c").Run()
	launch := exec.Command(adb, "shell", "am", "start", "-a", "android.intent.action.MAIN",
		"-c", "org.khronos.openxr.intent.category.IMMERSIVE_HMD",
		"-n", pkg+"/com.godot.game.GodotApp")
	launch.Stdout, launch.Stderr = os.Stdout, os.Stderr
	if err := launch.Run(); err != nil {
		return xray.New(err)
	}
	tail := exec.Command(adb, "logcat", "*:W")
	tail.Stdout, tail.Stderr = os.Stdout, os.Stderr
	return tail.Run()
}

// injectMetaQuest is the keystone: take the APK Godot just exported,
// drop in the Meta plugin classes (as a second dex), drop in the
// libopenxr_loader.so + vendor .so files, rewrite the manifest with
// the Quest features / permissions / intent filters, and zip it back
// up. After this returns the APK still needs to be re-signed.
func injectMetaQuest(apkPath string) error {
	work, err := os.MkdirTemp("", "metaquest-")
	if err != nil {
		return xray.New(err)
	}
	defer os.RemoveAll(work)

	loaderAAR, err := tooling.OpenXRLoaderAAR.Lookup()
	if err != nil {
		return xray.New(err)
	}
	vendorAAR, err := tooling.OpenXRVendorsMetaAAR.Lookup()
	if err != nil {
		return xray.New(err)
	}
	r8jar, err := tooling.R8.Lookup()
	if err != nil {
		return xray.New(err)
	}
	javaBin, err := tooling.Java.Lookup()
	if err != nil {
		return xray.New(err)
	}

	// Materialize the AAR contents we care about: classes.jar from
	// each, native libs from each. AARs are plain zip files.
	loaderDir := filepath.Join(work, "loader")
	vendorDir := filepath.Join(work, "vendor")
	if err := tooling.ExtractArchive(loaderAAR, loaderDir, "zip", "", false); err != nil {
		return xray.New(err)
	}
	if err := tooling.ExtractArchive(vendorAAR, vendorDir, "zip", "", false); err != nil {
		return xray.New(err)
	}

	// Use R8/D8 to compile the vendor plugin's classes.jar into a dex.
	// We pass --release so D8 strips debug info and uses smaller refs.
	// The loader AAR's classes.jar is small (mostly stubs) — bundle it
	// into the same D8 invocation so we end up with a single classes2.dex
	// to slot into the APK alongside Godot's existing classes.dex.
	dexOut := filepath.Join(work, "dex")
	if err := os.MkdirAll(dexOut, 0755); err != nil {
		return xray.New(err)
	}
	d8Args := []string{"-cp", r8jar, "com.android.tools.r8.D8", "--release",
		"--min-api", "29", // Quest 2 is API 29; Quest 3 is higher
		"--output", dexOut,
	}
	for _, dir := range []string{vendorDir, loaderDir} {
		jar := filepath.Join(dir, "classes.jar")
		if _, err := os.Stat(jar); err == nil {
			d8Args = append(d8Args, jar)
		}
	}
	d8 := exec.Command(javaBin, d8Args...)
	d8.Stdout, d8.Stderr = os.Stdout, os.Stderr
	if err := d8.Run(); err != nil {
		return xray.New(fmt.Errorf("d8 dex compile failed: %w", err))
	}
	// D8 writes classes.dex by default; rename to classes2.dex so it
	// merges with Godot's classes.dex as a secondary multidex slot.
	if err := os.Rename(filepath.Join(dexOut, "classes.dex"), filepath.Join(dexOut, "classes2.dex")); err != nil {
		return xray.New(err)
	}

	// Now decompile the APK so we can rewrite the binary AndroidManifest.xml
	// through apktool's smali/xml-tools, then repack. We don't touch the
	// existing dex — apktool will re-bundle it as-is.
	decompiled := filepath.Join(work, "decompiled")
	if err := tooling.AndroidPackageKitTool.Exec("d", apkPath, "-s", "-o", decompiled, "-f"); err != nil {
		return xray.New(err)
	}
	manifestPath := filepath.Join(decompiled, "AndroidManifest.xml")
	patched, err := patchManifestForQuest(manifestPath)
	if err != nil {
		return xray.New(err)
	}
	if err := os.WriteFile(manifestPath, patched, 0644); err != nil {
		return xray.New(err)
	}

	// Copy the OpenXR loader .so plus any vendor .so into the lib tree.
	for _, src := range []string{
		filepath.Join(loaderDir, "jni", "arm64-v8a", "libopenxr_loader.so"),
	} {
		if _, err := os.Stat(src); err != nil {
			continue
		}
		dst := filepath.Join(decompiled, "lib", "arm64-v8a", filepath.Base(src))
		if err := copyFile(src, dst); err != nil {
			return xray.New(err)
		}
	}
	// Vendor AAR may carry additional .so files (Meta runtime hooks).
	vendorLibs := filepath.Join(vendorDir, "jni", "arm64-v8a")
	if entries, err := os.ReadDir(vendorLibs); err == nil {
		for _, e := range entries {
			if e.IsDir() || !strings.HasSuffix(e.Name(), ".so") {
				continue
			}
			// The app gradle build (godot/platform/android/java/app/build.gradle:313)
			// explicitly deletes the vendor's bundled libopenxr_loader.so in favour
			// of the one from the openxr_loader_for_android dependency. Mirror that
			// — we already injected the loader above.
			if e.Name() == "libopenxr_loader.so" {
				continue
			}
			if err := copyFile(filepath.Join(vendorLibs, e.Name()),
				filepath.Join(decompiled, "lib", "arm64-v8a", e.Name())); err != nil {
				return xray.New(err)
			}
		}
	}

	// Apktool puts secondary dex files at the top level. Slot in our
	// freshly-compiled classes2.dex so the Meta plugin gets loaded by
	// the Android runtime alongside Godot's classes.dex.
	if err := copyFile(filepath.Join(dexOut, "classes2.dex"),
		filepath.Join(decompiled, "classes2.dex")); err != nil {
		return xray.New(err)
	}

	// Rebuild. Apktool's `b` re-encodes the manifest to binary AXML
	// and bundles everything in decompiled/ back into an APK.
	repacked := apkPath + ".meta-unsigned.apk"
	if err := tooling.AndroidPackageKitTool.Exec("b", decompiled, "-o", repacked); err != nil {
		return xray.New(err)
	}
	if err := os.Rename(repacked, apkPath); err != nil {
		return xray.New(err)
	}
	return nil
}

// patchManifestForQuest rewrites the decompiled AndroidManifest.xml
// emitted by apktool so it advertises the features / permissions
// Horizon OS requires, and registers the Meta OpenXR vendor plugin
// for Godot to discover at startup.
//
// We mutate the XML conservatively — additive entries, plus a single
// targeted intent-filter swap on the main activity so the launcher
// understands this is an immersive headset app.
func patchManifestForQuest(manifestPath string) ([]byte, error) {
	src, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, xray.New(err)
	}

	// Quick sanity check — apktool emits well-formed XML; if this ever
	// changes we want to fail loudly rather than corrupt the manifest.
	if !bytes.Contains(src, []byte("<manifest")) {
		return nil, fmt.Errorf("decompiled AndroidManifest.xml missing <manifest> root")
	}

	// 1. Add required uses-feature / uses-permission entries just after
	//    the opening <manifest ...> tag. Idempotent: skip lines that
	//    are already present.
	additions := strings.Builder{}
	for _, line := range []string{
		`<uses-feature android:name="android.hardware.vr.headtracking" android:required="true" android:version="1"/>`,
		`<uses-feature android:name="oculus.software.handtracking" android:required="false"/>`,
		`<uses-feature android:name="com.oculus.feature.PASSTHROUGH" android:required="false"/>`,
		`<uses-feature android:name="com.oculus.feature.RENDER_MODEL" android:required="false"/>`,
		`<uses-permission android:name="com.oculus.permission.HAND_TRACKING"/>`,
		`<uses-permission android:name="com.oculus.permission.USE_SCENE"/>`,
		`<uses-permission android:name="com.oculus.permission.USE_ANCHOR_API"/>`,
		`<uses-permission android:name="com.oculus.permission.RENDER_MODEL"/>`,
	} {
		// extract the "name" attribute and skip if the manifest already
		// has a tag for it (apktool keeps attributes in source order)
		if name := xmlNameAttr(line); name != "" && bytes.Contains(src, []byte(`android:name="`+name+`"`)) {
			continue
		}
		additions.WriteString("    " + line + "\n")
	}

	// 2. Splice the additions in directly after the <manifest ...> tag.
	if additions.Len() > 0 {
		close := bytes.Index(src, []byte(">"))
		if close < 0 {
			return nil, fmt.Errorf("malformed manifest: no '>' after <manifest")
		}
		// Walk forward to actual end of opening tag (in case of '>' in attrs).
		close = findTagEnd(src, bytes.Index(src, []byte("<manifest")))
		if close < 0 {
			return nil, fmt.Errorf("malformed manifest: cannot find end of <manifest> opening tag")
		}
		src = append(append(append([]byte{}, src[:close+1]...),
			append([]byte("\n"), []byte(additions.String())...)...),
			src[close+1:]...)
	}

	// 3. Make the main activity launchable on Quest. Godot's app
	//    template names the main activity .GodotApp with a normal
	//    LAUNCHER intent-filter. Quest's launcher looks for the
	//    IMMERSIVE_HMD + com.oculus.intent.category.VR pair instead.
	src = bytes.Replace(src,
		[]byte(`<category android:name="android.intent.category.LAUNCHER"/>`),
		[]byte(`<category android:name="android.intent.category.LAUNCHER"/>`+"\n"+
			`                <category android:name="com.oculus.intent.category.VR"/>`+"\n"+
			`                <category android:name="org.khronos.openxr.intent.category.IMMERSIVE_HMD"/>`),
		1)

	// 4. Add Quest-specific meta-data and the Godot OpenXR plugin
	//    registration just before </application>. The plugin name
	//    matches what the godot-openxr-vendors-meta AAR exposes
	//    (see godot/platform/android/java/editor/src/horizonos/AndroidManifest.xml).
	metaBlock := `        <meta-data android:name="com.oculus.supportedDevices" android:value="quest2|quest3|questpro|quest3s"/>` + "\n" +
		`        <meta-data android:name="com.oculus.ossplash" android:value="true"/>` + "\n" +
		`        <meta-data android:name="com.oculus.handtracking.version" android:value="V2.0"/>` + "\n" +
		`        <meta-data android:name="org.godotengine.plugin.v2.GodotOpenXR" android:value="org.godotengine.openxr.vendors.GodotOpenXR"/>` + "\n"
	src = bytes.Replace(src, []byte("</application>"), append([]byte(metaBlock), []byte("</application>")...), 1)

	// Final validation: still parseable XML?
	if err := xml.Unmarshal(src, new(struct {
		XMLName xml.Name `xml:"manifest"`
	})); err != nil {
		return nil, fmt.Errorf("patched manifest is not valid XML: %w", err)
	}
	return src, nil
}

// xmlNameAttr extracts the value of an `android:name="..."` attribute
// from a single XML tag string. Returns "" if not present. Cheap and
// loose; only used for de-duplication of additions, so false negatives
// are safe (they just produce a redundant tag).
func xmlNameAttr(tag string) string {
	const prefix = `android:name="`
	i := strings.Index(tag, prefix)
	if i < 0 {
		return ""
	}
	rest := tag[i+len(prefix):]
	end := strings.IndexByte(rest, '"')
	if end < 0 {
		return ""
	}
	return rest[:end]
}

// findTagEnd locates the '>' that closes the XML opening tag starting
// at start. Returns the index of '>' or -1 if not found. Handles single
// and double quoted attribute values so embedded '>' inside attributes
// don't confuse us.
func findTagEnd(src []byte, start int) int {
	if start < 0 || start >= len(src) {
		return -1
	}
	inQuote := byte(0)
	for i := start; i < len(src); i++ {
		c := src[i]
		switch {
		case inQuote != 0:
			if c == inQuote {
				inQuote = 0
			}
		case c == '"' || c == '\'':
			inQuote = c
		case c == '>':
			return i
		}
	}
	return -1
}

func copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

// signAPK runs apksigner with the given keystore. apksigner is one of
// the tools graphics.gd already manages (see tools.go:AndroidPackageSigner).
func signAPK(apkPath, keystore string) error {
	return tooling.AndroidPackageSigner.Exec(
		"sign", "--ks", keystore,
		"--ks-key-alias", "androiddebugkey", "--ks-pass", "pass:android",
		apkPath,
	)
}

func debugKeystore() string {
	switch runtime.GOOS {
	case "linux":
		return filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "keystores", "debug.keystore")
	case "windows":
		return filepath.Join(os.Getenv("APPDATA"), "Godot", "keystores", "debug.keystore")
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", "keystores", "debug.keystore")
	}
	return ""
}

// releaseKeystore is currently identical to debug — graphics.gd doesn't
// yet have a separate signing flow for Meta Quest release builds.
// Sideloading + Meta Store both accept debug-signed APKs in dev mode,
// and a real release flow can be added when needed.
func releaseKeystore() string { return debugKeystore() }

// Compile-time guard: ensure MetaQuest satisfies the Builder interface
// declared in main.go.
var _ interface {
	Run(...string) error
	Build(...string) error
	BuildMain(...string) error
	Test(...string) error
} = MetaQuest{}
