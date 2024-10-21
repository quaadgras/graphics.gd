package SystemFont

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
[SystemFont] loads a font from a system font with the first matching name from [member font_names].
It will attempt to match font style, but it's not guaranteed.
The returned font might be part of a font collection or be a variable font with OpenType "weight", "width" and/or "italic" features set.
You can create [FontVariation] of the system font for precise control over its features.
[b]Note:[/b] This class is implemented on iOS, Linux, macOS and Windows, on other platforms it will fallback to default theme font.

*/
type Simple [1]classdb.SystemFont
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
func (self Simple) SetOversampling(oversampling float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOversampling(gd.Float(oversampling))
}
func (self Simple) GetOversampling() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetOversampling()))
}
func (self Simple) GetFontNames() gd.PackedStringArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedStringArray(Expert(self).GetFontNames(gc))
}
func (self Simple) SetFontNames(names gd.PackedStringArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontNames(names)
}
func (self Simple) GetFontItalic() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetFontItalic())
}
func (self Simple) SetFontItalic(italic bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontItalic(italic)
}
func (self Simple) SetFontWeight(weight int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontWeight(gd.Int(weight))
}
func (self Simple) SetFontStretch(stretch int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFontStretch(gd.Int(stretch))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SystemFont
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetAntialiasing(antialiasing classdb.TextServerFontAntialiasing)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, antialiasing)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAntialiasing() classdb.TextServerFontAntialiasing {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerFontAntialiasing](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_antialiasing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDisableEmbeddedBitmaps() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGenerateMipmaps() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAllowSystemFallback() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_is_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsForceAutohinter() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_is_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_hinting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHinting() classdb.TextServerHinting {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerHinting](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_hinting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSubpixelPositioning() classdb.TextServerSubpixelPositioning {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerSubpixelPositioning](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsMultichannelSignedDistanceField() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_is_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsdfPixelRange() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_msdf_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMsdfSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_msdf_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOversampling() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetFontNames(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_font_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFontNames(names gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(names))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_font_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFontItalic() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_get_font_italic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFontItalic(italic bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, italic)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_font_italic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontWeight(weight gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, weight)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_font_weight, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFontStretch(stretch gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SystemFont.Bind_set_font_stretch, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsSystemFont() Expert { return self[0].AsSystemFont() }


//go:nosplit
func (self Simple) AsSystemFont() Simple { return self[0].AsSystemFont() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SystemFont", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
