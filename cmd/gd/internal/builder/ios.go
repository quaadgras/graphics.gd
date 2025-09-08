package builder

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

var (
	//go:embed bundled/ios/Info.plist
	info_plist string

	// ios_sdk is a manually prepared "IOS SDK" designed to support cross-compilation of Go/Godot code to arm64-ios targets using "zig cc".
	// the SDK was constructed by adding each undefined symbol / missing library observed from compilation/linking errors into .tbd files
	// placed in the expected location.
	//
	//go:embed bundled/ios
	ios_sdk embed.FS
)

type IOS struct{}

func (IOS) Build(args ...string) error {
	var GOARCH = "arm64"
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if err := project.SetupFiles(ios_sdk, "bundled/ios", filepath.Join(project.ReleasesDirectory, "ios", "sdk")); err != nil {
		return xray.New(err)
	}
	project.SetupFiles(macos_sdk, "bundled/macos", filepath.Join(project.ReleasesDirectory, "ios", "sdk"))
	DARWIN_SDK, err := filepath.Abs(filepath.Join(project.ReleasesDirectory, "ios", "sdk"))
	if err != nil {
		return xray.New(err)
	}
	switch GOARCH {
	case "arm64":
		if err := os.Setenv("CC", zig+" cc -target aarch64-ios -F "+DARWIN_SDK+"/Frameworks -L"+DARWIN_SDK+"/lib -I"+DARWIN_SDK+"/include -Wno-nullability-completeness"); err != nil {
			return xray.New(err)
		}
		if err := os.Setenv("GOARCH", "arm64"); err != nil {
			return xray.New(err)
		}
	default:
		return fmt.Errorf("gd build: cannot cross-compile linux %v on %v", GOARCH, runtime.GOOS)
	}
	fmt.Println(args)
	if err := tooling.Go.Action("build", args, "-buildmode=c-archive", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("darwin_%v.a", GOARCH))); err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll(filepath.Join(project.GraphicsDirectory, "go.xcframework", "ios-arm64"), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Rename(
		filepath.Join(project.GraphicsDirectory, "/darwin_arm64.a"),
		filepath.Join(project.GraphicsDirectory, "go.xcframework", "ios-arm64", "libgo.a"),
	); err != nil {
		return xray.New(err)
	}
	if err := project.SetupFile(true, filepath.Join(project.GraphicsDirectory, "go.xcframework", "Info.plist"), info_plist); err != nil {
		return xray.New(err)
	}
	return nil
}

func (ios IOS) BuildMain(args ...string) error {
	if err := ios.Build(args...); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-release", "iOS"); err != nil {
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
