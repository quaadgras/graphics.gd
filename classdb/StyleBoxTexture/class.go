// Package StyleBoxTexture provides methods for working with StyleBoxTexture object instances.
package StyleBoxTexture

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/StyleBox"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A texture-based nine-patch [StyleBox], in a way similar to [NinePatchRect]. This stylebox performs a 3×3 scaling of a texture, where only the center cell is fully stretched. This makes it possible to design bordered styles regardless of the stylebox's size.
*/
type Instance [1]gdclass.StyleBoxTexture

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStyleBoxTexture() Instance
}

/*
Sets the margin to [param size] pixels for all sides.
*/
func (self Instance) SetTextureMarginAll(size Float.X) { //gd:StyleBoxTexture.set_texture_margin_all
	class(self).SetTextureMarginAll(float64(size))
}

/*
Sets the expand margin to [param size] pixels for all sides.
*/
func (self Instance) SetExpandMarginAll(size Float.X) { //gd:StyleBoxTexture.set_expand_margin_all
	class(self).SetExpandMarginAll(float64(size))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StyleBoxTexture

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBoxTexture"))
	casted := Instance{*(*gdclass.StyleBoxTexture)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Texture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value [1]gdclass.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) TextureMarginLeft() Float.X {
	return Float.X(Float.X(class(self).GetTextureMargin(0)))
}

func (self Instance) SetTextureMarginLeft(value Float.X) {
	class(self).SetTextureMargin(0, float64(value))
}

func (self Instance) TextureMarginTop() Float.X {
	return Float.X(Float.X(class(self).GetTextureMargin(1)))
}

func (self Instance) SetTextureMarginTop(value Float.X) {
	class(self).SetTextureMargin(1, float64(value))
}

func (self Instance) TextureMarginRight() Float.X {
	return Float.X(Float.X(class(self).GetTextureMargin(2)))
}

func (self Instance) SetTextureMarginRight(value Float.X) {
	class(self).SetTextureMargin(2, float64(value))
}

func (self Instance) TextureMarginBottom() Float.X {
	return Float.X(Float.X(class(self).GetTextureMargin(3)))
}

func (self Instance) SetTextureMarginBottom(value Float.X) {
	class(self).SetTextureMargin(3, float64(value))
}

func (self Instance) ExpandMarginLeft() Float.X {
	return Float.X(Float.X(class(self).GetExpandMargin(0)))
}

func (self Instance) SetExpandMarginLeft(value Float.X) {
	class(self).SetExpandMargin(0, float64(value))
}

func (self Instance) ExpandMarginTop() Float.X {
	return Float.X(Float.X(class(self).GetExpandMargin(1)))
}

func (self Instance) SetExpandMarginTop(value Float.X) {
	class(self).SetExpandMargin(1, float64(value))
}

func (self Instance) ExpandMarginRight() Float.X {
	return Float.X(Float.X(class(self).GetExpandMargin(2)))
}

func (self Instance) SetExpandMarginRight(value Float.X) {
	class(self).SetExpandMargin(2, float64(value))
}

func (self Instance) ExpandMarginBottom() Float.X {
	return Float.X(Float.X(class(self).GetExpandMargin(3)))
}

func (self Instance) SetExpandMarginBottom(value Float.X) {
	class(self).SetExpandMargin(3, float64(value))
}

func (self Instance) AxisStretchHorizontal() gdclass.StyleBoxTextureAxisStretchMode {
	return gdclass.StyleBoxTextureAxisStretchMode(class(self).GetHAxisStretchMode())
}

func (self Instance) SetAxisStretchHorizontal(value gdclass.StyleBoxTextureAxisStretchMode) {
	class(self).SetHAxisStretchMode(value)
}

func (self Instance) AxisStretchVertical() gdclass.StyleBoxTextureAxisStretchMode {
	return gdclass.StyleBoxTextureAxisStretchMode(class(self).GetVAxisStretchMode())
}

func (self Instance) SetAxisStretchVertical(value gdclass.StyleBoxTextureAxisStretchMode) {
	class(self).SetVAxisStretchMode(value)
}

func (self Instance) RegionRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRegionRect())
}

func (self Instance) SetRegionRect(value Rect2.PositionSize) {
	class(self).SetRegionRect(Rect2.PositionSize(value))
}

func (self Instance) ModulateColor() Color.RGBA {
	return Color.RGBA(class(self).GetModulate())
}

func (self Instance) SetModulateColor(value Color.RGBA) {
	class(self).SetModulate(Color.RGBA(value))
}

func (self Instance) DrawCenter() bool {
	return bool(class(self).IsDrawCenterEnabled())
}

func (self Instance) SetDrawCenter(value bool) {
	class(self).SetDrawCenter(value)
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture2D) { //gd:StyleBoxTexture.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture2D { //gd:StyleBoxTexture.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the margin to [param size] pixels for the specified [enum Side].
*/
//go:nosplit
func (self class) SetTextureMargin(margin Side, size float64) { //gd:StyleBoxTexture.set_texture_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_texture_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the margin to [param size] pixels for all sides.
*/
//go:nosplit
func (self class) SetTextureMarginAll(size float64) { //gd:StyleBoxTexture.set_texture_margin_all
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_texture_margin_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the margin size of the specified [enum Side].
*/
//go:nosplit
func (self class) GetTextureMargin(margin Side) float64 { //gd:StyleBoxTexture.get_texture_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_texture_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the expand margin to [param size] pixels for the specified [enum Side].
*/
//go:nosplit
func (self class) SetExpandMargin(margin Side, size float64) { //gd:StyleBoxTexture.set_expand_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_expand_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the expand margin to [param size] pixels for all sides.
*/
//go:nosplit
func (self class) SetExpandMarginAll(size float64) { //gd:StyleBoxTexture.set_expand_margin_all
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_expand_margin_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the expand margin size of the specified [enum Side].
*/
//go:nosplit
func (self class) GetExpandMargin(margin Side) float64 { //gd:StyleBoxTexture.get_expand_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_expand_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegionRect(region Rect2.PositionSize) { //gd:StyleBoxTexture.set_region_rect
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_region_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRegionRect() Rect2.PositionSize { //gd:StyleBoxTexture.get_region_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[Rect2.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_region_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawCenter(enable bool) { //gd:StyleBoxTexture.set_draw_center
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_draw_center, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawCenterEnabled() bool { //gd:StyleBoxTexture.is_draw_center_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_is_draw_center_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModulate(color Color.RGBA) { //gd:StyleBoxTexture.set_modulate
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModulate() Color.RGBA { //gd:StyleBoxTexture.get_modulate
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHAxisStretchMode(mode gdclass.StyleBoxTextureAxisStretchMode) { //gd:StyleBoxTexture.set_h_axis_stretch_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHAxisStretchMode() gdclass.StyleBoxTextureAxisStretchMode { //gd:StyleBoxTexture.get_h_axis_stretch_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.StyleBoxTextureAxisStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVAxisStretchMode(mode gdclass.StyleBoxTextureAxisStretchMode) { //gd:StyleBoxTexture.set_v_axis_stretch_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_set_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVAxisStretchMode() gdclass.StyleBoxTextureAxisStretchMode { //gd:StyleBoxTexture.get_v_axis_stretch_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.StyleBoxTextureAxisStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxTexture.Bind_get_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StyleBox.Advanced(self.AsStyleBox()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StyleBox.Instance(self.AsStyleBox()), name)
	}
}
func init() {
	gdclass.Register("StyleBoxTexture", func(ptr gd.Object) any {
		return [1]gdclass.StyleBoxTexture{*(*gdclass.StyleBoxTexture)(unsafe.Pointer(&ptr))}
	})
}

type AxisStretchMode = gdclass.StyleBoxTextureAxisStretchMode //gd:StyleBoxTexture.AxisStretchMode

const (
	/*Stretch the stylebox's texture. This results in visible distortion unless the texture size matches the stylebox's size perfectly.*/
	AxisStretchModeStretch AxisStretchMode = 0
	/*Repeats the stylebox's texture to match the stylebox's size according to the nine-patch system.*/
	AxisStretchModeTile AxisStretchMode = 1
	/*Repeats the stylebox's texture to match the stylebox's size according to the nine-patch system. Unlike [constant AXIS_STRETCH_MODE_TILE], the texture may be slightly stretched to make the nine-patch texture tile seamlessly.*/
	AxisStretchModeTileFit AxisStretchMode = 2
)

type Side int

const (
	/*Left side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideLeft Side = 0
	/*Top side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideTop Side = 1
	/*Right side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideRight Side = 2
	/*Bottom side, usually used for [Control] or [StyleBox]-derived classes.*/
	SideBottom Side = 3
)
