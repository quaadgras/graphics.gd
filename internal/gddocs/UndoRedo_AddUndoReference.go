/*
var node = $Node2D
undo_redo.create_action("Remove node")
undo_redo.add_do_method(remove_child.bind(node))
undo_redo.add_undo_method(add_child.bind(node))
undo_redo.add_undo_reference(node)
undo_redo.commit_action()
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/UndoRedo"
)

func ExampleUndoRedoAddUndoReference(self Node.Instance, node Node2D.Instance, undoRedo UndoRedo.Instance) {
	undoRedo.CreateAction("Remove node")
	undoRedo.AddDoMethod(func() { self.RemoveChild(node.AsNode()) })
	undoRedo.AddUndoMethod(func() { self.AddChild(node.AsNode()) })
	undoRedo.AddUndoReference(node.AsObject())
	undoRedo.CommitAction()
}
