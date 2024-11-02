//go:build !generate

package gd

import (
	float "grow.graphics/gd/variant/Float"
)

// NewRect2 constructs a Rect2 by setting its position to (x, y), and its size to (width, height).
func NewRect2(x, y, width, height Float) Rect2 {
	return Rect2{
		Position: Vector2{float.X(x), float.X(y)},
		Size:     Vector2{float.X(width), float.X(height)},
	}
}

func NewRect2i(x, y, width, height int32) Rect2i {
	return Rect2i{
		Position: Vector2i{x, y},
		Size:     Vector2i{width, height},
	}
}

func NewTransform2D(a, b, c, d, e, f Float) Transform2D {
	return Transform2D{
		Vector2{float.X(a), float.X(b)},
		Vector2{float.X(c), float.X(d)},
		Vector2{float.X(e), float.X(f)},
	}
}

func NewTransform3D(a, b, c, d, e, f, g, h, i, j, k, l Float) Transform3D {
	return Transform3D{
		Basis: Basis{
			Vector3{float.X(a), float.X(b), float.X(c)},
			Vector3{float.X(d), float.X(e), float.X(f)},
			Vector3{float.X(g), float.X(h), float.X(i)},
		},
		Origin: Vector3{float.X(j), float.X(k), float.X(l)},
	}
}
