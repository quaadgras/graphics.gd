//go:build go1.26 && (amd64 || arm64)

package jumponly

import (
	"reflect"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/ring"
	"graphics.gd/internal/threadcheck"
)

// PtrcallFn holds the address of gdextension_object_method_bind_ptrcall,
// set during engine CORE-level init.
var PtrcallFn uintptr

// Call invokes a trivial Godot method via a direct assembly trampoline,
// bypassing cgo and the ring buffer. The method must have zero function
// calls in its C++ body (field reads/writes only).
//
// Call has the same signature as noescape.Call[T].
func Call[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	var result T
	// Flush any pending ring buffer entries to maintain ordering.
	if ring.Main.Pending() && threadcheck.Main() {
		ring.Main.Flush()
	}
	var argptr unsafe.Pointer
	if args != nil {
		argptr = reflect.ValueOf(args).UnsafePointer()
	}
	// Decode shape to build the ptrcall args pointer array.
	var ptrs [16]unsafe.Pointer
	offset := uintptr(0)
	for i := 1; i < 16; i++ {
		code := gdextension.Shape((uint64(shape) >> (i * 4)) & 0xF)
		if code == gdextension.ShapeEmpty {
			break
		}
		align := uintptr(code.Alignment())
		offset = (offset + align - 1) &^ (align - 1)
		ptrs[i-1] = unsafe.Add(argptr, offset)
		offset += uintptr(code.SizeResult())
	}
	call(uintptr(method), uintptr(object), unsafe.Pointer(&ptrs[0]), unsafe.Pointer(&result))
	return result
}

// CallStatic invokes a static trivial method (no object receiver).
func CallStatic[T any](method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return Call[T](0, method, shape, args)
}

// call is the assembly trampoline that directly CALLs the ptrcall function
// pointer on the goroutine stack.
//
//go:nosplit
//go:noescape
func call(method, obj uintptr, args, ret unsafe.Pointer)
