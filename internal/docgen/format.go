package docgen

import "encoding/xml"

type XML struct {
	XMLName xml.Name `xml:"class"`

	Name     string `xml:"name,attr"`
	Inherits string `xml:"inherits,attr"`
	Version  string `xml:"version,attr"`

	BriefDescription string     `xml:"brief_description"`
	Description      string     `xml:"description"`
	Tutorials        []Link     `xml:"tutorials>link"`
	Methods          []Method   `xml:"methods>method"`
	Members          []Member   `xml:"members>member"`
	Signals          []Signal   `xml:"signals>signal"`
	Constants        []Constant `xml:"constants>constant"`
}

type Link struct {
	Title string `xml:"title,attr"`
	Value string `xml:",chardata"`
}

type Method struct {
	Name        string     `xml:"name,attr"`
	Return      Return     `xml:"return"`
	Arguments   []Argument `xml:"param"`
	Description string     `xml:"description"`
}

type Return struct {
	Type string `xml:"type,attr"`
}

type Argument struct {
	Index   int    `xml:"index,attr"`
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Default string `xml:"default,attr"`
}

type Signal struct {
	Name        string     `xml:"name,attr"`
	Param       []Argument `xml:"param"`
	Description string     `xml:"description"`
}

type Member struct {
	Name        string `xml:"name,attr"`
	Type        string `xml:"type,attr"`
	Setter      string `xml:"setter,attr"`
	Getter      string `xml:"getter,attr"`
	Description string `xml:",chardata"`
}

type Constant struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	Enum        string `xml:"enum,attr"`
	Description string `xml:",chardata"`
}
