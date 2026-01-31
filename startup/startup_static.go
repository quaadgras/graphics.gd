//go:build musl || archive

package startup

// #include "libgodot.h"
// GDExtensionBool cgo_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization);
//
// extern void go_main();
//
// int main(int argc, char *argv[]) {
//		go_main();
// }
import "C"
import (
	"os"

	"graphics.gd/classdb/Startup"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/pointers"
)

// libgodot MVP
func init() {
	running_as_gdextension = true
	weNeedToStartupTheEngine = true
	startupTheEngine = func() {
		var cargs []*C.char
		for _, arg := range os.Args {
			cargs = append(cargs, C.CString(arg))
		}
		ptr := C.libgodot_create_godot_instance(C.int(len(os.Args)), &cargs[0], (C.GDExtensionInitializationFunction)(C.cgo_extension_init))
		if ptr == nil {
			return
		}
		engine := Startup.Instance([1]gdclass.Startup{gdclass.NewStartup(pointers.Raw[gd.Object]([3]uint64{uint64(uintptr(ptr))}))})
		engine.Start()
		for !engine.Iteration() {
		}
		os.Exit(0)
	}
}
