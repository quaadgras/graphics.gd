//go:build precision_double || (!amd64 && !arm64)

package Angle

import (
	"math"

	"graphics.gd/variant/Float"
)

// Cos returns the cosine of angle x in radians.
func Cos(x Radians) Float.X { return Float.X(math.Cos(float64(x))) } //gd:cos

// Sin returns the sine of angle x in radians.
func Sin(x Radians) Float.X { return Float.X(math.Sin(float64(x))) } //gd:sin

// AsVector2 creates a unit Vector2 rotated to the given angle in radians. This is equivalent
// to doing:
//
//	Vector2.New(math.Cos(angle), math.Sin(angle))
//	Vector2.Rotated(Vector2.Right, angle).
func (angle Radians) AsVector2() vector2 { //gd:Vector2.from_angle
	return vector2{float(math.Cos(float64(angle))), float(math.Sin(float64(angle)))}
}
