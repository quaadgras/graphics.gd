//go:build wasm

package gdunsafe

import (
	"sync"
	"unsafe"

	"graphics.gd/variant"
)

type Pointer uint32

type MutablePointer uint32

type (
	functionID          uint32
	extensionClassID    uint32
	extensionInstanceID uint32
	callableID          uint32
	scriptInstance      uint32
	taskID              uint32
)

var (
	onWorkerThreadPoolTask      func(taskID)
	onWorkerThreadPoolGroupTask func(taskID, int32)
	onEditorClassDetection      func(PackedArray[String]) PackedArray[String]
)

// Cross-memory helpers for transferring data between Go and engine address spaces.

//go:wasmimport gd memory_malloc
func gd_memory_malloc(int64) uint32

//go:wasmimport gd memory_free
func gd_memory_free(uint32)

//go:wasmimport gd memory_memset
func gd_memory_memset(uint32, int64)

//go:wasmimport gd memory_load_byte
func gd_memory_load_byte(uint32) uint32

//go:wasmimport gd memory_load_u16
func gd_memory_load_u16(uint32) uint32

//go:wasmimport gd memory_load_u32
func gd_memory_load_u32(uint32) uint32

//go:wasmimport gd memory_load_u64
func gd_memory_load_u64(uint32) uint64

//go:wasmimport gd memory_edit_byte
func gd_memory_edit_byte(uint32, uint32)

//go:wasmimport gd memory_edit_u16
func gd_memory_edit_u16(uint32, uint32)

//go:wasmimport gd memory_edit_u32
func gd_memory_edit_u32(uint32, uint32)

//go:wasmimport gd memory_edit_u64
func gd_memory_edit_u64(uint32, uint64)

//go:wasmimport gd memory_edit_128
func gd_memory_edit_128(uint32, uint64, uint64)

func engineLoadByte(ptr uint32) byte         { return byte(gd_memory_load_byte(ptr)) }
func engineLoadU16(ptr uint32) uint16        { return uint16(gd_memory_load_u16(ptr)) }
func engineLoadU32(ptr uint32) uint32        { return gd_memory_load_u32(ptr) }
func engineLoadU64(ptr uint32) uint64        { return gd_memory_load_u64(ptr) }
func engineStoreByte(ptr uint32, v byte)     { gd_memory_edit_byte(ptr, uint32(v)) }
func engineStoreU16(ptr uint32, v uint16)    { gd_memory_edit_u16(ptr, uint32(v)) }
func engineStoreU32(ptr uint32, v uint32)    { gd_memory_edit_u32(ptr, v) }
func engineStoreU64(ptr uint32, v uint64)    { gd_memory_edit_u64(ptr, v) }
func engineStore128(ptr uint32, a, b uint64) { gd_memory_edit_128(ptr, a, b) }

func copyBufToEngine(buf []byte) uint32 {
	ptr := gd_memory_malloc(int64(len(buf)))
	off := uint32(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			engineStoreU64(ptr+off, *(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			engineStoreU32(ptr+off, *(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			engineStoreU16(ptr+off, *(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		default:
			engineStoreByte(ptr+off, buf[0])
			buf = buf[1:]
			off += 1
		}
	}
	return ptr
}

func copyBufToGo(ptr uint32, buf []byte) {
	off := uint32(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 4:
			*(*uint32)(unsafe.Pointer(&buf[0])) = engineLoadU32(ptr + off)
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = engineLoadU16(ptr + off)
			buf = buf[2:]
			off += 2
		default:
			buf[0] = engineLoadByte(ptr + off)
			buf = buf[1:]
			off += 1
		}
	}
	gd_memory_free(ptr)
}

// copyBufToGo2 copies from engine memory without freeing.
func copyBufToGo2(ptr uint32, buf []byte) {
	off := uint32(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 4:
			*(*uint32)(unsafe.Pointer(&buf[0])) = engineLoadU32(ptr + off)
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			*(*uint16)(unsafe.Pointer(&buf[0])) = engineLoadU16(ptr + off)
			buf = buf[2:]
			off += 2
		default:
			buf[0] = engineLoadByte(ptr + off)
			buf = buf[1:]
			off += 1
		}
	}
}

// copyBufToEngine2 copies to engine memory at an existing pointer.
func copyBufToEngine2(ptr uint32, buf []byte) {
	off := uint32(0)
	for len(buf) > 0 {
		switch {
		case len(buf) >= 8:
			engineStoreU64(ptr+off, *(*uint64)(unsafe.Pointer(&buf[0])))
			buf = buf[8:]
			off += 8
		case len(buf) >= 4:
			engineStoreU32(ptr+off, *(*uint32)(unsafe.Pointer(&buf[0])))
			buf = buf[4:]
			off += 4
		case len(buf) >= 2:
			engineStoreU16(ptr+off, *(*uint16)(unsafe.Pointer(&buf[0])))
			buf = buf[2:]
			off += 2
		default:
			engineStoreByte(ptr+off, buf[0])
			buf = buf[1:]
			off += 1
		}
	}
}

var wasmResultBufs [2]uint32
var wasmResultIdx int
var wasmArgBuf uint32
var wasmSelfBuf uint32

var wasmSetup = sync.OnceFunc(func() {
	wasmArgBuf = gd_memory_malloc(64 * 64)
	wasmSelfBuf = gd_memory_malloc(64)
	for i := range wasmResultBufs {
		wasmResultBufs[i] = gd_memory_malloc(64 * 64)
		gd_memory_memset(wasmResultBufs[i], 64*64)
	}
})

func makeResult(shape Shape) uint32 {
	wasmSetup()
	wasmResultIdx ^= 1
	return wasmResultBufs[wasmResultIdx]
}

func loadResult[T ~unsafe.Pointer | *Variant | *Error](shape Shape, result T, from uint32) {
	wasmSetup()
	if from == 0 {
		panic("nil pointer dereference")
	}
	data := unsafe.Pointer(result)
	done := uint32(0)
	size := shape.SizeResult()
	if size == 0 {
		return
	}
	defer gd_memory_memset(from, int64(size))
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(data, done)) = engineLoadU32(from + done)
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(data, done)) = engineLoadU16(from + done)
			done += 2
			size -= 2
		default:
			*(*uint8)(unsafe.Add(data, done)) = engineLoadByte(from + done)
			done += 1
			size -= 1
		}
	}
}

func copyVariants[T ~unsafe.Pointer | *Variant](args T, n int) uint32 {
	wasmSetup()
	var offset uint32
	var data = unsafe.Pointer(args)
	for i := range n {
		pair := *(*[2]uint64)(unsafe.Add(data, uintptr(i*24)))
		engineStore128(wasmArgBuf+offset, pair[0], pair[1])
		engineStoreU64(wasmArgBuf+offset+16, *(*uint64)(unsafe.Add(data, uintptr(i*24+16))))
		offset += 24
	}
	return wasmArgBuf
}

func copySelf(selfShape Shape, self unsafe.Pointer) uint32 {
	wasmSetup()
	if self == nil {
		return 0
	}
	size := selfShape.SizeResult()
	if size == 0 {
		return 0
	}
	buf := unsafe.Slice((*byte)(self), size)
	copyBufToEngine2(wasmSelfBuf, buf)
	return wasmSelfBuf
}

func copyArgumentsTo(shape Shape, args unsafe.Pointer, target uint32) uint32 {
	if args == nil {
		return 0
	}
	bytes := shape.SizeArguments()
	buf := unsafe.Slice((*byte)(args), bytes)
	copyBufToEngine2(target, buf)
	return target
}

func copyArguments(shape Shape, args unsafe.Pointer) uint32 {
	wasmSetup()
	return copyArgumentsTo(shape, args, wasmArgBuf)
}

func writeVariant(addr uint32, v Variant) {
	engineStoreU64(addr, v[0])
	engineStoreU64(addr+8, v[1])
	engineStoreU64(addr+16, v[2])
}

func writeCallError(addr uint32, e Error) {
	engineStoreU32(addr, uint32(e.error))
	engineStoreU32(addr+4, uint32(e.argument))
	engineStoreU32(addr+8, uint32(e.expected))
}

func readVariant(addr uint32) Variant {
	if addr == 0 {
		panic("nil pointer dereference")
	}
	var v Variant
	v[0] = engineLoadU64(addr)
	v[1] = engineLoadU64(addr + 8)
	v[2] = engineLoadU64(addr + 16)
	return v
}

// LibraryLocation returns a string representing the location of the current extension.

//go:wasmimport gd library_location
func gd_library_location() uint32

func LibraryLocation() String { return String(gd_library_location()) }

func (args Variants) Index(i int) Variant {
	if args.count > 0 && (i >= args.count || i < 0) {
		panic("index out of range")
	}
	ptr := engineLoadU32(uint32(args.first) + uint32(i)*4)
	if ptr == 0 {
		return Variant{}
	}
	return readVariant(ptr)
}

// Array

//go:wasmimport gd array_set
func gd_array_set(array uint32, index int64, v1, v2, v3 uint64)

func (array Array) SetIndex(index int, value Variant) {
	gd_array_set(uint32(array), int64(index), value[0], value[1], value[2])
}

//go:wasmimport gd array_get
func gd_array_get(array uint32, index int64, result uint32)

func (array Array) Index(index int) Variant {
	var value Variant
	result := makeResult(ShapeVariant)
	gd_array_get(uint32(array), int64(index), result)
	loadResult(ShapeVariant, &value, result)
	return value
}

//go:wasmimport gd variant_type_setup_array
func gd_variant_type_setup_array(array uint32, vtype uint32, className uint32, v1, v2, v3 uint64)

func (array Array) SetType(t Type) {
	var script Variant
	gd_variant_type_setup_array(uint32(array), uint32(uint32(t.vtype)), uint32(t.class), script[0], script[1], script[2])
}

func (t Type) Size() uintptr { return uintptr(t.shape.SizeResult()) }

// Version

//go:wasmimport gd version_string
func gd_version_string() uint32

func Version() String { return String(gd_version_string()) }

//go:wasmimport gd version_major
func VersionMajor() uint32

//go:wasmimport gd version_minor
func VersionMinor() uint32

//go:wasmimport gd version_patch
func VersionPatch() uint32

//go:wasmimport gd version_hex
func gd_version_hex() uint32

func VersionHexed() uint32 { return gd_version_hex() }

//go:wasmimport gd version_status
func gd_version_status() uint32

func VersionState() String { return String(gd_version_status()) }

//go:wasmimport gd version_build
func gd_version_build() uint32

func VersionBuild() String { return String(gd_version_build()) }

//go:wasmimport gd version_hash
func gd_version_hash() uint32

func VersionCommit() String { return String(gd_version_hash()) }

//go:wasmimport gd gd_version_timestamp
func VersionTimestamp() uint64

// Memory

//go:wasmimport gd memory_resize
func gd_memory_resize(uint32, int64) uint32

func Malloc(size uintptr) MutablePointer {
	return MutablePointer(gd_memory_malloc(int64(size)))
}

func Resize(ptr MutablePointer, size uintptr) MutablePointer {
	return MutablePointer(gd_memory_resize(uint32(ptr), int64(size)))
}

func Memset(ptr MutablePointer, size uintptr, value byte) {
	gd_memory_memset(uint32(ptr), int64(size))
}

func (ptr MutablePointer) Free() { gd_memory_free(uint32(ptr)) }

func (ptr PointerTo[T]) Get() T {
	var v T
	buf := unsafe.Slice((*byte)(unsafe.Pointer(&v)), unsafe.Sizeof(v))
	copyBufToGo2(uint32(ptr), buf)
	return v
}

func (ptr MutablePointerTo[T]) Set(v T) {
	buf := unsafe.Slice((*byte)(unsafe.Pointer(&v)), unsafe.Sizeof(v))
	copyBufToEngine2(uint32(ptr), buf)
}

// String operations

//go:wasmimport gd string_access
func gd_string_access(s uint32, idx int64) int32

func (s String) Index(idx int) rune {
	return rune(gd_string_access(uint32(s), int64(idx)))
}

//go:wasmimport gd string_unsafe
func gd_string_unsafe(s uint32) uint32

func (s String) SetIndex(idx int, char rune) {
	ptr := gd_string_unsafe(uint32(s))
	engineStoreU32(ptr+uint32(idx)*4, uint32(char))
}

//go:wasmimport gd string_resize
func gd_string_resize(s uint32, size int64) uint32

func (s String) Resize(size int) String {
	return String(gd_string_resize(uint32(s), int64(size)))
}

func (s String) Pointer() PointerTo[rune] {
	return PointerTo[rune](gd_string_unsafe(uint32(s)))
}

func (s String) MutablePointer() MutablePointerTo[rune] {
	return MutablePointerTo[rune](gd_string_unsafe(uint32(s)))
}

//go:wasmimport gd string_append
func gd_string_append(s uint32, other uint32) uint32

func (s *String) Append(other String) {
	*s = String(gd_string_append(uint32(*s), uint32(other)))
}

//go:wasmimport gd string_append_rune
func gd_string_append_rune(s uint32, ch int32) uint32

func (s *String) AppendRune(ch rune) {
	*s = String(gd_string_append_rune(uint32(*s), int32(ch)))
}

// Encoding

//go:wasmimport gd string_decode
func gd_string_decode(enc uint32, s uint32, length int64) uint32

//go:wasmimport gd string_encode
func gd_string_encode(enc uint32, s uint32, buf uint32, cap int64) int64

//go:wasmimport gd string_intern
func gd_string_intern(enc uint32, s uint32, length int64) uint32

// Encoding — Latin1

func (enc latin1) Decode(s String, buf []byte) int {
	ebuf := copyBufToEngine(buf)
	n := int(gd_string_encode(0, uint32(s), ebuf, int64(len(buf))))
	copyBufToGo(ebuf, buf)
	return n
}

func (enc latin1) String(s string) String {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := String(gd_string_decode(0, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

func (enc latin1) Intern(s string) StringName {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := StringName(gd_string_intern(0, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

// Encoding — UTF8

func (enc utf8) Decode(s String, buf []byte) int {
	ebuf := copyBufToEngine(buf)
	n := int(gd_string_encode(1, uint32(s), ebuf, int64(len(buf))))
	copyBufToGo(ebuf, buf)
	return n
}

func (enc utf8) String(s string) String {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := String(gd_string_decode(1, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

func (enc utf8) Intern(s string) StringName {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := StringName(gd_string_intern(1, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

// Encoding — UTF16

func (enc utf16) Decode(s String, buf []byte) int {
	ebuf := copyBufToEngine(buf)
	n := int(gd_string_encode(2, uint32(s), ebuf, int64(len(buf))))
	copyBufToGo(ebuf, buf)
	return n
}

func (enc utf16) String(s string) String {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := String(gd_string_decode(2, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

func (enc utf16) Intern(s string) StringName {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := StringName(gd_string_intern(2, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

// Encoding — UTF32

func (enc utf32) Decode(s String, buf []byte) int {
	ebuf := copyBufToEngine(buf)
	n := int(gd_string_encode(4, uint32(s), ebuf, int64(len(buf))))
	copyBufToGo(ebuf, buf)
	return n
}

func (enc utf32) String(s string) String {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := String(gd_string_decode(4, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

func (enc utf32) Intern(s string) StringName {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := StringName(gd_string_intern(4, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

// Encoding — Wide

func (enc wide) Decode(s String, buf []byte) int {
	ebuf := copyBufToEngine(buf)
	n := int(gd_string_encode(5, uint32(s), ebuf, int64(len(buf))))
	copyBufToGo(ebuf, buf)
	return n
}

func (enc wide) String(s string) String {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := String(gd_string_decode(5, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

func (enc wide) Intern(s string) StringName {
	ebuf := copyBufToEngine(unsafe.Slice(unsafe.StringData(s), len(s)))
	result := StringName(gd_string_intern(5, ebuf, int64(len(s))))
	gd_memory_free(ebuf)
	return result
}

// Log

//go:wasmimport gd log
func gd_log(level uint32, text uint32, text_len int32, code uint32, code_len int32, fn uint32, fn_len int32, file uint32, file_len int32, line int32, notify_editor uint32)

func Log(level LogLevel, text, code, fn, file string, line int32, notify_editor bool) {
	etext := copyBufToEngine(unsafe.Slice(unsafe.StringData(text), len(text)))
	ecode := copyBufToEngine(unsafe.Slice(unsafe.StringData(code), len(code)))
	efn := copyBufToEngine(unsafe.Slice(unsafe.StringData(fn), len(fn)))
	efile := copyBufToEngine(unsafe.Slice(unsafe.StringData(file), len(file)))
	var ne uint32
	if notify_editor {
		ne = 1
	}
	gd_log(uint32(level), etext, int32(len(text)), ecode, int32(len(code)), efn, int32(len(fn)), efile, int32(len(file)), line, ne)
	gd_memory_free(etext)
	gd_memory_free(ecode)
	gd_memory_free(efn)
	gd_memory_free(efile)
}

// PackedArray

//go:wasmimport gd packed_array_access
func gd_packed_array_access(t uint32, a1, a2 uint32, idx int64) uint32

//go:wasmimport gd packed_array_modify
func gd_packed_array_modify(t uint32, a1, a2 uint32, idx int64) uint32

func (p PackedArray[T]) Index(idx int64) T {
	ptr := gd_packed_array_access(uint32(p.Type()), p[0], p[1], idx)
	return PointerTo[T](ptr).Get()
}

func (p PackedArray[T]) SetIndex(idx int64, val T) {
	ptr := gd_packed_array_modify(uint32(p.Type()), p[0], p[1], idx)
	MutablePointerTo[T](ptr).Set(val)
}

func (p PackedArray[T]) Pointer() PointerTo[T] {
	return PointerTo[T](gd_packed_array_access(uint32(p.Type()), p[0], p[1], 0))
}

func (p PackedArray[T]) MutablePointer() MutablePointerTo[T] {
	return MutablePointerTo[T](gd_packed_array_modify(uint32(p.Type()), p[0], p[1], 0))
}

// Variant constructors

//go:wasmimport gd variant_type_make
func gd_variant_type_make(t uint32, result uint32, arg_count int64, args uint32, err uint32)

func MakeVariant(vtype variant.Type, args ...Variant) (Variant, Error) {
	param := copyVariants(unsafe.SliceData(args), len(args))
	result := makeResult(ShapeVariant)
	result_err := makeResult(ShapeError)
	gd_variant_type_make(uint32(vtype), result, int64(len(args)), param, result_err)
	var value Variant
	var callErr Error
	loadResult(ShapeVariant, &value, result)
	loadResult(ShapeError, &callErr, result_err)
	return value, callErr
}

//go:wasmimport gd variant_type_call
func gd_variant_type_call(t uint32, method uint32, result uint32, argc int64, args uint32, err uint32)

func Call[T Any](method StringName, args ...Variant) (Variant, Error) {
	param := copyVariants(unsafe.SliceData(args), len(args))
	result := makeResult(ShapeVariant)
	result_err := makeResult(ShapeError)
	gd_variant_type_call(uint32(variantTypeOf[T]()), uint32(method), result, int64(len(args)), param, result_err)
	var value Variant
	var callErr Error
	loadResult(ShapeVariant, &value, result)
	loadResult(ShapeError, &callErr, result_err)
	return value, callErr
}

//go:wasmimport gd variant_type_convertable
func gd_variant_type_convertable(t uint32, to uint32, strict uint32) uint32

func Convertable[A, B Any](strict bool) bool {
	var s uint32
	if strict {
		s = 1
	}
	return gd_variant_type_convertable(uint32(variantTypeOf[A]()), uint32(variantTypeOf[B]()), s) != 0
}

//go:wasmimport gd builtin_name
func gd_builtin_name(utility uint32, hash int64) uint32

//go:wasmimport gd builtin_call
func gd_builtin_call(fn uint32, result uint32, shape Shape, args uint32)

func Utility(utility StringName, hash int64) func(result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	fn := gd_builtin_name(uint32(utility), hash)
	return func(result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
		mem_result := makeResult(Shape(shape))
		mem_args := copyArguments(Shape(shape), args)
		gd_builtin_call(fn, mem_result, shape, mem_args)
		loadResult(Shape(shape), result, mem_result)
	}
}

//go:wasmimport gd variant_type_fetch_constant
func gd_variant_type_fetch_constant(vtype uint32, constant uint32, result uint32)

func Constant[T, E Any](name StringName) E {
	mem := makeResult(ShapeVariant)
	gd_variant_type_fetch_constant(uint32(variantTypeOf[T]()), uint32(name), mem)
	var result E
	loadResult(shapeOf[E](), unsafe.Pointer(&result), mem)
	return result
}

//go:wasmimport gd variant_type_unsafe_constructor
func gd_variant_type_unsafe_constructor(vtype uint32, n int64) uint32

//go:wasmimport gd variant_type_unsafe_make
func gd_variant_type_unsafe_make(constructor uint32, result uint32, shape Shape, args uint32)

func Constructor[T Any](n int) func(shape Shape, args unsafe.Pointer) T {
	fn := gd_variant_type_unsafe_constructor(uint32(variantTypeOf[T]()), int64(n))
	return func(shape Shape, args unsafe.Pointer) T {
		mem_result := makeResult(shape)
		mem_args := copyArguments(shape, args)
		gd_variant_type_unsafe_make(fn, mem_result, shape, mem_args)
		var result T
		loadResult(Shape(shape), unsafe.Pointer(&result), mem_result)
		return result
	}
}

//go:wasmimport gd variant_type_evaluator
func gd_variant_type_evaluator(op uint32, a, b uint32) uint32

//go:wasmimport gd variant_unsafe_eval
func gd_variant_unsafe_eval(fn uint32, result uint32, shape Shape, args uint32)

func Evaluator[A, B, R Any](op VariantOperator) func(a A, b B) R {
	fn := gd_variant_type_evaluator(uint32(op), uint32(variantTypeOf[A]()), uint32(variantTypeOf[B]()))
	shapeA := shapeOf[A]()
	shapeB := shapeOf[B]()
	shape := shapeOf[R]() | shapeA<<4 | shapeB<<8
	return func(a A, b B) R {
		mem_result := makeResult(Shape(shape))
		mem_args := copyArguments(Shape(shape), unsafe.Pointer(&struct {
			A A
			B B
		}{a, b}))
		gd_variant_unsafe_eval(fn, mem_result, shape, mem_args)
		var result R
		loadResult(Shape(shape), unsafe.Pointer(&result), mem_result)
		return result
	}
}

//go:wasmimport gd variant_type_setter
func gd_variant_type_setter(vtype uint32, property uint32) uint32

//go:wasmimport gd variant_unsafe_set_field
func gd_variant_unsafe_set_field(setter uint32, shape Shape, args uint32)

func Setter[T Any, E Any](field StringName) func(v T, val E) {
	fn := gd_variant_type_setter(uint32(variantTypeOf[T]()), uint32(field))
	shapeT := shapeOf[T]()
	shapeE := shapeOf[E]()
	shape := shapeT<<4 | shapeE<<8
	return func(v T, val E) {
		mem_args := copyArguments(Shape(shape), unsafe.Pointer(&struct {
			T T
			E E
		}{v, val}))
		gd_variant_unsafe_set_field(fn, shape, mem_args)
	}
}

//go:wasmimport gd variant_type_getter
func gd_variant_type_getter(vtype uint32, property uint32) uint32

//go:wasmimport gd variant_unsafe_get_field
func gd_variant_unsafe_get_field(getter uint32, result uint32, shape Shape, args uint32)

func Getter[T Any, E Any](field StringName) func(v T) E {
	fn := gd_variant_type_getter(uint32(variantTypeOf[T]()), uint32(field))
	shapeT := shapeOf[T]()
	shape := shapeOf[E]() | shapeT<<4
	return func(v T) E {
		mem_result := makeResult(Shape(shape))
		mem_args := copyArguments(shapeT, unsafe.Pointer(&v))
		gd_variant_unsafe_get_field(fn, mem_result, shape, mem_args)
		var result E
		loadResult(Shape(shape), unsafe.Pointer(&result), mem_result)
		return result
	}
}

//go:wasmimport gd variant_type_has_property
func gd_variant_type_has_property(vtype uint32, property uint32) uint32

func PropertyExists[T Any](property StringName) bool {
	return gd_variant_type_has_property(uint32(variantTypeOf[T]()), uint32(property)) != 0
}

//go:wasmimport gd variant_type_builtin_method
func gd_variant_type_builtin_method(vtype uint32, method uint32, hash int64) uint32

//go:wasmimport gd variant_type_unsafe_call
func gd_variant_type_unsafe_call(self uint32, fn uint32, result uint32, shape Shape, args uint32)

func BuiltinMethod[T Any](method StringName, hash int64) func(self *T, ret unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	fn := gd_variant_type_builtin_method(uint32(variantTypeOf[T]()), uint32(method), hash)
	return func(self *T, ret unsafe.Pointer, shape Shape, args unsafe.Pointer) {
		selfShape := Shape(shape) >> 4
		mem_self := copySelf(selfShape, unsafe.Pointer(self))
		mem_result := makeResult(Shape(shape))
		mem_args := copyArguments(selfShape, args)
		gd_variant_type_unsafe_call(mem_self, fn, mem_result, shape, mem_args)
		// copy self back (may have been mutated)
		if mem_self != 0 {
			size := selfShape.SizeResult()
			buf := unsafe.Slice((*byte)(unsafe.Pointer(self)), size)
			copyBufToGo2(mem_self, buf)
		}
		loadResult(Shape(shape), ret, mem_result)
	}
}

//go:wasmimport gd variant_unsafe_set_array
func gd_variant_unsafe_set_array(vtype uint32, idx int64, shape Shape, args uint32)

func SetIndex[T, V Any](self T, index int64, value V) {
	shapeT := shapeOf[T]()
	shapeV := shapeOf[V]()
	shape := shapeT<<4 | shapeV<<8
	mem_args := copyArguments(Shape(shape), unsafe.Pointer(&struct {
		T T
		V V
	}{self, value}))
	gd_variant_unsafe_set_array(uint32(variantTypeOf[T]()), index, shape, mem_args)
}

//go:wasmimport gd variant_unsafe_get_array
func gd_variant_unsafe_get_array(vtype uint32, idx int64, result uint32, shape Shape, args uint32)

func Index[T, V Any](self T, index int64) V {
	shapeT := shapeOf[T]()
	shape := shapeOf[V]() | shapeT<<4
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(shapeT, unsafe.Pointer(&self))
	gd_variant_unsafe_get_array(uint32(variantTypeOf[T]()), index, mem_result, shape, mem_args)
	var result V
	loadResult(Shape(shape), unsafe.Pointer(&result), mem_result)
	return result
}

//go:wasmimport gd variant_unsafe_set_index
func gd_variant_unsafe_set_index(vtype uint32, shape Shape, args uint32)

func Insert[T Any](self T, index, value Variant) {
	shapeT := shapeOf[T]()
	shape := shapeT<<4 | ShapeVariant<<8 | ShapeVariant<<12
	mem_args := copyArguments(Shape(shape), unsafe.Pointer(&struct {
		T T
		K Variant
		V Variant
	}{self, index, value}))
	gd_variant_unsafe_set_index(uint32(variantTypeOf[T]()), shape, mem_args)
}

//go:wasmimport gd variant_unsafe_get_index
func gd_variant_unsafe_get_index(vtype uint32, result uint32, shape Shape, args uint32)

func Lookup[T Any](self T, key Variant) Variant {
	shapeT := shapeOf[T]()
	shape := ShapeVariant | shapeT<<4 | ShapeVariant<<8
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), unsafe.Pointer(&struct {
		T T
		K Variant
	}{self, key}))
	gd_variant_unsafe_get_index(uint32(variantTypeOf[T]()), mem_result, shape, mem_args)
	var result Variant
	loadResult(Shape(shape), unsafe.Pointer(&result), mem_result)
	return result
}

//go:wasmimport gd variant_type_unsafe_free
func gd_variant_type_unsafe_free(vtype uint32, shape Shape, args uint32)

func Free[T Any](val T) {
	vtype := variantTypeOf[T]()
	shape := shapeOf[T]()
	mem_args := copyArguments(shape, unsafe.Pointer(&val))
	gd_variant_type_unsafe_free(uint32(vtype), shape<<4, mem_args)
}

// Callable

//go:wasmimport gd callable_create
func gd_callable_create(id uint32, object uint64, result uint32)

func MakeCallable(impl ExtensionCallable, obj ObjectID) Callable {
	result := makeResult(ShapeCallable)
	gd_callable_create(uint32(callables.New(impl)), uint64(obj), result)
	var c Callable
	loadResult(ShapeCallable, unsafe.Pointer(&c), result)
	return c
}

// Object

//go:wasmimport gd object_type
func gd_object_type(name uint32) uint32

func (class Class) Tag() ClassTag {
	return ClassTag(gd_object_type(uint32(class)))
}

//go:wasmimport gd object_make
func gd_object_make(name uint32) uint32

func New(name Class) Object {
	return Object(gd_object_make(uint32(name)))
}

//go:wasmimport gd object_name
func gd_object_name(obj uint32) uint32

func (obj Object) Class() Class {
	return Class(gd_object_name(uint32(obj)))
}

//go:wasmimport gd object_cast
func gd_object_cast(obj uint32, to uint32) uint32

func (obj Object) Cast(to ClassTag) Object {
	return Object(gd_object_cast(uint32(obj), uint32(to)))
}

//go:wasmimport gd object_script_fetch
func gd_object_script_fetch(obj uint32, language uint32) uint32

func (obj Object) Script(lang ScriptLanguage) Script {
	return Script(Object(gd_object_script_fetch(uint32(obj), uint32(Object(lang)))))
}

//go:wasmimport gd object_script_setup
func gd_object_script_setup(obj uint32, script uint32)

func (obj Object) AttachScript(script Script) {
	gd_object_script_setup(uint32(obj), uint32(Object(script)))
}

//go:wasmimport gd object_id
func gd_object_id(obj uint32) uint64

func (obj Object) ID() ObjectID {
	return ObjectID(gd_object_id(uint32(obj)))
}

//go:wasmimport gd object_unsafe_free
func gd_object_unsafe_free(obj uint32)

func (obj Object) Free() { gd_object_unsafe_free(uint32(obj)) }

//go:wasmimport gd object_lookup
func gd_object_lookup(id uint64) uint32

func (id ObjectID) Object() Object {
	return Object(gd_object_lookup(uint64(id)))
}

//go:wasmimport gd object_global
func gd_object_global(name uint32) uint32

func Singleton(name StringName) Object {
	return Object(gd_object_global(uint32(name)))
}

//go:wasmimport gd object_method_lookup
func gd_object_method_lookup(class, method uint32, hash int64) uint32

func Method(class, method StringName, hash int64) MethodPointer {
	return MethodPointer(gd_object_method_lookup(uint32(class), uint32(method), hash))
}

//go:wasmimport gd object_call
func gd_object_call(obj uint32, method uint32, result uint32, argc int64, args uint32, err uint32)

func (obj Object) Call(method MethodPointer, args ...Variant) (Variant, Error) {
	mem_result := makeResult(ShapeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(ShapeError)
	gd_object_call(uint32(obj), uint32(method), mem_result, int64(len(args)), mem_args, mem_err)
	var result Variant
	var errResult Error
	loadResult(ShapeVariant, &result, mem_result)
	loadResult(ShapeError, &errResult, mem_err)
	return result, errResult
}

//go:wasmimport gd object_shaped_call
func gd_object_shaped_call(obj uint32, fn uint32, result uint32, shape Shape, args uint32)

func (obj Object) ShapedCall(method MethodPointer, result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	mem_result := makeResult(Shape(shape))
	mem_args := copyArguments(Shape(shape), args)
	gd_object_shaped_call(uint32(obj), uint32(method), mem_result, shape, mem_args)
	loadResult(Shape(shape), result, mem_result)
}

//go:wasmimport gd object_extension_setup
func gd_object_extension_setup(obj uint32, name uint32, inst uint32)

func (obj Object) SetupExtension(name StringName, inst ExtensionInstance) {
	gd_object_extension_setup(uint32(obj), uint32(name), uint32(instances.New(inst)))
}

//go:wasmimport gd object_extension_fetch
func gd_object_extension_fetch(obj uint32) uint32

func (obj Object) ExtensionInstance() ExtensionInstance {
	return instances.Get(uintptr(gd_object_extension_fetch(uint32(obj))))
}

// Script

//go:wasmimport gd object_script_make
func gd_object_script_make(fn uint32) uint32

func MakeScript(fn ExtensionScript) Script {
	return Script(Object(gd_object_script_make(uint32(instances.New(fn)))))
}

//go:wasmimport gd object_script_call
func gd_object_script_call(obj uint32, name uint32, result uint32, argc int64, args uint32, err uint32)

func (obj Script) Call(name StringName, args ...Variant) (Variant, Error) {
	mem_result := makeResult(ShapeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(ShapeError)
	gd_object_script_call(uint32(Object(obj)), uint32(name), mem_result, int64(len(args)), mem_args, mem_err)
	var result Variant
	var err Error
	loadResult(ShapeVariant, &result, mem_result)
	loadResult(ShapeError, &err, mem_err)
	return result, err
}

//go:wasmimport gd object_script_defines_method
func gd_object_script_defines_method(obj uint32, method uint32) uint32

func (obj Script) HasMethod(method StringName) bool {
	return gd_object_script_defines_method(uint32(Object(obj)), uint32(method)) != 0
}

//go:wasmimport gd object_script_placeholder_create
func gd_object_script_placeholder_create(language, script, owner uint32) uint32

func MakePlaceholderScript(language ScriptLanguage, script Script, owner Object) Script {
	return Script(Object(gd_object_script_placeholder_create(uint32(Object(language)), uint32(Object(script)), uint32(owner))))
}

//go:wasmimport gd object_script_placeholder_update
func gd_object_script_placeholder_update(script uint32, array uint32, dict uint32)

func (s Script) UpdatePlaceholder(array Array, dict Dictionary) {
	gd_object_script_placeholder_update(uint32(Object(s)), uint32(array), uint32(dict))
}

// Variant operations

//go:wasmimport gd variant_zero
func gd_variant_zero(result uint32)

func Nil() Variant {
	mem := makeResult(ShapeVariant)
	gd_variant_zero(mem)
	var result Variant
	loadResult(ShapeVariant, &result, mem)
	return result
}

//go:wasmimport gd variant_copy
func gd_variant_copy(v1, v2, v3 uint64, result uint32)

//go:wasmimport gd variant_deep_copy
func gd_variant_deep_copy(v1, v2, v3 uint64, result uint32)

func (v Variant) Copy(deep bool) Variant {
	mem := makeResult(ShapeVariant)
	if deep {
		gd_variant_deep_copy(v[0], v[1], v[2], mem)
	} else {
		gd_variant_copy(v[0], v[1], v[2], mem)
	}
	var result Variant
	loadResult(ShapeVariant, &result, mem)
	return result
}

//go:wasmimport gd variant_call
func gd_variant_call(v1, v2, v3 uint64, method uint32, result uint32, argc int64, args uint32, err uint32)

func (v Variant) Call(method StringName, args ...Variant) (Variant, Error) {
	mem_result := makeResult(ShapeVariant)
	mem_args := copyVariants(unsafe.SliceData(args), len(args))
	mem_err := makeResult(ShapeError)
	gd_variant_call(v[0], v[1], v[2], uint32(method), mem_result, int64(len(args)), mem_args, mem_err)
	var result Variant
	var err Error
	loadResult(ShapeVariant, &result, mem_result)
	loadResult(ShapeError, &err, mem_err)
	return result, err
}

//go:wasmimport gd variant_deep_hash
func gd_variant_deep_hash(v1, v2, v3 uint64, depth int64) int64

func (v Variant) Hash(depth int64) int64 { return gd_variant_deep_hash(v[0], v[1], v[2], depth) }

//go:wasmimport gd variant_bool
func gd_variant_bool(v1, v2, v3 uint64) uint32

func (v Variant) Bool() bool {
	return gd_variant_bool(v[0], v[1], v[2]) != 0
}

//go:wasmimport gd variant_text
func gd_variant_text(v1, v2, v3 uint64) uint32

func (v Variant) UnsafeString() String {
	return String(gd_variant_text(v[0], v[1], v[2]))
}

//go:wasmimport gd variant_type
func gd_variant_type(v1, v2, v3 uint64) uint32

func (v Variant) Type() variant.Type {
	return variant.Type(gd_variant_type(v[0], v[1], v[2]))
}

//go:wasmimport gd object_id_inside_variant
func gd_object_id_inside_variant(v1, v2, v3 uint64) uint64

func (v Variant) ObjectID() ObjectID {
	return ObjectID(gd_object_id_inside_variant(v[0], v[1], v[2]))
}

//go:wasmimport gd variant_get_index
func gd_variant_get_index(v1, v2, v3, k1, k2, k3 uint64, result uint32) uint32

func (v Variant) Lookup(key Variant) (Variant, bool) {
	mem := makeResult(ShapeVariant)
	r := gd_variant_get_index(v[0], v[1], v[2], key[0], key[1], key[2], mem)
	var result Variant
	loadResult(ShapeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_get_array
func gd_variant_get_array(v1, v2, v3 uint64, idx int64, result uint32, err uint32) uint32

func (v Variant) Index(idx int) (Variant, bool, Error) {
	mem_result := makeResult(ShapeVariant)
	mem_err := makeResult(ShapeError)
	r := gd_variant_get_array(v[0], v[1], v[2], int64(idx), mem_result, mem_err)
	var result Variant
	var callErr Error
	loadResult(ShapeVariant, &result, mem_result)
	loadResult(ShapeError, &callErr, mem_err)
	return result, r != 0, callErr
}

//go:wasmimport gd variant_get_field
func gd_variant_get_field(v1, v2, v3 uint64, field uint32, result uint32) uint32

func (v Variant) Field(field StringName) (Variant, bool) {
	mem := makeResult(ShapeVariant)
	r := gd_variant_get_field(v[0], v[1], v[2], uint32(field), mem)
	var result Variant
	loadResult(ShapeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_set_index
func gd_variant_set_index(v1, v2, v3, k1, k2, k3, val1, val2, val3 uint64) uint32

func (v Variant) Insert(key, val Variant) bool {
	return gd_variant_set_index(v[0], v[1], v[2], key[0], key[1], key[2], val[0], val[1], val[2]) != 0
}

//go:wasmimport gd variant_set_array
func gd_variant_set_array(v1, v2, v3 uint64, idx int64, val1, val2, val3 uint64, err uint32) uint32

func (v Variant) SetIndex(idx int, val Variant) (bool, Error) {
	mem_err := makeResult(ShapeError)
	r := gd_variant_set_array(v[0], v[1], v[2], int64(idx), val[0], val[1], val[2], mem_err)
	var callErr Error
	loadResult(ShapeError, &callErr, mem_err)
	return r != 0, callErr
}

//go:wasmimport gd variant_set_field
func gd_variant_set_field(v1, v2, v3 uint64, field uint32, val1, val2, val3 uint64) uint32

func (v Variant) SetField(field StringName, value Variant) bool {
	return gd_variant_set_field(v[0], v[1], v[2], uint32(field), value[0], value[1], value[2]) != 0
}

//go:wasmimport gd variant_has_index
func gd_variant_has_index(v1, v2, v3, i1, i2, i3 uint64) uint32

func (v Variant) Has(index Variant) bool {
	return gd_variant_has_index(v[0], v[1], v[2], index[0], index[1], index[2]) != 0
}

//go:wasmimport gd variant_has_method
func gd_variant_has_method(v1, v2, v3 uint64, method uint32) uint32

func (v Variant) HasMethod(method StringName) bool {
	return gd_variant_has_method(v[0], v[1], v[2], uint32(method)) != 0
}

//go:wasmimport gd variant_unsafe_free
func gd_variant_unsafe_free(v1, v2, v3 uint64)

func (v Variant) Free() {
	gd_variant_unsafe_free(v[0], v[1], v[2])
}

//go:wasmimport gd variant_eval
func gd_variant_eval(op uint32, a1, a2, a3, b1, b2, b3 uint64, result uint32) uint32

func (op VariantOperator) Evaluate(a, b Variant) (Variant, bool) {
	mem := makeResult(ShapeVariant)
	r := gd_variant_eval(uint32(op), a[0], a[1], a[2], b[0], b[1], b[2], mem)
	var result Variant
	loadResult(ShapeVariant, &result, mem)
	return result, r != 0
}

//go:wasmimport gd variant_unsafe_make_native
func gd_variant_unsafe_make_native(vtype uint32, v1, v2, v3 uint64, shape Shape, result uint32)

func VariantInto[T Any](v Variant) T {
	vtype := variantTypeOf[T]()
	shape := shapeOf[T]()
	mem := makeResult(shape)
	gd_variant_unsafe_make_native(uint32(vtype), v[0], v[1], v[2], shape<<4, mem)
	var result T
	loadResult(shape, unsafe.Pointer(&result), mem)
	return result
}

//go:wasmimport gd variant_unsafe_from_native
func gd_variant_unsafe_from_native(vtype uint32, result uint32, shape Shape, args uint32)

func VariantFrom[T Any](native T) Variant {
	vtype := variantTypeOf[T]()
	shape := shapeOf[T]()
	mem_result := makeResult(ShapeVariant)
	mem_args := copyArguments(shape<<4, unsafe.Pointer(&native))
	gd_variant_unsafe_from_native(uint32(vtype), mem_result, shape<<4, mem_args)
	var result Variant
	loadResult(ShapeVariant, &result, mem_result)
	return result
}

//go:wasmimport gd variant_unsafe_internal_pointer
func gd_variant_unsafe_internal_pointer(vtype uint32, v1, v2, v3 uint64) uint32

func PointerIntoVariant[T Any](v Variant) PointerTo[T] {
	return PointerTo[T](gd_variant_unsafe_internal_pointer(uint32(variantTypeOf[T]()), v[0], v[1], v[2]))
}

// Dictionary

//go:wasmimport gd packed_dictionary_access
func gd_packed_dictionary_access(d uint32, k1, k2, k3 uint64, result uint32)

func (d Dictionary) Lookup(key Variant) Variant {
	mem := makeResult(ShapeVariant)
	gd_packed_dictionary_access(uint32(d), key[0], key[1], key[2], mem)
	var result Variant
	loadResult(ShapeVariant, &result, mem)
	return result
}

//go:wasmimport gd packed_dictionary_modify
func gd_packed_dictionary_modify(d uint32, k1, k2, k3, v1, v2, v3 uint64)

func (d Dictionary) Insert(key, val Variant) {
	gd_packed_dictionary_modify(uint32(d), key[0], key[1], key[2], val[0], val[1], val[2])
}

//go:wasmimport gd variant_type_setup_dictionary
func gd_variant_type_setup_dictionary(dict uint32, keyType uint32, keyClassName uint32, ks1, ks2, ks3 uint64, valType uint32, valClassName uint32, vs1, vs2, vs3 uint64)

func (dict Dictionary) SetType(key, val Type) {
	var keyScript, valScript Variant
	gd_variant_type_setup_dictionary(uint32(dict),
		uint32(key.vtype), uint32(key.class),
		keyScript[0], keyScript[1], keyScript[2],
		uint32(val.vtype), uint32(val.class),
		valScript[0], valScript[1], valScript[2])
}

// RefCounted

//go:wasmimport gd ref_get_object
func gd_ref_get_object(ref uint32) uint32

func (ref RefCounted) Get() Object {
	return Object(gd_ref_get_object(uint32(ref)))
}

//go:wasmimport gd ref_set_object
func gd_ref_set_object(ref uint32, obj uint32)

func (ref RefCounted) Set(obj Object) {
	gd_ref_set_object(uint32(ref), uint32(obj))
}

// Editor

//go:wasmimport gd editor_add_documentation
func EditorDocumentation(xml string)

//go:wasmimport gd editor_add_plugin
func gd_editor_add_plugin(name uint32)

func EnableEditorPlugin(name Class) {
	gd_editor_add_plugin(uint32(name))
}

//go:wasmimport gd editor_end_plugin
func gd_editor_end_plugin(name uint32)

func RemoveEditorPlugin(name Class) {
	gd_editor_end_plugin(uint32(name))
}

// PropertyList

//go:wasmimport gd property_list_make
func gd_property_list_make(n int64) uint32

func MakePropertyList(n int64) PropertyList {
	return PropertyList(gd_property_list_make(n))
}

//go:wasmimport gd property_list_push
func gd_property_list_push(list uint32, vtype uint32, name uint32, className uint32, hint uint32, hintString uint32, usage uint32, meta uint32)

func (p PropertyList) Push(prop Property) {
	gd_property_list_push(uint32(p), uint32(prop.Type), uint32(prop.Name), uint32(prop.ClassName), prop.Hint, uint32(prop.HintString), prop.Usage, 0)
}

//go:wasmimport gd property_list_free
func gd_property_list_free(list uint32)

func (p PropertyList) Free() { gd_property_list_free(uint32(p)) }

// MethodList

//go:wasmimport gd method_list_make
func gd_method_list_make(n int64) uint32

func MakeMethodList(n int64) MethodList {
	return MethodList(gd_method_list_make(n))
}

//go:wasmimport gd method_list_push
func gd_method_list_push(list uint32, name uint32, call uint32, flags uint32, returnInfo uint32, argsInfo uint32, count int64, defaults uint32)

func (m MethodList) Push(name StringName, call ExtensionFunction, flags uint32, returnInfo PropertyList, argsInfo PropertyList, count int64, defaults unsafe.Pointer) {
	var def uint32
	if defaults != nil {
		def = *(*uint32)(defaults)
	}
	gd_method_list_push(uint32(m), uint32(name), uint32(functions.New(call)), flags, uint32(returnInfo), uint32(argsInfo), count, def)
}

//go:wasmimport gd method_list_free
func gd_method_list_free(list uint32)

func (m MethodList) Free() { gd_method_list_free(uint32(m)) }

// ClassDB registration

//go:wasmimport gd classdb_register
func gd_classdb_register(class, parent uint32, id uint32, virtual, abstract, exposed, runtime, icon uint32)

func RegisterClass(class string, id ExtensionClass) Class {
	class_name := UTF8.Intern(class)
	parent := id.Parent()
	var vb, ab, eb, rb uint32
	if id.Virtual() {
		vb = 1
	}
	if id.Abstract() {
		ab = 1
	}
	if id.Exposed() {
		eb = 1
	}
	if id.Runtime() {
		rb = 1
	}
	gd_classdb_register(uint32(class_name), uint32(parent), uint32(classes.New(id)), vb, ab, eb, rb, uint32(id.Icon()))
	return Class(class_name)
}

//go:wasmimport gd classdb_register_methods
func gd_classdb_register_methods(class uint32, methods uint32)

func (class Class) RegisterMethods(methods MethodList) {
	gd_classdb_register_methods(uint32(class), uint32(methods))
}

//go:wasmimport gd classdb_register_constant
func gd_classdb_register_constant(class, enum, name uint32, value int64, bitfield uint32)

func (class Class) RegisterConstant(enum, name StringName, value int64, bitfield bool) {
	var bf uint32
	if bitfield {
		bf = 1
	}
	gd_classdb_register_constant(uint32(class), uint32(enum), uint32(name), value, bf)
}

//go:wasmimport gd classdb_register_property
func gd_classdb_register_property(class uint32, property uint32, setter, getter uint32)

func (class Class) RegisterProperty(property Property, setter, getter StringName) {
	pl := MakePropertyList(1)
	pl.Push(property)
	gd_classdb_register_property(uint32(class), uint32(pl), uint32(setter), uint32(getter))
}

//go:wasmimport gd classdb_register_property_indexed
func gd_classdb_register_property_indexed(class uint32, property uint32, setter, getter uint32, index int64)

func (class Class) RegisterPropertyIndexed(property Property, setter, getter String, index int) {
	pl := MakePropertyList(1)
	pl.Push(property)
	gd_classdb_register_property_indexed(uint32(class), uint32(pl), uint32(setter), uint32(getter), int64(index))
}

//go:wasmimport gd classdb_register_property_group
func gd_classdb_register_property_group(class, group, prefix uint32)

func (class Class) RegisterPropertyGroup(group, prefix String) {
	gd_classdb_register_property_group(uint32(class), uint32(group), uint32(prefix))
}

//go:wasmimport gd classdb_register_property_sub_group
func gd_classdb_register_property_sub_group(class, subgroup, prefix uint32)

func (class Class) RegisterPropertySubgroup(subgroup, prefix String) {
	gd_classdb_register_property_sub_group(uint32(class), uint32(subgroup), uint32(prefix))
}

//go:wasmimport gd classdb_register_signal
func gd_classdb_register_signal(class, signal uint32, args uint32)

func (class Class) RegisterSignal(signal StringName, args PropertyList) {
	gd_classdb_register_signal(uint32(class), uint32(signal), uint32(args))
}

//go:wasmimport gd classdb_register_removal
func gd_classdb_register_removal(class uint32)

func (class Class) Free() {
	gd_classdb_register_removal(uint32(class))
}

// Iterator

//go:wasmimport gd iterator_make
func gd_iterator_make(v1, v2, v3 uint64, result uint32, err uint32)

func (v Variant) Iterator() (Iterator, Error) {
	mem_result := makeResult(ShapeVariant)
	mem_err := makeResult(ShapeError)
	gd_iterator_make(v[0], v[1], v[2], mem_result, mem_err)
	var iter Variant
	var callErr Error
	loadResult(ShapeVariant, &iter, mem_result)
	loadResult(ShapeError, &callErr, mem_err)
	return Iterator(iter), callErr
}

//go:wasmimport gd iterator_next
func gd_iterator_next(v1, v2, v3 uint64, iter uint32, err uint32) uint32

func (iter *Iterator) Next() (bool, Error) {
	v := Variant(*iter)
	mem_iter := copyArguments(ShapeVariant, unsafe.Pointer(iter))
	mem_err := makeResult(ShapeError)
	r := gd_iterator_next(v[0], v[1], v[2], mem_iter, mem_err)
	loadResult(ShapeVariant, unsafe.Pointer(iter), mem_iter)
	var callErr Error
	loadResult(ShapeError, &callErr, mem_err)
	return r != 0, callErr
}

//go:wasmimport gd iterator_load
func gd_iterator_load(v1, v2, v3, i1, i2, i3 uint64, result uint32, err uint32)

func (iter Iterator) Value() (Variant, Error) {
	v := Variant(iter)
	mem_result := makeResult(ShapeVariant)
	mem_err := makeResult(ShapeError)
	gd_iterator_load(v[0], v[1], v[2], v[0], v[1], v[2], mem_result, mem_err)
	var result Variant
	var callErr Error
	loadResult(ShapeVariant, &result, mem_result)
	loadResult(ShapeError, &callErr, mem_err)
	return result, callErr
}

// Callable callbacks

//go:wasmexport gd_on_callable_called
func gd_on_callable_called(c uint32, ret uint32, argc int64, args uint32, err uint32) {
	r, e := callables.Get(uintptr(c)).Call(Variants{first: PointerTo[PointerTo[Variant]](args), count: int(argc)})
	writeVariant(ret, r)
	writeCallError(err, e)
}

//go:wasmexport gd_on_callable_verify
func gd_on_callable_verify(c uint32) uint32 {
	if callables.Get(uintptr(c)).IsValid() {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_callable_delete
func gd_on_callable_delete(c uint32) { callables.Del(uintptr(c)) }

//go:wasmexport gd_on_callable_hashed
func gd_on_callable_hashed(c uint32) uint32 {
	return callables.Get(uintptr(c)).Hash()
}

//go:wasmexport gd_on_callable_sorted
func gd_on_callable_sorted(a, b uint32) int64 {
	return int64(callables.Get(uintptr(a)).Compare(callables.Get(uintptr(b))))
}

//go:wasmexport gd_on_callable_string
func gd_on_callable_string(c uint32) uint32 {
	return uint32(callables.Get(uintptr(c)).UnsafeString())
}

//go:wasmexport gd_on_callable_length
func gd_on_callable_length(c uint32) int64 {
	return int64(callables.Get(uintptr(c)).ArgumentCount())
}

// Extension binding callbacks

//go:wasmexport gd_on_extension_binding_created
func gd_on_extension_binding_created(p0 uint32) uint32 { return 0 }

//go:wasmexport gd_on_extension_binding_removed
func gd_on_extension_binding_removed(p0, p1 uint32) {}

//go:wasmexport gd_on_extension_binding_reference
func gd_on_extension_binding_reference(p0, p1 uint32) uint32 { return 0 }

// Extension class callbacks

//go:wasmexport gd_on_extension_class_create
func gd_on_extension_class_create(p0, p1 uint32) uint32 {
	return uint32(classes.Get(uintptr(p0)).Create(p1 != 0))
}

//go:wasmexport gd_on_extension_class_method
func gd_on_extension_class_method(p0, p1, p2 uint32) uint32 {
	fn := classes.Get(uintptr(p0)).Method(StringName(p1), p2)
	if fn == nil {
		return 0
	}
	return uint32(functions.New(fn))
}

//go:wasmexport gd_on_extension_class_caller
func gd_on_extension_class_caller(p0, p1, p2 uint32) uint32 {
	fn := classes.Get(uintptr(p0)).Method(StringName(p1), p2)
	if fn == nil {
		return 0
	}
	return uint32(functions.New(fn))
}

// Extension instance callbacks

//go:wasmexport gd_on_extension_instance_set
func gd_on_extension_instance_set(p0, p1 uint32, p2, p3, p4 uint64) uint32 {
	if instances.Get(uintptr(p0)).Set(StringName(p1), Variant{p2, p3, p4}) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_get
func gd_on_extension_instance_get(p0, p1, p2 uint32) uint32 {
	v, ok := instances.Get(uintptr(p0)).Get(StringName(p1))
	if ok {
		writeVariant(p2, v)
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_property_list
func gd_on_extension_instance_property_list(p0 uint32) uint32 {
	return uint32(instances.Get(uintptr(p0)).PropertyList())
}

//go:wasmexport gd_on_extension_instance_property_has_default
func gd_on_extension_instance_property_has_default(p0, p1 uint32) uint32 {
	if instances.Get(uintptr(p0)).HasDefault(StringName(p1)) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_property_get_default
func gd_on_extension_instance_property_get_default(p0, p1, p2 uint32) uint32 {
	v, ok := instances.Get(uintptr(p0)).GetDefault(StringName(p1))
	if ok {
		writeVariant(p2, v)
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_property_validation
func gd_on_extension_instance_property_validation(p0, p1 uint32) uint32 {
	if instances.Get(uintptr(p0)).ValidateProperty(Property{Name: StringName(p1)}) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_notification
func gd_on_extension_instance_notification(p0 uint32, p1 int32, p2 uint32) {
	instances.Get(uintptr(p0)).Notification(p1, p2 != 0)
}

//go:wasmexport gd_on_extension_instance_stringify
func gd_on_extension_instance_stringify(p0 uint32) uint32 {
	return uint32(instances.Get(uintptr(p0)).UnsafeString())
}

//go:wasmexport gd_on_extension_instance_reference
func gd_on_extension_instance_reference(p0, p1 uint32) uint32 {
	if instances.Get(uintptr(p0)).Reference(p1 != 0) {
		return 1
	}
	return 0
}

//go:wasmexport gd_on_extension_instance_rid
func gd_on_extension_instance_rid(p0 uint32) uint64 {
	return uint64(instances.Get(uintptr(p0)).RID())
}

//go:wasmexport gd_on_extension_instance_checked_call
func gd_on_extension_instance_checked_call(p0, p1, p2, p3 uint32) {
	var inst ExtensionInstance
	if p0 != 0 {
		inst = instances.Get(uintptr(p0))
	}
	functions.Get(uintptr(p1)).PointerCall(inst, MutablePointer(p3), PointerTo[Pointer](p2))
}

//go:wasmexport gd_on_extension_instance_called
func gd_on_extension_instance_called(p0, p1, p2, p3 uint32) {
	inst := instances.Get(uintptr(p0))
	functions.Get(uintptr(p1)).PointerCall(inst, MutablePointer(p3), PointerTo[Pointer](p2))
}

//go:wasmexport gd_on_extension_instance_variant_call
func gd_on_extension_instance_variant_call(p0, p1, p2, p3 uint32) {
	var inst ExtensionInstance
	if p0 != 0 {
		inst = instances.Get(uintptr(p0))
	}
	v := functions.Get(uintptr(p1)).CheckedCall(inst, Variants{
		first: PointerTo[PointerTo[Variant]](p3),
		count: -1,
	})
	writeVariant(p2, v)
}

//go:wasmexport gd_on_extension_instance_dynamic_call
func gd_on_extension_instance_dynamic_call(p0, p1, p2 uint32, p3 int64, p4, p5 uint32) {
	var inst ExtensionInstance
	if p0 != 0 {
		inst = instances.Get(uintptr(p0))
	}
	v, err := functions.Get(uintptr(p1)).DynamicCall(inst, Variants{
		first: PointerTo[PointerTo[Variant]](p4),
		count: int(p3),
	})
	writeVariant(p2, v)
	writeCallError(p5, err)
}

//go:wasmexport gd_on_extension_instance_free
func gd_on_extension_instance_free(p0 uint32) {
	inst := instances.Get(uintptr(p0))
	if f, ok := inst.(interface{ Free() }); ok {
		f.Free()
	}
	instances.Del(uintptr(p0))
}

// Extension script callbacks

//go:wasmimport gd object_script_property_state_add
func gd_object_script_property_state_add(fn uint32, arg uint32, name uint32, s1, s2, s3 uint64)

//go:wasmexport gd_on_extension_script_categorization
func gd_on_extension_script_categorization(p0, p1 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
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
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		writeCallError(p2, Error{error: errorInvalidMethod})
		return 0
	}
	return uint32(script.PropertyType(StringName(p1)))
}

//go:wasmexport gd_on_extension_script_get_owner
func gd_on_extension_script_get_owner(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.Owner())
}

//go:wasmexport gd_on_extension_script_get_property_state
func gd_on_extension_script_get_property_state(p0, p1, p2 uint32) {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return
	}
	script.ExportedProperties(func(name StringName, value Variant) bool {
		gd_object_script_property_state_add(p1, p2, uint32(name), value[0], value[1], value[2])
		return true
	})
}

//go:wasmexport gd_on_extension_script_get_methods
func gd_on_extension_script_get_methods(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.MethodList())
}

//go:wasmexport gd_on_extension_script_has_method
func gd_on_extension_script_has_method(p0, p1 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
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
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return int64(script.MethodArgumentCount(StringName(p1)))
}

//go:wasmexport gd_on_extension_script_get
func gd_on_extension_script_get(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.Script())
}

//go:wasmexport gd_on_extension_script_is_placeholder
func gd_on_extension_script_is_placeholder(p0 uint32) uint32 {
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
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
	script, ok := instances.Get(uintptr(p0)).(ExtensionScript)
	if !ok {
		return 0
	}
	return uint32(script.ScriptLanguage())
}

// Lifecycle callbacks

//go:wasmexport gd_on_engine_init
func gd_on_engine_init(p0 uint32) {
	level := InitializationLevel(p0)
	for _, fn := range onEngineInit {
		fn(level)
	}
}

//go:wasmexport gd_on_engine_exit
func gd_on_engine_exit(p0 uint32) {
	level := InitializationLevel(p0)
	for _, fn := range onEngineExit {
		fn(level)
	}
}

//go:wasmexport gd_on_first_frame
func gd_on_first_frame() {
	for _, fn := range onFirstFrame {
		fn()
	}
}

//go:wasmexport gd_on_every_frame
func gd_on_every_frame() {
	for _, fn := range onEveryFrame {
		fn()
	}
}

//go:wasmexport gd_on_final_frame
func gd_on_final_frame() {
	for _, fn := range onFinalFrame {
		fn()
	}
}

//go:wasmexport gd_on_worker_thread_pool_task
func gd_on_worker_thread_pool_task(p0 uint32) {
	if onWorkerThreadPoolTask != nil {
		onWorkerThreadPoolTask(taskID(p0))
	}
}

//go:wasmexport gd_on_worker_thread_pool_group_task
func gd_on_worker_thread_pool_group_task(p0, p1 uint32) {
	if onWorkerThreadPoolGroupTask != nil {
		onWorkerThreadPoolGroupTask(taskID(p0), int32(p1))
	}
}

//go:wasmexport gd_on_editor_class_in_use_detection
func gd_on_editor_class_in_use_detection(p0, p1, p2 uint32) {
	if onEditorClassDetection != nil {
		result := onEditorClassDetection(PackedArray[String]{p0, p1})
		engineStoreU32(p2, result[0])
		engineStoreU32(p2+4, result[1])
	}
}
