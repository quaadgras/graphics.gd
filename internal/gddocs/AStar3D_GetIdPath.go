/*
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(0, 0, 0))
astar.add_point(2, Vector3(0, 1, 0), 1) # Default weight is 1
astar.add_point(3, Vector3(1, 1, 0))
astar.add_point(4, Vector3(2, 0, 0))

astar.connect_points(1, 2, false)
astar.connect_points(2, 3, false)
astar.connect_points(4, 3, false)
astar.connect_points(1, 4, false)

var res = astar.get_id_path(1, 3) # Returns [1, 2, 3]
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(0, 0, 0));
astar.AddPoint(2, new Vector3(0, 1, 0), 1); // Default weight is 1
astar.AddPoint(3, new Vector3(1, 1, 0));
astar.AddPoint(4, new Vector3(2, 0, 0));
astar.ConnectPoints(1, 2, false);
astar.ConnectPoints(2, 3, false);
astar.ConnectPoints(4, 3, false);
astar.ConnectPoints(1, 4, false);
long[] res = astar.GetIdPath(1, 3); // Returns [1, 2, 3]
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar3D"
	"graphics.gd/variant/Vector3"
)

func AStar3D_GetIdPath() {
	var astar = AStar3D.New()
	astar.AddPoint(1, Vector3.New(0, 0, 0))
	astar.MoreArgs().AddPoint(2, Vector3.New(0, 1, 0), 1) // Default weight is 1
	astar.AddPoint(3, Vector3.New(1, 1, 0))
	astar.AddPoint(4, Vector3.New(2, 0, 0))
	astar.MoreArgs().ConnectPoints(1, 2, false)
	astar.MoreArgs().ConnectPoints(2, 3, false)
	astar.MoreArgs().ConnectPoints(4, 3, false)
	astar.MoreArgs().ConnectPoints(1, 4, false)
	var res = astar.GetIdPath(1, 3) // Returns [1, 2, 3]
	_ = res
}
