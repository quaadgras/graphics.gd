package PinJoint3D

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
A physics joint that attaches two 3D physics bodies at a single point, allowing them to freely rotate. For example, a [RigidBody3D] can be attached to a [StaticBody3D] to create a pendulum or a seesaw.

*/
type Simple [1]classdb.PinJoint3D
func (self Simple) SetParam(param classdb.PinJoint3DParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParam(param, gd.Float(value))
}
func (self Simple) GetParam(param classdb.PinJoint3DParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParam(param)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PinJoint3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the value of the specified parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.PinJoint3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.PinJoint3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PinJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPinJoint3D() Expert { return self[0].AsPinJoint3D() }


//go:nosplit
func (self Simple) AsPinJoint3D() Simple { return self[0].AsPinJoint3D() }


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
func init() {classdb.Register("PinJoint3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Param = classdb.PinJoint3DParam

const (
/*The force with which the pinned objects stay in positional relation to each other. The higher, the stronger.*/
	ParamBias Param = 0
/*The force with which the pinned objects stay in velocity relation to each other. The higher, the stronger.*/
	ParamDamping Param = 1
/*If above 0, this value is the maximum value for an impulse that this Joint3D produces.*/
	ParamImpulseClamp Param = 2
)
