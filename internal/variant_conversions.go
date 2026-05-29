//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"time"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"
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
func NewVariant(v any) gdunsafe.Variant {
	return CutVariant(v, false)
}

// CutVariant is like NewVariant but when cut is true, releases the ownership
// of the given value. Use it on return values passed back to the engine.
//
// used to fix cases of https://github.com/quaadgras/graphics.gd/issues/147
func CutVariant(v any, cut bool) gdunsafe.Variant {
	if v == nil {
		return gdunsafe.Variant{}
	}
	var ret gdunsafe.Variant
	if enum, ok := v.(Enum.Any); ok {
		v = enum.Int()
	}
	rtype := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	switch rtype.Kind() {
	case reflect.Bool:
		ret = gdunsafe.VariantFrom(value.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = gdunsafe.VariantFrom(value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		ret = gdunsafe.VariantFrom(Int(value.Uint()))
	case reflect.Uint64:
		if instance := value.MethodByName("Instance"); instance.IsValid() && instance.Type().NumOut() == 2 && instance.Type().NumIn() == 0 {
			result := instance.Call(nil)
			if !result[1].Bool() {
				return gdunsafe.Variant{}
			}
			obj := result[0].Interface().(IsClass).AsObject()
			if gdreference.GetObject(gdreference.Object(obj[0])) == (gdunsafe.Object{}) {
				return gdunsafe.Variant{}
			}
			var arg = gdreference.CutObject(gdreference.Object(obj[0]), cut)
			if cut {
				ExtensionInstanceGoOnly(arg, false)
			}
			ret = gdunsafe.VariantFrom(gdunsafe.Object(arg))
		} else {
			ret = gdunsafe.VariantFrom(RID(value.Uint()))
		}
	case reflect.Float32, reflect.Float64:
		ret = gdunsafe.VariantFrom(value.Float())
	case reflect.Complex64, reflect.Complex128:
		ret = gdunsafe.VariantFrom(Vector2{
			X: FloatType.X(real(value.Complex())),
			Y: FloatType.X(imag(value.Complex())),
		})
	case reflect.Pointer:
		if value.IsNil() {
			return gdunsafe.Variant{}
		}
		if rtype.Implements(reflect.TypeFor[[0]IsClass]().Elem()) {
			obj := value.Interface().(IsClass).AsObject()
			if gdreference.GetObject(obj[0]) == (gdunsafe.Object{}) {
				return gdunsafe.Variant{}
			}
			var arg = gdreference.GetObject(obj[0])
			if cut {
				ExtensionInstanceGoOnly(arg, false)
			}
			ret = gdunsafe.VariantFrom(gdunsafe.Object(arg))
		} else {
			return NewVariant(value.Elem().Interface())
		}
	case reflect.Array:
		if rtype.Elem().Implements(reflect.TypeFor[ObjectAny]()) {
			rtype = rtype.Elem()
			value = value.Index(0)
		}
		if rtype.Implements(reflect.TypeFor[ObjectAny]()) {
			anyobj, _ := reflect.TypeAssert[ObjectAny](value)
			obj := anyobj.AsObject()
			if gdreference.GetObject(obj[0]) == (gdunsafe.Object{}) {
				return gdunsafe.Variant{}
			}
			var arg = gdreference.CutObject(obj[0], cut)
			if cut {
				ExtensionInstanceGoOnly(arg, false)
			}
			ret = gdunsafe.VariantFrom(gdunsafe.Object(arg))
		} else {
			ret = gdunsafe.VariantFrom(gdunsafe.Array(pointers.Cut(newArray(value), cut)[0]))
		}
	case reflect.Slice:
		switch value := value.Interface().(type) {
		case []byte:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[byte](pointers.Cut(NewPackedByteSlice(value), cut)))
		case []int32:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[int32](pointers.Cut(NewPackedInt32Slice(value), cut)))
		case []int64:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[int64](pointers.Cut(NewPackedInt64Slice(value), cut)))
		case []float32:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[float32](pointers.Cut(NewPackedFloat32Slice(value), cut)))
		case []float64:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[float64](pointers.Cut(NewPackedFloat64Slice(value), cut)))
		case []string:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[gdunsafe.String](pointers.Cut(NewPackedStringSlice(value), cut)))
		case []StringType.Unicode:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[gdunsafe.String](pointers.Cut(NewPackedReadableStringSlice(value), cut)))
		case []Vector2:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector2](pointers.Cut(NewPackedVector2Slice(value), cut)))
		case []Vector3:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector3](pointers.Cut(NewPackedVector3Slice(value), cut)))
		case []Color:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Color](pointers.Cut(NewPackedColorSlice(value), cut)))
		case []Vector4:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector4](pointers.Cut(NewPackedVector4Slice(value), cut)))
		default:
			ret = gdunsafe.VariantFrom(gdunsafe.Array(pointers.Cut(newArray(reflect.ValueOf(value)), cut)[0]))
		}
	case reflect.Func:
		if value.IsNil() {
			return Variant{}
		}
		ret = gdunsafe.VariantFrom(gdunsafe.Callable(pointers.Get(NewCallable(value))))
	case reflect.Map:
		if value.IsNil() {
			return Variant{}
		}
		ret = gdunsafe.VariantFrom(gdunsafe.Dictionary(pointers.Cut(newDictionary(value), cut)[0]))
	case reflect.String:
		switch rtype {
		case reflect.TypeFor[Path.ToNode]():
			ret = gdunsafe.VariantFrom(gdunsafe.NodePath(pointers.Cut(NewString(value.String()).NodePath(), cut)[0]))
		case reflect.TypeFor[StringType.Name]():
			ret = gdunsafe.VariantFrom(gdunsafe.StringName(pointers.Cut(NewStringName(value.String()), cut)[0]))
		case reflect.TypeFor[StringType.Unicode]():
			ret = gdunsafe.VariantFrom(gdunsafe.String(pointers.Cut(NewString(value.String()), cut)[0]))
		default:
			ret = gdunsafe.VariantFrom(gdunsafe.String(pointers.Cut(NewString(value.String()), cut)[0]))
		}
	case reflect.Struct:
		switch val := v.(type) {
		case PackedType.Bytes:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[byte](pointers.Cut(InternalPacked[PackedByteArray, byte](PackedType.Array[byte](val.Array)), cut)))
		case PackedType.Array[byte]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[byte](pointers.Cut(InternalPacked[PackedByteArray, byte](val), cut)))
		case PackedType.Array[int32]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[int32](pointers.Cut(InternalPacked[PackedInt32Array, int32](val), cut)))
		case PackedType.Array[int64]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[int64](pointers.Cut(InternalPacked[PackedInt64Array, int64](val), cut)))
		case PackedType.Array[float32]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[float32](pointers.Cut(InternalPacked[PackedFloat32Array, float32](val), cut)))
		case PackedType.Array[float64]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[float64](pointers.Cut(InternalPacked[PackedFloat64Array, float64](val), cut)))
		case PackedType.Strings:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[gdunsafe.String](pointers.Cut(InternalPackedStrings(val), cut)))
		case PackedType.Array[Vector2]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector2](pointers.Cut(InternalPacked[PackedVector2Array, Vector2](val), cut)))
		case PackedType.Array[Vector3]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector3](pointers.Cut(InternalPacked[PackedVector3Array, Vector3](val), cut)))
		case PackedType.Array[Color]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Color](pointers.Cut(InternalPacked[PackedColorArray, Color](val), cut)))
		case PackedType.Array[Vector4]:
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector4](pointers.Cut(InternalPacked[PackedVector4Array, Vector4](val), cut)))
		case ArrayType.Any:
			ret = gdunsafe.VariantFrom(gdunsafe.Array(pointers.Cut(InternalArray(val), cut)[0]))
		case ArrayType.Interface:
			ret = gdunsafe.VariantFrom(gdunsafe.Array(pointers.Cut(InternalArray(val.Any()), cut)[0]))
		case DictionaryType.Any:
			ret = gdunsafe.VariantFrom(gdunsafe.Dictionary(pointers.Cut(InternalDictionary(val), cut)[0]))
		case DictionaryType.Interface:
			ret = gdunsafe.VariantFrom(gdunsafe.Dictionary(pointers.Cut(InternalDictionary(val.Any()), cut)[0]))
		case StringType.Unicode:
			ret = gdunsafe.VariantFrom(gdunsafe.String(pointers.Cut(InternalString(val), cut)[0]))
		case Path.ToNode:
			ret = gdunsafe.VariantFrom(gdunsafe.NodePath(pointers.Cut(InternalNodePath(val), cut)[0]))
		case StringType.Name:
			ret = gdunsafe.VariantFrom(gdunsafe.StringName(pointers.Cut(InternalStringName(val), cut)[0]))
		case CallableType.Function:
			ret = gdunsafe.VariantFrom(gdunsafe.Callable(pointers.Cut(InternalCallable(val), cut)))
		case SignalType.Any:
			ret = gdunsafe.VariantFrom(gdunsafe.Signal(pointers.Cut(InternalSignal(val), cut)))
		case Variant:
			return val
		case VariantPkg.Any:
			return CutVariant(val.Interface(), cut)
		case Vector2:
			ret = gdunsafe.VariantFrom(val)
		case Vector2i:
			ret = gdunsafe.VariantFrom(val)
		case Rect2:
			ret = gdunsafe.VariantFrom(val)
		case Rect2i:
			ret = gdunsafe.VariantFrom(val)
		case Vector3:
			ret = gdunsafe.VariantFrom(val)
		case Euler.Radians:
			ret = gdunsafe.VariantFrom[Vector3](Vector3{X: FloatType.X(val.X), Y: FloatType.X(val.Y), Z: FloatType.X(val.Z)})
		case Euler.Degrees:
			ret = gdunsafe.VariantFrom[Vector3](Vector3{X: FloatType.X(val.X), Y: FloatType.X(val.Y), Z: FloatType.X(val.Z)})
		case Vector3i:
			ret = gdunsafe.VariantFrom(val)
		case Transform2D:
			ret = gdunsafe.VariantFrom(val)
		case Vector4:
			ret = gdunsafe.VariantFrom(val)
		case Vector4i:
			ret = gdunsafe.VariantFrom(val)
		case Plane:
			ret = gdunsafe.VariantFrom(val)
		case Quaternion:
			ret = gdunsafe.VariantFrom(val)
		case AABB:
			ret = gdunsafe.VariantFrom(val)
		case Basis:
			ret = gdunsafe.VariantFrom(val)
		case Transform3D:
			ret = gdunsafe.VariantFrom(val)
		case Projection:
			ret = gdunsafe.VariantFrom(val)
		case Color:
			ret = gdunsafe.VariantFrom(val)
		case String:
			ret = gdunsafe.VariantFrom(gdunsafe.String(pointers.Cut(val, cut)[0]))
		case StringName:
			ret = gdunsafe.VariantFrom(gdunsafe.StringName(pointers.Cut(val, cut)[0]))
		case NodePath:
			ret = gdunsafe.VariantFrom(gdunsafe.NodePath(pointers.Cut(val, cut)[0]))
		case gdreference.Object:
			var arg = gdreference.CutObject(val, cut)
			if cut {
				ExtensionInstanceGoOnly(arg, false)
			}
			ret = gdunsafe.VariantFrom(arg)
		case Callable:
			if pointers.Get(val) == ([2]uint64{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.Callable(pointers.Cut(val, cut)))
		case Signal:
			if pointers.Get(val) == ([2]uint64{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.Signal(pointers.Cut(val, cut)))
		case Dictionary:
			if pointers.Get(val) == (gdextension.Dictionary{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.Dictionary(pointers.Cut(val, cut)[0]))
		case Array:
			if pointers.Get(val) == (gdextension.Array{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.Array(pointers.Cut(val, cut)[0]))
		case PackedByteArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[byte](pointers.Cut(val, cut)))
		case PackedInt32Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[int32](pointers.Cut(val, cut)))
		case PackedInt64Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[int64](pointers.Cut(val, cut)))
		case PackedFloat32Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[float32](pointers.Cut(val, cut)))
		case PackedFloat64Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[float64](pointers.Cut(val, cut)))
		case PackedStringArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[gdunsafe.String](pointers.Cut(val, cut)))
		case PackedVector2Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector2](pointers.Cut(val, cut)))
		case PackedVector3Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector3](pointers.Cut(val, cut)))
		case PackedVector4Array:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Vector4](pointers.Cut(val, cut)))
		case PackedColorArray:
			if pointers.Get(val) == (PackedPointers{}) {
				return Variant{}
			}
			ret = gdunsafe.VariantFrom(gdunsafe.PackedArray[Color](pointers.Cut(val, cut)))
		case time.Time:
			ret = gdunsafe.VariantFrom(val.UnixNano())
		default:
			ret = gdunsafe.VariantFrom(gdunsafe.Dictionary(pointers.Cut(newDictionary(value), cut)[0]))
		}

	default:
		panic("gd.Variant: unsupported type " + reflect.TypeOf(v).String())
	}
	if cut {
		return pointers.Let[Variant](ret)
	} else {
		return pointers.New[Variant](ret)
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
		for field, rvalue := range val.Fields() {
			if !field.IsExported() {
				continue
			}
			name := field.Name
			if tag := field.Tag.Get("gd"); tag != "" {
				name = tag
			}
			dict.SetIndex(NewVariant(name), NewVariant(rvalue.Interface()))
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
	gdunsafe.Array(pointers.Get(array)[0]).SetType(gdunsafe.TypeFrom(vtype))
	array.Resize(Int(val.Len()))
	for i := 0; i < val.Len(); i++ {
		array.SetIndex(Int(i), NewVariant(val.Index(i).Interface()))
	}
	return array
}

// Interface returns the variant's value as one of the the native Godot values
// (as defined) in the gd package.
func (v Variant) Interface() any {
	raw := gdunsafe.Variant(pointers.Get(v))
	switch vtype := v.Type(); vtype {
	case variant.TypeNil:
		return nil
	case variant.TypeBool:
		return gdunsafe.VariantInto[bool](raw)
	case variant.TypeInt:
		return gdunsafe.VariantInto[Int](raw)
	case variant.TypeFloat:
		return gdunsafe.VariantInto[float64](raw)
	case variant.TypeString:
		s := pointers.New[String](gdextension.String{gdextension.Pointer(gdunsafe.VariantInto[gdunsafe.String](raw))})
		return StringType.Via(StringProxy{}, pointers.Pack(s))
	case variant.TypeVector2:
		return gdunsafe.VariantInto[Vector2](raw)
	case variant.TypeVector2i:
		return gdunsafe.VariantInto[Vector2i](raw)
	case variant.TypeRect2:
		return gdunsafe.VariantInto[Rect2](raw)
	case variant.TypeRect2i:
		return gdunsafe.VariantInto[Rect2i](raw)
	case variant.TypeVector3:
		return gdunsafe.VariantInto[Vector3](raw)
	case variant.TypeVector3i:
		return gdunsafe.VariantInto[Vector3i](raw)
	case variant.TypeTransform2D:
		return gdunsafe.VariantInto[Transform2D](raw)
	case variant.TypeVector4:
		return gdunsafe.VariantInto[Vector4](raw)
	case variant.TypeVector4i:
		return gdunsafe.VariantInto[Vector4i](raw)
	case variant.TypePlane:
		return gdunsafe.VariantInto[Plane](raw)
	case variant.TypeQuaternion:
		return gdunsafe.VariantInto[Quaternion](raw)
	case variant.TypeAABB:
		return gdunsafe.VariantInto[AABB](raw)
	case variant.TypeBasis:
		return BasisType.Transposed(gdunsafe.VariantInto[Basis](raw))
	case variant.TypeTransform3D:
		return Transposed(gdunsafe.VariantInto[Transform3D](raw))
	case variant.TypeProjection:
		return gdunsafe.VariantInto[Projection](raw)
	case variant.TypeColor:
		return gdunsafe.VariantInto[Color](raw)
	case variant.TypeStringName:
		s := pointers.New[StringName](gdextension.StringName{gdextension.Pointer(gdunsafe.VariantInto[gdunsafe.StringName](raw))})
		return StringType.Name(StringType.Via(StringNameProxy{}, pointers.Pack(s)))
	case variant.TypeNodePath:
		s := pointers.New[NodePath](gdextension.NodePath{gdextension.Pointer(gdunsafe.VariantInto[gdunsafe.NodePath](raw))})
		return Path.ToNode(StringType.Via(NodePathProxy{}, pointers.Pack(s)))
	case variant.TypeRID:
		return gdunsafe.VariantInto[RID](raw)
	case variant.TypeObject:
		var obj = VariantAsObject(v)
		if gdreference.BadObject(obj) {
			return nil
		}
		return ObjectAs(ObjectGetClass(obj).String(), obj)
	case variant.TypeCallable:
		callable := pointers.New[Callable](gdextension.Callable(gdunsafe.VariantInto[gdunsafe.Callable](raw)))
		return CallableType.Through(CallableProxy{}, pointers.Pack(callable))
	case variant.TypeSignal:
		signal := pointers.New[Signal](gdextension.Signal(gdunsafe.VariantInto[gdunsafe.Signal](raw)))
		return SignalType.Via(SignalProxy{}, pointers.Pack(signal))
	case variant.TypeDictionary:
		dict := pointers.New[Dictionary](gdextension.Dictionary{gdextension.Pointer(gdunsafe.VariantInto[gdunsafe.Dictionary](raw))})
		return DictionaryType.Through(DictionaryProxy[VariantPkg.Any, VariantPkg.Any]{}, pointers.Pack(dict))
	case variant.TypeArray:
		array := pointers.New[Array](gdextension.Array{gdextension.Pointer(gdunsafe.VariantInto[gdunsafe.Array](raw))})
		return ArrayType.Through(ArrayProxy[VariantPkg.Any]{}, pointers.Pack(array))
	case variant.TypePackedByteArray:
		array := pointers.New[PackedByteArray](gdextension.PackedArray[byte](gdunsafe.VariantInto[gdunsafe.PackedArray[byte]](raw)))
		return PackedType.Bytes{Array: PackedType.Array[byte](ArrayType.Through(PackedProxy[PackedByteArray, byte]{}, pointers.Pack(array)))}
	case variant.TypePackedInt32Array:
		array := pointers.New[PackedInt32Array](gdextension.PackedArray[int32](gdunsafe.VariantInto[gdunsafe.PackedArray[int32]](raw)))
		return PackedType.Array[int32](ArrayType.Through(PackedProxy[PackedInt32Array, int32]{}, pointers.Pack(array)))
	case variant.TypePackedInt64Array:
		array := pointers.New[PackedInt64Array](gdextension.PackedArray[int64](gdunsafe.VariantInto[gdunsafe.PackedArray[int64]](raw)))
		return PackedType.Array[int64](ArrayType.Through(PackedProxy[PackedInt64Array, int64]{}, pointers.Pack(array)))
	case variant.TypePackedFloat32Array:
		array := pointers.New[PackedFloat32Array](gdextension.PackedArray[float32](gdunsafe.VariantInto[gdunsafe.PackedArray[float32]](raw)))
		return PackedType.Array[float32](ArrayType.Through(PackedProxy[PackedFloat32Array, float32]{}, pointers.Pack(array)))
	case variant.TypePackedFloat64Array:
		array := pointers.New[PackedFloat64Array](gdextension.PackedArray[float64](gdunsafe.VariantInto[gdunsafe.PackedArray[float64]](raw)))
		return PackedType.Array[float64](ArrayType.Through(PackedProxy[PackedFloat64Array, float64]{}, pointers.Pack(array)))
	case variant.TypePackedStringArray:
		array := pointers.New[PackedStringArray](gdextension.PackedArray[gdextension.String](gdunsafe.VariantInto[gdunsafe.PackedArray[gdunsafe.String]](raw)))
		return PackedType.Strings(ArrayType.Through(PackedStringArrayProxy{}, pointers.Pack(array)))
	case variant.TypePackedVector2Array:
		array := pointers.New[PackedVector2Array](gdextension.PackedArray[Vector2](gdunsafe.VariantInto[gdunsafe.PackedArray[Vector2]](raw)))
		return PackedType.Array[Vector2](ArrayType.Through(PackedProxy[PackedVector2Array, Vector2]{}, pointers.Pack(array)))
	case variant.TypePackedVector3Array:
		array := pointers.New[PackedVector3Array](gdextension.PackedArray[Vector3](gdunsafe.VariantInto[gdunsafe.PackedArray[Vector3]](raw)))
		return PackedType.Array[Vector3](ArrayType.Through(PackedProxy[PackedVector3Array, Vector3]{}, pointers.Pack(array)))
	case variant.TypePackedVector4Array:
		array := pointers.New[PackedVector4Array](gdextension.PackedArray[Vector4](gdunsafe.VariantInto[gdunsafe.PackedArray[Vector4]](raw)))
		return PackedType.Array[Vector4](ArrayType.Through(PackedProxy[PackedVector4Array, Vector4]{}, pointers.Pack(array)))
	case variant.TypePackedColorArray:
		array := pointers.New[PackedColorArray](gdextension.PackedArray[Color](gdunsafe.VariantInto[gdunsafe.PackedArray[Color]](raw)))
		return PackedType.Array[Color](ArrayType.Through(PackedProxy[PackedColorArray, Color]{}, pointers.Pack(array)))
	default:
		panic("gd.Variant.Interface: invalid variant type " + fmt.Sprint(uint32(vtype)))
	}
}

func VariantAsObject(variant Variant) gdreference.Object {
	return gdreference.LetObject(gdunsafe.VariantInto[gdunsafe.Object](gdunsafe.Variant(pointers.Get(variant))))
}

var ObjectAs = func(name string, ptr gdreference.Object) any {
	return ptr
}
