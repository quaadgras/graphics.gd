/*
[gdscript]
var cancel_event = InputEventAction.new()
cancel_event.action = "ui_cancel"
cancel_event.pressed = true
Input.parse_input_event(cancel_event)
[/gdscript]
[csharp]
var cancelEvent = new InputEventAction();
cancelEvent.Action = "ui_cancel";
cancelEvent.Pressed = true;
Input.ParseInputEvent(cancelEvent);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEventAction"
)

func Input_ParseInputEvent() {
	var cancelEvent = InputEventAction.New()
	cancelEvent.SetAction("ui_cancel")
	cancelEvent.SetPressed(true)
	Input.ParseInputEvent(cancelEvent.AsInputEvent())
}
