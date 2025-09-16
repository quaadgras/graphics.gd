/*
func _set_state(data):
	zoom = data.get("zoom", 1.0)
	preferred_color = data.get("my_color", Color.WHITE)
*/

package main

import "graphics.gd/variant/Color"

func EditorPlugin_SetState() {
	var zoom float64
	var preferred_color Color.RGBA
	SetState := func(data map[any]any) {
		zoom = data["zoom"].(float64)
		preferred_color = data["my_color"].(Color.RGBA)
	}
	_ = SetState
	_ = zoom
	_ = preferred_color
}
