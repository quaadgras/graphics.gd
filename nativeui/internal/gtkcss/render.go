package gtkcss

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"math"
	"strings"
	"sync"
)

// Extents is how far drawing reaches outside the border box (outset shadows
// and outlines), per side: top, right, bottom, left.
type Extents [4]int

// Outside returns how far st draws outside its border box.
func (st Style) Outside() Extents {
	var e Extents
	grow := func(top, right, bottom, left float64) {
		for i, v := range [4]float64{top, right, bottom, left} {
			if n := int(math.Ceil(v)); n > e[i] {
				e[i] = n
			}
		}
	}
	for _, sh := range st.Shadows {
		if sh.Inset || sh.Color.A == 0 {
			continue
		}
		r := sh.Blur + sh.Spread
		grow(r-sh.Y, r+sh.X, r+sh.Y, r-sh.X)
	}
	if st.OutlineWidth > 0 && st.OutlineColor.A > 0 {
		o := st.OutlineWidth + st.OutlineOffset
		grow(o, o, o, o)
	}
	return e
}

// rrect is a rectangle with per-corner radii (top-left, top-right, bottom-right, bottom-left).
type rrect struct {
	x0, y0, x1, y1 float64
	r              [4]float64
}

func (r rrect) inset(top, right, bottom, left float64) rrect {
	out := rrect{r.x0 + left, r.y0 + top, r.x1 - right, r.y1 - bottom, r.r}
	out.r[0] = math.Max(0, r.r[0]-math.Max(top, left))
	out.r[1] = math.Max(0, r.r[1]-math.Max(top, right))
	out.r[2] = math.Max(0, r.r[2]-math.Max(bottom, right))
	out.r[3] = math.Max(0, r.r[3]-math.Max(bottom, left))
	return out
}

func (r rrect) shift(dx, dy float64) rrect {
	r.x0, r.x1, r.y0, r.y1 = r.x0+dx, r.x1+dx, r.y0+dy, r.y1+dy
	return r
}

// dist is a signed distance, negative inside.
func (r rrect) dist(x, y float64) float64 {
	cx, cy := (r.x0+r.x1)/2, (r.y0+r.y1)/2
	hw, hh := (r.x1-r.x0)/2, (r.y1-r.y0)/2
	if hw < 0 || hh < 0 {
		return math.Inf(1)
	}
	var rad float64
	switch {
	case x < cx && y < cy:
		rad = r.r[0]
	case x >= cx && y < cy:
		rad = r.r[1]
	case x >= cx && y >= cy:
		rad = r.r[2]
	default:
		rad = r.r[3]
	}
	rad = math.Min(rad, math.Min(hw, hh))
	qx := math.Abs(x-cx) - hw + rad
	qy := math.Abs(y-cy) - hh + rad
	return math.Hypot(math.Max(qx, 0), math.Max(qy, 0)) + math.Min(math.Max(qx, qy), 0) - rad
}

func (r rrect) inside(x, y float64) bool { return r.dist(x, y) < 0 }

// canvas accumulates premultiplied color.
type canvas struct {
	w, h int
	px   []float64
}

func (c *canvas) blend(x, y int, col RGBA, cov float64) {
	if x < 0 || y < 0 || x >= c.w || y >= c.h || cov <= 0 {
		return
	}
	a := col.A * cov
	i := (y*c.w + x) * 4
	c.px[i] = col.R*a + c.px[i]*(1-a)
	c.px[i+1] = col.G*a + c.px[i+1]*(1-a)
	c.px[i+2] = col.B*a + c.px[i+2]*(1-a)
	c.px[i+3] = a + c.px[i+3]*(1-a)
}

const ss = 4 // supersampling per axis

// paint blends col where in(x, y) holds, with supersampled coverage.
func (c *canvas) paint(col RGBA, in func(x, y float64) bool) {
	if col.A <= 0 {
		return
	}
	for y := 0; y < c.h; y++ {
		for x := 0; x < c.w; x++ {
			n := 0
			for sy := 0; sy < ss; sy++ {
				for sx := 0; sx < ss; sx++ {
					if in(float64(x)+(float64(sx)+0.5)/ss, float64(y)+(float64(sy)+0.5)/ss) {
						n++
					}
				}
			}
			c.blend(x, y, col, float64(n)/(ss*ss))
		}
	}
}

// paintFn blends a per-pixel color (sampled at the pixel center) where in holds.
func (c *canvas) paintFn(fn func(x, y float64) RGBA, in func(x, y float64) bool) {
	for y := 0; y < c.h; y++ {
		for x := 0; x < c.w; x++ {
			n := 0
			for sy := 0; sy < ss; sy++ {
				for sx := 0; sx < ss; sx++ {
					if in(float64(x)+(float64(sx)+0.5)/ss, float64(y)+(float64(sy)+0.5)/ss) {
						n++
					}
				}
			}
			if n > 0 {
				c.blend(x, y, fn(float64(x)+0.5, float64(y)+0.5), float64(n)/(ss*ss))
			}
		}
	}
}

// soft blends col with coverage from a signed distance blurred over blur px.
func (c *canvas) soft(col RGBA, dist func(x, y float64) float64, blur float64) {
	for y := 0; y < c.h; y++ {
		for x := 0; x < c.w; x++ {
			d := dist(float64(x)+0.5, float64(y)+0.5)
			cov := clamp01(0.5 - d/blur)
			cov = cov * cov * (3 - 2*cov)
			c.blend(x, y, col, cov)
		}
	}
}

func (c *canvas) image() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, c.w, c.h))
	for i := 0; i < c.w*c.h; i++ {
		a := c.px[i*4+3]
		if a <= 0 {
			continue
		}
		img.Pix[i*4+0] = uint8(clamp01(c.px[i*4+0]/a)*255 + 0.5)
		img.Pix[i*4+1] = uint8(clamp01(c.px[i*4+1]/a)*255 + 0.5)
		img.Pix[i*4+2] = uint8(clamp01(c.px[i*4+2]/a)*255 + 0.5)
		img.Pix[i*4+3] = uint8(clamp01(a)*255 + 0.5)
	}
	return img
}

var images sync.Map // path -> image.Image (or nil)

// LoadImage loads (and caches) a theme asset.
func LoadImage(path string) image.Image {
	if v, ok := images.Load(path); ok {
		img, _ := v.(image.Image)
		return img
	}
	var img image.Image
	if strings.HasSuffix(strings.ToLower(path), ".svg") {
		images.Store(path, img) // no SVG decoder: callers fall back to PNG versions
		return nil
	}
	if data, err := readFile(path); err == nil {
		img, _, _ = image.Decode(bytes.NewReader(data))
	}
	images.Store(path, img)
	return img
}

func sampleAt(img image.Image, x, y int) RGBA {
	b := img.Bounds()
	x = max(b.Min.X, min(b.Max.X-1, b.Min.X+x))
	y = max(b.Min.Y, min(b.Max.Y-1, b.Min.Y+y))
	c := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
	return RGBA{float64(c.R) / 255, float64(c.G) / 255, float64(c.B) / 255, float64(c.A) / 255}
}

// Render draws a node with the given border-box size. The image includes the
// node's Outside extents around the border box.
func Render(st Style, w, h int) *image.NRGBA {
	e := st.Outside()
	c := &canvas{w: w + e[1] + e[3], h: h + e[0] + e[2]}
	c.px = make([]float64, c.w*c.h*4)
	radius := st.Radius
	for i, pct := range st.RadiusPct {
		if pct > 0 {
			radius[i] = pct * math.Min(float64(w), float64(h))
		}
	}
	box := rrect{float64(e[3]), float64(e[0]), float64(e[3] + w), float64(e[0] + h), radius}
	bw := st.BorderWidth
	useImage := st.BorderImage.Kind == "url" && LoadImage(st.BorderImage.Path) != nil
	pad := box.inset(bw[0], bw[1], bw[2], bw[3])

	// outset shadows, last layer first, never under the box itself.
	for i := len(st.Shadows) - 1; i >= 0; i-- {
		sh := st.Shadows[i]
		if sh.Inset {
			continue
		}
		s := box.shift(sh.X, sh.Y).inset(-sh.Spread, -sh.Spread, -sh.Spread, -sh.Spread)
		if sh.Blur > 0 {
			c.soft(sh.Color, func(x, y float64) float64 {
				if box.inside(x, y) {
					return math.Inf(1)
				}
				return s.dist(x, y)
			}, sh.Blur)
		} else {
			c.paint(sh.Color, func(x, y float64) bool { return s.inside(x, y) && !box.inside(x, y) })
		}
	}

	// background
	c.paint(st.BackgroundColor, box.inside)
	for i := len(st.BackgroundImage) - 1; i >= 0; i-- {
		img := st.BackgroundImage[i]
		switch img.Kind {
		case "color":
			c.paint(img.Color, box.inside)
		case "linear-gradient":
			c.paintFn(func(x, y float64) RGBA { return gradientAt(img, box, x, y) }, box.inside)
		case "url":
			if src := LoadImage(img.Path); src != nil {
				sb := src.Bounds()
				c.paintFn(func(x, y float64) RGBA {
					return sampleAt(src, int(x-box.x0)%sb.Dx(), int(y-box.y0)%sb.Dy())
				}, box.inside)
			}
		}
	}

	// inset shadows, last layer first, clipped to the padding box.
	for i := len(st.Shadows) - 1; i >= 0; i-- {
		sh := st.Shadows[i]
		if !sh.Inset {
			continue
		}
		hole := pad.shift(sh.X, sh.Y).inset(sh.Spread, sh.Spread, sh.Spread, sh.Spread)
		if sh.Blur > 0 {
			c.soft(sh.Color, func(x, y float64) float64 {
				if !pad.inside(x, y) {
					return math.Inf(1)
				}
				return -hole.dist(x, y)
			}, sh.Blur)
		} else {
			c.paint(sh.Color, func(x, y float64) bool { return pad.inside(x, y) && !hole.inside(x, y) })
		}
	}

	// border
	if useImage {
		nineSlice(c, LoadImage(st.BorderImage.Path), st.BorderImageSlice, box, bw)
	} else if bw != [4]float64{} {
		for side := 0; side < 4; side++ {
			if bw[side] <= 0 || st.BorderColor[side].A <= 0 {
				continue
			}
			side := side
			in := func(x, y float64) bool {
				if !box.inside(x, y) || pad.inside(x, y) {
					return false
				}
				return borderSide(box, bw, x, y) == side && styleOn(st.BorderStyle[side], box, x, y, side)
			}
			c.paint(st.BorderColor[side], in)
		}
	}

	// outline
	if st.OutlineWidth > 0 && st.OutlineColor.A > 0 {
		o := st.OutlineOffset
		outer := box.inset(-o-st.OutlineWidth, -o-st.OutlineWidth, -o-st.OutlineWidth, -o-st.OutlineWidth)
		inner := box.inset(-o, -o, -o, -o)
		c.paint(st.OutlineColor, func(x, y float64) bool {
			if !outer.inside(x, y) || inner.inside(x, y) {
				return false
			}
			return styleOn(st.OutlineStyle, outer, x, y, borderSide(outer, [4]float64{1, 1, 1, 1}, x, y))
		})
	}
	return c.image()
}

// borderSide picks which side's border a point in the border area belongs to:
// the one it is relatively closest to, which splits corners diagonally.
func borderSide(box rrect, bw [4]float64, x, y float64) int {
	d := [4]float64{y - box.y0, box.x1 - x, box.y1 - y, x - box.x0}
	best, bestV := 0, math.Inf(1)
	for i, v := range d {
		if bw[i] <= 0 {
			continue
		}
		if r := v / bw[i]; r < bestV {
			best, bestV = i, r
		}
	}
	return best
}

// styleOn applies dotted/dashed patterns along a side.
func styleOn(style string, box rrect, x, y float64, side int) bool {
	var along float64
	if side == 0 || side == 2 {
		along = x - box.x0
	} else {
		along = y - box.y0
	}
	switch style {
	case "dotted":
		return int(math.Floor(along))%2 == 0
	case "dashed":
		return int(math.Floor(along/3))%2 == 0
	}
	return true
}

func gradientAt(img Image, box rrect, x, y float64) RGBA {
	// CSS: angle 0 points up, 90 right; the gradient line passes through the center.
	a := img.Angle * math.Pi / 180
	dx, dy := math.Sin(a), -math.Cos(a)
	w, h := box.x1-box.x0, box.y1-box.y0
	length := math.Abs(w*dx) + math.Abs(h*dy)
	cx, cy := (box.x0+box.x1)/2, (box.y0+box.y1)/2
	t := ((x-cx)*dx+(y-cy)*dy)/length + 0.5
	stops := img.Stops
	if t <= stops[0].Pos {
		return stops[0].Color
	}
	for i := 1; i < len(stops); i++ {
		if t <= stops[i].Pos {
			a, b := stops[i-1], stops[i]
			f := 0.0
			if b.Pos > a.Pos {
				f = (t - a.Pos) / (b.Pos - a.Pos)
			}
			return RGBA{a.Color.R + (b.Color.R-a.Color.R)*f, a.Color.G + (b.Color.G-a.Color.G)*f,
				a.Color.B + (b.Color.B-a.Color.B)*f, a.Color.A + (b.Color.A-a.Color.A)*f}
		}
	}
	return stops[len(stops)-1].Color
}

// nineSlice draws a border-image: the source's slices (top, right, bottom,
// left) stretched over the border widths; the middle is not drawn (no "fill").
func nineSlice(c *canvas, src image.Image, slice [4]float64, box rrect, bw [4]float64) {
	sb := src.Bounds()
	sw, sh := float64(sb.Dx()), float64(sb.Dy())
	// source and destination edges along each axis
	sx := [4]float64{0, slice[3], sw - slice[1], sw}
	sy := [4]float64{0, slice[0], sh - slice[2], sh}
	dx := [4]float64{box.x0, box.x0 + bw[3], box.x1 - bw[1], box.x1}
	dy := [4]float64{box.y0, box.y0 + bw[0], box.y1 - bw[2], box.y1}
	region := func(v float64, d [4]float64) int {
		switch {
		case v < d[1]:
			return 0
		case v < d[2]:
			return 1
		}
		return 2
	}
	for y := int(box.y0); y < int(math.Ceil(box.y1)); y++ {
		for x := int(box.x0); x < int(math.Ceil(box.x1)); x++ {
			px, py := float64(x)+0.5, float64(y)+0.5
			rx, ry := region(px, dx), region(py, dy)
			if rx == 1 && ry == 1 {
				continue
			}
			u := sx[rx] + (px-dx[rx])/math.Max(dx[rx+1]-dx[rx], 1e-9)*(sx[rx+1]-sx[rx])
			v := sy[ry] + (py-dy[ry])/math.Max(dy[ry+1]-dy[ry], 1e-9)*(sy[ry+1]-sy[ry])
			col := sampleAt(src, int(u), int(v))
			c.blend(x, y, col, 1)
		}
	}
}
