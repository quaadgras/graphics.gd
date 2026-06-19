/*
[gdscript]
var tween = create_tween()
tween.tween_property(...)
tween.parallel().tween_property(...)
tween.parallel().tween_property(...)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(...);
tween.Parallel().TweenProperty(...);
tween.Parallel().TweenProperty(...);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenParallel(node Node.Instance) {
	var tween = node.CreateTween()
	PropertyTweener.Make(tween, node.AsObject(), "position", Vector2.New(100, 0), 1)
	PropertyTweener.Make(tween.Parallel(), node.AsObject(), "modulate", Color.W3C.Red, 1)
	PropertyTweener.Make(tween.Parallel(), node.AsObject(), "scale", Vector2.One, 1)
}
