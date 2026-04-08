package Startup

import (
	gdunsafe "graphics.gd"
	gd "graphics.gd/internal"
)

func init() {
	gd.LinkStartup = func() {
		sname = gdunsafe.UTF8.Intern("GodotInstance")
		otype = gdunsafe.Class(sname).Tag()
		gd.LinkMethods(sname, &methods, false)
	}
}
