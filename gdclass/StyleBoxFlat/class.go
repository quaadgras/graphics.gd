package StyleBoxFlat

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/StyleBox"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Go [1]classdb.StyleBoxFlat

/*
Sets the border width to [param width] pixels for all sides.
*/
func (self Go) SetBorderWidthAll(width int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderWidthAll(gd.Int(width))
}

/*
Returns the smallest border width out of all four borders.
*/
func (self Go) GetBorderWidthMin() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetBorderWidthMin()))
}

/*
Sets the corner radius to [param radius] pixels for all corners.
*/
func (self Go) SetCornerRadiusAll(radius int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCornerRadiusAll(gd.Int(radius))
}

/*
Sets the expand margin to [param size] pixels for all sides.
*/
func (self Go) SetExpandMarginAll(size float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpandMarginAll(gd.Float(size))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.StyleBoxFlat
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("StyleBoxFlat"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BgColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetBgColor())
}

func (self Go) SetBgColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBgColor(value)
}

func (self Go) DrawCenter() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawCenterEnabled())
}

func (self Go) SetDrawCenter(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawCenter(value)
}

func (self Go) Skew() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetSkew())
}

func (self Go) SetSkew(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSkew(value)
}

func (self Go) BorderWidthLeft() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBorderWidth(0)))
}

func (self Go) SetBorderWidthLeft(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderWidth(0, gd.Int(value))
}

func (self Go) BorderWidthTop() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBorderWidth(1)))
}

func (self Go) SetBorderWidthTop(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderWidth(1, gd.Int(value))
}

func (self Go) BorderWidthRight() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBorderWidth(2)))
}

func (self Go) SetBorderWidthRight(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderWidth(2, gd.Int(value))
}

func (self Go) BorderWidthBottom() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetBorderWidth(3)))
}

func (self Go) SetBorderWidthBottom(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderWidth(3, gd.Int(value))
}

func (self Go) BorderColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetBorderColor())
}

func (self Go) SetBorderColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderColor(value)
}

func (self Go) BorderBlend() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetBorderBlend())
}

func (self Go) SetBorderBlend(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBorderBlend(value)
}

func (self Go) CornerRadiusTopLeft() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCornerRadius(0)))
}

func (self Go) SetCornerRadiusTopLeft(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCornerRadius(0, gd.Int(value))
}

func (self Go) CornerRadiusTopRight() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCornerRadius(1)))
}

func (self Go) SetCornerRadiusTopRight(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCornerRadius(1, gd.Int(value))
}

func (self Go) CornerRadiusBottomRight() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCornerRadius(2)))
}

func (self Go) SetCornerRadiusBottomRight(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCornerRadius(2, gd.Int(value))
}

func (self Go) CornerRadiusBottomLeft() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCornerRadius(3)))
}

func (self Go) SetCornerRadiusBottomLeft(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCornerRadius(3, gd.Int(value))
}

func (self Go) CornerDetail() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetCornerDetail()))
}

func (self Go) SetCornerDetail(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetCornerDetail(gd.Int(value))
}

func (self Go) ExpandMarginLeft() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetExpandMargin(0)))
}

func (self Go) SetExpandMarginLeft(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpandMargin(0, gd.Float(value))
}

func (self Go) ExpandMarginTop() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetExpandMargin(1)))
}

func (self Go) SetExpandMarginTop(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpandMargin(1, gd.Float(value))
}

func (self Go) ExpandMarginRight() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetExpandMargin(2)))
}

func (self Go) SetExpandMarginRight(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpandMargin(2, gd.Float(value))
}

func (self Go) ExpandMarginBottom() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetExpandMargin(3)))
}

func (self Go) SetExpandMarginBottom(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetExpandMargin(3, gd.Float(value))
}

func (self Go) ShadowColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetShadowColor())
}

func (self Go) SetShadowColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadowColor(value)
}

func (self Go) ShadowSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetShadowSize()))
}

func (self Go) SetShadowSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadowSize(gd.Int(value))
}

func (self Go) ShadowOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Vector2(class(self).GetShadowOffset())
}

func (self Go) SetShadowOffset(value gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetShadowOffset(value)
}

func (self Go) AntiAliasing() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsAntiAliased())
}

func (self Go) SetAntiAliasing(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAntiAliased(value)
}

func (self Go) AntiAliasingSize() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetAaSize()))
}

func (self Go) SetAntiAliasingSize(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetAaSize(gd.Float(value))
}

//go:nosplit
func (self class) SetBgColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBgColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_bg_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBorderColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBorderColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_border_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the border width to [param width] pixels for all sides.
*/
//go:nosplit
func (self class) SetBorderWidthAll(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_border_width_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the smallest border width out of all four borders.
*/
//go:nosplit
func (self class) GetBorderWidthMin() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_border_width_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the specified [enum Side]'s border width to [param width] pixels.
*/
//go:nosplit
func (self class) SetBorderWidth(margin gd.Side, width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_border_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the specified [enum Side]'s border width.
*/
//go:nosplit
func (self class) GetBorderWidth(margin gd.Side) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_border_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBorderBlend(blend bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, blend)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_border_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBorderBlend() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_border_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the corner radius to [param radius] pixels for all corners.
*/
//go:nosplit
func (self class) SetCornerRadiusAll(radius gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_corner_radius_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the corner radius to [param radius] pixels for the given [param corner]. See [enum Corner] for possible values.
*/
//go:nosplit
func (self class) SetCornerRadius(corner gd.Corner, radius gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, corner)
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_corner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the given [param corner]'s radius. See [enum Corner] for possible values.
*/
//go:nosplit
func (self class) GetCornerRadius(corner gd.Corner) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, corner)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_corner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the expand margin to [param size] pixels for the specified [enum Side].
*/
//go:nosplit
func (self class) SetExpandMargin(margin gd.Side, size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_expand_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the expand margin to [param size] pixels for all sides.
*/
//go:nosplit
func (self class) SetExpandMarginAll(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_expand_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the size of the specified [enum Side]'s expand margin.
*/
//go:nosplit
func (self class) GetExpandMargin(margin gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_expand_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawCenter(draw_center bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, draw_center)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_draw_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawCenterEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_is_draw_center_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkew(skew gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, skew)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_skew, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkew() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_skew, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowColor(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_shadow_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowColor() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_shadow_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_shadow_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_shadow_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetShadowOffset(offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_shadow_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShadowOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_shadow_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAntiAliased(anti_aliased bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, anti_aliased)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_anti_aliased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAntiAliased() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_is_anti_aliased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAaSize(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_aa_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAaSize() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_aa_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCornerDetail(detail gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, detail)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_set_corner_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCornerDetail() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxFlat.Bind_get_corner_detail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsStyleBoxFlat() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStyleBoxFlat() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsStyleBox() StyleBox.GD { return *((*StyleBox.GD)(unsafe.Pointer(&self))) }
func (self Go) AsStyleBox() StyleBox.Go { return *((*StyleBox.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsStyleBox(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsStyleBox(), name)
	}
}
func init() {classdb.Register("StyleBoxFlat", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
