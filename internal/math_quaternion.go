//go:build !generate

package gd

type Quaternion [4]float32

func (q Quaternion) LengthSquared() Float {
	return Float(dot4(q))
}

func (q Quaternion) Basis() Basis {
	var d = float32(q.LengthSquared())
	var s = 2.0 / d
	var xs, ys, zs = q[x] * s, q[y] * s, q[z] * s
	var wx, wy, wz = q[w] * xs, q[w] * ys, q[w] * zs
	var xx, xy, xz = q[x] * xs, q[x] * ys, q[x] * zs
	var yy, yz, zz = q[y] * ys, q[y] * zs, q[z] * zs
	return Basis{
		{1.0 - (yy + zz), xy - wz, xz + wy},
		{xy + wz, 1.0 - (xx + zz), yz - wx},
		{xz - wy, yz + wx, 1.0 - (xx + yy)},
	}
}
