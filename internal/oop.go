//go:build !generate

package gd

import (
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"graphics.gd/internal/gdextension"
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

func PointerWithOwnershipTransferredToGo[T pointers.Generic[T, [3]uint64]](ptr gdextension.Object) T {
	if ptr == 0 {
		return T{}
	}
	return pointers.New[T]([3]uint64{uint64(ptr)})
}

func PointerBorrowedTemporarily[T pointers.Generic[T, [3]uint64]](ptr gdextension.Object) T {
	if ptr == 0 {
		return T{}
	}
	return pointers.Let[T]([3]uint64{uint64(ptr)})
}

func PointerWithOwnershipTransferredToGodot[T pointers.Generic[T, [3]uint64]](ptr T) EnginePointer {
	raw := pointers.Get(ptr)
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(gdextension.Object(raw[0]), gdextension.CallReturns[gdextension.ObjectID](unsafe.Pointer(&id)))
	pointers.Set(ptr, [3]uint64{raw[0], uint64(id)})
	pointers.Lay(ptr)
	if raw[1] != 0 {
		panic("illegal transfer of ownership from Go -> Godot")
	}
	return EnginePointer(raw[0])
}

func PointerQueueFree[T pointers.Generic[T, [3]uint64]](ptr T) {
	raw := pointers.Get(ptr)
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(gdextension.Object(raw[0]), gdextension.CallReturns[gdextension.ObjectID](unsafe.Pointer(&id)))
	pointers.Set(ptr, [3]uint64{raw[0], uint64(id)})
	pointers.Lay(ptr)
}

func PointerMustAssertInstanceID[T pointers.Generic[T, [3]uint64]](ptr gdextension.Object) T {
	if ptr == 0 {
		return T{}
	}
	var id gdextension.ObjectID
	gdextension.Host.Objects.ID.Get(gdextension.Object(ptr), gdextension.CallReturns[gdextension.ObjectID](unsafe.Pointer(&id)))
	return pointers.Let[T]([3]uint64{uint64(ptr), uint64(id)})
}

func PointerLifetimeBoundTo[T pointers.Generic[T, [3]uint64]](obj [1]Object, ptr gdextension.Object) T {
	if ptr == 0 {
		return T{}
	}
	return pointers.Let[T]([3]uint64{uint64(ptr), 0})
}

func CallerIncrements(obj [1]Object) gdextension.Object {
	RefCounted(obj[0]).Reference()
	return ObjectChecked([1]Object{Object(obj[0])})
}

func ObjectChecked(obj [1]Object) gdextension.Object {
	raw := pointers.Get(obj[0])
	if !obj[0].IsAlive(raw) {
		panic("use after free")
	}
	if raw == [3]uint64{} {
		return 0
	}
	return gdextension.Object(raw[0])
}

func (self Object) AsObject() [1]Object {
	return [1]Object{self}
}

func (class Object) Virtual(s string) reflect.Value {
	return reflect.Value{}
}

func (self RefCounted) AsObject() [1]Object {
	return *(*[1]Object)(unsafe.Pointer(&self))
}

func (class RefCounted) Virtual(s string) reflect.Value {
	return reflect.Value{}
}

func (self RefCounted) Free() {
	_, ok := pointers.End(Object(self))
	if !ok {
		return
	}
}
func (self *RefCounted) SetObject(obj [1]Object) bool {
	ref := gdextension.Host.Objects.Cast(gdextension.Object(pointers.Get(obj[0])[0]), Global.refCountedClassTag)
	if ref != 0 {
		*self = RefCounted(obj[0])
		return true
	}
	return false
}

func (self Object) IsAlive(raw [3]uint64) bool {
	return raw[1] == 0 || gdextension.Host.Objects.Lookup(gdextension.ObjectID(raw[1])) != 0
}

func (self Object) Free() {
	raw, ok := pointers.End(self)
	if !ok {
		return
	}
	//fmt.Println(raw)
	//fmt.Fprintln(os.Stderr, "FREE ", pointers.Trio[Object](self), pointers.Raw[Object](raw).GetClass().String())
	//fmt.Println(runtime.Caller(2))
	ref := gdextension.Host.Objects.Cast(gdextension.Object(raw[0]), Global.refCountedClassTag)
	if ref != 0 {
		// Important that we don't destroy RefCounted objects, instead
		// they should be unreferenced instead.
		if last := RefCounted(pointers.Raw[Object]([3]uint64{uint64(ref)})).Unreference(); !last {
			//fmt.Println("not last reference, not freeing")
			return
		}
	}
	gdextension.Host.Objects.Unsafe.Free(gdextension.Object(raw[0]))
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

type IsClassCastable interface {
	SetObject([1]Object) bool
}
