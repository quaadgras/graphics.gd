package BaseButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

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
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Called when the button is toggled (only if [member toggle_mode] is active).
*/
func (Go) _toggled(impl func(ptr unsafe.Pointer, toggled_on bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var toggled_on = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, toggled_on)
		gc.End()
	}
}

/*
Changes the [member button_pressed] state of the button, without emitting [signal toggled]. Use when you just want to change the state of the button without sending the pressed event (e.g. when initializing scene). Only works if [member toggle_mode] is [code]true[/code].
[b]Note:[/b] This method doesn't unpress other buttons in [member button_group].
*/
func (self Go) SetPressedNoSignal(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPressedNoSignal(pressed)
}

/*
Returns [code]true[/code] if the mouse has entered the button and has not left it yet.
*/
func (self Go) IsHovered() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsHovered())
}

/*
Returns the visual state used to draw the button. This is useful mainly when implementing your own draw code by either overriding _draw() or connecting to "draw" signal. The visual state of the button is defined by the [enum DrawMode] enum.
*/
func (self Go) GetDrawMode() classdb.BaseButtonDrawMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseButtonDrawMode(class(self).GetDrawMode())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.BaseButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("BaseButton"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Disabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDisabled())
}

func (self Go) SetDisabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisabled(value)
}

func (self Go) ToggleMode() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsToggleMode())
}

func (self Go) SetToggleMode(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetToggleMode(value)
}

func (self Go) ButtonPressed() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPressed())
}

func (self Go) SetButtonPressed(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPressed(value)
}

func (self Go) ActionMode() classdb.BaseButtonActionMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.BaseButtonActionMode(class(self).GetActionMode())
}

func (self Go) SetActionMode(value classdb.BaseButtonActionMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetActionMode(value)
}

func (self Go) ButtonMask() gd.MouseButtonMask {
	gc := gd.GarbageCollector(); _ = gc
		return gd.MouseButtonMask(class(self).GetButtonMask())
}

func (self Go) SetButtonMask(value gd.MouseButtonMask) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetButtonMask(value)
}

func (self Go) KeepPressedOutside() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsKeepPressedOutside())
}

func (self Go) SetKeepPressedOutside(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetKeepPressedOutside(value)
}

func (self Go) ButtonGroup() gdclass.ButtonGroup {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.ButtonGroup(class(self).GetButtonGroup(gc))
}

func (self Go) SetButtonGroup(value gdclass.ButtonGroup) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetButtonGroup(value)
}

func (self Go) Shortcut() gdclass.Shortcut {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Shortcut(class(self).GetShortcut(gc))
}

func (self Go) SetShortcut(value gdclass.Shortcut) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShortcut(value)
}

func (self Go) ShortcutFeedback() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShortcutFeedback())
}

func (self Go) SetShortcutFeedback(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShortcutFeedback(value)
}

func (self Go) ShortcutInTooltip() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShortcutInTooltipEnabled())
}

func (self Go) SetShortcutInTooltip(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShortcutInTooltip(value)
}

/*
Called when the button is pressed. If you need to know the button's pressed state (and [member toggle_mode] is active), use [method _toggled] instead.
*/
func (class) _pressed(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Called when the button is toggled (only if [member toggle_mode] is active).
*/
func (class) _toggled(impl func(ptr unsafe.Pointer, toggled_on bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var toggled_on = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, toggled_on)
		ctx.End()
	}
}

//go:nosplit
func (self class) SetPressed(pressed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPressed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_is_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_pressed_no_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the mouse has entered the button and has not left it yet.
*/
//go:nosplit
func (self class) IsHovered() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_is_hovered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetToggleMode(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsToggleMode() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_is_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcutInTooltip(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_shortcut_in_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShortcutInTooltipEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_is_shortcut_in_tooltip_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisabled(disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_is_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetActionMode(mode classdb.BaseButtonActionMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_action_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetActionMode() classdb.BaseButtonActionMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseButtonActionMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_get_action_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetButtonMask(mask gd.MouseButtonMask)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetButtonMask() gd.MouseButtonMask {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.MouseButtonMask](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_get_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the visual state used to draw the button. This is useful mainly when implementing your own draw code by either overriding _draw() or connecting to "draw" signal. The visual state of the button is defined by the [enum DrawMode] enum.
*/
//go:nosplit
func (self class) GetDrawMode() classdb.BaseButtonDrawMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.BaseButtonDrawMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_get_draw_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepPressedOutside(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_keep_pressed_outside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsKeepPressedOutside() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_is_keep_pressed_outside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcutFeedback(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_shortcut_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShortcutFeedback() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_is_shortcut_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcut(shortcut gdclass.Shortcut)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShortcut(ctx gd.Lifetime) gdclass.Shortcut {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_get_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Shortcut
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetButtonGroup(button_group gdclass.ButtonGroup)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(button_group[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_button_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetButtonGroup(ctx gd.Lifetime) gdclass.ButtonGroup {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_get_button_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ButtonGroup
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
func (self Go) OnPressed(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("pressed"), gc.Callable(cb), 0)
}


func (self Go) OnButtonUp(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("button_up"), gc.Callable(cb), 0)
}


func (self Go) OnButtonDown(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("button_down"), gc.Callable(cb), 0)
}


func (self Go) OnToggled(cb func(toggled_on bool)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("toggled"), gc.Callable(cb), 0)
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
func init() {classdb.Register("BaseButton", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
