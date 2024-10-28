package ConeTwistJoint3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Joint3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A physics joint that connects two 3D physics bodies in a way that simulates a ball-and-socket joint. The twist axis is initiated as the X axis of the [ConeTwistJoint3D]. Once the physics bodies swing, the twist axis is calculated as the middle of the X axes of the joint in the local space of the two physics bodies. Useful for limbs like shoulders and hips, lamps hanging off a ceiling, etc.

*/
type Go [1]classdb.ConeTwistJoint3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ConeTwistJoint3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ConeTwistJoint3D"))
	return Go{classdb.ConeTwistJoint3D(object)}
}

func (self Go) SwingSpan() float64 {
		return float64(float64(class(self).GetParam(0)))
}

func (self Go) SetSwingSpan(value float64) {
	class(self).SetParam(0, gd.Float(value))
}

func (self Go) TwistSpan() float64 {
		return float64(float64(class(self).GetParam(1)))
}

func (self Go) SetTwistSpan(value float64) {
	class(self).SetParam(1, gd.Float(value))
}

func (self Go) Bias() float64 {
		return float64(float64(class(self).GetParam(2)))
}

func (self Go) SetBias(value float64) {
	class(self).SetParam(2, gd.Float(value))
}

func (self Go) Softness() float64 {
		return float64(float64(class(self).GetParam(3)))
}

func (self Go) SetSoftness(value float64) {
	class(self).SetParam(3, gd.Float(value))
}

func (self Go) Relaxation() float64 {
		return float64(float64(class(self).GetParam(4)))
}

func (self Go) SetRelaxation(value float64) {
	class(self).SetParam(4, gd.Float(value))
}

/*
Sets the value of the specified parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.ConeTwistJoint3DParam, value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConeTwistJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.ConeTwistJoint3DParam) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ConeTwistJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsConeTwistJoint3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsConeTwistJoint3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsJoint3D() Joint3D.GD { return *((*Joint3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsJoint3D() Joint3D.Go { return *((*Joint3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsJoint3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsJoint3D(), name)
	}
}
func init() {classdb.Register("ConeTwistJoint3D", func(ptr gd.Object) any { return classdb.ConeTwistJoint3D(ptr) })}
type Param = classdb.ConeTwistJoint3DParam

const (
/*Swing is rotation from side to side, around the axis perpendicular to the twist axis.
The swing span defines, how much rotation will not get corrected along the swing axis.
Could be defined as looseness in the [ConeTwistJoint3D].
If below 0.05, this behavior is locked.*/
	ParamSwingSpan Param = 0
/*Twist is the rotation around the twist axis, this value defined how far the joint can twist.
Twist is locked if below 0.05.*/
	ParamTwistSpan Param = 1
/*The speed with which the swing or twist will take place.
The higher, the faster.*/
	ParamBias Param = 2
/*The ease with which the joint starts to twist. If it's too low, it takes more force to start twisting the joint.*/
	ParamSoftness Param = 3
/*Defines, how fast the swing- and twist-speed-difference on both sides gets synced.*/
	ParamRelaxation Param = 4
/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 5
)
