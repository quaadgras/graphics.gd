//go:build !generate

package gd

import (
	"reflect"

	"runtime.link/api/call"
	"runtime.link/mmm"
)

// Variant returns a variant from the given value, which must be one of the
// basic godot types defined in the gd package.
func (ctx Context) Variant(v any) Variant {
	var godot = ctx.API()
	if v == nil {
		return godot.Variants.NewNil(ctx)
	}
	var frame = call.New()
	var ret = call.Ret[[3]uintptr](frame)
	switch val := v.(type) {
	case Variant:
		frame.Free()
		return godot.Variants.NewCopy(ctx, val)
	case bool:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeBool](ret, arg)
	case Int:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeInt](ret, arg)
	case Float:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeFloat](ret, arg)
	case String:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypeString](ret, arg)
	case Vector2:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeVector2](ret, arg)
	case Vector2i:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeVector2i](ret, arg)
	case Rect2:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeRect2](ret, arg)
	case Rect2i:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeRect2i](ret, arg)
	case Vector3:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeVector3](ret, arg)
	case Vector3i:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeVector3i](ret, arg)
	case Transform2D:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeTransform2d](ret, arg)
	case Vector4:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeVector4](ret, arg)
	case Vector4i:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeVector4i](ret, arg)
	case Plane:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypePlane](ret, arg)
	case Quaternion:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeQuaternion](ret, arg)
	case AABB:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeAabb](ret, arg)
	case Basis:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeBasis](ret, arg)
	case Transform3D:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeTransform3d](ret, arg)
	case Projection:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeProjection](ret, arg)
	case Color:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeColor](ret, arg)
	case StringName:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypeStringName](ret, arg)
	case NodePath:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypeStringName](ret, arg)
	case RID:
		var arg = call.Arg(frame, val)
		godot.variant.FromType[TypeRid](ret, arg)
	case Object:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypeObject](ret, arg)
	case Callable:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypeCallable](ret, arg)
	case Signal:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypeSignal](ret, arg)
	case Dictionary:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypeDictionary](ret, arg)
	case Array:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypeArray](ret, arg)
	case PackedByteArray:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedByteArray](ret, arg)
	case PackedInt32Array:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedInt32Array](ret, arg)
	case PackedInt64Array:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedInt64Array](ret, arg)
	case PackedFloat32Array:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedFloat32Array](ret, arg)
	case PackedFloat64Array:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedFloat64Array](ret, arg)
	case PackedStringArray:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedStringArray](ret, arg)
	case PackedVector2Array:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedVector2Array](ret, arg)
	case PackedVector3Array:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedVector3Array](ret, arg)
	case PackedColorArray:
		var arg = call.Arg(frame, val.Pointer())
		godot.variant.FromType[TypePackedColorArray](ret, arg)
	default:
		class, ok := v.(IsClass)
		if ok {
			var arg = call.Arg(frame, class.AsPointer().Pointer())
			godot.variant.FromType[TypeObject](ret, arg)
		} else {
			panic("gd.Variant: unsupported type " + reflect.TypeOf(v).String())
		}
	}
	var variant = mmm.Make[API, Variant](ctx, godot, ret.Get())
	frame.Free()
	return variant
}

// Interface returns the variant's value as one of the the native Godot values
// (as defined) in the gd package.
func (variant Variant) Interface(ctx mmm.Context) any {
	switch vtype := variant.Type(); vtype {
	case TypeNil:
		return nil
	case TypeBool:
		return variantAsValueType[bool](variant, vtype)
	case TypeInt:
		return variantAsValueType[Int](variant, vtype)
	case TypeFloat:
		return variantAsValueType[Float](variant, vtype)
	case TypeString:
		return variantAsPointerType[String](ctx, variant, vtype)
	case TypeVector2:
		return variantAsValueType[Vector2](variant, vtype)
	case TypeVector2i:
		return variantAsValueType[Vector2i](variant, vtype)
	case TypeRect2:
		return variantAsValueType[Rect2](variant, vtype)
	case TypeRect2i:
		return variantAsValueType[Rect2i](variant, vtype)
	case TypeVector3:
		return variantAsValueType[Vector3](variant, vtype)
	case TypeVector3i:
		return variantAsValueType[Vector3i](variant, vtype)
	case TypeTransform2d:
		return variantAsValueType[Transform2D](variant, vtype)
	case TypeVector4:
		return variantAsValueType[Vector4](variant, vtype)
	case TypeVector4i:
		return variantAsValueType[Vector4i](variant, vtype)
	case TypePlane:
		return variantAsValueType[Plane](variant, vtype)
	case TypeQuaternion:
		return variantAsValueType[Quaternion](variant, vtype)
	case TypeAabb:
		return variantAsValueType[AABB](variant, vtype)
	case TypeBasis:
		return variantAsValueType[Basis](variant, vtype)
	case TypeTransform3d:
		return variantAsValueType[Transform3D](variant, vtype)
	case TypeProjection:
		return variantAsValueType[Projection](variant, vtype)
	case TypeColor:
		return variantAsValueType[Color](variant, vtype)
	case TypeStringName:
		return variantAsPointerType[StringName](ctx, variant, vtype)
	case TypeNodePath:
		return variantAsPointerType[NodePath](ctx, variant, vtype)
	case TypeRid:
		return variantAsValueType[RID](variant, vtype)
	case TypeObject:
		var obj Object
		obj.SetPointer(variantAsPointerType[Pointer](ctx, variant, vtype))
		return obj
	case TypeCallable:
		return variantAsPointerType[Callable](ctx, variant, vtype)
	case TypeSignal:
		return variantAsPointerType[Signal](ctx, variant, vtype)
	case TypeDictionary:
		return variantAsPointerType[Dictionary](ctx, variant, vtype)
	case TypeArray:
		return variantAsPointerType[Array](ctx, variant, vtype)
	case TypePackedByteArray:
		return variantAsPointerType[PackedByteArray](ctx, variant, vtype)
	case TypePackedInt32Array:
		return variantAsPointerType[PackedInt32Array](ctx, variant, vtype)
	case TypePackedInt64Array:
		return variantAsPointerType[PackedInt64Array](ctx, variant, vtype)
	case TypePackedFloat32Array:
		return variantAsPointerType[PackedFloat32Array](ctx, variant, vtype)
	case TypePackedFloat64Array:
		return variantAsPointerType[PackedFloat64Array](ctx, variant, vtype)
	case TypePackedStringArray:
		return variantAsPointerType[PackedStringArray](ctx, variant, vtype)
	case TypePackedVector2Array:
		return variantAsPointerType[PackedVector2Array](ctx, variant, vtype)
	case TypePackedVector3Array:
		return variantAsPointerType[PackedVector3Array](ctx, variant, vtype)
	case TypePackedColorArray:
		return variantAsPointerType[PackedColorArray](ctx, variant, vtype)
	default:
		panic("gd.Variant.Interface: invalid variant type")
	}
}

func variantAsValueType[T comparable](variant Variant, vtype VariantType) T {
	var godot = variant.API
	var frame = call.New()
	var r_ret = call.Ret[T](frame)
	godot.variant.IntoType[vtype](r_ret, call.Arg(frame, variant.Pointer()))
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

func variantAsPointerType[T mmm.IsPointerAlias[API, Kind], Kind comparable](ctx mmm.Context, variant Variant, vtype VariantType) T {
	var godot = variant.API
	var frame = call.New()
	var r_ret = call.Ret[Kind](frame)
	godot.variant.IntoType[vtype](r_ret, call.Arg(frame, variant.Pointer()))
	var ret = r_ret.Get()
	frame.Free()
	return mmm.Make[API, T, Kind](ctx, godot, ret)
}
