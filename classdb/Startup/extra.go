package Startup

import (
	gdunsafe "graphics.gd"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
)

func init() {
	gd.LinkStartup = func() {
		sname = gdextension.StringName{gdextension.Pointer(gdunsafe.UTF8.Intern("GodotInstance"))}
		otype = gdunsafe.ObjectTypeTag(gdunsafe.StringName(sname[0]))
		gd.LinkMethods(sname, &methods, false)
	}
}
