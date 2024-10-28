package Color

import (
	"grow.graphics/gd/gdmaths/Float"
)

// RGBE9995 decodes a Color from a RGBE9995 format integer where the three color components have
// 9 bits of precision and all three share a single 5-bit exponent.
func RGBE9995(rgbe uint32) RGBA { //gd:Color.from_rgbe9995
	var r = Float.X(rgbe & 0x1ff)
	var g = Float.X((rgbe >> 9) & 0x1ff)
	var b = Float.X((rgbe >> 18) & 0x1ff)
	var e = Float.X((rgbe >> 27))
	var m = Float.Pow(2.0, e-15.0-9.0)
	var (
		rd = r * m
		gd = g * m
		bd = b * m
	)
	return RGBA{rd, gd, bd, 1.0}
}

// HSV constructs a color from an HSV profile. The hue (h), saturation (s), and value (v) are typically
// between 0.0 and 1.0.
func HSV(h, s, v Float.X) RGBA { //gd:Color.from_hsv
	var (
		i, f, p, q, t Float.X
		a             Float.X = 1
	)
	if s == 0.0 {
		return RGBA{v, v, v, a} // Achromatic (gray)
	}
	h *= 6.0
	h = Float.Mod(h, 6)
	i = Float.Floor(h)
	f = h - i
	p = v * (1.0 - s)
	q = v * (1.0 - s*f)
	t = v * (1.0 - s*(1.0-f))
	switch int(i) {
	case 0: // Red is the dominant color
		return RGBA{v, t, p, a}
	case 1: // Green is the dominant color
		return RGBA{q, v, p, a}
	case 2:
		return RGBA{p, v, t, a}
	case 3: // Blue is the dominant color
		return RGBA{p, q, v, a}
	case 4:
		return RGBA{t, p, v, a}
	default: // (5) Red is the dominant color
		return RGBA{v, p, q, a}
	}
}

// HSVA constructs a color from an HSV profile. The hue (h), saturation (s), and value (v) are typically
// between 0.0 and 1.0. Includes alpha.
func HSVA(h, s, v, a Float.X) RGBA {
	c := HSV(h, s, v)
	c.A = a
	return c
}

// Uint32 returns the Color associated with the provided hex integer in 32-bit RGBA format (8 bits per channel).
//
// The int is best visualized with hexadecimal notation ("0x" prefix, making it "0xRRGGBBAA").
func Uint32(hex uint32) RGBA { //gd:Color.hex
	var a = Float.X(hex&0xFF) / 255
	hex >>= 8
	var b = Float.X(hex&0xFF) / 255
	hex >>= 8
	var g = Float.X(hex&0xFF) / 255
	hex >>= 8
	var r = Float.X(hex&0xFF) / 255
	return RGBA{r, g, b, a}
}

// AsUint32 returns the color as a 32-bit RGBA integer.
func AsUint32(c RGBA) uint32 { //gd:Color.to_rgba32
	var r = uint32(c.R * 255)
	var g = uint32(c.G * 255)
	var b = uint32(c.B * 255)
	var a = uint32(c.A * 255)
	return r | g<<8 | b<<16 | a<<24
}

// Uint64 returns the Color associated with the provided hex integer in 64-bit RGBA format (16 bits per channel).
//
// The int is best visualized with hexadecimal notation ("0x" prefix, making it "0xRRRRGGGGBBBBAAAA").
func Uint64(hex int64) RGBA { //gd:Color.hex64
	var a = Float.X(hex&0xFFFF) / 65535
	hex >>= 16
	var b = Float.X(hex&0xFFFF) / 65535
	hex >>= 16
	var g = Float.X(hex&0xFFFF) / 65535
	hex >>= 16
	var r = Float.X(hex&0xFFFF) / 65535
	return RGBA{r, g, b, a}
}

// Hex returns a new color from rgba, an HTML hexadecimal color string. rgba is not case-sensitive
// and may be prefixed by a hash sign (#).
//
// rgba must be a valid three-digit or six-digit hexadecimal color string, and may contain an alpha
// channel value. If rgba does not contain an alpha channel value, an alpha channel value of 1.0 is
// applied. If rgba is invalid, returns an empty color.
func Hex(rgba string) RGBA { //gd:Color.html
	var color = rgba
	if len(color) == 0 {
		return RGBA{}
	}
	if color[0] == '#' {
		color = color[1:]
	}
	// If enabled, use 1 hex digit per channel instead of 2.
	// Other sizes aren't in the HTML/CSS spec but we could add them if desired.
	var (
		is_shorthand = len(color) < 5
		alpha        = false
	)
	if len(color) == 8 {
		alpha = true
	} else if len(color) == 6 {
		alpha = false
	} else if len(color) == 4 {
		alpha = true
	} else if len(color) == 3 {
		alpha = false
	}
	var r, g, b, a Float.X = 1.0, 1.0, 1.0, 1.0
	if is_shorthand {
		r = _parse_col4(color, 0) / 15
		g = _parse_col4(color, 1) / 15
		b = _parse_col4(color, 2) / 15
		if alpha {
			a = _parse_col4(color, 3) / 15
		}
	} else {
		r = _parse_col8(color, 0) / 255
		g = _parse_col8(color, 2) / 255
		b = _parse_col8(color, 4) / 255
		if alpha {
			a = _parse_col8(color, 6) / 255
		}
	}
	return RGBA{r, g, b, a}
}

func _parse_col4(s string, ofs int) Float.X {
	var character = s[ofs]
	if character >= '0' && character <= '9' {
		return Float.X(character - '0')
	} else if character >= 'a' && character <= 'f' {
		return Float.X(character) + Float.X(10-'a')
	} else if character >= 'A' && character <= 'F' {
		return Float.X(character) + Float.X(10-'A')
	}
	return -1
}

func _parse_col8(s string, ofs int) Float.X {
	return _parse_col4(s, ofs)*16 + _parse_col4(s, ofs+1)
}

// ValidHex returns true if color is a valid HTML hexadecimal color string. The string must be a hexadecimal
// value (case-insensitive) of either 3, 4, 6 or 8 digits, and may be prefixed by a hash sign (#).
func ValidHex(color string) bool { //gd:Color.html_is_valid
	if len(color) == 0 {
		return false
	}
	if color[0] == '#' {
		color = color[1:]
	}
	// Check if the amount of hex digits is valid.
	if !(len(color) == 3 || len(color) == 4 || len(color) == 6 || len(color) == 8) {
		return false
	}
	// Check if each hex digit is valid.
	for i := 0; i < len(color); i++ {
		if _parse_col4(color, i) == -1 {
			return false
		}
	}
	return true
}
