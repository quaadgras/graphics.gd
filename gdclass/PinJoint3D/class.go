package PinJoint3D

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
A physics joint that attaches two 3D physics bodies at a single point, allowing them to freely rotate. For example, a [RigidBody3D] can be attached to a [StaticBody3D] to create a pendulum or a seesaw.

*/
type Go [1]classdb.PinJoint3D

/*
Sets the value of the specified parameter.
*/
func (self Go) SetParam(param classdb.PinJoint3DParam, value float64) {
	class(self).SetParam(param, gd.Float(value))
}

/*
Returns the value of the specified parameter.
*/
func (self Go) GetParam(param classdb.PinJoint3DParam) float64 {
	return float64(float64(class(self).GetParam(param)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PinJoint3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PinJoint3D"))
	return Go{classdb.PinJoint3D(object)}
}

/*
Sets the value of the specified parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.PinJoint3DParam, value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.PinJoint3DParam) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPinJoint3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPinJoint3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PinJoint3D", func(ptr gd.Object) any { return classdb.PinJoint3D(ptr) })}
type Param = classdb.PinJoint3DParam

const (
/*The force with which the pinned objects stay in positional relation to each other. The higher, the stronger.*/
	ParamBias Param = 0
/*The force with which the pinned objects stay in velocity relation to each other. The higher, the stronger.*/
	ParamDamping Param = 1
/*If above 0, this value is the maximum value for an impulse that this Joint3D produces.*/
	ParamImpulseClamp Param = 2
)
