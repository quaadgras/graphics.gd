package TileSetSource

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Exposes a set of tiles for a [TileSet] resource.
Tiles in a source are indexed with two IDs, coordinates ID (of type Vector2i) and an alternative ID (of type int), named according to their use in the [TileSetAtlasSource] class.
Depending on the TileSet source type, those IDs might have restrictions on their values, this is why the base [TileSetSource] class only exposes getters for them.
You can iterate over all tiles exposed by a TileSetSource by first iterating over coordinates IDs using [method get_tiles_count] and [method get_tile_id], then over alternative IDs using [method get_alternative_tiles_count] and [method get_alternative_tile_id].
[b]Warning:[/b] [TileSetSource] can only be added to one TileSet at the same time. Calling [method TileSet.add_source] on a second [TileSet] will remove the source from the first one.

*/
type Go [1]classdb.TileSetSource

/*
Returns how many tiles this atlas source defines (not including alternative tiles).
*/
func (self Go) GetTilesCount() int {
	return int(int(class(self).GetTilesCount()))
}

/*
Returns the tile coordinates ID of the tile with index [param index].
*/
func (self Go) GetTileId(index int) gd.Vector2i {
	return gd.Vector2i(class(self).GetTileId(gd.Int(index)))
}

/*
Returns if this atlas has a tile with coordinates ID [param atlas_coords].
*/
func (self Go) HasTile(atlas_coords gd.Vector2i) bool {
	return bool(class(self).HasTile(atlas_coords))
}

/*
Returns the number of alternatives tiles for the coordinates ID [param atlas_coords].
For [TileSetAtlasSource], this always return at least 1, as the base tile with ID 0 is always part of the alternatives list.
Returns -1 if there is not tile at the given coords.
*/
func (self Go) GetAlternativeTilesCount(atlas_coords gd.Vector2i) int {
	return int(int(class(self).GetAlternativeTilesCount(atlas_coords)))
}

/*
Returns the alternative ID for the tile with coordinates ID [param atlas_coords] at index [param index].
*/
func (self Go) GetAlternativeTileId(atlas_coords gd.Vector2i, index int) int {
	return int(int(class(self).GetAlternativeTileId(atlas_coords, gd.Int(index))))
}

/*
Returns if the base tile at coordinates [param atlas_coords] has an alternative with ID [param alternative_tile].
*/
func (self Go) HasAlternativeTile(atlas_coords gd.Vector2i, alternative_tile int) bool {
	return bool(class(self).HasAlternativeTile(atlas_coords, gd.Int(alternative_tile)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TileSetSource
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileSetSource"))
	return Go{classdb.TileSetSource(object)}
}

/*
Returns how many tiles this atlas source defines (not including alternative tiles).
*/
//go:nosplit
func (self class) GetTilesCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetSource.Bind_get_tiles_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tile coordinates ID of the tile with index [param index].
*/
//go:nosplit
func (self class) GetTileId(index gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetSource.Bind_get_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if this atlas has a tile with coordinates ID [param atlas_coords].
*/
//go:nosplit
func (self class) HasTile(atlas_coords gd.Vector2i) bool {
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetSource.Bind_has_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetSource.Bind_get_alternative_tiles_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the alternative ID for the tile with coordinates ID [param atlas_coords] at index [param index].
*/
//go:nosplit
func (self class) GetAlternativeTileId(atlas_coords gd.Vector2i, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetSource.Bind_get_alternative_tile_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if the base tile at coordinates [param atlas_coords] has an alternative with ID [param alternative_tile].
*/
//go:nosplit
func (self class) HasAlternativeTile(atlas_coords gd.Vector2i, alternative_tile gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetSource.Bind_has_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTileSetSource() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTileSetSource() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("TileSetSource", func(ptr gd.Object) any { return classdb.TileSetSource(ptr) })}
