/*
var scene_root = EditorInterface.get_edited_scene_root()
var undo_redo = EditorInterface.get_editor_undo_redo()
undo_redo.clear_history(undo_redo.get_object_history_id(scene_root))
*/

package main

import (
	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/EditorUndoRedoManager"
)

func EditorUndoRedoManager_ClearHistory() {
	var scene_root = EditorInterface.GetEditedSceneRoot()
	var undo_redo = EditorInterface.GetEditorUndoRedo()
	EditorUndoRedoManager.Expanded(undo_redo).ClearHistory(undo_redo.GetObjectHistoryId(scene_root.AsObject()), false)
}
