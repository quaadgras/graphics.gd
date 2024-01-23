package gd

import (
	"math"
	"unsafe"

	gd "grow.graphics/gd/internal"
)

const (
	Pi  = math.Pi
	Tau = math.Pi * 2

	cmpEpsilon = 0.00001
)

// ʕ is a little ternary operator for porting C code.
func ʕ[T any](q bool, a T, b T) T {
	if q {
		return a
	}
	return b
}

func sign[T ~float32 | ~float64 | ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

func fract[T ~float32 | ~float64](x T) T {
	return x - Floorf(x)
}

// Abs returns the absolute value of the parameter x (i.e. non-negative value).
func Abs[T gd.ComponentWise[T]](x T) T { return gd.Abs(x) }

// Absf returns the absolute value of the float parameter x (i.e. non-negative value).
func Absf[T ~float32 | ~float64](x T) T { return T(math.Abs(float64(x))) }

// Absi returns the absolute value of the integer parameter x (i.e. non-negative value).
func Absi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Acos returns the arc cosine of x in radians. Use to get the angle of cosine x.
// x will be clamped between -1.0 and 1.0 (inclusive), in order to prevent acos
// from returning NaN.
func Acos[T ~float32 | ~float64](x T) T {
	return T(math.Acos(float64(Clamp(x, -1.0, 1.0))))
}

// Acosh returns the hyperbolic arc (also called inverse) cosine of x, returning a value
// in radians. Use it to get the angle from an angle's cosine in hyperbolic space if x
// is larger or equal to 1. For values of x lower than 1, it will return 0, in order to
// prevent acosh from returning NaN.
func Acosh[T ~float32 | ~float64](x T) T {
	if x < 1 {
		return 0
	}
	return T(math.Acosh(float64(x)))
}

// AngleDifference returns the difference between the two angles, in the range of [-Pi, +Pi].
// When from and to are opposite, returns -Pi if from is smaller than to, or Pi otherwise.
func AngleDifference[T ~float32 | ~float64](from, to T) T {
	difference := Fmod(from, to)
	return Fmod(2*difference, Tau) - difference
}

// Asin returns the arc sine of x in radians. Use to get the angle of sine x. x will be clamped
// between -1.0 and 1.0 (inclusive), in order to prevent asin from returning NaN.
func Asin[T ~float32 | ~float64](x T) T {
	return T(math.Asin(float64(Clamp(x, -1.0, 1.0))))
}

// Asinh returns the hyperbolic arc (also called inverse) sine of x, returning a value in radians.
// Use it to get the angle from an angle's sine in hyperbolic space.
func Asinh[T ~float32 | ~float64](x T) T {
	return T(math.Asinh(float64(x)))
}

// Atan returns the arc tangent of x in radians. Use it to get the angle from an angle's tangent in
// trigonometry. The method cannot know in which quadrant the angle should fall. See atan2 if you
// have both y and x.
func Atan[T ~float32 | ~float64](x T) T {
	return T(math.Atan(float64(x)))
}

// Atan2 returns the arc tangent of y/x in radians. Use to get the angle of tangent y/x. To compute
// the value, the method takes into account the sign of both arguments in order to determine the quadrant.
//
// Important note: The Y coordinate comes first, by convention.
func Atan2[T ~float32 | ~float64](y, x T) T {
	return T(math.Atan2(float64(y), float64(x)))
}

// Atanh returns the hyperbolic arc (also called inverse) tangent of x, returning a value in radians. Use
// it to get the angle from an angle's tangent in hyperbolic space if x is between -1 and 1 (non-inclusive).
//
// In mathematics, the inverse hyperbolic tangent is only defined for -1 < x < 1 in the real set, so values
// equal or lower to -1 for x return -INF and values equal or higher than 1 return +INF in order to prevent
// atanh from returning NaN.
func Atanh[T ~float32 | ~float64](x T) T {
	if x <= -1 {
		return T(math.Inf(-1))
	}
	if x >= 1 {
		return T(math.Inf(1))
	}
	return T(math.Atanh(float64(x)))
}

// BezierDerivative returns the derivative at the given t on a one-dimensional Bézier curve defined by the given
// control_1, control_2, and end points.
func BezierDerivative[T ~float32 | ~float64](start, control_1, control_2, end, t T) T {
	/* Formula from Wikipedia article on Bezier curves. */
	omt := (1.0 - t)
	omt2 := omt * omt
	t2 := t * t
	return T((control_1-start)*3.0*omt2 + (control_2-control_1)*6.0*omt*t + (end-control_2)*3.0*t2)
}

// BezierInterpolate returns the point at the given t on a one-dimensional Bézier curve defined by the given
// control_1, control_2, and end points.
func BezierInterpolate[T ~float32 | ~float64](start, control_1, control_2, end, t T) T {
	/* Formula from Wikipedia article on Bezier curves. */
	omt := (1.0 - t)
	omt2 := omt * omt
	omt3 := omt2 * omt
	t2 := t * t
	t3 := t2 * t
	return start*omt3 + control_1*omt2*t*3.0 + control_2*omt*t2*3.0 + end*t3
}

// Ceil rounds x upward (towards positive infinity), returning the smallest whole number that is not less than x.
func Ceil[T gd.ComponentWise[T]](x T) T {
	return gd.Ceil(x)
}

// Ceilf rounds x upward (towards positive infinity), returning the smallest whole number that is not less than x.
func Ceilf[T ~float32 | ~float64](x T) T { return T(math.Ceil(float64(x))) }

// Ceili rounds x upward (towards positive infinity), returning the smallest whole number that is not less than x.
func Ceili[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return x }

type ordered interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float32 | ~float64
}

// Clamp clamps the value, returning a Variant not less than min and not more than max. Any values that can be
// compared with the less than and greater than operators will work.
func Clamp[T ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Clampf clamps the value, returning a float not less than min and not more than max.
func Clampf[T ~float32 | ~float64](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Clampi clamps the value, returning an integer not less than min and not more than max.
func Clampi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Cos returns the cosine of angle x in radians.
func Cos[T ~float32 | ~float64](x T) T { return T(math.Cos(float64(x))) }

// Cosh returns the hyperbolic cosine of x in radians.
func Cosh[T ~float32 | ~float64](x T) T { return T(math.Cosh(float64(x))) }

// CubicInterpolate cubic interpolates between two values by the factor defined in weightSee also
// with pre and post values.
func CubicInterpolate[T ~float32 | ~float64](from, to, pre, post, weight T) T {
	return 0.5 *
		((from * 2.0) +
			(-pre+to)*weight +
			(2.0*pre-5.0*from+4.0*to-post)*(weight*weight) +
			(-pre+3.0*from-3.0*to+post)*(weight*weight*weight))
}

// CubicInterpolateAngle cubic interpolates between two rotation values with shortest path
// by the factor defined in weight with pre and post values. See also [LerpAngle].
func CubicInterpolateAngle[T ~float32 | ~float64](from, to, pre, post, weight T) T {
	from_rot := Fmod(from, Tau)
	var (
		pre_diff = Fmod(pre-from_rot, Tau)
		pre_rot  = from_rot + Fmod(2.0*pre_diff, Tau) - pre_diff
	)
	var (
		to_diff = Fmod(to-from_rot, Tau)
		to_rot  = from_rot + Fmod(2.0*to_diff, Tau) - to_diff
	)
	var (
		post_diff = Fmod(post-to_rot, Tau)
		post_rot  = to_rot + Fmod(2.0*post_diff, Tau) - post_diff
	)
	return CubicInterpolate(from_rot, to_rot, pre_rot, post_rot, weight)
}

// CubicInterpolateAngleInTime cubic interpolates between two rotation values with shortest path
// by the factor defined in weight with pre and post values. See also [LerpAngle].
//
// It can perform smoother interpolation than [CubicInterpolate] by the time values.
func CubicInterpolateAngleInTime[T ~float32 | ~float64](from, to, pre, post, weight, to_t, pre_t, post_t T) T {
	from_rot := Fmod(from, Tau)
	var (
		pre_diff = Fmod(pre-from_rot, Tau)
		pre_rot  = from_rot + Fmod(2.0*pre_diff, Tau) - pre_diff
	)
	var (
		to_diff = Fmod(to-from_rot, Tau)
		to_rot  = from_rot + Fmod(2.0*to_diff, Tau) - to_diff
	)
	var (
		post_diff = Fmod(post-to_rot, Tau)
		post_rot  = to_rot + Fmod(2.0*post_diff, Tau) - post_diff
	)
	return CubicInterpolateInTime(from_rot, to_rot, pre_rot, post_rot, weight, to_t, pre_t, post_t)
}

// CubicInterpolateInTime cubic interpolates between two values by the factor defined in weight with
// pre and post values.
//
// It can perform smoother interpolation than cubic_interpolate by the time values.
func CubicInterpolateInTime[T ~float32 | ~float64](from, to, pre, post, weight, to_t, pre_t, post_t T) T {
	/* Barry-Goldman method */
	t := Lerpf(0.0, to_t, weight)
	a1 := Lerpf(pre, from, ʕ(pre_t == 0, 0.0, (t-pre_t)/-pre_t))
	a2 := Lerpf(from, to, ʕ(to_t == 0, 0.5, t/to_t))
	a3 := Lerpf(to, post, post_t-ʕ(to_t == 0, 1.0, (t-to_t)/(post_t-to_t)))
	b1 := Lerpf(a1, a2, to_t-ʕ(pre_t == 0, 0.0, (t-pre_t)/(to_t-pre_t)))
	b2 := Lerpf(a2, a3, ʕ(post_t == 0, 1.0, t/post_t))
	return Lerpf(b1, b2, ʕ(to_t == 0, 0.5, t/to_t))
}

// DecibelsToLinear converts from decibels to linear energy (audio).
func DecibelsToLinear[T ~float32 | ~float64](db T) T {
	return T(math.Exp(float64(db) * 0.11512925464970228420089957273422))
}

// DegreesToRadians converts an angle expressed in degrees to radians.
func DegreesToRadians[T ~float32 | ~float64](deg T) T {
	return deg * (Pi / 180.0)
}

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
func Ease[T ~float32 | ~float64](x, curve T) T {
	if x < 0 {
		x = 0
	} else if x > 1.0 {
		x = 1.0
	}
	if curve > 0 {
		if curve < 1.0 {
			return T(1.0 - math.Pow(1.0-float64(x), 1.0/float64(curve)))
		} else {
			return T(math.Pow(float64(x), float64(curve)))
		}
	} else if curve < 0 {
		//inout ease

		if x < 0.5 {
			return T(math.Pow(float64(x)*2.0, float64(-curve)) * 0.5)
		} else {
			return T(1.0-math.Pow(1.0-(float64(x)-0.5)*2.0, float64(-curve)))*0.5 + 0.5
		}
	} else {
		return 0 // no ease (raw)
	}
}

// Exp raises the mathematical constant e to the power of x and returns it.
// e has an approximate value of 2.71828, and can be obtained with Exp(1).
//
// For exponents to other bases use the method pow.
func Exp[T ~float32 | ~float64](x T) T { return T(math.Exp(float64(x))) }

// Floor rounds x downward (towards negative infinity), returning the largest whole number that is not
// more than x.
func Floor[T gd.ComponentWise[T]](x T) T { return gd.Floor(x) }

// Floorf rounds x downward (towards negative infinity), returning the largest whole number that is not
// more than x.
func Floorf[T ~float32 | ~float64](x T) T { return T(math.Floor(float64(x))) }

// Floori rounds x downward (towards negative infinity), returning the largest whole number that is not
// more than x.
func Floori[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return x }

// Fmod returns the floating-point remainder of x divided by y, keeping the sign of x.
func Fmod[T ~float32 | ~float64](x, y T) T {
	return T(math.Mod(float64(x), float64(y)))
}

// Fposmod returns the floating-point modulus of x divided by y, wrapping equally in positive and negative.
func Fposmod[T ~float32 | ~float64](x, y T) T {
	value := Fmod(x, y)
	if ((value < 0) && (y > 0)) || ((value > 0) && (y < 0)) {
		value += y
	}
	value += 0.0
	return value
}

// InverseLerp returns an interpolation or extrapolation factor considering the range specified in from and to,
// and the interpolated value specified in weight. The returned value will be between 0.0 and 1.0 if weight is
// between from and to (inclusive). If weight is located outside this range, then an extrapolation factor will
// be returned (return value lower than 0.0 or greater than 1.0). Use [Clamp] on the result of [InverseLerp] if
// this is not desired.
func InverseLerp[T ~float32 | ~float64](from, to, weight T) T { return (weight - from) / (to - from) }

// IsApproximatelyEqual returns true if a and b are approximately equal to each other.
//
// Here, "approximately equal" means that a and b are within a small internal epsilon of each other, which scales
// with the magnitude of the numbers.
//
// Infinity values of the same sign are considered equal.
func IsApproximatelyEqual[T ~float32 | ~float64](a, b T) bool {
	// Check for exact equality first, required to handle "infinity" values.
	if a == b {
		return true
	}
	// Then check for approximate equality.
	tolerance := cmpEpsilon * Absf(a)
	if tolerance < cmpEpsilon {
		tolerance = cmpEpsilon
	}
	return Absf(a-b) < tolerance
}

// IsFinite returns whether x is a finite value, i.e. it is not NaN, +INF, or -INF.
func IsFinite[T ~float32 | ~float64](x T) bool {
	return !math.IsInf(float64(x), 0) && !math.IsNaN(float64(x))
}

// IsInfinity returns whether x is an infinite value, i.e. it is +INF or -INF.
func IsInfinity[T ~float32 | ~float64](x T) bool {
	return math.IsInf(float64(x), 0)
}

// IsNaN returns whether x is NaN (not a number).
func IsNaN[T ~float32 | ~float64](x T) bool {
	return math.IsNaN(float64(x))
}

// IsApproximatelyZero Returns true if x is zero or almost zero. The comparison is done using a
// tolerance calculation with a small internal epsilon. This function is faster than using
// [IsApproximatelyEqual] with one value as zero.
func IsApproximatelyZero[T ~float32 | ~float64](x T) bool {
	return Absf(x) < cmpEpsilon
}

// Lerp linearly interpolates between two values by the factor defined in weight. To perform interpolation,
// weight should be between 0.0 and 1.0 (inclusive). However, values outside this range are allowed and can
// be used to perform extrapolation. If this is not desired, use [Clampf] on the result of this function.
//
// See also [InverseLerp] which performs the reverse of this operation. To perform eased interpolation with
// [Lerp], combine it with ease or smoothstep.
func Lerp[T gd.Lerpable[T]](from, to T, weight Float) T { return gd.Lerp(from, to, weight) }

// LerpAngle linearly interpolates between two angles (in radians) by a weight value between 0.0 and 1.0.
// Similar to lerp, but interpolates correctly when the angles wrap around [Tau]. To perform eased
// interpolation with [LerpAngle], combine it with [Ease] or [Smoothstep].
//
// Note: This function lerps through the shortest path between from and to. However, when these two angles
// are approximately Pi + k * Tau apart for any integer k, it's not obvious which way they lerp due to
// floating-point precision errors. For example, LerpAngle(0, Pi, weight) lerps counter-clockwise, while
// LerpAngle(0, Pi + 5 * Tau, weight) lerps clockwise.
func LerpAngle[T ~float32 | ~float64](from, to, weight T) T {
	return from + AngleDifference(from, to)*weight
}

// Lerpf linearly interpolates between two values by the factor defined in weight. To perform interpolation,
// weight should be between 0.0 and 1.0 (inclusive). However, values outside this range are allowed and can
// be used to perform extrapolation. If this is not desired, use [Clampf] on the result of this function.
//
// See also [InverseLerp] which performs the reverse of this operation. To perform eased interpolation with
// [Lerpf], combine it with ease or smoothstep.
func Lerpf[T ~float32 | ~float64](from, to, weight T) T { return from + (to-from)*weight }

// LinearToDecibels converts from linear energy (audio) to decibels.
func LinearToDecibels[T ~float32 | ~float64](energy T) T {
	return T(math.Log(float64(energy)) * 8.6858896380650365530225783783321)
}

// Log returns the natural logarithm of x (base e, with e being approximately 2.71828). This is the amount of
// time needed to reach a certain level of continuous growth.
//
// Note: This is not the same as the "log" function on most calculators, which uses a base 10 logarithm. To use
// base 10 logarithm, use Log(x) / Log(10).
//
// Note: The logarithm of 0 returns -inf, while negative values return -NaN.
func Log[T ~float32 | ~float64](x T) T { return T(math.Log(float64(x))) }

// MoveToward moves from toward to by the delta amount. Will not go past to.
// Use a negative delta value to move away.
func MoveToward[T ~float32 | ~float64](from, to, delta T) T {
	if Absf(to-from) <= delta {
		return to
	}
	return from + sign(to-from)*delta
}

// NearestPowerOfTwo returns the nearest power of two to the given value.
//
// Warning: Due to its implementation, this method returns 0 rather than 1 for values
// less than or equal to 0, with an exception for value being the smallest negative
// 64-bit integer (-9223372036854775808) in which case the value is returned unchanged.
func NearestPowerOfTwo(x Int) Int {
	get_shift_from_power_of_2 := func(bits uintptr) int {
		for i := 0; i < 32; i++ {
			if bits == (1 << i) {
				return i
			}
		}
		return -1
	}
	x--
	num := get_shift_from_power_of_2(unsafe.Sizeof(x)) + 3
	for i := 0; i < num; i++ {
		x |= x >> (1 << i)
	}
	return x + 1
}

// PingPong wraps value between 0 and the length. If the limit is reached, the next value
// the function returns is decreased to the 0 side or increased to the length side (like
// a triangle wave). If length is less than zero, it becomes positive.
func PingPong[T ~float32 | ~float64](value, length T) T {
	if length != 0 {
		return Absf(fract((value-length)/(length*2.0))*length*2.0 - length)
	}
	return 0
}

// Posmod returns the integer modulus of x divided by y that wraps equally in positive and negative.
func Posmod[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x, y T) T {
	value := x % y
	if ((value < 0) && (y > 0)) || ((value > 0) && (y < 0)) {
		value += y
	}
	return value
}

// Pow returns the result of base raised to the power of exp.
func Pow[T ~float32 | ~float64](base, exp T) T { return T(math.Pow(float64(base), float64(exp))) }

// RadiansToDegrees converts an angle expressed in radians to degrees.
func RadiansToDegrees[T ~float32 | ~float64](rad T) T { return rad * (180.0 / Pi) }

// Remap maps a value from range (istart, istop) to (ostart, ostop). See also [Lerp] and [InverseLerp].
// If value is outside (istart, istop), then the resulting value will also be outside (ostart, ostop).
// If this is not desired, use [Clamp] on the result of this function.
func Remap[T ~float32 | ~float64](value, istart, istop, ostart, ostop T) T {
	return Lerpf(ostart, ostop, InverseLerp(istart, istop, value))
}

// RotateToward rotates from toward to by the delta amount. Will not go past to.
//
// Similar to move_toward, but interpolates correctly when the angles wrap around Tau.
//
// If delta is negative, this function will rotate away from to, toward the opposite angle,
// and will not go past the opposite angle.
func RotateToward[T ~float32 | ~float64](from, to, delta T) T {
	var difference = AngleDifference(from, to)
	var abs_difference = Absf(difference)
	// When `delta < 0` move no further than to Pi radians away from `to` (as Pi is the max possible angle distance).
	return from + Clamp(delta, abs_difference-Pi, abs_difference)*(ʕ[T](difference >= 0.0, 1.0, -1.0))
}

// Round rounds x to the nearest whole number, with halfway cases rounded away from 0.
func Round[T gd.ComponentWise[T]](x T) T { return gd.Round(x) }

// Roundf rounds x to the nearest whole number, with halfway cases rounded away from 0.
func Roundf[T ~float32 | ~float64](x T) T { return T(math.Round(float64(x))) }

// Roundi rounds x to the nearest whole number, with halfway cases rounded away from 0.
func Roundi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return x }

// Sign returns the same type of value as x, with -1 for negative values, 1 for positive
// values, and 0 for zeros. For nan values it returns 0.
func Sign[T gd.ComponentWise[T]](x T) T { return gd.Sign(x) }

// Signf returns -1.0 if x is negative, 1.0 if x is positive, and 0.0 if x is zero. For NaN
// values of x it returns 0.0.
func Signf[T ~float32 | ~float64](x T) T { return T(sign(x)) }

// Signi returns -1 if x is negative, 1 if x is positive, and 0 if x is zero. For NaN values
// of x it returns 0.
func Signi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return T(sign(x)) }

// Sin returns the sine of angle x in radians.
func Sin[T ~float32 | ~float64](x T) T { return T(math.Sin(float64(x))) }

// Sinh returns the hyperbolic sine of x in radians.
func Sinh[T ~float32 | ~float64](x T) T { return T(math.Sinh(float64(x))) }

// Smoothstep returns the result of smoothly interpolating the value of x between 0 and 1,
// based on the where x lies with respect to the edges from and to.
//
// The return value is 0 if x <= from, and 1 if x >= to. If x lies between from and to,
// the returned value follows an S-shaped curve that maps x between 0 and 1.
//
// This S-shaped curve is the cubic Hermite interpolator, given by
//
//	(y) = 3*y^2 - 2*y^3 where y = (x-from) / (to-from).
func Smoothstep[T ~float32 | ~float64](from, to, x T) T {
	if IsApproximatelyEqual(from, to) {
		return from
	}
	var s = Clamp((x-from)/(to-from), 0.0, 1.0)
	return s * s * (3.0 - 2.0*s)
}

// Snapped returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snapped[T gd.ComponentWise[T]](x, step T) T { return gd.Snapped(x, step) }

// Snappedf returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snappedf[T ~float32 | ~float64](x, step T) T { return T(Snapped(Float(x), Float(step))) }

// Snappedi returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snappedi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x, step T) T {
	return T(Snapped(Int(x), Int(step)))
}

// Sqrt returns the square root of x. Where x is negative, the result is NaN.
func Sqrt[T ~float32 | ~float64](x T) T { return T(math.Sqrt(float64(x))) }

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

// StepDecimals returns the position of the first non-zero digit, after the decimal point. Note that
// the maximum return value is 10, which is a design decision in the implementation.
func StepDecimals[T ~float32 | ~float64](x T) Int {
	var abs = Absf(x)
	var decs = abs - T((int)(abs)) // Strip away integer part
	for i := 0; i < maxn; i++ {
		if decs >= T(sd[i]) {
			return Int(i)
		}
	}
	return 0
}

// Tan returns the tangent of angle x in radians.
func Tan[T ~float32 | ~float64](x T) T { return T(math.Tan(float64(x))) }

// Tanh returns the hyperbolic tangent of x in radians.
func Tanh[T ~float32 | ~float64](x T) T { return T(math.Tanh(float64(x))) }

// Wrapi value between min and max. Can be used for creating loop-alike behavior or infinite surfaces.
func Wrapi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](value, min, max T) T {
	var diff = max - min
	if diff == 0 {
		return min
	}
	return min + ((((value - min) % diff) + diff) % diff)
}

// Wrapf value between min and max. Can be used for creating loop-alike behavior or infinite surfaces.
func Wrapf[T ~float32 | ~float64](value, min, max T) T {
	var diff = max - min
	if IsApproximatelyZero(diff) {
		return min
	}
	var result = value - (diff * Floorf((value-min)/diff))
	if IsApproximatelyEqual(result, max) {
		return min
	}
	return result
}
