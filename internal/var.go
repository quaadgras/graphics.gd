//go:build !generate

package gd

import (
	"reflect"

	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"

	float "graphics.gd/variant/Float"
	gdVector3 "graphics.gd/variant/Vector3"
)

type (
	Bool     bool
	Float    float64
	Int      int64
	Vector2  = struct{ X, Y float.X }
	Vector2i = struct{ X, Y int32 }
	Rect2    = struct {
		Position Vector2
		Size     Vector2
	}
	Rect2i = struct {
		Position Vector2i
		Size     Vector2i
	}
	Vector3     = struct{ X, Y, Z float.X }
	Vector3i    = struct{ X, Y, Z int32 }
	Transform2D = struct {
		X, Y   Vector2
		Origin Vector2
	}
	Vector4  = struct{ X, Y, Z, W float.X }
	Vector4i = struct{ X, Y, Z, W int32 }
	Plane    = struct {
		Normal Vector3
		D      float.X
	}
	Quaternion = struct{ I, J, K, X float.X }
	AABB       = struct {
		Position Vector3
		Size     Vector3
	}
	Basis       = struct{ X, Y, Z Vector3 }
	Transform3D = struct {
		Basis  Basis
		Origin Vector3
	}
	Projection = struct{ X, Y, Z, W Vector4 }
)

type Color = struct{ R, G, B, A float.X }

type RID uint64

func (c Callable) Free() {
	ptr, ok := pointers.End(c)
	if !ok {
		return
	}
	var frame = callframe.New()
	Global.typeset.destruct.Callable(callframe.Arg(frame, ptr).Addr())
	frame.Free()
}

func (s Variant) Free() {
	Global.Variants.Destroy(s)
}

type Iterator struct {
	self Variant
	iter Variant
}

func (iter Iterator) Next() bool {
	return Global.Variants.IteratorNext(iter.self, iter.iter)
}

func (iter Iterator) Value() Variant {
	val, ok := Global.Variants.IteratorGet(iter.self, iter.iter)
	if !ok {
		panic("failed to get iterator value")
	}
	return val
}

func variantTypeFromName(s string) (VariantType, reflect.Type) {
	switch s {
	case "Nil":
		return TypeNil, nil
	case "bool", "Bool":
		return TypeBool, reflect.TypeOf(false)
	case "int", "Int":
		return TypeInt, reflect.TypeOf(int64(0))
	case "float", "Float":
		return TypeFloat, reflect.TypeOf(Float(0))
	case "String":
		return TypeString, reflect.TypeOf(String{})
	case "Vector2":
		return TypeVector2, reflect.TypeOf(Vector2{})
	case "Vector2i":
		return TypeVector2i, reflect.TypeOf(Vector2i{})
	case "Rect2":
		return TypeRect2, reflect.TypeOf(Rect2{})
	case "Rect2i":
		return TypeRect2i, reflect.TypeOf(Rect2i{})
	case "Vector3":
		return TypeVector3, reflect.TypeOf(Vector3{})
	case "Vector3i":
		return TypeVector3i, reflect.TypeOf(Vector3i{})
	case "Transform2D":
		return TypeTransform2D, reflect.TypeOf(Transform2D{})
	case "Vector4":
		return TypeVector4, reflect.TypeOf(Vector4{})
	case "Vector4i":
		return TypeVector4i, reflect.TypeOf(Vector4i{})
	case "Plane":
		return TypePlane, reflect.TypeOf(Plane{})
	case "Quaternion":
		return TypeQuaternion, reflect.TypeOf(Quaternion{})
	case "AABB":
		return TypeAABB, reflect.TypeOf(AABB{})
	case "Basis":
		return TypeBasis, reflect.TypeOf(Basis{})
	case "Transform3D":
		return TypeTransform3D, reflect.TypeOf(Transform3D{})
	case "Projection":
		return TypeProjection, reflect.TypeOf(Projection{})
	case "Color":
		return TypeColor, reflect.TypeOf(Color{})
	case "StringName":
		return TypeStringName, reflect.TypeOf(StringName{})
	case "NodePath":
		return TypeNodePath, reflect.TypeOf(NodePath{})
	case "RID":
		return TypeRID, reflect.TypeOf(RID(0))
	case "Object":
		return TypeObject, reflect.TypeOf(uintptr(0))
	case "Callable":
		return TypeCallable, reflect.TypeOf(Callable{})
	case "Signal":
		return TypeSignal, reflect.TypeOf(Signal{})
	case "Dictionary":
		return TypeDictionary, reflect.TypeOf(Dictionary{})
	case "Array":
		return TypeArray, reflect.TypeOf(Array{})
	case "PackedByteArray":
		return TypePackedByteArray, reflect.TypeOf(PackedByteArray{})
	case "PackedInt32Array":
		return TypePackedInt32Array, reflect.TypeOf(PackedInt32Array{})
	case "PackedInt64Array":
		return TypePackedInt64Array, reflect.TypeOf(PackedInt64Array{})
	case "PackedFloat32Array":
		return TypePackedFloat32Array, reflect.TypeOf(PackedFloat32Array{})
	case "PackedFloat64Array":
		return TypePackedFloat64Array, reflect.TypeOf(PackedFloat64Array{})
	case "PackedStringArray":
		return TypePackedStringArray, reflect.TypeOf(PackedStringArray{})
	case "PackedVector2Array":
		return TypePackedVector2Array, reflect.TypeOf(PackedVector2Array{})
	case "PackedVector3Array":
		return TypePackedVector3Array, reflect.TypeOf(PackedVector3Array{})
	case "PackedVector4Array":
		return TypePackedVector4Array, reflect.TypeOf(PackedVector4Array{})
	case "PackedColorArray":
		return TypePackedColorArray, reflect.TypeOf(PackedColorArray{})
	default:
		panic("gdextension.variantTypeFromName: unknown type " + s)
	}
}

func operatoTypeFromName(name string) Operator {
	switch name {
	case "Equals":
		return Equal
	case "NotEqual":
		return NotEqual
	case "Less":
		return Less
	case "LessEqual":
		return LessEqual
	case "Greater":
		return Greater
	case "GreaterEqual":
		return GreaterEqual
	case "Add":
		return Add
	case "Subtract":
		return Subtract
	case "Multiply":
		return Multiply
	case "Divide":
		return Divide
	case "Negate":
		return Negate
	case "Module":
		return Module
	case "Power":
		return Power
	case "ShiftLeft":
		return ShiftLeft
	case "ShiftRight":
		return ShiftRight
	case "BitAnd":
		return BitAnd
	case "BitOr":
		return BitOr
	case "BitXor":
		return BitXor
	case "BitNegate":
		return BitNegate
	case "And":
		return LogicalAnd
	case "Or":
		return LogicalOr
	case "Xor":
		return LogicalXor
	case "Not":
		return LogicalNegate
	case "In":
		return In
	default:
		panic("gdextension.operatoTypeFromName: unknown operator " + name)
	}
}

type Vector3Axis = gdVector3.Axis

const (
	/*Variable is [code]null[/code].*/
	TypeNil VariantType = 0
	/*Variable is of type [bool].*/
	TypeBool VariantType = 1
	/*Variable is of type [int].*/
	TypeInt VariantType = 2
	/*Variable is of type [float].*/
	TypeFloat VariantType = 3
	/*Variable is of type [String].*/
	TypeString VariantType = 4
	/*Variable is of type [Vector2].*/
	TypeVector2 VariantType = 5
	/*Variable is of type [Vector2i].*/
	TypeVector2i VariantType = 6
	/*Variable is of type [Rect2].*/
	TypeRect2 VariantType = 7
	/*Variable is of type [Rect2i].*/
	TypeRect2i VariantType = 8
	/*Variable is of type [Vector3].*/
	TypeVector3 VariantType = 9
	/*Variable is of type [Vector3i].*/
	TypeVector3i VariantType = 10
	/*Variable is of type [Transform2D].*/
	TypeTransform2D VariantType = 11
	/*Variable is of type [Vector4].*/
	TypeVector4 VariantType = 12
	/*Variable is of type [Vector4i].*/
	TypeVector4i VariantType = 13
	/*Variable is of type [Plane].*/
	TypePlane VariantType = 14
	/*Variable is of type [Quaternion].*/
	TypeQuaternion VariantType = 15
	/*Variable is of type [AABB].*/
	TypeAABB VariantType = 16
	/*Variable is of type [Basis].*/
	TypeBasis VariantType = 17
	/*Variable is of type [Transform3D].*/
	TypeTransform3D VariantType = 18
	/*Variable is of type [Projection].*/
	TypeProjection VariantType = 19
	/*Variable is of type [Color].*/
	TypeColor VariantType = 20
	/*Variable is of type [StringName].*/
	TypeStringName VariantType = 21
	/*Variable is of type [NodePath].*/
	TypeNodePath VariantType = 22
	/*Variable is of type [RID].*/
	TypeRID VariantType = 23
	/*Variable is of type [Object].*/
	TypeObject VariantType = 24
	/*Variable is of type [Callable].*/
	TypeCallable VariantType = 25
	/*Variable is of type [Signal].*/
	TypeSignal VariantType = 26
	/*Variable is of type [Dictionary].*/
	TypeDictionary VariantType = 27
	/*Variable is of type [Array].*/
	TypeArray VariantType = 28
	/*Variable is of type [PackedByteArray].*/
	TypePackedByteArray VariantType = 29
	/*Variable is of type [PackedInt32Array].*/
	TypePackedInt32Array VariantType = 30
	/*Variable is of type [PackedInt64Array].*/
	TypePackedInt64Array VariantType = 31
	/*Variable is of type [PackedFloat32Array].*/
	TypePackedFloat32Array VariantType = 32
	/*Variable is of type [PackedFloat64Array].*/
	TypePackedFloat64Array VariantType = 33
	/*Variable is of type [PackedStringArray].*/
	TypePackedStringArray VariantType = 34
	/*Variable is of type [PackedVector2Array].*/
	TypePackedVector2Array VariantType = 35
	/*Variable is of type [PackedVector3Array].*/
	TypePackedVector3Array VariantType = 36
	/*Variable is of type [PackedColorArray].*/
	TypePackedColorArray VariantType = 37
	/*Variable is of type [PackedVector4Array].*/
	TypePackedVector4Array VariantType = 38
	/*Represents the size of the [enum Variant.Type] enum.*/
	TypeMax VariantType = 39
)
