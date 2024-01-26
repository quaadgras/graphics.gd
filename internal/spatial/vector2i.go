package spatial

import "math"

/*
Vector2i is a 2-element structure that can be used to represent 2D grid coordinates
or any other pair of integers.

It uses integer coordinates and is therefore preferable to Vector2 when exact precision
is required. Note that the values are limited to 32 bits, and unlike Vector2 this cannot
be configured with an engine build option. Use [Int] or [PackedInt64Array] if 64-bit values
are needed.

Note: In a boolean context, a [Vector2i] will evaluate to false if it's equal to Vector2i{0, 0}.
Otherwise, a [Vector2i] will always evaluate to true.
*/
type Vector2i [2]int32

func (v Vector2i) X() Int     { return Int(v[0]) }
func (v Vector2i) Y() Int     { return Int(v[1]) }
func (v Vector2i) SetX(x Int) { v[0] = int32(x) }
func (v Vector2i) SetY(y Int) { v[1] = int32(y) }

// "Constants"
func (v Vector2i) ZERO() Vector2i  { return Vector2i{} }
func (v Vector2i) ONE() Vector2i   { return Vector2i{1, 1} }
func (v Vector2i) MIN() Vector2i   { return Vector2i{math.MinInt32, math.MinInt32} }
func (v Vector2i) MAX() Vector2i   { return Vector2i{math.MaxInt32, math.MaxInt32} }
func (v Vector2i) LEFT() Vector2i  { return Vector2i{-1, 0} }
func (v Vector2i) RIGHT() Vector2i { return Vector2i{1, 0} }
func (v Vector2i) UP() Vector2i    { return Vector2i{0, -1} }
func (v Vector2i) DOWN() Vector2i  { return Vector2i{0, 1} }

// Vector2 constructs a new Vector2 from Vector2i.
func (v Vector2i) Vector2() Vector2 { return Vector2{float(v[0]), float(v[1])} }

// Abs returns a new vector with all components in absolute values (i.e. positive).
func (v Vector2i) Abs() Vector2i { return v.abs() }

// Aspect returns the aspect ratio of this vector, the ratio of x to y.
func (v Vector2i) Aspect() Float { return Float(v[0]) / Float(v[1]) }

// Clamp returns a new vector with all components clamped between the components of min
// and max, by running clamp on each component.
func (v Vector2i) Clamp(min, max Vector2i) Vector2i {
	return Vector2i{
		Clampi(v[0], min[0], max[0]),
		Clampi(v[1], min[1], max[1]),
	}
}

// Length returns the length (magnitude) of this vector.
func (v Vector2i) Length() Float { return Sqrt(Float(v.LengthSquared())) }

// LengthSquared returns the squared length (squared magnitude) of this vector.
//
// This method runs faster than length, so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func (v Vector2i) LengthSquared() Int { return Int(v[x]*v[x] + v[y]*v[y]) }

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If
// all components are equal, this method returns [X].
func (v Vector2i) MaxAxis() Axis {
	if v[Y] > v[X] {
		return Y
	}
	return X
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If
// all components are equal, this method returns [Y].
func (v Vector2i) MinAxis() Axis {
	if v[Y] < v[X] {
		return Y
	}
	return X
}

// Sign returns a new vector with each component set to 1 if it's positive, -1 if it's
// negative, and 0 if it's zero. The result is identical to calling [Signi] on each component.
func (v Vector2i) Sign() Vector2i { return v.sign() }

// Snapped returns a new vector with each component snapped to the closest multiple of the corresponding
// component in step.
func (v Vector2i) Snapped(step Vector2i) Vector2i {
	return Vector2i{
		Snappedi(v[0], step[0]),
		Snappedi(v[1], step[1]),
	}
}

func (v Vector2i) Add(other Vector2i) Vector2i { return Vector2i{v[x] + other[x], v[y] + other[y]} }
func (v Vector2i) Sub(other Vector2i) Vector2i { return Vector2i{v[x] - other[x], v[y] - other[y]} }
func (v Vector2i) Mul(other Vector2i) Vector2i { return Vector2i{v[x] * other[x], v[y] * other[y]} }
func (v Vector2i) Div(other Vector2i) Vector2i { return Vector2i{v[x] / other[x], v[y] / other[y]} }
func (v Vector2i) Mod(other Vector2i) Vector2i { return Vector2i{v[x] % other[x], v[y] % other[y]} }
func (v Vector2i) Addi(other Int) Vector2i     { return Vector2i{v[x] + int32(other), v[y] + int32(other)} }
func (v Vector2i) Subi(other Int) Vector2i     { return Vector2i{v[x] - int32(other), v[y] - int32(other)} }
func (v Vector2i) Muli(other Int) Vector2i     { return Vector2i{v[x] * int32(other), v[y] * int32(other)} }
func (v Vector2i) Divi(other Int) Vector2i     { return Vector2i{v[x] / int32(other), v[y] / int32(other)} }
func (v Vector2i) Modi(other Int) Vector2i     { return Vector2i{v[x] % int32(other), v[y] % int32(other)} }
func (v Vector2i) Neg() Vector2i               { return Vector2i{-v[x], -v[y]} }
