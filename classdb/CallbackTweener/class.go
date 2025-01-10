// Package CallbackTweener provides methods for working with CallbackTweener object instances.
package CallbackTweener

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Tweener"
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
[CallbackTweener] is used to call a method in a tweening sequence. See [method Tween.tween_callback] for more usage information.
The tweener will finish automatically if the callback's target object is freed.
[b]Note:[/b] [method Tween.tween_callback] is the only correct way to create [CallbackTweener]. Any [CallbackTweener] created manually will not function correctly.
*/
type Instance [1]gdclass.CallbackTweener

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCallbackTweener() Instance
}

/*
Makes the callback call delayed by given time in seconds.
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_callback(queue_free).set_delay(2) #this will call queue_free() after 2 seconds
[/codeblock]
*/
func (self Instance) SetDelay(delay Float.X) [1]gdclass.CallbackTweener {
	return [1]gdclass.CallbackTweener(class(self).SetDelay(gd.Float(delay)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CallbackTweener

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CallbackTweener"))
	casted := Instance{*(*gdclass.CallbackTweener)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Makes the callback call delayed by given time in seconds.
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_callback(queue_free).set_delay(2) #this will call queue_free() after 2 seconds
[/codeblock]
*/
//go:nosplit
func (self class) SetDelay(delay gd.Float) [1]gdclass.CallbackTweener {
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CallbackTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.CallbackTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.CallbackTweener](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsCallbackTweener() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCallbackTweener() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTweener() Tweener.Advanced    { return *((*Tweener.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTweener() Tweener.Instance {
	return *((*Tweener.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Tweener.Advanced(self.AsTweener()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Tweener.Instance(self.AsTweener()), name)
	}
}
func init() {
	gdclass.Register("CallbackTweener", func(ptr gd.Object) any {
		return [1]gdclass.CallbackTweener{*(*gdclass.CallbackTweener)(unsafe.Pointer(&ptr))}
	})
}
