/*
[gdscript]
var tween = get_tree().create_tween().set_loops()
tween.tween_callback(shoot).set_delay(1.0)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween().SetLoops();
tween.TweenCallback(Callable.From(Shoot)).SetDelay(1.0f);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
)

func ExampleTweenCallbackDelay(node Node.Instance, shoot func()) {
	var tween = SceneTree.Get(node).CreateTween().SetLoops()
	tween.TweenCallback(shoot).SetDelay(1.0)
}
