/*
[gdscript]
func _can_drop_data(position, data):
    return typeof(data) == TYPE_DICTIONARY and data.has("color")

func _drop_data(position, data):
    var color = data["color"]
[/gdscript]
[csharp]
public override bool _CanDropData(Vector2 atPosition, Variant data)
{
    return data.VariantType == Variant.Type.Dictionary && data.AsGodotDictionary().ContainsKey("color");
}

public override void _DropData(Vector2 atPosition, Variant data)
{
    Color color = data.AsGodotDictionary()["color"].AsColor();
}
[/csharp]
*/

package main

import (
	"reflect"

	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
)

var color Color.RGBA

func Control_DropData() {
	CanDropData := func(position Vector2.XY, data any) bool {
		return reflect.TypeOf(data).Kind() == reflect.Map && data.(map[any]any)["color"] != nil
	}
	DropData := func(position Vector2.XY, data any) {
		color = data.(map[any]any)["color"].(Color.RGBA)
	}
	_ = CanDropData
	_ = DropData
}
