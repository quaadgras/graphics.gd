/*
[gdscript]
print(ProjectSettings.get_setting_with_override("application/config/name"))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSettingWithOverride("application/config/name"));
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/ProjectSettings"
)

func ProjectSettings_GetSettingWithOverride() {
	fmt.Println(ProjectSettings.GetSettingWithOverride("application/config/name"))
}
