/*
[gdscript]
func _ready():
	var tree = Tree.new()
	var root = tree.create_item()
	tree.hide_root = true
	var child1 = tree.create_item(root)
	var child2 = tree.create_item(root)
	var subchild1 = tree.create_item(child1)
	subchild1.set_text(0, "Subchild1")
[/gdscript]
[csharp]
public override void _Ready()
{
	var tree = new Tree();
	TreeItem root = tree.CreateItem();
	tree.HideRoot = true;
	TreeItem child1 = tree.CreateItem(root);
	TreeItem child2 = tree.CreateItem(root);
	TreeItem subchild1 = tree.CreateItem(child1);
	subchild1.SetText(0, "Subchild1");
}
[/csharp]
*/

package main

import "graphics.gd/classdb/Tree"

func ExampleTree() {
	var tree = Tree.New()
	var root = tree.CreateItem()
	tree.SetHideRoot(true)
	var child1 = Tree.Expanded(tree).CreateItem(root, -1)
	var child2 = Tree.Expanded(tree).CreateItem(root, -1)
	var subchild1 = Tree.Expanded(tree).CreateItem(child1, -1)
	subchild1.SetText(0, "Subchild1")
	_ = child2
}
