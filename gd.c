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
    #define UINT64_FROM(v) ((uint64_t)(v##_1) << 32) | ((uint64_t)(v##_2))
    #define UINT64_MAKE(v) (uint32_t)(v >> 32), (uint32_t)(v & 0xFFFFFFFF)
    #define INT64_FROM(v) std::__bit_cast<int64_t>(((uint64_t)(v##_1) << 32) | ((uint64_t)(v##_2)))
    #define BUFFER_POINTER(s) (char*)(s)
    #define EXPORT EMSCRIPTEN_KEEPALIVE
    #define RETURN(T, v) *(T*)result = v; return
    extern "C" {
#else
	#ifdef _WIN32
	#define EXPORT __declspec(dllexport)
	#else
	#define EXPORT
	#endif
    #define UINT64_FROM(v) (v)
    #define UINT64_MAKE(v) (v)
    #define INT64_FROM(v) (v)
    #define BUFFER_POINTER(s) (s)
    #define RETURN(T, v) return v
#endif

#define VARIANT_ARG_GET(n) (Variant){n##_1, {n##_2, n##_3}}
#define VARIANT_ARG_PUT(v) v.tag, v.payload[0], v.payload[1]
#define PACKED_ARRAY_ARG_GET(n) (PackedArray){n##_1, n##_2}

// X-macro list: each entry is X(snake_name, GDExtensionInterfaceType)
// Declares variables and loads proc addresses from a single source of truth.
#define GD_PROC_LIST(X) \
    X(mem_alloc, GDExtensionInterfaceMemAlloc) \
    X(mem_realloc, GDExtensionInterfaceMemRealloc) \
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

static void engine_exit(void *ignore, GDExtensionInitializationLevel level) {
    gd_on_engine_exit(level);
}
static void extension_instance_dynamic_call(void *method_userdata, GDExtensionClassInstancePtr p_instance, const GDExtensionConstVariantPtr *p_args, GDExtensionInt p_argument_count, GDExtensionVariantPtr r_return, GDExtensionCallError *r_error) {
    gd_on_extension_instance_dynamic_call((uintptr_t)p_instance, (uintptr_t)method_userdata, (Variant*)r_return, p_argument_count, (VariadicVariants)p_args, (CallError*)r_error);
}
static void extension_instance_checked_call(void *method_userdata, GDExtensionClassInstancePtr p_instance, const GDExtensionConstTypePtr *p_args, GDExtensionTypePtr r_ret) {
    gd_on_extension_instance_checked_call((ExtensionInstanceID)p_instance, (FunctionID)method_userdata, (UnsafePointer)r_ret, (UnsafePointer)p_args);
}
static void *extension_binding_created(void *p_token, void *p_instance) {
    return (void *)gd_on_extension_binding_created((uintptr_t)p_instance);
}
static void extension_binding_removed(void *p_token, void *p_instance, void *p_binding) {
    gd_on_extension_binding_removed((uintptr_t)p_instance, (uintptr_t)p_binding);
}
static GDExtensionBool extension_binding_reference(void *p_token, void *p_binding, GDExtensionBool p_reference) {
    return gd_on_extension_binding_reference((uintptr_t)p_binding, p_reference);
}
static GDExtensionInstanceBindingCallbacks instance_binding_callbacks = {
    .create_callback = extension_binding_created,
    .free_callback = extension_binding_removed,
   // .reference_callback = extension_binding_reference
};

GDExtensionTypeFromVariantConstructorFunc type_from_variant_constructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionVariantFromTypeConstructorFunc variant_from_type_constructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrDestructor variant_ptr_destructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionVariantGetInternalPtrFunc variant_internal_ptr_funcs[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrIndexedSetter variant_ptr_indexed_setters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrIndexedGetter variant_ptr_indexed_getters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrKeyedSetter variant_ptr_keyed_setters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrKeyedGetter variant_ptr_keyed_getters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];

static void engine_init(void *ignore, GDExtensionInitializationLevel level) {
    gd_on_engine_init(level);
}

EXPORT GDExtensionBool gd_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization) {
    #define LOAD_PROC(name, type) gdextension_##name = (type)p_get_proc_address(#name);
    GD_PROC_LIST(LOAD_PROC)
    #undef LOAD_PROC
    gd_library = p_library;
    r_initialization->userdata = 0;
    r_initialization->minimum_initialization_level = GDEXTENSION_INITIALIZATION_CORE;
    r_initialization->initialize = engine_init;
    r_initialization->deinitialize = engine_exit;
    gdextension_get_godot_version2(&gd_godot_version_cached);
    GDExtensionMainLoopCallbacks callbacks = {
    	.startup_func = gd_on_first_frame,
    	.shutdown_func = gd_on_final_frame,
    	.frame_func = gd_on_every_frame,
    };
    gdextension_register_main_loop_callbacks(p_library, &callbacks);
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
static void prepare_variants(void **frame, uint32_t argc, ANY args) {
    uint8_t *head = (uint8_t*)args;
    for (int i = 0; i < argc; i++) {
        frame[i] = head;
        head += 24;
    }
}
// Helper macro to align a value to the next multiple of 'align'
#define ALIGN_UP(value, align) (((value) + ((align) - 1)) & ~((align) - 1))
static uint8_t prepare_callframe(int skip, void **frame, uint64_t shape, ANY args) {
    uint8_t *head = (uint8_t *)args;
    ptrdiff_t offset = 0; // Track current offset in the frame
    for (int i = skip; i < 16; i++) {
        Shape code = (Shape)((shape >> (i * 4)) & 0xF);
        uint32_t size;
        uint32_t align;
        // Determine size based on code
        switch (code) {
            case ShapeEmpty: size = 0; frame[i-skip] = NULL; return i-skip;
            case ShapeBytes1: size = 1; align = 1; break;
            case ShapeBytes2: size = 2; align = 2; break;
            case ShapeBytes4: size = 4; align = 4; break;
            case ShapeBytes8: size = 8; align = 8; break;
            case ShapeBytes4x2: size = 4*2; align = 4; break;
            case ShapeBytes4x3: size = 4*3; align = 4; break;
            case ShapeBytes8x2: size = 8*2; align = 8; break;
            case ShapeBytes4x4: size = 4*4; align = 4; break;
            case ShapeBytes8x3: size = 8*3; align = 8; break;
            case ShapeBytes4x6: size = 4*6; align = 4; break;
            case ShapeBytes4x9: size = 4*9; align = 4; break;
            case ShapeBytes4x12: size = 4*12; align = 4; break;
            case ShapeBytes4x16: size = 4*16; align = 4; break;
        }
        offset = ALIGN_UP(offset, align);
        frame[i-skip] = head + offset;     // Set frame pointer to the aligned address
        offset += size;                 // Move offset forward by the size of the current argument
    }
    return 16-skip;
}

static void callable_called(void *c, const GDExtensionConstVariantPtr *args, GDExtensionInt argc, GDExtensionVariantPtr ret, GDExtensionCallError *err) {
	gd_on_callable_called((CallableID)c, (Variant*)ret, argc, (VariadicVariants)args, (CallError*)err);
}
static GDExtensionBool callable_verify(void *c) {
    return gd_on_callable_verify((CallableID)c);
}
static void callable_delete(void *c) {
    gd_on_callable_delete((CallableID)c);
}
static uint32_t callable_hashed(void *c) {
    return gd_on_callable_hashed((CallableID)c);
}
static GDExtensionBool callable_compare(void *a, void *b) {
    return gd_on_callable_sorted((CallableID)a, (CallableID)b) == 0;
}
static GDExtensionBool callable_less_than(void *a, void *b) {
    return gd_on_callable_sorted((CallableID)a, (CallableID)b) < 0;
}
static void callable_string(void *c, GDExtensionBool *r_is_valid, GDExtensionStringPtr r_out) {
    String s = gd_on_callable_string((CallableID)c);
    *r_is_valid = (GDExtensionBool)(s != 0);
    *((uintptr_t*)r_out) = s;
}
static GDExtensionInt callable_length(void *callable_userdata, GDExtensionBool *r_is_valid) {
    *r_is_valid = true;
    return (GDExtensionInt)(gd_on_callable_length((CallableID)callable_userdata));
}
EXPORT void gd_callable_create(CallableID id, ObjectID object, Callable* result) {
    GDExtensionCallableCustomInfo2 info = {
        .callable_userdata = (void*)id,
        .token = gd_library,
        .object_id = object,
        .call_func = callable_called,
        .is_valid_func = callable_verify,
        .free_func = callable_delete,
        .hash_func = callable_hashed,
        .equal_func = callable_compare,
        .less_than_func = callable_less_than,
        .to_string_func = callable_string,
        .get_argument_count_func = callable_length,
    };
    gdextension_callable_custom_create2(result, &info);
}
EXPORT CallableID gd_callable_lookup(CALLABLE_ARG(c)) {
    Callable c = CALLABLE_ARG_GET(c);
    return (CallableID)gdextension_callable_custom_get_userdata(&c, gd_library);
};

EXPORT uintptr_t gd_builtin_name(uintptr_t name, int64_t hash) { return (uintptr_t)gdextension_variant_get_ptr_utility_function((GDExtensionConstStringNamePtr)&name, hash);}
EXPORT void gd_builtin_call(uintptr_t fn, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; uint8_t argc = prepare_callframe(1, &points[0], shape, (ANY)args);
    ((GDExtensionPtrUtilityFunction)fn)((GDExtensionTypePtr)result, (GDExtensionConstTypePtr*)&points[0], argc);
}
EXPORT uintptr_t gd_library_location() {
    uintptr_t s;
    gdextension_get_library_path(gd_library, &s);
    return s;
}
EXPORT void gd_classdb_FileAccess_write(Object FileAccess, char* buf, Int len) {
    gdextension_file_access_store_buffer((GDExtensionObjectPtr)FileAccess, (const uint8_t *)BUFFER_POINTER(buf), len);
};
EXPORT Int gd_classdb_FileAccess_read(Object FileAccess, char* buf, Int len) {
    return gdextension_file_access_get_buffer((GDExtensionObjectPtr)FileAccess, (uint8_t *)BUFFER_POINTER(buf), len);
};
EXPORT UnsafePointer gd_classdb_Image_unsafe(Object Image) {
    return (UnsafePointer)gdextension_image_ptrw((GDExtensionObjectPtr)Image);
};
EXPORT uint8_t gd_classdb_Image_access(Object Image, Int offset) {
    return gdextension_image_ptr((GDExtensionObjectPtr)Image)[offset];
};
typedef struct {
    int32_t push;
    int32_t size;
    GDExtensionClassMethodInfo *info;
} method_list;
typedef struct {
    int32_t push;
    int32_t size;
    GDExtensionPropertyInfo *info;
    GDExtensionClassMethodArgumentMetadata *meta;
} property_list;
EXPORT uintptr_t gd_method_list_make(Int length) {
    method_list *list = (method_list *)gdextension_mem_alloc(sizeof(method_list));
    list->push = 0;
    list->size = length;
    list->info = (GDExtensionClassMethodInfo*)gdextension_mem_alloc(sizeof(GDExtensionClassMethodInfo) * length);
    return (uintptr_t)list;
};
EXPORT void gd_method_list_push(uintptr_t list_p, uintptr_t name, uintptr_t method, uint32_t method_flags, uintptr_t return_value_info, uintptr_t arguments_info, Int default_argument_count, ANY default_arguments) {
    method_list *list = (method_list *)list_p;
    if (list->push >= list->size) return;
    GDExtensionClassMethodInfo *info = &list->info[list->push++];
    property_list *return_value = (property_list *)return_value_info;
    property_list *arguments = (property_list *)arguments_info;
    uintptr_t *name_allocated = (uintptr_t *)gdextension_mem_alloc(sizeof(uintptr_t));
    *name_allocated = name;
    info->name = (GDExtensionStringNamePtr)name_allocated;
    info->method_userdata = (void *)method;
    info->call_func = extension_instance_dynamic_call;
    info->ptrcall_func = extension_instance_checked_call;
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
    void **points = (void **)gdextension_mem_alloc(sizeof(void*) * default_argument_count);
    prepare_variants(&points[0], default_argument_count, default_arguments);
    info->default_argument_count = default_argument_count;
    info->default_arguments = points;
};
EXPORT void gd_method_list_free(uintptr_t list_p) {
    method_list *list = (method_list *)list_p;
    for (int i = 0; i < list->push; i++) {
        gdextension_mem_free(list->info[i].name);
        if (list->info[i].default_arguments) {
            gdextension_mem_free(list->info[i].default_arguments);
        }
    }
    gdextension_mem_free(list->info); gdextension_mem_free(list);
};
EXPORT uintptr_t gd_property_list_make(Int length) {
    property_list *list = (property_list*)gdextension_mem_alloc(sizeof(property_list));
    list->push = 0;
    list->size = length;
    list->info = (GDExtensionPropertyInfo*)gdextension_mem_alloc(sizeof(GDExtensionPropertyInfo) * length);
    list->meta = (GDExtensionClassMethodArgumentMetadata*)gdextension_mem_alloc(sizeof(GDExtensionClassMethodArgumentMetadata) * length);
    return (uintptr_t)list;
};
EXPORT void gd_property_list_push(uintptr_t list_p, uint32_t vtype, uintptr_t name, uintptr_t class_name, uint32_t hint, uintptr_t hint_string, uint32_t usage, uint32_t meta) {
    property_list *list = (property_list *)list_p;
    if (list->push >= list->size) return;
    GDExtensionPropertyInfo *info = &list->info[list->push++];
    GDExtensionClassMethodArgumentMetadata *meta_info = &list->meta[list->push - 1];
    uintptr_t *name_allocated = (uintptr_t *)gdextension_mem_alloc(sizeof(uintptr_t));
    *name_allocated = name;
    uintptr_t *class_name_allocated = (uintptr_t *)gdextension_mem_alloc(sizeof(uintptr_t));
    *class_name_allocated = class_name;
    uintptr_t *hint_string_allocated = (uintptr_t *)gdextension_mem_alloc(sizeof(uintptr_t));
    *hint_string_allocated = hint_string;
    info->type = (GDExtensionVariantType)vtype;
    info->name = (GDExtensionStringNamePtr)name_allocated;
    info->class_name = (GDExtensionStringNamePtr)class_name_allocated;
    info->hint = hint;
    info->hint_string = (GDExtensionStringPtr)hint_string_allocated;
    info->usage = usage;
    *meta_info = (GDExtensionClassMethodArgumentMetadata)meta;
};
EXPORT void gd_property_list_free(uintptr_t list_p) {
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
static GDExtensionBool extension_instance_set(GDExtensionClassInstancePtr instance, GDExtensionConstStringNamePtr field, GDExtensionConstVariantPtr value) {
    Variant v = *(Variant*)value;
    return gd_on_extension_instance_set((ExtensionInstanceID)instance, *(StringName*)field, VARIANT_ARG_PUT(v));
}
static GDExtensionBool extension_instance_get(GDExtensionClassInstancePtr instance, GDExtensionConstStringNamePtr field, GDExtensionVariantPtr value) {
    return gd_on_extension_instance_get((uintptr_t)instance, *(uintptr_t*)field, (Variant*)value);
}
static GDExtensionBool extension_instance_property_has_default(GDExtensionClassInstancePtr instance, GDExtensionConstStringNamePtr field) {
    return gd_on_extension_instance_property_has_default((uintptr_t)instance, *(uintptr_t*)field);
}
static GDExtensionBool extension_instance_property_get_default(GDExtensionClassInstancePtr instance, GDExtensionConstStringNamePtr field, GDExtensionVariantPtr value) {
    return gd_on_extension_instance_property_get_default((uintptr_t)instance, *(uintptr_t*)field, (Variant*)value);
}
static const GDExtensionPropertyInfo *extension_instance_property_list(GDExtensionClassInstancePtr instance, uint32_t *count) {
    property_list *list = (property_list*)gd_on_extension_instance_property_list((uintptr_t)instance);
    GDExtensionPropertyInfo *info = list ? list->info : NULL;
    *count = list ? list->push : 0;
    if (list && list->meta) {
        gdextension_mem_free(list->meta);
    }
    return info;
}
static void class_free_property_list_func(GDExtensionClassInstancePtr instance, const GDExtensionPropertyInfo *list, uint32_t count) {
   if (list) gdextension_mem_free((void*)list);
}
static GDExtensionBool extension_instance_property_validation(GDExtensionClassInstancePtr instance, GDExtensionPropertyInfo *field) {
    property_list list = {
        .push = 1,
        .size = 1,
        .info = field,
        .meta = NULL
    };
    return gd_on_extension_instance_property_validation((uintptr_t)instance, (uintptr_t)&list);
}
static void extension_instance_stringify(GDExtensionClassInstancePtr instance, GDExtensionBool *ok, GDExtensionStringPtr s) {
    uint32_t result = gd_on_extension_instance_stringify((uintptr_t)instance);
    if (result) {
        *(uint32_t*)s = result;
        *ok = true;
    } else {
        gdextension_string_new_with_latin1_chars(s, ""); // FIXME/TODO remove in 4.5 (where my PR to fix this has been merged https://github.com/godotengine/godot/pull/105546)
        *ok = false;
    }
}
static void extension_instance_reference(GDExtensionClassInstancePtr instance) {
    gd_on_extension_instance_reference((uintptr_t)instance, true);
}
static GDExtensionBool extension_instance_unreference(GDExtensionClassInstancePtr instance) {
    return gd_on_extension_instance_reference((uintptr_t)instance, false);
}
static GDExtensionObjectPtr extension_class_create(void *user_data, GDExtensionBool notify_postinitialize) {
    return (GDExtensionObjectPtr)gd_on_extension_class_create((uintptr_t)user_data, notify_postinitialize);
}
static void *extension_class_caller(void *user_data, GDExtensionConstStringNamePtr name, uint32_t hash) {
    return (void*)gd_on_extension_class_caller((uintptr_t)user_data, *(uintptr_t*)name, hash);
}
static void extension_instance_called(GDExtensionClassInstancePtr p_instance, GDExtensionConstStringNamePtr p_name, void *p_virtual_call_userdata, const GDExtensionConstTypePtr *p_args, GDExtensionTypePtr r_ret) {
    gd_on_extension_instance_called((ExtensionInstanceID)p_instance, (FunctionID)p_virtual_call_userdata, (UnsafePointer)r_ret, (UnsafePointer)p_args);
}
static void extension_instance_free(void *p_class_userdata, GDExtensionClassInstancePtr p_instance) {
    gd_on_extension_instance_free((uintptr_t)p_instance);
}
EXPORT void gd_classdb_register(uintptr_t class_name, uintptr_t parent, uintptr_t id, bool is_virtual, bool abstract, bool exposed, bool runtime, uintptr_t icon_path) {
    GDExtensionClassCreationInfo5 info = {
        .is_virtual = is_virtual,
        .is_abstract = abstract,
        .is_exposed = exposed,
        .is_runtime = runtime,
        .icon_path = (GDExtensionConstStringNamePtr)&icon_path,
        .set_func = extension_instance_set,
        .get_func = extension_instance_get,
        .get_property_list_func = extension_instance_property_list,
        .free_property_list_func = class_free_property_list_func,
        .property_can_revert_func = extension_instance_property_has_default,
        .property_get_revert_func = extension_instance_property_get_default,
        .validate_property_func = extension_instance_property_validation,
        .notification_func = (GDExtensionClassNotification2)gd_on_extension_instance_notification,
        .to_string_func = extension_instance_stringify,
        //.reference_func = (GDExtensionClassReference)class_reference, // FIXME JavaScript error: null function or function signature mismatch
        //.unreference_func = (GDExtensionClassUnreference)class_unreference, // FIXME JavaScript error: null function or function signature mismatch
        .create_instance_func = extension_class_create,
        .free_instance_func = extension_instance_free,
        .get_virtual_call_data_func = extension_class_caller,
        .call_virtual_with_data_func = extension_instance_called,
        .class_userdata = (void *)id,
    };
    gdextension_classdb_register_extension_class5(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&parent, &info);
};
EXPORT void gd_classdb_register_methods(uintptr_t class_name, uintptr_t methods) {
    method_list *list = (method_list *)methods;
    for (int i = 0; i < list->push; i++) {
        GDExtensionClassMethodInfo *info = &list->info[i];
        gdextension_classdb_register_extension_class_method(gd_library, (GDExtensionConstStringNamePtr)&class_name, info);
    }
};
EXPORT void gd_classdb_register_constant(uintptr_t class_name, uintptr_t enum_name, uintptr_t name, INT64(value), bool bitfield) {
    gdextension_classdb_register_extension_class_integer_constant(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&enum_name, (GDExtensionConstStringNamePtr)&name, INT64_FROM(value), bitfield);
};
EXPORT void gd_classdb_register_property(uintptr_t class_name, uintptr_t info, uintptr_t setter, uintptr_t getter) {
    property_list *list = (property_list *)info;
    gdextension_classdb_register_extension_class_property(gd_library, (GDExtensionConstStringNamePtr)&class_name, list->info, (GDExtensionConstStringNamePtr)&setter, (GDExtensionConstStringNamePtr)&getter);
};
EXPORT void gd_classdb_register_property_indexed(uintptr_t class_name, uintptr_t info, uintptr_t setter, uintptr_t getter, INT64(index)) {
    property_list *list = (property_list *)info;
    gdextension_classdb_register_extension_class_property_indexed(gd_library, (GDExtensionConstStringNamePtr)&class_name, list->info, (GDExtensionConstStringNamePtr)&setter, (GDExtensionConstStringNamePtr)&getter, INT64_FROM(index));
};
EXPORT void gd_classdb_register_property_group(uintptr_t class_name, uintptr_t group, uintptr_t prefix) {
    gdextension_classdb_register_extension_class_property_group(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&group, (GDExtensionConstStringPtr)&prefix);
};
EXPORT void gd_classdb_register_property_sub_group(uintptr_t class_name, uintptr_t subgroup, uintptr_t prefix) {
    gdextension_classdb_register_extension_class_property_subgroup(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&subgroup, (GDExtensionConstStringPtr)&prefix);
};
EXPORT void gd_classdb_register_signal(uintptr_t class_name, uintptr_t name, uintptr_t args) {
    property_list *list = (property_list *)args;
    gdextension_classdb_register_extension_class_signal(gd_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&name, list->info, list->push);
};
EXPORT void gd_classdb_register_removal(uintptr_t class_name) {
    gdextension_classdb_unregister_extension_class(gd_library, (GDExtensionConstStringNamePtr)&class_name);
};
EXPORT void gd_classdb_WorkerThreadPool_add_task(uintptr_t WorkerPool, uintptr_t task_id, bool priority, uintptr_t description) {
    gdextension_worker_thread_pool_add_native_task((GDExtensionObjectPtr)WorkerPool, (GDExtensionWorkerThreadPoolTask)gd_on_worker_thread_pool_task, (void *)task_id, priority, (GDExtensionConstStringNamePtr)&description);
};
EXPORT void gd_classdb_WorkerThreadPool_add_group_task(uintptr_t WorkerPool, uintptr_t task_id, int32_t elements, int32_t tasks, bool priority, uintptr_t description) {
    gdextension_worker_thread_pool_add_native_group_task((GDExtensionObjectPtr)WorkerPool, (GDExtensionWorkerThreadPoolGroupTask)gd_on_worker_thread_pool_group_task, (void *)task_id, elements, tasks, priority, (GDExtensionConstStringNamePtr)&description);
};
EXPORT Int gd_classdb_XMLParser_load(uintptr_t XMLParser, char* buf, Int len) {
    return gdextension_xml_parser_open_buffer((GDExtensionObjectPtr)XMLParser, (const uint8_t *)BUFFER_POINTER(buf), len);
};
EXPORT void gd_packed_dictionary_access(Dictionary dict, VARIANT_ARG(k), Variant* result) {
    Variant key = VARIANT_ARG_GET(k);
    Variant* value = (Variant*)gdextension_dictionary_operator_index_const((GDExtensionTypePtr)&dict, &key);
    if (!value) return;
    *result = *value;
};
EXPORT void gd_packed_dictionary_modify(Dictionary dict, VARIANT_ARG(k), VARIANT_ARG(v)) {
    Variant key = VARIANT_ARG_GET(k);
    Variant *value = (Variant*)gdextension_dictionary_operator_index((GDExtensionTypePtr)&dict, (GDExtensionVariantPtr)&key);
    *value = VARIANT_ARG_GET(v);
};
EXPORT void gd_editor_add_documentation(const char* xml, uint32_t len) {
    gdextension_editor_help_load_xml_from_utf8_chars_and_len(xml, len);
};
EXPORT void gd_editor_add_plugin(uintptr_t class_name) {
    gdextension_editor_add_plugin((GDExtensionConstStringNamePtr)&class_name);
};
EXPORT void gd_editor_end_plugin(uintptr_t class_name) {
    gdextension_editor_remove_plugin((GDExtensionConstStringNamePtr)&class_name);
};
EXPORT void gd_iterator_make(VARIANT_ARG(v), UnsafePointer result, UnsafePointer err) {
    Variant self = VARIANT_ARG_GET(v);
    GDExtensionBool valid = true;
    valid = valid && gdextension_variant_iter_init(&self, (void*)result, &valid);
    if (valid) return;
    ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
};
EXPORT bool gd_iterator_next(VARIANT_ARG(v), UnsafePointer iter, UnsafePointer err) {
    Variant self = VARIANT_ARG_GET(v);
    GDExtensionBool valid = false;
    GDExtensionBool ok = gdextension_variant_iter_next(&self, (GDExtensionVariantPtr)iter, &valid);
    if (ok) {
        return true;
    }
    ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return false;
};
EXPORT void gd_iterator_load(VARIANT_ARG(v), VARIANT_ARG(i), UnsafePointer result, UnsafePointer err) {
    Variant self = VARIANT_ARG_GET(v);
    uint64_t iter[3] = {i_1, i_2, i_3};
    void *points[16]; prepare_variants(&points[0], 1, (ANY)result);
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
        buf = (char *)gdextension_mem_alloc(len + 1); // +1 for null-terminator
    }
    memcpy(buf, str, len);
    buf[len] = '\0'; // null-terminate the string
    return buf;
}
EXPORT void gd_log(LogLevel level,
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
EXPORT UnsafePointer gd_memory_malloc(Int size) {
    return (UnsafePointer)gdextension_mem_alloc(size);
};
EXPORT Int gd_memory_sizeof(StringName name) {
    return gdextension_get_native_struct_size((GDExtensionConstStringNamePtr)&name);
};
EXPORT UnsafePointer gd_memory_resize(UnsafePointer addr, Int size) {
    return (UnsafePointer)gdextension_mem_realloc((void *)addr, size);
};
EXPORT void gd_memory_free(UnsafePointer addr) {
    gdextension_mem_free((void *)addr);
};
EXPORT void gd_memory_clear(UnsafePointer addr, Int size) {
    if (size <= 0) return;
    memset((void *)addr, 0, size);
};
EXPORT void gd_memory_edit_byte(UnsafePointer addr, uint8_t b) { *(uint8_t *)addr = b; };
EXPORT void gd_memory_edit_u16(UnsafePointer addr, uint16_t b) { *(uint16_t *)addr = b; };
EXPORT void gd_memory_edit_u32(UnsafePointer addr, uint32_t b) { *(uint32_t *)addr = b; };
EXPORT void gd_memory_edit_u64(UnsafePointer addr, uint64_t b) { *(uint64_t *)addr = b; };
EXPORT void gd_memory_edit_128(UnsafePointer addr, uint64_t a, uint64_t b) {
    uint64_t *ptr = (uint64_t *)addr;
    ptr[0] = a; ptr[1] = b;
};
EXPORT void gd_memory_edit_256(UnsafePointer addr, uint64_t a, uint64_t b, uint64_t c, uint64_t d) {
    uint64_t *ptr = (uint64_t *)addr;
    ptr[0] = a; ptr[1] = b; ptr[2] = c; ptr[3] = d;
};
EXPORT void gd_memory_edit_512(UnsafePointer addr, uint64_t a, uint64_t b, uint64_t c, uint64_t d, uint64_t e, uint64_t f, uint64_t g, uint64_t h) {
    uint64_t *ptr = (uint64_t *)addr;
    ptr[0] = a; ptr[1] = b; ptr[2] = c; ptr[3] = d;
    ptr[4] = e; ptr[5] = f; ptr[6] = g; ptr[7] = h;
};
EXPORT uint8_t gd_memory_load_byte(UnsafePointer addr) { return *(uint8_t *)addr; };
EXPORT uint16_t gd_memory_load_u16(UnsafePointer addr) { return *(uint16_t *)addr; };
EXPORT uint32_t gd_memory_load_u32(UnsafePointer addr) { return *(uint32_t *)addr; };
EXPORT uint64_t gd_memory_load_u64(UnsafePointer addr) { return *(uint64_t *)addr; };
EXPORT uintptr_t gd_object_make(uintptr_t name) {
    return (uintptr_t)gdextension_classdb_construct_object2((GDExtensionConstStringNamePtr)&name);
};
EXPORT void gd_object_call(Object obj, MethodForClass fn, Variant* result, Int argc, Variant args[], CallError* err) {
    void *points[16]; prepare_variants(&points[0], (uint32_t)argc, (ANY)args);
    gdextension_object_method_bind_call((GDExtensionMethodBindPtr)fn, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};
EXPORT result_24 gd_object_call_24(uintptr_t obj, uintptr_t fn, Int argc, ANY args, ANY err) {
    void *points[16]; prepare_variants(&points[0], argc, args);
    result_24 result = {};
    gdextension_object_method_bind_call((GDExtensionMethodBindPtr)fn, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)&result, (GDExtensionCallError*)err);
    return result;
};
EXPORT void gd_object_unsafe_free(uintptr_t obj) {
    gdextension_object_destroy((GDExtensionObjectPtr)obj);
};
EXPORT uintptr_t gd_object_name(uintptr_t obj) {
    uintptr_t name = 0;
    gdextension_object_get_class_name((GDExtensionObjectPtr)obj, gd_library, (GDExtensionUninitializedStringNamePtr)&name);
    return name;
};
EXPORT uintptr_t gd_object_type(uintptr_t name) {
    return (uintptr_t)gdextension_classdb_get_class_tag((GDExtensionConstStringNamePtr)&name);
};
EXPORT uintptr_t gd_object_cast(uintptr_t obj, uintptr_t tag) {
    return (uintptr_t)gdextension_object_cast_to((GDExtensionObjectPtr)obj, (void *)tag);
};
EXPORT Object gd_object_lookup(ObjectID id) {
    return (Object)gdextension_object_get_instance_from_id((GDObjectInstanceID)id);
};
EXPORT uintptr_t gd_object_global(uintptr_t name) {
    return (uintptr_t)gdextension_global_get_singleton((GDExtensionConstStringNamePtr)&name);
};
EXPORT void gd_object_extension_setup(uintptr_t obj, uintptr_t name, uintptr_t instance) {
    gdextension_object_set_instance((GDExtensionObjectPtr)obj, (GDExtensionConstStringNamePtr)&name, (GDExtensionClassInstancePtr)instance);
    gdextension_object_set_instance_binding((GDExtensionObjectPtr)obj, gd_library, (void *)instance, &instance_binding_callbacks);
};
EXPORT uintptr_t gd_object_extension_fetch(uintptr_t obj) {
    return (uintptr_t)gdextension_object_get_instance_binding((GDExtensionObjectPtr)obj, gd_library, &instance_binding_callbacks);
};
EXPORT void gd_object_extension_close(uintptr_t obj) {
    gdextension_object_free_instance_binding((GDExtensionObjectPtr)obj, gd_library);
};
EXPORT ObjectID gd_object_id(Object obj) {
    return gdextension_object_get_instance_id((GDExtensionObjectPtr)obj);
};
EXPORT ObjectID gd_object_id_inside_variant(VARIANT_ARG(v)) {
    Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_get_object_instance_id(&self);
};
EXPORT MethodForClass gd_object_method_lookup(StringName class_name, StringName method, int64_t hash) {
    return (MethodForClass)gdextension_classdb_get_method_bind((GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&method, hash);
};
EXPORT void gd_object_shaped_call(Object obj, MethodForClass fn, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)fn, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)result);
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
        prepare_callframe(1, &points[0], e->shape, (ANY)e->args);
        gdextension_object_method_bind_ptrcall(
            (GDExtensionMethodBindPtr)e->method,
            (GDExtensionObjectPtr)e->object,
            (const GDExtensionConstTypePtr*)&points[0],
            (GDExtensionTypePtr)e->result
        );
    }
    *crash_index = 0xFFFFFFFF;
}
EXPORT uint64_t gd_object_unsafe_call_8(uintptr_t obj, uintptr_t method, UINT64(shape), ANY args) {
    void *points[16]; prepare_callframe(1, &points[0], UINT64_FROM(shape), args);
    uint64_t result = 0;
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)&result);
    return result;
};
EXPORT result_16 gd_object_unsafe_call_16(uintptr_t obj, uintptr_t method, UINT64(shape), ANY args) {
    void *points[16]; prepare_callframe(1, &points[0], UINT64_FROM(shape), args);
    result_16 result = {};
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)&result);
    return result;
};
EXPORT result_32 gd_object_unsafe_call_32(uintptr_t obj, uintptr_t method, UINT64(shape), ANY args) {
    void *points[16]; prepare_callframe(1, &points[0], UINT64_FROM(shape), args);
    result_32 result = {};
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)&result);
    return result;
};
EXPORT result_64 gd_object_unsafe_call_64(uintptr_t obj, uintptr_t method, UINT64(shape), ANY args) {
    void *points[16]; prepare_callframe(1, &points[0], UINT64_FROM(shape), args);
    result_64 result = {};
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)&result);
    return result;
};
EXPORT uintptr_t gd_ptrcall_fn_addr() {
    return (uintptr_t)gdextension_object_method_bind_ptrcall;
}
static GDExtensionBool extension_script_categorization(GDExtensionScriptInstanceDataPtr instance, GDExtensionPropertyInfo *info) {
    property_list list = { .info = info };
    return gd_on_extension_script_categorization((uintptr_t)instance, (uintptr_t)&list);
}
static GDExtensionObjectPtr extension_script_get_owner(GDExtensionScriptInstanceDataPtr instance) {
    return (GDExtensionObjectPtr)gd_on_extension_script_get_owner((uintptr_t)instance);
}
EXPORT void gd_object_script_property_state_add(FunctionID fn, uintptr_t arg, StringName name, VARIANT_ARG(state)) {
    uint64_t v[3] = {state_1, state_2, state_3};
    ((GDExtensionScriptInstancePropertyStateAdd)fn)(&name, &v, (void *)arg);
}
static void extension_script_get_property_state(GDExtensionScriptInstanceDataPtr p_instance, GDExtensionScriptInstancePropertyStateAdd p_add_func, void *p_userdata) {
    gd_on_extension_script_get_property_state((uintptr_t)p_instance, (uintptr_t)p_add_func, (uintptr_t)p_userdata);
}
static GDExtensionBool extension_script_has_method(GDExtensionScriptInstanceDataPtr instance, GDExtensionConstStringNamePtr method) {
    return gd_on_extension_script_has_method((uintptr_t)instance, *(uintptr_t*)method);
}

static GDExtensionInt extension_script_get_method_argument_count(GDExtensionScriptInstanceDataPtr instance, GDExtensionConstStringNamePtr method, GDExtensionBool *valid) {
    *valid = 1;
    return gd_on_extension_script_get_method_argument_count((uintptr_t)instance, *(uintptr_t*)method);
}
static GDExtensionObjectPtr extension_script_get(GDExtensionScriptInstanceDataPtr instance) {
    return (GDExtensionObjectPtr)gd_on_extension_script_get((uintptr_t)instance);
}
static GDExtensionBool extension_script_is_placeholder(GDExtensionScriptInstanceDataPtr instance) {
    return gd_on_extension_script_is_placeholder((uintptr_t)instance);
}
static GDExtensionObjectPtr extension_script_get_language(GDExtensionScriptInstanceDataPtr instance) {
    return (GDExtensionObjectPtr)gd_on_extension_script_get_language((uintptr_t)instance);
}
uintptr_t gd_object_script_make(uintptr_t instance) {
    GDExtensionScriptInstanceInfo3 info = {
        .set_func = extension_instance_set,
        .get_func = extension_instance_get,
        .get_property_list_func = extension_instance_property_list,
        .free_property_list_func = class_free_property_list_func,
        .get_class_category_func = extension_script_categorization,
        .property_can_revert_func = extension_instance_property_has_default,
        .property_get_revert_func = extension_instance_property_get_default,
        .get_owner_func = extension_script_get_owner,
        .get_property_state_func = extension_script_get_property_state,
        .validate_property_func = extension_instance_property_validation,
        .has_method_func = extension_script_has_method,
        .get_method_argument_count_func = extension_script_get_method_argument_count,
        .call_func = (GDExtensionScriptInstanceCall)extension_instance_dynamic_call,
        .notification_func = (GDExtensionClassNotification2)gd_on_extension_instance_notification,
        .to_string_func = extension_instance_stringify,
        //.refcount_incremented_func = class_reference,
        //.refcount_decremented_func = class_unreference,
        .get_script_func = extension_script_get,
        .is_placeholder_func = extension_script_is_placeholder,
        .get_language_func = extension_script_get_language,
        .free_func = (GDExtensionScriptInstanceFree)gd_on_extension_instance_free,
    };
    return (uintptr_t)gdextension_script_instance_create3(&info, (GDExtensionScriptInstanceDataPtr)&instance);
};
EXPORT void gd_object_script_call(Object obj, StringName method, Variant* result, Int argc, Variant args[], CallError* err) {
    void *points[16]; points[1] = 0; prepare_variants(&points[0], (uint32_t)argc, (ANY)args);
    gdextension_object_call_script_method((GDExtensionObjectPtr)obj, (GDExtensionConstStringNamePtr)&method, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};
EXPORT void gd_object_script_setup(uintptr_t obj, uintptr_t instance) {
    gdextension_object_set_script_instance((GDExtensionScriptInstanceDataPtr)obj, (GDExtensionClassInstancePtr)instance);
};
EXPORT uintptr_t gd_object_script_fetch(uintptr_t obj, uintptr_t language) {
    return (uintptr_t)gdextension_object_get_script_instance((GDExtensionScriptInstanceDataPtr)obj, (GDExtensionClassInstancePtr)language);
};
EXPORT bool gd_object_script_defines_method(uintptr_t obj, uintptr_t method) {
    return gdextension_object_has_script_method((GDExtensionScriptInstanceDataPtr)obj, (GDExtensionConstStringNamePtr)&method);
};
EXPORT uintptr_t gd_object_script_placeholder_create(uintptr_t language, uintptr_t script, uintptr_t owner) {
    return (uintptr_t)gdextension_placeholder_script_instance_create((GDExtensionScriptInstanceDataPtr)language, (GDExtensionObjectPtr)script, (GDExtensionObjectPtr)owner);
};
EXPORT void gd_object_script_placeholder_update(uintptr_t p_placeholder, uintptr_t p_properties, uintptr_t p_values) {
    gdextension_placeholder_script_instance_update((GDExtensionScriptInstanceDataPtr)p_placeholder, (GDExtensionConstTypePtr)&p_properties, (GDExtensionConstTypePtr)&p_values);
};
EXPORT UnsafePointer gd_packed_array_access(uint32_t type, PACKED_ARRAY_ARG(pa), Int i) {
    PackedArray pa = PACKED_ARRAY_ARG_GET(pa);
    switch (type) {
    case 29: return (UnsafePointer)gdextension_packed_byte_array_operator_index_const(&pa, i);
    case 30: return (UnsafePointer)gdextension_packed_int32_array_operator_index_const(&pa, i);
    case 31: return (UnsafePointer)gdextension_packed_int64_array_operator_index_const(&pa, i);
    case 32: return (UnsafePointer)gdextension_packed_float32_array_operator_index_const(&pa, i);
    case 33: return (UnsafePointer)gdextension_packed_float64_array_operator_index_const(&pa, i);
    case 34: return (UnsafePointer)gdextension_packed_string_array_operator_index_const(&pa, i);
    case 35: return (UnsafePointer)gdextension_packed_vector2_array_operator_index_const(&pa, i);
    case 36: return (UnsafePointer)gdextension_packed_vector3_array_operator_index_const(&pa, i);
    case 37: return (UnsafePointer)gdextension_packed_color_array_operator_index_const(&pa, i);
    case 38: return (UnsafePointer)gdextension_packed_vector4_array_operator_index_const(&pa, i);
    }
    return 0;
};
EXPORT UnsafePointer gd_packed_array_modify(uint32_t type, PACKED_ARRAY_ARG(pa), Int i) {
    PackedArray pa = PACKED_ARRAY_ARG_GET(pa);
    switch (type) {
    case 29: return (UnsafePointer)gdextension_packed_byte_array_operator_index(&pa, i);
    case 30: return (UnsafePointer)gdextension_packed_int32_array_operator_index(&pa, i);
    case 31: return (UnsafePointer)gdextension_packed_int64_array_operator_index(&pa, i);
    case 32: return (UnsafePointer)gdextension_packed_float32_array_operator_index(&pa, i);
    case 33: return (UnsafePointer)gdextension_packed_float64_array_operator_index(&pa, i);
    case 34: return (UnsafePointer)gdextension_packed_string_array_operator_index(&pa, i);
    case 35: return (UnsafePointer)gdextension_packed_vector2_array_operator_index(&pa, i);
    case 36: return (UnsafePointer)gdextension_packed_vector3_array_operator_index(&pa, i);
    case 37: return (UnsafePointer)gdextension_packed_color_array_operator_index(&pa, i);
    case 38: return (UnsafePointer)gdextension_packed_vector4_array_operator_index(&pa, i);
    }
    return 0;
};
EXPORT void gd_array_set(Array a, Int i, VARIANT_ARG(v)) {
    *((Variant*)gdextension_array_operator_index(&a, i)) = VARIANT_ARG_GET(v);
};
EXPORT void gd_array_get(Array a, Int i, Variant* result) {
    *result = *(Variant*)gdextension_array_operator_index_const(&a, i);
};
EXPORT uintptr_t gd_ref_get_object(uintptr_t ref) {
    return (uintptr_t)gdextension_ref_get_object((GDExtensionObjectPtr)ref);
};
EXPORT void gd_ref_set_object(uintptr_t ref, uintptr_t obj) {
    gdextension_ref_set_object((GDExtensionObjectPtr)ref, (GDExtensionObjectPtr)obj);
};
EXPORT int32_t gd_string_access(String s, Int i) {
    return *(int32_t *)gdextension_string_operator_index_const((GDExtensionStringPtr)s, i);
};
EXPORT String gd_string_resize(String s, Int size) {
    String ptr = s;
    gdextension_string_resize((GDExtensionStringPtr)&ptr, size);
    return ptr;
};
EXPORT UnsafePointer gd_string_unsafe(String s) {
    return (UnsafePointer)gdextension_string_operator_index((GDExtensionStringPtr)s, 0);
};
EXPORT String gd_string_append(String s, String other) {
    String ptr = s;
    gdextension_string_operator_plus_eq_string((GDExtensionStringPtr)&ptr, (GDExtensionConstStringPtr)&other);
    return ptr;
};
EXPORT String gd_string_append_rune(String s, int32_t rune) {
    String ptr = s;
    gdextension_string_operator_plus_eq_char((GDExtensionStringPtr)&ptr, rune);
    return ptr;
};
EXPORT String gd_string_decode(uint8_t enc, const char* s, Int len) {
    String ptr = 0;
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
EXPORT Int gd_string_encode(uint8_t enc, String s, char* addr, Int len) {
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
EXPORT StringName gd_string_intern(uint8_t enc, const char* s, Int len) {
    StringName ptr = 0;
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
EXPORT String gd_variant_type_name(uint32_t vtype) {
    String name = 0;
    gdextension_variant_get_type_name((GDExtensionVariantType)vtype, &name);
    return name;
};
EXPORT void gd_variant_unsafe_make_native(uint32_t vtype, VARIANT_ARG(v), uint64_t shape, UnsafePointer result) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)result);
    uint64_t self[4] = {v_1, v_2, v_3};
    type_from_variant_constructors[vtype](points[0], &self[0]);
};
EXPORT void gd_variant_unsafe_from_native(uint32_t vtype, Variant* result, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    variant_from_type_constructors[vtype]((GDExtensionUninitializedVariantPtr)result, points[0]);
};
EXPORT void gd_variant_type_call(uint32_t vtype, uintptr_t name, Variant* result, Int argc, Variant args[], CallError* err) {
    void *points[16]; prepare_variants(&points[0], argc, (ANY)args);
    gdextension_variant_call_static((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};
EXPORT bool gd_variant_type_convertable(uint32_t a, uint32_t b, bool strict) {
    if (strict) {
        return gdextension_variant_can_convert_strict((GDExtensionVariantType)a, (GDExtensionVariantType)b);
    } else {
        return gdextension_variant_can_convert((GDExtensionVariantType)a, (GDExtensionVariantType)b);
    }
};
EXPORT void gd_variant_type_setup_array(uintptr_t array, uint32_t vtype, uintptr_t class_name, uint64_t v_1, uint64_t v_2, uint64_t v_3) {
    uint64_t script[3] = {v_1, v_2, v_3};
    gdextension_array_set_typed(&array, (GDExtensionVariantType)vtype, &class_name, &script[0]);
};
EXPORT void gd_variant_type_setup_dictionary(uintptr_t dict, uint32_t ktype, uintptr_t kclass_name, uint64_t key_script_1, uint64_t key_script_2, uint64_t key_script_3, uint32_t vtype, uintptr_t vclass_name, uint64_t val_script_1, uint64_t val_script_2, uint64_t val_script_3) {
    uint64_t k_script[3] = {key_script_1, key_script_2, key_script_3};
    uint64_t v_script[3] = {val_script_1, val_script_2, val_script_3};
    gdextension_dictionary_set_typed(&dict, (GDExtensionVariantType)ktype, &kclass_name, &k_script[0], (GDExtensionVariantType)vtype, &vclass_name, &v_script[0]);
};
EXPORT void gd_variant_type_fetch_constant(uint32_t vtype, uintptr_t name, UnsafePointer result) {
    void *points[16]; prepare_variants(&points[0], 1, (ANY)result);
    gdextension_variant_get_constant_value((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name, (GDExtensionTypePtr)points[0]);
};
EXPORT uintptr_t gd_variant_type_builtin_method(uint32_t vtype, uintptr_t name, int64_t hash) {
    return (uintptr_t)gdextension_variant_get_ptr_builtin_method((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name, hash);
};
EXPORT uintptr_t gd_variant_type_unsafe_constructor(uint32_t vtype, Int n) {
    return (uintptr_t)gdextension_variant_get_ptr_constructor((GDExtensionVariantType)vtype, n);
};
EXPORT void gd_variant_type_unsafe_free(uint32_t vtype, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    variant_ptr_destructors[vtype](points[0]);
};
EXPORT void gd_variant_zero(Variant* result) {
    gdextension_variant_new_nil((GDExtensionUninitializedVariantPtr)result);
};
EXPORT void gd_variant_copy(VARIANT_ARG(v), Variant* result) {
    Variant self = VARIANT_ARG_GET(v);
    gdextension_variant_new_copy((GDExtensionUninitializedVariantPtr)result, &self);
};
EXPORT void gd_variant_call(VARIANT_ARG(v), uintptr_t name, Variant* result, Int argc, Variant args[], CallError* err) {
    Variant self = VARIANT_ARG_GET(v);
    void *points[16]; prepare_variants(&points[0], argc, (ANY)args);
    gdextension_variant_call(&self, (GDExtensionConstStringNamePtr)&name, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};
EXPORT bool gd_variant_eval(uint32_t op, VARIANT_ARG(a), VARIANT_ARG(b), Variant* result) {
    uint64_t a[3] = {a_1, a_2, a_3};
    uint64_t b[3] = {b_1, b_2, b_3};
    GDExtensionBool valid = false;
    gdextension_variant_evaluate((GDExtensionVariantOperator)op, &a[0], &b[0], (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};
EXPORT Int gd_variant_hash(VARIANT_ARG(v)) {
    Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_hash(&self);
};
EXPORT bool gd_variant_bool(VARIANT_ARG(v)) {
    Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_booleanize(&self);
};
EXPORT uintptr_t gd_variant_text(VARIANT_ARG(v)) {
    Variant self = VARIANT_ARG_GET(v);
    uintptr_t text = 0;
    gdextension_variant_stringify(&self, &text);
    return text;
};
EXPORT uint32_t gd_variant_type(VARIANT_ARG(v)) {
    Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_get_type(&self);
};
EXPORT void gd_variant_deep_copy(VARIANT_ARG(v), Variant* result) {
    Variant self = VARIANT_ARG_GET(v);
    gdextension_variant_duplicate(&self, (GDExtensionVariantPtr)result, true);
};
EXPORT Int gd_variant_deep_hash(VARIANT_ARG(v), Int recursion) {
    Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_recursive_hash(&self, recursion);
};
EXPORT bool gd_variant_get_index(VARIANT_ARG(v), VARIANT_ARG(key), Variant* result) {
    Variant self = VARIANT_ARG_GET(v);
    Variant index = VARIANT_ARG_GET(key);
    GDExtensionBool valid = false;
    gdextension_variant_get_keyed(&self, &index, (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};
EXPORT bool gd_variant_get_array(VARIANT_ARG(v), Int i, Variant* result, CallError* err) {
    Variant self = VARIANT_ARG_GET(v);
    GDExtensionBool valid = false;
    GDExtensionBool oob = false;
    gdextension_variant_get_indexed(&self, i, (GDExtensionUninitializedVariantPtr)result, &valid, &oob);
    if (oob) ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return valid;
};
EXPORT bool gd_variant_get_field(VARIANT_ARG(v), uintptr_t name, Variant* result) {
    Variant self = VARIANT_ARG_GET(v);
    GDExtensionBool valid = false;
    gdextension_variant_get_named(&self, (GDExtensionConstStringNamePtr)&name, (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};
EXPORT bool gd_variant_type_has_property(uint32_t vtype, uintptr_t name) {
    return gdextension_variant_has_member((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name);
};
EXPORT bool gd_variant_has_index(VARIANT_ARG(v), VARIANT_ARG(idx)) {
    Variant self = VARIANT_ARG_GET(v);
    Variant index = VARIANT_ARG_GET(idx);
    GDExtensionBool valid = false;
    gdextension_variant_has_key(&self, &index, &valid);
    return valid;
};
EXPORT bool gd_variant_has_method(VARIANT_ARG(v), uintptr_t name) {
    Variant self = VARIANT_ARG_GET(v);
    return gdextension_variant_has_method(&self, (GDExtensionConstStringNamePtr)&name);
};
EXPORT bool gd_variant_set_index(VARIANT_ARG(v), VARIANT_ARG(key), VARIANT_ARG(val)) {
    Variant self = VARIANT_ARG_GET(v);
    Variant index = VARIANT_ARG_GET(key);
    Variant value = VARIANT_ARG_GET(val);
    GDExtensionBool valid = false;
    gdextension_variant_set_keyed(&self, &index, &value, &valid);
    return valid;
};
EXPORT bool gd_variant_set_array(VARIANT_ARG(v), Int i, VARIANT_ARG(val), UnsafePointer err) {
    Variant self = VARIANT_ARG_GET(v);
    Variant value = VARIANT_ARG_GET(val);
    GDExtensionBool valid = false;
    GDExtensionBool oob = false;
    gdextension_variant_set_indexed(&self, i, &value, &valid, &oob);
    if (oob) ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return valid;
};
EXPORT bool gd_variant_set_field(VARIANT_ARG(v), uintptr_t name, VARIANT_ARG(val)) {
    Variant self = VARIANT_ARG_GET(v);
    Variant value = VARIANT_ARG_GET(val);
    GDExtensionBool valid = false;
    gdextension_variant_set_named(&self, (GDExtensionConstStringNamePtr)&name, &value, &valid);
    return valid;
};
EXPORT void gd_variant_unsafe_eval(uintptr_t fn, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    ((GDExtensionPtrOperatorEvaluator)fn)(points[0], points[1], (GDExtensionTypePtr)result);
};
EXPORT void gd_variant_unsafe_call(uintptr_t fn, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; uint8_t argc = prepare_callframe(1, &points[0], shape, (ANY)args);
    ((GDExtensionPtrBuiltInMethod)fn)((GDExtensionTypePtr*)points[0], (const GDExtensionConstTypePtr*)&points[1], (GDExtensionTypePtr)result, argc);
};
EXPORT void gd_variant_unsafe_free(VARIANT_ARG(v)) {
    Variant self = VARIANT_ARG_GET(v);
    gdextension_variant_destroy(&self);
};
EXPORT uintptr_t gd_variant_unsafe_internal_pointer(uint32_t vtype, VARIANT_ARG(v)) {
    Variant self = VARIANT_ARG_GET(v);
    return (uintptr_t)variant_internal_ptr_funcs[vtype](&self);
};
EXPORT void gd_variant_unsafe_get_field(uintptr_t getter, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    ((GDExtensionPtrGetter)getter)(points[0], (void*)&result);
};
EXPORT void gd_variant_unsafe_get_array(uint32_t vtype, Int i, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    variant_ptr_indexed_getters[vtype]((GDExtensionConstTypePtr)result, i, points[0]);
};
EXPORT void gd_variant_unsafe_get_index(uint32_t vtype, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    variant_ptr_keyed_getters[vtype](points[0], points[1], (GDExtensionTypePtr)result);
};
EXPORT void gd_variant_unsafe_set_field(uintptr_t setter, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    ((GDExtensionPtrSetter)setter)(points[1], points[2]);
};
EXPORT void gd_variant_unsafe_set_array(uint32_t vtype, Int i, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    variant_ptr_indexed_setters[vtype](points[0], i, points[1]);
};
EXPORT void gd_variant_unsafe_set_index(uint32_t vtype, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    variant_ptr_keyed_setters[vtype](points[0], points[1], points[2]);
};
EXPORT uintptr_t gd_variant_type_setter(uint32_t vtype, uintptr_t name) {
    return (uintptr_t)gdextension_variant_get_ptr_setter((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name);
}
EXPORT uintptr_t gd_variant_type_getter(uint32_t vtype, uintptr_t name) {
    return (uintptr_t)gdextension_variant_get_ptr_getter((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name);
}
EXPORT void gd_variant_type_unsafe_make(uintptr_t fn, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; prepare_callframe(1, &points[0], shape, (ANY)args);
    ((GDExtensionPtrConstructor)fn)((GDExtensionUninitializedTypePtr)result, (const GDExtensionConstTypePtr *)&points[0]);
}
EXPORT void gd_variant_type_make(uint32_t vtype, Variant* result, Int argc, Variant args[], CallError* err) {
    void *points[16]; prepare_variants(&points[0], argc, (ANY)args);
    gdextension_variant_construct((GDExtensionVariantType)vtype, (GDExtensionUninitializedVariantPtr)result, (const GDExtensionConstVariantPtr *)&points[0], argc, (GDExtensionCallError*)err);
}
EXPORT void gd_variant_type_unsafe_call(UnsafePointer self, uintptr_t fn, UnsafePointer result, uint64_t shape, UnsafePointer args) {
    void *points[16]; uint8_t argc = prepare_callframe(2, &points[0], shape, (ANY)args);
    ((GDExtensionPtrBuiltInMethod)fn)((GDExtensionTypePtr*)self, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)result, argc);
}
EXPORT uintptr_t gd_variant_type_evaluator(uint32_t op, uint32_t a, uint32_t b) {
    return (uintptr_t)gdextension_variant_get_ptr_operator_evaluator((GDExtensionVariantOperator)op, (GDExtensionVariantType)a, (GDExtensionVariantType)b);
}

EXPORT uint32_t gd_version_major() {return gd_godot_version_cached.major;};
EXPORT uint32_t gd_version_minor() {return gd_godot_version_cached.minor;};
EXPORT uint32_t gd_version_patch() {return gd_godot_version_cached.patch;};
EXPORT uint32_t gd_version_hex() {return gd_godot_version_cached.hex;};
EXPORT uintptr_t gd_version_status() {
    uintptr_t s;
    gdextension_string_new_with_latin1_chars(&s, gd_godot_version_cached.status);
    return s;
};
EXPORT uintptr_t gd_version_build() {
    uintptr_t s;
    gdextension_string_new_with_latin1_chars(&s, gd_godot_version_cached.build);
    return s;
};
EXPORT uintptr_t gd_version_hash() {
    uintptr_t s;
    gdextension_string_new_with_latin1_chars(&s, gd_godot_version_cached.hash);
    return s;
};
EXPORT uint64_t gd_version_timestamp(void) { return gd_godot_version_cached.timestamp; };
EXPORT uintptr_t gd_version_string() {
    uintptr_t s;
    gdextension_string_new_with_latin1_chars(&s, gd_godot_version_cached.string);
    return s;
};

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

using namespace emscripten;
EMSCRIPTEN_BINDINGS(my_module) {
	function("gd_classdb_register", &gd_classdb_register, allow_raw_pointers());
	function("gd_classdb_register_methods", &gd_classdb_register_methods, allow_raw_pointers());
	function("gd_classdb_register_constant", &gd_classdb_register_constant, allow_raw_pointers());
	function("gd_classdb_register_property", &gd_classdb_register_property, allow_raw_pointers());
	function("gd_classdb_register_property_indexed", &gd_classdb_register_property_indexed, allow_raw_pointers());
	function("gd_classdb_register_property_group", &gd_classdb_register_property_group, allow_raw_pointers());
	function("gd_classdb_register_property_sub_group", &gd_classdb_register_property_sub_group, allow_raw_pointers());
	function("gd_classdb_register_signal", &gd_classdb_register_signal, allow_raw_pointers());
	function("gd_classdb_register_removal", &gd_classdb_register_removal, allow_raw_pointers());
}

#endif // __EMSCRIPTEN__
