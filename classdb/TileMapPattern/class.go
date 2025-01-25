// Package TileMapPattern provides methods for working with TileMapPattern object instances.
package TileMapPattern

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This resource holds a set of cells to help bulk manipulations of [TileMap].
A pattern always start at the [code](0,0)[/code] coordinates and cannot have cells with negative coordinates.
*/
type Instance [1]gdclass.TileMapPattern

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTileMapPattern() Instance
}

/*
Sets the tile identifiers for the cell at coordinates [param coords]. See [method TileMap.set_cell].
*/
func (self Instance) SetCell(coords Vector2i.XY) {
	class(self).SetCell(gd.Vector2i(coords), gd.Int(-1), gd.Vector2i(gd.Vector2i{-1, -1}), gd.Int(-1))
}

/*
Returns whether the pattern has a tile at the given coordinates.
*/
func (self Instance) HasCell(coords Vector2i.XY) bool {
	return bool(class(self).HasCell(gd.Vector2i(coords)))
}

/*
Remove the cell at the given coordinates.
*/
func (self Instance) RemoveCell(coords Vector2i.XY, update_size bool) {
	class(self).RemoveCell(gd.Vector2i(coords), update_size)
}

/*
Returns the tile source ID of the cell at [param coords].
*/
func (self Instance) GetCellSourceId(coords Vector2i.XY) int {
	return int(int(class(self).GetCellSourceId(gd.Vector2i(coords))))
}

/*
Returns the tile atlas coordinates ID of the cell at [param coords].
*/
func (self Instance) GetCellAtlasCoords(coords Vector2i.XY) Vector2i.XY {
	return Vector2i.XY(class(self).GetCellAtlasCoords(gd.Vector2i(coords)))
}

/*
Returns the tile alternative ID of the cell at [param coords].
*/
func (self Instance) GetCellAlternativeTile(coords Vector2i.XY) int {
	return int(int(class(self).GetCellAlternativeTile(gd.Vector2i(coords))))
}

/*
Returns the list of used cell coordinates in the pattern.
*/
func (self Instance) GetUsedCells() []Vector2i.XY {
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](class(self).GetUsedCells()))
}

/*
Returns the size, in cells, of the pattern.
*/
func (self Instance) GetSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetSize())
}

/*
Sets the size of the pattern.
*/
func (self Instance) SetSize(size Vector2i.XY) {
	class(self).SetSize(gd.Vector2i(size))
}

/*
Returns whether the pattern is empty or not.
*/
func (self Instance) IsEmpty() bool {
	return bool(class(self).IsEmpty())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TileMapPattern

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileMapPattern"))
	casted := Instance{*(*gdclass.TileMapPattern)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Sets the tile identifiers for the cell at coordinates [param coords]. See [method TileMap.set_cell].
*/
//go:nosplit
func (self class) SetCell(coords gd.Vector2i, source_id gd.Int, atlas_coords gd.Vector2i, alternative_tile gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_set_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_has_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Remove the cell at the given coordinates.
*/
//go:nosplit
func (self class) RemoveCell(coords gd.Vector2i, update_size bool) {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, update_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_remove_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_cell_source_id, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_cell_atlas_coords, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_cell_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the size of the pattern.
*/
//go:nosplit
func (self class) SetSize(size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the pattern is empty or not.
*/
//go:nosplit
func (self class) IsEmpty() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMapPattern.Bind_is_empty, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTileMapPattern() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTileMapPattern() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("TileMapPattern", func(ptr gd.Object) any {
		return [1]gdclass.TileMapPattern{*(*gdclass.TileMapPattern)(unsafe.Pointer(&ptr))}
	})
}
