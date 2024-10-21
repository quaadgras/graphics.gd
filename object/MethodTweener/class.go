package MethodTweener

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Tweener"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[MethodTweener] is similar to a combination of [CallbackTweener] and [PropertyTweener]. It calls a method providing an interpolated value as a parameter. See [method Tween.tween_method] for more usage information.
The tweener will finish automatically if the callback's target object is freed.
[b]Note:[/b] [method Tween.tween_method] is the only correct way to create [MethodTweener]. Any [MethodTweener] created manually will not function correctly.

*/
type Simple [1]classdb.MethodTweener
func (self Simple) SetDelay(delay float64) [1]classdb.MethodTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.MethodTweener(Expert(self).SetDelay(gc, gd.Float(delay)))
}
func (self Simple) SetTrans(trans classdb.TweenTransitionType) [1]classdb.MethodTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.MethodTweener(Expert(self).SetTrans(gc, trans))
}
func (self Simple) SetEase(ease classdb.TweenEaseType) [1]classdb.MethodTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.MethodTweener(Expert(self).SetEase(gc, ease))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MethodTweener
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the time in seconds after which the [MethodTweener] will start interpolating. By default there's no delay.
*/
//go:nosplit
func (self class) SetDelay(ctx gd.Lifetime, delay gd.Float) object.MethodTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MethodTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.MethodTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetTrans(ctx gd.Lifetime, trans classdb.TweenTransitionType) object.MethodTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, trans)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MethodTweener.Bind_set_trans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.MethodTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetEase(ctx gd.Lifetime, ease classdb.TweenEaseType) object.MethodTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ease)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MethodTweener.Bind_set_ease, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.MethodTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsMethodTweener() Expert { return self[0].AsMethodTweener() }


//go:nosplit
func (self Simple) AsMethodTweener() Simple { return self[0].AsMethodTweener() }


//go:nosplit
func (self class) AsTweener() Tweener.Expert { return self[0].AsTweener() }


//go:nosplit
func (self Simple) AsTweener() Tweener.Simple { return self[0].AsTweener() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("MethodTweener", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
