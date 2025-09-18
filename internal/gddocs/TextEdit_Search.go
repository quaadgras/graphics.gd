/*
[gdscript]
var result = search("print", SEARCH_WHOLE_WORDS, 0, 0)
if result.x != -1:
	# Result found.
	var line_number = result.y
	var column_number = result.x
[/gdscript]
[csharp]
Vector2I result = Search("print", (uint)TextEdit.SearchFlags.WholeWords, 0, 0);
if (result.X != -1)
{
	// Result found.
	int lineNumber = result.Y;
	int columnNumber = result.X;
}
[/csharp]
*/

package main

import "graphics.gd/classdb/TextEdit"

var textEdit TextEdit.Instance

func TextEdit_Search() {
	line, column := textEdit.Search("print", TextEdit.SearchWholeWords, 0, 0)
	_ = line
	_ = column
}
