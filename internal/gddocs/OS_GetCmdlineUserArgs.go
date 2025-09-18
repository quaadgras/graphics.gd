/*
# Godot has been executed with the following command:
# godot --fullscreen -- --level=2 --hardcore

OS.get_cmdline_args()      # Returns ["--fullscreen", "--level=2", "--hardcore"]
OS.get_cmdline_user_args() # Returns ["--level=2", "--hardcore"]
*/

package main

import "graphics.gd/classdb/OS"

func OS_GetCmdlineUserArgs() {
	// The application has been executed with the following command:
	// myprogram --fullscreen -- --level=2 --hardcore
	OS.GetCmdlineArgs()     // Returns ["--fullscreen", "--level=2", "--hardcore"]
	OS.GetCmdlineUserArgs() // Returns ["--level=2", "--hardcore"]
}
