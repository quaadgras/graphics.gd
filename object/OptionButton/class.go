package OptionButton

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Button"
import "grow.graphics/gd/object/BaseButton"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[OptionButton] is a type of button that brings up a dropdown with selectable items when pressed. The item selected becomes the "current" item and is displayed as the button text.
See also [BaseButton] which contains common properties and methods associated with this node.
[b]Note:[/b] The ID values used for items are limited to 32 bits, not full 64 bits of [int]. This has a range of [code]-2^32[/code] to [code]2^32 - 1[/code], i.e. [code]-2147483648[/code] to [code]2147483647[/code].
[b]Note:[/b] The [member Button.text] and [member Button.icon] properties are set automatically based on the selected item. They shouldn't be changed manually.

*/
type Simple [1]classdb.OptionButton
func (self Simple) AddItem(label string, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddItem(gc.String(label), gd.Int(id))
}
func (self Simple) AddIconItem(texture [1]classdb.Texture2D, label string, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddIconItem(texture, gc.String(label), gd.Int(id))
}
func (self Simple) SetItemText(idx int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemText(gd.Int(idx), gc.String(text))
}
func (self Simple) SetItemIcon(idx int, texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIcon(gd.Int(idx), texture)
}
func (self Simple) SetItemDisabled(idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemDisabled(gd.Int(idx), disabled)
}
func (self Simple) SetItemId(idx int, id int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemId(gd.Int(idx), gd.Int(id))
}
func (self Simple) SetItemMetadata(idx int, metadata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemMetadata(gd.Int(idx), metadata)
}
func (self Simple) SetItemTooltip(idx int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemTooltip(gd.Int(idx), gc.String(tooltip))
}
func (self Simple) GetItemText(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemText(gc, gd.Int(idx)).String())
}
func (self Simple) GetItemIcon(idx int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetItemIcon(gc, gd.Int(idx)))
}
func (self Simple) GetItemId(idx int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemId(gd.Int(idx))))
}
func (self Simple) GetItemIndex(id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemIndex(gd.Int(id))))
}
func (self Simple) GetItemMetadata(idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetItemMetadata(gc, gd.Int(idx)))
}
func (self Simple) GetItemTooltip(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemTooltip(gc, gd.Int(idx)).String())
}
func (self Simple) IsItemDisabled(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemDisabled(gd.Int(idx)))
}
func (self Simple) IsItemSeparator(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemSeparator(gd.Int(idx)))
}
func (self Simple) AddSeparator(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddSeparator(gc.String(text))
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) Select(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Select(gd.Int(idx))
}
func (self Simple) GetSelected() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSelected()))
}
func (self Simple) GetSelectedId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSelectedId()))
}
func (self Simple) GetSelectedMetadata() gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetSelectedMetadata(gc))
}
func (self Simple) RemoveItem(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveItem(gd.Int(idx))
}
func (self Simple) GetPopup() [1]classdb.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PopupMenu(Expert(self).GetPopup(gc))
}
func (self Simple) ShowPopup() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ShowPopup()
}
func (self Simple) SetItemCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemCount(gd.Int(count))
}
func (self Simple) GetItemCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemCount()))
}
func (self Simple) HasSelectableItems() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasSelectableItems())
}
func (self Simple) GetSelectableItem(from_last bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSelectableItem(from_last)))
}
func (self Simple) SetFitToLongestItem(fit bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFitToLongestItem(fit)
}
func (self Simple) IsFitToLongestItem() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFitToLongestItem())
}
func (self Simple) SetAllowReselect(allow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowReselect(allow)
}
func (self Simple) GetAllowReselect() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAllowReselect())
}
func (self Simple) SetDisableShortcuts(disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableShortcuts(disabled)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OptionButton
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Adds an item, with text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
//go:nosplit
func (self class) AddItem(label gd.String, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an item, with a [param texture] icon, text [param label] and (optionally) [param id]. If no [param id] is passed, the item index will be used as the item's ID. New items are appended at the end.
*/
//go:nosplit
func (self class) AddIconItem(texture object.Texture2D, label gd.String, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(label))
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the text of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemText(idx gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the icon of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemIcon(idx gd.Int, texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets whether the item at index [param idx] is disabled.
Disabled items are drawn differently in the dropdown and are not selectable by the user. If the current selected item is set as disabled, it will remain selected.
*/
//go:nosplit
func (self class) SetItemDisabled(idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the ID of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemId(idx gd.Int, id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the metadata of an item. Metadata may be of any type and can be used to store extra information about an item, such as an external string ID.
*/
//go:nosplit
func (self class) SetItemMetadata(idx gd.Int, metadata gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(metadata))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the tooltip of the item at index [param idx].
*/
//go:nosplit
func (self class) SetItemTooltip(idx gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the text of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemText(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the icon of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemIcon(ctx gd.Lifetime, idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the ID of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemId(idx gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_item_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the item with the given [param id].
*/
//go:nosplit
func (self class) GetItemIndex(id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_item_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieves the metadata of an item. Metadata may be any type and can be used to store extra information about an item, such as an external string ID.
*/
//go:nosplit
func (self class) GetItemMetadata(ctx gd.Lifetime, idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the tooltip of the item at index [param idx].
*/
//go:nosplit
func (self class) GetItemTooltip(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is disabled.
*/
//go:nosplit
func (self class) IsItemDisabled(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the item at index [param idx] is marked as a separator.
*/
//go:nosplit
func (self class) IsItemSeparator(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_is_item_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a separator to the list of items. Separators help to group items, and can optionally be given a [param text] header. A separator also gets an index assigned, and is appended at the end of the item list.
*/
//go:nosplit
func (self class) AddSeparator(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_add_separator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all the items in the [OptionButton].
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Selects an item by index and makes it the current item. This will work even if the item is disabled.
Passing [code]-1[/code] as the index deselects any currently selected item.
*/
//go:nosplit
func (self class) Select(idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSelected() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID of the selected item, or [code]-1[/code] if no item is selected.
*/
//go:nosplit
func (self class) GetSelectedId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_selected_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the metadata of the selected item. Metadata for items can be set using [method set_item_metadata].
*/
//go:nosplit
func (self class) GetSelectedMetadata(ctx gd.Lifetime) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_selected_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes the item at index [param idx].
*/
//go:nosplit
func (self class) RemoveItem(idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [PopupMenu] contained in this button.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup(ctx gd.Lifetime) object.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PopupMenu
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adjusts popup position and sizing for the [OptionButton], then shows the [PopupMenu]. Prefer this over using [code]get_popup().popup()[/code].
*/
//go:nosplit
func (self class) ShowPopup()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_show_popup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetItemCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetItemCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this button contains at least one item which is not disabled, or marked as a separator.
*/
//go:nosplit
func (self class) HasSelectableItems() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_has_selectable_items, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_last)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_selectable_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFitToLongestItem(fit bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fit)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_fit_to_longest_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFitToLongestItem() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_is_fit_to_longest_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowReselect(allow bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowReselect() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_get_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OptionButton.Bind_set_disable_shortcuts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsOptionButton() Expert { return self[0].AsOptionButton() }


//go:nosplit
func (self Simple) AsOptionButton() Simple { return self[0].AsOptionButton() }


//go:nosplit
func (self class) AsButton() Button.Expert { return self[0].AsButton() }


//go:nosplit
func (self Simple) AsButton() Button.Simple { return self[0].AsButton() }


//go:nosplit
func (self class) AsBaseButton() BaseButton.Expert { return self[0].AsBaseButton() }


//go:nosplit
func (self Simple) AsBaseButton() BaseButton.Simple { return self[0].AsBaseButton() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("OptionButton", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
