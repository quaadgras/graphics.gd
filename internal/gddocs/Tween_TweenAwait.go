/*
var tween = create_tween()
tween.tween_callback(launch)
tween.tween_await(collided).set_timeout(4.0)
tween.tween_callback(explode)
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Signal"
)

func ExampleTweenAwait(node Node.Instance, collided Signal.Any, launch, explode func()) {
	var tween = node.CreateTween()
	tween.TweenCallback(launch)
	tween.TweenAwait(collided).SetTimeout(4.0)
	tween.TweenCallback(explode)
}
