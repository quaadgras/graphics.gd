package NativeMenu

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[NativeMenu] handles low-level access to the OS native global menu bar and popup menus.
[b]Note:[/b] This is low-level API, consider using [MenuBar] with [member MenuBar.prefer_global_menu] set to [code]true[/code], and [PopupMenu] with [member PopupMenu.prefer_native_menu] set to [code]true[/code].
To create a menu, use [method create_menu], add menu items using [code]add_*_item[/code] methods. To remove a menu, use [method free_menu].
[codeblock]
var menu

func _menu_callback(item_id):
    if item_id == "ITEM_CUT":
        cut()
    elif item_id == "ITEM_COPY":
        copy()
    elif item_id == "ITEM_PASTE":
        paste()

func _enter_tree():
    # Create new menu and add items:
    menu = NativeMenu.create_menu()
    NativeMenu.add_item(menu, "Cut", _menu_callback, Callable(), "ITEM_CUT")
    NativeMenu.add_item(menu, "Copy", _menu_callback, Callable(), "ITEM_COPY")
    NativeMenu.add_separator(menu)
    NativeMenu.add_item(menu, "Paste", _menu_callback, Callable(), "ITEM_PASTE")

func _on_button_pressed():
    # Show popup menu at mouse position:
    NativeMenu.popup(menu, DisplayServer.mouse_get_position())

func _exit_tree():
    # Remove menu when it's no longer needed:
    NativeMenu.free_menu(menu)
[/codeblock]

*/
type Simple [1]classdb.NativeMenu
func (self Simple) HasFeature(feature classdb.NativeMenuFeature) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasFeature(feature))
}
func (self Simple) HasSystemMenu(menu_id classdb.NativeMenuSystemMenus) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasSystemMenu(menu_id))
}
func (self Simple) GetSystemMenu(menu_id classdb.NativeMenuSystemMenus) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetSystemMenu(menu_id))
}
func (self Simple) GetSystemMenuName(menu_id classdb.NativeMenuSystemMenus) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSystemMenuName(gc, menu_id).String())
}
func (self Simple) CreateMenu() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).CreateMenu())
}
func (self Simple) HasMenu(rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasMenu(rid))
}
func (self Simple) FreeMenu(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FreeMenu(rid)
}
func (self Simple) GetSize(rid gd.RID) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetSize(rid))
}
func (self Simple) Popup(rid gd.RID, position gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Popup(rid, position)
}
func (self Simple) SetInterfaceDirection(rid gd.RID, is_rtl bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInterfaceDirection(rid, is_rtl)
}
func (self Simple) SetPopupOpenCallback(rid gd.RID, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPopupOpenCallback(rid, callback)
}
func (self Simple) GetPopupOpenCallback(rid gd.RID) gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Callable(Expert(self).GetPopupOpenCallback(gc, rid))
}
func (self Simple) SetPopupCloseCallback(rid gd.RID, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPopupCloseCallback(rid, callback)
}
func (self Simple) GetPopupCloseCallback(rid gd.RID) gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Callable(Expert(self).GetPopupCloseCallback(gc, rid))
}
func (self Simple) SetMinimumWidth(rid gd.RID, width float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinimumWidth(rid, gd.Float(width))
}
func (self Simple) GetMinimumWidth(rid gd.RID) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMinimumWidth(rid)))
}
func (self Simple) IsOpened(rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsOpened(rid))
}
func (self Simple) AddSubmenuItem(rid gd.RID, label string, submenu_rid gd.RID, tag gd.Variant, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddSubmenuItem(rid, gc.String(label), submenu_rid, tag, gd.Int(index))))
}
func (self Simple) AddItem(rid gd.RID, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddItem(rid, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) AddCheckItem(rid gd.RID, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddCheckItem(rid, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) AddIconItem(rid gd.RID, icon [1]classdb.Texture2D, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddIconItem(rid, icon, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) AddIconCheckItem(rid gd.RID, icon [1]classdb.Texture2D, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddIconCheckItem(rid, icon, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) AddRadioCheckItem(rid gd.RID, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddRadioCheckItem(rid, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) AddIconRadioCheckItem(rid gd.RID, icon [1]classdb.Texture2D, label string, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddIconRadioCheckItem(rid, icon, gc.String(label), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) AddMultistateItem(rid gd.RID, label string, max_states int, default_state int, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddMultistateItem(rid, gc.String(label), gd.Int(max_states), gd.Int(default_state), callback, key_callback, tag, accelerator, gd.Int(index))))
}
func (self Simple) AddSeparator(rid gd.RID, index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddSeparator(rid, gd.Int(index))))
}
func (self Simple) FindItemIndexWithText(rid gd.RID, text string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindItemIndexWithText(rid, gc.String(text))))
}
func (self Simple) FindItemIndexWithTag(rid gd.RID, tag gd.Variant) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindItemIndexWithTag(rid, tag)))
}
func (self Simple) FindItemIndexWithSubmenu(rid gd.RID, submenu_rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).FindItemIndexWithSubmenu(rid, submenu_rid)))
}
func (self Simple) IsItemChecked(rid gd.RID, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemChecked(rid, gd.Int(idx)))
}
func (self Simple) IsItemCheckable(rid gd.RID, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemCheckable(rid, gd.Int(idx)))
}
func (self Simple) IsItemRadioCheckable(rid gd.RID, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemRadioCheckable(rid, gd.Int(idx)))
}
func (self Simple) GetItemCallback(rid gd.RID, idx int) gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Callable(Expert(self).GetItemCallback(gc, rid, gd.Int(idx)))
}
func (self Simple) GetItemKeyCallback(rid gd.RID, idx int) gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Callable(Expert(self).GetItemKeyCallback(gc, rid, gd.Int(idx)))
}
func (self Simple) GetItemTag(rid gd.RID, idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetItemTag(gc, rid, gd.Int(idx)))
}
func (self Simple) GetItemText(rid gd.RID, idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemText(gc, rid, gd.Int(idx)).String())
}
func (self Simple) GetItemSubmenu(rid gd.RID, idx int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetItemSubmenu(rid, gd.Int(idx)))
}
func (self Simple) GetItemAccelerator(rid gd.RID, idx int) gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GetItemAccelerator(rid, gd.Int(idx)))
}
func (self Simple) IsItemDisabled(rid gd.RID, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemDisabled(rid, gd.Int(idx)))
}
func (self Simple) IsItemHidden(rid gd.RID, idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemHidden(rid, gd.Int(idx)))
}
func (self Simple) GetItemTooltip(rid gd.RID, idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemTooltip(gc, rid, gd.Int(idx)).String())
}
func (self Simple) GetItemState(rid gd.RID, idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemState(rid, gd.Int(idx))))
}
func (self Simple) GetItemMaxStates(rid gd.RID, idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemMaxStates(rid, gd.Int(idx))))
}
func (self Simple) GetItemIcon(rid gd.RID, idx int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetItemIcon(gc, rid, gd.Int(idx)))
}
func (self Simple) GetItemIndentationLevel(rid gd.RID, idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemIndentationLevel(rid, gd.Int(idx))))
}
func (self Simple) SetItemChecked(rid gd.RID, idx int, checked bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemChecked(rid, gd.Int(idx), checked)
}
func (self Simple) SetItemCheckable(rid gd.RID, idx int, checkable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemCheckable(rid, gd.Int(idx), checkable)
}
func (self Simple) SetItemRadioCheckable(rid gd.RID, idx int, checkable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemRadioCheckable(rid, gd.Int(idx), checkable)
}
func (self Simple) SetItemCallback(rid gd.RID, idx int, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemCallback(rid, gd.Int(idx), callback)
}
func (self Simple) SetItemHoverCallbacks(rid gd.RID, idx int, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemHoverCallbacks(rid, gd.Int(idx), callback)
}
func (self Simple) SetItemKeyCallback(rid gd.RID, idx int, key_callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemKeyCallback(rid, gd.Int(idx), key_callback)
}
func (self Simple) SetItemTag(rid gd.RID, idx int, tag gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemTag(rid, gd.Int(idx), tag)
}
func (self Simple) SetItemText(rid gd.RID, idx int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemText(rid, gd.Int(idx), gc.String(text))
}
func (self Simple) SetItemSubmenu(rid gd.RID, idx int, submenu_rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemSubmenu(rid, gd.Int(idx), submenu_rid)
}
func (self Simple) SetItemAccelerator(rid gd.RID, idx int, keycode gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemAccelerator(rid, gd.Int(idx), keycode)
}
func (self Simple) SetItemDisabled(rid gd.RID, idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemDisabled(rid, gd.Int(idx), disabled)
}
func (self Simple) SetItemHidden(rid gd.RID, idx int, hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemHidden(rid, gd.Int(idx), hidden)
}
func (self Simple) SetItemTooltip(rid gd.RID, idx int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemTooltip(rid, gd.Int(idx), gc.String(tooltip))
}
func (self Simple) SetItemState(rid gd.RID, idx int, state int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemState(rid, gd.Int(idx), gd.Int(state))
}
func (self Simple) SetItemMaxStates(rid gd.RID, idx int, max_states int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemMaxStates(rid, gd.Int(idx), gd.Int(max_states))
}
func (self Simple) SetItemIcon(rid gd.RID, idx int, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIcon(rid, gd.Int(idx), icon)
}
func (self Simple) SetItemIndentationLevel(rid gd.RID, idx int, level int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIndentationLevel(rid, gd.Int(idx), gd.Int(level))
}
func (self Simple) GetItemCount(rid gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemCount(rid)))
}
func (self Simple) IsSystemMenu(rid gd.RID) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSystemMenu(rid))
}
func (self Simple) RemoveItem(rid gd.RID, idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveItem(rid, gd.Int(idx))
}
func (self Simple) Clear(rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear(rid)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NativeMenu
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns [code]true[/code] if the specified [param feature] is supported by the current [NativeMenu], [code]false[/code] otherwise.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) HasFeature(feature classdb.NativeMenuFeature) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a special system menu is supported.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) HasSystemMenu(menu_id classdb.NativeMenuSystemMenus) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu_id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_has_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns RID of a special system menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetSystemMenu(menu_id classdb.NativeMenuSystemMenus) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu_id)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns readable name of a special system menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetSystemMenuName(ctx gd.Lifetime, menu_id classdb.NativeMenuSystemMenus) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, menu_id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_system_menu_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a new global menu object.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) CreateMenu() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_create_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [param rid] is valid global menu.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) HasMenu(rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_has_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Frees a global menu object created by this [NativeMenu].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) FreeMenu(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_free_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns global menu size.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetSize(rid gd.RID) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Shows the global menu at [param position] in the screen coordinates.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) Popup(rid gd.RID, position gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the menu text layout direction from right-to-left if [param is_rtl] is [code]true[/code].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetInterfaceDirection(rid gd.RID, is_rtl bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, is_rtl)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_interface_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers callable to emit after the menu is closed.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetPopupOpenCallback(rid gd.RID, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_popup_open_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns global menu open callback.
b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetPopupOpenCallback(ctx gd.Lifetime, rid gd.RID) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_popup_open_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Registers callable to emit when the menu is about to show.
[b]Note:[/b] The OS can simulate menu opening to track menu item changes and global shortcuts, in which case the corresponding close callback is not triggered. Use [method is_opened] to check if the menu is currently opened.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetPopupCloseCallback(rid gd.RID, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_popup_close_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns global menu close callback.
b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetPopupCloseCallback(ctx gd.Lifetime, rid gd.RID) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_popup_close_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the minimum width of the global menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetMinimumWidth(rid gd.RID, width gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_minimum_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns global menu minimum width.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetMinimumWidth(rid gd.RID) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_minimum_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the menu is currently opened.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) IsOpened(rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_is_opened, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an item that will act as a submenu of the global menu [param rid]. The [param submenu_rid] argument is the RID of the global menu that will be shown when the item is clicked.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) AddSubmenuItem(rid gd.RID, label gd.String, submenu_rid gd.RID, tag gd.Variant, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, submenu_rid)
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_submenu_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new item with text [param label] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
//go:nosplit
func (self class) AddItem(rid gd.RID, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new checkable item with text [param label] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
//go:nosplit
func (self class) AddCheckItem(rid gd.RID, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new item with text [param label] and icon [param icon] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
//go:nosplit
func (self class) AddIconItem(rid gd.RID, icon object.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new checkable item with text [param label] and icon [param icon] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
//go:nosplit
func (self class) AddIconCheckItem(rid gd.RID, icon object.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_icon_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new radio-checkable item with text [param label] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] Radio-checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
//go:nosplit
func (self class) AddRadioCheckItem(rid gd.RID, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new radio-checkable item with text [param label] and icon [param icon] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] Radio-checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
//go:nosplit
func (self class) AddIconRadioCheckItem(rid gd.RID, icon object.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_icon_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new item with text [param label] to the global menu [param rid].
Contrarily to normal binary items, multistate items can have more than two states, as defined by [param max_states]. Each press or activate of the item will increase the state by one. The default value is defined by [param default_state].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] By default, there's no indication of the current item state, it should be changed manually.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
//go:nosplit
func (self class) AddMultistateItem(rid gd.RID, label gd.String, max_states gd.Int, default_state gd.Int, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, mmm.Get(callback))
	callframe.Arg(frame, mmm.Get(key_callback))
	callframe.Arg(frame, mmm.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_multistate_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a separator between items to the global menu [param rid]. Separators also occupy an index.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) AddSeparator(rid gd.RID, index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_add_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the item with the specified [param text]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) FindItemIndexWithText(rid gd.RID, text gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_find_item_index_with_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the item with the specified [param tag]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) FindItemIndexWithTag(rid gd.RID, tag gd.Variant) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, mmm.Get(tag))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_find_item_index_with_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the item with the submenu specified by [param submenu_rid]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) FindItemIndexWithSubmenu(rid gd.RID, submenu_rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, submenu_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_find_item_index_with_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is checked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) IsItemChecked(rid gd.RID, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_is_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) IsItemCheckable(rid gd.RID, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_is_item_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) IsItemRadioCheckable(rid gd.RID, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_is_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the callback of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemCallback(ctx gd.Lifetime, rid gd.RID, idx gd.Int) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetItemKeyCallback(ctx gd.Lifetime, rid gd.RID, idx gd.Int) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemTag(ctx gd.Lifetime, rid gd.RID, idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemText(ctx gd.Lifetime, rid gd.RID, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the submenu ID of the item at index [param idx]. See [method add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemSubmenu(rid gd.RID, idx gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the accelerator of the item at index [param idx]. Accelerators are special combinations of keys that activate the item, no matter which control is focused.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetItemAccelerator(rid gd.RID, idx gd.Int) gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method set_item_disabled] for more info on how to disable an item.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) IsItemDisabled(rid gd.RID, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is hidden.
See [method set_item_hidden] for more info on how to hide an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) IsItemHidden(rid gd.RID, idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_is_item_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tooltip associated with the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetItemTooltip(ctx gd.Lifetime, rid gd.RID, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemState(rid gd.RID, idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of states of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemMaxStates(rid gd.RID, idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_max_states, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the icon of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemIcon(ctx gd.Lifetime, rid gd.RID, idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetItemIndentationLevel(rid gd.RID, idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_indentation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the checkstate status of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemChecked(rid gd.RID, idx gd.Int, checked bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the item at index [param idx] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemCheckable(rid gd.RID, idx gd.Int, checkable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checkable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the type of the item at the specified index [param idx] to radio button. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemRadioCheckable(rid gd.RID, idx gd.Int, checkable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checkable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. Callback is emitted when an item is pressed.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemCallback(rid gd.RID, idx gd.Int, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. The callback is emitted when an item is hovered.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemHoverCallbacks(rid gd.RID, idx gd.Int, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_hover_callbacks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. Callback is emitted when its accelerator is activated.
[b]Note:[/b] The [param key_callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemKeyCallback(rid gd.RID, idx gd.Int, key_callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(key_callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemTag(rid gd.RID, idx gd.Int, tag gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(tag))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemText(rid gd.RID, idx gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the submenu RID of the item at index [param idx]. The submenu is a global menu that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemSubmenu(rid gd.RID, idx gd.Int, submenu_rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, submenu_rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemAccelerator(rid gd.RID, idx gd.Int, keycode gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, keycode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables/disables the item at index [param idx]. When it is disabled, it can't be selected and its action can't be invoked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemDisabled(rid gd.RID, idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Hides/shows the item at index [param idx]. When it is hidden, an item does not appear in a menu and its action cannot be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemHidden(rid gd.RID, idx gd.Int, hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [String] tooltip of the item at the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemTooltip(rid gd.RID, idx gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemState(rid gd.RID, idx gd.Int, state gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, state)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets number of state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemMaxStates(rid gd.RID, idx gd.Int, max_states gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, max_states)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_max_states, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the [Texture2D] icon of the specified [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] This method is not supported by macOS Dock menu items.
*/
//go:nosplit
func (self class) SetItemIcon(rid gd.RID, idx gd.Int, icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemIndentationLevel(rid gd.RID, idx gd.Int, level gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, level)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_set_item_indentation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns number of items in the global menu [param rid].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemCount(rid gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Return [code]true[/code] is global menu is a special system menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) IsSystemMenu(rid gd.RID) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_is_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the item at index [param idx] from the global menu [param rid].
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) RemoveItem(rid gd.RID, idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all items from the global menu [param rid].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) Clear(rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NativeMenu.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsNativeMenu() Expert { return self[0].AsNativeMenu() }


//go:nosplit
func (self Simple) AsNativeMenu() Simple { return self[0].AsNativeMenu() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("NativeMenu", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Feature = classdb.NativeMenuFeature

const (
/*[NativeMenu] supports native global main menu.*/
	FeatureGlobalMenu Feature = 0
/*[NativeMenu] supports native popup menus.*/
	FeaturePopupMenu Feature = 1
/*[NativeMenu] supports menu open and close callbacks.*/
	FeatureOpenCloseCallback Feature = 2
/*[NativeMenu] supports menu item hover callback.*/
	FeatureHoverCallback Feature = 3
/*[NativeMenu] supports menu item accelerator/key callback.*/
	FeatureKeyCallback Feature = 4
)
type SystemMenus = classdb.NativeMenuSystemMenus

const (
/*Invalid special system menu ID.*/
	InvalidMenuId SystemMenus = 0
/*Global main menu ID.*/
	MainMenuId SystemMenus = 1
/*Application (first menu after "Apple" menu on macOS) menu ID.*/
	ApplicationMenuId SystemMenus = 2
/*"Window" menu ID (on macOS this menu includes standard window control items and a list of open windows).*/
	WindowMenuId SystemMenus = 3
/*"Help" menu ID (on macOS this menu includes help search bar).*/
	HelpMenuId SystemMenus = 4
/*Dock icon right-click menu ID (on macOS this menu include standard application control items and a list of open windows).*/
	DockMenuId SystemMenus = 5
)
