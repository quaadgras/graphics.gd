package gdextension

import (
	"errors"
	"fmt"
	"unsafe"

	"grow.graphics/gd"
	"runtime.link/api"
	"runtime.link/api/call"
	"runtime.link/api/stub"
)
import "C"

var classDB ClassDB
var dlsymGD func(string) unsafe.Pointer

// Link returns a handle to the [API] and the global [ClassDB].
// The [bool] return value is [true] if the API has been
// linked with Godot successfully.
func Link() (API, ClassDB, bool) {
	if dlsymGD == nil {
		return api.Import[API](stub.API, "", errors.New("gdextension not linked")), 0, false
	}
	return api.Import[API](call.API, "", call.Options{
		LookupSymbol: func(name string) (unsafe.Pointer, error) {
			sym := dlsymGD(name)
			if sym == nil {
				return nil, errors.New("'" + name + "' symbol not found")
			}
			return sym, nil
		},
	}), classDB, true
}

//export loadExtension
func loadExtension(lookupFunc, classes, configuration unsafe.Pointer) uint8 {
	dlsym, err := call.Make[func(string) unsafe.Pointer](lookupFunc, "func(&char)~void")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	dlsymGD = dlsym
	classDB = ClassDB(classes)

	GDextension := api.Import[API](call.API, "", call.Options{
		LookupSymbol: func(name string) (unsafe.Pointer, error) {
			sym := dlsym(name)
			if sym == nil {
				return nil, errors.New("'" + name + "' symbol not found")
			}
			return sym, nil
		},
	})

	var name gd.StringName
	GDextension.StringNames.New(&name, "AudioFrame")

	GDextension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&name))

	init := (*ExtensionInitialization[uintptr])(configuration)
	*init = ExtensionInitialization[uintptr]{}
	init.MinimumInitializationLevel = gd.ExtensionInitializationLevelScene
	init.Initialize.Set(func(userdata uintptr, level gd.ExtensionInitializationLevel) {
		if level == 2 {
			main()
		}
		/*if level == 3 {

			// playground
			GDextension.StringNames.New(&name, "Engine")
			var Engine = GDextension.Object.GetSingleton(&name)

			var method gd.StringName
			GDextension.StringNames.New(&method, "get_license_text")
			var bind = GDextension.ClassDB.GetMethodBind(&name, &method, 201670096)

			GDextension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&name))
			GDextension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&method))

			fmt.Println("bind", bind)

			fmt.Println(Engine)

			var result gd.String
			GDextension.Object.UnsafeCall(bind, Engine, nil, unsafe.Pointer(&result))
			fmt.Println("result=", result)

			var buf = make([]byte, 1024)
			GDextension.Strings.Get(&result, buf)
			fmt.Println(string(buf))
		}*/
	})
	init.Deinitialize.Set(func(userdata uintptr, level gd.ExtensionInitializationLevel) {
		if level == 3 {
			init.Initialize.Free()
		}
		if level == 0 {
			init.Deinitialize.Free()
		}
	})
	return 1
}

//go:linkname main main.main
func main()
