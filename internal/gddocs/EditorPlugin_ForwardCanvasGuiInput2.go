/*
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_canvas_gui_input(event):
	if (event is InputEventMouseMotion):
		return true
	return false
[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override bool _ForwardCanvasGuiInput(InputEvent @event)
{
	if (@event is InputEventMouseMotion)
	{
		return true;
	}
	return false;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/variant/Object"
)

type editorPluginForwardCanvasGuiInput struct {
	EditorPlugin.Extension[editorPluginForwardCanvasGuiInput]
}

// Consumes InputEventMouseMotion and forwards other InputEvent types.
func (n editorPluginForwardCanvasGuiInput) ForwardCanvasGuiInput(event InputEvent.Instance) bool {
	if _, ok := Object.As[InputEventMouseMotion.Instance](event); ok {
		return true
	}
	return false
}
