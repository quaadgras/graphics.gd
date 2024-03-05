//go:build !generate

package gd

import (
	"reflect"
	"unsafe"

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

func (class Class[T, S]) Virtual(s string) reflect.Value {
	return class.super.Virtual(s)
}

func VirtualByName(class IsClass, name string) reflect.Value {
	return class.Virtual(name)
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T IsClass](godot Context, class IsClass) (T, bool) {
	if ref, ok := godot.API.Instances[mmm.Get(class.AsPointer())].(T); ok {
		return ref, true
	}
	var rtype = reflect.TypeOf([0]T{}).Elem()
	if rtype.Kind() == reflect.Pointer {
		rtype = rtype.Elem()
	}
	var classtag = godot.API.ClassDB.GetClassTag(godot.StringName(rtype.Name()))
	casted := godot.API.Object.CastTo(godot, class.AsObject(), classtag)
	if mmm.Get(casted.AsPointer()) != 0 {
		mmm.End(class.AsPointer()) // lifetime of the class has transferred to the return value.
		return (*(*T)(unsafe.Pointer(&casted))), true
	}
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
	obj.ptr = ptr

	API := mmm.API(ptr)

	ctx := newContext(API)
	defer ctx.End()

	// Important that we don't destroy RefCounted objects, instead
	// they should be unreferenced instead.
	ref := API.Object.CastTo(ctx, obj, API.refCountedClassTag)
	if mmm.Get(ref.AsPointer()) != 0 {
		(*(*RefCounted)(unsafe.Pointer(&ref))).Unreference()
	} else {
		API.Object.Destroy(obj)
	}
	mmm.End(ref.AsPointer())
	mmm.End(obj.AsPointer())
}

func (ptr Pointer) Virtual(string) reflect.Value {
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
	AsPointer() Pointer
	Virtual(string) reflect.Value
	AsObject() Object
}

type IsPointer interface {
	AsPointer() Pointer
}
