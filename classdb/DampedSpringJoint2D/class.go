// Package DampedSpringJoint2D provides methods for working with DampedSpringJoint2D object instances.
package DampedSpringJoint2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
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

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
A physics joint that connects two 2D physics bodies with a spring-like force. This resembles a spring that always wants to stretch to a given length.
*/
type Instance [1]gdclass.DampedSpringJoint2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsDampedSpringJoint2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.DampedSpringJoint2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("DampedSpringJoint2D"))
	casted := Instance{*(*gdclass.DampedSpringJoint2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Length() Float.X {
	return Float.X(Float.X(class(self).GetLength()))
}

func (self Instance) SetLength(value Float.X) {
	class(self).SetLength(gd.Float(value))
}

func (self Instance) RestLength() Float.X {
	return Float.X(Float.X(class(self).GetRestLength()))
}

func (self Instance) SetRestLength(value Float.X) {
	class(self).SetRestLength(gd.Float(value))
}

func (self Instance) Stiffness() Float.X {
	return Float.X(Float.X(class(self).GetStiffness()))
}

func (self Instance) SetStiffness(value Float.X) {
	class(self).SetStiffness(gd.Float(value))
}

func (self Instance) Damping() Float.X {
	return Float.X(Float.X(class(self).GetDamping()))
}

func (self Instance) SetDamping(value Float.X) {
	class(self).SetDamping(gd.Float(value))
}

//go:nosplit
func (self class) SetLength(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRestLength(rest_length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, rest_length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_rest_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRestLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_rest_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStiffness(stiffness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, stiffness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStiffness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDamping(damping gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, damping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDamping() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsDampedSpringJoint2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsDampedSpringJoint2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsJoint2D() Joint2D.Advanced        { return *((*Joint2D.Advanced)(unsafe.Pointer(&self))) }
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
	gdclass.Register("DampedSpringJoint2D", func(ptr gd.Object) any {
		return [1]gdclass.DampedSpringJoint2D{*(*gdclass.DampedSpringJoint2D)(unsafe.Pointer(&ptr))}
	})
}
