package TextEdit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Rect2"
import "grow.graphics/gd/objects/Control"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"
import "grow.graphics/gd/variant/Array"
import "grow.graphics/gd/variant/Vector2i"
import "grow.graphics/gd/variant/Vector2"
import "grow.graphics/gd/variant/Rect2i"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Color"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A multiline text editor. It also has limited facilities for editing code, such as syntax highlighting support. For more advanced facilities for editing code, see [CodeEdit].
[b]Note:[/b] Most viewport, caret, and edit methods contain a [code]caret_index[/code] argument for [member caret_multiple] support. The argument should be one of the following: [code]-1[/code] for all carets, [code]0[/code] for the main caret, or greater than [code]0[/code] for secondary carets in the order they were created.
[b]Note:[/b] When holding down [kbd]Alt[/kbd], the vertical scroll wheel will scroll 5 times as fast as it would normally do. This also works in the Godot script editor.

	// TextEdit methods that can be overridden by a [Class] that extends it.
	type TextEdit interface {
		//Override this method to define what happens when the user types in the provided key [param unicode_char].
		HandleUnicodeInput(unicode_char int, caret_index int)
		//Override this method to define what happens when the user presses the backspace key.
		Backspace(caret_index int)
		//Override this method to define what happens when the user performs a cut operation.
		Cut(caret_index int)
		//Override this method to define what happens when the user performs a copy operation.
		Copy(caret_index int)
		//Override this method to define what happens when the user performs a paste operation.
		Paste(caret_index int)
		//Override this method to define what happens when the user performs a paste operation with middle mouse button.
		//[b]Note:[/b] This method is only implemented on Linux.
		PastePrimaryClipboard(caret_index int)
	}
*/
type Instance [1]classdb.TextEdit

/*
Override this method to define what happens when the user types in the provided key [param unicode_char].
*/
func (Instance) _handle_unicode_input(impl func(ptr unsafe.Pointer, unicode_char int, caret_index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var unicode_char = gd.UnsafeGet[gd.Int](p_args, 0)
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(unicode_char), int(caret_index))
	}
}

/*
Override this method to define what happens when the user presses the backspace key.
*/
func (Instance) _backspace(impl func(ptr unsafe.Pointer, caret_index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(caret_index))
	}
}

/*
Override this method to define what happens when the user performs a cut operation.
*/
func (Instance) _cut(impl func(ptr unsafe.Pointer, caret_index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(caret_index))
	}
}

/*
Override this method to define what happens when the user performs a copy operation.
*/
func (Instance) _copy(impl func(ptr unsafe.Pointer, caret_index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(caret_index))
	}
}

/*
Override this method to define what happens when the user performs a paste operation.
*/
func (Instance) _paste(impl func(ptr unsafe.Pointer, caret_index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(caret_index))
	}
}

/*
Override this method to define what happens when the user performs a paste operation with middle mouse button.
[b]Note:[/b] This method is only implemented on Linux.
*/
func (Instance) _paste_primary_clipboard(impl func(ptr unsafe.Pointer, caret_index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(caret_index))
	}
}

/*
Returns [code]true[/code] if the user has text in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME).
*/
func (self Instance) HasImeText() bool {
	return bool(class(self).HasImeText())
}

/*
Closes the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME) if it is open. Any text in the IME will be lost.
*/
func (self Instance) CancelIme() {
	class(self).CancelIme()
}

/*
Applies text from the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME) to each caret and closes the IME if it is open.
*/
func (self Instance) ApplyIme() {
	class(self).ApplyIme()
}

/*
Sets the tab size for the [TextEdit] to use.
*/
func (self Instance) SetTabSize(size int) {
	class(self).SetTabSize(gd.Int(size))
}

/*
Returns the [TextEdit]'s' tab size.
*/
func (self Instance) GetTabSize() int {
	return int(int(class(self).GetTabSize()))
}

/*
If [code]true[/code], sets the user into overtype mode. When the user types in this mode, it will override existing text.
*/
func (self Instance) SetOvertypeModeEnabled(enabled bool) {
	class(self).SetOvertypeModeEnabled(enabled)
}

/*
Returns whether the user is in overtype mode.
*/
func (self Instance) IsOvertypeModeEnabled() bool {
	return bool(class(self).IsOvertypeModeEnabled())
}

/*
Performs a full reset of [TextEdit], including undo history.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Returns the number of lines in the text.
*/
func (self Instance) GetLineCount() int {
	return int(int(class(self).GetLineCount()))
}

/*
Sets the text for a specific [param line].
Carets on the line will attempt to keep their visual x position.
*/
func (self Instance) SetLine(line int, new_text string) {
	class(self).SetLine(gd.Int(line), gd.NewString(new_text))
}

/*
Returns the text of a specific line.
*/
func (self Instance) GetLine(line int) string {
	return string(class(self).GetLine(gd.Int(line)).String())
}

/*
Returns the width in pixels of the [param wrap_index] on [param line].
*/
func (self Instance) GetLineWidth(line int) int {
	return int(int(class(self).GetLineWidth(gd.Int(line), gd.Int(-1))))
}

/*
Returns the maximum value of the line height among all lines.
[b]Note:[/b] The return value is influenced by [theme_item line_spacing] and [theme_item font_size]. And it will not be less than [code]1[/code].
*/
func (self Instance) GetLineHeight() int {
	return int(int(class(self).GetLineHeight()))
}

/*
Returns the number of spaces and [code]tab * tab_size[/code] before the first char.
*/
func (self Instance) GetIndentLevel(line int) int {
	return int(int(class(self).GetIndentLevel(gd.Int(line))))
}

/*
Returns the first column containing a non-whitespace character.
*/
func (self Instance) GetFirstNonWhitespaceColumn(line int) int {
	return int(int(class(self).GetFirstNonWhitespaceColumn(gd.Int(line))))
}

/*
Swaps the two lines. Carets will be swapped with the lines.
*/
func (self Instance) SwapLines(from_line int, to_line int) {
	class(self).SwapLines(gd.Int(from_line), gd.Int(to_line))
}

/*
Inserts a new line with [param text] at [param line].
*/
func (self Instance) InsertLineAt(line int, text string) {
	class(self).InsertLineAt(gd.Int(line), gd.NewString(text))
}

/*
Removes the line of text at [param line]. Carets on this line will attempt to match their previous visual x position.
If [param move_carets_down] is [code]true[/code] carets will move to the next line down, otherwise carets will move up.
*/
func (self Instance) RemoveLineAt(line int) {
	class(self).RemoveLineAt(gd.Int(line), true)
}

/*
Insert the specified text at the caret position.
*/
func (self Instance) InsertTextAtCaret(text string) {
	class(self).InsertTextAtCaret(gd.NewString(text), gd.Int(-1))
}

/*
Inserts the [param text] at [param line] and [param column].
If [param before_selection_begin] is [code]true[/code], carets and selections that begin at [param line] and [param column] will moved to the end of the inserted text, along with all carets after it.
If [param before_selection_end] is [code]true[/code], selections that end at [param line] and [param column] will be extended to the end of the inserted text. These parameters can be used to insert text inside of or outside of selections.
*/
func (self Instance) InsertText(text string, line int, column int) {
	class(self).InsertText(gd.NewString(text), gd.Int(line), gd.Int(column), true, false)
}

/*
Removes text between the given positions.
*/
func (self Instance) RemoveText(from_line int, from_column int, to_line int, to_column int) {
	class(self).RemoveText(gd.Int(from_line), gd.Int(from_column), gd.Int(to_line), gd.Int(to_column))
}

/*
Returns the last unhidden line in the entire [TextEdit].
*/
func (self Instance) GetLastUnhiddenLine() int {
	return int(int(class(self).GetLastUnhiddenLine()))
}

/*
Returns the count to the next visible line from [param line] to [code]line + visible_amount[/code]. Can also count backwards. For example if a [TextEdit] has 5 lines with lines 2 and 3 hidden, calling this with [code]line = 1, visible_amount = 1[/code] would return 3.
*/
func (self Instance) GetNextVisibleLineOffsetFrom(line int, visible_amount int) int {
	return int(int(class(self).GetNextVisibleLineOffsetFrom(gd.Int(line), gd.Int(visible_amount))))
}

/*
Similar to [method get_next_visible_line_offset_from], but takes into account the line wrap indexes. In the returned vector, [code]x[/code] is the line, [code]y[/code] is the wrap index.
*/
func (self Instance) GetNextVisibleLineIndexOffsetFrom(line int, wrap_index int, visible_amount int) Vector2i.XY {
	return Vector2i.XY(class(self).GetNextVisibleLineIndexOffsetFrom(gd.Int(line), gd.Int(wrap_index), gd.Int(visible_amount)))
}

/*
Called when the user presses the backspace key. Can be overridden with [method _backspace].
*/
func (self Instance) Backspace() {
	class(self).Backspace(gd.Int(-1))
}

/*
Cut's the current selection. Can be overridden with [method _cut].
*/
func (self Instance) Cut() {
	class(self).Cut(gd.Int(-1))
}

/*
Copies the current text selection. Can be overridden with [method _copy].
*/
func (self Instance) Copy() {
	class(self).Copy(gd.Int(-1))
}

/*
Paste at the current location. Can be overridden with [method _paste].
*/
func (self Instance) Paste() {
	class(self).Paste(gd.Int(-1))
}

/*
Pastes the primary clipboard.
*/
func (self Instance) PastePrimaryClipboard() {
	class(self).PastePrimaryClipboard(gd.Int(-1))
}

/*
Starts an action, will end the current action if [param action] is different.
An action will also end after a call to [method end_action], after [member ProjectSettings.gui/timers/text_edit_idle_detect_sec] is triggered or a new undoable step outside the [method start_action] and [method end_action] calls.
*/
func (self Instance) StartAction(action classdb.TextEditEditAction) {
	class(self).StartAction(action)
}

/*
Marks the end of steps in the current action started with [method start_action].
*/
func (self Instance) EndAction() {
	class(self).EndAction()
}

/*
Starts a multipart edit. All edits will be treated as one action until [method end_complex_operation] is called.
*/
func (self Instance) BeginComplexOperation() {
	class(self).BeginComplexOperation()
}

/*
Ends a multipart edit, started with [method begin_complex_operation]. If called outside a complex operation, the current operation is pushed onto the undo/redo stack.
*/
func (self Instance) EndComplexOperation() {
	class(self).EndComplexOperation()
}

/*
Returns [code]true[/code] if an "undo" action is available.
*/
func (self Instance) HasUndo() bool {
	return bool(class(self).HasUndo())
}

/*
Returns [code]true[/code] if a "redo" action is available.
*/
func (self Instance) HasRedo() bool {
	return bool(class(self).HasRedo())
}

/*
Perform undo operation.
*/
func (self Instance) Undo() {
	class(self).Undo()
}

/*
Perform redo operation.
*/
func (self Instance) Redo() {
	class(self).Redo()
}

/*
Clears the undo history.
*/
func (self Instance) ClearUndoHistory() {
	class(self).ClearUndoHistory()
}

/*
Tag the current version as saved.
*/
func (self Instance) TagSavedVersion() {
	class(self).TagSavedVersion()
}

/*
Returns the current version of the [TextEdit]. The version is a count of recorded operations by the undo/redo history.
*/
func (self Instance) GetVersion() int {
	return int(int(class(self).GetVersion()))
}

/*
Returns the last tagged saved version from [method tag_saved_version].
*/
func (self Instance) GetSavedVersion() int {
	return int(int(class(self).GetSavedVersion()))
}

/*
Sets the search text. See [method set_search_flags].
*/
func (self Instance) SetSearchText(search_text string) {
	class(self).SetSearchText(gd.NewString(search_text))
}

/*
Sets the search [param flags]. This is used with [method set_search_text] to highlight occurrences of the searched text. Search flags can be specified from the [enum SearchFlags] enum.
*/
func (self Instance) SetSearchFlags(flags int) {
	class(self).SetSearchFlags(gd.Int(flags))
}

/*
Perform a search inside the text. Search flags can be specified in the [enum SearchFlags] enum.
In the returned vector, [code]x[/code] is the column, [code]y[/code] is the line. If no results are found, both are equal to [code]-1[/code].
[codeblocks]
[gdscript]
var result = search("print", SEARCH_WHOLE_WORDS, 0, 0)
if result.x != -1:

	# Result found.
	var line_number = result.y
	var column_number = result.x

[/gdscript]
[csharp]
Vector2I result = Search("print", (uint)TextEdit.SearchFlags.WholeWords, 0, 0);
if (result.X != -1)

	{
	    // Result found.
	    int lineNumber = result.Y;
	    int columnNumber = result.X;
	}

[/csharp]
[/codeblocks]
*/
func (self Instance) Search(text string, flags int, from_line int, from_column int) Vector2i.XY {
	return Vector2i.XY(class(self).Search(gd.NewString(text), gd.Int(flags), gd.Int(from_line), gd.Int(from_column)))
}

/*
Provide custom tooltip text. The callback method must take the following args: [code]hovered_word: String[/code].
*/
func (self Instance) SetTooltipRequestFunc(callback func(hovered_word string) string) {
	class(self).SetTooltipRequestFunc(gd.NewCallable(callback))
}

/*
Returns the local mouse position adjusted for the text direction.
*/
func (self Instance) GetLocalMousePos() Vector2.XY {
	return Vector2.XY(class(self).GetLocalMousePos())
}

/*
Returns the word at [param position].
*/
func (self Instance) GetWordAtPos(position Vector2.XY) string {
	return string(class(self).GetWordAtPos(gd.Vector2(position)).String())
}

/*
Returns the line and column at the given position. In the returned vector, [code]x[/code] is the column, [code]y[/code] is the line. If [param allow_out_of_bounds] is [code]false[/code] and the position is not over the text, both vector values will be set to [code]-1[/code].
*/
func (self Instance) GetLineColumnAtPos(position Vector2i.XY) Vector2i.XY {
	return Vector2i.XY(class(self).GetLineColumnAtPos(gd.Vector2i(position), true))
}

/*
Returns the local position for the given [param line] and [param column]. If [code]x[/code] or [code]y[/code] of the returned vector equal [code]-1[/code], the position is outside of the viewable area of the control.
[b]Note:[/b] The Y position corresponds to the bottom side of the line. Use [method get_rect_at_line_column] to get the top side position.
*/
func (self Instance) GetPosAtLineColumn(line int, column int) Vector2i.XY {
	return Vector2i.XY(class(self).GetPosAtLineColumn(gd.Int(line), gd.Int(column)))
}

/*
Returns the local position and size for the grapheme at the given [param line] and [param column]. If [code]x[/code] or [code]y[/code] position of the returned rect equal [code]-1[/code], the position is outside of the viewable area of the control.
[b]Note:[/b] The Y position of the returned rect corresponds to the top side of the line, unlike [method get_pos_at_line_column] which returns the bottom side.
*/
func (self Instance) GetRectAtLineColumn(line int, column int) Rect2i.PositionSize {
	return Rect2i.PositionSize(class(self).GetRectAtLineColumn(gd.Int(line), gd.Int(column)))
}

/*
Returns the equivalent minimap line at [param position].
*/
func (self Instance) GetMinimapLineAtPos(position Vector2i.XY) int {
	return int(int(class(self).GetMinimapLineAtPos(gd.Vector2i(position))))
}

/*
Returns [code]true[/code] if the user is dragging their mouse for scrolling, selecting, or text dragging.
*/
func (self Instance) IsDraggingCursor() bool {
	return bool(class(self).IsDraggingCursor())
}

/*
Returns whether the mouse is over selection. If [param edges] is [code]true[/code], the edges are considered part of the selection.
*/
func (self Instance) IsMouseOverSelection(edges bool) bool {
	return bool(class(self).IsMouseOverSelection(edges, gd.Int(-1)))
}

/*
Adds a new caret at the given location. Returns the index of the new caret, or [code]-1[/code] if the location is invalid.
*/
func (self Instance) AddCaret(line int, column int) int {
	return int(int(class(self).AddCaret(gd.Int(line), gd.Int(column))))
}

/*
Removes the given caret index.
[b]Note:[/b] This can result in adjustment of all other caret indices.
*/
func (self Instance) RemoveCaret(caret int) {
	class(self).RemoveCaret(gd.Int(caret))
}

/*
Removes all additional carets.
*/
func (self Instance) RemoveSecondaryCarets() {
	class(self).RemoveSecondaryCarets()
}

/*
Returns the number of carets in this [TextEdit].
*/
func (self Instance) GetCaretCount() int {
	return int(int(class(self).GetCaretCount()))
}

/*
Adds an additional caret above or below every caret. If [param below] is [code]true[/code] the new caret will be added below and above otherwise.
*/
func (self Instance) AddCaretAtCarets(below bool) {
	class(self).AddCaretAtCarets(below)
}

/*
Returns the carets sorted by selection beginning from lowest line and column to highest (from top to bottom of text).
If [param include_ignored_carets] is [code]false[/code], carets from [method multicaret_edit_ignore_caret] will be ignored.
*/
func (self Instance) GetSortedCarets() []int32 {
	return []int32(class(self).GetSortedCarets(false).AsSlice())
}

/*
Collapse all carets in the given range to the [param from_line] and [param from_column] position.
[param inclusive] applies to both ends.
If [method is_in_mulitcaret_edit] is [code]true[/code], carets that are collapsed will be [code]true[/code] for [method multicaret_edit_ignore_caret].
[method merge_overlapping_carets] will be called if any carets were collapsed.
*/
func (self Instance) CollapseCarets(from_line int, from_column int, to_line int, to_column int) {
	class(self).CollapseCarets(gd.Int(from_line), gd.Int(from_column), gd.Int(to_line), gd.Int(to_column), false)
}

/*
Merges any overlapping carets. Will favor the newest caret, or the caret with a selection.
If [method is_in_mulitcaret_edit] is [code]true[/code], the merge will be queued to happen at the end of the multicaret edit. See [method begin_multicaret_edit] and [method end_multicaret_edit].
[b]Note:[/b] This is not called when a caret changes position but after certain actions, so it is possible to get into a state where carets overlap.
*/
func (self Instance) MergeOverlappingCarets() {
	class(self).MergeOverlappingCarets()
}

/*
Starts an edit for multiple carets. The edit must be ended with [method end_multicaret_edit]. Multicaret edits can be used to edit text at multiple carets and delay merging the carets until the end, so the caret indexes aren't affected immediately. [method begin_multicaret_edit] and [method end_multicaret_edit] can be nested, and the merge will happen at the last [method end_multicaret_edit].
Example usage:
[codeblock]
begin_complex_operation()
begin_multicaret_edit()
for i in range(get_caret_count()):

	if multicaret_edit_ignore_caret(i):
	    continue
	# Logic here.

end_multicaret_edit()
end_complex_operation()
[/codeblock]
*/
func (self Instance) BeginMulticaretEdit() {
	class(self).BeginMulticaretEdit()
}

/*
Ends an edit for multiple carets, that was started with [method begin_multicaret_edit]. If this was the last [method end_multicaret_edit] and [method merge_overlapping_carets] was called, carets will be merged.
*/
func (self Instance) EndMulticaretEdit() {
	class(self).EndMulticaretEdit()
}

/*
Returns [code]true[/code] if a [method begin_multicaret_edit] has been called and [method end_multicaret_edit] has not yet been called.
*/
func (self Instance) IsInMulitcaretEdit() bool {
	return bool(class(self).IsInMulitcaretEdit())
}

/*
Returns [code]true[/code] if the given [param caret_index] should be ignored as part of a multicaret edit. See [method begin_multicaret_edit] and [method end_multicaret_edit]. Carets that should be ignored are ones that were part of removed text and will likely be merged at the end of the edit, or carets that were added during the edit.
It is recommended to [code]continue[/code] within a loop iterating on multiple carets if a caret should be ignored.
*/
func (self Instance) MulticaretEditIgnoreCaret(caret_index int) bool {
	return bool(class(self).MulticaretEditIgnoreCaret(gd.Int(caret_index)))
}

/*
Returns [code]true[/code] if the caret is visible on the screen.
*/
func (self Instance) IsCaretVisible() bool {
	return bool(class(self).IsCaretVisible(gd.Int(0)))
}

/*
Returns the caret pixel draw position.
*/
func (self Instance) GetCaretDrawPos() Vector2.XY {
	return Vector2.XY(class(self).GetCaretDrawPos(gd.Int(0)))
}

/*
Moves the caret to the specified [param line] index. The caret column will be moved to the same visual position it was at the last time [method set_caret_column] was called, or clamped to the end of the line.
If [param adjust_viewport] is [code]true[/code], the viewport will center at the caret position after the move occurs.
If [param can_be_hidden] is [code]true[/code], the specified [param line] can be hidden.
If [param wrap_index] is [code]-1[/code], the caret column will be clamped to the [param line]'s length. If [param wrap_index] is greater than [code]-1[/code], the column will be moved to attempt to match the visual x position on the line's [param wrap_index] to the position from the last time [method set_caret_column] was called.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
func (self Instance) SetCaretLine(line int) {
	class(self).SetCaretLine(gd.Int(line), true, true, gd.Int(0), gd.Int(0))
}

/*
Returns the line the editing caret is on.
*/
func (self Instance) GetCaretLine() int {
	return int(int(class(self).GetCaretLine(gd.Int(0))))
}

/*
Moves the caret to the specified [param column] index.
If [param adjust_viewport] is [code]true[/code], the viewport will center at the caret position after the move occurs.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
func (self Instance) SetCaretColumn(column int) {
	class(self).SetCaretColumn(gd.Int(column), true, gd.Int(0))
}

/*
Returns the column the editing caret is at.
*/
func (self Instance) GetCaretColumn() int {
	return int(int(class(self).GetCaretColumn(gd.Int(0))))
}

/*
Returns the wrap index the editing caret is on.
*/
func (self Instance) GetCaretWrapIndex() int {
	return int(int(class(self).GetCaretWrapIndex(gd.Int(0))))
}

/*
Returns a [String] text with the word under the caret's location.
*/
func (self Instance) GetWordUnderCaret() string {
	return string(class(self).GetWordUnderCaret(gd.Int(-1)).String())
}

/*
Sets the current selection mode.
*/
func (self Instance) SetSelectionMode(mode classdb.TextEditSelectionMode) {
	class(self).SetSelectionMode(mode)
}

/*
Returns the current selection mode.
*/
func (self Instance) GetSelectionMode() classdb.TextEditSelectionMode {
	return classdb.TextEditSelectionMode(class(self).GetSelectionMode())
}

/*
Select all the text.
If [member selecting_enabled] is [code]false[/code], no selection will occur.
*/
func (self Instance) SelectAll() {
	class(self).SelectAll()
}

/*
Selects the word under the caret.
*/
func (self Instance) SelectWordUnderCaret() {
	class(self).SelectWordUnderCaret(gd.Int(-1))
}

/*
Adds a selection and a caret for the next occurrence of the current selection. If there is no active selection, selects word under caret.
*/
func (self Instance) AddSelectionForNextOccurrence() {
	class(self).AddSelectionForNextOccurrence()
}

/*
Moves a selection and a caret for the next occurrence of the current selection. If there is no active selection, moves to the next occurrence of the word under caret.
*/
func (self Instance) SkipSelectionForNextOccurrence() {
	class(self).SkipSelectionForNextOccurrence()
}

/*
Selects text from [param origin_line] and [param origin_column] to [param caret_line] and [param caret_column] for the given [param caret_index]. This moves the selection origin and the caret. If the positions are the same, the selection will be deselected.
If [member selecting_enabled] is [code]false[/code], no selection will occur.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
func (self Instance) Select(origin_line int, origin_column int, caret_line int, caret_column int) {
	class(self).Select(gd.Int(origin_line), gd.Int(origin_column), gd.Int(caret_line), gd.Int(caret_column), gd.Int(0))
}

/*
Returns [code]true[/code] if the user has selected text.
*/
func (self Instance) HasSelection() bool {
	return bool(class(self).HasSelection(gd.Int(-1)))
}

/*
Returns the text inside the selection of a caret, or all the carets if [param caret_index] is its default value [code]-1[/code].
*/
func (self Instance) GetSelectedText() string {
	return string(class(self).GetSelectedText(gd.Int(-1)).String())
}

/*
Returns the caret index of the selection at the given [param line] and [param column], or [code]-1[/code] if there is none.
If [param include_edges] is [code]false[/code], the position must be inside the selection and not at either end. If [param only_selections] is [code]false[/code], carets without a selection will also be considered.
*/
func (self Instance) GetSelectionAtLineColumn(line int, column int) int {
	return int(int(class(self).GetSelectionAtLineColumn(gd.Int(line), gd.Int(column), true, true)))
}

/*
Returns an [Array] of line ranges where [code]x[/code] is the first line and [code]y[/code] is the last line. All lines within these ranges will have a caret on them or be part of a selection. Each line will only be part of one line range, even if it has multiple carets on it.
If a selection's end column ([method get_selection_to_column]) is at column [code]0[/code], that line will not be included. If a selection begins on the line after another selection ends and [param merge_adjacent] is [code]true[/code], or they begin and end on the same line, one line range will include both selections.
*/
func (self Instance) GetLineRangesFromCarets() gd.Array {
	return gd.Array(class(self).GetLineRangesFromCarets(false, true))
}

/*
Returns the origin line of the selection. This is the opposite end from the caret.
*/
func (self Instance) GetSelectionOriginLine() int {
	return int(int(class(self).GetSelectionOriginLine(gd.Int(0))))
}

/*
Returns the origin column of the selection. This is the opposite end from the caret.
*/
func (self Instance) GetSelectionOriginColumn() int {
	return int(int(class(self).GetSelectionOriginColumn(gd.Int(0))))
}

/*
Sets the selection origin line to the [param line] for the given [param caret_index]. If the selection origin is moved to the caret position, the selection will deselect.
If [param can_be_hidden] is [code]false[/code], The line will be set to the nearest unhidden line below or above.
If [param wrap_index] is [code]-1[/code], the selection origin column will be clamped to the [param line]'s length. If [param wrap_index] is greater than [code]-1[/code], the column will be moved to attempt to match the visual x position on the line's [param wrap_index] to the position from the last time [method set_selection_origin_column] or [method select] was called.
*/
func (self Instance) SetSelectionOriginLine(line int) {
	class(self).SetSelectionOriginLine(gd.Int(line), true, gd.Int(-1), gd.Int(0))
}

/*
Sets the selection origin column to the [param column] for the given [param caret_index]. If the selection origin is moved to the caret position, the selection will deselect.
*/
func (self Instance) SetSelectionOriginColumn(column int) {
	class(self).SetSelectionOriginColumn(gd.Int(column), gd.Int(0))
}

/*
Returns the selection begin line. Returns the caret line if there is no selection.
*/
func (self Instance) GetSelectionFromLine() int {
	return int(int(class(self).GetSelectionFromLine(gd.Int(0))))
}

/*
Returns the selection begin column. Returns the caret column if there is no selection.
*/
func (self Instance) GetSelectionFromColumn() int {
	return int(int(class(self).GetSelectionFromColumn(gd.Int(0))))
}

/*
Returns the selection end line. Returns the caret line if there is no selection.
*/
func (self Instance) GetSelectionToLine() int {
	return int(int(class(self).GetSelectionToLine(gd.Int(0))))
}

/*
Returns the selection end column. Returns the caret column if there is no selection.
*/
func (self Instance) GetSelectionToColumn() int {
	return int(int(class(self).GetSelectionToColumn(gd.Int(0))))
}

/*
Returns [code]true[/code] if the caret of the selection is after the selection origin. This can be used to determine the direction of the selection.
*/
func (self Instance) IsCaretAfterSelectionOrigin() bool {
	return bool(class(self).IsCaretAfterSelectionOrigin(gd.Int(0)))
}

/*
Deselects the current selection.
*/
func (self Instance) Deselect() {
	class(self).Deselect(gd.Int(-1))
}

/*
Deletes the selected text.
*/
func (self Instance) DeleteSelection() {
	class(self).DeleteSelection(gd.Int(-1))
}

/*
Returns if the given line is wrapped.
*/
func (self Instance) IsLineWrapped(line int) bool {
	return bool(class(self).IsLineWrapped(gd.Int(line)))
}

/*
Returns the number of times the given line is wrapped.
*/
func (self Instance) GetLineWrapCount(line int) int {
	return int(int(class(self).GetLineWrapCount(gd.Int(line))))
}

/*
Returns the wrap index of the given line column.
*/
func (self Instance) GetLineWrapIndexAtColumn(line int, column int) int {
	return int(int(class(self).GetLineWrapIndexAtColumn(gd.Int(line), gd.Int(column))))
}

/*
Returns an array of [String]s representing each wrapped index.
*/
func (self Instance) GetLineWrappedText(line int) []string {
	return []string(class(self).GetLineWrappedText(gd.Int(line)).Strings())
}

/*
Returns the [VScrollBar] of the [TextEdit].
*/
func (self Instance) GetVScrollBar() objects.VScrollBar {
	return objects.VScrollBar(class(self).GetVScrollBar())
}

/*
Returns the [HScrollBar] used by [TextEdit].
*/
func (self Instance) GetHScrollBar() objects.HScrollBar {
	return objects.HScrollBar(class(self).GetHScrollBar())
}

/*
Returns the scroll position for [param wrap_index] of [param line].
*/
func (self Instance) GetScrollPosForLine(line int) Float.X {
	return Float.X(Float.X(class(self).GetScrollPosForLine(gd.Int(line), gd.Int(0))))
}

/*
Positions the [param wrap_index] of [param line] at the top of the viewport.
*/
func (self Instance) SetLineAsFirstVisible(line int) {
	class(self).SetLineAsFirstVisible(gd.Int(line), gd.Int(0))
}

/*
Returns the first visible line.
*/
func (self Instance) GetFirstVisibleLine() int {
	return int(int(class(self).GetFirstVisibleLine()))
}

/*
Positions the [param wrap_index] of [param line] at the center of the viewport.
*/
func (self Instance) SetLineAsCenterVisible(line int) {
	class(self).SetLineAsCenterVisible(gd.Int(line), gd.Int(0))
}

/*
Positions the [param wrap_index] of [param line] at the bottom of the viewport.
*/
func (self Instance) SetLineAsLastVisible(line int) {
	class(self).SetLineAsLastVisible(gd.Int(line), gd.Int(0))
}

/*
Returns the last visible line. Use [method get_last_full_visible_line_wrap_index] for the wrap index.
*/
func (self Instance) GetLastFullVisibleLine() int {
	return int(int(class(self).GetLastFullVisibleLine()))
}

/*
Returns the last visible wrap index of the last visible line.
*/
func (self Instance) GetLastFullVisibleLineWrapIndex() int {
	return int(int(class(self).GetLastFullVisibleLineWrapIndex()))
}

/*
Returns the number of visible lines, including wrapped text.
*/
func (self Instance) GetVisibleLineCount() int {
	return int(int(class(self).GetVisibleLineCount()))
}

/*
Returns the total number of visible + wrapped lines between the two lines.
*/
func (self Instance) GetVisibleLineCountInRange(from_line int, to_line int) int {
	return int(int(class(self).GetVisibleLineCountInRange(gd.Int(from_line), gd.Int(to_line))))
}

/*
Returns the number of lines that may be drawn.
*/
func (self Instance) GetTotalVisibleLineCount() int {
	return int(int(class(self).GetTotalVisibleLineCount()))
}

/*
Adjust the viewport so the caret is visible.
*/
func (self Instance) AdjustViewportToCaret() {
	class(self).AdjustViewportToCaret(gd.Int(0))
}

/*
Centers the viewport on the line the editing caret is at. This also resets the [member scroll_horizontal] value to [code]0[/code].
*/
func (self Instance) CenterViewportToCaret() {
	class(self).CenterViewportToCaret(gd.Int(0))
}

/*
Returns the number of lines that may be drawn on the minimap.
*/
func (self Instance) GetMinimapVisibleLines() int {
	return int(int(class(self).GetMinimapVisibleLines()))
}

/*
Register a new gutter to this [TextEdit]. Use [param at] to have a specific gutter order. A value of [code]-1[/code] appends the gutter to the right.
*/
func (self Instance) AddGutter() {
	class(self).AddGutter(gd.Int(-1))
}

/*
Removes the gutter from this [TextEdit].
*/
func (self Instance) RemoveGutter(gutter int) {
	class(self).RemoveGutter(gd.Int(gutter))
}

/*
Returns the number of gutters registered.
*/
func (self Instance) GetGutterCount() int {
	return int(int(class(self).GetGutterCount()))
}

/*
Sets the name of the gutter.
*/
func (self Instance) SetGutterName(gutter int, name string) {
	class(self).SetGutterName(gd.Int(gutter), gd.NewString(name))
}

/*
Returns the name of the gutter at the given index.
*/
func (self Instance) GetGutterName(gutter int) string {
	return string(class(self).GetGutterName(gd.Int(gutter)).String())
}

/*
Sets the type of gutter. Gutters can contain icons, text, or custom visuals. See [enum TextEdit.GutterType] for options.
*/
func (self Instance) SetGutterType(gutter int, atype classdb.TextEditGutterType) {
	class(self).SetGutterType(gd.Int(gutter), atype)
}

/*
Returns the type of the gutter at the given index. Gutters can contain icons, text, or custom visuals. See [enum TextEdit.GutterType] for options.
*/
func (self Instance) GetGutterType(gutter int) classdb.TextEditGutterType {
	return classdb.TextEditGutterType(class(self).GetGutterType(gd.Int(gutter)))
}

/*
Set the width of the gutter.
*/
func (self Instance) SetGutterWidth(gutter int, width int) {
	class(self).SetGutterWidth(gd.Int(gutter), gd.Int(width))
}

/*
Returns the width of the gutter at the given index.
*/
func (self Instance) GetGutterWidth(gutter int) int {
	return int(int(class(self).GetGutterWidth(gd.Int(gutter))))
}

/*
Sets whether the gutter should be drawn.
*/
func (self Instance) SetGutterDraw(gutter int, draw bool) {
	class(self).SetGutterDraw(gd.Int(gutter), draw)
}

/*
Returns whether the gutter is currently drawn.
*/
func (self Instance) IsGutterDrawn(gutter int) bool {
	return bool(class(self).IsGutterDrawn(gd.Int(gutter)))
}

/*
Sets the gutter as clickable. This will change the mouse cursor to a pointing hand when hovering over the gutter.
*/
func (self Instance) SetGutterClickable(gutter int, clickable bool) {
	class(self).SetGutterClickable(gd.Int(gutter), clickable)
}

/*
Returns whether the gutter is clickable.
*/
func (self Instance) IsGutterClickable(gutter int) bool {
	return bool(class(self).IsGutterClickable(gd.Int(gutter)))
}

/*
Sets the gutter to overwritable. See [method merge_gutters].
*/
func (self Instance) SetGutterOverwritable(gutter int, overwritable bool) {
	class(self).SetGutterOverwritable(gd.Int(gutter), overwritable)
}

/*
Returns whether the gutter is overwritable.
*/
func (self Instance) IsGutterOverwritable(gutter int) bool {
	return bool(class(self).IsGutterOverwritable(gd.Int(gutter)))
}

/*
Merge the gutters from [param from_line] into [param to_line]. Only overwritable gutters will be copied.
*/
func (self Instance) MergeGutters(from_line int, to_line int) {
	class(self).MergeGutters(gd.Int(from_line), gd.Int(to_line))
}

/*
Set a custom draw method for the gutter. The callback method must take the following args: [code]line: int, gutter: int, Area: Rect2[/code]. This only works when the gutter type is [constant GUTTER_TYPE_CUSTOM] (see [method set_gutter_type]).
*/
func (self Instance) SetGutterCustomDraw(column int, draw_callback func(line int, gutter int, area Rect2.PositionSize)) {
	class(self).SetGutterCustomDraw(gd.Int(column), gd.NewCallable(draw_callback))
}

/*
Returns the total width of all gutters and internal padding.
*/
func (self Instance) GetTotalGutterWidth() int {
	return int(int(class(self).GetTotalGutterWidth()))
}

/*
Sets the metadata for [param gutter] on [param line] to [param metadata].
*/
func (self Instance) SetLineGutterMetadata(line int, gutter int, metadata any) {
	class(self).SetLineGutterMetadata(gd.Int(line), gd.Int(gutter), gd.NewVariant(metadata))
}

/*
Returns the metadata currently in [param gutter] at [param line].
*/
func (self Instance) GetLineGutterMetadata(line int, gutter int) any {
	return any(class(self).GetLineGutterMetadata(gd.Int(line), gd.Int(gutter)).Interface())
}

/*
Sets the text for [param gutter] on [param line] to [param text]. This only works when the gutter type is [constant GUTTER_TYPE_STRING] (see [method set_gutter_type]).
*/
func (self Instance) SetLineGutterText(line int, gutter int, text string) {
	class(self).SetLineGutterText(gd.Int(line), gd.Int(gutter), gd.NewString(text))
}

/*
Returns the text currently in [param gutter] at [param line]. This only works when the gutter type is [constant GUTTER_TYPE_STRING] (see [method set_gutter_type]).
*/
func (self Instance) GetLineGutterText(line int, gutter int) string {
	return string(class(self).GetLineGutterText(gd.Int(line), gd.Int(gutter)).String())
}

/*
Sets the icon for [param gutter] on [param line] to [param icon]. This only works when the gutter type is [constant GUTTER_TYPE_ICON] (see [method set_gutter_type]).
*/
func (self Instance) SetLineGutterIcon(line int, gutter int, icon objects.Texture2D) {
	class(self).SetLineGutterIcon(gd.Int(line), gd.Int(gutter), icon)
}

/*
Returns the icon currently in [param gutter] at [param line]. This only works when the gutter type is [constant GUTTER_TYPE_ICON] (see [method set_gutter_type]).
*/
func (self Instance) GetLineGutterIcon(line int, gutter int) objects.Texture2D {
	return objects.Texture2D(class(self).GetLineGutterIcon(gd.Int(line), gd.Int(gutter)))
}

/*
Sets the color for [param gutter] on [param line] to [param color].
*/
func (self Instance) SetLineGutterItemColor(line int, gutter int, color Color.RGBA) {
	class(self).SetLineGutterItemColor(gd.Int(line), gd.Int(gutter), gd.Color(color))
}

/*
Returns the color currently in [param gutter] at [param line].
*/
func (self Instance) GetLineGutterItemColor(line int, gutter int) Color.RGBA {
	return Color.RGBA(class(self).GetLineGutterItemColor(gd.Int(line), gd.Int(gutter)))
}

/*
If [param clickable] is [code]true[/code], makes the [param gutter] on [param line] clickable. See [signal gutter_clicked].
*/
func (self Instance) SetLineGutterClickable(line int, gutter int, clickable bool) {
	class(self).SetLineGutterClickable(gd.Int(line), gd.Int(gutter), clickable)
}

/*
Returns whether the gutter on the given line is clickable.
*/
func (self Instance) IsLineGutterClickable(line int, gutter int) bool {
	return bool(class(self).IsLineGutterClickable(gd.Int(line), gd.Int(gutter)))
}

/*
Sets the current background color of the line. Set to [code]Color(0, 0, 0, 0)[/code] for no color.
*/
func (self Instance) SetLineBackgroundColor(line int, color Color.RGBA) {
	class(self).SetLineBackgroundColor(gd.Int(line), gd.Color(color))
}

/*
Returns the current background color of the line. [code]Color(0, 0, 0, 0)[/code] is returned if no color is set.
*/
func (self Instance) GetLineBackgroundColor(line int) Color.RGBA {
	return Color.RGBA(class(self).GetLineBackgroundColor(gd.Int(line)))
}

/*
Returns the [PopupMenu] of this [TextEdit]. By default, this menu is displayed when right-clicking on the [TextEdit].
You can add custom menu items or remove standard ones. Make sure your IDs don't conflict with the standard ones (see [enum MenuItems]). For example:
[codeblocks]
[gdscript]
func _ready():

	var menu = get_menu()
	# Remove all items after "Redo".
	menu.item_count = menu.get_item_index(MENU_REDO) + 1
	# Add custom items.
	menu.add_separator()
	menu.add_item("Insert Date", MENU_MAX + 1)
	# Connect callback.
	menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):

	if id == MENU_MAX + 1:
	    insert_text_at_caret(Time.get_date_string_from_system())

[/gdscript]
[csharp]
public override void _Ready()

	{
	    var menu = GetMenu();
	    // Remove all items after "Redo".
	    menu.ItemCount = menu.GetItemIndex(TextEdit.MenuItems.Redo) + 1;
	    // Add custom items.
	    menu.AddSeparator();
	    menu.AddItem("Insert Date", TextEdit.MenuItems.Max + 1);
	    // Add event handler.
	    menu.IdPressed += OnItemPressed;
	}

public void OnItemPressed(int id)

	{
	    if (id == TextEdit.MenuItems.Max + 1)
	    {
	        InsertTextAtCaret(Time.GetDateStringFromSystem());
	    }
	}

[/csharp]
[/codeblocks]
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Instance) GetMenu() objects.PopupMenu {
	return objects.PopupMenu(class(self).GetMenu())
}

/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
func (self Instance) IsMenuVisible() bool {
	return bool(class(self).IsMenuVisible())
}

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
func (self Instance) MenuOption(option int) {
	class(self).MenuOption(gd.Int(option))
}

/*
This method does nothing.
*/
func (self Instance) AdjustCaretsAfterEdit(caret int, from_line int, from_col int, to_line int, to_col int) {
	class(self).AdjustCaretsAfterEdit(gd.Int(caret), gd.Int(from_line), gd.Int(from_col), gd.Int(to_line), gd.Int(to_col))
}

/*
Returns a list of caret indexes in their edit order, this done from bottom to top. Edit order refers to the way actions such as [method insert_text_at_caret] are applied.
*/
func (self Instance) GetCaretIndexEditOrder() []int32 {
	return []int32(class(self).GetCaretIndexEditOrder().AsSlice())
}

/*
Returns the original start line of the selection.
*/
func (self Instance) GetSelectionLine() int {
	return int(int(class(self).GetSelectionLine(gd.Int(0))))
}

/*
Returns the original start column of the selection.
*/
func (self Instance) GetSelectionColumn() int {
	return int(int(class(self).GetSelectionColumn(gd.Int(0))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TextEdit

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextEdit"))
	return Instance{classdb.TextEdit(object)}
}

func (self Instance) Text() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Instance) PlaceholderText() string {
	return string(class(self).GetPlaceholder().String())
}

func (self Instance) SetPlaceholderText(value string) {
	class(self).SetPlaceholder(gd.NewString(value))
}

func (self Instance) Editable() bool {
	return bool(class(self).IsEditable())
}

func (self Instance) SetEditable(value bool) {
	class(self).SetEditable(value)
}

func (self Instance) ContextMenuEnabled() bool {
	return bool(class(self).IsContextMenuEnabled())
}

func (self Instance) SetContextMenuEnabled(value bool) {
	class(self).SetContextMenuEnabled(value)
}

func (self Instance) ShortcutKeysEnabled() bool {
	return bool(class(self).IsShortcutKeysEnabled())
}

func (self Instance) SetShortcutKeysEnabled(value bool) {
	class(self).SetShortcutKeysEnabled(value)
}

func (self Instance) SelectingEnabled() bool {
	return bool(class(self).IsSelectingEnabled())
}

func (self Instance) SetSelectingEnabled(value bool) {
	class(self).SetSelectingEnabled(value)
}

func (self Instance) DeselectOnFocusLossEnabled() bool {
	return bool(class(self).IsDeselectOnFocusLossEnabled())
}

func (self Instance) SetDeselectOnFocusLossEnabled(value bool) {
	class(self).SetDeselectOnFocusLossEnabled(value)
}

func (self Instance) DragAndDropSelectionEnabled() bool {
	return bool(class(self).IsDragAndDropSelectionEnabled())
}

func (self Instance) SetDragAndDropSelectionEnabled(value bool) {
	class(self).SetDragAndDropSelectionEnabled(value)
}

func (self Instance) VirtualKeyboardEnabled() bool {
	return bool(class(self).IsVirtualKeyboardEnabled())
}

func (self Instance) SetVirtualKeyboardEnabled(value bool) {
	class(self).SetVirtualKeyboardEnabled(value)
}

func (self Instance) MiddleMousePasteEnabled() bool {
	return bool(class(self).IsMiddleMousePasteEnabled())
}

func (self Instance) SetMiddleMousePasteEnabled(value bool) {
	class(self).SetMiddleMousePasteEnabled(value)
}

func (self Instance) WrapMode() classdb.TextEditLineWrappingMode {
	return classdb.TextEditLineWrappingMode(class(self).GetLineWrappingMode())
}

func (self Instance) SetWrapMode(value classdb.TextEditLineWrappingMode) {
	class(self).SetLineWrappingMode(value)
}

func (self Instance) AutowrapMode() classdb.TextServerAutowrapMode {
	return classdb.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Instance) SetAutowrapMode(value classdb.TextServerAutowrapMode) {
	class(self).SetAutowrapMode(value)
}

func (self Instance) IndentWrappedLines() bool {
	return bool(class(self).IsIndentWrappedLines())
}

func (self Instance) SetIndentWrappedLines(value bool) {
	class(self).SetIndentWrappedLines(value)
}

func (self Instance) ScrollSmooth() bool {
	return bool(class(self).IsSmoothScrollEnabled())
}

func (self Instance) SetScrollSmooth(value bool) {
	class(self).SetSmoothScrollEnabled(value)
}

func (self Instance) ScrollVScrollSpeed() Float.X {
	return Float.X(Float.X(class(self).GetVScrollSpeed()))
}

func (self Instance) SetScrollVScrollSpeed(value Float.X) {
	class(self).SetVScrollSpeed(gd.Float(value))
}

func (self Instance) ScrollPastEndOfFile() bool {
	return bool(class(self).IsScrollPastEndOfFileEnabled())
}

func (self Instance) SetScrollPastEndOfFile(value bool) {
	class(self).SetScrollPastEndOfFileEnabled(value)
}

func (self Instance) ScrollVertical() Float.X {
	return Float.X(Float.X(class(self).GetVScroll()))
}

func (self Instance) SetScrollVertical(value Float.X) {
	class(self).SetVScroll(gd.Float(value))
}

func (self Instance) ScrollHorizontal() int {
	return int(int(class(self).GetHScroll()))
}

func (self Instance) SetScrollHorizontal(value int) {
	class(self).SetHScroll(gd.Int(value))
}

func (self Instance) ScrollFitContentHeight() bool {
	return bool(class(self).IsFitContentHeightEnabled())
}

func (self Instance) SetScrollFitContentHeight(value bool) {
	class(self).SetFitContentHeightEnabled(value)
}

func (self Instance) MinimapDraw() bool {
	return bool(class(self).IsDrawingMinimap())
}

func (self Instance) SetMinimapDraw(value bool) {
	class(self).SetDrawMinimap(value)
}

func (self Instance) MinimapWidth() int {
	return int(int(class(self).GetMinimapWidth()))
}

func (self Instance) SetMinimapWidth(value int) {
	class(self).SetMinimapWidth(gd.Int(value))
}

func (self Instance) CaretType() classdb.TextEditCaretType {
	return classdb.TextEditCaretType(class(self).GetCaretType())
}

func (self Instance) SetCaretType(value classdb.TextEditCaretType) {
	class(self).SetCaretType(value)
}

func (self Instance) CaretBlink() bool {
	return bool(class(self).IsCaretBlinkEnabled())
}

func (self Instance) SetCaretBlink(value bool) {
	class(self).SetCaretBlinkEnabled(value)
}

func (self Instance) CaretBlinkInterval() Float.X {
	return Float.X(Float.X(class(self).GetCaretBlinkInterval()))
}

func (self Instance) SetCaretBlinkInterval(value Float.X) {
	class(self).SetCaretBlinkInterval(gd.Float(value))
}

func (self Instance) CaretDrawWhenEditableDisabled() bool {
	return bool(class(self).IsDrawingCaretWhenEditableDisabled())
}

func (self Instance) SetCaretDrawWhenEditableDisabled(value bool) {
	class(self).SetDrawCaretWhenEditableDisabled(value)
}

func (self Instance) CaretMoveOnRightClick() bool {
	return bool(class(self).IsMoveCaretOnRightClickEnabled())
}

func (self Instance) SetCaretMoveOnRightClick(value bool) {
	class(self).SetMoveCaretOnRightClickEnabled(value)
}

func (self Instance) CaretMidGrapheme() bool {
	return bool(class(self).IsCaretMidGraphemeEnabled())
}

func (self Instance) SetCaretMidGrapheme(value bool) {
	class(self).SetCaretMidGraphemeEnabled(value)
}

func (self Instance) CaretMultiple() bool {
	return bool(class(self).IsMultipleCaretsEnabled())
}

func (self Instance) SetCaretMultiple(value bool) {
	class(self).SetMultipleCaretsEnabled(value)
}

func (self Instance) UseDefaultWordSeparators() bool {
	return bool(class(self).IsDefaultWordSeparatorsEnabled())
}

func (self Instance) SetUseDefaultWordSeparators(value bool) {
	class(self).SetUseDefaultWordSeparators(value)
}

func (self Instance) UseCustomWordSeparators() bool {
	return bool(class(self).IsCustomWordSeparatorsEnabled())
}

func (self Instance) SetUseCustomWordSeparators(value bool) {
	class(self).SetUseCustomWordSeparators(value)
}

func (self Instance) CustomWordSeparators() string {
	return string(class(self).GetCustomWordSeparators().String())
}

func (self Instance) SetCustomWordSeparators(value string) {
	class(self).SetCustomWordSeparators(gd.NewString(value))
}

func (self Instance) SyntaxHighlighter() objects.SyntaxHighlighter {
	return objects.SyntaxHighlighter(class(self).GetSyntaxHighlighter())
}

func (self Instance) SetSyntaxHighlighter(value objects.SyntaxHighlighter) {
	class(self).SetSyntaxHighlighter(value)
}

func (self Instance) HighlightAllOccurrences() bool {
	return bool(class(self).IsHighlightAllOccurrencesEnabled())
}

func (self Instance) SetHighlightAllOccurrences(value bool) {
	class(self).SetHighlightAllOccurrences(value)
}

func (self Instance) HighlightCurrentLine() bool {
	return bool(class(self).IsHighlightCurrentLineEnabled())
}

func (self Instance) SetHighlightCurrentLine(value bool) {
	class(self).SetHighlightCurrentLine(value)
}

func (self Instance) DrawControlChars() bool {
	return bool(class(self).GetDrawControlChars())
}

func (self Instance) SetDrawControlChars(value bool) {
	class(self).SetDrawControlChars(value)
}

func (self Instance) DrawTabs() bool {
	return bool(class(self).IsDrawingTabs())
}

func (self Instance) SetDrawTabs(value bool) {
	class(self).SetDrawTabs(value)
}

func (self Instance) DrawSpaces() bool {
	return bool(class(self).IsDrawingSpaces())
}

func (self Instance) SetDrawSpaces(value bool) {
	class(self).SetDrawSpaces(value)
}

func (self Instance) TextDirection() classdb.ControlTextDirection {
	return classdb.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value classdb.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(gd.NewString(value))
}

func (self Instance) StructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	return classdb.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Instance) SetStructuredTextBidiOverride(value classdb.TextServerStructuredTextParser) {
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Instance) StructuredTextBidiOverrideOptions() Array.Any {
	return Array.Any(class(self).GetStructuredTextBidiOverrideOptions())
}

func (self Instance) SetStructuredTextBidiOverrideOptions(value Array.Any) {
	class(self).SetStructuredTextBidiOverrideOptions(value)
}

/*
Override this method to define what happens when the user types in the provided key [param unicode_char].
*/
func (class) _handle_unicode_input(impl func(ptr unsafe.Pointer, unicode_char gd.Int, caret_index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var unicode_char = gd.UnsafeGet[gd.Int](p_args, 0)
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, unicode_char, caret_index)
	}
}

/*
Override this method to define what happens when the user presses the backspace key.
*/
func (class) _backspace(impl func(ptr unsafe.Pointer, caret_index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, caret_index)
	}
}

/*
Override this method to define what happens when the user performs a cut operation.
*/
func (class) _cut(impl func(ptr unsafe.Pointer, caret_index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, caret_index)
	}
}

/*
Override this method to define what happens when the user performs a copy operation.
*/
func (class) _copy(impl func(ptr unsafe.Pointer, caret_index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, caret_index)
	}
}

/*
Override this method to define what happens when the user performs a paste operation.
*/
func (class) _paste(impl func(ptr unsafe.Pointer, caret_index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, caret_index)
	}
}

/*
Override this method to define what happens when the user performs a paste operation with middle mouse button.
[b]Note:[/b] This method is only implemented on Linux.
*/
func (class) _paste_primary_clipboard(impl func(ptr unsafe.Pointer, caret_index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var caret_index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, caret_index)
	}
}

/*
Returns [code]true[/code] if the user has text in the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME).
*/
//go:nosplit
func (self class) HasImeText() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_has_ime_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Closes the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME) if it is open. Any text in the IME will be lost.
*/
//go:nosplit
func (self class) CancelIme() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_cancel_ime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies text from the [url=https://en.wikipedia.org/wiki/Input_method]Input Method Editor[/url] (IME) to each caret and closes the IME if it is open.
*/
//go:nosplit
func (self class) ApplyIme() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_apply_ime, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetEditable(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser classdb.TextServerStructuredTextParser) {
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(args))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the tab size for the [TextEdit] to use.
*/
//go:nosplit
func (self class) SetTabSize(size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_tab_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [TextEdit]'s' tab size.
*/
//go:nosplit
func (self class) GetTabSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_tab_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIndentWrappedLines(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_indent_wrapped_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsIndentWrappedLines() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_indent_wrapped_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], sets the user into overtype mode. When the user types in this mode, it will override existing text.
*/
//go:nosplit
func (self class) SetOvertypeModeEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_overtype_mode_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the user is in overtype mode.
*/
//go:nosplit
func (self class) IsOvertypeModeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_overtype_mode_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContextMenuEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsContextMenuEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShortcutKeysEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsShortcutKeysEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVirtualKeyboardEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_virtual_keyboard_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVirtualKeyboardEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_virtual_keyboard_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMiddleMousePasteEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_middle_mouse_paste_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsMiddleMousePasteEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_middle_mouse_paste_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Performs a full reset of [TextEdit], including undo history.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetText(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the number of lines in the text.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlaceholder(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlaceholder() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the text for a specific [param line].
Carets on the line will attempt to keep their visual x position.
*/
//go:nosplit
func (self class) SetLine(line gd.Int, new_text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, pointers.Get(new_text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the text of a specific line.
*/
//go:nosplit
func (self class) GetLine(line gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the width in pixels of the [param wrap_index] on [param line].
*/
//go:nosplit
func (self class) GetLineWidth(line gd.Int, wrap_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the maximum value of the line height among all lines.
[b]Note:[/b] The return value is influenced by [theme_item line_spacing] and [theme_item font_size]. And it will not be less than [code]1[/code].
*/
//go:nosplit
func (self class) GetLineHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of spaces and [code]tab * tab_size[/code] before the first char.
*/
//go:nosplit
func (self class) GetIndentLevel(line gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_indent_level, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the first column containing a non-whitespace character.
*/
//go:nosplit
func (self class) GetFirstNonWhitespaceColumn(line gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_first_non_whitespace_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Swaps the two lines. Carets will be swapped with the lines.
*/
//go:nosplit
func (self class) SwapLines(from_line gd.Int, to_line gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, to_line)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_swap_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Inserts a new line with [param text] at [param line].
*/
//go:nosplit
func (self class) InsertLineAt(line gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_insert_line_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the line of text at [param line]. Carets on this line will attempt to match their previous visual x position.
If [param move_carets_down] is [code]true[/code] carets will move to the next line down, otherwise carets will move up.
*/
//go:nosplit
func (self class) RemoveLineAt(line gd.Int, move_carets_down bool) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, move_carets_down)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_remove_line_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Insert the specified text at the caret position.
*/
//go:nosplit
func (self class) InsertTextAtCaret(text gd.String, caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_insert_text_at_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Inserts the [param text] at [param line] and [param column].
If [param before_selection_begin] is [code]true[/code], carets and selections that begin at [param line] and [param column] will moved to the end of the inserted text, along with all carets after it.
If [param before_selection_end] is [code]true[/code], selections that end at [param line] and [param column] will be extended to the end of the inserted text. These parameters can be used to insert text inside of or outside of selections.
*/
//go:nosplit
func (self class) InsertText(text gd.String, line gd.Int, column gd.Int, before_selection_begin bool, before_selection_end bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	callframe.Arg(frame, before_selection_begin)
	callframe.Arg(frame, before_selection_end)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_insert_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes text between the given positions.
*/
//go:nosplit
func (self class) RemoveText(from_line gd.Int, from_column gd.Int, to_line gd.Int, to_column gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, from_column)
	callframe.Arg(frame, to_line)
	callframe.Arg(frame, to_column)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_remove_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the last unhidden line in the entire [TextEdit].
*/
//go:nosplit
func (self class) GetLastUnhiddenLine() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_last_unhidden_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the count to the next visible line from [param line] to [code]line + visible_amount[/code]. Can also count backwards. For example if a [TextEdit] has 5 lines with lines 2 and 3 hidden, calling this with [code]line = 1, visible_amount = 1[/code] would return 3.
*/
//go:nosplit
func (self class) GetNextVisibleLineOffsetFrom(line gd.Int, visible_amount gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, visible_amount)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_next_visible_line_offset_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Similar to [method get_next_visible_line_offset_from], but takes into account the line wrap indexes. In the returned vector, [code]x[/code] is the line, [code]y[/code] is the wrap index.
*/
//go:nosplit
func (self class) GetNextVisibleLineIndexOffsetFrom(line gd.Int, wrap_index gd.Int, visible_amount gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	callframe.Arg(frame, visible_amount)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_next_visible_line_index_offset_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Called when the user presses the backspace key. Can be overridden with [method _backspace].
*/
//go:nosplit
func (self class) Backspace(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_backspace, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Cut's the current selection. Can be overridden with [method _cut].
*/
//go:nosplit
func (self class) Cut(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_cut, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Copies the current text selection. Can be overridden with [method _copy].
*/
//go:nosplit
func (self class) Copy(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_copy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Paste at the current location. Can be overridden with [method _paste].
*/
//go:nosplit
func (self class) Paste(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_paste, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Pastes the primary clipboard.
*/
//go:nosplit
func (self class) PastePrimaryClipboard(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_paste_primary_clipboard, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Starts an action, will end the current action if [param action] is different.
An action will also end after a call to [method end_action], after [member ProjectSettings.gui/timers/text_edit_idle_detect_sec] is triggered or a new undoable step outside the [method start_action] and [method end_action] calls.
*/
//go:nosplit
func (self class) StartAction(action classdb.TextEditEditAction) {
	var frame = callframe.New()
	callframe.Arg(frame, action)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_start_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Marks the end of steps in the current action started with [method start_action].
*/
//go:nosplit
func (self class) EndAction() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_end_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Starts a multipart edit. All edits will be treated as one action until [method end_complex_operation] is called.
*/
//go:nosplit
func (self class) BeginComplexOperation() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_begin_complex_operation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Ends a multipart edit, started with [method begin_complex_operation]. If called outside a complex operation, the current operation is pushed onto the undo/redo stack.
*/
//go:nosplit
func (self class) EndComplexOperation() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_end_complex_operation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if an "undo" action is available.
*/
//go:nosplit
func (self class) HasUndo() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_has_undo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a "redo" action is available.
*/
//go:nosplit
func (self class) HasRedo() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_has_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Perform undo operation.
*/
//go:nosplit
func (self class) Undo() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_undo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Perform redo operation.
*/
//go:nosplit
func (self class) Redo() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears the undo history.
*/
//go:nosplit
func (self class) ClearUndoHistory() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_clear_undo_history, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Tag the current version as saved.
*/
//go:nosplit
func (self class) TagSavedVersion() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_tag_saved_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current version of the [TextEdit]. The version is a count of recorded operations by the undo/redo history.
*/
//go:nosplit
func (self class) GetVersion() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the last tagged saved version from [method tag_saved_version].
*/
//go:nosplit
func (self class) GetSavedVersion() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_saved_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the search text. See [method set_search_flags].
*/
//go:nosplit
func (self class) SetSearchText(search_text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(search_text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_search_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the search [param flags]. This is used with [method set_search_text] to highlight occurrences of the searched text. Search flags can be specified from the [enum SearchFlags] enum.
*/
//go:nosplit
func (self class) SetSearchFlags(flags gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_search_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Perform a search inside the text. Search flags can be specified in the [enum SearchFlags] enum.
In the returned vector, [code]x[/code] is the column, [code]y[/code] is the line. If no results are found, both are equal to [code]-1[/code].
[codeblocks]
[gdscript]
var result = search("print", SEARCH_WHOLE_WORDS, 0, 0)
if result.x != -1:
    # Result found.
    var line_number = result.y
    var column_number = result.x
[/gdscript]
[csharp]
Vector2I result = Search("print", (uint)TextEdit.SearchFlags.WholeWords, 0, 0);
if (result.X != -1)
{
    // Result found.
    int lineNumber = result.Y;
    int columnNumber = result.X;
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) Search(text gd.String, flags gd.Int, from_line gd.Int, from_column gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, flags)
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, from_column)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_search, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Provide custom tooltip text. The callback method must take the following args: [code]hovered_word: String[/code].
*/
//go:nosplit
func (self class) SetTooltipRequestFunc(callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_tooltip_request_func, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the local mouse position adjusted for the text direction.
*/
//go:nosplit
func (self class) GetLocalMousePos() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_local_mouse_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the word at [param position].
*/
//go:nosplit
func (self class) GetWordAtPos(position gd.Vector2) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_word_at_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the line and column at the given position. In the returned vector, [code]x[/code] is the column, [code]y[/code] is the line. If [param allow_out_of_bounds] is [code]false[/code] and the position is not over the text, both vector values will be set to [code]-1[/code].
*/
//go:nosplit
func (self class) GetLineColumnAtPos(position gd.Vector2i, allow_out_of_bounds bool) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, allow_out_of_bounds)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_column_at_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local position for the given [param line] and [param column]. If [code]x[/code] or [code]y[/code] of the returned vector equal [code]-1[/code], the position is outside of the viewable area of the control.
[b]Note:[/b] The Y position corresponds to the bottom side of the line. Use [method get_rect_at_line_column] to get the top side position.
*/
//go:nosplit
func (self class) GetPosAtLineColumn(line gd.Int, column gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_pos_at_line_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local position and size for the grapheme at the given [param line] and [param column]. If [code]x[/code] or [code]y[/code] position of the returned rect equal [code]-1[/code], the position is outside of the viewable area of the control.
[b]Note:[/b] The Y position of the returned rect corresponds to the top side of the line, unlike [method get_pos_at_line_column] which returns the bottom side.
*/
//go:nosplit
func (self class) GetRectAtLineColumn(line gd.Int, column gd.Int) gd.Rect2i {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_rect_at_line_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the equivalent minimap line at [param position].
*/
//go:nosplit
func (self class) GetMinimapLineAtPos(position gd.Vector2i) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_minimap_line_at_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the user is dragging their mouse for scrolling, selecting, or text dragging.
*/
//go:nosplit
func (self class) IsDraggingCursor() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_dragging_cursor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the mouse is over selection. If [param edges] is [code]true[/code], the edges are considered part of the selection.
*/
//go:nosplit
func (self class) IsMouseOverSelection(edges bool, caret_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, edges)
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_mouse_over_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretType(atype classdb.TextEditCaretType) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_caret_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCaretType() classdb.TextEditCaretType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextEditCaretType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_caret_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretBlinkEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_caret_blink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCaretBlinkEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_caret_blink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretBlinkInterval(interval gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_caret_blink_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCaretBlinkInterval() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_caret_blink_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawCaretWhenEditableDisabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_draw_caret_when_editable_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawingCaretWhenEditableDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_drawing_caret_when_editable_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMoveCaretOnRightClickEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_move_caret_on_right_click_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsMoveCaretOnRightClickEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_move_caret_on_right_click_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretMidGraphemeEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_caret_mid_grapheme_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCaretMidGraphemeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_caret_mid_grapheme_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMultipleCaretsEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_multiple_carets_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsMultipleCaretsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_multiple_carets_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a new caret at the given location. Returns the index of the new caret, or [code]-1[/code] if the location is invalid.
*/
//go:nosplit
func (self class) AddCaret(line gd.Int, column gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_add_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the given caret index.
[b]Note:[/b] This can result in adjustment of all other caret indices.
*/
//go:nosplit
func (self class) RemoveCaret(caret gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_remove_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all additional carets.
*/
//go:nosplit
func (self class) RemoveSecondaryCarets() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_remove_secondary_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of carets in this [TextEdit].
*/
//go:nosplit
func (self class) GetCaretCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_caret_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds an additional caret above or below every caret. If [param below] is [code]true[/code] the new caret will be added below and above otherwise.
*/
//go:nosplit
func (self class) AddCaretAtCarets(below bool) {
	var frame = callframe.New()
	callframe.Arg(frame, below)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_add_caret_at_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the carets sorted by selection beginning from lowest line and column to highest (from top to bottom of text).
If [param include_ignored_carets] is [code]false[/code], carets from [method multicaret_edit_ignore_caret] will be ignored.
*/
//go:nosplit
func (self class) GetSortedCarets(include_ignored_carets bool) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, include_ignored_carets)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_sorted_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Collapse all carets in the given range to the [param from_line] and [param from_column] position.
[param inclusive] applies to both ends.
If [method is_in_mulitcaret_edit] is [code]true[/code], carets that are collapsed will be [code]true[/code] for [method multicaret_edit_ignore_caret].
[method merge_overlapping_carets] will be called if any carets were collapsed.
*/
//go:nosplit
func (self class) CollapseCarets(from_line gd.Int, from_column gd.Int, to_line gd.Int, to_column gd.Int, inclusive bool) {
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, from_column)
	callframe.Arg(frame, to_line)
	callframe.Arg(frame, to_column)
	callframe.Arg(frame, inclusive)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_collapse_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Merges any overlapping carets. Will favor the newest caret, or the caret with a selection.
If [method is_in_mulitcaret_edit] is [code]true[/code], the merge will be queued to happen at the end of the multicaret edit. See [method begin_multicaret_edit] and [method end_multicaret_edit].
[b]Note:[/b] This is not called when a caret changes position but after certain actions, so it is possible to get into a state where carets overlap.
*/
//go:nosplit
func (self class) MergeOverlappingCarets() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_merge_overlapping_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Starts an edit for multiple carets. The edit must be ended with [method end_multicaret_edit]. Multicaret edits can be used to edit text at multiple carets and delay merging the carets until the end, so the caret indexes aren't affected immediately. [method begin_multicaret_edit] and [method end_multicaret_edit] can be nested, and the merge will happen at the last [method end_multicaret_edit].
Example usage:
[codeblock]
begin_complex_operation()
begin_multicaret_edit()
for i in range(get_caret_count()):
    if multicaret_edit_ignore_caret(i):
        continue
    # Logic here.
end_multicaret_edit()
end_complex_operation()
[/codeblock]
*/
//go:nosplit
func (self class) BeginMulticaretEdit() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_begin_multicaret_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Ends an edit for multiple carets, that was started with [method begin_multicaret_edit]. If this was the last [method end_multicaret_edit] and [method merge_overlapping_carets] was called, carets will be merged.
*/
//go:nosplit
func (self class) EndMulticaretEdit() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_end_multicaret_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if a [method begin_multicaret_edit] has been called and [method end_multicaret_edit] has not yet been called.
*/
//go:nosplit
func (self class) IsInMulitcaretEdit() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_in_mulitcaret_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [param caret_index] should be ignored as part of a multicaret edit. See [method begin_multicaret_edit] and [method end_multicaret_edit]. Carets that should be ignored are ones that were part of removed text and will likely be merged at the end of the edit, or carets that were added during the edit.
It is recommended to [code]continue[/code] within a loop iterating on multiple carets if a caret should be ignored.
*/
//go:nosplit
func (self class) MulticaretEditIgnoreCaret(caret_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_multicaret_edit_ignore_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the caret is visible on the screen.
*/
//go:nosplit
func (self class) IsCaretVisible(caret_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_caret_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the caret pixel draw position.
*/
//go:nosplit
func (self class) GetCaretDrawPos(caret_index gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_caret_draw_pos, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves the caret to the specified [param line] index. The caret column will be moved to the same visual position it was at the last time [method set_caret_column] was called, or clamped to the end of the line.
If [param adjust_viewport] is [code]true[/code], the viewport will center at the caret position after the move occurs.
If [param can_be_hidden] is [code]true[/code], the specified [param line] can be hidden.
If [param wrap_index] is [code]-1[/code], the caret column will be clamped to the [param line]'s length. If [param wrap_index] is greater than [code]-1[/code], the column will be moved to attempt to match the visual x position on the line's [param wrap_index] to the position from the last time [method set_caret_column] was called.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
//go:nosplit
func (self class) SetCaretLine(line gd.Int, adjust_viewport bool, can_be_hidden bool, wrap_index gd.Int, caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, adjust_viewport)
	callframe.Arg(frame, can_be_hidden)
	callframe.Arg(frame, wrap_index)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_caret_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the line the editing caret is on.
*/
//go:nosplit
func (self class) GetCaretLine(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_caret_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves the caret to the specified [param column] index.
If [param adjust_viewport] is [code]true[/code], the viewport will center at the caret position after the move occurs.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
//go:nosplit
func (self class) SetCaretColumn(column gd.Int, adjust_viewport bool, caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, adjust_viewport)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_caret_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the column the editing caret is at.
*/
//go:nosplit
func (self class) GetCaretColumn(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_caret_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the wrap index the editing caret is on.
*/
//go:nosplit
func (self class) GetCaretWrapIndex(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_caret_wrap_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [String] text with the word under the caret's location.
*/
//go:nosplit
func (self class) GetWordUnderCaret(caret_index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_word_under_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseDefaultWordSeparators(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_use_default_word_separators, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDefaultWordSeparatorsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_default_word_separators_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseCustomWordSeparators(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_use_custom_word_separators, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCustomWordSeparatorsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_custom_word_separators_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCustomWordSeparators(custom_word_separators gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(custom_word_separators))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_custom_word_separators, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomWordSeparators() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_custom_word_separators, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectingEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_selecting_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelectingEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_selecting_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeselectOnFocusLossEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDeselectOnFocusLossEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragAndDropSelectionEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDragAndDropSelectionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the current selection mode.
*/
//go:nosplit
func (self class) SetSelectionMode(mode classdb.TextEditSelectionMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_selection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current selection mode.
*/
//go:nosplit
func (self class) GetSelectionMode() classdb.TextEditSelectionMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextEditSelectionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Select all the text.
If [member selecting_enabled] is [code]false[/code], no selection will occur.
*/
//go:nosplit
func (self class) SelectAll() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_select_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Selects the word under the caret.
*/
//go:nosplit
func (self class) SelectWordUnderCaret(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_select_word_under_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a selection and a caret for the next occurrence of the current selection. If there is no active selection, selects word under caret.
*/
//go:nosplit
func (self class) AddSelectionForNextOccurrence() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_add_selection_for_next_occurrence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Moves a selection and a caret for the next occurrence of the current selection. If there is no active selection, moves to the next occurrence of the word under caret.
*/
//go:nosplit
func (self class) SkipSelectionForNextOccurrence() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_skip_selection_for_next_occurrence, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Selects text from [param origin_line] and [param origin_column] to [param caret_line] and [param caret_column] for the given [param caret_index]. This moves the selection origin and the caret. If the positions are the same, the selection will be deselected.
If [member selecting_enabled] is [code]false[/code], no selection will occur.
[b]Note:[/b] If supporting multiple carets this will not check for any overlap. See [method merge_overlapping_carets].
*/
//go:nosplit
func (self class) Select(origin_line gd.Int, origin_column gd.Int, caret_line gd.Int, caret_column gd.Int, caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, origin_line)
	callframe.Arg(frame, origin_column)
	callframe.Arg(frame, caret_line)
	callframe.Arg(frame, caret_column)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the user has selected text.
*/
//go:nosplit
func (self class) HasSelection(caret_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_has_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text inside the selection of a caret, or all the carets if [param caret_index] is its default value [code]-1[/code].
*/
//go:nosplit
func (self class) GetSelectedText(caret_index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selected_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the caret index of the selection at the given [param line] and [param column], or [code]-1[/code] if there is none.
If [param include_edges] is [code]false[/code], the position must be inside the selection and not at either end. If [param only_selections] is [code]false[/code], carets without a selection will also be considered.
*/
//go:nosplit
func (self class) GetSelectionAtLineColumn(line gd.Int, column gd.Int, include_edges bool, only_selections bool) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	callframe.Arg(frame, include_edges)
	callframe.Arg(frame, only_selections)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_at_line_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an [Array] of line ranges where [code]x[/code] is the first line and [code]y[/code] is the last line. All lines within these ranges will have a caret on them or be part of a selection. Each line will only be part of one line range, even if it has multiple carets on it.
If a selection's end column ([method get_selection_to_column]) is at column [code]0[/code], that line will not be included. If a selection begins on the line after another selection ends and [param merge_adjacent] is [code]true[/code], or they begin and end on the same line, one line range will include both selections.
*/
//go:nosplit
func (self class) GetLineRangesFromCarets(only_selections bool, merge_adjacent bool) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, only_selections)
	callframe.Arg(frame, merge_adjacent)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_ranges_from_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the origin line of the selection. This is the opposite end from the caret.
*/
//go:nosplit
func (self class) GetSelectionOriginLine(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_origin_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the origin column of the selection. This is the opposite end from the caret.
*/
//go:nosplit
func (self class) GetSelectionOriginColumn(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_origin_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the selection origin line to the [param line] for the given [param caret_index]. If the selection origin is moved to the caret position, the selection will deselect.
If [param can_be_hidden] is [code]false[/code], The line will be set to the nearest unhidden line below or above.
If [param wrap_index] is [code]-1[/code], the selection origin column will be clamped to the [param line]'s length. If [param wrap_index] is greater than [code]-1[/code], the column will be moved to attempt to match the visual x position on the line's [param wrap_index] to the position from the last time [method set_selection_origin_column] or [method select] was called.
*/
//go:nosplit
func (self class) SetSelectionOriginLine(line gd.Int, can_be_hidden bool, wrap_index gd.Int, caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, can_be_hidden)
	callframe.Arg(frame, wrap_index)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_selection_origin_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the selection origin column to the [param column] for the given [param caret_index]. If the selection origin is moved to the caret position, the selection will deselect.
*/
//go:nosplit
func (self class) SetSelectionOriginColumn(column gd.Int, caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_selection_origin_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the selection begin line. Returns the caret line if there is no selection.
*/
//go:nosplit
func (self class) GetSelectionFromLine(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_from_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the selection begin column. Returns the caret column if there is no selection.
*/
//go:nosplit
func (self class) GetSelectionFromColumn(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_from_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the selection end line. Returns the caret line if there is no selection.
*/
//go:nosplit
func (self class) GetSelectionToLine(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_to_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the selection end column. Returns the caret column if there is no selection.
*/
//go:nosplit
func (self class) GetSelectionToColumn(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_to_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the caret of the selection is after the selection origin. This can be used to determine the direction of the selection.
*/
//go:nosplit
func (self class) IsCaretAfterSelectionOrigin(caret_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_caret_after_selection_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Deselects the current selection.
*/
//go:nosplit
func (self class) Deselect(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Deletes the selected text.
*/
//go:nosplit
func (self class) DeleteSelection(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_delete_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetLineWrappingMode(mode classdb.TextEditLineWrappingMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_wrapping_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLineWrappingMode() classdb.TextEditLineWrappingMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextEditLineWrappingMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_wrapping_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode classdb.TextServerAutowrapMode) {
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutowrapMode() classdb.TextServerAutowrapMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns if the given line is wrapped.
*/
//go:nosplit
func (self class) IsLineWrapped(line gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_line_wrapped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of times the given line is wrapped.
*/
//go:nosplit
func (self class) GetLineWrapCount(line gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_wrap_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the wrap index of the given line column.
*/
//go:nosplit
func (self class) GetLineWrapIndexAtColumn(line gd.Int, column gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_wrap_index_at_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of [String]s representing each wrapped index.
*/
//go:nosplit
func (self class) GetLineWrappedText(line gd.Int) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_wrapped_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSmoothScrollEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_smooth_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSmoothScrollEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_smooth_scroll_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [VScrollBar] of the [TextEdit].
*/
//go:nosplit
func (self class) GetVScrollBar() objects.VScrollBar {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.VScrollBar{classdb.VScrollBar(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the [HScrollBar] used by [TextEdit].
*/
//go:nosplit
func (self class) GetHScrollBar() objects.HScrollBar {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_h_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.HScrollBar{classdb.HScrollBar(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVScroll(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_v_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVScroll() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_v_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHScroll(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_h_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHScroll() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_h_scroll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollPastEndOfFileEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_scroll_past_end_of_file_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsScrollPastEndOfFileEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_scroll_past_end_of_file_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVScrollSpeed(speed gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, speed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_v_scroll_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVScrollSpeed() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_v_scroll_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFitContentHeightEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_fit_content_height_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFitContentHeightEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_fit_content_height_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scroll position for [param wrap_index] of [param line].
*/
//go:nosplit
func (self class) GetScrollPosForLine(line gd.Int, wrap_index gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_scroll_pos_for_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Positions the [param wrap_index] of [param line] at the top of the viewport.
*/
//go:nosplit
func (self class) SetLineAsFirstVisible(line gd.Int, wrap_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_as_first_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the first visible line.
*/
//go:nosplit
func (self class) GetFirstVisibleLine() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_first_visible_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Positions the [param wrap_index] of [param line] at the center of the viewport.
*/
//go:nosplit
func (self class) SetLineAsCenterVisible(line gd.Int, wrap_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_as_center_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Positions the [param wrap_index] of [param line] at the bottom of the viewport.
*/
//go:nosplit
func (self class) SetLineAsLastVisible(line gd.Int, wrap_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, wrap_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_as_last_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the last visible line. Use [method get_last_full_visible_line_wrap_index] for the wrap index.
*/
//go:nosplit
func (self class) GetLastFullVisibleLine() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_last_full_visible_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the last visible wrap index of the last visible line.
*/
//go:nosplit
func (self class) GetLastFullVisibleLineWrapIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_last_full_visible_line_wrap_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of visible lines, including wrapped text.
*/
//go:nosplit
func (self class) GetVisibleLineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_visible_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total number of visible + wrapped lines between the two lines.
*/
//go:nosplit
func (self class) GetVisibleLineCountInRange(from_line gd.Int, to_line gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, to_line)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_visible_line_count_in_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of lines that may be drawn.
*/
//go:nosplit
func (self class) GetTotalVisibleLineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_total_visible_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adjust the viewport so the caret is visible.
*/
//go:nosplit
func (self class) AdjustViewportToCaret(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_adjust_viewport_to_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Centers the viewport on the line the editing caret is at. This also resets the [member scroll_horizontal] value to [code]0[/code].
*/
//go:nosplit
func (self class) CenterViewportToCaret(caret_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_center_viewport_to_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetDrawMinimap(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_draw_minimap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawingMinimap() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_drawing_minimap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinimapWidth(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_minimap_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinimapWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_minimap_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of lines that may be drawn on the minimap.
*/
//go:nosplit
func (self class) GetMinimapVisibleLines() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_minimap_visible_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Register a new gutter to this [TextEdit]. Use [param at] to have a specific gutter order. A value of [code]-1[/code] appends the gutter to the right.
*/
//go:nosplit
func (self class) AddGutter(at gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, at)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_add_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the gutter from this [TextEdit].
*/
//go:nosplit
func (self class) RemoveGutter(gutter gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_remove_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of gutters registered.
*/
//go:nosplit
func (self class) GetGutterCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_gutter_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the name of the gutter.
*/
//go:nosplit
func (self class) SetGutterName(gutter gd.Int, name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_gutter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the name of the gutter at the given index.
*/
//go:nosplit
func (self class) GetGutterName(gutter gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_gutter_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the type of gutter. Gutters can contain icons, text, or custom visuals. See [enum TextEdit.GutterType] for options.
*/
//go:nosplit
func (self class) SetGutterType(gutter gd.Int, atype classdb.TextEditGutterType) {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_gutter_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the type of the gutter at the given index. Gutters can contain icons, text, or custom visuals. See [enum TextEdit.GutterType] for options.
*/
//go:nosplit
func (self class) GetGutterType(gutter gd.Int) classdb.TextEditGutterType {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[classdb.TextEditGutterType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_gutter_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set the width of the gutter.
*/
//go:nosplit
func (self class) SetGutterWidth(gutter gd.Int, width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_gutter_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the width of the gutter at the given index.
*/
//go:nosplit
func (self class) GetGutterWidth(gutter gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_gutter_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the gutter should be drawn.
*/
//go:nosplit
func (self class) SetGutterDraw(gutter gd.Int, draw bool) {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, draw)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_gutter_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the gutter is currently drawn.
*/
//go:nosplit
func (self class) IsGutterDrawn(gutter gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_gutter_drawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the gutter as clickable. This will change the mouse cursor to a pointing hand when hovering over the gutter.
*/
//go:nosplit
func (self class) SetGutterClickable(gutter gd.Int, clickable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, clickable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_gutter_clickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the gutter is clickable.
*/
//go:nosplit
func (self class) IsGutterClickable(gutter gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_gutter_clickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the gutter to overwritable. See [method merge_gutters].
*/
//go:nosplit
func (self class) SetGutterOverwritable(gutter gd.Int, overwritable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, overwritable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_gutter_overwritable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the gutter is overwritable.
*/
//go:nosplit
func (self class) IsGutterOverwritable(gutter gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_gutter_overwritable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Merge the gutters from [param from_line] into [param to_line]. Only overwritable gutters will be copied.
*/
//go:nosplit
func (self class) MergeGutters(from_line gd.Int, to_line gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, to_line)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_merge_gutters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Set a custom draw method for the gutter. The callback method must take the following args: [code]line: int, gutter: int, Area: Rect2[/code]. This only works when the gutter type is [constant GUTTER_TYPE_CUSTOM] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) SetGutterCustomDraw(column gd.Int, draw_callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(draw_callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_gutter_custom_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the total width of all gutters and internal padding.
*/
//go:nosplit
func (self class) GetTotalGutterWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_total_gutter_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the metadata for [param gutter] on [param line] to [param metadata].
*/
//go:nosplit
func (self class) SetLineGutterMetadata(line gd.Int, gutter gd.Int, metadata gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, pointers.Get(metadata))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_gutter_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the metadata currently in [param gutter] at [param line].
*/
//go:nosplit
func (self class) GetLineGutterMetadata(line gd.Int, gutter gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_gutter_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the text for [param gutter] on [param line] to [param text]. This only works when the gutter type is [constant GUTTER_TYPE_STRING] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) SetLineGutterText(line gd.Int, gutter gd.Int, text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_gutter_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the text currently in [param gutter] at [param line]. This only works when the gutter type is [constant GUTTER_TYPE_STRING] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) GetLineGutterText(line gd.Int, gutter gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_gutter_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the icon for [param gutter] on [param line] to [param icon]. This only works when the gutter type is [constant GUTTER_TYPE_ICON] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) SetLineGutterIcon(line gd.Int, gutter gd.Int, icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_gutter_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the icon currently in [param gutter] at [param line]. This only works when the gutter type is [constant GUTTER_TYPE_ICON] (see [method set_gutter_type]).
*/
//go:nosplit
func (self class) GetLineGutterIcon(line gd.Int, gutter gd.Int) objects.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_gutter_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets the color for [param gutter] on [param line] to [param color].
*/
//go:nosplit
func (self class) SetLineGutterItemColor(line gd.Int, gutter gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_gutter_item_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the color currently in [param gutter] at [param line].
*/
//go:nosplit
func (self class) GetLineGutterItemColor(line gd.Int, gutter gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_gutter_item_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param clickable] is [code]true[/code], makes the [param gutter] on [param line] clickable. See [signal gutter_clicked].
*/
//go:nosplit
func (self class) SetLineGutterClickable(line gd.Int, gutter gd.Int, clickable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	callframe.Arg(frame, clickable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_gutter_clickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the gutter on the given line is clickable.
*/
//go:nosplit
func (self class) IsLineGutterClickable(line gd.Int, gutter gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, gutter)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_line_gutter_clickable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the current background color of the line. Set to [code]Color(0, 0, 0, 0)[/code] for no color.
*/
//go:nosplit
func (self class) SetLineBackgroundColor(line gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_line_background_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the current background color of the line. [code]Color(0, 0, 0, 0)[/code] is returned if no color is set.
*/
//go:nosplit
func (self class) GetLineBackgroundColor(line gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_line_background_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSyntaxHighlighter(syntax_highlighter objects.SyntaxHighlighter) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(syntax_highlighter[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSyntaxHighlighter() objects.SyntaxHighlighter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.SyntaxHighlighter{classdb.SyntaxHighlighter(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHighlightCurrentLine(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_highlight_current_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsHighlightCurrentLineEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_highlight_current_line_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHighlightAllOccurrences(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_highlight_all_occurrences, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsHighlightAllOccurrencesEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_highlight_all_occurrences_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetDrawControlChars() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_draw_control_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawControlChars(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_draw_control_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetDrawTabs(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_draw_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawingTabs() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_drawing_tabs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawSpaces(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_set_draw_spaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawingSpaces() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_drawing_spaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [PopupMenu] of this [TextEdit]. By default, this menu is displayed when right-clicking on the [TextEdit].
You can add custom menu items or remove standard ones. Make sure your IDs don't conflict with the standard ones (see [enum MenuItems]). For example:
[codeblocks]
[gdscript]
func _ready():
    var menu = get_menu()
    # Remove all items after "Redo".
    menu.item_count = menu.get_item_index(MENU_REDO) + 1
    # Add custom items.
    menu.add_separator()
    menu.add_item("Insert Date", MENU_MAX + 1)
    # Connect callback.
    menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):
    if id == MENU_MAX + 1:
        insert_text_at_caret(Time.get_date_string_from_system())
[/gdscript]
[csharp]
public override void _Ready()
{
    var menu = GetMenu();
    // Remove all items after "Redo".
    menu.ItemCount = menu.GetItemIndex(TextEdit.MenuItems.Redo) + 1;
    // Add custom items.
    menu.AddSeparator();
    menu.AddItem("Insert Date", TextEdit.MenuItems.Max + 1);
    // Add event handler.
    menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
    if (id == TextEdit.MenuItems.Max + 1)
    {
        InsertTextAtCaret(Time.GetDateStringFromSystem());
    }
}
[/csharp]
[/codeblocks]
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetMenu() objects.PopupMenu {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.PopupMenu{classdb.PopupMenu(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
//go:nosplit
func (self class) IsMenuVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_is_menu_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
//go:nosplit
func (self class) MenuOption(option gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_menu_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
This method does nothing.
*/
//go:nosplit
func (self class) AdjustCaretsAfterEdit(caret gd.Int, from_line gd.Int, from_col gd.Int, to_line gd.Int, to_col gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, caret)
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, from_col)
	callframe.Arg(frame, to_line)
	callframe.Arg(frame, to_col)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_adjust_carets_after_edit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of caret indexes in their edit order, this done from bottom to top. Edit order refers to the way actions such as [method insert_text_at_caret] are applied.
*/
//go:nosplit
func (self class) GetCaretIndexEditOrder() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_caret_index_edit_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the original start line of the selection.
*/
//go:nosplit
func (self class) GetSelectionLine(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the original start column of the selection.
*/
//go:nosplit
func (self class) GetSelectionColumn(caret_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, caret_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextEdit.Bind_get_selection_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnTextSet(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("text_set"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTextChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("text_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnLinesEditedFrom(cb func(from_line int, to_line int)) {
	self[0].AsObject().Connect(gd.NewStringName("lines_edited_from"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCaretChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("caret_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGutterClicked(cb func(line int, gutter int)) {
	self[0].AsObject().Connect(gd.NewStringName("gutter_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGutterAdded(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("gutter_added"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGutterRemoved(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("gutter_removed"), gd.NewCallable(cb), 0)
}

func (self class) AsTextEdit() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextEdit() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_handle_unicode_input":
		return reflect.ValueOf(self._handle_unicode_input)
	case "_backspace":
		return reflect.ValueOf(self._backspace)
	case "_cut":
		return reflect.ValueOf(self._cut)
	case "_copy":
		return reflect.ValueOf(self._copy)
	case "_paste":
		return reflect.ValueOf(self._paste)
	case "_paste_primary_clipboard":
		return reflect.ValueOf(self._paste_primary_clipboard)
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_handle_unicode_input":
		return reflect.ValueOf(self._handle_unicode_input)
	case "_backspace":
		return reflect.ValueOf(self._backspace)
	case "_cut":
		return reflect.ValueOf(self._cut)
	case "_copy":
		return reflect.ValueOf(self._copy)
	case "_paste":
		return reflect.ValueOf(self._paste)
	case "_paste_primary_clipboard":
		return reflect.ValueOf(self._paste_primary_clipboard)
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() { classdb.Register("TextEdit", func(ptr gd.Object) any { return classdb.TextEdit(ptr) }) }

type MenuItems = classdb.TextEditMenuItems

const (
	/*Cuts (copies and clears) the selected text.*/
	MenuCut MenuItems = 0
	/*Copies the selected text.*/
	MenuCopy MenuItems = 1
	/*Pastes the clipboard text over the selected text (or at the cursor's position).*/
	MenuPaste MenuItems = 2
	/*Erases the whole [TextEdit] text.*/
	MenuClear MenuItems = 3
	/*Selects the whole [TextEdit] text.*/
	MenuSelectAll MenuItems = 4
	/*Undoes the previous action.*/
	MenuUndo MenuItems = 5
	/*Redoes the previous action.*/
	MenuRedo MenuItems = 6
	/*ID of "Text Writing Direction" submenu.*/
	MenuSubmenuTextDir MenuItems = 7
	/*Sets text direction to inherited.*/
	MenuDirInherited MenuItems = 8
	/*Sets text direction to automatic.*/
	MenuDirAuto MenuItems = 9
	/*Sets text direction to left-to-right.*/
	MenuDirLtr MenuItems = 10
	/*Sets text direction to right-to-left.*/
	MenuDirRtl MenuItems = 11
	/*Toggles control character display.*/
	MenuDisplayUcc MenuItems = 12
	/*ID of "Insert Control Character" submenu.*/
	MenuSubmenuInsertUcc MenuItems = 13
	/*Inserts left-to-right mark (LRM) character.*/
	MenuInsertLrm MenuItems = 14
	/*Inserts right-to-left mark (RLM) character.*/
	MenuInsertRlm MenuItems = 15
	/*Inserts start of left-to-right embedding (LRE) character.*/
	MenuInsertLre MenuItems = 16
	/*Inserts start of right-to-left embedding (RLE) character.*/
	MenuInsertRle MenuItems = 17
	/*Inserts start of left-to-right override (LRO) character.*/
	MenuInsertLro MenuItems = 18
	/*Inserts start of right-to-left override (RLO) character.*/
	MenuInsertRlo MenuItems = 19
	/*Inserts pop direction formatting (PDF) character.*/
	MenuInsertPdf MenuItems = 20
	/*Inserts Arabic letter mark (ALM) character.*/
	MenuInsertAlm MenuItems = 21
	/*Inserts left-to-right isolate (LRI) character.*/
	MenuInsertLri MenuItems = 22
	/*Inserts right-to-left isolate (RLI) character.*/
	MenuInsertRli MenuItems = 23
	/*Inserts first strong isolate (FSI) character.*/
	MenuInsertFsi MenuItems = 24
	/*Inserts pop direction isolate (PDI) character.*/
	MenuInsertPdi MenuItems = 25
	/*Inserts zero width joiner (ZWJ) character.*/
	MenuInsertZwj MenuItems = 26
	/*Inserts zero width non-joiner (ZWNJ) character.*/
	MenuInsertZwnj MenuItems = 27
	/*Inserts word joiner (WJ) character.*/
	MenuInsertWj MenuItems = 28
	/*Inserts soft hyphen (SHY) character.*/
	MenuInsertShy MenuItems = 29
	/*Represents the size of the [enum MenuItems] enum.*/
	MenuMax MenuItems = 30
)

type EditAction = classdb.TextEditEditAction

const (
	/*No current action.*/
	ActionNone EditAction = 0
	/*A typing action.*/
	ActionTyping EditAction = 1
	/*A backwards delete action.*/
	ActionBackspace EditAction = 2
	/*A forward delete action.*/
	ActionDelete EditAction = 3
)

type SearchFlags = classdb.TextEditSearchFlags

const (
	/*Match case when searching.*/
	SearchMatchCase SearchFlags = 1
	/*Match whole words when searching.*/
	SearchWholeWords SearchFlags = 2
	/*Search from end to beginning.*/
	SearchBackwards SearchFlags = 4
)

type CaretType = classdb.TextEditCaretType

const (
	/*Vertical line caret.*/
	CaretTypeLine CaretType = 0
	/*Block caret.*/
	CaretTypeBlock CaretType = 1
)

type SelectionMode = classdb.TextEditSelectionMode

const (
	/*Not selecting.*/
	SelectionModeNone SelectionMode = 0
	/*Select as if [code]shift[/code] is pressed.*/
	SelectionModeShift SelectionMode = 1
	/*Select single characters as if the user single clicked.*/
	SelectionModePointer SelectionMode = 2
	/*Select whole words as if the user double clicked.*/
	SelectionModeWord SelectionMode = 3
	/*Select whole lines as if the user triple clicked.*/
	SelectionModeLine SelectionMode = 4
)

type LineWrappingMode = classdb.TextEditLineWrappingMode

const (
	/*Line wrapping is disabled.*/
	LineWrappingNone LineWrappingMode = 0
	/*Line wrapping occurs at the control boundary, beyond what would normally be visible.*/
	LineWrappingBoundary LineWrappingMode = 1
)

type GutterType = classdb.TextEditGutterType

const (
	/*When a gutter is set to string using [method set_gutter_type], it is used to contain text set via the [method set_line_gutter_text] method.*/
	GutterTypeString GutterType = 0
	/*When a gutter is set to icon using [method set_gutter_type], it is used to contain an icon set via the [method set_line_gutter_icon] method.*/
	GutterTypeIcon GutterType = 1
	/*When a gutter is set to custom using [method set_gutter_type], it is used to contain custom visuals controlled by a callback method set via the [method set_gutter_custom_draw] method.*/
	GutterTypeCustom GutterType = 2
)
