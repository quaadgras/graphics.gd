package gdextension

import (
	"errors"
	"os"
	"runtime"
	"unsafe"

	internal "grow.graphics/gd/internal"

	"runtime.link/api"
	"runtime.link/api/stub"
	"runtime.link/mmm"
)

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

type pinner mmm.Pointer[runtime.Pinner, pinner, [0]uintptr]

func (p pinner) Free() { mmm.API(p).Unpin(); mmm.End(p) }

var classDB internal.ExtensionToken
var dlsymGD func(string) unsafe.Pointer

var background = internal.NewContext(&godot)

// Link returns a handle to the [API] and the global [ClassDB].
// The [bool] return value is [true] if the API has been
// linked with Godot successfully.
func Link() (internal.Lifetime, bool) {
	if dlsymGD == nil {
		return background, false
	}
	return background, true
}

var (
	godot = api.Import[internal.API](stub.API, "", errors.New("gdextension not linked"))
)

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	dlsymGD = func(s string) unsafe.Pointer {
		return get_proc_address(lookupFunc, s)
	}
	classDB = internal.ExtensionToken(classes)
	godot.ExtensionToken = classDB
	linkCGO(&godot)
	init := (*initialization)(configuration)
	*init = initialization{}
	init.minimum_initialization_level = initializationLevel(internal.GDExtensionInitializationLevelScene)
	doInitialization(init)
	return 1
}

//go:linkname main main.main
func main()
