/*
[gdscript]
var thread = Thread.new()

func _ready():
	thread.start(
		func():
			var driver_info = OS.get_video_adapter_driver_info()
			if not driver_info.is_empty():
				print("Driver: %s %s" % [driver_info[0], driver_info[1]])
			else:
				print("Driver: (unknown)")
	)

func _exit_tree():
	thread.wait_to_finish()
[/gdscript]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/OS"
)

func OS_GetVideoAdapterDriverInfo() {
	go func() {
		var driverInfo = OS.GetVideoAdapterDriverInfo()
		if len(driverInfo) != 0 {
			fmt.Printf("Driver: %s %s\n", driverInfo[0], driverInfo[1])
		} else {
			fmt.Println("Driver: (unknown)")
		}
	}()
}
