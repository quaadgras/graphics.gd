/*
[gdscript]
var polygon_path_finder = PolygonPathFinder.new()
var points = [Vector2(0.0, 0.0), Vector2(1.0, 0.0), Vector2(0.0, 1.0)]
var connections = [0, 1, 1, 2, 2, 0]
polygon_path_finder.setup(points, connections)
print(polygon_path_finder.is_point_inside(Vector2(0.2, 0.2))) # Prints true
print(polygon_path_finder.is_point_inside(Vector2(1.0, 1.0))) # Prints false
[/gdscript]
[csharp]
var polygonPathFinder = new PolygonPathFinder();
Vector2[] points =
[
	new Vector2(0.0f, 0.0f),
	new Vector2(1.0f, 0.0f),
	new Vector2(0.0f, 1.0f)
];
int[] connections = [0, 1, 1, 2, 2, 0];
polygonPathFinder.Setup(points, connections);
GD.Print(polygonPathFinder.IsPointInside(new Vector2(0.2f, 0.2f))); // Prints True
GD.Print(polygonPathFinder.IsPointInside(new Vector2(1.0f, 1.0f))); // Prints False
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/PolygonPathFinder"
	"graphics.gd/variant/Vector2"
)

func PolygonPathFinder_IsPointInside() {
	var polygonPathFinder = PolygonPathFinder.New()
	points := []Vector2.XY{{0.0, 0.0}, {1.0, 0.0}, {0.0, 1.0}}
	connections := []int32{0, 1, 1, 2, 2, 0}
	polygonPathFinder.Setup(points, connections)
	fmt.Println(polygonPathFinder.IsPointInside(Vector2.New(0.2, 0.2))) // Prints true
	fmt.Println(polygonPathFinder.IsPointInside(Vector2.New(1.0, 1.0))) // Prints false
}
