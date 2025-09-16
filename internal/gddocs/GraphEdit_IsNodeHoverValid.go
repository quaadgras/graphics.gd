/*
[gdscript]
func _is_node_hover_valid(from, from_port, to, to_port):
	return from != to
[/gdscript]
[csharp]
public override bool _IsNodeHoverValid(StringName fromNode, int fromPort, StringName toNode, int toPort)
{
	return fromNode != toNode;
}
[/csharp]
*/

package main

func GraphEdit_IsNodeHoverValid() {
	IsNodeHoverValid := func(from string, from_port int, to string, to_port int) bool {
		return from != to
	}
	_ = IsNodeHoverValid
}
