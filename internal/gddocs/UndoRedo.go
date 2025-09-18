/*
[gdscript]
var undo_redo = UndoRedo.new()

func do_something():
	pass # Put your code here.

func undo_something():
	pass # Put here the code that reverts what's done by "do_something()".

func _on_my_button_pressed():
	var node = get_node("MyNode2D")
	undo_redo.create_action("Move the node")
	undo_redo.add_do_method(do_something)
	undo_redo.add_undo_method(undo_something)
	undo_redo.add_do_property(node, "position", Vector2(100, 100))
	undo_redo.add_undo_property(node, "position", node.position)
	undo_redo.commit_action()
[/gdscript]
[csharp]
private UndoRedo _undoRedo;

public override void _Ready()
{
	_undoRedo = new UndoRedo();
}

public void DoSomething()
{
	// Put your code here.
}

public void UndoSomething()
{
	// Put here the code that reverts what's done by "DoSomething()".
}

private void OnMyButtonPressed()
{
	var node = GetNode<Node2D>("MyNode2D");
	_undoRedo.CreateAction("Move the node");
	_undoRedo.AddDoMethod(new Callable(this, MethodName.DoSomething));
	_undoRedo.AddUndoMethod(new Callable(this, MethodName.UndoSomething));
	_undoRedo.AddDoProperty(node, "position", new Vector2(100, 100));
	_undoRedo.AddUndoProperty(node, "position", node.Position);
	_undoRedo.CommitAction();
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/UndoRedo"
	"graphics.gd/variant/Vector2"
)

var undoRedo = UndoRedo.New()

func ExampleUndoRedo(node Node2D.Instance) {
	undoRedo.CreateAction("Move the node")
	undoRedo.AddDoMethod(func() {})
	undoRedo.AddUndoMethod(func() {})
	undoRedo.AddDoProperty(node.AsObject(), "position", Vector2.New(100.0, 100.0))
	undoRedo.AddUndoProperty(node.AsObject(), "position", node.Position())
	undoRedo.CommitAction()
}
