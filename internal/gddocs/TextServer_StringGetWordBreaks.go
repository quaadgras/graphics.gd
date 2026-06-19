/*
var ts = TextServerManager.get_primary_interface()
# Corresponds to the substrings "The", "Godot", "Engine", and "4".
print(ts.string_get_word_breaks("The Godot Engine, 4")) # Prints [0, 3, 4, 9, 10, 16, 18, 19]
# Corresponds to the substrings "The", "Godot", "Engin", and "e, 4".
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 5)) # Prints [0, 3, 4, 9, 10, 15, 15, 19]
# Corresponds to the substrings "The Godot" and "Engine, 4".
print(ts.string_get_word_breaks("The Godot Engine, 4", "en", 10)) # Prints [0, 9, 10, 19]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/TextServerManager"
)

func ExampleStringGetWordBreaks() {
	var ts = TextServerManager.GetPrimaryInterface()
	fmt.Println(ts.StringGetWordBreaks("The Godot Engine, 4"))                      // Prints [0 3 4 9 10 16 18 19]
	fmt.Println(ts.MoreArgs().StringGetWordBreaks("The Godot Engine, 4", "en", 5))  // Prints [0 3 4 9 10 15 15 19]
	fmt.Println(ts.MoreArgs().StringGetWordBreaks("The Godot Engine, 4", "en", 10)) // Prints [0 9 10 19]
}
