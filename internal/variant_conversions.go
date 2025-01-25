//go:build !generate

package gd

import (
	"fmt"
	"reflect"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	FloatType "graphics.gd/variant/Float"
	NodePathType "graphics.gd/variant/NodePath"
)

// Variant returns a variant from the given value, which must be one of the
// basic godot types defined in the gd package.
func NewVariant(v any) Variant {
	if v == nil {
		return Global.Variants.NewNil()
	}
	var frame = callframe.New()
	var ret = callframe.Ret[VariantPointers](frame)
	rtype := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	switch rtype.Kind() {
	case reflect.Bool:
		var arg = callframe.Arg(frame, value.Bool())
		Global.variant.FromType[TypeBool](ret, arg.Addr())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var arg = callframe.Arg(frame, Int(value.Int()))
		Global.variant.FromType[TypeInt](ret, arg.Addr())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		var arg = callframe.Arg(frame, Int(value.Uint()))
		Global.variant.FromType[TypeInt](ret, arg.Addr())
	case reflect.Uint64:
		var arg = callframe.Arg(frame, RID(value.Uint()))
		Global.variant.FromType[TypeRID](ret, arg.Addr())
	case reflect.Float32, reflect.Float64:
		var arg = callframe.Arg(frame, Float(value.Float()))
		Global.variant.FromType[TypeFloat](ret, arg.Addr())
	case reflect.Complex64, reflect.Complex128:
		var arg = callframe.Arg(frame, Vector2{
			X: FloatType.X(real(value.Complex())),
			Y: FloatType.X(imag(value.Complex())),
		})
		Global.variant.FromType[TypeVector2](ret, arg.Addr())
	case reflect.Pointer:
		if value.IsNil() {
			return Global.Variants.NewNil()
		}
		if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			obj := value.Interface().(IsClass).AsObject()
			if pointers.Get(obj[0]) == ([3]uint64{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(obj[0]))
			Global.variant.FromType[TypeObject](ret, arg.Addr())
		} else {
			return NewVariant(value.Elem().Interface())
		}
	case reflect.Array:
		if rtype.Elem().Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			obj := value.Index(0).Interface().(IsClass).AsObject()
			if pointers.Get(obj[0]) == ([3]uint64{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(obj[0]))
			Global.variant.FromType[TypeObject](ret, arg.Addr())
		} else {
			var arg = callframe.Arg(frame, pointers.Get(newArray(value)))
			Global.variant.FromType[TypeArray](ret, arg.Addr())
		}
	case reflect.Slice:
		switch value := value.Interface().(type) {
		case []byte:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedByteSlice(value)))
			Global.variant.FromType[TypePackedByteArray](ret, arg.Addr())
		case []int32:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedInt32Slice(value)))
			Global.variant.FromType[TypePackedInt32Array](ret, arg.Addr())
		case []int64:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedInt64Slice(value)))
			Global.variant.FromType[TypePackedInt64Array](ret, arg.Addr())
		case []float32:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedFloat32Slice(value)))
			Global.variant.FromType[TypePackedFloat32Array](ret, arg.Addr())
		case []float64:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedFloat64Slice(value)))
			Global.variant.FromType[TypePackedFloat64Array](ret, arg.Addr())
		case []string:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedStringSlice(value)))
			Global.variant.FromType[TypePackedStringArray](ret, arg.Addr())
		case []Vector2:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedVector2Slice(value)))
			Global.variant.FromType[TypePackedVector2Array](ret, arg.Addr())
		case []Vector3:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedVector3Slice(value)))
			Global.variant.FromType[TypePackedVector3Array](ret, arg.Addr())
		case []Color:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedColorSlice(value)))
			Global.variant.FromType[TypePackedColorArray](ret, arg.Addr())
		case []Vector4:
			var arg = callframe.Arg(frame, pointers.Get(NewPackedVector4Slice(value)))
			Global.variant.FromType[TypePackedVector4Array](ret, arg.Addr())
		default:
			var arg = callframe.Arg(frame, pointers.Get(newArray(reflect.ValueOf(value))))
			Global.variant.FromType[TypeArray](ret, arg.Addr())
		}
	case reflect.Func:
		if value.IsNil() {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, NewCallable(value))
		Global.variant.FromType[TypeCallable](ret, arg.Addr())
	case reflect.Map:
		if value.IsNil() {
			return Global.Variants.NewNil()
		}
		var arg = callframe.Arg(frame, pointers.Get(newDictionary(value)))
		Global.variant.FromType[TypeDictionary](ret, arg.Addr())
	case reflect.String:
		switch rtype {
		case reflect.TypeFor[NodePathType.String]():
			var s = NewString(value.String()).NodePath()
			var arg = callframe.Arg(frame, pointers.Get(s))
			Global.variant.FromType[TypeNodePath](ret, arg.Addr())
		default:
			var s = NewString(value.String())
			var arg = callframe.Arg(frame, pointers.Get(s))
			Global.variant.FromType[TypeString](ret, arg.Addr())
		}
	case reflect.Struct:
		switch val := v.(type) {
		case ArrayType.Any:
			var arg = callframe.Arg(frame, pointers.Get(InternalArray(val)))
			Global.variant.FromType[TypeArray](ret, arg.Addr())
		case Variant:
			return val
		case VariantPkg.Any:
			return NewVariant(val.Interface())
		case Vector2:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeVector2](ret, arg.Addr())
		case Vector2i:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeVector2i](ret, arg.Addr())
		case Rect2:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeRect2](ret, arg.Addr())
		case Rect2i:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeRect2i](ret, arg.Addr())
		case Vector3:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeVector3](ret, arg.Addr())
		case Vector3i:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeVector3i](ret, arg.Addr())
		case Transform2D:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeTransform2D](ret, arg.Addr())
		case Vector4:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeVector4](ret, arg.Addr())
		case Vector4i:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeVector4i](ret, arg.Addr())
		case Plane:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypePlane](ret, arg.Addr())
		case Quaternion:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeQuaternion](ret, arg.Addr())
		case AABB:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeAABB](ret, arg.Addr())
		case Basis:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeBasis](ret, arg.Addr())
		case Transform3D:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeTransform3D](ret, arg.Addr())
		case Projection:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeProjection](ret, arg.Addr())
		case Color:
			var arg = callframe.Arg(frame, val)
			Global.variant.FromType[TypeColor](ret, arg.Addr())
		case String:
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypeString](ret, arg.Addr())
		case StringName:
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypeStringName](ret, arg.Addr())
		case NodePath:
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypeStringName](ret, arg.Addr())
		case Object:

			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypeObject](ret, arg.Addr())
		case Callable:
			if pointers.Get(val) == ([2]uint64{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypeCallable](ret, arg.Addr())
		case Signal:
			if pointers.Get(val) == ([2]uint64{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypeSignal](ret, arg.Addr())
		case Dictionary:
			if pointers.Get(val) == ([1]EnginePointer{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypeDictionary](ret, arg.Addr())
		case Array:
			if pointers.Get(val) == ([1]EnginePointer{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypeArray](ret, arg.Addr())
		case PackedByteArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedByteArray](ret, arg.Addr())
		case PackedInt32Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedInt32Array](ret, arg.Addr())
		case PackedInt64Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedInt64Array](ret, arg.Addr())
		case PackedFloat32Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedFloat32Array](ret, arg.Addr())
		case PackedFloat64Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedFloat64Array](ret, arg.Addr())
		case PackedStringArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedStringArray](ret, arg.Addr())
		case PackedVector2Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedVector2Array](ret, arg.Addr())
		case PackedVector3Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedVector3Array](ret, arg.Addr())
		case PackedVector4Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedVector4Array](ret, arg.Addr())
		case PackedColorArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Global.Variants.NewNil()
			}
			var arg = callframe.Arg(frame, pointers.Get(val))
			Global.variant.FromType[TypePackedColorArray](ret, arg.Addr())
		default:
			var arg = callframe.Arg(frame, pointers.Get(newDictionary(value)))
			Global.variant.FromType[TypeDictionary](ret, arg.Addr())
		}

	default:
		panic("gd.Variant: unsupported type " + reflect.TypeOf(v).String())
	}
	var variant = pointers.New[Variant](ret.Get())
	frame.Free()
	return variant
}

func newDictionary(val reflect.Value) Dictionary {
	var dict = NewDictionary()
	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			dict.SetIndex(NewVariant(key.Interface()), NewVariant(val.MapIndex(key).Interface()))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Type().Field(i)
			if !field.IsExported() {
				continue
			}
			name := field.Name
			if tag := field.Tag.Get("gd"); tag != "" {
				name = tag
			}
			dict.SetIndex(NewVariant(name), NewVariant(val.Field(i).Interface()))
		}
	}
	return dict
}

func newArray(val reflect.Value) Array {
	vtype, ok := VariantTypeOf(val.Type().Elem())
	if !ok {
		panic("gd.Variant: unsupported array element type " + val.Type().Elem().String())
	}
	var array = NewArray()
	Global.Array.SetTyped(array, vtype, StringName{}, Object{})
	array.Resize(Int(val.Len()))
	for i := 0; i < val.Len(); i++ {
		array.SetIndex(Int(i), NewVariant(val.Index(i).Interface()))
	}
	return array
}

// Interface returns the variant's value as one of the the native Godot values
// (as defined) in the gd package.
func (variant Variant) Interface() any {
	switch vtype := variant.Type(); vtype {
	case TypeNil:
		return nil
	case TypeBool:
		return Bool(variantAsValueType[bool](variant, vtype))
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
		array := variantAsPointerType[Array](variant, vtype)
		return ArrayType.Through(ArrayProxy[VariantPkg.Any]{}, pointers.Pack(array))
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
	Global.variant.IntoType[vtype](r_ret.Addr(), callframe.Arg(frame, pointers.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

func variantAsPointerType[T pointers.Generic[T, Size], Size pointers.Size](variant Variant, vtype VariantType) T {
	var frame = callframe.New()
	var r_ret = callframe.Ret[Size](frame)
	Global.variant.IntoType[vtype](r_ret.Addr(), callframe.Arg(frame, pointers.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return pointers.New[T](ret)
}

func LetVariantAsPointerType[T pointers.Generic[T, Size], Size pointers.Size](variant Variant, vtype VariantType) T {
	var frame = callframe.New()
	var r_ret = callframe.Ret[Size](frame)
	Global.variant.IntoType[vtype](r_ret.Addr(), callframe.Arg(frame, pointers.Get(variant)))
	var ret = r_ret.Get()
	frame.Free()
	return pointers.Let[T](ret)
}

var ObjectAs = func(name string, ptr Object) any {
	return ptr
}
