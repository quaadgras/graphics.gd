package OptionButton

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Button"
import "graphics.gd/objects/BaseButton"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[OptionButton] is a type of button that brings up a dropdown with selectable items when pressed. The item selected becomes the "current" item and is displayed as the button text.
See also [BaseButton] which contains common properties and methods associated with this node.
[b]Note:[/b] The ID values used for items are limited to 32 bits, not full 64 bits of [int]. This has a range of [code]-2^32[/code] to [code]2^32 - 1[/code], i.e. [code]-2147483648[/code] to [code]2147483647[/code].
[b]Note:[/b] The [member Button.text] and [member Button.icon] properties are set automatically based on the selected item. They shouldn't be changed manually.
*/
type Instance [1]classdb.OptionButton
type Any interface {
	gd.IsClass
	AsOptionButton() Instance
}

/*
Adds an item, with text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
func (self Instance) AddItem(label string) {
	class(self).AddItem(gd.NewString(label), gd.Int(-1))
}

/*
Adds an item, with a [param texture] icon, text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
func (self Instance) AddIconItem(texture objects.Texture2D, label string) {
	class(self).AddIconItem(texture, gd.NewString(label), gd.Int(-1))
}

/*
Sets the text of the item at index [param idx].
*/
func (self Instance) SetItemText(idx int, text string) {
	class(self).SetItemText(gd.Int(idx), gd.NewString(text))
}

/*
Sets the icon of the item at index [param idx].
*/
func (self Instance) SetItemIcon(idx int, texture objects.Texture2D) {
	class(self).SetItemIcon(gd.Int(idx), texture)
}

/*
Sets whether the item at index [param idx] is disabled.
Disabled items are drawn differently in the dropdown and are not selectable by the user. If the current selected item is set as disabled, it will remain selected.
*/
func (self Instance) SetItemDisabled(idx int, disabled bool) {
	class(self).SetItemDisabled(gd.Int(idx), disabled)
}

/*
Sets the ID of the item at index [param idx].
*/
func (self Instance) SetItemId(idx int, id int) {
	class(self).SetItemId(gd.Int(idx), gd.Int(id))
}

/*
Sets the metadata of an item. Metadata may be of any type and can be used to store extra information about an item, such as an external string ID.
*/
func (self Instance) SetItemMetadata(idx int, metadata any) {
	class(self).SetItemMetadata(gd.Int(idx), gd.NewVariant(metadata))
}

/*
Sets the tooltip of the item at index [param idx].
*/
func (self Instance) SetItemTooltip(idx int, tooltip string) {
	class(self).SetItemTooltip(gd.Int(idx), gd.NewString(tooltip))
}

/*
Returns the text of the item at index [param idx].
*/
func (self Instance) GetItemText(idx int) string {
	return string(class(self).GetItemText(gd.Int(idx)).String())
}

/*
Returns the icon of the item at index [param idx].
*/
func (self Instance) GetItemIcon(idx int) objects.Texture2D {
	return objects.Texture2D(class(self).GetItemIcon(gd.Int(idx)))
}

/*
Returns the ID of the item at index [param idx].
*/
func (self Instance) GetItemId(idx int) int {
	return int(int(class(self).GetItemId(gd.Int(idx))))
}

/*
Returns the index of the item with the given [param id].
*/
func (self Instance) GetItemIndex(id int) int {
	return int(int(class(self).GetItemIndex(gd.Int(id))))
}

/*
Retrieves the metadata of an item. Metadata may be any type and can be used to store extra information about an item, such as an external string ID.
*/
func (self Instance) GetItemMetadata(idx int) any {
	return any(class(self).GetItemMetadata(gd.Int(idx)).Interface())
}

/*
Returns the tooltip of the item at index [param idx].
*/
func (self Instance) GetItemTooltip(idx int) string {
	return string(class(self).GetItemTooltip(gd.Int(idx)).String())
}

/*
Returns [code]true[/code] if the item at index [param idx] is disabled.
*/
func (self Instance) IsItemDisabled(idx int) bool {
	return bool(class(self).IsItemDisabled(gd.Int(idx)))
}

/*
Returns [code]true[/code] if the item at index [param idx] is marked as a separator.
*/
func (self Instance) IsItemSeparator(idx int) bool {
	return bool(class(self).IsItemSeparator(gd.Int(idx)))
}

/*
Adds a separator to the list of items. Separators help to group items, and can optionally be given a [param text] header. A separator also gets an index assigned, and is appended at the end of the item list.
*/
func (self Instance) AddSeparator() {
	class(self).AddSeparator(gd.NewString(""))
}

/*
Clears all the items in the [OptionButton].
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Selects an item by index and makes it the current item. This will work even if the item is disabled.
Passing [code]-1[/code] as the index deselects any currently selected item.
*/
func (self Instance) Select(idx int) {
	class(self).Select(gd.Int(idx))
}

/*
Returns the ID of the selected item, or [code]-1[/code] if no item is selected.
*/
func (self Instance) GetSelectedId() int {
	return int(int(class(self).GetSelectedId()))
}

/*
Gets the metadata of the selected item. Metadata for items can be set using [method set_item_metadata].
*/
func (self Instance) GetSelectedMetadata() any {
	return any(class(self).GetSelectedMetadata().Interface())
}

/*
Removes the item at index [param idx].
*/
func (self Instance) RemoveItem(idx int) {
	class(self).RemoveItem(gd.Int(idx))
}

/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Instance) GetPopup() objects.PopupMenu {
	return objects.PopupMenu(class(self).GetPopup())
}

/*
Adjusts popup position and sizing for the [OptionButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
func (self Instance) ShowPopup() {
	class(self).ShowPopup()
}

/*
Returns [code]true[/code] if this button contains at least one item which is not disabled, or marked as a separator.
*/
func (self Instance) HasSelectableItems() bool {
	return bool(class(self).HasSelectableItems())
}

/*
Returns the index of the first item which is not disabled, or marked as a separator. If [param from_last] is [code]true[/code], the items will be searched in reverse order.
Returns [code]-1[/code] if no item is found.
*/
func (self Instance) GetSelectableItem() int {
	return int(int(class(self).GetSelectableItem(false)))
}

/*
If [code]true[/code], shortcuts are disabled and cannot be used to trigger the button.
*/
func (self Instance) SetDisableShortcuts(disabled bool) {
	class(self).SetDisableShortcuts(disabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OptionButton

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OptionButton"))
	return Instance{*(*classdb.OptionButton)(unsafe.Pointer(&object))}
}

func (self Instance) Selected() int {
	return int(int(class(self).GetSelected()))
}

func (self Instance) FitToLongestItem() bool {
	return bool(class(self).IsFitToLongestItem())
}

func (self Instance) SetFitToLongestItem(value bool) {
	class(self).SetFitToLongestItem(value)
}

func (self Instance) AllowReselect() bool {
	return bool(class(self).GetAllowReselect())
}

func (self Instance) SetAllowReselect(value bool) {
	class(self).SetAllowReselect(value)
}

func (self Instance) ItemCount() int {
	return int(int(class(self).GetItemCount()))
}

func (self Instance) SetItemCount(value int) {
	class(self).SetItemCount(gd.Int(value))
}

/*
Adds an item, with text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
//go:nosplit
func (self class) AddItem(label gd.String, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds an item, with a [param texture] icon, text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
//go:nosplit
func (self class) AddIconItem(texture objects.Texture2D, label gd.String, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	callframe.Arg(frame, pointers.Get(label))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the text of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemText(idx gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the icon of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemIcon(idx gd.Int, texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets whether the item at index [param idx] is disabled.
Disabled items are drawn differently in the dropdown and are not selectable by the user. If the current selected item is set as disabled, it will remain selected.
*/
//go:nosplit
func (self class) SetItemDisabled(idx gd.Int, disabled bool) {
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
func (self class) SetItemId(idx gd.Int, id gd.Int) {
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
func (self class) SetItemMetadata(idx gd.Int, metadata gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(metadata))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the tooltip of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemTooltip(idx gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(tooltip))
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
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the icon of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemIcon(idx gd.Int) objects.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{gd.PointerWithOwnershipTransferredToGo[classdb.Texture2D](r_ret.Get())}
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
	var ret = pointers.New[gd.Variant](r_ret.Get())
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
	var ret = pointers.New[gd.String](r_ret.Get())
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
func (self class) AddSeparator(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_add_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears all the items in the [OptionButton].
*/
//go:nosplit
func (self class) Clear() {
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
func (self class) Select(idx gd.Int) {
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
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes the item at index [param idx].
*/
//go:nosplit
func (self class) RemoveItem(idx gd.Int) {
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
func (self class) GetPopup() objects.PopupMenu {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.PopupMenu{gd.PointerLifetimeBoundTo[classdb.PopupMenu](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adjusts popup position and sizing for the [OptionButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
//go:nosplit
func (self class) ShowPopup() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_show_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetItemCount(count gd.Int) {
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
func (self class) SetFitToLongestItem(fit bool) {
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
func (self class) SetAllowReselect(allow bool) {
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
func (self class) SetDisableShortcuts(disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OptionButton.Bind_set_disable_shortcuts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnItemSelected(cb func(index int)) {
	self[0].AsObject().Connect(gd.NewStringName("item_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemFocused(cb func(index int)) {
	self[0].AsObject().Connect(gd.NewStringName("item_focused"), gd.NewCallable(cb), 0)
}

func (self class) AsOptionButton() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOptionButton() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsButton() Button.Advanced    { return *((*Button.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsButton() Button.Instance { return *((*Button.Instance)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.Advanced {
	return *((*BaseButton.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBaseButton() BaseButton.Instance {
	return *((*BaseButton.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsButton(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsButton(), name)
	}
}
func init() {
	classdb.Register("OptionButton", func(ptr gd.Object) any {
		return [1]classdb.OptionButton{*(*classdb.OptionButton)(unsafe.Pointer(&ptr))}
	})
}
