package nativeui

import (
	"image"
	"math"

	"graphics.gd/classdb/Image"
	"graphics.gd/classdb/ImageTexture"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/variant/Color"
)

// canvas is a tiny signed-distance-field rasterizer used to draw the theme's
// icons (check marks, switch knobs, slider thumbs...) at the theme's scale,
// so that they stay crisp instead of being bitmaps scaled by the engine.
type canvas struct {
	w, h int
	px   []float64 // premultiplied RGBA
}

func newCanvas(w, h int) *canvas {
	return &canvas{w: w, h: h, px: make([]float64, w*h*4)}
}

type sdf func(x, y float64) float64

// fill composites col over the canvas wherever the distance is negative,
// anti-aliased over one pixel (or over blur pixels, for soft shadows).
func (c *canvas) fill(d sdf, col Color.RGBA, blur float64) {
	if blur < 1 {
		blur = 1
	}
	for y := 0; y < c.h; y++ {
		for x := 0; x < c.w; x++ {
			dist := d(float64(x)+0.5, float64(y)+0.5)
			cov := clamp(0.5-dist/blur, 0, 1)
			if blur > 1 {
				cov = cov * cov * (3 - 2*cov)
			}
			if cov <= 0 {
				continue
			}
			a := float64(col.A) * cov
			i := (y*c.w + x) * 4
			c.px[i+0] = float64(col.R)*a + c.px[i+0]*(1-a)
			c.px[i+1] = float64(col.G)*a + c.px[i+1]*(1-a)
			c.px[i+2] = float64(col.B)*a + c.px[i+2]*(1-a)
			c.px[i+3] = a + c.px[i+3]*(1-a)
		}
	}
}

func (c *canvas) texture() Texture2D.Instance {
	buf := make([]byte, c.w*c.h*4)
	for i := 0; i < c.w*c.h; i++ {
		a := c.px[i*4+3]
		if a > 0 {
			buf[i*4+0] = byte(clamp(c.px[i*4+0]/a, 0, 1)*255 + 0.5)
			buf[i*4+1] = byte(clamp(c.px[i*4+1]/a, 0, 1)*255 + 0.5)
			buf[i*4+2] = byte(clamp(c.px[i*4+2]/a, 0, 1)*255 + 0.5)
		}
		buf[i*4+3] = byte(clamp(a, 0, 1)*255 + 0.5)
	}
	img := Image.CreateFromData(c.w, c.h, false, Image.FormatRgba8, buf)
	return ImageTexture.CreateFromImage(img).AsTexture2D()
}

func clamp(v, lo, hi float64) float64 { return math.Max(lo, math.Min(hi, v)) }

// roundRect is centered on (cx, cy) with half extents (hw, hh) and radius r.
func roundRect(cx, cy, hw, hh, r float64) sdf {
	return func(x, y float64) float64 {
		qx := math.Abs(x-cx) - hw + r
		qy := math.Abs(y-cy) - hh + r
		return math.Hypot(math.Max(qx, 0), math.Max(qy, 0)) + math.Min(math.Max(qx, qy), 0) - r
	}
}

func circle(cx, cy, r float64) sdf {
	return func(x, y float64) float64 { return math.Hypot(x-cx, y-cy) - r }
}

// inset shrinks a shape by w pixels.
func inset(d sdf, w float64) sdf { return func(x, y float64) float64 { return d(x, y) + w } }

// ring is the band of width w just inside the shape's edge.
func ring(d sdf, w float64) sdf {
	return func(x, y float64) float64 { v := d(x, y); return math.Max(v, -(v + w)) }
}

func offset(d sdf, dx, dy float64) sdf { return func(x, y float64) float64 { return d(x-dx, y-dy) } }

// polyline is a stroke of width w through the given points (x0, y0, x1, y1, ...).
func polyline(w float64, pts ...float64) sdf {
	return func(x, y float64) float64 {
		best := math.Inf(1)
		for i := 0; i+3 < len(pts); i += 2 {
			ax, ay, bx, by := pts[i], pts[i+1], pts[i+2], pts[i+3]
			px, py := x-ax, y-ay
			dx, dy := bx-ax, by-ay
			t := clamp((px*dx+py*dy)/(dx*dx+dy*dy), 0, 1)
			best = math.Min(best, math.Hypot(px-dx*t, py-dy*t))
		}
		return best - w/2
	}
}

// canvasFrom starts a canvas from an existing image.
func canvasFrom(img *image.NRGBA) *canvas {
	b := img.Bounds()
	c := newCanvas(b.Dx(), b.Dy())
	for i := 0; i < c.w*c.h; i++ {
		a := float64(img.Pix[i*4+3]) / 255
		c.px[i*4+0] = float64(img.Pix[i*4+0]) / 255 * a
		c.px[i*4+1] = float64(img.Pix[i*4+1]) / 255 * a
		c.px[i*4+2] = float64(img.Pix[i*4+2]) / 255 * a
		c.px[i*4+3] = a
	}
	return c
}

func (c *canvas) nrgba() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, c.w, c.h))
	for i := 0; i < c.w*c.h; i++ {
		a := c.px[i*4+3]
		if a > 0 {
			img.Pix[i*4+0] = byte(clamp(c.px[i*4+0]/a, 0, 1)*255 + 0.5)
			img.Pix[i*4+1] = byte(clamp(c.px[i*4+1]/a, 0, 1)*255 + 0.5)
			img.Pix[i*4+2] = byte(clamp(c.px[i*4+2]/a, 0, 1)*255 + 0.5)
		}
		img.Pix[i*4+3] = byte(clamp(a, 0, 1)*255 + 0.5)
	}
	return img
}
