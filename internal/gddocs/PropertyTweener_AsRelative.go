/*
[gdscript]
var tween = get_tree().create_tween()
tween.tween_property(self, "position", Vector2.RIGHT * 100, 1).as_relative()
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween();
tween.TweenProperty(this, "position", Vector2.Right * 100.0f, 1.0f).AsRelative();
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/variant/Vector2"
)

func PropertyTweener_AsRelative() {
	var tween = SceneTree.Get(node).CreateTween()
	PropertyTweener.Make(tween, self, "position", Vector2.MulX(Vector2.Right, 100), 1).AsRelative()
}
