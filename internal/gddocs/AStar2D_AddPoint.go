/*
[gdscript]
var astar = AStar2D.new()
astar.add_point(1, Vector2(1, 0), 4) # Adds the point (1, 0) with weight_scale 4 and id 1
[/gdscript]
[csharp]
var astar = new AStar2D();
astar.AddPoint(1, new Vector2(1, 0), 4); // Adds the point (1, 0) with weight_scale 4 and id 1
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AStar2D"
	"graphics.gd/variant/Vector2"
)

func AStar2D_AddPoint() {
	var astar = AStar2D.New()
	AStar2D.Expanded(astar).AddPoint(1, Vector2.New(1, 0), 4) // Adds the point (1, 0) with weight_scale 4 and id 1
}
