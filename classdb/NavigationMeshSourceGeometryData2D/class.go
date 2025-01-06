// Package NavigationMeshSourceGeometryData2D provides methods for working with NavigationMeshSourceGeometryData2D object instances.
package NavigationMeshSourceGeometryData2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Array"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Container for parsed source geometry data used in navigation mesh baking.
*/
type Instance [1]gdclass.NavigationMeshSourceGeometryData2D
type Any interface {
	gd.IsClass
	AsNavigationMeshSourceGeometryData2D() Instance
}

/*
Clears the internal data.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Returns [code]true[/code] when parsed source geometry data exists.
*/
func (self Instance) HasData() bool {
	return bool(class(self).HasData())
}

/*
Appends another array of [param traversable_outlines] at the end of the existing traversable outlines array.
*/
func (self Instance) AppendTraversableOutlines(traversable_outlines gd.Array) {
	class(self).AppendTraversableOutlines(traversable_outlines)
}

/*
Appends another array of [param obstruction_outlines] at the end of the existing obstruction outlines array.
*/
func (self Instance) AppendObstructionOutlines(obstruction_outlines gd.Array) {
	class(self).AppendObstructionOutlines(obstruction_outlines)
}

/*
Adds the outline points of a shape as traversable area.
*/
func (self Instance) AddTraversableOutline(shape_outline []Vector2.XY) {
	class(self).AddTraversableOutline(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&shape_outline))))
}

/*
Adds the outline points of a shape as obstructed area.
*/
func (self Instance) AddObstructionOutline(shape_outline []Vector2.XY) {
	class(self).AddObstructionOutline(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&shape_outline))))
}

/*
Adds the geometry data of another [NavigationMeshSourceGeometryData2D] to the navigation mesh baking data.
*/
func (self Instance) Merge(other_geometry [1]gdclass.NavigationMeshSourceGeometryData2D) {
	class(self).Merge(other_geometry)
}

/*
Adds a projected obstruction shape to the source geometry. If [param carve] is [code]true[/code] the carved shape will not be affected by additional offsets (e.g. agent radius) of the navigation mesh baking process.
*/
func (self Instance) AddProjectedObstruction(vertices []Vector2.XY, carve bool) {
	class(self).AddProjectedObstruction(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&vertices))), carve)
}

/*
Clears all projected obstructions.
*/
func (self Instance) ClearProjectedObstructions() {
	class(self).ClearProjectedObstructions()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationMeshSourceGeometryData2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationMeshSourceGeometryData2D"))
	casted := Instance{*(*gdclass.NavigationMeshSourceGeometryData2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) TraversableOutlines() gd.Array {
	return gd.Array(class(self).GetTraversableOutlines())
}

func (self Instance) SetTraversableOutlines(value gd.Array) {
	class(self).SetTraversableOutlines(value)
}

func (self Instance) ObstructionOutlines() gd.Array {
	return gd.Array(class(self).GetObstructionOutlines())
}

func (self Instance) SetObstructionOutlines(value gd.Array) {
	class(self).SetObstructionOutlines(value)
}

func (self Instance) ProjectedObstructions() Array.Any {
	return Array.Any(class(self).GetProjectedObstructions())
}

func (self Instance) SetProjectedObstructions(value Array.Any) {
	class(self).SetProjectedObstructions(value)
}

/*
Clears the internal data.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] when parsed source geometry data exists.
*/
//go:nosplit
func (self class) HasData() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_has_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets all the traversable area outlines arrays.
*/
//go:nosplit
func (self class) SetTraversableOutlines(traversable_outlines gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(traversable_outlines))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_set_traversable_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns all the traversable area outlines arrays.
*/
//go:nosplit
func (self class) GetTraversableOutlines() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_get_traversable_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets all the obstructed area outlines arrays.
*/
//go:nosplit
func (self class) SetObstructionOutlines(obstruction_outlines gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obstruction_outlines))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_set_obstruction_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns all the obstructed area outlines arrays.
*/
//go:nosplit
func (self class) GetObstructionOutlines() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_get_obstruction_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Appends another array of [param traversable_outlines] at the end of the existing traversable outlines array.
*/
//go:nosplit
func (self class) AppendTraversableOutlines(traversable_outlines gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(traversable_outlines))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_append_traversable_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Appends another array of [param obstruction_outlines] at the end of the existing obstruction outlines array.
*/
//go:nosplit
func (self class) AppendObstructionOutlines(obstruction_outlines gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obstruction_outlines))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_append_obstruction_outlines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds the outline points of a shape as traversable area.
*/
//go:nosplit
func (self class) AddTraversableOutline(shape_outline gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape_outline))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_add_traversable_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds the outline points of a shape as obstructed area.
*/
//go:nosplit
func (self class) AddObstructionOutline(shape_outline gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape_outline))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_add_obstruction_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds the geometry data of another [NavigationMeshSourceGeometryData2D] to the navigation mesh baking data.
*/
//go:nosplit
func (self class) Merge(other_geometry [1]gdclass.NavigationMeshSourceGeometryData2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(other_geometry[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_merge, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a projected obstruction shape to the source geometry. If [param carve] is [code]true[/code] the carved shape will not be affected by additional offsets (e.g. agent radius) of the navigation mesh baking process.
*/
//go:nosplit
func (self class) AddProjectedObstruction(vertices gd.PackedVector2Array, carve bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertices))
	callframe.Arg(frame, carve)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_add_projected_obstruction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears all projected obstructions.
*/
//go:nosplit
func (self class) ClearProjectedObstructions() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_clear_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) SetProjectedObstructions(projected_obstructions gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(projected_obstructions))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_set_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the projected obstructions as an [Array] of dictionaries. Each [Dictionary] contains the following entries:
- [code]vertices[/code] - A [PackedFloat32Array] that defines the outline points of the projected shape.
- [code]carve[/code] - A [bool] that defines how the projected shape affects the navigation mesh baking. If [code]true[/code] the projected shape will not be affected by addition offsets, e.g. agent radius.
*/
//go:nosplit
func (self class) GetProjectedObstructions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMeshSourceGeometryData2D.Bind_get_projected_obstructions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsNavigationMeshSourceGeometryData2D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsNavigationMeshSourceGeometryData2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
	gdclass.Register("NavigationMeshSourceGeometryData2D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationMeshSourceGeometryData2D{*(*gdclass.NavigationMeshSourceGeometryData2D)(unsafe.Pointer(&ptr))}
	})
}
