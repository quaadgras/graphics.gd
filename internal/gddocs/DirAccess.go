/*
# Standard
var dir = DirAccess.open("user://levels")
dir.make_dir("world1")
# Static
DirAccess.make_dir_absolute("user://levels/world1")
*/

package main

import "graphics.gd/classdb/DirAccess"

func ExampleDirectoryMake() {
	var dir = DirAccess.Open("user://levels")
	dir.MakeDir("world1")
	DirAccess.MakeDirAbsolute("user://levels/world1")
}
