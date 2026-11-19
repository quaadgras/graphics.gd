package gtkcss

import (
	"sort"
	"strings"
)

// Node is one element of a GTK CSS node tree, e.g. button.suggested-action:hover.
type Node struct {
	Name    string
	Classes []string
	States  []string // hover, active, checked, disabled, focus, selected, backdrop, indeterminate...

	// Position among siblings, for :first-child and friends.
	First, Last bool
}

// N builds a node from GTK's selector-like notation: "button.text-button:hover".
func N(spec string) Node {
	var n Node
	n.First, n.Last = true, true
	cur := &n.Name
	var buf strings.Builder
	flush := func() {
		if cur == &n.Name {
			n.Name = buf.String()
		}
		buf.Reset()
	}
	kind := byte(0)
	for i := 0; i <= len(spec); i++ {
		if i == len(spec) || spec[i] == '.' || spec[i] == ':' {
			switch kind {
			case 0:
				flush()
			case '.':
				n.Classes = append(n.Classes, buf.String())
				buf.Reset()
			case ':':
				n.States = append(n.States, buf.String())
				buf.Reset()
			}
			if i < len(spec) {
				kind = spec[i]
				cur = nil
			}
			continue
		}
		buf.WriteByte(spec[i])
	}
	return n
}

// Path builds an ancestor chain from space separated nodes, root first. A
// node written with a leading "+" has an earlier sibling (so it is not
// :first-child or :only-child), a trailing "+" a later one.
func Path(spec string) []Node {
	var path []Node
	for _, f := range strings.Fields(spec) {
		before, after := strings.HasPrefix(f, "+"), strings.HasSuffix(f, "+")
		n := N(strings.Trim(f, "+"))
		n.First, n.Last = !before, !after
		path = append(path, n)
	}
	return path
}

func has(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

// Selector is a parsed complex selector, subject last.
type Selector struct {
	Parts       []compound
	Combinators []byte // between Parts[i] and Parts[i+1]: ' ' descendant, '>' child, '+'/'~' siblings
	Spec        [3]int // ids, classes/pseudo-classes, types
}

type compound struct {
	name    string // "" or "*" match any
	classes []string
	ids     []string
	pseudo  []pseudo
	never   bool // uses something we can't match (attributes...)
}

type pseudo struct {
	name string
	not  []compound // for :not(), each alternative
	arg  string
}

func parseSelectorList(s string) []Selector {
	var list []Selector
	for len(s) > 0 {
		end := scanTo(s, ",")
		part := s
		if end >= 0 {
			part, s = s[:end], s[end+1:]
		} else {
			s = ""
		}
		if sel, ok := parseSelector(strings.TrimSpace(part)); ok {
			list = append(list, sel)
		}
	}
	return list
}

func parseSelector(s string) (Selector, bool) {
	var sel Selector
	if s == "" {
		return sel, false
	}
	pendingComb := byte(0)
	for i := 0; i < len(s); {
		c := s[i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			if len(sel.Parts) > 0 && pendingComb == 0 {
				pendingComb = ' '
			}
			i++
			continue
		}
		if c == '>' || c == '+' || c == '~' {
			pendingComb = c
			i++
			continue
		}
		// compound runs until whitespace or a combinator at depth 0
		j := i
		depth := 0
		for j < len(s) {
			if s[j] == '(' {
				depth++
			} else if s[j] == ')' {
				depth--
			} else if depth == 0 && strings.IndexByte(" \t\n\r>+~", s[j]) >= 0 {
				break
			}
			j++
		}
		comp, spec := parseCompound(s[i:j])
		if len(sel.Parts) > 0 {
			if pendingComb == 0 {
				pendingComb = ' '
			}
			sel.Combinators = append(sel.Combinators, pendingComb)
		}
		pendingComb = 0
		sel.Parts = append(sel.Parts, comp)
		for k := range spec {
			sel.Spec[k] += spec[k]
		}
		i = j
	}
	return sel, len(sel.Parts) > 0
}

func parseCompound(s string) (compound, [3]int) {
	var c compound
	var spec [3]int
	i := 0
	readIdent := func() string {
		j := i
		for j < len(s) && (s[j] == '-' || s[j] == '_' || s[j] >= '0' && s[j] <= '9' || s[j] >= 'a' && s[j] <= 'z' || s[j] >= 'A' && s[j] <= 'Z' || s[j] >= 0x80) {
			j++
		}
		id := s[i:j]
		i = j
		return id
	}
	if i < len(s) && s[i] == '*' {
		c.name = "*"
		i++
	} else if i < len(s) && s[i] != '.' && s[i] != ':' && s[i] != '#' && s[i] != '[' {
		c.name = readIdent()
		spec[2]++
	}
	for i < len(s) {
		switch s[i] {
		case '.':
			i++
			c.classes = append(c.classes, readIdent())
			spec[1]++
		case '#':
			i++
			c.ids = append(c.ids, readIdent())
			spec[0]++
		case ':':
			i++
			if i < len(s) && s[i] == ':' { // pseudo-element: unsupported
				c.never = true
				i++
			}
			p := pseudo{name: readIdent()}
			if i < len(s) && s[i] == '(' {
				end := strings.IndexByte(s[i:], ')')
				// find the matching paren
				depth := 0
				for k := i; k < len(s); k++ {
					if s[k] == '(' {
						depth++
					} else if s[k] == ')' {
						depth--
						if depth == 0 {
							end = k - i
							break
						}
					}
				}
				p.arg = s[i+1 : i+end]
				i += end + 1
			}
			if p.name == "not" {
				for _, alt := range parseSelectorList(p.arg) {
					if len(alt.Parts) == 1 {
						p.not = append(p.not, alt.Parts[0])
						for k := range alt.Spec {
							spec[k] += alt.Spec[k]
						}
					}
				}
			} else {
				spec[1]++
			}
			c.pseudo = append(c.pseudo, p)
		case '[':
			end := strings.IndexByte(s[i:], ']')
			if end < 0 {
				end = len(s) - i - 1
			}
			i += end + 1
			c.never = true
			spec[1]++
		default:
			i++ // unexpected character, skip
		}
	}
	return c, spec
}

func (c compound) matches(n Node) bool {
	if c.never {
		return false
	}
	if c.name != "" && c.name != "*" && c.name != n.Name {
		return false
	}
	for _, class := range c.classes {
		if !has(n.Classes, class) {
			return false
		}
	}
	if len(c.ids) > 0 {
		return false // we never give nodes ids
	}
	for _, p := range c.pseudo {
		switch p.name {
		case "not":
			for _, alt := range p.not {
				if alt.matches(n) {
					return false
				}
			}
		case "first-child":
			if !n.First {
				return false
			}
		case "last-child":
			if !n.Last {
				return false
			}
		case "only-child":
			if !n.First || !n.Last {
				return false
			}
		case "dir":
			if strings.TrimSpace(p.arg) != "ltr" {
				return false
			}
		case "nth-child", "nth-last-child":
			if a := strings.TrimSpace(p.arg); a != "1" && a != "odd" && a != "first" {
				return false
			}
		case "prelight":
			if !has(n.States, "hover") {
				return false
			}
		case "insensitive":
			if !has(n.States, "disabled") {
				return false
			}
		default:
			if !has(n.States, p.name) {
				return false
			}
		}
	}
	return true
}

// Matches reports whether the selector matches the last node of path.
func (s Selector) Matches(path []Node) bool {
	return s.match(len(s.Parts)-1, path, len(path)-1)
}

func (s Selector) match(part int, path []Node, at int) bool {
	if at < 0 || !s.Parts[part].matches(path[at]) {
		return false
	}
	if part == 0 {
		return true
	}
	switch s.Combinators[part-1] {
	case '>':
		return s.match(part-1, path, at-1)
	case ' ':
		for k := at - 1; k >= 0; k-- {
			if s.match(part-1, path, k) {
				return true
			}
		}
		return false
	default: // sibling combinators: our paths have no siblings
		return false
	}
}

type match struct {
	decl  Decl
	spec  [3]int
	order int
}

// Declared returns the cascaded declarations for the last node of path
// (no inheritance), property name -> declaration.
func (s *Sheet) Declared(path []Node) map[string]Decl {
	var matches []match
	for _, r := range s.Rules {
		best := [3]int{-1}
		for _, sel := range r.Selectors {
			if sel.Matches(path) && less(best, sel.Spec) {
				best = sel.Spec
			}
		}
		if best[0] < 0 {
			continue
		}
		for _, d := range r.Decls {
			matches = append(matches, match{decl: d, spec: best, order: r.Order})
		}
	}
	sort.SliceStable(matches, func(i, j int) bool {
		if matches[i].spec != matches[j].spec {
			return less(matches[i].spec, matches[j].spec)
		}
		return matches[i].order < matches[j].order
	})
	out := map[string]Decl{}
	for _, m := range matches {
		for _, d := range expand(m.decl) {
			out[d.Name] = d
		}
	}
	return out
}

func less(a, b [3]int) bool {
	for i := range a {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return false
}
