/*
func _popup_menu(paths):
	var popup_menu = PopupMenu.new()
	popup_menu.add_item("Blue")
	popup_menu.add_item("White")
	popup_menu.id_pressed.connect(_on_color_submenu_option)

	add_context_submenu_item("Set Node Color", popup_menu)
*/

package main

import "graphics.gd/classdb/PopupMenu"

func EditorContextMenuPlugin_AddContextSubmenuItem() {
	var popup_menu = PopupMenu.New()
	popup_menu.AddItem("Blue")
	popup_menu.AddItem("White")
	popup_menu.OnIdPressed(func(id int) {})
	editorContextMenuPlugin.AddContextSubmenuItem("Set Node Color", popup_menu)
}
