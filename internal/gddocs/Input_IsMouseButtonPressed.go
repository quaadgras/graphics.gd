/*
[gdscript]
func _input(event):
	if event is InputEventMouseButton and event.is_pressed() and event.button_index == MOUSE_BUTTON_LEFT:
		pass # Your code here.
[/gdscript]
[csharp]
public override void _Input(InputEvent @event)
{
	if (@event is InputEventMouseButton eventMouseButton && eventMouseButton.Pressed && eventMouseButton.ButtonIndex == MouseButton.Left)
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
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Object"
)

type inputIsMouseButtonPressed struct {
	Node.Extension[inputIsMouseButtonPressed]
}

func (n inputIsMouseButtonPressed) Input(event InputEvent.Instance) {
	mb, ok := Object.As[InputEventMouseButton.Instance](event)
	if ok && event.IsPressed() && mb.ButtonIndex() == Input.MouseButtonLeft {
		// Your code here.
	}
}
