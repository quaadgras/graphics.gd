//go:build !generate

package gd

import (
	"fmt"
	"reflect"

	"runtime.link/mmm"
)

func registerMethods(godot Context, class StringName, rtype reflect.Type) {
	classTypePtr := reflect.PointerTo(rtype)
	for i := 0; i < classTypePtr.NumMethod(); i++ {
		i := i
		method := classTypePtr.Method(i)
		if !method.IsExported() || method.Type.NumIn() < 2 || method.Type.In(1) != reflect.TypeOf(Context{}) || method.Type.NumOut() > 0 {
			continue
		}
		var arguments = make([]PropertyInfo, 0, method.Type.NumIn()-2)
		var metadatas = make([]ClassMethodArgumentMetadata, 0, method.Type.NumIn()-2)
		for i := 2; i < method.Type.NumIn(); i++ {
			arguments = append(arguments, PropertyInfo{
				Type:      variantTypeOf(method.Type.In(i)),
				Name:      godot.StringName("arg" + fmt.Sprint(i)),
				ClassName: godot.StringName(classNameOf(method.Type.In(i))),
			})
			metadatas = append(metadatas, 0)
		}
		var returns *PropertyInfo
		var returnMetadata ClassMethodArgumentMetadata
		if method.Type.NumOut() > 0 {
			returns = &PropertyInfo{
				Type:      variantTypeOf(method.Type.Out(0)),
				Name:      godot.StringName("result"),
				ClassName: godot.StringName(classNameOf(method.Type.Out(0))),
			}
			returnMetadata = 0
		}
		godot.API.ClassDB.RegisterClassMethod(godot, godot.API.ExtensionToken, class, Method{
			Name: godot.StringName(method.Name),
			// FIXME type check and return an error if argumetns are invalid.
			Call: func(instance any, v ...Variant) (Variant, error) {
				ctx := NewContext(godot.API)
				defer ctx.End()

				var args = make([]reflect.Value, len(v)+1)
				args[0] = reflect.ValueOf(ctx)
				for i := 0; i < len(v); i++ {
					if method.Type.In(i + 2).Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
						var obj = reflect.New(method.Type.In(i + 2))
						obj.Interface().(PointerToClass).SetPointer(letVariantAsPointerType[Pointer](ctx, v[i], TypeObject))
						args[i+1] = obj.Elem()
					} else {
						args[i+1] = reflect.ValueOf(v[i].Interface(ctx))
					}
				}
				rets := reflect.ValueOf(instance.(*instanceImplementation).Value).Method(i).Call(args)
				if len(rets) > 0 {
					return ctx.Variant(rets[0].Interface()), nil
				}
				return Variant{}, nil
			},
			PointerCall: func(instance any, args UnsafeArgs, ret UnsafeBack) {
				slowCall(godot.API, reflect.ValueOf(instance.(*instanceImplementation).Value).Method(i), args, ret)
			},
			Arguments:           arguments,
			ArgumentsMetadata:   metadatas,
			ReturnValueInfo:     returns,
			ReturnValueMetadata: returnMetadata,
		})
	}
}

func slowCall(godot *API, method reflect.Value, p_args UnsafeArgs, p_ret UnsafeBack) {
	var ctx = NewContext(godot)
	defer ctx.End()
	var (
		args = make([]reflect.Value, method.Type().NumIn())
	)
	args[0] = reflect.ValueOf(ctx)
	for i := 1; i < method.Type().NumIn(); i++ {
		switch method.Type().In(i) {
		case reflect.TypeOf(Bool(false)):
			args[i] = reflect.ValueOf(UnsafeGet[Bool](p_args, i-1))
		case reflect.TypeOf(Int(0)):
			args[i] = reflect.ValueOf(UnsafeGet[Int](p_args, i-1))
		case reflect.TypeOf(Float(0)):
			args[i] = reflect.ValueOf(UnsafeGet[Float](p_args, i-1))
		case reflect.TypeOf(String{}):
			ptr := UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[String](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(Vector2{}):
			args[i] = reflect.ValueOf(UnsafeGet[Vector2](p_args, i-1))
		case reflect.TypeOf(Vector2i{}):
			args[i] = reflect.ValueOf(UnsafeGet[Vector2i](p_args, i-1))
		case reflect.TypeOf(Rect2{}):
			args[i] = reflect.ValueOf(UnsafeGet[Rect2](p_args, i-1))
		case reflect.TypeOf(Rect2i{}):
			args[i] = reflect.ValueOf(UnsafeGet[Rect2i](p_args, i-1))
		case reflect.TypeOf(Vector3{}):
			args[i] = reflect.ValueOf(UnsafeGet[Vector3](p_args, i-1))
		case reflect.TypeOf(Vector3i{}):
			args[i] = reflect.ValueOf(UnsafeGet[Vector3i](p_args, i-1))
		case reflect.TypeOf(Transform2D{}):
			args[i] = reflect.ValueOf(UnsafeGet[Transform2D](p_args, i-1))
		case reflect.TypeOf(Vector4{}):
			args[i] = reflect.ValueOf(UnsafeGet[Vector4](p_args, i-1))
		case reflect.TypeOf(Vector4i{}):
			args[i] = reflect.ValueOf(UnsafeGet[Vector4i](p_args, i-1))
		case reflect.TypeOf(Plane{}):
			args[i] = reflect.ValueOf(UnsafeGet[Plane](p_args, i-1))
		case reflect.TypeOf(Quaternion{}):
			args[i] = reflect.ValueOf(UnsafeGet[Quaternion](p_args, i-1))
		case reflect.TypeOf(AABB{}):
			args[i] = reflect.ValueOf(UnsafeGet[AABB](p_args, i-1))
		case reflect.TypeOf(Basis{}):
			args[i] = reflect.ValueOf(UnsafeGet[Basis](p_args, i-1))
		case reflect.TypeOf(Transform3D{}):
			args[i] = reflect.ValueOf(UnsafeGet[Transform3D](p_args, i-1))
		case reflect.TypeOf(Projection{}):
			args[i] = reflect.ValueOf(UnsafeGet[Projection](p_args, i-1))
		case reflect.TypeOf(Color{}):
			args[i] = reflect.ValueOf(UnsafeGet[Color](p_args, i-1))
		case reflect.TypeOf(StringName{}):
			ptr := UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[StringName](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(NodePath{}):
			ptr := UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[NodePath](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(RID(0)):
			args[i] = reflect.ValueOf(UnsafeGet[RID](p_args, i-1))
		case reflect.TypeOf(Object{}):
			ptr := UnsafeGet[uintptr](p_args, i-1)
			var obj Object
			obj.SetPointer(mmm.Let[Pointer](ctx.Lifetime, ctx.API, ptr))
			args[i] = reflect.ValueOf(obj)
		case reflect.TypeOf(Callable{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[Callable](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(Signal{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[Signal](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(Dictionary{}):
			ptr := UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[Dictionary](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(Array{}):
			ptr := UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedByteArray{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedByteArray](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedInt32Array{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedInt32Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedInt64Array{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedInt64Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedFloat32Array{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedFloat32Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedFloat64Array{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedFloat64Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedStringArray{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedStringArray](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedVector2Array{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedVector2Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedVector3Array{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedVector3Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedColorArray{}):
			ptr := UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedColorArray](ctx.Lifetime, ctx.API, ptr))
		default:
			panic(fmt.Sprintf("gdextension: unsupported Godot -> Go type %v", method.Type().In(i)))
		}
	}
	rrets := method.Call(args)
	if len(rrets) > 0 {
		switch val := rrets[0].Interface().(type) {
		case Bool:
			UnsafeSet[Bool](p_ret, val)
		case Int:
			UnsafeSet[Int](p_ret, val)
		case Float:
			UnsafeSet[Float](p_ret, val)
		case String:
			UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case Vector2:
			UnsafeSet[Vector2](p_ret, val)
		case Vector2i:
			UnsafeSet[Vector2i](p_ret, val)
		case Rect2:
			UnsafeSet[Rect2](p_ret, val)
		case Rect2i:
			UnsafeSet[Rect2i](p_ret, val)
		case Vector3:
			UnsafeSet[Vector3](p_ret, val)
		case Vector3i:
			UnsafeSet[Vector3i](p_ret, val)
		case Transform2D:
			UnsafeSet[Transform2D](p_ret, val)
		case Vector4:
			UnsafeSet[Vector4](p_ret, val)
		case Vector4i:
			UnsafeSet[Vector4i](p_ret, val)
		case Plane:
			UnsafeSet[Plane](p_ret, val)
		case Quaternion:
			UnsafeSet[Quaternion](p_ret, val)
		case AABB:
			UnsafeSet[AABB](p_ret, val)
		case Basis:
			UnsafeSet[Basis](p_ret, val)
		case Transform3D:
			UnsafeSet[Transform3D](p_ret, val)
		case Projection:
			UnsafeSet[Projection](p_ret, val)
		case Color:
			UnsafeSet[Color](p_ret, val)
		case StringName:
			UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case NodePath:
			UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case RID:
			UnsafeSet[RID](p_ret, val)
		case Object:
			UnsafeSet[uintptr](p_ret, mmm.Get(val.AsPointer()))
		case Callable:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case Signal:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case Dictionary:
			UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case Array:
			UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case PackedByteArray:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedInt32Array:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedInt64Array:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedFloat32Array:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedFloat64Array:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedStringArray:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedVector2Array:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedVector3Array:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedColorArray:
			UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		default:
			panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", method.Type().Out(0)))
		}
	}
}
