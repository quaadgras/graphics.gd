package Geometry3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Provides a set of helper functions to create geometric shapes, compute intersections between shapes, and process various other geometric operations in 3D.

*/
type Simple [1]classdb.Geometry3D
func (self Simple) ComputeConvexMeshPoints(planes gd.ArrayOf[gd.Plane]) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).ComputeConvexMeshPoints(gc, planes))
}
func (self Simple) BuildBoxPlanes(extents gd.Vector3) gd.ArrayOf[gd.Plane] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Plane](Expert(self).BuildBoxPlanes(gc, extents))
}
func (self Simple) BuildCylinderPlanes(radius float64, height float64, sides int, axis gd.Vector3Axis) gd.ArrayOf[gd.Plane] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Plane](Expert(self).BuildCylinderPlanes(gc, gd.Float(radius), gd.Float(height), gd.Int(sides), axis))
}
func (self Simple) BuildCapsulePlanes(radius float64, height float64, sides int, lats int, axis gd.Vector3Axis) gd.ArrayOf[gd.Plane] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Plane](Expert(self).BuildCapsulePlanes(gc, gd.Float(radius), gd.Float(height), gd.Int(sides), gd.Int(lats), axis))
}
func (self Simple) GetClosestPointsBetweenSegments(p1 gd.Vector3, p2 gd.Vector3, q1 gd.Vector3, q2 gd.Vector3) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).GetClosestPointsBetweenSegments(gc, p1, p2, q1, q2))
}
func (self Simple) GetClosestPointToSegment(point gd.Vector3, s1 gd.Vector3, s2 gd.Vector3) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetClosestPointToSegment(point, s1, s2))
}
func (self Simple) GetClosestPointToSegmentUncapped(point gd.Vector3, s1 gd.Vector3, s2 gd.Vector3) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetClosestPointToSegmentUncapped(point, s1, s2))
}
func (self Simple) GetTriangleBarycentricCoords(point gd.Vector3, a gd.Vector3, b gd.Vector3, c gd.Vector3) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(Expert(self).GetTriangleBarycentricCoords(point, a, b, c))
}
func (self Simple) RayIntersectsTriangle(from gd.Vector3, dir gd.Vector3, a gd.Vector3, b gd.Vector3, c gd.Vector3) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).RayIntersectsTriangle(gc, from, dir, a, b, c))
}
func (self Simple) SegmentIntersectsTriangle(from gd.Vector3, to gd.Vector3, a gd.Vector3, b gd.Vector3, c gd.Vector3) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(Expert(self).SegmentIntersectsTriangle(gc, from, to, a, b, c))
}
func (self Simple) SegmentIntersectsSphere(from gd.Vector3, to gd.Vector3, sphere_position gd.Vector3, sphere_radius float64) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).SegmentIntersectsSphere(gc, from, to, sphere_position, gd.Float(sphere_radius)))
}
func (self Simple) SegmentIntersectsCylinder(from gd.Vector3, to gd.Vector3, height float64, radius float64) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).SegmentIntersectsCylinder(gc, from, to, gd.Float(height), gd.Float(radius)))
}
func (self Simple) SegmentIntersectsConvex(from gd.Vector3, to gd.Vector3, planes gd.ArrayOf[gd.Plane]) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).SegmentIntersectsConvex(gc, from, to, planes))
}
func (self Simple) ClipPolygon(points gd.PackedVector3Array, plane gd.Plane) gd.PackedVector3Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedVector3Array(Expert(self).ClipPolygon(gc, points, plane))
}
func (self Simple) TetrahedralizeDelaunay(points gd.PackedVector3Array) gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).TetrahedralizeDelaunay(gc, points))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Geometry3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Calculates and returns all the vertex points of a convex shape defined by an array of [param planes].
*/
//go:nosplit
func (self class) ComputeConvexMeshPoints(ctx gd.Lifetime, planes gd.ArrayOf[gd.Plane]) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(planes))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_compute_convex_mesh_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with 6 [Plane]s that describe the sides of a box centered at the origin. The box size is defined by [param extents], which represents one (positive) corner of the box (i.e. half its actual size).
*/
//go:nosplit
func (self class) BuildBoxPlanes(ctx gd.Lifetime, extents gd.Vector3) gd.ArrayOf[gd.Plane] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_build_box_planes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Plane](ret)
}
/*
Returns an array of [Plane]s closely bounding a faceted cylinder centered at the origin with radius [param radius] and height [param height]. The parameter [param sides] defines how many planes will be generated for the round part of the cylinder. The parameter [param axis] describes the axis along which the cylinder is oriented (0 for X, 1 for Y, 2 for Z).
*/
//go:nosplit
func (self class) BuildCylinderPlanes(ctx gd.Lifetime, radius gd.Float, height gd.Float, sides gd.Int, axis gd.Vector3Axis) gd.ArrayOf[gd.Plane] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	callframe.Arg(frame, height)
	callframe.Arg(frame, sides)
	callframe.Arg(frame, axis)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_build_cylinder_planes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Plane](ret)
}
/*
Returns an array of [Plane]s closely bounding a faceted capsule centered at the origin with radius [param radius] and height [param height]. The parameter [param sides] defines how many planes will be generated for the side part of the capsule, whereas [param lats] gives the number of latitudinal steps at the bottom and top of the capsule. The parameter [param axis] describes the axis along which the capsule is oriented (0 for X, 1 for Y, 2 for Z).
*/
//go:nosplit
func (self class) BuildCapsulePlanes(ctx gd.Lifetime, radius gd.Float, height gd.Float, sides gd.Int, lats gd.Int, axis gd.Vector3Axis) gd.ArrayOf[gd.Plane] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	callframe.Arg(frame, height)
	callframe.Arg(frame, sides)
	callframe.Arg(frame, lats)
	callframe.Arg(frame, axis)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_build_capsule_planes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Plane](ret)
}
/*
Given the two 3D segments ([param p1], [param p2]) and ([param q1], [param q2]), finds those two points on the two segments that are closest to each other. Returns a [PackedVector3Array] that contains this point on ([param p1], [param p2]) as well the accompanying point on ([param q1], [param q2]).
*/
//go:nosplit
func (self class) GetClosestPointsBetweenSegments(ctx gd.Lifetime, p1 gd.Vector3, p2 gd.Vector3, q1 gd.Vector3, q2 gd.Vector3) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, p1)
	callframe.Arg(frame, p2)
	callframe.Arg(frame, q1)
	callframe.Arg(frame, q2)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_get_closest_points_between_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the 3D point on the 3D segment ([param s1], [param s2]) that is closest to [param point]. The returned point will always be inside the specified segment.
*/
//go:nosplit
func (self class) GetClosestPointToSegment(point gd.Vector3, s1 gd.Vector3, s2 gd.Vector3) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, s1)
	callframe.Arg(frame, s2)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_get_closest_point_to_segment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the 3D point on the 3D line defined by ([param s1], [param s2]) that is closest to [param point]. The returned point can be inside the segment ([param s1], [param s2]) or outside of it, i.e. somewhere on the line extending from the segment.
*/
//go:nosplit
func (self class) GetClosestPointToSegmentUncapped(point gd.Vector3, s1 gd.Vector3, s2 gd.Vector3) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, s1)
	callframe.Arg(frame, s2)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_get_closest_point_to_segment_uncapped, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [Vector3] containing weights based on how close a 3D position ([param point]) is to a triangle's different vertices ([param a], [param b] and [param c]). This is useful for interpolating between the data of different vertices in a triangle. One example use case is using this to smoothly rotate over a mesh instead of relying solely on face normals.
[url=https://en.wikipedia.org/wiki/Barycentric_coordinate_system]Here is a more detailed explanation of barycentric coordinates.[/url]
*/
//go:nosplit
func (self class) GetTriangleBarycentricCoords(point gd.Vector3, a gd.Vector3, b gd.Vector3, c gd.Vector3) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, a)
	callframe.Arg(frame, b)
	callframe.Arg(frame, c)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_get_triangle_barycentric_coords, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Tests if the 3D ray starting at [param from] with the direction of [param dir] intersects the triangle specified by [param a], [param b] and [param c]. If yes, returns the point of intersection as [Vector3]. If no intersection takes place, returns [code]null[/code].
*/
//go:nosplit
func (self class) RayIntersectsTriangle(ctx gd.Lifetime, from gd.Vector3, dir gd.Vector3, a gd.Vector3, b gd.Vector3, c gd.Vector3) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, dir)
	callframe.Arg(frame, a)
	callframe.Arg(frame, b)
	callframe.Arg(frame, c)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_ray_intersects_triangle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Tests if the segment ([param from], [param to]) intersects the triangle [param a], [param b], [param c]. If yes, returns the point of intersection as [Vector3]. If no intersection takes place, returns [code]null[/code].
*/
//go:nosplit
func (self class) SegmentIntersectsTriangle(ctx gd.Lifetime, from gd.Vector3, to gd.Vector3, a gd.Vector3, b gd.Vector3, c gd.Vector3) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, a)
	callframe.Arg(frame, b)
	callframe.Arg(frame, c)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_segment_intersects_triangle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Checks if the segment ([param from], [param to]) intersects the sphere that is located at [param sphere_position] and has radius [param sphere_radius]. If no, returns an empty [PackedVector3Array]. If yes, returns a [PackedVector3Array] containing the point of intersection and the sphere's normal at the point of intersection.
*/
//go:nosplit
func (self class) SegmentIntersectsSphere(ctx gd.Lifetime, from gd.Vector3, to gd.Vector3, sphere_position gd.Vector3, sphere_radius gd.Float) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, sphere_position)
	callframe.Arg(frame, sphere_radius)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_segment_intersects_sphere, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Checks if the segment ([param from], [param to]) intersects the cylinder with height [param height] that is centered at the origin and has radius [param radius]. If no, returns an empty [PackedVector3Array]. If an intersection takes place, the returned array contains the point of intersection and the cylinder's normal at the point of intersection.
*/
//go:nosplit
func (self class) SegmentIntersectsCylinder(ctx gd.Lifetime, from gd.Vector3, to gd.Vector3, height gd.Float, radius gd.Float) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, height)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_segment_intersects_cylinder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Given a convex hull defined though the [Plane]s in the array [param planes], tests if the segment ([param from], [param to]) intersects with that hull. If an intersection is found, returns a [PackedVector3Array] containing the point the intersection and the hull's normal. Otherwise, returns an empty array.
*/
//go:nosplit
func (self class) SegmentIntersectsConvex(ctx gd.Lifetime, from gd.Vector3, to gd.Vector3, planes gd.ArrayOf[gd.Plane]) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, mmm.Get(planes))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_segment_intersects_convex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clips the polygon defined by the points in [param points] against the [param plane] and returns the points of the clipped polygon.
*/
//go:nosplit
func (self class) ClipPolygon(ctx gd.Lifetime, points gd.PackedVector3Array, plane gd.Plane) gd.PackedVector3Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	callframe.Arg(frame, plane)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_clip_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedVector3Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Tetrahedralizes the volume specified by a discrete set of [param points] in 3D space, ensuring that no point lies within the circumsphere of any resulting tetrahedron. The method returns a [PackedInt32Array] where each tetrahedron consists of four consecutive point indices into the [param points] array (resulting in an array with [code]n * 4[/code] elements, where [code]n[/code] is the number of tetrahedra found). If the tetrahedralization is unsuccessful, an empty [PackedInt32Array] is returned.
*/
//go:nosplit
func (self class) TetrahedralizeDelaunay(ctx gd.Lifetime, points gd.PackedVector3Array) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(points))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Geometry3D.Bind_tetrahedralize_delaunay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsGeometry3D() Expert { return self[0].AsGeometry3D() }


//go:nosplit
func (self Simple) AsGeometry3D() Simple { return self[0].AsGeometry3D() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Geometry3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
