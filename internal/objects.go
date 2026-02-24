//go:build !generate

package gd

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/ring"
)

type ExtensionClassCallVirtualFunc func(any, gdextension.Pointer, gdextension.Pointer)

var ExtensionInstanceLookup func(gdextension.Object) any
var ExtensionInstanceGoOnly func(gdextension.Object, bool) (gdreference.Object, bool)

type NotificationType int32

func PointerWithOwnershipTransferredToGo(ptr gdextension.Object) gdreference.Object {
	if ptr == 0 {
		return gdreference.Object{}
	}
	if obj, ok := ExtensionInstanceGoOnly(ptr, true); ok {
		return obj
	}
	return gdreference.OwnObject(ptr, Free)
}

func PointerBorrowedTemporarily(ptr gdextension.Object) gdreference.Object {
	if ptr == 0 {
		return gdreference.Object{}
	}
	return gdreference.LetObject(ptr)
}

func PointerWithOwnershipTransferredToGodot(obj gdreference.Object) EnginePointer {
	raw, ok := gdreference.EndObject(obj)
	if !ok {
		panic("illegal transfer of ownership from Go -> Godot")
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(raw, gdextension.CallReturns[gdextension.ObjectID](unsafe.Pointer(&id)))
	ExtensionInstanceGoOnly(raw, false)
	return EnginePointer(raw)
}

func PointerQueueFree(ptr gdreference.Object) {
	gdreference.EndObject(ptr)
}

func PointerMustAssertInstanceID[T pointers.Generic[T, [3]uint64]](ptr gdextension.Object) T {
	if ptr == 0 {
		return T{}
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(gdextension.Object(ptr), gdextension.CallReturns[gdextension.ObjectID](unsafe.Pointer(&id)))
	return pointers.Let[T]([3]uint64{uint64(ptr), uint64(id)})
}

func PointerLifetimeBoundTo(obj [1]gdreference.Object, ptr gdextension.Object) gdreference.Object {
	if ptr == 0 {
		return gdreference.Object{}
	}
	return gdreference.LetObject(ptr)
}

func CallerIncrements(obj [1]gdreference.Object) gdextension.Object {
	RefCounted(obj[0]).Reference()
	return ObjectChecked([1]gdreference.Object{gdreference.Object(obj[0])})
}

func ObjectChecked(obj [1]gdreference.Object) gdextension.Object {
	raw := gdreference.GetObject(obj[0])
	if raw == 0 {
		panic("use of an invalid reference (please read https://the.graphics.gd/guide/memory)")
	}
	return raw
}

func (self RefCounted) AsObject() [1]gdreference.Object {
	return *(*[1]gdreference.Object)(unsafe.Pointer(&self))
}

func (class RefCounted) Virtual(s string) reflect.Value {
	return reflect.Value{}
}

func (self RefCounted) Free() {
	ObjectFree(self.AsObject()[0])
}

func (self *RefCounted) SetObject(obj [1]gdreference.Object) bool {
	ref := gdextension.Host.Objects.Cast(gdreference.GetObject(obj[0]), refCountedClassTag)
	if ref != 0 {
		*self = RefCounted(obj[0])
		return true
	}
	return false
}

func ObjectIsAlive(raw [3]uint64) bool {
	return raw[1] == 0 || gdextension.Host.Objects.Lookup(gdextension.ObjectID(raw[1])) != 0
}

var debugFree = strings.Contains(os.Getenv("GDDEBUG"), "free")

func Free(raw gdextension.Object) {
	if raw == 0 {
		return
	}
	ref := gdextension.Host.Objects.Cast(raw, refCountedClassTag)
	if ref != 0 {
		// Important that we don't destroy RefCounted objects, instead
		// they should be unreferenced instead.
		if last := RefCounted(gdreference.RawObject(ref)).Unreference(); !last {
			return
		}
	}
	if debugFree {
		fmt.Println(raw)
		fmt.Fprintln(os.Stderr, "FREE ", ObjectGetClass(gdreference.RawObject(raw)).String())
		fmt.Println(runtime.Caller(2))
	}
	ring.Main.Flush()
	gdextension.Host.Objects.Unsafe.Free(raw)
}

func ObjectFree(self gdreference.Object) {
	raw, ok := gdreference.EndObject(self)
	if !ok {
		return
	}
	Free(raw)
}

type Class[T any, S IsClass] struct {
	_     [0]*T
	super S
}

func (class *Class[T, S]) AsObject() [1]gdreference.Object {
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

type Singleton interface {
	IsSingleton()
}

type Extends[T IsClass] interface {
	class() T

	Virtual(string) reflect.Value
}

type PointerToClass interface {
	IsClass
	SetPointer([1]gdreference.Object)
}

type IsClass interface {
	Virtual(string) reflect.Value
	AsObject() [1]gdreference.Object
}

type IsClassCastable interface {
	SetObject([1]gdreference.Object) bool
}
