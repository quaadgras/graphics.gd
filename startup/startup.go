// Package startup provides a runtime for connecting to the graphics engine.
package startup

import (
	"iter"
	"runtime"

	"graphics.gd/classdb"
	EngineClass "graphics.gd/classdb/Engine"
	MainLoopClass "graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/ProjectSettings"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Float"
)

var mainloop MainLoopClass.Interface

// MainLoop uses the given struct as the main loop implementation. This will take care of initialising
// the Go runtime correctly, blocks until the main loop has shutdown.
func MainLoop(loop MainLoopClass.Interface) {
	if EngineClass.IsEditorHint() {
		doneInit <- struct{}{}
		runtime.Goexit()
	}
	startupEngine = true
	classdb.Register[goMainLoop]()
	className := classdb.NameFor[goMainLoop]()
	ProjectSettings.SetInitialValue("application/run/main_loop_type", className)
	ProjectSettings.SetSetting("application/run/main_loop_type", className)
	mainloop = loop
	doneInit <- struct{}{}
	<-done
}

var done = make(chan struct{})

// Wait until the engine has been fully started up.
func Wait() {
	if EngineClass.IsEditorHint() {
		doneInit <- struct{}{}
		runtime.Goexit()
	}
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
	if EngineClass.IsEditorHint() {
		doneInit <- struct{}{}
		runtime.Goexit()
	}
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
func (loop goMainLoop) Initialize() {
	if mainloop != nil {
		mainloop.Initialize()
	} else {
		close(starts)
	}
}

// Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (loop goMainLoop) PhysicsProcess(delta Float.X) bool {
	if mainloop != nil {
		return mainloop.PhysicsProcess(delta)
	}
	return false
}

// Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (loop goMainLoop) Process(delta Float.X) bool {
	defer pointers.Cycle()
	if mainloop != nil {
		return mainloop.Process(delta)
	}
	frames <- delta
	return <-breaks
}

// Called before the program exits.
func (loop goMainLoop) Finalize() {
	if mainloop != nil {
		mainloop.Finalize()
	} else {
		close(finish)
	}
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
	if EngineClass.IsEditorHint() {
		doneInit <- struct{}{}
		runtime.Goexit()
	}
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
