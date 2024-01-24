package gd

import "math"

// Radians represents an angle in radians.
type Radians Float

// Radians converts an angle expressed in radians to degrees.
func (rad Radians) Degrees() Degrees { return Degrees(rad * (180.0 / Pi)) }

// Vector2 creates a unit Vector2 rotated to the given angle in radians. This is equivalent to
// doing NewVector(Cos(angle),Sin(angle)) or Const(Vector2.Right).Rotated(angle).
func (angle Radians) Vector2() Vector2 {
	return Vector2{float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))}
}

// Degrees represents an angle in degrees.
type Degrees Float

func (deg Degrees) Radians() Radians { return Radians(deg * (Pi / 180.0)) }

// Vector2 creates a unit Vector2 rotated to the given angle in degrees. This is equivalent to
// doing angle.Radians().Vector2()
func (deg Degrees) Vector2() Vector2 {
	angle := deg.Radians()
	return Vector2{float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))}
}
