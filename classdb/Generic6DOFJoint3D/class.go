// Package Generic6DOFJoint3D provides methods for working with Generic6DOFJoint3D object instances.
package Generic6DOFJoint3D

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/Joint3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
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
var _ Dictionary.Any

/*
The [Generic6DOFJoint3D] (6 Degrees Of Freedom) joint allows for implementing custom types of joints by locking the rotation and translation of certain axes.
The first 3 DOF represent the linear motion of the physics bodies and the last 3 DOF represent the angular motion of the physics bodies. Each axis can be either locked, or limited.
*/
type Instance [1]gdclass.Generic6DOFJoint3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGeneric6DOFJoint3D() Instance
}

func (self Instance) SetParamX(param gdclass.Generic6DOFJoint3DParam, value Float.X) { //gd:Generic6DOFJoint3D.set_param_x
	class(self).SetParamX(param, gd.Float(value))
}
func (self Instance) GetParamX(param gdclass.Generic6DOFJoint3DParam) Float.X { //gd:Generic6DOFJoint3D.get_param_x
	return Float.X(Float.X(class(self).GetParamX(param)))
}
func (self Instance) SetParamY(param gdclass.Generic6DOFJoint3DParam, value Float.X) { //gd:Generic6DOFJoint3D.set_param_y
	class(self).SetParamY(param, gd.Float(value))
}
func (self Instance) GetParamY(param gdclass.Generic6DOFJoint3DParam) Float.X { //gd:Generic6DOFJoint3D.get_param_y
	return Float.X(Float.X(class(self).GetParamY(param)))
}
func (self Instance) SetParamZ(param gdclass.Generic6DOFJoint3DParam, value Float.X) { //gd:Generic6DOFJoint3D.set_param_z
	class(self).SetParamZ(param, gd.Float(value))
}
func (self Instance) GetParamZ(param gdclass.Generic6DOFJoint3DParam) Float.X { //gd:Generic6DOFJoint3D.get_param_z
	return Float.X(Float.X(class(self).GetParamZ(param)))
}
func (self Instance) SetFlagX(flag gdclass.Generic6DOFJoint3DFlag, value bool) { //gd:Generic6DOFJoint3D.set_flag_x
	class(self).SetFlagX(flag, value)
}
func (self Instance) GetFlagX(flag gdclass.Generic6DOFJoint3DFlag) bool { //gd:Generic6DOFJoint3D.get_flag_x
	return bool(class(self).GetFlagX(flag))
}
func (self Instance) SetFlagY(flag gdclass.Generic6DOFJoint3DFlag, value bool) { //gd:Generic6DOFJoint3D.set_flag_y
	class(self).SetFlagY(flag, value)
}
func (self Instance) GetFlagY(flag gdclass.Generic6DOFJoint3DFlag) bool { //gd:Generic6DOFJoint3D.get_flag_y
	return bool(class(self).GetFlagY(flag))
}
func (self Instance) SetFlagZ(flag gdclass.Generic6DOFJoint3DFlag, value bool) { //gd:Generic6DOFJoint3D.set_flag_z
	class(self).SetFlagZ(flag, value)
}
func (self Instance) GetFlagZ(flag gdclass.Generic6DOFJoint3DFlag) bool { //gd:Generic6DOFJoint3D.get_flag_z
	return bool(class(self).GetFlagZ(flag))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Generic6DOFJoint3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Generic6DOFJoint3D"))
	casted := Instance{*(*gdclass.Generic6DOFJoint3D)(unsafe.Pointer(&object))}
	return casted
}

//go:nosplit
func (self class) SetParamX(param gdclass.Generic6DOFJoint3DParam, value gd.Float) { //gd:Generic6DOFJoint3D.set_param_x
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_param_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParamX(param gdclass.Generic6DOFJoint3DParam) gd.Float { //gd:Generic6DOFJoint3D.get_param_x
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_param_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParamY(param gdclass.Generic6DOFJoint3DParam, value gd.Float) { //gd:Generic6DOFJoint3D.set_param_y
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_param_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParamY(param gdclass.Generic6DOFJoint3DParam) gd.Float { //gd:Generic6DOFJoint3D.get_param_y
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_param_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParamZ(param gdclass.Generic6DOFJoint3DParam, value gd.Float) { //gd:Generic6DOFJoint3D.set_param_z
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_param_z, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetParamZ(param gdclass.Generic6DOFJoint3DParam) gd.Float { //gd:Generic6DOFJoint3D.get_param_z
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_param_z, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlagX(flag gdclass.Generic6DOFJoint3DFlag, value bool) { //gd:Generic6DOFJoint3D.set_flag_x
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_flag_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlagX(flag gdclass.Generic6DOFJoint3DFlag) bool { //gd:Generic6DOFJoint3D.get_flag_x
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_flag_x, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlagY(flag gdclass.Generic6DOFJoint3DFlag, value bool) { //gd:Generic6DOFJoint3D.set_flag_y
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_flag_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlagY(flag gdclass.Generic6DOFJoint3DFlag) bool { //gd:Generic6DOFJoint3D.get_flag_y
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_flag_y, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlagZ(flag gdclass.Generic6DOFJoint3DFlag, value bool) { //gd:Generic6DOFJoint3D.set_flag_z
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_flag_z, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlagZ(flag gdclass.Generic6DOFJoint3DFlag) bool { //gd:Generic6DOFJoint3D.get_flag_z
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_flag_z, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGeneric6DOFJoint3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGeneric6DOFJoint3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsJoint3D() Joint3D.Advanced       { return *((*Joint3D.Advanced)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Joint3D.Advanced(self.AsJoint3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Joint3D.Instance(self.AsJoint3D()), name)
	}
}
func init() {
	gdclass.Register("Generic6DOFJoint3D", func(ptr gd.Object) any {
		return [1]gdclass.Generic6DOFJoint3D{*(*gdclass.Generic6DOFJoint3D)(unsafe.Pointer(&ptr))}
	})
}

type Param = gdclass.Generic6DOFJoint3DParam //gd:Generic6DOFJoint3D.Param

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
	ParamLinearMotorForceLimit        Param = 6
	ParamLinearSpringStiffness        Param = 7
	ParamLinearSpringDamping          Param = 8
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
	ParamAngularMotorForceLimit        Param = 18
	ParamAngularSpringStiffness        Param = 19
	ParamAngularSpringDamping          Param = 20
	ParamAngularSpringEquilibriumPoint Param = 21
	/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 22
)

type Flag = gdclass.Generic6DOFJoint3DFlag //gd:Generic6DOFJoint3D.Flag

const (
	/*If enabled, linear motion is possible within the given limits.*/
	FlagEnableLinearLimit Flag = 0
	/*If enabled, rotational motion is possible within the given limits.*/
	FlagEnableAngularLimit  Flag = 1
	FlagEnableLinearSpring  Flag = 3
	FlagEnableAngularSpring Flag = 2
	/*If enabled, there is a rotational motor across these axes.*/
	FlagEnableMotor Flag = 4
	/*If enabled, there is a linear motor across these axes.*/
	FlagEnableLinearMotor Flag = 5
	/*Represents the size of the [enum Flag] enum.*/
	FlagMax Flag = 6
)
