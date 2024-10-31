package TreeItem

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A single item of a [Tree] control. It can contain other [TreeItem]s as children, which allows it to create a hierarchy. It can also contain text and buttons. [TreeItem] is not a [Node], it is internal to the [Tree].
To create a [TreeItem], use [method Tree.create_item] or [method TreeItem.create_child]. To remove a [TreeItem], use [method Object.free].
[b]Note:[/b] The ID values used for buttons are 32-bit, unlike [int] which is always 64-bit. They go from [code]-2147483648[/code] to [code]2147483647[/code].
*/
type Instance [1]classdb.TreeItem

/*
Sets the given column's cell mode to [param mode]. This determines how the cell is displayed and edited. See [enum TreeCellMode] constants for details.
*/
func (self Instance) SetCellMode(column int, mode classdb.TreeItemTreeCellMode) {
	class(self).SetCellMode(gd.Int(column), mode)
}

/*
Returns the column's cell mode.
*/
func (self Instance) GetCellMode(column int) classdb.TreeItemTreeCellMode {
	return classdb.TreeItemTreeCellMode(class(self).GetCellMode(gd.Int(column)))
}

/*
If [param multiline] is [code]true[/code], the given [param column] is multiline editable.
[b]Note:[/b] This option only affects the type of control ([LineEdit] or [TextEdit]) that appears when editing the column. You can set multiline values with [method set_text] even if the column is not multiline editable.
*/
func (self Instance) SetEditMultiline(column int, multiline bool) {
	class(self).SetEditMultiline(gd.Int(column), multiline)
}

/*
Returns [code]true[/code] if the given [param column] is multiline editable.
*/
func (self Instance) IsEditMultiline(column int) bool {
	return bool(class(self).IsEditMultiline(gd.Int(column)))
}

/*
If [param checked] is [code]true[/code], the given [param column] is checked. Clears column's indeterminate status.
*/
func (self Instance) SetChecked(column int, checked bool) {
	class(self).SetChecked(gd.Int(column), checked)
}

/*
If [param indeterminate] is [code]true[/code], the given [param column] is marked indeterminate.
[b]Note:[/b] If set [code]true[/code] from [code]false[/code], then column is cleared of checked status.
*/
func (self Instance) SetIndeterminate(column int, indeterminate bool) {
	class(self).SetIndeterminate(gd.Int(column), indeterminate)
}

/*
Returns [code]true[/code] if the given [param column] is checked.
*/
func (self Instance) IsChecked(column int) bool {
	return bool(class(self).IsChecked(gd.Int(column)))
}

/*
Returns [code]true[/code] if the given [param column] is indeterminate.
*/
func (self Instance) IsIndeterminate(column int) bool {
	return bool(class(self).IsIndeterminate(gd.Int(column)))
}

/*
Propagates this item's checked status to its children and parents for the given [param column]. It is possible to process the items affected by this method call by connecting to [signal Tree.check_propagated_to_item]. The order that the items affected will be processed is as follows: the item invoking this method, children of that item, and finally parents of that item. If [param emit_signal] is [code]false[/code], then [signal Tree.check_propagated_to_item] will not be emitted.
*/
func (self Instance) PropagateCheck(column int) {
	class(self).PropagateCheck(gd.Int(column), true)
}

/*
Sets the given column's text value.
*/
func (self Instance) SetText(column int, text string) {
	class(self).SetText(gd.Int(column), gd.NewString(text))
}

/*
Returns the given column's text.
*/
func (self Instance) GetText(column int) string {
	return string(class(self).GetText(gd.Int(column)).String())
}

/*
Sets item's text base writing direction.
*/
func (self Instance) SetTextDirection(column int, direction classdb.ControlTextDirection) {
	class(self).SetTextDirection(gd.Int(column), direction)
}

/*
Returns item's text base writing direction.
*/
func (self Instance) GetTextDirection(column int) classdb.ControlTextDirection {
	return classdb.ControlTextDirection(class(self).GetTextDirection(gd.Int(column)))
}

/*
Sets the autowrap mode in the given [param column]. If set to something other than [constant TextServer.AUTOWRAP_OFF], the text gets wrapped inside the cell's bounding rectangle.
*/
func (self Instance) SetAutowrapMode(column int, autowrap_mode classdb.TextServerAutowrapMode) {
	class(self).SetAutowrapMode(gd.Int(column), autowrap_mode)
}

/*
Returns the text autowrap mode in the given [param column]. By default it is [constant TextServer.AUTOWRAP_OFF].
*/
func (self Instance) GetAutowrapMode(column int) classdb.TextServerAutowrapMode {
	return classdb.TextServerAutowrapMode(class(self).GetAutowrapMode(gd.Int(column)))
}

/*
Sets the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column].
*/
func (self Instance) SetTextOverrunBehavior(column int, overrun_behavior classdb.TextServerOverrunBehavior) {
	class(self).SetTextOverrunBehavior(gd.Int(column), overrun_behavior)
}

/*
Returns the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column]. By default it is [constant TextServer.OVERRUN_TRIM_ELLIPSIS].
*/
func (self Instance) GetTextOverrunBehavior(column int) classdb.TextServerOverrunBehavior {
	return classdb.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior(gd.Int(column)))
}

/*
Set BiDi algorithm override for the structured text. Has effect for cells that display text.
*/
func (self Instance) SetStructuredTextBidiOverride(column int, parser classdb.TextServerStructuredTextParser) {
	class(self).SetStructuredTextBidiOverride(gd.Int(column), parser)
}

/*
Returns the BiDi algorithm override set for this cell.
*/
func (self Instance) GetStructuredTextBidiOverride(column int) classdb.TextServerStructuredTextParser {
	return classdb.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride(gd.Int(column)))
}

/*
Set additional options for BiDi override. Has effect for cells that display text.
*/
func (self Instance) SetStructuredTextBidiOverrideOptions(column int, args gd.Array) {
	class(self).SetStructuredTextBidiOverrideOptions(gd.Int(column), args)
}

/*
Returns the additional BiDi options set for this cell.
*/
func (self Instance) GetStructuredTextBidiOverrideOptions(column int) gd.Array {
	return gd.Array(class(self).GetStructuredTextBidiOverrideOptions(gd.Int(column)))
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Instance) SetLanguage(column int, language string) {
	class(self).SetLanguage(gd.Int(column), gd.NewString(language))
}

/*
Returns item's text language code.
*/
func (self Instance) GetLanguage(column int) string {
	return string(class(self).GetLanguage(gd.Int(column)).String())
}

/*
Sets a string to be shown after a column's value (for example, a unit abbreviation).
*/
func (self Instance) SetSuffix(column int, text string) {
	class(self).SetSuffix(gd.Int(column), gd.NewString(text))
}

/*
Gets the suffix string shown after the column value.
*/
func (self Instance) GetSuffix(column int) string {
	return string(class(self).GetSuffix(gd.Int(column)).String())
}

/*
Sets the given cell's icon [Texture2D]. The cell has to be in [constant CELL_MODE_ICON] mode.
*/
func (self Instance) SetIcon(column int, texture gdclass.Texture2D) {
	class(self).SetIcon(gd.Int(column), texture)
}

/*
Returns the given column's icon [Texture2D]. Error if no icon is set.
*/
func (self Instance) GetIcon(column int) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetIcon(gd.Int(column)))
}

/*
Sets the given column's icon's texture region.
*/
func (self Instance) SetIconRegion(column int, region gd.Rect2) {
	class(self).SetIconRegion(gd.Int(column), region)
}

/*
Returns the icon [Texture2D] region as [Rect2].
*/
func (self Instance) GetIconRegion(column int) gd.Rect2 {
	return gd.Rect2(class(self).GetIconRegion(gd.Int(column)))
}

/*
Sets the maximum allowed width of the icon in the given [param column]. This limit is applied on top of the default size of the icon and on top of [theme_item Tree.icon_max_width]. The height is adjusted according to the icon's ratio.
*/
func (self Instance) SetIconMaxWidth(column int, width int) {
	class(self).SetIconMaxWidth(gd.Int(column), gd.Int(width))
}

/*
Returns the maximum allowed width of the icon in the given [param column].
*/
func (self Instance) GetIconMaxWidth(column int) int {
	return int(int(class(self).GetIconMaxWidth(gd.Int(column))))
}

/*
Modulates the given column's icon with [param modulate].
*/
func (self Instance) SetIconModulate(column int, modulate gd.Color) {
	class(self).SetIconModulate(gd.Int(column), modulate)
}

/*
Returns the [Color] modulating the column's icon.
*/
func (self Instance) GetIconModulate(column int) gd.Color {
	return gd.Color(class(self).GetIconModulate(gd.Int(column)))
}

/*
Sets the value of a [constant CELL_MODE_RANGE] column.
*/
func (self Instance) SetRange(column int, value float64) {
	class(self).SetRange(gd.Int(column), gd.Float(value))
}

/*
Returns the value of a [constant CELL_MODE_RANGE] column.
*/
func (self Instance) GetRange(column int) float64 {
	return float64(float64(class(self).GetRange(gd.Int(column))))
}

/*
Sets the range of accepted values for a column. The column must be in the [constant CELL_MODE_RANGE] mode.
If [param expr] is [code]true[/code], the edit mode slider will use an exponential scale as with [member Range.exp_edit].
*/
func (self Instance) SetRangeConfig(column int, min float64, max float64, step float64) {
	class(self).SetRangeConfig(gd.Int(column), gd.Float(min), gd.Float(max), gd.Float(step), false)
}

/*
Returns a dictionary containing the range parameters for a given column. The keys are "min", "max", "step", and "expr".
*/
func (self Instance) GetRangeConfig(column int) gd.Dictionary {
	return gd.Dictionary(class(self).GetRangeConfig(gd.Int(column)))
}

/*
Sets the metadata value for the given column, which can be retrieved later using [method get_metadata]. This can be used, for example, to store a reference to the original data.
*/
func (self Instance) SetMetadata(column int, meta gd.Variant) {
	class(self).SetMetadata(gd.Int(column), meta)
}

/*
Returns the metadata value that was set for the given column using [method set_metadata].
*/
func (self Instance) GetMetadata(column int) gd.Variant {
	return gd.Variant(class(self).GetMetadata(gd.Int(column)))
}

/*
Sets the given column's custom draw callback to the [param callback] method on [param object].
The method named [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
func (self Instance) SetCustomDraw(column int, obj gd.Object, callback string) {
	class(self).SetCustomDraw(gd.Int(column), obj, gd.NewStringName(callback))
}

/*
Sets the given column's custom draw callback. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback. The cell has to be in [constant CELL_MODE_CUSTOM] to use this feature.
The [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
func (self Instance) SetCustomDrawCallback(column int, callback gd.Callable) {
	class(self).SetCustomDrawCallback(gd.Int(column), callback)
}

/*
Returns the custom callback of column [param column].
*/
func (self Instance) GetCustomDrawCallback(column int) gd.Callable {
	return gd.Callable(class(self).GetCustomDrawCallback(gd.Int(column)))
}

/*
Collapses or uncollapses this [TreeItem] and all the descendants of this item.
*/
func (self Instance) SetCollapsedRecursive(enable bool) {
	class(self).SetCollapsedRecursive(enable)
}

/*
Returns [code]true[/code] if this [TreeItem], or any of its descendants, is collapsed.
If [param only_visible] is [code]true[/code] it ignores non-visible [TreeItem]s.
*/
func (self Instance) IsAnyCollapsed() bool {
	return bool(class(self).IsAnyCollapsed(false))
}

/*
Returns [code]true[/code] if [member visible] is [code]true[/code] and all its ancestors are also visible.
*/
func (self Instance) IsVisibleInTree() bool {
	return bool(class(self).IsVisibleInTree())
}

/*
Uncollapses all [TreeItem]s necessary to reveal this [TreeItem], i.e. all ancestor [TreeItem]s.
*/
func (self Instance) UncollapseTree() {
	class(self).UncollapseTree()
}

/*
If [param selectable] is [code]true[/code], the given [param column] is selectable.
*/
func (self Instance) SetSelectable(column int, selectable bool) {
	class(self).SetSelectable(gd.Int(column), selectable)
}

/*
Returns [code]true[/code] if the given [param column] is selectable.
*/
func (self Instance) IsSelectable(column int) bool {
	return bool(class(self).IsSelectable(gd.Int(column)))
}

/*
Returns [code]true[/code] if the given [param column] is selected.
*/
func (self Instance) IsSelected(column int) bool {
	return bool(class(self).IsSelected(gd.Int(column)))
}

/*
Selects the given [param column].
*/
func (self Instance) Select(column int) {
	class(self).Select(gd.Int(column))
}

/*
Deselects the given column.
*/
func (self Instance) Deselect(column int) {
	class(self).Deselect(gd.Int(column))
}

/*
If [param enabled] is [code]true[/code], the given [param column] is editable.
*/
func (self Instance) SetEditable(column int, enabled bool) {
	class(self).SetEditable(gd.Int(column), enabled)
}

/*
Returns [code]true[/code] if the given [param column] is editable.
*/
func (self Instance) IsEditable(column int) bool {
	return bool(class(self).IsEditable(gd.Int(column)))
}

/*
Sets the given column's custom color.
*/
func (self Instance) SetCustomColor(column int, color gd.Color) {
	class(self).SetCustomColor(gd.Int(column), color)
}

/*
Returns the custom color of column [param column].
*/
func (self Instance) GetCustomColor(column int) gd.Color {
	return gd.Color(class(self).GetCustomColor(gd.Int(column)))
}

/*
Resets the color for the given column to default.
*/
func (self Instance) ClearCustomColor(column int) {
	class(self).ClearCustomColor(gd.Int(column))
}

/*
Sets custom font used to draw text in the given [param column].
*/
func (self Instance) SetCustomFont(column int, font gdclass.Font) {
	class(self).SetCustomFont(gd.Int(column), font)
}

/*
Returns custom font used to draw text in the column [param column].
*/
func (self Instance) GetCustomFont(column int) gdclass.Font {
	return gdclass.Font(class(self).GetCustomFont(gd.Int(column)))
}

/*
Sets custom font size used to draw text in the given [param column].
*/
func (self Instance) SetCustomFontSize(column int, font_size int) {
	class(self).SetCustomFontSize(gd.Int(column), gd.Int(font_size))
}

/*
Returns custom font size used to draw text in the column [param column].
*/
func (self Instance) GetCustomFontSize(column int) int {
	return int(int(class(self).GetCustomFontSize(gd.Int(column))))
}

/*
Sets the given column's custom background color and whether to just use it as an outline.
*/
func (self Instance) SetCustomBgColor(column int, color gd.Color) {
	class(self).SetCustomBgColor(gd.Int(column), color, false)
}

/*
Resets the background color for the given column to default.
*/
func (self Instance) ClearCustomBgColor(column int) {
	class(self).ClearCustomBgColor(gd.Int(column))
}

/*
Returns the custom background color of column [param column].
*/
func (self Instance) GetCustomBgColor(column int) gd.Color {
	return gd.Color(class(self).GetCustomBgColor(gd.Int(column)))
}

/*
Makes a cell with [constant CELL_MODE_CUSTOM] display as a non-flat button with a [StyleBox].
*/
func (self Instance) SetCustomAsButton(column int, enable bool) {
	class(self).SetCustomAsButton(gd.Int(column), enable)
}

/*
Returns [code]true[/code] if the cell was made into a button with [method set_custom_as_button].
*/
func (self Instance) IsCustomSetAsButton(column int) bool {
	return bool(class(self).IsCustomSetAsButton(gd.Int(column)))
}

/*
Adds a button with [Texture2D] [param button] at column [param column]. The [param id] is used to identify the button in the according [signal Tree.button_clicked] signal and can be different from the buttons index. If not specified, the next available index is used, which may be retrieved by calling [method get_button_count] immediately before this method. Optionally, the button can be [param disabled] and have a [param tooltip_text].
*/
func (self Instance) AddButton(column int, button gdclass.Texture2D) {
	class(self).AddButton(gd.Int(column), button, gd.Int(-1), false, gd.NewString(""))
}

/*
Returns the number of buttons in column [param column].
*/
func (self Instance) GetButtonCount(column int) int {
	return int(int(class(self).GetButtonCount(gd.Int(column))))
}

/*
Returns the tooltip text for the button at index [param button_index] in column [param column].
*/
func (self Instance) GetButtonTooltipText(column int, button_index int) string {
	return string(class(self).GetButtonTooltipText(gd.Int(column), gd.Int(button_index)).String())
}

/*
Returns the ID for the button at index [param button_index] in column [param column].
*/
func (self Instance) GetButtonId(column int, button_index int) int {
	return int(int(class(self).GetButtonId(gd.Int(column), gd.Int(button_index))))
}

/*
Returns the button index if there is a button with ID [param id] in column [param column], otherwise returns -1.
*/
func (self Instance) GetButtonById(column int, id int) int {
	return int(int(class(self).GetButtonById(gd.Int(column), gd.Int(id))))
}

/*
Returns the color of the button with ID [param id] in column [param column]. If the specified button does not exist, returns [constant Color.BLACK].
*/
func (self Instance) GetButtonColor(column int, id int) gd.Color {
	return gd.Color(class(self).GetButtonColor(gd.Int(column), gd.Int(id)))
}

/*
Returns the [Texture2D] of the button at index [param button_index] in column [param column].
*/
func (self Instance) GetButton(column int, button_index int) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetButton(gd.Int(column), gd.Int(button_index)))
}

/*
Sets the tooltip text for the button at index [param button_index] in the given [param column].
*/
func (self Instance) SetButtonTooltipText(column int, button_index int, tooltip string) {
	class(self).SetButtonTooltipText(gd.Int(column), gd.Int(button_index), gd.NewString(tooltip))
}

/*
Sets the given column's button [Texture2D] at index [param button_index] to [param button].
*/
func (self Instance) SetButton(column int, button_index int, button gdclass.Texture2D) {
	class(self).SetButton(gd.Int(column), gd.Int(button_index), button)
}

/*
Removes the button at index [param button_index] in column [param column].
*/
func (self Instance) EraseButton(column int, button_index int) {
	class(self).EraseButton(gd.Int(column), gd.Int(button_index))
}

/*
If [code]true[/code], disables the button at index [param button_index] in the given [param column].
*/
func (self Instance) SetButtonDisabled(column int, button_index int, disabled bool) {
	class(self).SetButtonDisabled(gd.Int(column), gd.Int(button_index), disabled)
}

/*
Sets the given column's button color at index [param button_index] to [param color].
*/
func (self Instance) SetButtonColor(column int, button_index int, color gd.Color) {
	class(self).SetButtonColor(gd.Int(column), gd.Int(button_index), color)
}

/*
Returns [code]true[/code] if the button at index [param button_index] for the given [param column] is disabled.
*/
func (self Instance) IsButtonDisabled(column int, button_index int) bool {
	return bool(class(self).IsButtonDisabled(gd.Int(column), gd.Int(button_index)))
}

/*
Sets the given column's tooltip text.
*/
func (self Instance) SetTooltipText(column int, tooltip string) {
	class(self).SetTooltipText(gd.Int(column), gd.NewString(tooltip))
}

/*
Returns the given column's tooltip text.
*/
func (self Instance) GetTooltipText(column int) string {
	return string(class(self).GetTooltipText(gd.Int(column)).String())
}

/*
Sets the given column's text alignment. See [enum HorizontalAlignment] for possible values.
*/
func (self Instance) SetTextAlignment(column int, text_alignment gdconst.HorizontalAlignment) {
	class(self).SetTextAlignment(gd.Int(column), text_alignment)
}

/*
Returns the given column's text alignment.
*/
func (self Instance) GetTextAlignment(column int) gdconst.HorizontalAlignment {
	return gdconst.HorizontalAlignment(class(self).GetTextAlignment(gd.Int(column)))
}

/*
If [param enable] is [code]true[/code], the given [param column] is expanded to the right.
*/
func (self Instance) SetExpandRight(column int, enable bool) {
	class(self).SetExpandRight(gd.Int(column), enable)
}

/*
Returns [code]true[/code] if [code]expand_right[/code] is set.
*/
func (self Instance) GetExpandRight(column int) bool {
	return bool(class(self).GetExpandRight(gd.Int(column)))
}

/*
Creates an item and adds it as a child.
The new item will be inserted as position [param index] (the default value [code]-1[/code] means the last position), or it will be the last child if [param index] is higher than the child count.
*/
func (self Instance) CreateChild() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).CreateChild(gd.Int(-1)))
}

/*
Adds a previously unparented [TreeItem] as a direct child of this one. The [param child] item must not be a part of any [Tree] or parented to any [TreeItem]. See also [method remove_child].
*/
func (self Instance) AddChild(child gdclass.TreeItem) {
	class(self).AddChild(child)
}

/*
Removes the given child [TreeItem] and all its children from the [Tree]. Note that it doesn't free the item from memory, so it can be reused later (see [method add_child]). To completely remove a [TreeItem] use [method Object.free].
[b]Note:[/b] If you want to move a child from one [Tree] to another, then instead of removing and adding it manually you can use [method move_before] or [method move_after].
*/
func (self Instance) RemoveChild(child gdclass.TreeItem) {
	class(self).RemoveChild(child)
}

/*
Returns the [Tree] that owns this TreeItem.
*/
func (self Instance) GetTree() gdclass.Tree {
	return gdclass.Tree(class(self).GetTree())
}

/*
Returns the next sibling TreeItem in the tree or a null object if there is none.
*/
func (self Instance) GetNext() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetNext())
}

/*
Returns the previous sibling TreeItem in the tree or a null object if there is none.
*/
func (self Instance) GetPrev() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetPrev())
}

/*
Returns the parent TreeItem or a null object if there is none.
*/
func (self Instance) GetParent() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetParent())
}

/*
Returns the TreeItem's first child.
*/
func (self Instance) GetFirstChild() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetFirstChild())
}

/*
Returns the next TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first element in the tree when called on the last element, otherwise it returns [code]null[/code].
*/
func (self Instance) GetNextInTree() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetNextInTree(false))
}

/*
Returns the previous TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
func (self Instance) GetPrevInTree() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetPrevInTree(false))
}

/*
Returns the next visible TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first visible element in the tree when called on the last visible element, otherwise it returns [code]null[/code].
*/
func (self Instance) GetNextVisible() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetNextVisible(false))
}

/*
Returns the previous visible sibling TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last visible element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
func (self Instance) GetPrevVisible() gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetPrevVisible(false))
}

/*
Returns a child item by its [param index] (see [method get_child_count]). This method is often used for iterating all children of an item.
Negative indices access the children from the last one.
*/
func (self Instance) GetChild(index int) gdclass.TreeItem {
	return gdclass.TreeItem(class(self).GetChild(gd.Int(index)))
}

/*
Returns the number of child items.
*/
func (self Instance) GetChildCount() int {
	return int(int(class(self).GetChildCount()))
}

/*
Returns an array of references to the item's children.
*/
func (self Instance) GetChildren() gd.Array {
	return gd.Array(class(self).GetChildren())
}

/*
Returns the node's order in the tree. For example, if called on the first child item the position is [code]0[/code].
*/
func (self Instance) GetIndex() int {
	return int(int(class(self).GetIndex()))
}

/*
Moves this TreeItem right before the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
func (self Instance) MoveBefore(item gdclass.TreeItem) {
	class(self).MoveBefore(item)
}

/*
Moves this TreeItem right after the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
func (self Instance) MoveAfter(item gdclass.TreeItem) {
	class(self).MoveAfter(item)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TreeItem

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TreeItem"))
	return Instance{classdb.TreeItem(object)}
}

func (self Instance) Collapsed() bool {
	return bool(class(self).IsCollapsed())
}

func (self Instance) SetCollapsed(value bool) {
	class(self).SetCollapsed(value)
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisible(value)
}

func (self Instance) DisableFolding() bool {
	return bool(class(self).IsFoldingDisabled())
}

func (self Instance) SetDisableFolding(value bool) {
	class(self).SetDisableFolding(value)
}

func (self Instance) CustomMinimumHeight() int {
	return int(int(class(self).GetCustomMinimumHeight()))
}

func (self Instance) SetCustomMinimumHeight(value int) {
	class(self).SetCustomMinimumHeight(gd.Int(value))
}

/*
Sets the given column's cell mode to [param mode]. This determines how the cell is displayed and edited. See [enum TreeCellMode] constants for details.
*/
//go:nosplit
func (self class) SetCellMode(column gd.Int, mode classdb.TreeItemTreeCellMode) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_cell_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the column's cell mode.
*/
//go:nosplit
func (self class) GetCellMode(column gd.Int) classdb.TreeItemTreeCellMode {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TreeItemTreeCellMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_cell_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param multiline] is [code]true[/code], the given [param column] is multiline editable.
[b]Note:[/b] This option only affects the type of control ([LineEdit] or [TextEdit]) that appears when editing the column. You can set multiline values with [method set_text] even if the column is not multiline editable.
*/
//go:nosplit
func (self class) SetEditMultiline(column gd.Int, multiline bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, multiline)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_edit_multiline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param column] is multiline editable.
*/
//go:nosplit
func (self class) IsEditMultiline(column gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_edit_multiline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param checked] is [code]true[/code], the given [param column] is checked. Clears column's indeterminate status.
*/
//go:nosplit
func (self class) SetChecked(column gd.Int, checked bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If [param indeterminate] is [code]true[/code], the given [param column] is marked indeterminate.
[b]Note:[/b] If set [code]true[/code] from [code]false[/code], then column is cleared of checked status.
*/
//go:nosplit
func (self class) SetIndeterminate(column gd.Int, indeterminate bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, indeterminate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param column] is checked.
*/
//go:nosplit
func (self class) IsChecked(column gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [param column] is indeterminate.
*/
//go:nosplit
func (self class) IsIndeterminate(column gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Propagates this item's checked status to its children and parents for the given [param column]. It is possible to process the items affected by this method call by connecting to [signal Tree.check_propagated_to_item]. The order that the items affected will be processed is as follows: the item invoking this method, children of that item, and finally parents of that item. If [param emit_signal] is [code]false[/code], then [signal Tree.check_propagated_to_item] will not be emitted.
*/
//go:nosplit
func (self class) PropagateCheck(column gd.Int, emit_signal bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, emit_signal)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_propagate_check, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the given column's text value.
*/
//go:nosplit
func (self class) SetText(column gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the given column's text.
*/
//go:nosplit
func (self class) GetText(column gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetTextDirection(column gd.Int, direction classdb.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetTextDirection(column gd.Int) classdb.ControlTextDirection {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the autowrap mode in the given [param column]. If set to something other than [constant TextServer.AUTOWRAP_OFF], the text gets wrapped inside the cell's bounding rectangle.
*/
//go:nosplit
func (self class) SetAutowrapMode(column gd.Int, autowrap_mode classdb.TextServerAutowrapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the text autowrap mode in the given [param column]. By default it is [constant TextServer.AUTOWRAP_OFF].
*/
//go:nosplit
func (self class) GetAutowrapMode(column gd.Int) classdb.TextServerAutowrapMode {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column].
*/
//go:nosplit
func (self class) SetTextOverrunBehavior(column gd.Int, overrun_behavior classdb.TextServerOverrunBehavior) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column]. By default it is [constant TextServer.OVERRUN_TRIM_ELLIPSIS].
*/
//go:nosplit
func (self class) GetTextOverrunBehavior(column gd.Int) classdb.TextServerOverrunBehavior {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set BiDi algorithm override for the structured text. Has effect for cells that display text.
*/
//go:nosplit
func (self class) SetStructuredTextBidiOverride(column gd.Int, parser classdb.TextServerStructuredTextParser) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the BiDi algorithm override set for this cell.
*/
//go:nosplit
func (self class) GetStructuredTextBidiOverride(column gd.Int) classdb.TextServerStructuredTextParser {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set additional options for BiDi override. Has effect for cells that display text.
*/
//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(column gd.Int, args gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(args))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the additional BiDi options set for this cell.
*/
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(column gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetLanguage(column gd.Int, language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetLanguage(column gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets a string to be shown after a column's value (for example, a unit abbreviation).
*/
//go:nosplit
func (self class) SetSuffix(column gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Gets the suffix string shown after the column value.
*/
//go:nosplit
func (self class) GetSuffix(column gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the given cell's icon [Texture2D]. The cell has to be in [constant CELL_MODE_ICON] mode.
*/
//go:nosplit
func (self class) SetIcon(column gd.Int, texture gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the given column's icon [Texture2D]. Error if no icon is set.
*/
//go:nosplit
func (self class) GetIcon(column gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets the given column's icon's texture region.
*/
//go:nosplit
func (self class) SetIconRegion(column gd.Int, region gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, region)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the icon [Texture2D] region as [Rect2].
*/
//go:nosplit
func (self class) GetIconRegion(column gd.Int) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum allowed width of the icon in the given [param column]. This limit is applied on top of the default size of the icon and on top of [theme_item Tree.icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetIconMaxWidth(column gd.Int, width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the maximum allowed width of the icon in the given [param column].
*/
//go:nosplit
func (self class) GetIconMaxWidth(column gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Modulates the given column's icon with [param modulate].
*/
//go:nosplit
func (self class) SetIconModulate(column gd.Int, modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Color] modulating the column's icon.
*/
//go:nosplit
func (self class) GetIconModulate(column gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the value of a [constant CELL_MODE_RANGE] column.
*/
//go:nosplit
func (self class) SetRange(column gd.Int, value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the value of a [constant CELL_MODE_RANGE] column.
*/
//go:nosplit
func (self class) GetRange(column gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the range of accepted values for a column. The column must be in the [constant CELL_MODE_RANGE] mode.
If [param expr] is [code]true[/code], the edit mode slider will use an exponential scale as with [member Range.exp_edit].
*/
//go:nosplit
func (self class) SetRangeConfig(column gd.Int, min gd.Float, max gd.Float, step gd.Float, expr bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, min)
	callframe.Arg(frame, max)
	callframe.Arg(frame, step)
	callframe.Arg(frame, expr)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_range_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a dictionary containing the range parameters for a given column. The keys are "min", "max", "step", and "expr".
*/
//go:nosplit
func (self class) GetRangeConfig(column gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_range_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the metadata value for the given column, which can be retrieved later using [method get_metadata]. This can be used, for example, to store a reference to the original data.
*/
//go:nosplit
func (self class) SetMetadata(column gd.Int, meta gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the metadata value that was set for the given column using [method set_metadata].
*/
//go:nosplit
func (self class) GetMetadata(column gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the given column's custom draw callback to the [param callback] method on [param object].
The method named [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
//go:nosplit
func (self class) SetCustomDraw(column gd.Int, obj gd.Object, callback gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the given column's custom draw callback. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback. The cell has to be in [constant CELL_MODE_CUSTOM] to use this feature.
The [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
//go:nosplit
func (self class) SetCustomDrawCallback(column gd.Int, callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_draw_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the custom callback of column [param column].
*/
//go:nosplit
func (self class) GetCustomDrawCallback(column gd.Int) gd.Callable {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_draw_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Callable](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollapsed(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollapsed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Collapses or uncollapses this [TreeItem] and all the descendants of this item.
*/
//go:nosplit
func (self class) SetCollapsedRecursive(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_collapsed_recursive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if this [TreeItem], or any of its descendants, is collapsed.
If [param only_visible] is [code]true[/code] it ignores non-visible [TreeItem]s.
*/
//go:nosplit
func (self class) IsAnyCollapsed(only_visible bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, only_visible)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_any_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisible(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [member visible] is [code]true[/code] and all its ancestors are also visible.
*/
//go:nosplit
func (self class) IsVisibleInTree() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_visible_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Uncollapses all [TreeItem]s necessary to reveal this [TreeItem], i.e. all ancestor [TreeItem]s.
*/
//go:nosplit
func (self class) UncollapseTree() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_uncollapse_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCustomMinimumHeight(height gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_minimum_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomMinimumHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_minimum_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param selectable] is [code]true[/code], the given [param column] is selectable.
*/
//go:nosplit
func (self class) SetSelectable(column gd.Int, selectable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, selectable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param column] is selectable.
*/
//go:nosplit
func (self class) IsSelectable(column gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [param column] is selected.
*/
//go:nosplit
func (self class) IsSelected(column gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Selects the given [param column].
*/
//go:nosplit
func (self class) Select(column gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Deselects the given column.
*/
//go:nosplit
func (self class) Deselect(column gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If [param enabled] is [code]true[/code], the given [param column] is editable.
*/
//go:nosplit
func (self class) SetEditable(column gd.Int, enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param column] is editable.
*/
//go:nosplit
func (self class) IsEditable(column gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the given column's custom color.
*/
//go:nosplit
func (self class) SetCustomColor(column gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the custom color of column [param column].
*/
//go:nosplit
func (self class) GetCustomColor(column gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resets the color for the given column to default.
*/
//go:nosplit
func (self class) ClearCustomColor(column gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_clear_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets custom font used to draw text in the given [param column].
*/
//go:nosplit
func (self class) SetCustomFont(column gd.Int, font gdclass.Font) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns custom font used to draw text in the column [param column].
*/
//go:nosplit
func (self class) GetCustomFont(column gd.Int) gdclass.Font {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Font{classdb.Font(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets custom font size used to draw text in the given [param column].
*/
//go:nosplit
func (self class) SetCustomFontSize(column gd.Int, font_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns custom font size used to draw text in the column [param column].
*/
//go:nosplit
func (self class) GetCustomFontSize(column gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the given column's custom background color and whether to just use it as an outline.
*/
//go:nosplit
func (self class) SetCustomBgColor(column gd.Int, color gd.Color, just_outline bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, color)
	callframe.Arg(frame, just_outline)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Resets the background color for the given column to default.
*/
//go:nosplit
func (self class) ClearCustomBgColor(column gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_clear_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the custom background color of column [param column].
*/
//go:nosplit
func (self class) GetCustomBgColor(column gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Makes a cell with [constant CELL_MODE_CUSTOM] display as a non-flat button with a [StyleBox].
*/
//go:nosplit
func (self class) SetCustomAsButton(column gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_as_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the cell was made into a button with [method set_custom_as_button].
*/
//go:nosplit
func (self class) IsCustomSetAsButton(column gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_custom_set_as_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a button with [Texture2D] [param button] at column [param column]. The [param id] is used to identify the button in the according [signal Tree.button_clicked] signal and can be different from the buttons index. If not specified, the next available index is used, which may be retrieved by calling [method get_button_count] immediately before this method. Optionally, the button can be [param disabled] and have a [param tooltip_text].
*/
//go:nosplit
func (self class) AddButton(column gd.Int, button gdclass.Texture2D, id gd.Int, disabled bool, tooltip_text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(button[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, disabled)
	callframe.Arg(frame, pointers.Get(tooltip_text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_add_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of buttons in column [param column].
*/
//go:nosplit
func (self class) GetButtonCount(column gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tooltip text for the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButtonTooltipText(column gd.Int, button_index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the ID for the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButtonId(column gd.Int, button_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the button index if there is a button with ID [param id] in column [param column], otherwise returns -1.
*/
//go:nosplit
func (self class) GetButtonById(column gd.Int, id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_by_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the color of the button with ID [param id] in column [param column]. If the specified button does not exist, returns [constant Color.BLACK].
*/
//go:nosplit
func (self class) GetButtonColor(column gd.Int, id gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Texture2D] of the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButton(column gd.Int, button_index gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets the tooltip text for the button at index [param button_index] in the given [param column].
*/
//go:nosplit
func (self class) SetButtonTooltipText(column gd.Int, button_index gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_button_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the given column's button [Texture2D] at index [param button_index] to [param button].
*/
//go:nosplit
func (self class) SetButton(column gd.Int, button_index gd.Int, button gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, pointers.Get(button[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) EraseButton(column gd.Int, button_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_erase_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If [code]true[/code], disables the button at index [param button_index] in the given [param column].
*/
//go:nosplit
func (self class) SetButtonDisabled(column gd.Int, button_index gd.Int, disabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_button_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the given column's button color at index [param button_index] to [param color].
*/
//go:nosplit
func (self class) SetButtonColor(column gd.Int, button_index gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_button_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the button at index [param button_index] for the given [param column] is disabled.
*/
//go:nosplit
func (self class) IsButtonDisabled(column gd.Int, button_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_button_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the given column's tooltip text.
*/
//go:nosplit
func (self class) SetTooltipText(column gd.Int, tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the given column's tooltip text.
*/
//go:nosplit
func (self class) GetTooltipText(column gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the given column's text alignment. See [enum HorizontalAlignment] for possible values.
*/
//go:nosplit
func (self class) SetTextAlignment(column gd.Int, text_alignment gdconst.HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, text_alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_text_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the given column's text alignment.
*/
//go:nosplit
func (self class) GetTextAlignment(column gd.Int) gdconst.HorizontalAlignment {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gdconst.HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_text_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enable] is [code]true[/code], the given [param column] is expanded to the right.
*/
//go:nosplit
func (self class) SetExpandRight(column gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_expand_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if [code]expand_right[/code] is set.
*/
//go:nosplit
func (self class) GetExpandRight(column gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_expand_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableFolding(disable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_disable_folding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFoldingDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_folding_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates an item and adds it as a child.
The new item will be inserted as position [param index] (the default value [code]-1[/code] means the last position), or it will be the last child if [param index] is higher than the child count.
*/
//go:nosplit
func (self class) CreateChild(index gd.Int) gdclass.TreeItem {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_create_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Adds a previously unparented [TreeItem] as a direct child of this one. The [param child] item must not be a part of any [Tree] or parented to any [TreeItem]. See also [method remove_child].
*/
//go:nosplit
func (self class) AddChild(child gdclass.TreeItem) {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(child[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_add_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the given child [TreeItem] and all its children from the [Tree]. Note that it doesn't free the item from memory, so it can be reused later (see [method add_child]). To completely remove a [TreeItem] use [method Object.free].
[b]Note:[/b] If you want to move a child from one [Tree] to another, then instead of removing and adding it manually you can use [method move_before] or [method move_after].
*/
//go:nosplit
func (self class) RemoveChild(child gdclass.TreeItem) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(child[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_remove_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Tree] that owns this TreeItem.
*/
//go:nosplit
func (self class) GetTree() gdclass.Tree {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Tree{classdb.Tree(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the next sibling TreeItem in the tree or a null object if there is none.
*/
//go:nosplit
func (self class) GetNext() gdclass.TreeItem {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the previous sibling TreeItem in the tree or a null object if there is none.
*/
//go:nosplit
func (self class) GetPrev() gdclass.TreeItem {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_prev, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the parent TreeItem or a null object if there is none.
*/
//go:nosplit
func (self class) GetParent() gdclass.TreeItem {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the TreeItem's first child.
*/
//go:nosplit
func (self class) GetFirstChild() gdclass.TreeItem {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_first_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the next TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first element in the tree when called on the last element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetNextInTree(wrap bool) gdclass.TreeItem {
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_next_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the previous TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetPrevInTree(wrap bool) gdclass.TreeItem {
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_prev_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the next visible TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first visible element in the tree when called on the last visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetNextVisible(wrap bool) gdclass.TreeItem {
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_next_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the previous visible sibling TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last visible element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetPrevVisible(wrap bool) gdclass.TreeItem {
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_prev_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns a child item by its [param index] (see [method get_child_count]). This method is often used for iterating all children of an item.
Negative indices access the children from the last one.
*/
//go:nosplit
func (self class) GetChild(index gd.Int) gdclass.TreeItem {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TreeItem{classdb.TreeItem(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the number of child items.
*/
//go:nosplit
func (self class) GetChildCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_child_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of references to the item's children.
*/
//go:nosplit
func (self class) GetChildren() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_children, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the node's order in the tree. For example, if called on the first child item the position is [code]0[/code].
*/
//go:nosplit
func (self class) GetIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves this TreeItem right before the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
//go:nosplit
func (self class) MoveBefore(item gdclass.TreeItem) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_move_before, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Moves this TreeItem right after the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
//go:nosplit
func (self class) MoveAfter(item gdclass.TreeItem) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_move_after, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsTreeItem() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTreeItem() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() { classdb.Register("TreeItem", func(ptr gd.Object) any { return classdb.TreeItem(ptr) }) }

type TreeCellMode = classdb.TreeItemTreeCellMode

const (
	/*Cell shows a string label. When editable, the text can be edited using a [LineEdit], or a [TextEdit] popup if [method set_edit_multiline] is used.*/
	CellModeString TreeCellMode = 0
	/*Cell shows a checkbox, optionally with text. The checkbox can be pressed, released, or indeterminate (via [method set_indeterminate]). The checkbox can't be clicked unless the cell is editable.*/
	CellModeCheck TreeCellMode = 1
	/*Cell shows a numeric range. When editable, it can be edited using a range slider. Use [method set_range] to set the value and [method set_range_config] to configure the range.
	  This cell can also be used in a text dropdown mode when you assign a text with [method set_text]. Separate options with a comma, e.g. [code]"Option1,Option2,Option3"[/code].*/
	CellModeRange TreeCellMode = 2
	/*Cell shows an icon. It can't be edited nor display text.*/
	CellModeIcon TreeCellMode = 3
	/*Cell shows as a clickable button. It will display an arrow similar to [OptionButton], but doesn't feature a dropdown (for that you can use [constant CELL_MODE_RANGE]). Clicking the button emits the [signal Tree.item_edited] signal. The button is flat by default, you can use [method set_custom_as_button] to display it with a [StyleBox].
	  This mode also supports custom drawing using [method set_custom_draw_callback].*/
	CellModeCustom TreeCellMode = 4
)
