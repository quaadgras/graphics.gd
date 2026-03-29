// Package gdmemory provides safe access to engine memory that may belong in a different address space.
package gdmemory

import (
	"fmt"
	"reflect"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

// Pointer is a typed pointer to a value of type T in engine memory.
type Pointer[T any] struct {
	_   [0]func(T) // prevent comparison/conversion
	raw gdextension.Pointer
	rev uint64
	cap uintptr
}

// Uintptr returns the underlying pointer as a raw [uintptr] value.
func (ptr *Pointer[T]) Uintptr() uintptr {
	return uintptr(ptr.raw)
}

// Put the given value into the given pointer.
func (ptr *Pointer[T]) Put(val T) {
	copyIntoEngine(int(unsafe.Sizeof([1]T{})), unsafe.Pointer(&val), ptr.raw)
}

// Dereference the pointer, returning a copy of the value being
// pointed to.
func (ptr *Pointer[T]) Dereference() T {
	var value T
	CopyBufferToGo(ptr.raw, unsafe.Slice((*byte)(unsafe.Pointer(&value)), unsafe.Sizeof([1]T{})))
	return value
}

func safetyCheck(top, rtype reflect.Type) {
	switch rtype.Kind() {
	case reflect.Struct:
		for field := range rtype.Fields() {
			safetyCheck(top, field.Type)
		}
	case reflect.Array:
		safetyCheck(top, rtype.Elem())
	case reflect.Pointer, reflect.Chan:
		panic(fmt.Sprintf("Engine.Allocate[%s]: cannot allocate a value containing pointers/channels", top))
	}
}

// New is the implementation for Engine.Allocate[T]
func New[T comparable](val T) Pointer[T] {
	rtype := reflect.TypeOf(val)
	safetyCheck(rtype, rtype)
	cap := rtype.Size()
	return Pointer[T]{
		raw: CopyBufferToEngine(unsafe.Slice((*byte)(unsafe.Pointer(&val)), cap)),
		cap: cap,
	}
}

// WrapPointer wraps a raw engine pointer into a barrier-checked [Pointer].
func WrapPointer[T any](ptr gdextension.Pointer) Pointer[T] {
	return Pointer[T]{
		raw: ptr,
		rev: barrier,
		cap: unsafe.Sizeof([1]T{}),
	}
}

// UnwrapPointer extracts the raw engine pointer from a [Pointer].
func UnwrapPointer[T any](p Pointer[T]) gdextension.Pointer {
	if p.rev != barrier {
		panic("invalidated Engine.Pointer")
	}
	return p.raw
}
