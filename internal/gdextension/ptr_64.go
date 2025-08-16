//go:build !js

package gdextension

import "unsafe"

type Pointer uintptr
type Returns[T any] unsafe.Pointer
type Accepts[T any] unsafe.Pointer
type PackedArray [2]Pointer

const (
	SizeString      Shape = SizeBytes8
	SizeObject      Shape = SizeBytes8
	SizeArray       Shape = SizeBytes8
	SizePackedArray Shape = SizeBytes16
	SizeDictionary  Shape = SizeBytes8
	SizeStringName  Shape = SizeBytes8
	SizeNodePath    Shape = SizeBytes8
	SizePointer     Shape = SizeBytes8
)

func ToPackedArray(p [2]uint64) PackedArray {
	return PackedArray{Pointer(p[0]), Pointer(p[1])}
}
