package project

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"graphics.gd/cmd/gd/internal/tooling"

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

	//go:embed graphics/gdscript_export_presets.cfg
	gdscript_export_presets_cfg string

	//go:embed graphics/gitignore
	gitignore string

	//go:embed graphics/icon.svg
	icon string
)

var (
	Name              string // Name of the current project (the name of the directory where go.mod is located).
	Directory         string // Directory of the current project (where go.mod is located).
	GraphicsDirectory string // Graphics directory.
	ReleasesDirectory string // Releases directory (Directory + "/releases"
	Version           string // extracted from project.godot config/version

	IncludesGo bool
)

func AndroidSafePackageName(name string) string {
	return strings.ReplaceAll(name, "-", "_")
}

func AppleSafePackageName(name string) string {
	return strings.ReplaceAll(name, "_", "")
}

func SetupVersion() {
	project, err := os.ReadFile(filepath.Join(GraphicsDirectory, "project.godot"))
	if err != nil {
		Version = "0.0.0"
		return
	}
	for line := range bytes.SplitSeq(project, []byte("\n")) {
		if bytes.HasPrefix(line, []byte("config/version=\"")) {
			Version = strings.TrimSuffix(strings.TrimPrefix(string(line), "config/version=\""), "\"")
			return
		}
	}
}

func Setup() error {
	defer SetupVersion()
	wd, err := os.Getwd()
	if err != nil {
		return xray.New(err)
	}
	var specificGoFile bool
	var isRun bool
	for _, arg := range os.Args[1:] {
		if strings.HasSuffix(arg, ".go") {
			specificGoFile = true
		}
		if arg == "run" {
			isRun = true
		}
	}
	wd, hasGoMod, err := findGoMod(wd)
	if err != nil {
		return xray.New(err)
	}
	var runningSpecificGoFile = specificGoFile && isRun
	if !runningSpecificGoFile && !hasGoMod {
		if _, err := os.Stat(filepath.Join(wd, "project.godot")); err == nil {
			Name = filepath.Base(wd)
			Directory = wd
			GraphicsDirectory = wd
			ReleasesDirectory = filepath.Join(wd, "releases")
			SetupFile(false, filepath.Join(ReleasesDirectory, ".gdignore"), "")
			if err := SetupFile(false, filepath.Join(GraphicsDirectory, "export_presets.cfg"), gdscript_export_presets_cfg, filepath.Base(wd), AndroidSafePackageName(filepath.Base(wd)), AppleSafePackageName(filepath.Base(wd))); err != nil {
				return xray.New(err)
			}
			return nil
		}
		return fmt.Errorf("gd requires your project to have either a go.mod file or a project.godot")
	}
	IncludesGo = true
	Name = filepath.Base(wd)
	Directory = wd
	GraphicsDirectory = filepath.Join(wd, "graphics")
	ReleasesDirectory = filepath.Join(wd, "releases")
	if runtime.GOOS == "android" {
		GraphicsDirectory = "/sdcard/gd/" + filepath.Base(wd) // Godot project needs to be in an accessible location
	}
	if err := os.MkdirAll(GraphicsDirectory, 0755); err != nil {
		return xray.New(err)
	}
	if _, err := os.Stat(filepath.Join(GraphicsDirectory, "project.godot")); os.IsNotExist(err) {
		// only create the main scene if the project.godot file doesn't exist yet
		if err := SetupFile(false, filepath.Join(GraphicsDirectory, "main.tscn"), main_tscn); err != nil {
			return xray.New(err)
		}
		if err := SetupFile(false, filepath.Join(GraphicsDirectory, "project.godot"), project_godot, filepath.Base(wd)); err != nil {
			return xray.New(err)
		}
	}
	if err := SetupFile(false, filepath.Join(GraphicsDirectory, "export_presets.cfg"), export_presets_cfg, filepath.Base(wd), AndroidSafePackageName(filepath.Base(wd)), AppleSafePackageName(filepath.Base(wd))); err != nil {
		return xray.New(err)
	}
	if err := SetupFile(false, filepath.Join(GraphicsDirectory, ".gitignore"), gitignore); err != nil {
		return xray.New(err)
	}
	gdextension_version, err := tooling.Godot.Output(tooling.Godot.VersionFlag)
	if err != nil {
		return xray.New(err)
	}
	if tooling.Godot.Name == "blazium" {
		gdextension_version = "4.1.0"
	}
	if err := SetupFile(true, filepath.Join(GraphicsDirectory, "library.gdextension"), library_gdextension, gdextension_version); err != nil {
		return xray.New(err)
	}
	if _, err := os.Stat(filepath.Join(GraphicsDirectory, ".godot")); os.IsNotExist(err) {
		current, err := os.Getwd()
		if err != nil {
			return xray.New(err)
		}
		if err := os.Chdir(GraphicsDirectory); err != nil {
			return xray.New(err)
		}
		if err := tooling.Godot.Exec("--import", "--headless"); err != nil {
			return xray.New(err)
		}
		if err := os.Chdir(current); err != nil {
			return xray.New(err)
		}
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

func SetupIcon() error {
	return SetupFile(false, filepath.Join(GraphicsDirectory, "icon.svg"), icon)
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

func findGoMod(wd string) (string, bool, error) {
	og := wd
	for last := ""; last != wd; last, wd = wd, filepath.Dir(wd) { // look for a go.mod file
		_, err := os.Stat(filepath.Join(wd, "go.mod"))
		if err == nil {
			return wd, true, nil
		} else if os.IsNotExist(err) {
			if _, err := os.Stat(filepath.Join(wd, "graphics")); err == nil {
				return wd, true, nil
			}
			continue
		} else {
			return wd, false, err
		}
	}
	wd = og
	for last := ""; last != wd; last, wd = wd, filepath.Dir(wd) {
		_, err := os.Stat(filepath.Join(wd, "project.godot"))
		if err == nil {
			return wd, false, nil
		}
	}
	return og, false, nil
}

// CopyDir recursively copies a directory tree from src to dst.
// It returns an error if the copy operation fails.
func CopyDir(src, dst string) error {
	// Get source directory info
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Create destination directory
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	// Read source directory entries
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Iterate through directory entries
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectories
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// Copy files
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// CopyFile copies a single file from src to dst
func CopyFile(src, dst string) error {
	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy file contents
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	// Copy file permissions
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}
