package gd

import (
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

// All pointer types from the engine need to be defined here.
// Each pointer of a given shape (number of uintptrs) must implement
// a Free method and have a unique [pointers.Type] array size, this
// is required for the pointers garbage collector to call the correct
// Free method.

type Object pointers.Trio[Object]     // pointer, objectID, reserved
type RefCounted pointers.Trio[Object] // pointer, objectID, reserved

type enginePointer = EnginePointer
type packedPointers = PackedPointers
type VariantPointers = [3]uint64

type Variant pointers.Type[Variant, gdextension.Variant]
type Signal pointers.Type[Signal, gdextension.Signal]
type Callable pointers.Type[Callable, gdextension.Callable]

type Dictionary pointers.Type[Dictionary, gdextension.Dictionary]
type Array pointers.Type[Array, gdextension.Array]
type String pointers.Type[String, gdextension.String]
type StringName pointers.Type[StringName, gdextension.StringName]
type NodePath pointers.Type[NodePath, gdextension.NodePath]

type PackedByteArray pointers.Type[PackedByteArray, gdextension.PackedArray[byte]]
type PackedInt32Array pointers.Type[PackedInt32Array, gdextension.PackedArray[int32]]
type PackedInt64Array pointers.Type[PackedInt64Array, gdextension.PackedArray[int64]]
type PackedFloat32Array pointers.Type[PackedFloat32Array, gdextension.PackedArray[float32]]
type PackedFloat64Array pointers.Type[PackedFloat64Array, gdextension.PackedArray[float64]]
type PackedStringArray pointers.Type[PackedStringArray, gdextension.PackedArray[gdextension.String]]
type PackedVector2Array pointers.Type[PackedVector2Array, gdextension.PackedArray[Vector2]]
type PackedVector3Array pointers.Type[PackedVector3Array, gdextension.PackedArray[Vector3]]
type PackedVector4Array pointers.Type[PackedVector4Array, gdextension.PackedArray[Vector4]]
type PackedColorArray pointers.Type[PackedColorArray, gdextension.PackedArray[Color]]
