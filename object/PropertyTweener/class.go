package PropertyTweener

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
[PropertyTweener] is used to interpolate a property in an object. See [method Tween.tween_property] for more usage information.
[b]Note:[/b] [method Tween.tween_property] is the only correct way to create [PropertyTweener]. Any [PropertyTweener] created manually will not function correctly.

*/
type Simple [1]classdb.PropertyTweener
func (self Simple) From(value gd.Variant) [1]classdb.PropertyTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PropertyTweener(Expert(self).From(gc, value))
}
func (self Simple) FromCurrent() [1]classdb.PropertyTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PropertyTweener(Expert(self).FromCurrent(gc))
}
func (self Simple) AsRelative() [1]classdb.PropertyTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PropertyTweener(Expert(self).AsRelative(gc))
}
func (self Simple) SetTrans(trans classdb.TweenTransitionType) [1]classdb.PropertyTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PropertyTweener(Expert(self).SetTrans(gc, trans))
}
func (self Simple) SetEase(ease classdb.TweenEaseType) [1]classdb.PropertyTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PropertyTweener(Expert(self).SetEase(gc, ease))
}
func (self Simple) SetCustomInterpolator(interpolator_method gd.Callable) [1]classdb.PropertyTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PropertyTweener(Expert(self).SetCustomInterpolator(gc, interpolator_method))
}
func (self Simple) SetDelay(delay float64) [1]classdb.PropertyTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PropertyTweener(Expert(self).SetDelay(gc, gd.Float(delay)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PropertyTweener
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets a custom initial value to the [PropertyTweener].
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_property(self, "position", Vector2(200, 100), 1).from(Vector2(100, 100)) #this will move the node from position (100, 100) to (200, 100)
[/codeblock]
*/
//go:nosplit
func (self class) From(ctx gd.Lifetime, value gd.Variant) object.PropertyTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(value))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PropertyTweener.Bind_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PropertyTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Makes the [PropertyTweener] use the current property value (i.e. at the time of creating this [PropertyTweener]) as a starting point. This is equivalent of using [method from] with the current value. These two calls will do the same:
[codeblock]
tween.tween_property(self, "position", Vector2(200, 100), 1).from(position)
tween.tween_property(self, "position", Vector2(200, 100), 1).from_current()
[/codeblock]
*/
//go:nosplit
func (self class) FromCurrent(ctx gd.Lifetime) object.PropertyTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PropertyTweener.Bind_from_current, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PropertyTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
When called, the final value will be used as a relative value instead.
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_property(self, "position", Vector2.RIGHT * 100, 1).as_relative() #the node will move by 100 pixels to the right
[/codeblock]
*/
//go:nosplit
func (self class) AsRelative(ctx gd.Lifetime) object.PropertyTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PropertyTweener.Bind_as_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PropertyTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetTrans(ctx gd.Lifetime, trans classdb.TweenTransitionType) object.PropertyTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, trans)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PropertyTweener.Bind_set_trans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PropertyTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetEase(ctx gd.Lifetime, ease classdb.TweenEaseType) object.PropertyTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ease)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PropertyTweener.Bind_set_ease, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PropertyTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Allows interpolating the value with a custom easing function. The provided [param interpolator_method] will be called with a value ranging from [code]0.0[/code] to [code]1.0[/code] and is expected to return a value within the same range (values outside the range can be used for overshoot). The return value of the method is then used for interpolation between initial and final value. Note that the parameter passed to the method is still subject to the tweener's own easing.
[b]Example:[/b]
[codeblock]
@export var curve: Curve

func _ready():
    var tween = create_tween()
    # Interpolate the value using a custom curve.
    tween.tween_property(self, "position:x", 300, 1).as_relative().set_custom_interpolator(tween_curve)

func tween_curve(v):
    return curve.sample_baked(v)
[/codeblock]
*/
//go:nosplit
func (self class) SetCustomInterpolator(ctx gd.Lifetime, interpolator_method gd.Callable) object.PropertyTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interpolator_method))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PropertyTweener.Bind_set_custom_interpolator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PropertyTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the time in seconds after which the [PropertyTweener] will start interpolating. By default there's no delay.
*/
//go:nosplit
func (self class) SetDelay(ctx gd.Lifetime, delay gd.Float) object.PropertyTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PropertyTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PropertyTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPropertyTweener() Expert { return self[0].AsPropertyTweener() }


//go:nosplit
func (self Simple) AsPropertyTweener() Simple { return self[0].AsPropertyTweener() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PropertyTweener", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
