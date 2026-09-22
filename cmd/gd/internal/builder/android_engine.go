package builder

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/shim"
	"graphics.gd/cmd/gd/internal/tooling"
)

// androidABIs maps the ABIs an export preset can enable, to the GOARCH
// and scons arch gd is able to build the engine for.
var androidABIs = map[string]struct{ GOARCH, scons string }{
	"arm64-v8a": {"arm64", "arm64"},
	"x86_64":    {"amd64", "x86_64"},
}

// android builds the engine's libgodot_android.so for the given ABI.
func (custom engine) android(abi string, debug bool) (string, error) {
	arch, ok := androidABIs[abi]
	if !ok {
		return "", fmt.Errorf("gd: a custom engine cannot be built for android %s yet, disable this architecture in the export preset", abi)
	}
	src, commit, err := custom.checkout()
	if err != nil {
		return "", err
	}
	library := custom.artifact(commit, "android", arch.GOARCH, debug, "libgodot_android.so")
	if _, err := os.Stat(library); err == nil {
		return library, nil
	}
	warnEngineVersion(src)
	triple, err := androidTriple(arch.GOARCH)
	if err != nil {
		return "", err
	}
	sdk, err := setupAndroidSDK(filepath.Join(gdpaths.Lib, "android", "sdk"), triple)
	if err != nil {
		return "", err
	}
	home, err := androidShimNDK(src, triple, sdk)
	if err != nil {
		return "", err
	}
	fmt.Printf("gd: building engine %s for android/%s (%s), this will take a while\n", commit[:12], arch.GOARCH, engineTarget(debug))
	// swappy is only distributed as binaries built against the NDK's C++
	// standard library, which cannot be linked with zig's.
	if err := scons(src, []string{"ANDROID_HOME=" + home},
		"platform=android", "arch="+arch.scons, "target="+engineTarget(debug), "swappy=no",
	); err != nil {
		return "", fmt.Errorf("gd: failed to build the custom engine: %w", err)
	}
	variant := "release"
	if debug {
		variant = "debug"
	}
	built := filepath.Join(src, "platform", "android", "java", "lib", "libs", variant, abi, "libgodot_android.so")
	return library, copyFile(built, library)
}

// androidShimNDK lays out just enough of an Android SDK + NDK for the
// engine's platform/android/detect.py to be satisfied, with zig standing
// in for the NDK's compilers, returning the directory for ANDROID_HOME.
func androidShimNDK(src, triple string, sdk androidSDK) (string, error) {
	detect, err := os.ReadFile(filepath.Join(src, "platform", "android", "detect.py"))
	if err != nil {
		return "", err
	}
	// the build looks inside of the exact NDK version that it asks for.
	match := regexp.MustCompile(`def get_ndk_version\(\):\s*return "([^"]+)"`).FindSubmatch(detect)
	if match == nil {
		return "", errors.New("gd: cannot determine which NDK version the engine expects (platform/android/detect.py)")
	}
	var host string
	switch runtime.GOOS {
	case "linux":
		host = "linux-x86_64"
	case "darwin":
		host = "darwin-x86_64"
	case "windows":
		host = "windows-x86_64"
	default:
		return "", fmt.Errorf("gd: building a custom engine for android is not supported on %s", runtime.GOOS)
	}
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return "", err
	}
	home := filepath.Join(gdpaths.Lib, "android", "home")
	prebuilt := filepath.Join(home, "ndk", string(match[1]), "toolchains", "llvm", "prebuilt", host)
	config := shim.Config{
		Zig:   zig,
		Cache: filepath.Join(gdpaths.Lib, "android", "cache"),
		Targets: map[string]shim.Target{
			triple: {LibC: sdk.LibC, Lib: sdk.Lib},
		},
	}
	if err := shim.Install(filepath.Join(prebuilt, "bin"), config, "clang", "clang++", "llvm-ar", "llvm-ranlib"); err != nil {
		return "", err
	}
	// The build copies the NDK's libc++_shared.so next to the engine. zig
	// links the C++ standard library statically, nothing loads this.
	placeholder := filepath.Join(prebuilt, "sysroot", "usr", "lib", triple, "libc++_shared.so")
	if err := os.MkdirAll(filepath.Dir(placeholder), 0755); err != nil {
		return "", err
	}
	return home, os.WriteFile(placeholder, nil, 0644)
}

// androidEngineTemplate is where the export template with the custom
// engine is kept, relative to the graphics directory.
func androidEngineTemplate(debug bool) string {
	return ".godot/godot.android." + engineTarget(debug) + ".apk"
}

// useAndroidEngine makes sure the named export preset will export with the
// project's custom engine (if any), building it as necessary.
func useAndroidEngine(presetName string, debug bool) error {
	option := "custom_template/release"
	if debug {
		option = "custom_template/debug"
	}
	custom, ok := customEngine()
	if !ok {
		// leave any template the user has set themselves alone.
		if current, _ := presetOption(presetName, option); current == androidEngineTemplate(debug) {
			return setPresetOption(presetName, option, "")
		}
		return nil
	}
	if runtime.GOOS == "android" {
		return errors.New("gd: projects with a custom engine (project.godot gd/engine/repository) cannot be exported on-device yet")
	}
	presets, err := loadAndroidPresets()
	if err != nil {
		return err
	}
	var abis []string
	for _, preset := range presets {
		if preset.name == presetName {
			for abi := range preset.archs {
				abis = append(abis, abi)
			}
		}
	}
	if len(abis) == 0 {
		return fmt.Errorf("gd: export preset %q does not enable any architectures", presetName)
	}
	slices.Sort(abis)
	libraries := map[string]string{}
	for _, abi := range abis {
		library, err := custom.android(abi, debug)
		if err != nil {
			return err
		}
		libraries["lib/"+abi+"/libgodot_android.so"] = library
	}
	templates, ok := gdpaths.ExportTemplates(tooling.Godot.InstalledVersion())
	if !ok {
		return fmt.Errorf("gd: no export templates on %s", runtime.GOOS)
	}
	stock := filepath.Join(templates, "android_release.apk")
	if debug {
		stock = filepath.Join(templates, "android_debug.apk")
	}
	template := filepath.Join(project.GraphicsDirectory, filepath.FromSlash(androidEngineTemplate(debug)))
	// only the engine is replaced, everything else about the template (the
	// java side, resources) is Godot's own: nothing about it is tied to the
	// native library beyond the JNI functions both were built from.
	stamp := stock
	for _, abi := range abis {
		stamp += "\n" + libraries["lib/"+abi+"/libgodot_android.so"]
	}
	if existing, err := os.ReadFile(template + ".txt"); err != nil || string(existing) != stamp {
		if err := replaceInZip(stock, template, libraries); err != nil {
			return err
		}
		if err := os.WriteFile(template+".txt", []byte(stamp), 0644); err != nil {
			return err
		}
	}
	return setPresetOption(presetName, option, androidEngineTemplate(debug))
}

// exportAndroid runs the Godot export of an android preset, mode is either
// --export-release or --export-debug.
func exportAndroid(mode, presetName string) error {
	if err := useAndroidEngine(presetName, mode == "--export-debug"); err != nil {
		return err
	}
	return tooling.Godot.Exec("--headless", mode, presetName)
}

// replaceInZip copies the zip at src to dst, with the named entries read
// from the given files instead. Native libraries that are not replaced
// are dropped, so that an export fails when it asks for an architecture
// which would otherwise silently ship with the stock engine.
func replaceInZip(src, dst string, replacements map[string]string) error {
	in, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer in.Close()
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	out, err := os.Create(dst + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()
	zw := zip.NewWriter(out)
	replaced, remaining := map[string]bool{}, len(replacements)
	for name := range replacements {
		replaced[filepath.ToSlash(filepath.Dir(name))] = true
	}
	for _, f := range in.File {
		if path, ok := replacements[f.Name]; ok {
			w, err := zw.CreateHeader(&zip.FileHeader{Name: f.Name, Method: zip.Deflate})
			if err != nil {
				return err
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(w, file)
			file.Close()
			if err != nil {
				return err
			}
			remaining--
			continue
		}
		if strings.HasPrefix(f.Name, "lib/") && !replaced[filepath.ToSlash(filepath.Dir(f.Name))] {
			continue
		}
		w, err := zw.CreateRaw(&f.FileHeader)
		if err != nil {
			return err
		}
		r, err := f.OpenRaw()
		if err != nil {
			return err
		}
		if _, err := io.Copy(w, r); err != nil {
			return err
		}
	}
	if remaining != 0 {
		return fmt.Errorf("gd: %s does not have every architecture that was asked for", filepath.Base(src))
	}
	if err := zw.Close(); err != nil {
		return err
	}
	if err := out.Close(); err != nil {
		return err
	}
	return os.Rename(dst+".tmp", dst)
}
