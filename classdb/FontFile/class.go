// Package FontFile provides methods for working with FontFile object instances.
package FontFile

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
import "graphics.gd/classdb/Font"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Rect2"

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

/*
[FontFile] contains a set of glyphs to represent Unicode characters imported from a font file, as well as a cache of rasterized glyphs, and a set of fallback [Font]s to use.
Use [FontVariation] to access specific OpenType variation of the font, create simulated bold / slanted version, and draw lines of text.
For more complex text processing, use [FontVariation] in conjunction with [TextLine] or [TextParagraph].
Supported font formats:
- Dynamic font importer: TrueType (.ttf), TrueType collection (.ttc), OpenType (.otf), OpenType collection (.otc), WOFF (.woff), WOFF2 (.woff2), Type 1 (.pfb, .pfm).
- Bitmap font importer: AngelCode BMFont (.fnt, .font), text and binary (version 3) format variants.
- Monospace image font importer: All supported image formats.
[b]Note:[/b] A character is a symbol that represents an item (letter, digit etc.) in an abstract way.
[b]Note:[/b] A glyph is a bitmap or a shape used to draw one or more characters in a context-dependent manner. Glyph indices are bound to the specific font data source.
[b]Note:[/b] If none of the font data sources contain glyphs for a character used in a string, the character in question will be replaced with a box displaying its hexadecimal code.
[codeblocks]
[gdscript]
var f = load("res://BarlowCondensed-Bold.ttf")
$Label.add_theme_font_override("font", f)
$Label.add_theme_font_size_override("font_size", 64)
[/gdscript]
[csharp]
var f = ResourceLoader.Load<FontFile>("res://BarlowCondensed-Bold.ttf");
GetNode("Label").AddThemeFontOverride("font", f);
GetNode("Label").AddThemeFontSizeOverride("font_size", 64);
[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.FontFile

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFontFile() Instance
}

/*
Loads an AngelCode BMFont (.fnt, .font) bitmap font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
func (self Instance) LoadBitmapFont(path string) error { //gd:FontFile.load_bitmap_font
	return error(gd.ToError(class(self).LoadBitmapFont(gd.NewString(path))))
}

/*
Loads a TrueType (.ttf), OpenType (.otf), WOFF (.woff), WOFF2 (.woff2) or Type 1 (.pfb, .pfm) dynamic font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
func (self Instance) LoadDynamicFont(path string) error { //gd:FontFile.load_dynamic_font
	return error(gd.ToError(class(self).LoadDynamicFont(gd.NewString(path))))
}

/*
Returns number of the font cache entries.
*/
func (self Instance) GetCacheCount() int { //gd:FontFile.get_cache_count
	return int(int(class(self).GetCacheCount()))
}

/*
Removes all font cache entries.
*/
func (self Instance) ClearCache() { //gd:FontFile.clear_cache
	class(self).ClearCache()
}

/*
Removes specified font cache entry.
*/
func (self Instance) RemoveCache(cache_index int) { //gd:FontFile.remove_cache
	class(self).RemoveCache(gd.Int(cache_index))
}

/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
func (self Instance) GetSizeCacheList(cache_index int) []Vector2i.XY { //gd:FontFile.get_size_cache_list
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetSizeCacheList(gd.Int(cache_index)))))
}

/*
Removes all font sizes from the cache entry
*/
func (self Instance) ClearSizeCache(cache_index int) { //gd:FontFile.clear_size_cache
	class(self).ClearSizeCache(gd.Int(cache_index))
}

/*
Removes specified font size from the cache entry.
*/
func (self Instance) RemoveSizeCache(cache_index int, size Vector2i.XY) { //gd:FontFile.remove_size_cache
	class(self).RemoveSizeCache(gd.Int(cache_index), gd.Vector2i(size))
}

/*
Sets variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
func (self Instance) SetVariationCoordinates(cache_index int, variation_coordinates map[string]float32) { //gd:FontFile.set_variation_coordinates
	class(self).SetVariationCoordinates(gd.Int(cache_index), gd.DictionaryFromMap(variation_coordinates))
}

/*
Returns variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
func (self Instance) GetVariationCoordinates(cache_index int) map[string]float32 { //gd:FontFile.get_variation_coordinates
	return map[string]float32(gd.DictionaryAs[map[string]float32](class(self).GetVariationCoordinates(gd.Int(cache_index))))
}

/*
Sets embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (self Instance) SetEmbolden(cache_index int, strength Float.X) { //gd:FontFile.set_embolden
	class(self).SetEmbolden(gd.Int(cache_index), gd.Float(strength))
}

/*
Returns embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (self Instance) GetEmbolden(cache_index int) Float.X { //gd:FontFile.get_embolden
	return Float.X(Float.X(class(self).GetEmbolden(gd.Int(cache_index))))
}

/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
func (self Instance) SetTransform(cache_index int, transform Transform2D.OriginXY) { //gd:FontFile.set_transform
	class(self).SetTransform(gd.Int(cache_index), gd.Transform2D(transform))
}

/*
Returns 2D transform, applied to the font outlines, can be used for slanting, flipping and rotating glyphs.
*/
func (self Instance) GetTransform(cache_index int) Transform2D.OriginXY { //gd:FontFile.get_transform
	return Transform2D.OriginXY(class(self).GetTransform(gd.Int(cache_index)))
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
func (self Instance) SetExtraSpacing(cache_index int, spacing gdclass.TextServerSpacingType, value int) { //gd:FontFile.set_extra_spacing
	class(self).SetExtraSpacing(gd.Int(cache_index), spacing, gd.Int(value))
}

/*
Returns spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
func (self Instance) GetExtraSpacing(cache_index int, spacing gdclass.TextServerSpacingType) int { //gd:FontFile.get_extra_spacing
	return int(int(class(self).GetExtraSpacing(gd.Int(cache_index), spacing)))
}

/*
Sets extra baseline offset (as a fraction of font height).
*/
func (self Instance) SetExtraBaselineOffset(cache_index int, baseline_offset Float.X) { //gd:FontFile.set_extra_baseline_offset
	class(self).SetExtraBaselineOffset(gd.Int(cache_index), gd.Float(baseline_offset))
}

/*
Returns extra baseline offset (as a fraction of font height).
*/
func (self Instance) GetExtraBaselineOffset(cache_index int) Float.X { //gd:FontFile.get_extra_baseline_offset
	return Float.X(Float.X(class(self).GetExtraBaselineOffset(gd.Int(cache_index))))
}

/*
Sets an active face index in the TrueType / OpenType collection.
*/
func (self Instance) SetFaceIndex(cache_index int, face_index int) { //gd:FontFile.set_face_index
	class(self).SetFaceIndex(gd.Int(cache_index), gd.Int(face_index))
}

/*
Returns an active face index in the TrueType / OpenType collection.
*/
func (self Instance) GetFaceIndex(cache_index int) int { //gd:FontFile.get_face_index
	return int(int(class(self).GetFaceIndex(gd.Int(cache_index))))
}

/*
Sets the font ascent (number of pixels above the baseline).
*/
func (self Instance) SetCacheAscent(cache_index int, size int, ascent Float.X) { //gd:FontFile.set_cache_ascent
	class(self).SetCacheAscent(gd.Int(cache_index), gd.Int(size), gd.Float(ascent))
}

/*
Returns the font ascent (number of pixels above the baseline).
*/
func (self Instance) GetCacheAscent(cache_index int, size int) Float.X { //gd:FontFile.get_cache_ascent
	return Float.X(Float.X(class(self).GetCacheAscent(gd.Int(cache_index), gd.Int(size))))
}

/*
Sets the font descent (number of pixels below the baseline).
*/
func (self Instance) SetCacheDescent(cache_index int, size int, descent Float.X) { //gd:FontFile.set_cache_descent
	class(self).SetCacheDescent(gd.Int(cache_index), gd.Int(size), gd.Float(descent))
}

/*
Returns the font descent (number of pixels below the baseline).
*/
func (self Instance) GetCacheDescent(cache_index int, size int) Float.X { //gd:FontFile.get_cache_descent
	return Float.X(Float.X(class(self).GetCacheDescent(gd.Int(cache_index), gd.Int(size))))
}

/*
Sets pixel offset of the underline below the baseline.
*/
func (self Instance) SetCacheUnderlinePosition(cache_index int, size int, underline_position Float.X) { //gd:FontFile.set_cache_underline_position
	class(self).SetCacheUnderlinePosition(gd.Int(cache_index), gd.Int(size), gd.Float(underline_position))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Instance) GetCacheUnderlinePosition(cache_index int, size int) Float.X { //gd:FontFile.get_cache_underline_position
	return Float.X(Float.X(class(self).GetCacheUnderlinePosition(gd.Int(cache_index), gd.Int(size))))
}

/*
Sets thickness of the underline in pixels.
*/
func (self Instance) SetCacheUnderlineThickness(cache_index int, size int, underline_thickness Float.X) { //gd:FontFile.set_cache_underline_thickness
	class(self).SetCacheUnderlineThickness(gd.Int(cache_index), gd.Int(size), gd.Float(underline_thickness))
}

/*
Returns thickness of the underline in pixels.
*/
func (self Instance) GetCacheUnderlineThickness(cache_index int, size int) Float.X { //gd:FontFile.get_cache_underline_thickness
	return Float.X(Float.X(class(self).GetCacheUnderlineThickness(gd.Int(cache_index), gd.Int(size))))
}

/*
Sets scaling factor of the color bitmap font.
*/
func (self Instance) SetCacheScale(cache_index int, size int, scale Float.X) { //gd:FontFile.set_cache_scale
	class(self).SetCacheScale(gd.Int(cache_index), gd.Int(size), gd.Float(scale))
}

/*
Returns scaling factor of the color bitmap font.
*/
func (self Instance) GetCacheScale(cache_index int, size int) Float.X { //gd:FontFile.get_cache_scale
	return Float.X(Float.X(class(self).GetCacheScale(gd.Int(cache_index), gd.Int(size))))
}

/*
Returns number of textures used by font cache entry.
*/
func (self Instance) GetTextureCount(cache_index int, size Vector2i.XY) int { //gd:FontFile.get_texture_count
	return int(int(class(self).GetTextureCount(gd.Int(cache_index), gd.Vector2i(size))))
}

/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method remove_glyph] to remove them manually.
*/
func (self Instance) ClearTextures(cache_index int, size Vector2i.XY) { //gd:FontFile.clear_textures
	class(self).ClearTextures(gd.Int(cache_index), gd.Vector2i(size))
}

/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture. Remove them manually using [method remove_glyph].
*/
func (self Instance) RemoveTexture(cache_index int, size Vector2i.XY, texture_index int) { //gd:FontFile.remove_texture
	class(self).RemoveTexture(gd.Int(cache_index), gd.Vector2i(size), gd.Int(texture_index))
}

/*
Sets font cache texture image.
*/
func (self Instance) SetTextureImage(cache_index int, size Vector2i.XY, texture_index int, image [1]gdclass.Image) { //gd:FontFile.set_texture_image
	class(self).SetTextureImage(gd.Int(cache_index), gd.Vector2i(size), gd.Int(texture_index), image)
}

/*
Returns a copy of the font cache texture image.
*/
func (self Instance) GetTextureImage(cache_index int, size Vector2i.XY, texture_index int) [1]gdclass.Image { //gd:FontFile.get_texture_image
	return [1]gdclass.Image(class(self).GetTextureImage(gd.Int(cache_index), gd.Vector2i(size), gd.Int(texture_index)))
}

/*
Sets array containing glyph packing data.
*/
func (self Instance) SetTextureOffsets(cache_index int, size Vector2i.XY, texture_index int, offset []int32) { //gd:FontFile.set_texture_offsets
	class(self).SetTextureOffsets(gd.Int(cache_index), gd.Vector2i(size), gd.Int(texture_index), gd.NewPackedInt32Slice(offset))
}

/*
Returns a copy of the array containing glyph packing data.
*/
func (self Instance) GetTextureOffsets(cache_index int, size Vector2i.XY, texture_index int) []int32 { //gd:FontFile.get_texture_offsets
	return []int32(class(self).GetTextureOffsets(gd.Int(cache_index), gd.Vector2i(size), gd.Int(texture_index)).AsSlice())
}

/*
Returns list of rendered glyphs in the cache entry.
*/
func (self Instance) GetGlyphList(cache_index int, size Vector2i.XY) []int32 { //gd:FontFile.get_glyph_list
	return []int32(class(self).GetGlyphList(gd.Int(cache_index), gd.Vector2i(size)).AsSlice())
}

/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
func (self Instance) ClearGlyphs(cache_index int, size Vector2i.XY) { //gd:FontFile.clear_glyphs
	class(self).ClearGlyphs(gd.Int(cache_index), gd.Vector2i(size))
}

/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
func (self Instance) RemoveGlyph(cache_index int, size Vector2i.XY, glyph int) { //gd:FontFile.remove_glyph
	class(self).RemoveGlyph(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph))
}

/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Instance) SetGlyphAdvance(cache_index int, size int, glyph int, advance Vector2.XY) { //gd:FontFile.set_glyph_advance
	class(self).SetGlyphAdvance(gd.Int(cache_index), gd.Int(size), gd.Int(glyph), gd.Vector2(advance))
}

/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Instance) GetGlyphAdvance(cache_index int, size int, glyph int) Vector2.XY { //gd:FontFile.get_glyph_advance
	return Vector2.XY(class(self).GetGlyphAdvance(gd.Int(cache_index), gd.Int(size), gd.Int(glyph)))
}

/*
Sets glyph offset from the baseline.
*/
func (self Instance) SetGlyphOffset(cache_index int, size Vector2i.XY, glyph int, offset Vector2.XY) { //gd:FontFile.set_glyph_offset
	class(self).SetGlyphOffset(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph), gd.Vector2(offset))
}

/*
Returns glyph offset from the baseline.
*/
func (self Instance) GetGlyphOffset(cache_index int, size Vector2i.XY, glyph int) Vector2.XY { //gd:FontFile.get_glyph_offset
	return Vector2.XY(class(self).GetGlyphOffset(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph)))
}

/*
Sets glyph size.
*/
func (self Instance) SetGlyphSize(cache_index int, size Vector2i.XY, glyph int, gl_size Vector2.XY) { //gd:FontFile.set_glyph_size
	class(self).SetGlyphSize(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph), gd.Vector2(gl_size))
}

/*
Returns glyph size.
*/
func (self Instance) GetGlyphSize(cache_index int, size Vector2i.XY, glyph int) Vector2.XY { //gd:FontFile.get_glyph_size
	return Vector2.XY(class(self).GetGlyphSize(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph)))
}

/*
Sets rectangle in the cache texture containing the glyph.
*/
func (self Instance) SetGlyphUvRect(cache_index int, size Vector2i.XY, glyph int, uv_rect Rect2.PositionSize) { //gd:FontFile.set_glyph_uv_rect
	class(self).SetGlyphUvRect(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph), gd.Rect2(uv_rect))
}

/*
Returns rectangle in the cache texture containing the glyph.
*/
func (self Instance) GetGlyphUvRect(cache_index int, size Vector2i.XY, glyph int) Rect2.PositionSize { //gd:FontFile.get_glyph_uv_rect
	return Rect2.PositionSize(class(self).GetGlyphUvRect(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph)))
}

/*
Sets index of the cache texture containing the glyph.
*/
func (self Instance) SetGlyphTextureIdx(cache_index int, size Vector2i.XY, glyph int, texture_idx int) { //gd:FontFile.set_glyph_texture_idx
	class(self).SetGlyphTextureIdx(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph), gd.Int(texture_idx))
}

/*
Returns index of the cache texture containing the glyph.
*/
func (self Instance) GetGlyphTextureIdx(cache_index int, size Vector2i.XY, glyph int) int { //gd:FontFile.get_glyph_texture_idx
	return int(int(class(self).GetGlyphTextureIdx(gd.Int(cache_index), gd.Vector2i(size), gd.Int(glyph))))
}

/*
Returns list of the kerning overrides.
*/
func (self Instance) GetKerningList(cache_index int, size int) []Vector2i.XY { //gd:FontFile.get_kerning_list
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetKerningList(gd.Int(cache_index), gd.Int(size)))))
}

/*
Removes all kerning overrides.
*/
func (self Instance) ClearKerningMap(cache_index int, size int) { //gd:FontFile.clear_kerning_map
	class(self).ClearKerningMap(gd.Int(cache_index), gd.Int(size))
}

/*
Removes kerning override for the pair of glyphs.
*/
func (self Instance) RemoveKerning(cache_index int, size int, glyph_pair Vector2i.XY) { //gd:FontFile.remove_kerning
	class(self).RemoveKerning(gd.Int(cache_index), gd.Int(size), gd.Vector2i(glyph_pair))
}

/*
Sets kerning for the pair of glyphs.
*/
func (self Instance) SetKerning(cache_index int, size int, glyph_pair Vector2i.XY, kerning Vector2.XY) { //gd:FontFile.set_kerning
	class(self).SetKerning(gd.Int(cache_index), gd.Int(size), gd.Vector2i(glyph_pair), gd.Vector2(kerning))
}

/*
Returns kerning for the pair of glyphs.
*/
func (self Instance) GetKerning(cache_index int, size int, glyph_pair Vector2i.XY) Vector2.XY { //gd:FontFile.get_kerning
	return Vector2.XY(class(self).GetKerning(gd.Int(cache_index), gd.Int(size), gd.Vector2i(glyph_pair)))
}

/*
Renders the range of characters to the font cache texture.
*/
func (self Instance) RenderRange(cache_index int, size Vector2i.XY, start int, end int) { //gd:FontFile.render_range
	class(self).RenderRange(gd.Int(cache_index), gd.Vector2i(size), gd.Int(start), gd.Int(end))
}

/*
Renders specified glyph to the font cache texture.
*/
func (self Instance) RenderGlyph(cache_index int, size Vector2i.XY, index int) { //gd:FontFile.render_glyph
	class(self).RenderGlyph(gd.Int(cache_index), gd.Vector2i(size), gd.Int(index))
}

/*
Adds override for [method Font.is_language_supported].
*/
func (self Instance) SetLanguageSupportOverride(language string, supported bool) { //gd:FontFile.set_language_support_override
	class(self).SetLanguageSupportOverride(gd.NewString(language), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
func (self Instance) GetLanguageSupportOverride(language string) bool { //gd:FontFile.get_language_support_override
	return bool(class(self).GetLanguageSupportOverride(gd.NewString(language)))
}

/*
Remove language support override.
*/
func (self Instance) RemoveLanguageSupportOverride(language string) { //gd:FontFile.remove_language_support_override
	class(self).RemoveLanguageSupportOverride(gd.NewString(language))
}

/*
Returns list of language support overrides.
*/
func (self Instance) GetLanguageSupportOverrides() []string { //gd:FontFile.get_language_support_overrides
	return []string(class(self).GetLanguageSupportOverrides().Strings())
}

/*
Adds override for [method Font.is_script_supported].
*/
func (self Instance) SetScriptSupportOverride(script string, supported bool) { //gd:FontFile.set_script_support_override
	class(self).SetScriptSupportOverride(gd.NewString(script), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
func (self Instance) GetScriptSupportOverride(script string) bool { //gd:FontFile.get_script_support_override
	return bool(class(self).GetScriptSupportOverride(gd.NewString(script)))
}

/*
Removes script support override.
*/
func (self Instance) RemoveScriptSupportOverride(script string) { //gd:FontFile.remove_script_support_override
	class(self).RemoveScriptSupportOverride(gd.NewString(script))
}

/*
Returns list of script support overrides.
*/
func (self Instance) GetScriptSupportOverrides() []string { //gd:FontFile.get_script_support_overrides
	return []string(class(self).GetScriptSupportOverrides().Strings())
}

/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
func (self Instance) GetGlyphIndex(size int, char int, variation_selector int) int { //gd:FontFile.get_glyph_index
	return int(int(class(self).GetGlyphIndex(gd.Int(size), gd.Int(char), gd.Int(variation_selector))))
}

/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method get_glyph_index].
*/
func (self Instance) GetCharFromGlyphIndex(size int, glyph_index int) int { //gd:FontFile.get_char_from_glyph_index
	return int(int(class(self).GetCharFromGlyphIndex(gd.Int(size), gd.Int(glyph_index))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FontFile

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FontFile"))
	casted := Instance{*(*gdclass.FontFile)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Data() []byte {
	return []byte(class(self).GetData().Bytes())
}

func (self Instance) SetData(value []byte) {
	class(self).SetData(gd.NewPackedByteSlice(value))
}

func (self Instance) GenerateMipmaps() bool {
	return bool(class(self).GetGenerateMipmaps())
}

func (self Instance) SetGenerateMipmaps(value bool) {
	class(self).SetGenerateMipmaps(value)
}

func (self Instance) DisableEmbeddedBitmaps() bool {
	return bool(class(self).GetDisableEmbeddedBitmaps())
}

func (self Instance) SetDisableEmbeddedBitmaps(value bool) {
	class(self).SetDisableEmbeddedBitmaps(value)
}

func (self Instance) Antialiasing() gdclass.TextServerFontAntialiasing {
	return gdclass.TextServerFontAntialiasing(class(self).GetAntialiasing())
}

func (self Instance) SetAntialiasing(value gdclass.TextServerFontAntialiasing) {
	class(self).SetAntialiasing(value)
}

func (self Instance) SetFontName(value string) {
	class(self).SetFontName(gd.NewString(value))
}

func (self Instance) SetStyleName(value string) {
	class(self).SetFontStyleName(gd.NewString(value))
}

func (self Instance) SetFontStyle(value gdclass.TextServerFontStyle) {
	class(self).SetFontStyle(value)
}

func (self Instance) SetFontWeight(value int) {
	class(self).SetFontWeight(gd.Int(value))
}

func (self Instance) SetFontStretch(value int) {
	class(self).SetFontStretch(gd.Int(value))
}

func (self Instance) SubpixelPositioning() gdclass.TextServerSubpixelPositioning {
	return gdclass.TextServerSubpixelPositioning(class(self).GetSubpixelPositioning())
}

func (self Instance) SetSubpixelPositioning(value gdclass.TextServerSubpixelPositioning) {
	class(self).SetSubpixelPositioning(value)
}

func (self Instance) MultichannelSignedDistanceField() bool {
	return bool(class(self).IsMultichannelSignedDistanceField())
}

func (self Instance) SetMultichannelSignedDistanceField(value bool) {
	class(self).SetMultichannelSignedDistanceField(value)
}

func (self Instance) MsdfPixelRange() int {
	return int(int(class(self).GetMsdfPixelRange()))
}

func (self Instance) SetMsdfPixelRange(value int) {
	class(self).SetMsdfPixelRange(gd.Int(value))
}

func (self Instance) MsdfSize() int {
	return int(int(class(self).GetMsdfSize()))
}

func (self Instance) SetMsdfSize(value int) {
	class(self).SetMsdfSize(gd.Int(value))
}

func (self Instance) AllowSystemFallback() bool {
	return bool(class(self).IsAllowSystemFallback())
}

func (self Instance) SetAllowSystemFallback(value bool) {
	class(self).SetAllowSystemFallback(value)
}

func (self Instance) ForceAutohinter() bool {
	return bool(class(self).IsForceAutohinter())
}

func (self Instance) SetForceAutohinter(value bool) {
	class(self).SetForceAutohinter(value)
}

func (self Instance) Hinting() gdclass.TextServerHinting {
	return gdclass.TextServerHinting(class(self).GetHinting())
}

func (self Instance) SetHinting(value gdclass.TextServerHinting) {
	class(self).SetHinting(value)
}

func (self Instance) Oversampling() Float.X {
	return Float.X(Float.X(class(self).GetOversampling()))
}

func (self Instance) SetOversampling(value Float.X) {
	class(self).SetOversampling(gd.Float(value))
}

func (self Instance) FixedSize() int {
	return int(int(class(self).GetFixedSize()))
}

func (self Instance) SetFixedSize(value int) {
	class(self).SetFixedSize(gd.Int(value))
}

func (self Instance) FixedSizeScaleMode() gdclass.TextServerFixedSizeScaleMode {
	return gdclass.TextServerFixedSizeScaleMode(class(self).GetFixedSizeScaleMode())
}

func (self Instance) SetFixedSizeScaleMode(value gdclass.TextServerFixedSizeScaleMode) {
	class(self).SetFixedSizeScaleMode(value)
}

func (self Instance) OpentypeFeatureOverrides() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetOpentypeFeatureOverrides()))
}

func (self Instance) SetOpentypeFeatureOverrides(value map[any]any) {
	class(self).SetOpentypeFeatureOverrides(gd.DictionaryFromMap(value))
}

/*
Loads an AngelCode BMFont (.fnt, .font) bitmap font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
//go:nosplit
func (self class) LoadBitmapFont(path gd.String) gd.Error { //gd:FontFile.load_bitmap_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_load_bitmap_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads a TrueType (.ttf), OpenType (.otf), WOFF (.woff), WOFF2 (.woff2) or Type 1 (.pfb, .pfm) dynamic font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
//go:nosplit
func (self class) LoadDynamicFont(path gd.String) gd.Error { //gd:FontFile.load_dynamic_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_load_dynamic_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetData(data gd.PackedByteArray) { //gd:FontFile.set_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetData() gd.PackedByteArray { //gd:FontFile.get_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFontName(name gd.String) { //gd:FontFile.set_font_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontStyleName(name gd.String) { //gd:FontFile.set_font_style_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_style_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontStyle(style gdclass.TextServerFontStyle) { //gd:FontFile.set_font_style
	var frame = callframe.New()
	callframe.Arg(frame, style)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_style, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontWeight(weight gd.Int) { //gd:FontFile.set_font_weight
	var frame = callframe.New()
	callframe.Arg(frame, weight)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_weight, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontStretch(stretch gd.Int) { //gd:FontFile.set_font_stretch
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_stretch, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAntialiasing(antialiasing gdclass.TextServerFontAntialiasing) { //gd:FontFile.set_antialiasing
	var frame = callframe.New()
	callframe.Arg(frame, antialiasing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAntialiasing() gdclass.TextServerFontAntialiasing { //gd:FontFile.get_antialiasing
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerFontAntialiasing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableEmbeddedBitmaps(disable_embedded_bitmaps bool) { //gd:FontFile.set_disable_embedded_bitmaps
	var frame = callframe.New()
	callframe.Arg(frame, disable_embedded_bitmaps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisableEmbeddedBitmaps() bool { //gd:FontFile.get_disable_embedded_bitmaps
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGenerateMipmaps(generate_mipmaps bool) { //gd:FontFile.set_generate_mipmaps
	var frame = callframe.New()
	callframe.Arg(frame, generate_mipmaps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGenerateMipmaps() bool { //gd:FontFile.get_generate_mipmaps
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMultichannelSignedDistanceField(msdf bool) { //gd:FontFile.set_multichannel_signed_distance_field
	var frame = callframe.New()
	callframe.Arg(frame, msdf)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMultichannelSignedDistanceField() bool { //gd:FontFile.is_multichannel_signed_distance_field
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsdfPixelRange(msdf_pixel_range gd.Int) { //gd:FontFile.set_msdf_pixel_range
	var frame = callframe.New()
	callframe.Arg(frame, msdf_pixel_range)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsdfPixelRange() gd.Int { //gd:FontFile.get_msdf_pixel_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsdfSize(msdf_size gd.Int) { //gd:FontFile.set_msdf_size
	var frame = callframe.New()
	callframe.Arg(frame, msdf_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_msdf_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsdfSize() gd.Int { //gd:FontFile.get_msdf_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_msdf_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFixedSize(fixed_size gd.Int) { //gd:FontFile.set_fixed_size
	var frame = callframe.New()
	callframe.Arg(frame, fixed_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_fixed_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFixedSize() gd.Int { //gd:FontFile.get_fixed_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_fixed_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFixedSizeScaleMode(fixed_size_scale_mode gdclass.TextServerFixedSizeScaleMode) { //gd:FontFile.set_fixed_size_scale_mode
	var frame = callframe.New()
	callframe.Arg(frame, fixed_size_scale_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFixedSizeScaleMode() gdclass.TextServerFixedSizeScaleMode { //gd:FontFile.get_fixed_size_scale_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerFixedSizeScaleMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowSystemFallback(allow_system_fallback bool) { //gd:FontFile.set_allow_system_fallback
	var frame = callframe.New()
	callframe.Arg(frame, allow_system_fallback)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllowSystemFallback() bool { //gd:FontFile.is_allow_system_fallback
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetForceAutohinter(force_autohinter bool) { //gd:FontFile.set_force_autohinter
	var frame = callframe.New()
	callframe.Arg(frame, force_autohinter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsForceAutohinter() bool { //gd:FontFile.is_force_autohinter
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHinting(hinting gdclass.TextServerHinting) { //gd:FontFile.set_hinting
	var frame = callframe.New()
	callframe.Arg(frame, hinting)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_hinting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHinting() gdclass.TextServerHinting { //gd:FontFile.get_hinting
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerHinting](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_hinting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubpixelPositioning(subpixel_positioning gdclass.TextServerSubpixelPositioning) { //gd:FontFile.set_subpixel_positioning
	var frame = callframe.New()
	callframe.Arg(frame, subpixel_positioning)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubpixelPositioning() gdclass.TextServerSubpixelPositioning { //gd:FontFile.get_subpixel_positioning
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerSubpixelPositioning](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOversampling(oversampling gd.Float) { //gd:FontFile.set_oversampling
	var frame = callframe.New()
	callframe.Arg(frame, oversampling)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOversampling() gd.Float { //gd:FontFile.get_oversampling
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of the font cache entries.
*/
//go:nosplit
func (self class) GetCacheCount() gd.Int { //gd:FontFile.get_cache_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes all font cache entries.
*/
//go:nosplit
func (self class) ClearCache() { //gd:FontFile.clear_cache
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified font cache entry.
*/
//go:nosplit
func (self class) RemoveCache(cache_index gd.Int) { //gd:FontFile.remove_cache
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
//go:nosplit
func (self class) GetSizeCacheList(cache_index gd.Int) Array.Contains[gd.Vector2i] { //gd:FontFile.get_size_cache_list
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_size_cache_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Vector2i]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes all font sizes from the cache entry
*/
//go:nosplit
func (self class) ClearSizeCache(cache_index gd.Int) { //gd:FontFile.clear_size_cache
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_size_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified font size from the cache entry.
*/
//go:nosplit
func (self class) RemoveSizeCache(cache_index gd.Int, size gd.Vector2i) { //gd:FontFile.remove_size_cache
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_size_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) SetVariationCoordinates(cache_index gd.Int, variation_coordinates Dictionary.Any) { //gd:FontFile.set_variation_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(variation_coordinates)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) GetVariationCoordinates(cache_index gd.Int) Dictionary.Any { //gd:FontFile.get_variation_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) SetEmbolden(cache_index gd.Int, strength gd.Float) { //gd:FontFile.set_embolden
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_embolden, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) GetEmbolden(cache_index gd.Int) gd.Float { //gd:FontFile.get_embolden
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_embolden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
//go:nosplit
func (self class) SetTransform(cache_index gd.Int, transform gd.Transform2D) { //gd:FontFile.set_transform
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns 2D transform, applied to the font outlines, can be used for slanting, flipping and rotating glyphs.
*/
//go:nosplit
func (self class) GetTransform(cache_index gd.Int) gd.Transform2D { //gd:FontFile.get_transform
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) SetExtraSpacing(cache_index gd.Int, spacing gdclass.TextServerSpacingType, value gd.Int) { //gd:FontFile.set_extra_spacing
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_extra_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) GetExtraSpacing(cache_index gd.Int, spacing gdclass.TextServerSpacingType) gd.Int { //gd:FontFile.get_extra_spacing
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_extra_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) SetExtraBaselineOffset(cache_index gd.Int, baseline_offset gd.Float) { //gd:FontFile.set_extra_baseline_offset
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, baseline_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_extra_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) GetExtraBaselineOffset(cache_index gd.Int) gd.Float { //gd:FontFile.get_extra_baseline_offset
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_extra_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) SetFaceIndex(cache_index gd.Int, face_index gd.Int) { //gd:FontFile.set_face_index
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, face_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_face_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) GetFaceIndex(cache_index gd.Int) gd.Int { //gd:FontFile.get_face_index
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_face_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) SetCacheAscent(cache_index gd.Int, size gd.Int, ascent gd.Float) { //gd:FontFile.set_cache_ascent
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, ascent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) GetCacheAscent(cache_index gd.Int, size gd.Int) gd.Float { //gd:FontFile.get_cache_ascent
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) SetCacheDescent(cache_index gd.Int, size gd.Int, descent gd.Float) { //gd:FontFile.set_cache_descent
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, descent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) GetCacheDescent(cache_index gd.Int, size gd.Int) gd.Float { //gd:FontFile.get_cache_descent
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) SetCacheUnderlinePosition(cache_index gd.Int, size gd.Int, underline_position gd.Float) { //gd:FontFile.set_cache_underline_position
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetCacheUnderlinePosition(cache_index gd.Int, size gd.Int) gd.Float { //gd:FontFile.get_cache_underline_position
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets thickness of the underline in pixels.
*/
//go:nosplit
func (self class) SetCacheUnderlineThickness(cache_index gd.Int, size gd.Int, underline_thickness gd.Float) { //gd:FontFile.set_cache_underline_thickness
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_thickness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns thickness of the underline in pixels.
*/
//go:nosplit
func (self class) GetCacheUnderlineThickness(cache_index gd.Int, size gd.Int) gd.Float { //gd:FontFile.get_cache_underline_thickness
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) SetCacheScale(cache_index gd.Int, size gd.Int, scale gd.Float) { //gd:FontFile.set_cache_scale
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) GetCacheScale(cache_index gd.Int, size gd.Int) gd.Float { //gd:FontFile.get_cache_scale
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of textures used by font cache entry.
*/
//go:nosplit
func (self class) GetTextureCount(cache_index gd.Int, size gd.Vector2i) gd.Int { //gd:FontFile.get_texture_count
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method remove_glyph] to remove them manually.
*/
//go:nosplit
func (self class) ClearTextures(cache_index gd.Int, size gd.Vector2i) { //gd:FontFile.clear_textures
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_textures, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture. Remove them manually using [method remove_glyph].
*/
//go:nosplit
func (self class) RemoveTexture(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int) { //gd:FontFile.remove_texture
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets font cache texture image.
*/
//go:nosplit
func (self class) SetTextureImage(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int, image [1]gdclass.Image) { //gd:FontFile.set_texture_image
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_texture_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a copy of the font cache texture image.
*/
//go:nosplit
func (self class) GetTextureImage(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int) [1]gdclass.Image { //gd:FontFile.get_texture_image
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets array containing glyph packing data.
*/
//go:nosplit
func (self class) SetTextureOffsets(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int, offset gd.PackedInt32Array) { //gd:FontFile.set_texture_offsets
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, pointers.Get(offset))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a copy of the array containing glyph packing data.
*/
//go:nosplit
func (self class) GetTextureOffsets(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int) gd.PackedInt32Array { //gd:FontFile.get_texture_offsets
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns list of rendered glyphs in the cache entry.
*/
//go:nosplit
func (self class) GetGlyphList(cache_index gd.Int, size gd.Vector2i) gd.PackedInt32Array { //gd:FontFile.get_glyph_list
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) ClearGlyphs(cache_index gd.Int, size gd.Vector2i) { //gd:FontFile.clear_glyphs
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_glyphs, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) RemoveGlyph(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) { //gd:FontFile.remove_glyph
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_glyph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) SetGlyphAdvance(cache_index gd.Int, size gd.Int, glyph gd.Int, advance gd.Vector2) { //gd:FontFile.set_glyph_advance
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, advance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) GetGlyphAdvance(cache_index gd.Int, size gd.Int, glyph gd.Int) gd.Vector2 { //gd:FontFile.get_glyph_advance
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets glyph offset from the baseline.
*/
//go:nosplit
func (self class) SetGlyphOffset(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, offset gd.Vector2) { //gd:FontFile.set_glyph_offset
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns glyph offset from the baseline.
*/
//go:nosplit
func (self class) GetGlyphOffset(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Vector2 { //gd:FontFile.get_glyph_offset
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets glyph size.
*/
//go:nosplit
func (self class) SetGlyphSize(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, gl_size gd.Vector2) { //gd:FontFile.set_glyph_size
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, gl_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns glyph size.
*/
//go:nosplit
func (self class) GetGlyphSize(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Vector2 { //gd:FontFile.get_glyph_size
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) SetGlyphUvRect(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, uv_rect gd.Rect2) { //gd:FontFile.set_glyph_uv_rect
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, uv_rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) GetGlyphUvRect(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Rect2 { //gd:FontFile.get_glyph_uv_rect
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) SetGlyphTextureIdx(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, texture_idx gd.Int) { //gd:FontFile.set_glyph_texture_idx
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, texture_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) GetGlyphTextureIdx(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Int { //gd:FontFile.get_glyph_texture_idx
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns list of the kerning overrides.
*/
//go:nosplit
func (self class) GetKerningList(cache_index gd.Int, size gd.Int) Array.Contains[gd.Vector2i] { //gd:FontFile.get_kerning_list
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_kerning_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Vector2i]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes all kerning overrides.
*/
//go:nosplit
func (self class) ClearKerningMap(cache_index gd.Int, size gd.Int) { //gd:FontFile.clear_kerning_map
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_kerning_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes kerning override for the pair of glyphs.
*/
//go:nosplit
func (self class) RemoveKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i) { //gd:FontFile.remove_kerning
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) SetKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i, kerning gd.Vector2) { //gd:FontFile.set_kerning
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	callframe.Arg(frame, kerning)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) GetKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i) gd.Vector2 { //gd:FontFile.get_kerning
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renders the range of characters to the font cache texture.
*/
//go:nosplit
func (self class) RenderRange(cache_index gd.Int, size gd.Vector2i, start gd.Int, end gd.Int) { //gd:FontFile.render_range
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_render_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Renders specified glyph to the font cache texture.
*/
//go:nosplit
func (self class) RenderGlyph(cache_index gd.Int, size gd.Vector2i, index gd.Int) { //gd:FontFile.render_glyph
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_render_glyph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds override for [method Font.is_language_supported].
*/
//go:nosplit
func (self class) SetLanguageSupportOverride(language gd.String, supported bool) { //gd:FontFile.set_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	callframe.Arg(frame, supported)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
//go:nosplit
func (self class) GetLanguageSupportOverride(language gd.String) bool { //gd:FontFile.get_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Remove language support override.
*/
//go:nosplit
func (self class) RemoveLanguageSupportOverride(language gd.String) { //gd:FontFile.remove_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns list of language support overrides.
*/
//go:nosplit
func (self class) GetLanguageSupportOverrides() gd.PackedStringArray { //gd:FontFile.get_language_support_overrides
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_language_support_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds override for [method Font.is_script_supported].
*/
//go:nosplit
func (self class) SetScriptSupportOverride(script gd.String, supported bool) { //gd:FontFile.set_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script))
	callframe.Arg(frame, supported)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
//go:nosplit
func (self class) GetScriptSupportOverride(script gd.String) bool { //gd:FontFile.get_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes script support override.
*/
//go:nosplit
func (self class) RemoveScriptSupportOverride(script gd.String) { //gd:FontFile.remove_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns list of script support overrides.
*/
//go:nosplit
func (self class) GetScriptSupportOverrides() gd.PackedStringArray { //gd:FontFile.get_script_support_overrides
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_script_support_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOpentypeFeatureOverrides(overrides Dictionary.Any) { //gd:FontFile.set_opentype_feature_overrides
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(overrides)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOpentypeFeatureOverrides() Dictionary.Any { //gd:FontFile.get_opentype_feature_overrides
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
//go:nosplit
func (self class) GetGlyphIndex(size gd.Int, char gd.Int, variation_selector gd.Int) gd.Int { //gd:FontFile.get_glyph_index
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, char)
	callframe.Arg(frame, variation_selector)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method get_glyph_index].
*/
//go:nosplit
func (self class) GetCharFromGlyphIndex(size gd.Int, glyph_index gd.Int) gd.Int { //gd:FontFile.get_char_from_glyph_index
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_char_from_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFontFile() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFontFile() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsFont() Font.Advanced    { return *((*Font.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFont() Font.Instance { return *((*Font.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Font.Advanced(self.AsFont()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Font.Instance(self.AsFont()), name)
	}
}
func init() {
	gdclass.Register("FontFile", func(ptr gd.Object) any { return [1]gdclass.FontFile{*(*gdclass.FontFile)(unsafe.Pointer(&ptr))} })
}

type Error = gd.Error //gd:Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
