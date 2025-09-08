//go:build cgo

package gdextension

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
)

func Call[T any](object Object, method MethodForClass, shape Shape, args any) T {
	var argptr unsafe.Pointer = nil
	var result T
	if args != nil {
		argptr = reflect.ValueOf(args).UnsafePointer()
	}
	switch {
	case unsafe.Sizeof(result) == 0:
		call_noescape(object, method, unsafe.Pointer(&result), shape, argptr)
		return result
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

func CallStatic[T any](method MethodForClass, shape Shape, args any) T {
	return Call[T](0, method, shape, args)
}

//go:noescape
func call_noescape(object Object, method MethodForClass, result unsafe.Pointer, shape Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/gdextension.call_noescape
//go:nosplit
func call(object Object, method MethodForClass, result unsafe.Pointer, shape Shape, args unsafe.Pointer) {
	C.gd_object_unsafe_call(C.uintptr_t(object), C.uintptr_t(method), result, C.uint64_t(shape), args)
}

//go:noescape
func call_8_noescape(object Object, method MethodForClass, shape Shape, args unsafe.Pointer) uint64

//go:linkname call_8 graphics.gd/internal/gdextension.call_8_noescape
//go:nosplit
func call_8(object Object, method MethodForClass, shape Shape, args unsafe.Pointer) uint64 {
	return uint64(C.gd_object_unsafe_call_8(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args))
}

//go:noescape
func call_16_noescape(object Object, method MethodForClass, shape Shape, args unsafe.Pointer) (result C.result_16)

//go:linkname call_16 graphics.gd/internal/gdextension.call_16_noescape
//go:nosplit
func call_16(object Object, method MethodForClass, shape Shape, args unsafe.Pointer) C.result_16 {
	return C.gd_object_unsafe_call_16(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args)
}

//go:noescape
func call_32_noescape(object Object, method MethodForClass, shape Shape, args unsafe.Pointer) (result C.result_32)

//go:linkname call_32 graphics.gd/internal/gdextension.call_32_noescape
//go:nosplit
func call_32(object Object, method MethodForClass, shape Shape, args unsafe.Pointer) C.result_32 {
	return C.gd_object_unsafe_call_32(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args)
}

//go:noescape
func call_64_noescape(object Object, method MethodForClass, shape Shape, args unsafe.Pointer) (result C.result_64)

//go:linkname call_64 graphics.gd/internal/gdextension.call_64_noescape
//go:nosplit
func call_64(object Object, method MethodForClass, shape Shape, args unsafe.Pointer) C.result_64 {
	return C.gd_object_unsafe_call_64(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args)
}

func (method MethodForClass) Call(self Object, args ...Variant) (Variant, error) {
	var result Variant
	var err CallError
	object_method_call_noescape(self, method, &result, args, &err)
	return result, err.Err()
}

//go:noescape
func object_method_call_noescape(object Object, method MethodForClass, result *Variant, args []Variant, err *CallError)

//go:linkname object_method_call graphics.gd/internal/gdextension.object_method_call_noescape
//go:nosplit
func object_method_call(object Object, method MethodForClass, result *Variant, args []Variant, err *CallError) {
	raw := C.gd_object_call_24(C.uintptr_t(object), C.uintptr_t(method), C.int64_t(len(args)), unsafe.Pointer(unsafe.SliceData(args)), unsafe.Pointer(err))
	*result = *(*Variant)(unsafe.Pointer(&raw))
}
