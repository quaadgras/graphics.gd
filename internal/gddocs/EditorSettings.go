/*
[gdscript]
var settings = EditorInterface.get_editor_settings()
# `settings.set("some/property", 10)` also works as this class overrides `_set()` internally.
settings.set_setting("some/property", 10)
# `settings.get("some/property")` also works as this class overrides `_get()` internally.
settings.get_setting("some/property")
var list_of_settings = settings.get_property_list()
[/gdscript]
[csharp]
EditorSettings settings = EditorInterface.Singleton.GetEditorSettings();
// `settings.set("some/property", value)` also works as this class overrides `_set()` internally.
settings.SetSetting("some/property", Value);
// `settings.get("some/property", value)` also works as this class overrides `_get()` internally.
settings.GetSetting("some/property");
Godot.Collections.Array<Godot.Collections.Dictionary> listOfSettings = settings.GetPropertyList();
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/variant/Object"
)

func ExampleEditorSettings() {
	var settings = EditorInterface.GetEditorSettings()
	settings.SetSetting("some/property", 10)
	settings.GetSetting("some/property")
	var list_of_settings = Object.GetPropertyList(settings)
	_ = list_of_settings
}
