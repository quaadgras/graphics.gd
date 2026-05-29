package gd

import (
	"reflect"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdreference"
	"graphics.gd/variant"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	CallableType "graphics.gd/variant/Callable"
	"graphics.gd/variant/Color"
	DictionaryType "graphics.gd/variant/Dictionary"
	FloatType "graphics.gd/variant/Float"
	PackedType "graphics.gd/variant/Packed"
	"graphics.gd/variant/Path"
	SignalType "graphics.gd/variant/Signal"
	StringType "graphics.gd/variant/String"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
)

func ConvieniantGoTypeOf(vtype variant.Type) reflect.Type {
	switch vtype {
	case variant.TypeNil:
		return reflect.TypeFor[any]()
	case variant.TypeBool:
		return reflect.TypeFor[bool]()
	case variant.TypeInt:
		return reflect.TypeFor[int]()
	case variant.TypeFloat:
		return reflect.TypeFor[FloatType.X]()
	case variant.TypeString:
		return reflect.TypeFor[string]()
	case variant.TypeVector2:
		return reflect.TypeFor[Vector2]()
	case variant.TypeVector2i:
		return reflect.TypeFor[Vector2i]()
	case variant.TypeRect2:
		return reflect.TypeFor[Rect2]()
	case variant.TypeRect2i:
		return reflect.TypeFor[Rect2i]()
	case variant.TypeVector3:
		return reflect.TypeFor[Vector3]()
	case variant.TypeVector3i:
		return reflect.TypeFor[Vector3i]()
	case variant.TypeTransform2D:
		return reflect.TypeFor[Transform2D]()
	case variant.TypeVector4:
		return reflect.TypeFor[Vector4]()
	case variant.TypeVector4i:
		return reflect.TypeFor[Vector4i]()
	case variant.TypePlane:
		return reflect.TypeFor[Plane]()
	case variant.TypeQuaternion:
		return reflect.TypeFor[Quaternion]()
	case variant.TypeAABB:
		return reflect.TypeFor[AABB]()
	case variant.TypeBasis:
		return reflect.TypeFor[Basis]()
	case variant.TypeTransform3D:
		return reflect.TypeFor[Transform3D]()
	case variant.TypeProjection:
		return reflect.TypeFor[Projection]()
	case variant.TypeColor:
		return reflect.TypeFor[Color]()
	case variant.TypeStringName:
		return reflect.TypeFor[string]()
	case variant.TypeNodePath:
		return reflect.TypeFor[Path.ToNode]()
	case variant.TypeRID:
		return reflect.TypeFor[RID]()
	case variant.TypeObject:
		return reflect.TypeFor[gdreference.Object]()
	case variant.TypeCallable:
		return reflect.TypeFor[CallableType.Function]()
	case variant.TypeSignal:
		return reflect.TypeFor[SignalType.Any]()
	case variant.TypeDictionary:
		return reflect.TypeFor[DictionaryType.Any]()
	case variant.TypeArray:
		return reflect.TypeFor[ArrayType.Any]()
	case variant.TypePackedByteArray:
		return reflect.TypeFor[PackedType.Bytes]()
	case variant.TypePackedInt32Array:
		return reflect.TypeFor[PackedType.Array[int32]]()
	case variant.TypePackedInt64Array:
		return reflect.TypeFor[PackedType.Array[int64]]()
	case variant.TypePackedFloat32Array:
		return reflect.TypeFor[PackedType.Array[float32]]()
	case variant.TypePackedFloat64Array:
		return reflect.TypeFor[PackedType.Array[float64]]()
	case variant.TypePackedStringArray:
		return reflect.TypeFor[PackedType.Strings]()
	case variant.TypePackedVector2Array:
		return reflect.TypeFor[PackedType.Array[Vector2]]()
	case variant.TypePackedVector3Array:
		return reflect.TypeFor[PackedType.Array[Vector3]]()
	case variant.TypePackedColorArray:
		return reflect.TypeFor[PackedType.Array[Color]]()
	case variant.TypePackedVector4Array:
		return reflect.TypeFor[PackedType.Array[Vector4]]()
	default:
		return nil
	}
}

func VariantTypeOf(rtype reflect.Type) (vtype variant.Type, ok bool) {
	if rtype == reflect.TypeFor[error]() {
		return variant.TypeInt, true
	}
	switch rtype.Kind() {
	case reflect.Bool:
		return variant.TypeBool, true
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint:
		return variant.TypeInt, true
	case reflect.Uint64:
		return variant.TypeRID, true
	case reflect.Float32, reflect.Float64:
		return variant.TypeFloat, true
	case reflect.Complex64, reflect.Complex128:
		return variant.TypeVector2, true
	case reflect.Pointer:
		if rtype.Implements(reflect.TypeFor[IsClass]()) {
			return variant.TypeObject, true
		}
		return VariantTypeOf(rtype.Elem())
	case reflect.Func:
		return variant.TypeCallable, true
	case reflect.Array:
		if rtype.Implements(reflect.TypeFor[IsClass]()) {
			return variant.TypeObject, true
		}
		return variant.TypeArray, true
	case reflect.String:
		if rtype == reflect.TypeFor[Path.ToNode]() {
			return variant.TypeNodePath, true
		}
		return variant.TypeString, true
	case reflect.Slice:
		switch rtype.Elem().Kind() {
		case reflect.Uint8:
			return variant.TypePackedByteArray, true
		case reflect.Int32:
			return variant.TypePackedInt32Array, true
		case reflect.Int64:
			return variant.TypePackedInt64Array, true
		case reflect.Float32:
			return variant.TypePackedFloat32Array, true
		case reflect.Float64:
			return variant.TypePackedFloat64Array, true
		case reflect.String:
			return variant.TypePackedStringArray, true
		default:
			switch rtype.Elem() {
			case reflect.TypeFor[Vector2]():
				return variant.TypePackedVector2Array, true
			case reflect.TypeFor[Vector3]():
				return variant.TypePackedVector3Array, true
			case reflect.TypeFor[Color]():
				return variant.TypePackedColorArray, true
			case reflect.TypeFor[Vector4]():
				return variant.TypePackedVector4Array, true
			default:
				return variant.TypeArray, true
			}
		}
	case reflect.Map:
		return variant.TypeDictionary, true
	case reflect.Interface:
		if rtype == reflect.TypeFor[any]() {
			return variant.TypeNil, true
		}
		return variant.TypeNil, false
	case reflect.Struct:
		switch rtype {
		case reflect.TypeFor[VariantPkg.Any]():
			return variant.TypeNil, true
		case reflect.TypeFor[PackedType.Bytes]():
			return variant.TypePackedByteArray, true
		case reflect.TypeFor[PackedType.Array[int32]]():
			return variant.TypePackedInt32Array, true
		case reflect.TypeFor[PackedType.Array[int64]]():
			return variant.TypePackedInt64Array, true
		case reflect.TypeFor[PackedType.Array[float32]]():
			return variant.TypePackedFloat32Array, true
		case reflect.TypeFor[PackedType.Array[float64]]():
			return variant.TypePackedFloat64Array, true
		case reflect.TypeFor[PackedType.Strings]():
			return variant.TypePackedStringArray, true
		case reflect.TypeFor[PackedType.Array[Vector2]]():
			return variant.TypePackedVector2Array, true
		case reflect.TypeFor[PackedType.Array[Vector3]]():
			return variant.TypePackedVector3Array, true
		case reflect.TypeFor[PackedType.Array[Color]]():
			return variant.TypePackedColorArray, true
		case reflect.TypeFor[PackedType.Array[Vector4]]():
			return variant.TypePackedVector4Array, true
		case reflect.TypeFor[ArrayType.Any]():
			return variant.TypeArray, true
		case reflect.TypeFor[StringType.Unicode]():
			return variant.TypeString, true
		case reflect.TypeFor[Path.ToNode]():
			return variant.TypeNodePath, true
		case reflect.TypeFor[StringType.Name]():
			return variant.TypeStringName, true
		case reflect.TypeFor[DictionaryType.Any]():
			return variant.TypeDictionary, true
		case reflect.TypeFor[SignalType.Any]():
			return variant.TypeSignal, true
		case reflect.TypeFor[[0]gdunsafe.Variant]().Elem():
			vtype = variant.TypeNil
		case reflect.TypeFor[[0]bool]().Elem():
			vtype = variant.TypeBool
		case reflect.TypeFor[[0]Int]().Elem():
			vtype = variant.TypeInt
		case reflect.TypeFor[[0]Float]().Elem():
			vtype = variant.TypeFloat
		case reflect.TypeFor[[0]gdunsafe.String]().Elem():
			vtype = variant.TypeString
		case reflect.TypeFor[[0]Vector2.XY]().Elem():
			vtype = variant.TypeVector2
		case reflect.TypeFor[[0]Vector2i]().Elem():
			vtype = variant.TypeVector2i
		case reflect.TypeFor[[0]Rect2]().Elem():
			vtype = variant.TypeRect2
		case reflect.TypeFor[[0]Rect2i]().Elem():
			vtype = variant.TypeRect2i
		case reflect.TypeFor[[0]Vector3.XYZ]().Elem():
			vtype = variant.TypeVector3
		case reflect.TypeFor[[0]Vector3i]().Elem():
			vtype = variant.TypeVector3i
		case reflect.TypeFor[[0]Transform2D]().Elem():
			vtype = variant.TypeTransform2D
		case reflect.TypeFor[[0]Vector4]().Elem():
			vtype = variant.TypeVector4
		case reflect.TypeFor[[0]Vector4i]().Elem():
			vtype = variant.TypeVector4i
		case reflect.TypeFor[[0]Plane]().Elem():
			vtype = variant.TypePlane
		case reflect.TypeFor[[0]Quaternion]().Elem():
			vtype = variant.TypeQuaternion
		case reflect.TypeFor[[0]AABB]().Elem():
			vtype = variant.TypeAABB
		case reflect.TypeFor[[0]Basis]().Elem():
			vtype = variant.TypeBasis
		case reflect.TypeFor[[0]Transform3D]().Elem():
			vtype = variant.TypeTransform3D
		case reflect.TypeFor[[0]Projection]().Elem():
			vtype = variant.TypeProjection
		case reflect.TypeFor[[0]Color.RGBA]().Elem():
			vtype = variant.TypeColor
		case reflect.TypeFor[[0]gdunsafe.StringName]().Elem():
			vtype = variant.TypeStringName
		case reflect.TypeFor[[0]gdunsafe.NodePath]().Elem():
			vtype = variant.TypeNodePath
		case reflect.TypeFor[[0]RID]().Elem():
			vtype = variant.TypeRID
		case reflect.TypeFor[[0]gdreference.Object]().Elem():
			vtype = variant.TypeObject
		case reflect.TypeFor[gdunsafe.Callable](), reflect.TypeFor[CallableType.Function]():
			vtype = variant.TypeCallable
		case reflect.TypeFor[[0]gdunsafe.Dictionary]().Elem():
			vtype = variant.TypeDictionary
		case reflect.TypeFor[[0]gdunsafe.Array]().Elem():
			vtype = variant.TypeArray
		case reflect.TypeFor[[0]gdunsafe.PackedArray[byte]]().Elem():
			vtype = variant.TypePackedByteArray
		case reflect.TypeFor[[0]gdunsafe.PackedArray[int32]]().Elem():
			vtype = variant.TypePackedInt32Array
		case reflect.TypeFor[[0]gdunsafe.PackedArray[int64]]().Elem():
			vtype = variant.TypePackedInt64Array
		case reflect.TypeFor[[0]gdunsafe.PackedArray[float32]]().Elem():
			vtype = variant.TypePackedFloat32Array
		case reflect.TypeFor[[0]gdunsafe.PackedArray[float64]]().Elem():
			vtype = variant.TypePackedFloat64Array
		case reflect.TypeFor[[0]gdunsafe.PackedArray[gdunsafe.String]]().Elem():
			vtype = variant.TypePackedStringArray
		case reflect.TypeFor[[0]gdunsafe.PackedArray[Vector2.XY]]().Elem():
			vtype = variant.TypePackedVector2Array
		case reflect.TypeFor[[0]gdunsafe.PackedArray[Vector3.XYZ]]().Elem():
			vtype = variant.TypePackedVector3Array
		case reflect.TypeFor[[0]gdunsafe.PackedArray[Color.RGBA]]().Elem():
			vtype = variant.TypePackedColorArray
		case reflect.TypeFor[VariantPkg.Any]():
			vtype = variant.TypeNil
		case reflect.TypeFor[[0]unsafe.Pointer]().Elem():
			vtype = variant.TypeNil
		case reflect.TypeFor[[0]*ScriptLanguageExtensionProfilingInfo]().Elem():
			vtype = variant.TypeNil
		default:
			switch {
			case rtype.Implements(reflect.TypeFor[[0]IsClass]().Elem()):
				vtype = variant.TypeObject
			case reflect.PointerTo(rtype).Implements(reflect.TypeFor[SignalType.Pointer]()):
				vtype = variant.TypeSignal
			case rtype.Implements(reflect.TypeFor[ArrayType.Interface]()):
				vtype = variant.TypeArray
			case rtype.Implements(reflect.TypeFor[DictionaryType.Interface]()):
				vtype = variant.TypeDictionary
			default:
				vtype = variant.TypeDictionary
			}
		}
		return vtype, true
	default:
		return variant.TypeNil, false
	}
}
