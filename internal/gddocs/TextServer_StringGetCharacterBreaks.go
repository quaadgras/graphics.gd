/*
var ts = TextServerManager.get_primary_interface()
print(ts.string_get_character_breaks("Test ❤️‍🔥 Test")) # Prints [1, 2, 3, 4, 5, 9, 10, 11, 12, 13, 14]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/TextServerManager"
)

func ExampleStringGetCharacterBreaks() {
	var ts = TextServerManager.GetPrimaryInterface()
	fmt.Println(ts.StringGetCharacterBreaks("Test ❤️‍🔥 Test")) // Prints [1 2 3 4 5 9 10 11 12 13 14]
}
