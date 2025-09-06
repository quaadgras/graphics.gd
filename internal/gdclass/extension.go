package gdclass

import (
	"reflect"
	"sync"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

type Receiver unsafe.Pointer

type Interface interface {
	superType() reflect.Type
	getObject() [1]gd.Object
	Virtual(string) reflect.Value
}

type Pointer interface {
	gd.IsClass

	getObject() [1]gd.Object
	setObject([1]gd.Object)

	superType() reflect.Type
}

func SuperType(class Interface) reflect.Type {
	return class.superType()
}

func SetObject(class Pointer, obj [1]gd.Object) {
	class.setObject(obj)
}

func GetObject(class Interface) [1]gd.Object {
	return class.getObject()
}

var Registered sync.Map

type Constructor interface {
	CreateGoInstanceFrom(reflect.Value, bool) [1]gd.Object
}

type Extension[T Interface, S gd.IsClass] struct {
	gd.Class[T, S]
}

func (class Extension[T, S]) super() S {
	return class.Super()
}

// Deprecated: use the class-specific 'AsClass' method instead.
func (class *Extension[T, S]) Super() S {
	class.AsObject()
	return *class.Class.Super()
}

func (class Extension[T, S]) getObject() [1]gd.Object {
	return *(*[1]gd.Object)(unsafe.Pointer(class.Class.Super()))
}

func (class *Extension[T, S]) setObject(obj [1]gd.Object) {
	*(*[1]gd.Object)(unsafe.Pointer(class.Class.Super())) = obj
}

func (class Extension[T, S]) superType() reflect.Type {
	return reflect.TypeFor[S]()
}

func (class *Extension[T, S]) AsObject() [1]gd.Object {
	obj := class.getObject()
	if obj == ([1]gd.Object{}) {
		impl, ok := Registered.Load(reflect.TypeFor[T]())
		if ok {
			instancer := impl.(Constructor)
			obj = instancer.CreateGoInstanceFrom(reflect.NewAt(reflect.TypeFor[T](), unsafe.Pointer(class)), true)
			class.setObject(obj)
		}
	}
	pointers.Get(obj[0])
	return obj
}

func (class *Extension[T, S]) UnsafePointer() unsafe.Pointer {
	return unsafe.Pointer(class)
}
