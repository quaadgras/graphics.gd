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
	Pointer    uintptr

	VariantType uint32
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
func VersionTimestamp() uint64  { return uint64(C.gd_version_timestamp()) }
func VersionString() String    { return String(C.gd_version_string()) }
func LibraryLocation() String  { return String(C.gd_library_location()) }

func Malloc(size Int) Pointer  { return Pointer(C.gd_memory_malloc(C.int64_t(size))) }
func Sizeof(name StringName) Int { return Int(C.gd_memory_sizeof(C.uintptr_t(name))) }
func Resize(ptr Pointer, size Int) Pointer { return Pointer(C.gd_memory_resize(C.uintptr_t(ptr), C.int64_t(size))) }
func Clear(ptr Pointer, size Int) { C.gd_memory_clear(C.uintptr_t(ptr), C.int64_t(size)) }

func (ptr Pointer) Byte() byte       { return *(*byte)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint16() uint16   { return *(*uint16)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint32() uint32   { return *(*uint32)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }
func (ptr Pointer) Uint64() uint64   { return *(*uint64)(*(*unsafe.Pointer)(unsafe.Pointer(&ptr))) }

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

func (ptr Pointer) Free() { C.gd_memory_free(C.uintptr_t(ptr)) }

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
