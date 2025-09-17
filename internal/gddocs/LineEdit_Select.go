/*
[gdscript]
text = "Welcome"
select() # Will select "Welcome".
select(4) # Will select "ome".
select(2, 5) # Will select "lco".
[/gdscript]
[csharp]
Text = "Welcome";
Select(); // Will select "Welcome".
Select(4); // Will select "ome".
Select(2, 5); // Will select "lco".
[/csharp]
*/

package main

import "graphics.gd/classdb/LineEdit"

var line_edit LineEdit.Instance

func LineEdit_Select() {
	line_edit.SetText("Welcome")
	line_edit.Select()                 // Will select "Welcome".
	line_edit.MoreArgs().Select(4, -1) // Will select "ome".
	line_edit.MoreArgs().Select(2, 5)  // Will select "lco".
}
