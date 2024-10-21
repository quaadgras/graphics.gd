package Generic6DOFJoint3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
The [Generic6DOFJoint3D] (6 Degrees Of Freedom) joint allows for implementing custom types of joints by locking the rotation and translation of certain axes.
The first 3 DOF represent the linear motion of the physics bodies and the last 3 DOF represent the angular motion of the physics bodies. Each axis can be either locked, or limited.

*/
type Simple [1]classdb.Generic6DOFJoint3D
func (self Simple) SetParamX(param classdb.Generic6DOFJoint3DParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParamX(param, gd.Float(value))
}
func (self Simple) GetParamX(param classdb.Generic6DOFJoint3DParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParamX(param)))
}
func (self Simple) SetParamY(param classdb.Generic6DOFJoint3DParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParamY(param, gd.Float(value))
}
func (self Simple) GetParamY(param classdb.Generic6DOFJoint3DParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParamY(param)))
}
func (self Simple) SetParamZ(param classdb.Generic6DOFJoint3DParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParamZ(param, gd.Float(value))
}
func (self Simple) GetParamZ(param classdb.Generic6DOFJoint3DParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParamZ(param)))
}
func (self Simple) SetFlagX(flag classdb.Generic6DOFJoint3DFlag, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlagX(flag, value)
}
func (self Simple) GetFlagX(flag classdb.Generic6DOFJoint3DFlag) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlagX(flag))
}
func (self Simple) SetFlagY(flag classdb.Generic6DOFJoint3DFlag, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlagY(flag, value)
}
func (self Simple) GetFlagY(flag classdb.Generic6DOFJoint3DFlag) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlagY(flag))
}
func (self Simple) SetFlagZ(flag classdb.Generic6DOFJoint3DFlag, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlagZ(flag, value)
}
func (self Simple) GetFlagZ(flag classdb.Generic6DOFJoint3DFlag) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFlagZ(flag))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Generic6DOFJoint3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetParamX(param classdb.Generic6DOFJoint3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_set_param_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParamX(param classdb.Generic6DOFJoint3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_get_param_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParamY(param classdb.Generic6DOFJoint3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_set_param_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParamY(param classdb.Generic6DOFJoint3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_get_param_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParamZ(param classdb.Generic6DOFJoint3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_set_param_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParamZ(param classdb.Generic6DOFJoint3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_get_param_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlagX(flag classdb.Generic6DOFJoint3DFlag, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_set_flag_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlagX(flag classdb.Generic6DOFJoint3DFlag) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_get_flag_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlagY(flag classdb.Generic6DOFJoint3DFlag, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_set_flag_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlagY(flag classdb.Generic6DOFJoint3DFlag) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_get_flag_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlagZ(flag classdb.Generic6DOFJoint3DFlag, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_set_flag_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlagZ(flag classdb.Generic6DOFJoint3DFlag) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Generic6DOFJoint3D.Bind_get_flag_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsGeneric6DOFJoint3D() Expert { return self[0].AsGeneric6DOFJoint3D() }


//go:nosplit
func (self Simple) AsGeneric6DOFJoint3D() Simple { return self[0].AsGeneric6DOFJoint3D() }


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
func init() {classdb.Register("Generic6DOFJoint3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Param = classdb.Generic6DOFJoint3DParam

const (
/*The minimum difference between the pivot points' axes.*/
	ParamLinearLowerLimit Param = 0
/*The maximum difference between the pivot points' axes.*/
	ParamLinearUpperLimit Param = 1
/*A factor applied to the movement across the axes. The lower, the slower the movement.*/
	ParamLinearLimitSoftness Param = 2
/*The amount of restitution on the axes' movement. The lower, the more momentum gets lost.*/
	ParamLinearRestitution Param = 3
/*The amount of damping that happens at the linear motion across the axes.*/
	ParamLinearDamping Param = 4
/*The velocity the linear motor will try to reach.*/
	ParamLinearMotorTargetVelocity Param = 5
/*The maximum force the linear motor will apply while trying to reach the velocity target.*/
	ParamLinearMotorForceLimit Param = 6
	ParamLinearSpringStiffness Param = 7
	ParamLinearSpringDamping Param = 8
	ParamLinearSpringEquilibriumPoint Param = 9
/*The minimum rotation in negative direction to break loose and rotate around the axes.*/
	ParamAngularLowerLimit Param = 10
/*The minimum rotation in positive direction to break loose and rotate around the axes.*/
	ParamAngularUpperLimit Param = 11
/*The speed of all rotations across the axes.*/
	ParamAngularLimitSoftness Param = 12
/*The amount of rotational damping across the axes. The lower, the more damping occurs.*/
	ParamAngularDamping Param = 13
/*The amount of rotational restitution across the axes. The lower, the more restitution occurs.*/
	ParamAngularRestitution Param = 14
/*The maximum amount of force that can occur, when rotating around the axes.*/
	ParamAngularForceLimit Param = 15
/*When rotating across the axes, this error tolerance factor defines how much the correction gets slowed down. The lower, the slower.*/
	ParamAngularErp Param = 16
/*Target speed for the motor at the axes.*/
	ParamAngularMotorTargetVelocity Param = 17
/*Maximum acceleration for the motor at the axes.*/
	ParamAngularMotorForceLimit Param = 18
	ParamAngularSpringStiffness Param = 19
	ParamAngularSpringDamping Param = 20
	ParamAngularSpringEquilibriumPoint Param = 21
/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 22
)
type Flag = classdb.Generic6DOFJoint3DFlag

const (
/*If enabled, linear motion is possible within the given limits.*/
	FlagEnableLinearLimit Flag = 0
/*If enabled, rotational motion is possible within the given limits.*/
	FlagEnableAngularLimit Flag = 1
	FlagEnableLinearSpring Flag = 3
	FlagEnableAngularSpring Flag = 2
/*If enabled, there is a rotational motor across these axes.*/
	FlagEnableMotor Flag = 4
/*If enabled, there is a linear motor across these axes.*/
	FlagEnableLinearMotor Flag = 5
/*Represents the size of the [enum Flag] enum.*/
	FlagMax Flag = 6
)
