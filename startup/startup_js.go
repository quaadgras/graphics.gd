package startup

import (
	gd "graphics.gd/internal"
	internal "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

func init() {
	gdextension.On.Engine = gdextension.CallbacksForEngine{
		Init: func(level gdextension.InitializationLevel) {
			internal.Linked = true
			gd.Global.Init(level)
			if level == 2 {
				for _, fn := range gd.StartupFunctions {
					fn()
				}
				close(intialized)
				for _, fn := range gd.PostStartupFunctions {
					fn()
				}
			}
		},
		Exit: func(level gdextension.InitializationLevel) {
			if level == 2 {
				for _, cleanup := range gd.Cleanups() {
					cleanup()
				}
				pointers.Cycle()
				pointers.Cycle()
				close(shutdown)
				internal.Linked = false
			}
		},
	}
}
