package builder

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/shim"
	"graphics.gd/cmd/gd/internal/tooling"
)

// engine is a custom build of Godot that a project asks to be exported
// with, configured in the [gd] section of its project.godot:
//
//	[gd]
//
//	engine/repository="https://github.com/example/godot"
//	engine/ref="my-branch"
//	engine/strip_unused_classes=true
//	engine/classes="Label,Sprite2D"
//	engine/options="optimize=size lto=full"
//
// The repository (with its branch, tag or commit) is what gets built, when
// there is none the engine is the stock Godot that the project is exported
// with. strip_unused_classes leaves out every class the project does not
// use (see [engine.profile]), classes names those that the project does
// use, but in a way that cannot be detected. options are passed to scons.
//
// gd builds the engine from source with zig, no platform SDKs are needed
// beyond the ones gd bundles.
type engine struct {
	Repository string // git URL, or a path relative to the project.
	Ref        string // branch, tag or commit, defaults to the remote's HEAD.

	StripUnusedClasses bool
	Classes            []string // kept, regardless of what is detected.
	Options            []string // for scons, last so that they take precedence.
}

// BuildMode of the project, configured in the [gd] section of its
// project.godot:
//
//	[gd]
//
//	build/mode="libgodot"
//
// In the default mode ("c-shared") the Go code is built as a shared
// library that the engine loads as an extension, and exports are the
// stock engine (or the custom one) with that library alongside. In
// libgodot mode the Go program is the entry point, with the engine linked
// into it as a library: exports are one executable, which is what the
// musl builds have always been, so that is what a linux export becomes.
func BuildMode() string {
	if mode := gdSettings()["build/mode"]; mode != "" {
		return mode
	}
	return "c-shared"
}

// gdSettings returns the [gd] section of the project's project.godot.
func gdSettings() map[string]string {
	return projectSettings("gd")
}

// projectSettings reads the named section of the project's project.godot.
func projectSettings(name string) map[string]string {
	settings := map[string]string{}
	data, err := os.ReadFile(filepath.Join(project.GraphicsDirectory, "project.godot"))
	if err != nil {
		return settings
	}
	section := ""
	for line := range strings.SplitSeq(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") {
			section = line
			continue
		}
		if section != "["+name+"]" {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key, value = strings.TrimSpace(key), strings.TrimSpace(value)
		if unquoted, err := strconv.Unquote(value); err == nil {
			value = unquoted
		}
		settings[key] = value
	}
	return settings
}

// customEngine returns the engine configured for the project, if any.
func customEngine() (engine, bool) {
	var custom engine
	configured := false
	for key, value := range gdSettings() {
		switch key {
		case "engine/repository":
			custom.Repository = value
		case "engine/ref":
			custom.Ref = value
		case "engine/strip_unused_classes":
			custom.StripUnusedClasses = value == "true"
		case "engine/classes":
			custom.Classes = strings.Split(value, ",")
		case "engine/options":
			custom.Options = strings.Fields(value)
		default:
			continue
		}
		configured = true
	}
	if !configured || (custom.Repository == "" && !custom.StripUnusedClasses && len(custom.Options) == 0) {
		return custom, false
	}
	if custom.Repository == "" {
		// the stock engine, at the version everything else about an export
		// (the templates, the editor) comes from.
		version := tooling.Godot.InstalledVersion()
		if version == "" {
			return custom, false
		}
		custom.Repository = "https://github.com/godotengine/godot"
		custom.Ref = version + "-stable"
	}
	return custom, true
}

// directory name for the engine, unique to its repository.
func (custom engine) directory() string {
	name := strings.TrimSuffix(filepath.Base(filepath.ToSlash(custom.Repository)), ".git")
	name = regexp.MustCompile(`[^A-Za-z0-9._-]+`).ReplaceAllString(name, "_")
	sum := sha256.Sum256([]byte(custom.Repository))
	return name + "-" + hex.EncodeToString(sum[:4])
}

// checkout the engine's source code, returning where it is along with
// the commit that has been checked out.
func (custom engine) checkout() (dir, commit string, err error) {
	dir = filepath.Join(filepath.Dir(gdpaths.Lib), "src", custom.directory())
	repository := custom.Repository
	if !strings.Contains(repository, ":") && !filepath.IsAbs(repository) {
		if repository, err = filepath.Abs(filepath.Join(project.Directory, repository)); err != nil {
			return "", "", err
		}
	}
	git := func(args ...string) (string, error) {
		cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		out, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("git %s: %w\n%s", strings.Join(args, " "), err, stderr.String())
		}
		return strings.TrimSpace(string(out)), nil
	}
	if _, err := os.Stat(filepath.Join(dir, ".git")); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", "", err
		}
		if _, err := git("init", "-q"); err != nil {
			return "", "", err
		}
		if _, err := git("remote", "add", "origin", repository); err != nil {
			return "", "", err
		}
	}
	if _, err := git("remote", "set-url", "origin", repository); err != nil {
		return "", "", err
	}
	head, _ := git("rev-parse", "--verify", "-q", "HEAD")
	if head != "" && head == custom.Ref {
		return dir, head, nil // pinned to the commit that is already checked out.
	}
	ref := custom.Ref
	if ref == "" {
		ref = "HEAD"
	}
	fmt.Println("gd: fetching engine", custom.Repository, custom.Ref)
	if _, err := git("fetch", "-q", "--depth", "1", "origin", ref); err != nil {
		if head == "" {
			return "", "", err
		}
		// Most likely offline, the last checkout is the best there is.
		fmt.Fprintln(os.Stderr, "gd: could not update the engine, continuing with", head[:12])
		fmt.Fprintln(os.Stderr, err)
		return dir, head, nil
	}
	if _, err := git("checkout", "-q", "--detach", "FETCH_HEAD"); err != nil {
		return "", "", err
	}
	commit, err = git("rev-parse", "HEAD")
	return dir, commit, err
}

// version of the engine's source code, ie. "4.7.2"
func engineVersion(src string) string {
	data, err := os.ReadFile(filepath.Join(src, "version.py"))
	if err != nil {
		return ""
	}
	var parts []string
	for _, key := range []string{"major", "minor", "patch"} {
		match := regexp.MustCompile(`(?m)^` + key + `\s*=\s*(\d+)`).FindSubmatch(data)
		if match == nil {
			return ""
		}
		parts = append(parts, string(match[1]))
	}
	return strings.TrimSuffix(strings.Join(parts, "."), ".0")
}

// scons runs the engine's build system inside of src, with the project's
// build profile and options on top of the given arguments.
func (custom engine) scons(src string, env []string, args ...string) error {
	if custom.StripUnusedClasses {
		profile, err := custom.profile(src)
		if err != nil {
			return err
		}
		args = append(args, "build_profile="+profile)
	}
	return scons(src, env, append(args, custom.Options...)...)
}

// scons runs the engine's build system inside of src.
func scons(src string, env []string, args ...string) error {
	name, prefix, python, err := tooling.SCons()
	if err != nil {
		return err
	}
	cmd := exec.Command(name, append(prefix, append(args, "-j"+strconv.Itoa(runtime.NumCPU()))...)...)
	cmd.Dir = src
	cmd.Env = append(append(os.Environ(), python...), env...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

// engineTarget returns the scons target to build.
func engineTarget(debug bool) string {
	if debug {
		return "template_debug"
	}
	return "template_release"
}

// engineRecipe is incremented whenever gd changes how it builds an engine
// (flags, SDKs), so that engines built any other way are not reused.
const engineRecipe = "r2"

// sdkSum is a hash of the SDKs gd bundles, which an engine is built against.
var sdkSum = sync.OnceValue(func() string {
	hash := sha256.New()
	for _, bundle := range []struct {
		fs   embed.FS
		root string
	}{{android_sdk, "bundled/android"}, {swappy_src, "bundled/swappy"}, {ios_sdk, "bundled/ios"}, {macos_sdk, "bundled/macos"}} {
		fs.WalkDir(bundle.fs, bundle.root, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return err
			}
			data, _ := bundle.fs.ReadFile(path)
			fmt.Fprintln(hash, path, len(data))
			hash.Write(data)
			return nil
		})
	}
	return hex.EncodeToString(hash.Sum(nil))
})

// artifact returns where a built file of the engine is kept, builds are
// only ever made once for each commit and way of building it (which the
// bundled SDKs are part of).
func (custom engine) artifact(src, commit, platform, arch string, debug bool, name string) (string, error) {
	hash := sha256.New()
	fmt.Fprintln(hash, engineRecipe, custom.Options, sdkSum())
	if custom.StripUnusedClasses {
		profile, err := custom.profile(src)
		if err != nil {
			return "", err
		}
		data, err := os.ReadFile(profile)
		if err != nil {
			return "", err
		}
		hash.Write(data)
	}
	variant := commit + "-" + hex.EncodeToString(hash.Sum(nil))[:8]
	return filepath.Join(gdpaths.Lib, "engine", custom.directory(), variant, platform, arch, engineTarget(debug), name), nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	tmp := dst + ".tmp"
	if err := os.WriteFile(tmp, data, 0755); err != nil {
		return err
	}
	return os.Rename(tmp, dst)
}

// warnEngineVersion when the engine is not the version of the Godot that
// everything else about an export (templates, the editor) comes from.
func warnEngineVersion(src string) {
	stock, custom := tooling.Godot.InstalledVersion(), engineVersion(src)
	if stock != "" && custom != "" && stock != custom {
		fmt.Fprintf(os.Stderr, "gd: warning: the custom engine is Godot %s but gd is exporting with Godot %s, these should match\n", custom, stock)
	}
}

// archiver returns the go build arguments, with `-buildmode=c-archive`
// told to pack its archive with zig's ar (through gd's shim) instead of
// the host's: macOS's ar leaves the objects of other systems out of the
// archive's index, and their linkers then find nothing in it. The flag is
// folded into any -ldflags among the arguments, as go takes the last.
func archiver(args []string) ([]string, error) {
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return nil, err
	}
	bin := filepath.Join(gdpaths.Lib, "ar")
	if err := shim.Install(bin, shim.Config{Zig: zig, Cache: filepath.Join(gdpaths.Lib, "cache")}, "ar"); err != nil {
		return nil, err
	}
	extar := "-extar=" + filepath.Join(bin, "ar")
	args = slices.Clone(args)
	for i, arg := range args {
		if strings.HasPrefix(arg, "-ldflags=") {
			args[i] = arg + " " + extar
			return args, nil
		}
	}
	return append(args, "-ldflags="+extar), nil
}
