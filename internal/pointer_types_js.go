package gd

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

type gdptr uint32

type EnginePointer = uint32
type PackedPointers = [1]uint64

func UnsafeGet[T any](frame Address, index int) T {
	// frame is a list of pointers, so we need to get the pointer at the index
	var ptr = gdextension.Pointer(uintptr(frame) + uintptr(index)*unsafe.Sizeof(gdextension.Pointer(0)))
	var addr = gdextension.Pointer(gdextension.Host.Memory.Load.Uint32(ptr))
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

func UnsafeSet[T any](frame Address, value T) {
	var addr = gdextension.Pointer(uintptr(frame))
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
