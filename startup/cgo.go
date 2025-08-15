//go:build cgo

package startup

import (
	"errors"
	"fmt"
	"iter"
	"os"
	"runtime/debug"
	"unsafe"

	internal "graphics.gd/internal"

	"runtime.link/api"
	"runtime.link/api/stub"
)

//#include "gdextension_interface.h"
//#include "startup_cgo_v2.h"
import "C"

// little hack to enable `gd test` to work, we strip away the headless flag
// so that 'go test' doesn't complain on startup.
func init() {
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "--headless" {
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
		}
	}
}

var classDB internal.ExtensionToken
var dlsymGD func(string) unsafe.Pointer

func init() {
	internal.Global = api.Import[internal.API](stub.API, "", errors.New("gdextension not linked"))
}

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	C.cgo_extension_init(
		*(*C.GDExtensionInterfaceGetProcAddress)(unsafe.Pointer(&lookupFunc)),
		*(*C.GDExtensionClassLibraryPtr)(classes),
		(*C.GDExtensionInitialization)(configuration),
	)
	dlsymGD = func(s string) unsafe.Pointer {
		return get_proc_address(lookupFunc, s)
	}
	classDB = internal.ExtensionToken(classes)
	internal.Global.ExtensionToken = classDB
	linkCGO(&internal.Global)
	internal.Linked = true
	init := (*initialization)(configuration)
	*init = initialization{}
	init.minimum_initialization_level = initializationLevel(internal.GDExtensionInitializationLevelScene)
	doInitialization(init)
	return 1
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
