/*
[gdscript]
# The snippet below assumes the child node "MyButton" has a StyleBoxFlat assigned.
# Resources are shared across instances, so we need to duplicate it
# to avoid modifying the appearance of all other buttons.
var new_stylebox_normal = $MyButton.get_theme_stylebox("normal").duplicate()
new_stylebox_normal.border_width_top = 3
new_stylebox_normal.border_color = Color(0, 1, 0.5)
$MyButton.add_theme_stylebox_override("normal", new_stylebox_normal)
# Remove the stylebox override.
$MyButton.remove_theme_stylebox_override("normal")
[/gdscript]
[csharp]
// The snippet below assumes the child node "MyButton" has a StyleBoxFlat assigned.
// Resources are shared across instances, so we need to duplicate it
// to avoid modifying the appearance of all other buttons.
StyleBoxFlat newStyleboxNormal = GetNode<Button>("MyButton").GetThemeStylebox("normal").Duplicate() as StyleBoxFlat;
newStyleboxNormal.BorderWidthTop = 3;
newStyleboxNormal.BorderColor = new Color(0, 1, 0.5f);
GetNode<Button>("MyButton").AddThemeStyleboxOverride("normal", newStyleboxNormal);
// Remove the stylebox override.
GetNode<Button>("MyButton").RemoveThemeStyleboxOverride("normal");
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Resource"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Object"
)

var MyButton Button.Instance

func Control_AddThemeStyleboxOverride() {
	// The snippet below assumes the child node "MyButton" has a StyleBoxFlat assigned.
	// Resources are shared across instances, so we need to duplicate it
	// to avoid modifying the appearance of all other buttons.
	var newStyleboxNormal = Resource.Duplicate(MyButton.AsControl().GetThemeStylebox("normal"))
	Object.Set(newStyleboxNormal, "border_width_top", 3)
	Object.Set(newStyleboxNormal, "border_color", Color.RGBA{0, 1, 0.5, 1})
	MyButton.AsControl().AddThemeStyleboxOverride("normal", newStyleboxNormal)
	// Remove the stylebox override.
	MyButton.AsControl().RemoveThemeStyleboxOverride("normal")
}
