package gd

import (
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
)

type gdptr uint32

type EnginePointer = uint32
type PackedPointers = [2]uint32

func UnsafeGet[T any](frame gdextension.Pointer, index int) T {
	// frame is a list of pointers, so we need to get the pointer at the index
	var ptr = frame + gdextension.Pointer(uintptr(index)*unsafe.Sizeof(gdextension.Pointer(0)))
	var addr = gdunsafe.PointerTo[gdunsafe.Pointer](ptr).Get()
	var zero T
	var done = 0
	var size = unsafe.Sizeof([1]T{})
	for size > 0 {
		switch {
		case size >= 4:
			*(*uint32)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.PointerTo[uint32](addr).Get()
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			*(*uint16)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.PointerTo[uint16](addr).Get()
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			*(*uint8)(unsafe.Add(unsafe.Pointer(&zero), done)) = gdunsafe.PointerTo[uint8](addr).Get()
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
			gdunsafe.MutablePointerTo[uint64](addr).Set(*(*uint64)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 8
			done += 8
			size -= 8
		case size >= 4:
			gdunsafe.MutablePointerTo[uint32](addr).Set(*(*uint32)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 4
			done += 4
			size -= 4
		case size >= 2:
			gdunsafe.MutablePointerTo[uint16](addr).Set(*(*uint16)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 2
			done += 2
			size -= 2
		case size >= 1:
			gdunsafe.MutablePointerTo[uint8](addr).Set(*(*uint8)(unsafe.Add(unsafe.Pointer(&value), done)))
			addr += 1
			done += 1
			size -= 1
		}
	}
}
