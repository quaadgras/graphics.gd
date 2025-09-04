package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling/godot"
	"graphics.gd/cmd/gd/internal/tooling/golang"
	"graphics.gd/cmd/gd/internal/tooling/zig"
	"runtime.link/api/xray"
)

type Windows struct{}

func (Windows) Build(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "windows" || runtime.GOARCH != GOARCH {
		if err := zig.Assert(); err != nil {
			return xray.New(err)
		}
		switch GOARCH {
		case "amd64":
			if err := os.Setenv("CC", zig.Executable+" cc -target x86_64-windows-gnu"); err != nil {
				return xray.New(err)
			}
		case "arm64":
			if err := os.Setenv("CC", zig.Executable+" cc -target aarch64-windows-gnu"); err != nil {
				return xray.New(err)
			}
		default:
			return fmt.Errorf("gd build: cannot cross-compile windows %v on %v", GOARCH, runtime.GOOS)
		}
	}
	return golang.Build(args, "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("windows_%v.dll", GOARCH)))
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
	if err := godot.Run(export...); err != nil {
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
	return godot.Run()
}

func (Windows) Test(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "windows" || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd test: cannot run windows/%v tests on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	if err := golang.Test(args, "-c", "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("windows_%v.so", GOARCH))); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return godot.Test(args...)
}
