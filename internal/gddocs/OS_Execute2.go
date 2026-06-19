/*
[gdscript]
var output = []
OS.execute("CMD.exe", ["/C", "cd %TEMP% && dir"], output)
[/gdscript]
[csharp]
Godot.Collections.Array output = [];
OS.Execute("CMD.exe", ["/C", "cd %TEMP% && dir"], output);
[/csharp]
*/

package main

import "graphics.gd/classdb/OS"

func ExampleOSExecute() {
	output, _ := OS.Execute("CMD.exe", []string{"/C", "cd %TEMP% && dir"}, false, false)
	_ = output
}
