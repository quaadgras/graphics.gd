/*
Rect2(Vector2(), get_size())
*/

package main

import (
	"graphics.gd/classdb/BitMap"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

var bitmap BitMap.Instance

func BitMap_OpaqueToPolygons() {
	var size = Rect2.PositionSize{Vector2.Zero, Vector2.From(bitmap.GetSize())}
	_ = size
}
