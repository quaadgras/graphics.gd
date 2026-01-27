/*
# Godot has been executed with the following command:
# godot --headless --verbose --scene my_scene.tscn --custom
OS.get_cmdline_args() # Returns ["--scene", "my_scene.tscn", "--custom"]
*/

package main

import (
	"graphics.gd/classdb/OS"
)

func OS_GetCmdlineArgs() {
	// Engine has been executed with the following command:
	// binary --headless --verbose --scene my_scene.tscn --custom
	OS.GetCmdlineArgs() // Returns ["--scene", "my_scene.tscn", "--custom"]
}
