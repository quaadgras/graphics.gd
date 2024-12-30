package NativeMenu

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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
var self objects.NativeMenu
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.NativeMenu)
	self = *(*objects.NativeMenu)(unsafe.Pointer(&obj))
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
func GetSystemMenu(menu_id classdb.NativeMenuSystemMenus) Resource.ID {
	once.Do(singleton)
	return Resource.ID(class(self).GetSystemMenu(menu_id))
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
func CreateMenu() Resource.ID {
	once.Do(singleton)
	return Resource.ID(class(self).CreateMenu())
}

/*
Returns [code]true[/code] if [param rid] is valid global menu.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func HasMenu(rid Resource.ID) bool {
	once.Do(singleton)
	return bool(class(self).HasMenu(rid))
}

/*
Frees a global menu object created by this [NativeMenu].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func FreeMenu(rid Resource.ID) {
	once.Do(singleton)
	class(self).FreeMenu(rid)
}

/*
Returns global menu size.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetSize(rid Resource.ID) Vector2.XY {
	once.Do(singleton)
	return Vector2.XY(class(self).GetSize(rid))
}

/*
Shows the global menu at [param position] in the screen coordinates.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func Popup(rid Resource.ID, position Vector2i.XY) {
	once.Do(singleton)
	class(self).Popup(rid, gd.Vector2i(position))
}

/*
Sets the menu text layout direction from right-to-left if [param is_rtl] is [code]true[/code].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetInterfaceDirection(rid Resource.ID, is_rtl bool) {
	once.Do(singleton)
	class(self).SetInterfaceDirection(rid, is_rtl)
}

/*
Registers callable to emit after the menu is closed.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetPopupOpenCallback(rid Resource.ID, callback func()) {
	once.Do(singleton)
	class(self).SetPopupOpenCallback(rid, gd.NewCallable(callback))
}

/*
Returns global menu open callback.
b]Note:[/b] This method is implemented only on macOS.
*/
func GetPopupOpenCallback(rid Resource.ID) Callable.Any {
	once.Do(singleton)
	return Callable.Any(class(self).GetPopupOpenCallback(rid))
}

/*
Registers callable to emit when the menu is about to show.
[b]Note:[/b] The OS can simulate menu opening to track menu item changes and global shortcuts, in which case the corresponding close callback is not triggered. Use [method is_opened] to check if the menu is currently opened.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetPopupCloseCallback(rid Resource.ID, callback func()) {
	once.Do(singleton)
	class(self).SetPopupCloseCallback(rid, gd.NewCallable(callback))
}

/*
Returns global menu close callback.
b]Note:[/b] This method is implemented only on macOS.
*/
func GetPopupCloseCallback(rid Resource.ID) Callable.Any {
	once.Do(singleton)
	return Callable.Any(class(self).GetPopupCloseCallback(rid))
}

/*
Sets the minimum width of the global menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetMinimumWidth(rid Resource.ID, width Float.X) {
	once.Do(singleton)
	class(self).SetMinimumWidth(rid, gd.Float(width))
}

/*
Returns global menu minimum width.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetMinimumWidth(rid Resource.ID) Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).GetMinimumWidth(rid)))
}

/*
Returns [code]true[/code] if the menu is currently opened.
[b]Note:[/b] This method is implemented only on macOS.
*/
func IsOpened(rid Resource.ID) bool {
	once.Do(singleton)
	return bool(class(self).IsOpened(rid))
}

/*
Adds an item that will act as a submenu of the global menu [param rid]. The [param submenu_rid] argument is the RID of the global menu that will be shown when the item is clicked.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func AddSubmenuItem(rid Resource.ID, label string, submenu_rid Resource.ID) int {
	once.Do(singleton)
	return int(int(class(self).AddSubmenuItem(rid, gd.NewString(label), submenu_rid, gd.NewVariant(gd.NewVariant(([1]any{}[0]))), gd.Int(-1))))
}

/*
Adds a new item with text [param label] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
func AddItem(rid Resource.ID, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddItem(rid, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new checkable item with text [param label] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
func AddCheckItem(rid Resource.ID, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddCheckItem(rid, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new item with text [param label] and icon [param icon] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
func AddIconItem(rid Resource.ID, icon objects.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddIconItem(rid, icon, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new checkable item with text [param label] and icon [param icon] to the global menu [param rid].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On Windows, [param accelerator] and [param key_callback] are ignored.
*/
func AddIconCheckItem(rid Resource.ID, icon objects.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddIconCheckItem(rid, icon, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
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
func AddRadioCheckItem(rid Resource.ID, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddRadioCheckItem(rid, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
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
func AddIconRadioCheckItem(rid Resource.ID, icon objects.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).AddIconRadioCheckItem(rid, icon, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
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
func AddMultistateItem(rid Resource.ID, label string, max_states int, default_state int) int {
	once.Do(singleton)
	return int(int(class(self).AddMultistateItem(rid, gd.NewString(label), gd.Int(max_states), gd.Int(default_state), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a separator between items to the global menu [param rid]. Separators also occupy an index.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func AddSeparator(rid Resource.ID) int {
	once.Do(singleton)
	return int(int(class(self).AddSeparator(rid, gd.Int(-1))))
}

/*
Returns the index of the item with the specified [param text]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func FindItemIndexWithText(rid Resource.ID, text string) int {
	once.Do(singleton)
	return int(int(class(self).FindItemIndexWithText(rid, gd.NewString(text))))
}

/*
Returns the index of the item with the specified [param tag]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func FindItemIndexWithTag(rid Resource.ID, tag any) int {
	once.Do(singleton)
	return int(int(class(self).FindItemIndexWithTag(rid, gd.NewVariant(tag))))
}

/*
Returns the index of the item with the submenu specified by [param submenu_rid]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func FindItemIndexWithSubmenu(rid Resource.ID, submenu_rid Resource.ID) int {
	once.Do(singleton)
	return int(int(class(self).FindItemIndexWithSubmenu(rid, submenu_rid)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is checked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func IsItemChecked(rid Resource.ID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemChecked(rid, gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func IsItemCheckable(rid Resource.ID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemCheckable(rid, gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func IsItemRadioCheckable(rid Resource.ID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemRadioCheckable(rid, gd.Int(idx)))
}

/*
Returns the callback of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemCallback(rid Resource.ID, idx int) Callable.Any {
	once.Do(singleton)
	return Callable.Any(class(self).GetItemCallback(rid, gd.Int(idx)))
}

/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetItemKeyCallback(rid Resource.ID, idx int) Callable.Any {
	once.Do(singleton)
	return Callable.Any(class(self).GetItemKeyCallback(rid, gd.Int(idx)))
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemTag(rid Resource.ID, idx int) any {
	once.Do(singleton)
	return any(class(self).GetItemTag(rid, gd.Int(idx)).Interface())
}

/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemText(rid Resource.ID, idx int) string {
	once.Do(singleton)
	return string(class(self).GetItemText(rid, gd.Int(idx)).String())
}

/*
Returns the submenu ID of the item at index [param idx]. See [method add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemSubmenu(rid Resource.ID, idx int) Resource.ID {
	once.Do(singleton)
	return Resource.ID(class(self).GetItemSubmenu(rid, gd.Int(idx)))
}

/*
Returns the accelerator of the item at index [param idx]. Accelerators are special combinations of keys that activate the item, no matter which control is focused.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetItemAccelerator(rid Resource.ID, idx int) Key {
	once.Do(singleton)
	return Key(class(self).GetItemAccelerator(rid, gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method set_item_disabled] for more info on how to disable an item.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func IsItemDisabled(rid Resource.ID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemDisabled(rid, gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is hidden.
See [method set_item_hidden] for more info on how to hide an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
func IsItemHidden(rid Resource.ID, idx int) bool {
	once.Do(singleton)
	return bool(class(self).IsItemHidden(rid, gd.Int(idx)))
}

/*
Returns the tooltip associated with the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetItemTooltip(rid Resource.ID, idx int) string {
	once.Do(singleton)
	return string(class(self).GetItemTooltip(rid, gd.Int(idx)).String())
}

/*
Returns the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemState(rid Resource.ID, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GetItemState(rid, gd.Int(idx))))
}

/*
Returns number of states of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemMaxStates(rid Resource.ID, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GetItemMaxStates(rid, gd.Int(idx))))
}

/*
Returns the icon of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemIcon(rid Resource.ID, idx int) objects.Texture2D {
	once.Do(singleton)
	return objects.Texture2D(class(self).GetItemIcon(rid, gd.Int(idx)))
}

/*
Returns the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GetItemIndentationLevel(rid Resource.ID, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GetItemIndentationLevel(rid, gd.Int(idx))))
}

/*
Sets the checkstate status of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemChecked(rid Resource.ID, idx int, checked bool) {
	once.Do(singleton)
	class(self).SetItemChecked(rid, gd.Int(idx), checked)
}

/*
Sets whether the item at index [param idx] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemCheckable(rid Resource.ID, idx int, checkable bool) {
	once.Do(singleton)
	class(self).SetItemCheckable(rid, gd.Int(idx), checkable)
}

/*
Sets the type of the item at the specified index [param idx] to radio button. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemRadioCheckable(rid Resource.ID, idx int, checkable bool) {
	once.Do(singleton)
	class(self).SetItemRadioCheckable(rid, gd.Int(idx), checkable)
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when an item is pressed.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemCallback(rid Resource.ID, idx int, callback func(tag any)) {
	once.Do(singleton)
	class(self).SetItemCallback(rid, gd.Int(idx), gd.NewCallable(callback))
}

/*
Sets the callback of the item at index [param idx]. The callback is emitted when an item is hovered.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemHoverCallbacks(rid Resource.ID, idx int, callback func(tag any)) {
	once.Do(singleton)
	class(self).SetItemHoverCallbacks(rid, gd.Int(idx), gd.NewCallable(callback))
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when its accelerator is activated.
[b]Note:[/b] The [param key_callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemKeyCallback(rid Resource.ID, idx int, key_callback func(tag any)) {
	once.Do(singleton)
	class(self).SetItemKeyCallback(rid, gd.Int(idx), gd.NewCallable(key_callback))
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemTag(rid Resource.ID, idx int, tag any) {
	once.Do(singleton)
	class(self).SetItemTag(rid, gd.Int(idx), gd.NewVariant(tag))
}

/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemText(rid Resource.ID, idx int, text string) {
	once.Do(singleton)
	class(self).SetItemText(rid, gd.Int(idx), gd.NewString(text))
}

/*
Sets the submenu RID of the item at index [param idx]. The submenu is a global menu that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemSubmenu(rid Resource.ID, idx int, submenu_rid Resource.ID) {
	once.Do(singleton)
	class(self).SetItemSubmenu(rid, gd.Int(idx), submenu_rid)
}

/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemAccelerator(rid Resource.ID, idx int, keycode Key) {
	once.Do(singleton)
	class(self).SetItemAccelerator(rid, gd.Int(idx), keycode)
}

/*
Enables/disables the item at index [param idx]. When it is disabled, it can't be selected and its action can't be invoked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemDisabled(rid Resource.ID, idx int, disabled bool) {
	once.Do(singleton)
	class(self).SetItemDisabled(rid, gd.Int(idx), disabled)
}

/*
Hides/shows the item at index [param idx]. When it is hidden, an item does not appear in a menu and its action cannot be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemHidden(rid Resource.ID, idx int, hidden bool) {
	once.Do(singleton)
	class(self).SetItemHidden(rid, gd.Int(idx), hidden)
}

/*
Sets the [String] tooltip of the item at the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemTooltip(rid Resource.ID, idx int, tooltip string) {
	once.Do(singleton)
	class(self).SetItemTooltip(rid, gd.Int(idx), gd.NewString(tooltip))
}

/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemState(rid Resource.ID, idx int, state int) {
	once.Do(singleton)
	class(self).SetItemState(rid, gd.Int(idx), gd.Int(state))
}

/*
Sets number of state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func SetItemMaxStates(rid Resource.ID, idx int, max_states int) {
	once.Do(singleton)
	class(self).SetItemMaxStates(rid, gd.Int(idx), gd.Int(max_states))
}

/*
Replaces the [Texture2D] icon of the specified [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] This method is not supported by macOS Dock menu items.
*/
func SetItemIcon(rid Resource.ID, idx int, icon objects.Texture2D) {
	once.Do(singleton)
	class(self).SetItemIcon(rid, gd.Int(idx), icon)
}

/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func SetItemIndentationLevel(rid Resource.ID, idx int, level int) {
	once.Do(singleton)
	class(self).SetItemIndentationLevel(rid, gd.Int(idx), gd.Int(level))
}

/*
Returns number of items in the global menu [param rid].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetItemCount(rid Resource.ID) int {
	once.Do(singleton)
	return int(int(class(self).GetItemCount(rid)))
}

/*
Return [code]true[/code] is global menu is a special system menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func IsSystemMenu(rid Resource.ID) bool {
	once.Do(singleton)
	return bool(class(self).IsSystemMenu(rid))
}

/*
Removes the item at index [param idx] from the global menu [param rid].
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func RemoveItem(rid Resource.ID, idx int) {
	once.Do(singleton)
	class(self).RemoveItem(rid, gd.Int(idx))
}

/*
Removes all items from the global menu [param rid].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func Clear(rid Resource.ID) {
	once.Do(singleton)
	class(self).Clear(rid)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.NativeMenu

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

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
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) FreeMenu(rid gd.RID) {
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
func (self class) Popup(rid gd.RID, position gd.Vector2i) {
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
func (self class) SetInterfaceDirection(rid gd.RID, is_rtl bool) {
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
func (self class) SetPopupOpenCallback(rid gd.RID, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(callback))
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
	var ret = pointers.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}

/*
Registers callable to emit when the menu is about to show.
[b]Note:[/b] The OS can simulate menu opening to track menu item changes and global shortcuts, in which case the corresponding close callback is not triggered. Use [method is_opened] to check if the menu is currently opened.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetPopupCloseCallback(rid gd.RID, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(callback))
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
	var ret = pointers.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the minimum width of the global menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetMinimumWidth(rid gd.RID, width gd.Float) {
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
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, submenu_rid)
	callframe.Arg(frame, pointers.Get(tag))
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
func (self class) AddItem(rid gd.RID, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
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
func (self class) AddCheckItem(rid gd.RID, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
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
func (self class) AddIconItem(rid gd.RID, icon objects.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
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
func (self class) AddIconCheckItem(rid gd.RID, icon objects.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
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
func (self class) AddRadioCheckItem(rid gd.RID, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
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
func (self class) AddIconRadioCheckItem(rid gd.RID, icon objects.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
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
func (self class) AddMultistateItem(rid gd.RID, label gd.String, max_states gd.Int, default_state gd.Int, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
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
	callframe.Arg(frame, pointers.Get(text))
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
	callframe.Arg(frame, pointers.Get(tag))
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
	var ret = pointers.New[gd.Callable](r_ret.Get())
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
	var ret = pointers.New[gd.Callable](r_ret.Get())
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
	var ret = pointers.New[gd.Variant](r_ret.Get())
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
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) GetItemAccelerator(rid gd.RID, idx gd.Int) Key {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Key](frame)
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
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) GetItemIcon(rid gd.RID, idx gd.Int) objects.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
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
func (self class) SetItemChecked(rid gd.RID, idx gd.Int, checked bool) {
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
func (self class) SetItemCheckable(rid gd.RID, idx gd.Int, checkable bool) {
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
func (self class) SetItemRadioCheckable(rid gd.RID, idx gd.Int, checkable bool) {
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
func (self class) SetItemCallback(rid gd.RID, idx gd.Int, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(callback))
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
func (self class) SetItemHoverCallbacks(rid gd.RID, idx gd.Int, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(callback))
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
func (self class) SetItemKeyCallback(rid gd.RID, idx gd.Int, key_callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(key_callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemTag(rid gd.RID, idx gd.Int, tag gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(tag))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemText(rid gd.RID, idx gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the submenu RID of the item at index [param idx]. The submenu is a global menu that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemSubmenu(rid gd.RID, idx gd.Int, submenu_rid gd.RID) {
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
func (self class) SetItemAccelerator(rid gd.RID, idx gd.Int, keycode Key) {
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
func (self class) SetItemDisabled(rid gd.RID, idx gd.Int, disabled bool) {
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
func (self class) SetItemHidden(rid gd.RID, idx gd.Int, hidden bool) {
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
func (self class) SetItemTooltip(rid gd.RID, idx gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) SetItemState(rid gd.RID, idx gd.Int, state gd.Int) {
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
func (self class) SetItemMaxStates(rid gd.RID, idx gd.Int, max_states gd.Int) {
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
func (self class) SetItemIcon(rid gd.RID, idx gd.Int, icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) SetItemIndentationLevel(rid gd.RID, idx gd.Int, level gd.Int) {
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
func (self class) RemoveItem(rid gd.RID, idx gd.Int) {
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
func (self class) Clear(rid gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NativeMenu.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("NativeMenu", func(ptr gd.Object) any { return classdb.NativeMenu(ptr) })
}

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

type Key int

const (
	/*Enum value which doesn't correspond to any key. This is used to initialize [enum Key] properties with a generic state.*/
	KeyNone Key = 0
	/*Keycodes with this bit applied are non-printable.*/
	KeySpecial Key = 4194304
	/*Escape key.*/
	KeyEscape Key = 4194305
	/*Tab key.*/
	KeyTab Key = 4194306
	/*Shift + Tab key.*/
	KeyBacktab Key = 4194307
	/*Backspace key.*/
	KeyBackspace Key = 4194308
	/*Return key (on the main keyboard).*/
	KeyEnter Key = 4194309
	/*Enter key on the numeric keypad.*/
	KeyKpEnter Key = 4194310
	/*Insert key.*/
	KeyInsert Key = 4194311
	/*Delete key.*/
	KeyDelete Key = 4194312
	/*Pause key.*/
	KeyPause Key = 4194313
	/*Print Screen key.*/
	KeyPrint Key = 4194314
	/*System Request key.*/
	KeySysreq Key = 4194315
	/*Clear key.*/
	KeyClear Key = 4194316
	/*Home key.*/
	KeyHome Key = 4194317
	/*End key.*/
	KeyEnd Key = 4194318
	/*Left arrow key.*/
	KeyLeft Key = 4194319
	/*Up arrow key.*/
	KeyUp Key = 4194320
	/*Right arrow key.*/
	KeyRight Key = 4194321
	/*Down arrow key.*/
	KeyDown Key = 4194322
	/*Page Up key.*/
	KeyPageup Key = 4194323
	/*Page Down key.*/
	KeyPagedown Key = 4194324
	/*Shift key.*/
	KeyShift Key = 4194325
	/*Control key.*/
	KeyCtrl Key = 4194326
	/*Meta key.*/
	KeyMeta Key = 4194327
	/*Alt key.*/
	KeyAlt Key = 4194328
	/*Caps Lock key.*/
	KeyCapslock Key = 4194329
	/*Num Lock key.*/
	KeyNumlock Key = 4194330
	/*Scroll Lock key.*/
	KeyScrolllock Key = 4194331
	/*F1 key.*/
	KeyF1 Key = 4194332
	/*F2 key.*/
	KeyF2 Key = 4194333
	/*F3 key.*/
	KeyF3 Key = 4194334
	/*F4 key.*/
	KeyF4 Key = 4194335
	/*F5 key.*/
	KeyF5 Key = 4194336
	/*F6 key.*/
	KeyF6 Key = 4194337
	/*F7 key.*/
	KeyF7 Key = 4194338
	/*F8 key.*/
	KeyF8 Key = 4194339
	/*F9 key.*/
	KeyF9 Key = 4194340
	/*F10 key.*/
	KeyF10 Key = 4194341
	/*F11 key.*/
	KeyF11 Key = 4194342
	/*F12 key.*/
	KeyF12 Key = 4194343
	/*F13 key.*/
	KeyF13 Key = 4194344
	/*F14 key.*/
	KeyF14 Key = 4194345
	/*F15 key.*/
	KeyF15 Key = 4194346
	/*F16 key.*/
	KeyF16 Key = 4194347
	/*F17 key.*/
	KeyF17 Key = 4194348
	/*F18 key.*/
	KeyF18 Key = 4194349
	/*F19 key.*/
	KeyF19 Key = 4194350
	/*F20 key.*/
	KeyF20 Key = 4194351
	/*F21 key.*/
	KeyF21 Key = 4194352
	/*F22 key.*/
	KeyF22 Key = 4194353
	/*F23 key.*/
	KeyF23 Key = 4194354
	/*F24 key.*/
	KeyF24 Key = 4194355
	/*F25 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF25 Key = 4194356
	/*F26 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF26 Key = 4194357
	/*F27 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF27 Key = 4194358
	/*F28 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF28 Key = 4194359
	/*F29 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF29 Key = 4194360
	/*F30 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF30 Key = 4194361
	/*F31 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF31 Key = 4194362
	/*F32 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF32 Key = 4194363
	/*F33 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF33 Key = 4194364
	/*F34 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF34 Key = 4194365
	/*F35 key. Only supported on macOS and Linux due to a Windows limitation.*/
	KeyF35 Key = 4194366
	/*Multiply (*) key on the numeric keypad.*/
	KeyKpMultiply Key = 4194433
	/*Divide (/) key on the numeric keypad.*/
	KeyKpDivide Key = 4194434
	/*Subtract (-) key on the numeric keypad.*/
	KeyKpSubtract Key = 4194435
	/*Period (.) key on the numeric keypad.*/
	KeyKpPeriod Key = 4194436
	/*Add (+) key on the numeric keypad.*/
	KeyKpAdd Key = 4194437
	/*Number 0 on the numeric keypad.*/
	KeyKp0 Key = 4194438
	/*Number 1 on the numeric keypad.*/
	KeyKp1 Key = 4194439
	/*Number 2 on the numeric keypad.*/
	KeyKp2 Key = 4194440
	/*Number 3 on the numeric keypad.*/
	KeyKp3 Key = 4194441
	/*Number 4 on the numeric keypad.*/
	KeyKp4 Key = 4194442
	/*Number 5 on the numeric keypad.*/
	KeyKp5 Key = 4194443
	/*Number 6 on the numeric keypad.*/
	KeyKp6 Key = 4194444
	/*Number 7 on the numeric keypad.*/
	KeyKp7 Key = 4194445
	/*Number 8 on the numeric keypad.*/
	KeyKp8 Key = 4194446
	/*Number 9 on the numeric keypad.*/
	KeyKp9 Key = 4194447
	/*Context menu key.*/
	KeyMenu Key = 4194370
	/*Hyper key. (On Linux/X11 only).*/
	KeyHyper Key = 4194371
	/*Help key.*/
	KeyHelp Key = 4194373
	/*Media back key. Not to be confused with the Back button on an Android device.*/
	KeyBack Key = 4194376
	/*Media forward key.*/
	KeyForward Key = 4194377
	/*Media stop key.*/
	KeyStop Key = 4194378
	/*Media refresh key.*/
	KeyRefresh Key = 4194379
	/*Volume down key.*/
	KeyVolumedown Key = 4194380
	/*Mute volume key.*/
	KeyVolumemute Key = 4194381
	/*Volume up key.*/
	KeyVolumeup Key = 4194382
	/*Media play key.*/
	KeyMediaplay Key = 4194388
	/*Media stop key.*/
	KeyMediastop Key = 4194389
	/*Previous song key.*/
	KeyMediaprevious Key = 4194390
	/*Next song key.*/
	KeyMedianext Key = 4194391
	/*Media record key.*/
	KeyMediarecord Key = 4194392
	/*Home page key.*/
	KeyHomepage Key = 4194393
	/*Favorites key.*/
	KeyFavorites Key = 4194394
	/*Search key.*/
	KeySearch Key = 4194395
	/*Standby key.*/
	KeyStandby Key = 4194396
	/*Open URL / Launch Browser key.*/
	KeyOpenurl Key = 4194397
	/*Launch Mail key.*/
	KeyLaunchmail Key = 4194398
	/*Launch Media key.*/
	KeyLaunchmedia Key = 4194399
	/*Launch Shortcut 0 key.*/
	KeyLaunch0 Key = 4194400
	/*Launch Shortcut 1 key.*/
	KeyLaunch1 Key = 4194401
	/*Launch Shortcut 2 key.*/
	KeyLaunch2 Key = 4194402
	/*Launch Shortcut 3 key.*/
	KeyLaunch3 Key = 4194403
	/*Launch Shortcut 4 key.*/
	KeyLaunch4 Key = 4194404
	/*Launch Shortcut 5 key.*/
	KeyLaunch5 Key = 4194405
	/*Launch Shortcut 6 key.*/
	KeyLaunch6 Key = 4194406
	/*Launch Shortcut 7 key.*/
	KeyLaunch7 Key = 4194407
	/*Launch Shortcut 8 key.*/
	KeyLaunch8 Key = 4194408
	/*Launch Shortcut 9 key.*/
	KeyLaunch9 Key = 4194409
	/*Launch Shortcut A key.*/
	KeyLauncha Key = 4194410
	/*Launch Shortcut B key.*/
	KeyLaunchb Key = 4194411
	/*Launch Shortcut C key.*/
	KeyLaunchc Key = 4194412
	/*Launch Shortcut D key.*/
	KeyLaunchd Key = 4194413
	/*Launch Shortcut E key.*/
	KeyLaunche Key = 4194414
	/*Launch Shortcut F key.*/
	KeyLaunchf Key = 4194415
	/*"Globe" key on Mac / iPad keyboard.*/
	KeyGlobe Key = 4194416
	/*"On-screen keyboard" key on iPad keyboard.*/
	KeyKeyboard Key = 4194417
	/*英数 key on Mac keyboard.*/
	KeyJisEisu Key = 4194418
	/*かな key on Mac keyboard.*/
	KeyJisKana Key = 4194419
	/*Unknown key.*/
	KeyUnknown Key = 8388607
	/*Space key.*/
	KeySpace Key = 32
	/*! key.*/
	KeyExclam Key = 33
	/*" key.*/
	KeyQuotedbl Key = 34
	/*# key.*/
	KeyNumbersign Key = 35
	/*$ key.*/
	KeyDollar Key = 36
	/*% key.*/
	KeyPercent Key = 37
	/*& key.*/
	KeyAmpersand Key = 38
	/*' key.*/
	KeyApostrophe Key = 39
	/*( key.*/
	KeyParenleft Key = 40
	/*) key.*/
	KeyParenright Key = 41
	/** key.*/
	KeyAsterisk Key = 42
	/*+ key.*/
	KeyPlus Key = 43
	/*, key.*/
	KeyComma Key = 44
	/*- key.*/
	KeyMinus Key = 45
	/*. key.*/
	KeyPeriod Key = 46
	/*/ key.*/
	KeySlash Key = 47
	/*Number 0 key.*/
	Key0 Key = 48
	/*Number 1 key.*/
	Key1 Key = 49
	/*Number 2 key.*/
	Key2 Key = 50
	/*Number 3 key.*/
	Key3 Key = 51
	/*Number 4 key.*/
	Key4 Key = 52
	/*Number 5 key.*/
	Key5 Key = 53
	/*Number 6 key.*/
	Key6 Key = 54
	/*Number 7 key.*/
	Key7 Key = 55
	/*Number 8 key.*/
	Key8 Key = 56
	/*Number 9 key.*/
	Key9 Key = 57
	/*: key.*/
	KeyColon Key = 58
	/*; key.*/
	KeySemicolon Key = 59
	/*< key.*/
	KeyLess Key = 60
	/*= key.*/
	KeyEqual Key = 61
	/*> key.*/
	KeyGreater Key = 62
	/*? key.*/
	KeyQuestion Key = 63
	/*@ key.*/
	KeyAt Key = 64
	/*A key.*/
	KeyA Key = 65
	/*B key.*/
	KeyB Key = 66
	/*C key.*/
	KeyC Key = 67
	/*D key.*/
	KeyD Key = 68
	/*E key.*/
	KeyE Key = 69
	/*F key.*/
	KeyF Key = 70
	/*G key.*/
	KeyG Key = 71
	/*H key.*/
	KeyH Key = 72
	/*I key.*/
	KeyI Key = 73
	/*J key.*/
	KeyJ Key = 74
	/*K key.*/
	KeyK Key = 75
	/*L key.*/
	KeyL Key = 76
	/*M key.*/
	KeyM Key = 77
	/*N key.*/
	KeyN Key = 78
	/*O key.*/
	KeyO Key = 79
	/*P key.*/
	KeyP Key = 80
	/*Q key.*/
	KeyQ Key = 81
	/*R key.*/
	KeyR Key = 82
	/*S key.*/
	KeyS Key = 83
	/*T key.*/
	KeyT Key = 84
	/*U key.*/
	KeyU Key = 85
	/*V key.*/
	KeyV Key = 86
	/*W key.*/
	KeyW Key = 87
	/*X key.*/
	KeyX Key = 88
	/*Y key.*/
	KeyY Key = 89
	/*Z key.*/
	KeyZ Key = 90
	/*[ key.*/
	KeyBracketleft Key = 91
	/*\ key.*/
	KeyBackslash Key = 92
	/*] key.*/
	KeyBracketright Key = 93
	/*^ key.*/
	KeyAsciicircum Key = 94
	/*_ key.*/
	KeyUnderscore Key = 95
	/*` key.*/
	KeyQuoteleft Key = 96
	/*{ key.*/
	KeyBraceleft Key = 123
	/*| key.*/
	KeyBar Key = 124
	/*} key.*/
	KeyBraceright Key = 125
	/*~ key.*/
	KeyAsciitilde Key = 126
	/*¥ key.*/
	KeyYen Key = 165
	/*§ key.*/
	KeySection Key = 167
)
