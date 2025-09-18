/*
[gdscript]
ProjectSettings.set_setting("application/config/name", "Example")
[/gdscript]
[csharp]
ProjectSettings.SetSetting("application/config/name", "Example");
[/csharp]
*/

package main

import "graphics.gd/classdb/ProjectSettings"

func ProjectSettings_SetSetting() {
	ProjectSettings.SetSetting("application/config/name", "Example")
}
