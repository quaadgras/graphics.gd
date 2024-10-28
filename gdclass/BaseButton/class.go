package BaseButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[BaseButton] is an abstract base class for GUI buttons. It doesn't display anything by itself.
	// BaseButton methods that can be overridden by a [Class] that extends it.
	type BaseButton interface {
		//Called when the button is pressed. If you need to know the button's pressed state (and [member toggle_mode] is active), use [method _toggled] instead.
		Pressed() 
		//Called when the button is toggled (only if [member toggle_mode] is active).
		Toggled(toggled_on bool) 
	}

*/
type Go [1]classdb.BaseButton

/*
Called when the button is pressed. If you need to know the button's pressed state (and [member toggle_mode] is active), use [method _toggled] instead.
*/
func (Go) _pressed(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
	}
}

/*
Called when the button is toggled (only if [member toggle_mode] is active).
*/
func (Go) _toggled(impl func(ptr unsafe.Pointer, toggled_on bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var toggled_on = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, toggled_on)
	}
}

/*
Changes the [member button_pressed] state of the button, without emitting [signal toggled]. Use when you just want to change the state of the button without sending the pressed event (e.g. when initializing scene). Only works if [member toggle_mode] is [code]true[/code].
[b]Note:[/b] This method doesn't unpress other buttons in [member button_group].
*/
func (self Go) SetPressedNoSignal(pressed bool) {
	class(self).SetPressedNoSignal(pressed)
}

/*
Returns [code]true[/code] if the mouse has entered the button and has not left it yet.
*/
func (self Go) IsHovered() bool {
	return bool(class(self).IsHovered())
}

/*
Returns the visual state used to draw the button. This is useful mainly when implementing your own draw code by either overriding _draw() or connecting to "draw" signal. The visual state of the button is defined by the [enum DrawMode] enum.
*/
func (self Go) GetDrawMode() classdb.BaseButtonDrawMode {
	return classdb.BaseButtonDrawMode(class(self).GetDrawMode())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.BaseButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BaseButton"))
	return Go{classdb.BaseButton(object)}
}

func (self Go) Disabled() bool {
		return bool(class(self).IsDisabled())
}

func (self Go) SetDisabled(value bool) {
	class(self).SetDisabled(value)
}

func (self Go) ToggleMode() bool {
		return bool(class(self).IsToggleMode())
}

func (self Go) SetToggleMode(value bool) {
	class(self).SetToggleMode(value)
}

func (self Go) ButtonPressed() bool {
		return bool(class(self).IsPressed())
}

func (self Go) SetButtonPressed(value bool) {
	class(self).SetPressed(value)
}

func (self Go) ActionMode() classdb.BaseButtonActionMode {
		return classdb.BaseButtonActionMode(class(self).GetActionMode())
}

func (self Go) SetActionMode(value classdb.BaseButtonActionMode) {
	class(self).SetActionMode(value)
}

func (self Go) ButtonMask() gd.MouseButtonMask {
		return gd.MouseButtonMask(class(self).GetButtonMask())
}

func (self Go) SetButtonMask(value gd.MouseButtonMask) {
	class(self).SetButtonMask(value)
}

func (self Go) KeepPressedOutside() bool {
		return bool(class(self).IsKeepPressedOutside())
}

func (self Go) SetKeepPressedOutside(value bool) {
	class(self).SetKeepPressedOutside(value)
}

func (self Go) ButtonGroup() gdclass.ButtonGroup {
		return gdclass.ButtonGroup(class(self).GetButtonGroup())
}

func (self Go) SetButtonGroup(value gdclass.ButtonGroup) {
	class(self).SetButtonGroup(value)
}

func (self Go) Shortcut() gdclass.Shortcut {
		return gdclass.Shortcut(class(self).GetShortcut())
}

func (self Go) SetShortcut(value gdclass.Shortcut) {
	class(self).SetShortcut(value)
}

func (self Go) ShortcutFeedback() bool {
		return bool(class(self).IsShortcutFeedback())
}

func (self Go) SetShortcutFeedback(value bool) {
	class(self).SetShortcutFeedback(value)
}

func (self Go) ShortcutInTooltip() bool {
		return bool(class(self).IsShortcutInTooltipEnabled())
}

func (self Go) SetShortcutInTooltip(value bool) {
	class(self).SetShortcutInTooltip(value)
}

/*
Called when the button is pressed. If you need to know the button's pressed state (and [member toggle_mode] is active), use [method _toggled] instead.
*/
func (class) _pressed(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
	}
}

/*
Called when the button is toggled (only if [member toggle_mode] is active).
*/
func (class) _toggled(impl func(ptr unsafe.Pointer, toggled_on bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var toggled_on = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, toggled_on)
	}
}

//go:nosplit
func (self class) SetPressed(pressed bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPressed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_is_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Changes the [member button_pressed] state of the button, without emitting [signal toggled]. Use when you just want to change the state of the button without sending the pressed event (e.g. when initializing scene). Only works if [member toggle_mode] is [code]true[/code].
[b]Note:[/b] This method doesn't unpress other buttons in [member button_group].
*/
//go:nosplit
func (self class) SetPressedNoSignal(pressed bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_pressed_no_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the mouse has entered the button and has not left it yet.
*/
//go:nosplit
func (self class) IsHovered() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_is_hovered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetToggleMode(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsToggleMode() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_is_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcutInTooltip(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_shortcut_in_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShortcutInTooltipEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_is_shortcut_in_tooltip_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisabled(disabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_is_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetActionMode(mode classdb.BaseButtonActionMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_action_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetActionMode() classdb.BaseButtonActionMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseButtonActionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_get_action_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetButtonMask(mask gd.MouseButtonMask)  {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetButtonMask() gd.MouseButtonMask {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.MouseButtonMask](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_get_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the visual state used to draw the button. This is useful mainly when implementing your own draw code by either overriding _draw() or connecting to "draw" signal. The visual state of the button is defined by the [enum DrawMode] enum.
*/
//go:nosplit
func (self class) GetDrawMode() classdb.BaseButtonDrawMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseButtonDrawMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_get_draw_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepPressedOutside(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_keep_pressed_outside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsKeepPressedOutside() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_is_keep_pressed_outside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcutFeedback(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_shortcut_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShortcutFeedback() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_is_shortcut_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcut(shortcut gdclass.Shortcut)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(shortcut[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShortcut() gdclass.Shortcut {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_get_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Shortcut{classdb.Shortcut(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetButtonGroup(button_group gdclass.ButtonGroup)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(button_group[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_button_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetButtonGroup() gdclass.ButtonGroup {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_get_button_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.ButtonGroup{classdb.ButtonGroup(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self Go) OnPressed(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("pressed"), gd.NewCallable(cb), 0)
}


func (self Go) OnButtonUp(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("button_up"), gd.NewCallable(cb), 0)
}


func (self Go) OnButtonDown(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("button_down"), gd.NewCallable(cb), 0)
}


func (self Go) OnToggled(cb func(toggled_on bool)) {
	self[0].AsObject().Connect(gd.NewStringName("toggled"), gd.NewCallable(cb), 0)
}


func (self class) AsBaseButton() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsBaseButton() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_pressed": return reflect.ValueOf(self._pressed);
	case "_toggled": return reflect.ValueOf(self._toggled);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_pressed": return reflect.ValueOf(self._pressed);
	case "_toggled": return reflect.ValueOf(self._toggled);
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("BaseButton", func(ptr gd.Object) any { return classdb.BaseButton(ptr) })}
type DrawMode = classdb.BaseButtonDrawMode

const (
/*The normal state (i.e. not pressed, not hovered, not toggled and enabled) of buttons.*/
	DrawNormal DrawMode = 0
/*The state of buttons are pressed.*/
	DrawPressed DrawMode = 1
/*The state of buttons are hovered.*/
	DrawHover DrawMode = 2
/*The state of buttons are disabled.*/
	DrawDisabled DrawMode = 3
/*The state of buttons are both hovered and pressed.*/
	DrawHoverPressed DrawMode = 4
)
type ActionMode = classdb.BaseButtonActionMode

const (
/*Require just a press to consider the button clicked.*/
	ActionModeButtonPress ActionMode = 0
/*Require a press and a subsequent release before considering the button clicked.*/
	ActionModeButtonRelease ActionMode = 1
)
