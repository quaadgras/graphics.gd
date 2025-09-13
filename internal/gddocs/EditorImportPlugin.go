/*
[gdscript]
@tool
extends EditorImportPlugin

func _get_importer_name():
    return "my.special.plugin"

func _get_visible_name():
    return "Special Mesh"

func _get_recognized_extensions():
    return ["special", "spec"]

func _get_save_extension():
    return "mesh"

func _get_resource_type():
    return "Mesh"

func _get_preset_count():
    return 1

func _get_preset_name(preset_index):
    return "Default"

func _get_import_options(path, preset_index):
    return [{"name": "my_option", "default_value": false}]

func _import(source_file, save_path, options, platform_variants, gen_files):
    var file = FileAccess.open(source_file, FileAccess.READ)
    if file == null:
        return FAILED
    var mesh = ArrayMesh.new()
    # Fill the Mesh with data read in "file", left as an exercise to the reader.

    var filename = save_path + "." + _get_save_extension()
    return ResourceSaver.save(mesh, filename)
[/gdscript]
[csharp]
using Godot;

public partial class MySpecialPlugin : EditorImportPlugin
{
    public override string _GetImporterName()
    {
        return "my.special.plugin";
    }

    public override string _GetVisibleName()
    {
        return "Special Mesh";
    }

    public override string[] _GetRecognizedExtensions()
    {
        return ["special", "spec"];
    }

    public override string _GetSaveExtension()
    {
        return "mesh";
    }

    public override string _GetResourceType()
    {
        return "Mesh";
    }

    public override int _GetPresetCount()
    {
        return 1;
    }

    public override string _GetPresetName(int presetIndex)
    {
        return "Default";
    }

    public override Godot.Collections.Array<Godot.Collections.Dictionary> _GetImportOptions(string path, int presetIndex)
    {
        return
        [
            new Godot.Collections.Dictionary
            {
                { "name", "myOption" },
                { "default_value", false },
            },
        ];
    }

    public override Error _Import(string sourceFile, string savePath, Godot.Collections.Dictionary options, Godot.Collections.Array<string> platformVariants, Godot.Collections.Array<string> genFiles)
    {
        using var file = FileAccess.Open(sourceFile, FileAccess.ModeFlags.Read);
        if (file.GetError() != Error.Ok)
        {
            return Error.Failed;
        }

        var mesh = new ArrayMesh();
        // Fill the Mesh with data read in "file", left as an exercise to the reader.
        string filename = $"{savePath}.{_GetSaveExtension()}";
        return ResourceSaver.Save(mesh, filename);
    }
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/ArrayMesh"
	"graphics.gd/classdb/EditorImportPlugin"
	"graphics.gd/classdb/FileAccess"
	"graphics.gd/classdb/ResourceSaver"
)

type MySpecialPlugin struct {
	EditorImportPlugin.Extension[MySpecialPlugin]
}

func (m *MySpecialPlugin) GetImporterName() string              { return "my.special.plugin" }
func (m *MySpecialPlugin) GetVisibleName() string               { return "Special Mesh" }
func (m *MySpecialPlugin) GetRecognizedExtensions() []string    { return []string{"special", "spec"} }
func (m *MySpecialPlugin) GetSaveExtension() string             { return "mesh" }
func (m *MySpecialPlugin) GetResourceType() string              { return "Mesh" }
func (m *MySpecialPlugin) GetPresetCount() int                  { return 1 }
func (m *MySpecialPlugin) GetPresetName(presetIndex int) string { return "Default" }
func (m *MySpecialPlugin) GetImportOptions(path string, presetIndex int) []map[any]any {
	return []map[any]any{
		{
			"name":          "my_option",
			"default_value": false,
		},
	}
}
func (m *MySpecialPlugin) Import(sourceFile, savePath string, options map[any]any, platformVariants, genFiles []string) error {
	var file = FileAccess.Open(sourceFile, FileAccess.Read)
	if err := file.GetError(); err != nil {
		return err
	}
	var mesh = ArrayMesh.New()
	// Fill the Mesh with data read in "file", left as an exercise to the reader.
	var filename = savePath + "." + m.GetSaveExtension()
	return ResourceSaver.Save(mesh.AsResource(), filename, 0)
}
