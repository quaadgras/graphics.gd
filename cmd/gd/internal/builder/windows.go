package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

type Windows struct{}

func (Windows) Build(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "windows" || runtime.GOARCH != GOARCH {
		zig, err := tooling.Zig.Lookup()
		if err != nil {
			return xray.New(err)
		}
		switch GOARCH {
		case "amd64":
			if err := os.Setenv("CC", zig+" cc -target x86_64-windows-gnu"); err != nil {
				return xray.New(err)
			}
		case "arm64":
			if err := os.Setenv("CC", zig+" cc -target aarch64-windows-gnu"); err != nil {
				return xray.New(err)
			}
		default:
			return fmt.Errorf("gd build: cannot cross-compile windows %v on %v", GOARCH, runtime.GOOS)
		}
	}
	return tooling.Go.Action("build", args, "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("windows_%v.dll", GOARCH)))
}

func (windows Windows) BuildMain(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if err := windows.Build(args...); err != nil {
		return xray.New(err)
	}
	var export []string
	switch GOARCH {
	case "amd64":
		export = []string{"--headless", "--export-release", "Windows x86_64"}
	case "arm64":
		export = []string{"--headless", "--export-release", "Windows arm64"}
	default:
		return fmt.Errorf("gd export: cannot export windows %v", GOARCH)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec(export...); err != nil {
		return xray.New(err)
	}
	return nil
}

func (windows Windows) Run(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "windows" || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd test: cannot run windows/%v executable on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	if err := windows.Build(args...); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return tooling.Godot.Exec()
}

func (Windows) Test(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "windows" || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd test: cannot run windows/%v tests on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	if err := tooling.Go.Action("build", args, "-c", "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("windows_%v.so", GOARCH))); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return tooling.Godot.Exec(args...)
}
