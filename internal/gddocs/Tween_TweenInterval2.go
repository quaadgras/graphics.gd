/*
[gdscript]
var tween = create_tween().set_loops()
tween.tween_property($Sprite, "position:x", 200.0, 1.0).as_relative()
tween.tween_callback(jump)
tween.tween_interval(2)
tween.tween_property($Sprite, "position:x", -200.0, 1.0).as_relative()
tween.tween_callback(jump)
tween.tween_interval(2)
[/gdscript]
[csharp]
Tween tween = CreateTween().SetLoops();
tween.TweenProperty(GetNode("Sprite"), "position:x", 200.0f, 1.0f).AsRelative();
tween.TweenCallback(Callable.From(Jump));
tween.TweenInterval(2.0f);
tween.TweenProperty(GetNode("Sprite"), "position:x", -200.0f, 1.0f).AsRelative();
tween.TweenCallback(Callable.From(Jump));
tween.TweenInterval(2.0f);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/Sprite2D"
)

func ExampleTweenIntervalLoop(node Node.Instance, sprite Sprite2D.Instance, jump func()) {
	var tween = node.CreateTween().SetLoops()
	PropertyTweener.Make(tween, sprite.AsObject(), "position:x", 200.0, 1).AsRelative()
	tween.TweenCallback(jump)
	tween.TweenInterval(2)
	PropertyTweener.Make(tween, sprite.AsObject(), "position:x", -200.0, 1).AsRelative()
	tween.TweenCallback(jump)
	tween.TweenInterval(2)
}
