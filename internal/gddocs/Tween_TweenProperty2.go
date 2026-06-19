/*
[gdscript]
var tween = create_tween()
tween.tween_property($Sprite, "position", Vector2.RIGHT * 300, 1.0).as_relative().set_trans(Tween.TRANS_SINE)
tween.tween_property($Sprite, "position", Vector2.RIGHT * 300, 1.0).as_relative().from_current().set_trans(Tween.TRANS_EXPO)
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenProperty(GetNode("Sprite"), "position", Vector2.Right * 300.0f, 1.0f).AsRelative().SetTrans(Tween.TransitionType.Sine);
tween.TweenProperty(GetNode("Sprite"), "position", Vector2.Right * 300.0f, 1.0f).AsRelative().FromCurrent().SetTrans(Tween.TransitionType.Expo);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/Tween"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenPropertyRelative(node Node.Instance, sprite Sprite2D.Instance) {
	var tween = node.CreateTween()
	PropertyTweener.Make(tween, sprite.AsObject(), "position", Vector2.New(300, 0), 1).AsRelative().SetTrans(Tween.TransSine)
	PropertyTweener.Make(tween, sprite.AsObject(), "position", Vector2.New(300, 0), 1).AsRelative().FromCurrent().SetTrans(Tween.TransExpo)
}
