/*
var plugin_control

func _enter_tree():
	plugin_control = preload("my_plugin_control.tscn").instantiate()
	EditorInterface.get_editor_main_screen().add_child(plugin_control)
	plugin_control.hide()

func _has_main_screen():
	return true

func _make_visible(visible):
	plugin_control.visible = visible

func _get_plugin_name():
	return "My Super Cool Plugin 3000"

func _get_plugin_icon():
	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")
*/

package main

import (
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture2D"
)

type EditorPluginWithMainScreen struct {
	EditorPlugin.Extension[EditorPluginWithMainScreen]

	pluginControl Control.Instance
}

func (e *EditorPluginWithMainScreen) EnterTree() {
	e.pluginControl = Resource.Load[PackedScene.Is[Control.Instance]]("my_plugin_control.tscn").Instantiate()
	EditorInterface.GetEditorMainScreen().AsNode().AddChild(e.pluginControl.AsNode())
	e.pluginControl.AsCanvasItem().Hide()
}

func (e *EditorPluginWithMainScreen) HasMainScreen() bool { return true }

func (e *EditorPluginWithMainScreen) MakeVisible(visible bool) {
	e.pluginControl.AsCanvasItem().SetVisible(visible)
}

func (e *EditorPluginWithMainScreen) GetPluginName() string { return "My Super Cool Plugin 3000" }

func (e *EditorPluginWithMainScreen) GetPluginIcon() Texture2D.Instance {
	return EditorInterface.GetEditorTheme().GetIcon("Node", "EditorIcons")
}
