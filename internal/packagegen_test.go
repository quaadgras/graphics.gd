//go:build !generate

package gd_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/ResourceSaver"
	"graphics.gd/internal/packagegen"
)

// packagegen walks the graphics directory on the filesystem, it is a build-time
// tool that runs on the host. A running game may have no such filesystem to walk:
// under js/wasm t.TempDir isn't implemented and res:// lives inside the packed
// .pck, and on android the suite runs from the exported .apk, where res:// has no
// OS directory behind it and the app's sandbox has nowhere to save a scene to.
func skipWithoutFilesystem(t *testing.T) {
	t.Helper()
	switch runtime.GOOS {
	case "js", "android":
		t.Skipf("packagegen is a host-side tool, %v has no project directory to generate from", runtime.GOOS)
	}
}

func TestGraphicsPackageGeneration(t *testing.T) {
	skipWithoutFilesystem(t)
	dir := t.TempDir()
	scene := makeSceneMain(t)
	if err := ResourceSaver.Save(scene.AsResource(), filepath.Join(dir, "main.tscn"), 0); err != nil {
		t.Fatal(err)
	}
	source, err := packagegen.Directory(dir, dir+"/", "graphics")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"package graphics",
		"var TSCN = Resource.Library[struct {",
		"PackedScene.Is[SceneMain]",
		// The tag holds a quoted Go string (reflect.StructTag.Get unquotes it),
		// so it has to be compared in that form: on Windows dir contains
		// backslashes, which the generator escapes and Get puts back.
		fmt.Sprintf("`gd:%q`", dir+"/main.tscn"),
		"type SceneMain struct {",
		"Node2D.Instance",
		"Bullets Node2D.Instance",
		"Player  struct {",
		"Area2D.Instance",
		"Sprite Node2D.Instance",
		`"graphics.gd/classdb/Area2D"`,
		`"graphics.gd/classdb/PackedScene"`,
		`"graphics.gd/classdb/Resource"`,
	} {
		if !strings.Contains(source, want) {
			t.Fatalf("generated source is missing %q:\n%s", want, source)
		}
	}
}

// Asset directories are named by artists, not by Go programmers, so the
// package name a directory implies can be a keyword. Generating "package
// interface" makes a file that does not parse, and one such directory used
// to abort generation for the whole project.
func TestGraphicsPackageGenerationKeywordDirectory(t *testing.T) {
	skipWithoutFilesystem(t)
	dir := t.TempDir()
	scene := makeSceneMain(t)
	if err := ResourceSaver.Save(scene.AsResource(), filepath.Join(dir, "main.tscn"), 0); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"interface", "map", "graphics"} {
		source, err := packagegen.Directory(dir, dir+"/", name)
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		want := "package " + name
		if name != "graphics" {
			want += "_"
		}
		if !strings.Contains(source, want+"\n") {
			t.Fatalf("generated source is missing %q:\n%s", want, source)
		}
	}
}

// TestGraphicsPackageGenerationProject runs generation against the test
// project's own graphics directory, exercising res:// loading of imported
// resources (icon.svg has no direct loader, only its import remap).
func TestGraphicsPackageGenerationProject(t *testing.T) {
	skipWithoutFilesystem(t)
	source, err := packagegen.Directory(ProjectSettings.GlobalizePath("res://"), "res://", "graphics")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"package graphics",
		"Icon Texture2D.Instance `gd:\"res://icon.svg\"`",
		"Main PackedScene.Is[SceneMain] `gd:\"res://main.tscn\"`",
		"type SceneMain struct {",
		"Node.Instance `gd:\"res://main.tscn\"`",
		"ActionUI_Left",
		`= "ui_left"`,
		"ActionUI_Accept",
	} {
		if !strings.Contains(source, want) {
			t.Fatalf("generated source is missing %q:\n%s", want, source)
		}
	}
}
