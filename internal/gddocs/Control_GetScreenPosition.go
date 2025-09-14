/*
popup_menu.position = get_screen_position() + get_local_mouse_position()
popup_menu.reset_size()
popup_menu.popup()
*/

package main

import (
	"graphics.gd/classdb/PopupMenu"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

var popup_menu PopupMenu.Instance

func Control_GetScreenPosition() {
	popup_menu.AsWindow().SetPosition(Vector2i.From(Vector2.Add(control.GetScreenPosition(), control.AsCanvasItem().GetLocalMousePosition())))
	popup_menu.AsWindow().ResetSize()
	popup_menu.AsWindow().Popup()
}
