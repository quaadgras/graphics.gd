/*
# Subtween will rotate the object.
var subtween = create_tween()
subtween.tween_property(self, "rotation_degrees", 45.0, 1.0)
subtween.tween_property(self, "rotation_degrees", 0.0, 1.0)

# Parent tween will execute the subtween as one of its steps.
var tween = create_tween()
tween.tween_property(self, "position:x", 500, 3.0)
tween.tween_subtween(subtween)
tween.tween_property(self, "position:x", 300, 2.0)
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
)

func ExampleTweenSubtween(node Node.Instance) {
	// Subtween will rotate the object.
	var subtween = node.CreateTween()
	PropertyTweener.Make(subtween, node.AsObject(), "rotation_degrees", 45.0, 1)
	PropertyTweener.Make(subtween, node.AsObject(), "rotation_degrees", 0.0, 1)

	// Parent tween will execute the subtween as one of its steps.
	var tween = node.CreateTween()
	PropertyTweener.Make(tween, node.AsObject(), "position:x", 500, 3)
	tween.TweenSubtween(subtween)
	PropertyTweener.Make(tween, node.AsObject(), "position:x", 300, 2)
}
