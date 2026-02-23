//go:build cgo

//go:generate go run ./internal/cmd/generate
//go:generate go fmt .
package startup

/*

#include "startup_cgo.h"
*/
import "C"

import (
	"fmt"
	"iter"
	"os"
	"runtime/debug"
	"slices"
	"testing"

	"graphics.gd/classdb"
	EngineClass "graphics.gd/classdb/Engine"
	"graphics.gd/classdb/SceneTree"
	internal "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/threadcheck"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
)

var initDone = false
var exitDone = false
var toolUsed = false

func init() {
	gdextension.On.Engine = gdextension.CallbacksForEngine{
		Init: func(level gdextension.InitializationLevel) {
			if startup == nil {
				startup = engineLoadingSharedGo{}
				// little hack to enable `gd test` to work, we strip away the headless flag
				// so that 'go test' doesn't complain on startup.
				for i := 0; i < len(os.Args); i++ {
					switch os.Args[i] {
					case "--headless", "-race":
						os.Args = append(os.Args[:i], os.Args[i+1:]...)
						i--
					}
				}
			}
			internal.Init(level)
			if level == 2 && !initDone {
				for _, fn := range internal.StartupFunctions {
					fn()
				}
				if _, ok := startup.(engineLoadingSharedGo); ok {
					if testing.Testing() {
						classdb.Register[goSceneTree]()
					} else {
						resume_main, stop_main = iter.Pull(call_main_in_steps())
						resume_main()
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
				for _, cleanup := range slices.Backward(internal.Cleanups()) {
					cleanup()
				}
				pointers.Cycle()
				pointers.Cycle()
				if theMainFunctionIsWaitingForTheEngineToShutDown {
					resume_main()
				}
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
		if testing.Testing() && !toolUsed {
			go main()
		}
		Callable.Cycle()
		if EngineClass.IsEditorHint() {
			editorSetup()
		}
		if pause_main != nil {
			resume_main()
		}
	}
}

type goMain struct {
	SceneTree.Extension[goSceneTree] `gd:"GoMainLoop"`
}

func (loop goMain) Initialize() {
	Callable.Cycle()
	resume_main()
}

func (loop goMain) PhysicsProcess(delta Float.X) bool {
	return false
}

func (loop goMain) Process(delta Float.X) bool {
	defer Callable.Cycle()
	defer keep_reachable_instances_alive()
	defer pointers.Cycle()
	dt = delta
	close, _ := resume_main()
	return close
}

func (loop goMain) Finalize() {
	resume_main()
}
