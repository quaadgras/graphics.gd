/*
[gdscript]
var source_id = tile_map_layer.get_cell_source_id(Vector2i(x, y))
if source_id > -1:
    var scene_source = tile_map_layer.tile_set.get_source(source_id)
    if scene_source is TileSetScenesCollectionSource:
        var alt_id = tile_map_layer.get_cell_alternative_tile(Vector2i(x, y))
        # The assigned PackedScene.
        var scene = scene_source.get_scene_tile_scene(alt_id)
[/gdscript]
[csharp]
int sourceId = tileMapLayer.GetCellSourceId(new Vector2I(x, y));
if (sourceId > -1)
{
    TileSetSource source = tileMapLayer.TileSet.GetSource(sourceId);
    if (source is TileSetScenesCollectionSource sceneSource)
    {
        int altId = tileMapLayer.GetCellAlternativeTile(new Vector2I(x, y));
        // The assigned PackedScene.
        PackedScene scene = sceneSource.GetSceneTileScene(altId);
    }
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/TileMapLayer"
	"graphics.gd/classdb/TileSetScenesCollectionSource"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2i"
)

func ExampleTileSetScenesCollectionSource(tileMapLayer TileMapLayer.Instance, x, y int) {
	sourceId := tileMapLayer.GetCellSourceId(Vector2i.New(x, y))
	if sourceId > -1 {
		var scene_source = tileMapLayer.TileSet().GetSource(sourceId)
		if scene_source, ok := Object.As[TileSetScenesCollectionSource.Instance](scene_source); ok {
			altId := tileMapLayer.GetCellAlternativeTile(Vector2i.New(x, y))
			var scene = scene_source.GetSceneTileScene(altId)
			_ = scene
		}
	}
}
