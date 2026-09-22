package builder

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/shim"
	"graphics.gd/cmd/gd/internal/tooling"
)

// macos builds the engine as a static library for macOS. The build only
// proceeds on another OS with OSXCROSS_ROOT set, taking its tools from
// there by name, which gd's shim takes on.
func (custom engine) macos(GOARCH string) (string, error) {
	var arch, triple, prefix string
	switch GOARCH {
	case "arm64":
		arch, triple, prefix = "arm64", "aarch64-macos", "arm64-apple-darwin16-"
	case "amd64":
		arch, triple, prefix = "x86_64", "x86_64-macos", "x86_64-apple-darwin16-"
	default:
		return "", fmt.Errorf("gd: a custom engine cannot be built for macos/%s yet", GOARCH)
	}
	src, commit, err := custom.checkout()
	if err != nil {
		return "", err
	}
	library, err := custom.artifact(src, commit, "darwin", GOARCH, false, "libgodot.a")
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(library); err == nil {
		return library, nil
	}
	warnEngineVersion(src)
	sdk, err := setupMacOSSDK(filepath.Join(gdpaths.Lib, "macos", "sdk"))
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
	root := filepath.Join(gdpaths.Lib, "macos", "osxcross")
	config := shim.Config{
		Zig:   zig,
		LLVM:  llvm,
		Cache: filepath.Join(gdpaths.Lib, "macos", "cache"),
		Targets: map[string]shim.Target{
			triple: {
				System:     []string{filepath.Join(sdk, "usr", "include")},
				Frameworks: []string{filepath.Join(sdk, "System", "Library", "Frameworks")},
				Flags:      []string{"-U_LIBCPP_HAS_VENDOR_AVAILABILITY_ANNOTATIONS", "-D_LIBCPP_HAS_VENDOR_AVAILABILITY_ANNOTATIONS=1"},
			},
		},
	}
	var names []string
	for _, tool := range []string{"cc", "c++", "ar", "ranlib", "as"} {
		names = append(names, prefix+tool)
	}
	if err := shim.Install(filepath.Join(root, "target", "bin"), config, names...); err != nil {
		return "", err
	}
	fmt.Printf("gd: building engine %s for macos/%s (%s), this will take a while\n", commit[:12], GOARCH, engineTarget(false))
	// Vulkan on macOS is MoltenVK, a prebuilt SDK the build would have to be
	// pointed at (vulkan_sdk_path, through engine/options), Metal remains.
	if err := custom.scons(src, []string{"OSXCROSS_ROOT=" + root},
		"platform=macos", "arch="+arch, "target="+engineTarget(false), "library_type=static_library", "osxcross_sdk=darwin16", "vulkan=no",
	); err != nil {
		return "", fmt.Errorf("gd: failed to build the custom engine: %w", err)
	}
	return library, custom.combineDarwin(llvm, src, library, ".macos."+engineTarget(false)+"."+arch)
}

// combineDarwin is [engine.combine] for Mach-O static libraries, which
// Apple's libtool (llvm's likeness of it) archives with the symbol table
// their linker expects.
func (custom engine) combineDarwin(llvm, src, library, infix string) error {
	var parts []string
	if err := filepath.WalkDir(filepath.Join(src, "bin"), func(path string, d os.DirEntry, err error) error {
		if err == nil && !d.IsDir() && strings.HasSuffix(path, ".a") && strings.Contains(path, infix) {
			parts = append(parts, path)
		}
		return err
	}); err != nil {
		return err
	}
	if len(parts) == 0 {
		return fmt.Errorf("gd: the engine's build left no static libraries in %s", filepath.Join(src, "bin"))
	}
	if err := os.MkdirAll(filepath.Dir(library), 0755); err != nil {
		return err
	}
	cmd := exec.Command(llvm, append([]string{"libtool-darwin", "-static", "-o", library}, parts...)...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("archiving the engine: %w\n%s", err, output)
	}
	return nil
}

// setupMacOSSDK lays gd's bundled macOS SDK out the way Xcode does one.
// Most of the frameworks are those of the iOS SDK, whose declarations
// are the same on macOS, with the macOS ones on top.
func setupMacOSSDK(dir string) (string, error) {
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	fmt.Fprintln(hash, dir, zig)
	for _, bundle := range []struct {
		fs   fs.FS
		root string
	}{{ios_sdk, "bundled/ios"}, {macos_sdk, "bundled/macos"}} {
		err = fs.WalkDir(bundle.fs, bundle.root, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			data, err := fs.ReadFile(bundle.fs, path)
			fmt.Fprintln(hash, path, len(data))
			hash.Write(data)
			return err
		})
		if err != nil {
			return "", err
		}
	}
	stamp, sum := filepath.Join(dir, "stamp"), hex.EncodeToString(hash.Sum(nil))
	if existing, err := os.ReadFile(stamp); err == nil && string(existing) == sum {
		return dir, nil
	}
	if err := os.RemoveAll(dir); err != nil {
		return "", err
	}
	if err := project.CopyDir(filepath.Join(zigLibDir(zig), "libc", "include", "any-macos-any"), filepath.Join(dir, "usr", "include")); err != nil {
		return "", fmt.Errorf("gd: cannot find zig's darwin headers: %w", err)
	}
	// the C library stubs of the iOS SDK declare their macOS targets too.
	for _, from := range []string{"bundled/ios/include", "bundled/ios/lib"} {
		if err := project.SetupFiles(ios_sdk, from, filepath.Join(dir, "usr", strings.TrimPrefix(from, "bundled/ios/"))); err != nil {
			return "", err
		}
	}
	// the iOS frameworks' headers, other than UIKit's, apply to macOS too.
	frameworks, err := fs.ReadDir(ios_sdk, "bundled/ios/Frameworks")
	if err != nil {
		return "", err
	}
	for _, framework := range frameworks {
		headers := "bundled/ios/Frameworks/" + framework.Name() + "/Headers"
		if framework.Name() == "UIKit.framework" {
			continue
		}
		if _, err := fs.Stat(ios_sdk, headers); err != nil {
			continue
		}
		if err := project.SetupFiles(ios_sdk, headers, filepath.Join(dir, "System", "Library", "Frameworks", framework.Name(), "Headers")); err != nil {
			return "", err
		}
	}
	for from, to := range map[string]string{
		"bundled/macos/include/mach-o": filepath.Join(dir, "usr", "include", "mach-o"),
		"bundled/macos/lib":            filepath.Join(dir, "usr", "lib"),
		"bundled/macos/Frameworks":     filepath.Join(dir, "System", "Library", "Frameworks"),
	} { // the rest of bundled/macos/include is the Go build's, superseded by the frameworks here.
		if err := project.SetupFiles(macos_sdk, from, to); err != nil {
			return "", err
		}
	}
	return dir, os.WriteFile(stamp, []byte(sum), 0644)
}

// zigLibDir returns where zig keeps its bundled libraries.
func zigLibDir(zig string) string {
	lib := filepath.Join(gdpaths.Bin, "lib")
	if env, err := exec.Command(zig, "env").Output(); err == nil {
		if _, after, ok := strings.Cut(string(env), "lib_dir"); ok {
			if _, value, ok := strings.Cut(after, "\""); ok {
				if value, _, ok := strings.Cut(value, "\""); ok {
					lib = value
				}
			}
		}
	}
	return lib
}
