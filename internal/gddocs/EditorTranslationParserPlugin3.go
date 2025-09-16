/*
[gdscript]
func _parse_file(path):
	var res = ResourceLoader.load(path, "Script")
	var text = res.source_code
	# Parsing logic.

func _get_recognized_extensions():
	return ["gd"]
[/gdscript]
[csharp]
public override Godot.Collections.Array<string[]> _ParseFile(string path)
{
	var res = ResourceLoader.Load<Script>(path, "Script");
	string text = res.SourceCode;
	// Parsing logic.
}

public override string[] _GetRecognizedExtensions()
{
	return ["gd"];
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/EditorTranslationParserPlugin"
	"graphics.gd/classdb/ResourceLoader"
	"graphics.gd/classdb/Script"
	"graphics.gd/variant/Object"
)

type MyEditorTranslationParserPlugin struct {
	EditorTranslationParserPlugin.Extension[MyEditorTranslationParserPlugin]
}

func (m *MyEditorTranslationParserPlugin) ParseFile(path string) [][]string {
	var res = Object.To[Script.Instance](ResourceLoader.Load(path, "Script"))
	var text = res.SourceCode()
	// Parsing logic.
	_ = text
	return nil
}

func (m *MyEditorTranslationParserPlugin) GetRecognizedExtensions() []string {
	return []string{"gd"}
}
