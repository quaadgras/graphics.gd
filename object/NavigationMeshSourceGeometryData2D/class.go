package NavigationMeshSourceGeometryData2D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
Container for parsed source geometry data used in navigation mesh baking.

*/
type Simple [1]classdb.NavigationMeshSourceGeometryData2D
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) HasData() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasData())
}
func (self Simple) SetTraversableOutlines(traversable_outlines gd.ArrayOf[gd.PackedVector2Array]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTraversableOutlines(traversable_outlines)
}
func (self Simple) GetTraversableOutlines() gd.ArrayOf[gd.PackedVector2Array] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.PackedVector2Array](Expert(self).GetTraversableOutlines(gc))
}
func (self Simple) SetObstructionOutlines(obstruction_outlines gd.ArrayOf[gd.PackedVector2Array]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetObstructionOutlines(obstruction_outlines)
}
func (self Simple) GetObstructionOutlines() gd.ArrayOf[gd.PackedVector2Array] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.PackedVector2Array](Expert(self).GetObstructionOutlines(gc))
}
func (self Simple) AppendTraversableOutlines(traversable_outlines gd.ArrayOf[gd.PackedVector2Array]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AppendTraversableOutlines(traversable_outlines)
}
func (self Simple) AppendObstructionOutlines(obstruction_outlines gd.ArrayOf[gd.PackedVector2Array]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AppendObstructionOutlines(obstruction_outlines)
}
func (self Simple) AddTraversableOutline(shape_outline gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddTraversableOutline(shape_outline)
}
func (self Simple) AddObstructionOutline(shape_outline gd.PackedVector2Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddObstructionOutline(shape_outline)
}
func (self Simple) Merge(other_geometry [1]classdb.NavigationMeshSourceGeometryData2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Merge(other_geometry)
}
func (self Simple) AddProjectedObstruction(vertices gd.PackedVector2Array, carve bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddProjectedObstruction(vertices, carve)
}
func (self Simple) ClearProjectedObstructions() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearProjectedObstructions()
}
func (self Simple) SetProjectedObstructions(projected_obstructions gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProjectedObstructions(projected_obstructions)
}
func (self Simple) GetProjectedObstructions() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetProjectedObstructions(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationMeshSourceGeometryData2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Clears the internal data.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] when parsed source geometry data exists.
*/
//go:nosplit
func (self class) HasData() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_has_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets all the traversable area outlines arrays.
*/
//go:nosplit
func (self class) SetTraversableOutlines(traversable_outlines gd.ArrayOf[gd.PackedVector2Array])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(traversable_outlines))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_set_traversable_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns all the traversable area outlines arrays.
*/
//go:nosplit
func (self class) GetTraversableOutlines(ctx gd.Lifetime) gd.ArrayOf[gd.PackedVector2Array] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_get_traversable_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.PackedVector2Array](ret)
}
/*
Sets all the obstructed area outlines arrays.
*/
//go:nosplit
func (self class) SetObstructionOutlines(obstruction_outlines gd.ArrayOf[gd.PackedVector2Array])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(obstruction_outlines))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_set_obstruction_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns all the obstructed area outlines arrays.
*/
//go:nosplit
func (self class) GetObstructionOutlines(ctx gd.Lifetime) gd.ArrayOf[gd.PackedVector2Array] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_get_obstruction_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.PackedVector2Array](ret)
}
/*
Appends another array of [param traversable_outlines] at the end of the existing traversable outlines array.
*/
//go:nosplit
func (self class) AppendTraversableOutlines(traversable_outlines gd.ArrayOf[gd.PackedVector2Array])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(traversable_outlines))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_append_traversable_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Appends another array of [param obstruction_outlines] at the end of the existing obstruction outlines array.
*/
//go:nosplit
func (self class) AppendObstructionOutlines(obstruction_outlines gd.ArrayOf[gd.PackedVector2Array])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(obstruction_outlines))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_append_obstruction_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds the outline points of a shape as traversable area.
*/
//go:nosplit
func (self class) AddTraversableOutline(shape_outline gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape_outline))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_add_traversable_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds the outline points of a shape as obstructed area.
*/
//go:nosplit
func (self class) AddObstructionOutline(shape_outline gd.PackedVector2Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shape_outline))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_add_obstruction_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds the geometry data of another [NavigationMeshSourceGeometryData2D] to the navigation mesh baking data.
*/
//go:nosplit
func (self class) Merge(other_geometry object.NavigationMeshSourceGeometryData2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(other_geometry[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_merge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a projected obstruction shape to the source geometry. If [param carve] is [code]true[/code] the carved shape will not be affected by additional offsets (e.g. agent radius) of the navigation mesh baking process.
*/
//go:nosplit
func (self class) AddProjectedObstruction(vertices gd.PackedVector2Array, carve bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	callframe.Arg(frame, carve)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_add_projected_obstruction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all projected obstructions.
*/
//go:nosplit
func (self class) ClearProjectedObstructions()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_clear_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the projected obstructions with an Array of Dictionaries with the following key value pairs:
[codeblocks]
[gdscript]
"vertices" : PackedFloat32Array
"carve" : bool
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) SetProjectedObstructions(projected_obstructions gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(projected_obstructions))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_set_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the projected obstructions as an [Array] of dictionaries. Each [Dictionary] contains the following entries:
- [code]vertices[/code] - A [PackedFloat32Array] that defines the outline points of the projected shape.
- [code]carve[/code] - A [bool] that defines how the projected shape affects the navigation mesh baking. If [code]true[/code] the projected shape will not be affected by addition offsets, e.g. agent radius.
*/
//go:nosplit
func (self class) GetProjectedObstructions(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationMeshSourceGeometryData2D.Bind_get_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsNavigationMeshSourceGeometryData2D() Expert { return self[0].AsNavigationMeshSourceGeometryData2D() }


//go:nosplit
func (self Simple) AsNavigationMeshSourceGeometryData2D() Simple { return self[0].AsNavigationMeshSourceGeometryData2D() }


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
func init() {classdb.Register("NavigationMeshSourceGeometryData2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
