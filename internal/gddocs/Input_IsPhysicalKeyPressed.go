/*
[gdscript]
func _input(event):
	if event is InputEventKey and not event.is_echo() and event.is_pressed() and event.physical_keycode == KEY_SPACE:
		pass # Your code here.
[/gdscript]
[csharp]
public override void _Input(InputEvent @event)
{
	if (@event is InputEventKey eventKey && !eventKey.IsEcho() && eventKey.Pressed && eventKey.PhysicalKeycode == Key.Space)
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
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Object"
)

type inputIsPhysicalKeyPressed struct {
	Node.Extension[inputIsPhysicalKeyPressed]
}

func (n inputIsPhysicalKeyPressed) Input(event InputEvent.Instance) {
	key, ok := Object.As[InputEventKey.Instance](event)
	if ok && !event.IsEcho() && event.IsPressed() && key.PhysicalKeycode() == Input.KeySpace {
		// Your code here.
	}
}
