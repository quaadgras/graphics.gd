// Package ColorPickerButton provides methods for working with ColorPickerButton object instances.
package ColorPickerButton

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Button"
import "graphics.gd/classdb/BaseButton"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Color"

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
Encapsulates a [ColorPicker], making it accessible by pressing a button. Pressing the button will toggle the [ColorPicker]'s visibility.
See also [BaseButton] which contains common properties and methods associated with this node.
[b]Note:[/b] By default, the button may not be wide enough for the color preview swatch to be visible. Make sure to set [member Control.custom_minimum_size] to a big enough value to give the button enough space.
*/
type Instance [1]gdclass.ColorPickerButton

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsColorPickerButton() Instance
}

/*
Returns the [ColorPicker] that this node toggles.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetPicker() [1]gdclass.ColorPicker {
	return [1]gdclass.ColorPicker(class(self).GetPicker())
}

/*
Returns the control's [PopupPanel] which allows you to connect to popup signals. This allows you to handle events when the ColorPicker is shown or hidden.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Instance) GetPopup() [1]gdclass.PopupPanel {
	return [1]gdclass.PopupPanel(class(self).GetPopup())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ColorPickerButton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ColorPickerButton"))
	casted := Instance{*(*gdclass.ColorPickerButton)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetPickColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetPickColor(gd.Color(value))
}

func (self Instance) EditAlpha() bool {
	return bool(class(self).IsEditingAlpha())
}

func (self Instance) SetEditAlpha(value bool) {
	class(self).SetEditAlpha(value)
}

//go:nosplit
func (self class) SetPickColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPickerButton.Bind_set_pick_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPickColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPickerButton.Bind_get_pick_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [ColorPicker] that this node toggles.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetPicker() [1]gdclass.ColorPicker {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPickerButton.Bind_get_picker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.ColorPicker{gd.PointerLifetimeBoundTo[gdclass.ColorPicker](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the control's [PopupPanel] which allows you to connect to popup signals. This allows you to handle events when the ColorPicker is shown or hidden.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup() [1]gdclass.PopupPanel {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPickerButton.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.PopupPanel{gd.PointerLifetimeBoundTo[gdclass.PopupPanel](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEditAlpha(show bool) {
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPickerButton.Bind_set_edit_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditingAlpha() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ColorPickerButton.Bind_is_editing_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnColorChanged(cb func(color Color.RGBA)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("color_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPopupClosed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("popup_closed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPickerCreated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("picker_created"), gd.NewCallable(cb), 0)
}

func (self class) AsColorPickerButton() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsColorPickerButton() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsButton() Button.Advanced        { return *((*Button.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsButton() Button.Instance     { return *((*Button.Instance)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.Advanced {
	return *((*BaseButton.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBaseButton() BaseButton.Instance {
	return *((*BaseButton.Instance)(unsafe.Pointer(&self)))
}
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
	default:
		return gd.VirtualByName(Button.Advanced(self.AsButton()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Button.Instance(self.AsButton()), name)
	}
}
func init() {
	gdclass.Register("ColorPickerButton", func(ptr gd.Object) any {
		return [1]gdclass.ColorPickerButton{*(*gdclass.ColorPickerButton)(unsafe.Pointer(&ptr))}
	})
}
