package SceneTree

import (
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Window"
	"graphics.gd/variant/Object"
)

// Add the given node to the scene tree. If the tree is busy setting up other
// children (for example, the initial scene is still entering the tree during
// the first few frames), the node is added through the engine's deferred
// queue instead and will be inside the tree by the end of a later frame.
func Add(node Node.Any) {
	if tree, ok := Object.As[Instance](Engine.GetMainLoop()); ok {
		if root := tree.Root(); root != Window.Nil {
			n := node.AsNode()
			root.AsNode().AddChild(n)
			if n.GetParent() == Node.Nil {
				// The engine rejects add_child while the root is busy setting
				// up children: the deferred queue flushes once it is idle.
				// A rejected add leaves the node without a parent. Checking
				// IsInsideTree instead would also retry after a successful
				// add from main, where the root itself is not in the tree
				// yet (the engine attaches it right after the MainLoop
				// initialize callback that resumes main), and that retry
				// fails with "already has a parent 'root'".
				Object.Call(root, "call_deferred", "add_child", n)
			}
		}
	}
}

// AddNamed adds the given node to the scene tree with the given name, with
// the same busy-tree fallback as [Add].
func AddNamed(name string, node Node.Any) {
	node.AsNode().SetName(name)
	Add(node)
}
