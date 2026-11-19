// Package cosmic reads the COSMIC desktop's theme: the RON files under
// ~/.config/cosmic/com.system76.CosmicTheme.{Dark,Light}/v1/ that describe
// its colours (per container and per component state), corner radii and
// spacing.
package cosmic

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Value is a parsed RON value: map[string]Value for structs, []Value for
// tuples and lists, float64, string, bool or nil.
type Value = any

// Parse reads a RON document.
func Parse(src string) (Value, error) {
	p := &parser{s: src}
	v, err := p.value()
	if err != nil {
		return nil, err
	}
	return v, nil
}

type parser struct {
	s   string
	pos int
}

func (p *parser) skip() {
	for p.pos < len(p.s) {
		c := p.s[p.pos]
		switch {
		case c == ' ' || c == '\t' || c == '\n' || c == '\r':
			p.pos++
		case strings.HasPrefix(p.s[p.pos:], "//"):
			for p.pos < len(p.s) && p.s[p.pos] != '\n' {
				p.pos++
			}
		case strings.HasPrefix(p.s[p.pos:], "/*"):
			end := strings.Index(p.s[p.pos+2:], "*/")
			if end < 0 {
				p.pos = len(p.s)
				return
			}
			p.pos += end + 4
		default:
			return
		}
	}
}

func (p *parser) ident() string {
	start := p.pos
	for p.pos < len(p.s) {
		c := p.s[p.pos]
		if c == '_' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
			p.pos++
			continue
		}
		break
	}
	return p.s[start:p.pos]
}

func (p *parser) value() (Value, error) {
	p.skip()
	if p.pos >= len(p.s) {
		return nil, fmt.Errorf("cosmic: unexpected end")
	}
	c := p.s[p.pos]
	switch {
	case c == '(' || c == '[' || c == '{':
		return p.group()
	case c == '"':
		p.pos++
		var b strings.Builder
		for p.pos < len(p.s) && p.s[p.pos] != '"' {
			if p.s[p.pos] == '\\' && p.pos+1 < len(p.s) {
				p.pos++
			}
			b.WriteByte(p.s[p.pos])
			p.pos++
		}
		p.pos++
		return b.String(), nil
	case c == '-' || c == '+' || c == '.' || c >= '0' && c <= '9':
		start := p.pos
		p.pos++
		for p.pos < len(p.s) && strings.IndexByte("0123456789.eE+-_", p.s[p.pos]) >= 0 {
			p.pos++
		}
		f, err := strconv.ParseFloat(strings.ReplaceAll(p.s[start:p.pos], "_", ""), 64)
		return f, err
	}
	id := p.ident()
	if id == "" {
		return nil, fmt.Errorf("cosmic: unexpected %q at %d", c, p.pos)
	}
	switch id {
	case "true":
		return true, nil
	case "false":
		return false, nil
	case "None":
		return nil, nil
	}
	p.skip()
	if p.pos < len(p.s) && p.s[p.pos] == '(' { // Some(x), or a named struct/enum variant
		v, err := p.group()
		if list, ok := v.([]Value); ok && len(list) == 1 {
			return list[0], err
		}
		return v, err
	}
	return id, nil // a unit enum variant
}

// group parses (...), [...] or {...}: a struct if it has "name:" fields,
// otherwise a list.
func (p *parser) group() (Value, error) {
	open := p.s[p.pos]
	close := map[byte]byte{'(': ')', '[': ']', '{': '}'}[open]
	p.pos++
	fields := map[string]Value{}
	var list []Value
	for {
		p.skip()
		if p.pos >= len(p.s) {
			return nil, fmt.Errorf("cosmic: unterminated %c", open)
		}
		if p.s[p.pos] == close {
			p.pos++
			break
		}
		// "name: value" ?
		save := p.pos
		if id := p.ident(); id != "" {
			p.skip()
			if p.pos < len(p.s) && p.s[p.pos] == ':' {
				p.pos++
				v, err := p.value()
				if err != nil {
					return nil, err
				}
				fields[id] = v
				p.skip()
				if p.pos < len(p.s) && p.s[p.pos] == ',' {
					p.pos++
				}
				continue
			}
		}
		p.pos = save
		v, err := p.value()
		if err != nil {
			return nil, err
		}
		list = append(list, v)
		p.skip()
		if p.pos < len(p.s) && p.s[p.pos] == ',' {
			p.pos++
		}
	}
	if len(fields) > 0 {
		return fields, nil
	}
	return list, nil
}

// RGBA is a colour, 0-1.
type RGBA struct{ R, G, B, A float64 }

// Get walks a path of struct fields.
func Get(v Value, path ...string) Value {
	for _, k := range path {
		m, ok := v.(map[string]Value)
		if !ok {
			return nil
		}
		v = m[k]
	}
	return v
}

// Color reads a {red, green, blue, alpha} struct at a path.
func Color(v Value, path ...string) (RGBA, bool) {
	m, ok := Get(v, path...).(map[string]Value)
	if !ok {
		return RGBA{}, false
	}
	f := func(k string, def float64) float64 {
		if x, ok := m[k].(float64); ok {
			return x
		}
		return def
	}
	if _, ok := m["red"]; !ok {
		return RGBA{}, false
	}
	return RGBA{f("red", 0), f("green", 0), f("blue", 0), f("alpha", 1)}, true
}

// Theme is a loaded COSMIC theme: each file of the theme directory, parsed.
type Theme struct {
	Dark  bool
	Files map[string]Value // "background", "primary", "accent", "button", "corner_radii"...
}

// Load reads the current theme from a config directory (~/.config).
func Load(configHome string) (*Theme, error) {
	root := filepath.Join(configHome, "cosmic")
	t := &Theme{Files: map[string]Value{}}
	if b, err := os.ReadFile(filepath.Join(root, "com.system76.CosmicTheme.Mode", "v1", "is_dark")); err == nil {
		t.Dark = strings.TrimSpace(string(b)) == "true"
	}
	name := "com.system76.CosmicTheme.Light"
	if t.Dark {
		name = "com.system76.CosmicTheme.Dark"
	}
	dir := filepath.Join(root, name, "v1")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		if v, err := Parse(string(b)); err == nil {
			t.Files[e.Name()] = v
		}
	}
	return t, nil
}

// Color reads a colour from one of the theme's files.
func (t *Theme) Color(file string, path ...string) (RGBA, bool) {
	return Color(t.Files[file], path...)
}

// Radius reads a corner radius (the first of its four values).
func (t *Theme) Radius(name string, def float64) float64 {
	if list, ok := Get(t.Files["corner_radii"], name).([]Value); ok && len(list) > 0 {
		if f, ok := list[0].(float64); ok {
			return f
		}
	}
	return def
}
