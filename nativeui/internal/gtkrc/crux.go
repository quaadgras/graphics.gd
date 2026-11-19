package gtkrc

import (
	"image"
	"math"
)

// Crux: soft rounded gradient buttons with translucent bevels, a dark
// two-line entry frame, selection-coloured sliders and progress bars.
// Implemented from a behavioural specification of gtk-engines' Crux, not
// from its source.

func init() { engines["crux-engine"] = crux }

var (
	cruxHighlight = RGBA{1, 1, 1, 0.7}
	cruxShadow    = RGBA{0, 0, 0, 0.2}
	cruxEntryLine = RGBA{54.0 / 255, 52.0 / 255, 54.0 / 255, 1}
)

func cruxOutline(c RGBA) RGBA { return shade(c, 0.5) }

// cruxStdShadow is the standard shadow table.
func cruxStdShadow(dst *image.NRGBA, r image.Rectangle, bg RGBA, shadow string) {
	b := cruxOutline(bg)
	in := r.Inset(1)
	switch shadow {
	case "OUT":
		ring(dst, r, b)
		simpleBorder(dst, in, cruxHighlight, cruxShadow, true)
	case "IN":
		simpleBorder(dst, r, cruxShadow, cruxHighlight, true)
		ring(dst, in, b)
	case "ETCHED_IN":
		simpleBorder(dst, r, cruxShadow, cruxHighlight, true)
		simpleBorder(dst, in, cruxHighlight, cruxShadow, true)
	case "ETCHED_OUT":
		simpleBorder(dst, r, cruxHighlight, cruxShadow, true)
		simpleBorder(dst, in, cruxShadow, cruxHighlight, true)
	}
}

func crux(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) bool {
	st := c.State
	x, y, w, h := float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy())
	bg := s.Bg[st]
	switch c.Function {
	case "SHADOW":
		if c.Detail == "entry" {
			cruxEntry(s, dst, r, st, c.Focused)
			return true
		}
		cruxStdShadow(dst, r, bg, c.Shadow)
		return true
	case "BOX", "SLIDER":
		switch {
		case c.Detail == "button" || c.Detail == "optionmenu":
			cruxButton(s, dst, x, y, w, h, st, c.Shadow, true)
			return true
		case c.Detail == "spinbutton_up" || c.Detail == "spinbutton_down":
			if c.Detail == "spinbutton_up" {
				cruxButton(s, dst, x, y+3, w-3, h-3, st, c.Shadow, false)
			} else {
				cruxButton(s, dst, x, y+1, w-3, h-4, st, c.Shadow, false)
			}
			return true
		case c.Detail == "trough" && c.Widget == "GtkHScale":
			y, h = y+1, h-2
			c1, c2 := shade(s.Bg[Normal], 1.1), shade(s.Bg[Normal], 0.7)
			X, Y, W, H := x+0.5, y+0.5, w-1, h-1
			fillRR(dst, X, Y, W, H, 2, allCorners, vgrad(Y+2, Y+H-2, []float64{0, 1}, []RGBA{c1, c2}))
			strokeRR(dst, X, Y, W, H, 2, allCorners, 1, solid(cruxOutline(s.Bg[Normal])), nil)
			strokeRR(dst, X+1, Y+1, W-2, H-2, 2, allCorners, 1, vgrad(Y, Y+H, []float64{0, 1}, []RGBA{{1, 1, 1, 1}, {1, 1, 1, 100.0 / 255}}), nil)
			segment(dst, X+2, Y+H+1, X+W-2, Y+H+1, 1, shade(s.Bg[Normal], 0.8), true)
			return true
		case c.Function == "SLIDER" || (c.Detail == "hscale"):
			base := s.Bg[Selected]
			top, bot := shade(base, 1.5), shade(base, 0.9)
			if st == Prelight {
				top = shade(top, 1.2)
			}
			fillRR(dst, x+0.5, y+0.5, w-1, h-1, 2, allCorners, vgrad(y, y+h, []float64{0, 1}, []RGBA{top, bot}))
			strokeRR(dst, x+0.5, y+0.5, w-1, h-1, 2, allCorners, 1, solid(bot), nil)
			strokeRR(dst, x+1.5, y+1.5, w-3, h-3, 2, allCorners, 1, vgrad(y, y+h, []float64{0, 1}, []RGBA{{1, 1, 1, 0.2}, {1, 1, 1, 0.1}}), nil)
			return true
		case c.Detail == "bar":
			base := s.Base[Selected]
			c1, c2 := shade(base, 1.5), shade(base, 0.9)
			fillRR(dst, x, y, w, h, 0, allCorners, vgrad(y, y+h, []float64{0, 1}, []RGBA{c1, c2}))
			X, Y, W, H := x+0.5, y+0.5, w-2, h-1
			strokeRR(dst, X, Y, W, H, 0, allCorners, 1, vgrad(Y, Y+H, []float64{0, 1}, []RGBA{{1, 1, 1, 0.3}, {1, 1, 1, 0.1}}), nil)
			vline(dst, int(X+W+1), int(y), int(y+h-1), c2)
			for i, a := range []float64{0.22, 0.12, 0.03} {
				vline(dst, int(X+W+2)+i, int(y), int(y+h-1), RGBA{0, 0, 0, a})
			}
			return true
		case c.Detail == "menu":
			ring(dst, r, cruxEntryLine)
			in := r.Inset(1)
			hline(dst, in.Min.X, in.Max.X-1, in.Min.Y, s.light(st))
			vline(dst, in.Min.X, in.Min.Y, in.Max.Y-1, s.light(st))
			hline(dst, in.Min.X, in.Max.X-1, in.Max.Y-1, s.dark(st))
			vline(dst, in.Max.X-1, in.Min.Y, in.Max.Y-1, s.dark(st))
			return true
		case c.Detail == "menuitem":
			cruxStdShadow(dst, r, bg, "IN")
			l := shade(bg, 1.3)
			fillRR(dst, x+2, y+2, w-4, h-4, 0, allCorners, vgrad(y+2, y+h-2, []float64{0, 0.28, 0.35, 1}, []RGBA{bg, l, l, shade(l, 0.6)}))
			return true
		}
		fill(dst, r, bg)
		if c.Detail == "trough" && c.Widget == "GtkProgressBar" {
			cx, cy, cw, ch := x+1.5, y+1.5, w-3, h-3
			polyline(dst, 1, RGBA{0, 0, 0, 0.22}, [2]float64{cx + 1, cy + ch - 1}, [2]float64{cx + 1, cy + 1}, [2]float64{cx + cw - 1, cy + 1})
			polyline(dst, 1, RGBA{0, 0, 0, 0.12}, [2]float64{cx + 2, cy + ch - 1}, [2]float64{cx + 2, cy + 2}, [2]float64{cx + cw - 1, cy + 2}, [2]float64{cx + cw - 1, cy + ch - 1})
		}
		cruxStdShadow(dst, r, bg, c.Shadow)
		return true
	case "FOCUS":
		if c.Detail == "entry" {
			ring(dst, r, s.Base[Selected])
		}
		return c.Detail == "entry" || c.Detail == "button"
	case "CHECK":
		outline, dk := cruxOutline(s.Bg[Normal]), shade(s.Bg[Normal], 0.9)
		cx, cy, cw, ch := x+0.5, y+0.5, w-1, h-1
		menu := c.Detail == "check"
		if !menu {
			if st == Insensitive {
				fillRR(dst, cx, cy, cw, ch, 2, allCorners, solid(s.Bg[Insensitive]))
				strokeRR(dst, cx, cy, cw, ch, 2, allCorners, 1, solid(s.Fg[Insensitive]), nil)
			} else {
				a, b := RGBA{1, 1, 1, 1}, dk
				if st == Active {
					a, b = b, a
				}
				fillRR(dst, cx, cy, cw, ch, 2, allCorners, vgrad(cy, cy+ch, []float64{0, 1}, []RGBA{a, b}))
				strokeRR(dst, cx, cy, cw, ch, 2, allCorners, 1, solid(outline), nil)
				strokeRR(dst, cx+1, cy+1, cw-2, ch-2, 1, allCorners, 1, solid(cruxHighlight), nil)
			}
			cx, cy, cw, ch = cx+2, cy+2, cw-4, ch-4
		}
		switch c.Shadow {
		case "IN":
			pts := [][2]float64{{cx + math.Floor(cw/4), cy + ch/2}, {cx + math.Floor(cw/2), cy + math.Floor(0.75*ch)}, {cx + cw, cy + math.Floor(ch/4)}}
			if st == Insensitive {
				polyline(dst, 2, s.Fg[Insensitive], pts...)
			} else {
				shifted := [][2]float64{{pts[0][0], pts[0][1] + 1}, {pts[1][0], pts[1][1] + 1}, {pts[2][0], pts[2][1] + 1}}
				polyline(dst, 2, RGBA{0, 0, 0, 0.3}, shifted...)
				polyline(dst, 2, s.Bg[Selected], pts...)
			}
		case "ETCHED_IN":
			fillRegion(dst, r, RGBA{0, 0, 0, 0.3}, inPolygon([][2]float64{{x + float64(int(w)/4), y + float64(int(h)/3) + 0.5}, {x + w - float64(int(w)/4), y + float64(int(h)/3) + 0.5}, {x + w - float64(int(w)/4), y + float64(int(h)/3+int(h)/4) + 0.5}, {x + float64(int(w)/4), y + float64(int(h)/3+int(h)/4) + 0.5}}))
			fill(dst, image.Rect(int(x)+int(w)/4, int(y)+int(h)/3, int(x+w)-int(w)/4, int(y)+int(h)/3+int(h)/4), s.Bg[Selected])
		}
		return true
	case "OPTION":
		outline, c2 := cruxOutline(s.Bg[Normal]), shade(s.Bg[Normal], 0.7)
		cx, cy := x+float64(int(h)/2), y+float64(int(h)/2)
		rho := float64(int(h)/2) - 0.5
		if c.Detail != "option" {
			if st == Insensitive {
				fillDisc(dst, cx, cy, rho, s.Bg[Insensitive], nil)
				strokeCircle(dst, cx, cy, rho, 1, s.Fg[Insensitive])
			} else {
				fillRegionFunc(dst, r.Inset(-1), func(px, py float64) RGBA {
					t := ((px - x) + (py - y)) / (2 * h)
					if st == Active {
						return lerpRGBA(c2, RGBA{1, 1, 1, 1}, t)
					}
					return stops(t, []float64{0, 0.2, 1}, []RGBA{{1, 1, 1, 1}, {1, 1, 1, 1}, c2})
				}, func(px, py float64) bool { return math.Hypot(px-cx, py-cy) <= rho })
				strokeCircle(dst, cx, cy, rho, 1, outline)
				strokeCircle(dst, cx, cy, rho-1, 1, RGBA{1, 1, 1, 0.5})
			}
			rho /= 2
		}
		if c.Shadow == "IN" {
			fill, stroke := s.light(Selected), s.dark(Selected)
			if st == Insensitive {
				fill, stroke = s.Fg[Insensitive], s.Fg[Insensitive]
			}
			fillDisc(dst, cx, cy, rho, fill, nil)
			strokeCircle(dst, cx, cy, rho, 1, stroke)
		}
		return true
	case "ARROW":
		ax, ay := float64(r.Min.X+r.Dx()/2-2), float64(r.Min.Y+r.Dy()/2-2)
		if c.Detail == "spinbutton" {
			ax--
			if c.ArrowDirection == "UP" {
				ay++
			}
		}
		var pts [][2]float64
		switch c.ArrowDirection {
		case "UP":
			pts = [][2]float64{{ax, ay + 4}, {ax + 5, ay + 4}, {ax + 2.5, ay}}
		case "LEFT":
			pts = [][2]float64{{ax + 4, ay}, {ax, ay + 2.5}, {ax + 4, ay + 5}}
		case "RIGHT":
			pts = [][2]float64{{ax, ay}, {ax + 4, ay + 2.5}, {ax, ay + 5}}
		default:
			pts = [][2]float64{{ax, ay}, {ax + 5, ay}, {ax + 2.5, ay + 4}}
		}
		fillPolygon(dst, pts, s.Fg[st])
		return true
	case "HLINE":
		yy := r.Min.Y + r.Dy()/2 - 1
		hline(dst, r.Min.X, r.Max.X-1, yy, shade(bg, 0.88))
		hline(dst, r.Min.X, r.Max.X-1, yy+1, shade(bg, 1.12))
		return true
	case "VLINE":
		xx := r.Min.X + r.Dx()/2 - 1
		vline(dst, xx, r.Min.Y, r.Max.Y-1, shade(bg, 0.88))
		vline(dst, xx+1, r.Min.Y, r.Max.Y-1, shade(bg, 1.12))
		return true
	case "BOX_GAP":
		fill(dst, r, bg)
		cruxStdShadow(dst, r, bg, c.Shadow)
		return true
	case "EXTENSION":
		cruxTab(s, dst, r, c)
		return true
	}
	return false
}

func cruxButton(s *Style, dst *image.NRGBA, x, y, w, h float64, st int, shadow string, halo bool) {
	k := s.Bg[st]
	stroke := lightnessOffset(k, -0.52)
	r := 2.5
	if halo {
		col := RGBA{0, 0, 0, 0.1}
		if st == Insensitive {
			col = RGBA{0, 0, 0, 0.05}
		}
		fillRR(dst, x, y, w, h, 3.5, allCorners, solid(col))
		x, y, w, h = x+1, y+1, w-2, h-2
	}
	switch shadow {
	case "OUT", "ETCHED_OUT":
		if st == Insensitive {
			fillRR(dst, x, y, w, h, r, allCorners, solid(s.Fg[Insensitive]))
			fillRR(dst, x+1, y+1, w-2, h-2, r-1, allCorners, solid(s.Bg[Insensitive]))
			return
		}
		fillRR(dst, x, y, w, h, r, allCorners, solid(stroke))
		fillRR(dst, x+1, y+1, w-2, h-2, r-1, allCorners, vgrad(y+1, y+h-1, []float64{0, 1}, []RGBA{lightnessOffset(k, 0.156), k}))
		ix, iy, iw, ih := x+1.5, y+1.5, w-3, h-3
		fillRR(dst, ix, iy, iw, ih, 0, allCorners, vgrad(iy, iy+ih, []float64{0, 0.14, 0.86, 1}, []RGBA{lightnessOffset(k, 0.117), lightnessOffset(k, 0.117), lightnessOffset(k, -0.117), lightnessOffset(k, -0.117)}))
	case "IN", "ETCHED_IN":
		if st == Insensitive {
			fillRR(dst, x, y, w, h, r, allCorners, solid(s.Fg[Insensitive]))
			fillRR(dst, x+1, y+1, w-2, h-2, r, allCorners, solid(k))
			return
		}
		fillRR(dst, x, y, w, h, r, allCorners, solid(stroke))
		fillRR(dst, x+1, y+1, w-2, h-2, r, allCorners, vgrad(y+1, y+h-1, []float64{0, 0.3, 0.9, 1}, []RGBA{lightnessOffset(k, -0.05), lightnessOffset(k, -0.05), lightnessOffset(k, -0.2), lightnessOffset(k, -0.2)}))
		px, py, pw, ph := x+1.5, y+1.5, w-3, h-3
		polyline(dst, 1, RGBA{0, 0, 0, 0.2}, [2]float64{px, py + ph}, [2]float64{px, py}, [2]float64{px + pw, py})
		polyline(dst, 1, RGBA{0, 0, 0, 0.05}, [2]float64{px + 1, py + ph}, [2]float64{px + 1, py + 1}, [2]float64{px + pw, py + 1}, [2]float64{px + pw, py + ph})
	}
}

func cruxEntry(s *Style, dst *image.NRGBA, r image.Rectangle, st int, focused bool) {
	x, y, h := r.Min.X, r.Min.Y, r.Dy()
	if focused {
		ring(dst, r, s.light(Selected))
	} else {
		top, bottom := RGBA{185.0 / 255, 189.0 / 255, 182.0 / 255, 1}, RGBA{238.0 / 255, 238.0 / 255, 236.0 / 255, 1}
		for yy := r.Min.Y; yy < r.Max.Y; yy++ {
			col := lerpRGBA(top, bottom, float64(yy-y)/math.Max(float64(h-1), 1))
			if yy == r.Min.Y || yy == r.Max.Y-1 {
				hline(dst, r.Min.X, r.Max.X-1, yy, col)
			} else {
				setPixel(dst, r.Min.X, yy, col)
				setPixel(dst, r.Max.X-1, yy, col)
			}
		}
	}
	inner := cruxEntryLine
	base := s.Base[Normal]
	if st == Insensitive {
		inner, base = s.Fg[Insensitive], s.Base[Insensitive]
	}
	ring(dst, r.Inset(1), inner)
	ring(dst, r.Inset(2), base)
	polyline(dst, 1, RGBA{0, 0, 0, 0.12}, [2]float64{float64(x) + 2.5, float64(r.Max.Y) - 2.5}, [2]float64{float64(x) + 2.5, float64(y) + 2.5}, [2]float64{float64(r.Max.X) - 2.5, float64(y) + 2.5})
}

func cruxTab(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) {
	st := c.State
	bg := s.Bg[st]
	x, y, w, h := float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy())
	h += float64(s.YThickness) // tabs above the page
	corners := [4]bool{true, true, false, false}
	clip := image.NewNRGBA(dst.Bounds())
	fillRR(clip, x+0.5, y+0.5, w-1, h-1, 2, corners, vgrad(y, y+h, []float64{0, 0.3, 1}, []RGBA{shade(bg, 1.1), bg, bg}))
	strokeRR(clip, x+0.5, y+0.5, w-1, h-1, 2, corners, 1, solid(cruxOutline(bg)), nil)
	X, Y, W, H := x+1.5, y+1.5, w-3, h-3
	if st == Normal {
		H++
		polyline(clip, 1, cruxHighlight, [2]float64{X, Y + H}, [2]float64{X, Y + 1}, [2]float64{X + 1, Y}, [2]float64{X + W - 1, Y})
		polyline(clip, 1, cruxShadow, [2]float64{X + W, Y + 1}, [2]float64{X + W, Y + H - 1})
		blit(dst, clip, r.Union(image.Rect(r.Min.X, r.Min.Y, r.Max.X, r.Max.Y+s.YThickness)))
		return
	}
	strokeRR(clip, X, Y, W, H, 1, corners, 1, solid(RGBA{1, 1, 1, 0.5}), nil)
	blit(dst, clip, r)
}
