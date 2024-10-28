package AcceptDialog

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

/*
The default use of [AcceptDialog] is to allow it to only be accepted or closed, with the same result. However, the [signal confirmed] and [signal canceled] signals allow to make the two actions different, and the [method add_button] method allows to add custom buttons and actions.

*/
type Go [1]classdb.AcceptDialog

/*
Returns the OK [Button] instance.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetOkButton() gdclass.Button {
	return gdclass.Button(class(self).GetOkButton())
}

/*
Returns the label used for built-in text.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetLabel() gdclass.Label {
	return gdclass.Label(class(self).GetLabel())
}

/*
Adds a button with label [param text] and a custom [param action] to the dialog and returns the created button. [param action] will be passed to the [signal custom_action] signal when pressed.
If [code]true[/code], [param right] will place the button to the right of any sibling buttons.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
func (self Go) AddButton(text string) gdclass.Button {
	return gdclass.Button(class(self).AddButton(gd.NewString(text), false, gd.NewString("")))
}

/*
Adds a button with label [param name] and a cancel action to the dialog and returns the created button.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
func (self Go) AddCancelButton(name string) gdclass.Button {
	return gdclass.Button(class(self).AddCancelButton(gd.NewString(name)))
}

/*
Removes the [param button] from the dialog. Does NOT free the [param button]. The [param button] must be a [Button] added with [method add_button] or [method add_cancel_button] method. After removal, pressing the [param button] will no longer emit this dialog's [signal custom_action] or [signal canceled] signals.
*/
func (self Go) RemoveButton(button gdclass.Button) {
	class(self).RemoveButton(button)
}

/*
Registers a [LineEdit] in the dialog. When the enter key is pressed, the dialog will be accepted.
*/
func (self Go) RegisterTextEnter(line_edit gdclass.LineEdit) {
	class(self).RegisterTextEnter(line_edit)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AcceptDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AcceptDialog"))
	return Go{classdb.AcceptDialog(object)}
}

func (self Go) OkButtonText() string {
		return string(class(self).GetOkButtonText().String())
}

func (self Go) SetOkButtonText(value string) {
	class(self).SetOkButtonText(gd.NewString(value))
}

func (self Go) DialogText() string {
		return string(class(self).GetText().String())
}

func (self Go) SetDialogText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Go) DialogHideOnOk() bool {
		return bool(class(self).GetHideOnOk())
}

func (self Go) SetDialogHideOnOk(value bool) {
	class(self).SetHideOnOk(value)
}

func (self Go) DialogCloseOnEscape() bool {
		return bool(class(self).GetCloseOnEscape())
}

func (self Go) SetDialogCloseOnEscape(value bool) {
	class(self).SetCloseOnEscape(value)
}

func (self Go) DialogAutowrap() bool {
		return bool(class(self).HasAutowrap())
}

func (self Go) SetDialogAutowrap(value bool) {
	class(self).SetAutowrap(value)
}

/*
Returns the OK [Button] instance.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetOkButton() gdclass.Button {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_ok_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Button{classdb.Button(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the label used for built-in text.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLabel() gdclass.Label {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Label{classdb.Label(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHideOnOk(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_set_hide_on_ok, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHideOnOk() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_hide_on_ok, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCloseOnEscape(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_set_close_on_escape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCloseOnEscape() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_close_on_escape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AddButton(text gd.String, right bool, action gd.String) gdclass.Button {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(text))
	callframe.Arg(frame, right)
	callframe.Arg(frame, discreet.Get(action))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_add_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Button{classdb.Button(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Adds a button with label [param name] and a cancel action to the dialog and returns the created button.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
//go:nosplit
func (self class) AddCancelButton(name gd.String) gdclass.Button {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_add_cancel_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Button{classdb.Button(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Removes the [param button] from the dialog. Does NOT free the [param button]. The [param button] must be a [Button] added with [method add_button] or [method add_cancel_button] method. After removal, pressing the [param button] will no longer emit this dialog's [signal custom_action] or [signal canceled] signals.
*/
//go:nosplit
func (self class) RemoveButton(button gdclass.Button)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(button[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_remove_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers a [LineEdit] in the dialog. When the enter key is pressed, the dialog will be accepted.
*/
//go:nosplit
func (self class) RegisterTextEnter(line_edit gdclass.LineEdit)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(line_edit[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_register_text_enter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetText(text gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutowrap(autowrap bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, autowrap)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_set_autowrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasAutowrap() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_has_autowrap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOkButtonText(text gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_set_ok_button_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOkButtonText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_ok_button_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self Go) OnConfirmed(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("confirmed"), gd.NewCallable(cb), 0)
}


func (self Go) OnCanceled(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("canceled"), gd.NewCallable(cb), 0)
}


func (self Go) OnCustomAction(cb func(action string)) {
	self[0].AsObject().Connect(gd.NewStringName("custom_action"), gd.NewCallable(cb), 0)
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
func init() {classdb.Register("AcceptDialog", func(ptr gd.Object) any { return classdb.AcceptDialog(ptr) })}
