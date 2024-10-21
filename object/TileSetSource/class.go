package TileSetSource

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Exposes a set of tiles for a [TileSet] resource.
Tiles in a source are indexed with two IDs, coordinates ID (of type Vector2i) and an alternative ID (of type int), named according to their use in the [TileSetAtlasSource] class.
Depending on the TileSet source type, those IDs might have restrictions on their values, this is why the base [TileSetSource] class only exposes getters for them.
You can iterate over all tiles exposed by a TileSetSource by first iterating over coordinates IDs using [method get_tiles_count] and [method get_tile_id], then over alternative IDs using [method get_alternative_tiles_count] and [method get_alternative_tile_id].
[b]Warning:[/b] [TileSetSource] can only be added to one TileSet at the same time. Calling [method TileSet.add_source] on a second [TileSet] will remove the source from the first one.

*/
type Simple [1]classdb.TileSetSource
func (self Simple) GetTilesCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTilesCount()))
}
func (self Simple) GetTileId(index int) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetTileId(gd.Int(index)))
}
func (self Simple) HasTile(atlas_coords gd.Vector2i) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasTile(atlas_coords))
}
func (self Simple) GetAlternativeTilesCount(atlas_coords gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAlternativeTilesCount(atlas_coords)))
}
func (self Simple) GetAlternativeTileId(atlas_coords gd.Vector2i, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetAlternativeTileId(atlas_coords, gd.Int(index))))
}
func (self Simple) HasAlternativeTile(atlas_coords gd.Vector2i, alternative_tile int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAlternativeTile(atlas_coords, gd.Int(alternative_tile)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TileSetSource
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns how many tiles this atlas source defines (not including alternative tiles).
*/
//go:nosplit
func (self class) GetTilesCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetSource.Bind_get_tiles_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tile coordinates ID of the tile with index [param index].
*/
//go:nosplit
func (self class) GetTileId(index gd.Int) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetSource.Bind_get_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if this atlas has a tile with coordinates ID [param atlas_coords].
*/
//go:nosplit
func (self class) HasTile(atlas_coords gd.Vector2i) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetSource.Bind_has_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of alternatives tiles for the coordinates ID [param atlas_coords].
For [TileSetAtlasSource], this always return at least 1, as the base tile with ID 0 is always part of the alternatives list.
Returns -1 if there is not tile at the given coords.
*/
//go:nosplit
func (self class) GetAlternativeTilesCount(atlas_coords gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetSource.Bind_get_alternative_tiles_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the alternative ID for the tile with coordinates ID [param atlas_coords] at index [param index].
*/
//go:nosplit
func (self class) GetAlternativeTileId(atlas_coords gd.Vector2i, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetSource.Bind_get_alternative_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if the base tile at coordinates [param atlas_coords] has an alternative with ID [param alternative_tile].
*/
//go:nosplit
func (self class) HasAlternativeTile(atlas_coords gd.Vector2i, alternative_tile gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileSetSource.Bind_has_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTileSetSource() Expert { return self[0].AsTileSetSource() }


//go:nosplit
func (self Simple) AsTileSetSource() Simple { return self[0].AsTileSetSource() }


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
func init() {classdb.Register("TileSetSource", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
