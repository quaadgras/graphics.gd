//go:build cgo

package gdunsafe

// #include "gd.h"
//
// GD_EXTENSION(go)
import "C"
import "unsafe"

type (
	gd_addr                 = C.gd_addr
	gd_shape                = C.gd_shape
	gd_float                = C.gd_float
	gd_initialization_level = C.gd_initialization_level
	gd_error                = C.gd_error

	gdVariant            = C.struct_Variant
	gdString             = C.struct_String
	gdVector2            = C.struct_Vector2
	gdVector2i           = C.struct_Vector2i
	gdRect2              = C.struct_Rect2
	gdRect2i             = C.struct_Rect2i
	gdVector3            = C.struct_Vector3
	gdVector3i           = C.struct_Vector3i
	gdTransform2D        = C.struct_Transform2D
	gdVector4            = C.struct_Vector4
	gdVector4i           = C.struct_Vector4i
	gdPlane              = C.struct_Plane
	gdQuaternion         = C.struct_Quaternion
	gdAABB               = C.struct_AABB
	gdBasis              = C.struct_Basis
	gdTransform3D        = C.struct_Transform3D
	gdProjection         = C.struct_Projection
	gdColor              = C.struct_Color
	gdCallable           = C.struct_Callable
	gdStringName         = C.struct_StringName
	gdNodePath           = C.struct_NodePath
	gdRID                = C.RID
	gdObject             = C.struct_Object
	gdSignal             = C.struct_Signal
	gdDictionary         = C.struct_Dictionary
	gdArray              = C.struct_Array
	gdPackedArray        = C.struct_PackedArray
	gdPackedByteArray    = C.struct_PackedByteArray
	gdPackedInt32Array   = C.struct_PackedInt32Array
	gdPackedInt64Array   = C.struct_PackedInt64Array
	gdPackedFloat32Array = C.struct_PackedFloat32Array
	gdPackedFloat64Array = C.struct_PackedFloat64Array
	gdPackedStringArray  = C.struct_PackedStringArray
	gdPackedVector2Array = C.struct_PackedVector2Array
	gdPackedVector3Array = C.struct_PackedVector3Array
	gdPackedColorArray   = C.struct_PackedColorArray
	gdPackedVector4Array = C.struct_PackedVector4Array

	gdRefCounted = C.struct_RefCounted
	gdScript     = C.struct_Script
	gdClassTag   = C.struct_ClassTag

	gd_constructor_id = C.gd_constructor_id
	gd_evaluator_id   = C.gd_evaluator_id
	gd_method_id      = C.gd_method_id
	gd_getter_id      = C.gd_getter_id
	gd_setter_id      = C.gd_setter_id
	gd_caller_id      = C.gd_caller_id
	gd_function_t     = C.gd_function_t

	gd_extension_object_id  = C.gd_extension_object_id
	gd_extension_method_id  = C.gd_extension_method_id
	gd_extension_class_id   = C.gd_extension_class_id
	gd_extension_task_id    = C.gd_extension_task_id
	gd_extension_script_id  = C.gd_extension_script_id
	gd_extension_binding_id = C.gd_extension_binding_id
	gd_extension_callable_t = C.gd_extension_callable_t

	gd_property_list_t     = C.gd_property_list_t
	gd_method_list_t       = C.gd_method_list_t
	gd_property_iterator_t = C.gd_property_iterator_t

	gd_encoding  = C.gd_encoding
	gd_log_level = C.gd_log_level

	gdVariantType      = C.VariantType
	gdVariantOperator  = C.VariantOperator
	gdMethodFlags      = C.MethodFlags
	gdArgumentMetadata = C.ArgumentMetadata
	gdObjectID         = C.ObjectID
)

func gd_addrOf(shape Shape, ptr unsafe.Pointer) gd_addr {
	return gd_addr(ptr)
}

//
// Version Information
//

func gd_version() gdString       { return C.gd_version() }
func gd_version_major() uint32   { return uint32(C.gd_version_major()) }
func gd_version_minor() uint32   { return uint32(C.gd_version_minor()) }
func gd_version_patch() uint32   { return uint32(C.gd_version_patch()) }
func gd_version_hexed() uint32   { return uint32(C.gd_version_hexed()) }
func gd_version_state() gdString { return C.gd_version_state() }
func gd_version_build() gdString { return C.gd_version_build() }
func gd_version_nanos() int64    { return int64(C.gd_version_nanos()) }
func gd_version_stamp() gdString { return C.gd_version_stamp() }

//
// Unsafe Engine Memory
//

func gd_sizeof(type_name gdStringName) int64  { return int64(C.gd_sizeof(type_name)) }
func gd_malloc(size int64, pad8 bool) gd_addr { return C.gd_malloc(C.int64_t(size), C.bool(pad8)) }
func gd_resize(addr gd_addr, size int64, pad8 bool) gd_addr {
	return C.gd_resize(addr, C.int64_t(size), C.bool(pad8))
}
func gd_memset(addr gd_addr, value byte, size int64) {
	C.gd_memset(addr, C.uint8_t(value), C.int64_t(size))
}

func gd_memory_bytes1(addr gd_addr) byte   { return byte(C.gd_memory_bytes1(addr)) }
func gd_memory_bytes2(addr gd_addr) uint16 { return uint16(C.gd_memory_bytes2(addr)) }
func gd_memory_bytes4(addr gd_addr) uint32 { return uint32(C.gd_memory_bytes4(addr)) }
func gd_memory_bytes8(addr gd_addr) uint64 { return uint64(C.gd_memory_bytes8(addr)) }

func gd_store_bytes1(addr gd_addr, v byte)   { C.gd_store_bytes1(addr, C.uint8_t(v)) }
func gd_store_bytes2(addr gd_addr, v uint16) { C.gd_store_bytes2(addr, C.uint16_t(v)) }
func gd_store_bytes4(addr gd_addr, v uint32) { C.gd_store_bytes4(addr, C.uint32_t(v)) }
func gd_store_bytes8(addr gd_addr, v uint64) { C.gd_store_bytes8(addr, C.uint64_t(v)) }
func gd_store_pair64(addr gd_addr, a, b uint64) {
	C.gd_store_pair64(addr, C.uint64_t(a), C.uint64_t(b))
}
func gd_store_quad64(addr gd_addr, a, b, c, d uint64) {
	C.gd_store_quad64(addr, C.uint64_t(a), C.uint64_t(b), C.uint64_t(c), C.uint64_t(d))
}
func gd_store_octo64(addr gd_addr, a, b, c, d, e, f, g, h uint64) {
	C.gd_store_octo64(addr, C.uint64_t(a), C.uint64_t(b), C.uint64_t(c), C.uint64_t(d), C.uint64_t(e), C.uint64_t(f), C.uint64_t(g), C.uint64_t(h))
}

func gd_free(addr gd_addr) { C.gd_free(addr) }

//
// Variants
//

func gd_variant_zero(result *gdVariant) { C.gd_variant_zero(result) }
func gd_variant_copy(v_1, v_2, v_3 uint64, result *gdVariant, deep bool) {
	C.gd_variant_copy(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), result, C.bool(deep))
}
func gd_variant_hash(v_1, v_2, v_3 uint64, recursion_count int64) int64 {
	return int64(C.gd_variant_hash(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), C.int64_t(recursion_count)))
}
func gd_variant_bool(v_1, v_2, v_3 uint64) bool {
	return bool(C.gd_variant_bool(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3)))
}
func gd_variant_text(v_1, v_2, v_3 uint64) gdString {
	return C.gd_variant_text(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3))
}
func gd_variant_type(v_1, v_2, v_3 uint64) uint32 {
	return uint32(C.gd_variant_type(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3)))
}
func gd_variant_free(v_1, v_2, v_3 uint64) {
	C.gd_variant_free(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3))
}
func gd_variant_from(vtype uint32, result *gdVariant, shape gd_shape, args gd_addr) {
	C.gd_variant_from(C.VariantType(vtype), result, shape, args)
}
func gd_variant_make(t uint32, result *gdVariant, arg_count int64, args *gdVariant, err *gd_error) {
	C.gd_variant_make(C.VariantType(t), result, C.int64_t(arg_count), args, err)
}
func gd_variant_data(vtype uint32, v_1, v_2, v_3 uint64) gd_addr {
	return C.gd_variant_data(C.VariantType(vtype), C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3))
}
func gd_variant_call(v_1, v_2, v_3 uint64, method gdStringName, result *gdVariant, arg_count int64, args *gdVariant, err *gd_error) {
	C.gd_variant_call(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), method, result, C.int64_t(arg_count), args, err)
}
func gd_variant_eval(op uint32, a_1, a_2, a_3, b_1, b_2, b_3 uint64, result *gdVariant) bool {
	return bool(C.gd_variant_eval(C.VariantOperator(op), C.uint64_t(a_1), C.uint64_t(a_2), C.uint64_t(a_3), C.uint64_t(b_1), C.uint64_t(b_2), C.uint64_t(b_3), result))
}
func gd_variant_get_keyed(v_1, v_2, v_3, key_1, key_2, key_3 uint64, result *gdVariant) bool {
	return bool(C.gd_variant_get_keyed(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), C.uint64_t(key_1), C.uint64_t(key_2), C.uint64_t(key_3), result))
}
func gd_variant_get_index(v_1, v_2, v_3 uint64, idx int64, result *gdVariant, err *gd_error) bool {
	return bool(C.gd_variant_get_index(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), C.int64_t(idx), result, err))
}
func gd_variant_get_field(v_1, v_2, v_3 uint64, field gdStringName, result *gdVariant) bool {
	return bool(C.gd_variant_get_field(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), field, result))
}
func gd_variant_has_key(v_1, v_2, v_3, idx_1, idx_2, idx_3 uint64) bool {
	return bool(C.gd_variant_has_key(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), C.uint64_t(idx_1), C.uint64_t(idx_2), C.uint64_t(idx_3)))
}
func gd_variant_has_method(v_1, v_2, v_3 uint64, method gdStringName) bool {
	return bool(C.gd_variant_has_method(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), method))
}
func gd_variant_set_keyed(v_1, v_2, v_3, key_1, key_2, key_3, val_1, val_2, val_3 uint64) bool {
	return bool(C.gd_variant_set_keyed(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), C.uint64_t(key_1), C.uint64_t(key_2), C.uint64_t(key_3), C.uint64_t(val_1), C.uint64_t(val_2), C.uint64_t(val_3)))
}
func gd_variant_set_index(v_1, v_2, v_3 uint64, idx int64, val_1, val_2, val_3 uint64, err gd_addr) bool {
	return bool(C.gd_variant_set_index(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), C.int64_t(idx), C.uint64_t(val_1), C.uint64_t(val_2), C.uint64_t(val_3), err))
}
func gd_variant_set_field(v_1, v_2, v_3 uint64, field gdStringName, val_1, val_2, val_3 uint64) bool {
	return bool(C.gd_variant_set_field(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), field, C.uint64_t(val_1), C.uint64_t(val_2), C.uint64_t(val_3)))
}

//
// Builtin Types
//

func gd_evaluator(op uint32, a, b uint32) gd_evaluator_id {
	return C.gd_evaluator(C.VariantOperator(op), C.VariantType(a), C.VariantType(b))
}
func gd_setter(t uint32, property gdStringName) gd_setter_id {
	return C.gd_setter(C.VariantType(t), property)
}
func gd_getter(t uint32, property gdStringName) gd_getter_id {
	return C.gd_getter(C.VariantType(t), property)
}
func gd_constructor(t uint32, n int64) gd_constructor_id {
	return C.gd_constructor(C.VariantType(t), C.int64_t(n))
}
func gd_builtin_method(t uint32, method gdStringName, hash int64) gd_caller_id {
	return C.gd_builtin_method(C.VariantType(t), method, C.int64_t(hash))
}

func gd_builtin_call(self gd_addr, fn gd_caller_id, result gd_addr, shape gd_shape, args gd_addr) {
	C.gd_builtin_call(self, fn, result, shape, args)
}
func gd_builtin_make(constructor gd_constructor_id, result gd_addr, shape gd_shape, args gd_addr) {
	C.gd_builtin_make(constructor, result, shape, args)
}
func gd_builtin_free(t uint32, value gd_addr) {
	C.gd_builtin_free(C.VariantType(t), value)
}
func gd_builtin_from(vtype uint32, v_1, v_2, v_3 uint64, result gd_addr) {
	C.gd_builtin_from(C.VariantType(vtype), C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), result)
}
func gd_builtin_eval(op gd_evaluator_id, result gd_addr, shape gd_shape, args gd_addr) {
	C.gd_builtin_eval(op, result, shape, args)
}
func gd_builtin_get_field(getter gd_getter_id, result gd_addr, shape gd_shape, args gd_addr) {
	C.gd_builtin_get_field(getter, result, shape, args)
}
func gd_builtin_get_array(vtype uint32, idx int64, result gd_addr, shape gd_shape, args gd_addr) {
	C.gd_builtin_get_array(C.VariantType(vtype), C.int64_t(idx), result, shape, args)
}
func gd_builtin_get_keyed(vtype uint32, result gd_addr, shape gd_shape, args gd_addr) {
	C.gd_builtin_get_keyed(C.VariantType(vtype), result, shape, args)
}
func gd_builtin_set_field(setter gd_setter_id, shape gd_shape, args gd_addr) {
	C.gd_builtin_set_field(setter, shape, args)
}
func gd_builtin_set_array(vtype uint32, idx int64, shape gd_shape, args gd_addr) {
	C.gd_builtin_set_array(C.VariantType(vtype), C.int64_t(idx), shape, args)
}
func gd_builtin_set_keyed(vtype uint32, shape gd_shape, args gd_addr) {
	C.gd_builtin_set_keyed(C.VariantType(vtype), shape, args)
}

//
// Variant Type Operations
//

func gd_variant_type_name(t uint32) gdString {
	return C.gd_variant_type_name(C.VariantType(t))
}
func gd_variant_type_call(t uint32, static_method gdStringName, result *gdVariant, arg_count int64, args *gdVariant, err *gd_error) {
	C.gd_variant_type_call(C.VariantType(t), static_method, result, C.int64_t(arg_count), args, err)
}
func gd_variant_type_convertable(t, to uint32, strict bool) bool {
	return bool(C.gd_variant_type_convertable(C.VariantType(t), C.VariantType(to), C.bool(strict)))
}
func gd_variant_type_has_property(t uint32, property gdStringName) bool {
	return bool(C.gd_variant_type_has_property(C.VariantType(t), property))
}
func gd_variant_type_setup_array(a gdArray, elem uint32, class_name gdStringName, v_1, v_2, v_3 uint64) {
	C.gd_variant_type_setup_array(a, C.VariantType(elem), class_name, C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3))
}
func gd_variant_type_setup_dictionary(d gdDictionary, key uint32, key_class gdStringName, ks_1, ks_2, ks_3 uint64, val uint32, val_class gdStringName, vs_1, vs_2, vs_3 uint64) {
	C.gd_variant_type_setup_dictionary(d, C.VariantType(key), key_class, C.uint64_t(ks_1), C.uint64_t(ks_2), C.uint64_t(ks_3), C.VariantType(val), val_class, C.uint64_t(vs_1), C.uint64_t(vs_2), C.uint64_t(vs_3))
}
func gd_variant_type_constant(t uint32, constant gdStringName, result gd_addr) {
	C.gd_variant_type_constant(C.VariantType(t), constant, result)
}

//
// Packed Arrays
//

func gd_packed_array_access(vtype uint32, pa_1, pa_2 uintptr, idx int64) gd_addr {
	return C.gd_packed_array_access(C.VariantType(vtype), C.uintptr_t(pa_1), C.uintptr_t(pa_2), C.int64_t(idx))
}
func gd_packed_array_modify(vtype uint32, pa_1, pa_2 uintptr, idx int64) gd_addr {
	return C.gd_packed_array_modify(C.VariantType(vtype), C.uintptr_t(pa_1), C.uintptr_t(pa_2), C.int64_t(idx))
}

//
// Arrays
//

func gd_array_get_index(a gdArray, i int64, result *gdVariant) {
	C.gd_array_get_index(a, C.int64_t(i), result)
}
func gd_array_set_index(a gdArray, i int64, v_1, v_2, v_3 uint64) {
	C.gd_array_set_index(a, C.int64_t(i), C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3))
}

//
// GlobalScope Functions
//

func gd_function(utility gdStringName, hash int64) gd_function_t {
	return C.gd_function(utility, C.int64_t(hash))
}
func gd_call(fn gd_function_t, result gd_addr, shape gd_shape, args gd_addr) {
	C.gd_call(fn, result, shape, args)
}

//
// ClassDB
//

func gd_method_list_make(count int64) gd_method_list_t {
	return C.gd_method_list_make(C.int64_t(count))
}
func gd_method_list_free(list gd_method_list_t) { C.gd_method_list_free(list) }
func gd_method_list_push(list gd_method_list_t, name gdStringName, call gd_extension_method_id, flags uint32, return_info gd_property_list_t, args_info gd_property_list_t, count int64, defaults gd_addr) {
	C.gd_method_list_push(list, name, call, C.MethodFlags(flags), return_info, args_info, C.int64_t(count), defaults)
}

func gd_property_list_make(count int64) gd_property_list_t {
	return C.gd_property_list_make(C.int64_t(count))
}
func gd_property_list_push(list gd_property_list_t, t uint32, name gdStringName, class_name gdStringName, hint uint32, hint_string gdString, usage uint32, meta uint32) {
	C.gd_property_list_push(list, C.VariantType(t), name, class_name, C.uint32_t(hint), hint_string, C.uint32_t(usage), C.ArgumentMetadata(meta))
}
func gd_property_list_free(list gd_property_list_t) { C.gd_property_list_free(list) }

func gd_classdb_register(class_name, parent gdStringName, id gd_extension_class_id, is_virtual, is_abstract, is_exposed, is_runtime bool, icon_path gdString) {
	C.gd_classdb_register(class_name, parent, id, C.bool(is_virtual), C.bool(is_abstract), C.bool(is_exposed), C.bool(is_runtime), icon_path)
}
func gd_classdb_register_methods(class_name gdStringName, methods gd_method_list_t) {
	C.gd_classdb_register_methods(class_name, methods)
}
func gd_classdb_register_constant(class_name, enum_name, constant_name gdStringName, value int64, bitfield bool) {
	C.gd_classdb_register_constant(class_name, enum_name, constant_name, C.int64_t(value), C.bool(bitfield))
}
func gd_classdb_register_property(class_name gdStringName, property gd_property_list_t, setter, getter gdStringName) {
	C.gd_classdb_register_property(class_name, property, setter, getter)
}
func gd_classdb_register_property_indexed(class_name gdStringName, property gd_property_list_t, setter, getter gdStringName, index int64) {
	C.gd_classdb_register_property_indexed(class_name, property, setter, getter, C.int64_t(index))
}
func gd_classdb_register_property_group(class_name gdStringName, group, prefix gdString) {
	C.gd_classdb_register_property_group(class_name, group, prefix)
}
func gd_classdb_register_property_sub_group(class_name gdStringName, subgroup, prefix gdString) {
	C.gd_classdb_register_property_sub_group(class_name, subgroup, prefix)
}
func gd_classdb_register_signal(class_name, signal gdStringName, args gd_property_list_t) {
	C.gd_classdb_register_signal(class_name, signal, args)
}
func gd_classdb_register_removal(class_name gdStringName) {
	C.gd_classdb_register_removal(class_name)
}

func gd_classdb_FileAccess_write(file gdObject, buf *byte, length int64) {
	C.gd_classdb_FileAccess_write(file, (*C.char)(unsafe.Pointer(buf)), C.int64_t(length))
}
func gd_classdb_FileAccess_read(file gdObject, buf *byte, cap_ int64) int64 {
	return int64(C.gd_classdb_FileAccess_read(file, (*C.char)(unsafe.Pointer(buf)), C.int64_t(cap_)))
}

func gd_classdb_Image_memory(img gdObject) gd_addr { return C.gd_classdb_Image_memory(img) }
func gd_classdb_Image_access(img gdObject, offset int64) byte {
	return byte(C.gd_classdb_Image_access(img, C.int64_t(offset)))
}

func gd_classdb_WorkerThreadPool_add_task(pool gdObject, task gd_extension_task_id, priority bool, description gdString) {
	C.gd_classdb_WorkerThreadPool_add_task(pool, task, C.bool(priority), description)
}
func gd_classdb_WorkerThreadPool_add_group_task(pool gdObject, task gd_extension_task_id, elements, arg int32, priority bool, description gdString) {
	C.gd_classdb_WorkerThreadPool_add_group_task(pool, task, C.int32_t(elements), C.int32_t(arg), C.bool(priority), description)
}

func gd_classdb_XMLParser_load(parser gdObject, buf *byte, cap_ int64) int64 {
	return int64(C.gd_classdb_XMLParser_load(parser, (*C.char)(unsafe.Pointer(buf)), C.int64_t(cap_)))
}

//
// Dictionaries
//

func gd_packed_dictionary_access(d gdDictionary, key_1, key_2, key_3 uint64, result *gdVariant) {
	C.gd_packed_dictionary_access(d, C.uint64_t(key_1), C.uint64_t(key_2), C.uint64_t(key_3), result)
}
func gd_packed_dictionary_modify(d gdDictionary, key_1, key_2, key_3, val_1, val_2, val_3 uint64) {
	C.gd_packed_dictionary_modify(d, C.uint64_t(key_1), C.uint64_t(key_2), C.uint64_t(key_3), C.uint64_t(val_1), C.uint64_t(val_2), C.uint64_t(val_3))
}

//
// Editor
//

func gd_editor_add_documentation(xml *byte, length uint32) {
	C.gd_editor_add_documentation((*C.char)(unsafe.Pointer(xml)), C.uint32_t(length))
}
func gd_editor_add_plugin(class_name gdStringName) { C.gd_editor_add_plugin(class_name) }
func gd_editor_end_plugin(class_name gdStringName) { C.gd_editor_end_plugin(class_name) }

//
// Iterators
//

func gd_iterator_make(v_1, v_2, v_3 uint64, result *gdVariant, err *gd_error) {
	C.gd_iterator_make(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), result, err)
}
func gd_iterator_next(v_1, v_2, v_3 uint64, iter *gdVariant, err *gd_error) bool {
	return bool(C.gd_iterator_next(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), iter, err))
}
func gd_iterator_load(v_1, v_2, v_3, i_1, i_2, i_3 uint64, result *gdVariant, err *gd_error) {
	C.gd_iterator_load(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3), C.uint64_t(i_1), C.uint64_t(i_2), C.uint64_t(i_3), result, err)
}

//
// Logging
//

func gd_log(level gd_log_level, text *byte, text_len uint32, code *byte, code_len uint32, fn *byte, fn_len uint32, file *byte, file_len uint32, line int32, notify_editor bool) {
	C.gd_log(level,
		(*C.char)(unsafe.Pointer(text)), C.uint32_t(text_len),
		(*C.char)(unsafe.Pointer(code)), C.uint32_t(code_len),
		(*C.char)(unsafe.Pointer(fn)), C.uint32_t(fn_len),
		(*C.char)(unsafe.Pointer(file)), C.uint32_t(file_len),
		C.int32_t(line), C.bool(notify_editor),
	)
}

//
// Objects
//

func gd_object_make(name gdStringName) gdObject           { return C.gd_object_make(name) }
func gd_object_name(obj gdObject) gdStringName            { return C.gd_object_name(obj) }
func gd_object_type(name gdStringName) gdClassTag         { return C.gd_object_type(name) }
func gd_object_cast(obj gdObject, to gdClassTag) gdObject { return C.gd_object_cast(obj, to) }
func gd_object_lookup(id uint64) gdObject                 { return C.gd_object_lookup(C.ObjectID(id)) }
func gd_object_global(name gdStringName) gdObject         { return C.gd_object_global(name) }
func gd_object_call(obj gdObject, method gd_method_id, result *gdVariant, arg_count int64, args *gdVariant, err *gd_error) {
	C.gd_object_call(obj, method, result, C.int64_t(arg_count), args, err)
}
func gd_object_id(obj gdObject) uint64 { return uint64(C.gd_object_id(obj)) }
func gd_object_id_inside_variant(v_1, v_2, v_3 uint64) uint64 {
	return uint64(C.gd_object_id_inside_variant(C.uint64_t(v_1), C.uint64_t(v_2), C.uint64_t(v_3)))
}
func gd_object_free(obj gdObject) { C.gd_object_free(obj) }

func gd_method(class_name, method gdStringName, hash int64) gd_method_id {
	return C.gd_method(class_name, method, C.int64_t(hash))
}
func gd_method_call(obj gdObject, fn gd_method_id, result gd_addr, shape gd_shape, args gd_addr) {
	C.gd_method_call(obj, fn, result, shape, args)
}

//
// Extension Scripts
//

func gd_script(obj gdObject, language gdObject) gd_extension_script_id {
	return C.gd_script(obj, language)
}
func gd_script_call(obj gdObject, name gdStringName, result *gdVariant, arg_count int64, args *gdVariant, err *gd_error) {
	C.gd_script_call(obj, name, result, C.int64_t(arg_count), args, err)
}
func gd_script_setup(obj gdObject, script gd_extension_script_id) {
	C.gd_script_setup(obj, script)
}
func gd_script_defines_method(obj gdObject, method gdStringName) bool {
	return bool(C.gd_script_defines_method(obj, method))
}
func gd_object_script_placeholder_create(language, script, owner gdObject) gd_extension_script_id {
	return C.gd_object_script_placeholder_create(language, script, owner)
}
func gd_object_script_placeholder_update(script gd_extension_script_id, array gdArray, dict gdDictionary) {
	C.gd_object_script_placeholder_update(script, array, dict)
}
func gd_script_yield_property(fn gd_property_iterator_t, arg uintptr, name gdStringName, state_1, state_2, state_3 uint64) {
	C.gd_script_yield_property(fn, C.uintptr_t(arg), name, C.uint64_t(state_1), C.uint64_t(state_2), C.uint64_t(state_3))
}

//
// Extension Instances
//

func gd_extension_object_setup(obj gdObject, name gdStringName, inst gd_extension_object_id) {
	C.gd_extension_object_setup(obj, name, inst)
}

//
// Extension Bindings
//

func gd_object_attach_extension_binding(obj gdObject, binding gd_extension_binding_id) {
	C.gd_object_attach_extension_binding(obj, binding)
}
func gd_object_detach_extension_binding(obj gdObject) {
	C.gd_object_detach_extension_binding(obj)
}

//
// RefCounted
//

func gd_ref_get_object(ref gdRefCounted) gdObject      { return C.gd_ref_get_object(ref) }
func gd_ref_set_object(ref gdRefCounted, obj gdObject) { C.gd_ref_set_object(ref, obj) }

//
// Strings
//

func gd_string_access(s gdString, idx int64) int32 {
	return int32(C.gd_string_access(s, C.int64_t(idx)))
}
func gd_string_memory(s gdString) gd_addr { return C.gd_string_memory(s) }
func gd_string_decode(enc gd_encoding, s *byte, length int64) gdString {
	return C.gd_string_decode(enc, (*C.char)(unsafe.Pointer(s)), C.int64_t(length))
}
func gd_string_encode(enc gd_encoding, s gdString, buf *byte, cap_ int64) int64 {
	return int64(C.gd_string_encode(enc, s, (*C.char)(unsafe.Pointer(buf)), C.int64_t(cap_)))
}
func gd_string_intern(enc gd_encoding, s *byte, length int64) gdStringName {
	return C.gd_string_intern(enc, (*C.char)(unsafe.Pointer(s)), C.int64_t(length))
}
func gd_string_resize(s gdString, size int64) gdString {
	return C.gd_string_resize(s, C.int64_t(size))
}
func gd_string_append(s gdString, other gdString) gdString { return C.gd_string_append(s, other) }
func gd_string_append_rune(s gdString, ch int32) gdString {
	return C.gd_string_append_rune(s, C.int32_t(ch))
}

//
// Callables
//

func gd_callable_create(id gd_extension_callable_t, owner uint64, result *gdCallable) {
	C.gd_callable_create(id, C.ObjectID(owner), result)
}
func gd_callable_lookup(c_1, c_2 uint64) gd_extension_callable_t {
	return C.gd_callable_lookup(C.uint64_t(c_1), C.uint64_t(c_2))
}

//
// Extension Lifecycle
//

func gd_extension_library_location() gdString { return C.gd_extension_library_location() }
