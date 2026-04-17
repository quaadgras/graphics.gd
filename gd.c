#ifdef __EMSCRIPTEN__
    #include "core/extension/gdextension_interface.h"
    #include <emscripten/bind.h>
    #include <emscripten/emscripten.h>
    #include <bit>
#else
    #include "gdextension_interface.h"
#endif
#include "gd.h"
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <stdio.h>

// emscripten can only expose std::string to JS, not char *
// we want emscripten to be able to deal with 64bit values without resorting to slow big ints.
#ifdef __EMSCRIPTEN__
    #define EXPORT EMSCRIPTEN_KEEPALIVE
    extern "C" {
#else
	#ifdef _WIN32
	#define EXPORT __declspec(dllexport)
	#else
	#define EXPORT
	#endif
#endif

#define VARIANT_ARG_GET(n) (struct Variant){n##_1, {n##_2, n##_3}}
#define VARIANT_ARG_PUT(v) v.tag, v.payload[0], v.payload[1]
#define PACKED_ARRAY_ARG_GET(n) (struct PackedArray){{n##_1, n##_2}}

// X-macro list: each entry is X(snake_name, GDExtensionInterfaceType)
// Declares variables and loads proc addresses from a single source of truth.
#define GD_PROC_LIST(X) \
    X(mem_alloc2, GDExtensionInterfaceMemAlloc2) \
    X(mem_realloc2, GDExtensionInterfaceMemRealloc2) \
    X(mem_free, GDExtensionInterfaceMemFree) \
    X(print_error, GDExtensionInterfacePrintError) \
    X(print_error_with_message, GDExtensionInterfacePrintErrorWithMessage) \
    X(print_warning, GDExtensionInterfacePrintWarning) \
    X(print_warning_with_message, GDExtensionInterfacePrintWarningWithMessage) \
    X(print_script_error, GDExtensionInterfacePrintScriptError) \
    X(print_script_error_with_message, GDExtensionInterfacePrintScriptErrorWithMessage) \
    X(get_native_struct_size, GDExtensionInterfaceGetNativeStructSize) \
    X(get_godot_version2, GDExtensionInterfaceGetGodotVersion2) \
    X(variant_new_copy, GDExtensionInterfaceVariantNewCopy) \
    X(variant_new_nil, GDExtensionInterfaceVariantNewNil) \
    X(variant_destroy, GDExtensionInterfaceVariantDestroy) \
    X(variant_call, GDExtensionInterfaceVariantCall) \
    X(variant_call_static, GDExtensionInterfaceVariantCallStatic) \
    X(variant_evaluate, GDExtensionInterfaceVariantEvaluate) \
    X(variant_set, GDExtensionInterfaceVariantSet) \
    X(variant_set_named, GDExtensionInterfaceVariantSetNamed) \
    X(variant_set_keyed, GDExtensionInterfaceVariantSetKeyed) \
    X(variant_set_indexed, GDExtensionInterfaceVariantSetIndexed) \
    X(variant_get, GDExtensionInterfaceVariantGet) \
    X(variant_get_named, GDExtensionInterfaceVariantGetNamed) \
    X(variant_get_keyed, GDExtensionInterfaceVariantGetKeyed) \
    X(variant_get_indexed, GDExtensionInterfaceVariantGetIndexed) \
    X(variant_iter_init, GDExtensionInterfaceVariantIterInit) \
    X(variant_iter_next, GDExtensionInterfaceVariantIterNext) \
    X(variant_iter_get, GDExtensionInterfaceVariantIterGet) \
    X(variant_hash, GDExtensionInterfaceVariantHash) \
    X(variant_recursive_hash, GDExtensionInterfaceVariantRecursiveHash) \
    X(variant_hash_compare, GDExtensionInterfaceVariantHashCompare) \
    X(variant_booleanize, GDExtensionInterfaceVariantBooleanize) \
    X(variant_duplicate, GDExtensionInterfaceVariantDuplicate) \
    X(variant_stringify, GDExtensionInterfaceVariantStringify) \
    X(variant_get_type, GDExtensionInterfaceVariantGetType) \
    X(variant_has_method, GDExtensionInterfaceVariantHasMethod) \
    X(variant_has_member, GDExtensionInterfaceVariantHasMember) \
    X(variant_has_key, GDExtensionInterfaceVariantHasKey) \
    X(variant_get_object_instance_id, GDExtensionInterfaceVariantGetObjectInstanceId) \
    X(variant_get_type_name, GDExtensionInterfaceVariantGetTypeName) \
    X(variant_can_convert, GDExtensionInterfaceVariantCanConvert) \
    X(variant_can_convert_strict, GDExtensionInterfaceVariantCanConvertStrict) \
    X(get_variant_from_type_constructor, GDExtensionInterfaceGetVariantFromTypeConstructor) \
    X(get_variant_to_type_constructor, GDExtensionInterfaceGetVariantToTypeConstructor) \
    X(variant_get_ptr_internal_getter, GDExtensionInterfaceGetVariantGetInternalPtrFunc) \
    X(variant_get_ptr_operator_evaluator, GDExtensionInterfaceVariantGetPtrOperatorEvaluator) \
    X(variant_get_ptr_builtin_method, GDExtensionInterfaceVariantGetPtrBuiltinMethod) \
    X(variant_get_ptr_constructor, GDExtensionInterfaceVariantGetPtrConstructor) \
    X(variant_get_ptr_destructor, GDExtensionInterfaceVariantGetPtrDestructor) \
    X(variant_construct, GDExtensionInterfaceVariantConstruct) \
    X(variant_get_ptr_setter, GDExtensionInterfaceVariantGetPtrSetter) \
    X(variant_get_ptr_getter, GDExtensionInterfaceVariantGetPtrGetter) \
    X(variant_get_ptr_indexed_setter, GDExtensionInterfaceVariantGetPtrIndexedSetter) \
    X(variant_get_ptr_indexed_getter, GDExtensionInterfaceVariantGetPtrIndexedGetter) \
    X(variant_get_ptr_keyed_setter, GDExtensionInterfaceVariantGetPtrKeyedSetter) \
    X(variant_get_ptr_keyed_getter, GDExtensionInterfaceVariantGetPtrKeyedGetter) \
    X(variant_get_ptr_keyed_checker, GDExtensionInterfaceVariantGetPtrKeyedChecker) \
    X(variant_get_constant_value, GDExtensionInterfaceVariantGetConstantValue) \
    X(variant_get_ptr_utility_function, GDExtensionInterfaceVariantGetPtrUtilityFunction) \
    X(string_new_with_latin1_chars, GDExtensionInterfaceStringNewWithLatin1Chars) \
    X(string_new_with_utf8_chars, GDExtensionInterfaceStringNewWithUtf8Chars) \
    X(string_new_with_utf16_chars, GDExtensionInterfaceStringNewWithUtf16Chars) \
    X(string_new_with_utf32_chars, GDExtensionInterfaceStringNewWithUtf32Chars) \
    X(string_new_with_wide_chars, GDExtensionInterfaceStringNewWithWideChars) \
    X(string_new_with_latin1_chars_and_len, GDExtensionInterfaceStringNewWithLatin1CharsAndLen) \
    X(string_new_with_utf8_chars_and_len, GDExtensionInterfaceStringNewWithUtf8CharsAndLen) \
    X(string_new_with_utf8_chars_and_len2, GDExtensionInterfaceStringNewWithUtf8CharsAndLen2) \
    X(string_new_with_utf16_chars_and_len, GDExtensionInterfaceStringNewWithUtf16CharsAndLen) \
    X(string_new_with_utf16_chars_and_len2, GDExtensionInterfaceStringNewWithUtf16CharsAndLen2) \
    X(string_new_with_utf32_chars_and_len, GDExtensionInterfaceStringNewWithUtf32CharsAndLen) \
    X(string_new_with_wide_chars_and_len, GDExtensionInterfaceStringNewWithWideCharsAndLen) \
    X(string_to_latin1_chars, GDExtensionInterfaceStringToLatin1Chars) \
    X(string_to_utf8_chars, GDExtensionInterfaceStringToUtf8Chars) \
    X(string_to_utf16_chars, GDExtensionInterfaceStringToUtf16Chars) \
    X(string_to_utf32_chars, GDExtensionInterfaceStringToUtf32Chars) \
    X(string_to_wide_chars, GDExtensionInterfaceStringToWideChars) \
    X(string_operator_index, GDExtensionInterfaceStringOperatorIndex) \
    X(string_operator_index_const, GDExtensionInterfaceStringOperatorIndexConst) \
    X(string_operator_plus_eq_string, GDExtensionInterfaceStringOperatorPlusEqString) \
    X(string_operator_plus_eq_char, GDExtensionInterfaceStringOperatorPlusEqChar) \
    X(string_operator_plus_eq_cstr, GDExtensionInterfaceStringOperatorPlusEqCstr) \
    X(string_operator_plus_eq_wcstr, GDExtensionInterfaceStringOperatorPlusEqWcstr) \
    X(string_operator_plus_eq_c32str, GDExtensionInterfaceStringOperatorPlusEqC32str) \
    X(string_resize, GDExtensionInterfaceStringResize) \
    X(string_name_new_with_latin1_chars, GDExtensionInterfaceStringNameNewWithLatin1Chars) \
    X(string_name_new_with_utf8_chars_and_len, GDExtensionInterfaceStringNameNewWithUtf8CharsAndLen) \
    X(xml_parser_open_buffer, GDExtensionInterfaceXmlParserOpenBuffer) \
    X(file_access_store_buffer, GDExtensionInterfaceFileAccessStoreBuffer) \
    X(file_access_get_buffer, GDExtensionInterfaceFileAccessGetBuffer) \
    X(worker_thread_pool_add_native_group_task, GDExtensionInterfaceWorkerThreadPoolAddNativeGroupTask) \
    X(worker_thread_pool_add_native_task, GDExtensionInterfaceWorkerThreadPoolAddNativeTask) \
    X(packed_byte_array_operator_index, GDExtensionInterfacePackedByteArrayOperatorIndex) \
    X(packed_byte_array_operator_index_const, GDExtensionInterfacePackedByteArrayOperatorIndexConst) \
    X(packed_color_array_operator_index, GDExtensionInterfacePackedColorArrayOperatorIndex) \
    X(packed_color_array_operator_index_const, GDExtensionInterfacePackedColorArrayOperatorIndexConst) \
    X(packed_float32_array_operator_index, GDExtensionInterfacePackedFloat32ArrayOperatorIndex) \
    X(packed_float32_array_operator_index_const, GDExtensionInterfacePackedFloat32ArrayOperatorIndexConst) \
    X(packed_float64_array_operator_index, GDExtensionInterfacePackedFloat64ArrayOperatorIndex) \
    X(packed_float64_array_operator_index_const, GDExtensionInterfacePackedFloat64ArrayOperatorIndexConst) \
    X(packed_int32_array_operator_index, GDExtensionInterfacePackedInt32ArrayOperatorIndex) \
    X(packed_int32_array_operator_index_const, GDExtensionInterfacePackedInt32ArrayOperatorIndexConst) \
    X(packed_int64_array_operator_index, GDExtensionInterfacePackedInt64ArrayOperatorIndex) \
    X(packed_int64_array_operator_index_const, GDExtensionInterfacePackedInt64ArrayOperatorIndexConst) \
    X(packed_string_array_operator_index, GDExtensionInterfacePackedStringArrayOperatorIndex) \
    X(packed_string_array_operator_index_const, GDExtensionInterfacePackedStringArrayOperatorIndexConst) \
    X(packed_vector2_array_operator_index, GDExtensionInterfacePackedVector2ArrayOperatorIndex) \
    X(packed_vector2_array_operator_index_const, GDExtensionInterfacePackedVector2ArrayOperatorIndexConst) \
    X(packed_vector3_array_operator_index, GDExtensionInterfacePackedVector3ArrayOperatorIndex) \
    X(packed_vector3_array_operator_index_const, GDExtensionInterfacePackedVector3ArrayOperatorIndexConst) \
    X(packed_vector4_array_operator_index, GDExtensionInterfacePackedVector4ArrayOperatorIndex) \
    X(packed_vector4_array_operator_index_const, GDExtensionInterfacePackedVector4ArrayOperatorIndexConst) \
    X(array_operator_index, GDExtensionInterfaceArrayOperatorIndex) \
    X(array_operator_index_const, GDExtensionInterfaceArrayOperatorIndexConst) \
    X(array_set_typed, GDExtensionInterfaceArraySetTyped) \
    X(dictionary_operator_index, GDExtensionInterfaceDictionaryOperatorIndex) \
    X(dictionary_operator_index_const, GDExtensionInterfaceDictionaryOperatorIndexConst) \
    X(dictionary_set_typed, GDExtensionInterfaceDictionarySetTyped) \
    X(object_method_bind_call, GDExtensionInterfaceObjectMethodBindCall) \
    X(object_method_bind_ptrcall, GDExtensionInterfaceObjectMethodBindPtrcall) \
    X(object_destroy, GDExtensionInterfaceObjectDestroy) \
    X(global_get_singleton, GDExtensionInterfaceGlobalGetSingleton) \
    X(object_get_instance_binding, GDExtensionInterfaceObjectGetInstanceBinding) \
    X(object_set_instance_binding, GDExtensionInterfaceObjectSetInstanceBinding) \
    X(object_free_instance_binding, GDExtensionInterfaceObjectFreeInstanceBinding) \
    X(object_set_instance, GDExtensionInterfaceObjectSetInstance) \
    X(object_get_class_name, GDExtensionInterfaceObjectGetClassName) \
    X(object_cast_to, GDExtensionInterfaceObjectCastTo) \
    X(object_get_instance_from_id, GDExtensionInterfaceObjectGetInstanceFromId) \
    X(object_get_instance_id, GDExtensionInterfaceObjectGetInstanceId) \
    X(object_has_script_method, GDExtensionInterfaceObjectHasScriptMethod) \
    X(object_call_script_method, GDExtensionInterfaceObjectCallScriptMethod) \
    X(callable_custom_create2, GDExtensionInterfaceCallableCustomCreate2) \
    X(callable_custom_get_userdata, GDExtensionInterfaceCallableCustomGetUserData) \
    X(ref_get_object, GDExtensionInterfaceRefGetObject) \
    X(ref_set_object, GDExtensionInterfaceRefSetObject) \
    X(script_instance_create3, GDExtensionInterfaceScriptInstanceCreate3) \
    X(placeholder_script_instance_create, GDExtensionInterfacePlaceHolderScriptInstanceCreate) \
    X(placeholder_script_instance_update, GDExtensionInterfacePlaceHolderScriptInstanceUpdate) \
    X(object_get_script_instance, GDExtensionInterfaceObjectGetScriptInstance) \
    X(object_set_script_instance, GDExtensionInterfaceObjectSetScriptInstance) \
    X(classdb_construct_object2, GDExtensionInterfaceClassdbConstructObject2) \
    X(classdb_get_method_bind, GDExtensionInterfaceClassdbGetMethodBind) \
    X(classdb_get_class_tag, GDExtensionInterfaceClassdbGetClassTag) \
    X(classdb_register_extension_class5, GDExtensionInterfaceClassdbRegisterExtensionClass5) \
    X(classdb_register_extension_class_method, GDExtensionInterfaceClassdbRegisterExtensionClassMethod) \
    X(classdb_register_extension_class_virtual_method, GDExtensionInterfaceClassdbRegisterExtensionClassVirtualMethod) \
    X(classdb_register_extension_class_integer_constant, GDExtensionInterfaceClassdbRegisterExtensionClassIntegerConstant) \
    X(classdb_register_extension_class_property, GDExtensionInterfaceClassdbRegisterExtensionClassProperty) \
    X(classdb_register_extension_class_property_indexed, GDExtensionInterfaceClassdbRegisterExtensionClassPropertyIndexed) \
    X(classdb_register_extension_class_property_group, GDExtensionInterfaceClassdbRegisterExtensionClassPropertyGroup) \
    X(classdb_register_extension_class_property_subgroup, GDExtensionInterfaceClassdbRegisterExtensionClassPropertySubgroup) \
    X(classdb_register_extension_class_signal, GDExtensionInterfaceClassdbRegisterExtensionClassSignal) \
    X(classdb_unregister_extension_class, GDExtensionInterfaceClassdbUnregisterExtensionClass) \
    X(get_library_path, GDExtensionInterfaceGetLibraryPath) \
    X(editor_add_plugin, GDExtensionInterfaceEditorAddPlugin) \
    X(editor_remove_plugin, GDExtensionInterfaceEditorRemovePlugin) \
    X(editor_register_get_classes_used_callback, GDExtensionInterfaceEditorRegisterGetClassesUsedCallback) \
    X(editor_help_load_xml_from_utf8_chars, GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8Chars) \
    X(editor_help_load_xml_from_utf8_chars_and_len, GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8CharsAndLen) \
    X(image_ptrw, GDExtensionInterfaceImagePtrw) \
    X(image_ptr, GDExtensionInterfaceImagePtr) \
    X(register_main_loop_callbacks, GDExtensionInterfaceRegisterMainLoopCallbacks)

#define DECLARE_PROC(name, type) type gdextension_##name = NULL;
GD_PROC_LIST(DECLARE_PROC)
#undef DECLARE_PROC

static GDExtensionClassLibraryPtr gd_library = NULL;
static GDExtensionGodotVersion2 gd_godot_version_cached = {};

typedef struct { uint64_t part[2]; } result_16;
typedef struct { uint64_t part[3]; } result_24;
typedef struct { uint64_t part[4]; } result_32;
typedef struct { uint64_t part[8]; } result_64;

GDExtensionTypeFromVariantConstructorFunc type_from_variant_constructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionVariantFromTypeConstructorFunc variant_from_type_constructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrDestructor variant_ptr_destructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionVariantGetInternalPtrFunc variant_internal_ptr_funcs[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrIndexedSetter variant_ptr_indexed_setters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrIndexedGetter variant_ptr_indexed_getters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrKeyedSetter variant_ptr_keyed_setters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrKeyedGetter variant_ptr_keyed_getters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];

EXPORT bool gd_extension_init(void* p_get_proc_address, void* p_library, void* r_initialization, const gd_extension* extension) {
    GDExtensionInterfaceGetProcAddress proc = (GDExtensionInterfaceGetProcAddress)p_get_proc_address;
    #pragma GCC diagnostic push
    #pragma GCC diagnostic ignored "-Wcast-function-type"
    #define LOAD_PROC(name, type) gdextension_##name = (type)proc(#name);
    GD_PROC_LIST(LOAD_PROC)
    #undef LOAD_PROC
    #pragma GCC diagnostic pop
    gd_library = (GDExtensionClassLibraryPtr)p_library;
    GDExtensionInitialization *r_init = (GDExtensionInitialization *)r_initialization;
    r_init->userdata = (void*)extension;
    r_init->minimum_initialization_level = GDEXTENSION_INITIALIZATION_CORE;
    r_init->initialize = (GDExtensionInitializeCallback)extension->on_engine_init;
    r_init->deinitialize = (GDExtensionInitializeCallback)extension->on_engine_exit;
    gdextension_get_godot_version2(&gd_godot_version_cached);
    GDExtensionMainLoopCallbacks callbacks = {
    	.startup_func = extension->on_first_frame,
    	.shutdown_func = extension->on_final_frame,
    	.frame_func = extension->on_every_frame,
    };
    gdextension_register_main_loop_callbacks((GDExtensionClassLibraryPtr)p_library, &callbacks);
    for (int i = 1; i < GDEXTENSION_VARIANT_TYPE_VARIANT_MAX; i++) {
    	GDExtensionVariantType v = (GDExtensionVariantType)i;
        variant_from_type_constructors[i] = gdextension_get_variant_from_type_constructor(v);
        type_from_variant_constructors[i] = gdextension_get_variant_to_type_constructor(v);
        variant_ptr_destructors[i] = gdextension_variant_get_ptr_destructor(v);
        variant_internal_ptr_funcs[i] = gdextension_variant_get_ptr_internal_getter(v);
        variant_ptr_indexed_setters[i] = gdextension_variant_get_ptr_indexed_setter(v);
        variant_ptr_indexed_getters[i] = gdextension_variant_get_ptr_indexed_getter(v);
        variant_ptr_keyed_setters[i] = gdextension_variant_get_ptr_keyed_setter(v);
        variant_ptr_keyed_getters[i] = gdextension_variant_get_ptr_keyed_getter(v);
    }
    return true;
}
static void prepare_variants(void **frame, uint32_t argc, gd_addr args) {
    uint8_t *head = (uint8_t*)args;
    for (int i = 0; i < argc; i++) {
        frame[i] = head;
        head += 24;
    }
}
// Helper macro to align a value to the next multiple of 'align'
#define ALIGN_UP(value, align) (((value) + ((align) - 1)) & ~((align) - 1))
static uint8_t prepare_callframe(int skip, void **frame, uint64_t shape, gd_addr args) {
    uint8_t *head = (uint8_t *)args;
    ptrdiff_t offset = 0; // Track current offset in the frame
    for (int i = skip; i < 16; i++) {
        gd_shape code = (gd_shape)((shape >> (i * 4)) & 0xF);
        uint32_t size;
        uint32_t align;
        // Determine size based on code
        switch (code) {
            case 0: size = 0; frame[i-skip] = NULL; return i-skip;
            case GD_BYTES_1x1: size = 1; align = 1; break;
            case GD_BYTES_2x1: size = 2; align = 2; break;
            case GD_BYTES_4x1: size = 4; align = 4; break;
            case GD_BYTES_8x1: size = 8; align = 8; break;
            case GD_BYTES_4x2: size = 4*2; align = 4; break;
            case GD_BYTES_4x3: size = 4*3; align = 4; break;
            case GD_BYTES_8x2: size = 8*2; align = 8; break;
            case GD_BYTES_4x4: size = 4*4; align = 4; break;
            case GD_BYTES_8x3: size = 8*3; align = 8; break;
            case GD_BYTES_4x6: size = 4*6; align = 4; break;
            case GD_BYTES_4x9: size = 4*9; align = 4; break;
            case GD_BYTES_4x12: size = 4*12; align = 4; break;
            case GD_BYTES_4x16: size = 4*16; align = 4; break;
        }
        offset = ALIGN_UP(offset, align);
        frame[i-skip] = head + offset;     // Set frame pointer to the aligned address
        offset += size;                 // Move offset forward by the size of the current argument
    }
    return 16-skip;
}

//
// Version Information
//

EXPORT struct String gd_version() { struct String s; gdextension_string_new_with_latin1_chars(&s, gd_godot_version_cached.string); return s; };
EXPORT int64_t    gd_version_major() { return gd_godot_version_cached.major; };
EXPORT int64_t    gd_version_minor() { return gd_godot_version_cached.minor; };
EXPORT int64_t    gd_version_patch() { return gd_godot_version_cached.patch; };
EXPORT int64_t    gd_version_hexed() { return gd_godot_version_cached.hex; };
EXPORT struct String gd_version_state() { struct String s; gdextension_string_new_with_latin1_chars(&s, gd_godot_version_cached.status); return s; };
EXPORT struct String gd_version_build() { struct String s; gdextension_string_new_with_latin1_chars(&s, gd_godot_version_cached.build); return s; };
EXPORT struct String gd_version_stamp() { struct String s; gdextension_string_new_with_latin1_chars(&s, gd_godot_version_cached.hash); return s; };
EXPORT int64_t    gd_version_nanos() { return gd_godot_version_cached.timestamp * 1000000000; };

//
// Engine Memory Access
//

EXPORT int64_t    gd_sizeof(struct StringName name) { return gdextension_get_native_struct_size((GDExtensionConstStringNamePtr)&name); };
EXPORT gd_addr gd_malloc(int64_t size, bool pad8) { return (gd_addr)gdextension_mem_alloc2(size, pad8); };
EXPORT void*   gd_resize(gd_addr addr, int64_t size, bool pad8) { return (void*)gdextension_mem_realloc2((void *)addr, size, pad8); };
EXPORT void    gd_memset(gd_addr addr, uint8_t value, int64_t size) { if (size <= 0) return; for (int i = 0; i < size; i++) ((uint8_t *)addr)[i] = value; };
EXPORT uint8_t  gd_memory_bytes1(const gd_addr addr) { return *(uint8_t*)addr; };
EXPORT uint16_t gd_memory_bytes2(const gd_addr addr) { return *(uint16_t*)addr; };
EXPORT uint32_t gd_memory_bytes4(const gd_addr addr) { return *(uint32_t*)addr; };
EXPORT uint64_t gd_memory_bytes8(const gd_addr addr) { return *(uint64_t*)addr; };
EXPORT void   gd_store_bytes1(gd_addr addr, uint8_t b) { *(uint8_t *)addr = b; };
EXPORT void   gd_store_bytes2(gd_addr addr, uint16_t b) { *(uint16_t *)addr = b; };
EXPORT void   gd_store_bytes4(gd_addr addr, uint32_t b) { *(uint32_t *)addr = b; };
EXPORT void   gd_store_bytes8(gd_addr addr, uint64_t b) { *(uint64_t *)addr = b; };
EXPORT void   gd_store_pair64(gd_addr addr, uint64_t a, uint64_t b) { uint64_t *ptr = (uint64_t *)addr; ptr[0] = a; ptr[1] = b; };
EXPORT void   gd_store_quad64(gd_addr addr, uint64_t a, uint64_t b, uint64_t c, uint64_t d) { uint64_t *ptr = (uint64_t *)addr; ptr[0] = a; ptr[1] = b; ptr[2] = c; ptr[3] = d; };
EXPORT void   gd_store_octo64(gd_addr addr, uint64_t a, uint64_t b, uint64_t c, uint64_t d, uint64_t e, uint64_t f, uint64_t g, uint64_t h) { uint64_t *ptr = (uint64_t *)addr; ptr[0] = a; ptr[1] = b; ptr[2] = c; ptr[3] = d; ptr[4] = e; ptr[5] = f; ptr[6] = g; ptr[7] = h; };
EXPORT void   gd_free(gd_addr addr) { gdextension_mem_free((void *)addr);};

EXPORT void gd_callable_create(gd_extension* extension, gd_extension_callable_t id, ObjectID object, struct Callable* result) {
    GDExtensionCallableCustomInfo2 info = {
        .callable_userdata = (void*)id,
        .token = gd_library,
        .object_id = object,
        .call_func = (GDExtensionCallableCustomCall)extension->on_callable_called,
        .is_valid_func = (GDExtensionCallableCustomIsValid)extension->on_callable_verify,
        .free_func = (GDExtensionCallableCustomFree)extension->on_callable_delete,
        .hash_func = (GDExtensionCallableCustomHash)extension->on_callable_hashed,
        .equal_func = (GDExtensionCallableCustomEqual)extension->on_callable_equal,
        .less_than_func = (GDExtensionCallableCustomLessThan)extension->on_callable_less_than,
        .to_string_func = (GDExtensionCallableCustomToString)extension->on_callable_string,
        .get_argument_count_func = (GDExtensionCallableCustomGetArgumentCount)extension->on_callable_length,
    };
    gdextension_callable_custom_create2(result, &info);
}
EXPORT gd_extension_callable_t gd_callable_lookup(CALLABLE_ARG(c)) {
    struct Callable callable = {.opaque = {c_1, c_2}};
    return (gd_extension_callable_t)gdextension_callable_custom_get_userdata(&callable, gd_library);
};

EXPORT gd_function_t gd_function(struct StringName utility, int64_t hash) { return (gd_function_t)gdextension_variant_get_ptr_utility_function((GDExtensionConstStringNamePtr)&utility, hash);}
EXPORT void gd_call(gd_function_t fn, gd_addr result, gd_shape shape, gd_addr args) {
    void *points[16]; uint8_t argc = prepare_callframe(1, &points[0], shape, args);
    ((GDExtensionPtrUtilityFunction)fn)((GDExtensionTypePtr)result, (GDExtensionConstTypePtr*)&points[0], argc);
}
EXPORT struct String gd_extension_library_location() {
    struct String s;
    gdextension_get_library_path(gd_library, &s);
    return s;
}
EXPORT void gd_classdb_FileAccess_write(struct Object FileAccess, char* buf, int64_t len) {
    gdextension_file_access_store_buffer((GDExtensionObjectPtr)(FileAccess.opaque), (const uint8_t *)buf, len);
};
EXPORT int64_t gd_classdb_FileAccess_read(struct Object FileAccess, char* buf, int64_t len) {
    return gdextension_file_access_get_buffer((GDExtensionObjectPtr)(FileAccess.opaque), (uint8_t *)buf, len);
};
EXPORT gd_addr gd_classdb_Image_memory(struct Object Image) {
    return (gd_addr)gdextension_image_ptrw((GDExtensionObjectPtr)(Image.opaque));
};
EXPORT uint8_t gd_classdb_Image_access(struct Object Image, int64_t offset) {
    return gdextension_image_ptr((GDExtensionObjectPtr)(Image.opaque))[offset];
};
typedef struct {
    int32_t push;
    int32_t size;
    gd_extension* extension;
    GDExtensionClassMethodInfo *info;
} method_list;
typedef struct {
    int32_t push;
    int32_t size;
    GDExtensionPropertyInfo *info;
    GDExtensionClassMethodArgumentMetadata *meta;
} property_list;
EXPORT gd_method_list_t gd_method_list_make(gd_extension* extension, int64_t length) {
    method_list *list = (method_list *)gdextension_mem_alloc2(sizeof(method_list), false);
    list->push = 0;
    list->size = length;
    list->extension = extension;
    list->info = (GDExtensionClassMethodInfo*)gdextension_mem_alloc2(sizeof(GDExtensionClassMethodInfo) * length, false);
    return (gd_method_list_t)list;
};
EXPORT void gd_method_list_push(gd_method_list_t list_p, struct StringName name, gd_extension_method_id call, MethodFlags method_flags, gd_property_list_t return_value_info, gd_property_list_t arguments_info, int64_t default_argument_count, gd_addr default_arguments) {
    method_list *list = (method_list *)list_p;
    if (list->push >= list->size) return;
    GDExtensionClassMethodInfo *info = &list->info[list->push++];
    property_list *return_value = (property_list *)return_value_info;
    property_list *arguments = (property_list *)arguments_info;
    uintptr_t *name_allocated = (uintptr_t *)gdextension_mem_alloc2(sizeof(uintptr_t), false);
    *name_allocated = name.opaque;
    info->name = (GDExtensionStringNamePtr)name_allocated;
    info->method_userdata = (void *)call;
    info->call_func = (GDExtensionClassMethodCall)list->extension->on_extension_instance_dynamic_call;
    info->ptrcall_func = (GDExtensionClassMethodPtrCall)list->extension->on_extension_instance_checked_call;
    info->method_flags = method_flags;
    if (return_value && return_value->push > 0) {
        info->has_return_value = true;
        info->return_value_info = return_value->info;
        info->return_value_metadata = *return_value->meta;
    } else {
        info->has_return_value = false;
        info->return_value_info = NULL;
        info->return_value_metadata = (GDExtensionClassMethodArgumentMetadata)0;
    }
    if (arguments && arguments->push > 0) {
        info->argument_count = arguments->push;
        info->arguments_info = arguments->info;
        info->arguments_metadata = arguments->meta;
    } else {
        info->argument_count = 0;
        info->arguments_info = NULL;
        info->arguments_metadata = NULL;
    }
    void **points = (void **)gdextension_mem_alloc2(sizeof(void*) * default_argument_count, false);
    prepare_variants(&points[0], default_argument_count, default_arguments);
    info->default_argument_count = default_argument_count;
    info->default_arguments = points;
};
EXPORT void gd_method_list_free(gd_method_list_t list_p) {
    method_list *list = (method_list *)list_p;
    for (int i = 0; i < list->push; i++) {
        gdextension_mem_free(list->info[i].name);
        if (list->info[i].default_arguments) {
            gdextension_mem_free(list->info[i].default_arguments);
        }
    }
    gdextension_mem_free(list->info); gdextension_mem_free(list);
};
EXPORT gd_property_list_t gd_property_list_make(int64_t length) {
    property_list *list = (property_list*)gdextension_mem_alloc2(sizeof(property_list), false);
    list->push = 0;
    list->size = length;
    list->info = (GDExtensionPropertyInfo*)gdextension_mem_alloc2(sizeof(GDExtensionPropertyInfo) * length, false);
    list->meta = (GDExtensionClassMethodArgumentMetadata*)gdextension_mem_alloc2(sizeof(GDExtensionClassMethodArgumentMetadata) * length, false);
    return (gd_property_list_t)list;
};
EXPORT void gd_property_list_push(gd_property_list_t list_p,
    VariantType vtype, struct StringName name, struct StringName class_name,
    uint32_t hint, struct String hint_string, uint32_t usage, ArgumentMetadata meta
) {
    property_list *list = (property_list *)list_p;
    if (list->push >= list->size) return;
    GDExtensionPropertyInfo *info = &list->info[list->push++];
    GDExtensionClassMethodArgumentMetadata *meta_info = &list->meta[list->push - 1];
    uintptr_t *name_allocated = (uintptr_t *)gdextension_mem_alloc2(sizeof(uintptr_t), false);
    *name_allocated = name.opaque;
    uintptr_t *class_name_allocated = (uintptr_t *)gdextension_mem_alloc2(sizeof(uintptr_t), false);
    *class_name_allocated = class_name.opaque;
    uintptr_t *hint_string_allocated = (uintptr_t *)gdextension_mem_alloc2(sizeof(uintptr_t), false);
    *hint_string_allocated = hint_string.ptr;
    info->type = (GDExtensionVariantType)vtype;
    info->name = (GDExtensionStringNamePtr)name_allocated;
    info->class_name = (GDExtensionStringNamePtr)class_name_allocated;
    info->hint = hint;
    info->hint_string = (GDExtensionStringPtr)hint_string_allocated;
    info->usage = usage;
    *meta_info = (GDExtensionClassMethodArgumentMetadata)meta;
};
EXPORT void gd_property_list_free(gd_property_list_t list_p) {
    property_list *list = (property_list *)list_p;
    for (int i = 0; i < list->push; i++) {
        gdextension_mem_free(list->info[i].name);
        gdextension_mem_free(list->info[i].class_name);
        gdextension_mem_free(list->info[i].hint_string);
    }
    gdextension_mem_free(list->info);
    gdextension_mem_free(list->meta);
    gdextension_mem_free(list);
};
EXPORT uint32_t gd_property_info_type(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return list->info[list->push-1].type;
};
EXPORT uintptr_t gd_property_info_name(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return (uintptr_t)list->info[list->push-1].name;
};
EXPORT uintptr_t gd_property_info_class_name(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return (uintptr_t)list->info[list->push-1].class_name;
};
EXPORT uint32_t gd_property_info_hint(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return list->info[list->push-1].hint;
};
EXPORT uintptr_t gd_property_info_hint_string(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return (uintptr_t)list->info[list->push-1].hint_string;
};
EXPORT uint32_t gd_property_info_usage(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return list->info[list->push-1].usage;
};
static void free_property_list(void* instance, const GDExtensionPropertyInfo* list, uint32_t count) {
    for (uint32_t i = 0; i < count; i++) {
        gdextension_mem_free((void*)list[i].name);
        gdextension_mem_free((void*)list[i].class_name);
        gdextension_mem_free((void*)list[i].hint_string);
    }
    gdextension_mem_free((void*)list);
}

EXPORT void gd_classdb_register(gd_extension* extension, struct StringName class_name, struct StringName parent, gd_extension_class_id id, bool is_virtual, bool abstract, bool exposed, bool runtime, struct String icon_path) {
    GDExtensionClassCreationInfo5 info = {
        .is_virtual = is_virtual,
        .is_abstract = abstract,
        .is_exposed = exposed,
        .is_runtime = runtime,
        .icon_path = (GDExtensionConstStringPtr)&icon_path,
        .set_func = extension->on_extension_instance_set,
        .get_func = extension->on_extension_instance_get,
        .get_property_list_func = (GDExtensionClassGetPropertyList)extension->on_extension_instance_property_list,
        .free_property_list_func = free_property_list,
        .property_can_revert_func = extension->on_extension_instance_property_has_default,
        .property_get_revert_func = extension->on_extension_instance_property_get_default,
        .validate_property_func = (GDExtensionClassValidateProperty)extension->on_extension_instance_property_validation,
        .notification_func = extension->on_extension_instance_notification,
        .to_string_func = extension->on_extension_instance_stringify,
        .reference_func = extension->on_extension_instance_reference,
        .unreference_func = extension->on_extension_instance_unreference,
        .create_instance_func = extension->on_extension_class_create,
        .free_instance_func = extension->on_extension_instance_free,
        .get_virtual_call_data_func = extension->on_extension_class_caller,
        .call_virtual_with_data_func = extension->on_extension_instance_called,
        .class_userdata = (void *)id,
    };
    gdextension_classdb_register_extension_class5(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&parent, &info);
};
EXPORT void gd_classdb_register_methods(struct StringName class_name, gd_method_list_t methods) {
    method_list *list = (method_list *)methods;
    for (int i = 0; i < list->push; i++) {
        GDExtensionClassMethodInfo *info = &list->info[i];
        gdextension_classdb_register_extension_class_method(gd_library, (GDExtensionConstStringNamePtr)&class_name, info);
    }
};
EXPORT void gd_classdb_register_constant(struct StringName class_name, struct StringName enum_name, struct StringName name, int64_t value, bool bitfield) {
    gdextension_classdb_register_extension_class_integer_constant(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&enum_name, (GDExtensionConstStringNamePtr)&name, value, bitfield);
};
EXPORT void gd_classdb_register_property(struct StringName class_name, gd_property_list_t info, struct StringName setter, struct StringName getter) {
    property_list *list = (property_list *)info;
    gdextension_classdb_register_extension_class_property(gd_library, (GDExtensionConstStringNamePtr)&class_name, list->info, (GDExtensionConstStringNamePtr)&setter, (GDExtensionConstStringNamePtr)&getter);
};
EXPORT void gd_classdb_register_property_indexed(struct StringName class_name, gd_property_list_t info, struct StringName setter, struct StringName getter, int64_t index) {
    property_list *list = (property_list *)info;
    gdextension_classdb_register_extension_class_property_indexed(gd_library, (GDExtensionConstStringNamePtr)&class_name, list->info, (GDExtensionConstStringNamePtr)&setter, (GDExtensionConstStringNamePtr)&getter, index);
};
EXPORT void gd_classdb_register_property_group(struct StringName class_name, struct String group, struct String prefix) {
    gdextension_classdb_register_extension_class_property_group(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringPtr)&group, (GDExtensionConstStringPtr)&prefix);
};
EXPORT void gd_classdb_register_property_sub_group(struct StringName class_name, struct String subgroup, struct String prefix) {
    gdextension_classdb_register_extension_class_property_subgroup(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringPtr)&subgroup, (GDExtensionConstStringPtr)&prefix);
};
EXPORT void gd_classdb_register_signal(struct StringName class_name, struct StringName name, gd_property_list_t args) {
    property_list *list = (property_list *)args;
    gdextension_classdb_register_extension_class_signal(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&name, list->info, list->push);
};
EXPORT void gd_classdb_register_removal(struct StringName class_name) {
    gdextension_classdb_unregister_extension_class(gd_library, (GDExtensionConstStringNamePtr)&class_name);
};
EXPORT void gd_classdb_WorkerThreadPool_add_task(gd_extension* extension, struct Object pool, gd_extension_task_id task, bool priority, struct String description) {
    gdextension_worker_thread_pool_add_native_task((GDExtensionObjectPtr)(pool.opaque), (GDExtensionWorkerThreadPoolTask)extension->on_worker_thread_pool_task, (void *)task, priority, (GDExtensionConstStringPtr)&description);
};
EXPORT void gd_classdb_WorkerThreadPool_add_group_task(gd_extension* extension, struct Object pool, gd_extension_task_id task, int32_t elements, int32_t tasks, bool priority, struct String description) {
    gdextension_worker_thread_pool_add_native_group_task((GDExtensionObjectPtr)(pool.opaque), (GDExtensionWorkerThreadPoolGroupTask)extension->on_worker_thread_pool_group_task, (void *)task, elements, tasks, priority, (GDExtensionConstStringPtr)&description);
};
EXPORT int64_t gd_classdb_XMLParser_load(struct Object parser, char* buf, int64_t len) {
    return gdextension_xml_parser_open_buffer((GDExtensionObjectPtr)(parser.opaque), (const uint8_t *)buf, len);
};
EXPORT void gd_packed_dictionary_access(struct Dictionary dict, VARIANT_ARG(k), struct Variant* result) {
    struct Variant key = VARIANT_ARG_GET(k);
    struct Variant* value = (struct Variant*)gdextension_dictionary_operator_index_const((GDExtensionTypePtr)&dict, &key);
    if (!value) return;
    *result = *value;
};
EXPORT void gd_packed_dictionary_modify(struct Dictionary dict, VARIANT_ARG(k), VARIANT_ARG(v)) {
    struct Variant key = VARIANT_ARG_GET(k);
    struct Variant *value = (struct Variant*)gdextension_dictionary_operator_index((GDExtensionTypePtr)&dict, (GDExtensionVariantPtr)&key);
    *value = VARIANT_ARG_GET(v);
};
EXPORT void gd_editor_add_documentation(const char* xml, uint32_t len) {
    gdextension_editor_help_load_xml_from_utf8_chars_and_len(xml, len);
};
EXPORT void gd_editor_add_plugin(struct StringName class_name) {
    gdextension_editor_add_plugin((GDExtensionConstStringNamePtr)&class_name);
};
EXPORT void gd_editor_end_plugin(struct StringName class_name) {
    gdextension_editor_remove_plugin((GDExtensionConstStringNamePtr)&class_name);
};
EXPORT void gd_iterator_make(VARIANT_ARG(v), struct Variant* result_iter, gd_error* err) {
    struct Variant self = VARIANT_ARG_GET(v);
    GDExtensionBool valid = true;
    valid = valid && gdextension_variant_iter_init(&self, (void*)result_iter, &valid);
    if (valid) return;
    ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
};
EXPORT bool gd_iterator_next(VARIANT_ARG(v), struct Variant* iter, gd_error* err) {
    struct Variant self = VARIANT_ARG_GET(v);
    GDExtensionBool valid = false;
    GDExtensionBool ok = gdextension_variant_iter_next(&self, (GDExtensionVariantPtr)iter, &valid);
    if (ok) {
        return true;
    }
    ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return false;
};
EXPORT void gd_iterator_load(VARIANT_ARG(v), VARIANT_ARG(i), struct Variant* result, gd_error* err) {
    struct Variant self = VARIANT_ARG_GET(v);
    uint64_t iter[3] = {i_1, i_2, i_3};
    void *points[16]; prepare_variants(&points[0], 1, (gd_addr)result);
    GDExtensionBool ok = false;
    gdextension_variant_iter_get(&self, &iter, (GDExtensionUninitializedVariantPtr)points[0], &ok);
    if (ok) return;
    ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
};
static const char *fit_string(const char *str, uint64_t len, char *buf, size_t buf_size) {
    if (len == 0 || str == NULL) {
        return NULL;
    }
    if (len >= buf_size) {
        buf = (char *)gdextension_mem_alloc2(len + 1, false); // +1 for null-terminator
    }
    memcpy(buf, str, len);
    buf[len] = '\0'; // null-terminate the string
    return buf;
}
EXPORT void gd_log(gd_log_level level,
    const char* text, uint32_t text_len,
    const char* code, uint32_t code_len,
    const char* func, uint32_t func_len,
    const char* file, uint32_t file_len,
    int32_t line, bool notify_editor
) {
    char text_buf[256]; const char *text_ptr = fit_string(text, text_len, &text_buf[0], 256);
    char code_buf[100]; const char *code_ptr = fit_string(code, code_len, &code_buf[0], 100);
    char func_buf[100]; const char *func_ptr = fit_string(func, func_len, &func_buf[0], 100);
    char file_buf[100]; const char *file_ptr = fit_string(file, file_len, &file_buf[0], 100);
    switch (level) {
    case 0: gdextension_print_error_with_message(code_ptr, text_ptr, func_ptr, file_ptr, line, notify_editor); break;
    case 1: gdextension_print_warning_with_message(code_ptr, text_ptr, func_ptr, file_ptr, line, notify_editor); break;
    }
    if (text_ptr && text_ptr != text_buf) gdextension_mem_free((void *)text_ptr);
    if (code_ptr && code_ptr != code_buf) gdextension_mem_free((void *)code_ptr);
    if (func_ptr && func_ptr != func_buf) gdextension_mem_free((void *)func_ptr);
    if (file_ptr && file_ptr != file_buf) gdextension_mem_free((void *)file_ptr);
};
EXPORT struct Object gd_object_make(struct StringName name) {
    return (struct Object){(uintptr_t)gdextension_classdb_construct_object2((GDExtensionConstStringNamePtr)&name)};
};
EXPORT void gd_object_call(struct Object obj, gd_method_id fn, struct Variant* result, int64_t argc, struct Variant args[], gd_error* err) {
    void *points[16]; prepare_variants(&points[0], (uint32_t)argc, (gd_addr)args);
    gdextension_object_method_bind_call((GDExtensionMethodBindPtr)fn, (GDExtensionObjectPtr)(obj.opaque), (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};
EXPORT result_24 gd_object_call_24(uintptr_t obj, uintptr_t fn, int64_t argc, gd_addr args, gd_addr err) {
    void *points[16]; prepare_variants(&points[0], argc, args);
    result_24 result = {};
    gdextension_object_method_bind_call((GDExtensionMethodBindPtr)fn, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)&result, (GDExtensionCallError*)err);
    return result;
};
EXPORT void gd_object_free(struct Object obj) {
    gdextension_object_destroy((GDExtensionObjectPtr)(obj.opaque));
};
EXPORT struct StringName gd_object_name(struct Object obj) {
    struct StringName name = {0};
    gdextension_object_get_class_name((GDExtensionObjectPtr)(obj.opaque), gd_library, (GDExtensionUninitializedStringNamePtr)&name);
    return name;
};
EXPORT struct ClassTag gd_object_type(struct StringName name) {
    return (struct ClassTag){(uintptr_t)gdextension_classdb_get_class_tag((GDExtensionConstStringNamePtr)&name)};
};
EXPORT struct Object gd_object_cast(struct Object obj, struct ClassTag to) {
    return (struct Object){(uintptr_t)gdextension_object_cast_to((GDExtensionObjectPtr)(obj.opaque), (void *)(to.opaque))};
};
EXPORT struct Object gd_object_lookup(ObjectID id) {
    return (struct Object){(uintptr_t)gdextension_object_get_instance_from_id((GDObjectInstanceID)id)};
};
EXPORT struct Object gd_object_global(struct StringName name) {
    return (struct Object){(uintptr_t)gdextension_global_get_singleton((GDExtensionConstStringNamePtr)&name)};
};
EXPORT void gd_extension_object_setup(struct Object obj, struct StringName name, gd_extension_object_id instance) {
    gdextension_object_set_instance((GDExtensionObjectPtr)(obj.opaque), (GDExtensionConstStringNamePtr)&name, (GDExtensionClassInstancePtr)instance);
};
EXPORT gd_extension_binding_id gd_object_lookup_extension_binding(gd_extension* extension, struct Object obj) {
    GDExtensionInstanceBindingCallbacks instance_binding_callbacks = {
        .create_callback = extension->on_extension_binding_created,
        .reference_callback = extension->on_extension_binding_reference,
        .free_callback = extension->on_extension_binding_removed,
    };
    return (uintptr_t)gdextension_object_get_instance_binding((GDExtensionObjectPtr)(obj.opaque), gd_library, &instance_binding_callbacks);
};
EXPORT void gd_object_attach_extension_binding(gd_extension* extension, struct Object obj, gd_extension_binding_id binding) {
    GDExtensionInstanceBindingCallbacks instance_binding_callbacks = {
        .create_callback = extension->on_extension_binding_created,
        .reference_callback = extension->on_extension_binding_reference,
        .free_callback = extension->on_extension_binding_removed,
    };
    gdextension_object_set_instance_binding((GDExtensionObjectPtr)(obj.opaque), gd_library, (void *)binding, &instance_binding_callbacks);
};
EXPORT void gd_object_detach_extension_binding(struct Object obj) {
    gdextension_object_free_instance_binding((GDExtensionObjectPtr)(obj.opaque), gd_library);
};
EXPORT ObjectID gd_object_id(struct Object obj) {
    return gdextension_object_get_instance_id((GDExtensionObjectPtr)(obj.opaque));
};
EXPORT ObjectID gd_object_id_inside_variant(VARIANT_ARG(v)) {
    struct Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_get_object_instance_id(&self);
};
EXPORT gd_method_id gd_method(struct StringName class_name, struct StringName method, int64_t hash) {
    return (gd_method_id)gdextension_classdb_get_method_bind((GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&method, hash);
};
EXPORT void gd_method_call(struct Object obj, gd_method_id fn, gd_addr result, gd_shape shape, const gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)fn, (GDExtensionObjectPtr)(obj.opaque), (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)result);
};
typedef struct {
    uintptr_t object;
    uintptr_t method;
    uint64_t  shape;
    uint8_t   args[256];
    uint8_t   result[64];
    uint16_t  refs[16];
    uintptr_t owner;
    uintptr_t pc;
} ring_entry;
EXPORT void gd_ring_flush(void *entries, uint32_t tail, uint32_t head, uint32_t *crash_index) {
    ring_entry *ring = (ring_entry *)entries;
    for (uint32_t i = tail; i != head; i++) {
        ring_entry *e = &ring[i & 0xFF];
        *crash_index = i & 0xFF;
        void *points[16];
        prepare_callframe(1, &points[0], e->shape, (gd_addr)e->args);
        gdextension_object_method_bind_ptrcall(
            (GDExtensionMethodBindPtr)e->method,
            (GDExtensionObjectPtr)e->object,
            (const GDExtensionConstTypePtr*)&points[0],
            (GDExtensionTypePtr)e->result
        );
    }
    *crash_index = 0xFFFFFFFF;
}
EXPORT uint64_t gd_object_unsafe_call_8(uintptr_t obj, uintptr_t method, uint64_t shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    uint64_t result = 0;
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)&result);
    return result;
};
EXPORT result_16 gd_object_unsafe_call_16(uintptr_t obj, uintptr_t method, uint64_t shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    result_16 result = {};
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)&result);
    return result;
};
EXPORT result_32 gd_object_unsafe_call_32(uintptr_t obj, uintptr_t method, uint64_t shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    result_32 result = {};
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)&result);
    return result;
};
EXPORT result_64 gd_object_unsafe_call_64(uintptr_t obj, uintptr_t method, uint64_t shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    result_64 result = {};
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)&result);
    return result;
};
EXPORT uintptr_t gd_ptrcall_fn_addr() {
    return (uintptr_t)gdextension_object_method_bind_ptrcall;
}
EXPORT void gd_script_yield_property(gd_property_iterator_t fn, uintptr_t arg, struct StringName name, VARIANT_ARG(state)) {
    uint64_t v[3] = {state_1, state_2, state_3};
    ((GDExtensionScriptInstancePropertyStateAdd)fn)(&name, &v, (void *)arg);
}
struct Script gd_script_make(gd_extension* extension, gd_extension_script_id script) {
    GDExtensionScriptInstanceInfo3 info = {
        .set_func = extension->on_extension_instance_set,
        .get_func = extension->on_extension_instance_get,
        .get_property_list_func = (GDExtensionScriptInstanceGetPropertyList)extension->on_extension_instance_property_list,
        .free_property_list_func = free_property_list,
        .get_class_category_func = (GDExtensionScriptInstanceGetClassCategory)extension->on_extension_script_categorization,
        .property_can_revert_func = extension->on_extension_instance_property_has_default,
        .property_get_revert_func = extension->on_extension_instance_property_get_default,
        .get_owner_func = extension->on_extension_script_get_owner,
        .get_property_state_func = (GDExtensionScriptInstanceGetPropertyState)extension->on_extension_script_property_iter,
        .validate_property_func = (GDExtensionScriptInstanceValidateProperty)extension->on_extension_instance_property_validation,
        .has_method_func = extension->on_extension_script_has_method,
        .get_method_argument_count_func = extension->on_extension_script_get_method_argument_count,
        .call_func = (GDExtensionScriptInstanceCall)extension->on_extension_instance_dynamic_call,
        .notification_func = extension->on_extension_instance_notification,
        .to_string_func = extension->on_extension_instance_stringify,
        .refcount_incremented_func = extension->on_extension_instance_reference,
        .refcount_decremented_func = extension->on_extension_script_unreference,
        .get_script_func = extension->on_extension_script_get,
        .is_placeholder_func = extension->on_extension_script_is_placeholder,
        .get_language_func = extension->on_extension_script_get_language,
        .free_func = extension->on_extension_script_free,
    };
    return (struct Script){ (uintptr_t)gdextension_script_instance_create3(&info, (GDExtensionScriptInstanceDataPtr)script) };
};
EXPORT void gd_script_call(struct Object obj, struct StringName method, struct Variant* result, int64_t argc, struct Variant args[], gd_error* err) {
    void *points[16]; points[1] = 0; prepare_variants(&points[0], (uint32_t)argc, (gd_addr)args);
    gdextension_object_call_script_method((GDExtensionObjectPtr)(obj.opaque), (GDExtensionConstStringNamePtr)&method, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};
EXPORT void gd_script_setup(struct Object obj, gd_extension_script_id script) {
    gdextension_object_set_script_instance((GDExtensionObjectPtr)(obj.opaque), (GDExtensionScriptInstancePtr)script);
};
EXPORT gd_extension_script_id gd_script(struct Object obj, struct Object language) {
    return (gd_extension_script_id)gdextension_object_get_script_instance((GDExtensionObjectPtr)(obj.opaque), (GDExtensionObjectPtr)(language.opaque));
};
EXPORT bool gd_script_defines_method(struct Object obj, struct StringName method) {
    return gdextension_object_has_script_method((GDExtensionObjectPtr)(obj.opaque), (GDExtensionConstStringNamePtr)&method);
};
EXPORT gd_extension_script_id gd_object_script_placeholder_create(struct Object language, struct Object script, struct Object owner) {
    return (gd_extension_script_id)gdextension_placeholder_script_instance_create((GDExtensionObjectPtr)(language.opaque), (GDExtensionObjectPtr)(script.opaque), (GDExtensionObjectPtr)(owner.opaque));
};
EXPORT void gd_object_script_placeholder_update(gd_extension_script_id script, struct Array array, struct Dictionary dict) {
    gdextension_placeholder_script_instance_update((GDExtensionScriptInstancePtr)script, (GDExtensionConstTypePtr)&array, (GDExtensionConstTypePtr)&dict);
};
EXPORT gd_addr gd_packed_array_access(VariantType type, PACKED_ARRAY_ARG(pa), int64_t i) {
    struct PackedArray pa = PACKED_ARRAY_ARG_GET(pa);
    switch (type) {
    case 29: return (gd_addr)gdextension_packed_byte_array_operator_index_const(&pa, i);
    case 30: return (gd_addr)gdextension_packed_int32_array_operator_index_const(&pa, i);
    case 31: return (gd_addr)gdextension_packed_int64_array_operator_index_const(&pa, i);
    case 32: return (gd_addr)gdextension_packed_float32_array_operator_index_const(&pa, i);
    case 33: return (gd_addr)gdextension_packed_float64_array_operator_index_const(&pa, i);
    case 34: return (gd_addr)gdextension_packed_string_array_operator_index_const(&pa, i);
    case 35: return (gd_addr)gdextension_packed_vector2_array_operator_index_const(&pa, i);
    case 36: return (gd_addr)gdextension_packed_vector3_array_operator_index_const(&pa, i);
    case 37: return (gd_addr)gdextension_packed_color_array_operator_index_const(&pa, i);
    case 38: return (gd_addr)gdextension_packed_vector4_array_operator_index_const(&pa, i);
    }
    return 0;
};
EXPORT gd_addr gd_packed_array_modify(VariantType type, PACKED_ARRAY_ARG(pa), int64_t i) {
    struct PackedArray pa = PACKED_ARRAY_ARG_GET(pa);
    switch (type) {
    case 29: return (gd_addr)gdextension_packed_byte_array_operator_index(&pa, i);
    case 30: return (gd_addr)gdextension_packed_int32_array_operator_index(&pa, i);
    case 31: return (gd_addr)gdextension_packed_int64_array_operator_index(&pa, i);
    case 32: return (gd_addr)gdextension_packed_float32_array_operator_index(&pa, i);
    case 33: return (gd_addr)gdextension_packed_float64_array_operator_index(&pa, i);
    case 34: return (gd_addr)gdextension_packed_string_array_operator_index(&pa, i);
    case 35: return (gd_addr)gdextension_packed_vector2_array_operator_index(&pa, i);
    case 36: return (gd_addr)gdextension_packed_vector3_array_operator_index(&pa, i);
    case 37: return (gd_addr)gdextension_packed_color_array_operator_index(&pa, i);
    case 38: return (gd_addr)gdextension_packed_vector4_array_operator_index(&pa, i);
    }
    return 0;
};
EXPORT void gd_array_set_index(struct Array a, int64_t i, VARIANT_ARG(v)) {
    *((struct Variant*)gdextension_array_operator_index(&a, i)) = VARIANT_ARG_GET(v);
};
EXPORT void gd_array_get_index(struct Array a, int64_t i, struct Variant* result) {
    *result = *(struct Variant*)gdextension_array_operator_index_const(&a, i);
};
EXPORT struct Object gd_ref_get_object(struct RefCounted ref) {
    return (struct Object){(uintptr_t)gdextension_ref_get_object((GDExtensionRefPtr)(ref.opaque))};
};
EXPORT void gd_ref_set_object(struct RefCounted ref, struct Object obj) {
    gdextension_ref_set_object((GDExtensionRefPtr)(ref.opaque), (GDExtensionObjectPtr)(obj.opaque));
};
EXPORT int32_t gd_string_access(struct String s, int64_t i) {
    return *(int32_t *)gdextension_string_operator_index_const((GDExtensionStringPtr)&s, i);
};
EXPORT struct String gd_string_resize(struct String s, int64_t size) {
    struct String ptr = s;
    gdextension_string_resize((GDExtensionStringPtr)&ptr, size);
    return ptr;
};
EXPORT gd_addr gd_string_memory(struct String s) {
    return (gd_addr)gdextension_string_operator_index((GDExtensionStringPtr)&s, 0);
};
EXPORT struct String gd_string_append(struct String s, struct String other) {
    struct String ptr = s;
    gdextension_string_operator_plus_eq_string((GDExtensionStringPtr)&ptr, (GDExtensionConstStringPtr)&other);
    return ptr;
};
EXPORT struct String gd_string_append_rune(struct String s, int32_t rune) {
    struct String ptr = s;
    gdextension_string_operator_plus_eq_char((GDExtensionStringPtr)&ptr, rune);
    return ptr;
};
EXPORT struct String gd_string_decode(gd_encoding enc, const char* s, int64_t len) {
    struct String ptr = {0};
    switch (enc) {
    case 0: // Latin1
        gdextension_string_new_with_latin1_chars_and_len((GDExtensionUninitializedStringPtr)&ptr, s, len);
        break;
    case 1: // UTF8
        gdextension_string_new_with_utf8_chars_and_len((GDExtensionUninitializedStringPtr)&ptr, s, len);
        break;
    case 2: // UTF16LE
        gdextension_string_new_with_utf16_chars_and_len2((GDExtensionUninitializedStringPtr)&ptr, (const char16_t *)s, len, true);
        break;
    case 3: // UTF16BE
        gdextension_string_new_with_utf16_chars_and_len2((GDExtensionUninitializedStringPtr)&ptr, (const char16_t *)s, len, false);
        break;
    case 4: // UTF32
        gdextension_string_new_with_utf32_chars_and_len((GDExtensionUninitializedStringPtr)&ptr, (const char32_t *)s, len);
        break;
    case 5: // Wide
        gdextension_string_new_with_wide_chars_and_len((GDExtensionUninitializedStringPtr)&ptr, (const wchar_t *)s, len);
        break;
    }
    return ptr;
};
EXPORT int64_t gd_string_encode(gd_encoding enc, struct String s, char* addr, int64_t len) {
    switch (enc) {
    case 0: // Latin1
        return gdextension_string_to_latin1_chars((GDExtensionStringPtr)&s, addr, len);
    case 1: // UTF8
        return gdextension_string_to_utf8_chars((GDExtensionStringPtr)&s, addr, len);
    case 2: case 3: // UTF16LE, UTF16BE
        return gdextension_string_to_utf16_chars((GDExtensionStringPtr)&s, (char16_t *)addr, len/2);
    case 4: // UTF32
        return gdextension_string_to_utf32_chars((GDExtensionStringPtr)&s, (char32_t *)addr, len/4);
    case 5: // Wide
        return gdextension_string_to_wide_chars((GDExtensionStringPtr)&s, (wchar_t *)addr, len/sizeof(wchar_t));
    }
    return 0;
};
EXPORT struct StringName gd_string_intern(gd_encoding enc, const char* s, int64_t len) {
    struct StringName ptr = {0};
    switch (enc) {
    case 0: { // Latin1
        char buf[256]; const char *s_ptr = fit_string(s, len, &buf[0], 256);
        gdextension_string_name_new_with_latin1_chars((GDExtensionUninitializedStringNamePtr)&ptr, s_ptr, false);
        if (s_ptr != &buf[0]) gdextension_mem_free((void *)s_ptr);
        break;
    }
    case 1: // UTF8
        gdextension_string_name_new_with_utf8_chars_and_len((GDExtensionUninitializedStringNamePtr)&ptr, s, len);
        break;
    }
    return ptr;
};
EXPORT struct String gd_variant_type_name(VariantType vtype) {
    struct String name = {0};
    gdextension_variant_get_type_name((GDExtensionVariantType)vtype, &name);
    return name;
};
EXPORT void gd_builtin_from(VariantType vtype, VARIANT_ARG(v), gd_addr result) {
    uint64_t self[4] = {v_1, v_2, v_3};
    type_from_variant_constructors[vtype]((GDExtensionTypePtr)result, &self[0]);
};
EXPORT void gd_variant_from(VariantType vtype, struct Variant* result, gd_addr args) {
    variant_from_type_constructors[vtype]((GDExtensionUninitializedVariantPtr)result, &args);
};
EXPORT void gd_variant_type_call(VariantType vtype, struct StringName static_method_name, struct Variant* result, int64_t argc, struct Variant args[], gd_error* err) {
    void *points[16]; prepare_variants(&points[0], argc, (gd_addr)args);
    gdextension_variant_call_static((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&static_method_name, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};
EXPORT bool gd_variant_type_convertable(VariantType a, VariantType b, bool strict) {
    if (strict) {
        return gdextension_variant_can_convert_strict((GDExtensionVariantType)a, (GDExtensionVariantType)b);
    } else {
        return gdextension_variant_can_convert((GDExtensionVariantType)a, (GDExtensionVariantType)b);
    }
};
EXPORT void gd_variant_type_setup_array(struct Array a, VariantType elem, struct StringName class_name, VARIANT_ARG(v)) {
    uint64_t script[3] = {v_1, v_2, v_3};
    gdextension_array_set_typed(&a, (GDExtensionVariantType)elem, &class_name, &script[0]);
};
EXPORT void gd_variant_type_setup_dictionary(struct Dictionary d,
    VariantType key, struct StringName key_class_name, VARIANT_ARG(key_script),
    VariantType val, struct StringName val_class_name, VARIANT_ARG(val_script)
) {
    uint64_t k_script[3] = {key_script_1, key_script_2, key_script_3};
    uint64_t v_script[3] = {val_script_1, val_script_2, val_script_3};
    gdextension_dictionary_set_typed(&d, (GDExtensionVariantType)key, &key_class_name, &k_script[0], (GDExtensionVariantType)val, &val_class_name, &v_script[0]);
};
EXPORT void gd_variant_type_constant(VariantType t, struct StringName constant, struct Variant* result) {
    gdextension_variant_get_constant_value((GDExtensionVariantType)t, (GDExtensionConstStringNamePtr)&constant, &result);
};
EXPORT gd_caller_id gd_builtin_method(VariantType t, struct StringName method, int64_t hash) {
    return (gd_caller_id)gdextension_variant_get_ptr_builtin_method((GDExtensionVariantType)t, (GDExtensionConstStringNamePtr)&method, hash);
};
EXPORT gd_constructor_id gd_constructor(VariantType vtype, int64_t n) {
    return (uintptr_t)gdextension_variant_get_ptr_constructor((GDExtensionVariantType)vtype, n);
};
EXPORT void gd_builtin_free(VariantType t, const gd_addr value) {
    variant_ptr_destructors[t]((GDExtensionTypePtr)value);
};
EXPORT void gd_variant_zero(struct Variant* result) {
    gdextension_variant_new_nil((GDExtensionUninitializedVariantPtr)result);
};
EXPORT void gd_variant_copy(VARIANT_ARG(v), struct Variant* result, bool deep) {
    struct Variant self = VARIANT_ARG_GET(v);
    if (deep) {
        gdextension_variant_duplicate(&self, (GDExtensionVariantPtr)result, true);
    } else {
        gdextension_variant_new_copy((GDExtensionUninitializedVariantPtr)result, &self);
    }
};
EXPORT void gd_variant_call(VARIANT_ARG(v), struct StringName method, struct Variant* result, int64_t argc, struct Variant args[], gd_error* err) {
    struct Variant self = VARIANT_ARG_GET(v);
    void *points[16]; prepare_variants(&points[0], argc, (gd_addr)args);
    gdextension_variant_call(&self, (GDExtensionConstStringNamePtr)&method, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};
EXPORT bool gd_variant_eval(VariantOperator op, VARIANT_ARG(a), VARIANT_ARG(b), struct Variant* result) {
    uint64_t a[3] = {a_1, a_2, a_3};
    uint64_t b[3] = {b_1, b_2, b_3};
    GDExtensionBool valid = false;
    gdextension_variant_evaluate((GDExtensionVariantOperator)op, &a[0], &b[0], (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};
EXPORT int64_t gd_variant_hash(VARIANT_ARG(v), int64_t recursion_count) {
    struct Variant self = VARIANT_ARG_GET(v);
    if (recursion_count > 0) {
        return gdextension_variant_recursive_hash(&self, recursion_count);
    }
    return gdextension_variant_hash(&self);
};
EXPORT bool gd_variant_bool(VARIANT_ARG(v)) {
    struct Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_booleanize(&self);
};
EXPORT struct String gd_variant_text(VARIANT_ARG(v)) {
    struct Variant self = VARIANT_ARG_GET(v);
    struct String text = {0};
    gdextension_variant_stringify(&self, &text);
    return text;
};
EXPORT VariantType gd_variant_type(VARIANT_ARG(v)) {
    struct Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_get_type(&self);
};
// gd_variant_deep_copy and gd_variant_deep_hash merged into gd_variant_copy and gd_variant_hash above.
EXPORT bool gd_variant_get_keyed(VARIANT_ARG(v), VARIANT_ARG(key), struct Variant* result) {
    struct Variant self = VARIANT_ARG_GET(v);
    struct Variant index = VARIANT_ARG_GET(key);
    GDExtensionBool valid = false;
    gdextension_variant_get_keyed(&self, &index, (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};
EXPORT bool gd_variant_get_index(VARIANT_ARG(v), int64_t i, struct Variant* result, gd_error* err) {
    struct Variant self = VARIANT_ARG_GET(v);
    GDExtensionBool valid = false;
    GDExtensionBool oob = false;
    gdextension_variant_get_indexed(&self, i, (GDExtensionUninitializedVariantPtr)result, &valid, &oob);
    if (oob) ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return valid;
};
EXPORT bool gd_variant_get_field(VARIANT_ARG(v), struct StringName field, struct Variant* result) {
    struct Variant self = VARIANT_ARG_GET(v);
    GDExtensionBool valid = false;
    gdextension_variant_get_named(&self, (GDExtensionConstStringNamePtr)&field, (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};
EXPORT bool gd_variant_type_has_property(VariantType t, struct StringName property) {
    return gdextension_variant_has_member((GDExtensionVariantType)t, (GDExtensionConstStringNamePtr)&property);
};
EXPORT bool gd_variant_has_key(VARIANT_ARG(v), VARIANT_ARG(idx)) {
    struct Variant self = VARIANT_ARG_GET(v);
    struct Variant index = VARIANT_ARG_GET(idx);
    GDExtensionBool valid = false;
    gdextension_variant_has_key(&self, &index, &valid);
    return valid;
};
EXPORT bool gd_variant_has_method(VARIANT_ARG(v), struct StringName method) {
    struct Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_has_method(&self, (GDExtensionConstStringNamePtr)&method);
};
EXPORT bool gd_variant_set_keyed(VARIANT_ARG(v), VARIANT_ARG(key), VARIANT_ARG(val)) {
    struct Variant self = VARIANT_ARG_GET(v);
    struct Variant index = VARIANT_ARG_GET(key);
    struct Variant value = VARIANT_ARG_GET(val);
    GDExtensionBool valid = false;
    gdextension_variant_set_keyed(&self, &index, &value, &valid);
    return valid;
};
EXPORT bool gd_variant_set_index(VARIANT_ARG(v), int64_t i, VARIANT_ARG(val), gd_error* err) {
    struct Variant self = VARIANT_ARG_GET(v);
    struct Variant value = VARIANT_ARG_GET(val);
    GDExtensionBool valid = false;
    GDExtensionBool oob = false;
    gdextension_variant_set_indexed(&self, i, &value, &valid, &oob);
    if (oob) ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return valid;
};
EXPORT bool gd_variant_set_field(VARIANT_ARG(v), struct StringName field, VARIANT_ARG(val)) {
    struct Variant self = VARIANT_ARG_GET(v);
    struct Variant value = VARIANT_ARG_GET(val);
    GDExtensionBool valid = false;
    gdextension_variant_set_named(&self, (GDExtensionConstStringNamePtr)&field, &value, &valid);
    return valid;
};
EXPORT void gd_builtin_eval(gd_evaluator_id fn, gd_addr result, gd_shape shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    ((GDExtensionPtrOperatorEvaluator)fn)(points[0], points[1], (GDExtensionTypePtr)result);
};
// Removed: was a duplicate of gd_call, but with GDExtensionPtrBuiltInMethod — this logic is in gd_builtin_call now.
EXPORT void gd_variant_free(VARIANT_ARG(v)) {
    struct Variant self = VARIANT_ARG_GET(v);
    gdextension_variant_destroy(&self);
};
EXPORT gd_addr gd_variant_data(VariantType vtype, VARIANT_ARG(v)) {
    struct Variant self = VARIANT_ARG_GET(v);
    return (gd_addr)variant_internal_ptr_funcs[vtype](&self);
};
EXPORT void gd_builtin_get_field(gd_getter_id getter, gd_addr result, gd_shape shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    ((GDExtensionPtrGetter)getter)(points[0], (void*)&result);
};
EXPORT void gd_builtin_get_array(VariantType vtype, int64_t i, gd_addr result, gd_shape shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    variant_ptr_indexed_getters[vtype]((GDExtensionConstTypePtr)result, i, points[0]);
};
EXPORT void gd_builtin_get_keyed(VariantType vtype, gd_addr result, gd_shape shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    variant_ptr_keyed_getters[vtype](points[0], points[1], (GDExtensionTypePtr)result);
};
EXPORT void gd_builtin_set_field(gd_setter_id setter, gd_shape shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    ((GDExtensionPtrSetter)setter)(points[1], points[2]);
};
EXPORT void gd_builtin_set_array(VariantType vtype, int64_t i, gd_shape shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    variant_ptr_indexed_setters[vtype](points[0], i, points[1]);
};
EXPORT void gd_builtin_set_keyed(VariantType vtype, gd_shape shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    variant_ptr_keyed_setters[vtype](points[0], points[1], points[2]);
};
EXPORT gd_setter_id gd_setter(VariantType t, struct StringName property) {
    return (gd_setter_id)gdextension_variant_get_ptr_setter((GDExtensionVariantType)t, (GDExtensionConstStringNamePtr)&property);
}
EXPORT gd_getter_id gd_getter(VariantType t, struct StringName property) {
    return (gd_getter_id)gdextension_variant_get_ptr_getter((GDExtensionVariantType)t, (GDExtensionConstStringNamePtr)&property);
}
EXPORT void gd_builtin_make(gd_constructor_id fn, gd_addr result, gd_shape shape, gd_addr args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, args);
    ((GDExtensionPtrConstructor)fn)((GDExtensionUninitializedTypePtr)result, (const GDExtensionConstTypePtr *)&points[0]);
}
EXPORT void gd_variant_make(VariantType vtype, struct Variant* result, int64_t argc, struct Variant args[], gd_error* err) {
    void *points[16]; prepare_variants(&points[0], argc, (gd_addr)args);
    gdextension_variant_construct((GDExtensionVariantType)vtype, (GDExtensionUninitializedVariantPtr)result, (const GDExtensionConstVariantPtr *)&points[0], argc, (GDExtensionCallError*)err);
}
EXPORT void gd_builtin_call(gd_addr self, gd_caller_id fn, gd_addr result, gd_shape shape, gd_addr args) {
    void *points[16]; uint8_t argc = prepare_callframe(2, &points[0], shape, args);
    ((GDExtensionPtrBuiltInMethod)fn)((GDExtensionTypePtr*)self, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)result, argc);
}
EXPORT gd_evaluator_id gd_evaluator(VariantOperator op, VariantType a, VariantType b) {
    return (uintptr_t)gdextension_variant_get_ptr_operator_evaluator((GDExtensionVariantOperator)op, (GDExtensionVariantType)a, (GDExtensionVariantType)b);
}

#ifdef __EMSCRIPTEN__
}

// Raw WASM exports for direct WASM-to-WASM calls (bypassing embind JS trampolines).
// These use native types (real uint64_t instead of split uint32 pairs).

extern "C" {
    // WASM uses 8-byte uintptr: Object(0,8) Method(8,8) Shape(16,8) Args(24,256)
    EMSCRIPTEN_KEEPALIVE void wasm_gd_ring_flush(uint32_t ring_base, uint32_t entry_stride, uint32_t tail, uint32_t head) {
        for (uint32_t i = tail; i != head; i++) {
            uint8_t *entry = (uint8_t *)(uintptr_t)(ring_base + (i & 0xFF) * entry_stride);
            uint32_t object = *(uint32_t *)(entry + 0);
            uint32_t method = *(uint32_t *)(entry + 8);
            uint64_t shape = *(uint64_t *)(entry + 16);
            uint32_t args = (uint32_t)(uintptr_t)(entry + 24);
            void *points[16];
            prepare_callframe(1, &points[0], shape, args);
            gdextension_object_method_bind_ptrcall(
                (GDExtensionMethodBindPtr)(uintptr_t)method,
                (GDExtensionObjectPtr)(uintptr_t)object,
                (const GDExtensionConstTypePtr*)&points[0],
                NULL);
        }
    }
}
#endif // __EMSCRIPTEN__
