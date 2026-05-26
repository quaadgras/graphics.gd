//go:build unix && cgo && !android

package startup

/*
#cgo linux LDFLAGS: -Wl,-rpath=$ORIGIN
#include "../gdextension_interface.h"
#include <dlfcn.h>
#include <stdlib.h>
GDExtensionBool cgo_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization);


typedef GDExtensionObjectPtr (*libgodot_create_godot_instance_func_t)(int p_argc, char *p_argv[], GDExtensionInitializationFunction p_init_func);
typedef void (*libgodot_destroy_godot_instance_func_t)(GDExtensionObjectPtr p_godot_instance);

GDExtensionObjectPtr call_libgodot_create_godot_instance(void *p_func, int p_argc, char *p_argv[], GDExtensionInitializationFunction p_init_func) {
	libgodot_create_godot_instance_func_t func = (libgodot_create_godot_instance_func_t)p_func;
	return func(p_argc, p_argv, p_init_func);
}

void call_libgodot_destroy_godot_instance(void *p_func, GDExtensionObjectPtr p_godot_instance) {
	libgodot_destroy_godot_instance_func_t func = (libgodot_destroy_godot_instance_func_t)p_func;
	func(p_godot_instance);
}

static void *g_libgodot_destroy_func = 0;
static GDExtensionObjectPtr g_godot_instance = 0;

static void destroy_at_exit(void) {
	void *func = g_libgodot_destroy_func;
	GDExtensionObjectPtr inst = g_godot_instance;
	g_libgodot_destroy_func = 0;
	g_godot_instance = 0;
	if (func && inst) {
		((libgodot_destroy_godot_instance_func_t)func)(inst);
	}
}

void register_destroy_atexit(void *p_destroy_func, GDExtensionObjectPtr p_godot_instance) {
	static int registered = 0;
	g_libgodot_destroy_func = p_destroy_func;
	g_godot_instance = p_godot_instance;
	if (!registered) {
		registered = 1;
		atexit(destroy_at_exit);
	}
}
*/
import "C"
import (
	"fmt"
	"os"
	"runtime"
	"unsafe"

	"graphics.gd/classdb/Startup"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
)

func (engine *engineAsSharedLibrary) Start() {
	var ext string
	switch runtime.GOOS {
	case "linux":
		ext = ".so"
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
	destroyName := []byte("libgodot_destroy_godot_instance\000")
	var libgodot_destroy_godot_instance = C.dlsym(libgodot, (*C.char)(unsafe.Pointer(&destroyName[0])))
	var cargs []*C.char
	for _, arg := range os.Args {
		cargs = append(cargs, C.CString(arg))
	}
	ptr := C.call_libgodot_create_godot_instance(libgodot_create_godot_instance, C.int(len(os.Args)), &cargs[0], (C.GDExtensionInitializationFunction)(C.cgo_extension_init))
	if ptr == nil {
		return
	}
	engine.Library = Startup.Instance([1]gdclass.Startup{gdclass.NewStartup(gdreference.RawObject(gdextension.Object(uintptr(ptr))))})
	// Defer destroy to libc atexit: see the C preamble for rationale.
	// engine.destroy stays nil so engineAsLibrary.Scene's synchronous destroy
	// branch is skipped — the atexit handler owns teardown now.
	C.register_destroy_atexit(libgodot_destroy_godot_instance, ptr)
	engine.Library.Start()
}
