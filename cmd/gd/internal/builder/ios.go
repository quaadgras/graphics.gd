package builder

import (
	"archive/zip"
	"embed"
	_ "embed"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/mdp/qrterminal/v3"

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
		filepath.Join(project.GraphicsDirectory, "darwin_arm64.a"),
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

	if err := os.MkdirAll(filepath.Join(project.ReleasesDirectory, "ios", "arm64"), 0o755); err != nil {
		return xray.New(err)
	}

	// if the Xcode project already exists, we don't want to overwrite any configuration, in this case,
	// just copy over the new go.xcframework
	if _, err := os.Stat(filepath.Join(project.ReleasesDirectory, "ios", "arm64", project.Name+".xcodeproj")); os.IsNotExist(err) {
		if err := os.Chdir(project.GraphicsDirectory); err != nil {
			return xray.New(err)
		}
		if err := tooling.Godot.Exec("--headless", "--export-release", "iOS"); err != nil {
			return xray.New(err)
		}
	} else {
		if err := project.CopyDir(filepath.Join(project.GraphicsDirectory, "go.xcframework"), filepath.Join(project.ReleasesDirectory, "ios", "arm64", project.Name, "dylibs", "go.xcframework")); err != nil {
			return xray.New(err)
		}
	}

	if runtime.GOOS == "darwin" {
		return nil
	}

	GDPATH := os.Getenv("GOPATH")
	if GDPATH == "" {
		GDPATH = filepath.Join(os.Getenv("HOME"), "gd")
	}

	apple_name := project.AppleSafePackageName(project.Name)

	if err := os.Chdir(filepath.Join(project.ReleasesDirectory, "ios", "arm64")); err != nil {
		return xray.New(err)
	}
	if err := os.RemoveAll(apple_name + ".app"); err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll(apple_name+".app", 0o755); err != nil {
		return xray.New(err)
	}
	if err := tooling.Zig.Exec("cc", "-target", "aarch64-ios",
		filepath.Join(".", "MoltenVK.xcframework", "ios-arm64", "libMoltenVK.a"),
		filepath.Join(".", "hello_triangle.xcframework", "ios-arm64", "libgodot.a"),
		filepath.Join(".", "hello_triangle", "dylibs", "go.xcframework", "ios-arm64", "libgo.a"),
		filepath.Join(".", project.Name, "dummy.cpp"),
		"-o", filepath.Join(apple_name+".app", apple_name), "-F", filepath.Join("..", "sdk", "Frameworks"), "-L"+filepath.Join("..", "sdk", "lib"),
		"-lc", "-lobjc.A", "-framework", "IOSurface", "-framework", "OpenGLES", "-framework", "CoreText", "-framework", "CoreGraphics",
		"-framework", "CoreFoundation", "-framework", "QuartzCore", "-lc++.1", "-framework", "UIKit", "-framework", "Foundation",
		"-framework", "Metal", "-framework", "GameController", "-framework", "CoreMotion",
		"-framework", "CoreHaptics", "-framework", "AVFAudio", "-framework", "AudioToolbox"); err != nil {
		return xray.New(err)
	}
	info, err := os.ReadFile(filepath.Join(".", project.Name, project.Name+"-Info.plist"))
	if err != nil {
		return xray.New(err)
	}
	replacer := strings.NewReplacer(
		"$(INFOPLIST_KEY_CFBundleDisplayName)", project.Name,
		"$(EXECUTABLE_NAME)", project.AppleSafePackageName(project.Name),
		"$(PRODUCT_BUNDLE_IDENTIFIER)", "gd.graphics",
		"$(BUNDLE_NAME)", project.AppleSafePackageName(project.Name),
		"$(PRODUCT_NAME)", project.Name,
		"$(MARKETING_VERSION)", "v0.1.0",
		"$(CURRENT_PROJECT_VERSION)", "1",
	)
	if err := os.WriteFile(filepath.Join(apple_name+".app", "Info.plist"), []byte(replacer.Replace(string(info))), 0o644); err != nil {
		return xray.New(err)
	}
	if err := project.CopyFile(project.Name+".pck", filepath.Join(apple_name+".app", apple_name+".pck")); err != nil {
		return xray.New(err)
	}
	if err := project.CopyFile(filepath.Join(project.Name, "Images.xcassets", "AppIcon.appiconset", "Icon-58.png"), filepath.Join(apple_name+".app", "Icon.png")); err != nil {
		return xray.New(err)
	}
	if err := project.CopyFile(filepath.Join(project.Name, "Images.xcassets", "AppIcon.appiconset", "Icon-114.png"), filepath.Join(apple_name+".app", "Icon@2x.png")); err != nil {
		return xray.New(err)
	}
	// FIXME I don't think the storyboard / splash image is working, probably needs to be compiled on MacOS.
	if err := project.CopyFile(filepath.Join(project.Name, "Launch Screen.storyboard"), filepath.Join(apple_name+".app", "Launch Screen.storyboard")); err != nil {
		return xray.New(err)
	}
	if err := project.CopyFile(filepath.Join(project.Name, "Images.xcassets", "SplashImage.imageset", "splash@2x.png"), filepath.Join(apple_name+".app", "SplashImage@2x.png")); err != nil {
		return xray.New(err)
	}
	if err := os.WriteFile(filepath.Join(apple_name+".app", "PkgInfo"), []byte("APPL????"), 0o644); err != nil {
		return xray.New(err)
	}

	ipa, err := os.Create(apple_name + ".ipa")
	if err != nil {
		return xray.New(err)
	}
	zipped := zip.NewWriter(ipa)
	defer ipa.Close()
	defer zipped.Close()

	filepath.Walk(apple_name+".app", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(filepath.Dir(apple_name+".app"), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		var header = &zip.FileHeader{
			Name:     filepath.ToSlash(filepath.Join("Payload", rel)),
			Method:   zip.Deflate,
			Modified: time.Now(),
		}
		header.SetMode(info.Mode())
		header.CreatorVersion = (3 << 8) // Unix
		f, err := zipped.CreateHeader(header)
		if err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		return err
	})

	if err := zipped.Close(); err != nil {
		return xray.New(err)
	}
	if err := ipa.Close(); err != nil {
		return xray.New(err)
	}

	return nil
}

func GetLocalIP() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { // Check if it's an IPv4 address
				return ipnet.IP, nil
			}
		}
	}
	return nil, fmt.Errorf("no non-loopback IPv4 address found")
}

func (ios IOS) Run(args ...string) error {
	if err := ios.BuildMain(args...); err != nil {
		return xray.New(err)
	}

	if runtime.GOOS == "darwin" {
		return fmt.Errorf("gd run: ios on macos not supported (yet)")
	}

	host, err := GetLocalIP()
	if err != nil {
		return xray.New(err)
	}

	values := url.Values{}
	values.Add("url", "http://"+host.String()+":4431/"+project.AppleSafePackageName(project.Name)+".ipa")
	sidestore_url := url.URL{
		Scheme:   "sidestore",
		Path:     "install",
		RawQuery: values.Encode(),
	}

	fmt.Println("Scan the following QRCode with your iOS device to install the app: (you will need SideStore installed: https://sidestore.io)")
	fmt.Println(sidestore_url.String())

	qrterminal.Generate(sidestore_url.String(), qrterminal.L, os.Stdout)

	return http.ListenAndServe(":4431", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Serving", project.AppleSafePackageName(project.Name)+".ipa", "to", r.RemoteAddr)
		http.ServeFile(w, r, filepath.Join(project.ReleasesDirectory, "ios", "arm64", project.AppleSafePackageName(project.Name)+".ipa"))
	}))
}

func (IOS) Test(args ...string) error {
	return fmt.Errorf("gd test: ios not supported")
}
