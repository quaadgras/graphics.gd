package Tree

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
A control used to show a set of internal [TreeItem]s in a hierarchical structure. The tree items can be selected, expanded and collapsed. The tree can have multiple columns with custom controls like [LineEdit]s, buttons and popups. It can be useful for structured displays and interactions.
Trees are built via code, using [TreeItem] objects to create the structure. They have a single root, but multiple roots can be simulated with [member hide_root]:
[codeblocks]
[gdscript]
func _ready():
    var tree = Tree.new()
    var root = tree.create_item()
    tree.hide_root = true
    var child1 = tree.create_item(root)
    var child2 = tree.create_item(root)
    var subchild1 = tree.create_item(child1)
    subchild1.set_text(0, "Subchild1")
[/gdscript]
[csharp]
public override void _Ready()
{
    var tree = new Tree();
    TreeItem root = tree.CreateItem();
    tree.HideRoot = true;
    TreeItem child1 = tree.CreateItem(root);
    TreeItem child2 = tree.CreateItem(root);
    TreeItem subchild1 = tree.CreateItem(child1);
    subchild1.SetText(0, "Subchild1");
}
[/csharp]
[/codeblocks]
To iterate over all the [TreeItem] objects in a [Tree] object, use [method TreeItem.get_next] and [method TreeItem.get_first_child] after getting the root through [method get_root]. You can use [method Object.free] on a [TreeItem] to remove it from the [Tree].
[b]Incremental search:[/b] Like [ItemList] and [PopupMenu], [Tree] supports searching within the list while the control is focused. Press a key that matches the first letter of an item's name to select the first item starting with the given letter. After that point, there are two ways to perform incremental search: 1) Press the same key again before the timeout duration to select the next item starting with the same letter. 2) Press letter keys that match the rest of the word before the timeout duration to match to select the item in question directly. Both of these actions will be reset to the beginning of the list if the timeout duration has passed since the last keystroke was registered. You can adjust the timeout duration by changing [member ProjectSettings.gui/timers/incremental_search_max_interval_msec].

*/
type Simple [1]classdb.Tree
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) CreateItem(parent [1]classdb.TreeItem, index int) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).CreateItem(gc, parent, gd.Int(index)))
}
func (self Simple) GetRoot() [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetRoot(gc))
}
func (self Simple) SetColumnCustomMinimumWidth(column int, min_width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnCustomMinimumWidth(gd.Int(column), gd.Int(min_width))
}
func (self Simple) SetColumnExpand(column int, expand bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnExpand(gd.Int(column), expand)
}
func (self Simple) SetColumnExpandRatio(column int, ratio int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnExpandRatio(gd.Int(column), gd.Int(ratio))
}
func (self Simple) SetColumnClipContent(column int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnClipContent(gd.Int(column), enable)
}
func (self Simple) IsColumnExpanding(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsColumnExpanding(gd.Int(column)))
}
func (self Simple) IsColumnClippingContent(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsColumnClippingContent(gd.Int(column)))
}
func (self Simple) GetColumnExpandRatio(column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetColumnExpandRatio(gd.Int(column))))
}
func (self Simple) GetColumnWidth(column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetColumnWidth(gd.Int(column))))
}
func (self Simple) SetHideRoot(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHideRoot(enable)
}
func (self Simple) IsRootHidden() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRootHidden())
}
func (self Simple) GetNextSelected(from [1]classdb.TreeItem) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetNextSelected(gc, from))
}
func (self Simple) GetSelected() [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetSelected(gc))
}
func (self Simple) SetSelected(item [1]classdb.TreeItem, column int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelected(item, gd.Int(column))
}
func (self Simple) GetSelectedColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSelectedColumn()))
}
func (self Simple) GetPressedButton() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPressedButton()))
}
func (self Simple) SetSelectMode(mode classdb.TreeSelectMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelectMode(mode)
}
func (self Simple) GetSelectMode() classdb.TreeSelectMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TreeSelectMode(Expert(self).GetSelectMode())
}
func (self Simple) DeselectAll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DeselectAll()
}
func (self Simple) SetColumns(amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumns(gd.Int(amount))
}
func (self Simple) GetColumns() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetColumns()))
}
func (self Simple) GetEdited() [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetEdited(gc))
}
func (self Simple) GetEditedColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetEditedColumn()))
}
func (self Simple) EditSelected(force_edit bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).EditSelected(force_edit))
}
func (self Simple) GetCustomPopupRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetCustomPopupRect())
}
func (self Simple) GetItemAreaRect(item [1]classdb.TreeItem, column int, button_index int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetItemAreaRect(item, gd.Int(column), gd.Int(button_index)))
}
func (self Simple) GetItemAtPosition(position gd.Vector2) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetItemAtPosition(gc, position))
}
func (self Simple) GetColumnAtPosition(position gd.Vector2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetColumnAtPosition(position)))
}
func (self Simple) GetDropSectionAtPosition(position gd.Vector2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDropSectionAtPosition(position)))
}
func (self Simple) GetButtonIdAtPosition(position gd.Vector2) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetButtonIdAtPosition(position)))
}
func (self Simple) EnsureCursorIsVisible() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EnsureCursorIsVisible()
}
func (self Simple) SetColumnTitlesVisible(visible bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnTitlesVisible(visible)
}
func (self Simple) AreColumnTitlesVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).AreColumnTitlesVisible())
}
func (self Simple) SetColumnTitle(column int, title string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnTitle(gd.Int(column), gc.String(title))
}
func (self Simple) GetColumnTitle(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetColumnTitle(gc, gd.Int(column)).String())
}
func (self Simple) SetColumnTitleAlignment(column int, title_alignment gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnTitleAlignment(gd.Int(column), title_alignment)
}
func (self Simple) GetColumnTitleAlignment(column int) gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
	return gd.HorizontalAlignment(Expert(self).GetColumnTitleAlignment(gd.Int(column)))
}
func (self Simple) SetColumnTitleDirection(column int, direction classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnTitleDirection(gd.Int(column), direction)
}
func (self Simple) GetColumnTitleDirection(column int) classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlTextDirection(Expert(self).GetColumnTitleDirection(gd.Int(column)))
}
func (self Simple) SetColumnTitleLanguage(column int, language string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColumnTitleLanguage(gd.Int(column), gc.String(language))
}
func (self Simple) GetColumnTitleLanguage(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetColumnTitleLanguage(gc, gd.Int(column)).String())
}
func (self Simple) GetScroll() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetScroll())
}
func (self Simple) ScrollToItem(item [1]classdb.TreeItem, center_on_item bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScrollToItem(item, center_on_item)
}
func (self Simple) SetHScrollEnabled(h_scroll bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHScrollEnabled(h_scroll)
}
func (self Simple) IsHScrollEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHScrollEnabled())
}
func (self Simple) SetVScrollEnabled(h_scroll bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVScrollEnabled(h_scroll)
}
func (self Simple) IsVScrollEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVScrollEnabled())
}
func (self Simple) SetHideFolding(hide bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHideFolding(hide)
}
func (self Simple) IsFoldingHidden() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFoldingHidden())
}
func (self Simple) SetEnableRecursiveFolding(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnableRecursiveFolding(enable)
}
func (self Simple) IsRecursiveFoldingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRecursiveFoldingEnabled())
}
func (self Simple) SetDropModeFlags(flags int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDropModeFlags(gd.Int(flags))
}
func (self Simple) GetDropModeFlags() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDropModeFlags()))
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
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Tree
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Clears the tree. This removes all items.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates an item in the tree and adds it as a child of [param parent], which can be either a valid [TreeItem] or [code]null[/code].
If [param parent] is [code]null[/code], the root item will be the parent, or the new item will be the root itself if the tree is empty.
The new item will be the [param index]-th child of parent, or it will be the last child if there are not enough siblings.
*/
//go:nosplit
func (self class) CreateItem(ctx gd.Lifetime, parent object.TreeItem, index gd.Int) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(parent[0].AsPointer())[0])
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_create_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the tree's root item, or [code]null[/code] if the tree is empty.
*/
//go:nosplit
func (self class) GetRoot(ctx gd.Lifetime) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Overrides the calculated minimum width of a column. It can be set to [code]0[/code] to restore the default behavior. Columns that have the "Expand" flag will use their "min_width" in a similar fashion to [member Control.size_flags_stretch_ratio].
*/
//go:nosplit
func (self class) SetColumnCustomMinimumWidth(column gd.Int, min_width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, min_width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_custom_minimum_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], the column will have the "Expand" flag of [Control]. Columns that have the "Expand" flag will use their expand ratio in a similar fashion to [member Control.size_flags_stretch_ratio] (see [method set_column_expand_ratio]).
*/
//go:nosplit
func (self class) SetColumnExpand(column gd.Int, expand bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, expand)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_expand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the relative expand ratio for a column. See [method set_column_expand].
*/
//go:nosplit
func (self class) SetColumnExpandRatio(column gd.Int, ratio gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_expand_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Allows to enable clipping for column's content, making the content size ignored.
*/
//go:nosplit
func (self class) SetColumnClipContent(column gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_clip_content, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the column has enabled expanding (see [method set_column_expand]).
*/
//go:nosplit
func (self class) IsColumnExpanding(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_is_column_expanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the column has enabled clipping (see [method set_column_clip_content]).
*/
//go:nosplit
func (self class) IsColumnClippingContent(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_is_column_clipping_content, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the expand ratio assigned to the column.
*/
//go:nosplit
func (self class) GetColumnExpandRatio(column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_column_expand_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the column's width in pixels.
*/
//go:nosplit
func (self class) GetColumnWidth(column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_column_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHideRoot(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_hide_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRootHidden() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_is_root_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the next selected [TreeItem] after the given one, or [code]null[/code] if the end is reached.
If [param from] is [code]null[/code], this returns the first selected item.
*/
//go:nosplit
func (self class) GetNextSelected(ctx gd.Lifetime, from object.TreeItem) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(from[0].AsPointer())[0])
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_next_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the currently focused item, or [code]null[/code] if no item is focused.
In [constant SELECT_ROW] and [constant SELECT_SINGLE] modes, the focused item is same as the selected item. In [constant SELECT_MULTI] mode, the focused item is the item under the focus cursor, not necessarily selected.
To get the currently selected item(s), use [method get_next_selected].
*/
//go:nosplit
func (self class) GetSelected(ctx gd.Lifetime) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Selects the specified [TreeItem] and column.
*/
//go:nosplit
func (self class) SetSelected(item object.TreeItem, column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(item[0].AsPointer())[0])
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the currently focused column, or -1 if no column is focused.
In [constant SELECT_SINGLE] mode, the focused column is the selected column. In [constant SELECT_ROW] mode, the focused column is always 0 if any item is selected. In [constant SELECT_MULTI] mode, the focused column is the column under the focus cursor, and there are not necessarily any column selected.
To tell whether a column of an item is selected, use [method TreeItem.is_selected].
*/
//go:nosplit
func (self class) GetSelectedColumn() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_selected_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the last pressed button's index.
*/
//go:nosplit
func (self class) GetPressedButton() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_pressed_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectMode(mode classdb.TreeSelectMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_select_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSelectMode() classdb.TreeSelectMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TreeSelectMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_select_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Deselects all tree items (rows and columns). In [constant SELECT_MULTI] mode also removes selection cursor.
*/
//go:nosplit
func (self class) DeselectAll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_deselect_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetColumns(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColumns() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_columns, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the currently edited item. Can be used with [signal item_edited] to get the item that was modified.
[codeblocks]
[gdscript]
func _ready():
    $Tree.item_edited.connect(on_Tree_item_edited)

func on_Tree_item_edited():
    print($Tree.get_edited()) # This item just got edited (e.g. checked).
[/gdscript]
[csharp]
public override void _Ready()
{
    GetNode<Tree>("Tree").ItemEdited += OnTreeItemEdited;
}

public void OnTreeItemEdited()
{
    GD.Print(GetNode<Tree>("Tree").GetEdited()); // This item just got edited (e.g. checked).
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetEdited(ctx gd.Lifetime) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_edited, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the column for the currently edited item.
*/
//go:nosplit
func (self class) GetEditedColumn() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_edited_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Edits the selected tree item as if it was clicked.
Either the item must be set editable with [method TreeItem.set_editable] or [param force_edit] must be [code]true[/code].
Returns [code]true[/code] if the item could be edited. Fails if no item is selected.
*/
//go:nosplit
func (self class) EditSelected(force_edit bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force_edit)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_edit_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the rectangle for custom popups. Helper to create custom cell controls that display a popup. See [method TreeItem.set_cell_mode].
*/
//go:nosplit
func (self class) GetCustomPopupRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_custom_popup_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the rectangle area for the specified [TreeItem]. If [param column] is specified, only get the position and size of that column, otherwise get the rectangle containing all columns. If a button index is specified, the rectangle of that button will be returned.
*/
//go:nosplit
func (self class) GetItemAreaRect(item object.TreeItem, column gd.Int, button_index gd.Int) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(item[0].AsPointer())[0])
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_item_area_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tree item at the specified position (relative to the tree origin position).
*/
//go:nosplit
func (self class) GetItemAtPosition(ctx gd.Lifetime, position gd.Vector2) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_item_at_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the column index at [param position], or -1 if no item is there.
*/
//go:nosplit
func (self class) GetColumnAtPosition(position gd.Vector2) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_column_at_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the drop section at [param position], or -100 if no item is there.
Values -1, 0, or 1 will be returned for the "above item", "on item", and "below item" drop sections, respectively. See [enum DropModeFlags] for a description of each drop section.
To get the item which the returned drop section is relative to, use [method get_item_at_position].
*/
//go:nosplit
func (self class) GetDropSectionAtPosition(position gd.Vector2) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_drop_section_at_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the button ID at [param position], or -1 if no button is there.
*/
//go:nosplit
func (self class) GetButtonIdAtPosition(position gd.Vector2) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_button_id_at_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Makes the currently focused cell visible.
This will scroll the tree if necessary. In [constant SELECT_ROW] mode, this will not do horizontal scrolling, as all the cells in the selected row is focused logically.
[b]Note:[/b] Despite the name of this method, the focus cursor itself is only visible in [constant SELECT_MULTI] mode.
*/
//go:nosplit
func (self class) EnsureCursorIsVisible()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_ensure_cursor_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetColumnTitlesVisible(visible bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_titles_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) AreColumnTitlesVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_are_column_titles_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the title of a column.
*/
//go:nosplit
func (self class) SetColumnTitle(column gd.Int, title gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(title))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the column's title.
*/
//go:nosplit
func (self class) GetColumnTitle(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_column_title, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the column title alignment. Note that [constant @GlobalScope.HORIZONTAL_ALIGNMENT_FILL] is not supported for column titles.
*/
//go:nosplit
func (self class) SetColumnTitleAlignment(column gd.Int, title_alignment gd.HorizontalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, title_alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_title_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the column title alignment.
*/
//go:nosplit
func (self class) GetColumnTitleAlignment(column gd.Int) gd.HorizontalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_column_title_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets column title base writing direction.
*/
//go:nosplit
func (self class) SetColumnTitleDirection(column gd.Int, direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_title_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns column title base writing direction.
*/
//go:nosplit
func (self class) GetColumnTitleDirection(column gd.Int) classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_column_title_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets language code of column title used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetColumnTitleLanguage(column gd.Int, language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_column_title_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns column title language code.
*/
//go:nosplit
func (self class) GetColumnTitleLanguage(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_column_title_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the current scrolling position.
*/
//go:nosplit
func (self class) GetScroll() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Causes the [Tree] to jump to the specified [TreeItem].
*/
//go:nosplit
func (self class) ScrollToItem(item object.TreeItem, center_on_item bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(item[0].AsPointer())[0])
	callframe.Arg(frame, center_on_item)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_scroll_to_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetHScrollEnabled(h_scroll bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, h_scroll)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_h_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHScrollEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_is_h_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVScrollEnabled(h_scroll bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, h_scroll)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_v_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVScrollEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_is_v_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHideFolding(hide bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hide)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_hide_folding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFoldingHidden() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_is_folding_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnableRecursiveFolding(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_enable_recursive_folding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRecursiveFoldingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_is_recursive_folding_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDropModeFlags(flags gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_drop_mode_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDropModeFlags() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_drop_mode_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_allow_rmb_select, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowRmbSelect() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_allow_rmb_select, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowReselect() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_set_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAllowSearch() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Tree.Bind_get_allow_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTree() Expert { return self[0].AsTree() }


//go:nosplit
func (self Simple) AsTree() Simple { return self[0].AsTree() }


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
func init() {classdb.Register("Tree", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type SelectMode = classdb.TreeSelectMode

const (
/*Allows selection of a single cell at a time. From the perspective of items, only a single item is allowed to be selected. And there is only one column selected in the selected item.
The focus cursor is always hidden in this mode, but it is positioned at the current selection, making the currently selected item the currently focused item.*/
	SelectSingle SelectMode = 0
/*Allows selection of a single row at a time. From the perspective of items, only a single items is allowed to be selected. And all the columns are selected in the selected item.
The focus cursor is always hidden in this mode, but it is positioned at the first column of the current selection, making the currently selected item the currently focused item.*/
	SelectRow SelectMode = 1
/*Allows selection of multiple cells at the same time. From the perspective of items, multiple items are allowed to be selected. And there can be multiple columns selected in each selected item.
The focus cursor is visible in this mode, the item or column under the cursor is not necessarily selected.*/
	SelectMulti SelectMode = 2
)
type DropModeFlags = classdb.TreeDropModeFlags

const (
/*Disables all drop sections, but still allows to detect the "on item" drop section by [method get_drop_section_at_position].
[b]Note:[/b] This is the default flag, it has no effect when combined with other flags.*/
	DropModeDisabled DropModeFlags = 0
/*Enables the "on item" drop section. This drop section covers the entire item.
When combined with [constant DROP_MODE_INBETWEEN], this drop section halves the height and stays centered vertically.*/
	DropModeOnItem DropModeFlags = 1
/*Enables "above item" and "below item" drop sections. The "above item" drop section covers the top half of the item, and the "below item" drop section covers the bottom half.
When combined with [constant DROP_MODE_ON_ITEM], these drop sections halves the height and stays on top / bottom accordingly.*/
	DropModeInbetween DropModeFlags = 2
)
