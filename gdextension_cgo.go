//go:build cgo

package gdunsafe

// #include "gd.h"
import "C"
import (
	"unsafe"
)

type (
	String     uintptr
	StringName uintptr
	Array      uintptr
	Dictionary uintptr
	Pointer    uintptr

	VariantType uint32

	Object              uintptr
	ObjectType          uintptr
	MethodForClass      uintptr
	ScriptInstance      uintptr
	ExtensionInstanceID uintptr
	ExtensionClassID    uintptr
	ExtensionBindingID  uintptr
	FunctionID          uintptr
	PropertyList        uintptr
	MethodList          uintptr
)

func (args VariadicVariants) Index(i int) Variant {
	if i >= args.Count || i < 0 {
		panic("index out of range")
	}
	slot := unsafe.Pointer(uintptr(args.First) + unsafe.Sizeof(Pointer(0))*uintptr(i))
	return *(*Variant)(*(*unsafe.Pointer)(slot))
}

func (array Array) Set(index Int, value Variant) {
	C.gd_array_set(C.Array(array), C.int64_t(index), C.uint64_t(value[0]), C.uint64_t(value[1]), C.uint64_t(value[2]))
}

func (array Array) Get(index Int) Variant {
	r := C.gd_array_get(C.Array(array), C.int64_t(index))
	return Variant{uint64(r.tag), uint64(r.payload[0]), uint64(r.payload[1])}
}

func VersionMajor() uint32     { return uint32(C.gd_version_major()) }
func VersionMinor() uint32     { return uint32(C.gd_version_minor()) }
func VersionPatch() uint32     { return uint32(C.gd_version_patch()) }
func VersionHex() uint32       { return uint32(C.gd_version_hex()) }
func VersionStatus() String    { return String(C.gd_version_status()) }
func VersionBuild() String     { return String(C.gd_version_build()) }
func VersionHash() String      { return String(C.gd_version_hash()) }
func VersionTimestamp() uint64 { return uint64(C.gd_version_timestamp()) }
func VersionString() String    { return String(C.gd_version_string()) }
func LibraryLocation() String  { return String(C.gd_library_location()) }

func Malloc(size Int) Pointer    { return Pointer(C.gd_memory_malloc(C.int64_t(size))) }
func Sizeof(name StringName) Int { return Int(C.gd_memory_sizeof(C.uintptr_t(name))) }
func Resize(ptr Pointer, size Int) Pointer {
	return Pointer(C.gd_memory_resize(C.UnsafePointer(ptr), C.int64_t(size)))
}
func Clear(ptr Pointer, size Int) { C.gd_memory_clear(C.UnsafePointer(ptr), C.int64_t(size)) }

func (ptr Pointer) Byte() byte     { return *(*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint16() uint16 { return *(*uint16)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint32() uint32 { return *(*uint32)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint64() uint64 { return *(*uint64)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }

func (ptr Pointer) SetByte(v byte)     { *(*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) = v }
func (ptr Pointer) SetUint16(v uint16) { *(*uint16)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) = v }
func (ptr Pointer) SetUint32(v uint32) { *(*uint32)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) = v }
func (ptr Pointer) SetUint64(v uint64) { *(*uint64)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) = v }

func (ptr Pointer) SetBits128(val [2]uint64) {
	p := *(*unsafe.Pointer)(unsafe.Pointer(&ptr))
	*(*[2]uint64)(p) = val
}
func (ptr Pointer) SetBits256(val [4]uint64) {
	p := *(*unsafe.Pointer)(unsafe.Pointer(&ptr))
	*(*[4]uint64)(p) = val
}
func (ptr Pointer) SetBits512(val [8]uint64) {
	p := *(*unsafe.Pointer)(unsafe.Pointer(&ptr))
	*(*[8]uint64)(p) = val
}

func (ptr Pointer) Free() { C.gd_memory_free(C.UnsafePointer(ptr)) }

func LogError(text, code, fn, file string, line int32, notify_editor bool) {
	C.gd_log_error((*C.char)(unsafe.Pointer(unsafe.StringData(text))), C.int64_t(len(text)), (*C.char)(unsafe.Pointer(unsafe.StringData(code))), C.int64_t(len(code)), (*C.char)(unsafe.Pointer(unsafe.StringData(fn))), C.int64_t(len(fn)), (*C.char)(unsafe.Pointer(unsafe.StringData(file))), C.int64_t(len(file)), C.int32_t(line), C._Bool(notify_editor))
}
func LogWarning(text, code, fn, file string, line int32, notify_editor bool) {
	C.gd_log_warning((*C.char)(unsafe.Pointer(unsafe.StringData(text))), C.int64_t(len(text)), (*C.char)(unsafe.Pointer(unsafe.StringData(code))), C.int64_t(len(code)), (*C.char)(unsafe.Pointer(unsafe.StringData(fn))), C.int64_t(len(fn)), (*C.char)(unsafe.Pointer(unsafe.StringData(file))), C.int64_t(len(file)), C.int32_t(line), C._Bool(notify_editor))
}

func (t VariantType) Name() String {
	return String(C.gd_variant_type_name(C.uint32_t(t)))
}

func (t VariantType) Make(args ...Variant) (value Variant, err CallError) {
	var result Variant
	var cerr CallError
	var argsPtr unsafe.Pointer
	if len(args) > 0 {
		argsPtr = unsafe.Pointer(&args[0])
	}
	C.gd_variant_type_make(C.uint32_t(t), unsafe.Pointer(&result), C.int64_t(len(args)), argsPtr, unsafe.Pointer(&cerr))
	return result, cerr
}

type (
	Callable   [2]uint64
	CallableID uintptr
)

func MakeCallable(impl CallableImplementation, obj ObjectID) Callable {
	var c C.Callable
	C.gd_callable_create(C.CallableID(callables.New(impl)), C.ObjectID(obj), &c)
	return Callable{uint64(c.opaque[0]), uint64(c.opaque[1])}
}

//export gd_on_callable_called
func gd_on_callable_called(c C.CallableID, ret *C.Variant, argc C.Int, args C.VariadicVariants, err *C.CallError) {
	r, e := callables.Get(CallableID(c)).Call(VariadicVariants{First: PointerTo[PointerTo[Variant]](unsafe.Pointer(args)), Count: int(argc)})
	*ret = C.Variant{C.uint64_t(r[0]), [2]C.uint64_t{C.uint64_t(r[1]), C.uint64_t(r[2])}}
	*err = C.CallError{C.uint32_t(e.Type), C.int32_t(e.Argument), C.int32_t(e.Expected)}
}

//export gd_on_callable_verify
func gd_on_callable_verify(c C.CallableID) C.bool {
	return C.bool(callables.Get(CallableID(c)).IsValid())
}

//export gd_on_callable_delete
func gd_on_callable_delete(c C.CallableID) { callables.Del(CallableID(c)) }

//export gd_on_callable_hashed
func gd_on_callable_hashed(c C.CallableID) C.uint32_t {
	return C.uint32_t(callables.Get(CallableID(c)).Hash())
}

//export gd_on_callable_sorted
func gd_on_callable_sorted(a, b C.CallableID) C.Int {
	return C.Int(callables.Get(CallableID(a)).Compare(callables.Get(CallableID(b))))
}

//export gd_on_callable_string
func gd_on_callable_string(c C.CallableID) C.String {
	return C.String(callables.Get(CallableID(c)).UnsafeString())
}

//export gd_on_callable_length
func gd_on_callable_length(c CallableID) C.Int { return C.Int(callables.Get(CallableID(c)).NumIn()) }

// Object construction and identity

func MakeObject(name StringName) Object {
	return Object(C.gd_object_make(C.StringName(name)))
}
func (obj Object) Name() StringName {
	return StringName(C.gd_object_name(C.Object(obj)))
}
func ObjectTypeTag(name StringName) ObjectType {
	return ObjectType(C.gd_object_type(C.StringName(name)))
}
func (obj Object) Cast(to ObjectType) Object {
	return Object(C.gd_object_cast(C.Object(obj), C.ObjectType(to)))
}
func (id ObjectID) Lookup() Object {
	return Object(C.gd_object_lookup(C.ObjectID(id)))
}
func ObjectGlobal(name StringName) Object {
	return Object(C.gd_object_global(C.StringName(name)))
}
func (obj Object) ID() ObjectID {
	return ObjectID(C.gd_object_id(C.Object(obj)))
}
func ObjectIDInsideVariant(v Variant) ObjectID {
	return ObjectID(C.gd_object_id_inside_variant(C.uint64_t(v[0]), C.uint64_t(v[1]), C.uint64_t(v[2])))
}
func (obj Object) Free() {
	C.gd_object_unsafe_free(C.Object(obj))
}

// Object method calls

func MethodLookup(class, method StringName, hash int64) MethodForClass {
	return MethodForClass(C.gd_object_method_lookup(C.StringName(class), C.StringName(method), C.int64_t(hash)))
}
func (obj Object) Call(method MethodForClass, result unsafe.Pointer, argc Int, args unsafe.Pointer, err unsafe.Pointer) {
	C.gd_object_call(C.Object(obj), C.MethodForClass(method), C.UnsafePointer(result), C.int64_t(argc), C.UnsafePointer(args), C.UnsafePointer(err))
}
func (obj Object) UnsafeCall(fn MethodForClass, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	C.gd_object_unsafe_call(C.Object(obj), C.MethodForClass(fn), C.UnsafePointer(result), C.uint64_t(shape), C.UnsafePointer(args))
}

// Extension instance management

func (obj Object) ExtensionSetup(name StringName, inst ExtensionInstanceID) {
	C.gd_object_extension_setup(C.Object(obj), C.StringName(name), C.ExtensionInstanceID(inst))
}
func (obj Object) ExtensionFetch() ExtensionInstanceID {
	return ExtensionInstanceID(C.gd_object_extension_fetch(C.Object(obj)))
}
func (obj Object) ExtensionClose() {
	C.gd_object_extension_close(C.Object(obj))
}

// Script instance management

func ScriptMake(fn ExtensionInstanceID) ScriptInstance {
	return ScriptInstance(C.gd_object_script_make(C.ExtensionInstanceID(fn)))
}
func (obj Object) ScriptCall(name StringName, result unsafe.Pointer, argc Int, args unsafe.Pointer, err unsafe.Pointer) {
	C.gd_object_script_call(C.Object(obj), C.StringName(name), C.UnsafePointer(result), C.int64_t(argc), C.UnsafePointer(args), C.UnsafePointer(err))
}
func (obj Object) ScriptSetup(script ScriptInstance) {
	C.gd_object_script_setup(C.Object(obj), C.ScriptInstance(script))
}
func (obj Object) ScriptFetch(language Object) ScriptInstance {
	return ScriptInstance(C.gd_object_script_fetch(C.Object(obj), C.Object(language)))
}
func (obj Object) ScriptDefinesMethod(method StringName) bool {
	return bool(C.gd_object_script_defines_method(C.Object(obj), C.StringName(method)))
}
func ScriptPropertyStateAdd(fn FunctionID, arg Pointer, name StringName, state Variant) {
	C.gd_object_script_property_state_add(C.FunctionID(fn), C.uintptr_t(arg), C.StringName(name), C.uint64_t(state[0]), C.uint64_t(state[1]), C.uint64_t(state[2]))
}
func ScriptPlaceholderCreate(language, script, owner Object) ScriptInstance {
	return ScriptInstance(C.gd_object_script_placeholder_create(C.Object(language), C.Object(script), C.Object(owner)))
}
func ScriptPlaceholderUpdate(script ScriptInstance, array Array, dict Dictionary) {
	C.gd_object_script_placeholder_update(C.ScriptInstance(script), C.Array(array), C.Dictionary(dict))
}
