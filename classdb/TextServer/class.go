// Package TextServer provides methods for working with TextServer object instances.
package TextServer

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/Packed"
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
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

/*
[TextServer] is the API backend for managing fonts and rendering text.
[b]Note:[/b] This is a low-level API, consider using [TextLine], [TextParagraph], and [Font] classes instead.
This is an abstract class, so to get the currently active [TextServer] instance, use the following code:
[codeblocks]
[gdscript]
var ts = TextServerManager.get_primary_interface()
[/gdscript]
[csharp]
var ts = TextServerManager.GetPrimaryInterface();
[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.TextServer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTextServer() Instance
}

/*
Returns [code]true[/code] if the server supports a feature.
*/
func (self Instance) HasFeature(feature gdclass.TextServerFeature) bool { //gd:TextServer.has_feature
	return bool(class(self).HasFeature(feature))
}

/*
Returns the name of the server interface.
*/
func (self Instance) GetName() string { //gd:TextServer.get_name
	return string(class(self).GetName().String())
}

/*
Returns text server features, see [enum Feature].
*/
func (self Instance) GetFeatures() int { //gd:TextServer.get_features
	return int(int(class(self).GetFeatures()))
}

/*
Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
[b]Note:[/b] This function should be called before any other TextServer functions used, otherwise it won't have any effect.
*/
func (self Instance) LoadSupportData(filename string) bool { //gd:TextServer.load_support_data
	return bool(class(self).LoadSupportData(String.New(filename)))
}

/*
Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
*/
func (self Instance) GetSupportDataFilename() string { //gd:TextServer.get_support_data_filename
	return string(class(self).GetSupportDataFilename().String())
}

/*
Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
*/
func (self Instance) GetSupportDataInfo() string { //gd:TextServer.get_support_data_info
	return string(class(self).GetSupportDataInfo().String())
}

/*
Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
[b]Note:[/b] This function is used by during project export, to include TextServer database.
*/
func (self Instance) SaveSupportData(filename string) bool { //gd:TextServer.save_support_data
	return bool(class(self).SaveSupportData(String.New(filename)))
}

/*
Returns [code]true[/code] if locale is right-to-left.
*/
func (self Instance) IsLocaleRightToLeft(locale string) bool { //gd:TextServer.is_locale_right_to_left
	return bool(class(self).IsLocaleRightToLeft(String.New(locale)))
}

/*
Converts readable feature, variation, script, or language name to OpenType tag.
*/
func (self Instance) NameToTag(name string) int { //gd:TextServer.name_to_tag
	return int(int(class(self).NameToTag(String.New(name))))
}

/*
Converts OpenType tag to readable feature, variation, script, or language name.
*/
func (self Instance) TagToName(tag int) string { //gd:TextServer.tag_to_name
	return string(class(self).TagToName(gd.Int(tag)).String())
}

/*
Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
*/
func (self Instance) Has(rid RID.Any) bool { //gd:TextServer.has
	return bool(class(self).Has(gd.RID(rid)))
}

/*
Frees an object created by this [TextServer].
*/
func (self Instance) FreeRid(rid RID.Any) { //gd:TextServer.free_rid
	class(self).FreeRid(gd.RID(rid))
}

/*
Creates a new, empty font cache entry resource. To free the resulting resource, use the [method free_rid] method.
*/
func (self Instance) CreateFont() RID.Font { //gd:TextServer.create_font
	return RID.Font(class(self).CreateFont())
}

/*
Creates a new variation existing font which is reusing the same glyph cache and font data. To free the resulting resource, use the [method free_rid] method.
*/
func (self Instance) CreateFontLinkedVariation(font_rid RID.Font) RID.Font { //gd:TextServer.create_font_linked_variation
	return RID.Font(class(self).CreateFontLinkedVariation(gd.RID(font_rid)))
}

/*
Sets font source data, e.g contents of the dynamic font source file.
*/
func (self Instance) FontSetData(font_rid RID.Font, data []byte) { //gd:TextServer.font_set_data
	class(self).FontSetData(gd.RID(font_rid), Packed.Bytes(Packed.New(data...)))
}

/*
Sets an active face index in the TrueType / OpenType collection.
*/
func (self Instance) FontSetFaceIndex(font_rid RID.Font, face_index int) { //gd:TextServer.font_set_face_index
	class(self).FontSetFaceIndex(gd.RID(font_rid), gd.Int(face_index))
}

/*
Returns an active face index in the TrueType / OpenType collection.
*/
func (self Instance) FontGetFaceIndex(font_rid RID.Font) int { //gd:TextServer.font_get_face_index
	return int(int(class(self).FontGetFaceIndex(gd.RID(font_rid))))
}

/*
Returns number of faces in the TrueType / OpenType collection.
*/
func (self Instance) FontGetFaceCount(font_rid RID.Font) int { //gd:TextServer.font_get_face_count
	return int(int(class(self).FontGetFaceCount(gd.RID(font_rid))))
}

/*
Sets the font style flags, see [enum FontStyle].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], [method font_set_embolden], or [method font_set_transform] instead.
*/
func (self Instance) FontSetStyle(font_rid RID.Font, style gdclass.TextServerFontStyle) { //gd:TextServer.font_set_style
	class(self).FontSetStyle(gd.RID(font_rid), style)
}

/*
Returns font style flags, see [enum FontStyle].
*/
func (self Instance) FontGetStyle(font_rid RID.Font) gdclass.TextServerFontStyle { //gd:TextServer.font_get_style
	return gdclass.TextServerFontStyle(class(self).FontGetStyle(gd.RID(font_rid)))
}

/*
Sets the font family name.
*/
func (self Instance) FontSetName(font_rid RID.Font, name string) { //gd:TextServer.font_set_name
	class(self).FontSetName(gd.RID(font_rid), String.New(name))
}

/*
Returns font family name.
*/
func (self Instance) FontGetName(font_rid RID.Font) string { //gd:TextServer.font_get_name
	return string(class(self).FontGetName(gd.RID(font_rid)).String())
}

/*
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
func (self Instance) FontGetOtNameStrings(font_rid RID.Font) map[string]string { //gd:TextServer.font_get_ot_name_strings
	return map[string]string(gd.DictionaryAs[map[string]string](class(self).FontGetOtNameStrings(gd.RID(font_rid))))
}

/*
Sets the font style name.
*/
func (self Instance) FontSetStyleName(font_rid RID.Font, name string) { //gd:TextServer.font_set_style_name
	class(self).FontSetStyleName(gd.RID(font_rid), String.New(name))
}

/*
Returns font style name.
*/
func (self Instance) FontGetStyleName(font_rid RID.Font) string { //gd:TextServer.font_get_style_name
	return string(class(self).FontGetStyleName(gd.RID(font_rid)).String())
}

/*
Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], or [method font_set_embolden] instead.
*/
func (self Instance) FontSetWeight(font_rid RID.Font, weight int) { //gd:TextServer.font_set_weight
	class(self).FontSetWeight(gd.RID(font_rid), gd.Int(weight))
}

/*
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
func (self Instance) FontGetWeight(font_rid RID.Font) int { //gd:TextServer.font_get_weight
	return int(int(class(self).FontGetWeight(gd.RID(font_rid))))
}

/*
Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], or [method font_set_transform] instead.
*/
func (self Instance) FontSetStretch(font_rid RID.Font, weight int) { //gd:TextServer.font_set_stretch
	class(self).FontSetStretch(gd.RID(font_rid), gd.Int(weight))
}

/*
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
func (self Instance) FontGetStretch(font_rid RID.Font) int { //gd:TextServer.font_get_stretch
	return int(int(class(self).FontGetStretch(gd.RID(font_rid))))
}

/*
Sets font anti-aliasing mode.
*/
func (self Instance) FontSetAntialiasing(font_rid RID.Font, antialiasing gdclass.TextServerFontAntialiasing) { //gd:TextServer.font_set_antialiasing
	class(self).FontSetAntialiasing(gd.RID(font_rid), antialiasing)
}

/*
Returns font anti-aliasing mode.
*/
func (self Instance) FontGetAntialiasing(font_rid RID.Font) gdclass.TextServerFontAntialiasing { //gd:TextServer.font_get_antialiasing
	return gdclass.TextServerFontAntialiasing(class(self).FontGetAntialiasing(gd.RID(font_rid)))
}

/*
If set to [code]true[/code], embedded font bitmap loading is disabled (bitmap-only and color fonts ignore this property).
*/
func (self Instance) FontSetDisableEmbeddedBitmaps(font_rid RID.Font, disable_embedded_bitmaps bool) { //gd:TextServer.font_set_disable_embedded_bitmaps
	class(self).FontSetDisableEmbeddedBitmaps(gd.RID(font_rid), disable_embedded_bitmaps)
}

/*
Returns whether the font's embedded bitmap loading is disabled.
*/
func (self Instance) FontGetDisableEmbeddedBitmaps(font_rid RID.Font) bool { //gd:TextServer.font_get_disable_embedded_bitmaps
	return bool(class(self).FontGetDisableEmbeddedBitmaps(gd.RID(font_rid)))
}

/*
If set to [code]true[/code] font texture mipmap generation is enabled.
*/
func (self Instance) FontSetGenerateMipmaps(font_rid RID.Font, generate_mipmaps bool) { //gd:TextServer.font_set_generate_mipmaps
	class(self).FontSetGenerateMipmaps(gd.RID(font_rid), generate_mipmaps)
}

/*
Returns [code]true[/code] if font texture mipmap generation is enabled.
*/
func (self Instance) FontGetGenerateMipmaps(font_rid RID.Font) bool { //gd:TextServer.font_get_generate_mipmaps
	return bool(class(self).FontGetGenerateMipmaps(gd.RID(font_rid)))
}

/*
If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
[b]Note:[/b] MSDF font rendering does not render glyphs with overlapping shapes correctly. Overlapping shapes are not valid per the OpenType standard, but are still commonly found in many font files, especially those converted by Google Fonts. To avoid issues with overlapping glyphs, consider downloading the font file directly from the type foundry instead of relying on Google Fonts.
*/
func (self Instance) FontSetMultichannelSignedDistanceField(font_rid RID.Font, msdf bool) { //gd:TextServer.font_set_multichannel_signed_distance_field
	class(self).FontSetMultichannelSignedDistanceField(gd.RID(font_rid), msdf)
}

/*
Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
*/
func (self Instance) FontIsMultichannelSignedDistanceField(font_rid RID.Font) bool { //gd:TextServer.font_is_multichannel_signed_distance_field
	return bool(class(self).FontIsMultichannelSignedDistanceField(gd.RID(font_rid)))
}

/*
Sets the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (self Instance) FontSetMsdfPixelRange(font_rid RID.Font, msdf_pixel_range int) { //gd:TextServer.font_set_msdf_pixel_range
	class(self).FontSetMsdfPixelRange(gd.RID(font_rid), gd.Int(msdf_pixel_range))
}

/*
Returns the width of the range around the shape between the minimum and maximum representable signed distance.
*/
func (self Instance) FontGetMsdfPixelRange(font_rid RID.Font) int { //gd:TextServer.font_get_msdf_pixel_range
	return int(int(class(self).FontGetMsdfPixelRange(gd.RID(font_rid))))
}

/*
Sets source font size used to generate MSDF textures.
*/
func (self Instance) FontSetMsdfSize(font_rid RID.Font, msdf_size int) { //gd:TextServer.font_set_msdf_size
	class(self).FontSetMsdfSize(gd.RID(font_rid), gd.Int(msdf_size))
}

/*
Returns source font size used to generate MSDF textures.
*/
func (self Instance) FontGetMsdfSize(font_rid RID.Font) int { //gd:TextServer.font_get_msdf_size
	return int(int(class(self).FontGetMsdfSize(gd.RID(font_rid))))
}

/*
Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
*/
func (self Instance) FontSetFixedSize(font_rid RID.Font, fixed_size int) { //gd:TextServer.font_set_fixed_size
	class(self).FontSetFixedSize(gd.RID(font_rid), gd.Int(fixed_size))
}

/*
Returns bitmap font fixed size.
*/
func (self Instance) FontGetFixedSize(font_rid RID.Font) int { //gd:TextServer.font_get_fixed_size
	return int(int(class(self).FontGetFixedSize(gd.RID(font_rid))))
}

/*
Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
*/
func (self Instance) FontSetFixedSizeScaleMode(font_rid RID.Font, fixed_size_scale_mode gdclass.TextServerFixedSizeScaleMode) { //gd:TextServer.font_set_fixed_size_scale_mode
	class(self).FontSetFixedSizeScaleMode(gd.RID(font_rid), fixed_size_scale_mode)
}

/*
Returns bitmap font scaling mode.
*/
func (self Instance) FontGetFixedSizeScaleMode(font_rid RID.Font) gdclass.TextServerFixedSizeScaleMode { //gd:TextServer.font_get_fixed_size_scale_mode
	return gdclass.TextServerFixedSizeScaleMode(class(self).FontGetFixedSizeScaleMode(gd.RID(font_rid)))
}

/*
If set to [code]true[/code], system fonts can be automatically used as fallbacks.
*/
func (self Instance) FontSetAllowSystemFallback(font_rid RID.Font, allow_system_fallback bool) { //gd:TextServer.font_set_allow_system_fallback
	class(self).FontSetAllowSystemFallback(gd.RID(font_rid), allow_system_fallback)
}

/*
Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
*/
func (self Instance) FontIsAllowSystemFallback(font_rid RID.Font) bool { //gd:TextServer.font_is_allow_system_fallback
	return bool(class(self).FontIsAllowSystemFallback(gd.RID(font_rid)))
}

/*
If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
*/
func (self Instance) FontSetForceAutohinter(font_rid RID.Font, force_autohinter bool) { //gd:TextServer.font_set_force_autohinter
	class(self).FontSetForceAutohinter(gd.RID(font_rid), force_autohinter)
}

/*
Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting. Used by dynamic fonts only.
*/
func (self Instance) FontIsForceAutohinter(font_rid RID.Font) bool { //gd:TextServer.font_is_force_autohinter
	return bool(class(self).FontIsForceAutohinter(gd.RID(font_rid)))
}

/*
Sets font hinting mode. Used by dynamic fonts only.
*/
func (self Instance) FontSetHinting(font_rid RID.Font, hinting gdclass.TextServerHinting) { //gd:TextServer.font_set_hinting
	class(self).FontSetHinting(gd.RID(font_rid), hinting)
}

/*
Returns the font hinting mode. Used by dynamic fonts only.
*/
func (self Instance) FontGetHinting(font_rid RID.Font) gdclass.TextServerHinting { //gd:TextServer.font_get_hinting
	return gdclass.TextServerHinting(class(self).FontGetHinting(gd.RID(font_rid)))
}

/*
Sets font subpixel glyph positioning mode.
*/
func (self Instance) FontSetSubpixelPositioning(font_rid RID.Font, subpixel_positioning gdclass.TextServerSubpixelPositioning) { //gd:TextServer.font_set_subpixel_positioning
	class(self).FontSetSubpixelPositioning(gd.RID(font_rid), subpixel_positioning)
}

/*
Returns font subpixel glyph positioning mode.
*/
func (self Instance) FontGetSubpixelPositioning(font_rid RID.Font) gdclass.TextServerSubpixelPositioning { //gd:TextServer.font_get_subpixel_positioning
	return gdclass.TextServerSubpixelPositioning(class(self).FontGetSubpixelPositioning(gd.RID(font_rid)))
}

/*
Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (self Instance) FontSetEmbolden(font_rid RID.Font, strength Float.X) { //gd:TextServer.font_set_embolden
	class(self).FontSetEmbolden(gd.RID(font_rid), gd.Float(strength))
}

/*
Returns font embolden strength.
*/
func (self Instance) FontGetEmbolden(font_rid RID.Font) Float.X { //gd:TextServer.font_get_embolden
	return Float.X(Float.X(class(self).FontGetEmbolden(gd.RID(font_rid))))
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
func (self Instance) FontSetSpacing(font_rid RID.Font, spacing gdclass.TextServerSpacingType, value int) { //gd:TextServer.font_set_spacing
	class(self).FontSetSpacing(gd.RID(font_rid), spacing, gd.Int(value))
}

/*
Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
func (self Instance) FontGetSpacing(font_rid RID.Font, spacing gdclass.TextServerSpacingType) int { //gd:TextServer.font_get_spacing
	return int(int(class(self).FontGetSpacing(gd.RID(font_rid), spacing)))
}

/*
Sets extra baseline offset (as a fraction of font height).
*/
func (self Instance) FontSetBaselineOffset(font_rid RID.Font, baseline_offset Float.X) { //gd:TextServer.font_set_baseline_offset
	class(self).FontSetBaselineOffset(gd.RID(font_rid), gd.Float(baseline_offset))
}

/*
Returns extra baseline offset (as a fraction of font height).
*/
func (self Instance) FontGetBaselineOffset(font_rid RID.Font) Float.X { //gd:TextServer.font_get_baseline_offset
	return Float.X(Float.X(class(self).FontGetBaselineOffset(gd.RID(font_rid))))
}

/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
For example, to simulate italic typeface by slanting, apply the following transform [code]Transform2D(1.0, slant, 0.0, 1.0, 0.0, 0.0)[/code].
*/
func (self Instance) FontSetTransform(font_rid RID.Font, transform Transform2D.OriginXY) { //gd:TextServer.font_set_transform
	class(self).FontSetTransform(gd.RID(font_rid), gd.Transform2D(transform))
}

/*
Returns 2D transform applied to the font outlines.
*/
func (self Instance) FontGetTransform(font_rid RID.Font) Transform2D.OriginXY { //gd:TextServer.font_get_transform
	return Transform2D.OriginXY(class(self).FontGetTransform(gd.RID(font_rid)))
}

/*
Sets variation coordinates for the specified font cache entry. See [method font_supported_variation_list] for more info.
*/
func (self Instance) FontSetVariationCoordinates(font_rid RID.Font, variation_coordinates map[string]float32) { //gd:TextServer.font_set_variation_coordinates
	class(self).FontSetVariationCoordinates(gd.RID(font_rid), gd.DictionaryFromMap(variation_coordinates))
}

/*
Returns variation coordinates for the specified font cache entry. See [method font_supported_variation_list] for more info.
*/
func (self Instance) FontGetVariationCoordinates(font_rid RID.Font) map[string]float32 { //gd:TextServer.font_get_variation_coordinates
	return map[string]float32(gd.DictionaryAs[map[string]float32](class(self).FontGetVariationCoordinates(gd.RID(font_rid))))
}

/*
Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (self Instance) FontSetOversampling(font_rid RID.Font, oversampling Float.X) { //gd:TextServer.font_set_oversampling
	class(self).FontSetOversampling(gd.RID(font_rid), gd.Float(oversampling))
}

/*
Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
func (self Instance) FontGetOversampling(font_rid RID.Font) Float.X { //gd:TextServer.font_get_oversampling
	return Float.X(Float.X(class(self).FontGetOversampling(gd.RID(font_rid))))
}

/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
func (self Instance) FontGetSizeCacheList(font_rid RID.Font) []Vector2i.XY { //gd:TextServer.font_get_size_cache_list
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).FontGetSizeCacheList(gd.RID(font_rid)))))
}

/*
Removes all font sizes from the cache entry.
*/
func (self Instance) FontClearSizeCache(font_rid RID.Font) { //gd:TextServer.font_clear_size_cache
	class(self).FontClearSizeCache(gd.RID(font_rid))
}

/*
Removes specified font size from the cache entry.
*/
func (self Instance) FontRemoveSizeCache(font_rid RID.Font, size Vector2i.XY) { //gd:TextServer.font_remove_size_cache
	class(self).FontRemoveSizeCache(gd.RID(font_rid), gd.Vector2i(size))
}

/*
Sets the font ascent (number of pixels above the baseline).
*/
func (self Instance) FontSetAscent(font_rid RID.Font, size int, ascent Float.X) { //gd:TextServer.font_set_ascent
	class(self).FontSetAscent(gd.RID(font_rid), gd.Int(size), gd.Float(ascent))
}

/*
Returns the font ascent (number of pixels above the baseline).
*/
func (self Instance) FontGetAscent(font_rid RID.Font, size int) Float.X { //gd:TextServer.font_get_ascent
	return Float.X(Float.X(class(self).FontGetAscent(gd.RID(font_rid), gd.Int(size))))
}

/*
Sets the font descent (number of pixels below the baseline).
*/
func (self Instance) FontSetDescent(font_rid RID.Font, size int, descent Float.X) { //gd:TextServer.font_set_descent
	class(self).FontSetDescent(gd.RID(font_rid), gd.Int(size), gd.Float(descent))
}

/*
Returns the font descent (number of pixels below the baseline).
*/
func (self Instance) FontGetDescent(font_rid RID.Font, size int) Float.X { //gd:TextServer.font_get_descent
	return Float.X(Float.X(class(self).FontGetDescent(gd.RID(font_rid), gd.Int(size))))
}

/*
Sets pixel offset of the underline below the baseline.
*/
func (self Instance) FontSetUnderlinePosition(font_rid RID.Font, size int, underline_position Float.X) { //gd:TextServer.font_set_underline_position
	class(self).FontSetUnderlinePosition(gd.RID(font_rid), gd.Int(size), gd.Float(underline_position))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Instance) FontGetUnderlinePosition(font_rid RID.Font, size int) Float.X { //gd:TextServer.font_get_underline_position
	return Float.X(Float.X(class(self).FontGetUnderlinePosition(gd.RID(font_rid), gd.Int(size))))
}

/*
Sets thickness of the underline in pixels.
*/
func (self Instance) FontSetUnderlineThickness(font_rid RID.Font, size int, underline_thickness Float.X) { //gd:TextServer.font_set_underline_thickness
	class(self).FontSetUnderlineThickness(gd.RID(font_rid), gd.Int(size), gd.Float(underline_thickness))
}

/*
Returns thickness of the underline in pixels.
*/
func (self Instance) FontGetUnderlineThickness(font_rid RID.Font, size int) Float.X { //gd:TextServer.font_get_underline_thickness
	return Float.X(Float.X(class(self).FontGetUnderlineThickness(gd.RID(font_rid), gd.Int(size))))
}

/*
Sets scaling factor of the color bitmap font.
*/
func (self Instance) FontSetScale(font_rid RID.Font, size int, scale Float.X) { //gd:TextServer.font_set_scale
	class(self).FontSetScale(gd.RID(font_rid), gd.Int(size), gd.Float(scale))
}

/*
Returns scaling factor of the color bitmap font.
*/
func (self Instance) FontGetScale(font_rid RID.Font, size int) Float.X { //gd:TextServer.font_get_scale
	return Float.X(Float.X(class(self).FontGetScale(gd.RID(font_rid), gd.Int(size))))
}

/*
Returns number of textures used by font cache entry.
*/
func (self Instance) FontGetTextureCount(font_rid RID.Font, size Vector2i.XY) int { //gd:TextServer.font_get_texture_count
	return int(int(class(self).FontGetTextureCount(gd.RID(font_rid), gd.Vector2i(size))))
}

/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method font_remove_glyph] to remove them manually.
*/
func (self Instance) FontClearTextures(font_rid RID.Font, size Vector2i.XY) { //gd:TextServer.font_clear_textures
	class(self).FontClearTextures(gd.RID(font_rid), gd.Vector2i(size))
}

/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, remove them manually, using [method font_remove_glyph].
*/
func (self Instance) FontRemoveTexture(font_rid RID.Font, size Vector2i.XY, texture_index int) { //gd:TextServer.font_remove_texture
	class(self).FontRemoveTexture(gd.RID(font_rid), gd.Vector2i(size), gd.Int(texture_index))
}

/*
Sets font cache texture image data.
*/
func (self Instance) FontSetTextureImage(font_rid RID.Font, size Vector2i.XY, texture_index int, image [1]gdclass.Image) { //gd:TextServer.font_set_texture_image
	class(self).FontSetTextureImage(gd.RID(font_rid), gd.Vector2i(size), gd.Int(texture_index), image)
}

/*
Returns font cache texture image data.
*/
func (self Instance) FontGetTextureImage(font_rid RID.Font, size Vector2i.XY, texture_index int) [1]gdclass.Image { //gd:TextServer.font_get_texture_image
	return [1]gdclass.Image(class(self).FontGetTextureImage(gd.RID(font_rid), gd.Vector2i(size), gd.Int(texture_index)))
}

/*
Sets array containing glyph packing data.
*/
func (self Instance) FontSetTextureOffsets(font_rid RID.Font, size Vector2i.XY, texture_index int, offset []int32) { //gd:TextServer.font_set_texture_offsets
	class(self).FontSetTextureOffsets(gd.RID(font_rid), gd.Vector2i(size), gd.Int(texture_index), Packed.New(offset...))
}

/*
Returns array containing glyph packing data.
*/
func (self Instance) FontGetTextureOffsets(font_rid RID.Font, size Vector2i.XY, texture_index int) []int32 { //gd:TextServer.font_get_texture_offsets
	return []int32(slices.Collect(class(self).FontGetTextureOffsets(gd.RID(font_rid), gd.Vector2i(size), gd.Int(texture_index)).Values()))
}

/*
Returns list of rendered glyphs in the cache entry.
*/
func (self Instance) FontGetGlyphList(font_rid RID.Font, size Vector2i.XY) []int32 { //gd:TextServer.font_get_glyph_list
	return []int32(slices.Collect(class(self).FontGetGlyphList(gd.RID(font_rid), gd.Vector2i(size)).Values()))
}

/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method font_remove_texture] to remove them manually.
*/
func (self Instance) FontClearGlyphs(font_rid RID.Font, size Vector2i.XY) { //gd:TextServer.font_clear_glyphs
	class(self).FontClearGlyphs(gd.RID(font_rid), gd.Vector2i(size))
}

/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method font_remove_texture] to remove them manually.
*/
func (self Instance) FontRemoveGlyph(font_rid RID.Font, size Vector2i.XY, glyph int) { //gd:TextServer.font_remove_glyph
	class(self).FontRemoveGlyph(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph))
}

/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Instance) FontGetGlyphAdvance(font_rid RID.Font, size int, glyph int) Vector2.XY { //gd:TextServer.font_get_glyph_advance
	return Vector2.XY(class(self).FontGetGlyphAdvance(gd.RID(font_rid), gd.Int(size), gd.Int(glyph)))
}

/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Instance) FontSetGlyphAdvance(font_rid RID.Font, size int, glyph int, advance Vector2.XY) { //gd:TextServer.font_set_glyph_advance
	class(self).FontSetGlyphAdvance(gd.RID(font_rid), gd.Int(size), gd.Int(glyph), gd.Vector2(advance))
}

/*
Returns glyph offset from the baseline.
*/
func (self Instance) FontGetGlyphOffset(font_rid RID.Font, size Vector2i.XY, glyph int) Vector2.XY { //gd:TextServer.font_get_glyph_offset
	return Vector2.XY(class(self).FontGetGlyphOffset(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph)))
}

/*
Sets glyph offset from the baseline.
*/
func (self Instance) FontSetGlyphOffset(font_rid RID.Font, size Vector2i.XY, glyph int, offset Vector2.XY) { //gd:TextServer.font_set_glyph_offset
	class(self).FontSetGlyphOffset(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph), gd.Vector2(offset))
}

/*
Returns size of the glyph.
*/
func (self Instance) FontGetGlyphSize(font_rid RID.Font, size Vector2i.XY, glyph int) Vector2.XY { //gd:TextServer.font_get_glyph_size
	return Vector2.XY(class(self).FontGetGlyphSize(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph)))
}

/*
Sets size of the glyph.
*/
func (self Instance) FontSetGlyphSize(font_rid RID.Font, size Vector2i.XY, glyph int, gl_size Vector2.XY) { //gd:TextServer.font_set_glyph_size
	class(self).FontSetGlyphSize(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph), gd.Vector2(gl_size))
}

/*
Returns rectangle in the cache texture containing the glyph.
*/
func (self Instance) FontGetGlyphUvRect(font_rid RID.Font, size Vector2i.XY, glyph int) Rect2.PositionSize { //gd:TextServer.font_get_glyph_uv_rect
	return Rect2.PositionSize(class(self).FontGetGlyphUvRect(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph)))
}

/*
Sets rectangle in the cache texture containing the glyph.
*/
func (self Instance) FontSetGlyphUvRect(font_rid RID.Font, size Vector2i.XY, glyph int, uv_rect Rect2.PositionSize) { //gd:TextServer.font_set_glyph_uv_rect
	class(self).FontSetGlyphUvRect(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph), gd.Rect2(uv_rect))
}

/*
Returns index of the cache texture containing the glyph.
*/
func (self Instance) FontGetGlyphTextureIdx(font_rid RID.Font, size Vector2i.XY, glyph int) int { //gd:TextServer.font_get_glyph_texture_idx
	return int(int(class(self).FontGetGlyphTextureIdx(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph))))
}

/*
Sets index of the cache texture containing the glyph.
*/
func (self Instance) FontSetGlyphTextureIdx(font_rid RID.Font, size Vector2i.XY, glyph int, texture_idx int) { //gd:TextServer.font_set_glyph_texture_idx
	class(self).FontSetGlyphTextureIdx(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph), gd.Int(texture_idx))
}

/*
Returns resource ID of the cache texture containing the glyph.
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
func (self Instance) FontGetGlyphTextureRid(font_rid RID.Font, size Vector2i.XY, glyph int) RID.Texture { //gd:TextServer.font_get_glyph_texture_rid
	return RID.Texture(class(self).FontGetGlyphTextureRid(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph)))
}

/*
Returns size of the cache texture containing the glyph.
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
func (self Instance) FontGetGlyphTextureSize(font_rid RID.Font, size Vector2i.XY, glyph int) Vector2.XY { //gd:TextServer.font_get_glyph_texture_size
	return Vector2.XY(class(self).FontGetGlyphTextureSize(gd.RID(font_rid), gd.Vector2i(size), gd.Int(glyph)))
}

/*
Returns outline contours of the glyph as a [Dictionary] with the following contents:
[code]points[/code]         - [PackedVector3Array], containing outline points. [code]x[/code] and [code]y[/code] are point coordinates. [code]z[/code] is the type of the point, using the [enum ContourPointTag] values.
[code]contours[/code]       - [PackedInt32Array], containing indices the end points of each contour.
[code]orientation[/code]    - [bool], contour orientation. If [code]true[/code], clockwise contours must be filled.
*/
func (self Instance) FontGetGlyphContours(font RID.Font, size int, index int) map[string]interface{} { //gd:TextServer.font_get_glyph_contours
	return map[string]interface{}(gd.DictionaryAs[map[string]interface{}](class(self).FontGetGlyphContours(gd.RID(font), gd.Int(size), gd.Int(index))))
}

/*
Returns list of the kerning overrides.
*/
func (self Instance) FontGetKerningList(font_rid RID.Font, size int) []Vector2i.XY { //gd:TextServer.font_get_kerning_list
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).FontGetKerningList(gd.RID(font_rid), gd.Int(size)))))
}

/*
Removes all kerning overrides.
*/
func (self Instance) FontClearKerningMap(font_rid RID.Font, size int) { //gd:TextServer.font_clear_kerning_map
	class(self).FontClearKerningMap(gd.RID(font_rid), gd.Int(size))
}

/*
Removes kerning override for the pair of glyphs.
*/
func (self Instance) FontRemoveKerning(font_rid RID.Font, size int, glyph_pair Vector2i.XY) { //gd:TextServer.font_remove_kerning
	class(self).FontRemoveKerning(gd.RID(font_rid), gd.Int(size), gd.Vector2i(glyph_pair))
}

/*
Sets kerning for the pair of glyphs.
*/
func (self Instance) FontSetKerning(font_rid RID.Font, size int, glyph_pair Vector2i.XY, kerning Vector2.XY) { //gd:TextServer.font_set_kerning
	class(self).FontSetKerning(gd.RID(font_rid), gd.Int(size), gd.Vector2i(glyph_pair), gd.Vector2(kerning))
}

/*
Returns kerning for the pair of glyphs.
*/
func (self Instance) FontGetKerning(font_rid RID.Font, size int, glyph_pair Vector2i.XY) Vector2.XY { //gd:TextServer.font_get_kerning
	return Vector2.XY(class(self).FontGetKerning(gd.RID(font_rid), gd.Int(size), gd.Vector2i(glyph_pair)))
}

/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector]. See [method font_get_char_from_glyph_index].
*/
func (self Instance) FontGetGlyphIndex(font_rid RID.Font, size int, char int, variation_selector int) int { //gd:TextServer.font_get_glyph_index
	return int(int(class(self).FontGetGlyphIndex(gd.RID(font_rid), gd.Int(size), gd.Int(char), gd.Int(variation_selector))))
}

/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method font_get_glyph_index].
*/
func (self Instance) FontGetCharFromGlyphIndex(font_rid RID.Font, size int, glyph_index int) int { //gd:TextServer.font_get_char_from_glyph_index
	return int(int(class(self).FontGetCharFromGlyphIndex(gd.RID(font_rid), gd.Int(size), gd.Int(glyph_index))))
}

/*
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
func (self Instance) FontHasChar(font_rid RID.Font, char int) bool { //gd:TextServer.font_has_char
	return bool(class(self).FontHasChar(gd.RID(font_rid), gd.Int(char)))
}

/*
Returns a string containing all the characters available in the font.
*/
func (self Instance) FontGetSupportedChars(font_rid RID.Font) string { //gd:TextServer.font_get_supported_chars
	return string(class(self).FontGetSupportedChars(gd.RID(font_rid)).String())
}

/*
Renders the range of characters to the font cache texture.
*/
func (self Instance) FontRenderRange(font_rid RID.Font, size Vector2i.XY, start int, end int) { //gd:TextServer.font_render_range
	class(self).FontRenderRange(gd.RID(font_rid), gd.Vector2i(size), gd.Int(start), gd.Int(end))
}

/*
Renders specified glyph to the font cache texture.
*/
func (self Instance) FontRenderGlyph(font_rid RID.Font, size Vector2i.XY, index int) { //gd:TextServer.font_render_glyph
	class(self).FontRenderGlyph(gd.RID(font_rid), gd.Vector2i(size), gd.Int(index))
}

/*
Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
[b]Note:[/b] Glyph index is specific to the font, use glyphs indices returned by [method shaped_text_get_glyphs] or [method font_get_glyph_index].
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
func (self Instance) FontDrawGlyph(font_rid RID.Font, canvas RID.Canvas, size int, pos Vector2.XY, index int) { //gd:TextServer.font_draw_glyph
	class(self).FontDrawGlyph(gd.RID(font_rid), gd.RID(canvas), gd.Int(size), gd.Vector2(pos), gd.Int(index), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
[b]Note:[/b] Glyph index is specific to the font, use glyphs indices returned by [method shaped_text_get_glyphs] or [method font_get_glyph_index].
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
func (self Instance) FontDrawGlyphOutline(font_rid RID.Font, canvas RID.Canvas, size int, outline_size int, pos Vector2.XY, index int) { //gd:TextServer.font_draw_glyph_outline
	class(self).FontDrawGlyphOutline(gd.RID(font_rid), gd.RID(canvas), gd.Int(size), gd.Int(outline_size), gd.Vector2(pos), gd.Int(index), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
func (self Instance) FontIsLanguageSupported(font_rid RID.Font, language string) bool { //gd:TextServer.font_is_language_supported
	return bool(class(self).FontIsLanguageSupported(gd.RID(font_rid), String.New(language)))
}

/*
Adds override for [method font_is_language_supported].
*/
func (self Instance) FontSetLanguageSupportOverride(font_rid RID.Font, language string, supported bool) { //gd:TextServer.font_set_language_support_override
	class(self).FontSetLanguageSupportOverride(gd.RID(font_rid), String.New(language), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
func (self Instance) FontGetLanguageSupportOverride(font_rid RID.Font, language string) bool { //gd:TextServer.font_get_language_support_override
	return bool(class(self).FontGetLanguageSupportOverride(gd.RID(font_rid), String.New(language)))
}

/*
Remove language support override.
*/
func (self Instance) FontRemoveLanguageSupportOverride(font_rid RID.Font, language string) { //gd:TextServer.font_remove_language_support_override
	class(self).FontRemoveLanguageSupportOverride(gd.RID(font_rid), String.New(language))
}

/*
Returns list of language support overrides.
*/
func (self Instance) FontGetLanguageSupportOverrides(font_rid RID.Font) []string { //gd:TextServer.font_get_language_support_overrides
	return []string(class(self).FontGetLanguageSupportOverrides(gd.RID(font_rid)).Strings())
}

/*
Returns [code]true[/code], if font supports given script (ISO 15924 code).
*/
func (self Instance) FontIsScriptSupported(font_rid RID.Font, script string) bool { //gd:TextServer.font_is_script_supported
	return bool(class(self).FontIsScriptSupported(gd.RID(font_rid), String.New(script)))
}

/*
Adds override for [method font_is_script_supported].
*/
func (self Instance) FontSetScriptSupportOverride(font_rid RID.Font, script string, supported bool) { //gd:TextServer.font_set_script_support_override
	class(self).FontSetScriptSupportOverride(gd.RID(font_rid), String.New(script), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
func (self Instance) FontGetScriptSupportOverride(font_rid RID.Font, script string) bool { //gd:TextServer.font_get_script_support_override
	return bool(class(self).FontGetScriptSupportOverride(gd.RID(font_rid), String.New(script)))
}

/*
Removes script support override.
*/
func (self Instance) FontRemoveScriptSupportOverride(font_rid RID.Font, script string) { //gd:TextServer.font_remove_script_support_override
	class(self).FontRemoveScriptSupportOverride(gd.RID(font_rid), String.New(script))
}

/*
Returns list of script support overrides.
*/
func (self Instance) FontGetScriptSupportOverrides(font_rid RID.Font) []string { //gd:TextServer.font_get_script_support_overrides
	return []string(class(self).FontGetScriptSupportOverrides(gd.RID(font_rid)).Strings())
}

/*
Sets font OpenType feature set override.
*/
func (self Instance) FontSetOpentypeFeatureOverrides(font_rid RID.Font, overrides map[string]string) { //gd:TextServer.font_set_opentype_feature_overrides
	class(self).FontSetOpentypeFeatureOverrides(gd.RID(font_rid), gd.DictionaryFromMap(overrides))
}

/*
Returns font OpenType feature set override.
*/
func (self Instance) FontGetOpentypeFeatureOverrides(font_rid RID.Font) map[string]string { //gd:TextServer.font_get_opentype_feature_overrides
	return map[string]string(gd.DictionaryAs[map[string]string](class(self).FontGetOpentypeFeatureOverrides(gd.RID(font_rid))))
}

/*
Returns the dictionary of the supported OpenType features.
*/
func (self Instance) FontSupportedFeatureList(font_rid RID.Font) map[string]string { //gd:TextServer.font_supported_feature_list
	return map[string]string(gd.DictionaryAs[map[string]string](class(self).FontSupportedFeatureList(gd.RID(font_rid))))
}

/*
Returns the dictionary of the supported OpenType variation coordinates.
*/
func (self Instance) FontSupportedVariationList(font_rid RID.Font) map[string]string { //gd:TextServer.font_supported_variation_list
	return map[string]string(gd.DictionaryAs[map[string]string](class(self).FontSupportedVariationList(gd.RID(font_rid))))
}

/*
Returns the font oversampling factor, shared by all fonts in the TextServer.
*/
func (self Instance) FontGetGlobalOversampling() Float.X { //gd:TextServer.font_get_global_oversampling
	return Float.X(Float.X(class(self).FontGetGlobalOversampling()))
}

/*
Sets oversampling factor, shared by all font in the TextServer.
[b]Note:[/b] This value can be automatically changed by display server.
*/
func (self Instance) FontSetGlobalOversampling(oversampling Float.X) { //gd:TextServer.font_set_global_oversampling
	class(self).FontSetGlobalOversampling(gd.Float(oversampling))
}

/*
Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
*/
func (self Instance) GetHexCodeBoxSize(size int, index int) Vector2.XY { //gd:TextServer.get_hex_code_box_size
	return Vector2.XY(class(self).GetHexCodeBoxSize(gd.Int(size), gd.Int(index)))
}

/*
Draws box displaying character hexadecimal code. Used for replacing missing characters.
*/
func (self Instance) DrawHexCodeBox(canvas RID.Canvas, size int, pos Vector2.XY, index int, color Color.RGBA) { //gd:TextServer.draw_hex_code_box
	class(self).DrawHexCodeBox(gd.RID(canvas), gd.Int(size), gd.Vector2(pos), gd.Int(index), gd.Color(color))
}

/*
Creates a new buffer for complex text layout, with the given [param direction] and [param orientation]. To free the resulting buffer, use [method free_rid] method.
[b]Note:[/b] Direction is ignored if server does not support [constant FEATURE_BIDI_LAYOUT] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] Orientation is ignored if server does not support [constant FEATURE_VERTICAL_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
func (self Instance) CreateShapedText() RID.TextBuffer { //gd:TextServer.create_shaped_text
	return RID.TextBuffer(class(self).CreateShapedText(0, 0))
}

/*
Clears text buffer (removes text and inline objects).
*/
func (self Instance) ShapedTextClear(rid RID.TextBuffer) { //gd:TextServer.shaped_text_clear
	class(self).ShapedTextClear(gd.RID(rid))
}

/*
Sets desired text direction. If set to [constant DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
[b]Note:[/b] Direction is ignored if server does not support [constant FEATURE_BIDI_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
func (self Instance) ShapedTextSetDirection(shaped RID.TextBuffer) { //gd:TextServer.shaped_text_set_direction
	class(self).ShapedTextSetDirection(gd.RID(shaped), 0)
}

/*
Returns direction of the text.
*/
func (self Instance) ShapedTextGetDirection(shaped RID.TextBuffer) gdclass.TextServerDirection { //gd:TextServer.shaped_text_get_direction
	return gdclass.TextServerDirection(class(self).ShapedTextGetDirection(gd.RID(shaped)))
}

/*
Returns direction of the text, inferred by the BiDi algorithm.
*/
func (self Instance) ShapedTextGetInferredDirection(shaped RID.TextBuffer) gdclass.TextServerDirection { //gd:TextServer.shaped_text_get_inferred_direction
	return gdclass.TextServerDirection(class(self).ShapedTextGetInferredDirection(gd.RID(shaped)))
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
func (self Instance) ShapedTextSetBidiOverride(shaped RID.TextBuffer, override []any) { //gd:TextServer.shaped_text_set_bidi_override
	class(self).ShapedTextSetBidiOverride(gd.RID(shaped), gd.EngineArrayFromSlice(override))
}

/*
Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (self Instance) ShapedTextSetCustomPunctuation(shaped RID.TextBuffer, punct string) { //gd:TextServer.shaped_text_set_custom_punctuation
	class(self).ShapedTextSetCustomPunctuation(gd.RID(shaped), String.New(punct))
}

/*
Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
func (self Instance) ShapedTextGetCustomPunctuation(shaped RID.TextBuffer) string { //gd:TextServer.shaped_text_get_custom_punctuation
	return string(class(self).ShapedTextGetCustomPunctuation(gd.RID(shaped)).String())
}

/*
Sets ellipsis character used for text clipping.
*/
func (self Instance) ShapedTextSetCustomEllipsis(shaped RID.TextBuffer, char int) { //gd:TextServer.shaped_text_set_custom_ellipsis
	class(self).ShapedTextSetCustomEllipsis(gd.RID(shaped), gd.Int(char))
}

/*
Returns ellipsis character used for text clipping.
*/
func (self Instance) ShapedTextGetCustomEllipsis(shaped RID.TextBuffer) int { //gd:TextServer.shaped_text_get_custom_ellipsis
	return int(int(class(self).ShapedTextGetCustomEllipsis(gd.RID(shaped))))
}

/*
Sets desired text orientation.
[b]Note:[/b] Orientation is ignored if server does not support [constant FEATURE_VERTICAL_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
func (self Instance) ShapedTextSetOrientation(shaped RID.TextBuffer) { //gd:TextServer.shaped_text_set_orientation
	class(self).ShapedTextSetOrientation(gd.RID(shaped), 0)
}

/*
Returns text orientation.
*/
func (self Instance) ShapedTextGetOrientation(shaped RID.TextBuffer) gdclass.TextServerOrientation { //gd:TextServer.shaped_text_get_orientation
	return gdclass.TextServerOrientation(class(self).ShapedTextGetOrientation(gd.RID(shaped)))
}

/*
If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
*/
func (self Instance) ShapedTextSetPreserveInvalid(shaped RID.TextBuffer, enabled bool) { //gd:TextServer.shaped_text_set_preserve_invalid
	class(self).ShapedTextSetPreserveInvalid(gd.RID(shaped), enabled)
}

/*
Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
[b]Note:[/b] If set to [code]false[/code], nothing is displayed in place of invalid characters.
*/
func (self Instance) ShapedTextGetPreserveInvalid(shaped RID.TextBuffer) bool { //gd:TextServer.shaped_text_get_preserve_invalid
	return bool(class(self).ShapedTextGetPreserveInvalid(gd.RID(shaped)))
}

/*
If set to [code]true[/code] text buffer will display control characters.
*/
func (self Instance) ShapedTextSetPreserveControl(shaped RID.TextBuffer, enabled bool) { //gd:TextServer.shaped_text_set_preserve_control
	class(self).ShapedTextSetPreserveControl(gd.RID(shaped), enabled)
}

/*
Returns [code]true[/code] if text buffer is configured to display control characters.
*/
func (self Instance) ShapedTextGetPreserveControl(shaped RID.TextBuffer) bool { //gd:TextServer.shaped_text_get_preserve_control
	return bool(class(self).ShapedTextGetPreserveControl(gd.RID(shaped)))
}

/*
Sets extra spacing added between glyphs or lines in pixels.
*/
func (self Instance) ShapedTextSetSpacing(shaped RID.TextBuffer, spacing gdclass.TextServerSpacingType, value int) { //gd:TextServer.shaped_text_set_spacing
	class(self).ShapedTextSetSpacing(gd.RID(shaped), spacing, gd.Int(value))
}

/*
Returns extra spacing added between glyphs or lines in pixels.
*/
func (self Instance) ShapedTextGetSpacing(shaped RID.TextBuffer, spacing gdclass.TextServerSpacingType) int { //gd:TextServer.shaped_text_get_spacing
	return int(int(class(self).ShapedTextGetSpacing(gd.RID(shaped), spacing)))
}

/*
Adds text span and font to draw it to the text buffer.
*/
func (self Instance) ShapedTextAddString(shaped RID.TextBuffer, text string, fonts []RID.TextBuffer, size int) bool { //gd:TextServer.shaped_text_add_string
	return bool(class(self).ShapedTextAddString(gd.RID(shaped), String.New(text), gd.ArrayFromSlice[Array.Contains[gd.RID]](fonts), gd.Int(size), Dictionary.Nil, String.New(""), gd.NewVariant(gd.NewVariant(([1]any{}[0])))))
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (self Instance) ShapedTextAddObject(shaped RID.TextBuffer, key any, size Vector2.XY) bool { //gd:TextServer.shaped_text_add_object
	return bool(class(self).ShapedTextAddObject(gd.RID(shaped), gd.NewVariant(key), gd.Vector2(size), 5, gd.Int(1), gd.Float(0.0)))
}

/*
Sets new size and alignment of embedded object.
*/
func (self Instance) ShapedTextResizeObject(shaped RID.TextBuffer, key any, size Vector2.XY) bool { //gd:TextServer.shaped_text_resize_object
	return bool(class(self).ShapedTextResizeObject(gd.RID(shaped), gd.NewVariant(key), gd.Vector2(size), 5, gd.Float(0.0)))
}

/*
Returns number of text spans added using [method shaped_text_add_string] or [method shaped_text_add_object].
*/
func (self Instance) ShapedGetSpanCount(shaped RID.TextBuffer) int { //gd:TextServer.shaped_get_span_count
	return int(int(class(self).ShapedGetSpanCount(gd.RID(shaped))))
}

/*
Returns text span metadata.
*/
func (self Instance) ShapedGetSpanMeta(shaped RID.TextBuffer, index int) any { //gd:TextServer.shaped_get_span_meta
	return any(class(self).ShapedGetSpanMeta(gd.RID(shaped), gd.Int(index)).Interface())
}

/*
Changes text span font, font size, and OpenType features, without changing the text.
*/
func (self Instance) ShapedSetSpanUpdateFont(shaped RID.TextBuffer, index int, fonts [][]RID.Font, size int) { //gd:TextServer.shaped_set_span_update_font
	class(self).ShapedSetSpanUpdateFont(gd.RID(shaped), gd.Int(index), gd.ArrayFromSlice[Array.Contains[gd.RID]](fonts), gd.Int(size), Dictionary.Nil)
}

/*
Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
*/
func (self Instance) ShapedTextSubstr(shaped RID.TextBuffer, start int, length int) RID.TextBuffer { //gd:TextServer.shaped_text_substr
	return RID.TextBuffer(class(self).ShapedTextSubstr(gd.RID(shaped), gd.Int(start), gd.Int(length)))
}

/*
Returns the parent buffer from which the substring originates.
*/
func (self Instance) ShapedTextGetParent(shaped RID.TextBuffer) RID.TextBuffer { //gd:TextServer.shaped_text_get_parent
	return RID.TextBuffer(class(self).ShapedTextGetParent(gd.RID(shaped)))
}

/*
Adjusts text width to fit to specified width, returns new text width.
*/
func (self Instance) ShapedTextFitToWidth(shaped RID.TextBuffer, width Float.X) Float.X { //gd:TextServer.shaped_text_fit_to_width
	return Float.X(Float.X(class(self).ShapedTextFitToWidth(gd.RID(shaped), gd.Float(width), 3)))
}

/*
Aligns shaped text to the given tab-stops.
*/
func (self Instance) ShapedTextTabAlign(shaped RID.TextBuffer, tab_stops []float32) Float.X { //gd:TextServer.shaped_text_tab_align
	return Float.X(Float.X(class(self).ShapedTextTabAlign(gd.RID(shaped), Packed.New(tab_stops...))))
}

/*
Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
[b]Note:[/b] It is not necessary to call this function manually, buffer will be shaped automatically as soon as any of its output data is requested.
*/
func (self Instance) ShapedTextShape(shaped RID.TextBuffer) bool { //gd:TextServer.shaped_text_shape
	return bool(class(self).ShapedTextShape(gd.RID(shaped)))
}

/*
Returns [code]true[/code] if buffer is successfully shaped.
*/
func (self Instance) ShapedTextIsReady(shaped RID.TextBuffer) bool { //gd:TextServer.shaped_text_is_ready
	return bool(class(self).ShapedTextIsReady(gd.RID(shaped)))
}

/*
Returns [code]true[/code] if text buffer contains any visible characters.
*/
func (self Instance) ShapedTextHasVisibleChars(shaped RID.TextBuffer) bool { //gd:TextServer.shaped_text_has_visible_chars
	return bool(class(self).ShapedTextHasVisibleChars(gd.RID(shaped)))
}

/*
Returns an array of glyphs in the visual order.
*/
func (self Instance) ShapedTextGetGlyphs(shaped RID.TextBuffer) []map[int]struct {
	X float32
	Y float32
} { //gd:TextServer.shaped_text_get_glyphs
	return []map[int]struct {
		X float32
		Y float32
	}(gd.ArrayAs[[]map[int]struct {
		X float32
		Y float32
	}](gd.InternalArray(class(self).ShapedTextGetGlyphs(gd.RID(shaped)))))
}

/*
Returns text glyphs in the logical order.
*/
func (self Instance) ShapedTextSortLogical(shaped RID.TextBuffer) []map[int]int { //gd:TextServer.shaped_text_sort_logical
	return []map[int]int(gd.ArrayAs[[]map[int]int](gd.InternalArray(class(self).ShapedTextSortLogical(gd.RID(shaped)))))
}

/*
Returns number of glyphs in the buffer.
*/
func (self Instance) ShapedTextGetGlyphCount(shaped RID.TextBuffer) int { //gd:TextServer.shaped_text_get_glyph_count
	return int(int(class(self).ShapedTextGetGlyphCount(gd.RID(shaped))))
}

/*
Returns substring buffer character range in the parent buffer.
*/
func (self Instance) ShapedTextGetRange(shaped RID.TextBuffer) Vector2i.XY { //gd:TextServer.shaped_text_get_range
	return Vector2i.XY(class(self).ShapedTextGetRange(gd.RID(shaped)))
}

/*
Breaks text to the lines and columns. Returns character ranges for each segment.
*/
func (self Instance) ShapedTextGetLineBreaksAdv(shaped RID.TextBuffer, width []float32) []int32 { //gd:TextServer.shaped_text_get_line_breaks_adv
	return []int32(slices.Collect(class(self).ShapedTextGetLineBreaksAdv(gd.RID(shaped), Packed.New(width...), gd.Int(0), true, 3).Values()))
}

/*
Breaks text to the lines and returns character ranges for each line.
*/
func (self Instance) ShapedTextGetLineBreaks(shaped RID.TextBuffer, width Float.X) []int32 { //gd:TextServer.shaped_text_get_line_breaks
	return []int32(slices.Collect(class(self).ShapedTextGetLineBreaks(gd.RID(shaped), gd.Float(width), gd.Int(0), 3).Values()))
}

/*
Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum GraphemeFlag]).
*/
func (self Instance) ShapedTextGetWordBreaks(shaped RID.TextBuffer) []int32 { //gd:TextServer.shaped_text_get_word_breaks
	return []int32(slices.Collect(class(self).ShapedTextGetWordBreaks(gd.RID(shaped), 264, 4).Values()))
}

/*
Returns the position of the overrun trim.
*/
func (self Instance) ShapedTextGetTrimPos(shaped RID.TextBuffer) int { //gd:TextServer.shaped_text_get_trim_pos
	return int(int(class(self).ShapedTextGetTrimPos(gd.RID(shaped))))
}

/*
Returns position of the ellipsis.
*/
func (self Instance) ShapedTextGetEllipsisPos(shaped RID.TextBuffer) int { //gd:TextServer.shaped_text_get_ellipsis_pos
	return int(int(class(self).ShapedTextGetEllipsisPos(gd.RID(shaped))))
}

/*
Returns array of the glyphs in the ellipsis.
*/
func (self Instance) ShapedTextGetEllipsisGlyphs(shaped RID.TextBuffer) []map[int]struct {
	X float32
	Y float32
} { //gd:TextServer.shaped_text_get_ellipsis_glyphs
	return []map[int]struct {
		X float32
		Y float32
	}(gd.ArrayAs[[]map[int]struct {
		X float32
		Y float32
	}](gd.InternalArray(class(self).ShapedTextGetEllipsisGlyphs(gd.RID(shaped)))))
}

/*
Returns number of glyphs in the ellipsis.
*/
func (self Instance) ShapedTextGetEllipsisGlyphCount(shaped RID.TextBuffer) int { //gd:TextServer.shaped_text_get_ellipsis_glyph_count
	return int(int(class(self).ShapedTextGetEllipsisGlyphCount(gd.RID(shaped))))
}

/*
Trims text if it exceeds the given width.
*/
func (self Instance) ShapedTextOverrunTrimToWidth(shaped RID.TextBuffer) { //gd:TextServer.shaped_text_overrun_trim_to_width
	class(self).ShapedTextOverrunTrimToWidth(gd.RID(shaped), gd.Float(0), 0)
}

/*
Returns array of inline objects.
*/
func (self Instance) ShapedTextGetObjects(shaped RID.TextBuffer) []any { //gd:TextServer.shaped_text_get_objects
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).ShapedTextGetObjects(gd.RID(shaped)))))
}

/*
Returns bounding rectangle of the inline object.
*/
func (self Instance) ShapedTextGetObjectRect(shaped RID.TextBuffer, key any) Rect2.PositionSize { //gd:TextServer.shaped_text_get_object_rect
	return Rect2.PositionSize(class(self).ShapedTextGetObjectRect(gd.RID(shaped), gd.NewVariant(key)))
}

/*
Returns the character range of the inline object.
*/
func (self Instance) ShapedTextGetObjectRange(shaped RID.TextBuffer, key any) Vector2i.XY { //gd:TextServer.shaped_text_get_object_range
	return Vector2i.XY(class(self).ShapedTextGetObjectRange(gd.RID(shaped), gd.NewVariant(key)))
}

/*
Returns the glyph index of the inline object.
*/
func (self Instance) ShapedTextGetObjectGlyph(shaped RID.TextBuffer, key any) int { //gd:TextServer.shaped_text_get_object_glyph
	return int(int(class(self).ShapedTextGetObjectGlyph(gd.RID(shaped), gd.NewVariant(key))))
}

/*
Returns size of the text.
*/
func (self Instance) ShapedTextGetSize(shaped RID.TextBuffer) Vector2.XY { //gd:TextServer.shaped_text_get_size
	return Vector2.XY(class(self).ShapedTextGetSize(gd.RID(shaped)))
}

/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
[b]Note:[/b] Overall ascent can be higher than font ascent, if some glyphs are displaced from the baseline.
*/
func (self Instance) ShapedTextGetAscent(shaped RID.TextBuffer) Float.X { //gd:TextServer.shaped_text_get_ascent
	return Float.X(Float.X(class(self).ShapedTextGetAscent(gd.RID(shaped))))
}

/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
[b]Note:[/b] Overall descent can be higher than font descent, if some glyphs are displaced from the baseline.
*/
func (self Instance) ShapedTextGetDescent(shaped RID.TextBuffer) Float.X { //gd:TextServer.shaped_text_get_descent
	return Float.X(Float.X(class(self).ShapedTextGetDescent(gd.RID(shaped))))
}

/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
func (self Instance) ShapedTextGetWidth(shaped RID.TextBuffer) Float.X { //gd:TextServer.shaped_text_get_width
	return Float.X(Float.X(class(self).ShapedTextGetWidth(gd.RID(shaped))))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Instance) ShapedTextGetUnderlinePosition(shaped RID.TextBuffer) Float.X { //gd:TextServer.shaped_text_get_underline_position
	return Float.X(Float.X(class(self).ShapedTextGetUnderlinePosition(gd.RID(shaped))))
}

/*
Returns thickness of the underline.
*/
func (self Instance) ShapedTextGetUnderlineThickness(shaped RID.TextBuffer) Float.X { //gd:TextServer.shaped_text_get_underline_thickness
	return Float.X(Float.X(class(self).ShapedTextGetUnderlineThickness(gd.RID(shaped))))
}

/*
Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
*/
func (self Instance) ShapedTextGetCarets(shaped RID.TextBuffer, position int) map[int]struct {
	X float32
	Y float32
} { //gd:TextServer.shaped_text_get_carets
	return map[int]struct {
		X float32
		Y float32
	}(gd.DictionaryAs[map[int]struct {
		X float32
		Y float32
	}](class(self).ShapedTextGetCarets(gd.RID(shaped), gd.Int(position))))
}

/*
Returns selection rectangles for the specified character range.
*/
func (self Instance) ShapedTextGetSelection(shaped RID.TextBuffer, start int, end int) []Vector2.XY { //gd:TextServer.shaped_text_get_selection
	return []Vector2.XY(slices.Collect(class(self).ShapedTextGetSelection(gd.RID(shaped), gd.Int(start), gd.Int(end)).Values()))
}

/*
Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
*/
func (self Instance) ShapedTextHitTestGrapheme(shaped RID.TextBuffer, coords Float.X) int { //gd:TextServer.shaped_text_hit_test_grapheme
	return int(int(class(self).ShapedTextHitTestGrapheme(gd.RID(shaped), gd.Float(coords))))
}

/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
func (self Instance) ShapedTextHitTestPosition(shaped RID.TextBuffer, coords Float.X) int { //gd:TextServer.shaped_text_hit_test_position
	return int(int(class(self).ShapedTextHitTestPosition(gd.RID(shaped), gd.Float(coords))))
}

/*
Returns composite character's bounds as offsets from the start of the line.
*/
func (self Instance) ShapedTextGetGraphemeBounds(shaped RID.TextBuffer, pos int) Vector2.XY { //gd:TextServer.shaped_text_get_grapheme_bounds
	return Vector2.XY(class(self).ShapedTextGetGraphemeBounds(gd.RID(shaped), gd.Int(pos)))
}

/*
Returns grapheme end position closest to the [param pos].
*/
func (self Instance) ShapedTextNextGraphemePos(shaped RID.TextBuffer, pos int) int { //gd:TextServer.shaped_text_next_grapheme_pos
	return int(int(class(self).ShapedTextNextGraphemePos(gd.RID(shaped), gd.Int(pos))))
}

/*
Returns grapheme start position closest to the [param pos].
*/
func (self Instance) ShapedTextPrevGraphemePos(shaped RID.TextBuffer, pos int) int { //gd:TextServer.shaped_text_prev_grapheme_pos
	return int(int(class(self).ShapedTextPrevGraphemePos(gd.RID(shaped), gd.Int(pos))))
}

/*
Returns array of the composite character boundaries.
*/
func (self Instance) ShapedTextGetCharacterBreaks(shaped RID.TextBuffer) []int32 { //gd:TextServer.shaped_text_get_character_breaks
	return []int32(slices.Collect(class(self).ShapedTextGetCharacterBreaks(gd.RID(shaped)).Values()))
}

/*
Returns composite character end position closest to the [param pos].
*/
func (self Instance) ShapedTextNextCharacterPos(shaped RID.TextBuffer, pos int) int { //gd:TextServer.shaped_text_next_character_pos
	return int(int(class(self).ShapedTextNextCharacterPos(gd.RID(shaped), gd.Int(pos))))
}

/*
Returns composite character start position closest to the [param pos].
*/
func (self Instance) ShapedTextPrevCharacterPos(shaped RID.TextBuffer, pos int) int { //gd:TextServer.shaped_text_prev_character_pos
	return int(int(class(self).ShapedTextPrevCharacterPos(gd.RID(shaped), gd.Int(pos))))
}

/*
Returns composite character position closest to the [param pos].
*/
func (self Instance) ShapedTextClosestCharacterPos(shaped RID.TextBuffer, pos int) int { //gd:TextServer.shaped_text_closest_character_pos
	return int(int(class(self).ShapedTextClosestCharacterPos(gd.RID(shaped), gd.Int(pos))))
}

/*
Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (self Instance) ShapedTextDraw(shaped RID.TextBuffer, canvas RID.Canvas, pos Vector2.XY) { //gd:TextServer.shaped_text_draw
	class(self).ShapedTextDraw(gd.RID(shaped), gd.RID(canvas), gd.Vector2(pos), gd.Float(-1), gd.Float(-1), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
func (self Instance) ShapedTextDrawOutline(shaped RID.TextBuffer, canvas RID.Canvas, pos Vector2.XY) { //gd:TextServer.shaped_text_draw_outline
	class(self).ShapedTextDrawOutline(gd.RID(shaped), gd.RID(canvas), gd.Vector2(pos), gd.Float(-1), gd.Float(-1), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Returns dominant direction of in the range of text.
*/
func (self Instance) ShapedTextGetDominantDirectionInRange(shaped RID.TextBuffer, start int, end int) gdclass.TextServerDirection { //gd:TextServer.shaped_text_get_dominant_direction_in_range
	return gdclass.TextServerDirection(class(self).ShapedTextGetDominantDirectionInRange(gd.RID(shaped), gd.Int(start), gd.Int(end)))
}

/*
Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
If [param language] is omitted, the active locale will be used.
*/
func (self Instance) FormatNumber(number string) string { //gd:TextServer.format_number
	return string(class(self).FormatNumber(String.New(number), String.New("")).String())
}

/*
Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
*/
func (self Instance) ParseNumber(number string) string { //gd:TextServer.parse_number
	return string(class(self).ParseNumber(String.New(number), String.New("")).String())
}

/*
Returns percent sign used in the [param language].
*/
func (self Instance) PercentSign() string { //gd:TextServer.percent_sign
	return string(class(self).PercentSign(String.New("")).String())
}

/*
Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
When [param chars_per_line] is greater than zero, line break boundaries are returned instead.
[codeblock]
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_word_breaks("The Godot Engine, 4")) # Prints [0, 3, 4, 9, 10, 16, 18, 19], which corresponds to the following substrings: "The", "Godot", "Engine", "4"
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 5)) # Prints [0, 3, 4, 9, 10, 15, 15, 19], which corresponds to the following substrings: "The", "Godot", "Engin", "e, 4"
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 10)) # Prints [0, 9, 10, 19], which corresponds to the following substrings: "The Godot", "Engine, 4"
[/codeblock]
*/
func (self Instance) StringGetWordBreaks(s string) []int32 { //gd:TextServer.string_get_word_breaks
	return []int32(slices.Collect(class(self).StringGetWordBreaks(String.New(s), String.New(""), gd.Int(0)).Values()))
}

/*
Returns array of the composite character boundaries.
[codeblock]
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_word_breaks("Test ❤️‍🔥 Test")) # Prints [1, 2, 3, 4, 5, 9, 10, 11, 12, 13, 14]
[/codeblock]
*/
func (self Instance) StringGetCharacterBreaks(s string) []int32 { //gd:TextServer.string_get_character_breaks
	return []int32(slices.Collect(class(self).StringGetCharacterBreaks(String.New(s), String.New("")).Values()))
}

/*
Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
[b]Note:[/b] This method doesn't detect invisible characters, for spoof detection use it in combination with [method spoof_check].
[b]Note:[/b] Always returns [code]-1[/code] if the server does not support the [constant FEATURE_UNICODE_SECURITY] feature.
*/
func (self Instance) IsConfusable(s string, dict []string) int { //gd:TextServer.is_confusable
	return int(int(class(self).IsConfusable(String.New(s), Packed.MakeStrings(dict...))))
}

/*
Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
[b]Note:[/b] Always returns [code]false[/code] if the server does not support the [constant FEATURE_UNICODE_SECURITY] feature.
*/
func (self Instance) SpoofCheck(s string) bool { //gd:TextServer.spoof_check
	return bool(class(self).SpoofCheck(String.New(s)))
}

/*
Strips diacritics from the string.
[b]Note:[/b] The result may be longer or shorter than the original.
*/
func (self Instance) StripDiacritics(s string) string { //gd:TextServer.strip_diacritics
	return string(class(self).StripDiacritics(String.New(s)).String())
}

/*
Returns [code]true[/code] if [param string] is a valid identifier.
If the text server supports the [constant FEATURE_UNICODE_IDENTIFIERS] feature, a valid identifier must:
- Conform to normalization form C.
- Begin with a Unicode character of class XID_Start or [code]"_"[/code].
- May contain Unicode characters of class XID_Continue in the other positions.
- Use UAX #31 recommended scripts only (mixed scripts are allowed).
If the [constant FEATURE_UNICODE_IDENTIFIERS] feature is not supported, a valid identifier must:
- Begin with a Unicode character of class XID_Start or [code]"_"[/code].
- May contain Unicode characters of class XID_Continue in the other positions.
*/
func (self Instance) IsValidIdentifier(s string) bool { //gd:TextServer.is_valid_identifier
	return bool(class(self).IsValidIdentifier(String.New(s)))
}

/*
Returns [code]true[/code] if the given code point is a valid letter, i.e. it belongs to the Unicode category "L".
*/
func (self Instance) IsValidLetter(unicode int) bool { //gd:TextServer.is_valid_letter
	return bool(class(self).IsValidLetter(gd.Int(unicode)))
}

/*
Returns the string converted to uppercase.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
func (self Instance) StringToUpper(s string) string { //gd:TextServer.string_to_upper
	return string(class(self).StringToUpper(String.New(s), String.New("")).String())
}

/*
Returns the string converted to lowercase.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
func (self Instance) StringToLower(s string) string { //gd:TextServer.string_to_lower
	return string(class(self).StringToLower(String.New(s), String.New("")).String())
}

/*
Returns the string converted to title case.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
func (self Instance) StringToTitle(s string) string { //gd:TextServer.string_to_title
	return string(class(self).StringToTitle(String.New(s), String.New("")).String())
}

/*
Default implementation of the BiDi algorithm override function. See [enum StructuredTextParser] for more info.
*/
func (self Instance) ParseStructuredText(parser_type gdclass.TextServerStructuredTextParser, args []any, text string) []Vector3i.XYZ { //gd:TextServer.parse_structured_text
	return []Vector3i.XYZ(gd.ArrayAs[[]Vector3i.XYZ](gd.InternalArray(class(self).ParseStructuredText(parser_type, gd.EngineArrayFromSlice(args), String.New(text)))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextServer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextServer"))
	casted := Instance{*(*gdclass.TextServer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns [code]true[/code] if the server supports a feature.
*/
//go:nosplit
func (self class) HasFeature(feature gdclass.TextServerFeature) bool { //gd:TextServer.has_feature
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_has_feature, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the server interface.
*/
//go:nosplit
func (self class) GetName() String.Readable { //gd:TextServer.get_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns text server features, see [enum Feature].
*/
//go:nosplit
func (self class) GetFeatures() gd.Int { //gd:TextServer.get_features
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_get_features, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads optional TextServer database (e.g. ICU break iterators and dictionaries).
[b]Note:[/b] This function should be called before any other TextServer functions used, otherwise it won't have any effect.
*/
//go:nosplit
func (self class) LoadSupportData(filename String.Readable) bool { //gd:TextServer.load_support_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(filename)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_load_support_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns default TextServer database (e.g. ICU break iterators and dictionaries) filename.
*/
//go:nosplit
func (self class) GetSupportDataFilename() String.Readable { //gd:TextServer.get_support_data_filename
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_get_support_data_filename, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns TextServer database (e.g. ICU break iterators and dictionaries) description.
*/
//go:nosplit
func (self class) GetSupportDataInfo() String.Readable { //gd:TextServer.get_support_data_info
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_get_support_data_info, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Saves optional TextServer database (e.g. ICU break iterators and dictionaries) to the file.
[b]Note:[/b] This function is used by during project export, to include TextServer database.
*/
//go:nosplit
func (self class) SaveSupportData(filename String.Readable) bool { //gd:TextServer.save_support_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(filename)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_save_support_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if locale is right-to-left.
*/
//go:nosplit
func (self class) IsLocaleRightToLeft(locale String.Readable) bool { //gd:TextServer.is_locale_right_to_left
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(locale)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_is_locale_right_to_left, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Converts readable feature, variation, script, or language name to OpenType tag.
*/
//go:nosplit
func (self class) NameToTag(name String.Readable) gd.Int { //gd:TextServer.name_to_tag
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_name_to_tag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Converts OpenType tag to readable feature, variation, script, or language name.
*/
//go:nosplit
func (self class) TagToName(tag gd.Int) String.Readable { //gd:TextServer.tag_to_name
	var frame = callframe.New()
	callframe.Arg(frame, tag)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_tag_to_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [param rid] is valid resource owned by this text server.
*/
//go:nosplit
func (self class) Has(rid gd.RID) bool { //gd:TextServer.has
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_has, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Frees an object created by this [TextServer].
*/
//go:nosplit
func (self class) FreeRid(rid gd.RID) { //gd:TextServer.free_rid
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_free_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a new, empty font cache entry resource. To free the resulting resource, use the [method free_rid] method.
*/
//go:nosplit
func (self class) CreateFont() gd.RID { //gd:TextServer.create_font
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_create_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new variation existing font which is reusing the same glyph cache and font data. To free the resulting resource, use the [method free_rid] method.
*/
//go:nosplit
func (self class) CreateFontLinkedVariation(font_rid gd.RID) gd.RID { //gd:TextServer.create_font_linked_variation
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_create_font_linked_variation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets font source data, e.g contents of the dynamic font source file.
*/
//go:nosplit
func (self class) FontSetData(font_rid gd.RID, data Packed.Bytes) { //gd:TextServer.font_set_data
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) FontSetFaceIndex(font_rid gd.RID, face_index gd.Int) { //gd:TextServer.font_set_face_index
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, face_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_face_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) FontGetFaceIndex(font_rid gd.RID) gd.Int { //gd:TextServer.font_get_face_index
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_face_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of faces in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) FontGetFaceCount(font_rid gd.RID) gd.Int { //gd:TextServer.font_get_face_count
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_face_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the font style flags, see [enum FontStyle].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], [method font_set_embolden], or [method font_set_transform] instead.
*/
//go:nosplit
func (self class) FontSetStyle(font_rid gd.RID, style gdclass.TextServerFontStyle) { //gd:TextServer.font_set_style
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, style)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_style, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font style flags, see [enum FontStyle].
*/
//go:nosplit
func (self class) FontGetStyle(font_rid gd.RID) gdclass.TextServerFontStyle { //gd:TextServer.font_get_style
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gdclass.TextServerFontStyle](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_style, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the font family name.
*/
//go:nosplit
func (self class) FontSetName(font_rid gd.RID, name String.Readable) { //gd:TextServer.font_set_name
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font family name.
*/
//go:nosplit
func (self class) FontGetName(font_rid gd.RID) String.Readable { //gd:TextServer.font_get_name
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [Dictionary] with OpenType font name strings (localized font names, version, description, license information, sample text, etc.).
*/
//go:nosplit
func (self class) FontGetOtNameStrings(font_rid gd.RID) Dictionary.Any { //gd:TextServer.font_get_ot_name_strings
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_ot_name_strings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the font style name.
*/
//go:nosplit
func (self class) FontSetStyleName(font_rid gd.RID, name String.Readable) { //gd:TextServer.font_set_style_name
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_style_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font style name.
*/
//go:nosplit
func (self class) FontGetStyleName(font_rid gd.RID) String.Readable { //gd:TextServer.font_get_style_name
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_style_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], or [method font_set_embolden] instead.
*/
//go:nosplit
func (self class) FontSetWeight(font_rid gd.RID, weight gd.Int) { //gd:TextServer.font_set_weight
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, weight)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_weight, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns weight (boldness) of the font. A value in the [code]100...999[/code] range, normal font weight is [code]400[/code], bold font weight is [code]700[/code].
*/
//go:nosplit
func (self class) FontGetWeight(font_rid gd.RID) gd.Int { //gd:TextServer.font_get_weight
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_weight, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
[b]Note:[/b] This value is used for font matching only and will not affect font rendering. Use [method font_set_face_index], [method font_set_variation_coordinates], or [method font_set_transform] instead.
*/
//go:nosplit
func (self class) FontSetStretch(font_rid gd.RID, weight gd.Int) { //gd:TextServer.font_set_stretch
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, weight)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_stretch, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font stretch amount, compared to a normal width. A percentage value between [code]50%[/code] and [code]200%[/code].
*/
//go:nosplit
func (self class) FontGetStretch(font_rid gd.RID) gd.Int { //gd:TextServer.font_get_stretch
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_stretch, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets font anti-aliasing mode.
*/
//go:nosplit
func (self class) FontSetAntialiasing(font_rid gd.RID, antialiasing gdclass.TextServerFontAntialiasing) { //gd:TextServer.font_set_antialiasing
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, antialiasing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font anti-aliasing mode.
*/
//go:nosplit
func (self class) FontGetAntialiasing(font_rid gd.RID) gdclass.TextServerFontAntialiasing { //gd:TextServer.font_get_antialiasing
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gdclass.TextServerFontAntialiasing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], embedded font bitmap loading is disabled (bitmap-only and color fonts ignore this property).
*/
//go:nosplit
func (self class) FontSetDisableEmbeddedBitmaps(font_rid gd.RID, disable_embedded_bitmaps bool) { //gd:TextServer.font_set_disable_embedded_bitmaps
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, disable_embedded_bitmaps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the font's embedded bitmap loading is disabled.
*/
//go:nosplit
func (self class) FontGetDisableEmbeddedBitmaps(font_rid gd.RID) bool { //gd:TextServer.font_get_disable_embedded_bitmaps
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code] font texture mipmap generation is enabled.
*/
//go:nosplit
func (self class) FontSetGenerateMipmaps(font_rid gd.RID, generate_mipmaps bool) { //gd:TextServer.font_set_generate_mipmaps
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, generate_mipmaps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if font texture mipmap generation is enabled.
*/
//go:nosplit
func (self class) FontGetGenerateMipmaps(font_rid gd.RID) bool { //gd:TextServer.font_get_generate_mipmaps
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data. MSDF rendering allows displaying the font at any scaling factor without blurriness, and without incurring a CPU cost when the font size changes (since the font no longer needs to be rasterized on the CPU). As a downside, font hinting is not available with MSDF. The lack of font hinting may result in less crisp and less readable fonts at small sizes.
[b]Note:[/b] MSDF font rendering does not render glyphs with overlapping shapes correctly. Overlapping shapes are not valid per the OpenType standard, but are still commonly found in many font files, especially those converted by Google Fonts. To avoid issues with overlapping glyphs, consider downloading the font file directly from the type foundry instead of relying on Google Fonts.
*/
//go:nosplit
func (self class) FontSetMultichannelSignedDistanceField(font_rid gd.RID, msdf bool) { //gd:TextServer.font_set_multichannel_signed_distance_field
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, msdf)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if glyphs of all sizes are rendered using single multichannel signed distance field generated from the dynamic font vector data.
*/
//go:nosplit
func (self class) FontIsMultichannelSignedDistanceField(font_rid gd.RID) bool { //gd:TextServer.font_is_multichannel_signed_distance_field
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_is_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the width of the range around the shape between the minimum and maximum representable signed distance.
*/
//go:nosplit
func (self class) FontSetMsdfPixelRange(font_rid gd.RID, msdf_pixel_range gd.Int) { //gd:TextServer.font_set_msdf_pixel_range
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, msdf_pixel_range)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the width of the range around the shape between the minimum and maximum representable signed distance.
*/
//go:nosplit
func (self class) FontGetMsdfPixelRange(font_rid gd.RID) gd.Int { //gd:TextServer.font_get_msdf_pixel_range
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets source font size used to generate MSDF textures.
*/
//go:nosplit
func (self class) FontSetMsdfSize(font_rid gd.RID, msdf_size gd.Int) { //gd:TextServer.font_set_msdf_size
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, msdf_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_msdf_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns source font size used to generate MSDF textures.
*/
//go:nosplit
func (self class) FontGetMsdfSize(font_rid gd.RID) gd.Int { //gd:TextServer.font_get_msdf_size
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_msdf_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets bitmap font fixed size. If set to value greater than zero, same cache entry will be used for all font sizes.
*/
//go:nosplit
func (self class) FontSetFixedSize(font_rid gd.RID, fixed_size gd.Int) { //gd:TextServer.font_set_fixed_size
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, fixed_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_fixed_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns bitmap font fixed size.
*/
//go:nosplit
func (self class) FontGetFixedSize(font_rid gd.RID) gd.Int { //gd:TextServer.font_get_fixed_size
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_fixed_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets bitmap font scaling mode. This property is used only if [code]fixed_size[/code] is greater than zero.
*/
//go:nosplit
func (self class) FontSetFixedSizeScaleMode(font_rid gd.RID, fixed_size_scale_mode gdclass.TextServerFixedSizeScaleMode) { //gd:TextServer.font_set_fixed_size_scale_mode
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, fixed_size_scale_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns bitmap font scaling mode.
*/
//go:nosplit
func (self class) FontGetFixedSizeScaleMode(font_rid gd.RID) gdclass.TextServerFixedSizeScaleMode { //gd:TextServer.font_get_fixed_size_scale_mode
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gdclass.TextServerFixedSizeScaleMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code], system fonts can be automatically used as fallbacks.
*/
//go:nosplit
func (self class) FontSetAllowSystemFallback(font_rid gd.RID, allow_system_fallback bool) { //gd:TextServer.font_set_allow_system_fallback
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, allow_system_fallback)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if system fonts can be automatically used as fallbacks.
*/
//go:nosplit
func (self class) FontIsAllowSystemFallback(font_rid gd.RID) bool { //gd:TextServer.font_is_allow_system_fallback
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_is_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code] auto-hinting is preferred over font built-in hinting.
*/
//go:nosplit
func (self class) FontSetForceAutohinter(font_rid gd.RID, force_autohinter bool) { //gd:TextServer.font_set_force_autohinter
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, force_autohinter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if auto-hinting is supported and preferred over font built-in hinting. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontIsForceAutohinter(font_rid gd.RID) bool { //gd:TextServer.font_is_force_autohinter
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_is_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets font hinting mode. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontSetHinting(font_rid gd.RID, hinting gdclass.TextServerHinting) { //gd:TextServer.font_set_hinting
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, hinting)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_hinting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the font hinting mode. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontGetHinting(font_rid gd.RID) gdclass.TextServerHinting { //gd:TextServer.font_get_hinting
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gdclass.TextServerHinting](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_hinting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets font subpixel glyph positioning mode.
*/
//go:nosplit
func (self class) FontSetSubpixelPositioning(font_rid gd.RID, subpixel_positioning gdclass.TextServerSubpixelPositioning) { //gd:TextServer.font_set_subpixel_positioning
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, subpixel_positioning)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font subpixel glyph positioning mode.
*/
//go:nosplit
func (self class) FontGetSubpixelPositioning(font_rid gd.RID) gdclass.TextServerSubpixelPositioning { //gd:TextServer.font_get_subpixel_positioning
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gdclass.TextServerSubpixelPositioning](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets font embolden strength. If [param strength] is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) FontSetEmbolden(font_rid gd.RID, strength gd.Float) { //gd:TextServer.font_set_embolden
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_embolden, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font embolden strength.
*/
//go:nosplit
func (self class) FontGetEmbolden(font_rid gd.RID) gd.Float { //gd:TextServer.font_get_embolden
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_embolden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) FontSetSpacing(font_rid gd.RID, spacing gdclass.TextServerSpacingType, value gd.Int) { //gd:TextServer.font_set_spacing
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) FontGetSpacing(font_rid gd.RID, spacing gdclass.TextServerSpacingType) gd.Int { //gd:TextServer.font_get_spacing
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) FontSetBaselineOffset(font_rid gd.RID, baseline_offset gd.Float) { //gd:TextServer.font_set_baseline_offset
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, baseline_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) FontGetBaselineOffset(font_rid gd.RID) gd.Float { //gd:TextServer.font_get_baseline_offset
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
For example, to simulate italic typeface by slanting, apply the following transform [code]Transform2D(1.0, slant, 0.0, 1.0, 0.0, 0.0)[/code].
*/
//go:nosplit
func (self class) FontSetTransform(font_rid gd.RID, transform gd.Transform2D) { //gd:TextServer.font_set_transform
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns 2D transform applied to the font outlines.
*/
//go:nosplit
func (self class) FontGetTransform(font_rid gd.RID) gd.Transform2D { //gd:TextServer.font_get_transform
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets variation coordinates for the specified font cache entry. See [method font_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) FontSetVariationCoordinates(font_rid gd.RID, variation_coordinates Dictionary.Any) { //gd:TextServer.font_set_variation_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(variation_coordinates)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns variation coordinates for the specified font cache entry. See [method font_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) FontGetVariationCoordinates(font_rid gd.RID) Dictionary.Any { //gd:TextServer.font_get_variation_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontSetOversampling(font_rid gd.RID, oversampling gd.Float) { //gd:TextServer.font_set_oversampling
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, oversampling)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font oversampling factor, if set to [code]0.0[/code] global oversampling factor is used instead. Used by dynamic fonts only.
*/
//go:nosplit
func (self class) FontGetOversampling(font_rid gd.RID) gd.Float { //gd:TextServer.font_get_oversampling
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
//go:nosplit
func (self class) FontGetSizeCacheList(font_rid gd.RID) Array.Contains[gd.Vector2i] { //gd:TextServer.font_get_size_cache_list
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_size_cache_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Vector2i]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes all font sizes from the cache entry.
*/
//go:nosplit
func (self class) FontClearSizeCache(font_rid gd.RID) { //gd:TextServer.font_clear_size_cache
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_clear_size_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified font size from the cache entry.
*/
//go:nosplit
func (self class) FontRemoveSizeCache(font_rid gd.RID, size gd.Vector2i) { //gd:TextServer.font_remove_size_cache
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_remove_size_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) FontSetAscent(font_rid gd.RID, size gd.Int, ascent gd.Float) { //gd:TextServer.font_set_ascent
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, ascent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) FontGetAscent(font_rid gd.RID, size gd.Int) gd.Float { //gd:TextServer.font_get_ascent
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) FontSetDescent(font_rid gd.RID, size gd.Int, descent gd.Float) { //gd:TextServer.font_set_descent
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, descent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) FontGetDescent(font_rid gd.RID, size gd.Int) gd.Float { //gd:TextServer.font_get_descent
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) FontSetUnderlinePosition(font_rid gd.RID, size gd.Int, underline_position gd.Float) { //gd:TextServer.font_set_underline_position
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) FontGetUnderlinePosition(font_rid gd.RID, size gd.Int) gd.Float { //gd:TextServer.font_get_underline_position
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets thickness of the underline in pixels.
*/
//go:nosplit
func (self class) FontSetUnderlineThickness(font_rid gd.RID, size gd.Int, underline_thickness gd.Float) { //gd:TextServer.font_set_underline_thickness
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_thickness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns thickness of the underline in pixels.
*/
//go:nosplit
func (self class) FontGetUnderlineThickness(font_rid gd.RID, size gd.Int) gd.Float { //gd:TextServer.font_get_underline_thickness
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) FontSetScale(font_rid gd.RID, size gd.Int, scale gd.Float) { //gd:TextServer.font_set_scale
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) FontGetScale(font_rid gd.RID, size gd.Int) gd.Float { //gd:TextServer.font_get_scale
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of textures used by font cache entry.
*/
//go:nosplit
func (self class) FontGetTextureCount(font_rid gd.RID, size gd.Vector2i) gd.Int { //gd:TextServer.font_get_texture_count
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_texture_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method font_remove_glyph] to remove them manually.
*/
//go:nosplit
func (self class) FontClearTextures(font_rid gd.RID, size gd.Vector2i) { //gd:TextServer.font_clear_textures
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_clear_textures, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, remove them manually, using [method font_remove_glyph].
*/
//go:nosplit
func (self class) FontRemoveTexture(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) { //gd:TextServer.font_remove_texture
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_remove_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets font cache texture image data.
*/
//go:nosplit
func (self class) FontSetTextureImage(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, image [1]gdclass.Image) { //gd:TextServer.font_set_texture_image
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_texture_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font cache texture image data.
*/
//go:nosplit
func (self class) FontGetTextureImage(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) [1]gdclass.Image { //gd:TextServer.font_get_texture_image
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_texture_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets array containing glyph packing data.
*/
//go:nosplit
func (self class) FontSetTextureOffsets(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int, offset Packed.Array[int32]) { //gd:TextServer.font_set_texture_offsets
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, gd.InternalPacked[gd.PackedInt32Array, int32](offset))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns array containing glyph packing data.
*/
//go:nosplit
func (self class) FontGetTextureOffsets(font_rid gd.RID, size gd.Vector2i, texture_index gd.Int) Packed.Array[int32] { //gd:TextServer.font_get_texture_offsets
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns list of rendered glyphs in the cache entry.
*/
//go:nosplit
func (self class) FontGetGlyphList(font_rid gd.RID, size gd.Vector2i) Packed.Array[int32] { //gd:TextServer.font_get_glyph_list
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method font_remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) FontClearGlyphs(font_rid gd.RID, size gd.Vector2i) { //gd:TextServer.font_clear_glyphs
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_clear_glyphs, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method font_remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) FontRemoveGlyph(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) { //gd:TextServer.font_remove_glyph
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_remove_glyph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) FontGetGlyphAdvance(font_rid gd.RID, size gd.Int, glyph gd.Int) gd.Vector2 { //gd:TextServer.font_get_glyph_advance
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) FontSetGlyphAdvance(font_rid gd.RID, size gd.Int, glyph gd.Int, advance gd.Vector2) { //gd:TextServer.font_set_glyph_advance
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, advance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns glyph offset from the baseline.
*/
//go:nosplit
func (self class) FontGetGlyphOffset(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2 { //gd:TextServer.font_get_glyph_offset
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets glyph offset from the baseline.
*/
//go:nosplit
func (self class) FontSetGlyphOffset(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, offset gd.Vector2) { //gd:TextServer.font_set_glyph_offset
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns size of the glyph.
*/
//go:nosplit
func (self class) FontGetGlyphSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2 { //gd:TextServer.font_get_glyph_size
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets size of the glyph.
*/
//go:nosplit
func (self class) FontSetGlyphSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, gl_size gd.Vector2) { //gd:TextServer.font_set_glyph_size
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, gl_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_glyph_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) FontGetGlyphUvRect(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Rect2 { //gd:TextServer.font_get_glyph_uv_rect
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) FontSetGlyphUvRect(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, uv_rect gd.Rect2) { //gd:TextServer.font_set_glyph_uv_rect
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, uv_rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) FontGetGlyphTextureIdx(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Int { //gd:TextServer.font_get_glyph_texture_idx
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) FontSetGlyphTextureIdx(font_rid gd.RID, size gd.Vector2i, glyph gd.Int, texture_idx gd.Int) { //gd:TextServer.font_set_glyph_texture_idx
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, texture_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns resource ID of the cache texture containing the glyph.
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
//go:nosplit
func (self class) FontGetGlyphTextureRid(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.RID { //gd:TextServer.font_get_glyph_texture_rid
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_texture_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns size of the cache texture containing the glyph.
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
//go:nosplit
func (self class) FontGetGlyphTextureSize(font_rid gd.RID, size gd.Vector2i, glyph gd.Int) gd.Vector2 { //gd:TextServer.font_get_glyph_texture_size
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_texture_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns outline contours of the glyph as a [Dictionary] with the following contents:
[code]points[/code]         - [PackedVector3Array], containing outline points. [code]x[/code] and [code]y[/code] are point coordinates. [code]z[/code] is the type of the point, using the [enum ContourPointTag] values.
[code]contours[/code]       - [PackedInt32Array], containing indices the end points of each contour.
[code]orientation[/code]    - [bool], contour orientation. If [code]true[/code], clockwise contours must be filled.
*/
//go:nosplit
func (self class) FontGetGlyphContours(font gd.RID, size gd.Int, index gd.Int) Dictionary.Any { //gd:TextServer.font_get_glyph_contours
	var frame = callframe.New()
	callframe.Arg(frame, font)
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_contours, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns list of the kerning overrides.
*/
//go:nosplit
func (self class) FontGetKerningList(font_rid gd.RID, size gd.Int) Array.Contains[gd.Vector2i] { //gd:TextServer.font_get_kerning_list
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_kerning_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Vector2i]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes all kerning overrides.
*/
//go:nosplit
func (self class) FontClearKerningMap(font_rid gd.RID, size gd.Int) { //gd:TextServer.font_clear_kerning_map
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_clear_kerning_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes kerning override for the pair of glyphs.
*/
//go:nosplit
func (self class) FontRemoveKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i) { //gd:TextServer.font_remove_kerning
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_remove_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) FontSetKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i, kerning gd.Vector2) { //gd:TextServer.font_set_kerning
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	callframe.Arg(frame, kerning)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) FontGetKerning(font_rid gd.RID, size gd.Int, glyph_pair gd.Vector2i) gd.Vector2 { //gd:TextServer.font_get_kerning
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector]. See [method font_get_char_from_glyph_index].
*/
//go:nosplit
func (self class) FontGetGlyphIndex(font_rid gd.RID, size gd.Int, char gd.Int, variation_selector gd.Int) gd.Int { //gd:TextServer.font_get_glyph_index
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, char)
	callframe.Arg(frame, variation_selector)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method font_get_glyph_index].
*/
//go:nosplit
func (self class) FontGetCharFromGlyphIndex(font_rid gd.RID, size gd.Int, glyph_index gd.Int) gd.Int { //gd:TextServer.font_get_char_from_glyph_index
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_char_from_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a Unicode [param char] is available in the font.
*/
//go:nosplit
func (self class) FontHasChar(font_rid gd.RID, char gd.Int) bool { //gd:TextServer.font_has_char
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, char)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_has_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a string containing all the characters available in the font.
*/
//go:nosplit
func (self class) FontGetSupportedChars(font_rid gd.RID) String.Readable { //gd:TextServer.font_get_supported_chars
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_supported_chars, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Renders the range of characters to the font cache texture.
*/
//go:nosplit
func (self class) FontRenderRange(font_rid gd.RID, size gd.Vector2i, start gd.Int, end gd.Int) { //gd:TextServer.font_render_range
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_render_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Renders specified glyph to the font cache texture.
*/
//go:nosplit
func (self class) FontRenderGlyph(font_rid gd.RID, size gd.Vector2i, index gd.Int) { //gd:TextServer.font_render_glyph
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_render_glyph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws single glyph into a canvas item at the position, using [param font_rid] at the size [param size].
[b]Note:[/b] Glyph index is specific to the font, use glyphs indices returned by [method shaped_text_get_glyphs] or [method font_get_glyph_index].
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
//go:nosplit
func (self class) FontDrawGlyph(font_rid gd.RID, canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) { //gd:TextServer.font_draw_glyph
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, size)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, index)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_draw_glyph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draws single glyph outline of size [param outline_size] into a canvas item at the position, using [param font_rid] at the size [param size].
[b]Note:[/b] Glyph index is specific to the font, use glyphs indices returned by [method shaped_text_get_glyphs] or [method font_get_glyph_index].
[b]Note:[/b] If there are pending glyphs to render, calling this function might trigger the texture cache update.
*/
//go:nosplit
func (self class) FontDrawGlyphOutline(font_rid gd.RID, canvas gd.RID, size gd.Int, outline_size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) { //gd:TextServer.font_draw_glyph_outline
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, size)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, index)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_draw_glyph_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code], if font supports given language ([url=https://en.wikipedia.org/wiki/ISO_639-1]ISO 639[/url] code).
*/
//go:nosplit
func (self class) FontIsLanguageSupported(font_rid gd.RID, language String.Readable) bool { //gd:TextServer.font_is_language_supported
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_is_language_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds override for [method font_is_language_supported].
*/
//go:nosplit
func (self class) FontSetLanguageSupportOverride(font_rid gd.RID, language String.Readable, supported bool) { //gd:TextServer.font_set_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	callframe.Arg(frame, supported)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
//go:nosplit
func (self class) FontGetLanguageSupportOverride(font_rid gd.RID, language String.Readable) bool { //gd:TextServer.font_get_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Remove language support override.
*/
//go:nosplit
func (self class) FontRemoveLanguageSupportOverride(font_rid gd.RID, language String.Readable) { //gd:TextServer.font_remove_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_remove_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns list of language support overrides.
*/
//go:nosplit
func (self class) FontGetLanguageSupportOverrides(font_rid gd.RID) Packed.Strings { //gd:TextServer.font_get_language_support_overrides
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_language_support_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if font supports given script (ISO 15924 code).
*/
//go:nosplit
func (self class) FontIsScriptSupported(font_rid gd.RID, script String.Readable) bool { //gd:TextServer.font_is_script_supported
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(script)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_is_script_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds override for [method font_is_script_supported].
*/
//go:nosplit
func (self class) FontSetScriptSupportOverride(font_rid gd.RID, script String.Readable, supported bool) { //gd:TextServer.font_set_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(script)))
	callframe.Arg(frame, supported)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
//go:nosplit
func (self class) FontGetScriptSupportOverride(font_rid gd.RID, script String.Readable) bool { //gd:TextServer.font_get_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(script)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes script support override.
*/
//go:nosplit
func (self class) FontRemoveScriptSupportOverride(font_rid gd.RID, script String.Readable) { //gd:TextServer.font_remove_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalString(script)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_remove_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns list of script support overrides.
*/
//go:nosplit
func (self class) FontGetScriptSupportOverrides(font_rid gd.RID) Packed.Strings { //gd:TextServer.font_get_script_support_overrides
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_script_support_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets font OpenType feature set override.
*/
//go:nosplit
func (self class) FontSetOpentypeFeatureOverrides(font_rid gd.RID, overrides Dictionary.Any) { //gd:TextServer.font_set_opentype_feature_overrides
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(overrides)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns font OpenType feature set override.
*/
//go:nosplit
func (self class) FontGetOpentypeFeatureOverrides(font_rid gd.RID) Dictionary.Any { //gd:TextServer.font_get_opentype_feature_overrides
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the dictionary of the supported OpenType features.
*/
//go:nosplit
func (self class) FontSupportedFeatureList(font_rid gd.RID) Dictionary.Any { //gd:TextServer.font_supported_feature_list
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_supported_feature_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the dictionary of the supported OpenType variation coordinates.
*/
//go:nosplit
func (self class) FontSupportedVariationList(font_rid gd.RID) Dictionary.Any { //gd:TextServer.font_supported_variation_list
	var frame = callframe.New()
	callframe.Arg(frame, font_rid)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_supported_variation_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the font oversampling factor, shared by all fonts in the TextServer.
*/
//go:nosplit
func (self class) FontGetGlobalOversampling() gd.Float { //gd:TextServer.font_get_global_oversampling
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_get_global_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets oversampling factor, shared by all font in the TextServer.
[b]Note:[/b] This value can be automatically changed by display server.
*/
//go:nosplit
func (self class) FontSetGlobalOversampling(oversampling gd.Float) { //gd:TextServer.font_set_global_oversampling
	var frame = callframe.New()
	callframe.Arg(frame, oversampling)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_font_set_global_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns size of the replacement character (box with character hexadecimal code that is drawn in place of invalid characters).
*/
//go:nosplit
func (self class) GetHexCodeBoxSize(size gd.Int, index gd.Int) gd.Vector2 { //gd:TextServer.get_hex_code_box_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_get_hex_code_box_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draws box displaying character hexadecimal code. Used for replacing missing characters.
*/
//go:nosplit
func (self class) DrawHexCodeBox(canvas gd.RID, size gd.Int, pos gd.Vector2, index gd.Int, color gd.Color) { //gd:TextServer.draw_hex_code_box
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, size)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, index)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_draw_hex_code_box, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Creates a new buffer for complex text layout, with the given [param direction] and [param orientation]. To free the resulting buffer, use [method free_rid] method.
[b]Note:[/b] Direction is ignored if server does not support [constant FEATURE_BIDI_LAYOUT] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] Orientation is ignored if server does not support [constant FEATURE_VERTICAL_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
//go:nosplit
func (self class) CreateShapedText(direction gdclass.TextServerDirection, orientation gdclass.TextServerOrientation) gd.RID { //gd:TextServer.create_shaped_text
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_create_shaped_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears text buffer (removes text and inline objects).
*/
//go:nosplit
func (self class) ShapedTextClear(rid gd.RID) { //gd:TextServer.shaped_text_clear
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets desired text direction. If set to [constant DIRECTION_AUTO], direction will be detected based on the buffer contents and current locale.
[b]Note:[/b] Direction is ignored if server does not support [constant FEATURE_BIDI_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
//go:nosplit
func (self class) ShapedTextSetDirection(shaped gd.RID, direction gdclass.TextServerDirection) { //gd:TextServer.shaped_text_set_direction
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_set_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns direction of the text.
*/
//go:nosplit
func (self class) ShapedTextGetDirection(shaped gd.RID) gdclass.TextServerDirection { //gd:TextServer.shaped_text_get_direction
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gdclass.TextServerDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns direction of the text, inferred by the BiDi algorithm.
*/
//go:nosplit
func (self class) ShapedTextGetInferredDirection(shaped gd.RID) gdclass.TextServerDirection { //gd:TextServer.shaped_text_get_inferred_direction
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gdclass.TextServerDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_inferred_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
//go:nosplit
func (self class) ShapedTextSetBidiOverride(shaped gd.RID, override Array.Any) { //gd:TextServer.shaped_text_set_bidi_override
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(override)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_set_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
//go:nosplit
func (self class) ShapedTextSetCustomPunctuation(shaped gd.RID, punct String.Readable) { //gd:TextServer.shaped_text_set_custom_punctuation
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pointers.Get(gd.InternalString(punct)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_set_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns custom punctuation character list, used for word breaking. If set to empty string, server defaults are used.
*/
//go:nosplit
func (self class) ShapedTextGetCustomPunctuation(shaped gd.RID) String.Readable { //gd:TextServer.shaped_text_get_custom_punctuation
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets ellipsis character used for text clipping.
*/
//go:nosplit
func (self class) ShapedTextSetCustomEllipsis(shaped gd.RID, char gd.Int) { //gd:TextServer.shaped_text_set_custom_ellipsis
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, char)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_set_custom_ellipsis, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns ellipsis character used for text clipping.
*/
//go:nosplit
func (self class) ShapedTextGetCustomEllipsis(shaped gd.RID) gd.Int { //gd:TextServer.shaped_text_get_custom_ellipsis
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_custom_ellipsis, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets desired text orientation.
[b]Note:[/b] Orientation is ignored if server does not support [constant FEATURE_VERTICAL_LAYOUT] feature (supported by [TextServerAdvanced]).
*/
//go:nosplit
func (self class) ShapedTextSetOrientation(shaped gd.RID, orientation gdclass.TextServerOrientation) { //gd:TextServer.shaped_text_set_orientation
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_set_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns text orientation.
*/
//go:nosplit
func (self class) ShapedTextGetOrientation(shaped gd.RID) gdclass.TextServerOrientation { //gd:TextServer.shaped_text_get_orientation
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gdclass.TextServerOrientation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code] text buffer will display invalid characters as hexadecimal codes, otherwise nothing is displayed.
*/
//go:nosplit
func (self class) ShapedTextSetPreserveInvalid(shaped gd.RID, enabled bool) { //gd:TextServer.shaped_text_set_preserve_invalid
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_set_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if text buffer is configured to display hexadecimal codes in place of invalid characters.
[b]Note:[/b] If set to [code]false[/code], nothing is displayed in place of invalid characters.
*/
//go:nosplit
func (self class) ShapedTextGetPreserveInvalid(shaped gd.RID) bool { //gd:TextServer.shaped_text_get_preserve_invalid
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set to [code]true[/code] text buffer will display control characters.
*/
//go:nosplit
func (self class) ShapedTextSetPreserveControl(shaped gd.RID, enabled bool) { //gd:TextServer.shaped_text_set_preserve_control
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_set_preserve_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if text buffer is configured to display control characters.
*/
//go:nosplit
func (self class) ShapedTextGetPreserveControl(shaped gd.RID) bool { //gd:TextServer.shaped_text_get_preserve_control
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_preserve_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets extra spacing added between glyphs or lines in pixels.
*/
//go:nosplit
func (self class) ShapedTextSetSpacing(shaped gd.RID, spacing gdclass.TextServerSpacingType, value gd.Int) { //gd:TextServer.shaped_text_set_spacing
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_set_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns extra spacing added between glyphs or lines in pixels.
*/
//go:nosplit
func (self class) ShapedTextGetSpacing(shaped gd.RID, spacing gdclass.TextServerSpacingType) gd.Int { //gd:TextServer.shaped_text_get_spacing
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds text span and font to draw it to the text buffer.
*/
//go:nosplit
func (self class) ShapedTextAddString(shaped gd.RID, text String.Readable, fonts Array.Contains[gd.RID], size gd.Int, opentype_features Dictionary.Any, language String.Readable, meta gd.Variant) bool { //gd:TextServer.shaped_text_add_string
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(fonts)))
	callframe.Arg(frame, size)
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(opentype_features)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_add_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
//go:nosplit
func (self class) ShapedTextAddObject(shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align InlineAlignment, length gd.Int, baseline gd.Float) bool { //gd:TextServer.shaped_text_add_object
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, length)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_add_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets new size and alignment of embedded object.
*/
//go:nosplit
func (self class) ShapedTextResizeObject(shaped gd.RID, key gd.Variant, size gd.Vector2, inline_align InlineAlignment, baseline gd.Float) bool { //gd:TextServer.shaped_text_resize_object
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_resize_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of text spans added using [method shaped_text_add_string] or [method shaped_text_add_object].
*/
//go:nosplit
func (self class) ShapedGetSpanCount(shaped gd.RID) gd.Int { //gd:TextServer.shaped_get_span_count
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_get_span_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns text span metadata.
*/
//go:nosplit
func (self class) ShapedGetSpanMeta(shaped gd.RID, index gd.Int) gd.Variant { //gd:TextServer.shaped_get_span_meta
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_get_span_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Changes text span font, font size, and OpenType features, without changing the text.
*/
//go:nosplit
func (self class) ShapedSetSpanUpdateFont(shaped gd.RID, index gd.Int, fonts Array.Contains[gd.RID], size gd.Int, opentype_features Dictionary.Any) { //gd:TextServer.shaped_set_span_update_font
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(fonts)))
	callframe.Arg(frame, size)
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(opentype_features)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_set_span_update_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns text buffer for the substring of the text in the [param shaped] text buffer (including inline objects).
*/
//go:nosplit
func (self class) ShapedTextSubstr(shaped gd.RID, start gd.Int, length gd.Int) gd.RID { //gd:TextServer.shaped_text_substr
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, start)
	callframe.Arg(frame, length)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_substr, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the parent buffer from which the substring originates.
*/
//go:nosplit
func (self class) ShapedTextGetParent(shaped gd.RID) gd.RID { //gd:TextServer.shaped_text_get_parent
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adjusts text width to fit to specified width, returns new text width.
*/
//go:nosplit
func (self class) ShapedTextFitToWidth(shaped gd.RID, width gd.Float, justification_flags gdclass.TextServerJustificationFlag) gd.Float { //gd:TextServer.shaped_text_fit_to_width
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, width)
	callframe.Arg(frame, justification_flags)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_fit_to_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Aligns shaped text to the given tab-stops.
*/
//go:nosplit
func (self class) ShapedTextTabAlign(shaped gd.RID, tab_stops Packed.Array[float32]) gd.Float { //gd:TextServer.shaped_text_tab_align
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, gd.InternalPacked[gd.PackedFloat32Array, float32](tab_stops))
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_tab_align, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Shapes buffer if it's not shaped. Returns [code]true[/code] if the string is shaped successfully.
[b]Note:[/b] It is not necessary to call this function manually, buffer will be shaped automatically as soon as any of its output data is requested.
*/
//go:nosplit
func (self class) ShapedTextShape(shaped gd.RID) bool { //gd:TextServer.shaped_text_shape
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if buffer is successfully shaped.
*/
//go:nosplit
func (self class) ShapedTextIsReady(shaped gd.RID) bool { //gd:TextServer.shaped_text_is_ready
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_is_ready, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if text buffer contains any visible characters.
*/
//go:nosplit
func (self class) ShapedTextHasVisibleChars(shaped gd.RID) bool { //gd:TextServer.shaped_text_has_visible_chars
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_has_visible_chars, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of glyphs in the visual order.
*/
//go:nosplit
func (self class) ShapedTextGetGlyphs(shaped gd.RID) Array.Contains[Dictionary.Any] { //gd:TextServer.shaped_text_get_glyphs
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_glyphs, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns text glyphs in the logical order.
*/
//go:nosplit
func (self class) ShapedTextSortLogical(shaped gd.RID) Array.Contains[Dictionary.Any] { //gd:TextServer.shaped_text_sort_logical
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_sort_logical, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns number of glyphs in the buffer.
*/
//go:nosplit
func (self class) ShapedTextGetGlyphCount(shaped gd.RID) gd.Int { //gd:TextServer.shaped_text_get_glyph_count
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_glyph_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns substring buffer character range in the parent buffer.
*/
//go:nosplit
func (self class) ShapedTextGetRange(shaped gd.RID) gd.Vector2i { //gd:TextServer.shaped_text_get_range
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Breaks text to the lines and columns. Returns character ranges for each segment.
*/
//go:nosplit
func (self class) ShapedTextGetLineBreaksAdv(shaped gd.RID, width Packed.Array[float32], start gd.Int, once bool, break_flags gdclass.TextServerLineBreakFlag) Packed.Array[int32] { //gd:TextServer.shaped_text_get_line_breaks_adv
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, gd.InternalPacked[gd.PackedFloat32Array, float32](width))
	callframe.Arg(frame, start)
	callframe.Arg(frame, once)
	callframe.Arg(frame, break_flags)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_line_breaks_adv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Breaks text to the lines and returns character ranges for each line.
*/
//go:nosplit
func (self class) ShapedTextGetLineBreaks(shaped gd.RID, width gd.Float, start gd.Int, break_flags gdclass.TextServerLineBreakFlag) Packed.Array[int32] { //gd:TextServer.shaped_text_get_line_breaks
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, width)
	callframe.Arg(frame, start)
	callframe.Arg(frame, break_flags)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_line_breaks, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Breaks text into words and returns array of character ranges. Use [param grapheme_flags] to set what characters are used for breaking (see [enum GraphemeFlag]).
*/
//go:nosplit
func (self class) ShapedTextGetWordBreaks(shaped gd.RID, grapheme_flags gdclass.TextServerGraphemeFlag, skip_grapheme_flags gdclass.TextServerGraphemeFlag) Packed.Array[int32] { //gd:TextServer.shaped_text_get_word_breaks
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, grapheme_flags)
	callframe.Arg(frame, skip_grapheme_flags)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_word_breaks, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the position of the overrun trim.
*/
//go:nosplit
func (self class) ShapedTextGetTrimPos(shaped gd.RID) gd.Int { //gd:TextServer.shaped_text_get_trim_pos
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_trim_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns position of the ellipsis.
*/
//go:nosplit
func (self class) ShapedTextGetEllipsisPos(shaped gd.RID) gd.Int { //gd:TextServer.shaped_text_get_ellipsis_pos
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_ellipsis_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns array of the glyphs in the ellipsis.
*/
//go:nosplit
func (self class) ShapedTextGetEllipsisGlyphs(shaped gd.RID) Array.Contains[Dictionary.Any] { //gd:TextServer.shaped_text_get_ellipsis_glyphs
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_ellipsis_glyphs, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns number of glyphs in the ellipsis.
*/
//go:nosplit
func (self class) ShapedTextGetEllipsisGlyphCount(shaped gd.RID) gd.Int { //gd:TextServer.shaped_text_get_ellipsis_glyph_count
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_ellipsis_glyph_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Trims text if it exceeds the given width.
*/
//go:nosplit
func (self class) ShapedTextOverrunTrimToWidth(shaped gd.RID, width gd.Float, overrun_trim_flags gdclass.TextServerTextOverrunFlag) { //gd:TextServer.shaped_text_overrun_trim_to_width
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, width)
	callframe.Arg(frame, overrun_trim_flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_overrun_trim_to_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns array of inline objects.
*/
//go:nosplit
func (self class) ShapedTextGetObjects(shaped gd.RID) Array.Any { //gd:TextServer.shaped_text_get_objects
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_objects, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns bounding rectangle of the inline object.
*/
//go:nosplit
func (self class) ShapedTextGetObjectRect(shaped gd.RID, key gd.Variant) gd.Rect2 { //gd:TextServer.shaped_text_get_object_rect
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_object_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the character range of the inline object.
*/
//go:nosplit
func (self class) ShapedTextGetObjectRange(shaped gd.RID, key gd.Variant) gd.Vector2i { //gd:TextServer.shaped_text_get_object_range
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_object_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the glyph index of the inline object.
*/
//go:nosplit
func (self class) ShapedTextGetObjectGlyph(shaped gd.RID, key gd.Variant) gd.Int { //gd:TextServer.shaped_text_get_object_glyph
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_object_glyph, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns size of the text.
*/
//go:nosplit
func (self class) ShapedTextGetSize(shaped gd.RID) gd.Vector2 { //gd:TextServer.shaped_text_get_size
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
[b]Note:[/b] Overall ascent can be higher than font ascent, if some glyphs are displaced from the baseline.
*/
//go:nosplit
func (self class) ShapedTextGetAscent(shaped gd.RID) gd.Float { //gd:TextServer.shaped_text_get_ascent
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
[b]Note:[/b] Overall descent can be higher than font descent, if some glyphs are displaced from the baseline.
*/
//go:nosplit
func (self class) ShapedTextGetDescent(shaped gd.RID) gd.Float { //gd:TextServer.shaped_text_get_descent
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
//go:nosplit
func (self class) ShapedTextGetWidth(shaped gd.RID) gd.Float { //gd:TextServer.shaped_text_get_width
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) ShapedTextGetUnderlinePosition(shaped gd.RID) gd.Float { //gd:TextServer.shaped_text_get_underline_position
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns thickness of the underline.
*/
//go:nosplit
func (self class) ShapedTextGetUnderlineThickness(shaped gd.RID) gd.Float { //gd:TextServer.shaped_text_get_underline_thickness
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns shapes of the carets corresponding to the character offset [param position] in the text. Returned caret shape is 1 pixel wide rectangle.
*/
//go:nosplit
func (self class) ShapedTextGetCarets(shaped gd.RID, position gd.Int) Dictionary.Any { //gd:TextServer.shaped_text_get_carets
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_carets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns selection rectangles for the specified character range.
*/
//go:nosplit
func (self class) ShapedTextGetSelection(shaped gd.RID, start gd.Int, end gd.Int) Packed.Array[Vector2.XY] { //gd:TextServer.shaped_text_get_selection
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns grapheme index at the specified pixel offset at the baseline, or [code]-1[/code] if none is found.
*/
//go:nosplit
func (self class) ShapedTextHitTestGrapheme(shaped gd.RID, coords gd.Float) gd.Int { //gd:TextServer.shaped_text_hit_test_grapheme
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_hit_test_grapheme, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
//go:nosplit
func (self class) ShapedTextHitTestPosition(shaped gd.RID, coords gd.Float) gd.Int { //gd:TextServer.shaped_text_hit_test_position
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_hit_test_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns composite character's bounds as offsets from the start of the line.
*/
//go:nosplit
func (self class) ShapedTextGetGraphemeBounds(shaped gd.RID, pos gd.Int) gd.Vector2 { //gd:TextServer.shaped_text_get_grapheme_bounds
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_grapheme_bounds, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns grapheme end position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextNextGraphemePos(shaped gd.RID, pos gd.Int) gd.Int { //gd:TextServer.shaped_text_next_grapheme_pos
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_next_grapheme_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns grapheme start position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextPrevGraphemePos(shaped gd.RID, pos gd.Int) gd.Int { //gd:TextServer.shaped_text_prev_grapheme_pos
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_prev_grapheme_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns array of the composite character boundaries.
*/
//go:nosplit
func (self class) ShapedTextGetCharacterBreaks(shaped gd.RID) Packed.Array[int32] { //gd:TextServer.shaped_text_get_character_breaks
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_character_breaks, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns composite character end position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextNextCharacterPos(shaped gd.RID, pos gd.Int) gd.Int { //gd:TextServer.shaped_text_next_character_pos
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_next_character_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns composite character start position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextPrevCharacterPos(shaped gd.RID, pos gd.Int) gd.Int { //gd:TextServer.shaped_text_prev_character_pos
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_prev_character_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns composite character position closest to the [param pos].
*/
//go:nosplit
func (self class) ShapedTextClosestCharacterPos(shaped gd.RID, pos gd.Int) gd.Int { //gd:TextServer.shaped_text_closest_character_pos
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, pos)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_closest_character_pos, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draw shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
//go:nosplit
func (self class) ShapedTextDraw(shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, color gd.Color) { //gd:TextServer.shaped_text_draw
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, clip_l)
	callframe.Arg(frame, clip_r)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Draw the outline of the shaped text into a canvas item at a given position, with [param color]. [param pos] specifies the leftmost point of the baseline (for horizontal layout) or topmost point of the baseline (for vertical layout).
*/
//go:nosplit
func (self class) ShapedTextDrawOutline(shaped gd.RID, canvas gd.RID, pos gd.Vector2, clip_l gd.Float, clip_r gd.Float, outline_size gd.Int, color gd.Color) { //gd:TextServer.shaped_text_draw_outline
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, clip_l)
	callframe.Arg(frame, clip_r)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_draw_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns dominant direction of in the range of text.
*/
//go:nosplit
func (self class) ShapedTextGetDominantDirectionInRange(shaped gd.RID, start gd.Int, end gd.Int) gdclass.TextServerDirection { //gd:TextServer.shaped_text_get_dominant_direction_in_range
	var frame = callframe.New()
	callframe.Arg(frame, shaped)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret = callframe.Ret[gdclass.TextServerDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_shaped_text_get_dominant_direction_in_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Converts a number from the Western Arabic (0..9) to the numeral systems used in [param language].
If [param language] is omitted, the active locale will be used.
*/
//go:nosplit
func (self class) FormatNumber(number String.Readable, language String.Readable) String.Readable { //gd:TextServer.format_number
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(number)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_format_number, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Converts [param number] from the numeral systems used in [param language] to Western Arabic (0..9).
*/
//go:nosplit
func (self class) ParseNumber(number String.Readable, language String.Readable) String.Readable { //gd:TextServer.parse_number
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(number)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_parse_number, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns percent sign used in the [param language].
*/
//go:nosplit
func (self class) PercentSign(language String.Readable) String.Readable { //gd:TextServer.percent_sign
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_percent_sign, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of the word break boundaries. Elements in the returned array are the offsets of the start and end of words. Therefore the length of the array is always even.
When [param chars_per_line] is greater than zero, line break boundaries are returned instead.
[codeblock]
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_word_breaks("The Godot Engine, 4")) # Prints [0, 3, 4, 9, 10, 16, 18, 19], which corresponds to the following substrings: "The", "Godot", "Engine", "4"
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 5)) # Prints [0, 3, 4, 9, 10, 15, 15, 19], which corresponds to the following substrings: "The", "Godot", "Engin", "e, 4"
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 10)) # Prints [0, 9, 10, 19], which corresponds to the following substrings: "The Godot", "Engine, 4"
[/codeblock]
*/
//go:nosplit
func (self class) StringGetWordBreaks(s String.Readable, language String.Readable, chars_per_line gd.Int) Packed.Array[int32] { //gd:TextServer.string_get_word_breaks
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	callframe.Arg(frame, chars_per_line)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_string_get_word_breaks, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns array of the composite character boundaries.
[codeblock]
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_word_breaks("Test ❤️‍🔥 Test")) # Prints [1, 2, 3, 4, 5, 9, 10, 11, 12, 13, 14]
[/codeblock]
*/
//go:nosplit
func (self class) StringGetCharacterBreaks(s String.Readable, language String.Readable) Packed.Array[int32] { //gd:TextServer.string_get_character_breaks
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_string_get_character_breaks, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns index of the first string in [param dict] which is visually confusable with the [param string], or [code]-1[/code] if none is found.
[b]Note:[/b] This method doesn't detect invisible characters, for spoof detection use it in combination with [method spoof_check].
[b]Note:[/b] Always returns [code]-1[/code] if the server does not support the [constant FEATURE_UNICODE_SECURITY] feature.
*/
//go:nosplit
func (self class) IsConfusable(s String.Readable, dict Packed.Strings) gd.Int { //gd:TextServer.is_confusable
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(dict)))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_is_confusable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [param string] is likely to be an attempt at confusing the reader.
[b]Note:[/b] Always returns [code]false[/code] if the server does not support the [constant FEATURE_UNICODE_SECURITY] feature.
*/
//go:nosplit
func (self class) SpoofCheck(s String.Readable) bool { //gd:TextServer.spoof_check
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_spoof_check, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Strips diacritics from the string.
[b]Note:[/b] The result may be longer or shorter than the original.
*/
//go:nosplit
func (self class) StripDiacritics(s String.Readable) String.Readable { //gd:TextServer.strip_diacritics
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_strip_diacritics, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [param string] is a valid identifier.
If the text server supports the [constant FEATURE_UNICODE_IDENTIFIERS] feature, a valid identifier must:
- Conform to normalization form C.
- Begin with a Unicode character of class XID_Start or [code]"_"[/code].
- May contain Unicode characters of class XID_Continue in the other positions.
- Use UAX #31 recommended scripts only (mixed scripts are allowed).
If the [constant FEATURE_UNICODE_IDENTIFIERS] feature is not supported, a valid identifier must:
- Begin with a Unicode character of class XID_Start or [code]"_"[/code].
- May contain Unicode characters of class XID_Continue in the other positions.
*/
//go:nosplit
func (self class) IsValidIdentifier(s String.Readable) bool { //gd:TextServer.is_valid_identifier
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_is_valid_identifier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given code point is a valid letter, i.e. it belongs to the Unicode category "L".
*/
//go:nosplit
func (self class) IsValidLetter(unicode gd.Int) bool { //gd:TextServer.is_valid_letter
	var frame = callframe.New()
	callframe.Arg(frame, unicode)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_is_valid_letter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the string converted to uppercase.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
//go:nosplit
func (self class) StringToUpper(s String.Readable, language String.Readable) String.Readable { //gd:TextServer.string_to_upper
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_string_to_upper, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the string converted to lowercase.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
//go:nosplit
func (self class) StringToLower(s String.Readable, language String.Readable) String.Readable { //gd:TextServer.string_to_lower
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_string_to_lower, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the string converted to title case.
[b]Note:[/b] Casing is locale dependent and context sensitive if server support [constant FEATURE_CONTEXT_SENSITIVE_CASE_CONVERSION] feature (supported by [TextServerAdvanced]).
[b]Note:[/b] The result may be longer or shorter than the original.
*/
//go:nosplit
func (self class) StringToTitle(s String.Readable, language String.Readable) String.Readable { //gd:TextServer.string_to_title
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_string_to_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Default implementation of the BiDi algorithm override function. See [enum StructuredTextParser] for more info.
*/
//go:nosplit
func (self class) ParseStructuredText(parser_type gdclass.TextServerStructuredTextParser, args Array.Any, text String.Readable) Array.Contains[gd.Vector3i] { //gd:TextServer.parse_structured_text
	var frame = callframe.New()
	callframe.Arg(frame, parser_type)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(args)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextServer.Bind_parse_structured_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Vector3i]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsTextServer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextServer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("TextServer", func(ptr gd.Object) any { return [1]gdclass.TextServer{*(*gdclass.TextServer)(unsafe.Pointer(&ptr))} })
}

type FontAntialiasing = gdclass.TextServerFontAntialiasing //gd:TextServer.FontAntialiasing

const (
	/*Font glyphs are rasterized as 1-bit bitmaps.*/
	FontAntialiasingNone FontAntialiasing = 0
	/*Font glyphs are rasterized as 8-bit grayscale anti-aliased bitmaps.*/
	FontAntialiasingGray FontAntialiasing = 1
	/*Font glyphs are rasterized for LCD screens.
	  LCD subpixel layout is determined by the value of [code]gui/theme/lcd_subpixel_layout[/code] project settings.
	  LCD subpixel anti-aliasing mode is suitable only for rendering horizontal, unscaled text in 2D.*/
	FontAntialiasingLcd FontAntialiasing = 2
)

type FontLCDSubpixelLayout = gdclass.TextServerFontLCDSubpixelLayout //gd:TextServer.FontLCDSubpixelLayout

const (
	/*Unknown or unsupported subpixel layout, LCD subpixel antialiasing is disabled.*/
	FontLcdSubpixelLayoutNone FontLCDSubpixelLayout = 0
	/*Horizontal RGB subpixel layout.*/
	FontLcdSubpixelLayoutHrgb FontLCDSubpixelLayout = 1
	/*Horizontal BGR subpixel layout.*/
	FontLcdSubpixelLayoutHbgr FontLCDSubpixelLayout = 2
	/*Vertical RGB subpixel layout.*/
	FontLcdSubpixelLayoutVrgb FontLCDSubpixelLayout = 3
	/*Vertical BGR subpixel layout.*/
	FontLcdSubpixelLayoutVbgr FontLCDSubpixelLayout = 4
	/*Represents the size of the [enum FontLCDSubpixelLayout] enum.*/
	FontLcdSubpixelLayoutMax FontLCDSubpixelLayout = 5
)

type Direction = gdclass.TextServerDirection //gd:TextServer.Direction

const (
	/*Text direction is determined based on contents and current locale.*/
	DirectionAuto Direction = 0
	/*Text is written from left to right.*/
	DirectionLtr Direction = 1
	/*Text is written from right to left.*/
	DirectionRtl Direction = 2
	/*Text writing direction is the same as base string writing direction. Used for BiDi override only.*/
	DirectionInherited Direction = 3
)

type Orientation = gdclass.TextServerOrientation //gd:TextServer.Orientation

const (
	/*Text is written horizontally.*/
	OrientationHorizontal Orientation = 0
	/*Left to right text is written vertically from top to bottom.
	  Right to left text is written vertically from bottom to top.*/
	OrientationVertical Orientation = 1
)

type JustificationFlag = gdclass.TextServerJustificationFlag //gd:TextServer.JustificationFlag

const (
	/*Do not justify text.*/
	JustificationNone JustificationFlag = 0
	/*Justify text by adding and removing kashidas.*/
	JustificationKashida JustificationFlag = 1
	/*Justify text by changing width of the spaces between the words.*/
	JustificationWordBound JustificationFlag = 2
	/*Remove trailing and leading spaces from the justified text.*/
	JustificationTrimEdgeSpaces JustificationFlag = 4
	/*Only apply justification to the part of the text after the last tab.*/
	JustificationAfterLastTab JustificationFlag = 8
	/*Apply justification to the trimmed line with ellipsis.*/
	JustificationConstrainEllipsis JustificationFlag = 16
	/*Do not apply justification to the last line of the paragraph.*/
	JustificationSkipLastLine JustificationFlag = 32
	/*Do not apply justification to the last line of the paragraph with visible characters (takes precedence over [constant JUSTIFICATION_SKIP_LAST_LINE]).*/
	JustificationSkipLastLineWithVisibleChars JustificationFlag = 64
	/*Always apply justification to the paragraphs with a single line ([constant JUSTIFICATION_SKIP_LAST_LINE] and [constant JUSTIFICATION_SKIP_LAST_LINE_WITH_VISIBLE_CHARS] are ignored).*/
	JustificationDoNotSkipSingleLine JustificationFlag = 128
)

type AutowrapMode = gdclass.TextServerAutowrapMode //gd:TextServer.AutowrapMode

const (
	/*Autowrap is disabled.*/
	AutowrapOff AutowrapMode = 0
	/*Wraps the text inside the node's bounding rectangle by allowing to break lines at arbitrary positions, which is useful when very limited space is available.*/
	AutowrapArbitrary AutowrapMode = 1
	/*Wraps the text inside the node's bounding rectangle by soft-breaking between words.*/
	AutowrapWord AutowrapMode = 2
	/*Behaves similarly to [constant AUTOWRAP_WORD], but force-breaks a word if that single word does not fit in one line.*/
	AutowrapWordSmart AutowrapMode = 3
)

type LineBreakFlag = gdclass.TextServerLineBreakFlag //gd:TextServer.LineBreakFlag

const (
	/*Do not break the line.*/
	BreakNone LineBreakFlag = 0
	/*Break the line at the line mandatory break characters (e.g. [code]"\n"[/code]).*/
	BreakMandatory LineBreakFlag = 1
	/*Break the line between the words.*/
	BreakWordBound LineBreakFlag = 2
	/*Break the line between any unconnected graphemes.*/
	BreakGraphemeBound LineBreakFlag = 4
	/*Should be used only in conjunction with [constant BREAK_WORD_BOUND], break the line between any unconnected graphemes, if it's impossible to break it between the words.*/
	BreakAdaptive LineBreakFlag = 8
	/*Remove edge spaces from the broken line segments.*/
	BreakTrimEdgeSpaces LineBreakFlag = 16
	/*Subtract first line indentation width from all lines after the first one.*/
	BreakTrimIndent LineBreakFlag = 32
)

type VisibleCharactersBehavior = gdclass.TextServerVisibleCharactersBehavior //gd:TextServer.VisibleCharactersBehavior

const (
	/*Trims text before the shaping. e.g, increasing [member Label.visible_characters] or [member RichTextLabel.visible_characters] value is visually identical to typing the text.*/
	VcCharsBeforeShaping VisibleCharactersBehavior = 0
	/*Displays glyphs that are mapped to the first [member Label.visible_characters] or [member RichTextLabel.visible_characters] characters from the beginning of the text.*/
	VcCharsAfterShaping VisibleCharactersBehavior = 1
	/*Displays [member Label.visible_ratio] or [member RichTextLabel.visible_ratio] glyphs, starting from the left or from the right, depending on [member Control.layout_direction] value.*/
	VcGlyphsAuto VisibleCharactersBehavior = 2
	/*Displays [member Label.visible_ratio] or [member RichTextLabel.visible_ratio] glyphs, starting from the left.*/
	VcGlyphsLtr VisibleCharactersBehavior = 3
	/*Displays [member Label.visible_ratio] or [member RichTextLabel.visible_ratio] glyphs, starting from the right.*/
	VcGlyphsRtl VisibleCharactersBehavior = 4
)

type OverrunBehavior = gdclass.TextServerOverrunBehavior //gd:TextServer.OverrunBehavior

const (
	/*No text trimming is performed.*/
	OverrunNoTrimming OverrunBehavior = 0
	/*Trims the text per character.*/
	OverrunTrimChar OverrunBehavior = 1
	/*Trims the text per word.*/
	OverrunTrimWord OverrunBehavior = 2
	/*Trims the text per character and adds an ellipsis to indicate that parts are hidden.*/
	OverrunTrimEllipsis OverrunBehavior = 3
	/*Trims the text per word and adds an ellipsis to indicate that parts are hidden.*/
	OverrunTrimWordEllipsis OverrunBehavior = 4
)

type TextOverrunFlag = gdclass.TextServerTextOverrunFlag //gd:TextServer.TextOverrunFlag

const (
	/*No trimming is performed.*/
	OverrunNoTrim TextOverrunFlag = 0
	/*Trims the text when it exceeds the given width.*/
	OverrunTrim TextOverrunFlag = 1
	/*Trims the text per word instead of per grapheme.*/
	OverrunTrimWordOnly TextOverrunFlag = 2
	/*Determines whether an ellipsis should be added at the end of the text.*/
	OverrunAddEllipsis TextOverrunFlag = 4
	/*Determines whether the ellipsis at the end of the text is enforced and may not be hidden.*/
	OverrunEnforceEllipsis TextOverrunFlag = 8
	/*Accounts for the text being justified before attempting to trim it (see [enum JustificationFlag]).*/
	OverrunJustificationAware TextOverrunFlag = 16
)

type GraphemeFlag = gdclass.TextServerGraphemeFlag //gd:TextServer.GraphemeFlag

const (
	/*Grapheme is supported by the font, and can be drawn.*/
	GraphemeIsValid GraphemeFlag = 1
	/*Grapheme is part of right-to-left or bottom-to-top run.*/
	GraphemeIsRtl GraphemeFlag = 2
	/*Grapheme is not part of source text, it was added by justification process.*/
	GraphemeIsVirtual GraphemeFlag = 4
	/*Grapheme is whitespace.*/
	GraphemeIsSpace GraphemeFlag = 8
	/*Grapheme is mandatory break point (e.g. [code]"\n"[/code]).*/
	GraphemeIsBreakHard GraphemeFlag = 16
	/*Grapheme is optional break point (e.g. space).*/
	GraphemeIsBreakSoft GraphemeFlag = 32
	/*Grapheme is the tabulation character.*/
	GraphemeIsTab GraphemeFlag = 64
	/*Grapheme is kashida.*/
	GraphemeIsElongation GraphemeFlag = 128
	/*Grapheme is punctuation character.*/
	GraphemeIsPunctuation GraphemeFlag = 256
	/*Grapheme is underscore character.*/
	GraphemeIsUnderscore GraphemeFlag = 512
	/*Grapheme is connected to the previous grapheme. Breaking line before this grapheme is not safe.*/
	GraphemeIsConnected GraphemeFlag = 1024
	/*It is safe to insert a U+0640 before this grapheme for elongation.*/
	GraphemeIsSafeToInsertTatweel GraphemeFlag = 2048
	/*Grapheme is an object replacement character for the embedded object.*/
	GraphemeIsEmbeddedObject GraphemeFlag = 4096
	/*Grapheme is a soft hyphen.*/
	GraphemeIsSoftHyphen GraphemeFlag = 8192
)

type Hinting = gdclass.TextServerHinting //gd:TextServer.Hinting

const (
	/*Disables font hinting (smoother but less crisp).*/
	HintingNone Hinting = 0
	/*Use the light font hinting mode.*/
	HintingLight Hinting = 1
	/*Use the default font hinting mode (crisper but less smooth).
	  [b]Note:[/b] This hinting mode changes both horizontal and vertical glyph metrics. If applied to monospace font, some glyphs might have different width.*/
	HintingNormal Hinting = 2
)

type SubpixelPositioning = gdclass.TextServerSubpixelPositioning //gd:TextServer.SubpixelPositioning

const (
	/*Glyph horizontal position is rounded to the whole pixel size, each glyph is rasterized once.*/
	SubpixelPositioningDisabled SubpixelPositioning = 0
	/*Glyph horizontal position is rounded based on font size.
	  - To one quarter of the pixel size if font size is smaller or equal to [constant SUBPIXEL_POSITIONING_ONE_QUARTER_MAX_SIZE].
	  - To one half of the pixel size if font size is smaller or equal to [constant SUBPIXEL_POSITIONING_ONE_HALF_MAX_SIZE].
	  - To the whole pixel size for larger fonts.*/
	SubpixelPositioningAuto SubpixelPositioning = 1
	/*Glyph horizontal position is rounded to one half of the pixel size, each glyph is rasterized up to two times.*/
	SubpixelPositioningOneHalf SubpixelPositioning = 2
	/*Glyph horizontal position is rounded to one quarter of the pixel size, each glyph is rasterized up to four times.*/
	SubpixelPositioningOneQuarter SubpixelPositioning = 3
	/*Maximum font size which will use one half of the pixel subpixel positioning in [constant SUBPIXEL_POSITIONING_AUTO] mode.*/
	SubpixelPositioningOneHalfMaxSize SubpixelPositioning = 20
	/*Maximum font size which will use one quarter of the pixel subpixel positioning in [constant SUBPIXEL_POSITIONING_AUTO] mode.*/
	SubpixelPositioningOneQuarterMaxSize SubpixelPositioning = 16
)

type Feature = gdclass.TextServerFeature //gd:TextServer.Feature

const (
	/*TextServer supports simple text layouts.*/
	FeatureSimpleLayout Feature = 1
	/*TextServer supports bidirectional text layouts.*/
	FeatureBidiLayout Feature = 2
	/*TextServer supports vertical layouts.*/
	FeatureVerticalLayout Feature = 4
	/*TextServer supports complex text shaping.*/
	FeatureShaping Feature = 8
	/*TextServer supports justification using kashidas.*/
	FeatureKashidaJustification Feature = 16
	/*TextServer supports complex line/word breaking rules (e.g. dictionary based).*/
	FeatureBreakIterators Feature = 32
	/*TextServer supports loading bitmap fonts.*/
	FeatureFontBitmap Feature = 64
	/*TextServer supports loading dynamic (TrueType, OpeType, etc.) fonts.*/
	FeatureFontDynamic Feature = 128
	/*TextServer supports multichannel signed distance field dynamic font rendering.*/
	FeatureFontMsdf Feature = 256
	/*TextServer supports loading system fonts.*/
	FeatureFontSystem Feature = 512
	/*TextServer supports variable fonts.*/
	FeatureFontVariable Feature = 1024
	/*TextServer supports locale dependent and context sensitive case conversion.*/
	FeatureContextSensitiveCaseConversion Feature = 2048
	/*TextServer require external data file for some features, see [method load_support_data].*/
	FeatureUseSupportData Feature = 4096
	/*TextServer supports UAX #31 identifier validation, see [method is_valid_identifier].*/
	FeatureUnicodeIdentifiers Feature = 8192
	/*TextServer supports [url=https://unicode.org/reports/tr36/]Unicode Technical Report #36[/url] and [url=https://unicode.org/reports/tr39/]Unicode Technical Standard #39[/url] based spoof detection features.*/
	FeatureUnicodeSecurity Feature = 16384
)

type ContourPointTag = gdclass.TextServerContourPointTag //gd:TextServer.ContourPointTag

const (
	/*Contour point is on the curve.*/
	ContourCurveTagOn ContourPointTag = 1
	/*Contour point isn't on the curve, but serves as a control point for a conic (quadratic) Bézier arc.*/
	ContourCurveTagOffConic ContourPointTag = 0
	/*Contour point isn't on the curve, but serves as a control point for a cubic Bézier arc.*/
	ContourCurveTagOffCubic ContourPointTag = 2
)

type SpacingType = gdclass.TextServerSpacingType //gd:TextServer.SpacingType

const (
	/*Spacing for each glyph.*/
	SpacingGlyph SpacingType = 0
	/*Spacing for the space character.*/
	SpacingSpace SpacingType = 1
	/*Spacing at the top of the line.*/
	SpacingTop SpacingType = 2
	/*Spacing at the bottom of the line.*/
	SpacingBottom SpacingType = 3
	/*Represents the size of the [enum SpacingType] enum.*/
	SpacingMax SpacingType = 4
)

type FontStyle = gdclass.TextServerFontStyle //gd:TextServer.FontStyle

const (
	/*Font is bold.*/
	FontBold FontStyle = 1
	/*Font is italic or oblique.*/
	FontItalic FontStyle = 2
	/*Font have fixed-width characters.*/
	FontFixedWidth FontStyle = 4
)

type StructuredTextParser = gdclass.TextServerStructuredTextParser //gd:TextServer.StructuredTextParser

const (
	/*Use default Unicode BiDi algorithm.*/
	StructuredTextDefault StructuredTextParser = 0
	/*BiDi override for URI.*/
	StructuredTextUri StructuredTextParser = 1
	/*BiDi override for file path.*/
	StructuredTextFile StructuredTextParser = 2
	/*BiDi override for email.*/
	StructuredTextEmail StructuredTextParser = 3
	/*BiDi override for lists. Structured text options: list separator [String].*/
	StructuredTextList StructuredTextParser = 4
	/*BiDi override for GDScript.*/
	StructuredTextGdscript StructuredTextParser = 5
	/*User defined structured text BiDi override function.*/
	StructuredTextCustom StructuredTextParser = 6
)

type FixedSizeScaleMode = gdclass.TextServerFixedSizeScaleMode //gd:TextServer.FixedSizeScaleMode

const (
	/*Bitmap font is not scaled.*/
	FixedSizeScaleDisable FixedSizeScaleMode = 0
	/*Bitmap font is scaled to the closest integer multiple of the font's fixed size. This is the recommended option for pixel art fonts.*/
	FixedSizeScaleIntegerOnly FixedSizeScaleMode = 1
	/*Bitmap font is scaled to an arbitrary (fractional) size. This is the recommended option for non-pixel art fonts.*/
	FixedSizeScaleEnabled FixedSizeScaleMode = 2
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
