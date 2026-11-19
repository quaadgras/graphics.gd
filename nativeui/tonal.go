package nativeui

import (
	"math"

	"graphics.gd/variant/Color"
)

// tonal builds Material's tonal colours: a hue and chroma at a given tone,
// where tone is CIE L* (0 black, 100 white), as Material's HCT tone is.
// Chroma is in CIE LCh(ab) units, which approximates HCT's well enough for
// deriving a dynamic scheme from one accent colour.
type tonal struct{ hue, chroma float64 }

func tonalOf(c Color.RGBA) tonal {
	l, a, b := toLab(c)
	_ = l
	return tonal{hue: math.Atan2(b, a), chroma: math.Hypot(a, b)}
}

func (p tonal) withChroma(c float64) tonal { p.chroma = c; return p }

// tone returns the colour at tone t, reducing chroma until it fits sRGB.
func (p tonal) tone(t float64) Color.RGBA {
	for c := p.chroma; ; c *= 0.9 {
		r, g, b, ok := fromLab(t, c*math.Cos(p.hue), c*math.Sin(p.hue))
		if ok || c < 0.5 {
			return Color.RGBA{R: float32(clamp(r, 0, 1)), G: float32(clamp(g, 0, 1)), B: float32(clamp(b, 0, 1)), A: 1}
		}
	}
}

func lin(c float64) float64 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}

func gam(c float64) float64 {
	if c <= 0.0031308 {
		return 12.92 * c
	}
	return 1.055*math.Pow(c, 1/2.4) - 0.055
}

func toLab(c Color.RGBA) (l, a, b float64) {
	r, g, bl := lin(float64(c.R)), lin(float64(c.G)), lin(float64(c.B))
	x := (0.4124*r + 0.3576*g + 0.1805*bl) / 0.95047
	y := 0.2126*r + 0.7152*g + 0.0722*bl
	z := (0.0193*r + 0.1192*g + 0.9505*bl) / 1.08883
	f := func(t float64) float64 {
		if t > 216.0/24389 {
			return math.Cbrt(t)
		}
		return (24389.0/27*t + 16) / 116
	}
	fx, fy, fz := f(x), f(y), f(z)
	return 116*fy - 16, 500 * (fx - fy), 200 * (fy - fz)
}

func fromLab(l, a, b float64) (r, g, bl float64, ok bool) {
	fy := (l + 16) / 116
	fx, fz := fy+a/500, fy-b/200
	inv := func(t float64) float64 {
		if t*t*t > 216.0/24389 {
			return t * t * t
		}
		return (116*t - 16) / (24389.0 / 27)
	}
	x, y, z := inv(fx)*0.95047, inv(fy), inv(fz)*1.08883
	r = gam(3.2406*x - 1.5372*y - 0.4986*z)
	g = gam(-0.9689*x + 1.8758*y + 0.0415*z)
	bl = gam(0.0557*x - 0.2040*y + 1.0570*z)
	const e = 0.002
	ok = r >= -e && r <= 1+e && g >= -e && g <= 1+e && bl >= -e && bl <= 1+e
	return r, g, bl, ok
}
