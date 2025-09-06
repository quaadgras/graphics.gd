//go:build !generate

package gd

import (
	"reflect"

	"graphics.gd/internal/gdextension"
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
	gdextension.Free(gdextension.TypeCallable, &ptr)
}

func (s Variant) Free() {
	ptr, ok := pointers.End(s)
	if !ok {
		return
	}
	gdextension.Host.Variants.Unsafe.Free(ptr)
}

type Iterator struct {
	self Variant
	iter iterator
}

type iterator pointers.Type[iterator, gdextension.Iterator]

func (iter iterator) Free() {
	if ptr, ok := pointers.End(iter); ok {
		gdextension.Host.Variants.Unsafe.Free(gdextension.Variant(ptr))
	}
}

func (iter Iterator) Next() bool {
	var err gdextension.CallError
	var raw = pointers.Get(iter.iter)
	next := gdextension.Host.Iterators.Next(pointers.Get(iter.self), gdextension.CallMutates[gdextension.Iterator](&raw), gdextension.CallReturns[gdextension.CallError](&err))
	pointers.Set(iter.iter, raw)
	return next
}

func (iter Iterator) Value() Variant {
	var err gdextension.CallError
	var raw gdextension.Variant
	gdextension.Host.Iterators.Load(pointers.Get(iter.self), pointers.Get(iter.iter), gdextension.CallReturns[gdextension.Variant](&raw), gdextension.CallReturns[gdextension.CallError](&err))
	if err.Type != 0 {
		panic("failed to get iterator value")
	}
	return pointers.New[Variant]([3]uint64(raw))
}

func variantTypeFromName(s string) (gdextension.VariantType, reflect.Type) {
	switch s {
	case "Nil":
		return gdextension.TypeNil, nil
	case "bool", "Bool":
		return gdextension.TypeBool, reflect.TypeOf(false)
	case "int", "Int":
		return gdextension.TypeInt, reflect.TypeOf(int64(0))
	case "float", "Float":
		return gdextension.TypeFloat, reflect.TypeOf(Float(0))
	case "String":
		return gdextension.TypeString, reflect.TypeOf(String{})
	case "Vector2":
		return gdextension.TypeVector2, reflect.TypeOf(Vector2{})
	case "Vector2i":
		return gdextension.TypeVector2i, reflect.TypeOf(Vector2i{})
	case "Rect2":
		return gdextension.TypeRect2, reflect.TypeOf(Rect2{})
	case "Rect2i":
		return gdextension.TypeRect2i, reflect.TypeOf(Rect2i{})
	case "Vector3":
		return gdextension.TypeVector3, reflect.TypeOf(Vector3{})
	case "Vector3i":
		return gdextension.TypeVector3i, reflect.TypeOf(Vector3i{})
	case "Transform2D":
		return gdextension.TypeTransform2D, reflect.TypeOf(Transform2D{})
	case "Vector4":
		return gdextension.TypeVector4, reflect.TypeOf(Vector4{})
	case "Vector4i":
		return gdextension.TypeVector4i, reflect.TypeOf(Vector4i{})
	case "Plane":
		return gdextension.TypePlane, reflect.TypeOf(Plane{})
	case "Quaternion":
		return gdextension.TypeQuaternion, reflect.TypeOf(Quaternion{})
	case "AABB":
		return gdextension.TypeAABB, reflect.TypeOf(AABB{})
	case "Basis":
		return gdextension.TypeBasis, reflect.TypeOf(Basis{})
	case "Transform3D":
		return gdextension.TypeTransform3D, reflect.TypeOf(Transform3D{})
	case "Projection":
		return gdextension.TypeProjection, reflect.TypeOf(Projection{})
	case "Color":
		return gdextension.TypeColor, reflect.TypeOf(Color{})
	case "StringName":
		return gdextension.TypeStringName, reflect.TypeOf(StringName{})
	case "NodePath":
		return gdextension.TypeNodePath, reflect.TypeOf(NodePath{})
	case "RID":
		return gdextension.TypeRID, reflect.TypeOf(RID(0))
	case "Object":
		return gdextension.TypeObject, reflect.TypeOf(uintptr(0))
	case "Callable":
		return gdextension.TypeCallable, reflect.TypeOf(Callable{})
	case "Signal":
		return gdextension.TypeSignal, reflect.TypeOf(Signal{})
	case "Dictionary":
		return gdextension.TypeDictionary, reflect.TypeOf(Dictionary{})
	case "Array":
		return gdextension.TypeArray, reflect.TypeOf(Array{})
	case "PackedByteArray":
		return gdextension.TypePackedByteArray, reflect.TypeOf(PackedByteArray{})
	case "PackedInt32Array":
		return gdextension.TypePackedInt32Array, reflect.TypeOf(PackedInt32Array{})
	case "PackedInt64Array":
		return gdextension.TypePackedInt64Array, reflect.TypeOf(PackedInt64Array{})
	case "PackedFloat32Array":
		return gdextension.TypePackedFloat32Array, reflect.TypeOf(PackedFloat32Array{})
	case "PackedFloat64Array":
		return gdextension.TypePackedFloat64Array, reflect.TypeOf(PackedFloat64Array{})
	case "PackedStringArray":
		return gdextension.TypePackedStringArray, reflect.TypeOf(PackedStringArray{})
	case "PackedVector2Array":
		return gdextension.TypePackedVector2Array, reflect.TypeOf(PackedVector2Array{})
	case "PackedVector3Array":
		return gdextension.TypePackedVector3Array, reflect.TypeOf(PackedVector3Array{})
	case "PackedVector4Array":
		return gdextension.TypePackedVector4Array, reflect.TypeOf(PackedVector4Array{})
	case "PackedColorArray":
		return gdextension.TypePackedColorArray, reflect.TypeOf(PackedColorArray{})
	default:
		panic("gdextension.variantTypeFromName: unknown type " + s)
	}
}
