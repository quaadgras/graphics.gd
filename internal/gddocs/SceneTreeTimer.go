/*
[gdscript]
func some_function():
	print("Timer started.")
	await get_tree().create_timer(1.0).timeout
	print("Timer ended.")
[/gdscript]
[csharp]
public async Task SomeFunction()
{
	GD.Print("Timer started.");
	await ToSignal(GetTree().CreateTimer(1.0f), SceneTreeTimer.SignalName.Timeout);
	GD.Print("Timer ended.");
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
)

func ExampleSceneTreeTimer(node Node.Instance) {
	fmt.Println("Timer started.")
	SceneTree.Get(node).CreateTimer(1.0).OnTimeout(func() {
		fmt.Println("Timer ended.")
	})

}
