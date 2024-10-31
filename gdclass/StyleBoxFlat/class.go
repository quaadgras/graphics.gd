package StyleBoxFlat

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/StyleBox"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
By configuring various properties of this style box, you can achieve many common looks without the need of a texture. This includes optionally rounded borders, antialiasing, shadows, and skew.
Setting corner radius to high values is allowed. As soon as corners overlap, the stylebox will switch to a relative system.
[b]Example:[/b]
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
type Instance [1]classdb.StyleBoxFlat

/*
Sets the border width to [param width] pixels for all sides.
*/
func (self Instance) SetBorderWidthAll(width int) {
	class(self).SetBorderWidthAll(gd.Int(width))
}

/*
Returns the smallest border width out of all four borders.
*/
func (self Instance) GetBorderWidthMin() int {
	return int(int(class(self).GetBorderWidthMin()))
}

/*
Sets the corner radius to [param radius] pixels for all corners.
*/
func (self Instance) SetCornerRadiusAll(radius int) {
	class(self).SetCornerRadiusAll(gd.Int(radius))
}

/*
Sets the expand margin to [param size] pixels for all sides.
*/
func (self Instance) SetExpandMarginAll(size float64) {
	class(self).SetExpandMarginAll(gd.Float(size))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StyleBoxFlat

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBoxFlat"))
	return Instance{classdb.StyleBoxFlat(object)}
}

func (self Instance) BgColor() gd.Color {
	return gd.Color(class(self).GetBgColor())
}

func (self Instance) SetBgColor(value gd.Color) {
	class(self).SetBgColor(value)
}

func (self Instance) DrawCenter() bool {
	return bool(class(self).IsDrawCenterEnabled())
}

func (self Instance) SetDrawCenter(value bool) {
	class(self).SetDrawCenter(value)
}

func (self Instance) Skew() gd.Vector2 {
	return gd.Vector2(class(self).GetSkew())
}

func (self Instance) SetSkew(value gd.Vector2) {
	class(self).SetSkew(value)
}

func (self Instance) BorderWidthLeft() int {
	return int(int(class(self).GetBorderWidth(0)))
}

func (self Instance) SetBorderWidthLeft(value int) {
	class(self).SetBorderWidth(0, gd.Int(value))
}

func (self Instance) BorderWidthTop() int {
	return int(int(class(self).GetBorderWidth(1)))
}

func (self Instance) SetBorderWidthTop(value int) {
	class(self).SetBorderWidth(1, gd.Int(value))
}

func (self Instance) BorderWidthRight() int {
	return int(int(class(self).GetBorderWidth(2)))
}

func (self Instance) SetBorderWidthRight(value int) {
	class(self).SetBorderWidth(2, gd.Int(value))
}

func (self Instance) BorderWidthBottom() int {
	return int(int(class(self).GetBorderWidth(3)))
}

func (self Instance) SetBorderWidthBottom(value int) {
	class(self).SetBorderWidth(3, gd.Int(value))
}

func (self Instance) BorderColor() gd.Color {
	return gd.Color(class(self).GetBorderColor())
}

func (self Instance) SetBorderColor(value gd.Color) {
	class(self).SetBorderColor(value)
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
	class(self).SetCornerRadius(0, gd.Int(value))
}

func (self Instance) CornerRadiusTopRight() int {
	return int(int(class(self).GetCornerRadius(1)))
}

func (self Instance) SetCornerRadiusTopRight(value int) {
	class(self).SetCornerRadius(1, gd.Int(value))
}

func (self Instance) CornerRadiusBottomRight() int {
	return int(int(class(self).GetCornerRadius(2)))
}

func (self Instance) SetCornerRadiusBottomRight(value int) {
	class(self).SetCornerRadius(2, gd.Int(value))
}

func (self Instance) CornerRadiusBottomLeft() int {
	return int(int(class(self).GetCornerRadius(3)))
}

func (self Instance) SetCornerRadiusBottomLeft(value int) {
	class(self).SetCornerRadius(3, gd.Int(value))
}

func (self Instance) CornerDetail() int {
	return int(int(class(self).GetCornerDetail()))
}

func (self Instance) SetCornerDetail(value int) {
	class(self).SetCornerDetail(gd.Int(value))
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

func (self Instance) ShadowColor() gd.Color {
	return gd.Color(class(self).GetShadowColor())
}

func (self Instance) SetShadowColor(value gd.Color) {
	class(self).SetShadowColor(value)
}

func (self Instance) ShadowSize() int {
	return int(int(class(self).GetShadowSize()))
}

func (self Instance) SetShadowSize(value int) {
	class(self).SetShadowSize(gd.Int(value))
}

func (self Instance) ShadowOffset() gd.Vector2 {
	return gd.Vector2(class(self).GetShadowOffset())
}

func (self Instance) SetShadowOffset(value gd.Vector2) {
	class(self).SetShadowOffset(value)
}

func (self Instance) AntiAliasing() bool {
	return bool(class(self).IsAntiAliased())
}

func (self Instance) SetAntiAliasing(value bool) {
	class(self).SetAntiAliased(value)
}

func (self Instance) AntiAliasingSize() float64 {
	return float64(float64(class(self).GetAaSize()))
}

func (self Instance) SetAntiAliasingSize(value float64) {
	class(self).SetAaSize(gd.Float(value))
}

//go:nosplit
func (self class) SetBgColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBgColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the border width to [param width] pixels for all sides.
*/
//go:nosplit
func (self class) SetBorderWidthAll(width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_border_width_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the smallest border width out of all four borders.
*/
//go:nosplit
func (self class) GetBorderWidthMin() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_border_width_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the specified [enum Side]'s border width to [param width] pixels.
*/
//go:nosplit
func (self class) SetBorderWidth(margin gdconst.Side, width gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_border_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the specified [enum Side]'s border width.
*/
//go:nosplit
func (self class) GetBorderWidth(margin gdconst.Side) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_border_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderBlend(blend bool) {
	var frame = callframe.New()
	callframe.Arg(frame, blend)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_border_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderBlend() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_border_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the corner radius to [param radius] pixels for all corners.
*/
//go:nosplit
func (self class) SetCornerRadiusAll(radius gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_corner_radius_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the corner radius to [param radius] pixels for the given [param corner]. See [enum Corner] for possible values.
*/
//go:nosplit
func (self class) SetCornerRadius(corner gdconst.Corner, radius gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, corner)
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_corner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the given [param corner]'s radius. See [enum Corner] for possible values.
*/
//go:nosplit
func (self class) GetCornerRadius(corner gdconst.Corner) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, corner)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_corner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_expand_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_expand_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the size of the specified [enum Side]'s expand margin.
*/
//go:nosplit
func (self class) GetExpandMargin(margin gdconst.Side) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_expand_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawCenter(draw_center bool) {
	var frame = callframe.New()
	callframe.Arg(frame, draw_center)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_draw_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawCenterEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_is_draw_center_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkew(skew gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, skew)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_skew, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkew() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_skew, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_shadow_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_shadow_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowSize(size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_shadow_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_shadow_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_shadow_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_shadow_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAntiAliased(anti_aliased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, anti_aliased)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_anti_aliased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAntiAliased() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_is_anti_aliased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAaSize(size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_aa_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAaSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_aa_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCornerDetail(detail gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, detail)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_set_corner_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCornerDetail() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxFlat.Bind_get_corner_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	classdb.Register("StyleBoxFlat", func(ptr gd.Object) any { return classdb.StyleBoxFlat(ptr) })
}
