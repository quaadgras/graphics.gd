package builder

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/tooling"

	"golang.org/x/mod/modfile"
)

// stub_main is the program that a Godot project with no Go code of its own
// is built with, wherever the engine is linked into a Go program (musl). It
// is graphics.gd's startup and nothing else, so the engine runs the project
// exactly as the official editor and export templates would.
const stub_main = `package main

import "graphics.gd/startup"

func main() { startup.Scene() }
`

// stubModule returns the directory of a Go module holding [stub_main], made
// (once for each graphics.gd) under gdpaths.Lib. It uses the graphics.gd that
// this gd was built from: the source tree itself when it is still there (a
// gd built in a checkout of graphics.gd, or installed from the module cache),
// otherwise the version that gd's build info records.
func stubModule() (string, error) {
	var require, replace, source string
	if dir, ok := graphicsSource(); ok {
		require, replace = "v0.0.0", dir
		// the stub follows the source tree's own requirements.
		gomod, err := os.ReadFile(filepath.Join(dir, "go.mod"))
		if err != nil {
			return "", err
		}
		source = string(gomod)
	} else if info, ok := debug.ReadBuildInfo(); ok && info.Main.Path == "graphics.gd" && info.Main.Version != "" && info.Main.Version != "(devel)" {
		if strings.HasSuffix(info.Main.Version, "+dirty") {
			return "", fmt.Errorf("gd was built from a graphics.gd checkout with uncommitted changes (%s) that is no longer where it was: rebuild gd, or install it with 'go install graphics.gd/cmd/gd@release'", info.Main.Version)
		}
		require = info.Main.Version
	} else {
		return "", errors.New("gd cannot tell which graphics.gd it was built from, to build a Godot project with: install gd with 'go install graphics.gd/cmd/gd@release'")
	}
	gomod := "module gdstub\n\ngo 1.24\n\nrequire graphics.gd " + require + "\n"
	if replace != "" {
		gomod += "\nreplace graphics.gd => " + modfile.AutoQuote(replace) + "\n"
	}
	stamped := gomod + stub_main + source
	sum := sha256.Sum256([]byte(stamped))
	dir := filepath.Join(gdpaths.Lib, "stub", hex.EncodeToString(sum[:])[:12])
	if was, err := os.ReadFile(filepath.Join(dir, "gd.stamp")); err == nil && string(was) == stamped {
		return dir, nil
	}
	// made beside where it goes and renamed into place, so that a gd running
	// at the same time never sees half of one.
	if err := os.MkdirAll(filepath.Dir(dir), 0o755); err != nil {
		return "", err
	}
	tmp, err := os.MkdirTemp(filepath.Dir(dir), "tmp-")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tmp)
	for name, content := range map[string]string{"go.mod": gomod, "main.go": stub_main} {
		if err := os.WriteFile(filepath.Join(tmp, name), []byte(content), 0o644); err != nil {
			return "", err
		}
	}
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if err := os.Chdir(tmp); err != nil {
		return "", err
	}
	defer os.Chdir(current)
	defer stubEnv()()
	// go.sum, and the go directive graphics.gd needs.
	if err := tooling.Go.Exec("mod", "tidy"); err != nil {
		return "", fmt.Errorf("preparing the Go program that hosts the engine: %w", err)
	}
	if err := os.WriteFile(filepath.Join(tmp, "gd.stamp"), []byte(stamped), 0o644); err != nil {
		return "", err
	}
	// one left by an earlier gd with no stamp, or a stale one.
	if err := os.RemoveAll(dir); err != nil {
		return "", err
	}
	if err := os.Rename(tmp, dir); err != nil {
		// another gd got there first.
		if was, statErr := os.ReadFile(filepath.Join(dir, "gd.stamp")); statErr != nil || string(was) != stamped {
			return "", err
		}
	}
	return dir, nil
}

// stubEnv sets up the environment of the go commands run in the stub module
// and returns a func that restores it: the stub belongs to no workspace of
// the user's and has no history of its own to stamp (every go command there
// would look for one: gd's own files live in the home directory, which may
// well be a git repository).
func stubEnv() (restore func()) {
	restoreFlags := setenv("GOFLAGS", strings.TrimSpace(os.Getenv("GOFLAGS")+" -buildvcs=false"))
	restoreWork := setenv("GOWORK", "off")
	return func() { restoreFlags(); restoreWork() }
}

// setenv sets the environment variable and returns a func that puts it back
// as it was, unset if it was.
func setenv(name, value string) (restore func()) {
	was, had := os.LookupEnv(name)
	os.Setenv(name, value)
	return func() {
		if had {
			os.Setenv(name, was)
		} else {
			os.Unsetenv(name)
		}
	}
}

// graphicsSource returns the root of the graphics.gd source tree this gd was
// compiled from, if it is still there.
func graphicsSource() (string, bool) {
	_, file, _, ok := runtime.Caller(0)
	if !ok || !filepath.IsAbs(file) {
		return "", false // -trimpath
	}
	// this file is cmd/gd/internal/builder/stub.go
	root := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(file)))))
	data, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil || modfile.ModulePath(data) != "graphics.gd" {
		return "", false
	}
	return root, true
}
