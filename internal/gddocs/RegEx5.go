/*
var regex = RegEx.new()
regex.compile("\\S+") # Negated whitespace character class.
var results = []
for result in regex.search_all("One  Two \n\tThree"):
    results.push_back(result.get_string())
# The `results` array now contains "One", "Two", and "Three".
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/RegEx"
)

func ExampleRegExSplit() {
	var regex = RegEx.New()
	regex.Compile(`\S+`) // Negated whitespace character class.
	var results = []string{}
	for _, result := range regex.SearchAll("One  Two \n\tThree") {
		results = append(results, result.GetString())
	}
	fmt.Println(results)
}
