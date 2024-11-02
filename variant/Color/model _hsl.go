package Color

import (
	"math"

	"grow.graphics/gd/variant/Float"
)

// HSL constructs a color from an OK HSL profile. The hue (h), saturation (s), and
// lightness (l) are typically between 0.0 and 1.0.
func HSL(h, s, l float64) RGBA { //gd:Color.from_ok_hsl
	if l == 1.0 {
		return RGBA{1, 1, 1, 1}
	} else if l == 0 {
		return RGBA{}
	}

	var a_ = math.Cos(2 * math.Pi * h)
	var b_ = math.Sin(2 * math.Pi * h)
	var L = toe_inv(l)

	var C_0, C_mid, C_max = get_Cs(L, a_, b_)

	var mid = 0.8
	var mid_inv = 1.25

	var C, t, k_0, k_1, k_2 float64

	if s < mid {
		t = mid_inv * s

		k_1 = mid * C_0
		k_2 = (1 - k_1/C_mid)

		C = t * k_1 / (1 - k_2*t)
	} else {
		t = (s - mid) / (1 - mid)

		k_0 = C_mid
		k_1 = (1 - mid) * C_mid * C_mid * mid_inv * mid_inv / C_0
		k_2 = (1 - (k_1)/(C_max-C_mid))

		C = k_0 + t*k_1/(1-k_2*t)
	}
	r, g, b := oklab_to_linear_srgb(L, C*a_, C*b_)
	return Clamp(RGBA{
		Float.X(srgb_transfer_function(r)),
		Float.X(srgb_transfer_function(g)),
		Float.X(srgb_transfer_function(b)),
		1,
	}, RGBA{}, RGBA{1, 1, 1, 1})
}

// HSLA constructs a color from an OK HSL profile. The hue (h), saturation (s), and
// lightness (l) are typically between 0.0 and 1.0. Includes Alpha.
func HSLA(h, s, l, a float64) RGBA {
	c := HSL(h, s, l)
	c.A = float32(a)
	return c
}

func toe_inv(x float64) float64 {
	const k_1 = 0.206
	const k_2 = 0.03
	const k_3 = (1 + k_1) / (1 + k_2)
	return (x*x + k_1*x) / (k_3 * (x + k_2))
}

func get_Cs(L, a_, b_ float64) (c0, cmid, cmax float64) {
	var L_cusp, C_cusp = find_cusp(a_, b_)

	var C_max = find_gamut_intersection(a_, b_, L, 1, L, L_cusp, C_cusp)
	var ST_max_S, ST_max_T = to_ST(L_cusp, C_cusp)

	// Scale factor to compensate for the curved part of gamut shape:
	var k = C_max / min((L*ST_max_S), (1-L)*ST_max_T)

	var C_mid float64
	{
		var ST_mid_S, ST_mid_T = get_ST_mid(a_, b_)

		// Use a soft minimum function, instead of a sharp triangle shape to get a smooth value for chroma.
		var C_a = L * ST_mid_S
		var C_b = (1. - L) * ST_mid_T
		C_mid = 0.9 * k * math.Sqrt(math.Sqrt(1/(1/(C_a*C_a*C_a*C_a)+1/(C_b*C_b*C_b*C_b))))
	}

	var C_0 float64
	{
		// for C_0, the shape is independent of hue, so ST are constant. Values picked to roughly be the average values of ST.
		var C_a = L * 0.4
		var C_b = (1 - L) * 0.8

		// Use a soft minimum function, instead of a sharp triangle shape to get a smooth value for chroma.
		C_0 = math.Sqrt(1 / (1/(C_a*C_a) + 1/(C_b*C_b)))
	}

	return C_0, C_mid, C_max
}

// finds L_cusp and C_cusp for a given hue
// a and b must be normalized so a^2 + b^2 == 1
func find_cusp(a, b float64) (L_cusp, C_cusp float64) {
	// First, find the maximum saturation (saturation S = C/L)
	var S_cusp = compute_max_saturation(a, b)
	// Convert to linear sRGB to find the first point where at least one of r,g or b >= 1:
	var sr, sg, sb = oklab_to_linear_srgb(1, S_cusp*a, S_cusp*b)
	L_cusp = math.Cbrt(1 / max(max(sr, sg), sb))
	C_cusp = L_cusp * S_cusp
	return
}

// Finds the maximum saturation possible for a given hue that fits in sRGB
// Saturation here is defined as S = C/L
// a and b must be normalized so a^2 + b^2 == 1
func compute_max_saturation(a, b float64) float64 {
	// Max saturation will be when one of r, g or b goes below zero.
	// Select different coefficients depending on which component goes below zero first
	var k0, k1, k2, k3, k4, wl, wm, ws float64
	if -1.88170328*a-0.80936493*b > 1 {
		// Red component
		k0 = +1.19086277
		k1 = +1.76576728
		k2 = +0.59662641
		k3 = +0.75515197
		k4 = +0.56771245
		wl = +4.0767416621
		wm = -3.3077115913
		ws = +0.2309699292
	} else if 1.81444104*a-1.19445276*b > 1 {
		// Green component
		k0 = +0.73956515
		k1 = -0.45954404
		k2 = +0.08285427
		k3 = +0.12541070
		k4 = +0.14503204
		wl = -1.2684380046
		wm = +2.6097574011
		ws = -0.3413193965
	} else {
		// Blue component
		k0 = +1.35733652
		k1 = -0.00915799
		k2 = -1.15130210
		k3 = -0.50559606
		k4 = +0.00692167
		wl = -0.0041960863
		wm = -0.7034186147
		ws = +1.7076147010
	}
	// Approximate max saturation using a polynomial:
	var S = k0 + k1*a + k2*b + k3*a*a + k4*a*b
	// Do one step Halley's method to get closer
	// this gives an error less than 10e6, except for some blue hues where the dS/dh is close to infinite
	// this should be sufficient for most applications, otherwise do two/three steps
	var (
		k_l = +0.3963377774*a + 0.2158037573*b
		k_m = -0.1055613458*a - 0.0638541728*b
		k_s = -0.0894841775*a - 1.2914855480*b
	)
	{
		var (
			l_ = 1 + S*k_l
			m_ = 1 + S*k_m
			s_ = 1 + S*k_s
		)
		var (
			l = l_ * l_ * l_
			m = m_ * m_ * m_
			s = s_ * s_ * s_
		)
		var (
			l_dS = 3 * k_l * l_ * l_
			m_dS = 3 * k_m * m_ * m_
			s_dS = 3 * k_s * s_ * s_
		)
		var (
			l_dS2 = 6 * k_l * k_l * l_
			m_dS2 = 6 * k_m * k_m * m_
			s_dS2 = 6 * k_s * k_s * s_
		)
		var (
			f  = wl*l + wm*m + ws*s
			f1 = wl*l_dS + wm*m_dS + ws*s_dS
			f2 = wl*l_dS2 + wm*m_dS2 + ws*s_dS2
		)
		S = S - f*f1/(f1*f1-0.5*f*f2)
	}
	return S
}

func oklab_to_linear_srgb(L, a, b_ float64) (r, g, b float64) {
	var (
		l_ = L + 0.3963377774*a + 0.2158037573*b_
		m_ = L - 0.1055613458*a - 0.0638541728*b_
		s_ = L - 0.0894841775*a - 1.2914855480*b_
	)
	var (
		l = l_ * l_ * l_
		m = m_ * m_ * m_
		s = s_ * s_ * s_
	)
	return +4.0767416621*l - 3.3077115913*m + 0.2309699292*s,
		-1.2684380046*l + 2.6097574011*m - 0.3413193965*s,
		-0.0041960863*l - 0.7034186147*m + 1.7076147010*s
}

// Finds intersection of the line defined by
// L = L0 * (1 - t) + t * L1;
// C = t * C1;
// a and b must be normalized so a^2 + b^2 == 1
func find_gamut_intersection(a, b, L1, C1, L0, L_cusp, C_cusp float64) float64 {
	// Find the intersection for upper and lower half seprately
	var t float64
	if ((L1-L0)*C_cusp - (L_cusp-L0)*C1) <= 0. {
		// Lower half

		t = C_cusp * L0 / (C1*L_cusp + C_cusp*(L0-L1))
	} else {
		// Upper half

		// First intersect with triangle
		t = C_cusp * (L0 - 1) / (C1*(L_cusp-1) + C_cusp*(L0-L1))

		// Then one step Halley's method
		{
			var dL = L1 - L0
			var dC = C1

			var k_l = +0.3963377774*a + 0.2158037573*b
			var k_m = -0.1055613458*a - 0.0638541728*b
			var k_s = -0.0894841775*a - 1.2914855480*b

			var l_dt = dL + dC*k_l
			var m_dt = dL + dC*k_m
			var s_dt = dL + dC*k_s

			// If higher accuracy is required, 2 or 3 iterations of the following block can be used:
			{
				var L = L0*(1-t) + t*L1
				var C = t * C1

				var l_ = L + C*k_l
				var m_ = L + C*k_m
				var s_ = L + C*k_s

				var l = l_ * l_ * l_
				var m = m_ * m_ * m_
				var s = s_ * s_ * s_

				var ldt = 3 * l_dt * l_ * l_
				var mdt = 3 * m_dt * m_ * m_
				var sdt = 3 * s_dt * s_ * s_

				var ldt2 = 6 * l_dt * l_dt * l_
				var mdt2 = 6 * m_dt * m_dt * m_
				var sdt2 = 6 * s_dt * s_dt * s_

				var r = 4.0767416621*l - 3.3077115913*m + 0.2309699292*s - 1
				var r1 = 4.0767416621*ldt - 3.3077115913*mdt + 0.2309699292*sdt
				var r2 = 4.0767416621*ldt2 - 3.3077115913*mdt2 + 0.2309699292*sdt2

				var u_r = r1 / (r1*r1 - 0.5*r*r2)
				var t_r = -r * u_r

				var g = -1.2684380046*l + 2.6097574011*m - 0.3413193965*s - 1
				var g1 = -1.2684380046*ldt + 2.6097574011*mdt - 0.3413193965*sdt
				var g2 = -1.2684380046*ldt2 + 2.6097574011*mdt2 - 0.3413193965*sdt2

				var u_g = g1 / (g1*g1 - 0.5*g*g2)
				var t_g = -g * u_g

				b = -0.0041960863*l - 0.7034186147*m + 1.7076147010*s - 1
				var b1 = -0.0041960863*ldt - 0.7034186147*mdt + 1.7076147010*sdt
				var b2 = -0.0041960863*ldt2 - 0.7034186147*mdt2 + 1.7076147010*sdt2

				var u_b = b1 / (b1*b1 - 0.5*b*b2)
				var t_b = -b * u_b

				if u_r < 0 {
					t_r = math.MaxFloat64
				}
				if u_g < 0 {
					t_g = math.MaxFloat64
				}
				if u_b < 0 {
					t_b = math.MaxFloat64
				}
				t += min(t_r, min(t_g, t_b))
			}
		}
	}
	return t
}

// Alternative representation of (L_cusp, C_cusp)
// Encoded so S = C_cusp/L_cusp and T = C_cusp/(1-L_cusp)
// The maximum value for C in the triangle is then found as fmin(S*L, T*(1-L)), for a given L
func to_ST(cusp_L, cusp_C float64) (ST_S, ST_T float64) {
	return cusp_C / cusp_L, cusp_C / (1 - cusp_L)
}

// Returns a smooth approximation of the location of the cusp
// This polynomial was created by an optimization process
// It has been designed so that S_mid < S_max and T_mid < T_max
func get_ST_mid(a_, b_ float64) (s, t float64) {
	var S = 0.11516993 + 1/(+7.44778970+4.15901240*b_+
		a_*(-2.19557347+1.75198401*b_+
			a_*(-2.13704948-10.02301043*b_+
				a_*(-4.24894561+5.38770819*b_+4.69891013*a_))))
	var T = 0.11239642 + 1/(+1.61320320-0.68124379*b_+
		a_*(+0.40370612+0.90148123*b_+
			a_*(-0.27087943+0.61223990*b_+
				a_*(+0.00299215-0.45399568*b_-0.14661872*a_))))
	return S, T
}

func srgb_transfer_function(a float64) float64 {
	if 0.0031308 >= a {
		return 12.92 * a
	}
	return 1.055*math.Pow(a, .4166666666666667) - .055
}
