package variant

import (
	"reflect"
	"unsafe"

	"graphics.gd/internal/pointers"
	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
	"graphics.gd/variant/Vector4"
	"graphics.gd/variant/Vector4i"
)

// New returns a new [Any] variant from the given interface value.
func New(val any) Any {
	if val == nil {
		return Nil
	}
	already, ok := val.(Any)
	if ok {
		return already
	}
	var local complex128
	rvalue := reflect.ValueOf(val)
	rtype := rvalue.Type()
	switch rtype.Kind() {
	case reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		*(*int64)(unsafe.Pointer(&local)) = rvalue.Int()
		return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
	case reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		*(*uint64)(unsafe.Pointer(&local)) = rvalue.Uint()
		return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
	case reflect.Float32, reflect.Float64:
		*(*float64)(unsafe.Pointer(&local)) = rvalue.Float()
		return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
	case reflect.Complex64, reflect.Complex128:
		*(*complex128)(unsafe.Pointer(&local)) = rvalue.Complex()
		return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
	case reflect.String:
		s := rvalue.String()
		*(*int)(unsafe.Pointer(&local)) = len(s)
		wrapped := isString(unsafe.StringData(s))
		return Any{value: wrapped, local: local}
	case reflect.Slice:
		if rvalue.IsNil() {
			return Any{}
		}
		switch rtype.Elem() {
		case reflect.TypeFor[byte]():
			b := rvalue.Bytes()
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[byte](unsafe.SliceData(b))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[int32]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[int32](unsafe.SliceData(val.([]int32)))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[int64]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[int64](unsafe.SliceData(val.([]int64)))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[float32]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[float32](unsafe.SliceData(val.([]float32)))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[float64]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[float64](unsafe.SliceData(val.([]float64)))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[string]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[string](unsafe.SliceData(val.([]string)))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[Vector2.XY]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[Vector2.XY](unsafe.SliceData(val.([]Vector2.XY)))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[Vector3.XYZ]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[Vector3.XYZ](unsafe.SliceData(val.([]Vector3.XYZ)))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[Color.RGBA]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[Color.RGBA](unsafe.SliceData(val.([]Color.RGBA)))
			return Any{value: wrapped, local: local}
		case reflect.TypeFor[Vector4.XYZW]():
			*(*lencap)(unsafe.Pointer(&local)) = lencap{len: rvalue.Len(), cap: rvalue.Cap()}
			wrapped := isPacked[Vector4.XYZW](unsafe.SliceData(val.([]Vector4.XYZW)))
			return Any{value: wrapped, local: local}
		}
	case reflect.Struct:
		switch val.(type) {
		case Vector2.XY:
			*(*Vector2.XY)(unsafe.Pointer(&local)) = val.(Vector2.XY)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Vector2i.XY:
			*(*Vector2i.XY)(unsafe.Pointer(&local)) = val.(Vector2i.XY)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Rect2.PositionSize:
			*(*Rect2.PositionSize)(unsafe.Pointer(&local)) = val.(Rect2.PositionSize)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Rect2i.PositionSize:
			*(*Rect2i.PositionSize)(unsafe.Pointer(&local)) = val.(Rect2i.PositionSize)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Plane.NormalD:
			*(*Plane.NormalD)(unsafe.Pointer(&local)) = val.(Plane.NormalD)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Vector3.XYZ:
			*(*Vector3.XYZ)(unsafe.Pointer(&local)) = val.(Vector3.XYZ)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Vector3i.XYZ:
			*(*Vector3i.XYZ)(unsafe.Pointer(&local)) = val.(Vector3i.XYZ)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Vector4.XYZW:
			*(*Vector4.XYZW)(unsafe.Pointer(&local)) = val.(Vector4.XYZW)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Vector4i.XYZW:
			*(*Vector4i.XYZW)(unsafe.Pointer(&local)) = val.(Vector4i.XYZW)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		case Color.RGBA:
			*(*Color.RGBA)(unsafe.Pointer(&local)) = val.(Color.RGBA)
			return Any{value: reflect.Zero(reflect.PointerTo(rtype)).Interface(), local: local}
		}
	case reflect.Pointer, reflect.UnsafePointer, reflect.Map, reflect.Func, reflect.Chan, reflect.Interface:
		if rvalue.IsNil() {
			return Any{}
		}
	}
	return Any{value: val, local: local}
}

type packable interface {
	byte | int32 | int64 | float32 | float64 | string | Vector2.XY | Vector3.XYZ | Color.RGBA | Vector4.XYZW
}

type isString *byte
type isPacked[T packable] *T

type lencap = struct {
	len int
	cap int
}

type Packer interface {
	PackVariant() (Loader, [16]byte)
}
type Loader interface {
	LoadVariant([16]byte) any
}

type Setter interface {
	SetVariant(Any)
}

type Getter interface {
	Variant() Any
}

// Get attempts to convert the given variant value to the given type.
/*func Get[T any](val gd.Variant) T {
	converted, err := gd.ConvertToDesiredGoType(val, reflect.TypeFor[T]())
	if err != nil {
		panic(err)
	}
	return converted.Interface().(T)
}*/

// Call attempts to call the given method on the given variant value with the given arguments.
func (a Any) Call(name string, args ...Any) Any {
	rvalue := reflect.ValueOf(a.value)
	method := rvalue.MethodByName(name)
	if !method.IsValid() {
		panic("method not found: " + name)
	}
	arguments := make([]reflect.Value, len(args))
	for i, arg := range args {
		arguments[i] = reflect.ValueOf(arg.Interface())
	}
	results := method.Call(arguments)
	if len(results) == 0 {
		return Nil
	}
	return New(results[0].Interface())
}

type Type int

const (
	TypeNil                Type = 0
	TypeBool               Type = 1
	TypeInt                Type = 2
	TypeFloat              Type = 3
	TypeString             Type = 4
	TypeVector2            Type = 5
	TypeVector2i           Type = 6
	TypeRect2              Type = 7
	TypeRect2i             Type = 8
	TypeVector3            Type = 9
	TypeVector3i           Type = 10
	TypeTransform2D        Type = 11
	TypeVector4            Type = 12
	TypeVector4i           Type = 13
	TypePlane              Type = 14
	TypeQuaternion         Type = 15
	TypeAABB               Type = 16
	TypeBasis              Type = 17
	TypeTransform3D        Type = 18
	TypeProjection         Type = 19
	TypeColor              Type = 20
	TypeStringName         Type = 21
	TypeNodePath           Type = 22
	TypeRID                Type = 23
	TypeObject             Type = 24
	TypeCallable           Type = 25
	TypeSignal             Type = 26
	TypeDictionary         Type = 27
	TypeArray              Type = 28
	TypePackedByteArray    Type = 29
	TypePackedInt32Array   Type = 30
	TypePackedInt64Array   Type = 31
	TypePackedFloat32Array Type = 32
	TypePackedFloat64Array Type = 33
	TypePackedStringArray  Type = 34
	TypePackedVector2Array Type = 35
	TypePackedVector3Array Type = 36
	TypePackedColorArray   Type = 37
	TypePackedVector4Array Type = 38
)

// Type returns the type of the given value.
func (a Any) Type() Type { //gd:typeof
	if a.value == nil {
		return TypeNil
	}
	rtype := reflect.TypeOf(a.value)
	switch a.value.(type) {
	case bool:
		return TypeBool
	case int8, *int16, *int32, *int64, *uint8, *uint16, *uint32:
		return TypeInt
	case *uint, *uint64, *uintptr:
		return TypeRID
	case *float32, *float64:
		return TypeFloat
	case isString:
		return TypeString
	case isPacked[byte]:
		return TypePackedByteArray
	case isPacked[int32]:
		return TypePackedInt32Array
	case isPacked[int64]:
		return TypePackedInt64Array
	case isPacked[float32]:
		return TypePackedFloat32Array
	case isPacked[float64]:
		return TypePackedFloat64Array
	case isPacked[string]:
		return TypePackedStringArray
	case isPacked[Vector2.XY]:
		return TypePackedVector2Array
	case isPacked[Vector3.XYZ]:
		return TypePackedVector3Array
	case isPacked[Color.RGBA]:
		return TypePackedColorArray
	case isPacked[Vector4.XYZW]:
		return TypePackedVector4Array
	case *Vector2.XY:
		return TypeVector2
	case *Vector2i.XY:
		return TypeVector2i
	case *Rect2.PositionSize:
		return TypeRect2
	case *Rect2i.PositionSize:
		return TypeRect2i
	case *Plane.NormalD:
		return TypePlane
	case *Vector3.XYZ:
		return TypeVector3
	case *Vector3i.XYZ:
		return TypeVector3i
	case *Vector4.XYZW:
		return TypeVector4
	case *Vector4i.XYZW:
		return TypeVector4i
	case Transform2D.OriginXY:
		return TypeTransform2D
	case Basis.XYZ:
		return TypeBasis
	case Transform3D.BasisOrigin:
		return TypeTransform3D
	case Projection.XYZW:
		return TypeProjection
	case Color.RGBA:
		return TypeColor
	case Quaternion.IJKX:
		return TypeQuaternion
	case AABB.PositionSize:
		return TypeAABB
	default:
		switch rtype.Kind() {
		case reflect.Slice:
			return TypeArray
		case reflect.Map:
			return TypeDictionary
		case reflect.Struct:
			return TypeDictionary
		case reflect.Func:
			return TypeCallable
		case reflect.Chan:
			return TypeSignal
		}
		return TypeObject
	}
}

// ConvertTo converts the given variant to the given type, using the Type values. This method is generous
// with how it handles types, it can automatically convert between array types, convert numeric Strings
// to int, and converting most things to String.
//
// If the type conversion cannot be done, this method will return the default value for that type, for
// example converting [Rect2.PositionSize] to [Vector2.XY] will always return [Vector2.Zero]. This method
// will never show error messages as long as type is a valid Variant type.
func ConvertTo(to reflect.Type, v any) any { //gd:type_convert
	panic("implement me")
}

// String returns a human-readable name of the given type, using the [Type] values.
func (t Type) String() string { //gd:type_string
	switch t {
	case TypeNil:
		return "Nil"
	case TypeBool:
		return "Bool"
	case TypeInt:
		return "Int"
	case TypeFloat:
		return "Float"
	case TypeString:
		return "String"
	case TypeVector2:
		return "Vector2"
	case TypeVector2i:
		return "Vector2i"
	case TypeRect2:
		return "Rect2"
	case TypeRect2i:
		return "Rect2i"
	case TypeVector3:
		return "Vector3"
	case TypeVector3i:
		return "Vector3i"
	case TypeVector4:
		return "Vector4"
	case TypeVector4i:
		return "Vector4i"
	case TypeTransform2D:
		return "Transform2D"
	case TypePlane:
		return "Plane"
	case TypeQuaternion:
		return "Quaternion"
	case TypeAABB:
		return "AABB"
	case TypeBasis:
		return "Basis"
	case TypeTransform3D:
		return "Transform3D"
	case TypeColor:
		return "Color"
	case TypeNodePath:
		return "NodePath"
	case TypeSignal:
		return "Signal"
	case TypeStringName:
		return "StringName"
	case TypeCallable:
		return "Callable"
	case TypeRID:
		return "RID"
	case TypeObject:
		return "Object"
	case TypeDictionary:
		return "Dictionary"
	case TypeArray:
		return "Array"
	case TypePackedByteArray:
		return "PackedByteArray"
	case TypePackedInt32Array:
		return "PackedInt32Array"
	case TypePackedInt64Array:
		return "PackedInt64Array"
	case TypePackedFloat32Array:
		return "PackedFloat32Array"
	case TypePackedFloat64Array:
		return "PackedFloat64Array"
	case TypePackedStringArray:
		return "PackedStringArray"
	case TypePackedVector2Array:
		return "PackedVector2Array"
	case TypePackedVector3Array:
		return "PackedVector3Array"
	case TypePackedColorArray:
		return "PackedColorArray"
	}
	return "Object"
}

// UnmarshalText converts a formatted string that was returned by [MarshalText] to the original value.
/*func UnmarshalText(s []byte) (any, error) { //gd:str_to_var
	return gd.StrToVar(gd.NewString(string(s))), nil
}

// MarshalText converts a Variant variable to a formatted String that can then be parsed
// using [UnmarshalText].
func MarshalText(v any) ([]byte, error) { //gd:var_to_str
	return []byte(gd.VarToStr(gd.NewVariant(v)).String()), nil
}

// Hash calculates the hash value for a Variant.
func Hash(v any) uint32 { //gd:hash
	return uint32(gd.NewVariant(v).Hash())
}

// Equal compares two Variants and returns true if they are equal.
func Equal(a, b any) bool { //gd:is_same
	return bool(gd.IsSame(gd.NewVariant(a), gd.NewVariant(b)))
}*/

type Operator int

const (
	OpEqual        Operator = 0  // ==
	OpNotEqual     Operator = 1  // !=
	OpLess         Operator = 2  // <
	OpLessEqual    Operator = 3  // <=
	OpGreater      Operator = 4  // >
	OpGreaterEqual Operator = 5  // >=
	OpAdd          Operator = 6  // +
	OpSubtract     Operator = 7  // -
	OpMultiply     Operator = 8  // *
	OpDivide       Operator = 9  // /
	OpNegate       Operator = 10 // -
	OpPositive     Operator = 11 // +
	OpModule       Operator = 12 // %
	OpPower        Operator = 13 // **
	OpShiftLeft    Operator = 14 // <<
	OpShiftRight   Operator = 15 // >>
	OpBitAnd       Operator = 16 // &
	OpBitOr        Operator = 17 // |
	OpBitXor       Operator = 18 // ^
	OpBitNegate    Operator = 19 // ~
	OpAnd          Operator = 20 // &&
	OpOr           Operator = 21 // ||
	OpXor          Operator = 22 // !=
	OpNot          Operator = 23 // !
	OpIn           Operator = 24 // []
)

func Use[T pointers.Generic[T, S], S pointers.Size](ptr T) {
	pointers.Get(ptr)
}

func Debug[T pointers.Generic[T, S], S pointers.Size](ptr T) {
	pointers.Debug(ptr)
}
