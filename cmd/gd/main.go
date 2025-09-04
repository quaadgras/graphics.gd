// The 'gd' command is designed as a drop-in replacement of the 'go' command when working
// with Godot-based projects. It will automatically download and install the supported
// version of Godot and ensure all go commands behave as expected (go build, go run, etc).
//
// The 'gd' command assumes that the Go module lives at the root of the project, an empty
// godot project will be created under a 'graphics' directory, this is where the user can
// keep the graphical representation of their project and manage their assets. Running the
// command without any command line arguments will launch the Godot editor for managing
// the assets in this directory.
package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"graphics.gd/cmd/gd/internal/builder"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling/godot"
	"graphics.gd/cmd/gd/internal/tooling/golang"
	"graphics.gd/internal/docgen"

	"runtime.link/api/xray"
)

func main() {
	if err := gd(os.Args...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type Builder interface {
	Run(...string) error       // go run
	Build(...string) error     // go build -buildmode=c-shared
	BuildMain(...string) error // go build
	Test(...string) error      // go test
}

func builderFor(goos string) Builder {
	switch goos {
	case "linux", "ubuntu", "arch", "debian", "nix":
		os.Setenv("GOOS", "linux")
		return builder.Linux{}
	case "windows", "win":
		os.Setenv("GOOS", "windows")
		return builder.Windows{}
	case "darwin", "macos":
		os.Setenv("GOOS", "darwin")
		return builder.MacOS{}
	case "ios":
		os.Setenv("GOOS", "darwin")
		return builder.IOS{}
	case "android":
		os.Setenv("GOOS", "android")
		return builder.Android{}
	case "browser", "js", "web", "wasm":
		os.Setenv("GOOS", "js")
		return builder.Browser{}
	default:
		fmt.Fprint(os.Stderr, "gd: unsupported GOOS '"+goos+"'\n")
		os.Exit(1)
		return nil
	}
}

func gd(args ...string) error {
	var GOARCH = runtime.GOARCH
	var GOOS = runtime.GOOS
	if goos := os.Getenv("GOOS"); goos != "" {
		GOOS = goos
	}
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if GOARCH != "amd64" && GOARCH != "arm64" && GOARCH != "wasm" {
		return errors.New("gd requires an amd64, wasm, or arm64 GOARCH")
	}
	if err := godot.Assert(); err != nil {
		return xray.New(err)
	}
	if err := project.Setup(); err != nil {
		return xray.New(err)
	}
	if err := docgen.Process(project.Directory); err != nil {
		return xray.New(err)
	}
	var platform = builderFor(GOOS)
	if goos := os.Getenv("GOOS"); goos != "" {
		GOOS = goos
	}
	if GOOS != "js" {
		if err := os.Setenv("CGO_ENABLED", "1"); err != nil {
			return xray.New(err)
		}
	}
	if zig, _ := exec.LookPath("zig"); zig != "" && os.Getenv("CC") == "" {
		if runtime.GOOS == "darwin" {
			if err := os.Setenv("CC", "clang"); err != nil {
				return xray.New(err)
			}
		} else {
			if err := os.Setenv("CC", "zig cc"); err != nil {
				return xray.New(err)
			}
		}
	}
	switch len(args) {
	case 1:
		if err := platform.Build(); err != nil {
			return xray.New(err)
		}
		if err := os.Chdir("./graphics"); err != nil {
			return xray.New(err)
		}
		return godot.Editor()
	default:
		switch args[1] {
		case "build":
			if err := os.MkdirAll(filepath.Join(project.ReleasesDirectory, GOOS, GOARCH), 0755); err != nil {
				return xray.New(err)
			}
			return platform.BuildMain(args[2:]...)
		case "run":
			return platform.Run(args[2:]...)
		case "test":
			return platform.Test(args[2:]...)
		default:
			return golang.Forwarded(args...)
		}
	}
}
