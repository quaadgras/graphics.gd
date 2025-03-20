package SceneTree

import (
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Window"
	"graphics.gd/variant/Object"
)

// Add the given node to the scene tree.
func Add(node Node.Any) {
	if tree, ok := Object.Is[Instance](MainLoop.Instance(Engine.GetMainLoop())); ok {
		Window.Instance(tree.Root()).AsNode().AddChild(node.AsNode())
	}
}

// AddNamed adds the given node to the scene tree with the given name.
func AddNamed(name string, node Node.Any) {
	if tree, ok := Object.Is[Instance](MainLoop.Instance(Engine.GetMainLoop())); ok {
		node := node.AsNode()
		node.SetName(name)
		Window.Instance(tree.Root()).AsNode().AddChild(node)
	}
}
