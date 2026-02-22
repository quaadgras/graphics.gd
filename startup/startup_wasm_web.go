//go:build js

package startup

import (
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
)

//go:wasmimport gd memory_malloc
func wasm_gd_memory_malloc(size uint32) uint32

//go:wasmimport gd memory_resize
func wasm_gd_memory_resize(addr uint32, size uint32) uint32

//go:wasmimport gd memory_clear
func wasm_gd_memory_clear(addr uint32, size uint32)

//go:wasmimport gd memory_free
func wasm_gd_memory_free(addr uint32)

//go:wasmimport gd memory_load_byte
func wasm_gd_memory_load_byte(addr uint32) uint32

//go:wasmimport gd memory_load_u16
func wasm_gd_memory_load_u16(addr uint32) uint32

//go:wasmimport gd memory_load_u32
func wasm_gd_memory_load_u32(addr uint32) uint32

//go:wasmimport gd memory_edit_byte
func wasm_gd_memory_edit_byte(addr uint32, value uint32)

//go:wasmimport gd memory_edit_u16
func wasm_gd_memory_edit_u16(addr uint32, value uint32)

//go:wasmimport gd memory_edit_u32
func wasm_gd_memory_edit_u32(addr uint32, value uint32)

//go:wasmimport gd memory_edit_u64
func wasm_gd_memory_edit_u64(addr uint32, value_hi uint32, value_lo uint32)

//go:wasmimport gd memory_edit_128
func wasm_gd_memory_edit_128(addr uint32, a_hi uint32, a_lo uint32, b_hi uint32, b_lo uint32)

//go:wasmimport gd memory_edit_256
func wasm_gd_memory_edit_256(addr uint32, a_hi uint32, a_lo uint32, b_hi uint32, b_lo uint32, c_hi uint32, c_lo uint32, d_hi uint32, d_lo uint32)

//go:wasmimport gd memory_edit_512
func wasm_gd_memory_edit_512(addr uint32, a_hi uint32, a_lo uint32, b_hi uint32, b_lo uint32, c_hi uint32, c_lo uint32, d_hi uint32, d_lo uint32, e_hi uint32, e_lo uint32, f_hi uint32, f_lo uint32, g_hi uint32, g_lo uint32, h_hi uint32, h_lo uint32)

//go:wasmimport gd object_unsafe_call
func wasm_gd_object_unsafe_call(obj uint32, method uint32, result uint32, shape_hi uint32, shape_lo uint32, args uint32)

func init() {
	// Override the syscall/js-based implementations set by startup_js_v2.go.
	// This init() runs after startup_js_v2.go's init() because "startup_wasm_web"
	// sorts alphabetically after "startup_js_v2".
	gdextension.Host.Memory.Malloc = func(size int) gdextension.Pointer {
		return gdextension.Pointer(wasm_gd_memory_malloc(uint32(size)))
	}
	gdextension.Host.Memory.Resize = func(addr gdextension.Pointer, size int) gdextension.Pointer {
		return gdextension.Pointer(wasm_gd_memory_resize(uint32(addr), uint32(size)))
	}
	gdextension.Host.Memory.Clear = func(addr gdextension.Pointer, size int) {
		wasm_gd_memory_clear(uint32(addr), uint32(size))
	}
	gdextension.Host.Memory.Free = func(addr gdextension.Pointer) {
		wasm_gd_memory_free(uint32(addr))
	}
	gdextension.Host.Memory.Load.Byte = func(addr gdextension.Pointer) byte {
		return byte(wasm_gd_memory_load_byte(uint32(addr)))
	}
	gdextension.Host.Memory.Load.Uint16 = func(addr gdextension.Pointer) uint16 {
		return uint16(wasm_gd_memory_load_u16(uint32(addr)))
	}
	gdextension.Host.Memory.Load.Uint32 = func(addr gdextension.Pointer) uint32 {
		return wasm_gd_memory_load_u32(uint32(addr))
	}
	gdextension.Host.Memory.Edit.Byte = func(addr gdextension.Pointer, value byte) {
		wasm_gd_memory_edit_byte(uint32(addr), uint32(value))
	}
	gdextension.Host.Memory.Edit.Uint16 = func(addr gdextension.Pointer, value uint16) {
		wasm_gd_memory_edit_u16(uint32(addr), uint32(value))
	}
	gdextension.Host.Memory.Edit.Uint32 = func(addr gdextension.Pointer, value uint32) {
		wasm_gd_memory_edit_u32(uint32(addr), value)
	}
	gdextension.Host.Memory.Edit.Uint64 = func(addr gdextension.Pointer, value uint64) {
		wasm_gd_memory_edit_u64(uint32(addr), uint32(value>>32), uint32(value&0xFFFFFFFF))
	}
	gdextension.Host.Memory.Edit.Bits128 = func(addr gdextension.Pointer, value [2]uint64) {
		wasm_gd_memory_edit_128(uint32(addr),
			uint32(value[0]>>32), uint32(value[0]&0xFFFFFFFF),
			uint32(value[1]>>32), uint32(value[1]&0xFFFFFFFF))
	}
	gdextension.Host.Memory.Edit.Bits256 = func(addr gdextension.Pointer, value [4]uint64) {
		wasm_gd_memory_edit_256(uint32(addr),
			uint32(value[0]>>32), uint32(value[0]&0xFFFFFFFF),
			uint32(value[1]>>32), uint32(value[1]&0xFFFFFFFF),
			uint32(value[2]>>32), uint32(value[2]&0xFFFFFFFF),
			uint32(value[3]>>32), uint32(value[3]&0xFFFFFFFF))
	}
	gdextension.Host.Memory.Edit.Bits512 = func(addr gdextension.Pointer, value [8]uint64) {
		wasm_gd_memory_edit_512(uint32(addr),
			uint32(value[0]>>32), uint32(value[0]&0xFFFFFFFF),
			uint32(value[1]>>32), uint32(value[1]&0xFFFFFFFF),
			uint32(value[2]>>32), uint32(value[2]&0xFFFFFFFF),
			uint32(value[3]>>32), uint32(value[3]&0xFFFFFFFF),
			uint32(value[4]>>32), uint32(value[4]&0xFFFFFFFF),
			uint32(value[5]>>32), uint32(value[5]&0xFFFFFFFF),
			uint32(value[6]>>32), uint32(value[6]&0xFFFFFFFF),
			uint32(value[7]>>32), uint32(value[7]&0xFFFFFFFF))
	}
	gdextension.Host.Objects.Unsafe.Call = func(p0 gdextension.Object, p1 gdextension.MethodForClass, p2 gdextension.CallReturns[interface{}], shape gdextension.Shape, p4 gdextension.CallAccepts[interface{}]) {
		mem2 := gdmemory.MakeResult(shape)
		mem4 := gdmemory.CopyArguments(shape, p4)
		wasm_gd_object_unsafe_call(uint32(p0), uint32(p1), uint32(mem2), uint32(shape>>32), uint32(shape&0xFFFFFFFF), uint32(mem4))
		gdmemory.LoadResult(shape, p2, mem2)
	}
}
