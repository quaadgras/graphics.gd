package gdextension

import (
	"reflect"
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
func RegisterStruct[Struct, Parent any](extension API, db ClassDB) {
	var classType = reflect.TypeOf([0]Struct{}).Elem()
	var superType = reflect.TypeOf([0]Parent{}).Elem()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named truct")
	}

	var className gd.StringName
	var superName gd.StringName
	extension.StringNames.New(&className, classType.Name())
	extension.StringNames.New(&superName, superType.Name())

	var info ClassCreationInfo
	info.IsExposed = true

	info.CreateInstance.Set(func(userdata unsafe.Pointer) Object {
		var super = extension.ClassDB.CreateObject(&superName)
		extension.Object.SetInstance(super, &className, nil)
		return super
	})
	info.FreeInstance.Set(func(userdata unsafe.Pointer, obj Object) {
		return
	})
	info.GetVirtual.Set(func(userdata unsafe.Pointer, name *gd.StringName) call.Back[func(instance unsafe.Pointer, args unsafe.Pointer, ret unsafe.Pointer)] {
		return call.Back[func(instance unsafe.Pointer, args unsafe.Pointer, ret unsafe.Pointer)]{}
	})

	extension.ClassDB.RegisterClass(db, &className, &superName, &info)
}
