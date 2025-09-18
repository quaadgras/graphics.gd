/*
# Calls "hide" to all nodes of the "enemies" group, at the end of the frame and in reverse tree order.
get_tree().call_group_flags(
		SceneTree.GROUP_CALL_DEFERRED | SceneTree.GROUP_CALL_REVERSE,
		"enemies", "hide")
*/

package main

import "graphics.gd/classdb/SceneTree"

func SceneTree_CallGroupFlags() {
	// Calls "hide" to all nodes of the "enemies" group, at the end of the frame and in reverse tree order.
	SceneTree.Get(node).CallGroupFlags(
		SceneTree.GroupCallDeferred|SceneTree.GroupCallReverse,
		"enemies", "hide")
}
