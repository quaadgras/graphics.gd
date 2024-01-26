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

// Abs returns an AABB with equivalent position and size, modified so that the most-negative corner is the origin and the
// size is positive.
func (a AABB) Abs() AABB {
	return AABB{
		Position: Vector3{a.Position[X] + min(a.Size[X], 0), a.Position[Y] + min(a.Size[Y], 0), a.Position[Z] + min(a.Size[Z], 0)},
		Size:     a.Size.Abs(),
	}
}
