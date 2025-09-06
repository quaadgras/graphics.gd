package gd

import (
	"reflect"
	"unsafe"

	"graphics.gd/internal/gdextension"
	VariantPkg "graphics.gd/variant"
	ArrayType "graphics.gd/variant/Array"
	CallableType "graphics.gd/variant/Callable"
	DictionaryType "graphics.gd/variant/Dictionary"
	FloatType "graphics.gd/variant/Float"
	PackedType "graphics.gd/variant/Packed"
	"graphics.gd/variant/Path"
	SignalType "graphics.gd/variant/Signal"
	StringType "graphics.gd/variant/String"
)

func ConvieniantGoTypeOf(vtype gdextension.VariantType) reflect.Type {
	switch vtype {
	case gdextension.TypeNil:
		return reflect.TypeFor[any]()
	case gdextension.TypeBool:
		return reflect.TypeFor[bool]()
	case gdextension.TypeInt:
		return reflect.TypeFor[int]()
	case gdextension.TypeFloat:
		return reflect.TypeFor[FloatType.X]()
	case gdextension.TypeString:
		return reflect.TypeFor[string]()
	case gdextension.TypeVector2:
		return reflect.TypeFor[Vector2]()
	case gdextension.TypeVector2i:
		return reflect.TypeFor[Vector2i]()
	case gdextension.TypeRect2:
		return reflect.TypeFor[Rect2]()
	case gdextension.TypeRect2i:
		return reflect.TypeFor[Rect2i]()
	case gdextension.TypeVector3:
		return reflect.TypeFor[Vector3]()
	case gdextension.TypeVector3i:
		return reflect.TypeFor[Vector3i]()
	case gdextension.TypeTransform2D:
		return reflect.TypeFor[Transform2D]()
	case gdextension.TypeVector4:
		return reflect.TypeFor[Vector4]()
	case gdextension.TypeVector4i:
		return reflect.TypeFor[Vector4i]()
	case gdextension.TypePlane:
		return reflect.TypeFor[Plane]()
	case gdextension.TypeQuaternion:
		return reflect.TypeFor[Quaternion]()
	case gdextension.TypeAABB:
		return reflect.TypeFor[AABB]()
	case gdextension.TypeBasis:
		return reflect.TypeFor[Basis]()
	case gdextension.TypeTransform3D:
		return reflect.TypeFor[Transform3D]()
	case gdextension.TypeProjection:
		return reflect.TypeFor[Projection]()
	case gdextension.TypeColor:
		return reflect.TypeFor[Color]()
	case gdextension.TypeStringName:
		return reflect.TypeFor[string]()
	case gdextension.TypeNodePath:
		return reflect.TypeFor[Path.ToNode]()
	case gdextension.TypeRID:
		return reflect.TypeFor[RID]()
	case gdextension.TypeObject:
		return reflect.TypeFor[Object]()
	case gdextension.TypeCallable:
		return reflect.TypeFor[CallableType.Function]()
	case gdextension.TypeSignal:
		return reflect.TypeFor[SignalType.Any]()
	case gdextension.TypeDictionary:
		return reflect.TypeFor[DictionaryType.Any]()
	case gdextension.TypeArray:
		return reflect.TypeFor[ArrayType.Any]()
	case gdextension.TypePackedByteArray:
		return reflect.TypeFor[PackedType.Bytes]()
	case gdextension.TypePackedInt32Array:
		return reflect.TypeFor[PackedType.Array[int32]]()
	case gdextension.TypePackedInt64Array:
		return reflect.TypeFor[PackedType.Array[int64]]()
	case gdextension.TypePackedFloat32Array:
		return reflect.TypeFor[PackedType.Array[float32]]()
	case gdextension.TypePackedFloat64Array:
		return reflect.TypeFor[PackedType.Array[float64]]()
	case gdextension.TypePackedStringArray:
		return reflect.TypeFor[PackedType.Strings]()
	case gdextension.TypePackedVector2Array:
		return reflect.TypeFor[PackedType.Array[Vector2]]()
	case gdextension.TypePackedVector3Array:
		return reflect.TypeFor[PackedType.Array[Vector3]]()
	case gdextension.TypePackedColorArray:
		return reflect.TypeFor[PackedType.Array[Color]]()
	case gdextension.TypePackedVector4Array:
		return reflect.TypeFor[PackedType.Array[Vector4]]()
	default:
		return nil
	}
}

func VariantTypeOf(rtype reflect.Type) (vtype gdextension.VariantType, ok bool) {
	switch rtype.Kind() {
	case reflect.Bool:
		return gdextension.TypeBool, true
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint:
		return gdextension.TypeInt, true
	case reflect.Uint64:
		return gdextension.TypeRID, true
	case reflect.Float32, reflect.Float64:
		return gdextension.TypeFloat, true
	case reflect.Complex64, reflect.Complex128:
		return gdextension.TypeVector2, true
	case reflect.Pointer:
		if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			return gdextension.TypeObject, true
		}
		return VariantTypeOf(rtype.Elem())
	case reflect.Func:
		return gdextension.TypeCallable, true
	case reflect.Array:
		if rtype.Elem().Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
			return gdextension.TypeObject, true
		}
		return gdextension.TypeArray, true
	case reflect.String:
		if rtype == reflect.TypeFor[Path.ToNode]() {
			return gdextension.TypeNodePath, true
		}
		return gdextension.TypeString, true
	case reflect.Slice:
		switch rtype.Elem().Kind() {
		case reflect.Uint8:
			return gdextension.TypePackedByteArray, true
		case reflect.Int32:
			return gdextension.TypePackedInt32Array, true
		case reflect.Int64:
			return gdextension.TypePackedInt64Array, true
		case reflect.Float32:
			return gdextension.TypePackedFloat32Array, true
		case reflect.Float64:
			return gdextension.TypePackedFloat64Array, true
		case reflect.String:
			return gdextension.TypePackedStringArray, true
		default:
			switch rtype.Elem() {
			case reflect.TypeFor[Vector2]():
				return gdextension.TypePackedVector2Array, true
			case reflect.TypeFor[Vector3]():
				return gdextension.TypePackedVector3Array, true
			case reflect.TypeFor[Color]():
				return gdextension.TypePackedColorArray, true
			case reflect.TypeFor[Vector4]():
				return gdextension.TypePackedVector4Array, true
			default:
				return gdextension.TypeArray, true
			}
		}
	case reflect.Map:
		return gdextension.TypeDictionary, true
	case reflect.Interface:
		if rtype == reflect.TypeFor[any]() {
			return gdextension.TypeNil, true
		}
		return gdextension.TypeNil, false
	case reflect.Struct:
		switch rtype {
		case reflect.TypeFor[VariantPkg.Any]():
			return gdextension.TypeNil, true
		case reflect.TypeFor[PackedType.Bytes]():
			return gdextension.TypePackedByteArray, true
		case reflect.TypeFor[PackedType.Array[int32]]():
			return gdextension.TypePackedInt32Array, true
		case reflect.TypeFor[PackedType.Array[int64]]():
			return gdextension.TypePackedInt64Array, true
		case reflect.TypeFor[PackedType.Array[float32]]():
			return gdextension.TypePackedFloat32Array, true
		case reflect.TypeFor[PackedType.Array[float64]]():
			return gdextension.TypePackedFloat64Array, true
		case reflect.TypeFor[PackedType.Strings]():
			return gdextension.TypePackedStringArray, true
		case reflect.TypeFor[PackedType.Array[Vector2]]():
			return gdextension.TypePackedVector2Array, true
		case reflect.TypeFor[PackedType.Array[Vector3]]():
			return gdextension.TypePackedVector3Array, true
		case reflect.TypeFor[PackedType.Array[Color]]():
			return gdextension.TypePackedColorArray, true
		case reflect.TypeFor[PackedType.Array[Vector4]]():
			return gdextension.TypePackedVector4Array, true
		case reflect.TypeFor[ArrayType.Any]():
			return gdextension.TypeArray, true
		case reflect.TypeFor[StringType.Readable]():
			return gdextension.TypeString, true
		case reflect.TypeFor[Path.ToNode]():
			return gdextension.TypeNodePath, true
		case reflect.TypeFor[StringType.Name]():
			return gdextension.TypeStringName, true
		case reflect.TypeFor[DictionaryType.Any]():
			return gdextension.TypeDictionary, true
		case reflect.TypeFor[SignalType.Any]():
			return gdextension.TypeSignal, true
		case reflect.TypeOf([0]Variant{}).Elem():
			vtype = gdextension.TypeNil
		case reflect.TypeOf([0]bool{}).Elem():
			vtype = gdextension.TypeBool
		case reflect.TypeOf([0]Int{}).Elem():
			vtype = gdextension.TypeInt
		case reflect.TypeOf([0]Float{}).Elem():
			vtype = gdextension.TypeFloat
		case reflect.TypeOf([0]String{}).Elem():
			vtype = gdextension.TypeString
		case reflect.TypeOf([0]Vector2{}).Elem():
			vtype = gdextension.TypeVector2
		case reflect.TypeOf([0]Vector2i{}).Elem():
			vtype = gdextension.TypeVector2i
		case reflect.TypeOf([0]Rect2{}).Elem():
			vtype = gdextension.TypeRect2
		case reflect.TypeOf([0]Rect2i{}).Elem():
			vtype = gdextension.TypeRect2i
		case reflect.TypeOf([0]Vector3{}).Elem():
			vtype = gdextension.TypeVector3
		case reflect.TypeOf([0]Vector3i{}).Elem():
			vtype = gdextension.TypeVector3i
		case reflect.TypeOf([0]Transform2D{}).Elem():
			vtype = gdextension.TypeTransform2D
		case reflect.TypeOf([0]Vector4{}).Elem():
			vtype = gdextension.TypeVector4
		case reflect.TypeOf([0]Vector4i{}).Elem():
			vtype = gdextension.TypeVector4i
		case reflect.TypeOf([0]Plane{}).Elem():
			vtype = gdextension.TypePlane
		case reflect.TypeOf([0]Quaternion{}).Elem():
			vtype = gdextension.TypeQuaternion
		case reflect.TypeOf([0]AABB{}).Elem():
			vtype = gdextension.TypeAABB
		case reflect.TypeOf([0]Basis{}).Elem():
			vtype = gdextension.TypeBasis
		case reflect.TypeOf([0]Transform3D{}).Elem():
			vtype = gdextension.TypeTransform3D
		case reflect.TypeOf([0]Projection{}).Elem():
			vtype = gdextension.TypeProjection
		case reflect.TypeOf([0]Color{}).Elem():
			vtype = gdextension.TypeColor
		case reflect.TypeOf([0]StringName{}).Elem():
			vtype = gdextension.TypeStringName
		case reflect.TypeOf([0]NodePath{}).Elem():
			vtype = gdextension.TypeNodePath
		case reflect.TypeOf([0]RID{}).Elem():
			vtype = gdextension.TypeRID
		case reflect.TypeOf([0]Object{}).Elem():
			vtype = gdextension.TypeObject
		case reflect.TypeFor[Callable](), reflect.TypeFor[CallableType.Function]():
			vtype = gdextension.TypeCallable
		case reflect.TypeOf([0]Dictionary{}).Elem():
			vtype = gdextension.TypeDictionary
		case reflect.TypeOf([0]Array{}).Elem():
			vtype = gdextension.TypeArray
		case reflect.TypeOf([0]PackedByteArray{}).Elem():
			vtype = gdextension.TypePackedByteArray
		case reflect.TypeOf([0]PackedInt32Array{}).Elem():
			vtype = gdextension.TypePackedInt32Array
		case reflect.TypeOf([0]PackedInt64Array{}).Elem():
			vtype = gdextension.TypePackedInt64Array
		case reflect.TypeOf([0]PackedFloat32Array{}).Elem():
			vtype = gdextension.TypePackedFloat32Array
		case reflect.TypeOf([0]PackedFloat64Array{}).Elem():
			vtype = gdextension.TypePackedFloat64Array
		case reflect.TypeOf([0]PackedStringArray{}).Elem():
			vtype = gdextension.TypePackedStringArray
		case reflect.TypeOf([0]PackedVector2Array{}).Elem():
			vtype = gdextension.TypePackedVector2Array
		case reflect.TypeOf([0]PackedVector3Array{}).Elem():
			vtype = gdextension.TypePackedVector3Array
		case reflect.TypeOf([0]PackedColorArray{}).Elem():
			vtype = gdextension.TypePackedColorArray
		case reflect.TypeFor[VariantPkg.Any]():
			vtype = gdextension.TypeNil
		case reflect.TypeOf([0]unsafe.Pointer{}).Elem():
			vtype = gdextension.TypeNil
		case reflect.TypeOf([0]*ScriptLanguageExtensionProfilingInfo{}).Elem():
			vtype = gdextension.TypeNil
		default:
			switch {
			case rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()):
				vtype = gdextension.TypeObject
			case reflect.PointerTo(rtype).Implements(reflect.TypeFor[SignalType.Pointer]()):
				vtype = gdextension.TypeSignal
			case rtype.Implements(reflect.TypeFor[ArrayType.Interface]()):
				vtype = gdextension.TypeArray
			case rtype.Implements(reflect.TypeFor[DictionaryType.Interface]()):
				vtype = gdextension.TypeDictionary
			default:
				vtype = gdextension.TypeDictionary
			}
		}
		return vtype, true
	default:
		return gdextension.TypeNil, false
	}
}
