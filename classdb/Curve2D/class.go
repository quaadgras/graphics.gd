// Package Curve2D provides methods for working with Curve2D object instances.
package Curve2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform2D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
This class describes a Bézier curve in 2D space. It is mainly used to give a shape to a [Path2D], but can be manually sampled for other purposes.
It keeps a cache of precalculated points along the curve, to speed up further calculations.
*/
type Instance [1]gdclass.Curve2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCurve2D() Instance
}

/*
Adds a point with the specified [param position] relative to the curve's own position, with control points [param in] and [param out]. Appends the new point at the end of the point list.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. Every existing point starting from [param index] is shifted further down the list of points. The index must be greater than or equal to [code]0[/code] and must not exceed the number of existing points in the line. See [member point_count].
*/
func (self Instance) AddPoint(position Vector2.XY) { //gd:Curve2D.add_point
	class(self).AddPoint(gd.Vector2(position), gd.Vector2(gd.Vector2{0, 0}), gd.Vector2(gd.Vector2{0, 0}), gd.Int(-1))
}

/*
Sets the position for the vertex [param idx]. If the index is out of bounds, the function sends an error to the console.
*/
func (self Instance) SetPointPosition(idx int, position Vector2.XY) { //gd:Curve2D.set_point_position
	class(self).SetPointPosition(gd.Int(idx), gd.Vector2(position))
}

/*
Returns the position of the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0)[/code].
*/
func (self Instance) GetPointPosition(idx int) Vector2.XY { //gd:Curve2D.get_point_position
	return Vector2.XY(class(self).GetPointPosition(gd.Int(idx)))
}

/*
Sets the position of the control point leading to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console. The position is relative to the vertex.
*/
func (self Instance) SetPointIn(idx int, position Vector2.XY) { //gd:Curve2D.set_point_in
	class(self).SetPointIn(gd.Int(idx), gd.Vector2(position))
}

/*
Returns the position of the control point leading to the vertex [param idx]. The returned position is relative to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0)[/code].
*/
func (self Instance) GetPointIn(idx int) Vector2.XY { //gd:Curve2D.get_point_in
	return Vector2.XY(class(self).GetPointIn(gd.Int(idx)))
}

/*
Sets the position of the control point leading out of the vertex [param idx]. If the index is out of bounds, the function sends an error to the console. The position is relative to the vertex.
*/
func (self Instance) SetPointOut(idx int, position Vector2.XY) { //gd:Curve2D.set_point_out
	class(self).SetPointOut(gd.Int(idx), gd.Vector2(position))
}

/*
Returns the position of the control point leading out of the vertex [param idx]. The returned position is relative to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0)[/code].
*/
func (self Instance) GetPointOut(idx int) Vector2.XY { //gd:Curve2D.get_point_out
	return Vector2.XY(class(self).GetPointOut(gd.Int(idx)))
}

/*
Deletes the point [param idx] from the curve. Sends an error to the console if [param idx] is out of bounds.
*/
func (self Instance) RemovePoint(idx int) { //gd:Curve2D.remove_point
	class(self).RemovePoint(gd.Int(idx))
}

/*
Removes all points from the curve.
*/
func (self Instance) ClearPoints() { //gd:Curve2D.clear_points
	class(self).ClearPoints()
}

/*
Returns the position between the vertex [param idx] and the vertex [code]idx + 1[/code], where [param t] controls if the point is the first vertex ([code]t = 0.0[/code]), the last vertex ([code]t = 1.0[/code]), or in between. Values of [param t] outside the range ([code]0.0 >= t <=1[/code]) give strange, but predictable results.
If [param idx] is out of bounds it is truncated to the first or last vertex, and [param t] is ignored. If the curve has no points, the function sends an error to the console, and returns [code](0, 0)[/code].
*/
func (self Instance) Sample(idx int, t Float.X) Vector2.XY { //gd:Curve2D.sample
	return Vector2.XY(class(self).Sample(gd.Int(idx), gd.Float(t)))
}

/*
Returns the position at the vertex [param fofs]. It calls [method sample] using the integer part of [param fofs] as [code]idx[/code], and its fractional part as [code]t[/code].
*/
func (self Instance) Samplef(fofs Float.X) Vector2.XY { //gd:Curve2D.samplef
	return Vector2.XY(class(self).Samplef(gd.Float(fofs)))
}

/*
Returns the total length of the curve, based on the cached points. Given enough density (see [member bake_interval]), it should be approximate enough.
*/
func (self Instance) GetBakedLength() Float.X { //gd:Curve2D.get_baked_length
	return Float.X(Float.X(class(self).GetBakedLength()))
}

/*
Returns a point within the curve at position [param offset], where [param offset] is measured as a pixel distance along the curve.
To do that, it finds the two cached points where the [param offset] lies between, then interpolates the values. This interpolation is cubic if [param cubic] is set to [code]true[/code], or linear if set to [code]false[/code].
Cubic interpolation tends to follow the curves better, but linear is faster (and often, precise enough).
*/
func (self Instance) SampleBaked() Vector2.XY { //gd:Curve2D.sample_baked
	return Vector2.XY(class(self).SampleBaked(gd.Float(0.0), false))
}

/*
Similar to [method sample_baked], but returns [Transform2D] that includes a rotation along the curve, with [member Transform2D.origin] as the point position and the [member Transform2D.x] vector pointing in the direction of the path at that point. Returns an empty transform if the length of the curve is [code]0[/code].
[codeblock]
var baked = curve.sample_baked_with_rotation(offset)
# The returned Transform2D can be set directly.
transform = baked
# You can also read the origin and rotation separately from the returned Transform2D.
position = baked.get_origin()
rotation = baked.get_rotation()
[/codeblock]
*/
func (self Instance) SampleBakedWithRotation() Transform2D.OriginXY { //gd:Curve2D.sample_baked_with_rotation
	return Transform2D.OriginXY(class(self).SampleBakedWithRotation(gd.Float(0.0), false))
}

/*
Returns the cache of points as a [PackedVector2Array].
*/
func (self Instance) GetBakedPoints() []Vector2.XY { //gd:Curve2D.get_baked_points
	return []Vector2.XY(class(self).GetBakedPoints().AsSlice())
}

/*
Returns the closest point on baked segments (in curve's local space) to [param to_point].
[param to_point] must be in this curve's local space.
*/
func (self Instance) GetClosestPoint(to_point Vector2.XY) Vector2.XY { //gd:Curve2D.get_closest_point
	return Vector2.XY(class(self).GetClosestPoint(gd.Vector2(to_point)))
}

/*
Returns the closest offset to [param to_point]. This offset is meant to be used in [method sample_baked].
[param to_point] must be in this curve's local space.
*/
func (self Instance) GetClosestOffset(to_point Vector2.XY) Float.X { //gd:Curve2D.get_closest_offset
	return Float.X(Float.X(class(self).GetClosestOffset(gd.Vector2(to_point))))
}

/*
Returns a list of points along the curve, with a curvature controlled point density. That is, the curvier parts will have more points than the straighter parts.
This approximation makes straight segments between each point, then subdivides those segments until the resulting shape is similar enough.
[param max_stages] controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care!
[param tolerance_degrees] controls how many degrees the midpoint of a segment may deviate from the real curve, before the segment has to be subdivided.
*/
func (self Instance) Tessellate() []Vector2.XY { //gd:Curve2D.tessellate
	return []Vector2.XY(class(self).Tessellate(gd.Int(5), gd.Float(4)).AsSlice())
}

/*
Returns a list of points along the curve, with almost uniform density. [param max_stages] controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care!
[param tolerance_length] controls the maximal distance between two neighboring points, before the segment has to be subdivided.
*/
func (self Instance) TessellateEvenLength() []Vector2.XY { //gd:Curve2D.tessellate_even_length
	return []Vector2.XY(class(self).TessellateEvenLength(gd.Int(5), gd.Float(20.0)).AsSlice())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Curve2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Curve2D"))
	casted := Instance{*(*gdclass.Curve2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BakeInterval() Float.X {
	return Float.X(Float.X(class(self).GetBakeInterval()))
}

func (self Instance) SetBakeInterval(value Float.X) {
	class(self).SetBakeInterval(gd.Float(value))
}

func (self Instance) PointCount() int {
	return int(int(class(self).GetPointCount()))
}

func (self Instance) SetPointCount(value int) {
	class(self).SetPointCount(gd.Int(value))
}

//go:nosplit
func (self class) GetPointCount() gd.Int { //gd:Curve2D.get_point_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPointCount(count gd.Int) { //gd:Curve2D.set_point_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_set_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a point with the specified [param position] relative to the curve's own position, with control points [param in] and [param out]. Appends the new point at the end of the point list.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. Every existing point starting from [param index] is shifted further down the list of points. The index must be greater than or equal to [code]0[/code] and must not exceed the number of existing points in the line. See [member point_count].
*/
//go:nosplit
func (self class) AddPoint(position gd.Vector2, in gd.Vector2, out gd.Vector2, index gd.Int) { //gd:Curve2D.add_point
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, in)
	callframe.Arg(frame, out)
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the position for the vertex [param idx]. If the index is out of bounds, the function sends an error to the console.
*/
//go:nosplit
func (self class) SetPointPosition(idx gd.Int, position gd.Vector2) { //gd:Curve2D.set_point_position
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_set_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0)[/code].
*/
//go:nosplit
func (self class) GetPointPosition(idx gd.Int) gd.Vector2 { //gd:Curve2D.get_point_position
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the control point leading to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console. The position is relative to the vertex.
*/
//go:nosplit
func (self class) SetPointIn(idx gd.Int, position gd.Vector2) { //gd:Curve2D.set_point_in
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_set_point_in, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the control point leading to the vertex [param idx]. The returned position is relative to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0)[/code].
*/
//go:nosplit
func (self class) GetPointIn(idx gd.Int) gd.Vector2 { //gd:Curve2D.get_point_in
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_point_in, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the control point leading out of the vertex [param idx]. If the index is out of bounds, the function sends an error to the console. The position is relative to the vertex.
*/
//go:nosplit
func (self class) SetPointOut(idx gd.Int, position gd.Vector2) { //gd:Curve2D.set_point_out
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_set_point_out, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the control point leading out of the vertex [param idx]. The returned position is relative to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0)[/code].
*/
//go:nosplit
func (self class) GetPointOut(idx gd.Int) gd.Vector2 { //gd:Curve2D.get_point_out
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_point_out, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Deletes the point [param idx] from the curve. Sends an error to the console if [param idx] is out of bounds.
*/
//go:nosplit
func (self class) RemovePoint(idx gd.Int) { //gd:Curve2D.remove_point
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all points from the curve.
*/
//go:nosplit
func (self class) ClearPoints() { //gd:Curve2D.clear_points
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_clear_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position between the vertex [param idx] and the vertex [code]idx + 1[/code], where [param t] controls if the point is the first vertex ([code]t = 0.0[/code]), the last vertex ([code]t = 1.0[/code]), or in between. Values of [param t] outside the range ([code]0.0 >= t <=1[/code]) give strange, but predictable results.
If [param idx] is out of bounds it is truncated to the first or last vertex, and [param t] is ignored. If the curve has no points, the function sends an error to the console, and returns [code](0, 0)[/code].
*/
//go:nosplit
func (self class) Sample(idx gd.Int, t gd.Float) gd.Vector2 { //gd:Curve2D.sample
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, t)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_sample, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position at the vertex [param fofs]. It calls [method sample] using the integer part of [param fofs] as [code]idx[/code], and its fractional part as [code]t[/code].
*/
//go:nosplit
func (self class) Samplef(fofs gd.Float) gd.Vector2 { //gd:Curve2D.samplef
	var frame = callframe.New()
	callframe.Arg(frame, fofs)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_samplef, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeInterval(distance gd.Float) { //gd:Curve2D.set_bake_interval
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_set_bake_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeInterval() gd.Float { //gd:Curve2D.get_bake_interval
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_bake_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total length of the curve, based on the cached points. Given enough density (see [member bake_interval]), it should be approximate enough.
*/
//go:nosplit
func (self class) GetBakedLength() gd.Float { //gd:Curve2D.get_baked_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_baked_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a point within the curve at position [param offset], where [param offset] is measured as a pixel distance along the curve.
To do that, it finds the two cached points where the [param offset] lies between, then interpolates the values. This interpolation is cubic if [param cubic] is set to [code]true[/code], or linear if set to [code]false[/code].
Cubic interpolation tends to follow the curves better, but linear is faster (and often, precise enough).
*/
//go:nosplit
func (self class) SampleBaked(offset gd.Float, cubic bool) gd.Vector2 { //gd:Curve2D.sample_baked
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, cubic)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_sample_baked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Similar to [method sample_baked], but returns [Transform2D] that includes a rotation along the curve, with [member Transform2D.origin] as the point position and the [member Transform2D.x] vector pointing in the direction of the path at that point. Returns an empty transform if the length of the curve is [code]0[/code].
[codeblock]
var baked = curve.sample_baked_with_rotation(offset)
# The returned Transform2D can be set directly.
transform = baked
# You can also read the origin and rotation separately from the returned Transform2D.
position = baked.get_origin()
rotation = baked.get_rotation()
[/codeblock]
*/
//go:nosplit
func (self class) SampleBakedWithRotation(offset gd.Float, cubic bool) gd.Transform2D { //gd:Curve2D.sample_baked_with_rotation
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, cubic)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_sample_baked_with_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the cache of points as a [PackedVector2Array].
*/
//go:nosplit
func (self class) GetBakedPoints() gd.PackedVector2Array { //gd:Curve2D.get_baked_points
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_baked_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the closest point on baked segments (in curve's local space) to [param to_point].
[param to_point] must be in this curve's local space.
*/
//go:nosplit
func (self class) GetClosestPoint(to_point gd.Vector2) gd.Vector2 { //gd:Curve2D.get_closest_point
	var frame = callframe.New()
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the closest offset to [param to_point]. This offset is meant to be used in [method sample_baked].
[param to_point] must be in this curve's local space.
*/
//go:nosplit
func (self class) GetClosestOffset(to_point gd.Vector2) gd.Float { //gd:Curve2D.get_closest_offset
	var frame = callframe.New()
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_get_closest_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of points along the curve, with a curvature controlled point density. That is, the curvier parts will have more points than the straighter parts.
This approximation makes straight segments between each point, then subdivides those segments until the resulting shape is similar enough.
[param max_stages] controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care!
[param tolerance_degrees] controls how many degrees the midpoint of a segment may deviate from the real curve, before the segment has to be subdivided.
*/
//go:nosplit
func (self class) Tessellate(max_stages gd.Int, tolerance_degrees gd.Float) gd.PackedVector2Array { //gd:Curve2D.tessellate
	var frame = callframe.New()
	callframe.Arg(frame, max_stages)
	callframe.Arg(frame, tolerance_degrees)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_tessellate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of points along the curve, with almost uniform density. [param max_stages] controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care!
[param tolerance_length] controls the maximal distance between two neighboring points, before the segment has to be subdivided.
*/
//go:nosplit
func (self class) TessellateEvenLength(max_stages gd.Int, tolerance_length gd.Float) gd.PackedVector2Array { //gd:Curve2D.tessellate_even_length
	var frame = callframe.New()
	callframe.Arg(frame, max_stages)
	callframe.Arg(frame, tolerance_length)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve2D.Bind_tessellate_even_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsCurve2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCurve2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Curve2D", func(ptr gd.Object) any { return [1]gdclass.Curve2D{*(*gdclass.Curve2D)(unsafe.Pointer(&ptr))} })
}
