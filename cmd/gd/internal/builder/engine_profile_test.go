package builder

import (
	"os"
	"path/filepath"
	"slices"
	"testing"

	"graphics.gd/cmd/gd/internal/project"
)

func TestDisabledClasses(t *testing.T) {
	classes := map[string]engineClass{
		"Object":          {},
		"Node":            {Inherits: "Object", Core: true},
		"CanvasItem":      {Inherits: "Node", Core: true},
		"Control":         {Inherits: "CanvasItem", Core: true},
		"Label":           {Inherits: "Control", Core: true},
		"Button":          {Inherits: "Control", Core: true},
		"MenuButton":      {Inherits: "Button", Core: true},
		"ColorPicker":     {Inherits: "Control", Core: true},
		"Node3D":          {Inherits: "Node", Core: true},
		"Camera3D":        {Inherits: "Node3D", Core: true},
		"Resource":        {Inherits: "Object", Core: true},
		"Font":            {Inherits: "Resource", Core: true},
		"FontFile":        {Inherits: "Font", Core: true},
		"Texture2D":       {Inherits: "Resource", Core: true},
		"EditorPlugin":    {Inherits: "Node", Core: false},
		"OS":              {Inherits: "Object", Core: true},
		"RenderingServer": {Inherits: "Object", Core: true},
	}
	dependencies := map[string][]string{"ColorPicker": {"MenuButton"}}
	disabled, dropped := disabledClasses(classes, dependencies, map[string]bool{"Label": true, "ColorPicker": true}, []string{" Texture2D ", ""})
	// Node3D is the root of the disabled 3D subtree, Camera3D under it is
	// implied. Button is kept as MenuButton (a dependency of ColorPicker)
	// inherits it. Font and its inheriters are always kept. Editor classes
	// and classes that are neither nodes nor resources are not the
	// profile's to disable.
	if want := []string{"Node3D"}; !slices.Equal(disabled, want) {
		t.Errorf("disabled %q, want %q", disabled, want)
	}
	if dropped != 2 {
		t.Errorf("dropped %d classes, want 2 (Node3D and Camera3D)", dropped)
	}
}

func TestEngineClasses(t *testing.T) {
	src := t.TempDir()
	write := func(path, content string) {
		path = filepath.Join(src, path)
		os.MkdirAll(filepath.Dir(path), 0755)
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	write("doc/classes/Label.xml", `<?xml version="1.0" encoding="UTF-8" ?>`+"\n"+`<class name="Label" inherits="Control" api_type="core" keywords="text" xmlns:xsi="x">`)
	write("doc/classes/Object.xml", `<class name="Object" xmlns:xsi="x">`)
	write("modules/gltf/doc_classes/GLTFState.xml", `<class name="GLTFState" inherits="Resource" api_type="core">`)
	write("platform/android/doc_classes/EditorExportPlatformAndroid.xml", `<class name="EditorExportPlatformAndroid" inherits="EditorExportPlatform" api_type="editor">`)
	write("scene/gui/color_picker.cpp", "void ColorPicker::_bind_methods() {\n\tADD_CLASS_DEPENDENCY(\"LineEdit\");\n\tADD_CLASS_DEPENDENCY(\"MenuButton\");\n}\nvoid ColorPickerButton::_bind_methods() {\n\tADD_CLASS_DEPENDENCY(\"ColorPicker\");\n}\n")
	classes, dependencies, err := engineClasses(src)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]engineClass{
		"Label":                       {Inherits: "Control", Core: true},
		"Object":                      {Core: true},
		"GLTFState":                   {Inherits: "Resource", Core: true},
		"EditorExportPlatformAndroid": {Inherits: "EditorExportPlatform", Core: false},
	}
	for name, class := range want {
		if classes[name] != class {
			t.Errorf("%s: got %+v, want %+v", name, classes[name], class)
		}
	}
	if !slices.Equal(dependencies["ColorPicker"], []string{"LineEdit", "MenuButton"}) || !slices.Equal(dependencies["ColorPickerButton"], []string{"ColorPicker"}) {
		t.Errorf("dependencies: %v", dependencies)
	}
}

func TestCustomEngine(t *testing.T) {
	dir := t.TempDir()
	defer func(d, g string) { project.Directory, project.GraphicsDirectory = d, g }(project.Directory, project.GraphicsDirectory)
	project.Directory, project.GraphicsDirectory = dir, dir
	write := func(content string) {
		if err := os.WriteFile(filepath.Join(dir, "project.godot"), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	write("[application]\nconfig/name=\"x\"\n")
	if _, ok := customEngine(); ok {
		t.Error("an engine was configured without a [gd] section")
	}
	write("[application]\nconfig/name=\"x\"\n\n[gd]\n\nengine/repository=\"../godot\"\nengine/ref=\"stencil\"\nengine/strip_unused_classes=true\nengine/classes=\"Label, Sprite2D\"\nengine/options=\"optimize=size  lto=full\"\n")
	custom, ok := customEngine()
	if !ok {
		t.Fatal("no engine configured")
	}
	if custom.Repository != "../godot" || custom.Ref != "stencil" || !custom.StripUnusedClasses ||
		!slices.Equal(custom.Classes, []string{"Label", " Sprite2D"}) || !slices.Equal(custom.Options, []string{"optimize=size", "lto=full"}) {
		t.Errorf("parsed %+v", custom)
	}
}
