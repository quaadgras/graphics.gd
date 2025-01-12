//go:build amd64 || arm64

package gd

import (
	"graphics.gd/internal/pointers"
)

type gdptr uint64

type Variant pointers.Trio[Variant]
type Signal pointers.Pair[Signal]
type Callable pointers.Pair[Callable]

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

type EnginePointer = uint64
type PackedPointers = [2]uint64
