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
	"archive/zip"
	"bufio"
	"bytes"
	"embed"
	"errors"
	"fmt"
	"go/build"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	lipo "github.com/konoui/lipo/cmd"
	"graphics.gd/cmd/gd/internal/golang"
	"graphics.gd/cmd/gd/internal/upx"
	"graphics.gd/cmd/gd/internal/vpk"
	"graphics.gd/internal/docgen"
	"runtime.link/api/xray"
)

// These are our initial Godot project template files, we create
// these automatically when the user runs the 'gd' command. They
// are minimally setup for including the Go shared library such
// that it will be executed on startup.
var (
	//go:embed graphics/project.godot
	project_godot string

	//go:embed graphics/library.gdextension
	library_gdextension string

	//go:embed graphics/main.tscn
	main_tscn string

	//go:embed graphics/.godot/extension_list.cfg
	extension_list_cfg string

	//go:embed graphics/export_presets.cfg
	export_presets_cfg string

	//go:embed graphics/gitignore
	gitignore string

	// macos_sdk is a manually prepared "MacOS SDK" designed to support cross-compilation of Go code to arm64/amd64 targets using "zig cc".
	// the SDK was constructed by adding each undefined symbol / missing library observed from compilation errors on Linux GOOS=darwin as
	// .tbd files placed in the expected locations.
	//
	//go:embed internal/macos
	macos_sdk embed.FS

	//go:embed internal/ios/Info.plist
	info_plist string
)

func setupFile(force bool, name, embed string, args ...any) error {
	if _, err := os.Stat(name); force || os.IsNotExist(err) {
		if len(args) > 0 {
			embed = fmt.Sprintf(embed, args...)
		}
		if err := os.WriteFile(name, []byte(embed), 0o644); err != nil {
			return xray.New(err)
		}
	}
	return nil
}

// setupFiles writes the contents of an embed.FS to the target directory on the OS filesystem.
func setupFiles(embedded embed.FS, embedRoot, targetDir string) error {
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}
	return fs.WalkDir(embedded, embedRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		targetPath := filepath.Join(targetDir, filepath.FromSlash(strings.TrimPrefix(path, embedRoot)))
		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}
		data, err := embedded.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(targetPath, data, 0644)
	})
}

func main() {
	if err := wrap(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func useEngine() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		gobin = filepath.Join(gopath, "bin")
	}
	engine, err := exec.LookPath(engineCmd)
	if err == nil {
		if current, err := exec.Command(engine, "--version").CombinedOutput(); err == nil {
			if strings.HasPrefix(string(current), version+".") {
				return engine, nil
			}
		}
	}
	// Use existing engine if available and the correct version.
	if binary, err := exec.LookPath(engineCmd + "-" + version); err == nil {
		return binary, nil
	}
	engineBin := filepath.Join(gobin + engineCmd + "-" + version)
	info, err := os.Stat(engineBin)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return "", xray.New(err)
		}
		fmt.Println(engineBin)
		engine, err := installEngine(gobin)
		if err != nil {
			return "", xray.New(err)
		}
		return engine, nil
	}
	if info.Mode()&0o111 == 0 {
		if err := os.Chmod(engineBin, 0o755); err != nil {
			return "", xray.New(err)
		}
	}
	return filepath.Join(gobin, engine+"-"+version), nil
}

func download(gobin, unzip, url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", xray.New(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", xray.New(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", xray.New(err)
	}
	archive, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return "", xray.New(err)
	}
	inZip, err := archive.Open(unzip)
	if err != nil {
		return "", xray.New(err)
	}
	defer inZip.Close()
	//executable
	binPath := filepath.Join(gobin, engineCmd+"-"+version)
	file, err := os.OpenFile(binPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return "", xray.New(err)
	}
	defer file.Close()
	if _, err = io.Copy(file, inZip); err != nil {
		return "", xray.New(err)
	}
	return binPath, nil
}

type WebServer struct {
	http.Handler
}

func (s WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
	w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
	s.Handler.ServeHTTP(w, r)
}

func wrap() error {
	GOOS, GOARCH := runtime.GOOS, runtime.GOARCH
	if os.Getenv("GOOS") != "" {
		GOOS = os.Getenv("GOOS")
	}
	if os.Getenv("GOARCH") != "" {
		GOARCH = os.Getenv("GOARCH")
	}
	if GOARCH != "amd64" && GOARCH != "arm64" && GOARCH != "wasm" {
		return errors.New("gd requires an amd64, wasm, or arm64 system")
	}
	engine, err := useEngine()
	if err != nil {
		return fmt.Errorf("gd requires %s v%s to be installed as a binary at $GOPATH/bin/%s-%s: %w", engineCmd, version, engineCmd, version, err)
	}
	wd, err := os.Getwd()
	if err != nil {
		return xray.New(err)
	}
	// look for a go.mod file
	for wd := wd; true; wd = filepath.Dir(wd) {
		if wd == "/" {
			return fmt.Errorf("gd requires your project to have a go.mod file")
		}
		_, err := os.Stat(wd + "/go.mod")
		if err == nil {
			break
		} else if os.IsNotExist(err) {
			continue
		} else {
			return xray.New(err)
		}
	}
	graphics := "./graphics"
	if GOOS == "android" {
		graphics = "/sdcard/gd/" + filepath.Base(wd)
	}
	setup := func() error {
		wd, err := os.Getwd()
		if err != nil {
			return xray.New(err)
		}
		err = docgen.Process(wd)
		if err != nil {
			return xray.New(err)
		}
		if GOOS == "js" {
			if err := os.MkdirAll(graphics+"/.godot/public", 0o755); err != nil {
				return xray.New(err)
			}
			path := filepath.Join(graphics, ".godot", "public", "wasm_exec.js")
			if _, err := os.Stat(path); os.IsNotExist(err) {
				GOROOT := golang.CMD.Env.GOROOT()
				wasm_exec, err := os.Open(filepath.Join(GOROOT, "lib", "wasm", "wasm_exec.js"))
				if err != nil && !os.IsNotExist(err) {
					return xray.New(err)
				}
				err1 := err
				if err != nil {
					wasm_exec, err = os.Open(filepath.Join(GOROOT, "misc", "wasm", "wasm_exec.js"))
					if err != nil {
						return xray.New(errors.Join(err1, err))
					}
				}
				defer wasm_exec.Close()
				out, err := os.Create(path)
				if err != nil {
					return xray.New(err)
				}
				defer out.Close()
				if _, err := io.Copy(out, wasm_exec); err != nil {
					return xray.New(err)
				}
			}
			template_path := filepath.Join(graphics, ".godot", "godot.web.template_debug.wasm32.zip")
			stat, statErr := os.Stat(template_path)
			resp, err := http.Get("https://graphics.gd/godot.web.template_debug.wasm32.zip")
			if err != nil && !os.IsNotExist(err) {
				return xray.New(err)
			} else {
				defer resp.Body.Close()
			}
			last_modified, err := time.Parse(http.TimeFormat, resp.Header.Get("Last-Modified"))
			if err != nil {
				return xray.New(err)
			}
			if (os.IsNotExist(statErr) || (statErr == nil && last_modified.After(stat.ModTime()))) && resp.Body != nil {
				fmt.Println("gd: downloading latest graphics.gd/godot.web.template_release.wasm32.zip")
				data, err := io.ReadAll(resp.Body)
				if err != nil {
					return xray.New(err)
				}
				if err := os.WriteFile(template_path, data, 0o644); err != nil {
					return xray.New(err)
				}
			} else if resp.Body == nil || resp.StatusCode != 200 {
				return fmt.Errorf("gd: failed to download godot.web.template_release.wasm32.zip: %v (is your gd command out of date?)", resp.Status)
			}
		}
		if err := setupFile(false, graphics+"/main.tscn", main_tscn); err != nil {
			return xray.New(err)
		}
		if err := setupFile(false, graphics+"/project.godot", project_godot, filepath.Base(wd)); err != nil {
			return xray.New(err)
		}
		if err := setupFile(false, graphics+"/export_presets.cfg", export_presets_cfg, filepath.Base(wd)); err != nil {
			return xray.New(err)
		}
		if err := setupFile(false, graphics+"/.gitignore", gitignore); err != nil {
			return xray.New(err)
		}
		var gdextension_version = version
		if engineCmd == "blazium" {
			gdextension_version = "4.1.0"
		}
		if err := setupFile(true, graphics+"/library.gdextension", library_gdextension, gdextension_version); err != nil {
			return xray.New(err)
		}
		if _, err := os.Stat(graphics + "/.godot"); os.IsNotExist(err) {
			engine := exec.Command(engine, "--import", "--headless")
			engine.Dir = graphics
			engine.Stderr = os.Stderr
			engine.Stdout = os.Stdout
			engine.Stdin = os.Stdin
			return xray.New(engine.Run())
		}
		if err := setupFile(false, graphics+"/.godot/extension_list.cfg", extension_list_cfg); err != nil {
			return xray.New(err)
		}
		return nil
	}
	var runGodotArgs []string
	var libraryPath = graphics + "/" + fmt.Sprintf("%v_%v", GOOS, GOARCH)
	switch GOOS {
	case "windows":
		libraryPath += ".dll"
	case "darwin":
		libraryPath += ".dylib"
	case "js":
		libraryPath = filepath.Join(graphics, ".godot", "public", "library.wasm")
		runGodotArgs = []string{"--headless", "--export-debug", "Web"}
	case "android":
		libraryPath = "lib" + libraryPath + ".so"
	case "ios":
		libraryPath += ".a"
	default:
		libraryPath += ".so"
	}
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "run")
		runGodotArgs = []string{"-e"}
	}
	args := make([]string, len(os.Args)-1)
	builds := [][]string{}
	switch os.Args[1] {
	case "fix":
		return fix()
	case "run", "build":
		copy(args, os.Args[1:])
		args[0] = "build"
		args = append(args, "-o", libraryPath)
		if GOOS == "ios" {
			args = append(args, "-buildmode=c-archive")
		} else if GOOS != "js" {
			args = append(args, "-buildmode=c-shared")
		}
		args = append(args, "-ldflags=-s -w")
	case "test":
		args = []string{"test", "-c", "-o", libraryPath}
		if GOOS != "js" {
			args = append(args, "-buildmode=c-shared")
		} else {
			os.Args[1] = "run"
		}
	case "release":
		if _, err := exec.LookPath("vpk"); err != nil {
			return fmt.Errorf("gd: vpk command not found")
		}
		if len(os.Args) < 3 {
			return fmt.Errorf("gd: release command requires a version argument")
		}
		version := os.Args[2]
		mod, err := os.Open(wd + "/go.mod")
		if err != nil {
			return xray.New(err)
		}
		defer mod.Close()
		scanner := bufio.NewScanner(mod)
		scanner.Split(bufio.ScanLines)
		scanner.Scan()
		module := scanner.Text()
		module = strings.TrimPrefix(module, "module ")
		org := path.Base(path.Dir(module))
		if strings.Count(org, ".") > 1 {
			org = ""
		}
		targets := []string{
			"linux/arm64",
			"linux/amd64",
			"windows/amd64",
			"windows/arm64",
			// "darwin/universal", // FIXME/TODO vpk does not support cross compilation yet
		}
		/*var icon string
		if _, err := os.Stat(wd + "/graphics/icon.svg"); err == nil {
			icon = wd + "/graphics/icon.svg"
		}
		if _, err := os.Stat(wd + "/graphics/icon.png"); err == nil {
			icon = wd + "/graphics/icon.png"
			}*/
		for _, target := range targets {
			if _, err := os.Stat(wd + "/releases/" + target); os.IsNotExist(err) {
				continue
			}
			var extension string
			var icon string
			var Velopack vpk.Commands
			switch {
			case strings.HasPrefix(target, "linux/"):
				Velopack = vpk.CMD.Linux
				icon = wd + "/graphics/icon.svg"
			case strings.HasPrefix(target, "darwin/"):
				Velopack = vpk.CMD.OSX
			case strings.HasPrefix(target, "windows/"):
				Velopack = vpk.CMD.Windows
				extension = ".exe"
			}
			if err := Velopack.Pack(vpk.Package{
				ID:      org + "." + path.Base(wd),
				Version: version,
				Dir:     wd + "/releases/" + target,
				MainEXE: path.Base(wd) + extension,
				Output:  wd + "/releases",
				Icon:    icon,
			}); err != nil {
				return xray.New(err)
			}
		}
		return nil
	default:
		copy(args, os.Args[1:])
	}
	builds = append(builds, args)
	arches := []string{GOARCH}
	if GOOS == "darwin" && (GOARCH == "amd64" || GOARCH == "arm64") {
		// GOARCH possible values = "amd64", "arm64"
		missingArch := "arm64"
		if GOARCH == "arm64" {
			missingArch = "amd64"
		}
		missingArgs := make([]string, len(os.Args)-1)
		missingLibraryName := fmt.Sprintf("%v_%v.dylib", GOOS, missingArch)
		missingArgs = append(args, "-buildmode=c-shared", "-o", graphics+"/"+missingLibraryName)
		builds = append(builds, missingArgs)
		arches = append(arches, missingArch)
	}
	for i, commandArgs := range builds {
		var undefinedSymbols []string
		var parsingDone = make(chan struct{})
		stderr, capture, err := os.Pipe()
		if err != nil {
			return xray.New(err)
		}
		go func() {
			defer stderr.Close()
			defer close(parsingDone)
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Fprintln(os.Stderr, line)
				_, ident, ok := strings.Cut(line, "undefined: ")
				if ok {
					undefinedSymbols = append(undefinedSymbols, ident)
				}
			}
		}()
		golang := exec.Command("go", commandArgs...)
		golang.Env = os.Environ()
		if GOOS != "js" {
			golang.Env = append(os.Environ(), "CGO_ENABLED=1")
		}
		golang.Env = append(golang.Env, "GOARCH="+arches[i])
		if GOOS != runtime.GOOS {
			_, err := exec.LookPath("zig")
			if err != nil {
				return fmt.Errorf("gd requires zig to be installed and available in your PATH for cross-compilation: %w", err)
			}
			switch GOOS {
			case "windows":
				switch arches[i] {
				case "amd64":
					golang.Env = append(golang.Env, "CC=zig cc -target x86_64-windows-gnu")
				case "arm64":
					golang.Env = append(golang.Env, "CC=zig cc -target aarch64-windows-gnu")
				}
			case "darwin":
				setupFiles(macos_sdk, "internal/macos", graphics+"/../releases/darwin/sdk")
				MACOS_SDK, err := filepath.Abs(graphics + "/../releases/darwin/sdk")
				if err != nil {
					return xray.New(err)
				}
				switch arches[i] {
				case "amd64":
					golang.Env = append(golang.Env, "CC=zig cc -target x86_64-macos -F "+MACOS_SDK+"/Frameworks -L"+MACOS_SDK+"/lib -I"+MACOS_SDK+"/include")
				case "arm64":
					golang.Env = append(golang.Env, "CC=zig cc -target aarch64-macos -F "+MACOS_SDK+"/Frameworks -L"+MACOS_SDK+"/lib -I"+MACOS_SDK+"/include")
				}
			case "linux":
				switch arches[i] {
				case "amd64":
					golang.Env = append(golang.Env, "CC=zig cc -target x86_64-linux-gnu")
				case "arm64":
					golang.Env = append(golang.Env, "CC=zig cc -target aarch64-linux-gnu")
				}
			case "ios":
				setupFiles(macos_sdk, "internal/macos", graphics+"/../releases/darwin/sdk")
				MACOS_SDK, err := filepath.Abs(graphics + "/../releases/darwin/sdk")
				if err != nil {
					return xray.New(err)
				}
				golang.Env = append(golang.Env, "CC=zig cc -target aarch64-macos -F "+MACOS_SDK+"/Frameworks -L"+MACOS_SDK+"/lib -I"+MACOS_SDK+"/include")
			}
		}
		golang.Stderr = capture
		golang.Stdout = os.Stdout
		golang.Stdin = os.Stdin
		if err := golang.Run(); err != nil {
			capture.Close()
			<-parsingDone
			checkForFixes(undefinedSymbols)
			return err
		}
		if GOOS == "ios" { // create the xcframework
			if err := os.MkdirAll(graphics+"/go.xcframework/ios_arm64", 0755); err != nil {
				return xray.New(err)
			}
			if err := os.Rename(graphics+"/ios_arm64.a", graphics+"/go.xcframework/ios_arm64/libgo.a"); err != nil {
				return xray.New(err)
			}
			if err := setupFile(true, graphics+"/go.xcframework/Info.plist", info_plist); err != nil {
				return xray.New(err)
			}
		}
	}
	if GOOS == "darwin" && (GOARCH == "amd64" || GOARCH == "arm64") {
		err := lipo.Execute(os.Stdout, os.Stdin, []string{"-create", graphics + "/darwin_amd64.dylib", graphics + "/darwin_arm64.dylib", "-output", graphics + "/darwin_universal.dylib"})
		if err != 0 {
			return errors.New("lipo execution failed")
		}
	}
	if err := setup(); err != nil {
		return xray.New(err)
	}
	switch os.Args[1] {
	case "run":
		godot := exec.Command(engine, runGodotArgs...)
		godot.Dir = graphics
		godot.Stderr = os.Stderr
		godot.Stdout = os.Stdout
		godot.Stdin = os.Stdin
		if err := godot.Run(); err != nil {
			return xray.New(err)
		}
		if GOOS == "js" {
			PORT := os.Getenv("PORT")
			if PORT == "" {
				PORT = "8080"
			}
			fmt.Println("gd: serving wasm/js on http://localhost:" + PORT)
			http.Handle("/", WebServer{http.FileServer(http.Dir(filepath.Join(graphics, ".godot/public")))})
			return xray.New(http.ListenAndServe(":"+PORT, nil))
		}
		return nil
	case "build":
		if _, err := exec.LookPath("upx"); err == nil {
			os.Chmod(libraryPath, 0o755)
			upx.CMD.Compress(libraryPath)
		}
		switch GOOS {
		case "windows":
			switch GOARCH {
			case "amd64":
				runGodotArgs = []string{"--headless", "--export-release", "Windows x86_64"}
			case "arm64":
				runGodotArgs = []string{"--headless", "--export-release", "Windows arm64"}
			}
		case "linux":
			switch GOARCH {
			case "amd64":
				runGodotArgs = []string{"--headless", "--export-release", "Linux x86_64"}
			case "arm64":
				runGodotArgs = []string{"--headless", "--export-release", "Linux arm64"}
			}
		case "darwin":
			switch GOARCH {
			case "amd64", "arm64":
				runGodotArgs = []string{"--headless", "--export-release", "macOS"}
				GOARCH = "universal"
			}
		case "ios":
			switch GOARCH {
			case "arm64":
				runGodotArgs = []string{"--headless", "--export-release", "iOS"}
			}
		case "android":
			switch GOARCH {
			case "arm64":
				runGodotArgs = []string{"--headless", "--export-release", "Android"}
			}
		case "js":
			switch GOARCH {
			case "wasm":
				runGodotArgs = []string{"--headless", "--export-release", "Web"}
			}
		}
		if runGodotArgs == nil {
			return fmt.Errorf("gd: cannot build for GOOS %s and GOARCH %s", GOOS, GOARCH)
		}
		if err := os.MkdirAll(graphics+"/../releases/"+GOOS+"/"+GOARCH, 0755); err != nil {
			return xray.New(err)
		}
		godot := exec.Command(engine, runGodotArgs...)
		godot.Dir = graphics
		godot.Stderr = os.Stderr
		godot.Stdout = os.Stdout
		godot.Stdin = os.Stdin
		if err := godot.Run(); err != nil {
			return xray.New(err)
		}
	case "test":
		args := []string{"--headless"}
		for _, arg := range os.Args[2:] {
			switch arg {
			case "-bench", "-benchmem", "-benchtime", "blockprofile",
				"-blockprofilerate", "-count", "-coverprofile", "-cpu",
				"-cpuprofile", "-failfast", "-fullpath", "-fuzz", "-fuzzcachedir",
				"-fuzzminimizetime", "-fuzztime", "-fuzzworker", "-gocoverdir",
				"-list", "-memprofile", "-memprofilerate", "-mutexprofile",
				"-mutexprofilefraction", "-outputdir", "-paniconexit0",
				"-parallel", "-run", "-short", "-shuffle", "-skip", "-testlogfile",
				"-timeout", "-trace", "-v":
				args = append(args, "-test."+strings.TrimPrefix(arg, "-"))
			default:
				args = append(args, arg)
			}
		}
		engine := exec.Command(engine, args...)
		engine.Dir = graphics
		engine.Stderr = os.Stderr
		engine.Stdout = os.Stdout
		engine.Stdin = os.Stdin
		return xray.New(engine.Run())
	}
	return nil
}
