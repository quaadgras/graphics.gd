/*
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(0, 0, 0))
astar.add_point(2, Vector3(0, 5, 0))
astar.connect_points(1, 2)
var res = astar.get_closest_position_in_segment(Vector3(3, 3, 0)) # Returns (0, 3, 0)
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(0, 0, 0));
astar.AddPoint(2, new Vector3(0, 5, 0));
astar.ConnectPoints(1, 2);
Vector3 res = astar.GetClosestPositionInSegment(new Vector3(3, 3, 0)); // Returns (0, 3, 0)
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar3D"
	"graphics.gd/variant/Vector3"
)

func AStar3D_GetClosestPositionInSegment() {
	var astar = AStar3D.New()
	astar.AddPoint(1, Vector3.New(0, 0, 0))
	astar.AddPoint(2, Vector3.New(0, 5, 0))
	astar.ConnectPoints(1, 2)
	var res = astar.GetClosestPositionInSegment(Vector3.New(3, 3, 0)) // Returns (0, 3, 0)
	_ = res
}
