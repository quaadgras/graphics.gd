//go:build wasm

package gdunsafe

import (
	"unsafe"
)

func (ptr PointerTo[T]) Get() T {
	var result T
	copyintoGo(unsafe.Slice((*byte)(unsafe.Pointer(&result)), int64(unsafe.Sizeof([1]T{}))), gdPointerToBuffer(ptr))
	return result
}
func (ptr MutablePointerTo[T]) Set(v T) {
	copyIntoEngine(unsafe.Slice((*byte)(unsafe.Pointer(&v)), int64(unsafe.Sizeof([1]T{}))), gdPointerToBuffer(ptr))
}

type (
	gd_addr  = uint32
	gd_shape = uint64
	gd_float = float32
	gd_error = struct {
		error    uint32
		argument int32
		expected int32
	}

	gdVariant    = [3]uint64
	gdString     = uint32
	gdStringName = uint32
	gdNodePath   = uint32
	gdObject     = uint32
	gdArray      = uint32
	gdDictionary = uint32
	gdCallable   = [2]uint64
	gdSignal     = [2]uint64
	gdRefCounted = uint32
	gdScript     = uint32
	gdClassTag   = uint32

	gd_initialization_level = uint32
	gd_constructor_id       = uint32
	gd_evaluator_id         = uint32
	gd_method_id            = uint32
	gd_getter_id            = uint32
	gd_setter_id            = uint32
	gd_caller_id            = uint32
	gd_function_t           = uint32
	gd_extension_object_id  = uint32
	gd_extension_method_id  = uint32
	gd_extension_class_id   = uint32
	gd_extension_task_id    = uint32
	gd_extension_script_id  = uint32
	gd_extension_binding_id = uint32
	gd_extension_callable_t = uint32
	gd_property_list_t      = uint32
	gd_method_list_t        = uint32
	gd_property_iterator_t  = uint32
	gd_encoding             = uint32
	gd_log_level            = uint32

	gd_bool = uint32

	gdPointerToVariant  = uint32
	gdPointerToBuffer   = uint32
	gdPointerToString   = uint32
	gdPointerToError    = uint32
	gdPointerToCallable = uint32
	gdPointerToVariants = uint32
)

func gdMakeBool(b bool) gd_bool {
	if b {
		return 1
	}
	return 0
}

func gdLoadBool(b gd_bool) bool {
	if b != 0 {
		return true
	}
	return false
}

func gdMakePointer(size int, ptr unsafe.Pointer) gd_addr {
	return gd_malloc(int64(size), 0)
}

func gdMakePointerToVariant(v *Variant) gdPointerToVariant {
	return gd_malloc(int64(unsafe.Sizeof(*v)), 0)
}
func gdMakePointerToError(e *Error) gdPointerToError {
	return gd_malloc(int64(unsafe.Sizeof(*e)), 0)
}
func gdMakePointerToCallable(c *Callable) gdPointerToCallable {
	return gd_malloc(int64(unsafe.Sizeof(*c)), 0)
}
func gdMakePointerToBuffer(buf []byte) (gdPointerToBuffer, func([]byte, gdPointerToBuffer)) {
	raw := gd_malloc(int64(len(buf)), 0)
	return gdPointerToBuffer(raw), copyintoGo
}
func gdMakePointerToVariants(args ...Variant) gdPointerToVariants {
	raw := gd_malloc(int64(unsafe.Sizeof(Variant{})*uintptr(len(args))), 0)
	copyIntoEngine(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(args))), unsafe.Sizeof(Variant{})*uintptr(len(args))), gdPointerToBuffer(raw))
	return raw
}

func gdFreePointerToVariant(ptr *Variant, raw gdPointerToVariant) {
	copyintoGo(unsafe.Slice((*byte)(unsafe.Pointer(ptr)), unsafe.Sizeof(Variant{})), gdPointerToBuffer(raw))
	gd_free(raw)
}
func gdFreePointerToError(ptr *Error, raw gdPointerToError) {
	copyintoGo(unsafe.Slice((*byte)(unsafe.Pointer(ptr)), unsafe.Sizeof(Error{})), gdPointerToBuffer(raw))
	gd_free(raw)
}
func gdFreePointerToCallable(ptr *Callable, raw gdPointerToCallable) {
	copyintoGo(unsafe.Slice((*byte)(unsafe.Pointer(ptr)), unsafe.Sizeof(Callable{})), gdPointerToBuffer(raw))
	gd_free(raw)
}
func gdFreePointerToVariants(raw gdPointerToVariants) {
	gd_free(raw)
}
func gdFreePointer(raw gd_addr) {
	gd_free(raw)
}

func copyintoPointer(ptr unsafe.Pointer, size int, raw gd_addr) {
	copyintoGo(unsafe.Slice((*byte)(ptr), size), gdPointerToBuffer(raw))
	gd_free(raw)
}

func copyintoGo(buf []byte, raw gdPointerToBuffer) {
	ptr := gd_addr(raw)
	if ptr == [1]gd_addr{}[0] {
		panic("nil pointer dereference")
	}
	var off gd_addr
	for len(buf) > 0 {
		switch {
		case len(buf) >= 4:
			*(*uint32)(unsafe.Pointer(&buf[0])) = gd_memory_bytes4(ptr + off)
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(gd_memory_bytes2(ptr + off))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			*(*uint8)(unsafe.Pointer(&buf[0])) = uint8(gd_memory_bytes1(ptr + off))
			buf = buf[1:]
			off += 1
		}
	}
}

func copyIntoEngine(buf []byte, raw gdPointerToBuffer) {
	ptr := gd_addr(raw)
	if ptr == 0 {
		panic("nil pointer dereference")
	}
	var off gd_addr
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			gd_store_bytes8(ptr+off, *(*uint64)(unsafe.Pointer(unsafe.SliceData(buf))))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			gd_store_bytes4(ptr+off, *(*uint32)(unsafe.Pointer(unsafe.SliceData(buf))))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			gd_store_bytes2(ptr+off, uint32(*(*uint16)(unsafe.Pointer(unsafe.SliceData(buf)))))
			buf = buf[2:]
			off += 2
		case len(buf) >= 1:
			gd_store_bytes1(ptr+off, uint32(*(*uint8)(unsafe.Pointer(unsafe.SliceData(buf)))))
			buf = buf[1:]
			off += 1
		}
	}
}

func gdMakePointerToString(buf string) (gdPointerToBuffer, func(gdPointerToBuffer)) {
	raw := gd_malloc(int64(len(buf)), 0)
	copyIntoEngine(unsafe.Slice(unsafe.StringData(buf), len(buf)), raw)
	return gdPointerToBuffer(raw), gd_free
}

//go:wasmimport gd version
func gd_version() gdString

//go:wasmimport gd version_major
func gd_version_major() uint32

//go:wasmimport gd version_minor
func gd_version_minor() uint32

//go:wasmimport gd version_patch
func gd_version_patch() uint32

//go:wasmimport gd version_hexed
func gd_version_hexed() uint32

//go:wasmimport gd version_state
func gd_version_state() uint32

//go:wasmimport gd version_build
func gd_version_build() uint32

//go:wasmimport gd version_stamp
func gd_version_stamp() uint32

//go:wasmimport gd version_nanos
func gd_version_nanos() int64

// Cross-memory helpers for transferring data between Go and engine address spaces.

//go:wasmimport gd malloc
func gd_malloc(size int64, pad8 gd_bool) gd_addr

//go:wasmimport gd gd_sizeof
func gd_sizeof(addr gdStringName) int64

//go:wasmimport gd resize
func gd_resize(addr gd_addr, size int64, pad8 gd_bool) gd_addr

//go:wasmimport gd memset
func gd_memset(addr gd_addr, value byte, size int64)

//go:wasmimport gd free
func gd_free(addr gd_addr)

//go:wasmimport gd memory_bytes1
func gd_memory_bytes1(gd_addr) uint32

//go:wasmimport gd memory_bytes2
func gd_memory_bytes2(gd_addr) uint32

//go:wasmimport gd memory_bytes4
func gd_memory_bytes4(gd_addr) uint32

//go:wasmimport gd memory_bytes8
func gd_memory_bytes8(gd_addr) uint64

//go:wasmimport gd store_bytes1
func gd_store_bytes1(gd_addr, uint32)

//go:wasmimport gd store_bytes2
func gd_store_bytes2(gd_addr, uint32)

//go:wasmimport gd store_bytes4
func gd_store_bytes4(gd_addr, uint32)

//go:wasmimport gd store_bytes8
func gd_store_bytes8(gd_addr, uint64)

//go:wasmimport gd store_pair64
func gd_store_pair64(gd_addr, uint64, uint64)

//go:wasmimport gd extension_library_location
func gd_extension_library_location() gdString

// Array

//go:wasmimport gd array_set
func gd_array_set_index(array uint32, index int64, v1, v2, v3 uint64)

//go:wasmimport gd array_get
func gd_array_get_index(array uint32, index int64, result uint32)

//go:wasmimport gd variant_type_setup_array
func gd_variant_type_setup_array(array uint32, vtype uint32, className uint32, v1, v2, v3 uint64)

// Memory

// String operations

//go:wasmimport gd string_access
func gd_string_access(s uint32, idx int64) int32

//go:wasmimport gd string_unsafe
func gd_string_memory(s uint32) uint32

//go:wasmimport gd string_resize
func gd_string_resize(s uint32, size int64) uint32

//go:wasmimport gd string_append
func gd_string_append(s uint32, other uint32) uint32

//go:wasmimport gd string_append_rune
func gd_string_append_rune(s uint32, ch int32) uint32

// Encoding

//go:wasmimport gd string_decode
func gd_string_decode(enc uint32, s uint32, length int64) uint32

//go:wasmimport gd string_encode
func gd_string_encode(enc uint32, s uint32, buf uint32, cap int64) int64

//go:wasmimport gd string_intern
func gd_string_intern(enc uint32, s uint32, length int64) uint32

//go:wasmimport gd log
func gd_log(level uint32, text string, code string, fn string, file string, line int32, notify_editor uint32)

// PackedArray

//go:wasmimport gd packed_array_access
func gd_packed_array_access(t uint32, a1, a2 uint32, idx int64) uint32

//go:wasmimport gd packed_array_modify
func gd_packed_array_modify(t uint32, a1, a2 uint32, idx int64) uint32

// Variant constructors

//go:wasmimport gd variant_make
func gd_variant_make(t uint32, result uint32, arg_count int64, args uint32, err uint32)

//go:wasmimport gd variant_type_call
func gd_variant_type_call(t uint32, method uint32, result uint32, argc int64, args uint32, err uint32)

//go:wasmimport gd variant_type_convertable
func gd_variant_type_convertable(t uint32, to uint32, strict uint32) uint32

//go:wasmimport gd function
func gd_function(utility uint32, hash int64) uint32

//go:wasmimport gd call
func gd_call(fn uint32, result uint32, shape gd_shape, args uint32)

//go:wasmimport gd variant_type_constant
func gd_variant_type_constant(vtype uint32, constant uint32, result uint32)

//go:wasmimport gd constructor
func gd_constructor(vtype uint32, n int64) uint32

//go:wasmimport gd builtin_make
func gd_builtin_make(constructor uint32, result uint32, shape gd_shape, args uint32)

//go:wasmimport gd evaluator
func gd_evaluator(op uint32, a, b uint32) uint32

//go:wasmimport gd builtin_eval
func gd_builtin_eval(fn uint32, result uint32, shape gd_shape, args uint32)

//go:wasmimport gd setter
func gd_setter(vtype uint32, property uint32) uint32

//go:wasmimport gd builtin_set_field
func gd_builtin_set_field(setter uint32, shape gd_shape, args uint32)

//go:wasmimport gd getter
func gd_getter(vtype uint32, property uint32) uint32

//go:wasmimport gd builtin_get_field
func gd_builtin_get_field(getter uint32, result uint32, shape gd_shape, args uint32)

//go:wasmimport gd variant_type_has_property
func gd_variant_type_has_property(vtype uint32, property uint32) uint32

//go:wasmimport gd builtin_method
func gd_builtin_method(vtype uint32, method uint32, hash int64) uint32

//go:wasmimport gd builtin_call
func gd_builtin_call(self uint32, fn uint32, result uint32, shape gd_shape, args uint32)

// BuiltinMethodByName returns a BuiltinMethodPointer that can be used to call the given builtin method on a value of type T.
//
//go:wasmimport gd builtin_set_array
func gd_builtin_set_array(vtype uint32, idx int64, shape gd_shape, args uint32)

//go:wasmimport gd builtin_get_array
func gd_builtin_get_array(vtype uint32, idx int64, result uint32, shape gd_shape, args uint32)

//go:wasmimport gd builtin_set_keyed
func gd_builtin_set_keyed(vtype uint32, shape gd_shape, args uint32)

//go:wasmimport gd builtin_get_keyed
func gd_builtin_get_keyed(vtype uint32, result uint32, shape gd_shape, args uint32)

//go:wasmimport gd builtin_free
func gd_builtin_free(vtype uint32, args uint32)

// Callable

//go:wasmimport gd callable_create
func gd_callable_create(id uint32, object uint64, result uint32)

// Object

//go:wasmimport gd object_type
func gd_object_type(name uint32) uint32

//go:wasmimport gd object_make
func gd_object_make(name uint32) uint32

//go:wasmimport gd object_name
func gd_object_name(obj uint32) uint32

//go:wasmimport gd object_cast
func gd_object_cast(obj uint32, to uint32) uint32

//go:wasmimport gd object_script_fetch
func gd_script(obj uint32, language uint32) uint32

//go:wasmimport gd object_script_setup
func gd_script_setup(obj uint32, script uint32)

//go:wasmimport gd object_id
func gd_object_id(obj uint32) uint64

//go:wasmimport gd object_free
func gd_object_free(obj uint32)

//go:wasmimport gd object_lookup
func gd_object_lookup(id uint64) uint32

//go:wasmimport gd object_global
func gd_object_global(name uint32) uint32

//go:wasmimport gd method
func gd_method(class, method uint32, hash int64) uint32

//go:wasmimport gd object_call
func gd_object_call(obj uint32, method uint32, result uint32, argc int64, args uint32, err uint32)

//go:wasmimport gd method_call
func gd_method_call(obj uint32, fn uint32, result uint32, shape gd_shape, args uint32)

//go:wasmimport gd extension_object_setup
func gd_extension_object_setup(obj uint32, name uint32, inst uint32)

//go:wasmimport gd object_lookup_extension_binding
func gd_object_lookup_extension_binding(obj uint32) uint32

// Script

//go:wasmimport gd script_make
func gd_script_make(fn uint32) uint32

//go:wasmimport gd script_call
func gd_script_call(obj uint32, name uint32, result uint32, argc int64, args uint32, err uint32)

//go:wasmimport gd script_defines_method
func gd_script_defines_method(obj uint32, method uint32) uint32

//go:wasmimport gd object_script_placeholder_create
func gd_object_script_placeholder_create(language, script, owner uint32) uint32

//go:wasmimport gd object_script_placeholder_update
func gd_object_script_placeholder_update(script uint32, array uint32, dict uint32)

// Variant operations

//go:wasmimport gd variant_zero
func gd_variant_zero(result uint32)

//go:wasmimport gd variant_copy
func gd_variant_copy(v1, v2, v3 uint64, result uint32, deep uint32)

//go:wasmimport gd variant_call
func gd_variant_call(v1, v2, v3 uint64, method uint32, result uint32, argc int64, args uint32, err uint32)

//go:wasmimport gd variant_hash
func gd_variant_hash(v1, v2, v3 uint64, depth int64) int64

//go:wasmimport gd variant_bool
func gd_variant_bool(v1, v2, v3 uint64) uint32

//go:wasmimport gd variant_text
func gd_variant_text(v1, v2, v3 uint64) uint32

//go:wasmimport gd variant_type
func gd_variant_type(v1, v2, v3 uint64) uint32

//go:wasmimport gd object_id_inside_variant
func gd_object_id_inside_variant(v1, v2, v3 uint64) uint64

//go:wasmimport gd variant_get_keyed
func gd_variant_get_keyed(v1, v2, v3, k1, k2, k3 uint64, result uint32) uint32

//go:wasmimport gd variant_get_index
func gd_variant_get_index(v1, v2, v3 uint64, idx int64, result uint32, err uint32) uint32

//go:wasmimport gd variant_get_field
func gd_variant_get_field(v1, v2, v3 uint64, field uint32, result uint32) uint32

//go:wasmimport gd variant_set_keyed
func gd_variant_set_keyed(v1, v2, v3, k1, k2, k3, val1, val2, val3 uint64) uint32

//go:wasmimport gd variant_set_index
func gd_variant_set_index(v1, v2, v3 uint64, idx int64, val1, val2, val3 uint64, err uint32) uint32

//go:wasmimport gd variant_set_field
func gd_variant_set_field(v1, v2, v3 uint64, field uint32, val1, val2, val3 uint64) uint32

//go:wasmimport gd variant_has_key
func gd_variant_has_key(v1, v2, v3, i1, i2, i3 uint64) uint32

//go:wasmimport gd variant_has_method
func gd_variant_has_method(v1, v2, v3 uint64, method uint32) uint32

//go:wasmimport gd variant_free
func gd_variant_free(v1, v2, v3 uint64)

//go:wasmimport gd variant_eval
func gd_variant_eval(op uint32, a1, a2, a3, b1, b2, b3 uint64, result uint32) uint32

//go:wasmimport gd builtin_from
func gd_builtin_from(vtype uint32, v1, v2, v3 uint64, result uint32)

//go:wasmimport gd variant_from
func gd_variant_from(vtype uint32, result uint32, args uint32)

//go:wasmimport gd variant_data
func gd_variant_data(vtype uint32, v1, v2, v3 uint64) uint32

// Dictionary

//go:wasmimport gd packed_dictionary_access
func gd_packed_dictionary_access(d uint32, k1, k2, k3 uint64, result uint32)

//go:wasmimport gd packed_dictionary_modify
func gd_packed_dictionary_modify(d uint32, k1, k2, k3, v1, v2, v3 uint64)

//go:wasmimport gd variant_type_setup_dictionary
func gd_variant_type_setup_dictionary(dict uint32, keyType uint32, keyClassName uint32, ks1, ks2, ks3 uint64, valType uint32, valClassName uint32, vs1, vs2, vs3 uint64)

// RefCounted

//go:wasmimport gd ref_get_object
func gd_ref_get_object(ref uint32) uint32

//go:wasmimport gd ref_set_object
func gd_ref_set_object(ref uint32, obj uint32)

// Editor

//go:wasmimport gd editor_add_documentation
func gd_editor_add_documentation(xml *byte, length uint32)

//go:wasmimport gd editor_add_plugin
func gd_editor_add_plugin(name uint32)

//go:wasmimport gd editor_end_plugin
func gd_editor_end_plugin(name uint32)

// PropertyList

//go:wasmimport gd property_list_make
func gd_property_list_make(n int64) uint32

//go:wasmimport gd property_list_push
func gd_property_list_push(list uint32, vtype uint32, name uint32, className uint32, hint uint32, hintString uint32, usage uint32, meta uint32)

//go:wasmimport gd property_list_free
func gd_property_list_free(list uint32)

// MethodList

//go:wasmimport gd method_list_make
func gd_method_list_make(n int64) uint32

//go:wasmimport gd method_list_push
func gd_method_list_push(list uint32, name uint32, call uint32, flags uint32, returnInfo uint32, argsInfo uint32, count int64, defaults uint32)

//go:wasmimport gd method_list_free
func gd_method_list_free(list uint32)

// ClassDB registration

//go:wasmimport gd classdb_register
func gd_classdb_register(class, parent uint32, id uint32, virtual, abstract, exposed, runtime, icon uint32)

//go:wasmimport gd classdb_register_methods
func gd_classdb_register_methods(class uint32, methods uint32)

//go:wasmimport gd classdb_register_constant
func gd_classdb_register_constant(class, enum, name uint32, value int64, bitfield uint32)

//go:wasmimport gd classdb_register_property
func gd_classdb_register_property(class uint32, property uint32, setter, getter uint32)

//go:wasmimport gd classdb_register_property_indexed
func gd_classdb_register_property_indexed(class uint32, property uint32, setter, getter uint32, index int64)

//go:wasmimport gd classdb_register_property_group
func gd_classdb_register_property_group(class, group, prefix uint32)

//go:wasmimport gd classdb_register_property_sub_group
func gd_classdb_register_property_sub_group(class, subgroup, prefix uint32)

//go:wasmimport gd classdb_register_signal
func gd_classdb_register_signal(class, signal uint32, args uint32)

//go:wasmimport gd classdb_register_removal
func gd_classdb_register_removal(class uint32)

// Iterator

//go:wasmimport gd iterator_make
func gd_iterator_make(v1, v2, v3 uint64, result uint32, err uint32)

//go:wasmimport gd iterator_next
func gd_iterator_next(v1, v2, v3 uint64, iter uint32, err uint32) uint32

//go:wasmimport gd iterator_load
func gd_iterator_load(v1, v2, v3, i1, i2, i3 uint64, result uint32, err uint32)

// Callable callbacks

//go:wasmexport go_on_callable_called
func go_on_callable_called(c uint32, ret uint32, argc int64, args uint32, err uint32) {
	r, e := callables.Get(uintptr(c)).Call(Variants{first: PointerTo[PointerTo[Variant]](args), count: int(argc)})
	MutablePointerTo[Variant](ret).Set(r)
	MutablePointerTo[Error](err).Set(e)
}

//go:wasmexport go_on_callable_verify
func go_on_callable_verify(c uint32) uint32 {
	if callables.Get(uintptr(c)).IsValid() {
		return 1
	}
	return 0
}

//go:wasmexport go_on_callable_delete
func go_on_callable_delete(c uint32) { callables.Del(uintptr(c)) }

//go:wasmexport go_on_callable_hashed
func go_on_callable_hashed(c uint32) uint32 {
	return callables.Get(uintptr(c)).Hash()
}

//go:wasmexport go_on_callable_equal
func go_on_callable_equal(a, b uint32) uint32 {
	if callables.Get(uintptr(a)).Compare(callables.Get(uintptr(b))) == 0 {
		return 1
	}
	return 0
}

//go:wasmexport go_on_callable_less_than
func go_on_callable_less_than(a, b uint32) uint32 {
	if callables.Get(uintptr(a)).Compare(callables.Get(uintptr(b))) < 0 {
		return 1
	}
	return 0
}

//go:wasmexport go_on_callable_string
func go_on_callable_string(c uint32, ok *uint32, out uint32) {
	s := callables.Get(uintptr(c)).UnsafeString()
	if s != (String{}) {
		*ok = 1
		gd_store_bytes4(out, uint32(s.raw))
	} else {
		*ok = 0
	}
}

//go:wasmexport go_on_callable_length
func go_on_callable_length(c uint32, ok *uint32) int64 {
	*ok = 1
	return int64(callables.Get(uintptr(c)).ArgumentCount())
}

// Extension binding callbacks

//go:wasmexport go_on_extension_binding_created
func go_on_extension_binding_created(p0 uint32) uint32 { return 0 }

//go:wasmexport go_on_extension_binding_removed
func go_on_extension_binding_removed(p0, p1 uint32) {}

//go:wasmexport go_on_extension_binding_reference
func go_on_extension_binding_reference(p0, p1 uint32) uint32 { return 0 }

// Extension class callbacks

//go:wasmexport go_on_extension_class_create
func go_on_extension_class_create(p0, p1 uint32) uint32 {
	return uint32(classes.Get(uintptr(p0)).Create(p1 != 0).raw)
}

//go:wasmexport go_on_extension_class_method
func go_on_extension_class_method(p0, p1, p2 uint32) uint32 {
	fn := classes.Get(uintptr(p0)).Method(StringName{raw: p1}, p2)
	if fn == nil {
		return 0
	}
	return uint32(functions.New(fn))
}

//go:wasmexport go_on_extension_class_caller
func go_on_extension_class_caller(p0, p1, p2 uint32) uint32 {
	fn := classes.Get(uintptr(p0)).Method(StringName{raw: p1}, p2)
	if fn == nil {
		return 0
	}
	return uint32(functions.New(fn))
}

// Extension instance callbacks

//go:wasmexport go_on_extension_instance_set
func go_on_extension_instance_set(p0, p1 uint32, p2, p3, p4 uint64) uint32 {
	if instances.Get(uintptr(p0)).Set(StringName{raw: p1}, Variant{p2, p3, p4}) {
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_instance_get
func go_on_extension_instance_get(p0, p1, p2 uint32) uint32 {
	v, ok := instances.Get(uintptr(p0)).Get(StringName{raw: p1})
	if ok {
		MutablePointerTo[Variant](p2).Set(v)
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_instance_property_list
func go_on_extension_instance_property_list(p0 uint32) uint32 {
	return uint32(instances.Get(uintptr(p0)).PropertyList())
}

//go:wasmexport go_on_extension_instance_property_has_default
func go_on_extension_instance_property_has_default(p0, p1 uint32) uint32 {
	if instances.Get(uintptr(p0)).HasDefault(StringName{raw: p1}) {
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_instance_property_get_default
func go_on_extension_instance_property_get_default(p0, p1, p2 uint32) uint32 {
	v, ok := instances.Get(uintptr(p0)).GetDefault(StringName{raw: p1})
	if ok {
		MutablePointerTo[Variant](p2).Set(v)
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_instance_property_validation
func go_on_extension_instance_property_validation(p0, p1 uint32) uint32 {
	if instances.Get(uintptr(p0)).ValidateProperty(Property{Name: StringName{raw: p1}}) {
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_instance_notification
func go_on_extension_instance_notification(p0 uint32, p1 int32, p2 uint32) {
	instances.Get(uintptr(p0)).Notification(p1, p2 != 0)
}

//go:wasmexport go_on_extension_instance_stringify
func go_on_extension_instance_stringify(p0 uint32, ok *uint32, out uint32) {
	s := instances.Get(uintptr(p0)).UnsafeString()
	if s != (String{}) {
		*ok = 1
		gd_store_bytes4(out, uint32(s.raw))
	} else {
		*ok = 0
	}
}

//go:wasmexport go_on_extension_instance_reference
func go_on_extension_instance_reference(p0 uint32) {
	instances.Get(uintptr(p0)).Reference(true)
}

//go:wasmexport go_on_extension_instance_unreference
func go_on_extension_instance_unreference(p0 uint32) uint32 {
	if instances.Get(uintptr(p0)).Reference(false) {
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_instance_rid
func go_on_extension_instance_rid(p0 uint32) uint64 {
	return uint64(instances.Get(uintptr(p0)).RID())
}

//go:wasmexport go_on_extension_instance_checked_call
func go_on_extension_instance_checked_call(p0, p1, p2, p3 uint32) {
	var inst ExtensionInstance
	if p0 != 0 {
		inst = instances.Get(uintptr(p0))
	}
	functions.Get(uintptr(p1)).PointerCall(inst, MutablePointer(p3), PointerTo[Pointer](p2))
}

//go:wasmexport go_on_extension_instance_called
func go_on_extension_instance_called(p0, p1, p2, p3 uint32) {
	inst := instances.Get(uintptr(p0))
	functions.Get(uintptr(p1)).PointerCall(inst, MutablePointer(p3), PointerTo[Pointer](p2))
}

//go:wasmexport go_on_extension_instance_variant_call
func go_on_extension_instance_variant_call(p0, p1, p2, p3 uint32) {
	var inst ExtensionInstance
	if p0 != 0 {
		inst = instances.Get(uintptr(p0))
	}
	v := functions.Get(uintptr(p1)).CheckedCall(inst, Variants{
		first: PointerTo[PointerTo[Variant]](p3),
		count: -1,
	})
	MutablePointerTo[Variant](p2).Set(v)
}

//go:wasmexport go_on_extension_instance_dynamic_call
func go_on_extension_instance_dynamic_call(p0, p1, p2 uint32, p3 int64, p4, p5 uint32) {
	var inst ExtensionInstance
	if p0 != 0 {
		inst = instances.Get(uintptr(p0))
	}
	v, err := functions.Get(uintptr(p1)).DynamicCall(inst, Variants{
		first: PointerTo[PointerTo[Variant]](p4),
		count: int(p3),
	})
	MutablePointerTo[Variant](p2).Set(v)
	MutablePointerTo[Error](p5).Set(err)
}

//go:wasmexport go_on_extension_instance_free
func go_on_extension_instance_free(p0 uint32) {
	inst := instances.Get(uintptr(p0))
	if f, ok := inst.(interface{ Free() }); ok {
		f.Free()
	}
	instances.Del(uintptr(p0))
}

// Extension script callbacks

//go:wasmimport gd script_yield_property
func gd_script_yield_property(fn uint32, arg uint32, name uint32, s1, s2, s3 uint64)

//go:wasmexport go_on_extension_script_categorization
func go_on_extension_script_categorization(p0, p1 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	if script.PropertyCategory() != (StringName{}) {
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_script_get_property_type
func go_on_extension_script_get_property_type(p0, p1, p2 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		MutablePointerTo[Error](p2).Set(Error{error: errorInvalidMethod})
		return 0
	}
	return uint32(script.PropertyType(StringName{raw: p1}))
}

//go:wasmexport go_on_extension_script_get_owner
func go_on_extension_script_get_owner(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.Owner().raw)
}

//go:wasmexport go_on_extension_script_get_property_state
func go_on_extension_script_get_property_state(p0, p1, p2 uint32) {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return
	}
	script.ExportedProperties(func(name StringName, value Variant) bool {
		gd_script_yield_property(p1, p2, uint32(name.raw), value[0], value[1], value[2])
		return true
	})
}

//go:wasmexport go_on_extension_script_get_methods
func go_on_extension_script_get_methods(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.MethodList())
}

//go:wasmexport go_on_extension_script_has_method
func go_on_extension_script_has_method(p0, p1 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	if script.HasMethod(StringName{raw: p1}) {
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_script_get_method_argument_count
func go_on_extension_script_get_method_argument_count(p0, p1 uint32) int64 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return int64(script.MethodArgumentCount(StringName{raw: p1}))
}

//go:wasmexport go_on_extension_script_get
func go_on_extension_script_get(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.Script().raw)
}

//go:wasmexport go_on_extension_script_is_placeholder
func go_on_extension_script_is_placeholder(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	if script.IsPlaceholder() {
		return 1
	}
	return 0
}

//go:wasmexport go_on_extension_script_get_language
func go_on_extension_script_get_language(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.ScriptLanguage().raw)
}

// Lifecycle callbacks

//go:wasmexport go_on_engine_init
func go_on_engine_init(p0 uint32) {
	level := InitializationLevel(p0)
	for _, fn := range onEngineInit {
		fn(level)
	}
}

//go:wasmexport go_on_engine_exit
func go_on_engine_exit(p0 uint32) {
	level := InitializationLevel(p0)
	for _, fn := range onEngineExit {
		fn(level)
	}
}

//go:wasmexport go_on_first_frame
func go_on_first_frame() {
	for _, fn := range onFirstFrame {
		fn()
	}
}

//go:wasmexport go_on_every_frame
func go_on_every_frame() {
	for _, fn := range onEveryFrame {
		fn()
	}
}

//go:wasmexport go_on_final_frame
func go_on_final_frame() {
	for _, fn := range onFinalFrame {
		fn()
	}
}
