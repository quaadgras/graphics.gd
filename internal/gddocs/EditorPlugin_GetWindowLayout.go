/*
func _get_window_layout(configuration):
	configuration.set_value("MyPlugin", "window_position", $Window.position)
	configuration.set_value("MyPlugin", "icon_color", $Icon.modulate)
*/

package main

import "graphics.gd/classdb/ConfigFile"

func EditorPlugin_GetWindowLayout() {
	GetWindowLayout := func(configuration ConfigFile.Instance) {
		configuration.SetValue("MyPlugin", "window_position", window.Position())
		configuration.SetValue("MyPlugin", "icon_color", control.AsCanvasItem().Modulate())
	}
	_ = GetWindowLayout
}
