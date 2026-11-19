package gtkrc

import (
	"image"
	"math"
	"strconv"
	"strings"
)

// Murrine: glazed gradients, soft 1px corners, light inner borders.
// Implemented from a behavioural specification of the Murrine engine
// (0.98), not from its source. Translucent (rgba), animation and the blurred
// glow/menu styles are not implemented.

func init() { engines["murrine"] = murrine }

type murrineOpts struct {
	contrast                         float64
	glazestyle, roundness            int
	reliefstyle, focusstyle          int
	lightborderstyle, arrowstyle     int
	menustyle, menuitemstyle         int
	progressbarstyle, separatorstyle int
	sliderstyle, menubarstyle        int
	gradient                         [4]float64
	highlight, lightborder, glow     float64
	focusColor                       *RGBA
	shade                            [9]RGBA
	spot                             [3]RGBA
	parent                           RGBA
}

func contrastShade(s, c float64) float64 {
	switch {
	case c == 1:
		return s
	case c < 1:
		if s < 1 {
			return s + (1-s)*(1-c)
		}
		return s - (s-1)*(1-c)
	default:
		if s < 1 {
			return s - s*(c-1)
		}
		return s + (s-1)*(c-1)
	}
}

func decShade(s, f float64) float64 {
	switch {
	case s > 1:
		return 1 + (s-1)/f
	case s < 1:
		return 1 - (1-s)/f
	}
	return s
}

func murrineOptions(s *Style) murrineOpts {
	o := murrineOpts{contrast: 1, glazestyle: 1, roundness: 1, reliefstyle: 2, focusstyle: 2,
		menustyle: 1, menuitemstyle: 1, progressbarstyle: 1,
		gradient: [4]float64{1.1, 1, 1, 1.1}, highlight: 1.1, lightborder: 1.1, glow: 1}
	num := func(name string, def float64) float64 {
		if v, ok := s.EngineParams[name]; ok {
			if f, err := strconv.ParseFloat(strings.TrimSpace(unquote(v)), 64); err == nil {
				return f
			}
		}
		return def
	}
	integer := func(name string, def int) int { return int(num(name, float64(def))) }
	o.contrast = num("contrast", 1)
	o.glazestyle = integer("glazestyle", 1)
	o.roundness = integer("roundness", 1)
	o.reliefstyle = integer("reliefstyle", 2)
	o.focusstyle = integer("focusstyle", 2)
	o.lightborderstyle = integer("lightborderstyle", 0)
	o.arrowstyle = integer("arrowstyle", 0)
	o.menustyle = integer("menustyle", 1)
	o.menuitemstyle = integer("menuitemstyle", 1)
	o.progressbarstyle = integer("progressbarstyle", 1)
	o.separatorstyle = integer("separatorstyle", 0)
	o.sliderstyle = integer("sliderstyle", 0)
	o.menubarstyle = integer("menubarstyle", 0)
	o.highlight = num("highlight_shade", num("highlight_ratio", 1.1))
	o.lightborder = contrastShade(num("lightborder_shade", num("lightborder_ratio", 1.1)), o.contrast)
	o.glow = num("glow_shade", 1)
	if v, ok := s.EngineParams["gradient_shades"]; ok {
		parts := strings.Split(strings.Trim(v, "{}"), ",")
		for i := 0; i < 4 && i < len(parts); i++ {
			if f, err := strconv.ParseFloat(strings.TrimSpace(parts[i]), 64); err == nil {
				o.gradient[i] = f
			}
		}
	}
	if o.glazestyle != 0 {
		avg := (o.gradient[1] + o.gradient[2]) / 2
		o.gradient[1], o.gradient[2] = avg, avg
	}
	if o.glazestyle == 2 && o.roundness > 1 {
		o.roundness = 1
	}
	if v, ok := s.EngineParams["focus_color"]; ok {
		if c, ok := (&Theme{Scheme: s.scheme}).color(v, s); ok {
			o.focusColor = &c
		}
	}
	for i, f := range []float64{1.065, 0.95, 0.896, 0.82, 0.75, 0.665, 0.5, 0.45, 0.4} {
		o.shade[i] = shade(s.Bg[Normal], contrastShade(f, o.contrast))
	}
	o.spot = [3]RGBA{shade(s.Bg[Selected], 1.42), s.Bg[Selected], shade(s.Bg[Selected], contrastShade(0.65, o.contrast))}
	o.parent = s.Bg[Normal]
	return o
}

// openRect strokes Murrine's "open" rounded rectangle: plain for r < 1,
// soft corners (edges stopping 0.35px short) for r == 1, arcs for r >= 2.
func murrineOpenRect(dst *image.NRGBA, x, y, w, h float64, r int, corners [4]bool, col func(px, py float64) RGBA) {
	switch {
	case r < 1 || corners == [4]bool{}:
		strokeRR(dst, x, y, w, h, 0, corners, 1, col, nil)
	case r == 1:
		d := func(c bool) float64 {
			if c {
				return 0.35
			}
			return -0.5
		}
		seg := func(x0, y0, x1, y1 float64) {
			c := col((x0+x1)/2, (y0+y1)/2)
			segment(dst, x0, y0, x1, y1, 1, c, false)
		}
		seg(x+d(corners[0]), y, x+w-d(corners[1]), y)
		seg(x+w, y+d(corners[1]), x+w, y+h-d(corners[2]))
		seg(x+w-d(corners[2]), y+h, x+d(corners[3]), y+h)
		seg(x, y+h-d(corners[3]), x, y+d(corners[0]))
	default:
		rr := math.Floor(math.Min(float64(r), math.Min(w/2, h/2)))
		strokeRR(dst, x, y, w, h, rr, corners, 1, col, nil)
	}
}

// closedRect is the fill/clip shape: plain for r < 2, arcs otherwise.
func murrineClosed(x, y, w, h float64, r int, corners [4]bool) func(px, py float64) bool {
	rad := 0.0
	if r >= 2 {
		rad = math.Floor(math.Min(float64(r), math.Min(w/2, h/2)))
	}
	d := rrSDF(x, y, w, h, rad, corners)
	return func(px, py float64) bool { return d(px, py) <= 0 }
}

func (o murrineOpts) baseGradient(c RGBA, g [4]float64, by, bh, alpha float64) func(px, py float64) RGBA {
	cols := []RGBA{shade(c, g[0]), shade(c, g[1]), shade(c, g[2]), shade(c, g[3])}
	for i := range cols {
		cols[i].A = alpha
	}
	return vgrad(by, by+bh, []float64{0, 0.49, 0.49, 1}, cols)
}

// glaze is the GLAZE routine: base gradient, highlight shape, light border,
// clipped to the interior.
func (o murrineOpts) glaze(dst *image.NRGBA, fillCol RGBA, bx, by, bw, bh float64, r int, corners [4]bool, g [4]float64, highlight, lightborder float64, clip func(px, py float64) bool) {
	bounds := image.Rect(int(bx)-1, int(by)-1, int(bx+bw)+2, int(by+bh)+2)
	fillRegionFunc(dst, bounds, o.baseGradient(fillCol, g, by, bh, 1), clip)
	if o.glazestyle != 0 || highlight != 1 {
		W, H := math.Floor(bw+2*bx), math.Floor(bh+2*by)
		hh, h5 := math.Floor(H/2), math.Floor(H/5)
		var shape [][2]float64
		switch o.glazestyle {
		case 0:
			shape = [][2]float64{{bx, by}, {bx + bw, by}, {bx + bw, by + math.Floor(bh/2)}, {bx, by + math.Floor(bh/2)}}
		case 3, 4:
			shape = append([][2]float64{{bx, by}}, bezier([2]float64{bx, by}, [2]float64{bx, hh - h5}, [2]float64{h5, hh}, [2]float64{hh, hh}, 10)...)
			shape = append(shape, [2]float64{W - hh, hh})
			shape = append(shape, bezier([2]float64{W - hh, hh}, [2]float64{W - bx - h5, hh}, [2]float64{W - bx - 0.5, hh - h5}, [2]float64{W - bx, by}, 10)...)
		default: // 1 (curved), also used by 2 and 5
			shape = append([][2]float64{{bx, H - by}}, bezier([2]float64{bx, H - by}, [2]float64{bx, hh + h5}, [2]float64{h5, hh}, [2]float64{hh, hh}, 10)...)
			shape = append(shape, [2]float64{W - hh, hh})
			shape = append(shape, bezier([2]float64{W - hh, hh}, [2]float64{W - bx - h5, hh}, [2]float64{W - bx - 0.5, hh - h5}, [2]float64{W - bx, by}, 10)...)
			shape = append(shape, [2]float64{bx, by})
		}
		in := inPolygon(shape)
		if o.glazestyle == 2 {
			inner := in
			in = func(px, py float64) bool { return !inner(px, py) }
		}
		fillRegionFunc(dst, bounds, o.baseGradient(shade(fillCol, highlight), g, by, bh, 0.8), func(px, py float64) bool {
			return clip(px, py) && in(px, py)
		})
	}
	if o.glazestyle != 4 && lightborder != 1 {
		o.lightBorder(dst, fillCol, bx+0.5, by+0.5, bw-1, bh-1, r, corners, g, lightborder*highlight)
	}
}

func (o murrineOpts) lightBorder(dst *image.NRGBA, fillCol RGBA, lx, ly, lw, lh float64, r int, corners [4]bool, g [4]float64, f float64) {
	rad := math.Min(float64(r), math.Min(lw/2, lh/2))
	fp := 1 - 1/lh
	if corners == allCorners && rad > 2 {
		fp = 1 - ((rad-1)/2)/lh
	}
	a := [4]float64{0.5, 0.5, 0.5, 0}
	switch o.lightborderstyle {
	case 1:
		a = [4]float64{0.5, 0.5, 0.5, 0.5}
	case 2:
		a = [4]float64{0.5, 0.3, 0.1, 0}
	case 3:
		a = [4]float64{0.5, 0.3, 0.1, 0.1}
	}
	l := [4]RGBA{}
	for i := range l {
		l[i] = shade(fillCol, f*g[i])
	}
	col := vgrad(ly, ly+lh, []float64{0, 0.49, 0.49, fp, fp, 1}, []RGBA{
		withAlpha(l[0], a[0]), withAlpha(l[1], a[1]), withAlpha(l[2], a[1]), withAlpha(l[3], a[2]), withAlpha(l[3], a[3]), withAlpha(l[3], a[3])})
	sr := 0.0
	if rad >= 2 {
		sr = rad - 1
	}
	strokeRR(dst, lx, ly, lw, lh, sr, corners, 1, col, nil)
}

// inset is INSET: a light bottom-right and dark top-left ring.
func (o murrineOpts) inset(dst *image.NRGBA, x, y, w, h float64, r int, corners [4]bool) {
	hl, sh := withAlpha(shade(o.parent, 1.3), 0.42), withAlpha(shade(o.parent, 0.74), 0.16)
	rad := math.Min(float64(r), math.Min(w/2, h/2))
	lower := func(px, py float64) bool {
		return (px-x)+(py-y) > math.Min(w, h) || px > x+w-rad*0.3 || py > y+h-rad*0.3
	}
	strokeRR(dst, x, y, w, h, rad, corners, 1, solid(hl), lower)
	strokeRR(dst, x, y, w, h, rad, corners, 1, solid(sh), func(px, py float64) bool { return !lower(px, py) })
}

func (o murrineOpts) button(s *Style, dst *image.NRGBA, x, y, w, h float64, st int, active bool, corners [4]bool, relief int) {
	fillCol := s.Bg[st]
	border0 := shade(fillCol, contrastShade(0.475, o.contrast))
	g, hl, lb := o.gradient, o.highlight, o.lightborder
	if st == Insensitive {
		for i := range g {
			g[i] = decShade(g[i], 3)
		}
		hl, lb = decShade(hl, 2), decShade(lb, 2)
		border0 = shade(fillCol, contrastShade(0.75, o.contrast))
	}
	os := 0.0
	if s.XThickness > 2 && s.YThickness > 2 {
		os = 1
	}
	if os == 1 {
		if !active && st != Insensitive && relief > 1 {
			murrineOpenRect(dst, x+os-0.5, y+os-0.5, w-2*os+1, h-2*os+1, o.roundness+1, corners, solid(withAlpha(border0, 0.08)))
		} else if relief != 0 {
			o.inset(dst, x+0.5, y+0.5, w-1, h-1, o.roundness+1, corners)
		}
	}
	border := mix(border0, fillCol, 0.4)
	ix, iy, iw, ih := x+os+1, y+os+1, w-2*os-2, h-2*os-2
	clip := murrineClosed(ix, iy, iw, ih, o.roundness-1, corners)
	light := lb
	if active {
		light = 1
	}
	o.glaze(dst, fillCol, ix, iy, iw, ih, o.roundness-1, corners, g, hl, light, clip)
	if active {
		sc := shade(fillCol, 0.94)
		fillRegionFunc(dst, image.Rect(int(ix), int(iy), int(ix+iw), int(iy+3)), func(px, py float64) RGBA {
			return withAlpha(sc, 0.58*(1-(py-iy)/3))
		}, clip)
		fillRegionFunc(dst, image.Rect(int(ix), int(iy), int(ix+3), int(iy+ih)), func(px, py float64) RGBA {
			return withAlpha(sc, 0.58*(1-(px-ix)/3))
		}, clip)
	}
	murrineOpenRect(dst, x+os+0.5, y+os+0.5, w-2*os-1, h-2*os-1, o.roundness, corners, solid(border))
}

func murrine(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) bool {
	o := murrineOptions(s)
	st := c.State
	x, y, w, h := float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy())
	switch c.Function {
	case "BOX", "SHADOW":
		switch c.Detail {
		case "button", "optionmenu":
			o.button(s, dst, x, y, w, h, st, c.Shadow == "IN", allCorners, o.reliefstyle)
			return true
		case "buttondefault":
			return true
		case "spinbutton_up", "spinbutton_down":
			corners := [4]bool{false, true, false, false}
			if c.Detail == "spinbutton_down" {
				corners = [4]bool{false, false, true, false}
			}
			o.button(s, dst, x-1, y, w+1, h, st, c.Shadow == "IN", corners, min(o.reliefstyle, 1))
			return true
		case "entry":
			o.entry(s, dst, x, y, w, h, st, c.Focused)
			return true
		case "trough", "trough-lower", "trough-upper":
			if c.Widget == "GtkProgressBar" {
				o.progressTrough(s, dst, x, y, w, h)
			} else {
				o.scaleTrough(s, dst, x, y, w, h, st, c.Detail == "trough-lower")
			}
			return true
		case "bar":
			o.progressBar(s, dst, x, y, w, h)
			return true
		case "menu":
			fill(dst, r, s.Bg[Normal])
			if o.menustyle == 1 {
				fill(dst, image.Rect(r.Min.X+1, r.Min.Y+1, r.Min.X+3, r.Max.Y-1), o.spot[1])
				vline(dst, r.Min.X+3, r.Min.Y+1, r.Max.Y-2, mix(o.spot[1], shade(o.spot[1], 0.5), 0.5))
			}
			ring(dst, r, o.shade[5])
			return true
		case "menuitem":
			if o.menustyle == 1 {
				x, w = x+3, w-3
			}
			clip := murrineClosed(x, y, w, h, o.roundness, allCorners)
			if o.menuitemstyle == 1 {
				o.glaze(dst, o.spot[1], x+1, y+1, w-2, h-2, o.roundness, allCorners, o.gradient, o.highlight, o.lightborder, clip)
			} else {
				fillRegionFunc(dst, r, o.baseGradient(o.spot[1], o.gradient, y, h, 1), clip)
			}
			murrineOpenRect(dst, x+0.5, y+0.5, w-1, h-1, o.roundness, allCorners, solid(withAlpha(o.spot[2], 0.8)))
			return true
		case "menubar":
			if c.Shadow == "NONE" {
				return true
			}
			switch o.menubarstyle {
			case 2:
				fillRR(dst, x, y, w, h, 0, allCorners, vgrad(y, y+h, []float64{0, 1}, []RGBA{s.Bg[Normal], shade(s.Bg[Normal], 0.95)}))
			default:
				fill(dst, r, s.Bg[Normal])
			}
			hline(dst, r.Min.X, r.Max.X-1, r.Max.Y-1, o.shade[3])
			return true
		case "scrolled_window", "viewport", "":
			if c.Function == "SHADOW" {
				ring(dst, r, o.shade[5])
				return true
			}
		}
		if c.Function == "SHADOW" {
			o.frame(s, dst, x, y, w, h, c.Shadow, o.shade[4], [4]bool{})
			return true
		}
		return false
	case "SLIDER":
		o.button(s, dst, x, y, w, h, st, false, allCorners, o.reliefstyle)
		if o.sliderstyle == 1 || o.sliderstyle == 3 {
			col := mix(shade(o.shade[6], 0.95), s.Bg[st], 0.4)
			n := 2
			if int(w)%2 == 1 {
				n = 3
			}
			bar := int(w)/2 - n
			for i := 0; i < n; i++ {
				vline(dst, r.Min.X+bar+3*i, r.Min.Y+5, r.Max.Y-6, col)
			}
		}
		return true
	case "CHECK", "OPTION":
		o.check(s, dst, x, y, w, h, st, c.Shadow, c.Function == "OPTION", c.Detail == "check" || c.Detail == "option")
		return true
	case "ARROW", "TAB":
		col := mix(s.Fg[st], s.Bg[st], 0.2)
		cx, cy := x+w/2, y+h/2
		double := c.Function == "TAB" || c.Widget == "GtkComboBox"
		if double { // a fixed 5×5 box
			cx, cy = math.Floor(x+w/2)-1+2.5, math.Floor(y+h/2)-2+2.5
		}
		draw := func(dx, dy float64, col RGBA) {
			if double {
				aw := math.Min(2*5.0/3, 5)
				clChevron(dst, cx+dx, cy-aw/2+dy, aw, "UP", col)
				clChevron(dst, cx+dx, cy+aw/2+dy, aw, "DOWN", col)
				return
			}
			clChevron(dst, cx+dx, cy+dy, math.Min(w, 7), c.ArrowDirection, col)
		}
		if st == Insensitive {
			draw(0.5, 0.5, o.shade[0])
		}
		draw(0, 0, col)
		return true
	case "HLINE", "VLINE":
		dark := withAlpha(shade(s.Bg[Normal], contrastShade(0.7, o.contrast)), 0.5)
		light := withAlpha(shade(s.Bg[Normal], contrastShade(1.3, o.contrast)), 0.5)
		if c.Function == "HLINE" {
			yy := r.Min.Y + r.Dy()/2 - 1
			hline(dst, r.Min.X, r.Max.X-1, yy, dark)
			if o.separatorstyle != 3 {
				hline(dst, r.Min.X, r.Max.X-2, yy+1, light)
			}
		} else {
			xx := r.Min.X + r.Dx()/2 - 1
			vline(dst, xx, r.Min.Y, r.Max.Y-1, dark)
			if o.separatorstyle != 3 {
				vline(dst, xx+1, r.Min.Y, r.Max.Y-1, light)
			}
		}
		return true
	case "BOX_GAP":
		corners := allCorners
		if o.roundness < 2 {
			corners = [4]bool{}
		}
		fillRegion(dst, r, s.Bg[Normal], murrineClosed(x, y, w, h, o.roundness, corners))
		o.frame(s, dst, x, y, w, h, c.Shadow, o.shade[5], corners)
		return true
	case "EXTENSION":
		o.tab(s, dst, r, c)
		return true
	case "FOCUS":
		return true // entries draw their own focus; buttons show none in nativeui
	}
	return false
}

func (o murrineOpts) entry(s *Style, dst *image.NRGBA, x, y, w, h float64, st int, focused bool) {
	border := o.shade[6]
	if st == Insensitive {
		border = o.shade[4]
	}
	if focused {
		border = o.spot[2]
		if o.focusColor != nil {
			border = *o.focusColor
		}
	}
	rr := max(0, min(o.roundness, 3))
	fill(dst, image.Rect(int(x), int(y), int(x+w), int(y+h)), o.parent)
	fill(dst, image.Rect(int(x)+2, int(y)+2, int(x+w)-2, int(y+h)-2), s.Base[st])
	if o.reliefstyle != 0 {
		o.inset(dst, x+0.5, y+0.5, w-1, h-1, rr+1, allCorners)
	}
	if focused {
		strokeRR(dst, x+2.5, y+2.5, w-5, h-5, 0, allCorners, 1, solid(withAlpha(shade(border, 1.54), 0.5)), nil)
	} else {
		a := 0.15
		if st == Insensitive {
			a = 0.05
		}
		col := withAlpha(shade(border, 0.925), a)
		vline(dst, int(x)+2, int(y)+2, int(y+h)-3, col)
		hline(dst, int(x)+3, int(x+w)-3, int(y)+2, col)
	}
	murrineOpenRect(dst, x+1.5, y+1.5, w-3, h-3, rr, allCorners, solid(border))
}

func (o murrineOpts) check(s *Style, dst *image.NRGBA, x, y, w, h float64, st int, shadow string, radio, inMenu bool) {
	checked := shadow == "IN" || shadow == "ETCHED_IN"
	border, bg, dot := o.shade[5], s.Base[Normal], s.Text[st]
	if checked {
		border, bg = o.spot[2], o.spot[1]
	}
	g, hl, lb := o.gradient, o.highlight, o.lightborder
	if st == Insensitive {
		border, bg, dot = o.shade[3], s.Bg[Normal], o.shade[3]
		for i := range g {
			g[i] = decShade(g[i], 3)
		}
		hl, lb = decShade(hl, 2), decShade(lb, 2)
	}
	radius := max(0, min(o.roundness, 2))
	if radio {
		radius = int(w + h)
	}
	if !inMenu {
		var clip func(px, py float64) bool
		if radio {
			cx, cy, rr := x+w/2, y+h/2, (w-3)/2
			clip = func(px, py float64) bool { return math.Hypot(px-cx, py-cy) <= rr }
		} else {
			clip = murrineClosed(x+1.5, y+1.5, w-3, h-3, radius, allCorners)
		}
		if checked {
			o.glaze(dst, bg, x+2, y+2, w-4, h-4, radius, allCorners, g, hl, lb, clip)
		} else {
			fillRegion(dst, image.Rect(int(x), int(y), int(x+w), int(y+h)), bg, clip)
		}
		if radio {
			strokeCircle(dst, x+w/2, y+h/2, (w-3)/2, 1, border)
		} else {
			murrineOpenRect(dst, x+1.5, y+1.5, w-3, h-3, radius, allCorners, solid(border))
		}
	}
	if shadow != "IN" {
		if shadow == "ETCHED_IN" && radio {
			segment(dst, x+5, y+h/2, x+w-5, y+h/2, 2, dot, false)
		}
		return
	}
	invert := func(c RGBA) RGBA {
		hh, l, ss := rgbToHLS(c.R, c.G, c.B)
		if l < 0.8 {
			l = 1
		} else {
			l = 0
		}
		r, g, b := hlsToRGB(hh, l, ss)
		return RGBA{r, g, b, 1}
	}
	if radio {
		cx, cy := x+w/2, y+h/2
		if inMenu {
			fillDisc(dst, cx, cy, (w+h)/4-4, dot, nil)
			return
		}
		a := 0.3
		if st == Insensitive {
			a *= 0.2
		}
		fillDisc(dst, cx, cy, (w+h)/4-4, withAlpha(invert(dot), a), nil)
		fillDisc(dst, cx, cy, (w+h)/4-5, dot, nil)
		return
	}
	sx, sy := w/18, h/18
	p := func(px, py float64) [2]float64 { return [2]float64{x + px*sx, y + py*sy} }
	ox, oy := 2.0, 3.0
	if !inMenu {
		a := 0.5
		if st == Insensitive {
			a *= 0.2
		}
		halo := [][2]float64{p(5, 5.65), p(8.95, 9.57), p(16, 2.54), p(16, 8.36), p(10.6, 15.1), p(7.6, 15.1), p(2.95, 10.48), p(2.95, 7.65)}
		fillPolygon(dst, halo, withAlpha(invert(dot), a))
		ox, oy = 4, 2
	}
	var tick [][2]float64
	for _, q := range [][2]float64{{0, 6}, {0, 8}, {4, 12}, {6, 12}, {15, 1}, {15, 0}, {14, 0}, {5, 9}, {1, 5}} {
		tick = append(tick, p(q[0]+ox, q[1]+oy))
	}
	fillPolygon(dst, tick, dot)
}

func (o murrineOpts) scaleTrough(s *Style, dst *image.NRGBA, x, y, w, h float64, st int, lower bool) {
	by := y + math.Floor(h/2) - 3
	lx, ly, W := x+0.5, by+0.5, w
	if o.reliefstyle != 0 {
		o.inset(dst, lx, ly, W-1, 5, o.roundness, allCorners)
	}
	fillCol, border := s.Bg[Active], shade(s.Bg[Active], contrastShade(0.82, o.contrast))
	if lower {
		fillCol = s.Bg[Selected]
		if st == Insensitive {
			fillCol = mix(s.Bg[Selected], o.parent, 0.25)
		}
		border = shade(fillCol, contrastShade(0.65, o.contrast))
	}
	fillRegion(dst, image.Rect(int(x), int(by), int(x+w), int(by+6)), fillCol, murrineClosed(lx+1, ly+1, W-2, 4, o.roundness, allCorners))
	murrineOpenRect(dst, lx+1, ly+1, W-3, 3, o.roundness, allCorners, solid(border))
}

func (o murrineOpts) progressTrough(s *Style, dst *image.NRGBA, x, y, w, h float64) {
	fill(dst, image.Rect(int(x), int(y), int(x+w), int(y+h)), o.parent)
	rr := int(math.Min(float64(o.roundness), math.Min((h-2)/2, (w-2)/2)))
	fillCol, border := s.Bg[Active], shade(s.Bg[Active], contrastShade(0.82, o.contrast))
	clip := murrineClosed(x+1, y+1, w-2, h-2, rr-1, allCorners)
	fillRegion(dst, image.Rect(int(x), int(y), int(x+w), int(y+h)), fillCol, clip)
	murrineOpenRect(dst, x+0.5, y+0.5, w-1, h-1, rr, allCorners, solid(border))
	sc := shade(border, 0.94)
	fillRegionFunc(dst, image.Rect(int(x+1), int(y+1), int(x+w-1), int(y+5)), func(px, py float64) RGBA {
		return withAlpha(sc, 0.24*(1-(py-y)/4))
	}, clip)
	fillRegionFunc(dst, image.Rect(int(x+1), int(y+1), int(x+5), int(y+h-1)), func(px, py float64) RGBA {
		return withAlpha(sc, 0.24*(1-(px-x)/4))
	}, clip)
}

func (o murrineOpts) progressBar(s *Style, dst *image.NRGBA, x, y, w, h float64) {
	X, W, H := x-1, w+2, h
	fillCol := o.spot[1]
	effect := shade(fillCol, contrastShade(0.65, o.contrast))
	border := mix(o.spot[2], fillCol, 0.28)
	rr := int(math.Min(float64(o.roundness-s.XThickness), H/2))
	clip := murrineClosed(X+2, y+1, W-4, H-2, rr-1, allCorners)
	o.glaze(dst, fillCol, X+2, y+1, W-4, H-2, rr, allCorners, o.gradient, o.highlight, o.lightborder, clip)
	if o.progressbarstyle == 1 {
		P := 2 * H
		fillRegion(dst, image.Rect(int(X), int(y), int(X+W), int(y+H)), withAlpha(effect, 0.15), func(px, py float64) bool {
			if !clip(px, py) {
				return false
			}
			lx, ly := px-X, py-y
			return math.Mod(lx+ly*(P/2)/H, P) >= P/2
		})
	}
	murrineOpenRect(dst, X+1.5, y+0.5, W-3, H-1, rr, allCorners, solid(border))
}

func (o murrineOpts) frame(s *Style, dst *image.NRGBA, x, y, w, h float64, shadow string, border RGBA, corners [4]bool) {
	light, darkBevel := shade(s.Bg[Normal], 1.04), shade(s.Bg[Normal], 0.96)
	lx, ly := x+0.5, y+0.5
	switch shadow {
	case "NONE", "":
		return
	case "ETCHED_IN":
		murrineOpenRect(dst, lx+1, ly+1, w-2, h-2, o.roundness, corners, solid(light))
		murrineOpenRect(dst, lx, ly, w-2, h-2, o.roundness, corners, solid(o.shade[3]))
		return
	case "ETCHED_OUT":
		murrineOpenRect(dst, lx, ly, w-2, h-2, o.roundness, corners, solid(light))
		murrineOpenRect(dst, lx+1, ly+1, w-2, h-2, o.roundness, corners, solid(o.shade[3]))
		return
	}
	tl, br := light, darkBevel
	if shadow == "IN" {
		tl, br = br, tl
	}
	segment(dst, lx+1, ly+h-2, lx+1, ly+1, 1, tl, false)
	segment(dst, lx+1, ly+1, lx+w-2, ly+1, 1, tl, false)
	segment(dst, lx+w-2, ly+1, lx+w-2, ly+h-2, 1, br, false)
	segment(dst, lx+w-2, ly+h-2, lx+1, ly+h-2, 1, br, false)
	murrineOpenRect(dst, lx, ly, w-1, h-1, o.roundness, corners, solid(border))
}

func (o murrineOpts) tab(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) {
	st := c.State
	bg := s.Bg[st]
	x, y, w, h := float64(r.Min.X)+0.5, float64(r.Min.Y)+0.5, float64(r.Dx()), float64(r.Dy())+3
	corners := [4]bool{true, true, false, false}
	var g [4]float64
	for i := range g {
		g[i] = decShade(o.gradient[i], 3)
	}
	hl := decShade(o.highlight, 2)
	clipImg := image.NewNRGBA(dst.Bounds())
	shape := murrineClosed(x, y, w-1, h-1, o.roundness, corners)
	bounds := image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Max.Y+3)
	var fillStops []RGBA
	var lbStops []RGBA
	if st == Normal {
		fillStops = []RGBA{shade(bg, g[0]*hl), shade(bg, g[1]*hl), shade(bg, g[2]), bg}
		hc := shade(bg, o.lightborder*hl)
		lbStops = []RGBA{withAlpha(shade(hc, g[0]*hl), 0.5), withAlpha(shade(hc, g[1]*hl), 0.5), withAlpha(shade(hc, g[2]), 0.5), shade(bg, 1.04)}
	} else {
		fillStops = []RGBA{shade(bg, g[0]*hl), shade(bg, g[1]*hl), shade(bg, g[2]), shade(bg, g[3])}
		hc := shade(bg, decShade(o.lightborder, 2)*hl)
		lbStops = []RGBA{withAlpha(shade(hc, g[0]*hl), 0.5), withAlpha(shade(hc, g[1]*hl), 0.5), withAlpha(shade(hc, g[2]), 0.5), withAlpha(shade(hc, g[3]), 0.5)}
	}
	offs := []float64{0, 0.45, 0.45, 1}
	fillRegionFunc(clipImg, bounds, vgrad(y, y+h, offs, fillStops), shape)
	strokeRR(clipImg, x+1, y+1, w-3, h-3, 0, corners, 1, vgrad(y+1, y+h-2, offs, lbStops), nil)
	border := o.shade[5]
	if st != Normal {
		border = o.shade[4]
	}
	murrineOpenRect(clipImg, x, y, w-1, h-1, o.roundness, corners, solid(border))
	blit(dst, clipImg, r)
}
