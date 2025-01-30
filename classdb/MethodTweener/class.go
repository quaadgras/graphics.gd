// Package MethodTweener provides methods for working with MethodTweener object instances.
package MethodTweener

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Tweener"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[MethodTweener] is similar to a combination of [CallbackTweener] and [PropertyTweener]. It calls a method providing an interpolated value as a parameter. See [method Tween.tween_method] for more usage information.
The tweener will finish automatically if the callback's target object is freed.
[b]Note:[/b] [method Tween.tween_method] is the only correct way to create [MethodTweener]. Any [MethodTweener] created manually will not function correctly.
*/
type Instance [1]gdclass.MethodTweener

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMethodTweener() Instance
}

/*
Sets the time in seconds after which the [MethodTweener] will start interpolating. By default there's no delay.
*/
func (self Instance) SetDelay(delay Float.X) [1]gdclass.MethodTweener { //gd:MethodTweener.set_delay
	return [1]gdclass.MethodTweener(class(self).SetDelay(float64(delay)))
}

/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
func (self Instance) SetTrans(trans gdclass.TweenTransitionType) [1]gdclass.MethodTweener { //gd:MethodTweener.set_trans
	return [1]gdclass.MethodTweener(class(self).SetTrans(trans))
}

/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
func (self Instance) SetEase(ease gdclass.TweenEaseType) [1]gdclass.MethodTweener { //gd:MethodTweener.set_ease
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
	casted := Instance{*(*gdclass.MethodTweener)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Sets the time in seconds after which the [MethodTweener] will start interpolating. By default there's no delay.
*/
//go:nosplit
func (self class) SetDelay(delay float64) [1]gdclass.MethodTweener { //gd:MethodTweener.set_delay
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MethodTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MethodTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.MethodTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetTrans(trans gdclass.TweenTransitionType) [1]gdclass.MethodTweener { //gd:MethodTweener.set_trans
	var frame = callframe.New()
	callframe.Arg(frame, trans)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MethodTweener.Bind_set_trans, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MethodTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.MethodTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetEase(ease gdclass.TweenEaseType) [1]gdclass.MethodTweener { //gd:MethodTweener.set_ease
	var frame = callframe.New()
	callframe.Arg(frame, ease)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MethodTweener.Bind_set_ease, self.AsObject(), frame.Array(0), r_ret.Addr())
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
