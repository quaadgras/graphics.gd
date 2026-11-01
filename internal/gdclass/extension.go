package gdclass

import (
	"reflect"
	"sync"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdreference"
)

type Receiver unsafe.Pointer

// ReceiverOf extracts the instance pointer from a virtual-call receiver
// interface without going through reflect. Virtual receivers are always
// pointer-shaped (the registered *T), so the interface data word is the
// receiver itself.
func ReceiverOf(class any) Receiver {
	return Receiver((*[2]unsafe.Pointer)(unsafe.Pointer(&class))[1])
}

type Interface interface {
	superType() reflect.Type
	goType() reflect.Type
	getObject() [1]gdreference.Object
	Virtual(string) reflect.Value
}

type Pointer interface {
	gd.IsClass

	getObject() [1]gdreference.Object
	setObject([1]gdreference.Object)

	superType() reflect.Type
	goType() reflect.Type
}

func SuperType(class Interface) reflect.Type {
	return class.superType()
}

func GoType(class Interface) reflect.Type {
	return class.goType()
}

func SetObject(class Pointer, obj [1]gdreference.Object) {
	class.setObject(obj)
}

func GetObjectFromInterface(class Interface) [1]gdreference.Object {
	return class.getObject()
}

func GetObject(class Object) [1]gdreference.Object {
	return [1]gdreference.Object{class}
}

var Registered sync.Map

type Constructor interface {
	CreateInstanceFrom(reflect.Value, bool, bool) [1]gdreference.Object
}

type Extension[T Interface, S gd.IsClass] struct {
	gd.Class[T, S]
}

func (class Extension[T, S]) super() S {
	return class.Super()
}

// Deprecated: use the class-specific 'AsClass' method instead.
func (class *Extension[T, S]) Super() S {
	s := class.Class.Super()
	if *(*[1]gdreference.Object)(unsafe.Pointer(s)) == ([1]gdreference.Object{}) {
		class.createObject()
	}
	gdreference.UseObject((*gdreference.Object)(unsafe.Pointer(s)))
	return *s
}

// createObject lazily instantiates the engine-side object; kept out of
// [Extension.Super] so the hot already-created path stays small.
//
//go:noinline
func (class *Extension[T, S]) createObject() {
	impl, ok := Registered.Load(reflect.TypeFor[T]())
	if ok {
		instancer := impl.(Constructor)
		class.setObject(instancer.CreateInstanceFrom(reflect.NewAt(reflect.TypeFor[T](), unsafe.Pointer(class)), true, false))
	}
}

func (class Extension[T, S]) getObject() [1]gdreference.Object {
	return *(*[1]gdreference.Object)(unsafe.Pointer(class.Class.Super()))
}

func (class *Extension[T, S]) setObject(obj [1]gdreference.Object) {
	*(*[1]gdreference.Object)(unsafe.Pointer(class.Class.Super())) = obj
}

func (class Extension[T, S]) superType() reflect.Type {
	return reflect.TypeFor[S]()
}

func (class Extension[T, S]) goType() reflect.Type {
	return reflect.TypeFor[T]()
}

func (class *Extension[T, S]) AsObject() [1]gdreference.Object {
	obj := class.getObject()
	if obj == ([1]gdreference.Object{}) {
		class.createObject()
		obj = class.getObject()
	}
	gdreference.UseObject((*gdreference.Object)(unsafe.Pointer(class.Class.Super())))
	return obj
}

// FIXME can we remove this?
func (class *Extension[T, S]) UnsafePointer() unsafe.Pointer {
	return unsafe.Pointer(class)
}

type ExtensionInherits[S, T Interface] struct {
	_     [0]*T
	super S
}

// Super returns the underlying Super class (of type S).
func (class *ExtensionInherits[S, T]) Super() *S {
	return &class.super
}

func (class ExtensionInherits[S, T]) Virtual(s string) reflect.Value {
	return class.super.Virtual(s)
}

func (class ExtensionInherits[S, T]) getObject() [1]gdreference.Object {
	return class.super.getObject()
}

func (class ExtensionInherits[S, T]) superType() reflect.Type {
	return reflect.TypeFor[S]()
}

func (class ExtensionInherits[S, T]) goType() reflect.Type {
	return reflect.TypeFor[T]()
}

func (class *ExtensionInherits[S, T]) setObject(obj [1]gdreference.Object) {
	*(*[1]gdreference.Object)(unsafe.Pointer(&class.super)) = obj
}

func (class *ExtensionInherits[S, T]) AsObject() [1]gdreference.Object {
	obj := class.getObject()
	if obj == ([1]gdreference.Object{}) {
		impl, ok := Registered.Load(reflect.TypeFor[T]())
		if ok {
			instancer := impl.(Constructor)
			obj = instancer.CreateInstanceFrom(reflect.NewAt(reflect.TypeFor[T](), unsafe.Pointer(class)), true, false)
			class.setObject(obj)
		}
	}
	gdreference.UseObject((*gdreference.Object)(unsafe.Pointer(&class.super)))
	return obj
}

// OnRegisterShader is set by the shaders package, classdb calls it when a
// ShaderMaterial extension (a shader written in Go) is registered, with the
// constructor of the class, so that its code can be generated in the
// background, before the first material is created.
var OnRegisterShader func(construct func() reflect.Value)
