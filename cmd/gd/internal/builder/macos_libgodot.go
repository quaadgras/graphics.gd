package builder

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	lipo "github.com/konoui/lipo/cmd"

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

// macosSlices are the architectures a macOS app is built for, universal
// (like the stock export) unless GOARCH narrows it down to one.
type macosSlice struct {
	GOARCH, target, arch, minimum string
}

var macosSlices = []macosSlice{
	{"arm64", "aarch64-macos", "arm64", "11.0"},
	{"amd64", "x86_64-macos", "x86_64", "10.13"},
}

// buildMainLibgodot exports the project as an app whose executable is the
// Go program with the engine linked into it.
func (macos MacOS) buildMainLibgodot() error {
	custom, ok := customEngine()
	if !ok {
		return fmt.Errorf("gd: build/mode=libgodot needs an engine to link, set engine/repository or engine/strip_unused_classes in project.godot")
	}
	targets, variant := macosSlices, "universal"
	if GOARCH := os.Getenv("GOARCH"); GOARCH != "" {
		variant, targets = GOARCH, nil
		for _, slice := range macosSlices {
			if slice.GOARCH == GOARCH {
				targets = append(targets, slice)
			}
		}
		if len(targets) == 0 {
			return fmt.Errorf("gd build: cannot cross-compile macos %v", GOARCH)
		}
	}
	// The app bundle is put together by hand (the way the iOS one is): the
	// stock export would need the Xcode-built template it ships with.
	name := project.AppleSafePackageName(project.Name)
	app := filepath.Join(project.ReleasesDirectory, "darwin", variant, name+".app")
	if err := os.RemoveAll(app); err != nil {
		return xray.New(err)
	}
	for _, dir := range []string{"MacOS", "Resources"} {
		if err := os.MkdirAll(filepath.Join(app, "Contents", dir), 0755); err != nil {
			return xray.New(err)
		}
	}
	var executables []string
	minimum := targets[0].minimum
	for _, slice := range targets {
		executable, err := macos.linkLibgodot(custom, slice)
		if err != nil {
			return err
		}
		executables = append(executables, executable)
		if slice.minimum < minimum { // the versions compare as strings, 10.13 < 11.0.
			minimum = slice.minimum
		}
	}
	executable := filepath.Join(app, "Contents", "MacOS", name)
	if len(executables) == 1 {
		if err := copyFile(executables[0], executable); err != nil {
			return xray.New(err)
		}
	} else if lipo.Execute(os.Stdout, os.Stderr, append(append([]string{"-create"}, executables...), "-output", executable)) != 0 {
		return errors.New("lipo execution failed")
	}
	if err := os.Chmod(executable, 0755); err != nil {
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
	icon, err := macosIcon(filepath.Join(app, "Contents", "Resources", "icon.icns"))
	if err != nil {
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
%[5]s	<key>CFBundleIdentifier</key>
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
`, name, identifier, project.Version, minimum, map[bool]string{true: "\t<key>CFBundleIconFile</key>\n\t<string>icon.icns</string>\n"}[icon])
	if err := os.WriteFile(filepath.Join(app, "Contents", "Info.plist"), []byte(plist), 0644); err != nil {
		return xray.New(err)
	}
	if err := os.WriteFile(filepath.Join(app, "Contents", "PkgInfo"), []byte("APPL????"), 0644); err != nil {
		return xray.New(err)
	}
	fmt.Println("gd: exported", app)
	return nil
}

// linkLibgodot builds the engine and the Go program for one architecture
// and links them into an executable, returning its path.
func (macos MacOS) linkLibgodot(custom engine, slice macosSlice) (string, error) {
	libgodot, err := custom.macos(slice.GOARCH)
	if err != nil {
		return "", xray.New(err)
	}
	sdk, err := setupMacOSSDK(filepath.Join(gdpaths.Lib, "macos", "sdk"))
	if err != nil {
		return "", xray.New(err)
	}
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return "", xray.New(err)
	}
	if err := os.Setenv("CC", zig+" cc -target "+slice.target+" -F "+filepath.Join(sdk, "System", "Library", "Frameworks")+" -L"+filepath.Join(sdk, "usr", "lib")+" -isystem "+filepath.Join(sdk, "usr", "include")); err != nil {
		return "", xray.New(err)
	}
	if err := os.Setenv("GOARCH", slice.GOARCH); err != nil {
		return "", xray.New(err)
	}
	tags := mergeTags("macos", "archive")
	libgo := filepath.Join(project.GraphicsDirectory, fmt.Sprintf("darwin_%v.a", slice.GOARCH))
	args, err := archiver(nil)
	if err != nil {
		return "", xray.New(err)
	}
	if err := tooling.Go.Action("build", args, append(fastcbFlags("macos", ""), "-tags", tags, "-buildmode=c-archive", "-o", libgo)...); err != nil {
		return "", xray.New(err)
	}
	// @available checks compile to __isPlatformVersionAtLeast from the
	// compiler's runtime library, the same one the iOS app is given.
	versions := filepath.Join(project.GraphicsDirectory, ".godot", "platform_version.c")
	if err := os.WriteFile(versions, []byte(platformVersionStub), 0o644); err != nil {
		return "", xray.New(err)
	}
	versionsObj := strings.TrimSuffix(versions, ".c") + "." + slice.arch + ".o"
	if err := tooling.Zig.Exec(append([]string{"cc", "-target", slice.target, "-c", versions, "-o", versionsObj}, tooling.CGOCFlags()...)...); err != nil {
		return "", xray.New(err)
	}
	// the ad-hoc signature takes its identifier from the file's name.
	executable := filepath.Join(project.GraphicsDirectory, ".godot", "macos", slice.arch, project.AppleSafePackageName(project.Name))
	if err := os.MkdirAll(filepath.Dir(executable), 0755); err != nil {
		return "", xray.New(err)
	}
	link := []string{
		"-arch", slice.arch,
		"-platform_version", "macos", slice.minimum, slice.minimum,
		"-adhoc_codesign", // which apple silicon insists on, to run anything at all.
		"--error-limit=0", // every missing symbol at once, for the stubs to be completed from.
		"-syslibroot", "/dev/null",
		"-o", executable,
		libgodot, libgo, versionsObj,
	}
	cgoLDFLAGS, err := tooling.Go.Output("list", "-tags", tags, "-deps", "-f", "{{range .CgoLDFLAGS}}{{println .}}{{end}}", ".")
	if err != nil {
		return "", xray.New(err)
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
		return "", xray.New(err)
	}
	return executable, nil
}

// macosIcon writes the app's icon to path, from the same settings the
// stock export takes it from, reporting whether there was one to write.
// An .icns is used as is, any other image is rendered at each size the
// icon needs (by Godot, which reads every format the settings accept)
// and packed into one.
func macosIcon(path string) (bool, error) {
	source, _ := presetOption("macOS", "application/icon")
	if source == "" {
		source = projectSettings("application")["config/macos_native_icon"]
	}
	if source == "" {
		source = projectSettings("application")["config/icon"]
	}
	if source == "" {
		return false, nil
	}
	if strings.HasSuffix(strings.ToLower(source), ".icns") {
		return true, copyFile(filepath.Join(project.GraphicsDirectory, filepath.FromSlash(strings.TrimPrefix(source, "res://"))), path)
	}
	// An SVG is rendered at the largest size rather than scaled up from
	// its imported (typically small) texture, as the stock export does.
	script := filepath.Join(project.GraphicsDirectory, ".godot", "gd_icon.gd")
	if err := os.WriteFile(script, []byte(`extends SceneTree

func _init() -> void:
	var args := OS.get_cmdline_user_args()
	var image := Image.new()
	if args[0].get_extension().to_lower() == "svg":
		var svg := FileAccess.get_file_as_bytes(args[0])
		image.load_svg_from_buffer(svg)
		var scale := 1024.0 / maxf(image.get_width(), image.get_height())
		if scale > 1.0:
			image.load_svg_from_buffer(svg, scale)
	else:
		image.load(args[0])
	image.convert(Image.FORMAT_RGBA8)
	for size in [1024, 512, 256, 128, 64, 32]:
		var copy := image.duplicate()
		copy.resize(size, size, Image.INTERPOLATE_LANCZOS)
		copy.save_png(args[1].path_join("icon_%d.png" % size))
	quit()
`), 0644); err != nil {
		return false, err
	}
	renders := filepath.Join(project.GraphicsDirectory, ".godot", "icon")
	if err := os.MkdirAll(renders, 0755); err != nil {
		return false, err
	}
	if err := tooling.Godot.Exec("--headless", "-s", "res://.godot/gd_icon.gd", "--", source, renders); err != nil {
		return false, err
	}
	// The icon's entries, each a PNG at a size (the @2x sizes double as the
	// next size's plain entry), see Godot's EditorExportPlatformMacOS.
	entries := []struct {
		kind string
		size int
	}{
		{"ic10", 1024}, // 512×512@2x
		{"ic09", 512},
		{"ic14", 512}, // 256×256@2x
		{"ic08", 256},
		{"ic13", 256}, // 128×128@2x
		{"ic07", 128},
		{"ic12", 64}, // 32×32@2x
		{"ic11", 32}, // 16×16@2x
	}
	var icns []byte
	icns = append(icns, "icns\x00\x00\x00\x00"...)
	for _, entry := range entries {
		png, err := os.ReadFile(filepath.Join(renders, fmt.Sprintf("icon_%d.png", entry.size)))
		if err != nil {
			return false, fmt.Errorf("gd: the icon was not rendered at %d pixels: %w", entry.size, err)
		}
		icns = append(icns, entry.kind...)
		icns = binary.BigEndian.AppendUint32(icns, uint32(8+len(png)))
		icns = append(icns, png...)
	}
	binary.BigEndian.PutUint32(icns[4:], uint32(len(icns)))
	return true, os.WriteFile(path, icns, 0644)
}
