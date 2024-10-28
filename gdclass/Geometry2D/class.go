package Geometry2D

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Provides a set of helper functions to create geometric shapes, compute intersections between shapes, and process various other geometric operations in 2D.

*/
var self gdclass.Geometry2D
var once sync.Once
func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Geometry2D)
	self = *(*gdclass.Geometry2D)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if [param point] is inside the circle or if it's located exactly [i]on[/i] the circle's boundary, otherwise returns [code]false[/code].
*/
func IsPointInCircle(point gd.Vector2, circle_position gd.Vector2, circle_radius float64) bool {
	once.Do(singleton)
	return bool(class(self).IsPointInCircle(point, circle_position, gd.Float(circle_radius)))
}

/*
Given the 2D segment ([param segment_from], [param segment_to]), returns the position on the segment (as a number between 0 and 1) at which the segment hits the circle that is located at position [param circle_position] and has radius [param circle_radius]. If the segment does not intersect the circle, -1 is returned (this is also the case if the line extending the segment would intersect the circle, but the segment does not).
*/
func SegmentIntersectsCircle(segment_from gd.Vector2, segment_to gd.Vector2, circle_position gd.Vector2, circle_radius float64) float64 {
	once.Do(singleton)
	return float64(float64(class(self).SegmentIntersectsCircle(segment_from, segment_to, circle_position, gd.Float(circle_radius))))
}

/*
Checks if the two segments ([param from_a], [param to_a]) and ([param from_b], [param to_b]) intersect. If yes, return the point of intersection as [Vector2]. If no intersection takes place, returns [code]null[/code].
*/
func SegmentIntersectsSegment(from_a gd.Vector2, to_a gd.Vector2, from_b gd.Vector2, to_b gd.Vector2) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).SegmentIntersectsSegment(from_a, to_a, from_b, to_b))
}

/*
Checks if the two lines ([param from_a], [param dir_a]) and ([param from_b], [param dir_b]) intersect. If yes, return the point of intersection as [Vector2]. If no intersection takes place, returns [code]null[/code].
[b]Note:[/b] The lines are specified using direction vectors, not end points.
*/
func LineIntersectsLine(from_a gd.Vector2, dir_a gd.Vector2, from_b gd.Vector2, dir_b gd.Vector2) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).LineIntersectsLine(from_a, dir_a, from_b, dir_b))
}

/*
Given the two 2D segments ([param p1], [param q1]) and ([param p2], [param q2]), finds those two points on the two segments that are closest to each other. Returns a [PackedVector2Array] that contains this point on ([param p1], [param q1]) as well the accompanying point on ([param p2], [param q2]).
*/
func GetClosestPointsBetweenSegments(p1 gd.Vector2, q1 gd.Vector2, p2 gd.Vector2, q2 gd.Vector2) []gd.Vector2 {
	once.Do(singleton)
	return []gd.Vector2(class(self).GetClosestPointsBetweenSegments(p1, q1, p2, q2).AsSlice())
}

/*
Returns the 2D point on the 2D segment ([param s1], [param s2]) that is closest to [param point]. The returned point will always be inside the specified segment.
*/
func GetClosestPointToSegment(point gd.Vector2, s1 gd.Vector2, s2 gd.Vector2) gd.Vector2 {
	once.Do(singleton)
	return gd.Vector2(class(self).GetClosestPointToSegment(point, s1, s2))
}

/*
Returns the 2D point on the 2D line defined by ([param s1], [param s2]) that is closest to [param point]. The returned point can be inside the segment ([param s1], [param s2]) or outside of it, i.e. somewhere on the line extending from the segment.
*/
func GetClosestPointToSegmentUncapped(point gd.Vector2, s1 gd.Vector2, s2 gd.Vector2) gd.Vector2 {
	once.Do(singleton)
	return gd.Vector2(class(self).GetClosestPointToSegmentUncapped(point, s1, s2))
}

/*
Returns if [param point] is inside the triangle specified by [param a], [param b] and [param c].
*/
func PointIsInsideTriangle(point gd.Vector2, a gd.Vector2, b gd.Vector2, c gd.Vector2) bool {
	once.Do(singleton)
	return bool(class(self).PointIsInsideTriangle(point, a, b, c))
}

/*
Returns [code]true[/code] if [param polygon]'s vertices are ordered in clockwise order, otherwise returns [code]false[/code].
[b]Note:[/b] Assumes a Cartesian coordinate system where [code]+x[/code] is right and [code]+y[/code] is up. If using screen coordinates ([code]+y[/code] is down), the result will need to be flipped (i.e. a [code]true[/code] result will indicate counter-clockwise).
*/
func IsPolygonClockwise(polygon []gd.Vector2) bool {
	once.Do(singleton)
	return bool(class(self).IsPolygonClockwise(gd.NewPackedVector2Slice(polygon)))
}

/*
Returns [code]true[/code] if [param point] is inside [param polygon] or if it's located exactly [i]on[/i] polygon's boundary, otherwise returns [code]false[/code].
*/
func IsPointInPolygon(point gd.Vector2, polygon []gd.Vector2) bool {
	once.Do(singleton)
	return bool(class(self).IsPointInPolygon(point, gd.NewPackedVector2Slice(polygon)))
}

/*
Triangulates the polygon specified by the points in [param polygon]. Returns a [PackedInt32Array] where each triangle consists of three consecutive point indices into [param polygon] (i.e. the returned array will have [code]n * 3[/code] elements, with [code]n[/code] being the number of found triangles). Output triangles will always be counter clockwise, and the contour will be flipped if it's clockwise. If the triangulation did not succeed, an empty [PackedInt32Array] is returned.
*/
func TriangulatePolygon(polygon []gd.Vector2) []int32 {
	once.Do(singleton)
	return []int32(class(self).TriangulatePolygon(gd.NewPackedVector2Slice(polygon)).AsSlice())
}

/*
Triangulates the area specified by discrete set of [param points] such that no point is inside the circumcircle of any resulting triangle. Returns a [PackedInt32Array] where each triangle consists of three consecutive point indices into [param points] (i.e. the returned array will have [code]n * 3[/code] elements, with [code]n[/code] being the number of found triangles). If the triangulation did not succeed, an empty [PackedInt32Array] is returned.
*/
func TriangulateDelaunay(points []gd.Vector2) []int32 {
	once.Do(singleton)
	return []int32(class(self).TriangulateDelaunay(gd.NewPackedVector2Slice(points)).AsSlice())
}

/*
Given an array of [Vector2]s, returns the convex hull as a list of points in counterclockwise order. The last point is the same as the first one.
*/
func ConvexHull(points []gd.Vector2) []gd.Vector2 {
	once.Do(singleton)
	return []gd.Vector2(class(self).ConvexHull(gd.NewPackedVector2Slice(points)).AsSlice())
}

/*
Decomposes the [param polygon] into multiple convex hulls and returns an array of [PackedVector2Array].
*/
func DecomposePolygonInConvex(polygon []gd.Vector2) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).DecomposePolygonInConvex(gd.NewPackedVector2Slice(polygon)))
}

/*
Merges (combines) [param polygon_a] and [param polygon_b] and returns an array of merged polygons. This performs [constant OPERATION_UNION] between polygons.
The operation may result in an outer polygon (boundary) and multiple inner polygons (holes) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
func MergePolygons(polygon_a []gd.Vector2, polygon_b []gd.Vector2) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).MergePolygons(gd.NewPackedVector2Slice(polygon_a), gd.NewPackedVector2Slice(polygon_b)))
}

/*
Clips [param polygon_a] against [param polygon_b] and returns an array of clipped polygons. This performs [constant OPERATION_DIFFERENCE] between polygons. Returns an empty array if [param polygon_b] completely overlaps [param polygon_a].
If [param polygon_b] is enclosed by [param polygon_a], returns an outer polygon (boundary) and inner polygon (hole) which could be distinguished by calling [method is_polygon_clockwise].
*/
func ClipPolygons(polygon_a []gd.Vector2, polygon_b []gd.Vector2) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).ClipPolygons(gd.NewPackedVector2Slice(polygon_a), gd.NewPackedVector2Slice(polygon_b)))
}

/*
Intersects [param polygon_a] with [param polygon_b] and returns an array of intersected polygons. This performs [constant OPERATION_INTERSECTION] between polygons. In other words, returns common area shared by polygons. Returns an empty array if no intersection occurs.
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
func IntersectPolygons(polygon_a []gd.Vector2, polygon_b []gd.Vector2) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).IntersectPolygons(gd.NewPackedVector2Slice(polygon_a), gd.NewPackedVector2Slice(polygon_b)))
}

/*
Mutually excludes common area defined by intersection of [param polygon_a] and [param polygon_b] (see [method intersect_polygons]) and returns an array of excluded polygons. This performs [constant OPERATION_XOR] between polygons. In other words, returns all but common area between polygons.
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
func ExcludePolygons(polygon_a []gd.Vector2, polygon_b []gd.Vector2) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).ExcludePolygons(gd.NewPackedVector2Slice(polygon_a), gd.NewPackedVector2Slice(polygon_b)))
}

/*
Clips [param polyline] against [param polygon] and returns an array of clipped polylines. This performs [constant OPERATION_DIFFERENCE] between the polyline and the polygon. This operation can be thought of as cutting a line with a closed shape.
*/
func ClipPolylineWithPolygon(polyline []gd.Vector2, polygon []gd.Vector2) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).ClipPolylineWithPolygon(gd.NewPackedVector2Slice(polyline), gd.NewPackedVector2Slice(polygon)))
}

/*
Intersects [param polyline] with [param polygon] and returns an array of intersected polylines. This performs [constant OPERATION_INTERSECTION] between the polyline and the polygon. This operation can be thought of as chopping a line with a closed shape.
*/
func IntersectPolylineWithPolygon(polyline []gd.Vector2, polygon []gd.Vector2) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).IntersectPolylineWithPolygon(gd.NewPackedVector2Slice(polyline), gd.NewPackedVector2Slice(polygon)))
}

/*
Inflates or deflates [param polygon] by [param delta] units (pixels). If [param delta] is positive, makes the polygon grow outward. If [param delta] is negative, shrinks the polygon inward. Returns an array of polygons because inflating/deflating may result in multiple discrete polygons. Returns an empty array if [param delta] is negative and the absolute value of it approximately exceeds the minimum bounding rectangle dimensions of the polygon.
Each polygon's vertices will be rounded as determined by [param join_type], see [enum PolyJoinType].
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
[b]Note:[/b] To translate the polygon's vertices specifically, multiply them to a [Transform2D]:
[codeblocks]
[gdscript]
var polygon = PackedVector2Array([Vector2(0, 0), Vector2(100, 0), Vector2(100, 100), Vector2(0, 100)])
var offset = Vector2(50, 50)
polygon = Transform2D(0, offset) * polygon
print(polygon) # prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/gdscript]
[csharp]
var polygon = new Vector2[] { new Vector2(0, 0), new Vector2(100, 0), new Vector2(100, 100), new Vector2(0, 100) };
var offset = new Vector2(50, 50);
polygon = new Transform2D(0, offset) * polygon;
GD.Print((Variant)polygon); // prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/csharp]
[/codeblocks]
*/
func OffsetPolygon(polygon []gd.Vector2, delta float64) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).OffsetPolygon(gd.NewPackedVector2Slice(polygon), gd.Float(delta), 0))
}

/*
Inflates or deflates [param polyline] by [param delta] units (pixels), producing polygons. If [param delta] is positive, makes the polyline grow outward. Returns an array of polygons because inflating/deflating may result in multiple discrete polygons. If [param delta] is negative, returns an empty array.
Each polygon's vertices will be rounded as determined by [param join_type], see [enum PolyJoinType].
Each polygon's endpoints will be rounded as determined by [param end_type], see [enum PolyEndType].
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
func OffsetPolyline(polyline []gd.Vector2, delta float64) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).OffsetPolyline(gd.NewPackedVector2Slice(polyline), gd.Float(delta), 0, 3))
}

/*
Given an array of [Vector2]s representing tiles, builds an atlas. The returned dictionary has two keys: [code]points[/code] is a [PackedVector2Array] that specifies the positions of each tile, [code]size[/code] contains the overall size of the whole atlas as [Vector2i].
*/
func MakeAtlas(sizes []gd.Vector2) gd.Dictionary {
	once.Do(singleton)
	return gd.Dictionary(class(self).MakeAtlas(gd.NewPackedVector2Slice(sizes)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.Geometry2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Returns [code]true[/code] if [param point] is inside the circle or if it's located exactly [i]on[/i] the circle's boundary, otherwise returns [code]false[/code].
*/
//go:nosplit
func (self class) IsPointInCircle(point gd.Vector2, circle_position gd.Vector2, circle_radius gd.Float) bool {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, circle_position)
	callframe.Arg(frame, circle_radius)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_is_point_in_circle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Given the 2D segment ([param segment_from], [param segment_to]), returns the position on the segment (as a number between 0 and 1) at which the segment hits the circle that is located at position [param circle_position] and has radius [param circle_radius]. If the segment does not intersect the circle, -1 is returned (this is also the case if the line extending the segment would intersect the circle, but the segment does not).
*/
//go:nosplit
func (self class) SegmentIntersectsCircle(segment_from gd.Vector2, segment_to gd.Vector2, circle_position gd.Vector2, circle_radius gd.Float) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, segment_from)
	callframe.Arg(frame, segment_to)
	callframe.Arg(frame, circle_position)
	callframe.Arg(frame, circle_radius)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_segment_intersects_circle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Checks if the two segments ([param from_a], [param to_a]) and ([param from_b], [param to_b]) intersect. If yes, return the point of intersection as [Vector2]. If no intersection takes place, returns [code]null[/code].
*/
//go:nosplit
func (self class) SegmentIntersectsSegment(from_a gd.Vector2, to_a gd.Vector2, from_b gd.Vector2, to_b gd.Vector2) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, from_a)
	callframe.Arg(frame, to_a)
	callframe.Arg(frame, from_b)
	callframe.Arg(frame, to_b)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_segment_intersects_segment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Checks if the two lines ([param from_a], [param dir_a]) and ([param from_b], [param dir_b]) intersect. If yes, return the point of intersection as [Vector2]. If no intersection takes place, returns [code]null[/code].
[b]Note:[/b] The lines are specified using direction vectors, not end points.
*/
//go:nosplit
func (self class) LineIntersectsLine(from_a gd.Vector2, dir_a gd.Vector2, from_b gd.Vector2, dir_b gd.Vector2) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, from_a)
	callframe.Arg(frame, dir_a)
	callframe.Arg(frame, from_b)
	callframe.Arg(frame, dir_b)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_line_intersects_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Given the two 2D segments ([param p1], [param q1]) and ([param p2], [param q2]), finds those two points on the two segments that are closest to each other. Returns a [PackedVector2Array] that contains this point on ([param p1], [param q1]) as well the accompanying point on ([param p2], [param q2]).
*/
//go:nosplit
func (self class) GetClosestPointsBetweenSegments(p1 gd.Vector2, q1 gd.Vector2, p2 gd.Vector2, q2 gd.Vector2) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, p1)
	callframe.Arg(frame, q1)
	callframe.Arg(frame, p2)
	callframe.Arg(frame, q2)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_get_closest_points_between_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the 2D point on the 2D segment ([param s1], [param s2]) that is closest to [param point]. The returned point will always be inside the specified segment.
*/
//go:nosplit
func (self class) GetClosestPointToSegment(point gd.Vector2, s1 gd.Vector2, s2 gd.Vector2) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, s1)
	callframe.Arg(frame, s2)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_get_closest_point_to_segment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the 2D point on the 2D line defined by ([param s1], [param s2]) that is closest to [param point]. The returned point can be inside the segment ([param s1], [param s2]) or outside of it, i.e. somewhere on the line extending from the segment.
*/
//go:nosplit
func (self class) GetClosestPointToSegmentUncapped(point gd.Vector2, s1 gd.Vector2, s2 gd.Vector2) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, s1)
	callframe.Arg(frame, s2)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_get_closest_point_to_segment_uncapped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if [param point] is inside the triangle specified by [param a], [param b] and [param c].
*/
//go:nosplit
func (self class) PointIsInsideTriangle(point gd.Vector2, a gd.Vector2, b gd.Vector2, c gd.Vector2) bool {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, a)
	callframe.Arg(frame, b)
	callframe.Arg(frame, c)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_point_is_inside_triangle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [param polygon]'s vertices are ordered in clockwise order, otherwise returns [code]false[/code].
[b]Note:[/b] Assumes a Cartesian coordinate system where [code]+x[/code] is right and [code]+y[/code] is up. If using screen coordinates ([code]+y[/code] is down), the result will need to be flipped (i.e. a [code]true[/code] result will indicate counter-clockwise).
*/
//go:nosplit
func (self class) IsPolygonClockwise(polygon gd.PackedVector2Array) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_is_polygon_clockwise, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if [param point] is inside [param polygon] or if it's located exactly [i]on[/i] polygon's boundary, otherwise returns [code]false[/code].
*/
//go:nosplit
func (self class) IsPointInPolygon(point gd.Vector2, polygon gd.PackedVector2Array) bool {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_is_point_in_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Triangulates the polygon specified by the points in [param polygon]. Returns a [PackedInt32Array] where each triangle consists of three consecutive point indices into [param polygon] (i.e. the returned array will have [code]n * 3[/code] elements, with [code]n[/code] being the number of found triangles). Output triangles will always be counter clockwise, and the contour will be flipped if it's clockwise. If the triangulation did not succeed, an empty [PackedInt32Array] is returned.
*/
//go:nosplit
func (self class) TriangulatePolygon(polygon gd.PackedVector2Array) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_triangulate_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Triangulates the area specified by discrete set of [param points] such that no point is inside the circumcircle of any resulting triangle. Returns a [PackedInt32Array] where each triangle consists of three consecutive point indices into [param points] (i.e. the returned array will have [code]n * 3[/code] elements, with [code]n[/code] being the number of found triangles). If the triangulation did not succeed, an empty [PackedInt32Array] is returned.
*/
//go:nosplit
func (self class) TriangulateDelaunay(points gd.PackedVector2Array) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(points))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_triangulate_delaunay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Given an array of [Vector2]s, returns the convex hull as a list of points in counterclockwise order. The last point is the same as the first one.
*/
//go:nosplit
func (self class) ConvexHull(points gd.PackedVector2Array) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(points))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_convex_hull, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Decomposes the [param polygon] into multiple convex hulls and returns an array of [PackedVector2Array].
*/
//go:nosplit
func (self class) DecomposePolygonInConvex(polygon gd.PackedVector2Array) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_decompose_polygon_in_convex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Merges (combines) [param polygon_a] and [param polygon_b] and returns an array of merged polygons. This performs [constant OPERATION_UNION] between polygons.
The operation may result in an outer polygon (boundary) and multiple inner polygons (holes) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) MergePolygons(polygon_a gd.PackedVector2Array, polygon_b gd.PackedVector2Array) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon_a))
	callframe.Arg(frame, discreet.Get(polygon_b))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_merge_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Clips [param polygon_a] against [param polygon_b] and returns an array of clipped polygons. This performs [constant OPERATION_DIFFERENCE] between polygons. Returns an empty array if [param polygon_b] completely overlaps [param polygon_a].
If [param polygon_b] is enclosed by [param polygon_a], returns an outer polygon (boundary) and inner polygon (hole) which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) ClipPolygons(polygon_a gd.PackedVector2Array, polygon_b gd.PackedVector2Array) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon_a))
	callframe.Arg(frame, discreet.Get(polygon_b))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_clip_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Intersects [param polygon_a] with [param polygon_b] and returns an array of intersected polygons. This performs [constant OPERATION_INTERSECTION] between polygons. In other words, returns common area shared by polygons. Returns an empty array if no intersection occurs.
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) IntersectPolygons(polygon_a gd.PackedVector2Array, polygon_b gd.PackedVector2Array) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon_a))
	callframe.Arg(frame, discreet.Get(polygon_b))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_intersect_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Mutually excludes common area defined by intersection of [param polygon_a] and [param polygon_b] (see [method intersect_polygons]) and returns an array of excluded polygons. This performs [constant OPERATION_XOR] between polygons. In other words, returns all but common area between polygons.
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) ExcludePolygons(polygon_a gd.PackedVector2Array, polygon_b gd.PackedVector2Array) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon_a))
	callframe.Arg(frame, discreet.Get(polygon_b))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_exclude_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Clips [param polyline] against [param polygon] and returns an array of clipped polylines. This performs [constant OPERATION_DIFFERENCE] between the polyline and the polygon. This operation can be thought of as cutting a line with a closed shape.
*/
//go:nosplit
func (self class) ClipPolylineWithPolygon(polyline gd.PackedVector2Array, polygon gd.PackedVector2Array) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polyline))
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_clip_polyline_with_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Intersects [param polyline] with [param polygon] and returns an array of intersected polylines. This performs [constant OPERATION_INTERSECTION] between the polyline and the polygon. This operation can be thought of as chopping a line with a closed shape.
*/
//go:nosplit
func (self class) IntersectPolylineWithPolygon(polyline gd.PackedVector2Array, polygon gd.PackedVector2Array) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polyline))
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_intersect_polyline_with_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Inflates or deflates [param polygon] by [param delta] units (pixels). If [param delta] is positive, makes the polygon grow outward. If [param delta] is negative, shrinks the polygon inward. Returns an array of polygons because inflating/deflating may result in multiple discrete polygons. Returns an empty array if [param delta] is negative and the absolute value of it approximately exceeds the minimum bounding rectangle dimensions of the polygon.
Each polygon's vertices will be rounded as determined by [param join_type], see [enum PolyJoinType].
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
[b]Note:[/b] To translate the polygon's vertices specifically, multiply them to a [Transform2D]:
[codeblocks]
[gdscript]
var polygon = PackedVector2Array([Vector2(0, 0), Vector2(100, 0), Vector2(100, 100), Vector2(0, 100)])
var offset = Vector2(50, 50)
polygon = Transform2D(0, offset) * polygon
print(polygon) # prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/gdscript]
[csharp]
var polygon = new Vector2[] { new Vector2(0, 0), new Vector2(100, 0), new Vector2(100, 100), new Vector2(0, 100) };
var offset = new Vector2(50, 50);
polygon = new Transform2D(0, offset) * polygon;
GD.Print((Variant)polygon); // prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) OffsetPolygon(polygon gd.PackedVector2Array, delta gd.Float, join_type classdb.Geometry2DPolyJoinType) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon))
	callframe.Arg(frame, delta)
	callframe.Arg(frame, join_type)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_offset_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Inflates or deflates [param polyline] by [param delta] units (pixels), producing polygons. If [param delta] is positive, makes the polyline grow outward. Returns an array of polygons because inflating/deflating may result in multiple discrete polygons. If [param delta] is negative, returns an empty array.
Each polygon's vertices will be rounded as determined by [param join_type], see [enum PolyJoinType].
Each polygon's endpoints will be rounded as determined by [param end_type], see [enum PolyEndType].
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) OffsetPolyline(polyline gd.PackedVector2Array, delta gd.Float, join_type classdb.Geometry2DPolyJoinType, end_type classdb.Geometry2DPolyEndType) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polyline))
	callframe.Arg(frame, delta)
	callframe.Arg(frame, join_type)
	callframe.Arg(frame, end_type)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_offset_polyline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Given an array of [Vector2]s representing tiles, builds an atlas. The returned dictionary has two keys: [code]points[/code] is a [PackedVector2Array] that specifies the positions of each tile, [code]size[/code] contains the overall size of the whole atlas as [Vector2i].
*/
//go:nosplit
func (self class) MakeAtlas(sizes gd.PackedVector2Array) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(sizes))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_make_atlas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("Geometry2D", func(ptr gd.Object) any { return classdb.Geometry2D(ptr) })}
type PolyBooleanOperation = classdb.Geometry2DPolyBooleanOperation

const (
/*Create regions where either subject or clip polygons (or both) are filled.*/
	OperationUnion PolyBooleanOperation = 0
/*Create regions where subject polygons are filled except where clip polygons are filled.*/
	OperationDifference PolyBooleanOperation = 1
/*Create regions where both subject and clip polygons are filled.*/
	OperationIntersection PolyBooleanOperation = 2
/*Create regions where either subject or clip polygons are filled but not where both are filled.*/
	OperationXor PolyBooleanOperation = 3
)
type PolyJoinType = classdb.Geometry2DPolyJoinType

const (
/*Squaring is applied uniformally at all convex edge joins at [code]1 * delta[/code].*/
	JoinSquare PolyJoinType = 0
/*While flattened paths can never perfectly trace an arc, they are approximated by a series of arc chords.*/
	JoinRound PolyJoinType = 1
/*There's a necessary limit to mitered joins since offsetting edges that join at very acute angles will produce excessively long and narrow "spikes". For any given edge join, when miter offsetting would exceed that maximum distance, "square" joining is applied.*/
	JoinMiter PolyJoinType = 2
)
type PolyEndType = classdb.Geometry2DPolyEndType

const (
/*Endpoints are joined using the [enum PolyJoinType] value and the path filled as a polygon.*/
	EndPolygon PolyEndType = 0
/*Endpoints are joined using the [enum PolyJoinType] value and the path filled as a polyline.*/
	EndJoined PolyEndType = 1
/*Endpoints are squared off with no extension.*/
	EndButt PolyEndType = 2
/*Endpoints are squared off and extended by [code]delta[/code] units.*/
	EndSquare PolyEndType = 3
/*Endpoints are rounded off and extended by [code]delta[/code] units.*/
	EndRound PolyEndType = 4
)
