package LineEdit

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
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
type Go [1]classdb.LineEdit

/*
Erases the [LineEdit]'s [member text].
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
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
func (self Go) Select() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Select(gd.Int(0), gd.Int(-1))
}

/*
Selects the whole [String].
*/
func (self Go) SelectAll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SelectAll()
}

/*
Clears the current selection.
*/
func (self Go) Deselect() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Deselect()
}

/*
Returns [code]true[/code] if the user has selected text.
*/
func (self Go) HasSelection() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasSelection())
}

/*
Returns the text inside the selection.
*/
func (self Go) GetSelectedText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSelectedText(gc).String())
}

/*
Returns the selection begin column.
*/
func (self Go) GetSelectionFromColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionFromColumn()))
}

/*
Returns the selection end column.
*/
func (self Go) GetSelectionToColumn() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionToColumn()))
}

/*
Returns the scroll offset due to [member caret_column], as a number of characters.
*/
func (self Go) GetScrollOffset() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetScrollOffset()))
}

/*
Inserts [param text] at the caret. If the resulting value is longer than [member max_length], nothing happens.
*/
func (self Go) InsertTextAtCaret(text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).InsertTextAtCaret(gc.String(text))
}

/*
Deletes one character at the caret's current position (equivalent to pressing [kbd]Delete[/kbd]).
*/
func (self Go) DeleteCharAtCaret() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DeleteCharAtCaret()
}

/*
Deletes a section of the [member text] going from position [param from_column] to [param to_column]. Both parameters should be within the text's length.
*/
func (self Go) DeleteText(from_column int, to_column int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DeleteText(gd.Int(from_column), gd.Int(to_column))
}

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
func (self Go) MenuOption(option int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MenuOption(gd.Int(option))
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
func (self Go) GetMenu() gdclass.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.PopupMenu(class(self).GetMenu(gc))
}

/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
func (self Go) IsMenuVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsMenuVisible())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.LineEdit
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("LineEdit"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Text() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetText(gc).String())
}

func (self Go) SetText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetText(gc.String(value))
}

func (self Go) PlaceholderText() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetPlaceholder(gc).String())
}

func (self Go) SetPlaceholderText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetPlaceholder(gc.String(value))
}

func (self Go) Alignment() gd.HorizontalAlignment {
	gc := gd.GarbageCollector(); _ = gc
		return gd.HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Go) SetAlignment(value gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHorizontalAlignment(value)
}

func (self Go) MaxLength() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxLength()))
}

func (self Go) SetMaxLength(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxLength(gd.Int(value))
}

func (self Go) Editable() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsEditable())
}

func (self Go) SetEditable(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEditable(value)
}

func (self Go) ExpandToTextLength() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsExpandToTextLengthEnabled())
}

func (self Go) SetExpandToTextLength(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpandToTextLengthEnabled(value)
}

func (self Go) ContextMenuEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsContextMenuEnabled())
}

func (self Go) SetContextMenuEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetContextMenuEnabled(value)
}

func (self Go) VirtualKeyboardEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsVirtualKeyboardEnabled())
}

func (self Go) SetVirtualKeyboardEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVirtualKeyboardEnabled(value)
}

func (self Go) VirtualKeyboardType() classdb.LineEditVirtualKeyboardType {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.LineEditVirtualKeyboardType(class(self).GetVirtualKeyboardType())
}

func (self Go) SetVirtualKeyboardType(value classdb.LineEditVirtualKeyboardType) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVirtualKeyboardType(value)
}

func (self Go) ClearButtonEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsClearButtonEnabled())
}

func (self Go) SetClearButtonEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetClearButtonEnabled(value)
}

func (self Go) ShortcutKeysEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShortcutKeysEnabled())
}

func (self Go) SetShortcutKeysEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShortcutKeysEnabled(value)
}

func (self Go) MiddleMousePasteEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMiddleMousePasteEnabled())
}

func (self Go) SetMiddleMousePasteEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMiddleMousePasteEnabled(value)
}

func (self Go) SelectingEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSelectingEnabled())
}

func (self Go) SetSelectingEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectingEnabled(value)
}

func (self Go) DeselectOnFocusLossEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDeselectOnFocusLossEnabled())
}

func (self Go) SetDeselectOnFocusLossEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDeselectOnFocusLossEnabled(value)
}

func (self Go) DragAndDropSelectionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDragAndDropSelectionEnabled())
}

func (self Go) SetDragAndDropSelectionEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDragAndDropSelectionEnabled(value)
}

func (self Go) RightIcon() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetRightIcon(gc))
}

func (self Go) SetRightIcon(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRightIcon(value)
}

func (self Go) Flat() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFlat())
}

func (self Go) SetFlat(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFlat(value)
}

func (self Go) DrawControlChars() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDrawControlChars())
}

func (self Go) SetDrawControlChars(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawControlChars(value)
}

func (self Go) SelectAllOnFocus() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSelectAllOnFocus())
}

func (self Go) SetSelectAllOnFocus(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectAllOnFocus(value)
}

func (self Go) CaretBlink() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCaretBlinkEnabled())
}

func (self Go) SetCaretBlink(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretBlinkEnabled(value)
}

func (self Go) CaretBlinkInterval() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetCaretBlinkInterval()))
}

func (self Go) SetCaretBlinkInterval(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretBlinkInterval(gd.Float(value))
}

func (self Go) CaretColumn() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCaretColumn()))
}

func (self Go) SetCaretColumn(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretColumn(gd.Int(value))
}

func (self Go) CaretForceDisplayed() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCaretForceDisplayed())
}

func (self Go) SetCaretForceDisplayed(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretForceDisplayed(value)
}

func (self Go) CaretMidGrapheme() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsCaretMidGraphemeEnabled())
}

func (self Go) SetCaretMidGrapheme(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCaretMidGraphemeEnabled(value)
}

func (self Go) Secret() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSecret())
}

func (self Go) SetSecret(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSecret(value)
}

func (self Go) SecretCharacter() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetSecretCharacter(gc).String())
}

func (self Go) SetSecretCharacter(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSecretCharacter(gc.String(value))
}

func (self Go) TextDirection() classdb.ControlTextDirection {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.ControlTextDirection(class(self).GetTextDirection())
}

func (self Go) SetTextDirection(value classdb.ControlTextDirection) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextDirection(value)
}

func (self Go) Language() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetLanguage(gc).String())
}

func (self Go) SetLanguage(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetLanguage(gc.String(value))
}

func (self Go) StructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Go) SetStructuredTextBidiOverride(value classdb.TextServerStructuredTextParser) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Go) StructuredTextBidiOverrideOptions() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetStructuredTextBidiOverrideOptions(gc))
}

func (self Go) SetStructuredTextBidiOverrideOptions(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetStructuredTextBidiOverrideOptions(value)
}

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
func (self class) GetMenu(ctx gd.Lifetime) gdclass.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.PopupMenu
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
func (self class) SetRightIcon(icon gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(icon[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_set_right_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRightIcon(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.LineEdit.Bind_get_right_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
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
func (self Go) OnTextChanged(cb func(new_text string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("text_changed"), gc.Callable(cb), 0)
}


func (self Go) OnTextChangeRejected(cb func(rejected_substring string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("text_change_rejected"), gc.Callable(cb), 0)
}


func (self Go) OnTextSubmitted(cb func(new_text string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("text_submitted"), gc.Callable(cb), 0)
}


func (self class) AsLineEdit() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLineEdit() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {classdb.Register("LineEdit", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
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
