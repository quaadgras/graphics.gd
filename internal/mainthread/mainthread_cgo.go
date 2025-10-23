//go:build cgo

package mainthread

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

// #include <stdint.h>
//
// typedef struct { uint64_t safe; uint64_t part[1]; } result_8;
// typedef struct { uint64_t safe; uint64_t part[2]; } result_16;
// typedef struct { uint64_t safe; uint64_t part[3]; } result_24;
// typedef struct { uint64_t safe; uint64_t part[4]; } result_32;
// typedef struct { uint64_t safe; uint64_t part[8]; } result_64;
//
// extern uint64_t gd_object_unsafe_call_0(uintptr_t obj, uintptr_t method, uint64_t shape, void *args, uintptr_t mainthread);
// extern result_8 gd_object_unsafe_call_8(uintptr_t obj, uintptr_t method, uint64_t shape, void *args, uintptr_t mainthread);
// extern result_16 gd_object_unsafe_call_16(uintptr_t obj, uintptr_t method, uint64_t shape, void *args, uintptr_t mainthread);
// extern result_24 gd_object_unsafe_call_24(uintptr_t obj, uintptr_t method, uint64_t shape, void *args, uintptr_t mainthread);
// extern result_32 gd_object_unsafe_call_32(uintptr_t obj, uintptr_t method, uint64_t shape, void *args, uintptr_t mainthread);
// extern result_64 gd_object_unsafe_call_64(uintptr_t obj, uintptr_t method, uint64_t shape, void *args, uintptr_t mainthread);
//
// extern result_24 gd_object_call_24(uintptr_t obj, uintptr_t fn, int64_t argc, void *args, void *err);
//
import "C"

var requests = make(chan request, runtime.NumCPU())

type request struct {
	object gdextension.Object
	method gdextension.MethodForClass
	shape  gdextension.Shape
	args   unsafe.Pointer
	length int
	result *response[[8]C.uint64_t]
}

type response[T [0]C.uint64_t | [1]C.uint64_t | [2]C.uint64_t | [4]C.uint64_t | [8]C.uint64_t] struct {
	ready sync.WaitGroup
	value T
}

func CallStatic[T any](method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return Call[T](0, method, shape, args)
}

func Yield() {
	for range len(requests) {
		req := <-requests
		func() {
			defer func() {
				recover()
				req.result.ready.Done()
			}()
			switch req.length {
			case 0:
				C.gd_object_unsafe_call_0(C.uintptr_t(req.object), C.uintptr_t(req.method), C.uint64_t(req.shape), req.args, 0)
			case 1:
				result := C.gd_object_unsafe_call_8(C.uintptr_t(req.object), C.uintptr_t(req.method), C.uint64_t(req.shape), req.args, 0).part
				copy(req.result.value[:], result[:])
			case 2:
				result := C.gd_object_unsafe_call_16(C.uintptr_t(req.object), C.uintptr_t(req.method), C.uint64_t(req.shape), req.args, 0).part
				copy(req.result.value[:], result[:])
			case 4:
				result := C.gd_object_unsafe_call_32(C.uintptr_t(req.object), C.uintptr_t(req.method), C.uint64_t(req.shape), req.args, 0).part
				copy(req.result.value[:], result[:])
			case 8:
				req.result.value = C.gd_object_unsafe_call_64(C.uintptr_t(req.object), C.uintptr_t(req.method), C.uint64_t(req.shape), req.args, 0).part
			}
		}()
	}
}

func Call[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	var argptr unsafe.Pointer = nil
	var result T
	if args != nil {
		argptr = reflect.ValueOf(args).UnsafePointer()
	}
	switch {
	case unsafe.Sizeof(result) == 0:
		var thread response[[0]C.uint64_t]
		call_noescape(object, method, shape, argptr, &thread)
		return result
	case unsafe.Sizeof(result) <= 8:
		var thread response[[1]C.uint64_t]
		var r8 = call_8_noescape(object, method, shape, argptr, &thread)
		result = *(*T)(unsafe.Pointer(&r8))
	case unsafe.Sizeof(result) <= 16:
		var thread response[[2]C.uint64_t]
		var r16 = call_16_noescape(object, method, shape, argptr, &thread)
		result = *(*T)(unsafe.Pointer(&r16))
	case unsafe.Sizeof(result) <= 32:
		var thread response[[4]C.uint64_t]
		var r32 = call_32_noescape(object, method, shape, argptr, &thread)
		result = *(*T)(unsafe.Pointer(&r32))
	case unsafe.Sizeof(result) <= 64:
		var thread response[[8]C.uint64_t]
		var r64 = call_64_noescape(object, method, shape, argptr, &thread)
		result = *(*T)(unsafe.Pointer(&r64))
	default:
		panic("return size too large")
	}
	return result
}

//go:noescape
func call_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[0]C.uint64_t])

//go:linkname call graphics.gd/internal/mainthread.call_noescape
//go:nosplit
func call(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[8]C.uint64_t]) {
	thread := uint64(C.gd_object_unsafe_call_0(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args, C.uintptr_t(uintptr(unsafe.Pointer(mainthread)))))
	if thread != 0 {
		mainthread.ready.Add(1)
		requests <- request{object, method, shape, args, 0, mainthread}
		mainthread.ready.Wait()
	} else {
		Yield()
	}
}

//go:noescape
func call_8_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[1]C.uint64_t]) uint64

//go:linkname call_8 graphics.gd/internal/mainthread.call_8_noescape
//go:nosplit
func call_8(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[8]C.uint64_t]) [1]C.uint64_t {
	result := C.gd_object_unsafe_call_8(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args, C.uintptr_t(uintptr(unsafe.Pointer(mainthread))))
	if thread := result.safe; thread != 0 {
		mainthread.ready.Add(1)
		requests <- request{object, method, shape, args, 1, mainthread}
		mainthread.ready.Wait()
		return [1]C.uint64_t(mainthread.value[:])
	}
	Yield()
	return result.part
}

//go:noescape
func call_16_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[2]C.uint64_t]) [2]uint64

//go:linkname call_16 graphics.gd/internal/mainthread.call_16_noescape
//go:nosplit
func call_16(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[8]C.uint64_t]) [2]C.uint64_t {
	result := C.gd_object_unsafe_call_16(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args, C.uintptr_t(uintptr(unsafe.Pointer(mainthread))))
	if thread := result.safe; thread != 0 {
		mainthread.ready.Add(1)
		requests <- request{object, method, shape, args, 2, mainthread}
		mainthread.ready.Wait()
		return [2]C.uint64_t(mainthread.value[:])
	}
	Yield()
	return result.part
}

//go:noescape
func call_32_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[4]C.uint64_t]) (result [4]C.uint64_t)

//go:linkname call_32 graphics.gd/internal/mainthread.call_32_noescape
//go:nosplit
func call_32(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[8]C.uint64_t]) [4]C.uint64_t {
	result := C.gd_object_unsafe_call_32(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args, C.uintptr_t(uintptr(unsafe.Pointer(mainthread))))
	if thread := result.safe; thread != 0 {
		mainthread.ready.Add(1)
		requests <- request{object, method, shape, args, 4, mainthread}
		mainthread.ready.Wait()
		return [4]C.uint64_t(mainthread.value[:])
	}
	Yield()
	return result.part
}

//go:noescape
func call_64_noescape(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[8]C.uint64_t]) (result [8]C.uint64_t)

//go:linkname call_64 graphics.gd/internal/mainthread.call_64_noescape
//go:nosplit
func call_64(object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args unsafe.Pointer, mainthread *response[[8]C.uint64_t]) [8]C.uint64_t {
	result := C.gd_object_unsafe_call_64(C.uintptr_t(object), C.uintptr_t(method), C.uint64_t(shape), args, C.uintptr_t(uintptr(unsafe.Pointer(mainthread))))
	if thread := result.safe; thread != 0 {
		mainthread.ready.Add(1)
		requests <- request{object, method, shape, args, 8, mainthread}
		mainthread.ready.Wait()
		return [8]C.uint64_t(mainthread.value[:])
	}
	Yield()
	return result.part
}
