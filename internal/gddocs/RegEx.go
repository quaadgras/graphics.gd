/*
var regex = RegEx.new()
regex.compile("\\w-(\\d+)")
# Shorthand to create and compile a regex (used in the examples below):
var regex2 = RegEx.create_from_string("\\w-(\\d+)")
*/

package main

import "graphics.gd/classdb/RegEx"

func ExampleRegEx() {
	var regex = RegEx.New()
	regex.Compile(`\w-(\d+)`)
	// Shorthand to create and compile a regex (used in the examples below):
	var regex2 = RegEx.CreateFromString("\\w-(\\d+)")
	_ = regex2
}
