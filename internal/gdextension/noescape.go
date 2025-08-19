package gdextension

import (
	"unsafe"
)

func (method MethodForClass) Call(self Object, args ...Variant) (Variant, error) {
	var result Variant
	var err CallError
	object_method_call_noescape(self, method, &result, args, &err)
	return result, err.Err()
}

//go:noescape
func object_method_call_noescape(object Object, method MethodForClass, result *Variant, args []Variant, err *CallError)

//go:linkname object_method_call graphics.gd/internal/gdextension.object_method_call_noescape
func object_method_call(object Object, method MethodForClass, result *Variant, args []Variant, err *CallError) {
	Host.Objects.Call(object, method, CallReturns[Variant](result), len(args), CallAccepts[Variant](unsafe.SliceData(args)), CallReturns[CallError](err))
}

func Call[T any](object Object, method MethodForClass, shape Shape, args unsafe.Pointer) T {
	var result T
	call_noescape(object, method, unsafe.Pointer(&result), shape, args)
	return result
}

func CallStatic[T any](method MethodForClass, shape Shape, args unsafe.Pointer) T {
	var result T
	call_noescape(0, method, unsafe.Pointer(&result), shape, args)
	return result
}

//go:noescape
func call_noescape(object Object, method MethodForClass, result unsafe.Pointer, shape Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/gdextension.call_noescape
func call(object Object, method MethodForClass, result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	Host.Objects.Unsafe.Call(object, method, CallReturns[any](result), shape, CallAccepts[any](args))
}
