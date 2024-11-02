package PhysicsMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Holds physics-related properties of a surface, namely its roughness and bounciness. This class is used to apply these properties to a physics body.
*/
type Instance [1]classdb.PhysicsMaterial

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PhysicsMaterial

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsMaterial"))
	return Instance{classdb.PhysicsMaterial(object)}
}

func (self Instance) Friction() float64 {
	return float64(float64(class(self).GetFriction()))
}

func (self Instance) SetFriction(value float64) {
	class(self).SetFriction(gd.Float(value))
}

func (self Instance) Rough() bool {
	return bool(class(self).IsRough())
}

func (self Instance) SetRough(value bool) {
	class(self).SetRough(value)
}

func (self Instance) Bounce() float64 {
	return float64(float64(class(self).GetBounce()))
}

func (self Instance) SetBounce(value float64) {
	class(self).SetBounce(gd.Float(value))
}

func (self Instance) Absorbent() bool {
	return bool(class(self).IsAbsorbent())
}

func (self Instance) SetAbsorbent(value bool) {
	class(self).SetAbsorbent(value)
}

//go:nosplit
func (self class) SetFriction(friction gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, friction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_set_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFriction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_get_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRough(rough bool) {
	var frame = callframe.New()
	callframe.Arg(frame, rough)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_set_rough, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsRough() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_is_rough, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBounce(bounce gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bounce)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_set_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBounce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_get_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAbsorbent(absorbent bool) {
	var frame = callframe.New()
	callframe.Arg(frame, absorbent)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_set_absorbent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAbsorbent() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_is_absorbent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicsMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("PhysicsMaterial", func(ptr gd.Object) any { return classdb.PhysicsMaterial(ptr) })
}
