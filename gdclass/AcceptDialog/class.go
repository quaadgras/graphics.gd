package AcceptDialog

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Window"
import "grow.graphics/gd/gdclass/Viewport"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The default use of [AcceptDialog] is to allow it to only be accepted or closed, with the same result. However, the [signal confirmed] and [signal canceled] signals allow to make the two actions different, and the [method add_button] method allows to add custom buttons and actions.

*/
type Go [1]classdb.AcceptDialog

/*
Returns the OK [Button] instance.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetOkButton() gdclass.Button {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Button(class(self).GetOkButton(gc))
}

/*
Returns the label used for built-in text.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetLabel() gdclass.Label {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Label(class(self).GetLabel(gc))
}

/*
Adds a button with label [param text] and a custom [param action] to the dialog and returns the created button. [param action] will be passed to the [signal custom_action] signal when pressed.
If [code]true[/code], [param right] will place the button to the right of any sibling buttons.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
func (self Go) AddButton(text string) gdclass.Button {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Button(class(self).AddButton(gc, gc.String(text), false, gc.String("")))
}

/*
Adds a button with label [param name] and a cancel action to the dialog and returns the created button.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
func (self Go) AddCancelButton(name string) gdclass.Button {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Button(class(self).AddCancelButton(gc, gc.String(name)))
}

/*
Removes the [param button] from the dialog. Does NOT free the [param button]. The [param button] must be a [Button] added with [method add_button] or [method add_cancel_button] method. After removal, pressing the [param button] will no longer emit this dialog's [signal custom_action] or [signal canceled] signals.
*/
func (self Go) RemoveButton(button gdclass.Button) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveButton(button)
}

/*
Registers a [LineEdit] in the dialog. When the enter key is pressed, the dialog will be accepted.
*/
func (self Go) RegisterTextEnter(line_edit gdclass.LineEdit) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RegisterTextEnter(line_edit)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AcceptDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AcceptDialog"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) OkButtonText() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetOkButtonText(gc).String())
}

func (self Go) SetOkButtonText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOkButtonText(gc.String(value))
}

func (self Go) DialogText() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetText(gc).String())
}

func (self Go) SetDialogText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetText(gc.String(value))
}

func (self Go) DialogHideOnOk() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetHideOnOk())
}

func (self Go) SetDialogHideOnOk(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHideOnOk(value)
}

func (self Go) DialogCloseOnEscape() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetCloseOnEscape())
}

func (self Go) SetDialogCloseOnEscape(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCloseOnEscape(value)
}

func (self Go) DialogAutowrap() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).HasAutowrap())
}

func (self Go) SetDialogAutowrap(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutowrap(value)
}

/*
Returns the OK [Button] instance.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetOkButton(ctx gd.Lifetime) gdclass.Button {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_get_ok_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Button
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the label used for built-in text.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLabel(ctx gd.Lifetime) gdclass.Label {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Label
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
func (self class) AddButton(ctx gd.Lifetime, text gd.String, right bool, action gd.String) gdclass.Button {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, right)
	callframe.Arg(frame, mmm.Get(action))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_add_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Button
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adds a button with label [param name] and a cancel action to the dialog and returns the created button.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
//go:nosplit
func (self class) AddCancelButton(ctx gd.Lifetime, name gd.String) gdclass.Button {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AcceptDialog.Bind_add_cancel_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Button
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Removes the [param button] from the dialog. Does NOT free the [param button]. The [param button] must be a [Button] added with [method add_button] or [method add_cancel_button] method. After removal, pressing the [param button] will no longer emit this dialog's [signal custom_action] or [signal canceled] signals.
*/
//go:nosplit
func (self class) RemoveButton(button gdclass.Button)  {
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
func (self class) RegisterTextEnter(line_edit gdclass.LineEdit)  {
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
func (self Go) OnConfirmed(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("confirmed"), gc.Callable(cb), 0)
}


func (self Go) OnCanceled(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("canceled"), gc.Callable(cb), 0)
}


func (self Go) OnCustomAction(cb func(action string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("custom_action"), gc.Callable(cb), 0)
}


func (self class) AsAcceptDialog() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAcceptDialog() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsWindow() Window.GD { return *((*Window.GD)(unsafe.Pointer(&self))) }
func (self Go) AsWindow() Window.Go { return *((*Window.Go)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.GD { return *((*Viewport.GD)(unsafe.Pointer(&self))) }
func (self Go) AsViewport() Viewport.Go { return *((*Viewport.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsWindow(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsWindow(), name)
	}
}
func init() {classdb.Register("AcceptDialog", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
