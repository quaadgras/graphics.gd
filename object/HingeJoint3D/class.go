package HingeJoint3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Joint3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A physics joint that restricts the rotation of a 3D physics body around an axis relative to another physics body. For example, Body A can be a [StaticBody3D] representing a door hinge that a [RigidBody3D] rotates around.

*/
type Simple [1]classdb.HingeJoint3D
func (self Simple) SetParam(param classdb.HingeJoint3DParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParam(param, gd.Float(value))
}
func (self Simple) GetParam(param classdb.HingeJoint3DParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParam(param)))
}
func (self Simple) SetFlag(flag classdb.HingeJoint3DFlag, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlag(flag, enabled)
}
func (self Simple) GetFlag(flag classdb.HingeJoint3DFlag) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlag(flag))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.HingeJoint3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the value of the specified parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.HingeJoint3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HingeJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.HingeJoint3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HingeJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], enables the specified flag.
*/
//go:nosplit
func (self class) SetFlag(flag classdb.HingeJoint3DFlag, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HingeJoint3D.Bind_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified flag.
*/
//go:nosplit
func (self class) GetFlag(flag classdb.HingeJoint3DFlag) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HingeJoint3D.Bind_get_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsHingeJoint3D() Expert { return self[0].AsHingeJoint3D() }


//go:nosplit
func (self Simple) AsHingeJoint3D() Simple { return self[0].AsHingeJoint3D() }


//go:nosplit
func (self class) AsJoint3D() Joint3D.Expert { return self[0].AsJoint3D() }


//go:nosplit
func (self Simple) AsJoint3D() Joint3D.Simple { return self[0].AsJoint3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


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
func init() {classdb.Register("HingeJoint3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Param = classdb.HingeJoint3DParam

const (
/*The speed with which the two bodies get pulled together when they move in different directions.*/
	ParamBias Param = 0
/*The maximum rotation. Only active if [member angular_limit/enable] is [code]true[/code].*/
	ParamLimitUpper Param = 1
/*The minimum rotation. Only active if [member angular_limit/enable] is [code]true[/code].*/
	ParamLimitLower Param = 2
/*The speed with which the rotation across the axis perpendicular to the hinge gets corrected.*/
	ParamLimitBias Param = 3
	ParamLimitSoftness Param = 4
/*The lower this value, the more the rotation gets slowed down.*/
	ParamLimitRelaxation Param = 5
/*Target speed for the motor.*/
	ParamMotorTargetVelocity Param = 6
/*Maximum acceleration for the motor.*/
	ParamMotorMaxImpulse Param = 7
/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 8
)
type Flag = classdb.HingeJoint3DFlag

const (
/*If [code]true[/code], the hinges maximum and minimum rotation, defined by [member angular_limit/lower] and [member angular_limit/upper] has effects.*/
	FlagUseLimit Flag = 0
/*When activated, a motor turns the hinge.*/
	FlagEnableMotor Flag = 1
/*Represents the size of the [enum Flag] enum.*/
	FlagMax Flag = 2
)
