package gdunsafe

import (
	"unsafe"

	"graphics.gd/internal/gdextension"
)

func Call[T any](object gdextension.Object, method gdextension.FunctionID, shape gdextension.Shape, args unsafe.Pointer) T {
	var result T
	call_noescape(object, method, unsafe.Pointer(&result), shape, args)
	return result
}

//go:noescape
func call_noescape(object gdextension.Object, method gdextension.FunctionID, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/gdunsafe.call_noescape
func call(object gdextension.Object, method gdextension.FunctionID, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer) {
	gdextension.Host.Objects.Unsafe.Call(object, method, gdextension.CallReturns[any](result), shape, gdextension.CallAccepts[any](args))
}
