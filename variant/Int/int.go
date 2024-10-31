// Package Int provides generic methods for working with integers.
package Int

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
