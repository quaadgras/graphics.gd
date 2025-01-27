// Package PinJoint2D provides methods for working with PinJoint2D object instances.
package PinJoint2D

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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Joint2D"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
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

/*
A physics joint that attaches two 2D physics bodies at a single point, allowing them to freely rotate. For example, a [RigidBody2D] can be attached to a [StaticBody2D] to create a pendulum or a seesaw.
*/
type Instance [1]gdclass.PinJoint2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPinJoint2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PinJoint2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PinJoint2D"))
	casted := Instance{*(*gdclass.PinJoint2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Softness() Float.X {
	return Float.X(Float.X(class(self).GetSoftness()))
}

func (self Instance) SetSoftness(value Float.X) {
	class(self).SetSoftness(gd.Float(value))
}

func (self Instance) AngularLimitEnabled() bool {
	return bool(class(self).IsAngularLimitEnabled())
}

func (self Instance) SetAngularLimitEnabled(value bool) {
	class(self).SetAngularLimitEnabled(value)
}

func (self Instance) AngularLimitLower() Float.X {
	return Float.X(Float.X(class(self).GetAngularLimitLower()))
}

func (self Instance) SetAngularLimitLower(value Float.X) {
	class(self).SetAngularLimitLower(gd.Float(value))
}

func (self Instance) AngularLimitUpper() Float.X {
	return Float.X(Float.X(class(self).GetAngularLimitUpper()))
}

func (self Instance) SetAngularLimitUpper(value Float.X) {
	class(self).SetAngularLimitUpper(gd.Float(value))
}

func (self Instance) MotorEnabled() bool {
	return bool(class(self).IsMotorEnabled())
}

func (self Instance) SetMotorEnabled(value bool) {
	class(self).SetMotorEnabled(value)
}

func (self Instance) MotorTargetVelocity() Float.X {
	return Float.X(Float.X(class(self).GetMotorTargetVelocity()))
}

func (self Instance) SetMotorTargetVelocity(value Float.X) {
	class(self).SetMotorTargetVelocity(gd.Float(value))
}

//go:nosplit
func (self class) SetSoftness(softness gd.Float) { //gd:PinJoint2D.set_softness
	var frame = callframe.New()
	callframe.Arg(frame, softness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_set_softness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSoftness() gd.Float { //gd:PinJoint2D.get_softness
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_get_softness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularLimitLower(angular_limit_lower gd.Float) { //gd:PinJoint2D.set_angular_limit_lower
	var frame = callframe.New()
	callframe.Arg(frame, angular_limit_lower)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_set_angular_limit_lower, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularLimitLower() gd.Float { //gd:PinJoint2D.get_angular_limit_lower
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_get_angular_limit_lower, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularLimitUpper(angular_limit_upper gd.Float) { //gd:PinJoint2D.set_angular_limit_upper
	var frame = callframe.New()
	callframe.Arg(frame, angular_limit_upper)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_set_angular_limit_upper, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularLimitUpper() gd.Float { //gd:PinJoint2D.get_angular_limit_upper
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_get_angular_limit_upper, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMotorTargetVelocity(motor_target_velocity gd.Float) { //gd:PinJoint2D.set_motor_target_velocity
	var frame = callframe.New()
	callframe.Arg(frame, motor_target_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_set_motor_target_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotorTargetVelocity() gd.Float { //gd:PinJoint2D.get_motor_target_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_get_motor_target_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMotorEnabled(enabled bool) { //gd:PinJoint2D.set_motor_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_set_motor_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMotorEnabled() bool { //gd:PinJoint2D.is_motor_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_is_motor_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularLimitEnabled(enabled bool) { //gd:PinJoint2D.set_angular_limit_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_set_angular_limit_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAngularLimitEnabled() bool { //gd:PinJoint2D.is_angular_limit_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PinJoint2D.Bind_is_angular_limit_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPinJoint2D() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPinJoint2D() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsJoint2D() Joint2D.Advanced { return *((*Joint2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJoint2D() Joint2D.Instance {
	return *((*Joint2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Joint2D.Advanced(self.AsJoint2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Joint2D.Instance(self.AsJoint2D()), name)
	}
}
func init() {
	gdclass.Register("PinJoint2D", func(ptr gd.Object) any { return [1]gdclass.PinJoint2D{*(*gdclass.PinJoint2D)(unsafe.Pointer(&ptr))} })
}
