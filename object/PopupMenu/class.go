package PopupMenu

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Popup"
import "grow.graphics/gd/object/Window"
import "grow.graphics/gd/object/Viewport"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[PopupMenu] is a modal window used to display a list of options. Useful for toolbars and context menus.
The size of a [PopupMenu] can be limited by using [member Window.max_size]. If the height of the list of items is larger than the maximum height of the [PopupMenu], a [ScrollContainer] within the popup will allow the user to scroll the contents. If no maximum size is set, or if it is set to [code]0[/code], the [PopupMenu] height will be limited by its parent rect.
All [code]set_*[/code] methods allow negative item indices, i.e. [code]-1[/code] to access the last item, [code]-2[/code] to select the second-to-last item, and so on.
[b]Incremental search:[/b] Like [ItemList] and [Tree], [PopupMenu] supports searching within the list while the control is focused. Press a key that matches the first letter of an item's name to select the first item starting with the given letter. After that point, there are two ways to perform incremental search: 1) Press the same key again before the timeout duration to select the next item starting with the same letter. 2) Press letter keys that match the rest of the word before the timeout duration to match to select the item in question directly. Both of these actions will be reset to the beginning of the list if the timeout duration has passed since the last keystroke was registered. You can adjust the timeout duration by changing [member ProjectSettings.gui/timers/incremental_search_max_interval_msec].
[b]Note:[/b] The ID values used for items are limited to 32 bits, not full 64 bits of [int]. This has a range of [code]-2^32[/code] to [code]2^32 - 1[/code], i.e. [code]-2147483648[/code] to [code]2147483647[/code].

*/
type Simple [1]classdb.PopupMenu
func (self Simple) ActivateItemByEvent(event [1]classdb.InputEvent, for_global_only bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).ActivateItemByEvent(event, for_global_only))
}
func (self Simple) SetPreferNativeMenu(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPreferNativeMenu(enabled)
}
func (self Simple) IsPreferNativeMenu() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsPreferNativeMenu())
}
func (self Simple) IsNativeMenu() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsNativeMenu())
}
func (self Simple) AddItem(label string, id int, accel gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddItem(gc.String(label), gd.Int(id), accel)
}
func (self Simple) AddIconItem(texture [1]classdb.Texture2D, label string, id int, accel gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddIconItem(texture, gc.String(label), gd.Int(id), accel)
}
func (self Simple) AddCheckItem(label string, id int, accel gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCheckItem(gc.String(label), gd.Int(id), accel)
}
func (self Simple) AddIconCheckItem(texture [1]classdb.Texture2D, label string, id int, accel gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddIconCheckItem(texture, gc.String(label), gd.Int(id), accel)
}
func (self Simple) AddRadioCheckItem(label string, id int, accel gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddRadioCheckItem(gc.String(label), gd.Int(id), accel)
}
func (self Simple) AddIconRadioCheckItem(texture [1]classdb.Texture2D, label string, id int, accel gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddIconRadioCheckItem(texture, gc.String(label), gd.Int(id), accel)
}
func (self Simple) AddMultistateItem(label string, max_states int, default_state int, id int, accel gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddMultistateItem(gc.String(label), gd.Int(max_states), gd.Int(default_state), gd.Int(id), accel)
}
func (self Simple) AddShortcut(shortcut [1]classdb.Shortcut, id int, global bool, allow_echo bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddShortcut(shortcut, gd.Int(id), global, allow_echo)
}
func (self Simple) AddIconShortcut(texture [1]classdb.Texture2D, shortcut [1]classdb.Shortcut, id int, global bool, allow_echo bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddIconShortcut(texture, shortcut, gd.Int(id), global, allow_echo)
}
func (self Simple) AddCheckShortcut(shortcut [1]classdb.Shortcut, id int, global bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCheckShortcut(shortcut, gd.Int(id), global)
}
func (self Simple) AddIconCheckShortcut(texture [1]classdb.Texture2D, shortcut [1]classdb.Shortcut, id int, global bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddIconCheckShortcut(texture, shortcut, gd.Int(id), global)
}
func (self Simple) AddRadioCheckShortcut(shortcut [1]classdb.Shortcut, id int, global bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddRadioCheckShortcut(shortcut, gd.Int(id), global)
}
func (self Simple) AddIconRadioCheckShortcut(texture [1]classdb.Texture2D, shortcut [1]classdb.Shortcut, id int, global bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddIconRadioCheckShortcut(texture, shortcut, gd.Int(id), global)
}
func (self Simple) AddSubmenuItem(label string, submenu string, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddSubmenuItem(gc.String(label), gc.String(submenu), gd.Int(id))
}
func (self Simple) AddSubmenuNodeItem(label string, submenu [1]classdb.PopupMenu, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddSubmenuNodeItem(gc.String(label), submenu, gd.Int(id))
}
func (self Simple) SetItemText(index int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemText(gd.Int(index), gc.String(text))
}
func (self Simple) SetItemTextDirection(index int, direction classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemTextDirection(gd.Int(index), direction)
}
func (self Simple) SetItemLanguage(index int, language string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemLanguage(gd.Int(index), gc.String(language))
}
func (self Simple) SetItemIcon(index int, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIcon(gd.Int(index), icon)
}
func (self Simple) SetItemIconMaxWidth(index int, width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIconMaxWidth(gd.Int(index), gd.Int(width))
}
func (self Simple) SetItemIconModulate(index int, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIconModulate(gd.Int(index), modulate)
}
func (self Simple) SetItemChecked(index int, checked bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemChecked(gd.Int(index), checked)
}
func (self Simple) SetItemId(index int, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemId(gd.Int(index), gd.Int(id))
}
func (self Simple) SetItemAccelerator(index int, accel gd.Key) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemAccelerator(gd.Int(index), accel)
}
func (self Simple) SetItemMetadata(index int, metadata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemMetadata(gd.Int(index), metadata)
}
func (self Simple) SetItemDisabled(index int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemDisabled(gd.Int(index), disabled)
}
func (self Simple) SetItemSubmenu(index int, submenu string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemSubmenu(gd.Int(index), gc.String(submenu))
}
func (self Simple) SetItemSubmenuNode(index int, submenu [1]classdb.PopupMenu) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemSubmenuNode(gd.Int(index), submenu)
}
func (self Simple) SetItemAsSeparator(index int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemAsSeparator(gd.Int(index), enable)
}
func (self Simple) SetItemAsCheckable(index int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemAsCheckable(gd.Int(index), enable)
}
func (self Simple) SetItemAsRadioCheckable(index int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemAsRadioCheckable(gd.Int(index), enable)
}
func (self Simple) SetItemTooltip(index int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemTooltip(gd.Int(index), gc.String(tooltip))
}
func (self Simple) SetItemShortcut(index int, shortcut [1]classdb.Shortcut, global bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemShortcut(gd.Int(index), shortcut, global)
}
func (self Simple) SetItemIndent(index int, indent int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIndent(gd.Int(index), gd.Int(indent))
}
func (self Simple) SetItemMultistate(index int, state int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemMultistate(gd.Int(index), gd.Int(state))
}
func (self Simple) SetItemMultistateMax(index int, max_states int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemMultistateMax(gd.Int(index), gd.Int(max_states))
}
func (self Simple) SetItemShortcutDisabled(index int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemShortcutDisabled(gd.Int(index), disabled)
}
func (self Simple) ToggleItemChecked(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ToggleItemChecked(gd.Int(index))
}
func (self Simple) ToggleItemMultistate(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ToggleItemMultistate(gd.Int(index))
}
func (self Simple) GetItemText(index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemText(gc, gd.Int(index)).String())
}
func (self Simple) GetItemTextDirection(index int) classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlTextDirection(Expert(self).GetItemTextDirection(gd.Int(index)))
}
func (self Simple) GetItemLanguage(index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemLanguage(gc, gd.Int(index)).String())
}
func (self Simple) GetItemIcon(index int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetItemIcon(gc, gd.Int(index)))
}
func (self Simple) GetItemIconMaxWidth(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemIconMaxWidth(gd.Int(index))))
}
func (self Simple) GetItemIconModulate(index int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetItemIconModulate(gd.Int(index)))
}
func (self Simple) IsItemChecked(index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemChecked(gd.Int(index)))
}
func (self Simple) GetItemId(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemId(gd.Int(index))))
}
func (self Simple) GetItemIndex(id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemIndex(gd.Int(id))))
}
func (self Simple) GetItemAccelerator(index int) gd.Key {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Key(Expert(self).GetItemAccelerator(gd.Int(index)))
}
func (self Simple) GetItemMetadata(index int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetItemMetadata(gc, gd.Int(index)))
}
func (self Simple) IsItemDisabled(index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemDisabled(gd.Int(index)))
}
func (self Simple) GetItemSubmenu(index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemSubmenu(gc, gd.Int(index)).String())
}
func (self Simple) GetItemSubmenuNode(index int) [1]classdb.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PopupMenu(Expert(self).GetItemSubmenuNode(gc, gd.Int(index)))
}
func (self Simple) IsItemSeparator(index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemSeparator(gd.Int(index)))
}
func (self Simple) IsItemCheckable(index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemCheckable(gd.Int(index)))
}
func (self Simple) IsItemRadioCheckable(index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemRadioCheckable(gd.Int(index)))
}
func (self Simple) IsItemShortcutDisabled(index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemShortcutDisabled(gd.Int(index)))
}
func (self Simple) GetItemTooltip(index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemTooltip(gc, gd.Int(index)).String())
}
func (self Simple) GetItemShortcut(index int) [1]classdb.Shortcut {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Shortcut(Expert(self).GetItemShortcut(gc, gd.Int(index)))
}
func (self Simple) GetItemIndent(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemIndent(gd.Int(index))))
}
func (self Simple) GetItemMultistateMax(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemMultistateMax(gd.Int(index))))
}
func (self Simple) GetItemMultistate(index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemMultistate(gd.Int(index))))
}
func (self Simple) SetFocusedItem(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFocusedItem(gd.Int(index))
}
func (self Simple) GetFocusedItem() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFocusedItem()))
}
func (self Simple) SetItemCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemCount(gd.Int(count))
}
func (self Simple) GetItemCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemCount()))
}
func (self Simple) ScrollToItem(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScrollToItem(gd.Int(index))
}
func (self Simple) RemoveItem(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveItem(gd.Int(index))
}
func (self Simple) AddSeparator(label string, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddSeparator(gc.String(label), gd.Int(id))
}
func (self Simple) Clear(free_submenus bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear(free_submenus)
}
func (self Simple) SetHideOnItemSelection(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHideOnItemSelection(enable)
}
func (self Simple) IsHideOnItemSelection() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHideOnItemSelection())
}
func (self Simple) SetHideOnCheckableItemSelection(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHideOnCheckableItemSelection(enable)
}
func (self Simple) IsHideOnCheckableItemSelection() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHideOnCheckableItemSelection())
}
func (self Simple) SetHideOnStateItemSelection(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHideOnStateItemSelection(enable)
}
func (self Simple) IsHideOnStateItemSelection() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHideOnStateItemSelection())
}
func (self Simple) SetSubmenuPopupDelay(seconds float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubmenuPopupDelay(gd.Float(seconds))
}
func (self Simple) GetSubmenuPopupDelay() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetSubmenuPopupDelay()))
}
func (self Simple) SetAllowSearch(allow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowSearch(allow)
}
func (self Simple) GetAllowSearch() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAllowSearch())
}
func (self Simple) IsSystemMenu() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSystemMenu())
}
func (self Simple) SetSystemMenu(system_menu_id classdb.NativeMenuSystemMenus) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSystemMenu(system_menu_id)
}
func (self Simple) GetSystemMenu() classdb.NativeMenuSystemMenus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.NativeMenuSystemMenus(Expert(self).GetSystemMenu())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PopupMenu
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Checks the provided [param event] against the [PopupMenu]'s shortcuts and accelerators, and activates the first item with matching events. If [param for_global_only] is [code]true[/code], only shortcuts and accelerators with [code]global[/code] set to [code]true[/code] will be called.
Returns [code]true[/code] if an item was successfully activated.
[b]Note:[/b] Certain [Control]s, such as [MenuButton], will call this method automatically.
*/
//go:nosplit
func (self class) ActivateItemByEvent(event object.InputEvent, for_global_only bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(event[0].AsPointer())[0])
	callframe.Arg(frame, for_global_only)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_activate_item_by_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPreferNativeMenu(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_prefer_native_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsPreferNativeMenu() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_prefer_native_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the system native menu is supported and currently used by this [PopupMenu].
*/
//go:nosplit
func (self class) IsNativeMenu() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_native_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] The provided [param id] is used only in [signal id_pressed] and [signal id_focused] signals. It's not related to the [code]index[/code] arguments in e.g. [method set_item_checked].
*/
//go:nosplit
func (self class) AddItem(label gd.String, id gd.Int, accel gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
*/
//go:nosplit
func (self class) AddIconItem(texture object.Texture2D, label gd.String, id gd.Int, accel gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new checkable item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddCheckItem(label gd.String, id gd.Int, accel gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new checkable item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddIconCheckItem(texture object.Texture2D, label gd.String, id gd.Int, accel gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_icon_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new radio check button with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddRadioCheckItem(label gd.String, id gd.Int, accel gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Same as [method add_icon_check_item], but uses a radio check button.
*/
//go:nosplit
func (self class) AddIconRadioCheckItem(texture object.Texture2D, label gd.String, id gd.Int, accel gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_icon_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new multistate item with text [param label].
Contrarily to normal binary items, multistate items can have more than two states, as defined by [param max_states]. The default value is defined by [param default_state].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Multistate items don't update their state automatically and must be done manually. See [method toggle_item_multistate], [method set_item_multistate] and [method get_item_multistate] for more info on how to control it.
Example usage:
[codeblock]
func _ready():
    add_multistate_item("Item", 3, 0)

    index_pressed.connect(func(index: int):
            toggle_item_multistate(index)
            match get_item_multistate(index):
                0:
                    print("First state")
                1:
                    print("Second state")
                2:
                    print("Third state")
        )
[/codeblock]
*/
//go:nosplit
func (self class) AddMultistateItem(label gd.String, max_states gd.Int, default_state gd.Int, id gd.Int, accel gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_multistate_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [Shortcut].
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
//go:nosplit
func (self class) AddShortcut(shortcut object.Shortcut, id gd.Int, global bool, allow_echo bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	callframe.Arg(frame, allow_echo)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
//go:nosplit
func (self class) AddIconShortcut(texture object.Texture2D, shortcut object.Shortcut, id gd.Int, global bool, allow_echo bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	callframe.Arg(frame, allow_echo)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_icon_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new checkable item and assigns the specified [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddCheckShortcut(shortcut object.Shortcut, id gd.Int, global bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new checkable item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddIconCheckShortcut(texture object.Texture2D, shortcut object.Shortcut, id gd.Int, global bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_icon_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a new radio check button and assigns a [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddRadioCheckShortcut(shortcut object.Shortcut, id gd.Int, global bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_radio_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Same as [method add_icon_check_shortcut], but uses a radio check button.
*/
//go:nosplit
func (self class) AddIconRadioCheckShortcut(texture object.Texture2D, shortcut object.Shortcut, id gd.Int, global bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_icon_radio_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. The [param submenu] argument must be the name of an existing [PopupMenu] that has been added as a child to this node. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
//go:nosplit
func (self class) AddSubmenuItem(label gd.String, submenu gd.String, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(submenu))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_submenu_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
[param submenu] must be either child of this [PopupMenu] or has no parent node (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
//go:nosplit
func (self class) AddSubmenuNodeItem(label gd.String, submenu object.PopupMenu, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, mmm.Get(submenu[0].AsPointer())[0])
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_submenu_node_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the text of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemText(index gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetItemTextDirection(index gd.Int, direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetItemLanguage(index gd.Int, language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the [Texture2D] icon of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemIcon(index gd.Int, icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the maximum allowed width of the icon for the item at the given [param index]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetItemIconMaxWidth(index gd.Int, width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a modulating [Color] of the item's icon at the given [param index].
*/
//go:nosplit
func (self class) SetItemIconModulate(index gd.Int, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the checkstate status of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemChecked(index gd.Int, checked bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param id] of the item at the given [param index].
The [param id] is used in [signal id_pressed] and [signal id_focused] signals.
*/
//go:nosplit
func (self class) SetItemId(index gd.Int, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. [param accel] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
*/
//go:nosplit
func (self class) SetItemAccelerator(index gd.Int, accel gd.Key)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_metadata], which provides a simple way of assigning context data to items.
*/
//go:nosplit
func (self class) SetItemMetadata(index gd.Int, metadata gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(metadata))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables/disables the item at the given [param index]. When it is disabled, it can't be selected and its action can't be invoked.
*/
//go:nosplit
func (self class) SetItemDisabled(index gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the submenu of the item at the given [param index]. The submenu is the name of a child [PopupMenu] node that would be shown when the item is clicked.
*/
//go:nosplit
func (self class) SetItemSubmenu(index gd.Int, submenu gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(submenu))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the submenu of the item at the given [param index]. The submenu is a [PopupMenu] node that would be shown when the item is clicked. It must either be a child of this [PopupMenu] or has no parent (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
*/
//go:nosplit
func (self class) SetItemSubmenuNode(index gd.Int, submenu object.PopupMenu)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(submenu[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_submenu_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Mark the item at the given [param index] as a separator, which means that it would be displayed as a line. If [code]false[/code], sets the type of the item to plain text.
*/
//go:nosplit
func (self class) SetItemAsSeparator(index gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_as_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the item at the given [param index] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
//go:nosplit
func (self class) SetItemAsCheckable(index gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_as_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the type of the item at the given [param index] to radio button. If [code]false[/code], sets the type of the item to plain text.
*/
//go:nosplit
func (self class) SetItemAsRadioCheckable(index gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_as_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [String] tooltip of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemTooltip(index gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a [Shortcut] for the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemShortcut(index gd.Int, shortcut object.Shortcut, global bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mmm.Get(shortcut[0].AsPointer())[0])
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the horizontal offset of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemIndent(index gd.Int, indent gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, indent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_indent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) SetItemMultistate(index gd.Int, state gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, state)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_multistate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the max states of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) SetItemMultistateMax(index gd.Int, max_states gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, max_states)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_multistate_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disables the [Shortcut] of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemShortcutDisabled(index gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_shortcut_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Toggles the check state of the item at the given [param index].
*/
//go:nosplit
func (self class) ToggleItemChecked(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_toggle_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Cycle to the next state of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) ToggleItemMultistate(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_toggle_item_multistate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the text of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemText(ctx gd.Lifetime, index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetItemTextDirection(index gd.Int) classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetItemLanguage(ctx gd.Lifetime, index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the icon of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIcon(ctx gd.Lifetime, index gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the maximum allowed width of the icon for the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIconMaxWidth(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [Color] modulating the item's icon at the given [param index].
*/
//go:nosplit
func (self class) GetItemIconModulate(index gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at the given [param index] is checked.
*/
//go:nosplit
func (self class) IsItemChecked(index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID of the item at the given [param index]. [code]id[/code] can be manually assigned, while index can not.
*/
//go:nosplit
func (self class) GetItemId(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the item containing the specified [param id]. Index is automatically assigned to each item by the engine and can not be set manually.
*/
//go:nosplit
func (self class) GetItemIndex(id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The return value is an integer which is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]). If no accelerator is defined for the specified [param index], [method get_item_accelerator] returns [code]0[/code] (corresponding to [constant @GlobalScope.KEY_NONE]).
*/
//go:nosplit
func (self class) GetItemAccelerator(index gd.Int) gd.Key {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Key](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_metadata], which provides a simple way of assigning context data to items.
*/
//go:nosplit
func (self class) GetItemMetadata(ctx gd.Lifetime, index gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at the given [param index] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method set_item_disabled] for more info on how to disable an item.
*/
//go:nosplit
func (self class) IsItemDisabled(index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the submenu name of the item at the given [param index]. See [method add_submenu_item] for more info on how to add a submenu.
*/
//go:nosplit
func (self class) GetItemSubmenu(ctx gd.Lifetime, index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the submenu of the item at the given [param index], or [code]null[/code] if no submenu was added. See [method add_submenu_node_item] for more info on how to add a submenu.
*/
//go:nosplit
func (self class) GetItemSubmenuNode(ctx gd.Lifetime, index gd.Int) object.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_submenu_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PopupMenu
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item is a separator. If it is, it will be displayed as a line. See [method add_separator] for more info on how to add a separator.
*/
//go:nosplit
func (self class) IsItemSeparator(index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_item_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at the given [param index] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] Checkable items just display a checkmark or radio button, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
//go:nosplit
func (self class) IsItemCheckable(index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_item_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at the given [param index] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
*/
//go:nosplit
func (self class) IsItemRadioCheckable(index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified item's shortcut is disabled.
*/
//go:nosplit
func (self class) IsItemShortcutDisabled(index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_item_shortcut_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tooltip associated with the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemTooltip(ctx gd.Lifetime, index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the [Shortcut] associated with the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemShortcut(ctx gd.Lifetime, index gd.Int) object.Shortcut {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Shortcut
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the horizontal offset of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIndent(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_indent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the max states of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemMultistateMax(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_multistate_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the state of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemMultistate(index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_multistate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the currently focused item as the given [param index].
Passing [code]-1[/code] as the index makes so that no item is focused.
*/
//go:nosplit
func (self class) SetFocusedItem(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_focused_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the currently focused item. Returns [code]-1[/code] if no item is focused.
*/
//go:nosplit
func (self class) GetFocusedItem() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_focused_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetItemCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetItemCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves the scroll view to make the item at the given [param index] visible.
*/
//go:nosplit
func (self class) ScrollToItem(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_scroll_to_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the item at the given [param index] from the menu.
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
*/
//go:nosplit
func (self class) RemoveItem(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a separator between items. Separators also occupy an index, which you can set by using the [param id] parameter.
A [param label] can optionally be provided, which will appear at the center of the separator.
*/
//go:nosplit
func (self class) AddSeparator(label gd.String, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_add_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all items from the [PopupMenu]. If [param free_submenus] is [code]true[/code], the submenu nodes are automatically freed.
*/
//go:nosplit
func (self class) Clear(free_submenus bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, free_submenus)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetHideOnItemSelection(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_hide_on_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHideOnItemSelection() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_hide_on_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHideOnCheckableItemSelection(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_hide_on_checkable_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHideOnCheckableItemSelection() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_hide_on_checkable_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHideOnStateItemSelection(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_hide_on_state_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHideOnStateItemSelection() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_hide_on_state_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubmenuPopupDelay(seconds gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_submenu_popup_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubmenuPopupDelay() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_submenu_popup_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowSearch(allow bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowSearch() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the menu is bound to the special system menu.
*/
//go:nosplit
func (self class) IsSystemMenu() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_is_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSystemMenu(system_menu_id classdb.NativeMenuSystemMenus)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, system_menu_id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_set_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSystemMenu() classdb.NativeMenuSystemMenus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NativeMenuSystemMenus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PopupMenu.Bind_get_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPopupMenu() Expert { return self[0].AsPopupMenu() }


//go:nosplit
func (self Simple) AsPopupMenu() Simple { return self[0].AsPopupMenu() }


//go:nosplit
func (self class) AsPopup() Popup.Expert { return self[0].AsPopup() }


//go:nosplit
func (self Simple) AsPopup() Popup.Simple { return self[0].AsPopup() }


//go:nosplit
func (self class) AsWindow() Window.Expert { return self[0].AsWindow() }


//go:nosplit
func (self Simple) AsWindow() Window.Simple { return self[0].AsWindow() }


//go:nosplit
func (self class) AsViewport() Viewport.Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Viewport.Simple { return self[0].AsViewport() }


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
func init() {classdb.Register("PopupMenu", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
