package builder

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	lipo "github.com/konoui/lipo/cmd"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling/godot"
	"graphics.gd/cmd/gd/internal/tooling/golang"
	"graphics.gd/cmd/gd/internal/tooling/zig"

	"runtime.link/api/xray"
)

var (
	// macos_sdk is a manually prepared "MacOS SDK" designed to support cross-compilation of Go code to arm64/amd64 targets using "zig cc".
	// the SDK was constructed by adding each undefined symbol / missing library observed from compilation errors on Linux GOOS=darwin as
	// .tbd files placed in the expected locations.
	//
	//go:embed bundled/macos
	macos_sdk embed.FS
)

type MacOS struct{}

func (MacOS) Build(args ...string) error {
	if runtime.GOOS != "darwin" {
		if err := zig.Assert(); err != nil {
			return xray.New(err)
		}
		project.SetupFiles(macos_sdk, "internal/macos", filepath.Join(project.ReleasesDirectory, "darwin", "sdk"))
		DARWIN_SDK, err := filepath.Abs(filepath.Join(project.ReleasesDirectory, "darwin", "sdk"))
		if err != nil {
			return xray.New(err)
		}
		if err := os.Setenv("CC", zig.Executable+"cc -target aarch64-macos -F "+DARWIN_SDK+"/Frameworks -L"+DARWIN_SDK+"/lib -I"+DARWIN_SDK+"/include"); err != nil {
			return xray.New(err)
		}
	}
	if err := os.Setenv("GOARCH", "arm64"); err != nil {
		return xray.New(err)
	}
	if err := golang.Build(args, "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, "darwin_arm64.dylib")); err != nil {
		return xray.New(err)
	}
	if runtime.GOOS != "darwin" {
		if err := zig.Assert(); err != nil {
			return xray.New(err)
		}
		DARWIN_SDK, err := filepath.Abs(filepath.Join(project.ReleasesDirectory, "darwin", "sdk"))
		if err != nil {
			return xray.New(err)
		}
		if err := os.Setenv("CC", zig.Executable+"cc -target x86_64-macos -F "+DARWIN_SDK+"/Frameworks -L"+DARWIN_SDK+"/lib -I"+DARWIN_SDK+"/include"); err != nil {
			return xray.New(err)
		}
	}
	if err := os.Setenv("GOARCH", "x86_64"); err != nil {
		return xray.New(err)
	}
	if err := golang.Build(args, "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, "darwin_amd64.dylib")); err != nil {
		return xray.New(err)
	}
	err := lipo.Execute(os.Stdout, os.Stderr,
		[]string{
			"-create",
			filepath.Join(project.GraphicsDirectory, "darwin_amd64.dylib"),
			filepath.Join(project.GraphicsDirectory, "darwin_arm64.dylib"),
			"-output",
			filepath.Join(project.GraphicsDirectory, "darwin_universal.dylib"),
		},
	)
	if err != 0 {
		return errors.New("lipo execution failed")
	}
	return nil
}

func (macos MacOS) BuildMain(...string) error {
	if err := macos.Build(); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := godot.Run("--headless", "--export-release", "macOS"); err != nil {
		return xray.New(err)
	}
	return nil
}

func (macos MacOS) Run(args ...string) error {
	if runtime.GOOS != "darwin" {
		return fmt.Errorf("gd run: cannot run darwin/universal executable on %v/%v", runtime.GOOS, runtime.GOARCH)
	}
	if err := golang.Build(args, "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("darwin_%v.dylib", runtime.GOARCH))); err != nil {
		return xray.New(err)
	}
	err := lipo.Execute(os.Stdout, os.Stderr,
		[]string{
			"-create",
			filepath.Join(project.GraphicsDirectory, "darwin_"+runtime.GOARCH+".dylib"),
			"-output",
			filepath.Join(project.GraphicsDirectory, "darwin_universal.dylib"),
		},
	)
	if err != 0 {
		return errors.New("lipo execution failed")
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return godot.Run()
}

func (MacOS) Test(args ...string) error {
	if runtime.GOOS != "darwin" {
		return fmt.Errorf("gd test: cannot run darwin/universal tests on %v/%v", runtime.GOOS, runtime.GOARCH)
	}
	if err := golang.Test(args, "-c", "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("darwin_%v.dylib", runtime.GOARCH))); err != nil {
		return xray.New(err)
	}
	err := lipo.Execute(os.Stdout, os.Stderr,
		[]string{
			"-create",
			filepath.Join(project.GraphicsDirectory, "darwin_"+runtime.GOARCH+".dylib"),
			"-output",
			filepath.Join(project.GraphicsDirectory, "darwin_universal.dylib"),
		},
	)
	if err != 0 {
		return errors.New("lipo execution failed")
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	return godot.Test(args...)
}
