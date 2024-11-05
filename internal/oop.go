//go:build !generate

package gd

import (
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"grow.graphics/gd/internal/pointers"
)

var ExtensionInstances sync.Map

type NotificationType int32

func PointerWithOwnershipTransferredToGo(ptr [1]uintptr) Object {
	return pointers.New[Object]([3]uintptr{ptr[0]})
}

func PointerWithOwnershipTransferredToGodot(ptr Object) uintptr {
	raw, _ := pointers.End(ptr)
	if raw[1] != 0 {
		panic("illegal transfer of ownership from Go -> Godot")
	}
	return raw[0]
}

func PointerMustAssertInstanceID(ptr [1]uintptr) Object {
	if ptr == ([1]uintptr{}) {
		return Object{}
	}
	initial := pointers.New[Object]([3]uintptr{ptr[0]})
	pointers.Set(&initial, [3]uintptr{ptr[0], uintptr(Global.Object.GetInstanceID(initial))})
	return initial
}

func PointerLifetimeBoundTo(obj Object, ptr [1]uintptr) Object {
	return pointers.New[Object]([3]uintptr{ptr[0], 0})
}

func (self Object) AsObject() Object {
	return self
}

func (self RefCounted) Free() {
	_, ok := pointers.End(self)
	if !ok {
		return
	}
}

func (self Object) Free() {
	_, ok := pointers.End(self)
	if !ok {
		return
	}
	// Important that we don't destroy RefCounted objects, instead
	// they should be unreferenced instead.
	ref := Global.Object.CastTo(self, Global.refCountedClassTag)
	if ref != (Object{}) {
		if RefCounted(ref).Unreference() {
			Global.Object.Destroy(self)
		}
	} else {
		Global.Object.Destroy(self)
	}
}

type Class[T any, S IsClass] struct {
	_     [0]*T
	super S
}

func (class *Class[T, S]) AsObject() Object {
	return class.super.AsObject()
}

func (class Class[T, S]) class() S { return class.super } //lint:ignore U1000 false positive.

func (class *Class[T, S]) Super() *S { return &class.super }

func (class Class[T, S]) Virtual(s string) reflect.Value {
	return class.super.Virtual(s)
}

func VirtualByName(class IsClass, name string) reflect.Value {
	return class.Virtual(name)
}

func classNameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr || rtype.Kind() == reflect.Array {
		return classNameOf(rtype.Elem())
	}
	if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
		if rtype.Field(0).Anonymous {
			if rename, ok := rtype.Field(0).Tag.Lookup("gd"); ok {
				return rename
			}
		}
		if rtype.Name() == "" && rtype.Field(0).Anonymous {
			return classNameOf(rtype.Field(0).Type)
		}
		return strings.TrimPrefix(rtype.Name(), "class")
	}
	return ""
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T IsClass](class IsClass) (T, bool) {
	/*if ref, ok := Global.Instances[pointers.Get(class.AsObject())[0]].(T); ok {
	extension := any(ref).(ExtensionClass)
	extension.SetPointer(class.AsObject())
	return ref, true
	}*/
	var zero T
	var rtype = reflect.TypeOf([0]T{}).Elem()
	if rtype.Kind() == reflect.Pointer {
		return zero, false
	}
	var classtag = Global.ClassDB.GetClassTag(NewStringName(classNameOf(rtype)))
	casted := Global.Object.CastTo(class.AsObject(), classtag)
	if casted != (Object{}) && pointers.Get(casted) != ([3]uintptr{}) {
		return (*(*T)(unsafe.Pointer(&casted))), true
	}
	return zero, false
}

func as[T any](v Variant) T {
	var zero T
	val := v.Interface()
	if val == nil {
		return zero
	}
	if obj, ok := val.(IsClass); ok {
		var classtag = Global.ClassDB.GetClassTag(obj.AsObject().GetClass().StringName())
		casted := Global.Object.CastTo(obj.AsObject(), classtag)
		if casted != (Object{}) && pointers.Get(casted) != ([3]uintptr{}) {
			any(&zero).(PointerToClass).SetPointer(casted)
		}
		return zero
	}
	return val.(T)
}

type Singleton interface {
	IsSingleton()
}

type Extends[T IsClass] interface {
	class() T

	Virtual(string) reflect.Value
}

type PointerToClass interface {
	IsClass
	SetPointer(Object)
}

type IsClass interface {
	Virtual(string) reflect.Value
	AsObject() Object
}
