package Gradient

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource describes a color transition by defining a set of colored points and how to interpolate between them.
See also [Curve] which supports more complex easing methods, but does not support colors.

*/
type Go [1]classdb.Gradient

/*
Adds the specified color to the gradient, with the specified offset.
*/
func (self Go) AddPoint(offset float64, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddPoint(gd.Float(offset), color)
}

/*
Removes the color at index [param point].
*/
func (self Go) RemovePoint(point int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemovePoint(gd.Int(point))
}

/*
Sets the offset for the gradient color at index [param point].
*/
func (self Go) SetOffset(point int, offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOffset(gd.Int(point), gd.Float(offset))
}

/*
Returns the offset of the gradient color at index [param point].
*/
func (self Go) GetOffset(point int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetOffset(gd.Int(point))))
}

/*
Reverses/mirrors the gradient.
[b]Note:[/b] This method mirrors all points around the middle of the gradient, which may produce unexpected results when [member interpolation_mode] is set to [constant GRADIENT_INTERPOLATE_CONSTANT].
*/
func (self Go) Reverse() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Reverse()
}

/*
Sets the color of the gradient color at index [param point].
*/
func (self Go) SetColor(point int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColor(gd.Int(point), color)
}

/*
Returns the color of the gradient color at index [param point].
*/
func (self Go) GetColor(point int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).GetColor(gd.Int(point)))
}

/*
Returns the interpolated color specified by [param offset].
*/
func (self Go) Sample(offset float64) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(class(self).Sample(gd.Float(offset)))
}

/*
Returns the number of colors in the gradient.
*/
func (self Go) GetPointCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPointCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Gradient
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("Gradient"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) InterpolationMode() classdb.GradientInterpolationMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GradientInterpolationMode(class(self).GetInterpolationMode())
}

func (self Go) SetInterpolationMode(value classdb.GradientInterpolationMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInterpolationMode(value)
}

func (self Go) InterpolationColorSpace() classdb.GradientColorSpace {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.GradientColorSpace(class(self).GetInterpolationColorSpace())
}

func (self Go) SetInterpolationColorSpace(value classdb.GradientColorSpace) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInterpolationColorSpace(value)
}

func (self Go) Offsets() []float32 {
	gc := gd.GarbageCollector(); _ = gc
		return []float32(class(self).GetOffsets(gc).AsSlice())
}

func (self Go) SetOffsets(value []float32) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetOffsets(gc.PackedFloat32Slice(value))
}

func (self Go) Colors() []gd.Color {
	gc := gd.GarbageCollector(); _ = gc
		return []gd.Color(class(self).GetColors(gc).AsSlice())
}

func (self Go) SetColors(value []gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetColors(gc.PackedColorSlice(value))
}

/*
Adds the specified color to the gradient, with the specified offset.
*/
//go:nosplit
func (self class) AddPoint(offset gd.Float, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the color at index [param point].
*/
//go:nosplit
func (self class) RemovePoint(point gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the offset for the gradient color at index [param point].
*/
//go:nosplit
func (self class) SetOffset(point gd.Int, offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the offset of the gradient color at index [param point].
*/
//go:nosplit
func (self class) GetOffset(point gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Reverses/mirrors the gradient.
[b]Note:[/b] This method mirrors all points around the middle of the gradient, which may produce unexpected results when [member interpolation_mode] is set to [constant GRADIENT_INTERPOLATE_CONSTANT].
*/
//go:nosplit
func (self class) Reverse()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_reverse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the color of the gradient color at index [param point].
*/
//go:nosplit
func (self class) SetColor(point gd.Int, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the color of the gradient color at index [param point].
*/
//go:nosplit
func (self class) GetColor(point gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the interpolated color specified by [param offset].
*/
//go:nosplit
func (self class) Sample(offset gd.Float) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_sample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of colors in the gradient.
*/
//go:nosplit
func (self class) GetPointCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOffsets(offsets gd.PackedFloat32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(offsets))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_set_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffsets(ctx gd.Lifetime) gd.PackedFloat32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_get_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColors(colors gd.PackedColorArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(colors))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_set_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColors(ctx gd.Lifetime) gd.PackedColorArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_get_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedColorArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInterpolationMode(interpolation_mode classdb.GradientInterpolationMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interpolation_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_set_interpolation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInterpolationMode() classdb.GradientInterpolationMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientInterpolationMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_get_interpolation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInterpolationColorSpace(interpolation_color_space classdb.GradientColorSpace)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interpolation_color_space)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_set_interpolation_color_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInterpolationColorSpace() classdb.GradientColorSpace {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientColorSpace](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Gradient.Bind_get_interpolation_color_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGradient() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGradient() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("Gradient", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type InterpolationMode = classdb.GradientInterpolationMode

const (
/*Linear interpolation.*/
	GradientInterpolateLinear InterpolationMode = 0
/*Constant interpolation, color changes abruptly at each point and stays uniform between. This might cause visible aliasing when used for a gradient texture in some cases.*/
	GradientInterpolateConstant InterpolationMode = 1
/*Cubic interpolation.*/
	GradientInterpolateCubic InterpolationMode = 2
)
type ColorSpace = classdb.GradientColorSpace

const (
/*sRGB color space.*/
	GradientColorSpaceSrgb ColorSpace = 0
/*Linear sRGB color space.*/
	GradientColorSpaceLinearSrgb ColorSpace = 1
/*[url=https://bottosson.github.io/posts/oklab/]Oklab[/url] color space. This color space provides a smooth and uniform-looking transition between colors.*/
	GradientColorSpaceOklab ColorSpace = 2
)
