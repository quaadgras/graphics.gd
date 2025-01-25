// Package SystemFont provides methods for working with SystemFont object instances.
package SystemFont

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
import "graphics.gd/classdb/Font"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
[SystemFont] loads a font from a system font with the first matching name from [member font_names].
It will attempt to match font style, but it's not guaranteed.
The returned font might be part of a font collection or be a variable font with OpenType "weight", "width" and/or "italic" features set.
You can create [FontVariation] of the system font for precise control over its features.
[b]Note:[/b] This class is implemented on iOS, Linux, macOS and Windows, on other platforms it will fallback to default theme font.
*/
type Instance [1]gdclass.SystemFont

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSystemFont() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SystemFont

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SystemFont"))
	casted := Instance{*(*gdclass.SystemFont)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) FontNames() []string {
	return []string(class(self).GetFontNames().Strings())
}

func (self Instance) SetFontNames(value []string) {
	class(self).SetFontNames(gd.NewPackedStringSlice(value))
}

func (self Instance) FontItalic() bool {
	return bool(class(self).GetFontItalic())
}

func (self Instance) SetFontItalic(value bool) {
	class(self).SetFontItalic(value)
}

func (self Instance) SetFontWeight(value int) {
	class(self).SetFontWeight(gd.Int(value))
}

func (self Instance) SetFontStretch(value int) {
	class(self).SetFontStretch(gd.Int(value))
}

func (self Instance) Antialiasing() gdclass.TextServerFontAntialiasing {
	return gdclass.TextServerFontAntialiasing(class(self).GetAntialiasing())
}

func (self Instance) SetAntialiasing(value gdclass.TextServerFontAntialiasing) {
	class(self).SetAntialiasing(value)
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

func (self Instance) Oversampling() Float.X {
	return Float.X(Float.X(class(self).GetOversampling()))
}

func (self Instance) SetOversampling(value Float.X) {
	class(self).SetOversampling(gd.Float(value))
}

//go:nosplit
func (self class) SetAntialiasing(antialiasing gdclass.TextServerFontAntialiasing) {
	var frame = callframe.New()
	callframe.Arg(frame, antialiasing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAntialiasing() gdclass.TextServerFontAntialiasing {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerFontAntialiasing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableEmbeddedBitmaps(disable_embedded_bitmaps bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disable_embedded_bitmaps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisableEmbeddedBitmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGenerateMipmaps(generate_mipmaps bool) {
	var frame = callframe.New()
	callframe.Arg(frame, generate_mipmaps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGenerateMipmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowSystemFallback(allow_system_fallback bool) {
	var frame = callframe.New()
	callframe.Arg(frame, allow_system_fallback)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllowSystemFallback() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_is_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetForceAutohinter(force_autohinter bool) {
	var frame = callframe.New()
	callframe.Arg(frame, force_autohinter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsForceAutohinter() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_is_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHinting(hinting gdclass.TextServerHinting) {
	var frame = callframe.New()
	callframe.Arg(frame, hinting)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_hinting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHinting() gdclass.TextServerHinting {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerHinting](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_hinting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubpixelPositioning(subpixel_positioning gdclass.TextServerSubpixelPositioning) {
	var frame = callframe.New()
	callframe.Arg(frame, subpixel_positioning)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubpixelPositioning() gdclass.TextServerSubpixelPositioning {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerSubpixelPositioning](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMultichannelSignedDistanceField(msdf bool) {
	var frame = callframe.New()
	callframe.Arg(frame, msdf)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMultichannelSignedDistanceField() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_is_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsdfPixelRange(msdf_pixel_range gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, msdf_pixel_range)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsdfPixelRange() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsdfSize(msdf_size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, msdf_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_msdf_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsdfSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_msdf_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOversampling(oversampling gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, oversampling)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOversampling() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetFontNames() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_font_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFontNames(names gd.PackedStringArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(names))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_font_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFontItalic() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_get_font_italic, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFontItalic(italic bool) {
	var frame = callframe.New()
	callframe.Arg(frame, italic)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_font_italic, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontWeight(weight gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, weight)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_font_weight, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontStretch(stretch gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SystemFont.Bind_set_font_stretch, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsSystemFont() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSystemFont() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsFont() Font.Advanced     { return *((*Font.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFont() Font.Instance  { return *((*Font.Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("SystemFont", func(ptr gd.Object) any { return [1]gdclass.SystemFont{*(*gdclass.SystemFont)(unsafe.Pointer(&ptr))} })
}
