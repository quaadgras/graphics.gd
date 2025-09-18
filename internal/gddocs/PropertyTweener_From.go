/*
[gdscript]
var tween = get_tree().create_tween()
tween.tween_property(self, "position", Vector2(200, 100), 1).from(Vector2(100, 100))
[/gdscript]
[csharp]
Tween tween = GetTree().CreateTween();
tween.TweenProperty(this, "position", new Vector2(200.0f, 100.0f), 1.0f).From(new Vector2(100.0f, 100.0f));
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

var self Object.Instance

func PropertyTweener_From() {
	var tween = SceneTree.Get(node).CreateTween()
	PropertyTweener.Make(tween, self, "position", Vector2.New(200, 100), 1).From(Vector2.New(100, 100))
}
