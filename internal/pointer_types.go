package gd

import (
	"grow.graphics/gd/internal/discreet"
)

// All pointer types from the engine need to be defined here.
// Each pointer of a given shape (number of uintptrs) must implement
// a Free method and have a unique [discreet.Type] array size, this
// is required for the discreet garbage collector to call the correct
// Free method.

type Variant discreet.PointerNamed[Variant, [3]uintptr, [0]discreet.Type]
type Object discreet.PointerNamed[Object, [3]uintptr, [1]discreet.Type]
type RefCounted discreet.PointerNamed[RefCounted, [3]uintptr, [1]discreet.Type]

type Dictionary discreet.PointerNamed[Dictionary, [1]uintptr, [0]discreet.Type]
type Array discreet.PointerNamed[Array, [1]uintptr, [1]discreet.Type]
type String discreet.PointerNamed[String, [1]uintptr, [2]discreet.Type]
type StringName discreet.PointerNamed[StringName, [1]uintptr, [3]discreet.Type]
type NodePath discreet.PointerNamed[NodePath, [1]uintptr, [4]discreet.Type]

type Signal discreet.PointerNamed[Signal, [2]uintptr, [0]discreet.Type]
type Callable discreet.PointerNamed[Callable, [2]uintptr, [1]discreet.Type]
type PackedByteArray discreet.PointerNamed[PackedByteArray, [2]uintptr, [2]discreet.Type]
type PackedInt32Array discreet.PointerNamed[PackedInt32Array, [2]uintptr, [3]discreet.Type]
type PackedInt64Array discreet.PointerNamed[PackedInt64Array, [2]uintptr, [4]discreet.Type]
type PackedFloat32Array discreet.PointerNamed[PackedFloat32Array, [2]uintptr, [5]discreet.Type]
type PackedFloat64Array discreet.PointerNamed[PackedFloat64Array, [2]uintptr, [6]discreet.Type]
type PackedStringArray discreet.PointerNamed[PackedStringArray, [2]uintptr, [7]discreet.Type]
type PackedVector2Array discreet.PointerNamed[PackedVector2Array, [2]uintptr, [8]discreet.Type]
type PackedVector3Array discreet.PointerNamed[PackedVector3Array, [2]uintptr, [9]discreet.Type]
type PackedVector4Array discreet.PointerNamed[PackedVector4Array, [2]uintptr, [10]discreet.Type]
type PackedColorArray discreet.PointerNamed[PackedColorArray, [2]uintptr, [11]discreet.Type]
