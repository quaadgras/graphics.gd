//go:build !generate

package gd

import (
	"fmt"
	"reflect"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
)

// Variant returns a variant from the given value, which must be one of the
// basic godot types defined in the gd package.
func NewVariant(v any) Variant {
	if v == nil {
		return Global.Variants.NewNil()
	}
	var frame = callframe.New()
	var ret = callframe.Ret[[3]uintptr](frame)
	switch val := v.(type) {
	case Variant:
		frame.Free()
		return Global.Variants.NewCopy(val)
	case Bool:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeBool](ret, arg.Uintptr())
	case Int:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeInt](ret, arg.Uintptr())
	case Float:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeFloat](ret, arg.Uintptr())
	case String:
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypeString](ret, arg.Uintptr())
	case Vector2:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeVector2](ret, arg.Uintptr())
	case Vector2i:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeVector2i](ret, arg.Uintptr())
	case Rect2:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeRect2](ret, arg.Uintptr())
	case Rect2i:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeRect2i](ret, arg.Uintptr())
	case Vector3:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeVector3](ret, arg.Uintptr())
	case Vector3i:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeVector3i](ret, arg.Uintptr())
	case Transform2D:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeTransform2D](ret, arg.Uintptr())
	case Vector4:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeVector4](ret, arg.Uintptr())
	case Vector4i:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeVector4i](ret, arg.Uintptr())
	case Plane:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypePlane](ret, arg.Uintptr())
	case Quaternion:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeQuaternion](ret, arg.Uintptr())
	case AABB:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeAABB](ret, arg.Uintptr())
	case Basis:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeBasis](ret, arg.Uintptr())
	case Transform3D:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeTransform3D](ret, arg.Uintptr())
	case Projection:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeProjection](ret, arg.Uintptr())
	case Color:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeColor](ret, arg.Uintptr())
	case StringName:
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypeStringName](ret, arg.Uintptr())
	case NodePath:
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypeStringName](ret, arg.Uintptr())
	case RID:
		var arg = callframe.Arg(frame, val)
		Global.variant.FromType[TypeRID](ret, arg.Uintptr())
	case Object:
		if pointers.Get(val) == ([3]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypeObject](ret, arg.Uintptr())
	case Callable:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypeCallable](ret, arg.Uintptr())
	case Signal:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypeSignal](ret, arg.Uintptr())
	case Dictionary:
		if pointers.Get(val) == ([1]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypeDictionary](ret, arg.Uintptr())
	case Array:
		if pointers.Get(val) == ([1]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypeArray](ret, arg.Uintptr())
	case PackedByteArray:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedByteArray](ret, arg.Uintptr())
	case PackedInt32Array:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedInt32Array](ret, arg.Uintptr())
	case PackedInt64Array:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedInt64Array](ret, arg.Uintptr())
	case PackedFloat32Array:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedFloat32Array](ret, arg.Uintptr())
	case PackedFloat64Array:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedFloat64Array](ret, arg.Uintptr())
	case PackedStringArray:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedStringArray](ret, arg.Uintptr())
	case PackedVector2Array:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedVector2Array](ret, arg.Uintptr())
	case PackedVector3Array:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedVector3Array](ret, arg.Uintptr())
	case PackedVector4Array:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedVector4Array](ret, arg.Uintptr())
	case PackedColorArray:
		if pointers.Get(val) == ([2]uintptr{}) {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(val))
		Global.variant.FromType[TypePackedColorArray](ret, arg.Uintptr())
	case interface{ Variant() Variant }:
		return val.Variant()
	default:
		class, ok := v.(IsClass)
		if ok {
			if reflect.ValueOf(v).IsZero() {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(class.AsObject()))
			Global.variant.FromType[TypeObject](ret, arg.Uintptr())
		} else {
			rtype := reflect.TypeOf(v)
			value := reflect.ValueOf(v)
			switch rtype.Kind() {
			case reflect.String:
				var s = NewString(value.String())
				var arg = callframe.Arg(frame, pointers.Get(s))
				Global.variant.FromType[TypeString](ret, arg.Uintptr())
			case reflect.Map:
				var dict = NewDictionary()
				for _, key := range value.MapKeys() {
					var k = NewVariant(key.Interface())
					var v = NewVariant(value.MapIndex(key).Interface())
					dict.SetIndex(k, v)
				}
				var arg = callframe.Arg(frame, pointers.Get(dict))
				Global.variant.FromType[TypeDictionary](ret, arg.Uintptr())
			default:
				panic("gd.Variant: unsupported type " + reflect.TypeOf(v).String())
			}
		}
	}
	var variant = pointers.New[Variant](ret.Get())
	frame.Free()
	return variant
}

// Interface returns the variant's value as one of the the native Godot values
// (as defined) in the gd package.
func (variant Variant) Interface() any {
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
		return variantAsPointerType[String](variant, vtype)
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
	case TypeTransform2D:
		return variantAsValueType[Transform2D](variant, vtype)
	case TypeVector4:
		return variantAsValueType[Vector4](variant, vtype)
	case TypeVector4i:
		return variantAsValueType[Vector4i](variant, vtype)
	case TypePlane:
		return variantAsValueType[Plane](variant, vtype)
	case TypeQuaternion:
		return variantAsValueType[Quaternion](variant, vtype)
	case TypeAABB:
		return variantAsValueType[AABB](variant, vtype)
	case TypeBasis:
		return variantAsValueType[Basis](variant, vtype)
	case TypeTransform3D:
		return variantAsValueType[Transform3D](variant, vtype)
	case TypeProjection:
		return variantAsValueType[Projection](variant, vtype)
	case TypeColor:
		return variantAsValueType[Color](variant, vtype)
	case TypeStringName:
		return variantAsPointerType[StringName](variant, vtype)
	case TypeNodePath:
		return variantAsPointerType[NodePath](variant, vtype)
	case TypeRID:
		return variantAsValueType[RID](variant, vtype)
	case TypeObject:
		var obj = variantAsPointerType[Object](variant, vtype)
		return ObjectAs(obj.GetClass().String(), obj)
	case TypeCallable:
		return variantAsPointerType[Callable](variant, vtype)
	case TypeSignal:
		return variantAsPointerType[Signal](variant, vtype)
	case TypeDictionary:
		return variantAsPointerType[Dictionary](variant, vtype)
	case TypeArray:
		return variantAsPointerType[Array](variant, vtype)
	case TypePackedByteArray:
		return variantAsPointerType[PackedByteArray](variant, vtype)
	case TypePackedInt32Array:
		return variantAsPointerType[PackedInt32Array](variant, vtype)
	case TypePackedInt64Array:
		return variantAsPointerType[PackedInt64Array](variant, vtype)
	case TypePackedFloat32Array:
		return variantAsPointerType[PackedFloat32Array](variant, vtype)
	case TypePackedFloat64Array:
		return variantAsPointerType[PackedFloat64Array](variant, vtype)
	case TypePackedStringArray:
		return variantAsPointerType[PackedStringArray](variant, vtype)
	case TypePackedVector2Array:
		return variantAsPointerType[PackedVector2Array](variant, vtype)
	case TypePackedVector3Array:
		return variantAsPointerType[PackedVector3Array](variant, vtype)
	case TypePackedVector4Array:
		return variantAsPointerType[PackedVector4Array](variant, vtype)
	case TypePackedColorArray:
		return variantAsPointerType[PackedColorArray](variant, vtype)
	default:
		panic("gd.Variant.Interface: invalid variant type " + fmt.Sprint(vtype))
	}
}

func variantAsValueType[T comparable](variant Variant, vtype VariantType) T {
	var frame = callframe.New()
	var r_ret = callframe.Ret[T](frame)
	Global.variant.IntoType[vtype](r_ret.Uintptr(), callframe.Arg(frame, pointers.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

func variantAsPointerType[T pointers.Pointer[T, Size, R], Size pointers.Shape, R pointers.PointerType](variant Variant, vtype VariantType) T {
	var frame = callframe.New()
	var r_ret = callframe.Ret[Size](frame)
	Global.variant.IntoType[vtype](r_ret.Uintptr(), callframe.Arg(frame, pointers.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return pointers.New[T](ret)
}

func LetVariantAsPointerType[T pointers.Pointer[T, Size, R], Size pointers.Shape, R pointers.PointerType](variant Variant, vtype VariantType) T {
	var frame = callframe.New()
	var r_ret = callframe.Ret[Size](frame)
	Global.variant.IntoType[vtype](r_ret.Uintptr(), callframe.Arg(frame, pointers.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return pointers.New[T](ret)
}

var ObjectAs = func(name string, ptr Object) any {
	return ptr
}
