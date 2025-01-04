// Package startup provides a runtime for connecting to the graphics engine.
package startup

import (
	"graphics.gd/classdb"
	MainLoopClass "graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/ProjectSettings"
	gd "graphics.gd/internal"
)

// MainLoop designates the given struct embedding classdb.Extension[T, MainLoop.Instance]
// as the main loop entrypoint. After the graphics engine has started up, it will use
// the given struct to execute as the main loop. Blocks until the engine shuts down.
func MainLoop[T classdb.ExtensionTo[M], M MainLoopClass.Any]() {
	startupEngine = true
	classdb.Register[T]()
	className := classdb.NameFor[T]()
	ProjectSettings.SetInitialValue("application/run/main_loop_type", className)
	ProjectSettings.SetSetting("application/run/main_loop_type", className)
	doneInit <- struct{}{}
	<-done
}

var done = make(chan struct{})

// Wait until the engine has been fully started up.
func Wait() {
	wait := make(chan struct{})
	doneInit <- struct{}{}
	gd.NewCallable(func() {
		close(wait)
	}).CallDeferred()
	<-wait
}

var startupEngine = false

// Engine starts up the engine and blocks until it shuts down.
func Engine() {
	startupEngine = true
	doneInit <- struct{}{}
	<-done
}
