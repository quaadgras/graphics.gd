/*
for result in regex.search_all("d01, d03, d0c, x3f and x42"):
	print(result.get_string("digit"))
# Would print 01 03 0 3f 42
*/

package main

import (
	"graphics.gd/classdb/RegEx"
)

func ExampleRegExSearchAll() {
	var regex = RegEx.New()
	regex.Compile(`d(?<digit>[0-9]+)|x(?<digit>[0-9a-f]+)`)
	for _, result := range regex.SearchAll("the numbers are d42 and x2f") {
		result.MoreArgs().GetString("digit")
	}
}
