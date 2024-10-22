package MenuButton

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
A button that brings up a [PopupMenu] when clicked. To create new items inside this [PopupMenu], use [code]get_popup().add_item("My Item Name")[/code]. You can also create them directly from Godot editor's inspector.
See also [BaseButton] which contains common properties and methods associated with this node.

*/
type Go [1]classdb.MenuButton

/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Go) GetPopup() gdclass.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.PopupMenu(class(self).GetPopup(gc))
}

/*
Adjusts popup position and sizing for the [MenuButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
func (self Go) ShowPopup() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ShowPopup()
}

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
func (self Go) SetDisableShortcuts(disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisableShortcuts(disabled)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MenuButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MenuButton"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) SwitchOnHover() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSwitchOnHover())
}

func (self Go) SetSwitchOnHover(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSwitchOnHover(value)
}

func (self Go) ItemCount() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetItemCount()))
}

func (self Go) SetItemCount(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetItemCount(gd.Int(value))
}

/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup(ctx gd.Lifetime) gdclass.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuButton.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PopupMenu
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adjusts popup position and sizing for the [MenuButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
//go:nosplit
func (self class) ShowPopup()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuButton.Bind_show_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSwitchOnHover(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuButton.Bind_set_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSwitchOnHover() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuButton.Bind_is_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
//go:nosplit
func (self class) SetDisableShortcuts(disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuButton.Bind_set_disable_shortcuts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetItemCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuButton.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetItemCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuButton.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnAboutToPopup(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("about_to_popup"), gc.Callable(cb), 0)
}


func (self class) AsMenuButton() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMenuButton() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("MenuButton", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
