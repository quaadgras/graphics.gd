/*
[gdscript]
var tween = create_tween()
tween.tween_property($Sprite, "position", Vector2(100, 200), 1.0)
tween.tween_property($Sprite, "position", Vector2(200, 300), 1.0)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(GetNode("Sprite"), "position", new Vector2(100.0f, 200.0f), 1.0f);
tween.TweenProperty(GetNode("Sprite"), "position", new Vector2(200.0f, 300.0f), 1.0f);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenProperty(node Node.Instance, sprite Sprite2D.Instance) {
	var tween = node.CreateTween()
	PropertyTweener.Make(tween, sprite.AsObject(), "position", Vector2.New(100, 200), 1)
	PropertyTweener.Make(tween, sprite.AsObject(), "position", Vector2.New(200, 300), 1)
}
