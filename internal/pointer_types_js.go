package gd

import (
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
)

type gdptr uint32

type EnginePointer = uint32
type PackedPointers = [1]uint64

func UnsafeGet[T any](frame gdextension.Pointer, index int) T {
	// frame is a list of pointers, so we need to get the pointer at the index
	var ptr = frame + gdextension.Pointer(uintptr(index)*unsafe.Sizeof(gdextension.Pointer(0)))
	var addr = gdextension.Pointer(gdunsafe.Pointer(ptr).Uint32())
	var zero T
	var done = 0
	var size = unsafe.Sizeof([1]T{})
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.Pointer(addr).Uint32()
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.Pointer(addr).Uint16()
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.Pointer(addr).Byte()
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
			gdunsafe.Pointer(addr).SetUint64(*(*uint64)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 8
			done += 8
			size -= 8
		case size >= 4:
			gdunsafe.Pointer(addr).SetUint32(*(*uint32)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			gdunsafe.Pointer(addr).SetUint16(*(*uint16)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			gdunsafe.Pointer(addr).SetByte(*(*uint8)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 1
			done += 1
			size -= 1
		}
	}
}
