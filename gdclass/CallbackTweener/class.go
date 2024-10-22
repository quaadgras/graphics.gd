package CallbackTweener

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Tweener"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[CallbackTweener] is used to call a method in a tweening sequence. See [method Tween.tween_callback] for more usage information.
The tweener will finish automatically if the callback's target object is freed.
[b]Note:[/b] [method Tween.tween_callback] is the only correct way to create [CallbackTweener]. Any [CallbackTweener] created manually will not function correctly.

*/
type Go [1]classdb.CallbackTweener

/*
Makes the callback call delayed by given time in seconds.
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_callback(queue_free).set_delay(2) #this will call queue_free() after 2 seconds
[/codeblock]
*/
func (self Go) SetDelay(delay float64) gdclass.CallbackTweener {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.CallbackTweener(class(self).SetDelay(gc, gd.Float(delay)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CallbackTweener
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CallbackTweener"))
	return *(*Go)(unsafe.Pointer(&object))
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
func (self class) SetDelay(ctx gd.Lifetime, delay gd.Float) gdclass.CallbackTweener {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CallbackTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.CallbackTweener
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsCallbackTweener() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCallbackTweener() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTweener() Tweener.GD { return *((*Tweener.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTweener() Tweener.Go { return *((*Tweener.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTweener(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTweener(), name)
	}
}
func init() {classdb.Register("CallbackTweener", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
