package main

type isClass interface {
	pointer() uintptr
}

type rawPointer uintptr

func (ptr rawPointer) pointer() uintptr {
	return uintptr(ptr)
}

type Class[T, S isClass] struct {
	_     [0]*T
	super S
}

func (class Class[T, S]) pointer() uintptr {
	return class.super.pointer()
}
func (class Class[T, S]) Super() *S { return &class.super }

type Animal struct {
	Class[Animal, rawPointer]

	kingdom string
}

type Dog struct {
	Class[Dog, Animal]
}

func main() {
	var dog Dog
	var animal = dog.Super()
	animal.kingdom = "Animalia"
}
