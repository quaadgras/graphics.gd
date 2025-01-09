package classdb

import (
	"encoding/xml"
)

type xmlDocumentation struct {
	XMLName xml.Name `xml:"class"`

	Name     string `xml:"name,attr"`
	Inherits string `xml:"inherits,attr"`
	Version  string `xml:"version,attr"`

	BriefDescription string        `xml:"brief_description"`
	Description      string        `xml:"description"`
	Tutorials        []xmlLink     `xml:"tutorials>link"`
	Methods          []xmlMethod   `xml:"methods>method"`
	Members          []xmlMember   `xml:"members>member"`
	Signals          []xmlSignal   `xml:"signals>signal"`
	Constants        []xmlConstant `xml:"constants>constant"`
}

type xmlLink struct {
	Title string `xml:"title,attr"`
	Value string `xml:",chardata"`
}

type xmlMethod struct {
	Name        string        `xml:"name,attr"`
	Return      xmlReturn     `xml:"return"`
	Arguments   []xmlArgument `xml:"param"`
	Description string        `xml:"description"`
}

type xmlReturn struct {
	Type string `xml:"type,attr"`
}

type xmlArgument struct {
	Index   int    `xml:"index,attr"`
	Name    string `xml:"name,attr"`
	Type    string `xml:"type,attr"`
	Default string `xml:"default,attr"`
}

type xmlSignal struct {
	Name        string        `xml:"name,attr"`
	Param       []xmlArgument `xml:"param"`
	Description string        `xml:"description"`
}

type xmlMember struct {
	Name        string `xml:"name,attr"`
	Type        string `xml:"type,attr"`
	Setter      string `xml:"setter,attr"`
	Getter      string `xml:"getter,attr"`
	Description string `xml:",chardata"`
}

type xmlConstant struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	Enum        string `xml:"enum,attr"`
	Description string `xml:",chardata"`
}
