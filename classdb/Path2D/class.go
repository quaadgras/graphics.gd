// Package Path2D provides methods for working with Path2D object instances.
package Path2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Can have [PathFollow2D] child nodes moving along the [Curve2D]. See [PathFollow2D] for more information on usage.
[b]Note:[/b] The path is considered as relative to the moved nodes (children of [PathFollow2D]). As such, the curve should usually start with a zero vector ([code](0, 0)[/code]).
*/
type Instance [1]gdclass.Path2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPath2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Path2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Path2D"))
	casted := Instance{*(*gdclass.Path2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Curve() [1]gdclass.Curve2D {
	return [1]gdclass.Curve2D(class(self).GetCurve())
}

func (self Instance) SetCurve(value [1]gdclass.Curve2D) {
	class(self).SetCurve(value)
}

//go:nosplit
func (self class) SetCurve(curve [1]gdclass.Curve2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Path2D.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurve() [1]gdclass.Curve2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Path2D.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Curve2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Curve2D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsPath2D() Advanced           { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPath2D() Instance        { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("Path2D", func(ptr gd.Object) any { return [1]gdclass.Path2D{*(*gdclass.Path2D)(unsafe.Pointer(&ptr))} })
}
