package TileSetSource

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2i"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Exposes a set of tiles for a [TileSet] resource.
Tiles in a source are indexed with two IDs, coordinates ID (of type Vector2i) and an alternative ID (of type int), named according to their use in the [TileSetAtlasSource] class.
Depending on the TileSet source type, those IDs might have restrictions on their values, this is why the base [TileSetSource] class only exposes getters for them.
You can iterate over all tiles exposed by a TileSetSource by first iterating over coordinates IDs using [method get_tiles_count] and [method get_tile_id], then over alternative IDs using [method get_alternative_tiles_count] and [method get_alternative_tile_id].
[b]Warning:[/b] [TileSetSource] can only be added to one TileSet at the same time. Calling [method TileSet.add_source] on a second [TileSet] will remove the source from the first one.
*/
type Instance [1]classdb.TileSetSource
type Any interface {
	gd.IsClass
	AsTileSetSource() Instance
}

/*
Returns how many tiles this atlas source defines (not including alternative tiles).
*/
func (self Instance) GetTilesCount() int {
	return int(int(class(self).GetTilesCount()))
}

/*
Returns the tile coordinates ID of the tile with index [param index].
*/
func (self Instance) GetTileId(index int) Vector2i.XY {
	return Vector2i.XY(class(self).GetTileId(gd.Int(index)))
}

/*
Returns if this atlas has a tile with coordinates ID [param atlas_coords].
*/
func (self Instance) HasTile(atlas_coords Vector2i.XY) bool {
	return bool(class(self).HasTile(gd.Vector2i(atlas_coords)))
}

/*
Returns the number of alternatives tiles for the coordinates ID [param atlas_coords].
For [TileSetAtlasSource], this always return at least 1, as the base tile with ID 0 is always part of the alternatives list.
Returns -1 if there is not tile at the given coords.
*/
func (self Instance) GetAlternativeTilesCount(atlas_coords Vector2i.XY) int {
	return int(int(class(self).GetAlternativeTilesCount(gd.Vector2i(atlas_coords))))
}

/*
Returns the alternative ID for the tile with coordinates ID [param atlas_coords] at index [param index].
*/
func (self Instance) GetAlternativeTileId(atlas_coords Vector2i.XY, index int) int {
	return int(int(class(self).GetAlternativeTileId(gd.Vector2i(atlas_coords), gd.Int(index))))
}

/*
Returns if the base tile at coordinates [param atlas_coords] has an alternative with ID [param alternative_tile].
*/
func (self Instance) HasAlternativeTile(atlas_coords Vector2i.XY, alternative_tile int) bool {
	return bool(class(self).HasAlternativeTile(gd.Vector2i(atlas_coords), gd.Int(alternative_tile)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TileSetSource

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileSetSource"))
	return Instance{classdb.TileSetSource(object)}
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
func (self class) AsTileSetSource() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTileSetSource() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("TileSetSource", func(ptr gd.Object) any { return classdb.TileSetSource(ptr) })
}
