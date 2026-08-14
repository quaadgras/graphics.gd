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
	"time"

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

type Builder interface {
	Run(...string) error       // go run
	Build(...string) error     // go build -buildmode=c-shared
	BuildMain(...string) error // go build
	Test(...string) error      // go test
}

func Setup(build_godot func() error) error {
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
	originalWd := wd
	wd, hasGoMod, err := findGoMod(wd)
	if err != nil {
		return xray.New(err)
	}
	var runningSpecificGoFile = specificGoFile && isRun
	// If there's no go.mod but there's a main.go in the current directory,
	// automatically run go mod init with the directory name.
	if !hasGoMod && !runningSpecificGoFile {
		if _, err := os.Stat(filepath.Join(originalWd, "main.go")); err == nil {
			if err := tooling.Go.Exec("mod", "init", filepath.Base(originalWd)); err != nil {
				return xray.New(err)
			}
			if err := tooling.Go.Exec("mod", "tidy"); err != nil {
				return xray.New(err)
			}
			if err := tooling.Go.Exec("get", "graphics.gd@release"); err != nil {
				return xray.New(err)
			}
			hasGoMod = true
			wd = originalWd
		}
	}
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
		// The Godot Android Editor app cannot read Termux's private home
		// directory, so the Godot project is staged on shared storage and
		// kept in sync with the repository's graphics directory: assets
		// added on the repository side reach the editor, and edits made in
		// the editor come back.
		local := GraphicsDirectory
		GraphicsDirectory = "/sdcard/gd/" + filepath.Base(wd) // Godot project needs to be in an accessible location
		if err := os.MkdirAll(GraphicsDirectory, 0755); err != nil {
			return fmt.Errorf("cannot create %s (in Termux, run 'termux-setup-storage' and grant storage access): %w", GraphicsDirectory, err)
		}
		if err := syncAndroidProject(local, GraphicsDirectory); err != nil {
			return xray.New(err)
		}
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
	// The icon is part of every project, not just the ones being packaged, so
	// it is written here rather than only by the icon builder.
	if err := SetupIcon(); err != nil {
		return xray.New(err)
	}
	if err := build_godot(); err != nil {
		return xray.New(err)
	}
	// On android (Termux) there is no godot binary to query or run as a
	// subprocess: the Godot Android Editor app opens the project instead and
	// imports resources itself when it does, so the compatibility version is
	// the one gd targets.
	gdextension_version := tooling.Godot.Version
	if runtime.GOOS != "android" {
		var err error
		gdextension_version, err = tooling.Godot.Output(tooling.Godot.VersionFlags...)
		if err != nil {
			return xray.New(err)
		}
		if tooling.Godot.Name == "blazium" {
			gdextension_version = "4.1.0"
		}
		if err := Import(); err != nil {
			return xray.New(err)
		}
	}
	if err := SetupFile(true, filepath.Join(GraphicsDirectory, "library.gdextension"), muslHostLibrary(library_gdextension), gdextension_version); err != nil {
		return xray.New(err)
	}
	// On android the .godot directory is not created by Import (which is
	// skipped there), so make sure it exists before writing into it.
	if err := os.MkdirAll(filepath.Join(GraphicsDirectory, ".godot"), 0755); err != nil {
		return xray.New(err)
	}
	if err := SetupFile(false, filepath.Join(GraphicsDirectory, ".godot", "extension_list.cfg"), extension_list_cfg); err != nil {
		return xray.New(err)
	}
	return nil
}

// muslHostLibrary scopes the linux entries of the library.gdextension template to
// template_debug/template_release on musl hosts. The editor there is a static binary
// with the project's Go code already linked in, so it must never dlopen the project's
// own linux .so: loading one left over from an older toolchain crashes the borrowed
// host loader mid-export, and even a fresh one is redundant. The editor (feature
// "editor") matches none of the scoped entries, while exported linux games (feature
// "template_*") still resolve theirs — so cross-platform and linux exports both keep
// working. Hosts with a glibc editor are left alone: dlopen'ing the .so is how the
// project loads there.
func muslHostLibrary(library string) string {
	if runtime.GOOS != "linux" {
		return library
	}
	// ldd reports "musl ..." on musl systems and "ldd (GNU libc) ..." on glibc ones.
	version, _ := tooling.ListDynamicDependencies.CombinedOutput("--version")
	if !strings.HasPrefix(strings.TrimSpace(version), "musl") {
		return library
	}
	var lines []string
	for _, line := range strings.Split(library, "\n") {
		if rest, ok := strings.CutPrefix(strings.TrimSpace(line), "linux."); ok {
			if arch, lib, ok := strings.Cut(rest, "="); ok {
				arch, lib = strings.TrimSpace(arch), strings.TrimSpace(lib)
				lines = append(lines,
					fmt.Sprintf("linux.template_debug.%s = %s", arch, lib),
					fmt.Sprintf("linux.template_release.%s = %s", arch, lib))
				continue
			}
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
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

// Import runs Godot's importer over the graphics directory when something in
// there has yet to be imported. Outside the editor a resource is only reachable
// through the remap the importer writes beside it — res://icon.svg resolves to
// its imported .ctex, and without one loading it fails with "No loader found" —
// so this has to happen before the engine runs. It is a no-op once everything
// is imported, and safe to call more than once.
//
// [Setup] calls this itself, but on musl hosts the whole test suite runs inside
// the build_godot callback (it is the static editor that runs the tests), which
// is too early: [builder.Musl] calls this again once it has an editor to run.
func Import() error {
	_, missing_godot_dir := os.Stat(filepath.Join(GraphicsDirectory, ".godot"))
	_, missing_icon_remap := os.Stat(filepath.Join(GraphicsDirectory, "icon.svg.import"))
	if !os.IsNotExist(missing_godot_dir) && !os.IsNotExist(missing_icon_remap) {
		return nil
	}
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
	return xray.New(os.Chdir(current))
}

// syncAndroidProject keeps the repository's graphics directory and the copy
// staged on shared storage (for the Godot Android Editor app to open) in
// sync: files missing on one side are copied there, and where both sides have
// a file the newer one wins. Nothing is ever deleted. Both directories live
// on the same device, so their timestamps are directly comparable; each copy
// carries the source's modification time so a synced file compares up to date
// on the next run. The .godot cache is skipped (each side rebuilds its own)
// and so are compiled libraries (the build writes them where they are used).
func syncAndroidProject(local, staged string) error {
	if err := os.MkdirAll(local, 0755); err != nil {
		return err
	}
	if err := syncNewerFiles(local, staged); err != nil {
		return err
	}
	return syncNewerFiles(staged, local)
}

var syncSkipSuffixes = []string{".so", ".a", ".editor", ".dll", ".dylib", ".lib", ".exe", ".xcframework"}

func syncSkip(name string, isDir bool) bool {
	if isDir {
		return name == ".godot" || name == ".git"
	}
	for _, suffix := range syncSkipSuffixes {
		if strings.HasSuffix(name, suffix) {
			return true
		}
	}
	return false
}

// syncNewerFiles copies every file under src that is missing under dst, or
// newer than its counterpart there (by more than the 2s resolution of shared
// storage timestamps), into dst. Shared storage supports neither file modes
// nor symlinks, so files are written plainly.
func syncNewerFiles(src, dst string) error {
	return fs.WalkDir(os.DirFS(src), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == "." {
			return nil
		}
		if syncSkip(d.Name(), d.IsDir()) {
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}
		target := filepath.Join(dst, filepath.FromSlash(path))
		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		existing, err := os.Stat(target)
		if err == nil && info.ModTime().Sub(existing.ModTime()) <= 2*time.Second {
			return nil // up to date, or the other side is newer (the reverse pass handles it)
		}
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		data, err := os.ReadFile(filepath.Join(src, filepath.FromSlash(path)))
		if err != nil {
			return err
		}
		if err := os.WriteFile(target, data, 0644); err != nil {
			return err
		}
		// Carry the source time over so the file compares up to date next
		// run; shared storage may refuse, in which case the next run copies
		// the same content again — harmless.
		_ = os.Chtimes(target, info.ModTime(), info.ModTime())
		fmt.Println("gd: synced", path, "to", dst)
		return nil
	})
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
