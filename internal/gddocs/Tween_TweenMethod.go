/*
[gdscript]
var tween = create_tween()
tween.tween_method(look_at.bind(Vector3.UP), Vector3(-1, 0, -1), Vector3(1, 0, -1), 1.0) # The look_at() method takes up vector as second argument.
[/gdscript]
[csharp]
Tween tween = CreateTween();
tween.TweenMethod(Callable.From((Vector3 target) => LookAt(target, Vector3.Up)), new Vector3(-1.0f, 0.0f, -1.0f), new Vector3(1.0f, 0.0f, -1.0f), 1.0f); // Use lambdas to bind additional arguments for the call.
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/MethodTweener"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Vector3"
)

func ExampleTweenMethod(node Node3D.Instance) {
	var tween = node.AsNode().CreateTween()
	// The look_at() method takes up vector as second argument.
	MethodTweener.Make(tween, Callable.New(func(target Vector3.XYZ) {
		node.LookAt(target)
	}), Vector3.New(-1, 0, -1), Vector3.New(1, 0, -1), 1)
}
