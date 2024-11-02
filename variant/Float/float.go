// Package Float provides a generic math library for floating-point numbers.
package Float

import (
	"math"
	"math/rand"
)

const Epsilon = 0.00001

var Inf X = X(math.Inf(1)) // Positive infinity.

// Any float.
type Any interface{ ~float32 | ~float64 }

// ʕ is a little ternary operator for porting C code.
func ʕ[T any](q bool, a T, b T) T {
	if q {
		return a
	}
	return b
}

// Abs returns the absolute value of the float parameter x (i.e. non-negative value).
func Abs[X Any](x X) X { return X(math.Abs(float64(x))) } //gd:absf abs

// BezierDerivative returns the derivative at the given t on a one-dimensional Bézier curve defined by the given
// control_1, control_2, and end points.
func BezierDerivative[T Any](start, control_1, control_2, end, t T) T { //gd:bezier_derivative
	/* Formula from Wikipedia article on Bezier curves. */
	omt := (1.0 - t)
	omt2 := omt * omt
	t2 := t * t
	return T((control_1-start)*3.0*omt2 + (control_2-control_1)*6.0*omt*t + (end-control_2)*3.0*t2)
}

// BezierInterpolate returns the point at the given t on a one-dimensional Bézier curve defined by the given
// control_1, control_2, and end points.
func BezierInterpolate[T Any](start, control_1, control_2, end, t T) T { //gd:bezier_interpolate
	/* Formula from Wikipedia article on Bezier curves. */
	omt := (1.0 - t)
	omt2 := omt * omt
	omt3 := omt2 * omt
	t2 := t * t
	t3 := t2 * t
	return start*omt3 + control_1*omt2*t*3.0 + control_2*omt*t2*3.0 + end*t3
}

// Ceil rounds x upward (towards positive infinity), returning the smallest whole number that is not less than x.
func Ceil[X Any](x X) X { return X(math.Ceil(float64(x))) } //gd:ceilf

// Clamp clamps the value, returning a Variant not less than min and not more than max. Any values that can be
// compared with the less than and greater than operators will work.
func Clamp[X Any](value, min, max X) X { //gd:clampf clamp
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// CubicInterpolate cubic interpolates between two values by the factor defined in weightSee also
// with pre and post values.
func CubicInterpolate[X Any](from, to, pre, post, weight X) X { //gd:cubic_interpolate
	return 0.5 *
		((from * 2.0) +
			(-pre+to)*weight +
			(2.0*pre-5.0*from+4.0*to-post)*(weight*weight) +
			(-pre+3.0*from-3.0*to+post)*(weight*weight*weight))
}

// CubicInterpolateInTime cubic interpolates between two values by the factor defined in weight with
// pre and post values.
//
// It can perform smoother interpolation than cubic_interpolate by the time values.
func CubicInterpolateInTime[X Any](from, to, pre, post, weight, to_t, pre_t, post_t X) X { //gd:cubic_interpolate_in_time
	/* Barry-Goldman method */
	t := Lerp(0.0, to_t, weight)
	a1 := Lerp(pre, from, ʕ(pre_t == 0, 0.0, (t-pre_t)/-pre_t))
	a2 := Lerp(from, to, ʕ(to_t == 0, 0.5, t/to_t))
	a3 := Lerp(to, post, post_t-ʕ(to_t == 0, 1.0, (t-to_t)/(post_t-to_t)))
	b1 := Lerp(a1, a2, to_t-ʕ(pre_t == 0, 0.0, (t-pre_t)/(to_t-pre_t)))
	b2 := Lerp(a2, a3, ʕ(post_t == 0, 1.0, t/post_t))
	return Lerp(b1, b2, ʕ(to_t == 0, 0.5, t/to_t))
}

// Floorf rounds x downward (towards negative infinity), returning the largest whole number that is not
// more than x.
func Floor[X Any](x X) X { return X(math.Floor(float64(x))) } //gd:floorf

// IsApproximatelyEqual returns true if a and b are approximately equal to each other.
//
// Here, "approximately equal" means that a and b are within a small internal epsilon of each other, which scales
// with the magnitude of the numbers.
//
// Infinity values of the same sign are considered equal.
func IsApproximatelyEqual[X Any](a, b X) bool { //gd:is_equal_approx
	// Check for exact equality first, required to handle "infinity" values.
	if a == b {
		return true
	}
	// Then check for approximate equality.
	tolerance := Epsilon * Abs(a)
	if tolerance < Epsilon {
		tolerance = Epsilon
	}
	return Abs(a-b) < tolerance
}

// IsApproximatelyZero Returns true if x is zero or almost zero. The comparison is done using a
// tolerance calculation with a small internal [Epsilon]. This function is faster than using
// [IsApproximatelyEqual] with one value as zero.
func IsApproximatelyZero[X Any](x X) bool { //gd:is_zero_approx
	return Abs(x) < Epsilon
}

// IsFinite returns whether x is a finite value, i.e. it is not NaN, +INF, or -INF.
func IsFinite[X Any](x X) bool { //gd:is_finite
	return !math.IsInf(float64(x), 0) && !math.IsNaN(float64(x))
}

// IsNaN returns whether x is a NaN (Not a Number) value.
func IsNaN[X Any](x X) bool { //gd:is_nan
	return math.IsNaN(float64(x))
}

// IsInf returns whether x is an infinite value, either +INF or -INF.
func IsInf[X Any](x X) bool { //gd:is_inf
	return math.IsInf(float64(x), 0)
}

// Log returns the natural logarithm of x.
func Log[X Any](x X) X { return X(math.Log(float64(x))) } //gd:log

// Exp returns e raised to the power of x.
func Exp[X Any](x X) X { return X(math.Exp(float64(x))) } //gd:exp

// Lerpf linearly interpolates between two values by the factor defined in weight. To perform interpolation,
// weight should be between 0.0 and 1.0 (inclusive). However, values outside this range are allowed and can
// be used to perform extrapolation. If this is not desired, use [Clampf] on the result of this function.
//
// See also [InverseLerp] which performs the reverse of this operation. To perform eased interpolation with
// [Lerpf], combine it with ease or smoothstep.
func Lerp[X Any](from, to, weight X) X { return from + (to-from)*weight } //gd:lerpf lerp

// Fmod returns the floating-point remainder of x divided by y, keeping the sign of x.
func Mod[X Any](x, y X) X { //gd:fmod
	return X(math.Mod(float64(x), float64(y)))
}

// Fposmod returns the floating-point modulus of x divided by y, wrapping equally in positive and negative.
func Posmod[X Any](x, y X) X { //gd:fposmod
	value := Mod(x, y)
	if ((value < 0) && (y > 0)) || ((value > 0) && (y < 0)) {
		value += y
	}
	value += 0.0
	return value
}

// Pow returns the result of base raised to the power of exp.
func Pow[X Any](base, exp X) X { return X(math.Pow(float64(base), float64(exp))) } //gd:pow

// Round rounds x to the nearest whole number, with halfway cases rounded away from 0.
func Round[X Any](x X) X { return X(math.Round(float64(x))) } //gd:roundf

// Sign returns -1.0 if x is negative, 1.0 if x is positive, and 0.0 if x is zero. For NaN
// values of x it returns 0.0.
func Sign[X Any](x X) X { //gd:signf sign
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

// Snapped returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snapped[X Any](x, step X) X { //gd:snappedf snapped
	if step != 0 {
		x = X(math.Floor(float64(x/step)+0.5)) * step
	}
	return x
}

// Sqrt returns the square root of x. Where x is negative, the result is NaN.
func Sqrt[X Any](x X) X { return X(math.Sqrt(float64(x))) } //gd:sqrt

// Ease returns an "eased" value of x based on an easing function defined with curve.
// This easing function is based on an exponent. The curve can be any floating-point number,
// with specific values leading to the following behaviors:
//
// - Lower than -1.0 (exclusive): Ease in-out
// - 1.0: Linear
// - Between -1.0 and 0.0 (exclusive): Ease out-in
// - 0.0: Constant
// - Between 0.0 to 1.0 (exclusive): Ease out
// - 1.0: Linear
// - Greater than 1.0 (exclusive): Ease in
//
// https://raw.githubusercontent.com/godotengine/godot-docs/master/img/ease_cheatsheet.png
//
// See also [Smoothstep]. If you need to perform more advanced transitions, use [Tween.InterpolateValue].
func Ease[X Any](x, curve X) X { //gd:ease
	if x < 0 {
		x = 0
	} else if x > 1.0 {
		x = 1.0
	}
	if curve > 0 {
		if curve < 1.0 {
			return X(1.0 - math.Pow(1.0-float64(x), 1.0/float64(curve)))
		} else {
			return X(math.Pow(float64(x), float64(curve)))
		}
	} else if curve < 0 {
		//inout ease
		if x < 0.5 {
			return X(math.Pow(float64(x)*2.0, float64(-curve)) * 0.5)
		} else {
			return X(1.0-math.Pow(1.0-(float64(x)-0.5)*2.0, float64(-curve)))*0.5 + 0.5
		}
	} else {
		return 0 // no ease (raw)
	}
}

// StepDecimals returns the position of the first non-zero digit, after the decimal point. Note that
// the maximum return value is 10, which is a design decision in the implementation.
func StepDecimals[X Any](x X) int { //gd:step_decimals
	const maxn int = 10
	var sd = [maxn]float64{
		0.9999, // somehow compensate for floating point error
		0.09999,
		0.009999,
		0.0009999,
		0.00009999,
		0.000009999,
		0.0000009999,
		0.00000009999,
		0.000000009999,
		0.0000000009999,
	}
	var abs = Abs(x)
	var decs = abs - X((int)(abs)) // Strip away integer part
	for i := 0; i < maxn; i++ {
		if decs >= X(sd[i]) {
			return int(i)
		}
	}
	return 0
}

// InverseLerp returns an interpolation or extrapolation factor considering the range specified in from and to,
// and the interpolated value specified in weight. The returned value will be between 0.0 and 1.0 if weight is
// between from and to (inclusive). If weight is located outside this range, then an extrapolation factor will
// be returned (return value lower than 0.0 or greater than 1.0). Use [Clamp] on the result of [InverseLerp] if
// this is not desired.
func InverseLerp[X Any](from, to, weight X) X { return (weight - from) / (to - from) } //gd:inverse_lerp

// Remap maps a value from range (istart, istop) to (ostart, ostop). See also [Lerp] and [InverseLerp].
// If value is outside (istart, istop), then the resulting value will also be outside (ostart, ostop).
// If this is not desired, use [Clamp] on the result of this function.
func Remap[X Any](value, istart, istop, ostart, ostop X) X { //gd:remap
	return Lerp(ostart, ostop, InverseLerp(istart, istop, value))
}

// Smoothstep returns the result of smoothly interpolating the value of x between 0 and 1,
// based on the where x lies with respect to the edges from and to.
//
// The return value is 0 if x <= from, and 1 if x >= to. If x lies between from and to,
// the returned value follows an S-shaped curve that maps x between 0 and 1.
//
// This S-shaped curve is the cubic Hermite interpolator, given by
//
//	(y) = 3*y^2 - 2*y^3 where y = (x-from) / (to-from).
func Smoothstep[X Any](from, to, x X) X { //gd:smoothstep
	if IsApproximatelyEqual(from, to) {
		return from
	}
	var s = Clamp((x-from)/(to-from), 0.0, 1.0)
	return s * s * (3.0 - 2.0*s)
}

// MoveToward moves from toward to by the delta amount. Will not go past to.
// Use a negative delta value to move away.
func MoveToward[X Any](from, to, delta X) X { //gd:move_toward
	if Abs(to-from) <= delta {
		return to
	}
	return from + Sign(to-from)*delta
}

// LinearToDecibels converts from linear energy (audio) to decibels.
func LinearToDecibels[X Any](energy X) X { //gd:linear_to_db
	return X(math.Log(float64(energy)) * 8.6858896380650365530225783783321)
}

// DecibelsToLinear converts from decibels to linear energy (audio).
func DecibelsToLinear[X Any](db X) X { //gd:db_to_linear
	return X(math.Exp(float64(db) * 0.11512925464970228420089957273422))
}

// Wrap wraps the float value between min and max. Can be used for creating loop-alike
// behavior or infinite surfaces.
func Wrap[X Any](value, min, max X) X { //gd:wrapf wrap
	var diff = max - min
	if IsApproximatelyZero(diff) {
		return min
	}
	var result = value - (diff * Floor((value-min)/diff))
	if IsApproximatelyEqual(result, max) {
		return min
	}
	return result
}

// Max returns the largest of the two values.
func Max[T Any](a, b T) T { //gd:maxf max
	return max(a, b)
}

// Min returns the smallest of the two values.
func Min[T Any](a, b T) T { //gd:minf min
	return min(a, b)
}

func fract[T ~float32 | ~float64](x T) T {
	return x - Floor(x)
}

// PingPong wraps value between 0 and the length. If the limit is reached, the next value
// the function returns is decreased to the 0 side or increased to the length side (like
// a triangle wave). If length is less than zero, it becomes positive.
func PingPong[X Any](value, length X) X { //gd:pingpong
	if length != 0 {
		return Abs(fract((value-length)/(length*2.0))*length*2.0 - length)
	}
	return 0
}

// Random returns a random integer between min and max (inclusive).
func Random() X { //gd:randf
	return X(rand.Float64())
}

// RandomBetween returns a random float between min and max (inclusive).
func RandomBetween(min, max X) X { //gd:randf_range
	return X(rand.Float64()*(float64(max)-float64(min)) + float64(min))
}

// RandomlyDistributed eeturns a normally-distributed, pseudo-random floating-point
// value from the specified mean and a standard deviation. This is also known as a
// Gaussian distribution.
func RandomlyDistributed(mean, deviation X) X { //gd:randfn
	return X(rand.NormFloat64()*float64(deviation) + float64(mean))
}
