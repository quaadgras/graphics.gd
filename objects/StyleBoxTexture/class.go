package StyleBoxTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/StyleBox"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A texture-based nine-patch [StyleBox], in a way similar to [NinePatchRect]. This stylebox performs a 3×3 scaling of a texture, where only the center cell is fully stretched. This makes it possible to design bordered styles regardless of the stylebox's size.
*/
type Instance [1]classdb.StyleBoxTexture

/*
Sets the margin to [param size] pixels for all sides.
*/
func (self Instance) SetTextureMarginAll(size float64) {
	class(self).SetTextureMarginAll(gd.Float(size))
}

/*
Sets the expand margin to [param size] pixels for all sides.
*/
func (self Instance) SetExpandMarginAll(size float64) {
	class(self).SetExpandMarginAll(gd.Float(size))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StyleBoxTexture

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBoxTexture"))
	return Instance{classdb.StyleBoxTexture(object)}
}

func (self Instance) Texture() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) TextureMarginLeft() float64 {
	return float64(float64(class(self).GetTextureMargin(0)))
}

func (self Instance) SetTextureMarginLeft(value float64) {
	class(self).SetTextureMargin(0, gd.Float(value))
}

func (self Instance) TextureMarginTop() float64 {
	return float64(float64(class(self).GetTextureMargin(1)))
}

func (self Instance) SetTextureMarginTop(value float64) {
	class(self).SetTextureMargin(1, gd.Float(value))
}

func (self Instance) TextureMarginRight() float64 {
	return float64(float64(class(self).GetTextureMargin(2)))
}

func (self Instance) SetTextureMarginRight(value float64) {
	class(self).SetTextureMargin(2, gd.Float(value))
}

func (self Instance) TextureMarginBottom() float64 {
	return float64(float64(class(self).GetTextureMargin(3)))
}

func (self Instance) SetTextureMarginBottom(value float64) {
	class(self).SetTextureMargin(3, gd.Float(value))
}

func (self Instance) ExpandMarginLeft() float64 {
	return float64(float64(class(self).GetExpandMargin(0)))
}

func (self Instance) SetExpandMarginLeft(value float64) {
	class(self).SetExpandMargin(0, gd.Float(value))
}

func (self Instance) ExpandMarginTop() float64 {
	return float64(float64(class(self).GetExpandMargin(1)))
}

func (self Instance) SetExpandMarginTop(value float64) {
	class(self).SetExpandMargin(1, gd.Float(value))
}

func (self Instance) ExpandMarginRight() float64 {
	return float64(float64(class(self).GetExpandMargin(2)))
}

func (self Instance) SetExpandMarginRight(value float64) {
	class(self).SetExpandMargin(2, gd.Float(value))
}

func (self Instance) ExpandMarginBottom() float64 {
	return float64(float64(class(self).GetExpandMargin(3)))
}

func (self Instance) SetExpandMarginBottom(value float64) {
	class(self).SetExpandMargin(3, gd.Float(value))
}

func (self Instance) AxisStretchHorizontal() classdb.StyleBoxTextureAxisStretchMode {
	return classdb.StyleBoxTextureAxisStretchMode(class(self).GetHAxisStretchMode())
}

func (self Instance) SetAxisStretchHorizontal(value classdb.StyleBoxTextureAxisStretchMode) {
	class(self).SetHAxisStretchMode(value)
}

func (self Instance) AxisStretchVertical() classdb.StyleBoxTextureAxisStretchMode {
	return classdb.StyleBoxTextureAxisStretchMode(class(self).GetVAxisStretchMode())
}

func (self Instance) SetAxisStretchVertical(value classdb.StyleBoxTextureAxisStretchMode) {
	class(self).SetVAxisStretchMode(value)
}

func (self Instance) RegionRect() gd.Rect2 {
	return gd.Rect2(class(self).GetRegionRect())
}

func (self Instance) SetRegionRect(value gd.Rect2) {
	class(self).SetRegionRect(value)
}

func (self Instance) ModulateColor() gd.Color {
	return gd.Color(class(self).GetModulate())
}

func (self Instance) SetModulateColor(value gd.Color) {
	class(self).SetModulate(value)
}

func (self Instance) DrawCenter() bool {
	return bool(class(self).IsDrawCenterEnabled())
}

func (self Instance) SetDrawCenter(value bool) {
	class(self).SetDrawCenter(value)
}

//go:nosplit
func (self class) SetTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets the margin to [param size] pixels for the specified [enum Side].
*/
//go:nosplit
func (self class) SetTextureMargin(margin gdconst.Side, size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_texture_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the margin to [param size] pixels for all sides.
*/
//go:nosplit
func (self class) SetTextureMarginAll(size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_texture_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the margin size of the specified [enum Side].
*/
//go:nosplit
func (self class) GetTextureMargin(margin gdconst.Side) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_texture_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the expand margin to [param size] pixels for the specified [enum Side].
*/
//go:nosplit
func (self class) SetExpandMargin(margin gdconst.Side, size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_expand_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the expand margin to [param size] pixels for all sides.
*/
//go:nosplit
func (self class) SetExpandMarginAll(size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_expand_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the expand margin size of the specified [enum Side].
*/
//go:nosplit
func (self class) GetExpandMargin(margin gdconst.Side) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_expand_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegionRect(region gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRegionRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawCenter(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_draw_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawCenterEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_is_draw_center_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModulate(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetModulate() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHAxisStretchMode(mode classdb.StyleBoxTextureAxisStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHAxisStretchMode() classdb.StyleBoxTextureAxisStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.StyleBoxTextureAxisStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVAxisStretchMode(mode classdb.StyleBoxTextureAxisStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVAxisStretchMode() classdb.StyleBoxTextureAxisStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.StyleBoxTextureAxisStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsStyleBoxTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStyleBoxTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStyleBox() StyleBox.Advanced {
	return *((*StyleBox.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStyleBox() StyleBox.Instance {
	return *((*StyleBox.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStyleBox(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStyleBox(), name)
	}
}
func init() {
	classdb.Register("StyleBoxTexture", func(ptr gd.Object) any { return classdb.StyleBoxTexture(ptr) })
}

type AxisStretchMode = classdb.StyleBoxTextureAxisStretchMode

const (
	/*Stretch the stylebox's texture. This results in visible distortion unless the texture size matches the stylebox's size perfectly.*/
	AxisStretchModeStretch AxisStretchMode = 0
	/*Repeats the stylebox's texture to match the stylebox's size according to the nine-patch system.*/
	AxisStretchModeTile AxisStretchMode = 1
	/*Repeats the stylebox's texture to match the stylebox's size according to the nine-patch system. Unlike [constant AXIS_STRETCH_MODE_TILE], the texture may be slightly stretched to make the nine-patch texture tile seamlessly.*/
	AxisStretchModeTileFit AxisStretchMode = 2
)
