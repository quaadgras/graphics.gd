//go:build !generate

package gd

import (
	"reflect"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"
	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/RID"
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
		return variant.TypeFloat, reflect.TypeFor[Float.X]()
	case "String":
		return variant.TypeString, reflect.TypeFor[gdunsafe.String]()
	case "Vector2":
		return variant.TypeVector2, reflect.TypeFor[Vector2.XY]()
	case "Vector2i":
		return variant.TypeVector2i, reflect.TypeFor[Vector2i.XY]()
	case "Rect2":
		return variant.TypeRect2, reflect.TypeFor[Rect2.PositionSize]()
	case "Rect2i":
		return variant.TypeRect2i, reflect.TypeFor[Rect2i.PositionSize]()
	case "Vector3":
		return variant.TypeVector3, reflect.TypeFor[Vector3.XYZ]()
	case "Vector3i":
		return variant.TypeVector3i, reflect.TypeFor[Vector3i.XYZ]()
	case "Transform2D":
		return variant.TypeTransform2D, reflect.TypeFor[Transform2D.OriginXY]()
	case "Vector4":
		return variant.TypeVector4, reflect.TypeFor[Vector4.XYZW]()
	case "Vector4i":
		return variant.TypeVector4i, reflect.TypeFor[Vector4i.XYZW]()
	case "Plane":
		return variant.TypePlane, reflect.TypeFor[Plane.NormalD]()
	case "Quaternion":
		return variant.TypeQuaternion, reflect.TypeFor[Quaternion.IJKX]()
	case "AABB":
		return variant.TypeAABB, reflect.TypeFor[AABB.PositionSize]()
	case "Basis":
		return variant.TypeBasis, reflect.TypeFor[Basis.XYZ]()
	case "Transform3D":
		return variant.TypeTransform3D, reflect.TypeFor[Transform3D.BasisOrigin]()
	case "Projection":
		return variant.TypeProjection, reflect.TypeFor[Projection.XYZW]()
	case "Color":
		return variant.TypeColor, reflect.TypeFor[Color.RGBA]()
	case "StringName":
		return variant.TypeStringName, reflect.TypeFor[gdunsafe.StringName]()
	case "NodePath":
		return variant.TypeNodePath, reflect.TypeFor[gdunsafe.NodePath]()
	case "RID":
		return variant.TypeRID, reflect.TypeFor[RID.Any]()
	case "Object":
		return variant.TypeObject, reflect.TypeFor[uintptr]()
	case "Callable":
		return variant.TypeCallable, reflect.TypeFor[gdunsafe.Callable]()
	case "Signal":
		return variant.TypeSignal, reflect.TypeFor[gdunsafe.Signal]()
	case "Dictionary":
		return variant.TypeDictionary, reflect.TypeFor[gdunsafe.Dictionary]()
	case "Array":
		return variant.TypeArray, reflect.TypeFor[gdunsafe.Array]()
	case "PackedByteArray":
		return variant.TypePackedByteArray, reflect.TypeFor[gdunsafe.PackedArray[byte]]()
	case "PackedInt32Array":
		return variant.TypePackedInt32Array, reflect.TypeFor[gdunsafe.PackedArray[int32]]()
	case "PackedInt64Array":
		return variant.TypePackedInt64Array, reflect.TypeFor[gdunsafe.PackedArray[int64]]()
	case "PackedFloat32Array":
		return variant.TypePackedFloat32Array, reflect.TypeFor[gdunsafe.PackedArray[float32]]()
	case "PackedFloat64Array":
		return variant.TypePackedFloat64Array, reflect.TypeFor[gdunsafe.PackedArray[float64]]()
	case "PackedStringArray":
		return variant.TypePackedStringArray, reflect.TypeFor[gdunsafe.PackedArray[gdunsafe.String]]()
	case "PackedVector2Array":
		return variant.TypePackedVector2Array, reflect.TypeFor[gdunsafe.PackedArray[Vector2.XY]]()
	case "PackedVector3Array":
		return variant.TypePackedVector3Array, reflect.TypeFor[gdunsafe.PackedArray[Vector3.XYZ]]()
	case "PackedVector4Array":
		return variant.TypePackedVector4Array, reflect.TypeFor[gdunsafe.PackedArray[Vector4.XYZW]]()
	case "PackedColorArray":
		return variant.TypePackedColorArray, reflect.TypeFor[gdunsafe.PackedArray[Color.RGBA]]()
	default:
		panic("gdextension.variantTypeFromName: unknown type " + s)
	}
}
