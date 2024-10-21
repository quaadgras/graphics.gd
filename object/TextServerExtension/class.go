package TextServerExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/TextServer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
External [TextServer] implementations should inherit from this class.
	// TextServerExtension methods that can be overridden by a [Class] that extends it.
	type TextServerExtension interface {
		//[b]Required.[/b]
		//Returns [code]true[/code] if the server supports a feature.
		HasFeature(feature classdb.TextServerFeature) bool
		//[b]Required.[/b]
		//Returns the name of the server interface.
		GetName() gd.String
		//[b]Required.[/b]
		//Returns text server features, see [enum TextServer.Feature].
		GetFeatures() gd.Int
		//[b]Required.[/b]
		//Frees an object created by this [TextServer].
		FreeRid(rid gd.RID) 
		//[b]Required.[/b]
		//Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
		Has(rid gd.RID) bool
		//[b]Optional.[/b]
		//Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
		LoadSupportData(filename gd.String) bool
		//[b]Optional.[/b]
		//Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
		GetSupportDataFilename() gd.String
		//[b]Optional.[/b]
		//Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
		GetSupportDataInfo() gd.String
		//[b]Optional.[/b]
		//Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
		SaveSupportData(filename gd.String) bool
		//[b]Required.[/b]
		//Returns [code]true[/code] if locale is right-to-left.
		IsLocaleRightToLeft(locale gd.String) bool
		//[b]Optional.[/b]
		//Converts readable feature, variation, script, or language name to OpenType tag.
		NameToTag(name gd.String) gd.Int
		//[b]Optional.[/b]
		//Converts OpenType tag to readable feature, variation, script, or language name.
		TagToName(tag gd.Int) gd.String
		//[b]Required.[/b]
		//Creates a new, empty font cache entry resource.
		CreateFont() gd.RID
		//Optional, implement if font supports extra spacing or baseline offset.
		//Creates a new variation existing font which is reusing the same glyph cache and font data.
		CreateFontLinkedVariation(font_rid gd.RID) gd.RID
		//[b]Optional.[/b]
		//Sets font source data, e.g contents of the dynamic font source file.
		FontSetData(font_rid gd.RID, data gd.PackedByteArray) 
		//[b]Optional.[/b]
		//Sets pointer to the font source data, e.g contents of the dynamic font source file.
		FontSetDataPtr(font_rid gd.RID, data_ptr unsafe.Pointer, data_size gd.Int) 
		//[b]Optional.[/b]
		//Sets an active face index in the TrueType / OpenType collection.
		FontSetFaceIndex(font_rid gd.RID, face_index gd.Int) 
		//[b]Optional.[/b]
		//Returns an active face index in the TrueType / OpenType collection.
		FontGetFaceIndex(font_rid gd.RID) gd.Int
		//[b]Optional.[/b]
		//Returns number of faces in the TrueType / OpenType collection.
		FontGetFaceCount(font_rid gd.RID) gd.Int
		//[b]Optional.[/b]
		//Sets the font style flags, see [enum TextServer.FontStyle].
		FontSetStyle(font_rid gd.RID, style classdb.TextServerFontStyle) 
		//[b]Optional.[/b]
		//Returns font style flags, see [enum TextServer.FontStyle].
		FontGetStyle(font_rid gd.RID) classdb.TextServerFontStyle
		//[b]Optional.[/b]
		//Sets the font family name.
		FontSetName(font_rid gd.RID, name gd.String) 
		//[b]Optional.[/b]
		//Returns font family name.
		FontGetName(font_rid gd.RID) gd.String
		//[b]Optional.[/b]
		//Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
		FontGetOtNameStrings(font_rid gd.RID) gd.Dictionary
		//[b]Optional.[/b]
		//Sets the font style name.
		FontSetStyleName(font_rid gd.RID, name_style gd.String) 
		//[b]Optional.[/b]
		//Returns font style name.
		FontGetStyleName(font_rid gd.RID) gd.String
		//[b]Optional.[/b]
		//Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
		FontSetWeight(font_rid gd.RID, weight gd.Int) 
		//[b]Optional.[/b]
		//Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
		FontGetWeight(font_rid gd.RID) gd.Int
		//[b]Optional.[/b]
		//Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
		FontSetStretch(font_rid gd.RID, stretch gd.Int) 
		//[b]Optional.[/b]
		//Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
		FontGetStretch(font_rid gd.RID) gd.Int
		//[b]Optional.[/b]
		//Sets font anti-aliasing mode.
		FontSetAntialiasing(font_rid gd.RID, antialiasing classdb.TextServerFontAntialiasing) 
		//[b]Optional.[/b]
		//Returns font anti-aliasing mode.
		FontGetAntialiasing(font_rid gd.RID) classdb.TextServerFontAntialiasing
		//[b]Optional.[/b]
		//If set to [code]true[/code], embedded font bitmap loading is disabled.
		FontSetDisableEmbeddedBitmaps(font_rid gd.RID, disable_embedded_bitmaps bool) 
		//[b]Optional.[/b]
		//Returns whether the font's embedded bitmap loading is disabled.
		FontGetDisableEmbeddedBitmaps(font_rid gd.RID) bool
		//[b]Optional.[/b]
		//If set to [code]true[/code] font texture mipmap generation is enabled.
		FontSetGenerateMipmaps(font_rid gd.RID, generate_mipmaps bool) 
		//[b]Optional.[/b]
		//Returns [code]true[/code] if font texture mipmap generation is enabled.
		FontGetGenerateMipmaps(font_rid gd.RID) bool
		//[b]Optional.[/b]
		//If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
		FontSetMultichannelSignedDistanceField(font_rid gd.RID, msdf bool) 
		//[b]Optional.[/b]
		//Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
		FontIsMultichannelSignedDistanceField(font_rid gd.RID) bool
		//[b]Optional.[/b]
		//Sets the width of the range around the shape between the minimum and maximum representable signed distance.
		FontSetMsdfPixelRange(font_rid gd.RID, msdf_pixel_range gd.Int) 
		//[b]Optional.[/b]
		//Returns the width of the range around the shape between the minimum and maximum representable signed distance.
		FontGetMsdfPixelRange(font_rid gd.RID) gd.Int
		//[b]Optional.[/b]
		//Sets source font size used to generate MSDF textures.
		FontSetMsdfSize(font_rid gd.RID, msdf_size gd.Int) 
		//[b]Optional.[/b]
		//Returns source font size used to generate MSDF textures.
		FontGetMsdfSize(font_rid gd.RID) gd.Int
		//[b]Required.[/b]
		//Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
		FontSetFixedSize(font_rid gd.RID, fixed_size gd.Int) 
		//[b]Required.[/b]
		//Returns bitmap font fixed size.
		FontGetFixedSize(font_rid gd.RID) gd.Int
		//[b]Required.[/b]
		//Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
		FontSetFixedSizeScaleMode(font_rid gd.RID, fixed_size_scale_mode classdb.TextServerFixedSizeScaleMode) 
		//[b]Required.[/b]
		//Returns bitmap font scaling mode.
		FontGetFixedSizeScaleMode(font_rid gd.RID) classdb.TextServerFixedSizeScaleMode
		//[b]Optional.[/b]
		//If set to [code]true[/code], system fonts can be automatically used as fallbacks.
		FontSetAllowSystemFallback(font_rid gd.RID, allow_system_fallback bool) 
		//[b]Optional.[/b]
		//Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
		FontIsAllowSystemFallback(font_rid gd.RID) bool
		//[b]Optional.[/b]
		//If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
		FontSetForceAutohinter(font_rid gd.RID, force_autohinter bool) 
		//[b]Optional.[/b]
		//Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting.
		FontIsForceAutohinter(font_rid gd.RID) bool
		//[b]Optional.[/b]
		//Sets font hinting mode. Used by dynamic fonts only.
		FontSetHinting(font_rid gd.RID, hinting classdb.TextServerHinting) 
		//[b]Optional.[/b]
		//Returns the font hinting mode. Used by dynamic fonts only.
		FontGetHinting(font_rid gd.RID) classdb.TextServerHinting
		//[b]Optional.[/b]
		//Sets font subpixel glyph positioning mode.
		FontSetSubpixelPositioning(font_rid gd.RID, subpixel_positioning classdb.TextServerSubpixelPositioning) 
		//[b]Optional.[/b]
		//Returns font subpixel glyph positioning mode.
		FontGetSubpixelPositioning(font_rid gd.RID) classdb.TextServerSubpixelPositioning
		//Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
		FontSetEmbolden(font_rid gd.RID, strength gd.Float) 
		//[b]Optional.[/b]
		//Returns font embolden strength.
		FontGetEmbolden(font_rid gd.RID) gd.Float
		//[b]Optional.[/b]
		//Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
		FontSetSpacing(font_rid gd.RID, spacing classdb.TextServerSpacingType, value gd.Int) 
		//[b]Optional.[/b]
		//Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
		FontGetSpacing(font_rid gd.RID, spacing classdb.TextServerSpacingType) gd.Int
		//[b]Optional.[/b]
		//Sets extra baseline offset (as a fraction of font height).
		FontSetBaselineOffset(font_rid gd.RID, baseline_offset gd.Float) 
		//[b]Optional.[/b]
		//Returns extra baseline offset (as a fraction of font height).
		FontGetBaselineOffset(font_rid gd.RID) gd.Float
		//[b]Optional.[/b]
		//Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
		FontSetTransform(font_rid gd.RID, transform gd.Transform2D) 
		//[b]Optional.[/b]
		//Returns 2D transform applied to the font outlines.
		FontGetTransform(font_rid gd.RID) gd.Transform2D
		//[b]Optional.[/b]
		//Sets variation coordinates for the specified font cache entry.
		FontSetVariationCoordinates(font_rid gd.RID, variation_coordinates gd.Dictionary) 
		//[b]Optional.[/b]
		//Returns variation coordinates for the specified font cache entry.
		FontGetVariationCoordinates(font_rid gd.RID) gd.Dictionary
		//[b]Optional.[/b]
		//Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
		FontSetOversampling(font_rid gd.RID, oversampling gd.Float) 
		//[b]Optional.[/b]
		//Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
		FontGetOversampling(font_rid gd.RID) gd.Float
		//[b]Required.[/b]
		//Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
		FontGetSizeCacheList(font_rid gd.RID) gd.ArrayOf[gd.Vector2i]
		//[b]Required.[/b]
		//Removes all font sizes from the cache entry.
		FontClearSizeCache(font_rid gd.RID) 
		//[b]Required.[/b]
		//Removes specified font size from the cache entry.
		FontRemoveSizeCache(font_rid gd.RID, size gd.Vector2i) 
		//[b]Required.[/b]
		//Sets the font ascent (number of pixels above the baseline).
		FontSetAscent(font_rid gd.RID, size gd.Int, ascent gd.Float) 
		//[b]Required.[/b]
		//Returns the font ascent (number of pixels above the baseline).
		FontGetAscent(font_rid gd.RID, size gd.Int) gd.Float
		//[b]Required.[/b]
		//Sets the font descent (number of pixels below the baseline).
		FontSetDescent(font_rid gd.RID, size gd.Int, descent gd.Float) 
		//[b]Required.[/b]
		//Returns the font descent (number of pixels below the baseline).
		FontGetDescent(font_rid gd.RID, size gd.Int) gd.Float
		//[b]Required.[/b]
		//Sets pixel offset of the underline below the baseline.
		FontSetUnderlinePosition(font_rid gd.RID, size gd.Int, underline_position gd.Float) 
		//[b]Required.[/b]
		//Returns pixel offset of the underline below the baseline.
		FontGetUnderlinePosition(font_rid gd.RID, size gd.Int) gd.Float
		//[b]Required.[/b]
		//Sets thickness of the underline in pixels.
		FontSetUnderlineThickness(font_rid gd.RID, size gd.Int, underline_thickness gd.Float) 
		//[b]Required.[/b]
		//Returns thickness of the underline in pixels.
		FontGetUnderlineThickness(font_rid gd.RID, size gd.Int) gd.Float
		//[b]Required.[/b]
		//Sets scaling factor of the color bitmap font.
		FontSetScale(font_rid gd.RID, size gd.Int, scale gd.Float) 
		//[b]Required.[/b]
		//Returns scaling factor of the color bitmap font.
		FontGetScale(font_rid gd.RID, size gd.Int) gd.Float
		//[b]Required.[/b]
		//Returns number of textures used by font cache entry.
		FontGetTextureCount(font_rid gd.RID, size gd.Vector2i) gd.Int
		//[b]Required.[/b]
		//Removes all textures from font cache entry.
		FontClearTextures(font_rid gd.RID, size gd.Vector2i) 
		//[b]Required.[/b]
		//Removes specified texture from the cache entry.
		FontRemoveTexture(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) 
		//[b]Required.[/b]
		//Sets font cache texture image data.
		FontSetTextureImage(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, image object.Image) 
		//[b]Required.[/b]
		//Returns font cache texture image data.
		FontGetTextureImage(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) object.Image
		//[b]Optional.[/b]
		//Sets array containing glyph packing data.
		FontSetTextureOffsets(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, offset gd.PackedInt32Array) 
		//[b]Optional.[/b]
		//Returns array containing glyph packing data.
		FontGetTextureOffsets(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) gd.PackedInt32Array
		//[b]Required.[/b]
		//Returns list of rendered glyphs in the cache entry.
		FontGetGlyphList(font_rid gd.RID, size gd.Vector2i) gd.PackedInt32Array
		//[b]Required.[/b]
		//Removes all rendered glyph information from the cache entry.
		FontClearGlyphs(font_rid gd.RID, size gd.Vector2i) 
		//[b]Required.[/b]
		//Removes specified rendered glyph information from the cache entry.
		FontRemoveGlyph(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) 
		//[b]Required.[/b]
		//Returns glyph advance (offset of the next glyph).
		FontGetGlyphAdvance(font_rid gd.RID, size gd.Int, glyph gd.Int) gd.Vector2
		//[b]Required.[/b]
		//Sets glyph advance (offset of the next glyph).
		FontSetGlyphAdvance(font_rid gd.RID, size gd.Int, glyph gd.Int, advance gd.Vector2) 
		//[b]Required.[/b]
		//Returns glyph offset from the baseline.
		FontGetGlyphOffset(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2
		//[b]Required.[/b]
		//Sets glyph offset from the baseline.
		FontSetGlyphOffset(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, offset gd.Vector2) 
		//[b]Required.[/b]
		//Returns size of the glyph.
		FontGetGlyphSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2
		//[b]Required.[/b]
		//Sets size of the glyph.
		FontSetGlyphSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, gl_size gd.Vector2) 
		//[b]Required.[/b]
		//Returns rectangle in the cache texture containing the glyph.
		FontGetGlyphUvRect(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Rect2
		//[b]Required.[/b]
		//Sets rectangle in the cache texture containing the glyph.
		FontSetGlyphUvRect(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, uv_rect gd.Rect2) 
		//[b]Required.[/b]
		//Returns index of the cache texture containing the glyph.
		FontGetGlyphTextureIdx(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Int
		//[b]Required.[/b]
		//Sets index of the cache texture containing the glyph.
		FontSetGlyphTextureIdx(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, texture_idx gd.Int) 
		//[b]Required.[/b]
		//Returns resource ID of the cache texture containing the glyph.
		FontGetGlyphTextureRid(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.RID
		//[b]Required.[/b]
		//Returns size of the cache texture containing the glyph.
		FontGetGlyphTextureSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2
		//[b]Optional.[/b]
		//Returns outline contours of the glyph.
		FontGetGlyphContours(font_rid gd.RID, size gd.Int, index gd.Int) gd.Dictionary
		//[b]Optional.[/b]
		//Returns list of the kerning overrides.
		FontGetKerningList(font_rid gd.RID, size gd.Int) gd.ArrayOf[gd.Vector2i]
		//[b]Optional.[/b]
		//Removes all kerning overrides.
		FontClearKerningMap(font_rid gd.RID, size gd.Int) 
		//[b]Optional.[/b]
		//Removes kerning override for the pair of glyphs.
		FontRemoveKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i) 
		//[b]Optional.[/b]
		//Sets kerning for the pair of glyphs.
		FontSetKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i, kerning gd.Vector2) 
		//[b]Optional.[/b]
		//Returns kerning for the pair of glyphs.
		FontGetKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i) gd.Vector2
		//[b]Required.[/b]
		//Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
		FontGetGlyphIndex(font_rid gd.RID, size gd.Int, char gd.Int, variation_selector gd.Int) gd.Int
		//[b]Required.[/b]
		//Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid.
		FontGetCharFromGlyphIndex(font_rid gd.RID, size gd.Int, glyph_index gd.Int) gd.Int
		//[b]Required.[/b]
		//Returns [code]true[/code] if a Unicode [param char] is available in the font.
		FontHasChar(font_rid gd.RID, char gd.Int) bool
		//[b]Required.[/b]
		//Returns a string containing all the characters available in the font.
		FontGetSupportedChars(font_rid gd.RID) gd.String
		//[b]Optional.[/b]
		//Renders the range of characters to the font cache texture.
		FontRenderRange(font_rid gd.RID, size gd.Vector2i, start gd.Int, end gd.Int) 
		//[b]Optional.[/b]
		//Renders specified glyph to the font cache texture.
		FontRenderGlyph(font_rid gd.RID, size gd.Vector2i, index gd.Int) 
		//[b]Required.[/b]
		//Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
		FontDrawGlyph(font_rid gd.RID, canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) 
		//[b]Required.[/b]
		//Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
		FontDrawGlyphOutline(font_rid gd.RID, canvas gd.RID, size gd.Int, outline_size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) 
		//[b]Optional.[/b]
		//Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
		FontIsLanguageSupported(font_rid gd.RID, language gd.String) bool
		//[b]Optional.[/b]
		//Adds override for [method _font_is_language_supported].
		FontSetLanguageSupportOverride(font_rid gd.RID, language gd.String, supported bool) 
		//[b]Optional.[/b]
		//Returns [code]true[/code] if support override is enabled for the [param language].
		FontGetLanguageSupportOverride(font_rid gd.RID, language gd.String) bool
		//[b]Optional.[/b]
		//Remove language support override.
		FontRemoveLanguageSupportOverride(font_rid gd.RID, language gd.String) 
		//[b]Optional.[/b]
		//Returns list of language support overrides.
		FontGetLanguageSupportOverrides(font_rid gd.RID) gd.PackedStringArray
		//[b]Optional.[/b]
		//Returns [code]true[/code], if font supports given script (ISO 15924 code).
		FontIsScriptSupported(font_rid gd.RID, script gd.String) bool
		//[b]Optional.[/b]
		//Adds override for [method _font_is_script_supported].
		FontSetScriptSupportOverride(font_rid gd.RID, script gd.String, supported bool) 
		//[b]Optional.[/b]
		//Returns [code]true[/code] if support override is enabled for the [param script].
		FontGetScriptSupportOverride(font_rid gd.RID, script gd.String) bool
		//[b]Optional.[/b]
		//Removes script support override.
		FontRemoveScriptSupportOverride(font_rid gd.RID, script gd.String) 
		//[b]Optional.[/b]
		//Returns list of script support overrides.
		FontGetScriptSupportOverrides(font_rid gd.RID) gd.PackedStringArray
		//[b]Optional.[/b]
		//Sets font OpenType feature set override.
		FontSetOpentypeFeatureOverrides(font_rid gd.RID, overrides gd.Dictionary) 
		//[b]Optional.[/b]
		//Returns font OpenType feature set override.
		FontGetOpentypeFeatureOverrides(font_rid gd.RID) gd.Dictionary
		//[b]Optional.[/b]
		//Returns the dictionary of the supported OpenType features.
		FontSupportedFeatureList(font_rid gd.RID) gd.Dictionary
		//[b]Optional.[/b]
		//Returns the dictionary of the supported OpenType variation coordinates.
		FontSupportedVariationList(font_rid gd.RID) gd.Dictionary
		//[b]Optional.[/b]
		//Returns the font oversampling factor, shared by all fonts in the TextServer.
		FontGetGlobalOversampling() gd.Float
		//[b]Optional.[/b]
		//Sets oversampling factor, shared by all font in the TextServer.
		FontSetGlobalOversampling(oversampling gd.Float) 
		//[b]Optional.[/b]
		//Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
		GetHexCodeBoxSize(size gd.Int, index gd.Int) gd.Vector2
		//[b]Optional.[/b]
		//Draws box displaying character hexadecimal code.
		DrawHexCodeBox(canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) 
		//[b]Required.[/b]
		//Creates a new buffer for complex text layout, with the given [param direction] and [param orientation].
		CreateShapedText(direction classdb.TextServerDirection, orientation classdb.TextServerOrientation) gd.RID
		//[b]Required.[/b]
		//Clears text buffer (removes text and inline objects).
		ShapedTextClear(shaped gd.RID) 
		//[b]Optional.[/b]
		//Sets desired text direction. If set to [constant TextServer.DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
		ShapedTextSetDirection(shaped gd.RID, direction classdb.TextServerDirection) 
		//[b]Optional.[/b]
		//Returns direction of the text.
		ShapedTextGetDirection(shaped gd.RID) classdb.TextServerDirection
		//[b]Optional.[/b]
		//Returns direction of the text, inferred by the BiDi algorithm.
		ShapedTextGetInferredDirection(shaped gd.RID) classdb.TextServerDirection
		//[b]Optional.[/b]
		//Overrides BiDi for the structured text.
		ShapedTextSetBidiOverride(shaped gd.RID, override gd.Array) 
		//[b]Optional.[/b]
		//Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
		ShapedTextSetCustomPunctuation(shaped gd.RID, punct gd.String) 
		//[b]Optional.[/b]
		//Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
		ShapedTextGetCustomPunctuation(shaped gd.RID) gd.String
		//[b]Optional.[/b]
		//Sets ellipsis character used for text clipping.
		ShapedTextSetCustomEllipsis(shaped gd.RID, char gd.Int) 
		//[b]Optional.[/b]
		//Returns ellipsis character used for text clipping.
		ShapedTextGetCustomEllipsis(shaped gd.RID) gd.Int
		//[b]Optional.[/b]
		//Sets desired text orientation.
		ShapedTextSetOrientation(shaped gd.RID, orientation classdb.TextServerOrientation) 
		//[b]Optional.[/b]
		//Returns text orientation.
		ShapedTextGetOrientation(shaped gd.RID) classdb.TextServerOrientation
		//[b]Optional.[/b]
		//If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
		ShapedTextSetPreserveInvalid(shaped gd.RID, enabled bool) 
		//[b]Optional.[/b]
		//Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
		ShapedTextGetPreserveInvalid(shaped gd.RID) bool
		//[b]Optional.[/b]
		//If set to [code]true[/code] text buffer will display control characters.
		ShapedTextSetPreserveControl(shaped gd.RID, enabled bool) 
		//[b]Optional.[/b]
		//Returns [code]true[/code] if text buffer is configured to display control characters.
		ShapedTextGetPreserveControl(shaped gd.RID) bool
		//[b]Optional.[/b]
		//Sets extra spacing added between glyphs or lines in pixels.
		ShapedTextSetSpacing(shaped gd.RID, spacing classdb.TextServerSpacingType, value gd.Int) 
		//[b]Optional.[/b]
		//Returns extra spacing added between glyphs or lines in pixels.
		ShapedTextGetSpacing(shaped gd.RID, spacing classdb.TextServerSpacingType) gd.Int
		//[b]Required.[/b]
		//Adds text span and font to draw it to the text buffer.
		ShapedTextAddString(shaped gd.RID, text gd.String, fonts gd.ArrayOf[gd.RID], size gd.Int, opentype_features gd.Dictionary, language gd.String, meta gd.Variant) bool
		//[b]Required.[/b]
		//Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
		ShapedTextAddObject(shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, length gd.Int, baseline gd.Float) bool
		//[b]Required.[/b]
		//Sets new size and alignment of embedded object.
		ShapedTextResizeObject(shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, baseline gd.Float) bool
		//[b]Required.[/b]
		//Returns number of text spans added using [method _shaped_text_add_string] or [method _shaped_text_add_object].
		ShapedGetSpanCount(shaped gd.RID) gd.Int
		//[b]Required.[/b]
		//Returns text span metadata.
		ShapedGetSpanMeta(shaped gd.RID, index gd.Int) gd.Variant
		//[b]Required.[/b]
		//Changes text span font, font size, and OpenType features, without changing the text.
		ShapedSetSpanUpdateFont(shaped gd.RID, index gd.Int, fonts gd.ArrayOf[gd.RID], size gd.Int, opentype_features gd.Dictionary) 
		//[b]Required.[/b]
		//Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
		ShapedTextSubstr(shaped gd.RID, start gd.Int, length gd.Int) gd.RID
		//[b]Required.[/b]
		//Returns the parent buffer from which the substring originates.
		ShapedTextGetParent(shaped gd.RID) gd.RID
		//[b]Optional.[/b]
		//Adjusts text width to fit to specified width, returns new text width.
		ShapedTextFitToWidth(shaped gd.RID, width gd.Float, justification_flags classdb.TextServerJustificationFlag) gd.Float
		//[b]Optional.[/b]
		//Aligns shaped text to the given tab-stops.
		ShapedTextTabAlign(shaped gd.RID, tab_stops gd.PackedFloat32Array) gd.Float
		//[b]Required.[/b]
		//Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
		ShapedTextShape(shaped gd.RID) bool
		//[b]Optional.[/b]
		//Updates break points in the shaped text. This method is called by default implementation of text breaking functions.
		ShapedTextUpdateBreaks(shaped gd.RID) bool
		//[b]Optional.[/b]
		//Updates justification points in the shaped text. This method is called by default implementation of text justification functions.
		ShapedTextUpdateJustificationOps(shaped gd.RID) bool
		//[b]Required.[/b]
		//Returns [code]true[/code] if buffer is successfully shaped.
		ShapedTextIsReady(shaped gd.RID) bool
		//[b]Required.[/b]
		//Returns an array of glyphs in the visual order.
		ShapedTextGetGlyphs(shaped gd.RID) *classdb. Glyph
		//[b]Required.[/b]
		//Returns text glyphs in the logical order.
		ShapedTextSortLogical(shaped gd.RID) *classdb. Glyph
		//[b]Required.[/b]
		//Returns number of glyphs in the buffer.
		ShapedTextGetGlyphCount(shaped gd.RID) gd.Int
		//[b]Required.[/b]
		//Returns substring buffer character range in the parent buffer.
		ShapedTextGetRange(shaped gd.RID) gd.Vector2i
		//[b]Optional.[/b]
		//Breaks text to the lines and columns. Returns character ranges for each segment.
		ShapedTextGetLineBreaksAdv(shaped gd.RID, width gd.PackedFloat32Array, start gd.Int, once bool, break_flags classdb.TextServerLineBreakFlag) gd.PackedInt32Array
		//[b]Optional.[/b]
		//Breaks text to the lines and returns character ranges for each line.
		ShapedTextGetLineBreaks(shaped gd.RID, width gd.Float, start gd.Int, break_flags classdb.TextServerLineBreakFlag) gd.PackedInt32Array
		//[b]Optional.[/b]
		//Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum TextServer.GraphemeFlag]).
		ShapedTextGetWordBreaks(shaped gd.RID, grapheme_flags classdb.TextServerGraphemeFlag, skip_grapheme_flags classdb.TextServerGraphemeFlag) gd.PackedInt32Array
		//[b]Required.[/b]
		//Returns the position of the overrun trim.
		ShapedTextGetTrimPos(shaped gd.RID) gd.Int
		//[b]Required.[/b]
		//Returns position of the ellipsis.
		ShapedTextGetEllipsisPos(shaped gd.RID) gd.Int
		//[b]Required.[/b]
		//Returns number of glyphs in the ellipsis.
		ShapedTextGetEllipsisGlyphCount(shaped gd.RID) gd.Int
		//[b]Required.[/b]
		//Returns array of the glyphs in the ellipsis.
		ShapedTextGetEllipsisGlyphs(shaped gd.RID) *classdb. Glyph
		//[b]Optional.[/b]
		//Trims text if it exceeds the given width.
		ShapedTextOverrunTrimToWidth(shaped gd.RID, width gd.Float, trim_flags classdb.TextServerTextOverrunFlag) 
		//[b]Required.[/b]
		//Returns array of inline objects.
		ShapedTextGetObjects(shaped gd.RID) gd.Array
		//[b]Required.[/b]
		//Returns bounding rectangle of the inline object.
		ShapedTextGetObjectRect(shaped gd.RID, key gd.Variant) gd.Rect2
		//[b]Required.[/b]
		//Returns the character range of the inline object.
		ShapedTextGetObjectRange(shaped gd.RID, key gd.Variant) gd.Vector2i
		//[b]Required.[/b]
		//Returns the glyph index of the inline object.
		ShapedTextGetObjectGlyph(shaped gd.RID, key gd.Variant) gd.Int
		//[b]Required.[/b]
		//Returns size of the text.
		ShapedTextGetSize(shaped gd.RID) gd.Vector2
		//[b]Required.[/b]
		//Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
		ShapedTextGetAscent(shaped gd.RID) gd.Float
		//[b]Required.[/b]
		//Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
		ShapedTextGetDescent(shaped gd.RID) gd.Float
		//[b]Required.[/b]
		//Returns width (for horizontal layout) or height (for vertical) of the text.
		ShapedTextGetWidth(shaped gd.RID) gd.Float
		//[b]Required.[/b]
		//Returns pixel offset of the underline below the baseline.
		ShapedTextGetUnderlinePosition(shaped gd.RID) gd.Float
		//[b]Required.[/b]
		//Returns thickness of the underline.
		ShapedTextGetUnderlineThickness(shaped gd.RID) gd.Float
		//[b]Optional.[/b]
		//Returns dominant direction of in the range of text.
		ShapedTextGetDominantDirectionInRange(shaped gd.RID, start gd.Int, end gd.Int) gd.Int
		//[b]Optional.[/b]
		//Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
		ShapedTextGetCarets(shaped gd.RID, position gd.Int, caret *classdb.CaretInfo) 
		//[b]Optional.[/b]
		//Returns selection rectangles for the specified character range.
		ShapedTextGetSelection(shaped gd.RID, start gd.Int, end gd.Int) gd.PackedVector2Array
		//[b]Optional.[/b]
		//Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
		ShapedTextHitTestGrapheme(shaped gd.RID, coord gd.Float) gd.Int
		//[b]Optional.[/b]
		//Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
		ShapedTextHitTestPosition(shaped gd.RID, coord gd.Float) gd.Int
		//[b]Optional.[/b]
		//Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
		ShapedTextDraw(shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, color gd.Color) 
		//[b]Optional.[/b]
		//Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
		ShapedTextDrawOutline(shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, outline_size gd.Int, color gd.Color) 
		//[b]Optional.[/b]
		//Returns composite character's bounds as offsets from the start of the line.
		ShapedTextGetGraphemeBounds(shaped gd.RID, pos gd.Int) gd.Vector2
		//[b]Optional.[/b]
		//Returns grapheme end position closest to the [param pos].
		ShapedTextNextGraphemePos(shaped gd.RID, pos gd.Int) gd.Int
		//[b]Optional.[/b]
		//Returns grapheme start position closest to the [param pos].
		ShapedTextPrevGraphemePos(shaped gd.RID, pos gd.Int) gd.Int
		//[b]Optional.[/b]
		//Returns array of the composite character boundaries.
		ShapedTextGetCharacterBreaks(shaped gd.RID) gd.PackedInt32Array
		//[b]Optional.[/b]
		//Returns composite character end position closest to the [param pos].
		ShapedTextNextCharacterPos(shaped gd.RID, pos gd.Int) gd.Int
		//[b]Optional.[/b]
		//Returns composite character start position closest to the [param pos].
		ShapedTextPrevCharacterPos(shaped gd.RID, pos gd.Int) gd.Int
		//[b]Optional.[/b]
		//Returns composite character position closest to the [param pos].
		ShapedTextClosestCharacterPos(shaped gd.RID, pos gd.Int) gd.Int
		//[b]Optional.[/b]
		//Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
		FormatNumber(number gd.String, language gd.String) gd.String
		//[b]Optional.[/b]
		//Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
		ParseNumber(number gd.String, language gd.String) gd.String
		//[b]Optional.[/b]
		//Returns percent sign used in the [param language].
		PercentSign(language gd.String) gd.String
		//[b]Optional.[/b]
		//Strips diacritics from the string.
		StripDiacritics(s gd.String) gd.String
		//[b]Optional.[/b]
		//Returns [code]true[/code] if [param string] is a valid identifier.
		IsValidIdentifier(s gd.String) bool
		IsValidLetter(unicode gd.Int) bool
		//[b]Optional.[/b]
		//Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
		StringGetWordBreaks(s gd.String, language gd.String, chars_per_line gd.Int) gd.PackedInt32Array
		//[b]Optional.[/b]
		//Returns array of the composite character boundaries.
		StringGetCharacterBreaks(s gd.String, language gd.String) gd.PackedInt32Array
		//[b]Optional.[/b]
		//Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
		IsConfusable(s gd.String, dict gd.PackedStringArray) gd.Int
		//[b]Optional.[/b]
		//Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
		SpoofCheck(s gd.String) bool
		//[b]Optional.[/b]
		//Returns the string converted to uppercase.
		StringToUpper(s gd.String, language gd.String) gd.String
		//[b]Optional.[/b]
		//Returns the string converted to lowercase.
		StringToLower(s gd.String, language gd.String) gd.String
		//[b]Optional.[/b]
		//Returns the string converted to title case.
		StringToTitle(s gd.String, language gd.String) gd.String
		//[b]Optional.[/b]
		//Default implementation of the BiDi algorithm override function. See [enum TextServer.StructuredTextParser] for more info.
		ParseStructuredText(parser_type classdb.TextServerStructuredTextParser, args gd.Array, text gd.String) gd.ArrayOf[gd.Vector3i]
		//[b]Optional.[/b]
		//This method is called before text server is unregistered.
		Cleanup() 
	}

*/
type Simple [1]classdb.TextServerExtension
func (Simple) _has_feature(impl func(ptr unsafe.Pointer, feature classdb.TextServerFeature) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var feature = gd.UnsafeGet[classdb.TextServerFeature](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, feature)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _get_features(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, rid)
		gc.End()
	}
}
func (Simple) _has(impl func(ptr unsafe.Pointer, rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _load_support_data(impl func(ptr unsafe.Pointer, filename string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var filename = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, filename.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_support_data_filename(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _get_support_data_info(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _save_support_data(impl func(ptr unsafe.Pointer, filename string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var filename = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, filename.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _is_locale_right_to_left(impl func(ptr unsafe.Pointer, locale string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var locale = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, locale.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _name_to_tag(impl func(ptr unsafe.Pointer, name string) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name.String())
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _tag_to_name(impl func(ptr unsafe.Pointer, tag int) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var tag = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(tag))
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _create_font(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _create_font_linked_variation(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_data(impl func(ptr unsafe.Pointer, font_rid gd.RID, data []byte) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var data = mmm.Let[gd.PackedByteArray](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, data.Bytes())
		gc.End()
	}
}
func (Simple) _font_set_data_ptr(impl func(ptr unsafe.Pointer, font_rid gd.RID, data_ptr unsafe.Pointer, data_size int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var data_ptr = gd.UnsafeGet[unsafe.Pointer](p_args,1)
		var data_size = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, data_ptr, int(data_size))
		gc.End()
	}
}
func (Simple) _font_set_face_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, face_index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var face_index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(face_index))
		gc.End()
	}
}
func (Simple) _font_get_face_index(impl func(ptr unsafe.Pointer, font_rid gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_get_face_count(impl func(ptr unsafe.Pointer, font_rid gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_set_style(impl func(ptr unsafe.Pointer, font_rid gd.RID, style classdb.TextServerFontStyle) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var style = gd.UnsafeGet[classdb.TextServerFontStyle](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, style)
		gc.End()
	}
}
func (Simple) _font_get_style(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerFontStyle, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_name(impl func(ptr unsafe.Pointer, font_rid gd.RID, name string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var name = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, name.String())
		gc.End()
	}
}
func (Simple) _font_get_name(impl func(ptr unsafe.Pointer, font_rid gd.RID) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _font_get_ot_name_strings(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_set_style_name(impl func(ptr unsafe.Pointer, font_rid gd.RID, name_style string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var name_style = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, name_style.String())
		gc.End()
	}
}
func (Simple) _font_get_style_name(impl func(ptr unsafe.Pointer, font_rid gd.RID) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _font_set_weight(impl func(ptr unsafe.Pointer, font_rid gd.RID, weight int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var weight = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(weight))
		gc.End()
	}
}
func (Simple) _font_get_weight(impl func(ptr unsafe.Pointer, font_rid gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_set_stretch(impl func(ptr unsafe.Pointer, font_rid gd.RID, stretch int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var stretch = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(stretch))
		gc.End()
	}
}
func (Simple) _font_get_stretch(impl func(ptr unsafe.Pointer, font_rid gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_set_antialiasing(impl func(ptr unsafe.Pointer, font_rid gd.RID, antialiasing classdb.TextServerFontAntialiasing) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var antialiasing = gd.UnsafeGet[classdb.TextServerFontAntialiasing](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, antialiasing)
		gc.End()
	}
}
func (Simple) _font_get_antialiasing(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerFontAntialiasing, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_disable_embedded_bitmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID, disable_embedded_bitmaps bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var disable_embedded_bitmaps = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, disable_embedded_bitmaps)
		gc.End()
	}
}
func (Simple) _font_get_disable_embedded_bitmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_generate_mipmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID, generate_mipmaps bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var generate_mipmaps = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, generate_mipmaps)
		gc.End()
	}
}
func (Simple) _font_get_generate_mipmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_multichannel_signed_distance_field(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var msdf = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, msdf)
		gc.End()
	}
}
func (Simple) _font_is_multichannel_signed_distance_field(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_msdf_pixel_range(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf_pixel_range int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var msdf_pixel_range = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(msdf_pixel_range))
		gc.End()
	}
}
func (Simple) _font_get_msdf_pixel_range(impl func(ptr unsafe.Pointer, font_rid gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_set_msdf_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf_size int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var msdf_size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(msdf_size))
		gc.End()
	}
}
func (Simple) _font_get_msdf_size(impl func(ptr unsafe.Pointer, font_rid gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_set_fixed_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, fixed_size int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var fixed_size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(fixed_size))
		gc.End()
	}
}
func (Simple) _font_get_fixed_size(impl func(ptr unsafe.Pointer, font_rid gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_set_fixed_size_scale_mode(impl func(ptr unsafe.Pointer, font_rid gd.RID, fixed_size_scale_mode classdb.TextServerFixedSizeScaleMode) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var fixed_size_scale_mode = gd.UnsafeGet[classdb.TextServerFixedSizeScaleMode](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, fixed_size_scale_mode)
		gc.End()
	}
}
func (Simple) _font_get_fixed_size_scale_mode(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerFixedSizeScaleMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_allow_system_fallback(impl func(ptr unsafe.Pointer, font_rid gd.RID, allow_system_fallback bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var allow_system_fallback = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, allow_system_fallback)
		gc.End()
	}
}
func (Simple) _font_is_allow_system_fallback(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_force_autohinter(impl func(ptr unsafe.Pointer, font_rid gd.RID, force_autohinter bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var force_autohinter = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, force_autohinter)
		gc.End()
	}
}
func (Simple) _font_is_force_autohinter(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_hinting(impl func(ptr unsafe.Pointer, font_rid gd.RID, hinting classdb.TextServerHinting) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var hinting = gd.UnsafeGet[classdb.TextServerHinting](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, hinting)
		gc.End()
	}
}
func (Simple) _font_get_hinting(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerHinting, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_subpixel_positioning(impl func(ptr unsafe.Pointer, font_rid gd.RID, subpixel_positioning classdb.TextServerSubpixelPositioning) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var subpixel_positioning = gd.UnsafeGet[classdb.TextServerSubpixelPositioning](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, subpixel_positioning)
		gc.End()
	}
}
func (Simple) _font_get_subpixel_positioning(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerSubpixelPositioning, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_embolden(impl func(ptr unsafe.Pointer, font_rid gd.RID, strength float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var strength = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, float64(strength))
		gc.End()
	}
}
func (Simple) _font_get_embolden(impl func(ptr unsafe.Pointer, font_rid gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_set_spacing(impl func(ptr unsafe.Pointer, font_rid gd.RID, spacing classdb.TextServerSpacingType, value int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var spacing = gd.UnsafeGet[classdb.TextServerSpacingType](p_args,1)
		var value = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, spacing, int(value))
		gc.End()
	}
}
func (Simple) _font_get_spacing(impl func(ptr unsafe.Pointer, font_rid gd.RID, spacing classdb.TextServerSpacingType) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var spacing = gd.UnsafeGet[classdb.TextServerSpacingType](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, spacing)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_set_baseline_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, baseline_offset float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var baseline_offset = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, float64(baseline_offset))
		gc.End()
	}
}
func (Simple) _font_get_baseline_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_set_transform(impl func(ptr unsafe.Pointer, font_rid gd.RID, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, transform)
		gc.End()
	}
}
func (Simple) _font_get_transform(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_variation_coordinates(impl func(ptr unsafe.Pointer, font_rid gd.RID, variation_coordinates gd.Dictionary) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var variation_coordinates = mmm.Let[gd.Dictionary](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, variation_coordinates)
		gc.End()
	}
}
func (Simple) _font_get_variation_coordinates(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_set_oversampling(impl func(ptr unsafe.Pointer, font_rid gd.RID, oversampling float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var oversampling = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, float64(oversampling))
		gc.End()
	}
}
func (Simple) _font_get_oversampling(impl func(ptr unsafe.Pointer, font_rid gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_get_size_cache_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.ArrayOf[gd.Vector2i], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _font_clear_size_cache(impl func(ptr unsafe.Pointer, font_rid gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid)
		gc.End()
	}
}
func (Simple) _font_remove_size_cache(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size)
		gc.End()
	}
}
func (Simple) _font_set_ascent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, ascent float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var ascent = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size), float64(ascent))
		gc.End()
	}
}
func (Simple) _font_get_ascent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_set_descent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, descent float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var descent = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size), float64(descent))
		gc.End()
	}
}
func (Simple) _font_get_descent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_set_underline_position(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, underline_position float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var underline_position = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size), float64(underline_position))
		gc.End()
	}
}
func (Simple) _font_get_underline_position(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_set_underline_thickness(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, underline_thickness float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var underline_thickness = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size), float64(underline_thickness))
		gc.End()
	}
}
func (Simple) _font_get_underline_thickness(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_set_scale(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, scale float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var scale = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size), float64(scale))
		gc.End()
	}
}
func (Simple) _font_get_scale(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_get_texture_count(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_clear_textures(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size)
		gc.End()
	}
}
func (Simple) _font_remove_texture(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(texture_index))
		gc.End()
	}
}
func (Simple) _font_set_texture_image(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index int, image [1]classdb.Image) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		var image [1]classdb.Image
		image[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,3)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(texture_index), image)
		gc.End()
	}
}
func (Simple) _font_get_texture_image(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index int) [1]classdb.Image, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(texture_index))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_set_texture_offsets(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index int, offset gd.PackedInt32Array) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		var offset = mmm.Let[gd.PackedInt32Array](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,3))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(texture_index), offset)
		gc.End()
	}
}
func (Simple) _font_get_texture_offsets(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index int) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(texture_index))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_get_glyph_list(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_clear_glyphs(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size)
		gc.End()
	}
}
func (Simple) _font_remove_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(glyph))
		gc.End()
	}
}
func (Simple) _font_get_glyph_advance(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, glyph int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), int(glyph))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_glyph_advance(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, glyph int, advance gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var advance = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size), int(glyph), advance)
		gc.End()
	}
}
func (Simple) _font_get_glyph_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_glyph_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int, offset gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var offset = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(glyph), offset)
		gc.End()
	}
}
func (Simple) _font_get_glyph_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_glyph_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int, gl_size gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var gl_size = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(glyph), gl_size)
		gc.End()
	}
}
func (Simple) _font_get_glyph_uv_rect(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int) gd.Rect2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_glyph_uv_rect(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int, uv_rect gd.Rect2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var uv_rect = gd.UnsafeGet[gd.Rect2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(glyph), uv_rect)
		gc.End()
	}
}
func (Simple) _font_get_glyph_texture_idx(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_set_glyph_texture_idx(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int, texture_idx int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var texture_idx = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(glyph), int(texture_idx))
		gc.End()
	}
}
func (Simple) _font_get_glyph_texture_rid(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_get_glyph_texture_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, int(glyph))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_get_glyph_contours(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, index int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), int(index))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_get_kerning_list(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int) gd.ArrayOf[gd.Vector2i], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size))
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _font_clear_kerning_map(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size))
		gc.End()
	}
}
func (Simple) _font_remove_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, glyph_pair gd.Vector2i) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size), glyph_pair)
		gc.End()
	}
}
func (Simple) _font_set_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, glyph_pair gd.Vector2i, kerning gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args,2)
		var kerning = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, int(size), glyph_pair, kerning)
		gc.End()
	}
}
func (Simple) _font_get_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, glyph_pair gd.Vector2i) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), glyph_pair)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_get_glyph_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, char int, variation_selector int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var char = gd.UnsafeGet[gd.Int](p_args,2)
		var variation_selector = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), int(char), int(variation_selector))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_get_char_from_glyph_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, size int, glyph_index int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph_index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(size), int(glyph_index))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _font_has_char(impl func(ptr unsafe.Pointer, font_rid gd.RID, char int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var char = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, int(char))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_get_supported_chars(impl func(ptr unsafe.Pointer, font_rid gd.RID) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _font_render_range(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, start int, end int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var start = gd.UnsafeGet[gd.Int](p_args,2)
		var end = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(start), int(end))
		gc.End()
	}
}
func (Simple) _font_render_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, index int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, int(index))
		gc.End()
	}
}
func (Simple) _font_draw_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, canvas gd.RID, size int, pos gd.Vector2, index int, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var canvas = gd.UnsafeGet[gd.RID](p_args,1)
		var size = gd.UnsafeGet[gd.Int](p_args,2)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,3)
		var index = gd.UnsafeGet[gd.Int](p_args,4)
		var color = gd.UnsafeGet[gd.Color](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, canvas, int(size), pos, int(index), color)
		gc.End()
	}
}
func (Simple) _font_draw_glyph_outline(impl func(ptr unsafe.Pointer, font_rid gd.RID, canvas gd.RID, size int, outline_size int, pos gd.Vector2, index int, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var canvas = gd.UnsafeGet[gd.RID](p_args,1)
		var size = gd.UnsafeGet[gd.Int](p_args,2)
		var outline_size = gd.UnsafeGet[gd.Int](p_args,3)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,4)
		var index = gd.UnsafeGet[gd.Int](p_args,5)
		var color = gd.UnsafeGet[gd.Color](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, canvas, int(size), int(outline_size), pos, int(index), color)
		gc.End()
	}
}
func (Simple) _font_is_language_supported(impl func(ptr unsafe.Pointer, font_rid gd.RID, language string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, language.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language string, supported bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var supported = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, language.String(), supported)
		gc.End()
	}
}
func (Simple) _font_get_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, language.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_remove_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, language.String())
		gc.End()
	}
}
func (Simple) _font_get_language_support_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_is_script_supported(impl func(ptr unsafe.Pointer, font_rid gd.RID, script string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var script = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, script.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_set_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script string, supported bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var script = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var supported = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, script.String(), supported)
		gc.End()
	}
}
func (Simple) _font_get_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var script = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, script.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _font_remove_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var script = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, script.String())
		gc.End()
	}
}
func (Simple) _font_get_script_support_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_set_opentype_feature_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID, overrides gd.Dictionary) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var overrides = mmm.Let[gd.Dictionary](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, overrides)
		gc.End()
	}
}
func (Simple) _font_get_opentype_feature_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_supported_feature_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_supported_variation_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _font_get_global_oversampling(impl func(ptr unsafe.Pointer) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _font_set_global_oversampling(impl func(ptr unsafe.Pointer, oversampling float64) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var oversampling = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, float64(oversampling))
		gc.End()
	}
}
func (Simple) _get_hex_code_box_size(impl func(ptr unsafe.Pointer, size int, index int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var size = gd.UnsafeGet[gd.Int](p_args,0)
		var index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(size), int(index))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _draw_hex_code_box(impl func(ptr unsafe.Pointer, canvas gd.RID, size int, pos gd.Vector2, index int, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var canvas = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,2)
		var index = gd.UnsafeGet[gd.Int](p_args,3)
		var color = gd.UnsafeGet[gd.Color](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, canvas, int(size), pos, int(index), color)
		gc.End()
	}
}
func (Simple) _create_shaped_text(impl func(ptr unsafe.Pointer, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var direction = gd.UnsafeGet[classdb.TextServerDirection](p_args,0)
		var orientation = gd.UnsafeGet[classdb.TextServerOrientation](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, direction, orientation)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_clear(impl func(ptr unsafe.Pointer, shaped gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped)
		gc.End()
	}
}
func (Simple) _shaped_text_set_direction(impl func(ptr unsafe.Pointer, shaped gd.RID, direction classdb.TextServerDirection) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var direction = gd.UnsafeGet[classdb.TextServerDirection](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, direction)
		gc.End()
	}
}
func (Simple) _shaped_text_get_direction(impl func(ptr unsafe.Pointer, shaped gd.RID) classdb.TextServerDirection, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_inferred_direction(impl func(ptr unsafe.Pointer, shaped gd.RID) classdb.TextServerDirection, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_set_bidi_override(impl func(ptr unsafe.Pointer, shaped gd.RID, override gd.Array) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var override = mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, override)
		gc.End()
	}
}
func (Simple) _shaped_text_set_custom_punctuation(impl func(ptr unsafe.Pointer, shaped gd.RID, punct string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var punct = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, punct.String())
		gc.End()
	}
}
func (Simple) _shaped_text_get_custom_punctuation(impl func(ptr unsafe.Pointer, shaped gd.RID) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _shaped_text_set_custom_ellipsis(impl func(ptr unsafe.Pointer, shaped gd.RID, char int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var char = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, int(char))
		gc.End()
	}
}
func (Simple) _shaped_text_get_custom_ellipsis(impl func(ptr unsafe.Pointer, shaped gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_set_orientation(impl func(ptr unsafe.Pointer, shaped gd.RID, orientation classdb.TextServerOrientation) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var orientation = gd.UnsafeGet[classdb.TextServerOrientation](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, orientation)
		gc.End()
	}
}
func (Simple) _shaped_text_get_orientation(impl func(ptr unsafe.Pointer, shaped gd.RID) classdb.TextServerOrientation, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_set_preserve_invalid(impl func(ptr unsafe.Pointer, shaped gd.RID, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var enabled = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, enabled)
		gc.End()
	}
}
func (Simple) _shaped_text_get_preserve_invalid(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_set_preserve_control(impl func(ptr unsafe.Pointer, shaped gd.RID, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var enabled = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, enabled)
		gc.End()
	}
}
func (Simple) _shaped_text_get_preserve_control(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_set_spacing(impl func(ptr unsafe.Pointer, shaped gd.RID, spacing classdb.TextServerSpacingType, value int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var spacing = gd.UnsafeGet[classdb.TextServerSpacingType](p_args,1)
		var value = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, spacing, int(value))
		gc.End()
	}
}
func (Simple) _shaped_text_get_spacing(impl func(ptr unsafe.Pointer, shaped gd.RID, spacing classdb.TextServerSpacingType) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var spacing = gd.UnsafeGet[classdb.TextServerSpacingType](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, spacing)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_add_string(impl func(ptr unsafe.Pointer, shaped gd.RID, text string, fonts gd.ArrayOf[gd.RID], size int, opentype_features gd.Dictionary, language string, meta gd.Variant) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var text = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var fonts = gd.TypedArray[gd.RID](mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2)))
		var size = gd.UnsafeGet[gd.Int](p_args,3)
		var opentype_features = mmm.Let[gd.Dictionary](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,4))
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,5))
		var meta = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,6))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, text.String(), fonts, int(size), opentype_features, language.String(), meta)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_add_object(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, length int, baseline float64) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		var size = gd.UnsafeGet[gd.Vector2](p_args,2)
		var inline_align = gd.UnsafeGet[gd.InlineAlignment](p_args,3)
		var length = gd.UnsafeGet[gd.Int](p_args,4)
		var baseline = gd.UnsafeGet[gd.Float](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key, size, inline_align, int(length), float64(baseline))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_resize_object(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, baseline float64) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		var size = gd.UnsafeGet[gd.Vector2](p_args,2)
		var inline_align = gd.UnsafeGet[gd.InlineAlignment](p_args,3)
		var baseline = gd.UnsafeGet[gd.Float](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key, size, inline_align, float64(baseline))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_get_span_count(impl func(ptr unsafe.Pointer, shaped gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_get_span_meta(impl func(ptr unsafe.Pointer, shaped gd.RID, index int) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(index))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _shaped_set_span_update_font(impl func(ptr unsafe.Pointer, shaped gd.RID, index int, fonts gd.ArrayOf[gd.RID], size int, opentype_features gd.Dictionary) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var index = gd.UnsafeGet[gd.Int](p_args,1)
		var fonts = gd.TypedArray[gd.RID](mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2)))
		var size = gd.UnsafeGet[gd.Int](p_args,3)
		var opentype_features = mmm.Let[gd.Dictionary](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,4))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, int(index), fonts, int(size), opentype_features)
		gc.End()
	}
}
func (Simple) _shaped_text_substr(impl func(ptr unsafe.Pointer, shaped gd.RID, start int, length int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var start = gd.UnsafeGet[gd.Int](p_args,1)
		var length = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(start), int(length))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_parent(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_fit_to_width(impl func(ptr unsafe.Pointer, shaped gd.RID, width float64, justification_flags classdb.TextServerJustificationFlag) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var width = gd.UnsafeGet[gd.Float](p_args,1)
		var justification_flags = gd.UnsafeGet[classdb.TextServerJustificationFlag](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, float64(width), justification_flags)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_tab_align(impl func(ptr unsafe.Pointer, shaped gd.RID, tab_stops gd.PackedFloat32Array) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var tab_stops = mmm.Let[gd.PackedFloat32Array](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, tab_stops)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_shape(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_update_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_update_justification_ops(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_is_ready(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_glyphs(impl func(ptr unsafe.Pointer, shaped gd.RID) *classdb. Glyph, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_sort_logical(impl func(ptr unsafe.Pointer, shaped gd.RID) *classdb. Glyph, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_glyph_count(impl func(ptr unsafe.Pointer, shaped gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_range(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Vector2i, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_line_breaks_adv(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.PackedFloat32Array, start int, once bool, break_flags classdb.TextServerLineBreakFlag) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var width = mmm.Let[gd.PackedFloat32Array](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		var start = gd.UnsafeGet[gd.Int](p_args,2)
		var once = gd.UnsafeGet[bool](p_args,3)
		var break_flags = gd.UnsafeGet[classdb.TextServerLineBreakFlag](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, width, int(start), once, break_flags)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_line_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID, width float64, start int, break_flags classdb.TextServerLineBreakFlag) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var width = gd.UnsafeGet[gd.Float](p_args,1)
		var start = gd.UnsafeGet[gd.Int](p_args,2)
		var break_flags = gd.UnsafeGet[classdb.TextServerLineBreakFlag](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, float64(width), int(start), break_flags)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_word_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID, grapheme_flags classdb.TextServerGraphemeFlag, skip_grapheme_flags classdb.TextServerGraphemeFlag) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var grapheme_flags = gd.UnsafeGet[classdb.TextServerGraphemeFlag](p_args,1)
		var skip_grapheme_flags = gd.UnsafeGet[classdb.TextServerGraphemeFlag](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, grapheme_flags, skip_grapheme_flags)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_trim_pos(impl func(ptr unsafe.Pointer, shaped gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_ellipsis_pos(impl func(ptr unsafe.Pointer, shaped gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_ellipsis_glyph_count(impl func(ptr unsafe.Pointer, shaped gd.RID) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_ellipsis_glyphs(impl func(ptr unsafe.Pointer, shaped gd.RID) *classdb. Glyph, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_overrun_trim_to_width(impl func(ptr unsafe.Pointer, shaped gd.RID, width float64, trim_flags classdb.TextServerTextOverrunFlag) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var width = gd.UnsafeGet[gd.Float](p_args,1)
		var trim_flags = gd.UnsafeGet[classdb.TextServerTextOverrunFlag](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, float64(width), trim_flags)
		gc.End()
	}
}
func (Simple) _shaped_text_get_objects(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_object_rect(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) gd.Rect2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_object_range(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) gd.Vector2i, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_object_glyph(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_size(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_ascent(impl func(ptr unsafe.Pointer, shaped gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_descent(impl func(ptr unsafe.Pointer, shaped gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_width(impl func(ptr unsafe.Pointer, shaped gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_underline_position(impl func(ptr unsafe.Pointer, shaped gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_underline_thickness(impl func(ptr unsafe.Pointer, shaped gd.RID) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, gd.Float(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_dominant_direction_in_range(impl func(ptr unsafe.Pointer, shaped gd.RID, start int, end int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var start = gd.UnsafeGet[gd.Int](p_args,1)
		var end = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(start), int(end))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_carets(impl func(ptr unsafe.Pointer, shaped gd.RID, position int, caret *classdb.CaretInfo) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var position = gd.UnsafeGet[gd.Int](p_args,1)
		var caret = gd.UnsafeGet[*classdb.CaretInfo](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, int(position), caret)
		gc.End()
	}
}
func (Simple) _shaped_text_get_selection(impl func(ptr unsafe.Pointer, shaped gd.RID, start int, end int) gd.PackedVector2Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var start = gd.UnsafeGet[gd.Int](p_args,1)
		var end = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(start), int(end))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_hit_test_grapheme(impl func(ptr unsafe.Pointer, shaped gd.RID, coord float64) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var coord = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, float64(coord))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_hit_test_position(impl func(ptr unsafe.Pointer, shaped gd.RID, coord float64) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var coord = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, float64(coord))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_draw(impl func(ptr unsafe.Pointer, shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l float64, clip_r float64, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var canvas = gd.UnsafeGet[gd.RID](p_args,1)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,2)
		var clip_l = gd.UnsafeGet[gd.Float](p_args,3)
		var clip_r = gd.UnsafeGet[gd.Float](p_args,4)
		var color = gd.UnsafeGet[gd.Color](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, canvas, pos, float64(clip_l), float64(clip_r), color)
		gc.End()
	}
}
func (Simple) _shaped_text_draw_outline(impl func(ptr unsafe.Pointer, shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l float64, clip_r float64, outline_size int, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var canvas = gd.UnsafeGet[gd.RID](p_args,1)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,2)
		var clip_l = gd.UnsafeGet[gd.Float](p_args,3)
		var clip_r = gd.UnsafeGet[gd.Float](p_args,4)
		var outline_size = gd.UnsafeGet[gd.Int](p_args,5)
		var color = gd.UnsafeGet[gd.Color](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, canvas, pos, float64(clip_l), float64(clip_r), int(outline_size), color)
		gc.End()
	}
}
func (Simple) _shaped_text_get_grapheme_bounds(impl func(ptr unsafe.Pointer, shaped gd.RID, pos int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _shaped_text_next_grapheme_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_prev_grapheme_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_get_character_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_next_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_prev_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _shaped_text_closest_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos int) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, int(pos))
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _format_number(impl func(ptr unsafe.Pointer, number string, language string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var number = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, number.String(), language.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _parse_number(impl func(ptr unsafe.Pointer, number string, language string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var number = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, number.String(), language.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _percent_sign(impl func(ptr unsafe.Pointer, language string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, language.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _strip_diacritics(impl func(ptr unsafe.Pointer, s string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _is_valid_identifier(impl func(ptr unsafe.Pointer, s string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _is_valid_letter(impl func(ptr unsafe.Pointer, unicode int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var unicode = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(unicode))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _string_get_word_breaks(impl func(ptr unsafe.Pointer, s string, language string, chars_per_line int) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var chars_per_line = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String(), int(chars_per_line))
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _string_get_character_breaks(impl func(ptr unsafe.Pointer, s string, language string) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String())
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _is_confusable(impl func(ptr unsafe.Pointer, s string, dict gd.PackedStringArray) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var dict = mmm.Let[gd.PackedStringArray](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), dict)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _spoof_check(impl func(ptr unsafe.Pointer, s string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _string_to_upper(impl func(ptr unsafe.Pointer, s string, language string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _string_to_lower(impl func(ptr unsafe.Pointer, s string, language string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _string_to_title(impl func(ptr unsafe.Pointer, s string, language string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var s = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s.String(), language.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _parse_structured_text(impl func(ptr unsafe.Pointer, parser_type classdb.TextServerStructuredTextParser, args gd.Array, text string) gd.ArrayOf[gd.Vector3i], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var parser_type = gd.UnsafeGet[classdb.TextServerStructuredTextParser](p_args,0)
		var args = mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var text = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parser_type, args, text.String())
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}
func (Simple) _cleanup(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TextServerExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
[b]Required.[/b]
Returns [code]true[/code] if the server supports a feature.
*/
func (class) _has_feature(impl func(ptr unsafe.Pointer, feature classdb.TextServerFeature) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var feature = gd.UnsafeGet[classdb.TextServerFeature](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, feature)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the name of the server interface.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns text server features, see [enum TextServer.Feature].
*/
func (class) _get_features(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Frees an object created by this [TextServer].
*/
func (class) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, rid)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
*/
func (class) _has(impl func(ptr unsafe.Pointer, rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
*/
func (class) _load_support_data(impl func(ptr unsafe.Pointer, filename gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var filename = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, filename)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
*/
func (class) _get_support_data_filename(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
*/
func (class) _get_support_data_info(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
*/
func (class) _save_support_data(impl func(ptr unsafe.Pointer, filename gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var filename = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, filename)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if locale is right-to-left.
*/
func (class) _is_locale_right_to_left(impl func(ptr unsafe.Pointer, locale gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var locale = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, locale)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Converts readable feature, variation, script, or language name to OpenType tag.
*/
func (class) _name_to_tag(impl func(ptr unsafe.Pointer, name gd.String) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Converts OpenType tag to readable feature, variation, script, or language name.
*/
func (class) _tag_to_name(impl func(ptr unsafe.Pointer, tag gd.Int) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var tag = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, tag)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Required.[/b]
Creates a new, empty font cache entry resource.
*/
func (class) _create_font(impl func(ptr unsafe.Pointer) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Optional, implement if font supports extra spacing or baseline offset.
Creates a new variation existing font which is reusing the same glyph cache and font data.
*/
func (class) _create_font_linked_variation(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets font source data, e.g contents of the dynamic font source file.
*/
func (class) _font_set_data(impl func(ptr unsafe.Pointer, font_rid gd.RID, data gd.PackedByteArray) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var data = mmm.Let[gd.PackedByteArray](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, data)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets pointer to the font source data, e.g contents of the dynamic font source file.
*/
func (class) _font_set_data_ptr(impl func(ptr unsafe.Pointer, font_rid gd.RID, data_ptr unsafe.Pointer, data_size gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var data_ptr = gd.UnsafeGet[unsafe.Pointer](p_args,1)
		var data_size = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, data_ptr, data_size)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets an active face index in the TrueType / OpenType collection.
*/
func (class) _font_set_face_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, face_index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var face_index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, face_index)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns an active face index in the TrueType / OpenType collection.
*/
func (class) _font_get_face_index(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns number of faces in the TrueType / OpenType collection.
*/
func (class) _font_get_face_count(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets the font style flags, see [enum TextServer.FontStyle].
*/
func (class) _font_set_style(impl func(ptr unsafe.Pointer, font_rid gd.RID, style classdb.TextServerFontStyle) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var style = gd.UnsafeGet[classdb.TextServerFontStyle](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, style)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font style flags, see [enum TextServer.FontStyle].
*/
func (class) _font_get_style(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerFontStyle, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets the font family name.
*/
func (class) _font_set_name(impl func(ptr unsafe.Pointer, font_rid gd.RID, name gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var name = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, name)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font family name.
*/
func (class) _font_get_name(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
func (class) _font_get_ot_name_strings(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets the font style name.
*/
func (class) _font_set_style_name(impl func(ptr unsafe.Pointer, font_rid gd.RID, name_style gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var name_style = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, name_style)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font style name.
*/
func (class) _font_get_style_name(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (class) _font_set_weight(impl func(ptr unsafe.Pointer, font_rid gd.RID, weight gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var weight = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, weight)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (class) _font_get_weight(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (class) _font_set_stretch(impl func(ptr unsafe.Pointer, font_rid gd.RID, stretch gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var stretch = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, stretch)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (class) _font_get_stretch(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets font anti-aliasing mode.
*/
func (class) _font_set_antialiasing(impl func(ptr unsafe.Pointer, font_rid gd.RID, antialiasing classdb.TextServerFontAntialiasing) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var antialiasing = gd.UnsafeGet[classdb.TextServerFontAntialiasing](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, antialiasing)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font anti-aliasing mode.
*/
func (class) _font_get_antialiasing(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerFontAntialiasing, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], embedded font bitmap loading is disabled.
*/
func (class) _font_set_disable_embedded_bitmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID, disable_embedded_bitmaps bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var disable_embedded_bitmaps = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, disable_embedded_bitmaps)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns whether the font's embedded bitmap loading is disabled.
*/
func (class) _font_get_disable_embedded_bitmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] font texture mipmap generation is enabled.
*/
func (class) _font_set_generate_mipmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID, generate_mipmaps bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var generate_mipmaps = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, generate_mipmaps)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if font texture mipmap generation is enabled.
*/
func (class) _font_get_generate_mipmaps(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
*/
func (class) _font_set_multichannel_signed_distance_field(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var msdf = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, msdf)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
*/
func (class) _font_is_multichannel_signed_distance_field(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (class) _font_set_msdf_pixel_range(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf_pixel_range gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var msdf_pixel_range = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, msdf_pixel_range)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (class) _font_get_msdf_pixel_range(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets source font size used to generate MSDF textures.
*/
func (class) _font_set_msdf_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, msdf_size gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var msdf_size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, msdf_size)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns source font size used to generate MSDF textures.
*/
func (class) _font_get_msdf_size(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
*/
func (class) _font_set_fixed_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, fixed_size gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var fixed_size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, fixed_size)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns bitmap font fixed size.
*/
func (class) _font_get_fixed_size(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
*/
func (class) _font_set_fixed_size_scale_mode(impl func(ptr unsafe.Pointer, font_rid gd.RID, fixed_size_scale_mode classdb.TextServerFixedSizeScaleMode) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var fixed_size_scale_mode = gd.UnsafeGet[classdb.TextServerFixedSizeScaleMode](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, fixed_size_scale_mode)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns bitmap font scaling mode.
*/
func (class) _font_get_fixed_size_scale_mode(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerFixedSizeScaleMode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code], system fonts can be automatically used as fallbacks.
*/
func (class) _font_set_allow_system_fallback(impl func(ptr unsafe.Pointer, font_rid gd.RID, allow_system_fallback bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var allow_system_fallback = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, allow_system_fallback)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
*/
func (class) _font_is_allow_system_fallback(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
*/
func (class) _font_set_force_autohinter(impl func(ptr unsafe.Pointer, font_rid gd.RID, force_autohinter bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var force_autohinter = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, force_autohinter)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting.
*/
func (class) _font_is_force_autohinter(impl func(ptr unsafe.Pointer, font_rid gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets font hinting mode. Used by dynamic fonts only.
*/
func (class) _font_set_hinting(impl func(ptr unsafe.Pointer, font_rid gd.RID, hinting classdb.TextServerHinting) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var hinting = gd.UnsafeGet[classdb.TextServerHinting](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, hinting)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the font hinting mode. Used by dynamic fonts only.
*/
func (class) _font_get_hinting(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerHinting, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets font subpixel glyph positioning mode.
*/
func (class) _font_set_subpixel_positioning(impl func(ptr unsafe.Pointer, font_rid gd.RID, subpixel_positioning classdb.TextServerSubpixelPositioning) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var subpixel_positioning = gd.UnsafeGet[classdb.TextServerSubpixelPositioning](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, subpixel_positioning)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font subpixel glyph positioning mode.
*/
func (class) _font_get_subpixel_positioning(impl func(ptr unsafe.Pointer, font_rid gd.RID) classdb.TextServerSubpixelPositioning, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (class) _font_set_embolden(impl func(ptr unsafe.Pointer, font_rid gd.RID, strength gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var strength = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, strength)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font embolden strength.
*/
func (class) _font_get_embolden(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
func (class) _font_set_spacing(impl func(ptr unsafe.Pointer, font_rid gd.RID, spacing classdb.TextServerSpacingType, value gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var spacing = gd.UnsafeGet[classdb.TextServerSpacingType](p_args,1)
		var value = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, spacing, value)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
func (class) _font_get_spacing(impl func(ptr unsafe.Pointer, font_rid gd.RID, spacing classdb.TextServerSpacingType) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var spacing = gd.UnsafeGet[classdb.TextServerSpacingType](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, spacing)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets extra baseline offset (as a fraction of font height).
*/
func (class) _font_set_baseline_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, baseline_offset gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var baseline_offset = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, baseline_offset)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns extra baseline offset (as a fraction of font height).
*/
func (class) _font_get_baseline_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
func (class) _font_set_transform(impl func(ptr unsafe.Pointer, font_rid gd.RID, transform gd.Transform2D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform2D](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, transform)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns 2D transform applied to the font outlines.
*/
func (class) _font_get_transform(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Transform2D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets variation coordinates for the specified font cache entry.
*/
func (class) _font_set_variation_coordinates(impl func(ptr unsafe.Pointer, font_rid gd.RID, variation_coordinates gd.Dictionary) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var variation_coordinates = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, variation_coordinates)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns variation coordinates for the specified font cache entry.
*/
func (class) _font_get_variation_coordinates(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (class) _font_set_oversampling(impl func(ptr unsafe.Pointer, font_rid gd.RID, oversampling gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var oversampling = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, oversampling)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (class) _font_get_oversampling(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
func (class) _font_get_size_cache_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.ArrayOf[gd.Vector2i], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
[b]Required.[/b]
Removes all font sizes from the cache entry.
*/
func (class) _font_clear_size_cache(impl func(ptr unsafe.Pointer, font_rid gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Removes specified font size from the cache entry.
*/
func (class) _font_remove_size_cache(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets the font ascent (number of pixels above the baseline).
*/
func (class) _font_set_ascent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, ascent gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var ascent = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, ascent)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the font ascent (number of pixels above the baseline).
*/
func (class) _font_get_ascent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets the font descent (number of pixels below the baseline).
*/
func (class) _font_set_descent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, descent gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var descent = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, descent)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the font descent (number of pixels below the baseline).
*/
func (class) _font_get_descent(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets pixel offset of the underline below the baseline.
*/
func (class) _font_set_underline_position(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, underline_position gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var underline_position = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, underline_position)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns pixel offset of the underline below the baseline.
*/
func (class) _font_get_underline_position(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets thickness of the underline in pixels.
*/
func (class) _font_set_underline_thickness(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, underline_thickness gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var underline_thickness = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, underline_thickness)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns thickness of the underline in pixels.
*/
func (class) _font_get_underline_thickness(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets scaling factor of the color bitmap font.
*/
func (class) _font_set_scale(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, scale gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var scale = gd.UnsafeGet[gd.Float](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, scale)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns scaling factor of the color bitmap font.
*/
func (class) _font_get_scale(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns number of textures used by font cache entry.
*/
func (class) _font_get_texture_count(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Removes all textures from font cache entry.
*/
func (class) _font_clear_textures(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Removes specified texture from the cache entry.
*/
func (class) _font_remove_texture(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, texture_index)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets font cache texture image data.
*/
func (class) _font_set_texture_image(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, image object.Image) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		var image object.Image
		image[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,3)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, texture_index, image)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns font cache texture image data.
*/
func (class) _font_get_texture_image(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) object.Image, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, texture_index)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets array containing glyph packing data.
*/
func (class) _font_set_texture_offsets(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, offset gd.PackedInt32Array) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		var offset = mmm.Let[gd.PackedInt32Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,3))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, texture_index, offset)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns array containing glyph packing data.
*/
func (class) _font_get_texture_offsets(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var texture_index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, texture_index)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns list of rendered glyphs in the cache entry.
*/
func (class) _font_get_glyph_list(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Required.[/b]
Removes all rendered glyph information from the cache entry.
*/
func (class) _font_clear_glyphs(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Removes specified rendered glyph information from the cache entry.
*/
func (class) _font_remove_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, glyph)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns glyph advance (offset of the next glyph).
*/
func (class) _font_get_glyph_advance(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets glyph advance (offset of the next glyph).
*/
func (class) _font_set_glyph_advance(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph gd.Int, advance gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var advance = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, glyph, advance)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns glyph offset from the baseline.
*/
func (class) _font_get_glyph_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets glyph offset from the baseline.
*/
func (class) _font_set_glyph_offset(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int, offset gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var offset = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, glyph, offset)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns size of the glyph.
*/
func (class) _font_get_glyph_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets size of the glyph.
*/
func (class) _font_set_glyph_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int, gl_size gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var gl_size = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, glyph, gl_size)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns rectangle in the cache texture containing the glyph.
*/
func (class) _font_get_glyph_uv_rect(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Rect2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets rectangle in the cache texture containing the glyph.
*/
func (class) _font_set_glyph_uv_rect(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int, uv_rect gd.Rect2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var uv_rect = gd.UnsafeGet[gd.Rect2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, glyph, uv_rect)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns index of the cache texture containing the glyph.
*/
func (class) _font_get_glyph_texture_idx(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets index of the cache texture containing the glyph.
*/
func (class) _font_set_glyph_texture_idx(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int, texture_idx gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		var texture_idx = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, glyph, texture_idx)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns resource ID of the cache texture containing the glyph.
*/
func (class) _font_get_glyph_texture_rid(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns size of the cache texture containing the glyph.
*/
func (class) _font_get_glyph_texture_size(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var glyph = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns outline contours of the glyph.
*/
func (class) _font_get_glyph_contours(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, index gd.Int) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, index)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns list of the kerning overrides.
*/
func (class) _font_get_kerning_list(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) gd.ArrayOf[gd.Vector2i], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Removes all kerning overrides.
*/
func (class) _font_clear_kerning_map(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Removes kerning override for the pair of glyphs.
*/
func (class) _font_remove_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, glyph_pair)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets kerning for the pair of glyphs.
*/
func (class) _font_set_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i, kerning gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args,2)
		var kerning = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, glyph_pair, kerning)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns kerning for the pair of glyphs.
*/
func (class) _font_get_kerning(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph_pair = gd.UnsafeGet[gd.Vector2i](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph_pair)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
func (class) _font_get_glyph_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, char gd.Int, variation_selector gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var char = gd.UnsafeGet[gd.Int](p_args,2)
		var variation_selector = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, char, variation_selector)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid.
*/
func (class) _font_get_char_from_glyph_index(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Int, glyph_index gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var glyph_index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, size, glyph_index)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
func (class) _font_has_char(impl func(ptr unsafe.Pointer, font_rid gd.RID, char gd.Int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var char = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, char)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns a string containing all the characters available in the font.
*/
func (class) _font_get_supported_chars(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Renders the range of characters to the font cache texture.
*/
func (class) _font_render_range(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, start gd.Int, end gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var start = gd.UnsafeGet[gd.Int](p_args,2)
		var end = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, start, end)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Renders specified glyph to the font cache texture.
*/
func (class) _font_render_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, size gd.Vector2i, index gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var index = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, size, index)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
*/
func (class) _font_draw_glyph(impl func(ptr unsafe.Pointer, font_rid gd.RID, canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var canvas = gd.UnsafeGet[gd.RID](p_args,1)
		var size = gd.UnsafeGet[gd.Int](p_args,2)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,3)
		var index = gd.UnsafeGet[gd.Int](p_args,4)
		var color = gd.UnsafeGet[gd.Color](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, canvas, size, pos, index, color)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
*/
func (class) _font_draw_glyph_outline(impl func(ptr unsafe.Pointer, font_rid gd.RID, canvas gd.RID, size gd.Int, outline_size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var canvas = gd.UnsafeGet[gd.RID](p_args,1)
		var size = gd.UnsafeGet[gd.Int](p_args,2)
		var outline_size = gd.UnsafeGet[gd.Int](p_args,3)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,4)
		var index = gd.UnsafeGet[gd.Int](p_args,5)
		var color = gd.UnsafeGet[gd.Color](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, canvas, size, outline_size, pos, index, color)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
func (class) _font_is_language_supported(impl func(ptr unsafe.Pointer, font_rid gd.RID, language gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, language)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Adds override for [method _font_is_language_supported].
*/
func (class) _font_set_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language gd.String, supported bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var supported = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, language, supported)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if support override is enabled for the [param language].
*/
func (class) _font_get_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, language)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Remove language support override.
*/
func (class) _font_remove_language_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, language gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, language)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns list of language support overrides.
*/
func (class) _font_get_language_support_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code], if font supports given script (ISO 15924 code).
*/
func (class) _font_is_script_supported(impl func(ptr unsafe.Pointer, font_rid gd.RID, script gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var script = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, script)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Adds override for [method _font_is_script_supported].
*/
func (class) _font_set_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script gd.String, supported bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var script = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var supported = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, script, supported)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if support override is enabled for the [param script].
*/
func (class) _font_get_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var script = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid, script)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Removes script support override.
*/
func (class) _font_remove_script_support_override(impl func(ptr unsafe.Pointer, font_rid gd.RID, script gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var script = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, script)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns list of script support overrides.
*/
func (class) _font_get_script_support_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets font OpenType feature set override.
*/
func (class) _font_set_opentype_feature_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID, overrides gd.Dictionary) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		var overrides = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, font_rid, overrides)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns font OpenType feature set override.
*/
func (class) _font_get_opentype_feature_overrides(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the dictionary of the supported OpenType features.
*/
func (class) _font_supported_feature_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the dictionary of the supported OpenType variation coordinates.
*/
func (class) _font_supported_variation_list(impl func(ptr unsafe.Pointer, font_rid gd.RID) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var font_rid = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, font_rid)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the font oversampling factor, shared by all fonts in the TextServer.
*/
func (class) _font_get_global_oversampling(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets oversampling factor, shared by all font in the TextServer.
*/
func (class) _font_set_global_oversampling(impl func(ptr unsafe.Pointer, oversampling gd.Float) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var oversampling = gd.UnsafeGet[gd.Float](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, oversampling)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
*/
func (class) _get_hex_code_box_size(impl func(ptr unsafe.Pointer, size gd.Int, index gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var size = gd.UnsafeGet[gd.Int](p_args,0)
		var index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, size, index)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Draws box displaying character hexadecimal code.
*/
func (class) _draw_hex_code_box(impl func(ptr unsafe.Pointer, canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var canvas = gd.UnsafeGet[gd.RID](p_args,0)
		var size = gd.UnsafeGet[gd.Int](p_args,1)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,2)
		var index = gd.UnsafeGet[gd.Int](p_args,3)
		var color = gd.UnsafeGet[gd.Color](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, canvas, size, pos, index, color)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Creates a new buffer for complex text layout, with the given [param direction] and [param orientation].
*/
func (class) _create_shaped_text(impl func(ptr unsafe.Pointer, direction classdb.TextServerDirection, orientation classdb.TextServerOrientation) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var direction = gd.UnsafeGet[classdb.TextServerDirection](p_args,0)
		var orientation = gd.UnsafeGet[classdb.TextServerOrientation](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, direction, orientation)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Clears text buffer (removes text and inline objects).
*/
func (class) _shaped_text_clear(impl func(ptr unsafe.Pointer, shaped gd.RID) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets desired text direction. If set to [constant TextServer.DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
*/
func (class) _shaped_text_set_direction(impl func(ptr unsafe.Pointer, shaped gd.RID, direction classdb.TextServerDirection) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var direction = gd.UnsafeGet[classdb.TextServerDirection](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, direction)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns direction of the text.
*/
func (class) _shaped_text_get_direction(impl func(ptr unsafe.Pointer, shaped gd.RID) classdb.TextServerDirection, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns direction of the text, inferred by the BiDi algorithm.
*/
func (class) _shaped_text_get_inferred_direction(impl func(ptr unsafe.Pointer, shaped gd.RID) classdb.TextServerDirection, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Overrides BiDi for the structured text.
*/
func (class) _shaped_text_set_bidi_override(impl func(ptr unsafe.Pointer, shaped gd.RID, override gd.Array) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var override = mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, override)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (class) _shaped_text_set_custom_punctuation(impl func(ptr unsafe.Pointer, shaped gd.RID, punct gd.String) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var punct = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, punct)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (class) _shaped_text_get_custom_punctuation(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets ellipsis character used for text clipping.
*/
func (class) _shaped_text_set_custom_ellipsis(impl func(ptr unsafe.Pointer, shaped gd.RID, char gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var char = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, char)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns ellipsis character used for text clipping.
*/
func (class) _shaped_text_get_custom_ellipsis(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets desired text orientation.
*/
func (class) _shaped_text_set_orientation(impl func(ptr unsafe.Pointer, shaped gd.RID, orientation classdb.TextServerOrientation) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var orientation = gd.UnsafeGet[classdb.TextServerOrientation](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, orientation)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns text orientation.
*/
func (class) _shaped_text_get_orientation(impl func(ptr unsafe.Pointer, shaped gd.RID) classdb.TextServerOrientation, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
*/
func (class) _shaped_text_set_preserve_invalid(impl func(ptr unsafe.Pointer, shaped gd.RID, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var enabled = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, enabled)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
*/
func (class) _shaped_text_get_preserve_invalid(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
If set to [code]true[/code] text buffer will display control characters.
*/
func (class) _shaped_text_set_preserve_control(impl func(ptr unsafe.Pointer, shaped gd.RID, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var enabled = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, enabled)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if text buffer is configured to display control characters.
*/
func (class) _shaped_text_get_preserve_control(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Sets extra spacing added between glyphs or lines in pixels.
*/
func (class) _shaped_text_set_spacing(impl func(ptr unsafe.Pointer, shaped gd.RID, spacing classdb.TextServerSpacingType, value gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var spacing = gd.UnsafeGet[classdb.TextServerSpacingType](p_args,1)
		var value = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, spacing, value)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns extra spacing added between glyphs or lines in pixels.
*/
func (class) _shaped_text_get_spacing(impl func(ptr unsafe.Pointer, shaped gd.RID, spacing classdb.TextServerSpacingType) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var spacing = gd.UnsafeGet[classdb.TextServerSpacingType](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, spacing)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Adds text span and font to draw it to the text buffer.
*/
func (class) _shaped_text_add_string(impl func(ptr unsafe.Pointer, shaped gd.RID, text gd.String, fonts gd.ArrayOf[gd.RID], size gd.Int, opentype_features gd.Dictionary, language gd.String, meta gd.Variant) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var text = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var fonts = gd.TypedArray[gd.RID](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2)))
		var size = gd.UnsafeGet[gd.Int](p_args,3)
		var opentype_features = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,4))
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,5))
		var meta = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,6))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, text, fonts, size, opentype_features, language, meta)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (class) _shaped_text_add_object(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, length gd.Int, baseline gd.Float) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		var size = gd.UnsafeGet[gd.Vector2](p_args,2)
		var inline_align = gd.UnsafeGet[gd.InlineAlignment](p_args,3)
		var length = gd.UnsafeGet[gd.Int](p_args,4)
		var baseline = gd.UnsafeGet[gd.Float](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key, size, inline_align, length, baseline)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Sets new size and alignment of embedded object.
*/
func (class) _shaped_text_resize_object(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, baseline gd.Float) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		var size = gd.UnsafeGet[gd.Vector2](p_args,2)
		var inline_align = gd.UnsafeGet[gd.InlineAlignment](p_args,3)
		var baseline = gd.UnsafeGet[gd.Float](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key, size, inline_align, baseline)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns number of text spans added using [method _shaped_text_add_string] or [method _shaped_text_add_object].
*/
func (class) _shaped_get_span_count(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns text span metadata.
*/
func (class) _shaped_get_span_meta(impl func(ptr unsafe.Pointer, shaped gd.RID, index gd.Int) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var index = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, index)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Required.[/b]
Changes text span font, font size, and OpenType features, without changing the text.
*/
func (class) _shaped_set_span_update_font(impl func(ptr unsafe.Pointer, shaped gd.RID, index gd.Int, fonts gd.ArrayOf[gd.RID], size gd.Int, opentype_features gd.Dictionary) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var index = gd.UnsafeGet[gd.Int](p_args,1)
		var fonts = gd.TypedArray[gd.RID](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2)))
		var size = gd.UnsafeGet[gd.Int](p_args,3)
		var opentype_features = mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,4))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, index, fonts, size, opentype_features)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
*/
func (class) _shaped_text_substr(impl func(ptr unsafe.Pointer, shaped gd.RID, start gd.Int, length gd.Int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var start = gd.UnsafeGet[gd.Int](p_args,1)
		var length = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, start, length)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the parent buffer from which the substring originates.
*/
func (class) _shaped_text_get_parent(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Adjusts text width to fit to specified width, returns new text width.
*/
func (class) _shaped_text_fit_to_width(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.Float, justification_flags classdb.TextServerJustificationFlag) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var width = gd.UnsafeGet[gd.Float](p_args,1)
		var justification_flags = gd.UnsafeGet[classdb.TextServerJustificationFlag](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, width, justification_flags)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Aligns shaped text to the given tab-stops.
*/
func (class) _shaped_text_tab_align(impl func(ptr unsafe.Pointer, shaped gd.RID, tab_stops gd.PackedFloat32Array) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var tab_stops = mmm.Let[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, tab_stops)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
*/
func (class) _shaped_text_shape(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Updates break points in the shaped text. This method is called by default implementation of text breaking functions.
*/
func (class) _shaped_text_update_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Updates justification points in the shaped text. This method is called by default implementation of text justification functions.
*/
func (class) _shaped_text_update_justification_ops(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns [code]true[/code] if buffer is successfully shaped.
*/
func (class) _shaped_text_is_ready(impl func(ptr unsafe.Pointer, shaped gd.RID) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns an array of glyphs in the visual order.
*/
func (class) _shaped_text_get_glyphs(impl func(ptr unsafe.Pointer, shaped gd.RID) *classdb. Glyph, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns text glyphs in the logical order.
*/
func (class) _shaped_text_sort_logical(impl func(ptr unsafe.Pointer, shaped gd.RID) *classdb. Glyph, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns number of glyphs in the buffer.
*/
func (class) _shaped_text_get_glyph_count(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns substring buffer character range in the parent buffer.
*/
func (class) _shaped_text_get_range(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Vector2i, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Breaks text to the lines and columns. Returns character ranges for each segment.
*/
func (class) _shaped_text_get_line_breaks_adv(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.PackedFloat32Array, start gd.Int, once bool, break_flags classdb.TextServerLineBreakFlag) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var width = mmm.Let[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		var start = gd.UnsafeGet[gd.Int](p_args,2)
		var once = gd.UnsafeGet[bool](p_args,3)
		var break_flags = gd.UnsafeGet[classdb.TextServerLineBreakFlag](p_args,4)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, width, start, once, break_flags)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Breaks text to the lines and returns character ranges for each line.
*/
func (class) _shaped_text_get_line_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.Float, start gd.Int, break_flags classdb.TextServerLineBreakFlag) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var width = gd.UnsafeGet[gd.Float](p_args,1)
		var start = gd.UnsafeGet[gd.Int](p_args,2)
		var break_flags = gd.UnsafeGet[classdb.TextServerLineBreakFlag](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, width, start, break_flags)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum TextServer.GraphemeFlag]).
*/
func (class) _shaped_text_get_word_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID, grapheme_flags classdb.TextServerGraphemeFlag, skip_grapheme_flags classdb.TextServerGraphemeFlag) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var grapheme_flags = gd.UnsafeGet[classdb.TextServerGraphemeFlag](p_args,1)
		var skip_grapheme_flags = gd.UnsafeGet[classdb.TextServerGraphemeFlag](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, grapheme_flags, skip_grapheme_flags)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the position of the overrun trim.
*/
func (class) _shaped_text_get_trim_pos(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns position of the ellipsis.
*/
func (class) _shaped_text_get_ellipsis_pos(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns number of glyphs in the ellipsis.
*/
func (class) _shaped_text_get_ellipsis_glyph_count(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns array of the glyphs in the ellipsis.
*/
func (class) _shaped_text_get_ellipsis_glyphs(impl func(ptr unsafe.Pointer, shaped gd.RID) *classdb. Glyph, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Trims text if it exceeds the given width.
*/
func (class) _shaped_text_overrun_trim_to_width(impl func(ptr unsafe.Pointer, shaped gd.RID, width gd.Float, trim_flags classdb.TextServerTextOverrunFlag) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var width = gd.UnsafeGet[gd.Float](p_args,1)
		var trim_flags = gd.UnsafeGet[classdb.TextServerTextOverrunFlag](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, width, trim_flags)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns array of inline objects.
*/
func (class) _shaped_text_get_objects(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns bounding rectangle of the inline object.
*/
func (class) _shaped_text_get_object_rect(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) gd.Rect2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the character range of the inline object.
*/
func (class) _shaped_text_get_object_range(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) gd.Vector2i, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the glyph index of the inline object.
*/
func (class) _shaped_text_get_object_glyph(impl func(ptr unsafe.Pointer, shaped gd.RID, key gd.Variant) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var key = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, key)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns size of the text.
*/
func (class) _shaped_text_get_size(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (class) _shaped_text_get_ascent(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (class) _shaped_text_get_descent(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
func (class) _shaped_text_get_width(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns pixel offset of the underline below the baseline.
*/
func (class) _shaped_text_get_underline_position(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Required.[/b]
Returns thickness of the underline.
*/
func (class) _shaped_text_get_underline_thickness(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns dominant direction of in the range of text.
*/
func (class) _shaped_text_get_dominant_direction_in_range(impl func(ptr unsafe.Pointer, shaped gd.RID, start gd.Int, end gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var start = gd.UnsafeGet[gd.Int](p_args,1)
		var end = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, start, end)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
*/
func (class) _shaped_text_get_carets(impl func(ptr unsafe.Pointer, shaped gd.RID, position gd.Int, caret *classdb.CaretInfo) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var position = gd.UnsafeGet[gd.Int](p_args,1)
		var caret = gd.UnsafeGet[*classdb.CaretInfo](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, position, caret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns selection rectangles for the specified character range.
*/
func (class) _shaped_text_get_selection(impl func(ptr unsafe.Pointer, shaped gd.RID, start gd.Int, end gd.Int) gd.PackedVector2Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var start = gd.UnsafeGet[gd.Int](p_args,1)
		var end = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, start, end)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
*/
func (class) _shaped_text_hit_test_grapheme(impl func(ptr unsafe.Pointer, shaped gd.RID, coord gd.Float) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var coord = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, coord)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
func (class) _shaped_text_hit_test_position(impl func(ptr unsafe.Pointer, shaped gd.RID, coord gd.Float) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var coord = gd.UnsafeGet[gd.Float](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, coord)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (class) _shaped_text_draw(impl func(ptr unsafe.Pointer, shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var canvas = gd.UnsafeGet[gd.RID](p_args,1)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,2)
		var clip_l = gd.UnsafeGet[gd.Float](p_args,3)
		var clip_r = gd.UnsafeGet[gd.Float](p_args,4)
		var color = gd.UnsafeGet[gd.Color](p_args,5)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, canvas, pos, clip_l, clip_r, color)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (class) _shaped_text_draw_outline(impl func(ptr unsafe.Pointer, shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, outline_size gd.Int, color gd.Color) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var canvas = gd.UnsafeGet[gd.RID](p_args,1)
		var pos = gd.UnsafeGet[gd.Vector2](p_args,2)
		var clip_l = gd.UnsafeGet[gd.Float](p_args,3)
		var clip_r = gd.UnsafeGet[gd.Float](p_args,4)
		var outline_size = gd.UnsafeGet[gd.Int](p_args,5)
		var color = gd.UnsafeGet[gd.Color](p_args,6)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, shaped, canvas, pos, clip_l, clip_r, outline_size, color)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns composite character's bounds as offsets from the start of the line.
*/
func (class) _shaped_text_get_grapheme_bounds(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns grapheme end position closest to the [param pos].
*/
func (class) _shaped_text_next_grapheme_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns grapheme start position closest to the [param pos].
*/
func (class) _shaped_text_prev_grapheme_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns array of the composite character boundaries.
*/
func (class) _shaped_text_get_character_breaks(impl func(ptr unsafe.Pointer, shaped gd.RID) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns composite character end position closest to the [param pos].
*/
func (class) _shaped_text_next_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns composite character start position closest to the [param pos].
*/
func (class) _shaped_text_prev_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns composite character position closest to the [param pos].
*/
func (class) _shaped_text_closest_character_pos(impl func(ptr unsafe.Pointer, shaped gd.RID, pos gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var shaped = gd.UnsafeGet[gd.RID](p_args,0)
		var pos = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shaped, pos)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
*/
func (class) _format_number(impl func(ptr unsafe.Pointer, number gd.String, language gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var number = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, number, language)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
*/
func (class) _parse_number(impl func(ptr unsafe.Pointer, number gd.String, language gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var number = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, number, language)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns percent sign used in the [param language].
*/
func (class) _percent_sign(impl func(ptr unsafe.Pointer, language gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, language)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Strips diacritics from the string.
*/
func (class) _strip_diacritics(impl func(ptr unsafe.Pointer, s gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if [param string] is a valid identifier.
*/
func (class) _is_valid_identifier(impl func(ptr unsafe.Pointer, s gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _is_valid_letter(impl func(ptr unsafe.Pointer, unicode gd.Int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var unicode = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, unicode)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
*/
func (class) _string_get_word_breaks(impl func(ptr unsafe.Pointer, s gd.String, language gd.String, chars_per_line gd.Int) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var chars_per_line = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language, chars_per_line)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns array of the composite character boundaries.
*/
func (class) _string_get_character_breaks(impl func(ptr unsafe.Pointer, s gd.String, language gd.String) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
*/
func (class) _is_confusable(impl func(ptr unsafe.Pointer, s gd.String, dict gd.PackedStringArray) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var dict = mmm.Let[gd.PackedStringArray](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, dict)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
*/
func (class) _spoof_check(impl func(ptr unsafe.Pointer, s gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the string converted to uppercase.
*/
func (class) _string_to_upper(impl func(ptr unsafe.Pointer, s gd.String, language gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the string converted to lowercase.
*/
func (class) _string_to_lower(impl func(ptr unsafe.Pointer, s gd.String, language gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Returns the string converted to title case.
*/
func (class) _string_to_title(impl func(ptr unsafe.Pointer, s gd.String, language gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var s = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var language = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, s, language)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
Default implementation of the BiDi algorithm override function. See [enum TextServer.StructuredTextParser] for more info.
*/
func (class) _parse_structured_text(impl func(ptr unsafe.Pointer, parser_type classdb.TextServerStructuredTextParser, args gd.Array, text gd.String) gd.ArrayOf[gd.Vector3i], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var parser_type = gd.UnsafeGet[classdb.TextServerStructuredTextParser](p_args,0)
		var args = mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var text = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parser_type, args, text)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
[b]Optional.[/b]
This method is called before text server is unregistered.
*/
func (class) _cleanup(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}


//go:nosplit
func (self class) AsTextServerExtension() Expert { return self[0].AsTextServerExtension() }


//go:nosplit
func (self Simple) AsTextServerExtension() Simple { return self[0].AsTextServerExtension() }


//go:nosplit
func (self class) AsTextServer() TextServer.Expert { return self[0].AsTextServer() }


//go:nosplit
func (self Simple) AsTextServer() TextServer.Simple { return self[0].AsTextServer() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_has_feature": return reflect.ValueOf(self._has_feature);
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_get_features": return reflect.ValueOf(self._get_features);
	case "_free_rid": return reflect.ValueOf(self._free_rid);
	case "_has": return reflect.ValueOf(self._has);
	case "_load_support_data": return reflect.ValueOf(self._load_support_data);
	case "_get_support_data_filename": return reflect.ValueOf(self._get_support_data_filename);
	case "_get_support_data_info": return reflect.ValueOf(self._get_support_data_info);
	case "_save_support_data": return reflect.ValueOf(self._save_support_data);
	case "_is_locale_right_to_left": return reflect.ValueOf(self._is_locale_right_to_left);
	case "_name_to_tag": return reflect.ValueOf(self._name_to_tag);
	case "_tag_to_name": return reflect.ValueOf(self._tag_to_name);
	case "_create_font": return reflect.ValueOf(self._create_font);
	case "_create_font_linked_variation": return reflect.ValueOf(self._create_font_linked_variation);
	case "_font_set_data": return reflect.ValueOf(self._font_set_data);
	case "_font_set_data_ptr": return reflect.ValueOf(self._font_set_data_ptr);
	case "_font_set_face_index": return reflect.ValueOf(self._font_set_face_index);
	case "_font_get_face_index": return reflect.ValueOf(self._font_get_face_index);
	case "_font_get_face_count": return reflect.ValueOf(self._font_get_face_count);
	case "_font_set_style": return reflect.ValueOf(self._font_set_style);
	case "_font_get_style": return reflect.ValueOf(self._font_get_style);
	case "_font_set_name": return reflect.ValueOf(self._font_set_name);
	case "_font_get_name": return reflect.ValueOf(self._font_get_name);
	case "_font_get_ot_name_strings": return reflect.ValueOf(self._font_get_ot_name_strings);
	case "_font_set_style_name": return reflect.ValueOf(self._font_set_style_name);
	case "_font_get_style_name": return reflect.ValueOf(self._font_get_style_name);
	case "_font_set_weight": return reflect.ValueOf(self._font_set_weight);
	case "_font_get_weight": return reflect.ValueOf(self._font_get_weight);
	case "_font_set_stretch": return reflect.ValueOf(self._font_set_stretch);
	case "_font_get_stretch": return reflect.ValueOf(self._font_get_stretch);
	case "_font_set_antialiasing": return reflect.ValueOf(self._font_set_antialiasing);
	case "_font_get_antialiasing": return reflect.ValueOf(self._font_get_antialiasing);
	case "_font_set_disable_embedded_bitmaps": return reflect.ValueOf(self._font_set_disable_embedded_bitmaps);
	case "_font_get_disable_embedded_bitmaps": return reflect.ValueOf(self._font_get_disable_embedded_bitmaps);
	case "_font_set_generate_mipmaps": return reflect.ValueOf(self._font_set_generate_mipmaps);
	case "_font_get_generate_mipmaps": return reflect.ValueOf(self._font_get_generate_mipmaps);
	case "_font_set_multichannel_signed_distance_field": return reflect.ValueOf(self._font_set_multichannel_signed_distance_field);
	case "_font_is_multichannel_signed_distance_field": return reflect.ValueOf(self._font_is_multichannel_signed_distance_field);
	case "_font_set_msdf_pixel_range": return reflect.ValueOf(self._font_set_msdf_pixel_range);
	case "_font_get_msdf_pixel_range": return reflect.ValueOf(self._font_get_msdf_pixel_range);
	case "_font_set_msdf_size": return reflect.ValueOf(self._font_set_msdf_size);
	case "_font_get_msdf_size": return reflect.ValueOf(self._font_get_msdf_size);
	case "_font_set_fixed_size": return reflect.ValueOf(self._font_set_fixed_size);
	case "_font_get_fixed_size": return reflect.ValueOf(self._font_get_fixed_size);
	case "_font_set_fixed_size_scale_mode": return reflect.ValueOf(self._font_set_fixed_size_scale_mode);
	case "_font_get_fixed_size_scale_mode": return reflect.ValueOf(self._font_get_fixed_size_scale_mode);
	case "_font_set_allow_system_fallback": return reflect.ValueOf(self._font_set_allow_system_fallback);
	case "_font_is_allow_system_fallback": return reflect.ValueOf(self._font_is_allow_system_fallback);
	case "_font_set_force_autohinter": return reflect.ValueOf(self._font_set_force_autohinter);
	case "_font_is_force_autohinter": return reflect.ValueOf(self._font_is_force_autohinter);
	case "_font_set_hinting": return reflect.ValueOf(self._font_set_hinting);
	case "_font_get_hinting": return reflect.ValueOf(self._font_get_hinting);
	case "_font_set_subpixel_positioning": return reflect.ValueOf(self._font_set_subpixel_positioning);
	case "_font_get_subpixel_positioning": return reflect.ValueOf(self._font_get_subpixel_positioning);
	case "_font_set_embolden": return reflect.ValueOf(self._font_set_embolden);
	case "_font_get_embolden": return reflect.ValueOf(self._font_get_embolden);
	case "_font_set_spacing": return reflect.ValueOf(self._font_set_spacing);
	case "_font_get_spacing": return reflect.ValueOf(self._font_get_spacing);
	case "_font_set_baseline_offset": return reflect.ValueOf(self._font_set_baseline_offset);
	case "_font_get_baseline_offset": return reflect.ValueOf(self._font_get_baseline_offset);
	case "_font_set_transform": return reflect.ValueOf(self._font_set_transform);
	case "_font_get_transform": return reflect.ValueOf(self._font_get_transform);
	case "_font_set_variation_coordinates": return reflect.ValueOf(self._font_set_variation_coordinates);
	case "_font_get_variation_coordinates": return reflect.ValueOf(self._font_get_variation_coordinates);
	case "_font_set_oversampling": return reflect.ValueOf(self._font_set_oversampling);
	case "_font_get_oversampling": return reflect.ValueOf(self._font_get_oversampling);
	case "_font_get_size_cache_list": return reflect.ValueOf(self._font_get_size_cache_list);
	case "_font_clear_size_cache": return reflect.ValueOf(self._font_clear_size_cache);
	case "_font_remove_size_cache": return reflect.ValueOf(self._font_remove_size_cache);
	case "_font_set_ascent": return reflect.ValueOf(self._font_set_ascent);
	case "_font_get_ascent": return reflect.ValueOf(self._font_get_ascent);
	case "_font_set_descent": return reflect.ValueOf(self._font_set_descent);
	case "_font_get_descent": return reflect.ValueOf(self._font_get_descent);
	case "_font_set_underline_position": return reflect.ValueOf(self._font_set_underline_position);
	case "_font_get_underline_position": return reflect.ValueOf(self._font_get_underline_position);
	case "_font_set_underline_thickness": return reflect.ValueOf(self._font_set_underline_thickness);
	case "_font_get_underline_thickness": return reflect.ValueOf(self._font_get_underline_thickness);
	case "_font_set_scale": return reflect.ValueOf(self._font_set_scale);
	case "_font_get_scale": return reflect.ValueOf(self._font_get_scale);
	case "_font_get_texture_count": return reflect.ValueOf(self._font_get_texture_count);
	case "_font_clear_textures": return reflect.ValueOf(self._font_clear_textures);
	case "_font_remove_texture": return reflect.ValueOf(self._font_remove_texture);
	case "_font_set_texture_image": return reflect.ValueOf(self._font_set_texture_image);
	case "_font_get_texture_image": return reflect.ValueOf(self._font_get_texture_image);
	case "_font_set_texture_offsets": return reflect.ValueOf(self._font_set_texture_offsets);
	case "_font_get_texture_offsets": return reflect.ValueOf(self._font_get_texture_offsets);
	case "_font_get_glyph_list": return reflect.ValueOf(self._font_get_glyph_list);
	case "_font_clear_glyphs": return reflect.ValueOf(self._font_clear_glyphs);
	case "_font_remove_glyph": return reflect.ValueOf(self._font_remove_glyph);
	case "_font_get_glyph_advance": return reflect.ValueOf(self._font_get_glyph_advance);
	case "_font_set_glyph_advance": return reflect.ValueOf(self._font_set_glyph_advance);
	case "_font_get_glyph_offset": return reflect.ValueOf(self._font_get_glyph_offset);
	case "_font_set_glyph_offset": return reflect.ValueOf(self._font_set_glyph_offset);
	case "_font_get_glyph_size": return reflect.ValueOf(self._font_get_glyph_size);
	case "_font_set_glyph_size": return reflect.ValueOf(self._font_set_glyph_size);
	case "_font_get_glyph_uv_rect": return reflect.ValueOf(self._font_get_glyph_uv_rect);
	case "_font_set_glyph_uv_rect": return reflect.ValueOf(self._font_set_glyph_uv_rect);
	case "_font_get_glyph_texture_idx": return reflect.ValueOf(self._font_get_glyph_texture_idx);
	case "_font_set_glyph_texture_idx": return reflect.ValueOf(self._font_set_glyph_texture_idx);
	case "_font_get_glyph_texture_rid": return reflect.ValueOf(self._font_get_glyph_texture_rid);
	case "_font_get_glyph_texture_size": return reflect.ValueOf(self._font_get_glyph_texture_size);
	case "_font_get_glyph_contours": return reflect.ValueOf(self._font_get_glyph_contours);
	case "_font_get_kerning_list": return reflect.ValueOf(self._font_get_kerning_list);
	case "_font_clear_kerning_map": return reflect.ValueOf(self._font_clear_kerning_map);
	case "_font_remove_kerning": return reflect.ValueOf(self._font_remove_kerning);
	case "_font_set_kerning": return reflect.ValueOf(self._font_set_kerning);
	case "_font_get_kerning": return reflect.ValueOf(self._font_get_kerning);
	case "_font_get_glyph_index": return reflect.ValueOf(self._font_get_glyph_index);
	case "_font_get_char_from_glyph_index": return reflect.ValueOf(self._font_get_char_from_glyph_index);
	case "_font_has_char": return reflect.ValueOf(self._font_has_char);
	case "_font_get_supported_chars": return reflect.ValueOf(self._font_get_supported_chars);
	case "_font_render_range": return reflect.ValueOf(self._font_render_range);
	case "_font_render_glyph": return reflect.ValueOf(self._font_render_glyph);
	case "_font_draw_glyph": return reflect.ValueOf(self._font_draw_glyph);
	case "_font_draw_glyph_outline": return reflect.ValueOf(self._font_draw_glyph_outline);
	case "_font_is_language_supported": return reflect.ValueOf(self._font_is_language_supported);
	case "_font_set_language_support_override": return reflect.ValueOf(self._font_set_language_support_override);
	case "_font_get_language_support_override": return reflect.ValueOf(self._font_get_language_support_override);
	case "_font_remove_language_support_override": return reflect.ValueOf(self._font_remove_language_support_override);
	case "_font_get_language_support_overrides": return reflect.ValueOf(self._font_get_language_support_overrides);
	case "_font_is_script_supported": return reflect.ValueOf(self._font_is_script_supported);
	case "_font_set_script_support_override": return reflect.ValueOf(self._font_set_script_support_override);
	case "_font_get_script_support_override": return reflect.ValueOf(self._font_get_script_support_override);
	case "_font_remove_script_support_override": return reflect.ValueOf(self._font_remove_script_support_override);
	case "_font_get_script_support_overrides": return reflect.ValueOf(self._font_get_script_support_overrides);
	case "_font_set_opentype_feature_overrides": return reflect.ValueOf(self._font_set_opentype_feature_overrides);
	case "_font_get_opentype_feature_overrides": return reflect.ValueOf(self._font_get_opentype_feature_overrides);
	case "_font_supported_feature_list": return reflect.ValueOf(self._font_supported_feature_list);
	case "_font_supported_variation_list": return reflect.ValueOf(self._font_supported_variation_list);
	case "_font_get_global_oversampling": return reflect.ValueOf(self._font_get_global_oversampling);
	case "_font_set_global_oversampling": return reflect.ValueOf(self._font_set_global_oversampling);
	case "_get_hex_code_box_size": return reflect.ValueOf(self._get_hex_code_box_size);
	case "_draw_hex_code_box": return reflect.ValueOf(self._draw_hex_code_box);
	case "_create_shaped_text": return reflect.ValueOf(self._create_shaped_text);
	case "_shaped_text_clear": return reflect.ValueOf(self._shaped_text_clear);
	case "_shaped_text_set_direction": return reflect.ValueOf(self._shaped_text_set_direction);
	case "_shaped_text_get_direction": return reflect.ValueOf(self._shaped_text_get_direction);
	case "_shaped_text_get_inferred_direction": return reflect.ValueOf(self._shaped_text_get_inferred_direction);
	case "_shaped_text_set_bidi_override": return reflect.ValueOf(self._shaped_text_set_bidi_override);
	case "_shaped_text_set_custom_punctuation": return reflect.ValueOf(self._shaped_text_set_custom_punctuation);
	case "_shaped_text_get_custom_punctuation": return reflect.ValueOf(self._shaped_text_get_custom_punctuation);
	case "_shaped_text_set_custom_ellipsis": return reflect.ValueOf(self._shaped_text_set_custom_ellipsis);
	case "_shaped_text_get_custom_ellipsis": return reflect.ValueOf(self._shaped_text_get_custom_ellipsis);
	case "_shaped_text_set_orientation": return reflect.ValueOf(self._shaped_text_set_orientation);
	case "_shaped_text_get_orientation": return reflect.ValueOf(self._shaped_text_get_orientation);
	case "_shaped_text_set_preserve_invalid": return reflect.ValueOf(self._shaped_text_set_preserve_invalid);
	case "_shaped_text_get_preserve_invalid": return reflect.ValueOf(self._shaped_text_get_preserve_invalid);
	case "_shaped_text_set_preserve_control": return reflect.ValueOf(self._shaped_text_set_preserve_control);
	case "_shaped_text_get_preserve_control": return reflect.ValueOf(self._shaped_text_get_preserve_control);
	case "_shaped_text_set_spacing": return reflect.ValueOf(self._shaped_text_set_spacing);
	case "_shaped_text_get_spacing": return reflect.ValueOf(self._shaped_text_get_spacing);
	case "_shaped_text_add_string": return reflect.ValueOf(self._shaped_text_add_string);
	case "_shaped_text_add_object": return reflect.ValueOf(self._shaped_text_add_object);
	case "_shaped_text_resize_object": return reflect.ValueOf(self._shaped_text_resize_object);
	case "_shaped_get_span_count": return reflect.ValueOf(self._shaped_get_span_count);
	case "_shaped_get_span_meta": return reflect.ValueOf(self._shaped_get_span_meta);
	case "_shaped_set_span_update_font": return reflect.ValueOf(self._shaped_set_span_update_font);
	case "_shaped_text_substr": return reflect.ValueOf(self._shaped_text_substr);
	case "_shaped_text_get_parent": return reflect.ValueOf(self._shaped_text_get_parent);
	case "_shaped_text_fit_to_width": return reflect.ValueOf(self._shaped_text_fit_to_width);
	case "_shaped_text_tab_align": return reflect.ValueOf(self._shaped_text_tab_align);
	case "_shaped_text_shape": return reflect.ValueOf(self._shaped_text_shape);
	case "_shaped_text_update_breaks": return reflect.ValueOf(self._shaped_text_update_breaks);
	case "_shaped_text_update_justification_ops": return reflect.ValueOf(self._shaped_text_update_justification_ops);
	case "_shaped_text_is_ready": return reflect.ValueOf(self._shaped_text_is_ready);
	case "_shaped_text_get_glyphs": return reflect.ValueOf(self._shaped_text_get_glyphs);
	case "_shaped_text_sort_logical": return reflect.ValueOf(self._shaped_text_sort_logical);
	case "_shaped_text_get_glyph_count": return reflect.ValueOf(self._shaped_text_get_glyph_count);
	case "_shaped_text_get_range": return reflect.ValueOf(self._shaped_text_get_range);
	case "_shaped_text_get_line_breaks_adv": return reflect.ValueOf(self._shaped_text_get_line_breaks_adv);
	case "_shaped_text_get_line_breaks": return reflect.ValueOf(self._shaped_text_get_line_breaks);
	case "_shaped_text_get_word_breaks": return reflect.ValueOf(self._shaped_text_get_word_breaks);
	case "_shaped_text_get_trim_pos": return reflect.ValueOf(self._shaped_text_get_trim_pos);
	case "_shaped_text_get_ellipsis_pos": return reflect.ValueOf(self._shaped_text_get_ellipsis_pos);
	case "_shaped_text_get_ellipsis_glyph_count": return reflect.ValueOf(self._shaped_text_get_ellipsis_glyph_count);
	case "_shaped_text_get_ellipsis_glyphs": return reflect.ValueOf(self._shaped_text_get_ellipsis_glyphs);
	case "_shaped_text_overrun_trim_to_width": return reflect.ValueOf(self._shaped_text_overrun_trim_to_width);
	case "_shaped_text_get_objects": return reflect.ValueOf(self._shaped_text_get_objects);
	case "_shaped_text_get_object_rect": return reflect.ValueOf(self._shaped_text_get_object_rect);
	case "_shaped_text_get_object_range": return reflect.ValueOf(self._shaped_text_get_object_range);
	case "_shaped_text_get_object_glyph": return reflect.ValueOf(self._shaped_text_get_object_glyph);
	case "_shaped_text_get_size": return reflect.ValueOf(self._shaped_text_get_size);
	case "_shaped_text_get_ascent": return reflect.ValueOf(self._shaped_text_get_ascent);
	case "_shaped_text_get_descent": return reflect.ValueOf(self._shaped_text_get_descent);
	case "_shaped_text_get_width": return reflect.ValueOf(self._shaped_text_get_width);
	case "_shaped_text_get_underline_position": return reflect.ValueOf(self._shaped_text_get_underline_position);
	case "_shaped_text_get_underline_thickness": return reflect.ValueOf(self._shaped_text_get_underline_thickness);
	case "_shaped_text_get_dominant_direction_in_range": return reflect.ValueOf(self._shaped_text_get_dominant_direction_in_range);
	case "_shaped_text_get_carets": return reflect.ValueOf(self._shaped_text_get_carets);
	case "_shaped_text_get_selection": return reflect.ValueOf(self._shaped_text_get_selection);
	case "_shaped_text_hit_test_grapheme": return reflect.ValueOf(self._shaped_text_hit_test_grapheme);
	case "_shaped_text_hit_test_position": return reflect.ValueOf(self._shaped_text_hit_test_position);
	case "_shaped_text_draw": return reflect.ValueOf(self._shaped_text_draw);
	case "_shaped_text_draw_outline": return reflect.ValueOf(self._shaped_text_draw_outline);
	case "_shaped_text_get_grapheme_bounds": return reflect.ValueOf(self._shaped_text_get_grapheme_bounds);
	case "_shaped_text_next_grapheme_pos": return reflect.ValueOf(self._shaped_text_next_grapheme_pos);
	case "_shaped_text_prev_grapheme_pos": return reflect.ValueOf(self._shaped_text_prev_grapheme_pos);
	case "_shaped_text_get_character_breaks": return reflect.ValueOf(self._shaped_text_get_character_breaks);
	case "_shaped_text_next_character_pos": return reflect.ValueOf(self._shaped_text_next_character_pos);
	case "_shaped_text_prev_character_pos": return reflect.ValueOf(self._shaped_text_prev_character_pos);
	case "_shaped_text_closest_character_pos": return reflect.ValueOf(self._shaped_text_closest_character_pos);
	case "_format_number": return reflect.ValueOf(self._format_number);
	case "_parse_number": return reflect.ValueOf(self._parse_number);
	case "_percent_sign": return reflect.ValueOf(self._percent_sign);
	case "_strip_diacritics": return reflect.ValueOf(self._strip_diacritics);
	case "_is_valid_identifier": return reflect.ValueOf(self._is_valid_identifier);
	case "_is_valid_letter": return reflect.ValueOf(self._is_valid_letter);
	case "_string_get_word_breaks": return reflect.ValueOf(self._string_get_word_breaks);
	case "_string_get_character_breaks": return reflect.ValueOf(self._string_get_character_breaks);
	case "_is_confusable": return reflect.ValueOf(self._is_confusable);
	case "_spoof_check": return reflect.ValueOf(self._spoof_check);
	case "_string_to_upper": return reflect.ValueOf(self._string_to_upper);
	case "_string_to_lower": return reflect.ValueOf(self._string_to_lower);
	case "_string_to_title": return reflect.ValueOf(self._string_to_title);
	case "_parse_structured_text": return reflect.ValueOf(self._parse_structured_text);
	case "_cleanup": return reflect.ValueOf(self._cleanup);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_has_feature": return reflect.ValueOf(self._has_feature);
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_get_features": return reflect.ValueOf(self._get_features);
	case "_free_rid": return reflect.ValueOf(self._free_rid);
	case "_has": return reflect.ValueOf(self._has);
	case "_load_support_data": return reflect.ValueOf(self._load_support_data);
	case "_get_support_data_filename": return reflect.ValueOf(self._get_support_data_filename);
	case "_get_support_data_info": return reflect.ValueOf(self._get_support_data_info);
	case "_save_support_data": return reflect.ValueOf(self._save_support_data);
	case "_is_locale_right_to_left": return reflect.ValueOf(self._is_locale_right_to_left);
	case "_name_to_tag": return reflect.ValueOf(self._name_to_tag);
	case "_tag_to_name": return reflect.ValueOf(self._tag_to_name);
	case "_create_font": return reflect.ValueOf(self._create_font);
	case "_create_font_linked_variation": return reflect.ValueOf(self._create_font_linked_variation);
	case "_font_set_data": return reflect.ValueOf(self._font_set_data);
	case "_font_set_data_ptr": return reflect.ValueOf(self._font_set_data_ptr);
	case "_font_set_face_index": return reflect.ValueOf(self._font_set_face_index);
	case "_font_get_face_index": return reflect.ValueOf(self._font_get_face_index);
	case "_font_get_face_count": return reflect.ValueOf(self._font_get_face_count);
	case "_font_set_style": return reflect.ValueOf(self._font_set_style);
	case "_font_get_style": return reflect.ValueOf(self._font_get_style);
	case "_font_set_name": return reflect.ValueOf(self._font_set_name);
	case "_font_get_name": return reflect.ValueOf(self._font_get_name);
	case "_font_get_ot_name_strings": return reflect.ValueOf(self._font_get_ot_name_strings);
	case "_font_set_style_name": return reflect.ValueOf(self._font_set_style_name);
	case "_font_get_style_name": return reflect.ValueOf(self._font_get_style_name);
	case "_font_set_weight": return reflect.ValueOf(self._font_set_weight);
	case "_font_get_weight": return reflect.ValueOf(self._font_get_weight);
	case "_font_set_stretch": return reflect.ValueOf(self._font_set_stretch);
	case "_font_get_stretch": return reflect.ValueOf(self._font_get_stretch);
	case "_font_set_antialiasing": return reflect.ValueOf(self._font_set_antialiasing);
	case "_font_get_antialiasing": return reflect.ValueOf(self._font_get_antialiasing);
	case "_font_set_disable_embedded_bitmaps": return reflect.ValueOf(self._font_set_disable_embedded_bitmaps);
	case "_font_get_disable_embedded_bitmaps": return reflect.ValueOf(self._font_get_disable_embedded_bitmaps);
	case "_font_set_generate_mipmaps": return reflect.ValueOf(self._font_set_generate_mipmaps);
	case "_font_get_generate_mipmaps": return reflect.ValueOf(self._font_get_generate_mipmaps);
	case "_font_set_multichannel_signed_distance_field": return reflect.ValueOf(self._font_set_multichannel_signed_distance_field);
	case "_font_is_multichannel_signed_distance_field": return reflect.ValueOf(self._font_is_multichannel_signed_distance_field);
	case "_font_set_msdf_pixel_range": return reflect.ValueOf(self._font_set_msdf_pixel_range);
	case "_font_get_msdf_pixel_range": return reflect.ValueOf(self._font_get_msdf_pixel_range);
	case "_font_set_msdf_size": return reflect.ValueOf(self._font_set_msdf_size);
	case "_font_get_msdf_size": return reflect.ValueOf(self._font_get_msdf_size);
	case "_font_set_fixed_size": return reflect.ValueOf(self._font_set_fixed_size);
	case "_font_get_fixed_size": return reflect.ValueOf(self._font_get_fixed_size);
	case "_font_set_fixed_size_scale_mode": return reflect.ValueOf(self._font_set_fixed_size_scale_mode);
	case "_font_get_fixed_size_scale_mode": return reflect.ValueOf(self._font_get_fixed_size_scale_mode);
	case "_font_set_allow_system_fallback": return reflect.ValueOf(self._font_set_allow_system_fallback);
	case "_font_is_allow_system_fallback": return reflect.ValueOf(self._font_is_allow_system_fallback);
	case "_font_set_force_autohinter": return reflect.ValueOf(self._font_set_force_autohinter);
	case "_font_is_force_autohinter": return reflect.ValueOf(self._font_is_force_autohinter);
	case "_font_set_hinting": return reflect.ValueOf(self._font_set_hinting);
	case "_font_get_hinting": return reflect.ValueOf(self._font_get_hinting);
	case "_font_set_subpixel_positioning": return reflect.ValueOf(self._font_set_subpixel_positioning);
	case "_font_get_subpixel_positioning": return reflect.ValueOf(self._font_get_subpixel_positioning);
	case "_font_set_embolden": return reflect.ValueOf(self._font_set_embolden);
	case "_font_get_embolden": return reflect.ValueOf(self._font_get_embolden);
	case "_font_set_spacing": return reflect.ValueOf(self._font_set_spacing);
	case "_font_get_spacing": return reflect.ValueOf(self._font_get_spacing);
	case "_font_set_baseline_offset": return reflect.ValueOf(self._font_set_baseline_offset);
	case "_font_get_baseline_offset": return reflect.ValueOf(self._font_get_baseline_offset);
	case "_font_set_transform": return reflect.ValueOf(self._font_set_transform);
	case "_font_get_transform": return reflect.ValueOf(self._font_get_transform);
	case "_font_set_variation_coordinates": return reflect.ValueOf(self._font_set_variation_coordinates);
	case "_font_get_variation_coordinates": return reflect.ValueOf(self._font_get_variation_coordinates);
	case "_font_set_oversampling": return reflect.ValueOf(self._font_set_oversampling);
	case "_font_get_oversampling": return reflect.ValueOf(self._font_get_oversampling);
	case "_font_get_size_cache_list": return reflect.ValueOf(self._font_get_size_cache_list);
	case "_font_clear_size_cache": return reflect.ValueOf(self._font_clear_size_cache);
	case "_font_remove_size_cache": return reflect.ValueOf(self._font_remove_size_cache);
	case "_font_set_ascent": return reflect.ValueOf(self._font_set_ascent);
	case "_font_get_ascent": return reflect.ValueOf(self._font_get_ascent);
	case "_font_set_descent": return reflect.ValueOf(self._font_set_descent);
	case "_font_get_descent": return reflect.ValueOf(self._font_get_descent);
	case "_font_set_underline_position": return reflect.ValueOf(self._font_set_underline_position);
	case "_font_get_underline_position": return reflect.ValueOf(self._font_get_underline_position);
	case "_font_set_underline_thickness": return reflect.ValueOf(self._font_set_underline_thickness);
	case "_font_get_underline_thickness": return reflect.ValueOf(self._font_get_underline_thickness);
	case "_font_set_scale": return reflect.ValueOf(self._font_set_scale);
	case "_font_get_scale": return reflect.ValueOf(self._font_get_scale);
	case "_font_get_texture_count": return reflect.ValueOf(self._font_get_texture_count);
	case "_font_clear_textures": return reflect.ValueOf(self._font_clear_textures);
	case "_font_remove_texture": return reflect.ValueOf(self._font_remove_texture);
	case "_font_set_texture_image": return reflect.ValueOf(self._font_set_texture_image);
	case "_font_get_texture_image": return reflect.ValueOf(self._font_get_texture_image);
	case "_font_set_texture_offsets": return reflect.ValueOf(self._font_set_texture_offsets);
	case "_font_get_texture_offsets": return reflect.ValueOf(self._font_get_texture_offsets);
	case "_font_get_glyph_list": return reflect.ValueOf(self._font_get_glyph_list);
	case "_font_clear_glyphs": return reflect.ValueOf(self._font_clear_glyphs);
	case "_font_remove_glyph": return reflect.ValueOf(self._font_remove_glyph);
	case "_font_get_glyph_advance": return reflect.ValueOf(self._font_get_glyph_advance);
	case "_font_set_glyph_advance": return reflect.ValueOf(self._font_set_glyph_advance);
	case "_font_get_glyph_offset": return reflect.ValueOf(self._font_get_glyph_offset);
	case "_font_set_glyph_offset": return reflect.ValueOf(self._font_set_glyph_offset);
	case "_font_get_glyph_size": return reflect.ValueOf(self._font_get_glyph_size);
	case "_font_set_glyph_size": return reflect.ValueOf(self._font_set_glyph_size);
	case "_font_get_glyph_uv_rect": return reflect.ValueOf(self._font_get_glyph_uv_rect);
	case "_font_set_glyph_uv_rect": return reflect.ValueOf(self._font_set_glyph_uv_rect);
	case "_font_get_glyph_texture_idx": return reflect.ValueOf(self._font_get_glyph_texture_idx);
	case "_font_set_glyph_texture_idx": return reflect.ValueOf(self._font_set_glyph_texture_idx);
	case "_font_get_glyph_texture_rid": return reflect.ValueOf(self._font_get_glyph_texture_rid);
	case "_font_get_glyph_texture_size": return reflect.ValueOf(self._font_get_glyph_texture_size);
	case "_font_get_glyph_contours": return reflect.ValueOf(self._font_get_glyph_contours);
	case "_font_get_kerning_list": return reflect.ValueOf(self._font_get_kerning_list);
	case "_font_clear_kerning_map": return reflect.ValueOf(self._font_clear_kerning_map);
	case "_font_remove_kerning": return reflect.ValueOf(self._font_remove_kerning);
	case "_font_set_kerning": return reflect.ValueOf(self._font_set_kerning);
	case "_font_get_kerning": return reflect.ValueOf(self._font_get_kerning);
	case "_font_get_glyph_index": return reflect.ValueOf(self._font_get_glyph_index);
	case "_font_get_char_from_glyph_index": return reflect.ValueOf(self._font_get_char_from_glyph_index);
	case "_font_has_char": return reflect.ValueOf(self._font_has_char);
	case "_font_get_supported_chars": return reflect.ValueOf(self._font_get_supported_chars);
	case "_font_render_range": return reflect.ValueOf(self._font_render_range);
	case "_font_render_glyph": return reflect.ValueOf(self._font_render_glyph);
	case "_font_draw_glyph": return reflect.ValueOf(self._font_draw_glyph);
	case "_font_draw_glyph_outline": return reflect.ValueOf(self._font_draw_glyph_outline);
	case "_font_is_language_supported": return reflect.ValueOf(self._font_is_language_supported);
	case "_font_set_language_support_override": return reflect.ValueOf(self._font_set_language_support_override);
	case "_font_get_language_support_override": return reflect.ValueOf(self._font_get_language_support_override);
	case "_font_remove_language_support_override": return reflect.ValueOf(self._font_remove_language_support_override);
	case "_font_get_language_support_overrides": return reflect.ValueOf(self._font_get_language_support_overrides);
	case "_font_is_script_supported": return reflect.ValueOf(self._font_is_script_supported);
	case "_font_set_script_support_override": return reflect.ValueOf(self._font_set_script_support_override);
	case "_font_get_script_support_override": return reflect.ValueOf(self._font_get_script_support_override);
	case "_font_remove_script_support_override": return reflect.ValueOf(self._font_remove_script_support_override);
	case "_font_get_script_support_overrides": return reflect.ValueOf(self._font_get_script_support_overrides);
	case "_font_set_opentype_feature_overrides": return reflect.ValueOf(self._font_set_opentype_feature_overrides);
	case "_font_get_opentype_feature_overrides": return reflect.ValueOf(self._font_get_opentype_feature_overrides);
	case "_font_supported_feature_list": return reflect.ValueOf(self._font_supported_feature_list);
	case "_font_supported_variation_list": return reflect.ValueOf(self._font_supported_variation_list);
	case "_font_get_global_oversampling": return reflect.ValueOf(self._font_get_global_oversampling);
	case "_font_set_global_oversampling": return reflect.ValueOf(self._font_set_global_oversampling);
	case "_get_hex_code_box_size": return reflect.ValueOf(self._get_hex_code_box_size);
	case "_draw_hex_code_box": return reflect.ValueOf(self._draw_hex_code_box);
	case "_create_shaped_text": return reflect.ValueOf(self._create_shaped_text);
	case "_shaped_text_clear": return reflect.ValueOf(self._shaped_text_clear);
	case "_shaped_text_set_direction": return reflect.ValueOf(self._shaped_text_set_direction);
	case "_shaped_text_get_direction": return reflect.ValueOf(self._shaped_text_get_direction);
	case "_shaped_text_get_inferred_direction": return reflect.ValueOf(self._shaped_text_get_inferred_direction);
	case "_shaped_text_set_bidi_override": return reflect.ValueOf(self._shaped_text_set_bidi_override);
	case "_shaped_text_set_custom_punctuation": return reflect.ValueOf(self._shaped_text_set_custom_punctuation);
	case "_shaped_text_get_custom_punctuation": return reflect.ValueOf(self._shaped_text_get_custom_punctuation);
	case "_shaped_text_set_custom_ellipsis": return reflect.ValueOf(self._shaped_text_set_custom_ellipsis);
	case "_shaped_text_get_custom_ellipsis": return reflect.ValueOf(self._shaped_text_get_custom_ellipsis);
	case "_shaped_text_set_orientation": return reflect.ValueOf(self._shaped_text_set_orientation);
	case "_shaped_text_get_orientation": return reflect.ValueOf(self._shaped_text_get_orientation);
	case "_shaped_text_set_preserve_invalid": return reflect.ValueOf(self._shaped_text_set_preserve_invalid);
	case "_shaped_text_get_preserve_invalid": return reflect.ValueOf(self._shaped_text_get_preserve_invalid);
	case "_shaped_text_set_preserve_control": return reflect.ValueOf(self._shaped_text_set_preserve_control);
	case "_shaped_text_get_preserve_control": return reflect.ValueOf(self._shaped_text_get_preserve_control);
	case "_shaped_text_set_spacing": return reflect.ValueOf(self._shaped_text_set_spacing);
	case "_shaped_text_get_spacing": return reflect.ValueOf(self._shaped_text_get_spacing);
	case "_shaped_text_add_string": return reflect.ValueOf(self._shaped_text_add_string);
	case "_shaped_text_add_object": return reflect.ValueOf(self._shaped_text_add_object);
	case "_shaped_text_resize_object": return reflect.ValueOf(self._shaped_text_resize_object);
	case "_shaped_get_span_count": return reflect.ValueOf(self._shaped_get_span_count);
	case "_shaped_get_span_meta": return reflect.ValueOf(self._shaped_get_span_meta);
	case "_shaped_set_span_update_font": return reflect.ValueOf(self._shaped_set_span_update_font);
	case "_shaped_text_substr": return reflect.ValueOf(self._shaped_text_substr);
	case "_shaped_text_get_parent": return reflect.ValueOf(self._shaped_text_get_parent);
	case "_shaped_text_fit_to_width": return reflect.ValueOf(self._shaped_text_fit_to_width);
	case "_shaped_text_tab_align": return reflect.ValueOf(self._shaped_text_tab_align);
	case "_shaped_text_shape": return reflect.ValueOf(self._shaped_text_shape);
	case "_shaped_text_update_breaks": return reflect.ValueOf(self._shaped_text_update_breaks);
	case "_shaped_text_update_justification_ops": return reflect.ValueOf(self._shaped_text_update_justification_ops);
	case "_shaped_text_is_ready": return reflect.ValueOf(self._shaped_text_is_ready);
	case "_shaped_text_get_glyphs": return reflect.ValueOf(self._shaped_text_get_glyphs);
	case "_shaped_text_sort_logical": return reflect.ValueOf(self._shaped_text_sort_logical);
	case "_shaped_text_get_glyph_count": return reflect.ValueOf(self._shaped_text_get_glyph_count);
	case "_shaped_text_get_range": return reflect.ValueOf(self._shaped_text_get_range);
	case "_shaped_text_get_line_breaks_adv": return reflect.ValueOf(self._shaped_text_get_line_breaks_adv);
	case "_shaped_text_get_line_breaks": return reflect.ValueOf(self._shaped_text_get_line_breaks);
	case "_shaped_text_get_word_breaks": return reflect.ValueOf(self._shaped_text_get_word_breaks);
	case "_shaped_text_get_trim_pos": return reflect.ValueOf(self._shaped_text_get_trim_pos);
	case "_shaped_text_get_ellipsis_pos": return reflect.ValueOf(self._shaped_text_get_ellipsis_pos);
	case "_shaped_text_get_ellipsis_glyph_count": return reflect.ValueOf(self._shaped_text_get_ellipsis_glyph_count);
	case "_shaped_text_get_ellipsis_glyphs": return reflect.ValueOf(self._shaped_text_get_ellipsis_glyphs);
	case "_shaped_text_overrun_trim_to_width": return reflect.ValueOf(self._shaped_text_overrun_trim_to_width);
	case "_shaped_text_get_objects": return reflect.ValueOf(self._shaped_text_get_objects);
	case "_shaped_text_get_object_rect": return reflect.ValueOf(self._shaped_text_get_object_rect);
	case "_shaped_text_get_object_range": return reflect.ValueOf(self._shaped_text_get_object_range);
	case "_shaped_text_get_object_glyph": return reflect.ValueOf(self._shaped_text_get_object_glyph);
	case "_shaped_text_get_size": return reflect.ValueOf(self._shaped_text_get_size);
	case "_shaped_text_get_ascent": return reflect.ValueOf(self._shaped_text_get_ascent);
	case "_shaped_text_get_descent": return reflect.ValueOf(self._shaped_text_get_descent);
	case "_shaped_text_get_width": return reflect.ValueOf(self._shaped_text_get_width);
	case "_shaped_text_get_underline_position": return reflect.ValueOf(self._shaped_text_get_underline_position);
	case "_shaped_text_get_underline_thickness": return reflect.ValueOf(self._shaped_text_get_underline_thickness);
	case "_shaped_text_get_dominant_direction_in_range": return reflect.ValueOf(self._shaped_text_get_dominant_direction_in_range);
	case "_shaped_text_get_carets": return reflect.ValueOf(self._shaped_text_get_carets);
	case "_shaped_text_get_selection": return reflect.ValueOf(self._shaped_text_get_selection);
	case "_shaped_text_hit_test_grapheme": return reflect.ValueOf(self._shaped_text_hit_test_grapheme);
	case "_shaped_text_hit_test_position": return reflect.ValueOf(self._shaped_text_hit_test_position);
	case "_shaped_text_draw": return reflect.ValueOf(self._shaped_text_draw);
	case "_shaped_text_draw_outline": return reflect.ValueOf(self._shaped_text_draw_outline);
	case "_shaped_text_get_grapheme_bounds": return reflect.ValueOf(self._shaped_text_get_grapheme_bounds);
	case "_shaped_text_next_grapheme_pos": return reflect.ValueOf(self._shaped_text_next_grapheme_pos);
	case "_shaped_text_prev_grapheme_pos": return reflect.ValueOf(self._shaped_text_prev_grapheme_pos);
	case "_shaped_text_get_character_breaks": return reflect.ValueOf(self._shaped_text_get_character_breaks);
	case "_shaped_text_next_character_pos": return reflect.ValueOf(self._shaped_text_next_character_pos);
	case "_shaped_text_prev_character_pos": return reflect.ValueOf(self._shaped_text_prev_character_pos);
	case "_shaped_text_closest_character_pos": return reflect.ValueOf(self._shaped_text_closest_character_pos);
	case "_format_number": return reflect.ValueOf(self._format_number);
	case "_parse_number": return reflect.ValueOf(self._parse_number);
	case "_percent_sign": return reflect.ValueOf(self._percent_sign);
	case "_strip_diacritics": return reflect.ValueOf(self._strip_diacritics);
	case "_is_valid_identifier": return reflect.ValueOf(self._is_valid_identifier);
	case "_is_valid_letter": return reflect.ValueOf(self._is_valid_letter);
	case "_string_get_word_breaks": return reflect.ValueOf(self._string_get_word_breaks);
	case "_string_get_character_breaks": return reflect.ValueOf(self._string_get_character_breaks);
	case "_is_confusable": return reflect.ValueOf(self._is_confusable);
	case "_spoof_check": return reflect.ValueOf(self._spoof_check);
	case "_string_to_upper": return reflect.ValueOf(self._string_to_upper);
	case "_string_to_lower": return reflect.ValueOf(self._string_to_lower);
	case "_string_to_title": return reflect.ValueOf(self._string_to_title);
	case "_parse_structured_text": return reflect.ValueOf(self._parse_structured_text);
	case "_cleanup": return reflect.ValueOf(self._cleanup);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TextServerExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
