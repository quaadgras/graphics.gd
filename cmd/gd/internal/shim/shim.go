// Package shim lets gd stand in for the compilers of a proprietary SDK.
//
// Godot's build system insists on the toolchain of the platform's official
// SDK: platform/android/detect.py for example, assigns CC to a path inside
// the NDK no matter what is passed on the command line. Rather than
// requiring every engine fork to be patched, gd lays out a directory with
// the structure the build expects, where each tool is a link back to gd.
// [Run] notices when gd has been started through such a link and forwards
// to zig, translating whatever the two compiler drivers disagree on.
package shim

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"
)

// ConfigName of the file [Install] writes next to the tools, a link is only
// treated as a shim when it has this file beside it.
const ConfigName = "gd-shim.json"

// Config for the tools of a directory.
type Config struct {
	Zig   string `json:"zig"`   // path to the zig executable.
	LLVM  string `json:"llvm"`  // path to gd's llvm multi-tool, needed for libtool.
	Cache string `json:"cache"` // zig cache directory, build systems rarely pass $HOME through.

	// Targets are keyed by zig target triple (without any version).
	Targets map[string]Target `json:"targets"`

	// Default target, for builds that never name one (a linux build expects
	// the compiler on its PATH to know what it is for).
	Default string `json:"default,omitempty"`

	// Swift is an Objective-C source file that swift-frontend compiles
	// in place of whatever Swift it has been asked to, zig has no Swift.
	Swift string `json:"swift,omitempty"`
}

// Target specific configuration.
type Target struct {
	LibC string `json:"libc,omitempty"` // ZIG_LIBC file describing the target's libc.
	Lib  string `json:"lib,omitempty"`  // directory to search for libraries.

	// System and Frameworks directories are searched after the C++
	// standard library's, an -I naming one of them is dropped: clang
	// does as much for the directories of its sysroot, as libc++ only
	// works when its own headers are found ahead of the C library's.
	System     []string `json:"system,omitempty"`
	Frameworks []string `json:"frameworks,omitempty"`

	Flags []string `json:"flags,omitempty"` // added to every compilation.
}

// tools maps the name gd can be invoked as, to the zig subcommand it
// forwards to (or for those zig has none for, the name of the tool).
var tools = map[string]string{
	"clang":          "cc",
	"clang++":        "c++",
	"llvm-ar":        "ar",
	"llvm-ranlib":    "ranlib",
	"ar":             "ar",
	"ranlib":         "ranlib",
	"libtool":        "libtool",
	"swift-frontend": "swift-frontend",
}

// Install the named tools into dir, along with their config.
func Install(dir string, config Config, names ...string) error {
	gd, err := os.Executable()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, ConfigName), data, 0644); err != nil {
		return err
	}
	for _, name := range names {
		if _, ok := tools[name]; !ok {
			return fmt.Errorf("gd: no shim for %q", name)
		}
		if runtime.GOOS == "windows" {
			name += ".exe"
		}
		path := filepath.Join(dir, name)
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			return err
		}
		if err := os.Symlink(gd, path); err == nil {
			continue
		}
		// symlinks need privileges on windows.
		if err := os.Link(gd, path); err == nil {
			continue
		}
		exe, err := os.ReadFile(gd)
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, exe, 0755); err != nil {
			return err
		}
	}
	return nil
}

// Run does not return if gd was started as one of its shims.
func Run() {
	name := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
	tool, ok := tools[name]
	if !ok {
		return
	}
	// The config sits beside the link gd was started through, which is the
	// name alone when a build found it on the PATH (os.Executable would
	// follow the link to gd itself).
	link := os.Args[0]
	if !strings.ContainsRune(link, os.PathSeparator) {
		if found, err := exec.LookPath(link); err == nil {
			link = found
		}
	}
	data, err := os.ReadFile(filepath.Join(filepath.Dir(link), ConfigName))
	if err != nil {
		return
	}
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		fatal(err)
	}
	env := append(os.Environ(), "ZIG_GLOBAL_CACHE_DIR="+config.Cache, "ZIG_LOCAL_CACHE_DIR="+config.Cache)
	switch tool {
	case "cc", "c++":
		args, target := compiler([]string{config.Zig, tool}, os.Args[1:], config)
		if target.LibC != "" {
			env = append(env, "ZIG_LIBC="+target.LibC)
		}
		fatal(execute(config.Zig, args, env))
	case "swift-frontend":
		args, target := swift([]string{config.Zig, "cc"}, os.Args[1:], config)
		if target.LibC != "" {
			env = append(env, "ZIG_LIBC="+target.LibC)
		}
		fatal(execute(config.Zig, args, env))
	case "libtool":
		// Apple's libtool, of which only `-static -o` is ever asked for.
		fatal(execute(config.LLVM, append([]string{config.LLVM, "libtool-darwin"}, os.Args[1:]...), env))
	default:
		fatal(execute(config.Zig, append([]string{config.Zig, tool}, os.Args[1:]...), env))
	}
}

var (
	// androidTarget splits the API level off an NDK style triple.
	androidTarget = regexp.MustCompile(`^([a-z0-9_]+)-linux-(android(?:eabi)?)([0-9]*)$`)

	// appleTarget is how swift names a target, arm64-apple-ios14.0-simulator
	appleTarget = regexp.MustCompile(`^([a-z0-9_]+)-apple-(ios|macosx|xros)([0-9.]*)(-simulator)?$`)
)

// architectures that clang has a name for, which zig knows by another.
var architectures = map[string]string{
	"armv7a": "arm",
	"i686":   "x86",
	"arm64":  "aarch64",
}

func architecture(name string) string {
	if zig, ok := architectures[name]; ok {
		return zig
	}
	return name
}

// retarget a clang triple to a zig one, along with its configuration key.
func retarget(triple string) (zig, key string) {
	if match := androidTarget.FindStringSubmatch(triple); match != nil {
		key = architecture(match[1]) + "-linux-" + match[2]
		if match[3] != "" { // zig separates the API level.
			return key + "." + match[3], key
		}
		return key, key
	}
	if match := appleTarget.FindStringSubmatch(triple); match != nil {
		platform := strings.NewReplacer("macosx", "macos", "xros", "visionos").Replace(match[2])
		return apple(architecture(match[1]), platform, match[3], match[4] != "")
	}
	return triple, triple
}

// apple returns the zig target for an apple platform.
func apple(arch, platform, version string, simulator bool) (zig, key string) {
	key = arch + "-" + platform
	zig = key
	if version != "" {
		zig += "." + version
	}
	if simulator {
		zig += "-simulator"
		key += "-simulator"
	}
	return zig, key
}

// compiler translates clang arguments into zig cc ones.
func compiler(args, original []string, config Config) ([]string, Target) {
	var (
		key       string
		arch      string // -arch, which is how Xcode's clang is told the target.
		platform  string
		version   string
		simulator bool
		debug     bool
		rest      []string
	)
	for i := 0; i < len(original); i++ {
		arg := original[i]
		switch {
		case arg == "-target" && i+1 < len(original):
			i++
			var zig string
			zig, key = retarget(original[i])
			rest = append(rest, arg, zig)
		case strings.HasPrefix(arg, "--target="):
			var zig string
			zig, key = retarget(strings.TrimPrefix(arg, "--target="))
			rest = append(rest, "--target="+zig)
		case arg == "-arch" && i+1 < len(original):
			i++
			arch = architecture(original[i])
		case strings.HasPrefix(arg, "-miphoneos-version-min="):
			platform, version = "ios", strings.TrimPrefix(arg, "-miphoneos-version-min=")
		case strings.HasPrefix(arg, "-mios-simulator-version-min="):
			platform, version, simulator = "ios", strings.TrimPrefix(arg, "-mios-simulator-version-min="), true
		case strings.HasPrefix(arg, "-mmacosx-version-min="):
			platform, version = "macos", strings.TrimPrefix(arg, "-mmacosx-version-min=")
		case strings.HasSuffix(arg, ".os"):
			// zig decides what an input is from its extension and does
			// not know the one SCons gives to shared objects.
			if object, err := alias(arg); err == nil {
				arg = object
			}
			rest = append(rest, arg)
		default:
			debug = debug || strings.HasPrefix(arg, "-g")
			rest = append(rest, arg)
		}
	}
	if arch != "" && platform != "" {
		var zig string
		zig, key = apple(arch, platform, version, simulator)
		args = append(args, "-target", zig)
	}
	if key == "" && config.Default != "" {
		key = config.Default
		args = append(args, "-target", key)
	}
	target := config.Targets[key]
	for _, arg := range rest {
		if dir, ok := strings.CutPrefix(arg, "-I"); ok && contains(target.System, dir) {
			continue
		}
		args = append(args, arg)
	}
	if !debug {
		args = append(args, "-g0") // zig includes debug information unless told otherwise.
	}
	return append(args, target.flags()...), target
}

// flags every compilation for the target needs.
func (target Target) flags() []string {
	flags := slices.Clone(target.Flags)
	for _, dir := range target.System {
		flags = append(flags, "-isystem", dir)
	}
	for _, dir := range target.Frameworks {
		flags = append(flags, "-iframework", dir)
	}
	if target.Lib != "" {
		flags = append(flags, "-L"+target.Lib)
	}
	return flags
}

// swift translates a swift-frontend invocation into the compilation of
// the Objective-C that gd has to take the place of an engine's Swift.
func swift(args, original []string, config Config) ([]string, Target) {
	var key, zig, output string
	for i := 0; i+1 < len(original); i++ {
		switch original[i] {
		case "-target":
			zig, key = retarget(original[i+1])
		case "-o":
			output = original[i+1]
		}
	}
	target := config.Targets[key]
	args = append(args, "-target", zig, "-fobjc-arc", "-fblocks", "-fvisibility=hidden", "-O2", "-g0", "-c", config.Swift, "-o", output)
	return append(args, target.flags()...), target
}

func contains(dirs []string, dir string) bool {
	dir = filepath.Clean(dir)
	for _, other := range dirs {
		if filepath.Clean(other) == dir {
			return true
		}
	}
	return false
}

// alias returns a name ending in .o for the object at path.
func alias(path string) (string, error) {
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	object := path + ".o"
	os.Remove(object)
	if err := os.Symlink(filepath.Base(path), object); err == nil {
		return object, nil
	}
	if err := os.Link(path, object); err == nil {
		return object, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return object, os.WriteFile(object, data, 0644)
}

func fatal(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "gd shim:", err)
		os.Exit(1)
	}
	os.Exit(0)
}
