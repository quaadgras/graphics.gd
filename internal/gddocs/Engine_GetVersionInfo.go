/*
[gdscript]
if Engine.get_version_info().hex >= 0x040100:
	pass # Do things specific to version 4.1 or later.
else:
	pass # Do things specific to versions before 4.1.
[/gdscript]
[csharp]
if ((int)Engine.GetVersionInfo()["hex"] >= 0x040100)
{
	// Do things specific to version 4.1 or later.
}
else
{
	// Do things specific to versions before 4.1.
}
[/csharp]
*/

package main

import "graphics.gd/classdb/Engine"

func Engine_GetVersionInfo() {
	if Engine.GetVersionInfo().Hex >= 0x040100 {
		// Do things specific to version 4.1 or later.
	} else {
		// Do things specific to versions before 4.1.
	}
}
