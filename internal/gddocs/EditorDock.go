/*
@tool
extends EditorPlugin

# Dock reference.
var dock

# Plugin initialization.
func _enter_tree():
	dock = EditorDock.new()
	dock.title = "My Dock"
	dock.dock_icon = preload("./dock_icon.png")
	dock.default_slot = EditorDock.DOCK_SLOT_RIGHT_UL
	var dock_content = preload("./dock_content.tscn").instantiate()
	dock.add_child(dock_content)
	add_dock(dock)

# Plugin clean-up.
func _exit_tree():
	remove_dock(dock)
	dock.queue_free()
	dock = null
*/

package main

import (
	"graphics.gd/classdb/EditorDock"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture2D"
)

type MyDockPlugin struct {
	EditorPlugin.Extension[MyDockPlugin]

	dock EditorDock.Instance
}

func (plugin MyDockPlugin) EnterTree() {
	plugin.dock = EditorDock.New()
	plugin.dock.SetTitle("My Dock")
	plugin.dock.SetDockIcon(Resource.Load[Texture2D.Instance]("./dock_icon.png"))
	plugin.dock.SetDefaultSlot(EditorDock.DockSlotRightUl)
	var dock_content = Resource.Load[PackedScene.Instance]("./dock_content.tscn").Instantiate()
	plugin.dock.AsNode().AddChild(dock_content)
	plugin.AsEditorPlugin().AddDock(plugin.dock)
}

func (plugin MyDockPlugin) ExitTree() {
	plugin.AsEditorPlugin().RemoveDock(plugin.dock)
	plugin.dock.AsNode().QueueFree()
	plugin.dock = EditorDock.Nil
}
