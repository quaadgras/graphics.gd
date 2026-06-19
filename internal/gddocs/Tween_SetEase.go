/*
[gdscript]
var tween = create_tween()
tween.tween_property(self, "position", Vector2(300, 0), 0.5) # Uses EASE_IN_OUT.
tween.set_ease(Tween.EASE_IN)
tween.tween_property(self, "rotation_degrees", 45.0, 0.5) # Uses EASE_IN.
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(this, "position", new Vector2(300, 0), 0.5); // Uses EaseType.InOut.
tween.SetEase(Tween.EaseType.In);
tween.TweenProperty(this, "rotation_degrees", 45.0, 0.5); // Uses EaseType.In.
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/Tween"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenSetEase(node Node.Instance) {
	var tween = node.CreateTween()
	PropertyTweener.Make(tween, node.AsObject(), "position", Vector2.New(300, 0), 0.5) // Uses EaseInOut.
	tween.SetEase(Tween.EaseIn)
	PropertyTweener.Make(tween, node.AsObject(), "rotation_degrees", 45.0, 0.5) // Uses EaseIn.
}
