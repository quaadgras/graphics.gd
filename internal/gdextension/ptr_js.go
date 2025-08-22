//go:build js

package gdextension

import (
	"unsafe"

	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

type Pointer = uint32
type Returns[T any] Pointer
type Accepts[T any] Pointer
type PackedArray[T byte | int32 | int64 | float32 | float64 | Color.RGBA | Vector2.XY | Vector3.XYZ | Vector4.XYZW | String] [1]uint64

const (
	SizeString      Shape = SizeBytes4
	SizeObject      Shape = SizeBytes4
	SizeArray       Shape = SizeBytes4
	SizePackedArray Shape = SizeBytes8
	SizeDictionary  Shape = SizeBytes4
	SizeStringName  Shape = SizeBytes4
	SizeNodePath    Shape = SizeBytes4
	SizePointer     Shape = SizeBytes4
)

func (p PackedArray[T]) JS() [2]uint32 {
	return *(*[2]uint32)(unsafe.Pointer(&p))
}
