/*
[gdscript]
var polygon_path_finder = PolygonPathFinder.new()
var points = [Vector2(0.0, 0.0), Vector2(1.0, 0.0), Vector2(0.0, 1.0)]
var connections = [0, 1, 1, 2, 2, 0]
polygon_path_finder.setup(points, connections)
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
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/PolygonPathFinder"
	"graphics.gd/variant/Vector2"
)

func PolygonPathFinder_Setup() {
	var polygon_path_finder = PolygonPathFinder.New()
	var points = []Vector2.XY{{0.0, 0.0}, {1.0, 0.0}, {0.0, 1.0}}
	var connections = []int32{0, 1, 1, 2, 2, 0}
	polygon_path_finder.Setup(points, connections)
}
