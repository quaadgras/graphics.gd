package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/shim"
	"graphics.gd/cmd/gd/internal/tooling"
)

// linux builds the engine as a static library for linux, against musl so
// that the executable gd links it into runs on any distribution.
func (custom engine) linux(GOARCH string) (string, error) {
	var arch, triple string
	switch GOARCH {
	case "amd64":
		arch, triple = "x86_64", "x86_64-linux-musl"
	case "arm64":
		arch, triple = "arm64", "aarch64-linux-musl"
	default:
		return "", fmt.Errorf("gd: a custom engine cannot be built for linux/%s yet", GOARCH)
	}
	src, commit, err := custom.checkout()
	if err != nil {
		return "", err
	}
	library, err := custom.artifact(src, commit, "linux", GOARCH, false, "libgodot.a")
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(library); err == nil {
		return library, nil
	}
	warnEngineVersion(src)
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return "", err
	}
	// The build takes its compiler from the PATH, where gd puts zig by the
	// names the build knows clang by (use_llvm), told what it is for by the
	// configuration, as nothing about a linux build says so.
	bin := filepath.Join(gdpaths.Lib, "linux", "toolchain", triple)
	config := shim.Config{
		Zig:     zig,
		Cache:   filepath.Join(gdpaths.Lib, "linux", "cache"),
		Default: triple,
		Targets: map[string]shim.Target{triple: {}},
	}
	if err := shim.Install(bin, config, "clang", "clang++", "ar", "ranlib"); err != nil {
		return "", err
	}
	fmt.Printf("gd: building engine %s for linux/%s (%s), this will take a while\n", commit[:12], GOARCH, engineTarget(false))
	if err := custom.scons(src, []string{"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH")},
		"platform=linuxbsd", "arch="+arch, "target="+engineTarget(false), "library_type=static_library",
		"use_llvm=yes", "use_static_cpp=yes", "execinfo=no",
	); err != nil {
		return "", fmt.Errorf("gd: failed to build the custom engine: %w", err)
	}
	// Some engines archive everything into one libgodot (the graphics.gd
	// fork does), otherwise the build leaves a static library per part of
	// the engine which gd archives together.
	pattern := ".linuxbsd." + engineTarget(false) + "." + arch
	if combined, _ := filepath.Glob(filepath.Join(src, "bin", "libgodot"+pattern+"*.a")); len(combined) == 1 {
		return library, copyFile(combined[0], library)
	}
	var parts []string
	if err := filepath.WalkDir(filepath.Join(src, "bin"), func(path string, d os.DirEntry, err error) error {
		if err == nil && !d.IsDir() && strings.HasSuffix(path, ".a") && strings.Contains(path, pattern) {
			parts = append(parts, path)
		}
		return err
	}); err != nil {
		return "", err
	}
	if len(parts) == 0 {
		return "", fmt.Errorf("gd: the engine's build left no static libraries in %s", filepath.Join(src, "bin"))
	}
	slices.Sort(parts)
	return library, archive(zig, library, parts)
}

// archive the given static libraries into one at path, with an ar script.
func archive(zig, path string, libraries []string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	var script strings.Builder
	fmt.Fprintf(&script, "CREATE %s\n", path)
	for _, library := range libraries {
		fmt.Fprintf(&script, "ADDLIB %s\n", library)
	}
	script.WriteString("SAVE\nEND\n")
	cmd := exec.Command(zig, "ar", "-M")
	cmd.Stdin = strings.NewReader(script.String())
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("archiving the engine: %w\n%s", err, output)
	}
	return nil
}

// buildModeError for the platforms libgodot mode is yet to reach.
func buildModeError(GOOS string) error {
	return fmt.Errorf("gd: build/mode=libgodot is only supported for linux so far, not %s", GOOS)
}
