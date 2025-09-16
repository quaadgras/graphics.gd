/*
[gdscript]
const MyInspectorPlugin = preload("res://addons/your_addon/path/to/your/script.gd")
var inspector_plugin = MyInspectorPlugin.new()

func _enter_tree():
	add_inspector_plugin(inspector_plugin)

func _exit_tree():
	remove_inspector_plugin(inspector_plugin)
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/EditorInspectorPlugin"
	"graphics.gd/classdb/EditorPlugin"
)

type MyInspectorPlugin struct {
	EditorInspectorPlugin.Extension[MyInspectorPlugin]
}

type MyEditorPlugin struct {
	EditorPlugin.Extension[MyEditorPlugin]

	inspector_plugin *MyInspectorPlugin
}

func (e *MyEditorPlugin) EnterTree() {
	e.inspector_plugin = new(MyInspectorPlugin)
	e.AsEditorPlugin().AddInspectorPlugin(e.inspector_plugin.AsEditorInspectorPlugin())
}

func (e *MyEditorPlugin) ExitTree() {
	e.AsEditorPlugin().RemoveInspectorPlugin(e.inspector_plugin.AsEditorInspectorPlugin())
}
