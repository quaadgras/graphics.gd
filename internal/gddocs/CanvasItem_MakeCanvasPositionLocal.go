/*
var viewport_point = get_global_transform_with_canvas() * local_point
*/

package main

import (
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Vector2"
)

var local_point Vector2.XY

func CanvasItem_MakeCanvasPositionLocal() {
	var viewport_point = Transform2D.Vector(local_point, canvas_item.GetGlobalTransformWithCanvas())
	_ = viewport_point
}
