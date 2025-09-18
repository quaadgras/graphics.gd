/*
[gdscript]
var file_to_remove = "user://slot1.save"
OS.move_to_trash(ProjectSettings.globalize_path(file_to_remove))
[/gdscript]
[csharp]
var fileToRemove = "user://slot1.save";
OS.MoveToTrash(ProjectSettings.GlobalizePath(fileToRemove));
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/ProjectSettings"
)

func OS_MoveToTrash() {
	var file_to_remove = "user://slot1.save"
	OS.MoveToTrash(ProjectSettings.GlobalizePath(file_to_remove))
}
