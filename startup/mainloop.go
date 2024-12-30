// Package startup provides a runtime for connecting to the graphics engine.
package startup

import (
	"graphics.gd/defined"
	gd "graphics.gd/internal"
	MainLoopClass "graphics.gd/objects/MainLoop"
	"graphics.gd/objects/ProjectSettings"
	"graphics.gd/variant/String"
)

// MainLoop designates the given struct embedding defined.Object[T, MainLoop.Instance]
// as the main loop entrypoint. After the graphics engine has started up, it will use
// the given struct to execute as the main loop.
func MainLoop[T defined.ExtensionTo[M], M MainLoopClass.Any]() {
	defined.InEditor[T]()
	className := gd.NewVariant(String.New(defined.NameFor[T]()))
	main_loop_type := String.New("application/run/main_loop_type")
	ProjectSettings.Advanced().SetInitialValue(main_loop_type, className)
	ProjectSettings.Advanced().SetSetting(main_loop_type, className)
}
