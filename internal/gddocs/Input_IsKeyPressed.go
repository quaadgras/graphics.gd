/*
[gdscript]
func _input(event):
	if event is InputEventKey and not event.is_echo() and event.is_pressed() and event.keycode == KEY_SPACE:
		pass # Your code here.
[/gdscript]
[csharp]
public override void _Input(InputEvent @event)
{
	if (@event is InputEventKey eventKey && !eventKey.IsEcho() && eventKey.Pressed && eventKey.Keycode == Key.Space)
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

type inputIsKeyPressed struct {
	Node.Extension[inputIsKeyPressed]
}

func (n inputIsKeyPressed) Input(event InputEvent.Instance) {
	key, ok := Object.As[InputEventKey.Instance](event)
	if ok && !event.IsEcho() && event.IsPressed() && key.Keycode() == Input.KeySpace {
		// Your code here.
	}
}
