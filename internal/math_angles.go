//go:build !generate

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

// EulerAngles represents a rotation in 3D space using Euler angles.
type EulerAngles [3]Radians

// Quaternion constructs a Quaternion from Euler angles in YXZ rotation order.
func (e EulerAngles) Quaternion() Quaternion {
	var (
		half_a1 = e[Y] * 0.5
		half_a2 = e[X] * 0.5
		half_a3 = e[Z] * 0.5
	)
	// R = Y(a1).X(a2).Z(a3) convention for Euler angles.
	// Conversion to quaternion as listed in https://ntrs.nasa.gov/archive/nasa/casi.ntrs.nasa.gov/19770024290.pdf (page A-6)
	// a3 is the angle of the first rotation, following the notation in this reference.
	var (
		cos_a1 = Cos(half_a1)
		sin_a1 = Sin(half_a1)
		cos_a2 = Cos(half_a2)
		sin_a2 = Sin(half_a2)
		cos_a3 = Cos(half_a3)
		sin_a3 = Sin(half_a3)
	)
	return Quaternion{
		float32(sin_a1*cos_a2*sin_a3 + cos_a1*sin_a2*cos_a3),
		float32(sin_a1*cos_a2*cos_a3 - cos_a1*sin_a2*sin_a3),
		float32(-sin_a1*sin_a2*cos_a3 + cos_a1*cos_a2*sin_a3),
		float32(sin_a1*sin_a2*sin_a3 + cos_a1*cos_a2*cos_a3),
	}
}
