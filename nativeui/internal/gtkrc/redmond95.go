package gtkrc

import (
	"image"
	"math"
)

// Redmond95: the Windows 95 look, 1px bevels in four tones. Implemented from
// a behavioural specification of gtk-engines' Redmond95, not from its source.

func init() { engines["redmond95"] = redmond95 }

func (s *Style) veryDark(st int) RGBA { return shade(s.Bg[st], 0.3) }

// redmondRing is RING(tl, br) with corner ownership.
func redmondRing(dst *image.NRGBA, r image.Rectangle, tl, br RGBA, tlWins bool) {
	simpleBorder(dst, r, tl, br, tlWins)
}

func hatch(dst *image.NRGBA, r image.Rectangle, a, b RGBA, in func(px, py float64) bool) {
	r = r.Intersect(dst.Bounds())
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if in != nil && !in(float64(x)+0.5, float64(y)+0.5) {
				continue
			}
			if (x+y)%2 == 0 {
				setPixel(dst, x, y, a)
			} else {
				setPixel(dst, x, y, b)
			}
		}
	}
}

// redmondShadow is B3, the border vocabulary.
func redmondShadow(s *Style, dst *image.NRGBA, r image.Rectangle, st int, shadow, detail string) {
	inner := r.Inset(1)
	light, dark, bg, black := s.light(st), s.dark(st), s.Bg[st], s.veryDark(st)
	switch shadow {
	case "ETCHED_IN":
		redmondRing(dst, r, dark, light, true)
		redmondRing(dst, inner, light, dark, true)
	case "ETCHED_OUT":
		redmondRing(dst, r, light, dark, true)
		redmondRing(dst, inner, dark, light, true)
	case "IN":
		switch detail {
		case "buttondefault":
			ring(dst, r, black)
		case "vscrollbar", "hscrollbar":
			ring(dst, r, dark)
		default:
			redmondRing(dst, r, dark, light, false)
			redmondRing(dst, inner, black, bg, false)
		}
	case "OUT":
		switch detail {
		case "vscrollbar", "hscrollbar", "slider", "spinbutton_up", "spinbutton_down", "menu", "combobutton":
			redmondRing(dst, r, bg, black, false)
			redmondRing(dst, inner, light, dark, false)
		case "frame":
			redmondRing(dst, r, light, dark, false)
		default:
			redmondRing(dst, r, light, black, false)
			redmondRing(dst, inner, bg, dark, false)
		}
	}
}

func redmondTick(dst *image.NRGBA, r image.Rectangle, c RGBA) {
	X, Y, W, H := r.Min.X, r.Min.Y, r.Dx(), r.Dy()
	scale := min(W, H)
	odd := scale % 2
	factor := 10 - odd
	if scale <= factor+2 {
		scale = factor
	}
	left := float64(X+(W-scale)/2) + 0.5
	top := float64(Y+(H-scale)/2) + 0.5
	u := func(n int) float64 { return float64(n * scale / factor) }
	pts := [][2]float64{
		{left + u(1), top + u(4-odd)}, {left + u(1), top + u(6-odd)}, {left + u(3), top + u(8-odd)},
		{left + u(8-odd), top + u(3)}, {left + u(8-odd), top + u(1)}, {left + u(3), top + u(6-odd)},
	}
	fillPolygon(dst, pts, c)
	for i := range pts {
		j := (i + 1) % len(pts)
		segment(dst, pts[i][0], pts[i][1], pts[j][0], pts[j][1], 0.5, c, false)
	}
}

func redmond95(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) bool {
	st := c.State
	x, y, w, h := r.Min.X, r.Min.Y, r.Dx(), r.Dy()
	switch c.Function {
	case "SHADOW":
		redmondShadow(s, dst, r, st, c.Shadow, c.Detail)
		return true
	case "BOX", "SLIDER":
		switch {
		case c.Detail == "trough" && c.Widget == "GtkHScale":
			fill(dst, r, s.Bg[Normal])
			yy := y + h/2
			hline(dst, x, x+w-3, yy, s.veryDark(st))
			redmondShadow(s, dst, image.Rect(x, yy-1, x+w, yy+3), st, "IN", "trough")
			return true
		case c.Detail == "bar" && c.Widget == "GtkProgressBar":
			fill(dst, r, s.Bg[Selected])
			return true
		case c.Detail == "spinbutton_up" || c.Detail == "spinbutton_down":
			if st != Insensitive {
				st = Normal
			}
			fill(dst, r, s.Bg[st])
			redmondShadow(s, dst, r, st, c.Shadow, c.Detail)
			return true
		}
		fill(dst, r, s.Bg[st])
		redmondShadow(s, dst, r, st, c.Shadow, c.Detail)
		return true
	case "CHECK":
		if c.Detail == "check" && c.Shadow != "ETCHED_IN" {
			if c.Shadow == "IN" {
				redmondTick(dst, image.Rect(x+2, y+2, x+w-2, y+h-2), s.Text[Normal])
			}
			return true
		}
		switch {
		case st == Active || st == Insensitive:
			fill(dst, image.Rect(x, y, x+w-1, y+h-1), s.Bg[Normal])
		case c.Shadow == "ETCHED_IN":
			hatch(dst, r, s.Bg[Normal], s.light(Normal), nil)
		default:
			fill(dst, image.Rect(x, y, x+w-1, y+h-1), s.Base[Normal])
		}
		tick := image.Rect(x+2, y+2, x+w-2, y+h-2)
		if c.Shadow == "ETCHED_IN" || st == Insensitive {
			redmondTick(dst, tick, s.Fg[Insensitive])
		} else if c.Shadow == "IN" {
			redmondTick(dst, tick, s.Fg[Normal])
		}
		redmondShadow(s, dst, r, Normal, "IN", c.Detail)
		return true
	case "OPTION":
		cx, cy := float64(x+w/2), float64(y+h/2)
		R := math.Max(float64(min(w, h)/2), 6)
		b := math.Max((R-2)*0.5, 1)
		upper := func(px, py float64) bool { return (px-cx)+(py-cy) < 0 }
		lower := func(px, py float64) bool { return !upper(px, py) }
		fillDisc(dst, cx, cy, R, s.dark(Normal), upper)
		fillDisc(dst, cx, cy, R, s.light(Normal), lower)
		fillDisc(dst, cx, cy, R-1, s.veryDark(Normal), upper)
		fillDisc(dst, cx, cy, R-1, s.Bg[Normal], lower)
		switch {
		case st == Active || st == Insensitive:
			fillDisc(dst, cx, cy, R-2, s.Bg[Normal], nil)
		case c.Shadow == "ETCHED_IN":
			hatch(dst, r, s.Bg[Normal], s.light(Normal), func(px, py float64) bool { return math.Hypot(px-cx, py-cy) <= R-2 })
		default:
			fillDisc(dst, cx, cy, R-2, s.Base[Normal], nil)
		}
		if c.Shadow == "ETCHED_IN" || st == Insensitive {
			fillDisc(dst, cx, cy, b, s.Fg[Insensitive], nil)
		} else if c.Shadow == "IN" {
			fillDisc(dst, cx, cy, b, s.Text[Normal], nil)
		}
		return true
	case "ARROW":
		if c.Detail != "menuitem" && st != Insensitive {
			st = Normal
		}
		v := gtkArrow(r, c.ArrowDirection)
		if st == Insensitive {
			var off [3][2]int
			for i := range v {
				off[i] = [2]int{v[i][0] + 1, v[i][1] + 1}
			}
			drawArrowShape(dst, off, s.light(Insensitive))
		}
		drawArrowShape(dst, v, s.Fg[st])
		return true
	case "TAB":
		return c.Detail == "optionmenutab"
	case "HLINE":
		yy := y + h/2 - 1
		hline(dst, x+1, x+w-2, yy, s.dark(st))
		hline(dst, x+1, x+w-2, yy+1, s.light(st))
		return true
	case "VLINE":
		xx := x + w/2 - 1
		vline(dst, xx, y+1, y+h-2, s.dark(st))
		vline(dst, xx+1, y+1, y+h-2, s.light(st))
		return true
	case "EXTENSION":
		redmondTab(s, dst, r, c)
		return true
	case "BOX_GAP":
		// stock GTK: bg fill and the two-pixel bevel, black outermost on the bottom/right
		fill(dst, r, s.Bg[st])
		redmondRing(dst, r, s.light(st), RGBA{0, 0, 0, 1}, false)
		redmondRing(dst, r.Inset(1), s.Bg[st], s.dark(st), false)
		return true
	}
	return false
}

func redmondTab(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) {
	st := c.State
	x, y, w, h := r.Min.X, r.Min.Y, r.Dx(), r.Dy()
	fill(dst, r, s.Bg[Normal])
	if c.Shadow == "NONE" {
		return
	}
	c1, c2, c3, c4 := s.light(st), s.Bg[st], s.dark(st), s.veryDark(st)
	switch c.Shadow {
	case "IN":
		c1, c2, c3, c4 = s.dark(st), s.veryDark(st), s.Bg[st], s.light(st)
	case "ETCHED_IN":
		c1, c2, c3, c4 = s.dark(st), s.light(st), s.dark(st), s.light(st)
	case "ETCHED_OUT":
		c1, c2, c3, c4 = s.light(st), s.dark(st), s.light(st), s.dark(st)
	}
	xt, yt := s.XThickness, s.YThickness
	switch c.GapSide {
	case "TOP":
		fill(dst, image.Rect(x+xt, y, x+w-xt, y+h-yt), s.Bg[st])
		vline(dst, x, y, y+h-2, c1)
		vline(dst, x+1, y, y+h-2, c2)
		hline(dst, x+2, x+w-2, y+h-2, c3)
		vline(dst, x+w-2, y, y+h-2, c3)
		hline(dst, x+1, x+w-2, y+h-1, c4)
		vline(dst, x+w-1, y, y+h-2, c4)
	default: // BOTTOM: tabs above the page
		fill(dst, image.Rect(x+xt, y+yt, x+w-xt, y+h), s.Bg[st])
		hline(dst, x+2, x+w-3, y, c1)
		vline(dst, x, y+2, y+h-1, c1)
		hline(dst, x+1, x+w-2, y+1, c2)
		vline(dst, x+1, y+1, y+h-1, c2)
		segment(dst, float64(x)+0.5, float64(y+2)+0.5, float64(x+1)+0.5, float64(y+1)+0.5, 1, c1, false)
		vline(dst, x+w-2, y+2, y+h-1, c3)
		vline(dst, x+w-1, y+2, y+h-1, c4)
		segment(dst, float64(x+w-1)+0.5, float64(y+2)+0.5, float64(x+w-2)+0.5, float64(y+1)+0.5, 1, c4, false)
	}
}
