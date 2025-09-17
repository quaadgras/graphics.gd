/*
# Assuming the following are children of this node, in order:
# First, Middle, Last.

var a = get_child(0).name  # a is "First"
var b = get_child(1).name  # b is "Middle"
var b = get_child(2).name  # b is "Last"
var c = get_child(-1).name # c is "Last"
*/

package main

func Node_GetChild() {
	// Assuming the following are children of this node, in order:
	// First, Middle, Last.
	var a = node.GetChild(0).Name()  // a is "First"
	var b = node.GetChild(1).Name()  // b is "Middle"
	b = node.GetChild(2).Name()      // b is "Last"
	var c = node.GetChild(-1).Name() // c is "Last"
	_ = a
	_ = b
	_ = c
}
