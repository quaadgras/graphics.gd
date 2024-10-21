package TreeItem

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A single item of a [Tree] control. It can contain other [TreeItem]s as children, which allows it to create a hierarchy. It can also contain text and buttons. [TreeItem] is not a [Node], it is internal to the [Tree].
To create a [TreeItem], use [method Tree.create_item] or [method TreeItem.create_child]. To remove a [TreeItem], use [method Object.free].
[b]Note:[/b] The ID values used for buttons are 32-bit, unlike [int] which is always 64-bit. They go from [code]-2147483648[/code] to [code]2147483647[/code].

*/
type Simple [1]classdb.TreeItem
func (self Simple) SetCellMode(column int, mode classdb.TreeItemTreeCellMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellMode(gd.Int(column), mode)
}
func (self Simple) GetCellMode(column int) classdb.TreeItemTreeCellMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TreeItemTreeCellMode(Expert(self).GetCellMode(gd.Int(column)))
}
func (self Simple) SetEditMultiline(column int, multiline bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditMultiline(gd.Int(column), multiline)
}
func (self Simple) IsEditMultiline(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEditMultiline(gd.Int(column)))
}
func (self Simple) SetChecked(column int, checked bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetChecked(gd.Int(column), checked)
}
func (self Simple) SetIndeterminate(column int, indeterminate bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIndeterminate(gd.Int(column), indeterminate)
}
func (self Simple) IsChecked(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsChecked(gd.Int(column)))
}
func (self Simple) IsIndeterminate(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsIndeterminate(gd.Int(column)))
}
func (self Simple) PropagateCheck(column int, emit_signal bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PropagateCheck(gd.Int(column), emit_signal)
}
func (self Simple) SetText(column int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetText(gd.Int(column), gc.String(text))
}
func (self Simple) GetText(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetText(gc, gd.Int(column)).String())
}
func (self Simple) SetTextDirection(column int, direction classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextDirection(gd.Int(column), direction)
}
func (self Simple) GetTextDirection(column int) classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlTextDirection(Expert(self).GetTextDirection(gd.Int(column)))
}
func (self Simple) SetAutowrapMode(column int, autowrap_mode classdb.TextServerAutowrapMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutowrapMode(gd.Int(column), autowrap_mode)
}
func (self Simple) GetAutowrapMode(column int) classdb.TextServerAutowrapMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerAutowrapMode(Expert(self).GetAutowrapMode(gd.Int(column)))
}
func (self Simple) SetTextOverrunBehavior(column int, overrun_behavior classdb.TextServerOverrunBehavior) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextOverrunBehavior(gd.Int(column), overrun_behavior)
}
func (self Simple) GetTextOverrunBehavior(column int) classdb.TextServerOverrunBehavior {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerOverrunBehavior(Expert(self).GetTextOverrunBehavior(gd.Int(column)))
}
func (self Simple) SetStructuredTextBidiOverride(column int, parser classdb.TextServerStructuredTextParser) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStructuredTextBidiOverride(gd.Int(column), parser)
}
func (self Simple) GetStructuredTextBidiOverride(column int) classdb.TextServerStructuredTextParser {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerStructuredTextParser(Expert(self).GetStructuredTextBidiOverride(gd.Int(column)))
}
func (self Simple) SetStructuredTextBidiOverrideOptions(column int, args gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStructuredTextBidiOverrideOptions(gd.Int(column), args)
}
func (self Simple) GetStructuredTextBidiOverrideOptions(column int) gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetStructuredTextBidiOverrideOptions(gc, gd.Int(column)))
}
func (self Simple) SetLanguage(column int, language string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLanguage(gd.Int(column), gc.String(language))
}
func (self Simple) GetLanguage(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLanguage(gc, gd.Int(column)).String())
}
func (self Simple) SetSuffix(column int, text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSuffix(gd.Int(column), gc.String(text))
}
func (self Simple) GetSuffix(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSuffix(gc, gd.Int(column)).String())
}
func (self Simple) SetIcon(column int, texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIcon(gd.Int(column), texture)
}
func (self Simple) GetIcon(column int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetIcon(gc, gd.Int(column)))
}
func (self Simple) SetIconRegion(column int, region gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIconRegion(gd.Int(column), region)
}
func (self Simple) GetIconRegion(column int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetIconRegion(gd.Int(column)))
}
func (self Simple) SetIconMaxWidth(column int, width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIconMaxWidth(gd.Int(column), gd.Int(width))
}
func (self Simple) GetIconMaxWidth(column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetIconMaxWidth(gd.Int(column))))
}
func (self Simple) SetIconModulate(column int, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIconModulate(gd.Int(column), modulate)
}
func (self Simple) GetIconModulate(column int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetIconModulate(gd.Int(column)))
}
func (self Simple) SetRange(column int, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRange(gd.Int(column), gd.Float(value))
}
func (self Simple) GetRange(column int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRange(gd.Int(column))))
}
func (self Simple) SetRangeConfig(column int, min float64, max float64, step float64, expr bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRangeConfig(gd.Int(column), gd.Float(min), gd.Float(max), gd.Float(step), expr)
}
func (self Simple) GetRangeConfig(column int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetRangeConfig(gc, gd.Int(column)))
}
func (self Simple) SetMetadata(column int, meta gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMetadata(gd.Int(column), meta)
}
func (self Simple) GetMetadata(column int) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).GetMetadata(gc, gd.Int(column)))
}
func (self Simple) SetCustomDraw(column int, obj gd.Object, callback string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomDraw(gd.Int(column), obj, gc.StringName(callback))
}
func (self Simple) SetCustomDrawCallback(column int, callback gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomDrawCallback(gd.Int(column), callback)
}
func (self Simple) GetCustomDrawCallback(column int) gd.Callable {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Callable(Expert(self).GetCustomDrawCallback(gc, gd.Int(column)))
}
func (self Simple) SetCollapsed(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollapsed(enable)
}
func (self Simple) IsCollapsed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCollapsed())
}
func (self Simple) SetCollapsedRecursive(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollapsedRecursive(enable)
}
func (self Simple) IsAnyCollapsed(only_visible bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAnyCollapsed(only_visible))
}
func (self Simple) SetVisible(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisible(enable)
}
func (self Simple) IsVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVisible())
}
func (self Simple) IsVisibleInTree() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVisibleInTree())
}
func (self Simple) UncollapseTree() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UncollapseTree()
}
func (self Simple) SetCustomMinimumHeight(height int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomMinimumHeight(gd.Int(height))
}
func (self Simple) GetCustomMinimumHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCustomMinimumHeight()))
}
func (self Simple) SetSelectable(column int, selectable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelectable(gd.Int(column), selectable)
}
func (self Simple) IsSelectable(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSelectable(gd.Int(column)))
}
func (self Simple) IsSelected(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSelected(gd.Int(column)))
}
func (self Simple) Select(column int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Select(gd.Int(column))
}
func (self Simple) Deselect(column int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Deselect(gd.Int(column))
}
func (self Simple) SetEditable(column int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditable(gd.Int(column), enabled)
}
func (self Simple) IsEditable(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEditable(gd.Int(column)))
}
func (self Simple) SetCustomColor(column int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomColor(gd.Int(column), color)
}
func (self Simple) GetCustomColor(column int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetCustomColor(gd.Int(column)))
}
func (self Simple) ClearCustomColor(column int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearCustomColor(gd.Int(column))
}
func (self Simple) SetCustomFont(column int, font [1]classdb.Font) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomFont(gd.Int(column), font)
}
func (self Simple) GetCustomFont(column int) [1]classdb.Font {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Font(Expert(self).GetCustomFont(gc, gd.Int(column)))
}
func (self Simple) SetCustomFontSize(column int, font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomFontSize(gd.Int(column), gd.Int(font_size))
}
func (self Simple) GetCustomFontSize(column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCustomFontSize(gd.Int(column))))
}
func (self Simple) SetCustomBgColor(column int, color gd.Color, just_outline bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomBgColor(gd.Int(column), color, just_outline)
}
func (self Simple) ClearCustomBgColor(column int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearCustomBgColor(gd.Int(column))
}
func (self Simple) GetCustomBgColor(column int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetCustomBgColor(gd.Int(column)))
}
func (self Simple) SetCustomAsButton(column int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomAsButton(gd.Int(column), enable)
}
func (self Simple) IsCustomSetAsButton(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCustomSetAsButton(gd.Int(column)))
}
func (self Simple) AddButton(column int, button [1]classdb.Texture2D, id int, disabled bool, tooltip_text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddButton(gd.Int(column), button, gd.Int(id), disabled, gc.String(tooltip_text))
}
func (self Simple) GetButtonCount(column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetButtonCount(gd.Int(column))))
}
func (self Simple) GetButtonTooltipText(column int, button_index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetButtonTooltipText(gc, gd.Int(column), gd.Int(button_index)).String())
}
func (self Simple) GetButtonId(column int, button_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetButtonId(gd.Int(column), gd.Int(button_index))))
}
func (self Simple) GetButtonById(column int, id int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetButtonById(gd.Int(column), gd.Int(id))))
}
func (self Simple) GetButtonColor(column int, id int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetButtonColor(gd.Int(column), gd.Int(id)))
}
func (self Simple) GetButton(column int, button_index int) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetButton(gc, gd.Int(column), gd.Int(button_index)))
}
func (self Simple) SetButtonTooltipText(column int, button_index int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetButtonTooltipText(gd.Int(column), gd.Int(button_index), gc.String(tooltip))
}
func (self Simple) SetButton(column int, button_index int, button [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetButton(gd.Int(column), gd.Int(button_index), button)
}
func (self Simple) EraseButton(column int, button_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseButton(gd.Int(column), gd.Int(button_index))
}
func (self Simple) SetButtonDisabled(column int, button_index int, disabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetButtonDisabled(gd.Int(column), gd.Int(button_index), disabled)
}
func (self Simple) SetButtonColor(column int, button_index int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetButtonColor(gd.Int(column), gd.Int(button_index), color)
}
func (self Simple) IsButtonDisabled(column int, button_index int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsButtonDisabled(gd.Int(column), gd.Int(button_index)))
}
func (self Simple) SetTooltipText(column int, tooltip string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTooltipText(gd.Int(column), gc.String(tooltip))
}
func (self Simple) GetTooltipText(column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTooltipText(gc, gd.Int(column)).String())
}
func (self Simple) SetTextAlignment(column int, text_alignment gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextAlignment(gd.Int(column), text_alignment)
}
func (self Simple) GetTextAlignment(column int) gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
	return gd.HorizontalAlignment(Expert(self).GetTextAlignment(gd.Int(column)))
}
func (self Simple) SetExpandRight(column int, enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExpandRight(gd.Int(column), enable)
}
func (self Simple) GetExpandRight(column int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetExpandRight(gd.Int(column)))
}
func (self Simple) SetDisableFolding(disable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableFolding(disable)
}
func (self Simple) IsFoldingDisabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFoldingDisabled())
}
func (self Simple) CreateChild(index int) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).CreateChild(gc, gd.Int(index)))
}
func (self Simple) AddChild(child [1]classdb.TreeItem) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddChild(child)
}
func (self Simple) RemoveChild(child [1]classdb.TreeItem) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveChild(child)
}
func (self Simple) GetTree() [1]classdb.Tree {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Tree(Expert(self).GetTree(gc))
}
func (self Simple) GetNext() [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetNext(gc))
}
func (self Simple) GetPrev() [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetPrev(gc))
}
func (self Simple) GetParent() [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetParent(gc))
}
func (self Simple) GetFirstChild() [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetFirstChild(gc))
}
func (self Simple) GetNextInTree(wrap bool) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetNextInTree(gc, wrap))
}
func (self Simple) GetPrevInTree(wrap bool) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetPrevInTree(gc, wrap))
}
func (self Simple) GetNextVisible(wrap bool) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetNextVisible(gc, wrap))
}
func (self Simple) GetPrevVisible(wrap bool) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetPrevVisible(gc, wrap))
}
func (self Simple) GetChild(index int) [1]classdb.TreeItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TreeItem(Expert(self).GetChild(gc, gd.Int(index)))
}
func (self Simple) GetChildCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetChildCount()))
}
func (self Simple) GetChildren() gd.ArrayOf[[1]classdb.TreeItem] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.TreeItem](Expert(self).GetChildren(gc))
}
func (self Simple) GetIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetIndex()))
}
func (self Simple) MoveBefore(item [1]classdb.TreeItem) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveBefore(item)
}
func (self Simple) MoveAfter(item [1]classdb.TreeItem) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveAfter(item)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TreeItem
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the given column's cell mode to [param mode]. This determines how the cell is displayed and edited. See [enum TreeCellMode] constants for details.
*/
//go:nosplit
func (self class) SetCellMode(column gd.Int, mode classdb.TreeItemTreeCellMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_cell_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the column's cell mode.
*/
//go:nosplit
func (self class) GetCellMode(column gd.Int) classdb.TreeItemTreeCellMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TreeItemTreeCellMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_cell_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param multiline] is [code]true[/code], the given [param column] is multiline editable.
[b]Note:[/b] This option only affects the type of control ([LineEdit] or [TextEdit]) that appears when editing the column. You can set multiline values with [method set_text] even if the column is not multiline editable.
*/
//go:nosplit
func (self class) SetEditMultiline(column gd.Int, multiline bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, multiline)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_edit_multiline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param column] is multiline editable.
*/
//go:nosplit
func (self class) IsEditMultiline(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_edit_multiline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param checked] is [code]true[/code], the given [param column] is checked. Clears column's indeterminate status.
*/
//go:nosplit
func (self class) SetChecked(column gd.Int, checked bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, checked)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param indeterminate] is [code]true[/code], the given [param column] is marked indeterminate.
[b]Note:[/b] If set [code]true[/code] from [code]false[/code], then column is cleared of checked status.
*/
//go:nosplit
func (self class) SetIndeterminate(column gd.Int, indeterminate bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, indeterminate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param column] is checked.
*/
//go:nosplit
func (self class) IsChecked(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_checked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given [param column] is indeterminate.
*/
//go:nosplit
func (self class) IsIndeterminate(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Propagates this item's checked status to its children and parents for the given [param column]. It is possible to process the items affected by this method call by connecting to [signal Tree.check_propagated_to_item]. The order that the items affected will be processed is as follows: the item invoking this method, children of that item, and finally parents of that item. If [param emit_signal] is [code]false[/code], then [signal Tree.check_propagated_to_item] will not be emitted.
*/
//go:nosplit
func (self class) PropagateCheck(column gd.Int, emit_signal bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, emit_signal)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_propagate_check, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given column's text value.
*/
//go:nosplit
func (self class) SetText(column gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given column's text.
*/
//go:nosplit
func (self class) GetText(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetTextDirection(column gd.Int, direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetTextDirection(column gd.Int) classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the autowrap mode in the given [param column]. If set to something other than [constant TextServer.AUTOWRAP_OFF], the text gets wrapped inside the cell's bounding rectangle.
*/
//go:nosplit
func (self class) SetAutowrapMode(column gd.Int, autowrap_mode classdb.TextServerAutowrapMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the text autowrap mode in the given [param column]. By default it is [constant TextServer.AUTOWRAP_OFF].
*/
//go:nosplit
func (self class) GetAutowrapMode(column gd.Int) classdb.TextServerAutowrapMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column].
*/
//go:nosplit
func (self class) SetTextOverrunBehavior(column gd.Int, overrun_behavior classdb.TextServerOverrunBehavior)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column]. By default it is [constant TextServer.OVERRUN_TRIM_ELLIPSIS].
*/
//go:nosplit
func (self class) GetTextOverrunBehavior(column gd.Int) classdb.TextServerOverrunBehavior {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set BiDi algorithm override for the structured text. Has effect for cells that display text.
*/
//go:nosplit
func (self class) SetStructuredTextBidiOverride(column gd.Int, parser classdb.TextServerStructuredTextParser)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the BiDi algorithm override set for this cell.
*/
//go:nosplit
func (self class) GetStructuredTextBidiOverride(column gd.Int) classdb.TextServerStructuredTextParser {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Set additional options for BiDi override. Has effect for cells that display text.
*/
//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(column gd.Int, args gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(args))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the additional BiDi options set for this cell.
*/
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(ctx gd.Lifetime, column gd.Int) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetLanguage(column gd.Int, language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetLanguage(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets a string to be shown after a column's value (for example, a unit abbreviation).
*/
//go:nosplit
func (self class) SetSuffix(column gd.Int, text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets the suffix string shown after the column value.
*/
//go:nosplit
func (self class) GetSuffix(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_suffix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the given cell's icon [Texture2D]. The cell has to be in [constant CELL_MODE_ICON] mode.
*/
//go:nosplit
func (self class) SetIcon(column gd.Int, texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given column's icon [Texture2D]. Error if no icon is set.
*/
//go:nosplit
func (self class) GetIcon(ctx gd.Lifetime, column gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the given column's icon's texture region.
*/
//go:nosplit
func (self class) SetIconRegion(column gd.Int, region gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, region)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the icon [Texture2D] region as [Rect2].
*/
//go:nosplit
func (self class) GetIconRegion(column gd.Int) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_icon_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the maximum allowed width of the icon in the given [param column]. This limit is applied on top of the default size of the icon and on top of [theme_item Tree.icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetIconMaxWidth(column gd.Int, width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the maximum allowed width of the icon in the given [param column].
*/
//go:nosplit
func (self class) GetIconMaxWidth(column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Modulates the given column's icon with [param modulate].
*/
//go:nosplit
func (self class) SetIconModulate(column gd.Int, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Color] modulating the column's icon.
*/
//go:nosplit
func (self class) GetIconModulate(column gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of a [constant CELL_MODE_RANGE] column.
*/
//go:nosplit
func (self class) SetRange(column gd.Int, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of a [constant CELL_MODE_RANGE] column.
*/
//go:nosplit
func (self class) GetRange(column gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the range of accepted values for a column. The column must be in the [constant CELL_MODE_RANGE] mode.
If [param expr] is [code]true[/code], the edit mode slider will use an exponential scale as with [member Range.exp_edit].
*/
//go:nosplit
func (self class) SetRangeConfig(column gd.Int, min gd.Float, max gd.Float, step gd.Float, expr bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, min)
	callframe.Arg(frame, max)
	callframe.Arg(frame, step)
	callframe.Arg(frame, expr)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_range_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a dictionary containing the range parameters for a given column. The keys are "min", "max", "step", and "expr".
*/
//go:nosplit
func (self class) GetRangeConfig(ctx gd.Lifetime, column gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_range_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the metadata value for the given column, which can be retrieved later using [method get_metadata]. This can be used, for example, to store a reference to the original data.
*/
//go:nosplit
func (self class) SetMetadata(column gd.Int, meta gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(meta))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the metadata value that was set for the given column using [method set_metadata].
*/
//go:nosplit
func (self class) GetMetadata(ctx gd.Lifetime, column gd.Int) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the given column's custom draw callback to the [param callback] method on [param object].
The method named [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
//go:nosplit
func (self class) SetCustomDraw(column gd.Int, obj gd.Object, callback gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given column's custom draw callback. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback. The cell has to be in [constant CELL_MODE_CUSTOM] to use this feature.
The [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
//go:nosplit
func (self class) SetCustomDrawCallback(column gd.Int, callback gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(callback))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_draw_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom callback of column [param column].
*/
//go:nosplit
func (self class) GetCustomDrawCallback(ctx gd.Lifetime, column gd.Int) gd.Callable {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_draw_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollapsed(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollapsed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Collapses or uncollapses this [TreeItem] and all the descendants of this item.
*/
//go:nosplit
func (self class) SetCollapsedRecursive(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_collapsed_recursive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if this [TreeItem], or any of its descendants, is collapsed.
If [param only_visible] is [code]true[/code] it ignores non-visible [TreeItem]s.
*/
//go:nosplit
func (self class) IsAnyCollapsed(only_visible bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, only_visible)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_any_collapsed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisible(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [member visible] is [code]true[/code] and all its ancestors are also visible.
*/
//go:nosplit
func (self class) IsVisibleInTree() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_visible_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Uncollapses all [TreeItem]s necessary to reveal this [TreeItem], i.e. all ancestor [TreeItem]s.
*/
//go:nosplit
func (self class) UncollapseTree()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_uncollapse_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCustomMinimumHeight(height gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_minimum_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomMinimumHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_minimum_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param selectable] is [code]true[/code], the given [param column] is selectable.
*/
//go:nosplit
func (self class) SetSelectable(column gd.Int, selectable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, selectable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param column] is selectable.
*/
//go:nosplit
func (self class) IsSelectable(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_selectable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given [param column] is selected.
*/
//go:nosplit
func (self class) IsSelected(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Selects the given [param column].
*/
//go:nosplit
func (self class) Select(column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deselects the given column.
*/
//go:nosplit
func (self class) Deselect(column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [param enabled] is [code]true[/code], the given [param column] is editable.
*/
//go:nosplit
func (self class) SetEditable(column gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given [param column] is editable.
*/
//go:nosplit
func (self class) IsEditable(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the given column's custom color.
*/
//go:nosplit
func (self class) SetCustomColor(column gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom color of column [param column].
*/
//go:nosplit
func (self class) GetCustomColor(column gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Resets the color for the given column to default.
*/
//go:nosplit
func (self class) ClearCustomColor(column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_clear_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets custom font used to draw text in the given [param column].
*/
//go:nosplit
func (self class) SetCustomFont(column gd.Int, font object.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns custom font used to draw text in the column [param column].
*/
//go:nosplit
func (self class) GetCustomFont(ctx gd.Lifetime, column gd.Int) object.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets custom font size used to draw text in the given [param column].
*/
//go:nosplit
func (self class) SetCustomFontSize(column gd.Int, font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns custom font size used to draw text in the column [param column].
*/
//go:nosplit
func (self class) GetCustomFontSize(column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the given column's custom background color and whether to just use it as an outline.
*/
//go:nosplit
func (self class) SetCustomBgColor(column gd.Int, color gd.Color, just_outline bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, color)
	callframe.Arg(frame, just_outline)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Resets the background color for the given column to default.
*/
//go:nosplit
func (self class) ClearCustomBgColor(column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_clear_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom background color of column [param column].
*/
//go:nosplit
func (self class) GetCustomBgColor(column gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Makes a cell with [constant CELL_MODE_CUSTOM] display as a non-flat button with a [StyleBox].
*/
//go:nosplit
func (self class) SetCustomAsButton(column gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_custom_as_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the cell was made into a button with [method set_custom_as_button].
*/
//go:nosplit
func (self class) IsCustomSetAsButton(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_custom_set_as_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a button with [Texture2D] [param button] at column [param column]. The [param id] is used to identify the button in the according [signal Tree.button_clicked] signal and can be different from the buttons index. If not specified, the next available index is used, which may be retrieved by calling [method get_button_count] immediately before this method. Optionally, the button can be [param disabled] and have a [param tooltip_text].
*/
//go:nosplit
func (self class) AddButton(column gd.Int, button object.Texture2D, id gd.Int, disabled bool, tooltip_text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(button[0].AsPointer())[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, disabled)
	callframe.Arg(frame, mmm.Get(tooltip_text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_add_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of buttons in column [param column].
*/
//go:nosplit
func (self class) GetButtonCount(column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tooltip text for the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButtonTooltipText(ctx gd.Lifetime, column gd.Int, button_index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the ID for the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButtonId(column gd.Int, button_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the button index if there is a button with ID [param id] in column [param column], otherwise returns -1.
*/
//go:nosplit
func (self class) GetButtonById(column gd.Int, id gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_by_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the color of the button with ID [param id] in column [param column]. If the specified button does not exist, returns [constant Color.BLACK].
*/
//go:nosplit
func (self class) GetButtonColor(column gd.Int, id gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Texture2D] of the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButton(ctx gd.Lifetime, column gd.Int, button_index gd.Int) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the tooltip text for the button at index [param button_index] in the given [param column].
*/
//go:nosplit
func (self class) SetButtonTooltipText(column gd.Int, button_index gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_button_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given column's button [Texture2D] at index [param button_index] to [param button].
*/
//go:nosplit
func (self class) SetButton(column gd.Int, button_index gd.Int, button object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, mmm.Get(button[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) EraseButton(column gd.Int, button_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_erase_button, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], disables the button at index [param button_index] in the given [param column].
*/
//go:nosplit
func (self class) SetButtonDisabled(column gd.Int, button_index gd.Int, disabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_button_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the given column's button color at index [param button_index] to [param color].
*/
//go:nosplit
func (self class) SetButtonColor(column gd.Int, button_index gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_button_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the button at index [param button_index] for the given [param column] is disabled.
*/
//go:nosplit
func (self class) IsButtonDisabled(column gd.Int, button_index gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_button_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the given column's tooltip text.
*/
//go:nosplit
func (self class) SetTooltipText(column gd.Int, tooltip gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mmm.Get(tooltip))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given column's tooltip text.
*/
//go:nosplit
func (self class) GetTooltipText(ctx gd.Lifetime, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the given column's text alignment. See [enum HorizontalAlignment] for possible values.
*/
//go:nosplit
func (self class) SetTextAlignment(column gd.Int, text_alignment gd.HorizontalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, text_alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_text_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given column's text alignment.
*/
//go:nosplit
func (self class) GetTextAlignment(column gd.Int) gd.HorizontalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_text_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param enable] is [code]true[/code], the given [param column] is expanded to the right.
*/
//go:nosplit
func (self class) SetExpandRight(column gd.Int, enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_expand_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if [code]expand_right[/code] is set.
*/
//go:nosplit
func (self class) GetExpandRight(column gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_expand_right, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisableFolding(disable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_set_disable_folding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFoldingDisabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_is_folding_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an item and adds it as a child.
The new item will be inserted as position [param index] (the default value [code]-1[/code] means the last position), or it will be the last child if [param index] is higher than the child count.
*/
//go:nosplit
func (self class) CreateChild(ctx gd.Lifetime, index gd.Int) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_create_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adds a previously unparented [TreeItem] as a direct child of this one. The [param child] item must not be a part of any [Tree] or parented to any [TreeItem]. See also [method remove_child].
*/
//go:nosplit
func (self class) AddChild(child object.TreeItem)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(child[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_add_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given child [TreeItem] and all its children from the [Tree]. Note that it doesn't free the item from memory, so it can be reused later (see [method add_child]). To completely remove a [TreeItem] use [method Object.free].
[b]Note:[/b] If you want to move a child from one [Tree] to another, then instead of removing and adding it manually you can use [method move_before] or [method move_after].
*/
//go:nosplit
func (self class) RemoveChild(child object.TreeItem)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(child[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_remove_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Tree] that owns this TreeItem.
*/
//go:nosplit
func (self class) GetTree(ctx gd.Lifetime) object.Tree {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Tree
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the next sibling TreeItem in the tree or a null object if there is none.
*/
//go:nosplit
func (self class) GetNext(ctx gd.Lifetime) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the previous sibling TreeItem in the tree or a null object if there is none.
*/
//go:nosplit
func (self class) GetPrev(ctx gd.Lifetime) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_prev, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the parent TreeItem or a null object if there is none.
*/
//go:nosplit
func (self class) GetParent(ctx gd.Lifetime) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the TreeItem's first child.
*/
//go:nosplit
func (self class) GetFirstChild(ctx gd.Lifetime) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_first_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the next TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first element in the tree when called on the last element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetNextInTree(ctx gd.Lifetime, wrap bool) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_next_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the previous TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetPrevInTree(ctx gd.Lifetime, wrap bool) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_prev_in_tree, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the next visible TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first visible element in the tree when called on the last visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetNextVisible(ctx gd.Lifetime, wrap bool) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_next_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the previous visible sibling TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last visible element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetPrevVisible(ctx gd.Lifetime, wrap bool) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_prev_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a child item by its [param index] (see [method get_child_count]). This method is often used for iterating all children of an item.
Negative indices access the children from the last one.
*/
//go:nosplit
func (self class) GetChild(ctx gd.Lifetime, index gd.Int) object.TreeItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_child, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TreeItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the number of child items.
*/
//go:nosplit
func (self class) GetChildCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_child_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of references to the item's children.
*/
//go:nosplit
func (self class) GetChildren(ctx gd.Lifetime) gd.ArrayOf[object.TreeItem] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_children, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.TreeItem](ret)
}
/*
Returns the node's order in the tree. For example, if called on the first child item the position is [code]0[/code].
*/
//go:nosplit
func (self class) GetIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves this TreeItem right before the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
//go:nosplit
func (self class) MoveBefore(item object.TreeItem)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(item[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_move_before, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves this TreeItem right after the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
//go:nosplit
func (self class) MoveAfter(item object.TreeItem)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(item[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TreeItem.Bind_move_after, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsTreeItem() Expert { return self[0].AsTreeItem() }


//go:nosplit
func (self Simple) AsTreeItem() Simple { return self[0].AsTreeItem() }

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
func init() {classdb.Register("TreeItem", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
