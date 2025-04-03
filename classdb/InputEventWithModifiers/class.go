// Package InputEventWithModifiers provides methods for working with InputEventWithModifiers object instances.
package InputEventWithModifiers

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/InputEvent"
import "graphics.gd/classdb/InputEventFromWindow"
import "graphics.gd/classdb/Resource"
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
Stores information about mouse, keyboard, and touch gesture input events. This includes information about which modifier keys are pressed, such as [kbd]Shift[/kbd] or [kbd]Alt[/kbd]. See [method Node._input].
[b]Note:[/b] Modifier keys are considered modifiers only when used in combination with another key. As a result, their corresponding member variables, such as [member ctrl_pressed], will return [code]false[/code] if the key is pressed on its own.
*/
type Instance [1]gdclass.InputEventWithModifiers

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsInputEventWithModifiers() Instance
}

/*
On macOS, returns [code]true[/code] if [kbd]Meta[/kbd] ([kbd]Cmd[/kbd]) is pressed.
On other platforms, returns [code]true[/code] if [kbd]Ctrl[/kbd] is pressed.
*/
func (self Instance) IsCommandOrControlPressed() bool { //gd:InputEventWithModifiers.is_command_or_control_pressed
	return bool(Advanced(self).IsCommandOrControlPressed())
}

/*
Returns the keycode combination of modifier keys.
*/
func (self Instance) GetModifiersMask() KeyModifierMask { //gd:InputEventWithModifiers.get_modifiers_mask
	return KeyModifierMask(Advanced(self).GetModifiersMask())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEventWithModifiers

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventWithModifiers"))
	casted := Instance{*(*gdclass.InputEventWithModifiers)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) CommandOrControlAutoremap() bool {
	return bool(class(self).IsCommandOrControlAutoremap())
}

func (self Instance) SetCommandOrControlAutoremap(value bool) {
	class(self).SetCommandOrControlAutoremap(value)
}

func (self Instance) AltPressed() bool {
	return bool(class(self).IsAltPressed())
}

func (self Instance) SetAltPressed(value bool) {
	class(self).SetAltPressed(value)
}

func (self Instance) ShiftPressed() bool {
	return bool(class(self).IsShiftPressed())
}

func (self Instance) SetShiftPressed(value bool) {
	class(self).SetShiftPressed(value)
}

func (self Instance) CtrlPressed() bool {
	return bool(class(self).IsCtrlPressed())
}

func (self Instance) SetCtrlPressed(value bool) {
	class(self).SetCtrlPressed(value)
}

func (self Instance) MetaPressed() bool {
	return bool(class(self).IsMetaPressed())
}

func (self Instance) SetMetaPressed(value bool) {
	class(self).SetMetaPressed(value)
}

//go:nosplit
func (self class) SetCommandOrControlAutoremap(enable bool) { //gd:InputEventWithModifiers.set_command_or_control_autoremap
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_command_or_control_autoremap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCommandOrControlAutoremap() bool { //gd:InputEventWithModifiers.is_command_or_control_autoremap
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_command_or_control_autoremap, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
On macOS, returns [code]true[/code] if [kbd]Meta[/kbd] ([kbd]Cmd[/kbd]) is pressed.
On other platforms, returns [code]true[/code] if [kbd]Ctrl[/kbd] is pressed.
*/
//go:nosplit
func (self class) IsCommandOrControlPressed() bool { //gd:InputEventWithModifiers.is_command_or_control_pressed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_command_or_control_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAltPressed(pressed bool) { //gd:InputEventWithModifiers.set_alt_pressed
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_alt_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAltPressed() bool { //gd:InputEventWithModifiers.is_alt_pressed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_alt_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShiftPressed(pressed bool) { //gd:InputEventWithModifiers.set_shift_pressed
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_shift_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShiftPressed() bool { //gd:InputEventWithModifiers.is_shift_pressed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_shift_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCtrlPressed(pressed bool) { //gd:InputEventWithModifiers.set_ctrl_pressed
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_ctrl_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCtrlPressed() bool { //gd:InputEventWithModifiers.is_ctrl_pressed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_ctrl_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMetaPressed(pressed bool) { //gd:InputEventWithModifiers.set_meta_pressed
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_set_meta_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMetaPressed() bool { //gd:InputEventWithModifiers.is_meta_pressed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_is_meta_pressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the keycode combination of modifier keys.
*/
//go:nosplit
func (self class) GetModifiersMask() KeyModifierMask { //gd:InputEventWithModifiers.get_modifiers_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[KeyModifierMask](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventWithModifiers.Bind_get_modifiers_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventWithModifiers() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventWithModifiers() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEventFromWindow() InputEventFromWindow.Advanced {
	return *((*InputEventFromWindow.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventFromWindow() InputEventFromWindow.Instance {
	return *((*InputEventFromWindow.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEvent() InputEvent.Advanced {
	return *((*InputEvent.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEvent() InputEvent.Instance {
	return *((*InputEvent.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(InputEventFromWindow.Advanced(self.AsInputEventFromWindow()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEventFromWindow.Instance(self.AsInputEventFromWindow()), name)
	}
}
func init() {
	gdclass.Register("InputEventWithModifiers", func(ptr gd.Object) any {
		return [1]gdclass.InputEventWithModifiers{*(*gdclass.InputEventWithModifiers)(unsafe.Pointer(&ptr))}
	})
}

type KeyModifierMask int

const (
	/*Key Code mask.*/
	KeyCodeMask KeyModifierMask = 8388607
	/*Modifier key mask.*/
	KeyModifierMaskDefault KeyModifierMask = 2130706432
	/*Automatically remapped to [constant KEY_META] on macOS and [constant KEY_CTRL] on other platforms, this mask is never set in the actual events, and should be used for key mapping only.*/
	KeyMaskCmdOrCtrl KeyModifierMask = 16777216
	/*Shift key mask.*/
	KeyMaskShift KeyModifierMask = 33554432
	/*Alt or Option (on macOS) key mask.*/
	KeyMaskAlt KeyModifierMask = 67108864
	/*Command (on macOS) or Meta/Windows key mask.*/
	KeyMaskMeta KeyModifierMask = 134217728
	/*Control key mask.*/
	KeyMaskCtrl KeyModifierMask = 268435456
	/*Keypad key mask.*/
	KeyMaskKpad KeyModifierMask = 536870912
	/*Group Switch key mask.*/
	KeyMaskGroupSwitch KeyModifierMask = 1073741824
)
