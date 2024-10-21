package TileSetScenesCollectionSource

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/TileSetSource"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
When placed on a [TileMap], tiles from [TileSetScenesCollectionSource] will automatically instantiate an associated scene at the cell's position in the TileMap.
Scenes are instantiated as children of the [TileMap] when it enters the tree. If you add/remove a scene tile in the [TileMap] that is already inside the tree, the [TileMap] will automatically instantiate/free the scene accordingly.
[b]Note:[/b] Scene tiles all occupy one tile slot and instead use alternate tile ID to identify scene index. [method TileSetSource.get_tiles_count] will always return [code]1[/code]. Use [method get_scene_tiles_count] to get a number of scenes in a [TileSetScenesCollectionSource].
Use this code if you want to find the scene path at a given tile in [TileMapLayer]:
[codeblocks]
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
[/codeblocks]

*/
type Simple [1]classdb.TileSetScenesCollectionSource
func (self Simple) GetSceneTilesCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSceneTilesCount()))
}
func (self Simple) GetSceneTileId(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSceneTileId(gd.Int(index))))
}
func (self Simple) HasSceneTileId(id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasSceneTileId(gd.Int(id)))
}
func (self Simple) CreateSceneTile(packed_scene [1]classdb.PackedScene, id_override int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).CreateSceneTile(packed_scene, gd.Int(id_override))))
}
func (self Simple) SetSceneTileId(id int, new_id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSceneTileId(gd.Int(id), gd.Int(new_id))
}
func (self Simple) SetSceneTileScene(id int, packed_scene [1]classdb.PackedScene) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSceneTileScene(gd.Int(id), packed_scene)
}
func (self Simple) GetSceneTileScene(id int) [1]classdb.PackedScene {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PackedScene(Expert(self).GetSceneTileScene(gc, gd.Int(id)))
}
func (self Simple) SetSceneTileDisplayPlaceholder(id int, display_placeholder bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSceneTileDisplayPlaceholder(gd.Int(id), display_placeholder)
}
func (self Simple) GetSceneTileDisplayPlaceholder(id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSceneTileDisplayPlaceholder(gd.Int(id)))
}
func (self Simple) RemoveSceneTile(id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveSceneTile(gd.Int(id))
}
func (self Simple) GetNextSceneTileId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetNextSceneTileId()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TileSetScenesCollectionSource
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the number or scene tiles this TileSet source has.
*/
//go:nosplit
func (self class) GetSceneTilesCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_get_scene_tiles_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the scene tile ID of the scene tile at [param index].
*/
//go:nosplit
func (self class) GetSceneTileId(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_get_scene_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether this TileSet source has a scene tile with [param id].
*/
//go:nosplit
func (self class) HasSceneTileId(id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_has_scene_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a scene-based tile out of the given scene.
Returns a newly generated unique ID.
*/
//go:nosplit
func (self class) CreateSceneTile(packed_scene object.PackedScene, id_override gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(packed_scene[0].AsPointer())[0])
	callframe.Arg(frame, id_override)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_create_scene_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Changes a scene tile's ID from [param id] to [param new_id]. This will fail if there is already a tile with an ID equal to [param new_id].
*/
//go:nosplit
func (self class) SetSceneTileId(id gd.Int, new_id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, new_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_set_scene_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Assigns a [PackedScene] resource to the scene tile with [param id]. This will fail if the scene does not extend CanvasItem, as positioning properties are needed to place the scene on the TileMap.
*/
//go:nosplit
func (self class) SetSceneTileScene(id gd.Int, packed_scene object.PackedScene)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(packed_scene[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_set_scene_tile_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [PackedScene] resource of scene tile with [param id].
*/
//go:nosplit
func (self class) GetSceneTileScene(ctx gd.Lifetime, id gd.Int) object.PackedScene {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_get_scene_tile_scene, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PackedScene
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets whether or not the scene tile with [param id] should display a placeholder in the editor. This might be useful for scenes that are not visible.
*/
//go:nosplit
func (self class) SetSceneTileDisplayPlaceholder(id gd.Int, display_placeholder bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, display_placeholder)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_set_scene_tile_display_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the scene tile with [param id] displays a placeholder in the editor.
*/
//go:nosplit
func (self class) GetSceneTileDisplayPlaceholder(id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_get_scene_tile_display_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Remove the scene tile with [param id].
*/
//go:nosplit
func (self class) RemoveSceneTile(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_remove_scene_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the scene ID a following call to [method create_scene_tile] would return.
*/
//go:nosplit
func (self class) GetNextSceneTileId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetScenesCollectionSource.Bind_get_next_scene_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTileSetScenesCollectionSource() Expert { return self[0].AsTileSetScenesCollectionSource() }


//go:nosplit
func (self Simple) AsTileSetScenesCollectionSource() Simple { return self[0].AsTileSetScenesCollectionSource() }


//go:nosplit
func (self class) AsTileSetSource() TileSetSource.Expert { return self[0].AsTileSetSource() }


//go:nosplit
func (self Simple) AsTileSetSource() TileSetSource.Simple { return self[0].AsTileSetSource() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TileSetScenesCollectionSource", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
