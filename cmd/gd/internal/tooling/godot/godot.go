package godot

import (
	"errors"
	"go/build"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"runtime.link/api/xray"
)

func Editor() error {
	if err := Assert(); err != nil {
		return xray.New(err)
	}
	godot := exec.Command(Executable, "-e")
	godot.Stderr = os.Stderr
	godot.Stdout = os.Stdout
	godot.Stdin = os.Stdin
	return godot.Run()
}

func Run(args ...string) error {
	if err := Assert(); err != nil {
		return xray.New(err)
	}
	godot := exec.Command(Executable, args...)
	godot.Stderr = os.Stderr
	godot.Stdout = os.Stdout
	godot.Stdin = os.Stdin
	return godot.Run()
}

func Test(args ...string) error {
	if err := Assert(); err != nil {
		return xray.New(err)
	}
	converted := []string{"--headless"}
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
			converted = append(converted, "-test."+strings.TrimPrefix(arg, "-"))
		default:
			converted = append(converted, arg)
		}
	}
	godot := exec.Command(Executable, converted...)
	godot.Stderr = os.Stderr
	godot.Stdout = os.Stdout
	godot.Stdin = os.Stdin
	return godot.Run()
}

var Executable string

func Assert() error {
	if Executable != "" {
		return nil
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		gobin = filepath.Join(gopath, "bin")
	}
	engine, err := exec.LookPath(Name)
	if err == nil {
		if current, err := exec.Command(engine, "--version").CombinedOutput(); err == nil {
			if strings.HasPrefix(string(current), Version+".") {
				Executable = engine
				return nil
			}
		}
	}
	// Use existing engine if available and the correct version.
	if binary, err := exec.LookPath(Name + "-" + Version); err == nil {
		Executable = binary
		return nil
	}
	engineBin := filepath.Join(gobin + Name + "-" + Version)
	info, err := os.Stat(engineBin)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return xray.New(err)
		}
		engine, err := install(gobin)
		if err != nil {
			return xray.New(err)
		}
		Executable = engine
		return nil
	}
	if info.Mode()&0o111 == 0 {
		if err := os.Chmod(engineBin, 0o755); err != nil {
			return xray.New(err)
		}
	}
	Executable = filepath.Join(gobin, engine+"-"+Version)
	return nil
}
