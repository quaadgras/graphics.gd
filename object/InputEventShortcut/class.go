package InputEventShortcut

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/InputEvent"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
InputEventShortcut is a special event that can be received in [method Node._input], [method Node._shortcut_input], and [method Node._unhandled_input]. It is typically sent by the editor's Command Palette to trigger actions, but can also be sent manually using [method Viewport.push_input].

*/
type Simple [1]classdb.InputEventShortcut
func (self Simple) SetShortcut(shortcut [1]classdb.Shortcut) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShortcut(shortcut)
}
func (self Simple) GetShortcut() [1]classdb.Shortcut {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Shortcut(Expert(self).GetShortcut(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.InputEventShortcut
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetShortcut(shortcut object.Shortcut)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventShortcut.Bind_set_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShortcut(ctx gd.Lifetime) object.Shortcut {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventShortcut.Bind_get_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Shortcut
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsInputEventShortcut() Expert { return self[0].AsInputEventShortcut() }


//go:nosplit
func (self Simple) AsInputEventShortcut() Simple { return self[0].AsInputEventShortcut() }


//go:nosplit
func (self class) AsInputEvent() InputEvent.Expert { return self[0].AsInputEvent() }


//go:nosplit
func (self Simple) AsInputEvent() InputEvent.Simple { return self[0].AsInputEvent() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


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
func init() {classdb.Register("InputEventShortcut", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
