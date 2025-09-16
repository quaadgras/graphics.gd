/*
var regex = RegEx.new()
regex.compile("d(?<digit>[0-9]+)|x(?<digit>[0-9a-f]+)")
var result = regex.search("the number is x2f")
if result:
	print(result.get_string("digit")) # Would print 2f
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/RegEx"
	"graphics.gd/classdb/RegExMatch"
)

func ExampleRegExCapture() {
	var regex = RegEx.New()
	regex.Compile(`d(?<digit>[0-9]+)|x(?<digit>[0-9a-f]+)`)
	result := regex.Search("the number is x2f")
	if result != RegExMatch.Nil {
		fmt.Print(RegExMatch.Expanded(result).GetString("digit")) // Would print 2f
	}
}
