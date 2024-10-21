package Gradient

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This resource describes a color transition by defining a set of colored points and how to interpolate between them.
See also [Curve] which supports more complex easing methods, but does not support colors.

*/
type Simple [1]classdb.Gradient
func (self Simple) AddPoint(offset float64, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddPoint(gd.Float(offset), color)
}
func (self Simple) RemovePoint(point int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemovePoint(gd.Int(point))
}
func (self Simple) SetOffset(point int, offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOffset(gd.Int(point), gd.Float(offset))
}
func (self Simple) GetOffset(point int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetOffset(gd.Int(point))))
}
func (self Simple) Reverse() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Reverse()
}
func (self Simple) SetColor(point int, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColor(gd.Int(point), color)
}
func (self Simple) GetColor(point int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetColor(gd.Int(point)))
}
func (self Simple) Sample(offset float64) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).Sample(gd.Float(offset)))
}
func (self Simple) GetPointCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPointCount()))
}
func (self Simple) SetOffsets(offsets gd.PackedFloat32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOffsets(offsets)
}
func (self Simple) GetOffsets() gd.PackedFloat32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedFloat32Array(Expert(self).GetOffsets(gc))
}
func (self Simple) SetColors(colors gd.PackedColorArray) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetColors(colors)
}
func (self Simple) GetColors() gd.PackedColorArray {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedColorArray(Expert(self).GetColors(gc))
}
func (self Simple) SetInterpolationMode(interpolation_mode classdb.GradientInterpolationMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInterpolationMode(interpolation_mode)
}
func (self Simple) GetInterpolationMode() classdb.GradientInterpolationMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GradientInterpolationMode(Expert(self).GetInterpolationMode())
}
func (self Simple) SetInterpolationColorSpace(interpolation_color_space classdb.GradientColorSpace) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetInterpolationColorSpace(interpolation_color_space)
}
func (self Simple) GetInterpolationColorSpace() classdb.GradientColorSpace {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.GradientColorSpace(Expert(self).GetInterpolationColorSpace())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Gradient
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

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

//go:nosplit
func (self class) AsGradient() Expert { return self[0].AsGradient() }


//go:nosplit
func (self Simple) AsGradient() Simple { return self[0].AsGradient() }


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
func init() {classdb.Register("Gradient", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
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
