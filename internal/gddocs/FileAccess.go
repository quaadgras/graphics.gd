/*
[gdscript]
func save_to_file(content):
	var file = FileAccess.open("user://save_game.dat", FileAccess.WRITE)
	file.store_string(content)

func load_from_file():
	var file = FileAccess.open("user://save_game.dat", FileAccess.READ)
	var content = file.get_as_text()
	return content
[/gdscript]
[csharp]
public void SaveToFile(string content)
{
	using var file = FileAccess.Open("user://save_game.dat", FileAccess.ModeFlags.Write);
	file.StoreString(content);
}

public string LoadFromFile()
{
	using var file = FileAccess.Open("user://save_game.dat", FileAccess.ModeFlags.Read);
	string content = file.GetAsText();
	return content;
}
[/csharp]
*/

package main

import "graphics.gd/classdb/FileAccess"

func SaveToFile(content string) {
	var file = FileAccess.Open("user://save_game.dat", FileAccess.Write)
	file.StoreString(content)
}

func LoadFromFile() string {
	var file = FileAccess.Open("user://save_game.dat", FileAccess.Read)
	var content = file.GetAsText()
	return content
}
