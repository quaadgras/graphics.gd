// Package Vector4i provides a 4D vector using integer coordinates.
package Vector4i

import (
	"math"

	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
)

// XYZW is a 4-element structure that can be used to represent 4D grid
// coordinates or any other quadruplet of integers.
//
// It uses integer coordinates and is therefore preferable to Vector4
// when exact precision is required. Note that the values are limited
// to 32 bits, and unlike Vector4 this cannot be configured with an
// engine build option. Use int or PackedInt64Array if 64-bit values
// are needed.
type XYZW = struct {
	X int32 // The vector's X component.
	Y int32 // The vector's Y component.
	Z int32 // The vector's Z component.
	W int32 // The vector's W component.
}

type Axis int

const (
	X Axis = iota // Enumerated value for the X axis. Returned by [MaxAxis] and [MinAxis].
	Y             // Enumerated value for the Y axis. Returned by [MaxAxis] and [MinAxis].
	Z             // Enumerated value for the Z axis. Returned by [MaxAxis] and [MinAxis].
	W             // Enumerated value for the W axis. Returned by [MaxAxis] and [MinAxis].
)

var (
	Zero    = XYZW{0, 0, 0, 0}                                                 // Zero vector, a vector with all components set to 0.
	One     = XYZW{1, 1, 1, 1}                                                 // One vector, a vector with all components set to 1.
	MinXYZW = XYZW{math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32} // Min vector, a vector with all components equal to [math.MinInt32]. Can be used as a negative integer equivalent of [Vector4.Inf].
	MaxXYZW = XYZW{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32} // Max vector, a vector with all components equal to [math.MaxInt32]. Can be used as a positive integer equivalent of [Vector4.Inf].
)

// New returns a [XYZW] with the given components.
func New[X Int.Any | Float.Any](x, y, z, w X) XYZW { //gd:Vector4i(x:float,y:float,z:float,w:float)
	return XYZW{int32(x), int32(y), int32(z), int32(w)}
}

// Abs returns a new vector with all components in absolute values (i.e. positive).
func Abs(v XYZW) XYZW { //gd:Vector4i.abs
	return XYZW{
		Int.Abs(v.X), Int.Abs(v.Y), Int.Abs(v.Z), Int.Abs(v.W),
	}
}

// Clamp returns a new vector with all components clamped between the components of min and max,
// by running [Int.Clamp] on each component.
func Clamp(v, min, max XYZW) XYZW { //gd:Vector4i.clamp
	return XYZW{
		Int.Clamp(v.X, min.X, max.X),
		Int.Clamp(v.Y, min.Y, max.Y),
		Int.Clamp(v.Z, min.Z, max.Z),
		Int.Clamp(v.W, min.W, max.W),
	}
}

// Clampi returns a new vector with all components clamped between the components of min and max,
// by running [Int.Clamp] on each component.
func Clampi[X Int.Any](v XYZW, min, max X) XYZW { //gd:Vector4i.clampi
	return XYZW{
		Int.Clamp(v.X, int32(min), int32(max)),
		Int.Clamp(v.Y, int32(min), int32(max)),
		Int.Clamp(v.Z, int32(min), int32(max)),
		Int.Clamp(v.W, int32(min), int32(max)),
	}
}

// DistanceSquared returns the squared distance between this vector and to.
//
// This method runs faster than [Vector3.DistanceTo], so prefer it if you need to compare vectors or need the squared distance for
// some formula.
func DistanceSquared(v, to XYZW) int { //gd:Vector4i.distance_squared_to
	return LengthSquared(Sub(v, to))
}

// Distance returns the distance between this vector and to.
func Distance(v, to XYZW) Float.X { //gd:Vector4i.distance_to
	return Length(Sub(v, to))
}

// Length the length (magnitude) of this vector.
func Length(v XYZW) Float.X { //gd:Vector4i.length
	return Float.X(math.Sqrt(float64(LengthSquared(v))))
}

// LengthSquared returns the squared length (squared magnitude) of this vector.
//
// This method runs faster than length, so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func LengthSquared(v XYZW) int { //gd:Vector4i.length_squared
	return int(v.X)*int(v.X) + int(v.Y)*int(v.Y) + int(v.Z)*int(v.Z) + int(v.W)*int(v.W)
}

// Max returns the component-wise maximum of this and with.
func Max(a, b XYZW) XYZW { //gd:Vector4i.max
	return XYZW{max(a.X, b.X), max(a.Y, b.Y), max(a.Z, b.Z), max(a.W, b.W)}
}

// Max returns the component-wise maximum of this and with.
func Maxi[X Int.Any](a XYZW, b X) XYZW { //gd:Vector4i.maxi
	return XYZW{max(a.X, int32(b)), max(a.Y, int32(b)), max(a.Z, int32(b)), max(a.W, int32(b))}
}

// Min returns the component-wise minimum of this and with.
func Min(a, b XYZW) XYZW { //gd:Vector4i.min
	return XYZW{min(a.X, b.X), min(a.Y, b.Y), min(a.Z, b.Z), min(a.W, b.W)}
}

// Mini returns the component-wise minimum of this and with.
func Mini[X Int.Any](a XYZW, b X) XYZW { //gd:Vector4i.mini
	return XYZW{min(a.X, int32(b)), min(a.Y, int32(b)), min(a.Z, int32(b)), min(a.W, int32(b))}
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all components are
// equal, this method returns [X].
func MaxAxis(v XYZW) Axis { //gd:Vector4i.max_axis_index
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
func MinAxis(v XYZW) Axis { //gd:Vector4i.min_axis_index
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

// Sign returns a new vector with each component set to 1 if it's positive, -1 if it's
// negative, and 0 if it's zero. The result is identical to calling [Signi] on each component.
func Sign(v XYZW) XYZW { return XYZW{Int.Sign(v.X), Int.Sign(v.Y), Int.Sign(v.Z), Int.Sign(v.W)} } //gd:Vector4i.sign

// Snapped returns a new vector with each component snapped to the closest multiple of the corresponding
// component in step.
func Snapped(v, step XYZW) XYZW { //gd:Vector4i.snapped
	return XYZW{
		Int.Snapped(v.X, step.X),
		Int.Snapped(v.Y, step.Y),
		Int.Snapped(v.Z, step.Z),
		Int.Snapped(v.W, step.W),
	}
}

// Snappedi returns a new vector with each component snapped to the closest multiple of the corresponding
// component in step.
func Snappedi[X Int.Any](v XYZW, step X) XYZW { //gd:Vector4i.snappedi
	return XYZW{
		Int.Snapped(v.X, int32(step)),
		Int.Snapped(v.Y, int32(step)),
		Int.Snapped(v.Z, int32(step)),
		Int.Snapped(v.W, int32(step)),
	}
}

func Add(a, b XYZW) XYZW { return XYZW{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W} } //gd:Vector4i+(right:Vector4i)
func Sub(a, b XYZW) XYZW { return XYZW{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W} } //gd:Vector4i-(right:Vector4i)
func Mul(a, b XYZW) XYZW { return XYZW{a.X * b.X, a.Y * b.Y, a.Z * b.Z, a.W * b.W} } //gd:Vector4i*(right:Vector4i)
func Div(a, b XYZW) XYZW { return XYZW{a.X / b.X, a.Y / b.Y, a.Z / b.Z, a.W / b.W} } //gd:Vector4i/(right:Vector4i)
func Mod(a, b XYZW) XYZW { return XYZW{a.X % b.X, a.Y % b.Y, a.Z % b.Z, a.W % b.W} } //gd:Vector4i%(right:Vector4i)

func AddX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4i+(right:float)
	return XYZW{a.X + int32(b), a.Y + int32(b), a.Z + int32(b), a.W + int32(b)}
}
func SubX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4i-(right:float)
	return XYZW{a.X - int32(b), a.Y - int32(b), a.Z - int32(b), a.W - int32(b)}
}
func MulX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4i*(right:float)
	return XYZW{a.X * int32(b), a.Y * int32(b), a.Z * int32(b), a.W * int32(b)}
}
func DivX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4i/(right:float)
	return XYZW{a.X / int32(b), a.Y / int32(b), a.Z / int32(b), a.W / int32(b)}
}
func ModX[X Int.Any | Float.Any](a XYZW, b X) XYZW { //gd:Vector4i%(right:int)
	return XYZW{a.X % int32(b), a.Y % int32(b), a.Z % int32(b), a.W - int32(b)}
}

func Neg(v XYZW) XYZW { return XYZW{-v.X, -v.Y, -v.Z, -v.W} } //gd:Vector4i-(unary)

func AsArray(vec XYZW) [4]int32 { return [4]int32{vec.X, vec.Y, vec.Z, vec.W} }

func Index[I Int.Any](v XYZW, i I) int { //gd:Vector4i[](index:int)
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
