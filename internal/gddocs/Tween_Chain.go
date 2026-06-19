/*
[gdscript]
var tween = create_tween().set_parallel(true)
tween.tween_property(...)
tween.tween_property(...) # Will run parallelly with above.
tween.chain().tween_property(...) # Will run after two above are finished.
[/gdscript]
[csharp]
Tween tween = CreateTween().SetParallel(true);
tween.TweenProperty(...);
tween.TweenProperty(...); // Will run parallelly with above.
tween.Chain().TweenProperty(...); // Will run after two above are finished.
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

func ExampleTweenChain(node Node.Instance) {
	var tween = node.CreateTween().SetParallel()
	PropertyTweener.Make(tween, node.AsObject(), "position", Vector2.New(100, 0), 1)
	PropertyTweener.Make(tween, node.AsObject(), "modulate", Color.W3C.Red, 1)    // Will run parallelly with above.
	PropertyTweener.Make(tween.Chain(), node.AsObject(), "scale", Vector2.One, 1) // Will run after two above are finished.
}
