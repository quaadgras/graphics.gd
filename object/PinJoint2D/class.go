package PinJoint2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Joint2D"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A physics joint that attaches two 2D physics bodies at a single point, allowing them to freely rotate. For example, a [RigidBody2D] can be attached to a [StaticBody2D] to create a pendulum or a seesaw.

*/
type Simple [1]classdb.PinJoint2D
func (self Simple) SetSoftness(softness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSoftness(gd.Float(softness))
}
func (self Simple) GetSoftness() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSoftness()))
}
func (self Simple) SetAngularLimitLower(angular_limit_lower float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularLimitLower(gd.Float(angular_limit_lower))
}
func (self Simple) GetAngularLimitLower() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAngularLimitLower()))
}
func (self Simple) SetAngularLimitUpper(angular_limit_upper float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularLimitUpper(gd.Float(angular_limit_upper))
}
func (self Simple) GetAngularLimitUpper() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAngularLimitUpper()))
}
func (self Simple) SetMotorTargetVelocity(motor_target_velocity float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMotorTargetVelocity(gd.Float(motor_target_velocity))
}
func (self Simple) GetMotorTargetVelocity() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMotorTargetVelocity()))
}
func (self Simple) SetMotorEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMotorEnabled(enabled)
}
func (self Simple) IsMotorEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMotorEnabled())
}
func (self Simple) SetAngularLimitEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAngularLimitEnabled(enabled)
}
func (self Simple) IsAngularLimitEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAngularLimitEnabled())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PinJoint2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsPinJoint2D() Expert { return self[0].AsPinJoint2D() }


//go:nosplit
func (self Simple) AsPinJoint2D() Simple { return self[0].AsPinJoint2D() }


//go:nosplit
func (self class) AsJoint2D() Joint2D.Expert { return self[0].AsJoint2D() }


//go:nosplit
func (self Simple) AsJoint2D() Joint2D.Simple { return self[0].AsJoint2D() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PinJoint2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
