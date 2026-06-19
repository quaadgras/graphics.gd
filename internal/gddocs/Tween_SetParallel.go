/*
tween.tween_property(self, "position", Vector2(300, 0), 0.5)
tween.set_parallel()
tween.tween_property(self, "modulate", Color.GREEN, 0.5) # Runs together with the position tweener.
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenSetParallel(node Node.Instance) {
	var tween = node.CreateTween()
	PropertyTweener.Make(tween, node.AsObject(), "position", Vector2.New(300, 0), 0.5)
	tween.SetParallel()
	PropertyTweener.Make(tween, node.AsObject(), "modulate", Color.W3C.Green, 0.5) // Runs together with the position tweener.
}
