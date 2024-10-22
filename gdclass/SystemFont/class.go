package SystemFont

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
[SystemFont] loads a font from a system font with the first matching name from [member font_names].
It will attempt to match font style, but it's not guaranteed.
The returned font might be part of a font collection or be a variable font with OpenType "weight", "width" and/or "italic" features set.
You can create [FontVariation] of the system font for precise control over its features.
[b]Note:[/b] This class is implemented on iOS, Linux, macOS and Windows, on other platforms it will fallback to default theme font.

*/
type Go [1]classdb.SystemFont
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SystemFont
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("SystemFont"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) FontNames() []string {
	gc := gd.GarbageCollector(); _ = gc
		return []string(class(self).GetFontNames(gc).Strings(gc))
}

func (self Go) SetFontNames(value []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFontNames(gc.PackedStringSlice(value))
}

func (self Go) FontItalic() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetFontItalic())
}

func (self Go) SetFontItalic(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFontItalic(value)
}

func (self Go) Antialiasing() classdb.TextServerFontAntialiasing {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerFontAntialiasing(class(self).GetAntialiasing())
}

func (self Go) SetAntialiasing(value classdb.TextServerFontAntialiasing) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAntialiasing(value)
}

func (self Go) GenerateMipmaps() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetGenerateMipmaps())
}

func (self Go) SetGenerateMipmaps(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetGenerateMipmaps(value)
}

func (self Go) DisableEmbeddedBitmaps() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetDisableEmbeddedBitmaps())
}

func (self Go) SetDisableEmbeddedBitmaps(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDisableEmbeddedBitmaps(value)
}

func (self Go) AllowSystemFallback() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAllowSystemFallback())
}

func (self Go) SetAllowSystemFallback(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAllowSystemFallback(value)
}

func (self Go) ForceAutohinter() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsForceAutohinter())
}

func (self Go) SetForceAutohinter(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetForceAutohinter(value)
}

func (self Go) Hinting() classdb.TextServerHinting {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerHinting(class(self).GetHinting())
}

func (self Go) SetHinting(value classdb.TextServerHinting) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHinting(value)
}

func (self Go) SubpixelPositioning() classdb.TextServerSubpixelPositioning {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.TextServerSubpixelPositioning(class(self).GetSubpixelPositioning())
}

func (self Go) SetSubpixelPositioning(value classdb.TextServerSubpixelPositioning) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSubpixelPositioning(value)
}

func (self Go) MultichannelSignedDistanceField() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsMultichannelSignedDistanceField())
}

func (self Go) SetMultichannelSignedDistanceField(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMultichannelSignedDistanceField(value)
}

func (self Go) MsdfPixelRange() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMsdfPixelRange()))
}

func (self Go) SetMsdfPixelRange(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMsdfPixelRange(gd.Int(value))
}

func (self Go) MsdfSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMsdfSize()))
}

func (self Go) SetMsdfSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMsdfSize(gd.Int(value))
}

func (self Go) Oversampling() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetOversampling()))
}

func (self Go) SetOversampling(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOversampling(gd.Float(value))
}

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
func (self class) AsSystemFont() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSystemFont() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("SystemFont", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
