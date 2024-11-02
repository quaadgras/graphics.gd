package CapsuleMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/PrimitiveMesh"
import "grow.graphics/gd/objects/Mesh"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Class representing a capsule-shaped [PrimitiveMesh].
*/
type Instance [1]classdb.CapsuleMesh

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CapsuleMesh

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CapsuleMesh"))
	return Instance{classdb.CapsuleMesh(object)}
}

func (self Instance) Radius() float64 {
	return float64(float64(class(self).GetRadius()))
}

func (self Instance) SetRadius(value float64) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) Height() float64 {
	return float64(float64(class(self).GetHeight()))
}

func (self Instance) SetHeight(value float64) {
	class(self).SetHeight(gd.Float(value))
}

func (self Instance) RadialSegments() int {
	return int(int(class(self).GetRadialSegments()))
}

func (self Instance) SetRadialSegments(value int) {
	class(self).SetRadialSegments(gd.Int(value))
}

func (self Instance) Rings() int {
	return int(int(class(self).GetRings()))
}

func (self Instance) SetRings(value int) {
	class(self).SetRings(gd.Int(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleMesh.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleMesh.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeight(height gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleMesh.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleMesh.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadialSegments(segments gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleMesh.Bind_set_radial_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadialSegments() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleMesh.Bind_get_radial_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRings(rings gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, rings)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleMesh.Bind_set_rings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRings() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleMesh.Bind_get_rings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCapsuleMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCapsuleMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}
func init() {
	classdb.Register("CapsuleMesh", func(ptr gd.Object) any { return classdb.CapsuleMesh(ptr) })
}
