package gdunsafe

import (
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

type Use struct{}

func Call[T any](object [1]gd.Object, method gd.MethodBind, shape gdextension.Shape, args unsafe.Pointer) T {
	var result T
	call_noescape(gdextension.Object(pointers.Get(object[0].AsObject()[0])[0]), gdextension.FunctionID(method), unsafe.Pointer(&result), shape, args)
	return result
}

func CallStatic[T any](method gd.MethodBind, shape gdextension.Shape, args unsafe.Pointer) T {
	var result T
	call_noescape(0, gdextension.FunctionID(method), unsafe.Pointer(&result), shape, args)
	return result
}

//go:noescape
func call_noescape(object gdextension.Object, method gdextension.FunctionID, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/gdunsafe.call_noescape
func call(object gdextension.Object, method gdextension.FunctionID, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer) {
	gdextension.Host.Objects.Unsafe.Call(object, method, gdextension.CallReturns[any](result), shape, gdextension.CallAccepts[any](args))
}
