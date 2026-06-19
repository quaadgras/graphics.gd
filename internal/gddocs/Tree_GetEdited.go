/*
[gdscript]
func _ready():
	$Tree.item_edited.connect(on_Tree_item_edited)

func on_Tree_item_edited():
	print($Tree.get_edited()) # This item just got edited (e.g. checked).
[/gdscript]
[csharp]
public override void _Ready()
{
	GetNode<Tree>("Tree").ItemEdited += OnTreeItemEdited;
}

public void OnTreeItemEdited()
{
	GD.Print(GetNode<Tree>("Tree").GetEdited()); // This item just got edited (e.g. checked).
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Tree"
)

type treeGetEditedExample struct {
	Node.Extension[treeGetEditedExample]

	Tree Tree.Instance
}

func (n treeGetEditedExample) Ready() {
	n.Tree.OnItemEdited(n.onTreeItemEdited)
}

func (n treeGetEditedExample) onTreeItemEdited() {
	fmt.Println(n.Tree.GetEdited()) // This item just got edited (e.g. checked).
}
