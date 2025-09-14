/*
[gdscript]
# Set region, using Path2D node.
DisplayServer.window_set_mouse_passthrough($Path2D.curve.get_baked_points())

# Set region, using Polygon2D node.
DisplayServer.window_set_mouse_passthrough($Polygon2D.polygon)

# Reset region to default.
DisplayServer.window_set_mouse_passthrough([])
[/gdscript]
[csharp]
// Set region, using Path2D node.
DisplayServer.WindowSetMousePassthrough(GetNode<Path2D>("Path2D").Curve.GetBakedPoints());

// Set region, using Polygon2D node.
DisplayServer.WindowSetMousePassthrough(GetNode<Polygon2D>("Polygon2D").Polygon);

// Reset region to default.
DisplayServer.WindowSetMousePassthrough([]);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Path2D"
	"graphics.gd/classdb/Polygon2D"
)

var path2d Path2D.Instance
var polygon2d Polygon2D.Instance

func DisplayServer_WindowSetMousePassthrough() {
	// Set region, using Path2D node.
	DisplayServer.WindowSetMousePassthrough(path2d.Curve().GetBakedPoints(), 0)

	// Set region, using Polygon2D node.
	DisplayServer.WindowSetMousePassthrough(polygon2d.Polygon(), 0)

	// Reset region to default.
	DisplayServer.WindowSetMousePassthrough(nil, 0)
}
