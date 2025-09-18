/*
[gdscript]
func _input(event):
	if event is InputEventMouseButton and event.pressed and event.button_index == MOUSE_BUTTON_LEFT:
		if get_rect().has_point(to_local(event.position)):
			print("A click!")
[/gdscript]
[csharp]
public override void _Input(InputEvent @event)
{
	if (@event is InputEventMouseButton inputEventMouse)
	{
		if (inputEventMouse.Pressed && inputEventMouse.ButtonIndex == MouseButton.Left)
		{
			if (GetRect().HasPoint(ToLocal(inputEventMouse.Position)))
			{
				GD.Print("A click!");
			}
		}
	}
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventMouseButton"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Rect2"
)

func Sprite2D_GetRect() {
	Input := func(event InputEvent.Instance) {
		if inputEventMouse, ok := Object.As[InputEventMouseButton.Instance](event); ok {
			if Rect2.HasPoint(sprite.GetRect(), sprite.AsNode2D().ToLocal(inputEventMouse.AsInputEventMouse().Position())) {
				fmt.Println("A click!")
			}
		}
	}
	_ = Input
}
