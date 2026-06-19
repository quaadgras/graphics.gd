/*
var regex = RegEx.create_from_string("\\w-(\\d+)")
var result = regex.search("abc n-0123")
if result:
	print(result.get_string()) # Prints "n-0123"
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/RegEx"
	"graphics.gd/classdb/RegExMatch"
)

func ExampleRegExSearch() {
	var regex = RegEx.CreateFromString(`\w-(\d+)`)
	var result = regex.Search("abc n-0123")
	if result != RegExMatch.Nil {
		fmt.Println(result.GetString()) // Prints "n-0123"
	}
}
