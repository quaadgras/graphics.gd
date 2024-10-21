package CallbackTweener

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
[CallbackTweener] is used to call a method in a tweening sequence. See [method Tween.tween_callback] for more usage information.
The tweener will finish automatically if the callback's target object is freed.
[b]Note:[/b] [method Tween.tween_callback] is the only correct way to create [CallbackTweener]. Any [CallbackTweener] created manually will not function correctly.

*/
type Simple [1]classdb.CallbackTweener
func (self Simple) SetDelay(delay float64) [1]classdb.CallbackTweener {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.CallbackTweener(Expert(self).SetDelay(gc, gd.Float(delay)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CallbackTweener
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Makes the callback call delayed by given time in seconds.
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_callback(queue_free).set_delay(2) #this will call queue_free() after 2 seconds
[/codeblock]
*/
//go:nosplit
func (self class) SetDelay(ctx gd.Lifetime, delay gd.Float) object.CallbackTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CallbackTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.CallbackTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCallbackTweener() Expert { return self[0].AsCallbackTweener() }


//go:nosplit
func (self Simple) AsCallbackTweener() Simple { return self[0].AsCallbackTweener() }


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
func init() {classdb.Register("CallbackTweener", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
