// Package startup provides a runtime for connecting to the graphics engine.
package startup

import (
	"iter"

	"graphics.gd/classdb"
	EngineClass "graphics.gd/classdb/Engine"
	MainLoopClass "graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
)

var mainloop MainLoopClass.Interface

var intialized = make(chan struct{})
var shutdown = make(chan struct{})

// MainLoop uses the given struct as the main loop implementation. This will take care of initialising
// the Go runtime correctly, blocks until the main loop has shutdown.
func MainLoop(loop MainLoopClass.Interface) {
	if pause_main != nil {
		if EngineClass.IsEditorHint() {
			stop_main()
		}
		theMainFunctionIsWaitingForTheEngineToShutDown = true
	} else {
		<-intialized
	}
	classdb.Register[goMainLoop]()
	mainloop = loop
	if pause_main != nil {
		pause_main(false) // We pause here until the engine has fully started up.
	} else {
		<-shutdown
	}
}

var theMainFunctionIsWaitingForTheEngineToShutDown = false

// Deprecated: Use [Scene] instead.
func Engine() { Scene() }

// Deprecated: Use [LoadingScene] instead.
func Loader() { LoadingScene() }

// Scene starts up the SceneTree and blocks until the engine shuts down.
func Scene() {
	if !loadingSceneWasCalled {
		LoadingScene()
	}
	if pause_main != nil {
		theMainFunctionIsWaitingForTheEngineToShutDown = true
		pause_main(false)
	} else {
		<-shutdown
	}
}

var loadingSceneWasCalled bool

// LoadingScene starts up loading the main scene after this function is called, all
// graphics functions will be available to use.
//
// A subsequent call to [Scene] is required to startup the scene.
//
// Blocks indefinitely if the editor is running. As such, make sure to register all
// editor-accessible classes before calling this function if you want them to be
// available in the editor.
func LoadingScene() {
	loadingSceneWasCalled = true
	classdb.Register[goSceneTree]()
	if pause_main != nil {
		gd.NewCallable(func() {
			resume_main()
		}).CallDeferred()
		pause_main(false)
		if EngineClass.IsEditorHint() {
			stop_main()
		}
	} else {
		<-intialized
		var loaded = make(chan struct{})
		gd.NewCallable(func() {
			if !hasLoaded {
				close(loaded)
				hasLoaded = true
			}
		}).CallDeferred()
		<-loaded
	}
}

var hasLoaded bool

// There are two main loop implementations, we decide on which one to use based on
// what startup functions are called in the main function.

type goMainLoop struct {
	classdb.Extension[goMainLoop, MainLoopClass.Instance] `gd:"GoMainLoop"`
}
type goSceneTree struct {
	classdb.Extension[goSceneTree, SceneTree.Instance] `gd:"GoMainLoop"`
}

var main_loop_initialized = make(chan struct{})

// Called once during initialization.
func (loop goMainLoop) Initialize() {
	if mainloop != nil {
		mainloop.Initialize()
	} else if pause_main != nil {
		resume_main()
	}
	Callable.Cycle()
}

// Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (loop goMainLoop) PhysicsProcess(delta Float.X) bool {
	if mainloop != nil {
		return mainloop.PhysicsProcess(delta)
	}
	return false
}

var dt Float.X

var frame_ready = make(chan bool)

// Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (loop goMainLoop) Process(delta Float.X) bool {
	defer Callable.Cycle()
	defer pointers.Cycle()
	if mainloop != nil {
		return mainloop.Process(delta)
	}
	dt = delta
	if pause_main != nil {
		close, _ := resume_main()
		return close
	}
	frame_ready <- false
	return <-frame_ready
}

var main_loop_shutdown = make(chan struct{})

// Called before the program exits.
func (loop goMainLoop) Finalize() {
	if mainloop != nil {
		mainloop.Finalize()
	} else if pause_main != nil {
		resume_main()
	} else {
		close(main_loop_shutdown)
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
	classdb.Register[goMainLoop]()
	if pause_main != nil {
		if EngineClass.IsEditorHint() {
			stop_main()
		}
		pause_main(false) // We pause here until the engine has fully started up.
		return func(yield func(Float.X) bool) {
			pause_main(false) // we pause here until the MainLoop initialize function is called.
			for {
				pause_main(false) // we pause here until the next frame is ready (next Process callback).
				if !yield(dt) {
					break
				}
			}
			pause_main(true) // we pause here until the engine has fully shut down.
		}
	} else {
		<-frame_ready
		return func(yield func(Float.X) bool) {
			frame_ready <- false
			for {
				<-frame_ready // we pause here until the next frame is ready (next Process callback).
				if !yield(dt) {
					frame_ready <- true
					break
				}
				frame_ready <- false
			}
			<-main_loop_shutdown
		}
	}
}

var (
	pause_main  func(bool) bool
	resume_main func() (bool, bool)
	stop_main   func()
)
