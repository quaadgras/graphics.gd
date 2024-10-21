package RichTextLabel

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
A control for displaying text that can contain custom fonts, images, and basic formatting. [RichTextLabel] manages these as an internal tag stack. It also adapts itself to given width/heights.
[b]Note:[/b] Assignments to [member text] clear the tag stack and reconstruct it from the property's contents. Any edits made to [member text] will erase previous edits made from other manual sources such as [method append_text] and the [code]push_*[/code] / [method pop] methods.
[b]Note:[/b] RichTextLabel doesn't support entangled BBCode tags. For example, instead of using [code skip-lint][b]bold[i]bold italic[/b]italic[/i][/code], use [code skip-lint][b]bold[i]bold italic[/i][/b][i]italic[/i][/code].
[b]Note:[/b] [code]push_pop_*[/code] functions won't affect BBCode.
[b]Note:[/b] Unlike [Label], [RichTextLabel] doesn't have a [i]property[/i] to horizontally align text to the center. Instead, enable [member bbcode_enabled] and surround the text in a [code skip-lint][center][/code] tag as follows: [code skip-lint][center]Example[/center][/code]. There is currently no built-in way to vertically align text either, but this can be emulated by relying on anchors/containers and the [member fit_content] property.

*/
type Simple [1]classdb.RichTextLabel
func (self Simple) GetParsedText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetParsedText(gc).String())
}
func (self Simple) AddText(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddText(gc.String(text))
}
func (self Simple) SetText(text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetText(gc.String(text))
}
func (self Simple) AddImage(image [1]classdb.Texture2D, width int, height int, color gd.Color, inline_align gd.InlineAlignment, region gd.Rect2, key gd.Variant, pad bool, tooltip string, size_in_percent bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddImage(image, gd.Int(width), gd.Int(height), color, inline_align, region, key, pad, gc.String(tooltip), size_in_percent)
}
func (self Simple) UpdateImage(key gd.Variant, mask classdb.RichTextLabelImageUpdateMask, image [1]classdb.Texture2D, width int, height int, color gd.Color, inline_align gd.InlineAlignment, region gd.Rect2, pad bool, tooltip string, size_in_percent bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateImage(key, mask, image, gd.Int(width), gd.Int(height), color, inline_align, region, pad, gc.String(tooltip), size_in_percent)
}
func (self Simple) Newline() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Newline()
}
func (self Simple) RemoveParagraph(paragraph int, no_invalidate bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).RemoveParagraph(gd.Int(paragraph), no_invalidate))
}
func (self Simple) InvalidateParagraph(paragraph int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).InvalidateParagraph(gd.Int(paragraph)))
}
func (self Simple) PushFont(font [1]classdb.Font, font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushFont(font, gd.Int(font_size))
}
func (self Simple) PushFontSize(font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushFontSize(gd.Int(font_size))
}
func (self Simple) PushNormal() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushNormal()
}
func (self Simple) PushBold() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushBold()
}
func (self Simple) PushBoldItalics() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushBoldItalics()
}
func (self Simple) PushItalics() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushItalics()
}
func (self Simple) PushMono() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushMono()
}
func (self Simple) PushColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushColor(color)
}
func (self Simple) PushOutlineSize(outline_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushOutlineSize(gd.Int(outline_size))
}
func (self Simple) PushOutlineColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushOutlineColor(color)
}
func (self Simple) PushParagraph(alignment gd.HorizontalAlignment, base_direction classdb.ControlTextDirection, language string, st_parser classdb.TextServerStructuredTextParser, justification_flags classdb.TextServerJustificationFlag, tab_stops gd.PackedFloat32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushParagraph(alignment, base_direction, gc.String(language), st_parser, justification_flags, tab_stops)
}
func (self Simple) PushIndent(level int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushIndent(gd.Int(level))
}
func (self Simple) PushList(level int, atype classdb.RichTextLabelListType, capitalize bool, bullet string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushList(gd.Int(level), atype, capitalize, gc.String(bullet))
}
func (self Simple) PushMeta(data gd.Variant, underline_mode classdb.RichTextLabelMetaUnderline) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushMeta(data, underline_mode)
}
func (self Simple) PushHint(description string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushHint(gc.String(description))
}
func (self Simple) PushLanguage(language string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushLanguage(gc.String(language))
}
func (self Simple) PushUnderline() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushUnderline()
}
func (self Simple) PushStrikethrough() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushStrikethrough()
}
func (self Simple) PushTable(columns int, inline_align gd.InlineAlignment, align_to_row int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushTable(gd.Int(columns), inline_align, gd.Int(align_to_row))
}
func (self Simple) PushDropcap(s string, font [1]classdb.Font, size int, dropcap_margins gd.Rect2, color gd.Color, outline_size int, outline_color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushDropcap(gc.String(s), font, gd.Int(size), dropcap_margins, color, gd.Int(outline_size), outline_color)
}
func (self Simple) SetTableColumnExpand(column int, expand bool, ratio int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTableColumnExpand(gd.Int(column), expand, gd.Int(ratio))
}
func (self Simple) SetCellRowBackgroundColor(odd_row_bg gd.Color, even_row_bg gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellRowBackgroundColor(odd_row_bg, even_row_bg)
}
func (self Simple) SetCellBorderColor(color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellBorderColor(color)
}
func (self Simple) SetCellSizeOverride(min_size gd.Vector2, max_size gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellSizeOverride(min_size, max_size)
}
func (self Simple) SetCellPadding(padding gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellPadding(padding)
}
func (self Simple) PushCell() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushCell()
}
func (self Simple) PushFgcolor(fgcolor gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushFgcolor(fgcolor)
}
func (self Simple) PushBgcolor(bgcolor gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushBgcolor(bgcolor)
}
func (self Simple) PushCustomfx(effect [1]classdb.RichTextEffect, env gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushCustomfx(effect, env)
}
func (self Simple) PushContext() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PushContext()
}
func (self Simple) PopContext() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopContext()
}
func (self Simple) Pop() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Pop()
}
func (self Simple) PopAll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).PopAll()
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
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
func (self Simple) SetAutowrapMode(autowrap_mode classdb.TextServerAutowrapMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutowrapMode(autowrap_mode)
}
func (self Simple) GetAutowrapMode() classdb.TextServerAutowrapMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerAutowrapMode(Expert(self).GetAutowrapMode())
}
func (self Simple) SetMetaUnderline(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMetaUnderline(enable)
}
func (self Simple) IsMetaUnderlined() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMetaUnderlined())
}
func (self Simple) SetHintUnderline(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHintUnderline(enable)
}
func (self Simple) IsHintUnderlined() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsHintUnderlined())
}
func (self Simple) SetScrollActive(active bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScrollActive(active)
}
func (self Simple) IsScrollActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsScrollActive())
}
func (self Simple) SetScrollFollow(follow bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScrollFollow(follow)
}
func (self Simple) IsScrollFollowing() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsScrollFollowing())
}
func (self Simple) GetVScrollBar() [1]classdb.VScrollBar {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.VScrollBar(Expert(self).GetVScrollBar(gc))
}
func (self Simple) ScrollToLine(line int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScrollToLine(gd.Int(line))
}
func (self Simple) ScrollToParagraph(paragraph int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScrollToParagraph(gd.Int(paragraph))
}
func (self Simple) ScrollToSelection() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScrollToSelection()
}
func (self Simple) SetTabSize(spaces int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTabSize(gd.Int(spaces))
}
func (self Simple) GetTabSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTabSize()))
}
func (self Simple) SetFitContent(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFitContent(enabled)
}
func (self Simple) IsFitContentEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsFitContentEnabled())
}
func (self Simple) SetSelectionEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSelectionEnabled(enabled)
}
func (self Simple) IsSelectionEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSelectionEnabled())
}
func (self Simple) SetContextMenuEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContextMenuEnabled(enabled)
}
func (self Simple) IsContextMenuEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsContextMenuEnabled())
}
func (self Simple) SetShortcutKeysEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShortcutKeysEnabled(enabled)
}
func (self Simple) IsShortcutKeysEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsShortcutKeysEnabled())
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
func (self Simple) GetSelectionFrom() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSelectionFrom()))
}
func (self Simple) GetSelectionTo() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSelectionTo()))
}
func (self Simple) SelectAll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SelectAll()
}
func (self Simple) GetSelectedText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSelectedText(gc).String())
}
func (self Simple) Deselect() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Deselect()
}
func (self Simple) ParseBbcode(bbcode string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ParseBbcode(gc.String(bbcode))
}
func (self Simple) AppendText(bbcode string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AppendText(gc.String(bbcode))
}
func (self Simple) GetText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetText(gc).String())
}
func (self Simple) IsReady() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsReady())
}
func (self Simple) SetThreaded(threaded bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetThreaded(threaded)
}
func (self Simple) IsThreaded() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsThreaded())
}
func (self Simple) SetProgressBarDelay(delay_ms int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProgressBarDelay(gd.Int(delay_ms))
}
func (self Simple) GetProgressBarDelay() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetProgressBarDelay()))
}
func (self Simple) SetVisibleCharacters(amount int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisibleCharacters(gd.Int(amount))
}
func (self Simple) GetVisibleCharacters() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetVisibleCharacters()))
}
func (self Simple) GetVisibleCharactersBehavior() classdb.TextServerVisibleCharactersBehavior {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerVisibleCharactersBehavior(Expert(self).GetVisibleCharactersBehavior())
}
func (self Simple) SetVisibleCharactersBehavior(behavior classdb.TextServerVisibleCharactersBehavior) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisibleCharactersBehavior(behavior)
}
func (self Simple) SetVisibleRatio(ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVisibleRatio(gd.Float(ratio))
}
func (self Simple) GetVisibleRatio() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetVisibleRatio()))
}
func (self Simple) GetCharacterLine(character int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCharacterLine(gd.Int(character))))
}
func (self Simple) GetCharacterParagraph(character int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCharacterParagraph(gd.Int(character))))
}
func (self Simple) GetTotalCharacterCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTotalCharacterCount()))
}
func (self Simple) SetUseBbcode(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseBbcode(enable)
}
func (self Simple) IsUsingBbcode() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsUsingBbcode())
}
func (self Simple) GetLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLineCount()))
}
func (self Simple) GetVisibleLineCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetVisibleLineCount()))
}
func (self Simple) GetParagraphCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetParagraphCount()))
}
func (self Simple) GetVisibleParagraphCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetVisibleParagraphCount()))
}
func (self Simple) GetContentHeight() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetContentHeight()))
}
func (self Simple) GetContentWidth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetContentWidth()))
}
func (self Simple) GetLineOffset(line int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetLineOffset(gd.Int(line))))
}
func (self Simple) GetParagraphOffset(paragraph int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParagraphOffset(gd.Int(paragraph))))
}
func (self Simple) ParseExpressionsForValues(expressions gd.PackedStringArray) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).ParseExpressionsForValues(gc, expressions))
}
func (self Simple) SetEffects(effects gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEffects(effects)
}
func (self Simple) GetEffects() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetEffects(gc))
}
func (self Simple) InstallEffect(effect gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InstallEffect(effect)
}
func (self Simple) GetMenu() [1]classdb.PopupMenu {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.PopupMenu(Expert(self).GetMenu(gc))
}
func (self Simple) IsMenuVisible() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMenuVisible())
}
func (self Simple) MenuOption(option int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MenuOption(gd.Int(option))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.RichTextLabel
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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
func (self class) AddImage(image object.Texture2D, width gd.Int, height gd.Int, color gd.Color, inline_align gd.InlineAlignment, region gd.Rect2, key gd.Variant, pad bool, tooltip gd.String, size_in_percent bool)  {
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
func (self class) UpdateImage(key gd.Variant, mask classdb.RichTextLabelImageUpdateMask, image object.Texture2D, width gd.Int, height gd.Int, color gd.Color, inline_align gd.InlineAlignment, region gd.Rect2, pad bool, tooltip gd.String, size_in_percent bool)  {
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
func (self class) PushFont(font object.Font, font_size gd.Int)  {
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
func (self class) PushDropcap(s gd.String, font object.Font, size gd.Int, dropcap_margins gd.Rect2, color gd.Color, outline_size gd.Int, outline_color gd.Color)  {
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
func (self class) PushCustomfx(effect object.RichTextEffect, env gd.Dictionary)  {
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
func (self class) GetVScrollBar(ctx gd.Lifetime) object.VScrollBar {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_v_scroll_bar, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.VScrollBar
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
func (self class) GetMenu(ctx gd.Lifetime) object.PopupMenu {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.RichTextLabel.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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

//go:nosplit
func (self class) AsRichTextLabel() Expert { return self[0].AsRichTextLabel() }


//go:nosplit
func (self Simple) AsRichTextLabel() Simple { return self[0].AsRichTextLabel() }


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
func init() {classdb.Register("RichTextLabel", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
