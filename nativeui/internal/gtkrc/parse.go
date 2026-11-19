// Package gtkrc reads GTK 2 themes (gtk-2.0/gtkrc), resolves the style GTK 2
// would give a widget and draws its parts the way the theme's engine would:
// the pixmap engine's images, or GTK 2's default bevels for everything else.
// It is pure Go so that it can be tested without GTK or the engine.
package gtkrc

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// States, in GTK 2's order.
const (
	Normal = iota
	Active
	Prelight
	Selected
	Insensitive
	numStates
)

var stateNames = map[string]int{"NORMAL": Normal, "ACTIVE": Active, "PRELIGHT": Prelight, "SELECTED": Selected, "INSENSITIVE": Insensitive}

// RGBA is a color, components 0-1.
type RGBA struct{ R, G, B, A float64 }

// Style is a gtkrc style, possibly merged from several.
type Style struct {
	Name               string
	Fg, Bg, Base, Text [numStates]RGBA
	set                [4][numStates]bool // which of Fg, Bg, Base, Text were given
	XThickness         int
	YThickness         int
	thicknessSet       [2]bool
	Font               string
	Props              map[string]string // "GtkButton::default-border" -> raw value
	Engine             string
	Images             []Image           // pixmap engine
	EngineParams       map[string]string // other engines' parameters
	Colors             map[string]string // color["name"] = ...
	scheme             map[string]string // the theme\'s gtk_color_scheme, for engine options naming its colours
}

// Image is one pixmap engine image rule.
type Image struct {
	Function, Detail, ArrowDirection, GapSide, Orientation string
	State                                                  int // -1: any
	Shadow                                                 string
	File, OverlayFile, GapFile, GapStartFile, GapEndFile   string
	Border, OverlayBorder, GapBorder                       [4]int // left, right, top, bottom
	Stretch, OverlayStretch                                bool
	stretchSet                                             bool
}

type binding struct {
	kind     int // 0 class, 1 widget_class, 2 widget
	pattern  string
	style    string
	priority int
	order    int
}

// Theme is a parsed gtkrc.
type Theme struct {
	Dir        string
	Scheme     map[string]string // gtk_color_scheme name -> color
	Styles     map[string]*Style
	bindings   []binding
	pixmapPath []string
}

// Load parses a gtkrc file and its includes.
func Load(path string) (*Theme, error) {
	t := &Theme{Dir: filepath.Dir(path), Scheme: map[string]string{}, Styles: map[string]*Style{}}
	if err := t.load(path, 0); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Theme) load(path string, depth int) error {
	if depth > 16 {
		return nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	p := &parser{toks: tokenize(string(data)), dir: filepath.Dir(path), theme: t, depth: depth}
	return p.top()
}

// tokens

type token struct {
	kind byte // 's' string, 'i' identifier/number, 'p' punctuation
	text string
}

func tokenize(src string) []token {
	var toks []token
	for i := 0; i < len(src); {
		c := src[i]
		switch {
		case c == '#':
			for i < len(src) && src[i] != '\n' {
				i++
			}
		case c == '/' && i+1 < len(src) && src[i+1] == '*':
			end := strings.Index(src[i+2:], "*/")
			if end < 0 {
				return toks
			}
			i += end + 4
		case c == ' ' || c == '\t' || c == '\n' || c == '\r':
			i++
		case c == '"':
			var b strings.Builder
			i++
			for i < len(src) && src[i] != '"' {
				if src[i] == '\\' && i+1 < len(src) {
					i++
					switch src[i] {
					case 'n':
						b.WriteByte('\n')
					case 't':
						b.WriteByte('\t')
					default:
						b.WriteByte(src[i])
					}
				} else {
					b.WriteByte(src[i])
				}
				i++
			}
			i++
			toks = append(toks, token{'s', b.String()})
		case strings.IndexByte("{}[]=,()@:;<>", c) >= 0:
			if c == ':' && i+1 < len(src) && src[i+1] == ':' { // Class::property
				toks = append(toks, token{'p', "::"})
				i += 2
				continue
			}
			toks = append(toks, token{'p', string(c)})
			i++
		default:
			j := i
			for j < len(src) && strings.IndexByte(" \t\n\r{}[]=,()@:;\"#<>", src[j]) < 0 {
				j++
			}
			if j == i {
				j++
			}
			toks = append(toks, token{'i', src[i:j]})
			i = j
		}
	}
	return toks
}

type parser struct {
	toks  []token
	pos   int
	dir   string
	theme *Theme
	depth int
}

func (p *parser) peek() token {
	if p.pos < len(p.toks) {
		return p.toks[p.pos]
	}
	return token{}
}

func (p *parser) next() token {
	t := p.peek()
	p.pos++
	return t
}

func (p *parser) accept(text string) bool {
	if p.peek().text == text && p.peek().kind != 's' {
		p.pos++
		return true
	}
	return false
}

// skipBlock skips a balanced {...} starting at the current "{".
func (p *parser) skipBlock() {
	depth := 0
	for p.pos < len(p.toks) {
		t := p.next()
		if t.kind == 'p' && t.text == "{" {
			depth++
		} else if t.kind == 'p' && t.text == "}" {
			depth--
			if depth == 0 {
				return
			}
		}
	}
}

// value reads a value up to the end of a statement: a string, a number, an
// identifier, a {..} list or a color expression. It returns the raw text.
func (p *parser) value() string {
	t := p.peek()
	if t.kind == 'p' && t.text == "{" {
		var parts []string
		p.next()
		for p.pos < len(p.toks) && !(p.peek().kind == 'p' && p.peek().text == "}") {
			parts = append(parts, p.value())
			p.accept(",")
		}
		p.accept("}")
		return "{" + strings.Join(parts, ",") + "}"
	}
	if t.kind == 'p' && t.text == "@" {
		p.next()
		return "@" + p.next().text
	}
	p.next()
	if t.kind == 's' {
		return strconv.Quote(t.text)
	}
	if t.kind == 'i' && p.peek().kind == 'p' && p.peek().text == "(" { // shade (0.5, @bg)
		var parts []string
		p.next()
		for p.pos < len(p.toks) && !(p.peek().kind == 'p' && p.peek().text == ")") {
			parts = append(parts, p.value())
			p.accept(",")
		}
		p.accept(")")
		return t.text + "(" + strings.Join(parts, ",") + ")"
	}
	if t.text == "-" && p.peek().kind == 'i' { // negative number split by the tokenizer? (not split, but be safe)
		return "-" + p.next().text
	}
	return t.text
}

func unquote(s string) string {
	if u, err := strconv.Unquote(s); err == nil {
		return u
	}
	return s
}

func (p *parser) top() error {
	order := len(p.theme.bindings)
	for p.pos < len(p.toks) {
		t := p.next()
		switch {
		case t.kind == 'i' && t.text == "include":
			file := p.next().text
			path := file
			if !filepath.IsAbs(path) {
				path = filepath.Join(p.dir, file)
			}
			if err := p.theme.load(path, p.depth+1); err != nil && !os.IsNotExist(err) {
				return fmt.Errorf("gtkrc: include %s: %w", file, err)
			}
		case t.kind == 'i' && (t.text == "gtk_color_scheme" || t.text == "gtk-color-scheme"):
			p.accept("=")
			// entries are separated by newlines or semicolons
			for _, line := range strings.FieldsFunc(unquote(p.value()), func(r rune) bool { return r == '\n' || r == ';' }) {
				if k, v, ok := strings.Cut(line, ":"); ok {
					p.theme.Scheme[strings.TrimSpace(k)] = strings.TrimSpace(v)
				}
			}
		case t.kind == 'i' && t.text == "pixmap_path":
			p.theme.pixmapPath = append(p.theme.pixmapPath, filepath.SplitList(p.next().text)...)
		case t.kind == 'i' && t.text == "style":
			p.style()
		case t.kind == 'i' && (t.text == "class" || t.text == "widget_class" || t.text == "widget"):
			kind := map[string]int{"class": 0, "widget_class": 1, "widget": 2}[t.text]
			pattern := p.next().text
			p.accept("style")
			priority := 0
			if p.accept(":") {
				priority = map[string]int{"lowest": 0, "gtk": 4, "application": 10, "theme": 6, "rc": 8, "highest": 15}[p.next().text]
			}
			name := p.next().text
			p.theme.bindings = append(p.theme.bindings, binding{kind, pattern, name, priority, order})
			order++
		case t.kind == 'i' && t.text == "binding":
			p.next()
			p.skipBlock()
		case t.kind == 'p' && t.text == "{":
			p.pos--
			p.skipBlock()
		case t.kind == 'i' && p.peek().text == "=": // gtk-setting = value
			p.next()
			p.value()
		}
	}
	return nil
}

func (p *parser) style() {
	name := p.next().text
	st := &Style{Name: name, Props: map[string]string{}, EngineParams: map[string]string{}, Colors: map[string]string{}}
	if p.accept("=") {
		if parent, ok := p.theme.Styles[p.next().text]; ok {
			cp := *parent
			cp.Name = name
			cp.Props = copyMap(parent.Props)
			cp.EngineParams = copyMap(parent.EngineParams)
			cp.Colors = copyMap(parent.Colors)
			cp.Images = append([]Image(nil), parent.Images...)
			st = &cp
		}
	}
	if !p.accept("{") {
		return
	}
	for p.pos < len(p.toks) && !p.accept("}") {
		t := p.next()
		if t.kind != 'i' {
			continue
		}
		switch t.text {
		case "fg", "bg", "base", "text", "bg_pixmap":
			state := Normal
			if p.accept("[") {
				state = stateNames[p.next().text]
				p.accept("]")
			}
			p.accept("=")
			v := p.value()
			which := map[string]int{"fg": 0, "bg": 1, "base": 2, "text": 3}[t.text]
			if t.text == "bg_pixmap" {
				continue
			}
			if c, ok := p.theme.color(v, st); ok {
				[]*[numStates]RGBA{&st.Fg, &st.Bg, &st.Base, &st.Text}[which][state] = c
				st.set[which][state] = true
			}
		case "xthickness", "ythickness":
			p.accept("=")
			n, _ := strconv.Atoi(p.value())
			if t.text == "xthickness" {
				st.XThickness, st.thicknessSet[0] = n, true
			} else {
				st.YThickness, st.thicknessSet[1] = n, true
			}
		case "font_name", "font":
			p.accept("=")
			st.Font = unquote(p.value())
		case "color":
			p.accept("[")
			key := p.next().text
			p.accept("]")
			p.accept("=")
			st.Colors[key] = p.value()
		case "engine":
			st.Engine = p.next().text
			st.Images = nil
			if p.accept("{") {
				p.engine(st)
			}
		case "stock", "icon_name":
			p.accept("[")
			p.next()
			p.accept("]")
			p.accept("=")
			p.value()
		default:
			if p.accept("::") { // Class::property = value
				prop := p.next().text
				p.accept("=")
				st.Props[t.text+"::"+strings.ReplaceAll(prop, "_", "-")] = p.value()
			} else if p.accept("=") {
				p.value()
			}
		}
	}
	p.theme.Styles[name] = st
}

func copyMap(m map[string]string) map[string]string {
	out := make(map[string]string, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}

func (p *parser) engine(st *Style) {
	for p.pos < len(p.toks) && !p.accept("}") {
		t := p.next()
		if t.kind == 'i' && t.text == "image" && p.accept("{") {
			img := Image{State: -1, Stretch: true, OverlayStretch: false}
			for p.pos < len(p.toks) && !p.accept("}") {
				key := p.next().text
				p.accept("=")
				v := p.value()
				u := unquote(v)
				switch key {
				case "function":
					img.Function = u
				case "detail":
					img.Detail = u
				case "state":
					if s, ok := stateNames[u]; ok {
						img.State = s
					}
				case "shadow":
					img.Shadow = u
				case "arrow_direction":
					img.ArrowDirection = u
				case "gap_side":
					img.GapSide = u
				case "orientation":
					img.Orientation = u
				case "file":
					img.File = p.theme.resolve(p.dir, u)
				case "overlay_file":
					img.OverlayFile = p.theme.resolve(p.dir, u)
				case "gap_file":
					img.GapFile = p.theme.resolve(p.dir, u)
				case "gap_start_file":
					img.GapStartFile = p.theme.resolve(p.dir, u)
				case "gap_end_file":
					img.GapEndFile = p.theme.resolve(p.dir, u)
				case "border":
					img.Border = ints4(v)
				case "overlay_border":
					img.OverlayBorder = ints4(v)
				case "gap_border":
					img.GapBorder = ints4(v)
				case "stretch":
					img.Stretch, img.stretchSet = u == "TRUE", true
				case "overlay_stretch":
					img.OverlayStretch = u == "TRUE"
				}
			}
			st.Images = append(st.Images, img)
			continue
		}
		if t.kind == 'i' && p.accept("=") {
			st.EngineParams[t.text] = p.value()
		} else if t.kind == 'p' && t.text == "{" {
			p.pos--
			p.skipBlock()
		}
	}
}

func ints4(v string) [4]int {
	var out [4]int
	parts := strings.Split(strings.Trim(v, "{}"), ",")
	for i := 0; i < 4 && i < len(parts); i++ {
		out[i], _ = strconv.Atoi(strings.TrimSpace(parts[i]))
	}
	return out
}

func (t *Theme) resolve(dir, file string) string {
	if filepath.IsAbs(file) {
		return file
	}
	for _, d := range append([]string{dir}, t.pixmapPath...) {
		if !filepath.IsAbs(d) {
			d = filepath.Join(t.Dir, d)
		}
		path := filepath.Join(d, file)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return filepath.Join(dir, file)
}
