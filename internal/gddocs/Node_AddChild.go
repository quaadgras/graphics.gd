/*
[gdscript]
var child_node = get_child(0)
if child_node.get_parent():
	child_node.get_parent().remove_child(child_node)
add_child(child_node)
[/gdscript]
[csharp]
Node childNode = GetChild(0);
if (childNode.GetParent() != null)
{
	childNode.GetParent().RemoveChild(childNode);
}
AddChild(childNode);
[/csharp]
*/

package main

import "graphics.gd/classdb/Node"

func Node_AddChild() {
	var childNode = node.GetChild(0)
	if childNode.GetParent() != Node.Nil {
		childNode.GetParent().RemoveChild(childNode)
	}
	node.AddChild(childNode)
}
