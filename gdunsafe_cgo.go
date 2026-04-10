//go:build cgo

package gdunsafe

// #include "gd.h"
import "C"
import (
	"unsafe"

	"graphics.gd/variant"
)

type Pointer uintptr

type MutablePointer uintptr

type taskID = uintptr

var (
	onWorkerThreadPoolTask      func(taskID)
	onWorkerThreadPoolGroupTask func(taskID, int32)
	onEditorClassDetection      func(PackedArray[String]) PackedArray[String]
)

func toVariant(v C.Variant) Variant {
	return Variant{uint64(v.tag), uint64(v.payload[0]), uint64(v.payload[1])}
}
func toCallError(cerr C.CallError) Error {
	return Error{error: errorType(cerr.error), expected: int32(cerr.expected), argument: int32(cerr.argument)}
}

// LibraryLocation returns a string representing the location of the current extension.
func LibraryLocation() String { return String(C.gd_library_location()) }

func (args Variants) Index(i int) Variant {
	if args.count > 0 && (i >= args.count || i < 0) {
		panic("index out of range")
	}
	slot := unsafe.Pointer(uintptr(args.first) + unsafe.Sizeof(Pointer(0))*uintptr(i))
	return *(*Variant)(*(*unsafe.Pointer)(slot))
}

func (array Array) Index(index int) Variant {
	var r C.Variant
	C.gd_array_get(C.Array(array), C.int64_t(index), &r)
	return toVariant(r)
}

func (array Array) SetIndex(index int, value Variant) {
	C.gd_array_set(C.Array(array), C.int64_t(index), C.uint64_t(value[0]), C.uint64_t(value[1]), C.uint64_t(value[2]))
}

func (array Array) SetType(t Type) {
	var script Variant = t.script
	C.gd_variant_type_setup_array(C.uintptr_t(array), C.uint32_t(uint32(t.vtype)), C.uintptr_t(t.class),
		C.uint64_t(script[0]), C.uint64_t(script[1]), C.uint64_t(script[2]))
}

func (t Type) Size() uintptr { return uintptr(t.shape.SizeResult()) }

// Version

func Version() String          { return String(C.gd_version_string()) }
func VersionMajor() uint32     { return uint32(C.gd_version_major()) }
func VersionMinor() uint32     { return uint32(C.gd_version_minor()) }
func VersionPatch() uint32     { return uint32(C.gd_version_patch()) }
func VersionHexed() uint32     { return uint32(C.gd_version_hex()) }
func VersionState() String     { return String(C.gd_version_status()) }
func VersionBuild() String     { return String(C.gd_version_build()) }
func VersionCommit() String    { return String(C.gd_version_hash()) }
func VersionTimestamp() uint64 { return uint64(C.gd_version_timestamp()) }

// Memory

func Malloc(size uintptr) MutablePointer {
	return MutablePointer(C.gd_memory_malloc(C.int64_t(size)))
}

func Resize(ptr MutablePointer, size uintptr) MutablePointer {
	return MutablePointer(C.gd_memory_resize(C.UnsafePointer(ptr), C.int64_t(size)))
}

func Memset(ptr MutablePointer, size uintptr, value byte) {
	C.gd_memory_clear(C.UnsafePointer(ptr), C.int64_t(size))
}

func (ptr MutablePointer) Free() { C.gd_memory_free(C.UnsafePointer(ptr)) }

func (ptr PointerTo[T]) Get() T         { return *(*T)(unsafe.Pointer(ptr)) }
func (ptr MutablePointerTo[T]) Set(v T) { *(*T)(unsafe.Pointer(ptr)) = v }

// String operations

func (s String) Index(idx int) rune {
	return rune(C.gd_string_access(C.uintptr_t(s), C.int64_t(idx)))
}

func (s String) SetIndex(idx int, char rune) {
	ptr := unsafe.Pointer(C.gd_string_unsafe(C.uintptr_t(s)))
	*(*int32)(unsafe.Add(ptr, uintptr(idx)*4)) = int32(char)
}

func (s String) Resize(size int) String {
	return String(C.gd_string_resize(C.uintptr_t(s), C.int64_t(size)))
}

func (s String) Pointer() PointerTo[rune] {
	return PointerTo[rune](C.gd_string_unsafe(C.uintptr_t(s)))
}

func (s String) MutablePointer() MutablePointerTo[rune] {
	return MutablePointerTo[rune](C.gd_string_unsafe(C.uintptr_t(s)))
}

func (s *String) Append(other String) {
	*s = String(C.gd_string_append(C.uintptr_t(*s), C.uintptr_t(other)))
}

func (s *String) AppendRune(ch rune) {
	*s = String(C.gd_string_append_rune(C.uintptr_t(*s), C.int32_t(ch)))
}

// Encoding — Latin1

func (enc latin1) Decode(s String, buf []byte) int {
	return int(C.gd_string_encode(C.uint8_t(0), C.uintptr_t(s),
		(*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf))))
}

func (enc latin1) String(s string) String {
	return String(C.gd_string_decode(C.uint8_t(0),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

func (enc latin1) Intern(s string) StringName {
	return StringName(C.gd_string_intern(C.uint8_t(0),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

// Encoding — UTF8

func (enc utf8) Decode(s String, buf []byte) int {
	return int(C.gd_string_encode(C.uint8_t(1), C.uintptr_t(s),
		(*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf))))
}

func (enc utf8) String(s string) String {
	return String(C.gd_string_decode(C.uint8_t(1),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

func (enc utf8) Intern(s string) StringName {
	return StringName(C.gd_string_intern(C.uint8_t(1),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

// Encoding — UTF16

func (enc utf16) Decode(s String, buf []byte) int {
	return int(C.gd_string_encode(C.uint8_t(2), C.uintptr_t(s),
		(*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf))))
}

func (enc utf16) String(s string) String {
	return String(C.gd_string_decode(C.uint8_t(2),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

func (enc utf16) Intern(s string) StringName {
	return StringName(C.gd_string_intern(C.uint8_t(2),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

// Encoding — UTF32

func (enc utf32) Decode(s String, buf []byte) int {
	return int(C.gd_string_encode(C.uint8_t(4), C.uintptr_t(s),
		(*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf))))
}

func (enc utf32) String(s string) String {
	return String(C.gd_string_decode(C.uint8_t(4),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

func (enc utf32) Intern(s string) StringName {
	return StringName(C.gd_string_intern(C.uint8_t(4),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

// Encoding — Wide

func (enc wide) Decode(s String, buf []byte) int {
	return int(C.gd_string_encode(C.uint8_t(5), C.uintptr_t(s),
		(*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), C.int64_t(len(buf))))
}

func (enc wide) String(s string) String {
	return String(C.gd_string_decode(C.uint8_t(5),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

func (enc wide) Intern(s string) StringName {
	return StringName(C.gd_string_intern(C.uint8_t(5),
		(*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.int64_t(len(s))))
}

// Log

func Log(level LogLevel, text, code, fn, file string, line int32, notify_editor bool) {
	C.gd_log(C.uint32_t(level),
		(*C.char)(unsafe.Pointer(unsafe.StringData(text))), C.uint32_t(len(text)),
		(*C.char)(unsafe.Pointer(unsafe.StringData(code))), C.uint32_t(len(code)),
		(*C.char)(unsafe.Pointer(unsafe.StringData(fn))), C.uint32_t(len(fn)),
		(*C.char)(unsafe.Pointer(unsafe.StringData(file))), C.uint32_t(len(file)),
		C.int32_t(line), C._Bool(notify_editor))
}

// PackedArray

func (p PackedArray[T]) Index(idx int64) T {
	ptr := PointerTo[T](C.gd_packed_array_access(C.uint32_t(p.Type()), C.uintptr_t(p[0]), C.uintptr_t(p[1]), C.int64_t(idx)))
	return ptr.Get()
}

func (p PackedArray[T]) SetIndex(idx int64, val T) {
	ptr := MutablePointerTo[T](C.gd_packed_array_modify(C.uint32_t(p.Type()), C.uintptr_t(p[0]), C.uintptr_t(p[1]), C.int64_t(idx)))
	ptr.Set(val)
}

func (p PackedArray[T]) Pointer() PointerTo[T] {
	return PointerTo[T](C.gd_packed_array_access(C.uint32_t(p.Type()), C.uintptr_t(p[0]), C.uintptr_t(p[1]), 0))
}

func (p PackedArray[T]) MutablePointer() MutablePointerTo[T] {
	return MutablePointerTo[T](C.gd_packed_array_modify(C.uint32_t(p.Type()), C.uintptr_t(p[0]), C.uintptr_t(p[1]), 0))
}

// Variant constructors

func MakeVariant(vtype variant.Type, args ...Variant) (Variant, Error) {
	var value C.Variant
	var err C.CallError
	C.gd_variant_type_make(C.uint32_t(uint32(vtype)), &value, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(value), toCallError(err)
}

func Call[T Any](method StringName, args ...Variant) (Variant, Error) {
	var value C.Variant
	var err C.CallError
	C.gd_variant_type_call(C.uint32_t(uint32(variantTypeOf[T]())), C.uintptr_t(method), &value, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(value), toCallError(err)
}

func Convertable[A, B Any](strict bool) bool {
	return bool(C.gd_variant_type_convertable(C.uint32_t(uint32(variantTypeOf[A]())), C.uint32_t(uint32(variantTypeOf[B]())), C.bool(strict)))
}

func Utility(utility StringName, hash int64) func(result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	fn := C.gd_builtin_name(C.uintptr_t(utility), C.int64_t(hash))
	return func(result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
		C.gd_builtin_call(fn, C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
	}
}

func Constant[T, E Any](name StringName) E {
	var result E
	C.gd_variant_type_fetch_constant(C.uint32_t(uint32(variantTypeOf[T]())), C.uintptr_t(name), C.UnsafePointer(unsafe.Pointer(&result)))
	return result
}

func Constructor[T Any](n int) func(shape Shape, args unsafe.Pointer) T {
	fn := C.gd_variant_type_unsafe_constructor(C.uint32_t(uint32(variantTypeOf[T]())), C.int64_t(n))
	return func(shape Shape, args unsafe.Pointer) T {
		var result T
		C.gd_variant_type_unsafe_make(fn, C.UnsafePointer(unsafe.Pointer(&result)), C.uint64_t(shape), C.UnsafePointer(args))
		return result
	}
}

func Evaluator[A, B, R Any](op VariantOperator) func(a A, b B) R {
	fn := C.gd_variant_type_evaluator(C.uint32_t(op), C.uint32_t(uint32(variantTypeOf[A]())), C.uint32_t(uint32(variantTypeOf[B]())))
	shapeA := shapeOf[A]()
	shapeB := shapeOf[B]()
	shape := uint64(shapeOf[R]()) | uint64(shapeA)<<4 | uint64(shapeB)<<8
	return func(a A, b B) R {
		var result R
		C.gd_variant_unsafe_eval(fn, C.UnsafePointer(unsafe.Pointer(&result)), C.uint64_t(shape), C.UnsafePointer(unsafe.Pointer(&struct {
			A A
			B B
		}{A: a, B: b})))
		return result
	}
}

func Setter[T Any, E Any](field StringName) func(v T, val E) {
	fn := C.gd_variant_type_setter(C.uint32_t(uint32(variantTypeOf[T]())), C.uintptr_t(field))
	shapeT := shapeOf[T]()
	shapeE := shapeOf[E]()
	shape := uint64(shapeT)<<4 | uint64(shapeE)<<8
	return func(v T, val E) {
		C.gd_variant_unsafe_set_field(fn, C.uint64_t(shape), C.UnsafePointer(unsafe.Pointer(&struct {
			T T
			E E
		}{T: v, E: val})))
	}
}

func Getter[T Any, E Any](field StringName) func(v T) E {
	fn := C.gd_variant_type_getter(C.uint32_t(uint32(variantTypeOf[T]())), C.uintptr_t(field))
	shapeT := shapeOf[T]()
	shape := uint64(shapeOf[E]()) | uint64(shapeT)<<4
	return func(v T) E {
		var result E
		C.gd_variant_unsafe_get_field(fn, C.UnsafePointer(unsafe.Pointer(&result)), C.uint64_t(shape), C.UnsafePointer(unsafe.Pointer(&v)))
		return result
	}
}

func PropertyExists[T Any](property StringName) bool {
	return bool(C.gd_variant_type_has_property(C.uint32_t(uint32(variantTypeOf[T]())), C.uintptr_t(property)))
}

func BuiltinMethod[T Any](method StringName, hash int64) func(self *T, ret unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	fn := C.gd_variant_type_builtin_method(C.uint32_t(uint32(variantTypeOf[T]())), C.uintptr_t(method), C.int64_t(hash))
	return func(self *T, ret unsafe.Pointer, shape Shape, args unsafe.Pointer) {
		C.gd_variant_type_unsafe_call(C.UnsafePointer(unsafe.Pointer(self)), fn, C.UnsafePointer(ret), C.uint64_t(shape), C.UnsafePointer(args))
	}
}

func SetIndex[T, V Any](self T, index int64, value V) {
	shapeT := shapeOf[T]()
	shapeV := shapeOf[V]()
	shape := uint64(shapeT)<<4 | uint64(shapeV)<<8
	C.gd_variant_unsafe_set_array(C.uint32_t(uint32(variantTypeOf[T]())), C.int64_t(index), C.uint64_t(shape), C.UnsafePointer(unsafe.Pointer(&struct {
		T T
		V V
	}{T: self, V: value})))
}

func Index[T, V Any](self T, index int64) V {
	shapeT := shapeOf[T]()
	shape := uint64(shapeOf[V]()) | uint64(shapeT)<<4
	var result V
	C.gd_variant_unsafe_get_array(C.uint32_t(uint32(variantTypeOf[T]())), C.int64_t(index), C.UnsafePointer(unsafe.Pointer(&result)), C.uint64_t(shape), C.UnsafePointer(unsafe.Pointer(&self)))
	return result
}

func Insert[T Any](self T, index, value Variant) {
	shapeT := shapeOf[T]()
	shape := uint64(shapeT)<<4 | uint64(ShapeVariant)<<8 | uint64(ShapeVariant)<<12
	C.gd_variant_unsafe_set_index(C.uint32_t(uint32(variantTypeOf[T]())), C.uint64_t(shape), C.UnsafePointer(unsafe.Pointer(&struct {
		T     T
		Index Variant
		Value Variant
	}{T: self, Index: index, Value: value})))
}

func Lookup[T Any](self T, key Variant) Variant {
	shapeT := shapeOf[T]()
	shape := uint64(ShapeVariant) | uint64(shapeT)<<4 | uint64(ShapeVariant)<<8
	var result Variant
	C.gd_variant_unsafe_get_index(C.uint32_t(uint32(variantTypeOf[T]())), C.UnsafePointer(unsafe.Pointer(&result)), C.uint64_t(shape), C.UnsafePointer(unsafe.Pointer(&struct {
		T   T
		Key Variant
	}{
		T:   self,
		Key: key,
	})))
	return result
}

func Free[T Any](val T) {
	C.gd_variant_type_unsafe_free(C.uint32_t(uint32(variantTypeOf[T]())), C.uint64_t(uint64(shapeOf[T]())<<4), C.UnsafePointer(unsafe.Pointer(&val)))
}

// Callable

func MakeCallable(impl ExtensionCallable, obj ObjectID) Callable {
	var c C.Callable
	C.gd_callable_create(C.CallableID(callables.New(impl)), C.ObjectID(obj), &c)
	return Callable{uint64(c.opaque[0]), uint64(c.opaque[1])}
}

// Object

func (class Class) Tag() ClassTag {
	return ClassTag(C.gd_object_type(C.StringName(class)))
}

func New(name Class) Object {
	return Object(C.gd_object_make(C.StringName(name)))
}

func (obj Object) Class() Class {
	return Class(C.gd_object_name(C.Object(obj)))
}

func (obj Object) Cast(to ClassTag) Object {
	return Object(C.gd_object_cast(C.Object(obj), C.ObjectType(to)))
}

func (obj Object) Script(lang ScriptLanguage) Script {
	return Script(Object(C.gd_object_script_fetch(C.Object(obj), C.Object(Object(lang)))))
}

func (obj Object) AttachScript(script Script) {
	C.gd_object_script_setup(C.Object(obj), C.ScriptInstance(Object(script)))
}

func (obj Object) ID() ObjectID {
	return ObjectID(C.gd_object_id(C.Object(obj)))
}

func (obj Object) Free() {
	C.gd_object_unsafe_free(C.Object(obj))
}

func (id ObjectID) Object() Object {
	return Object(C.gd_object_lookup(C.ObjectID(id)))
}

func Singleton(name StringName) Object {
	return Object(C.gd_object_global(C.StringName(name)))
}

func Method(class, method StringName, hash int64) MethodPointer {
	return MethodPointer(C.gd_object_method_lookup(C.StringName(class), C.StringName(method), C.int64_t(hash)))
}

func (obj Object) Call(method MethodPointer, args ...Variant) (Variant, Error) {
	var ret C.Variant
	var err C.CallError
	C.gd_object_call(C.Object(obj), C.MethodForClass(method), &ret, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(ret), toCallError(err)
}

func (obj Object) ShapedCall(method MethodPointer, result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	C.gd_object_shaped_call(C.Object(obj), C.MethodForClass(method), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}

func (obj Object) SetupExtension(name StringName, inst ExtensionInstance) {
	C.gd_object_extension_setup(C.Object(obj), C.StringName(name), C.ExtensionInstanceID(instances.New(inst)))
}

func (obj Object) ExtensionInstance() ExtensionInstance {
	return instances.Get(uintptr(C.gd_object_extension_fetch(C.Object(obj))))
}

// Script

func MakeScript(fn ExtensionScript) Script {
	return Script(Object(C.gd_object_script_make(C.ExtensionInstanceID(instances.New(fn)))))
}

func (obj Script) Call(name StringName, args ...Variant) (Variant, Error) {
	var ret C.Variant
	var err C.CallError
	C.gd_object_script_call(C.Object(Object(obj)), C.StringName(name), &ret, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(ret), toCallError(err)
}

func (obj Script) HasMethod(method StringName) bool {
	return bool(C.gd_object_script_defines_method(C.Object(Object(obj)), C.StringName(method)))
}

func MakePlaceholderScript(language ScriptLanguage, script Script, owner Object) Script {
	return Script(Object(C.gd_object_script_placeholder_create(C.Object(Object(language)), C.Object(Object(script)), C.Object(owner))))
}

func (s Script) UpdatePlaceholder(array Array, dict Dictionary) {
	C.gd_object_script_placeholder_update(C.ScriptInstance(Object(s)), C.Array(array), C.Dictionary(dict))
}

// Variant operations

func Nil() Variant {
	var zero C.Variant
	C.gd_variant_zero(&zero)
	return toVariant(zero)
}

func (v Variant) Copy(deep bool) Variant {
	var result C.Variant
	if deep {
		C.gd_variant_deep_copy(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), &result)
	} else {
		C.gd_variant_copy(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), &result)
	}
	return toVariant(result)
}

func (v Variant) Call(method StringName, args ...Variant) (Variant, Error) {
	var result C.Variant
	var err C.CallError
	C.gd_variant_call(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.StringName(method), &result, C.int64_t(len(args)), (*C.Variant)(unsafe.Pointer(unsafe.SliceData(args))), &err)
	return toVariant(result), toCallError(err)
}

func (v Variant) Hash(depth int64) int64 {
	return int64(C.gd_variant_deep_hash(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.int64_t(depth)))
}

func (v Variant) Bool() bool {
	return bool(C.gd_variant_bool(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}

func (v Variant) UnsafeString() String {
	return String(C.gd_variant_text(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}

func (v Variant) Type() variant.Type {
	return variant.Type(C.gd_variant_type(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}

func (v Variant) ObjectID() ObjectID {
	return ObjectID(C.gd_object_id_inside_variant(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}

func (v Variant) Lookup(key Variant) (Variant, bool) {
	var result C.Variant
	ok := bool(C.gd_variant_get_index(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(key[0]), C.uint64_t(key[1]), C.uint64_t(key[2]), &result))
	return toVariant(result), ok
}

func (v Variant) Index(idx int) (Variant, bool, Error) {
	var result C.Variant
	var err C.CallError
	ok := bool(C.gd_variant_get_array(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.int64_t(idx), &result, &err))
	return toVariant(result), ok, toCallError(err)
}

func (v Variant) Field(field StringName) (Variant, bool) {
	var result C.Variant
	ok := bool(C.gd_variant_get_field(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.StringName(field), &result))
	return toVariant(result), ok
}

func (v Variant) Insert(key, val Variant) bool {
	return bool(C.gd_variant_set_index(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(key[0]), C.uint64_t(key[1]), C.uint64_t(key[2]), C.uint64_t(val[0]), C.uint64_t(val[1]), C.uint64_t(val[2])))
}

func (v Variant) SetIndex(idx int, val Variant) (bool, Error) {
	var err Error
	ok := bool(C.gd_variant_set_array(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.int64_t(idx), C.uint64_t(val[0]), C.uint64_t(val[1]), C.uint64_t(val[2]), C.UnsafePointer(unsafe.Pointer(&err))))
	return ok, err
}

func (v Variant) SetField(field StringName, value Variant) bool {
	return bool(C.gd_variant_set_field(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.StringName(field), C.uint64_t(value[0]), C.uint64_t(value[1]), C.uint64_t(value[2])))
}

func (v Variant) Has(index Variant) bool {
	return bool(C.gd_variant_has_index(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(index[0]), C.uint64_t(index[1]), C.uint64_t(index[2])))
}

func (v Variant) HasMethod(method StringName) bool {
	return bool(C.gd_variant_has_method(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.StringName(method)))
}

func (v Variant) Free() {
	C.gd_variant_unsafe_free(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]))
}

func (op VariantOperator) Evaluate(a, b Variant) (Variant, bool) {
	var result C.Variant
	ok := bool(C.gd_variant_eval(C.uint32_t(op), C.uint64_t(a[0]), C.uint64_t(a[1]), C.uint64_t(a[2]), C.uint64_t(b[0]), C.uint64_t(b[1]), C.uint64_t(b[2]), &result))
	return toVariant(result), ok
}

func VariantInto[T Any](v Variant) T {
	var result T
	C.gd_variant_unsafe_make_native(C.uint32_t(uint32(variantTypeOf[T]())), C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]), C.uint64_t(uint64(shapeOf[T]())<<4), C.UnsafePointer(unsafe.Pointer(&result)))
	return result
}

func VariantFrom[T Any](native T) Variant {
	var result C.Variant
	C.gd_variant_unsafe_from_native(C.uint32_t(uint32(variantTypeOf[T]())), &result, C.uint64_t(uint64(shapeOf[T]())<<4), C.UnsafePointer(unsafe.Pointer(&native)))
	return toVariant(result)
}

func PointerIntoVariant[T Any](v Variant) PointerTo[T] {
	return PointerTo[T](C.gd_variant_unsafe_internal_pointer(C.uint32_t(uint32(variantTypeOf[T]())), C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}

// Dictionary

func (d Dictionary) Lookup(key Variant) Variant {
	var result C.Variant
	C.gd_packed_dictionary_access(C.uintptr_t(d), C.uint64_t(key[0]), C.uint64_t(key[1]), C.uint64_t(key[2]), &result)
	return toVariant(result)
}

func (d Dictionary) Insert(key, val Variant) {
	C.gd_packed_dictionary_modify(C.uintptr_t(d), C.uint64_t(key[0]), C.uint64_t(key[1]), C.uint64_t(key[2]), C.uint64_t(val[0]), C.uint64_t(val[1]), C.uint64_t(val[2]))
}

func (dict Dictionary) SetType(key, val Type) {
	var keyScript, valScript = key.script, val.script
	C.gd_variant_type_setup_dictionary(C.uintptr_t(dict),
		C.uint32_t(uint32(key.vtype)), C.uintptr_t(key.class),
		C.uint64_t(keyScript[0]), C.uint64_t(keyScript[1]), C.uint64_t(keyScript[2]),
		C.uint32_t(uint32(val.vtype)), C.uintptr_t(val.class),
		C.uint64_t(valScript[0]), C.uint64_t(valScript[1]), C.uint64_t(valScript[2]))
}

// RefCounted

func (ref RefCounted) Get() Object {
	return Object(C.gd_ref_get_object(C.uintptr_t(ref)))
}

func (ref RefCounted) Set(obj Object) {
	C.gd_ref_set_object(C.uintptr_t(ref), C.uintptr_t(obj))
}

// Editor

func EditorDocumentation(xml string) {
	C.gd_editor_add_documentation((*C.char)(unsafe.Pointer(unsafe.StringData(xml))), C.uint32_t(len(xml)))
}

func EnableEditorPlugin(name Class) {
	C.gd_editor_add_plugin(C.uintptr_t(name))
}

func RemoveEditorPlugin(name Class) {
	C.gd_editor_end_plugin(C.uintptr_t(name))
}

// PropertyList

func MakePropertyList(n int64) PropertyList {
	return PropertyList(C.gd_property_list_make(C.int64_t(n)))
}

func (p PropertyList) Push(prop Property) {
	C.gd_property_list_push(C.uintptr_t(p), C.uint32_t(uint32(prop.Type)), C.uintptr_t(prop.Name), C.uintptr_t(prop.ClassName), C.uint32_t(prop.Hint), C.uintptr_t(prop.HintString), C.uint32_t(prop.Usage), 0)
}

func (p PropertyList) Free() {
	C.gd_property_list_free(C.uintptr_t(p))
}

// MethodList

func MakeMethodList(n int64) MethodList {
	return MethodList(C.gd_method_list_make(C.int64_t(n)))
}

func (m MethodList) Push(name StringName, call ExtensionFunction, flags uint32, returnInfo PropertyList, argsInfo PropertyList, count int64, defaults unsafe.Pointer) {
	C.gd_method_list_push(C.uintptr_t(m), C.uintptr_t(name), C.uintptr_t(functions.New(call)), C.uint32_t(flags), C.uintptr_t(returnInfo), C.uintptr_t(argsInfo), C.int64_t(count), C.UnsafePointer(defaults))
}

func (m MethodList) Free() {
	C.gd_method_list_free(C.uintptr_t(m))
}

// ClassDB registration

func RegisterClass(class string, id ExtensionClass) Class {
	class_name := UTF8.Intern(class)
	parent := id.Parent()
	C.gd_classdb_register(C.StringName(class_name), C.StringName(parent),
		C.ExtensionClassID(classes.New(id)),
		C.bool(id.Virtual()), C.bool(id.Abstract()), C.bool(id.Exposed()), C.bool(id.Runtime()),
		C.String(id.Icon()))
	return Class(class_name)
}

func (class Class) RegisterMethods(methods MethodList) {
	C.gd_classdb_register_methods(C.uintptr_t(class), C.uintptr_t(methods))
}

func (class Class) RegisterConstant(enum, name StringName, value int64, bitfield bool) {
	C.gd_classdb_register_constant(C.uintptr_t(class), C.uintptr_t(enum), C.uintptr_t(name), C.int64_t(value), C.bool(bitfield))
}

func (class Class) RegisterProperty(property Property, setter, getter StringName) {
	pl := MakePropertyList(1)
	pl.Push(property)
	C.gd_classdb_register_property(C.uintptr_t(class), C.uintptr_t(pl), C.uintptr_t(setter), C.uintptr_t(getter))
}

func (class Class) RegisterPropertyIndexed(property Property, setter, getter StringName, index int) {
	pl := MakePropertyList(1)
	pl.Push(property)
	C.gd_classdb_register_property_indexed(C.uintptr_t(class), C.uintptr_t(pl), C.uintptr_t(setter), C.uintptr_t(getter), C.int64_t(index))
}

func (class Class) RegisterPropertyGroup(group, prefix String) {
	C.gd_classdb_register_property_group(C.uintptr_t(class), C.uintptr_t(group), C.uintptr_t(prefix))
}

func (class Class) RegisterPropertySubgroup(subgroup, prefix String) {
	C.gd_classdb_register_property_sub_group(C.uintptr_t(class), C.uintptr_t(subgroup), C.uintptr_t(prefix))
}

func (class Class) RegisterSignal(signal StringName, args PropertyList) {
	C.gd_classdb_register_signal(C.uintptr_t(class), C.uintptr_t(signal), C.uintptr_t(args))
}

func (class Class) Free() {
	C.gd_classdb_register_removal(C.uintptr_t(class))
}

// Iterator

func (v Variant) Iterator() (Iterator, Error) {
	var iter Variant
	var callErr Error
	C.gd_iterator_make(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]),
		C.UnsafePointer(unsafe.Pointer(&iter)), C.UnsafePointer(unsafe.Pointer(&callErr)))
	return Iterator(iter), callErr
}

func (iter *Iterator) Next() (bool, Error) {
	v := Variant(*iter)
	var callErr Error
	ok := bool(C.gd_iterator_next(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]),
		C.UnsafePointer(unsafe.Pointer(iter)), C.UnsafePointer(unsafe.Pointer(&callErr))))
	return ok, callErr
}

func (iter Iterator) Value() (Variant, Error) {
	v := Variant(iter)
	var result Variant
	var callErr Error
	C.gd_iterator_load(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]),
		C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2]),
		C.UnsafePointer(unsafe.Pointer(&result)), C.UnsafePointer(unsafe.Pointer(&callErr)))
	return result, callErr
}

// Callable callbacks

//export gd_on_callable_called
func gd_on_callable_called(c C.CallableID, ret *C.Variant, argc C.Int, args C.VariadicVariants, err *C.CallError) {
	r, e := callables.Get(uintptr(c)).Call(Variants{first: PointerTo[PointerTo[Variant]](unsafe.Pointer(args)), count: int(argc)})
	*ret = C.Variant{C.uint64_t(r[0]), [2]C.uint64_t{C.uint64_t(r[1]), C.uint64_t(r[2])}}
	*err = C.CallError{C.uint32_t(e.error), C.int32_t(e.argument), C.int32_t(e.expected)}
}

//export gd_on_callable_verify
func gd_on_callable_verify(c C.CallableID) C.bool {
	return C.bool(callables.Get(uintptr(c)).IsValid())
}

//export gd_on_callable_delete
func gd_on_callable_delete(c C.CallableID) { callables.Del(uintptr(c)) }

//export gd_on_callable_hashed
func gd_on_callable_hashed(c C.CallableID) C.uint32_t {
	return C.uint32_t(callables.Get(uintptr(c)).Hash())
}

//export gd_on_callable_sorted
func gd_on_callable_sorted(a, b C.CallableID) C.Int {
	return C.Int(callables.Get(uintptr(a)).Compare(callables.Get(uintptr(b))))
}

//export gd_on_callable_string
func gd_on_callable_string(c C.CallableID) C.String {
	return C.String(callables.Get(uintptr(c)).UnsafeString())
}

//export gd_on_callable_length
func gd_on_callable_length(c C.CallableID) C.Int {
	return C.Int(callables.Get(uintptr(c)).ArgumentCount())
}

// Extension binding callbacks

//export gd_on_extension_binding_created
func gd_on_extension_binding_created(p0 C.uintptr_t) C.uintptr_t { return 0 }

//export gd_on_extension_binding_removed
func gd_on_extension_binding_removed(p0, p1 C.uintptr_t) {}

//export gd_on_extension_binding_reference
func gd_on_extension_binding_reference(p0 C.uintptr_t, p1 C.bool) C.bool { return false }

// Extension class callbacks

//export gd_on_extension_class_create
func gd_on_extension_class_create(p0 C.uintptr_t, p1 C.bool) C.uintptr_t {
	return C.uintptr_t(classes.Get(uintptr(p0)).Create(bool(p1)))
}

//export gd_on_extension_class_method
func gd_on_extension_class_method(p0 C.uintptr_t, p1 C.uintptr_t, p2 C.uint32_t) C.uintptr_t {
	fn := classes.Get(uintptr(p0)).Method(StringName(p1), uint32(p2))
	if fn == nil {
		return 0
	}
	return C.uintptr_t(functions.New(fn))
}

//export gd_on_extension_class_caller
func gd_on_extension_class_caller(p0 C.uintptr_t, p1 C.uintptr_t, p2 C.uint32_t) C.uintptr_t {
	fn := classes.Get(uintptr(p0)).Method(StringName(p1), uint32(p2))
	if fn == nil {
		return 0
	}
	return C.uintptr_t(functions.New(fn))
}

// Extension instance callbacks

//export gd_on_extension_instance_set
func gd_on_extension_instance_set(p0 C.uintptr_t, p1 C.uintptr_t, p2, p3, p4 C.uint64_t) C.bool {
	return C.bool(instances.Get(uintptr(p0)).Set(
		StringName(p1), Variant{uint64(p2), uint64(p3), uint64(p4)}))
}

//export gd_on_extension_instance_get
func gd_on_extension_instance_get(p0 C.uintptr_t, p1 C.uintptr_t, p2 *C.Variant) C.bool {
	v, ok := instances.Get(uintptr(p0)).Get(StringName(p1))
	if ok {
		*p2 = C.Variant{C.uint64_t(v[0]), [2]C.uint64_t{C.uint64_t(v[1]), C.uint64_t(v[2])}}
	}
	return C.bool(ok)
}

//export gd_on_extension_instance_property_list
func gd_on_extension_instance_property_list(p0 C.uintptr_t) C.uintptr_t {
	return C.uintptr_t(instances.Get(uintptr(p0)).PropertyList())
}

//export gd_on_extension_instance_property_has_default
func gd_on_extension_instance_property_has_default(p0 C.uintptr_t, p1 C.uintptr_t) C.bool {
	return C.bool(instances.Get(uintptr(p0)).HasDefault(StringName(p1)))
}

//export gd_on_extension_instance_property_get_default
func gd_on_extension_instance_property_get_default(p0 C.uintptr_t, p1 C.uintptr_t, p2 *C.Variant) C.bool {
	v, ok := instances.Get(uintptr(p0)).GetDefault(StringName(p1))
	if ok {
		*p2 = C.Variant{C.uint64_t(v[0]), [2]C.uint64_t{C.uint64_t(v[1]), C.uint64_t(v[2])}}
	}
	return C.bool(ok)
}

//export gd_on_extension_instance_property_validation
func gd_on_extension_instance_property_validation(p0 C.uintptr_t, p1 C.uintptr_t) C.bool {
	return C.bool(instances.Get(uintptr(p0)).ValidateProperty(Property{Name: StringName(p1)}))
}

//export gd_on_extension_instance_notification
func gd_on_extension_instance_notification(p0 C.uintptr_t, p1 C.int32_t, p2 C.bool) {
	instances.Get(uintptr(p0)).Notification(int32(p1), bool(p2))
}

//export gd_on_extension_instance_stringify
func gd_on_extension_instance_stringify(p0 C.uintptr_t) C.uintptr_t {
	return C.uintptr_t(instances.Get(uintptr(p0)).UnsafeString())
}

//export gd_on_extension_instance_reference
func gd_on_extension_instance_reference(p0 C.uintptr_t, p1 C.bool) C.bool {
	return C.bool(instances.Get(uintptr(p0)).Reference(bool(p1)))
}

//export gd_on_extension_instance_rid
func gd_on_extension_instance_rid(p0 C.uintptr_t) C.uint64_t {
	return C.uint64_t(instances.Get(uintptr(p0)).RID())
}

//export gd_on_extension_instance_checked_call
func gd_on_extension_instance_checked_call(p0, p1 C.uintptr_t, p2, p3 C.UnsafePointer) {
	var inst ExtensionInstance
	if uintptr(p0) != 0 {
		inst = instances.Get(uintptr(p0))
	}
	functions.Get(uintptr(p1)).PointerCall(inst, MutablePointer(uintptr(p3)), PointerTo[Pointer](uintptr(p2)))
}

//export gd_on_extension_instance_called
func gd_on_extension_instance_called(p0, p1 C.uintptr_t, p2, p3 C.UnsafePointer) {
	inst := instances.Get(uintptr(p0))
	functions.Get(uintptr(p1)).PointerCall(inst, MutablePointer(uintptr(p3)), PointerTo[Pointer](uintptr(p2)))
}

//export gd_on_extension_instance_variant_call
func gd_on_extension_instance_variant_call(p0 C.uintptr_t, p1 C.uintptr_t, p2 *C.Variant, p3 C.VariadicVariants) {
	var inst ExtensionInstance
	if uintptr(p0) != 0 {
		inst = instances.Get(uintptr(p0))
	}
	v := functions.Get(uintptr(p1)).CheckedCall(inst, Variants{
		first: PointerTo[PointerTo[Variant]](unsafe.Pointer(p3)),
		count: -1,
	})
	*p2 = C.Variant{C.uint64_t(v[0]), [2]C.uint64_t{C.uint64_t(v[1]), C.uint64_t(v[2])}}
}

//export gd_on_extension_instance_dynamic_call
func gd_on_extension_instance_dynamic_call(p0 C.uintptr_t, p1 C.uintptr_t, p2 *C.Variant, p3 C.int64_t, p4 C.VariadicVariants, p5 *C.CallError) {
	var inst ExtensionInstance
	if uintptr(p0) != 0 {
		inst = instances.Get(uintptr(p0))
	}
	v, err := functions.Get(uintptr(p1)).DynamicCall(inst, Variants{
		first: PointerTo[PointerTo[Variant]](unsafe.Pointer(p4)),
		count: int(p3),
	})
	*p2 = C.Variant{C.uint64_t(v[0]), [2]C.uint64_t{C.uint64_t(v[1]), C.uint64_t(v[2])}}
	*p5 = C.CallError{C.uint32_t(err.error), C.int32_t(err.argument), C.int32_t(err.expected)}
}

//export gd_on_extension_instance_free
func gd_on_extension_instance_free(p0 C.uintptr_t) {
	inst := instances.Get(uintptr(p0))
	if f, ok := inst.(interface{ Free() }); ok {
		f.Free()
	}
	instances.Del(uintptr(p0))
}

// Extension script callbacks

//export gd_on_extension_script_categorization
func gd_on_extension_script_categorization(p0 C.uintptr_t, p1 C.uintptr_t) C.bool {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return false
	}
	return C.bool(script.PropertyCategory() != 0)
}

//export gd_on_extension_script_get_property_type
func gd_on_extension_script_get_property_type(p0 C.uintptr_t, name C.uintptr_t, p1 *C.CallError) C.uint32_t {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		*p1 = C.CallError{C.uint32_t(errorInvalidMethod), 0, 0}
		return 0
	}
	return C.uint32_t(script.PropertyType(StringName(name)))
}

//export gd_on_extension_script_get_owner
func gd_on_extension_script_get_owner(p0 C.uintptr_t) C.uintptr_t {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.uintptr_t(script.Owner())
}

//export gd_on_extension_script_get_property_state
func gd_on_extension_script_get_property_state(p0 C.uintptr_t, p1 C.uintptr_t, p2 C.uintptr_t) {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return
	}
	script.ExportedProperties(func(name StringName, value Variant) bool {
		C.gd_object_script_property_state_add(C.FunctionID(p1), C.uintptr_t(p2), C.StringName(name), C.uint64_t(value[0]), C.uint64_t(value[1]), C.uint64_t(value[2]))
		return true
	})
}

//export gd_on_extension_script_get_methods
func gd_on_extension_script_get_methods(p0 C.uintptr_t) C.uintptr_t {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.uintptr_t(script.MethodList())
}

//export gd_on_extension_script_has_method
func gd_on_extension_script_has_method(p0 C.uintptr_t, p1 C.uintptr_t) C.bool {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return false
	}
	return C.bool(script.HasMethod(StringName(p1)))
}

//export gd_on_extension_script_get_method_argument_count
func gd_on_extension_script_get_method_argument_count(p0 C.uintptr_t, p1 C.uintptr_t) C.int64_t {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.int64_t(script.MethodArgumentCount(StringName(p1)))
}

//export gd_on_extension_script_get
func gd_on_extension_script_get(p0 C.uintptr_t) C.uintptr_t {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.uintptr_t(script.Script())
}

//export gd_on_extension_script_is_placeholder
func gd_on_extension_script_is_placeholder(p0 C.uintptr_t) C.bool {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return false
	}
	return C.bool(script.IsPlaceholder())
}

//export gd_on_extension_script_get_language
func gd_on_extension_script_get_language(p0 C.uintptr_t) C.uintptr_t {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return C.uintptr_t(script.ScriptLanguage())
}

// Lifecycle callbacks

//export gd_on_engine_init
func gd_on_engine_init(p0 C.uint32_t) {
	level := InitializationLevel(p0)
	for _, fn := range onEngineInit {
		fn(level)
	}
}

//export gd_on_engine_exit
func gd_on_engine_exit(p0 C.uint32_t) {
	level := InitializationLevel(p0)
	for _, fn := range onEngineExit {
		fn(level)
	}
}

//export gd_on_first_frame
func gd_on_first_frame() {
	for _, fn := range onFirstFrame {
		fn()
	}
}

//export gd_on_every_frame
func gd_on_every_frame() {
	for _, fn := range onEveryFrame {
		fn()
	}
}

//export gd_on_final_frame
func gd_on_final_frame() {
	for _, fn := range onFinalFrame {
		fn()
	}
}

//export gd_on_worker_thread_pool_task
func gd_on_worker_thread_pool_task(p0 C.uintptr_t) {
	if onWorkerThreadPoolTask != nil {
		onWorkerThreadPoolTask(taskID(p0))
	}
}

//export gd_on_worker_thread_pool_group_task
func gd_on_worker_thread_pool_group_task(p0 C.uintptr_t, p1 C.uint32_t) {
	if onWorkerThreadPoolGroupTask != nil {
		onWorkerThreadPoolGroupTask(taskID(p0), int32(p1))
	}
}

//export gd_on_editor_class_in_use_detection
func gd_on_editor_class_in_use_detection(p0, p1 C.uintptr_t, p2 *C.PackedStringArray) {
	if onEditorClassDetection != nil {
		result := onEditorClassDetection(PackedArray[String]{uint64(p0), uint64(p1)})
		p2.array = C.uint64_t(result[0])
		p2.length = C.uint64_t(result[1])
	}
}
