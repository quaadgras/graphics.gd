/*
[gdscript]
# ... some code
await create_tween().tween_interval(2).finished
# ... more code
[/gdscript]
[csharp]
// ... some code
await ToSignal(CreateTween().TweenInterval(2.0f), Tween.SignalName.Finished);
// ... more code
[/csharp]
*/

package main

import "graphics.gd/classdb/Node"

func ExampleTweenInterval(node Node.Instance) {
	// ... some code
	node.CreateTween().TweenInterval(2) // await ....finished
	// ... more code
}
