//go:build wasm

package gdunsafe

import (
	"sync"
	"unsafe"

	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

type (
	String     uint32
	StringName uint32
	Array      uint32
	Dictionary uint32
	Pointer    uint32

	PackedArray[T byte | int32 | int64 | float32 | float64 | Color.RGBA | Vector2.XY | Vector3.XYZ | Vector4.XYZW | String] [2]uint32

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
	result := makeResult(SizeVariant)
	gd_array_get(array, index, uint32(result))
	loadResult(SizeVariant, &value, result)
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

// String operations — buffer copy helpers

func copyBufToEngine(buf []byte) Pointer {
	ptr := Malloc(Int(len(buf)))
	off := Pointer(0)
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
			Pointer(ptr + off).SetByte(buf[0])
			buf = buf[1:]
			off += 1
		}
	}
	return ptr
}

func copyBufToGo(ptr Pointer, buf []byte) {
	off := 0
	for len(buf) > 0 {
		switch {
		case len(buf) >= 4:
			*(*uint32)(unsafe.Pointer(&buf[0])) = Pointer(ptr + Pointer(off)).Uint32()
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = Pointer(ptr + Pointer(off)).Uint16()
			buf = buf[2:]
			off += 2
		default:
			buf[0] = Pointer(ptr + Pointer(off)).Byte()
			buf = buf[1:]
			off += 1
		}
	}
	ptr.Free()
}

// String basic operations

//go:wasmimport gd string_access
func gd_string_access_wasm(s uint32, idx int64) int32

func (s String) Access(idx Int) int32 {
	return gd_string_access_wasm(uint32(s), int64(idx))
}

//go:wasmimport gd string_resize
func gd_string_resize_wasm(s uint32, size int64) uint32

func (s String) Resize(size Int) String {
	return String(gd_string_resize_wasm(uint32(s), int64(size)))
}

//go:wasmimport gd string_unsafe
func gd_string_unsafe_wasm(s uint32) uint32

func (s String) UnsafePtr() Pointer {
	return Pointer(gd_string_unsafe_wasm(uint32(s)))
}

//go:wasmimport gd string_append
func gd_string_append_wasm(s uint32, other uint32) uint32

func (s String) Append(other String) String {
	return String(gd_string_append_wasm(uint32(s), uint32(other)))
}

//go:wasmimport gd string_append_rune
func gd_string_append_rune_wasm(s uint32, ch int32) uint32

func (s String) AppendRune(ch int32) String {
	return String(gd_string_append_rune_wasm(uint32(s), ch))
}

//go:wasmimport gd string_decode
func gd_string_decode_wasm(enc uint32, s uint32, length int64) uint32

//go:wasmimport gd string_encode
func gd_string_encode_wasm(enc uint32, s uint32, buf uint32, cap int64) int64

//go:wasmimport gd string_intern
func gd_string_intern_wasm(enc uint32, s uint32, length int64) uint32

func (enc StringEncoding) String(s string) String {
	buf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := String(gd_string_decode_wasm(uint32(enc), uint32(buf), int64(len(s))))
	buf.Free()
	return result
}

func (s String) Encode(enc StringEncoding, buf []byte) Int {
	ebuf := copyBufToEngine(buf)
	n := Int(gd_string_encode_wasm(uint32(enc), uint32(s), uint32(ebuf), int64(len(buf))))
	copyBufToGo(ebuf, buf)
	return n
}

func (enc StringEncoding) Intern(s string) StringName {
	buf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := StringName(gd_string_intern_wasm(uint32(enc), uint32(buf), int64(len(s))))
	buf.Free()
	return result
}

//go:wasmimport gd log
func gd_log_wasm(level uint32, text uint32, text_len int32, code uint32, code_len int32, fn uint32, fn_len int32, file uint32, file_len int32, line int32, notify_editor uint32)

func Log(level LogLevel, text, code, fn, file string, line int32, notify_editor bool) {
	etext := copyBufToEngine(unsafe.Slice(unsafe.StringData(text), len(text)))
	ecode := copyBufToEngine(unsafe.Slice(unsafe.StringData(code), len(code)))
	efn := copyBufToEngine(unsafe.Slice(unsafe.StringData(fn), len(fn)))
	efile := copyBufToEngine(unsafe.Slice(unsafe.StringData(file), len(file)))
	var ne uint32
	if notify_editor {
		ne = 1
	}
	gd_log_wasm(uint32(level), uint32(etext), int32(len(text)), uint32(ecode), int32(len(code)), uint32(efn), int32(len(fn)), uint32(efile), int32(len(file)), line, ne)
	etext.Free()
	ecode.Free()
	efn.Free()
	efile.Free()
}

//go:wasmimport gd packed_array_access
func gd_packed_array_access(t uint32, a1, a2 uint32, idx int64) uint32

//go:wasmimport gd packed_array_modify
func gd_packed_array_modify(t uint32, a1, a2 uint32, idx int64) uint32

func (p PackedArray[T]) Access(idx Int) PointerTo[T] {
	return PointerTo[T](gd_packed_array_access(uint32(p.Type()), p[0], p[1], int64(idx)))
}

func (p PackedArray[T]) Modify(idx Int) PointerTo[T] {
	return PointerTo[T](gd_packed_array_modify(uint32(p.Type()), p[0], p[1], int64(idx)))
}

func (ptr PointerTo[T]) Get() T {
	var v T
	buf := unsafe.Slice((*byte)(unsafe.Pointer(&v)), unsafe.Sizeof(v))
	copyBufToGo2(Pointer(ptr), buf)
	return v
}

func (ptr PointerTo[T]) Set(v T) {
	buf := unsafe.Slice((*byte)(unsafe.Pointer(&v)), unsafe.Sizeof(v))
	copyBufToEngine2(Pointer(ptr), buf)
}

func copyBufToGo2(ptr Pointer, buf []byte) {
	off := 0
	for len(buf) > 0 {
		switch {
		case len(buf) >= 4:
			*(*uint32)(unsafe.Pointer(&buf[0])) = Pointer(ptr + Pointer(off)).Uint32()
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = Pointer(ptr + Pointer(off)).Uint16()
			buf = buf[2:]
			off += 2
		default:
			buf[0] = Pointer(ptr + Pointer(off)).Byte()
			buf = buf[1:]
			off += 1
		}
	}
}

func copyBufToEngine2(ptr Pointer, buf []byte) {
	off := Pointer(0)
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
			Pointer(ptr + off).SetByte(buf[0])
			buf = buf[1:]
			off += 1
		}
	}
}

//go:wasmimport gd variant_type_name
func (t VariantType) Name() String

//go:wasmimport gd variant_type_make
func gd_variant_type_make(t VariantType, result Pointer, arg_count Int, args, err Pointer)

func (t VariantType) Make(args ...Variant) (value Variant, err CallError) {
	var param = copyVariants(unsafe.SliceData(args), len(args))
	result := makeResult(SizeVariant)
	result_err := makeResult(SizeCallError)
	gd_variant_type_make(t, Pointer(result), Int(len(args)), Pointer(param), Pointer(result_err))
	loadResult(SizeVariant, &value, result)
	loadResult(SizeCallError, &value, result_err)
	return value, err
}

//go:wasmimport gd variant_type_call
func gd_variant_type_call_wasm(t VariantType, method StringName, result Pointer, argc Int, args Pointer, err Pointer)

func (t VariantType) StaticCall(method StringName, args ...Variant) (value Variant, callErr CallError) {
	param := copyVariants(unsafe.SliceData(args), len(args))
	result := makeResult(SizeVariant)
	result_err := makeResult(SizeCallError)
	gd_variant_type_call_wasm(t, method, Pointer(result), Int(len(args)), Pointer(param), Pointer(result_err))
	loadResult(SizeVariant, &value, result)
	loadResult(SizeCallError, &callErr, result_err)
	return value, callErr
}

//go:wasmimport gd variant_type_convertable
func gd_variant_type_convertable_wasm(t VariantType, to VariantType, strict uint32) uint32

func (t VariantType) Convertable(to VariantType, strict bool) bool {
	var s uint32
	if strict {
		s = 1
	}
	return gd_variant_type_convertable_wasm(t, to, s) != 0
}

//go:wasmimport gd builtin_name
func gd_builtin_name_wasm(utility StringName, hash int64) FunctionID

func BuiltinName(utility StringName, hash int64) FunctionID {
	return gd_builtin_name_wasm(utility, hash)
}

//go:wasmimport gd builtin_call
func gd_builtin_call_wasm(fn FunctionID, result Pointer, shape uint64, args Pointer)

func BuiltinCall(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_builtin_call_wasm(fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_type_setup_array
func gd_variant_type_setup_array_wasm(array Array, vtype VariantType, className StringName, v1, v2, v3 uint64)

func VariantTypeSetupArray(array Array, vtype VariantType, className StringName, script Variant) {
	gd_variant_type_setup_array_wasm(array, vtype, className, script[0], script[1], script[2])
}

//go:wasmimport gd variant_type_setup_dictionary
func gd_variant_type_setup_dictionary_wasm(dict Dictionary, keyType VariantType, keyClassName StringName, ks1, ks2, ks3 uint64, valType VariantType, valClassName StringName, vs1, vs2, vs3 uint64)

func VariantTypeSetupDictionary(dict Dictionary, keyType VariantType, keyClassName StringName, keyScript Variant, valType VariantType, valClassName StringName, valScript Variant) {
	gd_variant_type_setup_dictionary_wasm(dict, keyType, keyClassName, keyScript[0], keyScript[1], keyScript[2], valType, valClassName, valScript[0], valScript[1], valScript[2])
}

//go:wasmimport gd variant_type_fetch_constant
func gd_variant_type_fetch_constant_wasm(vtype VariantType, constant StringName, result Pointer)

func VariantTypeFetchConstant(vtype VariantType, constant StringName, result unsafe.Pointer) {
	mem := makeResult(SizeVariant)
	gd_variant_type_fetch_constant_wasm(vtype, constant, Pointer(mem))
	loadResult(SizeVariant, result, mem)
}

//go:wasmimport gd variant_type_unsafe_constructor
func gd_variant_type_unsafe_constructor_wasm(vtype VariantType, n Int) FunctionID

func VariantTypeConstructor(vtype VariantType, n Int) FunctionID {
	return gd_variant_type_unsafe_constructor_wasm(vtype, n)
}

//go:wasmimport gd variant_type_evaluator
func gd_variant_type_evaluator_wasm(op VariantOperator, a, b VariantType) FunctionID

func VariantTypeEvaluator(op VariantOperator, a, b VariantType) FunctionID {
	return gd_variant_type_evaluator_wasm(op, a, b)
}

//go:wasmimport gd variant_type_setter
func gd_variant_type_setter_wasm(vtype VariantType, property StringName) FunctionID

func VariantTypeSetter(vtype VariantType, property StringName) FunctionID {
	return gd_variant_type_setter_wasm(vtype, property)
}

//go:wasmimport gd variant_type_getter
func gd_variant_type_getter_wasm(vtype VariantType, property StringName) FunctionID

func VariantTypeGetter(vtype VariantType, property StringName) FunctionID {
	return gd_variant_type_getter_wasm(vtype, property)
}

//go:wasmimport gd variant_type_has_property
func gd_variant_type_has_property_wasm(vtype VariantType, property StringName) uint32

func VariantTypeHasProperty(vtype VariantType, property StringName) bool {
	return gd_variant_type_has_property_wasm(vtype, property) != 0
}

//go:wasmimport gd variant_type_builtin_method
func gd_variant_type_builtin_method_wasm(vtype VariantType, method StringName, hash int64) FunctionID

func VariantTypeMethod(vtype VariantType, method StringName, hash int64) FunctionID {
	return gd_variant_type_builtin_method_wasm(vtype, method, hash)
}

//go:wasmimport gd variant_type_unsafe_call
func gd_variant_type_unsafe_call_wasm(self Pointer, fn FunctionID, result Pointer, shape uint64, args Pointer)

func VariantTypeUnsafeCall(self unsafe.Pointer, fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	selfShape := Shape(shape) >> 4
	mem_self := copySelf(selfShape, self)
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(selfShape, args)
	gd_variant_type_unsafe_call_wasm(Pointer(mem_self), fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(selfShape, self, mem_self)
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_type_unsafe_make
func gd_variant_type_unsafe_make_wasm(constructor FunctionID, result Pointer, shape uint64, args Pointer)

func VariantTypeUnsafeMake(constructor FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_type_unsafe_make_wasm(constructor, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_type_unsafe_free
func gd_variant_type_unsafe_free_wasm(vtype VariantType, shape uint64, args Pointer)

func VariantTypeUnsafeFree(vtype VariantType, shape uint64, args unsafe.Pointer) {
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_type_unsafe_free_wasm(vtype, shape, Pointer(mem_args))
}

func (args VariadicVariants) Index(i int) Variant {
	if i >= args.Count || i < 0 {
		panic("index out of range")
	}
	// args.First points to an engine-side array of pointers-to-Variant.
	// Read the i-th pointer, then read the Variant it points to.
	ptr := Pointer(Pointer(args.First) + Pointer(i)*Pointer(4)).Uint32()
	if ptr == 0 {
		return Variant{}
	}
	return readVariant(Pointer(ptr))
}

type (
	Callable   [2]uint64
	CallableID uint32
)

//go:wasmimport gd callable_create
func gd_callable_create(id CallableID, object ObjectID, result Pointer)

func MakeCallable(impl ExtensionCallable, obj ObjectID) Callable {
	result := makeResult(SizeCallable)
	gd_callable_create(callables.New(impl), obj, Pointer(result))
	var c Callable
	loadResult(SizeCallable, unsafe.Pointer(&c), result)
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
	return int64(callables.Get(c).ArgumentCount())
}

// Cross-memory helpers for transferring data between Go and engine address spaces.

var wasmResultBufs [2]Pointer
var wasmResultIdx int
var wasmArgBuf Pointer
var wasmSelfBuf Pointer

var wasmSetup = sync.OnceFunc(func() {
	wasmArgBuf = Malloc(64 * 64)
	wasmSelfBuf = Malloc(64)
	for i := range wasmResultBufs {
		wasmResultBufs[i] = Malloc(64 * 64)
		Clear(wasmResultBufs[i], 64*64)
	}
})

func makeResult(shape Shape) Pointer {
	wasmSetup()
	wasmResultIdx ^= 1
	return wasmResultBufs[wasmResultIdx]
}

func loadResult[T ~unsafe.Pointer | *Variant | *CallError](shape Shape, result T, from Pointer) {
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
			*(*uint32)(unsafe.Add(data, done)) = Pointer(from + Pointer(done)).Uint32()
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(data, done)) = Pointer(from + Pointer(done)).Uint16()
			done += 2
			size -= 2
		default:
			*(*uint8)(unsafe.Add(data, done)) = Pointer(from + Pointer(done)).Byte()
			done += 1
			size -= 1
		}
	}
}

func copyVariants[T ~unsafe.Pointer | *Variant](args T, n int) Pointer {
	wasmSetup()
	var offset int
	var data = unsafe.Pointer(args)
	for i := range n {
		Pointer(wasmArgBuf + Pointer(offset)).SetBits128(*(*[2]uint64)(unsafe.Add(data, uintptr(i*24))))
		Pointer(wasmArgBuf + Pointer(offset+16)).SetUint64(*(*uint64)(unsafe.Add(data, uintptr(i*24+16))))
		offset += 24
	}
	return wasmArgBuf
}

// copySelf copies self data to wasmSelfBuf. Uses SizeResult() because the
// self type is at nibble 0 of the (already shifted) shape.
func copySelf(selfShape Shape, self unsafe.Pointer) Pointer {
	wasmSetup()
	if self == nil {
		return 0
	}
	size := selfShape.SizeResult()
	if size == 0 {
		return 0
	}
	buf := unsafe.Slice((*byte)(self), size)
	off := Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			Pointer(wasmSelfBuf + off).SetUint64(*(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			Pointer(wasmSelfBuf + off).SetUint32(*(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			Pointer(wasmSelfBuf + off).SetUint16(*(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		default:
			Pointer(wasmSelfBuf + off).SetByte(*(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return wasmSelfBuf
}

func copyArgumentsTo(shape Shape, args unsafe.Pointer, target Pointer) Pointer {
	if args == nil {
		return 0
	}
	bytes := shape.SizeArguments()
	buf := unsafe.Slice((*byte)(args), bytes)
	off := Pointer(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			Pointer(target + off).SetUint64(*(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			Pointer(target + off).SetUint32(*(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			Pointer(target + off).SetUint16(*(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		default:
			Pointer(target + off).SetByte(*(*uint8)(unsafe.Pointer(&buf[0])))
			buf = buf[1:]
			off += 1
		}
	}
	return target
}

func copyArguments(shape Shape, args unsafe.Pointer) Pointer {
	wasmSetup()
	return copyArgumentsTo(shape, args, wasmArgBuf)
}

func writeVariant(addr Pointer, v Variant) {
	Pointer(addr).SetUint64(v[0])
	Pointer(addr + 8).SetUint64(v[1])
	Pointer(addr + 16).SetUint64(v[2])
}

func writeCallError(addr Pointer, e CallError) {
	Pointer(addr).SetUint32(uint32(e.Type))
	Pointer(addr + 4).SetInt32(e.Argument)
	Pointer(addr + 8).SetInt32(e.Expected)
}

func readVariant(addr Pointer) Variant {
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

func (obj Object) Call(method MethodForClass, args ...Variant) (Variant, CallError) {
	mem_result := makeResult(SizeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(SizeCallError)
	gd_object_call(obj, method, Pointer(mem_result), Int(len(args)), Pointer(mem_args), Pointer(mem_err))
	var result Variant
	var errResult CallError
	loadResult(SizeVariant, &result, mem_result)
	loadResult(SizeCallError, &errResult, mem_err)
	return result, errResult
}

//go:wasmimport gd object_shaped_call
func gd_object_shaped_call(obj Object, fn MethodForClass, result Pointer, shape uint64, args Pointer)

func (obj Object) ShapedCall(fn MethodForClass, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_object_shaped_call(obj, fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

// Extension instance management

//go:wasmimport gd object_extension_setup
func gd_object_extension_setup(obj Object, name StringName, inst ExtensionInstanceID)

func (obj Object) ExtensionSetup(name StringName, inst ExtensionInstance) {
	gd_object_extension_setup(obj, name, instances.New(inst))
}

//go:wasmimport gd object_extension_fetch
func gd_object_extension_fetch(obj Object) ExtensionInstanceID

func (obj Object) ExtensionFetch() ExtensionInstance {
	return instances.Get(gd_object_extension_fetch(obj))
}

//go:wasmimport gd object_extension_close
func gd_object_extension_close(obj Object)

func (obj Object) ExtensionClose() { gd_object_extension_close(obj) }

// Script instance management

//go:wasmimport gd object_script_make
func gd_object_script_make(fn ExtensionInstanceID) ScriptInstance

func ScriptMake(fn ExtensionInstanceID) ScriptInstance { return gd_object_script_make(fn) }

//go:wasmimport gd object_script_call
func gd_object_script_call(obj Object, name StringName, result Pointer, argc Int, args Pointer, err Pointer)

func (obj Object) ScriptCall(name StringName, args ...Variant) (Variant, CallError) {
	mem_result := makeResult(SizeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(SizeCallError)
	gd_object_script_call(obj, name, Pointer(mem_result), Int(len(args)), Pointer(mem_args), Pointer(mem_err))
	var result Variant
	var err CallError
	loadResult(SizeVariant, &result, mem_result)
	loadResult(SizeCallError, &err, mem_err)
	return result, err
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

// Variant operations

type VariantOperator = uint32

//go:wasmimport gd variant_zero
func gd_variant_zero(result Pointer)

func ZeroVariant() Variant {
	mem := makeResult(SizeVariant)
	gd_variant_zero(Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result
}

//go:wasmimport gd variant_copy
func gd_variant_copy(v1, v2, v3 uint64, result Pointer)

func (v Variant) Copy() Variant {
	mem := makeResult(SizeVariant)
	gd_variant_copy(v[0], v[1], v[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result
}

//go:wasmimport gd variant_call
func gd_variant_call(v1, v2, v3 uint64, method StringName, result Pointer, argc Int, args Pointer, err Pointer)

func (v Variant) VariantCall(method StringName, args ...Variant) (Variant, CallError) {
	mem_result := makeResult(SizeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(SizeCallError)
	gd_variant_call(v[0], v[1], v[2], method, Pointer(mem_result), Int(len(args)), Pointer(mem_args), Pointer(mem_err))
	var result Variant
	var err CallError
	loadResult(SizeVariant, &result, mem_result)
	loadResult(SizeCallError, &err, mem_err)
	return result, err
}

//go:wasmimport gd variant_eval
func gd_variant_eval(op uint32, a1, a2, a3, b1, b2, b3 uint64, result Pointer) uint32

func VariantEval(op VariantOperator, a, b Variant) (Variant, bool) {
	mem := makeResult(SizeVariant)
	r := gd_variant_eval(op, a[0], a[1], a[2], b[0], b[1], b[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_hash
func gd_variant_hash(v1, v2, v3 uint64) int64

func (v Variant) Hash() Int { return gd_variant_hash(v[0], v[1], v[2]) }

//go:wasmimport gd variant_bool
func gd_variant_bool(v1, v2, v3 uint64) uint32

func (v Variant) Bool() bool {
	return gd_variant_bool(v[0], v[1], v[2]) != 0
}

//go:wasmimport gd variant_text
func gd_variant_text(v1, v2, v3 uint64) String

func (v Variant) Text() String {
	return gd_variant_text(v[0], v[1], v[2])
}

//go:wasmimport gd variant_type
func gd_variant_type(v1, v2, v3 uint64) VariantType

func (v Variant) Type() VariantType {
	return gd_variant_type(v[0], v[1], v[2])
}

// Deep variant operations

//go:wasmimport gd variant_deep_copy
func gd_variant_deep_copy(v1, v2, v3 uint64, result Pointer)

func (v Variant) DeepCopy() Variant {
	mem := makeResult(SizeVariant)
	gd_variant_deep_copy(v[0], v[1], v[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result
}

//go:wasmimport gd variant_deep_hash
func gd_variant_deep_hash(v1, v2, v3 uint64, recursion Int) Int

func (v Variant) DeepHash(recursion Int) Int {
	return gd_variant_deep_hash(v[0], v[1], v[2], recursion)
}

// Variant get/set/has

//go:wasmimport gd variant_get_index
func gd_variant_get_index(v1, v2, v3, k1, k2, k3 uint64, result Pointer) uint32

func (v Variant) GetIndex(key Variant) (Variant, bool) {
	mem := makeResult(SizeVariant)
	r := gd_variant_get_index(v[0], v[1], v[2], key[0], key[1], key[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_get_array
func gd_variant_get_array(v1, v2, v3 uint64, idx Int, result Pointer, err Pointer) uint32

func (v Variant) GetArray(idx Int) (Variant, bool, CallError) {
	mem_result := makeResult(SizeVariant)
	mem_err := makeResult(SizeCallError)
	r := gd_variant_get_array(v[0], v[1], v[2], idx, Pointer(mem_result), Pointer(mem_err))
	var result Variant
	var callErr CallError
	loadResult(SizeVariant, &result, mem_result)
	loadResult(SizeCallError, &callErr, mem_err)
	return result, r != 0, callErr
}

//go:wasmimport gd variant_get_field
func gd_variant_get_field(v1, v2, v3 uint64, field StringName, result Pointer) uint32

func (v Variant) GetField(field StringName) (Variant, bool) {
	mem := makeResult(SizeVariant)
	r := gd_variant_get_field(v[0], v[1], v[2], field, Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_set_index
func gd_variant_set_index(v1, v2, v3, k1, k2, k3, val1, val2, val3 uint64) uint32

func (v Variant) SetIndex(key, val Variant) bool {
	return gd_variant_set_index(v[0], v[1], v[2], key[0], key[1], key[2], val[0], val[1], val[2]) != 0
}

//go:wasmimport gd variant_set_array
func gd_variant_set_array(v1, v2, v3 uint64, idx Int, val1, val2, val3 uint64, err Pointer) uint32

func (v Variant) SetArray(idx Int, val Variant) (bool, CallError) {
	mem_err := makeResult(SizeCallError)
	r := gd_variant_set_array(v[0], v[1], v[2], idx, val[0], val[1], val[2], Pointer(mem_err))
	var callErr CallError
	loadResult(SizeCallError, &callErr, mem_err)
	return r != 0, callErr
}

//go:wasmimport gd variant_set_field
func gd_variant_set_field(v1, v2, v3 uint64, field StringName, val1, val2, val3 uint64) uint32

func (v Variant) SetField(field StringName, value Variant) bool {
	return gd_variant_set_field(v[0], v[1], v[2], field, value[0], value[1], value[2]) != 0
}

//go:wasmimport gd variant_has_index
func gd_variant_has_index(v1, v2, v3, i1, i2, i3 uint64) uint32

func (v Variant) HasIndex(index Variant) bool {
	return gd_variant_has_index(v[0], v[1], v[2], index[0], index[1], index[2]) != 0
}

//go:wasmimport gd variant_has_method
func gd_variant_has_method(v1, v2, v3 uint64, method StringName) uint32

func (v Variant) HasMethod(method StringName) bool {
	return gd_variant_has_method(v[0], v[1], v[2], method) != 0
}

// Unsafe variant operations

//go:wasmimport gd variant_unsafe_call
func gd_variant_unsafe_call(fn FunctionID, result Pointer, shape uint64, args Pointer)

func VariantUnsafeCall(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_call(fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_unsafe_eval
func gd_variant_unsafe_eval(fn FunctionID, result Pointer, shape uint64, args Pointer)

func VariantUnsafeEval(fn FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_eval(fn, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_unsafe_free
func gd_variant_unsafe_free(v1, v2, v3 uint64)

func (v Variant) UnsafeFree() {
	gd_variant_unsafe_free(v[0], v[1], v[2])
}

//go:wasmimport gd variant_unsafe_make_native
func gd_variant_unsafe_make_native(vtype VariantType, v1, v2, v3 uint64, shape uint64, result Pointer)

func VariantUnsafeMakeNative(vtype VariantType, v Variant, shape uint64, result unsafe.Pointer) {
	mem := makeResult(Shape(shape))
	gd_variant_unsafe_make_native(vtype, v[0], v[1], v[2], shape, Pointer(mem))
	loadResult(Shape(shape), result, mem)
}

//go:wasmimport gd variant_unsafe_from_native
func gd_variant_unsafe_from_native(vtype VariantType, result Pointer, shape uint64, args Pointer)

func VariantUnsafeFromNative(vtype VariantType, shape uint64, args unsafe.Pointer) Variant {
	mem_result := makeResult(SizeVariant)
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_from_native(vtype, Pointer(mem_result), shape, Pointer(mem_args))
	var result Variant
	loadResult(SizeVariant, &result, mem_result)
	return result
}

//go:wasmimport gd variant_unsafe_internal_pointer
func gd_variant_unsafe_internal_pointer(vtype VariantType, v1, v2, v3 uint64) Pointer

func VariantUnsafeInternalPointer(vtype VariantType, v Variant) Pointer {
	return gd_variant_unsafe_internal_pointer(vtype, v[0], v[1], v[2])
}

//go:wasmimport gd variant_unsafe_set_field
func gd_variant_unsafe_set_field(setter FunctionID, shape uint64, args Pointer)

func VariantUnsafeSetField(setter FunctionID, shape uint64, args unsafe.Pointer) {
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_set_field(setter, shape, Pointer(mem_args))
}

//go:wasmimport gd variant_unsafe_set_array
func gd_variant_unsafe_set_array(vtype VariantType, idx Int, shape uint64, args Pointer)

func VariantUnsafeSetArray(vtype VariantType, idx Int, shape uint64, args unsafe.Pointer) {
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_set_array(vtype, idx, shape, Pointer(mem_args))
}

//go:wasmimport gd variant_unsafe_set_index
func gd_variant_unsafe_set_index(vtype VariantType, shape uint64, args Pointer)

func VariantUnsafeSetIndex(vtype VariantType, shape uint64, args unsafe.Pointer) {
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_set_index(vtype, shape, Pointer(mem_args))
}

//go:wasmimport gd variant_unsafe_get_field
func gd_variant_unsafe_get_field(getter FunctionID, result Pointer, shape uint64, args Pointer)

func VariantUnsafeGetField(getter FunctionID, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_get_field(getter, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_unsafe_get_array
func gd_variant_unsafe_get_array(vtype VariantType, idx Int, result Pointer, shape uint64, args Pointer)

func VariantUnsafeGetArray(vtype VariantType, idx Int, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_get_array(vtype, idx, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd variant_unsafe_get_index
func gd_variant_unsafe_get_index(vtype VariantType, result Pointer, shape uint64, args Pointer)

func VariantUnsafeGetIndex(vtype VariantType, result unsafe.Pointer, shape uint64, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_variant_unsafe_get_index(vtype, Pointer(mem_result), shape, Pointer(mem_args))
	loadResult(Shape(shape), result, mem_result)
}

// Iterator operations

//go:wasmimport gd iterator_make
func gd_iterator_make(v1, v2, v3 uint64, result Pointer, err Pointer)

func (v Variant) IteratorMake(result unsafe.Pointer, err unsafe.Pointer) {
	mem_result := makeResult(SizeVariant)
	mem_err := makeResult(SizeCallError)
	gd_iterator_make(v[0], v[1], v[2], Pointer(mem_result), Pointer(mem_err))
	loadResult(SizeVariant, result, mem_result)
	loadResult(SizeCallError, err, mem_err)
}

//go:wasmimport gd iterator_next
func gd_iterator_next(v1, v2, v3 uint64, iter Pointer, err Pointer) uint32

func (v Variant) IteratorNext(iter unsafe.Pointer, err unsafe.Pointer) bool {
	mem_iter := makeResult(SizeVariant)
	// Copy iter into engine memory, call, then copy back.
	mem_args := copyArguments(SizeVariant, iter)
	mem_err := makeResult(SizeCallError)
	r := gd_iterator_next(v[0], v[1], v[2], Pointer(mem_args), Pointer(mem_err))
	loadResult(SizeVariant, iter, mem_args)
	_ = mem_iter
	loadResult(SizeCallError, err, mem_err)
	return r != 0
}

//go:wasmimport gd iterator_load
func gd_iterator_load(v1, v2, v3, i1, i2, i3 uint64, result Pointer, err Pointer)

func (v Variant) IteratorLoad(iter Variant, result unsafe.Pointer, err unsafe.Pointer) {
	mem_result := makeResult(SizeVariant)
	mem_err := makeResult(SizeCallError)
	gd_iterator_load(v[0], v[1], v[2], iter[0], iter[1], iter[2], Pointer(mem_result), Pointer(mem_err))
	loadResult(SizeVariant, result, mem_result)
	loadResult(SizeCallError, err, mem_err)
}

// Dictionary operations

//go:wasmimport gd packed_dictionary_access
func gd_packed_dictionary_access(d Dictionary, k1, k2, k3 uint64, result Pointer)

func (d Dictionary) Access(key Variant) Variant {
	mem := makeResult(SizeVariant)
	gd_packed_dictionary_access(d, key[0], key[1], key[2], Pointer(mem))
	var result Variant
	loadResult(SizeVariant, &result, mem)
	return result
}

//go:wasmimport gd packed_dictionary_modify
func gd_packed_dictionary_modify(d Dictionary, k1, k2, k3, v1, v2, v3 uint64)

func (d Dictionary) Modify(key, val Variant) {
	gd_packed_dictionary_modify(d, key[0], key[1], key[2], val[0], val[1], val[2])
}

// RefCounted operations

//go:wasmimport gd ref_get_object
func gd_ref_get_object(ref Pointer) Object

func RefGet(ref Pointer) Object { return gd_ref_get_object(ref) }

//go:wasmimport gd ref_set_object
func gd_ref_set_object(ref Pointer, obj Object)

func RefSet(ref Pointer, obj Object) { gd_ref_set_object(ref, obj) }

// Editor operations

//go:wasmimport gd editor_add_documentation
func EditorAddDocumentation(xml string)

//go:wasmimport gd editor_add_plugin
func EditorAddPlugin(name StringName)

//go:wasmimport gd editor_end_plugin
func EditorEndPlugin(name StringName)

// PropertyList operations

//go:wasmimport gd property_list_make
func MakePropertyList(n Int) PropertyList

//go:wasmimport gd property_list_push
func gd_property_list_push(list PropertyList, vtype VariantType, name StringName, className StringName, hint uint32, hintString String, usage uint32, meta uint32)

func (p PropertyList) Push(vtype VariantType, name StringName, className StringName, hint uint32, hintString String, usage uint32, meta uint32) {
	gd_property_list_push(p, vtype, name, className, hint, hintString, usage, meta)
}

//go:wasmimport gd property_list_free
func gd_property_list_free(list PropertyList)

func (p PropertyList) Free() { gd_property_list_free(p) }

//go:wasmimport gd property_info_type
func gd_property_info_type(info PropertyList) VariantType

func (p PropertyList) InfoType() VariantType { return gd_property_info_type(p) }

//go:wasmimport gd property_info_name
func gd_property_info_name(info PropertyList) StringName

func (p PropertyList) InfoName() StringName { return gd_property_info_name(p) }

//go:wasmimport gd property_info_class_name
func gd_property_info_class_name(info PropertyList) StringName

func (p PropertyList) InfoClassName() StringName { return gd_property_info_class_name(p) }

//go:wasmimport gd property_info_hint
func gd_property_info_hint(info PropertyList) uint32

func (p PropertyList) InfoHint() uint32 { return gd_property_info_hint(p) }

//go:wasmimport gd property_info_hint_string
func gd_property_info_hint_string(info PropertyList) String

func (p PropertyList) InfoHintString() String { return gd_property_info_hint_string(p) }

//go:wasmimport gd property_info_usage
func gd_property_info_usage(info PropertyList) uint32

func (p PropertyList) InfoUsage() uint32 { return gd_property_info_usage(p) }

// MethodList operations

//go:wasmimport gd method_list_make
func gd_method_list_make(n int64) MethodList

func MakeMethodList(n Int) MethodList { return gd_method_list_make(int64(n)) }

//go:wasmimport gd method_list_push
func gd_method_list_push(list MethodList, name StringName, call FunctionID, flags uint32, returnInfo PropertyList, argsInfo PropertyList, count int64, defaults Pointer)

func (m MethodList) Push(name StringName, call ExtensionFunction, flags uint32, returnInfo PropertyList, argsInfo PropertyList, count Int, defaults unsafe.Pointer) {
	var def Pointer
	if defaults != nil {
		def = Pointer(*(*uint32)(defaults))
	}
	gd_method_list_push(m, name, functions.New(call), flags, returnInfo, argsInfo, int64(count), def)
}

//go:wasmimport gd method_list_free
func gd_method_list_free(list MethodList)

func (m MethodList) Free() { gd_method_list_free(m) }

// ClassDB registration

//go:wasmimport gd classdb_register
func gd_classdb_register(class, parent uint32, id uint32, virtual, abstract, exposed, runtime, icon uint32)

func RegisterClass(class, parent StringName, id ExtensionClass, virtual, abstract, exposed, runtime bool, icon String) {
	var vb, ab, eb, rb uint32
	if virtual {
		vb = 1
	}
	if abstract {
		ab = 1
	}
	if exposed {
		eb = 1
	}
	if runtime {
		rb = 1
	}
	gd_classdb_register(uint32(class), uint32(parent), uint32(classes.New(id)), vb, ab, eb, rb, uint32(icon))
}

//go:wasmimport gd classdb_register_methods
func gd_classdb_register_methods(class uint32, methods uint32)

func RegisterMethods(class StringName, methods MethodList) {
	gd_classdb_register_methods(uint32(class), uint32(methods))
}

//go:wasmimport gd classdb_register_constant
func gd_classdb_register_constant(class, enum, name uint32, value int64, bitfield uint32)

func RegisterConstant(class, enum, name StringName, value int64, bitfield bool) {
	var bf uint32
	if bitfield {
		bf = 1
	}
	gd_classdb_register_constant(uint32(class), uint32(enum), uint32(name), value, bf)
}

//go:wasmimport gd classdb_register_property
func gd_classdb_register_property(class uint32, property uint32, setter, getter uint32)

func RegisterProperty(class StringName, property PropertyList, setter, getter StringName) {
	gd_classdb_register_property(uint32(class), uint32(property), uint32(setter), uint32(getter))
}

//go:wasmimport gd classdb_register_property_indexed
func gd_classdb_register_property_indexed(class uint32, property uint32, setter, getter uint32, index int64)

func RegisterPropertyIndexed(class StringName, property PropertyList, setter, getter StringName, index int) {
	gd_classdb_register_property_indexed(uint32(class), uint32(property), uint32(setter), uint32(getter), int64(index))
}

//go:wasmimport gd classdb_register_property_group
func gd_classdb_register_property_group(class, group, prefix uint32)

func RegisterPropertyGroup(class StringName, group, prefix String) {
	gd_classdb_register_property_group(uint32(class), uint32(group), uint32(prefix))
}

//go:wasmimport gd classdb_register_property_sub_group
func gd_classdb_register_property_sub_group(class, subgroup, prefix uint32)

func RegisterPropertySubgroup(class StringName, subgroup, prefix String) {
	gd_classdb_register_property_sub_group(uint32(class), uint32(subgroup), uint32(prefix))
}

//go:wasmimport gd classdb_register_signal
func gd_classdb_register_signal(class, signal uint32, args uint32)

func RegisterSignal(class, signal StringName, args PropertyList) {
	gd_classdb_register_signal(uint32(class), uint32(signal), uint32(args))
}

//go:wasmimport gd classdb_register_removal
func gd_classdb_register_removal(class uint32)

func RegisterRemoval(class StringName) {
	gd_classdb_register_removal(uint32(class))
}

// ClassDB sub-API operations

//go:wasmimport gd classdb_FileAccess_write
func gd_classdb_FileAccess_write(file Object, buf Pointer, length int64)

func FileAccessWrite(file Object, buf []byte) {
	ebuf := copyBufToEngine(buf)
	gd_classdb_FileAccess_write(file, ebuf, int64(len(buf)))
	ebuf.Free()
}

//go:wasmimport gd classdb_FileAccess_read
func gd_classdb_FileAccess_read(file Object, buf Pointer, cap int64) int64

func FileAccessRead(file Object, buf []byte) int {
	ebuf := copyBufToEngine(buf)
	n := int(gd_classdb_FileAccess_read(file, ebuf, int64(len(buf))))
	copyBufToGo(ebuf, buf)
	return n
}

//go:wasmimport gd classdb_Image_unsafe
func gd_classdb_Image_unsafe(img Object) Pointer

func ImageUnsafe(img Object) Pointer { return gd_classdb_Image_unsafe(img) }

//go:wasmimport gd classdb_Image_access
func gd_classdb_Image_access(img Object, offset int64) uint32

func ImageAccess(img Object, offset Int) byte {
	return byte(gd_classdb_Image_access(img, int64(offset)))
}

//go:wasmimport gd classdb_XMLParser_load
func gd_classdb_XMLParser_load(parser Object, buf Pointer, cap int64) int64

func XMLParserLoad(parser Object, buf []byte) int {
	ebuf := copyBufToEngine(buf)
	n := int(gd_classdb_XMLParser_load(parser, ebuf, int64(len(buf))))
	copyBufToGo(ebuf, buf)
	return n
}

//go:wasmimport gd classdb_WorkerThreadPool_add_task
func gd_classdb_WorkerThreadPool_add_task(pool Object, task Pointer, priority uint32, description String)

func WorkerThreadPoolAddTask(pool Object, task Pointer, priority bool, description String) {
	var p uint32
	if priority {
		p = 1
	}
	gd_classdb_WorkerThreadPool_add_task(pool, task, p, description)
}

//go:wasmimport gd classdb_WorkerThreadPool_add_group_task
func gd_classdb_WorkerThreadPool_add_group_task(pool Object, task Pointer, elements int32, arg int32, priority uint32, description String)

func WorkerThreadPoolAddGroupTask(pool Object, task Pointer, elements, arg int32, priority bool, description String) {
	var p uint32
	if priority {
		p = 1
	}
	gd_classdb_WorkerThreadPool_add_group_task(pool, task, elements, arg, p, description)
}

// Extension binding callbacks (no-ops)

//go:wasmexport gd_on_extension_binding_created
func gd_on_extension_binding_created(p0 uint32) uint32 { return 0 }

//go:wasmexport gd_on_extension_binding_removed
func gd_on_extension_binding_removed(p0, p1 uint32) {}

//go:wasmexport gd_on_extension_binding_reference
func gd_on_extension_binding_reference(p0, p1 uint32) uint32 { return 0 }

// Extension class callbacks

//go:wasmexport gd_on_extension_class_create
func gd_on_extension_class_create(p0, p1 uint32) uint32 {
	return uint32(classes.Get(ExtensionClassID(p0)).Create(p1 != 0))
}

//go:wasmexport gd_on_extension_class_method
func gd_on_extension_class_method(p0, p1, p2 uint32) uint32 {
	fn := classes.Get(ExtensionClassID(p0)).Method(StringName(p1), p2)
	if fn == nil {
		return 0
	}
	return uint32(functions.New(fn))
}

//go:wasmexport gd_on_extension_class_caller
func gd_on_extension_class_caller(p0, p1, p2 uint32) uint32 {
	fn := classes.Get(ExtensionClassID(p0)).Method(StringName(p1), p2)
	if fn == nil {
		return 0
	}
	return uint32(functions.New(fn))
}

// Extension instance callbacks

//go:wasmexport gd_on_extension_instance_set
func gd_on_extension_instance_set(p0, p1 uint32, p2, p3, p4 uint64) uint32 {
	if instances.Get(ExtensionInstanceID(p0)).Set(StringName(p1), Variant{p2, p3, p4}) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_get
func gd_on_extension_instance_get(p0, p1, p2 uint32) uint32 {
	v, ok := instances.Get(ExtensionInstanceID(p0)).Get(StringName(p1))
	if ok {
		writeVariant(Pointer(p2), v)
	}
	if ok {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_property_list
func gd_on_extension_instance_property_list(p0 uint32) uint32 {
	return uint32(instances.Get(ExtensionInstanceID(p0)).PropertyList())
}

//go:wasmexport gd_on_extension_instance_property_has_default
func gd_on_extension_instance_property_has_default(p0, p1 uint32) uint32 {
	if instances.Get(ExtensionInstanceID(p0)).HasDefault(StringName(p1)) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_property_get_default
func gd_on_extension_instance_property_get_default(p0, p1, p2 uint32) uint32 {
	v, ok := instances.Get(ExtensionInstanceID(p0)).GetDefault(StringName(p1))
	if ok {
		writeVariant(Pointer(p2), v)
	}
	if ok {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_property_validation
func gd_on_extension_instance_property_validation(p0, p1 uint32) uint32 {
	if instances.Get(ExtensionInstanceID(p0)).ValidateProperty(StringName(p1)) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_notification
func gd_on_extension_instance_notification(p0 uint32, p1 int32, p2 uint32) {
	instances.Get(ExtensionInstanceID(p0)).Notification(p1, p2 != 0)
}

//go:wasmexport gd_on_extension_instance_stringify
func gd_on_extension_instance_stringify(p0 uint32) uint32 {
	return uint32(instances.Get(ExtensionInstanceID(p0)).UnsafeString())
}

//go:wasmexport gd_on_extension_instance_reference
func gd_on_extension_instance_reference(p0, p1 uint32) uint32 {
	if instances.Get(ExtensionInstanceID(p0)).Reference(p1 != 0) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_rid
func gd_on_extension_instance_rid(p0 uint32) uint64 {
	return uint64(instances.Get(ExtensionInstanceID(p0)).RID())
}

//go:wasmexport gd_on_extension_instance_checked_call
func gd_on_extension_instance_checked_call(p0, p1, p2, p3 uint32) {
	var inst ExtensionInstance
	if ExtensionInstanceID(p0) != 0 {
		inst = instances.Get(ExtensionInstanceID(p0))
	}
	functions.Get(FunctionID(p1)).PointerCall(inst, Pointer(p3), Pointer(p2))
}

//go:wasmexport gd_on_extension_instance_called
func gd_on_extension_instance_called(p0, p1, p2, p3 uint32) {
	inst := instances.Get(ExtensionInstanceID(p0))
	functions.Get(FunctionID(p1)).PointerCall(inst, Pointer(p3), Pointer(p2))
}

//go:wasmexport gd_on_extension_instance_variant_call
func gd_on_extension_instance_variant_call(p0, p1, p2, p3 uint32) {
	var inst ExtensionInstance
	if ExtensionInstanceID(p0) != 0 {
		inst = instances.Get(ExtensionInstanceID(p0))
	}
	v := functions.Get(FunctionID(p1)).CheckedCall(inst, VariadicVariants{
		First: PointerTo[PointerTo[Variant]](p3),
	})
	writeVariant(Pointer(p2), v)
}

//go:wasmexport gd_on_extension_instance_dynamic_call
func gd_on_extension_instance_dynamic_call(p0, p1, p2 uint32, p3 int64, p4, p5 uint32) {
	var inst ExtensionInstance
	if ExtensionInstanceID(p0) != 0 {
		inst = instances.Get(ExtensionInstanceID(p0))
	}
	v, err := functions.Get(FunctionID(p1)).DynamicCall(inst, VariadicVariants{
		First: PointerTo[PointerTo[Variant]](p4),
		Count: int(p3),
	})
	writeVariant(Pointer(p2), v)
	writeCallError(Pointer(p5), err)
}

//go:wasmexport gd_on_extension_instance_free
func gd_on_extension_instance_free(p0 uint32) {
	inst := instances.Get(ExtensionInstanceID(p0))
	if f, ok := inst.(interface{ Free() }); ok {
		f.Free()
	}
	instances.Del(ExtensionInstanceID(p0))
}

// Extension script callbacks

//go:wasmexport gd_on_extension_script_categorization
func gd_on_extension_script_categorization(p0, p1 uint32) uint32 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	if script.PropertyCategory() != 0 {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_script_get_property_type
func gd_on_extension_script_get_property_type(p0, p1, p2 uint32) uint32 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		writeCallError(Pointer(p2), CallError{Type: CallInvalidMethod})
		return 0
	}
	return uint32(script.PropertyType(StringName(p1)))
}

//go:wasmexport gd_on_extension_script_get_owner
func gd_on_extension_script_get_owner(p0 uint32) uint32 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.Owner())
}

//go:wasmexport gd_on_extension_script_get_property_state
func gd_on_extension_script_get_property_state(p0, p1, p2 uint32) {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return
	}
	script.ExportedProperties(func(name StringName, value Variant) bool {
		ScriptPropertyStateAdd(FunctionID(p1), Pointer(p2), name, value)
		return true
	})
}

//go:wasmexport gd_on_extension_script_get_methods
func gd_on_extension_script_get_methods(p0 uint32) uint32 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.MethodList())
}

//go:wasmexport gd_on_extension_script_has_method
func gd_on_extension_script_has_method(p0, p1 uint32) uint32 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	if script.HasMethod(StringName(p1)) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_script_get_method_argument_count
func gd_on_extension_script_get_method_argument_count(p0, p1 uint32) int64 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return int64(script.MethodArgumentCount(StringName(p1)))
}

//go:wasmexport gd_on_extension_script_get
func gd_on_extension_script_get(p0 uint32) uint32 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.Script())
}

//go:wasmexport gd_on_extension_script_is_placeholder
func gd_on_extension_script_is_placeholder(p0 uint32) uint32 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	if script.IsPlaceholder() {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_script_get_language
func gd_on_extension_script_get_language(p0 uint32) uint32 {
	script, ok := instances.Get(ExtensionInstanceID(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.ScriptLanguage())
}

// Non-extension callbacks

//go:wasmexport gd_on_engine_init
func gd_on_engine_init(p0 uint32) { onEngineInit(InitializationLevel(p0)) }

//go:wasmexport gd_on_engine_exit
func gd_on_engine_exit(p0 uint32) { onEngineExit(InitializationLevel(p0)) }

//go:wasmexport gd_on_first_frame
func gd_on_first_frame() { onFirstFrame() }

//go:wasmexport gd_on_every_frame
func gd_on_every_frame() { onEveryFrame() }

//go:wasmexport gd_on_final_frame
func gd_on_final_frame() { onFinalFrame() }

//go:wasmexport gd_on_worker_thread_pool_task
func gd_on_worker_thread_pool_task(p0 uint32) { onWorkerThreadPoolTask(TaskID(p0)) }

//go:wasmexport gd_on_worker_thread_pool_group_task
func gd_on_worker_thread_pool_group_task(p0, p1 uint32) {
	onWorkerThreadPoolGroupTask(TaskID(p0), int32(p1))
}

//go:wasmexport gd_on_editor_class_in_use_detection
func gd_on_editor_class_in_use_detection(p0, p1, p2 uint32) {
	if onEditorClassDetection != nil {
		result := onEditorClassDetection(PackedArray[String]{p0, p1})
		Pointer(p2).SetUint32(result[0])
		Pointer(p2 + 4).SetUint32(result[1])
	}
}
