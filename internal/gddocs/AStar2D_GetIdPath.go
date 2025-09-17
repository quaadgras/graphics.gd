/*
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 1), 1) # Default weight is 1
astar.add_point(3, Vector2(1, 1))
astar.add_point(4, Vector2(2, 0))

astar.connect_points(1, 2, false)
astar.connect_points(2, 3, false)
astar.connect_points(4, 3, false)
astar.connect_points(1, 4, false)

var res = astar.get_id_path(1, 3) # Returns [1, 2, 3]
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 1), 1); // Default weight is 1
astar.AddPoint(3, new Vector2(1, 1));
astar.AddPoint(4, new Vector2(2, 0));

astar.ConnectPoints(1, 2, false);
astar.ConnectPoints(2, 3, false);
astar.ConnectPoints(4, 3, false);
astar.ConnectPoints(1, 4, false);
long[] res = astar.GetIdPath(1, 3); // Returns [1, 2, 3]
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar2D"
	"graphics.gd/variant/Vector2"
)

func AStar2D_GetIdPath() {
	var astar = AStar2D.New()
	astar.AddPoint(1, Vector2.New(0, 0))
	astar.MoreArgs().AddPoint(2, Vector2.New(0, 1), 1) // Default weight is 1
	astar.AddPoint(3, Vector2.New(1, 1))
	astar.AddPoint(4, Vector2.New(2, 0))

	astar.MoreArgs().ConnectPoints(1, 2, false)
	astar.MoreArgs().ConnectPoints(2, 3, false)
	astar.MoreArgs().ConnectPoints(4, 3, false)
	astar.MoreArgs().ConnectPoints(1, 4, false)

	var res = astar.GetIdPath(1, 3) // Returns [1, 2, 3]
	_ = res
}
