/*
var data = JSON.parse_string(json_string) # Returns null if parsing failed.
*/

package main

import "graphics.gd/classdb/JSON"

func ExampleParseString(json_string string) {
	var data = JSON.ParseString(json_string) // Returns null if parsing failed.
	_ = data
}
