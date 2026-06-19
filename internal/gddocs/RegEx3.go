/*
var regex = RegEx.create_from_string("d(?<digit>[0-9]+)|x(?<digit>[0-9a-f]+)")
var result = regex.search("the number is x2f")
if result:
	print(result.get_string("digit")) # Prints "2f"
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/RegEx"
	"graphics.gd/classdb/RegExMatch"
)

func ExampleRegExNamedGroup() {
	var regex = RegEx.CreateFromString("d(?<digit>[0-9]+)|x(?<digit>[0-9a-f]+)")
	var result = regex.Search("the number is x2f")
	if result != RegExMatch.Nil {
		fmt.Println(result.MoreArgs().GetString("digit")) // Prints "2f"
	}
}
