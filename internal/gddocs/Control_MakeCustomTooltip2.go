/*
[gdscript]
func _make_custom_tooltip(for_text):
	var tooltip = preload("res://some_tooltip_scene.tscn").instantiate()
	tooltip.get_node("Label").text = for_text
	return tooltip
[/gdscript]
[csharp]
public override Control _MakeCustomTooltip(string forText)
{
	Node tooltip = ResourceLoader.Load<PackedScene>("res://some_tooltip_scene.tscn").Instantiate();
	tooltip.GetNode<Label>("Label").Text = forText;
	return tooltip;
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/variant/Object"
)

type controlMakeCustomTooltip struct {
	Control.Extension[controlMakeCustomTooltip]

	TooltipScene PackedScene.Instance // preload("res://some_tooltip_scene.tscn")
}

func (n controlMakeCustomTooltip) MakeCustomTooltip(forText string) Object.Instance {
	var tooltip = n.TooltipScene.Instantiate()
	if label, ok := Object.As[Label.Instance](tooltip.GetNode("Label")); ok {
		label.SetText(forText)
	}
	return tooltip.AsObject()
}
