package SceneTree

import (
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Window"
	gd "graphics.gd/internal"
	"graphics.gd/variant/Object"
)

// Add queues the given node to be added to the scene tree.
func Add(node Node.Any) {
	gd.NewCallable(func() {
		if tree, ok := Object.Is[Instance](MainLoop.Instance(Engine.GetMainLoop())); ok {
			Window.Instance(tree.Root()).AsNode().AddChild(node.AsNode())
		}
	}).CallDeferred()
}
