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

type Linux struct{}

func (Linux) Build(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "linux" || runtime.GOARCH != GOARCH {
		if err := zig.Assert(); err != nil {
			return xray.New(err)
		}
		switch GOARCH {
		case "amd64":
			if err := os.Setenv("CC", zig.Executable+" cc -target x86_64-linux-gnu"); err != nil {
				return xray.New(err)
			}
		case "arm64":
			if err := os.Setenv("CC", zig.Executable+" cc -target aarch64-linux-gnu"); err != nil {
				return xray.New(err)
			}
		default:
			return fmt.Errorf("gd build: cannot cross-compile linux %v on %v", GOARCH, runtime.GOOS)
		}
	}
	return golang.Build(args, "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("linux_%v.so", GOARCH)))
}

func (linux Linux) BuildMain(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if err := linux.Build(args...); err != nil {
		return xray.New(err)
	}
	var export []string
	switch GOARCH {
	case "amd64":
		export = []string{"--headless", "--export-release", "Linux x86_64"}
	case "arm64":
		export = []string{"--headless", "--export-release", "Linux arm64"}
	default:
		return fmt.Errorf("gd export: cannot export linux %v", GOARCH)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := godot.Run(export...); err != nil {
		return xray.New(err)
	}
	return nil
}

func (linux Linux) Run(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "linux" || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd run: cannot run linux/%v executable on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	if err := linux.Build(args...); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return godot.Run()
}

func (Linux) Test(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "linux" || runtime.GOARCH != GOARCH {
		return fmt.Errorf("gd test: cannot run linux/%v tests on %v/%v", GOARCH, runtime.GOOS, runtime.GOARCH)
	}
	if err := golang.Test(args, "-c", "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("linux_%v.so", GOARCH))); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return godot.Test(args...)
}
