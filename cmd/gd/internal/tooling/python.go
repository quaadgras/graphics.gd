package tooling

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
)

// SCons is the build system of the engine, which needs a python to run
// on: gd installs a self-contained CPython (the builds of
// python-build-standalone) with SCons beside it under $GDPATH/lib/python,
// so that building an engine depends on nothing the host has installed.
// With GOTOOLCHAIN=local the host's scons is used, or its python when
// SCons has been installed as a package there.
//
// The command to run is returned along with any environment it needs.
func SCons() (name string, args []string, env []string, err error) {
	if Local() {
		if _, err := exec.LookPath("scons"); err == nil {
			return "scons", nil, nil, nil
		}
		python := "python3"
		if _, err := exec.LookPath(python); err != nil {
			python = "python"
		}
		if exec.Command(python, "-c", "import SCons").Run() != nil {
			return "", nil, nil, errors.New("gd: building a custom engine needs scons, which can be installed with 'pip install scons'\n(see https://scons.org/pages/download.html)")
		}
		return python, []string{"-m", "SCons"}, nil, nil
	}
	dir := filepath.Join(gdpaths.Lib, "python")
	python := filepath.Join(dir, "bin", "python3")
	if runtime.GOOS == "windows" {
		python = filepath.Join(dir, "python.exe")
	}
	// the interpreter and SCons are installed together, whenever either
	// version changes both are replaced.
	stamp, versions := filepath.Join(dir, "gd-python.txt"), pythonVersion+"+"+pythonRelease+" scons "+sconsVersion
	if existing, err := os.ReadFile(stamp); err != nil || string(existing) != versions {
		if err := installPython(dir); err != nil {
			return "", nil, nil, err
		}
		if err := os.WriteFile(stamp, []byte(versions), 0644); err != nil {
			return "", nil, nil, err
		}
	}
	scons := filepath.Join(dir, "scons")
	if existing := os.Getenv("PYTHONPATH"); existing != "" {
		scons += string(os.PathListSeparator) + existing
	}
	return python, []string{"-m", "SCons"}, []string{"PYTHONPATH=" + scons}, nil
}

const (
	pythonVersion = "3.13.15"
	pythonRelease = "20260901" // of python-build-standalone
	sconsVersion  = "4.11.1"
	sconsWheel    = "https://files.pythonhosted.org/packages/8e/43/d6285848e893c19682c06e92679dc1a07d37ff7ea148747b1df681ec496c/scons-4.11.1-py3-none-any.whl"
)

// installPython downloads the interpreter for the host into dir, with
// SCons extracted from its wheel (a zip of pure python) into dir/scons.
func installPython(dir string) error {
	const required, hint = "building a custom engine", "https://www.python.org/downloads/ and 'pip install scons'"
	arch, ok := map[string]string{"amd64": "x86_64", "arm64": "aarch64"}[runtime.GOARCH]
	if !ok {
		return fmt.Errorf("gd: no python available for %s/%s (required for %s), please install it, ie. %s", runtime.GOOS, runtime.GOARCH, required, hint)
	}
	var system string
	switch runtime.GOOS {
	case "linux":
		system = "unknown-linux-gnu"
		if version, _ := exec.Command("ldd", "--version").CombinedOutput(); strings.HasPrefix(strings.TrimSpace(string(version)), "musl") {
			system = "unknown-linux-musl"
		}
	case "darwin":
		system = "apple-darwin"
	case "windows":
		system = "pc-windows-msvc"
	default:
		return fmt.Errorf("gd: no python available for %s/%s (required for %s), please install it, ie. %s", runtime.GOOS, runtime.GOARCH, required, hint)
	}
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	url := fmt.Sprintf("https://github.com/astral-sh/python-build-standalone/releases/download/%s/cpython-%s+%s-%s-%s-install_only_stripped.tar.gz",
		pythonRelease, pythonVersion, pythonRelease, arch, system)
	archive := filepath.Join(filepath.Dir(dir), "python."+pythonVersion+".download")
	if err := download("python", pythonVersion, required, hint, url, archive); err != nil {
		return err
	}
	if err := ExtractArchive(archive, dir, "tar.gz", "", true); err != nil {
		return err
	}
	if err := os.Remove(archive); err != nil {
		return err
	}
	archive = filepath.Join(filepath.Dir(dir), "scons."+sconsVersion+".download")
	if err := download("scons", sconsVersion, required, hint, sconsWheel, archive); err != nil {
		return err
	}
	if err := ExtractArchive(archive, filepath.Join(dir, "scons"), "zip", "", false); err != nil {
		return err
	}
	return os.Remove(archive)
}
