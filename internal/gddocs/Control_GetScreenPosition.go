/*
popup_menu.position = get_screen_position() + get_screen_transform().basis_xform(get_local_mouse_position())

# The above code is equivalent to:
popup_menu.position = get_screen_transform() * get_local_mouse_position()

popup_menu.reset_size()
popup_menu.popup()
*/

package main

import (
	"graphics.gd/classdb/PopupMenu"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

var popup_menu PopupMenu.Instance

func Control_GetScreenPosition() {
	popup_menu.AsWindow().SetPosition(Vector2i.From(Vector2.Add(control.GetScreenPosition(),
		Transform2D.BasisTransform(
			control.AsCanvasItem().GetScreenTransform(), control.AsCanvasItem().GetLocalMousePosition()))))

	// The above code is equivalent to:
	popup_menu.AsWindow().SetPosition(Vector2i.From(Vector2.Add(control.GetScreenPosition(), control.AsCanvasItem().GetLocalMousePosition())))

	popup_menu.AsWindow().ResetSize()
	popup_menu.AsWindow().Popup()
}
