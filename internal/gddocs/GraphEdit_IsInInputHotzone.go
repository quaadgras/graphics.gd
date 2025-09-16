/*
func _is_in_input_hotzone(in_node, in_port, mouse_position):
	var port_size = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
	var port_pos = in_node.get_position() + in_node.get_input_port_position(in_port) - port_size / 2
	var rect = Rect2(port_pos, port_size)

	return rect.has_point(mouse_position)
*/

package main

import (
	"graphics.gd/classdb/GraphNode"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

func GraphEdit_IsInInputHotzone() {
	IsInInputHotzone := func(in Object.Instance, in_port int, mouse_position Vector2.XY) bool {
		in_node := Object.To[GraphNode.Instance](in)
		port_size := Vector2.New(control.GetThemeConstant("port_grab_distance_horizontal"), control.GetThemeConstant("port_grab_distance_vertical"))
		port_pos := Vector2.Sub(Vector2.Add(in_node.AsControl().Position(), in_node.GetInputPortPosition(in_port)), Vector2.DivX(port_size, 2))
		var rect = Rect2.PositionSize{port_pos, port_size}

		return Rect2.HasPoint(rect, mouse_position)
	}
	_ = IsInInputHotzone
}
