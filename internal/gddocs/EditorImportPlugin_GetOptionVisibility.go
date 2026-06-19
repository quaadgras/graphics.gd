/*
[gdscript]
func _get_option_visibility(path, option_name, options):
	# Only show the lossy quality setting if the compression mode is set to "Lossy".
	if option_name == "compress/lossy_quality" and options.has("compress/mode"):
		return int(options["compress/mode"]) == COMPRESS_LOSSY # This is a constant that you set

	return true
[/gdscript]
[csharp]
public override bool _GetOptionVisibility(string path, StringName optionName, Godot.Collections.Dictionary options)
{
	// Only show the lossy quality setting if the compression mode is set to "Lossy".
	if (optionName == "compress/lossy_quality" && options.ContainsKey("compress/mode"))
	{
		return (int)options["compress/mode"] == CompressLossy; // This is a constant you set
	}

	return true;
}
[/csharp]
*/

package main

import "graphics.gd/classdb/EditorImportPlugin"

const compressLossy = 0 // This is a constant that you set.

type editorImportPluginOptionVisibility struct {
	EditorImportPlugin.Extension[editorImportPluginOptionVisibility]
}

func (n editorImportPluginOptionVisibility) GetOptionVisibility(path string, optionName string, options map[string]any) bool {
	// Only show the lossy quality setting if the compression mode is set to "Lossy".
	if optionName == "compress/lossy_quality" {
		if mode, ok := options["compress/mode"]; ok {
			v, _ := mode.(int64)
			return int(v) == compressLossy
		}
	}
	return true
}
