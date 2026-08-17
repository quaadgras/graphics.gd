// The 'gd' command is designed as a drop-in replacement of the 'go' command when working
// with Godot-based projects. It will automatically download and install the supported
// version of Godot and ensure all go commands behave as expected (go build, go run, etc).
//
// The 'gd' command assumes that the Go module lives at the root of the project, an empty
// godot project will be created under a 'graphics' directory, this is where the user can
// keep the graphical representation of their project and manage their assets. Running the
// command without any command line arguments will launch the Godot editor for managing
// the assets in this directory.
package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"slices"
	"strings"

	"graphics.gd/cmd/gd/internal/builder"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"

	"graphics.gd/internal/docgen"

	"runtime.link/api/xray"
)

func main() {
	/*if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" && info.Main.Version != "" {
	if dir, goModPath, ok := findProjectGoMod(); ok {
		if required := readGoModGraphicsVersion(goModPath); required != "" && required != info.Main.Version {
			if goPath, err := exec.LookPath("go"); err == nil {
				ensureGoToolGd(goPath, dir, goModPath)
				cmd := exec.Command(goPath, append([]string{"tool", "gd"}, os.Args[1:]...)...)
				cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
				if err := cmd.Run(); err != nil {
					if exitErr, ok := err.(*exec.ExitError); ok {
						os.Exit(exitErr.ExitCode())
					}
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}
	}*/
	if err := gd(os.Args[1:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "\nis this error unexpected? open an issue! https://github.com/quaadgras/graphics.gd/issues/new/choose")
		os.Exit(1)
	}
}

// setupDeterministicCGO disables floating-point contraction for every
// C compiler gd drives. A fused multiply-add rounds once where separate
// mul+add round twice, so two machines whose compilers fuse differently
// (a -march=native distro, a stock mingw on Windows, arm64 where fmadd
// is baseline) compute different floats from identical source — which
// silently breaks games that rely on cross-machine deterministic
// simulation. gcc, clang, zig cc, mingw and emscripten all accept the
// flag. Users keep the last word: flags they set in CGO_CFLAGS are
// preserved, including their own -ffp-contract choice. Restates Go's
// built-in "-g -O2" default because setting CGO_CFLAGS replaces it.
func setupDeterministicCGO() error {
	for _, env := range []string{"CGO_CFLAGS", "CGO_CXXFLAGS"} {
		flags := os.Getenv(env)
		switch {
		case flags == "":
			flags = "-g -O2 -ffp-contract=off"
		case strings.Contains(flags, "-ffp-contract"):
			continue
		default:
			flags += " -ffp-contract=off"
		}
		if err := os.Setenv(env, flags); err != nil {
			return err
		}
	}
	return nil
}

type Builder interface {
	Run(...string) error       // go run
	Build(...string) error     // go build -buildmode=c-shared
	BuildMain(...string) error // go build
	Test(...string) error      // go test
}

func builderFor(goos string) Builder {
	switch goos {
	case "linux", "ubuntu", "arch", "debian", "nix", "musl":
		if goos == "musl" {
			return &builder.Musl{}
		}
		return builder.Linux{}
	case "windows", "win":
		os.Setenv("GOOS", "windows")
		return builder.Windows{}
	case "darwin", "macos":
		os.Setenv("GOOS", "darwin")
		return builder.MacOS{}
	case "ios", "iphone":
		os.Setenv("GOOS", "ios")
		if os.Getenv("GOARCH") == "" {
			os.Setenv("GOARCH", "arm64")
		}
		return builder.IOS{}
	case "android":
		os.Setenv("GOOS", "android")
		if os.Getenv("GOARCH") == "" {
			os.Setenv("GOARCH", "arm64")
		}
		return builder.Android{}
	case "metaquest", "quest", "meta":
		// Pseudo-GOOS: builds an android/arm64 APK, then post-
		// processes it to inject the Khronos OpenXR Android loader
		// and the GodotVR Meta vendor plugin (both Apache-2.0,
		// fetched from Maven Central at build time). No gradle, no
		// Android SDK required beyond what graphics.gd already
		// manages.
		os.Setenv("GOOS", "android")
		os.Setenv("GOARCH", "arm64")
		return builder.MetaQuest{}
	case "browser", "js", "web", "wasm":
		os.Setenv("GOOS", "js")
		return builder.Browser{}
	default:
		fmt.Fprint(os.Stderr, "gd: unsupported GOOS '"+goos+"'\n")
		os.Exit(1)
		return nil
	}
}

func testArgs(args ...string) []string {
	converted := []string{}
	var benchmark bool
	for _, arg := range args {
		switch arg {
		case "-bench":
			benchmark = true
			fallthrough
		case "-benchmem", "-benchtime", "blockprofile",
			"-blockprofilerate", "-count", "-coverprofile", "-cpu",
			"-cpuprofile", "-failfast", "-fullpath", "-fuzz", "-fuzzcachedir",
			"-fuzzminimizetime", "-fuzztime", "-fuzzworker", "-gocoverdir",
			"-list", "-memprofile", "-memprofilerate", "-mutexprofile",
			"-mutexprofilefraction", "-outputdir", "-paniconexit0",
			"-parallel", "-run", "-short", "-shuffle", "-skip", "-testlogfile",
			"-timeout", "-trace", "-v":
			converted = append(converted, "-test."+strings.TrimPrefix(arg, "-"))
		default:
			converted = append(converted, arg)
		}
	}
	if !benchmark {
		converted = append([]string{"-gcflags=graphics.gd/classdb/...=-N -l"}, converted...)
	}
	return converted
}

func gd(args ...string) error {
	if err := setupDeterministicCGO(); err != nil {
		return xray.New(err)
	}
	// Pass through go commands that don't need project setup.
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		switch args[0] {
		case "version":
			version := "(devel)"
			if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
				version = info.Main.Version
			}
			fmt.Println("gd version", version)
			return tooling.Go.Exec(args...)
		case "doc":
			return doc(args[1:]...)
		case "build", "run", "test":
		default:
			return tooling.Go.Exec(args...)
		}
	}
	// On-device (Termux) there is no /tmp; the engine and other tooling honor
	// TMPDIR, which Termux sets in its own sessions but which may be missing
	// in stripped-down environments (ssh with a bare command, cron).
	if runtime.GOOS == "android" && os.Getenv("TMPDIR") == "" {
		tmp := "/data/data/com.termux/files/usr/tmp"
		if prefix := os.Getenv("PREFIX"); prefix != "" {
			tmp = filepath.Join(prefix, "tmp")
		}
		if err := os.MkdirAll(tmp, 0755); err == nil {
			os.Setenv("TMPDIR", tmp)
		}
	}
	var GOARCH = runtime.GOARCH
	var GOOS = runtime.GOOS
	if goos := os.Getenv("GOOS"); goos != "" {
		GOOS = goos
	}
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if GOARCH != "amd64" && GOARCH != "arm64" && GOARCH != "wasm" {
		return errors.New("gd requires an amd64, wasm, or arm64 GOARCH")
	}
	// Hot reloading is the default development experience: the editor
	// flow (plain `gd`, no arguments) builds the project with the
	// reloads tag, so code changes swap in live. This covers the
	// engine-as-library startup path (musl hosts run the
	// statically-linked editor binary directly) and the c-shared path
	// (the extension loaded into a godot editor process). Set
	// GD_NO_RELOAD=1 to opt out. Cross-compile targets keep their
	// native builds (mergeTags ignores GD_EXTRA_TAGS off the allowed
	// list, and the env is only set for the local development flow).
	// No reload mode on-device (Termux): the reloads host rebuilds the guest
	// with the Go toolchain, which the Godot Android Editor app that loads
	// the extension does not have.
	reloadMode := len(args) == 0 && os.Getenv("GD_NO_RELOAD") == "" && os.Getenv("GOOS") == "" && runtime.GOOS != "android"
	if reloadMode {
		// The reloads host is the project built with the reloads tag; on
		// musl hosts the build happens inside project.Setup via the musl
		// builder, which merges GD_EXTRA_TAGS into its tag set.
		if extra := os.Getenv("GD_EXTRA_TAGS"); extra != "" {
			os.Setenv("GD_EXTRA_TAGS", extra+",reloads")
		} else {
			os.Setenv("GD_EXTRA_TAGS", "reloads")
		}
	}
	var build_godot = func() error { return nil }
	var muslHost bool
	if runtime.GOOS == "linux" {
		version, err := tooling.ListDynamicDependencies.CombinedOutput("--version")
		muslHost = strings.HasPrefix(version, "musl")
		if !muslHost && err != nil {
			return xray.New(err)
		}
	} else if runtime.GOOS == "android" && len(args) > 0 {
		// On-device (Termux) there is no godot binary for android hosts, but
		// the statically-linked musl arm64 editor runs on the Android kernel,
		// so it stands in as the engine tooling: `gd test` runs the suite in
		// it headless (the musl-host flow) and `gd build` uses it for the
		// headless android export while the project itself is compiled for
		// android. The plain `gd` editor flow keeps using the Godot Android
		// Editor app: the static editor cannot load the device's bionic GPU
		// drivers, so it is headless-only.
		target := os.Getenv("GOOS")
		muslHost = (args[0] == "test" && (target == "" || target == "musl")) ||
			((args[0] == "build" || args[0] == "run") && (target == "" || target == "musl" || target == "android"))
	}
	if muslHost {
		if len(args) > 0 && args[0] == "test" {
			build_godot = func() error {
				// The host editor must be built for the host arch even when
				// the test target is a cross arch (e.g. android/arm64);
				// otherwise Musl.Test refuses to "run linux/arm64 on amd64".
				targetARCH := os.Getenv("GOARCH")
				os.Setenv("GOARCH", runtime.GOARCH)
				defer os.Setenv("GOARCH", targetARCH)
				musl_args := args
				current, err := os.Getwd()
				if err != nil {
					return xray.New(err)
				}
				os.Chdir(project.Directory)
				defer os.Chdir(current)
				var faster_compile = []string{"-gcflags=graphics.gd/classdb/...=-N -l"}
				if slices.Contains(musl_args, "-bench") {
					faster_compile = nil
				}
				if GOOS != "musl" && GOOS != "" {
					musl_args = []string{"test", "-test.skip", "."}
				}
				return builder.Musl{}.Test(append(faster_compile, testArgs(musl_args[1:]...)...)...)
			}
		} else {
			build_godot = func() error {
				GOARCH := os.Getenv("GOARCH")
				os.Setenv("GOARCH", runtime.GOARCH)
				defer os.Setenv("GOARCH", GOARCH)
				current, err := os.Getwd()
				if err != nil {
					return xray.New(err)
				}
				os.Chdir(project.Directory)
				defer os.Chdir(current)
				return builder.Musl{}.Build("-gcflags=graphics.gd/classdb/...=-N -l")
			}
		}
		if os.Getenv("GOOS") == "" {
			GOOS = "musl"
			if runtime.GOOS == "android" && len(args) > 0 && (args[0] == "build" || args[0] == "run") {
				// The default on-device build/run target is android; musl is
				// only the tooling that performs the headless export.
				GOOS = "android"
			}
		}
	}
	var platform = builderFor(GOOS)
	if goos := os.Getenv("GOOS"); goos != "" {
		GOOS = goos
	}
	if arch := os.Getenv("GOARCH"); arch != "" {
		GOARCH = arch
	}
	if GOOS != "js" {
		if err := os.Setenv("CGO_ENABLED", "1"); err != nil {
			return xray.New(err)
		}
	}
	if GOOS == "windows" && os.Getenv("CC") == "" {
		zig, err := tooling.Zig.Lookup()
		if err != nil {
			return xray.New(err)
		}
		if err := os.Setenv("CC", zig+" cc"); err != nil {
			return xray.New(err)
		}
	} else {
		if zig, _ := exec.LookPath("zig"); zig != "" && os.Getenv("CC") == "" {
			// On darwin and android the native clang is the working
			// compiler: zig can't find bionic's libc on-device (Termux)
			// and fails with LibCRuntimeNotFound.
			if runtime.GOOS == "darwin" || runtime.GOOS == "android" {
				if err := os.Setenv("CC", "clang"); err != nil {
					return xray.New(err)
				}
			} else {
				if err := os.Setenv("CC", "zig cc"); err != nil {
					return xray.New(err)
				}
			}
		}
	}
	if err := project.Setup(build_godot); err != nil {
		// In reload mode a compile error must not stop the session: fall
		// back to the previous host build if one exists — its file
		// watcher rebuilds the wasm guest as soon as the source
		// compiles again, so fixes hot-swap in as usual.
		stale := filepath.Join(project.GraphicsDirectory, "musl_"+GOARCH+".editor")
		if _, statErr := os.Stat(stale); !reloadMode || statErr != nil {
			return err
		}
		fmt.Fprintln(os.Stderr, "gd: build failed, launching the previous build — hot reload picks up fixes:", err)
		tooling.Godot.Path = stale
	}
	if project.IncludesGo {
		if err := docgen.Process(project.Directory); err != nil {
			return xray.New(err)
		}
	}
	var editorArgs []string
	if len(args) > 0 && strings.HasPrefix(args[0], "-") {
		editorArgs = args
		args = nil
	}
	switch len(args) {
	case 0:
		if err := os.Chdir(project.Directory); err != nil {
			return xray.New(err)
		}
		if err := platform.Build("-gcflags=graphics.gd/classdb/...=-N -l"); err != nil {
			if !reloadMode {
				return xray.New(err)
			}
			// A previously built extension (if any) still loads, and
			// once the source compiles again its watcher swaps the
			// fixed code in.
			fmt.Fprintln(os.Stderr, "gd: build failed, launching with the previous build — hot reload picks up fixes:", err)
		}
		if err := os.Chdir(project.GraphicsDirectory); err != nil {
			return xray.New(err)
		}
		if os.Getenv("RUNNING_INSIDE_GODOT") != "" {
			return nil
		}
		if runtime.GOOS == "android" {
			return openAndroidEditorApp(editorArgs)
		}
		return tooling.Godot.Exec(append([]string{"-e"}, editorArgs...)...)
	default:
		// On-device (Termux) these all need godot to run headless (--import,
		// --export-*), and the Godot Android Editor app strips command-line
		// arguments from external intents by design — so the static musl
		// editor stands in as the engine tooling: `gd test` runs the suite in
		// it (GOOS "musl") while `gd build` and `gd run` use it to export the
		// android project (run then installs the APK through the system
		// package installer and launches it). Anything else cannot run
		// on-device.
		if runtime.GOOS == "android" && GOOS != "musl" &&
			!(GOOS == "android" && (args[0] == "build" || args[0] == "run")) {
			return fmt.Errorf("gd %[1]s is not supported on android (Termux): the Godot editor app cannot run headless exports.\nRun 'gd' to open the project in the editor and use its play/export buttons, or run 'gd %[1]s android' from a desktop", args[0])
		}
		switch args[0] {
		case "build":
			if err := os.Chdir(project.Directory); err != nil {
				return xray.New(err)
			}
			// we need to make sure export templates are installed before building.
			if err := AssertExportTemplates(tooling.Godot.InstalledVersion(), GOOS); err != nil {
				return xray.New(err)
			}
			if err := os.MkdirAll(filepath.Join(project.ReleasesDirectory, GOOS, GOARCH), 0755); err != nil {
				return xray.New(err)
			}
			return platform.BuildMain(append([]string{"-ldflags=-s -w"}, args[1:]...)...)
		case "run":
			if err := os.Chdir(project.Directory); err != nil {
				return xray.New(err)
			}
			// android runs as an exported APK, which needs the export
			// template installed (like `gd build` does).
			if GOOS == "android" {
				if err := AssertExportTemplates(tooling.Godot.InstalledVersion(), GOOS); err != nil {
					return xray.New(err)
				}
			}
			return platform.Run(append([]string{"-gcflags=graphics.gd/classdb/...=-N -l"}, args[1:]...)...)
		case "test":
			if !project.IncludesGo {
				return errors.New("cannot run 'gd test' on a project that does not include Go code")
			}
			// android/ios run the suite as an exported app on a device,
			// emulator or simulator, so they need the export template installed
			// (like `gd build` does).
			if GOOS == "android" || GOOS == "ios" {
				if err := AssertExportTemplates(tooling.Godot.InstalledVersion(), GOOS); err != nil {
					return xray.New(err)
				}
			}
			return platform.Test(testArgs(args[1:]...)...)
		}
	}
	return nil
}

// openAndroidEditorApp opens the staged project in the Godot Android Editor
// app (the Termux flow, where there is no godot binary to exec). Arbitrary
// engine arguments cannot be passed: the editor strips the command-line extra
// from intents sent by other apps, but its exported ProjectManager activity
// accepts an ACTION_VIEW intent on a project.godot file and opens that project
// in the editor.
func openAndroidEditorApp(editorArgs []string) error {
	if len(editorArgs) > 0 {
		fmt.Fprintln(os.Stderr, "gd: ignoring", strings.Join(editorArgs, " "), "— the Godot editor app does not accept command-line arguments from other apps")
	}
	am, err := exec.LookPath("am")
	if err != nil {
		am = "/system/bin/am"
	}
	component := os.Getenv("GD_ANDROID_EDITOR")
	if component == "" {
		component = "org.godotengine.editor.v4/org.godotengine.editor.ProjectManager"
	}
	uri := "file://" + filepath.Join(project.GraphicsDirectory, "project.godot")
	out, err := exec.Command(am, "start", "-a", "android.intent.action.VIEW", "-n", component, "-d", uri).CombinedOutput()
	os.Stdout.Write(out)
	// `am start` reports failures like a missing activity on stdout with a
	// zero exit status, so scan the output as well.
	if err != nil || strings.Contains(string(out), "Error") {
		fmt.Fprintln(os.Stderr, "gd: could not open the Godot editor app.")
		fmt.Fprintln(os.Stderr, "	1. install the Godot Editor 4.x app: https://godotengine.org/download/android/")
		fmt.Fprintln(os.Stderr, "	2. grant it 'All files access' so it can open "+project.GraphicsDirectory)
		fmt.Fprintln(os.Stderr, "	(installed under a different package name? set GD_ANDROID_EDITOR=<package>/<activity>)")
		if err == nil {
			err = errors.New(strings.TrimSpace(string(out)))
		}
		return err
	}
	fmt.Println("gd: project synced with", project.GraphicsDirectory)
	return nil
}

// findProjectGoMod walks up from the current working directory looking for a go.mod file.
// Returns the directory containing go.mod, the full path to go.mod, and whether one was found.
func findProjectGoMod() (dir string, goModPath string, ok bool) {
	wd, err := os.Getwd()
	if err != nil {
		return "", "", false
	}
	for last := ""; last != wd; last, wd = wd, filepath.Dir(wd) {
		path := filepath.Join(wd, "go.mod")
		if _, err := os.Stat(path); err == nil {
			return wd, path, true
		}
	}
	return "", "", false
}

// readGoModGraphicsVersion reads a go.mod file and returns the required version of graphics.gd,
// or an empty string if graphics.gd is not a dependency.
func readGoModGraphicsVersion(goModPath string) string {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "graphics.gd ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "graphics.gd "))
		}
	}
	return ""
}

// ensureGoToolGd checks if graphics.gd/cmd/gd is registered as a tool in go.mod,
// and if not, runs "go get -tool graphics.gd/cmd/gd" to add it.
func ensureGoToolGd(goPath, dir, goModPath string) {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return
	}
	if strings.Contains(string(data), "graphics.gd/cmd/gd") {
		return
	}
	cmd := exec.Command(goPath, "get", "-tool", "graphics.gd/cmd/gd")
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
