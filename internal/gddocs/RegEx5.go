/*
var regex = RegEx.create_from_string("\\S+") # Negated whitespace character class.
var results = []
for result in regex.search_all("One  Two \n\tThree"):
	results.push_back(result.get_string())
print(results) # Prints ["One", "Two", "Three"]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/RegEx"
)

func ExampleRegExSearchAll() {
	var regex = RegEx.CreateFromString(`\S+`) // Negated whitespace character class.
	var results []string
	for _, result := range regex.SearchAll("One  Two \n\tThree") {
		results = append(results, result.GetString())
	}
	fmt.Println(results) // Prints ["One" "Two" "Three"]
}
