package gtkrc

import (
	"image"
	"math"
	"strconv"
	"strings"
)

// Clearlooks, the GUMMY style (the installed theme's) in detail; the other
// styles fall back to it. Implemented from a behavioural specification of
// gtk-engines' Clearlooks, not from its source.

func init() { engines["clearlooks"] = clearlooks }

type clPalette struct {
	shade        [9]RGBA
	spot         [3]RGBA
	radius       float64
	relief       bool
	menubarStyle int
	parent       RGBA
}

func clearlooksPalette(s *Style) clPalette {
	var p clPalette
	contrast := 1.0
	if v, ok := s.EngineParams["contrast"]; ok {
		if f, err := strconv.ParseFloat(unquote(v), 64); err == nil && f > 0 {
			contrast = f
		}
	}
	for i, f := range []float64{1.15, 0.95, 0.896, 0.82, 0.70, 0.665, 0.475, 0.45, 0.40} {
		if f < 1 {
			f /= contrast
		} else {
			f *= contrast
		}
		p.shade[i] = shade(s.Bg[Normal], f)
	}
	for i, f := range []float64{1.25, 1.05, 0.65} {
		p.spot[i] = shade(s.Bg[Selected], f)
	}
	p.radius = 3
	if v, ok := s.EngineParams["radius"]; ok {
		if f, err := strconv.ParseFloat(unquote(v), 64); err == nil {
			p.radius = math.Max(0, math.Min(10, f))
		}
	}
	if v, ok := s.EngineParams["reliefstyle"]; ok {
		p.relief = strings.TrimSpace(unquote(v)) != "0"
	}
	if v, ok := s.EngineParams["menubarstyle"]; ok {
		p.menubarStyle, _ = strconv.Atoi(strings.TrimSpace(unquote(v)))
	}
	p.parent = s.Bg[Normal] // the colour behind the widget: the window's
	return p
}

// gummy is the four-stop "Gummy gradient" of a colour.
func gummy(c RGBA, a, b, d float64) ([]float64, []RGBA) {
	return []float64{0, 0.5, 0.5, 1}, []RGBA{shade(c, a), shade(c, b), c, shade(c, d)}
}

// clInset is INSET: the outermost ring split along the diagonal, darker
// upper-left, lighter lower-right.
func clInset(dst *image.NRGBA, ox, oy, w, h, r float64, corners [4]bool, b RGBA) {
	m := math.Min(w, h)
	upper := func(px, py float64) bool {
		lx := px - ox
		var y float64
		switch {
		case lx < m/2:
			y = h - lx
		case lx > w-m/2:
			y = w - lx
		default:
			y = h - m/2 - (lx-m/2)*(h-m)/math.Max(w-m, 1e-9)
		}
		return py-oy < y
	}
	strokeRR(dst, ox+0.5, oy+0.5, w-1, h-1, r, corners, 1, solid(shade(b, 0.94)), upper)
	strokeRR(dst, ox+0.5, oy+0.5, w-1, h-1, r, corners, 1, solid(shade(b, 1.06)), func(px, py float64) bool { return !upper(px, py) })
}

// clTLH is the top-left highlight: a 1px line up the left edge and along the
// top of a box, in shade(c, 1.3) at alpha a.
func clTLH(dst *image.NRGBA, x, y, w, h, r float64, corners [4]bool, c RGBA, a float64) {
	col := withAlpha(shade(c, 1.3), a)
	bottom, right := y+h, x+w
	if corners[3] {
		bottom -= r
	}
	if corners[1] {
		right -= r
	}
	strokeRR(dst, x+0.5, y+0.5, w, h, r, corners, 1, solid(col), func(px, py float64) bool {
		return (px < x+1.5 && py < bottom) || (py < y+1.5 && px < right)
	})
}

func clCorners(tl, tr, br, bl bool) [4]bool { return [4]bool{tl, tr, br, bl} }

func clearlooks(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) bool {
	p := clearlooksPalette(s)
	st := c.State
	x, y, w, h := float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy())
	disabled := st == Insensitive
	prelight := st == Prelight
	switch c.Function {
	case "BOX", "SHADOW":
		switch c.Detail {
		case "button", "optionmenu":
			corners := allCorners
			relief := p.relief
			if c.Widget == "GtkComboBoxEntry" {
				corners, relief = clCorners(false, true, true, false), false
			}
			clButton(s, p, dst, x, y, w, h, st, c.Shadow == "IN", relief, corners)
			return true
		case "buttondefault":
			return true
		case "spinbutton_up", "spinbutton_down":
			corners := clCorners(false, true, false, false)
			if c.Detail == "spinbutton_down" {
				corners = clCorners(false, false, true, false)
			}
			if st != Active {
				clButton(s, p, dst, x-1, y, w+1, h, Normal, false, false, corners)
				return true
			}
			rr := math.Min(p.radius, math.Min((w-4)/2, (h-4)/2))
			fillRR(dst, x+2, y+2, w-4, h-4, rr, corners, solid(s.Bg[st]))
			fillRR(dst, x+2, y+2, w-4, h-4, rr, corners, vgrad(y, y+h, []float64{0, 1}, []RGBA{shade(s.Bg[Normal], 0.8), withAlpha(shade(s.Bg[Normal], 0.8), 0)}))
			return true
		case "entry":
			clEntry(s, p, dst, x, y, w, h, st, c.Focused)
			return true
		case "trough", "trough-lower", "trough-upper":
			if c.Widget == "GtkProgressBar" {
				clProgressTrough(p, dst, x, y, w, h)
			} else {
				clScaleTrough(p, dst, x, y, w, h, c.Detail == "trough-lower")
			}
			return true
		case "bar":
			clProgressBar(s, p, dst, x, y, w, h)
			return true
		case "menu":
			fill(dst, r, s.Bg[Normal])
			ring(dst, r, p.shade[5])
			return true
		case "menuitem":
			offs, cols := gummy(p.spot[1], 1.08, 1.02, 0.94)
			fillRR(dst, x+0.5, y+0.5, w-1, h-1, 0, allCorners, vgrad(y, y+h, offs, cols))
			ring(dst, r, p.spot[2])
			return true
		case "menubar":
			if c.Shadow == "NONE" {
				return true
			}
			if p.menubarStyle == 2 || p.menubarStyle == 1 {
				fillRR(dst, x, y, w, h, 0, allCorners, vgrad(y, y+h, []float64{0, 1}, []RGBA{s.Bg[Normal], shade(s.Bg[Normal], 0.96)}))
			}
			hline(dst, r.Min.X, r.Max.X-1, r.Max.Y-1, p.shade[3])
			return true
		case "scrolled_window", "viewport", "":
			if c.Function == "SHADOW" {
				ring(dst, r, p.shade[5])
				return true
			}
		case "frame":
			if c.Function == "SHADOW" {
				clFrame(s, p, dst, x, y, w, h, c.Shadow, [4]bool{}, p.shade[4])
				return true
			}
		}
		if c.Function == "SHADOW" {
			clFrame(s, p, dst, x, y, w, h, c.Shadow, allCorners, p.shade[5])
			return true
		}
		return false
	case "SLIDER":
		clSlider(s, p, dst, x, y, w, h, st)
		return true
	case "CHECK", "OPTION":
		clCheck(s, p, dst, x, y, w, h, st, c.Shadow, c.Function == "OPTION", disabled, prelight)
		return true
	case "ARROW", "TAB":
		col := s.Fg[st]
		cx, cy := x+w/2, y+h/2
		double := c.Function == "TAB" || c.Widget == "GtkComboBox"
		draw := func(dx, dy float64, col RGBA) {
			if double {
				aw := math.Min(2*h/3, w)
				clChevron(dst, cx+dx, cy-aw/2+dy, aw, "UP", col)
				clChevron(dst, cx+dx, cy+aw/2+dy, aw, "DOWN", col)
			} else {
				clChevron(dst, cx+dx, cy+dy, w, c.ArrowDirection, col)
			}
		}
		if disabled {
			draw(0.5, 0.5, p.shade[0])
			col = s.Fg[Insensitive]
		}
		draw(0, 0, col)
		return true
	case "HLINE":
		yy := r.Min.Y + r.Dy()/2
		if c.Detail == "menuitem" {
			hline(dst, r.Min.X, r.Max.X-1, yy, p.shade[5])
			return true
		}
		hline(dst, r.Min.X, r.Max.X-1, yy-1, p.shade[3])
		hline(dst, r.Min.X, r.Max.X-1, yy, shade(p.shade[3], 1.3))
		return true
	case "VLINE":
		xx := r.Min.X + r.Dx()/2 - 1
		vline(dst, xx, r.Min.Y, r.Max.Y-1, p.shade[3])
		vline(dst, xx+1, r.Min.Y, r.Max.Y-1, shade(p.shade[3], 1.3))
		return true
	case "BOX_GAP":
		corners := clCorners(false, true, true, true) // tabs on top, starting at the left edge
		fillRR(dst, x, y, w, h, p.radius, corners, solid(s.Bg[Normal]))
		clFrame(s, p, dst, x, y, w, h, c.Shadow, corners, p.shade[5])
		return true
	case "EXTENSION":
		clTab(s, p, dst, r, c)
		return true
	case "FOCUS":
		return true // Gummy draws button focus outside the button; nativeui shows none
	case "FLAT_BOX":
		if (c.Detail == "checkbutton" || c.Detail == "radiobutton") && prelight {
			return true
		}
	}
	return false
}

func clButton(s *Style, p clPalette, dst *image.NRGBA, x, y, w, h float64, st int, active, relief bool, corners [4]bool) {
	fillCol := s.Bg[st]
	disabled := st == Insensitive
	o := 0.0
	if s.XThickness >= 3 {
		o = 1
	}
	r := math.Min(p.radius, math.Min((w-2-2*o)/2, (h-2-2*o)/2))
	if s.XThickness >= 3 || s.YThickness >= 3 {
		if relief && !active && !disabled {
			r = math.Min(p.radius, math.Min((w-2-2*o)/2-1, (h-2-2*o)/2-1))
			strokeRR(dst, x+0.5, y+0.5, w-1, h-1, r+1, corners, 1, solid(shade(p.parent, 0.97)), nil)
			strokeRR(dst, x+1.5, y+1.5, w-2, h-2, r+1, corners, 1, solid(shade(p.parent, 0.93)), nil)
		} else {
			clInset(dst, x, y, w, h, p.radius+1, corners, p.parent)
		}
	}
	bx, by, bw, bh := x+o+1, y+o+1, w-2*o-2, h-2*o-2
	g := []float64{1.08, 1.02, 0.94}
	if disabled {
		g = []float64{1.04, 1.01, 0.96}
	}
	offs, cols := gummy(fillCol, g[0], g[1], g[2])
	if disabled {
		cols[2] = shade(fillCol, 0.99)
	}
	fillRR(dst, bx, by, bw, bh, math.Floor(r), corners, vgrad(by, by+bh, offs, cols))
	if active {
		sc := shade(fillCol, 0.92)
		clip := rrSDF(bx, by, bw, h, r, [4]bool{corners[0], corners[1], false, corners[3]})
		in := func(px, py float64) bool { return clip(px, py) <= 0 }
		fillRegionFunc(dst, image.Rect(int(bx), int(by), int(bx+bw), int(by+3)), func(px, py float64) RGBA {
			return withAlpha(sc, 0.58*(1-(py-by)/3))
		}, in)
		fillRegionFunc(dst, image.Rect(int(bx), int(by), int(bx+3), int(by+bh)), func(px, py float64) RGBA {
			return withAlpha(sc, 0.58*(1-(px-bx)/3))
		}, in)
	}
	border := mix(p.shade[6], fillCol, 0.2)
	if disabled {
		border = p.shade[4]
	}
	strokeRR(dst, x+o+0.5, y+o+0.5, w-2*o-1, h-2*o-1, r, corners, 1, solid(border), nil)
	if !active {
		clTLH(dst, bx, by, bw, bh, r, corners, fillCol, 0.4)
	}
}

func clEntry(s *Style, p clPalette, dst *image.NRGBA, x, y, w, h float64, st int, focused bool) {
	base := s.Base[st]
	focus := p.spot[2]
	if v, ok := s.EngineParams["focus_color"]; ok {
		if c, ok := (&Theme{Scheme: s.scheme}).color(v, s); ok {
			focus = c
		}
	}
	bd := p.shade[6]
	if st == Insensitive {
		bd = p.shade[4]
	}
	if focused {
		bd = focus
	}
	r := math.Min(p.radius, math.Min((w-4)/2, (h-4)/2))
	o := 0.0
	fill(dst, image.Rect(int(x), int(y), int(x+w), int(y+h)), p.parent)
	if s.XThickness >= 3 && s.YThickness >= 3 {
		clInset(dst, x, y, w, h, r+1, allCorners, p.parent)
		o = 1
	}
	ri := math.Max(0, r-1)
	fillRR(dst, x+o+1, y+o+1, w-2*(o+1), h-2*(o+1), ri, allCorners, solid(base))
	if focused {
		strokeRR(dst, x+o+1.5, y+o+1.5, w-2*(o+1)-1, h-2*(o+1)-1, ri, allCorners, 1, solid(mix(base, shade(bd, 1.61), 0.5)), nil)
	} else {
		a := 0.18
		if st == Insensitive {
			a = 0.09
		}
		strokeRR(dst, x+o+1.5, y+o+1.5, w, h, ri, allCorners, 1, solid(withAlpha(shade(bd, 0.92), a)), func(px, py float64) bool {
			return (px < x+o+2 && py < y+h-r) || (py < y+o+2 && px < x+w-r)
		})
	}
	strokeRR(dst, x+o+0.5, y+o+0.5, w-2*o-1, h-2*o-1, r, allCorners, 1, solid(bd), nil)
}

func clCheck(s *Style, p clPalette, dst *image.NRGBA, x, y, w, h float64, st int, shadow string, radio, disabled, prelight bool) {
	border, mark := p.shade[6], s.Text[Normal]
	if prelight {
		border = p.spot[2]
	}
	if disabled {
		border, mark = p.shade[5], p.shade[6]
	}
	fillCol := s.Base[Normal]
	if prelight {
		fillCol = mix(s.Base[Normal], p.spot[1], 0.5)
	}
	if !radio {
		r1 := 0.0
		if p.radius > 0 {
			r1 = 1
		}
		bx, by, bw, bh := x+0.5, y+0.5, w-1, h-1
		if s.XThickness >= 3 && s.YThickness >= 3 {
			clInset(dst, x, y, w, h, r1, allCorners, p.parent)
			bx, by, bw, bh = x+1.5, y+1.5, w-3, h-3
		}
		if !disabled {
			fillRR(dst, bx, by, bw, bh, r1, allCorners, solid(fillCol))
		}
		strokeRR(dst, bx, by, bw, bh, r1, allCorners, 1, solid(border), nil)
		switch shadow {
		case "ETCHED_IN":
			segment(dst, x+3, y+h/2, x+w-3, y+h/2, 2, mark, true)
		case "IN":
			pts := [][2]float64{{x + 0.5 + 0.2*w, y + 0.5*h}, {x + 0.5 + 0.4*w, y + 0.7*h}}
			pts = append(pts, bezier(pts[1], [2]float64{x + 0.5 + 0.4*w, y + 0.7*h}, [2]float64{x + 0.5 + 0.5*w, y + 0.4*h}, [2]float64{x + 0.5 + 0.7*w, y + 0.25*h}, 10)...)
			polyline(dst, 1.7, mark, pts...)
		}
		return
	}
	sh, hl := shade(p.parent, 0.9), shade(p.parent, 1.1)
	cx, cy := w/2, h/2
	R := math.Min(w, h) / 2
	ccx, ccy := x+math.Ceil(cx), y+math.Ceil(cy)
	swapped := RGBA{sh.R, sh.B, sh.G, 1}
	lw := math.Max(1, math.Floor(R/3))
	rr := math.Floor(R - 0.1)
	fillRegionFunc(dst, image.Rect(int(x)-1, int(y)-1, int(x+w)+2, int(y+h)+2), func(px, py float64) RGBA {
		t := ((px - x) + (py - y)) / (4 * R)
		return stops(t, []float64{0, 0.5, 0.5, 1}, []RGBA{swapped, withAlpha(swapped, 0.5), withAlpha(hl, 0.5), hl})
	}, func(px, py float64) bool { return math.Abs(math.Hypot(px-ccx, py-ccy)-rr) <= lw/2 })
	dr := math.Max(1, math.Ceil(R)-1.5)
	if !disabled {
		fillDisc(dst, ccx, ccy, dr, fillCol, nil)
	}
	strokeCircle(dst, ccx, ccy, dr, math.Max(1, math.Floor(R/6)), border)
	switch shadow {
	case "IN":
		fillDisc(dst, ccx, ccy, math.Floor(R/2), mark, nil)
		fillDisc(dst, x+math.Floor(cx-R/10), y+math.Floor(cy-R/10), math.Floor(R/6), withAlpha(hl, 0.5), nil)
	case "ETCHED_IN":
		lw := math.Floor(2 * R / 3)
		yy := y + math.Ceil(cy-lw) + lw
		segment(dst, x+math.Ceil(cx-R/3-lw)+lw, yy, x+math.Floor(cx+R/3+lw)-lw, yy, lw, mark, false)
	}
}

// clChevron fills a down-pointing "V" (rotated for other directions) centred
// at (cx, cy) in a box aw wide.
func clChevron(dst *image.NRGBA, cx, cy, aw float64, dir string, c RGBA) {
	t := math.Max(1, math.Ceil(aw/3)/2) / 2
	ah := aw/2 + t
	var pts [][2]float64
	pts = append(pts, [2]float64{-aw / 2, t}, [2]float64{-aw/2 + t, 0})
	ncx, ncy := 0.0, ah-2*t-2*math.Sqrt2*t
	for a := 135.0; a >= 45; a -= 9 {
		rad := a * math.Pi / 180
		pts = append(pts, [2]float64{ncx + 2*t*math.Cos(rad), ncy + 2*t*math.Sin(rad)})
	}
	pts = append(pts, [2]float64{aw/2 - t, 0}, [2]float64{aw / 2, t}, [2]float64{0, ah})
	angle := 0.0
	switch dir {
	case "UP":
		angle = math.Pi
	case "LEFT":
		angle = math.Pi / 2
	case "RIGHT":
		angle = -math.Pi / 2
	}
	sin, cos := math.Sin(angle), math.Cos(angle)
	for i, q := range pts {
		qx, qy := q[0], q[1]-ah/2
		pts[i] = [2]float64{cx + qx*cos - qy*sin, cy + qx*sin + qy*cos}
	}
	fillPolygon(dst, pts, c)
}

func clScaleTrough(p clPalette, dst *image.NRGBA, x, y, w, h float64, lower bool) {
	ty := y + math.Floor(h/2) - 3
	th := 7.0
	clInset(dst, x, ty, w, th, 0, [4]bool{}, p.parent)
	fillCol, border, in := shade(p.parent, 0.896), p.shade[6], true
	if lower {
		fillCol, border, in = p.spot[1], p.spot[2], false
	}
	a, b := 0.95, 1.05
	if !in {
		a, b = 1.1, 0.9
	}
	fillRR(dst, x+1, ty+1, w-2, th-2, 0, allCorners, vgrad(ty+0.5, ty+th+1, []float64{0, 1}, []RGBA{shade(fillCol, a), shade(fillCol, b)}))
	strokeRR(dst, x+1.5, ty+1.5, w-3, th-3, 0, allCorners, 1, solid(mix(border, fillCol, 0.2)), nil)
}

func clSlider(s *Style, p clPalette, dst *image.NRGBA, x, y, w, h float64, st int) {
	r0 := math.Min(p.radius, math.Min((w-1)/2, (h-1)/2))
	shadow := withAlpha(shade(p.shade[6], 0.92), 0.1)
	strokeRR(dst, x-0.5, y-0.5, w, h, r0, allCorners, 1, solid(shadow), func(px, py float64) bool {
		return px > x+w-1 || py > y+h-1
	})
	bx, by, W, H := x+1, y+1, w-2, h-2
	f := s.Bg[st]
	if st == Prelight {
		f = shade(f, 1.04)
	}
	offs, cols := gummy(f, 1.08, 1.02, 0.94)
	fillRR(dst, bx+1, by+1, W-2, H-2, 0, allCorners, vgrad(by+1, by+H-2, offs, cols))
	border := mix(p.shade[7], f, 0.2)
	if st == Prelight {
		border = p.spot[2]
	}
	strokeRR(dst, bx+0.5, by+0.5, W-1, H-1, 2.5, allCorners, 1, solid(border), nil)
	sEven := 0
	if int(W)%2 == 0 {
		sEven = 1
	}
	for i := 0; i < 3-sEven; i++ {
		lx := bx + math.Floor(W/2) - 3 + float64(sEven) + float64(3*i)
		vline(dst, int(lx), int(by+4), int(by+H-5), p.shade[7])
	}
	clTLH(dst, bx+1, by+1, W-2, H-2, 2, allCorners, f, 0.4)
}

func clProgressTrough(p clPalette, dst *image.NRGBA, x, y, w, h float64) {
	r := math.Min(p.radius, math.Min((h-2)/2, (w-2)/2))
	fill(dst, image.Rect(int(x), int(y), int(x+w), int(y+h)), p.parent)
	fillRR(dst, x+1, y+1, w-2, h-2, r, allCorners, solid(p.shade[2]))
	strokeRR(dst, x+0.5, y+0.5, w-1, h-1, r, allCorners, 1, solid(mix(p.shade[7], p.shade[2], 0.3)), nil)
	sc := shade(p.shade[7], 0.92)
	clip := rrSDF(x+1, y+1, w-2, h-2, r, allCorners)
	in := func(px, py float64) bool { return clip(px, py) <= 0 }
	fillRegionFunc(dst, image.Rect(int(x+1), int(y+1), int(x+w-1), int(y+5)), func(px, py float64) RGBA {
		return withAlpha(sc, 0.3*(1-(py-y)/4))
	}, in)
	fillRegionFunc(dst, image.Rect(int(x+1), int(y+1), int(x+5), int(y+h-1)), func(px, py float64) RGBA {
		return withAlpha(sc, 0.3*(1-(px-x)/4))
	}, in)
}

func clProgressBar(s *Style, p clPalette, dst *image.NRGBA, x, y, w, h float64) {
	r := math.Min(math.Max(0, p.radius-float64(s.XThickness)), h/2)
	shape := func(px, py float64) bool {
		return rrSDF(x, y, w+r, h, r, clCorners(true, false, false, true))(px, py) <= 0 &&
			rrSDF(x-r, y, w+r, h, r, clCorners(false, true, true, false))(px, py) <= 0
	}
	bounds := image.Rect(int(x), int(y), int(x+w), int(y+h))
	offs, cols := gummy(p.spot[1], 1.08, 1.02, 0.94)
	fillRegionFunc(dst, bounds, vgrad(y, y+h, offs, cols), shape)
	P := 2 * h
	fillRegion(dst, bounds, withAlpha(p.spot[2], 0.15), func(px, py float64) bool {
		if !shape(px, py) {
			return false
		}
		lx, ly := px-x, py-y
		v := math.Mod(lx+ly*(P/2)/h, P)
		return v >= P/2
	})
	strokeRR(dst, x+0.5, y+0.5, w-2, h-1, r, allCorners, 1, solid(withAlpha(p.spot[0], 0.2)), nil)
	vline(dst, int(x+w-1), int(y), int(y+h-1), withAlpha(p.spot[2], 0.6))
	vline(dst, int(x+w), int(y), int(y+h-1), withAlpha(shade(p.shade[7], 0.92), 0.2))
}

func clFrame(s *Style, p clPalette, dst *image.NRGBA, x, y, w, h float64, shadow string, corners [4]bool, border RGBA) {
	r := math.Min(p.radius, math.Min((w-2)/2, (h-2)/2))
	hilight := shade(s.Bg[Normal], 1.05)
	switch shadow {
	case "NONE", "":
		return
	case "ETCHED_IN":
		strokeRR(dst, x+1.5, y+1.5, w-2, h-2, r, corners, 1, solid(hilight), nil)
		strokeRR(dst, x+0.5, y+0.5, w-2, h-2, r, corners, 1, solid(p.shade[4]), nil)
		return
	case "ETCHED_OUT":
		strokeRR(dst, x+0.5, y+0.5, w-2, h-2, r, corners, 1, solid(hilight), nil)
		strokeRR(dst, x+1.5, y+1.5, w-2, h-2, r, corners, 1, solid(p.shade[4]), nil)
		return
	}
	hl, sh := shade(s.Bg[Normal], 1.06), shade(s.Bg[Normal], 0.94)
	if shadow == "IN" {
		hl, sh = sh, hl
	}
	strokeRR(dst, x+1.5, y+1.5, w-3, h-3, r, corners, 1, solid(hl), func(px, py float64) bool { return px-x < py-y || px < x+2.5 || py < y+2.5 })
	strokeRR(dst, x+1.5, y+1.5, w-3, h-3, r, corners, 1, solid(sh), func(px, py float64) bool { return px > x+w-2.5 || py > y+h-2.5 })
	strokeRR(dst, x+0.5, y+0.5, w-1, h-1, r, corners, 1, solid(border), nil)
}

func clTab(s *Style, p clPalette, dst *image.NRGBA, r image.Rectangle, c Call) {
	st := c.State
	x, y, w, h := float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy())
	rad := math.Min(p.radius, math.Min((w-2)/2, (h-2)/2))
	sz := 2.0
	if s.YThickness == 3 {
		sz = 3
	}
	// gap BOTTOM (tabs above the page): lengthen downwards, round the top corners
	H := h + 3
	corners := clCorners(true, true, false, false)
	fillCol := s.Bg[st]
	ox, oy := x+0.5, y+0.5
	clip := image.NewNRGBA(dst.Bounds())
	fillRR(clip, ox, oy, w-1, H-1, rad, corners, solid(fillCol))
	L := H
	if st != Normal { // unselected
		offs := []float64{0, 1 / L, 1 / L, 0.45, 0.45, 1}
		cols := []RGBA{shade(fillCol, 1.15), shade(fillCol, 1.15), shade(fillCol, 1.08), shade(fillCol, 1.02), fillCol, shade(fillCol, 0.94)}
		fillRR(clip, ox, oy, w-1, H-1, rad, corners, vgrad(oy+1, oy+H, offs, cols))
		strokeRR(clip, ox, oy, w-1, H-1, rad, corners, 1, solid(p.shade[5]), nil)
	} else {
		hl := withAlpha(shade(s.Bg[Normal], 1.2), 0.5)
		strokeRR(clip, ox+1, oy+1, w-1, H-1, rad, corners, 1, solid(hl), func(px, py float64) bool { return px < x+2.5 || py < y+2.5 })
		sf, sb := sz/(L-2), (sz+1)/(L-2)
		offs := []float64{0, sf, sf, sb, sb, 0.8}
		cols := []RGBA{withAlpha(p.spot[1], 0.6), p.spot[1], withAlpha(p.spot[2], 0.72), withAlpha(p.spot[2], 0.72), shade(fillCol, 1.08), withAlpha(fillCol, 0)}
		fillRR(clip, ox, oy, w-1, H-1, rad, corners, vgrad(oy, oy+H, offs, cols))
		strokeRR(clip, ox, oy, w-1, H-1, rad, corners, 1, vgrad(oy+2, oy+H, []float64{0, 0.8}, []RGBA{p.spot[2], p.shade[5]}), nil)
	}
	blit(dst, clip, r)
}
