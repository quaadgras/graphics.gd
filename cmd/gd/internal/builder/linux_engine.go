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
	return library, custom.combine(zig, src, library, ".linuxbsd."+engineTarget(false)+"."+arch)
}

// combine the static libraries that a build left in src/bin into one at
// library. The build leaves one per part of the engine under bin/obj, and
// the platform's own in bin, named with the given infix: unless that one
// is a combination of the parts already (which is what the graphics.gd
// fork's linux build leaves there, by the same name).
func (custom engine) combine(zig, src, library, infix string) error {
	var parts, platform []string
	if err := filepath.WalkDir(filepath.Join(src, "bin"), func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".a") || !strings.Contains(path, infix) {
			return err
		}
		if filepath.Dir(path) == filepath.Join(src, "bin") {
			platform = append(platform, path)
		} else {
			parts = append(parts, path)
		}
		return nil
	}); err != nil {
		return err
	}
	covered := map[string]bool{}
	for _, part := range parts {
		members, err := members(zig, part)
		if err != nil {
			return err
		}
		for _, member := range members {
			covered[member] = true
		}
	}
	for _, archive := range platform {
		members, err := members(zig, archive)
		if err != nil {
			return err
		}
		for _, member := range members {
			if !covered[member] {
				parts = append(parts, archive)
				break
			}
		}
	}
	if len(parts) == 0 {
		return fmt.Errorf("gd: the engine's build left no static libraries in %s", filepath.Join(src, "bin"))
	}
	slices.Sort(parts)
	return archive(zig, library, parts)
}

// members of a static library, by file name.
func members(zig, library string) ([]string, error) {
	out, err := exec.Command(zig, "ar", "t", library).Output()
	if err != nil {
		return nil, fmt.Errorf("listing %s: %w", library, err)
	}
	var names []string
	for line := range strings.SplitSeq(strings.TrimSpace(string(out)), "\n") {
		names = append(names, filepath.Base(line))
	}
	return names, nil
}

// archive the given static libraries into one at path, with an ar script.
func archive(zig, path string, libraries []string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	var script strings.Builder
	fmt.Fprintf(&script, "CREATE %s\n", path)
	for _, library := range libraries {
		// A thin archive only refers to its objects (the windows build makes
		// them), which have to be added themselves for the result to stand
		// on its own.
		if header, err := os.ReadFile(library); err == nil && strings.HasPrefix(string(header[:min(8, len(header))]), "!<thin>") {
			members, err := exec.Command(zig, "ar", "t", library).Output()
			if err != nil {
				return fmt.Errorf("listing %s: %w", library, err)
			}
			for member := range strings.SplitSeq(strings.TrimSpace(string(members)), "\n") {
				if !filepath.IsAbs(member) {
					member = filepath.Join(filepath.Dir(library), member)
				}
				fmt.Fprintf(&script, "ADDMOD %s\n", member)
			}
			continue
		}
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
