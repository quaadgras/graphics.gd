/*
[gdscript]
func _forward_3d_draw_over_viewport(overlay):
	# Draw a circle at the cursor's position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_3d_gui_input(camera, event):
	if event is InputEventMouseMotion:
		# Redraw the viewport when the cursor is moved.
		update_overlays()
		return EditorPlugin.AFTER_GUI_INPUT_STOP
	return EditorPlugin.AFTER_GUI_INPUT_PASS
[/gdscript]
[csharp]
public override void _Forward3DDrawOverViewport(Control viewportControl)
{
	// Draw a circle at the cursor's position.
	viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
}

public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D viewportCamera, InputEvent @event)
{
	if (@event is InputEventMouseMotion)
	{
		// Redraw the viewport when the cursor is moved.
		UpdateOverlays();
		return EditorPlugin.AfterGuiInput.Stop;
	}
	return EditorPlugin.AfterGuiInput.Pass;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Object"
)

func EditorPlugin_Forward3dDrawOverViewport() {
	Forward3dDrawOverViewport := func(overlay Control.Instance) {
		// Draw a circle at cursor position.
		overlay.AsCanvasItem().DrawCircle(overlay.AsCanvasItem().GetLocalMousePosition(), 64, Color.W3C.White)
	}
	Forward3dGuiInput := func(viewport_camera Camera3D.Instance, event InputEvent.Instance) EditorPlugin.AfterGUIInput {
		if Object.Is[InputEventMouseMotion.Instance](event) {
			// Redraw viewport when cursor is moved.
			editorPlugin.UpdateOverlays()
			return EditorPlugin.AfterGuiInputStop
		}
		return EditorPlugin.AfterGuiInputPass
	}
	_ = Forward3dDrawOverViewport
	_ = Forward3dGuiInput
}
