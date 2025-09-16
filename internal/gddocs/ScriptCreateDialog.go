/*
[gdscript]
func _ready():
	var dialog = ScriptCreateDialog.new();
	dialog.config("Node", "res://new_node.gd") # For in-engine types.
	dialog.config("\"res://base_node.gd\"", "res://derived_node.gd") # For script types.
	dialog.popup_centered()
[/gdscript]
[csharp]
public override void _Ready()
{
	var dialog = new ScriptCreateDialog();
	dialog.Config("Node", "res://NewNode.cs"); // For in-engine types.
	dialog.Config("\"res://BaseNode.cs\"", "res://DerivedNode.cs"); // For script types.
	dialog.PopupCentered();
}
[/csharp]
*/

package main

import "graphics.gd/classdb/ScriptCreateDialog"

func ExampleScriptCreateDialog() {
	var dialog = ScriptCreateDialog.New()
	dialog.Config("Node", "res://new_node.gd")                       // For in-engine types.
	dialog.Config("\"res://base_node.gd\"", "res://derived_node.gd") // For script types.
	dialog.AsWindow().PopupCentered()
}
