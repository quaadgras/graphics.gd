// Package Compositor provides methods for working with Compositor object instances.
package Compositor

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The compositor resource stores attributes used to customize how a [Viewport] is rendered.
*/
type Instance [1]gdclass.Compositor
type Any interface {
	gd.IsClass
	AsCompositor() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Compositor

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Compositor"))
	return Instance{*(*gdclass.Compositor)(unsafe.Pointer(&object))}
}

func (self Instance) CompositorEffects() gd.Array {
	return gd.Array(class(self).GetCompositorEffects())
}

func (self Instance) SetCompositorEffects(value gd.Array) {
	class(self).SetCompositorEffects(value)
}

//go:nosplit
func (self class) SetCompositorEffects(compositor_effects gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(compositor_effects))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Compositor.Bind_set_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCompositorEffects() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Compositor.Bind_get_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsCompositor() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCompositor() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	gdclass.Register("Compositor", func(ptr gd.Object) any { return [1]gdclass.Compositor{*(*gdclass.Compositor)(unsafe.Pointer(&ptr))} })
}
