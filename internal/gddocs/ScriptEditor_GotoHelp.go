/*
# Shows help for the Node class.
class_name:Node
# Shows help for the global min function.
# Global objects are accessible in the `@GlobalScope` namespace, shown here.
class_method:@GlobalScope:min
# Shows help for get_viewport in the Node class.
class_method:Node:get_viewport
# Shows help for the Input constant MOUSE_BUTTON_MIDDLE.
class_constant:Input:MOUSE_BUTTON_MIDDLE
# Shows help for the BaseButton signal pressed.
class_signal:BaseButton:pressed
# Shows help for the CanvasItem property visible.
class_property:CanvasItem:visible
# Shows help for the GDScript annotation export.
# Annotations should be prefixed with the `@` symbol in the descriptor, as shown here.
class_annotation:@GDScript:@export
# Shows help for the GraphNode theme item named panel_selected.
class_theme_item:GraphNode:panel_selected
*/

package main

import "graphics.gd/classdb/ScriptEditor"

var scriptEditor ScriptEditor.Instance

func ScriptEditor_GotoHelp() {
	var topics = []string{
		// Shows help for the Node class.
		"class_name:Node",
		// Shows help for the global min function.
		// Global objects are accessible in the GlobalScope namespace, shown here.
		"class_method:@GlobalScope:min",
		// Shows help for get_viewport in the Node class.
		"class_method:Node:get_viewport",
		// Shows help for the Input constant MOUSE_BUTTON_MIDDLE.
		"class_constant:Input:MOUSE_BUTTON_MIDDLE",
		// Shows help for the BaseButton signal pressed.
		"class_signal:BaseButton:pressed",
		// Shows help for the CanvasItem property visible.
		"class_property:CanvasItem:visible",
		// Shows help for the GDScript annotation export.
		// Annotations should be prefixed with the '@' symbol in the descriptor, as shown here.
		"class_annotation:@GDScript:@export",
		// Shows help for the GraphNode theme item named panel_selected.
		"class_theme_item:GraphNode:panel_selected",
	}
	for _, topic := range topics {
		scriptEditor.GotoHelp(topic)
	}
	_ = topics
}
