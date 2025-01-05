// Package MethodTweener provides methods for working with MethodTweener object instances.
package MethodTweener

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

/*
[MethodTweener] is similar to a combination of [CallbackTweener] and [PropertyTweener]. It calls a method providing an interpolated value as a parameter. See [method Tween.tween_method] for more usage information.
The tweener will finish automatically if the callback's target object is freed.
[b]Note:[/b] [method Tween.tween_method] is the only correct way to create [MethodTweener]. Any [MethodTweener] created manually will not function correctly.
*/
type Instance [1]gdclass.MethodTweener
type Any interface {
	gd.IsClass
	AsMethodTweener() Instance
}

/*
Sets the time in seconds after which the [MethodTweener] will start interpolating. By default there's no delay.
*/
func (self Instance) SetDelay(delay Float.X) [1]gdclass.MethodTweener {
	return [1]gdclass.MethodTweener(class(self).SetDelay(gd.Float(delay)))
}

/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
func (self Instance) SetTrans(trans gdclass.TweenTransitionType) [1]gdclass.MethodTweener {
	return [1]gdclass.MethodTweener(class(self).SetTrans(trans))
}

/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
func (self Instance) SetEase(ease gdclass.TweenEaseType) [1]gdclass.MethodTweener {
	return [1]gdclass.MethodTweener(class(self).SetEase(ease))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MethodTweener

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MethodTweener"))
	return Instance{*(*gdclass.MethodTweener)(unsafe.Pointer(&object))}
}

/*
Sets the time in seconds after which the [MethodTweener] will start interpolating. By default there's no delay.
*/
//go:nosplit
func (self class) SetDelay(delay gd.Float) [1]gdclass.MethodTweener {
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MethodTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.MethodTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.MethodTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetTrans(trans gdclass.TweenTransitionType) [1]gdclass.MethodTweener {
	var frame = callframe.New()
	callframe.Arg(frame, trans)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MethodTweener.Bind_set_trans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.MethodTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.MethodTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetEase(ease gdclass.TweenEaseType) [1]gdclass.MethodTweener {
	var frame = callframe.New()
	callframe.Arg(frame, ease)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MethodTweener.Bind_set_ease, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.MethodTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.MethodTweener](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsMethodTweener() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMethodTweener() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTweener() Tweener.Advanced  { return *((*Tweener.Advanced)(unsafe.Pointer(&self))) }
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
	gdclass.Register("MethodTweener", func(ptr gd.Object) any {
		return [1]gdclass.MethodTweener{*(*gdclass.MethodTweener)(unsafe.Pointer(&ptr))}
	})
}
