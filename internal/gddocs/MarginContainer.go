/*
[gdscript]
# This code sample assumes the current script is extending MarginContainer.
var margin_value = 100
add_theme_constant_override("margin_top", margin_value)
add_theme_constant_override("margin_left", margin_value)
add_theme_constant_override("margin_bottom", margin_value)
add_theme_constant_override("margin_right", margin_value)
[/gdscript]
[csharp]
// This code sample assumes the current script is extending MarginContainer.
int marginValue = 100;
AddThemeConstantOverride("margin_top", marginValue);
AddThemeConstantOverride("margin_left", marginValue);
AddThemeConstantOverride("margin_bottom", marginValue);
AddThemeConstantOverride("margin_right", marginValue);
[/csharp]
*/

package main

import "graphics.gd/classdb/MarginContainer"

func ExampleMarginContainer(c MarginContainer.Instance) {
	var marginValue = 100
	c.AsControl().AddThemeConstantOverride("margin_top", marginValue)
	c.AsControl().AddThemeConstantOverride("margin_left", marginValue)
	c.AsControl().AddThemeConstantOverride("margin_bottom", marginValue)
	c.AsControl().AddThemeConstantOverride("margin_right", marginValue)
}
