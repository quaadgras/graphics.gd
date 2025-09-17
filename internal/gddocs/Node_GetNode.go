/*
[gdscript]
get_node("Sword")
get_node("Backpack/Dagger")
get_node("../Swamp/Alligator")
get_node("/root/MyGame")
[/gdscript]
[csharp]
GetNode("Sword");
GetNode("Backpack/Dagger");
GetNode("../Swamp/Alligator");
GetNode("/root/MyGame");
[/csharp]
*/

package main

func Node_GetNode() {
	node.GetNode("Sword")
	node.GetNode("Backpack/Dagger")
	node.GetNode("../Swamp/Alligator")
	node.GetNode("/root/MyGame")
}
