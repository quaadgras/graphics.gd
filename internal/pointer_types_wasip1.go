//go:build wasip1

package gd

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

// The wasip1 build is the hot-reloading guest (see graphics.gd/startup
// with -tags reloads): it runs inside a wazero runtime embedded in the
// native host, and every gdextension.Pointer it is handed addresses the
// HOST's memory, not the guest's linear memory. Pointers are 64-bit on
// GOARCH=wasm so they survive the trip, but they must never be
// dereferenced directly — a virtual method's argument frame is an array
// of native pointers into engine memory, and reading it as a guest
// address traps with "out of bounds memory access". All accesses go
// through the host's memory bridge instead, like the js build does.

type gdptr uint64

type EnginePointer = uint64
type PackedPointers = [2]uint64

// loadPointer reads a native (64-bit) pointer from host memory.
func loadPointer(addr gdextension.Pointer) gdextension.Pointer {
	lo := gdextension.Host.Memory.Load.Uint32(addr)
	hi := gdextension.Host.Memory.Load.Uint32(addr + 4)
	return gdextension.Pointer(uint64(hi)<<32 | uint64(lo))
}

func UnsafeGet[T any](frame gdextension.Pointer, index int) T {
	// frame is a list of native pointers, one per argument.
	var addr = loadPointer(frame + gdextension.Pointer(index)*gdextension.Pointer(unsafe.Sizeof(uint64(0))))
	var zero T
	var done = 0
	var size = unsafe.Sizeof([1]T{})
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdextension.Host.Memory.Load.Uint32(addr)
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdextension.Host.Memory.Load.Uint16(addr)
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdextension.Host.Memory.Load.Byte(addr)
			addr += 1
			done += 1
			size -= 1
		}
	}
	return zero
}

func UnsafeSet[T any](addr gdextension.Pointer, value T) {
	var size = unsafe.Sizeof([1]T{})
	var done = 0
	for size > 0 {
		switch {
		case size >= 8:
			gdextension.Host.Memory.Edit.Uint64(addr, *(*uint64)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 8
			done += 8
			size -= 8
		case size >= 4:
			gdextension.Host.Memory.Edit.Uint32(addr, *(*uint32)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			gdextension.Host.Memory.Edit.Uint16(addr, *(*uint16)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			gdextension.Host.Memory.Edit.Byte(addr, *(*uint8)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 1
			done += 1
			size -= 1
		}
	}
}
