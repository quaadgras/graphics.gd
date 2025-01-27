// Package TextServerExtension provides methods for working with TextServerExtension object instances.
package TextServerExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/TextServer"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Vector3i"

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

/*
External [TextServer] implementations should inherit from this class.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=TextServerExtension)
*/
type Instance [1]gdclass.TextServerExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTextServerExtension() Instance
}
type Interface interface {
	//[b]Required.[/b]
	//Returns [code]true[/code] if the server supports a feature.
	HasFeature(feature gdclass.TextServerFeature) bool
	//[b]Required.[/b]
	//Returns the name of the server interface.
	GetName() string
	//[b]Required.[/b]
	//Returns text server features, see [enum TextServer.Feature].
	GetFeatures() int
	//[b]Required.[/b]
	//Frees an object created by this [TextServer].
	FreeRid(rid RID.Any)
	//[b]Required.[/b]
	//Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
	Has(rid RID.Any) bool
	//[b]Optional.[/b]
	//Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
	LoadSupportData(filename string) bool
	//[b]Optional.[/b]
	//Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
	GetSupportDataFilename() string
	//[b]Optional.[/b]
	//Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
	GetSupportDataInfo() string
	//[b]Optional.[/b]
	//Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
	SaveSupportData(filename string) bool
	//[b]Required.[/b]
	//Returns [code]true[/code] if locale is right-to-left.
	IsLocaleRightToLeft(locale string) bool
	//[b]Optional.[/b]
	//Converts readable feature, variation, script, or language name to OpenType tag.
	NameToTag(name string) int
	//[b]Optional.[/b]
	//Converts OpenType tag to readable feature, variation, script, or language name.
	TagToName(tag int) string
	//[b]Required.[/b]
	//Creates a new, empty font cache entry resource.
	CreateFont() RID.Any
	//Optional, implement if font supports extra spacing or baseline offset.
	//Creates a new variation existing font which is reusing the same glyph cache and font data.
	CreateFontLinkedVariation(font_rid RID.Any) RID.Any
	//[b]Optional.[/b]
	//Sets font source data, e.g contents of the dynamic font source file.
	FontSetData(font_rid RID.Any, data []byte)
	//[b]Optional.[/b]
	//Sets pointer to the font source data, e.g contents of the dynamic font source file.
	FontSetDataPtr(font_rid RID.Any, data_ptr unsafe.Pointer, data_size int)
	//[b]Optional.[/b]
	//Sets an active face index in the TrueType / OpenType collection.
	FontSetFaceIndex(font_rid RID.Any, face_index int)
	//[b]Optional.[/b]
	//Returns an active face index in the TrueType / OpenType collection.
	FontGetFaceIndex(font_rid RID.Any) int
	//[b]Optional.[/b]
	//Returns number of faces in the TrueType / OpenType collection.
	FontGetFaceCount(font_rid RID.Any) int
	//[b]Optional.[/b]
	//Sets the font style flags, see [enum TextServer.FontStyle].
	FontSetStyle(font_rid RID.Any, style gdclass.TextServerFontStyle)
	//[b]Optional.[/b]
	//Returns font style flags, see [enum TextServer.FontStyle].
	FontGetStyle(font_rid RID.Any) gdclass.TextServerFontStyle
	//[b]Optional.[/b]
	//Sets the font family name.
	FontSetName(font_rid RID.Any, name string)
	//[b]Optional.[/b]
	//Returns font family name.
	FontGetName(font_rid RID.Any) string
	//[b]Optional.[/b]
	//Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
	FontGetOtNameStrings(font_rid RID.Any) map[any]any
	//[b]Optional.[/b]
	//Sets the font style name.
	FontSetStyleName(font_rid RID.Any, name_style string)
	//[b]Optional.[/b]
	//Returns font style name.
	FontGetStyleName(font_rid RID.Any) string
	//[b]Optional.[/b]
	//Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
	FontSetWeight(font_rid RID.Any, weight int)
	//[b]Optional.[/b]
	//Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
	FontGetWeight(font_rid RID.Any) int
	//[b]Optional.[/b]
	//Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
	FontSetStretch(font_rid RID.Any, stretch int)
	//[b]Optional.[/b]
	//Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
	FontGetStretch(font_rid RID.Any) int
	//[b]Optional.[/b]
	//Sets font anti-aliasing mode.
	FontSetAntialiasing(font_rid RID.Any, antialiasing gdclass.TextServerFontAntialiasing)
	//[b]Optional.[/b]
	//Returns font anti-aliasing mode.
	FontGetAntialiasing(font_rid RID.Any) gdclass.TextServerFontAntialiasing
	//[b]Optional.[/b]
	//If set to [code]true[/code], embedded font bitmap loading is disabled.
	FontSetDisableEmbeddedBitmaps(font_rid RID.Any, disable_embedded_bitmaps bool)
	//[b]Optional.[/b]
	//Returns whether the font's embedded bitmap loading is disabled.
	FontGetDisableEmbeddedBitmaps(font_rid RID.Any) bool
	//[b]Optional.[/b]
	//If set to [code]true[/code] font texture mipmap generation is enabled.
	FontSetGenerateMipmaps(font_rid RID.Any, generate_mipmaps bool)
	//[b]Optional.[/b]
	//Returns [code]true[/code] if font texture mipmap generation is enabled.
	FontGetGenerateMipmaps(font_rid RID.Any) bool
	//[b]Optional.[/b]
	//If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
	FontSetMultichannelSignedDistanceField(font_rid RID.Any, msdf bool)
	//[b]Optional.[/b]
	//Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
	FontIsMultichannelSignedDistanceField(font_rid RID.Any) bool
	//[b]Optional.[/b]
	//Sets the width of the range around the shape between the minimum and maximum representable signed distance.
	FontSetMsdfPixelRange(font_rid RID.Any, msdf_pixel_range int)
	//[b]Optional.[/b]
	//Returns the width of the range around the shape between the minimum and maximum representable signed distance.
	FontGetMsdfPixelRange(font_rid RID.Any) int
	//[b]Optional.[/b]
	//Sets source font size used to generate MSDF textures.
	FontSetMsdfSize(font_rid RID.Any, msdf_size int)
	//[b]Optional.[/b]
	//Returns source font size used to generate MSDF textures.
	FontGetMsdfSize(font_rid RID.Any) int
	//[b]Required.[/b]
	//Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
	FontSetFixedSize(font_rid RID.Any, fixed_size int)
	//[b]Required.[/b]
	//Returns bitmap font fixed size.
	FontGetFixedSize(font_rid RID.Any) int
	//[b]Required.[/b]
	//Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
	FontSetFixedSizeScaleMode(font_rid RID.Any, fixed_size_scale_mode gdclass.TextServerFixedSizeScaleMode)
	//[b]Required.[/b]
	//Returns bitmap font scaling mode.
	FontGetFixedSizeScaleMode(font_rid RID.Any) gdclass.TextServerFixedSizeScaleMode
	//[b]Optional.[/b]
	//If set to [code]true[/code], system fonts can be automatically used as fallbacks.
	FontSetAllowSystemFallback(font_rid RID.Any, allow_system_fallback bool)
	//[b]Optional.[/b]
	//Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
	FontIsAllowSystemFallback(font_rid RID.Any) bool
	//[b]Optional.[/b]
	//If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
	FontSetForceAutohinter(font_rid RID.Any, force_autohinter bool)
	//[b]Optional.[/b]
	//Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting.
	FontIsForceAutohinter(font_rid RID.Any) bool
	//[b]Optional.[/b]
	//Sets font hinting mode. Used by dynamic fonts only.
	FontSetHinting(font_rid RID.Any, hinting gdclass.TextServerHinting)
	//[b]Optional.[/b]
	//Returns the font hinting mode. Used by dynamic fonts only.
	FontGetHinting(font_rid RID.Any) gdclass.TextServerHinting
	//[b]Optional.[/b]
	//Sets font subpixel glyph positioning mode.
	FontSetSubpixelPositioning(font_rid RID.Any, subpixel_positioning gdclass.TextServerSubpixelPositioning)
	//[b]Optional.[/b]
	//Returns font subpixel glyph positioning mode.
	FontGetSubpixelPositioning(font_rid RID.Any) gdclass.TextServerSubpixelPositioning
	//Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
	FontSetEmbolden(font_rid RID.Any, strength Float.X)
	//[b]Optional.[/b]
	//Returns font embolden strength.
	FontGetEmbolden(font_rid RID.Any) Float.X
	//[b]Optional.[/b]
	//Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
	FontSetSpacing(font_rid RID.Any, spacing gdclass.TextServerSpacingType, value int)
	//[b]Optional.[/b]
	//Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
	FontGetSpacing(font_rid RID.Any, spacing gdclass.TextServerSpacingType) int
	//[b]Optional.[/b]
	//Sets extra baseline offset (as a fraction of font height).
	FontSetBaselineOffset(font_rid RID.Any, baseline_offset Float.X)
	//[b]Optional.[/b]
	//Returns extra baseline offset (as a fraction of font height).
	FontGetBaselineOffset(font_rid RID.Any) Float.X
	//[b]Optional.[/b]
	//Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
	FontSetTransform(font_rid RID.Any, transform Transform2D.OriginXY)
	//[b]Optional.[/b]
	//Returns 2D transform applied to the font outlines.
	FontGetTransform(font_rid RID.Any) Transform2D.OriginXY
	//[b]Optional.[/b]
	//Sets variation coordinates for the specified font cache entry.
	FontSetVariationCoordinates(font_rid RID.Any, variation_coordinates map[any]any)
	//[b]Optional.[/b]
	//Returns variation coordinates for the specified font cache entry.
	FontGetVariationCoordinates(font_rid RID.Any) map[any]any
	//[b]Optional.[/b]
	//Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
	FontSetOversampling(font_rid RID.Any, oversampling Float.X)
	//[b]Optional.[/b]
	//Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
	FontGetOversampling(font_rid RID.Any) Float.X
	//[b]Required.[/b]
	//Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
	FontGetSizeCacheList(font_rid RID.Any) []Vector2i.XY
	//[b]Required.[/b]
	//Removes all font sizes from the cache entry.
	FontClearSizeCache(font_rid RID.Any)
	//[b]Required.[/b]
	//Removes specified font size from the cache entry.
	FontRemoveSizeCache(font_rid RID.Any, size Vector2i.XY)
	//[b]Required.[/b]
	//Sets the font ascent (number of pixels above the baseline).
	FontSetAscent(font_rid RID.Any, size int, ascent Float.X)
	//[b]Required.[/b]
	//Returns the font ascent (number of pixels above the baseline).
	FontGetAscent(font_rid RID.Any, size int) Float.X
	//[b]Required.[/b]
	//Sets the font descent (number of pixels below the baseline).
	FontSetDescent(font_rid RID.Any, size int, descent Float.X)
	//[b]Required.[/b]
	//Returns the font descent (number of pixels below the baseline).
	FontGetDescent(font_rid RID.Any, size int) Float.X
	//[b]Required.[/b]
	//Sets pixel offset of the underline below the baseline.
	FontSetUnderlinePosition(font_rid RID.Any, size int, underline_position Float.X)
	//[b]Required.[/b]
	//Returns pixel offset of the underline below the baseline.
	FontGetUnderlinePosition(font_rid RID.Any, size int) Float.X
	//[b]Required.[/b]
	//Sets thickness of the underline in pixels.
	FontSetUnderlineThickness(font_rid RID.Any, size int, underline_thickness Float.X)
	//[b]Required.[/b]
	//Returns thickness of the underline in pixels.
	FontGetUnderlineThickness(font_rid RID.Any, size int) Float.X
	//[b]Required.[/b]
	//Sets scaling factor of the color bitmap font.
	FontSetScale(font_rid RID.Any, size int, scale Float.X)
	//[b]Required.[/b]
	//Returns scaling factor of the color bitmap font.
	FontGetScale(font_rid RID.Any, size int) Float.X
	//[b]Required.[/b]
	//Returns number of textures used by font cache entry.
	FontGetTextureCount(font_rid RID.Any, size Vector2i.XY) int
	//[b]Required.[/b]
	//Removes all textures from font cache entry.
	FontClearTextures(font_rid RID.Any, size Vector2i.XY)
	//[b]Required.[/b]
	//Removes specified texture from the cache entry.
	FontRemoveTexture(font_rid RID.Any, size Vector2i.XY, texture_index int)
	//[b]Required.[/b]
	//Sets font cache texture image data.
	FontSetTextureImage(font_rid RID.Any, size Vector2i.XY, texture_index int, image [1]gdclass.Image)
	//[b]Required.[/b]
	//Returns font cache texture image data.
	FontGetTextureImage(font_rid RID.Any, size Vector2i.XY, texture_index int) [1]gdclass.Image
	//[b]Optional.[/b]
	//Sets array containing glyph packing data.
	FontSetTextureOffsets(font_rid RID.Any, size Vector2i.XY, texture_index int, offset []int32)
	//[b]Optional.[/b]
	//Returns array containing glyph packing data.
	FontGetTextureOffsets(font_rid RID.Any, size Vector2i.XY, texture_index int) []int32
	//[b]Required.[/b]
	//Returns list of rendered glyphs in the cache entry.
	FontGetGlyphList(font_rid RID.Any, size Vector2i.XY) []int32
	//[b]Required.[/b]
	//Removes all rendered glyph information from the cache entry.
	FontClearGlyphs(font_rid RID.Any, size Vector2i.XY)
	//[b]Required.[/b]
	//Removes specified rendered glyph information from the cache entry.
	FontRemoveGlyph(font_rid RID.Any, size Vector2i.XY, glyph int)
	//[b]Required.[/b]
	//Returns glyph advance (offset of the next glyph).
	FontGetGlyphAdvance(font_rid RID.Any, size int, glyph int) Vector2.XY
	//[b]Required.[/b]
	//Sets glyph advance (offset of the next glyph).
	FontSetGlyphAdvance(font_rid RID.Any, size int, glyph int, advance Vector2.XY)
	//[b]Required.[/b]
	//Returns glyph offset from the baseline.
	FontGetGlyphOffset(font_rid RID.Any, size Vector2i.XY, glyph int) Vector2.XY
	//[b]Required.[/b]
	//Sets glyph offset from the baseline.
	FontSetGlyphOffset(font_rid RID.Any, size Vector2i.XY, glyph int, offset Vector2.XY)
	//[b]Required.[/b]
	//Returns size of the glyph.
	FontGetGlyphSize(font_rid RID.Any, size Vector2i.XY, glyph int) Vector2.XY
	//[b]Required.[/b]
	//Sets size of the glyph.
	FontSetGlyphSize(font_rid RID.Any, size Vector2i.XY, glyph int, gl_size Vector2.XY)
	//[b]Required.[/b]
	//Returns rectangle in the cache texture containing the glyph.
	FontGetGlyphUvRect(font_rid RID.Any, size Vector2i.XY, glyph int) Rect2.PositionSize
	//[b]Required.[/b]
	//Sets rectangle in the cache texture containing the glyph.
	FontSetGlyphUvRect(font_rid RID.Any, size Vector2i.XY, glyph int, uv_rect Rect2.PositionSize)
	//[b]Required.[/b]
	//Returns index of the cache texture containing the glyph.
	FontGetGlyphTextureIdx(font_rid RID.Any, size Vector2i.XY, glyph int) int
	//[b]Required.[/b]
	//Sets index of the cache texture containing the glyph.
	FontSetGlyphTextureIdx(font_rid RID.Any, size Vector2i.XY, glyph int, texture_idx int)
	//[b]Required.[/b]
	//Returns resource ID of the cache texture containing the glyph.
	FontGetGlyphTextureRid(font_rid RID.Any, size Vector2i.XY, glyph int) RID.Any
	//[b]Required.[/b]
	//Returns size of the cache texture containing the glyph.
	FontGetGlyphTextureSize(font_rid RID.Any, size Vector2i.XY, glyph int) Vector2.XY
	//[b]Optional.[/b]
	//Returns outline contours of the glyph.
	FontGetGlyphContours(font_rid RID.Any, size int, index int) map[any]any
	//[b]Optional.[/b]
	//Returns list of the kerning overrides.
	FontGetKerningList(font_rid RID.Any, size int) []Vector2i.XY
	//[b]Optional.[/b]
	//Removes all kerning overrides.
	FontClearKerningMap(font_rid RID.Any, size int)
	//[b]Optional.[/b]
	//Removes kerning override for the pair of glyphs.
	FontRemoveKerning(font_rid RID.Any, size int, glyph_pair Vector2i.XY)
	//[b]Optional.[/b]
	//Sets kerning for the pair of glyphs.
	FontSetKerning(font_rid RID.Any, size int, glyph_pair Vector2i.XY, kerning Vector2.XY)
	//[b]Optional.[/b]
	//Returns kerning for the pair of glyphs.
	FontGetKerning(font_rid RID.Any, size int, glyph_pair Vector2i.XY) Vector2.XY
	//[b]Required.[/b]
	//Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
	FontGetGlyphIndex(font_rid RID.Any, size int, char int, variation_selector int) int
	//[b]Required.[/b]
	//Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid.
	FontGetCharFromGlyphIndex(font_rid RID.Any, size int, glyph_index int) int
	//[b]Required.[/b]
	//Returns [code]true[/code] if a Unicode [param char] is available in the font.
	FontHasChar(font_rid RID.Any, char int) bool
	//[b]Required.[/b]
	//Returns a string containing all the characters available in the font.
	FontGetSupportedChars(font_rid RID.Any) string
	//[b]Optional.[/b]
	//Renders the range of characters to the font cache texture.
	FontRenderRange(font_rid RID.Any, size Vector2i.XY, start int, end int)
	//[b]Optional.[/b]
	//Renders specified glyph to the font cache texture.
	FontRenderGlyph(font_rid RID.Any, size Vector2i.XY, index int)
	//[b]Required.[/b]
	//Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
	FontDrawGlyph(font_rid RID.Any, canvas RID.Any, size int, pos Vector2.XY, index int, color Color.RGBA)
	//[b]Required.[/b]
	//Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
	FontDrawGlyphOutline(font_rid RID.Any, canvas RID.Any, size int, outline_size int, pos Vector2.XY, index int, color Color.RGBA)
	//[b]Optional.[/b]
	//Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
	FontIsLanguageSupported(font_rid RID.Any, language string) bool
	//[b]Optional.[/b]
	//Adds override for [method _font_is_language_supported].
	FontSetLanguageSupportOverride(font_rid RID.Any, language string, supported bool)
	//[b]Optional.[/b]
	//Returns [code]true[/code] if support override is enabled for the [param language].
	FontGetLanguageSupportOverride(font_rid RID.Any, language string) bool
	//[b]Optional.[/b]
	//Remove language support override.
	FontRemoveLanguageSupportOverride(font_rid RID.Any, language string)
	//[b]Optional.[/b]
	//Returns list of language support overrides.
	FontGetLanguageSupportOverrides(font_rid RID.Any) []string
	//[b]Optional.[/b]
	//Returns [code]true[/code], if font supports given script (ISO 15924 code).
	FontIsScriptSupported(font_rid RID.Any, script string) bool
	//[b]Optional.[/b]
	//Adds override for [method _font_is_script_supported].
	FontSetScriptSupportOverride(font_rid RID.Any, script string, supported bool)
	//[b]Optional.[/b]
	//Returns [code]true[/code] if support override is enabled for the [param script].
	FontGetScriptSupportOverride(font_rid RID.Any, script string) bool
	//[b]Optional.[/b]
	//Removes script support override.
	FontRemoveScriptSupportOverride(font_rid RID.Any, script string)
	//[b]Optional.[/b]
	//Returns list of script support overrides.
	FontGetScriptSupportOverrides(font_rid RID.Any) []string
	//[b]Optional.[/b]
	//Sets font OpenType feature set override.
	FontSetOpentypeFeatureOverrides(font_rid RID.Any, overrides map[any]any)
	//[b]Optional.[/b]
	//Returns font OpenType feature set override.
	FontGetOpentypeFeatureOverrides(font_rid RID.Any) map[any]any
	//[b]Optional.[/b]
	//Returns the dictionary of the supported OpenType features.
	FontSupportedFeatureList(font_rid RID.Any) map[any]any
	//[b]Optional.[/b]
	//Returns the dictionary of the supported OpenType variation coordinates.
	FontSupportedVariationList(font_rid RID.Any) map[any]any
	//[b]Optional.[/b]
	//Returns the font oversampling factor, shared by all fonts in the TextServer.
	FontGetGlobalOversampling() Float.X
	//[b]Optional.[/b]
	//Sets oversampling factor, shared by all font in the TextServer.
	FontSetGlobalOversampling(oversampling Float.X)
	//[b]Optional.[/b]
	//Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
	GetHexCodeBoxSize(size int, index int) Vector2.XY
	//[b]Optional.[/b]
	//Draws box displaying character hexadecimal code.
	DrawHexCodeBox(canvas RID.Any, size int, pos Vector2.XY, index int, color Color.RGBA)
	//[b]Required.[/b]
	//Creates a new buffer for complex text layout, with the given [param direction] and [param orientation].
	CreateShapedText(direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) RID.Any
	//[b]Required.[/b]
	//Clears text buffer (removes text and inline objects).
	ShapedTextClear(shaped RID.Any)
	//[b]Optional.[/b]
	//Sets desired text direction. If set to [constant TextServer.DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
	ShapedTextSetDirection(shaped RID.Any, direction gdclass.TextServerDirection)
	//[b]Optional.[/b]
	//Returns direction of the text.
	ShapedTextGetDirection(shaped RID.Any) gdclass.TextServerDirection
	//[b]Optional.[/b]
	//Returns direction of the text, inferred by the BiDi algorithm.
	ShapedTextGetInferredDirection(shaped RID.Any) gdclass.TextServerDirection
	//[b]Optional.[/b]
	//Overrides BiDi for the structured text.
	ShapedTextSetBidiOverride(shaped RID.Any, override []any)
	//[b]Optional.[/b]
	//Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
	ShapedTextSetCustomPunctuation(shaped RID.Any, punct string)
	//[b]Optional.[/b]
	//Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
	ShapedTextGetCustomPunctuation(shaped RID.Any) string
	//[b]Optional.[/b]
	//Sets ellipsis character used for text clipping.
	ShapedTextSetCustomEllipsis(shaped RID.Any, char int)
	//[b]Optional.[/b]
	//Returns ellipsis character used for text clipping.
	ShapedTextGetCustomEllipsis(shaped RID.Any) int
	//[b]Optional.[/b]
	//Sets desired text orientation.
	ShapedTextSetOrientation(shaped RID.Any, orientation gdclass.TextServerOrientation)
	//[b]Optional.[/b]
	//Returns text orientation.
	ShapedTextGetOrientation(shaped RID.Any) gdclass.TextServerOrientation
	//[b]Optional.[/b]
	//If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
	ShapedTextSetPreserveInvalid(shaped RID.Any, enabled bool)
	//[b]Optional.[/b]
	//Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
	ShapedTextGetPreserveInvalid(shaped RID.Any) bool
	//[b]Optional.[/b]
	//If set to [code]true[/code] text buffer will display control characters.
	ShapedTextSetPreserveControl(shaped RID.Any, enabled bool)
	//[b]Optional.[/b]
	//Returns [code]true[/code] if text buffer is configured to display control characters.
	ShapedTextGetPreserveControl(shaped RID.Any) bool
	//[b]Optional.[/b]
	//Sets extra spacing added between glyphs or lines in pixels.
	ShapedTextSetSpacing(shaped RID.Any, spacing gdclass.TextServerSpacingType, value int)
	//[b]Optional.[/b]
	//Returns extra spacing added between glyphs or lines in pixels.
	ShapedTextGetSpacing(shaped RID.Any, spacing gdclass.TextServerSpacingType) int
	//[b]Required.[/b]
	//Adds text span and font to draw it to the text buffer.
	ShapedTextAddString(shaped RID.Any, text string, fonts []RID.Any, size int, opentype_features map[any]any, language string, meta any) bool
	//[b]Required.[/b]
	//Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
	ShapedTextAddObject(shaped RID.Any, key any, size Vector2.XY, inline_align InlineAlignment, length int, baseline Float.X) bool
	//[b]Required.[/b]
	//Sets new size and alignment of embedded object.
	ShapedTextResizeObject(shaped RID.Any, key any, size Vector2.XY, inline_align InlineAlignment, baseline Float.X) bool
	//[b]Required.[/b]
	//Returns number of text spans added using [method _shaped_text_add_string] or [method _shaped_text_add_object].
	ShapedGetSpanCount(shaped RID.Any) int
	//[b]Required.[/b]
	//Returns text span metadata.
	ShapedGetSpanMeta(shaped RID.Any, index int) any
	//[b]Required.[/b]
	//Changes text span font, font size, and OpenType features, without changing the text.
	ShapedSetSpanUpdateFont(shaped RID.Any, index int, fonts []RID.Any, size int, opentype_features map[any]any)
	//[b]Required.[/b]
	//Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
	ShapedTextSubstr(shaped RID.Any, start int, length int) RID.Any
	//[b]Required.[/b]
	//Returns the parent buffer from which the substring originates.
	ShapedTextGetParent(shaped RID.Any) RID.Any
	//[b]Optional.[/b]
	//Adjusts text width to fit to specified width, returns new text width.
	ShapedTextFitToWidth(shaped RID.Any, width Float.X, justification_flags gdclass.TextServerJustificationFlag) Float.X
	//[b]Optional.[/b]
	//Aligns shaped text to the given tab-stops.
	ShapedTextTabAlign(shaped RID.Any, tab_stops []float32) Float.X
	//[b]Required.[/b]
	//Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
	ShapedTextShape(shaped RID.Any) bool
	//[b]Optional.[/b]
	//Updates break points in the shaped text. This method is called by default implementation of text breaking functions.
	ShapedTextUpdateBreaks(shaped RID.Any) bool
	//[b]Optional.[/b]
	//Updates justification points in the shaped text. This method is called by default implementation of text justification functions.
	ShapedTextUpdateJustificationOps(shaped RID.Any) bool
	//[b]Required.[/b]
	//Returns [code]true[/code] if buffer is successfully shaped.
	ShapedTextIsReady(shaped RID.Any) bool
	//[b]Required.[/b]
	//Returns an array of glyphs in the visual order.
	ShapedTextGetGlyphs(shaped RID.Any) *Glyph
	//[b]Required.[/b]
	//Returns text glyphs in the logical order.
	ShapedTextSortLogical(shaped RID.Any) *Glyph
	//[b]Required.[/b]
	//Returns number of glyphs in the buffer.
	ShapedTextGetGlyphCount(shaped RID.Any) int
	//[b]Required.[/b]
	//Returns substring buffer character range in the parent buffer.
	ShapedTextGetRange(shaped RID.Any) Vector2i.XY
	//[b]Optional.[/b]
	//Breaks text to the lines and columns. Returns character ranges for each segment.
	ShapedTextGetLineBreaksAdv(shaped RID.Any, width []float32, start int, once bool, break_flags gdclass.TextServerLineBreakFlag) []int32
	//[b]Optional.[/b]
	//Breaks text to the lines and returns character ranges for each line.
	ShapedTextGetLineBreaks(shaped RID.Any, width Float.X, start int, break_flags gdclass.TextServerLineBreakFlag) []int32
	//[b]Optional.[/b]
	//Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum TextServer.GraphemeFlag]).
	ShapedTextGetWordBreaks(shaped RID.Any, grapheme_flags gdclass.TextServerGraphemeFlag, skip_grapheme_flags gdclass.TextServerGraphemeFlag) []int32
	//[b]Required.[/b]
	//Returns the position of the overrun trim.
	ShapedTextGetTrimPos(shaped RID.Any) int
	//[b]Required.[/b]
	//Returns position of the ellipsis.
	ShapedTextGetEllipsisPos(shaped RID.Any) int
	//[b]Required.[/b]
	//Returns number of glyphs in the ellipsis.
	ShapedTextGetEllipsisGlyphCount(shaped RID.Any) int
	//[b]Required.[/b]
	//Returns array of the glyphs in the ellipsis.
	ShapedTextGetEllipsisGlyphs(shaped RID.Any) *Glyph
	//[b]Optional.[/b]
	//Trims text if it exceeds the given width.
	ShapedTextOverrunTrimToWidth(shaped RID.Any, width Float.X, trim_flags gdclass.TextServerTextOverrunFlag)
	//[b]Required.[/b]
	//Returns array of inline objects.
	ShapedTextGetObjects(shaped RID.Any) []any
	//[b]Required.[/b]
	//Returns bounding rectangle of the inline object.
	ShapedTextGetObjectRect(shaped RID.Any, key any) Rect2.PositionSize
	//[b]Required.[/b]
	//Returns the character range of the inline object.
	ShapedTextGetObjectRange(shaped RID.Any, key any) Vector2i.XY
	//[b]Required.[/b]
	//Returns the glyph index of the inline object.
	ShapedTextGetObjectGlyph(shaped RID.Any, key any) int
	//[b]Required.[/b]
	//Returns size of the text.
	ShapedTextGetSize(shaped RID.Any) Vector2.XY
	//[b]Required.[/b]
	//Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
	ShapedTextGetAscent(shaped RID.Any) Float.X
	//[b]Required.[/b]
	//Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
	ShapedTextGetDescent(shaped RID.Any) Float.X
	//[b]Required.[/b]
	//Returns width (for horizontal layout) or height (for vertical) of the text.
	ShapedTextGetWidth(shaped RID.Any) Float.X
	//[b]Required.[/b]
	//Returns pixel offset of the underline below the baseline.
	ShapedTextGetUnderlinePosition(shaped RID.Any) Float.X
	//[b]Required.[/b]
	//Returns thickness of the underline.
	ShapedTextGetUnderlineThickness(shaped RID.Any) Float.X
	//[b]Optional.[/b]
	//Returns dominant direction of in the range of text.
	ShapedTextGetDominantDirectionInRange(shaped RID.Any, start int, end int) int
	//[b]Optional.[/b]
	//Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
	ShapedTextGetCarets(shaped RID.Any, position int, caret *CaretInfo)
	//[b]Optional.[/b]
	//Returns selection rectangles for the specified character range.
	ShapedTextGetSelection(shaped RID.Any, start int, end int) []Vector2.XY
	//[b]Optional.[/b]
	//Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
	ShapedTextHitTestGrapheme(shaped RID.Any, coord Float.X) int
	//[b]Optional.[/b]
	//Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
	ShapedTextHitTestPosition(shaped RID.Any, coord Float.X) int
	//[b]Optional.[/b]
	//Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
	ShapedTextDraw(shaped RID.Any, canvas RID.Any, pos Vector2.XY, clip_l Float.X, clip_r Float.X, color Color.RGBA)
	//[b]Optional.[/b]
	//Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
	ShapedTextDrawOutline(shaped RID.Any, canvas RID.Any, pos Vector2.XY, clip_l Float.X, clip_r Float.X, outline_size int, color Color.RGBA)
	//[b]Optional.[/b]
	//Returns composite character's bounds as offsets from the start of the line.
	ShapedTextGetGraphemeBounds(shaped RID.Any, pos int) Vector2.XY
	//[b]Optional.[/b]
	//Returns grapheme end position closest to the [param pos].
	ShapedTextNextGraphemePos(shaped RID.Any, pos int) int
	//[b]Optional.[/b]
	//Returns grapheme start position closest to the [param pos].
	ShapedTextPrevGraphemePos(shaped RID.Any, pos int) int
	//[b]Optional.[/b]
	//Returns array of the composite character boundaries.
	ShapedTextGetCharacterBreaks(shaped RID.Any) []int32
	//[b]Optional.[/b]
	//Returns composite character end position closest to the [param pos].
	ShapedTextNextCharacterPos(shaped RID.Any, pos int) int
	//[b]Optional.[/b]
	//Returns composite character start position closest to the [param pos].
	ShapedTextPrevCharacterPos(shaped RID.Any, pos int) int
	//[b]Optional.[/b]
	//Returns composite character position closest to the [param pos].
	ShapedTextClosestCharacterPos(shaped RID.Any, pos int) int
	//[b]Optional.[/b]
	//Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
	FormatNumber(number string, language string) string
	//[b]Optional.[/b]
	//Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
	ParseNumber(number string, language string) string
	//[b]Optional.[/b]
	//Returns percent sign used in the [param language].
	PercentSign(language string) string
	//[b]Optional.[/b]
	//Strips diacritics from the string.
	StripDiacritics(s string) string
	//[b]Optional.[/b]
	//Returns [code]true[/code] if [param string] is a valid identifier.
	IsValidIdentifier(s string) bool
	IsValidLetter(unicode int) bool
	//[b]Optional.[/b]
	//Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
	StringGetWordBreaks(s string, language string, chars_per_line int) []int32
	//[b]Optional.[/b]
	//Returns array of the composite character boundaries.
	StringGetCharacterBreaks(s string, language string) []int32
	//[b]Optional.[/b]
	//Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
	IsConfusable(s string, dict []string) int
	//[b]Optional.[/b]
	//Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
	SpoofCheck(s string) bool
	//[b]Optional.[/b]
	//Returns the string converted to uppercase.
	StringToUpper(s string, language string) string
	//[b]Optional.[/b]
	//Returns the string converted to lowercase.
	StringToLower(s string, language string) string
	//[b]Optional.[/b]
	//Returns the string converted to title case.
	StringToTitle(s string, language string) string
	//[b]Optional.[/b]
	//Default implementation of the BiDi algorithm override function. See [enum TextServer.StructuredTextParser] for more info.
	ParseStructuredText(parser_type gdclass.TextServerStructuredTextParser, args []any, text string) []Vector3i.XYZ
	//[b]Optional.[/b]
	//This method is called before text server is unregistered.
	Cleanup()
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) HasFeature(feature gdclass.TextServerFeature) (_ bool)  { return }
func (self implementation) GetName() (_ string)                                    { return }
func (self implementation) GetFeatures() (_ int)                                   { return }
func (self implementation) FreeRid(rid RID.Any)                                    { return }
func (self implementation) Has(rid RID.Any) (_ bool)                               { return }
func (self implementation) LoadSupportData(filename string) (_ bool)               { return }
func (self implementation) GetSupportDataFilename() (_ string)                     { return }
func (self implementation) GetSupportDataInfo() (_ string)                         { return }
func (self implementation) SaveSupportData(filename string) (_ bool)               { return }
func (self implementation) IsLocaleRightToLeft(locale string) (_ bool)             { return }
func (self implementation) NameToTag(name string) (_ int)                          { return }
func (self implementation) TagToName(tag int) (_ string)                           { return }
func (self implementation) CreateFont() (_ RID.Any)                                { return }
func (self implementation) CreateFontLinkedVariation(font_rid RID.Any) (_ RID.Any) { return }
func (self implementation) FontSetData(font_rid RID.Any, data []byte)              { return }
func (self implementation) FontSetDataPtr(font_rid RID.Any, data_ptr unsafe.Pointer, data_size int) {
	return
}
func (self implementation) FontSetFaceIndex(font_rid RID.Any, face_index int)                { return }
func (self implementation) FontGetFaceIndex(font_rid RID.Any) (_ int)                        { return }
func (self implementation) FontGetFaceCount(font_rid RID.Any) (_ int)                        { return }
func (self implementation) FontSetStyle(font_rid RID.Any, style gdclass.TextServerFontStyle) { return }
func (self implementation) FontGetStyle(font_rid RID.Any) (_ gdclass.TextServerFontStyle)    { return }
func (self implementation) FontSetName(font_rid RID.Any, name string)                        { return }
func (self implementation) FontGetName(font_rid RID.Any) (_ string)                          { return }
func (self implementation) FontGetOtNameStrings(font_rid RID.Any) (_ map[any]any)            { return }
func (self implementation) FontSetStyleName(font_rid RID.Any, name_style string)             { return }
func (self implementation) FontGetStyleName(font_rid RID.Any) (_ string)                     { return }
func (self implementation) FontSetWeight(font_rid RID.Any, weight int)                       { return }
func (self implementation) FontGetWeight(font_rid RID.Any) (_ int)                           { return }
func (self implementation) FontSetStretch(font_rid RID.Any, stretch int)                     { return }
func (self implementation) FontGetStretch(font_rid RID.Any) (_ int)                          { return }
func (self implementation) FontSetAntialiasing(font_rid RID.Any, antialiasing gdclass.TextServerFontAntialiasing) {
	return
}
func (self implementation) FontGetAntialiasing(font_rid RID.Any) (_ gdclass.TextServerFontAntialiasing) {
	return
}
func (self implementation) FontSetDisableEmbeddedBitmaps(font_rid RID.Any, disable_embedded_bitmaps bool) {
	return
}
func (self implementation) FontGetDisableEmbeddedBitmaps(font_rid RID.Any) (_ bool)        { return }
func (self implementation) FontSetGenerateMipmaps(font_rid RID.Any, generate_mipmaps bool) { return }
func (self implementation) FontGetGenerateMipmaps(font_rid RID.Any) (_ bool)               { return }
func (self implementation) FontSetMultichannelSignedDistanceField(font_rid RID.Any, msdf bool) {
	return
}
func (self implementation) FontIsMultichannelSignedDistanceField(font_rid RID.Any) (_ bool) { return }
func (self implementation) FontSetMsdfPixelRange(font_rid RID.Any, msdf_pixel_range int)    { return }
func (self implementation) FontGetMsdfPixelRange(font_rid RID.Any) (_ int)                  { return }
func (self implementation) FontSetMsdfSize(font_rid RID.Any, msdf_size int)                 { return }
func (self implementation) FontGetMsdfSize(font_rid RID.Any) (_ int)                        { return }
func (self implementation) FontSetFixedSize(font_rid RID.Any, fixed_size int)               { return }
func (self implementation) FontGetFixedSize(font_rid RID.Any) (_ int)                       { return }
func (self implementation) FontSetFixedSizeScaleMode(font_rid RID.Any, fixed_size_scale_mode gdclass.TextServerFixedSizeScaleMode) {
	return
}
func (self implementation) FontGetFixedSizeScaleMode(font_rid RID.Any) (_ gdclass.TextServerFixedSizeScaleMode) {
	return
}
func (self implementation) FontSetAllowSystemFallback(font_rid RID.Any, allow_system_fallback bool) {
	return
}
func (self implementation) FontIsAllowSystemFallback(font_rid RID.Any) (_ bool)            { return }
func (self implementation) FontSetForceAutohinter(font_rid RID.Any, force_autohinter bool) { return }
func (self implementation) FontIsForceAutohinter(font_rid RID.Any) (_ bool)                { return }
func (self implementation) FontSetHinting(font_rid RID.Any, hinting gdclass.TextServerHinting) {
	return
}
func (self implementation) FontGetHinting(font_rid RID.Any) (_ gdclass.TextServerHinting) { return }
func (self implementation) FontSetSubpixelPositioning(font_rid RID.Any, subpixel_positioning gdclass.TextServerSubpixelPositioning) {
	return
}
func (self implementation) FontGetSubpixelPositioning(font_rid RID.Any) (_ gdclass.TextServerSubpixelPositioning) {
	return
}
func (self implementation) FontSetEmbolden(font_rid RID.Any, strength Float.X) { return }
func (self implementation) FontGetEmbolden(font_rid RID.Any) (_ Float.X)       { return }
func (self implementation) FontSetSpacing(font_rid RID.Any, spacing gdclass.TextServerSpacingType, value int) {
	return
}
func (self implementation) FontGetSpacing(font_rid RID.Any, spacing gdclass.TextServerSpacingType) (_ int) {
	return
}
func (self implementation) FontSetBaselineOffset(font_rid RID.Any, baseline_offset Float.X)   { return }
func (self implementation) FontGetBaselineOffset(font_rid RID.Any) (_ Float.X)                { return }
func (self implementation) FontSetTransform(font_rid RID.Any, transform Transform2D.OriginXY) { return }
func (self implementation) FontGetTransform(font_rid RID.Any) (_ Transform2D.OriginXY)        { return }
func (self implementation) FontSetVariationCoordinates(font_rid RID.Any, variation_coordinates map[any]any) {
	return
}
func (self implementation) FontGetVariationCoordinates(font_rid RID.Any) (_ map[any]any) { return }
func (self implementation) FontSetOversampling(font_rid RID.Any, oversampling Float.X)   { return }
func (self implementation) FontGetOversampling(font_rid RID.Any) (_ Float.X)             { return }
func (self implementation) FontGetSizeCacheList(font_rid RID.Any) (_ []Vector2i.XY)      { return }
func (self implementation) FontClearSizeCache(font_rid RID.Any)                          { return }
func (self implementation) FontRemoveSizeCache(font_rid RID.Any, size Vector2i.XY)       { return }
func (self implementation) FontSetAscent(font_rid RID.Any, size int, ascent Float.X)     { return }
func (self implementation) FontGetAscent(font_rid RID.Any, size int) (_ Float.X)         { return }
func (self implementation) FontSetDescent(font_rid RID.Any, size int, descent Float.X)   { return }
func (self implementation) FontGetDescent(font_rid RID.Any, size int) (_ Float.X)        { return }
func (self implementation) FontSetUnderlinePosition(font_rid RID.Any, size int, underline_position Float.X) {
	return
}
func (self implementation) FontGetUnderlinePosition(font_rid RID.Any, size int) (_ Float.X) { return }
func (self implementation) FontSetUnderlineThickness(font_rid RID.Any, size int, underline_thickness Float.X) {
	return
}
func (self implementation) FontGetUnderlineThickness(font_rid RID.Any, size int) (_ Float.X) { return }
func (self implementation) FontSetScale(font_rid RID.Any, size int, scale Float.X)           { return }
func (self implementation) FontGetScale(font_rid RID.Any, size int) (_ Float.X)              { return }
func (self implementation) FontGetTextureCount(font_rid RID.Any, size Vector2i.XY) (_ int)   { return }
func (self implementation) FontClearTextures(font_rid RID.Any, size Vector2i.XY)             { return }
func (self implementation) FontRemoveTexture(font_rid RID.Any, size Vector2i.XY, texture_index int) {
	return
}
func (self implementation) FontSetTextureImage(font_rid RID.Any, size Vector2i.XY, texture_index int, image [1]gdclass.Image) {
	return
}
func (self implementation) FontGetTextureImage(font_rid RID.Any, size Vector2i.XY, texture_index int) (_ [1]gdclass.Image) {
	return
}
func (self implementation) FontSetTextureOffsets(font_rid RID.Any, size Vector2i.XY, texture_index int, offset []int32) {
	return
}
func (self implementation) FontGetTextureOffsets(font_rid RID.Any, size Vector2i.XY, texture_index int) (_ []int32) {
	return
}
func (self implementation) FontGetGlyphList(font_rid RID.Any, size Vector2i.XY) (_ []int32) { return }
func (self implementation) FontClearGlyphs(font_rid RID.Any, size Vector2i.XY)              { return }
func (self implementation) FontRemoveGlyph(font_rid RID.Any, size Vector2i.XY, glyph int)   { return }
func (self implementation) FontGetGlyphAdvance(font_rid RID.Any, size int, glyph int) (_ Vector2.XY) {
	return
}
func (self implementation) FontSetGlyphAdvance(font_rid RID.Any, size int, glyph int, advance Vector2.XY) {
	return
}
func (self implementation) FontGetGlyphOffset(font_rid RID.Any, size Vector2i.XY, glyph int) (_ Vector2.XY) {
	return
}
func (self implementation) FontSetGlyphOffset(font_rid RID.Any, size Vector2i.XY, glyph int, offset Vector2.XY) {
	return
}
func (self implementation) FontGetGlyphSize(font_rid RID.Any, size Vector2i.XY, glyph int) (_ Vector2.XY) {
	return
}
func (self implementation) FontSetGlyphSize(font_rid RID.Any, size Vector2i.XY, glyph int, gl_size Vector2.XY) {
	return
}
func (self implementation) FontGetGlyphUvRect(font_rid RID.Any, size Vector2i.XY, glyph int) (_ Rect2.PositionSize) {
	return
}
func (self implementation) FontSetGlyphUvRect(font_rid RID.Any, size Vector2i.XY, glyph int, uv_rect Rect2.PositionSize) {
	return
}
func (self implementation) FontGetGlyphTextureIdx(font_rid RID.Any, size Vector2i.XY, glyph int) (_ int) {
	return
}
func (self implementation) FontSetGlyphTextureIdx(font_rid RID.Any, size Vector2i.XY, glyph int, texture_idx int) {
	return
}
func (self implementation) FontGetGlyphTextureRid(font_rid RID.Any, size Vector2i.XY, glyph int) (_ RID.Any) {
	return
}
func (self implementation) FontGetGlyphTextureSize(font_rid RID.Any, size Vector2i.XY, glyph int) (_ Vector2.XY) {
	return
}
func (self implementation) FontGetGlyphContours(font_rid RID.Any, size int, index int) (_ map[any]any) {
	return
}
func (self implementation) FontGetKerningList(font_rid RID.Any, size int) (_ []Vector2i.XY) { return }
func (self implementation) FontClearKerningMap(font_rid RID.Any, size int)                  { return }
func (self implementation) FontRemoveKerning(font_rid RID.Any, size int, glyph_pair Vector2i.XY) {
	return
}
func (self implementation) FontSetKerning(font_rid RID.Any, size int, glyph_pair Vector2i.XY, kerning Vector2.XY) {
	return
}
func (self implementation) FontGetKerning(font_rid RID.Any, size int, glyph_pair Vector2i.XY) (_ Vector2.XY) {
	return
}
func (self implementation) FontGetGlyphIndex(font_rid RID.Any, size int, char int, variation_selector int) (_ int) {
	return
}
func (self implementation) FontGetCharFromGlyphIndex(font_rid RID.Any, size int, glyph_index int) (_ int) {
	return
}
func (self implementation) FontHasChar(font_rid RID.Any, char int) (_ bool)   { return }
func (self implementation) FontGetSupportedChars(font_rid RID.Any) (_ string) { return }
func (self implementation) FontRenderRange(font_rid RID.Any, size Vector2i.XY, start int, end int) {
	return
}
func (self implementation) FontRenderGlyph(font_rid RID.Any, size Vector2i.XY, index int) { return }
func (self implementation) FontDrawGlyph(font_rid RID.Any, canvas RID.Any, size int, pos Vector2.XY, index int, color Color.RGBA) {
	return
}
func (self implementation) FontDrawGlyphOutline(font_rid RID.Any, canvas RID.Any, size int, outline_size int, pos Vector2.XY, index int, color Color.RGBA) {
	return
}
func (self implementation) FontIsLanguageSupported(font_rid RID.Any, language string) (_ bool) {
	return
}
func (self implementation) FontSetLanguageSupportOverride(font_rid RID.Any, language string, supported bool) {
	return
}
func (self implementation) FontGetLanguageSupportOverride(font_rid RID.Any, language string) (_ bool) {
	return
}
func (self implementation) FontRemoveLanguageSupportOverride(font_rid RID.Any, language string) {
	return
}
func (self implementation) FontGetLanguageSupportOverrides(font_rid RID.Any) (_ []string)  { return }
func (self implementation) FontIsScriptSupported(font_rid RID.Any, script string) (_ bool) { return }
func (self implementation) FontSetScriptSupportOverride(font_rid RID.Any, script string, supported bool) {
	return
}
func (self implementation) FontGetScriptSupportOverride(font_rid RID.Any, script string) (_ bool) {
	return
}
func (self implementation) FontRemoveScriptSupportOverride(font_rid RID.Any, script string) { return }
func (self implementation) FontGetScriptSupportOverrides(font_rid RID.Any) (_ []string)     { return }
func (self implementation) FontSetOpentypeFeatureOverrides(font_rid RID.Any, overrides map[any]any) {
	return
}
func (self implementation) FontGetOpentypeFeatureOverrides(font_rid RID.Any) (_ map[any]any) { return }
func (self implementation) FontSupportedFeatureList(font_rid RID.Any) (_ map[any]any)        { return }
func (self implementation) FontSupportedVariationList(font_rid RID.Any) (_ map[any]any)      { return }
func (self implementation) FontGetGlobalOversampling() (_ Float.X)                           { return }
func (self implementation) FontSetGlobalOversampling(oversampling Float.X)                   { return }
func (self implementation) GetHexCodeBoxSize(size int, index int) (_ Vector2.XY)             { return }
func (self implementation) DrawHexCodeBox(canvas RID.Any, size int, pos Vector2.XY, index int, color Color.RGBA) {
	return
}
func (self implementation) CreateShapedText(direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) (_ RID.Any) {
	return
}
func (self implementation) ShapedTextClear(shaped RID.Any) { return }
func (self implementation) ShapedTextSetDirection(shaped RID.Any, direction gdclass.TextServerDirection) {
	return
}
func (self implementation) ShapedTextGetDirection(shaped RID.Any) (_ gdclass.TextServerDirection) {
	return
}
func (self implementation) ShapedTextGetInferredDirection(shaped RID.Any) (_ gdclass.TextServerDirection) {
	return
}
func (self implementation) ShapedTextSetBidiOverride(shaped RID.Any, override []any)    { return }
func (self implementation) ShapedTextSetCustomPunctuation(shaped RID.Any, punct string) { return }
func (self implementation) ShapedTextGetCustomPunctuation(shaped RID.Any) (_ string)    { return }
func (self implementation) ShapedTextSetCustomEllipsis(shaped RID.Any, char int)        { return }
func (self implementation) ShapedTextGetCustomEllipsis(shaped RID.Any) (_ int)          { return }
func (self implementation) ShapedTextSetOrientation(shaped RID.Any, orientation gdclass.TextServerOrientation) {
	return
}
func (self implementation) ShapedTextGetOrientation(shaped RID.Any) (_ gdclass.TextServerOrientation) {
	return
}
func (self implementation) ShapedTextSetPreserveInvalid(shaped RID.Any, enabled bool) { return }
func (self implementation) ShapedTextGetPreserveInvalid(shaped RID.Any) (_ bool)      { return }
func (self implementation) ShapedTextSetPreserveControl(shaped RID.Any, enabled bool) { return }
func (self implementation) ShapedTextGetPreserveControl(shaped RID.Any) (_ bool)      { return }
func (self implementation) ShapedTextSetSpacing(shaped RID.Any, spacing gdclass.TextServerSpacingType, value int) {
	return
}
func (self implementation) ShapedTextGetSpacing(shaped RID.Any, spacing gdclass.TextServerSpacingType) (_ int) {
	return
}
func (self implementation) ShapedTextAddString(shaped RID.Any, text string, fonts []RID.Any, size int, opentype_features map[any]any, language string, meta any) (_ bool) {
	return
}
func (self implementation) ShapedTextAddObject(shaped RID.Any, key any, size Vector2.XY, inline_align InlineAlignment, length int, baseline Float.X) (_ bool) {
	return
}
func (self implementation) ShapedTextResizeObject(shaped RID.Any, key any, size Vector2.XY, inline_align InlineAlignment, baseline Float.X) (_ bool) {
	return
}
func (self implementation) ShapedGetSpanCount(shaped RID.Any) (_ int)           { return }
func (self implementation) ShapedGetSpanMeta(shaped RID.Any, index int) (_ any) { return }
func (self implementation) ShapedSetSpanUpdateFont(shaped RID.Any, index int, fonts []RID.Any, size int, opentype_features map[any]any) {
	return
}
func (self implementation) ShapedTextSubstr(shaped RID.Any, start int, length int) (_ RID.Any) {
	return
}
func (self implementation) ShapedTextGetParent(shaped RID.Any) (_ RID.Any) { return }
func (self implementation) ShapedTextFitToWidth(shaped RID.Any, width Float.X, justification_flags gdclass.TextServerJustificationFlag) (_ Float.X) {
	return
}
func (self implementation) ShapedTextTabAlign(shaped RID.Any, tab_stops []float32) (_ Float.X) {
	return
}
func (self implementation) ShapedTextShape(shaped RID.Any) (_ bool)                  { return }
func (self implementation) ShapedTextUpdateBreaks(shaped RID.Any) (_ bool)           { return }
func (self implementation) ShapedTextUpdateJustificationOps(shaped RID.Any) (_ bool) { return }
func (self implementation) ShapedTextIsReady(shaped RID.Any) (_ bool)                { return }
func (self implementation) ShapedTextGetGlyphs(shaped RID.Any) (_ *Glyph)            { return }
func (self implementation) ShapedTextSortLogical(shaped RID.Any) (_ *Glyph)          { return }
func (self implementation) ShapedTextGetGlyphCount(shaped RID.Any) (_ int)           { return }
func (self implementation) ShapedTextGetRange(shaped RID.Any) (_ Vector2i.XY)        { return }
func (self implementation) ShapedTextGetLineBreaksAdv(shaped RID.Any, width []float32, start int, once bool, break_flags gdclass.TextServerLineBreakFlag) (_ []int32) {
	return
}
func (self implementation) ShapedTextGetLineBreaks(shaped RID.Any, width Float.X, start int, break_flags gdclass.TextServerLineBreakFlag) (_ []int32) {
	return
}
func (self implementation) ShapedTextGetWordBreaks(shaped RID.Any, grapheme_flags gdclass.TextServerGraphemeFlag, skip_grapheme_flags gdclass.TextServerGraphemeFlag) (_ []int32) {
	return
}
func (self implementation) ShapedTextGetTrimPos(shaped RID.Any) (_ int)            { return }
func (self implementation) ShapedTextGetEllipsisPos(shaped RID.Any) (_ int)        { return }
func (self implementation) ShapedTextGetEllipsisGlyphCount(shaped RID.Any) (_ int) { return }
func (self implementation) ShapedTextGetEllipsisGlyphs(shaped RID.Any) (_ *Glyph)  { return }
func (self implementation) ShapedTextOverrunTrimToWidth(shaped RID.Any, width Float.X, trim_flags gdclass.TextServerTextOverrunFlag) {
	return
}
func (self implementation) ShapedTextGetObjects(shaped RID.Any) (_ []any) { return }
func (self implementation) ShapedTextGetObjectRect(shaped RID.Any, key any) (_ Rect2.PositionSize) {
	return
}
func (self implementation) ShapedTextGetObjectRange(shaped RID.Any, key any) (_ Vector2i.XY) { return }
func (self implementation) ShapedTextGetObjectGlyph(shaped RID.Any, key any) (_ int)         { return }
func (self implementation) ShapedTextGetSize(shaped RID.Any) (_ Vector2.XY)                  { return }
func (self implementation) ShapedTextGetAscent(shaped RID.Any) (_ Float.X)                   { return }
func (self implementation) ShapedTextGetDescent(shaped RID.Any) (_ Float.X)                  { return }
func (self implementation) ShapedTextGetWidth(shaped RID.Any) (_ Float.X)                    { return }
func (self implementation) ShapedTextGetUnderlinePosition(shaped RID.Any) (_ Float.X)        { return }
func (self implementation) ShapedTextGetUnderlineThickness(shaped RID.Any) (_ Float.X)       { return }
func (self implementation) ShapedTextGetDominantDirectionInRange(shaped RID.Any, start int, end int) (_ int) {
	return
}
func (self implementation) ShapedTextGetCarets(shaped RID.Any, position int, caret *CaretInfo) {
	return
}
func (self implementation) ShapedTextGetSelection(shaped RID.Any, start int, end int) (_ []Vector2.XY) {
	return
}
func (self implementation) ShapedTextHitTestGrapheme(shaped RID.Any, coord Float.X) (_ int) { return }
func (self implementation) ShapedTextHitTestPosition(shaped RID.Any, coord Float.X) (_ int) { return }
func (self implementation) ShapedTextDraw(shaped RID.Any, canvas RID.Any, pos Vector2.XY, clip_l Float.X, clip_r Float.X, color Color.RGBA) {
	return
}
func (self implementation) ShapedTextDrawOutline(shaped RID.Any, canvas RID.Any, pos Vector2.XY, clip_l Float.X, clip_r Float.X, outline_size int, color Color.RGBA) {
	return
}
func (self implementation) ShapedTextGetGraphemeBounds(shaped RID.Any, pos int) (_ Vector2.XY) {
	return
}
func (self implementation) ShapedTextNextGraphemePos(shaped RID.Any, pos int) (_ int)     { return }
func (self implementation) ShapedTextPrevGraphemePos(shaped RID.Any, pos int) (_ int)     { return }
func (self implementation) ShapedTextGetCharacterBreaks(shaped RID.Any) (_ []int32)       { return }
func (self implementation) ShapedTextNextCharacterPos(shaped RID.Any, pos int) (_ int)    { return }
func (self implementation) ShapedTextPrevCharacterPos(shaped RID.Any, pos int) (_ int)    { return }
func (self implementation) ShapedTextClosestCharacterPos(shaped RID.Any, pos int) (_ int) { return }
func (self implementation) FormatNumber(number string, language string) (_ string)        { return }
func (self implementation) ParseNumber(number string, language string) (_ string)         { return }
func (self implementation) PercentSign(language string) (_ string)                        { return }
func (self implementation) StripDiacritics(s string) (_ string)                           { return }
func (self implementation) IsValidIdentifier(s string) (_ bool)                           { return }
func (self implementation) IsValidLetter(unicode int) (_ bool)                            { return }
func (self implementation) StringGetWordBreaks(s string, language string, chars_per_line int) (_ []int32) {
	return
}
func (self implementation) StringGetCharacterBreaks(s string, language string) (_ []int32) { return }
func (self implementation) IsConfusable(s string, dict []string) (_ int)                   { return }
func (self implementation) SpoofCheck(s string) (_ bool)                                   { return }
func (self implementation) StringToUpper(s string, language string) (_ string)             { return }
func (self implementation) StringToLower(s string, language string) (_ string)             { return }
func (self implementation) StringToTitle(s string, language string) (_ string)             { return }
func (self implementation) ParseStructuredText(parser_type gdclass.TextServerStructuredTextParser, args []any, text string) (_ []Vector3i.XYZ) {
	return
}
func (self implementation) Cleanup() { return }

/*
[b]Required.[/b]
Returns [code]true[/code] if the server supports a feature.
*/
func (Instance) _has_feature(impl func(ptr unsafe.Pointer, feature gdclass.TextServerFeature) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var feature = gd.UnsafeGet[gdclass.TextServerFeature](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, feature)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns the name of the server interface.
*/
func (Instance) _get_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns text server features, see [enum TextServer.Feature].
*/
func (Instance) _get_features(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Frees an object created by this [TextServer].
*/
func (Instance) _free_rid(impl func(ptr unsafe.Pointer, rid RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
*/
func (Instance) _has(impl func(ptr unsafe.Pointer, rid RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
*/
func (Instance) _load_support_data(impl func(ptr unsafe.Pointer, filename string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var filename = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(filename))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, filename.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
*/
func (Instance) _get_support_data_filename(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
*/
func (Instance) _get_support_data_info(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
*/
func (Instance) _save_support_data(impl func(ptr unsafe.Pointer, filename string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var filename = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(filename))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, filename.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if locale is right-to-left.
*/
func (Instance) _is_locale_right_to_left(impl func(ptr unsafe.Pointer, locale string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var locale = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(locale))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, locale.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Converts readable feature, variation, script, or language name to OpenType tag.
*/
func (Instance) _name_to_tag(impl func(ptr unsafe.Pointer, name string) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(name))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name.String())
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Converts OpenType tag to readable feature, variation, script, or language name.
*/
func (Instance) _tag_to_name(impl func(ptr unsafe.Pointer, tag int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var tag = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(tag))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Creates a new, empty font cache entry resource.
*/
func (Instance) _create_font(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Optional, implement if font supports extra spacing or baseline offset.
Creates a new variation existing font which is reusing the same glyph cache and font data.
*/
func (Instance) _create_font_linked_variation(impl func(ptr unsafe.Pointer, font_rid RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
[b]Optional.[/b]
Sets font source data, e.g contents of the dynamic font source file.
*/
func (Instance) _font_set_data(impl func(ptr unsafe.Pointer, font_rid RID.Any, data []byte)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var data = pointers.New[gd.PackedByteArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(data)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, data.Bytes())
	}
}

/*
[b]Optional.[/b]
Sets pointer to the font source data, e.g contents of the dynamic font source file.
*/
func (Instance) _font_set_data_ptr(impl func(ptr unsafe.Pointer, font_rid RID.Any, data_ptr unsafe.Pointer, data_size int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var data_ptr = gd.UnsafeGet[unsafe.Pointer](p_args, 1)

		var data_size = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, data_ptr, int(data_size))
	}
}

/*
[b]Optional.[/b]
Sets an active face index in the TrueType / OpenType collection.
*/
func (Instance) _font_set_face_index(impl func(ptr unsafe.Pointer, font_rid RID.Any, face_index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var face_index = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(face_index))
	}
}

/*
[b]Optional.[/b]
Returns an active face index in the TrueType / OpenType collection.
*/
func (Instance) _font_get_face_index(impl func(ptr unsafe.Pointer, font_rid RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Returns number of faces in the TrueType / OpenType collection.
*/
func (Instance) _font_get_face_count(impl func(ptr unsafe.Pointer, font_rid RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Sets the font style flags, see [enum TextServer.FontStyle].
*/
func (Instance) _font_set_style(impl func(ptr unsafe.Pointer, font_rid RID.Any, style gdclass.TextServerFontStyle)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var style = gd.UnsafeGet[gdclass.TextServerFontStyle](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, style)
	}
}

/*
[b]Optional.[/b]
Returns font style flags, see [enum TextServer.FontStyle].
*/
func (Instance) _font_get_style(impl func(ptr unsafe.Pointer, font_rid RID.Any) gdclass.TextServerFontStyle) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets the font family name.
*/
func (Instance) _font_set_name(impl func(ptr unsafe.Pointer, font_rid RID.Any, name string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(name))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, name.String())
	}
}

/*
[b]Optional.[/b]
Returns font family name.
*/
func (Instance) _font_get_name(impl func(ptr unsafe.Pointer, font_rid RID.Any) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
func (Instance) _font_get_ot_name_strings(impl func(ptr unsafe.Pointer, font_rid RID.Any) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets the font style name.
*/
func (Instance) _font_set_style_name(impl func(ptr unsafe.Pointer, font_rid RID.Any, name_style string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var name_style = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(name_style))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, name_style.String())
	}
}

/*
[b]Optional.[/b]
Returns font style name.
*/
func (Instance) _font_get_style_name(impl func(ptr unsafe.Pointer, font_rid RID.Any) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (Instance) _font_set_weight(impl func(ptr unsafe.Pointer, font_rid RID.Any, weight int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var weight = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(weight))
	}
}

/*
[b]Optional.[/b]
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (Instance) _font_get_weight(impl func(ptr unsafe.Pointer, font_rid RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (Instance) _font_set_stretch(impl func(ptr unsafe.Pointer, font_rid RID.Any, stretch int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var stretch = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(stretch))
	}
}

/*
[b]Optional.[/b]
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (Instance) _font_get_stretch(impl func(ptr unsafe.Pointer, font_rid RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Sets font anti-aliasing mode.
*/
func (Instance) _font_set_antialiasing(impl func(ptr unsafe.Pointer, font_rid RID.Any, antialiasing gdclass.TextServerFontAntialiasing)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var antialiasing = gd.UnsafeGet[gdclass.TextServerFontAntialiasing](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, antialiasing)
	}
}

/*
[b]Optional.[/b]
Returns font anti-aliasing mode.
*/
func (Instance) _font_get_antialiasing(impl func(ptr unsafe.Pointer, font_rid RID.Any) gdclass.TextServerFontAntialiasing) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], embedded font bitmap loading is disabled.
*/
func (Instance) _font_set_disable_embedded_bitmaps(impl func(ptr unsafe.Pointer, font_rid RID.Any, disable_embedded_bitmaps bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var disable_embedded_bitmaps = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, disable_embedded_bitmaps)
	}
}

/*
[b]Optional.[/b]
Returns whether the font's embedded bitmap loading is disabled.
*/
func (Instance) _font_get_disable_embedded_bitmaps(impl func(ptr unsafe.Pointer, font_rid RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] font texture mipmap generation is enabled.
*/
func (Instance) _font_set_generate_mipmaps(impl func(ptr unsafe.Pointer, font_rid RID.Any, generate_mipmaps bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var generate_mipmaps = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, generate_mipmaps)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if font texture mipmap generation is enabled.
*/
func (Instance) _font_get_generate_mipmaps(impl func(ptr unsafe.Pointer, font_rid RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
*/
func (Instance) _font_set_multichannel_signed_distance_field(impl func(ptr unsafe.Pointer, font_rid RID.Any, msdf bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var msdf = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, msdf)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
*/
func (Instance) _font_is_multichannel_signed_distance_field(impl func(ptr unsafe.Pointer, font_rid RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (Instance) _font_set_msdf_pixel_range(impl func(ptr unsafe.Pointer, font_rid RID.Any, msdf_pixel_range int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var msdf_pixel_range = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(msdf_pixel_range))
	}
}

/*
[b]Optional.[/b]
Returns the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (Instance) _font_get_msdf_pixel_range(impl func(ptr unsafe.Pointer, font_rid RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Sets source font size used to generate MSDF textures.
*/
func (Instance) _font_set_msdf_size(impl func(ptr unsafe.Pointer, font_rid RID.Any, msdf_size int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var msdf_size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(msdf_size))
	}
}

/*
[b]Optional.[/b]
Returns source font size used to generate MSDF textures.
*/
func (Instance) _font_get_msdf_size(impl func(ptr unsafe.Pointer, font_rid RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
*/
func (Instance) _font_set_fixed_size(impl func(ptr unsafe.Pointer, font_rid RID.Any, fixed_size int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var fixed_size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(fixed_size))
	}
}

/*
[b]Required.[/b]
Returns bitmap font fixed size.
*/
func (Instance) _font_get_fixed_size(impl func(ptr unsafe.Pointer, font_rid RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
*/
func (Instance) _font_set_fixed_size_scale_mode(impl func(ptr unsafe.Pointer, font_rid RID.Any, fixed_size_scale_mode gdclass.TextServerFixedSizeScaleMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var fixed_size_scale_mode = gd.UnsafeGet[gdclass.TextServerFixedSizeScaleMode](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, fixed_size_scale_mode)
	}
}

/*
[b]Required.[/b]
Returns bitmap font scaling mode.
*/
func (Instance) _font_get_fixed_size_scale_mode(impl func(ptr unsafe.Pointer, font_rid RID.Any) gdclass.TextServerFixedSizeScaleMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], system fonts can be automatically used as fallbacks.
*/
func (Instance) _font_set_allow_system_fallback(impl func(ptr unsafe.Pointer, font_rid RID.Any, allow_system_fallback bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var allow_system_fallback = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, allow_system_fallback)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
*/
func (Instance) _font_is_allow_system_fallback(impl func(ptr unsafe.Pointer, font_rid RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
*/
func (Instance) _font_set_force_autohinter(impl func(ptr unsafe.Pointer, font_rid RID.Any, force_autohinter bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var force_autohinter = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, force_autohinter)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting.
*/
func (Instance) _font_is_force_autohinter(impl func(ptr unsafe.Pointer, font_rid RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets font hinting mode. Used by dynamic fonts only.
*/
func (Instance) _font_set_hinting(impl func(ptr unsafe.Pointer, font_rid RID.Any, hinting gdclass.TextServerHinting)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var hinting = gd.UnsafeGet[gdclass.TextServerHinting](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, hinting)
	}
}

/*
[b]Optional.[/b]
Returns the font hinting mode. Used by dynamic fonts only.
*/
func (Instance) _font_get_hinting(impl func(ptr unsafe.Pointer, font_rid RID.Any) gdclass.TextServerHinting) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets font subpixel glyph positioning mode.
*/
func (Instance) _font_set_subpixel_positioning(impl func(ptr unsafe.Pointer, font_rid RID.Any, subpixel_positioning gdclass.TextServerSubpixelPositioning)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var subpixel_positioning = gd.UnsafeGet[gdclass.TextServerSubpixelPositioning](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, subpixel_positioning)
	}
}

/*
[b]Optional.[/b]
Returns font subpixel glyph positioning mode.
*/
func (Instance) _font_get_subpixel_positioning(impl func(ptr unsafe.Pointer, font_rid RID.Any) gdclass.TextServerSubpixelPositioning) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (Instance) _font_set_embolden(impl func(ptr unsafe.Pointer, font_rid RID.Any, strength Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var strength = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, Float.X(strength))
	}
}

/*
[b]Optional.[/b]
Returns font embolden strength.
*/
func (Instance) _font_get_embolden(impl func(ptr unsafe.Pointer, font_rid RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Optional.[/b]
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
func (Instance) _font_set_spacing(impl func(ptr unsafe.Pointer, font_rid RID.Any, spacing gdclass.TextServerSpacingType, value int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var spacing = gd.UnsafeGet[gdclass.TextServerSpacingType](p_args, 1)

		var value = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, spacing, int(value))
	}
}

/*
[b]Optional.[/b]
Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
func (Instance) _font_get_spacing(impl func(ptr unsafe.Pointer, font_rid RID.Any, spacing gdclass.TextServerSpacingType) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var spacing = gd.UnsafeGet[gdclass.TextServerSpacingType](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, spacing)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Sets extra baseline offset (as a fraction of font height).
*/
func (Instance) _font_set_baseline_offset(impl func(ptr unsafe.Pointer, font_rid RID.Any, baseline_offset Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var baseline_offset = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, Float.X(baseline_offset))
	}
}

/*
[b]Optional.[/b]
Returns extra baseline offset (as a fraction of font height).
*/
func (Instance) _font_get_baseline_offset(impl func(ptr unsafe.Pointer, font_rid RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Optional.[/b]
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
func (Instance) _font_set_transform(impl func(ptr unsafe.Pointer, font_rid RID.Any, transform Transform2D.OriginXY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, transform)
	}
}

/*
[b]Optional.[/b]
Returns 2D transform applied to the font outlines.
*/
func (Instance) _font_get_transform(impl func(ptr unsafe.Pointer, font_rid RID.Any) Transform2D.OriginXY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Transform2D(ret))
	}
}

/*
[b]Optional.[/b]
Sets variation coordinates for the specified font cache entry.
*/
func (Instance) _font_set_variation_coordinates(impl func(ptr unsafe.Pointer, font_rid RID.Any, variation_coordinates map[any]any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var variation_coordinates = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalDictionary(variation_coordinates))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, gd.DictionaryAs[map[any]any](variation_coordinates))
	}
}

/*
[b]Optional.[/b]
Returns variation coordinates for the specified font cache entry.
*/
func (Instance) _font_get_variation_coordinates(impl func(ptr unsafe.Pointer, font_rid RID.Any) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (Instance) _font_set_oversampling(impl func(ptr unsafe.Pointer, font_rid RID.Any, oversampling Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var oversampling = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, Float.X(oversampling))
	}
}

/*
[b]Optional.[/b]
Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (Instance) _font_get_oversampling(impl func(ptr unsafe.Pointer, font_rid RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
func (Instance) _font_get_size_cache_list(impl func(ptr unsafe.Pointer, font_rid RID.Any) []Vector2i.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[gd.Vector2i]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Removes all font sizes from the cache entry.
*/
func (Instance) _font_clear_size_cache(impl func(ptr unsafe.Pointer, font_rid RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid)
	}
}

/*
[b]Required.[/b]
Removes specified font size from the cache entry.
*/
func (Instance) _font_remove_size_cache(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size)
	}
}

/*
[b]Required.[/b]
Sets the font ascent (number of pixels above the baseline).
*/
func (Instance) _font_set_ascent(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, ascent Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var ascent = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size), Float.X(ascent))
	}
}

/*
[b]Required.[/b]
Returns the font ascent (number of pixels above the baseline).
*/
func (Instance) _font_get_ascent(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Sets the font descent (number of pixels below the baseline).
*/
func (Instance) _font_set_descent(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, descent Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var descent = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size), Float.X(descent))
	}
}

/*
[b]Required.[/b]
Returns the font descent (number of pixels below the baseline).
*/
func (Instance) _font_get_descent(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Sets pixel offset of the underline below the baseline.
*/
func (Instance) _font_set_underline_position(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, underline_position Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var underline_position = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size), Float.X(underline_position))
	}
}

/*
[b]Required.[/b]
Returns pixel offset of the underline below the baseline.
*/
func (Instance) _font_get_underline_position(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Sets thickness of the underline in pixels.
*/
func (Instance) _font_set_underline_thickness(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, underline_thickness Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var underline_thickness = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size), Float.X(underline_thickness))
	}
}

/*
[b]Required.[/b]
Returns thickness of the underline in pixels.
*/
func (Instance) _font_get_underline_thickness(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Sets scaling factor of the color bitmap font.
*/
func (Instance) _font_set_scale(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, scale Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var scale = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size), Float.X(scale))
	}
}

/*
[b]Required.[/b]
Returns scaling factor of the color bitmap font.
*/
func (Instance) _font_get_scale(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Returns number of textures used by font cache entry.
*/
func (Instance) _font_get_texture_count(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Removes all textures from font cache entry.
*/
func (Instance) _font_clear_textures(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size)
	}
}

/*
[b]Required.[/b]
Removes specified texture from the cache entry.
*/
func (Instance) _font_remove_texture(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, texture_index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(texture_index))
	}
}

/*
[b]Required.[/b]
Sets font cache texture image data.
*/
func (Instance) _font_set_texture_image(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, texture_index int, image [1]gdclass.Image)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}

		defer pointers.End(image[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(texture_index), image)
	}
}

/*
[b]Required.[/b]
Returns font cache texture image data.
*/
func (Instance) _font_get_texture_image(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, texture_index int) [1]gdclass.Image) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(texture_index))
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets array containing glyph packing data.
*/
func (Instance) _font_set_texture_offsets(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, texture_index int, offset []int32)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		var offset = pointers.New[gd.PackedInt32Array](gd.UnsafeGet[gd.PackedPointers](p_args, 3))
		defer pointers.End(offset)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(texture_index), offset.AsSlice())
	}
}

/*
[b]Optional.[/b]
Returns array containing glyph packing data.
*/
func (Instance) _font_get_texture_offsets(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, texture_index int) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(texture_index))
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns list of rendered glyphs in the cache entry.
*/
func (Instance) _font_get_glyph_list(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Removes all rendered glyph information from the cache entry.
*/
func (Instance) _font_clear_glyphs(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size)
	}
}

/*
[b]Required.[/b]
Removes specified rendered glyph information from the cache entry.
*/
func (Instance) _font_remove_glyph(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(glyph))
	}
}

/*
[b]Required.[/b]
Returns glyph advance (offset of the next glyph).
*/
func (Instance) _font_get_glyph_advance(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, glyph int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), int(glyph))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
[b]Required.[/b]
Sets glyph advance (offset of the next glyph).
*/
func (Instance) _font_set_glyph_advance(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, glyph int, advance Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var advance = gd.UnsafeGet[gd.Vector2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size), int(glyph), advance)
	}
}

/*
[b]Required.[/b]
Returns glyph offset from the baseline.
*/
func (Instance) _font_get_glyph_offset(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
[b]Required.[/b]
Sets glyph offset from the baseline.
*/
func (Instance) _font_set_glyph_offset(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int, offset Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var offset = gd.UnsafeGet[gd.Vector2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(glyph), offset)
	}
}

/*
[b]Required.[/b]
Returns size of the glyph.
*/
func (Instance) _font_get_glyph_size(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
[b]Required.[/b]
Sets size of the glyph.
*/
func (Instance) _font_set_glyph_size(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int, gl_size Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var gl_size = gd.UnsafeGet[gd.Vector2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(glyph), gl_size)
	}
}

/*
[b]Required.[/b]
Returns rectangle in the cache texture containing the glyph.
*/
func (Instance) _font_get_glyph_uv_rect(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int) Rect2.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, gd.Rect2(ret))
	}
}

/*
[b]Required.[/b]
Sets rectangle in the cache texture containing the glyph.
*/
func (Instance) _font_set_glyph_uv_rect(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int, uv_rect Rect2.PositionSize)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var uv_rect = gd.UnsafeGet[gd.Rect2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(glyph), uv_rect)
	}
}

/*
[b]Required.[/b]
Returns index of the cache texture containing the glyph.
*/
func (Instance) _font_get_glyph_texture_idx(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Sets index of the cache texture containing the glyph.
*/
func (Instance) _font_set_glyph_texture_idx(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int, texture_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var texture_idx = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(glyph), int(texture_idx))
	}
}

/*
[b]Required.[/b]
Returns resource ID of the cache texture containing the glyph.
*/
func (Instance) _font_get_glyph_texture_rid(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
[b]Required.[/b]
Returns size of the cache texture containing the glyph.
*/
func (Instance) _font_get_glyph_texture_size(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, glyph int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
[b]Optional.[/b]
Returns outline contours of the glyph.
*/
func (Instance) _font_get_glyph_contours(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, index int) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), int(index))
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns list of the kerning overrides.
*/
func (Instance) _font_get_kerning_list(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int) []Vector2i.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[gd.Vector2i]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Removes all kerning overrides.
*/
func (Instance) _font_clear_kerning_map(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size))
	}
}

/*
[b]Optional.[/b]
Removes kerning override for the pair of glyphs.
*/
func (Instance) _font_remove_kerning(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, glyph_pair Vector2i.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size), glyph_pair)
	}
}

/*
[b]Optional.[/b]
Sets kerning for the pair of glyphs.
*/
func (Instance) _font_set_kerning(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, glyph_pair Vector2i.XY, kerning Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args, 2)

		var kerning = gd.UnsafeGet[gd.Vector2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, int(size), glyph_pair, kerning)
	}
}

/*
[b]Optional.[/b]
Returns kerning for the pair of glyphs.
*/
func (Instance) _font_get_kerning(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, glyph_pair Vector2i.XY) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), glyph_pair)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
[b]Required.[/b]
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
func (Instance) _font_get_glyph_index(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, char int, variation_selector int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var char = gd.UnsafeGet[gd.Int](p_args, 2)

		var variation_selector = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), int(char), int(variation_selector))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid.
*/
func (Instance) _font_get_char_from_glyph_index(impl func(ptr unsafe.Pointer, font_rid RID.Any, size int, glyph_index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph_index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), int(glyph_index))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
func (Instance) _font_has_char(impl func(ptr unsafe.Pointer, font_rid RID.Any, char int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var char = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(char))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns a string containing all the characters available in the font.
*/
func (Instance) _font_get_supported_chars(impl func(ptr unsafe.Pointer, font_rid RID.Any) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Renders the range of characters to the font cache texture.
*/
func (Instance) _font_render_range(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, start int, end int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var start = gd.UnsafeGet[gd.Int](p_args, 2)

		var end = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(start), int(end))
	}
}

/*
[b]Optional.[/b]
Renders specified glyph to the font cache texture.
*/
func (Instance) _font_render_glyph(impl func(ptr unsafe.Pointer, font_rid RID.Any, size Vector2i.XY, index int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, int(index))
	}
}

/*
[b]Required.[/b]
Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
*/
func (Instance) _font_draw_glyph(impl func(ptr unsafe.Pointer, font_rid RID.Any, canvas RID.Any, size int, pos Vector2.XY, index int, color Color.RGBA)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var canvas = gd.UnsafeGet[gd.RID](p_args, 1)

		var size = gd.UnsafeGet[gd.Int](p_args, 2)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 3)

		var index = gd.UnsafeGet[gd.Int](p_args, 4)

		var color = gd.UnsafeGet[gd.Color](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, canvas, int(size), pos, int(index), color)
	}
}

/*
[b]Required.[/b]
Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
*/
func (Instance) _font_draw_glyph_outline(impl func(ptr unsafe.Pointer, font_rid RID.Any, canvas RID.Any, size int, outline_size int, pos Vector2.XY, index int, color Color.RGBA)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var canvas = gd.UnsafeGet[gd.RID](p_args, 1)

		var size = gd.UnsafeGet[gd.Int](p_args, 2)

		var outline_size = gd.UnsafeGet[gd.Int](p_args, 3)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 4)

		var index = gd.UnsafeGet[gd.Int](p_args, 5)

		var color = gd.UnsafeGet[gd.Color](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, canvas, int(size), int(outline_size), pos, int(index), color)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
func (Instance) _font_is_language_supported(impl func(ptr unsafe.Pointer, font_rid RID.Any, language string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, language.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Adds override for [method _font_is_language_supported].
*/
func (Instance) _font_set_language_support_override(impl func(ptr unsafe.Pointer, font_rid RID.Any, language string, supported bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		var supported = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, language.String(), supported)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if support override is enabled for the [param language].
*/
func (Instance) _font_get_language_support_override(impl func(ptr unsafe.Pointer, font_rid RID.Any, language string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, language.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Remove language support override.
*/
func (Instance) _font_remove_language_support_override(impl func(ptr unsafe.Pointer, font_rid RID.Any, language string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, language.String())
	}
}

/*
[b]Optional.[/b]
Returns list of language support overrides.
*/
func (Instance) _font_get_language_support_overrides(impl func(ptr unsafe.Pointer, font_rid RID.Any) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code], if font supports given script (ISO 15924 code).
*/
func (Instance) _font_is_script_supported(impl func(ptr unsafe.Pointer, font_rid RID.Any, script string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(script))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, script.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Adds override for [method _font_is_script_supported].
*/
func (Instance) _font_set_script_support_override(impl func(ptr unsafe.Pointer, font_rid RID.Any, script string, supported bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(script))
		var supported = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, script.String(), supported)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if support override is enabled for the [param script].
*/
func (Instance) _font_get_script_support_override(impl func(ptr unsafe.Pointer, font_rid RID.Any, script string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(script))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, script.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Removes script support override.
*/
func (Instance) _font_remove_script_support_override(impl func(ptr unsafe.Pointer, font_rid RID.Any, script string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(script))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, script.String())
	}
}

/*
[b]Optional.[/b]
Returns list of script support overrides.
*/
func (Instance) _font_get_script_support_overrides(impl func(ptr unsafe.Pointer, font_rid RID.Any) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets font OpenType feature set override.
*/
func (Instance) _font_set_opentype_feature_overrides(impl func(ptr unsafe.Pointer, font_rid RID.Any, overrides map[any]any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var overrides = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalDictionary(overrides))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, gd.DictionaryAs[map[any]any](overrides))
	}
}

/*
[b]Optional.[/b]
Returns font OpenType feature set override.
*/
func (Instance) _font_get_opentype_feature_overrides(impl func(ptr unsafe.Pointer, font_rid RID.Any) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the dictionary of the supported OpenType features.
*/
func (Instance) _font_supported_feature_list(impl func(ptr unsafe.Pointer, font_rid RID.Any) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the dictionary of the supported OpenType variation coordinates.
*/
func (Instance) _font_supported_variation_list(impl func(ptr unsafe.Pointer, font_rid RID.Any) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the font oversampling factor, shared by all fonts in the TextServer.
*/
func (Instance) _font_get_global_oversampling(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Optional.[/b]
Sets oversampling factor, shared by all font in the TextServer.
*/
func (Instance) _font_set_global_oversampling(impl func(ptr unsafe.Pointer, oversampling Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var oversampling = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(oversampling))
	}
}

/*
[b]Optional.[/b]
Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
*/
func (Instance) _get_hex_code_box_size(impl func(ptr unsafe.Pointer, size int, index int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var size = gd.UnsafeGet[gd.Int](p_args, 0)

		var index = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(size), int(index))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
[b]Optional.[/b]
Draws box displaying character hexadecimal code.
*/
func (Instance) _draw_hex_code_box(impl func(ptr unsafe.Pointer, canvas RID.Any, size int, pos Vector2.XY, index int, color Color.RGBA)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var canvas = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var index = gd.UnsafeGet[gd.Int](p_args, 3)

		var color = gd.UnsafeGet[gd.Color](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, canvas, int(size), pos, int(index), color)
	}
}

/*
[b]Required.[/b]
Creates a new buffer for complex text layout, with the given [param direction] and [param orientation].
*/
func (Instance) _create_shaped_text(impl func(ptr unsafe.Pointer, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var direction = gd.UnsafeGet[gdclass.TextServerDirection](p_args, 0)

		var orientation = gd.UnsafeGet[gdclass.TextServerOrientation](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, direction, orientation)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
[b]Required.[/b]
Clears text buffer (removes text and inline objects).
*/
func (Instance) _shaped_text_clear(impl func(ptr unsafe.Pointer, shaped RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped)
	}
}

/*
[b]Optional.[/b]
Sets desired text direction. If set to [constant TextServer.DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
*/
func (Instance) _shaped_text_set_direction(impl func(ptr unsafe.Pointer, shaped RID.Any, direction gdclass.TextServerDirection)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var direction = gd.UnsafeGet[gdclass.TextServerDirection](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, direction)
	}
}

/*
[b]Optional.[/b]
Returns direction of the text.
*/
func (Instance) _shaped_text_get_direction(impl func(ptr unsafe.Pointer, shaped RID.Any) gdclass.TextServerDirection) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns direction of the text, inferred by the BiDi algorithm.
*/
func (Instance) _shaped_text_get_inferred_direction(impl func(ptr unsafe.Pointer, shaped RID.Any) gdclass.TextServerDirection) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Overrides BiDi for the structured text.
*/
func (Instance) _shaped_text_set_bidi_override(impl func(ptr unsafe.Pointer, shaped RID.Any, override []any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var override = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalArray(override))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, gd.ArrayAs[[]any](gd.InternalArray(override)))
	}
}

/*
[b]Optional.[/b]
Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (Instance) _shaped_text_set_custom_punctuation(impl func(ptr unsafe.Pointer, shaped RID.Any, punct string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var punct = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(punct))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, punct.String())
	}
}

/*
[b]Optional.[/b]
Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (Instance) _shaped_text_get_custom_punctuation(impl func(ptr unsafe.Pointer, shaped RID.Any) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets ellipsis character used for text clipping.
*/
func (Instance) _shaped_text_set_custom_ellipsis(impl func(ptr unsafe.Pointer, shaped RID.Any, char int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var char = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, int(char))
	}
}

/*
[b]Optional.[/b]
Returns ellipsis character used for text clipping.
*/
func (Instance) _shaped_text_get_custom_ellipsis(impl func(ptr unsafe.Pointer, shaped RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Sets desired text orientation.
*/
func (Instance) _shaped_text_set_orientation(impl func(ptr unsafe.Pointer, shaped RID.Any, orientation gdclass.TextServerOrientation)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var orientation = gd.UnsafeGet[gdclass.TextServerOrientation](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, orientation)
	}
}

/*
[b]Optional.[/b]
Returns text orientation.
*/
func (Instance) _shaped_text_get_orientation(impl func(ptr unsafe.Pointer, shaped RID.Any) gdclass.TextServerOrientation) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
*/
func (Instance) _shaped_text_set_preserve_invalid(impl func(ptr unsafe.Pointer, shaped RID.Any, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var enabled = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, enabled)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
*/
func (Instance) _shaped_text_get_preserve_invalid(impl func(ptr unsafe.Pointer, shaped RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] text buffer will display control characters.
*/
func (Instance) _shaped_text_set_preserve_control(impl func(ptr unsafe.Pointer, shaped RID.Any, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var enabled = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, enabled)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if text buffer is configured to display control characters.
*/
func (Instance) _shaped_text_get_preserve_control(impl func(ptr unsafe.Pointer, shaped RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets extra spacing added between glyphs or lines in pixels.
*/
func (Instance) _shaped_text_set_spacing(impl func(ptr unsafe.Pointer, shaped RID.Any, spacing gdclass.TextServerSpacingType, value int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var spacing = gd.UnsafeGet[gdclass.TextServerSpacingType](p_args, 1)

		var value = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, spacing, int(value))
	}
}

/*
[b]Optional.[/b]
Returns extra spacing added between glyphs or lines in pixels.
*/
func (Instance) _shaped_text_get_spacing(impl func(ptr unsafe.Pointer, shaped RID.Any, spacing gdclass.TextServerSpacingType) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var spacing = gd.UnsafeGet[gdclass.TextServerSpacingType](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, spacing)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Adds text span and font to draw it to the text buffer.
*/
func (Instance) _shaped_text_add_string(impl func(ptr unsafe.Pointer, shaped RID.Any, text string, fonts []RID.Any, size int, opentype_features map[any]any, language string, meta any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var text = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(text))
		var fonts = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalArray(fonts))
		var size = gd.UnsafeGet[gd.Int](p_args, 3)

		var opentype_features = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 4))))
		defer pointers.End(gd.InternalDictionary(opentype_features))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 5))))
		defer pointers.End(gd.InternalString(language))
		var meta = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 6))
		defer pointers.End(meta)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, text.String(), gd.ArrayAs[[]RID.Any](gd.InternalArray(fonts)), int(size), gd.DictionaryAs[map[any]any](opentype_features), language.String(), meta.Interface())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (Instance) _shaped_text_add_object(impl func(ptr unsafe.Pointer, shaped RID.Any, key any, size Vector2.XY, inline_align InlineAlignment, length int, baseline Float.X) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		var size = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var inline_align = gd.UnsafeGet[InlineAlignment](p_args, 3)

		var length = gd.UnsafeGet[gd.Int](p_args, 4)

		var baseline = gd.UnsafeGet[gd.Float](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key.Interface(), size, inline_align, int(length), Float.X(baseline))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets new size and alignment of embedded object.
*/
func (Instance) _shaped_text_resize_object(impl func(ptr unsafe.Pointer, shaped RID.Any, key any, size Vector2.XY, inline_align InlineAlignment, baseline Float.X) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		var size = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var inline_align = gd.UnsafeGet[InlineAlignment](p_args, 3)

		var baseline = gd.UnsafeGet[gd.Float](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key.Interface(), size, inline_align, Float.X(baseline))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns number of text spans added using [method _shaped_text_add_string] or [method _shaped_text_add_object].
*/
func (Instance) _shaped_get_span_count(impl func(ptr unsafe.Pointer, shaped RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Returns text span metadata.
*/
func (Instance) _shaped_get_span_meta(impl func(ptr unsafe.Pointer, shaped RID.Any, index int) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var index = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(index))
		ptr, ok := pointers.End(gd.NewVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Changes text span font, font size, and OpenType features, without changing the text.
*/
func (Instance) _shaped_set_span_update_font(impl func(ptr unsafe.Pointer, shaped RID.Any, index int, fonts []RID.Any, size int, opentype_features map[any]any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var index = gd.UnsafeGet[gd.Int](p_args, 1)

		var fonts = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalArray(fonts))
		var size = gd.UnsafeGet[gd.Int](p_args, 3)

		var opentype_features = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 4))))
		defer pointers.End(gd.InternalDictionary(opentype_features))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, int(index), gd.ArrayAs[[]RID.Any](gd.InternalArray(fonts)), int(size), gd.DictionaryAs[map[any]any](opentype_features))
	}
}

/*
[b]Required.[/b]
Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
*/
func (Instance) _shaped_text_substr(impl func(ptr unsafe.Pointer, shaped RID.Any, start int, length int) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var start = gd.UnsafeGet[gd.Int](p_args, 1)

		var length = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(start), int(length))
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
[b]Required.[/b]
Returns the parent buffer from which the substring originates.
*/
func (Instance) _shaped_text_get_parent(impl func(ptr unsafe.Pointer, shaped RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
[b]Optional.[/b]
Adjusts text width to fit to specified width, returns new text width.
*/
func (Instance) _shaped_text_fit_to_width(impl func(ptr unsafe.Pointer, shaped RID.Any, width Float.X, justification_flags gdclass.TextServerJustificationFlag) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var width = gd.UnsafeGet[gd.Float](p_args, 1)

		var justification_flags = gd.UnsafeGet[gdclass.TextServerJustificationFlag](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, Float.X(width), justification_flags)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Optional.[/b]
Aligns shaped text to the given tab-stops.
*/
func (Instance) _shaped_text_tab_align(impl func(ptr unsafe.Pointer, shaped RID.Any, tab_stops []float32) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var tab_stops = pointers.New[gd.PackedFloat32Array](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(tab_stops)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, tab_stops.AsSlice())
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
*/
func (Instance) _shaped_text_shape(impl func(ptr unsafe.Pointer, shaped RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Updates break points in the shaped text. This method is called by default implementation of text breaking functions.
*/
func (Instance) _shaped_text_update_breaks(impl func(ptr unsafe.Pointer, shaped RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Updates justification points in the shaped text. This method is called by default implementation of text justification functions.
*/
func (Instance) _shaped_text_update_justification_ops(impl func(ptr unsafe.Pointer, shaped RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if buffer is successfully shaped.
*/
func (Instance) _shaped_text_is_ready(impl func(ptr unsafe.Pointer, shaped RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns an array of glyphs in the visual order.
*/
func (Instance) _shaped_text_get_glyphs(impl func(ptr unsafe.Pointer, shaped RID.Any) *Glyph) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns text glyphs in the logical order.
*/
func (Instance) _shaped_text_sort_logical(impl func(ptr unsafe.Pointer, shaped RID.Any) *Glyph) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns number of glyphs in the buffer.
*/
func (Instance) _shaped_text_get_glyph_count(impl func(ptr unsafe.Pointer, shaped RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Returns substring buffer character range in the parent buffer.
*/
func (Instance) _shaped_text_get_range(impl func(ptr unsafe.Pointer, shaped RID.Any) Vector2i.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Vector2i(ret))
	}
}

/*
[b]Optional.[/b]
Breaks text to the lines and columns. Returns character ranges for each segment.
*/
func (Instance) _shaped_text_get_line_breaks_adv(impl func(ptr unsafe.Pointer, shaped RID.Any, width []float32, start int, once bool, break_flags gdclass.TextServerLineBreakFlag) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var width = pointers.New[gd.PackedFloat32Array](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(width)
		var start = gd.UnsafeGet[gd.Int](p_args, 2)

		var once = gd.UnsafeGet[bool](p_args, 3)

		var break_flags = gd.UnsafeGet[gdclass.TextServerLineBreakFlag](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, width.AsSlice(), int(start), once, break_flags)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Breaks text to the lines and returns character ranges for each line.
*/
func (Instance) _shaped_text_get_line_breaks(impl func(ptr unsafe.Pointer, shaped RID.Any, width Float.X, start int, break_flags gdclass.TextServerLineBreakFlag) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var width = gd.UnsafeGet[gd.Float](p_args, 1)

		var start = gd.UnsafeGet[gd.Int](p_args, 2)

		var break_flags = gd.UnsafeGet[gdclass.TextServerLineBreakFlag](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, Float.X(width), int(start), break_flags)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum TextServer.GraphemeFlag]).
*/
func (Instance) _shaped_text_get_word_breaks(impl func(ptr unsafe.Pointer, shaped RID.Any, grapheme_flags gdclass.TextServerGraphemeFlag, skip_grapheme_flags gdclass.TextServerGraphemeFlag) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var grapheme_flags = gd.UnsafeGet[gdclass.TextServerGraphemeFlag](p_args, 1)

		var skip_grapheme_flags = gd.UnsafeGet[gdclass.TextServerGraphemeFlag](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, grapheme_flags, skip_grapheme_flags)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns the position of the overrun trim.
*/
func (Instance) _shaped_text_get_trim_pos(impl func(ptr unsafe.Pointer, shaped RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Returns position of the ellipsis.
*/
func (Instance) _shaped_text_get_ellipsis_pos(impl func(ptr unsafe.Pointer, shaped RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Returns number of glyphs in the ellipsis.
*/
func (Instance) _shaped_text_get_ellipsis_glyph_count(impl func(ptr unsafe.Pointer, shaped RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Returns array of the glyphs in the ellipsis.
*/
func (Instance) _shaped_text_get_ellipsis_glyphs(impl func(ptr unsafe.Pointer, shaped RID.Any) *Glyph) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Trims text if it exceeds the given width.
*/
func (Instance) _shaped_text_overrun_trim_to_width(impl func(ptr unsafe.Pointer, shaped RID.Any, width Float.X, trim_flags gdclass.TextServerTextOverrunFlag)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var width = gd.UnsafeGet[gd.Float](p_args, 1)

		var trim_flags = gd.UnsafeGet[gdclass.TextServerTextOverrunFlag](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, Float.X(width), trim_flags)
	}
}

/*
[b]Required.[/b]
Returns array of inline objects.
*/
func (Instance) _shaped_text_get_objects(impl func(ptr unsafe.Pointer, shaped RID.Any) []any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		ptr, ok := pointers.End(gd.InternalArray(gd.EngineArrayFromSlice(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns bounding rectangle of the inline object.
*/
func (Instance) _shaped_text_get_object_rect(impl func(ptr unsafe.Pointer, shaped RID.Any, key any) Rect2.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key.Interface())
		gd.UnsafeSet(p_back, gd.Rect2(ret))
	}
}

/*
[b]Required.[/b]
Returns the character range of the inline object.
*/
func (Instance) _shaped_text_get_object_range(impl func(ptr unsafe.Pointer, shaped RID.Any, key any) Vector2i.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key.Interface())
		gd.UnsafeSet(p_back, gd.Vector2i(ret))
	}
}

/*
[b]Required.[/b]
Returns the glyph index of the inline object.
*/
func (Instance) _shaped_text_get_object_glyph(impl func(ptr unsafe.Pointer, shaped RID.Any, key any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key.Interface())
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Required.[/b]
Returns size of the text.
*/
func (Instance) _shaped_text_get_size(impl func(ptr unsafe.Pointer, shaped RID.Any) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
[b]Required.[/b]
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (Instance) _shaped_text_get_ascent(impl func(ptr unsafe.Pointer, shaped RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (Instance) _shaped_text_get_descent(impl func(ptr unsafe.Pointer, shaped RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
func (Instance) _shaped_text_get_width(impl func(ptr unsafe.Pointer, shaped RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Returns pixel offset of the underline below the baseline.
*/
func (Instance) _shaped_text_get_underline_position(impl func(ptr unsafe.Pointer, shaped RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Required.[/b]
Returns thickness of the underline.
*/
func (Instance) _shaped_text_get_underline_thickness(impl func(ptr unsafe.Pointer, shaped RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
[b]Optional.[/b]
Returns dominant direction of in the range of text.
*/
func (Instance) _shaped_text_get_dominant_direction_in_range(impl func(ptr unsafe.Pointer, shaped RID.Any, start int, end int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var start = gd.UnsafeGet[gd.Int](p_args, 1)

		var end = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(start), int(end))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
*/
func (Instance) _shaped_text_get_carets(impl func(ptr unsafe.Pointer, shaped RID.Any, position int, caret *CaretInfo)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var position = gd.UnsafeGet[gd.Int](p_args, 1)

		var caret = gd.UnsafeGet[*CaretInfo](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, int(position), caret)
	}
}

/*
[b]Optional.[/b]
Returns selection rectangles for the specified character range.
*/
func (Instance) _shaped_text_get_selection(impl func(ptr unsafe.Pointer, shaped RID.Any, start int, end int) []Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var start = gd.UnsafeGet[gd.Int](p_args, 1)

		var end = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(start), int(end))
		ptr, ok := pointers.End(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&ret))))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
*/
func (Instance) _shaped_text_hit_test_grapheme(impl func(ptr unsafe.Pointer, shaped RID.Any, coord Float.X) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var coord = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, Float.X(coord))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
func (Instance) _shaped_text_hit_test_position(impl func(ptr unsafe.Pointer, shaped RID.Any, coord Float.X) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var coord = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, Float.X(coord))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (Instance) _shaped_text_draw(impl func(ptr unsafe.Pointer, shaped RID.Any, canvas RID.Any, pos Vector2.XY, clip_l Float.X, clip_r Float.X, color Color.RGBA)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var canvas = gd.UnsafeGet[gd.RID](p_args, 1)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var clip_l = gd.UnsafeGet[gd.Float](p_args, 3)

		var clip_r = gd.UnsafeGet[gd.Float](p_args, 4)

		var color = gd.UnsafeGet[gd.Color](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, canvas, pos, Float.X(clip_l), Float.X(clip_r), color)
	}
}

/*
[b]Optional.[/b]
Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (Instance) _shaped_text_draw_outline(impl func(ptr unsafe.Pointer, shaped RID.Any, canvas RID.Any, pos Vector2.XY, clip_l Float.X, clip_r Float.X, outline_size int, color Color.RGBA)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var canvas = gd.UnsafeGet[gd.RID](p_args, 1)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var clip_l = gd.UnsafeGet[gd.Float](p_args, 3)

		var clip_r = gd.UnsafeGet[gd.Float](p_args, 4)

		var outline_size = gd.UnsafeGet[gd.Int](p_args, 5)

		var color = gd.UnsafeGet[gd.Color](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, canvas, pos, Float.X(clip_l), Float.X(clip_r), int(outline_size), color)
	}
}

/*
[b]Optional.[/b]
Returns composite character's bounds as offsets from the start of the line.
*/
func (Instance) _shaped_text_get_grapheme_bounds(impl func(ptr unsafe.Pointer, shaped RID.Any, pos int) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
[b]Optional.[/b]
Returns grapheme end position closest to the [param pos].
*/
func (Instance) _shaped_text_next_grapheme_pos(impl func(ptr unsafe.Pointer, shaped RID.Any, pos int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Returns grapheme start position closest to the [param pos].
*/
func (Instance) _shaped_text_prev_grapheme_pos(impl func(ptr unsafe.Pointer, shaped RID.Any, pos int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Returns array of the composite character boundaries.
*/
func (Instance) _shaped_text_get_character_breaks(impl func(ptr unsafe.Pointer, shaped RID.Any) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns composite character end position closest to the [param pos].
*/
func (Instance) _shaped_text_next_character_pos(impl func(ptr unsafe.Pointer, shaped RID.Any, pos int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Returns composite character start position closest to the [param pos].
*/
func (Instance) _shaped_text_prev_character_pos(impl func(ptr unsafe.Pointer, shaped RID.Any, pos int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Returns composite character position closest to the [param pos].
*/
func (Instance) _shaped_text_closest_character_pos(impl func(ptr unsafe.Pointer, shaped RID.Any, pos int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
*/
func (Instance) _format_number(impl func(ptr unsafe.Pointer, number string, language string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var number = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(number))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, number.String(), language.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
*/
func (Instance) _parse_number(impl func(ptr unsafe.Pointer, number string, language string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var number = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(number))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, number.String(), language.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns percent sign used in the [param language].
*/
func (Instance) _percent_sign(impl func(ptr unsafe.Pointer, language string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, language.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Strips diacritics from the string.
*/
func (Instance) _strip_diacritics(impl func(ptr unsafe.Pointer, s string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if [param string] is a valid identifier.
*/
func (Instance) _is_valid_identifier(impl func(ptr unsafe.Pointer, s string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _is_valid_letter(impl func(ptr unsafe.Pointer, unicode int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var unicode = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(unicode))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
*/
func (Instance) _string_get_word_breaks(impl func(ptr unsafe.Pointer, s string, language string, chars_per_line int) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		var chars_per_line = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String(), int(chars_per_line))
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns array of the composite character boundaries.
*/
func (Instance) _string_get_character_breaks(impl func(ptr unsafe.Pointer, s string, language string) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String())
		ptr, ok := pointers.End(gd.NewPackedInt32Slice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
*/
func (Instance) _is_confusable(impl func(ptr unsafe.Pointer, s string, dict []string) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var dict = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(dict)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), dict.Strings())
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
*/
func (Instance) _spoof_check(impl func(ptr unsafe.Pointer, s string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns the string converted to uppercase.
*/
func (Instance) _string_to_upper(impl func(ptr unsafe.Pointer, s string, language string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the string converted to lowercase.
*/
func (Instance) _string_to_lower(impl func(ptr unsafe.Pointer, s string, language string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the string converted to title case.
*/
func (Instance) _string_to_title(impl func(ptr unsafe.Pointer, s string, language string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Default implementation of the BiDi algorithm override function. See [enum TextServer.StructuredTextParser] for more info.
*/
func (Instance) _parse_structured_text(impl func(ptr unsafe.Pointer, parser_type gdclass.TextServerStructuredTextParser, args []any, text string) []Vector3i.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var parser_type = gd.UnsafeGet[gdclass.TextServerStructuredTextParser](p_args, 0)

		var args = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalArray(args))
		var text = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(text))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parser_type, gd.ArrayAs[[]any](gd.InternalArray(args)), text.String())
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[gd.Vector3i]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
This method is called before text server is unregistered.
*/
func (Instance) _cleanup(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextServerExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextServerExtension"))
	casted := Instance{*(*gdclass.TextServerExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
[b]Required.[/b]
Returns [code]true[/code] if the server supports a feature.
*/
func (class) _has_feature(impl func(ptr unsafe.Pointer, feature gdclass.TextServerFeature) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var feature = gd.UnsafeGet[gdclass.TextServerFeature](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, feature)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns the name of the server interface.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns text server features, see [enum TextServer.Feature].
*/
func (class) _get_features(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Frees an object created by this [TextServer].
*/
func (class) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
*/
func (class) _has(impl func(ptr unsafe.Pointer, rid gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
*/
func (class) _load_support_data(impl func(ptr unsafe.Pointer, filename String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var filename = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(filename))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, filename)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
*/
func (class) _get_support_data_filename(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
*/
func (class) _get_support_data_info(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
*/
func (class) _save_support_data(impl func(ptr unsafe.Pointer, filename String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var filename = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(filename))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, filename)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if locale is right-to-left.
*/
func (class) _is_locale_right_to_left(impl func(ptr unsafe.Pointer, locale String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var locale = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(locale))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, locale)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Converts readable feature, variation, script, or language name to OpenType tag.
*/
func (class) _name_to_tag(impl func(ptr unsafe.Pointer, name String.Readable) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(name))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Converts OpenType tag to readable feature, variation, script, or language name.
*/
func (class) _tag_to_name(impl func(ptr unsafe.Pointer, tag gd.Int) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var tag = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, tag)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Creates a new, empty font cache entry resource.
*/
func (class) _create_font(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Optional, implement if font supports extra spacing or baseline offset.
Creates a new variation existing font which is reusing the same glyph cache and font data.
*/
func (class) _create_font_linked_variation(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets font source data, e.g contents of the dynamic font source file.
*/
func (class) _font_set_data(impl func(ptr unsafe.Pointer, font_rid gd.RID, data gd.PackedByteArray)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var data = pointers.New[gd.PackedByteArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(data)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, data)
	}
}

/*
[b]Optional.[/b]
Sets pointer to the font source data, e.g contents of the dynamic font source file.
*/
func (class) _font_set_data_ptr(impl func(ptr unsafe.Pointer, font_rid gd.RID, data_ptr unsafe.Pointer, data_size gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var data_ptr = gd.UnsafeGet[unsafe.Pointer](p_args, 1)

		var data_size = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, data_ptr, data_size)
	}
}

/*
[b]Optional.[/b]
Sets an active face index in the TrueType / OpenType collection.
*/
func (class) _font_set_face_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, face_index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var face_index = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, face_index)
	}
}

/*
[b]Optional.[/b]
Returns an active face index in the TrueType / OpenType collection.
*/
func (class) _font_get_face_index(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns number of faces in the TrueType / OpenType collection.
*/
func (class) _font_get_face_count(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets the font style flags, see [enum TextServer.FontStyle].
*/
func (class) _font_set_style(impl func(ptr unsafe.Pointer, font_rid gd.RID, style gdclass.TextServerFontStyle)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var style = gd.UnsafeGet[gdclass.TextServerFontStyle](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, style)
	}
}

/*
[b]Optional.[/b]
Returns font style flags, see [enum TextServer.FontStyle].
*/
func (class) _font_get_style(impl func(ptr unsafe.Pointer, font_rid gd.RID) gdclass.TextServerFontStyle) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets the font family name.
*/
func (class) _font_set_name(impl func(ptr unsafe.Pointer, font_rid gd.RID, name String.Readable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(name))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, name)
	}
}

/*
[b]Optional.[/b]
Returns font family name.
*/
func (class) _font_get_name(impl func(ptr unsafe.Pointer, font_rid gd.RID) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
func (class) _font_get_ot_name_strings(impl func(ptr unsafe.Pointer, font_rid gd.RID) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets the font style name.
*/
func (class) _font_set_style_name(impl func(ptr unsafe.Pointer, font_rid gd.RID, name_style String.Readable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var name_style = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(name_style))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, name_style)
	}
}

/*
[b]Optional.[/b]
Returns font style name.
*/
func (class) _font_get_style_name(impl func(ptr unsafe.Pointer, font_rid gd.RID) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (class) _font_set_weight(impl func(ptr unsafe.Pointer, font_rid gd.RID, weight gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var weight = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, weight)
	}
}

/*
[b]Optional.[/b]
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (class) _font_get_weight(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (class) _font_set_stretch(impl func(ptr unsafe.Pointer, font_rid gd.RID, stretch gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var stretch = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, stretch)
	}
}

/*
[b]Optional.[/b]
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (class) _font_get_stretch(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets font anti-aliasing mode.
*/
func (class) _font_set_antialiasing(impl func(ptr unsafe.Pointer, font_rid gd.RID, antialiasing gdclass.TextServerFontAntialiasing)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var antialiasing = gd.UnsafeGet[gdclass.TextServerFontAntialiasing](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, antialiasing)
	}
}

/*
[b]Optional.[/b]
Returns font anti-aliasing mode.
*/
func (class) _font_get_antialiasing(impl func(ptr unsafe.Pointer, font_rid gd.RID) gdclass.TextServerFontAntialiasing) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], embedded font bitmap loading is disabled.
*/
func (class) _font_set_disable_embedded_bitmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID, disable_embedded_bitmaps bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var disable_embedded_bitmaps = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, disable_embedded_bitmaps)
	}
}

/*
[b]Optional.[/b]
Returns whether the font's embedded bitmap loading is disabled.
*/
func (class) _font_get_disable_embedded_bitmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] font texture mipmap generation is enabled.
*/
func (class) _font_set_generate_mipmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID, generate_mipmaps bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var generate_mipmaps = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, generate_mipmaps)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if font texture mipmap generation is enabled.
*/
func (class) _font_get_generate_mipmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
*/
func (class) _font_set_multichannel_signed_distance_field(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var msdf = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, msdf)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
*/
func (class) _font_is_multichannel_signed_distance_field(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (class) _font_set_msdf_pixel_range(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf_pixel_range gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var msdf_pixel_range = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, msdf_pixel_range)
	}
}

/*
[b]Optional.[/b]
Returns the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (class) _font_get_msdf_pixel_range(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets source font size used to generate MSDF textures.
*/
func (class) _font_set_msdf_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf_size gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var msdf_size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, msdf_size)
	}
}

/*
[b]Optional.[/b]
Returns source font size used to generate MSDF textures.
*/
func (class) _font_get_msdf_size(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
*/
func (class) _font_set_fixed_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, fixed_size gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var fixed_size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, fixed_size)
	}
}

/*
[b]Required.[/b]
Returns bitmap font fixed size.
*/
func (class) _font_get_fixed_size(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
*/
func (class) _font_set_fixed_size_scale_mode(impl func(ptr unsafe.Pointer, font_rid gd.RID, fixed_size_scale_mode gdclass.TextServerFixedSizeScaleMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var fixed_size_scale_mode = gd.UnsafeGet[gdclass.TextServerFixedSizeScaleMode](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, fixed_size_scale_mode)
	}
}

/*
[b]Required.[/b]
Returns bitmap font scaling mode.
*/
func (class) _font_get_fixed_size_scale_mode(impl func(ptr unsafe.Pointer, font_rid gd.RID) gdclass.TextServerFixedSizeScaleMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], system fonts can be automatically used as fallbacks.
*/
func (class) _font_set_allow_system_fallback(impl func(ptr unsafe.Pointer, font_rid gd.RID, allow_system_fallback bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var allow_system_fallback = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, allow_system_fallback)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
*/
func (class) _font_is_allow_system_fallback(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
*/
func (class) _font_set_force_autohinter(impl func(ptr unsafe.Pointer, font_rid gd.RID, force_autohinter bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var force_autohinter = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, force_autohinter)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting.
*/
func (class) _font_is_force_autohinter(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets font hinting mode. Used by dynamic fonts only.
*/
func (class) _font_set_hinting(impl func(ptr unsafe.Pointer, font_rid gd.RID, hinting gdclass.TextServerHinting)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var hinting = gd.UnsafeGet[gdclass.TextServerHinting](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, hinting)
	}
}

/*
[b]Optional.[/b]
Returns the font hinting mode. Used by dynamic fonts only.
*/
func (class) _font_get_hinting(impl func(ptr unsafe.Pointer, font_rid gd.RID) gdclass.TextServerHinting) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets font subpixel glyph positioning mode.
*/
func (class) _font_set_subpixel_positioning(impl func(ptr unsafe.Pointer, font_rid gd.RID, subpixel_positioning gdclass.TextServerSubpixelPositioning)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var subpixel_positioning = gd.UnsafeGet[gdclass.TextServerSubpixelPositioning](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, subpixel_positioning)
	}
}

/*
[b]Optional.[/b]
Returns font subpixel glyph positioning mode.
*/
func (class) _font_get_subpixel_positioning(impl func(ptr unsafe.Pointer, font_rid gd.RID) gdclass.TextServerSubpixelPositioning) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (class) _font_set_embolden(impl func(ptr unsafe.Pointer, font_rid gd.RID, strength gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var strength = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, strength)
	}
}

/*
[b]Optional.[/b]
Returns font embolden strength.
*/
func (class) _font_get_embolden(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
func (class) _font_set_spacing(impl func(ptr unsafe.Pointer, font_rid gd.RID, spacing gdclass.TextServerSpacingType, value gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var spacing = gd.UnsafeGet[gdclass.TextServerSpacingType](p_args, 1)

		var value = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, spacing, value)
	}
}

/*
[b]Optional.[/b]
Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
func (class) _font_get_spacing(impl func(ptr unsafe.Pointer, font_rid gd.RID, spacing gdclass.TextServerSpacingType) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var spacing = gd.UnsafeGet[gdclass.TextServerSpacingType](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, spacing)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets extra baseline offset (as a fraction of font height).
*/
func (class) _font_set_baseline_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, baseline_offset gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var baseline_offset = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, baseline_offset)
	}
}

/*
[b]Optional.[/b]
Returns extra baseline offset (as a fraction of font height).
*/
func (class) _font_get_baseline_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
func (class) _font_set_transform(impl func(ptr unsafe.Pointer, font_rid gd.RID, transform gd.Transform2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, transform)
	}
}

/*
[b]Optional.[/b]
Returns 2D transform applied to the font outlines.
*/
func (class) _font_get_transform(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Transform2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets variation coordinates for the specified font cache entry.
*/
func (class) _font_set_variation_coordinates(impl func(ptr unsafe.Pointer, font_rid gd.RID, variation_coordinates Dictionary.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var variation_coordinates = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalDictionary(variation_coordinates))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, variation_coordinates)
	}
}

/*
[b]Optional.[/b]
Returns variation coordinates for the specified font cache entry.
*/
func (class) _font_get_variation_coordinates(impl func(ptr unsafe.Pointer, font_rid gd.RID) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (class) _font_set_oversampling(impl func(ptr unsafe.Pointer, font_rid gd.RID, oversampling gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var oversampling = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, oversampling)
	}
}

/*
[b]Optional.[/b]
Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (class) _font_get_oversampling(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
func (class) _font_get_size_cache_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) Array.Contains[gd.Vector2i]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Removes all font sizes from the cache entry.
*/
func (class) _font_clear_size_cache(impl func(ptr unsafe.Pointer, font_rid gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid)
	}
}

/*
[b]Required.[/b]
Removes specified font size from the cache entry.
*/
func (class) _font_remove_size_cache(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size)
	}
}

/*
[b]Required.[/b]
Sets the font ascent (number of pixels above the baseline).
*/
func (class) _font_set_ascent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, ascent gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var ascent = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, ascent)
	}
}

/*
[b]Required.[/b]
Returns the font ascent (number of pixels above the baseline).
*/
func (class) _font_get_ascent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets the font descent (number of pixels below the baseline).
*/
func (class) _font_set_descent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, descent gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var descent = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, descent)
	}
}

/*
[b]Required.[/b]
Returns the font descent (number of pixels below the baseline).
*/
func (class) _font_get_descent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets pixel offset of the underline below the baseline.
*/
func (class) _font_set_underline_position(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, underline_position gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var underline_position = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, underline_position)
	}
}

/*
[b]Required.[/b]
Returns pixel offset of the underline below the baseline.
*/
func (class) _font_get_underline_position(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets thickness of the underline in pixels.
*/
func (class) _font_set_underline_thickness(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, underline_thickness gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var underline_thickness = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, underline_thickness)
	}
}

/*
[b]Required.[/b]
Returns thickness of the underline in pixels.
*/
func (class) _font_get_underline_thickness(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets scaling factor of the color bitmap font.
*/
func (class) _font_set_scale(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, scale gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var scale = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, scale)
	}
}

/*
[b]Required.[/b]
Returns scaling factor of the color bitmap font.
*/
func (class) _font_get_scale(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns number of textures used by font cache entry.
*/
func (class) _font_get_texture_count(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Removes all textures from font cache entry.
*/
func (class) _font_clear_textures(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size)
	}
}

/*
[b]Required.[/b]
Removes specified texture from the cache entry.
*/
func (class) _font_remove_texture(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, texture_index)
	}
}

/*
[b]Required.[/b]
Sets font cache texture image data.
*/
func (class) _font_set_texture_image(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, image [1]gdclass.Image)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		var image = [1]gdclass.Image{pointers.New[gdclass.Image]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}

		defer pointers.End(image[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, texture_index, image)
	}
}

/*
[b]Required.[/b]
Returns font cache texture image data.
*/
func (class) _font_get_texture_image(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) [1]gdclass.Image) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, texture_index)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets array containing glyph packing data.
*/
func (class) _font_set_texture_offsets(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, offset gd.PackedInt32Array)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		var offset = pointers.New[gd.PackedInt32Array](gd.UnsafeGet[gd.PackedPointers](p_args, 3))
		defer pointers.End(offset)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, texture_index, offset)
	}
}

/*
[b]Optional.[/b]
Returns array containing glyph packing data.
*/
func (class) _font_get_texture_offsets(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var texture_index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, texture_index)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns list of rendered glyphs in the cache entry.
*/
func (class) _font_get_glyph_list(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Removes all rendered glyph information from the cache entry.
*/
func (class) _font_clear_glyphs(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size)
	}
}

/*
[b]Required.[/b]
Removes specified rendered glyph information from the cache entry.
*/
func (class) _font_remove_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, glyph)
	}
}

/*
[b]Required.[/b]
Returns glyph advance (offset of the next glyph).
*/
func (class) _font_get_glyph_advance(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets glyph advance (offset of the next glyph).
*/
func (class) _font_set_glyph_advance(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph gd.Int, advance gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var advance = gd.UnsafeGet[gd.Vector2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, glyph, advance)
	}
}

/*
[b]Required.[/b]
Returns glyph offset from the baseline.
*/
func (class) _font_get_glyph_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets glyph offset from the baseline.
*/
func (class) _font_set_glyph_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int, offset gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var offset = gd.UnsafeGet[gd.Vector2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, glyph, offset)
	}
}

/*
[b]Required.[/b]
Returns size of the glyph.
*/
func (class) _font_get_glyph_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets size of the glyph.
*/
func (class) _font_set_glyph_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int, gl_size gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var gl_size = gd.UnsafeGet[gd.Vector2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, glyph, gl_size)
	}
}

/*
[b]Required.[/b]
Returns rectangle in the cache texture containing the glyph.
*/
func (class) _font_get_glyph_uv_rect(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Rect2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets rectangle in the cache texture containing the glyph.
*/
func (class) _font_set_glyph_uv_rect(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int, uv_rect gd.Rect2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var uv_rect = gd.UnsafeGet[gd.Rect2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, glyph, uv_rect)
	}
}

/*
[b]Required.[/b]
Returns index of the cache texture containing the glyph.
*/
func (class) _font_get_glyph_texture_idx(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets index of the cache texture containing the glyph.
*/
func (class) _font_set_glyph_texture_idx(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int, texture_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		var texture_idx = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, glyph, texture_idx)
	}
}

/*
[b]Required.[/b]
Returns resource ID of the cache texture containing the glyph.
*/
func (class) _font_get_glyph_texture_rid(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns size of the cache texture containing the glyph.
*/
func (class) _font_get_glyph_texture_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var glyph = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns outline contours of the glyph.
*/
func (class) _font_get_glyph_contours(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, index gd.Int) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, index)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns list of the kerning overrides.
*/
func (class) _font_get_kerning_list(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) Array.Contains[gd.Vector2i]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Removes all kerning overrides.
*/
func (class) _font_clear_kerning_map(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size)
	}
}

/*
[b]Optional.[/b]
Removes kerning override for the pair of glyphs.
*/
func (class) _font_remove_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, glyph_pair)
	}
}

/*
[b]Optional.[/b]
Sets kerning for the pair of glyphs.
*/
func (class) _font_set_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i, kerning gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args, 2)

		var kerning = gd.UnsafeGet[gd.Vector2](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, glyph_pair, kerning)
	}
}

/*
[b]Optional.[/b]
Returns kerning for the pair of glyphs.
*/
func (class) _font_get_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph_pair)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
func (class) _font_get_glyph_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, char gd.Int, variation_selector gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var char = gd.UnsafeGet[gd.Int](p_args, 2)

		var variation_selector = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, char, variation_selector)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid.
*/
func (class) _font_get_char_from_glyph_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph_index gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var glyph_index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph_index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
func (class) _font_has_char(impl func(ptr unsafe.Pointer, font_rid gd.RID, char gd.Int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var char = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, char)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns a string containing all the characters available in the font.
*/
func (class) _font_get_supported_chars(impl func(ptr unsafe.Pointer, font_rid gd.RID) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Renders the range of characters to the font cache texture.
*/
func (class) _font_render_range(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, start gd.Int, end gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var start = gd.UnsafeGet[gd.Int](p_args, 2)

		var end = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, start, end)
	}
}

/*
[b]Optional.[/b]
Renders specified glyph to the font cache texture.
*/
func (class) _font_render_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, index gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)

		var index = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, size, index)
	}
}

/*
[b]Required.[/b]
Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
*/
func (class) _font_draw_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var canvas = gd.UnsafeGet[gd.RID](p_args, 1)

		var size = gd.UnsafeGet[gd.Int](p_args, 2)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 3)

		var index = gd.UnsafeGet[gd.Int](p_args, 4)

		var color = gd.UnsafeGet[gd.Color](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, canvas, size, pos, index, color)
	}
}

/*
[b]Required.[/b]
Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
*/
func (class) _font_draw_glyph_outline(impl func(ptr unsafe.Pointer, font_rid gd.RID, canvas gd.RID, size gd.Int, outline_size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var canvas = gd.UnsafeGet[gd.RID](p_args, 1)

		var size = gd.UnsafeGet[gd.Int](p_args, 2)

		var outline_size = gd.UnsafeGet[gd.Int](p_args, 3)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 4)

		var index = gd.UnsafeGet[gd.Int](p_args, 5)

		var color = gd.UnsafeGet[gd.Color](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, canvas, size, outline_size, pos, index, color)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
func (class) _font_is_language_supported(impl func(ptr unsafe.Pointer, font_rid gd.RID, language String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, language)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Adds override for [method _font_is_language_supported].
*/
func (class) _font_set_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language String.Readable, supported bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		var supported = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, language, supported)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if support override is enabled for the [param language].
*/
func (class) _font_get_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, language)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Remove language support override.
*/
func (class) _font_remove_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language String.Readable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, language)
	}
}

/*
[b]Optional.[/b]
Returns list of language support overrides.
*/
func (class) _font_get_language_support_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code], if font supports given script (ISO 15924 code).
*/
func (class) _font_is_script_supported(impl func(ptr unsafe.Pointer, font_rid gd.RID, script String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(script))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, script)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Adds override for [method _font_is_script_supported].
*/
func (class) _font_set_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script String.Readable, supported bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(script))
		var supported = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, script, supported)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if support override is enabled for the [param script].
*/
func (class) _font_get_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(script))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, script)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Removes script support override.
*/
func (class) _font_remove_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script String.Readable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(script))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, script)
	}
}

/*
[b]Optional.[/b]
Returns list of script support overrides.
*/
func (class) _font_get_script_support_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets font OpenType feature set override.
*/
func (class) _font_set_opentype_feature_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID, overrides Dictionary.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var overrides = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalDictionary(overrides))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, font_rid, overrides)
	}
}

/*
[b]Optional.[/b]
Returns font OpenType feature set override.
*/
func (class) _font_get_opentype_feature_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the dictionary of the supported OpenType features.
*/
func (class) _font_supported_feature_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the dictionary of the supported OpenType variation coordinates.
*/
func (class) _font_supported_variation_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var font_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the font oversampling factor, shared by all fonts in the TextServer.
*/
func (class) _font_get_global_oversampling(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets oversampling factor, shared by all font in the TextServer.
*/
func (class) _font_set_global_oversampling(impl func(ptr unsafe.Pointer, oversampling gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var oversampling = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, oversampling)
	}
}

/*
[b]Optional.[/b]
Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
*/
func (class) _get_hex_code_box_size(impl func(ptr unsafe.Pointer, size gd.Int, index gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var size = gd.UnsafeGet[gd.Int](p_args, 0)

		var index = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, size, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Draws box displaying character hexadecimal code.
*/
func (class) _draw_hex_code_box(impl func(ptr unsafe.Pointer, canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var canvas = gd.UnsafeGet[gd.RID](p_args, 0)

		var size = gd.UnsafeGet[gd.Int](p_args, 1)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var index = gd.UnsafeGet[gd.Int](p_args, 3)

		var color = gd.UnsafeGet[gd.Color](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, canvas, size, pos, index, color)
	}
}

/*
[b]Required.[/b]
Creates a new buffer for complex text layout, with the given [param direction] and [param orientation].
*/
func (class) _create_shaped_text(impl func(ptr unsafe.Pointer, direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var direction = gd.UnsafeGet[gdclass.TextServerDirection](p_args, 0)

		var orientation = gd.UnsafeGet[gdclass.TextServerOrientation](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, direction, orientation)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Clears text buffer (removes text and inline objects).
*/
func (class) _shaped_text_clear(impl func(ptr unsafe.Pointer, shaped gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped)
	}
}

/*
[b]Optional.[/b]
Sets desired text direction. If set to [constant TextServer.DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
*/
func (class) _shaped_text_set_direction(impl func(ptr unsafe.Pointer, shaped gd.RID, direction gdclass.TextServerDirection)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var direction = gd.UnsafeGet[gdclass.TextServerDirection](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, direction)
	}
}

/*
[b]Optional.[/b]
Returns direction of the text.
*/
func (class) _shaped_text_get_direction(impl func(ptr unsafe.Pointer, shaped gd.RID) gdclass.TextServerDirection) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns direction of the text, inferred by the BiDi algorithm.
*/
func (class) _shaped_text_get_inferred_direction(impl func(ptr unsafe.Pointer, shaped gd.RID) gdclass.TextServerDirection) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Overrides BiDi for the structured text.
*/
func (class) _shaped_text_set_bidi_override(impl func(ptr unsafe.Pointer, shaped gd.RID, override Array.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var override = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalArray(override))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, override)
	}
}

/*
[b]Optional.[/b]
Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (class) _shaped_text_set_custom_punctuation(impl func(ptr unsafe.Pointer, shaped gd.RID, punct String.Readable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var punct = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(punct))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, punct)
	}
}

/*
[b]Optional.[/b]
Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (class) _shaped_text_get_custom_punctuation(impl func(ptr unsafe.Pointer, shaped gd.RID) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Sets ellipsis character used for text clipping.
*/
func (class) _shaped_text_set_custom_ellipsis(impl func(ptr unsafe.Pointer, shaped gd.RID, char gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var char = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, char)
	}
}

/*
[b]Optional.[/b]
Returns ellipsis character used for text clipping.
*/
func (class) _shaped_text_get_custom_ellipsis(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets desired text orientation.
*/
func (class) _shaped_text_set_orientation(impl func(ptr unsafe.Pointer, shaped gd.RID, orientation gdclass.TextServerOrientation)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var orientation = gd.UnsafeGet[gdclass.TextServerOrientation](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, orientation)
	}
}

/*
[b]Optional.[/b]
Returns text orientation.
*/
func (class) _shaped_text_get_orientation(impl func(ptr unsafe.Pointer, shaped gd.RID) gdclass.TextServerOrientation) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
*/
func (class) _shaped_text_set_preserve_invalid(impl func(ptr unsafe.Pointer, shaped gd.RID, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var enabled = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, enabled)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
*/
func (class) _shaped_text_get_preserve_invalid(impl func(ptr unsafe.Pointer, shaped gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] text buffer will display control characters.
*/
func (class) _shaped_text_set_preserve_control(impl func(ptr unsafe.Pointer, shaped gd.RID, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var enabled = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, enabled)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if text buffer is configured to display control characters.
*/
func (class) _shaped_text_get_preserve_control(impl func(ptr unsafe.Pointer, shaped gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Sets extra spacing added between glyphs or lines in pixels.
*/
func (class) _shaped_text_set_spacing(impl func(ptr unsafe.Pointer, shaped gd.RID, spacing gdclass.TextServerSpacingType, value gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var spacing = gd.UnsafeGet[gdclass.TextServerSpacingType](p_args, 1)

		var value = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, spacing, value)
	}
}

/*
[b]Optional.[/b]
Returns extra spacing added between glyphs or lines in pixels.
*/
func (class) _shaped_text_get_spacing(impl func(ptr unsafe.Pointer, shaped gd.RID, spacing gdclass.TextServerSpacingType) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var spacing = gd.UnsafeGet[gdclass.TextServerSpacingType](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, spacing)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Adds text span and font to draw it to the text buffer.
*/
func (class) _shaped_text_add_string(impl func(ptr unsafe.Pointer, shaped gd.RID, text String.Readable, fonts Array.Contains[gd.RID], size gd.Int, opentype_features Dictionary.Any, language String.Readable, meta gd.Variant) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var text = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(text))
		var fonts = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalArray(fonts))
		var size = gd.UnsafeGet[gd.Int](p_args, 3)

		var opentype_features = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 4))))
		defer pointers.End(gd.InternalDictionary(opentype_features))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 5))))
		defer pointers.End(gd.InternalString(language))
		var meta = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 6))
		defer pointers.End(meta)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, text, fonts, size, opentype_features, language, meta)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (class) _shaped_text_add_object(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align InlineAlignment, length gd.Int, baseline gd.Float) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		var size = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var inline_align = gd.UnsafeGet[InlineAlignment](p_args, 3)

		var length = gd.UnsafeGet[gd.Int](p_args, 4)

		var baseline = gd.UnsafeGet[gd.Float](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key, size, inline_align, length, baseline)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Sets new size and alignment of embedded object.
*/
func (class) _shaped_text_resize_object(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align InlineAlignment, baseline gd.Float) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		var size = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var inline_align = gd.UnsafeGet[InlineAlignment](p_args, 3)

		var baseline = gd.UnsafeGet[gd.Float](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key, size, inline_align, baseline)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns number of text spans added using [method _shaped_text_add_string] or [method _shaped_text_add_object].
*/
func (class) _shaped_get_span_count(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns text span metadata.
*/
func (class) _shaped_get_span_meta(impl func(ptr unsafe.Pointer, shaped gd.RID, index gd.Int) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var index = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, index)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Changes text span font, font size, and OpenType features, without changing the text.
*/
func (class) _shaped_set_span_update_font(impl func(ptr unsafe.Pointer, shaped gd.RID, index gd.Int, fonts Array.Contains[gd.RID], size gd.Int, opentype_features Dictionary.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var index = gd.UnsafeGet[gd.Int](p_args, 1)

		var fonts = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalArray(fonts))
		var size = gd.UnsafeGet[gd.Int](p_args, 3)

		var opentype_features = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 4))))
		defer pointers.End(gd.InternalDictionary(opentype_features))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, index, fonts, size, opentype_features)
	}
}

/*
[b]Required.[/b]
Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
*/
func (class) _shaped_text_substr(impl func(ptr unsafe.Pointer, shaped gd.RID, start gd.Int, length gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var start = gd.UnsafeGet[gd.Int](p_args, 1)

		var length = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, start, length)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns the parent buffer from which the substring originates.
*/
func (class) _shaped_text_get_parent(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Adjusts text width to fit to specified width, returns new text width.
*/
func (class) _shaped_text_fit_to_width(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.Float, justification_flags gdclass.TextServerJustificationFlag) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var width = gd.UnsafeGet[gd.Float](p_args, 1)

		var justification_flags = gd.UnsafeGet[gdclass.TextServerJustificationFlag](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, width, justification_flags)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Aligns shaped text to the given tab-stops.
*/
func (class) _shaped_text_tab_align(impl func(ptr unsafe.Pointer, shaped gd.RID, tab_stops gd.PackedFloat32Array) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var tab_stops = pointers.New[gd.PackedFloat32Array](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(tab_stops)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, tab_stops)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
*/
func (class) _shaped_text_shape(impl func(ptr unsafe.Pointer, shaped gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Updates break points in the shaped text. This method is called by default implementation of text breaking functions.
*/
func (class) _shaped_text_update_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Updates justification points in the shaped text. This method is called by default implementation of text justification functions.
*/
func (class) _shaped_text_update_justification_ops(impl func(ptr unsafe.Pointer, shaped gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if buffer is successfully shaped.
*/
func (class) _shaped_text_is_ready(impl func(ptr unsafe.Pointer, shaped gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns an array of glyphs in the visual order.
*/
func (class) _shaped_text_get_glyphs(impl func(ptr unsafe.Pointer, shaped gd.RID) *Glyph) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns text glyphs in the logical order.
*/
func (class) _shaped_text_sort_logical(impl func(ptr unsafe.Pointer, shaped gd.RID) *Glyph) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns number of glyphs in the buffer.
*/
func (class) _shaped_text_get_glyph_count(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns substring buffer character range in the parent buffer.
*/
func (class) _shaped_text_get_range(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Vector2i) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Breaks text to the lines and columns. Returns character ranges for each segment.
*/
func (class) _shaped_text_get_line_breaks_adv(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.PackedFloat32Array, start gd.Int, once bool, break_flags gdclass.TextServerLineBreakFlag) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var width = pointers.New[gd.PackedFloat32Array](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(width)
		var start = gd.UnsafeGet[gd.Int](p_args, 2)

		var once = gd.UnsafeGet[bool](p_args, 3)

		var break_flags = gd.UnsafeGet[gdclass.TextServerLineBreakFlag](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, width, start, once, break_flags)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Breaks text to the lines and returns character ranges for each line.
*/
func (class) _shaped_text_get_line_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.Float, start gd.Int, break_flags gdclass.TextServerLineBreakFlag) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var width = gd.UnsafeGet[gd.Float](p_args, 1)

		var start = gd.UnsafeGet[gd.Int](p_args, 2)

		var break_flags = gd.UnsafeGet[gdclass.TextServerLineBreakFlag](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, width, start, break_flags)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum TextServer.GraphemeFlag]).
*/
func (class) _shaped_text_get_word_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID, grapheme_flags gdclass.TextServerGraphemeFlag, skip_grapheme_flags gdclass.TextServerGraphemeFlag) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var grapheme_flags = gd.UnsafeGet[gdclass.TextServerGraphemeFlag](p_args, 1)

		var skip_grapheme_flags = gd.UnsafeGet[gdclass.TextServerGraphemeFlag](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, grapheme_flags, skip_grapheme_flags)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns the position of the overrun trim.
*/
func (class) _shaped_text_get_trim_pos(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns position of the ellipsis.
*/
func (class) _shaped_text_get_ellipsis_pos(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns number of glyphs in the ellipsis.
*/
func (class) _shaped_text_get_ellipsis_glyph_count(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns array of the glyphs in the ellipsis.
*/
func (class) _shaped_text_get_ellipsis_glyphs(impl func(ptr unsafe.Pointer, shaped gd.RID) *Glyph) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Trims text if it exceeds the given width.
*/
func (class) _shaped_text_overrun_trim_to_width(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.Float, trim_flags gdclass.TextServerTextOverrunFlag)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var width = gd.UnsafeGet[gd.Float](p_args, 1)

		var trim_flags = gd.UnsafeGet[gdclass.TextServerTextOverrunFlag](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, width, trim_flags)
	}
}

/*
[b]Required.[/b]
Returns array of inline objects.
*/
func (class) _shaped_text_get_objects(impl func(ptr unsafe.Pointer, shaped gd.RID) Array.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Required.[/b]
Returns bounding rectangle of the inline object.
*/
func (class) _shaped_text_get_object_rect(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) gd.Rect2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns the character range of the inline object.
*/
func (class) _shaped_text_get_object_range(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) gd.Vector2i) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns the glyph index of the inline object.
*/
func (class) _shaped_text_get_object_glyph(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var key = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(key)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns size of the text.
*/
func (class) _shaped_text_get_size(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (class) _shaped_text_get_ascent(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (class) _shaped_text_get_descent(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
func (class) _shaped_text_get_width(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns pixel offset of the underline below the baseline.
*/
func (class) _shaped_text_get_underline_position(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Required.[/b]
Returns thickness of the underline.
*/
func (class) _shaped_text_get_underline_thickness(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns dominant direction of in the range of text.
*/
func (class) _shaped_text_get_dominant_direction_in_range(impl func(ptr unsafe.Pointer, shaped gd.RID, start gd.Int, end gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var start = gd.UnsafeGet[gd.Int](p_args, 1)

		var end = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, start, end)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
*/
func (class) _shaped_text_get_carets(impl func(ptr unsafe.Pointer, shaped gd.RID, position gd.Int, caret *CaretInfo)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var position = gd.UnsafeGet[gd.Int](p_args, 1)

		var caret = gd.UnsafeGet[*CaretInfo](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, position, caret)
	}
}

/*
[b]Optional.[/b]
Returns selection rectangles for the specified character range.
*/
func (class) _shaped_text_get_selection(impl func(ptr unsafe.Pointer, shaped gd.RID, start gd.Int, end gd.Int) gd.PackedVector2Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var start = gd.UnsafeGet[gd.Int](p_args, 1)

		var end = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, start, end)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
*/
func (class) _shaped_text_hit_test_grapheme(impl func(ptr unsafe.Pointer, shaped gd.RID, coord gd.Float) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var coord = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, coord)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
func (class) _shaped_text_hit_test_position(impl func(ptr unsafe.Pointer, shaped gd.RID, coord gd.Float) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var coord = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, coord)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (class) _shaped_text_draw(impl func(ptr unsafe.Pointer, shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, color gd.Color)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var canvas = gd.UnsafeGet[gd.RID](p_args, 1)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var clip_l = gd.UnsafeGet[gd.Float](p_args, 3)

		var clip_r = gd.UnsafeGet[gd.Float](p_args, 4)

		var color = gd.UnsafeGet[gd.Color](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, canvas, pos, clip_l, clip_r, color)
	}
}

/*
[b]Optional.[/b]
Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (class) _shaped_text_draw_outline(impl func(ptr unsafe.Pointer, shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, outline_size gd.Int, color gd.Color)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var canvas = gd.UnsafeGet[gd.RID](p_args, 1)

		var pos = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var clip_l = gd.UnsafeGet[gd.Float](p_args, 3)

		var clip_r = gd.UnsafeGet[gd.Float](p_args, 4)

		var outline_size = gd.UnsafeGet[gd.Int](p_args, 5)

		var color = gd.UnsafeGet[gd.Color](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shaped, canvas, pos, clip_l, clip_r, outline_size, color)
	}
}

/*
[b]Optional.[/b]
Returns composite character's bounds as offsets from the start of the line.
*/
func (class) _shaped_text_get_grapheme_bounds(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns grapheme end position closest to the [param pos].
*/
func (class) _shaped_text_next_grapheme_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns grapheme start position closest to the [param pos].
*/
func (class) _shaped_text_prev_grapheme_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns array of the composite character boundaries.
*/
func (class) _shaped_text_get_character_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns composite character end position closest to the [param pos].
*/
func (class) _shaped_text_next_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns composite character start position closest to the [param pos].
*/
func (class) _shaped_text_prev_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns composite character position closest to the [param pos].
*/
func (class) _shaped_text_closest_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shaped = gd.UnsafeGet[gd.RID](p_args, 0)

		var pos = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
*/
func (class) _format_number(impl func(ptr unsafe.Pointer, number String.Readable, language String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var number = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(number))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, number, language)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
*/
func (class) _parse_number(impl func(ptr unsafe.Pointer, number String.Readable, language String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var number = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(number))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, number, language)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns percent sign used in the [param language].
*/
func (class) _percent_sign(impl func(ptr unsafe.Pointer, language String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, language)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Strips diacritics from the string.
*/
func (class) _strip_diacritics(impl func(ptr unsafe.Pointer, s String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if [param string] is a valid identifier.
*/
func (class) _is_valid_identifier(impl func(ptr unsafe.Pointer, s String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _is_valid_letter(impl func(ptr unsafe.Pointer, unicode gd.Int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var unicode = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, unicode)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
*/
func (class) _string_get_word_breaks(impl func(ptr unsafe.Pointer, s String.Readable, language String.Readable, chars_per_line gd.Int) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		var chars_per_line = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language, chars_per_line)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns array of the composite character boundaries.
*/
func (class) _string_get_character_breaks(impl func(ptr unsafe.Pointer, s String.Readable, language String.Readable) gd.PackedInt32Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
*/
func (class) _is_confusable(impl func(ptr unsafe.Pointer, s String.Readable, dict gd.PackedStringArray) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var dict = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 1))
		defer pointers.End(dict)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, dict)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
*/
func (class) _spoof_check(impl func(ptr unsafe.Pointer, s String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
[b]Optional.[/b]
Returns the string converted to uppercase.
*/
func (class) _string_to_upper(impl func(ptr unsafe.Pointer, s String.Readable, language String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the string converted to lowercase.
*/
func (class) _string_to_lower(impl func(ptr unsafe.Pointer, s String.Readable, language String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Returns the string converted to title case.
*/
func (class) _string_to_title(impl func(ptr unsafe.Pointer, s String.Readable, language String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var s = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(s))
		var language = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(language))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
Default implementation of the BiDi algorithm override function. See [enum TextServer.StructuredTextParser] for more info.
*/
func (class) _parse_structured_text(impl func(ptr unsafe.Pointer, parser_type gdclass.TextServerStructuredTextParser, args Array.Any, text String.Readable) Array.Contains[gd.Vector3i]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var parser_type = gd.UnsafeGet[gdclass.TextServerStructuredTextParser](p_args, 0)

		var args = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalArray(args))
		var text = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(text))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parser_type, args, text)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
[b]Optional.[/b]
This method is called before text server is unregistered.
*/
func (class) _cleanup(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (self class) AsTextServerExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextServerExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTextServer() TextServer.Advanced {
	return *((*TextServer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextServer() TextServer.Instance {
	return *((*TextServer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_has_feature":
		return reflect.ValueOf(self._has_feature)
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_features":
		return reflect.ValueOf(self._get_features)
	case "_free_rid":
		return reflect.ValueOf(self._free_rid)
	case "_has":
		return reflect.ValueOf(self._has)
	case "_load_support_data":
		return reflect.ValueOf(self._load_support_data)
	case "_get_support_data_filename":
		return reflect.ValueOf(self._get_support_data_filename)
	case "_get_support_data_info":
		return reflect.ValueOf(self._get_support_data_info)
	case "_save_support_data":
		return reflect.ValueOf(self._save_support_data)
	case "_is_locale_right_to_left":
		return reflect.ValueOf(self._is_locale_right_to_left)
	case "_name_to_tag":
		return reflect.ValueOf(self._name_to_tag)
	case "_tag_to_name":
		return reflect.ValueOf(self._tag_to_name)
	case "_create_font":
		return reflect.ValueOf(self._create_font)
	case "_create_font_linked_variation":
		return reflect.ValueOf(self._create_font_linked_variation)
	case "_font_set_data":
		return reflect.ValueOf(self._font_set_data)
	case "_font_set_data_ptr":
		return reflect.ValueOf(self._font_set_data_ptr)
	case "_font_set_face_index":
		return reflect.ValueOf(self._font_set_face_index)
	case "_font_get_face_index":
		return reflect.ValueOf(self._font_get_face_index)
	case "_font_get_face_count":
		return reflect.ValueOf(self._font_get_face_count)
	case "_font_set_style":
		return reflect.ValueOf(self._font_set_style)
	case "_font_get_style":
		return reflect.ValueOf(self._font_get_style)
	case "_font_set_name":
		return reflect.ValueOf(self._font_set_name)
	case "_font_get_name":
		return reflect.ValueOf(self._font_get_name)
	case "_font_get_ot_name_strings":
		return reflect.ValueOf(self._font_get_ot_name_strings)
	case "_font_set_style_name":
		return reflect.ValueOf(self._font_set_style_name)
	case "_font_get_style_name":
		return reflect.ValueOf(self._font_get_style_name)
	case "_font_set_weight":
		return reflect.ValueOf(self._font_set_weight)
	case "_font_get_weight":
		return reflect.ValueOf(self._font_get_weight)
	case "_font_set_stretch":
		return reflect.ValueOf(self._font_set_stretch)
	case "_font_get_stretch":
		return reflect.ValueOf(self._font_get_stretch)
	case "_font_set_antialiasing":
		return reflect.ValueOf(self._font_set_antialiasing)
	case "_font_get_antialiasing":
		return reflect.ValueOf(self._font_get_antialiasing)
	case "_font_set_disable_embedded_bitmaps":
		return reflect.ValueOf(self._font_set_disable_embedded_bitmaps)
	case "_font_get_disable_embedded_bitmaps":
		return reflect.ValueOf(self._font_get_disable_embedded_bitmaps)
	case "_font_set_generate_mipmaps":
		return reflect.ValueOf(self._font_set_generate_mipmaps)
	case "_font_get_generate_mipmaps":
		return reflect.ValueOf(self._font_get_generate_mipmaps)
	case "_font_set_multichannel_signed_distance_field":
		return reflect.ValueOf(self._font_set_multichannel_signed_distance_field)
	case "_font_is_multichannel_signed_distance_field":
		return reflect.ValueOf(self._font_is_multichannel_signed_distance_field)
	case "_font_set_msdf_pixel_range":
		return reflect.ValueOf(self._font_set_msdf_pixel_range)
	case "_font_get_msdf_pixel_range":
		return reflect.ValueOf(self._font_get_msdf_pixel_range)
	case "_font_set_msdf_size":
		return reflect.ValueOf(self._font_set_msdf_size)
	case "_font_get_msdf_size":
		return reflect.ValueOf(self._font_get_msdf_size)
	case "_font_set_fixed_size":
		return reflect.ValueOf(self._font_set_fixed_size)
	case "_font_get_fixed_size":
		return reflect.ValueOf(self._font_get_fixed_size)
	case "_font_set_fixed_size_scale_mode":
		return reflect.ValueOf(self._font_set_fixed_size_scale_mode)
	case "_font_get_fixed_size_scale_mode":
		return reflect.ValueOf(self._font_get_fixed_size_scale_mode)
	case "_font_set_allow_system_fallback":
		return reflect.ValueOf(self._font_set_allow_system_fallback)
	case "_font_is_allow_system_fallback":
		return reflect.ValueOf(self._font_is_allow_system_fallback)
	case "_font_set_force_autohinter":
		return reflect.ValueOf(self._font_set_force_autohinter)
	case "_font_is_force_autohinter":
		return reflect.ValueOf(self._font_is_force_autohinter)
	case "_font_set_hinting":
		return reflect.ValueOf(self._font_set_hinting)
	case "_font_get_hinting":
		return reflect.ValueOf(self._font_get_hinting)
	case "_font_set_subpixel_positioning":
		return reflect.ValueOf(self._font_set_subpixel_positioning)
	case "_font_get_subpixel_positioning":
		return reflect.ValueOf(self._font_get_subpixel_positioning)
	case "_font_set_embolden":
		return reflect.ValueOf(self._font_set_embolden)
	case "_font_get_embolden":
		return reflect.ValueOf(self._font_get_embolden)
	case "_font_set_spacing":
		return reflect.ValueOf(self._font_set_spacing)
	case "_font_get_spacing":
		return reflect.ValueOf(self._font_get_spacing)
	case "_font_set_baseline_offset":
		return reflect.ValueOf(self._font_set_baseline_offset)
	case "_font_get_baseline_offset":
		return reflect.ValueOf(self._font_get_baseline_offset)
	case "_font_set_transform":
		return reflect.ValueOf(self._font_set_transform)
	case "_font_get_transform":
		return reflect.ValueOf(self._font_get_transform)
	case "_font_set_variation_coordinates":
		return reflect.ValueOf(self._font_set_variation_coordinates)
	case "_font_get_variation_coordinates":
		return reflect.ValueOf(self._font_get_variation_coordinates)
	case "_font_set_oversampling":
		return reflect.ValueOf(self._font_set_oversampling)
	case "_font_get_oversampling":
		return reflect.ValueOf(self._font_get_oversampling)
	case "_font_get_size_cache_list":
		return reflect.ValueOf(self._font_get_size_cache_list)
	case "_font_clear_size_cache":
		return reflect.ValueOf(self._font_clear_size_cache)
	case "_font_remove_size_cache":
		return reflect.ValueOf(self._font_remove_size_cache)
	case "_font_set_ascent":
		return reflect.ValueOf(self._font_set_ascent)
	case "_font_get_ascent":
		return reflect.ValueOf(self._font_get_ascent)
	case "_font_set_descent":
		return reflect.ValueOf(self._font_set_descent)
	case "_font_get_descent":
		return reflect.ValueOf(self._font_get_descent)
	case "_font_set_underline_position":
		return reflect.ValueOf(self._font_set_underline_position)
	case "_font_get_underline_position":
		return reflect.ValueOf(self._font_get_underline_position)
	case "_font_set_underline_thickness":
		return reflect.ValueOf(self._font_set_underline_thickness)
	case "_font_get_underline_thickness":
		return reflect.ValueOf(self._font_get_underline_thickness)
	case "_font_set_scale":
		return reflect.ValueOf(self._font_set_scale)
	case "_font_get_scale":
		return reflect.ValueOf(self._font_get_scale)
	case "_font_get_texture_count":
		return reflect.ValueOf(self._font_get_texture_count)
	case "_font_clear_textures":
		return reflect.ValueOf(self._font_clear_textures)
	case "_font_remove_texture":
		return reflect.ValueOf(self._font_remove_texture)
	case "_font_set_texture_image":
		return reflect.ValueOf(self._font_set_texture_image)
	case "_font_get_texture_image":
		return reflect.ValueOf(self._font_get_texture_image)
	case "_font_set_texture_offsets":
		return reflect.ValueOf(self._font_set_texture_offsets)
	case "_font_get_texture_offsets":
		return reflect.ValueOf(self._font_get_texture_offsets)
	case "_font_get_glyph_list":
		return reflect.ValueOf(self._font_get_glyph_list)
	case "_font_clear_glyphs":
		return reflect.ValueOf(self._font_clear_glyphs)
	case "_font_remove_glyph":
		return reflect.ValueOf(self._font_remove_glyph)
	case "_font_get_glyph_advance":
		return reflect.ValueOf(self._font_get_glyph_advance)
	case "_font_set_glyph_advance":
		return reflect.ValueOf(self._font_set_glyph_advance)
	case "_font_get_glyph_offset":
		return reflect.ValueOf(self._font_get_glyph_offset)
	case "_font_set_glyph_offset":
		return reflect.ValueOf(self._font_set_glyph_offset)
	case "_font_get_glyph_size":
		return reflect.ValueOf(self._font_get_glyph_size)
	case "_font_set_glyph_size":
		return reflect.ValueOf(self._font_set_glyph_size)
	case "_font_get_glyph_uv_rect":
		return reflect.ValueOf(self._font_get_glyph_uv_rect)
	case "_font_set_glyph_uv_rect":
		return reflect.ValueOf(self._font_set_glyph_uv_rect)
	case "_font_get_glyph_texture_idx":
		return reflect.ValueOf(self._font_get_glyph_texture_idx)
	case "_font_set_glyph_texture_idx":
		return reflect.ValueOf(self._font_set_glyph_texture_idx)
	case "_font_get_glyph_texture_rid":
		return reflect.ValueOf(self._font_get_glyph_texture_rid)
	case "_font_get_glyph_texture_size":
		return reflect.ValueOf(self._font_get_glyph_texture_size)
	case "_font_get_glyph_contours":
		return reflect.ValueOf(self._font_get_glyph_contours)
	case "_font_get_kerning_list":
		return reflect.ValueOf(self._font_get_kerning_list)
	case "_font_clear_kerning_map":
		return reflect.ValueOf(self._font_clear_kerning_map)
	case "_font_remove_kerning":
		return reflect.ValueOf(self._font_remove_kerning)
	case "_font_set_kerning":
		return reflect.ValueOf(self._font_set_kerning)
	case "_font_get_kerning":
		return reflect.ValueOf(self._font_get_kerning)
	case "_font_get_glyph_index":
		return reflect.ValueOf(self._font_get_glyph_index)
	case "_font_get_char_from_glyph_index":
		return reflect.ValueOf(self._font_get_char_from_glyph_index)
	case "_font_has_char":
		return reflect.ValueOf(self._font_has_char)
	case "_font_get_supported_chars":
		return reflect.ValueOf(self._font_get_supported_chars)
	case "_font_render_range":
		return reflect.ValueOf(self._font_render_range)
	case "_font_render_glyph":
		return reflect.ValueOf(self._font_render_glyph)
	case "_font_draw_glyph":
		return reflect.ValueOf(self._font_draw_glyph)
	case "_font_draw_glyph_outline":
		return reflect.ValueOf(self._font_draw_glyph_outline)
	case "_font_is_language_supported":
		return reflect.ValueOf(self._font_is_language_supported)
	case "_font_set_language_support_override":
		return reflect.ValueOf(self._font_set_language_support_override)
	case "_font_get_language_support_override":
		return reflect.ValueOf(self._font_get_language_support_override)
	case "_font_remove_language_support_override":
		return reflect.ValueOf(self._font_remove_language_support_override)
	case "_font_get_language_support_overrides":
		return reflect.ValueOf(self._font_get_language_support_overrides)
	case "_font_is_script_supported":
		return reflect.ValueOf(self._font_is_script_supported)
	case "_font_set_script_support_override":
		return reflect.ValueOf(self._font_set_script_support_override)
	case "_font_get_script_support_override":
		return reflect.ValueOf(self._font_get_script_support_override)
	case "_font_remove_script_support_override":
		return reflect.ValueOf(self._font_remove_script_support_override)
	case "_font_get_script_support_overrides":
		return reflect.ValueOf(self._font_get_script_support_overrides)
	case "_font_set_opentype_feature_overrides":
		return reflect.ValueOf(self._font_set_opentype_feature_overrides)
	case "_font_get_opentype_feature_overrides":
		return reflect.ValueOf(self._font_get_opentype_feature_overrides)
	case "_font_supported_feature_list":
		return reflect.ValueOf(self._font_supported_feature_list)
	case "_font_supported_variation_list":
		return reflect.ValueOf(self._font_supported_variation_list)
	case "_font_get_global_oversampling":
		return reflect.ValueOf(self._font_get_global_oversampling)
	case "_font_set_global_oversampling":
		return reflect.ValueOf(self._font_set_global_oversampling)
	case "_get_hex_code_box_size":
		return reflect.ValueOf(self._get_hex_code_box_size)
	case "_draw_hex_code_box":
		return reflect.ValueOf(self._draw_hex_code_box)
	case "_create_shaped_text":
		return reflect.ValueOf(self._create_shaped_text)
	case "_shaped_text_clear":
		return reflect.ValueOf(self._shaped_text_clear)
	case "_shaped_text_set_direction":
		return reflect.ValueOf(self._shaped_text_set_direction)
	case "_shaped_text_get_direction":
		return reflect.ValueOf(self._shaped_text_get_direction)
	case "_shaped_text_get_inferred_direction":
		return reflect.ValueOf(self._shaped_text_get_inferred_direction)
	case "_shaped_text_set_bidi_override":
		return reflect.ValueOf(self._shaped_text_set_bidi_override)
	case "_shaped_text_set_custom_punctuation":
		return reflect.ValueOf(self._shaped_text_set_custom_punctuation)
	case "_shaped_text_get_custom_punctuation":
		return reflect.ValueOf(self._shaped_text_get_custom_punctuation)
	case "_shaped_text_set_custom_ellipsis":
		return reflect.ValueOf(self._shaped_text_set_custom_ellipsis)
	case "_shaped_text_get_custom_ellipsis":
		return reflect.ValueOf(self._shaped_text_get_custom_ellipsis)
	case "_shaped_text_set_orientation":
		return reflect.ValueOf(self._shaped_text_set_orientation)
	case "_shaped_text_get_orientation":
		return reflect.ValueOf(self._shaped_text_get_orientation)
	case "_shaped_text_set_preserve_invalid":
		return reflect.ValueOf(self._shaped_text_set_preserve_invalid)
	case "_shaped_text_get_preserve_invalid":
		return reflect.ValueOf(self._shaped_text_get_preserve_invalid)
	case "_shaped_text_set_preserve_control":
		return reflect.ValueOf(self._shaped_text_set_preserve_control)
	case "_shaped_text_get_preserve_control":
		return reflect.ValueOf(self._shaped_text_get_preserve_control)
	case "_shaped_text_set_spacing":
		return reflect.ValueOf(self._shaped_text_set_spacing)
	case "_shaped_text_get_spacing":
		return reflect.ValueOf(self._shaped_text_get_spacing)
	case "_shaped_text_add_string":
		return reflect.ValueOf(self._shaped_text_add_string)
	case "_shaped_text_add_object":
		return reflect.ValueOf(self._shaped_text_add_object)
	case "_shaped_text_resize_object":
		return reflect.ValueOf(self._shaped_text_resize_object)
	case "_shaped_get_span_count":
		return reflect.ValueOf(self._shaped_get_span_count)
	case "_shaped_get_span_meta":
		return reflect.ValueOf(self._shaped_get_span_meta)
	case "_shaped_set_span_update_font":
		return reflect.ValueOf(self._shaped_set_span_update_font)
	case "_shaped_text_substr":
		return reflect.ValueOf(self._shaped_text_substr)
	case "_shaped_text_get_parent":
		return reflect.ValueOf(self._shaped_text_get_parent)
	case "_shaped_text_fit_to_width":
		return reflect.ValueOf(self._shaped_text_fit_to_width)
	case "_shaped_text_tab_align":
		return reflect.ValueOf(self._shaped_text_tab_align)
	case "_shaped_text_shape":
		return reflect.ValueOf(self._shaped_text_shape)
	case "_shaped_text_update_breaks":
		return reflect.ValueOf(self._shaped_text_update_breaks)
	case "_shaped_text_update_justification_ops":
		return reflect.ValueOf(self._shaped_text_update_justification_ops)
	case "_shaped_text_is_ready":
		return reflect.ValueOf(self._shaped_text_is_ready)
	case "_shaped_text_get_glyphs":
		return reflect.ValueOf(self._shaped_text_get_glyphs)
	case "_shaped_text_sort_logical":
		return reflect.ValueOf(self._shaped_text_sort_logical)
	case "_shaped_text_get_glyph_count":
		return reflect.ValueOf(self._shaped_text_get_glyph_count)
	case "_shaped_text_get_range":
		return reflect.ValueOf(self._shaped_text_get_range)
	case "_shaped_text_get_line_breaks_adv":
		return reflect.ValueOf(self._shaped_text_get_line_breaks_adv)
	case "_shaped_text_get_line_breaks":
		return reflect.ValueOf(self._shaped_text_get_line_breaks)
	case "_shaped_text_get_word_breaks":
		return reflect.ValueOf(self._shaped_text_get_word_breaks)
	case "_shaped_text_get_trim_pos":
		return reflect.ValueOf(self._shaped_text_get_trim_pos)
	case "_shaped_text_get_ellipsis_pos":
		return reflect.ValueOf(self._shaped_text_get_ellipsis_pos)
	case "_shaped_text_get_ellipsis_glyph_count":
		return reflect.ValueOf(self._shaped_text_get_ellipsis_glyph_count)
	case "_shaped_text_get_ellipsis_glyphs":
		return reflect.ValueOf(self._shaped_text_get_ellipsis_glyphs)
	case "_shaped_text_overrun_trim_to_width":
		return reflect.ValueOf(self._shaped_text_overrun_trim_to_width)
	case "_shaped_text_get_objects":
		return reflect.ValueOf(self._shaped_text_get_objects)
	case "_shaped_text_get_object_rect":
		return reflect.ValueOf(self._shaped_text_get_object_rect)
	case "_shaped_text_get_object_range":
		return reflect.ValueOf(self._shaped_text_get_object_range)
	case "_shaped_text_get_object_glyph":
		return reflect.ValueOf(self._shaped_text_get_object_glyph)
	case "_shaped_text_get_size":
		return reflect.ValueOf(self._shaped_text_get_size)
	case "_shaped_text_get_ascent":
		return reflect.ValueOf(self._shaped_text_get_ascent)
	case "_shaped_text_get_descent":
		return reflect.ValueOf(self._shaped_text_get_descent)
	case "_shaped_text_get_width":
		return reflect.ValueOf(self._shaped_text_get_width)
	case "_shaped_text_get_underline_position":
		return reflect.ValueOf(self._shaped_text_get_underline_position)
	case "_shaped_text_get_underline_thickness":
		return reflect.ValueOf(self._shaped_text_get_underline_thickness)
	case "_shaped_text_get_dominant_direction_in_range":
		return reflect.ValueOf(self._shaped_text_get_dominant_direction_in_range)
	case "_shaped_text_get_carets":
		return reflect.ValueOf(self._shaped_text_get_carets)
	case "_shaped_text_get_selection":
		return reflect.ValueOf(self._shaped_text_get_selection)
	case "_shaped_text_hit_test_grapheme":
		return reflect.ValueOf(self._shaped_text_hit_test_grapheme)
	case "_shaped_text_hit_test_position":
		return reflect.ValueOf(self._shaped_text_hit_test_position)
	case "_shaped_text_draw":
		return reflect.ValueOf(self._shaped_text_draw)
	case "_shaped_text_draw_outline":
		return reflect.ValueOf(self._shaped_text_draw_outline)
	case "_shaped_text_get_grapheme_bounds":
		return reflect.ValueOf(self._shaped_text_get_grapheme_bounds)
	case "_shaped_text_next_grapheme_pos":
		return reflect.ValueOf(self._shaped_text_next_grapheme_pos)
	case "_shaped_text_prev_grapheme_pos":
		return reflect.ValueOf(self._shaped_text_prev_grapheme_pos)
	case "_shaped_text_get_character_breaks":
		return reflect.ValueOf(self._shaped_text_get_character_breaks)
	case "_shaped_text_next_character_pos":
		return reflect.ValueOf(self._shaped_text_next_character_pos)
	case "_shaped_text_prev_character_pos":
		return reflect.ValueOf(self._shaped_text_prev_character_pos)
	case "_shaped_text_closest_character_pos":
		return reflect.ValueOf(self._shaped_text_closest_character_pos)
	case "_format_number":
		return reflect.ValueOf(self._format_number)
	case "_parse_number":
		return reflect.ValueOf(self._parse_number)
	case "_percent_sign":
		return reflect.ValueOf(self._percent_sign)
	case "_strip_diacritics":
		return reflect.ValueOf(self._strip_diacritics)
	case "_is_valid_identifier":
		return reflect.ValueOf(self._is_valid_identifier)
	case "_is_valid_letter":
		return reflect.ValueOf(self._is_valid_letter)
	case "_string_get_word_breaks":
		return reflect.ValueOf(self._string_get_word_breaks)
	case "_string_get_character_breaks":
		return reflect.ValueOf(self._string_get_character_breaks)
	case "_is_confusable":
		return reflect.ValueOf(self._is_confusable)
	case "_spoof_check":
		return reflect.ValueOf(self._spoof_check)
	case "_string_to_upper":
		return reflect.ValueOf(self._string_to_upper)
	case "_string_to_lower":
		return reflect.ValueOf(self._string_to_lower)
	case "_string_to_title":
		return reflect.ValueOf(self._string_to_title)
	case "_parse_structured_text":
		return reflect.ValueOf(self._parse_structured_text)
	case "_cleanup":
		return reflect.ValueOf(self._cleanup)
	default:
		return gd.VirtualByName(TextServer.Advanced(self.AsTextServer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_has_feature":
		return reflect.ValueOf(self._has_feature)
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_features":
		return reflect.ValueOf(self._get_features)
	case "_free_rid":
		return reflect.ValueOf(self._free_rid)
	case "_has":
		return reflect.ValueOf(self._has)
	case "_load_support_data":
		return reflect.ValueOf(self._load_support_data)
	case "_get_support_data_filename":
		return reflect.ValueOf(self._get_support_data_filename)
	case "_get_support_data_info":
		return reflect.ValueOf(self._get_support_data_info)
	case "_save_support_data":
		return reflect.ValueOf(self._save_support_data)
	case "_is_locale_right_to_left":
		return reflect.ValueOf(self._is_locale_right_to_left)
	case "_name_to_tag":
		return reflect.ValueOf(self._name_to_tag)
	case "_tag_to_name":
		return reflect.ValueOf(self._tag_to_name)
	case "_create_font":
		return reflect.ValueOf(self._create_font)
	case "_create_font_linked_variation":
		return reflect.ValueOf(self._create_font_linked_variation)
	case "_font_set_data":
		return reflect.ValueOf(self._font_set_data)
	case "_font_set_data_ptr":
		return reflect.ValueOf(self._font_set_data_ptr)
	case "_font_set_face_index":
		return reflect.ValueOf(self._font_set_face_index)
	case "_font_get_face_index":
		return reflect.ValueOf(self._font_get_face_index)
	case "_font_get_face_count":
		return reflect.ValueOf(self._font_get_face_count)
	case "_font_set_style":
		return reflect.ValueOf(self._font_set_style)
	case "_font_get_style":
		return reflect.ValueOf(self._font_get_style)
	case "_font_set_name":
		return reflect.ValueOf(self._font_set_name)
	case "_font_get_name":
		return reflect.ValueOf(self._font_get_name)
	case "_font_get_ot_name_strings":
		return reflect.ValueOf(self._font_get_ot_name_strings)
	case "_font_set_style_name":
		return reflect.ValueOf(self._font_set_style_name)
	case "_font_get_style_name":
		return reflect.ValueOf(self._font_get_style_name)
	case "_font_set_weight":
		return reflect.ValueOf(self._font_set_weight)
	case "_font_get_weight":
		return reflect.ValueOf(self._font_get_weight)
	case "_font_set_stretch":
		return reflect.ValueOf(self._font_set_stretch)
	case "_font_get_stretch":
		return reflect.ValueOf(self._font_get_stretch)
	case "_font_set_antialiasing":
		return reflect.ValueOf(self._font_set_antialiasing)
	case "_font_get_antialiasing":
		return reflect.ValueOf(self._font_get_antialiasing)
	case "_font_set_disable_embedded_bitmaps":
		return reflect.ValueOf(self._font_set_disable_embedded_bitmaps)
	case "_font_get_disable_embedded_bitmaps":
		return reflect.ValueOf(self._font_get_disable_embedded_bitmaps)
	case "_font_set_generate_mipmaps":
		return reflect.ValueOf(self._font_set_generate_mipmaps)
	case "_font_get_generate_mipmaps":
		return reflect.ValueOf(self._font_get_generate_mipmaps)
	case "_font_set_multichannel_signed_distance_field":
		return reflect.ValueOf(self._font_set_multichannel_signed_distance_field)
	case "_font_is_multichannel_signed_distance_field":
		return reflect.ValueOf(self._font_is_multichannel_signed_distance_field)
	case "_font_set_msdf_pixel_range":
		return reflect.ValueOf(self._font_set_msdf_pixel_range)
	case "_font_get_msdf_pixel_range":
		return reflect.ValueOf(self._font_get_msdf_pixel_range)
	case "_font_set_msdf_size":
		return reflect.ValueOf(self._font_set_msdf_size)
	case "_font_get_msdf_size":
		return reflect.ValueOf(self._font_get_msdf_size)
	case "_font_set_fixed_size":
		return reflect.ValueOf(self._font_set_fixed_size)
	case "_font_get_fixed_size":
		return reflect.ValueOf(self._font_get_fixed_size)
	case "_font_set_fixed_size_scale_mode":
		return reflect.ValueOf(self._font_set_fixed_size_scale_mode)
	case "_font_get_fixed_size_scale_mode":
		return reflect.ValueOf(self._font_get_fixed_size_scale_mode)
	case "_font_set_allow_system_fallback":
		return reflect.ValueOf(self._font_set_allow_system_fallback)
	case "_font_is_allow_system_fallback":
		return reflect.ValueOf(self._font_is_allow_system_fallback)
	case "_font_set_force_autohinter":
		return reflect.ValueOf(self._font_set_force_autohinter)
	case "_font_is_force_autohinter":
		return reflect.ValueOf(self._font_is_force_autohinter)
	case "_font_set_hinting":
		return reflect.ValueOf(self._font_set_hinting)
	case "_font_get_hinting":
		return reflect.ValueOf(self._font_get_hinting)
	case "_font_set_subpixel_positioning":
		return reflect.ValueOf(self._font_set_subpixel_positioning)
	case "_font_get_subpixel_positioning":
		return reflect.ValueOf(self._font_get_subpixel_positioning)
	case "_font_set_embolden":
		return reflect.ValueOf(self._font_set_embolden)
	case "_font_get_embolden":
		return reflect.ValueOf(self._font_get_embolden)
	case "_font_set_spacing":
		return reflect.ValueOf(self._font_set_spacing)
	case "_font_get_spacing":
		return reflect.ValueOf(self._font_get_spacing)
	case "_font_set_baseline_offset":
		return reflect.ValueOf(self._font_set_baseline_offset)
	case "_font_get_baseline_offset":
		return reflect.ValueOf(self._font_get_baseline_offset)
	case "_font_set_transform":
		return reflect.ValueOf(self._font_set_transform)
	case "_font_get_transform":
		return reflect.ValueOf(self._font_get_transform)
	case "_font_set_variation_coordinates":
		return reflect.ValueOf(self._font_set_variation_coordinates)
	case "_font_get_variation_coordinates":
		return reflect.ValueOf(self._font_get_variation_coordinates)
	case "_font_set_oversampling":
		return reflect.ValueOf(self._font_set_oversampling)
	case "_font_get_oversampling":
		return reflect.ValueOf(self._font_get_oversampling)
	case "_font_get_size_cache_list":
		return reflect.ValueOf(self._font_get_size_cache_list)
	case "_font_clear_size_cache":
		return reflect.ValueOf(self._font_clear_size_cache)
	case "_font_remove_size_cache":
		return reflect.ValueOf(self._font_remove_size_cache)
	case "_font_set_ascent":
		return reflect.ValueOf(self._font_set_ascent)
	case "_font_get_ascent":
		return reflect.ValueOf(self._font_get_ascent)
	case "_font_set_descent":
		return reflect.ValueOf(self._font_set_descent)
	case "_font_get_descent":
		return reflect.ValueOf(self._font_get_descent)
	case "_font_set_underline_position":
		return reflect.ValueOf(self._font_set_underline_position)
	case "_font_get_underline_position":
		return reflect.ValueOf(self._font_get_underline_position)
	case "_font_set_underline_thickness":
		return reflect.ValueOf(self._font_set_underline_thickness)
	case "_font_get_underline_thickness":
		return reflect.ValueOf(self._font_get_underline_thickness)
	case "_font_set_scale":
		return reflect.ValueOf(self._font_set_scale)
	case "_font_get_scale":
		return reflect.ValueOf(self._font_get_scale)
	case "_font_get_texture_count":
		return reflect.ValueOf(self._font_get_texture_count)
	case "_font_clear_textures":
		return reflect.ValueOf(self._font_clear_textures)
	case "_font_remove_texture":
		return reflect.ValueOf(self._font_remove_texture)
	case "_font_set_texture_image":
		return reflect.ValueOf(self._font_set_texture_image)
	case "_font_get_texture_image":
		return reflect.ValueOf(self._font_get_texture_image)
	case "_font_set_texture_offsets":
		return reflect.ValueOf(self._font_set_texture_offsets)
	case "_font_get_texture_offsets":
		return reflect.ValueOf(self._font_get_texture_offsets)
	case "_font_get_glyph_list":
		return reflect.ValueOf(self._font_get_glyph_list)
	case "_font_clear_glyphs":
		return reflect.ValueOf(self._font_clear_glyphs)
	case "_font_remove_glyph":
		return reflect.ValueOf(self._font_remove_glyph)
	case "_font_get_glyph_advance":
		return reflect.ValueOf(self._font_get_glyph_advance)
	case "_font_set_glyph_advance":
		return reflect.ValueOf(self._font_set_glyph_advance)
	case "_font_get_glyph_offset":
		return reflect.ValueOf(self._font_get_glyph_offset)
	case "_font_set_glyph_offset":
		return reflect.ValueOf(self._font_set_glyph_offset)
	case "_font_get_glyph_size":
		return reflect.ValueOf(self._font_get_glyph_size)
	case "_font_set_glyph_size":
		return reflect.ValueOf(self._font_set_glyph_size)
	case "_font_get_glyph_uv_rect":
		return reflect.ValueOf(self._font_get_glyph_uv_rect)
	case "_font_set_glyph_uv_rect":
		return reflect.ValueOf(self._font_set_glyph_uv_rect)
	case "_font_get_glyph_texture_idx":
		return reflect.ValueOf(self._font_get_glyph_texture_idx)
	case "_font_set_glyph_texture_idx":
		return reflect.ValueOf(self._font_set_glyph_texture_idx)
	case "_font_get_glyph_texture_rid":
		return reflect.ValueOf(self._font_get_glyph_texture_rid)
	case "_font_get_glyph_texture_size":
		return reflect.ValueOf(self._font_get_glyph_texture_size)
	case "_font_get_glyph_contours":
		return reflect.ValueOf(self._font_get_glyph_contours)
	case "_font_get_kerning_list":
		return reflect.ValueOf(self._font_get_kerning_list)
	case "_font_clear_kerning_map":
		return reflect.ValueOf(self._font_clear_kerning_map)
	case "_font_remove_kerning":
		return reflect.ValueOf(self._font_remove_kerning)
	case "_font_set_kerning":
		return reflect.ValueOf(self._font_set_kerning)
	case "_font_get_kerning":
		return reflect.ValueOf(self._font_get_kerning)
	case "_font_get_glyph_index":
		return reflect.ValueOf(self._font_get_glyph_index)
	case "_font_get_char_from_glyph_index":
		return reflect.ValueOf(self._font_get_char_from_glyph_index)
	case "_font_has_char":
		return reflect.ValueOf(self._font_has_char)
	case "_font_get_supported_chars":
		return reflect.ValueOf(self._font_get_supported_chars)
	case "_font_render_range":
		return reflect.ValueOf(self._font_render_range)
	case "_font_render_glyph":
		return reflect.ValueOf(self._font_render_glyph)
	case "_font_draw_glyph":
		return reflect.ValueOf(self._font_draw_glyph)
	case "_font_draw_glyph_outline":
		return reflect.ValueOf(self._font_draw_glyph_outline)
	case "_font_is_language_supported":
		return reflect.ValueOf(self._font_is_language_supported)
	case "_font_set_language_support_override":
		return reflect.ValueOf(self._font_set_language_support_override)
	case "_font_get_language_support_override":
		return reflect.ValueOf(self._font_get_language_support_override)
	case "_font_remove_language_support_override":
		return reflect.ValueOf(self._font_remove_language_support_override)
	case "_font_get_language_support_overrides":
		return reflect.ValueOf(self._font_get_language_support_overrides)
	case "_font_is_script_supported":
		return reflect.ValueOf(self._font_is_script_supported)
	case "_font_set_script_support_override":
		return reflect.ValueOf(self._font_set_script_support_override)
	case "_font_get_script_support_override":
		return reflect.ValueOf(self._font_get_script_support_override)
	case "_font_remove_script_support_override":
		return reflect.ValueOf(self._font_remove_script_support_override)
	case "_font_get_script_support_overrides":
		return reflect.ValueOf(self._font_get_script_support_overrides)
	case "_font_set_opentype_feature_overrides":
		return reflect.ValueOf(self._font_set_opentype_feature_overrides)
	case "_font_get_opentype_feature_overrides":
		return reflect.ValueOf(self._font_get_opentype_feature_overrides)
	case "_font_supported_feature_list":
		return reflect.ValueOf(self._font_supported_feature_list)
	case "_font_supported_variation_list":
		return reflect.ValueOf(self._font_supported_variation_list)
	case "_font_get_global_oversampling":
		return reflect.ValueOf(self._font_get_global_oversampling)
	case "_font_set_global_oversampling":
		return reflect.ValueOf(self._font_set_global_oversampling)
	case "_get_hex_code_box_size":
		return reflect.ValueOf(self._get_hex_code_box_size)
	case "_draw_hex_code_box":
		return reflect.ValueOf(self._draw_hex_code_box)
	case "_create_shaped_text":
		return reflect.ValueOf(self._create_shaped_text)
	case "_shaped_text_clear":
		return reflect.ValueOf(self._shaped_text_clear)
	case "_shaped_text_set_direction":
		return reflect.ValueOf(self._shaped_text_set_direction)
	case "_shaped_text_get_direction":
		return reflect.ValueOf(self._shaped_text_get_direction)
	case "_shaped_text_get_inferred_direction":
		return reflect.ValueOf(self._shaped_text_get_inferred_direction)
	case "_shaped_text_set_bidi_override":
		return reflect.ValueOf(self._shaped_text_set_bidi_override)
	case "_shaped_text_set_custom_punctuation":
		return reflect.ValueOf(self._shaped_text_set_custom_punctuation)
	case "_shaped_text_get_custom_punctuation":
		return reflect.ValueOf(self._shaped_text_get_custom_punctuation)
	case "_shaped_text_set_custom_ellipsis":
		return reflect.ValueOf(self._shaped_text_set_custom_ellipsis)
	case "_shaped_text_get_custom_ellipsis":
		return reflect.ValueOf(self._shaped_text_get_custom_ellipsis)
	case "_shaped_text_set_orientation":
		return reflect.ValueOf(self._shaped_text_set_orientation)
	case "_shaped_text_get_orientation":
		return reflect.ValueOf(self._shaped_text_get_orientation)
	case "_shaped_text_set_preserve_invalid":
		return reflect.ValueOf(self._shaped_text_set_preserve_invalid)
	case "_shaped_text_get_preserve_invalid":
		return reflect.ValueOf(self._shaped_text_get_preserve_invalid)
	case "_shaped_text_set_preserve_control":
		return reflect.ValueOf(self._shaped_text_set_preserve_control)
	case "_shaped_text_get_preserve_control":
		return reflect.ValueOf(self._shaped_text_get_preserve_control)
	case "_shaped_text_set_spacing":
		return reflect.ValueOf(self._shaped_text_set_spacing)
	case "_shaped_text_get_spacing":
		return reflect.ValueOf(self._shaped_text_get_spacing)
	case "_shaped_text_add_string":
		return reflect.ValueOf(self._shaped_text_add_string)
	case "_shaped_text_add_object":
		return reflect.ValueOf(self._shaped_text_add_object)
	case "_shaped_text_resize_object":
		return reflect.ValueOf(self._shaped_text_resize_object)
	case "_shaped_get_span_count":
		return reflect.ValueOf(self._shaped_get_span_count)
	case "_shaped_get_span_meta":
		return reflect.ValueOf(self._shaped_get_span_meta)
	case "_shaped_set_span_update_font":
		return reflect.ValueOf(self._shaped_set_span_update_font)
	case "_shaped_text_substr":
		return reflect.ValueOf(self._shaped_text_substr)
	case "_shaped_text_get_parent":
		return reflect.ValueOf(self._shaped_text_get_parent)
	case "_shaped_text_fit_to_width":
		return reflect.ValueOf(self._shaped_text_fit_to_width)
	case "_shaped_text_tab_align":
		return reflect.ValueOf(self._shaped_text_tab_align)
	case "_shaped_text_shape":
		return reflect.ValueOf(self._shaped_text_shape)
	case "_shaped_text_update_breaks":
		return reflect.ValueOf(self._shaped_text_update_breaks)
	case "_shaped_text_update_justification_ops":
		return reflect.ValueOf(self._shaped_text_update_justification_ops)
	case "_shaped_text_is_ready":
		return reflect.ValueOf(self._shaped_text_is_ready)
	case "_shaped_text_get_glyphs":
		return reflect.ValueOf(self._shaped_text_get_glyphs)
	case "_shaped_text_sort_logical":
		return reflect.ValueOf(self._shaped_text_sort_logical)
	case "_shaped_text_get_glyph_count":
		return reflect.ValueOf(self._shaped_text_get_glyph_count)
	case "_shaped_text_get_range":
		return reflect.ValueOf(self._shaped_text_get_range)
	case "_shaped_text_get_line_breaks_adv":
		return reflect.ValueOf(self._shaped_text_get_line_breaks_adv)
	case "_shaped_text_get_line_breaks":
		return reflect.ValueOf(self._shaped_text_get_line_breaks)
	case "_shaped_text_get_word_breaks":
		return reflect.ValueOf(self._shaped_text_get_word_breaks)
	case "_shaped_text_get_trim_pos":
		return reflect.ValueOf(self._shaped_text_get_trim_pos)
	case "_shaped_text_get_ellipsis_pos":
		return reflect.ValueOf(self._shaped_text_get_ellipsis_pos)
	case "_shaped_text_get_ellipsis_glyph_count":
		return reflect.ValueOf(self._shaped_text_get_ellipsis_glyph_count)
	case "_shaped_text_get_ellipsis_glyphs":
		return reflect.ValueOf(self._shaped_text_get_ellipsis_glyphs)
	case "_shaped_text_overrun_trim_to_width":
		return reflect.ValueOf(self._shaped_text_overrun_trim_to_width)
	case "_shaped_text_get_objects":
		return reflect.ValueOf(self._shaped_text_get_objects)
	case "_shaped_text_get_object_rect":
		return reflect.ValueOf(self._shaped_text_get_object_rect)
	case "_shaped_text_get_object_range":
		return reflect.ValueOf(self._shaped_text_get_object_range)
	case "_shaped_text_get_object_glyph":
		return reflect.ValueOf(self._shaped_text_get_object_glyph)
	case "_shaped_text_get_size":
		return reflect.ValueOf(self._shaped_text_get_size)
	case "_shaped_text_get_ascent":
		return reflect.ValueOf(self._shaped_text_get_ascent)
	case "_shaped_text_get_descent":
		return reflect.ValueOf(self._shaped_text_get_descent)
	case "_shaped_text_get_width":
		return reflect.ValueOf(self._shaped_text_get_width)
	case "_shaped_text_get_underline_position":
		return reflect.ValueOf(self._shaped_text_get_underline_position)
	case "_shaped_text_get_underline_thickness":
		return reflect.ValueOf(self._shaped_text_get_underline_thickness)
	case "_shaped_text_get_dominant_direction_in_range":
		return reflect.ValueOf(self._shaped_text_get_dominant_direction_in_range)
	case "_shaped_text_get_carets":
		return reflect.ValueOf(self._shaped_text_get_carets)
	case "_shaped_text_get_selection":
		return reflect.ValueOf(self._shaped_text_get_selection)
	case "_shaped_text_hit_test_grapheme":
		return reflect.ValueOf(self._shaped_text_hit_test_grapheme)
	case "_shaped_text_hit_test_position":
		return reflect.ValueOf(self._shaped_text_hit_test_position)
	case "_shaped_text_draw":
		return reflect.ValueOf(self._shaped_text_draw)
	case "_shaped_text_draw_outline":
		return reflect.ValueOf(self._shaped_text_draw_outline)
	case "_shaped_text_get_grapheme_bounds":
		return reflect.ValueOf(self._shaped_text_get_grapheme_bounds)
	case "_shaped_text_next_grapheme_pos":
		return reflect.ValueOf(self._shaped_text_next_grapheme_pos)
	case "_shaped_text_prev_grapheme_pos":
		return reflect.ValueOf(self._shaped_text_prev_grapheme_pos)
	case "_shaped_text_get_character_breaks":
		return reflect.ValueOf(self._shaped_text_get_character_breaks)
	case "_shaped_text_next_character_pos":
		return reflect.ValueOf(self._shaped_text_next_character_pos)
	case "_shaped_text_prev_character_pos":
		return reflect.ValueOf(self._shaped_text_prev_character_pos)
	case "_shaped_text_closest_character_pos":
		return reflect.ValueOf(self._shaped_text_closest_character_pos)
	case "_format_number":
		return reflect.ValueOf(self._format_number)
	case "_parse_number":
		return reflect.ValueOf(self._parse_number)
	case "_percent_sign":
		return reflect.ValueOf(self._percent_sign)
	case "_strip_diacritics":
		return reflect.ValueOf(self._strip_diacritics)
	case "_is_valid_identifier":
		return reflect.ValueOf(self._is_valid_identifier)
	case "_is_valid_letter":
		return reflect.ValueOf(self._is_valid_letter)
	case "_string_get_word_breaks":
		return reflect.ValueOf(self._string_get_word_breaks)
	case "_string_get_character_breaks":
		return reflect.ValueOf(self._string_get_character_breaks)
	case "_is_confusable":
		return reflect.ValueOf(self._is_confusable)
	case "_spoof_check":
		return reflect.ValueOf(self._spoof_check)
	case "_string_to_upper":
		return reflect.ValueOf(self._string_to_upper)
	case "_string_to_lower":
		return reflect.ValueOf(self._string_to_lower)
	case "_string_to_title":
		return reflect.ValueOf(self._string_to_title)
	case "_parse_structured_text":
		return reflect.ValueOf(self._parse_structured_text)
	case "_cleanup":
		return reflect.ValueOf(self._cleanup)
	default:
		return gd.VirtualByName(TextServer.Instance(self.AsTextServer()), name)
	}
}
func init() {
	gdclass.Register("TextServerExtension", func(ptr gd.Object) any {
		return [1]gdclass.TextServerExtension{*(*gdclass.TextServerExtension)(unsafe.Pointer(&ptr))}
	})
}

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
