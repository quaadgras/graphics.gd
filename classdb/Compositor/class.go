// Package Compositor provides methods for working with Compositor object instances.
package Compositor

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
import "graphics.gd/classdb/Resource"

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

/*
The compositor resource stores attributes used to customize how a [Viewport] is rendered.
*/
type Instance [1]gdclass.Compositor

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCompositor() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Compositor

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Compositor"))
	casted := Instance{*(*gdclass.Compositor)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) CompositorEffects() [][1]gdclass.CompositorEffect {
	return [][1]gdclass.CompositorEffect(gd.ArrayAs[[][1]gdclass.CompositorEffect](gd.InternalArray(class(self).GetCompositorEffects())))
}

func (self Instance) SetCompositorEffects(value [][1]gdclass.CompositorEffect) {
	class(self).SetCompositorEffects(gd.ArrayFromSlice[Array.Contains[[1]gdclass.CompositorEffect]](value))
}

//go:nosplit
func (self class) SetCompositorEffects(compositor_effects Array.Contains[[1]gdclass.CompositorEffect]) { //gd:Compositor.set_compositor_effects
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(compositor_effects)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Compositor.Bind_set_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCompositorEffects() Array.Contains[[1]gdclass.CompositorEffect] { //gd:Compositor.get_compositor_effects
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Compositor.Bind_get_compositor_effects, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.CompositorEffect]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Compositor", func(ptr gd.Object) any { return [1]gdclass.Compositor{*(*gdclass.Compositor)(unsafe.Pointer(&ptr))} })
}
