/*
[gdscript]
ProjectSettings.set("category/property_name", 0)

var property_info = {
	"name": "category/property_name",
	"type": TYPE_INT,
	"hint": PROPERTY_HINT_ENUM,
	"hint_string": "one,two,three"
}

ProjectSettings.add_property_info(property_info)
[/gdscript]
[csharp]
ProjectSettings.Singleton.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary
{
	{ "name", "category/propertyName" },
	{ "type", (int)Variant.Type.Int },
	{ "hint", (int)PropertyHint.Enum },
	{ "hint_string", "one,two,three" },
};

ProjectSettings.AddPropertyInfo(propertyInfo);
[/csharp]
*/

package main

import (
	"reflect"

	"graphics.gd/classdb"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/variant/Object"
)

func ProjectSettings_AddPropertyInfo() {
	ProjectSettings.SetSetting("category/property_name", 0)
	ProjectSettings.AddPropertyInfo(Object.PropertyInfo{
		Name:       "category/propertyName",
		Type:       reflect.TypeFor[int](),
		Hint:       int(classdb.PropertyHintEnum),
		HintString: "one,two,three",
	})
}
