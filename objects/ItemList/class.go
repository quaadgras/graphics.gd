package ItemList

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This control provides a vertical list of selectable items that may be in a single or in multiple columns, with each item having options for text and an icon. Tooltips are supported and may be different for every item in the list.
Selectable items in the list may be selected or deselected and multiple selection may be enabled. Selection with right mouse button may also be enabled to allow use of popup context menus. Items may also be "activated" by double-clicking them or by pressing [kbd]Enter[/kbd].
Item text only supports single-line strings. Newline characters (e.g. [code]\n[/code]) in the string won't produce a newline. Text wrapping is enabled in [constant ICON_MODE_TOP] mode, but the column's width is adjusted to fully fit its content by default. You need to set [member fixed_column_width] greater than zero to wrap the text.
All [code]set_*[/code] methods allow negative item indices, i.e. [code]-1[/code] to access the last item, [code]-2[/code] to select the second-to-last item, and so on.
[b]Incremental search:[/b] Like [PopupMenu] and [Tree], [ItemList] supports searching within the list while the control is focused. Press a key that matches the first letter of an item's name to select the first item starting with the given letter. After that point, there are two ways to perform incremental search: 1) Press the same key again before the timeout duration to select the next item starting with the same letter. 2) Press letter keys that match the rest of the word before the timeout duration to match to select the item in question directly. Both of these actions will be reset to the beginning of the list if the timeout duration has passed since the last keystroke was registered. You can adjust the timeout duration by changing [member ProjectSettings.gui/timers/incremental_search_max_interval_msec].
*/
type Instance [1]classdb.ItemList
type Any interface {
	gd.IsClass
	AsItemList() Instance
}

/*
Adds an item to the item list with specified text. Returns the index of an added item.
Specify an [param icon], or use [code]null[/code] as the [param icon] for a list item with no icon.
If selectable is [code]true[/code], the list item will be selectable.
*/
func (self Instance) AddItem(text string) int {
	return int(int(class(self).AddItem(gd.NewString(text), [1]objects.Texture2D{}[0], true)))
}

/*
Adds an item to the item list with no text, only an icon. Returns the index of an added item.
*/
func (self Instance) AddIconItem(icon objects.Texture2D) int {
	return int(int(class(self).AddIconItem(icon, true)))
}

/*
Sets text of the item associated with the specified index.
*/
func (self Instance) SetItemText(idx int, text string) {
	class(self).SetItemText(gd.Int(idx), gd.NewString(text))
}

/*
Returns the text associated with the specified index.
*/
func (self Instance) GetItemText(idx int) string {
	return string(class(self).GetItemText(gd.Int(idx)).String())
}

/*
Sets (or replaces) the icon's [Texture2D] associated with the specified index.
*/
func (self Instance) SetItemIcon(idx int, icon objects.Texture2D) {
	class(self).SetItemIcon(gd.Int(idx), icon)
}

/*
Returns the icon associated with the specified index.
*/
func (self Instance) GetItemIcon(idx int) objects.Texture2D {
	return objects.Texture2D(class(self).GetItemIcon(gd.Int(idx)))
}

/*
Sets item's text base writing direction.
*/
func (self Instance) SetItemTextDirection(idx int, direction classdb.ControlTextDirection) {
	class(self).SetItemTextDirection(gd.Int(idx), direction)
}

/*
Returns item's text base writing direction.
*/
func (self Instance) GetItemTextDirection(idx int) classdb.ControlTextDirection {
	return classdb.ControlTextDirection(class(self).GetItemTextDirection(gd.Int(idx)))
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Instance) SetItemLanguage(idx int, language string) {
	class(self).SetItemLanguage(gd.Int(idx), gd.NewString(language))
}

/*
Returns item's text language code.
*/
func (self Instance) GetItemLanguage(idx int) string {
	return string(class(self).GetItemLanguage(gd.Int(idx)).String())
}

/*
Sets whether the item icon will be drawn transposed.
*/
func (self Instance) SetItemIconTransposed(idx int, transposed bool) {
	class(self).SetItemIconTransposed(gd.Int(idx), transposed)
}

/*
Returns [code]true[/code] if the item icon will be drawn transposed, i.e. the X and Y axes are swapped.
*/
func (self Instance) IsItemIconTransposed(idx int) bool {
	return bool(class(self).IsItemIconTransposed(gd.Int(idx)))
}

/*
Sets the region of item's icon used. The whole icon will be used if the region has no area.
*/
func (self Instance) SetItemIconRegion(idx int, rect Rect2.PositionSize) {
	class(self).SetItemIconRegion(gd.Int(idx), gd.Rect2(rect))
}

/*
Returns the region of item's icon used. The whole icon will be used if the region has no area.
*/
func (self Instance) GetItemIconRegion(idx int) Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetItemIconRegion(gd.Int(idx)))
}

/*
Sets a modulating [Color] of the item associated with the specified index.
*/
func (self Instance) SetItemIconModulate(idx int, modulate Color.RGBA) {
	class(self).SetItemIconModulate(gd.Int(idx), gd.Color(modulate))
}

/*
Returns a [Color] modulating item's icon at the specified index.
*/
func (self Instance) GetItemIconModulate(idx int) Color.RGBA {
	return Color.RGBA(class(self).GetItemIconModulate(gd.Int(idx)))
}

/*
Allows or disallows selection of the item associated with the specified index.
*/
func (self Instance) SetItemSelectable(idx int, selectable bool) {
	class(self).SetItemSelectable(gd.Int(idx), selectable)
}

/*
Returns [code]true[/code] if the item at the specified index is selectable.
*/
func (self Instance) IsItemSelectable(idx int) bool {
	return bool(class(self).IsItemSelectable(gd.Int(idx)))
}

/*
Disables (or enables) the item at the specified index.
Disabled items cannot be selected and do not trigger activation signals (when double-clicking or pressing [kbd]Enter[/kbd]).
*/
func (self Instance) SetItemDisabled(idx int, disabled bool) {
	class(self).SetItemDisabled(gd.Int(idx), disabled)
}

/*
Returns [code]true[/code] if the item at the specified index is disabled.
*/
func (self Instance) IsItemDisabled(idx int) bool {
	return bool(class(self).IsItemDisabled(gd.Int(idx)))
}

/*
Sets a value (of any type) to be stored with the item associated with the specified index.
*/
func (self Instance) SetItemMetadata(idx int, metadata any) {
	class(self).SetItemMetadata(gd.Int(idx), gd.NewVariant(metadata))
}

/*
Returns the metadata value of the specified index.
*/
func (self Instance) GetItemMetadata(idx int) any {
	return any(class(self).GetItemMetadata(gd.Int(idx)).Interface())
}

/*
Sets the background color of the item specified by [param idx] index to the specified [Color].
*/
func (self Instance) SetItemCustomBgColor(idx int, custom_bg_color Color.RGBA) {
	class(self).SetItemCustomBgColor(gd.Int(idx), gd.Color(custom_bg_color))
}

/*
Returns the custom background color of the item specified by [param idx] index.
*/
func (self Instance) GetItemCustomBgColor(idx int) Color.RGBA {
	return Color.RGBA(class(self).GetItemCustomBgColor(gd.Int(idx)))
}

/*
Sets the foreground color of the item specified by [param idx] index to the specified [Color].
*/
func (self Instance) SetItemCustomFgColor(idx int, custom_fg_color Color.RGBA) {
	class(self).SetItemCustomFgColor(gd.Int(idx), gd.Color(custom_fg_color))
}

/*
Returns the custom foreground color of the item specified by [param idx] index.
*/
func (self Instance) GetItemCustomFgColor(idx int) Color.RGBA {
	return Color.RGBA(class(self).GetItemCustomFgColor(gd.Int(idx)))
}

/*
Returns the position and size of the item with the specified index, in the coordinate system of the [ItemList] node. If [param expand] is [code]true[/code] the last column expands to fill the rest of the row.
[b]Note:[/b] The returned value is unreliable if called right after modifying the [ItemList], before it redraws in the next frame.
*/
func (self Instance) GetItemRect(idx int) Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetItemRect(gd.Int(idx), true))
}

/*
Sets whether the tooltip hint is enabled for specified item index.
*/
func (self Instance) SetItemTooltipEnabled(idx int, enable bool) {
	class(self).SetItemTooltipEnabled(gd.Int(idx), enable)
}

/*
Returns [code]true[/code] if the tooltip is enabled for specified item index.
*/
func (self Instance) IsItemTooltipEnabled(idx int) bool {
	return bool(class(self).IsItemTooltipEnabled(gd.Int(idx)))
}

/*
Sets the tooltip hint for the item associated with the specified index.
*/
func (self Instance) SetItemTooltip(idx int, tooltip string) {
	class(self).SetItemTooltip(gd.Int(idx), gd.NewString(tooltip))
}

/*
Returns the tooltip hint associated with the specified index.
*/
func (self Instance) GetItemTooltip(idx int) string {
	return string(class(self).GetItemTooltip(gd.Int(idx)).String())
}

/*
Select the item at the specified index.
[b]Note:[/b] This method does not trigger the item selection signal.
*/
func (self Instance) Select(idx int) {
	class(self).Select(gd.Int(idx), true)
}

/*
Ensures the item associated with the specified index is not selected.
*/
func (self Instance) Deselect(idx int) {
	class(self).Deselect(gd.Int(idx))
}

/*
Ensures there are no items selected.
*/
func (self Instance) DeselectAll() {
	class(self).DeselectAll()
}

/*
Returns [code]true[/code] if the item at the specified index is currently selected.
*/
func (self Instance) IsSelected(idx int) bool {
	return bool(class(self).IsSelected(gd.Int(idx)))
}

/*
Returns an array with the indexes of the selected items.
*/
func (self Instance) GetSelectedItems() []int32 {
	return []int32(class(self).GetSelectedItems().AsSlice())
}

/*
Moves item from index [param from_idx] to [param to_idx].
*/
func (self Instance) MoveItem(from_idx int, to_idx int) {
	class(self).MoveItem(gd.Int(from_idx), gd.Int(to_idx))
}

/*
Removes the item specified by [param idx] index from the list.
*/
func (self Instance) RemoveItem(idx int) {
	class(self).RemoveItem(gd.Int(idx))
}

/*
Removes all items from the list.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Sorts items in the list by their text.
*/
func (self Instance) SortItemsByText() {
	class(self).SortItemsByText()
}

/*
Returns [code]true[/code] if one or more items are selected.
*/
func (self Instance) IsAnythingSelected() bool {
	return bool(class(self).IsAnythingSelected())
}

/*
Returns the item index at the given [param position].
When there is no item at that point, -1 will be returned if [param exact] is [code]true[/code], and the closest item index will be returned otherwise.
[b]Note:[/b] The returned value is unreliable if called right after modifying the [ItemList], before it redraws in the next frame.
*/
func (self Instance) GetItemAtPosition(position Vector2.XY) int {
	return int(int(class(self).GetItemAtPosition(gd.Vector2(position), false)))
}

/*
Ensure current selection is visible, adjusting the scroll position as necessary.
*/
func (self Instance) EnsureCurrentIsVisible() {
	class(self).EnsureCurrentIsVisible()
}

/*
Returns the vertical scrollbar.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetVScrollBar() objects.VScrollBar {
	return objects.VScrollBar(class(self).GetVScrollBar())
}

/*
Forces an update to the list size based on its items. This happens automatically whenever size of the items, or other relevant settings like [member auto_height], change. The method can be used to trigger the update ahead of next drawing pass.
*/
func (self Instance) ForceUpdateListSize() {
	class(self).ForceUpdateListSize()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ItemList

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ItemList"))
	return Instance{classdb.ItemList(object)}
}

func (self Instance) SelectMode() classdb.ItemListSelectMode {
	return classdb.ItemListSelectMode(class(self).GetSelectMode())
}

func (self Instance) SetSelectMode(value classdb.ItemListSelectMode) {
	class(self).SetSelectMode(value)
}

func (self Instance) AllowReselect() bool {
	return bool(class(self).GetAllowReselect())
}

func (self Instance) SetAllowReselect(value bool) {
	class(self).SetAllowReselect(value)
}

func (self Instance) AllowRmbSelect() bool {
	return bool(class(self).GetAllowRmbSelect())
}

func (self Instance) SetAllowRmbSelect(value bool) {
	class(self).SetAllowRmbSelect(value)
}

func (self Instance) AllowSearch() bool {
	return bool(class(self).GetAllowSearch())
}

func (self Instance) SetAllowSearch(value bool) {
	class(self).SetAllowSearch(value)
}

func (self Instance) MaxTextLines() int {
	return int(int(class(self).GetMaxTextLines()))
}

func (self Instance) SetMaxTextLines(value int) {
	class(self).SetMaxTextLines(gd.Int(value))
}

func (self Instance) AutoHeight() bool {
	return bool(class(self).HasAutoHeight())
}

func (self Instance) SetAutoHeight(value bool) {
	class(self).SetAutoHeight(value)
}

func (self Instance) TextOverrunBehavior() classdb.TextServerOverrunBehavior {
	return classdb.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior())
}

func (self Instance) SetTextOverrunBehavior(value classdb.TextServerOverrunBehavior) {
	class(self).SetTextOverrunBehavior(value)
}

func (self Instance) ItemCount() int {
	return int(int(class(self).GetItemCount()))
}

func (self Instance) SetItemCount(value int) {
	class(self).SetItemCount(gd.Int(value))
}

func (self Instance) MaxColumns() int {
	return int(int(class(self).GetMaxColumns()))
}

func (self Instance) SetMaxColumns(value int) {
	class(self).SetMaxColumns(gd.Int(value))
}

func (self Instance) SameColumnWidth() bool {
	return bool(class(self).IsSameColumnWidth())
}

func (self Instance) SetSameColumnWidth(value bool) {
	class(self).SetSameColumnWidth(value)
}

func (self Instance) FixedColumnWidth() int {
	return int(int(class(self).GetFixedColumnWidth()))
}

func (self Instance) SetFixedColumnWidth(value int) {
	class(self).SetFixedColumnWidth(gd.Int(value))
}

func (self Instance) IconMode() classdb.ItemListIconMode {
	return classdb.ItemListIconMode(class(self).GetIconMode())
}

func (self Instance) SetIconMode(value classdb.ItemListIconMode) {
	class(self).SetIconMode(value)
}

func (self Instance) IconScale() Float.X {
	return Float.X(Float.X(class(self).GetIconScale()))
}

func (self Instance) SetIconScale(value Float.X) {
	class(self).SetIconScale(gd.Float(value))
}

func (self Instance) FixedIconSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetFixedIconSize())
}

func (self Instance) SetFixedIconSize(value Vector2i.XY) {
	class(self).SetFixedIconSize(gd.Vector2i(value))
}

/*
Adds an item to the item list with specified text. Returns the index of an added item.
Specify an [param icon], or use [code]null[/code] as the [param icon] for a list item with no icon.
If selectable is [code]true[/code], the list item will be selectable.
*/
//go:nosplit
func (self class) AddItem(text gd.String, icon objects.Texture2D, selectable bool) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, selectable)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_add_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds an item to the item list with no text, only an icon. Returns the index of an added item.
*/
//go:nosplit
func (self class) AddIconItem(icon objects.Texture2D, selectable bool) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, selectable)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_add_icon_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets text of the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemText(idx gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the text associated with the specified index.
*/
//go:nosplit
func (self class) GetItemText(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets (or replaces) the icon's [Texture2D] associated with the specified index.
*/
//go:nosplit
func (self class) SetItemIcon(idx gd.Int, icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the icon associated with the specified index.
*/
//go:nosplit
func (self class) GetItemIcon(idx gd.Int) objects.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetItemTextDirection(idx gd.Int, direction classdb.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetItemTextDirection(idx gd.Int) classdb.ControlTextDirection {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetItemLanguage(idx gd.Int, language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetItemLanguage(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets whether the item icon will be drawn transposed.
*/
//go:nosplit
func (self class) SetItemIconTransposed(idx gd.Int, transposed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, transposed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_icon_transposed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the item icon will be drawn transposed, i.e. the X and Y axes are swapped.
*/
//go:nosplit
func (self class) IsItemIconTransposed(idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_is_item_icon_transposed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the region of item's icon used. The whole icon will be used if the region has no area.
*/
//go:nosplit
func (self class) SetItemIconRegion(idx gd.Int, rect gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the region of item's icon used. The whole icon will be used if the region has no area.
*/
//go:nosplit
func (self class) GetItemIconRegion(idx gd.Int) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a modulating [Color] of the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemIconModulate(idx gd.Int, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [Color] modulating item's icon at the specified index.
*/
//go:nosplit
func (self class) GetItemIconModulate(idx gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Allows or disallows selection of the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemSelectable(idx gd.Int, selectable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, selectable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the item at the specified index is selectable.
*/
//go:nosplit
func (self class) IsItemSelectable(idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_is_item_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Disables (or enables) the item at the specified index.
Disabled items cannot be selected and do not trigger activation signals (when double-clicking or pressing [kbd]Enter[/kbd]).
*/
//go:nosplit
func (self class) SetItemDisabled(idx gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the item at the specified index is disabled.
*/
//go:nosplit
func (self class) IsItemDisabled(idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_is_item_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a value (of any type) to be stored with the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemMetadata(idx gd.Int, metadata gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(metadata))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the metadata value of the specified index.
*/
//go:nosplit
func (self class) GetItemMetadata(idx gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the background color of the item specified by [param idx] index to the specified [Color].
*/
//go:nosplit
func (self class) SetItemCustomBgColor(idx gd.Int, custom_bg_color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, custom_bg_color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the custom background color of the item specified by [param idx] index.
*/
//go:nosplit
func (self class) GetItemCustomBgColor(idx gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the foreground color of the item specified by [param idx] index to the specified [Color].
*/
//go:nosplit
func (self class) SetItemCustomFgColor(idx gd.Int, custom_fg_color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, custom_fg_color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_custom_fg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the custom foreground color of the item specified by [param idx] index.
*/
//go:nosplit
func (self class) GetItemCustomFgColor(idx gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_custom_fg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, expand)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the tooltip hint is enabled for specified item index.
*/
//go:nosplit
func (self class) SetItemTooltipEnabled(idx gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_tooltip_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the tooltip is enabled for specified item index.
*/
//go:nosplit
func (self class) IsItemTooltipEnabled(idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_is_item_tooltip_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tooltip hint for the item associated with the specified index.
*/
//go:nosplit
func (self class) SetItemTooltip(idx gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the tooltip hint associated with the specified index.
*/
//go:nosplit
func (self class) GetItemTooltip(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Select the item at the specified index.
[b]Note:[/b] This method does not trigger the item selection signal.
*/
//go:nosplit
func (self class) Select(idx gd.Int, single bool) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, single)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Ensures the item associated with the specified index is not selected.
*/
//go:nosplit
func (self class) Deselect(idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Ensures there are no items selected.
*/
//go:nosplit
func (self class) DeselectAll() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_deselect_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the item at the specified index is currently selected.
*/
//go:nosplit
func (self class) IsSelected(idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array with the indexes of the selected items.
*/
//go:nosplit
func (self class) GetSelectedItems() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_selected_items, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Moves item from index [param from_idx] to [param to_idx].
*/
//go:nosplit
func (self class) MoveItem(from_idx gd.Int, to_idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from_idx)
	callframe.Arg(frame, to_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_move_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetItemCount(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetItemCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the item specified by [param idx] index from the list.
*/
//go:nosplit
func (self class) RemoveItem(idx gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_remove_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all items from the list.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sorts items in the list by their text.
*/
//go:nosplit
func (self class) SortItemsByText() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_sort_items_by_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFixedColumnWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_fixed_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFixedColumnWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_fixed_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSameColumnWidth(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_same_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSameColumnWidth() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_is_same_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxTextLines(lines gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, lines)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_max_text_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxTextLines() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_max_text_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxColumns(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_max_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxColumns() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_max_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectMode(mode classdb.ItemListSelectMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_select_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSelectMode() classdb.ItemListSelectMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ItemListSelectMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_select_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIconMode(mode classdb.ItemListIconMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_icon_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIconMode() classdb.ItemListIconMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ItemListIconMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_icon_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFixedIconSize(size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_fixed_icon_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFixedIconSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_fixed_icon_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIconScale(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_icon_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIconScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_icon_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowRmbSelect(allow bool) {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_allow_rmb_select, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAllowRmbSelect() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_allow_rmb_select, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowReselect(allow bool) {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAllowReselect() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowSearch(allow bool) {
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAllowSearch() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoHeight(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_auto_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasAutoHeight() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_has_auto_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if one or more items are selected.
*/
//go:nosplit
func (self class) IsAnythingSelected() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_is_anything_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, exact)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_item_at_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Ensure current selection is visible, adjusting the scroll position as necessary.
*/
//go:nosplit
func (self class) EnsureCurrentIsVisible() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_ensure_current_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the vertical scrollbar.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetVScrollBar() objects.VScrollBar {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.VScrollBar{classdb.VScrollBar(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior classdb.TextServerOverrunBehavior) {
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Forces an update to the list size based on its items. This happens automatically whenever size of the items, or other relevant settings like [member auto_height], change. The method can be used to trigger the update ahead of next drawing pass.
*/
//go:nosplit
func (self class) ForceUpdateListSize() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ItemList.Bind_force_update_list_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnItemSelected(cb func(index int)) {
	self[0].AsObject().Connect(gd.NewStringName("item_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnEmptyClicked(cb func(at_position Vector2.XY, mouse_button_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("empty_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemClicked(cb func(index int, at_position Vector2.XY, mouse_button_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("item_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMultiSelected(cb func(index int, selected bool)) {
	self[0].AsObject().Connect(gd.NewStringName("multi_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemActivated(cb func(index int)) {
	self[0].AsObject().Connect(gd.NewStringName("item_activated"), gd.NewCallable(cb), 0)
}

func (self class) AsItemList() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsItemList() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() { classdb.Register("ItemList", func(ptr gd.Object) any { return classdb.ItemList(ptr) }) }

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
