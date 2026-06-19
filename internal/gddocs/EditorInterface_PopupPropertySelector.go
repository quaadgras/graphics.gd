/*
func _ready():
	if Engine.is_editor_hint():
		EditorInterface.popup_property_selector(this, _on_property_selected, [TYPE_INT])

func _on_property_selected(property_path):
	if property_path.is_empty():
		print("property selection canceled")
	else:
		print("selected ", property_path)
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/Engine"
	"graphics.gd/variant/Object"
)

func ExamplePopupPropertySelector(self Object.Instance) {
	if Engine.IsEditorHint() {
		EditorInterface.PopupPropertySelector(self, onPropertySelected, []int32{2}, "") // 2 = TYPE_INT
	}
}

func onPropertySelected(propertyPath string) {
	if propertyPath == "" {
		fmt.Println("property selection canceled")
	} else {
		fmt.Println("selected ", propertyPath)
	}
}
