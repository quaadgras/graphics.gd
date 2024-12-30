// Package AABB provides a 3D axis-aligned bounding box.
package AABB

import (
	"unsafe"

	"graphics.gd/variant/Float"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Vector3"
)

// ʕ is a little ternary operator for porting C code.
func ʕ[T any](q bool, a T, b T) T {
	if q {
		return a
	}
	return b
}

// AABB built-in Variant type represents an axis-aligned bounding box in a
// 3D space. It is defined by its position and size, which are Vector3.
// It is frequently used for fast overlap tests (see intersects). Although
// AABB itself is axis-aligned, it can be combined with Transform3D to
// represent a rotated or skewed bounding box.
//
// It uses floating-point coordinates. The 2D counterpart to AABB is Rect2.
// There is no version of AABB that uses integer coordinates.
//
// Note: Negative values for size are not supported. With negative size,
// most functions do not work correctly. Use [Abs] to get an equivalent AABB
// with a non-negative size.
type PositionSize = struct {
	Position Vector3.XYZ // The origin point. This is usually the corner on the bottom-left and back of the bounding box.

	// The bounding box's width, height, and depth starting from position. Setting
	// this value also affects the end point.
	//
	// Note: It's recommended setting the width, height, and depth to non-negative
	// values. This is because most methods in Godot assume that the position is
	// the bottom-left-back corner, and the end is the top-right-forward corner.
	// To get an equivalent bounding box with non-negative size, use [Abs].
	Size Vector3.XYZ
}

// End returns the ending point. This is usually the corner on the top-right and forward
// of the bounding box, and is equivalent to position + size.
func End(rect PositionSize) Vector3.XYZ {
	return Vector3.Add(rect.Position, rect.Size)
}

// Abs returns an AABB equivalent to this bounding box, with its width, height, and depth
// modified to be non-negative values.
func Abs(a PositionSize) PositionSize { //gd:AABB.abs
	return PositionSize{
		Position: Vector3.New(
			a.Position.X+min(a.Size.X, 0),
			a.Position.Y+min(a.Size.Y, 0),
			a.Position.Z+min(a.Size.Z, 0),
		),
		Size: Vector3.Abs(a.Size),
	}
}

// Inside returns true if a completely encloses b.
func Inside(a, b PositionSize) bool { //gd:AABB.encloses
	var src_min = a.Position
	var src_max = Vector3.Add(a.Position, a.Size)
	var dst_min = b.Position
	var dst_max = Vector3.Add(b.Position, b.Size)
	return ((src_min.X <= dst_min.X) &&
		(src_max.X > dst_max.X) &&
		(src_min.Y <= dst_min.Y) &&
		(src_max.Y > dst_max.Y) &&
		(src_min.Z <= dst_min.Z) &&
		(src_max.Z > dst_max.Z))
}

// Expand returns a copy of this AABB expanded to include a given point.
//
//	// position (-3, 2, 0), size (1, 1, 1)
//	var box = AABB{Vector3{-3,2,0}, Vector3{1,1,1}}
//	// position (-3, -1, 0), size (3, 4, 2), so we fit both the original AABB and Vector3(0, -1, 2)
//	var box2 = box.Expand(Vector3{0, -1, 2})
func ExpandTo(to Vector3.XYZ, a PositionSize) PositionSize { //gd:AABB.expand
	var begin = a.Position
	var end = Vector3.Add(a.Position, a.Size)
	if to.X < begin.X {
		begin.X = to.X
	}
	if to.Y < begin.Y {
		begin.Y = to.Y
	}
	if to.Z < begin.Z {
		begin.Z = to.Z
	}
	if to.X > end.X {
		end.X = to.X
	}
	if to.Y > end.Y {
		end.Y = to.Y
	}
	if to.Z > end.Z {
		end.Z = to.Z
	}
	return PositionSize{
		Position: begin,
		Size:     Vector3.Sub(end, begin),
	}
}

// Center returns the center of the AABB, which is equal to position + (size / 2).
func Center(a PositionSize) Vector3.XYZ { //gd:AABB.get_center
	return Vector3.Add(a.Position, Vector3.DivX(a.Size, 2))
}

// Endpoint gets the position of the 8 endpoints of the AABB in space.
func Endpoint(a PositionSize, idx int) Vector3.XYZ { //gd:AABB.get_endpoint
	switch idx {
	case 0:
		return Vector3.XYZ{a.Position.X, a.Position.Y, a.Position.Z}
	case 1:
		return Vector3.XYZ{a.Position.X, a.Position.Y, a.Position.Z + a.Size.Z}
	case 2:
		return Vector3.XYZ{a.Position.X, a.Position.Y + a.Size.Y, a.Position.Z}
	case 3:
		return Vector3.XYZ{a.Position.X, a.Position.Y + a.Size.Y, a.Position.Z + a.Size.Z}
	case 4:
		return Vector3.XYZ{a.Position.X + a.Size.X, a.Position.Y, a.Position.Z}
	case 5:
		return Vector3.XYZ{a.Position.X + a.Size.X, a.Position.Y, a.Position.Z + a.Size.Z}
	case 6:
		return Vector3.XYZ{a.Position.X + a.Size.X, a.Position.Y + a.Size.Y, a.Position.Z}
	case 7:
		return Vector3.XYZ{a.Position.X + a.Size.X, a.Position.Y + a.Size.Y, a.Position.Z + a.Size.Z}
	default:
		panic("AABB.GetEndpoint(): invalid index")
	}
}

// LongestAxis returns the normalized longest axis of the AABB.
func LongestAxis(a PositionSize) Vector3.XYZ { //gd:AABB.get_longest_axis
	var (
		axis     = Vector3.XYZ{1, 0, 0}
		max_size = a.Size.X
	)
	if a.Size.Y > max_size {
		axis = Vector3.XYZ{0, 1, 0}
		max_size = a.Size.Y
	}
	if a.Size.Z > max_size {
		axis = Vector3.XYZ{0, 0, 1}
	}
	return axis
}

// LongestAxisIndex returns the index of the longest axis of the AABB (according to [Axis] constants).
func LongestAxisIndex(a PositionSize) Vector3.Axis { //gd:AABB.get_longest_axis_index
	var axis = Vector3.X
	var max_size = a.Size.X
	if a.Size.Y > max_size {
		axis = Vector3.Y
		max_size = a.Size.Y
	}
	if a.Size.Z > max_size {
		axis = Vector3.Z
	}
	return axis
}

// LongestAxisSize returns the scalar length of the longest axis of the AABB.
func LongestAxisSize(a PositionSize) Float.X { //gd:AABB.get_longest_axis_size
	var max_size = a.Size.X
	if a.Size.Y > max_size {
		max_size = a.Size.Y
	}
	if a.Size.Z > max_size {
		max_size = a.Size.Z
	}
	return max_size
}

// ShortestAxis returns the normalized shortest axis of the AABB.
func ShortestAxis(a PositionSize) Vector3.XYZ { //gd:AABB.get_shortest_axis
	var (
		axis     = Vector3.XYZ{1, 0, 0}
		min_size = a.Size.X
	)
	if a.Size.Y < min_size {
		axis = Vector3.XYZ{0, 1, 0}
		min_size = a.Size.Y
	}
	if a.Size.Z < min_size {
		axis = Vector3.XYZ{0, 0, 1}
	}
	return axis
}

// ShortestAxisIndex returns the index of the shortest axis of the AABB (according to [Axis] constants).
func ShortestAxisIndex(a PositionSize) Vector3.Axis { //gd:AABB.get_shortest_axis_index
	var axis = Vector3.X
	var min_size = a.Size.X
	if a.Size.Y < min_size {
		axis = Vector3.Y
		min_size = a.Size.Y
	}
	if a.Size.Z < min_size {
		axis = Vector3.Z
	}
	return axis
}

// ShortestAxisSize returns the scalar length of the shortest axis of the AABB.
func ShortestAxisSize(a PositionSize) Float.X { //gd:AABB.get_shortest_axis_size
	var min_size = a.Size.X
	if a.Size.Y < min_size {
		min_size = a.Size.Y
	}
	if a.Size.Z < min_size {
		min_size = a.Size.Z
	}
	return min_size
}

// Support returns the vertex of the AABB that's the farthest in a given direction. This point is commonly
// known as the support point in collision detection algorithms.
func Support(dir Vector3.XYZ, a PositionSize) Vector3.XYZ { //gd:AABB.get_support
	var (
		half_extents = Vector3.MulX(a.Size, 0.5)
		ofs          = Vector3.Add(a.Position, half_extents)
	)
	return Vector3.Add(Vector3.XYZ{
		ʕ(dir.X > 0, half_extents.X, -half_extents.X),
		ʕ(dir.Y > 0, half_extents.Y, -half_extents.Y),
		ʕ(dir.Z > 0, half_extents.Z, -half_extents.Z),
	}, ofs)
}

// Volume returns the volume of the AABB.
func Volume(a PositionSize) Float.X { return a.Size.X * a.Size.Y * a.Size.Z } //gd:AABB.get_volume

// Expand returns a copy of the AABB grown a given number of units towards all the sides.
func Expand[X Float.Any](a PositionSize, by X) PositionSize { //gd:AABB.grow
	a.Position = Vector3.SubX(a.Position, by)
	a.Size = Vector3.AddX(a.Size, by*2)
	return a
}

// HasPoint returns true if the AABB contains a point. Points on the faces of the AABB are considered included,
// though float-point precision errors may impact the accuracy of such checks.
//
// Note: This method is not reliable for AABB with a negative size. Use abs to get a positive sized equivalent
// AABB to check for contained points.
func HasPoint(point Vector3.XYZ, a PositionSize) bool { //gd:AABB.has_point
	if point.X < a.Position.X {
		return false
	}
	if point.Y < a.Position.Y {
		return false
	}
	if point.Z < a.Position.Z {
		return false
	}
	if point.X > a.Position.X+a.Size.X {
		return false
	}
	if point.Y > a.Position.Y+a.Size.Y {
		return false
	}
	if point.Z > a.Position.Z+a.Size.Z {
		return false
	}
	return true
}

// HasSurface returns true if the AABB has a surface or a length, and false if the AABB is empty (all components of
// size are zero or negative).
func HasSurface(a PositionSize) bool { return a.Size.X > 0 || a.Size.Y > 0 || a.Size.Z > 0 } //gd:AABB.has_surface

// HasVolume returns true if the AABB has a volume, and false if the AABB is flat, empty, or has a negative size.
func HasVolume(a PositionSize) bool { return a.Size.X > 0 && a.Size.Y > 0 && a.Size.Z > 0 } //gd:AABB.has_volume

// Intersection returns the intersection between two AABB. An empty AABB (size (0, 0, 0)) is returned on failure.
func Intersection(a, b PositionSize) PositionSize { //gd:AABB.intersection
	var (
		src_min = a.Position
		src_max = Vector3.Add(a.Position, a.Size)
		dst_min = b.Position
		dst_max = Vector3.Add(b.Position, a.Size)
	)
	var (
		min, max Vector3.XYZ
	)
	if src_min.X > dst_max.X || src_max.X < dst_min.X {
		return PositionSize{}
	} else {
		min.X = ʕ(src_min.X > dst_min.X, src_min.X, dst_min.X)
		max.X = ʕ(src_max.X < dst_max.X, src_max.X, dst_max.X)
	}
	if src_min.Y > dst_max.Y || src_max.Y < dst_min.Y {
		return PositionSize{}
	} else {
		min.Y = ʕ(src_min.Y > dst_min.Y, src_min.Y, dst_min.Y)
		max.Y = ʕ(src_max.Y < dst_max.Y, src_max.Y, dst_max.Y)
	}
	if src_min.Z > dst_max.Z || src_max.Z < dst_min.Z {
		return PositionSize{}
	} else {
		min.Z = ʕ(src_min.Z > dst_min.Z, src_min.Z, dst_min.Z)
		max.Z = ʕ(src_max.Z < dst_max.Z, src_max.Z, dst_max.Z)
	}
	return PositionSize{min, Vector3.Sub(max, min)}
}

// Intersects returns true if the AABB overlaps with another.
func Intersects(a, b PositionSize) bool { //gd:AABB.intersects
	if a.Position.X >= (b.Position.X + b.Size.X) {
		return false
	}
	if (a.Position.X + a.Size.X) <= b.Position.X {
		return false
	}
	if a.Position.Y >= (b.Position.Y + b.Size.Y) {
		return false
	}
	if (a.Position.Y + a.Size.Y) <= b.Position.Y {
		return false
	}
	if a.Position.Z >= (b.Position.Z + b.Size.Z) {
		return false
	}
	if (a.Position.Z + a.Size.Z) <= b.Position.Z {
		return false
	}
	return true
}

// IntersectsPlane returns true if the AABB is on both sides of a plane.
func IntersectsPlane(a PositionSize, plane Plane.NormalD) bool { //gd:AABB.intersects_plane
	var points = [8]Vector3.XYZ{
		{a.Position.X, a.Position.Y, a.Position.Z},
		{a.Position.X, a.Position.Y, a.Position.Z + a.Size.Z},
		{a.Position.X, a.Position.Y + a.Size.Y, a.Position.Z},
		{a.Position.X, a.Position.Y + a.Size.Y, a.Position.Z + a.Size.Z},
		{a.Position.X + a.Size.X, a.Position.Y, a.Position.Z},
		{a.Position.X + a.Size.X, a.Position.Y, a.Position.Z + a.Size.Z},
		{a.Position.X + a.Size.X, a.Position.Y + a.Size.Y, a.Position.Z},
		{a.Position.X + a.Size.X, a.Position.Y + a.Size.Y, a.Position.Z + a.Size.Z},
	}
	var (
		over  = false
		under = false
	)
	for i := range points {
		if Plane.DistanceToPoint(points[i], plane) > 0 {
			over = true
		} else {
			under = true
		}
	}
	return under && over
}

// IntersectsRay returns the point of intersection of the given ray with this AABB along with the normal
// or false if there is no intersection. Ray length is infinite.
func IntersectsRay(a PositionSize, fromv, dirv Vector3.XYZ) (clip, normal Vector3.XYZ, ok bool) { //gd:AABB.intersects_ray
	var (
		c1 = (*[3]Float.X)(unsafe.Pointer(&clip))
		n1 = (*[3]Float.X)(unsafe.Pointer(&normal))

		c2   [3]Float.X
		end               = Vector3.Add(a.Position, a.Size)
		near Float.X      = -1e20
		far  Float.X      = 1e20
		axis Vector3.Axis = 0

		from = *(*[3]Float.X)(unsafe.Pointer(&fromv))
		dir  = *(*[3]Float.X)(unsafe.Pointer(&dirv))
	)
	for i := Vector3.Axis(0); i < 3; i++ {
		if dir[i] == 0 {
			if (from[i] < Vector3.Index(a.Position, i)) || (from[i] > Vector3.Index(end, i)) {
				return Vector3.Zero, Vector3.Zero, false
			}
		} else { // ray not parallel to planes in this direction
			c1[i] = (Vector3.Index(a.Position, i) - from[i]) / dir[i]
			c2[i] = (Vector3.Index(end, i) - from[i]) / dir[i]
			if c1[i] > c2[i] {
				*c1, c2 = c2, *c1
			}
			if c1[i] > near {
				near = c1[i]
				axis = i
			}
			if c2[i] < far {
				far = c2[i]
			}
			if (near > far) || (far < 0) {
				return Vector3.Zero, Vector3.Zero, false
			}
		}
	}
	n1[axis] = ʕ[Float.X](dir[axis] != 0, -1, 1)
	return clip, normal, true
}

// IntesectsSegment returns the point of intersection between from and to along with this AABB along with the normal
// or false if there is no intersection.
func IntesectsSegment(a PositionSize, from, to Vector3.XYZ) (clip, normal Vector3.XYZ, ok bool) { //gd:AABB.intersects_segment
	var (
		min  Float.X      = 0
		max  Float.X      = 1
		axis Vector3.Axis = 0
		sign Float.X      = 0

		n1 = (*[3]Float.X)(unsafe.Pointer(&normal))
	)
	for i := Vector3.Axis(0); i < 3; i++ {
		var (
			seg_from   = Vector3.Index(from, i)
			seg_to     = Vector3.Index(to, i)
			box_begin  = Vector3.Index(a.Position, i)
			box_end    = box_begin + Vector3.Index(a.Size, i)
			cmin, cmax Float.X
			csign      Float.X
		)
		if seg_from < seg_to {
			if seg_from > box_end || seg_to < box_begin {
				return Vector3.Zero, Vector3.Zero, false
			}
			var length = seg_to - seg_from
			cmin = ʕ(seg_from < box_begin, ((box_begin - seg_from) / length), 0)
			cmax = ʕ(seg_to > box_end, ((box_end - seg_from) / length), 1)
			csign = -1.0
		} else {
			if seg_to > box_end || seg_from < box_begin {
				return Vector3.Zero, Vector3.Zero, false
			}
			var length = seg_to - seg_from
			cmin = ʕ(seg_from > box_end, (box_end-seg_from)/length, 0)
			cmax = ʕ(seg_to < box_begin, (box_begin-seg_from)/length, 1)
			csign = 1.0
		}
		if cmin > min {
			min = cmin
			axis = i
			sign = csign
		}
		if cmax < max {
			max = cmax
		}
		if max < min {
			return Vector3.Zero, Vector3.Zero, false
		}
	}
	var (
		rel = Vector3.Sub(to, from)
	)
	n1[axis] = sign
	return Vector3.Add(from, Vector3.MulX(rel, min)), normal, true
}

// IsAproximatelyEqual returns true if this AABB and other are approximately equal, by running
// [IsApproximatelyEqual] on each component.
func IsAproximatelyEqual(a, b PositionSize) bool { //gd:AABB.is_equal_approx
	return Vector3.IsApproximatelyEqual(a.Position, b.Position) && Vector3.IsApproximatelyEqual(a.Size, b.Size)
}

// IsFinite returns true if this AABB is finite, by calling [IsFinite] on each component.
func IsFinite(a PositionSize) bool { return Vector3.IsFinite(a.Position) && Vector3.IsFinite(a.Size) } //gd:AABB.is_finite

// Merge returns the smallest AABB enclosing both this AABB and b.
func Merge(a, b PositionSize) PositionSize { //gd:AABB.merge
	var beg_1, beg_2 Vector3.XYZ
	var end_1, end_2 Vector3.XYZ
	var min, max Vector3.XYZ
	beg_1 = a.Position
	beg_2 = b.Position
	end_1 = Vector3.Add(a.Size, beg_1)
	end_2 = Vector3.Add(b.Size, beg_2)
	min.X = ʕ(beg_1.X < beg_2.X, beg_1.X, beg_2.X)
	min.Y = ʕ(beg_1.Y < beg_2.Y, beg_1.Y, beg_2.Y)
	min.Z = ʕ(beg_1.Z < beg_2.Z, beg_1.Z, beg_2.Z)
	max.X = ʕ(end_1.X > end_2.X, end_1.X, end_2.X)
	max.Y = ʕ(end_1.Y > end_2.Y, end_1.Y, end_2.Y)
	max.Z = ʕ(end_1.Z > end_2.Z, end_1.Z, end_2.Z)
	return PositionSize{
		Position: min,
		Size:     Vector3.Sub(max, min),
	}
}
