//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

var Linked bool = false
var LinkedCore bool = false
var LinkedEditor bool = false

var Links []func()
var LinkStartup func()

// Link needs to be called once for the API to load in all of the
// dynamic function pointers. Typically, the link layer will take
// care of this (and you won't need to call it yourself).
func Init(level gdextension.InitializationLevel) {
	if !LinkedCore && level == gdextension.InitializationLevelCore {
		linkBuiltin()
		linkTypeset()
		linkTypesetCreation()
		LinkMethods(pointers.Get(NewStringName("Object")), &object_methods, false)
		if LinkStartup != nil {
			LinkStartup()
		}
		LinkedCore = true
	}
	if !Linked && level == gdextension.InitializationLevelScene {
		LinkMethods(pointers.Get(NewStringName("RefCounted")), &refcounted_methods, false)
		for _, fn := range Links {
			fn()
		}
		Linked = true
	}
	if !LinkedEditor && level == gdextension.InitializationLevelEditor {
		for _, fn := range EditorStartupFunctions {
			fn()
		}
		LinkedEditor = true
	}
}

// linkBuiltin is very similar to [Godot.linkMethods], except it loads in methods for the
// builtin Godot classes.
func linkBuiltin() {
	rvalue := reflect.ValueOf(&builtin).Elem()
	for i := 1; i < rvalue.NumField(); i++ {
		class := rvalue.Type().Field(i)
		value := reflect.NewAt(class.Type, unsafe.Add(rvalue.Addr().UnsafePointer(), class.Offset))
		for method := range class.Type.Fields() {
			method.Name = strings.TrimSuffix(method.Name, "_")
			direct := reflect.NewAt(method.Type, unsafe.Add(value.UnsafePointer(), method.Offset))
			methodName := NewStringName(method.Name)
			hash, err := strconv.ParseInt(method.Tag.Get("hash"), 10, 64)
			if err != nil {
				panic("gdextension.Link: invalid gd.API builtin function hash for " + method.Name + ": " + err.Error())
			}
			vtype, _ := variantTypeFromName(class.Name)
			*(direct.Interface().(*gdextension.MethodForBuiltinType)) = gdextension.MethodForBuiltinType(gdunsafe.VariantTypeMethod(gdunsafe.VariantType(vtype), gdunsafe.StringName(pointers.Get(methodName)[0]), int64(hash)))
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
	for method := range rvalue.Elem().Fields() {
		direct := reflect.NewAt(method.Type, unsafe.Add(rvalue.UnsafePointer(), method.Offset))

		method.Name = strings.TrimSuffix(method.Name, "_")

		methodName := NewStringName(method.Name)

		hash, err := strconv.ParseInt(method.Tag.Get("hash"), 10, 64)
		if err != nil {
			panic("gdextension.Link: invalid gd.API builtin function hash for " + method.Name + ": " + err.Error())
		}
		bind := gdextension.MethodForClass(gdunsafe.MethodLookup(gdunsafe.StringName(className[0]), gdunsafe.StringName(pointers.Get(methodName)[0]), hash))
		if bind == 0 {
			fmt.Println("null bind ", method.Name)
		}
		*(direct.Interface().(*gdextension.MethodForClass)) = bind

		methodName.Free()
	}
}

var refCountedClassTag gdunsafe.ObjectType

func linkTypeset() {
	refCountedClassTag = gdunsafe.ObjectTypeTag(gdunsafe.StringName(pointers.Get(NewStringName("RefCounted"))[0]))
}

// linkTypesetCreation, each field is an array of constructors.
func linkTypesetCreation() {
	rvalue := reflect.ValueOf(&builtin.creation).Elem()
	for field := range rvalue.Type().Fields() {
		esize := field.Type.Elem().Size()
		vtype, _ := variantTypeFromName(field.Name)
		for i := 0; i < field.Type.Len(); i++ {
			value := reflect.NewAt(field.Type.Elem(), unsafe.Add(rvalue.Addr().UnsafePointer(), field.Offset+uintptr(i)*esize))
			*(value.Interface().(*gdextension.FunctionID)) = gdextension.FunctionID(gdunsafe.VariantTypeConstructor(gdunsafe.VariantType(vtype), int64(i)))
		}
	}
}
