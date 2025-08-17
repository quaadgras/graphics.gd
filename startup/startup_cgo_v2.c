#ifdef __EMSCRIPTEN__
    #include "core/extension/gdextension_interface.h"
    #include "gdextension_web_cgo_callbacks.cpp"
    #include <emscripten/bind.h>
    #include <bit>
#else
    #include "gdextension_interface.h"
    #include "_cgo_export.h"
#endif
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <stdio.h>

// emscripten can only expose std::string to JS, not char *
// we want emscripten to be able to deal with 64bit values without resorting to slow big ints.
#ifdef __EMSCRIPTEN__
    #define UINT64 double
    #define UINT64_MAKE(v) std::__bit_cast<double>(v)
    #define UINT64_FROM(v) std::__bit_cast<uint64_t>(v)
    #define INT64 double
    #define INT64_MAKE(v) std::__bit_cast<double>(v)
    #define INT64_FROM(v) std::__bit_cast<int64_t>(v)
    #define INT int32_t
    #define STRING std::string
    #define STRING_POINTER(s) (s).c_str()
    #define BUFFER std::string
    #define BUFFER_POINTER(s) (s).data()
    #define ANY uint32_t
#else
    #define UINT64 uint64_t
    #define UINT64_MAKE(v) (v)
    #define UINT64_FROM(v) (v)
    #define INT64 int64_t
    #define INT64_MAKE(v) (v)
    #define INT64_FROM(v) (v)
    #define INT int64_t
    #define STRING const char*
    #define STRING_POINTER(s) (s)
    #define BUFFER char*
    #define BUFFER_POINTER(s) (s)
    #define ANY void*
#endif

#define LOAD_PROC_ADDRESS(m_name, m_type) gdextension_##m_name = (m_type)p_get_proc_address(#m_name);

GDExtensionClassLibraryPtr cgo_library = NULL;
GDExtensionGodotVersion2 cgo_cached_godot_version;

typedef struct {
    GDExtensionVariantPtr r_return;
    GDExtensionCallError *r_error;
    const GDExtensionConstVariantPtr *p_args;
    GDExtensionInt p_argument_count;
} reverse_variant_frame;

typedef struct {
    GDExtensionConstTypePtr r_return;
    const GDExtensionConstTypePtr *p_args;
} reverse_unsafe_frame;

void cgo_initialize(void *ignore, GDExtensionInitializationLevel level) { go_on_init(level); }
void cgo_deinitialize(void *ignore, GDExtensionInitializationLevel level) { go_on_exit(level); }
void cgo_callable_call_func(void *callable_userdata, const GDExtensionConstVariantPtr *p_args, GDExtensionInt p_argument_count, GDExtensionVariantPtr r_return, GDExtensionCallError *r_error) {
    go_on_callable_call((uintptr_t)callable_userdata, r_return, p_argument_count, (void *)p_args, r_error);
}
GDExtensionBool cgo_callable_is_valid_func(void *callable_userdata) {return go_on_callable_validation((uintptr_t)callable_userdata);}
void cgo_callable_free_func(void *callable_userdata) { go_on_callable_free((uintptr_t)callable_userdata); }
uint32_t cgo_callable_hash_func(void *callable_userdata) {return go_on_callable_hash((uintptr_t)callable_userdata);}
GDExtensionBool cgo_callable_equal_func(void *callable_userdata, void *other_userdata) {return go_on_callable_compare((uintptr_t)callable_userdata, (uintptr_t)other_userdata);}
GDExtensionBool cgo_callable_less_than_func(void *callable_userdata, void *other_userdata) {return go_on_callable_less_than((uintptr_t)callable_userdata, (uintptr_t)other_userdata);}
void cgo_callable_to_string_func(void *callable_userdata, GDExtensionBool *r_is_valid, GDExtensionStringPtr r_out) {
    uint64_t invalid = 0;
    uintptr_t s = go_on_callable_stringify((uintptr_t)callable_userdata, &invalid);
    *r_is_valid = !(GDExtensionBool)invalid;
    if (!invalid) *((uintptr_t*)r_out) = s;
}
GDExtensionInt cgo_callable_get_argument_count_func(void *callable_userdata, GDExtensionBool *r_is_valid) {
    uint64_t invalid = 0;
    int64_t count = go_on_callable_get_argument_count((uintptr_t)callable_userdata, &invalid);
    *r_is_valid = !(GDExtensionBool)invalid;
    if (invalid) return -1;
    return (GDExtensionInt)count;
}

void cgo_method_call_func(void *method_userdata, GDExtensionClassInstancePtr p_instance, const GDExtensionConstVariantPtr *p_args, GDExtensionInt p_argument_count, GDExtensionVariantPtr r_return, GDExtensionCallError *r_error) {
    go_on_extension_instance_call((uintptr_t)p_instance, (uintptr_t)method_userdata, r_return, p_argument_count, (void *)p_args, r_error);
}
void cgo_method_ptrcall_func(void *method_userdata, GDExtensionClassInstancePtr p_instance, const GDExtensionConstTypePtr *p_args, GDExtensionTypePtr r_ret) {
    go_on_extension_instance_unsafe_call((uintptr_t)p_instance, (uintptr_t)method_userdata, r_ret, (void *)p_args);
}

void *cgo_instance_binding_create_func(void *p_token, void *p_instance) {
    return (void *)go_on_extension_binding_created((uintptr_t)p_instance);
}

void cgo_instance_binding_free_func(void *p_token, void *p_instance, void *p_binding) {
    go_on_extension_binding_removed((uintptr_t)p_instance, (uintptr_t)p_binding);
}

GDExtensionBool cgo_instance_binding_reference_func(void *p_token, void *p_binding, GDExtensionBool p_reference) {
    return go_on_extension_binding_reference((uintptr_t)p_binding, p_reference);
}

GDExtensionInstanceBindingCallbacks instance_binding_callbacks = {
    .create_callback = cgo_instance_binding_create_func,
    .free_callback = cgo_instance_binding_free_func,
    .reference_callback = cgo_instance_binding_reference_func
};

GDExtensionTypeFromVariantConstructorFunc type_from_variant_constructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionVariantFromTypeConstructorFunc variant_from_type_constructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrDestructor variant_ptr_destructors[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionVariantGetInternalPtrFunc variant_internal_ptr_funcs[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrIndexedSetter variant_ptr_indexed_setters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrIndexedGetter variant_ptr_indexed_getters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrKeyedSetter variant_ptr_keyed_setters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];
GDExtensionPtrKeyedGetter variant_ptr_keyed_getters[GDEXTENSION_VARIANT_TYPE_VARIANT_MAX];

GDExtensionInterfaceMemAlloc gdextension_mem_alloc = NULL;
GDExtensionInterfaceMemRealloc gdextension_mem_realloc = NULL;
GDExtensionInterfaceMemFree gdextension_mem_free = NULL;
GDExtensionInterfacePrintError gdextension_print_error = NULL;
GDExtensionInterfacePrintErrorWithMessage gdextension_print_error_with_message = NULL;
GDExtensionInterfacePrintWarning gdextension_print_warning = NULL;
GDExtensionInterfacePrintWarningWithMessage gdextension_print_warning_with_message = NULL;
GDExtensionInterfacePrintScriptError gdextension_print_script_error = NULL;
GDExtensionInterfacePrintScriptErrorWithMessage gdextension_print_script_error_with_message = NULL;
GDExtensionInterfaceGetNativeStructSize gdextension_get_native_struct_size = NULL;
GDExtensionInterfaceVariantNewCopy gdextension_variant_new_copy = NULL;
GDExtensionInterfaceVariantNewNil gdextension_variant_new_nil = NULL;
GDExtensionInterfaceVariantDestroy gdextension_variant_destroy = NULL;
GDExtensionInterfaceVariantCall gdextension_variant_call = NULL;
GDExtensionInterfaceVariantCallStatic gdextension_variant_call_static = NULL;
GDExtensionInterfaceVariantEvaluate gdextension_variant_evaluate = NULL;
GDExtensionInterfaceVariantSet gdextension_variant_set = NULL;
GDExtensionInterfaceVariantSetNamed gdextension_variant_set_named = NULL;
GDExtensionInterfaceVariantSetKeyed gdextension_variant_set_keyed = NULL;
GDExtensionInterfaceVariantSetIndexed gdextension_variant_set_indexed = NULL;
GDExtensionInterfaceVariantGet gdextension_variant_get = NULL;
GDExtensionInterfaceVariantGetNamed gdextension_variant_get_named = NULL;
GDExtensionInterfaceVariantGetKeyed gdextension_variant_get_keyed = NULL;
GDExtensionInterfaceVariantGetIndexed gdextension_variant_get_indexed = NULL;
GDExtensionInterfaceVariantIterInit gdextension_variant_iter_init = NULL;
GDExtensionInterfaceVariantIterNext gdextension_variant_iter_next = NULL;
GDExtensionInterfaceVariantIterGet gdextension_variant_iter_get = NULL;
GDExtensionInterfaceVariantHash gdextension_variant_hash = NULL;
GDExtensionInterfaceVariantRecursiveHash gdextension_variant_recursive_hash = NULL;
GDExtensionInterfaceVariantHashCompare gdextension_variant_hash_compare = NULL;
GDExtensionInterfaceVariantBooleanize gdextension_variant_booleanize = NULL;
GDExtensionInterfaceVariantDuplicate gdextension_variant_duplicate = NULL;
GDExtensionInterfaceVariantStringify gdextension_variant_stringify = NULL;
GDExtensionInterfaceVariantGetType gdextension_variant_get_type = NULL;
GDExtensionInterfaceVariantHasMethod gdextension_variant_has_method = NULL;
GDExtensionInterfaceVariantHasMember gdextension_variant_has_member = NULL;
GDExtensionInterfaceVariantHasKey gdextension_variant_has_key = NULL;
GDExtensionInterfaceVariantGetObjectInstanceId gdextension_variant_get_object_instance_id = NULL;
GDExtensionInterfaceVariantGetTypeName gdextension_variant_get_type_name = NULL;
GDExtensionInterfaceVariantCanConvert gdextension_variant_can_convert = NULL;
GDExtensionInterfaceVariantCanConvertStrict gdextension_variant_can_convert_strict = NULL;
GDExtensionInterfaceGetVariantFromTypeConstructor gdextension_get_variant_from_type_constructor = NULL;
GDExtensionInterfaceGetVariantToTypeConstructor gdextension_get_variant_to_type_constructor = NULL;
GDExtensionInterfaceGetVariantGetInternalPtrFunc gdextension_variant_get_ptr_internal_getter = NULL;
GDExtensionInterfaceVariantGetPtrOperatorEvaluator gdextension_variant_get_ptr_operator_evaluator = NULL;
GDExtensionInterfaceVariantGetPtrBuiltinMethod gdextension_variant_get_ptr_builtin_method = NULL;
GDExtensionInterfaceVariantGetPtrConstructor gdextension_variant_get_ptr_constructor = NULL;
GDExtensionInterfaceVariantGetPtrDestructor gdextension_variant_get_ptr_destructor = NULL;
GDExtensionInterfaceVariantConstruct gdextension_variant_construct = NULL;
GDExtensionInterfaceVariantGetPtrSetter gdextension_variant_get_ptr_setter = NULL;
GDExtensionInterfaceVariantGetPtrGetter gdextension_variant_get_ptr_getter = NULL;
GDExtensionInterfaceVariantGetPtrIndexedSetter gdextension_variant_get_ptr_indexed_setter = NULL;
GDExtensionInterfaceVariantGetPtrIndexedGetter gdextension_variant_get_ptr_indexed_getter = NULL;
GDExtensionInterfaceVariantGetPtrKeyedSetter gdextension_variant_get_ptr_keyed_setter = NULL;
GDExtensionInterfaceVariantGetPtrKeyedGetter gdextension_variant_get_ptr_keyed_getter = NULL;
GDExtensionInterfaceVariantGetPtrKeyedChecker gdextension_variant_get_ptr_keyed_checker = NULL;
GDExtensionInterfaceVariantGetConstantValue gdextension_variant_get_constant_value = NULL;
GDExtensionInterfaceVariantGetPtrUtilityFunction gdextension_variant_get_ptr_utility_function = NULL;
GDExtensionInterfaceStringNewWithLatin1Chars gdextension_string_new_with_latin1_chars = NULL;
GDExtensionInterfaceStringNewWithUtf8Chars gdextension_string_new_with_utf8_chars = NULL;
GDExtensionInterfaceStringNewWithUtf16Chars gdextension_string_new_with_utf16_chars = NULL;
GDExtensionInterfaceStringNewWithUtf32Chars gdextension_string_new_with_utf32_chars = NULL;
GDExtensionInterfaceStringNewWithWideChars gdextension_string_new_with_wide_chars = NULL;
GDExtensionInterfaceStringNewWithLatin1CharsAndLen gdextension_string_new_with_latin1_chars_and_len = NULL;
GDExtensionInterfaceStringNewWithUtf8CharsAndLen gdextension_string_new_with_utf8_chars_and_len = NULL;
GDExtensionInterfaceStringNewWithUtf8CharsAndLen2 gdextension_string_new_with_utf8_chars_and_len2 = NULL;
GDExtensionInterfaceStringNewWithUtf16CharsAndLen gdextension_string_new_with_utf16_chars_and_len = NULL;
GDExtensionInterfaceStringNewWithUtf16CharsAndLen2 gdextension_string_new_with_utf16_chars_and_len2 = NULL;
GDExtensionInterfaceStringNewWithUtf32CharsAndLen gdextension_string_new_with_utf32_chars_and_len = NULL;
GDExtensionInterfaceStringNewWithWideCharsAndLen gdextension_string_new_with_wide_chars_and_len = NULL;
GDExtensionInterfaceStringToLatin1Chars gdextension_string_to_latin1_chars = NULL;
GDExtensionInterfaceStringToUtf8Chars gdextension_string_to_utf8_chars = NULL;
GDExtensionInterfaceStringToUtf16Chars gdextension_string_to_utf16_chars = NULL;
GDExtensionInterfaceStringToUtf32Chars gdextension_string_to_utf32_chars = NULL;
GDExtensionInterfaceStringToWideChars gdextension_string_to_wide_chars = NULL;
GDExtensionInterfaceStringOperatorIndex gdextension_string_operator_index = NULL;
GDExtensionInterfaceStringOperatorIndexConst gdextension_string_operator_index_const = NULL;
GDExtensionInterfaceStringOperatorPlusEqString gdextension_string_operator_plus_eq_string = NULL;
GDExtensionInterfaceStringOperatorPlusEqChar gdextension_string_operator_plus_eq_char = NULL;
GDExtensionInterfaceStringOperatorPlusEqCstr gdextension_string_operator_plus_eq_cstr = NULL;
GDExtensionInterfaceStringOperatorPlusEqWcstr gdextension_string_operator_plus_eq_wcstr = NULL;
GDExtensionInterfaceStringOperatorPlusEqC32str gdextension_string_operator_plus_eq_c32str = NULL;
GDExtensionInterfaceStringResize gdextension_string_resize = NULL;
GDExtensionInterfaceStringNameNewWithLatin1Chars gdextension_string_name_new_with_latin1_chars = NULL;
GDExtensionInterfaceStringNameNewWithUtf8CharsAndLen gdextension_string_name_new_with_utf8_chars_and_len = NULL;
GDExtensionInterfaceXmlParserOpenBuffer gdextension_xml_parser_open_buffer = NULL;
GDExtensionInterfaceFileAccessStoreBuffer gdextension_file_access_store_buffer = NULL;
GDExtensionInterfaceFileAccessGetBuffer gdextension_file_access_get_buffer = NULL;
GDExtensionInterfaceWorkerThreadPoolAddNativeGroupTask gdextension_worker_thread_pool_add_native_group_task = NULL;
GDExtensionInterfaceWorkerThreadPoolAddNativeTask gdextension_worker_thread_pool_add_native_task = NULL;
GDExtensionInterfacePackedByteArrayOperatorIndex gdextension_packed_byte_array_operator_index = NULL;
GDExtensionInterfacePackedByteArrayOperatorIndexConst gdextension_packed_byte_array_operator_index_const = NULL;
GDExtensionInterfacePackedColorArrayOperatorIndex gdextension_packed_color_array_operator_index = NULL;
GDExtensionInterfacePackedColorArrayOperatorIndexConst gdextension_packed_color_array_operator_index_const = NULL;
GDExtensionInterfacePackedFloat32ArrayOperatorIndex gdextension_packed_float32_array_operator_index = NULL;
GDExtensionInterfacePackedFloat32ArrayOperatorIndexConst gdextension_packed_float32_array_operator_index_const = NULL;
GDExtensionInterfacePackedFloat64ArrayOperatorIndex gdextension_packed_float64_array_operator_index = NULL;
GDExtensionInterfacePackedFloat64ArrayOperatorIndexConst gdextension_packed_float64_array_operator_index_const = NULL;
GDExtensionInterfacePackedInt32ArrayOperatorIndex gdextension_packed_int32_array_operator_index = NULL;
GDExtensionInterfacePackedInt32ArrayOperatorIndexConst gdextension_packed_int32_array_operator_index_const = NULL;
GDExtensionInterfacePackedInt64ArrayOperatorIndex gdextension_packed_int64_array_operator_index = NULL;
GDExtensionInterfacePackedInt64ArrayOperatorIndexConst gdextension_packed_int64_array_operator_index_const = NULL;
GDExtensionInterfacePackedStringArrayOperatorIndex gdextension_packed_string_array_operator_index = NULL;
GDExtensionInterfacePackedStringArrayOperatorIndexConst gdextension_packed_string_array_operator_index_const = NULL;
GDExtensionInterfacePackedVector2ArrayOperatorIndex gdextension_packed_vector2_array_operator_index = NULL;
GDExtensionInterfacePackedVector2ArrayOperatorIndexConst gdextension_packed_vector2_array_operator_index_const = NULL;
GDExtensionInterfacePackedVector3ArrayOperatorIndex gdextension_packed_vector3_array_operator_index = NULL;
GDExtensionInterfacePackedVector3ArrayOperatorIndexConst gdextension_packed_vector3_array_operator_index_const = NULL;
GDExtensionInterfacePackedVector4ArrayOperatorIndex gdextension_packed_vector4_array_operator_index = NULL;
GDExtensionInterfacePackedVector4ArrayOperatorIndexConst gdextension_packed_vector4_array_operator_index_const = NULL;
GDExtensionInterfaceArrayOperatorIndex gdextension_array_operator_index = NULL;
GDExtensionInterfaceArrayOperatorIndexConst gdextension_array_operator_index_const = NULL;
GDExtensionInterfaceArraySetTyped gdextension_array_set_typed = NULL;
GDExtensionInterfaceDictionaryOperatorIndex gdextension_dictionary_operator_index = NULL;
GDExtensionInterfaceDictionaryOperatorIndexConst gdextension_dictionary_operator_index_const = NULL;
GDExtensionInterfaceDictionarySetTyped gdextension_dictionary_set_typed = NULL;
GDExtensionInterfaceObjectMethodBindCall gdextension_object_method_bind_call = NULL;
GDExtensionInterfaceObjectMethodBindPtrcall gdextension_object_method_bind_ptrcall = NULL;
GDExtensionInterfaceObjectDestroy gdextension_object_destroy = NULL;
GDExtensionInterfaceGlobalGetSingleton gdextension_global_get_singleton = NULL;
GDExtensionInterfaceObjectGetInstanceBinding gdextension_object_get_instance_binding = NULL;
GDExtensionInterfaceObjectSetInstanceBinding gdextension_object_set_instance_binding = NULL;
GDExtensionInterfaceObjectFreeInstanceBinding gdextension_object_free_instance_binding = NULL;
GDExtensionInterfaceObjectSetInstance gdextension_object_set_instance = NULL;
GDExtensionInterfaceObjectGetClassName gdextension_object_get_class_name = NULL;
GDExtensionInterfaceObjectCastTo gdextension_object_cast_to = NULL;
GDExtensionInterfaceObjectGetInstanceFromId gdextension_object_get_instance_from_id = NULL;
GDExtensionInterfaceObjectGetInstanceId gdextension_object_get_instance_id = NULL;
GDExtensionInterfaceObjectHasScriptMethod gdextension_object_has_script_method = NULL;
GDExtensionInterfaceObjectCallScriptMethod gdextension_object_call_script_method = NULL;
GDExtensionInterfaceCallableCustomCreate2 gdextension_callable_custom_create2 = NULL;
GDExtensionInterfaceCallableCustomGetUserData gdextension_callable_custom_get_userdata = NULL;
GDExtensionInterfaceRefGetObject gdextension_ref_get_object = NULL;
GDExtensionInterfaceRefSetObject gdextension_ref_set_object = NULL;
GDExtensionInterfaceScriptInstanceCreate3 gdextension_script_instance_create3 = NULL;
GDExtensionInterfacePlaceHolderScriptInstanceCreate gdextension_placeholder_script_instance_create = NULL;
GDExtensionInterfacePlaceHolderScriptInstanceUpdate gdextension_placeholder_script_instance_update = NULL;
GDExtensionInterfaceObjectGetScriptInstance gdextension_object_get_script_instance = NULL;
GDExtensionInterfaceObjectSetScriptInstance gdextension_object_set_script_instance = NULL;
GDExtensionInterfaceClassdbConstructObject2 gdextension_classdb_construct_object2 = NULL;
GDExtensionInterfaceClassdbGetMethodBind gdextension_classdb_get_method_bind = NULL;
GDExtensionInterfaceClassdbGetClassTag gdextension_classdb_get_class_tag = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClass4 gdextension_classdb_register_extension_class4 = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClassMethod gdextension_classdb_register_extension_class_method = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClassVirtualMethod gdextension_classdb_register_extension_class_virtual_method = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClassIntegerConstant gdextension_classdb_register_extension_class_integer_constant = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClassProperty gdextension_classdb_register_extension_class_property = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClassPropertyIndexed gdextension_classdb_register_extension_class_property_indexed = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClassPropertyGroup gdextension_classdb_register_extension_class_property_group = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClassPropertySubgroup gdextension_classdb_register_extension_class_property_subgroup = NULL;
GDExtensionInterfaceClassdbRegisterExtensionClassSignal gdextension_classdb_register_extension_class_signal = NULL;
GDExtensionInterfaceClassdbUnregisterExtensionClass gdextension_classdb_unregister_extension_class = NULL;
GDExtensionInterfaceGetLibraryPath gdextension_get_library_path = NULL;
GDExtensionInterfaceEditorAddPlugin gdextension_editor_add_plugin = NULL;
GDExtensionInterfaceEditorRemovePlugin gdextension_editor_remove_plugin = NULL;
GDExtensionInterfaceEditorRegisterGetClassesUsedCallback gdextension_editor_register_get_classes_used_callback = NULL;
GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8Chars gdextension_editor_help_load_xml_from_utf8_chars = NULL;
GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8CharsAndLen gdextension_editor_help_load_xml_from_utf8_chars_and_len = NULL;
GDExtensionInterfaceImagePtrw gdextension_image_ptrw = NULL;
GDExtensionInterfaceImagePtr gdextension_image_ptr = NULL;
GDExtensionInterfaceRegisterMainLoopCallbacks gdextension_register_main_loop_callbacks = NULL;



GDExtensionBool cgo_extension_init(GDExtensionInterfaceGetProcAddress p_get_proc_address, GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization) {
    LOAD_PROC_ADDRESS(mem_alloc, GDExtensionInterfaceMemAlloc);
	LOAD_PROC_ADDRESS(mem_realloc, GDExtensionInterfaceMemRealloc);
	LOAD_PROC_ADDRESS(mem_free, GDExtensionInterfaceMemFree);
	LOAD_PROC_ADDRESS(print_error_with_message, GDExtensionInterfacePrintErrorWithMessage);
	LOAD_PROC_ADDRESS(print_warning, GDExtensionInterfacePrintWarning);
	LOAD_PROC_ADDRESS(print_warning_with_message, GDExtensionInterfacePrintWarningWithMessage);
	LOAD_PROC_ADDRESS(print_script_error, GDExtensionInterfacePrintScriptError);
	LOAD_PROC_ADDRESS(print_script_error_with_message, GDExtensionInterfacePrintScriptErrorWithMessage);
	LOAD_PROC_ADDRESS(get_native_struct_size, GDExtensionInterfaceGetNativeStructSize);
	LOAD_PROC_ADDRESS(variant_new_copy, GDExtensionInterfaceVariantNewCopy);
	LOAD_PROC_ADDRESS(variant_new_nil, GDExtensionInterfaceVariantNewNil);
	LOAD_PROC_ADDRESS(variant_destroy, GDExtensionInterfaceVariantDestroy);
	LOAD_PROC_ADDRESS(variant_call, GDExtensionInterfaceVariantCall);
	LOAD_PROC_ADDRESS(variant_call_static, GDExtensionInterfaceVariantCallStatic);
	LOAD_PROC_ADDRESS(variant_evaluate, GDExtensionInterfaceVariantEvaluate);
	LOAD_PROC_ADDRESS(variant_set, GDExtensionInterfaceVariantSet);
	LOAD_PROC_ADDRESS(variant_set_named, GDExtensionInterfaceVariantSetNamed);
	LOAD_PROC_ADDRESS(variant_set_keyed, GDExtensionInterfaceVariantSetKeyed);
	LOAD_PROC_ADDRESS(variant_set_indexed, GDExtensionInterfaceVariantSetIndexed);
	LOAD_PROC_ADDRESS(variant_get, GDExtensionInterfaceVariantGet);
	LOAD_PROC_ADDRESS(variant_get_named, GDExtensionInterfaceVariantGetNamed);
	LOAD_PROC_ADDRESS(variant_get_keyed, GDExtensionInterfaceVariantGetKeyed);
	LOAD_PROC_ADDRESS(variant_get_indexed, GDExtensionInterfaceVariantGetIndexed);
	LOAD_PROC_ADDRESS(variant_iter_init, GDExtensionInterfaceVariantIterInit);
	LOAD_PROC_ADDRESS(variant_iter_next, GDExtensionInterfaceVariantIterNext);
	LOAD_PROC_ADDRESS(variant_iter_get, GDExtensionInterfaceVariantIterGet);
	LOAD_PROC_ADDRESS(variant_hash, GDExtensionInterfaceVariantHash);
	LOAD_PROC_ADDRESS(variant_recursive_hash, GDExtensionInterfaceVariantRecursiveHash);
	LOAD_PROC_ADDRESS(variant_hash_compare, GDExtensionInterfaceVariantHashCompare);
	LOAD_PROC_ADDRESS(variant_booleanize, GDExtensionInterfaceVariantBooleanize);
	LOAD_PROC_ADDRESS(variant_duplicate, GDExtensionInterfaceVariantDuplicate);
	LOAD_PROC_ADDRESS(variant_stringify, GDExtensionInterfaceVariantStringify);
	LOAD_PROC_ADDRESS(variant_get_type, GDExtensionInterfaceVariantGetType);
	LOAD_PROC_ADDRESS(variant_has_method, GDExtensionInterfaceVariantHasMethod);
	LOAD_PROC_ADDRESS(variant_has_member, GDExtensionInterfaceVariantHasMember);
	LOAD_PROC_ADDRESS(variant_has_key, GDExtensionInterfaceVariantHasKey);
	LOAD_PROC_ADDRESS(variant_get_object_instance_id, GDExtensionInterfaceVariantGetObjectInstanceId);
	LOAD_PROC_ADDRESS(variant_get_type_name, GDExtensionInterfaceVariantGetTypeName);
	LOAD_PROC_ADDRESS(variant_can_convert, GDExtensionInterfaceVariantCanConvert);
	LOAD_PROC_ADDRESS(variant_can_convert_strict, GDExtensionInterfaceVariantCanConvertStrict);
	LOAD_PROC_ADDRESS(get_variant_from_type_constructor, GDExtensionInterfaceGetVariantFromTypeConstructor);
	LOAD_PROC_ADDRESS(get_variant_to_type_constructor, GDExtensionInterfaceGetVariantToTypeConstructor);
	LOAD_PROC_ADDRESS(variant_get_ptr_internal_getter, GDExtensionInterfaceGetVariantGetInternalPtrFunc);
	LOAD_PROC_ADDRESS(variant_get_ptr_operator_evaluator, GDExtensionInterfaceVariantGetPtrOperatorEvaluator);
	LOAD_PROC_ADDRESS(variant_get_ptr_builtin_method, GDExtensionInterfaceVariantGetPtrBuiltinMethod);
	LOAD_PROC_ADDRESS(variant_get_ptr_constructor, GDExtensionInterfaceVariantGetPtrConstructor);
	LOAD_PROC_ADDRESS(variant_get_ptr_destructor, GDExtensionInterfaceVariantGetPtrDestructor);
	LOAD_PROC_ADDRESS(variant_construct, GDExtensionInterfaceVariantConstruct);
	LOAD_PROC_ADDRESS(variant_get_ptr_setter, GDExtensionInterfaceVariantGetPtrSetter);
	LOAD_PROC_ADDRESS(variant_get_ptr_getter, GDExtensionInterfaceVariantGetPtrGetter);
	LOAD_PROC_ADDRESS(variant_get_ptr_indexed_setter, GDExtensionInterfaceVariantGetPtrIndexedSetter);
	LOAD_PROC_ADDRESS(variant_get_ptr_indexed_getter, GDExtensionInterfaceVariantGetPtrIndexedGetter);
	LOAD_PROC_ADDRESS(variant_get_ptr_keyed_setter, GDExtensionInterfaceVariantGetPtrKeyedSetter);
	LOAD_PROC_ADDRESS(variant_get_ptr_keyed_getter, GDExtensionInterfaceVariantGetPtrKeyedGetter);
	LOAD_PROC_ADDRESS(variant_get_ptr_keyed_checker, GDExtensionInterfaceVariantGetPtrKeyedChecker);
	LOAD_PROC_ADDRESS(variant_get_constant_value, GDExtensionInterfaceVariantGetConstantValue);
	LOAD_PROC_ADDRESS(variant_get_ptr_utility_function, GDExtensionInterfaceVariantGetPtrUtilityFunction);
	LOAD_PROC_ADDRESS(string_new_with_latin1_chars, GDExtensionInterfaceStringNewWithLatin1Chars);
	LOAD_PROC_ADDRESS(string_new_with_utf8_chars, GDExtensionInterfaceStringNewWithUtf8Chars);
	LOAD_PROC_ADDRESS(string_new_with_utf16_chars, GDExtensionInterfaceStringNewWithUtf16Chars);
	LOAD_PROC_ADDRESS(string_new_with_utf32_chars, GDExtensionInterfaceStringNewWithUtf32Chars);
	LOAD_PROC_ADDRESS(string_new_with_wide_chars, GDExtensionInterfaceStringNewWithWideChars);
	LOAD_PROC_ADDRESS(string_new_with_latin1_chars_and_len, GDExtensionInterfaceStringNewWithLatin1CharsAndLen);
	LOAD_PROC_ADDRESS(string_new_with_utf8_chars_and_len, GDExtensionInterfaceStringNewWithUtf8CharsAndLen);
	LOAD_PROC_ADDRESS(string_new_with_utf8_chars_and_len2, GDExtensionInterfaceStringNewWithUtf8CharsAndLen2);
	LOAD_PROC_ADDRESS(string_new_with_utf16_chars_and_len, GDExtensionInterfaceStringNewWithUtf16CharsAndLen);
	LOAD_PROC_ADDRESS(string_new_with_utf16_chars_and_len2, GDExtensionInterfaceStringNewWithUtf16CharsAndLen2);
	LOAD_PROC_ADDRESS(string_new_with_utf32_chars_and_len, GDExtensionInterfaceStringNewWithUtf32CharsAndLen);
	LOAD_PROC_ADDRESS(string_new_with_wide_chars_and_len, GDExtensionInterfaceStringNewWithWideCharsAndLen);
	LOAD_PROC_ADDRESS(string_to_latin1_chars, GDExtensionInterfaceStringToLatin1Chars);
	LOAD_PROC_ADDRESS(string_to_utf8_chars, GDExtensionInterfaceStringToUtf8Chars);
	LOAD_PROC_ADDRESS(string_to_utf16_chars, GDExtensionInterfaceStringToUtf16Chars);
	LOAD_PROC_ADDRESS(string_to_utf32_chars, GDExtensionInterfaceStringToUtf32Chars);
	LOAD_PROC_ADDRESS(string_to_wide_chars, GDExtensionInterfaceStringToWideChars);
	LOAD_PROC_ADDRESS(string_operator_index, GDExtensionInterfaceStringOperatorIndex);
	LOAD_PROC_ADDRESS(string_operator_index_const, GDExtensionInterfaceStringOperatorIndexConst);
	LOAD_PROC_ADDRESS(string_operator_plus_eq_string, GDExtensionInterfaceStringOperatorPlusEqString);
	LOAD_PROC_ADDRESS(string_operator_plus_eq_char, GDExtensionInterfaceStringOperatorPlusEqChar);
	LOAD_PROC_ADDRESS(string_operator_plus_eq_cstr, GDExtensionInterfaceStringOperatorPlusEqCstr);
	LOAD_PROC_ADDRESS(string_operator_plus_eq_wcstr, GDExtensionInterfaceStringOperatorPlusEqWcstr);
	LOAD_PROC_ADDRESS(string_operator_plus_eq_c32str, GDExtensionInterfaceStringOperatorPlusEqC32str);
	LOAD_PROC_ADDRESS(string_resize, GDExtensionInterfaceStringResize);
	LOAD_PROC_ADDRESS(string_name_new_with_latin1_chars, GDExtensionInterfaceStringNameNewWithLatin1Chars);
	LOAD_PROC_ADDRESS(string_name_new_with_utf8_chars_and_len, GDExtensionInterfaceStringNameNewWithUtf8CharsAndLen);
	LOAD_PROC_ADDRESS(xml_parser_open_buffer, GDExtensionInterfaceXmlParserOpenBuffer);
	LOAD_PROC_ADDRESS(file_access_store_buffer, GDExtensionInterfaceFileAccessStoreBuffer);
	LOAD_PROC_ADDRESS(file_access_get_buffer, GDExtensionInterfaceFileAccessGetBuffer);
	LOAD_PROC_ADDRESS(worker_thread_pool_add_native_group_task, GDExtensionInterfaceWorkerThreadPoolAddNativeGroupTask);
	LOAD_PROC_ADDRESS(worker_thread_pool_add_native_task, GDExtensionInterfaceWorkerThreadPoolAddNativeTask);
	LOAD_PROC_ADDRESS(packed_byte_array_operator_index, GDExtensionInterfacePackedByteArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_byte_array_operator_index_const, GDExtensionInterfacePackedByteArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_color_array_operator_index, GDExtensionInterfacePackedColorArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_color_array_operator_index_const, GDExtensionInterfacePackedColorArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_float32_array_operator_index, GDExtensionInterfacePackedFloat32ArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_float32_array_operator_index_const, GDExtensionInterfacePackedFloat32ArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_float64_array_operator_index, GDExtensionInterfacePackedFloat64ArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_float64_array_operator_index_const, GDExtensionInterfacePackedFloat64ArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_int32_array_operator_index, GDExtensionInterfacePackedInt32ArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_int32_array_operator_index_const, GDExtensionInterfacePackedInt32ArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_int64_array_operator_index, GDExtensionInterfacePackedInt64ArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_int64_array_operator_index_const, GDExtensionInterfacePackedInt64ArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_string_array_operator_index, GDExtensionInterfacePackedStringArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_string_array_operator_index_const, GDExtensionInterfacePackedStringArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_vector2_array_operator_index, GDExtensionInterfacePackedVector2ArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_vector2_array_operator_index_const, GDExtensionInterfacePackedVector2ArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_vector3_array_operator_index, GDExtensionInterfacePackedVector3ArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_vector3_array_operator_index_const, GDExtensionInterfacePackedVector3ArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(packed_vector4_array_operator_index, GDExtensionInterfacePackedVector4ArrayOperatorIndex);
	LOAD_PROC_ADDRESS(packed_vector4_array_operator_index_const, GDExtensionInterfacePackedVector4ArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(array_operator_index, GDExtensionInterfaceArrayOperatorIndex);
	LOAD_PROC_ADDRESS(array_operator_index_const, GDExtensionInterfaceArrayOperatorIndexConst);
	LOAD_PROC_ADDRESS(array_set_typed, GDExtensionInterfaceArraySetTyped);
	LOAD_PROC_ADDRESS(dictionary_operator_index, GDExtensionInterfaceDictionaryOperatorIndex);
	LOAD_PROC_ADDRESS(dictionary_operator_index_const, GDExtensionInterfaceDictionaryOperatorIndexConst);
	LOAD_PROC_ADDRESS(dictionary_set_typed, GDExtensionInterfaceDictionarySetTyped);
	LOAD_PROC_ADDRESS(object_method_bind_call, GDExtensionInterfaceObjectMethodBindCall);
	LOAD_PROC_ADDRESS(object_method_bind_ptrcall, GDExtensionInterfaceObjectMethodBindPtrcall);
	LOAD_PROC_ADDRESS(object_destroy, GDExtensionInterfaceObjectDestroy);
	LOAD_PROC_ADDRESS(global_get_singleton, GDExtensionInterfaceGlobalGetSingleton);
	LOAD_PROC_ADDRESS(object_get_instance_binding, GDExtensionInterfaceObjectGetInstanceBinding);
	LOAD_PROC_ADDRESS(object_set_instance_binding, GDExtensionInterfaceObjectSetInstanceBinding);
	LOAD_PROC_ADDRESS(object_free_instance_binding, GDExtensionInterfaceObjectFreeInstanceBinding);
	LOAD_PROC_ADDRESS(object_set_instance, GDExtensionInterfaceObjectSetInstance);
	LOAD_PROC_ADDRESS(object_get_class_name, GDExtensionInterfaceObjectGetClassName);
	LOAD_PROC_ADDRESS(object_cast_to, GDExtensionInterfaceObjectCastTo);
	LOAD_PROC_ADDRESS(object_get_instance_from_id, GDExtensionInterfaceObjectGetInstanceFromId);
	LOAD_PROC_ADDRESS(object_get_instance_id, GDExtensionInterfaceObjectGetInstanceId);
	LOAD_PROC_ADDRESS(object_has_script_method, GDExtensionInterfaceObjectHasScriptMethod);
	LOAD_PROC_ADDRESS(object_call_script_method, GDExtensionInterfaceObjectCallScriptMethod);
	LOAD_PROC_ADDRESS(callable_custom_create2, GDExtensionInterfaceCallableCustomCreate2);
	LOAD_PROC_ADDRESS(callable_custom_get_userdata, GDExtensionInterfaceCallableCustomGetUserData);
	LOAD_PROC_ADDRESS(ref_get_object, GDExtensionInterfaceRefGetObject);
	LOAD_PROC_ADDRESS(ref_set_object, GDExtensionInterfaceRefSetObject);
	LOAD_PROC_ADDRESS(script_instance_create3, GDExtensionInterfaceScriptInstanceCreate3);
	LOAD_PROC_ADDRESS(placeholder_script_instance_create, GDExtensionInterfacePlaceHolderScriptInstanceCreate);
	LOAD_PROC_ADDRESS(placeholder_script_instance_update, GDExtensionInterfacePlaceHolderScriptInstanceUpdate);
	LOAD_PROC_ADDRESS(object_get_script_instance, GDExtensionInterfaceObjectGetScriptInstance);
	//LOAD_PROC_ADDRESS(object_set_script_instance, GDExtensionInterfaceObjectSetScriptInstance);
	LOAD_PROC_ADDRESS(classdb_construct_object2, GDExtensionInterfaceClassdbConstructObject2);
	LOAD_PROC_ADDRESS(classdb_get_method_bind, GDExtensionInterfaceClassdbGetMethodBind);
	LOAD_PROC_ADDRESS(classdb_get_class_tag, GDExtensionInterfaceClassdbGetClassTag);
	LOAD_PROC_ADDRESS(classdb_register_extension_class4, GDExtensionInterfaceClassdbRegisterExtensionClass4);
	LOAD_PROC_ADDRESS(classdb_register_extension_class_method, GDExtensionInterfaceClassdbRegisterExtensionClassMethod);
	LOAD_PROC_ADDRESS(classdb_register_extension_class_virtual_method, GDExtensionInterfaceClassdbRegisterExtensionClassVirtualMethod);
	LOAD_PROC_ADDRESS(classdb_register_extension_class_integer_constant, GDExtensionInterfaceClassdbRegisterExtensionClassIntegerConstant);
	LOAD_PROC_ADDRESS(classdb_register_extension_class_property, GDExtensionInterfaceClassdbRegisterExtensionClassProperty);
	LOAD_PROC_ADDRESS(classdb_register_extension_class_property_indexed, GDExtensionInterfaceClassdbRegisterExtensionClassPropertyIndexed);
	LOAD_PROC_ADDRESS(classdb_register_extension_class_property_group, GDExtensionInterfaceClassdbRegisterExtensionClassPropertyGroup);
	LOAD_PROC_ADDRESS(classdb_register_extension_class_property_subgroup, GDExtensionInterfaceClassdbRegisterExtensionClassPropertySubgroup);
	LOAD_PROC_ADDRESS(classdb_register_extension_class_signal, GDExtensionInterfaceClassdbRegisterExtensionClassSignal);
	LOAD_PROC_ADDRESS(classdb_unregister_extension_class, GDExtensionInterfaceClassdbUnregisterExtensionClass);
	LOAD_PROC_ADDRESS(get_library_path, GDExtensionInterfaceGetLibraryPath);
	LOAD_PROC_ADDRESS(editor_add_plugin, GDExtensionInterfaceEditorAddPlugin);
	LOAD_PROC_ADDRESS(editor_remove_plugin, GDExtensionInterfaceEditorRemovePlugin);
	//LOAD_PROC_ADDRESS(editor_register_get_classes_used_callback, GDExtensionInterfaceEditorRegisterGetClassesUsedCallback);
	LOAD_PROC_ADDRESS(editor_help_load_xml_from_utf8_chars, GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8Chars);
	LOAD_PROC_ADDRESS(editor_help_load_xml_from_utf8_chars_and_len, GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8CharsAndLen);
	LOAD_PROC_ADDRESS(image_ptrw, GDExtensionInterfaceImagePtrw);
	LOAD_PROC_ADDRESS(image_ptr, GDExtensionInterfaceImagePtr);
	//LOAD_PROC_ADDRESS(register_main_loop_callbacks, GDExtensionInterfaceRegisterMainLoopCallbacks);
    cgo_library = p_library;
    r_initialization->userdata = 0;
    r_initialization->minimum_initialization_level = GDEXTENSION_INITIALIZATION_CORE;
    r_initialization->initialize = cgo_initialize;
    r_initialization->deinitialize = cgo_deinitialize;
    //((GDExtensionInterfaceGetGodotVersion2)p_get_proc_address("get_godot_version2"))(&cgo_cached_godot_version);

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

void prepare_variants(void **frame, uint32_t argc, ANY args) {
    uint8_t *head = (uint8_t*)args;
    for (int i = 0; i < argc; i++) {
        frame[i] = head;
        head += 24;
    }
}
// Helper macro to align a value to the next multiple of 'align'
#define ALIGN_UP(value, align) (((value) + ((align) - 1)) & ~((align) - 1))

uint8_t prepare_callframe(void **frame, UINT64 shape, ANY args) {
    uint8_t *head = (uint8_t *)args;
    ptrdiff_t offset = 0; // Track current offset in the frame
    for (int i = 1; i < 16; i++) {
        uint32_t code = (UINT64_FROM(shape) >> (i * 4)) & 0xF;
        uint32_t size;
        // Determine size based on code
        switch (code) {
            case 0: size = 0; frame[i-1] = NULL; return i; // Early return for zero-sized argument
            case 1: size = 1; break;
            case 2: size = 2; break;
            case 3: size = 4; break;
            case 4: size = 8; break;
            case 5: size = 12; break;
            case 6: size = 16; break;
            case 7: size = 24; break;
            case 8: size = 32; break;
            case 9: size = 36; break;
            case 10: size = 40; break;
            case 11: size = 48; break;
            case 12: size = 64; break;
            case 13: size = 72; break;
            case 14: size = 96; break;
            case 15: size = 128; break;
        }
        // Align to the minimum of the argument's size or 8 bytes (Go's max alignment on 64-bit systems)
        uint32_t alignment = (size > 8) ? 8 : size;
        offset = ALIGN_UP(offset, alignment);
        frame[i-1] = head + offset;     // Set frame pointer to the aligned address
        offset += size;                 // Move offset forward by the size of the current argument
    }
    return 15;
}

uintptr_t gd_builtin_name(uintptr_t name, INT64 hash) { return (uintptr_t)gdextension_variant_get_ptr_utility_function((GDExtensionConstStringNamePtr)name, INT64_FROM(hash));}
void gd_builtin_call(uintptr_t fn, void *result, UINT64 shape, ANY args) {
    void *points[16]; uint8_t argc = prepare_callframe(&points[0], shape, args);
    ((GDExtensionPtrUtilityFunction)fn)(result, (GDExtensionConstTypePtr*)&points[0], argc);
}

void gd_callable_create(uintptr_t id, UINT64 object, void *result) {
    GDExtensionCallableCustomInfo2 info = {
        .callable_userdata = (void *)id,
        .token = cgo_library,
        .object_id = UINT64_FROM(object),
        .call_func = cgo_callable_call_func,
        .is_valid_func = cgo_callable_is_valid_func,
        .free_func = cgo_callable_free_func,
        .hash_func = cgo_callable_hash_func,
        .equal_func = cgo_callable_equal_func,
        .less_than_func = cgo_callable_less_than_func,
        .to_string_func = cgo_callable_to_string_func,
        .get_argument_count_func = cgo_callable_get_argument_count_func,
    };
    gdextension_callable_custom_create2((uint64_t *)result+sizeof(uint64_t), &info);
}

uintptr_t gd_callable_lookup(UINT64 a, UINT64 b) {
    uint64_t callable[2] = {UINT64_FROM(a), UINT64_FROM(b)};
    return (uintptr_t)gdextension_callable_custom_get_userdata(&callable, cgo_library);
};

void gd_classdb_FileAccess_write(uintptr_t FileAccess, BUFFER buf, INT len) {
    gdextension_file_access_store_buffer((GDExtensionObjectPtr)FileAccess, (const uint8_t *)BUFFER_POINTER(buf), len);
};

INT gd_classdb_FileAccess_read(uintptr_t FileAccess, BUFFER buf, INT len) {
    return INT64_MAKE(gdextension_file_access_get_buffer((GDExtensionObjectPtr)FileAccess, (uint8_t *)BUFFER_POINTER(buf), len));
};

uintptr_t gd_classdb_Image_unsafe(uintptr_t Image) {
    return (uintptr_t)gdextension_image_ptrw((GDExtensionObjectPtr)Image);
};

uint8_t gd_classdb_Image_access(uintptr_t Image, INT offset) {
    return gdextension_image_ptr((GDExtensionObjectPtr)Image)[offset];
};

typedef struct {
    int32_t push;
    GDExtensionClassMethodInfo *info;
} method_list;

typedef struct {
    int32_t push;
    GDExtensionPropertyInfo *info;
    GDExtensionClassMethodArgumentMetadata *meta;
} property_list;

uintptr_t gd_method_list_make(INT length) {
    method_list *list = (method_list *)gdextension_mem_alloc(sizeof(method_list));
    list->push = 0;
    list->info = (GDExtensionClassMethodInfo*)gdextension_mem_alloc(sizeof(GDExtensionClassMethodInfo) * length);
    return (uintptr_t)list;
};

void gd_method_list_push(uintptr_t list_p, uintptr_t name, uintptr_t method, uint32_t method_flags, bool has_return_value, uintptr_t return_value_info, uint32_t argument_count, uintptr_t arguments_info, INT default_argument_count, void *default_arguments) {
    method_list *list = (method_list *)list_p;
    GDExtensionClassMethodInfo *info = &list->info[list->push++];
    property_list *return_value = (property_list *)return_value_info;
    property_list *arguments = (property_list *)arguments_info;
    uintptr_t *name_allocated = (uintptr_t *)gdextension_mem_alloc(sizeof(uintptr_t));
    *name_allocated = name;
    info->name = (GDExtensionStringNamePtr)name_allocated;
    info->method_userdata = (void *)method;
    info->call_func = cgo_method_call_func;
    info->ptrcall_func = cgo_method_ptrcall_func;
    info->method_flags = method_flags;
    info->has_return_value = has_return_value;
    info->return_value_info = return_value->info;
    info->return_value_metadata = *return_value->meta;
    info->argument_count = argument_count;
    info->arguments_info = arguments->info;
    info->arguments_metadata = arguments->meta;
    info->default_argument_count = default_argument_count;
    void *points[16]; prepare_variants(&points[0], default_argument_count, (ANY)default_arguments);
    info->default_arguments = &points[0];
};

void gd_method_list_free(uintptr_t list_p) {
    method_list *list = (method_list *)list_p;
    for (int i = 0; i < list->push; i++) {
        gdextension_mem_free(list->info[i].name);
    }
    gdextension_mem_free(list->info); gdextension_mem_free(list);
};

uintptr_t gd_property_list_make(INT length) {
    property_list *list = (property_list*)gdextension_mem_alloc(sizeof(property_list));
    list->push = 0;
    list->info = (GDExtensionPropertyInfo*)gdextension_mem_alloc(sizeof(GDExtensionPropertyInfo) * length);
    list->meta = (GDExtensionClassMethodArgumentMetadata*)gdextension_mem_alloc(sizeof(GDExtensionClassMethodArgumentMetadata) * length);
    return (uintptr_t)list;
};

void gd_property_list_push(uintptr_t list_p, uint32_t vtype, uintptr_t name, uintptr_t class_name, uint32_t hint, uintptr_t hint_string, uint32_t usage, uint32_t meta) {
    property_list *list = (property_list *)list_p;
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
    info->hint_string = (GDExtensionStringNamePtr)hint_string_allocated;
    info->usage = usage;
    *meta_info = (GDExtensionClassMethodArgumentMetadata)meta;
};

void gd_property_list_free(uintptr_t list_p) {
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

uint32_t gd_property_info_type(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return list->info[list->push].type;
};

uintptr_t gd_property_info_name(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return (uintptr_t)list->info[list->push].name;
};

uintptr_t gd_property_info_class_name(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return (uintptr_t)list->info[list->push].class_name;
};

uint32_t gd_property_info_hint(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return list->info[list->push].hint;
};

uintptr_t gd_property_info_hint_string(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return (uintptr_t)list->info[list->push].hint_string;
};

uint32_t gd_property_info_usage(uintptr_t list_p) {
    property_list *list = (property_list *)list_p;
    return list->info[list->push].usage;
};

GDExtensionBool cgo_class_set_func(GDExtensionClassInstancePtr instance, GDExtensionConstStringNamePtr field, GDExtensionConstVariantPtr value) {
    uint64_t *v = (uint64_t *)value;
    return go_on_extension_instance_set((uintptr_t)instance, *(uintptr_t*)field, v[0], v[1], v[2]);
}
GDExtensionBool cgo_class_get_func(GDExtensionClassInstancePtr instance, GDExtensionConstStringNamePtr field, GDExtensionVariantPtr value) {
    reverse_variant_frame frame = { .r_return = value, .r_error = NULL, .p_args = NULL, .p_argument_count = 0 };
    return go_on_extension_instance_get((uintptr_t)instance, *(uintptr_t*)field, &frame);
}
GDExtensionBool cgo_class_property_can_revert_func(GDExtensionClassInstancePtr instance, GDExtensionConstStringNamePtr field) {
    return go_on_extension_instance_property_has_default((uintptr_t)instance, *(uintptr_t*)field);
}
GDExtensionBool cgo_class_property_get_revert_func(GDExtensionClassInstancePtr instance, GDExtensionConstStringNamePtr field, GDExtensionVariantPtr value) {
    reverse_variant_frame frame = { .r_return = value, .r_error = NULL, .p_args = NULL, .p_argument_count = 0 };
    return go_on_extension_instance_property_get_default((uintptr_t)instance, *(uintptr_t*)field, &frame);
}

const GDExtensionPropertyInfo *cgo_class_get_property_list_func(GDExtensionClassInstancePtr instance, uint32_t *count) {
    property_list *list = (property_list*)go_on_extension_instance_property_list((uintptr_t)instance);
    GDExtensionPropertyInfo *info = list ? list->info : NULL;
    *count = list ? list->push : 0;
    gdextension_mem_free(list->meta);
    gdextension_mem_free(list);
    return info;
}

void cgo_class_free_property_list_func(GDExtensionClassInstancePtr instance, const GDExtensionPropertyInfo *list, uint32_t count) {
   gdextension_mem_free((void*)list);
}

GDExtensionBool cgo_class_validate_property_func(GDExtensionClassInstancePtr instance, GDExtensionPropertyInfo *field) {
    property_list list = {
        .push = 0,
        .info = field,
        .meta = NULL
    };
    return go_on_extension_instance_property_validation((uintptr_t)instance, (uintptr_t)&list);
}

void cgo_class_to_string_func(GDExtensionClassInstancePtr instance, GDExtensionBool *ok, GDExtensionStringPtr s) {
    *(uintptr_t*)s = go_on_extension_instance_stringify((uintptr_t)instance);
}

void cgo_class_reference_func(GDExtensionClassInstancePtr instance) {
    go_on_extension_instance_reference((uintptr_t)instance, true);
}

GDExtensionBool cgo_class_unreference_func(GDExtensionClassInstancePtr instance) {
    return go_on_extension_instance_reference((uintptr_t)instance, false);
}

GDExtensionObjectPtr cgo_class_create_instance_func(void *user_data, GDExtensionBool notify_postinitialize) {
    return (GDExtensionObjectPtr)go_on_extension_class_create((uintptr_t)user_data, notify_postinitialize);
}

void *cgo_class_get_virtual_call_data_func(void *user_data, GDExtensionConstStringNamePtr name, uint32_t hash) {
    return (void*)go_on_extension_class_method((uintptr_t)user_data, *(uintptr_t*)name, hash);
}

void cgo_class_call_virtual_with_data_func(GDExtensionClassInstancePtr p_instance, GDExtensionConstStringNamePtr p_name, void *p_virtual_call_userdata, const GDExtensionConstTypePtr *p_args, GDExtensionTypePtr r_ret) {
    go_on_extension_instance_unsafe_call((uintptr_t)p_instance, (uintptr_t)p_virtual_call_userdata, r_ret, (void *)p_args);
}

void cgo_class_free_instance_func(void *p_class_userdata, GDExtensionClassInstancePtr p_instance) {
    go_on_extension_instance_free((uintptr_t)p_instance);
}

void gd_classdb_register(uintptr_t class_name, uintptr_t parent, uintptr_t id, bool is_virtual, bool abstract, bool exposed, bool runtime, uintptr_t icon_path) {
    GDExtensionClassCreationInfo4 info = {
        .is_virtual = is_virtual,
        .is_abstract = abstract,
        .is_exposed = exposed,
        .is_runtime = runtime,
        .icon_path = (GDExtensionConstStringNamePtr)&icon_path,
        .set_func = cgo_class_set_func,
        .get_func = cgo_class_get_func,
        .get_property_list_func = cgo_class_get_property_list_func,
        .free_property_list_func = cgo_class_free_property_list_func,
        .property_can_revert_func = cgo_class_property_can_revert_func,
        .property_get_revert_func = cgo_class_property_get_revert_func,
        .validate_property_func = cgo_class_validate_property_func,
        .notification_func = (GDExtensionClassNotification2)go_on_extension_instance_notification,
        .to_string_func = cgo_class_to_string_func,
        .reference_func = (GDExtensionClassReference)cgo_class_reference_func,
        .unreference_func = (GDExtensionClassUnreference)cgo_class_unreference_func,
        .create_instance_func = cgo_class_create_instance_func,
        .free_instance_func = cgo_class_free_instance_func,
        .get_virtual_call_data_func = cgo_class_get_virtual_call_data_func,
        .call_virtual_with_data_func = cgo_class_call_virtual_with_data_func,
        .class_userdata = (void *)id,
    };
    gdextension_classdb_register_extension_class4(cgo_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&parent, &info);
};

void gd_classdb_register_methods(uintptr_t class_name, uintptr_t methods) {
    method_list *list = (method_list *)methods;
    for (int i = 0; i < list->push; i++) {
        GDExtensionClassMethodInfo *info = &list->info[i];
        gdextension_classdb_register_extension_class_method(cgo_library, (GDExtensionConstStringNamePtr)&class_name, info);
    }
};

void gd_classdb_register_constant(uintptr_t class_name, uintptr_t enum_name, uintptr_t name, INT64 value, bool bitfield) {
    gdextension_classdb_register_extension_class_integer_constant(cgo_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&enum_name, (GDExtensionConstStringNamePtr)&name, INT64_FROM(value), bitfield);
};

void gd_classdb_register_property(uintptr_t class_name, uintptr_t info, uintptr_t setter, uintptr_t getter) {
    property_list *list = (property_list *)info;
    gdextension_classdb_register_extension_class_property(cgo_library, (GDExtensionConstStringNamePtr)&class_name, list->info, (GDExtensionConstStringNamePtr)&setter, (GDExtensionConstStringNamePtr)&getter);
};

void gd_classdb_register_property_indexed(uintptr_t class_name, uintptr_t info, uintptr_t setter, uintptr_t getter, INT64 index) {
    property_list *list = (property_list *)info;
    gdextension_classdb_register_extension_class_property_indexed(cgo_library, (GDExtensionConstStringNamePtr)&class_name, list->info, (GDExtensionConstStringNamePtr)&setter, (GDExtensionConstStringNamePtr)&getter, INT64_FROM(index));
};

void gd_classdb_register_property_group(uintptr_t class_name, uintptr_t group, uintptr_t prefix) {
    gdextension_classdb_register_extension_class_property_group(cgo_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&group, (GDExtensionConstStringPtr)&prefix);
};

void gd_classdb_register_property_sub_group(uintptr_t class_name, uintptr_t subgroup, uintptr_t prefix) {
    gdextension_classdb_register_extension_class_property_subgroup(cgo_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&subgroup, (GDExtensionConstStringPtr)&prefix);
};

void gd_classdb_register_signal(uintptr_t class_name, uintptr_t name, uintptr_t args) {
    property_list *list = (property_list *)args;
    gdextension_classdb_register_extension_class_signal(cgo_library, (GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&name, list->info, list->push);
};

void gd_classdb_register_removal(uintptr_t class_name) {
    gdextension_classdb_unregister_extension_class(cgo_library, (GDExtensionConstStringNamePtr)&class_name);
};

void gd_classdb_WorkerThreadPool_add_task(uintptr_t WorkerPool, uintptr_t task_id, bool priority, uintptr_t description) {
    gdextension_worker_thread_pool_add_native_task((GDExtensionObjectPtr)WorkerPool, (GDExtensionWorkerThreadPoolTask)go_on_worker_thread_pool_task, (void *)task_id, priority, (GDExtensionConstStringNamePtr)&description);
};

void gd_classdb_WorkerThreadPool_add_group_task(uintptr_t WorkerPool, uintptr_t task_id, int32_t elements, int32_t tasks, bool priority, uintptr_t description) {
    gdextension_worker_thread_pool_add_native_group_task((GDExtensionObjectPtr)WorkerPool, (GDExtensionWorkerThreadPoolGroupTask)go_on_worker_thread_pool_group_task, (void *)task_id, elements, tasks, priority, (GDExtensionConstStringNamePtr)&description);
};

INT64 gd_classdb_XMLParser_load(uintptr_t XMLParser, BUFFER buf, INT len) {
    return INT64_MAKE(gdextension_xml_parser_open_buffer((GDExtensionObjectPtr)XMLParser, (const uint8_t *)BUFFER_POINTER(buf), len));
};

void gd_packed_dictionary_access(uintptr_t dict, UINT64 k1, UINT64 k2, UINT64 k3, ANY args) {
    uint64_t key[3] = {UINT64_FROM(k1), UINT64_FROM(k2), UINT64_FROM(k3)};
    uint64_t *value = (uint64_t*)gdextension_dictionary_operator_index_const((GDExtensionTypePtr)&dict, &key[0]);
    if (!value) return;
    uint64_t * result = (uint64_t*)args;
    result[0] = value[0];
    result[1] = value[1];
    result[2] = value[2];
};

void gd_packed_dictionary_modify(uintptr_t dict, UINT64 k1, UINT64 k2, UINT64 k3, UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t key[3] = {UINT64_FROM(k1), UINT64_FROM(k2), UINT64_FROM(k3)};
    uint64_t *value = (uint64_t*)gdextension_dictionary_operator_index((GDExtensionTypePtr)&dict, (GDExtensionVariantPtr)&key[0]);
    value[0] = UINT64_FROM(v1);
    value[1] = UINT64_FROM(v2);
    value[2] = UINT64_FROM(v3);
};

void gd_editor_add_documentation(STRING xml, INT len) {
    gdextension_editor_help_load_xml_from_utf8_chars_and_len(STRING_POINTER(xml), len);
};

void gd_editor_add_plugin(uintptr_t class_name) {
    gdextension_editor_add_plugin((GDExtensionConstStringNamePtr)&class_name);
};

void gd_editor_end_plugin(uintptr_t class_name) {
    gdextension_editor_remove_plugin((GDExtensionConstStringNamePtr)&class_name);
};

void gd_iterator_make(UINT64 v1, UINT64 v2, UINT64 v3, ANY result, void *err) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    GDExtensionBool valid = true;
    valid = valid && gdextension_variant_iter_init(&self, (void*)result, &valid);
    if (valid) return;
    ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
};

bool gd_iterator_next(UINT64 v1, UINT64 v2, UINT64 v3, void *recv, UINT64 i1, uint64_t i2, UINT64 i3, void *err) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    uint64_t iter[3] = {UINT64_FROM(i1), UINT64_FROM(i2), UINT64_FROM(i3)};
    GDExtensionBool valid = false;
    GDExtensionBool ok = gdextension_variant_iter_next(&self, &iter, &valid);
    if (ok) {
        ((uint64_t *)recv)[0] = iter[0];
        ((uint64_t *)recv)[1] = iter[1];
        ((uint64_t *)recv)[2] = iter[2];
        return true;
    }
    ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return false;
};

void gd_iterator_load(UINT64 v1, UINT64 v2, UINT64 v3, ANY args, UINT64 i1, UINT64 i2, UINT64 i3, void *err) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    uint64_t iter[3] = {UINT64_FROM(i1), UINT64_FROM(i2), UINT64_FROM(i3)};
    void *points[16]; prepare_variants(&points[0], 1, args);
    GDExtensionBool ok = false;
    gdextension_variant_iter_get(&self, &iter, (GDExtensionUninitializedVariantPtr)points[0], &ok);
    if (ok) return;
    ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
};

const char *fit_string(const char *str, uint64_t len, char *buf, size_t buf_size) {
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

void gd_log_error(
    STRING text, INT text_len,
    STRING code, INT code_len,
    STRING func, INT func_len,
    STRING file, INT file_len,
    int32_t line, bool notify_editor
) {
    char text_buf[256]; const char *text_ptr = fit_string(STRING_POINTER(text), text_len, &text_buf[0], 256);
    char code_buf[100]; const char *code_ptr = fit_string(STRING_POINTER(code), code_len, &code_buf[0], 100);
    char func_buf[100]; const char *func_ptr = fit_string(STRING_POINTER(func), func_len, &func_buf[0], 100);
    char file_buf[100]; const char *file_ptr = fit_string(STRING_POINTER(file), file_len, &file_buf[0], 100);
    gdextension_print_error_with_message(code_ptr, text_ptr, func_ptr, file_ptr, line, notify_editor);
    if (text_ptr && text_ptr != text_buf) gdextension_mem_free((void *)text_ptr);
    if (code_ptr && code_ptr != code_buf) gdextension_mem_free((void *)code_ptr);
    if (func_ptr && func_ptr != func_buf) gdextension_mem_free((void *)func_ptr);
    if (file_ptr && file_ptr != file_buf) gdextension_mem_free((void *)file_ptr);
};

void gd_log_warning(
    STRING text, INT text_len,
    STRING code, INT code_len,
    STRING func, INT func_len,
    STRING file, INT file_len,
    int32_t line, bool notify_editor
) {
    char text_buf[256]; const char *text_ptr = fit_string(STRING_POINTER(text), text_len, &text_buf[0], 256);
    char code_buf[100]; const char *code_ptr = fit_string(STRING_POINTER(code), code_len, &code_buf[0], 100);
    char func_buf[100]; const char *func_ptr = fit_string(STRING_POINTER(func), func_len, &func_buf[0], 100);
    char file_buf[100]; const char *file_ptr = fit_string(STRING_POINTER(file), file_len, &file_buf[0], 100);
    gdextension_print_warning_with_message(code_ptr, text_ptr, func_ptr, file_ptr, line, notify_editor);
    if (text_ptr && text_ptr != text_buf) gdextension_mem_free((void *)text_ptr);
    if (code_ptr && code_ptr != code_buf) gdextension_mem_free((void *)code_ptr);
    if (func_ptr && func_ptr != func_buf) gdextension_mem_free((void *)func_ptr);
    if (file_ptr && file_ptr != file_buf) gdextension_mem_free((void *)file_ptr);
};

uintptr_t gd_memory_malloc(INT size) {
    return (uintptr_t)gdextension_mem_alloc(size);
};

INT gd_memory_sizeof(uintptr_t name) {
    return gdextension_get_native_struct_size((GDExtensionConstStringNamePtr)&name);
};

uintptr_t gd_memory_resize(uintptr_t addr, INT size) {
    return (uintptr_t)gdextension_mem_realloc((void *)addr, size);
};

void gd_memory_free(uintptr_t addr) {
    gdextension_mem_free((void *)addr);
};

void gd_memory_edit_byte(uintptr_t addr, uint8_t b) { *(uint8_t *)addr = b; };
void gd_memory_edit_u16(uintptr_t addr, uint16_t b) { *(uint16_t *)addr = b; };
void gd_memory_edit_u32(uintptr_t addr, uint32_t b) { *(uint32_t *)addr = b; };
void gd_memory_edit_u64(uintptr_t addr, UINT64 b) { *(uint64_t *)addr = UINT64_FROM(b); };
void gd_memory_edit_128(uintptr_t addr, UINT64 a, UINT64 b) {
    uint64_t *ptr = (uint64_t *)addr;
    ptr[0] = UINT64_FROM(a); ptr[1] = UINT64_FROM(b);
};
void gd_memory_edit_256(uintptr_t addr, UINT64 a, UINT64 b, UINT64 c, UINT64 d) {
    uint64_t *ptr = (uint64_t *)addr;
    ptr[0] = UINT64_FROM(a); ptr[1] = UINT64_FROM(b); ptr[2] = UINT64_FROM(c); ptr[3] = UINT64_FROM(d);
};
void gd_memory_edit_512(uintptr_t addr, UINT64 a, UINT64 b, UINT64 c, UINT64 d, UINT64 e, UINT64 f, UINT64 g, UINT64 h) {
    uint64_t *ptr = (uint64_t *)addr;
    ptr[0] = UINT64_FROM(a); ptr[1] = UINT64_FROM(b); ptr[2] = UINT64_FROM(c); ptr[3] = UINT64_FROM(d);
    ptr[4] = UINT64_FROM(e); ptr[5] = UINT64_FROM(f); ptr[6] = UINT64_FROM(g); ptr[7] = UINT64_FROM(h);
};

uint8_t gd_memory_load_byte(uintptr_t addr) { return *(uint8_t *)addr; };
uint16_t gd_memory_load_u16(uintptr_t addr) { return *(uint16_t *)addr; };
uint32_t gd_memory_load_u32(uintptr_t addr) { return *(uint32_t *)addr; };
UINT64 gd_memory_load_u64(uintptr_t addr) { return UINT64_FROM(*(uint64_t *)(addr)); };

uintptr_t gd_object_make(uintptr_t name) {
    return (uintptr_t)gdextension_classdb_construct_object2((GDExtensionConstStringNamePtr)&name);
};

void gd_object_call(uintptr_t obj, uintptr_t fn, void *result, INT argc, ANY args, ANY err) {
    void *points[16]; prepare_variants(&points[0], argc, args);
    gdextension_object_method_bind_call((GDExtensionMethodBindPtr)fn, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};

void gd_object_unsafe_free(uintptr_t obj) {
    gdextension_object_destroy((GDExtensionObjectPtr)obj);
};

uintptr_t gd_object_name(uintptr_t obj) {
    uintptr_t name = 0;
    gdextension_object_get_class_name((GDExtensionObjectPtr)obj, cgo_library, (GDExtensionUninitializedStringNamePtr)&name);
    return name;
};

uintptr_t gd_object_type(uintptr_t name) {
    return (uintptr_t)gdextension_classdb_get_class_tag((GDExtensionConstStringNamePtr)&name);
};

uintptr_t gd_object_cast(uintptr_t obj, uintptr_t tag) {
    return (uintptr_t)gdextension_object_cast_to((GDExtensionObjectPtr)obj, (void *)tag);
};

uintptr_t gd_object_lookup(UINT64 id) {
    return (uintptr_t)gdextension_object_get_instance_from_id((GDObjectInstanceID)UINT64_FROM(id));
};

uintptr_t gd_object_global(uintptr_t name) {
    return (uintptr_t)gdextension_global_get_singleton((GDExtensionConstStringNamePtr)&name);
};

void gd_object_extension_setup(uintptr_t obj, uintptr_t name, uintptr_t instance) {
    gdextension_object_set_instance((GDExtensionObjectPtr)obj, (GDExtensionConstStringNamePtr)&name, (GDExtensionClassInstancePtr)instance);
    gdextension_object_set_instance_binding((GDExtensionObjectPtr)obj, cgo_library, (void *)instance, &instance_binding_callbacks);
};

uintptr_t gd_object_extension_fetch(uintptr_t obj) {
    return (uintptr_t)gdextension_object_get_instance_binding((GDExtensionObjectPtr)obj, cgo_library, &instance_binding_callbacks);
};

void gd_object_extension_close(uintptr_t obj) {
    gdextension_object_free_instance_binding((GDExtensionObjectPtr)obj, cgo_library);
};

uint64_t gd_object_id(uintptr_t obj) {
    return gdextension_object_get_instance_id((GDExtensionObjectPtr)obj);
};

UINT64 gd_object_id_inside_variant(UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    return UINT64_MAKE(gdextension_variant_get_object_instance_id(&self));
};

uintptr_t gd_object_method_lookup(uintptr_t class_name, uintptr_t method, INT64 hash) {
    return (uintptr_t)gdextension_classdb_get_method_bind((GDExtensionConstStringNamePtr)&class_name, (GDExtensionConstStringNamePtr)&method, INT64_FROM(hash));
};

void gd_object_unsafe_call(uintptr_t obj, uintptr_t method, ANY result, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    gdextension_object_method_bind_ptrcall((GDExtensionMethodBindPtr)method, (GDExtensionObjectPtr)obj, (const GDExtensionConstTypePtr*)&points[0], (GDExtensionTypePtr)result);
};

GDExtensionBool cgo_class_get_category_func(GDExtensionScriptInstanceDataPtr instance, GDExtensionPropertyInfo *info) {
    property_list list = { .info = info };
    return go_on_extension_script_categorization((uintptr_t)instance, (uintptr_t)&list);
}

GDExtensionObjectPtr cgo_class_get_owner_func(GDExtensionScriptInstanceDataPtr instance) {
    return (GDExtensionObjectPtr)go_on_extension_script_get_owner((uintptr_t)instance);
}

void gd_object_script_property_state_add(uintptr_t fn, uintptr_t arg, uintptr_t name, UINT64 s1, UINT64 s2, UINT64 s3) {
    uint64_t state[3] = {UINT64_FROM(s1), UINT64_FROM(s2), UINT64_FROM(s3)};
    ((GDExtensionScriptInstancePropertyStateAdd)fn)(&name, &state, (void *)arg);
}


void cgo_class_get_property_state_func(GDExtensionScriptInstanceDataPtr p_instance, GDExtensionScriptInstancePropertyStateAdd p_add_func, void *p_userdata) {
    go_on_extension_script_get_property_state((uintptr_t)p_instance, (uintptr_t)p_add_func, (uintptr_t)p_userdata);
}

GDExtensionBool cgo_class_has_method_func(GDExtensionScriptInstanceDataPtr instance, GDExtensionConstStringNamePtr method) {
    return go_on_extension_script_has_method((uintptr_t)instance, *(uintptr_t*)method);
}

GDExtensionInt cgo_class_get_method_argument_count_func(GDExtensionScriptInstanceDataPtr instance, GDExtensionConstStringNamePtr method, GDExtensionBool *valid) {
    *valid = 1;
    return go_on_extension_script_get_method_argument_count((uintptr_t)instance, *(uintptr_t*)method);
}

GDExtensionObjectPtr cgo_class_get_script_func(GDExtensionScriptInstanceDataPtr instance) {
    return (GDExtensionObjectPtr)go_on_extension_script_get((uintptr_t)instance);
}

GDExtensionBool cgo_class_is_placeholder_func(GDExtensionScriptInstanceDataPtr instance) {
    return go_on_extension_script_is_placeholder((uintptr_t)instance);
}

GDExtensionObjectPtr cgo_class_get_language_func(GDExtensionScriptInstanceDataPtr instance) {
    return (GDExtensionObjectPtr)go_on_extension_script_get_language((uintptr_t)instance);
}

uintptr_t gd_object_script_make(uintptr_t instance) {
    GDExtensionScriptInstanceInfo3 info = {
        .set_func = cgo_class_set_func,
        .get_func = cgo_class_get_func,
        .get_property_list_func = cgo_class_get_property_list_func,
        .free_property_list_func = cgo_class_free_property_list_func,
        .get_class_category_func = cgo_class_get_category_func,
        .property_can_revert_func = cgo_class_property_can_revert_func,
        .property_get_revert_func = cgo_class_property_get_revert_func,
        .get_owner_func = cgo_class_get_owner_func,
        .get_property_state_func = cgo_class_get_property_state_func,
        .validate_property_func = cgo_class_validate_property_func,
        .has_method_func = cgo_class_has_method_func,
        .get_method_argument_count_func = cgo_class_get_method_argument_count_func,
        .call_func = (GDExtensionScriptInstanceCall)cgo_method_call_func,
        .notification_func = (GDExtensionClassNotification2)go_on_extension_instance_notification,
        .to_string_func = cgo_class_to_string_func,
        .refcount_incremented_func = cgo_class_reference_func,
        .refcount_decremented_func = cgo_class_unreference_func,
        .get_script_func = cgo_class_get_script_func,
        .is_placeholder_func = cgo_class_is_placeholder_func,
        .get_language_func = cgo_class_get_language_func,
        .free_func = (GDExtensionScriptInstanceFree)go_on_extension_instance_free,
    };
    return (uintptr_t)gdextension_script_instance_create3(&info, (GDExtensionScriptInstanceDataPtr)&instance);
};

void gd_object_script_call(uintptr_t obj, uintptr_t method, void *result, INT argc, ANY args, ANY err) {
    void *points[16]; points[1] = 0; prepare_variants(&points[0], argc, args);
    gdextension_object_call_script_method((GDExtensionObjectPtr)obj, (GDExtensionConstStringNamePtr)method, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};

void gd_object_script_setup(uintptr_t obj, uintptr_t instance) {
    gdextension_object_set_script_instance((GDExtensionScriptInstanceDataPtr)obj, (GDExtensionClassInstancePtr)instance);
};

uintptr_t gd_object_script_fetch(uintptr_t obj, uintptr_t language) {
    return (uintptr_t)gdextension_object_get_script_instance((GDExtensionScriptInstanceDataPtr)obj, (GDExtensionClassInstancePtr)language);
};

bool gd_object_script_defines_method(uintptr_t obj, uintptr_t method) {
    return gdextension_object_has_script_method((GDExtensionScriptInstanceDataPtr)obj, (GDExtensionConstStringNamePtr)method);
};

uintptr_t gd_object_script_placeholder_create(uintptr_t language, uintptr_t script, uintptr_t owner) {
    return (uintptr_t)gdextension_placeholder_script_instance_create((GDExtensionScriptInstanceDataPtr)language, (GDExtensionObjectPtr)script, (GDExtensionObjectPtr)owner);
};

void gd_object_script_placeholder_update(uintptr_t p_placeholder, uintptr_t p_properties, uintptr_t p_values) {
    gdextension_placeholder_script_instance_update((GDExtensionScriptInstanceDataPtr)p_placeholder, (GDExtensionConstTypePtr)&p_properties, (GDExtensionConstTypePtr)&p_values);
};

uintptr_t gd_packed_byte_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_byte_array_operator_index(&packed_array[0], 0);
};

uint8_t gd_packed_byte_array_access(uintptr_t a1, uintptr_t a2, INT i) {
    uintptr_t packed_array[2] = {a1, a2};
    return *(uint8_t *)gdextension_packed_byte_array_operator_index_const(&packed_array[0], i);
};

uintptr_t gd_packed_color_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_color_array_operator_index(&packed_array[0], 0);
};

void gd_packed_color_array_access(uintptr_t a1, uintptr_t a2, INT i, ANY result) {
    uintptr_t packed_array[2] = {a1, a2};
    const float *color = (float *)gdextension_packed_color_array_operator_index_const(&packed_array[0], i);
    memcpy((float*)result, color, sizeof(float)*4);
};

uintptr_t gd_packed_float32_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_float32_array_operator_index(&packed_array[0], 0);
};

float gd_packed_float32_array_access(uintptr_t a1, uintptr_t a2, INT i) {
    uintptr_t packed_array[2] = {a1, a2};
    return *(float *)gdextension_packed_float32_array_operator_index_const(&packed_array[0], i);
};

uintptr_t gd_packed_float64_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_float64_array_operator_index(&packed_array[0], 0);
};

double gd_packed_float64_array_access(uintptr_t a1, uintptr_t a2, INT i) {
    uintptr_t packed_array[2] = {a1, a2};
    return *(double *)gdextension_packed_float64_array_operator_index_const(&packed_array[0], i);
};

uintptr_t gd_packed_int32_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_int32_array_operator_index(&packed_array[0], 0);
};

int32_t gd_packed_int32_array_access(uintptr_t a1, uintptr_t a2, INT i) {
    uintptr_t packed_array[2] = {a1, a2};
    return *(int32_t *)gdextension_packed_int32_array_operator_index_const(&packed_array[0], i);
};

uintptr_t gd_packed_int64_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_int64_array_operator_index(&packed_array[0], 0);
};

INT64 gd_packed_int64_array_access(uintptr_t a1, uintptr_t a2, INT i) {
    uintptr_t packed_array[2] = {a1, a2};
    return INT64_MAKE(*(int64_t *)gdextension_packed_int64_array_operator_index_const(&packed_array[0], i));
};

uintptr_t gd_packed_string_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_string_array_operator_index(&packed_array[0], 0);
};

uintptr_t gd_packed_string_array_access(uintptr_t a1, uintptr_t a2, INT i) {
    uintptr_t packed_array[2] = {a1, a2};
    return *(uintptr_t*)gdextension_packed_string_array_operator_index_const(&packed_array[0], i);
};

void gd_array_set(uintptr_t a, INT i, UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t *variant = (uint64_t*)gdextension_array_operator_index(&a, i);
    variant[0] = UINT64_FROM(v1);
    variant[1] = UINT64_FROM(v2);
    variant[2] = UINT64_FROM(v3);
};

void gd_array_get(uintptr_t a, INT i, ANY result) {
    uint64_t *packed = (uint64_t*)gdextension_array_operator_index_const(&a, i);
    memcpy((uint64_t*)result, packed, sizeof(uint64_t)*3);
};

uintptr_t gd_packed_vector2_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_vector2_array_operator_index(&packed_array[0], 0);
};

void gd_packed_vector2_array_access(uintptr_t a1, uintptr_t a2, INT i, ANY result) {
    uintptr_t packed_array[2] = {a1, a2};
    const float *vector = (const float *)gdextension_packed_vector2_array_operator_index_const(&packed_array[0], i);
    memcpy((float*)result, vector, sizeof(float)*2);
};

uintptr_t gd_packed_vector3_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_vector3_array_operator_index(&packed_array[0], 0);
};

void gd_packed_vector3_array_access(uintptr_t a1, uintptr_t a2, INT i, ANY result) {
    uintptr_t packed_array[2] = {a1, a2};
    const float *vector = (const float *)gdextension_packed_vector3_array_operator_index_const(&packed_array[0], i);
    memcpy((float*)result, vector, sizeof(float)*3);
};

uintptr_t gd_packed_vector4_array_unsafe(uintptr_t a1, uintptr_t a2) {
    uintptr_t packed_array[2] = {a1, a2};
    return (uintptr_t)gdextension_packed_vector4_array_operator_index(&packed_array[0], 0);
};

void gd_packed_vector4_array_access(uintptr_t a1, uintptr_t a2, INT i, ANY result) {
    uintptr_t packed_array[2] = {a1, a2};
    const float *vector = (const float *)gdextension_packed_vector4_array_operator_index_const(&packed_array[0], i);
    memcpy((float *)result, vector, sizeof(float)*4);
};

uintptr_t gd_ref_get_object(uintptr_t ref) {
    return (uintptr_t)gdextension_ref_get_object((GDExtensionObjectPtr)ref);
};

void gd_ref_set_object(uintptr_t ref, uintptr_t obj) {
    gdextension_ref_set_object((GDExtensionObjectPtr)ref, (GDExtensionObjectPtr)obj);
};

int32_t gd_string_access(uintptr_t s, INT i) {
    return *(int32_t *)gdextension_string_operator_index_const((GDExtensionStringPtr)s, i);
};

uintptr_t gd_string_resize(uintptr_t s, INT size) {
    uintptr_t ptr = s;
    gdextension_string_resize((GDExtensionStringPtr)&ptr, size);
    return ptr;
};

uintptr_t gd_string_unsafe(uintptr_t s) {
    return (uintptr_t)gdextension_string_operator_index((GDExtensionStringPtr)s, 0);
};

uintptr_t gd_string_append(uintptr_t s, uintptr_t other) {
    uintptr_t ptr = s;
    gdextension_string_operator_plus_eq_string((GDExtensionStringPtr)&ptr, (GDExtensionConstStringPtr)&other);
    return ptr;
};

uintptr_t gd_string_append_rune(uintptr_t s, int32_t rune) {
    uintptr_t ptr = s;
    gdextension_string_operator_plus_eq_char((GDExtensionStringPtr)&ptr, rune);
    return ptr;
};

uintptr_t gd_string_decode_latin1(STRING s, INT len) {
    uintptr_t ptr = 0;
    gdextension_string_new_with_latin1_chars_and_len((GDExtensionUninitializedStringPtr)&ptr, STRING_POINTER(s), len);
    return ptr;
};

uintptr_t gd_string_decode_utf8(STRING s, INT len) {
    uintptr_t ptr = 0;
    gdextension_string_new_with_utf8_chars_and_len((GDExtensionUninitializedStringPtr)&ptr, STRING_POINTER(s), len);
    return ptr;
};

uintptr_t gd_string_decode_utf16(STRING s, INT len, bool le) {
    uintptr_t ptr = 0;
    gdextension_string_new_with_utf16_chars_and_len2((GDExtensionUninitializedStringPtr)&ptr, (const char16_t *)STRING_POINTER(s), len, le);
    return ptr;
};

uintptr_t gd_string_decode_utf32(STRING s, INT len) {
    uintptr_t ptr = 0;
    gdextension_string_new_with_utf32_chars_and_len((GDExtensionUninitializedStringPtr)&ptr, (const char32_t *)STRING_POINTER(s), len);
    return ptr;
};

uintptr_t gd_string_decode_wide(STRING s, INT len) {
    uintptr_t ptr = 0;
    gdextension_string_new_with_wide_chars_and_len((GDExtensionUninitializedStringPtr)&ptr, (const wchar_t *)STRING_POINTER(s), len);
    return ptr;
};

INT gd_string_encode_latin1(uintptr_t s, BUFFER addr, INT len) {
    return gdextension_string_to_latin1_chars((GDExtensionStringPtr)&s, BUFFER_POINTER(addr), len);
};

INT gd_string_encode_utf8(uintptr_t s, BUFFER addr, INT len) {
    return gdextension_string_to_utf8_chars((GDExtensionStringPtr)&s, BUFFER_POINTER(addr), len);
};

INT gd_string_encode_utf16(uintptr_t s, BUFFER addr, INT len) {
    return gdextension_string_to_utf16_chars((GDExtensionStringPtr)&s, (char16_t *)BUFFER_POINTER(addr), len/2);
};

INT gd_string_encode_utf32(uintptr_t s, BUFFER addr, INT len) {
    return gdextension_string_to_utf32_chars((GDExtensionStringPtr)&s, (char32_t *)BUFFER_POINTER(addr), len/4);
};

INT gd_string_encode_wide(uintptr_t s, BUFFER addr, INT len) {
    return gdextension_string_to_wide_chars((GDExtensionStringPtr)&s, (wchar_t *)BUFFER_POINTER(addr), len/sizeof(wchar_t));
};

uintptr_t gd_string_intern_latin1(STRING s, INT len) {
    char buf[256]; const char *s_ptr = fit_string(STRING_POINTER(s), len, &buf[0], 256);
    uintptr_t ptr = 0;
    gdextension_string_name_new_with_latin1_chars((GDExtensionUninitializedStringNamePtr)&ptr, s_ptr, false);
    if (s_ptr != &buf[0]) gdextension_mem_free((void *)s_ptr); // Free the buffer if it was allocated
    return ptr;
};

uintptr_t gd_string_intern_utf8(STRING s, INT len) {
    uintptr_t ptr = 0;
    gdextension_string_name_new_with_utf8_chars_and_len((GDExtensionUninitializedStringNamePtr)&ptr, STRING_POINTER(s), len);
    return ptr;
};

uintptr_t gd_variant_type_name(uint32_t vtype) {
    uintptr_t name = 0;
    gdextension_variant_get_type_name((GDExtensionVariantType)vtype, &name);
    return name;
};

void gd_variant_unsafe_make_native(uint32_t vtype, UINT64 v1, UINT64 v2, UINT64 v3, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    type_from_variant_constructors[vtype](points[0], &self[0]);
};

void gd_variant_unsafe_from_native(uint32_t vtype, void *result, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    variant_from_type_constructors[vtype](result, points[0]);
};

void gd_variant_type_call(uint32_t vtype, uintptr_t name, void *result, INT argc, ANY args, ANY err) {
    void *points[16]; prepare_variants(&points[0], argc, args);
    gdextension_variant_call_static((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};

bool gd_variant_type_convertable(uint32_t a, uint32_t b, bool strict) {
    if (strict) {
        return gdextension_variant_can_convert_strict((GDExtensionVariantType)a, (GDExtensionVariantType)b);
    } else {
        return gdextension_variant_can_convert((GDExtensionVariantType)a, (GDExtensionVariantType)b);
    }
};

void gd_variant_type_setup_array(uintptr_t array, uint32_t vtype, uintptr_t class_name, UINT64 s1, UINT64 s2, UINT64 s3) {
    uint64_t script[3] = {UINT64_FROM(s1), UINT64_FROM(s2), UINT64_FROM(s3)};
    gdextension_array_set_typed(&array, (GDExtensionVariantType)vtype, &class_name, &script[0]);
};

void gd_variant_type_setup_dictionary(uintptr_t dict, uint32_t ktype, uintptr_t kclass_name, UINT64 ks1, UINT64 ks2, UINT64 ks3, uint32_t vtype, uintptr_t vclass_name, UINT64 vs1, UINT64 vs2, UINT64 vs3) {
    uint64_t k_script[3] = {UINT64_FROM(ks1), UINT64_FROM(ks2), UINT64_FROM(ks3)};
    uint64_t v_script[3] = {UINT64_FROM(vs1), UINT64_FROM(vs2), UINT64_FROM(vs3)};
    gdextension_dictionary_set_typed(&dict, (GDExtensionVariantType)ktype, &kclass_name, &k_script[0], (GDExtensionVariantType)vtype, &vclass_name, &v_script[0]);
};

void gd_variant_type_fetch_constant(uint32_t vtype, uintptr_t name, ANY result) {
    void *points[16]; prepare_variants(&points[0], 1, result);
    gdextension_variant_get_constant_value((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name, (GDExtensionTypePtr)points[0]);
};

uintptr_t gd_variant_type_builtin_method(uint32_t vtype, uintptr_t name, INT64 hash) {
    return (uintptr_t)gdextension_variant_get_ptr_builtin_method((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name, INT64_FROM(hash));
};

uintptr_t gd_variant_type_unsafe_constructor(uint32_t vtype, INT n) {
    return (uintptr_t)gdextension_variant_get_ptr_constructor((GDExtensionVariantType)vtype, n);
};

void gd_variant_type_unsafe_free(uint32_t vtype, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    variant_ptr_destructors[vtype](points[0]);
};

void gd_variant_zero(ANY result) {
    gdextension_variant_new_nil((GDExtensionUninitializedVariantPtr)result);
};

void gd_variant_copy(UINT64 v1, UINT64 v2, UINT64 v3, ANY result) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    gdextension_variant_new_copy((GDExtensionUninitializedVariantPtr)result, &self);
};

void gd_variant_call(UINT64 v1, UINT64 v2, UINT64 v3, uintptr_t name, void *result, INT argc, ANY args, void *err) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    void *points[16]; prepare_variants(&points[0], argc, args);
    gdextension_variant_call(&self, (GDExtensionConstStringNamePtr)&name, (const GDExtensionConstTypePtr*)&points[0], argc, (GDExtensionTypePtr)result, (GDExtensionCallError*)err);
};

bool gd_variant_eval(uint32_t op, UINT64 a1, UINT64 a2, UINT64 a3, UINT64 b1, UINT64 b2, UINT64 b3, ANY result) {
    uint64_t a[3] = {UINT64_FROM(a1), UINT64_FROM(a2), UINT64_FROM(a3)};
    uint64_t b[3] = {UINT64_FROM(b1), UINT64_FROM(b2), UINT64_FROM(b3)};
    GDExtensionBool valid = false;
    gdextension_variant_evaluate((GDExtensionVariantOperator)op, &a[0], &b[0], (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};

INT64 gd_variant_hash(UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    return INT64_MAKE(gdextension_variant_hash(&self));
};

bool gd_variant_bool(UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    return gdextension_variant_booleanize(&self);
};

uintptr_t gd_variant_text(UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    uintptr_t text = 0;
    gdextension_variant_stringify(&self, &text);
    return text;
};

uint32_t gd_variant_type(UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    return gdextension_variant_get_type(&self);
};

void gd_variant_deep_copy(UINT64 v1, UINT64 v2, UINT64 v3, void *result) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    gdextension_variant_duplicate(&self, result, true);
};

int64_t gd_variant_deep_hash(UINT64 v1, UINT64 v2, UINT64 v3, INT recursion) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    return gdextension_variant_recursive_hash(&self, recursion);
};

bool gd_variant_get_index(UINT64 v1, UINT64 v2, UINT64 v3, UINT64 i1, UINT64 i2, UINT64 i3, ANY result) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    uint64_t index[3] = {UINT64_FROM(i1), UINT64_FROM(i2), UINT64_FROM(i3)};
    GDExtensionBool valid = false;
    gdextension_variant_get_keyed(&self, &index[0], (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};

bool gd_variant_get_array(UINT64 v1, UINT64 v2, UINT64 v3, INT i, ANY result, ANY err) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    GDExtensionBool valid = false;
    GDExtensionBool oob = false;
    gdextension_variant_get_indexed(&self, i, (GDExtensionUninitializedVariantPtr)result, &valid, &oob);
    if (oob) ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return valid;
};

bool gd_variant_get_field(UINT64 v1, UINT64 v2, UINT64 v3, uintptr_t name, ANY result) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    GDExtensionBool valid = false;
    gdextension_variant_get_named(&self, (GDExtensionConstStringNamePtr)&name, (GDExtensionUninitializedVariantPtr)result, &valid);
    return valid;
};

bool gd_variant_type_has_property(uint32_t vtype, uintptr_t name) {
    return gdextension_variant_has_member((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name);
};

bool gd_variant_has_index(UINT64 v1, UINT64 v2, UINT64 v3, UINT64 i1, UINT64 i2, UINT64 i3) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    uint64_t index[3] = {UINT64_FROM(i1), UINT64_FROM(i2), UINT64_FROM(i3)};
    GDExtensionBool valid = false;
    gdextension_variant_has_key(&self, &index[0], &valid);
    return valid;
};

bool gd_variant_has_method(UINT64 v1, UINT64 v2, UINT64 v3, uintptr_t name) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    return gdextension_variant_has_method(&self, (GDExtensionConstStringNamePtr)&name);
};

bool gd_variant_set_index(UINT64 v1, UINT64 v2, UINT64 v3, UINT64 i1, UINT64 i2, UINT64 i3, UINT64 v1_set, UINT64 v2_set, UINT64 v3_set) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    uint64_t index[3] = {UINT64_FROM(i1), UINT64_FROM(i2), UINT64_FROM(i3)};
    uint64_t value[3] = {UINT64_FROM(v1_set), UINT64_FROM(v2_set), UINT64_FROM(v3_set)};
    GDExtensionBool valid = false;
    gdextension_variant_set_keyed(&self, &index[0], &value[0], &valid);
    return valid;
};

bool gd_variant_set_array(UINT64 v1, UINT64 v2, UINT64 v3, INT i, UINT64 v1_set, UINT64 v2_set, UINT64 v3_set, ANY err) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    uint64_t value[3] = {UINT64_FROM(v1_set), UINT64_FROM(v2_set), UINT64_FROM(v3_set)};
    GDExtensionBool valid = false;
    GDExtensionBool oob = false;
    gdextension_variant_set_indexed(&self, i, &value[0], &valid, &oob);
    if (oob) ((GDExtensionCallError*)err)->error = GDEXTENSION_CALL_ERROR_INVALID_ARGUMENT;
    return valid;
};

bool gd_variant_set_field(UINT64 v1, UINT64 v2, UINT64 v3, uintptr_t name, UINT64 v1_set, UINT64 v2_set, UINT64 v3_set) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    uint64_t value[3] = {UINT64_FROM(v1_set), UINT64_FROM(v2_set), UINT64_FROM(v3_set)};
    GDExtensionBool valid = false;
    gdextension_variant_set_named(&self, (GDExtensionConstStringNamePtr)&name, &value[0], &valid);
    return valid;
};

void gd_variant_unsafe_eval(uintptr_t fn, void *result, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    ((GDExtensionPtrOperatorEvaluator)fn)(points[0], points[1], result);
};

void gd_variant_unsafe_call(uintptr_t fn, void *result, UINT64 shape, ANY args) {
    void *points[16]; uint8_t argc = prepare_callframe(&points[0], shape, args);
    ((GDExtensionPtrBuiltInMethod)fn)((GDExtensionTypePtr*)points[0], (const GDExtensionConstTypePtr*)&points[1], (GDExtensionTypePtr)result, argc);
};

void gd_variant_unsafe_free(UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    gdextension_variant_destroy(&self);
};

uintptr_t gd_variant_unsafe_internal_pointer(uint32_t vtype, UINT64 v1, UINT64 v2, UINT64 v3) {
    uint64_t self[3] = {UINT64_FROM(v1), UINT64_FROM(v2), UINT64_FROM(v3)};
    return (uintptr_t)variant_internal_ptr_funcs[vtype](&self);
};

void gd_variant_unsafe_get_field(uintptr_t getter, void *result, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    ((GDExtensionPtrGetter)getter)(points[0], &result);

};

void gd_variant_unsafe_get_array(uint32_t vtype, INT i, void *result, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    variant_ptr_indexed_getters[vtype](result, i, points[0]);
};

void gd_variant_unsafe_get_index(uint32_t vtype, void *result, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    variant_ptr_keyed_getters[vtype](points[0], points[1], result);
};

void gd_variant_unsafe_set_field(uintptr_t setter, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    ((GDExtensionPtrSetter)setter)(points[1], points[2]);
};

void gd_variant_unsafe_set_array(uint32_t vtype, INT i, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    variant_ptr_indexed_setters[vtype](points[0], i, points[1]);
};

void gd_variant_unsafe_set_index(uint32_t vtype, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    variant_ptr_keyed_setters[vtype](points[0], points[1], points[2]);
};

uintptr_t gd_variant_type_setter(uint32_t vtype, uintptr_t name) {
    return (uintptr_t)gdextension_variant_get_ptr_setter((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name);
}

uintptr_t gd_variant_type_getter(uint32_t vtype, uintptr_t name) {
    return (uintptr_t)gdextension_variant_get_ptr_getter((GDExtensionVariantType)vtype, (GDExtensionConstStringNamePtr)&name);
}

void gd_variant_type_unsafe_make(uintptr_t fn, void* result, UINT64 shape, ANY args) {
    void *points[16]; prepare_callframe(&points[0], shape, args);
    ((GDExtensionPtrConstructor)fn)(result, (const GDExtensionConstTypePtr *)&points[0]);
}

void gd_variant_type_make(uint32_t vtype, void* result, INT argc, ANY args, ANY err) {
    void *points[16]; prepare_variants(&points[0], argc, args);
    gdextension_variant_construct((GDExtensionVariantType)vtype, result, (const GDExtensionConstVariantPtr *)&points[0], argc, (GDExtensionCallError*)err);
}

void gd_variant_type_unsafe_call(uintptr_t fn, void* result, UINT64 shape, ANY args) {
    void *points[16]; uint8_t argc = prepare_callframe(&points[0], shape, args);
    ((GDExtensionPtrBuiltInMethod)fn)((GDExtensionTypePtr*)points[0], (const GDExtensionConstTypePtr*)&points[1], (GDExtensionTypePtr)result, argc);
}

uintptr_t gd_variant_type_evaluator(uint32_t op, uint32_t a, uint32_t b) {
    return (uintptr_t)gdextension_variant_get_ptr_operator_evaluator((GDExtensionVariantOperator)op, (GDExtensionVariantType)a, (GDExtensionVariantType)b);
}

uint32_t gd_version_major() {return cgo_cached_godot_version.major;};
uint32_t gd_version_minor() {return cgo_cached_godot_version.minor;};
uint32_t gd_version_patch() {return cgo_cached_godot_version.patch;};
uint32_t gd_version_hex() {return cgo_cached_godot_version.hex;};
uintptr_t gd_version_status() {return (uintptr_t)cgo_cached_godot_version.status;};
uintptr_t gd_version_build() {return (uintptr_t)cgo_cached_godot_version.build;};
uintptr_t gd_version_hash() {return (uintptr_t)cgo_cached_godot_version.hash;};
uint64_t gd_version_timestamp() {return cgo_cached_godot_version.timestamp;};
uintptr_t gd_version_string() {return (uintptr_t)cgo_cached_godot_version.string;};

#ifdef __EMSCRIPTEN__
using namespace emscripten;
EMSCRIPTEN_BINDINGS(my_module) {
	function("gd_builtin_name", &gd_builtin_name, allow_raw_pointers());
	function("gd_builtin_call", &gd_builtin_call, allow_raw_pointers());
	function("gd_callable_create", &gd_callable_create, allow_raw_pointers());
	function("gd_callable_lookup", &gd_callable_lookup, allow_raw_pointers());
	function("gd_classdb_FileAccess_write", &gd_classdb_FileAccess_write, allow_raw_pointers());
	function("gd_classdb_FileAccess_read", &gd_classdb_FileAccess_read, allow_raw_pointers());
	function("gd_classdb_Image_unsafe", &gd_classdb_Image_unsafe, allow_raw_pointers());
	function("gd_classdb_Image_access", &gd_classdb_Image_access, allow_raw_pointers());
	function("gd_method_list_make", &gd_method_list_make, allow_raw_pointers());
	function("gd_method_list_push", &gd_method_list_push, allow_raw_pointers());
	function("gd_method_list_free", &gd_method_list_free, allow_raw_pointers());
	function("gd_property_list_make", &gd_property_list_make, allow_raw_pointers());
	function("gd_property_list_push", &gd_property_list_push, allow_raw_pointers());
	function("gd_property_list_free", &gd_property_list_free, allow_raw_pointers());
	function("gd_property_info_type", &gd_property_info_type, allow_raw_pointers());
	function("gd_property_info_name", &gd_property_info_name, allow_raw_pointers());
	function("gd_property_info_class_name", &gd_property_info_class_name, allow_raw_pointers());
	function("gd_property_info_hint", &gd_property_info_hint, allow_raw_pointers());
	function("gd_property_info_hint_string", &gd_property_info_hint_string, allow_raw_pointers());
	function("gd_property_info_usage", &gd_property_info_usage, allow_raw_pointers());
	function("gd_classdb_register", &gd_classdb_register, allow_raw_pointers());
	function("gd_classdb_register_methods", &gd_classdb_register_methods, allow_raw_pointers());
	function("gd_classdb_register_constant", &gd_classdb_register_constant, allow_raw_pointers());
	function("gd_classdb_register_property", &gd_classdb_register_property, allow_raw_pointers());
	function("gd_classdb_register_property_indexed", &gd_classdb_register_property_indexed, allow_raw_pointers());
	function("gd_classdb_register_property_group", &gd_classdb_register_property_group, allow_raw_pointers());
	function("gd_classdb_register_property_sub_group", &gd_classdb_register_property_sub_group, allow_raw_pointers());
	function("gd_classdb_register_signal", &gd_classdb_register_signal, allow_raw_pointers());
	function("gd_classdb_register_removal", &gd_classdb_register_removal, allow_raw_pointers());
	function("gd_classdb_WorkerThreadPool_add_task", &gd_classdb_WorkerThreadPool_add_task, allow_raw_pointers());
	function("gd_classdb_WorkerThreadPool_add_group_task", &gd_classdb_WorkerThreadPool_add_group_task, allow_raw_pointers());
	function("gd_classdb_XMLParser_load", &gd_classdb_XMLParser_load, allow_raw_pointers());
	function("gd_packed_dictionary_access", &gd_packed_dictionary_access, allow_raw_pointers());
	function("gd_packed_dictionary_modify", &gd_packed_dictionary_modify, allow_raw_pointers());
	function("gd_editor_add_documentation", &gd_editor_add_documentation, allow_raw_pointers());
	function("gd_editor_add_plugin", &gd_editor_add_plugin, allow_raw_pointers());
	function("gd_editor_end_plugin", &gd_editor_end_plugin, allow_raw_pointers());
	function("gd_iterator_make", &gd_iterator_make, allow_raw_pointers());
	function("gd_iterator_next", &gd_iterator_next, allow_raw_pointers());
	function("gd_iterator_load", &gd_iterator_load, allow_raw_pointers());
	function("gd_log_error", &gd_log_error, allow_raw_pointers());
	function("gd_log_warning", &gd_log_warning, allow_raw_pointers());
	function("gd_memory_malloc", &gd_memory_malloc, allow_raw_pointers());
	function("gd_memory_sizeof", &gd_memory_sizeof, allow_raw_pointers());
	function("gd_memory_resize", &gd_memory_resize, allow_raw_pointers());
	function("gd_memory_free", &gd_memory_free, allow_raw_pointers());
	function("gd_memory_edit_byte", &gd_memory_edit_byte, allow_raw_pointers());
	function("gd_memory_edit_u16", &gd_memory_edit_u16, allow_raw_pointers());
	function("gd_memory_edit_u32", &gd_memory_edit_u32, allow_raw_pointers());
	function("gd_memory_edit_u64", &gd_memory_edit_u64, allow_raw_pointers());
	function("gd_memory_edit_128", &gd_memory_edit_128, allow_raw_pointers());
	function("gd_memory_edit_256", &gd_memory_edit_256, allow_raw_pointers());
	function("gd_memory_edit_512", &gd_memory_edit_512, allow_raw_pointers());
	function("gd_memory_load_byte", &gd_memory_load_byte, allow_raw_pointers());
	function("gd_memory_load_u16", &gd_memory_load_u16, allow_raw_pointers());
	function("gd_memory_load_u32", &gd_memory_load_u32, allow_raw_pointers());
	function("gd_memory_load_u64", &gd_memory_load_u64, allow_raw_pointers());
	function("gd_object_make", &gd_object_make, allow_raw_pointers());
	function("gd_object_call", &gd_object_call, allow_raw_pointers());
	function("gd_object_name", &gd_object_name, allow_raw_pointers());
	function("gd_object_type", &gd_object_type, allow_raw_pointers());
	function("gd_object_cast", &gd_object_cast, allow_raw_pointers());
	function("gd_object_lookup", &gd_object_lookup, allow_raw_pointers());
	function("gd_object_global", &gd_object_global, allow_raw_pointers());
	function("gd_object_extension_setup", &gd_object_extension_setup, allow_raw_pointers());
	function("gd_object_extension_fetch", &gd_object_extension_fetch, allow_raw_pointers());
	function("gd_object_extension_close", &gd_object_extension_close, allow_raw_pointers());
	function("gd_object_id", &gd_object_id, allow_raw_pointers());
	function("gd_object_id_inside_variant", &gd_object_id_inside_variant, allow_raw_pointers());
	function("gd_object_method_lookup", &gd_object_method_lookup, allow_raw_pointers());
	function("gd_object_script_make", &gd_object_script_make, allow_raw_pointers());
	function("gd_object_script_call", &gd_object_script_call, allow_raw_pointers());
	function("gd_object_script_setup", &gd_object_script_setup, allow_raw_pointers());
	function("gd_object_script_fetch", &gd_object_script_fetch, allow_raw_pointers());
	function("gd_object_script_defines_method", &gd_object_script_defines_method, allow_raw_pointers());
	function("gd_object_script_property_state_add", &gd_object_script_property_state_add, allow_raw_pointers());
	function("gd_object_script_placeholder_create", &gd_object_script_placeholder_create, allow_raw_pointers());
	function("gd_object_script_placeholder_update", &gd_object_script_placeholder_update, allow_raw_pointers());
	function("gd_object_unsafe_call", &gd_object_unsafe_call, allow_raw_pointers());
	function("gd_object_unsafe_free", &gd_object_unsafe_free, allow_raw_pointers());
	function("gd_packed_byte_array_unsafe", &gd_packed_byte_array_unsafe, allow_raw_pointers());
	function("gd_packed_byte_array_access", &gd_packed_byte_array_access, allow_raw_pointers());
	function("gd_packed_color_array_unsafe", &gd_packed_color_array_unsafe, allow_raw_pointers());
	function("gd_packed_color_array_access", &gd_packed_color_array_access, allow_raw_pointers());
	function("gd_packed_float32_array_unsafe", &gd_packed_float32_array_unsafe, allow_raw_pointers());
	function("gd_packed_float32_array_access", &gd_packed_float32_array_access, allow_raw_pointers());
	function("gd_packed_float64_array_unsafe", &gd_packed_float64_array_unsafe, allow_raw_pointers());
	function("gd_packed_float64_array_access", &gd_packed_float64_array_access, allow_raw_pointers());
	function("gd_packed_int32_array_unsafe", &gd_packed_int32_array_unsafe, allow_raw_pointers());
	function("gd_packed_int32_array_access", &gd_packed_int32_array_access, allow_raw_pointers());
	function("gd_packed_int64_array_unsafe", &gd_packed_int64_array_unsafe, allow_raw_pointers());
	function("gd_packed_int64_array_access", &gd_packed_int64_array_access, allow_raw_pointers());
	function("gd_packed_string_array_unsafe", &gd_packed_string_array_unsafe, allow_raw_pointers());
	function("gd_packed_string_array_access", &gd_packed_string_array_access, allow_raw_pointers());
	function("gd_array_set", &gd_array_set, allow_raw_pointers());
	function("gd_array_get", &gd_array_get, allow_raw_pointers());
	function("gd_packed_vector2_array_unsafe", &gd_packed_vector2_array_unsafe, allow_raw_pointers());
	function("gd_packed_vector2_array_access", &gd_packed_vector2_array_access, allow_raw_pointers());
	function("gd_packed_vector3_array_unsafe", &gd_packed_vector3_array_unsafe, allow_raw_pointers());
	function("gd_packed_vector3_array_access", &gd_packed_vector3_array_access, allow_raw_pointers());
	function("gd_packed_vector4_array_unsafe", &gd_packed_vector4_array_unsafe, allow_raw_pointers());
	function("gd_packed_vector4_array_access", &gd_packed_vector4_array_access, allow_raw_pointers());
	function("gd_ref_get_object", &gd_ref_get_object, allow_raw_pointers());
	function("gd_ref_set_object", &gd_ref_set_object, allow_raw_pointers());
	function("gd_string_access", &gd_string_access, allow_raw_pointers());
	function("gd_string_resize", &gd_string_resize, allow_raw_pointers());
	function("gd_string_unsafe", &gd_string_unsafe, allow_raw_pointers());
	function("gd_string_append", &gd_string_append, allow_raw_pointers());
	function("gd_string_append_rune", &gd_string_append_rune, allow_raw_pointers());
	function("gd_string_decode_latin1", &gd_string_decode_latin1, allow_raw_pointers());
	function("gd_string_decode_utf8", &gd_string_decode_utf8, allow_raw_pointers());
	function("gd_string_decode_utf16", &gd_string_decode_utf16, allow_raw_pointers());
	function("gd_string_decode_utf32", &gd_string_decode_utf32, allow_raw_pointers());
	function("gd_string_decode_wide", &gd_string_decode_wide, allow_raw_pointers());
	function("gd_string_encode_latin1", &gd_string_encode_latin1, allow_raw_pointers());
	function("gd_string_encode_utf8", &gd_string_encode_utf8, allow_raw_pointers());
	function("gd_string_encode_utf16", &gd_string_encode_utf16, allow_raw_pointers());
	function("gd_string_encode_utf32", &gd_string_encode_utf32, allow_raw_pointers());
	function("gd_string_encode_wide", &gd_string_encode_wide, allow_raw_pointers());
	function("gd_string_intern_latin1", &gd_string_intern_latin1, allow_raw_pointers());
	function("gd_string_intern_utf8", &gd_string_intern_utf8, allow_raw_pointers());
	function("gd_variant_type_name", &gd_variant_type_name, allow_raw_pointers());
	function("gd_variant_type_make", &gd_variant_type_make, allow_raw_pointers());
	function("gd_variant_type_call", &gd_variant_type_call, allow_raw_pointers());
	function("gd_variant_type_convertable", &gd_variant_type_convertable, allow_raw_pointers());
	function("gd_variant_type_setup_array", &gd_variant_type_setup_array, allow_raw_pointers());
	function("gd_variant_type_setup_dictionary", &gd_variant_type_setup_dictionary, allow_raw_pointers());
	function("gd_variant_type_fetch_constant", &gd_variant_type_fetch_constant, allow_raw_pointers());
	function("gd_variant_type_builtin_method", &gd_variant_type_builtin_method, allow_raw_pointers());
	function("gd_variant_type_unsafe_constructor", &gd_variant_type_unsafe_constructor, allow_raw_pointers());
	function("gd_variant_type_evaluator", &gd_variant_type_evaluator, allow_raw_pointers());
	function("gd_variant_type_setter", &gd_variant_type_setter, allow_raw_pointers());
	function("gd_variant_type_getter", &gd_variant_type_getter, allow_raw_pointers());
	function("gd_variant_type_has_property", &gd_variant_type_has_property, allow_raw_pointers());
	function("gd_variant_type_unsafe_call", &gd_variant_type_unsafe_call, allow_raw_pointers());
	function("gd_variant_type_unsafe_make", &gd_variant_type_unsafe_make, allow_raw_pointers());
	function("gd_variant_type_unsafe_free", &gd_variant_type_unsafe_free, allow_raw_pointers());
	function("gd_variant_zero", &gd_variant_zero, allow_raw_pointers());
	function("gd_variant_copy", &gd_variant_copy, allow_raw_pointers());
	function("gd_variant_call", &gd_variant_call, allow_raw_pointers());
	function("gd_variant_eval", &gd_variant_eval, allow_raw_pointers());
	function("gd_variant_hash", &gd_variant_hash, allow_raw_pointers());
	function("gd_variant_bool", &gd_variant_bool, allow_raw_pointers());
	function("gd_variant_text", &gd_variant_text, allow_raw_pointers());
	function("gd_variant_type", &gd_variant_type, allow_raw_pointers());
	function("gd_variant_deep_copy", &gd_variant_deep_copy, allow_raw_pointers());
	function("gd_variant_deep_hash", &gd_variant_deep_hash, allow_raw_pointers());
	function("gd_variant_get_index", &gd_variant_get_index, allow_raw_pointers());
	function("gd_variant_get_array", &gd_variant_get_array, allow_raw_pointers());
	function("gd_variant_get_field", &gd_variant_get_field, allow_raw_pointers());
	function("gd_variant_has_index", &gd_variant_has_index, allow_raw_pointers());
	function("gd_variant_has_method", &gd_variant_has_method, allow_raw_pointers());
	function("gd_variant_set_index", &gd_variant_set_index, allow_raw_pointers());
	function("gd_variant_set_array", &gd_variant_set_array, allow_raw_pointers());
	function("gd_variant_set_field", &gd_variant_set_field, allow_raw_pointers());
	function("gd_variant_unsafe_call", &gd_variant_unsafe_call, allow_raw_pointers());
	function("gd_variant_unsafe_eval", &gd_variant_unsafe_eval, allow_raw_pointers());
	function("gd_variant_unsafe_free", &gd_variant_unsafe_free, allow_raw_pointers());
	function("gd_variant_unsafe_make_native", &gd_variant_unsafe_make_native, allow_raw_pointers());
	function("gd_variant_unsafe_from_native", &gd_variant_unsafe_from_native, allow_raw_pointers());
	function("gd_variant_unsafe_internal_pointer", &gd_variant_unsafe_internal_pointer, allow_raw_pointers());
	function("gd_variant_unsafe_get_field", &gd_variant_unsafe_get_field, allow_raw_pointers());
	function("gd_variant_unsafe_get_array", &gd_variant_unsafe_get_array, allow_raw_pointers());
	function("gd_variant_unsafe_get_index", &gd_variant_unsafe_get_index, allow_raw_pointers());
	function("gd_variant_unsafe_set_field", &gd_variant_unsafe_set_field, allow_raw_pointers());
	function("gd_variant_unsafe_set_array", &gd_variant_unsafe_set_array, allow_raw_pointers());
	function("gd_variant_unsafe_set_index", &gd_variant_unsafe_set_index, allow_raw_pointers());
	function("gd_version_major", &gd_version_major, allow_raw_pointers());
	function("gd_version_minor", &gd_version_minor, allow_raw_pointers());
	function("gd_version_patch", &gd_version_patch, allow_raw_pointers());
	function("gd_version_hex", &gd_version_hex, allow_raw_pointers());
	function("gd_version_status", &gd_version_status, allow_raw_pointers());
	function("gd_version_build", &gd_version_build, allow_raw_pointers());
	function("gd_version_hash", &gd_version_hash, allow_raw_pointers());
	function("gd_version_timestamp", &gd_version_timestamp, allow_raw_pointers());
	function("gd_version_string", &gd_version_string, allow_raw_pointers());
}
#endif // __EMSCRIPTEN__
