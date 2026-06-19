/*
[gdscript]
var connection = get_closest_connection_at_point(mouse_event.get_position())
[/gdscript]
*/

package main

import (
	"graphics.gd/classdb/GraphEdit"
	"graphics.gd/variant/Vector2"
)

func ExampleGraphEditClosestConnection(self GraphEdit.Instance, mousePosition Vector2.XY) {
	var connection = self.GetClosestConnectionAtPoint(mousePosition)
	_ = connection
}
