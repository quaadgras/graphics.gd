package TubeTrailMesh

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
[TubeTrailMesh] represents a straight tube-shaped mesh with variable width. The tube is composed of a number of cylindrical sections, each with the same [member section_length] and number of [member section_rings]. A [member curve] is sampled along the total length of the tube, meaning that the curve determines the radius of the tube along its length.
This primitive mesh is usually used for particle trails.

*/
type Go [1]classdb.TubeTrailMesh
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TubeTrailMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TubeTrailMesh"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Radius() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetRadius()))
}

func (self Go) SetRadius(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRadius(gd.Float(value))
}

func (self Go) RadialSteps() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetRadialSteps()))
}

func (self Go) SetRadialSteps(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRadialSteps(gd.Int(value))
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

func (self Go) SectionRings() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetSectionRings()))
}

func (self Go) SetSectionRings(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSectionRings(gd.Int(value))
}

func (self Go) CapTop() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCapTop())
}

func (self Go) SetCapTop(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCapTop(value)
}

func (self Go) CapBottom() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCapBottom())
}

func (self Go) SetCapBottom(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCapBottom(value)
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
func (self class) SetRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRadialSteps(radial_steps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radial_steps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_set_radial_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadialSteps() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_get_radial_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_set_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSections() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_get_sections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_set_section_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSectionLength() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_get_section_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSectionRings(section_rings gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, section_rings)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_set_section_rings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSectionRings() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_get_section_rings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCapTop(cap_top bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cap_top)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_set_cap_top, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCapTop() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_is_cap_top, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCapBottom(cap_bottom bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cap_bottom)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_set_cap_bottom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCapBottom() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_is_cap_bottom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurve(ctx gd.Lifetime) gdclass.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TubeTrailMesh.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsTubeTrailMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTubeTrailMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("TubeTrailMesh", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
