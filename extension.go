//go:build !generate

package gd

import "C"

import (
	"reflect"
	"sync/atomic"
	"unsafe"
)

// Instance of a builtin Godot object or an [ExtensionInstance].
type Instance interface {
	object() safeObject
	className() string
	virtual(reflect.Type, string) (reflect.Method, bool)

	// Free the instance, releasing any underlying resources.
	// May panic if the [Instance] is an [Extension] that is
	// leaking memory.
	Free()
}

// ExtensionInstance of a registered [Extension].
type ExtensionInstance interface {
	Instance

	extension() *Extension
}

// BelongsTo returns an [InstanceOwner] corresponding to the
// given [ExtensionInstance].
func BelongsTo(instance ExtensionInstance) InstanceOwner {
	return InstanceOwner{
		children: &instance.extension().count,
	}
}

// InstanceOwner refers to an owner of an [Instance].
// Exists primarily as a memory-leak detector.
type InstanceOwner struct {
	children *atomic.Int64 // pointer to children count.
}

// Extension can be added as a field in a Go struct to order to enable
// that struct to be registered as a Godot extension that can create
// new [Extension] instances.
type Extension struct {
	class extensionClassID
	index extensionClassInstanceID
	count atomic.Int64  // used as the InstanceOwner children count.
	owner InstanceOwner // InstanceOwner of the Extension's instance (parent).

	gdRef safeObject
}

func (ex Extension) virtual(reflect.Type, string) (method reflect.Method, ok bool) { return }
func (ex Extension) object() safeObject                                            { return ex.gdRef }
func (ex Extension) className() string                                             { return "" }

func (ex *Extension) extension() *Extension { return ex }

// Free the [Extension]'s instance, releasing any underlying resources.
// May panic if the [Extension]'s instance is leaking memory.
func (ex Extension) Free() {
	ex.gdRef.free()
}

// extensionClassID identifies an [Extension] class that has been
// registered with Godot. Whenever we register a class, we create
// a global [extensionClassHandler] that handles any Godot -> Go
// calls for any instances of that class.
type extensionClassID uint32

// extensionClassInstanceID is a unique identifier for an instance of
// an [Extension] class.
type extensionClassInstanceID uint32

// extensionClassDB is the [extensionClassDatabase] singleton.
var extensionClassDB extensionClassDatabase

// extensionClassDatabase contains registered [Extension] classes.
// [extensionClassID] will be a 1-based index into this array.
type extensionClassDatabase []extensionClassHandler

// add a new [Extension] class to the database.
func (db *extensionClassDatabase) add(extension extensionClassHandler) extensionClassID {
	*db = append(*db, extension)
	return extensionClassID(len(*db))
}

// get the [extensionClassHandler] for the given [extensionClassID].
func (db extensionClassDatabase) get(id extensionClassID) extensionClassHandler { return db[id-1] }

// extensionClassInstances stores a list of [ExtensionInstance]s.
type extensionClassInstances[Type Instance] []*Type

func (instances extensionClassInstances[T]) get(id extensionClassInstanceID) *T {
	return instances[id-1]
}

// extensionClassHandlerFor a given [Extension] that handles Godot -> Go function calls.
type extensionClassHandlerFor[Type Instance, Extends Instance] struct {
	id extensionClassID

	class string // class name, cached on [Register].

	vtype reflect.Type // [Instance] value type.
	ptype reflect.Type // [Instance] pointer type.

	build InstanceBuilder[Type, Extends]

	slice extensionClassInstances[Type] // slice of instantiated instances.
}

// extensionClassHandler that handles Godot -> Go function calls.
type extensionClassHandler interface {
	create() cObject
	delete(extensionClassInstanceID)
	lookup(name string) (reflect.Method, bool)
	call(instanceID extensionClassInstanceID, method uint8, args *unsafe.Pointer, result unsafe.Pointer)
}

// InstanceBuilder is a function that intializes a new [Instance] of a particular
// [Extension] 'Type'. This function should return a pointer to the [Instance]'s
// [Extension] field and a pointer to the builtin Godot object that the [Instance]
// is extending.
type InstanceBuilder[Type Instance, Extends Instance] func(*Type) (*Extension, Extends)

/*
Register a new [Extension] type with Godot, an [InstanceBuilder] should be provided
that will be used to initialize new instances of the [Extension]. Register returns
both New and NewAt constructors. These should be used to create new instances of
the [Extension] type.

Example:

	type HelloWorld struct {
		gd.Extension

		object gd.Object
	}

	var NewHelloWorld, NewHelloWorldAt = gd.Register(func(hello *HelloWorld) (*gd.Extension, *gd.Object) {
		return &hello.Extension, gd.NewObjectAt(&hello.object, gd.BelongsTo(hello))
	})
*/
func Register[Type Instance, Extends Instance](fn InstanceBuilder[Type, Extends]) (new func(InstanceOwner) *Type, newAt func(*Type, InstanceOwner) *Type) {
	var zero Type
	var parent Extends

	var rtype = reflect.TypeOf(zero)

	var extension = extensionClassHandlerFor[Type, Extends]{
		build: fn,

		vtype: rtype,
		ptype: reflect.PtrTo(rtype),
		class: rtype.Name() + "\000",
	}
	extension.id = extensionClassDB.add(&extension)

	deferred = append(deferred, func() {
		classDB.register_extension_class(extension.class, parent.className(), uintptr(extension.id))
	})

	return func(ctx InstanceOwner) *Type {
			_, ptr := extension.new(ctx, nil)
			return ptr
		}, func(dst *Type, ctx InstanceOwner) *Type {
			_, ptr := extension.new(ctx, dst)
			return ptr
		}
}

type instanceID struct {
	class extensionClassID
	index extensionClassInstanceID
}

func loadInstanceID(id uintptr) instanceID {
	return instanceID{
		class: extensionClassID(id >> 32),
		index: extensionClassInstanceID(id),
	}
}

func (id instanceID) pack() uintptr {
	return uintptr(id.class)<<32 | uintptr(id.index)
}

func (ext *extensionClassHandlerFor[Type, Extends]) lookup(name string) (reflect.Method, bool) {
	var zero Type
	var core Extends
	return core.virtual(reflect.PtrTo(reflect.TypeOf(zero)), name)
}

func (ext *extensionClassHandlerFor[Type, Extends]) new(ctx InstanceOwner, dst *Type) (*Extension, *Type) {
	var id instanceID
	id.class = ext.id

	if dst == nil {
		dst = new(Type)
	}

	// reuse existing slot if possible.
	for i, slot := range ext.slice {
		if slot == nil {
			id.index = extensionClassInstanceID(i + 1)
			ext.slice[i] = dst
			break
		}
	}

	if id.index == 0 {
		ext.slice = append(ext.slice, dst)
		id.index = extensionClassInstanceID(len(ext.slice))
	}

	instance, base := ext.build(ext.slice[id.index-1])
	instance.class = ext.id
	instance.index = id.index
	instance.owner = ctx

	obj := base.object()
	obj.get().set_instance(ext.class, id.pack())

	instance.gdRef = obj

	return instance, ext.slice[id.index-1]
}

var belongsToGodot = InstanceOwner{
	children: new(atomic.Int64),
}

func (ext *extensionClassHandlerFor[Type, Extends]) create() cObject {
	instance, _ := ext.new(belongsToGodot, nil)
	return instance.gdRef.get()
}

func (ext *extensionClassHandlerFor[Type, Extends]) delete(id extensionClassInstanceID) {
	obj := (*ext.slice[id-1]).object()
	obj.destroy()
	(*ext.slice[id-1]).Free()
	ext.slice[id-1] = nil
}

func (ext *extensionClassHandlerFor[Type, Extends]) call(instance extensionClassInstanceID, index uint8, args *unsafe.Pointer, result unsafe.Pointer) {
	rvalue := reflect.ValueOf(ext.slice[instance-1])
	method := rvalue.Method(int(index) - 1)
	rtype := method.Type()

	buffer := make([]reflect.Value, rtype.NumIn())

	if rtype.NumIn() > 1 {
		slice := unsafe.Slice(args, rtype.NumIn())
		for i := range slice {
			buffer[i] = reflect.New(rtype.In(i + 1)).Elem()
			into(buffer[i].Interface(), slice[i])
		}
	}

	results := method.Call(buffer)
	if len(results) > 1 {
		rtype := rvalue.Type()
		panic(rtype.String() + "." + rtype.Method(int(index)-1).Name + " returns too many values")
	}
	if len(results) == 1 {
		toResult(results[0].Interface(), result)
	}
}

func (ext *extensionClassHandlerFor[Type, Extends]) methods() []reflect.Method {
	var methods []reflect.Method
	for i := 0; i < ext.ptype.NumMethod(); i++ {
		methods = append(methods, ext.ptype.Method(i))
	}
	return methods
}
