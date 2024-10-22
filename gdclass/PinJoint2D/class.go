package PinJoint2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Joint2D"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A physics joint that attaches two 2D physics bodies at a single point, allowing them to freely rotate. For example, a [RigidBody2D] can be attached to a [StaticBody2D] to create a pendulum or a seesaw.

*/
type Go [1]classdb.PinJoint2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PinJoint2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PinJoint2D"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Softness() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetSoftness()))
}

func (self Go) SetSoftness(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSoftness(gd.Float(value))
}

func (self Go) AngularLimitEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAngularLimitEnabled())
}

func (self Go) SetAngularLimitEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularLimitEnabled(value)
}

func (self Go) AngularLimitLower() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAngularLimitLower()))
}

func (self Go) SetAngularLimitLower(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularLimitLower(gd.Float(value))
}

func (self Go) AngularLimitUpper() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAngularLimitUpper()))
}

func (self Go) SetAngularLimitUpper(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAngularLimitUpper(gd.Float(value))
}

func (self Go) MotorEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMotorEnabled())
}

func (self Go) SetMotorEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMotorEnabled(value)
}

func (self Go) MotorTargetVelocity() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetMotorTargetVelocity()))
}

func (self Go) SetMotorTargetVelocity(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMotorTargetVelocity(gd.Float(value))
}

//go:nosplit
func (self class) SetSoftness(softness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, softness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_set_softness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSoftness() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_get_softness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularLimitLower(angular_limit_lower gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_limit_lower)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_set_angular_limit_lower, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularLimitLower() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_get_angular_limit_lower, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularLimitUpper(angular_limit_upper gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, angular_limit_upper)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_set_angular_limit_upper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularLimitUpper() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_get_angular_limit_upper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMotorTargetVelocity(motor_target_velocity gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, motor_target_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_set_motor_target_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMotorTargetVelocity() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_get_motor_target_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMotorEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_set_motor_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMotorEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_is_motor_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularLimitEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_set_angular_limit_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAngularLimitEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint2D.Bind_is_angular_limit_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPinJoint2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPinJoint2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsJoint2D() Joint2D.GD { return *((*Joint2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsJoint2D() Joint2D.Go { return *((*Joint2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsJoint2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsJoint2D(), name)
	}
}
func init() {classdb.Register("PinJoint2D", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
