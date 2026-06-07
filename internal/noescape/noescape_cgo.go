//go:build cgo

package noescape

// #include <stdint.h>
//
// typedef struct { uint64_t part[2]; } result_16;
// typedef struct { uint64_t part[3]; } result_24;
// typedef struct { uint64_t part[4]; } result_32;
// typedef struct { uint64_t part[8]; } result_64;
//
// extern void gd_object_unsafe_call(uintptr_t obj, uintptr_t method, void *result, uint64_t shape, void *args);
//
// extern uint64_t gd_object_unsafe_call_8(uintptr_t obj, uintptr_t method, uint64_t shape, void *args);
// extern result_16 gd_object_unsafe_call_16(uintptr_t obj, uintptr_t method, uint64_t shape, void *args);
// extern result_24 gd_object_unsafe_call_24(uintptr_t obj, uintptr_t method, uint64_t shape, void *args);
// extern result_32 gd_object_unsafe_call_32(uintptr_t obj, uintptr_t method, uint64_t shape, void *args);
// extern result_64 gd_object_unsafe_call_64(uintptr_t obj, uintptr_t method, uint64_t shape, void *args);
//
// extern result_24 gd_object_call_24(uintptr_t obj, uintptr_t fn, int64_t argc, void *args, void *err);
//
import "C"
import (
	"reflect"
	"unsafe"

	"graphics.gd/internal/callerpc"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadcheck"
)

func Call[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	var argptr unsafe.Pointer = nil
	var result T
	if args != nil {
		argptr = reflect.ValueOf(args).UnsafePointer()
	}
	if unsafe.Sizeof(result) == 0 {
		if threadcheck.Main() {
			ring.Main.Buffer(uintptr(object), uintptr(method), uint64(shape), argptr, callerpc.Callerpc())
			return result
		}
		call_noescape(object, method, unsafe.Pointer(&result), shape, argptr)
		return result
	}
	if threadcheck.Main() && ring.Main.Pending() {
		ring.Main.Flush()
	}
	switch {
	case unsafe.Sizeof(result) <= 8:
		var r8 = call_8_noescape(object, method, shape, argptr)
		result = *(*T)(unsafe.Pointer(&r8))
	case unsafe.Sizeof(result) <= 16:
		var r16 = call_16_noescape(object, method, shape, argptr)
		result = *(*T)(unsafe.Pointer(&r16))
	case unsafe.Sizeof(result) <= 32:
		var r32 = call_32_noescape(object, method, shape, argptr)
		result = *(*T)(unsafe.Pointer(&r32))
	case unsafe.Sizeof(result) <= 64:
		var r64 = call_64_noescape(object, method, shape, argptr)
		result = *(*T)(unsafe.Pointer(&r64))
	default:
		panic("return size too large")
	}
	return result
}

//go:noescape
func call_noescape(object gdextension.Object, method gdextension.MethodForClass, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/noescape.call_noescape
//go:nosplit
func call(object gdextension.Object, method gdextension.MethodForClass, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer) {
	C.gd_object_unsafe_call(C.uintptr_t(object), C.uintptr_t(method), result, C.uint64_t(shape), args)
}

//go:noescape
func call_8_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer) uint64

//go:linkname call_8 graphics.gd/internal/noescape.call_8_noescape
//go:nosplit
func call_8(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer) uint64 {
	return uint64(C.gd_object_unsafe_call_8(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args))
}

//go:noescape
func call_16_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer) (result C.result_16)

//go:linkname call_16 graphics.gd/internal/noescape.call_16_noescape
//go:nosplit
func call_16(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer) C.result_16 {
	return C.gd_object_unsafe_call_16(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args)
}

//go:noescape
func call_32_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer) (result C.result_32)

//go:linkname call_32 graphics.gd/internal/noescape.call_32_noescape
//go:nosplit
func call_32(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer) C.result_32 {
	return C.gd_object_unsafe_call_32(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args)
}

//go:noescape
func call_64_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer) (result C.result_64)

//go:linkname call_64 graphics.gd/internal/noescape.call_64_noescape
//go:nosplit
func call_64(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer) C.result_64 {
	return C.gd_object_unsafe_call_64(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args)
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
//go:nosplit
func object_method_call(object gdextension.Object, method gdextension.MethodForClass, result *gdextension.Variant, args []gdextension.Variant, err *gdextension.CallError) {
	raw := C.gd_object_call_24(C.uintptr_t(object), C.uintptr_t(method), C.int64_t(len(args)), unsafe.Pointer(unsafe.SliceData(args)), unsafe.Pointer(err))
	*result = *(*gdextension.Variant)(unsafe.Pointer(&raw))
}
