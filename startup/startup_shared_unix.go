//go:build cgo && unix

package startup

/*
#cgo linux LDFLAGS: -Wl,-rpath=$ORIGIN
#include <gdextension_interface.h>
#include <dlfcn.h>
GDExtensionBool cgo_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization);


typedef GDExtensionObjectPtr (*libgodot_create_godot_instance_func_t)(int p_argc, char *p_argv[], GDExtensionInitializationFunction p_init_func);

GDExtensionObjectPtr call_libgodot_create_godot_instance(void *p_func, int p_argc, char *p_argv[], GDExtensionInitializationFunction p_init_func) {
	libgodot_create_godot_instance_func_t func = (libgodot_create_godot_instance_func_t)p_func;
	return func(p_argc, p_argv, p_init_func);
}
*/
import "C"
import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"graphics.gd/classdb/Startup"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/pointers"
)

func loadEngineAsSharedLibrary() {
	weNeedToStartupTheEngine = true
	var ext string
	switch runtime.GOOS {
	case "linux":
		ext = ".so"
	case "windows":
		ext = ".dll"
	case "darwin":
		ext = ".dylib"
	}
	path := []byte("libgodot" + ext + "\000")
	init := []byte("libgodot_create_godot_instance\000")
	var libgodot = C.dlopen((*C.char)(unsafe.Pointer(&path[0])), C.RTLD_LAZY)
	if libgodot == nil {
		fmt.Fprintln(os.Stderr, "failed to load libgodot"+ext)
		os.Exit(1)
	}
	var libgodot_create_godot_instance = C.dlsym(libgodot, (*C.char)(unsafe.Pointer(&init[0])))
	var cargs []*C.char
	for _, arg := range os.Args {
		cargs = append(cargs, C.CString(arg))
	}
	ptr := C.call_libgodot_create_godot_instance(libgodot_create_godot_instance, C.int(len(os.Args)), &cargs[0], (C.GDExtensionInitializationFunction)(C.cgo_extension_init))
	if ptr == nil {
		return
	}
	engine := Startup.Instance([1]gdclass.Startup{gdclass.NewStartup(pointers.Raw[gd.Object]([3]uint64{uint64(uintptr(ptr))}))})
	engine.Start()
	for !engine.Iteration() {
	}
	os.Exit(0)
}
