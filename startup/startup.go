// Package startup provides a runtime for connecting to the graphics engine.
package startup

import (
	"iter"

	"graphics.gd/classdb"
	MainLoopClass "graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/ProjectSettings"
	gd "graphics.gd/internal"
	"graphics.gd/variant/Float"
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

type goMainLoop struct {
	classdb.Extension[goMainLoop, MainLoopClass.Instance] `gd:"GoMainLoop"`
}

var frames = make(chan Float.X)
var breaks = make(chan bool)
var finish = make(chan struct{})
var starts = make(chan struct{})

// Called once during initialization.
func (goMainLoop) Initialize() {
	close(starts)
}

// Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (goMainLoop) PhysicsProcess(delta Float.X) bool {
	return false
}

// Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (goMainLoop) Process(delta Float.X) bool {
	frames <- delta
	return <-breaks
}

// Called before the program exits.
func (goMainLoop) Finalize() {
	close(finish)
}

// Rendering waits for the engine to startup and returns a frame iterator for the primary viewport that is
// ready for rendering. The iterator will block until the engine shuts down.
//
//		func main() {
//			frames := startup.Rendering()
//	    	// init.
//			for frame := range frames {
//				// render frame
//			}
//			// finalize
//		}
func Rendering() iter.Seq[Float.X] {
	classdb.Register[goMainLoop]()
	className := classdb.NameFor[goMainLoop]()
	ProjectSettings.SetInitialValue("application/run/main_loop_type", className)
	ProjectSettings.SetSetting("application/run/main_loop_type", className)
	doneInit <- struct{}{}
	<-starts
	return func(yield func(Float.X) bool) {
		for frame := range frames {
			if !yield(frame) {
				breaks <- true
				break
			}
			breaks <- false
		}
		<-finish
	}
}
