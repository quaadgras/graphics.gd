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
		linkTypeset()
		linkTypesetCreation()
		LinkMethods(gdunsafe.StringName(pointers.Get(NewStringName("Object"))[0]), &object_methods, false)
		if LinkStartup != nil {
			LinkStartup()
		}
		LinkedCore = true
	}
	if !Linked && level == gdextension.InitializationLevelScene {
		LinkMethods(gdunsafe.StringName(pointers.Get(NewStringName("RefCounted"))[0]), &refcounted_methods, false)
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

func LinkMethods(className gdunsafe.StringName, methods any, editor bool) {
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
		bind := gdunsafe.Method(className, gdunsafe.StringName(pointers.Get(methodName)[0]), hash)
		if bind == 0 {
			fmt.Println("null bind ", method.Name)
		}
		*(direct.Interface().(*gdextension.MethodForClass)) = bind

		methodName.Free()
	}
}

var refCountedClassTag gdunsafe.ClassTag

func linkTypeset() {
	refCountedClassTag = gdunsafe.Class(pointers.Get(NewStringName("RefCounted"))[0]).Tag()
}
