/*
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(1, 1))
astar.add_point(2, Vector2(0, 5))
astar.connect_points(1, 2, false)
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(1, 1));
astar.AddPoint(2, new Vector2(0, 5));
astar.ConnectPoints(1, 2, false);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar2D"
	"graphics.gd/variant/Vector2"
)

func AStar2D_ConnectPoints() {
	var astar = AStar2D.New()
	astar.AddPoint(1, Vector2.New(1, 1))
	astar.AddPoint(2, Vector2.New(0, 5))
	AStar2D.Expanded(astar).ConnectPoints(1, 2, false)
}
