package Curve

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
This resource describes a mathematical curve by defining a set of points and tangents at each point. By default, it ranges between [code]0[/code] and [code]1[/code] on the Y axis and positions points relative to the [code]0.5[/code] Y position.
See also [Gradient] which is designed for color interpolation. See also [Curve2D] and [Curve3D].

*/
type Simple [1]classdb.Curve
func (self Simple) GetPointCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPointCount()))
}
func (self Simple) SetPointCount(count int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointCount(gd.Int(count))
}
func (self Simple) AddPoint(position gd.Vector2, left_tangent float64, right_tangent float64, left_mode classdb.CurveTangentMode, right_mode classdb.CurveTangentMode) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).AddPoint(position, gd.Float(left_tangent), gd.Float(right_tangent), left_mode, right_mode)))
}
func (self Simple) RemovePoint(index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemovePoint(gd.Int(index))
}
func (self Simple) ClearPoints() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearPoints()
}
func (self Simple) GetPointPosition(index int) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetPointPosition(gd.Int(index)))
}
func (self Simple) SetPointValue(index int, y float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointValue(gd.Int(index), gd.Float(y))
}
func (self Simple) SetPointOffset(index int, offset float64) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).SetPointOffset(gd.Int(index), gd.Float(offset))))
}
func (self Simple) Sample(offset float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).Sample(gd.Float(offset))))
}
func (self Simple) SampleBaked(offset float64) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).SampleBaked(gd.Float(offset))))
}
func (self Simple) GetPointLeftTangent(index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPointLeftTangent(gd.Int(index))))
}
func (self Simple) GetPointRightTangent(index int) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetPointRightTangent(gd.Int(index))))
}
func (self Simple) GetPointLeftMode(index int) classdb.CurveTangentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CurveTangentMode(Expert(self).GetPointLeftMode(gd.Int(index)))
}
func (self Simple) GetPointRightMode(index int) classdb.CurveTangentMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CurveTangentMode(Expert(self).GetPointRightMode(gd.Int(index)))
}
func (self Simple) SetPointLeftTangent(index int, tangent float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointLeftTangent(gd.Int(index), gd.Float(tangent))
}
func (self Simple) SetPointRightTangent(index int, tangent float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointRightTangent(gd.Int(index), gd.Float(tangent))
}
func (self Simple) SetPointLeftMode(index int, mode classdb.CurveTangentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointLeftMode(gd.Int(index), mode)
}
func (self Simple) SetPointRightMode(index int, mode classdb.CurveTangentMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPointRightMode(gd.Int(index), mode)
}
func (self Simple) GetMinValue() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMinValue()))
}
func (self Simple) SetMinValue(min float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMinValue(gd.Float(min))
}
func (self Simple) GetMaxValue() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMaxValue()))
}
func (self Simple) SetMaxValue(max float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaxValue(gd.Float(max))
}
func (self Simple) CleanDupes() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).CleanDupes()
}
func (self Simple) Bake() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Bake()
}
func (self Simple) GetBakeResolution() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetBakeResolution()))
}
func (self Simple) SetBakeResolution(resolution int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBakeResolution(gd.Int(resolution))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Curve
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) GetPointCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPointCount(count gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a point to the curve. For each side, if the [code]*_mode[/code] is [constant TANGENT_LINEAR], the [code]*_tangent[/code] angle (in degrees) uses the slope of the curve halfway to the adjacent point. Allows custom assignments to the [code]*_tangent[/code] angle if [code]*_mode[/code] is set to [constant TANGENT_FREE].
*/
//go:nosplit
func (self class) AddPoint(position gd.Vector2, left_tangent gd.Float, right_tangent gd.Float, left_mode classdb.CurveTangentMode, right_mode classdb.CurveTangentMode) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, left_tangent)
	callframe.Arg(frame, right_tangent)
	callframe.Arg(frame, left_mode)
	callframe.Arg(frame, right_mode)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the point at [param index] from the curve.
*/
//go:nosplit
func (self class) RemovePoint(index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all points from the curve.
*/
//go:nosplit
func (self class) ClearPoints()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_clear_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the curve coordinates for the point at [param index].
*/
//go:nosplit
func (self class) GetPointPosition(index gd.Int) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Assigns the vertical position [param y] to the point at [param index].
*/
//go:nosplit
func (self class) SetPointValue(index gd.Int, y gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, y)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_point_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the offset from [code]0.5[/code].
*/
//go:nosplit
func (self class) SetPointOffset(index gd.Int, offset gd.Float) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_point_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve.
*/
//go:nosplit
func (self class) Sample(offset gd.Float) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_sample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve using the baked cache. Bakes the curve's points if not already baked.
*/
//go:nosplit
func (self class) SampleBaked(offset gd.Float) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_sample_baked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the left tangent angle (in degrees) for the point at [param index].
*/
//go:nosplit
func (self class) GetPointLeftTangent(index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_point_left_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the right tangent angle (in degrees) for the point at [param index].
*/
//go:nosplit
func (self class) GetPointRightTangent(index gd.Int) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_point_right_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the left [enum TangentMode] for the point at [param index].
*/
//go:nosplit
func (self class) GetPointLeftMode(index gd.Int) classdb.CurveTangentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[classdb.CurveTangentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_point_left_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the right [enum TangentMode] for the point at [param index].
*/
//go:nosplit
func (self class) GetPointRightMode(index gd.Int) classdb.CurveTangentMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[classdb.CurveTangentMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_point_right_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the left tangent angle for the point at [param index] to [param tangent].
*/
//go:nosplit
func (self class) SetPointLeftTangent(index gd.Int, tangent gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, tangent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_point_left_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the right tangent angle for the point at [param index] to [param tangent].
*/
//go:nosplit
func (self class) SetPointRightTangent(index gd.Int, tangent gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, tangent)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_point_right_tangent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the left [enum TangentMode] for the point at [param index] to [param mode].
*/
//go:nosplit
func (self class) SetPointLeftMode(index gd.Int, mode classdb.CurveTangentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_point_left_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the right [enum TangentMode] for the point at [param index] to [param mode].
*/
//go:nosplit
func (self class) SetPointRightMode(index gd.Int, mode classdb.CurveTangentMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_point_right_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinValue() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_min_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinValue(min gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, min)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_min_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxValue() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_max_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxValue(max gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_max_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes duplicate points, i.e. points that are less than 0.00001 units (engine epsilon value) away from their neighbor on the curve.
*/
//go:nosplit
func (self class) CleanDupes()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_clean_dupes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Recomputes the baked cache of points for the curve.
*/
//go:nosplit
func (self class) Bake()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_bake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBakeResolution() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_get_bake_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBakeResolution(resolution gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, resolution)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Curve.Bind_set_bake_resolution, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsCurve() Expert { return self[0].AsCurve() }


//go:nosplit
func (self Simple) AsCurve() Simple { return self[0].AsCurve() }


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
func init() {classdb.Register("Curve", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TangentMode = classdb.CurveTangentMode

const (
/*The tangent on this side of the point is user-defined.*/
	TangentFree TangentMode = 0
/*The curve calculates the tangent on this side of the point as the slope halfway towards the adjacent point.*/
	TangentLinear TangentMode = 1
/*The total number of available tangent modes.*/
	TangentModeCount TangentMode = 2
)
