package gtkrc

import (
	"image"
	"math"
	"strings"
)

// ThinIce: flat 1px bevels, corner-chopped slider with "//" grips, check
// boxes without a tick, radio buttons as a two-tone ring. Implemented from a
// behavioural specification of gtk-engines' ThinIce, not from its source.

func init() { engines["thinice"] = thinice }

func thiniceOpt(s *Style, name, def string) string {
	if v, ok := s.EngineParams[name]; ok {
		return strings.ToUpper(unquote(v))
	}
	return def
}

// thiniceBevel is gtk_paint_shadow (and the border of boxes).
func thiniceBevel(s *Style, dst *image.NRGBA, r image.Rectangle, st int, shadow, detail string) {
	switch shadow {
	case "IN":
		if (detail == "entry" || detail == "text") && (r.Min.X == 1 || r.Min.Y == 1) {
			simpleBorder(dst, r, s.Base[st], s.Base[st], false)
			return
		}
		simpleBorder(dst, r, s.dark(st), s.light(st), false)
	case "OUT":
		simpleBorder(dst, r, s.light(st), s.dark(st), false)
	case "ETCHED_IN", "ETCHED_OUT":
		first, second := s.light(st), s.dark(st)
		if shadow == "ETCHED_OUT" {
			first, second = second, first
		}
		inner := image.Rect(r.Min.X, r.Min.Y, r.Max.X-2, r.Max.Y-2)
		ring(dst, inner.Add(image.Pt(1, 1)), first)
		ring(dst, inner, second)
	}
}

func thinice(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) bool {
	st := c.State
	switch c.Function {
	case "SHADOW":
		thiniceBevel(s, dst, r, st, c.Shadow, c.Detail)
		return true
	case "BOX", "TAB":
		switch c.Detail {
		case "optionmenutab", "slider", "buttondefault":
			return true
		case "bar":
			fill(dst, r.Inset(1), s.Bg[st])
			return true
		case "handlebox_bin", "entry-progress":
			fill(dst, r, s.Bg[st])
			return true
		}
		fill(dst, r, s.Bg[st])
		thiniceBevel(s, dst, r, st, c.Shadow, c.Detail)
		return true
	case "CHECK":
		if st == Insensitive {
			ring(dst, r, s.dark(Insensitive))
			return true
		}
		fill(dst, r, s.Bg[st])
		thiniceBevel(s, dst, r, st, c.Shadow, c.Detail)
		if c.Shadow == "IN" {
			fill(dst, r.Inset(1), s.Bg[Active])
		}
		return true
	case "OPTION":
		thiniceRadio(s, dst, r, c)
		return true
	case "ARROW":
		if c.Detail == "vscrollbar" || c.Detail == "hscrollbar" {
			return false // mark_type2 steppers: not needed by nativeui
		}
		v := gtkArrow(r, c.ArrowDirection)
		if st == Insensitive {
			var off [3][2]int
			for i := range v {
				off[i] = [2]int{v[i][0] + 1, v[i][1] + 1}
			}
			drawArrowShape(dst, off, RGBA{1, 1, 1, 1})
		}
		drawArrowShape(dst, v, s.Fg[st])
		return true
	case "SLIDER":
		thiniceSlider(s, dst, r, c)
		return true
	case "BOX_GAP":
		fill(dst, r, s.Bg[st])
		var tl, br RGBA
		switch c.Shadow {
		case "IN":
			tl, br = s.dark(st), s.light(st)
		case "OUT":
			tl, br = s.light(st), s.dark(st)
		case "ETCHED_IN", "ETCHED_OUT":
			tl, br = s.dark(st), s.dark(st)
		default:
			return true
		}
		simpleBorder(dst, r, tl, br, false) // gap: nativeui draws the page without one
		return true
	case "EXTENSION":
		fill(dst, r, s.Bg[st])
		ext := r
		switch c.GapSide {
		case "TOP":
			ext.Min.Y--
		case "LEFT":
			ext.Min.X--
		case "BOTTOM":
			ext.Max.Y++
		case "RIGHT":
			ext.Max.X++
		}
		clipped := image.NewNRGBA(dst.Bounds())
		simpleBorder(clipped, ext, s.light(st), s.dark(st), false)
		blit(dst, clipped, r)
		return true
	case "HLINE":
		y := r.Min.Y + r.Dy()/2 - 1
		hline(dst, r.Min.X, r.Max.X-1, y, s.dark(st))
		hline(dst, r.Min.X, r.Max.X-1, y+1, s.light(st))
		return true
	case "VLINE":
		x := r.Min.X + r.Dx()/2 - 1
		vline(dst, x, r.Min.Y, r.Max.Y-1, s.dark(st))
		vline(dst, x+1, r.Min.Y, r.Max.Y-1, s.light(st))
		return true
	}
	return false
}

// blit copies src over dst within r (used to clip drawing to a rectangle).
func blit(dst, src *image.NRGBA, r image.Rectangle) {
	r = r.Intersect(dst.Bounds())
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if c := src.NRGBAAt(x, y); c.A > 0 {
				blend(dst, x, y, RGBA{float64(c.R) / 255, float64(c.G) / 255, float64(c.B) / 255, float64(c.A) / 255}, 1)
			}
		}
	}
}

func thiniceRadio(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) {
	st := c.State
	x, y, w, h := r.Min.X, r.Min.Y, r.Dx(), r.Dy()
	cx, cy := float64(x+w/2), float64(y+h/2)
	R := float64(min(w, h)/2 - 1)
	a, b, centre := s.light(st), s.dark(st), s.Bg[st]
	if c.Shadow == "IN" || c.Shadow == "ETCHED_IN" {
		a, b, centre = s.dark(st), s.light(st), s.Bg[Active]
	}
	switch c.Shadow {
	case "ETCHED_IN":
		strokeCircle(dst, cx+1, cy+1, R+1, 0.5, b)
		strokeCircle(dst, cx, cy, R+1, 0.5, a)
		return
	case "ETCHED_OUT":
		strokeCircle(dst, cx-1, cy-1, R+1, 0.5, a)
		strokeCircle(dst, cx, cy, R+1, 0.5, b)
		return
	}
	D := R + 2
	sn := D * 0.70711
	fillDisc(dst, cx, cy, R+1, b, inPolygon([][2]float64{{cx + D, cy + D}, {cx + sn, cy - sn}, {cx - sn, cy + sn}}))
	fillDisc(dst, cx, cy, R+1, a, inPolygon([][2]float64{{cx - D, cy - D}, {cx + sn, cy - sn}, {cx - sn, cy + sn}}))
	fillDisc(dst, cx, cy, R, centre, nil)
}

func thiniceSlider(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) {
	st := c.State
	x, y, w, h := r.Min.X, r.Min.Y, r.Dx(), r.Dy()
	horizontal := c.Orientation != "VERTICAL"
	rect := thiniceOpt(s, "rect_scrollbar", "FALSE") == "TRUE"
	marks := true
	if !rect && w <= 17 && h <= 17 {
		rect, marks = true, false
	}
	if rect {
		fill(dst, r, s.Bg[st])
		thiniceBevel(s, dst, r, st, c.Shadow, c.Detail)
	} else {
		cut := w - 17
		if !horizontal {
			cut = h - 17
		}
		cf := math.Max(0, math.Min(6, float64(cut)))
		fx, fy, fw, fh := float64(x), float64(y), float64(w), float64(h)
		p := [][2]float64{{fx, fy + fh - 1}, {fx, fy + cf}, {fx + cf, fy}, {fx + fw - 1, fy}, {fx + fw - 1, fy + fh - 1 - cf}, {fx + fw - 1 - cf, fy + fh - 1}}
		fillPolygon(dst, p, s.Bg[st])
		c5 := func(i int) [2]float64 { return [2]float64{p[i][0] + 0.5, p[i][1] + 0.5} }
		polyline(dst, 1, s.light(st), c5(0), c5(1), c5(2), c5(3))
		polyline(dst, 1, s.dark(st), c5(3), c5(4), c5(5), c5(0))
	}
	if !marks {
		return
	}
	first, second := s.light(st), s.dark(st)
	switch thiniceOpt(s, "mark_type1", "SLASH") {
	case "NOTHING":
		return
	case "INVSLASH", "INVDOT":
		first, second = second, first
	}
	switch thiniceOpt(s, "mark_type1", "SLASH") {
	case "DOT", "INVDOT":
		for _, o := range []int{-4, 0, 4} {
			X, Y := x+w/2, y+h/2
			if horizontal {
				X += o
			} else {
				Y += o
			}
			setPixel(dst, X+1, Y+1, first)
			setPixel(dst, X+1, Y, first)
			setPixel(dst, X, Y+1, first)
			setPixel(dst, X-1, Y-1, second)
			setPixel(dst, X-1, Y, second)
			setPixel(dst, X, Y-1, second)
		}
	default: // two slashes
		ox, oy := float64(x+w/2)+0.5, float64(y+h/2)+0.5
		t := float64((min(w, h)-1)/2) - 1.5
		for _, o := range []float64{-2, 2} {
			for k, col := range []RGBA{second, first} {
				dx, dy := 0.0, 0.0
				if w > h {
					dx = o + float64(k)
				} else {
					dy = o + float64(k)
				}
				segment(dst, ox-t+dx, oy+t+dy, ox+t+dx, oy-t+dy, 1, col, true)
			}
		}
	}
}
