/*
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(0, 0))
astar.add_point(2, Vector2(0, 1))
astar.add_point(3, Vector2(1, 1))
astar.add_point(4, Vector2(2, 0))

astar.connect_points(1, 2, true)
astar.connect_points(1, 3, true)

var neighbors = astar.get_point_connections(1) # Returns [2, 3]
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(0, 0));
astar.AddPoint(2, new Vector2(0, 1));
astar.AddPoint(3, new Vector2(1, 1));
astar.AddPoint(4, new Vector2(2, 0));

astar.ConnectPoints(1, 2, true);
astar.ConnectPoints(1, 3, true);

long[] neighbors = astar.GetPointConnections(1); // Returns [2, 3]
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar2D"
	"graphics.gd/variant/Vector2"
)

func AStar2D_GetPointConnections() {
	var astar = AStar2D.New()
	astar.AddPoint(1, Vector2.New(0, 0))
	astar.AddPoint(2, Vector2.New(0, 1))
	astar.AddPoint(3, Vector2.New(1, 1))
	astar.AddPoint(4, Vector2.New(2, 0))

	astar.MoreArgs().ConnectPoints(1, 2, true)
	astar.MoreArgs().ConnectPoints(1, 3, true)

	var neighbors = astar.GetPointConnections(1) // Returns [2, 3]
	_ = neighbors
}
