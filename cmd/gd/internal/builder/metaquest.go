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
// Dependencies (all Apache-2.0) are PRE-BUILT and embedded into the
// gd binary via cmd/gd/internal/builder/bundled/metaquest/:
//
//   - libopenxr_loader.so       (from org.khronos.openxr:openxr_loader_for_android)
//   - libgodotopenxrvendors.so  (from org.godotengine:godot-openxr-vendors-meta)
//   - classes2.dex              (vendor classes.jar D8-compiled offline)
//
// The user build needs no Maven access, no JDK, and no R8 — every
// asset ships inside the gd binary. Bumping the vendor version is
// a one-time regen of the bundled assets on a dev machine.
package builder

import (
	"archive/zip"
	"bytes"
	_ "embed"
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
	if err := tooling.Godot.Exec("--headless", "--export-release", "Meta Quest"); err != nil {
		return xray.New(err)
	}
	apk := filepath.Join(project.ReleasesDirectory, "metaquest", project.Name+".apk")
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
	if err := os.MkdirAll(filepath.Join(project.ReleasesDirectory, "metaquest"), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-debug", "Meta Quest"); err != nil {
		return xray.New(err)
	}
	apk := filepath.Join(project.ReleasesDirectory, "metaquest", project.Name+".apk")
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

// Pre-compiled Meta Quest assets. We could fetch the AARs from
// Maven Central and run R8/D8 at user-build time, but R8 8.x needs
// Java 11+ and dragging a JDK into the toolchain just to dex a few
// hundred classes for each user build was a lot. Instead we ship
// the artifacts pre-built:
//
//   - classes2.dex: vendor/classes.jar from
//     org.godotengine:godot-openxr-vendors-meta compiled with
//     `d8 --release --min-api 29` (any version of d8 / Java works
//     since the input was Java 8 bytecode).
//   - lib/arm64-v8a/libopenxr_loader.so: from
//     org.khronos.openxr:openxr_loader_for_android.
//   - lib/arm64-v8a/libgodotopenxrvendors.so: Meta vendor native
//     hook, from the same vendor AAR.
//
// All three are Apache-2.0. Regenerate by running the script at
// cmd/gd/internal/builder/bundled/metaquest/README.md (TODO add it)
// when the vendor releases a new version.
//
//go:embed bundled/metaquest/classes2.dex
var metaQuestDex []byte

//go:embed bundled/metaquest/lib/arm64-v8a/libopenxr_loader.so
var metaQuestLoaderSo []byte

//go:embed bundled/metaquest/lib/arm64-v8a/libgodotopenxrvendors.so
var metaQuestVendorSo []byte

// injectMetaQuest is the keystone: take the APK Godot just exported,
// patch its manifest with the Quest features / permissions / intent
// filters via apktool, drop in the pre-built native libs + dex, zip
// it back up. After this returns the APK still needs to be re-signed.
//
// No external downloads, no Java invocation — every Quest-specific
// asset is embedded directly in the gd binary.
func injectMetaQuest(apkPath string) error {
	work, err := os.MkdirTemp("", "metaquest-")
	if err != nil {
		return xray.New(err)
	}
	defer os.RemoveAll(work)

	// Decompile the APK so we can rewrite the binary AndroidManifest.xml
	// through apktool's xml-tools, then repack. We don't touch the
	// existing dex — apktool will re-bundle it as-is.
	decompiled := filepath.Join(work, "decompiled")
	if err := tooling.AndroidPackageKitTool.Exec("d", apkPath, "-s", "-o", decompiled, "-f"); err != nil {
		return xray.New(err)
	}
	// Godot's APK ships a themed_icon mipmap entry that aapt2 can't
	// recompile (it decodes as a non-reference value). Drop the
	// auto-generated mipmaps.xml + the matching public.xml row so
	// the resource recompile during `apktool b` succeeds. Same
	// cleanup BuildMain does for the AAB path in android.go.
	if err := os.Remove(filepath.Join(decompiled, "res", "values-anydpi-v26", "mipmaps.xml")); err != nil && !os.IsNotExist(err) {
		return xray.New(err)
	}
	publicPath := filepath.Join(decompiled, "res", "values", "public.xml")
	if public, err := os.ReadFile(publicPath); err == nil {
		public = bytes.Replace(public, []byte(`<public type="mipmap" name="themed_icon" id="0x7f0a0004" />`), nil, 1)
		if err := os.WriteFile(publicPath, public, 0644); err != nil {
			return xray.New(err)
		}
	}

	// apktool 2.12 emits a literal `\1` placeholder for some integer
	// manifest attributes it couldn't decode, and the decompiled
	// manifest also carries the placeholder package "com.godot.game"
	// instead of the project's real one. aapt2's link rejects both
	// when re-packing. Resolve the original via `aapt2 dump
	// packagename` (same fixups BuildMain does for the AAB path).
	originalPackageName, err := tooling.AndroidAssetPackagingTool.Output("dump", "packagename", apkPath)
	if err != nil {
		return xray.New(err)
	}

	manifestPath := filepath.Join(decompiled, "AndroidManifest.xml")
	patched, err := patchManifestForQuest(manifestPath)
	if err != nil {
		return xray.New(err)
	}
	patched = bytes.Replace(patched, []byte(`package="com.godot.game"`), []byte(`package="`+originalPackageName+`"`), 1)
	patched = bytes.Replace(patched, []byte(`android:name="com.godot.game"`), []byte(`android:name="`+originalPackageName+`"`), 1)
	patched = bytes.Replace(patched, []byte(`android:version="\1"`), []byte(`android:version="1"`), 1)
	if err := os.WriteFile(manifestPath, patched, 0644); err != nil {
		return xray.New(err)
	}

	// Write embedded native libs into the lib tree.
	if err := os.MkdirAll(filepath.Join(decompiled, "lib", "arm64-v8a"), 0755); err != nil {
		return xray.New(err)
	}
	for name, data := range map[string][]byte{
		"libopenxr_loader.so":      metaQuestLoaderSo,
		"libgodotopenxrvendors.so": metaQuestVendorSo,
	} {
		dst := filepath.Join(decompiled, "lib", "arm64-v8a", name)
		if err := os.WriteFile(dst, data, 0644); err != nil {
			return xray.New(err)
		}
	}

	// Apktool puts secondary dex files at the top level. Slot in our
	// pre-built classes2.dex so the Meta plugin gets loaded by the
	// Android runtime alongside Godot's classes.dex.
	if err := os.WriteFile(filepath.Join(decompiled, "classes2.dex"), metaQuestDex, 0644); err != nil {
		return xray.New(err)
	}

	// Rebuild. Apktool's `b` re-encodes the manifest to binary AXML
	// and bundles everything in decompiled/ back into an APK.
	//
	// apktool ships its own aapt2 baked into the jar but only as a
	// glibc x86_64 binary; on musl (Void, Alpine, …) the extract +
	// exec fails and it falls back to $PATH, which doesn't have one
	// either. Point it at the aapt2 graphics.gd already manages so
	// the apktool roundtrip works on every platform we support.
	aapt2, err := tooling.AndroidAssetPackagingTool.Lookup()
	if err != nil {
		return xray.New(err)
	}
	repacked := apkPath + ".meta-unsigned.apk"
	if err := tooling.AndroidPackageKitTool.Exec("b", decompiled, "--aapt", aapt2, "-o", repacked); err != nil {
		return xray.New(err)
	}
	// Android R+ (API 30+) refuses to install APKs whose
	// resources.arsc is compressed or not 4-byte aligned within
	// the zip, and historically wants .so files page-aligned for
	// mmap. apktool's output is Deflate-compressed and unaligned,
	// so re-pack with the alignment fixed before signing.
	aligned := apkPath + ".meta-aligned.apk"
	if err := zipalignAPK(repacked, aligned); err != nil {
		return xray.New(err)
	}
	_ = os.Remove(repacked)
	if err := os.Rename(aligned, apkPath); err != nil {
		return xray.New(err)
	}
	return nil
}

// zipalignAPK rewrites src to dst with two adjustments that
// Android's installer demands for R+ targets:
//
//   - resources.arsc is re-stored uncompressed (Method = Store)
//     and placed at a 4-byte aligned data offset.
//   - .so files (already typically Stored so they can be mmap'd
//     at runtime) are aligned to 4096-byte (one-page) boundaries.
//
// Other entries pass through with their original method intact;
// Deflate-compressed entries don't need alignment because the
// runtime doesn't mmap them, so we leave them alone.
func zipalignAPK(src, dst string) error {
	in, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("zipalign: open %s: %w", src, err)
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("zipalign: create %s: %w", dst, err)
	}
	defer out.Close()
	// Wrap with a counter so we know the current write offset when
	// computing per-entry alignment padding.
	counter := &countingWriter{w: out}
	zw := zip.NewWriter(counter)
	for _, f := range in.File {
		// archive/zip's Writer wraps the supplied io.Writer in a
		// bufio.Writer, so counter.n only reflects bytes that have
		// already been flushed from that buffer. Flush before
		// computing each entry's LFH start offset, otherwise our
		// padding math is based on a stale offset and the data
		// lands at the wrong alignment.
		if err := zw.Flush(); err != nil {
			return fmt.Errorf("zipalign: flush: %w", err)
		}
		var align int
		method := f.Method
		switch {
		case f.Name == "resources.arsc":
			method = zip.Store
			align = 4
		case strings.HasSuffix(f.Name, ".so") && method == zip.Store:
			align = 4096
		}
		fh := &zip.FileHeader{
			Name:           f.Name,
			Method:         method,
			CreatorVersion: f.CreatorVersion,
			ExternalAttrs:  f.ExternalAttrs,
			NonUTF8:        f.NonUTF8,
			Flags:          f.Flags &^ 0x8, // clear data-descriptor bit; we know sizes upfront
			ReaderVersion:  f.ReaderVersion,
			Comment:        f.Comment,
			// Deliberately NOT setting Modified — archive/zip's
			// Writer appends a 9-byte "Extended Timestamp" Extra
			// (tag 0x5455 "UT") whenever Modified is non-zero,
			// which silently shifts the data offset past whatever
			// alignment padding we computed. APKs don't use entry
			// modtimes for anything user-visible, so we drop them.
		}
		if align > 0 {
			// Local file header is fixed 30 bytes, then name, then extra,
			// then data. Pad Extra so the data offset lands on `align`.
			dataOffset := counter.n + 30 + int64(len(f.Name))
			if r := dataOffset % int64(align); r != 0 {
				fh.Extra = make([]byte, int64(align)-r)
			}
		}
		w, err := zw.CreateHeader(fh)
		if err != nil {
			return fmt.Errorf("zipalign: header %s: %w", f.Name, err)
		}
		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("zipalign: open entry %s: %w", f.Name, err)
		}
		if _, err := io.Copy(w, rc); err != nil {
			rc.Close()
			return fmt.Errorf("zipalign: copy %s: %w", f.Name, err)
		}
		rc.Close()
	}
	return zw.Close()
}

type countingWriter struct {
	w io.Writer
	n int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	cw.n += int64(n)
	return n, err
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
