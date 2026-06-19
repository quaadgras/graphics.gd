/*
var tween = create_tween()
tween.tween_callback(walk_to.bind(600.0))
tween.tween_await(destination_reached)
tween.tween_callback(say_dialogue.bind("Good day, sir!"))
tween.tween_await(dialogue_closed)
tween.tween_callback(walk_to.bind(0.0))
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Signal"
)

func ExampleTweenAwaitDialogue(node Node.Instance, destinationReached, dialogueClosed Signal.Any, walkTo func(float64), sayDialogue func(string)) {
	var tween = node.CreateTween()
	tween.TweenCallback(func() { walkTo(600.0) })
	tween.TweenAwait(destinationReached)
	tween.TweenCallback(func() { sayDialogue("Good day, sir!") })
	tween.TweenAwait(dialogueClosed)
	tween.TweenCallback(func() { walkTo(0.0) })
}
