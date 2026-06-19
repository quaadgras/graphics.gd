/*
[gdscript]
var rng = RandomNumberGenerator.new()

var my_array = ["one", "two", "three", "four"]
var weights = PackedFloat32Array([0.5, 1, 1, 2])

# Prints one of the four elements in `my_array`.
# It is more likely to print "four", and less likely to print "one".
print(my_array[rng.rand_weighted(weights)])
[/gdscript]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/RandomNumberGenerator"
)

func ExampleRandWeighted() {
	var rng = RandomNumberGenerator.New()
	var myArray = []string{"one", "two", "three", "four"}
	var weights = []float32{0.5, 1, 1, 2}
	// Prints one of the four elements in myArray.
	// It is more likely to print "four", and less likely to print "one".
	fmt.Println(myArray[rng.RandWeighted(weights)])
}
