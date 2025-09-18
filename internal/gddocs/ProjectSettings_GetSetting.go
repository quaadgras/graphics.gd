/*
[gdscript]
print(ProjectSettings.get_setting("application/config/name"))
print(ProjectSettings.get_setting("application/config/custom_description", "No description specified."))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSetting("application/config/name"));
GD.Print(ProjectSettings.GetSetting("application/config/custom_description", "No description specified."));
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/ProjectSettings"
)

func ProjectSettings_GetSetting() {
	fmt.Println(ProjectSettings.GetSetting("application/config/name", nil))
	fmt.Println(ProjectSettings.GetSetting("application/config/custom_description", "No description specified."))
}
