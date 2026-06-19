/*
[gdscript]
var tween = get_tree().create_tween().bind_node(self).set_trans(Tween.TRANS_ELASTIC)
tween.tween_property($Sprite, "modulate", Color.RED, 1.0)
tween.tween_property($Sprite, "scale", Vector2(), 1.0)
tween.tween_callback($Sprite.queue_free)
[/gdscript]
[csharp]
var tween = GetTree().CreateTween().BindNode(this).SetTrans(Tween.TransitionType.Elastic);
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
	"graphics.gd/classdb/Tween"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenBound(node Node.Instance, sprite Sprite2D.Instance) {
	var tween = Tween.Instance(Tween.Advanced(SceneTree.Get(node).CreateTween()).BindNode(node)).SetTrans(Tween.TransElastic)
	PropertyTweener.Make(tween, sprite.AsObject(), "modulate", Color.W3C.Red, 1)
	PropertyTweener.Make(tween, sprite.AsObject(), "scale", Vector2.Zero, 1)
	tween.TweenCallback(sprite.AsNode().QueueFree)
}
