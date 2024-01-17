package gdextension

import (
	"errors"
	"fmt"
	"unsafe"

	internal "grow.graphics/gd/internal"
	"runtime.link/api"
	"runtime.link/api/call"
	"runtime.link/api/stub"
)
import "C"

var classDB internal.ExtensionToken
var dlsymGD func(string) unsafe.Pointer

var background internal.Context

// Link returns a handle to the [API] and the global [ClassDB].
// The [bool] return value is [true] if the API has been
// linked with Godot successfully.
func Link() (internal.Context, bool) {
	if dlsymGD == nil {
		*godot = api.Import[internal.API](stub.API, "", errors.New("gdextension not linked"))
		return internal.NewContext(godot), false
	}
	return background, true
}

var (
	godot = new(internal.API)
)

//export loadExtension
func loadExtension(lookupFunc, classes, configuration unsafe.Pointer) uint8 {
	dlsym, err := call.Make[func(string) unsafe.Pointer](lookupFunc, "func(&char)~void")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	dlsymGD = dlsym
	classDB = internal.ExtensionToken(classes)

	*godot = api.Import[internal.API](call.API, "", call.Options{
		LookupSymbol: func(name string) (unsafe.Pointer, error) {
			sym := dlsymGD(name)
			if sym == nil {
				return nil, errors.New("'" + name + "' symbol not found")
			}
			return sym, nil
		},
	})
	godot.ExtensionToken = classDB

	background = internal.NewContext(godot)

	init := (*internal.ExtensionInitialization[uintptr])(configuration)
	*init = internal.ExtensionInitialization[uintptr]{}
	init.MinimumInitializationLevel = internal.GDExtensionInitializationLevelScene
	init.Initialize.Set(func(userdata uintptr, level internal.GDExtensionInitializationLevel) {
		godot.Link(level)
		if level == 2 {
			main()
		}
	})
	init.Deinitialize.Set(func(userdata uintptr, level internal.GDExtensionInitializationLevel) {
		if level == 3 {
			init.Initialize.Free()
		}
		if level == 0 {
			background.Free()
			init.Deinitialize.Free()
		}
	})
	return 1
}

//go:linkname main main.main
func main()
