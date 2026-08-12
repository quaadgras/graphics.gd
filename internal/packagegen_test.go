//go:build !generate

package gd_test

import (
	"path/filepath"
	"strings"
	"testing"

	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/ResourceSaver"
	"graphics.gd/internal/packagegen"
)

func TestGraphicsPackageGeneration(t *testing.T) {
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
		"`gd:\"" + dir + "/main.tscn\"`",
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

// TestGraphicsPackageGenerationProject runs generation against the test
// project's own graphics directory, exercising res:// loading of imported
// resources (icon.svg has no direct loader, only its import remap).
func TestGraphicsPackageGenerationProject(t *testing.T) {
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
