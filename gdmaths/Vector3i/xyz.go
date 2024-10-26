// Package Vector3 provides a 3D vector using integer coordinates.
package Vector3i

import (
	"math"

	"grow.graphics/gd/gdmaths/Float"
	"grow.graphics/gd/gdmaths/Int"
)

// XYZ is a 3-element structure that can be used to represent 3D grid
// coordinates or any other triplet of integers.
//
// It uses integer coordinates and is therefore preferable to Vector3
// when exact precision is required. Note that the values are limited
// to 32 bits, and unlike Vector3 this cannot be configured with a
// build tag. Use int or PackedInt64Array if 64-bit values are needed.
type XYZ = struct {
	X int32 // The vector's X component.
	Y int32 // The vector's Y component.
	Z int32 // The vector's Z component.
}

type Axis int

const (
	X Axis = iota // Enumerated value for the X axis. Returned by [MaxAxis] and [MinAxis].
	Y             // Enumerated value for the Y axis. Returned by [MaxAxis] and [MinAxis].
	Z             // Enumerated value for the Z axis. Returned by [MaxAxis] and [MinAxis].
)

var (
	Zero   = XYZ{0, 0, 0}                                     // Zero vector, a vector with all components set to 0.
	One    = XYZ{1, 1, 1}                                     // One vector, a vector with all components set to 1.
	MinXYZ = XYZ{math.MinInt32, math.MinInt32, math.MinInt32} // Min vector, a vector with all components equal to [math.MinInt32]. Can be used as a negative integer equivalent of Vector3.INF.
	MaxXYZ = XYZ{math.MaxInt32, math.MaxInt32, math.MaxInt32} // Max vector, a vector with all components equal to [math.MaxInt32]. Can be used as a positive integer equivalent of Vector3.INF.
	Left   = XYZ{-1, 0, 0}                                    // Left unit vector. Represents the local direction of left, and the global direction of west.
	Right  = XYZ{1, 0, 0}                                     // Right unit vector. Represents the local direction of right, and the global direction of east.
	Up     = XYZ{0, -1, 0}                                    // Up unit vector.
	Down   = XYZ{0, 1, 0}                                     // Down unit vector.

	// Forward unit vector. Represents the local direction of forward, and the global direction
	// of north. Keep in mind that the forward direction for lights, cameras, etc is different
	// from 3D assets like characters, which face towards the camera by convention. Use [ModelFront]
	// and similar constants when working in 3D asset space.
	Forward = XYZ{0, 0, -1}

	Back = XYZ{0, 0, 1} // Back unit vector. Represents the local direction of back, and the global direction of south.
)

// New returns a [XYZ] with the given components.
func New[X Int.Any](x, y, z X) XYZ { //gd:Vector3i(x: int, y: int, z: int)
	return XYZ{int32(x), int32(y), int32(z)}
}

// Abs returns a new vector with all components in absolute values (i.e. positive).
func Abs(v XYZ) XYZ { //gd:Vector3i.abs
	return XYZ{
		Int.Abs(v.X), Int.Abs(v.Y), Int.Abs(v.Z),
	}
}

// Clamp returns a new vector with all components clamped between the components of min and max,
// by running [Int.Clamp] on each component.
func Clamp(v, min, max XYZ) XYZ { //gd:Vector3i.clamp
	return XYZ{
		Int.Clamp(v.X, min.X, max.X),
		Int.Clamp(v.Y, min.Y, max.Y),
		Int.Clamp(v.Z, min.Z, max.Z),
	}
}

// Clampi returns a new vector with all components clamped between the components of min and max,
// by running [Int.Clamp] on each component.
func Clampi[X Int.Any](v XYZ, min, max X) XYZ { //gd:Vector3i.clampi
	return XYZ{
		Int.Clamp(v.X, int32(min), int32(max)),
		Int.Clamp(v.Y, int32(min), int32(max)),
		Int.Clamp(v.Z, int32(min), int32(max)),
	}
}

// DistanceSquared returns the squared distance between this vector and to.
//
// This method runs faster than [Vector3.DistanceTo], so prefer it if you need to compare vectors or need the squared distance for
// some formula.
func DistanceSquared(v, to XYZ) int { //gd:Vector3i.distance_squared_to
	return LengthSquared(Sub(v, to))
}

// Distance returns the distance between this vector and to.
func Distance(v, to XYZ) Float.X { //gd:Vector3i.distance_to
	return Length(Sub(v, to))
}

// Length the length (magnitude) of this vector.
func Length(v XYZ) Float.X { //gd:Vector3.length
	return Float.X(math.Sqrt(float64(LengthSquared(v))))
}

// LengthSquared returns the squared length (squared magnitude) of this vector.
//
// This method runs faster than length, so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func LengthSquared(v XYZ) int { //gd:Vector3.length_squared
	return int(v.X)*int(v.X) + int(v.Y)*int(v.Y) + int(v.Z)*int(v.Z)
}

// Max returns the component-wise maximum of this and with.
func Max(a, b XYZ) XYZ { //gd:Vector3i.max
	return XYZ{max(a.X, b.X), max(a.Y, b.Y), max(a.Z, b.Z)}
}

// Max returns the component-wise maximum of this and with.
func Maxi[X Int.Any](a XYZ, b X) XYZ { //gd:Vector3i.maxi
	return XYZ{max(a.X, int32(b)), max(a.Y, int32(b)), max(a.Z, int32(b))}
}

// Min returns the component-wise minimum of this and with.
func Min(a, b XYZ) XYZ { //gd:Vector3i.min
	return XYZ{min(a.X, b.X), min(a.Y, b.Y), min(a.Z, b.Z)}
}

// Mini returns the component-wise minimum of this and with.
func Mini[X Int.Any](a XYZ, b X) XYZ { //gd:Vector3i.mini
	return XYZ{min(a.X, int32(b)), min(a.Y, int32(b)), min(a.Z, int32(b))}
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all components are
// equal, this method returns [X].
func MaxAxis(v XYZ) Axis { //gd:Vector3i.max_axis_index
	if v.X < v.Y {
		if v.Y < v.Z {
			return Z
		}
		return Y
	}
	if v.X < v.Z {
		return Z
	}
	return X
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all components are
// equal, this method returns [Z].
func MinAxis(v XYZ) Axis { //gd:Vector3i.min_axis_index
	if v.X < v.Y {
		if v.Y < v.Z {
			return X
		}
		return Z
	}
	if v.X < v.Z {
		return Y
	}
	return Z
}

// Sign returns a new vector with each component set to 1 if it's positive, -1 if it's
// negative, and 0 if it's zero. The result is identical to calling [Signi] on each component.
func Sign(v XYZ) XYZ { return XYZ{Int.Sign(v.X), Int.Sign(v.Y), Int.Sign(v.Z)} } //gd:Vector3i.sign

// Snapped returns a new vector with each component snapped to the closest multiple of the corresponding
// component in step.
func Snapped(v, step XYZ) XYZ { //gd:Vector3i.snapped
	return XYZ{
		Int.Snapped(v.X, step.X),
		Int.Snapped(v.Y, step.Y),
		Int.Snapped(v.Z, step.Z),
	}
}

// Snappedi returns a new vector with each component snapped to the closest multiple of the corresponding
// component in step.
func Snappedi[X Int.Any](v XYZ, step X) XYZ { //gd:Vector3i.snappedi
	return XYZ{
		Int.Snapped(v.X, int32(step)),
		Int.Snapped(v.Y, int32(step)),
		Int.Snapped(v.Z, int32(step)),
	}
}

func Add(a, b XYZ) XYZ { return XYZ{a.X + b.X, a.Y + b.Y, a.Z + b.Z} } //gd:Vector3i+(right:Vector2i)
func Sub(a, b XYZ) XYZ { return XYZ{a.X - b.X, a.Y - b.Y, a.Z - b.Z} } //gd:Vector3i-(right:Vector2i)
func Mul(a, b XYZ) XYZ { return XYZ{a.X * b.X, a.Y * b.Y, a.Z * b.Z} } //gd:Vector3i*(right:Vector2i)
func Div(a, b XYZ) XYZ { return XYZ{a.X / b.X, a.Y / b.Y, a.Z / b.Z} } //gd:Vector3i/(right:Vector2i)
func Mod(a, b XYZ) XYZ { return XYZ{a.X % b.X, a.Y % b.Y, a.Z % b.Z} } //gd:Vector3i%(right:Vector2i)

func Addi[X Int.Any](a XYZ, b X) XYZ { return XYZ{a.X + int32(b), a.Y + int32(b), a.Z + int32(b)} } //gd:Vector3i+(right:int)
func Subi[X Int.Any](a XYZ, b X) XYZ { return XYZ{a.X - int32(b), a.Y - int32(b), a.Z - int32(b)} } //gd:Vector3i-(right:int)
func Muli[X Int.Any](a XYZ, b X) XYZ { return XYZ{a.X * int32(b), a.Y * int32(b), a.Z * int32(b)} } //gd:Vector3i*(right:int)
func Divi[X Int.Any](a XYZ, b X) XYZ { return XYZ{a.X / int32(b), a.Y / int32(b), a.Z / int32(b)} } //gd:Vector3i/(right:int)
func Modi[X Int.Any](a XYZ, b X) XYZ { return XYZ{a.X % int32(b), a.Y % int32(b), a.Z % int32(b)} } //gd:Vector3i%(right:int)

func Addf[X Float.Any](a XYZ, b X) XYZ { return XYZ{a.X + int32(b), a.Y + int32(b), a.Z + int32(b)} } //gd:Vector3i+(right:float)
func Subf[X Float.Any](a XYZ, b X) XYZ { return XYZ{a.X - int32(b), a.Y - int32(b), a.Z - int32(b)} } //gd:Vector3i-(right:float)
func Mulf[X Float.Any](a XYZ, b X) XYZ { return XYZ{a.X * int32(b), a.Y * int32(b), a.Z * int32(b)} } //gd:Vector3i*(right:float)
func Divf[X Float.Any](a XYZ, b X) XYZ { return XYZ{a.X / int32(b), a.Y / int32(b), a.Z / int32(b)} } //gd:Vector3i/(right:float)

func Neg(v XYZ) XYZ { return XYZ{-v.X, -v.Y, -v.Z} } //gd:Vector3i-(unary)

func Index[I Int.Any](v XYZ, i I) int { //gd:Vector3i[](index:int)
	switch Axis(i) {
	case X:
		return int(v.X)
	case Y:
		return int(v.Y)
	case Z:
		return int(v.Z)
	default:
		panic("index out of range")
	}
}
