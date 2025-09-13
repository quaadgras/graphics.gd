/*
[gdscript]
var spin_box = SpinBox.new()
add_child(spin_box)
var line_edit = spin_box.get_line_edit()
line_edit.context_menu_enabled = false
spin_box.horizontal_alignment = LineEdit.HORIZONTAL_ALIGNMENT_RIGHT
[/gdscript]
[csharp]
var spinBox = new SpinBox();
AddChild(spinBox);
var lineEdit = spinBox.GetLineEdit();
lineEdit.ContextMenuEnabled = false;
spinBox.AlignHorizontal = LineEdit.HorizontalAlignEnum.Right;
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SpinBox"
)

func ExampleSpinBox(parent Node.Instance) {
	var spinBox = SpinBox.New()
	parent.AddChild(spinBox.AsNode())
	var lineEdit = spinBox.GetLineEdit()
	lineEdit.SetContextMenuEnabled(false)
	SpinBox.Advanced(spinBox).SetHorizontalAlignment(GUI.HorizontalAlignmentRight)
}
