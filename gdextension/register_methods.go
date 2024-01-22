package gdextension

import (
	"fmt"
	"reflect"

	gd "grow.graphics/gd/internal"
	"runtime.link/mmm"
)

func registerMethods(godot gd.Context, class gd.StringName, rtype reflect.Type) {
	classTypePtr := reflect.PointerTo(rtype)
	for i := 0; i < classTypePtr.NumMethod(); i++ {
		i := i
		method := classTypePtr.Method(i)
		if !method.IsExported() || method.Type.NumIn() < 2 || method.Type.In(1) != reflect.TypeOf(gd.Context{}) || method.Type.NumOut() > 0 {
			continue
		}
		var arguments = make([]gd.PropertyInfo, 0, method.Type.NumIn()-2)
		var metadatas = make([]gd.ClassMethodArgumentMetadata, 0, method.Type.NumIn()-2)
		for i := 2; i < method.Type.NumIn(); i++ {
			arguments = append(arguments, gd.PropertyInfo{
				Type:      variantTypeOf(method.Type.In(i)),
				Name:      godot.StringName("arg" + fmt.Sprint(i)),
				ClassName: godot.StringName(classNameOf(method.Type.In(i))),
			})
			metadatas = append(metadatas, 0)
		}
		var returns *gd.PropertyInfo
		var returnMetadata gd.ClassMethodArgumentMetadata
		if method.Type.NumOut() > 0 {
			returns = &gd.PropertyInfo{
				Type:      variantTypeOf(method.Type.Out(0)),
				Name:      godot.StringName("result"),
				ClassName: godot.StringName(classNameOf(method.Type.Out(0))),
			}
			returnMetadata = 0
		}
		godot.API.ClassDB.RegisterClassMethod(godot, godot.API.ExtensionToken, class, gd.Method{
			Name: godot.StringName(method.Name),
			// FIXME type check and return an error if argumetns are invalid.
			Call: func(instance any, v ...gd.Variant) (gd.Variant, error) {
				ctx := gd.NewContext(godot.API)
				defer ctx.End()

				var args = make([]reflect.Value, len(v)+1)
				args[0] = reflect.ValueOf(ctx)
				for i := 0; i < len(v); i++ {
					args[i+1] = reflect.ValueOf(v[i].Interface(ctx))
				}
				rets := reflect.ValueOf(instance.(*instanceImplementation).Value).Method(i).Call(args)
				if len(rets) > 0 {
					return ctx.Variant(rets[0].Interface()), nil
				}
				return gd.Variant{}, nil
			},
			PointerCall: func(instance any, args gd.UnsafeArgs, ret gd.UnsafeBack) {
				slowCall(reflect.ValueOf(instance.(*instanceImplementation).Value).Method(i), args, ret)
			},
			Arguments:           arguments,
			ArgumentsMetadata:   metadatas,
			ReturnValueInfo:     returns,
			ReturnValueMetadata: returnMetadata,
		})
	}
}

func slowCall(method reflect.Value, p_args gd.UnsafeArgs, p_ret gd.UnsafeBack) {
	var ctx = gd.NewContext(godot)
	defer ctx.End()
	var (
		args = make([]reflect.Value, method.Type().NumIn())
	)
	args[0] = reflect.ValueOf(ctx)
	for i := 1; i < method.Type().NumIn(); i++ {
		switch method.Type().In(i) {
		case reflect.TypeOf(gd.Bool(false)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Bool](p_args, i-1))
		case reflect.TypeOf(gd.Int(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Int](p_args, i-1))
		case reflect.TypeOf(gd.Float(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Float](p_args, i-1))
		case reflect.TypeOf(gd.String{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.String](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.Vector2{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector2](p_args, i-1))
		case reflect.TypeOf(gd.Vector2i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector2i](p_args, i-1))
		case reflect.TypeOf(gd.Rect2{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Rect2](p_args, i-1))
		case reflect.TypeOf(gd.Rect2i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Rect2i](p_args, i-1))
		case reflect.TypeOf(gd.Vector3{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector3](p_args, i-1))
		case reflect.TypeOf(gd.Vector3i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector3i](p_args, i-1))
		case reflect.TypeOf(gd.Transform2D{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Transform2D](p_args, i-1))
		case reflect.TypeOf(gd.Vector4{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector4](p_args, i-1))
		case reflect.TypeOf(gd.Vector4i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Vector4i](p_args, i-1))
		case reflect.TypeOf(gd.Plane{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Plane](p_args, i-1))
		case reflect.TypeOf(gd.Quaternion{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Quaternion](p_args, i-1))
		case reflect.TypeOf(gd.AABB{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.AABB](p_args, i-1))
		case reflect.TypeOf(gd.Basis{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Basis](p_args, i-1))
		case reflect.TypeOf(gd.Transform3D{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Transform3D](p_args, i-1))
		case reflect.TypeOf(gd.Projection{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Projection](p_args, i-1))
		case reflect.TypeOf(gd.Color{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.Color](p_args, i-1))
		case reflect.TypeOf(gd.StringName{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.StringName](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.NodePath{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.NodePath](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.RID(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[gd.RID](p_args, i-1))
		case reflect.TypeOf(gd.Object{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			var obj gd.Object
			obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, ptr))
			args[i] = reflect.ValueOf(obj)
		case reflect.TypeOf(gd.Callable{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.Callable](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.Signal{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.Signal](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.Dictionary{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.Dictionary](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.Array{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedByteArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedByteArray](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedInt32Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedInt32Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedInt64Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedInt64Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedFloat32Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedFloat32Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedFloat64Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedFloat64Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedStringArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedStringArray](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedVector2Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedVector2Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedVector3Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedVector3Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.PackedColorArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.PackedColorArray](ctx.Lifetime, ctx.API, ptr))
		default:
			panic(fmt.Sprintf("gdextension: unsupported Godot -> Go type %v", method.Type().In(i)))
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
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
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
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case gd.NodePath:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case gd.RID:
			gd.UnsafeSet[gd.RID](p_ret, val)
		case gd.Object:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val.AsPointer()))
		case gd.Callable:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.Signal:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.Dictionary:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case gd.Array:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
		case gd.PackedByteArray:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.PackedInt32Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.PackedInt64Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.PackedFloat32Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.PackedFloat64Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.PackedStringArray:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.PackedVector2Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.PackedVector3Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case gd.PackedColorArray:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		default:
			panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", method.Type().Out(0)))
		}
	}
}
