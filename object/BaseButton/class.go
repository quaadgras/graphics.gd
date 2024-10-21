package BaseButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
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
type Simple [1]classdb.BaseButton
func (Simple) _pressed(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _toggled(impl func(ptr unsafe.Pointer, toggled_on bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var toggled_on = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, toggled_on)
		gc.End()
	}
}
func (self Simple) SetPressed(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPressed(pressed)
}
func (self Simple) IsPressed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPressed())
}
func (self Simple) SetPressedNoSignal(pressed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPressedNoSignal(pressed)
}
func (self Simple) IsHovered() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHovered())
}
func (self Simple) SetToggleMode(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetToggleMode(enabled)
}
func (self Simple) IsToggleMode() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsToggleMode())
}
func (self Simple) SetShortcutInTooltip(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShortcutInTooltip(enabled)
}
func (self Simple) IsShortcutInTooltipEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShortcutInTooltipEnabled())
}
func (self Simple) SetDisabled(disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisabled(disabled)
}
func (self Simple) IsDisabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDisabled())
}
func (self Simple) SetActionMode(mode classdb.BaseButtonActionMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetActionMode(mode)
}
func (self Simple) GetActionMode() classdb.BaseButtonActionMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseButtonActionMode(Expert(self).GetActionMode())
}
func (self Simple) SetButtonMask(mask gd.MouseButtonMask) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetButtonMask(mask)
}
func (self Simple) GetButtonMask() gd.MouseButtonMask {
	gc := gd.GarbageCollector(); _ = gc
	return gd.MouseButtonMask(Expert(self).GetButtonMask())
}
func (self Simple) GetDrawMode() classdb.BaseButtonDrawMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.BaseButtonDrawMode(Expert(self).GetDrawMode())
}
func (self Simple) SetKeepPressedOutside(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeepPressedOutside(enabled)
}
func (self Simple) IsKeepPressedOutside() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsKeepPressedOutside())
}
func (self Simple) SetShortcutFeedback(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShortcutFeedback(enabled)
}
func (self Simple) IsShortcutFeedback() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShortcutFeedback())
}
func (self Simple) SetShortcut(shortcut [1]classdb.Shortcut) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShortcut(shortcut)
}
func (self Simple) GetShortcut() [1]classdb.Shortcut {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Shortcut(Expert(self).GetShortcut(gc))
}
func (self Simple) SetButtonGroup(button_group [1]classdb.ButtonGroup) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetButtonGroup(button_group)
}
func (self Simple) GetButtonGroup() [1]classdb.ButtonGroup {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ButtonGroup(Expert(self).GetButtonGroup(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.BaseButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) SetShortcut(shortcut object.Shortcut)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShortcut(ctx gd.Lifetime) object.Shortcut {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_get_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Shortcut
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetButtonGroup(button_group object.ButtonGroup)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(button_group[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_set_button_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetButtonGroup(ctx gd.Lifetime) object.ButtonGroup {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.BaseButton.Bind_get_button_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ButtonGroup
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsBaseButton() Expert { return self[0].AsBaseButton() }


//go:nosplit
func (self Simple) AsBaseButton() Simple { return self[0].AsBaseButton() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_pressed": return reflect.ValueOf(self._pressed);
	case "_toggled": return reflect.ValueOf(self._toggled);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_pressed": return reflect.ValueOf(self._pressed);
	case "_toggled": return reflect.ValueOf(self._toggled);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("BaseButton", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
