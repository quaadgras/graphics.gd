package FontFile

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Font"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Simple [1]classdb.FontFile
func (self Simple) LoadBitmapFont(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadBitmapFont(gc.String(path)))
}
func (self Simple) LoadDynamicFont(path string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).LoadDynamicFont(gc.String(path)))
}
func (self Simple) SetData(data []byte) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetData(gc.PackedByteSlice(data))
}
func (self Simple) GetData() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(Expert(self).GetData(gc).Bytes())
}
func (self Simple) SetFontName(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontName(gc.String(name))
}
func (self Simple) SetFontStyleName(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontStyleName(gc.String(name))
}
func (self Simple) SetFontStyle(style classdb.TextServerFontStyle) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontStyle(style)
}
func (self Simple) SetFontWeight(weight int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontWeight(gd.Int(weight))
}
func (self Simple) SetFontStretch(stretch int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontStretch(gd.Int(stretch))
}
func (self Simple) SetAntialiasing(antialiasing classdb.TextServerFontAntialiasing) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAntialiasing(antialiasing)
}
func (self Simple) GetAntialiasing() classdb.TextServerFontAntialiasing {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerFontAntialiasing(Expert(self).GetAntialiasing())
}
func (self Simple) SetDisableEmbeddedBitmaps(disable_embedded_bitmaps bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDisableEmbeddedBitmaps(disable_embedded_bitmaps)
}
func (self Simple) GetDisableEmbeddedBitmaps() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetDisableEmbeddedBitmaps())
}
func (self Simple) SetGenerateMipmaps(generate_mipmaps bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGenerateMipmaps(generate_mipmaps)
}
func (self Simple) GetGenerateMipmaps() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetGenerateMipmaps())
}
func (self Simple) SetMultichannelSignedDistanceField(msdf bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMultichannelSignedDistanceField(msdf)
}
func (self Simple) IsMultichannelSignedDistanceField() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsMultichannelSignedDistanceField())
}
func (self Simple) SetMsdfPixelRange(msdf_pixel_range int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMsdfPixelRange(gd.Int(msdf_pixel_range))
}
func (self Simple) GetMsdfPixelRange() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMsdfPixelRange()))
}
func (self Simple) SetMsdfSize(msdf_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMsdfSize(gd.Int(msdf_size))
}
func (self Simple) GetMsdfSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetMsdfSize()))
}
func (self Simple) SetFixedSize(fixed_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFixedSize(gd.Int(fixed_size))
}
func (self Simple) GetFixedSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFixedSize()))
}
func (self Simple) SetFixedSizeScaleMode(fixed_size_scale_mode classdb.TextServerFixedSizeScaleMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFixedSizeScaleMode(fixed_size_scale_mode)
}
func (self Simple) GetFixedSizeScaleMode() classdb.TextServerFixedSizeScaleMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerFixedSizeScaleMode(Expert(self).GetFixedSizeScaleMode())
}
func (self Simple) SetAllowSystemFallback(allow_system_fallback bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAllowSystemFallback(allow_system_fallback)
}
func (self Simple) IsAllowSystemFallback() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAllowSystemFallback())
}
func (self Simple) SetForceAutohinter(force_autohinter bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetForceAutohinter(force_autohinter)
}
func (self Simple) IsForceAutohinter() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsForceAutohinter())
}
func (self Simple) SetHinting(hinting classdb.TextServerHinting) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHinting(hinting)
}
func (self Simple) GetHinting() classdb.TextServerHinting {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerHinting(Expert(self).GetHinting())
}
func (self Simple) SetSubpixelPositioning(subpixel_positioning classdb.TextServerSubpixelPositioning) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSubpixelPositioning(subpixel_positioning)
}
func (self Simple) GetSubpixelPositioning() classdb.TextServerSubpixelPositioning {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TextServerSubpixelPositioning(Expert(self).GetSubpixelPositioning())
}
func (self Simple) SetOversampling(oversampling float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOversampling(gd.Float(oversampling))
}
func (self Simple) GetOversampling() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetOversampling()))
}
func (self Simple) GetCacheCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCacheCount()))
}
func (self Simple) ClearCache() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearCache()
}
func (self Simple) RemoveCache(cache_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveCache(gd.Int(cache_index))
}
func (self Simple) GetSizeCacheList(cache_index int) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](Expert(self).GetSizeCacheList(gc, gd.Int(cache_index)))
}
func (self Simple) ClearSizeCache(cache_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearSizeCache(gd.Int(cache_index))
}
func (self Simple) RemoveSizeCache(cache_index int, size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveSizeCache(gd.Int(cache_index), size)
}
func (self Simple) SetVariationCoordinates(cache_index int, variation_coordinates gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVariationCoordinates(gd.Int(cache_index), variation_coordinates)
}
func (self Simple) GetVariationCoordinates(cache_index int) gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetVariationCoordinates(gc, gd.Int(cache_index)))
}
func (self Simple) SetEmbolden(cache_index int, strength float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmbolden(gd.Int(cache_index), gd.Float(strength))
}
func (self Simple) GetEmbolden(cache_index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEmbolden(gd.Int(cache_index))))
}
func (self Simple) SetTransform(cache_index int, transform gd.Transform2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTransform(gd.Int(cache_index), transform)
}
func (self Simple) GetTransform(cache_index int) gd.Transform2D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform2D(Expert(self).GetTransform(gd.Int(cache_index)))
}
func (self Simple) SetExtraSpacing(cache_index int, spacing classdb.TextServerSpacingType, value int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExtraSpacing(gd.Int(cache_index), spacing, gd.Int(value))
}
func (self Simple) GetExtraSpacing(cache_index int, spacing classdb.TextServerSpacingType) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetExtraSpacing(gd.Int(cache_index), spacing)))
}
func (self Simple) SetExtraBaselineOffset(cache_index int, baseline_offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetExtraBaselineOffset(gd.Int(cache_index), gd.Float(baseline_offset))
}
func (self Simple) GetExtraBaselineOffset(cache_index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetExtraBaselineOffset(gd.Int(cache_index))))
}
func (self Simple) SetFaceIndex(cache_index int, face_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFaceIndex(gd.Int(cache_index), gd.Int(face_index))
}
func (self Simple) GetFaceIndex(cache_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetFaceIndex(gd.Int(cache_index))))
}
func (self Simple) SetCacheAscent(cache_index int, size int, ascent float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCacheAscent(gd.Int(cache_index), gd.Int(size), gd.Float(ascent))
}
func (self Simple) GetCacheAscent(cache_index int, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCacheAscent(gd.Int(cache_index), gd.Int(size))))
}
func (self Simple) SetCacheDescent(cache_index int, size int, descent float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCacheDescent(gd.Int(cache_index), gd.Int(size), gd.Float(descent))
}
func (self Simple) GetCacheDescent(cache_index int, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCacheDescent(gd.Int(cache_index), gd.Int(size))))
}
func (self Simple) SetCacheUnderlinePosition(cache_index int, size int, underline_position float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCacheUnderlinePosition(gd.Int(cache_index), gd.Int(size), gd.Float(underline_position))
}
func (self Simple) GetCacheUnderlinePosition(cache_index int, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCacheUnderlinePosition(gd.Int(cache_index), gd.Int(size))))
}
func (self Simple) SetCacheUnderlineThickness(cache_index int, size int, underline_thickness float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCacheUnderlineThickness(gd.Int(cache_index), gd.Int(size), gd.Float(underline_thickness))
}
func (self Simple) GetCacheUnderlineThickness(cache_index int, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCacheUnderlineThickness(gd.Int(cache_index), gd.Int(size))))
}
func (self Simple) SetCacheScale(cache_index int, size int, scale float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCacheScale(gd.Int(cache_index), gd.Int(size), gd.Float(scale))
}
func (self Simple) GetCacheScale(cache_index int, size int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetCacheScale(gd.Int(cache_index), gd.Int(size))))
}
func (self Simple) GetTextureCount(cache_index int, size gd.Vector2i) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetTextureCount(gd.Int(cache_index), size)))
}
func (self Simple) ClearTextures(cache_index int, size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearTextures(gd.Int(cache_index), size)
}
func (self Simple) RemoveTexture(cache_index int, size gd.Vector2i, texture_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveTexture(gd.Int(cache_index), size, gd.Int(texture_index))
}
func (self Simple) SetTextureImage(cache_index int, size gd.Vector2i, texture_index int, image [1]classdb.Image) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureImage(gd.Int(cache_index), size, gd.Int(texture_index), image)
}
func (self Simple) GetTextureImage(cache_index int, size gd.Vector2i, texture_index int) [1]classdb.Image {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Image(Expert(self).GetTextureImage(gc, gd.Int(cache_index), size, gd.Int(texture_index)))
}
func (self Simple) SetTextureOffsets(cache_index int, size gd.Vector2i, texture_index int, offset gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureOffsets(gd.Int(cache_index), size, gd.Int(texture_index), offset)
}
func (self Simple) GetTextureOffsets(cache_index int, size gd.Vector2i, texture_index int) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetTextureOffsets(gc, gd.Int(cache_index), size, gd.Int(texture_index)))
}
func (self Simple) GetGlyphList(cache_index int, size gd.Vector2i) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetGlyphList(gc, gd.Int(cache_index), size))
}
func (self Simple) ClearGlyphs(cache_index int, size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearGlyphs(gd.Int(cache_index), size)
}
func (self Simple) RemoveGlyph(cache_index int, size gd.Vector2i, glyph int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveGlyph(gd.Int(cache_index), size, gd.Int(glyph))
}
func (self Simple) SetGlyphAdvance(cache_index int, size int, glyph int, advance gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlyphAdvance(gd.Int(cache_index), gd.Int(size), gd.Int(glyph), advance)
}
func (self Simple) GetGlyphAdvance(cache_index int, size int, glyph int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGlyphAdvance(gd.Int(cache_index), gd.Int(size), gd.Int(glyph)))
}
func (self Simple) SetGlyphOffset(cache_index int, size gd.Vector2i, glyph int, offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlyphOffset(gd.Int(cache_index), size, gd.Int(glyph), offset)
}
func (self Simple) GetGlyphOffset(cache_index int, size gd.Vector2i, glyph int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGlyphOffset(gd.Int(cache_index), size, gd.Int(glyph)))
}
func (self Simple) SetGlyphSize(cache_index int, size gd.Vector2i, glyph int, gl_size gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlyphSize(gd.Int(cache_index), size, gd.Int(glyph), gl_size)
}
func (self Simple) GetGlyphSize(cache_index int, size gd.Vector2i, glyph int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGlyphSize(gd.Int(cache_index), size, gd.Int(glyph)))
}
func (self Simple) SetGlyphUvRect(cache_index int, size gd.Vector2i, glyph int, uv_rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlyphUvRect(gd.Int(cache_index), size, gd.Int(glyph), uv_rect)
}
func (self Simple) GetGlyphUvRect(cache_index int, size gd.Vector2i, glyph int) gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetGlyphUvRect(gd.Int(cache_index), size, gd.Int(glyph)))
}
func (self Simple) SetGlyphTextureIdx(cache_index int, size gd.Vector2i, glyph int, texture_idx int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlyphTextureIdx(gd.Int(cache_index), size, gd.Int(glyph), gd.Int(texture_idx))
}
func (self Simple) GetGlyphTextureIdx(cache_index int, size gd.Vector2i, glyph int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetGlyphTextureIdx(gd.Int(cache_index), size, gd.Int(glyph))))
}
func (self Simple) GetKerningList(cache_index int, size int) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](Expert(self).GetKerningList(gc, gd.Int(cache_index), gd.Int(size)))
}
func (self Simple) ClearKerningMap(cache_index int, size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearKerningMap(gd.Int(cache_index), gd.Int(size))
}
func (self Simple) RemoveKerning(cache_index int, size int, glyph_pair gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveKerning(gd.Int(cache_index), gd.Int(size), glyph_pair)
}
func (self Simple) SetKerning(cache_index int, size int, glyph_pair gd.Vector2i, kerning gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKerning(gd.Int(cache_index), gd.Int(size), glyph_pair, kerning)
}
func (self Simple) GetKerning(cache_index int, size int, glyph_pair gd.Vector2i) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetKerning(gd.Int(cache_index), gd.Int(size), glyph_pair))
}
func (self Simple) RenderRange(cache_index int, size gd.Vector2i, start int, end int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RenderRange(gd.Int(cache_index), size, gd.Int(start), gd.Int(end))
}
func (self Simple) RenderGlyph(cache_index int, size gd.Vector2i, index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RenderGlyph(gd.Int(cache_index), size, gd.Int(index))
}
func (self Simple) SetLanguageSupportOverride(language string, supported bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLanguageSupportOverride(gc.String(language), supported)
}
func (self Simple) GetLanguageSupportOverride(language string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetLanguageSupportOverride(gc.String(language)))
}
func (self Simple) RemoveLanguageSupportOverride(language string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveLanguageSupportOverride(gc.String(language))
}
func (self Simple) GetLanguageSupportOverrides() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetLanguageSupportOverrides(gc))
}
func (self Simple) SetScriptSupportOverride(script string, supported bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScriptSupportOverride(gc.String(script), supported)
}
func (self Simple) GetScriptSupportOverride(script string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetScriptSupportOverride(gc.String(script)))
}
func (self Simple) RemoveScriptSupportOverride(script string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveScriptSupportOverride(gc.String(script))
}
func (self Simple) GetScriptSupportOverrides() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetScriptSupportOverrides(gc))
}
func (self Simple) SetOpentypeFeatureOverrides(overrides gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOpentypeFeatureOverrides(overrides)
}
func (self Simple) GetOpentypeFeatureOverrides() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(Expert(self).GetOpentypeFeatureOverrides(gc))
}
func (self Simple) GetGlyphIndex(size int, char int, variation_selector int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetGlyphIndex(gd.Int(size), gd.Int(char), gd.Int(variation_selector))))
}
func (self Simple) GetCharFromGlyphIndex(size int, glyph_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCharFromGlyphIndex(gd.Int(size), gd.Int(glyph_index))))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.FontFile
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Loads an AngelCode BMFont (.fnt, .font) bitmap font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
//go:nosplit
func (self class) LoadBitmapFont(path gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_load_bitmap_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_load_dynamic_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetData(data gd.PackedByteArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetData(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFontName(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_font_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontStyleName(name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_font_style_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontStyle(style classdb.TextServerFontStyle)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, style)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_font_style, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontWeight(weight gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_font_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontStretch(stretch gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_font_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetAntialiasing(antialiasing classdb.TextServerFontAntialiasing)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, antialiasing)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAntialiasing() classdb.TextServerFontAntialiasing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerFontAntialiasing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisableEmbeddedBitmaps(disable_embedded_bitmaps bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, disable_embedded_bitmaps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDisableEmbeddedBitmaps() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGenerateMipmaps(generate_mipmaps bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, generate_mipmaps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGenerateMipmaps() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMultichannelSignedDistanceField(msdf bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msdf)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMultichannelSignedDistanceField() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_is_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsdfPixelRange(msdf_pixel_range gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msdf_pixel_range)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsdfPixelRange() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMsdfSize(msdf_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, msdf_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_msdf_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsdfSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_msdf_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFixedSize(fixed_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fixed_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_fixed_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFixedSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_fixed_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFixedSizeScaleMode(fixed_size_scale_mode classdb.TextServerFixedSizeScaleMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, fixed_size_scale_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFixedSizeScaleMode() classdb.TextServerFixedSizeScaleMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerFixedSizeScaleMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAllowSystemFallback(allow_system_fallback bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, allow_system_fallback)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAllowSystemFallback() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_is_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetForceAutohinter(force_autohinter bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force_autohinter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsForceAutohinter() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_is_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHinting(hinting classdb.TextServerHinting)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hinting)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_hinting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHinting() classdb.TextServerHinting {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerHinting](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_hinting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSubpixelPositioning(subpixel_positioning classdb.TextServerSubpixelPositioning)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, subpixel_positioning)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubpixelPositioning() classdb.TextServerSubpixelPositioning {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerSubpixelPositioning](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOversampling(oversampling gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, oversampling)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOversampling() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of the font cache entries.
*/
//go:nosplit
func (self class) GetCacheCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_cache_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all font cache entries.
*/
//go:nosplit
func (self class) ClearCache()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_clear_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified font cache entry.
*/
//go:nosplit
func (self class) RemoveCache(cache_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_remove_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
//go:nosplit
func (self class) GetSizeCacheList(ctx gd.Lifetime, cache_index gd.Int) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_size_cache_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Removes all font sizes from the cache entry
*/
//go:nosplit
func (self class) ClearSizeCache(cache_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_clear_size_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified font size from the cache entry.
*/
//go:nosplit
func (self class) RemoveSizeCache(cache_index gd.Int, size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_remove_size_cache, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) SetVariationCoordinates(cache_index gd.Int, variation_coordinates gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, mmm.Get(variation_coordinates))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) GetVariationCoordinates(ctx gd.Lifetime, cache_index gd.Int) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) SetEmbolden(cache_index gd.Int, strength gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) GetEmbolden(cache_index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_embolden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
//go:nosplit
func (self class) SetTransform(cache_index gd.Int, transform gd.Transform2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns 2D transform, applied to the font outlines, can be used for slanting, flipping and rotating glyphs.
*/
//go:nosplit
func (self class) GetTransform(cache_index gd.Int) gd.Transform2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) SetExtraSpacing(cache_index gd.Int, spacing classdb.TextServerSpacingType, value gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_extra_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) GetExtraSpacing(cache_index gd.Int, spacing classdb.TextServerSpacingType) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_extra_spacing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) SetExtraBaselineOffset(cache_index gd.Int, baseline_offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, baseline_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_extra_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) GetExtraBaselineOffset(cache_index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_extra_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) SetFaceIndex(cache_index gd.Int, face_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, face_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) GetFaceIndex(cache_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) SetCacheAscent(cache_index gd.Int, size gd.Int, ascent gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, ascent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_cache_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) GetCacheAscent(cache_index gd.Int, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_cache_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) SetCacheDescent(cache_index gd.Int, size gd.Int, descent gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, descent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_cache_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) GetCacheDescent(cache_index gd.Int, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_cache_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) SetCacheUnderlinePosition(cache_index gd.Int, size gd.Int, underline_position gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_cache_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetCacheUnderlinePosition(cache_index gd.Int, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_cache_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets thickness of the underline in pixels.
*/
//go:nosplit
func (self class) SetCacheUnderlineThickness(cache_index gd.Int, size gd.Int, underline_thickness gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_thickness)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_cache_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns thickness of the underline in pixels.
*/
//go:nosplit
func (self class) GetCacheUnderlineThickness(cache_index gd.Int, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_cache_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) SetCacheScale(cache_index gd.Int, size gd.Int, scale gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_cache_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) GetCacheScale(cache_index gd.Int, size gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_cache_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of textures used by font cache entry.
*/
//go:nosplit
func (self class) GetTextureCount(cache_index gd.Int, size gd.Vector2i) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_texture_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_clear_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture. Remove them manually using [method remove_glyph].
*/
//go:nosplit
func (self class) RemoveTexture(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_remove_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets font cache texture image.
*/
//go:nosplit
func (self class) SetTextureImage(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int, image object.Image)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_texture_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a copy of the font cache texture image.
*/
//go:nosplit
func (self class) GetTextureImage(ctx gd.Lifetime, cache_index gd.Int, size gd.Vector2i, texture_index gd.Int) object.Image {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_texture_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Image
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets array containing glyph packing data.
*/
//go:nosplit
func (self class) SetTextureOffsets(cache_index gd.Int, size gd.Vector2i, texture_index gd.Int, offset gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, mmm.Get(offset))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a copy of the array containing glyph packing data.
*/
//go:nosplit
func (self class) GetTextureOffsets(ctx gd.Lifetime, cache_index gd.Int, size gd.Vector2i, texture_index gd.Int) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns list of rendered glyphs in the cache entry.
*/
//go:nosplit
func (self class) GetGlyphList(ctx gd.Lifetime, cache_index gd.Int, size gd.Vector2i) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_glyph_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) ClearGlyphs(cache_index gd.Int, size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_clear_glyphs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) RemoveGlyph(cache_index gd.Int, size gd.Vector2i, glyph gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_remove_glyph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) SetGlyphAdvance(cache_index gd.Int, size gd.Int, glyph gd.Int, advance gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, advance)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) GetGlyphAdvance(cache_index gd.Int, size gd.Int, glyph gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets glyph offset from the baseline.
*/
//go:nosplit
func (self class) SetGlyphOffset(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns glyph offset from the baseline.
*/
//go:nosplit
func (self class) GetGlyphOffset(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets glyph size.
*/
//go:nosplit
func (self class) SetGlyphSize(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, gl_size gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, gl_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_glyph_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns glyph size.
*/
//go:nosplit
func (self class) GetGlyphSize(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_glyph_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) SetGlyphUvRect(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, uv_rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, uv_rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) GetGlyphUvRect(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) SetGlyphTextureIdx(cache_index gd.Int, size gd.Vector2i, glyph gd.Int, texture_idx gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, texture_idx)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) GetGlyphTextureIdx(cache_index gd.Int, size gd.Vector2i, glyph gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns list of the kerning overrides.
*/
//go:nosplit
func (self class) GetKerningList(ctx gd.Lifetime, cache_index gd.Int, size gd.Int) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_kerning_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Removes all kerning overrides.
*/
//go:nosplit
func (self class) ClearKerningMap(cache_index gd.Int, size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_clear_kerning_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes kerning override for the pair of glyphs.
*/
//go:nosplit
func (self class) RemoveKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_remove_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) SetKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i, kerning gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	callframe.Arg(frame, kerning)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) GetKerning(cache_index gd.Int, size gd.Int, glyph_pair gd.Vector2i) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_kerning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Renders the range of characters to the font cache texture.
*/
//go:nosplit
func (self class) RenderRange(cache_index gd.Int, size gd.Vector2i, start gd.Int, end gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_render_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renders specified glyph to the font cache texture.
*/
//go:nosplit
func (self class) RenderGlyph(cache_index gd.Int, size gd.Vector2i, index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_render_glyph, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds override for [method Font.is_language_supported].
*/
//go:nosplit
func (self class) SetLanguageSupportOverride(language gd.String, supported bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	callframe.Arg(frame, supported)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
//go:nosplit
func (self class) GetLanguageSupportOverride(language gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Remove language support override.
*/
//go:nosplit
func (self class) RemoveLanguageSupportOverride(language gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_remove_language_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns list of language support overrides.
*/
//go:nosplit
func (self class) GetLanguageSupportOverrides(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_language_support_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Adds override for [method Font.is_script_supported].
*/
//go:nosplit
func (self class) SetScriptSupportOverride(script gd.String, supported bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(script))
	callframe.Arg(frame, supported)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
//go:nosplit
func (self class) GetScriptSupportOverride(script gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(script))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes script support override.
*/
//go:nosplit
func (self class) RemoveScriptSupportOverride(script gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(script))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_remove_script_support_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns list of script support overrides.
*/
//go:nosplit
func (self class) GetScriptSupportOverrides(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_script_support_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOpentypeFeatureOverrides(overrides gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(overrides))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_set_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOpentypeFeatureOverrides(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
//go:nosplit
func (self class) GetGlyphIndex(size gd.Int, char gd.Int, variation_selector gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, char)
	callframe.Arg(frame, variation_selector)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method get_glyph_index].
*/
//go:nosplit
func (self class) GetCharFromGlyphIndex(size gd.Int, glyph_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.FontFile.Bind_get_char_from_glyph_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsFontFile() Expert { return self[0].AsFontFile() }


//go:nosplit
func (self Simple) AsFontFile() Simple { return self[0].AsFontFile() }


//go:nosplit
func (self class) AsFont() Font.Expert { return self[0].AsFont() }


//go:nosplit
func (self Simple) AsFont() Font.Simple { return self[0].AsFont() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("FontFile", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
