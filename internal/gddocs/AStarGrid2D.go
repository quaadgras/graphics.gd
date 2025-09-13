/*
[gdscript]
var astar_grid = AStarGrid2D.new()
astar_grid.region = Rect2i(0, 0, 32, 32)
astar_grid.cell_size = Vector2(16, 16)
astar_grid.update()
print(astar_grid.get_id_path(Vector2i(0, 0), Vector2i(3, 4))) # Prints [(0, 0), (1, 1), (2, 2), (3, 3), (3, 4)]
print(astar_grid.get_point_path(Vector2i(0, 0), Vector2i(3, 4))) # Prints [(0, 0), (16, 16), (32, 32), (48, 48), (48, 64)]
[/gdscript]
[csharp]
AStarGrid2D astarGrid = new AStarGrid2D();
astarGrid.Region = new Rect2I(0, 0, 32, 32);
astarGrid.CellSize = new Vector2I(16, 16);
astarGrid.Update();
GD.Print(astarGrid.GetIdPath(Vector2I.Zero, new Vector2I(3, 4))); // Prints [(0, 0), (1, 1), (2, 2), (3, 3), (3, 4)]
GD.Print(astarGrid.GetPointPath(Vector2I.Zero, new Vector2I(3, 4))); // Prints [(0, 0), (16, 16), (32, 32), (48, 48), (48, 64)]
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/AStarGrid2D"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

func SetupGrid() {
	var astar_grid = AStarGrid2D.New()
	astar_grid.SetRegion(Rect2i.New(0, 0, 32, 32))
	astar_grid.SetCellSize(Vector2.New(16, 16))
	astar_grid.Update()
	fmt.Print(astar_grid.GetIdPath(Vector2i.New(0, 0), Vector2i.New(3, 4)))    // Prints [(0, 0), (1, 1), (2, 2), (3, 3), (3, 4)]
	fmt.Print(astar_grid.GetPointPath(Vector2i.New(0, 0), Vector2i.New(3, 4))) // Prints [(0, 0), (16, 16), (32, 32), (48, 48), (48, 64)]
}
