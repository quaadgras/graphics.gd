package SliderJoint3D

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
A physics joint that restricts the movement of a 3D physics body along an axis relative to another physics body. For example, Body A could be a [StaticBody3D] representing a piston base, while Body B could be a [RigidBody3D] representing the piston head, moving up and down.

*/
type Simple [1]classdb.SliderJoint3D
func (self Simple) SetParam(param classdb.SliderJoint3DParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParam(param, gd.Float(value))
}
func (self Simple) GetParam(param classdb.SliderJoint3DParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParam(param)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SliderJoint3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Assigns [param value] to the given parameter (see [enum Param] constants).
*/
//go:nosplit
func (self class) SetParam(param classdb.SliderJoint3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SliderJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the given parameter (see [enum Param] constants).
*/
//go:nosplit
func (self class) GetParam(param classdb.SliderJoint3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SliderJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSliderJoint3D() Expert { return self[0].AsSliderJoint3D() }


//go:nosplit
func (self Simple) AsSliderJoint3D() Simple { return self[0].AsSliderJoint3D() }


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
func init() {classdb.Register("SliderJoint3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
