package gdjson

import (
	"fmt"
	"os"
	"strings"

	"graphics.gd/internal/bbparser"
)

func DocsToGoDoc(docs string, className, codeblock string) string {
	docs = strings.ReplaceAll(docs, "*/", "")

	parser := bbparser.New()

	var block int
	code := func(tag bbparser.Tag, body string) string {
		block++
		if codeblock == "" {
			return body
		}
		name := fmt.Sprintf("%s.go", className)
		if block > 1 {
			name = fmt.Sprintf("%s%d.go", className, block)
		}
		trans, err := os.ReadFile("./internal/gddocs/" + name)
		s := string(trans)
		_, s, _ = strings.Cut(s, "package main\n")
		if err == nil {
			return "\tpackage main\n" + strings.ReplaceAll(string(s), "\n", "\n\t")
		}
		return ""
	}

	parser.AddTag("codeblocks", code)
	parser.AddTag("codeblocks", code)
	parser.AddTag("param", func(tag bbparser.Tag, body string) string {
		for key := range tag.Attributes {
			return fmt.Sprintf("'%s'", key)
		}
		return ""
	})
	parser.AddTag("method", func(tag bbparser.Tag, body string) string {
		for key := range tag.Attributes {
			return fmt.Sprintf("[Instance.%s]", ConvertName(key))
		}
		return ""
	})
	parser.AddTag("constant", func(tag bbparser.Tag, body string) string {
		for key := range tag.Attributes {
			return fmt.Sprintf("[%s]", ConvertName(key))
		}
		return ""
	})
	parser.AddTag("b", func(tag bbparser.Tag, body string) string {
		return body
	})
	parser.AddTag("code", func(t bbparser.Tag, s string) string {
		return s
	})

	parser.AddTag("enum", func(tag bbparser.Tag, body string) string {
		for key := range tag.Attributes {
			return fmt.Sprintf("[%s]", key)
		}
		return ""
	})

	parser.AddTag("PackedByteArray", func(tag bbparser.Tag, body string) string { return "byte slice" })

	return parser.Parse(docs)
}
