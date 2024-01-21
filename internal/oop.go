//go:build !generate

package gd

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

func (class Class[T, S]) AsObject() Object {
	var obj Object
	obj.SetPointer(class.super.AsPointer())
	return obj
}

func (class Class[T, S]) Pointer() uintptr {
	return class.super.Pointer()
}

// KeepAlive the class until the end of the specified context.
func (class Class[T, S]) KeepAlive(lt mmm.Lifetime) {
	mmm.Copy(class.super.AsPointer(), lt)
}

func (class Class[T, S]) Pin() Context {
	return Context{
		Lifetime: mmm.Life(class.AsPointer()),
		API:      mmm.API(class.AsPointer()),
	}
}

func (class Class[T, S]) class() S { return class.super }

func (class *Class[T, S]) Super() *S { return &class.super }

func (class Class[T, S]) virtual(s string) reflect.Value {
	return class.super.virtual(s)
}

func VirtualByName(class IsClass, name string) reflect.Value {
	return class.virtual(name)
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T IsClass](godot Context, class IsClass) (T, bool) {
	var rtype = reflect.TypeOf([0]T{}).Elem()
	if rtype.Kind() == reflect.Pointer {
		rtype = rtype.Elem()
	}
	var classtag = godot.API.ClassDB.GetClassTag(godot.StringName(rtype.Name()))
	godot.API.Object.CastTo(Context{Lifetime: mmm.Life(class.AsPointer()), API: godot.API}, class.AsObject(), classtag)
	var zero T
	return zero, false
}

type Singleton interface {
	IsSingleton()
}

type Pointer mmm.Pointer[API, Pointer, uintptr]

func (ptr Pointer) Pointer() uintptr {
	return mmm.Get(ptr)
}

func (ptr Pointer) Free() {
	var obj Object
	obj.super = ptr
	mmm.API(ptr).Object.Destroy(obj)
	mmm.End(obj.AsPointer())
}

func (ptr Pointer) virtual(string) reflect.Value {
	return reflect.Value{}
}

func (ptr Pointer) AsObject() Object {
	var obj Object
	obj.SetPointer(ptr)
	return obj
}

func (ptr Pointer) Pin() Context {
	return Context{
		Lifetime: mmm.Life(ptr),
		API:      mmm.API(ptr),
	}
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
	Pointer() uintptr
	AsPointer() Pointer
	virtual(string) reflect.Value

	Pin() Context
	AsObject() Object
}

type IsPointer interface {
	AsPointer() Pointer
}
