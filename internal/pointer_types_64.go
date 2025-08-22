//go:build amd64 || arm64 || wasip1

package gd

import (
	"unsafe"
)

type gdptr uint64

type EnginePointer = uint64
type PackedPointers = [2]uint64

func UnsafeGet[T any](frame Address, index int) T {
	ptrs := *(*unsafe.Pointer)(unsafe.Pointer(&frame))
	return *unsafe.Slice((**T)(ptrs), index+1)[index]
}

func UnsafeSet[T any](frame Address, value T) {
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&frame))
	*(*T)(ptr) = value
}
