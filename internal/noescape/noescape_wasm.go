//go:build wasm

package noescape

import (
	"reflect"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

func Call[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	var argptr unsafe.Pointer = nil
	var result T
	if args != nil {
		argptr = reflect.ValueOf(args).UnsafePointer()
	}
	call_noescape(object, method, unsafe.Pointer(&result), shape, argptr)
	return result
}

//go:noescape
func call_noescape(object gdextension.Object, method gdextension.MethodForClass, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/noescape.call_noescape
func call(object gdextension.Object, method gdextension.MethodForClass, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer) {
	gdextension.Host.Objects.Unsafe.Call(object, method, gdextension.CallReturns[any](result), shape, gdextension.CallAccepts[any](args))
}

func (method MethodForClass) Call(self gdextension.Object, args ...gdextension.Variant) (gdextension.Variant, error) {
	var result gdextension.Variant
	var err gdextension.CallError
	object_method_call_noescape(self, gdextension.MethodForClass(method), &result, args, &err)
	return result, err.Err()
}

//go:noescape
func object_method_call_noescape(object gdextension.Object, method gdextension.MethodForClass, result *gdextension.Variant, args []gdextension.Variant, err *gdextension.CallError)

//go:linkname object_method_call graphics.gd/internal/noescape.object_method_call_noescape
func object_method_call(object gdextension.Object, method gdextension.MethodForClass, result *gdextension.Variant, args []gdextension.Variant, err *gdextension.CallError) {
	gdextension.Host.Objects.Call(object, method, gdextension.CallReturns[gdextension.Variant](result), len(args), gdextension.CallAccepts[gdextension.Variant](unsafe.SliceData(args)), gdextension.CallReturns[gdextension.CallError](err))
}
