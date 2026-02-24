//go:build i386 || arm

package gd

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

type gdptr uint32

type EnginePointer = uint32
type PackedPointers = [2]uint32

func UnsafeGet[T any](frame gdextension.Pointer, index int) T {
	ptrs := *(*unsafe.Pointer)(unsafe.Pointer(&frame))
	return *unsafe.Slice((**T)(ptrs), index+1)[index]
}

func UnsafeSet[T any](frame gdextension.Pointer, value T) {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&frame))
	*(*T)(ptr) = value
}
