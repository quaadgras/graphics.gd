// Package PinJoint3D provides methods for working with PinJoint3D object instances.
package PinJoint3D

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
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
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

/*
A physics joint that attaches two 3D physics bodies at a single point, allowing them to freely rotate. For example, a [RigidBody3D] can be attached to a [StaticBody3D] to create a pendulum or a seesaw.
*/
type Instance [1]gdclass.PinJoint3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPinJoint3D() Instance
}

/*
Sets the value of the specified parameter.
*/
func (self Instance) SetParam(param gdclass.PinJoint3DParam, value Float.X) { //gd:PinJoint3D.set_param
	class(self).SetParam(param, gd.Float(value))
}

/*
Returns the value of the specified parameter.
*/
func (self Instance) GetParam(param gdclass.PinJoint3DParam) Float.X { //gd:PinJoint3D.get_param
	return Float.X(Float.X(class(self).GetParam(param)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PinJoint3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PinJoint3D"))
	casted := Instance{*(*gdclass.PinJoint3D)(unsafe.Pointer(&object))}
	return casted
}

/*
Sets the value of the specified parameter.
*/
//go:nosplit
func (self class) SetParam(param gdclass.PinJoint3DParam, value gd.Float) { //gd:PinJoint3D.set_param
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the value of the specified parameter.
*/
//go:nosplit
func (self class) GetParam(param gdclass.PinJoint3DParam) gd.Float { //gd:PinJoint3D.get_param
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPinJoint3D() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPinJoint3D() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("PinJoint3D", func(ptr gd.Object) any { return [1]gdclass.PinJoint3D{*(*gdclass.PinJoint3D)(unsafe.Pointer(&ptr))} })
}

type Param = gdclass.PinJoint3DParam //gd:PinJoint3D.Param

const (
	/*The force with which the pinned objects stay in positional relation to each other. The higher, the stronger.*/
	ParamBias Param = 0
	/*The force with which the pinned objects stay in velocity relation to each other. The higher, the stronger.*/
	ParamDamping Param = 1
	/*If above 0, this value is the maximum value for an impulse that this Joint3D produces.*/
	ParamImpulseClamp Param = 2
)
