/*
[gdscript]
func _ready():
    var menu = get_menu()
    # Remove all items after "Redo".
    menu.item_count = menu.get_item_index(MENU_REDO) + 1
    # Add custom items.
    menu.add_separator()
    menu.add_item("Insert Date", MENU_MAX + 1)
    # Connect callback.
    menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):
    if id == MENU_MAX + 1:
        insert_text_at_caret(Time.get_date_string_from_system())
[/gdscript]
[csharp]
public override void _Ready()
{
    var menu = GetMenu();
    // Remove all items after "Redo".
    menu.ItemCount = menu.GetItemIndex(LineEdit.MenuItems.Redo) + 1;
    // Add custom items.
    menu.AddSeparator();
    menu.AddItem("Insert Date", LineEdit.MenuItems.Max + 1);
    // Add event handler.
    menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
    if (id == LineEdit.MenuItems.Max + 1)
    {
        InsertTextAtCaret(Time.GetDateStringFromSystem());
    }
}
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/PopupMenu"
	"graphics.gd/classdb/Time"
)

func LineEdit_GetMenu() {
	Ready := func() {
		var menu = line_edit.GetMenu()
		// Remove all items after "Redo".
		menu.SetItemCount(menu.GetItemIndex(int(LineEdit.MenuRedo)) + 1)
		// Add custom items.
		menu.AddSeparator()
		PopupMenu.Expanded(menu).AddItem("Insert Date", int(LineEdit.MenuMax)+1, 0)
		// Connect callback.
		menu.OnIdPressed(func(id int) {
			if id == int(LineEdit.MenuMax)+1 {
				line_edit.InsertTextAtCaret(Time.GetDateStringFromSystem(false))
			}
		})
	}
	_ = Ready
}
