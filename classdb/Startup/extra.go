package Startup

import (
	gdunsafe "graphics.gd"
	gd "graphics.gd/internal"
)

func init() {
	gd.LinkStartup = func() {
		sname = gdunsafe.UTF8.Intern("GodotInstance")
		otype = gdunsafe.ObjectTypeTag(sname)
		gd.LinkMethods(sname, &methods, false)
	}
}
