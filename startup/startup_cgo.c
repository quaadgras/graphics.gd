#include "../gdextension_interface.h"

// Little backwards compatibility entrypoint SHIM.

#ifdef _WIN32
#define EXPORT __declspec(dllexport)
#include <windows.h>

// On Windows, pin our own DLL in-process so the host cannot unload us while
// Go's background runtime threads (sysmon, GC workers, etc.) are still
// executing inside our .text. Godot's tool/export path calls FreeLibrary on
// extensions during shutdown, which races with sysmon's perpetual
// runtime.usleep -> stdcall -> asmstdcall loop and produces an
// EXCEPTION_ACCESS_VIOLATION_EXEC when the next instruction lands in the
// unmapped DLL (issue #303).
//
// GET_MODULE_HANDLE_EX_FLAG_PIN bumps the module reference until process
// exit, so FreeLibrary effectively becomes a no-op for our DLL while still
// letting Godot run its extension teardown callbacks. Cost is one extra
// module reference for the lifetime of the process.
static void pin_self_dll(void) {
    HMODULE self = NULL;
    GetModuleHandleExA(
        GET_MODULE_HANDLE_EX_FLAG_PIN | GET_MODULE_HANDLE_EX_FLAG_FROM_ADDRESS,
        (LPCSTR)(void *)&pin_self_dll,
        &self);
}
#else
#define EXPORT
#endif

extern GDExtensionBool gd_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization);

EXPORT GDExtensionBool cgo_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization) {
#ifdef _WIN32
    pin_self_dll();
#endif
    return gd_extension_init(p_get_proc_address, p_library, r_initialization);
}
