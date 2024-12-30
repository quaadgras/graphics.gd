// Package Int provides generic methods for working with integers.
package Int

import (
	"math/rand"
	"unsafe"

	"graphics.gd/variant/Float"
)

// Any integer.
type Any interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

// Absi returns the absolute value of the integer parameter x (i.e. non-negative value).
func Abs[X Any](v X) X { //gd:absi
	if v < 0 {
		return -v
	}
	return v
}

// Clamp clamps the value, returning an integer not less than min and not more than max.
func Clamp[X Any](value, min, max X) X { //gd:clampi
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Floor rounds x downward (towards negative infinity), returning the largest whole number
// that is not more than x.
func Floor[X Float.Any](x X) int { //gd:floori floor
	return int(Float.Floor(x))
}

// Ceil rounds x upward (towards positive infinity), returning the smallest whole number
// that is not less than x.
func Ceil[X Float.Any](x X) int { //gd:ceili ceil
	return int(Float.Ceil(x))
}

// Round rounds x to the nearest whole number, rounding halfway cases away from zero.
func Round[X Float.Any](x X) int { //gd:roundi round
	return int(Float.Round(x))
}

// Sign returns -1 if x is negative, 1 if x is positive, and 0 if x is zero. For NaN values
// of x it returns 0.
func Sign[X Any](x X) X { //gd:signi
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}

// Posmod returns the integer modulus of x divided by y that wraps equally in positive and negative.
func Posmod[T Any](x, y T) T { //gd:posmod
	value := x % y
	if ((value < 0) && (y > 0)) || ((value > 0) && (y < 0)) {
		value += y
	}
	return value
}

// Snapped returns the multiple of step that is the closest to x. This can also be used to round a
// floating point number to an arbitrary number of decimals.
func Snapped[X Any](val, by X) X { //gd:snappedi
	if by != 0 {
		val = (val / by) * by
	}
	return val
}

// Wrap value between min and max. Can be used for creating loop-alike behavior or
// infinite surfaces.
func Wrap[T ~int8 | ~int16 | ~int32 | ~int64 | ~int](value, min, max T) T { //gd:wrapi
	var diff = max - min
	if diff == 0 {
		return min
	}
	return min + ((((value - min) % diff) + diff) % diff)
}

// Max returns the largest of the two integers.
func Max[T Any](a, b T) T { //gd:maxi
	return max(a, b)
}

// Min returns the smallest of the two values.
func Min[T Any](a, b T) T { //gd:mini
	return min(a, b)
}

// NearestPowerOfTwo returns the smallest integer power of 2 that is greater than
// or equal to value.
func NearestPowerOfTwo[T Any](x T) T { //gd:nearest_po2
	get_shift_from_power_of_2 := func(bits uintptr) int {
		for i := 0; i < int(unsafe.Sizeof(T(0))*8); i++ {
			if bits == (1 << i) {
				return i
			}
		}
		return -1
	}
	x--
	num := get_shift_from_power_of_2(unsafe.Sizeof(x)) + 3
	for i := 0; i < num; i++ {
		x |= x >> (1 << i)
	}
	return x + 1
}

// Random returns a random integer between min and max (inclusive).
func Random() int { //gd:randi
	return rand.Int()
}

// RandomBetween returns a random integer between min and max (inclusive).
func RandomBetween(min, max int) int { //gd:randi_range
	return rand.Intn(max-min+1) + min
}
