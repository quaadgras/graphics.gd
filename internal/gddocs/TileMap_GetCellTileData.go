/*
func get_clicked_tile_power():
	var clicked_cell = tile_map.local_to_map(tile_map.get_local_mouse_position())
	var data = tile_map.get_cell_tile_data(0, clicked_cell)
	if data:
		return data.get_custom_data("power")
	else:
		return 0
*/

package main

import (
	"graphics.gd/classdb/TileData"
	"graphics.gd/classdb/TileMap"
)

var tileMap TileMap.Instance

func TileMap_GetCellTileData() {
	GetClickedTilePower := func() int {
		var clickedCell = tileMap.LocalToMap(tileMap.AsCanvasItem().GetLocalMousePosition())
		var data = tileMap.GetCellTileData(0, clickedCell)
		if data != TileData.Nil {
			return data.GetCustomData("power").(int)
		} else {
			return 0
		}
	}
	_ = GetClickedTilePower
}
