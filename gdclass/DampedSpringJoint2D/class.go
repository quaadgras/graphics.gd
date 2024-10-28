package DampedSpringJoint2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Joint2D"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A physics joint that connects two 2D physics bodies with a spring-like force. This resembles a spring that always wants to stretch to a given length.

*/
type Go [1]classdb.DampedSpringJoint2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.DampedSpringJoint2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("DampedSpringJoint2D"))
	return Go{classdb.DampedSpringJoint2D(object)}
}

func (self Go) Length() float64 {
		return float64(float64(class(self).GetLength()))
}

func (self Go) SetLength(value float64) {
	class(self).SetLength(gd.Float(value))
}

func (self Go) RestLength() float64 {
		return float64(float64(class(self).GetRestLength()))
}

func (self Go) SetRestLength(value float64) {
	class(self).SetRestLength(gd.Float(value))
}

func (self Go) Stiffness() float64 {
		return float64(float64(class(self).GetStiffness()))
}

func (self Go) SetStiffness(value float64) {
	class(self).SetStiffness(gd.Float(value))
}

func (self Go) Damping() float64 {
		return float64(float64(class(self).GetDamping()))
}

func (self Go) SetDamping(value float64) {
	class(self).SetDamping(gd.Float(value))
}

//go:nosplit
func (self class) SetLength(length gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRestLength(rest_length gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, rest_length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_rest_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRestLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_rest_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStiffness(stiffness gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, stiffness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStiffness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_stiffness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDamping(damping gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, damping)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDamping() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_damping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsDampedSpringJoint2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsDampedSpringJoint2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsJoint2D() Joint2D.GD { return *((*Joint2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsJoint2D() Joint2D.Go { return *((*Joint2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsJoint2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsJoint2D(), name)
	}
}
func init() {classdb.Register("DampedSpringJoint2D", func(ptr gd.Object) any { return classdb.DampedSpringJoint2D(ptr) })}
