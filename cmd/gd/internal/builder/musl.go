package builder

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

var (
	//go:embed bundled/musl
	musl_sdk embed.FS
)

var built_musl bool

type Musl struct {
	lib string
	out string
}

func (musl Musl) Build(args ...string) (err error) {
	os.Remove(filepath.Join(project.GraphicsDirectory, "library.gdextension"))
	goos := os.Getenv("GOOS")
	os.Setenv("GOOS", "linux")
	defer os.Setenv("GOOS", goos)
	defer restoreCC()()
	if built_musl {
		return nil
	}
	defer func() {
		built_musl = true
	}()
	if !project.IncludesGo {
		return nil
	}
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if musl.lib == "" {
		libgodot, err := tooling.LibGodotEditor.LookupPlatform("musl", GOARCH)
		if err != nil {
			return xray.New(err)
		}
		musl.lib = libgodot
	}
	if musl.out == "" {
		musl.out = filepath.Join(project.GraphicsDirectory, "musl_"+GOARCH+".editor")
		// The static editor is the engine binary on hosts that have no other:
		// musl linux systems, and android (Termux), where it runs headless on
		// the Android kernel.
		hostRunsStaticEditor := runtime.GOOS == "android"
		if runtime.GOOS == "linux" {
			version, _ := tooling.ListDynamicDependencies.CombinedOutput("--version")
			hostRunsStaticEditor = strings.HasPrefix(version, "musl")
		}
		if hostRunsStaticEditor {
			defer func() {
				if err == nil {
					tooling.Godot.Path = musl.out
				}
			}()
		}
	}
	if err := project.SetupFiles(musl_sdk, "bundled/musl", filepath.Join(gdpaths.Lib, "musl")); err != nil {
		return xray.New(err)
	}
	if err := musl.patch(); err != nil {
		return xray.New(err)
	}
	GOROOT, err := tooling.Go.Output("env", "GOROOT")
	if err != nil {
		return xray.New(err)
	}
	overlay, err := muslOverlay(GOROOT)
	if err != nil {
		return xray.New(err)
	}
	var target string
	switch GOARCH {
	case "amd64":
		target = "x86_64-linux-musl"
		if err := os.Setenv("CC", zig+" cc -target x86_64-linux-musl -static"); err != nil {
			return xray.New(err)
		}
	case "arm64":
		target = "aarch64-linux-musl"
		if err := os.Setenv("CC", zig+" cc -target aarch64-linux-musl -static"); err != nil {
			return xray.New(err)
		}
	default:
		return fmt.Errorf("gd build: cannot cross-compile linux %v on %v", GOARCH, runtime.GOOS)
	}
	libgo := filepath.Join(project.GraphicsDirectory, fmt.Sprintf("musl_%v.a", GOARCH))
	if err := tooling.Go.Action("build", args, "-tags", muslTags(), "-buildmode=c-archive", "-overlay="+overlay, "-o", libgo); err != nil {
		return xray.New(err)
	}
	// Forward each linked package's #cgo LDFLAGS to the final link. `go build`
	// records them in the c-archive but does not apply them when an external tool
	// links it (zig here), so a package that statically links a C/C++ library would
	// otherwise fail with undefined symbols. `go list -deps` expands ${SRCDIR} to
	// absolute paths. -lc++ goes last so libc++ resolves any C++ symbols those
	// archives pull in.
	zigArgs := []string{"cc", "-target", target, musl.lib, libgo}
	cgoLDFLAGS, err := tooling.Go.Output("list", "-tags", muslTags(), "-deps", "-f", "{{range .CgoLDFLAGS}}{{println .}}{{end}}", ".")
	if err != nil {
		return xray.New(err)
	}
	for _, flag := range strings.Split(cgoLDFLAGS, "\n") {
		if flag = strings.TrimSpace(flag); flag != "" {
			zigArgs = append(zigArgs, flag)
		}
	}
	zigArgs = append(zigArgs, tooling.CGOLDFlags()...)
	zigArgs = append(zigArgs, "-lc++", "-o", musl.out)
	if err := tooling.Zig.Exec(zigArgs...); err != nil {
		return xray.New(err)
	}
	return nil
}

// restoreCC returns a func that restores CC to its value at the time of the
// call. The musl builder points CC at zig targeting musl for its own compile;
// when it runs as tooling for another target (the on-device android extension
// build relies on the ambient CC) that must not leak out.
func restoreCC() func() {
	cc, ok := os.LookupEnv("CC")
	return func() {
		if ok {
			os.Setenv("CC", cc)
		} else {
			os.Unsetenv("CC")
		}
	}
}

func (musl Musl) patch() error {
	my, err := user.Current()
	if err != nil {
		return xray.New(err)
	}
	HOME := my.HomeDir
	var GDPATH = os.Getenv("GDPATH")
	if GDPATH == "" {
		GDPATH = filepath.Join(HOME, "gd")
	}
	musl_malloc := filepath.Join(GDPATH, "bin", "lib", "libc", "musl", "src", "malloc", "mallocng", "malloc.c")
	file, err := os.ReadFile(musl_malloc)
	if err != nil {
		return xray.New(err)
	}
	file = bytes.Replace(file,
		[]byte(`struct malloc_context ctx = { 0 };`),
		[]byte(`struct malloc_context ctx = { .brk = -1 };`), 1)
	if err := os.WriteFile(musl_malloc, file, 0644); err != nil {
		return xray.New(err)
	}
	return nil
}

// ignoreEnabledExtensions disables every GDExtension for the duration of a musl export.
// A static musl binary can't dlopen a shared library (graphics.gd's loader borrows the
// host's dynamic loader), so any extension required at runtime must instead be statically
// linked and self-registered — e.g. GDExtensionManager.load_extension_from_function. We
// rename the project's *.gdextension files aside so Godot's import scan finds none: it then
// drops extension_list.cfg and the export bundles no extension loaders. The returned func
// restores everything for the editor and other-platform builds, so callers MUST defer it.
func ignoreEnabledExtensions() (restore func(), err error) {
	dir := project.GraphicsDirectory
	var found []string
	if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if d.IsDir() {
			if d.Name() == ".godot" { // huge import cache, never holds .gdextension files
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".gdextension") {
			found = append(found, path)
		}
		return nil
	}); err != nil {
		return func() {}, xray.New(err)
	}
	var disabled []string
	restore = func() {
		for _, path := range disabled {
			os.Rename(path+".disabled", path)
		}
	}
	for _, path := range found {
		if err := os.Rename(path, path+".disabled"); err != nil {
			restore()
			return func() {}, xray.New(err)
		}
		disabled = append(disabled, path)
	}
	// Godot regenerates extension_list.cfg from the (now-absent) .gdextension files — i.e.
	// removes it; keep the original so the editor's loaded-extension list survives the build.
	extList := filepath.Join(dir, ".godot", "extension_list.cfg")
	if saved, e := os.ReadFile(extList); e == nil {
		os.Remove(extList)
		inner := restore
		restore = func() {
			inner()
			os.WriteFile(extList, saved, 0644)
		}
	}
	return restore, nil
}

func (musl Musl) BuildMain(args ...string) error {
	os.Remove(filepath.Join(project.GraphicsDirectory, "library.gdextension"))
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	// The export preset's custom_template/release points at this arch-suffixed
	// name, so the two must agree (see graphics/export_presets.cfg).
	var arch string
	switch GOARCH {
	case "amd64":
		arch = "x86_64"
	case "arm64":
		arch = "arm64"
	default:
		return fmt.Errorf("gd export: cannot export musl %v", GOARCH)
	}
	var err error
	musl.out = filepath.Join(project.GraphicsDirectory, ".godot", "godot.musl.template_release."+arch)
	musl.lib, err = tooling.LibGodot.LookupPlatform("musl", GOARCH)
	if err != nil {
		return xray.New(err)
	}
	built_musl = false
	if err := musl.Build(args...); err != nil {
		return xray.New(err)
	}
	export := []string{"--headless", "--export-release", "Musl " + arch}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	restoreExtensions, err := ignoreEnabledExtensions()
	if err != nil {
		return xray.New(err)
	}
	defer restoreExtensions()
	if err := tooling.Godot.Exec(export...); err != nil {
		return xray.New(err)
	}
	return nil
}

func (musl Musl) Run(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "linux" || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd run: cannot run linux/%v executable on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	if err := musl.Build(args...); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return tooling.Godot.Exec(args...)
}

func (musl Musl) Test(args ...string) error {
	if built_musl {
		return nil
	}
	defer func() {
		built_musl = true
	}()
	os.Remove(filepath.Join(project.GraphicsDirectory, "library.gdextension"))
	goos := os.Getenv("GOOS")
	os.Setenv("GOOS", "linux")
	defer os.Setenv("GOOS", goos)
	defer restoreCC()()
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	// A static musl binary also runs on the Android kernel (Termux), headless.
	if (runtime.GOOS != "linux" && runtime.GOOS != "android") || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd test: cannot run linux/%v tests on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if err := project.SetupFiles(musl_sdk, "bundled/musl", filepath.Join(gdpaths.Lib, "musl")); err != nil {
		return xray.New(err)
	}
	if err := musl.patch(); err != nil {
		return xray.New(err)
	}
	GOROOT, err := tooling.Go.Output("env", "GOROOT")
	if err != nil {
		return xray.New(err)
	}
	overlay, err := muslOverlay(GOROOT)
	if err != nil {
		return xray.New(err)
	}
	var target string
	switch GOARCH {
	case "amd64":
		target = "x86_64-linux-musl"
		if err := os.Setenv("CC", zig+" cc -target x86_64-linux-musl"); err != nil {
			return xray.New(err)
		}
	case "arm64":
		target = "aarch64-linux-musl"
		if err := os.Setenv("CC", zig+" cc -target aarch64-linux-musl"); err != nil {
			return xray.New(err)
		}
	default:
		return fmt.Errorf("gd build: cannot cross-compile linux %v on %v", GOARCH, runtime.GOOS)
	}
	libgo := filepath.Join(project.GraphicsDirectory, fmt.Sprintf("musl_%v.a", GOARCH))
	if err := tooling.Go.Action("test", args, "-c", "-tags", muslTags(), "-buildmode=c-archive", "-overlay="+overlay, "-o", libgo); err != nil {
		return xray.New(err)
	}
	libgodot, err := tooling.LibGodotEditor.LookupPlatform("musl", GOARCH)
	if err != nil {
		return xray.New(err)
	}
	if err := tooling.Zig.Exec("c++", "-target", target, libgo, libgodot, "-o", filepath.Join(project.GraphicsDirectory, "musl_"+GOARCH+".editor")); err != nil {
		return xray.New(err)
	}
	tooling.Godot.Path = filepath.Join(project.GraphicsDirectory, "musl_"+GOARCH+".editor")

	// project.Setup imports the graphics directory after this callback returns,
	// which is too late for the suite we are about to run — now that there is an
	// editor to do it with, import anything still waiting for it.
	if err := project.Import(); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	args = append(args, "--headless")
	return tooling.Godot.Exec(args...)
}

// muslOverlay writes the GOROOT runtime-overlay file for musl builds and
// returns its path. The fastcb resident-callback cgocall.go replacement
// (bundled, see fastcb.go) is merged in by default, composing with the standard
// musl runtime1.go/os_linux.go overlays.
func muslOverlay(GOROOT string) (string, error) {
	replace := map[string]string{
		filepath.Join(GOROOT, "src", "runtime", "runtime1.go"): filepath.Join(gdpaths.Lib, "musl", "runtime1.go.overlay"),
		filepath.Join(GOROOT, "src", "runtime", "os_linux.go"): filepath.Join(gdpaths.Lib, "musl", "os_linux.go.overlay"),
	}
	if p := fastcbCgocall("musl"); p != "" {
		replace[filepath.Join(GOROOT, "src", "runtime", "cgocall.go")] = p
		// The bundled cgocall.go references the fused-crossing companions
		// (fastcbCallCFastPC), so the replacement and the additions are
		// all-or-nothing.
		companions, err := fastcbCallCFiles()
		if err != nil {
			return "", err
		}
		for name, path := range companions {
			replace[filepath.Join(GOROOT, "src", "runtime", name)] = path
		}
	}
	return writeOverlay("musl.json", replace)
}

// muslTags returns the build-tag list for musl builds; GD_EXTRA_TAGS appends
// additional tags (comma-separated), e.g. fastcboverlay.
func muslTags() string { return mergeTags("musl", "musl") }
