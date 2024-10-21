package LineEdit

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
[LineEdit] provides an input field for editing a single line of text. It features many built-in shortcuts that are always available ([kbd]Ctrl[/kbd] here maps to [kbd]Cmd[/kbd] on macOS):
- [kbd]Ctrl + C[/kbd]: Copy
- [kbd]Ctrl + X[/kbd]: Cut
- [kbd]Ctrl + V[/kbd] or [kbd]Ctrl + Y[/kbd]: Paste/"yank"
- [kbd]Ctrl + Z[/kbd]: Undo
- [kbd]Ctrl + ~[/kbd]: Swap input direction.
- [kbd]Ctrl + Shift + Z[/kbd]: Redo
- [kbd]Ctrl + U[/kbd]: Delete text from the caret position to the beginning of the line
- [kbd]Ctrl + K[/kbd]: Delete text from the caret position to the end of the line
- [kbd]Ctrl + A[/kbd]: Select all text
- [kbd]Up Arrow[/kbd]/[kbd]Down Arrow[/kbd]: Move the caret to the beginning/end of the line
On macOS, some extra keyboard shortcuts are available:
- [kbd]Cmd + F[/kbd]: Same as [kbd]Right Arrow[/kbd], move the caret one character right
- [kbd]Cmd + B[/kbd]: Same as [kbd]Left Arrow[/kbd], move the caret one character left
- [kbd]Cmd + P[/kbd]: Same as [kbd]Up Arrow[/kbd], move the caret to the previous line
- [kbd]Cmd + N[/kbd]: Same as [kbd]Down Arrow[/kbd], move the caret to the next line
- [kbd]Cmd + D[/kbd]: Same as [kbd]Delete[/kbd], delete the character on the right side of caret
- [kbd]Cmd + H[/kbd]: Same as [kbd]Backspace[/kbd], delete the character on the left side of the caret
- [kbd]Cmd + A[/kbd]: Same as [kbd]Home[/kbd], move the caret to the beginning of the line
- [kbd]Cmd + E[/kbd]: Same as [kbd]End[/kbd], move the caret to the end of the line
- [kbd]Cmd + Left Arrow[/kbd]: Same as [kbd]Home[/kbd], move the caret to the beginning of the line
- [kbd]Cmd + Right Arrow[/kbd]: Same as [kbd]End[/kbd], move the caret to the end of the line

*/
type Simple [1]classdb.LineEdit
func (self Simple) SetHorizontalAlignment(alignment gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHorizontalAlignment(alignment)
}
func (self Simple) GetHorizontalAlignment() gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
	return gd.HorizontalAlignment(Expert(self).GetHorizontalAlignment())
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) Select(from int, to int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Select(gd.Int(from), gd.Int(to))
}
func (self Simple) SelectAll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SelectAll()
}
func (self Simple) Deselect() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Deselect()
}
func (self Simple) HasSelection() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasSelection())
}
func (self Simple) GetSelectedText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSelectedText(gc).String())
}
func (self Simple) GetSelectionFromColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSelectionFromColumn()))
}
func (self Simple) GetSelectionToColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSelectionToColumn()))
}
func (self Simple) SetText(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetText(gc.String(text))
}
func (self Simple) GetText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetText(gc).String())
}
func (self Simple) GetDrawControlChars() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDrawControlChars())
}
func (self Simple) SetDrawControlChars(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDrawControlChars(enable)
}
func (self Simple) SetTextDirection(direction classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextDirection(direction)
}
func (self Simple) GetTextDirection() classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlTextDirection(Expert(self).GetTextDirection())
}
func (self Simple) SetLanguage(language string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLanguage(gc.String(language))
}
func (self Simple) GetLanguage() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLanguage(gc).String())
}
func (self Simple) SetStructuredTextBidiOverride(parser classdb.TextServerStructuredTextParser) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStructuredTextBidiOverride(parser)
}
func (self Simple) GetStructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerStructuredTextParser(Expert(self).GetStructuredTextBidiOverride())
}
func (self Simple) SetStructuredTextBidiOverrideOptions(args gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStructuredTextBidiOverrideOptions(args)
}
func (self Simple) GetStructuredTextBidiOverrideOptions() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetStructuredTextBidiOverrideOptions(gc))
}
func (self Simple) SetPlaceholder(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPlaceholder(gc.String(text))
}
func (self Simple) GetPlaceholder() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetPlaceholder(gc).String())
}
func (self Simple) SetCaretColumn(position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCaretColumn(gd.Int(position))
}
func (self Simple) GetCaretColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCaretColumn()))
}
func (self Simple) GetScrollOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetScrollOffset()))
}
func (self Simple) SetExpandToTextLengthEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExpandToTextLengthEnabled(enabled)
}
func (self Simple) IsExpandToTextLengthEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsExpandToTextLengthEnabled())
}
func (self Simple) SetCaretBlinkEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCaretBlinkEnabled(enabled)
}
func (self Simple) IsCaretBlinkEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCaretBlinkEnabled())
}
func (self Simple) SetCaretMidGraphemeEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCaretMidGraphemeEnabled(enabled)
}
func (self Simple) IsCaretMidGraphemeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCaretMidGraphemeEnabled())
}
func (self Simple) SetCaretForceDisplayed(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCaretForceDisplayed(enabled)
}
func (self Simple) IsCaretForceDisplayed() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCaretForceDisplayed())
}
func (self Simple) SetCaretBlinkInterval(interval float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCaretBlinkInterval(gd.Float(interval))
}
func (self Simple) GetCaretBlinkInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCaretBlinkInterval()))
}
func (self Simple) SetMaxLength(chars int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxLength(gd.Int(chars))
}
func (self Simple) GetMaxLength() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMaxLength()))
}
func (self Simple) InsertTextAtCaret(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InsertTextAtCaret(gc.String(text))
}
func (self Simple) DeleteCharAtCaret() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DeleteCharAtCaret()
}
func (self Simple) DeleteText(from_column int, to_column int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).DeleteText(gd.Int(from_column), gd.Int(to_column))
}
func (self Simple) SetEditable(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEditable(enabled)
}
func (self Simple) IsEditable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEditable())
}
func (self Simple) SetSecret(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSecret(enabled)
}
func (self Simple) IsSecret() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSecret())
}
func (self Simple) SetSecretCharacter(character string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSecretCharacter(gc.String(character))
}
func (self Simple) GetSecretCharacter() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSecretCharacter(gc).String())
}
func (self Simple) MenuOption(option int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MenuOption(gd.Int(option))
}
func (self Simple) GetMenu() [1]classdb.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PopupMenu(Expert(self).GetMenu(gc))
}
func (self Simple) IsMenuVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMenuVisible())
}
func (self Simple) SetContextMenuEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContextMenuEnabled(enable)
}
func (self Simple) IsContextMenuEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsContextMenuEnabled())
}
func (self Simple) SetVirtualKeyboardEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVirtualKeyboardEnabled(enable)
}
func (self Simple) IsVirtualKeyboardEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsVirtualKeyboardEnabled())
}
func (self Simple) SetVirtualKeyboardType(atype classdb.LineEditVirtualKeyboardType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVirtualKeyboardType(atype)
}
func (self Simple) GetVirtualKeyboardType() classdb.LineEditVirtualKeyboardType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.LineEditVirtualKeyboardType(Expert(self).GetVirtualKeyboardType())
}
func (self Simple) SetClearButtonEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetClearButtonEnabled(enable)
}
func (self Simple) IsClearButtonEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsClearButtonEnabled())
}
func (self Simple) SetShortcutKeysEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShortcutKeysEnabled(enable)
}
func (self Simple) IsShortcutKeysEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShortcutKeysEnabled())
}
func (self Simple) SetMiddleMousePasteEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMiddleMousePasteEnabled(enable)
}
func (self Simple) IsMiddleMousePasteEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMiddleMousePasteEnabled())
}
func (self Simple) SetSelectingEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelectingEnabled(enable)
}
func (self Simple) IsSelectingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSelectingEnabled())
}
func (self Simple) SetDeselectOnFocusLossEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDeselectOnFocusLossEnabled(enable)
}
func (self Simple) IsDeselectOnFocusLossEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDeselectOnFocusLossEnabled())
}
func (self Simple) SetDragAndDropSelectionEnabled(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDragAndDropSelectionEnabled(enable)
}
func (self Simple) IsDragAndDropSelectionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDragAndDropSelectionEnabled())
}
func (self Simple) SetRightIcon(icon [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRightIcon(icon)
}
func (self Simple) GetRightIcon() [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetRightIcon(gc))
}
func (self Simple) SetFlat(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFlat(enabled)
}
func (self Simple) IsFlat() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFlat())
}
func (self Simple) SetSelectAllOnFocus(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelectAllOnFocus(enabled)
}
func (self Simple) IsSelectAllOnFocus() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSelectAllOnFocus())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.LineEdit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetHorizontalAlignment(alignment gd.HorizontalAlignment)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHorizontalAlignment() gd.HorizontalAlignment {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Erases the [LineEdit]'s [member text].
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Selects characters inside [LineEdit] between [param from] and [param to]. By default, [param from] is at the beginning and [param to] at the end.
[codeblocks]
[gdscript]
text = "Welcome"
select() # Will select "Welcome".
select(4) # Will select "ome".
select(2, 5) # Will select "lco".
[/gdscript]
[csharp]
Text = "Welcome";
Select(); // Will select "Welcome".
Select(4); // Will select "ome".
Select(2, 5); // Will select "lco".
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) Select(from gd.Int, to gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Selects the whole [String].
*/
//go:nosplit
func (self class) SelectAll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_select_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the current selection.
*/
//go:nosplit
func (self class) Deselect()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the user has selected text.
*/
//go:nosplit
func (self class) HasSelection() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_has_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text inside the selection.
*/
//go:nosplit
func (self class) GetSelectedText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_selected_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the selection begin column.
*/
//go:nosplit
func (self class) GetSelectionFromColumn() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_selection_from_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the selection end column.
*/
//go:nosplit
func (self class) GetSelectionToColumn() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_selection_to_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetDrawControlChars() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_draw_control_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawControlChars(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_draw_control_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLanguage(language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLanguage(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser classdb.TextServerStructuredTextParser)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(args))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPlaceholder(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPlaceholder(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretColumn(position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_caret_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCaretColumn() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_caret_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the scroll offset due to [member caret_column], as a number of characters.
*/
//go:nosplit
func (self class) GetScrollOffset() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExpandToTextLengthEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_expand_to_text_length_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsExpandToTextLengthEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_expand_to_text_length_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretBlinkEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_caret_blink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCaretBlinkEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_caret_blink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretMidGraphemeEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_caret_mid_grapheme_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCaretMidGraphemeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_caret_mid_grapheme_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretForceDisplayed(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_caret_force_displayed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCaretForceDisplayed() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_caret_force_displayed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCaretBlinkInterval(interval gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_caret_blink_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCaretBlinkInterval() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_caret_blink_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxLength(chars gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, chars)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_max_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_max_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts [param text] at the caret. If the resulting value is longer than [member max_length], nothing happens.
*/
//go:nosplit
func (self class) InsertTextAtCaret(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_insert_text_at_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes one character at the caret's current position (equivalent to pressing [kbd]Delete[/kbd]).
*/
//go:nosplit
func (self class) DeleteCharAtCaret()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_delete_char_at_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes a section of the [member text] going from position [param from_column] to [param to_column]. Both parameters should be within the text's length.
*/
//go:nosplit
func (self class) DeleteText(from_column gd.Int, to_column gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from_column)
	callframe.Arg(frame, to_column)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_delete_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetEditable(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEditable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSecret(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_secret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSecret() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_secret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSecretCharacter(character gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(character))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_secret_character, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSecretCharacter(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_secret_character, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
//go:nosplit
func (self class) MenuOption(option gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_menu_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [PopupMenu] of this [LineEdit]. By default, this menu is displayed when right-clicking on the [LineEdit].
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
    menu.ItemCount = menu.GetItemIndex(LineEdit.MenuItems.Redo) + 1;
    // Add custom items.
    menu.AddSeparator();
    menu.AddItem("Insert Date", LineEdit.MenuItems.Max + 1);
    // Add event handler.
    menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
    if (id == LineEdit.MenuItems.Max + 1)
    {
        InsertTextAtCaret(Time.GetDateStringFromSystem());
    }
}
[/csharp]
[/codeblocks]
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetMenu(ctx gd.Lifetime) object.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.PopupMenu
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
//go:nosplit
func (self class) IsMenuVisible() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_menu_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContextMenuEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsContextMenuEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVirtualKeyboardEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_virtual_keyboard_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVirtualKeyboardEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_virtual_keyboard_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVirtualKeyboardType(atype classdb.LineEditVirtualKeyboardType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_virtual_keyboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVirtualKeyboardType() classdb.LineEditVirtualKeyboardType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.LineEditVirtualKeyboardType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_virtual_keyboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetClearButtonEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_clear_button_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsClearButtonEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_clear_button_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcutKeysEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShortcutKeysEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMiddleMousePasteEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_middle_mouse_paste_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMiddleMousePasteEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_middle_mouse_paste_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectingEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_selecting_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSelectingEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_selecting_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDeselectOnFocusLossEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeselectOnFocusLossEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDragAndDropSelectionEnabled(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDragAndDropSelectionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRightIcon(icon object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_right_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRightIcon(ctx gd.Lifetime) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_right_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlat(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFlat() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectAllOnFocus(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_select_all_on_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSelectAllOnFocus() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_is_select_all_on_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsLineEdit() Expert { return self[0].AsLineEdit() }


//go:nosplit
func (self Simple) AsLineEdit() Simple { return self[0].AsLineEdit() }


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
func init() {classdb.Register("LineEdit", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type MenuItems = classdb.LineEditMenuItems

const (
/*Cuts (copies and clears) the selected text.*/
	MenuCut MenuItems = 0
/*Copies the selected text.*/
	MenuCopy MenuItems = 1
/*Pastes the clipboard text over the selected text (or at the caret's position).
Non-printable escape characters are automatically stripped from the OS clipboard via [method String.strip_escapes].*/
	MenuPaste MenuItems = 2
/*Erases the whole [LineEdit] text.*/
	MenuClear MenuItems = 3
/*Selects the whole [LineEdit] text.*/
	MenuSelectAll MenuItems = 4
/*Undoes the previous action.*/
	MenuUndo MenuItems = 5
/*Reverse the last undo action.*/
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
type VirtualKeyboardType = classdb.LineEditVirtualKeyboardType

const (
/*Default text virtual keyboard.*/
	KeyboardTypeDefault VirtualKeyboardType = 0
/*Multiline virtual keyboard.*/
	KeyboardTypeMultiline VirtualKeyboardType = 1
/*Virtual number keypad, useful for PIN entry.*/
	KeyboardTypeNumber VirtualKeyboardType = 2
/*Virtual number keypad, useful for entering fractional numbers.*/
	KeyboardTypeNumberDecimal VirtualKeyboardType = 3
/*Virtual phone number keypad.*/
	KeyboardTypePhone VirtualKeyboardType = 4
/*Virtual keyboard with additional keys to assist with typing email addresses.*/
	KeyboardTypeEmailAddress VirtualKeyboardType = 5
/*Virtual keyboard for entering a password. On most platforms, this should disable autocomplete and autocapitalization.
[b]Note:[/b] This is not supported on Web. Instead, this behaves identically to [constant KEYBOARD_TYPE_DEFAULT].*/
	KeyboardTypePassword VirtualKeyboardType = 6
/*Virtual keyboard with additional keys to assist with typing URLs.*/
	KeyboardTypeUrl VirtualKeyboardType = 7
)
