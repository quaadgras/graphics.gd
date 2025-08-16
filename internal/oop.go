//go:build !generate

package gd

import (
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"graphics.gd/internal/pointers"
)

var ExtensionInstances sync.Map

func init() {
	RegisterCleanup(func() {
		ExtensionInstances.Range(func(key, value any) bool {
			pointers.Raw[Object]([3]uint64{key.(uint64)}).Free()
			return true
		})
	})
}

type NotificationType int32

func PointerWithOwnershipTransferredToGo[T pointers.Generic[T, [3]uint64]](ptr EnginePointer) T {
	if ptr == 0 {
		return T{}
	}
	if runtime.GOOS == "js" {
		// See https://github.com/quaadgras/graphics.gd/issues/103
		// in some cases, objects will crash on free after calling
		// certain methods if we don't call a known-safe method on
		// them first. Probably has something to do with complex
		// cross-language call stacks, ie.
		//
		// 	C -> JS -> Go -> JS -> C -> JS -> Go
		//
		// Maybe emscripten needs a chance to properly allocate the
		// object or something? or there could be a serious memory
		// corruption bug somewhere.
		pointers.Raw[Object]([3]uint64{uint64(ptr), 0}).IsBlockingSignals()
	}
	return pointers.New[T]([3]uint64{uint64(ptr)})
}

func PointerBorrowedTemporarily[T pointers.Generic[T, [3]uint64]](ptr EnginePointer) T {
	if ptr == 0 {
		return T{}
	}
	return pointers.Let[T]([3]uint64{uint64(ptr)})
}

func PointerWithOwnershipTransferredToGodot[T pointers.Generic[T, [3]uint64]](ptr T) EnginePointer {
	raw := pointers.Get(ptr)
	pointers.Set(ptr, [3]uint64{raw[0], uint64(Global.Object.GetInstanceID([1]Object{pointers.Raw[Object]([3]uint64{raw[0]})}))})
	pointers.Lay(ptr)
	if raw[1] != 0 {
		panic("illegal transfer of ownership from Go -> Godot")
	}
	return EnginePointer(raw[0])
}

func PointerMustAssertInstanceID[T pointers.Generic[T, [3]uint64]](ptr EnginePointer) T {
	if ptr == 0 {
		return T{}
	}
	return pointers.Let[T]([3]uint64{uint64(ptr), uint64(Global.Object.GetInstanceID([1]Object{pointers.Raw[Object]([3]uint64{uint64(ptr)})}))})
}

func PointerLifetimeBoundTo[T pointers.Generic[T, [3]uint64]](obj [1]Object, ptr EnginePointer) T {
	if ptr == 0 {
		return T{}
	}
	return pointers.Let[T]([3]uint64{uint64(ptr), 0})
}

func ObjectChecked(obj [1]Object) EnginePointer {
	raw := pointers.Get(obj[0])
	if !obj[0].IsAlive(raw) {
		panic("use after free")
	}
	return EnginePointer(raw[0])
}

func (self Object) AsObject() [1]Object {
	return [1]Object{self}
}

func (self RefCounted) Free() {
	_, ok := pointers.End(Object(self))
	if !ok {
		return
	}
}

func (self Object) IsAlive(raw [3]uint64) bool {
	return raw[1] == 0 || Global.Object.GetInstanceFromID(ObjectID(raw[1])) != [1]Object{}
}

func (self Object) Free() {
	raw, ok := pointers.End(self)
	if !ok {
		return
	}
	this := [1]Object{pointers.Raw[Object](raw)}
	//fmt.Println("FREE ", pointers.Trio[Object](self), this[0].GetClass().String())
	//	fmt.Println(runtime.Caller(2))
	// Important that we don't destroy RefCounted objects, instead
	// they should be unreferenced instead.
	ref := Global.Object.CastTo(this, Global.refCountedClassTag)
	if ref != ([1]Object{}) {
		if (*(*RefCounted)(unsafe.Pointer(&ref))).Unreference() {
			Global.Object.Destroy(this)
		}
	} else {
		Global.Object.Destroy(this)
	}
}

type Class[T any, S IsClass] struct {
	_     [0]*T
	super S
}

func (class *Class[T, S]) AsObject() [1]Object {
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
	if casted != ([1]Object{}) && pointers.Get(casted[0]) != ([3]uint64{}) {
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
		var classtag = Global.ClassDB.GetClassTag(obj.AsObject()[0].GetClass().StringName())
		casted := Global.Object.CastTo(obj.AsObject(), classtag)
		if casted != ([1]Object{}) && pointers.Get(casted[0]) != ([3]uint64{}) {
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
	SetPointer([1]Object)
}

type IsClass interface {
	Virtual(string) reflect.Value
	AsObject() [1]Object
}
