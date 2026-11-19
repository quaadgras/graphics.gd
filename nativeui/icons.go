package nativeui

import (
	"math"

	"graphics.gd/classdb/Texture2D"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

func faded(c Color.RGBA, disabled bool) Color.RGBA {
	if disabled {
		c.A *= 0.4
	}
	return c
}

// checkIcon draws a check box or radio button.
func (b *builder) checkIcon(checked, radio, disabled bool) Texture2D.Instance {
	c := b.s.check
	k := b.scale
	size := math.Ceil(c.size * k)
	pad := math.Ceil(2 * k) // room for the soft edge
	cv := newCanvas(int(size+2*pad), int(size+2*pad))
	mid := pad + size/2
	var shape sdf
	if radio {
		shape = circle(mid, mid, size/2)
	} else {
		shape = roundRect(mid, mid, size/2, size/2, c.radius*k)
	}
	if !checked {
		cv.fill(shape, faded(c.bg, disabled), 0)
		cv.fill(ring(shape, c.borderW*k), faded(c.border, disabled), 0)
		return cv.texture()
	}
	if radio && c.radioRing {
		if c.radioGap.A > 0 {
			cv.fill(shape, faded(c.radioGap, disabled), 0)
		}
		cv.fill(ring(shape, c.borderW*k), faded(b.s.accent, disabled), 0)
		cv.fill(circle(mid, mid, c.radioDot*size), faded(b.s.accent, disabled), 0)
		return cv.texture()
	}
	onFill, mark := b.s.accent, b.s.onAccent
	if c.onFill.A > 0 {
		onFill = c.onFill
	}
	if c.mark.A > 0 {
		mark = c.mark
	}
	cv.fill(shape, faded(onFill, disabled), 0)
	if c.onBorder.A > 0 {
		cv.fill(ring(shape, c.borderW*k), faded(c.onBorder, disabled), 0)
	}
	mark = faded(mark, disabled)
	if radio {
		cv.fill(circle(mid, mid, c.radioDot*size), mark, 0)
		return cv.texture()
	}
	u := size / 14 // mark is designed on a 14px box
	o := pad
	cv.fill(polyline(c.markW*k, o+3.5*u, o+7.2*u, o+6*u, o+9.8*u, o+10.6*u, o+4.4*u), mark, 0)
	return cv.texture()
}

// switchIcon draws a toggle switch (CheckButton).
func (b *builder) switchIcon(on, disabled bool) Texture2D.Instance {
	sw := b.s.toggle
	k := b.scale
	w, h := math.Ceil(sw.w*k), math.Ceil(sw.h*k)
	// Room for the shadow, and for a knob bigger than the track (One UI).
	pad := math.Ceil(max(3, math.Max(sw.knobOffR, sw.knobOnR)-sw.h/2+3) * k)
	cv := newCanvas(int(w+2*pad), int(h+2*pad))
	cx, cy := pad+w/2, pad+h/2
	track := roundRect(cx, cy, w/2, h/2, h/2)
	knobR := sw.knobOffR * k
	knob := sw.knobOff
	if on {
		onColor := sw.on
		if onColor.A == 0 {
			onColor = b.s.accent
		}
		cv.fill(track, faded(onColor, disabled), 0)
		knobR, knob = sw.knobOnR*k, sw.knobOn
	} else {
		cv.fill(track, faded(sw.off, disabled), 0)
		if sw.offBorderW > 0 {
			cv.fill(ring(track, sw.offBorderW*k), faded(sw.offBorder, disabled), 0)
		}
	}
	// The knob sits in the track's end cap, as far from the edge as it is
	// from the top; a knob wider than it is tall is a capsule (macOS 26+).
	gap := h/2 - knobR
	hw := math.Max(knobR, sw.knobW*k/2)
	kx := pad + gap + hw
	if on {
		kx = pad + w - gap - hw
	}
	shape := roundRect(kx, cy, hw, knobR, knobR)
	if sw.knobShadow {
		cv.fill(offset(shape, 0, 0.75*k), faded(black(0.25), disabled), 2*k)
	}
	cv.fill(shape, faded(knob, disabled), 0)
	return cv.texture()
}

// thumbIcon draws a slider grabber.
func (b *builder) thumbIcon(disabled bool) Texture2D.Instance {
	th := b.s.thumb
	k := b.scale
	if th.gapW > 0 {
		// A bar handle that cuts a gap out of the track on either side.
		bar, gap, h := th.w*k, th.gapW*k, math.Ceil(th.size*k)
		track := (b.s.sliderTrack.pad[1] + b.s.sliderTrack.pad[3]) * k
		w := math.Ceil(bar + 2*gap)
		cv := newCanvas(int(w), int(h))
		cv.fill(roundRect(w/2, h/2, w/2, track/2, 0), th.gap, 0)
		cv.fill(roundRect(w/2, h/2, bar/2, h/2, bar/2), faded(th.fill, disabled), 0)
		return cv.texture()
	}
	size := math.Ceil(th.size * k)
	width := size
	if th.w > 0 {
		width = math.Ceil(th.w * k)
	}
	pad := math.Ceil(3 * k)
	cv := newCanvas(int(width+2*pad), int(size+2*pad))
	mid, midX := pad+size/2, pad+width/2
	r, hw := size/2, width/2
	if th.shadow {
		cv.fill(offset(roundRect(midX, mid, hw-0.5*k, r-0.5*k, r-0.5*k), 0, 1*k), faded(black(0.25), disabled), 2.5*k)
	}
	if th.radius > 0 {
		outer := roundRect(midX, mid, hw, r, th.radius*k)
		cv.fill(outer, faded(th.border, disabled), 0)
		inner := inset(outer, 1*k)
		steps := int(size)
		for i := 0; i < steps; i++ { // vertical gradient top -> fill, a band at a time
			t := float64(i) / float64(max(1, steps-1))
			col := Color.RGBA{R: th.top.R + (th.fill.R-th.top.R)*Float.X(t), G: th.top.G + (th.fill.G-th.top.G)*Float.X(t), B: th.top.B + (th.fill.B-th.top.B)*Float.X(t), A: 1}
			y0 := pad + float64(i)
			cv.fill(func(x, y float64) float64 { return math.Max(inner(x, y), math.Max(y0-y, y-(y0+1))) }, faded(col, disabled), 0)
		}
	} else {
		cv.fill(roundRect(midX, mid, hw, r, r), faded(th.border, disabled), 0)
		bw := max(th.borderW, 1) * k
		cv.fill(roundRect(midX, mid, hw-bw, r-bw, r-bw), faded(th.fill, disabled), 0)
	}
	if th.dot > 0 {
		cv.fill(circle(midX, mid, th.dot*k), faded(b.s.accent, disabled), 0)
	}
	return cv.texture()
}

// arrowIcon is the OptionButton's indicator: a chevron, or on macOS up/down
// chevrons in an accent badge.
func (b *builder) arrowIcon() Texture2D.Instance {
	k := b.scale
	if b.s.arrowBadge {
		size := math.Ceil(16 * k)
		cv := newCanvas(int(size), int(size))
		m := size / 2
		cv.fill(roundRect(m, m, m, m, 4*k), b.s.accent, 0)
		u := size / 16
		cv.fill(polyline(1.6*k, 5*u, 6.8*u, 8*u, 4*u, 11*u, 6.8*u), b.s.onAccent, 0)
		cv.fill(polyline(1.6*k, 5*u, 9.2*u, 8*u, 12*u, 11*u, 9.2*u), b.s.onAccent, 0)
		return cv.texture()
	}
	size := math.Ceil(16 * k)
	cv := newCanvas(int(size), int(size))
	u := size / 16
	if b.s.arrowUpDown {
		cv.fill(polyline(1.3*k, 5.5*u, 6.5*u, 8*u, 4*u, 10.5*u, 6.5*u), b.s.arrowColor, 0)
		cv.fill(polyline(1.3*k, 5.5*u, 9.5*u, 8*u, 12*u, 10.5*u, 9.5*u), b.s.arrowColor, 0)
		return cv.texture()
	}
	cv.fill(polyline(1.3*k, 3.5*u, 6*u, 8*u, 10.5*u, 12.5*u, 6*u), b.s.arrowColor, 0)
	return cv.texture()
}

func (b *builder) chevronIcon(up bool) Texture2D.Instance {
	k := b.scale
	w, h := math.Ceil(14*k), math.Ceil(10*k)
	cv := newCanvas(int(w), int(h))
	u := w / 14
	if up {
		cv.fill(polyline(1.3*k, 3.5*u, 7*u, 7*u, 3.5*u, 10.5*u, 7*u), white(1), 0)
	} else {
		cv.fill(polyline(1.3*k, 3.5*u, 3*u, 7*u, 6.5*u, 10.5*u, 3*u), white(1), 0)
	}
	return cv.texture()
}

func (b *builder) updownIcon() Texture2D.Instance {
	k := b.scale
	w, h := math.Ceil(14*k), math.Ceil(20*k)
	cv := newCanvas(int(w), int(h))
	u := w / 14
	cv.fill(polyline(1.3*k, 3.5*u, 8*u, 7*u, 4.5*u, 10.5*u, 8*u), b.s.arrowColor, 0)
	cv.fill(polyline(1.3*k, 3.5*u, 12*u, 7*u, 15.5*u, 10.5*u, 12*u), b.s.arrowColor, 0)
	return cv.texture()
}

func (b *builder) menuCheckIcon() Texture2D.Instance {
	k := b.scale
	size := math.Ceil(14 * k)
	cv := newCanvas(int(size), int(size))
	u := size / 14
	cv.fill(polyline(1.6*k, 2.5*u, 7.5*u, 5.5*u, 10.5*u, 11.5*u, 3.5*u), b.s.text, 0)
	return cv.texture()
}

func (b *builder) blankIcon() Texture2D.Instance {
	size := math.Ceil(14 * b.scale)
	return newCanvas(int(size), int(size)).texture()
}
