// Package Vector2 provides a 2D vector using floating-point coordinates.
package Vector2

import (
	"math"

	"grow.graphics/gd/gdmaths/Angle"
	"grow.graphics/gd/gdmaths/Float"
)

// XY is a 2-element structure that can be used to represent 2D coordinates or any
// other pair of numeric values.
//
// It uses floating-point coordinates. By default, these floating-point values use
// 32-bit precision, unlike float which is always 64-bit. If double precision is
// needed, compile with the Go build tag 'precision_double'.
//
// See [Vector2i.XY] for its integer counterpart.
type XY = struct {
	X Float.X // The vector's X component.
	Y Float.X // The vector's Y component.
}

// New constructs a new Vector2 from the given x and y.
func New(x, y float64) XY { return XY{Float.X(x), Float.X(y)} } //gd:Vector2(x: float, y: float)

type Axis int

const (
	X Axis = iota // Enumerated value for the X axis. Returned by [MaxAxis] and [MinAxis].
	Y             // Enumerated value for the Y axis. Returned by [MaxAxis] and [MinAxis].
)

var (
	Zero  = XY{0, 0}                                       // Zero vector, a vector with all components set to 0.
	One   = XY{1, 1}                                       // One vector, a vector with all components set to 1.
	Inf   = XY{Float.X(math.Inf(1)), Float.X(math.Inf(1))} // Infinity vector, a vector with all components set to math.Inf(1).
	Left  = XY{-1, 0}                                      // Left unit vector. Represents the direction of left.
	Right = XY{1, 0}                                       // Right unit vector. Represents the direction of right.
	Up    = XY{0, -1}                                      // Up unit vector. Y is down in 2D, so this vector points -Y.
	Down  = XY{0, 1}                                       // Down unit vector. Y is down in 2D, so this vector points +Y.
)

// Abs returns a new vector with all components in absolute values (i.e. positive).
func Abs(vec XY) XY { //gd:Vector2.abs
	if vec.X < 0 {
		vec.X = -vec.X
	}
	if vec.Y < 0 {
		vec.Y = -vec.Y
	}
	return vec
}

// Angle Returns this vector's angle with respect to the positive X axis, or (1, 0) vector, in radians.
//
// For example, Const(Vector2.Right).Angle() will return zero, Const(Vector2.Down).Angle() will return
// PI / 2 (a quarter turn, or 90 degrees), and Vector2(1, -1).Angle() will return -Pi / 4
// (a negative eighth turn, or -45 degrees).
//
// Illustration of the returned angle.
// https://raw.githubusercontent.com/godotengine/godot-docs/master/img/vector2_angle.png
//
// Equivalent to the result of [Atan2] when called with the vector's y and x as parameters:
//
//	Atan2(y, x).
func AngleRadians(vec XY) Angle.Radians { return Float.Atan2(vec.Y, vec.X) } //gd:Vector2.angle

// AngleBetween returns the angle between the given vectors, in radians.
//
// Illustration of the returned angle.
// https://raw.githubusercontent.com/godotengine/godot-docs/master/img/vector2_angle_to.png
func AngleBetween(a, b XY) Angle.Radians { return Float.Atan2(Cross(a, b), Dot(a, b)) } //gd:Vector2.angle_to

// AngleToPoint returns the angle between the line connecting the two points and the X axis, in radians.
// a.AngleToPoint(b) is equivalent of doing (b - a).Angle().
//
// Illustration of the returned angle.
// https://raw.githubusercontent.com/godotengine/godot-docs/master/img/vector2_angle_to_point.png
func AngleToPoint(vec, p XY) Angle.Radians { return Angle.Radians(AngleRadians(Sub(vec, p))) } //gd:Vector2.angle_to_point

// Aspect returns the aspect ratio of this vector, the ratio of x to y.
func Aspect(vec XY) float64 { return float64(vec.X / vec.Y) } //gd:Vector2.aspect

// BezierDerivative returns the derivative at the given t on the Bézier curve defined by this
// vector and the given control_1, control_2, and end points.
func BezierDerivative(vec, control1, control2, end XY, t float64) XY { //gd:Vector2.bezier_derivative
	return XY{
		Float.BezierDerivative(vec.X, control1.X, control2.X, end.X, Float.X(t)),
		Float.BezierDerivative(vec.Y, control1.Y, control2.Y, end.Y, Float.X(t)),
	}
}

// BezierInterpolate returns the point at the given t on the Bézier curve defined by this vector
// and the given control_1, control_2, and end points.
func BezierInterpolate(vec, control1, control2, end XY, t float64) XY { //gd:Vector2.bezier_interpolate
	return XY{
		Float.BezierInterpolate(vec.X, control1.X, control2.X, end.X, Float.X(t)),
		Float.BezierInterpolate(vec.Y, control1.Y, control2.Y, end.Y, Float.X(t)),
	}
}

// Bounce returns a new vector "bounced off" from a plane defined by the given normal.
func Bounce(vec, normal XY) XY { return Sub(Zero, Reflect(vec, normal)) } //gd:Vector2.bounce

// Ceil returns a new vector with all components rounded up (towards positive infinity).
func Ceil(vec XY) XY { return XY{Float.Ceil(vec.X), Float.Ceil(vec.Y)} } //gd:Vector2.ceil

// Clamp returns a new vector with all components clamped to the given min and max.
func Clamp(vec, from, to XY) XY { //gd:Vector2.clamp
	return XY{Float.Clamp(vec.X, from.X, to.X), Float.Clamp(vec.Y, from.Y, to.Y)}
}

// Cross returns the 2D analog of the cross product for this vector and with.
//
// This is the signed area of the parallelogram formed by the two vectors. If the second vector
// is clockwise from the first vector, then the cross product is the positive area. If
// counter-clockwise, the cross product is the negative area.
//
// Note: Cross product is not defined in 2D mathematically. This method embeds the 2D
// vectors in the XY plane of 3D space and uses their cross product's Z component as the analog.
func Cross(a, b XY) float64 { return float64(a.X*b.Y - a.Y*b.X) } //gd:Vector2.cross

// CubicInterpolate performs a cubic interpolation between this vector and b using pre_a and
// post_b as handles, and returns the result at position weight. weight is on the range of
// 0.0 to 1.0, representing the amount of interpolation.
func CubicInterpolate(vec, b, preA, postB XY, weight float64) XY { //Vector2.cubic_interpolate
	return XY{
		Float.CubicInterpolate(vec.X, b.X, preA.X, postB.X, Float.X(weight)),
		Float.CubicInterpolate(vec.Y, b.Y, preA.Y, postB.Y, Float.X(weight)),
	}
}

// CubicInterpolateInTime performs a cubic interpolation between this vector and b using pre_a
// and post_b as handles, and returns the result at position weight. weight is on the range of
// 0.0 to 1.0, representing the amount of interpolation.
//
// It can perform smoother interpolation than cubic_interpolate by the time values.
func CubicInterpolateInTime(vec, b, preA, postB XY, weight, b_t, pre_a_t, post_b_t float64) XY { //gd:Vector2.cubic_interpolate_in_time
	return XY{
		Float.CubicInterpolateInTime(vec.X, b.X, preA.X, postB.X, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
		Float.CubicInterpolateInTime(vec.Y, b.Y, preA.Y, postB.Y, Float.X(weight), Float.X(b_t), Float.X(pre_a_t), Float.X(post_b_t)),
	}
}

// Direction returns the normalized vector pointing from this vector to to. This is equivalent
// to using (b - a).Normalized().
func Direction(a, b XY) XY { return Normalized(Sub(b, a)) } //gd:Vector2.direction_to

// DistanceSquared returns the squared distance between this vector and to.
//
// This method runs faster than distance_to, so prefer it if you need to compare vectors or
// need the squared distance for some formula.
func DistanceSquared(a, b XY) float64 { //gd:Vector2.distance_squared_to
	return float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

// Distance returns the distance between this vector and to.
func Distance(a, b XY) float64 { //gd:Vector2.distance_to
	return float64(math.Sqrt(float64(DistanceSquared(a, b))))
}

// Dot returns the dot product of this vector and with. This can be used to compare the angle
// between two vectors. For example, this can be used to determine whether an enemy is facing
// the player.
//
// The dot product will be 0 for a straight angle (90 degrees), greater than 0 for angles narrower
// than 90 degrees and lower than 0 for angles wider than 90 degrees.
//
// When using unit (normalized) vectors, the result will always be between -1.0 (180 degree angle)
// when the vectors are facing opposite directions, and 1.0 (0 degree angle) when the vectors are aligned.
//
// Note: Vector2.Dot(a,b) is equivalent to Vector2.Dot(b,a)
func Dot(a, b XY) float64 { return float64(a.X*b.X + a.Y*b.Y) } //gd:Vector2.dot

// Floor returns a new vector with all components rounded down (towards negative infinity).
func Floor(vec XY) XY { //gd:Vector2.floor
	return XY{Float.Floor(vec.X), Float.Floor(vec.Y)}
}

// IsAproximatelyEqual returns true if this vector and to are approximately equal, by running
// [IsAproximatelyEqual] on each component.
func IsApproximatelyEqual(a, b XY) bool { //gd:Vector2.is_equal_approx
	return Float.IsApproximatelyEqual(a.X, b.X) && Float.IsApproximatelyEqual(a.Y, b.Y)
}

// IsFinite returns true if this vector is finite, by calling [IsFinite] on each component.
func IsFinite(vec XY) bool { return Float.IsFinite(vec.X) && Float.IsFinite(vec.Y) } //gd:Vector2.is_finite

// IsNormalized returns true if the vector is normalized, i.e. its length is approximately equal to 1.
func IsNormalized(vec XY) bool { return Float.IsApproximatelyEqual(LengthSquared(vec), 1) } //gd:Vector2.is_normalized

// IsAproximatelyZero returns true if this vector is approximately equal to Vector2.Zero.
func IsApproximatelyZero(vec XY) bool { return IsApproximatelyEqual(vec, Zero) } //gd:Vector2.is_zero_approx

// Length returns the length (magnitude) of this vector.
func Length(vec XY) float64 { //gd:Vector2.length
	return float64(math.Sqrt(float64(LengthSquared(vec))))
}

// LengthSquared returns the squared length (squared magnitude) of this vector.
func LengthSquared(vec XY) float64 { return float64(vec.X*vec.X + vec.Y*vec.Y) } //gd:Vector2.length_squared

// Lerp returns the result of the linear interpolation between this vector and to by amount weight. weight
// is on the range of 0.0 to 1.0, representing the amount of interpolation.
func Lerp(vec, to XY, weight float64) XY { //gd:Vector2.lerp
	return XY{Float.Lerp(vec.X, to.X, Float.X(weight)), Float.Lerp(vec.Y, to.Y, Float.X(weight))}
}

// LengthLimited returns the vector with a maximum length by limiting its length to length.
func LengthLimited(vec XY, length float64) XY { //gd:Vector2.limit_length
	var l = Length(vec)
	if l > 0 && length < l {
		vec = Mulf(vec, 1/l)
		vec = Mulf(vec, length)
	}
	return vec
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all
// components are equal, this method returns [X].
func MaxAxis(vec XY) Axis { //gd:Vector2.max_axis_index
	if vec.Y > vec.X {
		return Y
	}
	return X
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all
// components are equal, this method returns [Y].
func MinAxis(vec XY) Axis { //gd:Vector2.min_axis_index
	if vec.X < vec.Y {
		return X
	}
	return Y
}

// Move returns a new vector moved toward to by the fixed delta amount. Will not go past
// the final value.
func Move(a, b XY, delta float64) XY { //gd:Vector2.move_toward
	var vd = Sub(b, a)
	var len = Length(vd)
	if len <= delta || len < Float.Epsilon {
		return b
	}
	return Add(a, Mulf(vd, len*delta))
}

// Normalized returns the result of scaling the vector to unit length. Equivalent to v / v.Length().
// See also is_normalized.
//
// Note: This function may return incorrect values if the input vector length is near zero.
func Normalized(vec XY) XY { //gd:Vector2.normalized
	length := Length(vec)
	if length == 0 {
		return Zero
	}
	return XY{Float.X(float64(vec.X) / length), Float.X(float64(vec.Y) / length)}
}

// Orthogonal returns a perpendicular vector rotated 90 degrees counter-clockwise compared to the original,
// with the same length.
func Orthogonal(vec XY) XY { return XY{vec.Y, -vec.X} } //gd:Vector2.orthogonal

// Posmode returns a vector composed of the [Fposmod] of this vector's components and [Mod].
func Posmod(vec XY, mod float64) XY { //gd:Vector2.posmod
	return XY{Float.Posmod(vec.X, Float.X(mod)), Float.Posmod(vec.Y, Float.X(mod))}
}

// Posmod returns a vector composed of the [Fposmod] of this vector's components and [Mod].
func PosmodVector(vec, mod XY) XY { //gd:Vector2.posmodv
	return XY{Float.Posmod(vec.X, mod.X), Float.Posmod(vec.Y, mod.Y)}
}

// Project returns the result of projecting the vector onto the given vector b.
func Project(a, b XY) XY { //gd:Vector2.project
	return Mulf(b, Dot(a, b)/LengthSquared(b))
}

// Reflect returns the result of reflecting the vector from a line defined by the given direction vector n.
func Reflect(v, n XY) XY { //gd:Vector2.reflect
	return Sub(Mulf(n, 2*Dot(v, n)), v)
}

// Rotated returns the result of rotating this vector by angle (in radians).
func Rotated(v XY, by Angle.Radians) XY { //gd:Vector2.rotated
	var cs = Float.X(Float.Cos(Float.X(by)))
	var sn = Float.X(Float.Sin(Float.X(by)))
	return XY{v.X*cs - v.Y*sn, v.X*sn + v.Y*cs}
}

// Round returns a new vector with all components rounded to the nearest integer, with halfway cases
// rounded away from zero.
func Round(v XY) XY { //gd:Vector2.round
	return XY{Float.Round(v.X), Float.Round(v.Y)}
}

// Sign returns a new vector with each component set to 1.0 if it's positive, -1.0 if it's negative,
// and 0.0 if it's zero. The result is identical to calling [Signf] on each component.
func Sign(vec XY) XY { //Vector2.sign
	return XY{Float.Sign(vec.X), Float.Sign(vec.Y)}
}

// Slerp returns the result of spherical linear interpolation between this vector and to, by amount weight.
// weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
//
// This method also handles interpolating the lengths if the input vectors have different lengths. For the
// special case of one or both input vectors having zero length, this method behaves like lerp.
func Slerp(vec, to XY, weight Angle.Radians) XY { //gd:Vector2.slerp
	var start_length_sq = LengthSquared(vec)
	var end_length_sq = LengthSquared(to)
	if start_length_sq == 0.0 || end_length_sq == 0.0 {
		// Zero length vectors have no angle, so the best we can do is either lerp or throw an error.
		return Lerp(vec, to, float64(weight))
	}
	var start_length = Float.Sqrt(start_length_sq)
	var result_length = Float.Lerp(start_length, Float.Sqrt(end_length_sq), float64(weight))
	var angle = AngleBetween(vec, to)
	return Mulf(Rotated(vec, angle*weight), result_length/start_length)
}

// Slide returns the result of sliding the vector along a plane defined by the given normal.
func Slide(v, n XY) XY { //gd:Vector2.slide
	return Sub(v, Mulf(n, Dot(v, n)))
}

// Snapped returns a new vector with all components snapped to the nearest multiple of step.
func Snapped(v, step XY) XY { //gd:Vector2.snapped
	return XY{
		Float.Snapped(v.X, step.X),
		Float.Snapped(v.Y, step.Y),
	}
}

func Add(a, b XY) XY { return XY{a.X + b.X, a.Y + b.Y} } //gd:(Vector2 + Vector2)
func Sub(a, b XY) XY { return XY{a.X - b.X, a.Y - b.Y} } //gd:(Vector2 - Vector2)
func Mul(a, b XY) XY { return XY{a.X * b.X, a.Y * b.Y} } //gd:(Vector2 * Vector2)
func Div(a, b XY) XY { return XY{a.X / b.X, a.Y / b.Y} } //gd:(Vector2 / Vector2)

func Addf(a XY, b float64) XY { //gd:(Vector2 + float)
	return XY{a.X + Float.X(b), a.Y + Float.X(b)}
}
func Subf(a XY, b float64) XY { //gd:(Vector2 - float)
	return XY{a.X - Float.X(b), a.Y - Float.X(b)}
}
func Mulf(a XY, b float64) XY { //gd:(Vector2 * float)
	return XY{a.X * Float.X(b), a.Y * Float.X(b)}
}
func Divf(a XY, b float64) XY { //gd:(Vector2 / float)
	return XY{a.X / Float.X(b), a.Y / Float.X(b)}
}

func Neg(v XY) XY { return XY{-v.X, -v.Y} } //gd:(-Vector2)

/*
func (v Vector2) Transform(t Transform2D) Vector2 { return Vector2{t.tdotx(v), t.tdoty(v)}.Add(t[2]) } //Transform2D * Vector2D
*/
