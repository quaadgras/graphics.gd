package bbparser

import (
	"strings"
)

const (
	startingAttr = "starting"
)

// Parser - parse bbcode to html
type Parser struct {
	// Handlers - pairs with tagname - handler function
	Handlers map[string]Handler
	// SpecialChars - contains chars/strings that should be replaced
	SpecialChars map[string]string
}

// New - return pointer on new Parser instance
func New() *Parser {
	parser := &Parser{}
	parser.Handlers = make(map[string]Handler)
	parser.SpecialChars = make(map[string]string)
	return parser
}

// NewDefault - return pointer on parser with all defaults tags and special chars
func NewDefault() *Parser {
	parser := &Parser{}
	parser.Handlers = make(map[string]Handler)
	parser.SpecialChars = make(map[string]string)

	//simple codes
	parser.Handlers["b"] = SimpleHandler
	parser.Handlers["i"] = SimpleHandler
	parser.Handlers["u"] = SimpleHandler
	parser.Handlers["s"] = SimpleHandler
	parser.Handlers["ul"] = SimpleHandler
	parser.Handlers["ol"] = SimpleHandler
	parser.Handlers["li"] = SimpleHandler

	//complex handlers
	parser.Handlers["left"] = TextAlignHandler
	parser.Handlers["center"] = TextAlignHandler
	parser.Handlers["right"] = TextAlignHandler
	parser.Handlers["color"] = FontColorHandler
	parser.Handlers["size"] = FontSizeHandler
	parser.Handlers["url"] = URLHandler
	parser.Handlers["ol"] = ListHandler
	parser.Handlers["ul"] = ListHandler
	parser.Handlers["list"] = ListHandler

	//add special characters or strings that should be replaced
	parser.SpecialChars["\n"] = "<br>"

	return parser
}

// AddTag - add tag to parser (if tag already exists - it will be override with new one)
func (p *Parser) AddTag(tagName string, handler Handler) {
	if p.Handlers == nil {
		p.Handlers = make(map[string]Handler)
	}
	p.Handlers[tagName] = handler
}

// AddSpecialString - add string to be replaced with (if string already in parser - it will be override with new one)
func (p *Parser) AddSpecialString(old, new string) {
	if p.SpecialChars == nil {
		p.SpecialChars[old] = new
	}
}

// Parse - parse given string with parser settings
func (p *Parser) Parse(str string) string {
	for i := 0; i < len(str)-1; i++ {
		if str[i] == '[' {
			for j := i + 1; j < len(str); j++ {
				var tag Tag
				if str[j] == ']' {
					tag = p.parseTag(str[i+1 : j])
					if tag.Name == "" {
						i = j
						break
					}
					if !tag.Closing {
						contentEnd, tagEnd, found := p.findEnd(str[j+1:], tag.Name)
						if !found {
							str = str[:i] + p.Handlers[tag.Name](tag, "") + str[j+1:]
							break
						}
						str = str[:i] + p.Handlers[tag.Name](tag, str[j+1:j+1+contentEnd]) + str[j+1+tagEnd+1:]
						break
					}
				}

			}
		}
	}
	str = p.processSpecialChars(str)
	return str
}

func (p *Parser) parseTag(raw string) Tag {
	var tag Tag
	if len(raw) < 1 {
		return tag
	}
	if raw[0] == '/' {
		tag.Closing = true
		tag.Name = raw[1:]
		return tag
	}

	//Split always return slice with length > 0 when separator != ""
	attributes := strings.Split(raw, " ")
	starting := strings.Split(attributes[0], "=")
	tag.Name = starting[0]
	if _, ok := p.Handlers[tag.Name]; !ok {
		tag.Name = ""
		return tag
	}
	tag.Attributes = make(map[string]string)
	if len(starting) > 1 {
		tag.Attributes[startingAttr] = starting[1]
	}
	for i := 1; i < len(attributes); i++ {
		attrStr := strings.Split(attributes[i], "=")
		if len(attrStr) > 1 {
			tag.Attributes[attrStr[0]] = attrStr[1]
		} else {
			tag.Attributes[attrStr[0]] = ""
		}
	}
	return tag
}

func (p *Parser) findEnd(str, tagName string) (int, int, bool) {
	var count int
	for i := 0; i < len(str)-1; i++ {
		if str[i] == '[' {
			for j := i + 1; j < len(str); j++ {
				var tag Tag
				if str[j] == ']' {
					tag = p.parseTag(str[i+1 : j])
					if tag.Name == tagName {
						if tag.Closing {
							if count == 0 {
								return i, j, true
							}
							count--
							i = j
							break
						}
						count++
					}
					i = j
					break
				}
			}
		}
	}
	return 0, 0, false
}

func (p *Parser) processSpecialChars(str string) string {
	for k, v := range p.SpecialChars {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}
