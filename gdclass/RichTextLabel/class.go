package RichTextLabel

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
A control for displaying text that can contain custom fonts, images, and basic formatting. [RichTextLabel] manages these as an internal tag stack. It also adapts itself to given width/heights.
[b]Note:[/b] Assignments to [member text] clear the tag stack and reconstruct it from the property's contents. Any edits made to [member text] will erase previous edits made from other manual sources such as [method append_text] and the [code]push_*[/code] / [method pop] methods.
[b]Note:[/b] RichTextLabel doesn't support entangled BBCode tags. For example, instead of using [code skip-lint][b]bold[i]bold italic[/b]italic[/i][/code], use [code skip-lint][b]bold[i]bold italic[/i][/b][i]italic[/i][/code].
[b]Note:[/b] [code]push_pop_*[/code] functions won't affect BBCode.
[b]Note:[/b] Unlike [Label], [RichTextLabel] doesn't have a [i]property[/i] to horizontally align text to the center. Instead, enable [member bbcode_enabled] and surround the text in a [code skip-lint][center][/code] tag as follows: [code skip-lint][center]Example[/center][/code]. There is currently no built-in way to vertically align text either, but this can be emulated by relying on anchors/containers and the [member fit_content] property.

*/
type Go [1]classdb.RichTextLabel

/*
Returns the text without BBCode mark-up.
*/
func (self Go) GetParsedText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetParsedText(gc).String())
}

/*
Adds raw non-BBCode-parsed text to the tag stack.
*/
func (self Go) AddText(text string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddText(gc.String(text))
}

/*
Adds an image's opening and closing tags to the tag stack, optionally providing a [param width] and [param height] to resize the image, a [param color] to tint the image and a [param region] to only use parts of the image.
If [param width] or [param height] is set to 0, the image size will be adjusted in order to keep the original aspect ratio.
If [param width] and [param height] are not set, but [param region] is, the region's rect will be used.
[param key] is an optional identifier, that can be used to modify the image via [method update_image].
If [param pad] is set, and the image is smaller than the size specified by [param width] and [param height], the image padding is added to match the size instead of upscaling.
If [param size_in_percent] is set, [param width] and [param height] values are percentages of the control width instead of pixels.
*/
func (self Go) AddImage(image gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddImage(image, gd.Int(0), gd.Int(0), gd.Color{1, 1, 1, 1}, 5, gd.NewRect2(0, 0, 0, 0), gc.Variant(([1]gd.Variant{}[0])), false, gc.String(""), false)
}

/*
Updates the existing images with the key [param key]. Only properties specified by [param mask] bits are updated. See [method add_image].
*/
func (self Go) UpdateImage(key gd.Variant, mask classdb.RichTextLabelImageUpdateMask, image gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UpdateImage(key, mask, image, gd.Int(0), gd.Int(0), gd.Color{1, 1, 1, 1}, 5, gd.NewRect2(0, 0, 0, 0), false, gc.String(""), false)
}

/*
Adds a newline tag to the tag stack.
*/
func (self Go) Newline() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Newline()
}

/*
Removes a paragraph of content from the label. Returns [code]true[/code] if the paragraph exists.
The [param paragraph] argument is the index of the paragraph to remove, it can take values in the interval [code][0, get_paragraph_count() - 1][/code].
If [param no_invalidate] is set to [code]true[/code], cache for the subsequent paragraphs is not invalidated. Use it for faster updates if deleted paragraph is fully self-contained (have no unclosed tags), or this call is part of the complex edit operation and [method invalidate_paragraph] will be called at the end of operation.
*/
func (self Go) RemoveParagraph(paragraph int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).RemoveParagraph(gd.Int(paragraph), false))
}

/*
Invalidates [param paragraph] and all subsequent paragraphs cache.
*/
func (self Go) InvalidateParagraph(paragraph int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).InvalidateParagraph(gd.Int(paragraph)))
}

/*
Adds a [code skip-lint][font][/code] tag to the tag stack. Overrides default fonts for its duration.
Passing [code]0[/code] to [param font_size] will use the existing default font size.
*/
func (self Go) PushFont(font gdclass.Font) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushFont(font, gd.Int(0))
}

/*
Adds a [code skip-lint][font_size][/code] tag to the tag stack. Overrides default font size for its duration.
*/
func (self Go) PushFontSize(font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushFontSize(gd.Int(font_size))
}

/*
Adds a [code skip-lint][font][/code] tag with a normal font to the tag stack.
*/
func (self Go) PushNormal() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushNormal()
}

/*
Adds a [code skip-lint][font][/code] tag with a bold font to the tag stack. This is the same as adding a [code skip-lint][b][/code] tag if not currently in a [code skip-lint][i][/code] tag.
*/
func (self Go) PushBold() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushBold()
}

/*
Adds a [code skip-lint][font][/code] tag with a bold italics font to the tag stack.
*/
func (self Go) PushBoldItalics() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushBoldItalics()
}

/*
Adds a [code skip-lint][font][/code] tag with an italics font to the tag stack. This is the same as adding an [code skip-lint][i][/code] tag if not currently in a [code skip-lint][b][/code] tag.
*/
func (self Go) PushItalics() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushItalics()
}

/*
Adds a [code skip-lint][font][/code] tag with a monospace font to the tag stack.
*/
func (self Go) PushMono() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushMono()
}

/*
Adds a [code skip-lint][color][/code] tag to the tag stack.
*/
func (self Go) PushColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushColor(color)
}

/*
Adds a [code skip-lint][outline_size][/code] tag to the tag stack. Overrides default text outline size for its duration.
*/
func (self Go) PushOutlineSize(outline_size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushOutlineSize(gd.Int(outline_size))
}

/*
Adds a [code skip-lint][outline_color][/code] tag to the tag stack. Adds text outline for its duration.
*/
func (self Go) PushOutlineColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushOutlineColor(color)
}

/*
Adds a [code skip-lint][p][/code] tag to the tag stack.
*/
func (self Go) PushParagraph(alignment gd.HorizontalAlignment) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushParagraph(alignment, 0, gc.String(""), 0, 163, gc.PackedFloat32Slice(([1][]float32{}[0])))
}

/*
Adds an [code skip-lint][indent][/code] tag to the tag stack. Multiplies [param level] by current [member tab_size] to determine new margin length.
*/
func (self Go) PushIndent(level int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushIndent(gd.Int(level))
}

/*
Adds [code skip-lint][ol][/code] or [code skip-lint][ul][/code] tag to the tag stack. Multiplies [param level] by current [member tab_size] to determine new margin length.
*/
func (self Go) PushList(level int, atype classdb.RichTextLabelListType, capitalize bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushList(gd.Int(level), atype, capitalize, gc.String("•"))
}

/*
Adds a meta tag to the tag stack. Similar to the BBCode [code skip-lint][url=something]{text}[/url][/code], but supports non-[String] metadata types.
If [member meta_underlined] is [code]true[/code], meta tags display an underline. This behavior can be customized with [param underline_mode].
[b]Note:[/b] Meta tags do nothing by default when clicked. To assign behavior when clicked, connect [signal meta_clicked] to a function that is called when the meta tag is clicked.
*/
func (self Go) PushMeta(data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushMeta(data, 1)
}

/*
Adds a [code skip-lint][hint][/code] tag to the tag stack. Same as BBCode [code skip-lint][hint=something]{text}[/hint][/code].
*/
func (self Go) PushHint(description string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushHint(gc.String(description))
}

/*
Adds language code used for text shaping algorithm and Open-Type font features.
*/
func (self Go) PushLanguage(language string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushLanguage(gc.String(language))
}

/*
Adds a [code skip-lint][u][/code] tag to the tag stack.
*/
func (self Go) PushUnderline() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushUnderline()
}

/*
Adds a [code skip-lint][s][/code] tag to the tag stack.
*/
func (self Go) PushStrikethrough() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushStrikethrough()
}

/*
Adds a [code skip-lint][table=columns,inline_align][/code] tag to the tag stack. Use [method set_table_column_expand] to set column expansion ratio. Use [method push_cell] to add cells.
*/
func (self Go) PushTable(columns int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushTable(gd.Int(columns), 0, gd.Int(-1))
}

/*
Adds a [code skip-lint][dropcap][/code] tag to the tag stack. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
func (self Go) PushDropcap(s string, font gdclass.Font, size int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushDropcap(gc.String(s), font, gd.Int(size), gd.NewRect2(0, 0, 0, 0), gd.Color{1, 1, 1, 1}, gd.Int(0), gd.Color{0, 0, 0, 0})
}

/*
Edits the selected column's expansion options. If [param expand] is [code]true[/code], the column expands in proportion to its expansion ratio versus the other columns' ratios.
For example, 2 columns with ratios 3 and 4 plus 70 pixels in available width would expand 30 and 40 pixels, respectively.
If [param expand] is [code]false[/code], the column will not contribute to the total ratio.
*/
func (self Go) SetTableColumnExpand(column int, expand bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTableColumnExpand(gd.Int(column), expand, gd.Int(1))
}

/*
Sets color of a table cell. Separate colors for alternating rows can be specified.
*/
func (self Go) SetCellRowBackgroundColor(odd_row_bg gd.Color, even_row_bg gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellRowBackgroundColor(odd_row_bg, even_row_bg)
}

/*
Sets color of a table cell border.
*/
func (self Go) SetCellBorderColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellBorderColor(color)
}

/*
Sets minimum and maximum size overrides for a table cell.
*/
func (self Go) SetCellSizeOverride(min_size gd.Vector2, max_size gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellSizeOverride(min_size, max_size)
}

/*
Sets inner padding of a table cell.
*/
func (self Go) SetCellPadding(padding gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCellPadding(padding)
}

/*
Adds a [code skip-lint][cell][/code] tag to the tag stack. Must be inside a [code skip-lint][table][/code] tag. See [method push_table] for details. Use [method set_table_column_expand] to set column expansion ratio, [method set_cell_border_color] to set cell border, [method set_cell_row_background_color] to set cell background, [method set_cell_size_override] to override cell size, and [method set_cell_padding] to set padding.
*/
func (self Go) PushCell() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushCell()
}

/*
Adds a [code skip-lint][fgcolor][/code] tag to the tag stack.
*/
func (self Go) PushFgcolor(fgcolor gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushFgcolor(fgcolor)
}

/*
Adds a [code skip-lint][bgcolor][/code] tag to the tag stack.
*/
func (self Go) PushBgcolor(bgcolor gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushBgcolor(bgcolor)
}

/*
Adds a custom effect tag to the tag stack. The effect does not need to be in [member custom_effects]. The environment is directly passed to the effect.
*/
func (self Go) PushCustomfx(effect gdclass.RichTextEffect, env gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushCustomfx(effect, env)
}

/*
Adds a context marker to the tag stack. See [method pop_context].
*/
func (self Go) PushContext() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PushContext()
}

/*
Terminates tags opened after the last [method push_context] call (including context marker), or all tags if there's no context marker on the stack.
*/
func (self Go) PopContext() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PopContext()
}

/*
Terminates the current tag. Use after [code]push_*[/code] methods to close BBCodes manually. Does not need to follow [code]add_*[/code] methods.
*/
func (self Go) Pop() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Pop()
}

/*
Terminates all tags opened by [code]push_*[/code] methods.
*/
func (self Go) PopAll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PopAll()
}

/*
Clears the tag stack, causing the label to display nothing.
[b]Note:[/b] This method does not affect [member text], and its contents will show again if the label is redrawn. However, setting [member text] to an empty [String] also clears the stack.
*/
func (self Go) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Clear()
}

/*
Returns the vertical scrollbar.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Go) GetVScrollBar() gdclass.VScrollBar {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.VScrollBar(class(self).GetVScrollBar(gc))
}

/*
Scrolls the window's top line to match [param line].
*/
func (self Go) ScrollToLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ScrollToLine(gd.Int(line))
}

/*
Scrolls the window's top line to match first line of the [param paragraph].
*/
func (self Go) ScrollToParagraph(paragraph int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ScrollToParagraph(gd.Int(paragraph))
}

/*
Scrolls to the beginning of the current selection.
*/
func (self Go) ScrollToSelection() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ScrollToSelection()
}

/*
Returns the current selection first character index if a selection is active, [code]-1[/code] otherwise. Does not include BBCodes.
*/
func (self Go) GetSelectionFrom() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionFrom()))
}

/*
Returns the current selection last character index if a selection is active, [code]-1[/code] otherwise. Does not include BBCodes.
*/
func (self Go) GetSelectionTo() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSelectionTo()))
}

/*
Select all the text.
If [member selection_enabled] is [code]false[/code], no selection will occur.
*/
func (self Go) SelectAll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SelectAll()
}

/*
Returns the current selection text. Does not include BBCodes.
*/
func (self Go) GetSelectedText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetSelectedText(gc).String())
}

/*
Clears the current selection.
*/
func (self Go) Deselect() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Deselect()
}

/*
The assignment version of [method append_text]. Clears the tag stack and inserts the new content.
*/
func (self Go) ParseBbcode(bbcode string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ParseBbcode(gc.String(bbcode))
}

/*
Parses [param bbcode] and adds tags to the tag stack as needed.
[b]Note:[/b] Using this method, you can't close a tag that was opened in a previous [method append_text] call. This is done to improve performance, especially when updating large RichTextLabels since rebuilding the whole BBCode every time would be slower. If you absolutely need to close a tag in a future method call, append the [member text] instead of using [method append_text].
*/
func (self Go) AppendText(bbcode string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AppendText(gc.String(bbcode))
}

/*
If [member threaded] is enabled, returns [code]true[/code] if the background thread has finished text processing, otherwise always return [code]true[/code].
*/
func (self Go) IsReady() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsReady())
}

/*
Returns the line number of the character position provided. Line and character numbers are both zero-indexed.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetCharacterLine(character int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCharacterLine(gd.Int(character))))
}

/*
Returns the paragraph number of the character position provided. Paragraph and character numbers are both zero-indexed.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetCharacterParagraph(character int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCharacterParagraph(gd.Int(character))))
}

/*
Returns the total number of characters from text tags. Does not include BBCodes.
*/
func (self Go) GetTotalCharacterCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetTotalCharacterCount()))
}

/*
Returns the total number of lines in the text. Wrapped text is counted as multiple lines.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetLineCount()))
}

/*
Returns the number of visible lines.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetVisibleLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetVisibleLineCount()))
}

/*
Returns the total number of paragraphs (newlines or [code]p[/code] tags in the tag stack's text tags). Considers wrapped text as one paragraph.
*/
func (self Go) GetParagraphCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetParagraphCount()))
}

/*
Returns the number of visible paragraphs. A paragraph is considered visible if at least one of its lines is visible.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetVisibleParagraphCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetVisibleParagraphCount()))
}

/*
Returns the height of the content.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetContentHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetContentHeight()))
}

/*
Returns the width of the content.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetContentWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetContentWidth()))
}

/*
Returns the vertical offset of the line found at the provided index.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetLineOffset(line int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetLineOffset(gd.Int(line))))
}

/*
Returns the vertical offset of the paragraph found at the provided index.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
func (self Go) GetParagraphOffset(paragraph int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetParagraphOffset(gd.Int(paragraph))))
}

/*
Parses BBCode parameter [param expressions] into a dictionary.
*/
func (self Go) ParseExpressionsForValues(expressions []string) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).ParseExpressionsForValues(gc, gc.PackedStringSlice(expressions)))
}

/*
Installs a custom effect. This can also be done in the RichTextLabel inspector using the [member custom_effects] property. [param effect] should be a valid [RichTextEffect].
Example RichTextEffect:
[codeblock]
# effect.gd
class_name MyCustomEffect
extends RichTextEffect

var bbcode = "my_custom_effect"

# ...
[/codeblock]
Registering the above effect in RichTextLabel from script:
[codeblock]
# rich_text_label.gd
extends RichTextLabel

func _ready():
    install_effect(MyCustomEffect.new())

    # Alternatively, if not using `class_name` in the script that extends RichTextEffect:
    install_effect(preload("res://effect.gd").new())
[/codeblock]
*/
func (self Go) InstallEffect(effect gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).InstallEffect(effect)
}

/*
Returns the [PopupMenu] of this [RichTextLabel]. By default, this menu is displayed when right-clicking on the [RichTextLabel].
You can add custom menu items or remove standard ones. Make sure your IDs don't conflict with the standard ones (see [enum MenuItems]). For example:
[codeblocks]
[gdscript]
func _ready():
    var menu = get_menu()
    # Remove "Select All" item.
    menu.remove_item(MENU_SELECT_ALL)
    # Add custom items.
    menu.add_separator()
    menu.add_item("Duplicate Text", MENU_MAX + 1)
    # Connect callback.
    menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):
    if id == MENU_MAX + 1:
        add_text("\n" + get_parsed_text())
[/gdscript]
[csharp]
public override void _Ready()
{
    var menu = GetMenu();
    // Remove "Select All" item.
    menu.RemoveItem(RichTextLabel.MenuItems.SelectAll);
    // Add custom items.
    menu.AddSeparator();
    menu.AddItem("Duplicate Text", RichTextLabel.MenuItems.Max + 1);
    // Add event handler.
    menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
    if (id == TextEdit.MenuItems.Max + 1)
    {
        AddText("\n" + GetParsedText());
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

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
func (self Go) MenuOption(option int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MenuOption(gd.Int(option))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RichTextLabel
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RichTextLabel"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BbcodeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsUsingBbcode())
}

func (self Go) SetBbcodeEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetUseBbcode(value)
}

func (self Go) Text() string {
	gc := gd.GarbageCollector(); _ = gc
		return string(class(self).GetText(gc).String())
}

func (self Go) SetText(value string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetText(gc.String(value))
}

func (self Go) FitContent() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsFitContentEnabled())
}

func (self Go) SetFitContent(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFitContent(value)
}

func (self Go) ScrollActive() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsScrollActive())
}

func (self Go) SetScrollActive(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScrollActive(value)
}

func (self Go) ScrollFollowing() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsScrollFollowing())
}

func (self Go) SetScrollFollowing(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetScrollFollow(value)
}

func (self Go) AutowrapMode() classdb.TextServerAutowrapMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Go) SetAutowrapMode(value classdb.TextServerAutowrapMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAutowrapMode(value)
}

func (self Go) TabSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetTabSize()))
}

func (self Go) SetTabSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTabSize(gd.Int(value))
}

func (self Go) ContextMenuEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsContextMenuEnabled())
}

func (self Go) SetContextMenuEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetContextMenuEnabled(value)
}

func (self Go) ShortcutKeysEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsShortcutKeysEnabled())
}

func (self Go) SetShortcutKeysEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShortcutKeysEnabled(value)
}

func (self Go) CustomEffects() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetEffects(gc))
}

func (self Go) SetCustomEffects(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetEffects(value)
}

func (self Go) MetaUnderlined() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMetaUnderlined())
}

func (self Go) SetMetaUnderlined(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMetaUnderline(value)
}

func (self Go) HintUnderlined() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsHintUnderlined())
}

func (self Go) SetHintUnderlined(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHintUnderline(value)
}

func (self Go) Threaded() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsThreaded())
}

func (self Go) SetThreaded(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetThreaded(value)
}

func (self Go) ProgressBarDelay() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetProgressBarDelay()))
}

func (self Go) SetProgressBarDelay(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProgressBarDelay(gd.Int(value))
}

func (self Go) SelectionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsSelectionEnabled())
}

func (self Go) SetSelectionEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSelectionEnabled(value)
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

func (self Go) VisibleCharacters() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetVisibleCharacters()))
}

func (self Go) SetVisibleCharacters(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibleCharacters(gd.Int(value))
}

func (self Go) VisibleCharactersBehavior() classdb.TextServerVisibleCharactersBehavior {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerVisibleCharactersBehavior(class(self).GetVisibleCharactersBehavior())
}

func (self Go) SetVisibleCharactersBehavior(value classdb.TextServerVisibleCharactersBehavior) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibleCharactersBehavior(value)
}

func (self Go) VisibleRatio() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVisibleRatio()))
}

func (self Go) SetVisibleRatio(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVisibleRatio(gd.Float(value))
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

/*
Returns the text without BBCode mark-up.
*/
//go:nosplit
func (self class) GetParsedText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_parsed_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds raw non-BBCode-parsed text to the tag stack.
*/
//go:nosplit
func (self class) AddText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_add_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetText(text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an image's opening and closing tags to the tag stack, optionally providing a [param width] and [param height] to resize the image, a [param color] to tint the image and a [param region] to only use parts of the image.
If [param width] or [param height] is set to 0, the image size will be adjusted in order to keep the original aspect ratio.
If [param width] and [param height] are not set, but [param region] is, the region's rect will be used.
[param key] is an optional identifier, that can be used to modify the image via [method update_image].
If [param pad] is set, and the image is smaller than the size specified by [param width] and [param height], the image padding is added to match the size instead of upscaling.
If [param size_in_percent] is set, [param width] and [param height] values are percentages of the control width instead of pixels.
*/
//go:nosplit
func (self class) AddImage(image gdclass.Texture2D, width gd.Int, height gd.Int, color gd.Color, inline_align gd.InlineAlignment, region gd.Rect2, key gd.Variant, pad bool, tooltip gd.String, size_in_percent bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, color)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, region)
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, pad)
	callframe.Arg(frame, mmm.Get(tooltip))
	callframe.Arg(frame, size_in_percent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_add_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Updates the existing images with the key [param key]. Only properties specified by [param mask] bits are updated. See [method add_image].
*/
//go:nosplit
func (self class) UpdateImage(key gd.Variant, mask classdb.RichTextLabelImageUpdateMask, image gdclass.Texture2D, width gd.Int, height gd.Int, color gd.Color, inline_align gd.InlineAlignment, region gd.Rect2, pad bool, tooltip gd.String, size_in_percent bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, mask)
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, color)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, region)
	callframe.Arg(frame, pad)
	callframe.Arg(frame, mmm.Get(tooltip))
	callframe.Arg(frame, size_in_percent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_update_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a newline tag to the tag stack.
*/
//go:nosplit
func (self class) Newline()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_newline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a paragraph of content from the label. Returns [code]true[/code] if the paragraph exists.
The [param paragraph] argument is the index of the paragraph to remove, it can take values in the interval [code][0, get_paragraph_count() - 1][/code].
If [param no_invalidate] is set to [code]true[/code], cache for the subsequent paragraphs is not invalidated. Use it for faster updates if deleted paragraph is fully self-contained (have no unclosed tags), or this call is part of the complex edit operation and [method invalidate_paragraph] will be called at the end of operation.
*/
//go:nosplit
func (self class) RemoveParagraph(paragraph gd.Int, no_invalidate bool) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, paragraph)
	callframe.Arg(frame, no_invalidate)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_remove_paragraph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Invalidates [param paragraph] and all subsequent paragraphs cache.
*/
//go:nosplit
func (self class) InvalidateParagraph(paragraph gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, paragraph)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_invalidate_paragraph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a [code skip-lint][font][/code] tag to the tag stack. Overrides default fonts for its duration.
Passing [code]0[/code] to [param font_size] will use the existing default font size.
*/
//go:nosplit
func (self class) PushFont(font gdclass.Font, font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][font_size][/code] tag to the tag stack. Overrides default font size for its duration.
*/
//go:nosplit
func (self class) PushFontSize(font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][font][/code] tag with a normal font to the tag stack.
*/
//go:nosplit
func (self class) PushNormal()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][font][/code] tag with a bold font to the tag stack. This is the same as adding a [code skip-lint][b][/code] tag if not currently in a [code skip-lint][i][/code] tag.
*/
//go:nosplit
func (self class) PushBold()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_bold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][font][/code] tag with a bold italics font to the tag stack.
*/
//go:nosplit
func (self class) PushBoldItalics()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_bold_italics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][font][/code] tag with an italics font to the tag stack. This is the same as adding an [code skip-lint][i][/code] tag if not currently in a [code skip-lint][b][/code] tag.
*/
//go:nosplit
func (self class) PushItalics()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_italics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][font][/code] tag with a monospace font to the tag stack.
*/
//go:nosplit
func (self class) PushMono()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_mono, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][color][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][outline_size][/code] tag to the tag stack. Overrides default text outline size for its duration.
*/
//go:nosplit
func (self class) PushOutlineSize(outline_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, outline_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_outline_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][outline_color][/code] tag to the tag stack. Adds text outline for its duration.
*/
//go:nosplit
func (self class) PushOutlineColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_outline_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][p][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushParagraph(alignment gd.HorizontalAlignment, base_direction classdb.ControlTextDirection, language gd.String, st_parser classdb.TextServerStructuredTextParser, justification_flags classdb.TextServerJustificationFlag, tab_stops gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, base_direction)
	callframe.Arg(frame, mmm.Get(language))
	callframe.Arg(frame, st_parser)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, mmm.Get(tab_stops))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_paragraph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an [code skip-lint][indent][/code] tag to the tag stack. Multiplies [param level] by current [member tab_size] to determine new margin length.
*/
//go:nosplit
func (self class) PushIndent(level gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, level)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_indent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds [code skip-lint][ol][/code] or [code skip-lint][ul][/code] tag to the tag stack. Multiplies [param level] by current [member tab_size] to determine new margin length.
*/
//go:nosplit
func (self class) PushList(level gd.Int, atype classdb.RichTextLabelListType, capitalize bool, bullet gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, level)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, capitalize)
	callframe.Arg(frame, mmm.Get(bullet))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a meta tag to the tag stack. Similar to the BBCode [code skip-lint][url=something]{text}[/url][/code], but supports non-[String] metadata types.
If [member meta_underlined] is [code]true[/code], meta tags display an underline. This behavior can be customized with [param underline_mode].
[b]Note:[/b] Meta tags do nothing by default when clicked. To assign behavior when clicked, connect [signal meta_clicked] to a function that is called when the meta tag is clicked.
*/
//go:nosplit
func (self class) PushMeta(data gd.Variant, underline_mode classdb.RichTextLabelMetaUnderline)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	callframe.Arg(frame, underline_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_meta, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][hint][/code] tag to the tag stack. Same as BBCode [code skip-lint][hint=something]{text}[/hint][/code].
*/
//go:nosplit
func (self class) PushHint(description gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(description))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds language code used for text shaping algorithm and Open-Type font features.
*/
//go:nosplit
func (self class) PushLanguage(language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][u][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushUnderline()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_underline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][s][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushStrikethrough()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_strikethrough, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][table=columns,inline_align][/code] tag to the tag stack. Use [method set_table_column_expand] to set column expansion ratio. Use [method push_cell] to add cells.
*/
//go:nosplit
func (self class) PushTable(columns gd.Int, inline_align gd.InlineAlignment, align_to_row gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, columns)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, align_to_row)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_table, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][dropcap][/code] tag to the tag stack. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
//go:nosplit
func (self class) PushDropcap(s gd.String, font gdclass.Font, size gd.Int, dropcap_margins gd.Rect2, color gd.Color, outline_size gd.Int, outline_color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(s))
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	callframe.Arg(frame, size)
	callframe.Arg(frame, dropcap_margins)
	callframe.Arg(frame, color)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, outline_color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_dropcap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Edits the selected column's expansion options. If [param expand] is [code]true[/code], the column expands in proportion to its expansion ratio versus the other columns' ratios.
For example, 2 columns with ratios 3 and 4 plus 70 pixels in available width would expand 30 and 40 pixels, respectively.
If [param expand] is [code]false[/code], the column will not contribute to the total ratio.
*/
//go:nosplit
func (self class) SetTableColumnExpand(column gd.Int, expand bool, ratio gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, expand)
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_table_column_expand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets color of a table cell. Separate colors for alternating rows can be specified.
*/
//go:nosplit
func (self class) SetCellRowBackgroundColor(odd_row_bg gd.Color, even_row_bg gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, odd_row_bg)
	callframe.Arg(frame, even_row_bg)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_cell_row_background_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets color of a table cell border.
*/
//go:nosplit
func (self class) SetCellBorderColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_cell_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets minimum and maximum size overrides for a table cell.
*/
//go:nosplit
func (self class) SetCellSizeOverride(min_size gd.Vector2, max_size gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min_size)
	callframe.Arg(frame, max_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_cell_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets inner padding of a table cell.
*/
//go:nosplit
func (self class) SetCellPadding(padding gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, padding)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_cell_padding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][cell][/code] tag to the tag stack. Must be inside a [code skip-lint][table][/code] tag. See [method push_table] for details. Use [method set_table_column_expand] to set column expansion ratio, [method set_cell_border_color] to set cell border, [method set_cell_row_background_color] to set cell background, [method set_cell_size_override] to override cell size, and [method set_cell_padding] to set padding.
*/
//go:nosplit
func (self class) PushCell()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][fgcolor][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushFgcolor(fgcolor gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fgcolor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_fgcolor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [code skip-lint][bgcolor][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushBgcolor(bgcolor gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bgcolor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_bgcolor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a custom effect tag to the tag stack. The effect does not need to be in [member custom_effects]. The environment is directly passed to the effect.
*/
//go:nosplit
func (self class) PushCustomfx(effect gdclass.RichTextEffect, env gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(effect[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(env))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_customfx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a context marker to the tag stack. See [method pop_context].
*/
//go:nosplit
func (self class) PushContext()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_push_context, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Terminates tags opened after the last [method push_context] call (including context marker), or all tags if there's no context marker on the stack.
*/
//go:nosplit
func (self class) PopContext()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_pop_context, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Terminates the current tag. Use after [code]push_*[/code] methods to close BBCodes manually. Does not need to follow [code]add_*[/code] methods.
*/
//go:nosplit
func (self class) Pop()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_pop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Terminates all tags opened by [code]push_*[/code] methods.
*/
//go:nosplit
func (self class) PopAll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_pop_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears the tag stack, causing the label to display nothing.
[b]Note:[/b] This method does not affect [member text], and its contents will show again if the label is redrawn. However, setting [member text] to an empty [String] also clears the stack.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser classdb.TextServerStructuredTextParser)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverride() classdb.TextServerStructuredTextParser {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerStructuredTextParser](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextDirection(direction classdb.ControlTextDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextDirection() classdb.ControlTextDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlTextDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLanguage(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode classdb.TextServerAutowrapMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutowrapMode() classdb.TextServerAutowrapMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerAutowrapMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMetaUnderline(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_meta_underline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMetaUnderlined() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_meta_underlined, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHintUnderline(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_hint_underline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHintUnderlined() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_hint_underlined, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScrollActive(active bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_scroll_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsScrollActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_scroll_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScrollFollow(follow bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, follow)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_scroll_follow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsScrollFollowing() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_scroll_following, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the vertical scrollbar.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetVScrollBar(ctx gd.Lifetime) gdclass.VScrollBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.VScrollBar
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Scrolls the window's top line to match [param line].
*/
//go:nosplit
func (self class) ScrollToLine(line gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_scroll_to_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Scrolls the window's top line to match first line of the [param paragraph].
*/
//go:nosplit
func (self class) ScrollToParagraph(paragraph gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, paragraph)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_scroll_to_paragraph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Scrolls to the beginning of the current selection.
*/
//go:nosplit
func (self class) ScrollToSelection()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_scroll_to_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTabSize(spaces gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, spaces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_tab_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTabSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_tab_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFitContent(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_fit_content, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFitContentEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_fit_content_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSelectionEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSelectionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContextMenuEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsContextMenuEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShortcutKeysEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsShortcutKeysEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDeselectOnFocusLossEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDragAndDropSelectionEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current selection first character index if a selection is active, [code]-1[/code] otherwise. Does not include BBCodes.
*/
//go:nosplit
func (self class) GetSelectionFrom() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_selection_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current selection last character index if a selection is active, [code]-1[/code] otherwise. Does not include BBCodes.
*/
//go:nosplit
func (self class) GetSelectionTo() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_selection_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Select all the text.
If [member selection_enabled] is [code]false[/code], no selection will occur.
*/
//go:nosplit
func (self class) SelectAll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_select_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current selection text. Does not include BBCodes.
*/
//go:nosplit
func (self class) GetSelectedText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_selected_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears the current selection.
*/
//go:nosplit
func (self class) Deselect()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
The assignment version of [method append_text]. Clears the tag stack and inserts the new content.
*/
//go:nosplit
func (self class) ParseBbcode(bbcode gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bbcode))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_parse_bbcode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Parses [param bbcode] and adds tags to the tag stack as needed.
[b]Note:[/b] Using this method, you can't close a tag that was opened in a previous [method append_text] call. This is done to improve performance, especially when updating large RichTextLabels since rebuilding the whole BBCode every time would be slower. If you absolutely need to close a tag in a future method call, append the [member text] instead of using [method append_text].
*/
//go:nosplit
func (self class) AppendText(bbcode gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(bbcode))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_append_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
If [member threaded] is enabled, returns [code]true[/code] if the background thread has finished text processing, otherwise always return [code]true[/code].
*/
//go:nosplit
func (self class) IsReady() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_ready, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetThreaded(threaded bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, threaded)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_threaded, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsThreaded() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_threaded, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetProgressBarDelay(delay_ms gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, delay_ms)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_progress_bar_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetProgressBarDelay() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_progress_bar_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibleCharacters(amount gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_visible_characters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibleCharacters() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_visible_characters, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetVisibleCharactersBehavior() classdb.TextServerVisibleCharactersBehavior {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerVisibleCharactersBehavior](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_visible_characters_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibleCharactersBehavior(behavior classdb.TextServerVisibleCharactersBehavior)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, behavior)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_visible_characters_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVisibleRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_visible_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibleRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_visible_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the line number of the character position provided. Line and character numbers are both zero-indexed.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetCharacterLine(character gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, character)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_character_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the paragraph number of the character position provided. Paragraph and character numbers are both zero-indexed.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetCharacterParagraph(character gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, character)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_character_paragraph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total number of characters from text tags. Does not include BBCodes.
*/
//go:nosplit
func (self class) GetTotalCharacterCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_total_character_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseBbcode(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_use_bbcode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingBbcode() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_using_bbcode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total number of lines in the text. Wrapped text is counted as multiple lines.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of visible lines.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetVisibleLineCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_visible_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the total number of paragraphs (newlines or [code]p[/code] tags in the tag stack's text tags). Considers wrapped text as one paragraph.
*/
//go:nosplit
func (self class) GetParagraphCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_paragraph_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of visible paragraphs. A paragraph is considered visible if at least one of its lines is visible.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetVisibleParagraphCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_visible_paragraph_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the height of the content.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetContentHeight() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_content_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the width of the content.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetContentWidth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_content_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the vertical offset of the line found at the provided index.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetLineOffset(line gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_line_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the vertical offset of the paragraph found at the provided index.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_ready] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetParagraphOffset(paragraph gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, paragraph)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_paragraph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Parses BBCode parameter [param expressions] into a dictionary.
*/
//go:nosplit
func (self class) ParseExpressionsForValues(ctx gd.Lifetime, expressions gd.PackedStringArray) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(expressions))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_parse_expressions_for_values, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEffects(effects gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(effects))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_set_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEffects(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_effects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Installs a custom effect. This can also be done in the RichTextLabel inspector using the [member custom_effects] property. [param effect] should be a valid [RichTextEffect].
Example RichTextEffect:
[codeblock]
# effect.gd
class_name MyCustomEffect
extends RichTextEffect

var bbcode = "my_custom_effect"

# ...
[/codeblock]
Registering the above effect in RichTextLabel from script:
[codeblock]
# rich_text_label.gd
extends RichTextLabel

func _ready():
    install_effect(MyCustomEffect.new())

    # Alternatively, if not using `class_name` in the script that extends RichTextEffect:
    install_effect(preload("res://effect.gd").new())
[/codeblock]
*/
//go:nosplit
func (self class) InstallEffect(effect gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(effect))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_install_effect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [PopupMenu] of this [RichTextLabel]. By default, this menu is displayed when right-clicking on the [RichTextLabel].
You can add custom menu items or remove standard ones. Make sure your IDs don't conflict with the standard ones (see [enum MenuItems]). For example:
[codeblocks]
[gdscript]
func _ready():
    var menu = get_menu()
    # Remove "Select All" item.
    menu.remove_item(MENU_SELECT_ALL)
    # Add custom items.
    menu.add_separator()
    menu.add_item("Duplicate Text", MENU_MAX + 1)
    # Connect callback.
    menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):
    if id == MENU_MAX + 1:
        add_text("\n" + get_parsed_text())
[/gdscript]
[csharp]
public override void _Ready()
{
    var menu = GetMenu();
    // Remove "Select All" item.
    menu.RemoveItem(RichTextLabel.MenuItems.SelectAll);
    // Add custom items.
    menu.AddSeparator();
    menu.AddItem("Duplicate Text", RichTextLabel.MenuItems.Max + 1);
    // Add event handler.
    menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
    if (id == TextEdit.MenuItems.Max + 1)
    {
        AddText("\n" + GetParsedText());
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_is_menu_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_menu_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnMetaClicked(cb func(meta gd.Variant)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("meta_clicked"), gc.Callable(cb), 0)
}


func (self Go) OnMetaHoverStarted(cb func(meta gd.Variant)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("meta_hover_started"), gc.Callable(cb), 0)
}


func (self Go) OnMetaHoverEnded(cb func(meta gd.Variant)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("meta_hover_ended"), gc.Callable(cb), 0)
}


func (self Go) OnFinished(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("finished"), gc.Callable(cb), 0)
}


func (self class) AsRichTextLabel() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRichTextLabel() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("RichTextLabel", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type ListType = classdb.RichTextLabelListType

const (
/*Each list item has a number marker.*/
	ListNumbers ListType = 0
/*Each list item has a letter marker.*/
	ListLetters ListType = 1
/*Each list item has a roman number marker.*/
	ListRoman ListType = 2
/*Each list item has a filled circle marker.*/
	ListDots ListType = 3
)
type MenuItems = classdb.RichTextLabelMenuItems

const (
/*Copies the selected text.*/
	MenuCopy MenuItems = 0
/*Selects the whole [RichTextLabel] text.*/
	MenuSelectAll MenuItems = 1
/*Represents the size of the [enum MenuItems] enum.*/
	MenuMax MenuItems = 2
)
type MetaUnderline = classdb.RichTextLabelMetaUnderline

const (
/*Meta tag does not display an underline, even if [member meta_underlined] is [code]true[/code].*/
	MetaUnderlineNever MetaUnderline = 0
/*If [member meta_underlined] is [code]true[/code], meta tag always display an underline.*/
	MetaUnderlineAlways MetaUnderline = 1
/*If [member meta_underlined] is [code]true[/code], meta tag display an underline when the mouse cursor is over it.*/
	MetaUnderlineOnHover MetaUnderline = 2
)
type ImageUpdateMask = classdb.RichTextLabelImageUpdateMask

const (
/*If this bit is set, [method update_image] changes image texture.*/
	UpdateTexture ImageUpdateMask = 1
/*If this bit is set, [method update_image] changes image size.*/
	UpdateSize ImageUpdateMask = 2
/*If this bit is set, [method update_image] changes image color.*/
	UpdateColor ImageUpdateMask = 4
/*If this bit is set, [method update_image] changes image inline alignment.*/
	UpdateAlignment ImageUpdateMask = 8
/*If this bit is set, [method update_image] changes image texture region.*/
	UpdateRegion ImageUpdateMask = 16
/*If this bit is set, [method update_image] changes image padding.*/
	UpdatePad ImageUpdateMask = 32
/*If this bit is set, [method update_image] changes image tooltip.*/
	UpdateTooltip ImageUpdateMask = 64
/*If this bit is set, [method update_image] changes image width from/to percents.*/
	UpdateWidthInPercent ImageUpdateMask = 128
)
