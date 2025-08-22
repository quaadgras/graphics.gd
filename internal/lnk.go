//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

// Link needs to be called once for the API to load in all of the
// dynamic function pointers. Typically, the link layer will take
// care of this (and you won't need to call it yourself).
func (Godot *API) Init(level GDExtensionInitializationLevel) {
	if level == GDExtensionInitializationLevelScene {
		Godot.linkBuiltin()
		Godot.linkSingletons()
		Godot.linkTypeset()
		Godot.linkTypesetCreation()
		Godot.linkMethods(false)
		Linked = true
	}
	if level == GDExtensionInitializationLevelEditor {
		Godot.linkMethods(true)
		for _, fn := range EditorStartupFunctions {
			fn()
		}
	}
}

func (Godot *API) linkSingletons() {
	rvalue := reflect.ValueOf(&Godot.Singletons).Elem()

	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		raw := pointers.Get(NewStringName(field.Name))
		rvalue.Field(i).Set(reflect.ValueOf(pointers.Raw[StringName](raw)))
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

// linkMethods, each field of cache.methods is a struct named after
// the class it represents. Each field of that struct needs to be
// filled in with a [MethodBind].
func (Godot *API) linkMethods(editor bool) {
	rvalue := reflect.ValueOf(&Godot.Methods).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		class := rvalue.Type().Field(i)
		value := reflect.NewAt(class.Type, unsafe.Add(rvalue.Addr().UnsafePointer(), class.Offset))

		isEditorMethod := false
		if strings.HasPrefix(class.Name, "Editor") {
			isEditorMethod = true
		}
		switch class.Name {
		case "FileSystemDock", "ScriptCreateDialog", "ScriptEditor", "ScriptEditorBase", "GridMapEditorPlugin":
			isEditorMethod = true
		case "JavaClassWrapper", "JavaScriptBridge", "JavaClass", "JavaObject":
			continue
		}
		if runtime.GOOS == "js" && (strings.HasPrefix(class.Name, "OpenXR") || class.Name == "ResourceImporterOggVorbis") {
			continue
		}
		if editor != isEditorMethod {
			continue
		}

		className := NewStringName(class.Name)
		for j := 0; j < class.Type.NumField(); j++ {
			method := class.Type.Field(j)
			method.Name = strings.TrimSuffix(method.Name, "_")
			method.Name = strings.TrimPrefix(method.Name, "Bind_")
			direct := reflect.NewAt(method.Type, unsafe.Add(value.UnsafePointer(), method.Offset))

			methodName := NewStringName(method.Name)

			hash, err := strconv.ParseInt(method.Tag.Get("hash"), 10, 64)
			if err != nil {
				panic("gdextension.Link: invalid gd.API builtin function hash for " + method.Name + ": " + err.Error())
			}
			bind := gdextension.Host.Objects.Method.Lookup(pointers.Get(className), pointers.Get(methodName), hash)
			if bind == 0 {
				fmt.Println("null bind ", class.Name, method.Name)
			}
			*(direct.Interface().(*gdextension.MethodForClass)) = bind

			methodName.Free()
		}
		className.Free()
	}
}

func (Godot *API) linkTypeset() {
	Godot.refCountedClassTag = Godot.ClassDB.GetClassTag(NewStringName("RefCounted"))
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
