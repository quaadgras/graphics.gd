// Package startup provides a runtime for connecting to the graphics engine.
package startup

import (
	"graphics.gd/defined"
	MainLoopClass "graphics.gd/objects/MainLoop"
	"graphics.gd/objects/ProjectSettings"
)

// MainLoop designates the given struct embedding defined.Object[T, MainLoop.Instance]
// as the main loop entrypoint. After the graphics engine has started up, it will use
// the given struct to execute as the main loop.
func MainLoop[T defined.ExtensionTo[M], M MainLoopClass.Any]() {
	defined.InEditor[T]()
	className := defined.NameFor[T]()
	ProjectSettings.SetInitialValue("application/run/main_loop_type", className)
	ProjectSettings.SetSetting("application/run/main_loop_type", className)
}
