package CodeEdit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/TextEdit"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
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
type Simple [1]classdb.CodeEdit
func (Simple) _confirm_code_completion(impl func(ptr unsafe.Pointer, replace bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var replace = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, replace)
		gc.End()
	}
}
func (Simple) _request_code_completion(impl func(ptr unsafe.Pointer, force bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var force = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		gc.End()
	}
}
func (Simple) _filter_code_completion_candidates(impl func(ptr unsafe.Pointer, candidates gd.ArrayOf[gd.Dictionary]) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (self Simple) SetIndentSize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIndentSize(gd.Int(size))
}
func (self Simple) GetIndentSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetIndentSize()))
}
func (self Simple) SetIndentUsingSpaces(use_spaces bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIndentUsingSpaces(use_spaces)
}
func (self Simple) IsIndentUsingSpaces() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsIndentUsingSpaces())
}
func (self Simple) SetAutoIndentEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoIndentEnabled(enable)
}
func (self Simple) IsAutoIndentEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAutoIndentEnabled())
}
func (self Simple) SetAutoIndentPrefixes(prefixes gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoIndentPrefixes(prefixes)
}
func (self Simple) GetAutoIndentPrefixes() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.String](Expert(self).GetAutoIndentPrefixes(gc))
}
func (self Simple) DoIndent() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DoIndent()
}
func (self Simple) IndentLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).IndentLines()
}
func (self Simple) UnindentLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnindentLines()
}
func (self Simple) ConvertIndent(from_line int, to_line int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ConvertIndent(gd.Int(from_line), gd.Int(to_line))
}
func (self Simple) SetAutoBraceCompletionEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoBraceCompletionEnabled(enable)
}
func (self Simple) IsAutoBraceCompletionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAutoBraceCompletionEnabled())
}
func (self Simple) SetHighlightMatchingBracesEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHighlightMatchingBracesEnabled(enable)
}
func (self Simple) IsHighlightMatchingBracesEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHighlightMatchingBracesEnabled())
}
func (self Simple) AddAutoBraceCompletionPair(start_key string, end_key string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddAutoBraceCompletionPair(gc.String(start_key), gc.String(end_key))
}
func (self Simple) SetAutoBraceCompletionPairs(pairs gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoBraceCompletionPairs(pairs)
}
func (self Simple) GetAutoBraceCompletionPairs() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetAutoBraceCompletionPairs(gc))
}
func (self Simple) HasAutoBraceCompletionOpenKey(open_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAutoBraceCompletionOpenKey(gc.String(open_key)))
}
func (self Simple) HasAutoBraceCompletionCloseKey(close_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasAutoBraceCompletionCloseKey(gc.String(close_key)))
}
func (self Simple) GetAutoBraceCompletionCloseKey(open_key string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetAutoBraceCompletionCloseKey(gc, gc.String(open_key)).String())
}
func (self Simple) SetDrawBreakpointsGutter(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDrawBreakpointsGutter(enable)
}
func (self Simple) IsDrawingBreakpointsGutter() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDrawingBreakpointsGutter())
}
func (self Simple) SetDrawBookmarksGutter(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDrawBookmarksGutter(enable)
}
func (self Simple) IsDrawingBookmarksGutter() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDrawingBookmarksGutter())
}
func (self Simple) SetDrawExecutingLinesGutter(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDrawExecutingLinesGutter(enable)
}
func (self Simple) IsDrawingExecutingLinesGutter() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDrawingExecutingLinesGutter())
}
func (self Simple) SetLineAsBreakpoint(line int, breakpointed bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLineAsBreakpoint(gd.Int(line), breakpointed)
}
func (self Simple) IsLineBreakpointed(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLineBreakpointed(gd.Int(line)))
}
func (self Simple) ClearBreakpointedLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearBreakpointedLines()
}
func (self Simple) GetBreakpointedLines() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetBreakpointedLines(gc))
}
func (self Simple) SetLineAsBookmarked(line int, bookmarked bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLineAsBookmarked(gd.Int(line), bookmarked)
}
func (self Simple) IsLineBookmarked(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLineBookmarked(gd.Int(line)))
}
func (self Simple) ClearBookmarkedLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearBookmarkedLines()
}
func (self Simple) GetBookmarkedLines() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetBookmarkedLines(gc))
}
func (self Simple) SetLineAsExecuting(line int, executing bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLineAsExecuting(gd.Int(line), executing)
}
func (self Simple) IsLineExecuting(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLineExecuting(gd.Int(line)))
}
func (self Simple) ClearExecutingLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearExecutingLines()
}
func (self Simple) GetExecutingLines() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetExecutingLines(gc))
}
func (self Simple) SetDrawLineNumbers(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDrawLineNumbers(enable)
}
func (self Simple) IsDrawLineNumbersEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDrawLineNumbersEnabled())
}
func (self Simple) SetLineNumbersZeroPadded(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLineNumbersZeroPadded(enable)
}
func (self Simple) IsLineNumbersZeroPadded() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLineNumbersZeroPadded())
}
func (self Simple) SetDrawFoldGutter(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDrawFoldGutter(enable)
}
func (self Simple) IsDrawingFoldGutter() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDrawingFoldGutter())
}
func (self Simple) SetLineFoldingEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLineFoldingEnabled(enabled)
}
func (self Simple) IsLineFoldingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLineFoldingEnabled())
}
func (self Simple) CanFoldLine(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).CanFoldLine(gd.Int(line)))
}
func (self Simple) FoldLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FoldLine(gd.Int(line))
}
func (self Simple) UnfoldLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnfoldLine(gd.Int(line))
}
func (self Simple) FoldAllLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FoldAllLines()
}
func (self Simple) UnfoldAllLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnfoldAllLines()
}
func (self Simple) ToggleFoldableLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ToggleFoldableLine(gd.Int(line))
}
func (self Simple) ToggleFoldableLinesAtCarets() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ToggleFoldableLinesAtCarets()
}
func (self Simple) IsLineFolded(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLineFolded(gd.Int(line)))
}
func (self Simple) GetFoldedLines() gd.ArrayOf[gd.Int] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Int](Expert(self).GetFoldedLines(gc))
}
func (self Simple) CreateCodeRegion() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CreateCodeRegion()
}
func (self Simple) GetCodeRegionStartTag() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCodeRegionStartTag(gc).String())
}
func (self Simple) GetCodeRegionEndTag() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetCodeRegionEndTag(gc).String())
}
func (self Simple) SetCodeRegionTags(start string, end string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCodeRegionTags(gc.String(start), gc.String(end))
}
func (self Simple) IsLineCodeRegionStart(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLineCodeRegionStart(gd.Int(line)))
}
func (self Simple) IsLineCodeRegionEnd(line int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLineCodeRegionEnd(gd.Int(line)))
}
func (self Simple) AddStringDelimiter(start_key string, end_key string, line_only bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddStringDelimiter(gc.String(start_key), gc.String(end_key), line_only)
}
func (self Simple) RemoveStringDelimiter(start_key string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveStringDelimiter(gc.String(start_key))
}
func (self Simple) HasStringDelimiter(start_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasStringDelimiter(gc.String(start_key)))
}
func (self Simple) SetStringDelimiters(string_delimiters gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStringDelimiters(string_delimiters)
}
func (self Simple) ClearStringDelimiters() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearStringDelimiters()
}
func (self Simple) GetStringDelimiters() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.String](Expert(self).GetStringDelimiters(gc))
}
func (self Simple) IsInString(line int, column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).IsInString(gd.Int(line), gd.Int(column))))
}
func (self Simple) AddCommentDelimiter(start_key string, end_key string, line_only bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCommentDelimiter(gc.String(start_key), gc.String(end_key), line_only)
}
func (self Simple) RemoveCommentDelimiter(start_key string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveCommentDelimiter(gc.String(start_key))
}
func (self Simple) HasCommentDelimiter(start_key string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasCommentDelimiter(gc.String(start_key)))
}
func (self Simple) SetCommentDelimiters(comment_delimiters gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCommentDelimiters(comment_delimiters)
}
func (self Simple) ClearCommentDelimiters() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearCommentDelimiters()
}
func (self Simple) GetCommentDelimiters() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.String](Expert(self).GetCommentDelimiters(gc))
}
func (self Simple) IsInComment(line int, column int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).IsInComment(gd.Int(line), gd.Int(column))))
}
func (self Simple) GetDelimiterStartKey(delimiter_index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetDelimiterStartKey(gc, gd.Int(delimiter_index)).String())
}
func (self Simple) GetDelimiterEndKey(delimiter_index int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetDelimiterEndKey(gc, gd.Int(delimiter_index)).String())
}
func (self Simple) GetDelimiterStartPosition(line int, column int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetDelimiterStartPosition(gd.Int(line), gd.Int(column)))
}
func (self Simple) GetDelimiterEndPosition(line int, column int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetDelimiterEndPosition(gd.Int(line), gd.Int(column)))
}
func (self Simple) SetCodeHint(code_hint string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCodeHint(gc.String(code_hint))
}
func (self Simple) SetCodeHintDrawBelow(draw_below bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCodeHintDrawBelow(draw_below)
}
func (self Simple) GetTextForCodeCompletion() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTextForCodeCompletion(gc).String())
}
func (self Simple) RequestCodeCompletion(force bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RequestCodeCompletion(force)
}
func (self Simple) AddCodeCompletionOption(atype classdb.CodeEditCodeCompletionKind, display_text string, insert_text string, text_color gd.Color, icon [1]classdb.Resource, value gd.Variant, location int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCodeCompletionOption(atype, gc.String(display_text), gc.String(insert_text), text_color, icon, value, gd.Int(location))
}
func (self Simple) UpdateCodeCompletionOptions(force bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateCodeCompletionOptions(force)
}
func (self Simple) GetCodeCompletionOptions() gd.ArrayOf[gd.Dictionary] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Dictionary](Expert(self).GetCodeCompletionOptions(gc))
}
func (self Simple) GetCodeCompletionOption(index int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetCodeCompletionOption(gc, gd.Int(index)))
}
func (self Simple) GetCodeCompletionSelectedIndex() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCodeCompletionSelectedIndex()))
}
func (self Simple) SetCodeCompletionSelectedIndex(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCodeCompletionSelectedIndex(gd.Int(index))
}
func (self Simple) ConfirmCodeCompletion(replace bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ConfirmCodeCompletion(replace)
}
func (self Simple) CancelCodeCompletion() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CancelCodeCompletion()
}
func (self Simple) SetCodeCompletionEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCodeCompletionEnabled(enable)
}
func (self Simple) IsCodeCompletionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCodeCompletionEnabled())
}
func (self Simple) SetCodeCompletionPrefixes(prefixes gd.ArrayOf[gd.String]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCodeCompletionPrefixes(prefixes)
}
func (self Simple) GetCodeCompletionPrefixes() gd.ArrayOf[gd.String] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.String](Expert(self).GetCodeCompletionPrefixes(gc))
}
func (self Simple) SetLineLengthGuidelines(guideline_columns gd.ArrayOf[gd.Int]) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLineLengthGuidelines(guideline_columns)
}
func (self Simple) GetLineLengthGuidelines() gd.ArrayOf[gd.Int] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Int](Expert(self).GetLineLengthGuidelines(gc))
}
func (self Simple) SetSymbolLookupOnClickEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSymbolLookupOnClickEnabled(enable)
}
func (self Simple) IsSymbolLookupOnClickEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSymbolLookupOnClickEnabled())
}
func (self Simple) GetTextForSymbolLookup() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTextForSymbolLookup(gc).String())
}
func (self Simple) GetTextWithCursorChar(line int, column int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTextWithCursorChar(gc, gd.Int(line), gd.Int(column)).String())
}
func (self Simple) SetSymbolLookupWordAsValid(valid bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSymbolLookupWordAsValid(valid)
}
func (self Simple) MoveLinesUp() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveLinesUp()
}
func (self Simple) MoveLinesDown() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveLinesDown()
}
func (self Simple) DeleteLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DeleteLines()
}
func (self Simple) DuplicateSelection() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DuplicateSelection()
}
func (self Simple) DuplicateLines() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DuplicateLines()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CodeEdit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) AddCodeCompletionOption(atype classdb.CodeEditCodeCompletionKind, display_text gd.String, insert_text gd.String, text_color gd.Color, icon object.Resource, value gd.Variant, location gd.Int)  {
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

//go:nosplit
func (self class) AsCodeEdit() Expert { return self[0].AsCodeEdit() }


//go:nosplit
func (self Simple) AsCodeEdit() Simple { return self[0].AsCodeEdit() }


//go:nosplit
func (self class) AsTextEdit() TextEdit.Expert { return self[0].AsTextEdit() }


//go:nosplit
func (self Simple) AsTextEdit() TextEdit.Simple { return self[0].AsTextEdit() }


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
	case "_confirm_code_completion": return reflect.ValueOf(self._confirm_code_completion);
	case "_request_code_completion": return reflect.ValueOf(self._request_code_completion);
	case "_filter_code_completion_candidates": return reflect.ValueOf(self._filter_code_completion_candidates);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_confirm_code_completion": return reflect.ValueOf(self._confirm_code_completion);
	case "_request_code_completion": return reflect.ValueOf(self._request_code_completion);
	case "_filter_code_completion_candidates": return reflect.ValueOf(self._filter_code_completion_candidates);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("CodeEdit", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
