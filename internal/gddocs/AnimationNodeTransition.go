/*
[gdscript]
# Play child animation connected to "state_2" port.
animation_tree.set("parameters/Transition/transition_request", "state_2")
# Alternative syntax (same result as above).
animation_tree["parameters/Transition/transition_request"] = "state_2"

# Get current state name (read-only).
animation_tree.get("parameters/Transition/current_state")
# Alternative syntax (same result as above).
animation_tree["parameters/Transition/current_state"]

# Get current state index (read-only).
animation_tree.get("parameters/Transition/current_index")
# Alternative syntax (same result as above).
animation_tree["parameters/Transition/current_index"]
[/gdscript]
[csharp]
// Play child animation connected to "state_2" port.
animationTree.Set("parameters/Transition/transition_request", "state_2");

// Get current state name (read-only).
animationTree.Get("parameters/Transition/current_state");

// Get current state index (read-only).
animationTree.Get("parameters/Transition/current_index");
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AnimationTree"
	"graphics.gd/variant/Object"
)

func ExampleAnimationNodeTransition(tree AnimationTree.Instance) {
	// Play child animation connected to "state_2" port.
	Object.Set(tree, "parameters/Transition/transition_request", "state_2")

	// Get current state name (read-only).
	Object.Get(tree, "parameters/Transition/current_state")

	// Get current state index (read-only).
	Object.Get(tree, "parameters/Transition/current_index")
}
