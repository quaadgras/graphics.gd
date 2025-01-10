package gd

import (
	"graphics.gd/internal/pointers"
)

// All pointer types from the engine need to be defined here.
// Each pointer of a given shape (number of uintptrs) must implement
// a Free method and have a unique [pointers.Type] array size, this
// is required for the pointers garbage collector to call the correct
// Free method.

type Object pointers.Trio[Object]
type RefCounted pointers.Trio[Object]

type Dictionary pointers.Solo[Dictionary]
type Array pointers.Solo[Array]
type String pointers.Solo[String]
type StringName pointers.Solo[StringName]
type NodePath pointers.Solo[NodePath]

type PackedByteArray pointers.Pair[PackedByteArray]
type PackedInt32Array pointers.Pair[PackedInt32Array]
type PackedInt64Array pointers.Pair[PackedInt64Array]
type PackedFloat32Array pointers.Pair[PackedFloat32Array]
type PackedFloat64Array pointers.Pair[PackedFloat64Array]
type PackedStringArray pointers.Pair[PackedStringArray]
type PackedVector2Array pointers.Pair[PackedVector2Array]
type PackedVector3Array pointers.Pair[PackedVector3Array]
type PackedVector4Array pointers.Pair[PackedVector4Array]
type PackedColorArray pointers.Pair[PackedColorArray]

type variantPointers = VariantPointers
type signalPointers = SignalPointers
type callablePointers = CallablePointers
