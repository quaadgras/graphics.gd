/*
func _popup_menu(paths):
    add_context_menu_item("File Custom options", handle, ICON)
*/

package main

func EditorContextMenuPlugin_AddContextMenuItem() {
	editorContextMenuPlugin.AddContextMenuItem("File Custom options", func(array []any) {})
}
