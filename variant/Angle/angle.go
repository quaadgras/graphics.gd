package Angle

import (
	"math"

	"grow.graphics/gd/variant/Float"
)

const Pi Radians = math.Pi
const Tau Radians = 2 * Pi

type Radians Float.X
type Degrees Float.X

type Euler3D = struct {
	X Radians
	Y Radians
	Z Radians
}

type vector2 = struct {
	X Float.X
	Y Float.X
}

// Acos returns the arc cosine of x in radians. Use to get the angle of cosine x.
// x will be clamped between -1.0 and 1.0 (inclusive), in order to prevent acos
// from returning NaN.
func Acos[X Float.Any](x X) Radians { //gd:acos
	return Radians(math.Acos(float64(Float.Clamp(x, -1.0, 1.0))))
}

// Asin returns the arc sine of x in radians. Use to get the angle of sine x. x will be clamped
// between -1.0 and 1.0 (inclusive), in order to prevent asin from returning NaN.
func Asin[X Float.Any](x X) Radians { //gd:asin
	return Radians(math.Asin(float64(Float.Clamp(x, -1.0, 1.0))))
}

// Difference returns the difference between the two angles, in the range of [-Pi, +Pi].
// When from and to are opposite, returns -Pi if from is smaller than to, or Pi otherwise.
func Difference(from, to Radians) Radians { //gd:angle_difference
	difference := Float.Mod(from, to)
	return Float.Mod(2*difference, Tau) - difference
}

// Tan returns the tangent of angle x in radians.
func Tan(x Radians) Float.X { return Float.X(math.Tan(float64(x))) } //tan

// Atan returns the arc tangent of x in radians. Use it to get the angle from an angle's tangent in
// trigonometry. The method cannot know in which quadrant the angle should fall. See atan2 if you
// have both y and x.
func Atan[X Float.Any](x X) Radians { //atan
	return Radians(math.Atan(float64(x)))
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

func InRadians(deg Degrees) Radians { return Radians(deg) * (Pi / 180.0) } //deg_to_rad
func InDegrees(rad Radians) Degrees { return Degrees(rad * (180.0 / Pi)) } //rad_to_deg
