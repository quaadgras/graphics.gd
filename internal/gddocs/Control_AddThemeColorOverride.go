/*
[gdscript]
# Given the child Label node "MyLabel", override its font color with a custom value.
$MyLabel.add_theme_color_override("font_color", Color(1, 0.5, 0))
# Reset the font color of the child label.
$MyLabel.remove_theme_color_override("font_color")
# Alternatively it can be overridden with the default value from the Label type.
$MyLabel.add_theme_color_override("font_color", get_theme_color("font_color", "Label"))
[/gdscript]
[csharp]
// Given the child Label node "MyLabel", override its font color with a custom value.
GetNode<Label>("MyLabel").AddThemeColorOverride("font_color", new Color(1, 0.5f, 0));
// Reset the font color of the child label.
GetNode<Label>("MyLabel").RemoveThemeColorOverride("font_color");
// Alternatively it can be overridden with the default value from the Label type.
GetNode<Label>("MyLabel").AddThemeColorOverride("font_color", GetThemeColor("font_color", "Label"));
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/variant/Color"
)

var MyLabel Label.Instance

func Control_AddThemeColorOverride() {
	// Given the child Label node "MyLabel", override its font color with a custom value.
	MyLabel.AsControl().AddThemeColorOverride("font_color", Color.RGBA{1, 0.5, 0, 1})
	// Reset the font color of the child label.
	MyLabel.AsControl().RemoveThemeColorOverride("font_color")
	// Alternatively it can be overridden with the default value from the Label type.
	MyLabel.AsControl().AddThemeColorOverride("font_color", Control.Expanded(MyLabel.AsControl()).GetThemeColor("font_color", "Label"))
}
