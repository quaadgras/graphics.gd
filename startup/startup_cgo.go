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

	internal "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

func init() {
	// little hack to enable `gd test` to work, we strip away the headless flag
	// so that 'go test' doesn't complain on startup.
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "--headless" {
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
		}
	}
	gdextension.On.Engine = gdextension.CallbacksForEngine{
		Init: func(level gdextension.InitializationLevel) {
			internal.Linked = true
			internal.Init(level)
			if level == 2 {
				for _, fn := range internal.StartupFunctions {
					fn()
				}
				close(intialized)
				resume_main, stop_main = iter.Pull(call_main_in_steps())
				resume_main()
				for _, fn := range internal.PostStartupFunctions {
					fn()
				}
			}
		},
		Exit: func(level gdextension.InitializationLevel) {
			if level == 2 {
				for _, cleanup := range internal.Cleanups() {
					cleanup()
				}
				pointers.Cycle()
				pointers.Cycle()
				if theMainFunctionIsWaitingForTheEngineToShutDown {
					resume_main()
				}
				internal.Linked = false
			}
		},
	}
}

//go:linkname main main.main
func main()

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
