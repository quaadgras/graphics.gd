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

// linkBuiltin loads in methods for the builtin Godot classes, using
// [gdunsafe.BuiltinMethod] to create typed closures for each method.
func linkBuiltin() {
	linkBuiltinType[gdunsafe.Array](&builtin.Array)
	linkBuiltinType[gdunsafe.Callable](&builtin.Callable)
	linkBuiltinType[gdunsafe.Dictionary](&builtin.Dictionary)
	linkBuiltinType[gdunsafe.PackedArray[byte]](&builtin.PackedByteArray)
	linkBuiltinType[gdunsafe.PackedArray[Color]](&builtin.PackedColorArray)
	linkBuiltinType[gdunsafe.PackedArray[float32]](&builtin.PackedFloat32Array)
	linkBuiltinType[gdunsafe.PackedArray[float64]](&builtin.PackedFloat64Array)
	linkBuiltinType[gdunsafe.PackedArray[int32]](&builtin.PackedInt32Array)
	linkBuiltinType[gdunsafe.PackedArray[gdunsafe.String]](&builtin.PackedStringArray)
	linkBuiltinType[gdunsafe.PackedArray[Vector2]](&builtin.PackedVector2Array)
	linkBuiltinType[gdunsafe.PackedArray[Vector3]](&builtin.PackedVector3Array)
	linkBuiltinType[gdunsafe.PackedArray[Vector4]](&builtin.PackedVector4Array)
	linkBuiltinType[gdunsafe.PackedArray[int64]](&builtin.PackedInt64Array)
	linkBuiltinType[gdunsafe.Signal](&builtin.Signal)
	linkBuiltinType[gdunsafe.String](&builtin.String)
	linkBuiltinType[gdunsafe.StringName](&builtin.StringName)
	linkBuiltinTypes()
}

// linkBuiltinType uses reflection to iterate over the struct fields of target,
// reads the hash from each field's struct tag, and sets the field to the closure
// returned by [gdunsafe.BuiltinMethod].
func linkBuiltinType[T gdunsafe.Any](target any) {
	rvalue := reflect.ValueOf(target)
	for method := range rvalue.Elem().Type().Fields() {
		name := strings.TrimSuffix(method.Name, "_")
		methodName := NewStringName(name)
		hash, err := strconv.ParseInt(method.Tag.Get("hash"), 10, 64)
		if err != nil {
			panic("linkBuiltinType: invalid hash for " + name + ": " + err.Error())
		}
		fn := gdunsafe.BuiltinMethod[T](gdunsafe.StringName(pointers.Get(methodName)[0]), hash)
		direct := reflect.NewAt(method.Type, unsafe.Add(rvalue.UnsafePointer(), method.Offset))
		function, ok := direct.Interface().(*func(*T, unsafe.Pointer, gdunsafe.Shape, unsafe.Pointer))
		if ok {
			*function = fn
		}
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
