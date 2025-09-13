/*
[gdscript]
# Play child animation from the start.
animation_tree.set("parameters/TimeSeek/seek_request", 0.0)
# Alternative syntax (same result as above).
animation_tree["parameters/TimeSeek/seek_request"] = 0.0

# Play child animation from 12 second timestamp.
animation_tree.set("parameters/TimeSeek/seek_request", 12.0)
# Alternative syntax (same result as above).
animation_tree["parameters/TimeSeek/seek_request"] = 12.0
[/gdscript]
[csharp]
// Play child animation from the start.
animationTree.Set("parameters/TimeSeek/seek_request", 0.0);

// Play child animation from 12 second timestamp.
animationTree.Set("parameters/TimeSeek/seek_request", 12.0);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/AnimationTree"
	"graphics.gd/variant/Object"
)

func ExampleAnimationNodeTimeSeek(tree AnimationTree.Instance) {
	// Play child animation from the start.
	Object.Set(tree, "parameters/TimeSeek/seek_request", 0.0)

	// Play child animation from 12 second timestamp.
	Object.Set(tree, "parameters/TimeSeek/seek_request", 12.0)
}
