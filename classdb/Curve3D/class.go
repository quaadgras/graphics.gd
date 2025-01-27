// Package Curve3D provides methods for working with Curve3D object instances.
package Curve3D

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform3D"

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

/*
This class describes a Bézier curve in 3D space. It is mainly used to give a shape to a [Path3D], but can be manually sampled for other purposes.
It keeps a cache of precalculated points along the curve, to speed up further calculations.
*/
type Instance [1]gdclass.Curve3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCurve3D() Instance
}

/*
Adds a point with the specified [param position] relative to the curve's own position, with control points [param in] and [param out]. Appends the new point at the end of the point list.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. Every existing point starting from [param index] is shifted further down the list of points. The index must be greater than or equal to [code]0[/code] and must not exceed the number of existing points in the line. See [member point_count].
*/
func (self Instance) AddPoint(position Vector3.XYZ) { //gd:Curve3D.add_point
	class(self).AddPoint(gd.Vector3(position), gd.Vector3(gd.Vector3{0, 0, 0}), gd.Vector3(gd.Vector3{0, 0, 0}), gd.Int(-1))
}

/*
Sets the position for the vertex [param idx]. If the index is out of bounds, the function sends an error to the console.
*/
func (self Instance) SetPointPosition(idx int, position Vector3.XYZ) { //gd:Curve3D.set_point_position
	class(self).SetPointPosition(gd.Int(idx), gd.Vector3(position))
}

/*
Returns the position of the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
*/
func (self Instance) GetPointPosition(idx int) Vector3.XYZ { //gd:Curve3D.get_point_position
	return Vector3.XYZ(class(self).GetPointPosition(gd.Int(idx)))
}

/*
Sets the tilt angle in radians for the point [param idx]. If the index is out of bounds, the function sends an error to the console.
The tilt controls the rotation along the look-at axis an object traveling the path would have. In the case of a curve controlling a [PathFollow3D], this tilt is an offset over the natural tilt the [PathFollow3D] calculates.
*/
func (self Instance) SetPointTilt(idx int, tilt Float.X) { //gd:Curve3D.set_point_tilt
	class(self).SetPointTilt(gd.Int(idx), gd.Float(tilt))
}

/*
Returns the tilt angle in radians for the point [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code]0[/code].
*/
func (self Instance) GetPointTilt(idx int) Float.X { //gd:Curve3D.get_point_tilt
	return Float.X(Float.X(class(self).GetPointTilt(gd.Int(idx))))
}

/*
Sets the position of the control point leading to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console. The position is relative to the vertex.
*/
func (self Instance) SetPointIn(idx int, position Vector3.XYZ) { //gd:Curve3D.set_point_in
	class(self).SetPointIn(gd.Int(idx), gd.Vector3(position))
}

/*
Returns the position of the control point leading to the vertex [param idx]. The returned position is relative to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
*/
func (self Instance) GetPointIn(idx int) Vector3.XYZ { //gd:Curve3D.get_point_in
	return Vector3.XYZ(class(self).GetPointIn(gd.Int(idx)))
}

/*
Sets the position of the control point leading out of the vertex [param idx]. If the index is out of bounds, the function sends an error to the console. The position is relative to the vertex.
*/
func (self Instance) SetPointOut(idx int, position Vector3.XYZ) { //gd:Curve3D.set_point_out
	class(self).SetPointOut(gd.Int(idx), gd.Vector3(position))
}

/*
Returns the position of the control point leading out of the vertex [param idx]. The returned position is relative to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
*/
func (self Instance) GetPointOut(idx int) Vector3.XYZ { //gd:Curve3D.get_point_out
	return Vector3.XYZ(class(self).GetPointOut(gd.Int(idx)))
}

/*
Deletes the point [param idx] from the curve. Sends an error to the console if [param idx] is out of bounds.
*/
func (self Instance) RemovePoint(idx int) { //gd:Curve3D.remove_point
	class(self).RemovePoint(gd.Int(idx))
}

/*
Removes all points from the curve.
*/
func (self Instance) ClearPoints() { //gd:Curve3D.clear_points
	class(self).ClearPoints()
}

/*
Returns the position between the vertex [param idx] and the vertex [code]idx + 1[/code], where [param t] controls if the point is the first vertex ([code]t = 0.0[/code]), the last vertex ([code]t = 1.0[/code]), or in between. Values of [param t] outside the range ([code]0.0 >= t <=1[/code]) give strange, but predictable results.
If [param idx] is out of bounds it is truncated to the first or last vertex, and [param t] is ignored. If the curve has no points, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
*/
func (self Instance) Sample(idx int, t Float.X) Vector3.XYZ { //gd:Curve3D.sample
	return Vector3.XYZ(class(self).Sample(gd.Int(idx), gd.Float(t)))
}

/*
Returns the position at the vertex [param fofs]. It calls [method sample] using the integer part of [param fofs] as [code]idx[/code], and its fractional part as [code]t[/code].
*/
func (self Instance) Samplef(fofs Float.X) Vector3.XYZ { //gd:Curve3D.samplef
	return Vector3.XYZ(class(self).Samplef(gd.Float(fofs)))
}

/*
Returns the total length of the curve, based on the cached points. Given enough density (see [member bake_interval]), it should be approximate enough.
*/
func (self Instance) GetBakedLength() Float.X { //gd:Curve3D.get_baked_length
	return Float.X(Float.X(class(self).GetBakedLength()))
}

/*
Returns a point within the curve at position [param offset], where [param offset] is measured as a distance in 3D units along the curve. To do that, it finds the two cached points where the [param offset] lies between, then interpolates the values. This interpolation is cubic if [param cubic] is set to [code]true[/code], or linear if set to [code]false[/code].
Cubic interpolation tends to follow the curves better, but linear is faster (and often, precise enough).
*/
func (self Instance) SampleBaked() Vector3.XYZ { //gd:Curve3D.sample_baked
	return Vector3.XYZ(class(self).SampleBaked(gd.Float(0.0), false))
}

/*
Returns a [Transform3D] with [code]origin[/code] as point position, [code]basis.x[/code] as sideway vector, [code]basis.y[/code] as up vector, [code]basis.z[/code] as forward vector. When the curve length is 0, there is no reasonable way to calculate the rotation, all vectors aligned with global space axes. See also [method sample_baked].
*/
func (self Instance) SampleBakedWithRotation() Transform3D.BasisOrigin { //gd:Curve3D.sample_baked_with_rotation
	return Transform3D.BasisOrigin(class(self).SampleBakedWithRotation(gd.Float(0.0), false, false))
}

/*
Returns an up vector within the curve at position [param offset], where [param offset] is measured as a distance in 3D units along the curve. To do that, it finds the two cached up vectors where the [param offset] lies between, then interpolates the values. If [param apply_tilt] is [code]true[/code], an interpolated tilt is applied to the interpolated up vector.
If the curve has no up vectors, the function sends an error to the console, and returns [code](0, 1, 0)[/code].
*/
func (self Instance) SampleBakedUpVector(offset Float.X) Vector3.XYZ { //gd:Curve3D.sample_baked_up_vector
	return Vector3.XYZ(class(self).SampleBakedUpVector(gd.Float(offset), false))
}

/*
Returns the cache of points as a [PackedVector3Array].
*/
func (self Instance) GetBakedPoints() []Vector3.XYZ { //gd:Curve3D.get_baked_points
	return []Vector3.XYZ(class(self).GetBakedPoints().AsSlice())
}

/*
Returns the cache of tilts as a [PackedFloat32Array].
*/
func (self Instance) GetBakedTilts() []float32 { //gd:Curve3D.get_baked_tilts
	return []float32(class(self).GetBakedTilts().AsSlice())
}

/*
Returns the cache of up vectors as a [PackedVector3Array].
If [member up_vector_enabled] is [code]false[/code], the cache will be empty.
*/
func (self Instance) GetBakedUpVectors() []Vector3.XYZ { //gd:Curve3D.get_baked_up_vectors
	return []Vector3.XYZ(class(self).GetBakedUpVectors().AsSlice())
}

/*
Returns the closest point on baked segments (in curve's local space) to [param to_point].
[param to_point] must be in this curve's local space.
*/
func (self Instance) GetClosestPoint(to_point Vector3.XYZ) Vector3.XYZ { //gd:Curve3D.get_closest_point
	return Vector3.XYZ(class(self).GetClosestPoint(gd.Vector3(to_point)))
}

/*
Returns the closest offset to [param to_point]. This offset is meant to be used in [method sample_baked] or [method sample_baked_up_vector].
[param to_point] must be in this curve's local space.
*/
func (self Instance) GetClosestOffset(to_point Vector3.XYZ) Float.X { //gd:Curve3D.get_closest_offset
	return Float.X(Float.X(class(self).GetClosestOffset(gd.Vector3(to_point))))
}

/*
Returns a list of points along the curve, with a curvature controlled point density. That is, the curvier parts will have more points than the straighter parts.
This approximation makes straight segments between each point, then subdivides those segments until the resulting shape is similar enough.
[param max_stages] controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care!
[param tolerance_degrees] controls how many degrees the midpoint of a segment may deviate from the real curve, before the segment has to be subdivided.
*/
func (self Instance) Tessellate() []Vector3.XYZ { //gd:Curve3D.tessellate
	return []Vector3.XYZ(class(self).Tessellate(gd.Int(5), gd.Float(4)).AsSlice())
}

/*
Returns a list of points along the curve, with almost uniform density. [param max_stages] controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care!
[param tolerance_length] controls the maximal distance between two neighboring points, before the segment has to be subdivided.
*/
func (self Instance) TessellateEvenLength() []Vector3.XYZ { //gd:Curve3D.tessellate_even_length
	return []Vector3.XYZ(class(self).TessellateEvenLength(gd.Int(5), gd.Float(0.2)).AsSlice())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Curve3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Curve3D"))
	casted := Instance{*(*gdclass.Curve3D)(unsafe.Pointer(&object))}
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

func (self Instance) UpVectorEnabled() bool {
	return bool(class(self).IsUpVectorEnabled())
}

func (self Instance) SetUpVectorEnabled(value bool) {
	class(self).SetUpVectorEnabled(value)
}

//go:nosplit
func (self class) GetPointCount() gd.Int { //gd:Curve3D.get_point_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPointCount(count gd.Int) { //gd:Curve3D.set_point_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_set_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a point with the specified [param position] relative to the curve's own position, with control points [param in] and [param out]. Appends the new point at the end of the point list.
If [param index] is given, the new point is inserted before the existing point identified by index [param index]. Every existing point starting from [param index] is shifted further down the list of points. The index must be greater than or equal to [code]0[/code] and must not exceed the number of existing points in the line. See [member point_count].
*/
//go:nosplit
func (self class) AddPoint(position gd.Vector3, in gd.Vector3, out gd.Vector3, index gd.Int) { //gd:Curve3D.add_point
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, in)
	callframe.Arg(frame, out)
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the position for the vertex [param idx]. If the index is out of bounds, the function sends an error to the console.
*/
//go:nosplit
func (self class) SetPointPosition(idx gd.Int, position gd.Vector3) { //gd:Curve3D.set_point_position
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_set_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
*/
//go:nosplit
func (self class) GetPointPosition(idx gd.Int) gd.Vector3 { //gd:Curve3D.get_point_position
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tilt angle in radians for the point [param idx]. If the index is out of bounds, the function sends an error to the console.
The tilt controls the rotation along the look-at axis an object traveling the path would have. In the case of a curve controlling a [PathFollow3D], this tilt is an offset over the natural tilt the [PathFollow3D] calculates.
*/
//go:nosplit
func (self class) SetPointTilt(idx gd.Int, tilt gd.Float) { //gd:Curve3D.set_point_tilt
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, tilt)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_set_point_tilt, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tilt angle in radians for the point [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code]0[/code].
*/
//go:nosplit
func (self class) GetPointTilt(idx gd.Int) gd.Float { //gd:Curve3D.get_point_tilt
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_point_tilt, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the control point leading to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console. The position is relative to the vertex.
*/
//go:nosplit
func (self class) SetPointIn(idx gd.Int, position gd.Vector3) { //gd:Curve3D.set_point_in
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_set_point_in, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the control point leading to the vertex [param idx]. The returned position is relative to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
*/
//go:nosplit
func (self class) GetPointIn(idx gd.Int) gd.Vector3 { //gd:Curve3D.get_point_in
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_point_in, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the control point leading out of the vertex [param idx]. If the index is out of bounds, the function sends an error to the console. The position is relative to the vertex.
*/
//go:nosplit
func (self class) SetPointOut(idx gd.Int, position gd.Vector3) { //gd:Curve3D.set_point_out
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_set_point_out, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the control point leading out of the vertex [param idx]. The returned position is relative to the vertex [param idx]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
*/
//go:nosplit
func (self class) GetPointOut(idx gd.Int) gd.Vector3 { //gd:Curve3D.get_point_out
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_point_out, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Deletes the point [param idx] from the curve. Sends an error to the console if [param idx] is out of bounds.
*/
//go:nosplit
func (self class) RemovePoint(idx gd.Int) { //gd:Curve3D.remove_point
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all points from the curve.
*/
//go:nosplit
func (self class) ClearPoints() { //gd:Curve3D.clear_points
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_clear_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position between the vertex [param idx] and the vertex [code]idx + 1[/code], where [param t] controls if the point is the first vertex ([code]t = 0.0[/code]), the last vertex ([code]t = 1.0[/code]), or in between. Values of [param t] outside the range ([code]0.0 >= t <=1[/code]) give strange, but predictable results.
If [param idx] is out of bounds it is truncated to the first or last vertex, and [param t] is ignored. If the curve has no points, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
*/
//go:nosplit
func (self class) Sample(idx gd.Int, t gd.Float) gd.Vector3 { //gd:Curve3D.sample
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, t)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_sample, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position at the vertex [param fofs]. It calls [method sample] using the integer part of [param fofs] as [code]idx[/code], and its fractional part as [code]t[/code].
*/
//go:nosplit
func (self class) Samplef(fofs gd.Float) gd.Vector3 { //gd:Curve3D.samplef
	var frame = callframe.New()
	callframe.Arg(frame, fofs)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_samplef, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeInterval(distance gd.Float) { //gd:Curve3D.set_bake_interval
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_set_bake_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeInterval() gd.Float { //gd:Curve3D.get_bake_interval
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_bake_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpVectorEnabled(enable bool) { //gd:Curve3D.set_up_vector_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_set_up_vector_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUpVectorEnabled() bool { //gd:Curve3D.is_up_vector_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_is_up_vector_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total length of the curve, based on the cached points. Given enough density (see [member bake_interval]), it should be approximate enough.
*/
//go:nosplit
func (self class) GetBakedLength() gd.Float { //gd:Curve3D.get_baked_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_baked_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a point within the curve at position [param offset], where [param offset] is measured as a distance in 3D units along the curve. To do that, it finds the two cached points where the [param offset] lies between, then interpolates the values. This interpolation is cubic if [param cubic] is set to [code]true[/code], or linear if set to [code]false[/code].
Cubic interpolation tends to follow the curves better, but linear is faster (and often, precise enough).
*/
//go:nosplit
func (self class) SampleBaked(offset gd.Float, cubic bool) gd.Vector3 { //gd:Curve3D.sample_baked
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, cubic)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_sample_baked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Transform3D] with [code]origin[/code] as point position, [code]basis.x[/code] as sideway vector, [code]basis.y[/code] as up vector, [code]basis.z[/code] as forward vector. When the curve length is 0, there is no reasonable way to calculate the rotation, all vectors aligned with global space axes. See also [method sample_baked].
*/
//go:nosplit
func (self class) SampleBakedWithRotation(offset gd.Float, cubic bool, apply_tilt bool) gd.Transform3D { //gd:Curve3D.sample_baked_with_rotation
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, cubic)
	callframe.Arg(frame, apply_tilt)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_sample_baked_with_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an up vector within the curve at position [param offset], where [param offset] is measured as a distance in 3D units along the curve. To do that, it finds the two cached up vectors where the [param offset] lies between, then interpolates the values. If [param apply_tilt] is [code]true[/code], an interpolated tilt is applied to the interpolated up vector.
If the curve has no up vectors, the function sends an error to the console, and returns [code](0, 1, 0)[/code].
*/
//go:nosplit
func (self class) SampleBakedUpVector(offset gd.Float, apply_tilt bool) gd.Vector3 { //gd:Curve3D.sample_baked_up_vector
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, apply_tilt)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_sample_baked_up_vector, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the cache of points as a [PackedVector3Array].
*/
//go:nosplit
func (self class) GetBakedPoints() gd.PackedVector3Array { //gd:Curve3D.get_baked_points
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_baked_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the cache of tilts as a [PackedFloat32Array].
*/
//go:nosplit
func (self class) GetBakedTilts() gd.PackedFloat32Array { //gd:Curve3D.get_baked_tilts
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_baked_tilts, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the cache of up vectors as a [PackedVector3Array].
If [member up_vector_enabled] is [code]false[/code], the cache will be empty.
*/
//go:nosplit
func (self class) GetBakedUpVectors() gd.PackedVector3Array { //gd:Curve3D.get_baked_up_vectors
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_baked_up_vectors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the closest point on baked segments (in curve's local space) to [param to_point].
[param to_point] must be in this curve's local space.
*/
//go:nosplit
func (self class) GetClosestPoint(to_point gd.Vector3) gd.Vector3 { //gd:Curve3D.get_closest_point
	var frame = callframe.New()
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the closest offset to [param to_point]. This offset is meant to be used in [method sample_baked] or [method sample_baked_up_vector].
[param to_point] must be in this curve's local space.
*/
//go:nosplit
func (self class) GetClosestOffset(to_point gd.Vector3) gd.Float { //gd:Curve3D.get_closest_offset
	var frame = callframe.New()
	callframe.Arg(frame, to_point)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_get_closest_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) Tessellate(max_stages gd.Int, tolerance_degrees gd.Float) gd.PackedVector3Array { //gd:Curve3D.tessellate
	var frame = callframe.New()
	callframe.Arg(frame, max_stages)
	callframe.Arg(frame, tolerance_degrees)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_tessellate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of points along the curve, with almost uniform density. [param max_stages] controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care!
[param tolerance_length] controls the maximal distance between two neighboring points, before the segment has to be subdivided.
*/
//go:nosplit
func (self class) TessellateEvenLength(max_stages gd.Int, tolerance_length gd.Float) gd.PackedVector3Array { //gd:Curve3D.tessellate_even_length
	var frame = callframe.New()
	callframe.Arg(frame, max_stages)
	callframe.Arg(frame, tolerance_length)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve3D.Bind_tessellate_even_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsCurve3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCurve3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Curve3D", func(ptr gd.Object) any { return [1]gdclass.Curve3D{*(*gdclass.Curve3D)(unsafe.Pointer(&ptr))} })
}
