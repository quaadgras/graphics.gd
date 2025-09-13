/*
var regex = RegEx.new()
regex.compile("\\w-(\\d+)")
*/

package main

import "graphics.gd/classdb/RegEx"

func ExampleRegEx() {
	var regex = RegEx.New()
	regex.Compile(`\w-(\d+)`)
}
