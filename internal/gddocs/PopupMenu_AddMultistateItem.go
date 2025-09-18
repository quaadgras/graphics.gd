/*
func _ready():
	add_multistate_item("Item", 3, 0)

	index_pressed.connect(func(index: int):
			toggle_item_multistate(index)
			match get_item_multistate(index):
				0:
					print("First state")
				1:
					print("Second state")
				2:
					print("Third state")
		)
*/

package main

import "fmt"

func PopupMenu_AddMultistateItem() {
	popup_menu.AddMultistateItem("Item", 3)
	popup_menu.OnIndexPressed(func(index int) {
		popup_menu.ToggleItemMultistate(index)
		switch popup_menu.GetItemMultistate(index) {
		case 0:
			fmt.Println("First state")
		case 1:
			fmt.Println("Second state")
		case 2:
			fmt.Println("Third state")
		}
	})
}
