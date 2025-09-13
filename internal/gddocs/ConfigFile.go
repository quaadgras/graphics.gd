/*
[gdscript]
# Create new ConfigFile object.
var config = ConfigFile.new()

# Store some values.
config.set_value("Player1", "player_name", "Steve")
config.set_value("Player1", "best_score", 10)
config.set_value("Player2", "player_name", "V3geta")
config.set_value("Player2", "best_score", 9001)

# Save it to a file (overwrite if already exists).
config.save("user://scores.cfg")
[/gdscript]
[csharp]
// Create new ConfigFile object.
var config = new ConfigFile();

// Store some values.
config.SetValue("Player1", "player_name", "Steve");
config.SetValue("Player1", "best_score", 10);
config.SetValue("Player2", "player_name", "V3geta");
config.SetValue("Player2", "best_score", 9001);

// Save it to a file (overwrite if already exists).
config.Save("user://scores.cfg");
[/csharp]
*/

package main

import "graphics.gd/classdb/ConfigFile"

func ExampleConfigFileSave() {
	// Create new ConfigFile object.
	var config = ConfigFile.New()

	// Store some values.
	config.SetValue("Player1", "player_name", "Steve")
	config.SetValue("Player1", "best_score", 10)
	config.SetValue("Player2", "player_name", "V3geta")
	config.SetValue("Player2", "best_score", 9001)

	// Save it to a file (overwrite if already exists).
	config.Save("user://scores.cfg")
}
