//go:build wasm

package gdunsafe

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
)

type (
	String     uint32
	StringName uint32
	Pointer    uint32
	Array      uint32

	VariantType uint32
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
	result := gdmemory.MakeResult(gdextension.SizeVariant)
	gd_array_get(array, index, uint32(result))
	gdmemory.LoadResult(gdextension.SizeVariant, &value, result)
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
	var param = gdmemory.CopyVariants(unsafe.SliceData(args), len(args))
	result := gdmemory.MakeResult(gdextension.SizeVariant)
	result_err := gdmemory.MakeResult(gdextension.SizeCallError)
	gd_variant_type_make(t, Pointer(result), Int(len(args)), Pointer(param), Pointer(result_err))
	gdmemory.LoadResult(gdextension.SizeVariant, &value, result)
	gdmemory.LoadResult(gdextension.SizeCallError, &value, result_err)
	return value, err
}
