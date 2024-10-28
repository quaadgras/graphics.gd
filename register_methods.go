//go:build !generate

package gd

import (
	"fmt"
	"reflect"

	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/internal/discreet"
)

func registerMethods(class StringName, rtype reflect.Type) {
	classTypePtr := reflect.PointerTo(rtype)
	for i := 0; i < classTypePtr.NumMethod(); i++ {
		i := i

		var hasContext bool = false

		method := classTypePtr.Method(i)
		if !method.IsExported() || method.Type.NumIn() < 1 {
			continue
		}
		parent, ok := rtype.FieldByName("Class")
		if !ok {
			panic(fmt.Sprintf("gdextension: %v does not have an embedded Class field", rtype))
		}
		if _, ok := reflect.PointerTo(parent.Type).MethodByName(method.Name); ok {
			continue
		}
		if method.Name == "OnRegister" {
			continue
		}
		var offset = 0
		var arguments = make([]gd.PropertyInfo, 0, method.Type.NumIn()-1-offset)
		var metadatas = make([]gd.ClassMethodArgumentMetadata, 0, method.Type.NumIn()-1-offset)
		for i := 1 + offset; i < method.Type.NumIn(); i++ {
			arguments = append(arguments,
				propertyOf(reflect.StructField{Name: "arg" + fmt.Sprint(i), Type: method.Type.In(i)}))
			metadatas = append(metadatas, 0)
		}
		var returns *gd.PropertyInfo
		var returnMetadata gd.ClassMethodArgumentMetadata
		if method.Type.NumOut() > 0 {
			property := propertyOf(reflect.StructField{Name: "result", Type: method.Type.Out(0)})
			returns = &property
			returnMetadata = 0
		}

		gd.Global.ClassDB.RegisterClassMethod(gd.Global.ExtensionToken, class, gd.Method{
			Name: gd.NewStringName(method.Name),
			// FIXME type check and return an error if arguments are invalid.
			Call: func(instance any, v ...gd.Variant) (gd.Variant, error) {
				var args = make([]reflect.Value, len(v)+offset)
				for i := 0; i < len(v); i++ {
					if method.Type.In(i + 1 + offset).Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()) {
						var obj = reflect.New(method.Type.In(i + 1 + offset))
						obj.Interface().(gd.PointerToClass).SetPointer(gd.LetVariantAsPointerType[gd.Object](v[i], TypeObject))
						args[i+offset] = obj.Elem()
					} else {
						args[i+offset] = reflect.ValueOf(v[i].Interface())
					}
				}
				extensionInstance := instance.(*instanceImplementation).Value
				rets := reflect.ValueOf(extensionInstance).Method(i).Call(args)
				if len(rets) > 0 {
					return gd.NewVariant(rets[0].Interface()), nil
				}
				return gd.Variant{}, nil
			},
			PointerCall: func(instance any, args gd.UnsafeArgs, ret gd.UnsafeBack) {
				extensionInstance := instance.(*instanceImplementation).Value
				slowCall(hasContext, reflect.ValueOf(extensionInstance).Method(i), args, ret)
			},
			Arguments:           arguments,
			ArgumentsMetadata:   metadatas,
			ReturnValueInfo:     returns,
			ReturnValueMetadata: returnMetadata,
		})
	}
}

func slowCall(hasContext bool, method reflect.Value, p_args gd.UnsafeArgs, p_ret gd.UnsafeBack) {
	godot := gd.Global
	var (
		args = make([]reflect.Value, method.Type().NumIn())
	)
	var offset = 0
	for i := offset; i < method.Type().NumIn(); i++ {
		switch method.Type().In(i) {
		case reflect.TypeOf(Bool(false)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Bool](p_args, i-offset))
		case reflect.TypeOf(Int(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Int](p_args, i-offset))
		case reflect.TypeOf(Float(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Float](p_args, i-offset))
		case reflect.TypeOf(String{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := discreet.New[String](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(Vector2{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector2](p_args, i-offset))
		case reflect.TypeOf(Vector2i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector2i](p_args, i-offset))
		case reflect.TypeOf(Rect2{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Rect2](p_args, i-offset))
		case reflect.TypeOf(Rect2i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Rect2i](p_args, i-offset))
		case reflect.TypeOf(Vector3{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector3](p_args, i-offset))
		case reflect.TypeOf(Vector3i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector3i](p_args, i-offset))
		case reflect.TypeOf(Transform2D{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Transform2D](p_args, i-offset))
		case reflect.TypeOf(Vector4{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector4](p_args, i-offset))
		case reflect.TypeOf(Vector4i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector4i](p_args, i-offset))
		case reflect.TypeOf(Plane{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Plane](p_args, i-offset))
		case reflect.TypeOf(Quaternion{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Quaternion](p_args, i-offset))
		case reflect.TypeOf(AABB{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[AABB](p_args, i-offset))
		case reflect.TypeOf(Basis{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Basis](p_args, i-offset))
		case reflect.TypeOf(Transform3D{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Transform3D](p_args, i-offset))
		case reflect.TypeOf(Projection{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Projection](p_args, i-offset))
		case reflect.TypeOf(Color{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Color](p_args, i-offset))
		case reflect.TypeOf(StringName{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := discreet.New[StringName](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(NodePath{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := discreet.New[NodePath](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(RID(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[RID](p_args, i-offset))
		case reflect.TypeOf(Object{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := discreet.New[Object]([3]uintptr{ptr[0]})
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(Callable{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[Callable](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.Signal{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[Signal](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(Dictionary{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := discreet.New[Dictionary](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(Array{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := discreet.New[Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedByteArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedByteArray](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedInt32Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedInt32Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedInt64Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedInt64Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedFloat32Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedFloat32Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedFloat64Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedFloat64Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedStringArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedStringArray](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedVector2Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedVector2Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedVector3Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedVector3Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(PackedColorArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := discreet.New[PackedColorArray](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		default:
			panic(fmt.Sprintf("gdextension: unsupported Godot -> Go type %v", method.Type().In(i)))
		}
	}
	rrets := method.Call(args)
	if len(rrets) > 0 {
		switch val := rrets[0].Interface().(type) {
		case Bool:
			gd.UnsafeSet[Bool](p_ret, val)
		case Int:
			gd.UnsafeSet[Int](p_ret, val)
		case Float:
			gd.UnsafeSet[Float](p_ret, val)
		case String:
			gd.UnsafeSet[[1]uintptr](p_ret, discreet.Get(val))
			discreet.End(val)
		case Vector2:
			gd.UnsafeSet[Vector2](p_ret, val)
		case Vector2i:
			gd.UnsafeSet[Vector2i](p_ret, val)
		case Rect2:
			gd.UnsafeSet[Rect2](p_ret, val)
		case Rect2i:
			gd.UnsafeSet[Rect2i](p_ret, val)
		case Vector3:
			gd.UnsafeSet[Vector3](p_ret, val)
		case Vector3i:
			gd.UnsafeSet[Vector3i](p_ret, val)
		case Transform2D:
			gd.UnsafeSet[Transform2D](p_ret, val)
		case Vector4:
			gd.UnsafeSet[Vector4](p_ret, val)
		case Vector4i:
			gd.UnsafeSet[Vector4i](p_ret, val)
		case Plane:
			gd.UnsafeSet[Plane](p_ret, val)
		case Quaternion:
			gd.UnsafeSet[Quaternion](p_ret, val)
		case AABB:
			gd.UnsafeSet[AABB](p_ret, val)
		case Basis:
			gd.UnsafeSet[Basis](p_ret, val)
		case Transform3D:
			gd.UnsafeSet[Transform3D](p_ret, val)
		case Projection:
			gd.UnsafeSet[Projection](p_ret, val)
		case Color:
			gd.UnsafeSet[Color](p_ret, val)
		case StringName:
			gd.UnsafeSet[[1]uintptr](p_ret, discreet.Get(val))
			discreet.End(val)
		case NodePath:
			gd.UnsafeSet[[1]uintptr](p_ret, discreet.Get(val))
			discreet.End(val)
		case RID:
			gd.UnsafeSet[RID](p_ret, val)
		case Object:
			gd.UnsafeSet[uintptr](p_ret, discreet.Get(val)[0])
			_, ok := godot.Instances[discreet.Get(val)[0]]
			if !ok {
				discreet.End(val)
			}
		case Callable:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
			discreet.End(val)
		case gd.Signal:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
			discreet.End(val)
		case Dictionary:
			gd.UnsafeSet[[1]uintptr](p_ret, discreet.Get(val))
			discreet.End(val)
		case Array:
			gd.UnsafeSet[[1]uintptr](p_ret, discreet.Get(val))
			discreet.End(val)
		case PackedByteArray:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		case PackedInt32Array:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		case PackedInt64Array:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		case PackedFloat32Array:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		case PackedFloat64Array:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		case PackedStringArray:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		case PackedVector2Array:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		case PackedVector3Array:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		case PackedColorArray:
			gd.UnsafeSet[[2]uintptr](p_ret, discreet.Get(val))
		default:
			panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", method.Type().Out(0)))
		}
	}
}
