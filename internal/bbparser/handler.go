package bbparser

import (
	"fmt"
	"strings"
)

// Tag - structure contains info about bbcode tag
type Tag struct {
	Name       string
	Closing    bool
	Attributes map[string]string
}

// Handler - function type for handling content between bbcode tags
type Handler func(Tag, string) string

// SimpleHandler - handle tags that need just add html tags at the begin and the end
func SimpleHandler(tag Tag, content string) string {
	return fmt.Sprintf("<%s>%s</%s>", tag.Name, content, tag.Name)
}

func TextAlignHandler(tag Tag, content string) string {
	return fmt.Sprintf("<div style=\"text-align:%s;\">%s</div>", tag.Name, content)
}

// FontColorHandler - parse color tag with color attribute
func FontColorHandler(tag Tag, content string) string {
	var colorAttr string
	if v, ok := tag.Attributes[startingAttr]; ok {
		colorAttr = "color=" + v
	}
	return fmt.Sprintf("<font %s>%s</font>", colorAttr, content)
}

// FontSizeHandler - parse size tag with size attribute
func FontSizeHandler(tag Tag, content string) string {
	var sizeAttr string
	if v, ok := tag.Attributes[startingAttr]; ok {
		sizeAttr = "size=" + v
	}
	return fmt.Sprintf("<font %s>%s</font>", sizeAttr, content)
}

// URLHandler - parse url tag with or without attribute
func URLHandler(tag Tag, content string) string {
	var link string
	if v, ok := tag.Attributes[startingAttr]; ok {
		link = v
	} else {
		link = content
	}
	return fmt.Sprintf("<a href=\"%s\">%s</a>", link, content)
}

// ListHandler - parse list with one-line tags (double tags like [li] parsed separated)
func ListHandler(tag Tag, content string) string {
	var listContent strings.Builder
	items := strings.SplitSeq(content, "\n")
	for item := range items {
		trimed := strings.TrimSpace(item)
		if strings.HasPrefix(trimed, "[*]") || strings.HasPrefix(trimed, "*") {
			trimed = strings.TrimPrefix(trimed, "[*]")
			trimed = strings.TrimPrefix(trimed, "*")
			listContent.WriteString(fmt.Sprintf("<li>%s</li>", trimed))
		} else {
			listContent.WriteString(trimed)
		}
	}
	return fmt.Sprintf("<%s>%s</%s>", tag.Name, listContent.String(), tag.Name)
}
