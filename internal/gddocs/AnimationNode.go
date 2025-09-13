/*
var current_length = $AnimationTree[parameters/AnimationNodeName/current_length]
var current_position = $AnimationTree[parameters/AnimationNodeName/current_position]
var current_delta = $AnimationTree[parameters/AnimationNodeName/current_delta]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/AnimationTree"
	"graphics.gd/variant/Object"
)

func ExampleAnimationNode(tree AnimationTree.Instance) {
	var current_length = Object.Get(tree, "parameters/AnimationNodeName/current_length")
	var current_position = Object.Get(tree, "parameters/AnimationNodeName/current_position")
	var current_delta = Object.Get(tree, "parameters/AnimationNodeName/current_delta")
	fmt.Print(current_length, current_position, current_delta)
}
