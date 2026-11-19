package gtkcss

import (
	"math"
	"strconv"
	"strings"
)

// RGBA is a non-premultiplied color, components 0-1.
type RGBA struct{ R, G, B, A float64 }

var named = map[string]RGBA{
	"transparent": {0, 0, 0, 0}, "black": {0, 0, 0, 1}, "white": {1, 1, 1, 1},
	"red": {1, 0, 0, 1}, "green": {0, 128.0 / 255, 0, 1}, "blue": {0, 0, 1, 1}, "lime": {0, 1, 0, 1},
	"yellow": {1, 1, 0, 1}, "cyan": {0, 1, 1, 1}, "aqua": {0, 1, 1, 1}, "magenta": {1, 0, 1, 1}, "fuchsia": {1, 0, 1, 1},
	"gray": {128.0 / 255, 128.0 / 255, 128.0 / 255, 1}, "grey": {128.0 / 255, 128.0 / 255, 128.0 / 255, 1},
	"silver": {192.0 / 255, 192.0 / 255, 192.0 / 255, 1}, "navy": {0, 0, 128.0 / 255, 1},
	"maroon": {128.0 / 255, 0, 0, 1}, "olive": {128.0 / 255, 128.0 / 255, 0, 1}, "purple": {128.0 / 255, 0, 128.0 / 255, 1},
	"teal": {0, 128.0 / 255, 128.0 / 255, 1}, "orange": {1, 165.0 / 255, 0, 1},
	"darkgray": {169.0 / 255, 169.0 / 255, 169.0 / 255, 1}, "lightgray": {211.0 / 255, 211.0 / 255, 211.0 / 255, 1},
}

// splitTop splits s on sep at parenthesis depth 0 (sep ' ' splits on any whitespace).
func splitTop(s string, sep byte) []string {
	var out []string
	depth := 0
	start := 0
	var quote byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case quote != 0:
			if c == quote {
				quote = 0
			}
		case c == '"' || c == '\'':
			quote = c
		case c == '(':
			depth++
		case c == ')':
			depth--
		case depth == 0 && (c == sep || sep == ' ' && (c == '\t' || c == '\n')):
			if part := strings.TrimSpace(s[start:i]); part != "" || sep != ' ' {
				out = append(out, part)
			}
			start = i + 1
		}
	}
	if part := strings.TrimSpace(s[start:]); part != "" || (sep != ' ' && len(out) > 0) {
		out = append(out, part)
	}
	return out
}

// call splits "name(args)" into name and comma separated args.
func call(s string) (string, []string, bool) {
	open := strings.IndexByte(s, '(')
	if open <= 0 || !strings.HasSuffix(s, ")") {
		return "", nil, false
	}
	return strings.TrimSpace(s[:open]), splitTop(s[open+1:len(s)-1], ','), true
}

// Color evaluates a GTK color expression.
func (s *Sheet) Color(expr string, current RGBA) (RGBA, bool) {
	return s.color(strings.TrimSpace(expr), current, 0)
}

func (s *Sheet) color(e string, current RGBA, depth int) (RGBA, bool) {
	if depth > 32 {
		return RGBA{}, false
	}
	switch {
	case e == "":
		return RGBA{}, false
	case e == "currentColor" || e == "currentcolor":
		return current, true
	case e[0] == '@':
		def, ok := s.Colors[e[1:]]
		if !ok {
			return RGBA{}, false
		}
		return s.color(strings.TrimSpace(def), current, depth+1)
	case e[0] == '#':
		return hexColor(e[1:])
	}
	if c, ok := named[strings.ToLower(e)]; ok {
		return c, true
	}
	name, args, ok := call(e)
	if !ok {
		return RGBA{}, false
	}
	num := func(i int) float64 {
		if i >= len(args) {
			return 0
		}
		v := strings.TrimSpace(args[i])
		if strings.HasSuffix(v, "%") {
			f, _ := strconv.ParseFloat(strings.TrimSuffix(v, "%"), 64)
			return f / 100
		}
		f, _ := strconv.ParseFloat(v, 64)
		return f
	}
	sub := func(i int) (RGBA, bool) {
		if i >= len(args) {
			return RGBA{}, false
		}
		return s.color(strings.TrimSpace(args[i]), current, depth+1)
	}
	channel := func(i int) float64 {
		v := strings.TrimSpace(args[i])
		if strings.HasSuffix(v, "%") {
			return num(i)
		}
		return num(i) / 255
	}
	switch name {
	case "rgb", "rgba":
		if len(args) < 3 {
			return RGBA{}, false
		}
		c := RGBA{channel(0), channel(1), channel(2), 1}
		if len(args) > 3 {
			c.A = num(3)
		}
		return c, true
	case "shade":
		c, ok := sub(0)
		return shade(c, num(1)), ok
	case "lighter":
		c, ok := sub(0)
		return shade(c, 1.3), ok
	case "darker":
		c, ok := sub(0)
		return shade(c, 0.7), ok
	case "alpha":
		c, ok := sub(0)
		c.A = clamp01(c.A * num(1))
		return c, ok
	case "mix":
		a, ok1 := sub(0)
		b, ok2 := sub(1)
		f := num(2)
		return RGBA{a.R + (b.R-a.R)*f, a.G + (b.G-a.G)*f, a.B + (b.B-a.B)*f, a.A + (b.A-a.A)*f}, ok1 && ok2
	}
	return RGBA{}, false
}

func hexColor(h string) (RGBA, bool) {
	parse := func(s string) float64 {
		v, _ := strconv.ParseUint(s, 16, 8)
		return float64(v) / 255
	}
	switch len(h) {
	case 3, 4:
		var c RGBA
		c.A = 1
		dst := []*float64{&c.R, &c.G, &c.B, &c.A}
		for i := range h {
			*dst[i] = parse(strings.Repeat(h[i:i+1], 2))
		}
		return c, true
	case 6, 8:
		var c RGBA
		c.A = 1
		dst := []*float64{&c.R, &c.G, &c.B, &c.A}
		for i := 0; i < len(h); i += 2 {
			*dst[i/2] = parse(h[i : i+2])
		}
		return c, true
	}
	return RGBA{}, false
}

func clamp01(v float64) float64 { return math.Max(0, math.Min(1, v)) }

// shade is GTK's: scale lightness and saturation in HLS.
func shade(c RGBA, f float64) RGBA {
	h, l, s := rgbToHLS(c.R, c.G, c.B)
	l = clamp01(l * f)
	s = clamp01(s * f)
	r, g, b := hlsToRGB(h, l, s)
	return RGBA{r, g, b, c.A}
}

func rgbToHLS(r, g, b float64) (h, l, s float64) {
	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	l = (max + min) / 2
	if max == min {
		return 0, l, 0
	}
	d := max - min
	if l <= 0.5 {
		s = d / (max + min)
	} else {
		s = d / (2 - max - min)
	}
	switch max {
	case r:
		h = (g - b) / d
	case g:
		h = 2 + (b-r)/d
	default:
		h = 4 + (r-g)/d
	}
	h *= 60
	if h < 0 {
		h += 360
	}
	return h, l, s
}

func hlsToRGB(h, l, s float64) (r, g, b float64) {
	if s == 0 {
		return l, l, l
	}
	var m2 float64
	if l <= 0.5 {
		m2 = l * (1 + s)
	} else {
		m2 = l + s - l*s
	}
	m1 := 2*l - m2
	f := func(hue float64) float64 {
		for hue > 360 {
			hue -= 360
		}
		for hue < 0 {
			hue += 360
		}
		switch {
		case hue < 60:
			return m1 + (m2-m1)*hue/60
		case hue < 180:
			return m2
		case hue < 240:
			return m1 + (m2-m1)*(240-hue)/60
		}
		return m1
	}
	return f(h + 120), f(h), f(h - 120)
}

// Length parses a CSS length in px (1em = 16px, 1pt = 4/3 px).
func Length(s string) (float64, bool) {
	s = strings.TrimSpace(s)
	unit := 1.0
	for _, u := range []struct {
		suffix string
		scale  float64
	}{{"px", 1}, {"pt", 4.0 / 3}, {"em", 16}, {"rem", 16}, {"ex", 8}} {
		if strings.HasSuffix(s, u.suffix) {
			s, unit = strings.TrimSuffix(s, u.suffix), u.scale
			break
		}
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, false
	}
	return f * unit, true
}

var borderStyles = map[string]bool{"none": true, "hidden": true, "solid": true, "dotted": true, "dashed": true,
	"double": true, "groove": true, "ridge": true, "inset": true, "outset": true}

var sides = [4]string{"top", "right", "bottom", "left"}

// boxValues expands 1-4 values to top, right, bottom, left.
func boxValues(v string) [4]string {
	f := splitTop(v, ' ')
	switch len(f) {
	case 0:
		return [4]string{}
	case 1:
		return [4]string{f[0], f[0], f[0], f[0]}
	case 2:
		return [4]string{f[0], f[1], f[0], f[1]}
	case 3:
		return [4]string{f[0], f[1], f[2], f[1]}
	}
	return [4]string{f[0], f[1], f[2], f[3]}
}

// expand turns a shorthand declaration into longhands.
func expand(d Decl) []Decl {
	mk := func(name, value string) Decl { return Decl{Name: name, Value: value, Base: d.Base} }
	var out []Decl
	switch d.Name {
	case "padding", "margin":
		v := boxValues(d.Value)
		for i, side := range sides {
			out = append(out, mk(d.Name+"-"+side, v[i]))
		}
	case "border-width", "border-color", "border-style":
		kind := strings.TrimPrefix(d.Name, "border-")
		v := boxValues(d.Value)
		for i, side := range sides {
			out = append(out, mk("border-"+side+"-"+kind, v[i]))
		}
	case "border-radius":
		// only the horizontal radii (before any "/") are used.
		v := d.Value
		if i := scanTo(v, "/"); i >= 0 {
			v = v[:i]
		}
		r := boxValues(v)
		for i, corner := range []string{"top-left", "top-right", "bottom-right", "bottom-left"} {
			out = append(out, mk("border-"+corner+"-radius", r[i]))
		}
	case "border", "border-top", "border-right", "border-bottom", "border-left":
		width, style, color := "0", "none", "currentColor"
		if d.Value != "none" && d.Value != "0" {
			width, style = "medium", "none"
			for _, f := range splitTop(d.Value, ' ') {
				if _, ok := Length(f); ok {
					width = f
				} else if borderStyles[f] {
					style = f
				} else {
					color = f
				}
			}
		}
		list := sides[:]
		if d.Name != "border" {
			list = []string{strings.TrimPrefix(d.Name, "border-")}
		}
		for _, side := range list {
			out = append(out, mk("border-"+side+"-width", width), mk("border-"+side+"-style", style), mk("border-"+side+"-color", color))
		}
		if d.Name == "border" {
			out = append(out, mk("border-image-source", "none"))
		}
	case "outline":
		width, style, color := "0", "none", "currentColor"
		if d.Value != "none" && d.Value != "0" {
			for _, f := range splitTop(d.Value, ' ') {
				if _, ok := Length(f); ok {
					width = f
				} else if borderStyles[f] || f == "auto" {
					style = f
				} else {
					color = f
				}
			}
		}
		out = append(out, mk("outline-width", width), mk("outline-style", style), mk("outline-color", color))
	case "background":
		color, image := "transparent", "none"
		for _, f := range splitTop(d.Value, ' ') {
			if f == "none" {
				continue
			}
			if isImage(f) {
				image = f
			} else if _, ok := Length(f); !ok && !strings.Contains("repeat no-repeat center top bottom left right cover contain", f) {
				color = f
			}
		}
		out = append(out, mk("background-color", color), mk("background-image", image))
	case "border-image":
		// border-image: source [slice] [/ width] [repeat]
		main := d.Value
		if i := scanTo(main, "/"); i >= 0 { // the "/" before border widths, not one inside url()
			main = main[:i]
		}
		var slices []string
		source := "none"
		for _, f := range splitTop(main, ' ') {
			if isImage(f) {
				source = f
			} else if _, err := strconv.ParseFloat(strings.TrimSuffix(f, "%"), 64); err == nil {
				slices = append(slices, f)
			}
		}
		out = append(out, mk("border-image-source", source))
		if len(slices) > 0 {
			out = append(out, mk("border-image-slice", strings.Join(slices, " ")))
		}
	default:
		out = append(out, d)
	}
	return out
}

func isImage(s string) bool {
	for _, prefix := range []string{"url(", "linear-gradient(", "radial-gradient(", "repeating-linear-gradient(", "-gtk-scaled(", "-gtk-icontheme(", "image(", "cross-fade(", "-gtk-gradient("} {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}
