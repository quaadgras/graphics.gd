/*
[gdscript]
var tween = get_tree().create_tween()
tween.tween_callback($Sprite.set_modulate.bind(Color.RED)).set_delay(2)
tween.tween_callback($Sprite.set_modulate.bind(Color.BLUE)).set_delay(2)
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween();
Sprite2D sprite = GetNode<Sprite2D>("Sprite");
tween.TweenCallback(Callable.From(() => sprite.Modulate = Colors.Red)).SetDelay(2.0f);
tween.TweenCallback(Callable.From(() => sprite.Modulate = Colors.Blue)).SetDelay(2.0f);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/variant/Color"
)

func ExampleTweenCallbackBind(node Node.Instance, sprite Sprite2D.Instance) {
	var tween = SceneTree.Get(node).CreateTween()
	tween.TweenCallback(func() { sprite.AsCanvasItem().SetModulate(Color.W3C.Red) }).SetDelay(2)
	tween.TweenCallback(func() { sprite.AsCanvasItem().SetModulate(Color.W3C.Blue) }).SetDelay(2)
}
