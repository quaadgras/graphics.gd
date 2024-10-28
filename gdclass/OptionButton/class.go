package OptionButton

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
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
var _ = discreet.Root

/*
[OptionButton] is a type of button that brings up a dropdown with selectable items when pressed. The item selected becomes the "current" item and is displayed as the button text.
See also [BaseButton] which contains common properties and methods associated with this node.
[b]Note:[/b] The ID values used for items are limited to 32 bits, not full 64 bits of [int]. This has a range of [code]-2^32[/code] to [code]2^32 - 1[/code], i.e. [code]-2147483648[/code] to [code]2147483647[/code].
[b]Note:[/b] The [member Button.text] and [member Button.icon] properties are set automatically based on the selected item. They shouldn't be changed manually.

*/
type Go [1]classdb.OptionButton

/*
Adds an item, with text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
func (self Go) AddItem(label string) {
	class(self).AddItem(gd.NewString(label), gd.Int(-1))
}

/*
Adds an item, with a [param texture] icon, text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
func (self Go) AddIconItem(texture gdclass.Texture2D, label string) {
	class(self).AddIconItem(texture, gd.NewString(label), gd.Int(-1))
}

/*
Sets the text of the item at index [param idx].
*/
func (self Go) SetItemText(idx int, text string) {
	class(self).SetItemText(gd.Int(idx), gd.NewString(text))
}

/*
Sets the icon of the item at index [param idx].
*/
func (self Go) SetItemIcon(idx int, texture gdclass.Texture2D) {
	class(self).SetItemIcon(gd.Int(idx), texture)
}

/*
Sets whether the item at index [param idx] is disabled.
Disabled items are drawn differently in the dropdown and are not selectable by the user. If the current selected item is set as disabled, it will remain selected.
*/
func (self Go) SetItemDisabled(idx int, disabled bool) {
	class(self).SetItemDisabled(gd.Int(idx), disabled)
}

/*
Sets the ID of the item at index [param idx].
*/
func (self Go) SetItemId(idx int, id int) {
	class(self).SetItemId(gd.Int(idx), gd.Int(id))
}

/*
Sets the metadata of an item. Metadata may be of any type and can be used to store extra information about an item, such as an external string ID.
*/
func (self Go) SetItemMetadata(idx int, metadata gd.Variant) {
	class(self).SetItemMetadata(gd.Int(idx), metadata)
}

/*
Sets the tooltip of the item at index [param idx].
*/
func (self Go) SetItemTooltip(idx int, tooltip string) {
	class(self).SetItemTooltip(gd.Int(idx), gd.NewString(tooltip))
}

/*
Returns the text of the item at index [param idx].
*/
func (self Go) GetItemText(idx int) string {
	return string(class(self).GetItemText(gd.Int(idx)).String())
}

/*
Returns the icon of the item at index [param idx].
*/
func (self Go) GetItemIcon(idx int) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetItemIcon(gd.Int(idx)))
}

/*
Returns the ID of the item at index [param idx].
*/
func (self Go) GetItemId(idx int) int {
	return int(int(class(self).GetItemId(gd.Int(idx))))
}

/*
Returns the index of the item with the given [param id].
*/
func (self Go) GetItemIndex(id int) int {
	return int(int(class(self).GetItemIndex(gd.Int(id))))
}

/*
Retrieves the metadata of an item. Metadata may be any type and can be used to store extra information about an item, such as an external string ID.
*/
func (self Go) GetItemMetadata(idx int) gd.Variant {
	return gd.Variant(class(self).GetItemMetadata(gd.Int(idx)))
}

/*
Returns the tooltip of the item at index [param idx].
*/
func (self Go) GetItemTooltip(idx int) string {
	return string(class(self).GetItemTooltip(gd.Int(idx)).String())
}

/*
Returns [code]true[/code] if the item at index [param idx] is disabled.
*/
func (self Go) IsItemDisabled(idx int) bool {
	return bool(class(self).IsItemDisabled(gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is marked as a separator.
*/
func (self Go) IsItemSeparator(idx int) bool {
	return bool(class(self).IsItemSeparator(gd.Int(idx)))
}

/*
Adds a separator to the list of items. Separators help to group items, and can optionally be given a [param text] header. A separator also gets an index assigned, and is appended at the end of the item list.
*/
func (self Go) AddSeparator() {
	class(self).AddSeparator(gd.NewString(""))
}

/*
Clears all the items in the [OptionButton].
*/
func (self Go) Clear() {
	class(self).Clear()
}

/*
Selects an item by index and makes it the current item. This will work even if the item is disabled.
Passing [code]-1[/code] as the index deselects any currently selected item.
*/
func (self Go) Select(idx int) {
	class(self).Select(gd.Int(idx))
}

/*
Returns the ID of the selected item, or [code]-1[/code] if no item is selected.
*/
func (self Go) GetSelectedId() int {
	return int(int(class(self).GetSelectedId()))
}

/*
Gets the metadata of the selected item. Metadata for items can be set using [method set_item_metadata].
*/
func (self Go) GetSelectedMetadata() gd.Variant {
	return gd.Variant(class(self).GetSelectedMetadata())
}

/*
Removes the item at index [param idx].
*/
func (self Go) RemoveItem(idx int) {
	class(self).RemoveItem(gd.Int(idx))
}

/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Go) GetPopup() gdclass.PopupMenu {
	return gdclass.PopupMenu(class(self).GetPopup())
}

/*
Adjusts popup position and sizing for the [OptionButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
func (self Go) ShowPopup() {
	class(self).ShowPopup()
}

/*
Returns [code]true[/code] if this button contains at least one item which is not disabled, or marked as a separator.
*/
func (self Go) HasSelectableItems() bool {
	return bool(class(self).HasSelectableItems())
}

/*
Returns the index of the first item which is not disabled, or marked as a separator. If [param from_last] is [code]true[/code], the items will be searched in reverse order.
Returns [code]-1[/code] if no item is found.
*/
func (self Go) GetSelectableItem() int {
	return int(int(class(self).GetSelectableItem(false)))
}

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
func (self Go) SetDisableShortcuts(disabled bool) {
	class(self).SetDisableShortcuts(disabled)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OptionButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OptionButton"))
	return Go{classdb.OptionButton(object)}
}

func (self Go) Selected() int {
		return int(int(class(self).GetSelected()))
}

func (self Go) FitToLongestItem() bool {
		return bool(class(self).IsFitToLongestItem())
}

func (self Go) SetFitToLongestItem(value bool) {
	class(self).SetFitToLongestItem(value)
}

func (self Go) AllowReselect() bool {
		return bool(class(self).GetAllowReselect())
}

func (self Go) SetAllowReselect(value bool) {
	class(self).SetAllowReselect(value)
}

func (self Go) ItemCount() int {
		return int(int(class(self).GetItemCount()))
}

func (self Go) SetItemCount(value int) {
	class(self).SetItemCount(gd.Int(value))
}

/*
Adds an item, with text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
//go:nosplit
func (self class) AddItem(label gd.String, id gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an item, with a [param texture] icon, text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
//go:nosplit
func (self class) AddIconItem(texture gdclass.Texture2D, label gd.String, id gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(texture[0])[0])
	callframe.Arg(frame, discreet.Get(label))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the text of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemText(idx gd.Int, text gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the icon of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemIcon(idx gd.Int, texture gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the item at index [param idx] is disabled.
Disabled items are drawn differently in the dropdown and are not selectable by the user. If the current selected item is set as disabled, it will remain selected.
*/
//go:nosplit
func (self class) SetItemDisabled(idx gd.Int, disabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the ID of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemId(idx gd.Int, id gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the metadata of an item. Metadata may be of any type and can be used to store extra information about an item, such as an external string ID.
*/
//go:nosplit
func (self class) SetItemMetadata(idx gd.Int, metadata gd.Variant)  {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(metadata))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the tooltip of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemTooltip(idx gd.Int, tooltip gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, discreet.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the text of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemText(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the icon of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemIcon(idx gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the ID of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemId(idx gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the item with the given [param id].
*/
//go:nosplit
func (self class) GetItemIndex(id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_item_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieves the metadata of an item. Metadata may be any type and can be used to store extra information about an item, such as an external string ID.
*/
//go:nosplit
func (self class) GetItemMetadata(idx gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the tooltip of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemTooltip(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is disabled.
*/
//go:nosplit
func (self class) IsItemDisabled(idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is marked as a separator.
*/
//go:nosplit
func (self class) IsItemSeparator(idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_is_item_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a separator to the list of items. Separators help to group items, and can optionally be given a [param text] header. A separator also gets an index assigned, and is appended at the end of the item list.
*/
//go:nosplit
func (self class) AddSeparator(text gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_add_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all the items in the [OptionButton].
*/
//go:nosplit
func (self class) Clear()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Selects an item by index and makes it the current item. This will work even if the item is disabled.
Passing [code]-1[/code] as the index deselects any currently selected item.
*/
//go:nosplit
func (self class) Select(idx gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSelected() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID of the selected item, or [code]-1[/code] if no item is selected.
*/
//go:nosplit
func (self class) GetSelectedId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_selected_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the metadata of the selected item. Metadata for items can be set using [method set_item_metadata].
*/
//go:nosplit
func (self class) GetSelectedMetadata() gd.Variant {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_selected_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes the item at index [param idx].
*/
//go:nosplit
func (self class) RemoveItem(idx gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup() gdclass.PopupMenu {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.PopupMenu{classdb.PopupMenu(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Adjusts popup position and sizing for the [OptionButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
//go:nosplit
func (self class) ShowPopup()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_show_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetItemCount(count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetItemCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this button contains at least one item which is not disabled, or marked as a separator.
*/
//go:nosplit
func (self class) HasSelectableItems() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_has_selectable_items, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the first item which is not disabled, or marked as a separator. If [param from_last] is [code]true[/code], the items will be searched in reverse order.
Returns [code]-1[/code] if no item is found.
*/
//go:nosplit
func (self class) GetSelectableItem(from_last bool) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, from_last)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_selectable_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFitToLongestItem(fit bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, fit)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_fit_to_longest_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFitToLongestItem() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_is_fit_to_longest_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowReselect(allow bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowReselect() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
//go:nosplit
func (self class) SetDisableShortcuts(disabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_disable_shortcuts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnItemSelected(cb func(index int)) {
	self[0].AsObject().Connect(gd.NewStringName("item_selected"), gd.NewCallable(cb), 0)
}


func (self Go) OnItemFocused(cb func(index int)) {
	self[0].AsObject().Connect(gd.NewStringName("item_focused"), gd.NewCallable(cb), 0)
}


func (self class) AsOptionButton() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOptionButton() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("OptionButton", func(ptr gd.Object) any { return classdb.OptionButton(ptr) })}
