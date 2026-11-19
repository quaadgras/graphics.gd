package gtkrc

import (
	"image"
	"math"
)

// Mist: flat 1px lines in GTK's derived light/dark colours, etched (all dark)
// outlines for most things. Implemented from a behavioural specification of
// gtk-engines' Mist, not from its source.

func init() { engines["mist"] = mist }

// mistShadow remaps the requested shadow type by detail.
func mistShadow(c Call) string {
	t := c.Shadow
	switch c.Detail {
	case "dockitem", "handlebox_bin":
		t = "NONE"
	case "spinbutton_up", "spinbutton_down":
		t = "OUT"
	case "button", "togglebutton", "notebook", "optionmenu":
	case "menu":
		t = "ETCHED_IN"
	default:
		if t != "NONE" && t != "" {
			t = "ETCHED_IN"
		}
	}
	if (c.Function == "SHADOW" || c.Function == "BOX") && c.State == Insensitive && t != "NONE" && t != "" {
		t = "ETCHED_IN"
	}
	return t
}

func mistBorder(s *Style, dst *image.NRGBA, r image.Rectangle, t string, st int) {
	switch t {
	case "ETCHED_IN", "ETCHED_OUT":
		ring(dst, r, s.dark(st))
	case "OUT":
		simpleBorder(dst, r, s.light(st), s.dark(st), false)
	case "IN":
		simpleBorder(dst, r, s.dark(st), s.light(st), false)
	}
}

func mist(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) bool {
	st := c.State
	x, y, w, h := r.Min.X, r.Min.Y, r.Dx(), r.Dy()
	switch c.Function {
	case "BOX", "SHADOW", "SLIDER":
		switch c.Detail {
		case "buttondefault":
			ring(dst, r, s.Fg[Normal])
			return true
		case "trough":
			fill(dst, image.Rect(x, y, x+w-1, y+h-1), s.Bg[st])
			ring(dst, r, s.dark(st))
			return true
		case "bar":
			if w > 1 && h > 1 {
				fill(dst, r.Inset(1), s.Base[Selected])
			}
			ring(dst, r, s.dark(Selected))
			return true
		case "menubar", "toolbar", "dockitem", "dockitem_bin", "handlebox":
			if c.Shadow != "NONE" {
				hline(dst, x, x+w-1, y+h-1, s.dark(Normal))
			}
			return true
		case "menuitem":
			st = Selected
			fill(dst, r, s.Bg[st])
			if c.Shadow != "NONE" {
				ring(dst, r, s.dark(st))
			}
			return true
		}
		if c.Function != "SHADOW" {
			fill(dst, r, s.Bg[st])
		}
		mistBorder(s, dst, r, mistShadow(c), st)
		if c.Function == "SLIDER" && (c.Detail == "hscale" || c.Detail == "vscale") {
			ty := s.YThickness
			if c.Orientation == "VERTICAL" {
				hline(dst, x+s.XThickness, x+w-s.XThickness-1, y+h/2, s.dark(st))
			} else {
				vline(dst, x+w/2, y+ty, y+h-ty-1, s.dark(st))
			}
		}
		return true
	case "FOCUS":
		ring(dst, r, s.dark(st))
		return true
	case "CHECK":
		base := s.Base[Normal]
		if st == Insensitive {
			base = s.Base[Insensitive]
		}
		fill(dst, image.Rect(x+1, y+1, x+w-2, y+h-2), base)
		ring(dst, image.Rect(x, y, x+w-1, y+h-1), s.dark(st))
		inner := image.Rect(x+2, y+2, x+w-3, y+h-3)
		switch c.Shadow {
		case "IN":
			mark := s.Base[Selected]
			if st == Insensitive {
				mark = s.dark(Insensitive)
			}
			fill(dst, inner, mark)
		case "ETCHED_IN":
			for py := inner.Min.Y; py < inner.Max.Y; py++ {
				for px := inner.Min.X; px < inner.Max.X; px++ {
					if (px+py)%2 == 1 {
						setPixel(dst, px, py, s.Base[Selected])
					}
				}
			}
		}
		return true
	case "OPTION":
		cx, cy := float64(x+w/2), float64(y+h/2)
		R := float64(w/2) - 0.5
		rim, base, dot := s.dark(Normal), s.Base[Normal], s.Base[Selected]
		if st == Insensitive {
			rim, base, dot = s.dark(Insensitive), s.Base[Insensitive], s.dark(Insensitive)
		}
		strokeCircle(dst, cx, cy, R, 1, rim)
		fillDisc(dst, cx, cy, R, base, nil)
		strokeCircle(dst, cx, cy, R, 1, RGBA{0, 0, 0, 0.35})
		switch c.Shadow {
		case "IN":
			fillDisc(dst, cx, cy, float64((w-7)/2+1), dot, nil)
		case "ETCHED_IN":
			t := (h + 1) / 3
			t -= t % 2
			odd := 0.0
			if h%2 == 1 {
				odd = 0.5
			}
			lw := float64(t) + 2*odd
			yy := float64(y+h/2) + 0.5
			x0, x1 := float64(x+t)-odd, float64(x+w-t)+odd
			segment(dst, x0, yy, x1, yy, lw, s.Bg[Selected], false)
			fillDisc(dst, x0, yy, lw/2, s.Bg[Selected], nil) // round caps
			fillDisc(dst, x1, yy, lw/2, s.Bg[Selected], nil)
		}
		return true
	case "TAB":
		if c.Detail == "optionmenutab" {
			vline(dst, x-5, y, y+h, s.dark(st))
			arrow(dst, image.Rect(x+1, y+1, x+w-1, y+h-1), "DOWN", s.Fg[st])
			return true
		}
		return false
	case "BOX_GAP":
		fill(dst, r, s.Bg[st])
		t := c.Shadow
		if c.Detail != "notebook" {
			t = mistShadow(c)
		}
		switch t {
		case "IN":
			simpleBorder(dst, r, s.dark(st), s.light(st), false)
		case "OUT":
			simpleBorder(dst, r, s.light(st), s.dark(st), false)
		case "ETCHED_IN", "ETCHED_OUT":
			ring(dst, r, s.dark(st))
		}
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
		yy := y + h/2 - 1
		switch c.Detail {
		case "hseparator", "menuitem", "slider", "vscale":
			hline(dst, x, x+w-2, yy, s.dark(st))
			return true
		}
		T := s.YThickness
		D, L := T-T/2, T/2
		x1, x2 := x, x+w-1
		for i := 0; i < D; i++ {
			hline(dst, x1, x2-i-1, yy+i, s.dark(st))
			hline(dst, x2-i, x2, yy+i, s.light(st))
		}
		for i := 0; i < L; i++ {
			hline(dst, x1, x1+L-i-2, yy+D+i, s.dark(st))
			hline(dst, x1+L-i-1, x2, yy+D+i, s.light(st))
		}
		return true
	case "VLINE":
		xx := x + w/2 - 1
		vline(dst, xx, y, y+h-2, s.dark(st))
		return true
	}
	_ = math.Pi
	return false
}
