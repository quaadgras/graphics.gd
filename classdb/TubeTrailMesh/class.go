// Package TubeTrailMesh provides methods for working with TubeTrailMesh object instances.
package TubeTrailMesh

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/PrimitiveMesh"
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
[TubeTrailMesh] represents a straight tube-shaped mesh with variable width. The tube is composed of a number of cylindrical sections, each with the same [member section_length] and number of [member section_rings]. A [member curve] is sampled along the total length of the tube, meaning that the curve determines the radius of the tube along its length.
This primitive mesh is usually used for particle trails.
*/
type Instance [1]gdclass.TubeTrailMesh

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTubeTrailMesh() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TubeTrailMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TubeTrailMesh"))
	casted := Instance{*(*gdclass.TubeTrailMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Radius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetRadius(value Float.X) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) RadialSteps() int {
	return int(int(class(self).GetRadialSteps()))
}

func (self Instance) SetRadialSteps(value int) {
	class(self).SetRadialSteps(gd.Int(value))
}

func (self Instance) Sections() int {
	return int(int(class(self).GetSections()))
}

func (self Instance) SetSections(value int) {
	class(self).SetSections(gd.Int(value))
}

func (self Instance) SectionLength() Float.X {
	return Float.X(Float.X(class(self).GetSectionLength()))
}

func (self Instance) SetSectionLength(value Float.X) {
	class(self).SetSectionLength(gd.Float(value))
}

func (self Instance) SectionRings() int {
	return int(int(class(self).GetSectionRings()))
}

func (self Instance) SetSectionRings(value int) {
	class(self).SetSectionRings(gd.Int(value))
}

func (self Instance) CapTop() bool {
	return bool(class(self).IsCapTop())
}

func (self Instance) SetCapTop(value bool) {
	class(self).SetCapTop(value)
}

func (self Instance) CapBottom() bool {
	return bool(class(self).IsCapBottom())
}

func (self Instance) SetCapBottom(value bool) {
	class(self).SetCapBottom(value)
}

func (self Instance) Curve() [1]gdclass.Curve {
	return [1]gdclass.Curve(class(self).GetCurve())
}

func (self Instance) SetCurve(value [1]gdclass.Curve) {
	class(self).SetCurve(value)
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) { //gd:TubeTrailMesh.set_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float { //gd:TubeTrailMesh.get_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadialSteps(radial_steps gd.Int) { //gd:TubeTrailMesh.set_radial_steps
	var frame = callframe.New()
	callframe.Arg(frame, radial_steps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_set_radial_steps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadialSteps() gd.Int { //gd:TubeTrailMesh.get_radial_steps
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_get_radial_steps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSections(sections gd.Int) { //gd:TubeTrailMesh.set_sections
	var frame = callframe.New()
	callframe.Arg(frame, sections)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_set_sections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSections() gd.Int { //gd:TubeTrailMesh.get_sections
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_get_sections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSectionLength(section_length gd.Float) { //gd:TubeTrailMesh.set_section_length
	var frame = callframe.New()
	callframe.Arg(frame, section_length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_set_section_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSectionLength() gd.Float { //gd:TubeTrailMesh.get_section_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_get_section_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSectionRings(section_rings gd.Int) { //gd:TubeTrailMesh.set_section_rings
	var frame = callframe.New()
	callframe.Arg(frame, section_rings)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_set_section_rings, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSectionRings() gd.Int { //gd:TubeTrailMesh.get_section_rings
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_get_section_rings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCapTop(cap_top bool) { //gd:TubeTrailMesh.set_cap_top
	var frame = callframe.New()
	callframe.Arg(frame, cap_top)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_set_cap_top, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCapTop() bool { //gd:TubeTrailMesh.is_cap_top
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_is_cap_top, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCapBottom(cap_bottom bool) { //gd:TubeTrailMesh.set_cap_bottom
	var frame = callframe.New()
	callframe.Arg(frame, cap_bottom)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_set_cap_bottom, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCapBottom() bool { //gd:TubeTrailMesh.is_cap_bottom
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_is_cap_bottom, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurve(curve [1]gdclass.Curve) { //gd:TubeTrailMesh.set_curve
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() [1]gdclass.Curve { //gd:TubeTrailMesh.get_curve
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TubeTrailMesh.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsTubeTrailMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTubeTrailMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPrimitiveMesh() PrimitiveMesh.Advanced {
	return *((*PrimitiveMesh.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPrimitiveMesh() PrimitiveMesh.Instance {
	return *((*PrimitiveMesh.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMesh() Mesh.Advanced    { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(PrimitiveMesh.Advanced(self.AsPrimitiveMesh()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PrimitiveMesh.Instance(self.AsPrimitiveMesh()), name)
	}
}
func init() {
	gdclass.Register("TubeTrailMesh", func(ptr gd.Object) any {
		return [1]gdclass.TubeTrailMesh{*(*gdclass.TubeTrailMesh)(unsafe.Pointer(&ptr))}
	})
}
