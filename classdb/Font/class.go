// Package Font provides methods for working with Font object instances.
package Font

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Abstract base class for different font types. It has methods for drawing text and font character introspection.
*/
type Instance [1]gdclass.Font

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFont() Instance
}

/*
Returns [TextServer] RID of the font cache for specific variation.
*/
func (self Instance) FindVariation(variation_coordinates Dictionary.Any) Resource.ID {
	return Resource.ID(class(self).FindVariation(variation_coordinates, gd.Int(0), gd.Float(0.0), gd.Transform2D(gd.NewTransform2D(1, 0, 0, 1, 0, 0)), gd.Int(0), gd.Int(0), gd.Int(0), gd.Int(0), gd.Float(0.0)))
}

/*
Returns [Array] of valid [Font] [RID]s, which can be passed to the [TextServer] methods.
*/
func (self Instance) GetRids() gd.Array {
	return gd.Array(class(self).GetRids())
}

/*
Returns the total average font height (ascent plus descent) in pixels.
[b]Note:[/b] Real height of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the height of empty line).
*/
func (self Instance) GetHeight() Float.X {
	return Float.X(Float.X(class(self).GetHeight(gd.Int(16))))
}

/*
Returns the average font ascent (number of pixels above the baseline).
[b]Note:[/b] Real ascent of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the ascent of empty line).
*/
func (self Instance) GetAscent() Float.X {
	return Float.X(Float.X(class(self).GetAscent(gd.Int(16))))
}

/*
Returns the average font descent (number of pixels below the baseline).
[b]Note:[/b] Real descent of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the descent of empty line).
*/
func (self Instance) GetDescent() Float.X {
	return Float.X(Float.X(class(self).GetDescent(gd.Int(16))))
}

/*
Returns average pixel offset of the underline below the baseline.
[b]Note:[/b] Real underline position of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate.
*/
func (self Instance) GetUnderlinePosition() Float.X {
	return Float.X(Float.X(class(self).GetUnderlinePosition(gd.Int(16))))
}

/*
Returns average thickness of the underline.
[b]Note:[/b] Real underline thickness of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate.
*/
func (self Instance) GetUnderlineThickness() Float.X {
	return Float.X(Float.X(class(self).GetUnderlineThickness(gd.Int(16))))
}

/*
Returns font family name.
*/
func (self Instance) GetFontName() string {
	return string(class(self).GetFontName().String())
}

/*
Returns font style name.
*/
func (self Instance) GetFontStyleName() string {
	return string(class(self).GetFontStyleName().String())
}

/*
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
func (self Instance) GetOtNameStrings() Dictionary.Any {
	return Dictionary.Any(class(self).GetOtNameStrings())
}

/*
Returns font style flags, see [enum TextServer.FontStyle].
*/
func (self Instance) GetFontStyle() gdclass.TextServerFontStyle {
	return gdclass.TextServerFontStyle(class(self).GetFontStyle())
}

/*
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (self Instance) GetFontWeight() int {
	return int(int(class(self).GetFontWeight()))
}

/*
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (self Instance) GetFontStretch() int {
	return int(int(class(self).GetFontStretch()))
}

/*
Returns the spacing for the given [code]type[/code] (see [enum TextServer.SpacingType]).
*/
func (self Instance) GetSpacing(spacing gdclass.TextServerSpacingType) int {
	return int(int(class(self).GetSpacing(spacing)))
}

/*
Returns a set of OpenType feature tags. More info: [url=https://docs.microsoft.com/en-us/typography/opentype/spec/featuretags]OpenType feature tags[/url].
*/
func (self Instance) GetOpentypeFeatures() Dictionary.Any {
	return Dictionary.Any(class(self).GetOpentypeFeatures())
}

/*
Sets LRU cache capacity for [code]draw_*[/code] methods.
*/
func (self Instance) SetCacheCapacity(single_line int, multi_line int) {
	class(self).SetCacheCapacity(gd.Int(single_line), gd.Int(multi_line))
}

/*
Returns the size of a bounding box of a single-line string, taking kerning, advance and subpixel positioning into account. See also [method get_multiline_string_size] and [method draw_string].
For example, to get the string size as displayed by a single-line Label, use:
[codeblocks]
[gdscript]
var string_size = $Label.get_theme_font("font").get_string_size($Label.text, HORIZONTAL_ALIGNMENT_LEFT, -1, $Label.get_theme_font_size("font_size"))
[/gdscript]
[csharp]
Label label = GetNode<Label>("Label");
Vector2 stringSize = label.GetThemeFont("font").GetStringSize(label.Text, HorizontalAlignment.Left, -1, label.GetThemeFontSize("font_size"));
[/csharp]
[/codeblocks]
[b]Note:[/b] Since kerning, advance and subpixel positioning are taken into account by [method get_string_size], using separate [method get_string_size] calls on substrings of a string then adding the results together will return a different result compared to using a single [method get_string_size] call on the full string.
[b]Note:[/b] Real height of the string is context-dependent and can be significantly different from the value returned by [method get_height].
*/
func (self Instance) GetStringSize(text string) Vector2.XY {
	return Vector2.XY(class(self).GetStringSize(gd.NewString(text), 0, gd.Float(-1), gd.Int(16), 3, 0, 0))
}

/*
Returns the size of a bounding box of a string broken into the lines, taking kerning and advance into account.
See also [method draw_multiline_string].
*/
func (self Instance) GetMultilineStringSize(text string) Vector2.XY {
	return Vector2.XY(class(self).GetMultilineStringSize(gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), 3, 3, 0, 0))
}

/*
Draw [param text] into a canvas item using the font, at a given position, with [param modulate] color, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_string].
*/
func (self Instance) DrawString(canvas_item Resource.ID, pos Vector2.XY, text string) {
	class(self).DrawString(canvas_item, gd.Vector2(pos), gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Color(gd.Color{1, 1, 1, 1}), 3, 0, 0)
}

/*
Breaks [param text] into lines using rules specified by [param brk_flags] and draws it into a canvas item using the font, at a given position, with [param modulate] color, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline of the first line, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_multiline_string].
*/
func (self Instance) DrawMultilineString(canvas_item Resource.ID, pos Vector2.XY, text string) {
	class(self).DrawMultilineString(canvas_item, gd.Vector2(pos), gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), gd.Color(gd.Color{1, 1, 1, 1}), 3, 3, 0, 0)
}

/*
Draw [param text] outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_string_outline].
*/
func (self Instance) DrawStringOutline(canvas_item Resource.ID, pos Vector2.XY, text string) {
	class(self).DrawStringOutline(canvas_item, gd.Vector2(pos), gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}), 3, 0, 0)
}

/*
Breaks [param text] to the lines using rules specified by [param brk_flags] and draws text outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline of the first line, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_multiline_string_outline].
*/
func (self Instance) DrawMultilineStringOutline(canvas_item Resource.ID, pos Vector2.XY, text string) {
	class(self).DrawMultilineStringOutline(canvas_item, gd.Vector2(pos), gd.NewString(text), 0, gd.Float(-1), gd.Int(16), gd.Int(-1), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}), 3, 3, 0, 0)
}

/*
Returns the size of a character. Does not take kerning into account.
[b]Note:[/b] Do not use this function to calculate width of the string character by character, use [method get_string_size] or [TextLine] instead. The height returned is the font height (see also [method get_height]) and has no relation to the glyph height.
*/
func (self Instance) GetCharSize(char int, font_size int) Vector2.XY {
	return Vector2.XY(class(self).GetCharSize(gd.Int(char), gd.Int(font_size)))
}

/*
Draw a single Unicode character [param char] into a canvas item using the font, at a given position, with [param modulate] color. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
[b]Note:[/b] Do not use this function to draw strings character by character, use [method draw_string] or [TextLine] instead.
*/
func (self Instance) DrawChar(canvas_item Resource.ID, pos Vector2.XY, char int, font_size int) Float.X {
	return Float.X(Float.X(class(self).DrawChar(canvas_item, gd.Vector2(pos), gd.Int(char), gd.Int(font_size), gd.Color(gd.Color{1, 1, 1, 1}))))
}

/*
Draw a single Unicode character [param char] outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
[b]Note:[/b] Do not use this function to draw strings character by character, use [method draw_string] or [TextLine] instead.
*/
func (self Instance) DrawCharOutline(canvas_item Resource.ID, pos Vector2.XY, char int, font_size int) Float.X {
	return Float.X(Float.X(class(self).DrawCharOutline(canvas_item, gd.Vector2(pos), gd.Int(char), gd.Int(font_size), gd.Int(-1), gd.Color(gd.Color{1, 1, 1, 1}))))
}

/*
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
func (self Instance) HasChar(char int) bool {
	return bool(class(self).HasChar(gd.Int(char)))
}

/*
Returns a string containing all the characters available in the font.
If a given character is included in more than one font data source, it appears only once in the returned string.
*/
func (self Instance) GetSupportedChars() string {
	return string(class(self).GetSupportedChars().String())
}

/*
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
func (self Instance) IsLanguageSupported(language string) bool {
	return bool(class(self).IsLanguageSupported(gd.NewString(language)))
}

/*
Returns [code]true[/code], if font supports given script ([url=https://en.wikipedia.org/wiki/ISO_15924]ISO 15924[/url] code).
*/
func (self Instance) IsScriptSupported(script string) bool {
	return bool(class(self).IsScriptSupported(gd.NewString(script)))
}

/*
Returns list of OpenType features supported by font.
*/
func (self Instance) GetSupportedFeatureList() Dictionary.Any {
	return Dictionary.Any(class(self).GetSupportedFeatureList())
}

/*
Returns list of supported [url=https://docs.microsoft.com/en-us/typography/opentype/spec/dvaraxisreg]variation coordinates[/url], each coordinate is returned as [code]tag: Vector3i(min_value,max_value,default_value)[/code].
Font variations allow for continuous change of glyph characteristics along some given design axis, such as weight, width or slant.
To print available variation axes of a variable font:
[codeblock]
var fv = FontVariation.new()
fv.base_font = load("res://RobotoFlex.ttf")
var variation_list = fv.get_supported_variation_list()
for tag in variation_list:

	var name = TextServerManager.get_primary_interface().tag_to_name(tag)
	var values = variation_list[tag]
	print("variation axis: %s (%d)\n\tmin, max, default: %s" % [name, tag, values])

[/codeblock]
[b]Note:[/b] To set and get variation coordinates of a [FontVariation], use [member FontVariation.variation_opentype].
*/
func (self Instance) GetSupportedVariationList() Dictionary.Any {
	return Dictionary.Any(class(self).GetSupportedVariationList())
}

/*
Returns number of faces in the TrueType / OpenType collection.
*/
func (self Instance) GetFaceCount() int {
	return int(int(class(self).GetFaceCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Font

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Font"))
	casted := Instance{*(*gdclass.Font)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Fallbacks() gd.Array {
	return gd.Array(class(self).GetFallbacks())
}

func (self Instance) SetFallbacks(value gd.Array) {
	class(self).SetFallbacks(value)
}

//go:nosplit
func (self class) SetFallbacks(fallbacks gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(fallbacks))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_set_fallbacks, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbacks() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_fallbacks, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [TextServer] RID of the font cache for specific variation.
*/
//go:nosplit
func (self class) FindVariation(variation_coordinates gd.Dictionary, face_index gd.Int, strength gd.Float, transform gd.Transform2D, spacing_top gd.Int, spacing_bottom gd.Int, spacing_space gd.Int, spacing_glyph gd.Int, baseline_offset gd.Float) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(variation_coordinates))
	callframe.Arg(frame, face_index)
	callframe.Arg(frame, strength)
	callframe.Arg(frame, transform)
	callframe.Arg(frame, spacing_top)
	callframe.Arg(frame, spacing_bottom)
	callframe.Arg(frame, spacing_space)
	callframe.Arg(frame, spacing_glyph)
	callframe.Arg(frame, baseline_offset)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_find_variation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [Array] of valid [Font] [RID]s, which can be passed to the [TextServer] methods.
*/
//go:nosplit
func (self class) GetRids() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_rids, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the total average font height (ascent plus descent) in pixels.
[b]Note:[/b] Real height of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the height of empty line).
*/
//go:nosplit
func (self class) GetHeight(font_size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the average font ascent (number of pixels above the baseline).
[b]Note:[/b] Real ascent of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the ascent of empty line).
*/
//go:nosplit
func (self class) GetAscent(font_size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the average font descent (number of pixels below the baseline).
[b]Note:[/b] Real descent of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate (e.g. as the descent of empty line).
*/
//go:nosplit
func (self class) GetDescent(font_size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns average pixel offset of the underline below the baseline.
[b]Note:[/b] Real underline position of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate.
*/
//go:nosplit
func (self class) GetUnderlinePosition(font_size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns average thickness of the underline.
[b]Note:[/b] Real underline thickness of the string is context-dependent and can be significantly different from the value returned by this function. Use it only as rough estimate.
*/
//go:nosplit
func (self class) GetUnderlineThickness(font_size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns font family name.
*/
//go:nosplit
func (self class) GetFontName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_font_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns font style name.
*/
//go:nosplit
func (self class) GetFontStyleName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_font_style_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
//go:nosplit
func (self class) GetOtNameStrings() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_ot_name_strings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns font style flags, see [enum TextServer.FontStyle].
*/
//go:nosplit
func (self class) GetFontStyle() gdclass.TextServerFontStyle {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerFontStyle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_font_style, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
//go:nosplit
func (self class) GetFontWeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_font_weight, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
//go:nosplit
func (self class) GetFontStretch() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_font_stretch, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the spacing for the given [code]type[/code] (see [enum TextServer.SpacingType]).
*/
//go:nosplit
func (self class) GetSpacing(spacing gdclass.TextServerSpacingType) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a set of OpenType feature tags. More info: [url=https://docs.microsoft.com/en-us/typography/opentype/spec/featuretags]OpenType feature tags[/url].
*/
//go:nosplit
func (self class) GetOpentypeFeatures() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_opentype_features, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets LRU cache capacity for [code]draw_*[/code] methods.
*/
//go:nosplit
func (self class) SetCacheCapacity(single_line gd.Int, multi_line gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, single_line)
	callframe.Arg(frame, multi_line)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_set_cache_capacity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the size of a bounding box of a single-line string, taking kerning, advance and subpixel positioning into account. See also [method get_multiline_string_size] and [method draw_string].
For example, to get the string size as displayed by a single-line Label, use:
[codeblocks]
[gdscript]
var string_size = $Label.get_theme_font("font").get_string_size($Label.text, HORIZONTAL_ALIGNMENT_LEFT, -1, $Label.get_theme_font_size("font_size"))
[/gdscript]
[csharp]
Label label = GetNode<Label>("Label");
Vector2 stringSize = label.GetThemeFont("font").GetStringSize(label.Text, HorizontalAlignment.Left, -1, label.GetThemeFontSize("font_size"));
[/csharp]
[/codeblocks]
[b]Note:[/b] Since kerning, advance and subpixel positioning are taken into account by [method get_string_size], using separate [method get_string_size] calls on substrings of a string then adding the results together will return a different result compared to using a single [method get_string_size] call on the full string.
[b]Note:[/b] Real height of the string is context-dependent and can be significantly different from the value returned by [method get_height].
*/
//go:nosplit
func (self class) GetStringSize(text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_string_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the size of a bounding box of a string broken into the lines, taking kerning and advance into account.
See also [method draw_multiline_string].
*/
//go:nosplit
func (self class) GetMultilineStringSize(text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, brk_flags gdclass.TextServerLineBreakFlag, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_multiline_string_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draw [param text] into a canvas item using the font, at a given position, with [param modulate] color, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_string].
*/
//go:nosplit
func (self class) DrawString(canvas_item gd.RID, pos gd.Vector2, text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, modulate gd.Color, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_draw_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Breaks [param text] into lines using rules specified by [param brk_flags] and draws it into a canvas item using the font, at a given position, with [param modulate] color, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline of the first line, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_multiline_string].
*/
//go:nosplit
func (self class) DrawMultilineString(canvas_item gd.RID, pos gd.Vector2, text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, modulate gd.Color, brk_flags gdclass.TextServerLineBreakFlag, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_draw_multiline_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draw [param text] outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_string_outline].
*/
//go:nosplit
func (self class) DrawStringOutline(canvas_item gd.RID, pos gd.Vector2, text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, size gd.Int, modulate gd.Color, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_draw_string_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Breaks [param text] to the lines using rules specified by [param brk_flags] and draws text outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size, optionally clipping the width and aligning horizontally. [param pos] specifies the baseline of the first line, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
See also [method CanvasItem.draw_multiline_string_outline].
*/
//go:nosplit
func (self class) DrawMultilineStringOutline(canvas_item gd.RID, pos gd.Vector2, text gd.String, alignment HorizontalAlignment, width gd.Float, font_size gd.Int, max_lines gd.Int, size gd.Int, modulate gd.Color, brk_flags gdclass.TextServerLineBreakFlag, justification_flags gdclass.TextServerJustificationFlag, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, width)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, max_lines)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	callframe.Arg(frame, brk_flags)
	callframe.Arg(frame, justification_flags)
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_draw_multiline_string_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the size of a character. Does not take kerning into account.
[b]Note:[/b] Do not use this function to calculate width of the string character by character, use [method get_string_size] or [TextLine] instead. The height returned is the font height (see also [method get_height]) and has no relation to the glyph height.
*/
//go:nosplit
func (self class) GetCharSize(char gd.Int, font_size gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, char)
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_char_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draw a single Unicode character [param char] into a canvas item using the font, at a given position, with [param modulate] color. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
[b]Note:[/b] Do not use this function to draw strings character by character, use [method draw_string] or [TextLine] instead.
*/
//go:nosplit
func (self class) DrawChar(canvas_item gd.RID, pos gd.Vector2, char gd.Int, font_size gd.Int, modulate gd.Color) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, char)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_draw_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draw a single Unicode character [param char] outline into a canvas item using the font, at a given position, with [param modulate] color and [param size] outline size. [param pos] specifies the baseline, not the top. To draw from the top, [i]ascent[/i] must be added to the Y axis.
[b]Note:[/b] Do not use this function to draw strings character by character, use [method draw_string] or [TextLine] instead.
*/
//go:nosplit
func (self class) DrawCharOutline(canvas_item gd.RID, pos gd.Vector2, char gd.Int, font_size gd.Int, size gd.Int, modulate gd.Color) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, char)
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, size)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_draw_char_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
//go:nosplit
func (self class) HasChar(char gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, char)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_has_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a string containing all the characters available in the font.
If a given character is included in more than one font data source, it appears only once in the returned string.
*/
//go:nosplit
func (self class) GetSupportedChars() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_supported_chars, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
//go:nosplit
func (self class) IsLanguageSupported(language gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_is_language_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if font supports given script ([url=https://en.wikipedia.org/wiki/ISO_15924]ISO 15924[/url] code).
*/
//go:nosplit
func (self class) IsScriptSupported(script gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_is_script_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns list of OpenType features supported by font.
*/
//go:nosplit
func (self class) GetSupportedFeatureList() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_supported_feature_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns list of supported [url=https://docs.microsoft.com/en-us/typography/opentype/spec/dvaraxisreg]variation coordinates[/url], each coordinate is returned as [code]tag: Vector3i(min_value,max_value,default_value)[/code].
Font variations allow for continuous change of glyph characteristics along some given design axis, such as weight, width or slant.
To print available variation axes of a variable font:
[codeblock]
var fv = FontVariation.new()
fv.base_font = load("res://RobotoFlex.ttf")
var variation_list = fv.get_supported_variation_list()
for tag in variation_list:
    var name = TextServerManager.get_primary_interface().tag_to_name(tag)
    var values = variation_list[tag]
    print("variation axis: %s (%d)\n\tmin, max, default: %s" % [name, tag, values])
[/codeblock]
[b]Note:[/b] To set and get variation coordinates of a [FontVariation], use [member FontVariation.variation_opentype].
*/
//go:nosplit
func (self class) GetSupportedVariationList() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_supported_variation_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns number of faces in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) GetFaceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Font.Bind_get_face_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFont() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFont() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Font", func(ptr gd.Object) any { return [1]gdclass.Font{*(*gdclass.Font)(unsafe.Pointer(&ptr))} })
}

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
