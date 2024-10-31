//go:generate go run ./internal/tool/generate
//go:generate go run ./internal/tool/generate/v2
//go:generate go fmt ./...
package gd

import (
	gd "grow.graphics/gd/internal"
)

type (
	Variant = gd.Variant

	Bool  = gd.Bool
	Int   = gd.Int
	Float = gd.Float

	Color = gd.Color

	Array      = gd.Array
	Dictionary = gd.Dictionary

	Object = gd.Object

	PackedByteArray    = gd.PackedByteArray
	PackedColorArray   = gd.PackedColorArray
	PackedFloat32Array = gd.PackedFloat32Array
	PackedFloat64Array = gd.PackedFloat64Array
	PackedInt32Array   = gd.PackedInt32Array
	PackedInt64Array   = gd.PackedInt64Array
	PackedStringArray  = gd.PackedStringArray
	PackedVector2Array = gd.PackedVector2Array
	PackedVector3Array = gd.PackedVector3Array
	PackedVector4Array = gd.PackedVector4Array

	Plane = gd.Plane

	Projection = gd.Projection
	Quaternion = gd.Quaternion

	AABB   = gd.AABB
	Rect2  = gd.Rect2
	Rect2i = gd.Rect2i

	RID = gd.RID

	Signal   = gd.Signal
	Callable = gd.Callable

	String     = gd.String
	StringName = gd.StringName
	NodePath   = gd.NodePath

	Basis       = gd.Basis
	Transform2D = gd.Transform2D
	Transform3D = gd.Transform3D

	Vector2  = gd.Vector2
	Vector2i = gd.Vector2i
	Vector3  = gd.Vector3
	Vector3i = gd.Vector3i
	Vector4  = gd.Vector4
	Vector4i = gd.Vector4i
)
