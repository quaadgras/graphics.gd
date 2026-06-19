/*
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(1, 1, 0))
astar.add_point(2, Vector3(0, 5, 0))
astar.connect_points(1, 2, false)
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(1, 1, 0));
astar.AddPoint(2, new Vector3(0, 5, 0));
astar.ConnectPoints(1, 2, false);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar3D"
	"graphics.gd/variant/Vector3"
)

func AStar3D_ConnectPoints() {
	var astar = AStar3D.New()
	astar.AddPoint(1, Vector3.New(1, 1, 0))
	astar.AddPoint(2, Vector3.New(0, 5, 0))
	astar.MoreArgs().ConnectPoints(1, 2, false)
}
