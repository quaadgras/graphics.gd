package gd

import (
	"grow.graphics/gd/internal/discreet"
)

// All pointer types from the engine need to be defined here.
// Each pointer of a given shape (number of uintptrs) must implement
// a Free method and have a unique [discreet.Type] array size, this
// is required for the discreet garbage collector to call the correct
// Free method.

type Variant discreet.PointerNamed[Variant, [3]uintptr, VariantPointerType]
type Object discreet.PointerNamed[Object, [3]uintptr, ObjectPointerType]
type RefCounted discreet.PointerNamed[RefCounted, [3]uintptr, RefCountedPointerType]

type VariantPointerType = [0]discreet.Type
type ObjectPointerType = [1]discreet.Type
type RefCountedPointerType = [1]discreet.Type

type Dictionary discreet.PointerNamed[Dictionary, [1]uintptr, DictionaryPointerType]
type Array discreet.PointerNamed[Array, [1]uintptr, ArrayPointerType]
type String discreet.PointerNamed[String, [1]uintptr, StringPointerType]
type StringName discreet.PointerNamed[StringName, [1]uintptr, StringNamePointerType]
type NodePath discreet.PointerNamed[NodePath, [1]uintptr, NodePathPointerType]

type DictionaryPointerType = [0]discreet.Type
type ArrayPointerType = [1]discreet.Type
type StringPointerType = [2]discreet.Type
type StringNamePointerType = [3]discreet.Type
type NodePathPointerType = [4]discreet.Type

type Signal discreet.PointerNamed[Signal, [2]uintptr, SignalPointerType]
type Callable discreet.PointerNamed[Callable, [2]uintptr, CallablePointerType]
type PackedByteArray discreet.PointerNamed[PackedByteArray, [2]uintptr, PackedByteArrayPointerType]
type PackedInt32Array discreet.PointerNamed[PackedInt32Array, [2]uintptr, PackedInt32ArrayPointerType]
type PackedInt64Array discreet.PointerNamed[PackedInt64Array, [2]uintptr, PackedInt64ArrayPointerType]
type PackedFloat32Array discreet.PointerNamed[PackedFloat32Array, [2]uintptr, PackedFloat32ArrayPointerType]
type PackedFloat64Array discreet.PointerNamed[PackedFloat64Array, [2]uintptr, PackedFloat64ArrayPointerType]
type PackedStringArray discreet.PointerNamed[PackedStringArray, [2]uintptr, PackedStringArrayPointerType]
type PackedVector2Array discreet.PointerNamed[PackedVector2Array, [2]uintptr, PackedVector2ArrayPointerType]
type PackedVector3Array discreet.PointerNamed[PackedVector3Array, [2]uintptr, PackedVector3ArrayPointerType]
type PackedVector4Array discreet.PointerNamed[PackedVector4Array, [2]uintptr, PackedVector4ArrayPointerType]
type PackedColorArray discreet.PointerNamed[PackedColorArray, [2]uintptr, PackedColorArrayPointerType]

type SignalPointerType = [0]discreet.Type
type CallablePointerType = [1]discreet.Type
type PackedByteArrayPointerType = [2]discreet.Type
type PackedInt32ArrayPointerType = [3]discreet.Type
type PackedInt64ArrayPointerType = [4]discreet.Type
type PackedFloat32ArrayPointerType = [5]discreet.Type
type PackedFloat64ArrayPointerType = [6]discreet.Type
type PackedStringArrayPointerType = [7]discreet.Type
type PackedVector2ArrayPointerType = [8]discreet.Type
type PackedVector3ArrayPointerType = [9]discreet.Type
type PackedVector4ArrayPointerType = [10]discreet.Type
type PackedColorArrayPointerType = [11]discreet.Type
