package RibbonTrailMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PrimitiveMesh"
import "grow.graphics/gd/gdclass/Mesh"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[RibbonTrailMesh] represents a straight ribbon-shaped mesh with variable width. The ribbon is composed of a number of flat or cross-shaped sections, each with the same [member section_length] and number of [member section_segments]. A [member curve] is sampled along the total length of the ribbon, meaning that the curve determines the size of the ribbon along its length.
This primitive mesh is usually used for particle trails.

*/
type Go [1]classdb.RibbonTrailMesh
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RibbonTrailMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RibbonTrailMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Shape() classdb.RibbonTrailMeshShape {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.RibbonTrailMeshShape(class(self).GetShape())
}

func (self Go) SetShape(value classdb.RibbonTrailMeshShape) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShape(value)
}

func (self Go) Size() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSize()))
}

func (self Go) SetSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSize(gd.Float(value))
}

func (self Go) Sections() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSections()))
}

func (self Go) SetSections(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSections(gd.Int(value))
}

func (self Go) SectionLength() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSectionLength()))
}

func (self Go) SetSectionLength(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSectionLength(gd.Float(value))
}

func (self Go) SectionSegments() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSectionSegments()))
}

func (self Go) SetSectionSegments(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSectionSegments(gd.Int(value))
}

func (self Go) Curve() gdclass.Curve {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Curve(class(self).GetCurve(gc))
}

func (self Go) SetCurve(value gdclass.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCurve(value)
}

//go:nosplit
func (self class) SetSize(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSections(sections gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sections)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_set_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSections() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_get_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSectionLength(section_length gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, section_length)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_set_section_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSectionLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_get_section_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSectionSegments(section_segments gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, section_segments)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_set_section_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSectionSegments() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_get_section_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCurve(curve gdclass.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurve(ctx gd.Lifetime) gdclass.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShape(shape classdb.RibbonTrailMeshShape)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShape() classdb.RibbonTrailMeshShape {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RibbonTrailMeshShape](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RibbonTrailMesh.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRibbonTrailMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRibbonTrailMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPrimitiveMesh() PrimitiveMesh.GD { return *((*PrimitiveMesh.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPrimitiveMesh() PrimitiveMesh.Go { return *((*PrimitiveMesh.Go)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.GD { return *((*Mesh.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMesh() Mesh.Go { return *((*Mesh.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}
func init() {classdb.Register("RibbonTrailMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Shape = classdb.RibbonTrailMeshShape

const (
/*Gives the mesh a single flat face.*/
	ShapeFlat Shape = 0
/*Gives the mesh two perpendicular flat faces, making a cross shape.*/
	ShapeCross Shape = 1
)
