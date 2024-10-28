package TileMapPattern

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
This resource holds a set of cells to help bulk manipulations of [TileMap].
A pattern always start at the [code](0,0)[/code] coordinates and cannot have cells with negative coordinates.

*/
type Go [1]classdb.TileMapPattern

/*
Sets the tile identifiers for the cell at coordinates [param coords]. See [method TileMap.set_cell].
*/
func (self Go) SetCell(coords gd.Vector2i) {
	class(self).SetCell(coords, gd.Int(-1), gd.Vector2i{-1, -1}, gd.Int(-1))
}

/*
Returns whether the pattern has a tile at the given coordinates.
*/
func (self Go) HasCell(coords gd.Vector2i) bool {
	return bool(class(self).HasCell(coords))
}

/*
Remove the cell at the given coordinates.
*/
func (self Go) RemoveCell(coords gd.Vector2i, update_size bool) {
	class(self).RemoveCell(coords, update_size)
}

/*
Returns the tile source ID of the cell at [param coords].
*/
func (self Go) GetCellSourceId(coords gd.Vector2i) int {
	return int(int(class(self).GetCellSourceId(coords)))
}

/*
Returns the tile atlas coordinates ID of the cell at [param coords].
*/
func (self Go) GetCellAtlasCoords(coords gd.Vector2i) gd.Vector2i {
	return gd.Vector2i(class(self).GetCellAtlasCoords(coords))
}

/*
Returns the tile alternative ID of the cell at [param coords].
*/
func (self Go) GetCellAlternativeTile(coords gd.Vector2i) int {
	return int(int(class(self).GetCellAlternativeTile(coords)))
}

/*
Returns the list of used cell coordinates in the pattern.
*/
func (self Go) GetUsedCells() gd.Array {
	return gd.Array(class(self).GetUsedCells())
}

/*
Returns the size, in cells, of the pattern.
*/
func (self Go) GetSize() gd.Vector2i {
	return gd.Vector2i(class(self).GetSize())
}

/*
Sets the size of the pattern.
*/
func (self Go) SetSize(size gd.Vector2i) {
	class(self).SetSize(size)
}

/*
Returns whether the pattern is empty or not.
*/
func (self Go) IsEmpty() bool {
	return bool(class(self).IsEmpty())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TileMapPattern
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileMapPattern"))
	return Go{classdb.TileMapPattern(object)}
}

/*
Sets the tile identifiers for the cell at coordinates [param coords]. See [method TileMap.set_cell].
*/
//go:nosplit
func (self class) SetCell(coords gd.Vector2i, source_id gd.Int, atlas_coords gd.Vector2i, alternative_tile gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_set_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the pattern has a tile at the given coordinates.
*/
//go:nosplit
func (self class) HasCell(coords gd.Vector2i) bool {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_has_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Remove the cell at the given coordinates.
*/
//go:nosplit
func (self class) RemoveCell(coords gd.Vector2i, update_size bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, update_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_remove_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tile source ID of the cell at [param coords].
*/
//go:nosplit
func (self class) GetCellSourceId(coords gd.Vector2i) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_cell_source_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tile atlas coordinates ID of the cell at [param coords].
*/
//go:nosplit
func (self class) GetCellAtlasCoords(coords gd.Vector2i) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_cell_atlas_coords, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tile alternative ID of the cell at [param coords].
*/
//go:nosplit
func (self class) GetCellAlternativeTile(coords gd.Vector2i) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_cell_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the list of used cell coordinates in the pattern.
*/
//go:nosplit
func (self class) GetUsedCells() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the size, in cells, of the pattern.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the size of the pattern.
*/
//go:nosplit
func (self class) SetSize(size gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the pattern is empty or not.
*/
//go:nosplit
func (self class) IsEmpty() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_is_empty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTileMapPattern() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTileMapPattern() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("TileMapPattern", func(ptr gd.Object) any { return classdb.TileMapPattern(ptr) })}
