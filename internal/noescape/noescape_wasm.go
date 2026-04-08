//go:build wasm

package noescape

import (
	"reflect"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/ring"
)

func Call[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	var argptr unsafe.Pointer = nil
	var result T
	if args != nil {
		argptr = reflect.ValueOf(args).UnsafePointer()
	}
	if unsafe.Sizeof(result) == 0 {
		ring.Main.Buffer(uintptr(object), uintptr(method), uint64(shape), argptr, 0)
		return result
	}
	if ring.Main.Pending() {
		ring.Main.Flush()
	}
	call_noescape(object, method, unsafe.Pointer(&result), shape, argptr)
	return result
}

//go:noescape
func call_noescape(object gdextension.Object, method gdextension.MethodForClass, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/noescape.call_noescape
func call(object gdextension.Object, method gdextension.MethodForClass, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer) {
	gdunsafe.Object(object).ShapedCall(method, result, shape, args)
}

func (method MethodForClass) Call(self gdextension.Object, args ...gdextension.Variant) (gdextension.Variant, error) {
	var result gdextension.Variant
	var err gdextension.CallError
	object_method_call_noescape(self, gdextension.MethodForClass(method), &result, args, &err)
	return result, err
}

//go:noescape
func object_method_call_noescape(object gdextension.Object, method gdextension.MethodForClass, result *gdextension.Variant, args []gdextension.Variant, err *gdextension.CallError)

//go:linkname object_method_call graphics.gd/internal/noescape.object_method_call_noescape
func object_method_call(object gdextension.Object, method gdextension.MethodForClass, result *gdextension.Variant, args []gdextension.Variant, err *gdextension.CallError) {
	*result, *err = gdunsafe.Object(object).Call(method, args...)
}
