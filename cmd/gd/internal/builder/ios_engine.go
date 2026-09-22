package builder

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/shim"
	"graphics.gd/cmd/gd/internal/tooling"
)

// ios builds the engine's libgodot.a for iOS devices.
func (custom engine) ios() (string, error) {
	if runtime.GOOS == "android" {
		return "", errors.New("gd: projects with a custom engine (project.godot gd/engine/repository) cannot be exported on-device yet")
	}
	src, commit, err := custom.checkout()
	if err != nil {
		return "", err
	}
	library, err := custom.artifact(src, commit, "ios", "arm64", false, "libgodot.a")
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(library); err == nil {
		return library, nil
	}
	warnEngineVersion(src)
	sdk, err := setupIOSSDK(filepath.Join(gdpaths.Lib, "ios", "sdk"))
	if err != nil {
		return "", err
	}
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return "", err
	}
	llvm, err := tooling.LLVM.Lookup()
	if err != nil {
		return "", err
	}
	// The build looks for Xcode's toolchain, which gd stands in for with zig
	// and llvm. The engine's Swift (the SwiftUI entry point of the app) is
	// swapped for bundled/ios/app.m, as zig is not able to compile Swift.
	toolchain := filepath.Join(gdpaths.Lib, "ios", "toolchain")
	target := shim.Target{
		System:     []string{filepath.Join(sdk, "usr", "include")},
		Frameworks: []string{filepath.Join(sdk, "System", "Library", "Frameworks")},
		// The app links the C++ standard library of the system, which is
		// older than zig's headers for it. With availability annotations,
		// they hold back to what the deployment target is known to have.
		Flags: []string{"-U_LIBCPP_HAS_VENDOR_AVAILABILITY_ANNOTATIONS", "-D_LIBCPP_HAS_VENDOR_AVAILABILITY_ANNOTATIONS=1"},
	}
	config := shim.Config{
		Zig:   zig,
		LLVM:  llvm,
		Cache: filepath.Join(gdpaths.Lib, "ios", "cache"),
		Swift: filepath.Join(sdk, "app.m"),
		Targets: map[string]shim.Target{
			"aarch64-ios": target,
		},
	}
	bin := filepath.Join(toolchain, "usr", "bin")
	if err := shim.Install(bin, config, "clang", "clang++", "ar", "ranlib", "libtool", "swift-frontend"); err != nil {
		return "", err
	}
	fmt.Printf("gd: building engine %s for ios/arm64 (%s), this will take a while\n", commit[:12], engineTarget(false))
	if err := custom.scons(src, []string{"OSXCROSS_IOS=1"}, // lets the build know not to look for Xcode.
		"platform=ios", "arch=arm64", "target="+engineTarget(false),
		"APPLE_TOOLCHAIN_PATH="+toolchain, "APPLE_SDK_PATH="+sdk, "SWIFT_FRONTEND="+filepath.Join(bin, "swift-frontend"),
	); err != nil {
		return "", fmt.Errorf("gd: failed to build the custom engine: %w", err)
	}
	return library, copyFile(filepath.Join(src, "bin", "libgodot.ios."+engineTarget(false)+".arm64.a"), library)
}

// setupIOSSDK lays gd's bundled iOS SDK out the way Xcode does one, which
// is what the engine's build expects of APPLE_SDK_PATH. The C library's
// headers are zig's (which it only offers up itself for macOS targets).
func setupIOSSDK(dir string) (string, error) {
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	fmt.Fprintln(hash, dir, zig)
	err = fs.WalkDir(ios_sdk, "bundled/ios", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := ios_sdk.ReadFile(path)
		fmt.Fprintln(hash, path, len(data))
		hash.Write(data)
		return err
	})
	if err != nil {
		return "", err
	}
	stamp, sum := filepath.Join(dir, "stamp"), hex.EncodeToString(hash.Sum(nil))
	if existing, err := os.ReadFile(stamp); err == nil && string(existing) == sum {
		return dir, nil
	}
	if err := os.RemoveAll(dir); err != nil {
		return "", err
	}
	env, err := exec.Command(zig, "env").Output()
	if err != nil {
		return "", err
	}
	lib := filepath.Join(gdpaths.Bin, "lib")
	if match := regexp.MustCompile(`"?lib_dir"?\s*[=:]\s*"([^"]+)"`).FindSubmatch(env); match != nil {
		lib = string(match[1])
	}
	if err := project.CopyDir(filepath.Join(lib, "libc", "include", "any-macos-any"), filepath.Join(dir, "usr", "include")); err != nil {
		return "", fmt.Errorf("gd: cannot find zig's darwin headers: %w", err)
	}
	for from, to := range map[string]string{
		"bundled/ios/include":    filepath.Join(dir, "usr", "include"),
		"bundled/ios/lib":        filepath.Join(dir, "usr", "lib"),
		"bundled/ios/Frameworks": filepath.Join(dir, "System", "Library", "Frameworks"),
	} {
		if err := project.SetupFiles(ios_sdk, from, to); err != nil {
			return "", err
		}
	}
	app, err := ios_sdk.ReadFile("bundled/ios/app.m")
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(dir, "app.m"), app, 0644); err != nil {
		return "", err
	}
	return dir, os.WriteFile(stamp, []byte(sum), 0644)
}
