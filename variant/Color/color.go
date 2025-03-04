// Package Color provides a color represented in RGBA format.
package Color

import (
	"math"

	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
)

// A color represented in RGBA format by a red (r), green (g), blue (b), and
// alpha (a) component. Each component is a 32-bit floating-point value,
// usually ranging from 0.0 to 1.0. In some cases, it may support values
// greater than 1.0, for overbright or HDR (High Dynamic Range) colors.
type RGBA = struct {
	R Float.X
	G Float.X
	B Float.X
	A Float.X
}

// Less compares two colors by first checking if the R value of the left color is less than the R value of the right color.
// If the R values are exactly equal, then it repeats this check with the G values of the two colors, B values of the two
// colors, and then with the A values. This operator is useful for sorting colors.
//
// Note: Colors with NaN elements don't behave the same as other colors. Therefore, the results from this operator may not
// be accurate if NaNs are included.
func Less(a, b RGBA) bool { //gd:Color<Color
	return a.R < b.R && a.G < b.G && a.B < b.B && a.A < b.A
}

// Bytes returns a color represented in RGBA format by a red (r), green
// (g), blue (b), and alpha (a) bytes from 0 to 255.
func Bytes(r, g, b, a byte) RGBA { //gd:Color.from_rgba8
	return RGBA{
		R: Float.X(r) / 255,
		G: Float.X(g) / 255,
		B: Float.X(b) / 255,
		A: Float.X(a) / 255,
	}
}

// ʕ is a little ternary operator for porting C code.
func ʕ[T any](q bool, a T, b T) T {
	if q {
		return a
	}
	return b
}

// Blend returns a new color resulting from overlaying this color over the given color.
// In a painting program, you can imagine it as the over color painted over this
// color (including alpha).
func Blend(a, b RGBA) RGBA { //gd:Color.blend
	var res RGBA
	var sa = 1.0 - b.A
	res.A = a.A*sa + b.A
	if res.A == 0 {
		return RGBA{}
	} else {
		res.R = (a.R*a.A*sa + b.R*b.A) / res.A
		res.G = (a.G*a.A*sa + b.G*b.A) / res.A
		res.B = (a.B*a.A*sa + b.B*b.A) / res.A
	}
	return res
}

// Clamp returns a new color with all components clamped between the components of minimum and maximum.
func Clamp(c RGBA, minimum, maximum RGBA) RGBA { //gd:Color.clamp
	return RGBA{
		Float.Clamp(c.R, minimum.R, maximum.R),
		Float.Clamp(c.G, minimum.G, maximum.G),
		Float.Clamp(c.B, minimum.B, maximum.B),
		Float.Clamp(c.A, minimum.A, maximum.A),
	}
}

// Darkened returns a new color resulting from making this color darker by the specified amount
// (ratio from 0.0 to 1.0). See also [Lightened].
func Darkened(c RGBA, amount Float.X) RGBA { //gd:Color.darkened
	return RGBA{
		c.R * (1.0 - amount),
		c.G * (1.0 - amount),
		c.B * (1.0 - amount),
		c.A,
	}
}

// Luminance returns the light intensity of the color, as a value between 0.0 and 1.0 (inclusive).
// This is useful when determining light or dark color. Colors with a luminance smaller than 0.5
// can be generally considered dark.
//
// Note: [Luminance] relies on the color being in the linear color space to return an accurate
// relative luminance value. If the color is in the sRGB color space, use srgb_to_linear to convert
// it to the linear color space first.
func Luminance(c RGBA) Float.X { //gd:Color.get_luminance
	return 0.2126*c.R + 0.7152*c.G + 0.0722*c.B
}

// Inverted returns the color with its r, g, and b components inverted ((1 - r, 1 - g, 1 - b, a)).
func Inverted(c RGBA) RGBA { return RGBA{1 - c.R, 1 - c.G, 1 - c.B, c.A} } //gd:Color.inverted

// IsApproximatelyEqual returns true if this color and to are approximately equal.
func IsApproximatelyEqual(a, b RGBA) bool { //gd:Color.is_equal_approx
	return Float.IsApproximatelyEqual(a.R, b.R) &&
		Float.IsApproximatelyEqual(a.G, b.G) &&
		Float.IsApproximatelyEqual(a.B, b.B) &&
		Float.IsApproximatelyEqual(a.A, b.A)
}

// Lerp returns the linear interpolation between this color's components and to's components.
// The interpolation factor weight should be between 0.0 and 1.0 (inclusive).
func Lerp(a, b RGBA, weight Float.X) RGBA { //gd:Color.lerp
	return RGBA{
		Float.Lerp(a.R, b.R, weight),
		Float.Lerp(a.G, b.G, weight),
		Float.Lerp(a.B, b.B, weight),
		Float.Lerp(a.A, b.A, weight),
	}
}

// Lightened returns a new color resulting from making this color lighter by the specified amount
// which should be a ratio from 0.0 to 1.0. See also [Darkened].
func Lightened(c RGBA, amount Float.X) RGBA { //gd:Color.lightened
	return RGBA{
		c.R + (1.0-c.R)*amount,
		c.G + (1.0-c.G)*amount,
		c.B + (1.0-c.B)*amount,
		c.A,
	}
}

// ToSRGB returns the color converted to the sRGB color space. This method assumes the original color is
// in the linear color space. See also [Color.Linear] which performs the opposite operation.
func ToSRGB(c RGBA) RGBA { //gd:Color.linear_to_srgb
	return RGBA{
		ʕ(c.R < 0.0031308, 12.92*c.R, (1.0+0.055)*Float.Pow(c.R, 1.0/2.4)-0.055),
		ʕ(c.G < 0.0031308, 12.92*c.G, (1.0+0.055)*Float.Pow(c.G, 1.0/2.4)-0.055),
		ʕ(c.B < 0.0031308, 12.92*c.B, (1.0+0.055)*Float.Pow(c.B, 1.0/2.4)-0.055),
		c.A,
	}
}

// ToLinear returns the color converted to the linear color space. This method assumes the original color
// already is in the sRGB color space. See also [SRGB] which performs the opposite operation.
func ToLinear(c RGBA) RGBA { //gd:Color.srgb_to_linear
	return RGBA{
		ʕ(c.R <= 0.04045, c.R/12.92, Float.Pow((c.R+0.055)/1.055, 2.4)),
		ʕ(c.G <= 0.04045, c.G/12.92, Float.Pow((c.G+0.055)/1.055, 2.4)),
		ʕ(c.B <= 0.04045, c.B/12.92, Float.Pow((c.B+0.055)/1.055, 2.4)),
		c.A,
	}
}

// AsABGR32 returns the color converted to a 32-bit integer in ABGR format (each component is 8 bits).
// ABGR is the reversed version of the default RGBA format.
func AsABGR32(c RGBA) uint32 { //gd:Color.to_abgr32
	var u32 = uint32(Float.Round(c.A * 255))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.B * 255.0))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.G * 255.0))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.R * 255.0))
	return u32
}

// AsABGR64 returns the color converted to a 64-bit integer in ABGR format (each component is 16 bits).
// ABGR is the reversed version of the default RGBA format.
func AsABGR64(c RGBA) uint64 { //gd:Color.to_abgr64
	var u64 = uint64(Float.Round(c.A * 65535))
	u64 <<= 16
	u64 |= uint64(Float.Round(float64(c.B) * 65535))
	u64 <<= 16
	u64 |= uint64(Float.Round(float64(c.G) * 65535))
	u64 <<= 16
	u64 |= uint64(Float.Round(float64(c.R) * 65535))
	return u64
}

// AsARGB32 returns the color converted to a 32-bit integer in ARGB format (each component is 8 bits).
// ARGB is more compatible with DirectX.
func AsARGB32(c RGBA) uint32 { //gd:Color.to_argb32
	var u32 = uint32(Float.Round(c.A * 255))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.R * 255.0))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.G * 255.0))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.B * 255.0))
	return u32
}

// AsARGB64 returns the color converted to a 64-bit integer in ARGB format (each component is 16 bits).
// ARGB is more compatible with DirectX.
func AsARGB64(c RGBA) uint64 { //gd:Color.to_argb64
	var u64 = uint64(Float.Round(float64(c.A) * 65535))
	u64 <<= 16
	u64 |= uint64(Float.Round(float64(c.R) * 65535))
	u64 <<= 16
	u64 |= uint64(Float.Round(float64(c.G) * 65535))
	u64 <<= 16
	u64 |= uint64(Float.Round(float64(c.B) * 65535))
	return u64
}

// AsHex returns the color converted to an HTML hexadecimal color String in RGBA format
// without the hash (#) prefix.
func AsHex(c RGBA) string { //gd:Color.to_html
	var txt string
	txt += _to_hex(c.R)
	txt += _to_hex(c.G)
	txt += _to_hex(c.B)
	const with_alpha = true
	if with_alpha {
		txt += _to_hex(c.A)
	}
	return txt
}

func _to_hex(val float32) string {
	v := rune(min(255, max(0, math.Round(float64(val*255)))))
	var ret string
	for i := 0; i < 2; i++ {
		var c = [2]rune{0, 0}
		var lv = v & 0xF
		if lv < 10 {
			c[0] = '0' + lv
		} else {
			c[0] = 'a' + lv - 10
		}
		v >>= 4
		var cs = string(c[:])
		ret = cs + ret
	}
	return ret
}

// AsRGBA32 returns the color converted to a 32-bit integer in RGBA format
// (each component is 8 bits).
func AsRGBA32(c RGBA) uint32 { // Color.to_rgba32
	var u32 = uint32(Float.Round(c.R * 255))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.G * 255.0))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.B * 255.0))
	u32 <<= 8
	u32 |= uint32(Float.Round(c.A * 255.0))
	return u32
}

// AsRGBA64 returns the color converted to a 64-bit integer in RGBA format
// (each component is 16 bits).
func AsRGBA64(c RGBA) int64 { //gd:Color.to_rgba64
	var u64 = int64(Float.Round(c.R * 65535))
	u64 <<= 16
	u64 |= int64(Float.Round(c.G * 65535))
	u64 <<= 16
	u64 |= int64(Float.Round(c.B * 65535))
	u64 <<= 16
	u64 |= int64(Float.Round(c.A * 65535))
	return u64
}

func Add(a, b RGBA) RGBA { //gd:Color+Color
	return RGBA{a.R + b.R, a.G + b.G, a.B + b.B, a.A + b.A}
}
func Sub(a, b RGBA) RGBA { //gd:Color-Color
	return RGBA{a.R - b.R, a.G - b.G, a.B - b.B, a.A - b.A}
}
func Mul(a, b RGBA) RGBA { //gd:Color*Color
	return RGBA{a.R * b.R, a.G * b.G, a.B * b.B, a.A * b.A}
}
func Div(a, b RGBA) RGBA { //gd:Color/Color
	return RGBA{a.R / b.R, a.G / b.G, a.B / b.B, a.A / b.A}
}
func Neg(a RGBA) RGBA { //gd:-Color
	return RGBA{-a.R, -a.G, -a.B, -a.A}
}
func AddX[X Float.Any | Int.Any](a RGBA, b X) RGBA { //gd:Color+float
	return RGBA{a.R + Float.X(b), a.G + Float.X(b), a.B + Float.X(b), a.A + Float.X(b)}
}
func SubX[X Float.Any | Int.Any](a RGBA, b X) RGBA { //gd:Color-float
	return RGBA{a.R - Float.X(b), a.G - Float.X(b), a.B - Float.X(b), a.A - Float.X(b)}
}
func MulX[X Float.Any | Int.Any](a RGBA, b X) RGBA { //gd:Color*float
	return RGBA{a.R * Float.X(b), a.G * Float.X(b), a.B * Float.X(b), a.A * Float.X(b)}
}
func DivX[X Float.Any | Int.Any](a RGBA, b X) RGBA { //gd:Color/float
	return RGBA{a.R / Float.X(b), a.G / Float.X(b), a.B / Float.X(b), a.A / Float.X(b)}
}
