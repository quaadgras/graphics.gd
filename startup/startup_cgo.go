//go:build cgo

//go:generate go run ./internal/cmd/generate
//go:generate go fmt .
package startup

import "C"

import (
	"fmt"
	"iter"
	"os"
	"runtime"
	"runtime/debug"
	"slices"
	"testing"

	_ "graphics.gd"

	"graphics.gd/classdb"
	EngineClass "graphics.gd/classdb/Engine"
	"graphics.gd/classdb/SceneTree"
	internal "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadcheck"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
)

var initDone = false
var exitDone = false
var toolUsed = false
var testMainStarted = false

func init() {
	gdextension.On.Engine = gdextension.CallbacksForEngine{
		Init: func(level gdextension.InitializationLevel) {
			if startup == nil {
				startup = engineLoadingSharedGo{}
				// little hack to enable `gd test` to work, we strip away the headless flag
				// so that 'go test' doesn't complain on startup. The engine has already
				// taken its own copy of the arguments by now, so it still sees them —
				// this only keeps them away from the testing package's flag parsing,
				// which is what lets the test editor also serve as the importer.
				for i := 0; i < len(os.Args); i++ {
					switch os.Args[i] {
					case "--headless", "--import", "-race":
						os.Args = append(os.Args[:i], os.Args[i+1:]...)
						i--
					}
				}
			}
			internal.Init(level)
			if level == 0 {
				initJumponly()
			}
			if level == 2 && !initDone {
				ring.Threads.Open()
				for _, fn := range internal.StartupFunctions {
					fn()
				}
				if _, ok := startup.(engineLoadingSharedGo); ok {
					if testing.Testing() {
						classdb.Register[goSceneTree]()
					} else {
						resume_main, stop_main = iter.Pull(call_main_in_steps())
						resumeMain()
					}
				}
				for _, fn := range internal.PostStartupFunctions {
					fn()
				}
				initDone = true
			}
		},
		Exit: func(level gdextension.InitializationLevel) {
			if !exitDone && level == 2 {
				// Run any still-buffered cross-thread calls while the engine
				// is alive, then poison the ring so goroutines parked on it
				// wake up instead of sleeping forever.
				ring.Threads.Close()
				if theMainFunctionIsWaitingForTheEngineToShutDown {
					resumeMain()
				}
				for _, cleanup := range slices.Backward(internal.Cleanups()) {
					cleanup()
				}
				pointers.Cycle()
				pointers.Cycle()
				internal.Linked = false
				exitDone = true
			}
		},
	}
}

//go:linkname main main.main
func main()

//export go_main
func go_main() {
	// libgodot calls go_main on the engine's main thread, which is not
	// necessarily the thread that initialised the Go runtime. Adopt it now:
	// engine calls made before the first frame (which re-runs Init) would
	// otherwise be routed through the cross-thread dispatch ring and block
	// waiting for this very thread to drain them.
	threadcheck.Init()
	if testing.Testing() {
		Scene()
	} else {
		main()
	}
}

// call_main_in_steps calls the main function on the main thread in steps,
// so that we can yield control back to the engine every frame and before
// and after startup.
func call_main_in_steps() iter.Seq[bool] {
	return func(yield func(bool) bool) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				debug.PrintStack()
			}
		}()
		pause_main = yield
		main()
	}
}

var (
	pause_main  func(bool) bool
	resume_main func() (bool, bool)
	stop_main   func()
)
var theMainFunctionIsWaitingForTheEngineToShutDown = false

// resumeMain resumes the main-function coroutine, restoring the OS-thread
// lock state it was created with for the duration of the switch.
//
// A coroutine may only be switched on the same thread and with the same
// internal and external thread-lock counts it was created with, or
// runtime.coroswitch_m throws "coro: OS thread locking must match locking at
// coroutine creation". iter.Pull creates this one during extension init, so
// it records no external lock. Resident-callback mode then takes one for the
// life of the process — fastcbFrame calls runtime.LockOSThread when it
// engages, which is what keeps the resident goroutine bound to the engine
// thread in the gaps between callbacks — so every resume after the first
// frame would switch with a count the coroutine has never seen. Rendering
// mode died that way on Windows on its first frame, once
// quaadgras/graphics.gd#321 stopped crashing ahead of it.
//
// Dropping that external lock across the switch is safe: the callback we are
// inside holds an internal lock (the resident paths take the same one the
// stock path does — see the fastcb overlay, which relies on it for exactly
// this reason), so the goroutine stays bound to this m while the coroutine
// runs, and the coroutine hands control back to it on this same thread by
// construction. Retaking the lock before the callback returns to C restores
// residency's invariant for the gap between callbacks.
func resumeMain() (bool, bool) {
	if !fastcbEngaged {
		return resume_main()
	}
	runtime.UnlockOSThread()
	closing, ok := resume_main()
	runtime.LockOSThread()
	return closing, ok
}

type engineLoadingSharedGo struct{}

func (engineLoadingSharedGo) Start() {
	pause_main(false)
	if EngineClass.IsEditorHint() {
		stop_main()
	}
}

func (engineLoadingSharedGo) Scene() {
	theMainFunctionIsWaitingForTheEngineToShutDown = true
	pause_main(false)
}

func (engineLoadingSharedGo) Rendering() iter.Seq[Float.X] {
	classdb.Register[goMain]()
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
}

func init() {
	gdextension.On.MainLoop.FirstFrame = func() {
		threadcheck.Init()
		// FirstFrame can fire more than once (e.g. on android), so guard the
		// test main: spawning it per frame would run the whole suite repeatedly.
		if testing.Testing() && !toolUsed && !testMainStarted {
			testMainStarted = true
			// On platforms where the process's stdout/stderr are discarded
			// (android), route the test output somewhere the harness can read
			// it back. No-op elsewhere. Must run before the test main starts.
			prepareTestOutput()
			go main()
		}
		Callable.Cycle()
		if EngineClass.IsEditorHint() {
			editorSetup()
		}
		if pause_main != nil {
			resumeMain()
		}
	}
}

type goMain struct {
	SceneTree.Extension[goSceneTree] `gd:"GoMainLoop"`
}

func (loop goMain) Initialize() {
	Callable.Cycle()
	resumeMain()
}

func (loop goMain) PhysicsProcess(delta Float.X) bool {
	return false
}

// Process intentionally does NOT run pointers.Cycle: the only per-frame
// pointer collection happens in EveryFrame (startup/garbage_collector.go),
// immediately after the cross-thread ring drain. A second cycle here would
// mean two cycles per frame with no drain in between, so a value used in a
// call buffered by a goroutine could expire AND be freed before the drain
// that executes the call.
func (loop goMain) Process(delta Float.X) bool {
	defer Callable.Cycle()
	defer keep_reachable_instances_alive()
	dt = delta
	close, _ := resumeMain()
	return close
}

func (loop goMain) Finalize() {
	resumeMain()
}
