package Startup

import (
	gdunsafe "graphics.gd"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
)

func init() {
	gd.LinkStartup = func() {
		sname = gdextension.Host.Strings.Intern.UTF8("GodotInstance")
		otype = gdunsafe.ObjectTypeTag(gdunsafe.StringName(sname[0]))
		gd.LinkMethods(sname, &methods, false)
	}
}
