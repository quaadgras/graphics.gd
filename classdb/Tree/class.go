// Package Tree provides methods for working with Tree object instances.
package Tree

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Vector2"

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
type Instance [1]gdclass.Tree

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTree() Instance
}

/*
Clears the tree. This removes all items.
*/
func (self Instance) Clear() { //gd:Tree.clear
	class(self).Clear()
}

/*
Creates an item in the tree and adds it as a child of [param parent], which can be either a valid [TreeItem] or [code]null[/code].
If [param parent] is [code]null[/code], the root item will be the parent, or the new item will be the root itself if the tree is empty.
The new item will be the [param index]-th child of parent, or it will be the last child if there are not enough siblings.
*/
func (self Instance) CreateItem() [1]gdclass.TreeItem { //gd:Tree.create_item
	return [1]gdclass.TreeItem(class(self).CreateItem([1][1]gdclass.TreeItem{}[0], gd.Int(-1)))
}

/*
Returns the tree's root item, or [code]null[/code] if the tree is empty.
*/
func (self Instance) GetRoot() [1]gdclass.TreeItem { //gd:Tree.get_root
	return [1]gdclass.TreeItem(class(self).GetRoot())
}

/*
Overrides the calculated minimum width of a column. It can be set to [code]0[/code] to restore the default behavior. Columns that have the "Expand" flag will use their "min_width" in a similar fashion to [member Control.size_flags_stretch_ratio].
*/
func (self Instance) SetColumnCustomMinimumWidth(column int, min_width int) { //gd:Tree.set_column_custom_minimum_width
	class(self).SetColumnCustomMinimumWidth(gd.Int(column), gd.Int(min_width))
}

/*
If [code]true[/code], the column will have the "Expand" flag of [Control]. Columns that have the "Expand" flag will use their expand ratio in a similar fashion to [member Control.size_flags_stretch_ratio] (see [method set_column_expand_ratio]).
*/
func (self Instance) SetColumnExpand(column int, expand bool) { //gd:Tree.set_column_expand
	class(self).SetColumnExpand(gd.Int(column), expand)
}

/*
Sets the relative expand ratio for a column. See [method set_column_expand].
*/
func (self Instance) SetColumnExpandRatio(column int, ratio int) { //gd:Tree.set_column_expand_ratio
	class(self).SetColumnExpandRatio(gd.Int(column), gd.Int(ratio))
}

/*
Allows to enable clipping for column's content, making the content size ignored.
*/
func (self Instance) SetColumnClipContent(column int, enable bool) { //gd:Tree.set_column_clip_content
	class(self).SetColumnClipContent(gd.Int(column), enable)
}

/*
Returns [code]true[/code] if the column has enabled expanding (see [method set_column_expand]).
*/
func (self Instance) IsColumnExpanding(column int) bool { //gd:Tree.is_column_expanding
	return bool(class(self).IsColumnExpanding(gd.Int(column)))
}

/*
Returns [code]true[/code] if the column has enabled clipping (see [method set_column_clip_content]).
*/
func (self Instance) IsColumnClippingContent(column int) bool { //gd:Tree.is_column_clipping_content
	return bool(class(self).IsColumnClippingContent(gd.Int(column)))
}

/*
Returns the expand ratio assigned to the column.
*/
func (self Instance) GetColumnExpandRatio(column int) int { //gd:Tree.get_column_expand_ratio
	return int(int(class(self).GetColumnExpandRatio(gd.Int(column))))
}

/*
Returns the column's width in pixels.
*/
func (self Instance) GetColumnWidth(column int) int { //gd:Tree.get_column_width
	return int(int(class(self).GetColumnWidth(gd.Int(column))))
}

/*
Returns the next selected [TreeItem] after the given one, or [code]null[/code] if the end is reached.
If [param from] is [code]null[/code], this returns the first selected item.
*/
func (self Instance) GetNextSelected(from [1]gdclass.TreeItem) [1]gdclass.TreeItem { //gd:Tree.get_next_selected
	return [1]gdclass.TreeItem(class(self).GetNextSelected(from))
}

/*
Returns the currently focused item, or [code]null[/code] if no item is focused.
In [constant SELECT_ROW] and [constant SELECT_SINGLE] modes, the focused item is same as the selected item. In [constant SELECT_MULTI] mode, the focused item is the item under the focus cursor, not necessarily selected.
To get the currently selected item(s), use [method get_next_selected].
*/
func (self Instance) GetSelected() [1]gdclass.TreeItem { //gd:Tree.get_selected
	return [1]gdclass.TreeItem(class(self).GetSelected())
}

/*
Selects the specified [TreeItem] and column.
*/
func (self Instance) SetSelected(item [1]gdclass.TreeItem, column int) { //gd:Tree.set_selected
	class(self).SetSelected(item, gd.Int(column))
}

/*
Returns the currently focused column, or -1 if no column is focused.
In [constant SELECT_SINGLE] mode, the focused column is the selected column. In [constant SELECT_ROW] mode, the focused column is always 0 if any item is selected. In [constant SELECT_MULTI] mode, the focused column is the column under the focus cursor, and there are not necessarily any column selected.
To tell whether a column of an item is selected, use [method TreeItem.is_selected].
*/
func (self Instance) GetSelectedColumn() int { //gd:Tree.get_selected_column
	return int(int(class(self).GetSelectedColumn()))
}

/*
Returns the last pressed button's index.
*/
func (self Instance) GetPressedButton() int { //gd:Tree.get_pressed_button
	return int(int(class(self).GetPressedButton()))
}

/*
Deselects all tree items (rows and columns). In [constant SELECT_MULTI] mode also removes selection cursor.
*/
func (self Instance) DeselectAll() { //gd:Tree.deselect_all
	class(self).DeselectAll()
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
func (self Instance) GetEdited() [1]gdclass.TreeItem { //gd:Tree.get_edited
	return [1]gdclass.TreeItem(class(self).GetEdited())
}

/*
Returns the column for the currently edited item.
*/
func (self Instance) GetEditedColumn() int { //gd:Tree.get_edited_column
	return int(int(class(self).GetEditedColumn()))
}

/*
Edits the selected tree item as if it was clicked.
Either the item must be set editable with [method TreeItem.set_editable] or [param force_edit] must be [code]true[/code].
Returns [code]true[/code] if the item could be edited. Fails if no item is selected.
*/
func (self Instance) EditSelected() bool { //gd:Tree.edit_selected
	return bool(class(self).EditSelected(false))
}

/*
Returns the rectangle for custom popups. Helper to create custom cell controls that display a popup. See [method TreeItem.set_cell_mode].
*/
func (self Instance) GetCustomPopupRect() Rect2.PositionSize { //gd:Tree.get_custom_popup_rect
	return Rect2.PositionSize(class(self).GetCustomPopupRect())
}

/*
Returns the rectangle area for the specified [TreeItem]. If [param column] is specified, only get the position and size of that column, otherwise get the rectangle containing all columns. If a button index is specified, the rectangle of that button will be returned.
*/
func (self Instance) GetItemAreaRect(item [1]gdclass.TreeItem) Rect2.PositionSize { //gd:Tree.get_item_area_rect
	return Rect2.PositionSize(class(self).GetItemAreaRect(item, gd.Int(-1), gd.Int(-1)))
}

/*
Returns the tree item at the specified position (relative to the tree origin position).
*/
func (self Instance) GetItemAtPosition(position Vector2.XY) [1]gdclass.TreeItem { //gd:Tree.get_item_at_position
	return [1]gdclass.TreeItem(class(self).GetItemAtPosition(gd.Vector2(position)))
}

/*
Returns the column index at [param position], or -1 if no item is there.
*/
func (self Instance) GetColumnAtPosition(position Vector2.XY) int { //gd:Tree.get_column_at_position
	return int(int(class(self).GetColumnAtPosition(gd.Vector2(position))))
}

/*
Returns the drop section at [param position], or -100 if no item is there.
Values -1, 0, or 1 will be returned for the "above item", "on item", and "below item" drop sections, respectively. See [enum DropModeFlags] for a description of each drop section.
To get the item which the returned drop section is relative to, use [method get_item_at_position].
*/
func (self Instance) GetDropSectionAtPosition(position Vector2.XY) int { //gd:Tree.get_drop_section_at_position
	return int(int(class(self).GetDropSectionAtPosition(gd.Vector2(position))))
}

/*
Returns the button ID at [param position], or -1 if no button is there.
*/
func (self Instance) GetButtonIdAtPosition(position Vector2.XY) int { //gd:Tree.get_button_id_at_position
	return int(int(class(self).GetButtonIdAtPosition(gd.Vector2(position))))
}

/*
Makes the currently focused cell visible.
This will scroll the tree if necessary. In [constant SELECT_ROW] mode, this will not do horizontal scrolling, as all the cells in the selected row is focused logically.
[b]Note:[/b] Despite the name of this method, the focus cursor itself is only visible in [constant SELECT_MULTI] mode.
*/
func (self Instance) EnsureCursorIsVisible() { //gd:Tree.ensure_cursor_is_visible
	class(self).EnsureCursorIsVisible()
}

/*
Sets the title of a column.
*/
func (self Instance) SetColumnTitle(column int, title string) { //gd:Tree.set_column_title
	class(self).SetColumnTitle(gd.Int(column), String.New(title))
}

/*
Returns the column's title.
*/
func (self Instance) GetColumnTitle(column int) string { //gd:Tree.get_column_title
	return string(class(self).GetColumnTitle(gd.Int(column)).String())
}

/*
Sets the column title alignment. Note that [constant @GlobalScope.HORIZONTAL_ALIGNMENT_FILL] is not supported for column titles.
*/
func (self Instance) SetColumnTitleAlignment(column int, title_alignment HorizontalAlignment) { //gd:Tree.set_column_title_alignment
	class(self).SetColumnTitleAlignment(gd.Int(column), title_alignment)
}

/*
Returns the column title alignment.
*/
func (self Instance) GetColumnTitleAlignment(column int) HorizontalAlignment { //gd:Tree.get_column_title_alignment
	return HorizontalAlignment(class(self).GetColumnTitleAlignment(gd.Int(column)))
}

/*
Sets column title base writing direction.
*/
func (self Instance) SetColumnTitleDirection(column int, direction gdclass.ControlTextDirection) { //gd:Tree.set_column_title_direction
	class(self).SetColumnTitleDirection(gd.Int(column), direction)
}

/*
Returns column title base writing direction.
*/
func (self Instance) GetColumnTitleDirection(column int) gdclass.ControlTextDirection { //gd:Tree.get_column_title_direction
	return gdclass.ControlTextDirection(class(self).GetColumnTitleDirection(gd.Int(column)))
}

/*
Sets language code of column title used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Instance) SetColumnTitleLanguage(column int, language string) { //gd:Tree.set_column_title_language
	class(self).SetColumnTitleLanguage(gd.Int(column), String.New(language))
}

/*
Returns column title language code.
*/
func (self Instance) GetColumnTitleLanguage(column int) string { //gd:Tree.get_column_title_language
	return string(class(self).GetColumnTitleLanguage(gd.Int(column)).String())
}

/*
Returns the current scrolling position.
*/
func (self Instance) GetScroll() Vector2.XY { //gd:Tree.get_scroll
	return Vector2.XY(class(self).GetScroll())
}

/*
Causes the [Tree] to jump to the specified [TreeItem].
*/
func (self Instance) ScrollToItem(item [1]gdclass.TreeItem) { //gd:Tree.scroll_to_item
	class(self).ScrollToItem(item, false)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Tree

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Tree"))
	casted := Instance{*(*gdclass.Tree)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Columns() int {
	return int(int(class(self).GetColumns()))
}

func (self Instance) SetColumns(value int) {
	class(self).SetColumns(gd.Int(value))
}

func (self Instance) ColumnTitlesVisible() bool {
	return bool(class(self).AreColumnTitlesVisible())
}

func (self Instance) SetColumnTitlesVisible(value bool) {
	class(self).SetColumnTitlesVisible(value)
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

func (self Instance) HideFolding() bool {
	return bool(class(self).IsFoldingHidden())
}

func (self Instance) SetHideFolding(value bool) {
	class(self).SetHideFolding(value)
}

func (self Instance) EnableRecursiveFolding() bool {
	return bool(class(self).IsRecursiveFoldingEnabled())
}

func (self Instance) SetEnableRecursiveFolding(value bool) {
	class(self).SetEnableRecursiveFolding(value)
}

func (self Instance) HideRoot() bool {
	return bool(class(self).IsRootHidden())
}

func (self Instance) SetHideRoot(value bool) {
	class(self).SetHideRoot(value)
}

func (self Instance) DropModeFlags() int {
	return int(int(class(self).GetDropModeFlags()))
}

func (self Instance) SetDropModeFlags(value int) {
	class(self).SetDropModeFlags(gd.Int(value))
}

func (self Instance) SelectMode() gdclass.TreeSelectMode {
	return gdclass.TreeSelectMode(class(self).GetSelectMode())
}

func (self Instance) SetSelectMode(value gdclass.TreeSelectMode) {
	class(self).SetSelectMode(value)
}

func (self Instance) ScrollHorizontalEnabled() bool {
	return bool(class(self).IsHScrollEnabled())
}

func (self Instance) SetScrollHorizontalEnabled(value bool) {
	class(self).SetHScrollEnabled(value)
}

func (self Instance) ScrollVerticalEnabled() bool {
	return bool(class(self).IsVScrollEnabled())
}

func (self Instance) SetScrollVerticalEnabled(value bool) {
	class(self).SetVScrollEnabled(value)
}

/*
Clears the tree. This removes all items.
*/
//go:nosplit
func (self class) Clear() { //gd:Tree.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates an item in the tree and adds it as a child of [param parent], which can be either a valid [TreeItem] or [code]null[/code].
If [param parent] is [code]null[/code], the root item will be the parent, or the new item will be the root itself if the tree is empty.
The new item will be the [param index]-th child of parent, or it will be the last child if there are not enough siblings.
*/
//go:nosplit
func (self class) CreateItem(parent [1]gdclass.TreeItem, index gd.Int) [1]gdclass.TreeItem { //gd:Tree.create_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parent[0])[0])
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_create_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the tree's root item, or [code]null[/code] if the tree is empty.
*/
//go:nosplit
func (self class) GetRoot() [1]gdclass.TreeItem { //gd:Tree.get_root
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Overrides the calculated minimum width of a column. It can be set to [code]0[/code] to restore the default behavior. Columns that have the "Expand" flag will use their "min_width" in a similar fashion to [member Control.size_flags_stretch_ratio].
*/
//go:nosplit
func (self class) SetColumnCustomMinimumWidth(column gd.Int, min_width gd.Int) { //gd:Tree.set_column_custom_minimum_width
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, min_width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_custom_minimum_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [code]true[/code], the column will have the "Expand" flag of [Control]. Columns that have the "Expand" flag will use their expand ratio in a similar fashion to [member Control.size_flags_stretch_ratio] (see [method set_column_expand_ratio]).
*/
//go:nosplit
func (self class) SetColumnExpand(column gd.Int, expand bool) { //gd:Tree.set_column_expand
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, expand)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_expand, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the relative expand ratio for a column. See [method set_column_expand].
*/
//go:nosplit
func (self class) SetColumnExpandRatio(column gd.Int, ratio gd.Int) { //gd:Tree.set_column_expand_ratio
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_expand_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Allows to enable clipping for column's content, making the content size ignored.
*/
//go:nosplit
func (self class) SetColumnClipContent(column gd.Int, enable bool) { //gd:Tree.set_column_clip_content
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_clip_content, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the column has enabled expanding (see [method set_column_expand]).
*/
//go:nosplit
func (self class) IsColumnExpanding(column gd.Int) bool { //gd:Tree.is_column_expanding
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_is_column_expanding, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the column has enabled clipping (see [method set_column_clip_content]).
*/
//go:nosplit
func (self class) IsColumnClippingContent(column gd.Int) bool { //gd:Tree.is_column_clipping_content
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_is_column_clipping_content, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the expand ratio assigned to the column.
*/
//go:nosplit
func (self class) GetColumnExpandRatio(column gd.Int) gd.Int { //gd:Tree.get_column_expand_ratio
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_column_expand_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the column's width in pixels.
*/
//go:nosplit
func (self class) GetColumnWidth(column gd.Int) gd.Int { //gd:Tree.get_column_width
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_column_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHideRoot(enable bool) { //gd:Tree.set_hide_root
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_hide_root, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRootHidden() bool { //gd:Tree.is_root_hidden
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_is_root_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next selected [TreeItem] after the given one, or [code]null[/code] if the end is reached.
If [param from] is [code]null[/code], this returns the first selected item.
*/
//go:nosplit
func (self class) GetNextSelected(from [1]gdclass.TreeItem) [1]gdclass.TreeItem { //gd:Tree.get_next_selected
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_next_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the currently focused item, or [code]null[/code] if no item is focused.
In [constant SELECT_ROW] and [constant SELECT_SINGLE] modes, the focused item is same as the selected item. In [constant SELECT_MULTI] mode, the focused item is the item under the focus cursor, not necessarily selected.
To get the currently selected item(s), use [method get_next_selected].
*/
//go:nosplit
func (self class) GetSelected() [1]gdclass.TreeItem { //gd:Tree.get_selected
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Selects the specified [TreeItem] and column.
*/
//go:nosplit
func (self class) SetSelected(item [1]gdclass.TreeItem, column gd.Int) { //gd:Tree.set_selected
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	callframe.Arg(frame, column)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the currently focused column, or -1 if no column is focused.
In [constant SELECT_SINGLE] mode, the focused column is the selected column. In [constant SELECT_ROW] mode, the focused column is always 0 if any item is selected. In [constant SELECT_MULTI] mode, the focused column is the column under the focus cursor, and there are not necessarily any column selected.
To tell whether a column of an item is selected, use [method TreeItem.is_selected].
*/
//go:nosplit
func (self class) GetSelectedColumn() gd.Int { //gd:Tree.get_selected_column
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_selected_column, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the last pressed button's index.
*/
//go:nosplit
func (self class) GetPressedButton() gd.Int { //gd:Tree.get_pressed_button
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_pressed_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectMode(mode gdclass.TreeSelectMode) { //gd:Tree.set_select_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_select_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSelectMode() gdclass.TreeSelectMode { //gd:Tree.get_select_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TreeSelectMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_select_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Deselects all tree items (rows and columns). In [constant SELECT_MULTI] mode also removes selection cursor.
*/
//go:nosplit
func (self class) DeselectAll() { //gd:Tree.deselect_all
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_deselect_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetColumns(amount gd.Int) { //gd:Tree.set_columns
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_columns, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColumns() gd.Int { //gd:Tree.get_columns
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_columns, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) GetEdited() [1]gdclass.TreeItem { //gd:Tree.get_edited
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_edited, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the column for the currently edited item.
*/
//go:nosplit
func (self class) GetEditedColumn() gd.Int { //gd:Tree.get_edited_column
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_edited_column, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) EditSelected(force_edit bool) bool { //gd:Tree.edit_selected
	var frame = callframe.New()
	callframe.Arg(frame, force_edit)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_edit_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the rectangle for custom popups. Helper to create custom cell controls that display a popup. See [method TreeItem.set_cell_mode].
*/
//go:nosplit
func (self class) GetCustomPopupRect() gd.Rect2 { //gd:Tree.get_custom_popup_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_custom_popup_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the rectangle area for the specified [TreeItem]. If [param column] is specified, only get the position and size of that column, otherwise get the rectangle containing all columns. If a button index is specified, the rectangle of that button will be returned.
*/
//go:nosplit
func (self class) GetItemAreaRect(item [1]gdclass.TreeItem, column gd.Int, button_index gd.Int) gd.Rect2 { //gd:Tree.get_item_area_rect
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_item_area_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tree item at the specified position (relative to the tree origin position).
*/
//go:nosplit
func (self class) GetItemAtPosition(position gd.Vector2) [1]gdclass.TreeItem { //gd:Tree.get_item_at_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_item_at_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the column index at [param position], or -1 if no item is there.
*/
//go:nosplit
func (self class) GetColumnAtPosition(position gd.Vector2) gd.Int { //gd:Tree.get_column_at_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_column_at_position, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) GetDropSectionAtPosition(position gd.Vector2) gd.Int { //gd:Tree.get_drop_section_at_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_drop_section_at_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the button ID at [param position], or -1 if no button is there.
*/
//go:nosplit
func (self class) GetButtonIdAtPosition(position gd.Vector2) gd.Int { //gd:Tree.get_button_id_at_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_button_id_at_position, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) EnsureCursorIsVisible() { //gd:Tree.ensure_cursor_is_visible
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_ensure_cursor_is_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetColumnTitlesVisible(visible bool) { //gd:Tree.set_column_titles_visible
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_titles_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) AreColumnTitlesVisible() bool { //gd:Tree.are_column_titles_visible
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_are_column_titles_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the title of a column.
*/
//go:nosplit
func (self class) SetColumnTitle(column gd.Int, title String.Readable) { //gd:Tree.set_column_title
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the column's title.
*/
//go:nosplit
func (self class) GetColumnTitle(column gd.Int) String.Readable { //gd:Tree.get_column_title
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_column_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the column title alignment. Note that [constant @GlobalScope.HORIZONTAL_ALIGNMENT_FILL] is not supported for column titles.
*/
//go:nosplit
func (self class) SetColumnTitleAlignment(column gd.Int, title_alignment HorizontalAlignment) { //gd:Tree.set_column_title_alignment
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, title_alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_title_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the column title alignment.
*/
//go:nosplit
func (self class) GetColumnTitleAlignment(column gd.Int) HorizontalAlignment { //gd:Tree.get_column_title_alignment
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_column_title_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets column title base writing direction.
*/
//go:nosplit
func (self class) SetColumnTitleDirection(column gd.Int, direction gdclass.ControlTextDirection) { //gd:Tree.set_column_title_direction
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_title_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns column title base writing direction.
*/
//go:nosplit
func (self class) GetColumnTitleDirection(column gd.Int) gdclass.ControlTextDirection { //gd:Tree.get_column_title_direction
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gdclass.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_column_title_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets language code of column title used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetColumnTitleLanguage(column gd.Int, language String.Readable) { //gd:Tree.set_column_title_language
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_column_title_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns column title language code.
*/
//go:nosplit
func (self class) GetColumnTitleLanguage(column gd.Int) String.Readable { //gd:Tree.get_column_title_language
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_column_title_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the current scrolling position.
*/
//go:nosplit
func (self class) GetScroll() gd.Vector2 { //gd:Tree.get_scroll
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_scroll, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Causes the [Tree] to jump to the specified [TreeItem].
*/
//go:nosplit
func (self class) ScrollToItem(item [1]gdclass.TreeItem, center_on_item bool) { //gd:Tree.scroll_to_item
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	callframe.Arg(frame, center_on_item)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_scroll_to_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetHScrollEnabled(h_scroll bool) { //gd:Tree.set_h_scroll_enabled
	var frame = callframe.New()
	callframe.Arg(frame, h_scroll)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_h_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHScrollEnabled() bool { //gd:Tree.is_h_scroll_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_is_h_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVScrollEnabled(h_scroll bool) { //gd:Tree.set_v_scroll_enabled
	var frame = callframe.New()
	callframe.Arg(frame, h_scroll)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_v_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVScrollEnabled() bool { //gd:Tree.is_v_scroll_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_is_v_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHideFolding(hide bool) { //gd:Tree.set_hide_folding
	var frame = callframe.New()
	callframe.Arg(frame, hide)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_hide_folding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFoldingHidden() bool { //gd:Tree.is_folding_hidden
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_is_folding_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableRecursiveFolding(enable bool) { //gd:Tree.set_enable_recursive_folding
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_enable_recursive_folding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRecursiveFoldingEnabled() bool { //gd:Tree.is_recursive_folding_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_is_recursive_folding_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDropModeFlags(flags gd.Int) { //gd:Tree.set_drop_mode_flags
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_drop_mode_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDropModeFlags() gd.Int { //gd:Tree.get_drop_mode_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_drop_mode_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowRmbSelect(allow bool) { //gd:Tree.set_allow_rmb_select
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_allow_rmb_select, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAllowRmbSelect() bool { //gd:Tree.get_allow_rmb_select
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_allow_rmb_select, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowReselect(allow bool) { //gd:Tree.set_allow_reselect
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAllowReselect() bool { //gd:Tree.get_allow_reselect
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_allow_reselect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowSearch(allow bool) { //gd:Tree.set_allow_search
	var frame = callframe.New()
	callframe.Arg(frame, allow)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_set_allow_search, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAllowSearch() bool { //gd:Tree.get_allow_search
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Tree.Bind_get_allow_search, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnItemSelected(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("item_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCellSelected(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("cell_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMultiSelected(cb func(item [1]gdclass.TreeItem, column int, selected bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("multi_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemMouseSelected(cb func(mouse_position Vector2.XY, mouse_button_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("item_mouse_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnEmptyClicked(cb func(click_position Vector2.XY, mouse_button_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("empty_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemEdited(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("item_edited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCustomItemClicked(cb func(mouse_button_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("custom_item_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemIconDoubleClicked(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("item_icon_double_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemCollapsed(cb func(item [1]gdclass.TreeItem)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("item_collapsed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCheckPropagatedToItem(cb func(item [1]gdclass.TreeItem, column int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("check_propagated_to_item"), gd.NewCallable(cb), 0)
}

func (self Instance) OnButtonClicked(cb func(item [1]gdclass.TreeItem, column int, id int, mouse_button_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("button_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCustomPopupEdited(cb func(arrow_clicked bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("custom_popup_edited"), gd.NewCallable(cb), 0)
}

func (self Instance) OnItemActivated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("item_activated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnColumnTitleClicked(cb func(column int, mouse_button_index int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("column_title_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNothingSelected(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("nothing_selected"), gd.NewCallable(cb), 0)
}

func (self class) AsTree() Advanced            { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTree() Instance         { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("Tree", func(ptr gd.Object) any { return [1]gdclass.Tree{*(*gdclass.Tree)(unsafe.Pointer(&ptr))} })
}

type SelectMode = gdclass.TreeSelectMode //gd:Tree.SelectMode

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

type DropModeFlags = gdclass.TreeDropModeFlags //gd:Tree.DropModeFlags

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

type HorizontalAlignment int

const (
	/*Horizontal left alignment, usually for text-derived classes.*/
	HorizontalAlignmentLeft HorizontalAlignment = 0
	/*Horizontal center alignment, usually for text-derived classes.*/
	HorizontalAlignmentCenter HorizontalAlignment = 1
	/*Horizontal right alignment, usually for text-derived classes.*/
	HorizontalAlignmentRight HorizontalAlignment = 2
	/*Expand row to fit width, usually for text-derived classes.*/
	HorizontalAlignmentFill HorizontalAlignment = 3
)
