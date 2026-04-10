package noescape

import (
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
)

type MethodForClass gdextension.MethodForClass
type Variant = gdunsafe.Variant

func CallStatic[T any](method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return Call[T](0, method, shape, args)
}

//go:noescape
func Pointer(ptr unsafe.Pointer) unsafe.Pointer

//go:linkname pointer graphics.gd/internal/noescape.Pointer
//go:nosplit
func pointer(ptr unsafe.Pointer) unsafe.Pointer {
	return ptr
}
