/*
func _ready():
	if Engine.is_editor_hint():
		EditorInterface.popup_node_selector(_on_node_selected, ["Button"])

func _on_node_selected(node_path):
	if node_path.is_empty():
		print("node selection canceled")
	else:
		print("selected ", node_path)
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Node"
)

func ExamplePopupNodeSelector() {
	if Engine.IsEditorHint() {
		EditorInterface.PopupNodeSelector(onNodeSelected, []string{"Button"}, Node.Nil)
	}
}

func onNodeSelected(nodePath string) {
	if nodePath == "" {
		fmt.Println("node selection canceled")
	} else {
		fmt.Println("selected ", nodePath)
	}
}
