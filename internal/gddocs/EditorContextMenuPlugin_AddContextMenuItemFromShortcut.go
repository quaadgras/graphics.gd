/*
func _init():
	add_menu_shortcut(SHORTCUT, handle)

func _popup_menu(paths):
	add_context_menu_item_from_shortcut("File Custom options", SHORTCUT, ICON)
*/

package main

import (
	"graphics.gd/classdb/Texture2D"
)

var icon Texture2D.Instance

func EditorContextMenuPlugin_AddContextMenuItemFromShortcut() {
	editorContextMenuPlugin.AddMenuShortcut(shortcut, func(array []any) {})
	editorContextMenuPlugin.MoreArgs().AddContextMenuItemFromShortcut("File Custom options", shortcut, icon)
}
