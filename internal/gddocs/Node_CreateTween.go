/*
[gdscript]
get_tree().create_tween().bind_node(self)
[/gdscript]
[csharp]
GetTree().CreateTween().BindNode(this);
[/csharp]
*/

package main

import "graphics.gd/classdb/SceneTree"

func Node_CreateTween() {
	node.BindToTween(SceneTree.Get(node).CreateTween())
}
