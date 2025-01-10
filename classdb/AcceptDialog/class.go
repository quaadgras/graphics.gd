// Package AcceptDialog provides methods for working with AcceptDialog object instances.
package AcceptDialog

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Window"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
The default use of [AcceptDialog] is to allow it to only be accepted or closed, with the same result. However, the [signal confirmed] and [signal canceled] signals allow to make the two actions different, and the [method add_button] method allows to add custom buttons and actions.
*/
type Instance [1]gdclass.AcceptDialog

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAcceptDialog() Instance
}

/*
Returns the OK [Button] instance.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetOkButton() [1]gdclass.Button {
	return [1]gdclass.Button(class(self).GetOkButton())
}

/*
Returns the label used for built-in text.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetLabel() [1]gdclass.Label {
	return [1]gdclass.Label(class(self).GetLabel())
}

/*
Adds a button with label [param text] and a custom [param action] to the dialog and returns the created button. [param action] will be passed to the [signal custom_action] signal when pressed.
If [code]true[/code], [param right] will place the button to the right of any sibling buttons.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
func (self Instance) AddButton(text string) [1]gdclass.Button {
	return [1]gdclass.Button(class(self).AddButton(gd.NewString(text), false, gd.NewString("")))
}

/*
Adds a button with label [param name] and a cancel action to the dialog and returns the created button.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
func (self Instance) AddCancelButton(name string) [1]gdclass.Button {
	return [1]gdclass.Button(class(self).AddCancelButton(gd.NewString(name)))
}

/*
Removes the [param button] from the dialog. Does NOT free the [param button]. The [param button] must be a [Button] added with [method add_button] or [method add_cancel_button] method. After removal, pressing the [param button] will no longer emit this dialog's [signal custom_action] or [signal canceled] signals.
*/
func (self Instance) RemoveButton(button [1]gdclass.Button) {
	class(self).RemoveButton(button)
}

/*
Registers a [LineEdit] in the dialog. When the enter key is pressed, the dialog will be accepted.
*/
func (self Instance) RegisterTextEnter(line_edit [1]gdclass.LineEdit) {
	class(self).RegisterTextEnter(line_edit)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AcceptDialog

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AcceptDialog"))
	casted := Instance{*(*gdclass.AcceptDialog)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) OkButtonText() string {
	return string(class(self).GetOkButtonText().String())
}

func (self Instance) SetOkButtonText(value string) {
	class(self).SetOkButtonText(gd.NewString(value))
}

func (self Instance) DialogText() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetDialogText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Instance) DialogHideOnOk() bool {
	return bool(class(self).GetHideOnOk())
}

func (self Instance) SetDialogHideOnOk(value bool) {
	class(self).SetHideOnOk(value)
}

func (self Instance) DialogCloseOnEscape() bool {
	return bool(class(self).GetCloseOnEscape())
}

func (self Instance) SetDialogCloseOnEscape(value bool) {
	class(self).SetCloseOnEscape(value)
}

func (self Instance) DialogAutowrap() bool {
	return bool(class(self).HasAutowrap())
}

func (self Instance) SetDialogAutowrap(value bool) {
	class(self).SetAutowrap(value)
}

/*
Returns the OK [Button] instance.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetOkButton() [1]gdclass.Button {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_ok_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Button{gd.PointerLifetimeBoundTo[gdclass.Button](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the label used for built-in text.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetLabel() [1]gdclass.Label {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_label, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Label{gd.PointerLifetimeBoundTo[gdclass.Label](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHideOnOk(enabled bool) {
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
func (self class) SetCloseOnEscape(enabled bool) {
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
func (self class) AddButton(text gd.String, right bool, action gd.String) [1]gdclass.Button {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, right)
	callframe.Arg(frame, pointers.Get(action))
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_add_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Button{gd.PointerLifetimeBoundTo[gdclass.Button](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds a button with label [param name] and a cancel action to the dialog and returns the created button.
You can use [method remove_button] method to remove a button created with this method from the dialog.
*/
//go:nosplit
func (self class) AddCancelButton(name gd.String) [1]gdclass.Button {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_add_cancel_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Button{gd.PointerLifetimeBoundTo[gdclass.Button](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Removes the [param button] from the dialog. Does NOT free the [param button]. The [param button] must be a [Button] added with [method add_button] or [method add_cancel_button] method. After removal, pressing the [param button] will no longer emit this dialog's [signal custom_action] or [signal canceled] signals.
*/
//go:nosplit
func (self class) RemoveButton(button [1]gdclass.Button) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(button[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_remove_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a [LineEdit] in the dialog. When the enter key is pressed, the dialog will be accepted.
*/
//go:nosplit
func (self class) RegisterTextEnter(line_edit [1]gdclass.LineEdit) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(line_edit[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_register_text_enter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetText(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutowrap(autowrap bool) {
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
func (self class) SetOkButtonText(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_set_ok_button_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOkButtonText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AcceptDialog.Bind_get_ok_button_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnConfirmed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("confirmed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCanceled(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("canceled"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCustomAction(cb func(action string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("custom_action"), gd.NewCallable(cb), 0)
}

func (self class) AsAcceptDialog() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAcceptDialog() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsWindow() Window.Advanced    { return *((*Window.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWindow() Window.Instance { return *((*Window.Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Window.Advanced(self.AsWindow()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Window.Instance(self.AsWindow()), name)
	}
}
func init() {
	gdclass.Register("AcceptDialog", func(ptr gd.Object) any {
		return [1]gdclass.AcceptDialog{*(*gdclass.AcceptDialog)(unsafe.Pointer(&ptr))}
	})
}
