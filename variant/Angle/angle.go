package Angle

import (
	"math"

	"graphics.gd/variant/Float"
)

const Pi Radians = math.Pi
const Tau Radians = 2 * Pi

type Radians Float.X
type Degrees Float.X

type Order int

const (
	OrderXYZ Order = 0
	OrderXZY Order = 1
	OrderYXZ Order = 2
	OrderYZX Order = 3
	OrderZXY Order = 4
	OrderZYX Order = 5
)

type Direction int

const (
	Clockwise        Direction = 0
	CounterClockwise Direction = 1
)

type vector2 = struct {
	X Float.X
	Y Float.X
}

// ʕ is a little ternary operator for porting C code.
func ʕ[T any](q bool, a T, b T) T {
	if q {
		return a
	}
	return b
}

// Acos returns the arc cosine of x in radians. Use to get the angle of cosine x.
// x will be clamped between -1.0 and 1.0 (inclusive), in order to prevent acos
// from returning NaN.
func Acos[X Float.Any](x X) Radians { //gd:acos
	return Radians(math.Acos(float64(Float.Clamp(x, -1.0, 1.0))))
}

// Acosh returns the hyperbolic arc cosine of x.
func Acosh[X Float.Any](x X) Radians { //gd:acosh
	return Radians(math.Acosh(float64(x)))
}

// Asin returns the arc sine of x in radians. Use to get the angle of sine x. x will be clamped
// between -1.0 and 1.0 (inclusive), in order to prevent asin from returning NaN.
func Asin[X Float.Any](x X) Radians { //gd:asin
	return Radians(math.Asin(float64(Float.Clamp(x, -1.0, 1.0))))
}

// Asinh returns the hyperbolic arc sine of x.
func Asinh[X Float.Any](x X) Radians { //gd:asinh
	return Radians(math.Asinh(float64(x)))
}

// Difference returns the difference between the two angles, in the range of [-Pi, +Pi].
// When from and to are opposite, returns -Pi if from is smaller than to, or Pi otherwise.
func Difference(from, to Radians) Radians { //gd:angle_difference
	diff := Float.Mod(to-from, Tau)
	return Float.Mod(2*diff, Tau) - diff
}

// Tan returns the tangent of angle x in radians.
func Tan(x Radians) Float.X { return Float.X(math.Tan(float64(x))) } //gd:tan

// Tanh returns the hyperbolic tangent of x.
func Tanh(x Radians) Float.X { //gd:tanh
	return Float.X(math.Tanh(float64(x)))
}

// Atan returns the arc tangent of x in radians. Use it to get the angle from an angle's tangent in
// trigonometry. The method cannot know in which quadrant the angle should fall. See atan2 if you
// have both y and x.
func Atan[X Float.Any](x X) Radians { //gd:atan
	return Radians(math.Atan(float64(x)))
}

// Atanh returns the hyperbolic arc tangent of x.
func Atanh[X Float.Any](x X) Radians { //gd:atanh
	return Radians(math.Atanh(float64(x)))
}

// Atan2 returns the arc tangent of y/x in radians. Use to get the angle of tangent y/x. To compute
// the value, the method takes into account the sign of both arguments in order to determine the quadrant.
//
// Important note: The Y coordinate comes first, by convention.
func Atan2[X Float.Any](y, x X) Radians { //gd:atan2
	return Radians(math.Atan2(float64(y), float64(x)))
}

// Cos returns the cosine of angle x in radians.
func Cos(x Radians) Float.X { return Float.X(math.Cos(float64(x))) } //gd:cos

// Sin returns the sine of angle x in radians.
func Sin(x Radians) Float.X { return Float.X(math.Sin(float64(x))) } //gd:sin

// Sinh returns the hyperbolic sine of x.
func Sinh[X Float.Any](x X) X { return X(math.Sinh(float64(x))) } //gd:sinh

// Cosh returns the hyperbolic cosine of x.
func Cosh[X Float.Any](x X) X { return X(math.Cosh(float64(x))) } //gd:cosh

// Lerp linearly interpolates between two angles (in radians) by a weight value between 0.0 and 1.0.
// Similar to lerp, but interpolates correctly when the angles wrap around [Tau]. To perform eased
// interpolation with [Lerp], combine it with [Float.Ease] or [Float.Smoothstep].
//
// Note: This function lerps through the shortest path between from and to. However, when these two angles
// are approximately Pi + k * Tau apart for any integer k, it's not obvious which way they lerp due to
// floating-point precision errors. For example, Lerp(0, Pi, weight) lerps counter-clockwise, while
// Lerp(0, Pi + 5 * Tau, weight) lerps clockwise.
func Lerp[X Float.Any](from, to Radians, weight X) Radians { //gd:lerp_angle
	return from + Difference(from, to)*Radians(weight)
}

// AsVector2 creates a unit Vector2 rotated to the given angle in radians. This is equivalent
// to doing:
//
//	Vector2.New(math.Cos(angle), math.Sin(angle))
//	Vector2.Rotated(Vector2.Right, angle).
func (angle Radians) AsVector2() vector2 { //gd:Vector2.from_angle
	return vector2{float(math.Cos(float64(angle))), float(math.Sin(float64(angle)))}
}

func InRadians(deg Degrees) Radians { return Radians(deg) * (Pi / 180.0) } //gd:deg_to_rad
func InDegrees(rad Radians) Degrees { return Degrees(rad * (180.0 / Pi)) } //gd:rad_to_deg

// CubicInterpolateAngle cubic interpolates between two rotation values with shortest path
// by the factor defined in weight with pre and post values. See also [Lerp].
func CubicInterpolate(from, to, pre, post, weight Radians) Radians { //gd:cubic_interpolate_angle
	from_rot := Float.Mod(from, Tau)
	var (
		pre_diff = Float.Mod(pre-from_rot, Tau)
		pre_rot  = from_rot + Float.Mod(2.0*pre_diff, Tau) - pre_diff
	)
	var (
		to_diff = Float.Mod(to-from_rot, Tau)
		to_rot  = from_rot + Float.Mod(2.0*to_diff, Tau) - to_diff
	)
	var (
		post_diff = Float.Mod(post-to_rot, Tau)
		post_rot  = to_rot + Float.Mod(2.0*post_diff, Tau) - post_diff
	)
	return CubicInterpolate(from_rot, to_rot, pre_rot, post_rot, weight)
}

// CubicInterpolateAngleInTime cubic interpolates between two rotation values with shortest path
// by the factor defined in weight with pre and post values. See also [Lerp].
//
// It can perform smoother interpolation than [CubicInterpolate] by the time values.
func CubicInterpolateAngleInTime(from, to, pre, post, weight, to_t, pre_t, post_t Radians) Radians { //gd:cubic_interpolate_angle_in_time
	from_rot := Float.Mod(from, Tau)
	var (
		pre_diff = Float.Mod(pre-from_rot, Tau)
		pre_rot  = from_rot + Float.Mod(2.0*pre_diff, Tau) - pre_diff
	)
	var (
		to_diff = Float.Mod(to-from_rot, Tau)
		to_rot  = from_rot + Float.Mod(2.0*to_diff, Tau) - to_diff
	)
	var (
		post_diff = Float.Mod(post-to_rot, Tau)
		post_rot  = to_rot + Float.Mod(2.0*post_diff, Tau) - post_diff
	)
	return Float.CubicInterpolateInTime(from_rot, to_rot, pre_rot, post_rot, weight, to_t, pre_t, post_t)
}

// RotateToward rotates from toward to by the delta amount. Will not go past to.
//
// Similar to move_toward, but interpolates correctly when the angles wrap around Tau.
//
// If delta is negative, this function will rotate away from to, toward the opposite angle,
// and will not go past the opposite angle.
func RotateToward[X Float.Any](from, to Radians, delta X) Radians { //gd:rotate_toward
	var difference = Difference(from, to)
	var abs_difference = Float.Abs(difference)
	// When `delta < 0` move no further than to Pi radians away from `to` (as Pi is the max possible angle distance).
	return from + Float.Clamp(Radians(delta), abs_difference-Pi, abs_difference)*(ʕ[Radians](difference >= 0.0, 1.0, -1.0))
}
