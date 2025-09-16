/*
var menu

func _menu_callback(item_id):
	if item_id == "ITEM_CUT":
		cut()
	elif item_id == "ITEM_COPY":
		copy()
	elif item_id == "ITEM_PASTE":
		paste()

func _enter_tree():
	# Create new menu and add items:
	menu = NativeMenu.create_menu()
	NativeMenu.add_item(menu, "Cut", _menu_callback, Callable(), "ITEM_CUT")
	NativeMenu.add_item(menu, "Copy", _menu_callback, Callable(), "ITEM_COPY")
	NativeMenu.add_separator(menu)
	NativeMenu.add_item(menu, "Paste", _menu_callback, Callable(), "ITEM_PASTE")

func _on_button_pressed():
	# Show popup menu at mouse position:
	NativeMenu.popup(menu, DisplayServer.mouse_get_position())

func _exit_tree():
	# Remove menu when it's no longer needed:
	NativeMenu.free_menu(menu)
*/

package main

import (
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/NativeMenu"
	"graphics.gd/variant/RID"
)

var menu RID.NativeMenu

func MenuCallback(item_id any) {
	switch item_id {
	case "ITEM_CUT":
	case "ITEM_COPY":
	case "ITEM_PASTE":
	}
}

func EnterTree() {
	menu = NativeMenu.CreateMenu()
	NativeMenu.AddItem(menu, "Cut", MenuCallback, nil, "ITEM_CUT", 0)
	NativeMenu.AddItem(menu, "Copy", MenuCallback, nil, "ITEM_COPY", 0)
	NativeMenu.AddSeparator(menu)
	NativeMenu.AddItem(menu, "Paste", MenuCallback, nil, "ITEM_PASTE", 0)
}

func OnButtonPressed() {
	NativeMenu.Popup(menu, DisplayServer.MouseGetPosition())
}

func ExitTree() {
	NativeMenu.FreeMenu(menu)
}
