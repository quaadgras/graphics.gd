#include <stdlib.h>
#include "gdextension_interface.h"

typedef uintptr_t pointer;
typedef const char* string;
typedef char32_t rune;
typedef uint8_t * bytes;

#ifdef _WIN32
#define EXPORT __declspec(dllexport)
#else
#define EXPORT
#endif


// TODO the amount of code here can probably be reduced by stenciling out the various
// function signatures, such that we can call every function that takes 1 pointer arg,
// 2 pointer args etc.

extern EXPORT void initialize(void *userdata, GDExtensionInitializationLevel p_level);
extern EXPORT void deinitialize(void *userdata, GDExtensionInitializationLevel p_level);

static inline void initialization(GDExtensionInitialization *p_init) {
	p_init->initialize = initialize;
	p_init->deinitialize = deinitialize;
}

static inline void *get_proc_address(pointer fn, string p_name) {
	return (void *)((GDExtensionInterfaceGetProcAddress)fn)(p_name);
}

static inline void get_godot_version(pointer fn, GDExtensionGodotVersion *r_version) {
	((GDExtensionInterfaceGetGodotVersion)fn)(r_version);
}
static inline void *mem_alloc(pointer fn, size_t size) {
	return ((GDExtensionInterfaceMemAlloc)fn)(size);
}
static inline void *mem_realloc(pointer fn, uintptr_t ptr, size_t size) {
	return ((GDExtensionInterfaceMemRealloc)fn)((void *)ptr, size);
}
static inline void mem_free(pointer fn, uintptr_t ptr) {
	((GDExtensionInterfaceMemFree)fn)((void *)ptr);
}
static inline void print_error(pointer fn, string p_description, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintError)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}
static inline void print_error_with_message(pointer fn, string p_description, string p_message, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintErrorWithMessage)fn)(p_description, p_message, p_function, p_file, p_line, p_notify_editor);
}
static inline void print_warning(pointer fn, string p_description, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintWarning)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}
static inline uint64_t get_native_struct_size(pointer fn, pointer p_classname) {
	return ((GDExtensionInterfaceGetNativeStructSize)fn)((GDExtensionConstStringNamePtr)p_classname);
}
static inline void variant_new_copy(pointer fn, pointer r_dest, pointer p_self) {
	((GDExtensionInterfaceVariantNewCopy)fn)((GDExtensionUninitializedVariantPtr)r_dest, (GDExtensionConstVariantPtr)p_self);
}
static inline void variant_new_nil(pointer fn, pointer r_dest) {
	((GDExtensionInterfaceVariantNewNil)fn)((GDExtensionUninitializedVariantPtr)r_dest);
}
static inline void variant_destroy(pointer fn, pointer p_self) {
	((GDExtensionInterfaceVariantDestroy)fn)((GDExtensionVariantPtr)p_self);
}
static inline void variant_call(pointer fn, pointer p_self, pointer p_method, pointer p_args, GDExtensionInt p_argument_count, pointer r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantCall)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_method, (GDExtensionConstVariantPtr)p_args, (GDExtensionInt)p_argument_count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
static inline void variant_call_static(pointer fn, GDExtensionVariantType p_type, pointer p_method, pointer p_args, GDExtensionInt p_argument_count, pointer r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantCallStatic)fn)((GDExtensionVariantType)p_type, (GDExtensionConstStringNamePtr)p_method, (GDExtensionConstVariantPtr)p_args, (GDExtensionInt)p_argument_count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
static inline void variant_evaluate(pointer fn, GDExtensionVariantOperator p_operator, pointer p_a, pointer p_b, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantEvaluate)fn)(p_operator, (GDExtensionConstVariantPtr)p_a, (GDExtensionConstVariantPtr)p_b, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
static inline void variant_set(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSet)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
static inline void variant_set_named(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSetNamed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
static inline void variant_set_keyed(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSetKeyed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
static inline void variant_set_indexed(pointer fn, pointer p_self, GDExtensionInt p_index, pointer p_value, pointer r_valid, pointer r_oob) {
	((GDExtensionInterfaceVariantSetIndexed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionInt)p_index, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid, (GDExtensionBool*)r_oob);
}
static inline void variant_get(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGet)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
static inline void variant_get_named(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGetNamed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
static inline void variant_get_keyed(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGetKeyed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
static inline void variant_get_indexed(pointer fn, pointer p_self, GDExtensionInt p_index, pointer r_ret, pointer r_valid, pointer r_oob) {
	((GDExtensionInterfaceVariantGetIndexed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionInt)p_index, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid, (GDExtensionBool*)r_oob);
}
static inline void variant_iter_init(pointer fn, pointer p_self, pointer r_iter, pointer r_valid) {
	((GDExtensionInterfaceVariantIterInit)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedVariantPtr)r_iter, (GDExtensionBool*)r_valid);
}
static inline void variant_iter_next(pointer fn, pointer p_self, pointer p_iter, pointer r_valid) {
	((GDExtensionInterfaceVariantIterNext)fn)((GDExtensionVariantPtr)p_self, (GDExtensionVariantPtr)p_iter, (GDExtensionBool*)r_valid);
}
static inline void variant_iter_get(pointer fn, pointer p_self, pointer p_iter, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantIterGet)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionVariantPtr)p_iter, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
static inline GDExtensionInt variant_hash(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantHash)fn)((GDExtensionConstVariantPtr)p_self);
}
static inline GDExtensionInt variant_recursive_hash(pointer fn, pointer p_self, GDExtensionInt p_count) {
	return ((GDExtensionInterfaceVariantRecursiveHash)fn)((GDExtensionConstVariantPtr)p_self, p_count);
}
static inline GDExtensionBool variant_hash_compare(pointer fn, pointer p_self, pointer p_other) {
	return ((GDExtensionInterfaceVariantHashCompare)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_other);
}
static inline GDExtensionBool variant_booleanize(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantBooleanize)fn)((GDExtensionConstVariantPtr)p_self);
}
static inline void variant_duplicate(pointer fn, pointer p_self, pointer r_ret, GDExtensionBool p_deep) {
	((GDExtensionInterfaceVariantDuplicate)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedVariantPtr)r_ret, p_deep);
}
static inline void variant_stringify(pointer fn, pointer p_self, pointer r_ret) {
	((GDExtensionInterfaceVariantStringify)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedStringPtr)r_ret);
}
static inline GDExtensionVariantType variant_get_type(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantGetType)fn)((GDExtensionConstVariantPtr)p_self);
}
static inline GDExtensionBool variant_has_method(pointer fn, pointer p_self, pointer p_method) {
	return ((GDExtensionInterfaceVariantHasMethod)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_method);
}
static inline GDExtensionBool variant_has_member(pointer fn, pointer p_self, pointer p_member) {
	return ((GDExtensionInterfaceVariantHasMember)fn)((GDExtensionVariantType)p_self, (GDExtensionConstStringNamePtr)p_member);
}
static inline GDExtensionBool variant_has_key(pointer fn, pointer p_self, pointer p_key, pointer r_valid) {
	return ((GDExtensionInterfaceVariantHasKey)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionBool*)r_valid);
}
static inline void variant_get_type_name(pointer fn, GDExtensionVariantType p_type, pointer r_ret) {
	((GDExtensionInterfaceVariantGetTypeName)fn)(p_type, (GDExtensionUninitializedStringPtr)r_ret);
}
static inline GDExtensionBool variant_can_convert(pointer fn, pointer p_self, GDExtensionVariantType p_type) {
	return ((GDExtensionInterfaceVariantCanConvert)fn)((GDExtensionVariantType)p_self, p_type);
}
static inline GDExtensionBool variant_can_convert_strict(pointer fn, pointer p_self, GDExtensionVariantType p_type) {
	return ((GDExtensionInterfaceVariantCanConvertStrict)fn)((GDExtensionVariantType)p_self, p_type);
}
static inline pointer get_variant_from_type_constructor(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceGetVariantFromTypeConstructor)fn)(p_type);
}
static inline void call_variant_from_type_constructor(pointer fn, pointer r_ret, pointer r_arg) {
	((GDExtensionVariantFromTypeConstructorFunc)fn)((GDExtensionUninitializedVariantPtr)r_ret, (void*)r_arg);
}
static inline pointer get_variant_to_type_constructor(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceGetVariantToTypeConstructor)fn)(p_type);
}
static inline void call_variant_to_type_constructor(pointer fn, pointer r_ret, const pointer r_arg) {
	((GDExtensionTypeFromVariantConstructorFunc)fn)((GDExtensionUninitializedVariantPtr)r_ret, (void*)r_arg);
}
static inline pointer variant_get_ptr_operator_evaluator(pointer fn, GDExtensionVariantOperator p_operator, GDExtensionVariantType p_type, GDExtensionVariantType p_rhs_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrOperatorEvaluator)fn)(p_operator, p_type, p_rhs_type);
}
static inline void call_variant_ptr_operator_evaluator(pointer fn, pointer r_ret, pointer r_a, pointer r_b) {
	((GDExtensionPtrOperatorEvaluator)fn)((GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionConstVariantPtr)r_a, (void*)r_b);
}
static inline pointer variant_get_ptr_builtin_method(pointer fn, GDExtensionVariantType p_type, pointer p_method, GDExtensionInt p_hash) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrBuiltinMethod)fn)(p_type, (GDExtensionConstStringNamePtr)p_method, p_hash);
}
static inline void call_variant_ptr_builtin_method(pointer fn, pointer p_base, pointer r_arg, pointer r_ret, int count) {
	((GDExtensionPtrBuiltInMethod)fn)((GDExtensionTypePtr)p_base, (const GDExtensionConstTypePtr *)r_arg, (GDExtensionTypePtr)r_ret, count);
}
static inline pointer variant_get_ptr_constructor(pointer fn, GDExtensionVariantType p_type, int32_t p_constructor) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrConstructor)fn)(p_type, p_constructor);
}
static inline void call_variant_ptr_constructor(pointer fn, pointer r_ret, pointer r_arg) {
	((GDExtensionPtrConstructor)fn)((GDExtensionUninitializedVariantPtr)r_ret, (void*)r_arg);
}
static inline pointer variant_get_ptr_destructor(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrDestructor)fn)(p_type);
}
static inline void call_variant_ptr_destructor(pointer fn, pointer r_arg) {
	((GDExtensionPtrDestructor)fn)((GDExtensionVariantPtr)r_arg);
}
static inline void variant_construct(pointer fn, GDExtensionVariantType p_type, pointer p_base, pointer p_args, int32_t p_argument_count, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantConstruct)fn)(p_type, (GDExtensionUninitializedVariantPtr)p_base, (GDExtensionConstVariantPtr)p_args, p_argument_count, r_error);
}
static inline pointer variant_get_ptr_setter(pointer fn, GDExtensionVariantType p_type, pointer p_member) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrSetter)fn)(p_type, (GDExtensionConstStringNamePtr)p_member);
}
static inline void call_variant_ptr_setter(pointer fn, pointer p_base, pointer p_value) {
	((GDExtensionPtrSetter)fn)((GDExtensionTypePtr)p_base, (GDExtensionConstTypePtr)p_value);
}
static inline pointer variant_get_ptr_getter(pointer fn, GDExtensionVariantType p_type, pointer p_member) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrGetter)fn)(p_type, (GDExtensionConstStringNamePtr)p_member);
}
static inline void call_variant_ptr_getter(pointer fn, pointer p_base, pointer r_value) {
	((GDExtensionPtrGetter)fn)((GDExtensionTypePtr)p_base, (GDExtensionTypePtr)r_value);
}
static inline pointer variant_get_ptr_indexed_setter(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrIndexedSetter)fn)(p_type);
}
static inline void call_variant_ptr_indexed_setter(pointer fn, pointer p_base, GDExtensionInt p_index, pointer r_value) {
	((GDExtensionPtrIndexedSetter)fn)((GDExtensionTypePtr)p_base, p_index, (GDExtensionConstTypePtr)r_value);
}
static inline pointer variant_get_ptr_indexed_getter(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrIndexedGetter)fn)(p_type);
}
static inline void call_variant_ptr_indexed_getter(pointer fn, pointer p_base, GDExtensionInt p_index, pointer r_ret) {
	((GDExtensionPtrIndexedGetter)fn)((GDExtensionTypePtr)p_base, p_index, (GDExtensionTypePtr)r_ret);
}
static inline pointer variant_get_ptr_keyed_setter(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrKeyedSetter)fn)(p_type);
}
static inline void call_variant_ptr_keyed_setter(pointer fn, pointer p_base, pointer r_arg, pointer r_value) {
	((GDExtensionPtrKeyedSetter)fn)((GDExtensionTypePtr)p_base, (const GDExtensionConstTypePtr *)r_arg, (GDExtensionConstTypePtr)r_value);
}
static inline pointer variant_get_ptr_keyed_getter(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrKeyedGetter)fn)(p_type);
}
static inline void call_variant_ptr_keyed_getter(pointer fn, pointer p_base, pointer r_arg, pointer r_ret) {
	((GDExtensionPtrKeyedGetter)fn)((GDExtensionTypePtr)p_base, (const GDExtensionConstTypePtr *)r_arg, (GDExtensionTypePtr)r_ret);
}
static inline pointer variant_get_ptr_keyed_checker(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrKeyedChecker)fn)(p_type);
}
static inline uint32_t call_variant_ptr_keyed_checker(pointer fn, pointer p_base, pointer r_arg) {
	return ((GDExtensionPtrKeyedChecker)fn)((GDExtensionTypePtr)p_base, (const GDExtensionConstTypePtr *)r_arg);
}
static inline void variant_get_constant_value(pointer fn, GDExtensionVariantType p_type, pointer p_constant, pointer r_ret) {
	((GDExtensionInterfaceVariantGetConstantValue)fn)(p_type, (GDExtensionConstStringNamePtr)p_constant, (GDExtensionUninitializedVariantPtr)r_ret);
}
static inline pointer variant_get_ptr_utility_function(pointer fn, pointer p_name, GDExtensionInt hash) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrUtilityFunction)fn)((GDExtensionConstStringNamePtr)p_name, hash);
}
static inline void call_variant_ptr_utility_function(pointer fn, pointer r_ret, pointer r_arg, int count) {
	((GDExtensionPtrUtilityFunction)fn)((GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionConstVariantPtr)r_arg, count);
}
static inline void string_new_with_utf8_chars_and_len(pointer fn, pointer r_ret, string p_chars, GDExtensionInt p_len) {
	((GDExtensionInterfaceStringNewWithUtf8CharsAndLen)fn)((GDExtensionUninitializedStringPtr)r_ret, p_chars, p_len);
}
static inline GDExtensionInt string_to_utf8_chars(pointer fn, pointer p_self, char *r_chars, GDExtensionInt p_max) {
	return ((GDExtensionInterfaceStringToUtf8Chars)fn)((GDExtensionConstStringPtr)p_self, r_chars, p_max);
}
static inline rune *string_operator_index(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return ((GDExtensionInterfaceStringOperatorIndex)fn)((GDExtensionStringPtr)p_self, p_index);
}
static inline const rune *string_operator_index_const(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return ((GDExtensionInterfaceStringOperatorIndexConst)fn)((GDExtensionConstStringPtr)p_self, p_index);
}
static inline void string_operator_plus_eq_string(pointer fn, pointer p_self, pointer p_other) {
	((GDExtensionInterfaceStringOperatorPlusEqString)fn)((GDExtensionStringPtr)p_self, (GDExtensionConstStringPtr)p_other);
}
static inline void string_operator_plus_eq_char(pointer fn, pointer p_self, rune p_other) {
	((GDExtensionInterfaceStringOperatorPlusEqChar)fn)((GDExtensionStringPtr)p_self, p_other);
}
static inline void string_resize(pointer fn, pointer p_self, GDExtensionInt p_size) {
	((GDExtensionInterfaceStringResize)fn)((GDExtensionStringPtr)p_self, p_size);
}
static inline void string_name_new_with_utf8_chars_and_len(pointer fn, pointer r_ret, string p_chars, GDExtensionInt p_len) {
	((GDExtensionInterfaceStringNameNewWithUtf8CharsAndLen)fn)((GDExtensionUninitializedStringNamePtr)r_ret, p_chars, p_len);
}
static inline GDExtensionInt xml_parser_open_buffer(pointer fn, pointer p_instance, const bytes p_buffer, size_t p_size) {
	return ((GDExtensionInterfaceXmlParserOpenBuffer)fn)((GDExtensionObjectPtr)p_instance, (const uint8_t *)p_buffer, p_size);
}
static inline void file_access_store_buffer(pointer fn, pointer p_self, const bytes p_buffer, uint64_t p_len) {
	((GDExtensionInterfaceFileAccessStoreBuffer)fn)((GDExtensionObjectPtr)p_self, (const uint8_t *)p_buffer, p_len);
}
static inline uint64_t file_access_get_buffer(pointer fn, pointer p_self, bytes r_buffer, uint64_t p_length) {
	return ((GDExtensionInterfaceFileAccessGetBuffer)fn)((GDExtensionConstObjectPtr)p_self, (uint8_t *)r_buffer, p_length);
}
static inline void *packed_T_operator_index(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return ((GDExtensionInterfacePackedByteArrayOperatorIndex)fn)((GDExtensionTypePtr)p_self, p_index);
}
static inline const void *packed_T_operator_index_const(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return ((GDExtensionInterfacePackedByteArrayOperatorIndexConst)fn)((GDExtensionTypePtr)p_self, p_index);
}
static inline void *array_operator_index(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return ((GDExtensionInterfaceArrayOperatorIndex)fn)((GDExtensionTypePtr)p_self, p_index);
}
static inline const void *array_operator_index_const(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return ((GDExtensionInterfaceArrayOperatorIndexConst)fn)((GDExtensionTypePtr)p_self, p_index);
}
static inline void array_ref(pointer fn, pointer p_self, pointer p_from) {
	((GDExtensionInterfaceArrayRef)fn)((GDExtensionTypePtr)p_self, (GDExtensionConstTypePtr)p_from);
}
static inline void array_set_typed(pointer fn, pointer p_self, GDExtensionVariantType p_type, pointer p_class_name, pointer p_script) {
	((GDExtensionInterfaceArraySetTyped)fn)((GDExtensionTypePtr)p_self, p_type, (GDExtensionConstStringNamePtr)p_class_name, (GDExtensionConstVariantPtr)p_script);
}
static inline void *dictionary_operator_index(pointer fn, pointer p_self, pointer p_key) {
	return ((GDExtensionInterfaceDictionaryOperatorIndex)fn)((GDExtensionTypePtr)p_self, (GDExtensionConstVariantPtr)p_key);
}
static inline void object_method_bind_call(pointer fn, pointer p_method_bind, pointer p_instance, pointer p_args, GDExtensionInt count, pointer r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceObjectMethodBindCall)fn)((GDExtensionMethodBindPtr)p_method_bind, (GDExtensionObjectPtr)p_instance, (GDExtensionConstVariantPtr)p_args, count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
static inline void object_method_bind_ptrcall(pointer fn, pointer p_method_bind, pointer p_instance, pointer p_args, pointer r_ret) {
	((GDExtensionInterfaceObjectMethodBindPtrcall)fn)((GDExtensionMethodBindPtr)p_method_bind, (GDExtensionObjectPtr)p_instance, (GDExtensionConstVariantPtr)p_args, (GDExtensionUninitializedVariantPtr)r_ret);
}
static inline void object_destroy(pointer fn, pointer p_self) {
	((GDExtensionInterfaceObjectDestroy)fn)((GDExtensionObjectPtr)p_self);
}
static inline pointer global_get_singleton(pointer fn, pointer p_name) {
	return (pointer)((GDExtensionInterfaceGlobalGetSingleton)fn)((GDExtensionConstStringNamePtr)p_name);
}
static inline pointer object_get_instance_binding(pointer fn, pointer p_o, pointer p_library, void *p_callbacks) {
	return (pointer)((GDExtensionInterfaceObjectGetInstanceBinding)fn)((GDExtensionObjectPtr)p_o, (void*)p_library, (const GDExtensionInstanceBindingCallbacks *)p_callbacks);
}
static inline void object_set_instance_binding(pointer fn, pointer p_o, pointer p_library, pointer p_binding, void *p_callbacks) {
	((GDExtensionInterfaceObjectSetInstanceBinding)fn)((GDExtensionObjectPtr)p_o, (void*)p_library, (void *)p_binding, (const GDExtensionInstanceBindingCallbacks *)p_callbacks);
}
static inline void object_free_instance_binding(pointer fn, pointer p_o, pointer p_library) {
	((GDExtensionInterfaceObjectFreeInstanceBinding)fn)((GDExtensionObjectPtr)p_o, (void*)p_library);
}
static inline void object_set_instance(pointer fn, pointer p_o, pointer p_classname, pointer p_instance) {
	((GDExtensionInterfaceObjectSetInstance)fn)((GDExtensionObjectPtr)p_o, (GDExtensionConstStringNamePtr)p_classname, (GDExtensionObjectPtr)p_instance);
}
static inline void object_get_class_name(pointer fn, pointer p_o, pointer p_token, pointer r_ret) {
	((GDExtensionInterfaceObjectGetClassName)fn)((GDExtensionObjectPtr)p_o, (void *)p_token, (GDExtensionUninitializedStringPtr)r_ret);
}
static inline pointer object_cast_to(pointer fn, pointer p_o, pointer p_class_tag) {
	return (pointer)((GDExtensionInterfaceObjectCastTo)fn)((GDExtensionObjectPtr)p_o, (void*)p_class_tag);
}
static inline uint64_t object_get_instance_id(pointer fn, pointer p_o) {
	return ((GDExtensionInterfaceObjectGetInstanceId)fn)((GDExtensionObjectPtr)p_o);
}
static inline pointer object_get_instance_from_id(pointer fn, pointer p_id) {
	return (pointer)((GDExtensionInterfaceObjectGetInstanceFromId)fn)(p_id);
}
static inline pointer ref_get_object(pointer fn, pointer p_r) {
	return (pointer)((GDExtensionInterfaceRefGetObject)fn)((GDExtensionRefPtr)p_r);
}
static inline void ref_set_object(pointer fn, pointer p_r, pointer p_o) {
	((GDExtensionInterfaceRefSetObject)fn)((GDExtensionRefPtr)p_r, (GDExtensionObjectPtr)p_o);
}


static inline pointer classdb_construct_object(pointer fn, pointer p_classname) {
	return (pointer)((GDExtensionInterfaceClassdbConstructObject)fn)((GDExtensionConstStringNamePtr)p_classname);
}
static inline pointer classdb_get_class_tag(pointer fn, pointer p_classname) {
	return (pointer)((GDExtensionInterfaceClassdbGetClassTag)fn)((GDExtensionConstStringNamePtr)p_classname);
}
static inline pointer classdb_get_method_bind(pointer fn, pointer p_classname, pointer p_methodname, GDExtensionInt hash) {
	return (pointer)((GDExtensionInterfaceClassdbGetMethodBind)fn)((GDExtensionConstStringNamePtr)p_classname, (GDExtensionConstStringNamePtr)p_methodname, hash);
}
static inline void get_library_path(pointer fn, pointer p_token, pointer r_ret) {
	((GDExtensionInterfaceGetLibraryPath)fn)((void *)p_token, (GDExtensionUninitializedStringPtr)r_ret);
}

extern EXPORT void callable_call(pointer p_userdata, void* args, GDExtensionInt count, void* r_ret, GDExtensionCallError *r_error);

static inline void callable_custom_create(pointer fn, pointer r_ret, GDExtensionCallableCustomInfo *p_callable_custom_info) {
	p_callable_custom_info->call_func = (void*)callable_call;
	((GDExtensionInterfaceCallableCustomCreate)fn)((GDExtensionUninitializedTypePtr)r_ret, p_callable_custom_info);
}

extern EXPORT GDExtensionBool set_func(pointer p_instance, void* p_name, void* p_value);
extern EXPORT GDExtensionBool get_func(pointer p_instance, void* p_name, void* r_ret);
extern EXPORT GDExtensionPropertyInfo *get_property_list_func(pointer p_instance, uint32_t *r_count);
extern EXPORT void free_property_list_func(pointer p_instance, GDExtensionPropertyInfo *p_list);
extern EXPORT GDExtensionBool property_can_revert_func(pointer p_instance, void* p_name);
extern EXPORT GDExtensionBool property_get_revert_func(pointer p_instance, void* p_name, void* r_ret);
//extern *validate_property_func;
extern EXPORT void notification_func(pointer p_instance, int32_t p_notification, GDExtensionBool reversed);
extern EXPORT void to_string_func(pointer p_instance, void* valid, void* r_ret);
extern EXPORT void reference_func(pointer p_instance);
extern EXPORT void unreference_func(pointer p_instance);
extern EXPORT pointer create_instance_func(pointer p_class);
extern EXPORT void free_instance_func(pointer p_class, pointer p_instance);
extern EXPORT pointer get_virtual_call_data_func(pointer p_class, void* name);
extern EXPORT void call_virtual_with_data_func(pointer p_instance, void* name, pointer userdata, void* args, void* ret);
extern EXPORT uint64_t get_rid_func(pointer p_instance);

static inline void classdb_register_extension_class2(pointer fn, pointer p_library, pointer p_class_name, pointer p_parent_class_name, GDExtensionClassCreationInfo2 *p_extension_funcs) {
	p_extension_funcs->set_func = (void*)set_func;
	p_extension_funcs->get_func = (void*)get_func;
	p_extension_funcs->get_property_list_func = (void*)get_property_list_func;
	p_extension_funcs->free_property_list_func = (void*)free_property_list_func;
	p_extension_funcs->property_can_revert_func = (void*)property_can_revert_func;
	p_extension_funcs->property_get_revert_func = (void*)property_get_revert_func;
	//p_extension_funcs.validate_property_func = (void*)validate_property_func;
	p_extension_funcs->notification_func = (void*)notification_func;
	p_extension_funcs->to_string_func = (void*)to_string_func;
	p_extension_funcs->reference_func = (void*)reference_func;
	p_extension_funcs->unreference_func = (void*)unreference_func;
	p_extension_funcs->create_instance_func = (void*)create_instance_func;
	p_extension_funcs->free_instance_func = (void*)free_instance_func;
	p_extension_funcs->recreate_instance_func = 0;
	p_extension_funcs->get_virtual_call_data_func = (void*)get_virtual_call_data_func;
	p_extension_funcs->call_virtual_with_data_func = (void*)call_virtual_with_data_func;
	p_extension_funcs->get_rid_func = (void*)get_rid_func;
	((GDExtensionInterfaceClassdbRegisterExtensionClass2)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, (GDExtensionConstStringNamePtr)p_parent_class_name, p_extension_funcs);
}

extern EXPORT void method_call(pointer p_method, pointer p_userdata, void* args, GDExtensionInt count, void* r_ret, GDExtensionCallError *r_error);
extern EXPORT void method_ptrcall(pointer p_method, pointer p_userdata, void* args, void* r_ret);

static inline void classdb_register_extension_class_method(pointer fn, pointer p_library, pointer p_class_name, GDExtensionClassMethodInfo *p_method_bind_info) {
	p_method_bind_info->call_func = (void*)method_call;
	p_method_bind_info->ptrcall_func = (void*)method_ptrcall;
	((GDExtensionInterfaceClassdbRegisterExtensionClassMethod)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, p_method_bind_info);
}
static inline void classdb_register_extension_class_integer_constant(pointer fn, pointer p_library, pointer p_class_name, pointer p_enum_name, pointer p_constant_name, int64_t p_constant_value, GDExtensionBool p_is_bitfield) {
	((GDExtensionInterfaceClassdbRegisterExtensionClassIntegerConstant)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, (GDExtensionConstStringNamePtr)p_enum_name, (GDExtensionConstStringNamePtr)p_constant_name, p_constant_value, p_is_bitfield);
}
static inline void classdb_register_extension_class_property(pointer fn, pointer p_library, pointer p_class_name, GDExtensionPropertyInfo *p_property_info, pointer setter, pointer getter) {
	((GDExtensionInterfaceClassdbRegisterExtensionClassProperty)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, p_property_info, (GDExtensionConstStringNamePtr)setter, (GDExtensionConstStringNamePtr)getter);
}
static inline void classdb_register_extension_class_property_indexed(pointer fn, pointer p_library, pointer p_class_name, GDExtensionPropertyInfo *p_property_info, pointer setter, pointer getter, GDExtensionInt p_index) {
	((GDExtensionInterfaceClassdbRegisterExtensionClassPropertyIndexed)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, p_property_info, (GDExtensionConstStringNamePtr)setter, (GDExtensionConstStringNamePtr)getter, p_index);
}
static inline void classdb_register_extension_class_property_group(pointer fn, pointer p_library, pointer p_class_name, pointer p_group_name, pointer p_prefix) {
	((GDExtensionInterfaceClassdbRegisterExtensionClassPropertyGroup)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, (GDExtensionConstStringPtr)p_group_name, (GDExtensionConstStringPtr)p_prefix);
}
static inline void classdb_register_extension_class_property_subgroup(pointer fn, pointer p_library, pointer p_class_name, pointer p_group_name, pointer p_subgroup_name) {
	((GDExtensionInterfaceClassdbRegisterExtensionClassPropertySubgroup)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, (GDExtensionConstStringPtr)p_group_name, (GDExtensionConstStringPtr)p_subgroup_name);
}
static inline void classdb_register_extension_class_signal(pointer fn, pointer p_library, pointer p_class_name, pointer p_signal_name, GDExtensionPropertyInfo *args, GDExtensionInt arg_count) {
	((GDExtensionInterfaceClassdbRegisterExtensionClassSignal)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, (GDExtensionConstStringNamePtr)p_signal_name, args, arg_count);
}
static inline void editor_add_plugin(pointer fn, pointer p_name) {
	((GDExtensionInterfaceEditorAddPlugin)fn)((GDExtensionConstStringNamePtr)p_name);
}
static inline void editor_remove_plugin(pointer fn, pointer p_name) {
	((GDExtensionInterfaceEditorRemovePlugin)fn)((GDExtensionConstStringNamePtr)p_name);
}
static inline void classdb_unregister_extension_class(pointer fn, pointer p_library, pointer p_class_name) {
	((GDExtensionInterfaceClassdbUnregisterExtensionClass)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name);
}
static inline void editor_help_load_xml_from_utf8_chars_and_len(pointer fn, string p_xml, GDExtensionInt p_len) {
	((GDExtensionsInterfaceEditorHelpLoadXmlFromUtf8CharsAndLen)fn)((const char *)p_xml, p_len);
}
