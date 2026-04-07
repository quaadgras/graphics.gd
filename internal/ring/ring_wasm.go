//go:build js

package ring

import (
	"unsafe"

	gdunsafe "graphics.gd"
)

var shadowRing gdunsafe.Pointer

//go:wasmimport gd ring_flush
func wasm_gd_ring_flush(entries gdunsafe.Pointer, stride uint32, tail uint32, head uint32)

//go:wasmimport gd bulk_copy
func wasm_gd_bulk_copy(godot_dst gdunsafe.Pointer, go_src uint32, length uint32)

func flush(entries unsafe.Pointer, tail, head uint32) {
	if shadowRing == 0 {
		shadowRing = gdunsafe.Malloc(int64(Size * uint32(unsafe.Sizeof(Entry{}))))
	}
	entrySize := uint32(unsafe.Sizeof(Entry{}))
	start := tail & Mask
	end := head & Mask
	base := uint32(uintptr(entries))
	if start < end {
		// Contiguous range
		wasm_gd_bulk_copy(
			shadowRing+gdunsafe.Pointer(start*entrySize),
			base+start*entrySize,
			(end-start)*entrySize,
		)
	} else {
		// Wraps around: copy [start..Size) then [0..end)
		wasm_gd_bulk_copy(
			shadowRing+gdunsafe.Pointer(start*entrySize),
			base+start*entrySize,
			(Size-start)*entrySize,
		)
		if end > 0 {
			wasm_gd_bulk_copy(
				shadowRing,
				base,
				end*entrySize,
			)
		}
	}
	wasm_gd_ring_flush(shadowRing, entrySize, tail, head)
}
