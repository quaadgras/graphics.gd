/*
[gdscript]
@tool
extends EditorTranslationParserPlugin

func _parse_file(path):
	var ret: Array[PackedStringArray] = []
	var file = FileAccess.open(path, FileAccess.READ)
	var text = file.get_as_text()
	var split_strs = text.split(",", false)
	for s in split_strs:
		ret.append(PackedStringArray([s]))
		#print("Extracted string: " + s)

	return ret

func _get_recognized_extensions():
	return ["csv"]
[/gdscript]
[csharp]
using Godot;

[Tool]
public partial class CustomParser : EditorTranslationParserPlugin
{
	public override Godot.Collections.Array<string[]> _ParseFile(string path)
	{
		Godot.Collections.Array<string[]> ret;
		using var file = FileAccess.Open(path, FileAccess.ModeFlags.Read);
		string text = file.GetAsText();
		string[] splitStrs = text.Split(",", allowEmpty: false);
		foreach (string s in splitStrs)
		{
			ret.Add([s]);
			//GD.Print($"Extracted string: {s}");
		}
		return ret;
	}

	public override string[] _GetRecognizedExtensions()
	{
		return ["csv"];
	}
}
[/csharp]
*/

package main

import (
	"strings"

	"graphics.gd/classdb/EditorTranslationParserPlugin"
	"graphics.gd/classdb/FileAccess"
)

type CustomParser struct {
	EditorTranslationParserPlugin.Extension[CustomParser]
}

func (c *CustomParser) ParseFile(path string) [][]string {
	var ret [][]string
	var file = FileAccess.Open(path, FileAccess.Read)
	var text = file.GetAsText()
	for s := range strings.SplitSeq(text, ",") {
		ret = append(ret, []string{s})
	}
	return ret
}

func (c *CustomParser) GetRecognizedExtensions() []string {
	return []string{"csv"}
}
