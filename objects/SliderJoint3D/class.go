package SliderJoint3D

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
A physics joint that restricts the movement of a 3D physics body along an axis relative to another physics body. For example, Body A could be a [StaticBody3D] representing a piston base, while Body B could be a [RigidBody3D] representing the piston head, moving up and down.
*/
type Instance [1]classdb.SliderJoint3D
type Any interface {
	gd.IsClass
	AsSliderJoint3D() Instance
}

/*
Assigns [param value] to the given parameter (see [enum Param] constants).
*/
func (self Instance) SetParam(param classdb.SliderJoint3DParam, value Float.X) {
	class(self).SetParam(param, gd.Float(value))
}

/*
Returns the value of the given parameter (see [enum Param] constants).
*/
func (self Instance) GetParam(param classdb.SliderJoint3DParam) Float.X {
	return Float.X(Float.X(class(self).GetParam(param)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SliderJoint3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SliderJoint3D"))
	return Instance{classdb.SliderJoint3D(object)}
}

/*
Assigns [param value] to the given parameter (see [enum Param] constants).
*/
//go:nosplit
func (self class) SetParam(param classdb.SliderJoint3DParam, value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SliderJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the value of the given parameter (see [enum Param] constants).
*/
//go:nosplit
func (self class) GetParam(param classdb.SliderJoint3DParam) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SliderJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSliderJoint3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSliderJoint3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsJoint3D() Joint3D.Advanced  { return *((*Joint3D.Advanced)(unsafe.Pointer(&self))) }
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
	classdb.Register("SliderJoint3D", func(ptr gd.Object) any { return classdb.SliderJoint3D(ptr) })
}

type Param = classdb.SliderJoint3DParam

const (
	/*Constant for accessing [member linear_limit/upper_distance]. The maximum difference between the pivot points on their X axis before damping happens.*/
	ParamLinearLimitUpper Param = 0
	/*Constant for accessing [member linear_limit/lower_distance]. The minimum difference between the pivot points on their X axis before damping happens.*/
	ParamLinearLimitLower Param = 1
	/*Constant for accessing [member linear_limit/softness]. A factor applied to the movement across the slider axis once the limits get surpassed. The lower, the slower the movement.*/
	ParamLinearLimitSoftness Param = 2
	/*Constant for accessing [member linear_limit/restitution]. The amount of restitution once the limits are surpassed. The lower, the more velocity-energy gets lost.*/
	ParamLinearLimitRestitution Param = 3
	/*Constant for accessing [member linear_limit/damping]. The amount of damping once the slider limits are surpassed.*/
	ParamLinearLimitDamping Param = 4
	/*Constant for accessing [member linear_motion/softness]. A factor applied to the movement across the slider axis as long as the slider is in the limits. The lower, the slower the movement.*/
	ParamLinearMotionSoftness Param = 5
	/*Constant for accessing [member linear_motion/restitution]. The amount of restitution inside the slider limits.*/
	ParamLinearMotionRestitution Param = 6
	/*Constant for accessing [member linear_motion/damping]. The amount of damping inside the slider limits.*/
	ParamLinearMotionDamping Param = 7
	/*Constant for accessing [member linear_ortho/softness]. A factor applied to the movement across axes orthogonal to the slider.*/
	ParamLinearOrthogonalSoftness Param = 8
	/*Constant for accessing [member linear_motion/restitution]. The amount of restitution when movement is across axes orthogonal to the slider.*/
	ParamLinearOrthogonalRestitution Param = 9
	/*Constant for accessing [member linear_motion/damping]. The amount of damping when movement is across axes orthogonal to the slider.*/
	ParamLinearOrthogonalDamping Param = 10
	/*Constant for accessing [member angular_limit/upper_angle]. The upper limit of rotation in the slider.*/
	ParamAngularLimitUpper Param = 11
	/*Constant for accessing [member angular_limit/lower_angle]. The lower limit of rotation in the slider.*/
	ParamAngularLimitLower Param = 12
	/*Constant for accessing [member angular_limit/softness]. A factor applied to the all rotation once the limit is surpassed.*/
	ParamAngularLimitSoftness Param = 13
	/*Constant for accessing [member angular_limit/restitution]. The amount of restitution of the rotation when the limit is surpassed.*/
	ParamAngularLimitRestitution Param = 14
	/*Constant for accessing [member angular_limit/damping]. The amount of damping of the rotation when the limit is surpassed.*/
	ParamAngularLimitDamping Param = 15
	/*Constant for accessing [member angular_motion/softness]. A factor applied to the all rotation in the limits.*/
	ParamAngularMotionSoftness Param = 16
	/*Constant for accessing [member angular_motion/restitution]. The amount of restitution of the rotation in the limits.*/
	ParamAngularMotionRestitution Param = 17
	/*Constant for accessing [member angular_motion/damping]. The amount of damping of the rotation in the limits.*/
	ParamAngularMotionDamping Param = 18
	/*Constant for accessing [member angular_ortho/softness]. A factor applied to the all rotation across axes orthogonal to the slider.*/
	ParamAngularOrthogonalSoftness Param = 19
	/*Constant for accessing [member angular_ortho/restitution]. The amount of restitution of the rotation across axes orthogonal to the slider.*/
	ParamAngularOrthogonalRestitution Param = 20
	/*Constant for accessing [member angular_ortho/damping]. The amount of damping of the rotation across axes orthogonal to the slider.*/
	ParamAngularOrthogonalDamping Param = 21
	/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 22
)
