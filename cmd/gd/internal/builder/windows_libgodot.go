package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

// windowsSystemLibraries the engine links, as platform/windows/detect.py
// lists them for MinGW (mincore and the accesskit ones aside, which zig's
// mingw-w64 has no import libraries for).
var windowsSystemLibraries = []string{
	"mingw32", "dsound", "ole32", "d3d9", "winmm", "gdi32", "iphlpapi", "shell32", "shlwapi", "shcore",
	"wsock32", "ws2_32", "kernel32", "oleaut32", "sapi", "dinput8", "dxguid", "ksuser", "imm32", "bcrypt",
	"crypt32", "avrt", "uuid", "dwmapi", "dwrite", "wbemuuid", "ntdll", "hid", "user32", "psapi", "dbghelp",
}

// buildMainLibgodot exports the project as one executable: the Go program
// with the engine linked into it, the way the musl build does for linux.
func (windows Windows) buildMainLibgodot(GOARCH string, args ...string) error {
	custom, ok := customEngine()
	if !ok {
		return fmt.Errorf("gd: build/mode=libgodot needs an engine to link, set engine/repository or engine/strip_unused_classes in project.godot")
	}
	var arch, target string
	switch GOARCH {
	case "amd64":
		arch, target = "x86_64", "x86_64-windows-gnu"
	case "arm64":
		arch, target = "arm64", "aarch64-windows-gnu"
	default:
		return fmt.Errorf("gd build: cannot cross-compile windows %v", GOARCH)
	}
	libgodot, err := custom.windows(GOARCH)
	if err != nil {
		return xray.New(err)
	}
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if err := os.Setenv("CC", zig+" cc -target "+target); err != nil {
		return xray.New(err)
	}
	tags := mergeTags("windows", "archive")
	libgo := filepath.Join(project.GraphicsDirectory, fmt.Sprintf("windows_%v.a", GOARCH))
	args, err = archiver(args)
	if err != nil {
		return xray.New(err)
	}
	if err := tooling.Go.Action("build", args, append(fastcbFlags("windows", ""), "-tags", tags, "-buildmode=c-archive", "-o", libgo)...); err != nil {
		return xray.New(err)
	}
	// The final link is zig's, so the #cgo LDFLAGS of every linked package
	// have to be forwarded by hand (as the musl build does). The subsystem
	// is windows (no console), with the C runtime's entry for a main().
	out := filepath.Join(project.GraphicsDirectory, ".godot", "godot.windows.template_release."+arch+".exe")
	pck, err := pckSection()
	if err != nil {
		return xray.New(err)
	}
	link := []string{"c++", "-target", target, "-o", out, pck, filepath.Join(filepath.Dir(libgodot), "godot_res.o"), libgodot, libgo}
	cgoLDFLAGS, err := tooling.Go.Output("list", "-tags", tags, "-deps", "-f", "{{range .CgoLDFLAGS}}{{println .}}{{end}}", ".")
	if err != nil {
		return xray.New(err)
	}
	for _, flag := range strings.Split(cgoLDFLAGS, "\n") {
		if flag = strings.TrimSpace(flag); flag != "" {
			link = append(link, flag)
		}
	}
	link = append(link, tooling.CGOLDFlags()...)
	// the import libraries the engine's build was given in place of the
	// ones zig's mingw-w64 lacks (see windows_engine.go).
	link = append(link, "-L", filepath.Join(gdpaths.Lib, "windows", "toolchain", target, "lib"))
	for _, library := range windowsSystemLibraries {
		link = append(link, "-l"+library)
	}
	link = append(link, "-static", "-Wl,--subsystem,windows", "-Wl,-e,mainCRTStartup")
	if err := tooling.Zig.Exec(link...); err != nil {
		return xray.New(err)
	}
	preset := "Windows " + arch
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := setPresetOption(preset, "custom_template/release", ".godot/"+filepath.Base(out)); err != nil {
		return xray.New(err)
	}
	if err := setPresetOption(preset, "binary_format/architecture", arch); err != nil {
		return xray.New(err)
	}
	// the Go code is in the executable, nothing is to be loaded as an extension.
	restoreExtensions, err := ignoreEnabledExtensions()
	if err != nil {
		return xray.New(err)
	}
	defer restoreExtensions()
	exportPath, err := presetExportPath(preset)
	if err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll(filepath.Dir(exportPath), 0755); err != nil {
		return xray.New(err)
	}
	os.Remove(exportPath)
	if err := tooling.Godot.Exec("--headless", "--export-release", preset); err != nil {
		return xray.New(err)
	}
	if _, err := os.Stat(exportPath); err != nil {
		return fmt.Errorf("gd: the export of preset %q did not produce %s", preset, exportPath)
	}
	fmt.Println("gd: exported", exportPath)
	return nil
}

// presetExportPath returns where the named export preset exports to.
func presetExportPath(presetName string) (string, error) {
	presets, err := loadAndroidPresets() // reads every preset, android or not.
	if err != nil {
		return "", err
	}
	for _, preset := range presets {
		if preset.name == presetName {
			return filepath.Join(project.GraphicsDirectory, filepath.FromSlash(preset.exportPath)), nil
		}
	}
	return "", fmt.Errorf("preset %q not found in export_presets.cfg", presetName)
}
