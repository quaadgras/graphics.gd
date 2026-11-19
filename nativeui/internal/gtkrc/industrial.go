package gtkrc

import (
	"image"
	"math"
	"strconv"
	"strings"
)

// Industrial: flat fills with a 1px outline in the foreground colour at a
// partial opacity, plus soft gradient shading. Implemented from a behavioural
// specification of gtk-engines' Industrial, not from its source.

func init() { engines["industrial"] = industrial }

type industrialOpts struct {
	contrast float64
	rounded  bool
}

func industrialOptions(s *Style) industrialOpts {
	o := industrialOpts{contrast: 1, rounded: true}
	if v, ok := s.EngineParams["contrast"]; ok {
		if f, err := strconv.ParseFloat(unquote(v), 64); err == nil {
			o.contrast = f
		}
	}
	if v, ok := s.EngineParams["rounded_buttons"]; ok && strings.EqualFold(unquote(v), "FALSE") {
		o.rounded = false
	}
	return o
}

func (o industrialOpts) a(v float64) float64 { return clamp01(v * o.contrast) }

func (o industrialOpts) corners() [4]bool {
	if o.rounded {
		return allCorners
	}
	return [4]bool{}
}

// industrialOutline is the OUTLINE primitive: an optional square fill inset
// by 1px and a 1px rounded stroke, opaque (pre-composited) when filled.
func industrialOutline(dst *image.NRGBA, r image.Rectangle, rad float64, border RGBA, fillCol *RGBA, corners [4]bool) {
	stroke := border
	if fillCol != nil {
		stroke = over(border, *fillCol)
		fill(dst, r.Inset(1), *fillCol)
	}
	if corners == ([4]bool{}) {
		rad = 0
	}
	strokeRoundRect(dst, r, rad, corners, stroke)
}

// industrialSoftFrame paints a band of depth g inside r, outer at the edge
// fading to inner at depth g (sides only: the corners are small enough at
// these sizes for the side gradients to cover them).
func industrialSoftFrame(dst *image.NRGBA, r image.Rectangle, g float64, inner, outer RGBA, clip func(px, py float64) bool) {
	fx, fy, fw, fh := float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy())
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			px, py := float64(x)+0.5, float64(y)+0.5
			if clip != nil && !clip(px, py) {
				continue
			}
			d := math.Min(math.Min(px-fx, fx+fw-px), math.Min(py-fy, fy+fh-py))
			if d >= g {
				continue
			}
			t := d / g
			c := RGBA{outer.R, outer.G, outer.B, outer.A + (inner.A-outer.A)*t}
			blend(dst, x, y, c, 1)
		}
	}
}

func industrialGrid(dst *image.NRGBA, r image.Rectangle, c RGBA) {
	painted := map[[2]int]float64{}
	row := 0
	for yy := r.Min.Y + 1; yy < r.Max.Y; yy += 2 {
		start := r.Min.X + 1
		if row%2 == 1 {
			start = r.Min.X + 3
		}
		for xx := start; xx < r.Max.X; xx += 4 {
			painted[[2]int{xx, yy}] = c.A
			for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				p := [2]int{xx + d[0], yy + d[1]}
				if painted[p] < c.A/2 {
					painted[p] = c.A / 2
				}
			}
		}
		row++
	}
	for p, a := range painted {
		blend(dst, p[0], p[1], withAlpha(c, a), 1)
	}
}

func industrial(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) bool {
	o := industrialOptions(s)
	st := c.State
	x, y, w, h := r.Min.X, r.Min.Y, r.Dx(), r.Dy()
	border := withAlpha(s.Fg[st], o.a(0.38))
	bg := s.Bg[st]
	switch c.Function {
	case "HLINE":
		yy := y + h/2
		if c.Detail == "menuitem" {
			yy++
		}
		hline(dst, x, x+w-1, yy, withAlpha(s.Fg[st], o.a(0.40)))
		return true
	case "VLINE":
		vline(dst, x+w/2, y, y+h-1, withAlpha(s.Fg[st], o.a(0.40)))
		return true
	case "FOCUS":
		if c.Detail != "entry" {
			return false
		}
		rad := 1.5
		if !o.rounded {
			rad = 0
		}
		strokeRoundRect(dst, r, rad, allCorners, s.Bg[Selected])
		strokeRoundRect(dst, r.Inset(1), math.Max(0, rad-1), allCorners, s.Bg[Selected])
		return true
	case "CHECK":
		m, f := s.Fg[st], s.Bg[st]
		if st == Normal {
			m, f = s.Text[Normal], s.Base[Normal]
		}
		S := min(w, h)
		sq := image.Rect(x+(w-S)/2, y+(h-S)/2, x+(w-S)/2+S, y+(h-S)/2+S)
		industrialOutline(dst, sq, 1.5, withAlpha(m, o.a(0.26)), &f, o.corners())
		fx, fy := float64(sq.Min.X), float64(sq.Min.Y)
		half := inPolygon([][2]float64{{fx, fy + float64(S)}, {fx + float64(S), fy}, {fx, fy}})
		hi := image.NewNRGBA(dst.Bounds())
		strokeRoundRect(hi, sq.Inset(1), 1.5, o.corners(), withAlpha(m, o.a(0.26)*0.3))
		for yy := sq.Min.Y; yy < sq.Max.Y; yy++ {
			for xx := sq.Min.X; xx < sq.Max.X; xx++ {
				if half(float64(xx)+0.5, float64(yy)+0.5) {
					if p := hi.NRGBAAt(xx, yy); p.A > 0 {
						blend(dst, xx, yy, RGBA{float64(p.R) / 255, float64(p.G) / 255, float64(p.B) / 255, float64(p.A) / 255}, 1)
					}
				}
			}
		}
		k := float64(S-4) / 7
		ox, oy := float64(sq.Min.X+2), float64(sq.Min.Y+2)
		u := func(px, py float64) [2]float64 { return [2]float64{ox + px*k, oy + py*k} }
		mark := withAlpha(m, o.a(1))
		switch c.Shadow {
		case "IN":
			pts := [][2]float64{u(0, 4.1), u(2.8, 6.65)}
			pts = append(pts, bezier(u(2.8, 6.65), u(3.7, 5.2), u(5.0, 2.65), u(6.9, 1.4), 12)...)
			pts = append(pts, u(6.4, 0.6))
			pts = append(pts, bezier(u(6.4, 0.6), u(4.5, 1.9), u(3.55, 3.4), u(2.6, 4.7), 12)...)
			pts = append(pts, u(0.8, 3.0))
			fillPolygon(dst, pts, mark)
		case "ETCHED_IN":
			a, b := u(0, 3.5), u(7, 3.5)
			segment(dst, a[0], a[1], b[0], b[1], 2*k, mark, false)
		}
		return true
	case "OPTION":
		m, f := s.Fg[st], s.Bg[st]
		if st == Normal {
			m, f = s.Text[Normal], s.Base[Normal]
		}
		rr := float64(min(w, h)) / 2
		cx, cy := float64(x)+float64(w)/2, float64(y)+float64(h)/2
		fillDisc(dst, cx, cy, rr-0.5, f, nil)
		fillRegion(dst, r.Inset(-1), withAlpha(m, o.a(0.5)), func(px, py float64) bool {
			return math.Hypot(px-cx, py-cy) <= rr && math.Hypot(px-cx-0.2, py-cy-0.2) > rr-1.2
		})
		switch c.Shadow {
		case "IN":
			fillDisc(dst, cx, cy, rr-3, withAlpha(m, o.a(1)), nil)
			hr := rr - 4
			if hr > 0 {
				for yy := int(cy - hr - 1); yy <= int(cy+hr+1); yy++ {
					for xx := int(cx - hr - 1); xx <= int(cx+hr+1); xx++ {
						px, py := float64(xx)+0.5, float64(yy)+0.5
						d := math.Hypot(px-cx, py-cy)
						if d <= hr && (px-cx)+(py-cy) < 0 {
							blend(dst, xx, yy, withAlpha(f, o.a(0.7)*d/hr), 1)
						}
					}
				}
			}
		case "ETCHED_IN":
			segment(dst, cx-rr+2, cy, cx+rr-2, cy, 2, withAlpha(m, o.a(0.5)), false)
		}
		return true
	case "EXTENSION":
		if c.Shadow == "NONE" {
			return true
		}
		ext := r
		switch c.GapSide {
		case "TOP":
			ext.Min.Y -= 3
		case "BOTTOM":
			ext.Max.Y += 3
		case "LEFT":
			ext.Min.X -= 3
		case "RIGHT":
			ext.Max.X += 3
		}
		clipped := image.NewNRGBA(dst.Bounds())
		industrialOutline(clipped, ext, 1.5, border, &bg, o.corners())
		if st != Normal {
			shade := withAlpha(s.Fg[st], o.a(0.38)*0.3)
			for i := 0; i < 4; i++ {
				a := shade.A * (1 - float64(i)/4)
				col := withAlpha(s.Fg[st], a)
				switch c.GapSide {
				case "BOTTOM":
					hline(clipped, x, x+w-1, y+h-1-i, col)
				case "TOP":
					hline(clipped, x, x+w-1, y+i, col)
				}
			}
		}
		blit(dst, clipped, r)
		return true
	case "BOX_GAP":
		if c.Shadow == "NONE" {
			return true
		}
		industrialOutline(dst, r, 1.5, border, &bg, o.corners())
		return true
	case "SLIDER":
		industrialOutline(dst, r, 1.5, border, &bg, o.corners())
		gw, gh := min(w-2, 19), min(h-2, 7)
		if c.Orientation == "VERTICAL" {
			gw, gh = min(w-2, 7), min(h-2, 19)
		}
		gx, gy := x+(w-gw)/2, y+(h-gh)/2
		industrialGrid(dst, image.Rect(gx, gy, gx+gw, gy+gh), withAlpha(s.Fg[st], o.a(0.38)))
		return true
	case "BOX", "SHADOW":
		switch c.Detail {
		case "button", "togglebutton", "optionmenu":
			industrialOutline(dst, r, 2.5, border, &bg, o.corners())
			sz := float64(min(w/2, h/2))
			fx, fy, fw, fh := float64(x), float64(y), float64(w), float64(h)
			if c.Shadow == "IN" {
				clip := inPolygon([][2]float64{{fx, fy + fh}, {fx + sz, fy + fh - sz}, {fx + fw - sz, fy + sz}, {fx + fw, fy}, {fx, fy}})
				industrialSoftFrame(dst, r.Inset(1), 3, withAlpha(s.Fg[st], 0), withAlpha(s.Fg[st], o.a(0.266)), clip)
			}
			if w > 18 && h > 18 && c.Shadow != "NONE" {
				rr := r
				if c.Shadow == "IN" {
					rr = rr.Add(image.Pt(1, 1))
				}
				fx, fy = float64(rr.Min.X), float64(rr.Min.Y)
				clip := inPolygon([][2]float64{{fx, fy + fh}, {fx + sz, fy + fh - sz}, {fx + fw - sz, fy + sz}, {fx + fw, fy}, {fx + fw, fy + fh}})
				industrialSoftFrame(dst, rr.Inset(1), 4, withAlpha(s.Fg[st], 0), withAlpha(s.Fg[st], o.a(0.10)), clip)
			}
			return true
		case "entry":
			fill(dst, r, s.Bg[Normal])
			base := s.Base[st]
			rad := 1.5
			industrialOutline(dst, r, rad, border, &base, o.corners())
			return true
		case "spinbutton_up", "spinbutton_down":
			in := image.Rect(x+1, y+1, x+w-1, y+h)
			if c.Detail == "spinbutton_down" {
				in = image.Rect(x+1, y, x+w-1, y+h-1)
			}
			fill(dst, in, bg)
			return true
		case "spinbutton":
			industrialOutline(dst, r, 1.5, border, &bg, [4]bool{false, o.rounded, o.rounded, false})
			return true
		case "trough", "trough-lower", "trough-upper", "trough-fill-level", "trough-fill-level-full":
			if c.Function == "SHADOW" {
				return true
			}
			fill(dst, r, s.Bg[Normal])
			ts := st
			if ts == Active {
				ts = Normal
			}
			interior := over(withAlpha(s.Fg[ts], o.a(0.15)), s.Bg[Normal])
			if c.Detail == "trough-lower" || strings.HasPrefix(c.Detail, "trough-fill") {
				interior = s.Bg[Selected]
			}
			industrialOutline(dst, r, 1.5, withAlpha(s.Fg[ts], o.a(0.21)), &interior, o.corners())
			return true
		case "menu", "toolbar", "dockitem", "dockitem_bin", "handlebox_bin":
			if c.Function == "SHADOW" {
				return true
			}
			industrialOutline(dst, r, 1.5, withAlpha(s.Fg[st], o.a(0.21)), &bg, o.corners())
			return true
		case "menuitem":
			b := withAlpha(s.Fg[Normal], o.a(0.38))
			if c.Function == "SHADOW" {
				industrialOutline(dst, r, 1.5, b, nil, o.corners())
			} else {
				industrialOutline(dst, r, 1.5, b, &bg, o.corners())
			}
			return true
		}
		if c.Shadow == "NONE" {
			if c.Function == "BOX" {
				fill(dst, r, bg)
			}
			return true
		}
		if c.Function == "SHADOW" {
			industrialOutline(dst, r, 1.5, border, nil, o.corners())
		} else {
			industrialOutline(dst, r, 1.5, border, &bg, o.corners())
		}
		return true
	}
	return false
}
