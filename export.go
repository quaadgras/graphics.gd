//go:build !generate

package gd

import (
	"reflect"
	"strings"
	"sync/atomic"
	"unsafe"
)

// Extension can be added as a field in a Go struct to order to enable
// that struct to be registered as a Godot extension that can create
// new [Extension] instances.
type Extension struct {
	class extensionClassID
	index extensionClassInstanceID
	count atomic.Int64 // used as the InstanceOwner children count.

	// depreciated
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

type extension interface {
	extension() *Extension
}

type exportedMethod struct {
	reflect.Method

	Arguments []string
	Result    string
}

var exportedMethods []exportedMethod

/*
	 Scripting psuedo-type used to define which methods are available to
	 the engine. For example:

		type MyExtension struct {
			gd.Extension
			gd.Scripting[struct{
				DoSomething gd.Method `gd:"do_something"`
			}]
		}]

		func (MyExtension) DoSomething() {
			// ...
		}
*/
type Scripting[T comparable] struct{}

type scripting interface {
	register(reflect.Type, string)
}

func (scripting Scripting[T]) register(extension reflect.Type, class string) {
	var zero T
	rtype := reflect.TypeOf(zero)
	if rtype.Kind() != reflect.Struct {
		panic("gd: Scripting type must be a struct")
	}

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)

		switch field.Type {
		case reflect.TypeOf(Method{}):
			method, ok := reflect.PtrTo(extension).MethodByName(field.Name)
			if !ok {
				panic("gd: Scripting method not found: " + field.Name + " (" + extension.Name() + ")")
			}

			tag := field.Tag.Get("gd")
			if tag == "" {
				panic("gd: Scripting method must have a gd tag: " + field.Name)
			}

			name, nargs, ok := strings.Cut(tag, "(")
			if !ok || len(nargs) == 0 {
				panic("gd: Scripting method must have a gd tag of form 'function_name(arg_name1, arg_name2) result_name'")
			}
			nargs, result, ok := strings.Cut(nargs, ")")
			if !ok {
				panic("gd: Scripting method must have a gd tag of form 'function_name(arg_name1, arg_name2) result_name'")
			}
			args := strings.Split(nargs, ",")

			exportedMethods = append(exportedMethods, exportedMethod{
				Method: method,

				Arguments: args,
				Result:    result,
			})

			classDB.register_extension_class_method(class, &cExtensionClassMethodInfo{
				Name:           name,
				MethodUserData: uintptr(len(exportedMethods)),
				ArgumentCount:  uint32(method.Type.NumIn() - 1), // exclude receiver.
				HasReturnValue: method.Type.NumOut() > 0,
			})

		default:
			panic("gd: Scripting type must only contain gd.Method fields")
		}
	}
}

// Method can be used in a [Scripting] struct to indicate that
// a named method should be exported. A struct tag should be
// included that will be used for the scripting name of the
// method and any of its arguments.
// ie.
//
//	`gd:"my_method(arg1_name, arg2_name) result_name"`
type Method struct{}

// Export the given extensions so that they are available to the engine
// (extensions should embed the gd.Extension type).
func Export(extensions ...extension) {
	for _, extension := range extensions {
		export(extension)
	}
}

// unsafeHandle is a lighterweight cgo.Handle, with the value
// identified by its index in an array.
type unsafeHandle struct {
	addr unsafe.Pointer
	used bool
}

// unsafeExtensionOffset records a gd.Object offset and the class name
// that has been placed inside a gd.Extension type. These are used to
// efficiently initialise an extension along with any objects it needs.
type unsafeOffset struct {
	class string
	bytes uintptr
	count bool // true if the object extends RefCounted.
}

/*
unsafeExtensionServer is responsible for handling requests from the Godot API for
any particular gd.Extension-derived type.
*/
type unsafeExtensionServer struct {
	index extensionClassID
	rtype reflect.Type // the value type of the extension.

	virtual func(reflect.Type, string) (method reflect.Method, ok bool) // virtual function lookup

	scripting scripting // the scripting type of the extension.

	extension unsafeOffset // gd.Extension location in the extension type.
	extending unsafeOffset // extended/embedded object location in the extension type.

	offsets []unsafeOffset
	handles []unsafeHandle // used to map index+1 to the instance.
}

// newServer creates a new server for the specified extension type.
// panics if the specified type is invalid.
func newServer(rtype reflect.Type) *unsafeExtensionServer {
	var server = unsafeExtensionServer{
		rtype: rtype,
	}

	if rtype.Kind() != reflect.Struct {
		panic("gd: extension must be a struct")
	}

	type extendable interface {
		virtual(reflect.Type, string) (method reflect.Method, ok bool)
		className() string
	}

	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)

		if field.Anonymous && field.Type == reflect.TypeOf(Extension{}) {
			server.extension = unsafeOffset{
				class: field.Tag.Get("gd"),
				bytes: field.Offset,
			}
			if server.extension.class == "" {
				server.extension.class = rtype.Name()
			}
			server.extension.class += "\000"
			continue
		}

		if field.Anonymous && reflect.PtrTo(field.Type).Implements(reflect.TypeOf((*extendable)(nil)).Elem()) {
			var zero = reflect.New(field.Type).Interface().(extendable)
			server.extending = unsafeOffset{
				class: zero.className(),
				bytes: field.Offset,
				// TODO count
			}
			server.virtual = zero.virtual
			continue
		}

		if reflect.PtrTo(field.Type).Implements(reflect.TypeOf((*extendable)(nil)).Elem()) {
			var zero = reflect.New(field.Type).Interface().(extendable)
			server.offsets = append(server.offsets, unsafeOffset{
				class: zero.className(),
				bytes: field.Offset,
				// TODO count
			})
		}

		if field.Anonymous && field.Type.Implements(reflect.TypeOf((*scripting)(nil)).Elem()) {
			server.scripting = reflect.New(field.Type).Interface().(scripting)
		}

	}

	if server.extension.class == "" {
		panic("gd: extension must embed a gd.Extension")
	}

	if server.extending.class == "" {
		panic("gd: extension must embed an extendable object")
	}

	return &server
}

// objectAt returns the *safeObject object at the given offset of an
// extension instance.
func (server *unsafeExtensionServer) objectAt(offset unsafeOffset, v reflect.Value) *safeObject {
	return (*safeObject)(unsafe.Add(v.UnsafePointer(), offset.bytes))
}

// extensionOf returns the extension field of the given value.
func (server *unsafeExtensionServer) extensionOf(val reflect.Value) *Extension {
	return (*Extension)(unsafe.Add(val.UnsafePointer(), server.extension.bytes))
}

/*
create a new instance of the extension 'class' and return the cObject
that was initialised as the instance's base godot object. The cObject
cannot be nil and MUST match the class name that the extension was
registered with.
*/
func (server *unsafeExtensionServer) create() cObject {
	instance := reflect.New(server.rtype)

	// TODO reuse existing.

	server.handles = append(server.handles, unsafeHandle{
		instance.UnsafePointer(),
		true,
	})

	ext := server.extensionOf(instance)
	owner := InstanceOwner{children: &ext.count}

	obj := server.objectAt(server.extending, instance)
	obj.new(owner, server.extending.class, server.extending.count)

	obj.cPtr.set_instance(server.extension.class, instanceID{
		class: server.index,
		index: extensionClassInstanceID(len(server.handles)),
	}.pack())

	return obj.cPtr
}

/*
delete should cleanup and free any resources associated with an
instance of the extension.
*/
func (server *unsafeExtensionServer) delete(id extensionClassInstanceID) {
	handle := server.handles[id-1]
	if !handle.used {
		return
	}

	instance := reflect.NewAt(server.rtype, handle.addr)
	server.objectAt(server.extending, instance).destroy()

	handle.used = false
}

// lookup a virtual method by name, if it is found, then the method
// definition should be returned, otherwise 'ok' should be set to false.
func (server *unsafeExtensionServer) lookup(name string) (method reflect.Method, ok bool) {
	return server.virtual(reflect.PtrTo(server.rtype), name)
}

// call the numbered method on the given instance, Godot will pass in arguments
// that match the definition that was registered. Therefore the args here is a
// pointer to an array of pointers to the arguments. The result is a pointer to
// the result and should be set before returning. If Godot expects the method
// not to return anything, result will be nil.
func (server *unsafeExtensionServer) call(id extensionClassInstanceID, methodNo uint8, args *unsafe.Pointer, result unsafe.Pointer) {
	handle := server.handles[id-1]
	if !handle.used {
		panic("gd: invalid extension instance")
	}

	rvalue := reflect.NewAt(server.rtype, handle.addr)
	method := rvalue.Method(int(methodNo) - 1)
	rtype := method.Type()

	buffer := make([]reflect.Value, rtype.NumIn())

	if rtype.NumIn() > 0 {
		slice := unsafe.Slice(args, rtype.NumIn())
		for i := range slice {
			arg := reflect.New(rtype.In(i))
			buffer[i] = arg.Elem()
			into(arg.Interface(), slice[i])
		}
	}

	results := method.Call(buffer)
	if len(results) > 1 {
		rtype := rvalue.Type()
		panic(rtype.String() + "." + rtype.Method(int(methodNo)-1).Name + " returns too many values")
	}
	if len(results) == 1 {
		if result == nil {
			rtype := rvalue.Type()
			panic(rtype.String() + "." + rtype.Method(int(methodNo)-1).Name + " returns a value but Godot expected no return value")
		}
		if !results[0].CanAddr() { // make it addressable
			old := results[0]
			results[0] = reflect.New(results[0].Type()).Elem()
			results[0].Set(old)
		}
		toResult(results[0].Addr().Interface(), result)
	}
}

func export(e extension) {
	rtype := reflect.TypeOf(e)
	if rtype.Kind() != reflect.Ptr {
		panic("gd: extension must be a pointer")
	}
	rtype = rtype.Elem()

	var server = newServer(rtype)

	id := extensionClassDB.add(server)

	server.index = id

	deferred = append(deferred, func() {
		classDB.register_extension_class(server.extension.class, server.extending.class, uintptr(id))
		if server.scripting != nil {
			server.scripting.register(server.rtype, server.extension.class)
		}
	})
}
