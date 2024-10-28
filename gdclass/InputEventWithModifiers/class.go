package InputEventWithModifiers

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEventFromWindow"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Stores information about mouse, keyboard, and touch gesture input events. This includes information about which modifier keys are pressed, such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See [method Node._input].

*/
type Go [1]classdb.InputEventWithModifiers

/*
On macOS, returns [code]true[/code] if [kbd]Meta[/kbd] ([kbd]Cmd[/kbd]) is pressed.
On other platforms, returns [code]true[/code] if [kbd]Ctrl[/kbd] is pressed.
*/
func (self Go) IsCommandOrControlPressed() bool {
	return bool(class(self).IsCommandOrControlPressed())
}

/*
Returns the keycode combination of modifier keys.
*/
func (self Go) GetModifiersMask() gd.KeyModifierMask {
	return gd.KeyModifierMask(class(self).GetModifiersMask())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEventWithModifiers
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventWithModifiers"))
	return Go{classdb.InputEventWithModifiers(object)}
}

func (self Go) CommandOrControlAutoremap() bool {
		return bool(class(self).IsCommandOrControlAutoremap())
}

func (self Go) SetCommandOrControlAutoremap(value bool) {
	class(self).SetCommandOrControlAutoremap(value)
}

func (self Go) AltPressed() bool {
		return bool(class(self).IsAltPressed())
}

func (self Go) SetAltPressed(value bool) {
	class(self).SetAltPressed(value)
}

func (self Go) ShiftPressed() bool {
		return bool(class(self).IsShiftPressed())
}

func (self Go) SetShiftPressed(value bool) {
	class(self).SetShiftPressed(value)
}

func (self Go) CtrlPressed() bool {
		return bool(class(self).IsCtrlPressed())
}

func (self Go) SetCtrlPressed(value bool) {
	class(self).SetCtrlPressed(value)
}

func (self Go) MetaPressed() bool {
		return bool(class(self).IsMetaPressed())
}

func (self Go) SetMetaPressed(value bool) {
	class(self).SetMetaPressed(value)
}

//go:nosplit
func (self class) SetCommandOrControlAutoremap(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_command_or_control_autoremap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCommandOrControlAutoremap() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_command_or_control_autoremap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_command_or_control_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAltPressed(pressed bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_alt_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAltPressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_alt_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShiftPressed(pressed bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_shift_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShiftPressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_shift_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCtrlPressed(pressed bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_ctrl_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCtrlPressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_ctrl_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMetaPressed(pressed bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_meta_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMetaPressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_meta_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the keycode combination of modifier keys.
*/
//go:nosplit
func (self class) GetModifiersMask() gd.KeyModifierMask {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.KeyModifierMask](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_get_modifiers_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventWithModifiers() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventWithModifiers() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEventFromWindow() InputEventFromWindow.GD { return *((*InputEventFromWindow.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventFromWindow() InputEventFromWindow.Go { return *((*InputEventFromWindow.Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEvent() InputEvent.GD { return *((*InputEvent.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEvent() InputEvent.Go { return *((*InputEvent.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEventFromWindow(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEventFromWindow(), name)
	}
}
func init() {classdb.Register("InputEventWithModifiers", func(ptr gd.Object) any { return classdb.InputEventWithModifiers(ptr) })}
