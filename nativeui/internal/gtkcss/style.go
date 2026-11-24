package gtkcss

import (
	"strconv"
	"strings"
)

// Shadow is one box-shadow layer.
type Shadow struct {
	Inset              bool
	X, Y, Blur, Spread float64
	Color              RGBA
}

// Stop is a gradient color stop; Pos is 0-1, or -1 when unspecified.
type Stop struct {
	Color RGBA
	Pos   float64
}

// Image is a background-image or border-image source.
type Image struct {
	Kind    string // "", "url", "linear-gradient", "icontheme", "color"
	Path    string // url: absolute file path; icontheme: icon name
	Recolor bool   // -gtk-recolor(): a symbolic image, tinted with the node's color
	Angle   float64
	Stops   []Stop
	Color   RGBA
}

// Style is the computed style of a node, lengths in px.
type Style struct {
	Color            RGBA
	BackgroundColor  RGBA
	BackgroundImage  []Image    // first layer on top
	BorderWidth      [4]float64 // top, right, bottom, left (0 when the style is none)
	BorderColor      [4]RGBA
	BorderStyle      [4]string
	Radius           [4]float64 // top-left, top-right, bottom-right, bottom-left
	RadiusPct        [4]float64 // radii given in %, of the box's smaller side
	Padding, Margin  [4]float64
	MinWidth         float64
	MinHeight        float64
	Shadows          []Shadow
	BorderImage      Image
	BorderImageSlice [4]float64
	IconSource       Image
	OutlineWidth     float64
	OutlineOffset    float64
	OutlineStyle     string
	OutlineColor     RGBA
	Opacity          float64
	FontSize         float64
	FontWeight       int
	FontFamily       string
}

// Compute returns the style GTK would give the last node of path, with
// inheritance from its ancestors.
func (s *Sheet) Compute(path []Node) Style {
	parent := Style{Color: RGBA{0, 0, 0, 1}, Opacity: 1, FontSize: 13.333, FontWeight: 400}
	var st Style
	for i := range path {
		st = s.computeOne(path[:i+1], parent)
		parent = st
	}
	return st
}

func (s *Sheet) computeOne(path []Node, parent Style) Style {
	d := s.Declared(path)
	st := Style{
		Color:      parent.Color,
		Opacity:    1,
		FontSize:   parent.FontSize,
		FontWeight: parent.FontWeight,
		FontFamily: parent.FontFamily,
	}
	get := func(name string) (Decl, bool) {
		v, ok := d[name]
		if ok && (v.Value == "inherit" || v.Value == "initial" || v.Value == "unset") {
			return v, false
		}
		return v, ok
	}
	length := func(name string) float64 {
		if v, ok := get(name); ok {
			f, _ := Length(v.Value)
			return f
		}
		return 0
	}
	if v, ok := get("color"); ok {
		if c, ok := s.Color(v.Value, parent.Color); ok {
			st.Color = c
		}
	}
	col := func(name string, def RGBA) RGBA {
		if v, ok := get(name); ok {
			if c, ok := s.Color(v.Value, st.Color); ok {
				return c
			}
		}
		return def
	}
	if v, ok := get("font-size"); ok {
		if f, ok := Length(v.Value); ok {
			st.FontSize = f
		}
	}
	if v, ok := get("font-weight"); ok {
		switch v.Value {
		case "bold":
			st.FontWeight = 700
		case "normal":
			st.FontWeight = 400
		default:
			if n, err := strconv.Atoi(v.Value); err == nil {
				st.FontWeight = n
			}
		}
	}
	if v, ok := get("font-family"); ok {
		st.FontFamily = strings.Trim(v.Value, `"'`)
	}
	st.BackgroundColor = col("background-color", RGBA{})
	if v, ok := get("background-image"); ok {
		for _, layer := range splitTop(v.Value, ',') {
			if img, ok := s.image(layer, v.Base, st.Color); ok {
				st.BackgroundImage = append(st.BackgroundImage, img)
			}
		}
	}
	for i, side := range sides {
		st.BorderStyle[i] = "none"
		if v, ok := get("border-" + side + "-style"); ok {
			st.BorderStyle[i] = v.Value
		}
		if st.BorderStyle[i] != "none" && st.BorderStyle[i] != "hidden" {
			if v, ok := get("border-" + side + "-width"); ok {
				if v.Value == "medium" {
					st.BorderWidth[i] = 3
				} else {
					st.BorderWidth[i], _ = Length(v.Value)
				}
			} else {
				st.BorderWidth[i] = 3 // CSS initial value "medium"
			}
		}
		st.BorderColor[i] = col("border-"+side+"-color", st.Color)
		st.Padding[i] = length("padding-" + side)
		st.Margin[i] = length("margin-" + side)
	}
	for i, corner := range []string{"top-left", "top-right", "bottom-right", "bottom-left"} {
		if v, ok := get("border-" + corner + "-radius"); ok && strings.HasSuffix(v.Value, "%") {
			pct, _ := strconv.ParseFloat(strings.TrimSuffix(v.Value, "%"), 64)
			st.RadiusPct[i] = pct / 100
			continue
		}
		st.Radius[i] = length("border-" + corner + "-radius")
	}
	st.MinWidth = length("min-width")
	st.MinHeight = length("min-height")
	if v, ok := get("box-shadow"); ok && v.Value != "none" {
		for _, layer := range splitTop(v.Value, ',') {
			if sh, ok := s.shadow(layer, st.Color); ok {
				st.Shadows = append(st.Shadows, sh)
			}
		}
	}
	if v, ok := get("border-image-source"); ok {
		st.BorderImage, _ = s.image(v.Value, v.Base, st.Color)
	}
	if v, ok := get("border-image-slice"); ok {
		b := boxValues(strings.ReplaceAll(v.Value, "fill", ""))
		for i := range b {
			st.BorderImageSlice[i], _ = strconv.ParseFloat(strings.TrimSuffix(b[i], "px"), 64)
		}
	}
	if v, ok := get("-gtk-icon-source"); ok {
		st.IconSource, _ = s.image(v.Value, v.Base, st.Color)
	}
	st.OutlineStyle = "none"
	if v, ok := get("outline-style"); ok {
		st.OutlineStyle = v.Value
	}
	// GTK only draws outlines around widgets with visible keyboard focus.
	if st.OutlineStyle != "none" && has(path[len(path)-1].States, "focus") {
		st.OutlineWidth = length("outline-width")
	}
	st.OutlineOffset = length("outline-offset")
	st.OutlineColor = col("outline-color", st.Color)
	if v, ok := get("opacity"); ok {
		st.Opacity, _ = strconv.ParseFloat(v.Value, 64)
	}
	return st
}

func (s *Sheet) shadow(layer string, current RGBA) (Shadow, bool) {
	var sh Shadow
	var lengths []float64
	color := current
	for _, f := range splitTop(layer, ' ') {
		if f == "inset" {
			sh.Inset = true
		} else if l, ok := Length(f); ok {
			lengths = append(lengths, l)
		} else if c, ok := s.Color(f, current); ok {
			color = c
		}
	}
	if len(lengths) < 2 {
		return sh, false
	}
	sh.X, sh.Y = lengths[0], lengths[1]
	if len(lengths) > 2 {
		sh.Blur = lengths[2]
	}
	if len(lengths) > 3 {
		sh.Spread = lengths[3]
	}
	sh.Color = color
	return sh, true
}

func (s *Sheet) image(v, base string, current RGBA) (Image, bool) {
	v = strings.TrimSpace(v)
	if v == "" || v == "none" {
		return Image{}, false
	}
	name, args, ok := call(v)
	if !ok {
		return Image{}, false
	}
	switch name {
	case "url":
		return Image{Kind: "url", Path: resolve(base, urlArg(v))}, true
	case "-gtk-scaled":
		if len(args) > 0 {
			return s.image(args[0], base, current) // the 1x image
		}
	case "-gtk-icontheme":
		if len(args) > 0 {
			return Image{Kind: "icontheme", Path: strings.Trim(args[0], `"' `)}, true
		}
	case "image":
		// image(fallback images..., color): the first one we can load, else the color.
		for _, a := range args {
			a = strings.TrimSpace(a)
			if isImage(a) || strings.HasPrefix(a, "-gtk-recolor(") {
				if img, ok := s.image(a, base, current); ok && (img.Kind != "url" || LoadImage(img.Path) != nil) {
					return img, true
				}
				continue
			}
			if c, ok := s.Color(a, current); ok {
				return Image{Kind: "color", Color: c}, true
			}
		}
	case "-gtk-recolor":
		if len(args) > 0 {
			img, ok := s.image(args[0], base, current)
			img.Recolor = true
			return img, ok
		}
	case "linear-gradient":
		img := Image{Kind: "linear-gradient", Angle: 180} // "to bottom"
		for i, a := range args {
			a = strings.TrimSpace(a)
			if i == 0 {
				if ang, ok := gradientAngle(a); ok {
					img.Angle = ang
					continue
				}
			}
			stop := Stop{Pos: -1}
			parts := splitTop(a, ' ')
			if len(parts) == 0 {
				continue
			}
			c, ok := s.Color(parts[0], current)
			if !ok {
				continue
			}
			stop.Color = c
			if len(parts) > 1 {
				if strings.HasSuffix(parts[1], "%") {
					f, _ := strconv.ParseFloat(strings.TrimSuffix(parts[1], "%"), 64)
					stop.Pos = f / 100
				}
			}
			img.Stops = append(img.Stops, stop)
		}
		if len(img.Stops) > 0 {
			fillStops(img.Stops)
			return img, true
		}
	}
	return Image{}, false
}

func gradientAngle(a string) (float64, bool) {
	switch a {
	case "to bottom":
		return 180, true
	case "to top":
		return 0, true
	case "to right":
		return 90, true
	case "to left":
		return 270, true
	}
	if strings.HasSuffix(a, "deg") {
		f, err := strconv.ParseFloat(strings.TrimSuffix(a, "deg"), 64)
		return f, err == nil
	}
	return 0, false
}

// fillStops gives unpositioned stops positions spaced evenly between their neighbours.
func fillStops(stops []Stop) {
	if stops[0].Pos < 0 {
		stops[0].Pos = 0
	}
	if last := len(stops) - 1; stops[last].Pos < 0 {
		stops[last].Pos = 1
	}
	for i := 1; i < len(stops); i++ {
		if stops[i].Pos >= 0 {
			continue
		}
		j := i
		for stops[j].Pos < 0 {
			j++
		}
		a, b := stops[i-1].Pos, stops[j].Pos
		for k := i; k < j; k++ {
			stops[k].Pos = a + (b-a)*float64(k-i+1)/float64(j-i+1)
		}
	}
}
