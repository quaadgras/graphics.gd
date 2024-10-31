package gd

import (
	"grow.graphics/gd/internal/pointers"
)

// All pointer types from the engine need to be defined here.
// Each pointer of a given shape (number of uintptrs) must implement
// a Free method and have a unique [pointers.Type] array size, this
// is required for the pointers garbage collector to call the correct
// Free method.

type Variant pointers.PointerNamed[Variant, [3]uintptr, VariantPointerType]
type Object pointers.PointerNamed[Object, [3]uintptr, ObjectPointerType]
type RefCounted pointers.PointerNamed[RefCounted, [3]uintptr, RefCountedPointerType]

type VariantPointerType = [0]pointers.Type
type ObjectPointerType = [1]pointers.Type
type RefCountedPointerType = [1]pointers.Type

type Dictionary pointers.PointerNamed[Dictionary, [1]uintptr, DictionaryPointerType]
type Array pointers.PointerNamed[Array, [1]uintptr, ArrayPointerType]
type String pointers.PointerNamed[String, [1]uintptr, StringPointerType]
type StringName pointers.PointerNamed[StringName, [1]uintptr, StringNamePointerType]
type NodePath pointers.PointerNamed[NodePath, [1]uintptr, NodePathPointerType]

type DictionaryPointerType = [0]pointers.Type
type ArrayPointerType = [1]pointers.Type
type StringPointerType = [2]pointers.Type
type StringNamePointerType = [3]pointers.Type
type NodePathPointerType = [4]pointers.Type

type Signal pointers.PointerNamed[Signal, [2]uintptr, SignalPointerType]
type Callable pointers.PointerNamed[Callable, [2]uintptr, CallablePointerType]
type PackedByteArray pointers.PointerNamed[PackedByteArray, [2]uintptr, PackedByteArrayPointerType]
type PackedInt32Array pointers.PointerNamed[PackedInt32Array, [2]uintptr, PackedInt32ArrayPointerType]
type PackedInt64Array pointers.PointerNamed[PackedInt64Array, [2]uintptr, PackedInt64ArrayPointerType]
type PackedFloat32Array pointers.PointerNamed[PackedFloat32Array, [2]uintptr, PackedFloat32ArrayPointerType]
type PackedFloat64Array pointers.PointerNamed[PackedFloat64Array, [2]uintptr, PackedFloat64ArrayPointerType]
type PackedStringArray pointers.PointerNamed[PackedStringArray, [2]uintptr, PackedStringArrayPointerType]
type PackedVector2Array pointers.PointerNamed[PackedVector2Array, [2]uintptr, PackedVector2ArrayPointerType]
type PackedVector3Array pointers.PointerNamed[PackedVector3Array, [2]uintptr, PackedVector3ArrayPointerType]
type PackedVector4Array pointers.PointerNamed[PackedVector4Array, [2]uintptr, PackedVector4ArrayPointerType]
type PackedColorArray pointers.PointerNamed[PackedColorArray, [2]uintptr, PackedColorArrayPointerType]

type SignalPointerType = [0]pointers.Type
type CallablePointerType = [1]pointers.Type
type PackedByteArrayPointerType = [2]pointers.Type
type PackedInt32ArrayPointerType = [3]pointers.Type
type PackedInt64ArrayPointerType = [4]pointers.Type
type PackedFloat32ArrayPointerType = [5]pointers.Type
type PackedFloat64ArrayPointerType = [6]pointers.Type
type PackedStringArrayPointerType = [7]pointers.Type
type PackedVector2ArrayPointerType = [8]pointers.Type
type PackedVector3ArrayPointerType = [9]pointers.Type
type PackedVector4ArrayPointerType = [10]pointers.Type
type PackedColorArrayPointerType = [11]pointers.Type
