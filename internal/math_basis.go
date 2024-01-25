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
