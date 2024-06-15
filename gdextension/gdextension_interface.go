package gdextension

/*
#include <stdlib.h>
#include "gdextension_interface.h"

typedef uintptr_t pointer;
typedef const char* string;
typedef char32_t rune;
typedef uint8_t * bytes;
typedef GDExtensionBool bool;

// TODO the amount of code here can probably be reduced by stenciling out the various
// function signatures, such that we can call every function that takes 1 pointer arg,
// 2 pointer args etc.

extern void initialize(void *userdata, GDExtensionInitializationLevel p_level);
extern void deinitialize(void *userdata, GDExtensionInitializationLevel p_level);

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
static inline void *mem_realloc(pointer fn, void *ptr, size_t size) {
	return ((GDExtensionInterfaceMemRealloc)fn)(ptr, size);
}
static inline void mem_free(pointer fn, void *ptr) {
	((GDExtensionInterfaceMemFree)fn)(ptr);
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

extern void callable_call(pointer p_userdata, void* args, GDExtensionInt count, void* r_ret, GDExtensionCallError *r_error);

static inline void callable_custom_create(pointer fn, pointer r_ret, GDExtensionCallableCustomInfo *p_callable_custom_info) {
	p_callable_custom_info->call_func = (void*)callable_call;
	((GDExtensionInterfaceCallableCustomCreate)fn)((GDExtensionUninitializedTypePtr)r_ret, p_callable_custom_info);
}

extern bool set_func(pointer p_instance, void* p_name, void* p_value);
extern bool get_func(pointer p_instance, void* p_name, void* r_ret);
extern GDExtensionPropertyInfo *get_property_list_func(pointer p_instance, uint32_t *r_count);
extern void free_property_list_func(pointer p_instance, GDExtensionPropertyInfo *p_list);
extern bool property_can_revert_func(pointer p_instance, void* p_name);
extern bool property_get_revert_func(pointer p_instance, void* p_name, void* r_ret);
//extern *validate_property_func;
extern void notification_func(pointer p_instance, int32_t p_notification, bool reversed);
extern void to_string_func(pointer p_instance, void* valid, void* r_ret);
extern void reference_func(pointer p_instance);
extern void unreference_func(pointer p_instance);
extern pointer create_instance_func(pointer p_class);
extern void free_instance_func(pointer p_class, pointer p_instance);
extern uintptr_t recreate_instance_func(uintptr_t p_class, uintptr_t p_instance);
extern pointer get_virtual_call_data_func(pointer p_class, void* name);
extern void call_virtual_with_data_func(pointer p_instance, void* name, pointer userdata, void* args, void* ret);
extern uint64_t get_rid_func(pointer p_instance);

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
	p_extension_funcs->recreate_instance_func = (void*)recreate_instance_func;
	p_extension_funcs->get_virtual_call_data_func = (void*)get_virtual_call_data_func;
	p_extension_funcs->call_virtual_with_data_func = (void*)call_virtual_with_data_func;
	p_extension_funcs->get_rid_func = (void*)get_rid_func;
	((GDExtensionInterfaceClassdbRegisterExtensionClass2)fn)((GDExtensionClassLibraryPtr)p_library, (GDExtensionConstStringNamePtr)p_class_name, (GDExtensionConstStringNamePtr)p_parent_class_name, p_extension_funcs);
}

extern void method_call(pointer p_method, pointer p_userdata, void* args, GDExtensionInt count, void* r_ret, GDExtensionCallError *r_error);
extern void method_ptrcall(pointer p_method, pointer p_userdata, void* args, void* r_ret);

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

*/
import "C"

import (
	"errors"
	"runtime"
	"runtime/cgo"
	"unsafe"

	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/internal/callframe"

	"runtime.link/mmm"
)

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

type initialization = C.GDExtensionInitialization
type initializationLevel = C.GDExtensionInitializationLevel

func doInitialization(init *initialization) {
	C.initialization(init)
}

//export initialize
func initialize(_ unsafe.Pointer, level initializationLevel) {
	godot.Init(background, gd.GDExtensionInitializationLevel(level))
	if level == 2 {
		main()
	}
}

//export deinitialize
func deinitialize(_ unsafe.Pointer, level initializationLevel) {
	if level == 0 {
		background.End()
	}
}

func get_proc_address(handle uintptr, name string) unsafe.Pointer {
	name = name + "\000"
	char := C.CString(name)
	defer C.free(unsafe.Pointer(char))
	return C.get_proc_address(C.pointer(handle), char)
}

// linkCGO implements the Godot GDExtension API via CGO.
func linkCGO(API *gd.API) {
	get_godot_version := dlsymGD("get_godot_version")
	API.GetGodotVersion = func() gd.Version {
		var version = new(C.GDExtensionGodotVersion)
		C.get_godot_version(C.uintptr_t(uintptr(get_godot_version)), version)
		return gd.Version{
			Major: uint32(version.major),
			Minor: uint32(version.minor),
			Patch: uint32(version.patch),
			Value: C.GoString(version.string),
		}
	}
	mem_alloc := dlsymGD("mem_alloc")
	API.Memory.Allocate = func(size uintptr) unsafe.Pointer {
		return unsafe.Pointer(C.mem_alloc(C.uintptr_t(uintptr(mem_alloc)), C.size_t(size)))
	}
	mem_realloc := dlsymGD("mem_realloc")
	API.Memory.Reallocate = func(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
		return unsafe.Pointer(C.mem_realloc(C.uintptr_t(uintptr(mem_realloc)), ptr, C.size_t(size)))
	}
	mem_free := dlsymGD("mem_free")
	API.Memory.Free = func(ptr unsafe.Pointer) {
		C.mem_free(C.uintptr_t(uintptr(mem_free)), ptr)
	}
	print_error := dlsymGD("print_error")
	API.PrintError = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error(
			C.uintptr_t(uintptr(print_error)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_error_with_message := dlsymGD("print_error_with_message")
	API.PrintErrorMessage = func(code, message, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_message := C.CString(message)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error_with_message(
			C.uintptr_t(uintptr(print_error_with_message)),
			p_description,
			p_message,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_message))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_warning := dlsymGD("print_warning")
	API.PrintWarning = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_warning(
			C.uintptr_t(uintptr(print_warning)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_script_error := dlsymGD("print_script_error")
	API.PrintScriptError = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error(
			C.uintptr_t(uintptr(print_script_error)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_script_error_with_message := dlsymGD("print_script_error_with_message")
	API.PrintScriptErrorMessage = func(code, message, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_message := C.CString(message)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error_with_message(
			C.uintptr_t(uintptr(print_script_error_with_message)),
			p_description,
			p_message,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_message))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	get_native_struct_size := dlsymGD("get_native_struct_size")
	API.GetNativeStructSize = func(name gd.StringName) uintptr {
		var frame = callframe.New()
		defer frame.Free()
		return uintptr(C.get_native_struct_size(
			C.uintptr_t(uintptr(get_native_struct_size)),
			C.uintptr_t(callframe.Arg(frame, mmm.Get(name)).Uintptr()),
		))
	}
	variant_new_copy := dlsymGD("variant_new_copy")
	API.Variants.NewCopy = func(ctx gd.Context, self gd.Variant) gd.Variant {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var r_dest = callframe.Ret[[3]uintptr](frame)
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(r_dest.Uintptr()),
			C.uintptr_t(p_self.Uintptr()),
		)
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_dest.Get())
		frame.Free()
		return ret
	}
	variant_new_nil := dlsymGD("variant_new_nil")
	API.Variants.NewNil = func(ctx gd.Context) gd.Variant {
		var frame = callframe.New()
		var r_dest = callframe.Ret[[3]uintptr](frame)
		C.variant_new_nil(
			C.uintptr_t(uintptr(variant_new_nil)),
			C.uintptr_t(r_dest.Uintptr()),
		)
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_dest.Get())
		frame.Free()
		return ret
	}
	variant_destroy := dlsymGD("variant_destroy")
	API.Variants.Destroy = func(self gd.Variant) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		C.variant_destroy(
			C.uintptr_t(uintptr(variant_destroy)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
	}
	variant_call := dlsymGD("variant_call")
	API.Variants.Call = func(ctx gd.Context, self gd.Variant, method gd.StringName, args ...gd.Variant) (gd.Variant, error) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_method = callframe.Arg(frame, mmm.Get(method))
		for _, arg := range args {
			callframe.Arg(frame, mmm.Get(arg))
		}
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.variant_call(
			C.uintptr_t(uintptr(variant_call)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(frame.Array(2).Uintptr()),
			C.GDExtensionInt(len(args)),
			C.uintptr_t(r_ret.Uintptr()),
			r_error,
		)
		if r_error.error != 0 {
			frame.Free()
			return gd.Variant{}, &gd.CallError{
				ErrorType: gd.CallErrorType(r_error.error),
				Argument:  int32(r_error.argument),
				Expected:  int32(r_error.expected),
			}
		}
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_call_static := dlsymGD("variant_call_static")
	API.Variants.CallStatic = func(ctx gd.Context, vtype gd.VariantType, method gd.StringName, args ...gd.Variant) (gd.Variant, error) {
		var frame = callframe.New()
		var p_method = callframe.Arg(frame, mmm.Get(method))
		for _, arg := range args {
			callframe.Arg(frame, mmm.Get(arg))
		}
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.variant_call_static(
			C.uintptr_t(uintptr(variant_call_static)),
			C.GDExtensionVariantType(vtype),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(frame.Array(1).Uintptr()),
			C.GDExtensionInt(len(args)),
			C.uintptr_t(r_ret.Uintptr()),
			r_error,
		)
		if r_error.error != 0 {
			frame.Free()
			return gd.Variant{}, &gd.CallError{
				ErrorType: gd.CallErrorType(r_error.error),
				Argument:  int32(r_error.argument),
				Expected:  int32(r_error.expected),
			}
		}
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_evaluate := dlsymGD("variant_evaluate")
	API.Variants.Evaluate = func(ctx gd.Context, operator gd.Operator, a, b gd.Variant) (ret gd.Variant, ok bool) {
		var frame = callframe.New()
		var p_a = callframe.Arg(frame, mmm.Get(a))
		var p_b = callframe.Arg(frame, mmm.Get(b))
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_valid = callframe.Ret[bool](frame)
		C.variant_evaluate(
			C.uintptr_t(uintptr(variant_evaluate)),
			C.GDExtensionVariantOperator(operator),
			C.uintptr_t(p_a.Uintptr()),
			C.uintptr_t(p_b.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret, r_valid.Get()
	}
	variant_set := dlsymGD("variant_set")
	API.Variants.Set = func(self, key, val gd.Variant) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var p_value = callframe.Arg(frame, mmm.Get(val))
		var r_valid = callframe.Ret[bool](frame)
		C.variant_set(
			C.uintptr_t(uintptr(variant_set)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		frame.Free()
		return r_valid.Get()
	}
	variant_set_named := dlsymGD("variant_set_named")
	API.Variants.SetNamed = func(self gd.Variant, key gd.StringName, val gd.Variant) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var p_value = callframe.Arg(frame, mmm.Get(val))
		var r_valid = callframe.Ret[bool](frame)
		C.variant_set_named(
			C.uintptr_t(uintptr(variant_set_named)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		frame.Free()
		return r_valid.Get()
	}
	variant_set_keyed := dlsymGD("variant_set_keyed")
	API.Variants.SetKeyed = func(self, key, val gd.Variant) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var p_value = callframe.Arg(frame, mmm.Get(val))
		var r_valid = callframe.Ret[bool](frame)
		C.variant_set_keyed(
			C.uintptr_t(uintptr(variant_set_keyed)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		frame.Free()
		return r_valid.Get()
	}
	variant_set_indexed := dlsymGD("variant_set_indexed")
	API.Variants.SetIndexed = func(self gd.Variant, index gd.Int, val gd.Variant) (ok, oob bool) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_value = callframe.Arg(frame, mmm.Get(val))
		var r_valid = callframe.Ret[bool](frame)
		var r_oob = callframe.Ret[bool](frame)
		C.variant_set_indexed(
			C.uintptr_t(uintptr(variant_set_indexed)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
			C.uintptr_t(r_oob.Uintptr()),
		)
		frame.Free()
		return r_valid.Get(), r_oob.Get()
	}
	variant_get := dlsymGD("variant_get")
	API.Variants.Get = func(ctx gd.Context, self, key gd.Variant) (val gd.Variant, ok bool) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_valid = callframe.Ret[bool](frame)
		C.variant_get(
			C.uintptr_t(uintptr(variant_get)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_named := dlsymGD("variant_get_named")
	API.Variants.GetNamed = func(ctx gd.Context, self gd.Variant, key gd.StringName) (val gd.Variant, ok bool) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_valid = callframe.Ret[bool](frame)
		C.variant_get_named(
			C.uintptr_t(uintptr(variant_get_named)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_keyed := dlsymGD("variant_get_keyed")
	API.Variants.GetKeyed = func(ctx gd.Context, self, key gd.Variant) (val gd.Variant, ok bool) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_valid = callframe.Ret[bool](frame)
		C.variant_get_keyed(
			C.uintptr_t(uintptr(variant_get_keyed)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_indexed := dlsymGD("variant_get_indexed")
	API.Variants.GetIndexed = func(ctx gd.Context, self gd.Variant, index gd.Int) (gd.Variant, bool, bool) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_valid = callframe.Ret[bool](frame)
		var r_oob = callframe.Ret[bool](frame)
		C.variant_get_indexed(
			C.uintptr_t(uintptr(variant_get_indexed)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
			C.uintptr_t(r_oob.Uintptr()),
		)
		frame.Free()
		return mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get()), r_valid.Get(), r_oob.Get()
	}
	variant_iter_init := dlsymGD("variant_iter_init")
	API.Variants.IteratorInitialize = func(ctx gd.Context, self gd.Variant) (gd.Variant, bool) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var r_valid = callframe.Ret[bool](frame)
		var r_iter = callframe.Ret[[3]uintptr](frame)
		C.variant_iter_init(
			C.uintptr_t(uintptr(variant_iter_init)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(r_iter.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_iter.Get())
		var ok = r_valid.Get()
		frame.Free()
		return ret, ok
	}
	variant_iter_next := dlsymGD("variant_iter_next")
	API.Variants.IteratorNext = func(self gd.Variant, iter gd.Variant) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_iter = callframe.Arg(frame, mmm.Get(iter))
		var r_valid = callframe.Ret[bool](frame)
		C.variant_iter_next(
			C.uintptr_t(uintptr(variant_iter_next)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_iter.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		var ok = r_valid.Get()
		frame.Free()
		return ok
	}
	variant_iter_get := dlsymGD("variant_iter_get")
	API.Variants.IteratorGet = func(ctx gd.Context, self, iter gd.Variant) (gd.Variant, bool) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_iter = callframe.Arg(frame, mmm.Get(iter))
		var r_valid = callframe.Ret[bool](frame)
		var r_ret = callframe.Ret[[3]uintptr](frame)
		C.variant_iter_get(
			C.uintptr_t(uintptr(variant_iter_get)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_iter.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		var ok = r_valid.Get()
		frame.Free()
		return ret, ok
	}
	variant_hash := dlsymGD("variant_hash")
	API.Variants.Hash = func(self gd.Variant) gd.Int {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var ret = C.variant_hash(
			C.uintptr_t(uintptr(variant_hash)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		return gd.Int(ret)
	}
	variant_recursive_hash := dlsymGD("variant_recursive_hash")
	API.Variants.RecursiveHash = func(self gd.Variant, count gd.Int) gd.Int {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var ret = C.variant_recursive_hash(
			C.uintptr_t(uintptr(variant_recursive_hash)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(count),
		)
		frame.Free()
		return gd.Int(ret)
	}
	variant_hash_compare := dlsymGD("variant_hash_compare")
	API.Variants.HashCompare = func(self, other gd.Variant) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_other = callframe.Arg(frame, mmm.Get(other))
		var ret = C.variant_hash_compare(
			C.uintptr_t(uintptr(variant_hash_compare)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_other.Uintptr()),
		)
		frame.Free()
		return ret != 0
	}
	variant_booleanize := dlsymGD("variant_booleanize")
	API.Variants.Booleanize = func(self gd.Variant) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var ret = C.variant_booleanize(
			C.uintptr_t(uintptr(variant_booleanize)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		return ret != 0
	}
	variant_duplicate := dlsymGD("variant_duplicate")
	API.Variants.Duplicate = func(ctx gd.Context, self gd.Variant, deep bool) gd.Variant {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var p_deep = 0
		if deep {
			p_deep = 1
		}
		C.variant_duplicate(
			C.uintptr_t(uintptr(variant_duplicate)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.GDExtensionBool(p_deep),
		)
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_stringify := dlsymGD("variant_stringify")
	API.Variants.Stringify = func(ctx gd.Context, self gd.Variant) gd.String {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var r_ret = callframe.Ret[uintptr](frame)
		C.variant_stringify(
			C.uintptr_t(uintptr(variant_stringify)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_get_type := dlsymGD("variant_get_type")
	API.Variants.GetType = func(self gd.Variant) gd.VariantType {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var ret = C.variant_get_type(
			C.uintptr_t(uintptr(variant_get_type)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		return gd.VariantType(ret)
	}
	variant_has_method := dlsymGD("variant_has_method")
	API.Variants.HasMethod = func(self gd.Variant, method gd.StringName) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_method = callframe.Arg(frame, mmm.Get(method))
		var ret = C.variant_has_method(
			C.uintptr_t(uintptr(variant_has_method)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		return ret != 0
	}
	variant_has_member := dlsymGD("variant_has_member")
	API.Variants.HasMember = func(self gd.Variant, member gd.StringName) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_member = callframe.Arg(frame, mmm.Get(member))
		var ret = C.variant_has_member(
			C.uintptr_t(uintptr(variant_has_member)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_member.Uintptr()),
		)
		frame.Free()
		return ret != 0
	}
	variant_has_key := dlsymGD("variant_has_key")
	API.Variants.HasKey = func(self gd.Variant, key gd.Variant) (hasKey, valid bool) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var r_valid = callframe.Ret[bool](frame)
		var ret = C.variant_has_key(
			C.uintptr_t(uintptr(variant_has_key)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		valid = r_valid.Get()
		frame.Free()
		if ret != 0 {
			return true, valid
		}
		return false, valid
	}
	variant_get_type_name := dlsymGD("variant_get_type_name")
	API.Variants.GetTypeName = func(ctx gd.Context, vtype gd.VariantType) gd.String {
		var frame = callframe.New()
		var r_ret = callframe.Ret[uintptr](frame)
		C.variant_get_type_name(
			C.uintptr_t(uintptr(variant_get_type_name)),
			C.GDExtensionVariantType(vtype),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_can_convert := dlsymGD("variant_can_convert")
	API.Variants.CanConvert = func(self gd.Variant, toType gd.VariantType) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var ret = C.variant_can_convert(
			C.uintptr_t(uintptr(variant_can_convert)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionVariantType(toType),
		)
		frame.Free()
		return ret != 0
	}
	variant_can_convert_strict := dlsymGD("variant_can_convert_strict")
	API.Variants.CanConvertStrict = func(self gd.Variant, toType gd.VariantType) bool {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var ret = C.variant_can_convert_strict(
			C.uintptr_t(uintptr(variant_can_convert_strict)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionVariantType(toType),
		)
		frame.Free()
		return ret != 0
	}
	get_variant_from_type_constructor := dlsymGD("get_variant_from_type_constructor")
	API.Variants.FromTypeConstructor = func(vt gd.VariantType) func(ret callframe.Ptr[[3]uintptr], arg uintptr) {
		fn := C.get_variant_from_type_constructor(
			C.uintptr_t(uintptr(get_variant_from_type_constructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret callframe.Ptr[[3]uintptr], arg uintptr) {
			C.call_variant_from_type_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(arg),
			)
		}
	}
	get_variant_to_type_constructor := dlsymGD("get_variant_to_type_constructor")
	API.Variants.ToTypeConstructor = func(vt gd.VariantType) func(ret uintptr, arg callframe.Ptr[[3]uintptr]) {
		fn := C.get_variant_to_type_constructor(
			C.uintptr_t(uintptr(get_variant_to_type_constructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret uintptr, arg callframe.Ptr[[3]uintptr]) {
			C.call_variant_to_type_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_operator_evaluator := dlsymGD("variant_get_ptr_operator_evaluator")
	API.Variants.PointerOperatorEvaluator = func(op gd.Operator, a, b gd.VariantType) func(a, b, ret uintptr) {
		fn := C.variant_get_ptr_operator_evaluator(
			C.uintptr_t(uintptr(variant_get_ptr_operator_evaluator)),
			C.GDExtensionVariantOperator(op),
			C.GDExtensionVariantType(a),
			C.GDExtensionVariantType(b),
		)
		return func(a, b, ret uintptr) {
			C.call_variant_ptr_operator_evaluator(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(a),
				C.uintptr_t(b),
				C.uintptr_t(ret),
			)
		}
	}
	variant_get_ptr_builtin_method := dlsymGD("variant_get_ptr_builtin_method")
	API.Variants.GetPointerBuiltinMethod = func(vt gd.VariantType, sn gd.StringName, i gd.Int) func(base uintptr, args callframe.Args, ret uintptr, c int32) {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, mmm.Get(sn))
		fn := C.variant_get_ptr_builtin_method(
			C.uintptr_t(uintptr(variant_get_ptr_builtin_method)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
			C.GDExtensionInt(i),
		)
		frame.Free()
		return func(base uintptr, args callframe.Args, ret uintptr, c int32) {
			C.call_variant_ptr_builtin_method(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base),
				C.uintptr_t(args.Uintptr()),
				C.uintptr_t(ret),
				C.int32_t(c),
			)
		}
	}
	variant_get_ptr_constructor := dlsymGD("variant_get_ptr_constructor")
	API.Variants.GetPointerConstructor = func(vt gd.VariantType, index int32) func(ret uintptr, args callframe.Args) {
		fn := C.variant_get_ptr_constructor(
			C.uintptr_t(uintptr(variant_get_ptr_constructor)),
			C.GDExtensionVariantType(vt),
			C.int32_t(index),
		)
		return func(ret uintptr, args callframe.Args) {
			C.call_variant_ptr_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret),
				C.uintptr_t(args.Uintptr()),
			)
		}
	}
	variant_get_ptr_destructor := dlsymGD("variant_get_ptr_destructor")
	API.Variants.GetPointerDestructor = func(vt gd.VariantType) func(ret uintptr) {
		fn := C.variant_get_ptr_destructor(
			C.uintptr_t(uintptr(variant_get_ptr_destructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret uintptr) {
			C.call_variant_ptr_destructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret),
			)
		}
	}
	variant_construct := dlsymGD("variant_construct")
	API.Variants.Construct = func(ctx gd.Context, t gd.VariantType, args ...gd.Variant) (gd.Variant, error) {
		var frame = callframe.New()
		for _, arg := range args {
			callframe.Arg(frame, mmm.Get(arg))
		}
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.variant_construct(
			C.uintptr_t(uintptr(variant_construct)),
			C.GDExtensionVariantType(t),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(frame.Array(0).Uintptr()),
			C.int32_t(len(args)),
			r_error,
		)
		if r_error.error != 0 {
			frame.Free()
			return gd.Variant{}, &gd.CallError{
				ErrorType: gd.CallErrorType(r_error.error),
				Argument:  int32(r_error.argument),
				Expected:  int32(r_error.expected),
			}
		}
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_get_ptr_setter := dlsymGD("variant_get_ptr_setter")
	API.Variants.GetPointerSetter = func(vt gd.VariantType, sn gd.StringName) func(base, arg uintptr) {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, mmm.Get(sn))
		fn := C.variant_get_ptr_setter(
			C.uintptr_t(uintptr(variant_get_ptr_setter)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		return func(base, arg uintptr) {
			C.call_variant_ptr_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base),
				C.uintptr_t(arg),
			)
		}
	}
	variant_get_ptr_getter := dlsymGD("variant_get_ptr_getter")
	API.Variants.GetPointerGetter = func(vt gd.VariantType, sn gd.StringName) func(base uintptr, ret uintptr) {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, mmm.Get(sn))
		fn := C.variant_get_ptr_getter(
			C.uintptr_t(uintptr(variant_get_ptr_getter)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		return func(base uintptr, ret uintptr) {
			C.call_variant_ptr_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base),
				C.uintptr_t(ret),
			)
		}
	}
	variant_get_ptr_indexed_setter := dlsymGD("variant_get_ptr_indexed_setter")
	API.Variants.GetPointerIndexedSetter = func(vt gd.VariantType) func(base uintptr, index gd.Int, arg uintptr) {
		fn := C.variant_get_ptr_indexed_setter(
			C.uintptr_t(uintptr(variant_get_ptr_indexed_setter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base uintptr, index gd.Int, arg uintptr) {
			C.call_variant_ptr_indexed_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base),
				C.GDExtensionInt(index),
				C.uintptr_t(arg),
			)
		}
	}
	variant_get_ptr_indexed_getter := dlsymGD("variant_get_ptr_indexed_getter")
	API.Variants.GetPointerIndexedGetter = func(vt gd.VariantType) func(base uintptr, index gd.Int, ret uintptr) {
		fn := C.variant_get_ptr_indexed_getter(
			C.uintptr_t(uintptr(variant_get_ptr_indexed_getter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base uintptr, index gd.Int, ret uintptr) {
			C.call_variant_ptr_indexed_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base),
				C.GDExtensionInt(index),
				C.uintptr_t(ret),
			)
		}
	}
	variant_get_ptr_keyed_setter := dlsymGD("variant_get_ptr_keyed_setter")
	API.Variants.GetPointerKeyedSetter = func(vt gd.VariantType) func(base uintptr, key uintptr, arg uintptr) {
		fn := C.variant_get_ptr_keyed_setter(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_setter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base uintptr, key uintptr, arg uintptr) {
			C.call_variant_ptr_keyed_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base),
				C.uintptr_t(key),
				C.uintptr_t(arg),
			)
		}
	}
	variant_get_ptr_keyed_getter := dlsymGD("variant_get_ptr_keyed_getter")
	API.Variants.GetPointerKeyedGetter = func(vt gd.VariantType) func(base uintptr, key uintptr, ret uintptr) {
		fn := C.variant_get_ptr_keyed_getter(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_getter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base uintptr, key uintptr, ret uintptr) {
			C.call_variant_ptr_keyed_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base),
				C.uintptr_t(key),
				C.uintptr_t(ret),
			)
		}
	}
	variant_get_ptr_keyed_checker := dlsymGD("variant_get_ptr_keyed_checker")
	API.Variants.GetPointerKeyedChecker = func(vt gd.VariantType) func(base uintptr, key uintptr) uint32 {
		fn := C.variant_get_ptr_keyed_checker(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_checker)),
			C.GDExtensionVariantType(vt),
		)
		return func(base uintptr, key uintptr) uint32 {
			return uint32(C.call_variant_ptr_keyed_checker(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base),
				C.uintptr_t(key),
			))
		}
	}
	variant_get_constant_value := dlsymGD("variant_get_constant_value")
	API.Variants.GetConstantValue = func(ctx gd.Context, vt gd.VariantType, sn gd.StringName) gd.Variant {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, mmm.Get(sn))
		var r_ret = callframe.Ret[[3]uintptr](frame)
		C.variant_get_constant_value(
			C.uintptr_t(uintptr(variant_get_constant_value)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_get_ptr_utility_function := dlsymGD("variant_get_ptr_utility_function")
	API.Variants.GetPointerUtilityFunction = func(sn gd.StringName, hash gd.Int) func(ret uintptr, args callframe.Args, c int32) {
		var frame = callframe.New()
		p_method := callframe.Arg(frame, mmm.Get(sn))
		fn := C.variant_get_ptr_utility_function(
			C.uintptr_t(uintptr(variant_get_ptr_utility_function)),
			C.uintptr_t(p_method.Uintptr()),
			C.GDExtensionInt(hash),
		)
		frame.Free()
		return func(ret uintptr, args callframe.Args, c int32) {
			C.call_variant_ptr_utility_function(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret),
				C.uintptr_t(args.Uintptr()),
				C.int32_t(c),
			)
		}
	}
	string_new_with_utf8_chars_and_len := dlsymGD("string_new_with_utf8_chars_and_len")
	API.Strings.New = func(ctx gd.Context, s string) gd.String {
		var str = C.CString(s)
		var frame = callframe.New()
		var r_ret = callframe.Ret[uintptr](frame)
		C.string_new_with_utf8_chars_and_len(
			C.uintptr_t(uintptr(string_new_with_utf8_chars_and_len)),
			C.uintptr_t(r_ret.Uintptr()),
			str,
			C.GDExtensionInt(len(s)),
		)
		var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		C.free(unsafe.Pointer(str))
		return ret
	}
	string_to_utf8_chars := dlsymGD("string_to_utf8_chars")
	API.Strings.Get = func(s gd.String) string {
		var length = s.Length()
		if length == 0 {
			return ""
		}
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(s))
		var r_ret = make([]byte, length)
		C.string_to_utf8_chars(
			C.uintptr_t(uintptr(string_to_utf8_chars)),
			C.uintptr_t(p_self.Uintptr()),
			(*C.char)(unsafe.Pointer(&r_ret[0])),
			C.GDExtensionInt(length),
		)
		frame.Free()
		return string(r_ret)
	}
	string_operator_index := dlsymGD("string_operator_index")
	API.Strings.SetIndex = func(s gd.String, index gd.Int, val rune) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(s))
		*C.string_operator_index(
			C.uintptr_t(uintptr(string_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
		) = C.rune(val)
		frame.Free()
	}
	string_operator_index_const := dlsymGD("string_operator_index")
	API.Strings.Index = func(s gd.String, index gd.Int) rune {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(s))
		var ret = C.string_operator_index_const(
			C.uintptr_t(uintptr(string_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
		)
		frame.Free()
		return rune(*ret)
	}
	string_operator_plus_eq_string := dlsymGD("string_operator_plus_eq_string")
	API.Strings.Append = func(ctx gd.Context, s gd.String, other gd.String) gd.String {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(s))
		var p_other = callframe.Arg(frame, mmm.Get(other))
		C.string_operator_plus_eq_string(
			C.uintptr_t(uintptr(string_operator_plus_eq_string)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_other.Uintptr()),
		)
		var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, p_self.Get())
		frame.Free()
		return ret
	}
	string_operator_plus_eq_char := dlsymGD("string_operator_plus_eq_char")
	API.Strings.AppendRune = func(s gd.String, other rune) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(s))
		C.string_operator_plus_eq_char(
			C.uintptr_t(uintptr(string_operator_plus_eq_char)),
			C.uintptr_t(p_self.Uintptr()),
			C.rune(other),
		)
		frame.Free()
	}
	string_resize := dlsymGD("string_resize")
	API.Strings.Resize = func(s gd.String, size gd.Int) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(s))
		var length = s.Length()
		C.string_resize(
			C.uintptr_t(uintptr(string_resize)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(size),
		)
		if size < length {
			API.Strings.SetIndex(s, size, 0)
		}
		frame.Free()
	}
	string_name_new_with_utf8_chars_and_len := dlsymGD("string_name_new_with_utf8_chars_and_len")
	API.StringNames.New = func(ctx gd.Context, s string) gd.StringName {
		var str = C.CString(s)
		var frame = callframe.New()
		var r_ret = callframe.Ret[uintptr](frame)
		C.string_name_new_with_utf8_chars_and_len(
			C.uintptr_t(uintptr(string_name_new_with_utf8_chars_and_len)),
			C.uintptr_t(r_ret.Uintptr()),
			str,
			C.GDExtensionInt(len(s)),
		)
		var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		C.free(unsafe.Pointer(str))
		return ret
	}
	xml_parser_open_buffer := dlsymGD("xml_parser_open_buffer")
	API.XMLParser.OpenBuffer = func(x gd.Object, b []byte) error {
		var pin runtime.Pinner
		pin.Pin(&b[0])
		mmm.Pin[pinner](mmm.Life(x.AsPointer()), &pin, [0]uintptr{})

		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(x.AsPointer()))
		var errCode = C.xml_parser_open_buffer(
			C.uintptr_t(uintptr(xml_parser_open_buffer)),
			C.uintptr_t(p_self.Uintptr()),
			(C.bytes)(unsafe.Pointer(&b[0])),
			C.size_t(len(b)),
		)
		frame.Free()
		if errCode != 0 {
			return errors.New("xml_parser_open_buffer failed") // TODO godot error code string?
		}
		return nil
	}
	file_access_store_buffer := dlsymGD("file_access_store_buffer")
	API.FileAccess.StoreBuffer = func(f gd.Object, b []byte) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(f.AsPointer()))
		C.file_access_store_buffer(
			C.uintptr_t(uintptr(file_access_store_buffer)),
			C.uintptr_t(p_self.Uintptr()),
			(C.bytes)(unsafe.Pointer(&b[0])),
			C.uint64_t(len(b)),
		)
		frame.Free()
	}
	file_access_get_buffer := dlsymGD("file_access_get_buffer")
	API.FileAccess.GetBuffer = func(f gd.Object, b []byte) int {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(f.AsPointer()))
		var length = C.file_access_get_buffer(
			C.uintptr_t(uintptr(file_access_get_buffer)),
			C.uintptr_t(p_self.Uintptr()),
			(C.bytes)(unsafe.Pointer(&b[0])),
			C.uint64_t(len(b)),
		)
		frame.Free()
		return int(length)
	}
	API.PackedByteArray = makePackedFunctions[*gd.PackedByteArray, byte]("byte_array")
	API.PackedColorArray = makePackedFunctions[*gd.PackedColorArray, gd.Color]("color_array")
	API.PackedFloat32Array = makePackedFunctions[*gd.PackedFloat32Array, float32]("float32_array")
	API.PackedFloat64Array = makePackedFunctions[*gd.PackedFloat64Array, float64]("float64_array")
	API.PackedInt32Array = makePackedFunctions[*gd.PackedInt32Array, int32]("int32_array")
	API.PackedInt64Array = makePackedFunctions[*gd.PackedInt64Array, int64]("int64_array")
	packed_string_array_operator_index_const := dlsymGD("packed_string_array_operator_index_const")
	API.PackedStringArray.Index = func(ctx gd.Context, psa gd.PackedStringArray, i gd.Int) gd.String {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(psa))
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_string_array_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		frame.Free()
		return mmm.Let[gd.String](ctx.Lifetime, ctx.API, uintptr(*(*uintptr)(ret)))
	}
	packed_string_array_operator_index := dlsymGD("packed_string_array_operator_index")
	API.PackedStringArray.SetIndex = func(psa gd.PackedStringArray, i gd.Int, v gd.String) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(psa))
		var ptr = C.packed_T_operator_index(
			C.uintptr_t(uintptr(packed_string_array_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		*(*uintptr)(ptr) = mmm.Get(v)
		frame.Free()
	}
	API.PackedVector2Array = makePackedFunctions[*gd.PackedVector2Array, gd.Vector2]("vector2_array")
	API.PackedVector3Array = makePackedFunctions[*gd.PackedVector3Array, gd.Vector3]("vector3_array")
	array_operator_index_const := dlsymGD("array_operator_index_const")
	API.Array.Index = func(ctx gd.Context, a gd.Array, i gd.Int) gd.Variant {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(a))
		var r_ret = C.array_operator_index_const(
			C.uintptr_t(uintptr(array_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		var ptr = (*[3]uintptr)(r_ret)
		var ret = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, *ptr)
		frame.Free()
		return ret.Copy(ctx)
	}
	array_operator_index := dlsymGD("array_operator_index")
	API.Array.SetIndex = func(a gd.Array, i gd.Int, v gd.Variant) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(a))
		var ptr = C.array_operator_index(
			C.uintptr_t(uintptr(array_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		var p_copy = callframe.Ret[[3]uintptr](frame)
		var p_value = callframe.Arg(frame, mmm.Get(v))
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(p_copy.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
		)
		*(*[3]uintptr)(ptr) = p_copy.Get()
		frame.Free()
	}
	array_ref := dlsymGD("array_ref")
	API.Array.Set = func(self gd.Array, from gd.Array) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_from = callframe.Arg(frame, mmm.Get(from))
		C.array_ref(
			C.uintptr_t(uintptr(array_ref)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_from.Uintptr()),
		)
		frame.Free()
	}
	array_set_typed := dlsymGD("array_set_typed")
	API.Array.SetTyped = func(self gd.Array, t gd.VariantType, className gd.StringName, script gd.Object) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(self))
		var p_className = callframe.Arg(frame, mmm.Get(className))
		var p_script = callframe.Arg(frame, mmm.Get(script.AsPointer()))
		C.array_set_typed(
			C.uintptr_t(uintptr(array_set_typed)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionVariantType(t),
			C.uintptr_t(p_className.Uintptr()),
			C.uintptr_t(p_script.Uintptr()), // FIXME should this be a variant?
		)
		frame.Free()
	}
	dictionary_operator_index := dlsymGD("dictionary_operator_index")
	API.Dictionary.Index = func(ctx gd.Context, d gd.Dictionary, key gd.Variant) gd.Variant {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(d))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var ptr = C.dictionary_operator_index(
			C.uintptr_t(uintptr(dictionary_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
		)
		var r_ret = *(*[3]uintptr)(ptr)
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret)
		frame.Free()
		return ret
	}
	API.Dictionary.SetIndex = func(dict gd.Dictionary, key, val gd.Variant) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(dict))
		var p_key = callframe.Arg(frame, mmm.Get(key))
		var ptr = C.dictionary_operator_index(
			C.uintptr_t(uintptr(dictionary_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
		)
		var p_copy = callframe.Ret[[3]uintptr](frame)
		var p_value = callframe.Arg(frame, mmm.Get(val))
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(p_copy.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
		)
		*(*[3]uintptr)(ptr) = p_copy.Get()
		frame.Free()
	}
	object_get_instance_from_id := dlsymGD("object_get_instance_from_id")
	API.Object.GetInstanceFromID = func(ctx gd.Context, id gd.ObjectID) gd.Object {
		var ret = C.object_get_instance_from_id(
			C.uintptr_t(uintptr(object_get_instance_from_id)),
			C.uintptr_t(id),
		)
		var obj gd.Object
		obj.SetPointer(gd.PointerMustAssertInstanceID(ctx, uintptr(ret)))
		return obj
	}
	object_method_bind_call := dlsymGD("object_method_bind_call")
	API.Object.MethodBindCall = func(ctx gd.Context, method gd.MethodBind, obj gd.Object, arg ...gd.Variant) (gd.Variant, error) {
		var self = mmm.Get(obj.AsPointer())
		if self[0] == 0 {
			return gd.Variant{}, errors.New("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				return gd.Variant{}, errors.New("use after free")
			}
		}
		var frame = callframe.New()
		for _, a := range arg {
			callframe.Arg(frame, mmm.Get(a))
		}
		var r_ret = callframe.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.object_method_bind_call(
			C.uintptr_t(uintptr(object_method_bind_call)),
			C.uintptr_t(method),
			C.uintptr_t(self[0]),
			C.uintptr_t(frame.Array(0).Uintptr()),
			C.GDExtensionInt(len(arg)),
			C.uintptr_t(r_ret.Uintptr()),
			r_error,
		)
		if r_error.error != 0 {
			frame.Free()
			return gd.Variant{}, &gd.CallError{
				ErrorType: gd.CallErrorType(r_error.error),
				Argument:  int32(r_error.argument),
				Expected:  int32(r_error.expected),
			}
		}
		var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret, nil
	}
	object_method_bind_ptrcall := dlsymGD("object_method_bind_ptrcall")
	API.Object.MethodBindPointerCall = func(method gd.MethodBind, obj gd.Object, arg callframe.Args, ret uintptr) {
		var self = mmm.Get(obj.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		C.object_method_bind_ptrcall(
			C.uintptr_t(uintptr(object_method_bind_ptrcall)),
			C.uintptr_t(method),
			C.uintptr_t(self[0]),
			C.uintptr_t(arg.Uintptr()),
			C.uintptr_t(ret),
		)
	}
	object_destroy := dlsymGD("object_destroy")
	API.Object.Destroy = func(o gd.Object) {
		var self = mmm.Get(o.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		C.object_destroy(
			C.uintptr_t(uintptr(object_destroy)),
			C.uintptr_t(self[0]),
		)
	}
	global_get_singleton := dlsymGD("global_get_singleton")
	API.Object.GetSingleton = func(ctx gd.Context, name gd.StringName) gd.Object {
		var frame = callframe.New()
		var p_name = callframe.Arg(frame, mmm.Get(name))
		var ret = C.global_get_singleton(
			C.uintptr_t(uintptr(global_get_singleton)),
			C.uintptr_t(p_name.Uintptr()),
		)
		var obj gd.Object
		obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{uintptr(ret)}))
		frame.Free()
		return obj
	}
	object_get_instance_binding := dlsymGD("object_get_instance_binding")
	API.Object.GetInstanceBinding = func(o gd.Object, et gd.ExtensionToken, ibt gd.InstanceBindingType) any {
		var self = mmm.Get(o.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var ret = C.object_get_instance_binding(
			C.uintptr_t(uintptr(object_get_instance_binding)),
			C.uintptr_t(self[0]),
			C.uintptr_t(et),
			unsafe.Pointer(ibt),
		)
		return cgo.Handle(ret).Value()
	}
	object_set_instance_binding := dlsymGD("object_set_instance_binding")
	API.Object.SetInstanceBinding = func(o gd.Object, et gd.ExtensionToken, val any, ibt gd.InstanceBindingType) {
		var self = mmm.Get(o.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}

		var pinner runtime.Pinner
		var bindings C.GDExtensionInstanceBindingCallbacks
		pinner.Pin(&bindings)
		defer pinner.Unpin()

		p_val := cgo.NewHandle(val)
		C.object_set_instance_binding(
			C.uintptr_t(uintptr(object_set_instance_binding)),
			C.uintptr_t(self[0]),
			C.uintptr_t(et),
			C.uintptr_t(p_val),
			unsafe.Pointer(&bindings),
		)
	}
	object_free_instance_binding := dlsymGD("object_free_instance_binding")
	API.Object.FreeInstanceBinding = func(o gd.Object, et gd.ExtensionToken) {
		var self = mmm.Get(o.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		C.object_free_instance_binding(
			C.uintptr_t(uintptr(object_free_instance_binding)),
			C.uintptr_t(self[0]),
			C.uintptr_t(et),
		)
	}
	object_set_instance := dlsymGD("object_set_instance")
	API.Object.SetInstance = func(o gd.Object, sn gd.StringName, a gd.ObjectInterface) {
		var self = mmm.Get(o.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var frame = callframe.New()
		var p_sn = callframe.Arg(frame, mmm.Get(sn))
		var p_val = cgo.NewHandle(a)
		C.object_set_instance(
			C.uintptr_t(uintptr(object_set_instance)),
			C.uintptr_t(self[0]),
			C.uintptr_t(p_sn.Uintptr()),
			C.uintptr_t(p_val),
		)
		frame.Free()
	}
	object_get_class_name := dlsymGD("object_get_class_name")
	API.Object.GetClassName = func(ctx gd.Context, o gd.Object, token gd.ExtensionToken) gd.String {
		var self = mmm.Get(o.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var frame = callframe.New()
		var r_ret = callframe.Ret[uintptr](frame)
		C.object_get_class_name(
			C.uintptr_t(uintptr(object_get_class_name)),
			C.uintptr_t(self[0]),
			C.uintptr_t(token),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret
	}
	object_cast_to := dlsymGD("object_cast_to")
	API.Object.CastTo = func(o gd.Object, ct gd.ClassTag) gd.Object {
		var self = mmm.Get(o.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var ret = C.object_cast_to(
			C.uintptr_t(uintptr(object_cast_to)),
			C.uintptr_t(self[0]),
			C.uintptr_t(ct),
		)
		if ret == 0 {
			return gd.Object{}
		}
		var obj gd.Object
		obj.SetPointer(o.AsPointer())
		return obj
	}
	object_get_instance_id := dlsymGD("object_get_instance_id")
	API.Object.GetInstanceID = func(o gd.Object) gd.ObjectID {
		var self = mmm.Get(o.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		return gd.ObjectID(C.object_get_instance_id(
			C.uintptr_t(uintptr(object_get_instance_id)),
			C.uintptr_t(self[0]),
		))
	}
	ref_get_object := dlsymGD("ref_get_object")
	API.RefCounted.GetObject = func(ctx gd.Context, rc gd.Object) gd.Object {
		var self = mmm.Get(rc.AsPointer())
		if self[0] == 0 {
			panic("nil gd.Object dereference")
		}
		if self[1] != 0 {
			var ret = C.object_get_instance_from_id(
				C.uintptr_t(uintptr(object_get_instance_from_id)),
				C.uintptr_t(self[1]),
			)
			if ret == 0 {
				panic("use after free")
			}
		}
		var ret = C.ref_get_object(
			C.uintptr_t(uintptr(ref_get_object)),
			C.uintptr_t(self[0]),
		)
		var obj gd.Object
		obj.SetPointer(mmm.New[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{uintptr(ret)}))
		return obj
	}
	ref_set_object := dlsymGD("ref_set_object")
	API.RefCounted.SetObject = func(rc gd.Object, o gd.Object) {
		C.ref_set_object(
			C.uintptr_t(uintptr(ref_set_object)),
			C.uintptr_t(mmm.Get(rc.AsPointer())[0]),
			C.uintptr_t(mmm.Get(o.AsPointer())[0]),
		)
	}
	classdb_construct_object := dlsymGD("classdb_construct_object")
	API.ClassDB.ConstructObject = func(ctx gd.Context, name gd.StringName) gd.Object {
		var frame = callframe.New()
		var obj gd.Object
		obj.SetPointer(mmm.New[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{uintptr(C.classdb_construct_object(
			C.uintptr_t(uintptr(classdb_construct_object)),
			C.uintptr_t(callframe.Arg(frame, mmm.Get(name)).Uintptr()),
		))}))
		frame.Free()
		return obj
	}
	classdb_get_class_tag := dlsymGD("classdb_get_class_tag")
	API.ClassDB.GetClassTag = func(name gd.StringName) gd.ClassTag {
		var frame = callframe.New()
		defer frame.Free()
		return gd.ClassTag(C.classdb_get_class_tag(
			C.uintptr_t(uintptr(classdb_get_class_tag)),
			C.uintptr_t(callframe.Arg(frame, mmm.Get(name)).Uintptr()),
		))
	}
	classdb_get_method_bind := dlsymGD("classdb_get_method_bind")
	API.ClassDB.GetMethodBind = func(class, method gd.StringName, hash gd.Int) gd.MethodBind {
		var frame = callframe.New()
		defer frame.Free()
		return gd.MethodBind(C.classdb_get_method_bind(
			C.uintptr_t(uintptr(classdb_get_method_bind)),
			C.uintptr_t(callframe.Arg(frame, mmm.Get(class)).Uintptr()),
			C.uintptr_t(callframe.Arg(frame, mmm.Get(method)).Uintptr()),
			C.GDExtensionInt(hash),
		))
	}
	get_library_path := dlsymGD("get_library_path")
	API.GetLibraryPath = func(ctx gd.Context, et gd.ExtensionToken) gd.String {
		var frame = callframe.New()
		var r_ret = callframe.Ret[uintptr](frame)
		C.get_library_path(
			C.uintptr_t(uintptr(get_library_path)),
			C.uintptr_t(et),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
		frame.Free()
		return ret
	}

	callable_custom_create := dlsymGD("callable_custom_create")
	API.Callables.Create = func(ctx gd.Context, fn func(...gd.Variant) (gd.Variant, error)) gd.Callable {
		var frame = callframe.New()
		var r_callable = callframe.Ret[[2]uintptr](frame)
		var info C.GDExtensionCallableCustomInfo
		info.token = C.uintptr_t(ctx.API.ExtensionToken)
		info.callable_userdata = C.uintptr_t(cgo.NewHandle(fn))
		C.callable_custom_create(
			C.uintptr_t(uintptr(callable_custom_create)),
			C.uintptr_t(r_callable.Uintptr()),
			&info,
		)
		var r_ret = mmm.New[gd.Callable](ctx.Lifetime, ctx.API, r_callable.Get())
		frame.Free()
		return r_ret
	}

	classdb_register_extension_class_signal := dlsymGD("classdb_register_extension_class_signal")
	API.ClassDB.RegisterClassSignal = func(library gd.ExtensionToken, class, signal gd.StringName, args []gd.PropertyInfo) {
		ctx := gd.NewContext(&godot)
		defer ctx.End()
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, mmm.Get(class))
		var p_signal = callframe.Arg(frame, mmm.Get(signal))
		var p_list = cPropertyList(ctx, args)
		C.classdb_register_extension_class_signal(
			C.uintptr_t(uintptr(classdb_register_extension_class_signal)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			C.uintptr_t(p_signal.Uintptr()),
			p_list,
			C.GDExtensionInt(len(args)),
		)
		frame.Free()
	}

	classdb_register_extension_class2 := dlsymGD("classdb_register_extension_class2")
	API.ClassDB.RegisterClass = func(library gd.ExtensionToken, name, extends gd.StringName, info gd.ClassInterface) {
		var frame = callframe.New()
		var p_name = callframe.Arg(frame, mmm.Get(name))
		var p_extends = callframe.Arg(frame, mmm.Get(extends))
		var is_virtual C.GDExtensionBool
		if info.IsVirtual() {
			is_virtual = 1
		}
		var is_abstract C.GDExtensionBool
		if info.IsAbstract() {
			is_abstract = 1
		}
		var is_exposed C.GDExtensionBool
		if info.IsExposed() {
			is_exposed = 1
		}
		var p_info = C.GDExtensionClassCreationInfo2{
			is_virtual:     is_virtual,
			is_abstract:    is_abstract,
			is_exposed:     is_exposed,
			class_userdata: C.uintptr_t(cgo.NewHandle(info)),
		}
		C.classdb_register_extension_class2(
			C.uintptr_t(uintptr(classdb_register_extension_class2)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_name.Uintptr()),
			C.uintptr_t(p_extends.Uintptr()),
			&p_info,
		)
		frame.Free()
	}

	classdb_register_extension_class_integer_constant := dlsymGD("classdb_register_extension_class_integer_constant")
	API.ClassDB.RegisterClassIntegerConstant = func(library gd.ExtensionToken, class, enum, constant gd.StringName, value int64, bitfield bool) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, mmm.Get(class))
		var p_enum = callframe.Arg(frame, mmm.Get(enum))
		var p_constant = callframe.Arg(frame, mmm.Get(constant))
		C.classdb_register_extension_class_integer_constant(
			C.uintptr_t(uintptr(classdb_register_extension_class_integer_constant)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			C.uintptr_t(p_enum.Uintptr()),
			C.uintptr_t(p_constant.Uintptr()),
			C.GDExtensionInt(value),
			C.GDExtensionBool(btoi(bitfield)),
		)
		frame.Free()
	}
	classdb_register_extension_class_property := dlsymGD("classdb_register_extension_class_property")
	API.ClassDB.RegisterClassProperty = func(library gd.ExtensionToken, class gd.StringName, info gd.PropertyInfo, getter, setter gd.StringName) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, mmm.Get(class))
		var p_getter = callframe.Arg(frame, mmm.Get(getter))
		var p_setter = callframe.Arg(frame, mmm.Get(setter))
		var pins runtime.Pinner
		defer pins.Unpin()
		var info_name = mmm.Get(info.Name)
		var info_className = mmm.Get(info.ClassName)
		var info_hintString = mmm.Get(info.HintString)
		pins.Pin(&info_name)
		pins.Pin(&info_className)
		pins.Pin(&info_hintString)
		var p_info = C.GDExtensionPropertyInfo{
			_type:       C.GDExtensionVariantType(info.Type),
			name:        (C.GDExtensionStringNamePtr)(unsafe.Pointer(&info_name)),
			class_name:  (C.GDExtensionStringNamePtr)(unsafe.Pointer(&info_className)),
			hint:        C.uint32_t(info.Hint),
			hint_string: (C.GDExtensionStringPtr)(unsafe.Pointer(&info_hintString)),
			usage:       C.uint32_t(info.Usage),
		}
		C.classdb_register_extension_class_property(
			C.uintptr_t(uintptr(classdb_register_extension_class_property)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			&p_info,
			C.uintptr_t(p_getter.Uintptr()),
			C.uintptr_t(p_setter.Uintptr()),
		)
		frame.Free()
	}
	classdb_register_extension_class_property_indexed := dlsymGD("classdb_register_extension_class_property_indexed")
	API.ClassDB.RegisterClassPropertyIndexed = func(library gd.ExtensionToken, class gd.StringName, info gd.PropertyInfo, getter, setter gd.StringName, index int64) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, mmm.Get(class))
		var p_getter = callframe.Arg(frame, mmm.Get(getter))
		var p_setter = callframe.Arg(frame, mmm.Get(setter))
		var pins runtime.Pinner
		defer pins.Unpin()
		var info_name = mmm.Get(info.Name)
		var info_className = mmm.Get(info.ClassName)
		var info_hintString = mmm.Get(info.HintString)
		pins.Pin(&info_name)
		pins.Pin(&info_className)
		pins.Pin(&info_hintString)
		var p_info = C.GDExtensionPropertyInfo{
			_type:       C.GDExtensionVariantType(info.Type),
			name:        (C.GDExtensionStringNamePtr)(unsafe.Pointer(&info_name)),
			class_name:  (C.GDExtensionStringNamePtr)(unsafe.Pointer(&info_className)),
			hint:        C.uint32_t(info.Hint),
			hint_string: (C.GDExtensionStringPtr)(unsafe.Pointer(&info_hintString)),
			usage:       C.uint32_t(info.Usage),
		}
		C.classdb_register_extension_class_property_indexed(
			C.uintptr_t(uintptr(classdb_register_extension_class_property_indexed)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			&p_info,
			C.uintptr_t(p_getter.Uintptr()),
			C.uintptr_t(p_setter.Uintptr()),
			C.GDExtensionInt(index),
		)
		frame.Free()
	}

	classdb_register_extension_class_property_group := dlsymGD("classdb_register_extension_class_property_group")
	API.ClassDB.RegisterClassPropertyGroup = func(library gd.ExtensionToken, class gd.StringName, group, prefix gd.String) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, mmm.Get(class))
		var p_group = callframe.Arg(frame, mmm.Get(group))
		var p_prefix = callframe.Arg(frame, mmm.Get(prefix))
		C.classdb_register_extension_class_property_group(
			C.uintptr_t(uintptr(classdb_register_extension_class_property_group)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			C.uintptr_t(p_group.Uintptr()),
			C.uintptr_t(p_prefix.Uintptr()),
		)
		frame.Free()
	}
	classdb_register_extension_class_property_subgroup := dlsymGD("classdb_register_extension_class_property_subgroup")
	API.ClassDB.RegisterClassPropertySubGroup = func(library gd.ExtensionToken, class gd.StringName, subGroup, prefix gd.String) {
		var frame = callframe.New()
		var p_class = callframe.Arg(frame, mmm.Get(class))
		var p_subGroup = callframe.Arg(frame, mmm.Get(subGroup))
		var p_prefix = callframe.Arg(frame, mmm.Get(prefix))
		C.classdb_register_extension_class_property_subgroup(
			C.uintptr_t(uintptr(classdb_register_extension_class_property_subgroup)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			C.uintptr_t(p_subGroup.Uintptr()),
			C.uintptr_t(p_prefix.Uintptr()),
		)
		frame.Free()
	}
	classdb_register_extension_class_method := dlsymGD("classdb_register_extension_class_method")
	API.ClassDB.RegisterClassMethod = func(ctx gd.Context, library gd.ExtensionToken, class gd.StringName, info gd.Method) {
		infoHandle := cgo.NewHandle(&info)
		releaseHandle := infoHandle.Delete
		mmm.Pin[onFree](ctx.Lifetime, &releaseHandle, [0]uintptr{})

		var pins runtime.Pinner
		mmm.Pin[pinner](ctx.Lifetime, &pins, [0]uintptr{})

		var name = mmm.Get(info.Name)
		pins.Pin(&name)

		var returnInfo *C.GDExtensionPropertyInfo

		var has_return_value C.GDExtensionBool
		if info.ReturnValueInfo != nil {
			has_return_value = 1

			var retName = mmm.Get(info.ReturnValueInfo.Name)
			pins.Pin(&retName)

			var className = mmm.Get(info.ReturnValueInfo.ClassName)
			pins.Pin(&className)

			var hintString = mmm.Get(info.ReturnValueInfo.HintString)
			pins.Pin(&hintString)

			returnInfo = &C.GDExtensionPropertyInfo{
				_type:       C.GDExtensionVariantType(info.ReturnValueInfo.Type),
				name:        (C.GDExtensionStringNamePtr)(unsafe.Pointer(&retName)),
				class_name:  (C.GDExtensionStringNamePtr)(unsafe.Pointer(&className)),
				hint:        C.uint32_t(info.ReturnValueInfo.Hint),
				hint_string: (C.GDExtensionStringPtr)(unsafe.Pointer(&hintString)),
				usage:       C.uint32_t(info.ReturnValueInfo.Usage),
			}

			pins.Pin(returnInfo)
		}

		var list = cPropertyList(ctx, info.Arguments)

		var firstMetadata *C.GDExtensionClassMethodArgumentMetadata
		var metadatas = make([]C.GDExtensionClassMethodArgumentMetadata, 0, len(info.ArgumentsMetadata))
		for _, metadata := range info.ArgumentsMetadata {
			metadatas = append(metadatas, C.GDExtensionClassMethodArgumentMetadata(metadata))
		}
		if len(metadatas) > 0 {
			firstMetadata = &metadatas[0]
			pins.Pin(&metadatas[0])
		}

		var firstDefaultArgument *C.GDExtensionVariantPtr
		var defaultArguments = make([]C.GDExtensionVariantPtr, 0, len(info.DefaultArguments))
		for _, arg := range info.DefaultArguments {
			var def = mmm.Get(arg)
			pins.Pin(&def)
			defaultArguments = append(defaultArguments, C.GDExtensionVariantPtr(unsafe.Pointer(&def)))
		}
		if len(defaultArguments) > 0 {
			firstDefaultArgument = &defaultArguments[0]
			pins.Pin(&defaultArguments[0])
		}

		var frame = callframe.New()
		var p_class = callframe.Arg(frame, mmm.Get(class))
		var p_info = C.GDExtensionClassMethodInfo{
			name:                   (C.GDExtensionStringNamePtr)(unsafe.Pointer(&name)),
			method_userdata:        C.uintptr_t(infoHandle), // FIXME leak
			method_flags:           C.uint32_t(info.MethodFlags),
			has_return_value:       has_return_value,
			return_value_info:      returnInfo,
			return_value_metadata:  C.GDExtensionClassMethodArgumentMetadata(info.ReturnValueMetadata),
			argument_count:         C.uint32_t(len(info.Arguments)),
			arguments_info:         list,
			arguments_metadata:     firstMetadata,
			default_argument_count: C.uint32_t(len(info.DefaultArguments)),
			default_arguments:      firstDefaultArgument,
		}
		C.classdb_register_extension_class_method(
			C.uintptr_t(uintptr(classdb_register_extension_class_method)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_class.Uintptr()),
			&p_info,
		)
		frame.Free()
	}

	editor_add_plugin := dlsymGD("editor_add_plugin")
	API.EditorPlugins.Add = func(plugin gd.StringName) {
		var frame = callframe.New()
		var p_plugin = callframe.Arg(frame, mmm.Get(plugin))
		C.editor_add_plugin(
			C.uintptr_t(uintptr(editor_add_plugin)),
			C.uintptr_t(p_plugin.Uintptr()),
		)
		frame.Free()
	}
	editor_remove_plugin := dlsymGD("editor_remove_plugin")
	API.EditorPlugins.Remove = func(plugin gd.StringName) {
		var frame = callframe.New()
		var p_plugin = callframe.Arg(frame, mmm.Get(plugin))
		C.editor_remove_plugin(
			C.uintptr_t(uintptr(editor_remove_plugin)),
			C.uintptr_t(p_plugin.Uintptr()),
		)
		frame.Free()
	}
	classdb_unregister_extension_class := dlsymGD("classdb_unregister_extension_class")
	API.ClassDB.UnregisterClass = func(library gd.ExtensionToken, name gd.StringName) {
		var frame = callframe.New()
		var p_name = callframe.Arg(frame, mmm.Get(name))
		C.classdb_unregister_extension_class(
			C.uintptr_t(uintptr(classdb_unregister_extension_class)),
			C.uintptr_t(uintptr(library)),
			C.uintptr_t(p_name.Uintptr()),
		)
		frame.Free()
	}
}

func makePackedFunctions[T gd.Packed, V comparable](prefix string) gd.PackedFunctionsFor[T, V] {
	var API gd.PackedFunctionsFor[T, V]
	packed_T_operator_index := dlsymGD("packed_" + prefix + "_operator_index")
	API.SetIndex = func(t T, i gd.Int, v V) {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(t))
		var ptr = C.packed_T_operator_index(
			C.uintptr_t(uintptr(packed_T_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		*(*V)(unsafe.Pointer(uintptr(ptr))) = v
		frame.Free()
	}
	packed_T_operator_index_const := dlsymGD("packed_" + prefix + "_operator_index_const")
	API.Index = func(t T, i gd.Int) V {
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(t))
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_T_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		frame.Free()
		return *(*V)(unsafe.Pointer(uintptr(ret)))
	}
	API.CopyAsSlice = func(t T) []V {
		var size = t.Size()
		if size == 0 {
			return nil
		}
		var frame = callframe.New()
		var p_self = callframe.Arg(frame, mmm.Get(t))
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_T_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(0),
		)
		var slice []V
		copy(slice, unsafe.Slice((*V)(unsafe.Pointer(uintptr(ret))), size))
		frame.Free()
		return slice
	}
	return API
}

//export set_func
func set_func(p_instance uintptr, p_name, p_value unsafe.Pointer) bool {
	ctx := gd.NewContext(&godot)
	defer ctx.End()

	name := mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *(*uintptr)(p_name))
	value := mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, *(*[3]uintptr)(p_value))
	return cgo.Handle(p_instance).Value().(gd.ObjectInterface).Set(name, value)
}

//export get_func
func get_func(p_instance uintptr, p_name, p_value unsafe.Pointer) bool {
	ctx := gd.NewContext(&godot)
	defer ctx.End()

	name := mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *(*uintptr)(p_name))
	variant, ok := cgo.Handle(p_instance).Value().(gd.ObjectInterface).Get(name)
	if !ok {
		return false
	}
	*(*[3]uintptr)(p_value) = mmm.End(variant)
	return true
}

var propertyLists = make(map[*C.GDExtensionPropertyInfo]gd.Context)

type onFree mmm.Pointer[func(), onFree, [0]uintptr]

func (cb onFree) Free() {
	(*mmm.API(cb))()
	mmm.End(cb)
}

func cPropertyList(ctx gd.Context, list []gd.PropertyInfo) *C.GDExtensionPropertyInfo {
	if len(list) == 0 {
		return nil
	}
	var pins runtime.Pinner
	var slice = make([]C.GDExtensionPropertyInfo, 0, len(list))
	for i := range list {
		property := &list[i]
		name := mmm.Get(property.Name)
		pins.Pin(&name)
		class_name := mmm.Get(property.ClassName)
		pins.Pin(&class_name)
		hint_string := mmm.Get(property.HintString)
		pins.Pin(&hint_string)
		slice = append(slice, C.GDExtensionPropertyInfo{
			_type:       C.GDExtensionVariantType(property.Type),
			name:        (C.GDExtensionStringNamePtr)(unsafe.Pointer(&name)),
			class_name:  (C.GDExtensionStringNamePtr)(unsafe.Pointer(&class_name)),
			hint:        C.uint32_t(property.Hint),
			hint_string: (C.GDExtensionStringPtr)(unsafe.Pointer(&hint_string)),
			usage:       C.uint32_t(property.Usage),
		})
	}
	pins.Pin(&slice[0])
	mmm.Pin[pinner](ctx.Lifetime, &pins, [0]uintptr{})
	propertyLists[&slice[0]] = ctx
	del := func() {
		delete(propertyLists, &slice[0])
	}
	mmm.Pin[onFree](ctx.Lifetime, &del, [0]uintptr{})
	return &slice[0]
}

//export get_property_list_func
func get_property_list_func(p_instance uintptr, p_length *uint32) *C.GDExtensionPropertyInfo {
	ctx := gd.NewContext(&godot)
	list := cgo.Handle(p_instance).Value().(gd.ObjectInterface).GetPropertyList(ctx)
	*p_length = uint32(len(list))
	return cPropertyList(ctx, list)
}

//export free_property_list_func
func free_property_list_func(_ uintptr, p_properties *C.GDExtensionPropertyInfo) {
	propertyLists[p_properties].End()
}

//export property_can_revert_func
func property_can_revert_func(p_instance uintptr, p_name unsafe.Pointer) bool {
	ctx := gd.NewContext(&godot)
	defer ctx.End()
	name := mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *(*uintptr)(p_name))
	return cgo.Handle(p_instance).Value().(gd.ObjectInterface).PropertyCanRevert(name)
}

//export property_get_revert_func
func property_get_revert_func(p_instance uintptr, p_name, p_value unsafe.Pointer) bool {
	ctx := gd.NewContext(&godot)
	defer ctx.End()
	name := mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *(*uintptr)(p_name))
	variant, ok := cgo.Handle(p_instance).Value().(gd.ObjectInterface).PropertyGetRevert(ctx, name)
	if ok {
		*(*[3]uintptr)(p_value) = mmm.End(variant)
	}
	return ok
}

//export notification_func
func notification_func(p_instance uintptr, p_notification int32, p_reversed bool) {
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).Notification(p_notification, p_reversed)
}

//export to_string_func
func to_string_func(p_instance uintptr, valid, out unsafe.Pointer) {
	ctx := gd.NewContext(&godot)
	defer ctx.End()
	s, ok := cgo.Handle(p_instance).Value().(gd.ObjectInterface).ToString()
	if !ok {
		*(*bool)(valid) = false
		return
	}
	*(*bool)(valid) = true
	*(*uintptr)(out) = mmm.Get(s)
}

//export reference_func
func reference_func(p_instance uintptr) {
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).Reference()
}

//export unreference_func
func unreference_func(p_instance uintptr) {
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).Unreference()
}

//export create_instance_func
func create_instance_func(p_class uintptr) uintptr {
	return mmm.Get(cgo.Handle(p_class).Value().(gd.ClassInterface).CreateInstance().AsPointer())[0]
}

//export free_instance_func
func free_instance_func(_, p_instance uintptr) {
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).Free()
}

//export recreate_instance_func
func recreate_instance_func(p_class, p_super uintptr) uintptr {
	ctx := gd.NewContext(&godot)
	var super gd.Object
	super.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{p_super}))
	return uintptr(cgo.NewHandle(cgo.Handle(p_class).Value().(gd.ClassInterface).ReloadInstance(super)))
}

//export get_virtual_call_data_func
func get_virtual_call_data_func(p_class uintptr, p_name unsafe.Pointer) uintptr {
	ctx := gd.NewContext(&godot)
	defer ctx.End()
	var name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *(*uintptr)(p_name))
	virtual := cgo.Handle(p_class).Value().(gd.ClassInterface).GetVirtual(name)
	if virtual == nil {
		return 0
	}
	return uintptr(cgo.NewHandle(virtual))
}

//export call_virtual_with_data_func
func call_virtual_with_data_func(p_instance uintptr, p_name unsafe.Pointer, p_data uintptr, p_args, p_ret unsafe.Pointer) {
	ctx := gd.NewContext(&godot)
	defer ctx.End()
	var name = mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *(*uintptr)(p_name))
	cgo.Handle(p_instance).Value().(gd.ObjectInterface).CallVirtual(name, cgo.Handle(p_data).Value(), gd.UnsafeArgs(p_args), gd.UnsafeBack(p_ret))
}

//export get_rid_func
func get_rid_func(p_instance uintptr) C.uint64_t {
	return C.uint64_t(cgo.Handle(p_instance).Value().(gd.ObjectInterface).GetRID())
}

//export callable_call
func callable_call(p_callable uintptr, p_args unsafe.Pointer, count C.GDExtensionInt, p_ret unsafe.Pointer, issue *C.GDExtensionCallError) {
	fn := cgo.Handle(p_callable).Value().(func(...gd.Variant) (gd.Variant, error))

	var slice = unsafe.Slice((*[3]uintptr)(p_args), int(count))

	ctx := gd.NewContext(&godot)
	defer ctx.End()

	var args = make([]gd.Variant, 0, len(slice))
	for _, elem := range slice {
		args = append(args, mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, elem))
	}
	ret, err := fn(args...)
	if err != nil {
		issue.error = 7 // TODO no generic error>
		return
	}
	*(*[3]uintptr)(p_ret) = mmm.Get(ret)
	*issue = C.GDExtensionCallError{}
}

//export method_call
func method_call(p_method uintptr, p_instance uintptr, p_args unsafe.Pointer, count C.GDExtensionInt, p_ret unsafe.Pointer, issue *C.GDExtensionCallError) {
	method := cgo.Handle(p_method).Value().(*gd.Method)
	var variants = make([]gd.Variant, 0, int(count))
	ctx := gd.NewContext(&godot)
	defer ctx.End()
	for _, elem := range unsafe.Slice((**[3]uintptr)(p_args), int(count)) {
		variants = append(variants, mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, *elem))
	}
	result, err := method.Call(ctx, cgo.Handle(p_instance).Value(), variants...)
	if err != nil {
		issue.error = 7 // TODO no generic error>
		return
	}
	*(*[3]uintptr)(p_ret) = mmm.End(result)
}

//export method_ptrcall
func method_ptrcall(p_method uintptr, p_instance uintptr, p_args unsafe.Pointer, p_ret unsafe.Pointer) {
	method := cgo.Handle(p_method).Value().(*gd.Method)
	method.PointerCall(cgo.Handle(p_instance).Value(), gd.UnsafeArgs(p_args), gd.UnsafeBack(p_ret))
}
