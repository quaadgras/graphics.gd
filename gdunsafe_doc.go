// Package gdunsafe provides unsafe access to the gdextension API (no protections for double-free, pointer-aliasing or use-after-free).
package gdunsafe

import (
	"unsafe"

	"graphics.gd/variant"
)

// Version of the engine.
func Version() String { return String{raw: gd_version()} }

// Major version of the engine.
func VersionMajor() uint32 { return gd_version_major() }

// Minor version of the engine.
func VersionMinor() uint32 { return gd_version_minor() }

// Patch version of the engine.
func VersionPatch() uint32 { return gd_version_patch() }

// Hexed version of the engine.
func VersionHexed() uint32 { return gd_version_hexed() }

// State of the engine (e.g. "stable", "beta", "alpha").
func VersionState() String { return String{raw: gd_version_state()} }

// Build type.
func VersionBuild() String { return String{raw: gd_version_build()} }

// Commit hash of the build.
func VersionStamp() String { return String{raw: gd_version_stamp()} }

// Timestamp of the engine build.
func VersionNanos() int64 { return gd_version_nanos() }

type MutableMemory struct {
	addr gd_addr
}

// Malloc allocates optionally 8-byte aligned memory.
func Malloc(size uintptr, pad8 bool) MutableMemory {
	return MutableMemory{addr: gd_malloc(int64(size), gdMakeBool(pad8))}
}

// Resize resizes optionally 8-byte aligned memory.
func Resize(ptr MutableMemory, size uintptr, pad8 bool) MutableMemory {
	return MutableMemory{addr: gd_resize(ptr.addr, int64(size), gdMakeBool(pad8))}
}

// Memset sets a block of memory to a given value.
func Memset(ptr MutableMemory, size uintptr, value byte) {
	gd_memset(ptr.addr, value, int64(size))
}

// LibraryPath returns a string representing the location of the current extension.
func LibraryPath() String { return String{raw: gd_extension_library_location()} }

// Tag returns the [ClassTag] of the class.
func (class Class) Tag() ClassTag { return ClassTag{raw: gd_object_type(StringName(class).raw)} }

// Free releases any resources associated with the given value of type T.
func Free[T Any](val T) {
	ptr := unsafe.Pointer(&val)
	raw := gdMakePointer(int(unsafe.Sizeof(val)), ptr)
	gd_builtin_free(uint32(variantTypeOf[T]()), raw)
	gdFreePointer(raw)
}

// Variants accessor.

//go:nosplit
func (args Variants) Index(i int) Variant {
	if i < 0 || (args.count != 0 && i >= args.count) {
		panic("index out of bounds")
	}
	return args.first.Add(uintptr(i) * unsafe.Sizeof([1]Pointer{})).Get().Get()
}

// Pointer operations.

func (ptr MutablePointer) Free() { gd_free(*(*gd_addr)(unsafe.Pointer(&ptr))) }

// String operations.

func (s String) Index(idx int) rune       { return rune(gd_string_access(s.raw, int64(idx))) }
func (s String) Resize(size int) String   { return String{raw: gd_string_resize(s.raw, int64(size))} }
func (s String) Pointer() PointerTo[rune] { return PointerTo[rune](gd_string_memory(s.raw)) }
func (s String) MutablePointer() MutablePointerTo[rune] {
	return MutablePointerTo[rune](gd_string_memory(s.raw))
}

func (s *String) Append(other String) { s.raw = gd_string_append(s.raw, other.raw) }
func (s *String) AppendRune(ch rune)  { s.raw = gd_string_append_rune(s.raw, int32(ch)) }

// Encoding operations.

func (enc latin1) Decode(s String, buf []byte) int {
	raw, free := gdMakePointerToBuffer(buf)
	defer free(buf, raw)
	return int(gd_string_encode(0, s.raw, raw, int64(len(buf))))
}
func (enc latin1) String(s string) String {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return String{raw: gd_string_decode(0, raw, int64(len(s)))}
}
func (enc latin1) Intern(s string) StringName {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return StringName{raw: gd_string_intern(0, raw, int64(len(s)))}
}

func (enc utf8) Decode(s String, buf []byte) int {
	raw, free := gdMakePointerToBuffer(buf)
	defer free(buf, raw)
	return int(gd_string_encode(1, s.raw, raw, int64(len(buf))))
}
func (enc utf8) String(s string) String {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return String{raw: gd_string_decode(1, raw, int64(len(s)))}
}
func (enc utf8) Intern(s string) StringName {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return StringName{raw: gd_string_intern(1, raw, int64(len(s)))}
}

func (enc utf16) Decode(s String, buf []byte) int {
	raw, free := gdMakePointerToBuffer(buf)
	defer free(buf, raw)
	return int(gd_string_encode(2, s.raw, raw, int64(len(buf))))
}
func (enc utf16) String(s string) String {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return String{raw: gd_string_decode(2, raw, int64(len(s)))}
}
func (enc utf16) Intern(s string) StringName {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return StringName{raw: gd_string_intern(2, raw, int64(len(s)))}
}

func (enc utf32) Decode(s String, buf []byte) int {
	raw, free := gdMakePointerToBuffer(buf)
	defer free(buf, raw)
	return int(gd_string_encode(4, s.raw, raw, int64(len(buf))))
}
func (enc utf32) String(s string) String {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return String{raw: gd_string_decode(4, raw, int64(len(s)))}
}
func (enc utf32) Intern(s string) StringName {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return StringName{raw: gd_string_intern(4, raw, int64(len(s)))}
}

func (enc wide) Decode(s String, buf []byte) int {
	raw, free := gdMakePointerToBuffer(buf)
	defer free(buf, raw)
	return int(gd_string_encode(5, s.raw, raw, int64(len(buf))))
}
func (enc wide) String(s string) String {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return String{raw: gd_string_decode(5, raw, int64(len(s)))}
}
func (enc wide) Intern(s string) StringName {
	raw, free := gdMakePointerToString(s)
	defer free(raw)
	return StringName{raw: gd_string_intern(5, raw, int64(len(s)))}
}

// Logging.

func Log(level LogLevel, text, code, fn, file string, line int32, notify_editor bool) {
	gd_log(gd_log_level(level), text, code, fn, file, line, gdMakeBool(notify_editor))
}

// Variant operations.

func Nil() Variant {
	var result Variant
	result_ptr := gdMakePointerToVariant(&result)
	gd_variant_zero(result_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	return result
}

func (v Variant) Copy(deep bool) Variant {
	var result Variant
	result_ptr := gdMakePointerToVariant(&result)
	gd_variant_copy(v[0], v[1], v[2], result_ptr, gdMakeBool(deep))
	gdFreePointerToVariant(&result, result_ptr)
	return result
}

func (v Variant) Hash(depth int64) int64 {
	return gd_variant_hash(v[0], v[1], v[2], depth)
}

func (v Variant) Bool() bool           { return gdLoadBool(gd_variant_bool(v[0], v[1], v[2])) }
func (v Variant) UnsafeString() String { return String{raw: gd_variant_text(v[0], v[1], v[2])} }
func (v Variant) Type() variant.Type   { return variant.Type(gd_variant_type(v[0], v[1], v[2])) }
func (v Variant) Free()                { gd_variant_free(v[0], v[1], v[2]) }

func (v Variant) ObjectID() ObjectID {
	return ObjectID(gd_object_id_inside_variant(v[0], v[1], v[2]))
}

func (v Variant) Call(method StringName, args ...Variant) (Variant, Error) {
	var result Variant
	var err Error
	result_ptr := gdMakePointerToVariant(&result)
	err_ptr := gdMakePointerToError(&err)
	args_ptr := gdMakePointerToVariants(args...)
	gd_variant_call(v[0], v[1], v[2], method.raw, result_ptr, int64(len(args)), args_ptr, err_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	gdFreePointerToError(&err, err_ptr)
	gdFreePointerToVariants(args_ptr)
	return result, err
}

func (v Variant) Lookup(key Variant) (Variant, bool) {
	var result Variant
	result_ptr := gdMakePointerToVariant(&result)
	ok := gd_variant_get_keyed(v[0], v[1], v[2], key[0], key[1], key[2], result_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	return result, gdLoadBool(ok)
}

func (v Variant) Index(idx int) (Variant, bool, Error) {
	var result Variant
	var err Error
	err_ptr := gdMakePointerToError(&err)
	result_ptr := gdMakePointerToVariant(&result)
	ok := gd_variant_get_index(v[0], v[1], v[2], int64(idx), result_ptr, err_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	gdFreePointerToError(&err, err_ptr)
	return result, gdLoadBool(ok), err
}

func (v Variant) Field(field StringName) (Variant, bool) {
	var result Variant
	result_ptr := gdMakePointerToVariant(&result)
	ok := gd_variant_get_field(v[0], v[1], v[2], field.raw, result_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	return result, gdLoadBool(ok)
}

func (v Variant) Has(index Variant) bool {
	return gdLoadBool(gd_variant_has_key(v[0], v[1], v[2], index[0], index[1], index[2]))
}

func (v Variant) HasMethod(method StringName) bool {
	return gdLoadBool(gd_variant_has_method(v[0], v[1], v[2], method.raw))
}

func (v Variant) Insert(key, val Variant) bool {
	return gdLoadBool(gd_variant_set_keyed(v[0], v[1], v[2], key[0], key[1], key[2], val[0], val[1], val[2]))
}

func (v Variant) SetIndex(idx int, val Variant) (bool, Error) {
	var err Error
	err_ptr := gdMakePointerToError(&err)
	ok := gd_variant_set_index(v[0], v[1], v[2], int64(idx), val[0], val[1], val[2], err_ptr)
	gdFreePointerToError(&err, err_ptr)
	return gdLoadBool(ok), err
}

func (v Variant) SetField(field StringName, value Variant) bool {
	return gdLoadBool(gd_variant_set_field(v[0], v[1], v[2], field.raw, value[0], value[1], value[2]))
}

func (op VariantOperator) Evaluate(a, b Variant) (Variant, bool) {
	var result Variant
	result_ptr := gdMakePointerToVariant(&result)
	ok := gd_variant_eval(uint32(op), a[0], a[1], a[2], b[0], b[1], b[2], result_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	return result, gdLoadBool(ok)
}

func MakeVariant(vtype variant.Type, args ...Variant) (Variant, Error) {
	var result Variant
	var err Error
	result_ptr := gdMakePointerToVariant(&result)
	err_ptr := gdMakePointerToError(&err)
	args_ptr := gdMakePointerToVariants(args...)
	gd_variant_make(uint32(vtype), result_ptr, int64(len(args)), args_ptr, err_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	gdFreePointerToError(&err, err_ptr)
	gdFreePointerToVariants(args_ptr)
	return result, err
}

func VariantInto[T Any](v Variant) T {
	var result T
	result_ptr := gdMakePointer(int(unsafe.Sizeof(result)), unsafe.Pointer(&result))
	gd_builtin_from(uint32(variantTypeOf[T]()), v[0], v[1], v[2], result_ptr)
	gdFreePointer(result_ptr)
	return result
}

func VariantFrom[T Any](native T) Variant {
	var result Variant
	native_ptr := gdMakePointer(int(unsafe.Sizeof(native)), unsafe.Pointer(&native))
	result_ptr := gdMakePointerToVariant(&result)
	gd_variant_from(uint32(variantTypeOf[T]()), result_ptr, native_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	gdFreePointer(native_ptr)
	return result
}

func PointerIntoVariant[T Any](v Variant) PointerTo[T] {
	return PointerTo[T](gd_variant_data(uint32(variantTypeOf[T]()), v[0], v[1], v[2]))
}

// Type info.

func (t Type) Size() uintptr {
	return uintptr(gd_sizeof(StringName(t.class).raw))
}

func Convertable[A, B Any](strict bool) bool {
	return gdLoadBool(gd_variant_type_convertable(uint32(variantTypeOf[A]()), uint32(variantTypeOf[B]()), gdMakeBool(strict)))
}

func PropertyExists[T Any](property StringName) bool {
	return gdLoadBool(gd_variant_type_has_property(uint32(variantTypeOf[T]()), property.raw))
}

func Constant[T Any](name StringName) Variant {
	var result Variant
	result_ptr := gdMakePointerToVariant(&result)
	gd_variant_type_constant(uint32(variantTypeOf[T]()), name.raw, result_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	return result
}

func Call[T Any](method StringName, args ...Variant) (Variant, Error) {
	var result Variant
	var err Error
	result_ptr := gdMakePointerToVariant(&result)
	err_ptr := gdMakePointerToError(&err)
	args_ptr := gdMakePointerToVariants(args...)
	gd_variant_type_call(uint32(variantTypeOf[T]()), method.raw, result_ptr, int64(len(args)), args_ptr, err_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	gdFreePointerToError(&err, err_ptr)
	gdFreePointerToVariants(args_ptr)
	return result, err
}

// Array / Dictionary / PackedArray.

func (array Array) Index(index int) Variant {
	var result Variant
	var result_ptr = gdMakePointerToVariant(&result)
	gd_array_get_index(array.raw, int64(index), result_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	return result
}

func (array Array) SetIndex(index int, value Variant) {
	gd_array_set_index(array.raw, int64(index), value[0], value[1], value[2])
}

func (array Array) SetType(t Type) {
	gd_variant_type_setup_array(array.raw, uint32(t.vtype), StringName(t.class).raw, t.script[0], t.script[1], t.script[2])
}

func (d Dictionary) Lookup(key Variant) Variant {
	var result Variant
	var result_ptr = gdMakePointerToVariant(&result)
	gd_packed_dictionary_access(d.raw, key[0], key[1], key[2], result_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	return result
}

func (d Dictionary) Insert(key, val Variant) {
	gd_packed_dictionary_modify(d.raw, key[0], key[1], key[2], val[0], val[1], val[2])
}

func (dict Dictionary) SetType(key, val Type) {
	gd_variant_type_setup_dictionary(dict.raw,
		uint32(key.vtype), StringName(key.class).raw, key.script[0], key.script[1], key.script[2],
		uint32(val.vtype), StringName(val.class).raw, val.script[0], val.script[1], val.script[2],
	)
}

func (v Variant) Iterator() (Iterator, Error) {
	var iter Iterator
	var err Error
	var iter_ptr = gdMakePointerToVariant((*Variant)(&iter))
	var err_ptr = gdMakePointerToError(&err)
	gd_iterator_make(v[0], v[1], v[2], iter_ptr, err_ptr)
	gdFreePointerToVariant((*Variant)(&iter), iter_ptr)
	gdFreePointerToError(&err, err_ptr)
	return iter, err
}

func (iter *Iterator) Next() (bool, Error) {
	var err Error
	v := Variant(*iter)
	var iter_ptr = gdMakePointerToVariant((*Variant)(iter))
	var err_ptr = gdMakePointerToError(&err)
	ok := gd_iterator_next(v[0], v[1], v[2], iter_ptr, err_ptr)
	gdFreePointerToVariant((*Variant)(iter), iter_ptr)
	gdFreePointerToError(&err, err_ptr)
	return gdLoadBool(ok), err
}

func (iter Iterator) Value() (Variant, Error) {
	var result Variant
	var err Error
	v := Variant(iter)
	var result_ptr = gdMakePointerToVariant(&result)
	var err_ptr = gdMakePointerToError(&err)
	gd_iterator_load(v[0], v[1], v[2], v[0], v[1], v[2], result_ptr, err_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	gdFreePointerToError(&err, err_ptr)
	return result, err
}

// Object operations.

func New(name Class) Object                { return Object{raw: gd_object_make(StringName(name).raw)} }
func (obj Object) Class() Class            { return Class(StringName{raw: gd_object_name(obj.raw)}) }
func (obj Object) Cast(to ClassTag) Object { return Object{raw: gd_object_cast(obj.raw, to.raw)} }
func (obj Object) ID() ObjectID            { return ObjectID(gd_object_id(obj.raw)) }
func (obj Object) Free()                   { gd_object_free(obj.raw) }
func (id ObjectID) Object() Object         { return Object{raw: gd_object_lookup(uint64(id))} }
func Singleton(name StringName) Object     { return Object{raw: gd_object_global(name.raw)} }

func Method(class, method StringName, hash int64) MethodPointer {
	return MethodPointer(gd_method(class.raw, method.raw, hash))
}

func (obj Object) Call(method MethodPointer, args ...Variant) (Variant, Error) {
	var result Variant
	var err Error
	var result_ptr = gdMakePointerToVariant(&result)
	var err_ptr = gdMakePointerToError(&err)
	args_ptr := gdMakePointerToVariants(args...)
	gd_object_call(obj.raw, gd_method_id(method), result_ptr, int64(len(args)), args_ptr, err_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	gdFreePointerToError(&err, err_ptr)
	gdFreePointerToVariants(args_ptr)
	return result, err
}

func (obj Object) ShapedCall(method MethodPointer, result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	args_ptr := gdMakePointer(shape.SizeArguments(), args)
	result_ptr := gdMakePointer(shape.SizeResult(), result)
	gd_method_call(obj.raw, gd_method_id(method), result_ptr, gd_shape(shape), args_ptr)
	gdFreePointer(args_ptr)
	gdFreePointer(result_ptr)
}

// Script operations.

func (obj Script) Call(name StringName, args ...Variant) (Variant, Error) {
	var result Variant
	var err Error
	var args_ptr = gdMakePointerToVariants(args...)
	var result_ptr = gdMakePointerToVariant(&result)
	var err_ptr = gdMakePointerToError(&err)
	gd_script_call(obj.raw, name.raw, result_ptr, int64(len(args)), args_ptr, err_ptr)
	gdFreePointerToVariant(&result, result_ptr)
	gdFreePointerToError(&err, err_ptr)
	gdFreePointerToVariants(args_ptr)
	return result, err
}

func (obj Script) HasMethod(method StringName) bool {
	return gdLoadBool(gd_script_defines_method(obj.raw, method.raw))
}

func (s Script) UpdatePlaceholder(array Array, dict Dictionary) {
	gd_object_script_placeholder_update(*(*gd_extension_script_id)(unsafe.Pointer(&s.raw)), array.raw, dict.raw)
}

// Callable.

func MakeCallable(impl ExtensionCallable, obj ObjectID) Callable {
	handle := callables.New(impl)
	var result Callable
	var result_ptr = gdMakePointerToCallable(&result)
	gd_callable_create(gd_extension_callable_t(handle), uint64(obj), result_ptr)
	gdFreePointerToCallable(&result, result_ptr)
	return result
}

// RefCounted.

func (ref RefCounted) Get() Object {
	return Object{raw: gd_ref_get_object(*(*gdRefCounted)(unsafe.Pointer(&ref.raw)))}
}
func (ref RefCounted) Set(obj Object) {
	gd_ref_set_object(*(*gdRefCounted)(unsafe.Pointer(&ref.raw)), obj.raw)
}

// Editor.

func EditorDocumentation(xml string) {
	b := []byte(xml)
	gd_editor_add_documentation(&b[0], uint32(len(b)))
}

func EnableEditorPlugin(name Class) { gd_editor_add_plugin(StringName(name).raw) }
func RemoveEditorPlugin(name Class) { gd_editor_end_plugin(StringName(name).raw) }

// Property/Method lists.

func MakePropertyList(n int64) PropertyList { return PropertyList(gd_property_list_make(n)) }
func (p PropertyList) Push(prop Property) {
	gd_property_list_push(gd_property_list_t(p), uint32(prop.Type), prop.Name.raw, prop.ClassName.raw, prop.Hint, prop.HintString.raw, prop.Usage, 0)
}
func (p PropertyList) Free() { gd_property_list_free(gd_property_list_t(p)) }

// Class registration.

func (class Class) RegisterMethods(methods MethodList) {
	gd_classdb_register_methods(StringName(class).raw, gd_method_list_t(methods))
}

func (class Class) RegisterConstant(enum, name StringName, value int64, bitfield bool) {
	gd_classdb_register_constant(StringName(class).raw, enum.raw, name.raw, value, gdMakeBool(bitfield))
}

func (class Class) RegisterProperty(property Property, setter, getter StringName) {
	pl := MakePropertyList(1)
	pl.Push(property)
	gd_classdb_register_property(StringName(class).raw, gd_property_list_t(pl), setter.raw, getter.raw)
}

func (class Class) RegisterPropertyIndexed(property Property, setter, getter StringName, index int) {
	pl := MakePropertyList(1)
	pl.Push(property)
	gd_classdb_register_property_indexed(StringName(class).raw, gd_property_list_t(pl), setter.raw, getter.raw, int64(index))
}

func (class Class) RegisterPropertyGroup(group, prefix String) {
	gd_classdb_register_property_group(StringName(class).raw, group.raw, prefix.raw)
}

func (class Class) RegisterPropertySubgroup(subgroup, prefix String) {
	gd_classdb_register_property_sub_group(StringName(class).raw, subgroup.raw, prefix.raw)
}

func (class Class) RegisterSignal(signal StringName, args PropertyList) {
	gd_classdb_register_signal(StringName(class).raw, signal.raw, gd_property_list_t(args))
}

func (class Class) Free() {
	gd_classdb_register_removal(StringName(class).raw)
}

// Method lists.

func MakeMethodList(n int64) MethodList { return MethodList(gd_method_list_make(n)) }
func (m MethodList) Push(name StringName, call ExtensionFunction, flags uint32, returnInfo PropertyList, argsInfo PropertyList, defaults []Variant) {
	fn := functions.New(call)
	var defaults_ptr gd_addr
	if len(defaults) > 0 {
		defaults_ptr = gd_malloc(int64(unsafe.Sizeof(Variant{})*uintptr(len(defaults))), gdMakeBool(false))
	}
	for i, v := range defaults {
		MutablePointerTo[Variant](PointerTo[Variant](defaults_ptr).Add(unsafe.Sizeof(Variant{}) * uintptr(i))).Set(v)
	}
	gd_method_list_push(gd_method_list_t(m), name.raw, gd_extension_method_id(fn), flags, gd_property_list_t(returnInfo), gd_property_list_t(argsInfo), int64(len(defaults)), defaults_ptr)
}
func (m MethodList) Free() { gd_method_list_free(gd_method_list_t(m)) }

// Class registration.

func RegisterClass(class string, id ExtensionClass) Class {
	handle := classes.New(id)
	name := UTF8.Intern(class)
	parent := id.Parent()
	gd_classdb_register(StringName(name).raw, StringName(parent).raw, gd_extension_class_id(handle),
		gdMakeBool(id.Virtual()), gdMakeBool(id.Abstract()), gdMakeBool(id.Exposed()), gdMakeBool(id.Runtime()), id.Icon().raw)
	return Class(name)
}

// Utility functions.

func Utility(utility StringName, hash int64) func(result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	fn := gd_function(utility.raw, hash)
	return func(result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
		result_ptr := gdMakePointer(shape.SizeResult(), result)
		args_ptr := gdMakePointer(shape.SizeArguments(), args)
		gd_call(fn, result_ptr, gd_shape(shape), args_ptr)
		gdFreePointer(result_ptr)
		gdFreePointer(args_ptr)
	}
}

// Constructor returns a constructor function for type T.
func Constructor[T Any](n int) func(shape Shape, args unsafe.Pointer) T {
	id := gd_constructor(uint32(variantTypeOf[T]()), int64(n))
	return func(shape Shape, args unsafe.Pointer) T {
		var result T
		var result_ptr = gdMakePointer(shape.SizeResult(), unsafe.Pointer(&result))
		var args_ptr = gdMakePointer(shape.SizeArguments(), args)
		gd_builtin_make(id, result_ptr, gd_shape(shape), args_ptr)
		gdFreePointer(result_ptr)
		gdFreePointer(args_ptr)
		return result
	}
}

// Evaluator returns an evaluator function for the given operator.
func Evaluator[A, B, R Any](op VariantOperator) func(a A, b B) R {
	id := gd_evaluator(uint32(op), uint32(variantTypeOf[A]()), uint32(variantTypeOf[B]()))
	shape := shapeFor[R]() | (shapeFor[A]() << 4) | (shapeFor[B]() << 8)
	return func(a A, b B) R {
		type pair struct {
			a A
			b B
		}
		args := pair{a, b}
		var result R
		var result_ptr = gdMakePointer(shape.SizeResult(), unsafe.Pointer(&result))
		var args_ptr = gdMakePointer(shape.SizeArguments(), unsafe.Pointer(&args))
		gd_builtin_eval(id, result_ptr, gd_shape(shape), args_ptr)
		gdFreePointer(result_ptr)
		gdFreePointer(args_ptr)
		return result
	}
}

// Setter returns a setter function for the given field on type T.
func Setter[T Any, E Any](field StringName) func(v T, val E) {
	id := gd_setter(uint32(variantTypeOf[T]()), field.raw)
	shape := shapeFor[T]() | (shapeFor[E]() << 4)
	return func(v T, val E) {
		type pair struct {
			e E
			t T
		}
		args := pair{val, v}
		var args_ptr = gdMakePointer(shape.SizeArguments(), unsafe.Pointer(&args))
		gd_builtin_set_field(id, gd_shape(shape), args_ptr)
		gdFreePointer(args_ptr)
	}
}

// Getter returns a getter function for the given field on type T.
func Getter[T Any, E Any](field StringName) func(v T) E {
	id := gd_getter(uint32(variantTypeOf[T]()), field.raw)
	shape := shapeFor[E]() | (shapeFor[T]() << 4)
	return func(v T) E {
		var result E
		var v_ptr = gdMakePointer(shape.SizeResult(), unsafe.Pointer(&v))
		var result_ptr = gdMakePointer(shape.SizeResult(), unsafe.Pointer(&result))
		gd_builtin_get_field(id, gd_addr(v_ptr), gd_shape(shape), gd_addr(result_ptr))
		gdFreePointer(v_ptr)
		gdFreePointer(result_ptr)
		return result
	}
}

// BuiltinMethod.Call calls the given builtin method.
func (builtin BuiltinMethod[T, Args, Result]) Call(self T, args Args) Result {
	var result Result
	type frame struct {
		self T
		args Args
	}
	f := frame{self, args}
	var f_ptr = gdMakePointer(builtin.shape.SizeArguments(), unsafe.Pointer(&f))
	var result_ptr = gdMakePointer(builtin.shape.SizeResult(), unsafe.Pointer(&result))
	gd_builtin_call(f_ptr, gd_caller_id(builtin.entry), result_ptr, gd_shape(builtin.shape), f_ptr)
	gdFreePointer(f_ptr)
	gdFreePointer(result_ptr)
	return result
}

// BuiltinMethodMutable.Call calls the given builtin method on a mutable receiver.
func (builtin BuiltinMethodMutable[T, Args, Result]) Call(self *T, check unsafe.Pointer, args Args) Result {
	if check != unsafe.Pointer(self) {
		panic("unsafe pointer mismatch")
	}
	var result Result
	var self_ptr = gdMakePointer(int(unsafe.Sizeof(*self)), unsafe.Pointer(self))
	var args_ptr = gdMakePointer(builtin.shape.SizeArguments(), unsafe.Pointer(&args))
	var result_ptr = gdMakePointer(builtin.shape.SizeResult(), unsafe.Pointer(&result))
	gd_builtin_call(self_ptr, gd_caller_id(builtin.entry), result_ptr, gd_shape(builtin.shape), args_ptr)
	gdFreePointer(args_ptr)
	gdFreePointer(result_ptr)
	gdFreePointer(self_ptr)
	return result
}

// Index/SetIndex/Insert/Lookup for builtin types.

func SetIndex[T, V Any](self T, index int64, value V) {
	type args struct {
		self  T
		value V
	}
	a := args{self, value}
	shape := shapeFor[T]() | (shapeFor[V]() << 4)
	args_ptr := gdMakePointer(shape.SizeArguments(), unsafe.Pointer(&a))
	gd_builtin_set_array(uint32(variantTypeOf[T]()), index, gd_shape(shape), args_ptr)
	gdFreePointer(args_ptr)
}

func Index[T, V Any](self T, index int64) V {
	var result V
	shape := shapeFor[V]() | (shapeFor[T]() << 4)
	args_ptr := gdMakePointer(shape.SizeArguments(), unsafe.Pointer(&self))
	result_ptr := gdMakePointer(shape.SizeResult(), unsafe.Pointer(&result))
	gd_builtin_get_array(uint32(variantTypeOf[T]()), index, result_ptr, gd_shape(shape), args_ptr)
	gdFreePointer(args_ptr)
	gdFreePointer(result_ptr)
	return result
}

func Insert[T Any](self T, key, value Variant) {
	type args struct {
		self  T
		key   Variant
		value Variant
	}
	a := args{self, key, value}
	shape := shapeFor[T]() | (shapeFor[Variant]() << 4) | (shapeFor[Variant]() << 8)
	args_ptr := gdMakePointer(shape.SizeArguments(), unsafe.Pointer(&a))
	gd_builtin_set_keyed(uint32(variantTypeOf[T]()), gd_shape(shape), args_ptr)
	gdFreePointer(args_ptr)
}

func Lookup[T Any](self T, key Variant) Variant {
	var result Variant
	shape := shapeFor[Variant]() | (shapeFor[T]() << 4) | (shapeFor[Variant]() << 8)
	args_ptr := gdMakePointer(shape.SizeArguments(), unsafe.Pointer(&self))
	result_ptr := gdMakePointer(shape.SizeArguments(), unsafe.Pointer(&result))
	gd_builtin_get_keyed(uint32(variantTypeOf[T]()), result_ptr, gd_shape(shape), args_ptr)
	gdFreePointer(args_ptr)
	gdFreePointer(result_ptr)
	return result
}

// PackedArray operations.

func (p PackedArray[T]) Index(idx int64) T {
	return PointerTo[T](gd_packed_array_access(uint32(p.Type()), p[0], p[1], idx)).Get()
}

func (p PackedArray[T]) SetIndex(idx int64, val T) {
	MutablePointerTo[T](gd_packed_array_modify(uint32(p.Type()), p[0], p[1], idx)).Set(val)
}

func (p PackedArray[T]) Pointer() PointerTo[T] {
	return PointerTo[T](gd_packed_array_access(uint32(p.Type()), p[0], p[1], 0))
}

func (p PackedArray[T]) MutablePointer() MutablePointerTo[T] {
	return MutablePointerTo[T](gd_packed_array_modify(uint32(p.Type()), p[0], p[1], 0))
}

// Pointer dereference operations.

func (ptr PointerTo[T]) Add(offset uintptr) PointerTo[T] {
	return PointerTo[T](ptr + PointerTo[T](offset))
}

// String.SetIndex sets the rune at the given index (requires mutable pointer).
func (s String) SetIndex(idx int, char rune) {
	MutablePointerTo[rune](gd_string_memory(s.raw)).Set(char)
}

// Object extension/script operations.

func (obj Object) SetupExtension(name StringName, inst ExtensionInstance) {
	handle := instances.New(inst)
	gd_extension_object_setup(obj.raw, name.raw, gd_extension_object_id(handle))
}

func (obj Object) ExtensionInstance() ExtensionInstance {
	handle := gd_object_lookup_extension_binding(obj.raw)
	if handle == 0 {
		return nil
	}
	return instances.Get(uintptr(handle))
}

func (obj Object) Script(lang ScriptLanguage) Script {
	id := gd_script(obj.raw, lang.raw)
	return Script{raw: *(*gdObject)(unsafe.Pointer(&id))}
}

func (obj Object) AttachScript(script Script) {
	gd_script_setup(obj.raw, *(*gd_extension_script_id)(unsafe.Pointer(&script.raw)))
}

func MakeScript(fn ExtensionScript) Script {
	handle := instances.New(fn)
	s := gd_script_make(gd_extension_script_id(handle))
	return Script{raw: *(*gdObject)(unsafe.Pointer(&s))}
}

func MakePlaceholderScript(language ScriptLanguage, script Script, owner Object) Script {
	id := gd_object_script_placeholder_create(language.raw, Object{raw: script.raw}.raw, owner.raw)
	return Script{raw: *(*gdObject)(unsafe.Pointer(&id))}
}
