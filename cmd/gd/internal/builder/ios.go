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

// swiftForceLoadStubs defines Swift FORCE_LOAD symbols for static-only libraries
// that exist in the Xcode toolchain but NOT as dylibs on the iOS device.
// All other Swift overlay FORCE_LOAD symbols are resolved via .tbd stubs in
// bundled/ios/lib/ which create proper LC_LOAD_DYLIB entries.
const swiftForceLoadStubs = `
void* _swift_FORCE_LOAD_$_swiftCompatibility56 = 0;
void* _swift_FORCE_LOAD_$_swiftCompatibilityConcurrency = 0;
void* _swift_FORCE_LOAD_$_swift_Builtin_float = 0;

// compiler-rt builtin: used by @available() checks in ObjC/Swift code.
// Normally statically linked by the compiler driver from libclang_rt.
#include <stdint.h>
extern int32_t __isPlatformVersionAtLeast(uint32_t platform, uint32_t major, uint32_t minor, uint32_t subminor) {
    // On iOS (platform 2), our deployment target is 14.0 and we only run on
    // devices with iOS 14+, so all @available checks for iOS <= deployment
    // target are satisfied. For runtime checks above deployment target, we
    // query the actual OS version via the dyld kernel info.
    (void)platform; (void)major; (void)minor; (void)subminor;
    return 1;
}
`

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
	if !project.IncludesGo {
		return nil
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

	// Check if this is a first build or subsequent build
	xcode_path := filepath.Join(project.ReleasesDirectory, "ios", "arm64", project.Name+".xcodeproj")
	existing_project := false
	if _, err := os.Stat(xcode_path); err == nil {
		existing_project = true
	}

	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}

	if existing_project {
		// Subsequent build: only export .pck, preserve Xcode project
		tooling.Godot.Exec("--headless", "--export-pack", "iOS", filepath.Join(project.ReleasesDirectory, "ios", "arm64", "shader.pck"))
	} else {
		// First build: full export
		tooling.Godot.Exec("--headless", "--export-release", "iOS")
	}

	// Copy the new go.xcframework
	if err := project.CopyDir(filepath.Join(project.GraphicsDirectory, "go.xcframework"), filepath.Join(project.ReleasesDirectory, "ios", "arm64", project.Name, "dylibs", "go.xcframework")); err != nil {
		return xray.New(err)
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
	// Write a small C file that defines Swift FORCE_LOAD symbols as no-ops.
	// These symbols are linker markers emitted by the Swift compiler to force-link
	// overlay libraries. Some of them (compatibility libs, Builtin_float) are static
	// libraries in the Xcode toolchain and do NOT exist as dylibs on the device,
	// so we must provide them directly rather than via .tbd stubs.
	swiftStubs := filepath.Join(apple_name+".app", "swift_stubs.c")
	if err := os.WriteFile(swiftStubs, []byte(swiftForceLoadStubs), 0o644); err != nil {
		return xray.New(err)
	}
	// Compile dummy.cpp and swift_stubs.c to .o files using zig cc.
	dummyObj := filepath.Join(apple_name+".app", "dummy.o")
	stubsObj := filepath.Join(apple_name+".app", "swift_stubs.o")
	if err := tooling.Zig.Exec("cc", "-c", "-target", "aarch64-ios", "-O2", filepath.Join(".", project.Name, "dummy.cpp"), "-o", dummyObj); err != nil {
		return xray.New(err)
	}
	if err := tooling.Zig.Exec("cc", "-c", "-target", "aarch64-ios", "-O2", swiftStubs, "-o", stubsObj); err != nil {
		return xray.New(err)
	}
	// Link with ld64.lld to produce a Mach-O binary with LC_DYLD_CHAINED_FIXUPS.
	var lld_args = []string{
		"-arch", "arm64",
		"-platform_version", "ios", "15.0.0", "15.0.0",
		"-fixup_chains", "-ignore_auto_link",
		"-syslibroot", "/dev/null",
		"-headerpad", "0x1000",
		filepath.Join(".", "MoltenVK.xcframework", "ios-arm64", "libMoltenVK.a"),
		filepath.Join(".", project.Name+".xcframework", "ios-arm64", "libgodot.a"),
		dummyObj,
		stubsObj,
	}
	if project.IncludesGo {
		lld_args = append(lld_args, filepath.Join(".", project.Name, "dylibs", "go.xcframework", "ios-arm64", "libgo.a"))
	}
	lld_args = append(lld_args,
		"-o", filepath.Join(apple_name+".app", apple_name),
		"-F", filepath.Join("..", "sdk", "Frameworks"), "-L", filepath.Join("..", "sdk", "lib"),
		"-lSystem", "-lobjc", "-lc++", "-lc++abi",
		"-framework", "IOSurface", "-framework", "OpenGLES", "-framework", "CoreText", "-framework", "CoreGraphics",
		"-framework", "CoreFoundation", "-framework", "QuartzCore", "-framework", "UIKit", "-framework", "Foundation",
		"-framework", "Metal", "-framework", "GameController", "-framework", "CoreMotion",
		"-framework", "CoreHaptics", "-framework", "AVFAudio", "-framework", "AudioToolbox",
		"-framework", "SwiftUI",
		"-lswiftCore", "-lswift_Concurrency", "-lswiftos",
		"-lswiftCoreFoundation", "-lswiftCoreImage", "-lswiftDarwin",
		"-lswiftDispatch", "-lswiftFoundation", "-lswiftMetal",
		"-lswiftOSLog", "-lswiftObjectiveC", "-lswiftQuartzCore",
		"-lswiftSpatial", "-lswiftUIKit", "-lswiftUniformTypeIdentifiers",
		"-lswiftXPC", "-lswiftsimd",
	)
	if err := tooling.LLVM.Exec(append([]string{"ld64.lld"}, lld_args...)...); err != nil {
		return xray.New(err)
	}
	// Clean up temp files before packaging the .app into an IPA.
	os.Remove(swiftStubs)
	os.Remove(dummyObj)
	os.Remove(stubsObj)
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
