/*
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_3d_gui_input(camera, event):
	return EditorPlugin.AFTER_GUI_INPUT_STOP if event is InputEventMouseMotion else EditorPlugin.AFTER_GUI_INPUT_PASS
[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)
{
	return @event is InputEventMouseMotion ? EditorPlugin.AfterGuiInput.Stop : EditorPlugin.AfterGuiInput.Pass;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/variant/Object"
)

type editorPluginForward3DGuiInput struct {
	EditorPlugin.Extension[editorPluginForward3DGuiInput]
}

// Consumes InputEventMouseMotion and forwards other InputEvent types.
func (n editorPluginForward3DGuiInput) Forward3dGuiInput(camera Camera3D.Instance, event InputEvent.Instance) EditorPlugin.AfterGUIInput {
	if _, ok := Object.As[InputEventMouseMotion.Instance](event); ok {
		return EditorPlugin.AfterGuiInputStop
	}
	return EditorPlugin.AfterGuiInputPass
}
