package ColorPickerButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Button"
import "grow.graphics/gd/gdclass/BaseButton"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Encapsulates a [ColorPicker], making it accessible by pressing a button. Pressing the button will toggle the [ColorPicker]'s visibility.
See also [BaseButton] which contains common properties and methods associated with this node.
[b]Note:[/b] By default, the button may not be wide enough for the color preview swatch to be visible. Make sure to set [member Control.custom_minimum_size] to a big enough value to give the button enough space.

*/
type Go [1]classdb.ColorPickerButton

/*
Returns the [ColorPicker] that this node toggles.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetPicker() gdclass.ColorPicker {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.ColorPicker(class(self).GetPicker(gc))
}

/*
Returns the control's [PopupPanel] which allows you to connect to popup signals. This allows you to handle events when the ColorPicker is shown or hidden.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Go) GetPopup() gdclass.PopupPanel {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.PopupPanel(class(self).GetPopup(gc))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ColorPickerButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ColorPickerButton"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Color() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetPickColor())
}

func (self Go) SetColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPickColor(value)
}

func (self Go) EditAlpha() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEditingAlpha())
}

func (self Go) SetEditAlpha(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditAlpha(value)
}

//go:nosplit
func (self class) SetPickColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPickerButton.Bind_set_pick_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPickColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPickerButton.Bind_get_pick_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [ColorPicker] that this node toggles.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetPicker(ctx gd.Lifetime) gdclass.ColorPicker {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPickerButton.Bind_get_picker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ColorPicker
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the control's [PopupPanel] which allows you to connect to popup signals. This allows you to handle events when the ColorPicker is shown or hidden.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup(ctx gd.Lifetime) gdclass.PopupPanel {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPickerButton.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PopupPanel
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEditAlpha(show bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, show)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPickerButton.Bind_set_edit_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditingAlpha() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ColorPickerButton.Bind_is_editing_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnColorChanged(cb func(color gd.Color)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("color_changed"), gc.Callable(cb), 0)
}


func (self Go) OnPopupClosed(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("popup_closed"), gc.Callable(cb), 0)
}


func (self Go) OnPickerCreated(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("picker_created"), gc.Callable(cb), 0)
}


func (self class) AsColorPickerButton() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsColorPickerButton() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsButton() Button.GD { return *((*Button.GD)(unsafe.Pointer(&self))) }
func (self Go) AsButton() Button.Go { return *((*Button.Go)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.GD { return *((*BaseButton.GD)(unsafe.Pointer(&self))) }
func (self Go) AsBaseButton() BaseButton.Go { return *((*BaseButton.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsButton(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsButton(), name)
	}
}
func init() {classdb.Register("ColorPickerButton", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
