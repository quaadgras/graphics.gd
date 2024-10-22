package Shortcut

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Shortcuts are commonly used for interacting with a [Control] element from an [InputEvent] (also known as hotkeys).
One shortcut can contain multiple [InputEvent]'s, allowing the possibility of triggering one action with multiple different inputs.

*/
type Go [1]classdb.Shortcut

/*
Returns whether [member events] contains an [InputEvent] which is valid.
*/
func (self Go) HasValidEvent() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasValidEvent())
}

/*
Returns whether any [InputEvent] in [member events] equals [param event].
*/
func (self Go) MatchesEvent(event gdclass.InputEvent) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).MatchesEvent(event))
}

/*
Returns the shortcut's first valid [InputEvent] as a [String].
*/
func (self Go) GetAsText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetAsText(gc).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Shortcut
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Shortcut"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Events() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetEvents(gc))
}

func (self Go) SetEvents(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEvents(value)
}

//go:nosplit
func (self class) SetEvents(events gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(events))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shortcut.Bind_set_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEvents(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shortcut.Bind_get_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [member events] contains an [InputEvent] which is valid.
*/
//go:nosplit
func (self class) HasValidEvent() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shortcut.Bind_has_valid_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether any [InputEvent] in [member events] equals [param event].
*/
//go:nosplit
func (self class) MatchesEvent(event gdclass.InputEvent) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shortcut.Bind_matches_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the shortcut's first valid [InputEvent] as a [String].
*/
//go:nosplit
func (self class) GetAsText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Shortcut.Bind_get_as_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsShortcut() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsShortcut() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Shortcut", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
