package startup

import (
	"errors"
	"os"
	"unsafe"

	gd "graphics.gd/internal"
	internal "graphics.gd/internal"

	"runtime.link/api"
	"runtime.link/api/stub"
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

var classDB internal.ExtensionToken
var dlsymGD func(string) unsafe.Pointer

func init() {
	internal.Global = api.Import[internal.API](stub.API, "", errors.New("gdextension not linked"))
}

//export loadExtension
func loadExtension(lookupFunc uintptr, classes, configuration unsafe.Pointer) uint8 {
	dlsymGD = func(s string) unsafe.Pointer {
		return get_proc_address(lookupFunc, s)
	}
	classDB = internal.ExtensionToken(classes)
	internal.Global.ExtensionToken = classDB
	linkCGO(&internal.Global)
	init := (*initialization)(configuration)
	*init = initialization{}
	init.minimum_initialization_level = initializationLevel(internal.GDExtensionInitializationLevelScene)
	doInitialization(init)
	return 1
}

// LibraryPath is the path to the shared library that contains the GD extension.
func LibraryPath() string {
	return gd.Global.GetLibraryPath(gd.Global.ExtensionToken).String()
}

//go:linkname main main.main
func main()
