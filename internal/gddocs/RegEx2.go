/*
var regex = RegEx.new()
regex.compile("\\w-(\\d+)")
var result = regex.search("abc n-0123")
if result:
	print(result.get_string()) # Would print n-0123
*/

package main

import (
	"graphics.gd/classdb/RegEx"
	"graphics.gd/classdb/RegExMatch"
)

func ExampleRegExSearch() {
	var regex = RegEx.New()
	regex.Compile(`\w-(\d+)`)
	result := regex.Search("abc n-0123")
	if result != RegExMatch.Nil {
		print(result.GetString()) // Would print n-0123
	}
}
