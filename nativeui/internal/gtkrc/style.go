package gtkrc

import (
	"math"
	"path"
	"sort"
	"strconv"
	"strings"
)

// color evaluates a gtkrc color expression in the context of a style.
func (t *Theme) color(v string, st *Style) (RGBA, bool) {
	return t.colorDepth(strings.TrimSpace(v), st, 0)
}

func (t *Theme) colorDepth(v string, st *Style, depth int) (RGBA, bool) {
	if depth > 16 || v == "" {
		return RGBA{}, false
	}
	if v[0] == '"' {
		v = unquote(v)
	}
	switch {
	case v[0] == '@':
		name := v[1:]
		if st != nil {
			if c, ok := st.Colors[name]; ok {
				return t.colorDepth(c, st, depth+1)
			}
		}
		if c, ok := t.Scheme[name]; ok {
			return t.colorDepth(c, st, depth+1)
		}
		return RGBA{}, false
	case v[0] == '#':
		return hexColor(v[1:])
	case v[0] == '{':
		parts := strings.Split(strings.Trim(v, "{}"), ",")
		if len(parts) != 3 {
			return RGBA{}, false
		}
		var c [3]float64
		for i, p := range parts {
			c[i], _ = strconv.ParseFloat(strings.TrimSpace(p), 64)
			if c[i] > 1 { // integer components
				c[i] /= 65535
			}
		}
		return RGBA{c[0], c[1], c[2], 1}, true
	}
	if open := strings.IndexByte(v, '('); open > 0 && strings.HasSuffix(v, ")") {
		name := v[:open]
		args := splitArgs(v[open+1 : len(v)-1])
		num := func(i int) float64 {
			f, _ := strconv.ParseFloat(strings.TrimSpace(args[i]), 64)
			return f
		}
		sub := func(i int) (RGBA, bool) { return t.colorDepth(strings.TrimSpace(args[i]), st, depth+1) }
		switch {
		case name == "shade" && len(args) == 2:
			c, ok := sub(1)
			return shade(c, num(0)), ok
		case name == "lighter" && len(args) == 1:
			c, ok := sub(0)
			return shade(c, 1.3), ok
		case name == "darker" && len(args) == 1:
			c, ok := sub(0)
			return shade(c, 0.7), ok
		case name == "mix" && len(args) == 3:
			a, ok1 := sub(1)
			b, ok2 := sub(2)
			f := num(0)
			// gtkrc's mix(f, a, b) = a*f + b*(1-f)
			return RGBA{a.R*f + b.R*(1-f), a.G*f + b.G*(1-f), a.B*f + b.B*(1-f), 1}, ok1 && ok2
		}
		return RGBA{}, false
	}
	if c, ok := named[strings.ToLower(v)]; ok {
		return c, true
	}
	return RGBA{}, false
}

var named = map[string]RGBA{"white": {1, 1, 1, 1}, "black": {0, 0, 0, 1}, "gray": {0.745, 0.745, 0.745, 1}, "grey": {0.745, 0.745, 0.745, 1},
	"red": {1, 0, 0, 1}, "green": {0, 1, 0, 1}, "blue": {0, 0, 1, 1}, "navy": {0, 0, 0.5, 1}}

func splitArgs(s string) []string {
	var out []string
	depth, start := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{':
			depth++
		case ')', '}':
			depth--
		case ',':
			if depth == 0 {
				out = append(out, s[start:i])
				start = i + 1
			}
		}
	}
	return append(out, s[start:])
}

func hexColor(h string) (RGBA, bool) {
	parse := func(s string) float64 {
		v, _ := strconv.ParseUint(s, 16, 16)
		return float64(v) / float64(uint64(1)<<(4*len(s))-1)
	}
	switch len(h) {
	case 3:
		return RGBA{parse(h[0:1]), parse(h[1:2]), parse(h[2:3]), 1}, true
	case 6:
		return RGBA{parse(h[0:2]), parse(h[2:4]), parse(h[4:6]), 1}, true
	case 12:
		return RGBA{parse(h[0:4]), parse(h[4:8]), parse(h[8:12]), 1}, true
	}
	return RGBA{}, false
}

func clamp01(v float64) float64 { return math.Max(0, math.Min(1, v)) }

// shade is GTK 2's gtk_style_shade: scale lightness and saturation in HLS.
func shade(c RGBA, f float64) RGBA {
	h, l, s := rgbToHLS(c.R, c.G, c.B)
	l, s = clamp01(l*f), clamp01(s*f)
	r, g, b := hlsToRGB(h, l, s)
	return RGBA{r, g, b, c.A}
}

func rgbToHLS(r, g, b float64) (h, l, s float64) {
	max, min := math.Max(r, math.Max(g, b)), math.Min(r, math.Min(g, b))
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
	return
}

func hlsToRGB(h, l, s float64) (r, g, b float64) {
	if s == 0 {
		return l, l, l
	}
	m2 := l + s - l*s
	if l <= 0.5 {
		m2 = l * (1 + s)
	}
	m1 := 2*l - m2
	f := func(hue float64) float64 {
		hue = math.Mod(hue+360, 360)
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

// Widget describes a widget for style matching.
type Widget struct {
	Class []string // the widget's class and its ancestors: GtkButton, GtkBin, GtkContainer, GtkWidget, GtkObject
	Path  []Widget // its ancestors in the widget tree, outermost first (only Class[0] of each is used, plus Name)
	Name  string   // widget name, defaults to the class name
}

func (w Widget) name() string {
	if w.Name != "" {
		return w.Name
	}
	return w.Class[0]
}

// Hierarchies of the classes nativeui asks about.
var hierarchy = map[string][]string{
	"GtkWidget":            {"GtkWidget", "GtkObject"},
	"GtkWindow":            {"GtkWindow", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkVBox":              {"GtkVBox", "GtkBox", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkHBox":              {"GtkHBox", "GtkBox", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkLabel":             {"GtkLabel", "GtkMisc", "GtkWidget", "GtkObject"},
	"GtkButton":            {"GtkButton", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkToggleButton":      {"GtkToggleButton", "GtkButton", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkCheckButton":       {"GtkCheckButton", "GtkToggleButton", "GtkButton", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkRadioButton":       {"GtkRadioButton", "GtkCheckButton", "GtkToggleButton", "GtkButton", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkEntry":             {"GtkEntry", "GtkWidget", "GtkObject"},
	"GtkSpinButton":        {"GtkSpinButton", "GtkEntry", "GtkWidget", "GtkObject"},
	"GtkComboBox":          {"GtkComboBox", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkHScale":            {"GtkHScale", "GtkScale", "GtkRange", "GtkWidget", "GtkObject"},
	"GtkProgressBar":       {"GtkProgressBar", "GtkProgress", "GtkWidget", "GtkObject"},
	"GtkNotebook":          {"GtkNotebook", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkMenu":              {"GtkMenu", "GtkMenuShell", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkMenuItem":          {"GtkMenuItem", "GtkItem", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkCheckMenuItem":     {"GtkCheckMenuItem", "GtkMenuItem", "GtkItem", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkSeparatorMenuItem": {"GtkSeparatorMenuItem", "GtkMenuItem", "GtkItem", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkHSeparator":        {"GtkHSeparator", "GtkSeparator", "GtkWidget", "GtkObject"},
	"GtkScrolledWindow":    {"GtkScrolledWindow", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkTextView":          {"GtkTextView", "GtkContainer", "GtkWidget", "GtkObject"},
	"GtkFrame":             {"GtkFrame", "GtkBin", "GtkContainer", "GtkWidget", "GtkObject"},
}

// W builds a Widget from a dotted class path, e.g. "GtkWindow.GtkVBox.GtkButton".
func W(classPath string) Widget {
	parts := strings.Split(classPath, ".")
	mk := func(class string) Widget {
		if h, ok := hierarchy[class]; ok {
			return Widget{Class: h}
		}
		return Widget{Class: []string{class, "GtkWidget", "GtkObject"}}
	}
	w := mk(parts[len(parts)-1])
	for _, p := range parts[:len(parts)-1] {
		w.Path = append(w.Path, mk(p))
	}
	return w
}

// Resolve merges the styles bound to a widget, as gtk_rc_get_style does:
// class bindings (base classes first), then widget_class, then widget, each
// later binding overriding earlier ones of the same priority.
func (t *Theme) Resolve(w Widget) *Style {
	type hit struct {
		b    binding
		rank int
	}
	var hits []hit
	full := append(append([]Widget{}, w.Path...), w)
	for _, b := range t.bindings {
		switch b.kind {
		case 0:
			for depth := len(w.Class) - 1; depth >= 0; depth-- {
				if ok, _ := path.Match(b.pattern, w.Class[depth]); ok {
					hits = append(hits, hit{b, len(w.Class) - depth})
					break
				}
			}
		case 1:
			if matchPath(b.pattern, full, true) {
				hits = append(hits, hit{b, 0})
			}
		case 2:
			if matchPath(b.pattern, full, false) {
				hits = append(hits, hit{b, 0})
			}
		}
	}
	sort.SliceStable(hits, func(i, j int) bool {
		a, b := hits[i], hits[j]
		if a.b.priority != b.b.priority {
			return a.b.priority < b.b.priority
		}
		if a.b.kind != b.b.kind {
			return a.b.kind < b.b.kind
		}
		if a.rank != b.rank {
			return a.rank < b.rank
		}
		return a.b.order < b.b.order
	})
	out := &Style{Name: "merged", Props: map[string]string{}, EngineParams: map[string]string{}, Colors: map[string]string{}}
	out.XThickness, out.YThickness = 2, 2
	out.scheme = t.Scheme
	for _, h := range hits {
		if st, ok := t.Styles[h.b.style]; ok {
			out.merge(st)
		}
	}
	out.defaults()
	return out
}

func (s *Style) merge(o *Style) {
	for which, dst := range []*[numStates]RGBA{&s.Fg, &s.Bg, &s.Base, &s.Text} {
		src := []*[numStates]RGBA{&o.Fg, &o.Bg, &o.Base, &o.Text}[which]
		for st := 0; st < numStates; st++ {
			if o.set[which][st] {
				dst[st] = src[st]
				s.set[which][st] = true
			}
		}
	}
	if o.thicknessSet[0] {
		s.XThickness, s.thicknessSet[0] = o.XThickness, true
	}
	if o.thicknessSet[1] {
		s.YThickness, s.thicknessSet[1] = o.YThickness, true
	}
	if o.Font != "" {
		s.Font = o.Font
	}
	for k, v := range o.Props {
		s.Props[k] = v
	}
	for k, v := range o.Colors {
		s.Colors[k] = v
	}
	if o.Engine != "" {
		if o.Engine == s.Engine {
			// the same engine merges: the later style's options override
			// key by key, and its pixmap images come before the earlier ones
			for k, v := range o.EngineParams {
				s.EngineParams[k] = v
			}
			s.Images = append(append([]Image(nil), o.Images...), s.Images...)
		} else {
			s.Engine = o.Engine
			s.Images = o.Images
			s.EngineParams = copyMap(o.EngineParams)
		}
	}
}

// defaults fills unset colors with GTK 2's built-in defaults.
func (s *Style) defaults() {
	def := [4][numStates]RGBA{
		{{0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}, {1, 1, 1, 1}, {0.459, 0.459, 0.459, 1}},                                                // fg
		{{0.863, 0.855, 0.835, 1}, {0.765, 0.757, 0.737, 1}, {0.933, 0.922, 0.906, 1}, {0.29, 0.565, 0.851, 1}, {0.863, 0.855, 0.835, 1}}, // bg
		{{1, 1, 1, 1}, {0.29, 0.565, 0.851, 1}, {0.933, 0.922, 0.906, 1}, {0.29, 0.565, 0.851, 1}, {0.863, 0.855, 0.835, 1}},              // base
		{{0, 0, 0, 1}, {1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 1}, {0.459, 0.459, 0.459, 1}},                                                // text
	}
	for which, dst := range []*[numStates]RGBA{&s.Fg, &s.Bg, &s.Base, &s.Text} {
		for st := 0; st < numStates; st++ {
			if !s.set[which][st] {
				dst[st] = def[which][st]
			}
		}
	}
}

// Prop returns a style property's raw value.
func (s *Style) Prop(class, name string) (string, bool) {
	v, ok := s.Props[class+"::"+name]
	return v, ok
}

// PropInt returns an integer style property, or def.
func (s *Style) PropInt(class, name string, def int) int {
	if v, ok := s.Prop(class, name); ok {
		if n, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
			return n
		}
	}
	return def
}

// matchPath matches a widget or widget_class pattern against a widget path.
// In widget_class patterns "<Class>" matches a widget of that class or a subclass.
func matchPath(pattern string, full []Widget, classes bool) bool {
	type seg struct {
		start, end int
		w          Widget
	}
	var b strings.Builder
	var segs []seg
	for i, w := range full {
		if i > 0 {
			b.WriteByte('.')
		}
		start := b.Len()
		if classes {
			b.WriteString(w.Class[0])
		} else {
			b.WriteString(w.name())
		}
		segs = append(segs, seg{start, b.Len(), w})
	}
	str := b.String()
	segAt := func(pos int) (seg, bool) {
		for _, s := range segs {
			if s.start == pos {
				return s, true
			}
		}
		return seg{}, false
	}
	var match func(pi, si int) bool
	match = func(pi, si int) bool {
		for pi < len(pattern) {
			switch c := pattern[pi]; {
			case c == '*':
				for k := si; k <= len(str); k++ {
					if match(pi+1, k) {
						return true
					}
				}
				return false
			case c == '?':
				if si >= len(str) {
					return false
				}
				pi, si = pi+1, si+1
			case c == '<' && classes:
				end := strings.IndexByte(pattern[pi:], '>')
				if end < 0 {
					return false
				}
				class := pattern[pi+1 : pi+end]
				s, ok := segAt(si)
				if !ok {
					return false
				}
				found := false
				for _, c := range s.w.Class {
					if c == class {
						found = true
					}
				}
				if !found {
					return false
				}
				pi, si = pi+end+1, s.end
			default:
				if si >= len(str) || str[si] != c {
					return false
				}
				pi, si = pi+1, si+1
			}
		}
		return si == len(str)
	}
	return match(0, 0)
}
