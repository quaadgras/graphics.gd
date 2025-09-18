/*
[gdscript]
func _ready():
	var menu = get_menu()
	# Remove "Select All" item.
	menu.remove_item(MENU_SELECT_ALL)
	# Add custom items.
	menu.add_separator()
	menu.add_item("Duplicate Text", MENU_MAX + 1)
	# Connect callback.
	menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):
	if id == MENU_MAX + 1:
		add_text("\n" + get_parsed_text())
[/gdscript]
[csharp]
public override void _Ready()
{
	var menu = GetMenu();
	// Remove "Select All" item.
	menu.RemoveItem(RichTextLabel.MenuItems.SelectAll);
	// Add custom items.
	menu.AddSeparator();
	menu.AddItem("Duplicate Text", RichTextLabel.MenuItems.Max + 1);
	// Add event handler.
	menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
	if (id == TextEdit.MenuItems.Max + 1)
	{
		AddText("\n" + GetParsedText());
	}
}
[/csharp]
*/

package main

import "graphics.gd/classdb/RichTextLabel"

var richTextLabel RichTextLabel.Instance

func RichTextLabel_GetMenu() {
	var menu = richTextLabel.GetMenu()
	// Remove "Select All" item.
	menu.RemoveItem(int(RichTextLabel.MenuSelectAll))
	// Add custom items.
	menu.AddSeparator()
	menu.MoreArgs().AddItem("Duplicate Text", int(RichTextLabel.MenuMax)+1, 0)
	// Add event handler.
	menu.OnIdPressed(func(id int) {
		if id == int(RichTextLabel.MenuMax)+1 {
			richTextLabel.AddText("\n" + richTextLabel.GetParsedText())
		}
	})
}
