/**************************************************************************/
/*  register_types.cpp                                                    */
/**************************************************************************/
/*                         This file is part of:                          */
/*                            BLAZIUM ENGINE                              */
/*                        https://blazium.app                             */
/**************************************************************************/
/* Copyright (c) 2024-present Blazium Engine contributors.                */
/* Copyright (c) 2024 Dragos Daian, Randolph William Aarseth II.          */
/*                                                                        */
/* Permission is hereby granted, free of charge, to any person obtaining  */
/* a copy of this software and associated documentation files (the        */
/* "Software"), to deal in the Software without restriction, including    */
/* without limitation the rights to use, copy, modify, merge, publish,    */
/* distribute, sublicense, and/or sell copies of the Software, and to     */
/* permit persons to whom the Software is furnished to do so, subject to  */
/* the following conditions:                                              */
/*                                                                        */
/* The above copyright notice and this permission notice shall be         */
/* included in all copies or substantial portions of the Software.        */
/*                                                                        */
/* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,        */
/* EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF     */
/* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. */
/* IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY   */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,   */
/* TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE      */
/* SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                 */
/**************************************************************************/

#include "register_types.h"
#include "core/object/class_db.h"
#include "core/extension/gdextension_interface.h"
#include "core/extension/gdextension_loader.h"
#include "core/extension/gdextension_manager.h"

#include <functional>

class GDExtensionStaticLibraryLoader : public GDExtensionLoader {
  friend class GDExtensionManager;
  friend class GDExtension;

private:
  GDExtensionInitializationFunction entry_funcptr = nullptr;
  String library_path;

public:
  void set_entry_funcptr(GDExtensionInitializationFunction p_entry_funcptr) {
    entry_funcptr = p_entry_funcptr;
  }
  virtual Error open_library(const String &p_path) override {
    library_path = p_path;
    return OK;
  }
  virtual Error
  initialize(GDExtensionInterfaceGetProcAddress p_get_proc_address,
             const Ref<GDExtension> &p_extension,
             GDExtensionInitialization *r_initialization) override {
    GDExtensionInitializationFunction initialization_function =
        (GDExtensionInitializationFunction)entry_funcptr;
    if (initialization_function == nullptr) {
      ERR_PRINT("GDExtension initialization function '" + library_path +
                "' is null.");
      return FAILED;
    }
    GDExtensionBool ret = initialization_function(
        p_get_proc_address, p_extension.ptr(), r_initialization);

    if (ret) {
      return OK;
    } else {
      ERR_PRINT("GDExtension initialization function '" + library_path +
                "' returned an error.");
      return FAILED;
    }
  }
  virtual void close_library() override {}
  virtual bool is_library_open() const override { return true; }
  virtual bool has_library_changed() const override { return false; }
  virtual bool library_exists() const override { return true; }
};

void initialize_cgo_extension_module(ModuleInitializationLevel p_level) {
    if (p_level != MODULE_INITIALIZATION_LEVEL_SERVERS || Engine::get_singleton()->is_project_manager_hint()) return;
	Ref<GDExtensionStaticLibraryLoader> loader;
    loader.instantiate();
    loader->set_entry_funcptr(&cgo_extension_init);
    GDExtensionManager::get_singleton()->load_extension_with_loader("go", loader);
}

void uninitialize_cgo_extension_module(ModuleInitializationLevel p_level) {}
