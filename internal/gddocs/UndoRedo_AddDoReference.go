/*
var node = Node2D.new()
undo_redo.create_action("Add node")
undo_redo.add_do_method(add_child.bind(node))
undo_redo.add_do_reference(node)
undo_redo.add_undo_method(remove_child.bind(node))
undo_redo.commit_action()
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/UndoRedo"
)

func ExampleUndoRedoAddDoReference(self Node.Instance, undoRedo UndoRedo.Instance) {
	var node = Node2D.New()
	undoRedo.CreateAction("Add node")
	undoRedo.AddDoMethod(func() { self.AddChild(node.AsNode()) })
	undoRedo.AddDoReference(node.AsObject())
	undoRedo.AddUndoMethod(func() { self.RemoveChild(node.AsNode()) })
	undoRedo.CommitAction()
}
