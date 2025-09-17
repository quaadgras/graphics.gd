/*
func _set_window_layout(configuration):
	$Window.position = configuration.get_value("MyPlugin", "window_position", Vector2())
	$Icon.modulate = configuration.get_value("MyPlugin", "icon_color", Color.WHITE)
*/

package main

import (
	"graphics.gd/classdb/ConfigFile"
	"graphics.gd/classdb/TextureRect"
	"graphics.gd/classdb/Window"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2i"
)

var window Window.Instance
var textureRect TextureRect.Instance

func EditorPlugin_SetWindowLayout() {
	SetWindowLayout := func(configuration ConfigFile.Instance) {
		window.SetPosition(configuration.MoreArgs().GetValue("MyPlugin", "window_position", Vector2i.Zero).(Vector2i.XY))
		textureRect.AsCanvasItem().SetModulate(configuration.MoreArgs().GetValue("MyPlugin", "icon_color", Color.W3C.White).(Color.RGBA))
	}
	_ = SetWindowLayout
}
