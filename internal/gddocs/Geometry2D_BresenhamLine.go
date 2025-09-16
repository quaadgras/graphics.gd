/*
func _draw():
	for pixel in Geometry2D.bresenham_line($MarkerA.position, $MarkerB.position):
		draw_rect(Rect2(pixel, Vector2.ONE), Color.WHITE)
*/

package main

import (
	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/Geometry2D"
	"graphics.gd/classdb/Marker2D"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

var markerA Marker2D.Instance
var markerB Marker2D.Instance
var canvas CanvasItem.Instance

func Geometry2D_BresenhamLine() {
	Draw := func() {
		for _, pixel := range Geometry2D.BresenhamLine(Vector2i.From(markerA.AsNode2D().Position()), Vector2i.From(markerB.AsNode2D().Position())) {
			canvas.DrawRect(Rect2.PositionSize{Vector2.From(pixel), Vector2.One}, Color.W3C.White)
		}
	}
	_ = Draw
}
