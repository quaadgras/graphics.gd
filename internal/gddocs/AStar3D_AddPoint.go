/*
[gdscript]
var astar = AStar3D.new()
astar.add_point(1, Vector3(1, 0, 0), 4) # Adds the point (1, 0, 0) with weight_scale 4 and id 1
[/gdscript]
[csharp]
var astar = new AStar3D();
astar.AddPoint(1, new Vector3(1, 0, 0), 4); // Adds the point (1, 0, 0) with weight_scale 4 and id 1
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar3D"
	"graphics.gd/variant/Vector3"
)

func AStar3D_AddPoint() {
	var astar = AStar3D.New()
	astar.MoreArgs().AddPoint(1, Vector3.New(1, 0, 0), 4) // Adds the point (1, 0, 0) with weight_scale 4 and id 1
}
