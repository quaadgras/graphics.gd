/*
[gdscript]
tween.tween_property(self, "position", Vector2(200, 100), 1).from(position)
tween.tween_property(self, "position", Vector2(200, 100), 1).from_current()
[/gdscript]
[csharp]
tween.TweenProperty(this, "position", new Vector2(200.0f, 100.0f), 1.0f).From(Position);
tween.TweenProperty(this, "position", new Vector2(200.0f, 100.0f), 1.0f).FromCurrent();
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/variant/Vector2"
)

func PropertyTweener_FromCurrent() {
	PropertyTweener.Make(tween, self, "position", Vector2.New(200, 100), 1).From(node2d.Position())
	PropertyTweener.Make(tween, self, "position", Vector2.New(200, 100), 1).FromCurrent()
}
