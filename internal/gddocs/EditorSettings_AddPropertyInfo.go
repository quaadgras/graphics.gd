/*
[gdscript]
var settings = EditorInterface.get_editor_settings()
settings.set("category/property_name", 0)

var property_info = {
	"name": "category/property_name",
	"type": TYPE_INT,
	"hint": PROPERTY_HINT_ENUM,
	"hint_string": "one,two,three"
}

settings.add_property_info(property_info)
[/gdscript]
[csharp]
var settings = GetEditorInterface().GetEditorSettings();
settings.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary
{
	{ "name", "category/propertyName" },
	{ "type", Variant.Type.Int },
	{ "hint", PropertyHint.Enum },
	{ "hint_string", "one,two,three" },
};

settings.AddPropertyInfo(propertyInfo);
[/csharp]
*/

package main

import (
	"reflect"

	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/variant/Object"
)

func ExampleEditorSettingsAddPropertyInfo() {
	var settings = EditorInterface.GetEditorSettings()
	settings.SetSetting("category/property_name", 0)

	var propertyInfo = Object.PropertyInfo{
		Name:       "category/property_name",
		Type:       reflect.TypeFor[int](),
		Hint:       2, // PROPERTY_HINT_ENUM
		HintString: "one,two,three",
	}

	settings.AddPropertyInfo(propertyInfo)
}
