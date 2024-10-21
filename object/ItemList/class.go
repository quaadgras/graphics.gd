package ItemList

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This control provides a vertical list of selectable items that may be in a single or in multiple columns, with each item having options for text and an icon. Tooltips are supported and may be different for every item in the list.
Selectable items in the list may be selected or deselected and multiple selection may be enabled. Selection with right mouse button may also be enabled to allow use of popup context menus. Items may also be "activated" by double-clicking them or by pressing [kbd]Enter[/kbd].
Item text only supports single-line strings. Newline characters (e.g. [code]\n[/code]) in the string won't produce a newline. Text wrapping is enabled in [constant ICON_MODE_TOP] mode, but the column's width is adjusted to fully fit its content by default. You need to set [member fixed_column_width] greater than zero to wrap the text.
All [code]set_*[/code] methods allow negative item indices, i.e. [code]-1[/code] to access the last item, [code]-2[/code] to select the second-to-last item, and so on.
[b]Incremental search:[/b] Like [PopupMenu] and [Tree], [ItemList] supports searching within the list while the control is focused. Press a key that matches the first letter of an item's name to select the first item starting with the given letter. After that point, there are two ways to perform incremental search: 1) Press the same key again before the timeout duration to select the next item starting with the same letter. 2) Press letter keys that match the rest of the word before the timeout duration to match to select the item in question directly. Both of these actions will be reset to the beginning of the list if the timeout duration has passed since the last keystroke was registered. You can adjust the timeout duration by changing [member ProjectSettings.gui/timers/incremental_search_max_interval_msec].

*/
type Simple [1]classdb.ItemList
func (self Simple) AddItem(text string, icon [1]classdb.Texture2D, selectable bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddItem(gc.String(text), icon, selectable)))
}
func (self Simple) AddIconItem(icon [1]classdb.Texture2D, selectable bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddIconItem(icon, selectable)))
}
func (self Simple) SetItemText(idx int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemText(gd.Int(idx), gc.String(text))
}
func (self Simple) GetItemText(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemText(gc, gd.Int(idx)).String())
}
func (self Simple) SetItemIcon(idx int, icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIcon(gd.Int(idx), icon)
}
func (self Simple) GetItemIcon(idx int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetItemIcon(gc, gd.Int(idx)))
}
func (self Simple) SetItemTextDirection(idx int, direction classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemTextDirection(gd.Int(idx), direction)
}
func (self Simple) GetItemTextDirection(idx int) classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlTextDirection(Expert(self).GetItemTextDirection(gd.Int(idx)))
}
func (self Simple) SetItemLanguage(idx int, language string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemLanguage(gd.Int(idx), gc.String(language))
}
func (self Simple) GetItemLanguage(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemLanguage(gc, gd.Int(idx)).String())
}
func (self Simple) SetItemIconTransposed(idx int, transposed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIconTransposed(gd.Int(idx), transposed)
}
func (self Simple) IsItemIconTransposed(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemIconTransposed(gd.Int(idx)))
}
func (self Simple) SetItemIconRegion(idx int, rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIconRegion(gd.Int(idx), rect)
}
func (self Simple) GetItemIconRegion(idx int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetItemIconRegion(gd.Int(idx)))
}
func (self Simple) SetItemIconModulate(idx int, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemIconModulate(gd.Int(idx), modulate)
}
func (self Simple) GetItemIconModulate(idx int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetItemIconModulate(gd.Int(idx)))
}
func (self Simple) SetItemSelectable(idx int, selectable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemSelectable(gd.Int(idx), selectable)
}
func (self Simple) IsItemSelectable(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemSelectable(gd.Int(idx)))
}
func (self Simple) SetItemDisabled(idx int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemDisabled(gd.Int(idx), disabled)
}
func (self Simple) IsItemDisabled(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemDisabled(gd.Int(idx)))
}
func (self Simple) SetItemMetadata(idx int, metadata gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemMetadata(gd.Int(idx), metadata)
}
func (self Simple) GetItemMetadata(idx int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetItemMetadata(gc, gd.Int(idx)))
}
func (self Simple) SetItemCustomBgColor(idx int, custom_bg_color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemCustomBgColor(gd.Int(idx), custom_bg_color)
}
func (self Simple) GetItemCustomBgColor(idx int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetItemCustomBgColor(gd.Int(idx)))
}
func (self Simple) SetItemCustomFgColor(idx int, custom_fg_color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemCustomFgColor(gd.Int(idx), custom_fg_color)
}
func (self Simple) GetItemCustomFgColor(idx int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetItemCustomFgColor(gd.Int(idx)))
}
func (self Simple) GetItemRect(idx int, expand bool) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetItemRect(gd.Int(idx), expand))
}
func (self Simple) SetItemTooltipEnabled(idx int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemTooltipEnabled(gd.Int(idx), enable)
}
func (self Simple) IsItemTooltipEnabled(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsItemTooltipEnabled(gd.Int(idx)))
}
func (self Simple) SetItemTooltip(idx int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemTooltip(gd.Int(idx), gc.String(tooltip))
}
func (self Simple) GetItemTooltip(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetItemTooltip(gc, gd.Int(idx)).String())
}
func (self Simple) Select(idx int, single bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Select(gd.Int(idx), single)
}
func (self Simple) Deselect(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Deselect(gd.Int(idx))
}
func (self Simple) DeselectAll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DeselectAll()
}
func (self Simple) IsSelected(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSelected(gd.Int(idx)))
}
func (self Simple) GetSelectedItems() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetSelectedItems(gc))
}
func (self Simple) MoveItem(from_idx int, to_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveItem(gd.Int(from_idx), gd.Int(to_idx))
}
func (self Simple) SetItemCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetItemCount(gd.Int(count))
}
func (self Simple) GetItemCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemCount()))
}
func (self Simple) RemoveItem(idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveItem(gd.Int(idx))
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) SortItemsByText() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SortItemsByText()
}
func (self Simple) SetFixedColumnWidth(width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFixedColumnWidth(gd.Int(width))
}
func (self Simple) GetFixedColumnWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFixedColumnWidth()))
}
func (self Simple) SetSameColumnWidth(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSameColumnWidth(enable)
}
func (self Simple) IsSameColumnWidth() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSameColumnWidth())
}
func (self Simple) SetMaxTextLines(lines int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxTextLines(gd.Int(lines))
}
func (self Simple) GetMaxTextLines() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxTextLines()))
}
func (self Simple) SetMaxColumns(amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxColumns(gd.Int(amount))
}
func (self Simple) GetMaxColumns() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxColumns()))
}
func (self Simple) SetSelectMode(mode classdb.ItemListSelectMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelectMode(mode)
}
func (self Simple) GetSelectMode() classdb.ItemListSelectMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ItemListSelectMode(Expert(self).GetSelectMode())
}
func (self Simple) SetIconMode(mode classdb.ItemListIconMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIconMode(mode)
}
func (self Simple) GetIconMode() classdb.ItemListIconMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ItemListIconMode(Expert(self).GetIconMode())
}
func (self Simple) SetFixedIconSize(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFixedIconSize(size)
}
func (self Simple) GetFixedIconSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetFixedIconSize())
}
func (self Simple) SetIconScale(scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIconScale(gd.Float(scale))
}
func (self Simple) GetIconScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetIconScale()))
}
func (self Simple) SetAllowRmbSelect(allow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowRmbSelect(allow)
}
func (self Simple) GetAllowRmbSelect() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAllowRmbSelect())
}
func (self Simple) SetAllowReselect(allow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowReselect(allow)
}
func (self Simple) GetAllowReselect() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAllowReselect())
}
func (self Simple) SetAllowSearch(allow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowSearch(allow)
}
func (self Simple) GetAllowSearch() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetAllowSearch())
}
func (self Simple) SetAutoHeight(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoHeight(enable)
}
func (self Simple) HasAutoHeight() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAutoHeight())
}
func (self Simple) IsAnythingSelected() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAnythingSelected())
}
func (self Simple) GetItemAtPosition(position gd.Vector2, exact bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetItemAtPosition(position, exact)))
}
func (self Simple) EnsureCurrentIsVisible() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnsureCurrentIsVisible()
}
func (self Simple) GetVScrollBar() [1]classdb.VScrollBar {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.VScrollBar(Expert(self).GetVScrollBar(gc))
}
func (self Simple) SetTextOverrunBehavior(overrun_behavior classdb.TextServerOverrunBehavior) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextOverrunBehavior(overrun_behavior)
}
func (self Simple) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerOverrunBehavior(Expert(self).GetTextOverrunBehavior())
}
func (self Simple) ForceUpdateListSize() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceUpdateListSize()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ItemList
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Adds an item to the item list with specified text. Returns the index of an added item.
Specify an [param icon], or use [code]null[/code] as the [param icon] for a list item with no icon.
If selectable is [code]true[/code], the list item will be selectable.
*/
//go:nosplit
func (self class) AddItem(text gd.String, icon object.Texture2D, selectable bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, selectable)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an item to the item list with no text, only an icon. Returns the index of an added item.
*/
//go:nosplit
func (self class) AddIconItem(icon object.Texture2D, selectable bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, selectable)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets text of the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemText(idx gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the text associated with the specified index.
*/
//go:nosplit
func (self class) GetItemText(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets (or replaces) the icon's [Texture2D] associated with the specified index.
*/
//go:nosplit
func (self class) SetItemIcon(idx gd.Int, icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the icon associated with the specified index.
*/
//go:nosplit
func (self class) GetItemIcon(ctx gd.Lifetime, idx gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetItemTextDirection(idx gd.Int, direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetItemTextDirection(idx gd.Int) classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetItemLanguage(idx gd.Int, language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetItemLanguage(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets whether the item icon will be drawn transposed.
*/
//go:nosplit
func (self class) SetItemIconTransposed(idx gd.Int, transposed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, transposed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_icon_transposed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the item icon will be drawn transposed, i.e. the X and Y axes are swapped.
*/
//go:nosplit
func (self class) IsItemIconTransposed(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_is_item_icon_transposed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the region of item's icon used. The whole icon will be used if the region has no area.
*/
//go:nosplit
func (self class) SetItemIconRegion(idx gd.Int, rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the region of item's icon used. The whole icon will be used if the region has no area.
*/
//go:nosplit
func (self class) GetItemIconRegion(idx gd.Int) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a modulating [Color] of the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemIconModulate(idx gd.Int, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Color] modulating item's icon at the specified index.
*/
//go:nosplit
func (self class) GetItemIconModulate(idx gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Allows or disallows selection of the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemSelectable(idx gd.Int, selectable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, selectable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the item at the specified index is selectable.
*/
//go:nosplit
func (self class) IsItemSelectable(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_is_item_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Disables (or enables) the item at the specified index.
Disabled items cannot be selected and do not trigger activation signals (when double-clicking or pressing [kbd]Enter[/kbd]).
*/
//go:nosplit
func (self class) SetItemDisabled(idx gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the item at the specified index is disabled.
*/
//go:nosplit
func (self class) IsItemDisabled(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a value (of any type) to be stored with the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemMetadata(idx gd.Int, metadata gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(metadata))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the metadata value of the specified index.
*/
//go:nosplit
func (self class) GetItemMetadata(ctx gd.Lifetime, idx gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the background color of the item specified by [param idx] index to the specified [Color].
*/
//go:nosplit
func (self class) SetItemCustomBgColor(idx gd.Int, custom_bg_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, custom_bg_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom background color of the item specified by [param idx] index.
*/
//go:nosplit
func (self class) GetItemCustomBgColor(idx gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the foreground color of the item specified by [param idx] index to the specified [Color].
*/
//go:nosplit
func (self class) SetItemCustomFgColor(idx gd.Int, custom_fg_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, custom_fg_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_custom_fg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom foreground color of the item specified by [param idx] index.
*/
//go:nosplit
func (self class) GetItemCustomFgColor(idx gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_custom_fg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position and size of the item with the specified index, in the coordinate system of the [ItemList] node. If [param expand] is [code]true[/code] the last column expands to fill the rest of the row.
[b]Note:[/b] The returned value is unreliable if called right after modifying the [ItemList], before it redraws in the next frame.
*/
//go:nosplit
func (self class) GetItemRect(idx gd.Int, expand bool) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, expand)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets whether the tooltip hint is enabled for specified item index.
*/
//go:nosplit
func (self class) SetItemTooltipEnabled(idx gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_tooltip_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the tooltip is enabled for specified item index.
*/
//go:nosplit
func (self class) IsItemTooltipEnabled(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_is_item_tooltip_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the tooltip hint for the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemTooltip(idx gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tooltip hint associated with the specified index.
*/
//go:nosplit
func (self class) GetItemTooltip(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Select the item at the specified index.
[b]Note:[/b] This method does not trigger the item selection signal.
*/
//go:nosplit
func (self class) Select(idx gd.Int, single bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, single)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ensures the item associated with the specified index is not selected.
*/
//go:nosplit
func (self class) Deselect(idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ensures there are no items selected.
*/
//go:nosplit
func (self class) DeselectAll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_deselect_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the item at the specified index is currently selected.
*/
//go:nosplit
func (self class) IsSelected(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array with the indexes of the selected items.
*/
//go:nosplit
func (self class) GetSelectedItems(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_selected_items, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Moves item from index [param from_idx] to [param to_idx].
*/
//go:nosplit
func (self class) MoveItem(from_idx gd.Int, to_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_idx)
	callframe.Arg(frame, to_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_move_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetItemCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetItemCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the item specified by [param idx] index from the list.
*/
//go:nosplit
func (self class) RemoveItem(idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all items from the list.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sorts items in the list by their text.
*/
//go:nosplit
func (self class) SortItemsByText()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_sort_items_by_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFixedColumnWidth(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_fixed_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFixedColumnWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_fixed_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSameColumnWidth(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_same_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSameColumnWidth() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_is_same_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxTextLines(lines gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lines)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_max_text_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxTextLines() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_max_text_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxColumns(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_max_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxColumns() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_max_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectMode(mode classdb.ItemListSelectMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_select_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSelectMode() classdb.ItemListSelectMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ItemListSelectMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_select_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIconMode(mode classdb.ItemListIconMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_icon_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIconMode() classdb.ItemListIconMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ItemListIconMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_icon_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFixedIconSize(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_fixed_icon_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFixedIconSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_fixed_icon_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIconScale(scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_icon_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIconScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_icon_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowRmbSelect(allow bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_allow_rmb_select, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowRmbSelect() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_allow_rmb_select, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowReselect() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowSearch() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoHeight(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_auto_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) HasAutoHeight() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_has_auto_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if one or more items are selected.
*/
//go:nosplit
func (self class) IsAnythingSelected() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_is_anything_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the item index at the given [param position].
When there is no item at that point, -1 will be returned if [param exact] is [code]true[/code], and the closest item index will be returned otherwise.
[b]Note:[/b] The returned value is unreliable if called right after modifying the [ItemList], before it redraws in the next frame.
*/
//go:nosplit
func (self class) GetItemAtPosition(position gd.Vector2, exact bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, exact)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_item_at_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Ensure current selection is visible, adjusting the scroll position as necessary.
*/
//go:nosplit
func (self class) EnsureCurrentIsVisible()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_ensure_current_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the vertical scrollbar.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetVScrollBar(ctx gd.Lifetime) object.VScrollBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.VScrollBar
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior classdb.TextServerOverrunBehavior)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces an update to the list size based on its items. This happens automatically whenever size of the items, or other relevant settings like [member auto_height], change. The method can be used to trigger the update ahead of next drawing pass.
*/
//go:nosplit
func (self class) ForceUpdateListSize()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ItemList.Bind_force_update_list_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsItemList() Expert { return self[0].AsItemList() }


//go:nosplit
func (self Simple) AsItemList() Simple { return self[0].AsItemList() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ItemList", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type IconMode = classdb.ItemListIconMode

const (
/*Icon is drawn above the text.*/
	IconModeTop IconMode = 0
/*Icon is drawn to the left of the text.*/
	IconModeLeft IconMode = 1
)
type SelectMode = classdb.ItemListSelectMode

const (
/*Only allow selecting a single item.*/
	SelectSingle SelectMode = 0
/*Allows selecting multiple items by holding [kbd]Ctrl[/kbd] or [kbd]Shift[/kbd].*/
	SelectMulti SelectMode = 1
)
