// Package Quaternion provides a unit quaternion used for representing 3D rotations.
package Quaternion

import (
	"math"

	"grow.graphics/gd/gdconst"

	"grow.graphics/gd/variant/Angle"
	"grow.graphics/gd/variant/Basis"
	"grow.graphics/gd/variant/Float"
	"grow.graphics/gd/variant/Int"
	"grow.graphics/gd/variant/Vector3"
)

// The Quaternion built-in Variant type is a 4D data structure that represents rotation
// in the form of a Hamilton convention quaternion. Compared to the Basis type which
// can store both rotation and scale, quaternions can only store rotation.
//
// A Quaternion is composed by 4 floating-point components: x, i, j, and k. These components
// are very compact in memory, and because of this some operations are more efficient and
// less likely to cause floating-point errors. Methods such as [Angle], [Axis], and [Slerp]
// are faster than their Basis counterparts.
//
// For a great introduction to quaternions, see this video by 3Blue1Brown. You do not need to
// know the math behind quaternions, as this package provides several helper methods that
// handle it for you. These include [Slerp] and [SphericalCubicInterpolate], as well as [Mul].
//
// Note: Quaternions must be normalized before being used for rotation (see normalized).
//
// Note: Similarly to Vector2 and Vector3, the components of a quaternion use 32-bit precision
// by default, unlike float which is always 64-bit. If double precision is
// needed, compile with the Go build tag 'precision_double'.
type IJKX = struct {
	I Float.X // I component of the quaternion. This is the value along the "imaginary" i axis.
	J Float.X // J component of the quaternion. This is the value along the "imaginary" j axis.
	K Float.X // K component of the quaternion. This is the value along the "imaginary" k axis.
	X Float.X // X component of the quaternion. This is the value along the "real" axis.
}

// Identity quaternion, representing no rotation. This has the same rotation as
// [Basis.Identity].
//
// If a Vector3 is rotated (multiplied) by this quaternion, it does not change.
var Identity = IJKX{X: 1}

// New constructs a Quaternion identical to the [Identity].
func New() IJKX { //gd:Quaternion()
	return Identity
}

// Euler constructs a Quaternion from Euler angles in YXZ rotation order.
func Euler(e Angle.Euler3D) IJKX { //gd:Quaternion.from_euler
	var (
		half_a1 = e.Y * 0.5
		half_a2 = e.X * 0.5
		half_a3 = e.Z * 0.5
	)
	// R = Y(a1).X(a2).Z(a3) convention for Euler angles.
	// Conversion to quaternion as listed in https://ntrs.nasa.gov/archive/nasa/casi.ntrs.nasa.gov/19770024290.pdf (page A-6)
	// a3 is the angle of the first rotation, following the notation in this reference.
	var (
		cos_a1 = Angle.Cos(half_a1)
		sin_a1 = Angle.Sin(half_a1)
		cos_a2 = Angle.Cos(half_a2)
		sin_a2 = Angle.Sin(half_a2)
		cos_a3 = Angle.Cos(half_a3)
		sin_a3 = Angle.Sin(half_a3)
	)
	return IJKX{
		sin_a1*cos_a2*sin_a3 + cos_a1*sin_a2*cos_a3,
		sin_a1*cos_a2*cos_a3 - cos_a1*sin_a2*sin_a3,
		-sin_a1*sin_a2*cos_a3 + cos_a1*cos_a2*sin_a3,
		sin_a1*sin_a2*sin_a3 + cos_a1*cos_a2*cos_a3,
	}
}

// AngleTo Returns the angle between this quaternion and to. This is the magnitude of the angle
// you would need to rotate by to get from one to the other.
//
// Note: The magnitude of the floating-point error for this method is abnormally high, so
// methods such as [IsApproximatelyZero] will not work reliably.
func AngleBetween(q, other IJKX) Angle.Radians { //gd:Quaternion.angle_to
	d := Dot(q, other)
	return Angle.Acos(d*d*2 - 1)
}

// Dot returns the dot product of this quaternion and other.
func Dot(q, other IJKX) Float.X { //gd:Quaternion.dot
	return q.I*other.I + q.J*other.J + q.K*other.K + q.X*other.X
}

// Exponential returns the exponential of this quaternion. The rotation axis of the result is the
// normalized rotation axis of this quaternion, the angle of the result is the length of
// the vector part of this quaternion.
func Exponential(q IJKX) IJKX { //gd:Quaternion.exp
	var v = Vector3.New(q.I, q.J, q.K)
	var theta = Vector3.Length(v)
	v = Vector3.Normalized(v)
	if theta < Float.Epsilon || !Vector3.IsNormalized(v) {
		return IJKX{0, 0, 0, 1}
	}
	return IJKX{v.X, v.Y, v.Z, theta}
}

// AngleInRadians returns the angle of the rotation represented by this quaternion.
//
// Note: The quaternion must be normalized.
func AngleInRadians(q IJKX) Angle.Radians { //gd:Quaternion.get_angle
	return Angle.Acos(q.X) * 2.0
}

func Axis(q IJKX) Vector3.XYZ { //gd:Quaternion.get_axis
	if Float.Abs(q.X) > 1-Float.Epsilon {
		return Vector3.New(q.I, q.J, q.K)
	}
	var r = (1) / Float.Sqrt(1-q.X*q.X)
	return Vector3.New(q.I*r, q.J*r, q.K*r)
}

// EulerAngles returns the quaternion's rotation in the form of Euler angles. The Euler order depends on the order
// parameter, for example using the YXZ convention: since this method decomposes, first Z, then X, and Y last. See
// the EulerOrder enum for possible values. The returned vector contains the rotation angles in the format
// (X angle, Y angle, Z angle).
func EulerAngles(order gdconst.EulerOrder, q IJKX) Angle.Euler3D { //gd:Quaternion.get_euler
	return Basis.AsEulerAngles(AsBasis(q), order)
}

// Inverse returns the inverse of the quaternion.
func Inverse(q IJKX) IJKX { //gd:Quaternion.inverse
	return IJKX{-q.I, -q.J, -q.K, q.X}
}

// IsApproximatelyEqual returns true if this quaternion and to are approximately equal, by running
// [IsApproximatelyEqual] on each component.
func IsApproximatelyEqual(q, other IJKX) bool { //gd:Quaternion.is_equal_approx
	if !Float.IsApproximatelyEqual(q.I, other.I) {
		return false
	}
	if !Float.IsApproximatelyEqual(q.J, other.J) {
		return false
	}
	if !Float.IsApproximatelyEqual(q.K, other.K) {
		return false
	}
	if !Float.IsApproximatelyEqual(q.X, other.X) {
		return false
	}
	return true
}

// IsFinite returns true if this quaternion is finite, by calling [IsFinite] on each component.
func IsFinite(q IJKX) bool { //gd:Quaternion.is_finite
	if !Float.IsFinite(q.I) {
		return false
	}
	if !Float.IsFinite(q.J) {
		return false
	}
	if !Float.IsFinite(q.K) {
		return false
	}
	if !Float.IsFinite(q.X) {
		return false
	}
	return true
}

// IsNormalized returns whether the quaternion is normalized or not.
func IsNormalized(q IJKX) bool { return Float.IsApproximatelyEqual(LengthSquared(q), 1) } //gd:Quaternion.is_normalized

// Length returns the length of the quaternion.
func Length(q IJKX) Float.X { return Float.Sqrt(LengthSquared(q)) } //gd:Quaternion.length

// LengthSquared returns the length of the quaternion, squared.
func LengthSquared(q IJKX) Float.X { //gd:Quaternion.length_squared
	return q.I*q.I + q.J*q.J + q.K*q.K + q.X*q.X
}

// Log returns the logarithm of this quaternion. Multiplies this quaternion's rotation axis by
// its rotation angle, and stores the result in the returned quaternion's vector part
// (x, y, and z). The returned quaternion's real part (w) is always 0.0.
func Log(q IJKX) IJKX { //gd:Quaternion.log
	v := Vector3.MulX(Axis(q), AngleInRadians(q))
	return IJKX{v.X, v.Y, v.Z, 0}
}

// Normalized returns a copy of the quaternion, normalized to unit length.
func Normalized(q IJKX) IJKX { return DivX(q, Length(q)) } //gd:Quaternion.normalized

// Slerp returns the result of the spherical linear interpolation between this quaternion
// and to by amount weight.
//
// Note: Both quaternions must be normalized.
func Slerp[X Float.Any](a, b IJKX, weight X) IJKX { //gd:Quaternion.slerp
	var to1 IJKX
	var (
		cosom, sinom, scale0, scale1 Float.X
		omega                        Angle.Radians
	)
	cosom = LengthSquared(b) // calc cosine
	if cosom < 0.0 {         // adjust signs (if necessary)
		cosom = -cosom
		to1 = Neg(b)
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
	return IJKX{ // calculate final values
		scale0*a.I + scale1*to1.I,
		scale0*a.J + scale1*to1.J,
		scale0*a.K + scale1*to1.K,
		scale0*a.X + scale1*to1.X,
	}
}

// Slerpni returns the result of the spherical linear interpolation between this quaternion and
// to by amount weight, but without checking if the rotation path is not bigger than 90 degrees.
func Slerpni[X Float.Any](a, b IJKX, weight X) IJKX { //gd:Quaternion.slerpni
	var dot = Dot(a, b)
	if Float.Abs(dot) > 0.9999 {
		return a
	}
	var theta = Angle.Acos(dot)
	var sinT = 1.0 / Angle.Sin(theta)
	var newFactor = Angle.Sin(Angle.Radians(weight)*theta) * sinT
	var invFactor = Angle.Sin(Angle.Radians(1.0-weight)*theta) * sinT
	return IJKX{
		invFactor*a.I + newFactor*b.I,
		invFactor*a.J + newFactor*b.J,
		invFactor*a.K + newFactor*b.K,
		invFactor*a.X + newFactor*b.X,
	}
}

// SphericalCubicInterpolate performs a spherical cubic interpolation between quaternions
// pre_a, this vector, b, and post_b, by the given amount weight.
func SphericalCubicInterpolate[X Float.Any](a, b, pre_a, post_b IJKX, weight X) IJKX { //gd:Quaternion.spherical_cubic_interpolate
	var (
		from_q = a
		pre_q  = pre_a
		to_q   = b
		post_q = post_b
	)
	// Align flip phases.
	from_q = Basis.AsQuaternion(AsBasis(from_q))
	pre_q = Basis.AsQuaternion(AsBasis(pre_q))
	to_q = Basis.AsQuaternion(AsBasis(to_q))
	post_q = Basis.AsQuaternion(AsBasis(post_q))
	// Flip quaternions to shortest path if necessary.
	var flip1 = math.Signbit(float64(Dot(from_q, pre_q)))
	if flip1 {
		pre_q = Neg(pre_q)
	}
	var flip2 = math.Signbit(float64(Dot(from_q, to_q)))
	if flip2 {
		to_q = Neg(to_q)
	}
	if flip2 {
		var flip3 = Dot(to_q, post_q) <= 0
		if flip3 {
			post_q = Neg(post_q)
		}
	} else {
		var flip3 = math.Signbit(float64(Dot(to_q, post_q)))
		if flip3 {
			post_q = Neg(post_q)
		}
	}
	// Calc by Expmap in from_q space.
	var ln_from IJKX
	var ln_to = Log(Mul(Inverse(from_q), to_q))
	var ln_pre = Log(Mul(Inverse(from_q), pre_q))
	var ln_post = Log(Mul(Inverse(from_q), post_q))
	var ln IJKX
	ln.I = Float.CubicInterpolate(ln_from.I, ln_to.I, ln_pre.I, ln_post.I, Float.X(weight))
	ln.J = Float.CubicInterpolate(ln_from.J, ln_to.J, ln_pre.J, ln_post.J, Float.X(weight))
	ln.K = Float.CubicInterpolate(ln_from.K, ln_to.K, ln_pre.K, ln_post.K, Float.X(weight))
	var q1 = Mul(from_q, Exponential(ln))
	// Calc by Expmap in to_q space.
	ln_from = Log(Mul(Inverse(to_q), from_q))
	ln_to = IJKX{}
	ln_pre = Log(Mul(Inverse(to_q), pre_q))
	ln_post = Log(Mul(Inverse(to_q), post_q))
	ln = IJKX{}
	ln.I = Float.CubicInterpolate(ln_from.I, ln_to.I, ln_pre.I, ln_post.I, Float.X(weight))
	ln.J = Float.CubicInterpolate(ln_from.J, ln_to.J, ln_pre.J, ln_post.J, Float.X(weight))
	ln.K = Float.CubicInterpolate(ln_from.K, ln_to.K, ln_pre.K, ln_post.K, Float.X(weight))
	var q2 = Mul(to_q, Exponential(ln))
	// To cancel error made by Expmap ambiguity, do blending.
	return Slerp(q1, q2, weight)
}

// SphericalCubicInterpolateInTime performs a spherical cubic interpolation between quaternions
// pre_a, this vector, b, and post_b, by the given amount weight.
//
// It can perform smoother interpolation than spherical_cubic_interpolate by the time values.
func SphericalCubicInterpolateInTime[X Float.Any](a, b, pre_a, post_b IJKX, weight, b_t, pre_a_t, post_b_t X) IJKX { //gd:Quaternion.spherical_cubic_interpolate_in_time
	var (
		from_q = a
		pre_q  = pre_a
		to_q   = b
		post_q = post_b
	)
	// Align flip phases.
	from_q = Basis.AsQuaternion(AsBasis(from_q))
	pre_q = Basis.AsQuaternion(AsBasis(pre_q))
	to_q = Basis.AsQuaternion(AsBasis(to_q))
	post_q = Basis.AsQuaternion(AsBasis(post_q))
	// Flip quaternions to shortest path if necessary.
	var flip1 = math.Signbit(float64(Dot(from_q, pre_q)))
	if flip1 {
		pre_q = Neg(pre_q)
	}
	var flip2 = math.Signbit(float64(Dot(from_q, to_q)))
	if flip2 {
		to_q = Neg(to_q)
	}
	if flip2 {
		var flip3 = Dot(to_q, post_q) <= 0
		if flip3 {
			post_q = Neg(post_q)
		}
	} else {
		var flip3 = math.Signbit(float64(Dot(to_q, post_q)))
		if flip3 {
			post_q = Neg(post_q)
		}
	}
	// Calc by Expmap in from_q space.
	var ln_from IJKX
	var ln_to = Log(Mul(Inverse(from_q), to_q))
	var ln_pre = Log(Mul(Inverse(from_q), pre_q))
	var ln_post = Log(Mul(Inverse(from_q), post_q))
	var ln IJKX
	ln.I = Float.CubicInterpolateInTime(ln_from.I, ln_to.I, ln_pre.I, ln_post.I, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t))
	ln.J = Float.CubicInterpolateInTime(ln_from.J, ln_to.J, ln_pre.J, ln_post.J, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t))
	ln.K = Float.CubicInterpolateInTime(ln_from.K, ln_to.K, ln_pre.K, ln_post.K, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t))
	var q1 = Mul(from_q, Exponential(ln))

	// Calc by Expmap in to_q space.
	ln_from = Log(Mul(Inverse(to_q), from_q))
	ln_to = IJKX{}
	ln_pre = Log(Mul(Inverse(to_q), pre_q))
	ln_post = Log(Mul(Inverse(to_q), post_q))
	ln = IJKX{}
	ln.I = Float.CubicInterpolateInTime(ln_from.I, ln_to.I, ln_pre.I, ln_post.I, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t))
	ln.J = Float.CubicInterpolateInTime(ln_from.J, ln_to.J, ln_pre.J, ln_post.J, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t))
	ln.K = Float.CubicInterpolateInTime(ln_from.K, ln_to.K, ln_pre.K, ln_post.K, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t))
	var q2 = Mul(to_q, Exponential(ln))

	// To cancel error made by Expmap ambiguity, do blending.
	return Slerp(q1, q2, weight)
}

func AsBasis(q IJKX) Basis.XYZ { //gd:Basis(Quaternion)
	var d = LengthSquared(q)
	var s = 2.0 / d
	var xs, ys, zs = q.I * s, q.J * s, q.K * s
	var wx, wy, wz = q.X * xs, q.X * ys, q.X * zs
	var xx, xy, xz = q.I * xs, q.I * ys, q.I * zs
	var yy, yz, zz = q.J * ys, q.J * zs, q.K * zs
	return Basis.XYZ{
		X: Vector3.New(1.0-(yy+zz), xy-wz, xz+wy),
		Y: Vector3.New(xy+wz, 1.0-(xx+zz), yz-wx),
		Z: Vector3.New(xz-wy, yz+wx, 1.0-(xx+yy)),
	}
}

func Mul(a, b IJKX) IJKX { //gd:Quaternion*(right:Quaternion)
	return IJKX{
		a.I*b.X + a.X*b.I + a.J*b.K - a.K*b.J,
		a.J*b.X + a.X*b.J + a.K*b.I - a.I*b.K,
		a.K*b.X + a.X*b.K + a.I*b.J - a.J*b.I,
		a.X*b.X - a.I*b.I - a.J*b.J - a.K*b.K,
	}
}

func Add(a, b IJKX) IJKX { return IJKX{a.I + b.I, a.J + b.J, a.K + b.K, a.X + b.X} } //gd:Quaternion+(right:Vector4i)
func Sub(a, b IJKX) IJKX { return IJKX{a.I - b.I, a.J - b.J, a.K - b.K, a.X - b.X} } //gd:Quaternion-(right:Vector4i)
func Div(a, b IJKX) IJKX { return Mul(a, Neg(b)) }                                   //gd:Quaternion/(right:Vector4i)

func AddX[X Int.Any | Float.Any](a IJKX, b X) IJKX { //gd:Quaternion+(right:float)
	return IJKX{a.I + Float.X(b), a.J + Float.X(b), a.K + Float.X(b), a.X + Float.X(b)}
}
func SubX[X Int.Any | Float.Any](a IJKX, b X) IJKX { //gd:Quaternion-(right:float)
	return IJKX{a.I - Float.X(b), a.J - Float.X(b), a.K - Float.X(b), a.X - Float.X(b)}
}
func MulX[X Int.Any | Float.Any](a IJKX, b X) IJKX { //gd:Quaternion*(right:float)
	return IJKX{a.I * Float.X(b), a.J * Float.X(b), a.K * Float.X(b), a.X * Float.X(b)}
}
func DivX[X Int.Any | Float.Any](a IJKX, b X) IJKX { //gd:Quaternion/(right:float)
	return IJKX{a.I / Float.X(b), a.J / Float.X(b), a.K / Float.X(b), a.X / Float.X(b)}
}

func Neg(v IJKX) IJKX { return IJKX{-v.I, -v.J, -v.K, -v.X} } //gd:Quaternion-(unary)

func AsArray(vec IJKX) [4]Float.X { return [4]Float.X{vec.I, vec.J, vec.K, vec.X} }
