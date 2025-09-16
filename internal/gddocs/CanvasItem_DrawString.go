/*
[gdscript]
# If using this method in a script that redraws constantly, move the
# `default_font` declaration to a member variable assigned in `_ready()`
# so the Control is only created once.
var default_font = ThemeDB.fallback_font
var default_font_size = ThemeDB.fallback_font_size
draw_string(default_font, Vector2(64, 64), "Hello world", HORIZONTAL_ALIGNMENT_LEFT, -1, default_font_size)
[/gdscript]
[csharp]
// If using this method in a script that redraws constantly, move the
// `default_font` declaration to a member variable assigned in `_Ready()`
// so the Control is only created once.
Font defaultFont = ThemeDB.FallbackFont;
int defaultFontSize = ThemeDB.FallbackFontSize;
DrawString(defaultFont, new Vector2(64, 64), "Hello world", HORIZONTAL_ALIGNMENT_LEFT, -1, defaultFontSize);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/CanvasItem"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/TextServer"
	"graphics.gd/classdb/ThemeDB"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

var canvas_item CanvasItem.Instance

func CanvasItem_DrawString() {
	// If using this method in a script that redraws constantly, move the
	// `default_font` declaration to a member variable assigned in `Ready()`
	// so the Control is only created once.
	var defaultFont = ThemeDB.FallbackFont()
	var defaultFontSize = ThemeDB.FallbackFontSize()
	CanvasItem.Expanded(canvas_item).DrawString(defaultFont, Vector2.New(64, 64), "Hello world", GUI.HorizontalAlignmentLeft, -1, defaultFontSize,
		Color.W3C.White, TextServer.JustificationKashida|TextServer.JustificationWordBound, 0, 0, 0)
}
