package InputEventShortcut

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
InputEventShortcut is a special event that can be received in [method Node._input], [method Node._shortcut_input], and [method Node._unhandled_input]. It is typically sent by the editor's Command Palette to trigger actions, but can also be sent manually using [method Viewport.push_input].

*/
type Go [1]classdb.InputEventShortcut
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEventShortcut
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("InputEventShortcut"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Shortcut() gdclass.Shortcut {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Shortcut(class(self).GetShortcut(gc))
}

func (self Go) SetShortcut(value gdclass.Shortcut) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShortcut(value)
}

//go:nosplit
func (self class) SetShortcut(shortcut gdclass.Shortcut)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventShortcut.Bind_set_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShortcut(ctx gd.Lifetime) gdclass.Shortcut {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventShortcut.Bind_get_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Shortcut
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsInputEventShortcut() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventShortcut() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEvent() InputEvent.GD { return *((*InputEvent.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEvent() InputEvent.Go { return *((*InputEvent.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEvent(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEvent(), name)
	}
}
func init() {classdb.Register("InputEventShortcut", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
