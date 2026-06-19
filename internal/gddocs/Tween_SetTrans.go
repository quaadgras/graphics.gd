/*
var tween = create_tween()
tween.tween_property(self, "position", Vector2(300, 0), 0.5) # Uses TRANS_LINEAR.
tween.set_trans(Tween.TRANS_SINE)
tween.tween_property(self, "rotation_degrees", 45.0, 0.5) # Uses TRANS_SINE.
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/Tween"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenSetTransition(node Node.Instance) {
	var tween = node.CreateTween()
	PropertyTweener.Make(tween, node.AsObject(), "position", Vector2.New(300, 0), 0.5) // Uses TransLinear.
	tween.SetTrans(Tween.TransSine)
	PropertyTweener.Make(tween, node.AsObject(), "rotation_degrees", 45.0, 0.5) // Uses TransSine.
}
