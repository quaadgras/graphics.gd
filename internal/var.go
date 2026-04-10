//go:build !generate

package gd

import (
	"reflect"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"

	float "graphics.gd/variant/Float"
	rid "graphics.gd/variant/RID"
)

type (
	Void     = struct{}
	Int      = int64
	Float    = float64
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

type RID = rid.Any

func (c Callable) Free() {
	ptr, ok := pointers.End(c)
	if !ok {
		return
	}
	gdunsafe.Free(gdunsafe.Callable(ptr))
}

func (v Variant) Free() {
	ptr, ok := pointers.End(v)
	if !ok {
		return
	}
	gdunsafe.Variant(ptr).Free()
}

type Iterator struct {
	self Variant
	iter iterator
}

type iterator pointers.Type[iterator, gdextension.Iterator]

func (iter iterator) Free() {
	if ptr, ok := pointers.End(iter); ok {
		gdunsafe.Variant(ptr).Free()
	}
}

func (iter Iterator) Next() bool {
	var raw = gdunsafe.Iterator(pointers.Get(iter.iter))
	next, _ := raw.Next()
	pointers.Set(iter.iter, raw)
	return next
}

func (iter Iterator) Value() Variant {
	var raw gdextension.Variant
	raw, err := gdunsafe.Iterator(pointers.Get(iter.self)).Value()
	if err != (gdextension.CallError{}) {
		panic("failed to get iterator value")
	}
	return pointers.New[Variant](raw)
}

func variantTypeFromName(s string) (variant.Type, reflect.Type) {
	switch s {
	case "Nil":
		return variant.TypeNil, nil
	case "bool", "Bool":
		return variant.TypeBool, reflect.TypeFor[bool]()
	case "int", "Int":
		return variant.TypeInt, reflect.TypeFor[int64]()
	case "float", "Float":
		return variant.TypeFloat, reflect.TypeFor[Float]()
	case "String":
		return variant.TypeString, reflect.TypeFor[String]()
	case "Vector2":
		return variant.TypeVector2, reflect.TypeFor[Vector2]()
	case "Vector2i":
		return variant.TypeVector2i, reflect.TypeFor[Vector2i]()
	case "Rect2":
		return variant.TypeRect2, reflect.TypeFor[Rect2]()
	case "Rect2i":
		return variant.TypeRect2i, reflect.TypeFor[Rect2i]()
	case "Vector3":
		return variant.TypeVector3, reflect.TypeFor[Vector3]()
	case "Vector3i":
		return variant.TypeVector3i, reflect.TypeFor[Vector3i]()
	case "Transform2D":
		return variant.TypeTransform2D, reflect.TypeFor[Transform2D]()
	case "Vector4":
		return variant.TypeVector4, reflect.TypeFor[Vector4]()
	case "Vector4i":
		return variant.TypeVector4i, reflect.TypeFor[Vector4i]()
	case "Plane":
		return variant.TypePlane, reflect.TypeFor[Plane]()
	case "Quaternion":
		return variant.TypeQuaternion, reflect.TypeFor[Quaternion]()
	case "AABB":
		return variant.TypeAABB, reflect.TypeFor[AABB]()
	case "Basis":
		return variant.TypeBasis, reflect.TypeFor[Basis]()
	case "Transform3D":
		return variant.TypeTransform3D, reflect.TypeFor[Transform3D]()
	case "Projection":
		return variant.TypeProjection, reflect.TypeFor[Projection]()
	case "Color":
		return variant.TypeColor, reflect.TypeFor[Color]()
	case "StringName":
		return variant.TypeStringName, reflect.TypeFor[StringName]()
	case "NodePath":
		return variant.TypeNodePath, reflect.TypeFor[NodePath]()
	case "RID":
		return variant.TypeRID, reflect.TypeFor[RID]()
	case "Object":
		return variant.TypeObject, reflect.TypeFor[uintptr]()
	case "Callable":
		return variant.TypeCallable, reflect.TypeFor[Callable]()
	case "Signal":
		return variant.TypeSignal, reflect.TypeFor[Signal]()
	case "Dictionary":
		return variant.TypeDictionary, reflect.TypeFor[Dictionary]()
	case "Array":
		return variant.TypeArray, reflect.TypeFor[Array]()
	case "PackedByteArray":
		return variant.TypePackedByteArray, reflect.TypeFor[PackedByteArray]()
	case "PackedInt32Array":
		return variant.TypePackedInt32Array, reflect.TypeFor[PackedInt32Array]()
	case "PackedInt64Array":
		return variant.TypePackedInt64Array, reflect.TypeFor[PackedInt64Array]()
	case "PackedFloat32Array":
		return variant.TypePackedFloat32Array, reflect.TypeFor[PackedFloat32Array]()
	case "PackedFloat64Array":
		return variant.TypePackedFloat64Array, reflect.TypeFor[PackedFloat64Array]()
	case "PackedStringArray":
		return variant.TypePackedStringArray, reflect.TypeFor[PackedStringArray]()
	case "PackedVector2Array":
		return variant.TypePackedVector2Array, reflect.TypeFor[PackedVector2Array]()
	case "PackedVector3Array":
		return variant.TypePackedVector3Array, reflect.TypeFor[PackedVector3Array]()
	case "PackedVector4Array":
		return variant.TypePackedVector4Array, reflect.TypeFor[PackedVector4Array]()
	case "PackedColorArray":
		return variant.TypePackedColorArray, reflect.TypeFor[PackedColorArray]()
	default:
		panic("gdextension.variantTypeFromName: unknown type " + s)
	}
}
