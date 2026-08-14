// Package startup provides a runtime for connecting to the graphics engine.
package startup

import (
	"iter"
	"time"

	"graphics.gd/classdb"
	EngineClass "graphics.gd/classdb/Engine"
	MainLoopClass "graphics.gd/classdb/MainLoop"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Startup"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Dictionary"
	"graphics.gd/variant/Float"
)

var mainloop MainLoopClass.Interface
var loadingSceneWasCalled bool

// engineStarted records whether the engine has been started, for the
// startup modes where starting is a real operation: the static library
// completes the embedded engine's setup, and the shared library creates
// the engine instance outright. More than one path believes it is the one
// that starts the engine — the project's own [LoadingScene] call and,
// under -tags reloads, the host session — and doing it twice re-runs
// setup over live singletons, or makes a second engine.
var engineStarted bool

// startingEngine reports whether this is the first attempt to start the
// engine, marking it started. Every startup mode that really does start
// one guards its Start with this so a second call is a no-op.
func startingEngine() bool {
	if engineStarted {
		return false
	}
	engineStarted = true
	return true
}

// MainLoop uses the given struct as the main loop implementation. This will take care of initialising
// the Go runtime correctly, blocks until the main loop has shutdown.
func MainLoop(loop MainLoopClass.Interface) {
	classdb.Register[goMainLoop]()
	mainloop = loop
	Scene()
}

// reloadsSession replaces the Scene flow when the reloads host is
// compiled in (-tags reloads): the host owns the engine and hands all
// class registration to a hot-reloadable wasm build of the project.
var reloadsSession func()

// Scene starts up the SceneTree and blocks until the engine shuts down.
func Scene() {
	if reloadsSession != nil {
		reloadsSession()
		return
	}
	if !loadingSceneWasCalled {
		LoadingScene()
	}
	startup.Scene()
}

// LoadingScene starts up loading the main scene after this function is called, all
// graphics.gd functionality will be available to use.
//
// A subsequent call to [Scene] is required to startup the scene.
//
// Blocks indefinitely if the editor is running. As such, make sure to register all
// editor-accessible classes before calling this function if you want them to be
// available in the editor.
func LoadingScene() {
	if startup == nil {
		startup = new(engineAsSharedLibrary)
	}
	loadingSceneWasCalled = true
	if mainloop == nil {
		classdb.Register[goSceneTree]()
	}
	startup.Start()
}

// There are two main loop implementations, we decide on which one to use based on
// what startup functions are called in the main function.

type goMainLoop struct {
	MainLoopClass.Extension[goMainLoop] `gd:"GoMainLoop"`
}
type goSceneTree struct {
	SceneTree.Extension[goSceneTree] `gd:"GoMainLoop"`
}

// Called once during initialization.
func (loop goMainLoop) Initialize() {
	Callable.Cycle()
	mainloop.Initialize()
}

// Called each physics frame with the time since the last physics frame as argument ([param delta], in seconds). Equivalent to [method Node._physics_process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
func (loop goMainLoop) PhysicsProcess(delta Float.X) bool {
	return mainloop.PhysicsProcess(delta)
}

var dt Float.X

var frame_ready = make(chan bool)

// Called each process (idle) frame with the time since the last process frame as argument (in seconds). Equivalent to [method Node._process].
// If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
// Process intentionally does NOT run pointers.Cycle — see goMain.Process
// in startup_cgo.go: per-frame pointer collection must only happen right
// after the cross-thread ring drain in EveryFrame.
func (loop goMainLoop) Process(delta Float.X) bool {
	defer Callable.Cycle()
	defer keep_reachable_instances_alive()
	return mainloop.Process(delta)
}

// Called before the program exits.
func (loop goMainLoop) Finalize() {
	mainloop.Finalize()
}

// Rendering waits for the engine to startup and returns a frame iterator for the primary viewport that is
// ready for rendering. The iterator will block until the engine shuts down.
//
//		func main() {
//			frames := startup.Rendering()
//	    	// init.
//			for dt := range frames {
//				// render frame
//			}
//			// finalize
//		}
func Rendering() iter.Seq[Float.X] {
	return startup.Rendering()
}

// AsExtension requests graphics.gd to startup the library as a GDExtension suitable for
// inclusion in Godot engine projects. Please note that only a single Go runtime can be
// active within an OS process, so all Go extensions within a project should be built
// together into a single library.
func AsExtension() {
	Scene() // TODO: investigate anything else that should be setup for pure extensions.
}

// OnSuspend registers a function to be called when the application is suspended, the
// dictionary populated by this function will be available to future [OnRestore] calls.
// Individual classes can also implement their own Suspend(Dictionary.Any) method. This
// function should not mutate any internal state and the application may continue to
// run after this has been called.
func OnSuspend(func(Dictionary.Any)) {

}

// OnRestore registers a function to be called when the application is being restored, the dictionary
// dictionary will be sourced from a previous call to [OnSuspend] (potentially a different
// version) Individual classes can also implement their own Restore(Dictionary.Any) method.
func OnRestore(func(Dictionary.Any)) {

}

var startup interface {
	Start()
	Scene()
	Rendering() iter.Seq[Float.X]
}

type engineAsLibrary struct {
	Library Startup.Instance
	destroy func()
}

type engineAsSharedLibrary struct {
	engineAsLibrary
}

func (engine *engineAsLibrary) Start() {}

func (engine *engineAsLibrary) Scene() {
	for !engine.Library.Iteration() {
	}
	if engine.destroy != nil {
		engine.destroy()
	}
}

func (engine *engineAsLibrary) Rendering() iter.Seq[Float.X] {
	classdb.Register[goSceneTree]()

	startup.Start()
	if EngineClass.IsEditorHint() {
		startup.Scene()
		return func(yield func(dt Float.X) bool) {}
	}
	return func(yield func(dt Float.X) bool) {
		var dt = time.Now()
		for !engine.Library.Iteration() {
			if !yield(Float.X(time.Since(dt).Seconds())) {
				break
			}
		}
		if engine.destroy != nil {
			engine.destroy()
		}
	}
}
