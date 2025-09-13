/*
[gdscript]
@tool # Needed so it runs in editor.
extends EditorScenePostImport

# This sample changes all node names.
# Called right after the scene is imported and gets the root node.
func _post_import(scene):
    # Change all node names to "modified_[oldnodename]"
    iterate(scene)
    return scene # Remember to return the imported scene

func iterate(node):
    if node != null:
        node.name = "modified_" + node.name
        for child in node.get_children():
            iterate(child)
[/gdscript]
[csharp]
using Godot;

// This sample changes all node names.
// Called right after the scene is imported and gets the root node.
[Tool]
public partial class NodeRenamer : EditorScenePostImport
{
    public override GodotObject _PostImport(Node scene)
    {
        // Change all node names to "modified_[oldnodename]"
        Iterate(scene);
        return scene; // Remember to return the imported scene
    }

    public void Iterate(Node node)
    {
        if (node != null)
        {
            node.Name = $"modified_{node.Name}";
            foreach (Node child in node.GetChildren())
            {
                Iterate(child);
            }
        }
    }
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/EditorScenePostImport"
	"graphics.gd/classdb/Node"
)

type NodeRenamer struct {
	EditorScenePostImport.Extension[NodeRenamer]
}

func (n *NodeRenamer) PostImport(scene Node.Instance) Node.Instance {
	n.Iterate(scene)
	return scene
}

func (n *NodeRenamer) Iterate(node Node.Instance) {
	if node != Node.Nil {
		node.SetName("modified_" + node.Name())
		for _, child := range node.GetChildren() {
			n.Iterate(child)
		}
	}
}
