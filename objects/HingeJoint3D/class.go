package HingeJoint3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Joint3D"
import "graphics.gd/objects/Node3D"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A physics joint that restricts the rotation of a 3D physics body around an axis relative to another physics body. For example, Body A can be a [StaticBody3D] representing a door hinge that a [RigidBody3D] rotates around.
*/
type Instance [1]classdb.HingeJoint3D
type Any interface {
	gd.IsClass
	AsHingeJoint3D() Instance
}

/*
Sets the value of the specified parameter.
*/
func (self Instance) SetParam(param classdb.HingeJoint3DParam, value Float.X) {
	class(self).SetParam(param, gd.Float(value))
}

/*
Returns the value of the specified parameter.
*/
func (self Instance) GetParam(param classdb.HingeJoint3DParam) Float.X {
	return Float.X(Float.X(class(self).GetParam(param)))
}

/*
If [code]true[/code], enables the specified flag.
*/
func (self Instance) SetFlag(flag classdb.HingeJoint3DFlag, enabled bool) {
	class(self).SetFlag(flag, enabled)
}

/*
Returns the value of the specified flag.
*/
func (self Instance) GetFlag(flag classdb.HingeJoint3DFlag) bool {
	return bool(class(self).GetFlag(flag))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.HingeJoint3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HingeJoint3D"))
	return Instance{classdb.HingeJoint3D(object)}
}

/*
Sets the value of the specified parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.HingeJoint3DParam, value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HingeJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the value of the specified parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.HingeJoint3DParam) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HingeJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], enables the specified flag.
*/
//go:nosplit
func (self class) SetFlag(flag classdb.HingeJoint3DFlag, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HingeJoint3D.Bind_set_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the value of the specified flag.
*/
//go:nosplit
func (self class) GetFlag(flag classdb.HingeJoint3DFlag) bool {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HingeJoint3D.Bind_get_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsHingeJoint3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHingeJoint3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsJoint3D() Joint3D.Advanced { return *((*Joint3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJoint3D() Joint3D.Instance {
	return *((*Joint3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsJoint3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsJoint3D(), name)
	}
}
func init() {
	classdb.Register("HingeJoint3D", func(ptr gd.Object) any { return [1]classdb.HingeJoint3D{classdb.HingeJoint3D(ptr)} })
}

type Param = classdb.HingeJoint3DParam

const (
	/*The speed with which the two bodies get pulled together when they move in different directions.*/
	ParamBias Param = 0
	/*The maximum rotation. Only active if [member angular_limit/enable] is [code]true[/code].*/
	ParamLimitUpper Param = 1
	/*The minimum rotation. Only active if [member angular_limit/enable] is [code]true[/code].*/
	ParamLimitLower Param = 2
	/*The speed with which the rotation across the axis perpendicular to the hinge gets corrected.*/
	ParamLimitBias     Param = 3
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
