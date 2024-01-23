//go:build !generate

package gd

import "math"

type Basis struct {
	rows [3]Vector3
}

func (m Basis) GetQuaternion() Quaternion {
	/* Allow getting a quaternion from an unnormalized transform */
	var (
		trace = m.rows[0][0] + m.rows[1][1] + m.rows[2][2]
		temp  [4]float32
	)
	if trace > 0.0 {
		var s = float32(math.Sqrt(float64(trace) + 1.0))
		temp[3] = (s * 0.5)
		s = 0.5 / s
		temp[0] = ((m.rows[2][1] - m.rows[1][2]) * s)
		temp[1] = ((m.rows[0][2] - m.rows[2][0]) * s)
		temp[2] = ((m.rows[1][0] - m.rows[0][1]) * s)
	} else {
		var i int
		if m.rows[0][0] < m.rows[1][1] {
			if m.rows[1][1] < m.rows[2][2] {
				i = 2
			} else {
				i = 1
			}
		} else {
			if m.rows[0][0] < m.rows[2][2] {
				i = 2
			} else {
				i = 0
			}
		}
		var (
			j int = (i + 1) % 3
			k int = (i + 2) % 3
		)
		var s = float32(math.Sqrt(float64(m.rows[i][i] - m.rows[j][j] - m.rows[k][k] + 1.0)))
		temp[i] = s * 0.5
		s = 0.5 / s
		temp[3] = (m.rows[k][j] - m.rows[j][k]) * s
		temp[j] = (m.rows[j][i] + m.rows[i][j]) * s
		temp[k] = (m.rows[k][i] + m.rows[i][k]) * s
	}
	return Quaternion{temp[0], temp[1], temp[2], temp[3]}
}
