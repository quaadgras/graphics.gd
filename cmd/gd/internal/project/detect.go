package project

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/mod/modfile"
)

// Kind is the kind of project that gd was run in.
type Kind int

const (
	// Graphics is a Go module that uses graphics.gd, with its Godot
	// project in the graphics directory beside go.mod.
	Graphics Kind = iota
	// Go is an ordinary Go module, which gd builds, runs and tests the way
	// go does, only with its compiler that shares generic instantiations.
	Go
	// Godot is an ordinary Godot project, with no Go code of its own, which
	// gd exports to every platform.
	Godot
	// Nothing is a directory with neither a Go module nor a Godot project.
	Nothing
)

// Detect works out which kind of project the directory wd belongs to, and
// where its root is. Whichever of go.mod and project.godot is nearer wins, a
// project.godot in the same directory as a go.mod included, unless the Go
// module uses graphics.gd: then it is a graphics.gd project, with its Godot
// project in the graphics directory (the nearest directory with one, which
// need not be the module's own). Without a go.mod, the Go files of wd
// make it a graphics.gd project when they import graphics.gd/startup, and an
// ordinary Go program otherwise. args are gd's arguments.
func Detect(wd string, args []string) (Kind, string) {
	goDir, hasGoMod := findUp(wd, "go.mod")
	godotDir, hasGodot := findUp(wd, "project.godot")
	// a module can hold more than one graphics.gd project (graphics.eg has
	// one per sample), each the directory beside its graphics directory.
	if dir, ok := findUp(wd, filepath.Join("graphics", "project.godot")); ok &&
		(!hasGodot || !within(godotDir, dir) || godotDir == filepath.Join(dir, "graphics")) &&
		(!hasGoMod || (within(dir, goDir) && usesGraphics(filepath.Join(goDir, "go.mod")))) {
		return Graphics, dir
	}
	if hasGoMod && usesGraphics(filepath.Join(goDir, "go.mod")) {
		// cwd inside the graphics directory (or a project.godot beside
		// go.mod) still means the graphics.gd project around it, but a
		// Godot project of its own further down a graphics.gd module is
		// not one: gd would build the module instead of exporting it.
		if !hasGodot || godotDir == goDir || godotDir == filepath.Join(goDir, "graphics") || !within(godotDir, goDir) {
			return Graphics, goDir
		}
		return Godot, godotDir
	}
	if hasGodot && (!hasGoMod || within(godotDir, goDir)) {
		return Godot, godotDir
	}
	if hasGoMod {
		return Go, goDir
	}
	// no go.mod: the Go files named on the command line (gd run main.go),
	// or else those of the current directory, say what this is.
	var files []string
	for _, arg := range args {
		if strings.HasSuffix(arg, ".go") {
			if !filepath.IsAbs(arg) {
				arg = filepath.Join(wd, arg)
			}
			files = append(files, arg)
		}
	}
	if len(files) == 0 {
		files, _ = filepath.Glob(filepath.Join(wd, "*.go"))
	}
	if len(files) == 0 {
		return Nothing, wd
	}
	for _, file := range files {
		if importsStartup(file) {
			return Graphics, wd
		}
	}
	return Go, wd
}

// findUp returns the nearest directory, wd or above, that holds name.
func findUp(wd, name string) (string, bool) {
	for last := ""; last != wd; last, wd = wd, filepath.Dir(wd) {
		if _, err := os.Stat(filepath.Join(wd, name)); err == nil {
			return wd, true
		}
	}
	return "", false
}

// within reports whether dir is parent or a directory below it.
func within(dir, parent string) bool {
	rel, err := filepath.Rel(parent, dir)
	return err == nil && rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}

// usesGraphics reports whether the go.mod at path requires graphics.gd, or
// is graphics.gd itself.
func usesGraphics(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		return false
	}
	mod, err := modfile.ParseLax(path, data, nil)
	if err != nil {
		// a go.mod that go itself would refuse, but that gd has always
		// taken for a graphics.gd project.
		return true
	}
	if mod.Module != nil && mod.Module.Mod.Path == "graphics.gd" {
		return true
	}
	for _, req := range mod.Require {
		if req.Mod.Path == "graphics.gd" {
			return true
		}
	}
	return false
}

// importsStartup reports whether the Go file imports graphics.gd/startup.
func importsStartup(path string) bool {
	file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
	if err != nil {
		return false
	}
	for _, spec := range file.Imports {
		if path, err := strconv.Unquote(spec.Path.Value); err == nil && path == "graphics.gd/startup" {
			return true
		}
	}
	return false
}
