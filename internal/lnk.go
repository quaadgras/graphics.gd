//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"unsafe"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
)

// Link needs to be called once for the API to load in all of the
// dynamic function pointers. Typically, the link layer will take
// care of this (and you won't need to call it yourself).
func (Godot *API) Init(level GDExtensionInitializationLevel) {
	if level == GDExtensionInitializationLevelScene {
		Godot.linkTypeset()
		Godot.linkVariant()
		Godot.linkUtility()
		Godot.linkBuiltin()
		Godot.linkSingletons()
		Godot.linkMethods(false)
		Linked = true
	}
	if level == GDExtensionInitializationLevelEditor {
		Godot.linkMethods(true)
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

// linkUtility, each field of cache.utility is a function value that
// needs to be loaded in dynamically.
func (Godot *API) linkUtility() {
	rvalue := reflect.ValueOf(&Godot.utility).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		value := reflect.NewAt(field.Type, unsafe.Add(rvalue.Addr().UnsafePointer(), field.Offset))
		name := NewStringName(field.Name)
		hash, err := strconv.ParseInt(field.Tag.Get("hash"), 10, 64)
		if err != nil {
			panic("gdextension.Link: invalid gd.API utility function hash for " + field.Name + ": " + err.Error())
		}
		*(value.Interface().(*func(ret callframe.Addr, args callframe.Args, c int32))) = Godot.Variants.GetPointerUtilityFunction(name, Int(hash))
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
			*(direct.Interface().(*func(base callframe.Addr, args callframe.Args, ret callframe.Addr, c int32))) = Godot.Variants.GetPointerBuiltinMethod(vtype, methodName, Int(hash))
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
		case "FileSystemDock", "ScriptCreateDialog", "ScriptEditor", "ScriptEditorBase":
			isEditorMethod = true
		case "JavaClassWrapper", "JavaScriptBridge":
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
			bind := Godot.ClassDB.GetMethodBind(className, methodName, Int(hash))
			if bind == 0 {
				fmt.Println("null bind ", class.Name, method.Name)
			}
			*(direct.Interface().(*MethodBind)) = bind

			methodName.Free()
		}
		className.Free()
	}
}

func (Godot *API) linkTypeset() {
	Godot.refCountedClassTag = Godot.ClassDB.GetClassTag(NewStringName("RefCounted"))

	Godot.linkTypesetCreation()
	Godot.linkTypesetOperator()
	Godot.linkTypesetDestruct()
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
			*(value.Interface().(*func(callframe.Addr, callframe.Args))) = Godot.Variants.GetPointerConstructor(vtype, int32(i))
		}
	}
}

// linkTypesetOperator, each field is a struct named after a builtin type, with a list of
// operator fields, named after the operator, each field is either a function (no right
// hand side) or a struct with a function field for each right hand side type.
func (Godot *API) linkTypesetOperator() {
	rvalue := reflect.ValueOf(&Godot.typeset.operator).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		class := rvalue.Type().Field(i)
		value := reflect.NewAt(class.Type, unsafe.Add(rvalue.Addr().UnsafePointer(), class.Offset))
		vtype, _ := variantTypeFromName(class.Name)
		for j := 0; j < class.Type.NumField(); j++ {
			op := class.Type.Field(j)
			otype := operatoTypeFromName(op.Name)
			switch op.Type.Kind() {
			case reflect.Func:
				direct := value.Elem().Field(j).Addr().Interface().(*func(a, b, ret callframe.Addr))
				*direct = Godot.Variants.PointerOperatorEvaluator(otype, vtype, TypeNil)
			default:
				for k := 0; k < op.Type.NumField(); k++ {
					rhs := op.Type.Field(k)
					rhsType, _ := variantTypeFromName(rhs.Name)
					unsafe := reflect.NewAt(rhs.Type,
						unsafe.Add(value.Elem().Field(j).Addr().UnsafePointer(), rhs.Offset))
					direct := unsafe.Interface().(*func(a, b, ret callframe.Addr))
					*direct = Godot.Variants.PointerOperatorEvaluator(otype, vtype, rhsType)
				}
			}
		}
	}
}

// linkTypesetDestruct, each field is a function value that needs to be loaded in dynamically.
func (Godot *API) linkTypesetDestruct() {
	rvalue := reflect.ValueOf(&Godot.typeset.destruct).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		value := reflect.NewAt(field.Type, unsafe.Add(rvalue.Addr().UnsafePointer(), field.Offset))
		vtype, _ := variantTypeFromName(field.Name)
		*(value.Interface().(*func(callframe.Addr))) = Godot.Variants.GetPointerDestructor(vtype)
	}
}

func (Godot *API) linkVariant() {
	for i := VariantType(1); i < TypeMax; i++ {
		Godot.variant.FromType[i] = Godot.Variants.FromTypeConstructor(i)
		Godot.variant.IntoType[i] = Godot.Variants.ToTypeConstructor(i)
	}
}
