/*
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 5))
astar.connect_points(1, 2)
var res = astar.get_closest_position_in_segment(Vector2(3, 3)) # Returns (0, 3)
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 5));
astar.ConnectPoints(1, 2);
Vector2 res = astar.GetClosestPositionInSegment(new Vector2(3, 3)); // Returns (0, 3)
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar2D"
	"graphics.gd/variant/Vector2"
)

func AStar2D_GetClosestPositionInSegment() {
	var astar = AStar2D.New()
	astar.AddPoint(1, Vector2.New(0, 0))
	astar.AddPoint(2, Vector2.New(0, 5))
	astar.ConnectPoints(1, 2)
	var res = astar.GetClosestPositionInSegment(Vector2.New(3, 3)) // Returns (0, 3)
	_ = res
}
