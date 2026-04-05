//go:build !generate

package gd

import (
	"reflect"
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/noescape"
	"graphics.gd/internal/pointers"

	float "graphics.gd/variant/Float"
	rid "graphics.gd/variant/RID"
)

type (
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
	noescape.Free(gdextension.TypeCallable, &ptr)
}

func (s Variant) Free() {
	ptr, ok := pointers.End(s)
	if !ok {
		return
	}
	gdunsafe.Variant(ptr).UnsafeFree()
}

type Iterator struct {
	self Variant
	iter iterator
}

type iterator pointers.Type[iterator, gdextension.Iterator]

func (iter iterator) Free() {
	if ptr, ok := pointers.End(iter); ok {
		gdunsafe.Variant(ptr).UnsafeFree()
	}
}

func (iter Iterator) Next() bool {
	var err gdextension.CallError
	var raw = pointers.Get(iter.iter)
	next := gdunsafe.Variant(pointers.Get(iter.self)).IteratorNext(unsafe.Pointer(&raw), unsafe.Pointer(&err))
	pointers.Set(iter.iter, raw)
	return next
}

func (iter Iterator) Value() Variant {
	var err gdextension.CallError
	var raw gdextension.Variant
	gdunsafe.Variant(pointers.Get(iter.self)).IteratorLoad(gdunsafe.Variant(pointers.Get(iter.iter)), unsafe.Pointer(&raw), unsafe.Pointer(&err))
	if err.Type != 0 {
		panic("failed to get iterator value")
	}
	return pointers.New[Variant](raw)
}

func variantTypeFromName(s string) (gdextension.VariantType, reflect.Type) {
	switch s {
	case "Nil":
		return gdextension.TypeNil, nil
	case "bool", "Bool":
		return gdextension.TypeBool, reflect.TypeFor[bool]()
	case "int", "Int":
		return gdextension.TypeInt, reflect.TypeFor[int64]()
	case "float", "Float":
		return gdextension.TypeFloat, reflect.TypeFor[Float]()
	case "String":
		return gdextension.TypeString, reflect.TypeFor[String]()
	case "Vector2":
		return gdextension.TypeVector2, reflect.TypeFor[Vector2]()
	case "Vector2i":
		return gdextension.TypeVector2i, reflect.TypeFor[Vector2i]()
	case "Rect2":
		return gdextension.TypeRect2, reflect.TypeFor[Rect2]()
	case "Rect2i":
		return gdextension.TypeRect2i, reflect.TypeFor[Rect2i]()
	case "Vector3":
		return gdextension.TypeVector3, reflect.TypeFor[Vector3]()
	case "Vector3i":
		return gdextension.TypeVector3i, reflect.TypeFor[Vector3i]()
	case "Transform2D":
		return gdextension.TypeTransform2D, reflect.TypeFor[Transform2D]()
	case "Vector4":
		return gdextension.TypeVector4, reflect.TypeFor[Vector4]()
	case "Vector4i":
		return gdextension.TypeVector4i, reflect.TypeFor[Vector4i]()
	case "Plane":
		return gdextension.TypePlane, reflect.TypeFor[Plane]()
	case "Quaternion":
		return gdextension.TypeQuaternion, reflect.TypeFor[Quaternion]()
	case "AABB":
		return gdextension.TypeAABB, reflect.TypeFor[AABB]()
	case "Basis":
		return gdextension.TypeBasis, reflect.TypeFor[Basis]()
	case "Transform3D":
		return gdextension.TypeTransform3D, reflect.TypeFor[Transform3D]()
	case "Projection":
		return gdextension.TypeProjection, reflect.TypeFor[Projection]()
	case "Color":
		return gdextension.TypeColor, reflect.TypeFor[Color]()
	case "StringName":
		return gdextension.TypeStringName, reflect.TypeFor[StringName]()
	case "NodePath":
		return gdextension.TypeNodePath, reflect.TypeFor[NodePath]()
	case "RID":
		return gdextension.TypeRID, reflect.TypeFor[RID]()
	case "Object":
		return gdextension.TypeObject, reflect.TypeFor[uintptr]()
	case "Callable":
		return gdextension.TypeCallable, reflect.TypeFor[Callable]()
	case "Signal":
		return gdextension.TypeSignal, reflect.TypeFor[Signal]()
	case "Dictionary":
		return gdextension.TypeDictionary, reflect.TypeFor[Dictionary]()
	case "Array":
		return gdextension.TypeArray, reflect.TypeFor[Array]()
	case "PackedByteArray":
		return gdextension.TypePackedByteArray, reflect.TypeFor[PackedByteArray]()
	case "PackedInt32Array":
		return gdextension.TypePackedInt32Array, reflect.TypeFor[PackedInt32Array]()
	case "PackedInt64Array":
		return gdextension.TypePackedInt64Array, reflect.TypeFor[PackedInt64Array]()
	case "PackedFloat32Array":
		return gdextension.TypePackedFloat32Array, reflect.TypeFor[PackedFloat32Array]()
	case "PackedFloat64Array":
		return gdextension.TypePackedFloat64Array, reflect.TypeFor[PackedFloat64Array]()
	case "PackedStringArray":
		return gdextension.TypePackedStringArray, reflect.TypeFor[PackedStringArray]()
	case "PackedVector2Array":
		return gdextension.TypePackedVector2Array, reflect.TypeFor[PackedVector2Array]()
	case "PackedVector3Array":
		return gdextension.TypePackedVector3Array, reflect.TypeFor[PackedVector3Array]()
	case "PackedVector4Array":
		return gdextension.TypePackedVector4Array, reflect.TypeFor[PackedVector4Array]()
	case "PackedColorArray":
		return gdextension.TypePackedColorArray, reflect.TypeFor[PackedColorArray]()
	default:
		panic("gdextension.variantTypeFromName: unknown type " + s)
	}
}
