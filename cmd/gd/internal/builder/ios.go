package builder

import (
	_ "embed"
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

var (
	//go:embed bundled/ios/Info.plist
	info_plist string
)

type IOS struct{}

func (IOS) Build(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "darwin" {
		if err := zig.Assert(); err != nil {
			return xray.New(err)
		}
		project.SetupFiles(macos_sdk, "internal/macos", filepath.Join(project.ReleasesDirectory, "darwin", "sdk"))
		DARWIN_SDK, err := filepath.Abs(filepath.Join(project.ReleasesDirectory, "darwin", "sdk"))
		if err != nil {
			return xray.New(err)
		}
		switch GOARCH {
		case "arm64":
			if err := os.Setenv("CC", zig.Executable+" cc -target aarch64-macos -F "+DARWIN_SDK+"/Frameworks -L"+DARWIN_SDK+"/lib -I"+DARWIN_SDK+"/include"); err != nil {
				return xray.New(err)
			}
		default:
			return fmt.Errorf("gd build: cannot cross-compile linux %v on %v", GOARCH, runtime.GOOS)
		}
	}
	if err := golang.Build(args, "-buildmode=c-archive", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("darwin_%v.a", GOARCH))); err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll(filepath.Join(project.GraphicsDirectory, "go.xcframework", "ios_arm64"), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Rename(
		filepath.Join(project.GraphicsDirectory, "/ios_arm64.a"),
		filepath.Join(project.GraphicsDirectory, "go.xcframework", "ios_arm64", "libgo.a"),
	); err != nil {
		return xray.New(err)
	}
	if err := project.SetupFile(true, filepath.Join(project.GraphicsDirectory, "go.xcframework", "Info.plist"), info_plist); err != nil {
		return xray.New(err)
	}
	return nil
}

func (ios IOS) BuildMain(...string) error {
	if err := ios.Build(); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := godot.Run("--headless", "--export-release", "iOS"); err != nil {
		return xray.New(err)
	}
	return nil
}

func (IOS) Run(args ...string) error {
	return fmt.Errorf("gd run: ios not supported")
}

func (IOS) Test(args ...string) error {
	return fmt.Errorf("gd test: ios not supported")
}
