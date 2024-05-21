//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"unsafe"

	"grow.graphics/gd/internal/callframe"

	"runtime.link/mmm"
)

// Variant returns a variant from the given value, which must be one of the
// basic godot types defined in the gd package.
func (godot Context) Variant(v any) Variant {
	if v == nil {
		return godot.API.Variants.NewNil(godot)
	}
	var frame = callframe.New()
	var ret = callframe.Ret[[3]uintptr](frame)
	switch val := v.(type) {
	case Variant:
		frame.Free()
		return godot.API.Variants.NewCopy(godot, val)
	case bool:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeBool](ret, arg.Uintptr())
	case Int:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeInt](ret, arg.Uintptr())
	case Float:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeFloat](ret, arg.Uintptr())
	case String:
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypeString](ret, arg.Uintptr())
	case Vector2:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeVector2](ret, arg.Uintptr())
	case Vector2i:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeVector2i](ret, arg.Uintptr())
	case Rect2:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeRect2](ret, arg.Uintptr())
	case Rect2i:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeRect2i](ret, arg.Uintptr())
	case Vector3:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeVector3](ret, arg.Uintptr())
	case Vector3i:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeVector3i](ret, arg.Uintptr())
	case Transform2D:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeTransform2d](ret, arg.Uintptr())
	case Vector4:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeVector4](ret, arg.Uintptr())
	case Vector4i:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeVector4i](ret, arg.Uintptr())
	case Plane:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypePlane](ret, arg.Uintptr())
	case Quaternion:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeQuaternion](ret, arg.Uintptr())
	case AABB:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeAabb](ret, arg.Uintptr())
	case Basis:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeBasis](ret, arg.Uintptr())
	case Transform3D:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeTransform3d](ret, arg.Uintptr())
	case Projection:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeProjection](ret, arg.Uintptr())
	case Color:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeColor](ret, arg.Uintptr())
	case StringName:
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypeStringName](ret, arg.Uintptr())
	case NodePath:
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypeStringName](ret, arg.Uintptr())
	case RID:
		var arg = callframe.Arg(frame, val)
		godot.API.variant.FromType[TypeRid](ret, arg.Uintptr())
	case Object:
		if mmm.Get(val.AsPointer()) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val.AsPointer()))
		godot.API.variant.FromType[TypeObject](ret, arg.Uintptr())
	case Callable:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypeCallable](ret, arg.Uintptr())
	case Signal:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypeSignal](ret, arg.Uintptr())
	case Dictionary:
		if mmm.Get(val) == 0 {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypeDictionary](ret, arg.Uintptr())
	case Array:
		if mmm.Get(val) == 0 {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypeArray](ret, arg.Uintptr())
	case PackedByteArray:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedByteArray](ret, arg.Uintptr())
	case PackedInt32Array:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedInt32Array](ret, arg.Uintptr())
	case PackedInt64Array:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedInt64Array](ret, arg.Uintptr())
	case PackedFloat32Array:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedFloat32Array](ret, arg.Uintptr())
	case PackedFloat64Array:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedFloat64Array](ret, arg.Uintptr())
	case PackedStringArray:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedStringArray](ret, arg.Uintptr())
	case PackedVector2Array:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedVector2Array](ret, arg.Uintptr())
	case PackedVector3Array:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedVector3Array](ret, arg.Uintptr())
	case PackedColorArray:
		if mmm.Get(val) == ([2]uintptr{}) {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypePackedColorArray](ret, arg.Uintptr())
	case IsArray:
		if mmm.Get(val) == 0 {
			return godot.API.Variants.NewNil(godot)
		}
		var arg = callframe.Arg(frame, mmm.Get(val))
		godot.API.variant.FromType[TypeArray](ret, arg.Uintptr())
	default:
		class, ok := v.(IsClass)
		if ok {
			if reflect.ValueOf(v).IsZero() {
				return godot.API.Variants.NewNil(godot)
			}
			var arg = callframe.Arg(frame, class.AsPointer().Pointer())
			godot.API.variant.FromType[TypeObject](ret, arg.Uintptr())
		} else {
			panic("gd.Variant: unsupported type " + reflect.TypeOf(v).String())
		}
	}
	var variant = mmm.New[Variant](godot.Lifetime, godot.API, ret.Get())
	frame.Free()
	return variant
}

// Interface returns the variant's value as one of the the native Godot values
// (as defined) in the gd package.
func (variant Variant) Interface(ctx Context) any {
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

		API := ctx.API
		ref := API.Object.CastTo(obj, API.refCountedClassTag)
		if mmm.Get(ref.AsPointer()) != ([2]uintptr{}) {
			(*(*RefCounted)(unsafe.Pointer(&ref))).Reference()
		}
		tmp := NewContext(API)
		defer tmp.End()
		return ObjectAs(obj.GetClass(tmp).String(), obj.AsPointer())
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
		panic("gd.Variant.Interface: invalid variant type " + fmt.Sprint(vtype))
	}
}

func variantAsValueType[T comparable](variant Variant, vtype VariantType) T {
	var frame = callframe.New()
	var r_ret = callframe.Ret[T](frame)
	mmm.API(variant).variant.IntoType[vtype](r_ret.Uintptr(), callframe.Arg(frame, mmm.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

func variantAsPointerType[T mmm.PointerWithFree[API, T, Size], Size mmm.PointerSize](ctx Context, variant Variant, vtype VariantType) T {
	var frame = callframe.New()
	var r_ret = callframe.Ret[Size](frame)
	mmm.API(variant).variant.IntoType[vtype](r_ret.Uintptr(), callframe.Arg(frame, mmm.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return mmm.New[T](ctx.Lifetime, ctx.API, ret)
}

func LetVariantAsPointerType[T mmm.PointerWithFree[API, T, Size], Size mmm.PointerSize](ctx Context, variant Variant, vtype VariantType) T {
	var frame = callframe.New()
	var r_ret = callframe.Ret[Size](frame)
	mmm.API(variant).variant.IntoType[vtype](r_ret.Uintptr(), callframe.Arg(frame, mmm.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return mmm.Let[T](ctx.Lifetime, ctx.API, ret)
}

var ObjectAs = func(name string, ptr Pointer) any {
	var obj Object
	obj.SetPointer(ptr)
	return obj
}
