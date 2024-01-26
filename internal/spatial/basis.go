package spatial

import "math"

/*
Basis is a 3×3 matrix used for representing 3D rotation and scale. Usually used as an orthogonal basis for a [Transform3D].

Contains 3 vector fields X, Y and Z as its columns, which are typically interpreted as the local basis vectors of a transformation.
For such use, it is composed of a scaling and a rotation matrix, in that order (M = R.S).

Basis can also be accessed as an array of 3D vectors. These vectors are usually orthogonal to each other, but are not necessarily
normalized (due to scaling).
*/
type Basis [3]Vector3

// "Fields"

func (b Basis) X() Vector3 { return b[0] }
func (b Basis) Y() Vector3 { return b[1] }
func (b Basis) Z() Vector3 { return b[2] }

func (b *Basis) SetX(x Vector3) { b[0] = x }
func (b *Basis) SetY(y Vector3) { b[1] = y }
func (b *Basis) SetZ(z Vector3) { b[2] = z }

// "Constants"

func (Basis) IDENTITY() Basis {
	return Basis{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}
func (Basis) FLIP_X() Basis {
	return Basis{
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}
func (Basis) FLIP_Y() Basis {
	return Basis{
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	}
}
func (Basis) FLIP_Z() Basis {
	return Basis{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, -1},
	}
}

// NewBasisScaledBy constructs a pure scale basis matrix with no rotation or shearing. The scale values are set as
// the diagonal of the matrix, and the other parts of the matrix are zero.
func NewBasisScaledBy(scale Vector3) Basis {
	return Basis{
		{scale[X], 0, 0},
		{0, scale[Y], 0},
		{0, 0, scale[Z]},
	}
}

// NewBasisRotatedAround constructs a pure rotation basis matrix, rotated around the given axis by angle (in radians).
// The axis must be a normalized vector.
func NewBasisRotatedAround(axis Vector3, angle Radians) Basis {
	var rows Basis
	var axis_sq = Vector3{axis[X] * axis[X], axis[Y] * axis[Y], axis[Z] * axis[Z]}
	var cosine = float(Cos(angle))
	rows[0][0] = axis_sq[X] + cosine*(1.0-axis_sq[X])
	rows[1][1] = axis_sq[Y] + cosine*(1.0-axis_sq[Y])
	rows[2][2] = axis_sq[Z] + cosine*(1.0-axis_sq[Z])

	var sine = float(Sin(angle))
	var t = 1 - cosine

	var xyzt = axis[X] * axis[Y] * t
	var zyxs = axis[Z] * sine
	rows[0][1] = xyzt - zyxs
	rows[1][0] = xyzt + zyxs

	xyzt = axis[X] * axis[Z] * t
	zyxs = axis[Y] * sine
	rows[0][2] = xyzt + zyxs
	rows[2][0] = xyzt - zyxs

	xyzt = axis[Y] * axis[Z] * t
	zyxs = axis[X] * sine
	rows[1][2] = xyzt - zyxs
	rows[2][1] = xyzt + zyxs
	return rows
}

func (b Basis) cofac(row1, col1, row2, col2 int) float {
	return (b[row1][col1]*b[row2][col2] - b[row1][col2]*b[row2][col1])
}

// "Methods"

// Determinant returns the determinant of the basis matrix. If the basis is uniformly scaled,
// its determinant is the square of the scale.
//
// A negative determinant means the basis has a negative scale. A zero determinant means the
// basis isn't invertible, and is usually considered invalid.
func (b Basis) Determinant() Float {
	return Float(b[0][0]*(b[1][1]*b[2][2]-b[2][1]*b[1][2]) -
		b[1][0]*(b[0][1]*b[2][2]-b[2][1]*b[0][2]) +
		b[2][0]*(b[0][1]*b[1][2]-b[1][1]*b[0][2]))
}

// EulerAngles returns the basis's rotation in the form of Euler angles. The Euler order depends
// on the order parameter, by default it uses the YXZ convention: when decomposing, first Z,
// then X, and Y last. The returned vector contains the rotation angles in the format
// (X angle, Y angle, Z angle).
//
// Consider using the [Basis.Quaternion] method instead, which returns a [Quaternion]
// quaternion instead of [EulerAngles].
func (b Basis) EulerAngles(order EulerOrder) EulerAngles {
	switch order {
	case EulerOrderXYZ:
		// Euler angles in XYZ convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cy*cz          -cy*sz           sy
		//        cz*sx*sy+cx*sz  cx*cz-sx*sy*sz -cy*sx
		//       -cx*cz*sy+sx*sz  cz*sx+cx*sy*sz  cx*cy
		var euler EulerAngles
		var sy = b[0][2]
		if sy < (1.0 - cmpEpsilon) {
			if sy > -(1.0 - cmpEpsilon) {
				// is this a pure Y rotation?
				if b[1][0] == 0 && b[0][1] == 0 && b[1][2] == 0 && b[2][1] == 0 && b[1][1] == 1 {
					// return the simplest form (human friendlier in editor and scripts)
					euler[X] = 0
					euler[Y] = Atan2(b[0][2], b[0][0])
					euler[Z] = 0
				} else {
					euler[X] = Atan2(-b[1][2], b[2][2])
					euler[Y] = Asin(sy)
					euler[Z] = Atan2(-b[0][1], b[0][0])
				}
			} else {
				euler[X] = Atan2(b[2][1], b[1][1])
				euler[Y] = -Pi / 2.0
				euler[Z] = 0.0
			}
		} else {
			euler[X] = Atan2(b[2][1], b[1][1])
			euler[Y] = Pi / 2.0
			euler[Z] = 0.0
		}
		return euler
	case EulerOrderXZY:
		// Euler angles in XZY convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cz*cy             -sz             cz*sy
		//        sx*sy+cx*cy*sz    cx*cz           cx*sz*sy-cy*sx
		//        cy*sx*sz          cz*sx           cx*cy+sx*sz*sy

		var euler EulerAngles
		var sz = b[0][1]
		if sz < (1.0 - cmpEpsilon) {
			if sz > -(1.0 - cmpEpsilon) {
				euler[X] = Atan2(b[2][1], b[1][1])
				euler[Y] = Atan2(b[0][2], b[0][0])
				euler[Z] = Asin(-sz)
			} else {
				// It's -1
				euler[X] = -Atan2(b[1][2], b[2][2])
				euler[Y] = 0.0
				euler[Z] = Pi / 2.0
			}
		} else {
			// It's 1
			euler[X] = -Atan2(b[1][2], b[2][2])
			euler[Y] = 0.0
			euler[Z] = -Pi / 2.0
		}
		return euler
	case EulerOrderYXZ:
		// Euler angles in YXZ convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cy*cz+sy*sx*sz    cz*sy*sx-cy*sz        cx*sy
		//        cx*sz             cx*cz                 -sx
		//        cy*sx*sz-cz*sy    cy*cz*sx+sy*sz        cy*cx

		var euler EulerAngles

		var m12 = b[1][2]

		if m12 < (1 - cmpEpsilon) {
			if m12 > -(1 - cmpEpsilon) {
				// is this a pure X rotation?
				if b[1][0] == 0 && b[0][1] == 0 && b[0][2] == 0 && b[2][0] == 0 && b[0][0] == 1 {
					// return the simplest form (human friendlier in editor and scripts)
					euler[X] = Atan2(-m12, b[1][1])
					euler[Y] = 0
					euler[Z] = 0
				} else {
					euler[X] = Asin(-m12)
					euler[Y] = Atan2(b[0][2], b[2][2])
					euler[Z] = Atan2(b[1][0], b[1][1])
				}
			} else { // m12 == -1
				euler[X] = Pi * 0.5
				euler[Y] = Atan2(b[0][1], b[0][0])
				euler[Z] = 0
			}
		} else { // m12 == 1
			euler[X] = -Pi * 0.5
			euler[Y] = -Atan2(b[0][1], b[0][0])
			euler[Z] = 0
		}

		return euler
	case EulerOrderYZX:
		// Euler angles in YZX convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cy*cz             sy*sx-cy*cx*sz     cx*sy+cy*sz*sx
		//        sz                cz*cx              -cz*sx
		//        -cz*sy            cy*sx+cx*sy*sz     cy*cx-sy*sz*sx

		var euler EulerAngles
		var sz = b[1][0]
		if sz < (1.0 - cmpEpsilon) {
			if sz > -(1.0 - cmpEpsilon) {
				euler[X] = Atan2(-b[1][2], b[1][1])
				euler[Y] = Atan2(-b[2][0], b[0][0])
				euler[Z] = Asin(sz)
			} else {
				// It's -1
				euler[X] = Atan2(b[2][1], b[2][2])
				euler[Y] = 0.0
				euler[Z] = -Pi / 2.0
			}
		} else {
			// It's 1
			euler[X] = Atan2(b[2][1], b[2][2])
			euler[Y] = 0.0
			euler[Z] = Pi / 2.0
		}
		return euler
	case EulerOrderZXY:
		// Euler angles in ZXY convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cz*cy-sz*sx*sy    -cx*sz                cz*sy+cy*sz*sx
		//        cy*sz+cz*sx*sy    cz*cx                 sz*sy-cz*cy*sx
		//        -cx*sy            sx                    cx*cy
		var euler EulerAngles
		var sx = b[2][1]
		if sx < (1.0 - cmpEpsilon) {
			if sx > -(1.0 - cmpEpsilon) {
				euler[X] = Asin(sx)
				euler[Y] = Atan2(-b[2][0], b[2][2])
				euler[Z] = Atan2(-b[0][1], b[1][1])
			} else {
				// It's -1
				euler[X] = -Pi / 2.0
				euler[Y] = Atan2(b[0][2], b[0][0])
				euler[Z] = 0
			}
		} else {
			// It's 1
			euler[X] = Pi / 2.0
			euler[Y] = Atan2(b[0][2], b[0][0])
			euler[Z] = 0
		}
		return euler
	case EulerOrderZYX:
		// Euler angles in ZYX convention.
		// See https://en.wikipedia.org/wiki/Euler_angles#Rotation_matrix
		//
		// rot =  cz*cy             cz*sy*sx-cx*sz        sz*sx+cz*cx*cy
		//        cy*sz             cz*cx+sz*sy*sx        cx*sz*sy-cz*sx
		//        -sy               cy*sx                 cy*cx
		var euler EulerAngles
		var sy = b[2][0]
		if sy < (1.0 - cmpEpsilon) {
			if sy > -(1.0 - cmpEpsilon) {
				euler[X] = Atan2(b[2][1], b[2][2])
				euler[Y] = Asin(-sy)
				euler[Z] = Atan2(b[1][0], b[0][0])
			} else {
				// It's -1
				euler[X] = 0
				euler[Y] = Pi / 2.0
				euler[Z] = -Atan2(b[0][1], b[1][1])
			}
		} else {
			// It's 1
			euler[X] = 0
			euler[Y] = -Pi / 2.0
			euler[Z] = -Atan2(b[0][1], b[1][1])
		}
		return euler
	default:
		panic("Invalid parameter for get_euler(order)")
	}
}

// Quaternion returns the basis's rotation in the form of a quaternion. See [Basis.EulerAngles] if you
// need Euler angles, but keep in mind quaternions should generally be preferred to [EulerAngles].
func (b Basis) Quaternion() Quaternion {
	// Assumes that the matrix can be decomposed into a proper rotation and scaling matrix as M = R.S,
	// and returns the Euler angles corresponding to the rotation part, complementing get_scale().
	// See the comment in get_scale() for further information.
	var m = b.Orthonormalized()
	var det = m.Determinant()
	if det < 0 {
		// Ensure that the determinant is 1, such that result is a proper rotation matrix which can be represented by Euler angles.
		m = m.Scaled(Vector3{-1, -1, -1})
	}
	var trace = m[0][0] + m[1][1] + m[2][2]
	var temp [4]float

	if trace > 0.0 {
		var s = Sqrt(trace + 1.0)
		temp[3] = (s * 0.5)
		s = 0.5 / s

		temp[0] = ((m[2][1] - m[1][2]) * s)
		temp[1] = ((m[0][2] - m[2][0]) * s)
		temp[2] = ((m[1][0] - m[0][1]) * s)
	} else {
		var i int
		if m[0][0] < m[1][1] {
			if m[1][1] < m[2][2] {
				i = 2
			} else {
				i = 1
			}
		} else {
			if m[0][0] < m[2][2] {
				i = 2
			} else {
				i = 0
			}
		}
		var j = (i + 1) % 3
		var k = (i + 2) % 3

		var s = Sqrt(m[i][i] - m[j][j] - m[k][k] + 1.0)
		temp[i] = s * 0.5
		s = 0.5 / s

		temp[3] = (m[k][j] - m[j][k]) * s
		temp[j] = (m[j][i] + m[i][j]) * s
		temp[k] = (m[k][i] + m[i][k]) * s
	}
	return Quaternion{temp[0], temp[1], temp[2], temp[3]}
}

// GetScale assuming that the matrix is the combination of a rotation and scaling, returns the absolute value
// of scaling factors along each axis.
func (b Basis) GetScale() Vector3 {
	// FIXME: We are assuming M = R.S (R is rotation and S is scaling), and use polar decomposition to extract R and S.
	// A polar decomposition is M = O.P, where O is an orthogonal matrix (meaning rotation and reflection) and
	// P is a positive semi-definite matrix (meaning it contains absolute values of scaling along its diagonal).
	//
	// Despite being different from what we want to achieve, we can nevertheless make use of polar decomposition
	// here as follows. We can split O into a rotation and a reflection as O = R.Q, and obtain M = R.S where
	// we defined S = Q.P. Now, R is a proper rotation matrix and S is a (signed) scaling matrix,
	// which can involve negative scalings. However, there is a catch: unlike the polar decomposition of M = O.P,
	// the decomposition of O into a rotation and reflection matrix as O = R.Q is not unique.
	// Therefore, we are going to do this decomposition by sticking to a particular convention.
	// This may lead to confusion for some users though.
	//
	// The convention we use here is to absorb the sign flip into the scaling matrix.
	// The same convention is also used in other similar functions such as get_rotation_axis_angle, get_rotation, ...
	//
	// A proper way to get rid of this issue would be to store the scaling values (or at least their signs)
	// as a part of Basis. However, if we go that path, we need to disable direct (write) access to the
	// matrix elements.
	//
	// The rotation part of this decomposition is returned by get_rotation* functions.
	var det_sign = Sign(b.Determinant())
	return Vector3{
		float(Vector3{b[0][0], b[1][0], b[2][0]}.Length()),
		float(Vector3{b[0][1], b[1][1], b[2][1]}.Length()),
		float(Vector3{b[0][2], b[1][2], b[2][2]}.Length()),
	}.Mulf(det_sign)
}

// Inverse returns the inverse of the matrix.
func (b Basis) Inverse() Basis {
	var co = [3]float{
		b.cofac(1, 1, 2, 2), b.cofac(1, 2, 2, 0), b.cofac(1, 0, 2, 1),
	}
	var det = b[0][0]*co[0] + b[0][1]*co[1] + b[0][2]*co[2]
	var (
		s = 1.0 / det
	)
	return Basis{
		{co[0] * s, b.cofac(0, 2, 2, 1) * s, b.cofac(0, 1, 1, 2) * s},
		{co[1] * s, b.cofac(0, 0, 2, 2) * s, b.cofac(0, 2, 1, 0) * s},
		{co[2] * s, b.cofac(0, 1, 2, 0) * s, b.cofac(0, 0, 1, 1) * s},
	}
}

// IsConformal returns true if the basis is conformal, meaning it preserves angles and distance ratios,
// and may only be composed of rotation and uniform scale. Returns false if the basis has non-uniform
// scale or shear/skew. This can be used to validate if the basis is non-distorted, which is important
// for physics and other use cases.
func (b Basis) IsConformal() bool {
	var (
		x = b[0]
		y = b[1]
		z = b[2]
	)
	x_len_sq := x.LengthSquared()
	return IsApproximatelyEqual(x_len_sq, y.LengthSquared()) && IsApproximatelyEqual(x_len_sq, z.LengthSquared()) &&
		IsApproximatelyZero(x.Dot(y)) && IsApproximatelyZero(x.Dot(z)) && IsApproximatelyZero(y.Dot(z))
}

// IsApproximatelyEqual returns true if this basis and b are approximately equal, by calling
// [IsApproximatelyEqual] on all vector components.
func (b Basis) IsApproximatelyEqual(other Basis) bool {
	return b[0].IsApproximatelyEqual(other[0]) && b[1].IsApproximatelyEqual(other[1]) && b[2].IsApproximatelyEqual(other[2])
}

// IsFinite returns true if this basis is finite, by calling [IsFinite] on all vector components.
func (b Basis) IsFinite() bool {
	return b[0].IsFinite() && b[1].IsFinite() && b[2].IsFinite()
}

// LookingAt creates a Basis with a rotation such that the forward axis (-Z) points towards the target position.
//
// The up axis (+Y) points as close to the up vector as possible while staying perpendicular to the forward axis.
// The resulting Basis is orthonormalized. The target and up vectors cannot be zero, and cannot be parallel to each other.
//
// If use_model_front is true, the +Z axis (asset front) is treated as forward (implies +X is left) and points toward the
// target position. By default, the -Z axis (camera forward) is treated as forward (implies +X is right).
func (b Basis) LookingAt(target, up Vector3, use_model_front bool) Basis {
	var v_z = target.Normalized()
	if !use_model_front {
		v_z = v_z.Neg()
	}
	var v_x = up.Cross(v_z)
	v_x = v_x.Normalized()
	var v_y = v_z.Cross(v_x)
	return Basis{v_x, v_y, v_z}
}

// Orthonormalized returns the orthonormalized version of the matrix (useful to call from time to time to avoid rounding error
// for orthogonal matrices). This performs a Gram-Schmidt orthonormalization on the basis of the matrix.
func (b Basis) Orthonormalized() Basis {
	var (
		x = b[0].Normalized()
		y = b[Y].Sub(Vector3.Mulf(x, Vector3.Dot(x, b[Y]))).Normalized()
		z = b[Z].Sub(Vector3.Mulf(x, Vector3.Dot(x, b[Z]))).Sub(Vector3.Mulf(y, Vector3.Dot(y, b[Z]))).Normalized()
	)
	return Basis{x, y, z}
}

// Rotated returns a copy of the basis rotated around the given axis by the given angle (in radians).
// The axis must be a normalized vector.
func (b Basis) Rotated(axis Vector3, angle Radians) Basis {
	return NewBasisRotatedAround(axis, angle).Mul(b)
}

// Scaled introduce an additional scaling specified by the given 3D scaling factor.
func (b Basis) Scaled(scale Vector3) Basis {
	b[0][0] *= scale[X]
	b[0][1] *= scale[X]
	b[0][2] *= scale[X]
	b[1][0] *= scale[Y]
	b[1][1] *= scale[Y]
	b[1][2] *= scale[Y]
	b[2][0] *= scale[Z]
	b[2][1] *= scale[Z]
	b[2][2] *= scale[Z]
	return b
}

// Slerp assuming that the matrix is a proper rotation matrix, slerp performs a spherical-linear interpolation
// with another rotation matrix.
func (b Basis) Slerp(to Basis, weight Float) Basis { return b.lerp(to, weight) }

// TransposedDotX returns the transposed dot product with the X axis of the matrix.
func (b Basis) TransposedDotX(v Vector3) Float {
	return Float(b[0][0]*v[0] + b[1][0]*v[1] + b[2][0]*v[2])
}

// TransposedDotY returns the transposed dot product with the Y axis of the matrix.
func (b Basis) TransposedDotY(v Vector3) Float {
	return Float(b[0][1]*v[0] + b[1][1]*v[1] + b[2][1]*v[2])
}

// TransposedDotZ returns the transposed dot product with the Z axis of the matrix.
func (b Basis) TransposedDotZ(v Vector3) Float {
	return Float(b[0][2]*v[0] + b[1][2]*v[1] + b[2][2]*v[2])
}

// Transposed returns the transposed version of the matrix.
func (m Basis) Transposed() Basis {
	return Basis{
		{m[0][0], m[1][0], m[2][0]},
		{m[0][1], m[1][1], m[2][1]},
		{m[0][2], m[1][2], m[2][2]},
	}
}

func (b Basis) Mul(other Basis) Basis {
	return Basis{
		{float(other.TransposedDotX(b[0])), float(other.TransposedDotY(b[0])), float(other.TransposedDotZ(b[0]))},
		{float(other.TransposedDotX(b[1])), float(other.TransposedDotY(b[1])), float(other.TransposedDotZ(b[1]))},
		{float(other.TransposedDotX(b[2])), float(other.TransposedDotY(b[2])), float(other.TransposedDotZ(b[2]))},
	}
}

func (m Basis) Transform(v Vector3) Vector3 {
	return Vector3{
		float(m[0].Dot(v)),
		float(m[1].Dot(v)),
		float(m[2].Dot(v)),
	}
}

func (m Basis) getQuaternion() Quaternion {
	/* Allow getting a quaternion from an unnormalized transform */
	var (
		trace = m[0][0] + m[1][1] + m[2][2]
		temp  [4]float
	)
	if trace > 0.0 {
		var s = float(math.Sqrt(float64(trace) + 1.0))
		temp[3] = (s * 0.5)
		s = 0.5 / s
		temp[0] = ((m[2][1] - m[1][2]) * s)
		temp[1] = ((m[0][2] - m[2][0]) * s)
		temp[2] = ((m[1][0] - m[0][1]) * s)
	} else {
		var i int
		if m[0][0] < m[1][1] {
			if m[1][1] < m[2][2] {
				i = 2
			} else {
				i = 1
			}
		} else {
			if m[0][0] < m[2][2] {
				i = 2
			} else {
				i = 0
			}
		}
		var (
			j int = (i + 1) % 3
			k int = (i + 2) % 3
		)
		var s = float(math.Sqrt(float64(m[i][i] - m[j][j] - m[k][k] + 1.0)))
		temp[i] = s * 0.5
		s = 0.5 / s
		temp[3] = (m[k][j] - m[j][k]) * s
		temp[j] = (m[j][i] + m[i][j]) * s
		temp[k] = (m[k][i] + m[i][k]) * s
	}
	return Quaternion{temp[0], temp[1], temp[2], temp[3]}
}
