/*
[gdscript]
var output = []
var exit_code = OS.execute("ls", ["-l", "/tmp"], output)
[/gdscript]
[csharp]
Godot.Collections.Array output = [];
int exitCode = OS.Execute("ls", ["-l", "/tmp"], output);
[/csharp]
*/

package main

import "graphics.gd/classdb/OS"

func OS_Execute() {
	output, exit_code := OS.Execute("ls", []string{"-l", "/tmp"}, false, false)
	_ = output
	_ = exit_code
}
