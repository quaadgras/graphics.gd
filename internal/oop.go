//go:build !generate

package internal

import (
	"reflect"

	"runtime.link/mmm"
)

type Class[T, S IsClass] struct {
	_     [0]*T
	super S
}

func (class *Class[T, S]) SetPointer(ptr Pointer) {
	any(&class.super).(PointerToClass).SetPointer(ptr)
}

func (class Class[T, S]) AsPointer() Pointer {
	return class.super.AsPointer()
}

func (class Class[T, S]) Context() mmm.Context {
	return class.super.Context()
}

func (class Class[T, S]) Pointer() uintptr {
	return class.super.Pointer()
}

// KeepAlive the class until the end of the specified context.
func (class Class[T, S]) KeepAlive(ctx mmm.Context) {
	mmm.Move(class.super.AsPointer(), ctx)
}

func (class Class[T, S]) class() S { return class.super }

func (class *Class[T, S]) Super() *S { return &class.super }

func (class Class[T, S]) virtual(s string) reflect.Value {
	return class.super.virtual(s)
}

func VirtualByName(class IsClass, name string) reflect.Value {
	return class.virtual(name)
}

type Singleton interface {
	isSingleton()
}

type Pointer mmm.Pointer[API, Pointer, uintptr]

func (ptr Pointer) Free() {
	ptr.API.Object.Destroy(ptr.Pointer())
}

func (ptr Pointer) virtual(string) reflect.Value {
	return reflect.Value{}
}

func (ptr Pointer) AsPointer() Pointer     { return ptr }
func (ptr *Pointer) SetPointer(to Pointer) { *ptr = to }

type Extends[T IsClass] interface {
	IsClass
	class() T
}

type PointerToClass interface {
	IsClass
	SetPointer(Pointer)
}

type IsClass interface {
	Context() mmm.Context
	Pointer() uintptr
	AsPointer() Pointer
	virtual(string) reflect.Value
}

type IsPointer interface {
	AsPointer() Pointer
}

// MarkFree marks the given class as being freed.
func MarkFree(class IsClass) {
	mmm.MarkFree(class.AsPointer())
}
