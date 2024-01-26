package spatial

import "math"

// Vector3i is A 3-element structure that can be used to represent 3D grid coordinates or any other triplet of integers.
//
// It uses integer coordinates and is therefore preferable to [Vector3] when exact precision is required. Note that the
// values are limited to 32 bits, and unlike [Vector3] this cannot be configured with an engine build option. Use [Int]
// or [PackedInt64Array] if 64-bit values are needed.
//
// Note: In a boolean context, a [Vector3i] will evaluate to false if it's equal to Vector3i{0, 0, 0}. Otherwise, a
// [Vector3i] will always evaluate to true.
type Vector3i [3]int32

func (v Vector3i) X() Int      { return Int(v[X]) }
func (v Vector3i) Y() Int      { return Int(v[Y]) }
func (v Vector3i) Z() Int      { return Int(v[Z]) }
func (v *Vector3i) SetX(x Int) { v[X] = int32(x) }
func (v *Vector3i) SetY(y Int) { v[Y] = int32(y) }
func (v *Vector3i) SetZ(z Int) { v[Z] = int32(z) }

// "Constants"
func (Vector3i) ZERO() Vector3i    { return Vector3i{} }
func (Vector3i) ONE() Vector3i     { return Vector3i{1, 1, 1} }
func (Vector3i) MIN() Vector3i     { return Vector3i{math.MinInt32, math.MinInt32, math.MinInt32} }
func (Vector3i) MAX() Vector3i     { return Vector3i{math.MaxInt32, math.MaxInt32, math.MaxInt32} }
func (Vector3i) LEFT() Vector3i    { return Vector3i{-1, 0, 0} }
func (Vector3i) RIGHT() Vector3i   { return Vector3i{1, 0, 0} }
func (Vector3i) UP() Vector3i      { return Vector3i{0, 1, 0} }
func (Vector3i) DOWN() Vector3i    { return Vector3i{0, -1, 0} }
func (Vector3i) FORWARD() Vector3i { return Vector3i{0, 0, 1} }
func (Vector3i) BACK() Vector3i    { return Vector3i{0, 0, -1} }

func (v Vector3i) Vector3() Vector3 { return Vector3{float(v[X]), float(v[Y]), float(v[Z])} }

// Abs returns a new vector with all components in absolute values (i.e. positive).
func (v Vector3i) Abs() Vector3i { return v.abs() }

// Clamp returns a new vector with all components clamped between the components of min and max,
// by running [Clampi] on each component.
func (v Vector3i) Clamp(min, max Vector3i) Vector3i {
	return Vector3i{
		Clampi(v[X], min[X], max[X]),
		Clampi(v[Y], min[Y], max[Y]),
		Clampi(v[Z], min[Z], max[Z]),
	}
}

// Length returns the length (magnitude) of this vector.
func (v Vector3i) Length() Float { return Sqrt(Float(v.LengthSquared())) }

// LengthSquared returns the squared length (squared magnitude) of this vector.
//
// This method runs faster than length, so prefer it if you need to compare vectors
// or need the squared distance for some formula.
func (v Vector3i) LengthSquared() Int {
	return Int(v[X])*Int(v[X]) + Int(v[Y])*Int(v[Y]) + Int(v[Z])*Int(v[Z])
}

// MaxAxis returns the axis of the vector's highest value. See [Axis] constants. If all components are
// equal, this method returns [X].
func (v Vector3i) MaxAxis() Axis {
	if v[X] < v[Y] {
		if v[Y] < v[Z] {
			return Z
		}
		return Y
	}
	if v[X] < v[Z] {
		return Z
	}
	return X
}

// MinAxis returns the axis of the vector's lowest value. See [Axis] constants. If all components are
// equal, this method returns [Z].
func (v Vector3i) MinAxis() Axis {
	if v[X] < v[Y] {
		if v[Y] < v[Z] {
			return X
		}
		return Z
	}
	if v[X] < v[Z] {
		return Y
	}
	return Z
}

// Sign returns a new vector with each component set to 1 if it's positive, -1 if it's negative, and 0 if
// it's zero. The result is identical to calling [Signi] on each component.
func (v Vector3i) Sign() Vector3i { return v.sign() }

// Snapped returns a new vector with each component snapped to the closest multiple of the corresponding
// component in step.
func (v Vector3i) Snapped(step Vector3i) Vector3i {
	return Vector3i{
		Snappedi(v[X], step[X]),
		Snappedi(v[Y], step[Y]),
		Snappedi(v[Z], step[Z]),
	}
}

func (v Vector3i) Add(other Vector3i) Vector3i {
	return Vector3i{v[x] + other[x], v[y] + other[y], v[z] + other[z]}
}
func (v Vector3i) Sub(other Vector3i) Vector3i {
	return Vector3i{v[x] - other[x], v[y] - other[y], v[z] - other[z]}
}
func (v Vector3i) Mul(other Vector3i) Vector3i {
	return Vector3i{v[x] * other[x], v[y] * other[y], v[z] * other[z]}
}
func (v Vector3i) Div(other Vector3i) Vector3i {
	return Vector3i{v[x] / other[x], v[y] / other[y], v[z] / other[z]}
}
func (v Vector3i) Mod(other Vector3i) Vector3i {
	return Vector3i{v[x] % other[x], v[y] % other[y], v[z] % other[z]}
}
func (v Vector3i) Addi(other Int) Vector3i {
	return Vector3i{v[x] + int32(other), v[y] + int32(other), v[z] + int32(other)}
}
func (v Vector3i) Subi(other Int) Vector3i {
	return Vector3i{v[x] - int32(other), v[y] - int32(other), v[z] - int32(other)}
}
func (v Vector3i) Muli(other Int) Vector3i {
	return Vector3i{v[x] * int32(other), v[y] * int32(other), v[z] * int32(other)}
}
func (v Vector3i) Divi(other Int) Vector3i {
	return Vector3i{v[x] / int32(other), v[y] / int32(other), v[z] / int32(other)}
}
func (v Vector3i) Modi(other Int) Vector3i {
	return Vector3i{v[x] % int32(other), v[y] % int32(other), v[z] % int32(other)}
}
func (v Vector3i) Neg() Vector3i { return Vector3i{-v[x], -v[y], -v[z]} }
