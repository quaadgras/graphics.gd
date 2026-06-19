/*
# Prints "01 03 0 3f 42"
for result in regex.search_all("d01, d03, d0c, x3f and x42"):
	print(result.get_string("digit"))
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/RegEx"
)

func ExampleRegExSearchAllNamed(regex RegEx.Instance) {
	// Prints "01 03 0 3f 42"
	for _, result := range regex.SearchAll("d01, d03, d0c, x3f and x42") {
		fmt.Println(result.MoreArgs().GetString("digit"))
	}
}
