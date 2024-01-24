package gd

import "math"

// Radians represents an angle in radians.
type Radians Float

// Vector2 creates a unit Vector2 rotated to the given angle in radians. This is equivalent to
// doing NewVector(Cos(angle),Sin(angle)) or Const(Vector2.Right).Rotated(angle).
func (angle Radians) Vector2() Vector2 {
	return Vector2{float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))}
}
