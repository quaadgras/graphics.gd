package main

import (
	"graphics.gd/classdb/AStarGrid2D"
	"graphics.gd/classdb/TileMapLayer"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
)

const BaseLineWidth = 3

var (
	TileStartPoint = Vector2i.XY{1, 0}
	TileEndPoint   = Vector2i.XY{2, 0}
	CellSize       = Vector2.XY{64, 64}
	DrawColor      = Color.Mul(Color.X11.White, Color.RGBA{1, 1, 1, 0.5})
)

type PathFindAStar struct {
	TileMapLayer.Extension[PathFindAStar]

	astar       AStarGrid2D.Instance
	start_point Vector2i.XY
	end_point   Vector2i.XY
	path        []Vector2.XY
}

func NewPathFindAStar() *PathFindAStar {
	return &PathFindAStar{
		astar: AStarGrid2D.New(),
	}
}

func (p *PathFindAStar) Ready() {
	// Region should match the size of the playable area plus one (in tiles).
	// In this demo, the playable area is 17×9 tiles, so the rect size is 18×10.
	// Depending on the setup TileMapLayer's GetUsedRect() can also be used.
	p.astar.SetRegion(Rect2i.New(0, 0, 18, 10))
	p.astar.SetCellSize(CellSize)
	p.astar.SetOffset(Vector2.MulX(CellSize, 0.5))
	p.astar.SetDefaultComputeHeuristic(AStarGrid2D.HeuristicManhattan)
	p.astar.SetDefaultEstimateHeuristic(AStarGrid2D.HeuristicManhattan)
	p.astar.SetDiagonalMode(AStarGrid2D.DiagonalModeNever)
	p.astar.Update()

	// Iterate over all cells on the tile map layer and mark them as
	// non-passable.
	for _, pos := range p.AsTileMapLayer().GetUsedCells() {
		p.astar.SetPointSolid(pos)
	}
}

func (p *PathFindAStar) Draw() {
	if len(p.path) == 0 {
		return
	}
	var last_point = p.path[0]
	for index := 1; index < len(p.path); index++ {
		var current_point = p.path[index]
		p.AsCanvasItem().MoreArgs().DrawLine(last_point, current_point, DrawColor, BaseLineWidth, false)
		p.AsCanvasItem().DrawCircle(current_point, BaseLineWidth*2, DrawColor)
		last_point = current_point
	}
}

func (p *PathFindAStar) RoundLocalPosition(local_position Vector2i.XY) Vector2i.XY {
	return Vector2i.From(p.AsTileMapLayer().MapToLocal(p.AsTileMapLayer().LocalToMap(Vector2.From(local_position))))
}

func (p *PathFindAStar) IsPointWalkable(point Vector2.XY) bool {
	var map_position = p.AsTileMapLayer().LocalToMap(point)
	if p.astar.IsInBoundsv(map_position) {
		return !p.astar.IsPointSolid(map_position)
	}
	return false
}

func (p *PathFindAStar) ClearPath() {
	if len(p.path) > 0 {
		p.path = nil
		p.AsTileMapLayer().EraseCell(p.start_point)
		p.AsTileMapLayer().EraseCell(p.end_point)
		p.AsCanvasItem().QueueRedraw()
	}
}

func (p *PathFindAStar) FindPath(start_point, end_point Vector2.XY) []Vector2.XY {
	p.ClearPath()
	p.start_point = p.AsTileMapLayer().LocalToMap(start_point)
	p.end_point = p.AsTileMapLayer().LocalToMap(end_point)
	p.path = p.astar.GetPointPath(p.start_point, p.end_point)
	if len(p.path) > 0 {
		p.AsTileMapLayer().MoreArgs().SetCell(p.start_point, 0, TileStartPoint, 0)
		p.AsTileMapLayer().MoreArgs().SetCell(p.end_point, 0, TileEndPoint, 0)
	}
	p.AsCanvasItem().QueueRedraw()
	return p.path
}
