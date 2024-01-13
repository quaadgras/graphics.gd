package gdextension

import (
	"fmt"
	"reflect"
	"runtime/cgo"
	"strings"
	"unsafe"

	"grow.graphics/gd"
	"runtime.link/api/call"
)

/*
RegisterStruct registers a struct available for use inside Godot
extending the given 'Parent' Godot class. The 'Struct' type must
be a named struct that implements a method with the same name as
the 'Parent' Godot class. This method must return the 'Parent'
class from a field within the struct.

	type MyClass struct {
		super gd.Node2D
	}

	func (m *MyClass) Node2D() gd.Node2D { return m.super }

Use this in a main or init function to register your Go structs
and they will become available within the Godot engine for use
in the editor and/or within scripts.

All exported fields and methods will be exposed to Godot, so
take caution when embedding types, as their fields and methods
will be promoted.
*/
func RegisterStruct[Struct gd.Extends[Parent], Parent gd.IsClass](extension gd.ExtensionAPI, db gd.ExtensionToken) {
	var classType = reflect.TypeOf([0]Struct{}).Elem()
	var superType = reflect.TypeOf([0]Parent{}).Elem()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named struct")
	}

	var godotFunctions = reflect.TypeOf(gd.API{})

	var className gd.StringName
	var superName gd.StringName
	extension.StringNames.New(&className, classType.Name())
	extension.StringNames.New(&superName, strings.TrimPrefix(superType.Name(), "class"))

	var info gd.ClassCreationInfo
	info.IsExposed = true

	info.CreateInstance.Set(func(userdata unsafe.Pointer) uintptr {
		var super = extension.ClassDB.CreateObject(&superName)
		var instance = reflect.New(classType)
		gd.MakePointer(instance.Interface().(gd.ClassContainer), gd.Pointer{
			Val: [2]uintptr{super, 0},
			API: godot,
		})
		injectDependenciesInto(extension, instance.Elem(), super, superType)
		var handle = cgo.NewHandle(instance.Interface())
		extension.Object.SetInstance(super, &className, handle)
		return super
	})
	info.FreeInstance.Set(func(userdata unsafe.Pointer, handle cgo.Handle) {
		handle.Delete()
	})

	// Dispatch virtual functions, these are functions that structs can
	// override with their own implementation.
	info.GetVirtual.Set(func(userdata unsafe.Pointer, sname *gd.StringName) call.Back[func(instance cgo.Handle, args unsafe.Pointer, ret unsafe.Pointer)] {
		var out = godot.String_NewFromStringName(*sname)
		var buf = make([]byte, godot.String_length(&out))
		extension.Strings.Get(&out, buf)
		vname := string(buf)

		var virtualType reflect.Type = reflect.PtrTo(superType)
		for {
			superMethod, ok := virtualType.MethodByName("Super")
			if !ok {
				break
			}
			virtualType = superMethod.Type.Out(0)
			virtualMethod, ok := godotFunctions.FieldByName(strings.TrimPrefix(virtualType.Elem().Name(), "class") + "_" + vname)
			if !ok {
				continue
			}
			GoName := convertName(vname)
			method, ok := reflect.PtrTo(classType).MethodByName(GoName)
			if !ok {
				continue
			}
			for i := 1; i < method.Type.NumIn(); i++ {
				if method.Type.In(i) != virtualMethod.Type.In(i) {
					panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s", classType.Name(), GoName, virtualType.Name(), vname))
				}
			}
			fmt.Println("Registering virtual function", classType.Name(), GoName, virtualType.Name(), vname)
			var callMethod call.Back[func(instance cgo.Handle, args unsafe.Pointer, ret unsafe.Pointer)]
			callMethod.Set(func(instance cgo.Handle, args unsafe.Pointer, ret unsafe.Pointer) {
				var in = make([]reflect.Value, method.Type.NumIn())
				in[0] = reflect.ValueOf(instance.Value())
				for i := 1; i < method.Type.NumIn(); i++ {
					in[i] = reflect.NewAt(method.Type.In(i), args).Elem()
				}
				var out = method.Func.Call(in)
				if len(out) > 0 {
					reflect.NewAt(method.Type.Out(0), ret).Elem().Set(out[0])
				}
			})
			// FIXME callMethod needs to be freed at some point?
			return callMethod
		}

		return call.Back[func(instance cgo.Handle, args unsafe.Pointer, ret unsafe.Pointer)]{}
	})

	extension.ClassDB.RegisterClass(db, &className, &superName, &info)
}

func convertName(fnName string) string {
	if fnName == "seek" {
		return "SeekTo"
	}
	if fnName == "type_string" {
		return "TypeToString"
	}
	fnName = strings.ToLower(fnName)
	joins := []string{}
	split := strings.Split(fnName, "_")
	for _, word := range split {
		joins = append(joins, strings.Title(word))
	}
	return strings.Join(joins, "")
}

// injectDependenciesInto the given freshly allocated value, this
// function makes sure any [gd.Object] types are instantiated and
// that any referenced singletons are injected.
func injectDependenciesInto(extension gd.ExtensionAPI, value reflect.Value, super uintptr, superType reflect.Type) {
	if value.CanAddr() && value.Kind() != reflect.Struct {
		panic("gdextension.injectDependenciesInto: value must be an addressable struct")
	}
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)

		// support private fields.
		fieldValue := reflect.NewAt(field.Type, unsafe.Add(value.Addr().UnsafePointer(), field.Offset)).Interface()

		container, ok := fieldValue.(gd.ClassContainer)
		if ok && (gd.ReadPointer(container) == [2]uintptr{}) {
			_, ok := fieldValue.(gd.Singleton)
			if ok {
				fmt.Println("making singleton")
				var name gd.StringName
				extension.StringNames.New(&name, strings.TrimPrefix(field.Type.Name(), "class"))
				singleton := extension.Object.GetSingleton(&name)
				extension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&name))
				gd.MakePointer(container, gd.Pointer{
					Val: [2]uintptr{singleton, 0},
					API: godot,
				})
			}
		}

	}
}
