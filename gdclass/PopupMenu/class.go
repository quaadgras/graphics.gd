package PopupMenu

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Popup"
import "grow.graphics/gd/gdclass/Window"
import "grow.graphics/gd/gdclass/Viewport"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[PopupMenu] is a modal window used to display a list of options. Useful for toolbars and context menus.
The size of a [PopupMenu] can be limited by using [member Window.max_size]. If the height of the list of items is larger than the maximum height of the [PopupMenu], a [ScrollContainer] within the popup will allow the user to scroll the contents. If no maximum size is set, or if it is set to [code]0[/code], the [PopupMenu] height will be limited by its parent rect.
All [code]set_*[/code] methods allow negative item indices, i.e. [code]-1[/code] to access the last item, [code]-2[/code] to select the second-to-last item, and so on.
[b]Incremental search:[/b] Like [ItemList] and [Tree], [PopupMenu] supports searching within the list while the control is focused. Press a key that matches the first letter of an item's name to select the first item starting with the given letter. After that point, there are two ways to perform incremental search: 1) Press the same key again before the timeout duration to select the next item starting with the same letter. 2) Press letter keys that match the rest of the word before the timeout duration to match to select the item in question directly. Both of these actions will be reset to the beginning of the list if the timeout duration has passed since the last keystroke was registered. You can adjust the timeout duration by changing [member ProjectSettings.gui/timers/incremental_search_max_interval_msec].
[b]Note:[/b] The ID values used for items are limited to 32 bits, not full 64 bits of [int]. This has a range of [code]-2^32[/code] to [code]2^32 - 1[/code], i.e. [code]-2147483648[/code] to [code]2147483647[/code].
*/
type Instance [1]classdb.PopupMenu

/*
Checks the provided [param event] against the [PopupMenu]'s shortcuts and accelerators, and activates the first item with matching events. If [param for_global_only] is [code]true[/code], only shortcuts and accelerators with [code]global[/code] set to [code]true[/code] will be called.
Returns [code]true[/code] if an item was successfully activated.
[b]Note:[/b] Certain [Control]s, such as [MenuButton], will call this method automatically.
*/
func (self Instance) ActivateItemByEvent(event gdclass.InputEvent) bool {
	return bool(class(self).ActivateItemByEvent(event, false))
}

/*
Returns [code]true[/code] if the system native menu is supported and currently used by this [PopupMenu].
*/
func (self Instance) IsNativeMenu() bool {
	return bool(class(self).IsNativeMenu())
}

/*
Adds a new item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] The provided [param id] is used only in [signal id_pressed] and [signal id_focused] signals. It's not related to the [code]index[/code] arguments in e.g. [method set_item_checked].
*/
func (self Instance) AddItem(label string) {
	class(self).AddItem(gd.NewString(label), gd.Int(-1), 0)
}

/*
Adds a new item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
*/
func (self Instance) AddIconItem(texture gdclass.Texture2D, label string) {
	class(self).AddIconItem(texture, gd.NewString(label), gd.Int(-1), 0)
}

/*
Adds a new checkable item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddCheckItem(label string) {
	class(self).AddCheckItem(gd.NewString(label), gd.Int(-1), 0)
}

/*
Adds a new checkable item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddIconCheckItem(texture gdclass.Texture2D, label string) {
	class(self).AddIconCheckItem(texture, gd.NewString(label), gd.Int(-1), 0)
}

/*
Adds a new radio check button with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddRadioCheckItem(label string) {
	class(self).AddRadioCheckItem(gd.NewString(label), gd.Int(-1), 0)
}

/*
Same as [method add_icon_check_item], but uses a radio check button.
*/
func (self Instance) AddIconRadioCheckItem(texture gdclass.Texture2D, label string) {
	class(self).AddIconRadioCheckItem(texture, gd.NewString(label), gd.Int(-1), 0)
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
func (self Instance) AddMultistateItem(label string, max_states int) {
	class(self).AddMultistateItem(gd.NewString(label), gd.Int(max_states), gd.Int(0), gd.Int(-1), 0)
}

/*
Adds a [Shortcut].
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
func (self Instance) AddShortcut(shortcut gdclass.Shortcut) {
	class(self).AddShortcut(shortcut, gd.Int(-1), false, false)
}

/*
Adds a new item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
func (self Instance) AddIconShortcut(texture gdclass.Texture2D, shortcut gdclass.Shortcut) {
	class(self).AddIconShortcut(texture, shortcut, gd.Int(-1), false, false)
}

/*
Adds a new checkable item and assigns the specified [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddCheckShortcut(shortcut gdclass.Shortcut) {
	class(self).AddCheckShortcut(shortcut, gd.Int(-1), false)
}

/*
Adds a new checkable item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddIconCheckShortcut(texture gdclass.Texture2D, shortcut gdclass.Shortcut) {
	class(self).AddIconCheckShortcut(texture, shortcut, gd.Int(-1), false)
}

/*
Adds a new radio check button and assigns a [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
func (self Instance) AddRadioCheckShortcut(shortcut gdclass.Shortcut) {
	class(self).AddRadioCheckShortcut(shortcut, gd.Int(-1), false)
}

/*
Same as [method add_icon_check_shortcut], but uses a radio check button.
*/
func (self Instance) AddIconRadioCheckShortcut(texture gdclass.Texture2D, shortcut gdclass.Shortcut) {
	class(self).AddIconRadioCheckShortcut(texture, shortcut, gd.Int(-1), false)
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. The [param submenu] argument must be the name of an existing [PopupMenu] that has been added as a child to this node. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
func (self Instance) AddSubmenuItem(label string, submenu string) {
	class(self).AddSubmenuItem(gd.NewString(label), gd.NewString(submenu), gd.Int(-1))
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
[param submenu] must be either child of this [PopupMenu] or has no parent node (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
func (self Instance) AddSubmenuNodeItem(label string, submenu gdclass.PopupMenu) {
	class(self).AddSubmenuNodeItem(gd.NewString(label), submenu, gd.Int(-1))
}

/*
Sets the text of the item at the given [param index].
*/
func (self Instance) SetItemText(index int, text string) {
	class(self).SetItemText(gd.Int(index), gd.NewString(text))
}

/*
Sets item's text base writing direction.
*/
func (self Instance) SetItemTextDirection(index int, direction classdb.ControlTextDirection) {
	class(self).SetItemTextDirection(gd.Int(index), direction)
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Instance) SetItemLanguage(index int, language string) {
	class(self).SetItemLanguage(gd.Int(index), gd.NewString(language))
}

/*
Replaces the [Texture2D] icon of the item at the given [param index].
*/
func (self Instance) SetItemIcon(index int, icon gdclass.Texture2D) {
	class(self).SetItemIcon(gd.Int(index), icon)
}

/*
Sets the maximum allowed width of the icon for the item at the given [param index]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
func (self Instance) SetItemIconMaxWidth(index int, width int) {
	class(self).SetItemIconMaxWidth(gd.Int(index), gd.Int(width))
}

/*
Sets a modulating [Color] of the item's icon at the given [param index].
*/
func (self Instance) SetItemIconModulate(index int, modulate gd.Color) {
	class(self).SetItemIconModulate(gd.Int(index), modulate)
}

/*
Sets the checkstate status of the item at the given [param index].
*/
func (self Instance) SetItemChecked(index int, checked bool) {
	class(self).SetItemChecked(gd.Int(index), checked)
}

/*
Sets the [param id] of the item at the given [param index].
The [param id] is used in [signal id_pressed] and [signal id_focused] signals.
*/
func (self Instance) SetItemId(index int, id int) {
	class(self).SetItemId(gd.Int(index), gd.Int(id))
}

/*
Sets the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. [param accel] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
*/
func (self Instance) SetItemAccelerator(index int, accel gdconst.Key) {
	class(self).SetItemAccelerator(gd.Int(index), accel)
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_metadata], which provides a simple way of assigning context data to items.
*/
func (self Instance) SetItemMetadata(index int, metadata gd.Variant) {
	class(self).SetItemMetadata(gd.Int(index), metadata)
}

/*
Enables/disables the item at the given [param index]. When it is disabled, it can't be selected and its action can't be invoked.
*/
func (self Instance) SetItemDisabled(index int, disabled bool) {
	class(self).SetItemDisabled(gd.Int(index), disabled)
}

/*
Sets the submenu of the item at the given [param index]. The submenu is the name of a child [PopupMenu] node that would be shown when the item is clicked.
*/
func (self Instance) SetItemSubmenu(index int, submenu string) {
	class(self).SetItemSubmenu(gd.Int(index), gd.NewString(submenu))
}

/*
Sets the submenu of the item at the given [param index]. The submenu is a [PopupMenu] node that would be shown when the item is clicked. It must either be a child of this [PopupMenu] or has no parent (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
*/
func (self Instance) SetItemSubmenuNode(index int, submenu gdclass.PopupMenu) {
	class(self).SetItemSubmenuNode(gd.Int(index), submenu)
}

/*
Mark the item at the given [param index] as a separator, which means that it would be displayed as a line. If [code]false[/code], sets the type of the item to plain text.
*/
func (self Instance) SetItemAsSeparator(index int, enable bool) {
	class(self).SetItemAsSeparator(gd.Int(index), enable)
}

/*
Sets whether the item at the given [param index] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
func (self Instance) SetItemAsCheckable(index int, enable bool) {
	class(self).SetItemAsCheckable(gd.Int(index), enable)
}

/*
Sets the type of the item at the given [param index] to radio button. If [code]false[/code], sets the type of the item to plain text.
*/
func (self Instance) SetItemAsRadioCheckable(index int, enable bool) {
	class(self).SetItemAsRadioCheckable(gd.Int(index), enable)
}

/*
Sets the [String] tooltip of the item at the given [param index].
*/
func (self Instance) SetItemTooltip(index int, tooltip string) {
	class(self).SetItemTooltip(gd.Int(index), gd.NewString(tooltip))
}

/*
Sets a [Shortcut] for the item at the given [param index].
*/
func (self Instance) SetItemShortcut(index int, shortcut gdclass.Shortcut) {
	class(self).SetItemShortcut(gd.Int(index), shortcut, false)
}

/*
Sets the horizontal offset of the item at the given [param index].
*/
func (self Instance) SetItemIndent(index int, indent int) {
	class(self).SetItemIndent(gd.Int(index), gd.Int(indent))
}

/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
*/
func (self Instance) SetItemMultistate(index int, state int) {
	class(self).SetItemMultistate(gd.Int(index), gd.Int(state))
}

/*
Sets the max states of a multistate item. See [method add_multistate_item] for details.
*/
func (self Instance) SetItemMultistateMax(index int, max_states int) {
	class(self).SetItemMultistateMax(gd.Int(index), gd.Int(max_states))
}

/*
Disables the [Shortcut] of the item at the given [param index].
*/
func (self Instance) SetItemShortcutDisabled(index int, disabled bool) {
	class(self).SetItemShortcutDisabled(gd.Int(index), disabled)
}

/*
Toggles the check state of the item at the given [param index].
*/
func (self Instance) ToggleItemChecked(index int) {
	class(self).ToggleItemChecked(gd.Int(index))
}

/*
Cycle to the next state of a multistate item. See [method add_multistate_item] for details.
*/
func (self Instance) ToggleItemMultistate(index int) {
	class(self).ToggleItemMultistate(gd.Int(index))
}

/*
Returns the text of the item at the given [param index].
*/
func (self Instance) GetItemText(index int) string {
	return string(class(self).GetItemText(gd.Int(index)).String())
}

/*
Returns item's text base writing direction.
*/
func (self Instance) GetItemTextDirection(index int) classdb.ControlTextDirection {
	return classdb.ControlTextDirection(class(self).GetItemTextDirection(gd.Int(index)))
}

/*
Returns item's text language code.
*/
func (self Instance) GetItemLanguage(index int) string {
	return string(class(self).GetItemLanguage(gd.Int(index)).String())
}

/*
Returns the icon of the item at the given [param index].
*/
func (self Instance) GetItemIcon(index int) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetItemIcon(gd.Int(index)))
}

/*
Returns the maximum allowed width of the icon for the item at the given [param index].
*/
func (self Instance) GetItemIconMaxWidth(index int) int {
	return int(int(class(self).GetItemIconMaxWidth(gd.Int(index))))
}

/*
Returns a [Color] modulating the item's icon at the given [param index].
*/
func (self Instance) GetItemIconModulate(index int) gd.Color {
	return gd.Color(class(self).GetItemIconModulate(gd.Int(index)))
}

/*
Returns [code]true[/code] if the item at the given [param index] is checked.
*/
func (self Instance) IsItemChecked(index int) bool {
	return bool(class(self).IsItemChecked(gd.Int(index)))
}

/*
Returns the ID of the item at the given [param index]. [code]id[/code] can be manually assigned, while index can not.
*/
func (self Instance) GetItemId(index int) int {
	return int(int(class(self).GetItemId(gd.Int(index))))
}

/*
Returns the index of the item containing the specified [param id]. Index is automatically assigned to each item by the engine and can not be set manually.
*/
func (self Instance) GetItemIndex(id int) int {
	return int(int(class(self).GetItemIndex(gd.Int(id))))
}

/*
Returns the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The return value is an integer which is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]). If no accelerator is defined for the specified [param index], [method get_item_accelerator] returns [code]0[/code] (corresponding to [constant @GlobalScope.KEY_NONE]).
*/
func (self Instance) GetItemAccelerator(index int) gdconst.Key {
	return gdconst.Key(class(self).GetItemAccelerator(gd.Int(index)))
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_metadata], which provides a simple way of assigning context data to items.
*/
func (self Instance) GetItemMetadata(index int) gd.Variant {
	return gd.Variant(class(self).GetItemMetadata(gd.Int(index)))
}

/*
Returns [code]true[/code] if the item at the given [param index] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method set_item_disabled] for more info on how to disable an item.
*/
func (self Instance) IsItemDisabled(index int) bool {
	return bool(class(self).IsItemDisabled(gd.Int(index)))
}

/*
Returns the submenu name of the item at the given [param index]. See [method add_submenu_item] for more info on how to add a submenu.
*/
func (self Instance) GetItemSubmenu(index int) string {
	return string(class(self).GetItemSubmenu(gd.Int(index)).String())
}

/*
Returns the submenu of the item at the given [param index], or [code]null[/code] if no submenu was added. See [method add_submenu_node_item] for more info on how to add a submenu.
*/
func (self Instance) GetItemSubmenuNode(index int) gdclass.PopupMenu {
	return gdclass.PopupMenu(class(self).GetItemSubmenuNode(gd.Int(index)))
}

/*
Returns [code]true[/code] if the item is a separator. If it is, it will be displayed as a line. See [method add_separator] for more info on how to add a separator.
*/
func (self Instance) IsItemSeparator(index int) bool {
	return bool(class(self).IsItemSeparator(gd.Int(index)))
}

/*
Returns [code]true[/code] if the item at the given [param index] is checkable in some way, i.e. if it has a checkbox or radio button.
[b]Note:[/b] Checkable items just display a checkmark or radio button, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
func (self Instance) IsItemCheckable(index int) bool {
	return bool(class(self).IsItemCheckable(gd.Int(index)))
}

/*
Returns [code]true[/code] if the item at the given [param index] has radio button-style checkability.
[b]Note:[/b] This is purely cosmetic; you must add the logic for checking/unchecking items in radio groups.
*/
func (self Instance) IsItemRadioCheckable(index int) bool {
	return bool(class(self).IsItemRadioCheckable(gd.Int(index)))
}

/*
Returns [code]true[/code] if the specified item's shortcut is disabled.
*/
func (self Instance) IsItemShortcutDisabled(index int) bool {
	return bool(class(self).IsItemShortcutDisabled(gd.Int(index)))
}

/*
Returns the tooltip associated with the item at the given [param index].
*/
func (self Instance) GetItemTooltip(index int) string {
	return string(class(self).GetItemTooltip(gd.Int(index)).String())
}

/*
Returns the [Shortcut] associated with the item at the given [param index].
*/
func (self Instance) GetItemShortcut(index int) gdclass.Shortcut {
	return gdclass.Shortcut(class(self).GetItemShortcut(gd.Int(index)))
}

/*
Returns the horizontal offset of the item at the given [param index].
*/
func (self Instance) GetItemIndent(index int) int {
	return int(int(class(self).GetItemIndent(gd.Int(index))))
}

/*
Returns the max states of the item at the given [param index].
*/
func (self Instance) GetItemMultistateMax(index int) int {
	return int(int(class(self).GetItemMultistateMax(gd.Int(index))))
}

/*
Returns the state of the item at the given [param index].
*/
func (self Instance) GetItemMultistate(index int) int {
	return int(int(class(self).GetItemMultistate(gd.Int(index))))
}

/*
Sets the currently focused item as the given [param index].
Passing [code]-1[/code] as the index makes so that no item is focused.
*/
func (self Instance) SetFocusedItem(index int) {
	class(self).SetFocusedItem(gd.Int(index))
}

/*
Returns the index of the currently focused item. Returns [code]-1[/code] if no item is focused.
*/
func (self Instance) GetFocusedItem() int {
	return int(int(class(self).GetFocusedItem()))
}

/*
Moves the scroll view to make the item at the given [param index] visible.
*/
func (self Instance) ScrollToItem(index int) {
	class(self).ScrollToItem(gd.Int(index))
}

/*
Removes the item at the given [param index] from the menu.
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
*/
func (self Instance) RemoveItem(index int) {
	class(self).RemoveItem(gd.Int(index))
}

/*
Adds a separator between items. Separators also occupy an index, which you can set by using the [param id] parameter.
A [param label] can optionally be provided, which will appear at the center of the separator.
*/
func (self Instance) AddSeparator() {
	class(self).AddSeparator(gd.NewString(""), gd.Int(-1))
}

/*
Removes all items from the [PopupMenu]. If [param free_submenus] is [code]true[/code], the submenu nodes are automatically freed.
*/
func (self Instance) Clear() {
	class(self).Clear(false)
}

/*
Returns [code]true[/code] if the menu is bound to the special system menu.
*/
func (self Instance) IsSystemMenu() bool {
	return bool(class(self).IsSystemMenu())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PopupMenu

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PopupMenu"))
	return Instance{classdb.PopupMenu(object)}
}

func (self Instance) HideOnItemSelection() bool {
	return bool(class(self).IsHideOnItemSelection())
}

func (self Instance) SetHideOnItemSelection(value bool) {
	class(self).SetHideOnItemSelection(value)
}

func (self Instance) HideOnCheckableItemSelection() bool {
	return bool(class(self).IsHideOnCheckableItemSelection())
}

func (self Instance) SetHideOnCheckableItemSelection(value bool) {
	class(self).SetHideOnCheckableItemSelection(value)
}

func (self Instance) HideOnStateItemSelection() bool {
	return bool(class(self).IsHideOnStateItemSelection())
}

func (self Instance) SetHideOnStateItemSelection(value bool) {
	class(self).SetHideOnStateItemSelection(value)
}

func (self Instance) SubmenuPopupDelay() float64 {
	return float64(float64(class(self).GetSubmenuPopupDelay()))
}

func (self Instance) SetSubmenuPopupDelay(value float64) {
	class(self).SetSubmenuPopupDelay(gd.Float(value))
}

func (self Instance) AllowSearch() bool {
	return bool(class(self).GetAllowSearch())
}

func (self Instance) SetAllowSearch(value bool) {
	class(self).SetAllowSearch(value)
}

func (self Instance) SystemMenuId() classdb.NativeMenuSystemMenus {
	return classdb.NativeMenuSystemMenus(class(self).GetSystemMenu())
}

func (self Instance) SetSystemMenuId(value classdb.NativeMenuSystemMenus) {
	class(self).SetSystemMenu(value)
}

func (self Instance) PreferNativeMenu() bool {
	return bool(class(self).IsPreferNativeMenu())
}

func (self Instance) SetPreferNativeMenu(value bool) {
	class(self).SetPreferNativeMenu(value)
}

func (self Instance) ItemCount() int {
	return int(int(class(self).GetItemCount()))
}

func (self Instance) SetItemCount(value int) {
	class(self).SetItemCount(gd.Int(value))
}

/*
Checks the provided [param event] against the [PopupMenu]'s shortcuts and accelerators, and activates the first item with matching events. If [param for_global_only] is [code]true[/code], only shortcuts and accelerators with [code]global[/code] set to [code]true[/code] will be called.
Returns [code]true[/code] if an item was successfully activated.
[b]Note:[/b] Certain [Control]s, such as [MenuButton], will call this method automatically.
*/
//go:nosplit
func (self class) ActivateItemByEvent(event gdclass.InputEvent, for_global_only bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	callframe.Arg(frame, for_global_only)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_activate_item_by_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreferNativeMenu(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_prefer_native_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsPreferNativeMenu() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_prefer_native_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the system native menu is supported and currently used by this [PopupMenu].
*/
//go:nosplit
func (self class) IsNativeMenu() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_native_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AddItem(label gd.String, id gd.Int, accel gdconst.Key) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
*/
//go:nosplit
func (self class) AddIconItem(texture gdclass.Texture2D, label gd.String, id gd.Int, accel gdconst.Key) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new checkable item with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddCheckItem(label gd.String, id gd.Int, accel gdconst.Key) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new checkable item with text [param label] and icon [param texture].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddIconCheckItem(texture gdclass.Texture2D, label gd.String, id gd.Int, accel gdconst.Key) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new radio check button with text [param label].
An [param id] can optionally be provided, as well as an accelerator ([param accel]). If no [param id] is provided, one will be created from the index. If no [param accel] is provided, then the default value of 0 (corresponding to [constant @GlobalScope.KEY_NONE]) will be assigned to the item (which means it won't have any accelerator). See [method get_item_accelerator] for more info on accelerators.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddRadioCheckItem(label gd.String, id gd.Int, accel gdconst.Key) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Same as [method add_icon_check_item], but uses a radio check button.
*/
//go:nosplit
func (self class) AddIconRadioCheckItem(texture gdclass.Texture2D, label gd.String, id gd.Int, accel gdconst.Key) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_radio_check_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AddMultistateItem(label gd.String, max_states gd.Int, default_state gd.Int, id gd.Int, accel gdconst.Key) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, max_states)
	callframe.Arg(frame, default_state)
	callframe.Arg(frame, id)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_multistate_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a [Shortcut].
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
//go:nosplit
func (self class) AddShortcut(shortcut gdclass.Shortcut, id gd.Int, global bool, allow_echo bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	callframe.Arg(frame, allow_echo)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
If [param allow_echo] is [code]true[/code], the shortcut can be activated with echo events.
*/
//go:nosplit
func (self class) AddIconShortcut(texture gdclass.Texture2D, shortcut gdclass.Shortcut, id gd.Int, global bool, allow_echo bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	callframe.Arg(frame, allow_echo)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new checkable item and assigns the specified [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddCheckShortcut(shortcut gdclass.Shortcut, id gd.Int, global bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new checkable item and assigns the specified [Shortcut] and icon [param texture] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddIconCheckShortcut(texture gdclass.Texture2D, shortcut gdclass.Shortcut, id gd.Int, global bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new radio check button and assigns a [Shortcut] to it. Sets the label of the checkbox to the [Shortcut]'s name.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually. See [method set_item_checked] for more info on how to control it.
*/
//go:nosplit
func (self class) AddRadioCheckShortcut(shortcut gdclass.Shortcut, id gd.Int, global bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_radio_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Same as [method add_icon_check_shortcut], but uses a radio check button.
*/
//go:nosplit
func (self class) AddIconRadioCheckShortcut(texture gdclass.Texture2D, shortcut gdclass.Shortcut, id gd.Int, global bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_icon_radio_check_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. The [param submenu] argument must be the name of an existing [PopupMenu] that has been added as a child to this node. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
//go:nosplit
func (self class) AddSubmenuItem(label gd.String, submenu gd.String, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(submenu))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_submenu_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds an item that will act as a submenu of the parent [PopupMenu] node when clicked. This submenu will be shown when the item is clicked, hovered for long enough, or activated using the [code]ui_select[/code] or [code]ui_right[/code] input actions.
[param submenu] must be either child of this [PopupMenu] or has no parent node (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
An [param id] can optionally be provided. If no [param id] is provided, one will be created from the index.
*/
//go:nosplit
func (self class) AddSubmenuNodeItem(label gd.String, submenu gdclass.PopupMenu, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, pointers.Get(submenu[0])[0])
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_submenu_node_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the text of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemText(index gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetItemTextDirection(index gd.Int, direction classdb.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetItemLanguage(index gd.Int, language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Replaces the [Texture2D] icon of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemIcon(index gd.Int, icon gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the maximum allowed width of the icon for the item at the given [param index]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetItemIconMaxWidth(index gd.Int, width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets a modulating [Color] of the item's icon at the given [param index].
*/
//go:nosplit
func (self class) SetItemIconModulate(index gd.Int, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the checkstate status of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemChecked(index gd.Int, checked bool) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the [param id] of the item at the given [param index].
The [param id] is used in [signal id_pressed] and [signal id_focused] signals.
*/
//go:nosplit
func (self class) SetItemId(index gd.Int, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. [param accel] is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]).
*/
//go:nosplit
func (self class) SetItemAccelerator(index gd.Int, accel gdconst.Key) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, accel)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the metadata of an item, which may be of any type. You can later get it with [method get_item_metadata], which provides a simple way of assigning context data to items.
*/
//go:nosplit
func (self class) SetItemMetadata(index gd.Int, metadata gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(metadata))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Enables/disables the item at the given [param index]. When it is disabled, it can't be selected and its action can't be invoked.
*/
//go:nosplit
func (self class) SetItemDisabled(index gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the submenu of the item at the given [param index]. The submenu is the name of a child [PopupMenu] node that would be shown when the item is clicked.
*/
//go:nosplit
func (self class) SetItemSubmenu(index gd.Int, submenu gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(submenu))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the submenu of the item at the given [param index]. The submenu is a [PopupMenu] node that would be shown when the item is clicked. It must either be a child of this [PopupMenu] or has no parent (in which case it will be automatically added as a child). If the [param submenu] popup has another parent, this method will fail.
*/
//go:nosplit
func (self class) SetItemSubmenuNode(index gd.Int, submenu gdclass.PopupMenu) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(submenu[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_submenu_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Mark the item at the given [param index] as a separator, which means that it would be displayed as a line. If [code]false[/code], sets the type of the item to plain text.
*/
//go:nosplit
func (self class) SetItemAsSeparator(index gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_as_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets whether the item at the given [param index] has a checkbox. If [code]false[/code], sets the type of the item to plain text.
[b]Note:[/b] Checkable items just display a checkmark, but don't have any built-in checking behavior and must be checked/unchecked manually.
*/
//go:nosplit
func (self class) SetItemAsCheckable(index gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_as_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the type of the item at the given [param index] to radio button. If [code]false[/code], sets the type of the item to plain text.
*/
//go:nosplit
func (self class) SetItemAsRadioCheckable(index gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_as_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the [String] tooltip of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemTooltip(index gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets a [Shortcut] for the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemShortcut(index gd.Int, shortcut gdclass.Shortcut, global bool) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	callframe.Arg(frame, global)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the horizontal offset of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemIndent(index gd.Int, indent gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, indent)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_indent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the state of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) SetItemMultistate(index gd.Int, state gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, state)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_multistate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the max states of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) SetItemMultistateMax(index gd.Int, max_states gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, max_states)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_multistate_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Disables the [Shortcut] of the item at the given [param index].
*/
//go:nosplit
func (self class) SetItemShortcutDisabled(index gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_shortcut_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Toggles the check state of the item at the given [param index].
*/
//go:nosplit
func (self class) ToggleItemChecked(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_toggle_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Cycle to the next state of a multistate item. See [method add_multistate_item] for details.
*/
//go:nosplit
func (self class) ToggleItemMultistate(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_toggle_item_multistate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the text of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemText(index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetItemTextDirection(index gd.Int) classdb.ControlTextDirection {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetItemLanguage(index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the icon of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIcon(index gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the maximum allowed width of the icon for the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIconMaxWidth(index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Color] modulating the item's icon at the given [param index].
*/
//go:nosplit
func (self class) GetItemIconModulate(index gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at the given [param index] is checked.
*/
//go:nosplit
func (self class) IsItemChecked(index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ID of the item at the given [param index]. [code]id[/code] can be manually assigned, while index can not.
*/
//go:nosplit
func (self class) GetItemId(index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the item containing the specified [param id]. Index is automatically assigned to each item by the engine and can not be set manually.
*/
//go:nosplit
func (self class) GetItemIndex(id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the accelerator of the item at the given [param index]. An accelerator is a keyboard shortcut that can be pressed to trigger the menu button even if it's not currently open. The return value is an integer which is generally a combination of [enum KeyModifierMask]s and [enum Key]s using bitwise OR such as [code]KEY_MASK_CTRL | KEY_A[/code] ([kbd]Ctrl + A[/kbd]). If no accelerator is defined for the specified [param index], [method get_item_accelerator] returns [code]0[/code] (corresponding to [constant @GlobalScope.KEY_NONE]).
*/
//go:nosplit
func (self class) GetItemAccelerator(index gd.Int) gdconst.Key {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gdconst.Key](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_accelerator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the metadata of the specified item, which might be of any type. You can set it with [method set_item_metadata], which provides a simple way of assigning context data to items.
*/
//go:nosplit
func (self class) GetItemMetadata(index gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item at the given [param index] is disabled. When it is disabled it can't be selected, or its action invoked.
See [method set_item_disabled] for more info on how to disable an item.
*/
//go:nosplit
func (self class) IsItemDisabled(index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the submenu name of the item at the given [param index]. See [method add_submenu_item] for more info on how to add a submenu.
*/
//go:nosplit
func (self class) GetItemSubmenu(index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_submenu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the submenu of the item at the given [param index], or [code]null[/code] if no submenu was added. See [method add_submenu_node_item] for more info on how to add a submenu.
*/
//go:nosplit
func (self class) GetItemSubmenuNode(index gd.Int) gdclass.PopupMenu {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_submenu_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.PopupMenu{classdb.PopupMenu(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the item is a separator. If it is, it will be displayed as a line. See [method add_separator] for more info on how to add a separator.
*/
//go:nosplit
func (self class) IsItemSeparator(index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_radio_checkable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified item's shortcut is disabled.
*/
//go:nosplit
func (self class) IsItemShortcutDisabled(index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_item_shortcut_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tooltip associated with the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemTooltip(index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [Shortcut] associated with the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemShortcut(index gd.Int) gdclass.Shortcut {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_shortcut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Shortcut{classdb.Shortcut(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the horizontal offset of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemIndent(index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_indent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the max states of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemMultistateMax(index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_multistate_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the state of the item at the given [param index].
*/
//go:nosplit
func (self class) GetItemMultistate(index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_multistate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the currently focused item as the given [param index].
Passing [code]-1[/code] as the index makes so that no item is focused.
*/
//go:nosplit
func (self class) SetFocusedItem(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_focused_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the index of the currently focused item. Returns [code]-1[/code] if no item is focused.
*/
//go:nosplit
func (self class) GetFocusedItem() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_focused_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetItemCount(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetItemCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves the scroll view to make the item at the given [param index] visible.
*/
//go:nosplit
func (self class) ScrollToItem(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_scroll_to_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the item at the given [param index] from the menu.
[b]Note:[/b] The indices of items after the removed item will be shifted by one.
*/
//go:nosplit
func (self class) RemoveItem(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a separator between items. Separators also occupy an index, which you can set by using the [param id] parameter.
A [param label] can optionally be provided, which will appear at the center of the separator.
*/
//go:nosplit
func (self class) AddSeparator(label gd.String, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_add_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all items from the [PopupMenu]. If [param free_submenus] is [code]true[/code], the submenu nodes are automatically freed.
*/
//go:nosplit
func (self class) Clear(free_submenus bool) {
	var frame = callframe.New()
	callframe.Arg(frame, free_submenus)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetHideOnItemSelection(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_hide_on_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsHideOnItemSelection() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_hide_on_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHideOnCheckableItemSelection(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_hide_on_checkable_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsHideOnCheckableItemSelection() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_hide_on_checkable_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHideOnStateItemSelection(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_hide_on_state_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsHideOnStateItemSelection() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_hide_on_state_item_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubmenuPopupDelay(seconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_submenu_popup_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubmenuPopupDelay() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_submenu_popup_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowSearch(allow bool) {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAllowSearch() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the menu is bound to the special system menu.
*/
//go:nosplit
func (self class) IsSystemMenu() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_is_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSystemMenu(system_menu_id classdb.NativeMenuSystemMenus) {
	var frame = callframe.New()
	callframe.Arg(frame, system_menu_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_set_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSystemMenu() classdb.NativeMenuSystemMenus {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.NativeMenuSystemMenus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PopupMenu.Bind_get_system_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnIdPressed(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("id_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnIdFocused(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("id_focused"), gd.NewCallable(cb), 0)
}

func (self Instance) OnIndexPressed(cb func(index int)) {
	self[0].AsObject().Connect(gd.NewStringName("index_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMenuChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("menu_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsPopupMenu() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPopupMenu() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPopup() Popup.Advanced      { return *((*Popup.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPopup() Popup.Instance   { return *((*Popup.Instance)(unsafe.Pointer(&self))) }
func (self class) AsWindow() Window.Advanced    { return *((*Window.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWindow() Window.Instance { return *((*Window.Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPopup(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPopup(), name)
	}
}
func init() { classdb.Register("PopupMenu", func(ptr gd.Object) any { return classdb.PopupMenu(ptr) }) }
