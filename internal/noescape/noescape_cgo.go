//go:build cgo && !(gd && wasm)

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
// extern result_24 gd_callable_callv_24(uintptr_t fn, uint64_t c1, uint64_t c2, uintptr_t array);
//
import "C"
import (
	"sync/atomic"
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
		argptr = argPointer(args)
	}
	if unsafe.Sizeof(result) == 0 {
		if threadcheck.Main() {
			ring.Main.Buffer(uintptr(object), uintptr(method), uint64(shape), argptr, callerpc.Callerpc())
			return result
		}
		if !threadcheck.Engine() {
			// user goroutine: the engine is not thread-safe, so record the
			// call in the cross-thread ring for the main thread to execute.
			ring.Threads.Buffer(uintptr(object), uintptr(method), uint64(shape), argptr, callerpc.Callerpc())
			return result
		}
		call_noescape(object, method, unsafe.Pointer(&result), shape, argptr)
		return result
	}
	if threadcheck.Main() {
		if ring.Main.Pending() {
			ring.Main.Flush()
		}
		// Resident fast path, inlined here rather than reached through
		// direct[T]: this is the per-frame hot route for every
		// result-bearing engine call, and direct would re-do the
		// threadcheck.Main this branch has already established.
		if n, ok := residentAvailable(shape); ok {
			return residentCall[T](object, method, shape, n, argptr)
		}
	} else if !threadcheck.Engine() {
		// user goroutine: queue the call and block until the main thread has
		// executed it and delivered the return value.
		ring.Threads.Call(uintptr(object), uintptr(method), uint64(shape), argptr, callerpc.Callerpc(), unsafe.Pointer(&result), unsafe.Sizeof(result))
		return result
	}
	return direct[T](object, method, shape, argptr)
}

// direct performs the engine crossing on the calling thread.
func direct[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, argptr unsafe.Pointer) T {
	if threadcheck.Main() {
		if n, ok := residentAvailable(shape); ok {
			// Resident-callback mode: cross via asmcgocall so nested engine->Go
			// callbacks take the resident fast path (see resident_cgo.go).
			return residentCall[T](object, method, shape, n, argptr)
		}
	}
	var result T
	if unsafe.Sizeof(result) == 0 {
		call_noescape(object, method, unsafe.Pointer(&result), shape, argptr)
		return result
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

// CallThreadSafe is like Call, for methods of the engine singletons that the
// engine guards internally against concurrent access (see
// gdfunc.ThreadSafeSingletons): goroutines cross into the engine directly
// instead of queueing for the main thread. On the main thread it behaves
// exactly like Call, preserving command-buffer ordering.
func CallThreadSafe[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	if threadcheck.Main() {
		return Call[T](object, method, shape, args)
	}
	var argptr unsafe.Pointer
	if args != nil {
		argptr = argPointer(args)
	}
	return direct[T](object, method, shape, argptr)
}

// CallThreadSafeIf is CallThreadSafe gated on a runtime condition, for
// singletons that are only thread-safe under certain project settings (the
// physics servers); when the condition is false it behaves like Call.
func CallThreadSafeIf[T any](safe *atomic.Bool, object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	if safe.Load() && !threadcheck.Main() {
		var argptr unsafe.Pointer
		if args != nil {
			argptr = argPointer(args)
		}
		return direct[T](object, method, shape, argptr)
	}
	return Call[T](object, method, shape, args)
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
	main := threadcheck.Main()
	if main && residentVariantAvailable(args) {
		// Resident-callback mode: script/vararg calls are the calls most
		// likely to re-enter Go (GDScript invoking a Go callable), so the
		// fast path for their nested callbacks matters most here.
		return residentVariantCall(self, method, args)
	}
	// result and err must not sit on this goroutine's stack: the engine
	// writes the error after the method has run, and a callback into Go
	// during the call can move the stack, so the write would land on the
	// copy that was left behind. They are on the heap only because the
	// closure below captures them and escapes. Keep it that way: if that
	// branch moves out of this function, allocate them explicitly.
	var result gdextension.Variant
	var err gdextension.CallError
	if main || threadcheck.Engine() {
		object_method_call_noescape(self, gdextension.MethodForClass(method), &result, args, &err)
	} else {
		// user goroutine: variadic variant calls cannot be encoded as a ring
		// entry, so run the whole call on the main thread in queue order.
		ring.Threads.Run(func() {
			object_method_call_noescape(self, gdextension.MethodForClass(method), &result, args, &err)
		})
	}
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

// CallableCallv is Callable.callv with the callable, the argument array and
// the result all passed by value, so that nothing the engine reads or writes
// during the call sits on the calling goroutine's stack: the function being
// called can be Go code, and Go code can move that stack.
func CallableCallv(callv gdextension.MethodForBuiltinType, callable gdextension.Callable, args gdextension.Array) gdextension.Variant {
	threadcheck.EnterCall()
	raw := C.gd_callable_callv_24(C.uintptr_t(callv), C.uint64_t(callable[0]), C.uint64_t(callable[1]), C.uintptr_t(args[0]))
	threadcheck.LeaveCall()
	return *(*gdextension.Variant)(unsafe.Pointer(&raw))
}
