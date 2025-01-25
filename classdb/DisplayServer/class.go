// Package DisplayServer provides methods for working with DisplayServer object instances.
package DisplayServer

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Rect2i"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector3i"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[DisplayServer] handles everything related to window management. It is separated from [OS] as a single operating system may support multiple display servers.
[b]Headless mode:[/b] Starting the engine with the [code]--headless[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url] disables all rendering and window management functions. Most functions from [DisplayServer] will return dummy values in this case.
*/
var self [1]gdclass.DisplayServer
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.DisplayServer)
	self = *(*[1]gdclass.DisplayServer)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if the specified [param feature] is supported by the current [DisplayServer], [code]false[/code] otherwise.
*/
func HasFeature(feature gdclass.DisplayServerFeature) bool {
	once.Do(singleton)
	return bool(class(self).HasFeature(feature))
}

/*
Returns the name of the [DisplayServer] currently in use. Most operating systems only have a single [DisplayServer], but Linux has access to more than one [DisplayServer] (currently X11 and Wayland).
The names of built-in display servers are [code]Windows[/code], [code]macOS[/code], [code]X11[/code] (Linux), [code]Wayland[/code] (Linux), [code]Android[/code], [code]iOS[/code], [code]web[/code] (HTML5), and [code]headless[/code] (when started with the [code]--headless[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url]).
*/
func GetName() string {
	once.Do(singleton)
	return string(class(self).GetName().String())
}

/*
Sets native help system search callbacks.
[param search_callback] has the following arguments: [code]String search_string, int result_limit[/code] and return a [Dictionary] with "key, display name" pairs for the search results. Called when the user enters search terms in the [code]Help[/code] menu.
[param action_callback] has the following arguments: [code]String key[/code]. Called when the user selects a search result in the [code]Help[/code] menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func HelpSetSearchCallbacks(search_callback func(search_string string, result_limit int) map[any]any, action_callback func(key string)) {
	once.Do(singleton)
	class(self).HelpSetSearchCallbacks(gd.NewCallable(search_callback), gd.NewCallable(action_callback))
}

/*
Registers callables to emit when the menu is respectively about to show or closed. Callback methods should have zero arguments.
*/
func GlobalMenuSetPopupCallbacks(menu_root string, open_callback func(), close_callback func()) {
	once.Do(singleton)
	class(self).GlobalMenuSetPopupCallbacks(gd.NewString(menu_root), gd.NewCallable(open_callback), gd.NewCallable(close_callback))
}

/*
Adds an item that will act as a submenu of the global menu [param menu_root]. The [param submenu] argument is the ID of the global menu root that will be shown when the item is clicked.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddSubmenuItem(menu_root string, label string, submenu string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddSubmenuItem(gd.NewString(menu_root), gd.NewString(label), gd.NewString(submenu), gd.Int(-1))))
}

/*
Adds a new item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddItem(menu_root string, label string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddItem(gd.NewString(menu_root), gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new checkable item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddCheckItem(menu_root string, label string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddCheckItem(gd.NewString(menu_root), gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddIconItem(menu_root string, icon [1]gdclass.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddIconItem(gd.NewString(menu_root), icon, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new checkable item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddIconCheckItem(menu_root string, icon [1]gdclass.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddIconCheckItem(gd.NewString(menu_root), icon, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new radio-checkable item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] Radio-checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method global_menu_set_item_checked] for more info on how to control it.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddRadioCheckItem(menu_root string, label string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddRadioCheckItem(gd.NewString(menu_root), gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new radio-checkable item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] Radio-checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method global_menu_set_item_checked] for more info on how to control it.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddIconRadioCheckItem(menu_root string, icon [1]gdclass.Texture2D, label string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddIconRadioCheckItem(gd.NewString(menu_root), icon, gd.NewString(label), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a new item with text [param label] to the global menu with ID [param menu_root].
Contrarily to normal binary items, multistate items can have more than two states, as defined by [param max_states]. Each press or activate of the item will increase the state by one. The default value is defined by [param default_state].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] By default, there's no indication of the current item state, it should be changed manually.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddMultistateItem(menu_root string, label string, max_states int, default_state int) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddMultistateItem(gd.NewString(menu_root), gd.NewString(label), gd.Int(max_states), gd.Int(default_state), gd.NewCallable(nil), gd.NewCallable(nil), gd.NewVariant(gd.NewVariant(([1]any{}[0]))), 0, gd.Int(-1))))
}

/*
Adds a separator between items to the global menu with ID [param menu_root]. Separators also occupy an index.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuAddSeparator(menu_root string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddSeparator(gd.NewString(menu_root), gd.Int(-1))))
}

/*
Returns the index of the item with the specified [param text]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemIndexFromText(menu_root string, text string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemIndexFromText(gd.NewString(menu_root), gd.NewString(text))))
}

/*
Returns the index of the item with the specified [param tag]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemIndexFromTag(menu_root string, tag any) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemIndexFromTag(gd.NewString(menu_root), gd.NewVariant(tag))))
}

/*
Returns [code]true[/code] if the item at index [param idx] is checked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemChecked(menu_root string, idx int) bool {
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemChecked(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemCheckable(menu_root string, idx int) bool {
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemCheckable(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemRadioCheckable(menu_root string, idx int) bool {
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemRadioCheckable(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns the callback of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemCallback(menu_root string, idx int) Callable.Any {
	once.Do(singleton)
	return Callable.Any(class(self).GlobalMenuGetItemCallback(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemKeyCallback(menu_root string, idx int) Callable.Any {
	once.Do(singleton)
	return Callable.Any(class(self).GlobalMenuGetItemKeyCallback(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method global_menu_set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemTag(menu_root string, idx int) any {
	once.Do(singleton)
	return any(class(self).GlobalMenuGetItemTag(gd.NewString(menu_root), gd.Int(idx)).Interface())
}

/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemText(menu_root string, idx int) string {
	once.Do(singleton)
	return string(class(self).GlobalMenuGetItemText(gd.NewString(menu_root), gd.Int(idx)).String())
}

/*
Returns the submenu ID of the item at index [param idx]. See [method global_menu_add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemSubmenu(menu_root string, idx int) string {
	once.Do(singleton)
	return string(class(self).GlobalMenuGetItemSubmenu(gd.NewString(menu_root), gd.Int(idx)).String())
}

/*
Returns the accelerator of the item at index [param idx]. Accelerators are special combinations of keys that activate the item, no matter which control is focused.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemAccelerator(menu_root string, idx int) Key {
	once.Do(singleton)
	return Key(class(self).GlobalMenuGetItemAccelerator(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method global_menu_set_item_disabled] for more info on how to disable an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemDisabled(menu_root string, idx int) bool {
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemDisabled(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is hidden.
See [method global_menu_set_item_hidden] for more info on how to hide an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemHidden(menu_root string, idx int) bool {
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemHidden(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns the tooltip associated with the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemTooltip(menu_root string, idx int) string {
	once.Do(singleton)
	return string(class(self).GlobalMenuGetItemTooltip(gd.NewString(menu_root), gd.Int(idx)).String())
}

/*
Returns the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemState(menu_root string, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemState(gd.NewString(menu_root), gd.Int(idx))))
}

/*
Returns number of states of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemMaxStates(menu_root string, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemMaxStates(gd.NewString(menu_root), gd.Int(idx))))
}

/*
Returns the icon of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemIcon(menu_root string, idx int) [1]gdclass.Texture2D {
	once.Do(singleton)
	return [1]gdclass.Texture2D(class(self).GlobalMenuGetItemIcon(gd.NewString(menu_root), gd.Int(idx)))
}

/*
Returns the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemIndentationLevel(menu_root string, idx int) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemIndentationLevel(gd.NewString(menu_root), gd.Int(idx))))
}

/*
Sets the checkstate status of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemChecked(menu_root string, idx int, checked bool) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemChecked(gd.NewString(menu_root), gd.Int(idx), checked)
}

/*
Sets whether the item at index [param idx] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemCheckable(menu_root string, idx int, checkable bool) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemCheckable(gd.NewString(menu_root), gd.Int(idx), checkable)
}

/*
Sets the type of the item at the specified index [param idx] to radio button. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemRadioCheckable(menu_root string, idx int, checkable bool) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemRadioCheckable(gd.NewString(menu_root), gd.Int(idx), checkable)
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when an item is pressed.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemCallback(menu_root string, idx int, callback func(tag any)) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemCallback(gd.NewString(menu_root), gd.Int(idx), gd.NewCallable(callback))
}

/*
Sets the callback of the item at index [param idx]. The callback is emitted when an item is hovered.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemHoverCallbacks(menu_root string, idx int, callback func(tag any)) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemHoverCallbacks(gd.NewString(menu_root), gd.Int(idx), gd.NewCallable(callback))
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when its accelerator is activated.
[b]Note:[/b] The [param key_callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemKeyCallback(menu_root string, idx int, key_callback func(tag any)) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemKeyCallback(gd.NewString(menu_root), gd.Int(idx), gd.NewCallable(key_callback))
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method global_menu_get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemTag(menu_root string, idx int, tag any) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemTag(gd.NewString(menu_root), gd.Int(idx), gd.NewVariant(tag))
}

/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemText(menu_root string, idx int, text string) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemText(gd.NewString(menu_root), gd.Int(idx), gd.NewString(text))
}

/*
Sets the submenu of the item at index [param idx]. The submenu is the ID of a global menu root that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemSubmenu(menu_root string, idx int, submenu string) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemSubmenu(gd.NewString(menu_root), gd.Int(idx), gd.NewString(submenu))
}

/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemAccelerator(menu_root string, idx int, keycode Key) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemAccelerator(gd.NewString(menu_root), gd.Int(idx), keycode)
}

/*
Enables/disables the item at index [param idx]. When it is disabled, it can't be selected and its action can't be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemDisabled(menu_root string, idx int, disabled bool) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemDisabled(gd.NewString(menu_root), gd.Int(idx), disabled)
}

/*
Hides/shows the item at index [param idx]. When it is hidden, an item does not appear in a menu and its action cannot be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemHidden(menu_root string, idx int, hidden bool) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemHidden(gd.NewString(menu_root), gd.Int(idx), hidden)
}

/*
Sets the [String] tooltip of the item at the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemTooltip(menu_root string, idx int, tooltip string) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemTooltip(gd.NewString(menu_root), gd.Int(idx), gd.NewString(tooltip))
}

/*
Sets the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemState(menu_root string, idx int, state int) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemState(gd.NewString(menu_root), gd.Int(idx), gd.Int(state))
}

/*
Sets number of state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemMaxStates(menu_root string, idx int, max_states int) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemMaxStates(gd.NewString(menu_root), gd.Int(idx), gd.Int(max_states))
}

/*
Replaces the [Texture2D] icon of the specified [param idx].
[b]Note:[/b] This method is implemented only on macOS.
[b]Note:[/b] This method is not supported by macOS "_dock" menu items.
*/
func GlobalMenuSetItemIcon(menu_root string, idx int, icon [1]gdclass.Texture2D) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemIcon(gd.NewString(menu_root), gd.Int(idx), icon)
}

/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemIndentationLevel(menu_root string, idx int, level int) {
	once.Do(singleton)
	class(self).GlobalMenuSetItemIndentationLevel(gd.NewString(menu_root), gd.Int(idx), gd.Int(level))
}

/*
Returns number of items in the global menu with ID [param menu_root].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemCount(menu_root string) int {
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemCount(gd.NewString(menu_root))))
}

/*
Removes the item at index [param idx] from the global menu [param menu_root].
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuRemoveItem(menu_root string, idx int) {
	once.Do(singleton)
	class(self).GlobalMenuRemoveItem(gd.NewString(menu_root), gd.Int(idx))
}

/*
Removes all items from the global menu with ID [param menu_root].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
func GlobalMenuClear(menu_root string) {
	once.Do(singleton)
	class(self).GlobalMenuClear(gd.NewString(menu_root))
}

/*
Returns Dictionary of supported system menu IDs and names.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetSystemMenuRoots() map[any]any {
	once.Do(singleton)
	return map[any]any(gd.DictionaryAs[any, any](class(self).GlobalMenuGetSystemMenuRoots()))
}

/*
Returns [code]true[/code] if the synthesizer is generating speech, or have utterance waiting in the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsIsSpeaking() bool {
	once.Do(singleton)
	return bool(class(self).TtsIsSpeaking())
}

/*
Returns [code]true[/code] if the synthesizer is in a paused state.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsIsPaused() bool {
	once.Do(singleton)
	return bool(class(self).TtsIsPaused())
}

/*
Returns an [Array] of voice information dictionaries.
Each [Dictionary] contains two [String] entries:
- [code]name[/code] is voice name.
- [code]id[/code] is voice identifier.
- [code]language[/code] is language code in [code]lang_Variant[/code] format. The [code]lang[/code] part is a 2 or 3-letter code based on the ISO-639 standard, in lowercase. The [code skip-lint]Variant[/code] part is an engine-dependent string describing country, region or/and dialect.
Note that Godot depends on system libraries for text-to-speech functionality. These libraries are installed by default on Windows and macOS, but not on all Linux distributions. If they are not present, this method will return an empty list. This applies to both Godot users on Linux, as well as end-users on Linux running Godot games that use text-to-speech.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsGetVoices() []map[any]any {
	once.Do(singleton)
	return []map[any]any(gd.ArrayAs[[]map[any]any](class(self).TtsGetVoices()))
}

/*
Returns an [PackedStringArray] of voice identifiers for the [param language].
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsGetVoicesForLanguage(language string) []string {
	once.Do(singleton)
	return []string(class(self).TtsGetVoicesForLanguage(gd.NewString(language)).Strings())
}

/*
Adds an utterance to the queue. If [param interrupt] is [code]true[/code], the queue is cleared first.
- [param voice] identifier is one of the [code]"id"[/code] values returned by [method tts_get_voices] or one of the values returned by [method tts_get_voices_for_language].
- [param volume] ranges from [code]0[/code] (lowest) to [code]100[/code] (highest).
- [param pitch] ranges from [code]0.0[/code] (lowest) to [code]2.0[/code] (highest), [code]1.0[/code] is default pitch for the current voice.
- [param rate] ranges from [code]0.1[/code] (lowest) to [code]10.0[/code] (highest), [code]1.0[/code] is a normal speaking rate. Other values act as a percentage relative.
- [param utterance_id] is passed as a parameter to the callback functions.
[b]Note:[/b] On Windows and Linux (X11/Wayland), utterance [param text] can use SSML markup. SSML support is engine and voice dependent. If the engine does not support SSML, you should strip out all XML markup before calling [method tts_speak].
[b]Note:[/b] The granularity of pitch, rate, and volume is engine and voice dependent. Values may be truncated.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsSpeak(text string, voice string) {
	once.Do(singleton)
	class(self).TtsSpeak(gd.NewString(text), gd.NewString(voice), gd.Int(50), gd.Float(1.0), gd.Float(1.0), gd.Int(0), false)
}

/*
Puts the synthesizer into a paused state.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsPause() {
	once.Do(singleton)
	class(self).TtsPause()
}

/*
Resumes the synthesizer if it was paused.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsResume() {
	once.Do(singleton)
	class(self).TtsResume()
}

/*
Stops synthesis in progress and removes all utterances from the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Linux), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsStop() {
	once.Do(singleton)
	class(self).TtsStop()
}

/*
Adds a callback, which is called when the utterance has started, finished, canceled or reached a text boundary.
- [constant TTS_UTTERANCE_STARTED], [constant TTS_UTTERANCE_ENDED], and [constant TTS_UTTERANCE_CANCELED] callable's method should take one [int] parameter, the utterance ID.
- [constant TTS_UTTERANCE_BOUNDARY] callable's method should take two [int] parameters, the index of the character and the utterance ID.
[b]Note:[/b] The granularity of the boundary callbacks is engine dependent.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsSetUtteranceCallback(event gdclass.DisplayServerTTSUtteranceEvent, callable func(int, int)) {
	once.Do(singleton)
	class(self).TtsSetUtteranceCallback(event, gd.NewCallable(callable))
}

/*
Returns [code]true[/code] if OS supports dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
func IsDarkModeSupported() bool {
	once.Do(singleton)
	return bool(class(self).IsDarkModeSupported())
}

/*
Returns [code]true[/code] if OS is using dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
func IsDarkMode() bool {
	once.Do(singleton)
	return bool(class(self).IsDarkMode())
}

/*
Returns OS theme accent color. Returns [code]Color(0, 0, 0, 0)[/code], if accent color is unknown.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetAccentColor() Color.RGBA {
	once.Do(singleton)
	return Color.RGBA(class(self).GetAccentColor())
}

/*
Returns the OS theme base color (default control background). Returns [code]Color(0, 0, 0, 0)[/code] if the base color is unknown.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func GetBaseColor() Color.RGBA {
	once.Do(singleton)
	return Color.RGBA(class(self).GetBaseColor())
}

/*
Sets the [param callable] that should be called when system theme settings are changed. Callback method should have zero arguments.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
func SetSystemThemeChangeCallback(callable func()) {
	once.Do(singleton)
	class(self).SetSystemThemeChangeCallback(gd.NewCallable(callable))
}

/*
Sets the current mouse mode. See also [method mouse_get_mode].
*/
func MouseSetMode(mouse_mode gdclass.DisplayServerMouseMode) {
	once.Do(singleton)
	class(self).MouseSetMode(mouse_mode)
}

/*
Returns the current mouse mode. See also [method mouse_set_mode].
*/
func MouseGetMode() gdclass.DisplayServerMouseMode {
	once.Do(singleton)
	return gdclass.DisplayServerMouseMode(class(self).MouseGetMode())
}

/*
Sets the mouse cursor position to the given [param position] relative to an origin at the upper left corner of the currently focused game Window Manager window.
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS, and Linux (X11/Wayland). It has no effect on Android, iOS, and Web.
*/
func WarpMouse(position Vector2i.XY) {
	once.Do(singleton)
	class(self).WarpMouse(gd.Vector2i(position))
}

/*
Returns the mouse cursor's current position in screen coordinates.
*/
func MouseGetPosition() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).MouseGetPosition())
}

/*
Returns the current state of mouse buttons (whether each button is pressed) as a bitmask. If multiple mouse buttons are pressed at the same time, the bits are added together. Equivalent to [method Input.get_mouse_button_mask].
*/
func MouseGetButtonState() MouseButtonMask {
	once.Do(singleton)
	return MouseButtonMask(class(self).MouseGetButtonState())
}

/*
Sets the user's clipboard content to the given string.
*/
func ClipboardSet(clipboard string) {
	once.Do(singleton)
	class(self).ClipboardSet(gd.NewString(clipboard))
}

/*
Returns the user's clipboard as a string if possible.
*/
func ClipboardGet() string {
	once.Do(singleton)
	return string(class(self).ClipboardGet().String())
}

/*
Returns the user's clipboard as an image if possible.
[b]Note:[/b] This method uses the copied pixel data, e.g. from a image editing software or a web browser, not an image file copied from file explorer.
*/
func ClipboardGetImage() [1]gdclass.Image {
	once.Do(singleton)
	return [1]gdclass.Image(class(self).ClipboardGetImage())
}

/*
Returns [code]true[/code] if there is a text content on the user's clipboard.
*/
func ClipboardHas() bool {
	once.Do(singleton)
	return bool(class(self).ClipboardHas())
}

/*
Returns [code]true[/code] if there is an image content on the user's clipboard.
*/
func ClipboardHasImage() bool {
	once.Do(singleton)
	return bool(class(self).ClipboardHasImage())
}

/*
Sets the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard content to the given string. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
func ClipboardSetPrimary(clipboard_primary string) {
	once.Do(singleton)
	class(self).ClipboardSetPrimary(gd.NewString(clipboard_primary))
}

/*
Returns the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard as a string if possible. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
func ClipboardGetPrimary() string {
	once.Do(singleton)
	return string(class(self).ClipboardGetPrimary().String())
}

/*
Returns an [Array] of [Rect2], each of which is the bounding rectangle for a display cutout or notch. These are non-functional areas on edge-to-edge screens used by cameras and sensors. Returns an empty array if the device does not have cutouts. See also [method get_display_safe_area].
[b]Note:[/b] Currently only implemented on Android. Other platforms will return an empty array even if they do have display cutouts or notches.
*/
func GetDisplayCutouts() []Rect2.PositionSize {
	once.Do(singleton)
	return []Rect2.PositionSize(gd.ArrayAs[[]Rect2.PositionSize](class(self).GetDisplayCutouts()))
}

/*
Returns the unobscured area of the display where interactive controls should be rendered. See also [method get_display_cutouts].
*/
func GetDisplaySafeArea() Rect2i.PositionSize {
	once.Do(singleton)
	return Rect2i.PositionSize(class(self).GetDisplaySafeArea())
}

/*
Returns the number of displays available.
*/
func GetScreenCount() int {
	once.Do(singleton)
	return int(int(class(self).GetScreenCount()))
}

/*
Returns index of the primary screen.
*/
func GetPrimaryScreen() int {
	once.Do(singleton)
	return int(int(class(self).GetPrimaryScreen()))
}

/*
Returns the index of the screen containing the window with the keyboard focus, or the primary screen if there's no focused window.
*/
func GetKeyboardFocusScreen() int {
	once.Do(singleton)
	return int(int(class(self).GetKeyboardFocusScreen()))
}

/*
Returns index of the screen which contains specified rectangle.
*/
func GetScreenFromRect(rect Rect2.PositionSize) int {
	once.Do(singleton)
	return int(int(class(self).GetScreenFromRect(gd.Rect2(rect))))
}

/*
Returns the screen's top-left corner position in pixels. On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
  - (0, 0)        +-------+
    |       |

+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
See also [method screen_get_size].
[b]Note:[/b] On Linux (Wayland) this method always returns [code](0, 0)[/code].
*/
func ScreenGetPosition() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).ScreenGetPosition(gd.Int(-1)))
}

/*
Returns the screen's size in pixels. See also [method screen_get_position] and [method screen_get_usable_rect].
*/
func ScreenGetSize() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).ScreenGetSize(gd.Int(-1)))
}

/*
Returns the portion of the screen that is not obstructed by a status bar in pixels. See also [method screen_get_size].
*/
func ScreenGetUsableRect() Rect2i.PositionSize {
	once.Do(singleton)
	return Rect2i.PositionSize(class(self).ScreenGetUsableRect(gd.Int(-1)))
}

/*
Returns the dots per inch density of the specified screen. If [param screen] is [constant SCREEN_OF_MAIN_WINDOW] (the default value), a screen with the main window will be used.
[b]Note:[/b] On macOS, returned value is inaccurate if fractional display scaling mode is used.
[b]Note:[/b] On Android devices, the actual screen densities are grouped into six generalized densities:
[codeblock lang=text]

	  ldpi - 120 dpi
	  mdpi - 160 dpi
	  hdpi - 240 dpi
	 xhdpi - 320 dpi
	xxhdpi - 480 dpi

xxxhdpi - 640 dpi
[/codeblock]
[b]Note:[/b] This method is implemented on Android, Linux (X11/Wayland), macOS and Windows. Returns [code]72[/code] on unsupported platforms.
*/
func ScreenGetDpi() int {
	once.Do(singleton)
	return int(int(class(self).ScreenGetDpi(gd.Int(-1))))
}

/*
Returns the scale factor of the specified screen by index.
[b]Note:[/b] On macOS, the returned value is [code]2.0[/code] for hiDPI (Retina) screens, and [code]1.0[/code] for all other cases.
[b]Note:[/b] On Linux (Wayland), the returned value is accurate only when [param screen] is [constant SCREEN_OF_MAIN_WINDOW]. Due to API limitations, passing a direct index will return a rounded-up integer, if the screen has a fractional scale (e.g. [code]1.25[/code] would get rounded up to [code]2.0[/code]).
[b]Note:[/b] This method is implemented only on macOS and Linux (Wayland).
*/
func ScreenGetScale() Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).ScreenGetScale(gd.Int(-1))))
}

/*
Returns [code]true[/code] if touch events are available (Android or iOS), the capability is detected on the Web platform or if [member ProjectSettings.input_devices/pointing/emulate_touch_from_mouse] is [code]true[/code].
*/
func IsTouchscreenAvailable() bool {
	once.Do(singleton)
	return bool(class(self).IsTouchscreenAvailable())
}

/*
Returns the greatest scale factor of all screens.
[b]Note:[/b] On macOS returned value is [code]2.0[/code] if there is at least one hiDPI (Retina) screen in the system, and [code]1.0[/code] in all other cases.
[b]Note:[/b] This method is implemented only on macOS.
*/
func ScreenGetMaxScale() Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).ScreenGetMaxScale()))
}

/*
Returns the current refresh rate of the specified screen. If [param screen] is [constant SCREEN_OF_MAIN_WINDOW] (the default value), a screen with the main window will be used.
[b]Note:[/b] Returns [code]-1.0[/code] if the DisplayServer fails to find the refresh rate for the specified screen. On Web, [method screen_get_refresh_rate] will always return [code]-1.0[/code] as there is no way to retrieve the refresh rate on that platform.
To fallback to a default refresh rate if the method fails, try:
[codeblock]
var refresh_rate = DisplayServer.screen_get_refresh_rate()
if refresh_rate < 0:

	refresh_rate = 60.0

[/codeblock]
*/
func ScreenGetRefreshRate() Float.X {
	once.Do(singleton)
	return Float.X(Float.X(class(self).ScreenGetRefreshRate(gd.Int(-1))))
}

/*
Returns color of the display pixel at the [param position].
[b]Note:[/b] This method is implemented on Linux (X11), macOS, and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
func ScreenGetPixel(position Vector2i.XY) Color.RGBA {
	once.Do(singleton)
	return Color.RGBA(class(self).ScreenGetPixel(gd.Vector2i(position)))
}

/*
Returns screenshot of the [param screen].
[b]Note:[/b] This method is implemented on Linux (X11), macOS, and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
func ScreenGetImage() [1]gdclass.Image {
	once.Do(singleton)
	return [1]gdclass.Image(class(self).ScreenGetImage(gd.Int(-1)))
}

/*
Sets the [param screen]'s [param orientation]. See also [method screen_get_orientation].
[b]Note:[/b] On iOS, this method has no effect if [member ProjectSettings.display/window/handheld/orientation] is not set to [constant SCREEN_SENSOR].
*/
func ScreenSetOrientation(orientation gdclass.DisplayServerScreenOrientation) {
	once.Do(singleton)
	class(self).ScreenSetOrientation(orientation, gd.Int(-1))
}

/*
Returns the [param screen]'s current orientation. See also [method screen_set_orientation].
[b]Note:[/b] This method is implemented on Android and iOS.
*/
func ScreenGetOrientation() gdclass.DisplayServerScreenOrientation {
	once.Do(singleton)
	return gdclass.DisplayServerScreenOrientation(class(self).ScreenGetOrientation(gd.Int(-1)))
}

/*
Sets whether the screen should never be turned off by the operating system's power-saving measures. See also [method screen_is_kept_on].
*/
func ScreenSetKeepOn(enable bool) {
	once.Do(singleton)
	class(self).ScreenSetKeepOn(enable)
}

/*
Returns [code]true[/code] if the screen should never be turned off by the operating system's power-saving measures. See also [method screen_set_keep_on].
*/
func ScreenIsKeptOn() bool {
	once.Do(singleton)
	return bool(class(self).ScreenIsKeptOn())
}

/*
Returns the list of Godot window IDs belonging to this process.
[b]Note:[/b] Native dialogs are not included in this list.
*/
func GetWindowList() []int32 {
	once.Do(singleton)
	return []int32(class(self).GetWindowList().AsSlice())
}

/*
Returns the ID of the window at the specified screen [param position] (in pixels). On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
  - (0, 0)        +-------+
    |       |

+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
*/
func GetWindowAtScreenPosition(position Vector2i.XY) int {
	once.Do(singleton)
	return int(int(class(self).GetWindowAtScreenPosition(gd.Vector2i(position))))
}

/*
Returns internal structure pointers for use in plugins.
[b]Note:[/b] This method is implemented on Android, Linux (X11/Wayland), macOS, and Windows.
*/
func WindowGetNativeHandle(handle_type gdclass.DisplayServerHandleType) int {
	once.Do(singleton)
	return int(int(class(self).WindowGetNativeHandle(handle_type, gd.Int(0))))
}

/*
Returns ID of the active popup window, or [constant INVALID_WINDOW_ID] if there is none.
*/
func WindowGetActivePopup() int {
	once.Do(singleton)
	return int(int(class(self).WindowGetActivePopup()))
}

/*
Sets the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system. Clicking this area will not auto-close this popup.
*/
func WindowSetPopupSafeRect(window int, rect Rect2i.PositionSize) {
	once.Do(singleton)
	class(self).WindowSetPopupSafeRect(gd.Int(window), gd.Rect2i(rect))
}

/*
Returns the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system.
*/
func WindowGetPopupSafeRect(window int) Rect2i.PositionSize {
	once.Do(singleton)
	return Rect2i.PositionSize(class(self).WindowGetPopupSafeRect(gd.Int(window)))
}

/*
Sets the title of the given window to [param title].
[b]Note:[/b] It's recommended to change this value using [member Window.title] instead.
[b]Note:[/b] Avoid changing the window title every frame, as this can cause performance issues on certain window managers. Try to change the window title only a few times per second at most.
*/
func WindowSetTitle(title string) {
	once.Do(singleton)
	class(self).WindowSetTitle(gd.NewString(title), gd.Int(0))
}

/*
Returns the estimated window title bar size (including text and window buttons) for the window specified by [param window_id] (in pixels). This method does not change the window title.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func WindowGetTitleSize(title string) Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetTitleSize(gd.NewString(title), gd.Int(0)))
}

/*
Sets a polygonal region of the window which accepts mouse events. Mouse events outside the region will be passed through.
Passing an empty array will disable passthrough support (all mouse events will be intercepted by the window, which is the default behavior).
[codeblocks]
[gdscript]
# Set region, using Path2D node.
DisplayServer.window_set_mouse_passthrough($Path2D.curve.get_baked_points())

# Set region, using Polygon2D node.
DisplayServer.window_set_mouse_passthrough($Polygon2D.polygon)

# Reset region to default.
DisplayServer.window_set_mouse_passthrough([])
[/gdscript]
[csharp]
// Set region, using Path2D node.
DisplayServer.WindowSetMousePassthrough(GetNode<Path2D>("Path2D").Curve.GetBakedPoints());

// Set region, using Polygon2D node.
DisplayServer.WindowSetMousePassthrough(GetNode<Polygon2D>("Polygon2D").Polygon);

// Reset region to default.
DisplayServer.WindowSetMousePassthrough(new Vector2[] {});
[/csharp]
[/codeblocks]
[b]Note:[/b] On Windows, the portion of a window that lies outside the region is not drawn, while on Linux (X11) and macOS it is.
[b]Note:[/b] This method is implemented on Linux (X11), macOS and Windows.
*/
func WindowSetMousePassthrough(region []Vector2.XY) {
	once.Do(singleton)
	class(self).WindowSetMousePassthrough(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&region))), gd.Int(0))
}

/*
Returns the screen the window specified by [param window_id] is currently positioned on. If the screen overlaps multiple displays, the screen where the window's center is located is returned. See also [method window_set_current_screen].
*/
func WindowGetCurrentScreen() int {
	once.Do(singleton)
	return int(int(class(self).WindowGetCurrentScreen(gd.Int(0))))
}

/*
Moves the window specified by [param window_id] to the specified [param screen]. See also [method window_get_current_screen].
*/
func WindowSetCurrentScreen(screen int) {
	once.Do(singleton)
	class(self).WindowSetCurrentScreen(gd.Int(screen), gd.Int(0))
}

/*
Returns the position of the client area of the given window on the screen.
*/
func WindowGetPosition() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetPosition(gd.Int(0)))
}

/*
Returns the position of the given window on the screen including the borders drawn by the operating system. See also [method window_get_position].
*/
func WindowGetPositionWithDecorations() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetPositionWithDecorations(gd.Int(0)))
}

/*
Sets the position of the given window to [param position]. On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
  - (0, 0)        +-------+
    |       |

+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
See also [method window_get_position] and [method window_set_size].
[b]Note:[/b] It's recommended to change this value using [member Window.position] instead.
[b]Note:[/b] On Linux (Wayland): this method is a no-op.
*/
func WindowSetPosition(position Vector2i.XY) {
	once.Do(singleton)
	class(self).WindowSetPosition(gd.Vector2i(position), gd.Int(0))
}

/*
Returns the size of the window specified by [param window_id] (in pixels), excluding the borders drawn by the operating system. This is also called the "client area". See also [method window_get_size_with_decorations], [method window_set_size] and [method window_get_position].
*/
func WindowGetSize() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetSize(gd.Int(0)))
}

/*
Sets the size of the given window to [param size] (in pixels). See also [method window_get_size] and [method window_get_position].
[b]Note:[/b] It's recommended to change this value using [member Window.size] instead.
*/
func WindowSetSize(size Vector2i.XY) {
	once.Do(singleton)
	class(self).WindowSetSize(gd.Vector2i(size), gd.Int(0))
}

/*
Sets the [param callback] that will be called when the window specified by [param window_id] is moved or resized.
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
func WindowSetRectChangedCallback(callback func(rect Rect2i.PositionSize)) {
	once.Do(singleton)
	class(self).WindowSetRectChangedCallback(gd.NewCallable(callback), gd.Int(0))
}

/*
Sets the [param callback] that will be called when an event occurs in the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
func WindowSetWindowEventCallback(callback func(event gdclass.DisplayServerWindowEvent)) {
	once.Do(singleton)
	class(self).WindowSetWindowEventCallback(gd.NewCallable(callback), gd.Int(0))
}

/*
Sets the [param callback] that should be called when any [InputEvent] is sent to the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
func WindowSetInputEventCallback(callback func(event [1]gdclass.InputEvent)) {
	once.Do(singleton)
	class(self).WindowSetInputEventCallback(gd.NewCallable(callback), gd.Int(0))
}

/*
Sets the [param callback] that should be called when text is entered using the virtual keyboard to the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
func WindowSetInputTextCallback(callback func(text string)) {
	once.Do(singleton)
	class(self).WindowSetInputTextCallback(gd.NewCallable(callback), gd.Int(0))
}

/*
Sets the [param callback] that should be called when files are dropped from the operating system's file manager to the window specified by [param window_id]. [param callback] should take one [PackedStringArray] argument, which is the list of dropped files.
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
[b]Note:[/b] This method is implemented on Windows, macOS, Linux (X11/Wayland), and Web.
*/
func WindowSetDropFilesCallback(callback func(tag any)) {
	once.Do(singleton)
	class(self).WindowSetDropFilesCallback(gd.NewCallable(callback), gd.Int(0))
}

/*
Returns the [method Object.get_instance_id] of the [Window] the [param window_id] is attached to.
*/
func WindowGetAttachedInstanceId() int {
	once.Do(singleton)
	return int(int(class(self).WindowGetAttachedInstanceId(gd.Int(0))))
}

/*
Returns the window's maximum size (in pixels). See also [method window_set_max_size].
*/
func WindowGetMaxSize() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetMaxSize(gd.Int(0)))
}

/*
Sets the maximum size of the window specified by [param window_id] in pixels. Normally, the user will not be able to drag the window to make it larger than the specified size. See also [method window_get_max_size].
[b]Note:[/b] It's recommended to change this value using [member Window.max_size] instead.
[b]Note:[/b] Using third-party tools, it is possible for users to disable window geometry restrictions and therefore bypass this limit.
*/
func WindowSetMaxSize(max_size Vector2i.XY) {
	once.Do(singleton)
	class(self).WindowSetMaxSize(gd.Vector2i(max_size), gd.Int(0))
}

/*
Returns the window's minimum size (in pixels). See also [method window_set_min_size].
*/
func WindowGetMinSize() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetMinSize(gd.Int(0)))
}

/*
Sets the minimum size for the given window to [param min_size] in pixels. Normally, the user will not be able to drag the window to make it smaller than the specified size. See also [method window_get_min_size].
[b]Note:[/b] It's recommended to change this value using [member Window.min_size] instead.
[b]Note:[/b] By default, the main window has a minimum size of [code]Vector2i(64, 64)[/code]. This prevents issues that can arise when the window is resized to a near-zero size.
[b]Note:[/b] Using third-party tools, it is possible for users to disable window geometry restrictions and therefore bypass this limit.
*/
func WindowSetMinSize(min_size Vector2i.XY) {
	once.Do(singleton)
	class(self).WindowSetMinSize(gd.Vector2i(min_size), gd.Int(0))
}

/*
Returns the size of the window specified by [param window_id] (in pixels), including the borders drawn by the operating system. See also [method window_get_size].
*/
func WindowGetSizeWithDecorations() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetSizeWithDecorations(gd.Int(0)))
}

/*
Returns the mode of the given window.
*/
func WindowGetMode() gdclass.DisplayServerWindowMode {
	once.Do(singleton)
	return gdclass.DisplayServerWindowMode(class(self).WindowGetMode(gd.Int(0)))
}

/*
Sets window mode for the given window to [param mode]. See [enum WindowMode] for possible values and how each mode behaves.
[b]Note:[/b] Setting the window to full screen forcibly sets the borderless flag to [code]true[/code], so make sure to set it back to [code]false[/code] when not wanted.
*/
func WindowSetMode(mode gdclass.DisplayServerWindowMode) {
	once.Do(singleton)
	class(self).WindowSetMode(mode, gd.Int(0))
}

/*
Enables or disables the given window's given [param flag]. See [enum WindowFlags] for possible values and their behavior.
*/
func WindowSetFlag(flag gdclass.DisplayServerWindowFlags, enabled bool) {
	once.Do(singleton)
	class(self).WindowSetFlag(flag, enabled, gd.Int(0))
}

/*
Returns the current value of the given window's [param flag].
*/
func WindowGetFlag(flag gdclass.DisplayServerWindowFlags) bool {
	once.Do(singleton)
	return bool(class(self).WindowGetFlag(flag, gd.Int(0)))
}

/*
When [constant WINDOW_FLAG_EXTEND_TO_TITLE] flag is set, set offset to the center of the first titlebar button.
[b]Note:[/b] This flag is implemented only on macOS.
*/
func WindowSetWindowButtonsOffset(offset Vector2i.XY) {
	once.Do(singleton)
	class(self).WindowSetWindowButtonsOffset(gd.Vector2i(offset), gd.Int(0))
}

/*
Returns left margins ([code]x[/code]), right margins ([code]y[/code]) and height ([code]z[/code]) of the title that are safe to use (contains no buttons or other elements) when [constant WINDOW_FLAG_EXTEND_TO_TITLE] flag is set.
*/
func WindowGetSafeTitleMargins() Vector3i.XYZ {
	once.Do(singleton)
	return Vector3i.XYZ(class(self).WindowGetSafeTitleMargins(gd.Int(0)))
}

/*
Makes the window specified by [param window_id] request attention, which is materialized by the window title and taskbar entry blinking until the window is focused. This usually has no visible effect if the window is currently focused. The exact behavior varies depending on the operating system.
*/
func WindowRequestAttention() {
	once.Do(singleton)
	class(self).WindowRequestAttention(gd.Int(0))
}

/*
Moves the window specified by [param window_id] to the foreground, so that it is visible over other windows.
*/
func WindowMoveToForeground() {
	once.Do(singleton)
	class(self).WindowMoveToForeground(gd.Int(0))
}

/*
Returns [code]true[/code] if the window specified by [param window_id] is focused.
*/
func WindowIsFocused() bool {
	once.Do(singleton)
	return bool(class(self).WindowIsFocused(gd.Int(0)))
}

/*
Returns [code]true[/code] if anything can be drawn in the window specified by [param window_id], [code]false[/code] otherwise. Using the [code]--disable-render-loop[/code] command line argument or a headless build will return [code]false[/code].
*/
func WindowCanDraw() bool {
	once.Do(singleton)
	return bool(class(self).WindowCanDraw(gd.Int(0)))
}

/*
Sets window transient parent. Transient window is will be destroyed with its transient parent and will return focus to their parent when closed. The transient window is displayed on top of a non-exclusive full-screen parent window. Transient windows can't enter full-screen mode.
[b]Note:[/b] It's recommended to change this value using [member Window.transient] instead.
[b]Note:[/b] The behavior might be different depending on the platform.
*/
func WindowSetTransient(window_id int, parent_window_id int) {
	once.Do(singleton)
	class(self).WindowSetTransient(gd.Int(window_id), gd.Int(parent_window_id))
}

/*
If set to [code]true[/code], this window will always stay on top of its parent window, parent window will ignore input while this window is opened.
[b]Note:[/b] On macOS, exclusive windows are confined to the same space (virtual desktop or screen) as the parent window.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func WindowSetExclusive(window_id int, exclusive bool) {
	once.Do(singleton)
	class(self).WindowSetExclusive(gd.Int(window_id), exclusive)
}

/*
Sets whether [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] should be enabled for the window specified by [param window_id]. See also [method window_set_ime_position].
*/
func WindowSetImeActive(active bool) {
	once.Do(singleton)
	class(self).WindowSetImeActive(active, gd.Int(0))
}

/*
Sets the position of the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] popup for the specified [param window_id]. Only effective if [method window_set_ime_active] was set to [code]true[/code] for the specified [param window_id].
*/
func WindowSetImePosition(position Vector2i.XY) {
	once.Do(singleton)
	class(self).WindowSetImePosition(gd.Vector2i(position), gd.Int(0))
}

/*
Sets the V-Sync mode of the given window. See also [member ProjectSettings.display/window/vsync/vsync_mode].
See [enum DisplayServer.VSyncMode] for possible values and how they affect the behavior of your application.
Depending on the platform and used renderer, the engine will fall back to [constant VSYNC_ENABLED] if the desired mode is not supported.
[b]Note:[/b] V-Sync modes other than [constant VSYNC_ENABLED] are only supported in the Forward+ and Mobile rendering methods, not Compatibility.
*/
func WindowSetVsyncMode(vsync_mode gdclass.DisplayServerVSyncMode) {
	once.Do(singleton)
	class(self).WindowSetVsyncMode(vsync_mode, gd.Int(0))
}

/*
Returns the V-Sync mode of the given window.
*/
func WindowGetVsyncMode() gdclass.DisplayServerVSyncMode {
	once.Do(singleton)
	return gdclass.DisplayServerVSyncMode(class(self).WindowGetVsyncMode(gd.Int(0)))
}

/*
Returns [code]true[/code] if the given window can be maximized (the maximize button is enabled).
*/
func WindowIsMaximizeAllowed() bool {
	once.Do(singleton)
	return bool(class(self).WindowIsMaximizeAllowed(gd.Int(0)))
}

/*
Returns [code]true[/code], if double-click on a window title should maximize it.
[b]Note:[/b] This method is implemented only on macOS.
*/
func WindowMaximizeOnTitleDblClick() bool {
	once.Do(singleton)
	return bool(class(self).WindowMaximizeOnTitleDblClick())
}

/*
Returns [code]true[/code], if double-click on a window title should minimize it.
[b]Note:[/b] This method is implemented only on macOS.
*/
func WindowMinimizeOnTitleDblClick() bool {
	once.Do(singleton)
	return bool(class(self).WindowMinimizeOnTitleDblClick())
}

/*
Returns the text selection in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] composition string, with the [Vector2i]'s [code]x[/code] component being the caret position and [code]y[/code] being the length of the selection.
[b]Note:[/b] This method is implemented only on macOS.
*/
func ImeGetSelection() Vector2i.XY {
	once.Do(singleton)
	return Vector2i.XY(class(self).ImeGetSelection())
}

/*
Returns the composition string contained within the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] window.
[b]Note:[/b] This method is implemented only on macOS.
*/
func ImeGetText() string {
	once.Do(singleton)
	return string(class(self).ImeGetText().String())
}

/*
Shows the virtual keyboard if the platform has one.
[param existing_text] parameter is useful for implementing your own [LineEdit] or [TextEdit], as it tells the virtual keyboard what text has already been typed (the virtual keyboard uses it for auto-correct and predictions).
[param position] parameter is the screen space [Rect2] of the edited text.
[param type] parameter allows configuring which type of virtual keyboard to show.
[param max_length] limits the number of characters that can be entered if different from [code]-1[/code].
[param cursor_start] can optionally define the current text cursor position if [param cursor_end] is not set.
[param cursor_start] and [param cursor_end] can optionally define the current text selection.
[b]Note:[/b] This method is implemented on Android, iOS and Web.
*/
func VirtualKeyboardShow(existing_text string) {
	once.Do(singleton)
	class(self).VirtualKeyboardShow(gd.NewString(existing_text), gd.Rect2(gd.NewRect2(0, 0, 0, 0)), 0, gd.Int(-1), gd.Int(-1), gd.Int(-1))
}

/*
Hides the virtual keyboard if it is shown, does nothing otherwise.
*/
func VirtualKeyboardHide() {
	once.Do(singleton)
	class(self).VirtualKeyboardHide()
}

/*
Returns the on-screen keyboard's height in pixels. Returns 0 if there is no keyboard or if it is currently hidden.
*/
func VirtualKeyboardGetHeight() int {
	once.Do(singleton)
	return int(int(class(self).VirtualKeyboardGetHeight()))
}

/*
Sets the default mouse cursor shape. The cursor's appearance will vary depending on the user's operating system and mouse cursor theme. See also [method cursor_get_shape] and [method cursor_set_custom_image].
*/
func CursorSetShape(shape gdclass.DisplayServerCursorShape) {
	once.Do(singleton)
	class(self).CursorSetShape(shape)
}

/*
Returns the default mouse cursor shape set by [method cursor_set_shape].
*/
func CursorGetShape() gdclass.DisplayServerCursorShape {
	once.Do(singleton)
	return gdclass.DisplayServerCursorShape(class(self).CursorGetShape())
}

/*
Sets a custom mouse cursor image for the given [param shape]. This means the user's operating system and mouse cursor theme will no longer influence the mouse cursor's appearance.
[param cursor] can be either a [Texture2D] or an [Image], and it should not be larger than 256×256 to display correctly. Optionally, [param hotspot] can be set to offset the image's position relative to the click point. By default, [param hotspot] is set to the top-left corner of the image. See also [method cursor_set_shape].
*/
func CursorSetCustomImage(cursor [1]gdclass.Resource) {
	once.Do(singleton)
	class(self).CursorSetCustomImage(cursor, 0, gd.Vector2(gd.Vector2{0, 0}))
}

/*
Returns [code]true[/code] if positions of [b]OK[/b] and [b]Cancel[/b] buttons are swapped in dialogs. This is enabled by default on Windows to follow interface conventions, and be toggled by changing [member ProjectSettings.gui/common/swap_cancel_ok].
[b]Note:[/b] This doesn't affect native dialogs such as the ones spawned by [method DisplayServer.dialog_show].
*/
func GetSwapCancelOk() bool {
	once.Do(singleton)
	return bool(class(self).GetSwapCancelOk())
}

/*
Allows the [param process_id] PID to steal focus from this window. In other words, this disables the operating system's focus stealing protection for the specified PID.
[b]Note:[/b] This method is implemented only on Windows.
*/
func EnableForStealingFocus(process_id int) {
	once.Do(singleton)
	class(self).EnableForStealingFocus(gd.Int(process_id))
}

/*
Shows a text dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [int] parameter which corresponds to the index of the pressed button.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG] feature. Supported platforms include macOS and Windows.
*/
func DialogShow(title string, description string, buttons []string, callback func(button int)) error {
	once.Do(singleton)
	return error(gd.ToError(class(self).DialogShow(gd.NewString(title), gd.NewString(description), gd.NewPackedStringSlice(buttons), gd.NewCallable(callback))))
}

/*
Shows a text input dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [String] parameter which contains the text field's contents.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_INPUT] feature. Supported platforms include macOS and Windows.
*/
func DialogInputText(title string, description string, existing_text string, callback func(text string)) error {
	once.Do(singleton)
	return error(gd.ToError(class(self).DialogInputText(gd.NewString(title), gd.NewString(description), gd.NewString(existing_text), gd.NewCallable(callback))))
}

/*
Displays OS native dialog for selecting files or directories in the file system.
Each filter string in the [param filters] array should be formatted like this: [code]*.txt,*.doc;Text Files[/code]. The description text of the filter is optional and can be omitted. See also [member FileDialog.filters].
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int[/code].
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE] feature. Supported platforms include Linux (X11/Wayland), Windows, and macOS.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] On Linux, [param show_hidden] is ignored.
[b]Note:[/b] On macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
func FileDialogShow(title string, current_directory string, filename string, show_hidden bool, mode gdclass.DisplayServerFileDialogMode, filters []string, callback func(status bool, selected_paths []string, selected_filter_index int)) error {
	once.Do(singleton)
	return error(gd.ToError(class(self).FileDialogShow(gd.NewString(title), gd.NewString(current_directory), gd.NewString(filename), show_hidden, mode, gd.NewPackedStringSlice(filters), gd.NewCallable(callback))))
}

/*
Displays OS native dialog for selecting files or directories in the file system with additional user selectable options.
Each filter string in the [param filters] array should be formatted like this: [code]*.txt,*.doc;Text Files[/code]. The description text of the filter is optional and can be omitted. See also [member FileDialog.filters].
[param options] is array of [Dictionary]s with the following keys:
- [code]"name"[/code] - option's name [String].
- [code]"values"[/code] - [PackedStringArray] of values. If empty, boolean option (check box) is used.
- [code]"default"[/code] - default selected option index ([int]) or default boolean value ([bool]).
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int, selected_option: Dictionary[/code].
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE] feature. Supported platforms include Linux (X11/Wayland), Windows, and macOS.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] On Linux (X11), [param show_hidden] is ignored.
[b]Note:[/b] On macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
func FileDialogWithOptionsShow(title string, current_directory string, root string, filename string, show_hidden bool, mode gdclass.DisplayServerFileDialogMode, filters []string, options []map[any]any, callback func(status bool, selected_paths []string, selected_filter_index int, selected_option map[any]any)) error {
	once.Do(singleton)
	return error(gd.ToError(class(self).FileDialogWithOptionsShow(gd.NewString(title), gd.NewString(current_directory), gd.NewString(root), gd.NewString(filename), show_hidden, mode, gd.NewPackedStringSlice(filters), gd.NewVariant(options).Interface().(gd.Array), gd.NewCallable(callback))))
}

/*
Returns the number of keyboard layouts.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetLayoutCount() int {
	once.Do(singleton)
	return int(int(class(self).KeyboardGetLayoutCount()))
}

/*
Returns active keyboard layout index.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS, and Windows.
*/
func KeyboardGetCurrentLayout() int {
	once.Do(singleton)
	return int(int(class(self).KeyboardGetCurrentLayout()))
}

/*
Sets the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardSetCurrentLayout(index int) {
	once.Do(singleton)
	class(self).KeyboardSetCurrentLayout(gd.Int(index))
}

/*
Returns the ISO-639/BCP-47 language code of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetLayoutLanguage(index int) string {
	once.Do(singleton)
	return string(class(self).KeyboardGetLayoutLanguage(gd.Int(index)).String())
}

/*
Returns the localized name of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetLayoutName(index int) string {
	once.Do(singleton)
	return string(class(self).KeyboardGetLayoutName(gd.Int(index)).String())
}

/*
Converts a physical (US QWERTY) [param keycode] to one in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetKeycodeFromPhysical(keycode Key) Key {
	once.Do(singleton)
	return Key(class(self).KeyboardGetKeycodeFromPhysical(keycode))
}

/*
Converts a physical (US QWERTY) [param keycode] to localized label printed on the key in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetLabelFromPhysical(keycode Key) Key {
	once.Do(singleton)
	return Key(class(self).KeyboardGetLabelFromPhysical(keycode))
}

/*
Perform window manager processing, including input flushing. See also [method force_process_and_drop_events], [method Input.flush_buffered_events] and [member Input.use_accumulated_input].
*/
func ProcessEvents() {
	once.Do(singleton)
	class(self).ProcessEvents()
}

/*
Forces window manager processing while ignoring all [InputEvent]s. See also [method process_events].
[b]Note:[/b] This method is implemented on Windows and macOS.
*/
func ForceProcessAndDropEvents() {
	once.Do(singleton)
	class(self).ForceProcessAndDropEvents()
}

/*
Sets the window icon (usually displayed in the top-left corner) in the operating system's [i]native[/i] format. The file at [param filename] must be in [code].ico[/code] format on Windows or [code].icns[/code] on macOS. By using specially crafted [code].ico[/code] or [code].icns[/code] icons, [method set_native_icon] allows specifying different icons depending on the size the icon is displayed at. This size is determined by the operating system and user preferences (including the display scale factor). To use icons in other formats, use [method set_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_NATIVE_ICON].
*/
func SetNativeIcon(filename string) {
	once.Do(singleton)
	class(self).SetNativeIcon(gd.NewString(filename))
}

/*
Sets the window icon (usually displayed in the top-left corner) with an [Image]. To use icons in the operating system's native format, use [method set_native_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_ICON].
*/
func SetIcon(image [1]gdclass.Image) {
	once.Do(singleton)
	class(self).SetIcon(image)
}

/*
Creates a new application status indicator with the specified icon, tooltip, and activation callback.
[param callback] should take two arguments: the pressed mouse button (one of the [enum MouseButton] constants) and the click position in screen coordinates (a [Vector2i]).
*/
func CreateStatusIndicator(icon [1]gdclass.Texture2D, tooltip string, callback func(button MouseButton, click_position Vector2i.XY)) int {
	once.Do(singleton)
	return int(int(class(self).CreateStatusIndicator(icon, gd.NewString(tooltip), gd.NewCallable(callback))))
}

/*
Sets the application status indicator icon.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func StatusIndicatorSetIcon(id int, icon [1]gdclass.Texture2D) {
	once.Do(singleton)
	class(self).StatusIndicatorSetIcon(gd.Int(id), icon)
}

/*
Sets the application status indicator tooltip.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func StatusIndicatorSetTooltip(id int, tooltip string) {
	once.Do(singleton)
	class(self).StatusIndicatorSetTooltip(gd.Int(id), gd.NewString(tooltip))
}

/*
Sets the application status indicator native popup menu.
[b]Note:[/b] On macOS, the menu is activated by any mouse button. Its activation callback is [i]not[/i] triggered.
[b]Note:[/b] On Windows, the menu is activated by the right mouse button, selecting the status icon and pressing [kbd]Shift + F10[/kbd], or the applications key. The menu's activation callback for the other mouse buttons is still triggered.
[b]Note:[/b] Native popup is only supported if [NativeMenu] supports the [constant NativeMenu.FEATURE_POPUP_MENU] feature.
*/
func StatusIndicatorSetMenu(id int, menu_rid Resource.ID) {
	once.Do(singleton)
	class(self).StatusIndicatorSetMenu(gd.Int(id), menu_rid)
}

/*
Sets the application status indicator activation callback. [param callback] should take two arguments: [int] mouse button index (one of [enum MouseButton] values) and [Vector2i] click position in screen coordinates.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func StatusIndicatorSetCallback(id int, callback func(button MouseButton, click_position Vector2i.XY)) {
	once.Do(singleton)
	class(self).StatusIndicatorSetCallback(gd.Int(id), gd.NewCallable(callback))
}

/*
Returns the rectangle for the given status indicator [param id] in screen coordinates. If the status indicator is not visible, returns an empty [Rect2].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func StatusIndicatorGetRect(id int) Rect2.PositionSize {
	once.Do(singleton)
	return Rect2.PositionSize(class(self).StatusIndicatorGetRect(gd.Int(id)))
}

/*
Removes the application status indicator.
*/
func DeleteStatusIndicator(id int) {
	once.Do(singleton)
	class(self).DeleteStatusIndicator(gd.Int(id))
}

/*
Returns the total number of available tablet drivers.
[b]Note:[/b] This method is implemented only on Windows.
*/
func TabletGetDriverCount() int {
	once.Do(singleton)
	return int(int(class(self).TabletGetDriverCount()))
}

/*
Returns the tablet driver name for the given index.
[b]Note:[/b] This method is implemented only on Windows.
*/
func TabletGetDriverName(idx int) string {
	once.Do(singleton)
	return string(class(self).TabletGetDriverName(gd.Int(idx)).String())
}

/*
Returns current active tablet driver name.
[b]Note:[/b] This method is implemented only on Windows.
*/
func TabletGetCurrentDriver() string {
	once.Do(singleton)
	return string(class(self).TabletGetCurrentDriver().String())
}

/*
Set active tablet driver name.
Supported drivers:
- [code]winink[/code]: Windows Ink API, default (Windows 8.1+ required).
- [code]wintab[/code]: Wacom Wintab API (compatible device driver required).
- [code]dummy[/code]: Dummy driver, tablet input is disabled.
[b]Note:[/b] This method is implemented only on Windows.
*/
func TabletSetCurrentDriver(name string) {
	once.Do(singleton)
	class(self).TabletSetCurrentDriver(gd.NewString(name))
}

/*
Returns [code]true[/code] if the window background can be made transparent. This method returns [code]false[/code] if [member ProjectSettings.display/window/per_pixel_transparency/allowed] is set to [code]false[/code], or if transparency is not supported by the renderer or OS compositor.
*/
func IsWindowTransparencyAvailable() bool {
	once.Do(singleton)
	return bool(class(self).IsWindowTransparencyAvailable())
}

/*
Registers an [Object] which represents an additional output that will be rendered too, beyond normal windows. The [Object] is only used as an identifier, which can be later passed to [method unregister_additional_output].
This can be used to prevent Godot from skipping rendering when no normal windows are visible.
*/
func RegisterAdditionalOutput(obj Object.Instance) {
	once.Do(singleton)
	class(self).RegisterAdditionalOutput(obj)
}

/*
Unregisters an [Object] representing an additional output, that was registered via [method register_additional_output].
*/
func UnregisterAdditionalOutput(obj Object.Instance) {
	once.Do(singleton)
	class(self).UnregisterAdditionalOutput(obj)
}

/*
Returns [code]true[/code] if any additional outputs have been registered via [method register_additional_output].
*/
func HasAdditionalOutputs() bool {
	once.Do(singleton)
	return bool(class(self).HasAdditionalOutputs())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.DisplayServer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns [code]true[/code] if the specified [param feature] is supported by the current [DisplayServer], [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) HasFeature(feature gdclass.DisplayServerFeature) bool {
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the [DisplayServer] currently in use. Most operating systems only have a single [DisplayServer], but Linux has access to more than one [DisplayServer] (currently X11 and Wayland).
The names of built-in display servers are [code]Windows[/code], [code]macOS[/code], [code]X11[/code] (Linux), [code]Wayland[/code] (Linux), [code]Android[/code], [code]iOS[/code], [code]web[/code] (HTML5), and [code]headless[/code] (when started with the [code]--headless[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url]).
*/
//go:nosplit
func (self class) GetName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets native help system search callbacks.
[param search_callback] has the following arguments: [code]String search_string, int result_limit[/code] and return a [Dictionary] with "key, display name" pairs for the search results. Called when the user enters search terms in the [code]Help[/code] menu.
[param action_callback] has the following arguments: [code]String key[/code]. Called when the user selects a search result in the [code]Help[/code] menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) HelpSetSearchCallbacks(search_callback gd.Callable, action_callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(search_callback))
	callframe.Arg(frame, pointers.Get(action_callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_help_set_search_callbacks, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers callables to emit when the menu is respectively about to show or closed. Callback methods should have zero arguments.
*/
//go:nosplit
func (self class) GlobalMenuSetPopupCallbacks(menu_root gd.String, open_callback gd.Callable, close_callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(open_callback))
	callframe.Arg(frame, pointers.Get(close_callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_popup_callbacks, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds an item that will act as a submenu of the global menu [param menu_root]. The [param submenu] argument is the ID of the global menu root that will be shown when the item is clicked.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddSubmenuItem(menu_root gd.String, label gd.String, submenu gd.String, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(submenu))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_submenu_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddItem(menu_root gd.String, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new checkable item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddCheckItem(menu_root gd.String, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_check_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddIconItem(menu_root gd.String, icon [1]gdclass.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new checkable item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddIconCheckItem(menu_root gd.String, icon [1]gdclass.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_icon_check_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new radio-checkable item with text [param label] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] Radio-checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method global_menu_set_item_checked] for more info on how to control it.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddRadioCheckItem(menu_root gd.String, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new radio-checkable item with text [param label] and icon [param icon] to the global menu with ID [param menu_root].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] Radio-checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method global_menu_set_item_checked] for more info on how to control it.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddIconRadioCheckItem(menu_root gd.String, icon [1]gdclass.Texture2D, label gd.String, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_icon_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new item with text [param label] to the global menu with ID [param menu_root].
Contrarily to normal binary items, multistate items can have more than two states, as defined by [param max_states]. Each press or activate of the item will increase the state by one. The default value is defined by [param default_state].
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
An [param accelerator] can optionally be defined, which is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The [param accelerator] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] By default, there's no indication of the current item state, it should be changed manually.
[b]Note:[/b] The [param callback] and [param key_callback] Callables need to accept exactly one Variant parameter, the parameter passed to the Callables will be the value passed to [param tag].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddMultistateItem(menu_root gd.String, label gd.String, max_states gd.Int, default_state gd.Int, callback gd.Callable, key_callback gd.Callable, tag gd.Variant, accelerator Key, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, pointers.Get(key_callback))
	callframe.Arg(frame, pointers.Get(tag))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_multistate_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a separator between items to the global menu with ID [param menu_root]. Separators also occupy an index.
Returns index of the inserted item, it's not guaranteed to be the same as [param index] value.
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuAddSeparator(menu_root gd.String, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_add_separator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the item with the specified [param text]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemIndexFromText(menu_root gd.String, text gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(text))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_index_from_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the item with the specified [param tag]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemIndexFromTag(menu_root gd.String, tag gd.Variant) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, pointers.Get(tag))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_index_from_tag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at index [param idx] is checked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemChecked(menu_root gd.String, idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_is_item_checked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at index [param idx] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemCheckable(menu_root gd.String, idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_is_item_checkable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at index [param idx] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemRadioCheckable(menu_root gd.String, idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_is_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the callback of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemCallback(menu_root gd.String, idx gd.Int) gd.Callable {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemKeyCallback(menu_root gd.String, idx gd.Int) gd.Callable {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method global_menu_set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemTag(menu_root gd.String, idx gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_tag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemText(menu_root gd.String, idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the submenu ID of the item at index [param idx]. See [method global_menu_add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemSubmenu(menu_root gd.String, idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_submenu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the accelerator of the item at index [param idx]. Accelerators are special combinations of keys that activate the item, no matter which control is focused.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemAccelerator(menu_root gd.String, idx gd.Int) Key {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Key](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at index [param idx] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method global_menu_set_item_disabled] for more info on how to disable an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemDisabled(menu_root gd.String, idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at index [param idx] is hidden.
See [method global_menu_set_item_hidden] for more info on how to hide an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuIsItemHidden(menu_root gd.String, idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_is_item_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tooltip associated with the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemTooltip(menu_root gd.String, idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemState(menu_root gd.String, idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of states of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemMaxStates(menu_root gd.String, idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_max_states, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the icon of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemIcon(menu_root gd.String, idx gd.Int) [1]gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemIndentationLevel(menu_root gd.String, idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_indentation_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the checkstate status of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemChecked(menu_root gd.String, idx gd.Int, checked bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checked)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_checked, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets whether the item at index [param idx] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemCheckable(menu_root gd.String, idx gd.Int, checkable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checkable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_checkable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the type of the item at the specified index [param idx] to radio button. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemRadioCheckable(menu_root gd.String, idx gd.Int, checkable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, checkable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when an item is pressed.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemCallback(menu_root gd.String, idx gd.Int, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the callback of the item at index [param idx]. The callback is emitted when an item is hovered.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemHoverCallbacks(menu_root gd.String, idx gd.Int, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_hover_callbacks, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when its accelerator is activated.
[b]Note:[/b] The [param key_callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemKeyCallback(menu_root gd.String, idx gd.Int, key_callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(key_callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method global_menu_get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemTag(menu_root gd.String, idx gd.Int, tag gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(tag))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_tag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemText(menu_root gd.String, idx gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the submenu of the item at index [param idx]. The submenu is the ID of a global menu root that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemSubmenu(menu_root gd.String, idx gd.Int, submenu gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(submenu))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_submenu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemAccelerator(menu_root gd.String, idx gd.Int, keycode Key) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Enables/disables the item at index [param idx]. When it is disabled, it can't be selected and its action can't be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemDisabled(menu_root gd.String, idx gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Hides/shows the item at index [param idx]. When it is hidden, an item does not appear in a menu and its action cannot be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemHidden(menu_root gd.String, idx gd.Int, hidden bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [String] tooltip of the item at the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemTooltip(menu_root gd.String, idx gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemState(menu_root gd.String, idx gd.Int, state gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, state)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets number of state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemMaxStates(menu_root gd.String, idx gd.Int, max_states gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, max_states)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_max_states, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Replaces the [Texture2D] icon of the specified [param idx].
[b]Note:[/b] This method is implemented only on macOS.
[b]Note:[/b] This method is not supported by macOS "_dock" menu items.
*/
//go:nosplit
func (self class) GlobalMenuSetItemIcon(menu_root gd.String, idx gd.Int, icon [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemIndentationLevel(menu_root gd.String, idx gd.Int, level gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, level)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_indentation_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns number of items in the global menu with ID [param menu_root].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemCount(menu_root gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the item at index [param idx] from the global menu [param menu_root].
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuRemoveItem(menu_root gd.String, idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_remove_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all items from the global menu with ID [param menu_root].
[b]Note:[/b] This method is implemented only on macOS.
[b]Supported system menu IDs:[/b]
[codeblock lang=text]
"_main" - Main menu (macOS).
"_dock" - Dock popup menu (macOS).
"_apple" - Apple menu (macOS, custom items added before "Services").
"_window" - Window menu (macOS, custom items added after "Bring All to Front").
"_help" - Help menu (macOS).
[/codeblock]
*/
//go:nosplit
func (self class) GlobalMenuClear(menu_root gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu_root))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns Dictionary of supported system menu IDs and names.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetSystemMenuRoots() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_system_menu_roots, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the synthesizer is generating speech, or have utterance waiting in the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsIsSpeaking() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_is_speaking, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the synthesizer is in a paused state.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsIsPaused() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_is_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an [Array] of voice information dictionaries.
Each [Dictionary] contains two [String] entries:
- [code]name[/code] is voice name.
- [code]id[/code] is voice identifier.
- [code]language[/code] is language code in [code]lang_Variant[/code] format. The [code]lang[/code] part is a 2 or 3-letter code based on the ISO-639 standard, in lowercase. The [code skip-lint]Variant[/code] part is an engine-dependent string describing country, region or/and dialect.
Note that Godot depends on system libraries for text-to-speech functionality. These libraries are installed by default on Windows and macOS, but not on all Linux distributions. If they are not present, this method will return an empty list. This applies to both Godot users on Linux, as well as end-users on Linux running Godot games that use text-to-speech.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsGetVoices() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_get_voices, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an [PackedStringArray] of voice identifiers for the [param language].
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsGetVoicesForLanguage(language gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_get_voices_for_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds an utterance to the queue. If [param interrupt] is [code]true[/code], the queue is cleared first.
- [param voice] identifier is one of the [code]"id"[/code] values returned by [method tts_get_voices] or one of the values returned by [method tts_get_voices_for_language].
- [param volume] ranges from [code]0[/code] (lowest) to [code]100[/code] (highest).
- [param pitch] ranges from [code]0.0[/code] (lowest) to [code]2.0[/code] (highest), [code]1.0[/code] is default pitch for the current voice.
- [param rate] ranges from [code]0.1[/code] (lowest) to [code]10.0[/code] (highest), [code]1.0[/code] is a normal speaking rate. Other values act as a percentage relative.
- [param utterance_id] is passed as a parameter to the callback functions.
[b]Note:[/b] On Windows and Linux (X11/Wayland), utterance [param text] can use SSML markup. SSML support is engine and voice dependent. If the engine does not support SSML, you should strip out all XML markup before calling [method tts_speak].
[b]Note:[/b] The granularity of pitch, rate, and volume is engine and voice dependent. Values may be truncated.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsSpeak(text gd.String, voice gd.String, volume gd.Int, pitch gd.Float, rate gd.Float, utterance_id gd.Int, interrupt bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, pointers.Get(voice))
	callframe.Arg(frame, volume)
	callframe.Arg(frame, pitch)
	callframe.Arg(frame, rate)
	callframe.Arg(frame, utterance_id)
	callframe.Arg(frame, interrupt)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_speak, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Puts the synthesizer into a paused state.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsPause() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_pause, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Resumes the synthesizer if it was paused.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsResume() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_resume, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops synthesis in progress and removes all utterances from the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Linux), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsStop() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a callback, which is called when the utterance has started, finished, canceled or reached a text boundary.
- [constant TTS_UTTERANCE_STARTED], [constant TTS_UTTERANCE_ENDED], and [constant TTS_UTTERANCE_CANCELED] callable's method should take one [int] parameter, the utterance ID.
- [constant TTS_UTTERANCE_BOUNDARY] callable's method should take two [int] parameters, the index of the character and the utterance ID.
[b]Note:[/b] The granularity of the boundary callbacks is engine dependent.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsSetUtteranceCallback(event gdclass.DisplayServerTTSUtteranceEvent, callable gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, event)
	callframe.Arg(frame, pointers.Get(callable))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_set_utterance_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if OS supports dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
//go:nosplit
func (self class) IsDarkModeSupported() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_is_dark_mode_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if OS is using dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
//go:nosplit
func (self class) IsDarkMode() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_is_dark_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns OS theme accent color. Returns [code]Color(0, 0, 0, 0)[/code], if accent color is unknown.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetAccentColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_accent_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the OS theme base color (default control background). Returns [code]Color(0, 0, 0, 0)[/code] if the base color is unknown.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) GetBaseColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_base_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [param callable] that should be called when system theme settings are changed. Callback method should have zero arguments.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
//go:nosplit
func (self class) SetSystemThemeChangeCallback(callable gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callable))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_set_system_theme_change_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the current mouse mode. See also [method mouse_get_mode].
*/
//go:nosplit
func (self class) MouseSetMode(mouse_mode gdclass.DisplayServerMouseMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mouse_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_mouse_set_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current mouse mode. See also [method mouse_set_mode].
*/
//go:nosplit
func (self class) MouseGetMode() gdclass.DisplayServerMouseMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.DisplayServerMouseMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_mouse_get_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the mouse cursor position to the given [param position] relative to an origin at the upper left corner of the currently focused game Window Manager window.
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS, and Linux (X11/Wayland). It has no effect on Android, iOS, and Web.
*/
//go:nosplit
func (self class) WarpMouse(position gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_warp_mouse, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the mouse cursor's current position in screen coordinates.
*/
//go:nosplit
func (self class) MouseGetPosition() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_mouse_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current state of mouse buttons (whether each button is pressed) as a bitmask. If multiple mouse buttons are pressed at the same time, the bits are added together. Equivalent to [method Input.get_mouse_button_mask].
*/
//go:nosplit
func (self class) MouseGetButtonState() MouseButtonMask {
	var frame = callframe.New()
	var r_ret = callframe.Ret[MouseButtonMask](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_mouse_get_button_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the user's clipboard content to the given string.
*/
//go:nosplit
func (self class) ClipboardSet(clipboard gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(clipboard))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the user's clipboard as a string if possible.
*/
//go:nosplit
func (self class) ClipboardGet() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_get, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the user's clipboard as an image if possible.
[b]Note:[/b] This method uses the copied pixel data, e.g. from a image editing software or a web browser, not an image file copied from file explorer.
*/
//go:nosplit
func (self class) ClipboardGetImage() [1]gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_get_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if there is a text content on the user's clipboard.
*/
//go:nosplit
func (self class) ClipboardHas() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_has, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if there is an image content on the user's clipboard.
*/
//go:nosplit
func (self class) ClipboardHasImage() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_has_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard content to the given string. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
//go:nosplit
func (self class) ClipboardSetPrimary(clipboard_primary gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(clipboard_primary))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_set_primary, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard as a string if possible. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
//go:nosplit
func (self class) ClipboardGetPrimary() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_get_primary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an [Array] of [Rect2], each of which is the bounding rectangle for a display cutout or notch. These are non-functional areas on edge-to-edge screens used by cameras and sensors. Returns an empty array if the device does not have cutouts. See also [method get_display_safe_area].
[b]Note:[/b] Currently only implemented on Android. Other platforms will return an empty array even if they do have display cutouts or notches.
*/
//go:nosplit
func (self class) GetDisplayCutouts() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_display_cutouts, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the unobscured area of the display where interactive controls should be rendered. See also [method get_display_cutouts].
*/
//go:nosplit
func (self class) GetDisplaySafeArea() gd.Rect2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_display_safe_area, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of displays available.
*/
//go:nosplit
func (self class) GetScreenCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_screen_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns index of the primary screen.
*/
//go:nosplit
func (self class) GetPrimaryScreen() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_primary_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the screen containing the window with the keyboard focus, or the primary screen if there's no focused window.
*/
//go:nosplit
func (self class) GetKeyboardFocusScreen() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_keyboard_focus_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns index of the screen which contains specified rectangle.
*/
//go:nosplit
func (self class) GetScreenFromRect(rect gd.Rect2) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_screen_from_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the screen's top-left corner position in pixels. On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
* (0, 0)        +-------+
                |       |
+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
See also [method screen_get_size].
[b]Note:[/b] On Linux (Wayland) this method always returns [code](0, 0)[/code].
*/
//go:nosplit
func (self class) ScreenGetPosition(screen gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the screen's size in pixels. See also [method screen_get_position] and [method screen_get_usable_rect].
*/
//go:nosplit
func (self class) ScreenGetSize(screen gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the portion of the screen that is not obstructed by a status bar in pixels. See also [method screen_get_size].
*/
//go:nosplit
func (self class) ScreenGetUsableRect(screen gd.Int) gd.Rect2i {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_usable_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the dots per inch density of the specified screen. If [param screen] is [constant SCREEN_OF_MAIN_WINDOW] (the default value), a screen with the main window will be used.
[b]Note:[/b] On macOS, returned value is inaccurate if fractional display scaling mode is used.
[b]Note:[/b] On Android devices, the actual screen densities are grouped into six generalized densities:
[codeblock lang=text]
   ldpi - 120 dpi
   mdpi - 160 dpi
   hdpi - 240 dpi
  xhdpi - 320 dpi
 xxhdpi - 480 dpi
xxxhdpi - 640 dpi
[/codeblock]
[b]Note:[/b] This method is implemented on Android, Linux (X11/Wayland), macOS and Windows. Returns [code]72[/code] on unsupported platforms.
*/
//go:nosplit
func (self class) ScreenGetDpi(screen gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_dpi, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scale factor of the specified screen by index.
[b]Note:[/b] On macOS, the returned value is [code]2.0[/code] for hiDPI (Retina) screens, and [code]1.0[/code] for all other cases.
[b]Note:[/b] On Linux (Wayland), the returned value is accurate only when [param screen] is [constant SCREEN_OF_MAIN_WINDOW]. Due to API limitations, passing a direct index will return a rounded-up integer, if the screen has a fractional scale (e.g. [code]1.25[/code] would get rounded up to [code]2.0[/code]).
[b]Note:[/b] This method is implemented only on macOS and Linux (Wayland).
*/
//go:nosplit
func (self class) ScreenGetScale(screen gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if touch events are available (Android or iOS), the capability is detected on the Web platform or if [member ProjectSettings.input_devices/pointing/emulate_touch_from_mouse] is [code]true[/code].
*/
//go:nosplit
func (self class) IsTouchscreenAvailable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_is_touchscreen_available, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the greatest scale factor of all screens.
[b]Note:[/b] On macOS returned value is [code]2.0[/code] if there is at least one hiDPI (Retina) screen in the system, and [code]1.0[/code] in all other cases.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) ScreenGetMaxScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_max_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current refresh rate of the specified screen. If [param screen] is [constant SCREEN_OF_MAIN_WINDOW] (the default value), a screen with the main window will be used.
[b]Note:[/b] Returns [code]-1.0[/code] if the DisplayServer fails to find the refresh rate for the specified screen. On Web, [method screen_get_refresh_rate] will always return [code]-1.0[/code] as there is no way to retrieve the refresh rate on that platform.
To fallback to a default refresh rate if the method fails, try:
[codeblock]
var refresh_rate = DisplayServer.screen_get_refresh_rate()
if refresh_rate < 0:
    refresh_rate = 60.0
[/codeblock]
*/
//go:nosplit
func (self class) ScreenGetRefreshRate(screen gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_refresh_rate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns color of the display pixel at the [param position].
[b]Note:[/b] This method is implemented on Linux (X11), macOS, and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
//go:nosplit
func (self class) ScreenGetPixel(position gd.Vector2i) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_pixel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns screenshot of the [param screen].
[b]Note:[/b] This method is implemented on Linux (X11), macOS, and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
//go:nosplit
func (self class) ScreenGetImage(screen gd.Int) [1]gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the [param screen]'s [param orientation]. See also [method screen_get_orientation].
[b]Note:[/b] On iOS, this method has no effect if [member ProjectSettings.display/window/handheld/orientation] is not set to [constant SCREEN_SENSOR].
*/
//go:nosplit
func (self class) ScreenSetOrientation(orientation gdclass.DisplayServerScreenOrientation, screen gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	callframe.Arg(frame, screen)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_set_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [param screen]'s current orientation. See also [method screen_set_orientation].
[b]Note:[/b] This method is implemented on Android and iOS.
*/
//go:nosplit
func (self class) ScreenGetOrientation(screen gd.Int) gdclass.DisplayServerScreenOrientation {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gdclass.DisplayServerScreenOrientation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the screen should never be turned off by the operating system's power-saving measures. See also [method screen_is_kept_on].
*/
//go:nosplit
func (self class) ScreenSetKeepOn(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_set_keep_on, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the screen should never be turned off by the operating system's power-saving measures. See also [method screen_set_keep_on].
*/
//go:nosplit
func (self class) ScreenIsKeptOn() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_is_kept_on, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of Godot window IDs belonging to this process.
[b]Note:[/b] Native dialogs are not included in this list.
*/
//go:nosplit
func (self class) GetWindowList() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_window_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the ID of the window at the specified screen [param position] (in pixels). On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
* (0, 0)        +-------+
                |       |
+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
*/
//go:nosplit
func (self class) GetWindowAtScreenPosition(position gd.Vector2i) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_window_at_screen_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns internal structure pointers for use in plugins.
[b]Note:[/b] This method is implemented on Android, Linux (X11/Wayland), macOS, and Windows.
*/
//go:nosplit
func (self class) WindowGetNativeHandle(handle_type gdclass.DisplayServerHandleType, window_id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, handle_type)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_native_handle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns ID of the active popup window, or [constant INVALID_WINDOW_ID] if there is none.
*/
//go:nosplit
func (self class) WindowGetActivePopup() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_active_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system. Clicking this area will not auto-close this popup.
*/
//go:nosplit
func (self class) WindowSetPopupSafeRect(window gd.Int, rect gd.Rect2i) {
	var frame = callframe.New()
	callframe.Arg(frame, window)
	callframe.Arg(frame, rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_popup_safe_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system.
*/
//go:nosplit
func (self class) WindowGetPopupSafeRect(window gd.Int) gd.Rect2i {
	var frame = callframe.New()
	callframe.Arg(frame, window)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_popup_safe_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the title of the given window to [param title].
[b]Note:[/b] It's recommended to change this value using [member Window.title] instead.
[b]Note:[/b] Avoid changing the window title every frame, as this can cause performance issues on certain window managers. Try to change the window title only a few times per second at most.
*/
//go:nosplit
func (self class) WindowSetTitle(title gd.String, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the estimated window title bar size (including text and window buttons) for the window specified by [param window_id] (in pixels). This method does not change the window title.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) WindowGetTitleSize(title gd.String, window_id gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_title_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a polygonal region of the window which accepts mouse events. Mouse events outside the region will be passed through.
Passing an empty array will disable passthrough support (all mouse events will be intercepted by the window, which is the default behavior).
[codeblocks]
[gdscript]
# Set region, using Path2D node.
DisplayServer.window_set_mouse_passthrough($Path2D.curve.get_baked_points())

# Set region, using Polygon2D node.
DisplayServer.window_set_mouse_passthrough($Polygon2D.polygon)

# Reset region to default.
DisplayServer.window_set_mouse_passthrough([])
[/gdscript]
[csharp]
// Set region, using Path2D node.
DisplayServer.WindowSetMousePassthrough(GetNode<Path2D>("Path2D").Curve.GetBakedPoints());

// Set region, using Polygon2D node.
DisplayServer.WindowSetMousePassthrough(GetNode<Polygon2D>("Polygon2D").Polygon);

// Reset region to default.
DisplayServer.WindowSetMousePassthrough(new Vector2[] {});
[/csharp]
[/codeblocks]
[b]Note:[/b] On Windows, the portion of a window that lies outside the region is not drawn, while on Linux (X11) and macOS it is.
[b]Note:[/b] This method is implemented on Linux (X11), macOS and Windows.
*/
//go:nosplit
func (self class) WindowSetMousePassthrough(region gd.PackedVector2Array, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(region))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_mouse_passthrough, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the screen the window specified by [param window_id] is currently positioned on. If the screen overlaps multiple displays, the screen where the window's center is located is returned. See also [method window_set_current_screen].
*/
//go:nosplit
func (self class) WindowGetCurrentScreen(window_id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_current_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves the window specified by [param window_id] to the specified [param screen]. See also [method window_get_current_screen].
*/
//go:nosplit
func (self class) WindowSetCurrentScreen(screen gd.Int, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_current_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the client area of the given window on the screen.
*/
//go:nosplit
func (self class) WindowGetPosition(window_id gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position of the given window on the screen including the borders drawn by the operating system. See also [method window_get_position].
*/
//go:nosplit
func (self class) WindowGetPositionWithDecorations(window_id gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_position_with_decorations, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the given window to [param position]. On multi-monitor setups, the screen position is relative to the virtual desktop area. On multi-monitor setups with different screen resolutions or orientations, the origin may be located outside any display like this:
[codeblock lang=text]
* (0, 0)        +-------+
                |       |
+-------------+ |       |
|             | |       |
|             | |       |
+-------------+ +-------+
[/codeblock]
See also [method window_get_position] and [method window_set_size].
[b]Note:[/b] It's recommended to change this value using [member Window.position] instead.
[b]Note:[/b] On Linux (Wayland): this method is a no-op.
*/
//go:nosplit
func (self class) WindowSetPosition(position gd.Vector2i, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the size of the window specified by [param window_id] (in pixels), excluding the borders drawn by the operating system. This is also called the "client area". See also [method window_get_size_with_decorations], [method window_set_size] and [method window_get_position].
*/
//go:nosplit
func (self class) WindowGetSize(window_id gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the size of the given window to [param size] (in pixels). See also [method window_get_size] and [method window_get_position].
[b]Note:[/b] It's recommended to change this value using [member Window.size] instead.
*/
//go:nosplit
func (self class) WindowSetSize(size gd.Vector2i, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [param callback] that will be called when the window specified by [param window_id] is moved or resized.
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
//go:nosplit
func (self class) WindowSetRectChangedCallback(callback gd.Callable, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_rect_changed_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [param callback] that will be called when an event occurs in the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
//go:nosplit
func (self class) WindowSetWindowEventCallback(callback gd.Callable, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_window_event_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [param callback] that should be called when any [InputEvent] is sent to the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
//go:nosplit
func (self class) WindowSetInputEventCallback(callback gd.Callable, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_input_event_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [param callback] that should be called when text is entered using the virtual keyboard to the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
//go:nosplit
func (self class) WindowSetInputTextCallback(callback gd.Callable, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_input_text_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the [param callback] that should be called when files are dropped from the operating system's file manager to the window specified by [param window_id]. [param callback] should take one [PackedStringArray] argument, which is the list of dropped files.
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
[b]Note:[/b] This method is implemented on Windows, macOS, Linux (X11/Wayland), and Web.
*/
//go:nosplit
func (self class) WindowSetDropFilesCallback(callback gd.Callable, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callback))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_drop_files_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [method Object.get_instance_id] of the [Window] the [param window_id] is attached to.
*/
//go:nosplit
func (self class) WindowGetAttachedInstanceId(window_id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_attached_instance_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the window's maximum size (in pixels). See also [method window_set_max_size].
*/
//go:nosplit
func (self class) WindowGetMaxSize(window_id gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum size of the window specified by [param window_id] in pixels. Normally, the user will not be able to drag the window to make it larger than the specified size. See also [method window_get_max_size].
[b]Note:[/b] It's recommended to change this value using [member Window.max_size] instead.
[b]Note:[/b] Using third-party tools, it is possible for users to disable window geometry restrictions and therefore bypass this limit.
*/
//go:nosplit
func (self class) WindowSetMaxSize(max_size gd.Vector2i, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_size)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the window's minimum size (in pixels). See also [method window_set_min_size].
*/
//go:nosplit
func (self class) WindowGetMinSize(window_id gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_min_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the minimum size for the given window to [param min_size] in pixels. Normally, the user will not be able to drag the window to make it smaller than the specified size. See also [method window_get_min_size].
[b]Note:[/b] It's recommended to change this value using [member Window.min_size] instead.
[b]Note:[/b] By default, the main window has a minimum size of [code]Vector2i(64, 64)[/code]. This prevents issues that can arise when the window is resized to a near-zero size.
[b]Note:[/b] Using third-party tools, it is possible for users to disable window geometry restrictions and therefore bypass this limit.
*/
//go:nosplit
func (self class) WindowSetMinSize(min_size gd.Vector2i, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, min_size)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_min_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the size of the window specified by [param window_id] (in pixels), including the borders drawn by the operating system. See also [method window_get_size].
*/
//go:nosplit
func (self class) WindowGetSizeWithDecorations(window_id gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_size_with_decorations, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the mode of the given window.
*/
//go:nosplit
func (self class) WindowGetMode(window_id gd.Int) gdclass.DisplayServerWindowMode {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gdclass.DisplayServerWindowMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets window mode for the given window to [param mode]. See [enum WindowMode] for possible values and how each mode behaves.
[b]Note:[/b] Setting the window to full screen forcibly sets the borderless flag to [code]true[/code], so make sure to set it back to [code]false[/code] when not wanted.
*/
//go:nosplit
func (self class) WindowSetMode(mode gdclass.DisplayServerWindowMode, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Enables or disables the given window's given [param flag]. See [enum WindowFlags] for possible values and their behavior.
*/
//go:nosplit
func (self class) WindowSetFlag(flag gdclass.DisplayServerWindowFlags, enabled bool, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, enabled)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current value of the given window's [param flag].
*/
//go:nosplit
func (self class) WindowGetFlag(flag gdclass.DisplayServerWindowFlags, window_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_flag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
When [constant WINDOW_FLAG_EXTEND_TO_TITLE] flag is set, set offset to the center of the first titlebar button.
[b]Note:[/b] This flag is implemented only on macOS.
*/
//go:nosplit
func (self class) WindowSetWindowButtonsOffset(offset gd.Vector2i, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_window_buttons_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns left margins ([code]x[/code]), right margins ([code]y[/code]) and height ([code]z[/code]) of the title that are safe to use (contains no buttons or other elements) when [constant WINDOW_FLAG_EXTEND_TO_TITLE] flag is set.
*/
//go:nosplit
func (self class) WindowGetSafeTitleMargins(window_id gd.Int) gd.Vector3i {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gd.Vector3i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_safe_title_margins, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Makes the window specified by [param window_id] request attention, which is materialized by the window title and taskbar entry blinking until the window is focused. This usually has no visible effect if the window is currently focused. The exact behavior varies depending on the operating system.
*/
//go:nosplit
func (self class) WindowRequestAttention(window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_request_attention, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves the window specified by [param window_id] to the foreground, so that it is visible over other windows.
*/
//go:nosplit
func (self class) WindowMoveToForeground(window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_move_to_foreground, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the window specified by [param window_id] is focused.
*/
//go:nosplit
func (self class) WindowIsFocused(window_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_is_focused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if anything can be drawn in the window specified by [param window_id], [code]false[/code] otherwise. Using the [code]--disable-render-loop[/code] command line argument or a headless build will return [code]false[/code].
*/
//go:nosplit
func (self class) WindowCanDraw(window_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_can_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets window transient parent. Transient window is will be destroyed with its transient parent and will return focus to their parent when closed. The transient window is displayed on top of a non-exclusive full-screen parent window. Transient windows can't enter full-screen mode.
[b]Note:[/b] It's recommended to change this value using [member Window.transient] instead.
[b]Note:[/b] The behavior might be different depending on the platform.
*/
//go:nosplit
func (self class) WindowSetTransient(window_id gd.Int, parent_window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	callframe.Arg(frame, parent_window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_transient, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If set to [code]true[/code], this window will always stay on top of its parent window, parent window will ignore input while this window is opened.
[b]Note:[/b] On macOS, exclusive windows are confined to the same space (virtual desktop or screen) as the parent window.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) WindowSetExclusive(window_id gd.Int, exclusive bool) {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	callframe.Arg(frame, exclusive)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_exclusive, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets whether [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] should be enabled for the window specified by [param window_id]. See also [method window_set_ime_position].
*/
//go:nosplit
func (self class) WindowSetImeActive(active bool, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_ime_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the position of the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] popup for the specified [param window_id]. Only effective if [method window_set_ime_active] was set to [code]true[/code] for the specified [param window_id].
*/
//go:nosplit
func (self class) WindowSetImePosition(position gd.Vector2i, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_ime_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the V-Sync mode of the given window. See also [member ProjectSettings.display/window/vsync/vsync_mode].
See [enum DisplayServer.VSyncMode] for possible values and how they affect the behavior of your application.
Depending on the platform and used renderer, the engine will fall back to [constant VSYNC_ENABLED] if the desired mode is not supported.
[b]Note:[/b] V-Sync modes other than [constant VSYNC_ENABLED] are only supported in the Forward+ and Mobile rendering methods, not Compatibility.
*/
//go:nosplit
func (self class) WindowSetVsyncMode(vsync_mode gdclass.DisplayServerVSyncMode, window_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, vsync_mode)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_vsync_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the V-Sync mode of the given window.
*/
//go:nosplit
func (self class) WindowGetVsyncMode(window_id gd.Int) gdclass.DisplayServerVSyncMode {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[gdclass.DisplayServerVSyncMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_vsync_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given window can be maximized (the maximize button is enabled).
*/
//go:nosplit
func (self class) WindowIsMaximizeAllowed(window_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_is_maximize_allowed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if double-click on a window title should maximize it.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) WindowMaximizeOnTitleDblClick() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_maximize_on_title_dbl_click, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if double-click on a window title should minimize it.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) WindowMinimizeOnTitleDblClick() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_minimize_on_title_dbl_click, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text selection in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] composition string, with the [Vector2i]'s [code]x[/code] component being the caret position and [code]y[/code] being the length of the selection.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) ImeGetSelection() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_ime_get_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the composition string contained within the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] window.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) ImeGetText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_ime_get_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Shows the virtual keyboard if the platform has one.
[param existing_text] parameter is useful for implementing your own [LineEdit] or [TextEdit], as it tells the virtual keyboard what text has already been typed (the virtual keyboard uses it for auto-correct and predictions).
[param position] parameter is the screen space [Rect2] of the edited text.
[param type] parameter allows configuring which type of virtual keyboard to show.
[param max_length] limits the number of characters that can be entered if different from [code]-1[/code].
[param cursor_start] can optionally define the current text cursor position if [param cursor_end] is not set.
[param cursor_start] and [param cursor_end] can optionally define the current text selection.
[b]Note:[/b] This method is implemented on Android, iOS and Web.
*/
//go:nosplit
func (self class) VirtualKeyboardShow(existing_text gd.String, position gd.Rect2, atype gdclass.DisplayServerVirtualKeyboardType, max_length gd.Int, cursor_start gd.Int, cursor_end gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(existing_text))
	callframe.Arg(frame, position)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, max_length)
	callframe.Arg(frame, cursor_start)
	callframe.Arg(frame, cursor_end)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_virtual_keyboard_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Hides the virtual keyboard if it is shown, does nothing otherwise.
*/
//go:nosplit
func (self class) VirtualKeyboardHide() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_virtual_keyboard_hide, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the on-screen keyboard's height in pixels. Returns 0 if there is no keyboard or if it is currently hidden.
*/
//go:nosplit
func (self class) VirtualKeyboardGetHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_virtual_keyboard_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the default mouse cursor shape. The cursor's appearance will vary depending on the user's operating system and mouse cursor theme. See also [method cursor_get_shape] and [method cursor_set_custom_image].
*/
//go:nosplit
func (self class) CursorSetShape(shape gdclass.DisplayServerCursorShape) {
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_cursor_set_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the default mouse cursor shape set by [method cursor_set_shape].
*/
//go:nosplit
func (self class) CursorGetShape() gdclass.DisplayServerCursorShape {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.DisplayServerCursorShape](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_cursor_get_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a custom mouse cursor image for the given [param shape]. This means the user's operating system and mouse cursor theme will no longer influence the mouse cursor's appearance.
[param cursor] can be either a [Texture2D] or an [Image], and it should not be larger than 256×256 to display correctly. Optionally, [param hotspot] can be set to offset the image's position relative to the click point. By default, [param hotspot] is set to the top-left corner of the image. See also [method cursor_set_shape].
*/
//go:nosplit
func (self class) CursorSetCustomImage(cursor [1]gdclass.Resource, shape gdclass.DisplayServerCursorShape, hotspot gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(cursor[0])[0])
	callframe.Arg(frame, shape)
	callframe.Arg(frame, hotspot)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_cursor_set_custom_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if positions of [b]OK[/b] and [b]Cancel[/b] buttons are swapped in dialogs. This is enabled by default on Windows to follow interface conventions, and be toggled by changing [member ProjectSettings.gui/common/swap_cancel_ok].
[b]Note:[/b] This doesn't affect native dialogs such as the ones spawned by [method DisplayServer.dialog_show].
*/
//go:nosplit
func (self class) GetSwapCancelOk() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_swap_cancel_ok, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Allows the [param process_id] PID to steal focus from this window. In other words, this disables the operating system's focus stealing protection for the specified PID.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) EnableForStealingFocus(process_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, process_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_enable_for_stealing_focus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Shows a text dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [int] parameter which corresponds to the index of the pressed button.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG] feature. Supported platforms include macOS and Windows.
*/
//go:nosplit
func (self class) DialogShow(title gd.String, description gd.String, buttons gd.PackedStringArray, callback gd.Callable) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	callframe.Arg(frame, pointers.Get(description))
	callframe.Arg(frame, pointers.Get(buttons))
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_dialog_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Shows a text input dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [String] parameter which contains the text field's contents.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_INPUT] feature. Supported platforms include macOS and Windows.
*/
//go:nosplit
func (self class) DialogInputText(title gd.String, description gd.String, existing_text gd.String, callback gd.Callable) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	callframe.Arg(frame, pointers.Get(description))
	callframe.Arg(frame, pointers.Get(existing_text))
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_dialog_input_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Displays OS native dialog for selecting files or directories in the file system.
Each filter string in the [param filters] array should be formatted like this: [code]*.txt,*.doc;Text Files[/code]. The description text of the filter is optional and can be omitted. See also [member FileDialog.filters].
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int[/code].
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE] feature. Supported platforms include Linux (X11/Wayland), Windows, and macOS.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] On Linux, [param show_hidden] is ignored.
[b]Note:[/b] On macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
//go:nosplit
func (self class) FileDialogShow(title gd.String, current_directory gd.String, filename gd.String, show_hidden bool, mode gdclass.DisplayServerFileDialogMode, filters gd.PackedStringArray, callback gd.Callable) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	callframe.Arg(frame, pointers.Get(current_directory))
	callframe.Arg(frame, pointers.Get(filename))
	callframe.Arg(frame, show_hidden)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, pointers.Get(filters))
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_file_dialog_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Displays OS native dialog for selecting files or directories in the file system with additional user selectable options.
Each filter string in the [param filters] array should be formatted like this: [code]*.txt,*.doc;Text Files[/code]. The description text of the filter is optional and can be omitted. See also [member FileDialog.filters].
[param options] is array of [Dictionary]s with the following keys:
- [code]"name"[/code] - option's name [String].
- [code]"values"[/code] - [PackedStringArray] of values. If empty, boolean option (check box) is used.
- [code]"default"[/code] - default selected option index ([int]) or default boolean value ([bool]).
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int, selected_option: Dictionary[/code].
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE] feature. Supported platforms include Linux (X11/Wayland), Windows, and macOS.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] On Linux (X11), [param show_hidden] is ignored.
[b]Note:[/b] On macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
//go:nosplit
func (self class) FileDialogWithOptionsShow(title gd.String, current_directory gd.String, root gd.String, filename gd.String, show_hidden bool, mode gdclass.DisplayServerFileDialogMode, filters gd.PackedStringArray, options gd.Array, callback gd.Callable) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(title))
	callframe.Arg(frame, pointers.Get(current_directory))
	callframe.Arg(frame, pointers.Get(root))
	callframe.Arg(frame, pointers.Get(filename))
	callframe.Arg(frame, show_hidden)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, pointers.Get(filters))
	callframe.Arg(frame, pointers.Get(options))
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_file_dialog_with_options_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of keyboard layouts.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLayoutCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_layout_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns active keyboard layout index.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS, and Windows.
*/
//go:nosplit
func (self class) KeyboardGetCurrentLayout() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_current_layout, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardSetCurrentLayout(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_set_current_layout, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the ISO-639/BCP-47 language code of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLayoutLanguage(index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_layout_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the localized name of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLayoutName(index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_layout_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Converts a physical (US QWERTY) [param keycode] to one in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetKeycodeFromPhysical(keycode Key) Key {
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Ret[Key](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_keycode_from_physical, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Converts a physical (US QWERTY) [param keycode] to localized label printed on the key in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLabelFromPhysical(keycode Key) Key {
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Ret[Key](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_label_from_physical, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Perform window manager processing, including input flushing. See also [method force_process_and_drop_events], [method Input.flush_buffered_events] and [member Input.use_accumulated_input].
*/
//go:nosplit
func (self class) ProcessEvents() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_process_events, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Forces window manager processing while ignoring all [InputEvent]s. See also [method process_events].
[b]Note:[/b] This method is implemented on Windows and macOS.
*/
//go:nosplit
func (self class) ForceProcessAndDropEvents() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_force_process_and_drop_events, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the window icon (usually displayed in the top-left corner) in the operating system's [i]native[/i] format. The file at [param filename] must be in [code].ico[/code] format on Windows or [code].icns[/code] on macOS. By using specially crafted [code].ico[/code] or [code].icns[/code] icons, [method set_native_icon] allows specifying different icons depending on the size the icon is displayed at. This size is determined by the operating system and user preferences (including the display scale factor). To use icons in other formats, use [method set_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_NATIVE_ICON].
*/
//go:nosplit
func (self class) SetNativeIcon(filename gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(filename))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_set_native_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the window icon (usually displayed in the top-left corner) with an [Image]. To use icons in the operating system's native format, use [method set_native_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_ICON].
*/
//go:nosplit
func (self class) SetIcon(image [1]gdclass.Image) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a new application status indicator with the specified icon, tooltip, and activation callback.
[param callback] should take two arguments: the pressed mouse button (one of the [enum MouseButton] constants) and the click position in screen coordinates (a [Vector2i]).
*/
//go:nosplit
func (self class) CreateStatusIndicator(icon [1]gdclass.Texture2D, tooltip gd.String, callback gd.Callable) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(tooltip))
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_create_status_indicator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the application status indicator icon.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorSetIcon(id gd.Int, icon [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_status_indicator_set_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the application status indicator tooltip.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorSetTooltip(id gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_status_indicator_set_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the application status indicator native popup menu.
[b]Note:[/b] On macOS, the menu is activated by any mouse button. Its activation callback is [i]not[/i] triggered.
[b]Note:[/b] On Windows, the menu is activated by the right mouse button, selecting the status icon and pressing [kbd]Shift + F10[/kbd], or the applications key. The menu's activation callback for the other mouse buttons is still triggered.
[b]Note:[/b] Native popup is only supported if [NativeMenu] supports the [constant NativeMenu.FEATURE_POPUP_MENU] feature.
*/
//go:nosplit
func (self class) StatusIndicatorSetMenu(id gd.Int, menu_rid gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, menu_rid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_status_indicator_set_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the application status indicator activation callback. [param callback] should take two arguments: [int] mouse button index (one of [enum MouseButton] values) and [Vector2i] click position in screen coordinates.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorSetCallback(id gd.Int, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_status_indicator_set_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the rectangle for the given status indicator [param id] in screen coordinates. If the status indicator is not visible, returns an empty [Rect2].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorGetRect(id gd.Int) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_status_indicator_get_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the application status indicator.
*/
//go:nosplit
func (self class) DeleteStatusIndicator(id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_delete_status_indicator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the total number of available tablet drivers.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletGetDriverCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tablet_get_driver_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tablet driver name for the given index.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletGetDriverName(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tablet_get_driver_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns current active tablet driver name.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletGetCurrentDriver() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tablet_get_current_driver, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Set active tablet driver name.
Supported drivers:
- [code]winink[/code]: Windows Ink API, default (Windows 8.1+ required).
- [code]wintab[/code]: Wacom Wintab API (compatible device driver required).
- [code]dummy[/code]: Dummy driver, tablet input is disabled.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletSetCurrentDriver(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tablet_set_current_driver, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the window background can be made transparent. This method returns [code]false[/code] if [member ProjectSettings.display/window/per_pixel_transparency/allowed] is set to [code]false[/code], or if transparency is not supported by the renderer or OS compositor.
*/
//go:nosplit
func (self class) IsWindowTransparencyAvailable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_is_window_transparency_available, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Registers an [Object] which represents an additional output that will be rendered too, beyond normal windows. The [Object] is only used as an identifier, which can be later passed to [method unregister_additional_output].
This can be used to prevent Godot from skipping rendering when no normal windows are visible.
*/
//go:nosplit
func (self class) RegisterAdditionalOutput(obj [1]gd.Object) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_register_additional_output, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters an [Object] representing an additional output, that was registered via [method register_additional_output].
*/
//go:nosplit
func (self class) UnregisterAdditionalOutput(obj [1]gd.Object) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_unregister_additional_output, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if any additional outputs have been registered via [method register_additional_output].
*/
//go:nosplit
func (self class) HasAdditionalOutputs() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_has_additional_outputs, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("DisplayServer", func(ptr gd.Object) any {
		return [1]gdclass.DisplayServer{*(*gdclass.DisplayServer)(unsafe.Pointer(&ptr))}
	})
}

type Feature = gdclass.DisplayServerFeature

const (
	/*Display server supports global menu. This allows the application to display its menu items in the operating system's top bar. [b]macOS[/b]*/
	FeatureGlobalMenu Feature = 0
	/*Display server supports multiple windows that can be moved outside of the main window. [b]Windows, macOS, Linux (X11)[/b]*/
	FeatureSubwindows Feature = 1
	/*Display server supports touchscreen input. [b]Windows, Linux (X11), Android, iOS, Web[/b]*/
	FeatureTouchscreen Feature = 2
	/*Display server supports mouse input. [b]Windows, macOS, Linux (X11/Wayland), Android, Web[/b]*/
	FeatureMouse Feature = 3
	/*Display server supports warping mouse coordinates to keep the mouse cursor constrained within an area, but looping when one of the edges is reached. [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureMouseWarp Feature = 4
	/*Display server supports setting and getting clipboard data. See also [constant FEATURE_CLIPBOARD_PRIMARY]. [b]Windows, macOS, Linux (X11/Wayland), Android, iOS, Web[/b]*/
	FeatureClipboard Feature = 5
	/*Display server supports popping up a virtual keyboard when requested to input text without a physical keyboard. [b]Android, iOS, Web[/b]*/
	FeatureVirtualKeyboard Feature = 6
	/*Display server supports setting the mouse cursor shape to be different from the default. [b]Windows, macOS, Linux (X11/Wayland), Android, Web[/b]*/
	FeatureCursorShape Feature = 7
	/*Display server supports setting the mouse cursor shape to a custom image. [b]Windows, macOS, Linux (X11/Wayland), Web[/b]*/
	FeatureCustomCursorShape Feature = 8
	/*Display server supports spawning text dialogs using the operating system's native look-and-feel. See [method dialog_show]. [b]Windows, macOS[/b]*/
	FeatureNativeDialog Feature = 9
	/*Display server supports [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url], which is commonly used for inputting Chinese/Japanese/Korean text. This is handled by the operating system, rather than by Godot. [b]Windows, macOS, Linux (X11)[/b]*/
	FeatureIme Feature = 10
	/*Display server supports windows can use per-pixel transparency to make windows behind them partially or fully visible. [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureWindowTransparency Feature = 11
	/*Display server supports querying the operating system's display scale factor. This allows for [i]reliable[/i] automatic hiDPI display detection, as opposed to guessing based on the screen resolution and reported display DPI (which can be unreliable due to broken monitor EDID). [b]Windows, Linux (Wayland), macOS[/b]*/
	FeatureHidpi Feature = 12
	/*Display server supports changing the window icon (usually displayed in the top-left corner). [b]Windows, macOS, Linux (X11)[/b]*/
	FeatureIcon Feature = 13
	/*Display server supports changing the window icon (usually displayed in the top-left corner). [b]Windows, macOS[/b]*/
	FeatureNativeIcon Feature = 14
	/*Display server supports changing the screen orientation. [b]Android, iOS[/b]*/
	FeatureOrientation Feature = 15
	/*Display server supports V-Sync status can be changed from the default (which is forced to be enabled platforms not supporting this feature). [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureSwapBuffers Feature = 16
	/*Display server supports Primary clipboard can be used. This is a different clipboard from [constant FEATURE_CLIPBOARD]. [b]Linux (X11/Wayland)[/b]*/
	FeatureClipboardPrimary Feature = 18
	/*Display server supports text-to-speech. See [code]tts_*[/code] methods. [b]Windows, macOS, Linux (X11/Wayland), Android, iOS, Web[/b]*/
	FeatureTextToSpeech Feature = 19
	/*Display server supports expanding window content to the title. See [constant WINDOW_FLAG_EXTEND_TO_TITLE]. [b]macOS[/b]*/
	FeatureExtendToTitle Feature = 20
	/*Display server supports reading screen pixels. See [method screen_get_pixel].*/
	FeatureScreenCapture Feature = 21
	/*Display server supports application status indicators.*/
	FeatureStatusIndicator Feature = 22
	/*Display server supports native help system search callbacks. See [method help_set_search_callbacks].*/
	FeatureNativeHelp Feature = 23
	/*Display server supports spawning text input dialogs using the operating system's native look-and-feel. See [method dialog_input_text]. [b]Windows, macOS[/b]*/
	FeatureNativeDialogInput Feature = 24
	/*Display server supports spawning dialogs for selecting files or directories using the operating system's native look-and-feel. See [method file_dialog_show] and [method file_dialog_with_options_show]. [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureNativeDialogFile Feature = 25
)

type MouseModeValue = gdclass.DisplayServerMouseMode

const (
	/*Makes the mouse cursor visible if it is hidden.*/
	MouseModeVisible MouseModeValue = 0
	/*Makes the mouse cursor hidden if it is visible.*/
	MouseModeHidden MouseModeValue = 1
	/*Captures the mouse. The mouse will be hidden and its position locked at the center of the window manager's window.
	  [b]Note:[/b] If you want to process the mouse's movement in this mode, you need to use [member InputEventMouseMotion.relative].*/
	MouseModeCaptured MouseModeValue = 2
	/*Confines the mouse cursor to the game window, and make it visible.*/
	MouseModeConfined MouseModeValue = 3
	/*Confines the mouse cursor to the game window, and make it hidden.*/
	MouseModeConfinedHidden MouseModeValue = 4
)

type ScreenOrientation = gdclass.DisplayServerScreenOrientation

const (
	/*Default landscape orientation.*/
	ScreenLandscape ScreenOrientation = 0
	/*Default portrait orientation.*/
	ScreenPortrait ScreenOrientation = 1
	/*Reverse landscape orientation (upside down).*/
	ScreenReverseLandscape ScreenOrientation = 2
	/*Reverse portrait orientation (upside down).*/
	ScreenReversePortrait ScreenOrientation = 3
	/*Automatic landscape orientation (default or reverse depending on sensor).*/
	ScreenSensorLandscape ScreenOrientation = 4
	/*Automatic portrait orientation (default or reverse depending on sensor).*/
	ScreenSensorPortrait ScreenOrientation = 5
	/*Automatic landscape or portrait orientation (default or reverse depending on sensor).*/
	ScreenSensor ScreenOrientation = 6
)

type VirtualKeyboardType = gdclass.DisplayServerVirtualKeyboardType

const (
	/*Default text virtual keyboard.*/
	KeyboardTypeDefault VirtualKeyboardType = 0
	/*Multiline virtual keyboard.*/
	KeyboardTypeMultiline VirtualKeyboardType = 1
	/*Virtual number keypad, useful for PIN entry.*/
	KeyboardTypeNumber VirtualKeyboardType = 2
	/*Virtual number keypad, useful for entering fractional numbers.*/
	KeyboardTypeNumberDecimal VirtualKeyboardType = 3
	/*Virtual phone number keypad.*/
	KeyboardTypePhone VirtualKeyboardType = 4
	/*Virtual keyboard with additional keys to assist with typing email addresses.*/
	KeyboardTypeEmailAddress VirtualKeyboardType = 5
	/*Virtual keyboard for entering a password. On most platforms, this should disable autocomplete and autocapitalization.
	  [b]Note:[/b] This is not supported on Web. Instead, this behaves identically to [constant KEYBOARD_TYPE_DEFAULT].*/
	KeyboardTypePassword VirtualKeyboardType = 6
	/*Virtual keyboard with additional keys to assist with typing URLs.*/
	KeyboardTypeUrl VirtualKeyboardType = 7
)

type CursorShape = gdclass.DisplayServerCursorShape

const (
	/*Arrow cursor shape. This is the default when not pointing anything that overrides the mouse cursor, such as a [LineEdit] or [TextEdit].*/
	CursorArrow CursorShape = 0
	/*I-beam cursor shape. This is used by default when hovering a control that accepts text input, such as [LineEdit] or [TextEdit].*/
	CursorIbeam CursorShape = 1
	/*Pointing hand cursor shape. This is used by default when hovering a [LinkButton] or a URL tag in a [RichTextLabel].*/
	CursorPointingHand CursorShape = 2
	/*Crosshair cursor. This is intended to be displayed when the user needs precise aim over an element, such as a rectangle selection tool or a color picker.*/
	CursorCross CursorShape = 3
	/*Wait cursor. On most cursor themes, this displays a spinning icon [i]besides[/i] the arrow. Intended to be used for non-blocking operations (when the user can do something else at the moment). See also [constant CURSOR_BUSY].*/
	CursorWait CursorShape = 4
	/*Wait cursor. On most cursor themes, this [i]replaces[/i] the arrow with a spinning icon. Intended to be used for blocking operations (when the user can't do anything else at the moment). See also [constant CURSOR_WAIT].*/
	CursorBusy CursorShape = 5
	/*Dragging hand cursor. This is displayed during drag-and-drop operations. See also [constant CURSOR_CAN_DROP].*/
	CursorDrag CursorShape = 6
	/*"Can drop" cursor. This is displayed during drag-and-drop operations if hovering over a [Control] that can accept the drag-and-drop event. On most cursor themes, this displays a dragging hand with an arrow symbol besides it. See also [constant CURSOR_DRAG].*/
	CursorCanDrop CursorShape = 7
	/*Forbidden cursor. This is displayed during drag-and-drop operations if the hovered [Control] can't accept the drag-and-drop event.*/
	CursorForbidden CursorShape = 8
	/*Vertical resize cursor. Intended to be displayed when the hovered [Control] can be vertically resized using the mouse. See also [constant CURSOR_VSPLIT].*/
	CursorVsize CursorShape = 9
	/*Horizontal resize cursor. Intended to be displayed when the hovered [Control] can be horizontally resized using the mouse. See also [constant CURSOR_HSPLIT].*/
	CursorHsize CursorShape = 10
	/*Secondary diagonal resize cursor (top-right/bottom-left). Intended to be displayed when the hovered [Control] can be resized on both axes at once using the mouse.*/
	CursorBdiagsize CursorShape = 11
	/*Main diagonal resize cursor (top-left/bottom-right). Intended to be displayed when the hovered [Control] can be resized on both axes at once using the mouse.*/
	CursorFdiagsize CursorShape = 12
	/*Move cursor. Intended to be displayed when the hovered [Control] can be moved using the mouse.*/
	CursorMove CursorShape = 13
	/*Vertical split cursor. This is displayed when hovering a [Control] with splits that can be vertically resized using the mouse, such as [VSplitContainer]. On some cursor themes, this cursor may have the same appearance as [constant CURSOR_VSIZE].*/
	CursorVsplit CursorShape = 14
	/*Horizontal split cursor. This is displayed when hovering a [Control] with splits that can be horizontally resized using the mouse, such as [HSplitContainer]. On some cursor themes, this cursor may have the same appearance as [constant CURSOR_HSIZE].*/
	CursorHsplit CursorShape = 15
	/*Help cursor. On most cursor themes, this displays a question mark icon instead of the mouse cursor. Intended to be used when the user has requested help on the next element that will be clicked.*/
	CursorHelp CursorShape = 16
	/*Represents the size of the [enum CursorShape] enum.*/
	CursorMax CursorShape = 17
)

type FileDialogMode = gdclass.DisplayServerFileDialogMode

const (
	/*The native file dialog allows selecting one, and only one file.*/
	FileDialogModeOpenFile FileDialogMode = 0
	/*The native file dialog allows selecting multiple files.*/
	FileDialogModeOpenFiles FileDialogMode = 1
	/*The native file dialog only allows selecting a directory, disallowing the selection of any file.*/
	FileDialogModeOpenDir FileDialogMode = 2
	/*The native file dialog allows selecting one file or directory.*/
	FileDialogModeOpenAny FileDialogMode = 3
	/*The native file dialog will warn when a file exists.*/
	FileDialogModeSaveFile FileDialogMode = 4
)

type WindowMode = gdclass.DisplayServerWindowMode

const (
	/*Windowed mode, i.e. [Window] doesn't occupy the whole screen (unless set to the size of the screen).*/
	WindowModeWindowed WindowMode = 0
	/*Minimized window mode, i.e. [Window] is not visible and available on window manager's window list. Normally happens when the minimize button is pressed.*/
	WindowModeMinimized WindowMode = 1
	/*Maximized window mode, i.e. [Window] will occupy whole screen area except task bar and still display its borders. Normally happens when the maximize button is pressed.*/
	WindowModeMaximized WindowMode = 2
	/*Full screen mode with full multi-window support.
	  Full screen window covers the entire display area of a screen and has no decorations. The display's video mode is not changed.
	  [b]On Windows:[/b] Multi-window full-screen mode has a 1px border of the [member ProjectSettings.rendering/environment/defaults/default_clear_color] color.
	  [b]On macOS:[/b] A new desktop is used to display the running project.
	  [b]Note:[/b] Regardless of the platform, enabling full screen will change the window size to match the monitor's size. Therefore, make sure your project supports [url=$DOCS_URL/tutorials/rendering/multiple_resolutions.html]multiple resolutions[/url] when enabling full screen mode.*/
	WindowModeFullscreen WindowMode = 3
	/*A single window full screen mode. This mode has less overhead, but only one window can be open on a given screen at a time (opening a child window or application switching will trigger a full screen transition).
	  Full screen window covers the entire display area of a screen and has no border or decorations. The display's video mode is not changed.
	  [b]On Windows:[/b] Depending on video driver, full screen transition might cause screens to go black for a moment.
	  [b]On macOS:[/b] A new desktop is used to display the running project. Exclusive full screen mode prevents Dock and Menu from showing up when the mouse pointer is hovering the edge of the screen.
	  [b]On Linux (X11):[/b] Exclusive full screen mode bypasses compositor.
	  [b]Note:[/b] Regardless of the platform, enabling full screen will change the window size to match the monitor's size. Therefore, make sure your project supports [url=$DOCS_URL/tutorials/rendering/multiple_resolutions.html]multiple resolutions[/url] when enabling full screen mode.*/
	WindowModeExclusiveFullscreen WindowMode = 4
)

type WindowFlags = gdclass.DisplayServerWindowFlags

const (
	/*The window can't be resized by dragging its resize grip. It's still possible to resize the window using [method window_set_size]. This flag is ignored for full screen windows.*/
	WindowFlagResizeDisabled WindowFlags = 0
	/*The window do not have native title bar and other decorations. This flag is ignored for full-screen windows.*/
	WindowFlagBorderless WindowFlags = 1
	/*The window is floating on top of all other windows. This flag is ignored for full-screen windows.*/
	WindowFlagAlwaysOnTop WindowFlags = 2
	/*The window background can be transparent.
	  [b]Note:[/b] This flag has no effect if [method is_window_transparency_available] returns [code]false[/code].
	  [b]Note:[/b] Transparency support is implemented on Linux (X11/Wayland), macOS, and Windows, but availability might vary depending on GPU driver, display manager, and compositor capabilities.*/
	WindowFlagTransparent WindowFlags = 3
	/*The window can't be focused. No-focus window will ignore all input, except mouse clicks.*/
	WindowFlagNoFocus WindowFlags = 4
	/*Window is part of menu or [OptionButton] dropdown. This flag can't be changed when the window is visible. An active popup window will exclusively receive all input, without stealing focus from its parent. Popup windows are automatically closed when uses click outside it, or when an application is switched. Popup window must have transient parent set (see [method window_set_transient]).*/
	WindowFlagPopup WindowFlags = 5
	/*Window content is expanded to the full size of the window. Unlike borderless window, the frame is left intact and can be used to resize the window, title bar is transparent, but have minimize/maximize/close buttons.
	  Use [method window_set_window_buttons_offset] to adjust minimize/maximize/close buttons offset.
	  Use [method window_get_safe_title_margins] to determine area under the title bar that is not covered by decorations.
	  [b]Note:[/b] This flag is implemented only on macOS.*/
	WindowFlagExtendToTitle WindowFlags = 6
	/*All mouse events are passed to the underlying window of the same application.*/
	WindowFlagMousePassthrough WindowFlags = 7
	/*Max value of the [enum WindowFlags].*/
	WindowFlagMax WindowFlags = 8
)

type WindowEvent = gdclass.DisplayServerWindowEvent

const (
	/*Sent when the mouse pointer enters the window.*/
	WindowEventMouseEnter WindowEvent = 0
	/*Sent when the mouse pointer exits the window.*/
	WindowEventMouseExit WindowEvent = 1
	/*Sent when the window grabs focus.*/
	WindowEventFocusIn WindowEvent = 2
	/*Sent when the window loses focus.*/
	WindowEventFocusOut WindowEvent = 3
	/*Sent when the user has attempted to close the window (e.g. close button is pressed).*/
	WindowEventCloseRequest WindowEvent = 4
	/*Sent when the device "Back" button is pressed.
	  [b]Note:[/b] This event is implemented only on Android.*/
	WindowEventGoBackRequest WindowEvent = 5
	/*Sent when the window is moved to the display with different DPI, or display DPI is changed.
	  [b]Note:[/b] This flag is implemented only on macOS.*/
	WindowEventDpiChange WindowEvent = 6
	/*Sent when the window title bar decoration is changed (e.g. [constant WINDOW_FLAG_EXTEND_TO_TITLE] is set or window entered/exited full screen mode).
	  [b]Note:[/b] This flag is implemented only on macOS.*/
	WindowEventTitlebarChange WindowEvent = 7
)

type VSyncMode = gdclass.DisplayServerVSyncMode

const (
	/*No vertical synchronization, which means the engine will display frames as fast as possible (tearing may be visible). Framerate is unlimited (regardless of [member Engine.max_fps]).*/
	VsyncDisabled VSyncMode = 0
	/*Default vertical synchronization mode, the image is displayed only on vertical blanking intervals (no tearing is visible). Framerate is limited by the monitor refresh rate (regardless of [member Engine.max_fps]).*/
	VsyncEnabled VSyncMode = 1
	/*Behaves like [constant VSYNC_DISABLED] when the framerate drops below the screen's refresh rate to reduce stuttering (tearing may be visible). Otherwise, vertical synchronization is enabled to avoid tearing. Framerate is limited by the monitor refresh rate (regardless of [member Engine.max_fps]). Behaves like [constant VSYNC_ENABLED] when using the Compatibility rendering method.*/
	VsyncAdaptive VSyncMode = 2
	/*Displays the most recent image in the queue on vertical blanking intervals, while rendering to the other images (no tearing is visible). Framerate is unlimited (regardless of [member Engine.max_fps]).
	  Although not guaranteed, the images can be rendered as fast as possible, which may reduce input lag (also called "Fast" V-Sync mode). [constant VSYNC_MAILBOX] works best when at least twice as many frames as the display refresh rate are rendered. Behaves like [constant VSYNC_ENABLED] when using the Compatibility rendering method.*/
	VsyncMailbox VSyncMode = 3
)

type HandleType = gdclass.DisplayServerHandleType

const (
	/*Display handle:
	  - Linux (X11): [code]X11::Display*[/code] for the display.
	  - Android: [code]EGLDisplay[/code] for the display.*/
	DisplayHandle HandleType = 0
	/*Window handle:
	  - Windows: [code]HWND[/code] for the window.
	  - Linux (X11): [code]X11::Window*[/code] for the window.
	  - macOS: [code]NSWindow*[/code] for the window.
	  - iOS: [code]UIViewController*[/code] for the view controller.
	  - Android: [code]jObject[/code] for the activity.*/
	WindowHandle HandleType = 1
	/*Window view:
	  - Windows: [code]HDC[/code] for the window (only with the GL Compatibility renderer).
	  - macOS: [code]NSView*[/code] for the window main view.
	  - iOS: [code]UIView*[/code] for the window main view.*/
	WindowView HandleType = 2
	/*OpenGL context (only with the GL Compatibility renderer):
	  - Windows: [code]HGLRC[/code] for the window (native GL), or [code]EGLContext[/code] for the window (ANGLE).
	  - Linux (X11): [code]GLXContext*[/code] for the window.
	  - macOS: [code]NSOpenGLContext*[/code] for the window (native GL), or [code]EGLContext[/code] for the window (ANGLE).
	  - Android: [code]EGLContext[/code] for the window.*/
	OpenglContext HandleType = 3
)

type TTSUtteranceEvent = gdclass.DisplayServerTTSUtteranceEvent

const (
	/*Utterance has begun to be spoken.*/
	TtsUtteranceStarted TTSUtteranceEvent = 0
	/*Utterance was successfully finished.*/
	TtsUtteranceEnded TTSUtteranceEvent = 1
	/*Utterance was canceled, or TTS service was unable to process it.*/
	TtsUtteranceCanceled TTSUtteranceEvent = 2
	/*Utterance reached a word or sentence boundary.*/
	TtsUtteranceBoundary TTSUtteranceEvent = 3
)

type Error = gd.Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
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

type MouseButton int

const (
	/*Enum value which doesn't correspond to any mouse button. This is used to initialize [enum MouseButton] properties with a generic state.*/
	MouseButtonNone MouseButton = 0
	/*Primary mouse button, usually assigned to the left button.*/
	MouseButtonLeft MouseButton = 1
	/*Secondary mouse button, usually assigned to the right button.*/
	MouseButtonRight MouseButton = 2
	/*Middle mouse button.*/
	MouseButtonMiddle MouseButton = 3
	/*Mouse wheel scrolling up.*/
	MouseButtonWheelUp MouseButton = 4
	/*Mouse wheel scrolling down.*/
	MouseButtonWheelDown MouseButton = 5
	/*Mouse wheel left button (only present on some mice).*/
	MouseButtonWheelLeft MouseButton = 6
	/*Mouse wheel right button (only present on some mice).*/
	MouseButtonWheelRight MouseButton = 7
	/*Extra mouse button 1. This is sometimes present, usually to the sides of the mouse.*/
	MouseButtonXbutton1 MouseButton = 8
	/*Extra mouse button 2. This is sometimes present, usually to the sides of the mouse.*/
	MouseButtonXbutton2 MouseButton = 9
)

type MouseButtonMask int

const (
	/*Primary mouse button mask, usually for the left button.*/
	MouseButtonMaskLeft MouseButtonMask = 1
	/*Secondary mouse button mask, usually for the right button.*/
	MouseButtonMaskRight MouseButtonMask = 2
	/*Middle mouse button mask.*/
	MouseButtonMaskMiddle MouseButtonMask = 4
	/*Extra mouse button 1 mask.*/
	MouseButtonMaskMbXbutton1 MouseButtonMask = 128
	/*Extra mouse button 2 mask.*/
	MouseButtonMaskMbXbutton2 MouseButtonMask = 256
)
