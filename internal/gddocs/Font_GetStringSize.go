/*
[gdscript]
var string_size = $Label.get_theme_font("font").get_string_size($Label.text, HORIZONTAL_ALIGNMENT_LEFT, -1, $Label.get_theme_font_size("font_size"))
[/gdscript]
[csharp]
Label label = GetNode<Label>("Label");
Vector2 stringSize = label.GetThemeFont("font").GetStringSize(label.Text, HorizontalAlignment.Left, -1, label.GetThemeFontSize("font_size"));
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Font"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/TextServer"
)

var label Label.Instance

func Font_GetStringSize() {
	var string_size = Font.Expanded(label.AsControl().GetThemeFont("font")).GetStringSize(label.Text(), GUI.HorizontalAlignmentLeft, -1, label.AsControl().GetThemeFontSize("font_size"),
		TextServer.JustificationKashida|TextServer.JustificationWordBound, 0, 0)
	_ = string_size
}
