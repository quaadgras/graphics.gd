// Package Vector2i provides a 2D vector using integer coordinates.
package Vector2i

import (
	"math"

	"grow.graphics/gd/gdmaths/Float"
	"grow.graphics/gd/gdmaths/Int"
)

// A 2-element structure that can be used to represent 2D grid
// coordinates or any other pair of integers.
//
// It uses integer coordinates and is therefore preferable to
// Vector2 when exact precision is required. Note that the values
// are limited to 32 bits, and unlike Vector2 this cannot be
// configured with an engine build option. Use int, int64 or
// PackedInt64Array if 64-bit values are needed.
type XY = struct {
	X int32
	Y int32
}

type Axis int

const (
	X Axis = iota // Enumerated value for the X axis. Returned by [MaxAxis] and [MinAxis].
	Y Axis = iota // Enumerated value for the Y axis. Returned by [MaxAxis] and [MinAxis].
)

var (
	Zero  = XY{0, 0}                         // Zero vector, a vector with all components set to 0.
	One   = XY{1, 1}                         // One vector, a vector with all components set to 1.
	MinXY = XY{math.MinInt32, math.MinInt32} // Min vector, a vector with all components equal to [math.MinInt32]. Can be used as a negative integer equivalent of [Vector2.Inf].
	MaxXY = XY{math.MaxInt32, math.MaxInt32} // Max vector, a vector with all components equal to [math.MaxInt32]. Can be used as a positive integer equivalent of [Vector2.Inf].
	Left  = XY{-1, 0}                        // Left vector, a unit vector pointing left (-1, 0).
	Right = XY{1, 0}                         // Right vector, a unit vector pointing right (1, 0).
	Up    = XY{0, -1}                        // Up vector, a unit vector pointing up (0, -1).
	Down  = XY{0, 1}                         // Down vector, a unit vector pointing down (0, 1).
)

// New constructs a new Vector2i from the given x and y.
func New(x, y int) XY { return XY{int32(x), int32(y)} } //gd:Vector2i(x:int,y:int)

// Abs returns a new vector with all components in absolute values (i.e. positive).
func Abs(v XY) XY { return XY{Int.Abs(v.X), Int.Abs(v.Y)} } //gd:Vector2i.abs

// Aspect returns the aspect ratio of this vector, the ratio of x to y.
func Aspect(v XY) Float.X { return Float.X(v.X) / Float.X(v.Y) } //gd:Vector2i.aspect

// Clamp returns a new vector with all components clamped between the components of min
// and max, by running clamp on each component.
func Clamp(vec, min, max XY) XY { //gd:Vector2i.clamp
	return XY{
		Int.Clamp(vec.X, min.X, max.X),
		Int.Clamp(vec.Y, min.Y, max.Y),
	}
}

// Clampi returns a new vector with all components clamped between the components of min
// and max.
func Clampi[X Int.Any](vec XY, min, max X) XY { //gd:Vector2i.clamp
	return XY{
		Int.Clamp(vec.X, int32(min), int32(max)),
		Int.Clamp(vec.Y, int32(min), int32(max)),
	}
}

// DistanceSquared returns the squared distance between this vector and to.
//
// This method runs faster than distance_to, so prefer it if you need to compare vectors or
// need the squared distance for some formula.
func DistanceSquared(a, b XY) int { //gd:Vector2i.distance_squared_to
	return int(a.X-b.X)*int(a.X-b.X) + int(a.Y-b.Y)*int(a.Y-b.Y)
}

// Distance returns the distance between this vector and to.
func Distance(a, b XY) Float.X { //gd:Vector2i.distance_to
	return Float.X(math.Sqrt(float64(DistanceSquared(a, b))))
}

// Length returns the length (magnitude) of this vector.
func Length(vec XY) Float.X { //gd:Vector2i.length
	return Float.X(math.Sqrt(float64(LengthSquared(vec))))
}

// LengthSquared returns the squared length (squared magnitude) of this vector.
func LengthSquared(vec XY) Float.X { //gd:Vector2i.length_squared
	return Float.X(vec.X)*Float.X(vec.X) + Float.X(vec.Y)*Float.X(vec.Y)
}

// Max returns the component-wise maximum of this and with.
func Max(a, b XY) XY { //gd:Vector2i.max
	return XY{max(a.X, b.X), max(a.Y, b.Y)}
}

// Max returns the component-wise maximum of this and with.
func Maxi[X Int.Any](a XY, b X) XY { //gd:Vector2i.maxi
	return XY{max(a.X, int32(b)), max(a.Y, int32(b))}
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all
// components are equal, this method returns [X].
func MaxAxis(vec XY) Axis { //gd:Vector2i.max_axis_index
	if vec.Y > vec.X {
		return Y
	}
	return X
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all
// components are equal, this method returns [Y].
func MinAxis(vec XY) Axis { //gd:Vector2i.min_axis_index
	if vec.X < vec.Y {
		return X
	}
	return Y
}

// Min returns the component-wise minimum of this and with.
func Min(a, b XY) XY { //gd:Vector2i.min
	return XY{min(a.X, b.X), min(a.Y, b.Y)}
}

// Mini returns the component-wise minimum of this and with.
func Mini[X Int.Any](a XY, b X) XY { //gd:Vector2i.mini
	return XY{min(a.X, int32(b)), min(a.Y, int32(b))}
}

// Sign returns a new vector with each component set to 1 if it's positive, -1 if it's
// negative, and 0 if it's zero. The result is identical to calling [Signi] on each component.
func Sign(v XY) XY { return XY{Int.Sign(v.X), Int.Sign(v.Y)} } //gd:Vector2i.sign

// Snapped returns a new vector with each component snapped to the closest multiple of the corresponding
// component in step.
func Snapped(v, step XY) XY { //gd:Vector2i.snapped
	return XY{
		Int.Snapped(v.X, step.X),
		Int.Snapped(v.Y, step.Y),
	}
}

// Snappedi returns a new vector with each component snapped to the closest multiple of the corresponding
// component in step.
func Snappedi[X Int.Any](v XY, step X) XY { //gd:Vector2i.snappedi
	return XY{
		Int.Snapped(v.X, int32(step)),
		Int.Snapped(v.Y, int32(step)),
	}
}

func Add(a, b XY) XY { return XY{a.X + b.X, a.Y + b.Y} } //gd:Vector2i+(right:Vector2i)
func Sub(a, b XY) XY { return XY{a.X - b.X, a.Y - b.Y} } //gd:operator-(right:Vector2i)
func Mul(a, b XY) XY { return XY{a.X * b.X, a.Y * b.Y} } //gd:Vector2i*(right:Vector2i)
func Div(a, b XY) XY { return XY{a.X / b.X, a.Y / b.Y} } //gd:Vector2i/(right:Vector2i)
func Mod(a, b XY) XY { return XY{a.X % b.X, a.Y % b.Y} } //gd:Vector2i%(right:Vector2i)

func Addi[X Int.Any](a XY, b X) XY { return XY{a.X + int32(b), a.Y + int32(b)} } //gd:Vector2i+(right:int)
func Subi[X Int.Any](a XY, b X) XY { return XY{a.X - int32(b), a.Y - int32(b)} } //gd:Vector2i-(right:int)
func Muli[X Int.Any](a XY, b X) XY { return XY{a.X * int32(b), a.Y * int32(b)} } //gd:Vector2i*(right:int)
func Divi[X Int.Any](a XY, b X) XY { return XY{a.X / int32(b), a.Y / int32(b)} } //gd:Vector2i/(right:int)
func Modi[X Int.Any](a XY, b X) XY { return XY{a.X % int32(b), a.Y % int32(b)} } //gd:Vector2i%(right:int)

func Addf[X Float.Any](a XY, b X) XY { return XY{a.X + int32(b), a.Y + int32(b)} } //gd:Vector2i+(right:float)
func Subf[X Float.Any](a XY, b X) XY { return XY{a.X - int32(b), a.Y - int32(b)} } //gd:Vector2i-(right:float)
func Mulf[X Float.Any](a XY, b X) XY { return XY{a.X * int32(b), a.Y * int32(b)} } //gd:Vector2i*(right:float)
func Divf[X Float.Any](a XY, b X) XY { return XY{a.X / int32(b), a.Y / int32(b)} } //gd:Vector2i/(right:float)

func Neg(v XY) XY { return XY{-v.X, -v.Y} } //gd:Vector2i-(unary)

func Index[I Int.Any](v XY, i I) int { //gd:Vector2i[](index:int)
	switch Axis(i) {
	case X:
		return int(v.X)
	case Y:
		return int(v.Y)
	default:
		panic("index out of range")
	}
}
