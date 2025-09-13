/*
[gdscript]
undo_redo.create_action("Add object")

# DO
undo_redo.add_do_method(_create_object)
undo_redo.add_do_method(_add_object_to_singleton)

# UNDO
undo_redo.add_undo_method(_remove_object_from_singleton)
undo_redo.add_undo_method(_destroy_that_object)

undo_redo.commit_action()
[/gdscript]
[csharp]
_undo_redo.CreateAction("Add object");

// DO
_undo_redo.AddDoMethod(new Callable(this, MethodName.CreateObject));
_undo_redo.AddDoMethod(new Callable(this, MethodName.AddObjectToSingleton));

// UNDO
_undo_redo.AddUndoMethod(new Callable(this, MethodName.RemoveObjectFromSingleton));
_undo_redo.AddUndoMethod(new Callable(this, MethodName.DestroyThatObject));

_undo_redo.CommitAction();
[/csharp]
*/

package main

func ExamplesUndoRedo() {
	undoRedo.CreateAction("Add object")

	// DO
	undoRedo.AddDoMethod(func() {}) // _create_object
	undoRedo.AddDoMethod(func() {}) // _add_object_to_singleton

	// UNDO
	undoRedo.AddUndoMethod(func() {}) // _remove_object_from_singleton
	undoRedo.AddUndoMethod(func() {}) // _destroy_that_object

	undoRedo.CommitAction()
}
