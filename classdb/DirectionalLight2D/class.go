// Package DirectionalLight2D provides methods for working with DirectionalLight2D object instances.
package DirectionalLight2D

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
import "graphics.gd/classdb/Light2D"
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

/*
A directional light is a type of [Light2D] node that models an infinite number of parallel rays covering the entire scene. It is used for lights with strong intensity that are located far away from the scene (for example: to model sunlight or moonlight).
[b]Note:[/b] [DirectionalLight2D] does not support light cull masks (but it supports shadow cull masks). It will always light up 2D nodes, regardless of the 2D node's [member CanvasItem.light_mask].
*/
type Instance [1]gdclass.DirectionalLight2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsDirectionalLight2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.DirectionalLight2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("DirectionalLight2D"))
	casted := Instance{*(*gdclass.DirectionalLight2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) MaxDistance() Float.X {
	return Float.X(Float.X(class(self).GetMaxDistance()))
}

func (self Instance) SetMaxDistance(value Float.X) {
	class(self).SetMaxDistance(gd.Float(value))
}

//go:nosplit
func (self class) SetMaxDistance(pixels gd.Float) { //gd:DirectionalLight2D.set_max_distance
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirectionalLight2D.Bind_set_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxDistance() gd.Float { //gd:DirectionalLight2D.get_max_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirectionalLight2D.Bind_get_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsDirectionalLight2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsDirectionalLight2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsLight2D() Light2D.Advanced       { return *((*Light2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLight2D() Light2D.Instance {
	return *((*Light2D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Light2D.Advanced(self.AsLight2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Light2D.Instance(self.AsLight2D()), name)
	}
}
func init() {
	gdclass.Register("DirectionalLight2D", func(ptr gd.Object) any {
		return [1]gdclass.DirectionalLight2D{*(*gdclass.DirectionalLight2D)(unsafe.Pointer(&ptr))}
	})
}
