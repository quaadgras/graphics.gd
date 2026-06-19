/*
@tool
extends EditorTranslationParserPlugin

func _customize_strings(strings):
	# Add new string.
	strings.append(["Test 1", "context", "test 1 plurals", "test 1 comment"])

	# Remove all strings that begin with $.
	strings = strings.filter(func(s): return not s[0].begins_with("$"))

	return strings
*/

package main

import (
	"strings"

	"graphics.gd/classdb/EditorTranslationParserPlugin"
)

type editorTranslationParserCustomizeStrings struct {
	EditorTranslationParserPlugin.Extension[editorTranslationParserCustomizeStrings]
}

func (n editorTranslationParserCustomizeStrings) CustomizeStrings(list [][]string) [][]string {
	// Add new string.
	list = append(list, []string{"Test 1", "context", "test 1 plurals", "test 1 comment"})

	// Remove all strings that begin with $.
	var filtered [][]string
	for _, s := range list {
		if len(s) > 0 && !strings.HasPrefix(s[0], "$") {
			filtered = append(filtered, s)
		}
	}
	return filtered
}
