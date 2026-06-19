/*
var tween = create_tween()

# Will move from 0 to 500 over 1 second.
position.x = 0.0
tween.tween_property(self, "position:x", 500, 1.0)

# Will be at (about) 250 when the timer finishes.
await get_tree().create_timer(0.5).timeout

# Will now move from (about) 250 to 500 over 1 second,
# thus at half the speed as before.
tween.stop()
tween.play()
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
)

func ExampleTweenStop(node Node.Instance) {
	var tween = node.CreateTween()
	// Will move from 0 to 500 over 1 second.
	PropertyTweener.Make(tween, node.AsObject(), "position:x", 500, 1)
	// await get_tree().create_timer(0.5).timeout
	tween.Stop()
	tween.Play()
}
