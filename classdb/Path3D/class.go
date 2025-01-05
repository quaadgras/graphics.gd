// Package Path3D provides methods for working with Path3D object instances.
package Path3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Can have [PathFollow3D] child nodes moving along the [Curve3D]. See [PathFollow3D] for more information on the usage.
Note that the path is considered as relative to the moved nodes (children of [PathFollow3D]). As such, the curve should usually start with a zero vector [code](0, 0, 0)[/code].
*/
type Instance [1]gdclass.Path3D
type Any interface {
	gd.IsClass
	AsPath3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Path3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Path3D"))
	return Instance{*(*gdclass.Path3D)(unsafe.Pointer(&object))}
}

func (self Instance) Curve() [1]gdclass.Curve3D {
	return [1]gdclass.Curve3D(class(self).GetCurve())
}

func (self Instance) SetCurve(value [1]gdclass.Curve3D) {
	class(self).SetCurve(value)
}

//go:nosplit
func (self class) SetCurve(curve [1]gdclass.Curve3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Path3D.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() [1]gdclass.Curve3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Path3D.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Curve3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve3D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self Instance) OnCurveChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("curve_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsPath3D() Advanced           { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPath3D() Instance        { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("Path3D", func(ptr gd.Object) any { return [1]gdclass.Path3D{*(*gdclass.Path3D)(unsafe.Pointer(&ptr))} })
}
