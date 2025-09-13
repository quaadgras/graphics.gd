/*
[gdscript]
var ts = TextServerManager.get_primary_interface()
[/gdscript]
[csharp]
var ts = TextServerManager.GetPrimaryInterface();
[/csharp]
*/

package main

import "graphics.gd/classdb/TextServerManager"

func ExampleTextServer() {
	var ts = TextServerManager.GetPrimaryInterface()
	_ = ts
}
