package variant

import (
	"reflect"

	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/variant/AABB"
	"grow.graphics/gd/variant/Basis"
	"grow.graphics/gd/variant/Path"
	"grow.graphics/gd/variant/Plane"
	"grow.graphics/gd/variant/Projection"
	"grow.graphics/gd/variant/Quaternion"
	"grow.graphics/gd/variant/Rect2"
	"grow.graphics/gd/variant/Rect2i"
	"grow.graphics/gd/variant/Transform2D"
	"grow.graphics/gd/variant/Transform3D"
	"grow.graphics/gd/variant/Vector2"
	"grow.graphics/gd/variant/Vector2i"
	"grow.graphics/gd/variant/Vector3"
	"grow.graphics/gd/variant/Vector3i"
	"grow.graphics/gd/variant/Vector4"
	"grow.graphics/gd/variant/Vector4i"
)

func New(val any) gd.Variant {
	return gd.NewVariant(val)
}

// Call attempts to call the given method on the given variant value with the given arguments.
func Call(val any, name string, args ...any) any {
	arguments := make([]gd.Variant, len(args))
	for i, arg := range args {
		arguments[i] = gd.NewVariant(arg)
	}
	result, err := gd.NewVariant(val).Call(gd.NewStringName(name), arguments...)
	if err != nil {
		panic(err)
	}
	return result.Interface()
}

type Type int

const (
	/*Variable is [code]null[/code].*/
	TypeNil Type = 0
	/*Variable is of type [bool].*/
	TypeBool Type = 1
	/*Variable is of type [int].*/
	TypeInt Type = 2
	/*Variable is of type [float].*/
	TypeFloat Type = 3
	/*Variable is of type [String].*/
	TypeString Type = 4
	/*Variable is of type [Vector2].*/
	TypeVector2 Type = 5
	/*Variable is of type [Vector2i].*/
	TypeVector2i Type = 6
	/*Variable is of type [Rect2].*/
	TypeRect2 Type = 7
	/*Variable is of type [Rect2i].*/
	TypeRect2i Type = 8
	/*Variable is of type [Vector3].*/
	TypeVector3 Type = 9
	/*Variable is of type [Vector3i].*/
	TypeVector3i Type = 10
	/*Variable is of type [Transform2D].*/
	TypeTransform2D Type = 11
	/*Variable is of type [Vector4].*/
	TypeVector4 Type = 12
	/*Variable is of type [Vector4i].*/
	TypeVector4i Type = 13
	/*Variable is of type [Plane].*/
	TypePlane Type = 14
	/*Variable is of type [Quaternion].*/
	TypeQuaternion Type = 15
	/*Variable is of type [AABB].*/
	TypeAABB Type = 16
	/*Variable is of type [Basis].*/
	TypeBasis Type = 17
	/*Variable is of type [Transform3D].*/
	TypeTransform3D Type = 18
	/*Variable is of type [Projection].*/
	TypeProjection Type = 19
	/*Variable is of type [Color].*/
	TypeColor Type = 20
	/*Variable is of type [StringName].*/
	TypeStringName Type = 21
	/*Variable is of type [NodePath].*/
	TypeNodePath Type = 22
	/*Variable is of type [RID].*/
	TypeRID Type = 23
	/*Variable is of type [Object].*/
	TypeObject Type = 24
	/*Variable is of type [Callable].*/
	TypeCallable Type = 25
	/*Variable is of type [Signal].*/
	TypeSignal Type = 26
	/*Variable is of type [Dictionary].*/
	TypeDictionary Type = 27
	/*Variable is of type [Array].*/
	TypeArray Type = 28
	/*Variable is of type [PackedByteArray].*/
	TypePackedByteArray Type = 29
	/*Variable is of type [PackedInt32Array].*/
	TypePackedInt32Array Type = 30
	/*Variable is of type [PackedInt64Array].*/
	TypePackedInt64Array Type = 31
	/*Variable is of type [PackedFloat32Array].*/
	TypePackedFloat32Array Type = 32
	/*Variable is of type [PackedFloat64Array].*/
	TypePackedFloat64Array Type = 33
	/*Variable is of type [PackedStringArray].*/
	TypePackedStringArray Type = 34
	/*Variable is of type [PackedVector2Array].*/
	TypePackedVector2Array Type = 35
	/*Variable is of type [PackedVector3Array].*/
	TypePackedVector3Array Type = 36
	/*Variable is of type [PackedColorArray].*/
	TypePackedColorArray Type = 37
	/*Variable is of type [PackedVector4Array].*/
	TypePackedVector4Array Type = 38
	/*Represents the size of the [enum Variant.Type] enum.*/
	TypeMax Type = 39
)

// TypeOf returns the type of the given value.
func TypeOf(val any) Type { //gd:typeof
	if val == nil {
		return TypeNil
	}
	rtype := reflect.TypeOf(val)
	switch rtype.Kind() {
	case reflect.Bool:
		return TypeBool
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return TypeInt
	case reflect.Float32, reflect.Float64:
		return TypeFloat
	case reflect.String:
		if rtype == reflect.TypeFor[Path.String]() {
			return TypeNodePath
		}
		return TypeString
	case reflect.Map:
		return TypeDictionary
	case reflect.Slice:
		switch rtype.Elem().Kind() {
		case reflect.Uint8:
			return TypePackedByteArray
		case reflect.Int32:
			return TypePackedInt32Array
		case reflect.Int64:
			return TypePackedInt64Array
		case reflect.Float32:
			return TypePackedFloat32Array
		case reflect.Float64:
			return TypePackedFloat64Array
		case reflect.String:
			return TypePackedStringArray
		case reflect.Struct:
			switch {
			case rtype.Elem().ConvertibleTo(reflect.TypeFor[Vector2.XY]()):
				return TypePackedVector2Array
			case rtype.Elem().ConvertibleTo(reflect.TypeFor[Vector3.XYZ]()):
				return TypePackedVector3Array
			case rtype.Elem().ConvertibleTo(reflect.TypeFor[gd.Color]()):
				return TypePackedColorArray
			case rtype.Elem().ConvertibleTo(reflect.TypeFor[Vector4.XYZW]()):
				return TypePackedVector4Array
			}
		}
		return TypeArray
	case reflect.Struct:
		switch {
		case rtype.ConvertibleTo(reflect.TypeFor[Vector2.XY]()):
			return TypeVector2
		case rtype.ConvertibleTo(reflect.TypeFor[Vector2i.XY]()):
			return TypeVector2i
		case rtype.ConvertibleTo(reflect.TypeFor[Rect2.PositionSize]()):
			return TypeRect2
		case rtype.ConvertibleTo(reflect.TypeFor[Rect2i.PositionSize]()):
			return TypeRect2i
		case rtype.ConvertibleTo(reflect.TypeFor[Vector3.XYZ]()):
			return TypeVector3
		case rtype.ConvertibleTo(reflect.TypeFor[Vector3i.XYZ]()):
			return TypeVector3i
		case rtype.ConvertibleTo(reflect.TypeFor[Transform2D.OriginXY]()):
			return TypeTransform2D
		case rtype.ConvertibleTo(reflect.TypeFor[Vector4.XYZW]()):
			return TypeVector4
		case rtype.ConvertibleTo(reflect.TypeFor[Vector4i.XYZW]()):
			return TypeVector4i
		case rtype.ConvertibleTo(reflect.TypeFor[Plane.NormalD]()):
			return TypePlane
		case rtype.ConvertibleTo(reflect.TypeFor[Quaternion.IJKX]()):
			return TypeQuaternion
		case rtype.ConvertibleTo(reflect.TypeFor[AABB.PositionSize]()):
			return TypeAABB
		case rtype.ConvertibleTo(reflect.TypeFor[Basis.XYZ]()):
			return TypeBasis
		case rtype.ConvertibleTo(reflect.TypeFor[Transform3D.BasisOrigin]()):
			return TypeTransform3D
		case rtype.ConvertibleTo(reflect.TypeFor[Projection.XYZW]()):
			return TypeProjection
		case rtype.ConvertibleTo(reflect.TypeFor[gd.Color]()):
			return TypeColor
		}
		return TypeDictionary
	}
	return TypeObject
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
func UnmarshalText(s []byte) (any, error) { //gd:str_to_var
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
}

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
