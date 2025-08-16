//go:build js

package gdextension

import "unsafe"

type Pointer uint32
type Returns[T any] uint32
type Accepts[T any] uint32
type PackedArray [2]Pointer

const (
	SizeString      Shape = SizeBytes4
	SizeObject      Shape = SizeBytes4
	SizeArray       Shape = SizeBytes4
	SizePackedArray Shape = SizeBytes8
	SizeDictionary  Shape = SizeBytes4
	SizeStringName  Shape = SizeBytes4
	SizeNodePath    Shape = SizeBytes4
)

func ToPackedArray(p [1]uint64) PackedArray {
	return *(*PackedArray)(unsafe.Pointer(&p))
}
