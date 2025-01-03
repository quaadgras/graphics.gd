package defined

import (
	"fmt"
	"reflect"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/String"
)

func registerMethods(class gd.StringName, rtype reflect.Type) {
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
						*(*gd.Object)(obj.UnsafePointer()) = gd.LetVariantAsPointerType[gd.Object](v[i], gd.TypeObject)
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
	var (
		args = make([]reflect.Value, method.Type().NumIn())
	)
	var offset = 0
	for i := offset; i < method.Type().NumIn(); i++ {
		switch method.Type().In(i) {
		case reflect.TypeOf(gd.Bool(false)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Bool](p_args, i-offset))
		case reflect.TypeOf(gd.Int(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Int](p_args, i-offset))
		case reflect.TypeOf(gd.Float(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Float](p_args, i-offset))
		case reflect.TypeOf(gd.String{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := pointers.Let[gd.String](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.Vector2{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector2](p_args, i-offset))
		case reflect.TypeOf(gd.Vector2i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector2i](p_args, i-offset))
		case reflect.TypeOf(gd.Rect2{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Rect2](p_args, i-offset))
		case reflect.TypeOf(gd.Rect2i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Rect2i](p_args, i-offset))
		case reflect.TypeOf(gd.Vector3{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector3](p_args, i-offset))
		case reflect.TypeOf(gd.Vector3i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector3i](p_args, i-offset))
		case reflect.TypeOf(gd.Transform2D{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Transform2D](p_args, i-offset))
		case reflect.TypeOf(gd.Vector4{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector4](p_args, i-offset))
		case reflect.TypeOf(gd.Vector4i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector4i](p_args, i-offset))
		case reflect.TypeOf(gd.Plane{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Plane](p_args, i-offset))
		case reflect.TypeOf(gd.Quaternion{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Quaternion](p_args, i-offset))
		case reflect.TypeOf(gd.AABB{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.AABB](p_args, i-offset))
		case reflect.TypeOf(gd.Basis{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Basis](p_args, i-offset))
		case reflect.TypeOf(gd.Transform3D{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Transform3D](p_args, i-offset))
		case reflect.TypeOf(gd.Projection{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Projection](p_args, i-offset))
		case reflect.TypeOf(gd.Color{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Color](p_args, i-offset))
		case reflect.TypeOf(gd.StringName{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := pointers.Let[gd.StringName](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.NodePath{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := pointers.Let[gd.NodePath](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.RID(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.RID](p_args, i-offset))
		case reflect.TypeOf(gd.Object{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := pointers.Let[gd.Object]([3]uintptr{ptr[0]})
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.Callable{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.Callable](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.Signal{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.Signal](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.Dictionary{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := pointers.Let[gd.Dictionary](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.Array{}):
			ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
			val := pointers.Let[gd.Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedByteArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedByteArray](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedInt32Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedInt32Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedInt64Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedInt64Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedFloat32Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedFloat32Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedFloat64Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedFloat64Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedStringArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedStringArray](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedVector2Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedVector2Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedVector3Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedVector3Array](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		case reflect.TypeOf(gd.PackedColorArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-offset)
			val := pointers.Let[gd.PackedColorArray](ptr)
			defer val.Free()
			args[i] = reflect.ValueOf(val)
		default:
			args[i] = reflect.New(method.Type().In(i)).Elem()
			switch method.Type().In(i).Kind() {
			case reflect.String:
				ptr := gd.UnsafeGet[[1]uintptr](p_args, i-offset)
				val := pointers.Let[gd.String](ptr)
				defer val.Free()
				args[i].SetString(val.String())
			default:
				panic(fmt.Sprintf("gdextension: unsupported Godot -> Go type %v", method.Type().In(i)))
			}
		}
	}
	rrets := method.Call(args)
	if len(rrets) > 0 {
		switch val := rrets[0].Interface().(type) {
		case gd.Bool:
			gd.UnsafeSet[gd.Bool](p_ret, val)
		case gd.Int:
			gd.UnsafeSet[gd.Int](p_ret, val)
		case gd.Float:
			gd.UnsafeSet[gd.Float](p_ret, val)
		case gd.String:
			gd.UnsafeSet[[1]uintptr](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Vector2:
			gd.UnsafeSet[gd.Vector2](p_ret, val)
		case gd.Vector2i:
			gd.UnsafeSet[gd.Vector2i](p_ret, val)
		case gd.Rect2:
			gd.UnsafeSet[gd.Rect2](p_ret, val)
		case gd.Rect2i:
			gd.UnsafeSet[gd.Rect2i](p_ret, val)
		case gd.Vector3:
			gd.UnsafeSet[gd.Vector3](p_ret, val)
		case gd.Vector3i:
			gd.UnsafeSet[gd.Vector3i](p_ret, val)
		case gd.Transform2D:
			gd.UnsafeSet[gd.Transform2D](p_ret, val)
		case gd.Vector4:
			gd.UnsafeSet[gd.Vector4](p_ret, val)
		case gd.Vector4i:
			gd.UnsafeSet[gd.Vector4i](p_ret, val)
		case gd.Plane:
			gd.UnsafeSet[gd.Plane](p_ret, val)
		case gd.Quaternion:
			gd.UnsafeSet[gd.Quaternion](p_ret, val)
		case gd.AABB:
			gd.UnsafeSet[gd.AABB](p_ret, val)
		case gd.Basis:
			gd.UnsafeSet[gd.Basis](p_ret, val)
		case gd.Transform3D:
			gd.UnsafeSet[gd.Transform3D](p_ret, val)
		case gd.Projection:
			gd.UnsafeSet[gd.Projection](p_ret, val)
		case gd.Color:
			gd.UnsafeSet[gd.Color](p_ret, val)
		case gd.StringName:
			gd.UnsafeSet[[1]uintptr](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.NodePath:
			gd.UnsafeSet[[1]uintptr](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.RID:
			gd.UnsafeSet[gd.RID](p_ret, val)
		case gd.Object:
			gd.UnsafeSet[uintptr](p_ret, pointers.Get(val)[0])
			_, ok := gd.ExtensionInstances.Load(pointers.Get(val)[0])
			if !ok {
				pointers.End(val)
			}
		case gd.Callable:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Signal:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Dictionary:
			gd.UnsafeSet[[1]uintptr](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Array:
			gd.UnsafeSet[[1]uintptr](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.PackedByteArray:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		case gd.PackedInt32Array:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		case gd.PackedInt64Array:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		case gd.PackedFloat32Array:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		case gd.PackedFloat64Array:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		case gd.PackedStringArray:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		case gd.PackedVector2Array:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		case gd.PackedVector3Array:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		case gd.PackedColorArray:
			gd.UnsafeSet[[2]uintptr](p_ret, pointers.Get(val))
		default:
			switch rrets[0].Type().Kind() {
			case reflect.String:
				gd.UnsafeSet[[1]uintptr](p_ret, pointers.Get(String.New(rrets[0])))
			default:
				panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", method.Type().Out(0)))
			}
		}
	}
}
