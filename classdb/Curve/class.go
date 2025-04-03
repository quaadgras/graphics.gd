// Package Curve provides methods for working with Curve object instances.
package Curve

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
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
This resource describes a mathematical curve by defining a set of points and tangents at each point. By default, it ranges between [code]0[/code] and [code]1[/code] on the X and Y axes, but these ranges can be changed.
Please note that many resources and nodes assume they are given [i]unit curves[/i]. A unit curve is a curve whose domain (the X axis) is between [code]0[/code] and [code]1[/code]. Some examples of unit curve usage are [member CPUParticles2D.angle_curve] and [member Line2D.width_curve].
*/
type Instance [1]gdclass.Curve
type Expanded [1]gdclass.Curve

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCurve() Instance
}

/*
Adds a point to the curve. For each side, if the [code]*_mode[/code] is [constant TANGENT_LINEAR], the [code]*_tangent[/code] angle (in degrees) uses the slope of the curve halfway to the adjacent point. Allows custom assignments to the [code]*_tangent[/code] angle if [code]*_mode[/code] is set to [constant TANGENT_FREE].
*/
func (self Instance) AddPoint(position Vector2.XY) int { //gd:Curve.add_point
	return int(int(Advanced(self).AddPoint(Vector2.XY(position), float64(0), float64(0), 0, 0)))
}

/*
Adds a point to the curve. For each side, if the [code]*_mode[/code] is [constant TANGENT_LINEAR], the [code]*_tangent[/code] angle (in degrees) uses the slope of the curve halfway to the adjacent point. Allows custom assignments to the [code]*_tangent[/code] angle if [code]*_mode[/code] is set to [constant TANGENT_FREE].
*/
func (self Expanded) AddPoint(position Vector2.XY, left_tangent Float.X, right_tangent Float.X, left_mode gdclass.CurveTangentMode, right_mode gdclass.CurveTangentMode) int { //gd:Curve.add_point
	return int(int(Advanced(self).AddPoint(Vector2.XY(position), float64(left_tangent), float64(right_tangent), left_mode, right_mode)))
}

/*
Removes the point at [param index] from the curve.
*/
func (self Instance) RemovePoint(index int) { //gd:Curve.remove_point
	Advanced(self).RemovePoint(int64(index))
}

/*
Removes all points from the curve.
*/
func (self Instance) ClearPoints() { //gd:Curve.clear_points
	Advanced(self).ClearPoints()
}

/*
Returns the curve coordinates for the point at [param index].
*/
func (self Instance) GetPointPosition(index int) Vector2.XY { //gd:Curve.get_point_position
	return Vector2.XY(Advanced(self).GetPointPosition(int64(index)))
}

/*
Assigns the vertical position [param y] to the point at [param index].
*/
func (self Instance) SetPointValue(index int, y Float.X) { //gd:Curve.set_point_value
	Advanced(self).SetPointValue(int64(index), float64(y))
}

/*
Sets the offset from [code]0.5[/code].
*/
func (self Instance) SetPointOffset(index int, offset Float.X) int { //gd:Curve.set_point_offset
	return int(int(Advanced(self).SetPointOffset(int64(index), float64(offset))))
}

/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve.
*/
func (self Instance) Sample(offset Float.X) Float.X { //gd:Curve.sample
	return Float.X(Float.X(Advanced(self).Sample(float64(offset))))
}

/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve using the baked cache. Bakes the curve's points if not already baked.
*/
func (self Instance) SampleBaked(offset Float.X) Float.X { //gd:Curve.sample_baked
	return Float.X(Float.X(Advanced(self).SampleBaked(float64(offset))))
}

/*
Returns the left tangent angle (in degrees) for the point at [param index].
*/
func (self Instance) GetPointLeftTangent(index int) Float.X { //gd:Curve.get_point_left_tangent
	return Float.X(Float.X(Advanced(self).GetPointLeftTangent(int64(index))))
}

/*
Returns the right tangent angle (in degrees) for the point at [param index].
*/
func (self Instance) GetPointRightTangent(index int) Float.X { //gd:Curve.get_point_right_tangent
	return Float.X(Float.X(Advanced(self).GetPointRightTangent(int64(index))))
}

/*
Returns the left [enum TangentMode] for the point at [param index].
*/
func (self Instance) GetPointLeftMode(index int) gdclass.CurveTangentMode { //gd:Curve.get_point_left_mode
	return gdclass.CurveTangentMode(Advanced(self).GetPointLeftMode(int64(index)))
}

/*
Returns the right [enum TangentMode] for the point at [param index].
*/
func (self Instance) GetPointRightMode(index int) gdclass.CurveTangentMode { //gd:Curve.get_point_right_mode
	return gdclass.CurveTangentMode(Advanced(self).GetPointRightMode(int64(index)))
}

/*
Sets the left tangent angle for the point at [param index] to [param tangent].
*/
func (self Instance) SetPointLeftTangent(index int, tangent Float.X) { //gd:Curve.set_point_left_tangent
	Advanced(self).SetPointLeftTangent(int64(index), float64(tangent))
}

/*
Sets the right tangent angle for the point at [param index] to [param tangent].
*/
func (self Instance) SetPointRightTangent(index int, tangent Float.X) { //gd:Curve.set_point_right_tangent
	Advanced(self).SetPointRightTangent(int64(index), float64(tangent))
}

/*
Sets the left [enum TangentMode] for the point at [param index] to [param mode].
*/
func (self Instance) SetPointLeftMode(index int, mode gdclass.CurveTangentMode) { //gd:Curve.set_point_left_mode
	Advanced(self).SetPointLeftMode(int64(index), mode)
}

/*
Sets the right [enum TangentMode] for the point at [param index] to [param mode].
*/
func (self Instance) SetPointRightMode(index int, mode gdclass.CurveTangentMode) { //gd:Curve.set_point_right_mode
	Advanced(self).SetPointRightMode(int64(index), mode)
}

/*
Returns the difference between [member min_value] and [member max_value].
*/
func (self Instance) GetValueRange() Float.X { //gd:Curve.get_value_range
	return Float.X(Float.X(Advanced(self).GetValueRange()))
}

/*
Returns the difference between [member min_domain] and [member max_domain].
*/
func (self Instance) GetDomainRange() Float.X { //gd:Curve.get_domain_range
	return Float.X(Float.X(Advanced(self).GetDomainRange()))
}

/*
Removes duplicate points, i.e. points that are less than 0.00001 units (engine epsilon value) away from their neighbor on the curve.
*/
func (self Instance) CleanDupes() { //gd:Curve.clean_dupes
	Advanced(self).CleanDupes()
}

/*
Recomputes the baked cache of points for the curve.
*/
func (self Instance) Bake() { //gd:Curve.bake
	Advanced(self).Bake()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Curve

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Curve"))
	casted := Instance{*(*gdclass.Curve)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) MinDomain() Float.X {
	return Float.X(Float.X(class(self).GetMinDomain()))
}

func (self Instance) SetMinDomain(value Float.X) {
	class(self).SetMinDomain(float64(value))
}

func (self Instance) MaxDomain() Float.X {
	return Float.X(Float.X(class(self).GetMaxDomain()))
}

func (self Instance) SetMaxDomain(value Float.X) {
	class(self).SetMaxDomain(float64(value))
}

func (self Instance) MinValue() Float.X {
	return Float.X(Float.X(class(self).GetMinValue()))
}

func (self Instance) SetMinValue(value Float.X) {
	class(self).SetMinValue(float64(value))
}

func (self Instance) MaxValue() Float.X {
	return Float.X(Float.X(class(self).GetMaxValue()))
}

func (self Instance) SetMaxValue(value Float.X) {
	class(self).SetMaxValue(float64(value))
}

func (self Instance) BakeResolution() int {
	return int(int(class(self).GetBakeResolution()))
}

func (self Instance) SetBakeResolution(value int) {
	class(self).SetBakeResolution(int64(value))
}

func (self Instance) PointCount() int {
	return int(int(class(self).GetPointCount()))
}

func (self Instance) SetPointCount(value int) {
	class(self).SetPointCount(int64(value))
}

//go:nosplit
func (self class) GetPointCount() int64 { //gd:Curve.get_point_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPointCount(count int64) { //gd:Curve.set_point_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a point to the curve. For each side, if the [code]*_mode[/code] is [constant TANGENT_LINEAR], the [code]*_tangent[/code] angle (in degrees) uses the slope of the curve halfway to the adjacent point. Allows custom assignments to the [code]*_tangent[/code] angle if [code]*_mode[/code] is set to [constant TANGENT_FREE].
*/
//go:nosplit
func (self class) AddPoint(position Vector2.XY, left_tangent float64, right_tangent float64, left_mode gdclass.CurveTangentMode, right_mode gdclass.CurveTangentMode) int64 { //gd:Curve.add_point
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, left_tangent)
	callframe.Arg(frame, right_tangent)
	callframe.Arg(frame, left_mode)
	callframe.Arg(frame, right_mode)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the point at [param index] from the curve.
*/
//go:nosplit
func (self class) RemovePoint(index int64) { //gd:Curve.remove_point
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all points from the curve.
*/
//go:nosplit
func (self class) ClearPoints() { //gd:Curve.clear_points
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_clear_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the curve coordinates for the point at [param index].
*/
//go:nosplit
func (self class) GetPointPosition(index int64) Vector2.XY { //gd:Curve.get_point_position
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Assigns the vertical position [param y] to the point at [param index].
*/
//go:nosplit
func (self class) SetPointValue(index int64, y float64) { //gd:Curve.set_point_value
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, y)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the offset from [code]0.5[/code].
*/
//go:nosplit
func (self class) SetPointOffset(index int64, offset float64) int64 { //gd:Curve.set_point_offset
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve.
*/
//go:nosplit
func (self class) Sample(offset float64) float64 { //gd:Curve.sample
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_sample, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve using the baked cache. Bakes the curve's points if not already baked.
*/
//go:nosplit
func (self class) SampleBaked(offset float64) float64 { //gd:Curve.sample_baked
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_sample_baked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the left tangent angle (in degrees) for the point at [param index].
*/
//go:nosplit
func (self class) GetPointLeftTangent(index int64) float64 { //gd:Curve.get_point_left_tangent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_left_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the right tangent angle (in degrees) for the point at [param index].
*/
//go:nosplit
func (self class) GetPointRightTangent(index int64) float64 { //gd:Curve.get_point_right_tangent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_right_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the left [enum TangentMode] for the point at [param index].
*/
//go:nosplit
func (self class) GetPointLeftMode(index int64) gdclass.CurveTangentMode { //gd:Curve.get_point_left_mode
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gdclass.CurveTangentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_left_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the right [enum TangentMode] for the point at [param index].
*/
//go:nosplit
func (self class) GetPointRightMode(index int64) gdclass.CurveTangentMode { //gd:Curve.get_point_right_mode
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gdclass.CurveTangentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_right_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the left tangent angle for the point at [param index] to [param tangent].
*/
//go:nosplit
func (self class) SetPointLeftTangent(index int64, tangent float64) { //gd:Curve.set_point_left_tangent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, tangent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_left_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the right tangent angle for the point at [param index] to [param tangent].
*/
//go:nosplit
func (self class) SetPointRightTangent(index int64, tangent float64) { //gd:Curve.set_point_right_tangent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, tangent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_right_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the left [enum TangentMode] for the point at [param index] to [param mode].
*/
//go:nosplit
func (self class) SetPointLeftMode(index int64, mode gdclass.CurveTangentMode) { //gd:Curve.set_point_left_mode
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_left_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the right [enum TangentMode] for the point at [param index] to [param mode].
*/
//go:nosplit
func (self class) SetPointRightMode(index int64, mode gdclass.CurveTangentMode) { //gd:Curve.set_point_right_mode
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_right_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinValue() float64 { //gd:Curve.get_min_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_min_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinValue(min float64) { //gd:Curve.set_min_value
	var frame = callframe.New()
	callframe.Arg(frame, min)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_min_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxValue() float64 { //gd:Curve.get_max_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_max_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxValue(max float64) { //gd:Curve.set_max_value
	var frame = callframe.New()
	callframe.Arg(frame, max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_max_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the difference between [member min_value] and [member max_value].
*/
//go:nosplit
func (self class) GetValueRange() float64 { //gd:Curve.get_value_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_value_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMinDomain() float64 { //gd:Curve.get_min_domain
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_min_domain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinDomain(min float64) { //gd:Curve.set_min_domain
	var frame = callframe.New()
	callframe.Arg(frame, min)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_min_domain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxDomain() float64 { //gd:Curve.get_max_domain
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_max_domain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxDomain(max float64) { //gd:Curve.set_max_domain
	var frame = callframe.New()
	callframe.Arg(frame, max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_max_domain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the difference between [member min_domain] and [member max_domain].
*/
//go:nosplit
func (self class) GetDomainRange() float64 { //gd:Curve.get_domain_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_domain_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes duplicate points, i.e. points that are less than 0.00001 units (engine epsilon value) away from their neighbor on the curve.
*/
//go:nosplit
func (self class) CleanDupes() { //gd:Curve.clean_dupes
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_clean_dupes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Recomputes the baked cache of points for the curve.
*/
//go:nosplit
func (self class) Bake() { //gd:Curve.bake
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_bake, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeResolution() int64 { //gd:Curve.get_bake_resolution
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_bake_resolution, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeResolution(resolution int64) { //gd:Curve.set_bake_resolution
	var frame = callframe.New()
	callframe.Arg(frame, resolution)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_bake_resolution, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnRangeChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("range_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDomainChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("domain_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsCurve() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCurve() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Curve", func(ptr gd.Object) any { return [1]gdclass.Curve{*(*gdclass.Curve)(unsafe.Pointer(&ptr))} })
}

type TangentMode = gdclass.CurveTangentMode //gd:Curve.TangentMode

const (
	/*The tangent on this side of the point is user-defined.*/
	TangentFree TangentMode = 0
	/*The curve calculates the tangent on this side of the point as the slope halfway towards the adjacent point.*/
	TangentLinear TangentMode = 1
	/*The total number of available tangent modes.*/
	TangentModeCount TangentMode = 2
)
