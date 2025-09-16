/*
[gdscript]
func _gui_input(event):
	if event is InputEventMouseButton:
		if event.button_index == MOUSE_BUTTON_LEFT and event.pressed:
			print("I've been clicked D:")
[/gdscript]
[csharp]
public override void _GuiInput(InputEvent @event)
{
	if (@event is InputEventMouseButton mb)
	{
		if (mb.ButtonIndex == MouseButton.Left && mb.Pressed)
		{
			GD.Print("I've been clicked D:");
		}
	}
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/variant/Object"
)

func Control_GuiInput() {
	GuiInput := func(event InputEvent.Instance) {
		if event, ok := Object.As[InputEventMouseButton.Instance](event); ok {
			if event.ButtonIndex() == Input.MouseButtonLeft && event.AsInputEvent().IsPressed() {
				fmt.Println("I've been clicked D:")
			}
		}
	}
	_ = GuiInput
}
