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
	for _, arg := range os.Args[1:] {
		if strings.HasSuffix(arg, ".go") {
			specificGoFile = true
		}
	}
	kind, wd := Detect(wd, os.Args[1:])
	switch kind {
	case Nothing:
		return fmt.Errorf("gd requires your project to have either a go.mod file or a project.godot")
	case Go:
		return fmt.Errorf("gd: %s is an ordinary Go module, not a graphics.gd project", wd)
	case Godot:
		return setupGodot(wd, build_godot)
	}
	// Go files that import graphics.gd/startup, with no go.mod around them:
	// make them a module that requires graphics.gd (unless gd was asked to
	// run a particular file, which go runs without one).
	if _, hasGoMod := findUp(wd, "go.mod"); !hasGoMod && !specificGoFile {
		if err := tooling.Go.Exec("mod", "init", filepath.Base(wd)); err != nil {
			return xray.New(err)
		}
		if err := tooling.Go.Exec("mod", "tidy"); err != nil {
			return xray.New(err)
		}
		if err := tooling.Go.Exec("get", "graphics.gd@release"); err != nil {
			return xray.New(err)
		}
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
		staged := "/sdcard/gd/" + filepath.Base(wd) // Godot project needs to be in an accessible location
		if termuxHeadlessCommand() {
			// The headless verbs run in the repository copy, but still pick
			// up edits made in the editor app when a staged copy exists;
			// they never create one (a pure headless workflow needs no
			// storage access).
			if _, err := os.Stat(staged); err == nil {
				if err := syncAndroidProject(local, staged); err != nil {
					return xray.New(err)
				}
			}
		} else {
			GraphicsDirectory = staged
			if err := os.MkdirAll(GraphicsDirectory, 0755); err != nil {
				return fmt.Errorf("cannot create %s (in Termux, run 'termux-setup-storage' and grant storage access): %w", GraphicsDirectory, err)
			}
			if err := syncAndroidProject(local, GraphicsDirectory); err != nil {
				return xray.New(err)
			}
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
	if err := SetupPresets(filepath.Join(GraphicsDirectory, "export_presets.cfg"), export_presets_cfg, filepath.Base(wd), AndroidSafePackageName(filepath.Base(wd)), AppleSafePackageName(filepath.Base(wd))); err != nil {
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
	// On android (Termux) there is normally no godot binary to query or run as
	// a subprocess: the Godot Android Editor app opens the project instead and
	// imports resources itself when it does, so the compatibility version is
	// the one gd targets. The exception is the headless verbs, where
	// build_godot has produced the static musl editor and pointed
	// tooling.Godot at it — then the normal query/import path applies.
	gdextension_version := tooling.Godot.Version
	if runtime.GOOS != "android" || tooling.Godot.Path != "" {
		var err error
		gdextension_version, err = tooling.Godot.Output(tooling.Godot.VersionFlags...)
		if err != nil {
			return xray.New(err)
		}
		if tooling.Godot.Name == "blazium" {
			gdextension_version = "4.1.0"
		}
		// only the numbers: the rest names the build ("stable.official",
		// "stable.custom_build.7f7882f94"), which differs between the
		// editors gd runs and would rewrite the file each time it does.
		gdextension_version = versionNumbers(gdextension_version)
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

// versionNumbers returns the leading numbers of a Godot version, ie. "4.7.2"
// of "4.7.2.stable.official.4b4c3ed45".
func versionNumbers(version string) string {
	parts := strings.Split(strings.TrimSpace(version), ".")
	n := 0
	for n < len(parts) && n < 3 && parts[n] != "" && strings.Trim(parts[n], "0123456789") == "" {
		n++
	}
	if n == 0 {
		return strings.TrimSpace(version)
	}
	return strings.Join(parts[:n], ".")
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
	switch runtime.GOOS {
	case "android":
		// On-device (Termux) the only engine binary matching the linux
		// feature set is the static musl editor doing headless work — same
		// deal as a musl host. The editor app and exported games match the
		// android entries, which are left untouched.
	case "linux":
		// ldd reports "musl ..." on musl systems and "ldd (GNU libc) ..." on glibc ones.
		version, _ := tooling.ListDynamicDependencies.CombinedOutput("--version")
		if !strings.HasPrefix(strings.TrimSpace(version), "musl") {
			return library
		}
	default:
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
		// an editor open on the project reacts to the file changing (for
		// library.gdextension, by loading the extension again), so a file
		// that already says this is left alone.
		if was, err := os.ReadFile(name); err == nil && string(was) == embed {
			return nil
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
func Import() error { return importProject(false) }

// importProject is [Import], which always imports when force is set.
func importProject(force bool) error {
	_, missing_godot_dir := os.Stat(filepath.Join(GraphicsDirectory, ".godot"))
	// the icon gd gives every graphics.gd project, which an ordinary Godot
	// project need not have.
	var missing_icon_remap error
	if _, err := os.Stat(filepath.Join(GraphicsDirectory, "icon.svg")); err == nil {
		_, missing_icon_remap = os.Stat(filepath.Join(GraphicsDirectory, "icon.svg.import"))
	}
	stale := staleImports(GraphicsDirectory)
	if !force && !os.IsNotExist(missing_godot_dir) && !os.IsNotExist(missing_icon_remap) && len(stale) == 0 {
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
	// Godot leaves an .import file alone when the asset's content is unchanged
	// (only its modification time moved, as after a git checkout), so bring the
	// .import files of the assets checked above up to date, or the next run
	// would import again for nothing.
	now := time.Now()
	for _, asset := range stale {
		if _, err := os.Stat(asset + ".import"); err == nil {
			os.Chtimes(asset+".import", now, now)
		}
	}
	return xray.New(os.Chdir(current))
}

// importable lists the file extensions that Godot imports as resources (and so
// writes an .import file beside), as opposed to loading them directly.
var importable = map[string]bool{
	// scenes and models
	".glb": true, ".gltf": true, ".obj": true, ".fbx": true, ".dae": true, ".blend": true,
	// images
	".png": true, ".jpg": true, ".jpeg": true, ".webp": true, ".svg": true, ".bmp": true,
	".tga": true, ".exr": true, ".hdr": true, ".dds": true, ".ktx": true, ".ktx2": true,
	// audio
	".wav": true, ".ogg": true, ".mp3": true,
	// fonts
	".ttf": true, ".otf": true, ".woff": true, ".woff2": true, ".pfb": true, ".pfm": true, ".fnt": true,
	// other
	".csv": true, ".bmfont": true,
}

// staleImports walks the graphics directory for importable assets that have no
// .import file beside them, or that were modified after it, and returns their
// paths. Like Godot, it skips hidden directories (such as the .godot cache)
// and directories marked with a .gdignore file.
func staleImports(dir string) []string {
	var stale []string
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			if path != dir && strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			if _, err := os.Stat(filepath.Join(path, ".gdignore")); err == nil {
				return filepath.SkipDir
			}
			return nil
		}
		if !importable[strings.ToLower(filepath.Ext(path))] {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return nil
		}
		imported, err := os.Stat(path + ".import")
		if err != nil || info.ModTime().After(imported.ModTime()) {
			stale = append(stale, path)
		}
		return nil
	})
	return stale
}

// termuxHeadlessCommand reports whether this on-device (Termux) invocation is
// a go verb (`gd build/run/test`, always os.Args[1]) rather than the plain
// `gd` editor flow. The verbs run headless through the statically-linked musl
// editor, a Termux process that reads the repository's graphics directory
// directly — staging on shared storage is only needed for the Godot Android
// Editor app, and /sdcard is mounted noexec, so the built editor binary could
// not run from the staged copy anyway.
func termuxHeadlessCommand() bool {
	return len(os.Args) > 1 && (os.Args[1] == "build" || os.Args[1] == "run" || os.Args[1] == "test")
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
