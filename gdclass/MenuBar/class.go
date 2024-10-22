package MenuBar

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A horizontal menu bar that creates a [MenuButton] for each [PopupMenu] child. New items are created by adding [PopupMenu]s to this node.

*/
type Go [1]classdb.MenuBar

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
func (self Go) SetDisableShortcuts(disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisableShortcuts(disabled)
}

/*
Returns [code]true[/code], if system global menu is supported and used by this [MenuBar].
*/
func (self Go) IsNativeMenu() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsNativeMenu())
}

/*
Returns number of menu items.
*/
func (self Go) GetMenuCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetMenuCount()))
}

/*
Sets menu item title.
*/
func (self Go) SetMenuTitle(menu int, title string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMenuTitle(gd.Int(menu), gc.String(title))
}

/*
Returns menu item title.
*/
func (self Go) GetMenuTitle(menu int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetMenuTitle(gc, gd.Int(menu)).String())
}

/*
Sets menu item tooltip.
*/
func (self Go) SetMenuTooltip(menu int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMenuTooltip(gd.Int(menu), gc.String(tooltip))
}

/*
Returns menu item tooltip.
*/
func (self Go) GetMenuTooltip(menu int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetMenuTooltip(gc, gd.Int(menu)).String())
}

/*
If [code]true[/code], menu item is disabled.
*/
func (self Go) SetMenuDisabled(menu int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMenuDisabled(gd.Int(menu), disabled)
}

/*
Returns [code]true[/code], if menu item is disabled.
*/
func (self Go) IsMenuDisabled(menu int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsMenuDisabled(gd.Int(menu)))
}

/*
If [code]true[/code], menu item is hidden.
*/
func (self Go) SetMenuHidden(menu int, hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMenuHidden(gd.Int(menu), hidden)
}

/*
Returns [code]true[/code], if menu item is hidden.
*/
func (self Go) IsMenuHidden(menu int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsMenuHidden(gd.Int(menu)))
}

/*
Returns [PopupMenu] associated with menu item.
*/
func (self Go) GetMenuPopup(menu int) gdclass.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.PopupMenu(class(self).GetMenuPopup(gc, gd.Int(menu)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MenuBar
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MenuBar"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Flat() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFlat())
}

func (self Go) SetFlat(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFlat(value)
}

func (self Go) StartIndex() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetStartIndex()))
}

func (self Go) SetStartIndex(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStartIndex(gd.Int(value))
}

func (self Go) SwitchOnHover() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSwitchOnHover())
}

func (self Go) SetSwitchOnHover(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSwitchOnHover(value)
}

func (self Go) PreferGlobalMenu() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsPreferGlobalMenu())
}

func (self Go) SetPreferGlobalMenu(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPreferGlobalMenu(value)
}

func (self Go) TextDirection() classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.ControlTextDirection(class(self).GetTextDirection())
}

func (self Go) SetTextDirection(value classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextDirection(value)
}

func (self Go) Language() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetLanguage(gc).String())
}

func (self Go) SetLanguage(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLanguage(gc.String(value))
}

//go:nosplit
func (self class) SetSwitchOnHover(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSwitchOnHover() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_is_switch_on_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_disable_shortcuts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetPreferGlobalMenu(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_prefer_global_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPreferGlobalMenu() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_is_prefer_global_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code], if system global menu is supported and used by this [MenuBar].
*/
//go:nosplit
func (self class) IsNativeMenu() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_is_native_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of menu items.
*/
//go:nosplit
func (self class) GetMenuCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_get_menu_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLanguage(language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLanguage(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlat(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFlat() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_is_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStartIndex(enabled gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_start_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStartIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_get_start_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets menu item title.
*/
//go:nosplit
func (self class) SetMenuTitle(menu gd.Int, title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_menu_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns menu item title.
*/
//go:nosplit
func (self class) GetMenuTitle(ctx gd.Lifetime, menu gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_get_menu_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets menu item tooltip.
*/
//go:nosplit
func (self class) SetMenuTooltip(menu gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_menu_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns menu item tooltip.
*/
//go:nosplit
func (self class) GetMenuTooltip(ctx gd.Lifetime, menu gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_get_menu_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If [code]true[/code], menu item is disabled.
*/
//go:nosplit
func (self class) SetMenuDisabled(menu gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_menu_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code], if menu item is disabled.
*/
//go:nosplit
func (self class) IsMenuDisabled(menu gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_is_menu_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], menu item is hidden.
*/
//go:nosplit
func (self class) SetMenuHidden(menu gd.Int, hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_set_menu_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code], if menu item is hidden.
*/
//go:nosplit
func (self class) IsMenuHidden(menu gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_is_menu_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [PopupMenu] associated with menu item.
*/
//go:nosplit
func (self class) GetMenuPopup(ctx gd.Lifetime, menu gd.Int) gdclass.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MenuBar.Bind_get_menu_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PopupMenu
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
func (self class) AsMenuBar() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMenuBar() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("MenuBar", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
