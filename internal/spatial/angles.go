package spatial

import "math"

// Radians represents an angle in radians.
type Radians Float

// Radians converts an angle expressed in radians to degrees.
func (rad Radians) Degrees() Degrees { return Degrees(rad * (180.0 / Pi)) }

// Vector2 creates a unit Vector2 rotated to the given angle in radians. This is equivalent to
// doing NewVector(Cos(angle),Sin(angle)) or Const(Vector2.Right).Rotated(angle).
func (angle Radians) Vector2() Vector2 {
	return Vector2{float(math.Cos(float64(angle))), float(math.Sin(float64(angle)))}
}

// Degrees represents an angle in degrees.
type Degrees Float

func (deg Degrees) Radians() Radians { return Radians(deg * (Pi / 180.0)) }

// Vector2 creates a unit Vector2 rotated to the given angle in degrees. This is equivalent to
// doing angle.Radians().Vector2()
func (deg Degrees) Vector2() Vector2 {
	angle := deg.Radians()
	return Vector2{float(math.Cos(float64(angle))), float(math.Sin(float64(angle)))}
}

// EulerAngles represents a rotation in 3D space using Euler angles.
type EulerAngles [3]Radians

type EulerOrder int64

const (
	EulerOrderXYZ EulerOrder = 0
	EulerOrderXZY EulerOrder = 1
	EulerOrderYXZ EulerOrder = 2
	EulerOrderYZX EulerOrder = 3
	EulerOrderZXY EulerOrder = 4
	EulerOrderZYX EulerOrder = 5
)

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
		float(sin_a1*cos_a2*sin_a3 + cos_a1*sin_a2*cos_a3),
		float(sin_a1*cos_a2*cos_a3 - cos_a1*sin_a2*sin_a3),
		float(-sin_a1*sin_a2*cos_a3 + cos_a1*cos_a2*sin_a3),
		float(sin_a1*sin_a2*sin_a3 + cos_a1*cos_a2*cos_a3),
	}
}

// Basis constructs a pure rotation Basis matrix from Euler angles in the specified Euler rotation order.
// By default, use YXZ order (most common). See the EulerOrder enum for possible values.
func (e EulerAngles) Basis(order EulerOrder) Basis {
	var c, s float

	c = float(Cos(e[X]))
	s = float(Sin(e[X]))
	var xmat = Basis{{1, 0, 0}, {0, c, -s}, {0, s, c}}

	c = float(Cos(e[Y]))
	s = float(Sin(e[Y]))
	var ymat = Basis{{c, 0, s}, {0, 1, 0}, {-s, 0, c}}

	c = float(Cos(e[Z]))
	s = float(Sin(e[Z]))
	var zmat = Basis{{c, -s, 0}, {s, c, 0}, {0, 0, 1}}

	switch order {
	case EulerOrderXYZ:
		return xmat.Mul(ymat.Mul(zmat))
	case EulerOrderXZY:
		return xmat.Mul(zmat.Mul(ymat))
	case EulerOrderYXZ:
		return ymat.Mul(xmat.Mul(zmat))
	case EulerOrderYZX:
		return ymat.Mul(zmat.Mul(xmat))
	case EulerOrderZXY:
		return zmat.Mul(xmat.Mul(ymat))
	case EulerOrderZYX:
		return zmat.Mul(ymat.Mul(xmat))
	default:
		panic("Invalid order parameter for EulerAngles.Basis()")
	}
}
