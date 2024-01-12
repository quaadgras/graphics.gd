package gdextension

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"grow.graphics/gd"
	"runtime.link/api"
	"runtime.link/api/call"
	"runtime.link/api/stub"
)
import "C"

var classDB gd.ExtensionToken
var dlsymGD func(string) unsafe.Pointer

// Link returns a handle to the [API] and the global [ClassDB].
// The [bool] return value is [true] if the API has been
// linked with Godot successfully.
func Link() (gd.ExtensionAPI, gd.ExtensionToken, bool) {
	if dlsymGD == nil {
		return api.Import[gd.ExtensionAPI](stub.API, "", errors.New("gdextension not linked")), 0, false
	}
	extension := api.Import[gd.ExtensionAPI](call.API, "", call.Options{
		LookupSymbol: func(name string) (unsafe.Pointer, error) {
			sym := dlsymGD(name)
			if sym == nil {
				return nil, errors.New("'" + name + "' symbol not found")
			}
			return sym, nil
		},
	})
	return extension, classDB, true
}

func variantTypeFromName(extension gd.ExtensionAPI, s string) (gd.VariantType, reflect.Type) {
	switch s {
	case "Bool":
		return gd.TypeBool, reflect.TypeOf(false)
	case "Int":
		return gd.TypeInt, reflect.TypeOf(int64(0))
	case "Float":
		return gd.TypeFloat, reflect.TypeOf(gd.Float(0))
	case "String":
		return gd.TypeString, reflect.TypeOf(gd.String{})
	case "Vector2":
		return gd.TypeVector2, reflect.TypeOf(gd.Vector2{})
	case "Vector2i":
		return gd.TypeVector2i, reflect.TypeOf(gd.Vector2i{})
	case "Rect2":
		return gd.TypeRect2, reflect.TypeOf(gd.Rect2{})
	case "Rect2i":
		return gd.TypeRect2i, reflect.TypeOf(gd.Rect2i{})
	case "Vector3":
		return gd.TypeVector3, reflect.TypeOf(gd.Vector3{})
	case "Vector3i":
		return gd.TypeVector3i, reflect.TypeOf(gd.Vector3i{})
	case "Transform2D":
		return gd.TypeTransform2d, reflect.TypeOf(gd.Transform2D{})
	case "Vector4":
		return gd.TypeVector4, reflect.TypeOf(gd.Vector4{})
	case "Vector4i":
		return gd.TypeVector4i, reflect.TypeOf(gd.Vector4i{})
	case "Plane":
		return gd.TypePlane, reflect.TypeOf(gd.Plane{})
	case "Quaternion":
		return gd.TypeQuaternion, reflect.TypeOf(gd.Quaternion{})
	case "AABB":
		return gd.TypeAabb, reflect.TypeOf(gd.AABB{})
	case "Basis":
		return gd.TypeBasis, reflect.TypeOf(gd.Basis{})
	case "Transform3D":
		return gd.TypeTransform3d, reflect.TypeOf(gd.Transform3D{})
	case "Projection":
		return gd.TypeProjection, reflect.TypeOf(gd.Projection{})
	case "Color":
		return gd.TypeColor, reflect.TypeOf(gd.Color{})
	case "StringName":
		return gd.TypeStringName, reflect.TypeOf(gd.StringName{})
	case "NodePath":
		return gd.TypeNodePath, reflect.TypeOf(gd.NodePath{})
	case "RID":
		return gd.TypeRid, reflect.TypeOf(gd.RID(0))
	case "Object":
		return gd.TypeObject, reflect.TypeOf(uintptr(0))
	case "Callable":
		return gd.TypeCallable, reflect.TypeOf(gd.Callable{})
	case "Signal":
		return gd.TypeSignal, reflect.TypeOf(gd.Signal{})
	case "Dictionary":
		return gd.TypeDictionary, reflect.TypeOf(gd.Dictionary{})
	case "Array":
		return gd.TypeArray, reflect.TypeOf(gd.Array{})
	case "PackedByteArray":
		return gd.TypePackedByteArray, reflect.TypeOf(gd.PackedByteArray{})
	case "PackedInt32Array":
		return gd.TypePackedInt32Array, reflect.TypeOf(gd.PackedInt32Array{})
	case "PackedInt64Array":
		return gd.TypePackedInt64Array, reflect.TypeOf(gd.PackedInt64Array{})
	case "PackedFloat32Array":
		return gd.TypePackedFloat32Array, reflect.TypeOf(gd.PackedFloat32Array{})
	case "PackedFloat64Array":
		return gd.TypePackedFloat64Array, reflect.TypeOf(gd.PackedFloat64Array{})
	case "PackedStringArray":
		return gd.TypePackedStringArray, reflect.TypeOf(gd.PackedStringArray{})
	case "PackedVector2Array":
		return gd.TypePackedVector2Array, reflect.TypeOf(gd.PackedVector2Array{})
	case "PackedVector3Array":
		return gd.TypePackedVector3Array, reflect.TypeOf(gd.PackedVector3Array{})
	case "PackedColorArray":
		return gd.TypePackedColorArray, reflect.TypeOf(gd.PackedColorArray{})
	default:
		panic("gdextension.variantTypeFromName: unknown type " + s)
	}
}

func link(extension gd.ExtensionAPI, godot *gd.API, structure api.Structure) {
	godot.Extension = extension
	var className gd.StringName
	var methodName gd.StringName
	for _, fn := range structure.Functions {
		fn := fn
		class, method, ok := strings.Cut(fn.Name, "_")
		if !ok {
			continue
		}
		if strings.HasPrefix(method, "_") {
			continue
		}
		if class == "Utility" || class == "JavaClassWrapper" || class == "JavaScriptBridge" {
			continue
		}
		// Builtin types need their methods loaded via a different method.
		switch class {
		case "Bool", "Int", "Float",
			"String", "StringName", "NodePath", "RID", "Color", "Basis",
			"Callable", "Signal", "Dictionary", "Array", "Projection",
			"Transform3D", "Transform2D", "Plane", "Quaternion", "AABB",
			"Rect2", "Rect2i",
			"Vector2", "Vector3", "Vector4", "Vector2i", "Vector3i", "Vector4i",
			"PackedByteArray", "PackedInt32Array", "PackedInt64Array",
			"PackedFloat32Array", "PackedFloat64Array", "PackedStringArray",
			"PackedVector2Array", "PackedVector3Array", "PackedColorArray":
			indexS, ok := fn.Tags.Lookup("index")
			if ok {
				index, err := strconv.ParseInt(indexS, 10, 32)
				if err != nil {
					panic("gdextension.Link: invalid gd.API constructor index for " + fn.Name + ": " + err.Error())
				}
				vtype, gtype := variantTypeFromName(extension, class)
				constructor := extension.Variants.Constructor(vtype, int32(index))
				fn.Make(func(args []reflect.Value) []reflect.Value {
					var ptrs = make([]unsafe.Pointer, len(args))
					for i := 0; i < len(args); i++ {
						var copy = reflect.New(args[i].Type())
						copy.Elem().Set(args[i])
						ptrs[i] = copy.UnsafePointer()
					}
					var ret = reflect.New(gtype)
					constructor(ret.UnsafePointer(), &ptrs[0])
					return []reflect.Value{ret.Elem()}
				})
			} else {
				hash, err := strconv.ParseInt(fn.Tags.Get("hash"), 10, 64)
				if err != nil {
					panic("gdextension.Link: invalid gd.API builtin function hash for " + fn.Name + ": " + err.Error())
				}
				extension.StringNames.New(&methodName, method)
				vtype, _ := variantTypeFromName(extension, class)
				caller := extension.Variants.BuiltinMethodCaller(vtype, &methodName, hash)
				fn.Make(func(args []reflect.Value) []reflect.Value {
					var ptrs = make([]unsafe.Pointer, len(args))
					for i := 0; i < len(args[1:]); i++ {
						var copy = reflect.New(args[i].Type())
						copy.Elem().Set(args[i])
						ptrs[i] = copy.UnsafePointer()
					}
					if fn.NumOut() > 0 {
						ret := reflect.New(fn.Type.Out(0))
						caller(args[0].UnsafePointer(), ptrs, ret.UnsafePointer())
						return []reflect.Value{ret.Elem()}
					} else {
						caller(args[0].UnsafePointer(), ptrs, nil)
						return []reflect.Value{}
					}
				})
			}
			continue
		}
		hash, err := strconv.ParseInt(fn.Tags.Get("hash"), 10, 64)
		if err != nil {
			panic("gdextension.Link: invalid gd.API function hash for " + fn.Name + ": " + err.Error())
		}
		extension.StringNames.New(&className, class)
		extension.StringNames.New(&methodName, method)
		var bind = extension.ClassDB.GetMethodBind(&className, &methodName, hash)
		if bind != 0 {
			fn.Make(func(args []reflect.Value) []reflect.Value {
				self := gd.ReadPointer(args[0].Interface().(gd.ClassReference))
				var ptrs = make([]unsafe.Pointer, len(args)-1)
				for i := 1; i < len(args); i++ {
					var copy = reflect.New(args[i].Type())
					copy.Elem().Set(args[i])
					ptrs[i-1] = copy.UnsafePointer()
				}
				if fn.NumOut() == 0 {
					extension.Object.UnsafeCall(bind, self, ptrs, nil)
					return []reflect.Value{}
				}
				var ret = reflect.New(fn.Type.Out(0))
				container, ok := ret.Interface().(gd.ClassContainer)
				if ok {
					var ptr uintptr
					extension.Object.UnsafeCall(bind, self, ptrs, unsafe.Pointer(&ptr))
					gd.MakePointer(container, gd.Pointer{Val: ptr, API: godot})
				} else {
					extension.Object.UnsafeCall(bind, self, ptrs, ret.UnsafePointer())
				}
				return []reflect.Value{ret.Elem()}
			})
		}
		extension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&className))
		extension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&methodName))
	}
}

var (
	godot = new(gd.API)
)

//export loadExtension
func loadExtension(lookupFunc, classes, configuration unsafe.Pointer) uint8 {
	dlsym, err := call.Make[func(string) unsafe.Pointer](lookupFunc, "func(&char)~void")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	dlsymGD = dlsym
	classDB = gd.ExtensionToken(classes)

	GDextension := api.Import[gd.ExtensionAPI](call.API, "", call.Options{
		LookupSymbol: func(name string) (unsafe.Pointer, error) {
			sym := dlsym(name)
			if sym == nil {
				return nil, errors.New("'" + name + "' symbol not found")
			}
			return sym, nil
		},
	})

	var name gd.StringName
	GDextension.StringNames.New(&name, "AudioFrame")

	GDextension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&name))

	init := (*gd.ExtensionInitialization[uintptr])(configuration)
	*init = gd.ExtensionInitialization[uintptr]{}
	init.MinimumInitializationLevel = gd.ExtensionInitializationLevelScene
	init.Initialize.Set(func(userdata uintptr, level gd.ExtensionInitializationLevel) {
		if level == 3 {
			link(GDextension, godot, api.StructureOf(godot))
		}
		if level == 2 {
			main()
		}
		/*if level == 3 {

			// playground
			GDextension.StringNames.New(&name, "Engine")
			var Engine = GDextension.Object.GetSingleton(&name)

			var method gd.StringName
			GDextension.StringNames.New(&method, "get_license_text")
			var bind = GDextension.ClassDB.GetMethodBind(&name, &method, 201670096)

			GDextension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&name))
			GDextension.Variants.Destructor(gd.TypeStringName)(unsafe.Pointer(&method))

			fmt.Println("bind", bind)

			fmt.Println(Engine)

			var result gd.String
			GDextension.Object.UnsafeCall(bind, Engine, nil, unsafe.Pointer(&result))
			fmt.Println("result=", result)

			var buf = make([]byte, 1024)
			GDextension.Strings.Get(&result, buf)
			fmt.Println(string(buf))
		}*/
	})
	init.Deinitialize.Set(func(userdata uintptr, level gd.ExtensionInitializationLevel) {
		if level == 3 {
			init.Initialize.Free()
		}
		if level == 0 {
			init.Deinitialize.Free()
		}
	})
	return 1
}

//go:linkname main main.main
func main()
