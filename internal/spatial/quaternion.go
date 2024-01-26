package spatial

import "math"

/*
Quaternions are similar to Basis, which implements the matrix representation of rotations. Unlike Basis, which stores
rotation, scale, and shearing, quaternions only store rotation.

Quaternions can be parametrized using both an axis-angle pair or Euler angles. Due to their compactness and the way
they are stored in memory, certain operations (obtaining axis-angle and performing SLERP, in particular) are more
efficient and robust against floating-point errors.

Note: Quaternions need to be normalized before being used for rotation.
*/
type Quaternion [4]float

// NewQuaternion constructs a quaternion representing the shortest arc between two points
// on the surface of a sphere with a radius of 1.0.
func NewQuaternion(arc_from, arc_to Vector3) Quaternion {
	var (
		c = arc_from.Cross(arc_to)
		d = arc_from.Dot(arc_to)
	)
	if d < -1.0+cmpEpsilon {
		return Quaternion{0, 1, 0, 0}
	}
	var s = float(Sqrt((1.0 + d) * 2.0))
	var rs = float(1.0 / s)
	return Quaternion{
		c[X] * rs,
		c[Y] * rs,
		c[Z] * rs,
		s * 0.5,
	}
}

// "Fields"

func (q Quaternion) X() Float { return Float(q[x]) }
func (q Quaternion) Y() Float { return Float(q[y]) }
func (q Quaternion) Z() Float { return Float(q[z]) }
func (q Quaternion) W() Float { return Float(q[w]) }

func (q Quaternion) SetX(v Float) { q[x] = float(v) }
func (q Quaternion) SetY(v Float) { q[y] = float(v) }
func (q Quaternion) SetZ(v Float) { q[z] = float(v) }
func (q Quaternion) SetW(v Float) { q[w] = float(v) }

// "Constants"

func (Quaternion) IDENTITY() Quaternion { return Quaternion{0, 0, 0, 1} }

// "Methods"

// AngleTo Returns the angle between this quaternion and to. This is the magnitude of the angle
// you would need to rotate by to get from one to the other.
//
// Note: The magnitude of the floating-point error for this method is abnormally high, so
// methods such as [IsApproximatelyZero] will not work reliably.
func (q Quaternion) AngleTo(other Quaternion) Radians {
	d := q.Dot(other)
	return Acos(d*d*2 - 1)
}

// Dot Returns the dot product of this quaternion and other.
func (q Quaternion) Dot(other Quaternion) Float { return Vector4(q).Dot(Vector4(other)) }

func (q Quaternion) Exp() Quaternion {
	var v = Vector3{q[X], q[Y], q[Z]}
	var theta = v.Length()
	v = v.Normalized()
	if theta < cmpEpsilon || !v.IsNormalized() {
		return Quaternion{0, 0, 0, 1}
	}
	return Quaternion{v[X], v[Y], v[Z], float(theta)}
}

func (q Quaternion) GetAngle() Radians {
	return Acos(q[W]) * 2.0
}

func (q Quaternion) GetAxis() Vector3 {
	if Absf(q[W]) > 1-cmpEpsilon {
		return Vector3{q[X], q[Y], q[Z]}
	}
	var r = (1) / Sqrt(1-q[W]*q[W])
	return Vector3{x * r, y * r, z * r}
}

// EulerAngles returns the quaternion's rotation in the form of Euler angles. The Euler order depends on the order
// parameter, for example using the YXZ convention: since this method decomposes, first Z, then X, and Y last. See
// the EulerOrder enum for possible values. The returned vector contains the rotation angles in the format
// (X angle, Y angle, Z angle).
func (q Quaternion) EulerAngles(order EulerOrder) EulerAngles {
	return q.Basis().EulerAngles(order)
}

// Inverse returns the inverse of the quaternion.
func (q Quaternion) Inverse() Quaternion {
	return Quaternion{-q[X], -q[Y], -q[Z], q[W]}
}

// IsApproximatelyEqual returns true if this quaternion and to are approximately equal, by running
// [IsApproximatelyEqual] on each component.
func (q Quaternion) IsApproximatelyEqual(other Quaternion) bool {
	return Vector4(q).IsApproximatelyEqual(Vector4(other))
}

// IsFinite returns true if this quaternion is finite, by calling [IsFinite] on each component.
func (q Quaternion) IsFinite() bool { return Vector4(q).IsFinite() }

// IsNormalized returns whether the quaternion is normalized or not.
func (q Quaternion) IsNormalized() bool { return IsApproximatelyEqual(q.LengthSquared(), 1) }

// Length returns the length of the quaternion.
func (q Quaternion) Length() Float { return Float(Sqrt(float64(q.LengthSquared()))) }

// LengthSquared returns the length of the quaternion, squared.
func (q Quaternion) LengthSquared() Float { return Float(dot4(q)) }

func (q Quaternion) Log() Quaternion {
	v := q.GetAxis().Mulf(Float(q.GetAngle()))
	return Quaternion{v[X], v[Y], v[Z], 0}
}

// Normalized returns a copy of the quaternion, normalized to unit length.
func (q Quaternion) Normalized() Quaternion { return q.Divf(q.Length()) }

// Slerp returns the result of the spherical linear interpolation between this quaternion
// and to by amount weight.
//
// Note: Both quaternions must be normalized.
func (q Quaternion) Slerp(to Quaternion, weight Float) Quaternion { return q.lerp(to, weight) }

// Slerpni returns the result of the spherical linear interpolation between this quaternion and
// to by amount weight, but without checking if the rotation path is not bigger than 90 degrees.
func (q Quaternion) Slerpni(to Quaternion, weight Float) Quaternion {
	var dot = q.Dot(to)
	if Absf(dot) > 0.9999 {
		return q
	}
	var theta = Acos(dot)
	var sinT = 1.0 / Sin(theta)
	var newFactor = float(Sin(Radians(weight)*theta) * sinT)
	var invFactor = float(Sin(Radians(1.0-weight)*theta) * sinT)
	return Quaternion{
		invFactor*q[X] + newFactor*to[X],
		invFactor*q[Y] + newFactor*to[Y],
		invFactor*q[Z] + newFactor*to[Z],
		invFactor*q[W] + newFactor*to[W],
	}
}

// SphericalCubicInterpolate performs a spherical cubic interpolation between quaternions
// pre_a, this vector, b, and post_b, by the given amount weight.
func (q Quaternion) SphericalCubicInterpolate(b, pre_a, post_b Quaternion, weight Float) Quaternion {
	var from_q = q
	var pre_q = pre_a
	var to_q = b
	var post_q = post_b

	// Align flip phases.
	from_q = from_q.Basis().Quaternion()
	pre_q = pre_q.Basis().Quaternion()
	to_q = to_q.Basis().Quaternion()
	post_q = post_q.Basis().Quaternion()

	// Flip quaternions to shortest path if necessary.
	var flip1 = math.Signbit(float64(from_q.Dot(pre_q)))
	if flip1 {
		pre_q = pre_q.Neg()
	}
	var flip2 = math.Signbit(float64(from_q.Dot(to_q)))
	if flip2 {
		to_q = to_q.Neg()
	}
	if flip2 {
		var flip3 = to_q.Dot(post_q) <= 0
		if flip3 {
			post_q = post_q.Neg()
		}
	} else {
		var flip3 = math.Signbit(float64(to_q.Dot(post_q)))
		if flip3 {
			post_q = post_q.Neg()
		}
	}

	// Calc by Expmap in from_q space.
	var ln_from Quaternion
	var ln_to = (from_q.Inverse().Mul(to_q)).Log()
	var ln_pre = (from_q.Inverse().Mul(pre_q)).Log()
	var ln_post = (from_q.Inverse().Mul(post_q)).Log()
	var ln Quaternion
	ln[X] = CubicInterpolate(ln_from[X], ln_to[X], ln_pre[X], ln_post[X], float(weight))
	ln[Y] = CubicInterpolate(ln_from[Y], ln_to[Y], ln_pre[Y], ln_post[Y], float(weight))
	ln[Z] = CubicInterpolate(ln_from[Z], ln_to[Z], ln_pre[Z], ln_post[Z], float(weight))
	var q1 = from_q.Mul(ln.Exp())

	// Calc by Expmap in to_q space.
	ln_from = (to_q.Inverse().Mul(from_q)).Log()
	ln_to = Quaternion{}
	ln_pre = (to_q.Inverse().Mul(pre_q)).Log()
	ln_post = (to_q.Inverse().Mul(post_q)).Log()
	ln = Quaternion{}
	ln[X] = CubicInterpolate(ln_from[X], ln_to[X], ln_pre[X], ln_post[X], float(weight))
	ln[Y] = CubicInterpolate(ln_from[Y], ln_to[Y], ln_pre[Y], ln_post[Y], float(weight))
	ln[Z] = CubicInterpolate(ln_from[Z], ln_to[Z], ln_pre[Z], ln_post[Z], float(weight))
	var q2 = to_q.Mul(ln.Exp())

	// To cancel error made by Expmap ambiguity, do blending.
	return q1.lerp(q2, weight)
}

// SphericalCubicInterpolateInTime performs a spherical cubic interpolation between quaternions
// pre_a, this vector, b, and post_b, by the given amount weight.
//
// It can perform smoother interpolation than spherical_cubic_interpolate by the time values.
func (q Quaternion) SphericalCubicInterpolateInTime(b, pre_a, post_b Quaternion, weight, b_t, pre_a_t, post_b_t Float) Quaternion {
	var from_q = q
	var pre_q = pre_a
	var to_q = b
	var post_q = post_b

	// Align flip phases.
	from_q = from_q.Basis().Quaternion()
	pre_q = pre_q.Basis().Quaternion()
	to_q = to_q.Basis().Quaternion()
	post_q = post_q.Basis().Quaternion()

	// Flip quaternions to shortest path if necessary.
	var flip1 = math.Signbit(float64(from_q.Dot(pre_q)))
	if flip1 {
		pre_q = pre_q.Neg()
	}
	var flip2 = math.Signbit(float64(from_q.Dot(to_q)))
	if flip2 {
		to_q = to_q.Neg()
	}
	if flip2 {
		var flip3 = to_q.Dot(post_q) <= 0
		if flip3 {
			post_q = post_q.Neg()
		}
	} else {
		var flip3 = math.Signbit(float64(to_q.Dot(post_q)))
		if flip3 {
			post_q = post_q.Neg()
		}
	}

	// Calc by Expmap in from_q space.
	var ln_from Quaternion
	var ln_to = (from_q.Inverse().Mul(to_q)).Log()
	var ln_pre = (from_q.Inverse().Mul(pre_q)).Log()
	var ln_post = (from_q.Inverse().Mul(post_q)).Log()
	var ln Quaternion
	ln[X] = CubicInterpolateInTime(ln_from[X], ln_to[X], ln_pre[X], ln_post[X], float(weight), float(b_t), float(pre_a_t), float(post_b_t))
	ln[Y] = CubicInterpolateInTime(ln_from[Y], ln_to[Y], ln_pre[Y], ln_post[Y], float(weight), float(b_t), float(pre_a_t), float(post_b_t))
	ln[Z] = CubicInterpolateInTime(ln_from[Z], ln_to[Z], ln_pre[Z], ln_post[Z], float(weight), float(b_t), float(pre_a_t), float(post_b_t))
	var q1 = from_q.Mul(ln.Exp())

	// Calc by Expmap in to_q space.
	ln_from = (to_q.Inverse().Mul(from_q)).Log()
	ln_to = Quaternion{}
	ln_pre = (to_q.Inverse().Mul(pre_q)).Log()
	ln_post = (to_q.Inverse().Mul(post_q)).Log()
	ln = Quaternion{}
	ln[X] = CubicInterpolateInTime(ln_from[X], ln_to[X], ln_pre[X], ln_post[X], float(weight), float(b_t), float(pre_a_t), float(post_b_t))
	ln[Y] = CubicInterpolateInTime(ln_from[Y], ln_to[Y], ln_pre[Y], ln_post[Y], float(weight), float(b_t), float(pre_a_t), float(post_b_t))
	ln[Z] = CubicInterpolateInTime(ln_from[Z], ln_to[Z], ln_pre[Z], ln_post[Z], float(weight), float(b_t), float(pre_a_t), float(post_b_t))
	var q2 = to_q.Mul(ln.Exp())

	// To cancel error made by Expmap ambiguity, do blending.
	return q1.lerp(q2, weight)
}

func (q Quaternion) Basis() Basis {
	var d = float(q.LengthSquared())
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

func (q Quaternion) Mul(other Quaternion) Quaternion {
	return Quaternion{
		q[X]*other[W] + q[W]*other[X] + q[Y]*other[Z] - q[Z]*other[Y],
		q[Y]*other[W] + q[W]*other[Y] + q[Z]*other[X] - q[X]*other[Z],
		q[Z]*other[W] + q[W]*other[Z] + q[X]*other[Y] - q[Y]*other[X],
		q[W]*other[W] - q[X]*other[X] - q[Y]*other[Y] - q[Z]*other[Z],
	}

}

func (q Quaternion) Add(other Quaternion) Quaternion {
	return Quaternion{q[X] + other[X], q[Y] + other[Y], q[Z] + other[Z], q[W] + other[W]}
}
func (q Quaternion) Sub(other Quaternion) Quaternion {
	return Quaternion{q[X] - other[X], q[Y] - other[Y], q[Z] - other[Z], q[W] - other[W]}
}
func (q Quaternion) Div(other Quaternion) Quaternion {
	return q.Mul(other.Neg())
}
func (q Quaternion) Addf(other Float) Quaternion {
	return Quaternion{q[X] + float(other), q[Y] + float(other), q[Z] + float(other), q[W] + float(other)}
}
func (q Quaternion) Subf(other Float) Quaternion {
	return Quaternion{q[X] - float(other), q[Y] - float(other), q[Z] - float(other), q[W] - float(other)}
}
func (q Quaternion) Mulf(other Float) Quaternion {
	return Quaternion{q[X] * float(other), q[Y] * float(other), q[Z] * float(other), q[W] * float(other)}
}
func (q Quaternion) Divf(other Float) Quaternion {
	return Quaternion{q[X] / float(other), q[Y] / float(other), q[Z] / float(other), q[W] / float(other)}
}

func (q Quaternion) Neg() Quaternion { return Quaternion{-q[X], -q[Y], -q[Z], -q[W]} }
