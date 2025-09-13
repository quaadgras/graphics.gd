/*
var rng = RandomNumberGenerator.new()
func _ready():
    var my_random_number = rng.randf_range(-10.0, 10.0)
*/

package main

import "graphics.gd/classdb/RandomNumberGenerator"

func ExampleRandomNumberGenerator() {
	var rng = RandomNumberGenerator.New()
	var my_random_number = rng.RandfRange(-10.0, 10.0)
	_ = my_random_number
}
