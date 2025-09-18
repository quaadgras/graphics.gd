/*
func get_clicked_tile_power():
	var clicked_cell = tile_map_layer.local_to_map(tile_map_layer.get_local_mouse_position())
	var data = tile_map_layer.get_cell_tile_data(clicked_cell)
	if data:
		return data.get_custom_data("power")
	else:
		return 0
*/

package main

import (
	"graphics.gd/classdb/TileData"
	"graphics.gd/classdb/TileMapLayer"
)

var tileMapLayer TileMapLayer.Instance

func TileMapLayer_GetCellTileData() {
	GetClickedTilePower := func() int {
		var clickedCell = tileMapLayer.LocalToMap(tileMapLayer.AsCanvasItem().GetLocalMousePosition())
		var data = tileMapLayer.GetCellTileData(clickedCell)
		if data != TileData.Nil {
			return data.GetCustomData("power").(int)
		} else {
			return 0
		}
	}
	_ = GetClickedTilePower
}
