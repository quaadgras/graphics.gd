/*
[gdscript]
func _make_custom_tooltip(for_text):
	var label = Label.new()
	label.text = for_text
	return label
[/gdscript]
[csharp]
public override Control _MakeCustomTooltip(string forText)
{
	var label = new Label();
	label.Text = forText;
	return label;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
)

func Control_MakeCustomTooltip() {
	MakeCustomTooltip := func(forText string) Control.Instance {
		var label = Label.New()
		label.SetText(forText)
		return label.AsControl()
	}
	_ = MakeCustomTooltip
}
