//go:build !generate

package gd

import "math"

// Vector3 is a 3-element structure that can be used to represent 3D coordinates or any other triplet of numeric values.
//
// It uses floating-point coordinates. By default, these floating-point values use 32-bit precision, unlike [Float] which
// is always 64-bit. If double precision is needed, compile the engine with the option precision=double.
//
// See Vector3i for its integer counterpart.
//
// Note: In a boolean context, a Vector3 will evaluate to false if it's equal to Vector3{0, 0, 0}. Otherwise, a [Vector3]
// will always evaluate to true.
type Vector3 [3]float32

func (v Vector3) X() Float      { return Float(v[X]) }
func (v Vector3) Y() Float      { return Float(v[Y]) }
func (v Vector3) Z() Float      { return Float(v[Z]) }
func (v *Vector3) SetX(x Float) { v[X] = float32(x) }
func (v *Vector3) SetY(y Float) { v[Y] = float32(y) }
func (v *Vector3) SetZ(z Float) { v[Z] = float32(z) }

// "Constants"
func (Vector3) ZERO() Vector3 { return Vector3{0, 0, 0} }
func (Vector3) ONE() Vector3  { return Vector3{1, 1, 1} }
func (Vector3) INF() Vector3 {
	return Vector3{float32(math.Inf(1)), float32(math.Inf(1)), float32(math.Inf(1))}
}
func (Vector3) LEFT() Vector3         { return Vector3{-1, 0, 0} }
func (Vector3) RIGHT() Vector3        { return Vector3{1, 0, 0} }
func (Vector3) UP() Vector3           { return Vector3{0, 1, 0} }
func (Vector3) DOWN() Vector3         { return Vector3{0, -1, 0} }
func (Vector3) FORWARD() Vector3      { return Vector3{0, 0, -1} }
func (Vector3) BACK() Vector3         { return Vector3{0, 0, 1} }
func (Vector3) MODEL_LEFT() Vector3   { return Vector3{1, 0, 0} }
func (Vector3) MODEL_RIGHT() Vector3  { return Vector3{-1, 0, 0} }
func (Vector3) MODEL_TOP() Vector3    { return Vector3{0, 1, 0} }
func (Vector3) MODEL_BOTTOM() Vector3 { return Vector3{0, -1, 0} }
func (Vector3) MODEL_FRONT() Vector3  { return Vector3{0, 0, -1} }
func (Vector3) MODEL_REAR() Vector3   { return Vector3{0, 0, 1} }

func (v Vector3) Vector3i() Vector3i { return Vector3i{int32(v[X]), int32(v[Y]), int32(v[Z])} }

// Abs returns a new vector with all components in absolute values (i.e. positive).
func (v Vector3) Abs() Vector3 { return v.abs() }

// AngleTo returns the unsigned minimum angle to the given vector, in radians.
func (v Vector3) AngleTo(to Vector3) Radians {
	return Atan2(v.Cross(to).Length(), v.Dot(to))
}

// BezierDerivative returns the derivative at the given t on the Bézier curve
// defined by this vector and the given control_1, control_2, and end points.
func (v Vector3) BezierDerivative(control1, control2, end Vector3, t Float) Vector3 {
	return Vector3{
		BezierDerivative(v[X], control1[X], control2[X], end[X], float32(t)),
		BezierDerivative(v[Y], control1[Y], control2[Y], end[Y], float32(t)),
		BezierDerivative(v[Z], control1[Z], control2[Z], end[Z], float32(t)),
	}
}

// BezierInterpolate returns the point at the given t on the Bézier curve defined by
// this vector and the given control_1, control_2, and end points.
func (v Vector3) BezierInterpolate(control1, control2, end Vector3, t Float) Vector3 {
	return Vector3{
		BezierInterpolate(v[X], control1[X], control2[X], end[X], float32(t)),
		BezierInterpolate(v[Y], control1[Y], control2[Y], end[Y], float32(t)),
		BezierInterpolate(v[Z], control1[Z], control2[Z], end[Z], float32(t)),
	}
}

// Bounce returns the vector "bounced off" from a plane defined by the given normal.
func (v Vector3) Bounce(n Vector3) Vector3 {
	return Vector3.Sub(Vector3{}, v.Reflect(n))
}

// Ceil returns a new vector with all components rounded up (towards positive infinity).
func (v Vector3) Ceil() Vector3 { return v.ceil() }

// Clamp returns a new vector with all components clamped between the components of min and max,
// by running [Clampf] on each component.
func (v Vector3) Clamp(min, max Vector3) Vector3 {
	return Vector3{
		Clampf(v[X], min[X], max[X]),
		Clampf(v[Y], min[Y], max[Y]),
		Clampf(v[Z], min[Z], max[Z]),
	}
}

// Cross returns the cross product of this vector and with.
func (v Vector3) Cross(with Vector3) Vector3 {
	return Vector3{
		v[Y]*with[Z] - v[Z]*with[Y],
		v[Z]*with[X] - v[X]*with[Z],
		v[X]*with[Y] - v[Y]*with[X],
	}
}

// CubicInterpolate performs a cubic interpolation between this vector and b using pre_a and post_b as handles
// and returns the result at position weight. weight is on the range of 0.0 to 1.0, representing the amount of
// interpolation.
func (v Vector3) CubicInterpolate(b, preA, postB Vector3, weight Float) Vector3 {
	return Vector3{
		CubicInterpolate(v[X], b[X], preA[X], postB[X], float32(weight)),
		CubicInterpolate(v[Y], b[Y], preA[Y], postB[Y], float32(weight)),
		CubicInterpolate(v[Z], b[Z], preA[Z], postB[Z], float32(weight)),
	}
}

// CubicInterpolateInTime performs a cubic interpolation between this vector and b using pre_a and post_b as handles
// and returns the result at position weight. weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
//
// It can perform smoother interpolation than [Vector3.CubicInterpolate] by the time values.
func (v Vector3) CubicInterpolateInTime(b, preA, postB Vector3, weight, b_t, pre_a_t, post_b_t Float) Vector3 {
	return Vector3{
		CubicInterpolateInTime(v[X], b[X], preA[X], postB[X], float32(weight), float32(b_t), float32(pre_a_t), float32(post_b_t)),
		CubicInterpolateInTime(v[Y], b[Y], preA[Y], postB[Y], float32(weight), float32(b_t), float32(pre_a_t), float32(post_b_t)),
		CubicInterpolateInTime(v[Z], b[Z], preA[Z], postB[Z], float32(weight), float32(b_t), float32(pre_a_t), float32(post_b_t)),
	}
}

// DirectionTo returns the normalized vector pointing from this vector to to. This is equivalent to using (b - a).Normalized().
func (v Vector3) DirectionTo(to Vector3) Vector3 {
	return Vector3.Sub(to, v).Normalized()
}

// DistanceSquaredTo returns the squared distance between this vector and to.
//
// This method runs faster than [Vector3.DistanceTo], so prefer it if you need to compare vectors or need the squared distance for
// some formula.
func (v Vector3) DistanceSquaredTo(to Vector3) Float {
	return Float(v.Sub(to).LengthSquared())
}

// DistanceTo returns the distance between this vector and to.
func (v Vector3) DistanceTo(to Vector3) Float {
	return Float(v.Sub(to).Length())
}

// Dot returns the dot product of this vector and with. This can be used to compare the angle between two vectors. For example, this
// can be used to determine whether an enemy is facing the player.
//
// The dot product will be 0 for a straight angle (90 degrees), greater than 0 for angles narrower than 90 degrees and lower than 0
// for angles wider than 90 degrees.
//
// When using unit (normalized) vectors, the result will always be between -1.0 (180 degree angle) when the vectors are facing opposite
// directions, and 1.0 (0 degree angle) when the vectors are aligned.
//
// Note: a.Dot(b) is equivalent to b.Dot(a).
func (v Vector3) Dot(with Vector3) Float {
	return Float(v[X]*with[X] + v[Y]*with[Y] + v[Z]*with[Z])
}

// Floor returns a new vector with all components rounded down (towards negative infinity).
func (v Vector3) Floor() Vector3 { return v.floor() }

// Inverse returns the inverse of the vector. This is the same as Vector3{1/v[X],1/v[Y],1/v[Z]}.
func (v Vector3) Inverse() Vector3 {
	return Vector3{1 / v[X], 1 / v[Y], 1 / v[Z]}
}

// IsApproximatelyEqual returns true if this vector and other are approximately equal, by running [IsApproximatelyEqual] on each component.
func (v Vector3) IsApproximatelyEqual(other Vector3) bool {
	if !IsApproximatelyEqual(v[X], other[X]) {
		return false
	}
	if !IsApproximatelyEqual(v[Y], other[Y]) {
		return false
	}
	if !IsApproximatelyEqual(v[Z], other[Z]) {
		return false
	}
	return true
}

// IsFinite returns true if this vector's values are finite, by running [IsFinite] on each component.
func (v Vector3) IsFinite() bool {
	if !IsFinite(v[X]) {
		return false
	}
	if !IsFinite(v[Y]) {
		return false
	}
	if !IsFinite(v[Z]) {
		return false
	}
	return true
}

// IsNormalized Returns true if the vector is normalized, i.e. its length is approximately equal to 1.
func (v Vector3) IsNormalized() bool { return IsApproximatelyEqual(v.LengthSquared(), 1) }

// IsApproximatelyZero returns true if this vector is approximately equal to zero, by running
// [IsApproximatelyZero] on each component.
func (v Vector3) IsApproximatelyZero() bool {
	if !IsApproximatelyZero(v[X]) {
		return false
	}
	if !IsApproximatelyZero(v[Y]) {
		return false
	}
	if !IsApproximatelyZero(v[Z]) {
		return false
	}
	return true
}

// Length the length (magnitude) of this vector.
func (v Vector3) Length() Float {
	return Float(math.Sqrt(float64(v.LengthSquared())))
}

// LengthSquared returns the squared length (squared magnitude) of this vector.
//
// This method runs faster than length, so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func (v Vector3) LengthSquared() Float {
	return Float(Float(v[x])*Float(v[x]) + Float(v[y])*Float(v[y]) + Float(v[z])*Float(v[z]))
}

// Lerp returns the result of the linear interpolation between this vector and to by amount weight.
// weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
func (v Vector3) Lerp(to Vector3, weight Float) Vector3 { return v.lerp(to, weight) }

// LimitLength returns the vector with a maximum length by limiting its length to length.
func (v Vector3) LimitLength(length Float) Vector3 {
	var l = v.Length()
	if l > 0 && length < l {
		v = v.Divf(l)
		v = v.Mulf(length)
	}
	return v
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all components are
// equal, this method returns [X].
func (v Vector3) MaxAxis() Axis {
	if v[X] < v[Y] {
		if v[Y] < v[Z] {
			return Z
		}
		return Y
	}
	if v[X] < v[Z] {
		return Z
	}
	return X
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all components are
// equal, this method returns [Z].
func (v Vector3) MinAxis() Axis {
	if v[X] < v[Y] {
		if v[Y] < v[Z] {
			return X
		}
		return Z
	}
	if v[X] < v[Z] {
		return Y
	}
	return Z
}

// MoveToward returns a new vector moved toward to by the fixed delta amount. Will not go past the final value.
func (v Vector3) MoveToward(to Vector3, delta Float) Vector3 {
	var vd = to.Sub(v)
	var l = vd.Length()
	if l <= delta || l < cmpEpsilon {
		return to
	}
	return v.Add(vd.Divf(l).Mulf(delta))
}

// Normalized returns the result of scaling the vector to unit length. Equivalent to v / v.Length().
// See also [Vector3.IsNormalized].
//
// Note: This function may return incorrect values if the input vector length is near zero.
func (v Vector3) Normalized() Vector3 {
	var l = v.Length()
	if l == 0 {
		return Vector3{}
	}
	return v.Divf(l)
}

// OctahedronDecode returns the [Vector3] from an octahedral-compressed form created using
// OctahedronEncode (stored as a [Vector2]).
func (v Vector3) OctahedronDecode(uv Vector2) Vector3 {
	var f = Vector3{uv[X]*2.0 - 1.0, uv[Y]*2.0 - 1.0}
	var n = Vector3{f[X], f[Y], 1.0 - Absf(f[X]) - Absf(f[Y])}
	var t = Clampf(-n[Z], 0.0, 1.0)
	if uv[X] >= 0 {
		n[X] += -t
	} else {
		n[X] += t
	}
	if uv[Y] >= 0 {
		n[Y] += -t
	} else {
		n[Y] += t
	}
	return n.Normalized()
}

// OctahedronEncode returns the octahedral-encoded (oct32) form of this [Vector3] as a [Vector2].
// Since a [Vector2] occupies 1/3 less memory compared to [Vector3], this form of compression can
// be used to pass greater amounts of normalized [Vector3]s without increasing storage or memory
// requirements. See also [Vector3.OctahedronDecode].
//
// Note: OctahedronEncode can only be used for normalized vectors. OctahedronEncode does not check
// whether this [Vector3] is normalized, and will return a value that does not decompress to the
// original value if the [Vector3] is not normalized.
//
// Note: Octahedral compression is lossy, although visual differences are rarely perceptible in
// real world scenarios.
func (v Vector3) OctahedronEncode() Vector2 {
	v = v.Divf(Float(Absf(v[X]) + Absf(v[Y]) + Absf(v[Z])))
	var o Vector2
	if v[Z] >= 0.0 {
		o[X] = v[X]
		o[Y] = v[Y]
	} else {
		if v[X] > 0 {
			o[X] = 1.0 - Absf(v[Y])
		} else {
			o[X] = -(1.0 - Absf(v[Y]))
		}
		if v[Y] > 0 {
			o[Y] = Absf(v[X])
		} else {
			o[Y] = -(1.0 - Absf(v[X]))
		}
	}
	o[X] = o[X]*0.5 + 0.5
	o[Y] = o[Y]*0.5 + 0.5
	return o
}

// Outer returns the outer product with with.
func (v Vector3) Outer(with Vector3) Basis {
	return Basis{
		Vector3{v[X] * with[X], v[X] * with[Y], v[X] * with[Z]},
		Vector3{y * with[X], y * with[Y], y * with[Z]},
		Vector3{z * with[X], z * with[Y], z * with[Z]},
	}
}

// Posmodf returns a vector composed of the Fposmod of this vector's components and mod.
func (v Vector3) Posmodf(mod Float) Vector3 {
	return Vector3{Fposmod(v[X], float32(mod)), Fposmod(v[Y], float32(mod)), Fposmod(v[Z], float32(mod))}
}

// Posmodv returns a vector composed of the Fposmod of this vector's components and mod.
func (v Vector3) Posmodv(mod Vector3) Vector3 {
	return Vector3{Fposmod(v[X], mod[X]), Fposmod(v[Y], mod[Y]), Fposmod(v[Z], mod[Z])}
}

// Project returns the result of projecting the vector onto the given vector b.
func (v Vector3) Project(b Vector3) Vector3 {
	return b.Mulf(v.Dot(b) / b.LengthSquared())
}

// Rotated returns the result of rotating this vector around a given axis by angle (in radians).
// The axis must be a normalized vector.
func (v Vector3) Rotated(axis Vector3, angle Radians) Vector3 {
	return NewRotationAroundAxis(axis, angle).Transform(v)
}

// Round returns a new vector with all components rounded to the nearest integer, with halfway cases
// rounded away from zero.
func (v Vector3) Round() Vector3 { return v.round() }

// SignedAngleTo returns the signed angle to the given vector, in radians. The sign of the angle is
// positive in a counter-clockwise direction and negative in a clockwise direction when viewed from
// the side specified by the axis.
func (v Vector3) SignedAngleTo(to Vector3, axis Vector3) Radians {
	var cross_to = v.Cross(to)
	var unsigned_angle = Atan2(cross_to.Length(), v.Dot(to))
	var sign = cross_to.Dot(axis)
	if sign < 0 {
		return -unsigned_angle
	}
	return unsigned_angle
}

// Slerp returns the result of spherical linear interpolation between this vector and to, by amount weight.
// weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
//
// This method also handles interpolating the lengths if the input vectors have different lengths. For the
// special case of one or both input vectors having zero length, this method behaves like lerp.
func (v Vector3) Slerp(to Vector3, weight Float) Vector3 {
	// This method seems more complicated than it really is, since we write out
	// the internals of some methods for efficiency (mainly, checking length).
	var start_length_sq = v.LengthSquared()
	var end_length_sq = to.LengthSquared()
	if start_length_sq == 0.0 || end_length_sq == 0.0 {
		// Zero length vectors have no angle, so the best we can do is either lerp or throw an error.
		return v.lerp(to, weight)
	}
	var axis Vector3 = v.Cross(to)
	var axis_length_sq = axis.LengthSquared()
	if axis_length_sq == 0.0 {
		// Colinear vectors have no rotation axis or angle between them, so the best we can do is lerp.
		return v.lerp(to, weight)
	}
	axis = axis.Divf(Sqrt(axis_length_sq))
	var start_length = Sqrt(start_length_sq)
	var result_length = Lerpf(start_length, Sqrt(end_length_sq), weight)
	var angle = v.AngleTo(to)
	return v.Rotated(axis, angle*Radians(weight)).Mulf(result_length / start_length)
}

// Slide returns a new vector slid along a plane defined by the given normal.
func (v Vector3) Slide(n Vector3) Vector3 {
	return Vector3.Sub(v, n.Mulf(v.Dot(n)))
}

// Snapped returns a new vector with each component snapped to the nearest multiple of the corresponding component
// in step. This can also be used to round the components to an arbitrary number of decimals.
func (v Vector3) Snapped(step Vector3) Vector3 {
	return Vector3{
		Snappedf(v[X], step[X]),
		Snappedf(v[Y], step[Y]),
		Snappedf(v[Z], step[Z]),
	}
}

// Reflect returns the result of reflecting the vector from a plane defined by the given normal n.
func (v Vector3) Reflect(n Vector3) Vector3 {
	return Vector3.Sub(n.Mulf(2.0*v.Dot(n)), v)
}

func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{v[x] + other[x], v[y] + other[y], v[z] + other[z]}
}
func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{v[x] - other[x], v[y] - other[y], v[z] - other[z]}
}
func (v Vector3) Mul(other Vector3) Vector3 {
	return Vector3{v[x] * other[x], v[y] * other[y], v[z] * other[z]}
}
func (v Vector3) Div(other Vector3) Vector3 {
	return Vector3{v[x] / other[x], v[y] / other[y], v[z] / other[z]}
}
func (v Vector3) Addf(other Float) Vector3 {
	return Vector3{v[x] + float32(other), v[y] + float32(other), v[z] + float32(other)}
}
func (v Vector3) Subf(other Float) Vector3 {
	return Vector3{v[x] - float32(other), v[y] - float32(other), v[z] - float32(other)}
}
func (v Vector3) Mulf(other Float) Vector3 {
	return Vector3{v[x] * float32(other), v[y] * float32(other), v[z] * float32(other)}
}
func (v Vector3) Divf(other Float) Vector3 {
	return Vector3{v[x] / float32(other), v[y] / float32(other), v[z] / float32(other)}
}
func (v Vector3) Neg() Vector3 { return Vector3{-v[x], -v[y], -v[z]} }

func (v Vector3) Transform(t Transform3D) Vector3 {
	return Vector3{
		float32(t.Basis[0].Dot(v)) + t.Origin[X],
		float32(t.Basis[1].Dot(v)) + t.Origin[Y],
		float32(t.Basis[2].Dot(v)) + t.Origin[Z],
	}
}
