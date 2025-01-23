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
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"io"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	_ "embed"

	"graphics.gd/cmd/gd/internal/golang"
	"runtime.link/api/xray"
)

const version = "4.3"

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

func main() {
	if err := wrap(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func installGodot(gobin string) (string, error) {
	switch runtime.GOOS {
	case "android":
		return "echo", nil
	case "darwin":
		if _, err := exec.LookPath("brew"); err != nil {
			return "", fmt.Errorf("gd: cannot install godot without homebrew")
		}
		fmt.Println("gd: installing Godot stable for macOS (via brew)")
		if err := exec.Command("brew", "install", "godot").Run(); err != nil {
			return "", fmt.Errorf("gd: failed to 'brew install godot': %w", err)
		}
		return "godot", nil
	case "linux":
		fmt.Println("gd: downloading Godot v" + version + " stable for linux")
		resp, err := http.Get("https://github.com/godotengine/godot-builds/releases/download/" + version + "-stable/Godot_v" + version + "-stable_linux.x86_64.zip")
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
		inZip, err := archive.Open("Godot_v" + version + "-stable_linux.x86_64")
		if err != nil {
			return "", xray.New(err)
		}
		defer inZip.Close()
		//executable
		binPath := filepath.Join(gobin, "godot-"+version)
		file, err := os.OpenFile(binPath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return "", xray.New(err)
		}
		defer file.Close()
		if _, err = io.Copy(file, inZip); err != nil {
			return "", xray.New(err)
		}
		return binPath, nil
	default:
		return "", fmt.Errorf("gd: installing godot for GOOS %v is not supported", runtime.GOOS)
	}
}

type WebServer struct {
	http.Handler
}

func (s WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
	w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
	s.Handler.ServeHTTP(w, r)
}

func useGodot() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		gobin = filepath.Join(gopath, "bin")
	}
	godot, err := exec.LookPath("godot")
	if err == nil {
		if current, err := exec.Command(godot, "--version").CombinedOutput(); err == nil {
			if strings.HasPrefix(string(current), version+".") {
				return godot, nil
			}
		}
	}
	// Use existing godot if available and the correct version.
	if binary, err := exec.LookPath("godot-" + version); err == nil {
		return binary, nil
	}
	godotBin := filepath.Join(gobin + "godot-" + version)
	info, err := os.Stat(godotBin)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return "", xray.New(err)
		}
		godot, err := installGodot(gobin)
		if err != nil {
			return "", xray.New(err)
		}
		return godot, nil
	}
	if info.Mode()&0o111 == 0 {
		if err := os.Chmod(godotBin, 0o755); err != nil {
			return "", xray.New(err)
		}
	}
	return filepath.Join(gobin, "godot-"+version), nil
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
	godot, err := useGodot()
	if err != nil {
		return fmt.Errorf("gd requires Godot v%s to be installed as a binary at $GOPATH/bin/godot-%s: %w", version, version, err)
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
		if GOOS == "js" {
			if err := os.MkdirAll(graphics+"/.godot/public", 0o755); err != nil {
				return xray.New(err)
			}
			path := filepath.Join(graphics, ".godot", "public", "wasm_exec.js")
			if _, err := os.Stat(path); os.IsNotExist(err) {
				GOROOT := golang.CMD.Env.GOROOT()
				wasm_exec, err := os.Open(filepath.Join(GOROOT, "lib", "wasm", "wasm_exec.js"))
				if err != nil {
					return xray.New(err)
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
				fmt.Println("gd: downloading latest graphics.gd/godot.web.template_debug.wasm32.zip")
				data, err := io.ReadAll(resp.Body)
				if err != nil {
					return xray.New(err)
				}
				if err := os.WriteFile(template_path, data, 0o644); err != nil {
					return xray.New(err)
				}
			} else if resp.Body == nil || resp.StatusCode != 200 {
				return fmt.Errorf("gd: failed to download godot.web.template_debug.wasm32.zip: %v (is your gd command out of date?)", resp.Status)
			}
		}
		if err := setupFile(false, graphics+"/main.tscn", main_tscn); err != nil {
			return xray.New(err)
		}
		if err := setupFile(false, graphics+"/project.godot", project_godot, filepath.Base(wd)); err != nil {
			return xray.New(err)
		}
		if err := setupFile(false, graphics+"/export_presets.cfg", export_presets_cfg, filepath.Join(
			".godot", "godot.web.template_debug.wasm32.zip",
		)); err != nil {
			return xray.New(err)
		}
		if err := setupFile(true, graphics+"/library.gdextension", library_gdextension); err != nil {
			return xray.New(err)
		}
		if _, err := os.Stat(graphics + "/.godot"); os.IsNotExist(err) {
			godot := exec.Command(godot, "--import", "--headless")
			godot.Dir = graphics
			godot.Stderr = os.Stderr
			godot.Stdout = os.Stdout
			godot.Stdin = os.Stdin
			return xray.New(godot.Run())
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
	case "run", "build":
		copy(args, os.Args[1:])
		args[0] = "build"
		args = append(args, "-o", libraryPath)
		if GOOS != "js" {
			args = append(args, "-buildmode=c-shared")
		}
	case "test":
		args = []string{"test", "-c", "-o", libraryPath}
		if GOOS != "js" {
			args = append(args, "-buildmode=c-shared")
		} else {
			os.Args[1] = "run"
		}
	default:
		copy(args, os.Args[1:])
	}
	builds = append(builds, args)
	if runtime.GOOS == "darwin" {
		// GOARCH possible values = "amd64", "arm64"
		missingArch := "arm64"
		if GOARCH == "arm64" {
			missingArch = "amd64"
		}
		missingArgs := make([]string, len(os.Args)-1)
		missingLibraryName := fmt.Sprintf("%v_%v", GOOS, missingArch)
		missingArgs = append(args, "-buildmode=c-shared", "-o", graphics+"/"+missingLibraryName)
		builds = append(builds, missingArgs)
	}
	for _, commandArgs := range builds {
		golang := exec.Command("go", commandArgs...)
		if GOOS != "js" {
			golang.Env = append(os.Environ(), "CGO_ENABLED=1")
		}
		golang.Stderr = os.Stderr
		golang.Stdout = os.Stdout
		golang.Stdin = os.Stdin
		if err := golang.Run(); err != nil {
			return err
		}
	}
	if runtime.GOOS == "darwin" {
		// check if command is available in the system
		_, err := exec.LookPath("lipo")
		if err != nil {
			return fmt.Errorf("gd: lipo command not found in the system, please install it!")
		}
		lipoCommand := exec.Command("lipo", "-create", graphics+"/darwin_amd64.dylib", graphics+"/darwing_arm64.dylib", "-output", graphics+"/universal.dylib")
		lipoCommand.Stderr = os.Stderr
		lipoCommand.Stdout = os.Stdout
		lipoCommand.Stdin = os.Stdin
		if err := lipoCommand.Run(); err != nil {
			return err
		}
	}
	if err := setup(); err != nil {
		return xray.New(err)
	}
	switch os.Args[1] {
	case "run":
		godot := exec.Command(godot, runGodotArgs...)
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
		godot := exec.Command(godot, args...)
		godot.Dir = graphics
		godot.Stderr = os.Stderr
		godot.Stdout = os.Stdout
		godot.Stdin = os.Stdin
		return xray.New(godot.Run())
	}
	return nil
}
