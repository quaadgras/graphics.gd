// Package MenuButton provides methods for working with MenuButton object instances.
package MenuButton

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/Button"
import "graphics.gd/classdb/BaseButton"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A button that brings up a [PopupMenu] when clicked. To create new items inside this [PopupMenu], use [code]get_popup().add_item("My Item Name")[/code]. You can also create them directly from Godot editor's inspector.
See also [BaseButton] which contains common properties and methods associated with this node.
*/
type Instance [1]gdclass.MenuButton

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMenuButton() Instance
}

/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Instance) GetPopup() [1]gdclass.PopupMenu { //gd:MenuButton.get_popup
	return [1]gdclass.PopupMenu(class(self).GetPopup())
}

/*
Adjusts popup position and sizing for the [MenuButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
func (self Instance) ShowPopup() { //gd:MenuButton.show_popup
	class(self).ShowPopup()
}

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
func (self Instance) SetDisableShortcuts(disabled bool) { //gd:MenuButton.set_disable_shortcuts
	class(self).SetDisableShortcuts(disabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MenuButton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MenuButton"))
	casted := Instance{*(*gdclass.MenuButton)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) SwitchOnHover() bool {
	return bool(class(self).IsSwitchOnHover())
}

func (self Instance) SetSwitchOnHover(value bool) {
	class(self).SetSwitchOnHover(value)
}

func (self Instance) ItemCount() int {
	return int(int(class(self).GetItemCount()))
}

func (self Instance) SetItemCount(value int) {
	class(self).SetItemCount(gd.Int(value))
}

/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup() [1]gdclass.PopupMenu { //gd:MenuButton.get_popup
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuButton.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PopupMenu{gd.PointerLifetimeBoundTo[gdclass.PopupMenu](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adjusts popup position and sizing for the [MenuButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
//go:nosplit
func (self class) ShowPopup() { //gd:MenuButton.show_popup
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuButton.Bind_show_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSwitchOnHover(enable bool) { //gd:MenuButton.set_switch_on_hover
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuButton.Bind_set_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSwitchOnHover() bool { //gd:MenuButton.is_switch_on_hover
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuButton.Bind_is_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
//go:nosplit
func (self class) SetDisableShortcuts(disabled bool) { //gd:MenuButton.set_disable_shortcuts
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuButton.Bind_set_disable_shortcuts, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetItemCount(count gd.Int) { //gd:MenuButton.set_item_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuButton.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetItemCount() gd.Int { //gd:MenuButton.get_item_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MenuButton.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnAboutToPopup(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("about_to_popup"), gd.NewCallable(cb), 0)
}

func (self class) AsMenuButton() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMenuButton() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsButton() Button.Advanced    { return *((*Button.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsButton() Button.Instance { return *((*Button.Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("MenuButton", func(ptr gd.Object) any { return [1]gdclass.MenuButton{*(*gdclass.MenuButton)(unsafe.Pointer(&ptr))} })
}
