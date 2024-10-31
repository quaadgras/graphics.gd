package BaseButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

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
type Instance [1]classdb.BaseButton

/*
Called when the button is pressed. If you need to know the button's pressed state (and [member toggle_mode] is active), use [method _toggled] instead.
*/
func (Instance) _pressed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the button is toggled (only if [member toggle_mode] is active).
*/
func (Instance) _toggled(impl func(ptr unsafe.Pointer, toggled_on bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var toggled_on = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, toggled_on)
	}
}

/*
Changes the [member button_pressed] state of the button, without emitting [signal toggled]. Use when you just want to change the state of the button without sending the pressed event (e.g. when initializing scene). Only works if [member toggle_mode] is [code]true[/code].
[b]Note:[/b] This method doesn't unpress other buttons in [member button_group].
*/
func (self Instance) SetPressedNoSignal(pressed bool) {
	class(self).SetPressedNoSignal(pressed)
}

/*
Returns [code]true[/code] if the mouse has entered the button and has not left it yet.
*/
func (self Instance) IsHovered() bool {
	return bool(class(self).IsHovered())
}

/*
Returns the visual state used to draw the button. This is useful mainly when implementing your own draw code by either overriding _draw() or connecting to "draw" signal. The visual state of the button is defined by the [enum DrawMode] enum.
*/
func (self Instance) GetDrawMode() classdb.BaseButtonDrawMode {
	return classdb.BaseButtonDrawMode(class(self).GetDrawMode())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.BaseButton

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BaseButton"))
	return Instance{classdb.BaseButton(object)}
}

func (self Instance) Disabled() bool {
	return bool(class(self).IsDisabled())
}

func (self Instance) SetDisabled(value bool) {
	class(self).SetDisabled(value)
}

func (self Instance) ToggleMode() bool {
	return bool(class(self).IsToggleMode())
}

func (self Instance) SetToggleMode(value bool) {
	class(self).SetToggleMode(value)
}

func (self Instance) ButtonPressed() bool {
	return bool(class(self).IsPressed())
}

func (self Instance) SetButtonPressed(value bool) {
	class(self).SetPressed(value)
}

func (self Instance) ActionMode() classdb.BaseButtonActionMode {
	return classdb.BaseButtonActionMode(class(self).GetActionMode())
}

func (self Instance) SetActionMode(value classdb.BaseButtonActionMode) {
	class(self).SetActionMode(value)
}

func (self Instance) ButtonMask() gdconst.MouseButtonMask {
	return gdconst.MouseButtonMask(class(self).GetButtonMask())
}

func (self Instance) SetButtonMask(value gdconst.MouseButtonMask) {
	class(self).SetButtonMask(value)
}

func (self Instance) KeepPressedOutside() bool {
	return bool(class(self).IsKeepPressedOutside())
}

func (self Instance) SetKeepPressedOutside(value bool) {
	class(self).SetKeepPressedOutside(value)
}

func (self Instance) ButtonGroup() gdclass.ButtonGroup {
	return gdclass.ButtonGroup(class(self).GetButtonGroup())
}

func (self Instance) SetButtonGroup(value gdclass.ButtonGroup) {
	class(self).SetButtonGroup(value)
}

func (self Instance) Shortcut() gdclass.Shortcut {
	return gdclass.Shortcut(class(self).GetShortcut())
}

func (self Instance) SetShortcut(value gdclass.Shortcut) {
	class(self).SetShortcut(value)
}

func (self Instance) ShortcutFeedback() bool {
	return bool(class(self).IsShortcutFeedback())
}

func (self Instance) SetShortcutFeedback(value bool) {
	class(self).SetShortcutFeedback(value)
}

func (self Instance) ShortcutInTooltip() bool {
	return bool(class(self).IsShortcutInTooltipEnabled())
}

func (self Instance) SetShortcutInTooltip(value bool) {
	class(self).SetShortcutInTooltip(value)
}

/*
Called when the button is pressed. If you need to know the button's pressed state (and [member toggle_mode] is active), use [method _toggled] instead.
*/
func (class) _pressed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the button is toggled (only if [member toggle_mode] is active).
*/
func (class) _toggled(impl func(ptr unsafe.Pointer, toggled_on bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var toggled_on = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, toggled_on)
	}
}

//go:nosplit
func (self class) SetPressed(pressed bool) {
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
func (self class) SetPressedNoSignal(pressed bool) {
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
func (self class) SetToggleMode(enabled bool) {
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
func (self class) SetShortcutInTooltip(enabled bool) {
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
func (self class) SetDisabled(disabled bool) {
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
func (self class) SetActionMode(mode classdb.BaseButtonActionMode) {
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
func (self class) SetButtonMask(mask gdconst.MouseButtonMask) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BaseButton.Bind_set_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetButtonMask() gdconst.MouseButtonMask {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdconst.MouseButtonMask](frame)
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
func (self class) SetKeepPressedOutside(enabled bool) {
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
func (self class) SetShortcutFeedback(enabled bool) {
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
func (self class) SetShortcut(shortcut gdclass.Shortcut) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
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
func (self class) SetButtonGroup(button_group gdclass.ButtonGroup) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(button_group[0])[0])
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
func (self Instance) OnPressed(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnButtonUp(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("button_up"), gd.NewCallable(cb), 0)
}

func (self Instance) OnButtonDown(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("button_down"), gd.NewCallable(cb), 0)
}

func (self Instance) OnToggled(cb func(toggled_on bool)) {
	self[0].AsObject().Connect(gd.NewStringName("toggled"), gd.NewCallable(cb), 0)
}

func (self class) AsBaseButton() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBaseButton() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_pressed":
		return reflect.ValueOf(self._pressed)
	case "_toggled":
		return reflect.ValueOf(self._toggled)
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_pressed":
		return reflect.ValueOf(self._pressed)
	case "_toggled":
		return reflect.ValueOf(self._toggled)
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {
	classdb.Register("BaseButton", func(ptr gd.Object) any { return classdb.BaseButton(ptr) })
}

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
