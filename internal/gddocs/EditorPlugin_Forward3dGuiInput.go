/*
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_3d_gui_input(camera, event):
	return EditorPlugin.AFTER_GUI_INPUT_STOP
[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)
{
	return EditorPlugin.AfterGuiInput.Stop;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/InputEvent"
)

func EditorPlugin_Forward3dGuiInput() {
	Forward3dGuiInput := func(viewport_camera Camera3D.Instance, event InputEvent.Instance) EditorPlugin.AfterGUIInput {
		return EditorPlugin.AfterGuiInputStop
	}
	_ = Forward3dGuiInput
}
