/*
[gdscript]
func some_function():
	print("start")
	await get_tree().create_timer(1.0).timeout
	print("end")
[/gdscript]
[csharp]
public async Task SomeFunction()
{
	GD.Print("start");
	await ToSignal(GetTree().CreateTimer(1.0f), SceneTreeTimer.SignalName.Timeout);
	GD.Print("end");
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/SceneTree"
	"graphics.gd/variant/Signal"
)

func SceneTree_CreateTimer() {
	fmt.Println("start")
	SceneTree.Get(node).CreateTimer(1.0).OnTimeout(func() {
		fmt.Println("end")
	}, Signal.OneShot)
}
