// Package Vector3 provides a 3D vector using floating-point coordinates.
package Vector3

import (
	"math"

	"grow.graphics/gd/gdmaths/Angle"
	"grow.graphics/gd/gdmaths/Float"
	"grow.graphics/gd/gdmaths/Int"
	"grow.graphics/gd/gdmaths/Vector2"
)

// XYZ is a 3-element structure that can be used to represent 3D coordinates
// or any other triplet of numeric values.
//
// It uses floating-point coordinates. By default, these floating-point values use
// 32-bit precision, unlike float which is always 64-bit. If double precision is
// needed, compile with the Go build tag 'precision_double'.
//
// See [Vector3i.XYZ] for its integer counterpart.
type XYZ = struct {
	X Float.X // The vector's X component.
	Y Float.X // The vector's Y component.
	Z Float.X // The vector's Z component.
}

type Axis int

const (
	X Axis = iota // Enumerated value for the X axis. Returned by [MaxAxis] and [MinAxis].
	Y             // Enumerated value for the Y axis. Returned by [MaxAxis] and [MinAxis].
	Z             // Enumerated value for the Z axis. Returned by [MaxAxis] and [MinAxis].
)

var (
	Zero  = XYZ{0, 0, 0}                         // Zero vector, a vector with all components set to 0.
	One   = XYZ{1, 1, 1}                         // One vector, a vector with all components set to 1.
	Inf   = XYZ{Float.Inf, Float.Inf, Float.Inf} // Inf vector, a vector with all components set to positive infinity.
	Left  = XYZ{-1, 0, 0}                        // Left unit vector. Represents the local direction of left, and the global direction of west.
	Right = XYZ{1, 0, 0}                         // Right unit vector. Represents the local direction of right, and the global direction of east.
	Up    = XYZ{0, -1, 0}                        // Up unit vector.
	Down  = XYZ{0, 1, 0}                         // Down unit vector.

	// Forward unit vector. Represents the local direction of forward, and the global direction
	// of north. Keep in mind that the forward direction for lights, cameras, etc is different
	// from 3D assets like characters, which face towards the camera by convention. Use [ModelFront]
	// and similar constants when working in 3D asset space.
	Forward = XYZ{0, 0, -1}

	Back = XYZ{0, 0, 1} // Back unit vector. Represents the local direction of back, and the global direction of south.

	ModelLeft   = XYZ{1, 0, 0}  // Unit vector pointing towards the left side of imported 3D assets.
	ModelRight  = XYZ{-1, 0, 0} // Unit vector pointing towards the right side of imported 3D assets.
	ModelTop    = XYZ{0, 1, 0}  // Unit vector pointing towards the top side of imported 3D assets.
	ModelBottom = XYZ{0, -1, 0} // Unit vector pointing towards the bottom side of imported 3D assets.
	ModelFront  = XYZ{0, 0, 1}  // Unit vector pointing towards the front side of imported 3D assets.
	ModelRear   = XYZ{0, 0, -1} // Unit vector pointing towards the rear side of imported 3D assets.
)

// New returns a [XYZ] with the given components.
func New[X Float.Any](x, y, z X) XYZ { //gd:Vector3(x:float,y:float,z:float)
	return XYZ{Float.X(x), Float.X(y), Float.X(z)}
}

// Abs returns a new vector with all components in absolute values (i.e. positive).
func Abs(v XYZ) XYZ { //gd:Vector3.abs
	return XYZ{
		Float.Abs(v.X), Float.Abs(v.Y), Float.Abs(v.Z),
	}
}

// AngleBetween returns the unsigned minimum angle to the given vector, in radians.
func AngleBetween(a, b XYZ) Angle.Radians { //gd:Vector3.angle_to
	return Float.Atan2(Length(Cross(a, b)), Dot(a, b))
}

// BezierDerivative returns the derivative at the given t on the Bézier curve
// defined by this vector and the given control_1, control_2, and end points.
func BezierDerivative[X Float.Any](v, control1, control2, end XYZ, t X) XYZ { //gd:Vector3.bezier_derivative
	return XYZ{
		Float.BezierDerivative(v.X, control1.X, control2.X, end.X, Float.X(t)),
		Float.BezierDerivative(v.Y, control1.Y, control2.Y, end.Y, Float.X(t)),
		Float.BezierDerivative(v.Z, control1.Z, control2.Z, end.Z, Float.X(t)),
	}
}

// BezierInterpolate returns the point at the given t on the Bézier curve defined by
// this vector and the given control_1, control_2, and end points.
func BezierInterpolate[X Float.Any](v, control1, control2, end XYZ, t X) XYZ { //gd:Vector3.bezier_interpolate
	return XYZ{
		Float.BezierInterpolate(v.X, control1.X, control2.X, end.X, Float.X(t)),
		Float.BezierInterpolate(v.Y, control1.Y, control2.Y, end.Y, Float.X(t)),
		Float.BezierInterpolate(v.Z, control1.Z, control2.Z, end.Z, Float.X(t)),
	}
}

// Bounce returns the vector "bounced off" from a plane defined by the given normal.
func Bounce(v, n XYZ) XYZ { //gd:Vector3.bounce
	return Sub(Zero, Reflect(v, n))
}

// Ceil returns a new vector with all components rounded up (towards positive infinity).
func Ceil(v XYZ) XYZ { //gd:Vector3.ceil
	return XYZ{
		Float.Ceil(v.X), Float.Ceil(v.Y), Float.Ceil(v.Z),
	}
}

// Clamp returns a new vector with all components clamped between the components of min and max,
// by running [Float.Clamp] on each component.
func Clamp(v, min, max XYZ) XYZ { //gd:Vector3.clamp
	return XYZ{
		Float.Clamp(v.X, min.X, max.X),
		Float.Clamp(v.Y, min.Y, max.Y),
		Float.Clamp(v.Z, min.Z, max.Z),
	}
}

// Cross returns the cross product of this vector and with.
func Cross(v, with XYZ) XYZ { //gd:Vector3.cross
	return XYZ{
		v.Y*with.Z - v.Z*with.Y,
		v.Z*with.X - v.X*with.Z,
		v.X*with.Y - v.Y*with.X,
	}
}

// CubicInterpolate performs a cubic interpolation between this vector and b using pre_a and post_b as handles
// and returns the result at position weight. weight is on the range of 0.0 to 1.0, representing the amount of
// interpolation.
func CubicInterpolate[X Float.Any](v, b, preA, postB XYZ, weight X) XYZ { //gd:Vector3.cubic_interpolate
	return XYZ{
		Float.CubicInterpolate(v.X, b.X, preA.X, postB.X, Float.X(weight)),
		Float.CubicInterpolate(v.Y, b.Y, preA.Y, postB.Y, Float.X(weight)),
		Float.CubicInterpolate(v.Z, b.Z, preA.Z, postB.Z, Float.X(weight)),
	}
}

// CubicInterpolateInTime performs a cubic interpolation between this vector and b using pre_a and post_b as handles
// and returns the result at position weight. weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
//
// It can perform smoother interpolation than [Vector3.CubicInterpolate] by the time values.
func CubicInterpolateInTime[X Float.Any](v, b, preA, postB XYZ, weight, b_t, pre_a_t, post_b_t X) XYZ { //gd:Vector3.cubic_interpolate_in_time
	return XYZ{
		Float.CubicInterpolateInTime(v.X, b.X, preA.X, postB.X, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
		Float.CubicInterpolateInTime(v.Y, b.Y, preA.Y, postB.Y, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
		Float.CubicInterpolateInTime(v.Z, b.Z, preA.Z, postB.Z, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
	}
}

// Direction returns the normalized vector pointing from this vector to to. This is equivalent to using Normalized(b - a).
func Direction(v, to XYZ) XYZ { //gd:Vector3.direction_to
	return Normalized(Sub(to, v))
}

// DistanceSquared returns the squared distance between this vector and to.
//
// This method runs faster than [Vector3.DistanceTo], so prefer it if you need to compare vectors or need the squared distance for
// some formula.
func DistanceSquared(v, to XYZ) Float.X { //gd:Vector3.distance_squared_to
	return LengthSquared(Sub(v, to))
}

// Distance returns the distance between this vector and to.
func Distance(v, to XYZ) Float.X { //gd:Vector3.distance_to
	return Length(Sub(v, to))
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
func Dot(a, b XYZ) Float.X { //gd:Vector3.dot
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Floor returns a new vector with all components rounded down (towards negative infinity).
func Floor(v XYZ) XYZ { //gd:Vector3.floor
	return XYZ{
		Float.Floor(v.X), Float.Floor(v.Y), Float.Floor(v.Z),
	}
}

// Inverse returns the inverse of the vector. This is the same as
//
//	XYZ{1/v.X,1/v.Y,1/v.Z}.
func Inverse(v XYZ) XYZ { //gd:Vector3.inverse
	return XYZ{1 / v.X, 1 / v.Y, 1 / v.Z}
}

// IsApproximatelyEqual returns true if this vector and other are approximately equal,
// by running [Float.IsApproximatelyEqual] on each component.
func IsApproximatelyEqual(a, b XYZ) bool { //gd:Vector3.is_equal_approx
	if !Float.IsApproximatelyEqual(a.X, b.X) {
		return false
	}
	if !Float.IsApproximatelyEqual(a.Y, b.Y) {
		return false
	}
	if !Float.IsApproximatelyEqual(a.Z, b.Z) {
		return false
	}
	return true
}

// IsFinite returns true if this vector's values are finite, by running [Float.IsFinite] on each component.
func IsFinite(v XYZ) bool { //gd:Vector3.is_finite
	if !Float.IsFinite(v.X) {
		return false
	}
	if !Float.IsFinite(v.Y) {
		return false
	}
	if !Float.IsFinite(v.Z) {
		return false
	}
	return true
}

// IsNormalized Returns true if the vector is normalized, i.e. its length is approximately equal to 1.
func IsNormalized(v XYZ) bool { return Float.IsApproximatelyEqual(LengthSquared(v), 1) } //gd:Vector3.is_normalized

// IsApproximatelyZero returns true if this vector is approximately equal to zero, by running
// [Float.IsApproximatelyZero] on each component.
func IsApproximatelyZero(v XYZ) bool { //gd:Vector3.is_zero_approx
	if !Float.IsApproximatelyZero(v.X) {
		return false
	}
	if !Float.IsApproximatelyZero(v.Y) {
		return false
	}
	if !Float.IsApproximatelyZero(v.Z) {
		return false
	}
	return true
}

// Length the length (magnitude) of this vector.
func Length(v XYZ) Float.X { //gd:Vector3.length
	return Float.X(math.Sqrt(float64(LengthSquared(v))))
}

// LengthSquared returns the squared length (squared magnitude) of this vector.
//
// This method runs faster than length, so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func LengthSquared(v XYZ) Float.X { //gd:Vector3.length_squared
	return Float.X(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Lerp returns the result of the linear interpolation between this vector and to by amount weight.
// weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
func Lerp[X Float.Any](v, to XYZ, weight X) XYZ { //gd:Vector3.lerp
	return XYZ{
		Float.Lerp(v.X, to.X, Float.X(weight)),
		Float.Lerp(v.Y, to.Y, Float.X(weight)),
		Float.Lerp(v.Z, to.Z, Float.X(weight)),
	}
}

// LimitLength returns the vector with a maximum length by limiting its length to length.
func LengthLimited[X Float.Any](v XYZ, length X) XYZ { //gd:Vector3.limit_length
	var l = Length(v)
	if l > 0 && Float.X(length) < l {
		v = Divf(v, l)
		v = Mulf(v, length)
	}
	return v
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all components are
// equal, this method returns [X].
func MaxAxis(v XYZ) Axis { //gd:Vector3.max_axis_index
	if v.X < v.Y {
		if v.Y < v.Z {
			return Z
		}
		return Y
	}
	if v.X < v.Z {
		return Z
	}
	return X
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all components are
// equal, this method returns [Z].
func MinAxis(v XYZ) Axis { //gd:Vector3.min_axis_index
	if v.X < v.Y {
		if v.Y < v.Z {
			return X
		}
		return Z
	}
	if v.X < v.Z {
		return Y
	}
	return Z
}

// Move returns a new vector moved toward to by the fixed delta amount. Will not go past the final value.
func Move[X Float.Any](a, b XYZ, delta X) XYZ { //gd:Vector3.move_toward
	var vd = Sub(b, a)
	var l = Length(vd)
	if l <= Float.X(delta) || l < Float.Epsilon {
		return b
	}
	return Add(a, Mulf(Divf(vd, l), delta))
}

// Normalized returns the result of scaling the vector to unit length. Equivalent to v / v.Length().
// See also [IsNormalized].
//
// Note: This function may return incorrect values if the input vector length is near zero.
func Normalized(v XYZ) XYZ { //gd:Vector3.normalized
	var l = Length(v)
	if l == 0 {
		return Zero
	}
	return Divf(v, l)
}

// OctahedronDecode returns the [XYZ] from an octahedral-compressed form created using
// OctahedronEncode (stored as a [Vector2.XY]).
func OctahedronDecode(uv Vector2.XY) XYZ { //gd:Vector3.octahedron_decode
	var f = XYZ{uv.X*2.0 - 1.0, uv.Y*2.0 - 1.0, 0}
	var n = XYZ{f.X, f.Y, 1.0 - Float.Abs(f.X) - Float.Abs(f.Y)}
	var t = Float.Clamp(-n.Z, 0.0, 1.0)
	if uv.X >= 0 {
		n.X += -t
	} else {
		n.X += t
	}
	if uv.Y >= 0 {
		n.Y += -t
	} else {
		n.Y += t
	}
	return Normalized(n)
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
func OctahedronEncode(v XYZ) Vector2.XY { //gd:Vector3.octahedron_encode
	v = Divf(v, float64(Float.Abs(v.X)+Float.Abs(v.Y)+Float.Abs(v.Z)))
	var o Vector2.XY
	if v.Z >= 0.0 {
		o.X = v.X
		o.Y = v.Y
	} else {
		if v.X > 0 {
			o.X = 1.0 - Float.Abs(v.Y)
		} else {
			o.X = -(1.0 - Float.Abs(v.Y))
		}
		if v.Y > 0 {
			o.Y = Float.Abs(v.X)
		} else {
			o.Y = -(1.0 - Float.Abs(v.X))
		}
	}
	o.X = o.X*0.5 + 0.5
	o.Y = o.Y*0.5 + 0.5
	return o
}

// Outer returns the outer product with with.
/*func Outer(v, with XYZ) Basis { //gd:Vector3.outer
return Basis{
	Vector3{v[X] * with[X], v[X] * with[Y], v[X] * with[Z]},
	Vector3{y * with[X], y * with[Y], y * with[Z]},
	Vector3{z * with[X], z * with[Y], z * with[Z]},
}
}*/

// Posmodf returns a vector composed of the Fposmod of this vector's components and mod.
func Posmodf[X Float.Any](v XYZ, mod X) XYZ { //gd:Vector3.posmod
	return XYZ{
		Float.Posmod(v.X, Float.X(mod)),
		Float.Posmod(v.Y, Float.X(mod)),
		Float.Posmod(v.Z, Float.X(mod)),
	}
}

// Posmod returns a vector composed of the Fposmod of this vector's components and mod.
func Posmod(v, mod XYZ) XYZ { //gd:Vector3.posmodv
	return XYZ{
		Float.Posmod(v.X, mod.X),
		Float.Posmod(v.Y, mod.Y),
		Float.Posmod(v.Z, mod.Z),
	}
}

// Project returns the result of projecting the vector onto the given vector b.
func Project(v, b XYZ) XYZ { //gd:Vector3.project
	return Mulf(b, Dot(v, b)/LengthSquared(b))
}

// Rotated returns the result of rotating this vector around a given axis by angle (in radians).
// The axis must be a normalized vector.
func Rotated(v, axis XYZ, angle Angle.Radians) XYZ { //gd:Vector3.rotated
	panic("not implemented")
	//return NewBasisRotatedAround(axis, angle).Transform(v)
}

// Round returns a new vector with all components rounded to the nearest integer, with halfway cases
// rounded away from zero.
func Round(v XYZ) XYZ { //gd:Vector3.round
	return XYZ{
		Float.Round(v.X), Float.Round(v.Y), Float.Round(v.Z),
	}
}

// SignedAngle returns the signed angle to the given vector, in radians. The sign of the angle is
// positive in a counter-clockwise direction and negative in a clockwise direction when viewed from
// the side specified by the axis.
func SignedAngle(v, to XYZ, axis XYZ) Angle.Radians { //gd:Vector3.signed_angle_to
	var cross_to = Cross(v, to)
	var unsigned_angle = Float.Atan2(Length(cross_to), Dot(v, to))
	var sign = Dot(cross_to, axis)
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
func Slerp[X Float.Any](v, to XYZ, weight X) XYZ { //gd:Vector3.slerp
	// This method seems more complicated than it really is, since we write out
	// the internals of some methods for efficiency (mainly, checking length).
	var start_length_sq = LengthSquared(v)
	var end_length_sq = LengthSquared(to)
	if start_length_sq == 0.0 || end_length_sq == 0.0 {
		// Zero length vectors have no angle, so the best we can do is either lerp or throw an error.
		return Lerp(v, to, weight)
	}
	var axis XYZ = Cross(v, to)
	var axis_length_sq = LengthSquared(axis)
	if axis_length_sq == 0.0 {
		// Colinear vectors have no rotation axis or angle between them, so the best we can do is lerp.
		return Lerp(v, to, weight)
	}
	axis = Divf(axis, Float.Sqrt(axis_length_sq))
	var start_length = Float.Sqrt(start_length_sq)
	var result_length = Float.Lerp(start_length, Float.Sqrt(end_length_sq), Float.X(weight))
	var angle = AngleBetween(v, to)
	return Mulf(Rotated(v, axis, angle*Angle.Radians(weight)), result_length/start_length)
}

// Slide returns a new vector slid along a plane defined by the given normal.
func Slide(v, n XYZ) XYZ { //gd:Vector3.slide
	return Sub(v, Mulf(n, Dot(v, n)))
}

// Snapped returns a new vector with each component snapped to the nearest multiple of the corresponding component
// in step. This can also be used to round the components to an arbitrary number of decimals.
func Snapped(v, step XYZ) XYZ { //gd:Vector3.snapped
	return XYZ{
		Float.Snapped(v.X, step.X),
		Float.Snapped(v.Y, step.Y),
		Float.Snapped(v.Z, step.Z),
	}
}

// Reflect returns the result of reflecting the vector from a plane defined by the given normal n.
func Reflect(v, n XYZ) XYZ { //gd:Vector3.reflect
	return Sub(Mulf(n, 2.0*Dot(v, n)), v)
}

func Add(a, b XYZ) XYZ { //gd:Vector3+(right:Vector3)
	return XYZ{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}
func Sub(a, b XYZ) XYZ { //gd:Vector3-(right:Vector3)
	return XYZ{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}
func Mul(a, b XYZ) XYZ { //gd:Vector3*(right:Vector3)
	return XYZ{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}
func Div(a, b XYZ) XYZ { //gd:Vector3/(right:Vector3)
	return XYZ{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

func Addf[X Float.Any](a XYZ, b X) XYZ { //gd:Vector3+(right:float)
	return XYZ{a.X + Float.X(b), a.Y + Float.X(b), a.Z + Float.X(b)}
}
func Subf[X Float.Any](a XYZ, b X) XYZ { //gd:Vector3-(right:float)
	return XYZ{a.X - Float.X(b), a.Y - Float.X(b), a.Z - Float.X(b)}
}
func Mulf[X Float.Any](a XYZ, b X) XYZ { //gd:Vector3*(right:float)
	return XYZ{a.X * Float.X(b), a.Y * Float.X(b), a.Z * Float.X(b)}
}
func Divf[X Float.Any](a XYZ, b X) XYZ { //gd:Vector3/(right:float)
	return XYZ{a.X / Float.X(b), a.Y / Float.X(b), a.Z / Float.X(b)}
}

func Addi[X Int.Any](a XYZ, b X) XYZ { //gd:Vector3+(right:int)
	return XYZ{a.X + Float.X(b), a.Y + Float.X(b), a.Z + Float.X(b)}
}
func Subi[X Int.Any](a XYZ, b X) XYZ { //gd:Vector3-(right:int)
	return XYZ{a.X - Float.X(b), a.Y - Float.X(b), a.Z - Float.X(b)}
}
func Muli[X Int.Any](a XYZ, b X) XYZ { //gd:Vector3*(right:int)
	return XYZ{a.X * Float.X(b), a.Y * Float.X(b), a.Z * Float.X(b)}
}
func Divi[X Int.Any](a XYZ, b X) XYZ { //gd:Vector3/(right:int)
	return XYZ{a.X / Float.X(b), a.Y / Float.X(b), a.Z / Float.X(b)}
}

func Neg(v XYZ) XYZ { return XYZ{-v.X, -v.Y, -v.Z} } //gd:Vector3-(unary)
