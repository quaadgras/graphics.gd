package gtkrc

import (
	"image"
	"math"
)

// Drawing primitives shared by the engine implementations, following the
// conventions of the gtk-engines support library: 1px lines through pixel
// centres with square caps, and the two-colour "simple border".

func (s *Style) mid(st int) RGBA {
	l, d := s.light(st), s.dark(st)
	return RGBA{(l.R + d.R) / 2, (l.G + d.G) / 2, (l.B + d.B) / 2, 1}
}

func mix(a, b RGBA, t float64) RGBA {
	return RGBA{a.R*(1-t) + b.R*t, a.G*(1-t) + b.G*t, a.B*(1-t) + b.B*t, 1}
}

func setPixel(dst *image.NRGBA, x, y int, c RGBA) { blend(dst, x, y, c, 1) }

// hline and vline paint whole pixels (butt caps).
func hline(dst *image.NRGBA, x1, x2, y int, c RGBA) {
	for x := x1; x <= x2; x++ {
		setPixel(dst, x, y, c)
	}
}

func vline(dst *image.NRGBA, x, y1, y2 int, c RGBA) {
	for y := y1; y <= y2; y++ {
		setPixel(dst, x, y, c)
	}
}

// simpleBorder draws the 1px L-shapes: tl along the left column and top row,
// br along the bottom row and right column. By default br is drawn second and
// owns the bottom-left and top-right pixels; overlap draws tl second.
func simpleBorder(dst *image.NRGBA, r image.Rectangle, tl, br RGBA, overlap bool) {
	if r.Dx() <= 0 || r.Dy() <= 0 {
		return
	}
	drawTL := func() {
		vline(dst, r.Min.X, r.Min.Y, r.Max.Y-1, tl)
		hline(dst, r.Min.X, r.Max.X-1, r.Min.Y, tl)
	}
	drawBR := func() {
		hline(dst, r.Min.X, r.Max.X-1, r.Max.Y-1, br)
		vline(dst, r.Max.X-1, r.Min.Y, r.Max.Y-1, br)
	}
	if overlap {
		drawBR()
		drawTL()
	} else {
		drawTL()
		drawBR()
	}
}

// ring paints the 1px border of an integer rectangle.
func ring(dst *image.NRGBA, r image.Rectangle, c RGBA) { simpleBorder(dst, r, c, c, false) }

// coverage is the fraction of a pixel inside a region, by 4×4 supersampling.
func coverage(x, y int, in func(px, py float64) bool) float64 {
	n := 0
	for sy := 0; sy < 4; sy++ {
		for sx := 0; sx < 4; sx++ {
			if in(float64(x)+(float64(sx)+0.5)/4, float64(y)+(float64(sy)+0.5)/4) {
				n++
			}
		}
	}
	return float64(n) / 16
}

// fillRegion blends c over every pixel of r by its coverage.
func fillRegion(dst *image.NRGBA, r image.Rectangle, c RGBA, in func(px, py float64) bool) {
	r = r.Intersect(dst.Bounds())
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if cov := coverage(x, y, in); cov > 0 {
				blend(dst, x, y, c, cov)
			}
		}
	}
}

// inPolygon is the even-odd point-in-polygon test over (x, y) pairs.
func inPolygon(pts [][2]float64) func(px, py float64) bool {
	return func(px, py float64) bool {
		in := false
		for i, j := 0, len(pts)-1; i < len(pts); j, i = i, i+1 {
			xi, yi, xj, yj := pts[i][0], pts[i][1], pts[j][0], pts[j][1]
			if (yi > py) != (yj > py) && px < (xj-xi)*(py-yi)/(yj-yi)+xi {
				in = !in
			}
		}
		return in
	}
}

func polyBounds(pts [][2]float64) image.Rectangle {
	minX, minY, maxX, maxY := math.Inf(1), math.Inf(1), math.Inf(-1), math.Inf(-1)
	for _, p := range pts {
		minX, maxX = math.Min(minX, p[0]), math.Max(maxX, p[0])
		minY, maxY = math.Min(minY, p[1]), math.Max(maxY, p[1])
	}
	return image.Rect(int(math.Floor(minX))-1, int(math.Floor(minY))-1, int(math.Ceil(maxX))+1, int(math.Ceil(maxY))+1)
}

func fillPolygon(dst *image.NRGBA, pts [][2]float64, c RGBA) {
	fillRegion(dst, polyBounds(pts), c, inPolygon(pts))
}

// segment strokes a line of width w from (x0, y0) to (x1, y1); square caps
// extend it by half the width at each end.
func segment(dst *image.NRGBA, x0, y0, x1, y1, w float64, c RGBA, square bool) {
	dx, dy := x1-x0, y1-y0
	l := math.Hypot(dx, dy)
	if l == 0 {
		fillRegion(dst, image.Rect(int(x0-w), int(y0-w), int(x0+w)+1, int(y0+w)+1), c, func(px, py float64) bool {
			return math.Abs(px-x0) <= w/2 && math.Abs(py-y0) <= w/2
		})
		return
	}
	ux, uy := dx/l, dy/l
	ext := 0.0
	if square {
		ext = w / 2
	}
	fillRegion(dst, image.Rect(int(math.Min(x0, x1)-w-1), int(math.Min(y0, y1)-w-1), int(math.Max(x0, x1)+w+2), int(math.Max(y0, y1)+w+2)), c,
		func(px, py float64) bool {
			t := (px-x0)*ux + (py-y0)*uy
			n := -(px-x0)*uy + (py-y0)*ux
			return t >= -ext && t <= l+ext && math.Abs(n) <= w/2
		})
}

// polyline strokes 1px (or w) lines through points, square caps.
func polyline(dst *image.NRGBA, w float64, c RGBA, pts ...[2]float64) {
	for i := 0; i+1 < len(pts); i++ {
		segment(dst, pts[i][0], pts[i][1], pts[i+1][0], pts[i+1][1], w, c, true)
	}
}

// fillDisc fills a circle, optionally clipped by another region.
func fillDisc(dst *image.NRGBA, cx, cy, r float64, c RGBA, clip func(px, py float64) bool) {
	b := image.Rect(int(cx-r)-1, int(cy-r)-1, int(cx+r)+2, int(cy+r)+2)
	fillRegion(dst, b, c, func(px, py float64) bool {
		if math.Hypot(px-cx, py-cy) > r {
			return false
		}
		return clip == nil || clip(px, py)
	})
}

// strokeCircle strokes a circle outline of width w.
func strokeCircle(dst *image.NRGBA, cx, cy, r, w float64, c RGBA) {
	b := image.Rect(int(cx-r-w)-1, int(cy-r-w)-1, int(cx+r+w)+2, int(cy+r+w)+2)
	fillRegion(dst, b, c, func(px, py float64) bool {
		return math.Abs(math.Hypot(px-cx, py-cy)-r) <= w/2
	})
}

// gtkArrow computes the triangle GTK 2 engines draw for an arrow in rect r
// (the shape shared by ThinIce and others), as integer pixel vertices.
func gtkArrow(r image.Rectangle, dir string) [3][2]int {
	x, y, w, h := r.Min.X, r.Min.Y, r.Dx(), r.Dy()
	odd := func(v int) int {
		if v%2 != 0 {
			return 1
		}
		return 0
	}
	if dir == "UP" || dir == "DOWN" {
		var aw, ah int
		t := (w+1)/2 - odd(h)
		if t > h {
			aw = 2*h - 1 - odd(h)
			ah = (aw + 1) / 2
		} else {
			ah = t
			aw = 2*ah - 1
		}
		if aw < 5 || ah < 3 {
			aw, ah = 5, 3
		}
		x += (w - aw) / 2
		y += (h - ah) / 2
		if aw%2 == 0 {
			aw--
		}
		yb, ya := y, y+ah-1
		if dir == "UP" {
			yb, ya = y+ah-1, y
		}
		return [3][2]int{{x, yb}, {x + aw - 1, yb}, {x + aw/2, ya}}
	}
	var aw, ah int
	t := (h+1)/2 - odd(w)
	if t > w {
		ah = 2*w - 1 - odd(w)
		aw = (ah + 1) / 2
	} else {
		aw = t
		ah = 2*aw - 1
	}
	if ah < 5 || aw < 3 {
		ah, aw = 5, 3
	}
	x += (w - aw) / 2
	y += (h - ah) / 2
	if ah%2 == 0 {
		ah--
	}
	xb, xa := x, x+aw-1
	if dir == "LEFT" {
		xb, xa = x+aw-1, x
	}
	return [3][2]int{{xb, y}, {xb, y + ah - 1}, {xa, y + ah/2}}
}

// drawArrowShape fills the arrow triangle through pixel centres, with a thin
// anti-aliased edge, as ThinIce does (stroke 0.5 then fill).
func drawArrowShape(dst *image.NRGBA, v [3][2]int, c RGBA) {
	pts := [][2]float64{
		{float64(v[0][0]) + 0.5, float64(v[0][1]) + 0.5},
		{float64(v[1][0]) + 0.5, float64(v[1][1]) + 0.5},
		{float64(v[2][0]) + 0.5, float64(v[2][1]) + 0.5},
	}
	for i := range pts {
		j := (i + 1) % 3
		segment(dst, pts[i][0], pts[i][1], pts[j][0], pts[j][1], 0.5, c, false)
	}
	fillPolygon(dst, pts, c)
}

// inRoundRect tests points against a rectangle with radius r on the given
// corners (top-left, top-right, bottom-right, bottom-left).
func inRoundRect(x, y, w, h, r float64, corners [4]bool) func(px, py float64) bool {
	return func(px, py float64) bool {
		if px < x || py < y || px > x+w || py > y+h {
			return false
		}
		if r <= 0 {
			return true
		}
		centres := [4][2]float64{{x + r, y + r}, {x + w - r, y + r}, {x + w - r, y + h - r}, {x + r, y + h - r}}
		for i, c := range centres {
			if !corners[i] {
				continue
			}
			inCorner := (i == 0 && px < c[0] && py < c[1]) || (i == 1 && px > c[0] && py < c[1]) ||
				(i == 2 && px > c[0] && py > c[1]) || (i == 3 && px < c[0] && py > c[1])
			if inCorner && math.Hypot(px-c[0], py-c[1]) > r {
				return false
			}
		}
		return true
	}
}

var allCorners = [4]bool{true, true, true, true}

// strokeRoundRect strokes, 1px wide, the path through pixel centres of the
// integer rect r, with radius rad on the given corners.
func strokeRoundRect(dst *image.NRGBA, r image.Rectangle, rad float64, corners [4]bool, c RGBA) {
	x, y, w, h := float64(r.Min.X), float64(r.Min.Y), float64(r.Dx()), float64(r.Dy())
	outer := inRoundRect(x, y, w, h, rad+0.5, corners)
	inner := inRoundRect(x+1, y+1, w-2, h-2, math.Max(0, rad-0.5), corners)
	if rad <= 0 {
		outer = inRoundRect(x, y, w, h, 0, corners)
		inner = inRoundRect(x+1, y+1, w-2, h-2, 0, corners)
	}
	fillRegion(dst, r.Inset(-1), c, func(px, py float64) bool { return outer(px, py) && !inner(px, py) })
}

func withAlpha(c RGBA, a float64) RGBA { c.A = clamp01(a); return c }

// over composites a translucent colour onto an opaque one.
func over(top, bottom RGBA) RGBA {
	return RGBA{bottom.R*(1-top.A) + top.R*top.A, bottom.G*(1-top.A) + top.G*top.A, bottom.B*(1-top.A) + top.B*top.A, 1}
}

// bezier samples a cubic Bézier into points (excluding the start).
func bezier(p0, p1, p2, p3 [2]float64, n int) [][2]float64 {
	var out [][2]float64
	for i := 1; i <= n; i++ {
		t := float64(i) / float64(n)
		u := 1 - t
		out = append(out, [2]float64{
			u*u*u*p0[0] + 3*u*u*t*p1[0] + 3*u*t*t*p2[0] + t*t*t*p3[0],
			u*u*u*p0[1] + 3*u*u*t*p1[1] + 3*u*t*t*p2[1] + t*t*t*p3[1],
		})
	}
	return out
}

// fillRegionFunc blends a per-pixel colour (sampled at the pixel centre)
// over the pixels of r inside the region.
func fillRegionFunc(dst *image.NRGBA, r image.Rectangle, colorAt func(px, py float64) RGBA, in func(px, py float64) bool) {
	r = r.Intersect(dst.Bounds())
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			if cov := coverage(x, y, in); cov > 0 {
				blend(dst, x, y, colorAt(float64(x)+0.5, float64(y)+0.5), cov)
			}
		}
	}
}

// lerpRGBA interpolates colours including alpha.
func lerpRGBA(a, b RGBA, t float64) RGBA {
	t = clamp01(t)
	return RGBA{a.R + (b.R-a.R)*t, a.G + (b.G-a.G)*t, a.B + (b.B-a.B)*t, a.A + (b.A-a.A)*t}
}

// stops evaluates a multi-stop gradient at t (stops sorted by offset).
func stops(t float64, offs []float64, cols []RGBA) RGBA {
	if t <= offs[0] {
		return cols[0]
	}
	for i := 1; i < len(offs); i++ {
		if t <= offs[i] {
			return lerpRGBA(cols[i-1], cols[i], (t-offs[i-1])/math.Max(offs[i]-offs[i-1], 1e-9))
		}
	}
	return cols[len(cols)-1]
}

// lightnessOffset sets HSL lightness to L+d without clamping L, then clamps
// the channels (Crux's colour offsets).
func lightnessOffset(c RGBA, d float64) RGBA {
	h, l, s := rgbToHLS(c.R, c.G, c.B)
	r, g, b := hlsToRGB(h, l+d, s)
	return RGBA{clamp01(r), clamp01(g), clamp01(b), c.A}
}

// rrSDF is the signed distance to a rounded rectangle path (negative inside)
// with radius r on the given corners, in float coordinates.
func rrSDF(x, y, w, h, r float64, corners [4]bool) func(px, py float64) float64 {
	return func(px, py float64) float64 {
		cx, cy := x+w/2, y+h/2
		var rad float64
		var corner int
		switch {
		case px < cx && py < cy:
			corner = 0
		case px >= cx && py < cy:
			corner = 1
		case px >= cx && py >= cy:
			corner = 2
		default:
			corner = 3
		}
		if corners[corner] {
			rad = math.Max(0, math.Min(r, math.Min(w, h)/2))
		}
		qx := math.Abs(px-cx) - w/2 + rad
		qy := math.Abs(py-cy) - h/2 + rad
		return math.Hypot(math.Max(qx, 0), math.Max(qy, 0)) + math.Min(math.Max(qx, qy), 0) - rad
	}
}

func rrBounds(x, y, w, h float64) image.Rectangle {
	return image.Rect(int(math.Floor(x))-2, int(math.Floor(y))-2, int(math.Ceil(x+w))+2, int(math.Ceil(y+h))+2)
}

// fillRR fills a rounded rectangle with a colour function.
func fillRR(dst *image.NRGBA, x, y, w, h, r float64, corners [4]bool, col func(px, py float64) RGBA) {
	if w <= 0 || h <= 0 {
		return
	}
	d := rrSDF(x, y, w, h, r, corners)
	fillRegionFunc(dst, rrBounds(x, y, w, h), col, func(px, py float64) bool { return d(px, py) <= 0 })
}

// strokeRR strokes a rounded rectangle path with a line of width lw,
// optionally masked to a region.
func strokeRR(dst *image.NRGBA, x, y, w, h, r float64, corners [4]bool, lw float64, col func(px, py float64) RGBA, mask func(px, py float64) bool) {
	if w <= 0 || h <= 0 {
		return
	}
	d := rrSDF(x, y, w, h, r, corners)
	fillRegionFunc(dst, rrBounds(x, y, w, h), col, func(px, py float64) bool {
		return math.Abs(d(px, py)) <= lw/2 && (mask == nil || mask(px, py))
	})
}

func solid(c RGBA) func(px, py float64) RGBA { return func(float64, float64) RGBA { return c } }

// vgrad is a vertical gradient between y0 and y1 through stops.
func vgrad(y0, y1 float64, offs []float64, cols []RGBA) func(px, py float64) RGBA {
	return func(px, py float64) RGBA { return stops((py-y0)/math.Max(y1-y0, 1e-9), offs, cols) }
}

func hgrad(x0, x1 float64, offs []float64, cols []RGBA) func(px, py float64) RGBA {
	return func(px, py float64) RGBA { return stops((px-x0)/math.Max(x1-x0, 1e-9), offs, cols) }
}
