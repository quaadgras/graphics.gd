//go:build !generate

package spatial

/*
AABB consists of a position, a size, and several utility functions. It is typically used for fast overlap tests.

It uses floating-point coordinates. The 2D counterpart to [AABB] is [Rect2].

Negative values for size are not supported and will not work for most methods. Use abs to get an [AABB] with a positive size.

Note: Unlike [Rect2], [AABB] does not have a variant that uses integer coordinates.
*/
type AABB struct {
	Position Vector3
	Size     Vector3
}

// "Fields"

func (a AABB) End() Vector3        { return a.Position.Add(a.Size) }
func (a *AABB) SetEnd(end Vector3) { a.Size = end.Sub(a.Position) }

// "Methods"

// Abs returns an [AABB] with equivalent position and size, modified so that the most-negative corner is the origin and the
// size is positive.
func (a AABB) Abs() AABB {
	return AABB{
		Position: Vector3{a.Position[X] + min(a.Size[X], 0), a.Position[Y] + min(a.Size[Y], 0), a.Position[Z] + min(a.Size[Z], 0)},
		Size:     a.Size.Abs(),
	}
}

// Encloses returns true if this [AABB] completely encloses another one.
func (a AABB) Encloses(b AABB) bool {
	var src_min = a.Position
	var src_max = a.Position.Add(a.Size)
	var dst_min = b.Position
	var dst_max = b.Position.Add(b.Size)
	return ((src_min[X] <= dst_min[X]) &&
		(src_max[X] > dst_max[X]) &&
		(src_min[Y] <= dst_min[Y]) &&
		(src_max[Y] > dst_max[Y]) &&
		(src_min[Z] <= dst_min[Z]) &&
		(src_max[Z] > dst_max[Z]))
}

// Expand returns a copy of this [AABB] expanded to include a given point.
//
//	// position (-3, 2, 0), size (1, 1, 1)
//	var box = AABB{Vector3{-3,2,0}, Vector3{1,1,1}}
//	// position (-3, -1, 0), size (3, 4, 2), so we fit both the original AABB and Vector3(0, -1, 2)
//	var box2 = box.Expand(Vector3{0, -1, 2})
func (a AABB) Expand(to Vector3) AABB {
	var begin = a.Position
	var end = a.Position.Add(a.Size)
	if to[X] < begin[X] {
		begin[X] = to[X]
	}
	if to[Y] < begin[Y] {
		begin[Y] = to[Y]
	}
	if to[Z] < begin[Z] {
		begin[Z] = to[Z]
	}
	if to[X] > end[X] {
		end[X] = to[X]
	}
	if to[Y] > end[Y] {
		end[Y] = to[Y]
	}
	if to[Z] > end[Z] {
		end[Z] = to[Z]
	}
	return AABB{
		Position: begin,
		Size:     end.Sub(begin),
	}
}

// GetCenter returns the center of the [AABB], which is equal to position + (size / 2).
func (a AABB) GetCenter() Vector3 {
	return a.Position.Add(a.Size.Divf(2))
}

// GetEndpoint gets the position of the 8 endpoints of the [AABB] in space.
func (a AABB) GetEndpoint(idx int) Vector3 {
	switch idx {
	case 0:
		return Vector3{a.Position[X], a.Position[Y], a.Position[Z]}
	case 1:
		return Vector3{a.Position[X], a.Position[Y], a.Position[Z] + a.Size[Z]}
	case 2:
		return Vector3{a.Position[X], a.Position[Y] + a.Size[Y], a.Position[Z]}
	case 3:
		return Vector3{a.Position[X], a.Position[Y] + a.Size[Y], a.Position[Z] + a.Size[Z]}
	case 4:
		return Vector3{a.Position[X] + a.Size[X], a.Position[Y], a.Position[Z]}
	case 5:
		return Vector3{a.Position[X] + a.Size[X], a.Position[Y], a.Position[Z] + a.Size[Z]}
	case 6:
		return Vector3{a.Position[X] + a.Size[X], a.Position[Y] + a.Size[Y], a.Position[Z]}
	case 7:
		return Vector3{a.Position[X] + a.Size[X], a.Position[Y] + a.Size[Y], a.Position[Z] + a.Size[Z]}
	default:
		panic("AABB.GetEndpoint(): invalid index")
	}
}

// GetLongestAxis returns the normalized longest axis of the [AABB].
func (a AABB) GetLongestAxis() Vector3 {
	var (
		axis     = Vector3{1, 0, 0}
		max_size = a.Size[X]
	)
	if a.Size[Y] > max_size {
		axis = Vector3{0, 1, 0}
		max_size = a.Size[Y]
	}
	if a.Size[Z] > max_size {
		axis = Vector3{0, 0, 1}
	}
	return axis
}

// GetLongestAxisIndex returns the index of the longest axis of the [AABB] (according to [Axis] constants).
func (a AABB) GetLongestAxisIndex() Axis {
	var axis = X
	var max_size = a.Size[X]
	if a.Size[Y] > max_size {
		axis = Y
		max_size = a.Size[Y]
	}
	if a.Size[Z] > max_size {
		axis = Z
	}
	return axis
}

// GetLongestAxisSize returns the scalar length of the longest axis of the AABB.
func (a AABB) GetLongestAxisSize() Float {
	var max_size = a.Size[x]
	if a.Size[Y] > max_size {
		max_size = a.Size[Y]
	}
	if a.Size[Z] > max_size {
		max_size = a.Size[Z]
	}
	return Float(max_size)
}

// GetShortestAxis returns the normalized shortest axis of the [AABB].
func (a AABB) GetShortestAxis() Vector3 {
	var (
		axis     = Vector3{1, 0, 0}
		min_size = a.Size[X]
	)
	if a.Size[Y] < min_size {
		axis = Vector3{0, 1, 0}
		min_size = a.Size[Y]
	}
	if a.Size[Z] < min_size {
		axis = Vector3{0, 0, 1}
	}
	return axis
}

// GetShortestAxisIndex returns the index of the shortest axis of the [AABB] (according to [Axis] constants).
func (a AABB) GetShortestAxisIndex() Axis {
	var axis = X
	var min_size = a.Size[X]
	if a.Size[Y] < min_size {
		axis = Y
		min_size = a.Size[Y]
	}
	if a.Size[Z] < min_size {
		axis = Z
	}
	return axis
}

// GetShortestAxisSize returns the scalar length of the shortest axis of the [AABB].
func (a AABB) GetShortestAxisSize() Float {
	var min_size = a.Size[x]
	if a.Size[Y] < min_size {
		min_size = a.Size[Y]
	}
	if a.Size[Z] < min_size {
		min_size = a.Size[Z]
	}
	return Float(min_size)
}

// GetSupport returns the vertex of the [AABB] that's the farthest in a given direction. This point is commonly
// known as the support point in collision detection algorithms.
func (a AABB) GetSupport(dir Vector3) Vector3 {
	var (
		half_extents = a.Size.Mulf(0.5)
		ofs          = a.Position.Add(half_extents)
	)
	return Vector3{
		ʕ(dir[X] > 0, half_extents[X], -half_extents[X]),
		ʕ(dir[Y] > 0, half_extents[Y], -half_extents[Y]),
		ʕ(dir[Z] > 0, half_extents[Z], -half_extents[Z]),
	}.Add(ofs)
}

// GetVolume returns the volume of the [AABB].
func (a AABB) GetVolume() Float { return Float(a.Size[X] * a.Size[Y] * a.Size[Z]) }

// Grow returns a copy of the AABB grown a given number of units towards all the sides.
func (a AABB) Grow(by Float) AABB {
	a.Position = a.Position.Subf(by)
	a.Size = a.Size.Addf(by * 2)
	return a
}

// HasPoint returns true if the [AABB] contains a point. Points on the faces of the [AABB] are considered included,
// though float-point precision errors may impact the accuracy of such checks.
//
// Note: This method is not reliable for [AABB] with a negative size. Use abs to get a positive sized equivalent
// [AABB] to check for contained points.
func (a AABB) HasPoint(point Vector3) bool {
	if point[X] < a.Position[X] {
		return false
	}
	if point[Y] < a.Position[Y] {
		return false
	}
	if point[Z] < a.Position[Z] {
		return false
	}
	if point[X] > a.Position[X]+a.Size[X] {
		return false
	}
	if point[Y] > a.Position[Y]+a.Size[Y] {
		return false
	}
	if point[Z] > a.Position[Z]+a.Size[Z] {
		return false
	}
	return true
}

// HasSurface returns true if the [AABB] has a surface or a length, and false if the [AABB] is empty (all components of
// size are zero or negative).
func (a AABB) HasSurface() bool { return a.Size[X] > 0 || a.Size[Y] > 0 || a.Size[Z] > 0 }

// HasVolume returns true if the [AABB] has a volume, and false if the [AABB] is flat, empty, or has a negative size.
func (a AABB) HasVolume() bool { return a.Size[X] > 0 && a.Size[Y] > 0 && a.Size[Z] > 0 }

// Intersection returns the intersection between two [AABB]. An empty [AABB] (size (0, 0, 0)) is returned on failure.
func (a AABB) Intersection(b AABB) AABB {
	var (
		src_min = a.Position
		src_max = a.Position.Add(a.Size)
		dst_min = b.Position
		dst_max = b.Position.Add(a.Size)
	)
	var (
		min, max Vector3
	)
	if src_min[X] > dst_max[X] || src_max[X] < dst_min[X] {
		return AABB{}
	} else {
		min[X] = ʕ(src_min[X] > dst_min[X], src_min[X], dst_min[X])
		max[X] = ʕ(src_max[X] < dst_max[X], src_max[X], dst_max[X])
	}
	if src_min[Y] > dst_max[Y] || src_max[Y] < dst_min[Y] {
		return AABB{}
	} else {
		min[Y] = ʕ(src_min[Y] > dst_min[Y], src_min[Y], dst_min[Y])
		max[Y] = ʕ(src_max[Y] < dst_max[Y], src_max[Y], dst_max[Y])
	}
	if src_min[Z] > dst_max[Z] || src_max[Z] < dst_min[Z] {
		return AABB{}
	} else {
		min[Z] = ʕ(src_min[Z] > dst_min[Z], src_min[Z], dst_min[Z])
		max[Z] = ʕ(src_max[Z] < dst_max[Z], src_max[Z], dst_max[Z])
	}
	return AABB{min, max.Sub(min)}
}

// Intersects returns true if the [AABB] overlaps with another.
func (a AABB) Intersects(b AABB) bool {
	if a.Position[X] >= (b.Position[X] + b.Size[X]) {
		return false
	}
	if (a.Position[X] + a.Size[X]) <= b.Position[X] {
		return false
	}
	if a.Position[Y] >= (b.Position[Y] + b.Size[Y]) {
		return false
	}
	if (a.Position[Y] + a.Size[Y]) <= b.Position[Y] {
		return false
	}
	if a.Position[Z] >= (b.Position[Z] + b.Size[Z]) {
		return false
	}
	if (a.Position[Z] + a.Size[Z]) <= b.Position[Z] {
		return false
	}
	return true
}

// IntersectsPlane returns true if the [AABB] is on both sides of a plane.
func (a AABB) IntersectsPlane(plane Plane) bool {
	var points = [8]Vector3{
		{a.Position[X], a.Position[Y], a.Position[Z]},
		{a.Position[X], a.Position[Y], a.Position[Z] + a.Size[Z]},
		{a.Position[X], a.Position[Y] + a.Size[Y], a.Position[Z]},
		{a.Position[X], a.Position[Y] + a.Size[Y], a.Position[Z] + a.Size[Z]},
		{a.Position[X] + a.Size[X], a.Position[Y], a.Position[Z]},
		{a.Position[X] + a.Size[X], a.Position[Y], a.Position[Z] + a.Size[Z]},
		{a.Position[X] + a.Size[X], a.Position[Y] + a.Size[Y], a.Position[Z]},
		{a.Position[X] + a.Size[X], a.Position[Y] + a.Size[Y], a.Position[Z] + a.Size[Z]},
	}
	var (
		over  = false
		under = false
	)
	for i := range points {
		if plane.DistanceTo(points[i]) > 0 {
			over = true
		} else {
			under = true
		}
	}
	return under && over
}

// IntersectsRay returns the point of intersection of the given ray with this [AABB] along with the normal
// or false if there is no intersection. Ray length is infinite.
func (a AABB) IntersectsRay(from, dir Vector3) (clip Vector3, normal Vector3, ok bool) {
	var (
		c1, c2 Vector3
		end          = a.Position.Add(a.Size)
		near   float = -1e20
		far    float = 1e20
		axis   Axis  = 0
	)
	for i := Axis(0); i < 3; i++ {
		if dir[i] == 0 {
			if (from[i] < a.Position[i]) || (from[i] > end[i]) {
				return Vector3{}, Vector3{}, false
			}
		} else { // ray not parallel to planes in this direction
			c1[i] = (a.Position[i] - from[i]) / dir[i]
			c2[i] = (end[i] - from[i]) / dir[i]
			if c1[i] > c2[i] {
				c1, c2 = c2, c1
			}
			if c1[i] > near {
				near = c1[i]
				axis = i
			}
			if c2[i] < far {
				far = c2[i]
			}
			if (near > far) || (far < 0) {
				return Vector3{}, Vector3{}, false
			}
		}
	}
	normal[axis] = ʕ[float](dir[axis] != 0, -1, 1)
	return c1, normal, true
}

// IntesectsSegment returns the point of intersection between from and to along with this AABB along with the normal
// or false if there is no intersection.
func (a AABB) IntesectsSegment(from, to Vector3) (clip Vector3, normal Vector3, ok bool) {
	var (
		min  float = 0
		max  float = 1
		axis Axis  = 0
		sign float = 0
	)
	for i := Axis(0); i < 3; i++ {
		var (
			seg_from   = from[i]
			seg_to     = to[i]
			box_begin  = a.Position[i]
			box_end    = box_begin + a.Size[i]
			cmin, cmax float
			csign      float
		)
		if seg_from < seg_to {
			if seg_from > box_end || seg_to < box_begin {
				return Vector3{}, Vector3{}, false
			}
			var length = seg_to - seg_from
			cmin = ʕ(seg_from < box_begin, ((box_begin - seg_from) / length), 0)
			cmax = ʕ(seg_to > box_end, ((box_end - seg_from) / length), 1)
			csign = -1.0
		} else {
			if seg_to > box_end || seg_from < box_begin {
				return Vector3{}, Vector3{}, false
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
			return Vector3{}, Vector3{}, false
		}
	}
	var (
		rel = to.Sub(from)
	)
	normal[axis] = sign
	return from.Add(rel.Mulf(Float(min))), normal, true
}

// IsAproxximatelyEqual returns true if this [AABB] and other are approximately equal, by running
// [IsApproximatelyEqual] on each component.
func (a AABB) IsAproxximatelyEqual(other AABB) bool {
	return a.Position.IsApproximatelyEqual(other.Position) && a.Size.IsApproximatelyEqual(other.Size)
}

// IsFinite returns true if this [AABB] is finite, by calling [IsFinite] on each component.
func (a AABB) IsFinite() bool { return a.Position.IsFinite() && a.Size.IsFinite() }

// Merge returns the smallest AABB enclosing both this AABB and b.
func (a AABB) Merge(b AABB) AABB {
	var beg_1, beg_2 Vector3
	var end_1, end_2 Vector3
	var min, max Vector3
	beg_1 = a.Position
	beg_2 = b.Position
	end_1 = a.Size.Add(beg_1)
	end_2 = b.Size.Add(beg_2)
	min[X] = ʕ(beg_1[X] < beg_2[X], beg_1[X], beg_2[X])
	min[Y] = ʕ(beg_1[Y] < beg_2[Y], beg_1[Y], beg_2[Y])
	min[Z] = ʕ(beg_1[Z] < beg_2[Z], beg_1[Z], beg_2[Z])
	max[X] = ʕ(end_1[X] > end_2[X], end_1[X], end_2[X])
	max[Y] = ʕ(end_1[Y] > end_2[Y], end_1[Y], end_2[Y])
	max[Z] = ʕ(end_1[Z] > end_2[Z], end_1[Z], end_2[Z])
	return AABB{
		Position: min,
		Size:     max.Sub(min),
	}
}

// Projection creates a new Projection that scales a given projection to fit around a given
// [AABB] in projection space.
func (a AABB) Projection() Projection {
	min := a.Position
	max := a.Position.Add(a.Size)
	return Projection{
		Vector4{2 / (max[X] - min[X]), 0, 0, -(max[X] + min[X]) / (max[X] - min[X])},
		Vector4{0, 2 / (max[Y] - min[Y]), 0, -(max[Y] + min[Y]) / (max[Y] - min[Y])},
		Vector4{0, 0, 2 / (max[Z] - min[Z]), -(max[Z] + min[Z]) / (max[Z] - min[Z])},
		Vector4{0, 0, 0, 1},
	}
}

func (a AABB) Transform(t Transform3D) AABB {
	/* https://dev.theomader.com/transform-bounding-boxes/ */
	var min = a.Position
	var max = a.Position.Add(a.Size)
	var tmin, tmax Vector3
	for i := 0; i < 3; i++ {
		tmin[i] = t.Origin[i]
		tmax[i] = t.Origin[i]
		for j := 0; j < 3; j++ {
			var e = t.Basis[i][j] * min[j]
			var f = t.Basis[i][j] * max[j]
			if e < f {
				tmin[i] += e
				tmax[i] += f
			} else {
				tmin[i] += f
				tmax[i] += e
			}
		}
	}
	return AABB{
		Position: tmin,
		Size:     tmax.Sub(tmin),
	}
}
