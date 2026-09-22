package builder

import (
	"fmt"
	"os"
	"path/filepath"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/shim"
	"graphics.gd/cmd/gd/internal/tooling"
)

// windows builds the engine as a static library for windows, with zig's
// mingw-w64 standing in for the MinGW toolchain the build expects.
func (custom engine) windows(GOARCH string) (string, error) {
	var arch, triple, prefix string
	switch GOARCH {
	case "amd64":
		arch, triple, prefix = "x86_64", "x86_64-windows-gnu", "x86_64-w64-mingw32-"
	case "arm64":
		arch, triple, prefix = "arm64", "aarch64-windows-gnu", "aarch64-w64-mingw32-"
	default:
		return "", fmt.Errorf("gd: a custom engine cannot be built for windows/%s yet", GOARCH)
	}
	src, commit, err := custom.checkout()
	if err != nil {
		return "", err
	}
	library, err := custom.artifact(src, commit, "windows", GOARCH, false, "libgodot.a")
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
	llvm, err := tooling.LLVM.Lookup()
	if err != nil {
		return "", err
	}
	bin := filepath.Join(gdpaths.Lib, "windows", "toolchain", triple)
	// mincore is an API set that forwards to kernel32 and ntdll, which zig's
	// mingw-w64 has no import library for: an empty one satisfies the link.
	// Nor has it sapi for arm64 (only x86), whose exports are the COM
	// registration of the speech API, which the engine reaches through COM
	// and the GUIDs of uuid instead.
	lib := filepath.Join(bin, "lib")
	if err := os.MkdirAll(lib, 0755); err != nil {
		return "", err
	}
	empty := []string{"libmincore.a"}
	if GOARCH == "arm64" {
		empty = append(empty, "libsapi.a")
	}
	for _, name := range empty {
		if _, err := os.Stat(filepath.Join(lib, name)); err != nil {
			if err := archive(zig, filepath.Join(lib, name), nil); err != nil {
				return "", err
			}
		}
	}
	config := shim.Config{
		Zig:     zig,
		LLVM:    llvm,
		Cache:   filepath.Join(gdpaths.Lib, "windows", "cache"),
		Default: triple,
		Targets: map[string]shim.Target{triple: {Lib: lib}},
	}
	var names []string
	for _, tool := range []string{"clang", "clang++", "ar", "ranlib", "windres", "dlltool"} {
		names = append(names, prefix+tool)
	}
	if err := shim.Install(bin, config, names...); err != nil {
		return "", err
	}
	fmt.Printf("gd: building engine %s for windows/%s (%s), this will take a while\n", commit[:12], GOARCH, engineTarget(false))
	// d3d12 needs a prebuilt Mesa that only the official builds are made
	// with, vulkan and opengl3 remain.
	if err := custom.scons(src, []string{"PATH=" + bin + string(os.PathListSeparator) + os.Getenv("PATH")},
		"platform=windows", "arch="+arch, "target="+engineTarget(false), "library_type=static_library",
		"use_mingw=yes", "use_llvm=yes", "use_static_cpp=yes", "d3d12=no",
	); err != nil {
		return "", fmt.Errorf("gd: failed to build the custom engine: %w", err)
	}
	if err := custom.combine(zig, src, library, ".windows."+engineTarget(false)+"."+arch); err != nil {
		return "", err
	}
	// The executable's version information and icon (which the export
	// rewrites, so a template has to have them) were compiled along with
	// the engine, into an object that defines no symbol: an archive would
	// never give it up, so it is kept beside the library to be linked.
	resources, _ := filepath.Glob(filepath.Join(src, "bin", "obj", "platform", "windows", "godot_res_template.windows*.o"))
	if len(resources) != 1 {
		return "", fmt.Errorf("gd: the engine's build left no resource object in %s", filepath.Join(src, "bin", "obj", "platform", "windows"))
	}
	return library, copyFile(resources[0], filepath.Join(filepath.Dir(library), "godot_res.o"))
}
