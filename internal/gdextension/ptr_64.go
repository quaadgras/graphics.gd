//go:build !js

package gdextension

import (
	"unsafe"

	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

type Pointer = uintptr
type Returns[T any] unsafe.Pointer
type Accepts[T any] unsafe.Pointer
type PackedArray[T byte | int32 | int64 | float32 | float64 | Color.RGBA | Vector2.XY | Vector3.XYZ | Vector4.XYZW | String] [2]uint64

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
