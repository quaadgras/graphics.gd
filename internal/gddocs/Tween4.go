/*
[gdscript]
var tween = create_tween()
for sprite in get_children():
	tween.tween_property(sprite, "position", Vector2(0, 0), 1.0)
[/gdscript]
[csharp]
Tween tween = CreateTween();
foreach (Node sprite in GetChildren())
	tween.TweenProperty(sprite, "position", Vector2.Zero, 1.0f);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenChildren(node Node.Instance) {
	var tween = node.CreateTween()
	for _, sprite := range node.GetChildren() {
		PropertyTweener.Make(tween, sprite.AsObject(), "position", Vector2.New(0, 0), 1)
	}
}
