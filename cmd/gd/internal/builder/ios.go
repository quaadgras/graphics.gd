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
	project.SetupIcon()
	project.SetupFiles(macos_sdk, "bundled/macos", filepath.Join(project.ReleasesDirectory, "ios", "sdk"))
	DARWIN_SDK, err := filepath.Abs(filepath.Join(project.ReleasesDirectory, "ios", "sdk"))
	if err != nil {
		return xray.New(err)
	}
	GDPATH := os.Getenv("GOPATH")
	if GDPATH == "" {
		GDPATH = filepath.Join(os.Getenv("HOME"), "gd")
	}
	ZIG_INCLUDES := filepath.Join(GDPATH, "bin", "lib", "libc", "include", "any-macos-any")
	switch GOARCH {
	case "arm64":
		if err := os.Setenv("CC", zig+" cc -target aarch64-ios -F "+DARWIN_SDK+"/Frameworks -L"+DARWIN_SDK+"/lib -I"+DARWIN_SDK+"/include -I"+ZIG_INCLUDES+" -Wno-nullability-completeness"); err != nil {
			return xray.New(err)
		}
		if err := os.Setenv("GOARCH", "arm64"); err != nil {
			return xray.New(err)
		}
	default:
		return fmt.Errorf("gd build: cannot cross-compile linux %v on %v", GOARCH, runtime.GOOS)
	}
	if err := tooling.Go.Action("build", args, "-tags=ios", "-buildmode=c-archive", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("darwin_%v.a", GOARCH))); err != nil {
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

	// TODO when we want to support standalone ios builds on any platform, this will do it.
	// Main blocker here is code-signing and being able to launch the app on an actual device.
	//
	// zig cc -target aarch64-ios ./MoltenVK.xcframework/ios-arm64/libMoltenVK.a ./hello_triangle.xcframework/ios-arm64/libgodot.a ./go.xcframework/ios-arm64/libgo.a -o ios_app -F ./releases/ios/sdk/Frameworks -L./releases/ios/sdk/lib
	// -lc -lobjc.A -framework IOSurface -framework OpenGLES -framework CoreText -framework CoreGraphics -framework CoreFoundation -framework QuartzCore -lc++.1  -framework UIKit -framework Foundation -framework Metal
	// -framework GameController -framework CoreMotion -framework CoreHaptics -framework AVFAudio -framework AudioToolbox

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
