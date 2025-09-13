/*
[gdscript]
# Use load() instead of preload() if the path isn't known at compile-time.
var scene = preload("res://scene.tscn").instantiate()
# Add the node as a child of the node the script is attached to.
add_child(scene)
[/gdscript]
[csharp]
// C# has no preload, so you have to always use ResourceLoader.Load<PackedScene>().
var scene = ResourceLoader.Load<PackedScene>("res://scene.tscn").Instantiate();
// Add the node as a child of the node the script is attached to.
AddChild(scene);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
)

func ExamplePackedSceneLoad(parent Node.Instance) {
	var scene = Resource.Load[PackedScene.Instance]("res://scene.tscn").Instantiate()
	parent.AddChild(scene)
}
