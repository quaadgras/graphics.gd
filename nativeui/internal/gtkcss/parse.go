// Package gtkcss reads GTK 3 themes (gtk-3.0/gtk.css and the files it
// imports) and computes the style GTK would give a widget node, without
// GTK. It is pure Go so that it can be tested without the engine.
package gtkcss

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Sheet is a parsed theme.
type Sheet struct {
	Dir    string            // directory of the root stylesheet, url()s resolve against the file that uses them
	Colors map[string]string // @define-color name -> unevaluated value
	Rules  []Rule
}

// Rule is one selector list with its declarations.
type Rule struct {
	Selectors []Selector
	Decls     []Decl
	Order     int // source order, for the cascade
}

// Decl is one property: value declaration. Base is the directory of the file
// it came from, for resolving url()s.
type Decl struct {
	Name, Value string
	Base        string
}

// Load parses the stylesheet at path, following @import.
func Load(path string) (*Sheet, error) {
	sheet := &Sheet{Dir: filepath.Dir(path), Colors: map[string]string{}}
	if err := sheet.load(path, map[string]bool{}); err != nil {
		return nil, err
	}
	return sheet, nil
}

// Parse parses a stylesheet from a string; base resolves url()s and imports.
func Parse(src, base string) (*Sheet, error) {
	sheet := &Sheet{Dir: base, Colors: map[string]string{}}
	return sheet, sheet.parse(src, base, map[string]bool{})
}

func (s *Sheet) load(path string, seen map[string]bool) error {
	if seen[path] {
		return nil
	}
	seen[path] = true
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return s.parse(string(data), filepath.Dir(path), seen)
}

func stripComments(src string) string {
	var b strings.Builder
	for {
		i := strings.Index(src, "/*")
		if i < 0 {
			b.WriteString(src)
			return b.String()
		}
		b.WriteString(src[:i])
		j := strings.Index(src[i+2:], "*/")
		if j < 0 {
			return b.String()
		}
		src = src[i+2+j+2:]
	}
}

// scanTo returns the index of the first stop byte at parenthesis depth 0
// outside of quotes, or -1.
func scanTo(s string, stops string) int {
	depth := 0
	var quote byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case quote != 0:
			if c == '\\' {
				i++
			} else if c == quote {
				quote = 0
			}
		case c == '"' || c == '\'':
			quote = c
		case c == '(':
			depth++
		case c == ')':
			depth--
		case depth == 0 && strings.IndexByte(stops, c) >= 0:
			return i
		}
	}
	return -1
}

// matchingBrace returns the index of the brace closing the one at s[0].
func matchingBrace(s string) int {
	depth := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

func (s *Sheet) parse(src, base string, seen map[string]bool) error {
	src = stripComments(src)
	for {
		src = strings.TrimSpace(src)
		if src == "" {
			return nil
		}
		if src[0] == '@' {
			end := scanTo(src, ";{")
			if end < 0 {
				return nil
			}
			if src[end] == '{' { // @keyframes, @binding-set...
				close := matchingBrace(src[end:])
				if close < 0 {
					return nil
				}
				src = src[end+close+1:]
				continue
			}
			stmt := strings.TrimSpace(src[:end])
			src = src[end+1:]
			switch {
			case strings.HasPrefix(stmt, "@import"):
				if file := urlArg(strings.TrimSpace(strings.TrimPrefix(stmt, "@import"))); file != "" {
					path := file
					if !filepath.IsAbs(path) {
						path = filepath.Join(base, file)
					}
					if err := s.load(path, seen); err != nil && !os.IsNotExist(err) {
						return fmt.Errorf("gtkcss: import %s: %w", file, err)
					}
				}
			case strings.HasPrefix(stmt, "@define-color"):
				fields := strings.Fields(strings.TrimPrefix(stmt, "@define-color"))
				if len(fields) >= 2 {
					s.Colors[fields[0]] = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(strings.TrimPrefix(stmt, "@define-color")), fields[0]))
				}
			}
			continue
		}
		open := scanTo(src, "{")
		if open < 0 {
			return nil
		}
		close := matchingBrace(src[open:])
		if close < 0 {
			return nil
		}
		selectors := parseSelectorList(src[:open])
		body := src[open+1 : open+close]
		src = src[open+close+1:]
		var decls []Decl
		for len(body) > 0 {
			end := scanTo(body, ";")
			part := body
			if end >= 0 {
				part, body = body[:end], body[end+1:]
			} else {
				body = ""
			}
			name, value, ok := strings.Cut(part, ":")
			if !ok {
				continue
			}
			value = strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(value), "!important"))
			decls = append(decls, Decl{Name: strings.ToLower(strings.TrimSpace(name)), Value: value, Base: base})
		}
		if len(selectors) > 0 && len(decls) > 0 {
			s.Rules = append(s.Rules, Rule{Selectors: selectors, Decls: decls, Order: len(s.Rules)})
		}
	}
}

// urlArg extracts the path from url("x"), url(x) or "x".
func urlArg(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "url(") && strings.HasSuffix(s, ")") {
		s = strings.TrimSpace(s[4 : len(s)-1])
	}
	return strings.Trim(s, `"'`)
}
