package StyleBoxTexture

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
A texture-based nine-patch [StyleBox], in a way similar to [NinePatchRect]. This stylebox performs a 3×3 scaling of a texture, where only the center cell is fully stretched. This makes it possible to design bordered styles regardless of the stylebox's size.

*/
type Go [1]classdb.StyleBoxTexture

/*
Sets the margin to [param size] pixels for all sides.
*/
func (self Go) SetTextureMarginAll(size float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureMarginAll(gd.Float(size))
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
type class [1]classdb.StyleBoxTexture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("StyleBoxTexture"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) Texture() gdclass.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.Texture2D(class(self).GetTexture(gc))
}

func (self Go) SetTexture(value gdclass.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTexture(value)
}

func (self Go) TextureMarginLeft() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTextureMargin(0)))
}

func (self Go) SetTextureMarginLeft(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureMargin(0, gd.Float(value))
}

func (self Go) TextureMarginTop() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTextureMargin(1)))
}

func (self Go) SetTextureMarginTop(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureMargin(1, gd.Float(value))
}

func (self Go) TextureMarginRight() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTextureMargin(2)))
}

func (self Go) SetTextureMarginRight(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureMargin(2, gd.Float(value))
}

func (self Go) TextureMarginBottom() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetTextureMargin(3)))
}

func (self Go) SetTextureMarginBottom(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTextureMargin(3, gd.Float(value))
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

func (self Go) AxisStretchHorizontal() classdb.StyleBoxTextureAxisStretchMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.StyleBoxTextureAxisStretchMode(class(self).GetHAxisStretchMode())
}

func (self Go) SetAxisStretchHorizontal(value classdb.StyleBoxTextureAxisStretchMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHAxisStretchMode(value)
}

func (self Go) AxisStretchVertical() classdb.StyleBoxTextureAxisStretchMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.StyleBoxTextureAxisStretchMode(class(self).GetVAxisStretchMode())
}

func (self Go) SetAxisStretchVertical(value classdb.StyleBoxTextureAxisStretchMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVAxisStretchMode(value)
}

func (self Go) RegionRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Rect2(class(self).GetRegionRect())
}

func (self Go) SetRegionRect(value gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRegionRect(value)
}

func (self Go) ModulateColor() gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Color(class(self).GetModulate())
}

func (self Go) SetModulateColor(value gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetModulate(value)
}

func (self Go) DrawCenter() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsDrawCenterEnabled())
}

func (self Go) SetDrawCenter(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetDrawCenter(value)
}

//go:nosplit
func (self class) SetTexture(texture gdclass.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexture(ctx gd.Lifetime) gdclass.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sets the margin to [param size] pixels for the specified [enum Side].
*/
//go:nosplit
func (self class) SetTextureMargin(margin gd.Side, size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_texture_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the margin to [param size] pixels for all sides.
*/
//go:nosplit
func (self class) SetTextureMarginAll(size gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_texture_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the margin size of the specified [enum Side].
*/
//go:nosplit
func (self class) GetTextureMargin(margin gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_get_texture_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_expand_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_expand_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the expand margin size of the specified [enum Side].
*/
//go:nosplit
func (self class) GetExpandMargin(margin gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_get_expand_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRegionRect(region gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRegionRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_get_region_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDrawCenter(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_draw_center, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDrawCenterEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_is_draw_center_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetModulate(color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetModulate() gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHAxisStretchMode(mode classdb.StyleBoxTextureAxisStretchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHAxisStretchMode() classdb.StyleBoxTextureAxisStretchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.StyleBoxTextureAxisStretchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_get_h_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVAxisStretchMode(mode classdb.StyleBoxTextureAxisStretchMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_set_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVAxisStretchMode() classdb.StyleBoxTextureAxisStretchMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.StyleBoxTextureAxisStretchMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBoxTexture.Bind_get_v_axis_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsStyleBoxTexture() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStyleBoxTexture() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("StyleBoxTexture", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type AxisStretchMode = classdb.StyleBoxTextureAxisStretchMode

const (
/*Stretch the stylebox's texture. This results in visible distortion unless the texture size matches the stylebox's size perfectly.*/
	AxisStretchModeStretch AxisStretchMode = 0
/*Repeats the stylebox's texture to match the stylebox's size according to the nine-patch system.*/
	AxisStretchModeTile AxisStretchMode = 1
/*Repeats the stylebox's texture to match the stylebox's size according to the nine-patch system. Unlike [constant AXIS_STRETCH_MODE_TILE], the texture may be slightly stretched to make the nine-patch texture tile seamlessly.*/
	AxisStretchModeTileFit AxisStretchMode = 2
)
