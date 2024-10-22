//go:build !generate

package gd

import "grow.graphics/xy"

// NewRect2 constructs a Rect2 by setting its position to (x, y), and its size to (width, height).
func NewRect2(x, y, width, height Float) Rect2 {
	return xy.NewRect2(x, y, width, height)
}

func NewRect2i(x, y, width, height int32) Rect2i {
	return xy.NewRect2i(x, y, width, height)
}

func NewTransform2D(a, b, c, d, e, f Float) Transform2D {
	return xy.Transform2D{
		xy.NewVector2(a, b),
		xy.NewVector2(c, d),
		xy.NewVector2(e, f),
	}
}

func NewTransform3D(a, b, c, d, e, f, g, h, i, j, k, l Float) Transform3D {
	return xy.Transform3D{
		Basis: xy.Basis{
			xy.NewVector3(a, b, c),
			xy.NewVector3(d, e, f),
			xy.NewVector3(g, h, i),
		},
		Origin: xy.NewVector3(j, k, l),
	}
}
