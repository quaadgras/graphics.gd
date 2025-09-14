/*
[gdscript]
func _forward_canvas_draw_over_viewport(overlay):
    # Draw a circle at cursor position.
    overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_canvas_gui_input(event):
    if event is InputEventMouseMotion:
        # Redraw viewport when cursor is moved.
        update_overlays()
        return true
    return false
[/gdscript]
[csharp]
public override void _ForwardCanvasDrawOverViewport(Control viewportControl)
{
    // Draw a circle at cursor position.
    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
}

public override bool _ForwardCanvasGuiInput(InputEvent @event)
{
    if (@event is InputEventMouseMotion)
    {
        // Redraw viewport when cursor is moved.
        UpdateOverlays();
        return true;
    }
    return false;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseMotion"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Object"
)

var editorPlugin EditorPlugin.Instance

func EditorPlugin_ForwardCanvasDrawOverViewport() {
	ForwardCanvasDrawOverViewport := func(overlay Control.Instance) {
		// Draw a circle at cursor position.
		overlay.AsCanvasItem().DrawCircle(overlay.AsCanvasItem().GetLocalMousePosition(), 64, Color.W3C.White)
	}
	ForwardCanvasGuiInput := func(event InputEvent.Instance) bool {
		if Object.Is[InputEventMouseMotion.Instance](event) {
			// Redraw viewport when cursor is moved.
			// Note: UpdateOverlays is a method of EditorPlugin, so you would call it on the instance of your plugin.
			// For this example, we'll just comment it out as we don't have the context here.
			editorPlugin.UpdateOverlays()
			return true
		}
		return false
	}
	_ = ForwardCanvasDrawOverViewport
	_ = ForwardCanvasGuiInput
}
