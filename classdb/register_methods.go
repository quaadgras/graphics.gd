package classdb

import (
	"fmt"
	"reflect"
	"strings"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Dictionary"
	"graphics.gd/variant/Packed"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/String"

	EngineClass "graphics.gd/classdb/Engine"
)

func registerMethods(class gd.StringName, rtype reflect.Type, renames map[uintptr]string) {
	classTypePtr := reflect.PointerTo(rtype)
	var elligible []reflect.Method
	for i := range classTypePtr.NumMethod() {
		i := i
		method := classTypePtr.Method(i)
		if !method.IsExported() || method.Type.NumIn() < 1 {
			continue
		}
		if strings.HasPrefix(method.Name, "As") || method.Name == "Super" || method.Name == "UnsafePointer" {
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
		elligible = append(elligible, method)
	}
	var minfo = gdextension.Host.ClassDB.MethodList.Make(len(elligible))
	for _, method := range elligible {
		if name, ok := renames[method.Func.Pointer()]; ok {
			method.Name = name
		} else {
			method.Name = String.ToSnakeCase(method.Name)
		}
		var offset = 0
		var arguments = gdextension.Host.ClassDB.PropertyList.Make(method.Type.NumIn() - 1 - offset)
		for i := 1 + offset; i < method.Type.NumIn(); i++ {
			if !propertyOf(class, reflect.StructField{Name: "arg" + fmt.Sprint(i), Type: method.Type.In(i)}, arguments) {
				panic(fmt.Sprintf("gdextension: method %s has an argument of unsupported type %v", method.Name, method.Type.In(i)))
			}
		}
		var returns = gdextension.Host.ClassDB.PropertyList.Make(method.Type.NumOut())
		if method.Type.NumOut() > 0 {
			if !propertyOf(class, reflect.StructField{Name: "result", Type: method.Type.Out(0)}, returns) {
				panic(fmt.Sprintf("gdextension: method %s has a return value of unsupported type %v", method.Name, method.Type.Out(0)))
			}
		}
		call := variantCall(method)
		gdextension.Host.ClassDB.MethodList.Push(minfo,
			pointers.Get(gd.NewStringName(method.Name)),
			gdextension.FunctionID(cgoNewHandle(&methodImplementation{
				arg_count: method.Type.NumIn(),
				dynamic:   call,
				checked: func(instance any, args gd.Address, ret gd.Address) {
					extensionInstance := instance.(*instanceImplementation).Value
					slowCall(false, reflect.ValueOf(extensionInstance).Method(method.Index), args, ret)
				},
				variant: func(instance any, v ...gd.Variant) gd.Variant {
					result, _ := call(instance, v...)
					return result
				},
			})),
			0,
			returns,
			arguments,
			0,
			nil,
		)
		defer gdextension.Host.ClassDB.PropertyList.Free(arguments)
		defer gdextension.Host.ClassDB.PropertyList.Free(returns)
	}
	gdextension.Host.ClassDB.Register.Methods(pointers.Get(class), minfo)
	gdextension.Host.ClassDB.MethodList.Free(minfo)
}

func registerStaticMethod(class gd.StringName, name string, fn reflect.Value) {
	ftype := fn.Type()
	var arguments = gdextension.Host.ClassDB.PropertyList.Make(ftype.NumIn())
	for i := 0; i < ftype.NumIn(); i++ {
		if !propertyOf(class, reflect.StructField{Name: "arg" + fmt.Sprint(i), Type: ftype.In(i)}, arguments) {
			panic(fmt.Sprintf("gdextension: method %s has an argument of unsupported type %v", name, ftype.In(i)))
		}
	}
	var returns = gdextension.Host.ClassDB.PropertyList.Make(ftype.NumIn())
	if ftype.NumOut() > 0 {
		if !propertyOf(class, reflect.StructField{Name: "result", Type: ftype.Out(0)}, returns) {
			panic(fmt.Sprintf("gdextension: method %s has a return value of unsupported type %v", name, ftype.Out(0)))
		}
	}
	var method = gdextension.Host.ClassDB.MethodList.Make(1)
	gdextension.Host.ClassDB.MethodList.Push(method,
		pointers.Get(gd.NewStringName(name)),
		gdextension.FunctionID(cgoNewHandle(variantCallStatic(fn))),
		gdextension.MethodFlagStatic,
		returns,
		arguments,
		0,
		nil,
	)
	gdextension.Host.ClassDB.Register.Methods(pointers.Get(class), method)
	gdextension.Host.ClassDB.MethodList.Free(method)
	gdextension.Host.ClassDB.PropertyList.Free(arguments)
	gdextension.Host.ClassDB.PropertyList.Free(returns)
}

type methodImplementation struct {
	arg_count int

	dynamic func(instance any, v ...gd.Variant) (gd.Variant, error)
	checked func(instance any, args gd.Address, ret gd.Address)
	variant func(instance any, v ...gd.Variant) gd.Variant
}

func variantCallStatic(fn reflect.Value) func(instance any, v ...gd.Variant) (gd.Variant, error) {
	return func(instance any, v ...gd.Variant) (result gd.Variant, err error) {
		var args = make([]reflect.Value, len(v))
		for i := 0; i < fn.Type().NumIn()-1; i++ {
			args[i], err = gd.ConvertToDesiredGoType(v[i], fn.Type().In(i+1))
			if err != nil {
				EngineClass.Raise(err)
				return gd.Variant{}, err
			}
		}
		rets := fn.Call(args)
		if len(rets) > 0 {
			return gd.NewVariant(rets[0].Interface()), nil
		}
		return gd.Variant{}, nil
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

func slowCall(hasContext bool, method reflect.Value, p_args gd.Address, p_ret gd.Address) {

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
			case gdextension.TypeBool:
				value = gd.UnsafeGet[bool](p_args, i-offset)
			case gdextension.TypeInt:
				value = gd.UnsafeGet[int64](p_args, i-offset)
			case gdextension.TypeFloat:
				value = gd.UnsafeGet[float64](p_args, i-offset)
			case gdextension.TypeString:
				ptr := gd.UnsafeGet[gdextension.String](p_args, i-offset)
				val := pointers.Let[gd.String](ptr)
				value = val
			case gdextension.TypeVector2:
				value = gd.UnsafeGet[gd.Vector2](p_args, i-offset)
			case gdextension.TypeVector2i:
				value = gd.UnsafeGet[gd.Vector2i](p_args, i-offset)
			case gdextension.TypeRect2:
				value = gd.UnsafeGet[gd.Rect2](p_args, i-offset)
			case gdextension.TypeRect2i:
				value = gd.UnsafeGet[gd.Rect2i](p_args, i-offset)
			case gdextension.TypeVector3:
				value = gd.UnsafeGet[gd.Vector3](p_args, i-offset)
			case gdextension.TypeVector3i:
				value = gd.UnsafeGet[gd.Vector3i](p_args, i-offset)
			case gdextension.TypeTransform2D:
				value = gd.UnsafeGet[gd.Transform2D](p_args, i-offset)
			case gdextension.TypeVector4:
				value = gd.UnsafeGet[gd.Vector4](p_args, i-offset)
			case gdextension.TypeVector4i:
				value = gd.UnsafeGet[gd.Vector4i](p_args, i-offset)
			case gdextension.TypePlane:
				value = gd.UnsafeGet[gd.Plane](p_args, i-offset)
			case gdextension.TypeQuaternion:
				value = gd.UnsafeGet[gd.Quaternion](p_args, i-offset)
			case gdextension.TypeAABB:
				value = gd.UnsafeGet[gd.AABB](p_args, i-offset)
			case gdextension.TypeBasis:
				value = gd.UnsafeGet[gd.Basis](p_args, i-offset)
			case gdextension.TypeTransform3D:
				value = gd.UnsafeGet[gd.Transform3D](p_args, i-offset)
			case gdextension.TypeProjection:
				value = gd.UnsafeGet[gd.Projection](p_args, i-offset)
			case gdextension.TypeColor:
				value = gd.UnsafeGet[gd.Color](p_args, i-offset)
			case gdextension.TypeStringName:
				ptr := gd.UnsafeGet[gdextension.StringName](p_args, i-offset)
				val := pointers.Let[gd.StringName](ptr)
				value = val
			case gdextension.TypeNodePath:
				ptr := gd.UnsafeGet[gdextension.NodePath](p_args, i-offset)
				val := pointers.Let[gd.NodePath](ptr)
				value = val
			case gdextension.TypeRID:
				value = gd.UnsafeGet[gd.RID](p_args, i-offset)
			case gdextension.TypeObject:
				ptr := gd.UnsafeGet[gd.EnginePointer](p_args, i-offset)
				val := [1]gd.Object{pointers.Let[gd.Object]([3]uint64{uint64(ptr)})}
				value = val
			case gdextension.TypeCallable:
				ptr := gd.UnsafeGet[[2]uint64](p_args, i-offset)
				val := pointers.Let[gd.Callable](ptr)
				defer val.Free()
				value = val
			case gdextension.TypeSignal:
				ptr := gd.UnsafeGet[[2]uint64](p_args, i-offset)
				val := pointers.Let[gd.Signal](ptr)
				value = val
			case gdextension.TypeDictionary:
				ptr := gd.UnsafeGet[gdextension.Dictionary](p_args, i-offset)
				val := pointers.Let[gd.Dictionary](ptr)
				value = val
			case gdextension.TypeArray:
				ptr := gd.UnsafeGet[gdextension.Array](p_args, i-offset)
				val := pointers.Let[gd.Array](ptr)
				value = val
			case gdextension.TypePackedByteArray:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedByteArray](ptr)
				value = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(val)))
			case gdextension.TypePackedInt32Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedInt32Array](ptr)
				value = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(val)))
			case gdextension.TypePackedInt64Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedInt64Array](ptr)
				value = Packed.Array[int64](Array.Through(gd.PackedProxy[gd.PackedInt64Array, int64]{}, pointers.Pack(val)))
			case gdextension.TypePackedFloat32Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedFloat32Array](ptr)
				value = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(val)))
			case gdextension.TypePackedFloat64Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedFloat64Array](ptr)
				value = Packed.Array[float64](Array.Through(gd.PackedProxy[gd.PackedFloat64Array, float64]{}, pointers.Pack(val)))
			case gdextension.TypePackedStringArray:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedStringArray](ptr)
				value = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(val)))
			case gdextension.TypePackedVector2Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedVector2Array](ptr)
				value = Packed.Array[gd.Vector2](Array.Through(gd.PackedProxy[gd.PackedVector2Array, gd.Vector2]{}, pointers.Pack(val)))
			case gdextension.TypePackedVector3Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedVector3Array](ptr)
				value = Packed.Array[gd.Vector3](Array.Through(gd.PackedProxy[gd.PackedVector3Array, gd.Vector3]{}, pointers.Pack(val)))
			case gdextension.TypePackedColorArray:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedColorArray](ptr)
				value = Packed.Array[gd.Color](Array.Through(gd.PackedProxy[gd.PackedColorArray, gd.Color]{}, pointers.Pack(val)))
			case gdextension.TypePackedVector4Array:
				ptr := gd.UnsafeGet[gd.PackedPointers](p_args, i-offset)
				val := pointers.Let[gd.PackedVector4Array](ptr)
				value = Packed.Array[gd.Vector4](Array.Through(gd.PackedProxy[gd.PackedVector4Array, gd.Vector4]{}, pointers.Pack(val)))
			case gdextension.TypeNil:
				value = nil
			default:
				panic(fmt.Sprintf("gdextension: unsupported Godot -> Go variant type %v", vtype))
			}
			args[i], err = gd.ConvertToDesiredGoType(value, rtype)
			if err != nil {
				fmt.Println(err)
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
		converted := false
		ivalue := result.Interface()
	retry:
		switch val := ivalue.(type) {
		case bool:
			gd.UnsafeSet[bool](p_ret, val)
		case gd.Int:
			gd.UnsafeSet[gd.Int](p_ret, val)
		case gd.Float:
			gd.UnsafeSet[gd.Float](p_ret, val)
		case gd.String:
			gd.UnsafeSet[gdextension.String](p_ret, pointers.Get(val))
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
			gd.UnsafeSet[gdextension.StringName](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.NodePath:
			gd.UnsafeSet[gdextension.NodePath](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Object:
			gd.UnsafeSet[gd.EnginePointer](p_ret, gd.PointerWithOwnershipTransferredToGodot(val))
		case [1]gd.Object:
			gd.UnsafeSet[gd.EnginePointer](p_ret, gd.PointerWithOwnershipTransferredToGodot(val[0]))
		case gd.Callable:
			gd.UnsafeSet[[2]uint64](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Signal:
			gd.UnsafeSet[[2]uint64](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Dictionary:
			gd.UnsafeSet[gdextension.Dictionary](p_ret, pointers.Get(val))
			pointers.End(val)
		case gd.Array:
			gd.UnsafeSet[gdextension.Array](p_ret, pointers.Get(val))
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
		case String.Readable:
			in := gd.InternalString(val)
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gdextension.String](p_ret, raw)
			} else {
				gd.UnsafeSet[gdextension.String](p_ret, pointers.Get(in))
			}
		case String.Name:
			in := gd.InternalStringName(val)
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gdextension.StringName](p_ret, raw)
			} else {
				gd.UnsafeSet[gdextension.StringName](p_ret, pointers.Get(in))
			}
		case Path.ToNode:
			in := gd.InternalNodePath(val)
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gdextension.NodePath](p_ret, raw)
			} else {
				gd.UnsafeSet[gdextension.NodePath](p_ret, pointers.Get(in))
			}
		case Callable.Function:
			in := gd.InternalCallable(val)
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[[2]uint64](p_ret, raw)
			} else {
				gd.UnsafeSet[[2]uint64](p_ret, pointers.Get(in))
			}
		case Dictionary.Any:
			in := gd.InternalDictionary(val)
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gdextension.Dictionary](p_ret, raw)
			} else {
				gd.UnsafeSet[gdextension.Dictionary](p_ret, pointers.Get(in))
			}
		case Array.Any:
			in := gd.InternalArray(val)
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gdextension.Array](p_ret, raw)
			} else {
				gd.UnsafeSet[gdextension.Array](p_ret, pointers.Get(in))
			}
		case Packed.Bytes:
			in := gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case RID.Any:
			gd.UnsafeSet[RID.Any](p_ret, val)
		case Packed.Array[int32]:
			in := gd.InternalPacked[gd.PackedInt32Array, int32](Packed.Array[int32](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case Packed.Array[int64]:
			in := gd.InternalPacked[gd.PackedInt64Array, int64](Packed.Array[int64](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case Packed.Array[float32]:
			in := gd.InternalPacked[gd.PackedFloat32Array, float32](Packed.Array[float32](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case Packed.Array[float64]:
			in := gd.InternalPacked[gd.PackedFloat64Array, float64](Packed.Array[float64](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case Packed.Strings:
			in := gd.InternalPackedStrings(val)
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case Packed.Array[gd.Vector2]:
			in := gd.InternalPacked[gd.PackedVector2Array, gd.Vector2](Packed.Array[gd.Vector2](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case Packed.Array[gd.Vector3]:
			in := gd.InternalPacked[gd.PackedVector3Array, gd.Vector3](Packed.Array[gd.Vector3](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case Packed.Array[gd.Color]:
			in := gd.InternalPacked[gd.PackedColorArray, gd.Color](Packed.Array[gd.Color](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		case Packed.Array[gd.Vector4]:
			in := gd.InternalPacked[gd.PackedVector4Array, gd.Vector4](Packed.Array[gd.Vector4](val))
			raw, ok := pointers.End(in)
			if ok {
				gd.UnsafeSet[gd.PackedPointers](p_ret, raw)
			} else {
				gd.UnsafeSet[gd.PackedPointers](p_ret, pointers.Get(in))
			}
		default:
			if !converted {
				vvalue := gd.NewVariant(result.Interface())
				if vvalue.Type() != vtype {
					panic(fmt.Sprintf("gdextension: expected %v, got %v (%s)", vtype, vvalue.Type(), result.Type()))
				}
				ivalue = vvalue.Interface()
				converted = true
				goto retry
			}
			panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", reflect.TypeOf(val)))
		}
	}
}
