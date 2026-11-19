package gtkrc

import (
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"math"
	"os"
	"sync"
)

// Call is one GTK 2 style drawing call (gtk_paint_box, gtk_paint_check...).
type Call struct {
	Function       string // BOX, FLAT_BOX, SHADOW, CHECK, OPTION, ARROW, SLIDER, BOX_GAP, EXTENSION, HLINE, FOCUS, TAB
	State          int
	Shadow         string // NONE, IN, OUT, ETCHED_IN, ETCHED_OUT
	Detail         string
	ArrowDirection string // UP, DOWN, LEFT, RIGHT
	GapSide        string // TOP, BOTTOM, LEFT, RIGHT
	Orientation    string // HORIZONTAL, VERTICAL
	Widget         string // the widget's class, for engines that look at it (GtkHScale, GtkProgressBar...)
	Focused        bool   // the widget has keyboard focus
}

// Match returns the pixmap engine image used for a call, if any.
func (s *Style) Match(c Call) (Image, bool) {
	if s.Engine != "pixmap" {
		return Image{}, false
	}
	for _, img := range s.Images {
		if img.Function != c.Function && !(img.Function == "BOX" && c.Function == "BOX_GAP" && false) {
			continue
		}
		if img.Detail != "" && img.Detail != c.Detail ||
			img.State >= 0 && img.State != c.State ||
			img.Shadow != "" && img.Shadow != c.Shadow ||
			img.ArrowDirection != "" && img.ArrowDirection != c.ArrowDirection ||
			img.GapSide != "" && img.GapSide != c.GapSide ||
			img.Orientation != "" && img.Orientation != c.Orientation {
			continue
		}
		return img, true
	}
	return Image{}, false
}

// Draw renders a call into a w×h image.
func (s *Style) Draw(c Call, w, h int) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	s.DrawInto(img, img.Bounds(), c)
	return img
}

// Engine draws the calls a theme engine handles itself, and reports false
// for those it leaves to GTK's default drawing.
type Engine func(s *Style, dst *image.NRGBA, r image.Rectangle, c Call) bool

// engines are the Go re-implementations of GTK 2 theme engines, by the name
// gtkrc uses for them (engine "clearlooks" {...}).
var engines = map[string]Engine{}

// DrawInto renders a call into a rectangle of dst.
func (s *Style) DrawInto(dst *image.NRGBA, r image.Rectangle, c Call) {
	if engine, ok := engines[s.Engine]; ok && engine(s, dst, r, c) {
		return
	}
	if pm, ok := s.Match(c); ok {
		drawn := false
		if src := loadImage(pm.File); src != nil {
			if pm.Stretch {
				nineSlice(dst, r, src, pm.Border)
			} else {
				tile(dst, r, src)
			}
			drawn = true
		}
		if src := loadImage(pm.OverlayFile); src != nil {
			if pm.OverlayStretch {
				nineSlice(dst, r, src, pm.OverlayBorder)
			} else {
				b := src.Bounds()
				at := image.Pt(r.Min.X+(r.Dx()-b.Dx())/2, r.Min.Y+(r.Dy()-b.Dy())/2)
				draw.Draw(dst, b.Sub(b.Min).Add(at), src, b.Min, draw.Over)
			}
			drawn = true
		}
		if drawn {
			return
		}
	}
	s.drawDefault(dst, r, c)
}

// default GTK 2 drawing

func (s *Style) light(st int) RGBA { return shade(s.Bg[st], 1.3) }
func (s *Style) dark(st int) RGBA  { return shade(s.Bg[st], 0.7) }

var black = RGBA{0, 0, 0, 1}

func (s *Style) drawDefault(dst *image.NRGBA, r image.Rectangle, c Call) {
	st := c.State
	switch c.Function {
	case "FLAT_BOX":
		col := s.Bg[st]
		switch c.Detail {
		case "entry_bg", "text", "viewportbin", "treeview", "base":
			col = s.Base[st]
		}
		fill(dst, r, col)
	case "BOX", "BOX_GAP", "SLIDER", "TAB":
		fill(dst, r, s.Bg[st])
		s.shadow(dst, r, st, c.Shadow, "")
		if c.Function == "TAB" {
			// the option menu indicator: a small raised box
		}
	case "EXTENSION": // a notebook tab: no border on the gap side
		fill(dst, r, s.Bg[st])
		s.shadow(dst, r, st, "OUT", c.GapSide)
	case "SHADOW":
		s.shadow(dst, r, st, c.Shadow, "")
	case "HLINE":
		y := r.Min.Y + r.Dy()/2 - 1
		fill(dst, image.Rect(r.Min.X, y, r.Max.X, y+1), s.dark(st))
		fill(dst, image.Rect(r.Min.X, y+1, r.Max.X, y+2), s.light(st))
	case "CHECK":
		fill(dst, r, s.Base[st])
		s.shadow(dst, r, st, "IN", "")
		if c.Shadow == "IN" {
			inner := r.Inset(3)
			mark := s.Text[st]
			// a 7x7 check mark scaled to the box
			w := float64(inner.Dx())
			pts := [][2]float64{{0, 0.45}, {0.35, 0.8}, {1, 0.1}}
			for i := 0; i+1 < len(pts); i++ {
				line(dst, float64(inner.Min.X)+pts[i][0]*w, float64(inner.Min.Y)+pts[i][1]*w,
					float64(inner.Min.X)+pts[i+1][0]*w, float64(inner.Min.Y)+pts[i+1][1]*w, math.Max(1.5, w/5), mark)
			}
		}
	case "OPTION":
		cx, cy := float64(r.Min.X)+float64(r.Dx())/2, float64(r.Min.Y)+float64(r.Dy())/2
		rad := math.Min(float64(r.Dx()), float64(r.Dy())) / 2
		disc(dst, cx, cy, rad, s.dark(st))
		disc(dst, cx+0.5, cy+0.5, rad-0.5, s.light(st))
		disc(dst, cx, cy, rad-1, s.Base[st])
		if c.Shadow == "IN" {
			disc(dst, cx, cy, rad*0.4, s.Text[st])
		}
	case "ARROW":
		arrow(dst, r, c.ArrowDirection, s.Fg[st])
	case "FOCUS":
		dotted(dst, r, s.Fg[st])
	}
}

// shadow draws GTK 2's default bevel, xthickness/ythickness pixels wide,
// leaving out the gap side.
func (s *Style) shadow(dst *image.NRGBA, r image.Rectangle, st int, kind, gap string) {
	var tl, br [2]RGBA // outer, inner
	switch kind {
	case "OUT":
		tl, br = [2]RGBA{s.light(st), s.Bg[st]}, [2]RGBA{black, s.dark(st)}
	case "IN":
		tl, br = [2]RGBA{s.dark(st), black}, [2]RGBA{s.light(st), s.Bg[st]}
	case "ETCHED_IN":
		tl, br = [2]RGBA{s.dark(st), s.light(st)}, [2]RGBA{s.light(st), s.dark(st)}
	case "ETCHED_OUT":
		tl, br = [2]RGBA{s.light(st), s.dark(st)}, [2]RGBA{s.dark(st), s.light(st)}
	default:
		return
	}
	if s.XThickness == 1 || s.YThickness == 1 {
		if kind == "OUT" {
			br[0] = s.dark(st)
		} else if kind == "IN" {
			tl[0] = s.dark(st)
		}
	}
	for i := 0; i < 2; i++ {
		x0, y0, x1, y1 := r.Min.X+i, r.Min.Y+i, r.Max.X-1-i, r.Max.Y-1-i
		if x1 < x0 || y1 < y0 {
			break
		}
		if i < s.YThickness && gap != "TOP" {
			fill(dst, image.Rect(x0, y0, x1+1, y0+1), tl[i])
		}
		if i < s.XThickness && gap != "LEFT" {
			fill(dst, image.Rect(x0, y0, x0+1, y1+1), tl[i])
		}
		if i < s.YThickness && gap != "BOTTOM" {
			fill(dst, image.Rect(x0, y1, x1+1, y1+1), br[i])
		}
		if i < s.XThickness && gap != "RIGHT" {
			fill(dst, image.Rect(x1, y0, x1+1, y1+1), br[i])
		}
	}
}

// pixel helpers

func nrgba(c RGBA) color.NRGBA {
	return color.NRGBA{uint8(clamp01(c.R)*255 + 0.5), uint8(clamp01(c.G)*255 + 0.5), uint8(clamp01(c.B)*255 + 0.5), uint8(clamp01(c.A)*255 + 0.5)}
}

func fill(dst *image.NRGBA, r image.Rectangle, c RGBA) {
	draw.Draw(dst, r, &image.Uniform{nrgba(c)}, image.Point{}, draw.Over)
}

func blend(dst *image.NRGBA, x, y int, c RGBA, cov float64) {
	if !(image.Point{x, y}.In(dst.Bounds())) || cov <= 0 {
		return
	}
	src := c
	src.A *= clamp01(cov)
	under := dst.NRGBAAt(x, y)
	ua := float64(under.A) / 255
	oa := src.A + ua*(1-src.A)
	if oa <= 0 {
		return
	}
	mix := func(sc float64, uc uint8) uint8 {
		return uint8(clamp01((sc*src.A+float64(uc)/255*ua*(1-src.A))/oa)*255 + 0.5)
	}
	dst.SetNRGBA(x, y, color.NRGBA{mix(src.R, under.R), mix(src.G, under.G), mix(src.B, under.B), uint8(oa*255 + 0.5)})
}

func disc(dst *image.NRGBA, cx, cy, r float64, c RGBA) {
	for y := int(cy - r - 1); y <= int(cy+r+1); y++ {
		for x := int(cx - r - 1); x <= int(cx+r+1); x++ {
			d := math.Hypot(float64(x)+0.5-cx, float64(y)+0.5-cy)
			blend(dst, x, y, c, r-d+0.5)
		}
	}
}

func line(dst *image.NRGBA, x0, y0, x1, y1, w float64, c RGBA) {
	minX, maxX := int(math.Min(x0, x1)-w), int(math.Max(x0, x1)+w)
	minY, maxY := int(math.Min(y0, y1)-w), int(math.Max(y0, y1)+w)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			px, py := float64(x)+0.5-x0, float64(y)+0.5-y0
			dx, dy := x1-x0, y1-y0
			t := clamp01((px*dx + py*dy) / (dx*dx + dy*dy))
			d := math.Hypot(px-dx*t, py-dy*t)
			blend(dst, x, y, c, w/2-d+0.5)
		}
	}
}

func arrow(dst *image.NRGBA, r image.Rectangle, dir string, c RGBA) {
	size := min(r.Dx(), r.Dy())
	if dir == "UP" || dir == "DOWN" {
		size = min(r.Dx(), r.Dy()*2-1)
	}
	if size%2 == 0 {
		size--
	}
	half := (size + 1) / 2
	cx, cy := r.Min.X+r.Dx()/2, r.Min.Y+r.Dy()/2
	for i := 0; i < half; i++ {
		switch dir {
		case "DOWN":
			fill(dst, image.Rect(cx-half+1+i, cy-half/2+i, cx+half-i, cy-half/2+i+1), c)
		case "UP":
			fill(dst, image.Rect(cx-half+1+i, cy+half/2-i, cx+half-i, cy+half/2-i+1), c)
		case "RIGHT":
			fill(dst, image.Rect(cx-half/2+i, cy-half+1+i, cx-half/2+i+1, cy+half-i), c)
		case "LEFT":
			fill(dst, image.Rect(cx+half/2-i, cy-half+1+i, cx+half/2-i+1, cy+half-i), c)
		}
	}
}

func dotted(dst *image.NRGBA, r image.Rectangle, c RGBA) {
	for x := r.Min.X; x < r.Max.X; x += 2 {
		blend(dst, x, r.Min.Y, c, 1)
		blend(dst, x, r.Max.Y-1, c, 1)
	}
	for y := r.Min.Y; y < r.Max.Y; y += 2 {
		blend(dst, r.Min.X, y, c, 1)
		blend(dst, r.Max.X-1, y, c, 1)
	}
}

// nineSlice draws src into r, keeping its border (left, right, top, bottom) unscaled.
func nineSlice(dst *image.NRGBA, r image.Rectangle, src image.Image, border [4]int) {
	sb := src.Bounds()
	l, rt, t, b := border[0], border[1], border[2], border[3]
	sx := [4]int{sb.Min.X, sb.Min.X + l, sb.Max.X - rt, sb.Max.X}
	sy := [4]int{sb.Min.Y, sb.Min.Y + t, sb.Max.Y - b, sb.Max.Y}
	dx := [4]int{r.Min.X, r.Min.X + l, r.Max.X - rt, r.Max.X}
	dy := [4]int{r.Min.Y, r.Min.Y + t, r.Max.Y - b, r.Max.Y}
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			sw, sh := sx[i+1]-sx[i], sy[j+1]-sy[j]
			dw, dh := dx[i+1]-dx[i], dy[j+1]-dy[j]
			if sw <= 0 || sh <= 0 || dw <= 0 || dh <= 0 {
				continue
			}
			for y := 0; y < dh; y++ {
				for x := 0; x < dw; x++ {
					c := color.NRGBAModel.Convert(src.At(sx[i]+x*sw/dw, sy[j]+y*sh/dh)).(color.NRGBA)
					if c.A == 0 {
						continue
					}
					blend(dst, dx[i]+x, dy[j]+y, RGBA{float64(c.R) / 255, float64(c.G) / 255, float64(c.B) / 255, float64(c.A) / 255}, 1)
				}
			}
		}
	}
}

func tile(dst *image.NRGBA, r image.Rectangle, src image.Image) {
	sb := src.Bounds()
	for y := r.Min.Y; y < r.Max.Y; y += sb.Dy() {
		for x := r.Min.X; x < r.Max.X; x += sb.Dx() {
			draw.Draw(dst, image.Rect(x, y, x+sb.Dx(), y+sb.Dy()).Intersect(r), src, sb.Min, draw.Over)
		}
	}
}

var images sync.Map

func loadImage(path string) image.Image {
	if path == "" {
		return nil
	}
	if v, ok := images.Load(path); ok {
		img, _ := v.(image.Image)
		return img
	}
	var img image.Image
	if f, err := os.Open(path); err == nil {
		img, _, _ = image.Decode(f)
		f.Close()
	}
	images.Store(path, img)
	return img
}
