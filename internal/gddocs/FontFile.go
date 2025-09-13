/*
[gdscript]
var f = load("res://BarlowCondensed-Bold.ttf")
$Label.add_theme_font_override("font", f)
$Label.add_theme_font_size_override("font_size", 64)
[/gdscript]
[csharp]
var f = ResourceLoader.Load<FontFile>("res://BarlowCondensed-Bold.ttf");
GetNode("Label").AddThemeFontOverride("font", f);
GetNode("Label").AddThemeFontSizeOverride("font_size", 64);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/FontFile"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Resource"
)

func ExampleFontFile(label Label.Instance) {
	var f = Resource.Load[FontFile.Instance]("res://BarlowCondensed-Bold.ttf")
	label.AsControl().AddThemeFontOverride("font", f.AsFont())
	label.AsControl().AddThemeFontSizeOverride("font_size", 64)
}
