/*
[gdscript]
var editor_settings = EditorInterface.get_editor_settings()
[/gdscript]
[csharp]
// In C# you can access it via the static Singleton property.
EditorSettings settings = EditorInterface.Singleton.GetEditorSettings();
[/csharp]
*/

package main

import "graphics.gd/classdb/EditorInterface"

func ExampleEditorInterface() {
	var editor_settings = EditorInterface.GetEditorSettings()
	_ = editor_settings
}
