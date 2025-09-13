/*
[gdscript]
var tween = get_tree().create_tween()
tween.tween_property($Sprite, "modulate", Color.RED, 1)
tween.tween_property($Sprite, "scale", Vector2(), 1)
tween.tween_callback($Sprite.queue_free)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween();
tween.TweenProperty(GetNode("Sprite"), "modulate", Colors.Red, 1.0f);
tween.TweenProperty(GetNode("Sprite"), "scale", Vector2.Zero, 1.0f);
tween.TweenCallback(Callable.From(GetNode("Sprite").QueueFree));
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

func ExampleTween(node Node.Instance, sprite Sprite2D.Instance) {
	var tween = SceneTree.Get(node).CreateTween()
	PropertyTweener.Make(tween, sprite.AsObject(), "modulate", Color.W3C.Red, 1)
	PropertyTweener.Make(tween, sprite.AsObject(), "scale", Vector2.Zero, 1)
	tween.TweenCallback(sprite.AsNode().QueueFree)
}
