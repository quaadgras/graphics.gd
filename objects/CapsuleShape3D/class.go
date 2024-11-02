package CapsuleShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Shape3D"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A 3D capsule shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D].
[b]Performance:[/b] [CapsuleShape3D] is fast to check collisions against. It is faster than [CylinderShape3D], but slower than [SphereShape3D] and [BoxShape3D].
*/
type Instance [1]classdb.CapsuleShape3D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CapsuleShape3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CapsuleShape3D"))
	return Instance{classdb.CapsuleShape3D(object)}
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

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleShape3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleShape3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeight(height gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleShape3D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleShape3D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCapsuleShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCapsuleShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.Advanced   { return *((*Shape3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape3D() Shape3D.Instance {
	return *((*Shape3D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsShape3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShape3D(), name)
	}
}
func init() {
	classdb.Register("CapsuleShape3D", func(ptr gd.Object) any { return classdb.CapsuleShape3D(ptr) })
}
