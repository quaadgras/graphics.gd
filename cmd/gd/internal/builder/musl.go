package builder

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

var (
	//go:embed bundled/musl
	musl_sdk embed.FS
)

var built bool

type Musl struct {
}

func (musl Musl) Build(args ...string) error {
	if built {
		return nil
	}
	defer func() {
		built = true
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
	var overlay = filepath.Join(gdpaths.Lib, "musl.json")
	if err := os.WriteFile(overlay, []byte(`{
		"Replace": {
			"`+filepath.Join(GOROOT, "src", "runtime", "runtime1.go")+`": "`+filepath.Join(gdpaths.Lib, "musl", "runtime1.go.overlay")+`",
			"`+filepath.Join(GOROOT, "src", "runtime", "os_linux.go")+`": "`+filepath.Join(gdpaths.Lib, "musl", "os_linux.go.overlay")+`"
		}
	}`), 0755); err != nil {
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
	if err := tooling.Go.Action("build", args, "-buildmode=c-archive", "-overlay="+overlay, "-o", libgo); err != nil {
		return xray.New(err)
	}
	libgodot, err := tooling.LibGodotEditor.LookupPlatform("musl", GOARCH)
	if err != nil {
		return xray.New(err)
	}
	if err := tooling.Zig.Exec("cc", "-target", target, libgodot, libgo, "-o", filepath.Join(project.GraphicsDirectory, "musl_"+GOARCH+".editor")); err != nil {
		return xray.New(err)
	}
	tooling.Godot.Path = filepath.Join(project.GraphicsDirectory, "musl_"+GOARCH+".editor")
	return nil
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

func (musl Musl) BuildMain(args ...string) error {
	return errors.New("not supported yet")
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
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "linux" || runtime.GOARCH != GOARCH {
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
	var overlay = filepath.Join(gdpaths.Lib, "musl.json")
	if err := os.WriteFile(overlay, []byte(`{
		"Replace": {
			"`+filepath.Join(GOROOT, "src", "runtime", "runtime1.go")+`": "`+filepath.Join(gdpaths.Lib, "musl", "runtime1.go.overlay")+`",
			"`+filepath.Join(GOROOT, "src", "runtime", "os_linux.go")+`": "`+filepath.Join(gdpaths.Lib, "musl", "os_linux.go.overlay")+`"
		}
	}`), 0755); err != nil {
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
	if err := tooling.Go.Action("test", args, "-c", "-buildmode=c-archive", "-overlay="+overlay, "-o", libgo); err != nil {
		return xray.New(err)
	}
	libgodot, err := tooling.LibGodotEditor.LookupPlatform("musl", GOARCH)
	if err != nil {
		return xray.New(err)
	}
	if err := tooling.Zig.Exec("cc", "-target", target, libgodot, libgo, "-o", filepath.Join(project.GraphicsDirectory, "musl_"+GOARCH+".editor")); err != nil {
		return xray.New(err)
	}
	tooling.Godot.Path = filepath.Join(project.GraphicsDirectory, "musl_"+GOARCH+".editor")

	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	args = append(args, "--headless")
	return tooling.Godot.Exec(args...)
}
