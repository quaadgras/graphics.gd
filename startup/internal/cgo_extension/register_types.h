#include "modules/register_module_types.h"

void initialize_cgo_extension_module(ModuleInitializationLevel p_level);
void uninitialize_cgo_extension_module(ModuleInitializationLevel p_level);

extern "C" {
    GDExtensionBool cgo_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization);
}
