package gd

import (
	"math"

	"grow.graphics/gd/internal/spatial"
)

const (
	Pi  = math.Pi
	Tau = math.Pi * 2
)

// Abs returns the absolute value of the parameter x (i.e. non-negative value).
func Abs[T spatial.ComponentWise[T]](x T) T { return spatial.Abs(x) }

// Absf returns the absolute value of the float parameter x (i.e. non-negative value).
func Absf[T ~float32 | ~float64](x T) T { return spatial.Absf(x) }

// Absi returns the absolute value of the integer parameter x (i.e. non-negative value).
func Absi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return spatial.Absi(x) }

// Acos returns the arc cosine of x in radians. Use to get the angle of cosine x.
// x will be clamped between -1.0 and 1.0 (inclusive), in order to prevent acos
// from returning NaN.
func Acos[T ~float32 | ~float64](x T) Radians { return spatial.Acos(x) }

// Acosh returns the hyperbolic arc (also called inverse) cosine of x, returning a value
// in radians. Use it to get the angle from an angle's cosine in hyperbolic space if x
// is larger or equal to 1. For values of x lower than 1, it will return 0, in order to
// prevent acosh from returning NaN.
func Acosh[T ~float32 | ~float64](x T) Radians { return spatial.Acosh(x) }

// AngleDifference returns the difference between the two angles, in the range of [-Pi, +Pi].
// When from and to are opposite, returns -Pi if from is smaller than to, or Pi otherwise.
func AngleDifference[T ~float32 | ~float64](from, to T) T { return spatial.AngleDifference(from, to) }

// Asin returns the arc sine of x in radians. Use to get the angle of sine x. x will be clamped
// between -1.0 and 1.0 (inclusive), in order to prevent asin from returning NaN.
func Asin[T ~float32 | ~float64](x T) Radians { return spatial.Asin(x) }

// Asinh returns the hyperbolic arc (also called inverse) sine of x, returning a value in radians.
// Use it to get the angle from an angle's sine in hyperbolic space.
func Asinh[T ~float32 | ~float64](x T) Radians { return spatial.Asinh(x) }

// Atan returns the arc tangent of x in radians. Use it to get the angle from an angle's tangent in
// trigonometry. The method cannot know in which quadrant the angle should fall. See atan2 if you
// have both y and x.
func Atan[T ~float32 | ~float64](x T) Radians { return spatial.Atan(x) }

// Atan2 returns the arc tangent of y/x in radians. Use to get the angle of tangent y/x. To compute
// the value, the method takes into account the sign of both arguments in order to determine the quadrant.
//
// Important note: The Y coordinate comes first, by convention.
func Atan2[T ~float32 | ~float64](y, x T) Radians { return spatial.Atan2(y, x) }

// Atanh returns the hyperbolic arc (also called inverse) tangent of x, returning a value in radians. Use
// it to get the angle from an angle's tangent in hyperbolic space if x is between -1 and 1 (non-inclusive).
//
// In mathematics, the inverse hyperbolic tangent is only defined for -1 < x < 1 in the real set, so values
// equal or lower to -1 for x return -INF and values equal or higher than 1 return +INF in order to prevent
// atanh from returning NaN.
func Atanh[T ~float32 | ~float64](x T) Radians { return spatial.Atanh(x) }

// BezierDerivative returns the derivative at the given t on a one-dimensional Bézier curve defined by the given
// control_1, control_2, and end points.
func BezierDerivative[T ~float32 | ~float64](start, control_1, control_2, end, t T) T {
	return spatial.BezierDerivative(start, control_1, control_2, end, t)
}

// BezierInterpolate returns the point at the given t on a one-dimensional Bézier curve defined by the given
// control_1, control_2, and end points.
func BezierInterpolate[T ~float32 | ~float64](start, control_1, control_2, end, t T) T {
	return spatial.BezierInterpolate(start, control_1, control_2, end, t)
}

// Ceil rounds x upward (towards positive infinity), returning the smallest whole number that is not less than x.
func Ceil[T spatial.ComponentWise[T]](x T) T { return spatial.Ceil(x) }

// Ceilf rounds x upward (towards positive infinity), returning the smallest whole number that is not less than x.
func Ceilf[T ~float32 | ~float64](x T) T { return spatial.Ceilf(x) }

// Ceili rounds x upward (towards positive infinity), returning the smallest whole number that is not less than x.
func Ceili[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return x }

type ordered interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float32 | ~float64
}

// Clamp clamps the value, returning a Variant not less than min and not more than max. Any values that can be
// compared with the less than and greater than operators will work.
func Clamp[T ordered](value, min, max T) T { return spatial.Clamp(value, min, max) }

// Clampf clamps the value, returning a float not less than min and not more than max.
func Clampf[T ~float32 | ~float64](value, min, max T) T { return spatial.Clampf(value, min, max) }

// Clampi clamps the value, returning an integer not less than min and not more than max.
func Clampi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](value, min, max T) T {
	return spatial.Clampi(value, min, max)
}

// Cos returns the cosine of angle x in radians.
func Cos[T ~float32 | ~float64](x T) T { return spatial.Cos(x) }

// Cosh returns the hyperbolic cosine of x in radians.
func Cosh[T ~float32 | ~float64](x T) T { return spatial.Cosh(x) }

// CubicInterpolate cubic interpolates between two values by the factor defined in weightSee also
// with pre and post values.
func CubicInterpolate[T ~float32 | ~float64](from, to, pre, post, weight T) T {
	return spatial.CubicInterpolate(from, to, pre, post, weight)
}

// CubicInterpolateAngle cubic interpolates between two rotation values with shortest path
// by the factor defined in weight with pre and post values. See also [LerpAngle].
func CubicInterpolateAngle[T ~float32 | ~float64](from, to, pre, post, weight T) T {
	return spatial.CubicInterpolateAngle(from, to, pre, post, weight)
}

// CubicInterpolateAngleInTime cubic interpolates between two rotation values with shortest path
// by the factor defined in weight with pre and post values. See also [LerpAngle].
//
// It can perform smoother interpolation than [CubicInterpolate] by the time values.
func CubicInterpolateAngleInTime[T ~float32 | ~float64](from, to, pre, post, weight, to_t, pre_t, post_t T) T {
	return spatial.CubicInterpolateAngleInTime(from, to, pre, post, weight, to_t, pre_t, post_t)
}

// CubicInterpolateInTime cubic interpolates between two values by the factor defined in weight with
// pre and post values.
//
// It can perform smoother interpolation than cubic_interpolate by the time values.
func CubicInterpolateInTime[T ~float32 | ~float64](from, to, pre, post, weight, to_t, pre_t, post_t T) T {
	return spatial.CubicInterpolateInTime(from, to, pre, post, weight, to_t, pre_t, post_t)
}

// DecibelsToLinear converts from decibels to linear energy (audio).
func DecibelsToLinear[T ~float32 | ~float64](db T) T { return spatial.DecibelsToLinear(db) }

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
func Ease[T ~float32 | ~float64](x, curve T) T { return spatial.Ease(x, curve) }

// Exp raises the mathematical constant e to the power of x and returns it.
// e has an approximate value of 2.71828, and can be obtained with Exp(1).
//
// For exponents to other bases use the method pow.
func Exp[T ~float32 | ~float64](x T) T { return spatial.Exp(x) }

// Floor rounds x downward (towards negative infinity), returning the largest whole number that is not
// more than x.
func Floor[T spatial.ComponentWise[T]](x T) T { return spatial.Floor(x) }

// Floorf rounds x downward (towards negative infinity), returning the largest whole number that is not
// more than x.
func Floorf[T ~float32 | ~float64](x T) T { return spatial.Floorf(x) }

// Floori rounds x downward (towards negative infinity), returning the largest whole number that is not
// more than x.
func Floori[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return x }

// Fmod returns the floating-point remainder of x divided by y, keeping the sign of x.
func Fmod[T ~float32 | ~float64](x, y T) T { return spatial.Fmod(x, y) }

// Fposmod returns the floating-point modulus of x divided by y, wrapping equally in positive and negative.
func Fposmod[T ~float32 | ~float64](x, y T) T { return spatial.Fposmod(x, y) }

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
	return spatial.IsApproximatelyEqual(a, b)
}

// IsFinite returns whether x is a finite value, i.e. it is not NaN, +INF, or -INF.
func IsFinite[T ~float32 | ~float64](x T) bool { return spatial.IsFinite(x) }

// IsInfinity returns whether x is an infinite value, i.e. it is +INF or -INF.
func IsInfinity[T ~float32 | ~float64](x T) bool { return spatial.IsInfinity(x) }

// IsNaN returns whether x is NaN (not a number).
func IsNaN[T ~float32 | ~float64](x T) bool { return spatial.IsNaN(x) }

// IsApproximatelyZero Returns true if x is zero or almost zero. The comparison is done using a
// tolerance calculation with a small internal epsilon. This function is faster than using
// [IsApproximatelyEqual] with one value as zero.
func IsApproximatelyZero[T ~float32 | ~float64](x T) bool { return spatial.IsApproximatelyZero(x) }

// Lerp linearly interpolates between two values by the factor defined in weight. To perform interpolation,
// weight should be between 0.0 and 1.0 (inclusive). However, values outside this range are allowed and can
// be used to perform extrapolation. If this is not desired, use [Clampf] on the result of this function.
//
// See also [InverseLerp] which performs the reverse of this operation. To perform eased interpolation with
// [Lerp], combine it with ease or smoothstep.
func Lerp[T spatial.Lerpable[T]](from, to T, weight Float) T {
	return spatial.Lerp(from, to, spatial.Float(weight))
}

// LerpAngle linearly interpolates between two angles (in radians) by a weight value between 0.0 and 1.0.
// Similar to lerp, but interpolates correctly when the angles wrap around [Tau]. To perform eased
// interpolation with [LerpAngle], combine it with [Ease] or [Smoothstep].
//
// Note: This function lerps through the shortest path between from and to. However, when these two angles
// are approximately Pi + k * Tau apart for any integer k, it's not obvious which way they lerp due to
// floating-point precision errors. For example, LerpAngle(0, Pi, weight) lerps counter-clockwise, while
// LerpAngle(0, Pi + 5 * Tau, weight) lerps clockwise.
func LerpAngle[T ~float32 | ~float64](from, to, weight T) T {
	return spatial.LerpAngle(from, to, weight)
}

// Lerpf linearly interpolates between two values by the factor defined in weight. To perform interpolation,
// weight should be between 0.0 and 1.0 (inclusive). However, values outside this range are allowed and can
// be used to perform extrapolation. If this is not desired, use [Clampf] on the result of this function.
//
// See also [InverseLerp] which performs the reverse of this operation. To perform eased interpolation with
// [Lerpf], combine it with ease or smoothstep.
func Lerpf[T ~float32 | ~float64](from, to, weight T) T { return spatial.Lerpf(from, to, weight) }

// LinearToDecibels converts from linear energy (audio) to decibels.
func LinearToDecibels[T ~float32 | ~float64](energy T) T { return spatial.LinearToDecibels(energy) }

// Log returns the natural logarithm of x (base e, with e being approximately 2.71828). This is the amount of
// time needed to reach a certain level of continuous growth.
//
// Note: This is not the same as the "log" function on most calculators, which uses a base 10 logarithm. To use
// base 10 logarithm, use Log(x) / Log(10).
//
// Note: The logarithm of 0 returns -inf, while negative values return -NaN.
func Log[T ~float32 | ~float64](x T) T { return spatial.Log(x) }

// MoveToward moves from toward to by the delta amount. Will not go past to.
// Use a negative delta value to move away.
func MoveToward[T ~float32 | ~float64](from, to, delta T) T {
	return spatial.MoveToward(from, to, delta)
}

// NearestPowerOfTwo returns the nearest power of two to the given value.
//
// Warning: Due to its implementation, this method returns 0 rather than 1 for values
// less than or equal to 0, with an exception for value being the smallest negative
// 64-bit integer (-9223372036854775808) in which case the value is returned unchanged.
func NearestPowerOfTwo(x Int) Int { return Int(spatial.NearestPowerOfTwo(spatial.Int(x))) }

// PingPong wraps value between 0 and the length. If the limit is reached, the next value
// the function returns is decreased to the 0 side or increased to the length side (like
// a triangle wave). If length is less than zero, it becomes positive.
func PingPong[T ~float32 | ~float64](value, length T) T { return spatial.PingPong(value, length) }

// Posmod returns the integer modulus of x divided by y that wraps equally in positive and negative.
func Posmod[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x, y T) T { return spatial.Posmod(x, y) }

// Pow returns the result of base raised to the power of exp.
func Pow[T ~float32 | ~float64](base, exp T) T { return spatial.Pow(base, exp) }

// Remap maps a value from range (istart, istop) to (ostart, ostop). See also [Lerp] and [InverseLerp].
// If value is outside (istart, istop), then the resulting value will also be outside (ostart, ostop).
// If this is not desired, use [Clamp] on the result of this function.
func Remap[T ~float32 | ~float64](value, istart, istop, ostart, ostop T) T {
	return spatial.Remap(value, istart, istop, ostart, ostop)
}

// RotateToward rotates from toward to by the delta amount. Will not go past to.
//
// Similar to move_toward, but interpolates correctly when the angles wrap around Tau.
//
// If delta is negative, this function will rotate away from to, toward the opposite angle,
// and will not go past the opposite angle.
func RotateToward[T ~float32 | ~float64](from, to, delta T) T {
	return spatial.RotateToward(from, to, delta)
}

// Round rounds x to the nearest whole number, with halfway cases rounded away from 0.
func Round[T spatial.ComponentWise[T]](x T) T { return spatial.Round(x) }

// Roundf rounds x to the nearest whole number, with halfway cases rounded away from 0.
func Roundf[T ~float32 | ~float64](x T) T { return spatial.Roundf(x) }

// Roundi rounds x to the nearest whole number, with halfway cases rounded away from 0.
func Roundi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return x }

// Sign returns the same type of value as x, with -1 for negative values, 1 for positive
// values, and 0 for zeros. For nan values it returns 0.
func Sign[T spatial.ComponentWise[T]](x T) T { return spatial.Sign(x) }

// Signf returns -1.0 if x is negative, 1.0 if x is positive, and 0.0 if x is zero. For NaN
// values of x it returns 0.0.
func Signf[T ~float32 | ~float64](x T) T { return spatial.Signf(x) }

// Signi returns -1 if x is negative, 1 if x is positive, and 0 if x is zero. For NaN values
// of x it returns 0.
func Signi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x T) T { return spatial.Signi(x) }

// Sin returns the sine of angle x in radians.
func Sin[T ~float32 | ~float64](x T) T { return spatial.Sin(x) }

// Sinh returns the hyperbolic sine of x in radians.
func Sinh[T ~float32 | ~float64](x T) T { return spatial.Sinh(x) }

// Smoothstep returns the result of smoothly interpolating the value of x between 0 and 1,
// based on the where x lies with respect to the edges from and to.
//
// The return value is 0 if x <= from, and 1 if x >= to. If x lies between from and to,
// the returned value follows an S-shaped curve that maps x between 0 and 1.
//
// This S-shaped curve is the cubic Hermite interpolator, given by
//
//	(y) = 3*y^2 - 2*y^3 where y = (x-from) / (to-from).
func Smoothstep[T ~float32 | ~float64](from, to, x T) T { return spatial.Smoothstep(from, to, x) }

// Snapped returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snapped[T spatial.ComponentWise[T]](x, step T) T { return spatial.Snapped(x, step) }

// Snappedf returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snappedf[T ~float32 | ~float64](x, step T) T {
	return T(spatial.Snapped(spatial.Float(x), spatial.Float(step)))
}

// Snappedi returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snappedi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](x, step T) T { return x }

// Sqrt returns the square root of x. Where x is negative, the result is NaN.
func Sqrt[T ~float32 | ~float64](x T) T { return spatial.Sqrt(x) }

// StepDecimals returns the position of the first non-zero digit, after the decimal point. Note that
// the maximum return value is 10, which is a design decision in the implementation.
func StepDecimals[T ~float32 | ~float64](x T) Int { return Int(spatial.StepDecimals(x)) }

// Tan returns the tangent of angle x in radians.
func Tan[T ~float32 | ~float64](x T) T { return spatial.Tan(x) }

// Tanh returns the hyperbolic tangent of x in radians.
func Tanh[T ~float32 | ~float64](x T) T { return spatial.Tanh(x) }

// Wrapi value between min and max. Can be used for creating loop-alike behavior or infinite surfaces.
func Wrapi[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](value, min, max T) T {
	return spatial.Wrapi(value, min, max)
}

// Wrapf value between min and max. Can be used for creating loop-alike behavior or infinite surfaces.
func Wrapf[T ~float32 | ~float64](value, min, max T) T { return spatial.Wrapf(value, min, max) }
