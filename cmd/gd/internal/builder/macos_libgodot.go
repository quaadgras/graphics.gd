package builder

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"runtime.link/api/xray"
)

// macosFrameworks the engine links, as platform/macos/detect.py lists them.
var macosFrameworks = []string{
	"Cocoa", "AppKit", "Foundation", "CoreFoundation", "Carbon", "CoreServices", "ApplicationServices",
	"CoreGraphics", "CoreText", "QuartzCore", "Metal", "OpenGL", "IOSurface", "AudioUnit", "AudioToolbox",
	"CoreAudio", "CoreMIDI", "IOKit", "ForceFeedback", "GameController", "CoreHaptics", "CoreVideo",
	"AVFoundation", "AVFAudio", "CoreMedia", "Security", "UniformTypeIdentifiers",
}

// buildMainLibgodot exports the project as an app whose executable is the
// Go program with the engine linked into it.
func (macos MacOS) buildMainLibgodot() error {
	custom, ok := customEngine()
	if !ok {
		return fmt.Errorf("gd: build/mode=libgodot needs an engine to link, set engine/repository or engine/strip_unused_classes in project.godot")
	}
	GOARCH := runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	var target, arch, minimum string
	switch GOARCH {
	case "arm64":
		target, arch, minimum = "aarch64-macos", "arm64", "11.0"
	case "amd64":
		target, arch, minimum = "x86_64-macos", "x86_64", "10.13"
	default:
		return fmt.Errorf("gd build: cannot cross-compile macos %v", GOARCH)
	}
	libgodot, err := custom.macos(GOARCH)
	if err != nil {
		return xray.New(err)
	}
	sdk, err := setupMacOSSDK(filepath.Join(gdpaths.Lib, "macos", "sdk"))
	if err != nil {
		return xray.New(err)
	}
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if err := os.Setenv("CC", zig+" cc -target "+target+" -F "+filepath.Join(sdk, "System", "Library", "Frameworks")+" -L"+filepath.Join(sdk, "usr", "lib")+" -isystem "+filepath.Join(sdk, "usr", "include")); err != nil {
		return xray.New(err)
	}
	if err := os.Setenv("GOARCH", GOARCH); err != nil {
		return xray.New(err)
	}
	tags := mergeTags("macos", "archive")
	libgo := filepath.Join(project.GraphicsDirectory, fmt.Sprintf("darwin_%v.a", GOARCH))
	if err := tooling.Go.Action("build", nil, append(fastcbFlags("macos", ""), "-tags", tags, "-buildmode=c-archive", "-o", libgo)...); err != nil {
		return xray.New(err)
	}
	// The app bundle is put together by hand (the way the iOS one is): the
	// stock export would need the Xcode-built template it ships with.
	name := project.AppleSafePackageName(project.Name)
	app := filepath.Join(project.ReleasesDirectory, "darwin", GOARCH, name+".app")
	if err := os.RemoveAll(app); err != nil {
		return xray.New(err)
	}
	for _, dir := range []string{"MacOS", "Resources"} {
		if err := os.MkdirAll(filepath.Join(app, "Contents", dir), 0755); err != nil {
			return xray.New(err)
		}
	}
	// @available checks compile to __isPlatformVersionAtLeast from the
	// compiler's runtime library, the same one the iOS app is given.
	versions := filepath.Join(project.GraphicsDirectory, ".godot", "platform_version.c")
	if err := os.WriteFile(versions, []byte(platformVersionStub), 0o644); err != nil {
		return xray.New(err)
	}
	versionsObj := strings.TrimSuffix(versions, ".c") + ".o"
	if err := tooling.Zig.Exec(append([]string{"cc", "-target", target, "-c", versions, "-o", versionsObj}, tooling.CGOCFlags()...)...); err != nil {
		return xray.New(err)
	}
	link := []string{
		"-arch", arch,
		"-platform_version", "macos", minimum, minimum,
		"-adhoc_codesign", // which apple silicon insists on, to run anything at all.
		"--error-limit=0", // every missing symbol at once, for the stubs to be completed from.
		"-syslibroot", "/dev/null",
		"-o", filepath.Join(app, "Contents", "MacOS", name),
		libgodot, libgo, versionsObj,
	}
	cgoLDFLAGS, err := tooling.Go.Output("list", "-tags", tags, "-deps", "-f", "{{range .CgoLDFLAGS}}{{println .}}{{end}}", ".")
	if err != nil {
		return xray.New(err)
	}
	for _, flag := range strings.Split(cgoLDFLAGS, "\n") {
		if flag = strings.TrimSpace(flag); flag != "" && !strings.HasPrefix(flag, "-framework") {
			link = append(link, flag)
		}
	}
	link = append(link, "-F", filepath.Join(sdk, "System", "Library", "Frameworks"), "-L", filepath.Join(sdk, "usr", "lib"),
		"-lSystem", "-lobjc", "-lc++", "-lc++abi", "-lresolv")
	for _, framework := range macosFrameworks {
		link = append(link, "-framework", framework)
	}
	if err := ld64(link...); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	restoreExtensions, err := ignoreEnabledExtensions()
	if err != nil {
		return xray.New(err)
	}
	defer restoreExtensions()
	if err := tooling.Godot.Exec("--headless", "--export-pack", "macOS", filepath.Join(app, "Contents", "Resources", name+".pck")); err != nil {
		return xray.New(err)
	}
	identifier, _ := presetOption("macOS", "application/bundle_identifier")
	if identifier == "" {
		identifier = "org.godotengine." + strings.ToLower(name)
	}
	plist := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleDevelopmentRegion</key>
	<string>en</string>
	<key>CFBundleExecutable</key>
	<string>%[1]s</string>
	<key>CFBundleIdentifier</key>
	<string>%[2]s</string>
	<key>CFBundleInfoDictionaryVersion</key>
	<string>6.0</string>
	<key>CFBundleName</key>
	<string>%[1]s</string>
	<key>CFBundlePackageType</key>
	<string>APPL</string>
	<key>CFBundleShortVersionString</key>
	<string>%[3]s</string>
	<key>CFBundleVersion</key>
	<string>%[3]s</string>
	<key>LSMinimumSystemVersion</key>
	<string>%[4]s</string>
	<key>NSHighResolutionCapable</key>
	<true/>
	<key>NSPrincipalClass</key>
	<string>NSApplication</string>
</dict>
</plist>
`, name, identifier, project.Version, minimum)
	if err := os.WriteFile(filepath.Join(app, "Contents", "Info.plist"), []byte(plist), 0644); err != nil {
		return xray.New(err)
	}
	if err := os.WriteFile(filepath.Join(app, "Contents", "PkgInfo"), []byte("APPL????"), 0644); err != nil {
		return xray.New(err)
	}
	fmt.Println("gd: exported", app)
	return nil
}
