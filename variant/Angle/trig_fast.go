//go:build !precision_double && (amd64 || arm64) && !O0

package Angle

import "graphics.gd/variant/Float"

// Cos returns the cosine of angle x in radians.
func Cos(x Radians) Float.X { return cos32(float32(x)) }

// Sin returns the sine of angle x in radians.
func Sin(x Radians) Float.X { return sin32(float32(x)) }

// AsVector2 creates a unit Vector2 rotated to the given angle in radians. This is equivalent
// to doing:
//
//	Vector2.New(math.Cos(angle), math.Sin(angle))
//	Vector2.Rotated(Vector2.Right, angle).
func (angle Radians) AsVector2() vector2 {
	sin, cos := sincos32(float32(angle))
	return vector2{float(cos), float(sin)}
}
