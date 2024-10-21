package MenuButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Button"
import "grow.graphics/gd/object/BaseButton"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A button that brings up a [PopupMenu] when clicked. To create new items inside this [PopupMenu], use [code]get_popup().add_item("My Item Name")[/code]. You can also create them directly from Godot editor's inspector.
See also [BaseButton] which contains common properties and methods associated with this node.

*/
type Simple [1]classdb.MenuButton
func (self Simple) GetPopup() [1]classdb.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PopupMenu(Expert(self).GetPopup(gc))
}
func (self Simple) ShowPopup() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShowPopup()
}
func (self Simple) SetSwitchOnHover(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSwitchOnHover(enable)
}
func (self Simple) IsSwitchOnHover() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSwitchOnHover())
}
func (self Simple) SetDisableShortcuts(disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableShortcuts(disabled)
}
func (self Simple) SetItemCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemCount(gd.Int(count))
}
func (self Simple) GetItemCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemCount()))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.MenuButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup(ctx gd.Lifetime) object.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuButton.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PopupMenu
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

//go:nosplit
func (self class) AsMenuButton() Expert { return self[0].AsMenuButton() }


//go:nosplit
func (self Simple) AsMenuButton() Simple { return self[0].AsMenuButton() }


//go:nosplit
func (self class) AsButton() Button.Expert { return self[0].AsButton() }


//go:nosplit
func (self Simple) AsButton() Button.Simple { return self[0].AsButton() }


//go:nosplit
func (self class) AsBaseButton() BaseButton.Expert { return self[0].AsBaseButton() }


//go:nosplit
func (self Simple) AsBaseButton() BaseButton.Simple { return self[0].AsBaseButton() }


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
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("MenuButton", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
