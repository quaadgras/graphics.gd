package CodeEdit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/TextEdit"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
CodeEdit is a specialized [TextEdit] designed for editing plain text code files. It has many features commonly found in code editors such as line numbers, line folding, code completion, indent management, and string/comment management.
[b]Note:[/b] Regardless of locale, [CodeEdit] will by default always use left-to-right text direction to correctly display source code.
	// CodeEdit methods that can be overridden by a [Class] that extends it.
	type CodeEdit interface {
		//Override this method to define how the selected entry should be inserted. If [param replace] is [code]true[/code], any existing text should be replaced.
		ConfirmCodeCompletion(replace bool) 
		//Override this method to define what happens when the user requests code completion. If [param force] is [code]true[/code], any checks should be bypassed.
		RequestCodeCompletion(force bool) 
		//Override this method to define what items in [param candidates] should be displayed.
		//Both [param candidates] and the return is a [Array] of [Dictionary], see [method get_code_completion_option] for [Dictionary] content.
		FilterCodeCompletionCandidates(candidates gd.ArrayOf[gd.Dictionary]) gd.ArrayOf[gd.Dictionary]
	}

*/
type Go [1]classdb.CodeEdit

/*
Override this method to define how the selected entry should be inserted. If [param replace] is [code]true[/code], any existing text should be replaced.
*/
func (Go) _confirm_code_completion(impl func(ptr unsafe.Pointer, replace bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var replace = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, replace)
		gc.End()
	}
}

/*
Override this method to define what happens when the user requests code completion. If [param force] is [code]true[/code], any checks should be bypassed.
*/
func (Go) _request_code_completion(impl func(ptr unsafe.Pointer, force bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var force = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		gc.End()
	}
}

/*
Override this method to define what items in [param candidates] should be displayed.
Both [param candidates] and the return is a [Array] of [Dictionary], see [method get_code_completion_option] for [Dictionary] content.
*/
func (Go) _filter_code_completion_candidates(impl func(ptr unsafe.Pointer, candidates gd.ArrayOf[gd.Dictionary]) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var candidates = gd.TypedArray[gd.Dictionary](mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0)))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, candidates)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}

/*
Perform an indent as if the user activated the "ui_text_indent" action.
*/
func (self Go) DoIndent() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DoIndent()
}

/*
Indents selected lines, or in the case of no selection the caret line by one.
*/
func (self Go) IndentLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).IndentLines()
}

/*
Unindents selected lines, or in the case of no selection the caret line by one. Same as performing "ui_text_unindent" action.
*/
func (self Go) UnindentLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UnindentLines()
}

/*
Converts the indents of lines between [param from_line] and [param to_line] to tabs or spaces as set by [member indent_use_spaces].
Values of [code]-1[/code] convert the entire text.
*/
func (self Go) ConvertIndent() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ConvertIndent(gd.Int(-1), gd.Int(-1))
}

/*
Adds a brace pair.
Both the start and end keys must be symbols. Only the start key has to be unique.
*/
func (self Go) AddAutoBraceCompletionPair(start_key string, end_key string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddAutoBraceCompletionPair(gc.String(start_key), gc.String(end_key))
}

/*
Returns [code]true[/code] if open key [param open_key] exists.
*/
func (self Go) HasAutoBraceCompletionOpenKey(open_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasAutoBraceCompletionOpenKey(gc.String(open_key)))
}

/*
Returns [code]true[/code] if close key [param close_key] exists.
*/
func (self Go) HasAutoBraceCompletionCloseKey(close_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasAutoBraceCompletionCloseKey(gc.String(close_key)))
}

/*
Gets the matching auto brace close key for [param open_key].
*/
func (self Go) GetAutoBraceCompletionCloseKey(open_key string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetAutoBraceCompletionCloseKey(gc, gc.String(open_key)).String())
}

/*
Sets the line as breakpointed.
*/
func (self Go) SetLineAsBreakpoint(line int, breakpointed bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineAsBreakpoint(gd.Int(line), breakpointed)
}

/*
Returns whether the line at the specified index is breakpointed or not.
*/
func (self Go) IsLineBreakpointed(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLineBreakpointed(gd.Int(line)))
}

/*
Clears all breakpointed lines.
*/
func (self Go) ClearBreakpointedLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearBreakpointedLines()
}

/*
Gets all breakpointed lines.
*/
func (self Go) GetBreakpointedLines() []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetBreakpointedLines(gc).AsSlice())
}

/*
Sets the line as bookmarked.
*/
func (self Go) SetLineAsBookmarked(line int, bookmarked bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineAsBookmarked(gd.Int(line), bookmarked)
}

/*
Returns whether the line at the specified index is bookmarked or not.
*/
func (self Go) IsLineBookmarked(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLineBookmarked(gd.Int(line)))
}

/*
Clears all bookmarked lines.
*/
func (self Go) ClearBookmarkedLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearBookmarkedLines()
}

/*
Gets all bookmarked lines.
*/
func (self Go) GetBookmarkedLines() []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetBookmarkedLines(gc).AsSlice())
}

/*
Sets the line as executing.
*/
func (self Go) SetLineAsExecuting(line int, executing bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineAsExecuting(gd.Int(line), executing)
}

/*
Returns whether the line at the specified index is marked as executing or not.
*/
func (self Go) IsLineExecuting(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLineExecuting(gd.Int(line)))
}

/*
Clears all executed lines.
*/
func (self Go) ClearExecutingLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearExecutingLines()
}

/*
Gets all executing lines.
*/
func (self Go) GetExecutingLines() []int32 {
	gc := gd.GarbageCollector(); _ = gc
	return []int32(class(self).GetExecutingLines(gc).AsSlice())
}

/*
Returns if the given line is foldable, that is, it has indented lines right below it or a comment / string block.
*/
func (self Go) CanFoldLine(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).CanFoldLine(gd.Int(line)))
}

/*
Folds the given line, if possible (see [method can_fold_line]).
*/
func (self Go) FoldLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FoldLine(gd.Int(line))
}

/*
Unfolds all lines that were previously folded.
*/
func (self Go) UnfoldLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UnfoldLine(gd.Int(line))
}

/*
Folds all lines that are possible to be folded (see [method can_fold_line]).
*/
func (self Go) FoldAllLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).FoldAllLines()
}

/*
Unfolds all lines, folded or not.
*/
func (self Go) UnfoldAllLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UnfoldAllLines()
}

/*
Toggle the folding of the code block at the given line.
*/
func (self Go) ToggleFoldableLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ToggleFoldableLine(gd.Int(line))
}

/*
Toggle the folding of the code block on all lines with a caret on them.
*/
func (self Go) ToggleFoldableLinesAtCarets() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ToggleFoldableLinesAtCarets()
}

/*
Returns whether the line at the specified index is folded or not.
*/
func (self Go) IsLineFolded(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLineFolded(gd.Int(line)))
}

/*
Returns all lines that are current folded.
*/
func (self Go) GetFoldedLines() gd.ArrayOf[gd.Int] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Int](class(self).GetFoldedLines(gc))
}

/*
Creates a new code region with the selection. At least one single line comment delimiter have to be defined (see [method add_comment_delimiter]).
A code region is a part of code that is highlighted when folded and can help organize your script.
Code region start and end tags can be customized (see [method set_code_region_tags]).
Code regions are delimited using start and end tags (respectively [code]region[/code] and [code]endregion[/code] by default) preceded by one line comment delimiter. (eg. [code]#region[/code] and [code]#endregion[/code])
*/
func (self Go) CreateCodeRegion() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CreateCodeRegion()
}

/*
Returns the code region start tag (without comment delimiter).
*/
func (self Go) GetCodeRegionStartTag() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetCodeRegionStartTag(gc).String())
}

/*
Returns the code region end tag (without comment delimiter).
*/
func (self Go) GetCodeRegionEndTag() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetCodeRegionEndTag(gc).String())
}

/*
Sets the code region start and end tags (without comment delimiter).
*/
func (self Go) SetCodeRegionTags() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCodeRegionTags(gc.String("region"), gc.String("endregion"))
}

/*
Returns whether the line at the specified index is a code region start.
*/
func (self Go) IsLineCodeRegionStart(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLineCodeRegionStart(gd.Int(line)))
}

/*
Returns whether the line at the specified index is a code region end.
*/
func (self Go) IsLineCodeRegionEnd(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsLineCodeRegionEnd(gd.Int(line)))
}

/*
Defines a string delimiter from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
func (self Go) AddStringDelimiter(start_key string, end_key string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddStringDelimiter(gc.String(start_key), gc.String(end_key), false)
}

/*
Removes the string delimiter with [param start_key].
*/
func (self Go) RemoveStringDelimiter(start_key string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveStringDelimiter(gc.String(start_key))
}

/*
Returns [code]true[/code] if string [param start_key] exists.
*/
func (self Go) HasStringDelimiter(start_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasStringDelimiter(gc.String(start_key)))
}

/*
Removes all string delimiters.
*/
func (self Go) ClearStringDelimiters() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearStringDelimiters()
}

/*
Returns the delimiter index if [param line] [param column] is in a string. If [param column] is not provided, will return the delimiter index if the entire [param line] is a string. Otherwise [code]-1[/code].
*/
func (self Go) IsInString(line int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).IsInString(gd.Int(line), gd.Int(-1))))
}

/*
Adds a comment delimiter from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
func (self Go) AddCommentDelimiter(start_key string, end_key string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddCommentDelimiter(gc.String(start_key), gc.String(end_key), false)
}

/*
Removes the comment delimiter with [param start_key].
*/
func (self Go) RemoveCommentDelimiter(start_key string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveCommentDelimiter(gc.String(start_key))
}

/*
Returns [code]true[/code] if comment [param start_key] exists.
*/
func (self Go) HasCommentDelimiter(start_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasCommentDelimiter(gc.String(start_key)))
}

/*
Removes all comment delimiters.
*/
func (self Go) ClearCommentDelimiters() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearCommentDelimiters()
}

/*
Returns delimiter index if [param line] [param column] is in a comment. If [param column] is not provided, will return delimiter index if the entire [param line] is a comment. Otherwise [code]-1[/code].
*/
func (self Go) IsInComment(line int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).IsInComment(gd.Int(line), gd.Int(-1))))
}

/*
Gets the start key for a string or comment region index.
*/
func (self Go) GetDelimiterStartKey(delimiter_index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetDelimiterStartKey(gc, gd.Int(delimiter_index)).String())
}

/*
Gets the end key for a string or comment region index.
*/
func (self Go) GetDelimiterEndKey(delimiter_index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetDelimiterEndKey(gc, gd.Int(delimiter_index)).String())
}

/*
If [param line] [param column] is in a string or comment, returns the start position of the region. If not or no start could be found, both [Vector2] values will be [code]-1[/code].
*/
func (self Go) GetDelimiterStartPosition(line int, column int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetDelimiterStartPosition(gd.Int(line), gd.Int(column)))
}

/*
If [param line] [param column] is in a string or comment, returns the end position of the region. If not or no end could be found, both [Vector2] values will be [code]-1[/code].
*/
func (self Go) GetDelimiterEndPosition(line int, column int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(class(self).GetDelimiterEndPosition(gd.Int(line), gd.Int(column)))
}

/*
Sets the code hint text. Pass an empty string to clear.
*/
func (self Go) SetCodeHint(code_hint string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCodeHint(gc.String(code_hint))
}

/*
Sets if the code hint should draw below the text.
*/
func (self Go) SetCodeHintDrawBelow(draw_below bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCodeHintDrawBelow(draw_below)
}

/*
Returns the full text with char [code]0xFFFF[/code] at the caret location.
*/
func (self Go) GetTextForCodeCompletion() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTextForCodeCompletion(gc).String())
}

/*
Emits [signal code_completion_requested], if [param force] is [code]true[/code] will bypass all checks. Otherwise will check that the caret is in a word or in front of a prefix. Will ignore the request if all current options are of type file path, node path, or signal.
*/
func (self Go) RequestCodeCompletion() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RequestCodeCompletion(false)
}

/*
Submits an item to the queue of potential candidates for the autocomplete menu. Call [method update_code_completion_options] to update the list.
[param location] indicates location of the option relative to the location of the code completion query. See [enum CodeEdit.CodeCompletionLocation] for how to set this value.
[b]Note:[/b] This list will replace all current candidates.
*/
func (self Go) AddCodeCompletionOption(atype classdb.CodeEditCodeCompletionKind, display_text string, insert_text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddCodeCompletionOption(atype, gc.String(display_text), gc.String(insert_text), gd.Color{1, 1, 1, 1}, ([1]gdclass.Resource{}[0]), gc.Variant(([1]gd.Variant{}[0])), gd.Int(1024))
}

/*
Submits all completion options added with [method add_code_completion_option]. Will try to force the autocomplete menu to popup, if [param force] is [code]true[/code].
[b]Note:[/b] This will replace all current candidates.
*/
func (self Go) UpdateCodeCompletionOptions(force bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UpdateCodeCompletionOptions(force)
}

/*
Gets all completion options, see [method get_code_completion_option] for return content.
*/
func (self Go) GetCodeCompletionOptions() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](class(self).GetCodeCompletionOptions(gc))
}

/*
Gets the completion option at [param index]. The return [Dictionary] has the following key-values:
[code]kind[/code]: [enum CodeCompletionKind]
[code]display_text[/code]: Text that is shown on the autocomplete menu.
[code]insert_text[/code]: Text that is to be inserted when this item is selected.
[code]font_color[/code]: Color of the text on the autocomplete menu.
[code]icon[/code]: Icon to draw on the autocomplete menu.
[code]default_value[/code]: Value of the symbol.
*/
func (self Go) GetCodeCompletionOption(index int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetCodeCompletionOption(gc, gd.Int(index)))
}

/*
Gets the index of the current selected completion option.
*/
func (self Go) GetCodeCompletionSelectedIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCodeCompletionSelectedIndex()))
}

/*
Sets the current selected completion option.
*/
func (self Go) SetCodeCompletionSelectedIndex(index int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCodeCompletionSelectedIndex(gd.Int(index))
}

/*
Inserts the selected entry into the text. If [param replace] is [code]true[/code], any existing text is replaced rather than merged.
*/
func (self Go) ConfirmCodeCompletion() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ConfirmCodeCompletion(false)
}

/*
Cancels the autocomplete menu.
*/
func (self Go) CancelCodeCompletion() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CancelCodeCompletion()
}

/*
Returns the full text with char [code]0xFFFF[/code] at the cursor location.
*/
func (self Go) GetTextForSymbolLookup() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTextForSymbolLookup(gc).String())
}

/*
Returns the full text with char [code]0xFFFF[/code] at the specified location.
*/
func (self Go) GetTextWithCursorChar(line int, column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetTextWithCursorChar(gc, gd.Int(line), gd.Int(column)).String())
}

/*
Sets the symbol emitted by [signal symbol_validate] as a valid lookup.
*/
func (self Go) SetSymbolLookupWordAsValid(valid bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSymbolLookupWordAsValid(valid)
}

/*
Moves all lines up that are selected or have a caret on them.
*/
func (self Go) MoveLinesUp() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MoveLinesUp()
}

/*
Moves all lines down that are selected or have a caret on them.
*/
func (self Go) MoveLinesDown() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MoveLinesDown()
}

/*
Deletes all lines that are selected or have a caret on them.
*/
func (self Go) DeleteLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DeleteLines()
}

/*
Duplicates all selected text and duplicates all lines with a caret on them.
*/
func (self Go) DuplicateSelection() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DuplicateSelection()
}

/*
Duplicates all lines currently selected with any caret. Duplicates the entire line beneath the current one no matter where the caret is within the line.
*/
func (self Go) DuplicateLines() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DuplicateLines()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CodeEdit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("CodeEdit"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) SymbolLookupOnClick() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSymbolLookupOnClickEnabled())
}

func (self Go) SetSymbolLookupOnClick(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSymbolLookupOnClickEnabled(value)
}

func (self Go) LineFolding() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsLineFoldingEnabled())
}

func (self Go) SetLineFolding(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineFoldingEnabled(value)
}

func (self Go) LineLengthGuidelines() gd.ArrayOf[gd.Int] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.Int](class(self).GetLineLengthGuidelines(gc))
}

func (self Go) SetLineLengthGuidelines(value gd.ArrayOf[gd.Int]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineLengthGuidelines(value)
}

func (self Go) GuttersDrawBreakpointsGutter() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawingBreakpointsGutter())
}

func (self Go) SetGuttersDrawBreakpointsGutter(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawBreakpointsGutter(value)
}

func (self Go) GuttersDrawBookmarks() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawingBookmarksGutter())
}

func (self Go) SetGuttersDrawBookmarks(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawBookmarksGutter(value)
}

func (self Go) GuttersDrawExecutingLines() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawingExecutingLinesGutter())
}

func (self Go) SetGuttersDrawExecutingLines(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawExecutingLinesGutter(value)
}

func (self Go) GuttersDrawLineNumbers() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawLineNumbersEnabled())
}

func (self Go) SetGuttersDrawLineNumbers(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawLineNumbers(value)
}

func (self Go) GuttersZeroPadLineNumbers() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsLineNumbersZeroPadded())
}

func (self Go) SetGuttersZeroPadLineNumbers(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLineNumbersZeroPadded(value)
}

func (self Go) GuttersDrawFoldGutter() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawingFoldGutter())
}

func (self Go) SetGuttersDrawFoldGutter(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawFoldGutter(value)
}

func (self Go) DelimiterStrings() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.String](class(self).GetStringDelimiters(gc))
}

func (self Go) SetDelimiterStrings(value gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStringDelimiters(value)
}

func (self Go) DelimiterComments() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.String](class(self).GetCommentDelimiters(gc))
}

func (self Go) SetDelimiterComments(value gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCommentDelimiters(value)
}

func (self Go) CodeCompletionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCodeCompletionEnabled())
}

func (self Go) SetCodeCompletionEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCodeCompletionEnabled(value)
}

func (self Go) CodeCompletionPrefixes() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.String](class(self).GetCodeCompletionPrefixes(gc))
}

func (self Go) SetCodeCompletionPrefixes(value gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCodeCompletionPrefixes(value)
}

func (self Go) IndentSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetIndentSize()))
}

func (self Go) SetIndentSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIndentSize(gd.Int(value))
}

func (self Go) IndentUseSpaces() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsIndentUsingSpaces())
}

func (self Go) SetIndentUseSpaces(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetIndentUsingSpaces(value)
}

func (self Go) IndentAutomatic() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAutoIndentEnabled())
}

func (self Go) SetIndentAutomatic(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoIndentEnabled(value)
}

func (self Go) IndentAutomaticPrefixes() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
		return gd.ArrayOf[gd.String](class(self).GetAutoIndentPrefixes(gc))
}

func (self Go) SetIndentAutomaticPrefixes(value gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoIndentPrefixes(value)
}

func (self Go) AutoBraceCompletionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAutoBraceCompletionEnabled())
}

func (self Go) SetAutoBraceCompletionEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoBraceCompletionEnabled(value)
}

func (self Go) AutoBraceCompletionHighlightMatching() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsHighlightMatchingBracesEnabled())
}

func (self Go) SetAutoBraceCompletionHighlightMatching(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHighlightMatchingBracesEnabled(value)
}

func (self Go) AutoBraceCompletionPairs() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Dictionary(class(self).GetAutoBraceCompletionPairs(gc))
}

func (self Go) SetAutoBraceCompletionPairs(value gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutoBraceCompletionPairs(value)
}

/*
Override this method to define how the selected entry should be inserted. If [param replace] is [code]true[/code], any existing text should be replaced.
*/
func (class) _confirm_code_completion(impl func(ptr unsafe.Pointer, replace bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var replace = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, replace)
		ctx.End()
	}
}

/*
Override this method to define what happens when the user requests code completion. If [param force] is [code]true[/code], any checks should be bypassed.
*/
func (class) _request_code_completion(impl func(ptr unsafe.Pointer, force bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

/*
Override this method to define what items in [param candidates] should be displayed.
Both [param candidates] and the return is a [Array] of [Dictionary], see [method get_code_completion_option] for [Dictionary] content.
*/
func (class) _filter_code_completion_candidates(impl func(ptr unsafe.Pointer, candidates gd.ArrayOf[gd.Dictionary]) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var candidates = gd.TypedArray[gd.Dictionary](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0)))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, candidates)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

//go:nosplit
func (self class) SetIndentSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_indent_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIndentSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_indent_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIndentUsingSpaces(use_spaces bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_spaces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_indent_using_spaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsIndentUsingSpaces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_indent_using_spaces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoIndentEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_auto_indent_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoIndentEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_auto_indent_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoIndentPrefixes(prefixes gd.ArrayOf[gd.String])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(prefixes))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_auto_indent_prefixes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoIndentPrefixes(ctx gd.Lifetime) gd.ArrayOf[gd.String] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_auto_indent_prefixes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.String](ret)
}
/*
Perform an indent as if the user activated the "ui_text_indent" action.
*/
//go:nosplit
func (self class) DoIndent()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_do_indent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Indents selected lines, or in the case of no selection the caret line by one.
*/
//go:nosplit
func (self class) IndentLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_indent_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unindents selected lines, or in the case of no selection the caret line by one. Same as performing "ui_text_unindent" action.
*/
//go:nosplit
func (self class) UnindentLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_unindent_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Converts the indents of lines between [param from_line] and [param to_line] to tabs or spaces as set by [member indent_use_spaces].
Values of [code]-1[/code] convert the entire text.
*/
//go:nosplit
func (self class) ConvertIndent(from_line gd.Int, to_line gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, to_line)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_convert_indent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAutoBraceCompletionEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_auto_brace_completion_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoBraceCompletionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_auto_brace_completion_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHighlightMatchingBracesEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_highlight_matching_braces_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHighlightMatchingBracesEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_highlight_matching_braces_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a brace pair.
Both the start and end keys must be symbols. Only the start key has to be unique.
*/
//go:nosplit
func (self class) AddAutoBraceCompletionPair(start_key gd.String, end_key gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	callframe.Arg(frame, mmm.Get(end_key))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_add_auto_brace_completion_pair, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAutoBraceCompletionPairs(pairs gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(pairs))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_auto_brace_completion_pairs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoBraceCompletionPairs(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_auto_brace_completion_pairs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if open key [param open_key] exists.
*/
//go:nosplit
func (self class) HasAutoBraceCompletionOpenKey(open_key gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(open_key))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_has_auto_brace_completion_open_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if close key [param close_key] exists.
*/
//go:nosplit
func (self class) HasAutoBraceCompletionCloseKey(close_key gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(close_key))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_has_auto_brace_completion_close_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the matching auto brace close key for [param open_key].
*/
//go:nosplit
func (self class) GetAutoBraceCompletionCloseKey(ctx gd.Lifetime, open_key gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(open_key))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_auto_brace_completion_close_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawBreakpointsGutter(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_draw_breakpoints_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawingBreakpointsGutter() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_drawing_breakpoints_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawBookmarksGutter(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_draw_bookmarks_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawingBookmarksGutter() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_drawing_bookmarks_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawExecutingLinesGutter(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_draw_executing_lines_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawingExecutingLinesGutter() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_drawing_executing_lines_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the line as breakpointed.
*/
//go:nosplit
func (self class) SetLineAsBreakpoint(line gd.Int, breakpointed bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, breakpointed)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_line_as_breakpoint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the line at the specified index is breakpointed or not.
*/
//go:nosplit
func (self class) IsLineBreakpointed(line gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_line_breakpointed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears all breakpointed lines.
*/
//go:nosplit
func (self class) ClearBreakpointedLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_clear_breakpointed_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets all breakpointed lines.
*/
//go:nosplit
func (self class) GetBreakpointedLines(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_breakpointed_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the line as bookmarked.
*/
//go:nosplit
func (self class) SetLineAsBookmarked(line gd.Int, bookmarked bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, bookmarked)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_line_as_bookmarked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the line at the specified index is bookmarked or not.
*/
//go:nosplit
func (self class) IsLineBookmarked(line gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_line_bookmarked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears all bookmarked lines.
*/
//go:nosplit
func (self class) ClearBookmarkedLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_clear_bookmarked_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets all bookmarked lines.
*/
//go:nosplit
func (self class) GetBookmarkedLines(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_bookmarked_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the line as executing.
*/
//go:nosplit
func (self class) SetLineAsExecuting(line gd.Int, executing bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, executing)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_line_as_executing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the line at the specified index is marked as executing or not.
*/
//go:nosplit
func (self class) IsLineExecuting(line gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_line_executing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears all executed lines.
*/
//go:nosplit
func (self class) ClearExecutingLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_clear_executing_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets all executing lines.
*/
//go:nosplit
func (self class) GetExecutingLines(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_executing_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawLineNumbers(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_draw_line_numbers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawLineNumbersEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_draw_line_numbers_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLineNumbersZeroPadded(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_line_numbers_zero_padded, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLineNumbersZeroPadded() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_line_numbers_zero_padded, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawFoldGutter(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_draw_fold_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawingFoldGutter() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_drawing_fold_gutter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLineFoldingEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_line_folding_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLineFoldingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_line_folding_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if the given line is foldable, that is, it has indented lines right below it or a comment / string block.
*/
//go:nosplit
func (self class) CanFoldLine(line gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_can_fold_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Folds the given line, if possible (see [method can_fold_line]).
*/
//go:nosplit
func (self class) FoldLine(line gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_fold_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unfolds all lines that were previously folded.
*/
//go:nosplit
func (self class) UnfoldLine(line gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_unfold_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Folds all lines that are possible to be folded (see [method can_fold_line]).
*/
//go:nosplit
func (self class) FoldAllLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_fold_all_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unfolds all lines, folded or not.
*/
//go:nosplit
func (self class) UnfoldAllLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_unfold_all_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Toggle the folding of the code block at the given line.
*/
//go:nosplit
func (self class) ToggleFoldableLine(line gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_toggle_foldable_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Toggle the folding of the code block on all lines with a caret on them.
*/
//go:nosplit
func (self class) ToggleFoldableLinesAtCarets()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_toggle_foldable_lines_at_carets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the line at the specified index is folded or not.
*/
//go:nosplit
func (self class) IsLineFolded(line gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_line_folded, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns all lines that are current folded.
*/
//go:nosplit
func (self class) GetFoldedLines(ctx gd.Lifetime) gd.ArrayOf[gd.Int] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_folded_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Int](ret)
}
/*
Creates a new code region with the selection. At least one single line comment delimiter have to be defined (see [method add_comment_delimiter]).
A code region is a part of code that is highlighted when folded and can help organize your script.
Code region start and end tags can be customized (see [method set_code_region_tags]).
Code regions are delimited using start and end tags (respectively [code]region[/code] and [code]endregion[/code] by default) preceded by one line comment delimiter. (eg. [code]#region[/code] and [code]#endregion[/code])
*/
//go:nosplit
func (self class) CreateCodeRegion()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_create_code_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the code region start tag (without comment delimiter).
*/
//go:nosplit
func (self class) GetCodeRegionStartTag(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_code_region_start_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the code region end tag (without comment delimiter).
*/
//go:nosplit
func (self class) GetCodeRegionEndTag(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_code_region_end_tag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the code region start and end tags (without comment delimiter).
*/
//go:nosplit
func (self class) SetCodeRegionTags(start gd.String, end gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start))
	callframe.Arg(frame, mmm.Get(end))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_code_region_tags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the line at the specified index is a code region start.
*/
//go:nosplit
func (self class) IsLineCodeRegionStart(line gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_line_code_region_start, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether the line at the specified index is a code region end.
*/
//go:nosplit
func (self class) IsLineCodeRegionEnd(line gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_line_code_region_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Defines a string delimiter from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
//go:nosplit
func (self class) AddStringDelimiter(start_key gd.String, end_key gd.String, line_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	callframe.Arg(frame, mmm.Get(end_key))
	callframe.Arg(frame, line_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_add_string_delimiter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the string delimiter with [param start_key].
*/
//go:nosplit
func (self class) RemoveStringDelimiter(start_key gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_remove_string_delimiter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if string [param start_key] exists.
*/
//go:nosplit
func (self class) HasStringDelimiter(start_key gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_has_string_delimiter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStringDelimiters(string_delimiters gd.ArrayOf[gd.String])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(string_delimiters))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_string_delimiters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all string delimiters.
*/
//go:nosplit
func (self class) ClearStringDelimiters()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_clear_string_delimiters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStringDelimiters(ctx gd.Lifetime) gd.ArrayOf[gd.String] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_string_delimiters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.String](ret)
}
/*
Returns the delimiter index if [param line] [param column] is in a string. If [param column] is not provided, will return the delimiter index if the entire [param line] is a string. Otherwise [code]-1[/code].
*/
//go:nosplit
func (self class) IsInString(line gd.Int, column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_in_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a comment delimiter from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
//go:nosplit
func (self class) AddCommentDelimiter(start_key gd.String, end_key gd.String, line_only bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	callframe.Arg(frame, mmm.Get(end_key))
	callframe.Arg(frame, line_only)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_add_comment_delimiter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the comment delimiter with [param start_key].
*/
//go:nosplit
func (self class) RemoveCommentDelimiter(start_key gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_remove_comment_delimiter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if comment [param start_key] exists.
*/
//go:nosplit
func (self class) HasCommentDelimiter(start_key gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(start_key))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_has_comment_delimiter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCommentDelimiters(comment_delimiters gd.ArrayOf[gd.String])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(comment_delimiters))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_comment_delimiters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all comment delimiters.
*/
//go:nosplit
func (self class) ClearCommentDelimiters()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_clear_comment_delimiters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCommentDelimiters(ctx gd.Lifetime) gd.ArrayOf[gd.String] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_comment_delimiters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.String](ret)
}
/*
Returns delimiter index if [param line] [param column] is in a comment. If [param column] is not provided, will return delimiter index if the entire [param line] is a comment. Otherwise [code]-1[/code].
*/
//go:nosplit
func (self class) IsInComment(line gd.Int, column gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_in_comment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the start key for a string or comment region index.
*/
//go:nosplit
func (self class) GetDelimiterStartKey(ctx gd.Lifetime, delimiter_index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delimiter_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_delimiter_start_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Gets the end key for a string or comment region index.
*/
//go:nosplit
func (self class) GetDelimiterEndKey(ctx gd.Lifetime, delimiter_index gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delimiter_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_delimiter_end_key, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If [param line] [param column] is in a string or comment, returns the start position of the region. If not or no start could be found, both [Vector2] values will be [code]-1[/code].
*/
//go:nosplit
func (self class) GetDelimiterStartPosition(line gd.Int, column gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_delimiter_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [param line] [param column] is in a string or comment, returns the end position of the region. If not or no end could be found, both [Vector2] values will be [code]-1[/code].
*/
//go:nosplit
func (self class) GetDelimiterEndPosition(line gd.Int, column gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_delimiter_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the code hint text. Pass an empty string to clear.
*/
//go:nosplit
func (self class) SetCodeHint(code_hint gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(code_hint))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_code_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets if the code hint should draw below the text.
*/
//go:nosplit
func (self class) SetCodeHintDrawBelow(draw_below bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_below)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_code_hint_draw_below, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the full text with char [code]0xFFFF[/code] at the caret location.
*/
//go:nosplit
func (self class) GetTextForCodeCompletion(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_text_for_code_completion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Emits [signal code_completion_requested], if [param force] is [code]true[/code] will bypass all checks. Otherwise will check that the caret is in a word or in front of a prefix. Will ignore the request if all current options are of type file path, node path, or signal.
*/
//go:nosplit
func (self class) RequestCodeCompletion(force bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_request_code_completion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Submits an item to the queue of potential candidates for the autocomplete menu. Call [method update_code_completion_options] to update the list.
[param location] indicates location of the option relative to the location of the code completion query. See [enum CodeEdit.CodeCompletionLocation] for how to set this value.
[b]Note:[/b] This list will replace all current candidates.
*/
//go:nosplit
func (self class) AddCodeCompletionOption(atype classdb.CodeEditCodeCompletionKind, display_text gd.String, insert_text gd.String, text_color gd.Color, icon gdclass.Resource, value gd.Variant, location gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, mmm.Get(display_text))
	callframe.Arg(frame, mmm.Get(insert_text))
	callframe.Arg(frame, text_color)
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(value))
	callframe.Arg(frame, location)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_add_code_completion_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Submits all completion options added with [method add_code_completion_option]. Will try to force the autocomplete menu to popup, if [param force] is [code]true[/code].
[b]Note:[/b] This will replace all current candidates.
*/
//go:nosplit
func (self class) UpdateCodeCompletionOptions(force bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_update_code_completion_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets all completion options, see [method get_code_completion_option] for return content.
*/
//go:nosplit
func (self class) GetCodeCompletionOptions(ctx gd.Lifetime) gd.ArrayOf[gd.Dictionary] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_code_completion_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Dictionary](ret)
}
/*
Gets the completion option at [param index]. The return [Dictionary] has the following key-values:
[code]kind[/code]: [enum CodeCompletionKind]
[code]display_text[/code]: Text that is shown on the autocomplete menu.
[code]insert_text[/code]: Text that is to be inserted when this item is selected.
[code]font_color[/code]: Color of the text on the autocomplete menu.
[code]icon[/code]: Icon to draw on the autocomplete menu.
[code]default_value[/code]: Value of the symbol.
*/
//go:nosplit
func (self class) GetCodeCompletionOption(ctx gd.Lifetime, index gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_code_completion_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Gets the index of the current selected completion option.
*/
//go:nosplit
func (self class) GetCodeCompletionSelectedIndex() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_code_completion_selected_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the current selected completion option.
*/
//go:nosplit
func (self class) SetCodeCompletionSelectedIndex(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_code_completion_selected_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Inserts the selected entry into the text. If [param replace] is [code]true[/code], any existing text is replaced rather than merged.
*/
//go:nosplit
func (self class) ConfirmCodeCompletion(replace bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, replace)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_confirm_code_completion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Cancels the autocomplete menu.
*/
//go:nosplit
func (self class) CancelCodeCompletion()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_cancel_code_completion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCodeCompletionEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_code_completion_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCodeCompletionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_code_completion_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCodeCompletionPrefixes(prefixes gd.ArrayOf[gd.String])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(prefixes))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_code_completion_prefixes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCodeCompletionPrefixes(ctx gd.Lifetime) gd.ArrayOf[gd.String] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_code_completion_prefixes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.String](ret)
}
//go:nosplit
func (self class) SetLineLengthGuidelines(guideline_columns gd.ArrayOf[gd.Int])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(guideline_columns))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_line_length_guidelines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLineLengthGuidelines(ctx gd.Lifetime) gd.ArrayOf[gd.Int] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_line_length_guidelines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Int](ret)
}
//go:nosplit
func (self class) SetSymbolLookupOnClickEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_symbol_lookup_on_click_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSymbolLookupOnClickEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_is_symbol_lookup_on_click_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the full text with char [code]0xFFFF[/code] at the cursor location.
*/
//go:nosplit
func (self class) GetTextForSymbolLookup(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_text_for_symbol_lookup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the full text with char [code]0xFFFF[/code] at the specified location.
*/
//go:nosplit
func (self class) GetTextWithCursorChar(ctx gd.Lifetime, line gd.Int, column gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_get_text_with_cursor_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the symbol emitted by [signal symbol_validate] as a valid lookup.
*/
//go:nosplit
func (self class) SetSymbolLookupWordAsValid(valid bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, valid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_set_symbol_lookup_word_as_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves all lines up that are selected or have a caret on them.
*/
//go:nosplit
func (self class) MoveLinesUp()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_move_lines_up, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves all lines down that are selected or have a caret on them.
*/
//go:nosplit
func (self class) MoveLinesDown()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_move_lines_down, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes all lines that are selected or have a caret on them.
*/
//go:nosplit
func (self class) DeleteLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_delete_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Duplicates all selected text and duplicates all lines with a caret on them.
*/
//go:nosplit
func (self class) DuplicateSelection()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_duplicate_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Duplicates all lines currently selected with any caret. Duplicates the entire line beneath the current one no matter where the caret is within the line.
*/
//go:nosplit
func (self class) DuplicateLines()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CodeEdit.Bind_duplicate_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnBreakpointToggled(cb func(line int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("breakpoint_toggled"), gc.Callable(cb), 0)
}


func (self Go) OnCodeCompletionRequested(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("code_completion_requested"), gc.Callable(cb), 0)
}


func (self Go) OnSymbolLookup(cb func(symbol string, line int, column int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("symbol_lookup"), gc.Callable(cb), 0)
}


func (self Go) OnSymbolValidate(cb func(symbol string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("symbol_validate"), gc.Callable(cb), 0)
}


func (self class) AsCodeEdit() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCodeEdit() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTextEdit() TextEdit.GD { return *((*TextEdit.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextEdit() TextEdit.Go { return *((*TextEdit.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_confirm_code_completion": return reflect.ValueOf(self._confirm_code_completion);
	case "_request_code_completion": return reflect.ValueOf(self._request_code_completion);
	case "_filter_code_completion_candidates": return reflect.ValueOf(self._filter_code_completion_candidates);
	default: return gd.VirtualByName(self.AsTextEdit(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_confirm_code_completion": return reflect.ValueOf(self._confirm_code_completion);
	case "_request_code_completion": return reflect.ValueOf(self._request_code_completion);
	case "_filter_code_completion_candidates": return reflect.ValueOf(self._filter_code_completion_candidates);
	default: return gd.VirtualByName(self.AsTextEdit(), name)
	}
}
func init() {classdb.Register("CodeEdit", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type CodeCompletionKind = classdb.CodeEditCodeCompletionKind

const (
/*Marks the option as a class.*/
	KindClass CodeCompletionKind = 0
/*Marks the option as a function.*/
	KindFunction CodeCompletionKind = 1
/*Marks the option as a Godot signal.*/
	KindSignal CodeCompletionKind = 2
/*Marks the option as a variable.*/
	KindVariable CodeCompletionKind = 3
/*Marks the option as a member.*/
	KindMember CodeCompletionKind = 4
/*Marks the option as an enum entry.*/
	KindEnum CodeCompletionKind = 5
/*Marks the option as a constant.*/
	KindConstant CodeCompletionKind = 6
/*Marks the option as a Godot node path.*/
	KindNodePath CodeCompletionKind = 7
/*Marks the option as a file path.*/
	KindFilePath CodeCompletionKind = 8
/*Marks the option as unclassified or plain text.*/
	KindPlainText CodeCompletionKind = 9
)
type CodeCompletionLocation = classdb.CodeEditCodeCompletionLocation

const (
/*The option is local to the location of the code completion query - e.g. a local variable. Subsequent value of location represent options from the outer class, the exact value represent how far they are (in terms of inner classes).*/
	LocationLocal CodeCompletionLocation = 0
/*The option is from the containing class or a parent class, relative to the location of the code completion query. Perform a bitwise OR with the class depth (e.g. [code]0[/code] for the local class, [code]1[/code] for the parent, [code]2[/code] for the grandparent, etc.) to store the depth of an option in the class or a parent class.*/
	LocationParentMask CodeCompletionLocation = 256
/*The option is from user code which is not local and not in a derived class (e.g. Autoload Singletons).*/
	LocationOtherUserCode CodeCompletionLocation = 512
/*The option is from other engine code, not covered by the other enum constants - e.g. built-in classes.*/
	LocationOther CodeCompletionLocation = 1024
)
