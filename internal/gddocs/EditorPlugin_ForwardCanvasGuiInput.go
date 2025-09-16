/*
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_canvas_gui_input(event):
	return true
[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override bool ForwardCanvasGuiInput(InputEvent @event)
{
	return true;
}
[/csharp]
*/

package main

import "graphics.gd/classdb/InputEvent"

func EditorPlugin_ForwardCanvasGuiInput() {
	ForwardCanvasGuiInput := func(event InputEvent.Instance) bool {
		return true // Prevents the InputEvent from reaching other Editor classes.
	}
	_ = ForwardCanvasGuiInput
}
