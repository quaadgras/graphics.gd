// Package Basis is a 3×3 matrix for representing 3D rotation and scale.
package Basis

import (
	"unsafe"

	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector3"
)

// The XYZ type is a 3×3 matrix used to represent 3D rotation, scale, and
// shear. It is frequently used within a Transform3D.
//
// A Basis is composed by 3 axis vectors, each representing a column of the
// matrix: x, y, and z. The length of each axis (Vector3.Length) influences
// the basis's scale, while the direction of all axes influence the rotation.
// Usually, these axes are perpendicular to one another. However, when you
// rotate any axis individually, the basis becomes sheared. Applying a
// sheared basis to a 3D model will make the model appear distorted.
//
// A Basis is orthogonal if its axes are perpendicular to each other. A basis
// is normalized if the length of every axis is 1. A basis is uniform if all
// axes share the same length (see get_scale). A basis is orthonormal if it
// is both orthogonal and normalized, which allows it to only represent
// rotations. A basis is conformal if it is both orthogonal and uniform,
// which ensures it is not distorted.
//
// For a general introduction, see the Matrices and transforms tutorial.
//
// Note: GD uses a right-handed coordinate system, which is a common standard.
// For directions, the convention for built-in types like Camera3D is for
// -Z to point forward (+X is right, +Y is up, and +Z is back). Other
// objects may use different direction conventions. For more information,
// see the 3D asset direction conventions tutorial.
type XYZ = struct {

	// The basis's X axis, and the column 0 of the matrix.
	//
	// On the identity basis, this vector points right ([Vector3.Right]).
	X Vector3.XYZ

	// The basis's Y axis, and the column 1 of the matrix.
	//
	// On the identity basis, this vector points up ([Vector3.Up]).
	Y Vector3.XYZ

	// The basis's Z axis, and the column 2 of the matrix.
	//
	// On the identity basis, this vector points back ([Vector3.Back]).
	Z Vector3.XYZ
}

var (
	// Identity basis. This is a basis with no rotation, no shear, and its
	// scale being 1. This means that:
	//
	//   - The x points right ([Vector3.Right]);
	//   - The y points up ([Vector3.Up]);
	//   - The z points back ([Vector3.Back]).
	Identity = XYZ{
		X: Vector3.Right,
		Y: Vector3.Up,
		Z: Vector3.Back,
	}

	// FlipX when any basis is multiplied by FlipX, it negates all components
	// of the x axis (the X column).
	//
	// When FlipX is multiplied by any basis, it negates the Vector3.X component
	// of all axes (the X row).
	FlipX = XYZ{
		X: Vector3.Left,
		Y: Vector3.Up,
		Z: Vector3.Back,
	}

	// FlipY when any basis is multiplied by FlipY, it negates all components
	// of the y axis (the Y column).
	//
	// When FlipY is multiplied by any basis, it negates the Vector3.Y component
	// of all axes (the Y row).
	FlipY = XYZ{
		X: Vector3.Right,
		Y: Vector3.Down,
		Z: Vector3.Back,
	}

	// FlipZ when any basis is multiplied by FlipZ, it negates all components
	// of the z axis (the Z column).
	//
	// When FlipZ is multiplied by any basis, it negates the Vector3.Z component
	// of all axes (the Z row).
	FlipZ = XYZ{
		X: Vector3.Right,
		Y: Vector3.Up,
		Z: Vector3.Forward,
	}
)

// New constructs a Basis identical to the [Identity].
func New() XYZ {
	return Identity
}

func cofac(basis XYZ, row1, col1, row2, col2 int) Float.X {
	b := *(*[3][3]Float.X)(unsafe.Pointer(&basis))
	return (b[row1][col1]*b[row2][col2] - b[row1][col2]*b[row2][col1])
}

// Scales constructs a pure scale basis matrix with no rotation or shearing.
// The scale values are set as the diagonal of the matrix, and the other
// parts of the matrix are zero.
func Scales(scale Vector3.XYZ) XYZ { //gd:Basis.from_scale
	return XYZ{
		Vector3.New(scale.X, 0, 0),
		Vector3.New(0, scale.Y, 0),
		Vector3.New(0, 0, scale.Z),
	}
}

// Outer returns the outer product with with.
func Outer(v, with Vector3.XYZ) XYZ { //gd:Vector3.outer
	return XYZ{
		Vector3.XYZ{v.X * with.X, v.X * with.Y, v.X * with.Z},
		Vector3.XYZ{v.Y * with.X, v.Y * with.Y, v.Y * with.Z},
		Vector3.XYZ{v.Z * with.X, v.Z * with.Y, v.Z * with.Z},
	}
}

// Euler constructs a pure rotation Basis matrix from Euler angles in the specified Euler rotation order.
// By default, use YXZ order (most common). See the EulerOrder enum for possible values.
func Euler(e Angle.Euler3D, order Angle.Order) XYZ { //gd:Basis.from_euler
	var (
		c, s Float.X
	)
	c = Angle.Cos(e.X)
	s = Angle.Sin(e.X)
	var xmat = XYZ{
		Vector3.New(1, 0, 0),
		Vector3.New(0, c, -s),
		Vector3.New(0, s, c),
	}
	c = Angle.Cos(e.Y)
	s = Angle.Sin(e.Y)
	var (
		ymat = XYZ{
			Vector3.New(c, 0, s),
			Vector3.New(0, 1, 0),
			Vector3.New(-s, 0, c),
		}
	)
	c = Angle.Cos(e.Z)
	s = Angle.Sin(e.Z)
	var (
		zmat = XYZ{
			Vector3.New(c, -s, 0),
			Vector3.New(s, c, 0),
			Vector3.New(0, 0, 1),
		}
	)
	switch order {
	case Angle.OrderXYZ:
		return Mul(xmat, Mul(ymat, zmat))
	case Angle.OrderXZY:
		return Mul(xmat, Mul(zmat, ymat))
	case Angle.OrderYXZ:
		return Mul(ymat, Mul(xmat, zmat))
	case Angle.OrderYZX:
		return Mul(ymat, Mul(zmat, xmat))
	case Angle.OrderZXY:
		return Mul(zmat, Mul(xmat, ymat))
	case Angle.OrderZYX:
		return Mul(zmat, Mul(ymat, xmat))
	default:
		panic("Invalid order parameter for EulerAngles.Basis()")
	}
}

// NewBasisRotatedAround constructs a pure rotation basis matrix, rotated around the given axis by angle (in radians).
// The axis must be a normalized vector.
func RotatesAxisAngle(axis Vector3.XYZ, angle Angle.Radians) XYZ { //gd:Basis(Vector3,float)
	var rows XYZ
	var c = Angle.Cos(angle)
	var s = Angle.Sin(angle)
	var t = 1 - c
	var ux, uy, uz = axis.X, axis.Y, axis.Z
	rows.X.X = c + ux*ux*t
	rows.X.Y = uy*ux*t + uz*s
	rows.X.Z = uz*ux*t - uy*s
	rows.Y.X = ux*uy*t - uz*s
	rows.Y.Y = c + uy*uy*t
	rows.Y.Z = uz*uy*t + ux*s
	rows.Z.X = ux*uz*t + uy*s
	rows.Z.Y = uy*uz*t - ux*s
	rows.Z.Z = c + uz*uz*t
	return rows
}

// LookingAt creates a Basis with a rotation such that the forward axis (-Z) points towards the target position.
//
// The up axis (+Y) points as close to the up vector as possible while staying perpendicular to the forward axis.
// The resulting Basis is orthonormalized. The target and up vectors cannot be zero, and cannot be parallel to each other.
//
// If use_model_front is true, the +Z axis (asset front) is treated as forward (implies +X is left) and points toward the
// target position. By default, the -Z axis (camera forward) is treated as forward (implies +X is right).
func LookingAt(target, up Vector3.XYZ) XYZ { //gd:Basis.looking_at
	vZ := Vector3.Normalized(target)                // Z column points to target
	vX := Vector3.Normalized(Vector3.Cross(up, vZ)) // X column perpendicular
	vY := Vector3.Cross(vZ, vX)                     // Y column completes the basis
	return XYZ{X: vX, Y: vY, Z: vZ}
}

// RotatesScales returns the basis from the given rotation Quaternion with the
// specified scale applied to it. This is equivalent to multiplying the basis
// with a scaling matrix.
func RotatesScales(q quaternion, s Vector3.XYZ) XYZ {
	return Mul(qAsBasis(q), Scales(s))
}

// Determinant returns the determinant of the basis matrix. If the basis is
// uniformly scaled, its determinant is the square of the scale.
//
// A negative determinant means the basis has a negative scale. A zero
// determinant means the basis isn't invertible, and is usually considered invalid.
func Determinant(m XYZ) Float.X { //gd:Basis.determinant
	return m.X.X*(m.Y.Y*m.Z.Z-m.Z.Y*m.Y.Z) -
		m.Y.X*(m.X.Y*m.Z.Z-m.Z.Y*m.X.Z) +
		m.Z.X*(m.X.Y*m.Y.Z-m.Y.Y*m.X.Z)
}

// AsEulerAngles returns the basis's rotation in the form of Euler angles. The Euler order depends
// on the order parameter, by default it uses the YXZ convention: when decomposing, first Z,
// then X, and Y last. The returned vector contains the rotation angles in the format
// (X angle, Y angle, Z angle).
//
// Consider using the [Basis.Quaternion] method instead, which returns a [Quaternion]
// quaternion instead of [EulerAngles].
func AsEulerAngles(b XYZ, order Angle.Order) Angle.Euler3D { //gd:Basis.get_euler
	switch order {
	case Angle.OrderXYZ:
		// Euler angles in XYZ convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cy*cz          -cy*sz           sy
		//        cz*sx*sy+cx*sz  cx*cz-sx*sy*sz -cy*sx
		//       -cx*cz*sy+sx*sz  cz*sx+cx*sy*sz  cx*cy
		var euler Angle.Euler3D
		var sy = b.X.Z
		if sy < (1.0 - Float.Epsilon) {
			if sy > -(1.0 - Float.Epsilon) {
				// is this a pure Y rotation?
				if b.Y.X == 0 && b.X.Y == 0 && b.Y.Z == 0 && b.Z.Y == 0 && b.X.X == 1 {
					// return the simplest form (human friendlier in editor and scripts)
					euler.X = 0
					euler.Y = Angle.Atan2(b.X.Z, b.X.X)
					euler.Z = 0
				} else {
					euler.X = Angle.Atan2(-b.Y.Z, b.Z.Z)
					euler.Y = Angle.Asin(sy)
					euler.Z = Angle.Atan2(-b.X.Y, b.X.X)
				}
			} else {
				euler.X = Angle.Atan2(b.Z.Y, b.Y.Y)
				euler.Y = -Angle.Pi / 2.0
				euler.Z = 0.0
			}
		} else {
			euler.X = Angle.Atan2(b.Z.Y, b.Y.Y)
			euler.Y = Angle.Pi / 2.0
			euler.Z = 0.0
		}
		return euler
	case Angle.OrderXZY:
		// Euler angles in XZY convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cz*cy             -sz             cz*sy
		//        sx*sy+cx*cy*sz    cx*cz           cx*sz*sy-cy*sx
		//        cy*sx*sz          cz*sx           cx*cy+sx*sz*sy

		var euler Angle.Euler3D
		var sz = b.X.Y
		if sz < (1.0 - Float.Epsilon) {
			if sz > -(1.0 - Float.Epsilon) {
				euler.X = Angle.Atan2(b.Z.Y, b.Y.Y)
				euler.Y = Angle.Atan2(b.X.Z, b.X.X)
				euler.Z = Angle.Asin(-sz)
			} else {
				// It's -1
				euler.X = -Angle.Atan2(b.Y.Z, b.Z.Z)
				euler.Y = 0.0
				euler.Z = Angle.Pi / 2.0
			}
		} else {
			// It's 1
			euler.X = -Angle.Atan2(b.Y.Z, b.Z.Z)
			euler.Y = 0.0
			euler.Z = -Angle.Pi / 2.0
		}
		return euler
	case Angle.OrderYXZ:
		// Euler angles in YXZ convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cy*cz+sy*sx*sz    cz*sy*sx-cy*sz        cx*sy
		//        cx*sz             cx*cz                 -sx
		//        cy*sx*sz-cz*sy    cy*cz*sx+sy*sz        cy*cx

		var euler Angle.Euler3D

		var m12 = b.Y.Z

		if m12 < (1 - Float.Epsilon) {
			if m12 > -(1 - Float.Epsilon) {
				// is this a pure X rotation?
				if b.Y.X == 0 && b.X.Y == 0 && b.X.Z == 0 && b.Z.X == 0 && b.X.X == 1 {
					// return the simplest form (human friendlier in editor and scripts)
					euler.X = Angle.Atan2(-m12, b.Y.Y)
					euler.Y = 0
					euler.Z = 0
				} else {
					euler.X = Angle.Asin(-m12)
					euler.Y = Angle.Atan2(b.X.Z, b.Z.Z)
					euler.Z = Angle.Atan2(b.Y.X, b.Y.Y)
				}
			} else { // m12 == -1
				euler.X = Angle.Pi * 0.5
				euler.Y = Angle.Atan2(b.X.Y, b.X.X)
				euler.Z = 0
			}
		} else { // m12 == 1
			euler.X = -Angle.Pi * 0.5
			euler.Y = -Angle.Atan2(b.X.Y, b.X.X)
			euler.Z = 0
		}

		return euler
	case Angle.OrderYZX:
		// Euler angles in YZX convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cy*cz             sy*sx-cy*cx*sz     cx*sy+cy*sz*sx
		//        sz                cz*cx              -cz*sx
		//        -cz*sy            cy*sx+cx*sy*sz     cy*cx-sy*sz*sx

		var euler Angle.Euler3D
		var sz = b.Y.X
		if sz < (1.0 - Float.Epsilon) {
			if sz > -(1.0 - Float.Epsilon) {
				euler.X = Angle.Atan2(-b.Y.Z, b.Y.Y)
				euler.Y = Angle.Atan2(-b.Z.X, b.X.X)
				euler.Z = Angle.Asin(sz)
			} else {
				// It's -1
				euler.X = Angle.Atan2(b.Z.Y, b.Z.Z)
				euler.Y = 0.0
				euler.Z = -Angle.Pi / 2.0
			}
		} else {
			// It's 1
			euler.X = Angle.Atan2(b.Z.Y, b.Z.Z)
			euler.Y = 0.0
			euler.Z = Angle.Pi / 2.0
		}
		return euler
	case Angle.OrderZXY:
		// Euler angles in ZXY convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cz*cy-sz*sx*sy    -cx*sz                cz*sy+cy*sz*sx
		//        cy*sz+cz*sx*sy    cz*cx                 sz*sy-cz*cy*sx
		//        -cx*sy            sx                    cx*cy
		var euler Angle.Euler3D
		var sx = b.Z.Y
		if sx < (1.0 - Float.Epsilon) {
			if sx > -(1.0 - Float.Epsilon) {
				euler.X = Angle.Asin(sx)
				euler.Y = Angle.Atan2(-b.Z.X, b.Z.Z)
				euler.Z = Angle.Atan2(-b.X.Y, b.Y.Y)
			} else {
				// It's -1
				euler.X = -Angle.Pi / 2.0
				euler.Y = Angle.Atan2(b.X.Z, b.X.X)
				euler.Z = 0
			}
		} else {
			// It's 1
			euler.X = Angle.Pi / 2.0
			euler.Y = Angle.Atan2(b.X.Z, b.X.X)
			euler.Z = 0
		}
		return euler
	case Angle.OrderZYX:
		// Euler angles in ZYX convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cz*cy             cz*sy*sx-cx*sz        sz*sx+cz*cx*cy
		//        cy*sz             cz*cx+sz*sy*sx        cx*sz*sy-cz*sx
		//        -sy               cy*sx                 cy*cx
		var euler Angle.Euler3D
		var sy = b.Z.X
		if sy < (1.0 - Float.Epsilon) {
			if sy > -(1.0 - Float.Epsilon) {
				euler.X = Angle.Atan2(b.Z.Y, b.Z.Z)
				euler.Y = Angle.Asin(-sy)
				euler.Z = Angle.Atan2(b.Y.X, b.X.X)
			} else {
				// It's -1
				euler.X = 0
				euler.Y = Angle.Pi / 2.0
				euler.Z = -Angle.Atan2(b.X.Y, b.Y.Y)
			}
		} else {
			// It's 1
			euler.X = 0
			euler.Y = -Angle.Pi / 2.0
			euler.Z = -Angle.Atan2(b.X.Y, b.Y.Y)
		}
		return euler
	default:
		panic("Invalid parameter for get_euler(order)")
	}
}

// Scale assuming that the matrix is the combination of a rotation and scaling, returns the absolute value
// of scaling factors along each axis.
func Scale(b XYZ) Vector3.XYZ { //gd:Basis.get_scale
	scale := Vector3.XYZ{
		X: Vector3.Length(b.X),
		Y: Vector3.Length(b.Y),
		Z: Vector3.Length(b.Z),
	}
	if Determinant(b) < 0 {
		scale.X = -scale.X // Adjust one axis for reflection
	}
	return scale
}

// Inverse returns the inverse of the matrix.
func Inverse(m XYZ) XYZ { //gd:Basis.inverse
	var det = Determinant(m)
	var (
		invDet = 1.0 / det
	)
	return XYZ{
		X: Vector3.XYZ{
			X: (m.Y.Y*m.Z.Z - m.Z.Y*m.Y.Z) * invDet,
			Y: (m.Z.Y*m.X.Z - m.X.Y*m.Z.Z) * invDet,
			Z: (m.X.Y*m.Y.Z - m.Y.Y*m.X.Z) * invDet,
		},
		Y: Vector3.XYZ{
			X: (m.Z.X*m.Y.Z - m.Y.X*m.Z.Z) * invDet,
			Y: (m.X.X*m.Z.Z - m.Z.X*m.X.Z) * invDet,
			Z: (m.Y.X*m.X.Z - m.X.X*m.Y.Z) * invDet,
		},
		Z: Vector3.XYZ{
			X: (m.Y.X*m.Z.Y - m.Z.X*m.Y.Y) * invDet,
			Y: (m.Z.X*m.X.Y - m.X.X*m.Z.Y) * invDet,
			Z: (m.X.X*m.Y.Y - m.Y.X*m.X.Y) * invDet,
		},
	}
}

// IsConformal returns true if the basis is conformal, meaning it preserves angles and distance ratios,
// and may only be composed of rotation and uniform scale. Returns false if the basis has non-uniform
// scale or shear/skew. This can be used to validate if the basis is non-distorted, which is important
// for physics and other use cases.
func IsConformal(b XYZ) bool { //gd:Basis.is_conformal
	var (
		x = b.X
		y = b.Y
		z = b.Z
	)
	x_len_sq := Vector3.LengthSquared(x)
	return Float.IsApproximatelyEqual(x_len_sq, Vector3.LengthSquared(y)) &&
		Float.IsApproximatelyEqual(x_len_sq, Vector3.LengthSquared(z)) &&
		Float.IsApproximatelyZero(Vector3.Dot(x, y)) &&
		Float.IsApproximatelyZero(Vector3.Dot(x, z)) &&
		Float.IsApproximatelyZero(Vector3.Dot(y, z))
}

// IsApproximatelyEqual returns true if this basis and b are approximately equal, by calling
// [IsApproximatelyEqual] on all vector components.
func IsApproximatelyEqual(a, b XYZ) bool { //gd:Basis.is_equal_approx
	return Vector3.IsApproximatelyEqual(a.X, b.X) && Vector3.IsApproximatelyEqual(a.Y, b.Y) && Vector3.IsApproximatelyEqual(a.Z, b.Z)
}

// IsFinite returns true if this basis is finite, by calling [IsFinite] on all vector components.
func IsFinite(b XYZ) bool { //gd:Basis.is_finite
	return Vector3.IsFinite(b.X) && Vector3.IsFinite(b.Y) && Vector3.IsFinite(b.Z)
}

// Orthonormalized returns the orthonormalized version of the matrix (useful to call from time to time to avoid rounding error
// for orthogonal matrices). This performs a Gram-Schmidt orthonormalization on the basis of the matrix.
func Orthonormalized(b XYZ) XYZ { //gd:Basis.orthonormalized
	var (
		x = Vector3.Normalized(b.X)
		y = Vector3.Normalized(Vector3.Sub(b.Y, Vector3.MulX(x, Vector3.Dot(x, b.Y))))
		z = Vector3.Normalized(Vector3.Sub(Vector3.Sub(b.Z, Vector3.MulX(x, Vector3.Dot(x, b.Z))), Vector3.MulX(y, Vector3.Dot(y, b.Z))))
	)
	return XYZ{x, y, z}
}

// Rotated returns a copy of the basis rotated around the given axis by the given angle (in radians).
// The axis must be a normalized vector.
func Rotated(b XYZ, axis Vector3.XYZ, angle Angle.Radians) XYZ { //gd:Basis.rotated
	rotation := RotatesAxisAngle(axis, angle) // Assume this constructs a rotation matrix
	return Mul(rotation, b)
}

// Scaled introduce an additional scaling specified by the given 3D scaling factor.
func Scaled(b XYZ, scale Vector3.XYZ) XYZ { //gd:Basis.scaled
	b.X.X *= scale.X
	b.X.Y *= scale.X
	b.X.Z *= scale.X
	b.Y.X *= scale.Y
	b.Y.Y *= scale.Y
	b.Y.Z *= scale.Y
	b.Z.X *= scale.Z
	b.Z.Y *= scale.Z
	b.Z.Z *= scale.Z
	return b
}

// Slerp assuming that the matrix is a proper rotation matrix, slerp performs a spherical-linear interpolation
// with another rotation matrix.
func Slerp[X Float.Any](a, b XYZ, weight X) XYZ { //gd:Basis.slerp
	//consider scale
	var from = AsQuaternion(a)
	var to = AsQuaternion(b)
	var result = qAsBasis(qSlerp(from, to, weight))
	result.X = Vector3.MulX(result.X, Float.Lerp(Vector3.Length(a.X), Vector3.Length(b.X), Float.X(weight)))
	result.Y = Vector3.MulX(result.Y, Float.Lerp(Vector3.Length(a.Y), Vector3.Length(b.Y), Float.X(weight)))
	result.Z = Vector3.MulX(result.Z, Float.Lerp(Vector3.Length(a.Z), Vector3.Length(b.Z), Float.X(weight)))
	return result
}

// TransposedDotX returns the transposed dot product with the X axis of the matrix.
func TransposedDotX(b XYZ, v Vector3.XYZ) Float.X { //gd:Basis.tdotx
	return b.X.X*v.X + b.Y.X*v.Y + b.Z.X*v.Z
}

// TransposedDotY returns the transposed dot product with the Y axis of the matrix.
func TransposedDotY(b XYZ, v Vector3.XYZ) Float.X { //gd:Basis.tdoty
	return b.X.Y*v.X + b.Y.Y*v.Y + b.Z.Y*v.Z
}

// TransposedDotZ returns the transposed dot product with the Z axis of the matrix.
func TransposedDotZ(b XYZ, v Vector3.XYZ) Float.X { //gd:Basis.tdotz
	return b.X.Z*v.X + b.Y.Z*v.Y + b.Z.Z*v.Z
}

// Transposed returns the transposed version of the matrix.
func Transposed(m XYZ) XYZ { //gd:Basis.transposed
	return XYZ{
		Vector3.New(m.X.X, m.Y.X, m.Z.X),
		Vector3.New(m.X.Y, m.Y.Y, m.Z.Y),
		Vector3.New(m.X.Z, m.Y.Z, m.Z.Z),
	}
}

func Mul(a, b XYZ) XYZ { //gd:Basis*Basis
	return XYZ{
		X: Transform(b.X, a),
		Y: Transform(b.Y, a),
		Z: Transform(b.Z, a),
	}
}

type quaternion = struct {
	I Float.X
	J Float.X
	K Float.X
	X Float.X
}

// AsQuaternion returns the basis's rotation in the form of a quaternion. See [AsEulerAngles] if you
// need Euler angles, but keep in mind quaternions should generally be preferred to [AsEulerAngles].
func AsQuaternion(b XYZ) quaternion { //gd:Basis.get_rotation_quaternion
	// Assumes that the matrix can be decomposed into a proper rotation and scaling matrix as M = R.S,
	// and returns the Euler angles corresponding to the rotation part, complementing get_scale().
	// See the comment in get_scale() for further information.
	var m = Orthonormalized(b)
	var det = Determinant(m)
	if det < 0 {
		// Ensure that the determinant is 1, such that result is a proper rotation matrix which can be represented by Euler angles.
		m = Scaled(m, Vector3.New(-1, -1, -1))
	}
	var (
		trace = m.X.X + m.Y.Y + m.Z.Z
		temp  [4]Float.X
	)
	if trace > 0.0 {
		var s = Float.Sqrt(trace + 1.0)
		temp[3] = (s * 0.5)
		s = 0.5 / s
		temp[0] = ((m.Z.Y - m.Y.Z) * s)
		temp[1] = ((m.X.Z - m.Z.X) * s)
		temp[2] = ((m.Y.X - m.X.Y) * s)
	} else {
		var i int
		if m.X.X < m.Y.Y {
			if m.Y.Y < m.Z.Z {
				i = 2
			} else {
				i = 1
			}
		} else {
			if m.X.X < m.Z.Z {
				i = 2
			} else {
				i = 0
			}
		}
		var (
			j = (i + 1) % 3
			k = (i + 2) % 3
			a = *(*[3][3]Float.X)(unsafe.Pointer(&m))
			s = Float.Sqrt(a[i][i] - a[j][j] - a[k][k] + 1.0)
		)
		temp[i] = s * 0.5
		s = 0.5 / s
		temp[3] = (a[k][j] - a[j][k]) * s
		temp[j] = (a[j][i] + a[i][j]) * s
		temp[k] = (a[k][i] + a[i][k]) * s
	}
	return quaternion{temp[0], temp[1], temp[2], temp[3]}
}

// Slerp returns the result of the spherical linear interpolation between this quaternion
// and to by amount weight.
//
// Note: Both quaternions must be normalized.
func qSlerp[X Float.Any](a, b quaternion, weight X) quaternion {
	var to1 quaternion
	var (
		cosom, sinom, scale0, scale1 Float.X
		omega                        Angle.Radians
	)
	cosom = qLengthSquared(b) // calc cosine
	if cosom < 0.0 {          // adjust signs (if necessary)
		cosom = -cosom
		to1 = qNeg(b)
	} else {
		to1 = b
	}
	// calculate coefficients
	if (1.0 - cosom) > Float.Epsilon { // standard case (slerp)
		omega = Angle.Acos(cosom)
		sinom = Angle.Sin(omega)
		scale0 = Angle.Sin(Angle.Radians(1.0-weight)*omega) / sinom
		scale1 = Angle.Sin(Angle.Radians(weight)*omega) / sinom
	} else {
		// "from" and "to" quaternions are very close
		//  ... so we can do a linear interpolation
		scale0 = 1.0 - Float.X(weight)
		scale1 = Float.X(weight)
	}
	return quaternion{ // calculate final values
		scale0*a.I + scale1*to1.I,
		scale0*a.J + scale1*to1.J,
		scale0*a.K + scale1*to1.K,
		scale0*a.X + scale1*to1.X,
	}
}

// LengthSquared returns the length of the quaternion, squared.
func qLengthSquared(q quaternion) Float.X {
	return q.I*q.I + q.J*q.J + q.K*q.K + q.X*q.X
}

func qNeg(v quaternion) quaternion { return quaternion{-v.I, -v.J, -v.K, -v.X} }

func qAsBasis(q quaternion) XYZ {
	var d = qLengthSquared(q)
	var s = 2.0 / d
	var xs, ys, zs = q.I * s, q.J * s, q.K * s
	var wx, wy, wz = q.X * xs, q.X * ys, q.X * zs
	var xx, xy, xz = q.I * xs, q.I * ys, q.I * zs
	var yy, yz, zz = q.J * ys, q.J * zs, q.K * zs
	return XYZ{
		X: Vector3.New(1.0-(yy+zz), xy-wz, xz+wy),
		Y: Vector3.New(xy+wz, 1.0-(xx+zz), yz-wx),
		Z: Vector3.New(xz-wy, yz+wx, 1.0-(xx+yy)),
	}
}

func Transform(v Vector3.XYZ, m XYZ) Vector3.XYZ { //gd:Basis*(right:Vector3)
	return Vector3.XYZ{
		X: m.X.X*v.X + m.Y.X*v.Y + m.Z.X*v.Z,
		Y: m.X.Y*v.X + m.Y.Y*v.Y + m.Z.Y*v.Z,
		Z: m.X.Z*v.X + m.Y.Z*v.Y + m.Z.Z*v.Z,
	}
}
