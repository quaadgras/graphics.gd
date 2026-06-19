/*
func _get_unsaved_status(for_scene):
	if not for_scene.is_empty():
		return ""
*/

package main

import "graphics.gd/classdb/EditorPlugin"

type editorPluginUnsavedStatus struct {
	EditorPlugin.Extension[editorPluginUnsavedStatus]
}

func (n editorPluginUnsavedStatus) GetUnsavedStatus(forScene string) string {
	if forScene != "" {
		return ""
	}
	return ""
}
