/*
[gdscript]
var polygon = PackedVector2Array([Vector2(0, 0), Vector2(100, 0), Vector2(100, 100), Vector2(0, 100)])
var offset = Vector2(50, 50)
polygon = Transform2D(0, offset) * polygon
print(polygon) # Prints [(50.0, 50.0), (150.0, 50.0), (150.0, 150.0), (50.0, 150.0)]
[/gdscript]
[csharp]
Vector2[] polygon = [new Vector2(0, 0), new Vector2(100, 0), new Vector2(100, 100), new Vector2(0, 100)];
var offset = new Vector2(50, 50);
polygon = new Transform2D(0, offset) * polygon;
GD.Print((Variant)polygon); // Prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Vector2"
)

func Geometry2D_OffsetPolygon() {
	var polygon = []Vector2.XY{Vector2.New(0, 0), Vector2.New(100, 0), Vector2.New(100, 100), Vector2.New(0, 100)}
	var offset = Vector2.New(50, 50)
	for i := range polygon {
		polygon[i] = Transform2D.Vector(polygon[i], Transform2D.OriginXY{Origin: offset})
	}
	fmt.Println(polygon) // Prints [(50, 50), (150, 50), (150, 150), (50, 150)]
}
