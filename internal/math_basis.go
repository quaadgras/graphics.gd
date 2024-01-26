//go:build !generate

package gd

import "math"

type Basis [3]Vector3

func NewRotationAroundAxis(axis Vector3, angle Radians) Basis {
	var rows Basis
	var axis_sq = Vector3{axis[X] * axis[X], axis[Y] * axis[Y], axis[Z] * axis[Z]}
	var cosine = float32(Cos(angle))
	rows[0][0] = axis_sq[X] + cosine*(1.0-axis_sq[X])
	rows[1][1] = axis_sq[Y] + cosine*(1.0-axis_sq[Y])
	rows[2][2] = axis_sq[Z] + cosine*(1.0-axis_sq[Z])

	var sine = float32(Sin(angle))
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

func (b Basis) cofac(row1, col1, row2, col2 int) float32 {
	return (b[row1][col1]*b[row2][col2] - b[row1][col2]*b[row2][col1])
}

func (b Basis) Inverse() Basis {
	var co = [3]float32{
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

func (m Basis) Transposed() Basis {
	return Basis{
		{m[0][0], m[1][0], m[2][0]},
		{m[0][1], m[1][1], m[2][1]},
		{m[0][2], m[1][2], m[2][2]},
	}
}

func (b Basis) Orthonormalized() Basis {
	var (
		x = b[0].Normalized()
		y = b[Y].Sub(Vector3.Mulf(x, Vector3.Dot(x, b[Y]))).Normalized()
		z = b[Z].Sub(Vector3.Mulf(x, Vector3.Dot(x, b[Z]))).Sub(Vector3.Mulf(y, Vector3.Dot(y, b[Z]))).Normalized()
	)
	return Basis{x, y, z}
}

func (b Basis) Determinant() Float {
	return Float(b[0][0]*(b[1][1]*b[2][2]-b[2][1]*b[1][2]) -
		b[1][0]*(b[0][1]*b[2][2]-b[2][1]*b[0][2]) +
		b[2][0]*(b[0][1]*b[1][2]-b[1][1]*b[0][2]))
}

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

func (b Basis) EulerAngles(order EulerOrder) EulerAngles {
	switch order {
	case EulerOrderXyz:
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
	case EulerOrderXzy:
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
	case EulerOrderYxz:
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
	case EulerOrderYzx:
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
	case EulerOrderZxy:
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
	case EulerOrderZyx:
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

func (b Basis) GetRotationQuaternion() Quaternion {
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
	var temp [4]float32

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

func (m Basis) Transform(v Vector3) Vector3 {
	return Vector3{
		float32(m[0].Dot(v)),
		float32(m[1].Dot(v)),
		float32(m[2].Dot(v)),
	}
}

func (m Basis) GetQuaternion() Quaternion {
	/* Allow getting a quaternion from an unnormalized transform */
	var (
		trace = m[0][0] + m[1][1] + m[2][2]
		temp  [4]float32
	)
	if trace > 0.0 {
		var s = float32(math.Sqrt(float64(trace) + 1.0))
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
		var s = float32(math.Sqrt(float64(m[i][i] - m[j][j] - m[k][k] + 1.0)))
		temp[i] = s * 0.5
		s = 0.5 / s
		temp[3] = (m[k][j] - m[j][k]) * s
		temp[j] = (m[j][i] + m[i][j]) * s
		temp[k] = (m[k][i] + m[i][k]) * s
	}
	return Quaternion{temp[0], temp[1], temp[2], temp[3]}
}
