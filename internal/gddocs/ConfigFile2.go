/*
[gdscript]
var score_data = {}
var config = ConfigFile.new()

# Load data from a file.
var err = config.load("user://scores.cfg")

# If the file didn't load, ignore it.
if err != OK:
	return

# Iterate over all sections.
for player in config.get_sections():
	# Fetch the data for each section.
	var player_name = config.get_value(player, "player_name")
	var player_score = config.get_value(player, "best_score")
	score_data[player_name] = player_score
[/gdscript]
[csharp]
var score_data = new Godot.Collections.Dictionary();
var config = new ConfigFile();

// Load data from a file.
Error err = config.Load("user://scores.cfg");

// If the file didn't load, ignore it.
if (err != Error.Ok)
{
	return;
}

// Iterate over all sections.
foreach (String player in config.GetSections())
{
	// Fetch the data for each section.
	var player_name = (String)config.GetValue(player, "player_name");
	var player_score = (int)config.GetValue(player, "best_score");
	score_data[player_name] = player_score;
}
[/csharp]
*/

package main

import "graphics.gd/classdb/ConfigFile"

func ExampleConfigFileLoad() error {
	var score_data = make(map[string]any)
	var config = ConfigFile.New()

	// Load data from a file.
	var err = config.Load("user://scores.cfg")
	if err != nil {
		return err
	}
	// Iterate over all sections.
	// for player in config.get_sections():
	for _, player := range config.GetSections() {
		// Fetch the data for each section.
		var player_name = config.GetValue(player, "player_name")
		var player_score = config.GetValue(player, "best_score")
		score_data[player_name.(string)] = player_score
	}
	return nil
}
