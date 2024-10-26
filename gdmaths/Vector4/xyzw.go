// Package Vector4i provides a 4D vector using integer coordinates.
package Vector4i

import (
	"math"

	"grow.graphics/gd/gdmaths/Float"
	"grow.graphics/gd/gdmaths/Int"
)

// XYZW is a 4-element structure that can be used to represent
// 4D coordinates or any other quadruplet of numeric values.
//
// It uses floating-point coordinates. By default, these
// floating-point values use 32-bit precision, unlike float
// which is always 64-bit. If double precision is needed,
// compile with the Go build tag 'precision_double'.
//
// See [Vector4i.XYZW] for its integer counterpart.
type XYZW = struct {
	X Float.X // The vector's X component.
	Y Float.X // The vector's Y component.
	Z Float.X // The vector's Z component.
	W Float.X // The vector's W component.
}

type Axis int

const (
	X Axis = iota // Enumerated value for the X axis. Returned by [MaxAxis] and [MinAxis].
	Y             // Enumerated value for the Y axis. Returned by [MaxAxis] and [MinAxis].
	Z             // Enumerated value for the Z axis. Returned by [MaxAxis] and [MinAxis].
	W             // Enumerated value for the W axis. Returned by [MaxAxis] and [MinAxis].
)

var (
	Zero = XYZW{0, 0, 0, 0}                                 // Zero vector, a vector with all components set to 0.
	One  = XYZW{1, 1, 1, 1}                                 // One vector, a vector with all components set to 1.
	Inf  = XYZW{Float.Inf, Float.Inf, Float.Inf, Float.Inf} // Inf vector, a vector with all components set to positive infinity.
)

// New returns a [XYZW] with the given components.
func New[X Int.Any | Float.Any](x, y, z, w X) XYZW { //gd:Vector4(x:float,y:float,z:float,w:float)
	return XYZW{Float.X(x), Float.X(y), Float.X(z), Float.X(w)}
}

// Abs returns a new vector with all components in absolute values (i.e. positive).
func Abs(v XYZW) XYZW { //gd:Vector4.abs
	return XYZW{
		Float.Abs(v.X), Float.Abs(v.Y), Float.Abs(v.Z), Float.Abs(v.W),
	}
}

// Ceil returns a new vector with all components rounded up (towards positive infinity).
func Ceil(v XYZW) XYZW { //gd:Vector4.ceil
	return XYZW{
		Float.Ceil(v.X), Float.Ceil(v.Y), Float.Ceil(v.Z), Float.Ceil(v.W),
	}
}

// Clamp returns a new vector with all components clamped between the components of min and max,
// by running [Float.Clamp] on each component.
func Clamp(v, min, max XYZW) XYZW { //gd:Vector4.clamp
	return XYZW{
		Float.Clamp(v.X, min.X, max.X),
		Float.Clamp(v.Y, min.Y, max.Y),
		Float.Clamp(v.Z, min.Z, max.Z),
		Float.Clamp(v.W, min.W, max.W),
	}
}

// Clamp returns a new vector with all components clamped between the components of min and max,
// by running [Float.Clamp] on each component.
func Clampf[X Float.Any](v XYZW, min, max X) XYZW { //gd:Vector4.clampf
	return XYZW{
		Float.Clamp(v.X, Float.X(min), Float.X(max)),
		Float.Clamp(v.Y, Float.X(min), Float.X(max)),
		Float.Clamp(v.Z, Float.X(min), Float.X(max)),
		Float.Clamp(v.W, Float.X(min), Float.X(max)),
	}
}

// CubicInterpolate performs a cubic interpolation between this vector and b using pre_a and post_b as handles
// and returns the result at position weight. weight is on the range of 0.0 to 1.0, representing the amount of
// interpolation.
func CubicInterpolate[X Float.Any](v, b, preA, postB XYZW, weight X) XYZW { //gd:Vector4.cubic_interpolate
	return XYZW{
		Float.CubicInterpolate(v.X, b.X, preA.X, postB.X, Float.X(weight)),
		Float.CubicInterpolate(v.Y, b.Y, preA.Y, postB.Y, Float.X(weight)),
		Float.CubicInterpolate(v.Z, b.Z, preA.Z, postB.Z, Float.X(weight)),
		Float.CubicInterpolate(v.W, b.W, preA.W, postB.W, Float.X(weight)),
	}
}

// CubicInterpolateInTime performs a cubic interpolation between this vector and b using pre_a and post_b as handles
// and returns the result at position weight. weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
//
// It can perform smoother interpolation than [Vector3.CubicInterpolate] by the time values.
func CubicInterpolateInTime[X Float.Any](v, b, preA, postB XYZW, weight, b_t, pre_a_t, post_b_t X) XYZW { //gd:Vector4.cubic_interpolate_in_time
	return XYZW{
		Float.CubicInterpolateInTime(v.X, b.X, preA.X, postB.X, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
		Float.CubicInterpolateInTime(v.Y, b.Y, preA.Y, postB.Y, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
		Float.CubicInterpolateInTime(v.Z, b.Z, preA.Z, postB.Z, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
		Float.CubicInterpolateInTime(v.W, b.W, preA.W, postB.W, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
	}
}

// Direction returns the normalized vector pointing from this vector to to. This is equivalent to using Normalized(b - a).
func Direction(v, to XYZW) XYZW { //gd:Vector4.direction_to
	return Normalized(Sub(to, v))
}

// DistanceSquared returns the squared distance between this vector and to.
//
// This method runs faster than [Vector3.DistanceTo], so prefer it if you need to compare vectors or need the squared distance for
// some formula.
func DistanceSquared(v, to XYZW) Float.X { //gd:Vector4.distance_squared_to
	return LengthSquared(Sub(v, to))
}

// Distance returns the distance between this vector and to.
func Distance(v, to XYZW) Float.X { //gd:Vector4.distance_to
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
func Dot(a, b XYZW) Float.X { //gd:Vector4.dot
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

// Floor returns a new vector with all components rounded down (towards negative infinity).
func Floor(v XYZW) XYZW { //gd:Vector4.floor
	return XYZW{
		Float.Floor(v.X), Float.Floor(v.Y), Float.Floor(v.Z), Float.Floor(v.W),
	}
}

// Inverse returns the inverse of the vector. This is the same as
//
//	XYZ{1/v.X,1/v.Y,1/v.Z}.
func Inverse(v XYZW) XYZW { //gd:Vector4.inverse
	return XYZW{1 / v.X, 1 / v.Y, 1 / v.Z, 1 / v.W}
}

// IsApproximatelyEqual returns true if this vector and other are approximately equal,
// by running [Float.IsApproximatelyEqual] on each component.
func IsApproximatelyEqual(a, b XYZW) bool { //gd:Vector4.is_equal_approx
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
func IsFinite(v XYZW) bool { //gd:Vector4.is_finite
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
func IsNormalized(v XYZW) bool { return Float.IsApproximatelyEqual(LengthSquared(v), 1) } //gd:Vector4.is_normalized

// IsApproximatelyZero returns true if this vector is approximately equal to zero, by running
// [Float.IsApproximatelyZero] on each component.
func IsApproximatelyZero(v XYZW) bool { //gd:Vector4.is_zero_approx
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
func Length(v XYZW) Float.X { //gd:Vector4.length
	return Float.X(math.Sqrt(float64(LengthSquared(v))))
}

// LengthSquared returns the squared length (squared magnitude) of this vector.
//
// This method runs faster than length, so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func LengthSquared(v XYZW) Float.X { //gd:Vector4.length_squared
	return Float.X(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Lerp returns the result of the linear interpolation between this vector and to by amount weight.
// weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
func Lerp[X Float.Any](v, to XYZW, weight X) XYZW { //gd:Vector4.lerp
	return XYZW{
		Float.Lerp(v.X, to.X, Float.X(weight)),
		Float.Lerp(v.Y, to.Y, Float.X(weight)),
		Float.Lerp(v.Z, to.Z, Float.X(weight)),
		Float.Lerp(v.W, to.W, Float.X(weight)),
	}
}

// Max returns the component-wise maximum of this and with, equivalent to
//
//	Vector4(max(x, with.x), max(y, with.y), max(z, with.z), max(w, with.w))
func Max(a, b XYZW) XYZW { //gd:Vector4.max
	return XYZW{
		max(a.X, b.X), max(a.Y, b.Y), max(a.Z, b.Z), max(a.W, b.W),
	}
}

// Maxf returns the component-wise maximum of this and with, equivalent to
//
//	Vector4(max(x, with), max(y, with), max(z, with), max(w, with)).
func Maxf[X Float.Any](v XYZW, with X) XYZW { //gd:Vector4.maxf
	return XYZW{
		max(v.X, Float.X(with)), max(v.Y, Float.X(with)), max(v.Z, Float.X(with)), max(v.W, Float.X(with)),
	}
}

// Min returns the component-wise minimum of this and with, equivalent to
//
//	Vector4(min(x, with.x), min(y, with.y), min(z, with.z), min(w, with.w))
func Min(a, b XYZW) XYZW { //gd:Vector4.min
	return XYZW{
		min(a.X, b.X), min(a.Y, b.Y), min(a.Z, b.Z), min(a.W, b.W),
	}
}

// Minf returns the component-wise minimum of this and with, equivalent to
//
//	Vector4(min(x, with), min(y, with), min(z, with), min(w, with)).
func Minf[X Float.Any](v XYZW, with X) XYZW { //gd:Vector4.minf
	return XYZW{
		min(v.X, Float.X(with)), min(v.Y, Float.X(with)), min(v.Z, Float.X(with)), min(v.W, Float.X(with)),
	}
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all components are
// equal, this method returns [X].
func MaxAxis(v XYZW) Axis { //gd:Vector4.max_axis_index
	var a = AsArray(v)
	var max_index = 0
	var max_value = a[X]
	for i := 1; i < len(a); i++ {
		if a[i] > max_value {
			max_index = i
			max_value = a[i]
		}
	}
	return Axis(max_index)
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all components are
// equal, this method returns [Z].
func MinAxis(v XYZW) Axis { //gd:Vector4.min_axis_index
	var a = AsArray(v)
	var min_index = 0
	var min_value = a[W]
	for i := 1; i < len(a); i++ {
		if a[i] < min_value {
			min_index = i
			min_value = a[i]
		}
	}
	return Axis(min_index)
}

// Normalized returns the result of scaling the vector to unit length. Equivalent to v / v.Length().
// See also [IsNormalized].
//
// Note: This function may return incorrect values if the input vector length is near zero.
func Normalized(v XYZW) XYZW { //gd:Vector4.normalized
	var l = Length(v)
	if l == 0 {
		return Zero
	}
	return DivX(v, l)
}

// Posmodf returns a vector composed of the Fposmod of this vector's components and mod.
func Posmodf[X Float.Any](v XYZW, mod X) XYZW { //gd:Vector4.posmod
	return XYZW{
		Float.Posmod(v.X, Float.X(mod)),
		Float.Posmod(v.Y, Float.X(mod)),
		Float.Posmod(v.Z, Float.X(mod)),
		Float.Posmod(v.W, Float.X(mod)),
	}
}

// Posmod returns a vector composed of the Fposmod of this vector's components and mod.
func Posmod(v, mod XYZW) XYZW { //gd:Vector4.posmodv
	return XYZW{
		Float.Posmod(v.X, mod.X),
		Float.Posmod(v.Y, mod.Y),
		Float.Posmod(v.Z, mod.Z),
		Float.Posmod(v.W, mod.W),
	}
}

// Round returns a new vector with all components rounded to the nearest integer, with halfway cases
// rounded away from zero.
func Round(v XYZW) XYZW { //gd:Vector4.round
	return XYZW{
		Float.Round(v.X), Float.Round(v.Y), Float.Round(v.Z), Float.Round(v.W),
	}
}

// Sign returns a new vector with each component set to 1.0 if it's positive, -1.0 if it's negative,
// and 0.0 if it's zero. The result is identical to calling [Float.Sign] on each component.
func Sign(vec XYZW) XYZW { //gd:Vector4.sign
	return XYZW{Float.Sign(vec.X), Float.Sign(vec.Y), Float.Sign(vec.Z), Float.Sign(vec.W)}
}

// Snapped returns a new vector with each component snapped to the nearest multiple of the corresponding component
// in step. This can also be used to round the components to an arbitrary number of decimals.
func Snapped(v, step XYZW) XYZW { //gd:Vector4.snapped
	return XYZW{
		Float.Snapped(v.X, step.X),
		Float.Snapped(v.Y, step.Y),
		Float.Snapped(v.Z, step.Z),
		Float.Snapped(v.W, step.W),
	}
}

// Snappedf returns a new vector with each component snapped to the nearest multiple of step.
// This can also be used to round the components to an arbitrary number of decimals.
func Snappedf[X Float.Any](v XYZW, step X) XYZW { //gd:Vector4.snappedf
	return XYZW{
		Float.Snapped(v.X, Float.X(step)),
		Float.Snapped(v.Y, Float.X(step)),
		Float.Snapped(v.Z, Float.X(step)),
		Float.Snapped(v.W, Float.X(step)),
	}
}

func Add(a, b XYZW) XYZW { return XYZW{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W} } //gd:Vector4+(right:Vector4i)
func Sub(a, b XYZW) XYZW { return XYZW{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W} } //gd:Vector4-(right:Vector4i)
func Mul(a, b XYZW) XYZW { return XYZW{a.X * b.X, a.Y * b.Y, a.Z * b.Z, a.W * b.W} } //gd:Vector4*(right:Vector4i)
func Div(a, b XYZW) XYZW { return XYZW{a.X / b.X, a.Y / b.Y, a.Z / b.Z, a.W / b.W} } //gd:Vector4/(right:Vector4i)

func AddX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4+(right:float)
	return XYZW{a.X + Float.X(b), a.Y + Float.X(b), a.Z + Float.X(b), a.W + Float.X(b)}
}
func SubX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4-(right:float)
	return XYZW{a.X - Float.X(b), a.Y - Float.X(b), a.Z - Float.X(b), a.W - Float.X(b)}
}
func MulX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4*(right:float)
	return XYZW{a.X * Float.X(b), a.Y * Float.X(b), a.Z * Float.X(b), a.W * Float.X(b)}
}
func DivX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4/(right:float)
	return XYZW{a.X / Float.X(b), a.Y / Float.X(b), a.Z / Float.X(b), a.W / Float.X(b)}
}

func Neg(v XYZW) XYZW { return XYZW{-v.X, -v.Y, -v.Z, -v.W} } //gd:Vector4-(unary)

func AsArray(vec XYZW) [4]Float.X { return [4]Float.X{vec.X, vec.Y, vec.Z, vec.W} }

func Index[I Int.Any](v XYZW, i I) int { //gd:Vector4[](index:int)
	switch Axis(i) {
	case X:
		return int(v.X)
	case Y:
		return int(v.Y)
	case Z:
		return int(v.Z)
	case W:
		return int(v.W)
	default:
		panic("index out of range")
	}
}
