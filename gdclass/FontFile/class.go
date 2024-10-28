package FontFile

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Font"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

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
type Go [1]classdb.FontFile

/*
Loads an AngelCode BMFont (.fnt, .font) bitmap font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
func (self Go) LoadBitmapFont(path string) gd.Error {
	return gd.Error(class(self).LoadBitmapFont(gd.NewString(path)))
}

/*
Loads a TrueType (.ttf), OpenType (.otf), WOFF (.woff), WOFF2 (.woff2) or Type 1 (.pfb, .pfm) dynamic font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
func (self Go) LoadDynamicFont(path string) gd.Error {
	return gd.Error(class(self).LoadDynamicFont(gd.NewString(path)))
}

/*
Returns number of the font cache entries.
*/
func (self Go) GetCacheCount() int {
	return int(int(class(self).GetCacheCount()))
}

/*
Removes all font cache entries.
*/
func (self Go) ClearCache() {
	class(self).ClearCache()
}

/*
Removes specified font cache entry.
*/
func (self Go) RemoveCache(cache_index int) {
	class(self).RemoveCache(gd.Int(cache_index))
}

/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
func (self Go) GetSizeCacheList(cache_index int) gd.Array {
	return gd.Array(class(self).GetSizeCacheList(gd.Int(cache_index)))
}

/*
Removes all font sizes from the cache entry
*/
func (self Go) ClearSizeCache(cache_index int) {
	class(self).ClearSizeCache(gd.Int(cache_index))
}

/*
Removes specified font size from the cache entry.
*/
func (self Go) RemoveSizeCache(cache_index int, size gd.Vector2i) {
	class(self).RemoveSizeCache(gd.Int(cache_index), size)
}

/*
Sets variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
func (self Go) SetVariationCoordinates(cache_index int, variation_coordinates gd.Dictionary) {
	class(self).SetVariationCoordinates(gd.Int(cache_index), variation_coordinates)
}

/*
Returns variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
func (self Go) GetVariationCoordinates(cache_index int) gd.Dictionary {
	return gd.Dictionary(class(self).GetVariationCoordinates(gd.Int(cache_index)))
}

/*
Sets embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (self Go) SetEmbolden(cache_index int, strength float64) {
	class(self).SetEmbolden(gd.Int(cache_index), gd.Float(strength))
}

/*
Returns embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (self Go) GetEmbolden(cache_index int) float64 {
	return float64(float64(class(self).GetEmbolden(gd.Int(cache_index))))
}

/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
func (self Go) SetTransform(cache_index int, transform gd.Transform2D) {
	class(self).SetTransform(gd.Int(cache_index), transform)
}

/*
Returns 2D transform, applied to the font outlines, can be used for slanting, flipping and rotating glyphs.
*/
func (self Go) GetTransform(cache_index int) gd.Transform2D {
	return gd.Transform2D(class(self).GetTransform(gd.Int(cache_index)))
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
func (self Go) SetExtraSpacing(cache_index int, spacing classdb.TextServerSpacingType, value int) {
	class(self).SetExtraSpacing(gd.Int(cache_index), spacing, gd.Int(value))
}

/*
Returns spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
func (self Go) GetExtraSpacing(cache_index int, spacing classdb.TextServerSpacingType) int {
	return int(int(class(self).GetExtraSpacing(gd.Int(cache_index), spacing)))
}

/*
Sets extra baseline offset (as a fraction of font height).
*/
func (self Go) SetExtraBaselineOffset(cache_index int, baseline_offset float64) {
	class(self).SetExtraBaselineOffset(gd.Int(cache_index), gd.Float(baseline_offset))
}

/*
Returns extra baseline offset (as a fraction of font height).
*/
func (self Go) GetExtraBaselineOffset(cache_index int) float64 {
	return float64(float64(class(self).GetExtraBaselineOffset(gd.Int(cache_index))))
}

/*
Sets an active face index in the TrueType / OpenType collection.
*/
func (self Go) SetFaceIndex(cache_index int, face_index int) {
	class(self).SetFaceIndex(gd.Int(cache_index), gd.Int(face_index))
}

/*
Returns an active face index in the TrueType / OpenType collection.
*/
func (self Go) GetFaceIndex(cache_index int) int {
	return int(int(class(self).GetFaceIndex(gd.Int(cache_index))))
}

/*
Sets the font ascent (number of pixels above the baseline).
*/
func (self Go) SetCacheAscent(cache_index int, size int, ascent float64) {
	class(self).SetCacheAscent(gd.Int(cache_index), gd.Int(size), gd.Float(ascent))
}

/*
Returns the font ascent (number of pixels above the baseline).
*/
func (self Go) GetCacheAscent(cache_index int, size int) float64 {
	return float64(float64(class(self).GetCacheAscent(gd.Int(cache_index), gd.Int(size))))
}

/*
Sets the font descent (number of pixels below the baseline).
*/
func (self Go) SetCacheDescent(cache_index int, size int, descent float64) {
	class(self).SetCacheDescent(gd.Int(cache_index), gd.Int(size), gd.Float(descent))
}

/*
Returns the font descent (number of pixels below the baseline).
*/
func (self Go) GetCacheDescent(cache_index int, size int) float64 {
	return float64(float64(class(self).GetCacheDescent(gd.Int(cache_index), gd.Int(size))))
}

/*
Sets pixel offset of the underline below the baseline.
*/
func (self Go) SetCacheUnderlinePosition(cache_index int, size int, underline_position float64) {
	class(self).SetCacheUnderlinePosition(gd.Int(cache_index), gd.Int(size), gd.Float(underline_position))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Go) GetCacheUnderlinePosition(cache_index int, size int) float64 {
	return float64(float64(class(self).GetCacheUnderlinePosition(gd.Int(cache_index), gd.Int(size))))
}

/*
Sets thickness of the underline in pixels.
*/
func (self Go) SetCacheUnderlineThickness(cache_index int, size int, underline_thickness float64) {
	class(self).SetCacheUnderlineThickness(gd.Int(cache_index), gd.Int(size), gd.Float(underline_thickness))
}

/*
Returns thickness of the underline in pixels.
*/
func (self Go) GetCacheUnderlineThickness(cache_index int, size int) float64 {
	return float64(float64(class(self).GetCacheUnderlineThickness(gd.Int(cache_index), gd.Int(size))))
}

/*
Sets scaling factor of the color bitmap font.
*/
func (self Go) SetCacheScale(cache_index int, size int, scale float64) {
	class(self).SetCacheScale(gd.Int(cache_index), gd.Int(size), gd.Float(scale))
}

/*
Returns scaling factor of the color bitmap font.
*/
func (self Go) GetCacheScale(cache_index int, size int) float64 {
	return float64(float64(class(self).GetCacheScale(gd.Int(cache_index), gd.Int(size))))
}

/*
Returns number of textures used by font cache entry.
*/
func (self Go) GetTextureCount(cache_index int, size gd.Vector2i) int {
	return int(int(class(self).GetTextureCount(gd.Int(cache_index), size)))
}

/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method remove_glyph] to remove them manually.
*/
func (self Go) ClearTextures(cache_index int, size gd.Vector2i) {
	class(self).ClearTextures(gd.Int(cache_index), size)
}

/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture. Remove them manually using [method remove_glyph].
*/
func (self Go) RemoveTexture(cache_index int, size gd.Vector2i, texture_index int) {
	class(self).RemoveTexture(gd.Int(cache_index), size, gd.Int(texture_index))
}

/*
Sets font cache texture image.
*/
func (self Go) SetTextureImage(cache_index int, size gd.Vector2i, texture_index int, image gdclass.Image) {
	class(self).SetTextureImage(gd.Int(cache_index), size, gd.Int(texture_index), image)
}

/*
Returns a copy of the font cache texture image.
*/
func (self Go) GetTextureImage(cache_index int, size gd.Vector2i, texture_index int) gdclass.Image {
	return gdclass.Image(class(self).GetTextureImage(gd.Int(cache_index), size, gd.Int(texture_index)))
}

/*
Sets array containing glyph packing data.
*/
func (self Go) SetTextureOffsets(cache_index int, size gd.Vector2i, texture_index int, offset []int32) {
	class(self).SetTextureOffsets(gd.Int(cache_index), size, gd.Int(texture_index), gd.NewPackedInt32Slice(offset))
}

/*
Returns a copy of the array containing glyph packing data.
*/
func (self Go) GetTextureOffsets(cache_index int, size gd.Vector2i, texture_index int) []int32 {
	return []int32(class(self).GetTextureOffsets(gd.Int(cache_index), size, gd.Int(texture_index)).AsSlice())
}

/*
Returns list of rendered glyphs in the cache entry.
*/
func (self Go) GetGlyphList(cache_index int, size gd.Vector2i) []int32 {
	return []int32(class(self).GetGlyphList(gd.Int(cache_index), size).AsSlice())
}

/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
func (self Go) ClearGlyphs(cache_index int, size gd.Vector2i) {
	class(self).ClearGlyphs(gd.Int(cache_index), size)
}

/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
func (self Go) RemoveGlyph(cache_index int, size gd.Vector2i, glyph int) {
	class(self).RemoveGlyph(gd.Int(cache_index), size, gd.Int(glyph))
}

/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Go) SetGlyphAdvance(cache_index int, size int, glyph int, advance gd.Vector2) {
	class(self).SetGlyphAdvance(gd.Int(cache_index), gd.Int(size), gd.Int(glyph), advance)
}

/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Go) GetGlyphAdvance(cache_index int, size int, glyph int) gd.Vector2 {
	return gd.Vector2(class(self).GetGlyphAdvance(gd.Int(cache_index), gd.Int(size), gd.Int(glyph)))
}

/*
Sets glyph offset from the baseline.
*/
func (self Go) SetGlyphOffset(cache_index int, size gd.Vector2i, glyph int, offset gd.Vector2) {
	class(self).SetGlyphOffset(gd.Int(cache_index), size, gd.Int(glyph), offset)
}

/*
Returns glyph offset from the baseline.
*/
func (self Go) GetGlyphOffset(cache_index int, size gd.Vector2i, glyph int) gd.Vector2 {
	return gd.Vector2(class(self).GetGlyphOffset(gd.Int(cache_index), size, gd.Int(glyph)))
}

/*
Sets glyph size.
*/
func (self Go) SetGlyphSize(cache_index int, size gd.Vector2i, glyph int, gl_size gd.Vector2) {
	class(self).SetGlyphSize(gd.Int(cache_index), size, gd.Int(glyph), gl_size)
}

/*
Returns glyph size.
*/
func (self Go) GetGlyphSize(cache_index int, size gd.Vector2i, glyph int) gd.Vector2 {
	return gd.Vector2(class(self).GetGlyphSize(gd.Int(cache_index), size, gd.Int(glyph)))
}

/*
Sets rectangle in the cache texture containing the glyph.
*/
func (self Go) SetGlyphUvRect(cache_index int, size gd.Vector2i, glyph int, uv_rect gd.Rect2) {
	class(self).SetGlyphUvRect(gd.Int(cache_index), size, gd.Int(glyph), uv_rect)
}

/*
Returns rectangle in the cache texture containing the glyph.
*/
func (self Go) GetGlyphUvRect(cache_index int, size gd.Vector2i, glyph int) gd.Rect2 {
	return gd.Rect2(class(self).GetGlyphUvRect(gd.Int(cache_index), size, gd.Int(glyph)))
}

/*
Sets index of the cache texture containing the glyph.
*/
func (self Go) SetGlyphTextureIdx(cache_index int, size gd.Vector2i, glyph int, texture_idx int) {
	class(self).SetGlyphTextureIdx(gd.Int(cache_index), size, gd.Int(glyph), gd.Int(texture_idx))
}

/*
Returns index of the cache texture containing the glyph.
*/
func (self Go) GetGlyphTextureIdx(cache_index int, size gd.Vector2i, glyph int) int {
	return int(int(class(self).GetGlyphTextureIdx(gd.Int(cache_index), size, gd.Int(glyph))))
}

/*
Returns list of the kerning overrides.
*/
func (self Go) GetKerningList(cache_index int, size int) gd.Array {
	return gd.Array(class(self).GetKerningList(gd.Int(cache_index), gd.Int(size)))
}

/*
Removes all kerning overrides.
*/
func (self Go) ClearKerningMap(cache_index int, size int) {
	class(self).ClearKerningMap(gd.Int(cache_index), gd.Int(size))
}

/*
Removes kerning override for the pair of glyphs.
*/
func (self Go) RemoveKerning(cache_index int, size int, glyph_pair gd.Vector2i) {
	class(self).RemoveKerning(gd.Int(cache_index), gd.Int(size), glyph_pair)
}

/*
Sets kerning for the pair of glyphs.
*/
func (self Go) SetKerning(cache_index int, size int, glyph_pair gd.Vector2i, kerning gd.Vector2) {
	class(self).SetKerning(gd.Int(cache_index), gd.Int(size), glyph_pair, kerning)
}

/*
Returns kerning for the pair of glyphs.
*/
func (self Go) GetKerning(cache_index int, size int, glyph_pair gd.Vector2i) gd.Vector2 {
	return gd.Vector2(class(self).GetKerning(gd.Int(cache_index), gd.Int(size), glyph_pair))
}

/*
Renders the range of characters to the font cache texture.
*/
func (self Go) RenderRange(cache_index int, size gd.Vector2i, start int, end int) {
	class(self).RenderRange(gd.Int(cache_index), size, gd.Int(start), gd.Int(end))
}

/*
Renders specified glyph to the font cache texture.
*/
func (self Go) RenderGlyph(cache_index int, size gd.Vector2i, index int) {
	class(self).RenderGlyph(gd.Int(cache_index), size, gd.Int(index))
}

/*
Adds override for [method Font.is_language_supported].
*/
func (self Go) SetLanguageSupportOverride(language string, supported bool) {
	class(self).SetLanguageSupportOverride(gd.NewString(language), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
func (self Go) GetLanguageSupportOverride(language string) bool {
	return bool(class(self).GetLanguageSupportOverride(gd.NewString(language)))
}

/*
Remove language support override.
*/
func (self Go) RemoveLanguageSupportOverride(language string) {
	class(self).RemoveLanguageSupportOverride(gd.NewString(language))
}

/*
Returns list of language support overrides.
*/
func (self Go) GetLanguageSupportOverrides() []string {
	return []string(class(self).GetLanguageSupportOverrides().Strings())
}

/*
Adds override for [method Font.is_script_supported].
*/
func (self Go) SetScriptSupportOverride(script string, supported bool) {
	class(self).SetScriptSupportOverride(gd.NewString(script), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
func (self Go) GetScriptSupportOverride(script string) bool {
	return bool(class(self).GetScriptSupportOverride(gd.NewString(script)))
}

/*
Removes script support override.
*/
func (self Go) RemoveScriptSupportOverride(script string) {
	class(self).RemoveScriptSupportOverride(gd.NewString(script))
}

/*
Returns list of script support overrides.
*/
func (self Go) GetScriptSupportOverrides() []string {
	return []string(class(self).GetScriptSupportOverrides().Strings())
}

/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
func (self Go) GetGlyphIndex(size int, char int, variation_selector int) int {
	return int(int(class(self).GetGlyphIndex(gd.Int(size), gd.Int(char), gd.Int(variation_selector))))
}

/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method get_glyph_index].
*/
func (self Go) GetCharFromGlyphIndex(size int, glyph_index int) int {
	return int(int(class(self).GetCharFromGlyphIndex(gd.Int(size), gd.Int(glyph_index))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.FontFile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FontFile"))
	return Go{classdb.FontFile(object)}
}

func (self Go) Data() []byte {
		return []byte(class(self).GetData().Bytes())
}

func (self Go) SetData(value []byte) {
	class(self).SetData(gd.NewPackedByteSlice(value))
}

func (self Go) GenerateMipmaps() bool {
		return bool(class(self).GetGenerateMipmaps())
}

func (self Go) SetGenerateMipmaps(value bool) {
	class(self).SetGenerateMipmaps(value)
}

func (self Go) DisableEmbeddedBitmaps() bool {
		return bool(class(self).GetDisableEmbeddedBitmaps())
}

func (self Go) SetDisableEmbeddedBitmaps(value bool) {
	class(self).SetDisableEmbeddedBitmaps(value)
}

func (self Go) Antialiasing() classdb.TextServerFontAntialiasing {
		return classdb.TextServerFontAntialiasing(class(self).GetAntialiasing())
}

func (self Go) SetAntialiasing(value classdb.TextServerFontAntialiasing) {
	class(self).SetAntialiasing(value)
}

func (self Go) SubpixelPositioning() classdb.TextServerSubpixelPositioning {
		return classdb.TextServerSubpixelPositioning(class(self).GetSubpixelPositioning())
}

func (self Go) SetSubpixelPositioning(value classdb.TextServerSubpixelPositioning) {
	class(self).SetSubpixelPositioning(value)
}

func (self Go) MultichannelSignedDistanceField() bool {
		return bool(class(self).IsMultichannelSignedDistanceField())
}

func (self Go) SetMultichannelSignedDistanceField(value bool) {
	class(self).SetMultichannelSignedDistanceField(value)
}

func (self Go) MsdfPixelRange() int {
		return int(int(class(self).GetMsdfPixelRange()))
}

func (self Go) SetMsdfPixelRange(value int) {
	class(self).SetMsdfPixelRange(gd.Int(value))
}

func (self Go) MsdfSize() int {
		return int(int(class(self).GetMsdfSize()))
}

func (self Go) SetMsdfSize(value int) {
	class(self).SetMsdfSize(gd.Int(value))
}

func (self Go) AllowSystemFallback() bool {
		return bool(class(self).IsAllowSystemFallback())
}

func (self Go) SetAllowSystemFallback(value bool) {
	class(self).SetAllowSystemFallback(value)
}

func (self Go) ForceAutohinter() bool {
		return bool(class(self).IsForceAutohinter())
}

func (self Go) SetForceAutohinter(value bool) {
	class(self).SetForceAutohinter(value)
}

func (self Go) Hinting() classdb.TextServerHinting {
		return classdb.TextServerHinting(class(self).GetHinting())
}

func (self Go) SetHinting(value classdb.TextServerHinting) {
	class(self).SetHinting(value)
}

func (self Go) Oversampling() float64 {
		return float64(float64(class(self).GetOversampling()))
}

func (self Go) SetOversampling(value float64) {
	class(self).SetOversampling(gd.Float(value))
}

func (self Go) FixedSize() int {
		return int(int(class(self).GetFixedSize()))
}

func (self Go) SetFixedSize(value int) {
	class(self).SetFixedSize(gd.Int(value))
}

func (self Go) FixedSizeScaleMode() classdb.TextServerFixedSizeScaleMode {
		return classdb.TextServerFixedSizeScaleMode(class(self).GetFixedSizeScaleMode())
}

func (self Go) SetFixedSizeScaleMode(value classdb.TextServerFixedSizeScaleMode) {
	class(self).SetFixedSizeScaleMode(value)
}

func (self Go) OpentypeFeatureOverrides() gd.Dictionary {
		return gd.Dictionary(class(self).GetOpentypeFeatureOverrides())
}

func (self Go) SetOpentypeFeatureOverrides(value gd.Dictionary) {
	class(self).SetOpentypeFeatureOverrides(value)
}

/*
Loads an AngelCode BMFont (.fnt, .font) bitmap font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
//go:nosplit
func (self class) LoadBitmapFont(path gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_load_bitmap_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Loads a TrueType (.ttf), OpenType (.otf), WOFF (.woff), WOFF2 (.woff2) or Type 1 (.pfb, .pfm) dynamic font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
//go:nosplit
func (self class) LoadDynamicFont(path gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_load_dynamic_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetData(data gd.PackedByteArray)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetData() gd.PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFontName(name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontStyleName(name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_style_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontStyle(style classdb.TextServerFontStyle)  {
	var frame = callframe.New()
	callframe.Arg(frame, style)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_style, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontWeight(weight gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontStretch(stretch gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAntialiasing(antialiasing classdb.TextServerFontAntialiasing)  {
	var frame = callframe.New()
	callframe.Arg(frame, antialiasing)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAntialiasing() classdb.TextServerFontAntialiasing {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerFontAntialiasing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisableEmbeddedBitmaps(disable_embedded_bitmaps bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, disable_embedded_bitmaps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDisableEmbeddedBitmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGenerateMipmaps(generate_mipmaps bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, generate_mipmaps)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGenerateMipmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMultichannelSignedDistanceField(msdf bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, msdf)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMultichannelSignedDistanceField() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsdfPixelRange(msdf_pixel_range gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, msdf_pixel_range)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsdfPixelRange() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsdfSize(msdf_size gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, msdf_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_msdf_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsdfSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_msdf_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFixedSize(fixed_size gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, fixed_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_fixed_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFixedSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_fixed_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFixedSizeScaleMode(fixed_size_scale_mode classdb.TextServerFixedSizeScaleMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, fixed_size_scale_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFixedSizeScaleMode() classdb.TextServerFixedSizeScaleMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerFixedSizeScaleMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowSystemFallback(allow_system_fallback bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, allow_system_fallback)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAllowSystemFallback() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetForceAutohinter(force_autohinter bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, force_autohinter)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsForceAutohinter() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHinting(hinting classdb.TextServerHinting)  {
	var frame = callframe.New()
	callframe.Arg(frame, hinting)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_hinting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHinting() classdb.TextServerHinting {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerHinting](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_hinting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubpixelPositioning(subpixel_positioning classdb.TextServerSubpixelPositioning)  {
	var frame = callframe.New()
	callframe.Arg(frame, subpixel_positioning)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubpixelPositioning() classdb.TextServerSubpixelPositioning {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerSubpixelPositioning](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOversampling(oversampling gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, oversampling)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOversampling() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of the font cache entries.
*/
//go:nosplit
func (self class) GetCacheCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all font cache entries.
*/
//go:nosplit
func (self class) ClearCache()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified font cache entry.
*/
//go:nosplit
func (self class) RemoveCache(cache_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
//go:nosplit
func (self class) GetSizeCacheList(cache_index gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_size_cache_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes all font sizes from the cache entry
*/
//go:nosplit
func (self class) ClearSizeCache(cache_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_size_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified font size from the cache entry.
*/
//go:nosplit
func (self class) RemoveSizeCache(cache_index gd.Int, size gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_size_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) SetVariationCoordinates(cache_index gd.Int, variation_coordinates gd.Dictionary)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, discreet.Get(variation_coordinates))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) GetVariationCoordinates(cache_index gd.Int) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) SetEmbolden(cache_index gd.Int, strength gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) GetEmbolden(cache_index gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
//go:nosplit
func (self class) SetTransform(cache_index gd.Int, transform gd.Transform2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns 2D transform, applied to the font outlines, can be used for slanting, flipping and rotating glyphs.
*/
//go:nosplit
func (self class) GetTransform(cache_index gd.Int) gd.Transform2D {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) SetExtraSpacing(cache_index gd.Int, spacing classdb.TextServerSpacingType, value gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_extra_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) GetExtraSpacing(cache_index gd.Int, spacing classdb.TextServerSpacingType) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_extra_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) SetExtraBaselineOffset(cache_index gd.Int, baseline_offset gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, baseline_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_extra_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) GetExtraBaselineOffset(cache_index gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_extra_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) SetFaceIndex(cache_index gd.Int, face_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, face_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) GetFaceIndex(cache_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) SetCacheAscent(cache_index gd.Int, size gd.Int, ascent gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, ascent)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) GetCacheAscent(cache_index gd.Int, size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) SetCacheDescent(cache_index gd.Int, size gd.Int, descent gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, descent)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) GetCacheDescent(cache_index gd.Int, size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) SetCacheUnderlinePosition(cache_index gd.Int, size gd.Int, underline_position gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetCacheUnderlinePosition(cache_index gd.Int, size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets thickness of the underline in pixels.
*/
//go:nosplit
func (self class) SetCacheUnderlineThickness(cache_index gd.Int, size gd.Int, underline_thickness gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_thickness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns thickness of the underline in pixels.
*/
//go:nosplit
func (self class) GetCacheUnderlineThickness(cache_index gd.Int, size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) SetCacheScale(cache_index gd.Int, size gd.Int, scale gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) GetCacheScale(cache_index gd.Int, size gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of textures used by font cache entry.
*/
//go:nosplit
func (self class) GetTextureCount(cache_index gd.Int, size gd.Vector2i) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method remove_glyph] to remove them manually.
*/
//go:nosplit
func (self class) ClearTextures(cache_index gd.Int, size gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture. Remove them manually using [method remove_glyph].
*/
//go:nosplit
func (self class) RemoveTexture(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets font cache texture image.
*/
//go:nosplit
func (self class) SetTextureImage(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int, image gdclass.Image)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, discreet.Get(image[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_texture_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a copy of the font cache texture image.
*/
//go:nosplit
func (self class) GetTextureImage(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int) gdclass.Image {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Image{classdb.Image(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Sets array containing glyph packing data.
*/
//go:nosplit
func (self class) SetTextureOffsets(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int, offset gd.PackedInt32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, discreet.Get(offset))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a copy of the array containing glyph packing data.
*/
//go:nosplit
func (self class) GetTextureOffsets(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns list of rendered glyphs in the cache entry.
*/
//go:nosplit
func (self class) GetGlyphList(cache_index gd.Int, size gd.Vector2i) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) ClearGlyphs(cache_index gd.Int, size gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_glyphs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) RemoveGlyph(cache_index gd.Int, size gd.Vector2i, glyph gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_glyph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) SetGlyphAdvance(cache_index gd.Int, size gd.Int, glyph gd.Int, advance gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, advance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) GetGlyphAdvance(cache_index gd.Int, size gd.Int, glyph gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets glyph offset from the baseline.
*/
//go:nosplit
func (self class) SetGlyphOffset(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, offset gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns glyph offset from the baseline.
*/
//go:nosplit
func (self class) GetGlyphOffset(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets glyph size.
*/
//go:nosplit
func (self class) SetGlyphSize(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, gl_size gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, gl_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns glyph size.
*/
//go:nosplit
func (self class) GetGlyphSize(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) SetGlyphUvRect(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, uv_rect gd.Rect2)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, uv_rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) GetGlyphUvRect(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) SetGlyphTextureIdx(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, texture_idx gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, texture_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) GetGlyphTextureIdx(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns list of the kerning overrides.
*/
//go:nosplit
func (self class) GetKerningList(cache_index gd.Int, size gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_kerning_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes all kerning overrides.
*/
//go:nosplit
func (self class) ClearKerningMap(cache_index gd.Int, size gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_kerning_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes kerning override for the pair of glyphs.
*/
//go:nosplit
func (self class) RemoveKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) SetKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i, kerning gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	callframe.Arg(frame, kerning)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) GetKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renders the range of characters to the font cache texture.
*/
//go:nosplit
func (self class) RenderRange(cache_index gd.Int, size gd.Vector2i, start gd.Int, end gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_render_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renders specified glyph to the font cache texture.
*/
//go:nosplit
func (self class) RenderGlyph(cache_index gd.Int, size gd.Vector2i, index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_render_glyph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds override for [method Font.is_language_supported].
*/
//go:nosplit
func (self class) SetLanguageSupportOverride(language gd.String, supported bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(language))
	callframe.Arg(frame, supported)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
//go:nosplit
func (self class) GetLanguageSupportOverride(language gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Remove language support override.
*/
//go:nosplit
func (self class) RemoveLanguageSupportOverride(language gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns list of language support overrides.
*/
//go:nosplit
func (self class) GetLanguageSupportOverrides() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_language_support_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds override for [method Font.is_script_supported].
*/
//go:nosplit
func (self class) SetScriptSupportOverride(script gd.String, supported bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(script))
	callframe.Arg(frame, supported)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
//go:nosplit
func (self class) GetScriptSupportOverride(script gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(script))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes script support override.
*/
//go:nosplit
func (self class) RemoveScriptSupportOverride(script gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(script))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns list of script support overrides.
*/
//go:nosplit
func (self class) GetScriptSupportOverrides() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_script_support_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOpentypeFeatureOverrides(overrides gd.Dictionary)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(overrides))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOpentypeFeatureOverrides() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
//go:nosplit
func (self class) GetGlyphIndex(size gd.Int, char gd.Int, variation_selector gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, char)
	callframe.Arg(frame, variation_selector)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method get_glyph_index].
*/
//go:nosplit
func (self class) GetCharFromGlyphIndex(size gd.Int, glyph_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_char_from_glyph_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFontFile() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsFontFile() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsFont() Font.GD { return *((*Font.GD)(unsafe.Pointer(&self))) }
func (self Go) AsFont() Font.Go { return *((*Font.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsFont(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsFont(), name)
	}
}
func init() {classdb.Register("FontFile", func(ptr gd.Object) any { return classdb.FontFile(ptr) })}
