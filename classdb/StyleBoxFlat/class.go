// Package StyleBoxFlat provides methods for working with StyleBoxFlat object instances.
package StyleBoxFlat

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
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"

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
By configuring various properties of this style box, you can achieve many common looks without the need of a texture. This includes optionally rounded borders, antialiasing, shadows, and skew.
Setting corner radius to high values is allowed. As soon as corners overlap, the stylebox will switch to a relative system:
[codeblock lang=text]
height = 30
corner_radius_top_left = 50
corner_radius_bottom_left = 100
[/codeblock]
The relative system now would take the 1:2 ratio of the two left corners to calculate the actual corner width. Both corners added will [b]never[/b] be more than the height. Result:
[codeblock lang=text]
corner_radius_top_left: 10
corner_radius_bottom_left: 20
[/codeblock]
*/
type Instance [1]gdclass.StyleBoxFlat

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStyleBoxFlat() Instance
}

/*
Sets the border width to [param width] pixels for all sides.
*/
func (self Instance) SetBorderWidthAll(width int) { //gd:StyleBoxFlat.set_border_width_all
	class(self).SetBorderWidthAll(int64(width))
}

/*
Returns the smallest border width out of all four borders.
*/
func (self Instance) GetBorderWidthMin() int { //gd:StyleBoxFlat.get_border_width_min
	return int(int(class(self).GetBorderWidthMin()))
}

/*
Sets the corner radius to [param radius] pixels for all corners.
*/
func (self Instance) SetCornerRadiusAll(radius int) { //gd:StyleBoxFlat.set_corner_radius_all
	class(self).SetCornerRadiusAll(int64(radius))
}

/*
Sets the expand margin to [param size] pixels for all sides.
*/
func (self Instance) SetExpandMarginAll(size Float.X) { //gd:StyleBoxFlat.set_expand_margin_all
	class(self).SetExpandMarginAll(float64(size))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StyleBoxFlat

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBoxFlat"))
	casted := Instance{*(*gdclass.StyleBoxFlat)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BgColor() Color.RGBA {
	return Color.RGBA(class(self).GetBgColor())
}

func (self Instance) SetBgColor(value Color.RGBA) {
	class(self).SetBgColor(Color.RGBA(value))
}

func (self Instance) DrawCenter() bool {
	return bool(class(self).IsDrawCenterEnabled())
}

func (self Instance) SetDrawCenter(value bool) {
	class(self).SetDrawCenter(value)
}

func (self Instance) Skew() Vector2.XY {
	return Vector2.XY(class(self).GetSkew())
}

func (self Instance) SetSkew(value Vector2.XY) {
	class(self).SetSkew(Vector2.XY(value))
}

func (self Instance) BorderWidthLeft() int {
	return int(int(class(self).GetBorderWidth(0)))
}

func (self Instance) SetBorderWidthLeft(value int) {
	class(self).SetBorderWidth(0, int64(value))
}

func (self Instance) BorderWidthTop() int {
	return int(int(class(self).GetBorderWidth(1)))
}

func (self Instance) SetBorderWidthTop(value int) {
	class(self).SetBorderWidth(1, int64(value))
}

func (self Instance) BorderWidthRight() int {
	return int(int(class(self).GetBorderWidth(2)))
}

func (self Instance) SetBorderWidthRight(value int) {
	class(self).SetBorderWidth(2, int64(value))
}

func (self Instance) BorderWidthBottom() int {
	return int(int(class(self).GetBorderWidth(3)))
}

func (self Instance) SetBorderWidthBottom(value int) {
	class(self).SetBorderWidth(3, int64(value))
}

func (self Instance) BorderColor() Color.RGBA {
	return Color.RGBA(class(self).GetBorderColor())
}

func (self Instance) SetBorderColor(value Color.RGBA) {
	class(self).SetBorderColor(Color.RGBA(value))
}

func (self Instance) BorderBlend() bool {
	return bool(class(self).GetBorderBlend())
}

func (self Instance) SetBorderBlend(value bool) {
	class(self).SetBorderBlend(value)
}

func (self Instance) CornerRadiusTopLeft() int {
	return int(int(class(self).GetCornerRadius(0)))
}

func (self Instance) SetCornerRadiusTopLeft(value int) {
	class(self).SetCornerRadius(0, int64(value))
}

func (self Instance) CornerRadiusTopRight() int {
	return int(int(class(self).GetCornerRadius(1)))
}

func (self Instance) SetCornerRadiusTopRight(value int) {
	class(self).SetCornerRadius(1, int64(value))
}

func (self Instance) CornerRadiusBottomRight() int {
	return int(int(class(self).GetCornerRadius(2)))
}

func (self Instance) SetCornerRadiusBottomRight(value int) {
	class(self).SetCornerRadius(2, int64(value))
}

func (self Instance) CornerRadiusBottomLeft() int {
	return int(int(class(self).GetCornerRadius(3)))
}

func (self Instance) SetCornerRadiusBottomLeft(value int) {
	class(self).SetCornerRadius(3, int64(value))
}

func (self Instance) CornerDetail() int {
	return int(int(class(self).GetCornerDetail()))
}

func (self Instance) SetCornerDetail(value int) {
	class(self).SetCornerDetail(int64(value))
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

func (self Instance) ShadowColor() Color.RGBA {
	return Color.RGBA(class(self).GetShadowColor())
}

func (self Instance) SetShadowColor(value Color.RGBA) {
	class(self).SetShadowColor(Color.RGBA(value))
}

func (self Instance) ShadowSize() int {
	return int(int(class(self).GetShadowSize()))
}

func (self Instance) SetShadowSize(value int) {
	class(self).SetShadowSize(int64(value))
}

func (self Instance) ShadowOffset() Vector2.XY {
	return Vector2.XY(class(self).GetShadowOffset())
}

func (self Instance) SetShadowOffset(value Vector2.XY) {
	class(self).SetShadowOffset(Vector2.XY(value))
}

func (self Instance) AntiAliasing() bool {
	return bool(class(self).IsAntiAliased())
}

func (self Instance) SetAntiAliasing(value bool) {
	class(self).SetAntiAliased(value)
}

func (self Instance) AntiAliasingSize() Float.X {
	return Float.X(Float.X(class(self).GetAaSize()))
}

func (self Instance) SetAntiAliasingSize(value Float.X) {
	class(self).SetAaSize(float64(value))
}

//go:nosplit
func (self class) SetBgColor(color Color.RGBA) { //gd:StyleBoxFlat.set_bg_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_bg_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBgColor() Color.RGBA { //gd:StyleBoxFlat.get_bg_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_bg_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderColor(color Color.RGBA) { //gd:StyleBoxFlat.set_border_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_border_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderColor() Color.RGBA { //gd:StyleBoxFlat.get_border_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_border_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the border width to [param width] pixels for all sides.
*/
//go:nosplit
func (self class) SetBorderWidthAll(width int64) { //gd:StyleBoxFlat.set_border_width_all
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_border_width_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the smallest border width out of all four borders.
*/
//go:nosplit
func (self class) GetBorderWidthMin() int64 { //gd:StyleBoxFlat.get_border_width_min
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_border_width_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the specified [enum Side]'s border width to [param width] pixels.
*/
//go:nosplit
func (self class) SetBorderWidth(margin Side, width int64) { //gd:StyleBoxFlat.set_border_width
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_border_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the specified [enum Side]'s border width.
*/
//go:nosplit
func (self class) GetBorderWidth(margin Side) int64 { //gd:StyleBoxFlat.get_border_width
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_border_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderBlend(blend bool) { //gd:StyleBoxFlat.set_border_blend
	var frame = callframe.New()
	callframe.Arg(frame, blend)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_border_blend, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderBlend() bool { //gd:StyleBoxFlat.get_border_blend
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_border_blend, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the corner radius to [param radius] pixels for all corners.
*/
//go:nosplit
func (self class) SetCornerRadiusAll(radius int64) { //gd:StyleBoxFlat.set_corner_radius_all
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_corner_radius_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the corner radius to [param radius] pixels for the given [param corner]. See [enum Corner] for possible values.
*/
//go:nosplit
func (self class) SetCornerRadius(corner Corner, radius int64) { //gd:StyleBoxFlat.set_corner_radius
	var frame = callframe.New()
	callframe.Arg(frame, corner)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_corner_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the given [param corner]'s radius. See [enum Corner] for possible values.
*/
//go:nosplit
func (self class) GetCornerRadius(corner Corner) int64 { //gd:StyleBoxFlat.get_corner_radius
	var frame = callframe.New()
	callframe.Arg(frame, corner)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_corner_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the expand margin to [param size] pixels for the specified [enum Side].
*/
//go:nosplit
func (self class) SetExpandMargin(margin Side, size float64) { //gd:StyleBoxFlat.set_expand_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_expand_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the expand margin to [param size] pixels for all sides.
*/
//go:nosplit
func (self class) SetExpandMarginAll(size float64) { //gd:StyleBoxFlat.set_expand_margin_all
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_expand_margin_all, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the size of the specified [enum Side]'s expand margin.
*/
//go:nosplit
func (self class) GetExpandMargin(margin Side) float64 { //gd:StyleBoxFlat.get_expand_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_expand_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawCenter(draw_center bool) { //gd:StyleBoxFlat.set_draw_center
	var frame = callframe.New()
	callframe.Arg(frame, draw_center)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_draw_center, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawCenterEnabled() bool { //gd:StyleBoxFlat.is_draw_center_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_is_draw_center_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkew(skew Vector2.XY) { //gd:StyleBoxFlat.set_skew
	var frame = callframe.New()
	callframe.Arg(frame, skew)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_skew, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkew() Vector2.XY { //gd:StyleBoxFlat.get_skew
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_skew, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowColor(color Color.RGBA) { //gd:StyleBoxFlat.set_shadow_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_shadow_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowColor() Color.RGBA { //gd:StyleBoxFlat.get_shadow_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_shadow_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowSize(size int64) { //gd:StyleBoxFlat.set_shadow_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_shadow_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowSize() int64 { //gd:StyleBoxFlat.get_shadow_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_shadow_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowOffset(offset Vector2.XY) { //gd:StyleBoxFlat.set_shadow_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_shadow_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowOffset() Vector2.XY { //gd:StyleBoxFlat.get_shadow_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_shadow_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAntiAliased(anti_aliased bool) { //gd:StyleBoxFlat.set_anti_aliased
	var frame = callframe.New()
	callframe.Arg(frame, anti_aliased)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_anti_aliased, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAntiAliased() bool { //gd:StyleBoxFlat.is_anti_aliased
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_is_anti_aliased, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAaSize(size float64) { //gd:StyleBoxFlat.set_aa_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_aa_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAaSize() float64 { //gd:StyleBoxFlat.get_aa_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_aa_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCornerDetail(detail int64) { //gd:StyleBoxFlat.set_corner_detail
	var frame = callframe.New()
	callframe.Arg(frame, detail)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_corner_detail, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCornerDetail() int64 { //gd:StyleBoxFlat.get_corner_detail
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_corner_detail, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsStyleBoxFlat() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStyleBoxFlat() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("StyleBoxFlat", func(ptr gd.Object) any {
		return [1]gdclass.StyleBoxFlat{*(*gdclass.StyleBoxFlat)(unsafe.Pointer(&ptr))}
	})
}

type Corner int

const (
	/*Top-left corner.*/
	CornerTopLeft Corner = 0
	/*Top-right corner.*/
	CornerTopRight Corner = 1
	/*Bottom-right corner.*/
	CornerBottomRight Corner = 2
	/*Bottom-left corner.*/
	CornerBottomLeft Corner = 3
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
