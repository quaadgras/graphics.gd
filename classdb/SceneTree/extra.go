package SceneTree

import (
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Object"
)

// Add the given node to the scene tree.
func Add(node Node.Any) {
	if tree, ok := Object.As[Instance](Engine.GetMainLoop()); ok {
		tree.Root().AsNode().AddChild(node.AsNode())
	}
}

// AddNamed adds the given node to the scene tree with the given name.
func AddNamed(name string, node Node.Any) {
	if tree, ok := Object.As[Instance](Engine.GetMainLoop()); ok {
		node := node.AsNode()
		node.SetName(name)
		tree.Root().AsNode().AddChild(node)
	}
}
