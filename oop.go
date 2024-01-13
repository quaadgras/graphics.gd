package gd

type Class[T, S IsClass] struct {
	_     [0]*T
	super S
}

func (class *Class[T, S]) alloc() {
	any(&class.super).(interface{ alloc() }).alloc()
}

func (class Class[T, S]) getPointer() Pointer {
	return class.super.getPointer()
}
func (class Class[T, S]) class() S { return class.super }

func (class *Class[T, S]) Super() *S { return &class.super }

type Singleton interface {
	isSingleton()
}

type Pointer struct {
	Val [2]uintptr
	API *API
	src *Pointer
}

func (ptr *Pointer) alloc()             { ptr.src = ptr }
func (ptr Pointer) getPointer() Pointer { return ptr }

func MakePointer(class ClassContainer, ptr Pointer) {
	class.alloc()
	class.getPointer().src.Val = ptr.Val
	class.getPointer().src.API = ptr.API
}

func InitPointer(class ClassContainer, ptr Pointer) {
	class.alloc()
}

func EditPointer(class ClassReference, ptr Pointer) {
	class.getPointer().src.Val = ptr.Val
	class.getPointer().src.API = ptr.API
}

func ReadPointer(class ClassReference) [2]uintptr {
	ptr := class.getPointer()
	if ptr.src == nil {
		return [2]uintptr{}
	}
	return ptr.src.Val
}

func FreePointer(class ClassReference) {
	class.getPointer().src.Val = [2]uintptr{}
}

type Extends[T IsClass] interface {
	IsClass
	class() T
}

type ClassReference interface {
	getPointer() Pointer
}

type ClassContainer interface {
	ClassReference

	alloc()
}

type IsClass interface {
	getPointer() Pointer
}
