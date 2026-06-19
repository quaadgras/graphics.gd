/*
var tween = create_tween()
tween.tween_await(signal)
tween.parallel().tween_callback(method_that_emits_signal)
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Signal"
)

func ExampleTweenAwaitSelf(node Node.Instance, signal Signal.Any, methodThatEmitsSignal func()) {
	var tween = node.CreateTween()
	tween.TweenAwait(signal)
	tween.Parallel().TweenCallback(methodThatEmitsSignal)
}
