//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

var Links []func()

// Link needs to be called once for the API to load in all of the
// dynamic function pointers. Typically, the link layer will take
// care of this (and you won't need to call it yourself).
func (Godot *API) Init(level gdextension.InitializationLevel) {
	if level == gdextension.InitializationLevelScene {
		Godot.linkBuiltin()
		Godot.linkTypeset()
		Godot.linkTypesetCreation()
		LinkMethods(pointers.Get(NewStringName("Object")), &object_methods, false)
		LinkMethods(pointers.Get(NewStringName("RefCounted")), &refcounted_methods, false)
		for _, fn := range Links {
			fn()
		}
		Linked = true
	}
	if level == gdextension.InitializationLevelEditor {
		for _, fn := range EditorStartupFunctions {
			fn()
		}
	}
}

// linkBuiltin is very similar to [Godot.linkMethods], except it loads in methods for the
// builtin Godot classes.
func (Godot *API) linkBuiltin() {
	rvalue := reflect.ValueOf(&Godot.builtin).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		class := rvalue.Type().Field(i)
		value := reflect.NewAt(class.Type, unsafe.Add(rvalue.Addr().UnsafePointer(), class.Offset))
		for j := 0; j < class.Type.NumField(); j++ {
			method := class.Type.Field(j)
			method.Name = strings.TrimSuffix(method.Name, "_")
			direct := reflect.NewAt(method.Type, unsafe.Add(value.UnsafePointer(), method.Offset))
			methodName := NewStringName(method.Name)
			hash, err := strconv.ParseInt(method.Tag.Get("hash"), 10, 64)
			if err != nil {
				panic("gdextension.Link: invalid gd.API builtin function hash for " + method.Name + ": " + err.Error())
			}
			vtype, _ := variantTypeFromName(class.Name)
			*(direct.Interface().(*gdextension.MethodForBuiltinType)) = gdextension.Host.Builtin.Types.Method(vtype, pointers.Get(methodName), Int(hash))
		}
	}
}

func LinkMethods(className gdextension.StringName, methods any, editor bool) {
	if editor {
		EditorStartupFunctions = append(EditorStartupFunctions, func() {
			LinkMethods(className, methods, false)
		})
		return
	}
	rvalue := reflect.ValueOf(methods)
	for i := range rvalue.Elem().NumField() {
		method := rvalue.Elem().Type().Field(i)
		direct := reflect.NewAt(method.Type, unsafe.Add(rvalue.UnsafePointer(), method.Offset))

		method.Name = strings.TrimSuffix(method.Name, "_")

		methodName := NewStringName(method.Name)

		hash, err := strconv.ParseInt(method.Tag.Get("hash"), 10, 64)
		if err != nil {
			panic("gdextension.Link: invalid gd.API builtin function hash for " + method.Name + ": " + err.Error())
		}
		bind := gdextension.Host.Objects.Method.Lookup(className, pointers.Get(methodName), hash)
		if bind == 0 {
			fmt.Println("null bind ", pointers.Raw[StringName](className), method.Name)
		}
		*(direct.Interface().(*gdextension.MethodForClass)) = bind

		methodName.Free()
	}
}

func (Godot *API) linkTypeset() {
	Godot.refCountedClassTag = gdextension.Host.Objects.Type(pointers.Get(NewStringName("RefCounted")))
}

// linkTypesetCreation, each field is an array of constructors.
func (Godot *API) linkTypesetCreation() {
	rvalue := reflect.ValueOf(&Godot.typeset.creation).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		esize := field.Type.Elem().Size()
		vtype, _ := variantTypeFromName(field.Name)
		for i := 0; i < field.Type.Len(); i++ {
			value := reflect.NewAt(field.Type.Elem(), unsafe.Add(rvalue.Addr().UnsafePointer(), field.Offset+uintptr(i)*esize))
			*(value.Interface().(*gdextension.FunctionID)) = gdextension.Host.Builtin.Types.Constructor(vtype, i)
		}
	}
}
