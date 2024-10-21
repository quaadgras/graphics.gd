package InputEventWithModifiers

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/InputEventFromWindow"
import "grow.graphics/gd/object/InputEvent"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Stores information about mouse, keyboard, and touch gesture input events. This includes information about which modifier keys are pressed, such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See [method Node._input].

*/
type Simple [1]classdb.InputEventWithModifiers
func (self Simple) SetCommandOrControlAutoremap(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCommandOrControlAutoremap(enable)
}
func (self Simple) IsCommandOrControlAutoremap() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCommandOrControlAutoremap())
}
func (self Simple) IsCommandOrControlPressed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCommandOrControlPressed())
}
func (self Simple) SetAltPressed(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAltPressed(pressed)
}
func (self Simple) IsAltPressed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAltPressed())
}
func (self Simple) SetShiftPressed(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShiftPressed(pressed)
}
func (self Simple) IsShiftPressed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShiftPressed())
}
func (self Simple) SetCtrlPressed(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCtrlPressed(pressed)
}
func (self Simple) IsCtrlPressed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCtrlPressed())
}
func (self Simple) SetMetaPressed(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMetaPressed(pressed)
}
func (self Simple) IsMetaPressed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMetaPressed())
}
func (self Simple) GetModifiersMask() gd.KeyModifierMask {
	gc := gd.GarbageCollector(); _ = gc
	return gd.KeyModifierMask(Expert(self).GetModifiersMask())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.InputEventWithModifiers
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetCommandOrControlAutoremap(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_set_command_or_control_autoremap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCommandOrControlAutoremap() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_is_command_or_control_autoremap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
On macOS, returns [code]true[/code] if [kbd]Meta[/kbd] ([kbd]Cmd[/kbd]) is pressed.
On other platforms, returns [code]true[/code] if [kbd]Ctrl[/kbd] is pressed.
*/
//go:nosplit
func (self class) IsCommandOrControlPressed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_is_command_or_control_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAltPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_set_alt_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAltPressed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_is_alt_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShiftPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_set_shift_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShiftPressed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_is_shift_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCtrlPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_set_ctrl_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCtrlPressed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_is_ctrl_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMetaPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_set_meta_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMetaPressed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_is_meta_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the keycode combination of modifier keys.
*/
//go:nosplit
func (self class) GetModifiersMask() gd.KeyModifierMask {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.KeyModifierMask](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.InputEventWithModifiers.Bind_get_modifiers_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsInputEventWithModifiers() Expert { return self[0].AsInputEventWithModifiers() }


//go:nosplit
func (self Simple) AsInputEventWithModifiers() Simple { return self[0].AsInputEventWithModifiers() }


//go:nosplit
func (self class) AsInputEventFromWindow() InputEventFromWindow.Expert { return self[0].AsInputEventFromWindow() }


//go:nosplit
func (self Simple) AsInputEventFromWindow() InputEventFromWindow.Simple { return self[0].AsInputEventFromWindow() }


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
func init() {classdb.Register("InputEventWithModifiers", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
