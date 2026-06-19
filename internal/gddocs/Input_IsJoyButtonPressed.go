/*
[gdscript]
func _input(event):
	if event is InputEventJoypadButton and event.is_pressed() and event.button_index == JOY_BUTTON_A:
		pass # Your code here.
[/gdscript]
[csharp]
public override void _Input(InputEvent @event)
{
	if (@event is InputEventJoypadButton eventButton && eventButton.Pressed && eventButton.ButtonIndex == JoyButton.A)
	{
		// Your code here.
	}
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventJoypadButton"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Object"
)

type inputIsJoyButtonPressed struct {
	Node.Extension[inputIsJoyButtonPressed]
}

func (n inputIsJoyButtonPressed) Input(event InputEvent.Instance) {
	jb, ok := Object.As[InputEventJoypadButton.Instance](event)
	if ok && event.IsPressed() && jb.ButtonIndex() == Input.JoyButtonA {
		// Your code here.
	}
}
