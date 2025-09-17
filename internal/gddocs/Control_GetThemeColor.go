/*
[gdscript]
func _ready():
	# Get the font color defined for the current Control's class, if it exists.
	modulate = get_theme_color("font_color")
	# Get the font color defined for the Button class.
	modulate = get_theme_color("font_color", "Button")
[/gdscript]
[csharp]
public override void _Ready()
{
	// Get the font color defined for the current Control's class, if it exists.
	Modulate = GetThemeColor("font_color");
	// Get the font color defined for the Button class.
	Modulate = GetThemeColor("font_color", "Button");
}
[/csharp]
*/

package main

func Control_GetThemeColor() {
	// Get the font color defined for the current Control's class, if it exists.
	control.AsCanvasItem().SetModulate(control.GetThemeColor("font_color"))
	// Get the font color defined for the Button class.
	control.MoreArgs().GetThemeColor("font_color", "Button")
}
