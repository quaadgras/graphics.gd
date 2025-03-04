// Package RichTextLabel provides methods for working with RichTextLabel object instances.
package RichTextLabel

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"

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
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A control for displaying text that can contain custom fonts, images, and basic formatting. [RichTextLabel] manages these as an internal tag stack. It also adapts itself to given width/heights.
[b]Note:[/b] [method newline], [method push_paragraph], [code]"\n"[/code], [code]"\r\n"[/code], [code]p[/code] tag, and alignment tags start a new paragraph. Each paragraph is processed independently, in its own BiDi context. If you want to force line wrapping within paragraph, any other line breaking character can be used, for example, Form Feed (U+000C), Next Line (U+0085), Line Separator (U+2028).
[b]Note:[/b] Assignments to [member text] clear the tag stack and reconstruct it from the property's contents. Any edits made to [member text] will erase previous edits made from other manual sources such as [method append_text] and the [code]push_*[/code] / [method pop] methods.
[b]Note:[/b] RichTextLabel doesn't support entangled BBCode tags. For example, instead of using [code skip-lint][b]bold[i]bold italic[/b]italic[/i][/code], use [code skip-lint][b]bold[i]bold italic[/i][/b][i]italic[/i][/code].
[b]Note:[/b] [code]push_pop_*[/code] functions won't affect BBCode.
[b]Note:[/b] Unlike [Label], [RichTextLabel] doesn't have a [i]property[/i] to horizontally align text to the center. Instead, enable [member bbcode_enabled] and surround the text in a [code skip-lint][center][/code] tag as follows: [code skip-lint][center]Example[/center][/code]. There is currently no built-in way to vertically align text either, but this can be emulated by relying on anchors/containers and the [member fit_content] property.
*/
type Instance [1]gdclass.RichTextLabel

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRichTextLabel() Instance
}

/*
Returns the text without BBCode mark-up.
*/
func (self Instance) GetParsedText() string { //gd:RichTextLabel.get_parsed_text
	return string(class(self).GetParsedText().String())
}

/*
Adds raw non-BBCode-parsed text to the tag stack.
*/
func (self Instance) AddText(text string) { //gd:RichTextLabel.add_text
	class(self).AddText(String.New(text))
}

/*
Adds an image's opening and closing tags to the tag stack, optionally providing a [param width] and [param height] to resize the image, a [param color] to tint the image and a [param region] to only use parts of the image.
If [param width] or [param height] is set to 0, the image size will be adjusted in order to keep the original aspect ratio.
If [param width] and [param height] are not set, but [param region] is, the region's rect will be used.
[param key] is an optional identifier, that can be used to modify the image via [method update_image].
If [param pad] is set, and the image is smaller than the size specified by [param width] and [param height], the image padding is added to match the size instead of upscaling.
If [param size_in_percent] is set, [param width] and [param height] values are percentages of the control width instead of pixels.
*/
func (self Instance) AddImage(image [1]gdclass.Texture2D) { //gd:RichTextLabel.add_image
	class(self).AddImage(image, int64(0), int64(0), Color.RGBA(gd.Color{1, 1, 1, 1}), 5, Rect2.PositionSize(gd.NewRect2(0, 0, 0, 0)), variant.New([1]any{}[0]), false, String.New(""), false)
}

/*
Updates the existing images with the key [param key]. Only properties specified by [param mask] bits are updated. See [method add_image].
*/
func (self Instance) UpdateImage(key any, mask gdclass.RichTextLabelImageUpdateMask, image [1]gdclass.Texture2D) { //gd:RichTextLabel.update_image
	class(self).UpdateImage(variant.New(key), mask, image, int64(0), int64(0), Color.RGBA(gd.Color{1, 1, 1, 1}), 5, Rect2.PositionSize(gd.NewRect2(0, 0, 0, 0)), false, String.New(""), false)
}

/*
Adds a newline tag to the tag stack.
*/
func (self Instance) Newline() { //gd:RichTextLabel.newline
	class(self).Newline()
}

/*
Removes a paragraph of content from the label. Returns [code]true[/code] if the paragraph exists.
The [param paragraph] argument is the index of the paragraph to remove, it can take values in the interval [code][0, get_paragraph_count() - 1][/code].
If [param no_invalidate] is set to [code]true[/code], cache for the subsequent paragraphs is not invalidated. Use it for faster updates if deleted paragraph is fully self-contained (have no unclosed tags), or this call is part of the complex edit operation and [method invalidate_paragraph] will be called at the end of operation.
*/
func (self Instance) RemoveParagraph(paragraph int) bool { //gd:RichTextLabel.remove_paragraph
	return bool(class(self).RemoveParagraph(int64(paragraph), false))
}

/*
Invalidates [param paragraph] and all subsequent paragraphs cache.
*/
func (self Instance) InvalidateParagraph(paragraph int) bool { //gd:RichTextLabel.invalidate_paragraph
	return bool(class(self).InvalidateParagraph(int64(paragraph)))
}

/*
Adds a [code skip-lint][font][/code] tag to the tag stack. Overrides default fonts for its duration.
Passing [code]0[/code] to [param font_size] will use the existing default font size.
*/
func (self Instance) PushFont(font [1]gdclass.Font) { //gd:RichTextLabel.push_font
	class(self).PushFont(font, int64(0))
}

/*
Adds a [code skip-lint][font_size][/code] tag to the tag stack. Overrides default font size for its duration.
*/
func (self Instance) PushFontSize(font_size int) { //gd:RichTextLabel.push_font_size
	class(self).PushFontSize(int64(font_size))
}

/*
Adds a [code skip-lint][font][/code] tag with a normal font to the tag stack.
*/
func (self Instance) PushNormal() { //gd:RichTextLabel.push_normal
	class(self).PushNormal()
}

/*
Adds a [code skip-lint][font][/code] tag with a bold font to the tag stack. This is the same as adding a [code skip-lint][b][/code] tag if not currently in a [code skip-lint][i][/code] tag.
*/
func (self Instance) PushBold() { //gd:RichTextLabel.push_bold
	class(self).PushBold()
}

/*
Adds a [code skip-lint][font][/code] tag with a bold italics font to the tag stack.
*/
func (self Instance) PushBoldItalics() { //gd:RichTextLabel.push_bold_italics
	class(self).PushBoldItalics()
}

/*
Adds a [code skip-lint][font][/code] tag with an italics font to the tag stack. This is the same as adding an [code skip-lint][i][/code] tag if not currently in a [code skip-lint][b][/code] tag.
*/
func (self Instance) PushItalics() { //gd:RichTextLabel.push_italics
	class(self).PushItalics()
}

/*
Adds a [code skip-lint][font][/code] tag with a monospace font to the tag stack.
*/
func (self Instance) PushMono() { //gd:RichTextLabel.push_mono
	class(self).PushMono()
}

/*
Adds a [code skip-lint][color][/code] tag to the tag stack.
*/
func (self Instance) PushColor(color Color.RGBA) { //gd:RichTextLabel.push_color
	class(self).PushColor(Color.RGBA(color))
}

/*
Adds a [code skip-lint][outline_size][/code] tag to the tag stack. Overrides default text outline size for its duration.
*/
func (self Instance) PushOutlineSize(outline_size int) { //gd:RichTextLabel.push_outline_size
	class(self).PushOutlineSize(int64(outline_size))
}

/*
Adds a [code skip-lint][outline_color][/code] tag to the tag stack. Adds text outline for its duration.
*/
func (self Instance) PushOutlineColor(color Color.RGBA) { //gd:RichTextLabel.push_outline_color
	class(self).PushOutlineColor(Color.RGBA(color))
}

/*
Adds a [code skip-lint][p][/code] tag to the tag stack.
*/
func (self Instance) PushParagraph(alignment HorizontalAlignment) { //gd:RichTextLabel.push_paragraph
	class(self).PushParagraph(alignment, 0, String.New(""), 0, 163, Packed.New([1][]float32{}[0]...))
}

/*
Adds an [code skip-lint][indent][/code] tag to the tag stack. Multiplies [param level] by current [member tab_size] to determine new margin length.
*/
func (self Instance) PushIndent(level int) { //gd:RichTextLabel.push_indent
	class(self).PushIndent(int64(level))
}

/*
Adds [code skip-lint][ol][/code] or [code skip-lint][ul][/code] tag to the tag stack. Multiplies [param level] by current [member tab_size] to determine new margin length.
*/
func (self Instance) PushList(level int, atype gdclass.RichTextLabelListType, capitalize bool) { //gd:RichTextLabel.push_list
	class(self).PushList(int64(level), atype, capitalize, String.New("•"))
}

/*
Adds a meta tag to the tag stack. Similar to the BBCode [code skip-lint][url=something]{text}[/url][/code], but supports non-[String] metadata types.
If [member meta_underlined] is [code]true[/code], meta tags display an underline. This behavior can be customized with [param underline_mode].
[b]Note:[/b] Meta tags do nothing by default when clicked. To assign behavior when clicked, connect [signal meta_clicked] to a function that is called when the meta tag is clicked.
*/
func (self Instance) PushMeta(data any) { //gd:RichTextLabel.push_meta
	class(self).PushMeta(variant.New(data), 1, String.New(""))
}

/*
Adds a [code skip-lint][hint][/code] tag to the tag stack. Same as BBCode [code skip-lint][hint=something]{text}[/hint][/code].
*/
func (self Instance) PushHint(description string) { //gd:RichTextLabel.push_hint
	class(self).PushHint(String.New(description))
}

/*
Adds language code used for text shaping algorithm and Open-Type font features.
*/
func (self Instance) PushLanguage(language string) { //gd:RichTextLabel.push_language
	class(self).PushLanguage(String.New(language))
}

/*
Adds a [code skip-lint][u][/code] tag to the tag stack.
*/
func (self Instance) PushUnderline() { //gd:RichTextLabel.push_underline
	class(self).PushUnderline()
}

/*
Adds a [code skip-lint][s][/code] tag to the tag stack.
*/
func (self Instance) PushStrikethrough() { //gd:RichTextLabel.push_strikethrough
	class(self).PushStrikethrough()
}

/*
Adds a [code skip-lint][table=columns,inline_align][/code] tag to the tag stack. Use [method set_table_column_expand] to set column expansion ratio. Use [method push_cell] to add cells.
*/
func (self Instance) PushTable(columns int) { //gd:RichTextLabel.push_table
	class(self).PushTable(int64(columns), 0, int64(-1))
}

/*
Adds a [code skip-lint][dropcap][/code] tag to the tag stack. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
func (self Instance) PushDropcap(s string, font [1]gdclass.Font, size int) { //gd:RichTextLabel.push_dropcap
	class(self).PushDropcap(String.New(s), font, int64(size), Rect2.PositionSize(gd.NewRect2(0, 0, 0, 0)), Color.RGBA(gd.Color{1, 1, 1, 1}), int64(0), Color.RGBA(gd.Color{0, 0, 0, 0}))
}

/*
Edits the selected column's expansion options. If [param expand] is [code]true[/code], the column expands in proportion to its expansion ratio versus the other columns' ratios.
For example, 2 columns with ratios 3 and 4 plus 70 pixels in available width would expand 30 and 40 pixels, respectively.
If [param expand] is [code]false[/code], the column will not contribute to the total ratio.
*/
func (self Instance) SetTableColumnExpand(column int, expand bool) { //gd:RichTextLabel.set_table_column_expand
	class(self).SetTableColumnExpand(int64(column), expand, int64(1), true)
}

/*
Sets color of a table cell. Separate colors for alternating rows can be specified.
*/
func (self Instance) SetCellRowBackgroundColor(odd_row_bg Color.RGBA, even_row_bg Color.RGBA) { //gd:RichTextLabel.set_cell_row_background_color
	class(self).SetCellRowBackgroundColor(Color.RGBA(odd_row_bg), Color.RGBA(even_row_bg))
}

/*
Sets color of a table cell border.
*/
func (self Instance) SetCellBorderColor(color Color.RGBA) { //gd:RichTextLabel.set_cell_border_color
	class(self).SetCellBorderColor(Color.RGBA(color))
}

/*
Sets minimum and maximum size overrides for a table cell.
*/
func (self Instance) SetCellSizeOverride(min_size Vector2.XY, max_size Vector2.XY) { //gd:RichTextLabel.set_cell_size_override
	class(self).SetCellSizeOverride(Vector2.XY(min_size), Vector2.XY(max_size))
}

/*
Sets inner padding of a table cell.
*/
func (self Instance) SetCellPadding(padding Rect2.PositionSize) { //gd:RichTextLabel.set_cell_padding
	class(self).SetCellPadding(Rect2.PositionSize(padding))
}

/*
Adds a [code skip-lint][cell][/code] tag to the tag stack. Must be inside a [code skip-lint][table][/code] tag. See [method push_table] for details. Use [method set_table_column_expand] to set column expansion ratio, [method set_cell_border_color] to set cell border, [method set_cell_row_background_color] to set cell background, [method set_cell_size_override] to override cell size, and [method set_cell_padding] to set padding.
*/
func (self Instance) PushCell() { //gd:RichTextLabel.push_cell
	class(self).PushCell()
}

/*
Adds a [code skip-lint][fgcolor][/code] tag to the tag stack.
*/
func (self Instance) PushFgcolor(fgcolor Color.RGBA) { //gd:RichTextLabel.push_fgcolor
	class(self).PushFgcolor(Color.RGBA(fgcolor))
}

/*
Adds a [code skip-lint][bgcolor][/code] tag to the tag stack.
*/
func (self Instance) PushBgcolor(bgcolor Color.RGBA) { //gd:RichTextLabel.push_bgcolor
	class(self).PushBgcolor(Color.RGBA(bgcolor))
}

/*
Adds a custom effect tag to the tag stack. The effect does not need to be in [member custom_effects]. The environment is directly passed to the effect.
*/
func (self Instance) PushCustomfx(effect [1]gdclass.RichTextEffect, env map[string]interface{}) { //gd:RichTextLabel.push_customfx
	class(self).PushCustomfx(effect, gd.DictionaryFromMap(env))
}

/*
Adds a context marker to the tag stack. See [method pop_context].
*/
func (self Instance) PushContext() { //gd:RichTextLabel.push_context
	class(self).PushContext()
}

/*
Terminates tags opened after the last [method push_context] call (including context marker), or all tags if there's no context marker on the stack.
*/
func (self Instance) PopContext() { //gd:RichTextLabel.pop_context
	class(self).PopContext()
}

/*
Terminates the current tag. Use after [code]push_*[/code] methods to close BBCodes manually. Does not need to follow [code]add_*[/code] methods.
*/
func (self Instance) Pop() { //gd:RichTextLabel.pop
	class(self).Pop()
}

/*
Terminates all tags opened by [code]push_*[/code] methods.
*/
func (self Instance) PopAll() { //gd:RichTextLabel.pop_all
	class(self).PopAll()
}

/*
Clears the tag stack, causing the label to display nothing.
[b]Note:[/b] This method does not affect [member text], and its contents will show again if the label is redrawn. However, setting [member text] to an empty [String] also clears the stack.
*/
func (self Instance) Clear() { //gd:RichTextLabel.clear
	class(self).Clear()
}

/*
Returns the vertical scrollbar.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetVScrollBar() [1]gdclass.VScrollBar { //gd:RichTextLabel.get_v_scroll_bar
	return [1]gdclass.VScrollBar(class(self).GetVScrollBar())
}

/*
Scrolls the window's top line to match [param line].
*/
func (self Instance) ScrollToLine(line int) { //gd:RichTextLabel.scroll_to_line
	class(self).ScrollToLine(int64(line))
}

/*
Scrolls the window's top line to match first line of the [param paragraph].
*/
func (self Instance) ScrollToParagraph(paragraph int) { //gd:RichTextLabel.scroll_to_paragraph
	class(self).ScrollToParagraph(int64(paragraph))
}

/*
Scrolls to the beginning of the current selection.
*/
func (self Instance) ScrollToSelection() { //gd:RichTextLabel.scroll_to_selection
	class(self).ScrollToSelection()
}

/*
Returns the current selection first character index if a selection is active, [code]-1[/code] otherwise. Does not include BBCodes.
*/
func (self Instance) GetSelectionFrom() int { //gd:RichTextLabel.get_selection_from
	return int(int(class(self).GetSelectionFrom()))
}

/*
Returns the current selection last character index if a selection is active, [code]-1[/code] otherwise. Does not include BBCodes.
*/
func (self Instance) GetSelectionTo() int { //gd:RichTextLabel.get_selection_to
	return int(int(class(self).GetSelectionTo()))
}

/*
Returns the current selection vertical line offset if a selection is active, [code]-1.0[/code] otherwise.
*/
func (self Instance) GetSelectionLineOffset() Float.X { //gd:RichTextLabel.get_selection_line_offset
	return Float.X(Float.X(class(self).GetSelectionLineOffset()))
}

/*
Select all the text.
If [member selection_enabled] is [code]false[/code], no selection will occur.
*/
func (self Instance) SelectAll() { //gd:RichTextLabel.select_all
	class(self).SelectAll()
}

/*
Returns the current selection text. Does not include BBCodes.
*/
func (self Instance) GetSelectedText() string { //gd:RichTextLabel.get_selected_text
	return string(class(self).GetSelectedText().String())
}

/*
Clears the current selection.
*/
func (self Instance) Deselect() { //gd:RichTextLabel.deselect
	class(self).Deselect()
}

/*
The assignment version of [method append_text]. Clears the tag stack and inserts the new content.
*/
func (self Instance) ParseBbcode(bbcode string) { //gd:RichTextLabel.parse_bbcode
	class(self).ParseBbcode(String.New(bbcode))
}

/*
Parses [param bbcode] and adds tags to the tag stack as needed.
[b]Note:[/b] Using this method, you can't close a tag that was opened in a previous [method append_text] call. This is done to improve performance, especially when updating large RichTextLabels since rebuilding the whole BBCode every time would be slower. If you absolutely need to close a tag in a future method call, append the [member text] instead of using [method append_text].
*/
func (self Instance) AppendText(bbcode string) { //gd:RichTextLabel.append_text
	class(self).AppendText(String.New(bbcode))
}

/*
If [member threaded] is enabled, returns [code]true[/code] if the background thread has finished text processing, otherwise always return [code]true[/code].
*/
func (self Instance) IsReady() bool { //gd:RichTextLabel.is_ready
	return bool(class(self).IsReady())
}

/*
If [member threaded] is enabled, returns [code]true[/code] if the background thread has finished text processing, otherwise always return [code]true[/code].
*/
func (self Instance) IsFinished() bool { //gd:RichTextLabel.is_finished
	return bool(class(self).IsFinished())
}

/*
Returns the line number of the character position provided. Line and character numbers are both zero-indexed.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetCharacterLine(character int) int { //gd:RichTextLabel.get_character_line
	return int(int(class(self).GetCharacterLine(int64(character))))
}

/*
Returns the paragraph number of the character position provided. Paragraph and character numbers are both zero-indexed.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetCharacterParagraph(character int) int { //gd:RichTextLabel.get_character_paragraph
	return int(int(class(self).GetCharacterParagraph(int64(character))))
}

/*
Returns the total number of characters from text tags. Does not include BBCodes.
*/
func (self Instance) GetTotalCharacterCount() int { //gd:RichTextLabel.get_total_character_count
	return int(int(class(self).GetTotalCharacterCount()))
}

/*
Returns the total number of lines in the text. Wrapped text is counted as multiple lines.
[b]Note:[/b] If [member visible_characters_behavior] is set to [constant TextServer.VC_CHARS_BEFORE_SHAPING] only visible wrapped lines are counted.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetLineCount() int { //gd:RichTextLabel.get_line_count
	return int(int(class(self).GetLineCount()))
}

/*
Returns the indexes of the first and last visible characters for the given [param line], as a [Vector2i].
[b]Note:[/b] If [member visible_characters_behavior] is set to [constant TextServer.VC_CHARS_BEFORE_SHAPING] only visible wrapped lines are counted.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetLineRange(line int) Vector2i.XY { //gd:RichTextLabel.get_line_range
	return Vector2i.XY(class(self).GetLineRange(int64(line)))
}

/*
Returns the number of visible lines.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetVisibleLineCount() int { //gd:RichTextLabel.get_visible_line_count
	return int(int(class(self).GetVisibleLineCount()))
}

/*
Returns the total number of paragraphs (newlines or [code]p[/code] tags in the tag stack's text tags). Considers wrapped text as one paragraph.
*/
func (self Instance) GetParagraphCount() int { //gd:RichTextLabel.get_paragraph_count
	return int(int(class(self).GetParagraphCount()))
}

/*
Returns the number of visible paragraphs. A paragraph is considered visible if at least one of its lines is visible.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetVisibleParagraphCount() int { //gd:RichTextLabel.get_visible_paragraph_count
	return int(int(class(self).GetVisibleParagraphCount()))
}

/*
Returns the height of the content.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetContentHeight() int { //gd:RichTextLabel.get_content_height
	return int(int(class(self).GetContentHeight()))
}

/*
Returns the width of the content.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetContentWidth() int { //gd:RichTextLabel.get_content_width
	return int(int(class(self).GetContentWidth()))
}

/*
Returns the vertical offset of the line found at the provided index.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetLineOffset(line int) Float.X { //gd:RichTextLabel.get_line_offset
	return Float.X(Float.X(class(self).GetLineOffset(int64(line))))
}

/*
Returns the vertical offset of the paragraph found at the provided index.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
func (self Instance) GetParagraphOffset(paragraph int) Float.X { //gd:RichTextLabel.get_paragraph_offset
	return Float.X(Float.X(class(self).GetParagraphOffset(int64(paragraph))))
}

/*
Parses BBCode parameter [param expressions] into a dictionary.
*/
func (self Instance) ParseExpressionsForValues(expressions []string) map[string]interface{} { //gd:RichTextLabel.parse_expressions_for_values
	return map[string]interface{}(gd.DictionaryAs[map[string]interface{}](class(self).ParseExpressionsForValues(Packed.MakeStrings(expressions...))))
}

/*
Installs a custom effect. This can also be done in the Inspector through the [member custom_effects] property. [param effect] should be a valid [RichTextEffect].
[b]Example:[/b] With the following script extending from [RichTextEffect]:
[codeblock]
# effect.gd
class_name MyCustomEffect
extends RichTextEffect

var bbcode = "my_custom_effect"

# ...
[/codeblock]
The above effect can be installed in [RichTextLabel] from a script:
[codeblock]
# rich_text_label.gd
extends RichTextLabel

func _ready():

	install_effect(MyCustomEffect.new())

	# Alternatively, if not using `class_name` in the script that extends RichTextEffect:
	install_effect(preload("res://effect.gd").new())

[/codeblock]
*/
func (self Instance) InstallEffect(effect any) { //gd:RichTextLabel.install_effect
	class(self).InstallEffect(variant.New(effect))
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
func (self Instance) GetMenu() [1]gdclass.PopupMenu { //gd:RichTextLabel.get_menu
	return [1]gdclass.PopupMenu(class(self).GetMenu())
}

/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
func (self Instance) IsMenuVisible() bool { //gd:RichTextLabel.is_menu_visible
	return bool(class(self).IsMenuVisible())
}

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
func (self Instance) MenuOption(option int) { //gd:RichTextLabel.menu_option
	class(self).MenuOption(int64(option))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RichTextLabel

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RichTextLabel"))
	casted := Instance{*(*gdclass.RichTextLabel)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) BbcodeEnabled() bool {
	return bool(class(self).IsUsingBbcode())
}

func (self Instance) SetBbcodeEnabled(value bool) {
	class(self).SetUseBbcode(value)
}

func (self Instance) Text() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetText(value string) {
	class(self).SetText(String.New(value))
}

func (self Instance) FitContent() bool {
	return bool(class(self).IsFitContentEnabled())
}

func (self Instance) SetFitContent(value bool) {
	class(self).SetFitContent(value)
}

func (self Instance) ScrollActive() bool {
	return bool(class(self).IsScrollActive())
}

func (self Instance) SetScrollActive(value bool) {
	class(self).SetScrollActive(value)
}

func (self Instance) ScrollFollowing() bool {
	return bool(class(self).IsScrollFollowing())
}

func (self Instance) SetScrollFollowing(value bool) {
	class(self).SetScrollFollow(value)
}

func (self Instance) AutowrapMode() gdclass.TextServerAutowrapMode {
	return gdclass.TextServerAutowrapMode(class(self).GetAutowrapMode())
}

func (self Instance) SetAutowrapMode(value gdclass.TextServerAutowrapMode) {
	class(self).SetAutowrapMode(value)
}

func (self Instance) TabSize() int {
	return int(int(class(self).GetTabSize()))
}

func (self Instance) SetTabSize(value int) {
	class(self).SetTabSize(int64(value))
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

func (self Instance) HorizontalAlignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Instance) SetHorizontalAlignment(value HorizontalAlignment) {
	class(self).SetHorizontalAlignment(value)
}

func (self Instance) VerticalAlignment() VerticalAlignment {
	return VerticalAlignment(class(self).GetVerticalAlignment())
}

func (self Instance) SetVerticalAlignment(value VerticalAlignment) {
	class(self).SetVerticalAlignment(value)
}

func (self Instance) JustificationFlags() gdclass.TextServerJustificationFlag {
	return gdclass.TextServerJustificationFlag(class(self).GetJustificationFlags())
}

func (self Instance) SetJustificationFlags(value gdclass.TextServerJustificationFlag) {
	class(self).SetJustificationFlags(value)
}

func (self Instance) TabStops() []float32 {
	return []float32(slices.Collect(class(self).GetTabStops().Values()))
}

func (self Instance) SetTabStops(value []float32) {
	class(self).SetTabStops(Packed.New(value...))
}

func (self Instance) CustomEffects() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetEffects())))
}

func (self Instance) SetCustomEffects(value []any) {
	class(self).SetEffects(gd.EngineArrayFromSlice(value))
}

func (self Instance) MetaUnderlined() bool {
	return bool(class(self).IsMetaUnderlined())
}

func (self Instance) SetMetaUnderlined(value bool) {
	class(self).SetMetaUnderline(value)
}

func (self Instance) HintUnderlined() bool {
	return bool(class(self).IsHintUnderlined())
}

func (self Instance) SetHintUnderlined(value bool) {
	class(self).SetHintUnderline(value)
}

func (self Instance) Threaded() bool {
	return bool(class(self).IsThreaded())
}

func (self Instance) SetThreaded(value bool) {
	class(self).SetThreaded(value)
}

func (self Instance) ProgressBarDelay() int {
	return int(int(class(self).GetProgressBarDelay()))
}

func (self Instance) SetProgressBarDelay(value int) {
	class(self).SetProgressBarDelay(int64(value))
}

func (self Instance) SelectionEnabled() bool {
	return bool(class(self).IsSelectionEnabled())
}

func (self Instance) SetSelectionEnabled(value bool) {
	class(self).SetSelectionEnabled(value)
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

func (self Instance) VisibleCharacters() int {
	return int(int(class(self).GetVisibleCharacters()))
}

func (self Instance) SetVisibleCharacters(value int) {
	class(self).SetVisibleCharacters(int64(value))
}

func (self Instance) VisibleCharactersBehavior() gdclass.TextServerVisibleCharactersBehavior {
	return gdclass.TextServerVisibleCharactersBehavior(class(self).GetVisibleCharactersBehavior())
}

func (self Instance) SetVisibleCharactersBehavior(value gdclass.TextServerVisibleCharactersBehavior) {
	class(self).SetVisibleCharactersBehavior(value)
}

func (self Instance) VisibleRatio() Float.X {
	return Float.X(Float.X(class(self).GetVisibleRatio()))
}

func (self Instance) SetVisibleRatio(value Float.X) {
	class(self).SetVisibleRatio(float64(value))
}

func (self Instance) TextDirection() gdclass.ControlTextDirection {
	return gdclass.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value gdclass.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(String.New(value))
}

func (self Instance) StructuredTextBidiOverride() gdclass.TextServerStructuredTextParser {
	return gdclass.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Instance) SetStructuredTextBidiOverride(value gdclass.TextServerStructuredTextParser) {
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Instance) StructuredTextBidiOverrideOptions() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetStructuredTextBidiOverrideOptions())))
}

func (self Instance) SetStructuredTextBidiOverrideOptions(value []any) {
	class(self).SetStructuredTextBidiOverrideOptions(gd.EngineArrayFromSlice(value))
}

/*
Returns the text without BBCode mark-up.
*/
//go:nosplit
func (self class) GetParsedText() String.Readable { //gd:RichTextLabel.get_parsed_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_parsed_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Adds raw non-BBCode-parsed text to the tag stack.
*/
//go:nosplit
func (self class) AddText(text String.Readable) { //gd:RichTextLabel.add_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_add_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetText(text String.Readable) { //gd:RichTextLabel.set_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AddImage(image [1]gdclass.Texture2D, width int64, height int64, color Color.RGBA, inline_align InlineAlignment, region Rect2.PositionSize, key variant.Any, pad bool, tooltip String.Readable, size_in_percent bool) { //gd:RichTextLabel.add_image
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, color)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, region)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(key)))
	callframe.Arg(frame, pad)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	callframe.Arg(frame, size_in_percent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_add_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Updates the existing images with the key [param key]. Only properties specified by [param mask] bits are updated. See [method add_image].
*/
//go:nosplit
func (self class) UpdateImage(key variant.Any, mask gdclass.RichTextLabelImageUpdateMask, image [1]gdclass.Texture2D, width int64, height int64, color Color.RGBA, inline_align InlineAlignment, region Rect2.PositionSize, pad bool, tooltip String.Readable, size_in_percent bool) { //gd:RichTextLabel.update_image
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(key)))
	callframe.Arg(frame, mask)
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, color)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, region)
	callframe.Arg(frame, pad)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	callframe.Arg(frame, size_in_percent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_update_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a newline tag to the tag stack.
*/
//go:nosplit
func (self class) Newline() { //gd:RichTextLabel.newline
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_newline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a paragraph of content from the label. Returns [code]true[/code] if the paragraph exists.
The [param paragraph] argument is the index of the paragraph to remove, it can take values in the interval [code][0, get_paragraph_count() - 1][/code].
If [param no_invalidate] is set to [code]true[/code], cache for the subsequent paragraphs is not invalidated. Use it for faster updates if deleted paragraph is fully self-contained (have no unclosed tags), or this call is part of the complex edit operation and [method invalidate_paragraph] will be called at the end of operation.
*/
//go:nosplit
func (self class) RemoveParagraph(paragraph int64, no_invalidate bool) bool { //gd:RichTextLabel.remove_paragraph
	var frame = callframe.New()
	callframe.Arg(frame, paragraph)
	callframe.Arg(frame, no_invalidate)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_remove_paragraph, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Invalidates [param paragraph] and all subsequent paragraphs cache.
*/
//go:nosplit
func (self class) InvalidateParagraph(paragraph int64) bool { //gd:RichTextLabel.invalidate_paragraph
	var frame = callframe.New()
	callframe.Arg(frame, paragraph)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_invalidate_paragraph, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a [code skip-lint][font][/code] tag to the tag stack. Overrides default fonts for its duration.
Passing [code]0[/code] to [param font_size] will use the existing default font size.
*/
//go:nosplit
func (self class) PushFont(font [1]gdclass.Font, font_size int64) { //gd:RichTextLabel.push_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][font_size][/code] tag to the tag stack. Overrides default font size for its duration.
*/
//go:nosplit
func (self class) PushFontSize(font_size int64) { //gd:RichTextLabel.push_font_size
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][font][/code] tag with a normal font to the tag stack.
*/
//go:nosplit
func (self class) PushNormal() { //gd:RichTextLabel.push_normal
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][font][/code] tag with a bold font to the tag stack. This is the same as adding a [code skip-lint][b][/code] tag if not currently in a [code skip-lint][i][/code] tag.
*/
//go:nosplit
func (self class) PushBold() { //gd:RichTextLabel.push_bold
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_bold, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][font][/code] tag with a bold italics font to the tag stack.
*/
//go:nosplit
func (self class) PushBoldItalics() { //gd:RichTextLabel.push_bold_italics
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_bold_italics, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][font][/code] tag with an italics font to the tag stack. This is the same as adding an [code skip-lint][i][/code] tag if not currently in a [code skip-lint][b][/code] tag.
*/
//go:nosplit
func (self class) PushItalics() { //gd:RichTextLabel.push_italics
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_italics, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][font][/code] tag with a monospace font to the tag stack.
*/
//go:nosplit
func (self class) PushMono() { //gd:RichTextLabel.push_mono
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_mono, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][color][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushColor(color Color.RGBA) { //gd:RichTextLabel.push_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][outline_size][/code] tag to the tag stack. Overrides default text outline size for its duration.
*/
//go:nosplit
func (self class) PushOutlineSize(outline_size int64) { //gd:RichTextLabel.push_outline_size
	var frame = callframe.New()
	callframe.Arg(frame, outline_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_outline_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][outline_color][/code] tag to the tag stack. Adds text outline for its duration.
*/
//go:nosplit
func (self class) PushOutlineColor(color Color.RGBA) { //gd:RichTextLabel.push_outline_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_outline_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][p][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushParagraph(alignment HorizontalAlignment, base_direction gdclass.ControlTextDirection, language String.Readable, st_parser gdclass.TextServerStructuredTextParser, justification_flags gdclass.TextServerJustificationFlag, tab_stops Packed.Array[float32]) { //gd:RichTextLabel.push_paragraph
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, base_direction)
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	callframe.Arg(frame, st_parser)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](tab_stops)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_paragraph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds an [code skip-lint][indent][/code] tag to the tag stack. Multiplies [param level] by current [member tab_size] to determine new margin length.
*/
//go:nosplit
func (self class) PushIndent(level int64) { //gd:RichTextLabel.push_indent
	var frame = callframe.New()
	callframe.Arg(frame, level)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_indent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds [code skip-lint][ol][/code] or [code skip-lint][ul][/code] tag to the tag stack. Multiplies [param level] by current [member tab_size] to determine new margin length.
*/
//go:nosplit
func (self class) PushList(level int64, atype gdclass.RichTextLabelListType, capitalize bool, bullet String.Readable) { //gd:RichTextLabel.push_list
	var frame = callframe.New()
	callframe.Arg(frame, level)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, capitalize)
	callframe.Arg(frame, pointers.Get(gd.InternalString(bullet)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a meta tag to the tag stack. Similar to the BBCode [code skip-lint][url=something]{text}[/url][/code], but supports non-[String] metadata types.
If [member meta_underlined] is [code]true[/code], meta tags display an underline. This behavior can be customized with [param underline_mode].
[b]Note:[/b] Meta tags do nothing by default when clicked. To assign behavior when clicked, connect [signal meta_clicked] to a function that is called when the meta tag is clicked.
*/
//go:nosplit
func (self class) PushMeta(data variant.Any, underline_mode gdclass.RichTextLabelMetaUnderline, tooltip String.Readable) { //gd:RichTextLabel.push_meta
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(data)))
	callframe.Arg(frame, underline_mode)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][hint][/code] tag to the tag stack. Same as BBCode [code skip-lint][hint=something]{text}[/hint][/code].
*/
//go:nosplit
func (self class) PushHint(description String.Readable) { //gd:RichTextLabel.push_hint
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(description)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds language code used for text shaping algorithm and Open-Type font features.
*/
//go:nosplit
func (self class) PushLanguage(language String.Readable) { //gd:RichTextLabel.push_language
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][u][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushUnderline() { //gd:RichTextLabel.push_underline
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_underline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][s][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushStrikethrough() { //gd:RichTextLabel.push_strikethrough
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_strikethrough, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][table=columns,inline_align][/code] tag to the tag stack. Use [method set_table_column_expand] to set column expansion ratio. Use [method push_cell] to add cells.
*/
//go:nosplit
func (self class) PushTable(columns int64, inline_align InlineAlignment, align_to_row int64) { //gd:RichTextLabel.push_table
	var frame = callframe.New()
	callframe.Arg(frame, columns)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, align_to_row)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_table, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][dropcap][/code] tag to the tag stack. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
//go:nosplit
func (self class) PushDropcap(s String.Readable, font [1]gdclass.Font, size int64, dropcap_margins Rect2.PositionSize, color Color.RGBA, outline_size int64, outline_color Color.RGBA) { //gd:RichTextLabel.push_dropcap
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, size)
	callframe.Arg(frame, dropcap_margins)
	callframe.Arg(frame, color)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, outline_color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_dropcap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Edits the selected column's expansion options. If [param expand] is [code]true[/code], the column expands in proportion to its expansion ratio versus the other columns' ratios.
For example, 2 columns with ratios 3 and 4 plus 70 pixels in available width would expand 30 and 40 pixels, respectively.
If [param expand] is [code]false[/code], the column will not contribute to the total ratio.
*/
//go:nosplit
func (self class) SetTableColumnExpand(column int64, expand bool, ratio int64, shrink bool) { //gd:RichTextLabel.set_table_column_expand
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, expand)
	callframe.Arg(frame, ratio)
	callframe.Arg(frame, shrink)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_table_column_expand, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets color of a table cell. Separate colors for alternating rows can be specified.
*/
//go:nosplit
func (self class) SetCellRowBackgroundColor(odd_row_bg Color.RGBA, even_row_bg Color.RGBA) { //gd:RichTextLabel.set_cell_row_background_color
	var frame = callframe.New()
	callframe.Arg(frame, odd_row_bg)
	callframe.Arg(frame, even_row_bg)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_cell_row_background_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets color of a table cell border.
*/
//go:nosplit
func (self class) SetCellBorderColor(color Color.RGBA) { //gd:RichTextLabel.set_cell_border_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_cell_border_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets minimum and maximum size overrides for a table cell.
*/
//go:nosplit
func (self class) SetCellSizeOverride(min_size Vector2.XY, max_size Vector2.XY) { //gd:RichTextLabel.set_cell_size_override
	var frame = callframe.New()
	callframe.Arg(frame, min_size)
	callframe.Arg(frame, max_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_cell_size_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets inner padding of a table cell.
*/
//go:nosplit
func (self class) SetCellPadding(padding Rect2.PositionSize) { //gd:RichTextLabel.set_cell_padding
	var frame = callframe.New()
	callframe.Arg(frame, padding)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_cell_padding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][cell][/code] tag to the tag stack. Must be inside a [code skip-lint][table][/code] tag. See [method push_table] for details. Use [method set_table_column_expand] to set column expansion ratio, [method set_cell_border_color] to set cell border, [method set_cell_row_background_color] to set cell background, [method set_cell_size_override] to override cell size, and [method set_cell_padding] to set padding.
*/
//go:nosplit
func (self class) PushCell() { //gd:RichTextLabel.push_cell
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][fgcolor][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushFgcolor(fgcolor Color.RGBA) { //gd:RichTextLabel.push_fgcolor
	var frame = callframe.New()
	callframe.Arg(frame, fgcolor)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_fgcolor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a [code skip-lint][bgcolor][/code] tag to the tag stack.
*/
//go:nosplit
func (self class) PushBgcolor(bgcolor Color.RGBA) { //gd:RichTextLabel.push_bgcolor
	var frame = callframe.New()
	callframe.Arg(frame, bgcolor)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_bgcolor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a custom effect tag to the tag stack. The effect does not need to be in [member custom_effects]. The environment is directly passed to the effect.
*/
//go:nosplit
func (self class) PushCustomfx(effect [1]gdclass.RichTextEffect, env Dictionary.Any) { //gd:RichTextLabel.push_customfx
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(effect[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(env)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_customfx, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a context marker to the tag stack. See [method pop_context].
*/
//go:nosplit
func (self class) PushContext() { //gd:RichTextLabel.push_context
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_push_context, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Terminates tags opened after the last [method push_context] call (including context marker), or all tags if there's no context marker on the stack.
*/
//go:nosplit
func (self class) PopContext() { //gd:RichTextLabel.pop_context
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_pop_context, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Terminates the current tag. Use after [code]push_*[/code] methods to close BBCodes manually. Does not need to follow [code]add_*[/code] methods.
*/
//go:nosplit
func (self class) Pop() { //gd:RichTextLabel.pop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_pop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Terminates all tags opened by [code]push_*[/code] methods.
*/
//go:nosplit
func (self class) PopAll() { //gd:RichTextLabel.pop_all
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_pop_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears the tag stack, causing the label to display nothing.
[b]Note:[/b] This method does not affect [member text], and its contents will show again if the label is redrawn. However, setting [member text] to an empty [String] also clears the stack.
*/
//go:nosplit
func (self class) Clear() { //gd:RichTextLabel.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser gdclass.TextServerStructuredTextParser) { //gd:RichTextLabel.set_structured_text_bidi_override
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverride() gdclass.TextServerStructuredTextParser { //gd:RichTextLabel.get_structured_text_bidi_override
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerStructuredTextParser](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args Array.Any) { //gd:RichTextLabel.set_structured_text_bidi_override_options
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(args)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions() Array.Any { //gd:RichTextLabel.get_structured_text_bidi_override_options
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextDirection(direction gdclass.ControlTextDirection) { //gd:RichTextLabel.set_text_direction
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() gdclass.ControlTextDirection { //gd:RichTextLabel.get_text_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language String.Readable) { //gd:RichTextLabel.set_language
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() String.Readable { //gd:RichTextLabel.get_language
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment HorizontalAlignment) { //gd:RichTextLabel.set_horizontal_alignment
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalAlignment() HorizontalAlignment { //gd:RichTextLabel.get_horizontal_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticalAlignment(alignment VerticalAlignment) { //gd:RichTextLabel.set_vertical_alignment
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticalAlignment() VerticalAlignment { //gd:RichTextLabel.get_vertical_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[VerticalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_vertical_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJustificationFlags(justification_flags gdclass.TextServerJustificationFlag) { //gd:RichTextLabel.set_justification_flags
	var frame = callframe.New()
	callframe.Arg(frame, justification_flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_justification_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJustificationFlags() gdclass.TextServerJustificationFlag { //gd:RichTextLabel.get_justification_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerJustificationFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_justification_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabStops(tab_stops Packed.Array[float32]) { //gd:RichTextLabel.set_tab_stops
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](tab_stops)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_tab_stops, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabStops() Packed.Array[float32] { //gd:RichTextLabel.get_tab_stops
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_tab_stops, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutowrapMode(autowrap_mode gdclass.TextServerAutowrapMode) { //gd:RichTextLabel.set_autowrap_mode
	var frame = callframe.New()
	callframe.Arg(frame, autowrap_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutowrapMode() gdclass.TextServerAutowrapMode { //gd:RichTextLabel.get_autowrap_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMetaUnderline(enable bool) { //gd:RichTextLabel.set_meta_underline
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_meta_underline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMetaUnderlined() bool { //gd:RichTextLabel.is_meta_underlined
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_meta_underlined, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHintUnderline(enable bool) { //gd:RichTextLabel.set_hint_underline
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_hint_underline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHintUnderlined() bool { //gd:RichTextLabel.is_hint_underlined
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_hint_underlined, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollActive(active bool) { //gd:RichTextLabel.set_scroll_active
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_scroll_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsScrollActive() bool { //gd:RichTextLabel.is_scroll_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_scroll_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollFollow(follow bool) { //gd:RichTextLabel.set_scroll_follow
	var frame = callframe.New()
	callframe.Arg(frame, follow)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_scroll_follow, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsScrollFollowing() bool { //gd:RichTextLabel.is_scroll_following
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_scroll_following, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the vertical scrollbar.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetVScrollBar() [1]gdclass.VScrollBar { //gd:RichTextLabel.get_v_scroll_bar
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.VScrollBar{gd.PointerLifetimeBoundTo[gdclass.VScrollBar](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Scrolls the window's top line to match [param line].
*/
//go:nosplit
func (self class) ScrollToLine(line int64) { //gd:RichTextLabel.scroll_to_line
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_scroll_to_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Scrolls the window's top line to match first line of the [param paragraph].
*/
//go:nosplit
func (self class) ScrollToParagraph(paragraph int64) { //gd:RichTextLabel.scroll_to_paragraph
	var frame = callframe.New()
	callframe.Arg(frame, paragraph)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_scroll_to_paragraph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Scrolls to the beginning of the current selection.
*/
//go:nosplit
func (self class) ScrollToSelection() { //gd:RichTextLabel.scroll_to_selection
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_scroll_to_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTabSize(spaces int64) { //gd:RichTextLabel.set_tab_size
	var frame = callframe.New()
	callframe.Arg(frame, spaces)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_tab_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabSize() int64 { //gd:RichTextLabel.get_tab_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_tab_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFitContent(enabled bool) { //gd:RichTextLabel.set_fit_content
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_fit_content, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFitContentEnabled() bool { //gd:RichTextLabel.is_fit_content_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_fit_content_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectionEnabled(enabled bool) { //gd:RichTextLabel.set_selection_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelectionEnabled() bool { //gd:RichTextLabel.is_selection_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContextMenuEnabled(enabled bool) { //gd:RichTextLabel.set_context_menu_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsContextMenuEnabled() bool { //gd:RichTextLabel.is_context_menu_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShortcutKeysEnabled(enabled bool) { //gd:RichTextLabel.set_shortcut_keys_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShortcutKeysEnabled() bool { //gd:RichTextLabel.is_shortcut_keys_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeselectOnFocusLossEnabled(enable bool) { //gd:RichTextLabel.set_deselect_on_focus_loss_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDeselectOnFocusLossEnabled() bool { //gd:RichTextLabel.is_deselect_on_focus_loss_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragAndDropSelectionEnabled(enable bool) { //gd:RichTextLabel.set_drag_and_drop_selection_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDragAndDropSelectionEnabled() bool { //gd:RichTextLabel.is_drag_and_drop_selection_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current selection first character index if a selection is active, [code]-1[/code] otherwise. Does not include BBCodes.
*/
//go:nosplit
func (self class) GetSelectionFrom() int64 { //gd:RichTextLabel.get_selection_from
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_selection_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current selection last character index if a selection is active, [code]-1[/code] otherwise. Does not include BBCodes.
*/
//go:nosplit
func (self class) GetSelectionTo() int64 { //gd:RichTextLabel.get_selection_to
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_selection_to, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current selection vertical line offset if a selection is active, [code]-1.0[/code] otherwise.
*/
//go:nosplit
func (self class) GetSelectionLineOffset() float64 { //gd:RichTextLabel.get_selection_line_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_selection_line_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Select all the text.
If [member selection_enabled] is [code]false[/code], no selection will occur.
*/
//go:nosplit
func (self class) SelectAll() { //gd:RichTextLabel.select_all
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_select_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current selection text. Does not include BBCodes.
*/
//go:nosplit
func (self class) GetSelectedText() String.Readable { //gd:RichTextLabel.get_selected_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_selected_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Clears the current selection.
*/
//go:nosplit
func (self class) Deselect() { //gd:RichTextLabel.deselect
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
The assignment version of [method append_text]. Clears the tag stack and inserts the new content.
*/
//go:nosplit
func (self class) ParseBbcode(bbcode String.Readable) { //gd:RichTextLabel.parse_bbcode
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(bbcode)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_parse_bbcode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Parses [param bbcode] and adds tags to the tag stack as needed.
[b]Note:[/b] Using this method, you can't close a tag that was opened in a previous [method append_text] call. This is done to improve performance, especially when updating large RichTextLabels since rebuilding the whole BBCode every time would be slower. If you absolutely need to close a tag in a future method call, append the [member text] instead of using [method append_text].
*/
//go:nosplit
func (self class) AppendText(bbcode String.Readable) { //gd:RichTextLabel.append_text
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(bbcode)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_append_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() String.Readable { //gd:RichTextLabel.get_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
If [member threaded] is enabled, returns [code]true[/code] if the background thread has finished text processing, otherwise always return [code]true[/code].
*/
//go:nosplit
func (self class) IsReady() bool { //gd:RichTextLabel.is_ready
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_ready, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [member threaded] is enabled, returns [code]true[/code] if the background thread has finished text processing, otherwise always return [code]true[/code].
*/
//go:nosplit
func (self class) IsFinished() bool { //gd:RichTextLabel.is_finished
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_finished, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetThreaded(threaded bool) { //gd:RichTextLabel.set_threaded
	var frame = callframe.New()
	callframe.Arg(frame, threaded)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_threaded, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsThreaded() bool { //gd:RichTextLabel.is_threaded
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_threaded, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProgressBarDelay(delay_ms int64) { //gd:RichTextLabel.set_progress_bar_delay
	var frame = callframe.New()
	callframe.Arg(frame, delay_ms)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_progress_bar_delay, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProgressBarDelay() int64 { //gd:RichTextLabel.get_progress_bar_delay
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_progress_bar_delay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibleCharacters(amount int64) { //gd:RichTextLabel.set_visible_characters
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_visible_characters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibleCharacters() int64 { //gd:RichTextLabel.get_visible_characters
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_visible_characters, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetVisibleCharactersBehavior() gdclass.TextServerVisibleCharactersBehavior { //gd:RichTextLabel.get_visible_characters_behavior
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerVisibleCharactersBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_visible_characters_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibleCharactersBehavior(behavior gdclass.TextServerVisibleCharactersBehavior) { //gd:RichTextLabel.set_visible_characters_behavior
	var frame = callframe.New()
	callframe.Arg(frame, behavior)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_visible_characters_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibleRatio(ratio float64) { //gd:RichTextLabel.set_visible_ratio
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_visible_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibleRatio() float64 { //gd:RichTextLabel.get_visible_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_visible_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the line number of the character position provided. Line and character numbers are both zero-indexed.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetCharacterLine(character int64) int64 { //gd:RichTextLabel.get_character_line
	var frame = callframe.New()
	callframe.Arg(frame, character)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_character_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the paragraph number of the character position provided. Paragraph and character numbers are both zero-indexed.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetCharacterParagraph(character int64) int64 { //gd:RichTextLabel.get_character_paragraph
	var frame = callframe.New()
	callframe.Arg(frame, character)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_character_paragraph, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total number of characters from text tags. Does not include BBCodes.
*/
//go:nosplit
func (self class) GetTotalCharacterCount() int64 { //gd:RichTextLabel.get_total_character_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_total_character_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseBbcode(enable bool) { //gd:RichTextLabel.set_use_bbcode
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_use_bbcode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingBbcode() bool { //gd:RichTextLabel.is_using_bbcode
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_using_bbcode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total number of lines in the text. Wrapped text is counted as multiple lines.
[b]Note:[/b] If [member visible_characters_behavior] is set to [constant TextServer.VC_CHARS_BEFORE_SHAPING] only visible wrapped lines are counted.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetLineCount() int64 { //gd:RichTextLabel.get_line_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the indexes of the first and last visible characters for the given [param line], as a [Vector2i].
[b]Note:[/b] If [member visible_characters_behavior] is set to [constant TextServer.VC_CHARS_BEFORE_SHAPING] only visible wrapped lines are counted.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetLineRange(line int64) Vector2i.XY { //gd:RichTextLabel.get_line_range
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_line_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of visible lines.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetVisibleLineCount() int64 { //gd:RichTextLabel.get_visible_line_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_visible_line_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total number of paragraphs (newlines or [code]p[/code] tags in the tag stack's text tags). Considers wrapped text as one paragraph.
*/
//go:nosplit
func (self class) GetParagraphCount() int64 { //gd:RichTextLabel.get_paragraph_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_paragraph_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of visible paragraphs. A paragraph is considered visible if at least one of its lines is visible.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetVisibleParagraphCount() int64 { //gd:RichTextLabel.get_visible_paragraph_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_visible_paragraph_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the height of the content.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetContentHeight() int64 { //gd:RichTextLabel.get_content_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_content_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the width of the content.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetContentWidth() int64 { //gd:RichTextLabel.get_content_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_content_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the vertical offset of the line found at the provided index.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetLineOffset(line int64) float64 { //gd:RichTextLabel.get_line_offset
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_line_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the vertical offset of the paragraph found at the provided index.
[b]Note:[/b] If [member threaded] is enabled, this method returns a value for the loaded part of the document. Use [method is_finished] or [signal finished] to determine whether document is fully loaded.
*/
//go:nosplit
func (self class) GetParagraphOffset(paragraph int64) float64 { //gd:RichTextLabel.get_paragraph_offset
	var frame = callframe.New()
	callframe.Arg(frame, paragraph)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_paragraph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Parses BBCode parameter [param expressions] into a dictionary.
*/
//go:nosplit
func (self class) ParseExpressionsForValues(expressions Packed.Strings) Dictionary.Any { //gd:RichTextLabel.parse_expressions_for_values
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(expressions)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_parse_expressions_for_values, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEffects(effects Array.Any) { //gd:RichTextLabel.set_effects
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(effects)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_set_effects, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEffects() Array.Any { //gd:RichTextLabel.get_effects
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_effects, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Installs a custom effect. This can also be done in the Inspector through the [member custom_effects] property. [param effect] should be a valid [RichTextEffect].
[b]Example:[/b] With the following script extending from [RichTextEffect]:
[codeblock]
# effect.gd
class_name MyCustomEffect
extends RichTextEffect

var bbcode = "my_custom_effect"

# ...
[/codeblock]
The above effect can be installed in [RichTextLabel] from a script:
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
func (self class) InstallEffect(effect variant.Any) { //gd:RichTextLabel.install_effect
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(effect)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_install_effect, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) GetMenu() [1]gdclass.PopupMenu { //gd:RichTextLabel.get_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PopupMenu{gd.PointerLifetimeBoundTo[gdclass.PopupMenu](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
//go:nosplit
func (self class) IsMenuVisible() bool { //gd:RichTextLabel.is_menu_visible
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_is_menu_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
//go:nosplit
func (self class) MenuOption(option int64) { //gd:RichTextLabel.menu_option
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RichTextLabel.Bind_menu_option, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnMetaClicked(cb func(meta any)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("meta_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMetaHoverStarted(cb func(meta any)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("meta_hover_started"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMetaHoverEnded(cb func(meta any)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("meta_hover_ended"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsRichTextLabel() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRichTextLabel() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced  { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
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
	gdclass.Register("RichTextLabel", func(ptr gd.Object) any {
		return [1]gdclass.RichTextLabel{*(*gdclass.RichTextLabel)(unsafe.Pointer(&ptr))}
	})
}

type ListType = gdclass.RichTextLabelListType //gd:RichTextLabel.ListType

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

type MenuItems = gdclass.RichTextLabelMenuItems //gd:RichTextLabel.MenuItems

const (
	/*Copies the selected text.*/
	MenuCopy MenuItems = 0
	/*Selects the whole [RichTextLabel] text.*/
	MenuSelectAll MenuItems = 1
	/*Represents the size of the [enum MenuItems] enum.*/
	MenuMax MenuItems = 2
)

type MetaUnderline = gdclass.RichTextLabelMetaUnderline //gd:RichTextLabel.MetaUnderline

const (
	/*Meta tag does not display an underline, even if [member meta_underlined] is [code]true[/code].*/
	MetaUnderlineNever MetaUnderline = 0
	/*If [member meta_underlined] is [code]true[/code], meta tag always display an underline.*/
	MetaUnderlineAlways MetaUnderline = 1
	/*If [member meta_underlined] is [code]true[/code], meta tag display an underline when the mouse cursor is over it.*/
	MetaUnderlineOnHover MetaUnderline = 2
)

type ImageUpdateMask = gdclass.RichTextLabelImageUpdateMask //gd:RichTextLabel.ImageUpdateMask

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

type InlineAlignment int

const (
	/*Aligns the top of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentTopTo InlineAlignment = 0
	/*Aligns the center of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentCenterTo InlineAlignment = 1
	/*Aligns the baseline (user defined) of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentBaselineTo InlineAlignment = 3
	/*Aligns the bottom of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentBottomTo InlineAlignment = 2
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the top of the text.*/
	InlineAlignmentToTop InlineAlignment = 0
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the center of the text.*/
	InlineAlignmentToCenter InlineAlignment = 4
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the baseline of the text.*/
	InlineAlignmentToBaseline InlineAlignment = 8
	/*Aligns inline object (e.g. image, table) to the bottom of the text.*/
	InlineAlignmentToBottom InlineAlignment = 12
	/*Aligns top of the inline object (e.g. image, table) to the top of the text. Equivalent to [code]INLINE_ALIGNMENT_TOP_TO | INLINE_ALIGNMENT_TO_TOP[/code].*/
	InlineAlignmentTop InlineAlignment = 0
	/*Aligns center of the inline object (e.g. image, table) to the center of the text. Equivalent to [code]INLINE_ALIGNMENT_CENTER_TO | INLINE_ALIGNMENT_TO_CENTER[/code].*/
	InlineAlignmentCenter InlineAlignment = 5
	/*Aligns bottom of the inline object (e.g. image, table) to the bottom of the text. Equivalent to [code]INLINE_ALIGNMENT_BOTTOM_TO | INLINE_ALIGNMENT_TO_BOTTOM[/code].*/
	InlineAlignmentBottom InlineAlignment = 14
	/*A bit mask for [code]INLINE_ALIGNMENT_*_TO[/code] alignment constants.*/
	InlineAlignmentImageMask InlineAlignment = 3
	/*A bit mask for [code]INLINE_ALIGNMENT_TO_*[/code] alignment constants.*/
	InlineAlignmentTextMask InlineAlignment = 12
)

type VerticalAlignment int

const (
	/*Vertical top alignment, usually for text-derived classes.*/
	VerticalAlignmentTop VerticalAlignment = 0
	/*Vertical center alignment, usually for text-derived classes.*/
	VerticalAlignmentCenter VerticalAlignment = 1
	/*Vertical bottom alignment, usually for text-derived classes.*/
	VerticalAlignmentBottom VerticalAlignment = 2
	/*Expand rows to fit height, usually for text-derived classes.*/
	VerticalAlignmentFill VerticalAlignment = 3
)
