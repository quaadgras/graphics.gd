package variant

import (
	"fmt"
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

// Any is like the standard [any] type except it can store the following types without allocating them on the heap:
//   - [bool]
//   - [int][int8][int16][int32][int64][uint][uint8][uint16][uint32][uint64][uintptr]
//   - [Float.X][float32][float64][complex64][complex128]
//   - [string][String][[]byte][[]int32][[]int64][[]float32][[]float64][[]string]
//   - [[]Vector2.XY][[]Vector3.XYZ][[]Color.RGBA][[]Vector4.XYZW]
//   - [Rect2.PositionSize][Rect2i.PositionSize][Plane.NormalD]
//   - [Vector2.XY][Vector2i.XY][Vector3.XYZ][Vector3i.XYZ][Vector4.XYZW][Vector4i.XYZW][Color.RGBA]
type Any struct {
	local complex128
	value any
}

// Nil reference value.
var Nil Any

func load[T any](a Any) T {
	_, ok := a.value.(*T)
	if !ok {
		panic("variant conversion: variant is " + a.Type().String() + ", not " + reflect.TypeFor[T]().String())
	}
	return *(*T)(unsafe.Pointer(&a.local))
}
func loadPacked[T packable](a Any) []T {
	ptr, ok := a.value.(isPacked[T])
	if !ok {
		panic("variant conversion: variant is " + a.Type().String() + ", not " + reflect.TypeFor[T]().String())
	}
	lencap := *(*lencap)(unsafe.Pointer(&a.local))
	return unsafe.Slice(ptr, lencap.cap)[:lencap.len:lencap.cap]
}

func (a Any) Bool() bool { return load[bool](a) }
func (a Any) Int() int {
	rtype := reflect.TypeOf(a.value)
	if rtype.Kind() != reflect.Ptr {
		panic("variant conversion: variant is " + a.Type().String() + ", not " + reflect.TypeFor[int]().String())
	}
	switch rtype.Elem().Kind() {
	case reflect.Int8:
		return int(*(*int8)(unsafe.Pointer(&a.local)))
	case reflect.Int16:
		return int(*(*int16)(unsafe.Pointer(&a.local)))
	case reflect.Int32:
		return int(*(*int32)(unsafe.Pointer(&a.local)))
	case reflect.Int64:
		return int(*(*int64)(unsafe.Pointer(&a.local)))
	case reflect.Int:
		return int(*(*int)(unsafe.Pointer(&a.local)))
	default:
		panic("variant conversion: variant is " + a.Type().String() + ", not " + reflect.TypeFor[int]().String())
	}
}
func (a Any) Int8() int8             { return load[int8](a) }
func (a Any) Int16() int16           { return load[int16](a) }
func (a Any) Int32() int32           { return load[int32](a) }
func (a Any) Int64() int64           { return load[int64](a) }
func (a Any) Uint() uint             { return uint(load[uint](a)) }
func (a Any) Uint8() uint8           { return load[uint8](a) }
func (a Any) Uint16() uint16         { return load[uint16](a) }
func (a Any) Uint32() uint32         { return load[uint32](a) }
func (a Any) Uint64() uint64         { return load[uint64](a) }
func (a Any) Float32() float32       { return load[float32](a) }
func (a Any) Float64() float64       { return load[float64](a) }
func (a Any) Uintptr() uintptr       { return uintptr(load[uintptr](a)) }
func (a Any) Complex64() complex64   { return load[complex64](a) }
func (a Any) Complex128() complex128 { return load[complex128](a) }
func (a Any) String() string {
	ptr, ok := a.value.(isString)
	if !ok {
		return a.toString()
	}
	l := *(*int)(unsafe.Pointer(&a.local))
	return unsafe.String(ptr, l)
}
func (a Any) Vector2() Vector2.XY                  { return load[Vector2.XY](a) }
func (a Any) Vector2i() Vector2i.XY                { return load[Vector2i.XY](a) }
func (a Any) Rect2() Rect2.PositionSize            { return load[Rect2.PositionSize](a) }
func (a Any) Rect2i() Rect2i.PositionSize          { return load[Rect2i.PositionSize](a) }
func (a Any) Vector3() Vector3.XYZ                 { return load[Vector3.XYZ](a) }
func (a Any) Vector3i() Vector3i.XYZ               { return load[Vector3i.XYZ](a) }
func (a Any) Transform2D() Transform2D.OriginXY    { return load[Transform2D.OriginXY](a) }
func (a Any) Vector4() Vector4.XYZW                { return load[Vector4.XYZW](a) }
func (a Any) Vector4i() Vector4i.XYZW              { return load[Vector4i.XYZW](a) }
func (a Any) Plane() Plane.NormalD                 { return load[Plane.NormalD](a) }
func (a Any) Quaternion() Quaternion.IJKX          { return a.value.(Quaternion.IJKX) }
func (a Any) AABB() AABB.PositionSize              { return a.value.(AABB.PositionSize) }
func (a Any) Basis() Basis.XYZ                     { return a.value.(Basis.XYZ) }
func (a Any) Transform3D() Transform3D.BasisOrigin { return a.value.(Transform3D.BasisOrigin) }
func (a Any) Projection() Projection.XYZW          { return a.value.(Projection.XYZW) }
func (a Any) Color() Color.RGBA                    { return load[Color.RGBA](a) }
func (a Any) Bytes() []byte                        { return loadPacked[byte](a) }

func (a Any) toString() string {
	return fmt.Sprint(a.Interface())
}

// Interface returns the value of the variant as an interface{}.
func (a Any) Interface() interface{} {
	rtype := reflect.TypeOf(a.value)
	rvalue := reflect.ValueOf(a.value)
	switch rtype.Kind() {
	case reflect.Pointer:
		if rvalue.Pointer() != 0 {
			switch value := a.value.(type) {
			case isString:
				l := *(*int)(unsafe.Pointer(&a.local))
				return unsafe.String(value, l)
			case isPacked[byte]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[int32]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[int64]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[float32]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[float64]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[string]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[Vector2.XY]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[Vector3.XYZ]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[Color.RGBA]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			case isPacked[Vector4.XYZW]:
				lencap := *(*lencap)(unsafe.Pointer(&a.local))
				return unsafe.Slice(value, lencap.cap)[:lencap.len:lencap.cap]
			}
			lencap := *(*lencap)(unsafe.Pointer(&a.local))
			if lencap.len > 0 {
				return reflect.SliceAt(rtype.Elem(), rvalue.UnsafePointer(), lencap.cap).Slice3(0, lencap.len, lencap.cap).Interface()
			}
			return a.value
		}
		var heap = a.local
		return reflect.NewAt(rtype.Elem(), unsafe.Pointer(&heap)).Elem().Interface()
	default:
		return a.value
	}
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
func ConvertTo(to Type, v any) any { //gd:type_convert
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
	/*Equality operator ([code]==[/code]).*/
	OpEqual Operator = 0
	/*Inequality operator ([code]!=[/code]).*/
	OpNotEqual Operator = 1
	/*Less than operator ([code]<[/code]).*/
	OpLess Operator = 2
	/*Less than or equal operator ([code]<=[/code]).*/
	OpLessEqual Operator = 3
	/*Greater than operator ([code]>[/code]).*/
	OpGreater Operator = 4
	/*Greater than or equal operator ([code]>=[/code]).*/
	OpGreaterEqual Operator = 5
	/*Addition operator ([code]+[/code]).*/
	OpAdd Operator = 6
	/*Subtraction operator ([code]-[/code]).*/
	OpSubtract Operator = 7
	/*Multiplication operator ([code]*[/code]).*/
	OpMultiply Operator = 8
	/*Division operator ([code]/[/code]).*/
	OpDivide Operator = 9
	/*Unary negation operator ([code]-[/code]).*/
	OpNegate Operator = 10
	/*Unary plus operator ([code]+[/code]).*/
	OpPositive Operator = 11
	/*Remainder/modulo operator ([code]%[/code]).*/
	OpModule Operator = 12
	/*Power operator ([code]**[/code]).*/
	OpPower Operator = 13
	/*Left shift operator ([code]<<[/code]).*/
	OpShiftLeft Operator = 14
	/*Right shift operator ([code]>>[/code]).*/
	OpShiftRight Operator = 15
	/*Bitwise AND operator ([code]&[/code]).*/
	OpBitAnd Operator = 16
	/*Bitwise OR operator ([code]|[/code]).*/
	OpBitOr Operator = 17
	/*Bitwise XOR operator ([code]^[/code]).*/
	OpBitXor Operator = 18
	/*Bitwise NOT operator ([code]~[/code]).*/
	OpBitNegate Operator = 19
	/*Logical AND operator ([code]and[/code] or [code]&&[/code]).*/
	OpAnd Operator = 20
	/*Logical OR operator ([code]or[/code] or [code]||[/code]).*/
	OpOr Operator = 21
	/*Logical XOR operator (not implemented in GDScript).*/
	OpXor Operator = 22
	/*Logical NOT operator ([code]not[/code] or [code]![/code]).*/
	OpNot Operator = 23
	/*Logical IN operator ([code]in[/code]).*/
	OpIn Operator = 24
	/*Represents the size of the [enum Variant.Operator] enum.*/
	OpMax Operator = 25
)

func Use[T pointers.Generic[T, S], S pointers.Size](ptr T) {
	pointers.Get(ptr)
}

func Debug[T pointers.Generic[T, S], S pointers.Size](ptr T) {
	pointers.Debug(ptr)
}
