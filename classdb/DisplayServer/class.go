// Package DisplayServer provides methods for working with DisplayServer object instances.
package DisplayServer

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Rect2i"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Vector3i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

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
func HasFeature(feature gdclass.DisplayServerFeature) bool { //gd:DisplayServer.has_feature
	once.Do(singleton)
	return bool(class(self).HasFeature(feature))
}

/*
Returns the name of the [DisplayServer] currently in use. Most operating systems only have a single [DisplayServer], but Linux has access to more than one [DisplayServer] (currently X11 and Wayland).
The names of built-in display servers are [code]Windows[/code], [code]macOS[/code], [code]X11[/code] (Linux), [code]Wayland[/code] (Linux), [code]Android[/code], [code]iOS[/code], [code]web[/code] (HTML5), and [code]headless[/code] (when started with the [code]--headless[/code] [url=$DOCS_URL/tutorials/editor/command_line_tutorial.html]command line argument[/url]).
*/
func GetName() string { //gd:DisplayServer.get_name
	once.Do(singleton)
	return string(class(self).GetName().String())
}

/*
Sets native help system search callbacks.
[param search_callback] has the following arguments: [code]String search_string, int result_limit[/code] and return a [Dictionary] with "key, display name" pairs for the search results. Called when the user enters search terms in the [code]Help[/code] menu.
[param action_callback] has the following arguments: [code]String key[/code]. Called when the user selects a search result in the [code]Help[/code] menu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func HelpSetSearchCallbacks(search_callback func(search_string string, result_limit int) map[any]any, action_callback func(key string)) { //gd:DisplayServer.help_set_search_callbacks
	once.Do(singleton)
	class(self).HelpSetSearchCallbacks(Callable.New(search_callback), Callable.New(action_callback))
}

/*
Registers callables to emit when the menu is respectively about to show or closed. Callback methods should have zero arguments.
*/
func GlobalMenuSetPopupCallbacks(menu_root string, open_callback func(), close_callback func()) { //gd:DisplayServer.global_menu_set_popup_callbacks
	once.Do(singleton)
	class(self).GlobalMenuSetPopupCallbacks(String.New(menu_root), Callable.New(open_callback), Callable.New(close_callback))
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
func GlobalMenuAddSubmenuItem(menu_root string, label string, submenu string) int { //gd:DisplayServer.global_menu_add_submenu_item
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddSubmenuItem(String.New(menu_root), String.New(label), String.New(submenu), int64(-1))))
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
func GlobalMenuAddItem(menu_root string, label string) int { //gd:DisplayServer.global_menu_add_item
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddItem(String.New(menu_root), String.New(label), Callable.New(Callable.Nil), Callable.New(Callable.Nil), variant.New([1]any{}[0]), 0, int64(-1))))
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
func GlobalMenuAddCheckItem(menu_root string, label string) int { //gd:DisplayServer.global_menu_add_check_item
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddCheckItem(String.New(menu_root), String.New(label), Callable.New(Callable.Nil), Callable.New(Callable.Nil), variant.New([1]any{}[0]), 0, int64(-1))))
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
func GlobalMenuAddIconItem(menu_root string, icon [1]gdclass.Texture2D, label string) int { //gd:DisplayServer.global_menu_add_icon_item
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddIconItem(String.New(menu_root), icon, String.New(label), Callable.New(Callable.Nil), Callable.New(Callable.Nil), variant.New([1]any{}[0]), 0, int64(-1))))
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
func GlobalMenuAddIconCheckItem(menu_root string, icon [1]gdclass.Texture2D, label string) int { //gd:DisplayServer.global_menu_add_icon_check_item
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddIconCheckItem(String.New(menu_root), icon, String.New(label), Callable.New(Callable.Nil), Callable.New(Callable.Nil), variant.New([1]any{}[0]), 0, int64(-1))))
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
func GlobalMenuAddRadioCheckItem(menu_root string, label string) int { //gd:DisplayServer.global_menu_add_radio_check_item
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddRadioCheckItem(String.New(menu_root), String.New(label), Callable.New(Callable.Nil), Callable.New(Callable.Nil), variant.New([1]any{}[0]), 0, int64(-1))))
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
func GlobalMenuAddIconRadioCheckItem(menu_root string, icon [1]gdclass.Texture2D, label string) int { //gd:DisplayServer.global_menu_add_icon_radio_check_item
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddIconRadioCheckItem(String.New(menu_root), icon, String.New(label), Callable.New(Callable.Nil), Callable.New(Callable.Nil), variant.New([1]any{}[0]), 0, int64(-1))))
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
func GlobalMenuAddMultistateItem(menu_root string, label string, max_states int, default_state int) int { //gd:DisplayServer.global_menu_add_multistate_item
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddMultistateItem(String.New(menu_root), String.New(label), int64(max_states), int64(default_state), Callable.New(Callable.Nil), Callable.New(Callable.Nil), variant.New([1]any{}[0]), 0, int64(-1))))
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
func GlobalMenuAddSeparator(menu_root string) int { //gd:DisplayServer.global_menu_add_separator
	once.Do(singleton)
	return int(int(class(self).GlobalMenuAddSeparator(String.New(menu_root), int64(-1))))
}

/*
Returns the index of the item with the specified [param text]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemIndexFromText(menu_root string, text string) int { //gd:DisplayServer.global_menu_get_item_index_from_text
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemIndexFromText(String.New(menu_root), String.New(text))))
}

/*
Returns the index of the item with the specified [param tag]. Indices are automatically assigned to each item by the engine, and cannot be set manually.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemIndexFromTag(menu_root string, tag any) int { //gd:DisplayServer.global_menu_get_item_index_from_tag
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemIndexFromTag(String.New(menu_root), variant.New(tag))))
}

/*
Returns [code]true[/code] if the item at index [param idx] is checked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemChecked(menu_root string, idx int) bool { //gd:DisplayServer.global_menu_is_item_checked
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemChecked(String.New(menu_root), int64(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemCheckable(menu_root string, idx int) bool { //gd:DisplayServer.global_menu_is_item_checkable
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemCheckable(String.New(menu_root), int64(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemRadioCheckable(menu_root string, idx int) bool { //gd:DisplayServer.global_menu_is_item_radio_checkable
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemRadioCheckable(String.New(menu_root), int64(idx)))
}

/*
Returns the callback of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemCallback(menu_root string, idx int) Callable.Function { //gd:DisplayServer.global_menu_get_item_callback
	once.Do(singleton)
	return Callable.Function(class(self).GlobalMenuGetItemCallback(String.New(menu_root), int64(idx)))
}

/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemKeyCallback(menu_root string, idx int) Callable.Function { //gd:DisplayServer.global_menu_get_item_key_callback
	once.Do(singleton)
	return Callable.Function(class(self).GlobalMenuGetItemKeyCallback(String.New(menu_root), int64(idx)))
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method global_menu_set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemTag(menu_root string, idx int) any { //gd:DisplayServer.global_menu_get_item_tag
	once.Do(singleton)
	return any(class(self).GlobalMenuGetItemTag(String.New(menu_root), int64(idx)).Interface())
}

/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemText(menu_root string, idx int) string { //gd:DisplayServer.global_menu_get_item_text
	once.Do(singleton)
	return string(class(self).GlobalMenuGetItemText(String.New(menu_root), int64(idx)).String())
}

/*
Returns the submenu ID of the item at index [param idx]. See [method global_menu_add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemSubmenu(menu_root string, idx int) string { //gd:DisplayServer.global_menu_get_item_submenu
	once.Do(singleton)
	return string(class(self).GlobalMenuGetItemSubmenu(String.New(menu_root), int64(idx)).String())
}

/*
Returns the accelerator of the item at index [param idx]. Accelerators are special combinations of keys that activate the item, no matter which control is focused.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemAccelerator(menu_root string, idx int) Key { //gd:DisplayServer.global_menu_get_item_accelerator
	once.Do(singleton)
	return Key(class(self).GlobalMenuGetItemAccelerator(String.New(menu_root), int64(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method global_menu_set_item_disabled] for more info on how to disable an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemDisabled(menu_root string, idx int) bool { //gd:DisplayServer.global_menu_is_item_disabled
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemDisabled(String.New(menu_root), int64(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is hidden.
See [method global_menu_set_item_hidden] for more info on how to hide an item.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuIsItemHidden(menu_root string, idx int) bool { //gd:DisplayServer.global_menu_is_item_hidden
	once.Do(singleton)
	return bool(class(self).GlobalMenuIsItemHidden(String.New(menu_root), int64(idx)))
}

/*
Returns the tooltip associated with the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemTooltip(menu_root string, idx int) string { //gd:DisplayServer.global_menu_get_item_tooltip
	once.Do(singleton)
	return string(class(self).GlobalMenuGetItemTooltip(String.New(menu_root), int64(idx)).String())
}

/*
Returns the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemState(menu_root string, idx int) int { //gd:DisplayServer.global_menu_get_item_state
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemState(String.New(menu_root), int64(idx))))
}

/*
Returns number of states of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemMaxStates(menu_root string, idx int) int { //gd:DisplayServer.global_menu_get_item_max_states
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemMaxStates(String.New(menu_root), int64(idx))))
}

/*
Returns the icon of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemIcon(menu_root string, idx int) [1]gdclass.Texture2D { //gd:DisplayServer.global_menu_get_item_icon
	once.Do(singleton)
	return [1]gdclass.Texture2D(class(self).GlobalMenuGetItemIcon(String.New(menu_root), int64(idx)))
}

/*
Returns the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemIndentationLevel(menu_root string, idx int) int { //gd:DisplayServer.global_menu_get_item_indentation_level
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemIndentationLevel(String.New(menu_root), int64(idx))))
}

/*
Sets the checkstate status of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemChecked(menu_root string, idx int, checked bool) { //gd:DisplayServer.global_menu_set_item_checked
	once.Do(singleton)
	class(self).GlobalMenuSetItemChecked(String.New(menu_root), int64(idx), checked)
}

/*
Sets whether the item at index [param idx] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemCheckable(menu_root string, idx int, checkable bool) { //gd:DisplayServer.global_menu_set_item_checkable
	once.Do(singleton)
	class(self).GlobalMenuSetItemCheckable(String.New(menu_root), int64(idx), checkable)
}

/*
Sets the type of the item at the specified index [param idx] to radio button. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemRadioCheckable(menu_root string, idx int, checkable bool) { //gd:DisplayServer.global_menu_set_item_radio_checkable
	once.Do(singleton)
	class(self).GlobalMenuSetItemRadioCheckable(String.New(menu_root), int64(idx), checkable)
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when an item is pressed.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemCallback(menu_root string, idx int, callback func(tag any)) { //gd:DisplayServer.global_menu_set_item_callback
	once.Do(singleton)
	class(self).GlobalMenuSetItemCallback(String.New(menu_root), int64(idx), Callable.New(callback))
}

/*
Sets the callback of the item at index [param idx]. The callback is emitted when an item is hovered.
[b]Note:[/b] The [param callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemHoverCallbacks(menu_root string, idx int, callback func(tag any)) { //gd:DisplayServer.global_menu_set_item_hover_callbacks
	once.Do(singleton)
	class(self).GlobalMenuSetItemHoverCallbacks(String.New(menu_root), int64(idx), Callable.New(callback))
}

/*
Sets the callback of the item at index [param idx]. Callback is emitted when its accelerator is activated.
[b]Note:[/b] The [param key_callback] Callable needs to accept exactly one Variant parameter, the parameter passed to the Callable will be the value passed to the [code]tag[/code] parameter when the menu item was created.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemKeyCallback(menu_root string, idx int, key_callback func(tag any)) { //gd:DisplayServer.global_menu_set_item_key_callback
	once.Do(singleton)
	class(self).GlobalMenuSetItemKeyCallback(String.New(menu_root), int64(idx), Callable.New(key_callback))
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method global_menu_get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemTag(menu_root string, idx int, tag any) { //gd:DisplayServer.global_menu_set_item_tag
	once.Do(singleton)
	class(self).GlobalMenuSetItemTag(String.New(menu_root), int64(idx), variant.New(tag))
}

/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemText(menu_root string, idx int, text string) { //gd:DisplayServer.global_menu_set_item_text
	once.Do(singleton)
	class(self).GlobalMenuSetItemText(String.New(menu_root), int64(idx), String.New(text))
}

/*
Sets the submenu of the item at index [param idx]. The submenu is the ID of a global menu root that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemSubmenu(menu_root string, idx int, submenu string) { //gd:DisplayServer.global_menu_set_item_submenu
	once.Do(singleton)
	class(self).GlobalMenuSetItemSubmenu(String.New(menu_root), int64(idx), String.New(submenu))
}

/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemAccelerator(menu_root string, idx int, keycode Key) { //gd:DisplayServer.global_menu_set_item_accelerator
	once.Do(singleton)
	class(self).GlobalMenuSetItemAccelerator(String.New(menu_root), int64(idx), keycode)
}

/*
Enables/disables the item at index [param idx]. When it is disabled, it can't be selected and its action can't be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemDisabled(menu_root string, idx int, disabled bool) { //gd:DisplayServer.global_menu_set_item_disabled
	once.Do(singleton)
	class(self).GlobalMenuSetItemDisabled(String.New(menu_root), int64(idx), disabled)
}

/*
Hides/shows the item at index [param idx]. When it is hidden, an item does not appear in a menu and its action cannot be invoked.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemHidden(menu_root string, idx int, hidden bool) { //gd:DisplayServer.global_menu_set_item_hidden
	once.Do(singleton)
	class(self).GlobalMenuSetItemHidden(String.New(menu_root), int64(idx), hidden)
}

/*
Sets the [String] tooltip of the item at the specified index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemTooltip(menu_root string, idx int, tooltip string) { //gd:DisplayServer.global_menu_set_item_tooltip
	once.Do(singleton)
	class(self).GlobalMenuSetItemTooltip(String.New(menu_root), int64(idx), String.New(tooltip))
}

/*
Sets the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemState(menu_root string, idx int, state int) { //gd:DisplayServer.global_menu_set_item_state
	once.Do(singleton)
	class(self).GlobalMenuSetItemState(String.New(menu_root), int64(idx), int64(state))
}

/*
Sets number of state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemMaxStates(menu_root string, idx int, max_states int) { //gd:DisplayServer.global_menu_set_item_max_states
	once.Do(singleton)
	class(self).GlobalMenuSetItemMaxStates(String.New(menu_root), int64(idx), int64(max_states))
}

/*
Replaces the [Texture2D] icon of the specified [param idx].
[b]Note:[/b] This method is implemented only on macOS.
[b]Note:[/b] This method is not supported by macOS "_dock" menu items.
*/
func GlobalMenuSetItemIcon(menu_root string, idx int, icon [1]gdclass.Texture2D) { //gd:DisplayServer.global_menu_set_item_icon
	once.Do(singleton)
	class(self).GlobalMenuSetItemIcon(String.New(menu_root), int64(idx), icon)
}

/*
Sets the horizontal offset of the item at the given [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuSetItemIndentationLevel(menu_root string, idx int, level int) { //gd:DisplayServer.global_menu_set_item_indentation_level
	once.Do(singleton)
	class(self).GlobalMenuSetItemIndentationLevel(String.New(menu_root), int64(idx), int64(level))
}

/*
Returns number of items in the global menu with ID [param menu_root].
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetItemCount(menu_root string) int { //gd:DisplayServer.global_menu_get_item_count
	once.Do(singleton)
	return int(int(class(self).GlobalMenuGetItemCount(String.New(menu_root))))
}

/*
Removes the item at index [param idx] from the global menu [param menu_root].
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuRemoveItem(menu_root string, idx int) { //gd:DisplayServer.global_menu_remove_item
	once.Do(singleton)
	class(self).GlobalMenuRemoveItem(String.New(menu_root), int64(idx))
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
func GlobalMenuClear(menu_root string) { //gd:DisplayServer.global_menu_clear
	once.Do(singleton)
	class(self).GlobalMenuClear(String.New(menu_root))
}

/*
Returns Dictionary of supported system menu IDs and names.
[b]Note:[/b] This method is implemented only on macOS.
*/
func GlobalMenuGetSystemMenuRoots() map[string]string { //gd:DisplayServer.global_menu_get_system_menu_roots
	once.Do(singleton)
	return map[string]string(gd.DictionaryAs[map[string]string](class(self).GlobalMenuGetSystemMenuRoots()))
}

/*
Returns [code]true[/code] if the synthesizer is generating speech, or have utterance waiting in the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsIsSpeaking() bool { //gd:DisplayServer.tts_is_speaking
	once.Do(singleton)
	return bool(class(self).TtsIsSpeaking())
}

/*
Returns [code]true[/code] if the synthesizer is in a paused state.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsIsPaused() bool { //gd:DisplayServer.tts_is_paused
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
func TtsGetVoices() []TextToSpeechVoice { //gd:DisplayServer.tts_get_voices
	once.Do(singleton)
	return []TextToSpeechVoice(gd.ArrayAs[[]TextToSpeechVoice](gd.InternalArray(class(self).TtsGetVoices())))
}

/*
Returns an [PackedStringArray] of voice identifiers for the [param language].
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsGetVoicesForLanguage(language string) []string { //gd:DisplayServer.tts_get_voices_for_language
	once.Do(singleton)
	return []string(class(self).TtsGetVoicesForLanguage(String.New(language)).Strings())
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
func TtsSpeak(text string, voice string) { //gd:DisplayServer.tts_speak
	once.Do(singleton)
	class(self).TtsSpeak(String.New(text), String.New(voice), int64(50), float64(1.0), float64(1.0), int64(0), false)
}

/*
Puts the synthesizer into a paused state.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsPause() { //gd:DisplayServer.tts_pause
	once.Do(singleton)
	class(self).TtsPause()
}

/*
Resumes the synthesizer if it was paused.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsResume() { //gd:DisplayServer.tts_resume
	once.Do(singleton)
	class(self).TtsResume()
}

/*
Stops synthesis in progress and removes all utterances from the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Linux), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
func TtsStop() { //gd:DisplayServer.tts_stop
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
func TtsSetUtteranceCallback(event gdclass.DisplayServerTTSUtteranceEvent, callable func(int, int)) { //gd:DisplayServer.tts_set_utterance_callback
	once.Do(singleton)
	class(self).TtsSetUtteranceCallback(event, Callable.New(callable))
}

/*
Returns [code]true[/code] if OS supports dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
func IsDarkModeSupported() bool { //gd:DisplayServer.is_dark_mode_supported
	once.Do(singleton)
	return bool(class(self).IsDarkModeSupported())
}

/*
Returns [code]true[/code] if OS is using dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
func IsDarkMode() bool { //gd:DisplayServer.is_dark_mode
	once.Do(singleton)
	return bool(class(self).IsDarkMode())
}

/*
Returns OS theme accent color. Returns [code]Color(0, 0, 0, 0)[/code], if accent color is unknown.
[b]Note:[/b] This method is implemented on macOS, Windows, and Android.
*/
func GetAccentColor() Color.RGBA { //gd:DisplayServer.get_accent_color
	once.Do(singleton)
	return Color.RGBA(class(self).GetAccentColor())
}

/*
Returns the OS theme base color (default control background). Returns [code]Color(0, 0, 0, 0)[/code] if the base color is unknown.
[b]Note:[/b] This method is implemented on macOS, Windows, and Android.
*/
func GetBaseColor() Color.RGBA { //gd:DisplayServer.get_base_color
	once.Do(singleton)
	return Color.RGBA(class(self).GetBaseColor())
}

/*
Sets the [param callable] that should be called when system theme settings are changed. Callback method should have zero arguments.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
func SetSystemThemeChangeCallback(callable func()) { //gd:DisplayServer.set_system_theme_change_callback
	once.Do(singleton)
	class(self).SetSystemThemeChangeCallback(Callable.New(callable))
}

/*
Sets the current mouse mode. See also [method mouse_get_mode].
*/
func MouseSetMode(mouse_mode gdclass.DisplayServerMouseMode) { //gd:DisplayServer.mouse_set_mode
	once.Do(singleton)
	class(self).MouseSetMode(mouse_mode)
}

/*
Returns the current mouse mode. See also [method mouse_set_mode].
*/
func MouseGetMode() gdclass.DisplayServerMouseMode { //gd:DisplayServer.mouse_get_mode
	once.Do(singleton)
	return gdclass.DisplayServerMouseMode(class(self).MouseGetMode())
}

/*
Sets the mouse cursor position to the given [param position] relative to an origin at the upper left corner of the currently focused game Window Manager window.
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS, and Linux (X11/Wayland). It has no effect on Android, iOS, and Web.
*/
func WarpMouse(position Vector2i.XY) { //gd:DisplayServer.warp_mouse
	once.Do(singleton)
	class(self).WarpMouse(Vector2i.XY(position))
}

/*
Returns the mouse cursor's current position in screen coordinates.
*/
func MouseGetPosition() Vector2i.XY { //gd:DisplayServer.mouse_get_position
	once.Do(singleton)
	return Vector2i.XY(class(self).MouseGetPosition())
}

/*
Returns the current state of mouse buttons (whether each button is pressed) as a bitmask. If multiple mouse buttons are pressed at the same time, the bits are added together. Equivalent to [method Input.get_mouse_button_mask].
*/
func MouseGetButtonState() MouseButtonMask { //gd:DisplayServer.mouse_get_button_state
	once.Do(singleton)
	return MouseButtonMask(class(self).MouseGetButtonState())
}

/*
Sets the user's clipboard content to the given string.
*/
func ClipboardSet(clipboard string) { //gd:DisplayServer.clipboard_set
	once.Do(singleton)
	class(self).ClipboardSet(String.New(clipboard))
}

/*
Returns the user's clipboard as a string if possible.
*/
func ClipboardGet() string { //gd:DisplayServer.clipboard_get
	once.Do(singleton)
	return string(class(self).ClipboardGet().String())
}

/*
Returns the user's clipboard as an image if possible.
[b]Note:[/b] This method uses the copied pixel data, e.g. from a image editing software or a web browser, not an image file copied from file explorer.
*/
func ClipboardGetImage() [1]gdclass.Image { //gd:DisplayServer.clipboard_get_image
	once.Do(singleton)
	return [1]gdclass.Image(class(self).ClipboardGetImage())
}

/*
Returns [code]true[/code] if there is a text content on the user's clipboard.
*/
func ClipboardHas() bool { //gd:DisplayServer.clipboard_has
	once.Do(singleton)
	return bool(class(self).ClipboardHas())
}

/*
Returns [code]true[/code] if there is an image content on the user's clipboard.
*/
func ClipboardHasImage() bool { //gd:DisplayServer.clipboard_has_image
	once.Do(singleton)
	return bool(class(self).ClipboardHasImage())
}

/*
Sets the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard content to the given string. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
func ClipboardSetPrimary(clipboard_primary string) { //gd:DisplayServer.clipboard_set_primary
	once.Do(singleton)
	class(self).ClipboardSetPrimary(String.New(clipboard_primary))
}

/*
Returns the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard as a string if possible. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
func ClipboardGetPrimary() string { //gd:DisplayServer.clipboard_get_primary
	once.Do(singleton)
	return string(class(self).ClipboardGetPrimary().String())
}

/*
Returns an [Array] of [Rect2], each of which is the bounding rectangle for a display cutout or notch. These are non-functional areas on edge-to-edge screens used by cameras and sensors. Returns an empty array if the device does not have cutouts. See also [method get_display_safe_area].
[b]Note:[/b] Currently only implemented on Android. Other platforms will return an empty array even if they do have display cutouts or notches.
*/
func GetDisplayCutouts() []Rect2.PositionSize { //gd:DisplayServer.get_display_cutouts
	once.Do(singleton)
	return []Rect2.PositionSize(gd.ArrayAs[[]Rect2.PositionSize](gd.InternalArray(class(self).GetDisplayCutouts())))
}

/*
Returns the unobscured area of the display where interactive controls should be rendered. See also [method get_display_cutouts].
[b]Note:[/b] Currently only implemented on Android and iOS. On other platforms, [code]screen_get_usable_rect(SCREEN_OF_MAIN_WINDOW)[/code] will be returned as a fallback. See also [method screen_get_usable_rect].
*/
func GetDisplaySafeArea() Rect2i.PositionSize { //gd:DisplayServer.get_display_safe_area
	once.Do(singleton)
	return Rect2i.PositionSize(class(self).GetDisplaySafeArea())
}

/*
Returns the number of displays available.
*/
func GetScreenCount() int { //gd:DisplayServer.get_screen_count
	once.Do(singleton)
	return int(int(class(self).GetScreenCount()))
}

/*
Returns index of the primary screen.
*/
func GetPrimaryScreen() int { //gd:DisplayServer.get_primary_screen
	once.Do(singleton)
	return int(int(class(self).GetPrimaryScreen()))
}

/*
Returns the index of the screen containing the window with the keyboard focus, or the primary screen if there's no focused window.
*/
func GetKeyboardFocusScreen() int { //gd:DisplayServer.get_keyboard_focus_screen
	once.Do(singleton)
	return int(int(class(self).GetKeyboardFocusScreen()))
}

/*
Returns the index of the screen that overlaps the most with the given rectangle. Returns [code]-1[/code] if the rectangle doesn't overlap with any screen or has no area.
*/
func GetScreenFromRect(rect Rect2.PositionSize) int { //gd:DisplayServer.get_screen_from_rect
	once.Do(singleton)
	return int(int(class(self).GetScreenFromRect(Rect2.PositionSize(rect))))
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
func ScreenGetPosition() Vector2i.XY { //gd:DisplayServer.screen_get_position
	once.Do(singleton)
	return Vector2i.XY(class(self).ScreenGetPosition(int64(-1)))
}

/*
Returns the screen's size in pixels. See also [method screen_get_position] and [method screen_get_usable_rect].
*/
func ScreenGetSize() Vector2i.XY { //gd:DisplayServer.screen_get_size
	once.Do(singleton)
	return Vector2i.XY(class(self).ScreenGetSize(int64(-1)))
}

/*
Returns the portion of the screen that is not obstructed by a status bar in pixels. See also [method screen_get_size].
*/
func ScreenGetUsableRect() Rect2i.PositionSize { //gd:DisplayServer.screen_get_usable_rect
	once.Do(singleton)
	return Rect2i.PositionSize(class(self).ScreenGetUsableRect(int64(-1)))
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
func ScreenGetDpi() int { //gd:DisplayServer.screen_get_dpi
	once.Do(singleton)
	return int(int(class(self).ScreenGetDpi(int64(-1))))
}

/*
Returns the scale factor of the specified screen by index.
[b]Note:[/b] On macOS, the returned value is [code]2.0[/code] for hiDPI (Retina) screens, and [code]1.0[/code] for all other cases.
[b]Note:[/b] On Linux (Wayland), the returned value is accurate only when [param screen] is [constant SCREEN_OF_MAIN_WINDOW]. Due to API limitations, passing a direct index will return a rounded-up integer, if the screen has a fractional scale (e.g. [code]1.25[/code] would get rounded up to [code]2.0[/code]).
[b]Note:[/b] This method is implemented on Android, iOS, Web, macOS, and Linux (Wayland).
*/
func ScreenGetScale() Float.X { //gd:DisplayServer.screen_get_scale
	once.Do(singleton)
	return Float.X(Float.X(class(self).ScreenGetScale(int64(-1))))
}

/*
Returns [code]true[/code] if touch events are available (Android or iOS), the capability is detected on the Web platform or if [member ProjectSettings.input_devices/pointing/emulate_touch_from_mouse] is [code]true[/code].
*/
func IsTouchscreenAvailable() bool { //gd:DisplayServer.is_touchscreen_available
	once.Do(singleton)
	return bool(class(self).IsTouchscreenAvailable())
}

/*
Returns the greatest scale factor of all screens.
[b]Note:[/b] On macOS returned value is [code]2.0[/code] if there is at least one hiDPI (Retina) screen in the system, and [code]1.0[/code] in all other cases.
[b]Note:[/b] This method is implemented only on macOS.
*/
func ScreenGetMaxScale() Float.X { //gd:DisplayServer.screen_get_max_scale
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
func ScreenGetRefreshRate() Float.X { //gd:DisplayServer.screen_get_refresh_rate
	once.Do(singleton)
	return Float.X(Float.X(class(self).ScreenGetRefreshRate(int64(-1))))
}

/*
Returns color of the display pixel at the [param position].
[b]Note:[/b] This method is implemented on Linux (X11), macOS, and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
func ScreenGetPixel(position Vector2i.XY) Color.RGBA { //gd:DisplayServer.screen_get_pixel
	once.Do(singleton)
	return Color.RGBA(class(self).ScreenGetPixel(Vector2i.XY(position)))
}

/*
Returns screenshot of the [param screen].
[b]Note:[/b] This method is implemented on Linux (X11), macOS, and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
func ScreenGetImage() [1]gdclass.Image { //gd:DisplayServer.screen_get_image
	once.Do(singleton)
	return [1]gdclass.Image(class(self).ScreenGetImage(int64(-1)))
}

/*
Returns screenshot of the screen [param rect].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
func ScreenGetImageRect(rect Rect2i.PositionSize) [1]gdclass.Image { //gd:DisplayServer.screen_get_image_rect
	once.Do(singleton)
	return [1]gdclass.Image(class(self).ScreenGetImageRect(Rect2i.PositionSize(rect)))
}

/*
Sets the [param screen]'s [param orientation]. See also [method screen_get_orientation].
[b]Note:[/b] On iOS, this method has no effect if [member ProjectSettings.display/window/handheld/orientation] is not set to [constant SCREEN_SENSOR].
*/
func ScreenSetOrientation(orientation gdclass.DisplayServerScreenOrientation) { //gd:DisplayServer.screen_set_orientation
	once.Do(singleton)
	class(self).ScreenSetOrientation(orientation, int64(-1))
}

/*
Returns the [param screen]'s current orientation. See also [method screen_set_orientation].
[b]Note:[/b] This method is implemented on Android and iOS.
*/
func ScreenGetOrientation() gdclass.DisplayServerScreenOrientation { //gd:DisplayServer.screen_get_orientation
	once.Do(singleton)
	return gdclass.DisplayServerScreenOrientation(class(self).ScreenGetOrientation(int64(-1)))
}

/*
Sets whether the screen should never be turned off by the operating system's power-saving measures. See also [method screen_is_kept_on].
*/
func ScreenSetKeepOn(enable bool) { //gd:DisplayServer.screen_set_keep_on
	once.Do(singleton)
	class(self).ScreenSetKeepOn(enable)
}

/*
Returns [code]true[/code] if the screen should never be turned off by the operating system's power-saving measures. See also [method screen_set_keep_on].
*/
func ScreenIsKeptOn() bool { //gd:DisplayServer.screen_is_kept_on
	once.Do(singleton)
	return bool(class(self).ScreenIsKeptOn())
}

/*
Returns the list of Godot window IDs belonging to this process.
[b]Note:[/b] Native dialogs are not included in this list.
*/
func GetWindowList() []int32 { //gd:DisplayServer.get_window_list
	once.Do(singleton)
	return []int32(slices.Collect(class(self).GetWindowList().Values()))
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
func GetWindowAtScreenPosition(position Vector2i.XY) int { //gd:DisplayServer.get_window_at_screen_position
	once.Do(singleton)
	return int(int(class(self).GetWindowAtScreenPosition(Vector2i.XY(position))))
}

/*
Returns internal structure pointers for use in plugins.
[b]Note:[/b] This method is implemented on Android, Linux (X11/Wayland), macOS, and Windows.
*/
func WindowGetNativeHandle(handle_type gdclass.DisplayServerHandleType) int { //gd:DisplayServer.window_get_native_handle
	once.Do(singleton)
	return int(int(class(self).WindowGetNativeHandle(handle_type, int64(0))))
}

/*
Returns ID of the active popup window, or [constant INVALID_WINDOW_ID] if there is none.
*/
func WindowGetActivePopup() int { //gd:DisplayServer.window_get_active_popup
	once.Do(singleton)
	return int(int(class(self).WindowGetActivePopup()))
}

/*
Sets the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system. Clicking this area will not auto-close this popup.
*/
func WindowSetPopupSafeRect(window int, rect Rect2i.PositionSize) { //gd:DisplayServer.window_set_popup_safe_rect
	once.Do(singleton)
	class(self).WindowSetPopupSafeRect(int64(window), Rect2i.PositionSize(rect))
}

/*
Returns the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system.
*/
func WindowGetPopupSafeRect(window int) Rect2i.PositionSize { //gd:DisplayServer.window_get_popup_safe_rect
	once.Do(singleton)
	return Rect2i.PositionSize(class(self).WindowGetPopupSafeRect(int64(window)))
}

/*
Sets the title of the given window to [param title].
[b]Note:[/b] It's recommended to change this value using [member Window.title] instead.
[b]Note:[/b] Avoid changing the window title every frame, as this can cause performance issues on certain window managers. Try to change the window title only a few times per second at most.
*/
func WindowSetTitle(title string) { //gd:DisplayServer.window_set_title
	once.Do(singleton)
	class(self).WindowSetTitle(String.New(title), int64(0))
}

/*
Returns the estimated window title bar size (including text and window buttons) for the window specified by [param window_id] (in pixels). This method does not change the window title.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func WindowGetTitleSize(title string) Vector2i.XY { //gd:DisplayServer.window_get_title_size
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetTitleSize(String.New(title), int64(0)))
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
DisplayServer.WindowSetMousePassthrough([]);
[/csharp]
[/codeblocks]
[b]Note:[/b] On Windows, the portion of a window that lies outside the region is not drawn, while on Linux (X11) and macOS it is.
[b]Note:[/b] This method is implemented on Linux (X11), macOS and Windows.
*/
func WindowSetMousePassthrough(region []Vector2.XY) { //gd:DisplayServer.window_set_mouse_passthrough
	once.Do(singleton)
	class(self).WindowSetMousePassthrough(Packed.New(region...), int64(0))
}

/*
Returns the screen the window specified by [param window_id] is currently positioned on. If the screen overlaps multiple displays, the screen where the window's center is located is returned. See also [method window_set_current_screen].
*/
func WindowGetCurrentScreen() int { //gd:DisplayServer.window_get_current_screen
	once.Do(singleton)
	return int(int(class(self).WindowGetCurrentScreen(int64(0))))
}

/*
Moves the window specified by [param window_id] to the specified [param screen]. See also [method window_get_current_screen].
*/
func WindowSetCurrentScreen(screen int) { //gd:DisplayServer.window_set_current_screen
	once.Do(singleton)
	class(self).WindowSetCurrentScreen(int64(screen), int64(0))
}

/*
Returns the position of the client area of the given window on the screen.
*/
func WindowGetPosition() Vector2i.XY { //gd:DisplayServer.window_get_position
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetPosition(int64(0)))
}

/*
Returns the position of the given window on the screen including the borders drawn by the operating system. See also [method window_get_position].
*/
func WindowGetPositionWithDecorations() Vector2i.XY { //gd:DisplayServer.window_get_position_with_decorations
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetPositionWithDecorations(int64(0)))
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
func WindowSetPosition(position Vector2i.XY) { //gd:DisplayServer.window_set_position
	once.Do(singleton)
	class(self).WindowSetPosition(Vector2i.XY(position), int64(0))
}

/*
Returns the size of the window specified by [param window_id] (in pixels), excluding the borders drawn by the operating system. This is also called the "client area". See also [method window_get_size_with_decorations], [method window_set_size] and [method window_get_position].
*/
func WindowGetSize() Vector2i.XY { //gd:DisplayServer.window_get_size
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetSize(int64(0)))
}

/*
Sets the size of the given window to [param size] (in pixels). See also [method window_get_size] and [method window_get_position].
[b]Note:[/b] It's recommended to change this value using [member Window.size] instead.
*/
func WindowSetSize(size Vector2i.XY) { //gd:DisplayServer.window_set_size
	once.Do(singleton)
	class(self).WindowSetSize(Vector2i.XY(size), int64(0))
}

/*
Sets the [param callback] that will be called when the window specified by [param window_id] is moved or resized.
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
func WindowSetRectChangedCallback(callback func(rect Rect2i.PositionSize)) { //gd:DisplayServer.window_set_rect_changed_callback
	once.Do(singleton)
	class(self).WindowSetRectChangedCallback(Callable.New(callback), int64(0))
}

/*
Sets the [param callback] that will be called when an event occurs in the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
func WindowSetWindowEventCallback(callback func(event gdclass.DisplayServerWindowEvent)) { //gd:DisplayServer.window_set_window_event_callback
	once.Do(singleton)
	class(self).WindowSetWindowEventCallback(Callable.New(callback), int64(0))
}

/*
Sets the [param callback] that should be called when any [InputEvent] is sent to the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
func WindowSetInputEventCallback(callback func(event [1]gdclass.InputEvent)) { //gd:DisplayServer.window_set_input_event_callback
	once.Do(singleton)
	class(self).WindowSetInputEventCallback(Callable.New(callback), int64(0))
}

/*
Sets the [param callback] that should be called when text is entered using the virtual keyboard to the window specified by [param window_id].
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
*/
func WindowSetInputTextCallback(callback func(text string)) { //gd:DisplayServer.window_set_input_text_callback
	once.Do(singleton)
	class(self).WindowSetInputTextCallback(Callable.New(callback), int64(0))
}

/*
Sets the [param callback] that should be called when files are dropped from the operating system's file manager to the window specified by [param window_id]. [param callback] should take one [PackedStringArray] argument, which is the list of dropped files.
[b]Warning:[/b] Advanced users only! Adding such a callback to a [Window] node will override its default implementation, which can introduce bugs.
[b]Note:[/b] This method is implemented on Windows, macOS, Linux (X11/Wayland), and Web.
*/
func WindowSetDropFilesCallback(callback func(tag any)) { //gd:DisplayServer.window_set_drop_files_callback
	once.Do(singleton)
	class(self).WindowSetDropFilesCallback(Callable.New(callback), int64(0))
}

/*
Returns the [method Object.get_instance_id] of the [Window] the [param window_id] is attached to.
*/
func WindowGetAttachedInstanceId() int { //gd:DisplayServer.window_get_attached_instance_id
	once.Do(singleton)
	return int(int(class(self).WindowGetAttachedInstanceId(int64(0))))
}

/*
Returns the window's maximum size (in pixels). See also [method window_set_max_size].
*/
func WindowGetMaxSize() Vector2i.XY { //gd:DisplayServer.window_get_max_size
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetMaxSize(int64(0)))
}

/*
Sets the maximum size of the window specified by [param window_id] in pixels. Normally, the user will not be able to drag the window to make it larger than the specified size. See also [method window_get_max_size].
[b]Note:[/b] It's recommended to change this value using [member Window.max_size] instead.
[b]Note:[/b] Using third-party tools, it is possible for users to disable window geometry restrictions and therefore bypass this limit.
*/
func WindowSetMaxSize(max_size Vector2i.XY) { //gd:DisplayServer.window_set_max_size
	once.Do(singleton)
	class(self).WindowSetMaxSize(Vector2i.XY(max_size), int64(0))
}

/*
Returns the window's minimum size (in pixels). See also [method window_set_min_size].
*/
func WindowGetMinSize() Vector2i.XY { //gd:DisplayServer.window_get_min_size
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetMinSize(int64(0)))
}

/*
Sets the minimum size for the given window to [param min_size] in pixels. Normally, the user will not be able to drag the window to make it smaller than the specified size. See also [method window_get_min_size].
[b]Note:[/b] It's recommended to change this value using [member Window.min_size] instead.
[b]Note:[/b] By default, the main window has a minimum size of [code]Vector2i(64, 64)[/code]. This prevents issues that can arise when the window is resized to a near-zero size.
[b]Note:[/b] Using third-party tools, it is possible for users to disable window geometry restrictions and therefore bypass this limit.
*/
func WindowSetMinSize(min_size Vector2i.XY) { //gd:DisplayServer.window_set_min_size
	once.Do(singleton)
	class(self).WindowSetMinSize(Vector2i.XY(min_size), int64(0))
}

/*
Returns the size of the window specified by [param window_id] (in pixels), including the borders drawn by the operating system. See also [method window_get_size].
*/
func WindowGetSizeWithDecorations() Vector2i.XY { //gd:DisplayServer.window_get_size_with_decorations
	once.Do(singleton)
	return Vector2i.XY(class(self).WindowGetSizeWithDecorations(int64(0)))
}

/*
Returns the mode of the given window.
*/
func WindowGetMode() gdclass.DisplayServerWindowMode { //gd:DisplayServer.window_get_mode
	once.Do(singleton)
	return gdclass.DisplayServerWindowMode(class(self).WindowGetMode(int64(0)))
}

/*
Sets window mode for the given window to [param mode]. See [enum WindowMode] for possible values and how each mode behaves.
[b]Note:[/b] On Android, setting it to [constant WINDOW_MODE_FULLSCREEN] or [constant WINDOW_MODE_EXCLUSIVE_FULLSCREEN] will enable immersive mode.
[b]Note:[/b] Setting the window to full screen forcibly sets the borderless flag to [code]true[/code], so make sure to set it back to [code]false[/code] when not wanted.
*/
func WindowSetMode(mode gdclass.DisplayServerWindowMode) { //gd:DisplayServer.window_set_mode
	once.Do(singleton)
	class(self).WindowSetMode(mode, int64(0))
}

/*
Enables or disables the given window's given [param flag]. See [enum WindowFlags] for possible values and their behavior.
*/
func WindowSetFlag(flag gdclass.DisplayServerWindowFlags, enabled bool) { //gd:DisplayServer.window_set_flag
	once.Do(singleton)
	class(self).WindowSetFlag(flag, enabled, int64(0))
}

/*
Returns the current value of the given window's [param flag].
*/
func WindowGetFlag(flag gdclass.DisplayServerWindowFlags) bool { //gd:DisplayServer.window_get_flag
	once.Do(singleton)
	return bool(class(self).WindowGetFlag(flag, int64(0)))
}

/*
When [constant WINDOW_FLAG_EXTEND_TO_TITLE] flag is set, set offset to the center of the first titlebar button.
[b]Note:[/b] This flag is implemented only on macOS.
*/
func WindowSetWindowButtonsOffset(offset Vector2i.XY) { //gd:DisplayServer.window_set_window_buttons_offset
	once.Do(singleton)
	class(self).WindowSetWindowButtonsOffset(Vector2i.XY(offset), int64(0))
}

/*
Returns left margins ([code]x[/code]), right margins ([code]y[/code]) and height ([code]z[/code]) of the title that are safe to use (contains no buttons or other elements) when [constant WINDOW_FLAG_EXTEND_TO_TITLE] flag is set.
*/
func WindowGetSafeTitleMargins() Vector3i.XYZ { //gd:DisplayServer.window_get_safe_title_margins
	once.Do(singleton)
	return Vector3i.XYZ(class(self).WindowGetSafeTitleMargins(int64(0)))
}

/*
Makes the window specified by [param window_id] request attention, which is materialized by the window title and taskbar entry blinking until the window is focused. This usually has no visible effect if the window is currently focused. The exact behavior varies depending on the operating system.
*/
func WindowRequestAttention() { //gd:DisplayServer.window_request_attention
	once.Do(singleton)
	class(self).WindowRequestAttention(int64(0))
}

/*
Moves the window specified by [param window_id] to the foreground, so that it is visible over other windows.
*/
func WindowMoveToForeground() { //gd:DisplayServer.window_move_to_foreground
	once.Do(singleton)
	class(self).WindowMoveToForeground(int64(0))
}

/*
Returns [code]true[/code] if the window specified by [param window_id] is focused.
*/
func WindowIsFocused() bool { //gd:DisplayServer.window_is_focused
	once.Do(singleton)
	return bool(class(self).WindowIsFocused(int64(0)))
}

/*
Returns [code]true[/code] if anything can be drawn in the window specified by [param window_id], [code]false[/code] otherwise. Using the [code]--disable-render-loop[/code] command line argument or a headless build will return [code]false[/code].
*/
func WindowCanDraw() bool { //gd:DisplayServer.window_can_draw
	once.Do(singleton)
	return bool(class(self).WindowCanDraw(int64(0)))
}

/*
Sets window transient parent. Transient window will be destroyed with its transient parent and will return focus to their parent when closed. The transient window is displayed on top of a non-exclusive full-screen parent window. Transient windows can't enter full-screen mode.
[b]Note:[/b] It's recommended to change this value using [member Window.transient] instead.
[b]Note:[/b] The behavior might be different depending on the platform.
*/
func WindowSetTransient(window_id int, parent_window_id int) { //gd:DisplayServer.window_set_transient
	once.Do(singleton)
	class(self).WindowSetTransient(int64(window_id), int64(parent_window_id))
}

/*
If set to [code]true[/code], this window will always stay on top of its parent window, parent window will ignore input while this window is opened.
[b]Note:[/b] On macOS, exclusive windows are confined to the same space (virtual desktop or screen) as the parent window.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func WindowSetExclusive(window_id int, exclusive bool) { //gd:DisplayServer.window_set_exclusive
	once.Do(singleton)
	class(self).WindowSetExclusive(int64(window_id), exclusive)
}

/*
Sets whether [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] should be enabled for the window specified by [param window_id]. See also [method window_set_ime_position].
*/
func WindowSetImeActive(active bool) { //gd:DisplayServer.window_set_ime_active
	once.Do(singleton)
	class(self).WindowSetImeActive(active, int64(0))
}

/*
Sets the position of the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] popup for the specified [param window_id]. Only effective if [method window_set_ime_active] was set to [code]true[/code] for the specified [param window_id].
*/
func WindowSetImePosition(position Vector2i.XY) { //gd:DisplayServer.window_set_ime_position
	once.Do(singleton)
	class(self).WindowSetImePosition(Vector2i.XY(position), int64(0))
}

/*
Sets the V-Sync mode of the given window. See also [member ProjectSettings.display/window/vsync/vsync_mode].
See [enum DisplayServer.VSyncMode] for possible values and how they affect the behavior of your application.
Depending on the platform and used renderer, the engine will fall back to [constant VSYNC_ENABLED] if the desired mode is not supported.
[b]Note:[/b] V-Sync modes other than [constant VSYNC_ENABLED] are only supported in the Forward+ and Mobile rendering methods, not Compatibility.
*/
func WindowSetVsyncMode(vsync_mode gdclass.DisplayServerVSyncMode) { //gd:DisplayServer.window_set_vsync_mode
	once.Do(singleton)
	class(self).WindowSetVsyncMode(vsync_mode, int64(0))
}

/*
Returns the V-Sync mode of the given window.
*/
func WindowGetVsyncMode() gdclass.DisplayServerVSyncMode { //gd:DisplayServer.window_get_vsync_mode
	once.Do(singleton)
	return gdclass.DisplayServerVSyncMode(class(self).WindowGetVsyncMode(int64(0)))
}

/*
Returns [code]true[/code] if the given window can be maximized (the maximize button is enabled).
*/
func WindowIsMaximizeAllowed() bool { //gd:DisplayServer.window_is_maximize_allowed
	once.Do(singleton)
	return bool(class(self).WindowIsMaximizeAllowed(int64(0)))
}

/*
Returns [code]true[/code], if double-click on a window title should maximize it.
[b]Note:[/b] This method is implemented only on macOS.
*/
func WindowMaximizeOnTitleDblClick() bool { //gd:DisplayServer.window_maximize_on_title_dbl_click
	once.Do(singleton)
	return bool(class(self).WindowMaximizeOnTitleDblClick())
}

/*
Returns [code]true[/code], if double-click on a window title should minimize it.
[b]Note:[/b] This method is implemented only on macOS.
*/
func WindowMinimizeOnTitleDblClick() bool { //gd:DisplayServer.window_minimize_on_title_dbl_click
	once.Do(singleton)
	return bool(class(self).WindowMinimizeOnTitleDblClick())
}

/*
Starts an interactive drag operation on the window with the given [param window_id], using the current mouse position. Call this method when handling a mouse button being pressed to simulate a pressed event on the window's title bar. Using this method allows the window to participate in space switching, tiling, and other system features.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS, and Windows.
*/
func WindowStartDrag() { //gd:DisplayServer.window_start_drag
	once.Do(singleton)
	class(self).WindowStartDrag(int64(0))
}

/*
Starts an interactive resize operation on the window with the given [param window_id], using the current mouse position. Call this method when handling a mouse button being pressed to simulate a pressed event on the window's edge.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS, and Windows.
*/
func WindowStartResize(edge gdclass.DisplayServerWindowResizeEdge) { //gd:DisplayServer.window_start_resize
	once.Do(singleton)
	class(self).WindowStartResize(edge, int64(0))
}

/*
Returns the text selection in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] composition string, with the [Vector2i]'s [code]x[/code] component being the caret position and [code]y[/code] being the length of the selection.
[b]Note:[/b] This method is implemented only on macOS.
*/
func ImeGetSelection() Vector2i.XY { //gd:DisplayServer.ime_get_selection
	once.Do(singleton)
	return Vector2i.XY(class(self).ImeGetSelection())
}

/*
Returns the composition string contained within the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] window.
[b]Note:[/b] This method is implemented only on macOS.
*/
func ImeGetText() string { //gd:DisplayServer.ime_get_text
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
func VirtualKeyboardShow(existing_text string) { //gd:DisplayServer.virtual_keyboard_show
	once.Do(singleton)
	class(self).VirtualKeyboardShow(String.New(existing_text), Rect2.PositionSize(gd.NewRect2(0, 0, 0, 0)), 0, int64(-1), int64(-1), int64(-1))
}

/*
Hides the virtual keyboard if it is shown, does nothing otherwise.
*/
func VirtualKeyboardHide() { //gd:DisplayServer.virtual_keyboard_hide
	once.Do(singleton)
	class(self).VirtualKeyboardHide()
}

/*
Returns the on-screen keyboard's height in pixels. Returns 0 if there is no keyboard or if it is currently hidden.
*/
func VirtualKeyboardGetHeight() int { //gd:DisplayServer.virtual_keyboard_get_height
	once.Do(singleton)
	return int(int(class(self).VirtualKeyboardGetHeight()))
}

/*
Returns [code]true[/code] if hardware keyboard is connected.
[b]Note:[/b] This method is implemented on Android and iOS, on other platforms this method always returns [code]true[/code].
*/
func HasHardwareKeyboard() bool { //gd:DisplayServer.has_hardware_keyboard
	once.Do(singleton)
	return bool(class(self).HasHardwareKeyboard())
}

/*
Sets the default mouse cursor shape. The cursor's appearance will vary depending on the user's operating system and mouse cursor theme. See also [method cursor_get_shape] and [method cursor_set_custom_image].
*/
func CursorSetShape(shape gdclass.DisplayServerCursorShape) { //gd:DisplayServer.cursor_set_shape
	once.Do(singleton)
	class(self).CursorSetShape(shape)
}

/*
Returns the default mouse cursor shape set by [method cursor_set_shape].
*/
func CursorGetShape() gdclass.DisplayServerCursorShape { //gd:DisplayServer.cursor_get_shape
	once.Do(singleton)
	return gdclass.DisplayServerCursorShape(class(self).CursorGetShape())
}

/*
Sets a custom mouse cursor image for the given [param shape]. This means the user's operating system and mouse cursor theme will no longer influence the mouse cursor's appearance.
[param cursor] can be either a [Texture2D] or an [Image], and it should not be larger than 256×256 to display correctly. Optionally, [param hotspot] can be set to offset the image's position relative to the click point. By default, [param hotspot] is set to the top-left corner of the image. See also [method cursor_set_shape].
*/
func CursorSetCustomImage(cursor [1]gdclass.Resource) { //gd:DisplayServer.cursor_set_custom_image
	once.Do(singleton)
	class(self).CursorSetCustomImage(cursor, 0, Vector2.XY(gd.Vector2{0, 0}))
}

/*
Returns [code]true[/code] if positions of [b]OK[/b] and [b]Cancel[/b] buttons are swapped in dialogs. This is enabled by default on Windows to follow interface conventions, and be toggled by changing [member ProjectSettings.gui/common/swap_cancel_ok].
[b]Note:[/b] This doesn't affect native dialogs such as the ones spawned by [method DisplayServer.dialog_show].
*/
func GetSwapCancelOk() bool { //gd:DisplayServer.get_swap_cancel_ok
	once.Do(singleton)
	return bool(class(self).GetSwapCancelOk())
}

/*
Allows the [param process_id] PID to steal focus from this window. In other words, this disables the operating system's focus stealing protection for the specified PID.
[b]Note:[/b] This method is implemented only on Windows.
*/
func EnableForStealingFocus(process_id int) { //gd:DisplayServer.enable_for_stealing_focus
	once.Do(singleton)
	class(self).EnableForStealingFocus(int64(process_id))
}

/*
Shows a text dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [int] parameter which corresponds to the index of the pressed button.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG] feature. Supported platforms include macOS, Windows, and Android.
*/
func DialogShow(title string, description string, buttons []string, callback func(button int)) error { //gd:DisplayServer.dialog_show
	once.Do(singleton)
	return error(gd.ToError(class(self).DialogShow(String.New(title), String.New(description), Packed.MakeStrings(buttons...), Callable.New(callback))))
}

/*
Shows a text input dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [String] parameter which contains the text field's contents.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_INPUT] feature. Supported platforms include macOS, Windows, and Android.
*/
func DialogInputText(title string, description string, existing_text string, callback func(text string)) error { //gd:DisplayServer.dialog_input_text
	once.Do(singleton)
	return error(gd.ToError(class(self).DialogInputText(String.New(title), String.New(description), String.New(existing_text), Callable.New(callback))))
}

/*
Displays OS native dialog for selecting files or directories in the file system.
Each filter string in the [param filters] array should be formatted like this: [code]*.png,*.jpg,*.jpeg;Image Files;image/png,image/jpeg[/code]. The description text of the filter is optional and can be omitted. It is recommended to set both file extension and MIME type. See also [member FileDialog.filters].
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int[/code]. [b]On Android,[/b] callback argument [code]selected_filter_index[/code] is always zero.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE] feature. Supported platforms include Linux (X11/Wayland), Windows, macOS, and Android.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] Embedded file dialog and Windows file dialog support only file extensions, while Android, Linux, and macOS file dialogs also support MIME types.
[b]Note:[/b] On Android and Linux, [param show_hidden] is ignored.
[b]Note:[/b] On Android and macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
func FileDialogShow(title string, current_directory string, filename string, show_hidden bool, mode gdclass.DisplayServerFileDialogMode, filters []string, callback func(status bool, selected_paths []string, selected_filter_index int)) error { //gd:DisplayServer.file_dialog_show
	once.Do(singleton)
	return error(gd.ToError(class(self).FileDialogShow(String.New(title), String.New(current_directory), String.New(filename), show_hidden, mode, Packed.MakeStrings(filters...), Callable.New(callback))))
}

/*
Displays OS native dialog for selecting files or directories in the file system with additional user selectable options.
Each filter string in the [param filters] array should be formatted like this: [code]*.png,*.jpg,*.jpeg;Image Files;image/png,image/jpeg[/code]. The description text of the filter is optional and can be omitted. It is recommended to set both file extension and MIME type. See also [member FileDialog.filters].
[param options] is array of [Dictionary]s with the following keys:
- [code]"name"[/code] - option's name [String].
- [code]"values"[/code] - [PackedStringArray] of values. If empty, boolean option (check box) is used.
- [code]"default"[/code] - default selected option index ([int]) or default boolean value ([bool]).
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int, selected_option: Dictionary[/code].
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE_EXTRA] feature. Supported platforms include Linux (X11/Wayland), Windows, and macOS.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] Embedded file dialog and Windows file dialog support only file extensions, while Android, Linux, and macOS file dialogs also support MIME types.
[b]Note:[/b] On Linux (X11), [param show_hidden] is ignored.
[b]Note:[/b] On macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
func FileDialogWithOptionsShow(title string, current_directory string, root string, filename string, show_hidden bool, mode gdclass.DisplayServerFileDialogMode, filters []string, options []FileDialogOption, callback func(status bool, selected_paths []string, selected_filter_index int, selected_option map[any]any)) error { //gd:DisplayServer.file_dialog_with_options_show
	once.Do(singleton)
	return error(gd.ToError(class(self).FileDialogWithOptionsShow(String.New(title), String.New(current_directory), String.New(root), String.New(filename), show_hidden, mode, Packed.MakeStrings(filters...), gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](options), Callable.New(callback))))
}

/*
Plays the beep sound from the operative system, if possible. Because it comes from the OS, the beep sound will be audible even if the application is muted. It may also be disabled for the entire OS by the user.
[b]Note:[/b] This method is implemented on macOS, Linux (X11/Wayland), and Windows.
*/
func Beep() { //gd:DisplayServer.beep
	once.Do(singleton)
	class(self).Beep()
}

/*
Returns the number of keyboard layouts.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetLayoutCount() int { //gd:DisplayServer.keyboard_get_layout_count
	once.Do(singleton)
	return int(int(class(self).KeyboardGetLayoutCount()))
}

/*
Returns active keyboard layout index.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS, and Windows.
*/
func KeyboardGetCurrentLayout() int { //gd:DisplayServer.keyboard_get_current_layout
	once.Do(singleton)
	return int(int(class(self).KeyboardGetCurrentLayout()))
}

/*
Sets the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardSetCurrentLayout(index int) { //gd:DisplayServer.keyboard_set_current_layout
	once.Do(singleton)
	class(self).KeyboardSetCurrentLayout(int64(index))
}

/*
Returns the ISO-639/BCP-47 language code of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetLayoutLanguage(index int) string { //gd:DisplayServer.keyboard_get_layout_language
	once.Do(singleton)
	return string(class(self).KeyboardGetLayoutLanguage(int64(index)).String())
}

/*
Returns the localized name of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetLayoutName(index int) string { //gd:DisplayServer.keyboard_get_layout_name
	once.Do(singleton)
	return string(class(self).KeyboardGetLayoutName(int64(index)).String())
}

/*
Converts a physical (US QWERTY) [param keycode] to one in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetKeycodeFromPhysical(keycode Key) Key { //gd:DisplayServer.keyboard_get_keycode_from_physical
	once.Do(singleton)
	return Key(class(self).KeyboardGetKeycodeFromPhysical(keycode))
}

/*
Converts a physical (US QWERTY) [param keycode] to localized label printed on the key in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
func KeyboardGetLabelFromPhysical(keycode Key) Key { //gd:DisplayServer.keyboard_get_label_from_physical
	once.Do(singleton)
	return Key(class(self).KeyboardGetLabelFromPhysical(keycode))
}

/*
Opens system emoji and symbol picker.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func ShowEmojiAndSymbolPicker() { //gd:DisplayServer.show_emoji_and_symbol_picker
	once.Do(singleton)
	class(self).ShowEmojiAndSymbolPicker()
}

/*
Perform window manager processing, including input flushing. See also [method force_process_and_drop_events], [method Input.flush_buffered_events] and [member Input.use_accumulated_input].
*/
func ProcessEvents() { //gd:DisplayServer.process_events
	once.Do(singleton)
	class(self).ProcessEvents()
}

/*
Forces window manager processing while ignoring all [InputEvent]s. See also [method process_events].
[b]Note:[/b] This method is implemented on Windows and macOS.
*/
func ForceProcessAndDropEvents() { //gd:DisplayServer.force_process_and_drop_events
	once.Do(singleton)
	class(self).ForceProcessAndDropEvents()
}

/*
Sets the window icon (usually displayed in the top-left corner) in the operating system's [i]native[/i] format. The file at [param filename] must be in [code].ico[/code] format on Windows or [code].icns[/code] on macOS. By using specially crafted [code].ico[/code] or [code].icns[/code] icons, [method set_native_icon] allows specifying different icons depending on the size the icon is displayed at. This size is determined by the operating system and user preferences (including the display scale factor). To use icons in other formats, use [method set_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_NATIVE_ICON].
*/
func SetNativeIcon(filename string) { //gd:DisplayServer.set_native_icon
	once.Do(singleton)
	class(self).SetNativeIcon(String.New(filename))
}

/*
Sets the window icon (usually displayed in the top-left corner) with an [Image]. To use icons in the operating system's native format, use [method set_native_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_ICON].
*/
func SetIcon(image [1]gdclass.Image) { //gd:DisplayServer.set_icon
	once.Do(singleton)
	class(self).SetIcon(image)
}

/*
Creates a new application status indicator with the specified icon, tooltip, and activation callback.
[param callback] should take two arguments: the pressed mouse button (one of the [enum MouseButton] constants) and the click position in screen coordinates (a [Vector2i]).
*/
func CreateStatusIndicator(icon [1]gdclass.Texture2D, tooltip string, callback func(button MouseButton, click_position Vector2i.XY)) int { //gd:DisplayServer.create_status_indicator
	once.Do(singleton)
	return int(int(class(self).CreateStatusIndicator(icon, String.New(tooltip), Callable.New(callback))))
}

/*
Sets the application status indicator icon.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func StatusIndicatorSetIcon(id int, icon [1]gdclass.Texture2D) { //gd:DisplayServer.status_indicator_set_icon
	once.Do(singleton)
	class(self).StatusIndicatorSetIcon(int64(id), icon)
}

/*
Sets the application status indicator tooltip.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func StatusIndicatorSetTooltip(id int, tooltip string) { //gd:DisplayServer.status_indicator_set_tooltip
	once.Do(singleton)
	class(self).StatusIndicatorSetTooltip(int64(id), String.New(tooltip))
}

/*
Sets the application status indicator native popup menu.
[b]Note:[/b] On macOS, the menu is activated by any mouse button. Its activation callback is [i]not[/i] triggered.
[b]Note:[/b] On Windows, the menu is activated by the right mouse button, selecting the status icon and pressing [kbd]Shift + F10[/kbd], or the applications key. The menu's activation callback for the other mouse buttons is still triggered.
[b]Note:[/b] Native popup is only supported if [NativeMenu] supports the [constant NativeMenu.FEATURE_POPUP_MENU] feature.
*/
func StatusIndicatorSetMenu(id int, menu_rid RID.NativeMenu) { //gd:DisplayServer.status_indicator_set_menu
	once.Do(singleton)
	class(self).StatusIndicatorSetMenu(int64(id), RID.Any(menu_rid))
}

/*
Sets the application status indicator activation callback. [param callback] should take two arguments: [int] mouse button index (one of [enum MouseButton] values) and [Vector2i] click position in screen coordinates.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func StatusIndicatorSetCallback(id int, callback func(button MouseButton, click_position Vector2i.XY)) { //gd:DisplayServer.status_indicator_set_callback
	once.Do(singleton)
	class(self).StatusIndicatorSetCallback(int64(id), Callable.New(callback))
}

/*
Returns the rectangle for the given status indicator [param id] in screen coordinates. If the status indicator is not visible, returns an empty [Rect2].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
func StatusIndicatorGetRect(id int) Rect2.PositionSize { //gd:DisplayServer.status_indicator_get_rect
	once.Do(singleton)
	return Rect2.PositionSize(class(self).StatusIndicatorGetRect(int64(id)))
}

/*
Removes the application status indicator.
*/
func DeleteStatusIndicator(id int) { //gd:DisplayServer.delete_status_indicator
	once.Do(singleton)
	class(self).DeleteStatusIndicator(int64(id))
}

/*
Returns the total number of available tablet drivers.
[b]Note:[/b] This method is implemented only on Windows.
*/
func TabletGetDriverCount() int { //gd:DisplayServer.tablet_get_driver_count
	once.Do(singleton)
	return int(int(class(self).TabletGetDriverCount()))
}

/*
Returns the tablet driver name for the given index.
[b]Note:[/b] This method is implemented only on Windows.
*/
func TabletGetDriverName(idx int) string { //gd:DisplayServer.tablet_get_driver_name
	once.Do(singleton)
	return string(class(self).TabletGetDriverName(int64(idx)).String())
}

/*
Returns current active tablet driver name.
[b]Note:[/b] This method is implemented only on Windows.
*/
func TabletGetCurrentDriver() string { //gd:DisplayServer.tablet_get_current_driver
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
func TabletSetCurrentDriver(name string) { //gd:DisplayServer.tablet_set_current_driver
	once.Do(singleton)
	class(self).TabletSetCurrentDriver(String.New(name))
}

/*
Returns [code]true[/code] if the window background can be made transparent. This method returns [code]false[/code] if [member ProjectSettings.display/window/per_pixel_transparency/allowed] is set to [code]false[/code], or if transparency is not supported by the renderer or OS compositor.
*/
func IsWindowTransparencyAvailable() bool { //gd:DisplayServer.is_window_transparency_available
	once.Do(singleton)
	return bool(class(self).IsWindowTransparencyAvailable())
}

/*
Registers an [Object] which represents an additional output that will be rendered too, beyond normal windows. The [Object] is only used as an identifier, which can be later passed to [method unregister_additional_output].
This can be used to prevent Godot from skipping rendering when no normal windows are visible.
*/
func RegisterAdditionalOutput(obj Object.Instance) { //gd:DisplayServer.register_additional_output
	once.Do(singleton)
	class(self).RegisterAdditionalOutput(obj)
}

/*
Unregisters an [Object] representing an additional output, that was registered via [method register_additional_output].
*/
func UnregisterAdditionalOutput(obj Object.Instance) { //gd:DisplayServer.unregister_additional_output
	once.Do(singleton)
	class(self).UnregisterAdditionalOutput(obj)
}

/*
Returns [code]true[/code] if any additional outputs have been registered via [method register_additional_output].
*/
func HasAdditionalOutputs() bool { //gd:DisplayServer.has_additional_outputs
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
func (self class) HasFeature(feature gdclass.DisplayServerFeature) bool { //gd:DisplayServer.has_feature
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
func (self class) GetName() String.Readable { //gd:DisplayServer.get_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
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
func (self class) HelpSetSearchCallbacks(search_callback Callable.Function, action_callback Callable.Function) { //gd:DisplayServer.help_set_search_callbacks
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(search_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(action_callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_help_set_search_callbacks, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Registers callables to emit when the menu is respectively about to show or closed. Callback methods should have zero arguments.
*/
//go:nosplit
func (self class) GlobalMenuSetPopupCallbacks(menu_root String.Readable, open_callback Callable.Function, close_callback Callable.Function) { //gd:DisplayServer.global_menu_set_popup_callbacks
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(open_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(close_callback)))
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
func (self class) GlobalMenuAddSubmenuItem(menu_root String.Readable, label String.Readable, submenu String.Readable, index int64) int64 { //gd:DisplayServer.global_menu_add_submenu_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(submenu)))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuAddItem(menu_root String.Readable, label String.Readable, callback Callable.Function, key_callback Callable.Function, tag variant.Any, accelerator Key, index int64) int64 { //gd:DisplayServer.global_menu_add_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(key_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuAddCheckItem(menu_root String.Readable, label String.Readable, callback Callable.Function, key_callback Callable.Function, tag variant.Any, accelerator Key, index int64) int64 { //gd:DisplayServer.global_menu_add_check_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(key_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuAddIconItem(menu_root String.Readable, icon [1]gdclass.Texture2D, label String.Readable, callback Callable.Function, key_callback Callable.Function, tag variant.Any, accelerator Key, index int64) int64 { //gd:DisplayServer.global_menu_add_icon_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(key_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuAddIconCheckItem(menu_root String.Readable, icon [1]gdclass.Texture2D, label String.Readable, callback Callable.Function, key_callback Callable.Function, tag variant.Any, accelerator Key, index int64) int64 { //gd:DisplayServer.global_menu_add_icon_check_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(key_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuAddRadioCheckItem(menu_root String.Readable, label String.Readable, callback Callable.Function, key_callback Callable.Function, tag variant.Any, accelerator Key, index int64) int64 { //gd:DisplayServer.global_menu_add_radio_check_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(key_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuAddIconRadioCheckItem(menu_root String.Readable, icon [1]gdclass.Texture2D, label String.Readable, callback Callable.Function, key_callback Callable.Function, tag variant.Any, accelerator Key, index int64) int64 { //gd:DisplayServer.global_menu_add_icon_radio_check_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(key_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuAddMultistateItem(menu_root String.Readable, label String.Readable, max_states int64, default_state int64, callback Callable.Function, key_callback Callable.Function, tag variant.Any, accelerator Key, index int64) int64 { //gd:DisplayServer.global_menu_add_multistate_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(label)))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(key_callback)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	callframe.Arg(frame, accelerator)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuAddSeparator(menu_root String.Readable, index int64) int64 { //gd:DisplayServer.global_menu_add_separator
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuGetItemIndexFromText(menu_root String.Readable, text String.Readable) int64 { //gd:DisplayServer.global_menu_get_item_index_from_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuGetItemIndexFromTag(menu_root String.Readable, tag variant.Any) int64 { //gd:DisplayServer.global_menu_get_item_index_from_tag
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuIsItemChecked(menu_root String.Readable, idx int64) bool { //gd:DisplayServer.global_menu_is_item_checked
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuIsItemCheckable(menu_root String.Readable, idx int64) bool { //gd:DisplayServer.global_menu_is_item_checkable
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuIsItemRadioCheckable(menu_root String.Readable, idx int64) bool { //gd:DisplayServer.global_menu_is_item_radio_checkable
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuGetItemCallback(menu_root String.Readable, idx int64) Callable.Function { //gd:DisplayServer.global_menu_get_item_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the callback of the item accelerator at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemKeyCallback(menu_root String.Readable, idx int64) Callable.Function { //gd:DisplayServer.global_menu_get_item_key_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method global_menu_set_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemTag(menu_root String.Readable, idx int64) variant.Any { //gd:DisplayServer.global_menu_get_item_tag
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_tag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemText(menu_root String.Readable, idx int64) String.Readable { //gd:DisplayServer.global_menu_get_item_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the submenu ID of the item at index [param idx]. See [method global_menu_add_submenu_item] for more info on how to add a submenu.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemSubmenu(menu_root String.Readable, idx int64) String.Readable { //gd:DisplayServer.global_menu_get_item_submenu
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_submenu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the accelerator of the item at index [param idx]. Accelerators are special combinations of keys that activate the item, no matter which control is focused.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemAccelerator(menu_root String.Readable, idx int64) Key { //gd:DisplayServer.global_menu_get_item_accelerator
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuIsItemDisabled(menu_root String.Readable, idx int64) bool { //gd:DisplayServer.global_menu_is_item_disabled
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuIsItemHidden(menu_root String.Readable, idx int64) bool { //gd:DisplayServer.global_menu_is_item_hidden
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuGetItemTooltip(menu_root String.Readable, idx int64) String.Readable { //gd:DisplayServer.global_menu_get_item_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetItemState(menu_root String.Readable, idx int64) int64 { //gd:DisplayServer.global_menu_get_item_state
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuGetItemMaxStates(menu_root String.Readable, idx int64) int64 { //gd:DisplayServer.global_menu_get_item_max_states
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuGetItemIcon(menu_root String.Readable, idx int64) [1]gdclass.Texture2D { //gd:DisplayServer.global_menu_get_item_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuGetItemIndentationLevel(menu_root String.Readable, idx int64) int64 { //gd:DisplayServer.global_menu_get_item_indentation_level
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuSetItemChecked(menu_root String.Readable, idx int64, checked bool) { //gd:DisplayServer.global_menu_set_item_checked
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemCheckable(menu_root String.Readable, idx int64, checkable bool) { //gd:DisplayServer.global_menu_set_item_checkable
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemRadioCheckable(menu_root String.Readable, idx int64, checkable bool) { //gd:DisplayServer.global_menu_set_item_radio_checkable
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemCallback(menu_root String.Readable, idx int64, callback Callable.Function) { //gd:DisplayServer.global_menu_set_item_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
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
func (self class) GlobalMenuSetItemHoverCallbacks(menu_root String.Readable, idx int64, callback Callable.Function) { //gd:DisplayServer.global_menu_set_item_hover_callbacks
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
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
func (self class) GlobalMenuSetItemKeyCallback(menu_root String.Readable, idx int64, key_callback Callable.Function) { //gd:DisplayServer.global_menu_set_item_key_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(key_callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_key_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method global_menu_get_item_tag], which provides a simple way of assigning context data to items.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemTag(menu_root String.Readable, idx int64, tag variant.Any) { //gd:DisplayServer.global_menu_set_item_tag
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(tag)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_tag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the text of the item at index [param idx].
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemText(menu_root String.Readable, idx int64, text String.Readable) { //gd:DisplayServer.global_menu_set_item_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the submenu of the item at index [param idx]. The submenu is the ID of a global menu root that would be shown when the item is clicked.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemSubmenu(menu_root String.Readable, idx int64, submenu String.Readable) { //gd:DisplayServer.global_menu_set_item_submenu
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalString(submenu)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_submenu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the accelerator of the item at index [param idx]. [param keycode] can be a single [enum Key], or a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemAccelerator(menu_root String.Readable, idx int64, keycode Key) { //gd:DisplayServer.global_menu_set_item_accelerator
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemDisabled(menu_root String.Readable, idx int64, disabled bool) { //gd:DisplayServer.global_menu_set_item_disabled
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemHidden(menu_root String.Readable, idx int64, hidden bool) { //gd:DisplayServer.global_menu_set_item_hidden
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemTooltip(menu_root String.Readable, idx int64, tooltip String.Readable) { //gd:DisplayServer.global_menu_set_item_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the state of a multistate item. See [method global_menu_add_multistate_item] for details.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuSetItemState(menu_root String.Readable, idx int64, state int64) { //gd:DisplayServer.global_menu_set_item_state
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemMaxStates(menu_root String.Readable, idx int64, max_states int64) { //gd:DisplayServer.global_menu_set_item_max_states
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemIcon(menu_root String.Readable, idx int64, icon [1]gdclass.Texture2D) { //gd:DisplayServer.global_menu_set_item_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuSetItemIndentationLevel(menu_root String.Readable, idx int64, level int64) { //gd:DisplayServer.global_menu_set_item_indentation_level
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuGetItemCount(menu_root String.Readable) int64 { //gd:DisplayServer.global_menu_get_item_count
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) GlobalMenuRemoveItem(menu_root String.Readable, idx int64) { //gd:DisplayServer.global_menu_remove_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
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
func (self class) GlobalMenuClear(menu_root String.Readable) { //gd:DisplayServer.global_menu_clear
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(menu_root)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns Dictionary of supported system menu IDs and names.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) GlobalMenuGetSystemMenuRoots() Dictionary.Any { //gd:DisplayServer.global_menu_get_system_menu_roots
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_global_menu_get_system_menu_roots, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the synthesizer is generating speech, or have utterance waiting in the queue.
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsIsSpeaking() bool { //gd:DisplayServer.tts_is_speaking
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
func (self class) TtsIsPaused() bool { //gd:DisplayServer.tts_is_paused
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
func (self class) TtsGetVoices() Array.Contains[Dictionary.Any] { //gd:DisplayServer.tts_get_voices
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_get_voices, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an [PackedStringArray] of voice identifiers for the [param language].
[b]Note:[/b] This method is implemented on Android, iOS, Web, Linux (X11/Wayland), macOS, and Windows.
[b]Note:[/b] [member ProjectSettings.audio/general/text_to_speech] should be [code]true[/code] to use text-to-speech.
*/
//go:nosplit
func (self class) TtsGetVoicesForLanguage(language String.Readable) Packed.Strings { //gd:DisplayServer.tts_get_voices_for_language
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_get_voices_for_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
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
func (self class) TtsSpeak(text String.Readable, voice String.Readable, volume int64, pitch float64, rate float64, utterance_id int64, interrupt bool) { //gd:DisplayServer.tts_speak
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(voice)))
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
func (self class) TtsPause() { //gd:DisplayServer.tts_pause
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
func (self class) TtsResume() { //gd:DisplayServer.tts_resume
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
func (self class) TtsStop() { //gd:DisplayServer.tts_stop
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
func (self class) TtsSetUtteranceCallback(event gdclass.DisplayServerTTSUtteranceEvent, callable Callable.Function) { //gd:DisplayServer.tts_set_utterance_callback
	var frame = callframe.New()
	callframe.Arg(frame, event)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tts_set_utterance_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if OS supports dark mode.
[b]Note:[/b] This method is implemented on Android, iOS, macOS, Windows, and Linux (X11/Wayland).
*/
//go:nosplit
func (self class) IsDarkModeSupported() bool { //gd:DisplayServer.is_dark_mode_supported
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
func (self class) IsDarkMode() bool { //gd:DisplayServer.is_dark_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_is_dark_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns OS theme accent color. Returns [code]Color(0, 0, 0, 0)[/code], if accent color is unknown.
[b]Note:[/b] This method is implemented on macOS, Windows, and Android.
*/
//go:nosplit
func (self class) GetAccentColor() Color.RGBA { //gd:DisplayServer.get_accent_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_accent_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the OS theme base color (default control background). Returns [code]Color(0, 0, 0, 0)[/code] if the base color is unknown.
[b]Note:[/b] This method is implemented on macOS, Windows, and Android.
*/
//go:nosplit
func (self class) GetBaseColor() Color.RGBA { //gd:DisplayServer.get_base_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
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
func (self class) SetSystemThemeChangeCallback(callable Callable.Function) { //gd:DisplayServer.set_system_theme_change_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_set_system_theme_change_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the current mouse mode. See also [method mouse_get_mode].
*/
//go:nosplit
func (self class) MouseSetMode(mouse_mode gdclass.DisplayServerMouseMode) { //gd:DisplayServer.mouse_set_mode
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
func (self class) MouseGetMode() gdclass.DisplayServerMouseMode { //gd:DisplayServer.mouse_get_mode
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
func (self class) WarpMouse(position Vector2i.XY) { //gd:DisplayServer.warp_mouse
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
func (self class) MouseGetPosition() Vector2i.XY { //gd:DisplayServer.mouse_get_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_mouse_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current state of mouse buttons (whether each button is pressed) as a bitmask. If multiple mouse buttons are pressed at the same time, the bits are added together. Equivalent to [method Input.get_mouse_button_mask].
*/
//go:nosplit
func (self class) MouseGetButtonState() MouseButtonMask { //gd:DisplayServer.mouse_get_button_state
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
func (self class) ClipboardSet(clipboard String.Readable) { //gd:DisplayServer.clipboard_set
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(clipboard)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the user's clipboard as a string if possible.
*/
//go:nosplit
func (self class) ClipboardGet() String.Readable { //gd:DisplayServer.clipboard_get
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_get, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the user's clipboard as an image if possible.
[b]Note:[/b] This method uses the copied pixel data, e.g. from a image editing software or a web browser, not an image file copied from file explorer.
*/
//go:nosplit
func (self class) ClipboardGetImage() [1]gdclass.Image { //gd:DisplayServer.clipboard_get_image
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
func (self class) ClipboardHas() bool { //gd:DisplayServer.clipboard_has
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
func (self class) ClipboardHasImage() bool { //gd:DisplayServer.clipboard_has_image
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
func (self class) ClipboardSetPrimary(clipboard_primary String.Readable) { //gd:DisplayServer.clipboard_set_primary
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(clipboard_primary)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_set_primary, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the user's [url=https://unix.stackexchange.com/questions/139191/whats-the-difference-between-primary-selection-and-clipboard-buffer]primary[/url] clipboard as a string if possible. This is the clipboard that is set when the user selects text in any application, rather than when pressing [kbd]Ctrl + C[/kbd]. The clipboard data can then be pasted by clicking the middle mouse button in any application that supports the primary clipboard mechanism.
[b]Note:[/b] This method is only implemented on Linux (X11/Wayland).
*/
//go:nosplit
func (self class) ClipboardGetPrimary() String.Readable { //gd:DisplayServer.clipboard_get_primary
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_clipboard_get_primary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an [Array] of [Rect2], each of which is the bounding rectangle for a display cutout or notch. These are non-functional areas on edge-to-edge screens used by cameras and sensors. Returns an empty array if the device does not have cutouts. See also [method get_display_safe_area].
[b]Note:[/b] Currently only implemented on Android. Other platforms will return an empty array even if they do have display cutouts or notches.
*/
//go:nosplit
func (self class) GetDisplayCutouts() Array.Contains[Rect2.PositionSize] { //gd:DisplayServer.get_display_cutouts
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_display_cutouts, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Rect2.PositionSize]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the unobscured area of the display where interactive controls should be rendered. See also [method get_display_cutouts].
[b]Note:[/b] Currently only implemented on Android and iOS. On other platforms, [code]screen_get_usable_rect(SCREEN_OF_MAIN_WINDOW)[/code] will be returned as a fallback. See also [method screen_get_usable_rect].
*/
//go:nosplit
func (self class) GetDisplaySafeArea() Rect2i.PositionSize { //gd:DisplayServer.get_display_safe_area
	var frame = callframe.New()
	var r_ret = callframe.Ret[Rect2i.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_display_safe_area, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of displays available.
*/
//go:nosplit
func (self class) GetScreenCount() int64 { //gd:DisplayServer.get_screen_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_screen_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns index of the primary screen.
*/
//go:nosplit
func (self class) GetPrimaryScreen() int64 { //gd:DisplayServer.get_primary_screen
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_primary_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the screen containing the window with the keyboard focus, or the primary screen if there's no focused window.
*/
//go:nosplit
func (self class) GetKeyboardFocusScreen() int64 { //gd:DisplayServer.get_keyboard_focus_screen
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_keyboard_focus_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the screen that overlaps the most with the given rectangle. Returns [code]-1[/code] if the rectangle doesn't overlap with any screen or has no area.
*/
//go:nosplit
func (self class) GetScreenFromRect(rect Rect2.PositionSize) int64 { //gd:DisplayServer.get_screen_from_rect
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) ScreenGetPosition(screen int64) Vector2i.XY { //gd:DisplayServer.screen_get_position
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the screen's size in pixels. See also [method screen_get_position] and [method screen_get_usable_rect].
*/
//go:nosplit
func (self class) ScreenGetSize(screen int64) Vector2i.XY { //gd:DisplayServer.screen_get_size
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the portion of the screen that is not obstructed by a status bar in pixels. See also [method screen_get_size].
*/
//go:nosplit
func (self class) ScreenGetUsableRect(screen int64) Rect2i.PositionSize { //gd:DisplayServer.screen_get_usable_rect
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[Rect2i.PositionSize](frame)
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
func (self class) ScreenGetDpi(screen int64) int64 { //gd:DisplayServer.screen_get_dpi
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_dpi, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scale factor of the specified screen by index.
[b]Note:[/b] On macOS, the returned value is [code]2.0[/code] for hiDPI (Retina) screens, and [code]1.0[/code] for all other cases.
[b]Note:[/b] On Linux (Wayland), the returned value is accurate only when [param screen] is [constant SCREEN_OF_MAIN_WINDOW]. Due to API limitations, passing a direct index will return a rounded-up integer, if the screen has a fractional scale (e.g. [code]1.25[/code] would get rounded up to [code]2.0[/code]).
[b]Note:[/b] This method is implemented on Android, iOS, Web, macOS, and Linux (Wayland).
*/
//go:nosplit
func (self class) ScreenGetScale(screen int64) float64 { //gd:DisplayServer.screen_get_scale
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if touch events are available (Android or iOS), the capability is detected on the Web platform or if [member ProjectSettings.input_devices/pointing/emulate_touch_from_mouse] is [code]true[/code].
*/
//go:nosplit
func (self class) IsTouchscreenAvailable() bool { //gd:DisplayServer.is_touchscreen_available
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
func (self class) ScreenGetMaxScale() float64 { //gd:DisplayServer.screen_get_max_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
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
func (self class) ScreenGetRefreshRate(screen int64) float64 { //gd:DisplayServer.screen_get_refresh_rate
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[float64](frame)
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
func (self class) ScreenGetPixel(position Vector2i.XY) Color.RGBA { //gd:DisplayServer.screen_get_pixel
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[Color.RGBA](frame)
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
func (self class) ScreenGetImage(screen int64) [1]gdclass.Image { //gd:DisplayServer.screen_get_image
	var frame = callframe.New()
	callframe.Arg(frame, screen)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns screenshot of the screen [param rect].
[b]Note:[/b] This method is implemented on macOS and Windows.
[b]Note:[/b] On macOS, this method requires "Screen Recording" permission, if permission is not granted it will return desktop wallpaper color.
*/
//go:nosplit
func (self class) ScreenGetImageRect(rect Rect2i.PositionSize) [1]gdclass.Image { //gd:DisplayServer.screen_get_image_rect
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_screen_get_image_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the [param screen]'s [param orientation]. See also [method screen_get_orientation].
[b]Note:[/b] On iOS, this method has no effect if [member ProjectSettings.display/window/handheld/orientation] is not set to [constant SCREEN_SENSOR].
*/
//go:nosplit
func (self class) ScreenSetOrientation(orientation gdclass.DisplayServerScreenOrientation, screen int64) { //gd:DisplayServer.screen_set_orientation
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
func (self class) ScreenGetOrientation(screen int64) gdclass.DisplayServerScreenOrientation { //gd:DisplayServer.screen_get_orientation
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
func (self class) ScreenSetKeepOn(enable bool) { //gd:DisplayServer.screen_set_keep_on
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
func (self class) ScreenIsKeptOn() bool { //gd:DisplayServer.screen_is_kept_on
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
func (self class) GetWindowList() Packed.Array[int32] { //gd:DisplayServer.get_window_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_get_window_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
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
func (self class) GetWindowAtScreenPosition(position Vector2i.XY) int64 { //gd:DisplayServer.get_window_at_screen_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) WindowGetNativeHandle(handle_type gdclass.DisplayServerHandleType, window_id int64) int64 { //gd:DisplayServer.window_get_native_handle
	var frame = callframe.New()
	callframe.Arg(frame, handle_type)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_native_handle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns ID of the active popup window, or [constant INVALID_WINDOW_ID] if there is none.
*/
//go:nosplit
func (self class) WindowGetActivePopup() int64 { //gd:DisplayServer.window_get_active_popup
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_active_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the bounding box of control, or menu item that was used to open the popup window, in the screen coordinate system. Clicking this area will not auto-close this popup.
*/
//go:nosplit
func (self class) WindowSetPopupSafeRect(window int64, rect Rect2i.PositionSize) { //gd:DisplayServer.window_set_popup_safe_rect
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
func (self class) WindowGetPopupSafeRect(window int64) Rect2i.PositionSize { //gd:DisplayServer.window_get_popup_safe_rect
	var frame = callframe.New()
	callframe.Arg(frame, window)
	var r_ret = callframe.Ret[Rect2i.PositionSize](frame)
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
func (self class) WindowSetTitle(title String.Readable, window_id int64) { //gd:DisplayServer.window_set_title
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
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
func (self class) WindowGetTitleSize(title String.Readable, window_id int64) Vector2i.XY { //gd:DisplayServer.window_get_title_size
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
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
DisplayServer.WindowSetMousePassthrough([]);
[/csharp]
[/codeblocks]
[b]Note:[/b] On Windows, the portion of a window that lies outside the region is not drawn, while on Linux (X11) and macOS it is.
[b]Note:[/b] This method is implemented on Linux (X11), macOS and Windows.
*/
//go:nosplit
func (self class) WindowSetMousePassthrough(region Packed.Array[Vector2.XY], window_id int64) { //gd:DisplayServer.window_set_mouse_passthrough
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](region)))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_mouse_passthrough, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the screen the window specified by [param window_id] is currently positioned on. If the screen overlaps multiple displays, the screen where the window's center is located is returned. See also [method window_set_current_screen].
*/
//go:nosplit
func (self class) WindowGetCurrentScreen(window_id int64) int64 { //gd:DisplayServer.window_get_current_screen
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_current_screen, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves the window specified by [param window_id] to the specified [param screen]. See also [method window_get_current_screen].
*/
//go:nosplit
func (self class) WindowSetCurrentScreen(screen int64, window_id int64) { //gd:DisplayServer.window_set_current_screen
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
func (self class) WindowGetPosition(window_id int64) Vector2i.XY { //gd:DisplayServer.window_get_position
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position of the given window on the screen including the borders drawn by the operating system. See also [method window_get_position].
*/
//go:nosplit
func (self class) WindowGetPositionWithDecorations(window_id int64) Vector2i.XY { //gd:DisplayServer.window_get_position_with_decorations
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
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
func (self class) WindowSetPosition(position Vector2i.XY, window_id int64) { //gd:DisplayServer.window_set_position
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
func (self class) WindowGetSize(window_id int64) Vector2i.XY { //gd:DisplayServer.window_get_size
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
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
func (self class) WindowSetSize(size Vector2i.XY, window_id int64) { //gd:DisplayServer.window_set_size
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
func (self class) WindowSetRectChangedCallback(callback Callable.Function, window_id int64) { //gd:DisplayServer.window_set_rect_changed_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
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
func (self class) WindowSetWindowEventCallback(callback Callable.Function, window_id int64) { //gd:DisplayServer.window_set_window_event_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
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
func (self class) WindowSetInputEventCallback(callback Callable.Function, window_id int64) { //gd:DisplayServer.window_set_input_event_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
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
func (self class) WindowSetInputTextCallback(callback Callable.Function, window_id int64) { //gd:DisplayServer.window_set_input_text_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
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
func (self class) WindowSetDropFilesCallback(callback Callable.Function, window_id int64) { //gd:DisplayServer.window_set_drop_files_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_set_drop_files_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [method Object.get_instance_id] of the [Window] the [param window_id] is attached to.
*/
//go:nosplit
func (self class) WindowGetAttachedInstanceId(window_id int64) int64 { //gd:DisplayServer.window_get_attached_instance_id
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_attached_instance_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the window's maximum size (in pixels). See also [method window_set_max_size].
*/
//go:nosplit
func (self class) WindowGetMaxSize(window_id int64) Vector2i.XY { //gd:DisplayServer.window_get_max_size
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
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
func (self class) WindowSetMaxSize(max_size Vector2i.XY, window_id int64) { //gd:DisplayServer.window_set_max_size
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
func (self class) WindowGetMinSize(window_id int64) Vector2i.XY { //gd:DisplayServer.window_get_min_size
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
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
func (self class) WindowSetMinSize(min_size Vector2i.XY, window_id int64) { //gd:DisplayServer.window_set_min_size
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
func (self class) WindowGetSizeWithDecorations(window_id int64) Vector2i.XY { //gd:DisplayServer.window_get_size_with_decorations
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_size_with_decorations, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the mode of the given window.
*/
//go:nosplit
func (self class) WindowGetMode(window_id int64) gdclass.DisplayServerWindowMode { //gd:DisplayServer.window_get_mode
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
[b]Note:[/b] On Android, setting it to [constant WINDOW_MODE_FULLSCREEN] or [constant WINDOW_MODE_EXCLUSIVE_FULLSCREEN] will enable immersive mode.
[b]Note:[/b] Setting the window to full screen forcibly sets the borderless flag to [code]true[/code], so make sure to set it back to [code]false[/code] when not wanted.
*/
//go:nosplit
func (self class) WindowSetMode(mode gdclass.DisplayServerWindowMode, window_id int64) { //gd:DisplayServer.window_set_mode
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
func (self class) WindowSetFlag(flag gdclass.DisplayServerWindowFlags, enabled bool, window_id int64) { //gd:DisplayServer.window_set_flag
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
func (self class) WindowGetFlag(flag gdclass.DisplayServerWindowFlags, window_id int64) bool { //gd:DisplayServer.window_get_flag
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
func (self class) WindowSetWindowButtonsOffset(offset Vector2i.XY, window_id int64) { //gd:DisplayServer.window_set_window_buttons_offset
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
func (self class) WindowGetSafeTitleMargins(window_id int64) Vector3i.XYZ { //gd:DisplayServer.window_get_safe_title_margins
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[Vector3i.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_get_safe_title_margins, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Makes the window specified by [param window_id] request attention, which is materialized by the window title and taskbar entry blinking until the window is focused. This usually has no visible effect if the window is currently focused. The exact behavior varies depending on the operating system.
*/
//go:nosplit
func (self class) WindowRequestAttention(window_id int64) { //gd:DisplayServer.window_request_attention
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
func (self class) WindowMoveToForeground(window_id int64) { //gd:DisplayServer.window_move_to_foreground
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
func (self class) WindowIsFocused(window_id int64) bool { //gd:DisplayServer.window_is_focused
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
func (self class) WindowCanDraw(window_id int64) bool { //gd:DisplayServer.window_can_draw
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_can_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets window transient parent. Transient window will be destroyed with its transient parent and will return focus to their parent when closed. The transient window is displayed on top of a non-exclusive full-screen parent window. Transient windows can't enter full-screen mode.
[b]Note:[/b] It's recommended to change this value using [member Window.transient] instead.
[b]Note:[/b] The behavior might be different depending on the platform.
*/
//go:nosplit
func (self class) WindowSetTransient(window_id int64, parent_window_id int64) { //gd:DisplayServer.window_set_transient
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
func (self class) WindowSetExclusive(window_id int64, exclusive bool) { //gd:DisplayServer.window_set_exclusive
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
func (self class) WindowSetImeActive(active bool, window_id int64) { //gd:DisplayServer.window_set_ime_active
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
func (self class) WindowSetImePosition(position Vector2i.XY, window_id int64) { //gd:DisplayServer.window_set_ime_position
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
func (self class) WindowSetVsyncMode(vsync_mode gdclass.DisplayServerVSyncMode, window_id int64) { //gd:DisplayServer.window_set_vsync_mode
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
func (self class) WindowGetVsyncMode(window_id int64) gdclass.DisplayServerVSyncMode { //gd:DisplayServer.window_get_vsync_mode
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
func (self class) WindowIsMaximizeAllowed(window_id int64) bool { //gd:DisplayServer.window_is_maximize_allowed
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
func (self class) WindowMaximizeOnTitleDblClick() bool { //gd:DisplayServer.window_maximize_on_title_dbl_click
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
func (self class) WindowMinimizeOnTitleDblClick() bool { //gd:DisplayServer.window_minimize_on_title_dbl_click
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_minimize_on_title_dbl_click, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Starts an interactive drag operation on the window with the given [param window_id], using the current mouse position. Call this method when handling a mouse button being pressed to simulate a pressed event on the window's title bar. Using this method allows the window to participate in space switching, tiling, and other system features.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS, and Windows.
*/
//go:nosplit
func (self class) WindowStartDrag(window_id int64) { //gd:DisplayServer.window_start_drag
	var frame = callframe.New()
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_start_drag, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Starts an interactive resize operation on the window with the given [param window_id], using the current mouse position. Call this method when handling a mouse button being pressed to simulate a pressed event on the window's edge.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS, and Windows.
*/
//go:nosplit
func (self class) WindowStartResize(edge gdclass.DisplayServerWindowResizeEdge, window_id int64) { //gd:DisplayServer.window_start_resize
	var frame = callframe.New()
	callframe.Arg(frame, edge)
	callframe.Arg(frame, window_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_window_start_resize, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the text selection in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] composition string, with the [Vector2i]'s [code]x[/code] component being the caret position and [code]y[/code] being the length of the selection.
[b]Note:[/b] This method is implemented only on macOS.
*/
//go:nosplit
func (self class) ImeGetSelection() Vector2i.XY { //gd:DisplayServer.ime_get_selection
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2i.XY](frame)
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
func (self class) ImeGetText() String.Readable { //gd:DisplayServer.ime_get_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_ime_get_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
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
func (self class) VirtualKeyboardShow(existing_text String.Readable, position Rect2.PositionSize, atype gdclass.DisplayServerVirtualKeyboardType, max_length int64, cursor_start int64, cursor_end int64) { //gd:DisplayServer.virtual_keyboard_show
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(existing_text)))
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
func (self class) VirtualKeyboardHide() { //gd:DisplayServer.virtual_keyboard_hide
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_virtual_keyboard_hide, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the on-screen keyboard's height in pixels. Returns 0 if there is no keyboard or if it is currently hidden.
*/
//go:nosplit
func (self class) VirtualKeyboardGetHeight() int64 { //gd:DisplayServer.virtual_keyboard_get_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_virtual_keyboard_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if hardware keyboard is connected.
[b]Note:[/b] This method is implemented on Android and iOS, on other platforms this method always returns [code]true[/code].
*/
//go:nosplit
func (self class) HasHardwareKeyboard() bool { //gd:DisplayServer.has_hardware_keyboard
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_has_hardware_keyboard, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the default mouse cursor shape. The cursor's appearance will vary depending on the user's operating system and mouse cursor theme. See also [method cursor_get_shape] and [method cursor_set_custom_image].
*/
//go:nosplit
func (self class) CursorSetShape(shape gdclass.DisplayServerCursorShape) { //gd:DisplayServer.cursor_set_shape
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
func (self class) CursorGetShape() gdclass.DisplayServerCursorShape { //gd:DisplayServer.cursor_get_shape
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
func (self class) CursorSetCustomImage(cursor [1]gdclass.Resource, shape gdclass.DisplayServerCursorShape, hotspot Vector2.XY) { //gd:DisplayServer.cursor_set_custom_image
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
func (self class) GetSwapCancelOk() bool { //gd:DisplayServer.get_swap_cancel_ok
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
func (self class) EnableForStealingFocus(process_id int64) { //gd:DisplayServer.enable_for_stealing_focus
	var frame = callframe.New()
	callframe.Arg(frame, process_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_enable_for_stealing_focus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Shows a text dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [int] parameter which corresponds to the index of the pressed button.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG] feature. Supported platforms include macOS, Windows, and Android.
*/
//go:nosplit
func (self class) DialogShow(title String.Readable, description String.Readable, buttons Packed.Strings, callback Callable.Function) Error.Code { //gd:DisplayServer.dialog_show
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(description)))
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(buttons)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_dialog_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Shows a text input dialog which uses the operating system's native look-and-feel. [param callback] should accept a single [String] parameter which contains the text field's contents.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_INPUT] feature. Supported platforms include macOS, Windows, and Android.
*/
//go:nosplit
func (self class) DialogInputText(title String.Readable, description String.Readable, existing_text String.Readable, callback Callable.Function) Error.Code { //gd:DisplayServer.dialog_input_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(description)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(existing_text)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_dialog_input_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Displays OS native dialog for selecting files or directories in the file system.
Each filter string in the [param filters] array should be formatted like this: [code]*.png,*.jpg,*.jpeg;Image Files;image/png,image/jpeg[/code]. The description text of the filter is optional and can be omitted. It is recommended to set both file extension and MIME type. See also [member FileDialog.filters].
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int[/code]. [b]On Android,[/b] callback argument [code]selected_filter_index[/code] is always zero.
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE] feature. Supported platforms include Linux (X11/Wayland), Windows, macOS, and Android.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] Embedded file dialog and Windows file dialog support only file extensions, while Android, Linux, and macOS file dialogs also support MIME types.
[b]Note:[/b] On Android and Linux, [param show_hidden] is ignored.
[b]Note:[/b] On Android and macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
//go:nosplit
func (self class) FileDialogShow(title String.Readable, current_directory String.Readable, filename String.Readable, show_hidden bool, mode gdclass.DisplayServerFileDialogMode, filters Packed.Strings, callback Callable.Function) Error.Code { //gd:DisplayServer.file_dialog_show
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(current_directory)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(filename)))
	callframe.Arg(frame, show_hidden)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(filters)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_file_dialog_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Displays OS native dialog for selecting files or directories in the file system with additional user selectable options.
Each filter string in the [param filters] array should be formatted like this: [code]*.png,*.jpg,*.jpeg;Image Files;image/png,image/jpeg[/code]. The description text of the filter is optional and can be omitted. It is recommended to set both file extension and MIME type. See also [member FileDialog.filters].
[param options] is array of [Dictionary]s with the following keys:
- [code]"name"[/code] - option's name [String].
- [code]"values"[/code] - [PackedStringArray] of values. If empty, boolean option (check box) is used.
- [code]"default"[/code] - default selected option index ([int]) or default boolean value ([bool]).
Callbacks have the following arguments: [code]status: bool, selected_paths: PackedStringArray, selected_filter_index: int, selected_option: Dictionary[/code].
[b]Note:[/b] This method is implemented if the display server has the [constant FEATURE_NATIVE_DIALOG_FILE_EXTRA] feature. Supported platforms include Linux (X11/Wayland), Windows, and macOS.
[b]Note:[/b] [param current_directory] might be ignored.
[b]Note:[/b] Embedded file dialog and Windows file dialog support only file extensions, while Android, Linux, and macOS file dialogs also support MIME types.
[b]Note:[/b] On Linux (X11), [param show_hidden] is ignored.
[b]Note:[/b] On macOS, native file dialogs have no title.
[b]Note:[/b] On macOS, sandboxed apps will save security-scoped bookmarks to retain access to the opened folders across multiple sessions. Use [method OS.get_granted_permissions] to get a list of saved bookmarks.
*/
//go:nosplit
func (self class) FileDialogWithOptionsShow(title String.Readable, current_directory String.Readable, root String.Readable, filename String.Readable, show_hidden bool, mode gdclass.DisplayServerFileDialogMode, filters Packed.Strings, options Array.Contains[Dictionary.Any], callback Callable.Function) Error.Code { //gd:DisplayServer.file_dialog_with_options_show
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(current_directory)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(root)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(filename)))
	callframe.Arg(frame, show_hidden)
	callframe.Arg(frame, mode)
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(filters)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(options)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_file_dialog_with_options_show, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Plays the beep sound from the operative system, if possible. Because it comes from the OS, the beep sound will be audible even if the application is muted. It may also be disabled for the entire OS by the user.
[b]Note:[/b] This method is implemented on macOS, Linux (X11/Wayland), and Windows.
*/
//go:nosplit
func (self class) Beep() { //gd:DisplayServer.beep
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_beep, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of keyboard layouts.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLayoutCount() int64 { //gd:DisplayServer.keyboard_get_layout_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) KeyboardGetCurrentLayout() int64 { //gd:DisplayServer.keyboard_get_current_layout
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) KeyboardSetCurrentLayout(index int64) { //gd:DisplayServer.keyboard_set_current_layout
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
func (self class) KeyboardGetLayoutLanguage(index int64) String.Readable { //gd:DisplayServer.keyboard_get_layout_language
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_layout_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the localized name of the keyboard layout at position [param index].
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetLayoutName(index int64) String.Readable { //gd:DisplayServer.keyboard_get_layout_name
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_layout_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Converts a physical (US QWERTY) [param keycode] to one in the active keyboard layout.
[b]Note:[/b] This method is implemented on Linux (X11/Wayland), macOS and Windows.
*/
//go:nosplit
func (self class) KeyboardGetKeycodeFromPhysical(keycode Key) Key { //gd:DisplayServer.keyboard_get_keycode_from_physical
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
func (self class) KeyboardGetLabelFromPhysical(keycode Key) Key { //gd:DisplayServer.keyboard_get_label_from_physical
	var frame = callframe.New()
	callframe.Arg(frame, keycode)
	var r_ret = callframe.Ret[Key](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_keyboard_get_label_from_physical, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Opens system emoji and symbol picker.
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) ShowEmojiAndSymbolPicker() { //gd:DisplayServer.show_emoji_and_symbol_picker
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_show_emoji_and_symbol_picker, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Perform window manager processing, including input flushing. See also [method force_process_and_drop_events], [method Input.flush_buffered_events] and [member Input.use_accumulated_input].
*/
//go:nosplit
func (self class) ProcessEvents() { //gd:DisplayServer.process_events
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
func (self class) ForceProcessAndDropEvents() { //gd:DisplayServer.force_process_and_drop_events
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
func (self class) SetNativeIcon(filename String.Readable) { //gd:DisplayServer.set_native_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(filename)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_set_native_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the window icon (usually displayed in the top-left corner) with an [Image]. To use icons in the operating system's native format, use [method set_native_icon] instead.
[b]Note:[/b] Requires support for [constant FEATURE_ICON].
*/
//go:nosplit
func (self class) SetIcon(image [1]gdclass.Image) { //gd:DisplayServer.set_icon
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
func (self class) CreateStatusIndicator(icon [1]gdclass.Texture2D, tooltip String.Readable, callback Callable.Function) int64 { //gd:DisplayServer.create_status_indicator
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) StatusIndicatorSetIcon(id int64, icon [1]gdclass.Texture2D) { //gd:DisplayServer.status_indicator_set_icon
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
func (self class) StatusIndicatorSetTooltip(id int64, tooltip String.Readable) { //gd:DisplayServer.status_indicator_set_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
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
func (self class) StatusIndicatorSetMenu(id int64, menu_rid RID.Any) { //gd:DisplayServer.status_indicator_set_menu
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
func (self class) StatusIndicatorSetCallback(id int64, callback Callable.Function) { //gd:DisplayServer.status_indicator_set_callback
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_status_indicator_set_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the rectangle for the given status indicator [param id] in screen coordinates. If the status indicator is not visible, returns an empty [Rect2].
[b]Note:[/b] This method is implemented on macOS and Windows.
*/
//go:nosplit
func (self class) StatusIndicatorGetRect(id int64) Rect2.PositionSize { //gd:DisplayServer.status_indicator_get_rect
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[Rect2.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_status_indicator_get_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the application status indicator.
*/
//go:nosplit
func (self class) DeleteStatusIndicator(id int64) { //gd:DisplayServer.delete_status_indicator
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
func (self class) TabletGetDriverCount() int64 { //gd:DisplayServer.tablet_get_driver_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
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
func (self class) TabletGetDriverName(idx int64) String.Readable { //gd:DisplayServer.tablet_get_driver_name
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tablet_get_driver_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns current active tablet driver name.
[b]Note:[/b] This method is implemented only on Windows.
*/
//go:nosplit
func (self class) TabletGetCurrentDriver() String.Readable { //gd:DisplayServer.tablet_get_current_driver
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tablet_get_current_driver, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
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
func (self class) TabletSetCurrentDriver(name String.Readable) { //gd:DisplayServer.tablet_set_current_driver
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DisplayServer.Bind_tablet_set_current_driver, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the window background can be made transparent. This method returns [code]false[/code] if [member ProjectSettings.display/window/per_pixel_transparency/allowed] is set to [code]false[/code], or if transparency is not supported by the renderer or OS compositor.
*/
//go:nosplit
func (self class) IsWindowTransparencyAvailable() bool { //gd:DisplayServer.is_window_transparency_available
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
func (self class) RegisterAdditionalOutput(obj [1]gd.Object) { //gd:DisplayServer.register_additional_output
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
func (self class) UnregisterAdditionalOutput(obj [1]gd.Object) { //gd:DisplayServer.unregister_additional_output
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
func (self class) HasAdditionalOutputs() bool { //gd:DisplayServer.has_additional_outputs
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

type Feature = gdclass.DisplayServerFeature //gd:DisplayServer.Feature

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
	/*Display server supports spawning dialogs for selecting files or directories using the operating system's native look-and-feel. See [method file_dialog_show]. [b]Windows, macOS, Linux (X11/Wayland), Android[/b]*/
	FeatureNativeDialogFile Feature = 25
	/*The display server supports all features of [constant FEATURE_NATIVE_DIALOG_FILE], with the added functionality of Options and native dialog file access to [code]res://[/code] and [code]user://[/code] paths. See [method file_dialog_show] and [method file_dialog_with_options_show]. [b]Windows, macOS, Linux (X11/Wayland)[/b]*/
	FeatureNativeDialogFileExtra Feature = 26
	/*The display server supports initiating window drag and resize operations on demand. See [method window_start_drag] and [method window_start_resize].*/
	FeatureWindowDrag Feature = 27
	/*Display server supports [constant WINDOW_FLAG_EXCLUDE_FROM_CAPTURE] window flag.*/
	FeatureScreenExcludeFromCapture Feature = 28
	/*Display server supports embedding a window from another process. [b]Windows, Linux (X11)[/b]*/
	FeatureWindowEmbedding Feature = 29
	/*Native file selection dialog supports MIME types as filters.*/
	FeatureNativeDialogFileMime Feature = 30
	/*Display server supports system emoji and symbol picker. [b]Windows, macOS[/b]*/
	FeatureEmojiAndSymbolPicker Feature = 31
)

type MouseModeValue = gdclass.DisplayServerMouseMode //gd:DisplayServer.MouseMode

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
	/*Max value of the [enum MouseMode].*/
	MouseModeMax MouseModeValue = 5
)

type ScreenOrientation = gdclass.DisplayServerScreenOrientation //gd:DisplayServer.ScreenOrientation

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

type VirtualKeyboardType = gdclass.DisplayServerVirtualKeyboardType //gd:DisplayServer.VirtualKeyboardType

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

type CursorShape = gdclass.DisplayServerCursorShape //gd:DisplayServer.CursorShape

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

type FileDialogMode = gdclass.DisplayServerFileDialogMode //gd:DisplayServer.FileDialogMode

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

type WindowMode = gdclass.DisplayServerWindowMode //gd:DisplayServer.WindowMode

const (
	/*Windowed mode, i.e. [Window] doesn't occupy the whole screen (unless set to the size of the screen).*/
	WindowModeWindowed WindowMode = 0
	/*Minimized window mode, i.e. [Window] is not visible and available on window manager's window list. Normally happens when the minimize button is pressed.*/
	WindowModeMinimized WindowMode = 1
	/*Maximized window mode, i.e. [Window] will occupy whole screen area except task bar and still display its borders. Normally happens when the maximize button is pressed.*/
	WindowModeMaximized WindowMode = 2
	/*Full screen mode with full multi-window support.
	  Full screen window covers the entire display area of a screen and has no decorations. The display's video mode is not changed.
	  [b]On Android:[/b] This enables immersive mode.
	  [b]On Windows:[/b] Multi-window full-screen mode has a 1px border of the [member ProjectSettings.rendering/environment/defaults/default_clear_color] color.
	  [b]On macOS:[/b] A new desktop is used to display the running project.
	  [b]Note:[/b] Regardless of the platform, enabling full screen will change the window size to match the monitor's size. Therefore, make sure your project supports [url=$DOCS_URL/tutorials/rendering/multiple_resolutions.html]multiple resolutions[/url] when enabling full screen mode.*/
	WindowModeFullscreen WindowMode = 3
	/*A single window full screen mode. This mode has less overhead, but only one window can be open on a given screen at a time (opening a child window or application switching will trigger a full screen transition).
	  Full screen window covers the entire display area of a screen and has no border or decorations. The display's video mode is not changed.
	  [b]On Android:[/b] This enables immersive mode.
	  [b]On Windows:[/b] Depending on video driver, full screen transition might cause screens to go black for a moment.
	  [b]On macOS:[/b] A new desktop is used to display the running project. Exclusive full screen mode prevents Dock and Menu from showing up when the mouse pointer is hovering the edge of the screen.
	  [b]On Linux (X11):[/b] Exclusive full screen mode bypasses compositor.
	  [b]On Linux (Wayland):[/b] Equivalent to [constant WINDOW_MODE_FULLSCREEN].
	  [b]Note:[/b] Regardless of the platform, enabling full screen will change the window size to match the monitor's size. Therefore, make sure your project supports [url=$DOCS_URL/tutorials/rendering/multiple_resolutions.html]multiple resolutions[/url] when enabling full screen mode.*/
	WindowModeExclusiveFullscreen WindowMode = 4
)

type WindowFlags = gdclass.DisplayServerWindowFlags //gd:DisplayServer.WindowFlags

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
	/*Window style is overridden, forcing sharp corners.
	  [b]Note:[/b] This flag is implemented only on Windows (11).*/
	WindowFlagSharpCorners WindowFlags = 8
	/*Windows is excluded from screenshots taken by [method screen_get_image], [method screen_get_image_rect], and [method screen_get_pixel].
	  [b]Note:[/b] This flag is implemented on macOS and Windows.
	  [b]Note:[/b] Setting this flag will [b]NOT[/b] prevent other apps from capturing an image, it should not be used as a security measure.*/
	WindowFlagExcludeFromCapture WindowFlags = 9
	/*Max value of the [enum WindowFlags].*/
	WindowFlagMax WindowFlags = 10
)

type WindowEvent = gdclass.DisplayServerWindowEvent //gd:DisplayServer.WindowEvent

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

type WindowResizeEdge = gdclass.DisplayServerWindowResizeEdge //gd:DisplayServer.WindowResizeEdge

const (
	/*Top-left edge of a window.*/
	WindowEdgeTopLeft WindowResizeEdge = 0
	/*Top edge of a window.*/
	WindowEdgeTop WindowResizeEdge = 1
	/*Top-right edge of a window.*/
	WindowEdgeTopRight WindowResizeEdge = 2
	/*Left edge of a window.*/
	WindowEdgeLeft WindowResizeEdge = 3
	/*Right edge of a window.*/
	WindowEdgeRight WindowResizeEdge = 4
	/*Bottom-left edge of a window.*/
	WindowEdgeBottomLeft WindowResizeEdge = 5
	/*Bottom edge of a window.*/
	WindowEdgeBottom WindowResizeEdge = 6
	/*Bottom-right edge of a window.*/
	WindowEdgeBottomRight WindowResizeEdge = 7
	/*Represents the size of the [enum WindowResizeEdge] enum.*/
	WindowEdgeMax WindowResizeEdge = 8
)

type VSyncMode = gdclass.DisplayServerVSyncMode //gd:DisplayServer.VSyncMode

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

type HandleType = gdclass.DisplayServerHandleType //gd:DisplayServer.HandleType

const (
	/*Display handle:
	  - Linux (X11): [code]X11::Display*[/code] for the display.
	  - Linux (Wayland): [code]wl_display[/code] for the display.
	  - Android: [code]EGLDisplay[/code] for the display.*/
	DisplayHandle HandleType = 0
	/*Window handle:
	  - Windows: [code]HWND[/code] for the window.
	  - Linux (X11): [code]X11::Window*[/code] for the window.
	  - Linux (Wayland): [code]wl_surface[/code] for the window.
	  - macOS: [code]NSWindow*[/code] for the window.
	  - iOS: [code]UIViewController*[/code] for the view controller.
	  - Android: [code]jObject[/code] for the activity.*/
	WindowHandle HandleType = 1
	/*Window view:
	  - Windows: [code]HDC[/code] for the window (only with the Compatibility renderer).
	  - macOS: [code]NSView*[/code] for the window main view.
	  - iOS: [code]UIView*[/code] for the window main view.*/
	WindowView HandleType = 2
	/*OpenGL context (only with the Compatibility renderer):
	  - Windows: [code]HGLRC[/code] for the window (native GL), or [code]EGLContext[/code] for the window (ANGLE).
	  - Linux (X11): [code]GLXContext*[/code] for the window.
	  - Linux (Wayland): [code]EGLContext[/code] for the window.
	  - macOS: [code]NSOpenGLContext*[/code] for the window (native GL), or [code]EGLContext[/code] for the window (ANGLE).
	  - Android: [code]EGLContext[/code] for the window.*/
	OpenglContext HandleType = 3
	/*- Windows: [code]EGLDisplay[/code] for the window (ANGLE).
	  - macOS: [code]EGLDisplay[/code] for the window (ANGLE).
	  - Linux (Wayland): [code]EGLDisplay[/code] for the window.*/
	EglDisplay HandleType = 4
	/*- Windows: [code]EGLConfig[/code] for the window (ANGLE).
	  - macOS: [code]EGLConfig[/code] for the window (ANGLE).
	  - Linux (Wayland): [code]EGLConfig[/code] for the window.*/
	EglConfig HandleType = 5
)

type TTSUtteranceEvent = gdclass.DisplayServerTTSUtteranceEvent //gd:DisplayServer.TTSUtteranceEvent

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
	/*Back key.*/
	KeyBack Key = 4194376
	/*Forward key.*/
	KeyForward Key = 4194377
	/*Media stop key.*/
	KeyStop Key = 4194378
	/*Refresh key.*/
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
	/*Exclamation mark ([code]![/code]) key.*/
	KeyExclam Key = 33
	/*Double quotation mark ([code]"[/code]) key.*/
	KeyQuotedbl Key = 34
	/*Number sign or [i]hash[/i] ([code]#[/code]) key.*/
	KeyNumbersign Key = 35
	/*Dollar sign ([code]$[/code]) key.*/
	KeyDollar Key = 36
	/*Percent sign ([code]%[/code]) key.*/
	KeyPercent Key = 37
	/*Ampersand ([code]&[/code]) key.*/
	KeyAmpersand Key = 38
	/*Apostrophe ([code]'[/code]) key.*/
	KeyApostrophe Key = 39
	/*Left parenthesis ([code]([/code]) key.*/
	KeyParenleft Key = 40
	/*Right parenthesis ([code])[/code]) key.*/
	KeyParenright Key = 41
	/*Asterisk ([code]*[/code]) key.*/
	KeyAsterisk Key = 42
	/*Plus ([code]+[/code]) key.*/
	KeyPlus Key = 43
	/*Comma ([code],[/code]) key.*/
	KeyComma Key = 44
	/*Minus ([code]-[/code]) key.*/
	KeyMinus Key = 45
	/*Period ([code].[/code]) key.*/
	KeyPeriod Key = 46
	/*Slash ([code]/[/code]) key.*/
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
	/*Colon ([code]:[/code]) key.*/
	KeyColon Key = 58
	/*Semicolon ([code];[/code]) key.*/
	KeySemicolon Key = 59
	/*Less-than sign ([code]<[/code]) key.*/
	KeyLess Key = 60
	/*Equal sign ([code]=[/code]) key.*/
	KeyEqual Key = 61
	/*Greater-than sign ([code]>[/code]) key.*/
	KeyGreater Key = 62
	/*Question mark ([code]?[/code]) key.*/
	KeyQuestion Key = 63
	/*At sign ([code]@[/code]) key.*/
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
	/*Left bracket ([code][lb][/code]) key.*/
	KeyBracketleft Key = 91
	/*Backslash ([code]\[/code]) key.*/
	KeyBackslash Key = 92
	/*Right bracket ([code][rb][/code]) key.*/
	KeyBracketright Key = 93
	/*Caret ([code]^[/code]) key.*/
	KeyAsciicircum Key = 94
	/*Underscore ([code]_[/code]) key.*/
	KeyUnderscore Key = 95
	/*Backtick ([code]`[/code]) key.*/
	KeyQuoteleft Key = 96
	/*Left brace ([code]{[/code]) key.*/
	KeyBraceleft Key = 123
	/*Vertical bar or [i]pipe[/i] ([code]|[/code]) key.*/
	KeyBar Key = 124
	/*Right brace ([code]}[/code]) key.*/
	KeyBraceright Key = 125
	/*Tilde ([code]~[/code]) key.*/
	KeyAsciitilde Key = 126
	/*Yen symbol ([code]¥[/code]) key.*/
	KeyYen Key = 165
	/*Section sign ([code]§[/code]) key.*/
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

type FileDialogOption struct {
	Name    string   `gd:"name"`
	Values  []string `gd:"values"`
	Default int      `gd:"default"`
}
type TextToSpeechVoice struct {
	Name     string `gd:"name"`
	ID       string `gd:"id"`
	Language string `gd:"language"`
}
