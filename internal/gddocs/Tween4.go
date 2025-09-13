/*
[gdscript]
var tween = create_tween()
for sprite in get_children():
    tween.tween_property(sprite, "position", Vector2(0, 0), 1)
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

func ExampleTweenObjects(node Node.Instance) {
	var tween = node.CreateTween()
	for _, sprite := range node.GetChildren() {
		PropertyTweener.Make(tween, sprite.AsObject(), "position", Vector2.Zero, 1.0)
	}
}
