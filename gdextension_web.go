//go:build wasm

package gdunsafe

import (
	"sync"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

type (
	String     uint32
	StringName uint32
	Array      uint32
	Dictionary uint32
	Pointer    uint32

	VariantType uint32

	Object              uint32
	ObjectType          uint32
	MethodForClass      uint32
	ScriptInstance      uint32
	ExtensionInstanceID uint32
	ExtensionClassID    uint32
	ExtensionBindingID  uint32
	FunctionID          uint32
	PropertyList        uint32
	MethodList          uint32
)

//go:wasmimport gd array_set
func gd_array_set(array Array, index Int, v1, v2, v3 uint64)

func (array Array) Set(index Int, value Variant) {
	gd_array_set(array, index, value[0], value[1], value[2])
}

//go:wasmimport gd array_get
func gd_array_get(array Array, index Int, result uint32)

func (array Array) Get(index Int) Variant {
	var value Variant
	result := makeResult(gdextension.SizeVariant)
	gd_array_get(array, index, uint32(result))
	loadResult(gdextension.SizeVariant, &value, result)
	return value
}

//go:wasmimport gd version_major
func VersionMajor() uint32

//go:wasmimport gd version_minor
func VersionMinor() uint32

//go:wasmimport gd version_patch
func VersionPatch() uint32

//go:wasmimport gd version_hex
func VersionHex() uint32

//go:wasmimport gd version_status
func VersionStatus() String

//go:wasmimport gd version_build
func VersionBuild() String

//go:wasmimport gd version_hash
func VersionHash() String

//go:wasmimport gd gd_version_timestamp
func VersionTimestamp() uint64

//go:wasmimport gd version_string
func VersionString() String

//go:wasmimport gd library_location
func LibraryLocation() String

//go:wasmimport gd memory_malloc
func Malloc(Int) Pointer

//go:wasmimport gd memory_sizeof
func Sizeof(StringName) Int

//go:wasmimport gd memory_resize
func Resize(Pointer, Int) Pointer

//go:wasmimport gd memory_clear
func Clear(Pointer, Int)

//go:wasmimport gd memory_load_byte
func gd_memory_load_byte(Pointer) uint32

//go:wasmimport gd memory_load_u16
func gd_memory_load_u16(Pointer) uint32

func (ptr Pointer) Byte() byte     { return byte(gd_memory_load_byte(ptr)) }
func (ptr Pointer) Uint16() uint16 { return uint16(gd_memory_load_u16(ptr)) }

//go:wasmimport gd memory_load_u32
func (ptr Pointer) Uint32() uint32

//go:wasmimport gd memory_load_u64
func (ptr Pointer) Uint64() uint64

//go:wasmimport gd memory_edit_byte
func gd_memory_edit_byte(Pointer, uint32)

//go:wasmimport gd memory_edit_u16
func gd_memory_edit_u16(Pointer, uint32)

func (ptr Pointer) SetByte(val byte)     { gd_memory_edit_byte(ptr, uint32(val)) }
func (ptr Pointer) SetUint16(val uint16) { gd_memory_edit_u16(ptr, uint32(val)) }

//go:wasmimport gd memory_edit_u32
func (ptr Pointer) SetUint32(val uint32)

//go:wasmimport gd memory_edit_u64
func (ptr Pointer) SetUint64(val uint64)

//go:wasmimport gd memory_edit_128
func gd_memory_edit_128(Pointer, uint64, uint64)

//go:wasmimport gd memory_edit_256
func gd_memory_edit_256(Pointer, uint64, uint64, uint64, uint64)

//go:wasmimport gd memory_edit_512
func gd_memory_edit_512(Pointer, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64)

func (ptr Pointer) SetBits128(val [2]uint64) {
	gd_memory_edit_128(ptr, val[0], val[1])
}
func (ptr Pointer) SetBits256(val [4]uint64) {
	gd_memory_edit_256(ptr, val[0], val[1], val[2], val[3])
}
func (ptr Pointer) SetBits512(val [8]uint64) {
	gd_memory_edit_512(ptr, val[0], val[1], val[2], val[3], val[4], val[5], val[6], val[7])
}

//go:wasmimport gd memory_free
func (ptr Pointer) Free()

//go:wasmimport gd log_error
func LogError(text, code, fn, file string, line int32, notify_editor bool)

//go:wasmimport gd log_warning
func LogWarning(text, code, fn, file string, line int32, notify_editor bool)

//go:wasmimport gd variant_type_name
func (t VariantType) Name() String

//go:wasmimport gd variant_type_make
func gd_variant_type_make(t VariantType, result Pointer, arg_count Int, args, err Pointer)

func (t VariantType) Make(args ...Variant) (value Variant, err CallError) {
	var param = copyVariants(unsafe.SliceData(args), len(args))
	result := makeResult(gdextension.SizeVariant)
	result_err := makeResult(gdextension.SizeCallError)
	gd_variant_type_make(t, Pointer(result), Int(len(args)), Pointer(param), Pointer(result_err))
	loadResult(gdextension.SizeVariant, &value, result)
	loadResult(gdextension.SizeCallError, &value, result_err)
	return value, err
}

func (args VariadicVariants) Index(i int) Variant {
	if i >= args.Count || i < 0 {
		panic("index out of range")
	}
	// args.First points to an engine-side array of pointers-to-Variant.
	// Read the i-th pointer, then read the Variant it points to.
	ptr := Pointer(gdextension.Pointer(args.First) + gdextension.Pointer(i)*gdextension.Pointer(4)).Uint32()
	if ptr == 0 {
		return Variant{}
	}
	return readVariant(gdextension.Pointer(ptr))
}

type (
	Callable   [2]uint64
	CallableID uint32
)

//go:wasmimport gd callable_create
func gd_callable_create(id CallableID, object ObjectID, result Pointer)

func MakeCallable(impl CallableImplementation, obj ObjectID) Callable {
	result := makeResult(gdextension.SizeCallable)
	gd_callable_create(callables.New(impl), obj, Pointer(result))
	var c Callable
	loadResult(gdextension.SizeCallable, unsafe.Pointer(&c), result)
	return c
}

//go:wasmexport gd_on_callable_called
func gd_on_callable_called(c CallableID, ret Pointer, argc Int, args Pointer, err Pointer) {
	r, e := callables.Get(c).Call(VariadicVariants{First: PointerTo[PointerTo[Variant]](args), Count: int(argc)})
	ret.SetBits128([2]uint64{r[0], r[1]})
	(ret + 16).SetUint64(r[2])
	(err + 0).SetUint32(uint32(e.Type))
	(err + 4).SetInt32(e.Argument)
	(err + 8).SetInt32(e.Expected)
}

//go:wasmexport gd_on_callable_verify
func gd_on_callable_verify(c CallableID) bool {
	return callables.Get(c).IsValid()
}

//go:wasmexport gd_on_callable_delete
func gd_on_callable_delete(c CallableID) { callables.Del(c) }

//go:wasmexport gd_on_callable_hashed
func gd_on_callable_hashed(c CallableID) uint32 {
	return callables.Get(c).Hash()
}

//go:wasmexport gd_on_callable_sorted
func gd_on_callable_sorted(a, b CallableID) int64 {
	return int64(callables.Get(a).Compare(callables.Get(b)))
}

//go:wasmexport gd_on_callable_string
func gd_on_callable_string(c CallableID) String {
	return callables.Get(c).UnsafeString()
}

//go:wasmexport gd_on_callable_length
func gd_on_callable_length(c CallableID) int64 {
	return int64(callables.Get(c).NumIn())
}

// Cross-memory helpers for transferring data between Go and engine address spaces.

var wasmResultBufs [2]gdextension.Pointer
var wasmResultIdx int
var wasmArgBuf gdextension.Pointer

var wasmSetup = sync.OnceFunc(func() {
	wasmArgBuf = gdextension.Pointer(Malloc(64 * 64))
	for i := range wasmResultBufs {
		wasmResultBufs[i] = gdextension.Pointer(Malloc(64 * 64))
		Clear(Pointer(wasmResultBufs[i]), 64*64)
	}
})

func makeResult(shape gdextension.Shape) gdextension.Pointer {
	wasmSetup()
	wasmResultIdx ^= 1
	return wasmResultBufs[wasmResultIdx]
}

func loadResult[T ~unsafe.Pointer | *gdextension.Variant](shape gdextension.Shape, result T, from gdextension.Pointer) {
	wasmSetup()
	if from == 0 {
		panic("nil pointer dereference")
	}
	data := unsafe.Pointer(result)
	done := 0
	size := shape.SizeResult()
	if size == 0 {
		return
	}
	defer Clear(Pointer(from), Int(size))
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(data, done)) = Pointer(from + gdextension.Pointer(done)).Uint32()
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(data, done)) = Pointer(from + gdextension.Pointer(done)).Uint16()
			done += 2
			size -= 2
		default:
			*(*uint8)(unsafe.Add(data, done)) = Pointer(from + gdextension.Pointer(done)).Byte()
			done += 1
			size -= 1
		}
	}
}

func copyVariants[T ~unsafe.Pointer | *gdextension.Variant](args T, n int) gdextension.Pointer {
	wasmSetup()
	var offset int
	var data = unsafe.Pointer(args)
	for i := range n {
		Pointer(wasmArgBuf + gdextension.Pointer(offset)).SetBits128(*(*[2]uint64)(unsafe.Add(data, uintptr(i*24))))
		Pointer(wasmArgBuf + gdextension.Pointer(offset+16)).SetUint64(*(*uint64)(unsafe.Add(data, uintptr(i*24+16))))
		offset += 24
	}
	return wasmArgBuf
}

func copyArguments(shape gdextension.Shape, args unsafe.Pointer) gdextension.Pointer {
	wasmSetup()
	if args == nil {
		return 0
	}
	bytes := shape.SizeArguments()
	buf := unsafe.Slice((*byte)(args), bytes)
	ptr := wasmArgBuf
	off := gdextension.Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			Pointer(ptr + off).SetUint64(*(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			Pointer(ptr + off).SetUint32(*(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			Pointer(ptr + off).SetUint16(*(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		default:
			Pointer(ptr + off).SetByte(*(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return ptr
}

func readVariant(addr gdextension.Pointer) Variant {
	if addr == 0 {
		panic("nil pointer dereference")
	}
	var v Variant
	v[0] = Pointer(addr).Uint64()
	v[1] = Pointer(addr + 8).Uint64()
	v[2] = Pointer(addr + 16).Uint64()
	return v
}

// Object construction and identity

//go:wasmimport gd object_make
func gd_object_make(name StringName) Object

func MakeObject(name StringName) Object { return gd_object_make(name) }

//go:wasmimport gd object_name
func gd_object_name(obj Object) StringName

func (obj Object) Name() StringName { return gd_object_name(obj) }

//go:wasmimport gd object_type
func gd_object_type(name StringName) ObjectType

func ObjectTypeTag(name StringName) ObjectType { return gd_object_type(name) }

//go:wasmimport gd object_cast
func gd_object_cast(obj Object, to ObjectType) Object

func (obj Object) Cast(to ObjectType) Object { return gd_object_cast(obj, to) }

//go:wasmimport gd object_lookup
func gd_object_lookup(id ObjectID) Object

func (id ObjectID) Lookup() Object { return gd_object_lookup(id) }

//go:wasmimport gd object_global
func gd_object_global(name StringName) Object

func ObjectGlobal(name StringName) Object { return gd_object_global(name) }

//go:wasmimport gd object_id
func gd_object_id(obj Object) ObjectID

func (obj Object) ID() ObjectID { return gd_object_id(obj) }

//go:wasmimport gd object_id_inside_variant
func gd_object_id_inside_variant(v1, v2, v3 uint64) ObjectID

func ObjectIDInsideVariant(v Variant) ObjectID {
	return gd_object_id_inside_variant(v[0], v[1], v[2])
}

//go:wasmimport gd object_unsafe_free
func gd_object_unsafe_free(obj Object)

func (obj Object) Free() { gd_object_unsafe_free(obj) }

// Object method calls

//go:wasmimport gd object_method_lookup
func gd_object_method_lookup(class, method StringName, hash int64) MethodForClass

func MethodLookup(class, method StringName, hash int64) MethodForClass {
	return gd_object_method_lookup(class, method, hash)
}

//go:wasmimport gd object_call
func gd_object_call(obj Object, method MethodForClass, result Pointer, argc Int, args Pointer, err Pointer)

func (obj Object) Call(method MethodForClass, result unsafe.Pointer, argc Int, args unsafe.Pointer, err unsafe.Pointer) {
	mem_result := makeResult(gdextension.SizeVariant)
	mem_args := copyVariants(args, int(argc))
	mem_err := makeResult(gdextension.SizeCallError)
	gd_object_call(obj, method, Pointer(mem_result), argc, Pointer(mem_args), Pointer(mem_err))
	loadResult(gdextension.SizeVariant, result, mem_result)
	loadResult(gdextension.SizeCallError, err, mem_err)
}

//go:wasmimport gd object_unsafe_call
func gd_object_unsafe_call(obj Object, fn MethodForClass, result Pointer, shape uint64, args Pointer)

func (obj Object) UnsafeCall(fn MethodForClass, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(gdextension.Shape(shape))
	mem_args := copyArguments(gdextension.Shape(shape), args)
	gd_object_unsafe_call(obj, fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(gdextension.Shape(shape), result, mem_result)
}

// Extension instance management

//go:wasmimport gd object_extension_setup
func gd_object_extension_setup(obj Object, name StringName, inst ExtensionInstanceID)

func (obj Object) ExtensionSetup(name StringName, inst ExtensionInstanceID) {
	gd_object_extension_setup(obj, name, inst)
}

//go:wasmimport gd object_extension_fetch
func gd_object_extension_fetch(obj Object) ExtensionInstanceID

func (obj Object) ExtensionFetch() ExtensionInstanceID { return gd_object_extension_fetch(obj) }

//go:wasmimport gd object_extension_close
func gd_object_extension_close(obj Object)

func (obj Object) ExtensionClose() { gd_object_extension_close(obj) }

// Script instance management

//go:wasmimport gd object_script_make
func gd_object_script_make(fn ExtensionInstanceID) ScriptInstance

func ScriptMake(fn ExtensionInstanceID) ScriptInstance { return gd_object_script_make(fn) }

//go:wasmimport gd object_script_call
func gd_object_script_call(obj Object, name StringName, result Pointer, argc Int, args Pointer, err Pointer)

func (obj Object) ScriptCall(name StringName, result unsafe.Pointer, argc Int, args unsafe.Pointer, err unsafe.Pointer) {
	mem_result := makeResult(gdextension.SizeVariant)
	mem_args := copyVariants(args, int(argc))
	mem_err := makeResult(gdextension.SizeCallError)
	gd_object_script_call(obj, name, Pointer(mem_result), argc, Pointer(mem_args), Pointer(mem_err))
	loadResult(gdextension.SizeVariant, result, mem_result)
	loadResult(gdextension.SizeCallError, err, mem_err)
}

//go:wasmimport gd object_script_setup
func gd_object_script_setup(obj Object, script ScriptInstance)

func (obj Object) ScriptSetup(script ScriptInstance) { gd_object_script_setup(obj, script) }

//go:wasmimport gd object_script_fetch
func gd_object_script_fetch(obj Object, language Object) ScriptInstance

func (obj Object) ScriptFetch(language Object) ScriptInstance {
	return gd_object_script_fetch(obj, language)
}

//go:wasmimport gd object_script_defines_method
func gd_object_script_defines_method(obj Object, method StringName) uint32

func (obj Object) ScriptDefinesMethod(method StringName) bool {
	return gd_object_script_defines_method(obj, method) != 0
}

//go:wasmimport gd object_script_property_state_add
func gd_object_script_property_state_add(fn FunctionID, arg uint32, name StringName, s1, s2, s3 uint64)

func ScriptPropertyStateAdd(fn FunctionID, arg Pointer, name StringName, state Variant) {
	gd_object_script_property_state_add(fn, uint32(arg), name, state[0], state[1], state[2])
}

//go:wasmimport gd object_script_placeholder_create
func gd_object_script_placeholder_create(language, script, owner Object) ScriptInstance

func ScriptPlaceholderCreate(language, script, owner Object) ScriptInstance {
	return gd_object_script_placeholder_create(language, script, owner)
}

//go:wasmimport gd object_script_placeholder_update
func gd_object_script_placeholder_update(script ScriptInstance, array Array, dict Dictionary)

func ScriptPlaceholderUpdate(script ScriptInstance, array Array, dict Dictionary) {
	gd_object_script_placeholder_update(script, array, dict)
}
