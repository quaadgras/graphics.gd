package gdextension

/*
#include <stdlib.h>
#include "gdextension_interface.h"

typedef uintptr_t pointer;
typedef const char* string;
typedef char32_t rune;
typedef uint8_t * bytes;

void get_godot_version(pointer fn, GDExtensionGodotVersion *r_version) {
	((GDExtensionInterfaceGetGodotVersion)fn)(r_version);
}
void *mem_alloc(pointer fn, size_t size) {
	return ((GDExtensionInterfaceMemAlloc)fn)(size);
}
void *mem_realloc(pointer fn, void *ptr, size_t size) {
	return ((GDExtensionInterfaceMemRealloc)fn)(ptr, size);
}
void mem_free(pointer fn, void *ptr) {
	((GDExtensionInterfaceMemFree)fn)(ptr);
}
void print_error(pointer fn, string p_description, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintError)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}
void print_error_with_message(pointer fn, string p_description, string p_message, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintErrorWithMessage)fn)(p_description, p_message, p_function, p_file, p_line, p_notify_editor);
}
void print_warning(pointer fn, string p_description, string p_function, string p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintWarning)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}
uint64_t get_native_struct_size(pointer fn, pointer p_classname) {
	return ((GDExtensionInterfaceGetNativeStructSize)fn)((GDExtensionConstStringNamePtr)p_classname);
}
void variant_new_copy(pointer fn, pointer r_dest, pointer p_self) {
	((GDExtensionInterfaceVariantNewCopy)fn)((GDExtensionUninitializedVariantPtr)r_dest, (GDExtensionConstVariantPtr)p_self);
}
void variant_new_nil(pointer fn, pointer r_dest) {
	((GDExtensionInterfaceVariantNewNil)fn)((GDExtensionUninitializedVariantPtr)r_dest);
}
void variant_destroy(pointer fn, pointer p_self) {
	((GDExtensionInterfaceVariantDestroy)fn)((GDExtensionVariantPtr)p_self);
}
void variant_call(pointer fn, pointer p_self, pointer p_method, pointer p_args, GDExtensionInt p_argument_count, pointer r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantCall)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_method, (GDExtensionConstVariantPtr)p_args, (GDExtensionInt)p_argument_count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
void variant_call_static(pointer fn, GDExtensionVariantType p_type, pointer p_method, pointer p_args, GDExtensionInt p_argument_count, pointer r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantCallStatic)fn)((GDExtensionVariantType)p_type, (GDExtensionConstStringNamePtr)p_method, (GDExtensionConstVariantPtr)p_args, (GDExtensionInt)p_argument_count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
void variant_evaluate(pointer fn, GDExtensionVariantOperator p_operator, pointer p_a, pointer p_b, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantEvaluate)fn)(p_operator, (GDExtensionConstVariantPtr)p_a, (GDExtensionConstVariantPtr)p_b, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_set(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSet)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
void variant_set_named(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSetNamed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
void variant_set_keyed(pointer fn, pointer p_self, pointer p_key, pointer p_value, pointer r_valid) {
	((GDExtensionInterfaceVariantSetKeyed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}
void variant_set_indexed(pointer fn, pointer p_self, GDExtensionInt p_index, pointer p_value, pointer r_valid, pointer r_oob) {
	((GDExtensionInterfaceVariantSetIndexed)fn)((GDExtensionVariantPtr)p_self, (GDExtensionInt)p_index, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid, (GDExtensionBool*)r_oob);
}
void variant_get(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGet)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_get_named(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGetNamed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_get_keyed(pointer fn, pointer p_self, pointer p_key, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantGetKeyed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_get_indexed(pointer fn, pointer p_self, GDExtensionInt p_index, pointer r_ret, pointer r_valid, pointer r_oob) {
	((GDExtensionInterfaceVariantGetIndexed)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionInt)p_index, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid, (GDExtensionBool*)r_oob);
}
void variant_iter_init(pointer fn, pointer p_self, pointer r_iter, pointer r_valid) {
	((GDExtensionInterfaceVariantIterInit)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedVariantPtr)r_iter, (GDExtensionBool*)r_valid);
}
void variant_iter_next(pointer fn, pointer p_self, pointer p_iter, pointer r_valid) {
	((GDExtensionInterfaceVariantIterNext)fn)((GDExtensionVariantPtr)p_self, (GDExtensionVariantPtr)p_iter, (GDExtensionBool*)r_valid);
}
void variant_iter_get(pointer fn, pointer p_self, pointer p_iter, pointer r_ret, pointer r_valid) {
	((GDExtensionInterfaceVariantIterGet)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionVariantPtr)p_iter, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
GDExtensionInt variant_hash(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantHash)fn)((GDExtensionConstVariantPtr)p_self);
}
GDExtensionInt variant_recursive_hash(pointer fn, pointer p_self, GDExtensionInt p_count) {
	return ((GDExtensionInterfaceVariantRecursiveHash)fn)((GDExtensionConstVariantPtr)p_self, p_count);
}
GDExtensionBool variant_hash_compare(pointer fn, pointer p_self, pointer p_other) {
	return ((GDExtensionInterfaceVariantHashCompare)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_other);
}
GDExtensionBool variant_booleanize(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantBooleanize)fn)((GDExtensionConstVariantPtr)p_self);
}
void variant_duplicate(pointer fn, pointer p_self, pointer r_ret, GDExtensionBool p_deep) {
	((GDExtensionInterfaceVariantDuplicate)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedVariantPtr)r_ret, p_deep);
}
void variant_stringify(pointer fn, pointer p_self, pointer r_ret) {
	((GDExtensionInterfaceVariantStringify)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionUninitializedStringPtr)r_ret);
}
GDExtensionVariantType variant_get_type(pointer fn, pointer p_self) {
	return ((GDExtensionInterfaceVariantGetType)fn)((GDExtensionConstVariantPtr)p_self);
}
GDExtensionBool variant_has_method(pointer fn, pointer p_self, pointer p_method) {
	return ((GDExtensionInterfaceVariantHasMethod)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_method);
}
GDExtensionBool variant_has_member(pointer fn, pointer p_self, pointer p_member) {
	return ((GDExtensionInterfaceVariantHasMember)fn)((GDExtensionVariantType)p_self, (GDExtensionConstStringNamePtr)p_member);
}
GDExtensionBool variant_has_key(pointer fn, pointer p_self, pointer p_key, pointer r_valid) {
	return ((GDExtensionInterfaceVariantHasKey)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionBool*)r_valid);
}
void variant_get_type_name(pointer fn, GDExtensionVariantType p_type, pointer r_ret) {
	((GDExtensionInterfaceVariantGetTypeName)fn)(p_type, (GDExtensionUninitializedStringPtr)r_ret);
}
GDExtensionBool variant_can_convert(pointer fn, pointer p_self, GDExtensionVariantType p_type) {
	return ((GDExtensionInterfaceVariantCanConvert)fn)((GDExtensionVariantType)p_self, p_type);
}
GDExtensionBool variant_can_convert_strict(pointer fn, pointer p_self, GDExtensionVariantType p_type) {
	return ((GDExtensionInterfaceVariantCanConvertStrict)fn)((GDExtensionVariantType)p_self, p_type);
}
pointer get_variant_from_type_constructor(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceGetVariantFromTypeConstructor)fn)(p_type);
}
void call_variant_from_type_constructor(pointer fn, pointer r_ret, pointer r_arg) {
	((GDExtensionVariantFromTypeConstructorFunc)fn)((GDExtensionUninitializedVariantPtr)r_ret, (void*)r_arg);
}
pointer get_variant_to_type_constructor(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceGetVariantToTypeConstructor)fn)(p_type);
}
void call_variant_to_type_constructor(pointer fn, pointer r_ret, const pointer r_arg) {
	((GDExtensionTypeFromVariantConstructorFunc)fn)((GDExtensionUninitializedVariantPtr)r_ret, (void*)r_arg);
}
pointer variant_get_ptr_operator_evaluator(pointer fn, GDExtensionVariantOperator p_operator, GDExtensionVariantType p_type, GDExtensionVariantType p_rhs_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrOperatorEvaluator)fn)(p_operator, p_type, p_rhs_type);
}
void call_variant_ptr_operator_evaluator(pointer fn, pointer r_ret, pointer r_a, pointer r_b) {
	((GDExtensionPtrOperatorEvaluator)fn)((GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionConstVariantPtr)r_a, (void*)r_b);
}
pointer variant_get_ptr_builtin_method(pointer fn, GDExtensionVariantType p_type, pointer p_method, GDExtensionInt p_hash) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrBuiltinMethod)fn)(p_type, (GDExtensionConstStringNamePtr)p_method, p_hash);
}
void call_variant_ptr_builtin_method(pointer fn, pointer p_base, pointer r_arg, pointer r_ret, int count) {
	((GDExtensionPtrBuiltInMethod)fn)((GDExtensionTypePtr)p_base, (const GDExtensionConstTypePtr *)r_arg, (GDExtensionTypePtr)r_ret, count);
}
pointer variant_get_ptr_constructor(pointer fn, GDExtensionVariantType p_type, int32_t p_constructor) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrConstructor)fn)(p_type, p_constructor);
}
void call_variant_ptr_constructor(pointer fn, pointer r_ret, pointer r_arg) {
	((GDExtensionPtrConstructor)fn)((GDExtensionUninitializedVariantPtr)r_ret, (void*)r_arg);
}
pointer variant_get_ptr_destructor(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrDestructor)fn)(p_type);
}
void call_variant_ptr_destructor(pointer fn, pointer r_arg) {
	((GDExtensionPtrDestructor)fn)((GDExtensionVariantPtr)r_arg);
}
void variant_construct(pointer fn, GDExtensionVariantType p_type, pointer p_base, pointer p_args, int32_t p_argument_count, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantConstruct)fn)(p_type, (GDExtensionUninitializedVariantPtr)p_base, (GDExtensionConstVariantPtr)p_args, p_argument_count, r_error);
}
pointer variant_get_ptr_setter(pointer fn, GDExtensionVariantType p_type, pointer p_member) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrSetter)fn)(p_type, (GDExtensionConstStringNamePtr)p_member);
}
void call_variant_ptr_setter(pointer fn, pointer p_base, pointer p_value) {
	((GDExtensionPtrSetter)fn)((GDExtensionTypePtr)p_base, (GDExtensionConstTypePtr)p_value);
}
pointer variant_get_ptr_getter(pointer fn, GDExtensionVariantType p_type, pointer p_member) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrGetter)fn)(p_type, (GDExtensionConstStringNamePtr)p_member);
}
void call_variant_ptr_getter(pointer fn, pointer p_base, pointer r_value) {
	((GDExtensionPtrGetter)fn)((GDExtensionTypePtr)p_base, (GDExtensionTypePtr)r_value);
}
pointer variant_get_ptr_indexed_setter(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrIndexedSetter)fn)(p_type);
}
void call_variant_ptr_indexed_setter(pointer fn, pointer p_base, GDExtensionInt p_index, pointer r_value) {
	((GDExtensionPtrIndexedSetter)fn)((GDExtensionTypePtr)p_base, p_index, (GDExtensionConstTypePtr)r_value);
}
pointer variant_get_ptr_indexed_getter(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrIndexedGetter)fn)(p_type);
}
void call_variant_ptr_indexed_getter(pointer fn, pointer p_base, GDExtensionInt p_index, pointer r_ret) {
	((GDExtensionPtrIndexedGetter)fn)((GDExtensionTypePtr)p_base, p_index, (GDExtensionTypePtr)r_ret);
}
pointer variant_get_ptr_keyed_setter(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrKeyedSetter)fn)(p_type);
}
void call_variant_ptr_keyed_setter(pointer fn, pointer p_base, pointer r_arg, pointer r_value) {
	((GDExtensionPtrKeyedSetter)fn)((GDExtensionTypePtr)p_base, (const GDExtensionConstTypePtr *)r_arg, (GDExtensionConstTypePtr)r_value);
}
pointer variant_get_ptr_keyed_getter(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrKeyedGetter)fn)(p_type);
}
void call_variant_ptr_keyed_getter(pointer fn, pointer p_base, pointer r_arg, pointer r_ret) {
	((GDExtensionPtrKeyedGetter)fn)((GDExtensionTypePtr)p_base, (const GDExtensionConstTypePtr *)r_arg, (GDExtensionTypePtr)r_ret);
}
pointer variant_get_ptr_keyed_checker(pointer fn, GDExtensionVariantType p_type) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrKeyedChecker)fn)(p_type);
}
uint32_t call_variant_ptr_keyed_checker(pointer fn, pointer p_base, pointer r_arg) {
	return ((GDExtensionPtrKeyedChecker)fn)((GDExtensionTypePtr)p_base, (const GDExtensionConstTypePtr *)r_arg);
}
void variant_get_constant_value(pointer fn, GDExtensionVariantType p_type, pointer p_constant, pointer r_ret) {
	((GDExtensionInterfaceVariantGetConstantValue)fn)(p_type, (GDExtensionConstStringNamePtr)p_constant, (GDExtensionUninitializedVariantPtr)r_ret);
}
pointer variant_get_ptr_utility_function(pointer fn, pointer p_name, GDExtensionInt hash) {
	return (pointer)((GDExtensionInterfaceVariantGetPtrUtilityFunction)fn)((GDExtensionConstStringNamePtr)p_name, hash);
}
void call_variant_ptr_utility_function(pointer fn, pointer r_ret, pointer r_arg, int count) {
	((GDExtensionPtrUtilityFunction)fn)((GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionConstVariantPtr)r_arg, count);
}
void string_new_with_utf8_chars_and_len(pointer fn, pointer r_ret, string p_chars, GDExtensionInt p_len) {
	((GDExtensionInterfaceStringNewWithUtf8CharsAndLen)fn)((GDExtensionUninitializedStringPtr)r_ret, p_chars, p_len);
}
GDExtensionInt string_to_utf8_chars(pointer fn, pointer p_self, char *r_chars, GDExtensionInt p_max) {
	return ((GDExtensionInterfaceStringToUtf8Chars)fn)((GDExtensionConstStringPtr)p_self, r_chars, p_max);
}
rune *string_operator_index(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return ((GDExtensionInterfaceStringOperatorIndex)fn)((GDExtensionStringPtr)p_self, p_index);
}
const rune *string_operator_index_const(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return ((GDExtensionInterfaceStringOperatorIndexConst)fn)((GDExtensionConstStringPtr)p_self, p_index);
}
void string_operator_plus_eq_string(pointer fn, pointer p_self, pointer p_other) {
	((GDExtensionInterfaceStringOperatorPlusEqString)fn)((GDExtensionStringPtr)p_self, (GDExtensionConstStringPtr)p_other);
}
void string_operator_plus_eq_char(pointer fn, pointer p_self, rune p_other) {
	((GDExtensionInterfaceStringOperatorPlusEqChar)fn)((GDExtensionStringPtr)p_self, p_other);
}
void string_resize(pointer fn, pointer p_self, GDExtensionInt p_size) {
	((GDExtensionInterfaceStringResize)fn)((GDExtensionStringPtr)p_self, p_size);
}
void string_name_new_with_utf8_chars_and_len(pointer fn, pointer r_ret, string p_chars, GDExtensionInt p_len) {
	((GDExtensionInterfaceStringNameNewWithUtf8CharsAndLen)fn)((GDExtensionUninitializedStringNamePtr)r_ret, p_chars, p_len);
}
GDExtensionInt xml_parser_open_buffer(pointer fn, pointer p_instance, const bytes p_buffer, size_t p_size) {
	return ((GDExtensionInterfaceXmlParserOpenBuffer)fn)((GDExtensionObjectPtr)p_instance, (const uint8_t *)p_buffer, p_size);
}
void file_access_store_buffer(pointer fn, pointer p_self, const bytes p_buffer, uint64_t p_len) {
	((GDExtensionInterfaceFileAccessStoreBuffer)fn)((GDExtensionObjectPtr)p_self, (const uint8_t *)p_buffer, p_len);
}
uint64_t file_access_get_buffer(pointer fn, pointer p_self, bytes r_buffer, uint64_t p_length) {
	return ((GDExtensionInterfaceFileAccessGetBuffer)fn)((GDExtensionConstObjectPtr)p_self, (uint8_t *)r_buffer, p_length);
}
pointer packed_T_operator_index(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return (pointer)((GDExtensionInterfacePackedByteArrayOperatorIndex)fn)((GDExtensionTypePtr)p_self, p_index);
}
const pointer packed_T_operator_index_const(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return (pointer)((GDExtensionInterfacePackedByteArrayOperatorIndexConst)fn)((GDExtensionTypePtr)p_self, p_index);
}
pointer array_operator_index(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return (pointer)((GDExtensionInterfaceArrayOperatorIndex)fn)((GDExtensionTypePtr)p_self, p_index);
}
const pointer array_operator_index_const(pointer fn, pointer p_self, GDExtensionInt p_index) {
	return (const pointer)((GDExtensionInterfaceArrayOperatorIndexConst)fn)((GDExtensionTypePtr)p_self, p_index);
}
void array_ref(pointer fn, pointer p_self, pointer p_from) {
	((GDExtensionInterfaceArrayRef)fn)((GDExtensionTypePtr)p_self, (GDExtensionConstTypePtr)p_from);
}
void array_set_typed(pointer fn, pointer p_self, GDExtensionVariantType p_type, pointer p_class_name, pointer p_script) {
	((GDExtensionInterfaceArraySetTyped)fn)((GDExtensionTypePtr)p_self, p_type, (GDExtensionConstStringNamePtr)p_class_name, (GDExtensionConstVariantPtr)p_script);
}
pointer dictionary_operator_index(pointer fn, pointer p_self, pointer p_key) {
	return (pointer)((GDExtensionInterfaceDictionaryOperatorIndex)fn)((GDExtensionTypePtr)p_self, (GDExtensionConstVariantPtr)p_key);
}
void object_method_bind_call(pointer fn, pointer p_method_bind, pointer p_instance, pointer p_args, GDExtensionInt count, pointer r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceObjectMethodBindCall)fn)((GDExtensionMethodBindPtr)p_method_bind, (GDExtensionObjectPtr)p_instance, (GDExtensionConstVariantPtr)p_args, count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
void object_method_bind_ptrcall(pointer fn, pointer p_method_bind, pointer p_instance, pointer p_args, pointer r_ret) {
	((GDExtensionInterfaceObjectMethodBindPtrcall)fn)((GDExtensionMethodBindPtr)p_method_bind, (GDExtensionObjectPtr)p_instance, (GDExtensionConstVariantPtr)p_args, (GDExtensionUninitializedVariantPtr)r_ret);
}
void object_destroy(pointer fn, pointer p_self) {
	((GDExtensionInterfaceObjectDestroy)fn)((GDExtensionObjectPtr)p_self);
}
pointer global_get_singleton(pointer fn, pointer p_name) {
	return (pointer)((GDExtensionInterfaceGlobalGetSingleton)fn)((GDExtensionConstStringNamePtr)p_name);
}
pointer object_get_instance_binding(pointer fn, pointer p_o, pointer p_library, void *p_callbacks) {
	return (pointer)((GDExtensionInterfaceObjectGetInstanceBinding)fn)((GDExtensionObjectPtr)p_o, (void*)p_library, (const GDExtensionInstanceBindingCallbacks *)p_callbacks);
}
void object_set_instance_binding(pointer fn, pointer p_o, pointer p_library, pointer p_binding, void *p_callbacks) {
	((GDExtensionInterfaceObjectSetInstanceBinding)fn)((GDExtensionObjectPtr)p_o, (void*)p_library, (void *)p_binding, (const GDExtensionInstanceBindingCallbacks *)p_callbacks);
}
void object_free_instance_binding(pointer fn, pointer p_o, pointer p_library) {
	((GDExtensionInterfaceObjectFreeInstanceBinding)fn)((GDExtensionObjectPtr)p_o, (void*)p_library);
}
void object_set_instance(pointer fn, pointer p_o, pointer p_classname, pointer p_instance) {
	((GDExtensionInterfaceObjectSetInstance)fn)((GDExtensionObjectPtr)p_o, (GDExtensionConstStringNamePtr)p_classname, (GDExtensionObjectPtr)p_instance);
}
void object_get_class_name(pointer fn, pointer p_o, pointer p_token, pointer r_ret) {
	((GDExtensionInterfaceObjectGetClassName)fn)((GDExtensionObjectPtr)p_o, (void *)p_token, (GDExtensionUninitializedStringPtr)r_ret);
}

pointer classdb_construct_object(pointer fn, pointer p_classname) {
	return (pointer)((GDExtensionInterfaceClassdbConstructObject)fn)((GDExtensionConstStringNamePtr)p_classname);
}
*/
import "C"

import (
	"errors"
	"runtime"
	"runtime/cgo"
	"unsafe"

	gd "grow.graphics/gd/internal"
	internal "grow.graphics/gd/internal"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

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
		var frame = call.New()
		defer frame.Free()
		return uintptr(C.get_native_struct_size(
			C.uintptr_t(uintptr(get_native_struct_size)),
			C.uintptr_t(call.Arg(frame, name.Pointer()).Uintptr()),
		))
	}
	variant_new_copy := dlsymGD("variant_new_copy")
	API.Variants.NewCopy = func(ctx gd.Context, self gd.Variant) gd.Variant {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_dest = call.Ret[[3]uintptr](frame)
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(r_dest.Uintptr()),
			C.uintptr_t(p_self.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_dest.Get())
		frame.Free()
		return ret
	}
	variant_new_nil := dlsymGD("variant_new_nil")
	API.Variants.NewNil = func(ctx gd.Context) gd.Variant {
		var frame = call.New()
		var r_dest = call.Ret[[3]uintptr](frame)
		C.variant_new_nil(
			C.uintptr_t(uintptr(variant_new_nil)),
			C.uintptr_t(r_dest.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_dest.Get())
		frame.Free()
		return ret
	}
	variant_destroy := dlsymGD("variant_destroy")
	API.Variants.Destroy = func(self gd.Variant) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		C.variant_destroy(
			C.uintptr_t(uintptr(variant_destroy)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
	}
	variant_call := dlsymGD("variant_call")
	API.Variants.Call = func(ctx gd.Context, self gd.Variant, method gd.StringName, args ...gd.Variant) (gd.Variant, error) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_method = call.Arg(frame, method.Pointer())
		for _, arg := range args {
			call.Arg(frame, arg.Pointer())
		}
		var r_ret = call.Ret[[3]uintptr](frame)
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
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_call_static := dlsymGD("variant_call_static")
	API.Variants.CallStatic = func(ctx gd.Context, vtype gd.VariantType, method gd.StringName, args ...gd.Variant) (gd.Variant, error) {
		var frame = call.New()
		var p_method = call.Arg(frame, method.Pointer())
		for _, arg := range args {
			call.Arg(frame, arg.Pointer())
		}
		var r_ret = call.Ret[[3]uintptr](frame)
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
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_evaluate := dlsymGD("variant_evaluate")
	API.Variants.Evaluate = func(ctx gd.Context, operator gd.Operator, a, b gd.Variant) (ret gd.Variant, ok bool) {
		var frame = call.New()
		var p_a = call.Arg(frame, a.Pointer())
		var p_b = call.Arg(frame, b.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_evaluate(
			C.uintptr_t(uintptr(variant_evaluate)),
			C.GDExtensionVariantOperator(operator),
			C.uintptr_t(p_a.Uintptr()),
			C.uintptr_t(p_b.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, r_valid.Get()
	}
	variant_set := dlsymGD("variant_set")
	API.Variants.Set = func(self, key, val gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
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
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
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
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
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
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
		var r_oob = call.Ret[bool](frame)
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
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_get(
			C.uintptr_t(uintptr(variant_get)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_named := dlsymGD("variant_get_named")
	API.Variants.GetNamed = func(ctx gd.Context, self gd.Variant, key gd.StringName) (val gd.Variant, ok bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_get_named(
			C.uintptr_t(uintptr(variant_get_named)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_keyed := dlsymGD("variant_get_keyed")
	API.Variants.GetKeyed = func(ctx gd.Context, self, key gd.Variant) (val gd.Variant, ok bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_get_keyed(
			C.uintptr_t(uintptr(variant_get_keyed)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		val = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		ok = r_valid.Get()
		frame.Free()
		return
	}
	variant_get_indexed := dlsymGD("variant_get_indexed")
	API.Variants.GetIndexed = func(ctx gd.Context, self gd.Variant, index gd.Int) (gd.Variant, bool, bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		var r_oob = call.Ret[bool](frame)
		C.variant_get_indexed(
			C.uintptr_t(uintptr(variant_get_indexed)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
			C.uintptr_t(r_oob.Uintptr()),
		)
		frame.Free()
		return mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get()), r_valid.Get(), r_oob.Get()
	}
	variant_iter_init := dlsymGD("variant_iter_init")
	API.Variants.IteratorInitialize = func(ctx gd.Context, self gd.Variant) (gd.Variant, bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_valid = call.Ret[bool](frame)
		var r_iter = call.Ret[[3]uintptr](frame)
		C.variant_iter_init(
			C.uintptr_t(uintptr(variant_iter_init)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(r_iter.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_iter.Get())
		var ok = r_valid.Get()
		frame.Free()
		return ret, ok
	}
	variant_iter_next := dlsymGD("variant_iter_next")
	API.Variants.IteratorNext = func(self gd.Variant, iter gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_iter = call.Arg(frame, iter.Pointer())
		var r_valid = call.Ret[bool](frame)
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
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_iter = call.Arg(frame, iter.Pointer())
		var r_valid = call.Ret[bool](frame)
		var r_ret = call.Ret[[3]uintptr](frame)
		C.variant_iter_get(
			C.uintptr_t(uintptr(variant_iter_get)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_iter.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		var ok = r_valid.Get()
		frame.Free()
		return ret, ok
	}
	variant_hash := dlsymGD("variant_hash")
	API.Variants.Hash = func(self gd.Variant) gd.Int {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_hash(
			C.uintptr_t(uintptr(variant_hash)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		return gd.Int(ret)
	}
	variant_recursive_hash := dlsymGD("variant_recursive_hash")
	API.Variants.RecursiveHash = func(self gd.Variant, count gd.Int) gd.Int {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
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
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_other = call.Arg(frame, other.Pointer())
		var ret = C.variant_hash_compare(
			C.uintptr_t(uintptr(variant_hash_compare)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_other.Uintptr()),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_booleanize := dlsymGD("variant_booleanize")
	API.Variants.Booleanize = func(self gd.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_booleanize(
			C.uintptr_t(uintptr(variant_booleanize)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_duplicate := dlsymGD("variant_duplicate")
	API.Variants.Duplicate = func(self gd.Variant, deep bool) gd.Variant {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
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
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](self.Context(), self.API, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_stringify := dlsymGD("variant_stringify")
	API.Variants.Stringify = func(self gd.Variant) gd.String {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_ret = call.Ret[uintptr](frame)
		C.variant_stringify(
			C.uintptr_t(uintptr(variant_stringify)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.String, uintptr](self.Context(), self.API, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_get_type := dlsymGD("variant_get_type")
	API.Variants.GetType = func(self gd.Variant) gd.VariantType {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_get_type(
			C.uintptr_t(uintptr(variant_get_type)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
		return gd.VariantType(ret)
	}
	variant_has_method := dlsymGD("variant_has_method")
	API.Variants.HasMethod = func(self gd.Variant, method gd.StringName) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_method = call.Arg(frame, method.Pointer())
		var ret = C.variant_has_method(
			C.uintptr_t(uintptr(variant_has_method)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_has_member := dlsymGD("variant_has_member")
	API.Variants.HasMember = func(self gd.Variant, member gd.StringName) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_member = call.Arg(frame, member.Pointer())
		var ret = C.variant_has_member(
			C.uintptr_t(uintptr(variant_has_member)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_member.Uintptr()),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_has_key := dlsymGD("variant_has_key")
	API.Variants.HasKey = func(self gd.Variant, key gd.Variant) (hasKey, valid bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_valid = call.Ret[bool](frame)
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
		var frame = call.New()
		var r_ret = call.Ret[uintptr](frame)
		C.variant_get_type_name(
			C.uintptr_t(uintptr(variant_get_type_name)),
			C.GDExtensionVariantType(vtype),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.String, uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret
	}
	variant_can_convert := dlsymGD("variant_can_convert")
	API.Variants.CanConvert = func(self gd.Variant, toType gd.VariantType) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_can_convert(
			C.uintptr_t(uintptr(variant_can_convert)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionVariantType(toType),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	variant_can_convert_strict := dlsymGD("variant_can_convert_strict")
	API.Variants.CanConvertStrict = func(self gd.Variant, toType gd.VariantType) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var ret = C.variant_can_convert_strict(
			C.uintptr_t(uintptr(variant_can_convert_strict)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionVariantType(toType),
		)
		frame.Free()
		if ret != 0 {
			return true
		}
		return false
	}
	get_variant_from_type_constructor := dlsymGD("get_variant_from_type_constructor")
	API.Variants.FromTypeConstructor = func(vt gd.VariantType) func(ret call.Ptr[[3]uintptr], arg call.Any) {
		fn := C.get_variant_from_type_constructor(
			C.uintptr_t(uintptr(get_variant_from_type_constructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret call.Ptr[[3]uintptr], arg call.Any) {
			C.call_variant_from_type_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	get_variant_to_type_constructor := dlsymGD("get_variant_to_type_constructor")
	API.Variants.ToTypeConstructor = func(vt gd.VariantType) func(ret call.Any, arg call.Ptr[[3]uintptr]) {
		fn := C.get_variant_to_type_constructor(
			C.uintptr_t(uintptr(get_variant_to_type_constructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret call.Any, arg call.Ptr[[3]uintptr]) {
			C.call_variant_to_type_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_operator_evaluator := dlsymGD("variant_get_ptr_operator_evaluator")
	API.Variants.PointerOperatorEvaluator = func(op gd.Operator, a, b gd.VariantType) func(a, b, ret call.Any) {
		fn := C.variant_get_ptr_operator_evaluator(
			C.uintptr_t(uintptr(variant_get_ptr_operator_evaluator)),
			C.GDExtensionVariantOperator(op),
			C.GDExtensionVariantType(a),
			C.GDExtensionVariantType(b),
		)
		return func(a, b, ret call.Any) {
			C.call_variant_ptr_operator_evaluator(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(a.Uintptr()),
				C.uintptr_t(b.Uintptr()),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_get_ptr_builtin_method := dlsymGD("variant_get_ptr_builtin_method")
	API.Variants.GetPointerBuiltinMethod = func(vt gd.VariantType, sn gd.StringName, i gd.Int) func(base call.Any, args call.Args, ret call.Any, c int32) {
		var frame = call.New()
		p_method := call.Arg(frame, sn.Pointer())
		fn := C.variant_get_ptr_builtin_method(
			C.uintptr_t(uintptr(variant_get_ptr_builtin_method)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
			C.GDExtensionInt(i),
		)
		frame.Free()
		return func(base call.Any, args call.Args, ret call.Any, c int32) {
			C.call_variant_ptr_builtin_method(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(args.Uintptr()),
				C.uintptr_t(ret.Uintptr()),
				C.int32_t(c),
			)
		}
	}
	variant_get_ptr_constructor := dlsymGD("variant_get_ptr_constructor")
	API.Variants.GetPointerConstructor = func(vt gd.VariantType, index int32) func(ret call.Any, args call.Args) {
		fn := C.variant_get_ptr_constructor(
			C.uintptr_t(uintptr(variant_get_ptr_constructor)),
			C.GDExtensionVariantType(vt),
			C.int32_t(index),
		)
		return func(ret call.Any, args call.Args) {
			C.call_variant_ptr_constructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(args.Uintptr()),
			)
		}
	}
	variant_get_ptr_destructor := dlsymGD("variant_get_ptr_destructor")
	API.Variants.GetPointerDestructor = func(vt gd.VariantType) func(ret call.Any) {
		fn := C.variant_get_ptr_destructor(
			C.uintptr_t(uintptr(variant_get_ptr_destructor)),
			C.GDExtensionVariantType(vt),
		)
		return func(ret call.Any) {
			C.call_variant_ptr_destructor(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_construct := dlsymGD("variant_construct")
	API.Variants.Construct = func(ctx gd.Context, t gd.VariantType, args ...gd.Variant) (gd.Variant, error) {
		var frame = call.New()
		for _, arg := range args {
			call.Arg(frame, arg.Pointer())
		}
		var r_ret = call.Ret[[3]uintptr](frame)
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
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, nil
	}
	variant_get_ptr_setter := dlsymGD("variant_get_ptr_setter")
	API.Variants.GetPointerSetter = func(vt gd.VariantType, sn gd.StringName) func(base, arg call.Any) {
		var frame = call.New()
		p_method := call.Arg(frame, sn.Pointer())
		fn := C.variant_get_ptr_setter(
			C.uintptr_t(uintptr(variant_get_ptr_setter)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		return func(base, arg call.Any) {
			C.call_variant_ptr_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_getter := dlsymGD("variant_get_ptr_getter")
	API.Variants.GetPointerGetter = func(vt gd.VariantType, sn gd.StringName) func(base call.Any, ret call.Any) {
		var frame = call.New()
		p_method := call.Arg(frame, sn.Pointer())
		fn := C.variant_get_ptr_getter(
			C.uintptr_t(uintptr(variant_get_ptr_getter)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
		)
		frame.Free()
		return func(base call.Any, ret call.Any) {
			C.call_variant_ptr_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_get_ptr_indexed_setter := dlsymGD("variant_get_ptr_indexed_setter")
	API.Variants.GetPointerIndexedSetter = func(vt gd.VariantType) func(base call.Any, index gd.Int, arg call.Any) {
		fn := C.variant_get_ptr_indexed_setter(
			C.uintptr_t(uintptr(variant_get_ptr_indexed_setter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base call.Any, index gd.Int, arg call.Any) {
			C.call_variant_ptr_indexed_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.GDExtensionInt(index),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_indexed_getter := dlsymGD("variant_get_ptr_indexed_getter")
	API.Variants.GetPointerIndexedGetter = func(vt gd.VariantType) func(base call.Any, index gd.Int, ret call.Any) {
		fn := C.variant_get_ptr_indexed_getter(
			C.uintptr_t(uintptr(variant_get_ptr_indexed_getter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base call.Any, index gd.Int, ret call.Any) {
			C.call_variant_ptr_indexed_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.GDExtensionInt(index),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_get_ptr_keyed_setter := dlsymGD("variant_get_ptr_keyed_setter")
	API.Variants.GetPointerKeyedSetter = func(vt gd.VariantType) func(base call.Any, key call.Any, arg call.Any) {
		fn := C.variant_get_ptr_keyed_setter(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_setter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base call.Any, key call.Any, arg call.Any) {
			C.call_variant_ptr_keyed_setter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(key.Uintptr()),
				C.uintptr_t(arg.Uintptr()),
			)
		}
	}
	variant_get_ptr_keyed_getter := dlsymGD("variant_get_ptr_keyed_getter")
	API.Variants.GetPointerKeyedGetter = func(vt gd.VariantType) func(base call.Any, key call.Any, ret call.Any) {
		fn := C.variant_get_ptr_keyed_getter(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_getter)),
			C.GDExtensionVariantType(vt),
		)
		return func(base call.Any, key call.Any, ret call.Any) {
			C.call_variant_ptr_keyed_getter(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(key.Uintptr()),
				C.uintptr_t(ret.Uintptr()),
			)
		}
	}
	variant_get_ptr_keyed_checker := dlsymGD("variant_get_ptr_keyed_checker")
	API.Variants.GetPointerKeyedChecker = func(vt gd.VariantType) func(base call.Any, key call.Any) uint32 {
		fn := C.variant_get_ptr_keyed_checker(
			C.uintptr_t(uintptr(variant_get_ptr_keyed_checker)),
			C.GDExtensionVariantType(vt),
		)
		return func(base call.Any, key call.Any) uint32 {
			return uint32(C.call_variant_ptr_keyed_checker(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(base.Uintptr()),
				C.uintptr_t(key.Uintptr()),
			))
		}
	}
	variant_get_constant_value := dlsymGD("variant_get_constant_value")
	API.Variants.GetConstantValue = func(ctx gd.Context, vt gd.VariantType, sn gd.StringName) gd.Variant {
		var frame = call.New()
		p_method := call.Arg(frame, sn.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		C.variant_get_constant_value(
			C.uintptr_t(uintptr(variant_get_constant_value)),
			C.GDExtensionVariantType(vt),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](nil, nil, r_ret.Get())
		frame.Free()
		return ret
	}
	variant_get_ptr_utility_function := dlsymGD("variant_get_ptr_utility_function")
	API.Variants.GetPointerUtilityFunction = func(sn gd.StringName, hash gd.Int) func(ret call.Any, args call.Args, c int32) {
		var frame = call.New()
		p_method := call.Arg(frame, sn.Pointer())
		fn := C.variant_get_ptr_utility_function(
			C.uintptr_t(uintptr(variant_get_ptr_utility_function)),
			C.uintptr_t(p_method.Uintptr()),
			C.GDExtensionInt(hash),
		)
		frame.Free()
		return func(ret call.Any, args call.Args, c int32) {
			C.call_variant_ptr_utility_function(
				C.uintptr_t(uintptr(fn)),
				C.uintptr_t(ret.Uintptr()),
				C.uintptr_t(args.Uintptr()),
				C.int32_t(c),
			)
		}
	}
	string_new_with_utf8_chars_and_len := dlsymGD("string_new_with_utf8_chars_and_len")
	API.Strings.New = func(ctx gd.Context, s string) gd.String {
		var str = C.CString(s)
		var frame = call.New()
		var r_ret = call.Ret[uintptr](frame)
		C.string_new_with_utf8_chars_and_len(
			C.uintptr_t(uintptr(string_new_with_utf8_chars_and_len)),
			C.uintptr_t(r_ret.Uintptr()),
			str,
			C.GDExtensionInt(len(s)),
		)
		var ret = mmm.Make[gd.API, gd.String, uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		C.free(unsafe.Pointer(str))
		return ret
	}
	string_to_utf8_chars := dlsymGD("string_to_utf8_chars")
	API.Strings.Get = func(s gd.String) string {
		var frame = call.New()
		var p_self = call.Arg(frame, s.Pointer())
		var length = s.Length()
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
		var frame = call.New()
		var p_self = call.Arg(frame, s.Pointer())
		*C.string_operator_index(
			C.uintptr_t(uintptr(string_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
		) = C.rune(val)
		frame.Free()
	}
	string_operator_index_const := dlsymGD("string_operator_index")
	API.Strings.Index = func(s gd.String, index gd.Int) rune {
		var frame = call.New()
		var p_self = call.Arg(frame, s.Pointer())
		var ret = C.string_operator_index_const(
			C.uintptr_t(uintptr(string_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(index),
		)
		frame.Free()
		return rune(*ret)
	}
	string_operator_plus_eq_string := dlsymGD("string_operator_plus_eq_string")
	API.Strings.Append = func(s gd.String, other gd.String) {
		var frame = call.New()
		var p_self = call.Arg(frame, s.Pointer())
		var p_other = call.Arg(frame, other.Pointer())
		C.string_operator_plus_eq_string(
			C.uintptr_t(uintptr(string_operator_plus_eq_string)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_other.Uintptr()),
		)
		frame.Free()
	}
	string_operator_plus_eq_char := dlsymGD("string_operator_plus_eq_char")
	API.Strings.AppendRune = func(s gd.String, other rune) {
		var frame = call.New()
		var p_self = call.Arg(frame, s.Pointer())
		C.string_operator_plus_eq_char(
			C.uintptr_t(uintptr(string_operator_plus_eq_char)),
			C.uintptr_t(p_self.Uintptr()),
			C.rune(other),
		)
		frame.Free()
	}
	string_resize := dlsymGD("string_resize")
	API.Strings.Resize = func(s gd.String, size gd.Int) {
		var frame = call.New()
		var p_self = call.Arg(frame, s.Pointer())
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
		var frame = call.New()
		var r_ret = call.Ret[uintptr](frame)
		C.string_name_new_with_utf8_chars_and_len(
			C.uintptr_t(uintptr(string_name_new_with_utf8_chars_and_len)),
			C.uintptr_t(r_ret.Uintptr()),
			str,
			C.GDExtensionInt(len(s)),
		)
		var ret = mmm.Make[gd.API, gd.StringName, uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		C.free(unsafe.Pointer(str))
		return ret
	}
	xml_parser_open_buffer := dlsymGD("xml_parser_open_buffer")
	API.XMLParser.OpenBuffer = func(x gd.XMLParser, b []byte) error {
		var pin runtime.Pinner
		pin.Pin(&b[0])
		x.Context().Defer(func() {
			pin.Unpin()
		})
		var frame = call.New()
		var p_self = call.Arg(frame, x.Pointer())
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
	API.FileAccess.StoreBuffer = func(f gd.FileAccess, b []byte) {
		var frame = call.New()
		var p_self = call.Arg(frame, f.Pointer())
		C.file_access_store_buffer(
			C.uintptr_t(uintptr(file_access_store_buffer)),
			C.uintptr_t(p_self.Uintptr()),
			(C.bytes)(unsafe.Pointer(&b[0])),
			C.uint64_t(len(b)),
		)
		frame.Free()
	}
	file_access_get_buffer := dlsymGD("file_access_get_buffer")
	API.FileAccess.GetBuffer = func(f gd.FileAccess, b []byte) int {
		var frame = call.New()
		var p_self = call.Arg(frame, f.Pointer())
		var length = C.file_access_get_buffer(
			C.uintptr_t(uintptr(file_access_get_buffer)),
			C.uintptr_t(p_self.Uintptr()),
			(C.bytes)(unsafe.Pointer(&b[0])),
			C.uint64_t(len(b)),
		)
		frame.Free()
		return int(length)
	}
	API.PackedByteArray = makePackedFunctions[gd.PackedByteArray, byte]("byte_array")
	API.PackedColorArray = makePackedFunctions[gd.PackedColorArray, gd.Color]("color_array")
	API.PackedFloat32Array = makePackedFunctions[gd.PackedFloat32Array, float32]("float32_array")
	API.PackedFloat64Array = makePackedFunctions[gd.PackedFloat64Array, float64]("float64_array")
	API.PackedInt32Array = makePackedFunctions[gd.PackedInt32Array, int32]("int32_array")
	API.PackedInt64Array = makePackedFunctions[gd.PackedInt64Array, int64]("int64_array")
	packed_string_array_operator_index_const := dlsymGD("packed_string_array_operator_index_const")
	API.PackedStringArray.Index = func(ctx gd.Context, psa gd.PackedStringArray, i int64) gd.String {
		var frame = call.New()
		var p_self = call.Arg(frame, psa.Pointer())
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_string_array_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		frame.Free()
		return mmm.Make[gd.API, gd.String, uintptr](nil, psa.API, uintptr(ret)).Copy(ctx)
	}
	packed_string_array_operator_index := dlsymGD("packed_string_array_operator_index")
	API.PackedStringArray.SetIndex = func(psa gd.PackedStringArray, i int64, v gd.String) {
		var frame = call.New()
		var p_self = call.Arg(frame, psa.Pointer())
		var ptr = C.packed_T_operator_index(
			C.uintptr_t(uintptr(packed_string_array_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		*(*gd.String)(unsafe.Pointer(uintptr(ptr))) = v
		frame.Free()
	}
	API.PackedStringArray.CopyAsSlice = func(ctx gd.Context, psa gd.PackedStringArray) []gd.String {
		var size = psa.Size()
		if size == 0 {
			return nil
		}
		var frame = call.New()
		var p_self = call.Arg(frame, psa.Pointer())
		var ret = C.packed_T_operator_index_const(
			C.uintptr_t(uintptr(packed_string_array_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(0),
		)
		var slice = make([]gd.String, size)
		for _, ptr := range unsafe.Slice((*uintptr)(unsafe.Pointer(uintptr(ret))), size) {
			slice = append(slice, mmm.Make[gd.API, gd.String, uintptr](nil, psa.API, ptr).Copy(ctx))
		}
		frame.Free()
		return slice
	}
	API.PackedVector2Array = makePackedFunctions[gd.PackedVector2Array, gd.Vector2]("vector2_array")
	API.PackedVector3Array = makePackedFunctions[gd.PackedVector3Array, gd.Vector3]("vector3_array")
	array_operator_index_const := dlsymGD("array_operator_index_const")
	API.Array.Index = func(ctx gd.Context, a gd.Array, i int64) gd.Variant {
		var frame = call.New()
		var p_self = call.Arg(frame, a.Pointer())
		var r_ret = C.array_operator_index_const(
			C.uintptr_t(uintptr(array_operator_index_const)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		var ptr = (*[3]uintptr)(unsafe.Pointer(uintptr(r_ret)))
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](nil, ctx.API(), *ptr)
		frame.Free()
		return ret.Copy(ctx)
	}
	array_operator_index := dlsymGD("array_operator_index")
	API.Array.SetIndex = func(a gd.Array, i gd.Int, v gd.Variant) {
		var frame = call.New()
		var p_self = call.Arg(frame, a.Pointer())
		var ptr = C.array_operator_index(
			C.uintptr_t(uintptr(array_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		var p_copy = call.Ret[[3]uintptr](frame)
		var p_value = call.Arg(frame, v.Pointer())
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(p_copy.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
		)
		*(*[3]uintptr)(unsafe.Pointer(uintptr(ptr))) = p_copy.Get()
		frame.Free()
	}
	array_ref := dlsymGD("array_ref")
	API.Array.Set = func(self gd.Array, from gd.Array) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_from = call.Arg(frame, from.Pointer())
		C.array_ref(
			C.uintptr_t(uintptr(array_ref)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_from.Uintptr()),
		)
		frame.Free()
	}
	array_set_typed := dlsymGD("array_set_typed")
	API.Array.SetTyped = func(self gd.Array, t gd.VariantType, className gd.StringName, script gd.Script) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_className = call.Arg(frame, className.Pointer())
		var p_script = call.Arg(frame, script.Pointer())
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
		var frame = call.New()
		var p_self = call.Arg(frame, d.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var ptr = C.dictionary_operator_index(
			C.uintptr_t(uintptr(dictionary_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
		)
		var r_ret = *(*[3]uintptr)(unsafe.Pointer(uintptr(ptr)))
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret)
		frame.Free()
		return ret
	}
	API.Dictionary.SetIndex = func(dict gd.Dictionary, key, val gd.Variant) {
		var frame = call.New()
		var p_self = call.Arg(frame, dict.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var ptr = C.dictionary_operator_index(
			C.uintptr_t(uintptr(dictionary_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
		)
		var p_copy = call.Ret[[3]uintptr](frame)
		var p_value = call.Arg(frame, val.Pointer())
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(p_copy.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
		)
		*(*[3]uintptr)(unsafe.Pointer(uintptr(ptr))) = p_copy.Get()
		frame.Free()
	}
	object_method_bind_call := dlsymGD("object_method_bind_call")
	API.Object.MethodBindCall = func(ctx gd.Context, method gd.MethodBind, obj gd.Object, arg ...gd.Variant) (gd.Variant, error) {
		var frame = call.New()
		for _, a := range arg {
			call.Arg(frame, a.Pointer())
		}
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.object_method_bind_call(
			C.uintptr_t(uintptr(object_method_bind_call)),
			C.uintptr_t(method),
			C.uintptr_t(obj.Pointer()),
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
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, nil
	}
	object_method_bind_ptrcall := dlsymGD("object_method_bind_ptrcall")
	API.Object.MethodBindPointerCall = func(method gd.MethodBind, obj gd.Object, arg call.Args, ret call.Any) {
		C.object_method_bind_ptrcall(
			C.uintptr_t(uintptr(object_method_bind_ptrcall)),
			C.uintptr_t(method),
			C.uintptr_t(obj.Pointer()),
			C.uintptr_t(arg.Uintptr()),
			C.uintptr_t(ret.Uintptr()),
		)
	}
	object_destroy := dlsymGD("object_destroy")
	API.Object.Destroy = func(o gd.Object) {
		C.object_destroy(
			C.uintptr_t(uintptr(object_destroy)),
			C.uintptr_t(o.Pointer()),
		)
	}
	global_get_singleton := dlsymGD("global_get_singleton")
	API.Object.GetSingleton = func(ctx gd.Context, name gd.StringName) gd.Object {
		var frame = call.New()
		var p_name = call.Arg(frame, name.Pointer())
		var ret = C.global_get_singleton(
			C.uintptr_t(uintptr(global_get_singleton)),
			C.uintptr_t(p_name.Uintptr()),
		)
		var obj gd.Object
		obj.SetPointer(mmm.Make[gd.API, gd.Pointer, uintptr](nil, ctx.API(), uintptr(ret)))
		frame.Free()
		return obj
	}
	object_get_instance_binding := dlsymGD("object_get_instance_binding")
	API.Object.GetInstanceBinding = func(o gd.Object, et gd.ExtensionToken, ibt gd.InstanceBindingType) any {
		var ret = C.object_get_instance_binding(
			C.uintptr_t(uintptr(object_get_instance_binding)),
			C.uintptr_t(o.Pointer()),
			C.uintptr_t(et),
			unsafe.Pointer(ibt),
		)
		return cgo.Handle(ret).Value()
	}
	object_set_instance_binding := dlsymGD("object_set_instance_binding")
	API.Object.SetInstanceBinding = func(o gd.Object, et gd.ExtensionToken, val any, ibt gd.InstanceBindingType) {
		p_val := cgo.NewHandle(val)
		C.object_set_instance_binding(
			C.uintptr_t(uintptr(object_set_instance_binding)),
			C.uintptr_t(o.Pointer()),
			C.uintptr_t(et),
			C.uintptr_t(p_val),
			unsafe.Pointer(ibt),
		)
	}
	object_free_instance_binding := dlsymGD("object_free_instance_binding")
	API.Object.FreeInstanceBinding = func(o gd.Object, et gd.ExtensionToken) {
		C.object_free_instance_binding(
			C.uintptr_t(uintptr(object_free_instance_binding)),
			C.uintptr_t(o.Pointer()),
			C.uintptr_t(et),
		)
	}
	object_set_instance := dlsymGD("object_set_instance")
	API.Object.SetInstance = func(o gd.Object, sn gd.StringName, a any) {
		var frame = call.New()
		var p_sn = call.Arg(frame, sn.Pointer())
		var p_val = cgo.NewHandle(a)
		C.object_set_instance(
			C.uintptr_t(uintptr(object_set_instance)),
			C.uintptr_t(o.Pointer()),
			C.uintptr_t(p_sn.Uintptr()),
			C.uintptr_t(p_val),
		)
		frame.Free()
	}
	object_get_class_name := dlsymGD("object_get_class_name")
	API.Object.GetClassName = func(ctx gd.Context, o gd.Object, token internal.ExtensionToken) gd.String {
		var frame = call.New()
		var r_ret = call.Ret[uintptr](frame)
		C.object_get_class_name(
			C.uintptr_t(uintptr(object_get_class_name)),
			C.uintptr_t(o.Pointer()),
			C.uintptr_t(token),
			C.uintptr_t(r_ret.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.String, uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret
	}

	classdb_construct_object := dlsymGD("classdb_construct_object")
	API.ClassDB.ConstructObject = func(ctx gd.Context, name gd.StringName) gd.Object {
		var frame = call.New()
		var obj gd.Object
		obj.SetPointer(mmm.Make[gd.API, gd.Pointer, uintptr](ctx, ctx.API(), uintptr(C.classdb_construct_object(
			C.uintptr_t(uintptr(classdb_construct_object)),
			C.uintptr_t(call.Arg(frame, name.Pointer()).Uintptr()),
		))))
		frame.Free()
		return obj
	}

}

func makePackedFunctions[T gd.Packed, V comparable](prefix string) gd.PackedFunctionsFor[T, V] {
	var API gd.PackedFunctionsFor[T, V]
	packed_T_operator_index := dlsymGD("packed_" + prefix + "_operator_index")
	API.SetIndex = func(t T, i int64, v V) {
		var frame = call.New()
		var p_self = call.Arg(frame, t.Pointer())
		var ptr = C.packed_T_operator_index(
			C.uintptr_t(uintptr(packed_T_operator_index)),
			C.uintptr_t(p_self.Uintptr()),
			C.GDExtensionInt(i),
		)
		*(*V)(unsafe.Pointer(uintptr(ptr))) = v
		frame.Free()
	}
	packed_T_operator_index_const := dlsymGD("packed_" + prefix + "_operator_index_const")
	API.Index = func(t T, i int64) V {
		var frame = call.New()
		var p_self = call.Arg(frame, t.Pointer())
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
		var frame = call.New()
		var p_self = call.Arg(frame, t.Pointer())
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
