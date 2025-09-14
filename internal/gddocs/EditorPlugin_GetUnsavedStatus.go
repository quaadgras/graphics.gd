/*
func _get_unsaved_status(for_scene):
    if not unsaved:
        return ""

    if for_scene.is_empty():
        return "Save changes in MyCustomPlugin before closing?"
    else:
        return "Scene %s has changes from MyCustomPlugin. Save before closing?" % for_scene.get_file()

func _save_external_data():
    unsaved = false
*/

package main

import "path/filepath"

func EditorPlugin_GetUnsavedStatus() {
	var unsaved bool
	GetUnsavedStatus := func(for_scene string) string {
		if !unsaved {
			return ""
		}
		if for_scene == "" {
			return "Save changes in MyCustomPlugin before closing?"
		}
		return "Scene " + filepath.Base(for_scene) + " has changes from MyCustomPlugin. Save before closing?"
	}
	SaveExternalData := func() {
		unsaved = false
	}
	_ = GetUnsavedStatus
	_ = SaveExternalData
}
