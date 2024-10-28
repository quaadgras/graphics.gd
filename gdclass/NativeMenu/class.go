package NativeMenu

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

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
var self gdclass.NativeMenu
var once sync.Once
func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.NativeMenu)
	self = *(*gdclass.NativeMenu)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if the specified [param feature] is supported by the current [NativeMenu], [code]false[/code] otherwise.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func HasFeature(feature classdb.NativeMenuFeature) bool {
	once.Do(singleton)
	return bool(class(self).HasFeature(feature))
}

/*
Returns [code]true[/code] if a special system menu is supported.
[b]Note:[/b] This method is implemented only on macOS.
*/
func HasSystemMenu(menu_id classdb.NativeMenuSystemMenus) bool {
	once.Do(singleton)
	return bool(class(self).HasSystemMenu(menu_id))
}

/*
Returns RID of a special system menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetSystemMenu(menu_id classdb.NativeMenuSystemMenus) gd.RID {
	once.Do(singleton)
	return gd.RID(class(self).GetSystemMenu(menu_id))
}

/*
Returns readable name of a special system menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetSystemMenuName(menu_id classdb.NativeMenuSystemMenus) string {
	once.Do(singleton)
	return string(class(self).GetSystemMenuName(menu_id).String())
}

/*
Creates a new global menu object.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func CreateMenu() gd.RID {
	once.Do(singleton)
	return gd.RID(class(self).CreateMenu())
}

/*
Returns [code]true[/code] if [param rid] is valid global menu.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func HasMenu(rid gd.RID) bool {
	once.Do(singleton)
	return bool(class(self).HasMenu(rid))
}

/*
Frees a global menu object created by this [NativeMenu].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func FreeMenu(rid gd.RID) {
	once.Do(singleton)
	class(self).FreeMenu(rid)
}

/*
Returns global menu size.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetSize(rid gd.RID) gd.Vector2 {
	once.Do(singleton)
	return gd.Vector2(class(self).GetSize(rid))
}

/*
Shows the global menu at [param position] in the screen coordinates.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func Popup(rid gd.RID, position gd.Vector2i) {
	once.Do(singleton)
	class(self).Popup(rid, position)
}

/*
Sets the menu text layout direction from right-to-left if [param is_rtl] is [code]true[/code].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetInterfaceDirection(rid gd.RID, is_rtl bool) {
	once.Do(singleton)
	class(self).SetInterfaceDirection(rid, is_rtl)
}

/*
Registers callable to emit after the menu is closed.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetPopupOpenCallback(rid gd.RID, callback gd.Callable) {
	once.Do(singleton)
	class(self).SetPopupOpenCallback(rid, callback)
}

/*
Returns global menu open callback.
b]Note:[/b] This method is implemented only on macOS.
*/
func GetPopupOpenCallback(rid gd.RID) gd.Callable {
	once.Do(singleton)
	return gd.Callable(class(self).GetPopupOpenCallback(rid))
}

/*
Registers callable to emit when the menu is about to show.
[b]Note:[/b] The OS can simulate menu opening to track menu item changes and global shortcuts, in which case the corresponding close callback is not triggered. Use [method is_opened] to check if the menu is currently opened.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetPopupCloseCallback(rid gd.RID, callback gd.Callable) {
	once.Do(singleton)
	class(self).SetPopupCloseCallback(rid, callback)
}

/*
Returns global menu close callback.
b]Note:[/b] This method is implemented only on macOS.
*/
func GetPopupCloseCallback(rid gd.RID) gd.Callable {
	once.Do(singleton)
	return gd.Callable(class(self).GetPopupCloseCallback(rid))
}

/*
Sets the minimum width of the global menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetMinimumWidth(rid gd.RID, width float64) {
	once.Do(singleton)
	class(self).SetMinimumWidth(rid, gd.Float(width))
}

/*
Returns global menu minimum width.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetMinimumWidth(rid gd.RID) float64 {
	once.Do(singleton)
	return float64(float64(class(self).GetMinimumWidth(rid)))
}

/*
Returns [code]true[/code] if the menu is currently opened.
[b]Note:[/b] This method is implemented only on macOS.
*/
func IsOpened(rid gd.RID) bool {
	once.Do(singleton)
	return bool(class(self).IsOpened(rid))
}

/*
Adds an item that will act as a submenu of the global menu [param rid]. The [param submenu_rid] argument is the RID of the global menu that will be shown when the item is clicked.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func AddSubmenuItem(rid gd.RID, label string, submenu_rid gd.RID) int {
	once.Do(singleton)
	return int(int(class(self).AddSubmenuItem(rid, gd.NewString(label), submenu_rid, gd.NewVariant(([1]gd.Variant{}[0])), gd.Int(-1))))
}

/*
Adds a new item with text [param label] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
func AddItem(rid gd.RID, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddItem(rid, gd.NewString(label), ([1]gd.Callable{}[0]), ([1]gd.Callable{}[0]), gd.NewVariant(([1]gd.Variant{}[0])), 0, gd.Int(-1))))
}

/*
Adds a new checkable item with text [param label] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
func AddCheckItem(rid gd.RID, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddCheckItem(rid, gd.NewString(label), ([1]gd.Callable{}[0]), ([1]gd.Callable{}[0]), gd.NewVariant(([1]gd.Variant{}[0])), 0, gd.Int(-1))))
}

/*
Adds a new item with text [param label] and icon [param icon] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
func AddIconItem(rid gd.RID, icon gdclass.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddIconItem(rid, icon, gd.NewString(label), ([1]gd.Callable{}[0]), ([1]gd.Callable{}[0]), gd.NewVariant(([1]gd.Variant{}[0])), 0, gd.Int(-1))))
}

/*
Adds a new checkable item with text [param label] and icon [param icon] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
func AddIconCheckItem(rid gd.RID, icon gdclass.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddIconCheckItem(rid, icon, gd.NewString(label), ([1]gd.Callable{}[0]), ([1]gd.Callable{}[0]), gd.NewVariant(([1]gd.Variant{}[0])), 0, gd.Int(-1))))
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
func AddRadioCheckItem(rid gd.RID, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddRadioCheckItem(rid, gd.NewString(label), ([1]gd.Callable{}[0]), ([1]gd.Callable{}[0]), gd.NewVariant(([1]gd.Variant{}[0])), 0, gd.Int(-1))))
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
func AddIconRadioCheckItem(rid gd.RID, icon gdclass.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddIconRadioCheckItem(rid, icon, gd.NewString(label), ([1]gd.Callable{}[0]), ([1]gd.Callable{}[0]), gd.NewVariant(([1]gd.Variant{}[0])), 0, gd.Int(-1))))
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
func AddMultistateItem(rid gd.RID, label string, max_states int, default_state int) int {
	once.Do(singleton)
	return int(int(class(self).AddMultistateItem(rid, gd.NewString(label), gd.Int(max_states), gd.Int(default_state), ([1]gd.Callable{}[0]), ([1]gd.Callable{}[0]), gd.NewVariant(([1]gd.Variant{}[0])), 0, gd.Int(-1))))
}

/*
Adds a separator between items to the global menu [param rid]. Separators also occupy an index.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func AddSeparator(rid gd.RID) int {
	once.Do(singleton)
	return int(int(class(self).AddSeparator(rid, gd.Int(-1))))
}

/*
Returns the index of the item with the specified [param text]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func FindItemIndexWithText(rid gd.RID, text string) int {
	once.Do(singleton)
	return int(int(class(self).FindItemIndexWithText(rid, gd.NewString(text))))
}

/*
Returns the index of the item with the specified [param tag]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func FindItemIndexWithTag(rid gd.RID, tag gd.Variant) int {
	once.Do(singleton)
	return int(int(class(self).FindItemIndexWithTag(rid, tag)))
}

/*
Returns the index of the item with the submenu specified by [param submenu_rid]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func FindItemIndexWithSubmenu(rid gd.RID, submenu_rid gd.RID) int {
	once.Do(singleton)
	return int(int(class(self).FindItemIndexWithSubmenu(rid, submenu_rid)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is checked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func IsItemChecked(rid gd.RID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemChecked(rid, gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func IsItemCheckable(rid gd.RID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemCheckable(rid, gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func IsItemRadioCheckable(rid gd.RID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemRadioCheckable(rid, gd.Int(idx)))
}

/*
Returns the callback of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemCallback(rid gd.RID, idx int) gd.Callable {
	once.Do(singleton)
	return gd.Callable(class(self).GetItemCallback(rid, gd.Int(idx)))
}

/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetItemKeyCallback(rid gd.RID, idx int) gd.Callable {
	once.Do(singleton)
	return gd.Callable(class(self).GetItemKeyCallback(rid, gd.Int(idx)))
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemTag(rid gd.RID, idx int) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).GetItemTag(rid, gd.Int(idx)))
}

/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemText(rid gd.RID, idx int) string {
	once.Do(singleton)
	return string(class(self).GetItemText(rid, gd.Int(idx)).String())
}

/*
Returns the submenu ID of the item at index [param idx]. See [method add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemSubmenu(rid gd.RID, idx int) gd.RID {
	once.Do(singleton)
	return gd.RID(class(self).GetItemSubmenu(rid, gd.Int(idx)))
}

/*
Returns the accelerator of the item at index [param idx]. Accelerators are special combinations of keys that activate the item, no matter which control is focused.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetItemAccelerator(rid gd.RID, idx int) gd.Key {
	once.Do(singleton)
	return gd.Key(class(self).GetItemAccelerator(rid, gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method set_item_disabled] for more info on how to disable an item.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func IsItemDisabled(rid gd.RID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemDisabled(rid, gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is hidden.
See [method set_item_hidden] for more info on how to hide an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
func IsItemHidden(rid gd.RID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemHidden(rid, gd.Int(idx)))
}

/*
Returns the tooltip associated with the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetItemTooltip(rid gd.RID, idx int) string {
	once.Do(singleton)
	return string(class(self).GetItemTooltip(rid, gd.Int(idx)).String())
}

/*
Returns the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemState(rid gd.RID, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GetItemState(rid, gd.Int(idx))))
}

/*
Returns number of states of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemMaxStates(rid gd.RID, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GetItemMaxStates(rid, gd.Int(idx))))
}

/*
Returns the icon of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemIcon(rid gd.RID, idx int) gdclass.Texture2D {
	once.Do(singleton)
	return gdclass.Texture2D(class(self).GetItemIcon(rid, gd.Int(idx)))
}

/*
Returns the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetItemIndentationLevel(rid gd.RID, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GetItemIndentationLevel(rid, gd.Int(idx))))
}

/*
Sets the checkstate status of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemChecked(rid gd.RID, idx int, checked bool) {
	once.Do(singleton)
	class(self).SetItemChecked(rid, gd.Int(idx), checked)
}

/*
Sets whether the item at index [param idx] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemCheckable(rid gd.RID, idx int, checkable bool) {
	once.Do(singleton)
	class(self).SetItemCheckable(rid, gd.Int(idx), checkable)
}

/*
Sets the type of the item at the specified index [param idx] to radio button. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemRadioCheckable(rid gd.RID, idx int, checkable bool) {
	once.Do(singleton)
	class(self).SetItemRadioCheckable(rid, gd.Int(idx), checkable)
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when an item is pressed.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemCallback(rid gd.RID, idx int, callback gd.Callable) {
	once.Do(singleton)
	class(self).SetItemCallback(rid, gd.Int(idx), callback)
}

/*
Sets the callback of the item at index [param idx]. The callback is emitted when an item is hovered.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemHoverCallbacks(rid gd.RID, idx int, callback gd.Callable) {
	once.Do(singleton)
	class(self).SetItemHoverCallbacks(rid, gd.Int(idx), callback)
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when its accelerator is activated.
[b]Note:[/b] The [param key_callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemKeyCallback(rid gd.RID, idx int, key_callback gd.Callable) {
	once.Do(singleton)
	class(self).SetItemKeyCallback(rid, gd.Int(idx), key_callback)
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemTag(rid gd.RID, idx int, tag gd.Variant) {
	once.Do(singleton)
	class(self).SetItemTag(rid, gd.Int(idx), tag)
}

/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemText(rid gd.RID, idx int, text string) {
	once.Do(singleton)
	class(self).SetItemText(rid, gd.Int(idx), gd.NewString(text))
}

/*
Sets the submenu RID of the item at index [param idx]. The submenu is a global menu that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemSubmenu(rid gd.RID, idx int, submenu_rid gd.RID) {
	once.Do(singleton)
	class(self).SetItemSubmenu(rid, gd.Int(idx), submenu_rid)
}

/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemAccelerator(rid gd.RID, idx int, keycode gd.Key) {
	once.Do(singleton)
	class(self).SetItemAccelerator(rid, gd.Int(idx), keycode)
}

/*
Enables/disables the item at index [param idx]. When it is disabled, it can't be selected and its action can't be invoked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemDisabled(rid gd.RID, idx int, disabled bool) {
	once.Do(singleton)
	class(self).SetItemDisabled(rid, gd.Int(idx), disabled)
}

/*
Hides/shows the item at index [param idx]. When it is hidden, an item does not appear in a menu and its action cannot be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemHidden(rid gd.RID, idx int, hidden bool) {
	once.Do(singleton)
	class(self).SetItemHidden(rid, gd.Int(idx), hidden)
}

/*
Sets the [String] tooltip of the item at the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemTooltip(rid gd.RID, idx int, tooltip string) {
	once.Do(singleton)
	class(self).SetItemTooltip(rid, gd.Int(idx), gd.NewString(tooltip))
}

/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemState(rid gd.RID, idx int, state int) {
	once.Do(singleton)
	class(self).SetItemState(rid, gd.Int(idx), gd.Int(state))
}

/*
Sets number of state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemMaxStates(rid gd.RID, idx int, max_states int) {
	once.Do(singleton)
	class(self).SetItemMaxStates(rid, gd.Int(idx), gd.Int(max_states))
}

/*
Replaces the [Texture2D] icon of the specified [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] This method is not supported by macOS Dock menu items.
*/
func SetItemIcon(rid gd.RID, idx int, icon gdclass.Texture2D) {
	once.Do(singleton)
	class(self).SetItemIcon(rid, gd.Int(idx), icon)
}

/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemIndentationLevel(rid gd.RID, idx int, level int) {
	once.Do(singleton)
	class(self).SetItemIndentationLevel(rid, gd.Int(idx), gd.Int(level))
}

/*
Returns number of items in the global menu [param rid].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemCount(rid gd.RID) int {
	once.Do(singleton)
	return int(int(class(self).GetItemCount(rid)))
}

/*
Return [code]true[/code] is global menu is a special system menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func IsSystemMenu(rid gd.RID) bool {
	once.Do(singleton)
	return bool(class(self).IsSystemMenu(rid))
}

/*
Removes the item at index [param idx] from the global menu [param rid].
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func RemoveItem(rid gd.RID, idx int) {
	once.Do(singleton)
	class(self).RemoveItem(rid, gd.Int(idx))
}

/*
Removes all items from the global menu [param rid].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func Clear(rid gd.RID) {
	once.Do(singleton)
	class(self).Clear(rid)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.NativeMenu
func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Returns [code]true[/code] if the specified [param feature] is supported by the current [NativeMenu], [code]false[/code] otherwise.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) HasFeature(feature classdb.NativeMenuFeature) bool {
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, menu_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_has_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, menu_id)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns readable name of a special system menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetSystemMenuName(menu_id classdb.NativeMenuSystemMenus) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, menu_id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_system_menu_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Creates a new global menu object.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) CreateMenu() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_create_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_has_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_free_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns global menu size.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetSize(rid gd.RID) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the menu text layout direction from right-to-left if [param is_rtl] is [code]true[/code].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetInterfaceDirection(rid gd.RID, is_rtl bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, is_rtl)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_interface_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers callable to emit after the menu is closed.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetPopupOpenCallback(rid gd.RID, callback gd.Callable)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_popup_open_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns global menu open callback.
b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetPopupOpenCallback(rid gd.RID) gd.Callable {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_popup_open_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Callable](r_ret.Get())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_popup_close_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns global menu close callback.
b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetPopupCloseCallback(rid gd.RID) gd.Callable {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_popup_close_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the minimum width of the global menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetMinimumWidth(rid gd.RID, width gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_minimum_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns global menu minimum width.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetMinimumWidth(rid gd.RID) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_minimum_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_is_opened, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, submenu_rid)
	callframe.Arg(frame, discreet.Get(tag))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_submenu_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, discreet.Get(callback))
	callframe.Arg(frame, discreet.Get(key_callback))
	callframe.Arg(frame, discreet.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, discreet.Get(callback))
	callframe.Arg(frame, discreet.Get(key_callback))
	callframe.Arg(frame, discreet.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AddIconItem(rid gd.RID, icon gdclass.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(icon[0])[0])
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, discreet.Get(callback))
	callframe.Arg(frame, discreet.Get(key_callback))
	callframe.Arg(frame, discreet.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AddIconCheckItem(rid gd.RID, icon gdclass.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(icon[0])[0])
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, discreet.Get(callback))
	callframe.Arg(frame, discreet.Get(key_callback))
	callframe.Arg(frame, discreet.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_icon_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, discreet.Get(callback))
	callframe.Arg(frame, discreet.Get(key_callback))
	callframe.Arg(frame, discreet.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AddIconRadioCheckItem(rid gd.RID, icon gdclass.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator gd.Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(icon[0])[0])
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, discreet.Get(callback))
	callframe.Arg(frame, discreet.Get(key_callback))
	callframe.Arg(frame, discreet.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_icon_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, discreet.Get(callback))
	callframe.Arg(frame, discreet.Get(key_callback))
	callframe.Arg(frame, discreet.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_multistate_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_add_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(text))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_find_item_index_with_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, discreet.Get(tag))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_find_item_index_with_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, submenu_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_find_item_index_with_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_is_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_is_item_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_is_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the callback of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemCallback(rid gd.RID, idx gd.Int) gd.Callable {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetItemKeyCallback(rid gd.RID, idx gd.Int) gd.Callable {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemTag(rid gd.RID, idx gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemText(rid gd.RID, idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the submenu ID of the item at index [param idx]. See [method add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemSubmenu(rid gd.RID, idx gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Key](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_is_item_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tooltip associated with the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetItemTooltip(rid gd.RID, idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemState(rid gd.RID, idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_max_states, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the icon of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemIcon(rid gd.RID, idx gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GetItemIndentationLevel(rid gd.RID, idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_indentation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the item at index [param idx] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemCheckable(rid gd.RID, idx gd.Int, checkable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checkable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the type of the item at the specified index [param idx] to radio button. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemRadioCheckable(rid gd.RID, idx gd.Int, checkable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checkable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. Callback is emitted when an item is pressed.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemCallback(rid gd.RID, idx gd.Int, callback gd.Callable)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. The callback is emitted when an item is hovered.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemHoverCallbacks(rid gd.RID, idx gd.Int, callback gd.Callable)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_hover_callbacks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the callback of the item at index [param idx]. Callback is emitted when its accelerator is activated.
[b]Note:[/b] The [param key_callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemKeyCallback(rid gd.RID, idx gd.Int, key_callback gd.Callable)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(key_callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemTag(rid gd.RID, idx gd.Int, tag gd.Variant)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(tag))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemText(rid gd.RID, idx gd.Int, text gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the submenu RID of the item at index [param idx]. The submenu is a global menu that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemSubmenu(rid gd.RID, idx gd.Int, submenu_rid gd.RID)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, submenu_rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemAccelerator(rid gd.RID, idx gd.Int, keycode gd.Key)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, keycode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables/disables the item at index [param idx]. When it is disabled, it can't be selected and its action can't be invoked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemDisabled(rid gd.RID, idx gd.Int, disabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Hides/shows the item at index [param idx]. When it is hidden, an item does not appear in a menu and its action cannot be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemHidden(rid gd.RID, idx gd.Int, hidden bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [String] tooltip of the item at the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemTooltip(rid gd.RID, idx gd.Int, tooltip gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemState(rid gd.RID, idx gd.Int, state gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, state)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets number of state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemMaxStates(rid gd.RID, idx gd.Int, max_states gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, max_states)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_max_states, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Replaces the [Texture2D] icon of the specified [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] This method is not supported by macOS Dock menu items.
*/
//go:nosplit
func (self class) SetItemIcon(rid gd.RID, idx gd.Int, icon gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemIndentationLevel(rid gd.RID, idx gd.Int, level gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, level)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_indentation_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns number of items in the global menu [param rid].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetItemCount(rid gd.RID) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_is_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all items from the global menu [param rid].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) Clear(rid gd.RID)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("NativeMenu", func(ptr gd.Object) any { return classdb.NativeMenu(ptr) })}
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
