//go:build musl || archive

package startup

// #include "libgodot.h"
// #include <stdlib.h>
//
// GDExtensionBool cgo_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization);
//
// extern void go_main();
// extern void go_destroy_engine();
//
// int main(int argc, char *argv[]) {
//		go_main();
//		go_destroy_engine();
// 		return 0;
// }
import "C"
import (
	"os"

	"graphics.gd/classdb/Startup"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
)

func init() {
	var cargs []*C.char
	for _, arg := range os.Args {
		if arg == "--export-release" {
			toolUsed = true
		}
		cargs = append(cargs, C.CString(arg))
	}
	ptr := C.libgodot_create_godot_instance(C.int(len(os.Args)), &cargs[0], (C.GDExtensionInitializationFunction)(C.cgo_extension_init))
	if ptr == nil {
		os.Exit(0)
		return
	}
	startup = &engineAsStaticLibrary{engineAsLibrary{
		Library: Startup.Instance([1]gdclass.Startup{gdclass.NewStartup(gdreference.RawObject(gdextension.Object(uintptr(ptr))))}),
		destroy: func() { C.libgodot_destroy_godot_instance(ptr) },
	}}
}

type engineAsStaticLibrary struct {
	engineAsLibrary
}

func (engine *engineAsStaticLibrary) Start() {
	engine.Library.Start()
}

func (engine *engineAsStaticLibrary) Scene() {
	for !engine.Library.Iteration() {
	}
}
