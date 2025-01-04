package SceneTree

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Window"
	gd "graphics.gd/internal"
)

// Add queues the given node to be added to the scene tree.
func Add(node Node.Any) {
	gd.NewCallable(func() {
		tree, ok := classdb.As[Instance](MainLoop.Instance(Engine.GetMainLoop()))
		if ok {
			Window.Instance(tree.Root()).AsNode().AddChild(node.AsNode())
		}
	}).CallDeferred()
}
