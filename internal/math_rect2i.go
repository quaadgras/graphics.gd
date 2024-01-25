//go:build !generate

package gd

/*
Rect2 represents an axis-aligned rectangle in a 2D space, using integer coordinates. It is defined by
its position and size, which are Vector2i. Because it does not rotate, it is frequently used for fast
overlap tests (see intersects).

For floating-point coordinates, see [Rect2].

Note: Negative values for size are not supported. With negative size, most [Rect2i] methods do not work
correctly. Use abs to get an equivalent [Rect2i] with a non-negative size.

Note: In a boolean context, a [Rect2i] evaluates to false if both position and size are zero (equal to
Const(Vector2i.ZERO)). Otherwise, it always evaluates to true.
*/
type Rect2i struct {
	Position Vector2i
	Size     Vector2i
}

// End is the ending point. This is usually the bottom-right corner of the rectangle, and is
// equivalent to position + size.
func (r Rect2i) End() Vector2i { return r.Position.Add(r.Size) }

// SetEnd sets the ending point. This is usually the bottom-right corner of the rectangle, and is
// equivalent to position + size. Setting the end point will change the size of the rectangle.
func (r *Rect2i) SetEnd(end Vector2i) { r.Size = end.Sub(r.Position) }

// Abs returns a Rect2 equivalent to this rectangle, with its width and height modified to be
// non-negative values, and with its position being the top-left corner of the rectangle.
//
//	var rect = NewRect2(25,25,-100,-50)
//	var absolute = rect.Abs() // absolute is Rect2(-75, -25, 100, 50)
//
// Note: It's recommended to use this method when size is negative, as most other methods in
// Godot assume that the position is the top-left corner, and the end is the bottom-right corner.
func (r Rect2i) Abs() Rect2i {
	r.Position = r.Position.Abs()
	r.Size = r.Size.Abs()
	return r
}

// Encloses returns true if this rectangle completely encloses the b rectangle.
func (r Rect2i) Encloses(b Rect2i) bool {
	return (b.Position[X] >= r.Position[Y]) && (b.Position[Y] >= r.Position[Y]) &&
		((b.Position[X] + b.Size[X]) <= (r.Position[X] + r.Size[X])) &&
		((b.Position[Y] + b.Size[Y]) <= (r.Position[Y] + r.Size[Y]))
}

// Expand returns a copy of this rectangle expanded to align the edges with the given to point, if
// necessary.
//
//	var rect = NewRect2(0, 0, 5, 2)
//	rect = rect.Expand(Vector2{10, 0}) // rect is now Rect2(0, 0, 10, 2)
//	rect = rect.Expand(Vector2{-5, 5}) // rect is now Rect2(-5, 0, 10, 5)
func (r Rect2i) Expand(to Vector2i) Rect2i {
	var begin = r.Position
	var end = Vector2i.Add(r.Position, r.Size)
	if to[X] < begin[X] {
		begin[X] = to[X]
	}
	if to[Y] < begin[Y] {
		begin[Y] = to[Y]
	}
	if to[X] > end[X] {
		end[X] = to[X]
	}
	if to[Y] > end[Y] {
		end[Y] = to[Y]
	}
	return Rect2i{
		Position: begin,
		Size:     end.Sub(begin),
	}
}

// GetArea returns the rectangle's area. This is equivalent to Size[X] * Size[Y]. See also [Rect2.HasArea].
func (r Rect2i) GetArea() Float { return Float(r.Size[X] * r.Size[Y]) }

// GetCenter returns the center point of the rectangle. This is the same as position + (size / 2.0).
func (r Rect2i) GetCenter() Vector2i { return r.Position.Add(r.Size.Divi(2)) }

// Grow returns a copy of this rectangle extended on all sides by the given amount. A negative amount
// shrinks the rectangle instead. See also [Rect2.GrowIndividual] and [Rect2.GrowSide].
//
//	var a = Rect2{4,4,8,8}.Grow(4) // a is Rect2(0, 0, 16, 16)
//	var b = Rect2{4,4,8,8}.Grow(2) // b is Rect2(-2, -2, 12, 8)
func (r Rect2i) Grow(amount Int) Rect2i {
	r.Position = r.Position.Subi(amount)
	r.Size = r.Size.Addi(amount * 2)
	return r
}

// GrowIndividual returns a copy of this rectangle with its left, top, right, and bottom sides extended by
// the given amounts. Negative values shrink the sides, instead. See also [Rect2.Grow] and [Rect2.GrowSide].
func (r Rect2i) GrowIndividual(left, top, right, bottom Int) Rect2i {
	r.Position[X] -= int32(left)
	r.Position[Y] -= int32(top)
	r.Size[X] += int32(left + right)
	r.Size[Y] += int32(top + bottom)
	return r
}

// GrowSide returns a copy of this rectangle with its side extended by the given amount (see [Side] constants).
// A negative amount shrinks the rectangle, instead. See also [Rect2.Grow] and [Rect2.GrowIndividual].
func (r Rect2i) GrowSide(side Side, amount Int) Rect2i {
	switch side {
	case SideLeft:
		r.Position[X] -= int32(amount)
		r.Size[X] += int32(amount)
	case SideTop:
		r.Position[Y] -= int32(amount)
		r.Size[Y] += int32(amount)
	case SideRight:
		r.Size[X] += int32(amount)
	case SideBottom:
		r.Size[Y] += int32(amount)
	}
	return r
}

// HasArea returns true if this rectangle has positive width and height. See also [GetArea].
func (r Rect2i) HasArea() bool { return r.Size[X] > 0 && r.Size[Y] > 0 }

// HasPoint returns true if the rectangle contains the given point. By convention, points on the right
// and bottom edges are not included.
//
// Note: This method is not reliable for [Rect2] with a negative size. Use abs first to get a valid rectangle.
func (r Rect2i) HasPoint(point Vector2i) bool {
	if point[X] < r.Position[X] {
		return false
	}
	if point[Y] < r.Position[Y] {
		return false
	}
	if point[X] >= r.Position[X]+r.Size[X] {
		return false
	}
	if point[Y] >= r.Position[Y]+r.Size[Y] {
		return false
	}
	return true
}

// Intersection returns the intersection between this rectangle and b. If the rectangles do not intersect,
// returns an empty [Rect2].
//
//	var rect1 = Rect2{0, 0, 5, 10}
//	var rect2 = Rect2{2, 0, 9, 4}
//
//	var a = rect1.Intersection(rect2) // a is Rect2(2, 0, 3, 4)
//
// Note: If you only need to know whether two rectangles are overlapping, use [Rect2.Intersects], instead.
func (r Rect2i) Intersection(b Rect2i) Rect2i {
	var new_rect = b
	if !r.Intersects(new_rect, false) {
		return Rect2i{}
	}
	new_rect.Position[X] = max(b.Position[X], r.Position[X])
	new_rect.Position[Y] = max(b.Position[Y], r.Position[Y])
	var (
		rect_end = b.Position.Add(b.Size)
		end      = r.Position.Add(r.Size)
	)
	new_rect.Size[X] = min(rect_end[X], end[X]) - new_rect.Position[X]
	new_rect.Size[Y] = min(rect_end[Y], end[Y]) - new_rect.Position[Y]
	return new_rect
}

// Intersects returns true if this rectangle overlaps with the b rectangle. The edges of both rectangles
// are excluded, unless include_borders is true.
func (r Rect2i) Intersects(b Rect2i, include_borders bool) bool {
	if include_borders {
		if r.Position[X] > (b.Position[X] + b.Size[X]) {
			return false
		}
		if (r.Position[X] + r.Size[X]) < b.Position[X] {
			return false
		}
		if r.Position[Y] > (b.Position[Y] + b.Size[Y]) {
			return false
		}
		if (r.Position[Y] + b.Size[Y]) < b.Position[Y] {
			return false
		}
	} else {
		if r.Position[X] >= (b.Position[X] + b.Size[X]) {
			return false
		}
		if (r.Position[X] + r.Size[X]) <= b.Position[X] {
			return false
		}
		if r.Position[Y] >= (b.Position[Y] + b.Size[Y]) {
			return false
		}
		if (r.Position[Y] + r.Size[Y]) <= b.Position[Y] {
			return false
		}
	}
	return true
}

// Merge returns a [Rect2] that encloses both this rectangle and b around the edges. See also [Rect2.Encloses].
func (r Rect2i) Merge(b Rect2i) Rect2i {
	var new_rect Rect2i
	new_rect.Position[X] = min(b.Position[X], r.Position[X])
	new_rect.Position[Y] = min(b.Position[Y], r.Position[Y])
	new_rect.Size[X] = max(b.Position[X]+b.Size[X], r.Position[X]+r.Size[X])
	new_rect.Size[Y] = max(b.Position[Y]+b.Size[Y], r.Position[Y]+r.Size[Y])
	new_rect.Size = new_rect.Size.Sub(new_rect.Position) // Make relative again.
	return new_rect
}
