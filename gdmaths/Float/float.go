// Package Float provides a generic math library for floating-point numbers.
package Float

import (
	"math"

	"grow.graphics/gd/gdmaths/Angle"
)

const Epsilon = 0.00001

// ʕ is a little ternary operator for porting C code.
func ʕ[T any](q bool, a T, b T) T {
	if q {
		return a
	}
	return b
}

// Abs returns the absolute value of the float parameter x (i.e. non-negative value).
func Abs[X ~float32 | ~float64](x X) X { return X(math.Abs(float64(x))) } //gd:absf

// Atan2 returns the arc tangent of y/x in radians. Use to get the angle of tangent y/x. To compute
// the value, the method takes into account the sign of both arguments in order to determine the quadrant.
//
// Important note: The Y coordinate comes first, by convention.
func Atan2[T ~float32 | ~float64](y, x T) Angle.Radians { //gd:atan2
	return Angle.Radians(math.Atan2(float64(y), float64(x)))
}

// BezierDerivative returns the derivative at the given t on a one-dimensional Bézier curve defined by the given
// control_1, control_2, and end points.
func BezierDerivative[T ~float32 | ~float64](start, control_1, control_2, end, t T) T { //gd:bezier_derivative
	/* Formula from Wikipedia article on Bezier curves. */
	omt := (1.0 - t)
	omt2 := omt * omt
	t2 := t * t
	return T((control_1-start)*3.0*omt2 + (control_2-control_1)*6.0*omt*t + (end-control_2)*3.0*t2)
}

// BezierInterpolate returns the point at the given t on a one-dimensional Bézier curve defined by the given
// control_1, control_2, and end points.
func BezierInterpolate[T ~float32 | ~float64](start, control_1, control_2, end, t T) T { //gd:bezier_interpolate
	/* Formula from Wikipedia article on Bezier curves. */
	omt := (1.0 - t)
	omt2 := omt * omt
	omt3 := omt2 * omt
	t2 := t * t
	t3 := t2 * t
	return start*omt3 + control_1*omt2*t*3.0 + control_2*omt*t2*3.0 + end*t3
}

// Ceil rounds x upward (towards positive infinity), returning the smallest whole number that is not less than x.
func Ceil[X float32 | float64](x X) X { return X(math.Ceil(float64(x))) } //gd:ceilf

// Clamp clamps the value, returning a Variant not less than min and not more than max. Any values that can be
// compared with the less than and greater than operators will work.
func Clamp[X float32 | float64](value, min, max X) X { //clamp
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Cos returns the cosine of angle x in radians.
func Cos[X ~float32 | ~float64](x X) Angle.Radians { return Angle.Radians(math.Cos(float64(x))) } //gd:cos

// CubicInterpolate cubic interpolates between two values by the factor defined in weightSee also
// with pre and post values.
func CubicInterpolate[X ~float32 | ~float64](from, to, pre, post, weight X) X { //gd:cubic_interpolate
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
func CubicInterpolateInTime[X ~float32 | ~float64](from, to, pre, post, weight, to_t, pre_t, post_t X) X { //gd:cubic_interpolate_in_time
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
func Floor[X ~float32 | ~float64](x X) X { return X(math.Floor(float64(x))) } //gd:floorf

// IsApproximatelyEqual returns true if a and b are approximately equal to each other.
//
// Here, "approximately equal" means that a and b are within a small internal epsilon of each other, which scales
// with the magnitude of the numbers.
//
// Infinity values of the same sign are considered equal.
func IsApproximatelyEqual[X ~float32 | ~float64](a, b X) bool { //gd:is_equal_approx
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

// IsFinite returns whether x is a finite value, i.e. it is not NaN, +INF, or -INF.
func IsFinite[X ~float32 | ~float64](x X) bool { //gd:is_finite
	return !math.IsInf(float64(x), 0) && !math.IsNaN(float64(x))
}

// Lerpf linearly interpolates between two values by the factor defined in weight. To perform interpolation,
// weight should be between 0.0 and 1.0 (inclusive). However, values outside this range are allowed and can
// be used to perform extrapolation. If this is not desired, use [Clampf] on the result of this function.
//
// See also [InverseLerp] which performs the reverse of this operation. To perform eased interpolation with
// [Lerpf], combine it with ease or smoothstep.
func Lerp[X ~float32 | ~float64](from, to, weight X) X { return from + (to-from)*weight } //gd:lerpf

// Fmod returns the floating-point remainder of x divided by y, keeping the sign of x.
func Mod[X ~float32 | ~float64](x, y X) X { //gd:fmod
	return X(math.Mod(float64(x), float64(y)))
}

// Fposmod returns the floating-point modulus of x divided by y, wrapping equally in positive and negative.
func Posmod[X ~float32 | ~float64](x, y X) X { //gd:fposmod
	value := Mod(x, y)
	if ((value < 0) && (y > 0)) || ((value > 0) && (y < 0)) {
		value += y
	}
	value += 0.0
	return value
}

// Round rounds x to the nearest whole number, with halfway cases rounded away from 0.
func Round[X ~float32 | ~float64](x X) X { return X(math.Round(float64(x))) } //gd:roundf

// Sign returns -1.0 if x is negative, 1.0 if x is positive, and 0.0 if x is zero. For NaN
// values of x it returns 0.0.
func Sign[X ~float32 | ~float64](x X) X { //gd:signf
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

// Sin returns the sine of angle x in radians.
func Sin[X ~float32 | ~float64](x X) Angle.Radians { return Angle.Radians(math.Sin(float64(x))) } //gd:sin

// Snapped returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snapped[X ~float32 | ~float64](x, step X) X { //gd:snappedf
	if step != 0 {
		x = X(math.Floor(float64(x/step)+0.5)) * step
	}
	return x
}

// Sqrt returns the square root of x. Where x is negative, the result is NaN.
func Sqrt[X ~float32 | ~float64](x X) X { return X(math.Sqrt(float64(x))) } //gd:sqrt
