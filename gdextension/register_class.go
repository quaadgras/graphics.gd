package gdextension

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/cgo"
	"strings"
	"unsafe"

	gd "grow.graphics/gd/internal"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

type pinner mmm.Pointer[runtime.Pinner, pinner, [0]uintptr]

func (p pinner) Free() { mmm.API(p).Unpin(); mmm.End(p) }

/*
RegisterClass registers a struct available for use inside Godot
extending the given 'Parent' Godot class. The 'Struct' type must
be a named struct that embeds a [gd.Class] field specifying the
parent class to extend.

	type MyClass struct {
		gd.Class[MyClass, gd.Node2D]
	}

Use this in a main or init function to register your Go structs
and they will become available within the Godot engine for use
in the editor and/or within scripts.

All exported fields and methods will be exposed to Godot, so
take caution when embedding types, as their fields and methods
will be promoted.

If the Struct extends [gd.EditorPlugin] then it will be added to
the editor as a plugin.
*/
func Register[Struct gd.Extends[Parent], Parent gd.IsClass](godot gd.Context) {
	var classType = reflect.TypeOf([0]Struct{}).Elem()
	var superType = reflect.TypeOf([0]Parent{}).Elem()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named struct")
	}

	var className = godot.StringName(classType.Name())
	var superName = godot.StringName(strings.TrimPrefix(superType.Name(), "class"))

	var info gd.ClassCreationInfo

	var pin runtime.Pinner
	mmm.Pin[pinner](godot.Lifetime, &pin, [0]uintptr{})

	pin.Pin(&info)

	attachProperties(godot, className, &info, classType)

	info.IsExposed = true

	info.CreateInstance.Set(func(userdata unsafe.Pointer) uintptr {
		ctx := gd.NewContext(godot.API)
		var super = godot.API.ClassDB.ConstructObject(ctx, superName)
		var instance = reflect.New(classType)
		instance.Interface().(gd.PointerToClass).SetPointer(
			mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, super.Pointer()))
		injectDependenciesInto(ctx, instance.Elem(), superType)
		godot.API.Object.SetInstance(super, className, instance.Interface())
		return super.Pointer()
	})
	info.FreeInstance.Set(func(userdata unsafe.Pointer, handle cgo.Handle) {
		//handle.Value().(gd.IsClass).Pin().End()
		handle.Delete()
	})

	// Dispatch virtual functions, these are functions that structs can
	// override with their own implementation.
	info.GetVirtual.Set(func(userdata unsafe.Pointer, ptr_name gd.StringNamePtr) gd.ExtensionClassCallVirtualFunc {
		ctx := gd.NewContext(godot.API)
		defer ctx.End()

		sname := mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, *ptr_name)
		vname := sname.String()
		var class Struct
		var virtual = gd.VirtualByName(class, vname)
		if !virtual.IsValid() {
			return gd.ExtensionClassCallVirtualFunc{}
		}
		var vtype = virtual.Type().In(0)
		GoName := convertName(vname)
		method, ok := reflect.PtrTo(classType).MethodByName(GoName)
		if !ok {
			return gd.ExtensionClassCallVirtualFunc{}
		}
		if method.Type.NumIn() != vtype.NumIn() {
			panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s", classType.Name(), GoName, virtual.Type().Name(), vname))
		}
		for i := 1; i < method.Type.NumIn(); i++ {
			if method.Type.In(i) != vtype.In(i) {
				panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s", classType.Name(), GoName, virtual.Type().Name(), vname))
			}
		}
		var copy = reflect.New(method.Type)
		copy.Elem().Set(method.Func)
		var fn = reflect.NewAt(vtype, copy.UnsafePointer()).Elem()
		return virtual.Call([]reflect.Value{fn, reflect.ValueOf(godot.API)})[0].Interface().(gd.ExtensionClassCallVirtualFunc)
	})

	var frame = call.New()
	godot.API.ClassDB.RegisterClass(godot.API.ExtensionToken,
		(gd.StringNamePtr)(unsafe.Pointer(call.Arg(frame, mmm.Get(className)).Uintptr())),
		(gd.StringNamePtr)(unsafe.Pointer(call.Arg(frame, mmm.Get(superName)).Uintptr())), &info)
	frame.Free()

	/*godot.Defer(func() {
		godot.API.ClassDB.UnregisterClass(godot.API.ExtensionToken, (gd.StringNamePtr)(unsafe.Pointer(&className)))
	})*/

	/*if superType.Implements(reflect.TypeOf([0]interface{ AsEditorPlugin() gd.EditorPlugin }{}).Elem()) {
		godot.API().EditorPlugins.Add((gd.StringNamePtr)(unsafe.Pointer(&className)))
		godot.Defer(func() {
			godot.API().EditorPlugins.Remove((gd.StringNamePtr)(unsafe.Pointer(&className)))
		})
	}*/

	registerSignals(godot, className, classType)
	registerMethods(godot, className, classType)
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
func injectDependenciesInto(godot gd.Context, value reflect.Value, superType reflect.Type) {
	if value.CanAddr() && value.Kind() != reflect.Struct {
		panic("gdextension.injectDependenciesInto: value must be an addressable struct")
	}
	localCtx := gd.NewContext(godot.API)
	defer localCtx.End()

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)

		// support private fields.
		fieldValue := reflect.NewAt(field.Type, unsafe.Add(value.Addr().UnsafePointer(), field.Offset)).Interface()

		container, ok := fieldValue.(gd.PointerToClass)
		if ok && container.Pointer() == 0 && field.Type.Implements(reflect.TypeOf([0]gd.Singleton{}).Elem()) {
			var name = localCtx.StringName(strings.TrimPrefix(field.Type.Name(), "class"))
			singleton := godot.API.Object.GetSingleton(localCtx, name)
			container.SetPointer(mmm.Let[gd.Pointer](godot.Lifetime, godot.API, singleton.Pointer()))
		}
	}
}
