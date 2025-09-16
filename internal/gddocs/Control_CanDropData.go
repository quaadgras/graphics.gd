/*
[gdscript]
func _can_drop_data(position, data):
	# Check position if it is relevant to you
	# Otherwise, just check data
	return typeof(data) == TYPE_DICTIONARY and data.has("expected")
[/gdscript]
[csharp]
public override bool _CanDropData(Vector2 atPosition, Variant data)
{
	// Check position if it is relevant to you
	// Otherwise, just check data
	return data.VariantType == Variant.Type.Dictionary && data.AsGodotDictionary().ContainsKey("expected");
}
[/csharp]
*/

package main

import (
	"reflect"

	"graphics.gd/variant/Vector2"
)

func Control_CanDropData() {
	CanDropData := func(position Vector2.XY, data any) bool {
		return reflect.TypeOf(data).Kind() == reflect.Map && data.(map[any]any)["expected"] != nil
	}
	_ = CanDropData
}
