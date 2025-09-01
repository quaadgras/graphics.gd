//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	BasisType "graphics.gd/variant/Basis"
	CallableType "graphics.gd/variant/Callable"
	DictionaryType "graphics.gd/variant/Dictionary"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Euler"
	FloatType "graphics.gd/variant/Float"
	PackedType "graphics.gd/variant/Packed"
	"graphics.gd/variant/Path"
	SignalType "graphics.gd/variant/Signal"
	StringType "graphics.gd/variant/String"
)

// Variant returns a variant from the given value, which must be one of the
// basic godot types defined in the gd package.
func NewVariant(v any) Variant {
	return CutVariant(v, false)
}

// CutVariant is like NewVariant but when cut is true, releases the ownership
// of the given value. Use it on return values passed back to the engine.
//
// used to fix cases of https://github.com/quaadgras/graphics.gd/issues/147
func CutVariant(v any, cut bool) Variant {
	if v == nil {
		return Variant{}
	}
	var ret gdextension.Variant
	if enum, ok := v.(Enum.Any); ok {
		v = enum.Int()
	}
	rtype := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	switch rtype.Kind() {
	case reflect.Bool:
		var arg = value.Bool()
		ret.LoadNative(TypeBool, gdextension.SizeBool, unsafe.Pointer(&arg))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var arg = Int(value.Int())
		ret.LoadNative(TypeInt, gdextension.SizeInt, unsafe.Pointer(&arg))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		var arg = Int(value.Uint())
		ret.LoadNative(TypeInt, gdextension.SizeInt, unsafe.Pointer(&arg))
	case reflect.Uint64:
		if instance := value.MethodByName("Instance"); instance.IsValid() && instance.Type().NumOut() == 2 && instance.Type().NumIn() == 0 {
			result := instance.Call(nil)
			if !result[1].Bool() {
				return Variant{}
			}
			obj := result[0].Interface().(IsClass).AsObject()
			if pointers.Get(obj[0]) == ([3]uint64{}) {
				return Variant{}
			}
			var arg = gdextension.Object(pointers.Cut(obj[0], cut)[0])
			ret.LoadNative(TypeObject, gdextension.SizeObject, unsafe.Pointer(&arg))
		} else {
			var arg = RID(value.Uint())
			ret.LoadNative(TypeRID, gdextension.SizeRID, unsafe.Pointer(&arg))
		}
	case reflect.Float32, reflect.Float64:
		var arg = Float(value.Float())
		ret.LoadNative(TypeFloat, gdextension.SizeFloat, unsafe.Pointer(&arg))
	case reflect.Complex64, reflect.Complex128:
		var arg = Vector2{
			X: FloatType.X(real(value.Complex())),
			Y: FloatType.X(imag(value.Complex())),
		}
		ret.LoadNative(TypeVector2, gdextension.SizeVector2, unsafe.Pointer(&arg))
	case reflect.Pointer:
		if value.IsNil() {
			return Variant{}
		}
		if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			obj := value.Interface().(IsClass).AsObject()
			if pointers.Get(obj[0]) == ([3]uint64{}) {
				return Variant{}
			}
			var arg = gdextension.Object(pointers.Get(obj[0])[0])
			ret.LoadNative(TypeObject, gdextension.SizeObject, unsafe.Pointer(&arg))
		} else {
			return NewVariant(value.Elem().Interface())
		}
	case reflect.Array:
		if rtype.Elem().Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			obj := value.Index(0).Interface().(IsClass).AsObject()
			if pointers.Get(obj[0]) == ([3]uint64{}) {
				return Variant{}
			}
			var arg = gdextension.Object(pointers.Cut(obj[0], cut)[0])
			ret.LoadNative(TypeObject, gdextension.SizeObject, unsafe.Pointer(&arg))
		} else {
			var arg = pointers.Cut(newArray(value), cut)
			ret.LoadNative(TypeArray, gdextension.SizeArray, unsafe.Pointer(&arg))
		}
	case reflect.Slice:
		switch value := value.Interface().(type) {
		case []byte:
			var arg = pointers.Cut(NewPackedByteSlice(value), cut)
			ret.LoadNative(TypePackedByteArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []int32:
			var arg = pointers.Cut(NewPackedInt32Slice(value), cut)
			ret.LoadNative(TypePackedInt32Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []int64:
			var arg = pointers.Cut(NewPackedInt64Slice(value), cut)
			ret.LoadNative(TypePackedInt64Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []float32:
			var arg = pointers.Cut(NewPackedFloat32Slice(value), cut)
			ret.LoadNative(TypePackedFloat32Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []float64:
			var arg = pointers.Cut(NewPackedFloat64Slice(value), cut)
			ret.LoadNative(TypePackedFloat64Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []string:
			var arg = pointers.Cut(NewPackedStringSlice(value), cut)
			ret.LoadNative(TypePackedStringArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []StringType.Readable:
			var arg = pointers.Cut(NewPackedReadableStringSlice(value), cut)
			ret.LoadNative(TypePackedStringArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []Vector2:
			var arg = pointers.Cut(NewPackedVector2Slice(value), cut)
			ret.LoadNative(TypePackedVector2Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []Vector3:
			var arg = pointers.Cut(NewPackedVector3Slice(value), cut)
			ret.LoadNative(TypePackedVector3Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []Color:
			var arg = pointers.Cut(NewPackedColorSlice(value), cut)
			ret.LoadNative(TypePackedColorArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case []Vector4:
			var arg = pointers.Cut(NewPackedVector4Slice(value), cut)
			ret.LoadNative(TypePackedVector4Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		default:
			var arg = pointers.Cut(newArray(reflect.ValueOf(value)), cut)
			ret.LoadNative(TypeArray, gdextension.SizeArray, unsafe.Pointer(&arg))
		}
	case reflect.Func:
		if value.IsNil() {
			return Variant{}
		}
		var arg = gdextension.Callable(pointers.Get(NewCallable(value)))
		ret.LoadNative(TypeCallable, gdextension.SizeCallable, unsafe.Pointer(&arg))
	case reflect.Map:
		if value.IsNil() {
			return Variant{}
		}
		var arg = pointers.Cut(newDictionary(value), cut)
		ret.LoadNative(TypeDictionary, gdextension.SizeDictionary, unsafe.Pointer(&arg))
	case reflect.String:
		switch rtype {
		case reflect.TypeFor[Path.ToNode]():
			var s = NewString(value.String()).NodePath()
			var arg = pointers.Cut(s, true)
			ret.LoadNative(TypeNodePath, gdextension.SizeNodePath, unsafe.Pointer(&arg))
		case reflect.TypeFor[StringType.Name]():
			var s = NewStringName(value.String())
			var arg = pointers.Cut(s, true)
			ret.LoadNative(TypeStringName, gdextension.SizeStringName, unsafe.Pointer(&arg))
		case reflect.TypeFor[StringType.Readable]():
			var s = NewString(value.String())
			var arg = pointers.Cut(s, true)
			ret.LoadNative(TypeString, gdextension.SizeString, unsafe.Pointer(&arg))
		default:
			var s = NewString(value.String())
			var arg = pointers.Cut(s, cut)
			ret.LoadNative(TypeString, gdextension.SizeString, unsafe.Pointer(&arg))
		}
	case reflect.Struct:
		switch val := v.(type) {
		case PackedType.Bytes:
			var arg = pointers.Cut(InternalPacked[PackedByteArray, byte](PackedType.Array[byte](val)), cut)
			ret.LoadNative(TypePackedByteArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[byte]:
			var arg = pointers.Cut(InternalPacked[PackedByteArray, byte](val), cut)
			ret.LoadNative(TypePackedByteArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[int32]:
			var arg = pointers.Cut(InternalPacked[PackedInt32Array, int32](val), cut)
			ret.LoadNative(TypePackedInt32Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[int64]:
			var arg = pointers.Cut(InternalPacked[PackedInt64Array, int64](val), cut)
			ret.LoadNative(TypePackedInt64Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[float32]:
			var arg = pointers.Cut(InternalPacked[PackedFloat32Array, float32](val), cut)
			ret.LoadNative(TypePackedFloat32Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[float64]:
			var arg = pointers.Cut(InternalPacked[PackedFloat64Array, float64](val), cut)
			ret.LoadNative(TypePackedFloat64Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Strings:
			var arg = pointers.Cut(InternalPackedStrings(val), cut)
			ret.LoadNative(TypePackedStringArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[Vector2]:
			var arg = pointers.Cut(InternalPacked[PackedVector2Array, Vector2](val), cut)
			ret.LoadNative(TypePackedVector2Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[Vector3]:
			var arg = pointers.Cut(InternalPacked[PackedVector3Array, Vector3](val), cut)
			ret.LoadNative(TypePackedVector3Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[Color]:
			var arg = pointers.Cut(InternalPacked[PackedColorArray, Color](val), cut)
			ret.LoadNative(TypePackedColorArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedType.Array[Vector4]:
			var arg = pointers.Cut(InternalPacked[PackedVector4Array, Vector4](val), cut)
			ret.LoadNative(TypePackedVector4Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case ArrayType.Any:
			var arg = pointers.Cut(InternalArray(val), cut)
			ret.LoadNative(TypeArray, gdextension.SizeArray, unsafe.Pointer(&arg))
		case ArrayType.Interface:
			var arg = pointers.Cut(InternalArray(val.Any()), cut)
			ret.LoadNative(TypeArray, gdextension.SizeArray, unsafe.Pointer(&arg))
		case DictionaryType.Any:
			var arg = pointers.Cut(InternalDictionary(val), cut)
			ret.LoadNative(TypeDictionary, gdextension.SizeDictionary, unsafe.Pointer(&arg))
		case DictionaryType.Interface:
			var arg = pointers.Cut(InternalDictionary(val.Any()), cut)
			ret.LoadNative(TypeDictionary, gdextension.SizeDictionary, unsafe.Pointer(&arg))
		case StringType.Readable:
			var arg = pointers.Cut(InternalString(val), cut)
			ret.LoadNative(TypeString, gdextension.SizeString, unsafe.Pointer(&arg))
		case Path.ToNode:
			var arg = pointers.Cut(InternalNodePath(val), cut)
			ret.LoadNative(TypeNodePath, gdextension.SizeNodePath, unsafe.Pointer(&arg))
		case StringType.Name:
			var arg = pointers.Cut(InternalStringName(val), cut)
			ret.LoadNative(TypeStringName, gdextension.SizeStringName, unsafe.Pointer(&arg))
		case CallableType.Function:
			var arg = gdextension.Callable(pointers.Cut(InternalCallable(val), cut))
			ret.LoadNative(TypeCallable, gdextension.SizeCallable, unsafe.Pointer(&arg))
		case SignalType.Any:
			var arg = gdextension.Signal(pointers.Cut(InternalSignal(val), cut))
			ret.LoadNative(TypeSignal, gdextension.SizeSignal, unsafe.Pointer(&arg))
		case Variant:
			return val
		case VariantPkg.Any:
			return CutVariant(val.Interface(), cut)
		case Vector2:
			var arg = val
			ret.LoadNative(TypeVector2, gdextension.SizeVector2, unsafe.Pointer(&arg))
		case Vector2i:
			var arg = val
			ret.LoadNative(TypeVector2i, gdextension.SizeVector2i, unsafe.Pointer(&arg))
		case Rect2:
			var arg = val
			ret.LoadNative(TypeRect2, gdextension.SizeRect2, unsafe.Pointer(&arg))
		case Rect2i:
			var arg = val
			ret.LoadNative(TypeRect2i, gdextension.SizeRect2i, unsafe.Pointer(&arg))
		case Vector3:
			var arg = val
			ret.LoadNative(TypeVector3, gdextension.SizeVector3, unsafe.Pointer(&arg))
		case Euler.Radians:
			var arg = val
			ret.LoadNative(TypeVector3, gdextension.SizeVector3, unsafe.Pointer(&arg))
		case Euler.Degrees:
			var arg = val
			ret.LoadNative(TypeVector3, gdextension.SizeVector3, unsafe.Pointer(&arg))
		case Vector3i:
			var arg = val
			ret.LoadNative(TypeVector3i, gdextension.SizeVector3i, unsafe.Pointer(&arg))
		case Transform2D:
			var arg = val
			ret.LoadNative(TypeTransform2D, gdextension.SizeTransform2D, unsafe.Pointer(&arg))
		case Vector4:
			var arg = val
			ret.LoadNative(TypeVector4, gdextension.SizeVector4, unsafe.Pointer(&arg))
		case Vector4i:
			var arg = val
			ret.LoadNative(TypeVector4i, gdextension.SizeVector4i, unsafe.Pointer(&arg))
		case Plane:
			var arg = val
			ret.LoadNative(TypePlane, gdextension.SizePlane, unsafe.Pointer(&arg))
		case Quaternion:
			var arg = val
			ret.LoadNative(TypeQuaternion, gdextension.SizeQuaternion, unsafe.Pointer(&arg))
		case AABB:
			var arg = val
			ret.LoadNative(TypeAABB, gdextension.SizeAABB, unsafe.Pointer(&arg))
		case Basis:
			var arg = BasisType.Transposed(val)
			ret.LoadNative(TypeBasis, gdextension.SizeBasis, unsafe.Pointer(&arg))
		case Transform3D:
			var arg = Transposed(val)
			ret.LoadNative(TypeTransform3D, gdextension.SizeTransform3D, unsafe.Pointer(&arg))
		case Projection:
			var arg = val
			ret.LoadNative(TypeProjection, gdextension.SizeProjection, unsafe.Pointer(&arg))
		case Color:
			var arg = val
			ret.LoadNative(TypeColor, gdextension.SizeColor, unsafe.Pointer(&arg))
		case String:
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypeString, gdextension.SizeString, unsafe.Pointer(&arg))
		case StringName:
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypeStringName, gdextension.SizeStringName, unsafe.Pointer(&arg))
		case NodePath:
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypeNodePath, gdextension.SizeNodePath, unsafe.Pointer(&arg))
		case Object:
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypeObject, gdextension.SizeObject, unsafe.Pointer(&arg))
		case Callable:
			if pointers.Get(val) == ([2]uint64{}) {
				return Variant{}
			}
			var arg = gdextension.Callable(pointers.Cut(val, cut))
			ret.LoadNative(TypeCallable, gdextension.SizeCallable, unsafe.Pointer(&arg))
		case Signal:
			if pointers.Get(val) == ([2]uint64{}) {
				return Variant{}
			}
			var arg = gdextension.Signal(pointers.Cut(val, cut))
			ret.LoadNative(TypeSignal, gdextension.SizeSignal, unsafe.Pointer(&arg))
		case Dictionary:
			if pointers.Get(val) == (gdextension.Dictionary{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypeDictionary, gdextension.SizeDictionary, unsafe.Pointer(&arg))
		case Array:
			if pointers.Get(val) == (gdextension.Array{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypeArray, gdextension.SizeArray, unsafe.Pointer(&arg))
		case PackedByteArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedByteArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedInt32Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedInt32Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedInt64Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedInt64Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedFloat32Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedFloat32Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedFloat64Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedFloat64Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedStringArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedStringArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedVector2Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedVector2Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedVector3Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedVector3Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedVector4Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedVector4Array, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		case PackedColorArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			var arg = pointers.Cut(val, cut)
			ret.LoadNative(TypePackedColorArray, gdextension.SizePackedArray, unsafe.Pointer(&arg))
		default:
			var arg = pointers.Cut(newDictionary(value), cut)
			ret.LoadNative(TypeDictionary, gdextension.SizeDictionary, unsafe.Pointer(&arg))
		}

	default:
		panic("gd.Variant: unsupported type " + reflect.TypeOf(v).String())
	}
	if cut {
		return pointers.Let[Variant]([3]uint64(ret))
	} else {
		return pointers.New[Variant]([3]uint64(ret))
	}
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
	gdextension.Host.Builtin.Types.SetupArray(pointers.Get(array), vtype, gdextension.StringName{}, [3]uint64{})
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
		return variantAsValueType[bool](variant, vtype)
	case TypeInt:
		return variantAsValueType[Int](variant, vtype)
	case TypeFloat:
		return variantAsValueType[Float](variant, vtype)
	case TypeString:
		s := variantAsPointerType[String](variant, vtype)
		return StringType.Via(StringProxy{}, pointers.Pack(s))
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
		return BasisType.Transposed(variantAsValueType[Basis](variant, vtype))
	case TypeTransform3D:
		return Transposed(variantAsValueType[Transform3D](variant, vtype))
	case TypeProjection:
		return variantAsValueType[Projection](variant, vtype)
	case TypeColor:
		return variantAsValueType[Color](variant, vtype)
	case TypeStringName:
		s := variantAsPointerType[StringName](variant, vtype)
		return StringType.Name(StringType.Via(StringNameProxy{}, pointers.Pack(s)))
	case TypeNodePath:
		s := variantAsPointerType[NodePath](variant, vtype)
		return Path.ToNode(StringType.Via(NodePathProxy{}, pointers.Pack(s)))
	case TypeRID:
		return variantAsValueType[RID](variant, vtype)
	case TypeObject:
		var obj = VariantAsObject(variant)
		return ObjectAs(obj.GetClass().String(), obj)
	case TypeCallable:
		callable := variantAsPointerType[Callable](variant, vtype)
		return CallableType.Through(CallableProxy{}, pointers.Pack(callable))
	case TypeSignal:
		signal := variantAsPointerType[Signal](variant, vtype)
		return SignalType.Via(SignalProxy{}, pointers.Pack(signal))
	case TypeDictionary:
		dict := variantAsPointerType[Dictionary](variant, vtype)
		return DictionaryType.Through(DictionaryProxy[VariantPkg.Any, VariantPkg.Any]{}, pointers.Pack(dict))
	case TypeArray:
		array := variantAsPointerType[Array](variant, vtype)
		return ArrayType.Through(ArrayProxy[VariantPkg.Any]{}, pointers.Pack(array))
	case TypePackedByteArray:
		array := variantAsPointerType[PackedByteArray](variant, vtype)
		return PackedType.Bytes(ArrayType.Through(PackedProxy[PackedByteArray, byte]{}, pointers.Pack(array)))
	case TypePackedInt32Array:
		array := variantAsPointerType[PackedInt32Array](variant, vtype)
		return PackedType.Array[int32](ArrayType.Through(PackedProxy[PackedInt32Array, int32]{}, pointers.Pack(array)))
	case TypePackedInt64Array:
		array := variantAsPointerType[PackedInt64Array](variant, vtype)
		return PackedType.Array[int64](ArrayType.Through(PackedProxy[PackedInt64Array, int64]{}, pointers.Pack(array)))
	case TypePackedFloat32Array:
		array := variantAsPointerType[PackedFloat32Array](variant, vtype)
		return PackedType.Array[float32](ArrayType.Through(PackedProxy[PackedFloat32Array, float32]{}, pointers.Pack(array)))
	case TypePackedFloat64Array:
		array := variantAsPointerType[PackedFloat64Array](variant, vtype)
		return PackedType.Array[float64](ArrayType.Through(PackedProxy[PackedFloat64Array, float64]{}, pointers.Pack(array)))
	case TypePackedStringArray:
		array := variantAsPointerType[PackedStringArray](variant, vtype)
		return PackedType.Strings(ArrayType.Through(PackedStringArrayProxy{}, pointers.Pack(array)))
	case TypePackedVector2Array:
		array := variantAsPointerType[PackedVector2Array](variant, vtype)
		return PackedType.Array[Vector2](ArrayType.Through(PackedProxy[PackedVector2Array, Vector2]{}, pointers.Pack(array)))
	case TypePackedVector3Array:
		array := variantAsPointerType[PackedVector3Array](variant, vtype)
		return PackedType.Array[Vector3](ArrayType.Through(PackedProxy[PackedVector3Array, Vector3]{}, pointers.Pack(array)))
	case TypePackedVector4Array:
		array := variantAsPointerType[PackedVector4Array](variant, vtype)
		return PackedType.Array[Vector4](ArrayType.Through(PackedProxy[PackedVector4Array, Vector4]{}, pointers.Pack(array)))
	case TypePackedColorArray:
		array := variantAsPointerType[PackedColorArray](variant, vtype)
		return PackedType.Array[Color](ArrayType.Through(PackedProxy[PackedColorArray, Color]{}, pointers.Pack(array)))
	default:
		panic("gd.Variant.Interface: invalid variant type " + fmt.Sprint(uint32(vtype)))
	}
}

func variantAsValueType[T gdextension.AnyVariant](variant Variant, vtype VariantType) T {
	return gdextension.LoadNative[T](vtype, gdextension.Variant(pointers.Get(variant)))
}

func variantAsPointerType[T pointers.Generic[T, Size], Size gdextension.AnyPointer](variant Variant, vtype VariantType) T {
	return pointers.New[T](gdextension.LoadNative[Size](vtype, gdextension.Variant(pointers.Get(variant))))
}

func VariantAsObject(variant Variant) Object {
	raw, kind := pointers.Ask(variant)
	obj := gdextension.LoadNative[gdextension.Object](TypeObject, raw)
	switch kind {
	case pointers.Letted:
		return pointers.Let[Object]([3]uint64{uint64(obj)})
	default:
		return PointerMustAssertInstanceID[Object](obj)
	}
}

func LetVariantAsPointerType[T pointers.Generic[T, Size], Size gdextension.AnyPointer](variant Variant, vtype VariantType) T {
	return pointers.Let[T](gdextension.LoadNative[Size](vtype, gdextension.Variant(pointers.Get(variant))))
}

var ObjectAs = func(name string, ptr Object) any {
	return ptr
}
