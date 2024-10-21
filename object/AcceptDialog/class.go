package AcceptDialog

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Window"
import "grow.graphics/gd/object/Viewport"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The default use of [AcceptDialog] is to allow it to only be accepted or closed, with the same result. However, the [signal confirmed] and [signal canceled] signals allow to make the two actions different, and the [method add_button] method allows to add custom buttons and actions.

*/
type Simple [1]classdb.AcceptDialog
func (self Simple) GetOkButton() [1]classdb.Button {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Button(Expert(self).GetOkButton(gc))
}
func (self Simple) GetLabel() [1]classdb.Label {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Label(Expert(self).GetLabel(gc))
}
func (self Simple) SetHideOnOk(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHideOnOk(enabled)
}
func (self Simple) GetHideOnOk() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetHideOnOk())
}
func (self Simple) SetCloseOnEscape(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCloseOnEscape(enabled)
}
func (self Simple) GetCloseOnEscape() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetCloseOnEscape())
}
func (self Simple) AddButton(text string, right bool, action string) [1]classdb.Button {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Button(Expert(self).AddButton(gc, gc.String(text), right, gc.String(action)))
}
func (self Simple) AddCancelButton(name string) [1]classdb.Button {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Button(Expert(self).AddCancelButton(gc, gc.String(name)))
}
func (self Simple) RemoveButton(button [1]classdb.Button) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveButton(button)
}
func (self Simple) RegisterTextEnter(line_edit [1]classdb.LineEdit) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegisterTextEnter(line_edit)
}
func (self Simple) SetText(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetText(gc.String(text))
}
func (self Simple) GetText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetText(gc).String())
}
func (self Simple) SetAutowrap(autowrap bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutowrap(autowrap)
}
func (self Simple) HasAutowrap() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAutowrap())
}
func (self Simple) SetOkButtonText(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOkButtonText(gc.String(text))
}
func (self Simple) GetOkButtonText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetOkButtonText(gc).String())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AcceptDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the OK [Button] instance.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetOkButton(ctx gd.Lifetime) object.Button {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_get_ok_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Button
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the label used for built-in text.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLabel(ctx gd.Lifetime) object.Label {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Label
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHideOnOk(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_set_hide_on_ok, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHideOnOk() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_get_hide_on_ok, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCloseOnEscape(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_set_close_on_escape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCloseOnEscape() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_get_close_on_escape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a button with label [param text] and a custom [param action] to the dialog and returns the created button. [param action] will be passed to the [signal custom_action] signal when pressed.
If [code]true[/code], [param right] will place the button to the right of any sibling buttons.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
//go:nosplit
func (self class) AddButton(ctx gd.Lifetime, text gd.String, right bool, action gd.String) object.Button {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, right)
	callframe.Arg(frame, mmm.Get(action))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_add_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Button
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adds a button with label [param name] and a cancel action to the dialog and returns the created button.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
//go:nosplit
func (self class) AddCancelButton(ctx gd.Lifetime, name gd.String) object.Button {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_add_cancel_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Button
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Removes the [param button] from the dialog. Does NOT free the [param button]. The [param button] must be a [Button] added with [method add_button] or [method add_cancel_button] method. After removal, pressing the [param button] will no longer emit this dialog's [signal custom_action] or [signal canceled] signals.
*/
//go:nosplit
func (self class) RemoveButton(button object.Button)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(button[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_remove_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers a [LineEdit] in the dialog. When the enter key is pressed, the dialog will be accepted.
*/
//go:nosplit
func (self class) RegisterTextEnter(line_edit object.LineEdit)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(line_edit[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_register_text_enter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutowrap(autowrap bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, autowrap)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_set_autowrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasAutowrap() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_has_autowrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOkButtonText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_set_ok_button_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOkButtonText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_get_ok_button_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAcceptDialog() Expert { return self[0].AsAcceptDialog() }


//go:nosplit
func (self Simple) AsAcceptDialog() Simple { return self[0].AsAcceptDialog() }


//go:nosplit
func (self class) AsWindow() Window.Expert { return self[0].AsWindow() }


//go:nosplit
func (self Simple) AsWindow() Window.Simple { return self[0].AsWindow() }


//go:nosplit
func (self class) AsViewport() Viewport.Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Viewport.Simple { return self[0].AsViewport() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("AcceptDialog", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
