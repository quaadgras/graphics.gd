package classdb

import (
	"fmt"
	"reflect"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"

	EngineClass "graphics.gd/classdb/Engine"
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
			vtype, ok := propertyOf(reflect.StructField{Name: "arg" + fmt.Sprint(i), Type: method.Type.In(i)})
			if ok {
				arguments = append(arguments, vtype)
				metadatas = append(metadatas, 0)
			}
		}
		var returns *gd.PropertyInfo
		var returnMetadata gd.ClassMethodArgumentMetadata
		if method.Type.NumOut() > 0 {
			property, ok := propertyOf(reflect.StructField{Name: "result", Type: method.Type.Out(0)})
			if ok {
				returns = &property
				returnMetadata = 0
			}
		}

		gd.Global.ClassDB.RegisterClassMethod(gd.Global.ExtensionToken, class, gd.Method{
			Name: gd.NewStringName(method.Name),
			Call: variantCall(method),
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

func variantCall(method reflect.Method) func(instance any, v ...gd.Variant) (gd.Variant, error) {
	return func(instance any, v ...gd.Variant) (result gd.Variant, err error) {
		var args = make([]reflect.Value, len(v))
		for i := 0; i < method.Type.NumIn()-1; i++ {
			args[i], err = gd.ConvertToDesiredGoType(v[i], method.Type.In(i+1))
			if err != nil {
				EngineClass.Raise(err)
				return gd.Variant{}, err
			}
		}
		extensionInstance := instance.(*instanceImplementation).Value
		rets := reflect.ValueOf(extensionInstance).Method(method.Index).Call(args)
		if len(rets) > 0 {
			return gd.NewVariant(rets[0].Interface()), nil
		}
		return gd.Variant{}, nil
	}
}

func slowCall(hasContext bool, method reflect.Value, p_args gd.UnsafeArgs, p_ret gd.UnsafeBack) {
	var (
		args = make([]reflect.Value, method.Type().NumIn())
	)
	var err error
	var offset = 0
	for i := offset; i < method.Type().NumIn(); i++ {
		rtype := method.Type().In(i)
		vtype, ok := gd.VariantTypeOf(rtype)
		if ok {
			var value any
			switch vtype {
			case gd.TypeBool:
				value = gd.UnsafeGet[gd.Bool](p_args, i-offset)
			case gd.TypeInt:
				value = gd.UnsafeGet[gd.Int](p_args, i-offset)
			case gd.TypeFloat:
				value = gd.UnsafeGet[gd.Float](p_args, i-offset)
			case gd.TypeString:
				ptr := gd.UnsafeGet[[1]gd.EnginePointer](p_args, i-offset)
				val := pointers.Let[gd.String](ptr)
				defer val.Free()
				value = val
			case gd.TypeVector2:
				value = gd.UnsafeGet[gd.Vector2](p_args, i-offset)
			case gd.TypeVector2i:
				value = gd.UnsafeGet[gd.Vector2i](p_args, i-offset)
			case gd.TypeRect2:
				value = gd.UnsafeGet[gd.Rect2](p_args, i-offset)
			case gd.TypeRect2i:
				value = gd.UnsafeGet[gd.Rect2i](p_args, i-offset)
			case gd.TypeVector3:
				value = gd.UnsafeGet[gd.Vector3](p_args, i-offset)
			case gd.TypeVector3i:
				value = gd.UnsafeGet[gd.Vector3i](p_args, i-offset)
			case gd.TypeTransform2D:
				value = gd.UnsafeGet[gd.Transform2D](p_args, i-offset)
			case gd.TypeVector4:
				value = gd.UnsafeGet[gd.Vector4](p_args, i-offset)
			case gd.TypeVector4i:
				value = gd.UnsafeGet[gd.Vector4i](p_args, i-offset)
			case gd.TypePlane:
				value = gd.UnsafeGet[gd.Plane](p_args, i-offset)
			case gd.TypeQuaternion:
				value = gd.UnsafeGet[gd.Quaternion](p_args, i-offset)
			case gd.TypeAABB:
				value = gd.UnsafeGet[gd.AABB](p_args, i-offset)
			case gd.TypeBasis:
				value = gd.UnsafeGet[gd.Basis](p_args, i-offset)
			case gd.TypeTransform3D:
				value = gd.UnsafeGet[gd.Transform3D](p_args, i-offset)
			case gd.TypeProjection:
				value = gd.UnsafeGet[gd.Projection](p_args, i-offset)
			case gd.TypeColor:
				value = gd.UnsafeGet[gd.Color](p_args, i-offset)
			case gd.TypeStringName:
				ptr := gd.UnsafeGet[[1]gd.EnginePointer](p_args, i-offset)
				val := pointers.Let[gd.StringName](ptr)
				value = val
			case gd.TypeNodePath:
				ptr := gd.UnsafeGet[[1]gd.EnginePointer](p_args, i-offset)
				val := pointers.Let[gd.NodePath](ptr)
				value = val
			case gd.TypeRID:
				value = gd.UnsafeGet[gd.RID](p_args, i-offset)
			case gd.TypeObject:
				ptr := gd.UnsafeGet[[3]uint64](p_args, i-offset)
				val := [1]gd.Object{pointers.Let[gd.Object](ptr)}
				value = val
			case gd.TypeCallable:
				ptr := gd.UnsafeGet[[2]uint64](p_args, i-offset)
				val := pointers.Let[gd.Callable](ptr)
				defer val.Free()
				value = val
			case gd.TypeSignal:
				ptr := gd.UnsafeGet[[2]uint64](p_args, i-offset)
				val := pointers.Let[gd.Signal](ptr)
				value = val
			case gd.TypeDictionary:
				ptr := gd.UnsafeGet[[1]gd.EnginePointer](p_args, i-offset)
				val := pointers.Let[gd.Dictionary](ptr)
				value = val
			case gd.TypeArray:
				ptr := gd.UnsafeGet[[1]gd.EnginePointer](p_args, i-offset)
				val := pointers.Let[gd.Array](ptr)
				value = val
			case gd.TypePackedByteArray:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedByteArray](ptr)
				value = val
			case gd.TypePackedInt32Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedInt32Array](ptr)
				value = val
			case gd.TypePackedInt64Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedInt64Array](ptr)
				value = val
			case gd.TypePackedFloat32Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedFloat32Array](ptr)
				value = val
			case gd.TypePackedFloat64Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedFloat64Array](ptr)
				value = val
			case gd.TypePackedStringArray:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedStringArray](ptr)
				value = val
			case gd.TypePackedVector2Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedVector2Array](ptr)
				value = val
			case gd.TypePackedVector3Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedVector3Array](ptr)
				value = val
			case gd.TypePackedColorArray:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedColorArray](ptr)
				value = val
			case gd.TypePackedVector4Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedVector4Array](ptr)
				value = val
			case gd.TypeNil:
				value = nil
			default:
				panic(fmt.Sprintf("gdextension: unsupported Godot -> Go variant type %v", vtype))
			}
			args[i], err = gd.ConvertToDesiredGoType(value, rtype)
			if err != nil {
				panic(err)
			}
		}
	}
	rrets := method.Call(args)
	if len(rrets) > 0 {
		result := rrets[0]
		vtype, ok := gd.VariantTypeOf(result.Type())
		if !ok {
			panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", result.Type()))
		}
		vvalue := gd.NewVariant(result.Interface())
		if vvalue.Type() != vtype {
			panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", result.Type()))
		}
		switch val := vvalue.Interface().(type) {
		case gd.Bool:
			gd.UnsafeSet[gd.Bool](p_ret, val)
		case gd.Int:
			gd.UnsafeSet[gd.Int](p_ret, val)
		case gd.Float:
			gd.UnsafeSet[gd.Float](p_ret, val)
		case gd.String:
			gd.UnsafeSet[[1]gd.EnginePointer](p_ret, pointers.Get(val))
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
			gd.UnsafeSet[[1]gd.EnginePointer](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.NodePath:
			gd.UnsafeSet[[1]gd.EnginePointer](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.RID:
			gd.UnsafeSet[gd.RID](p_ret, val)
		case [1]gd.Object:
			gd.UnsafeSet[gd.EnginePointer](p_ret, gd.EnginePointer(pointers.Get(val[0])[0]))
			_, ok := gd.ExtensionInstances.Load(pointers.Get(val[0])[0])
			if !ok {
				pointers.End(val[0])
			}
		case gd.Callable:
			gd.UnsafeSet[[2]uint64](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Signal:
			gd.UnsafeSet[[2]uint64](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Dictionary:
			gd.UnsafeSet[[1]gd.EnginePointer](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Array:
			gd.UnsafeSet[[1]gd.EnginePointer](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.PackedByteArray:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		case gd.PackedInt32Array:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		case gd.PackedInt64Array:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		case gd.PackedFloat32Array:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		case gd.PackedFloat64Array:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		case gd.PackedStringArray:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		case gd.PackedVector2Array:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		case gd.PackedVector3Array:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		case gd.PackedColorArray:
			gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(val))
		default:
			panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", method.Type().Out(0)))
		}
	}
}
