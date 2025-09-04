package project

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/tooling/godot"
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
)

var (
	Directory         string // Directory of the current project (where go.mod is located).
	GraphicsDirectory string // Graphics directory.
	ReleasesDirectory string // Releases directory (Directory + "/releases"
)

func AndroidSafePackageName(name string) string {
	return strings.ReplaceAll(name, "-", "_")
}

func Setup() error {
	if err := setupGoMod(); err != nil {
		return xray.New(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		return xray.New(err)
	}
	Directory = wd
	GraphicsDirectory = filepath.Join(wd, "graphics")
	ReleasesDirectory = filepath.Join(wd, "releases")
	if runtime.GOOS == "android" {
		GraphicsDirectory = "/sdcard/gd/" + filepath.Base(wd) // Godot project needs to be in an accessible location
	}
	if err := SetupFile(false, filepath.Join(GraphicsDirectory, "main.tscn"), main_tscn); err != nil {
		return xray.New(err)
	}
	if err := SetupFile(false, filepath.Join(GraphicsDirectory, "project.godot"), project_godot, filepath.Base(wd)); err != nil {
		return xray.New(err)
	}
	if err := SetupFile(false, filepath.Join(GraphicsDirectory, "export_presets.cfg"), export_presets_cfg, filepath.Base(wd), AndroidSafePackageName(filepath.Base(wd))); err != nil {
		return xray.New(err)
	}
	if err := SetupFile(false, filepath.Join(GraphicsDirectory, ".gitignore"), gitignore); err != nil {
		return xray.New(err)
	}
	var gdextension_version = godot.Version
	if godot.Name == "blazium" {
		gdextension_version = "4.1.0"
	}
	if err := SetupFile(true, filepath.Join(GraphicsDirectory, "library.gdextension"), library_gdextension, gdextension_version); err != nil {
		return xray.New(err)
	}
	if _, err := os.Stat(filepath.Join(GraphicsDirectory, ".godot")); os.IsNotExist(err) {
		engine := exec.Command(godot.Executable, "--import", "--headless")
		engine.Dir = GraphicsDirectory
		engine.Stderr = os.Stderr
		engine.Stdout = os.Stdout
		engine.Stdin = os.Stdin
		return xray.New(engine.Run())
	}
	if err := SetupFile(false, filepath.Join(GraphicsDirectory, ".godot", "extension_list.cfg"), extension_list_cfg); err != nil {
		return xray.New(err)
	}
	return nil
}

func SetupFile(force bool, name, embed string, args ...any) error {
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

// SetupFiles writes the contents of an embed.FS to the target directory on the OS filesystem.
func SetupFiles(embedded embed.FS, embedRoot, targetDir string) error {
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

func setupGoMod() error {
	wd, err := os.Getwd()
	if err != nil {
		return xray.New(err)
	}
	for wd := wd; true; wd = filepath.Dir(wd) { // look for a go.mod file
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
	return nil
}
