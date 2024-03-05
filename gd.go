package gd

import (
	internal "grow.graphics/gd/internal"
	"grow.graphics/xy"
)

type Axis int

const (
	X Axis = iota
	Y
	Z
	W
)

// NewVector2 constructs a new Vector2 from the given x and y.
func NewVector2(x, y Float) Vector2 { return xy.NewVector2(x, y) }

// NewVector2i constructs a new Vector2i from the given x and y.
func NewVector2i(x, y Int) Vector2i { return Vector2i{int32(x), int32(y)} }

// NewRect2 constructs a Rect2 by setting its position to (x, y), and its size to (width, height).
func NewRect2(x, y, width, height Float) Rect2 {
	return xy.NewRect2(x, y, width, height)
}

// NewRect2i constructs a Rect2i by setting its position to (x, y), and its size to (width, height).
func NewRect2i(x, y, width, height Int) Rect2i {
	return Rect2i{Position: Vector2i{int32(x), int32(y)}, Size: Vector2i{int32(width), int32(height)}}
}

// NewVector3 constructs a new Vector3 from the given x, y, and z.
func NewVector3(x, y, z Float) Vector3 { return xy.NewVector3(x, y, z) }

// NewVector3i constructs a new Vector3i from the given x, y, and z.
func NewVector3i(x, y, z Int) Vector3i { return Vector3i{int32(x), int32(y), int32(z)} }

// NewTransform2D constructs a new Transform2D from the given rotation and position.
func NewTransform2D(rotation Radians, scale Vector2, skew Radians, position Vector2) Transform2D {
	return xy.NewTransform2D(rotation, scale, skew, position)
}

// NewVector4 constructs a new Vector4 from the given x, y, z, and w.
func NewVector4(x, y, z, w Float) Vector4 {
	return xy.NewVector4(x, y, z, w)
}

// NewVector4i constructs a new Vector4i from the given x, y, z, and w.
func NewVector4i(x, y, z, w Int) Vector4i {
	return Vector4i{int32(x), int32(y), int32(z), int32(w)}
}

// NewPlane creates a plane from the three points, given in clockwise order.
func NewPlane(a, b, c Vector3) Plane { return xy.NewPlane(a, b, c) }

// NewBasisScaledBy constructs a pure scale basis matrix with no rotation or shearing. The scale values are set as
// the diagonal of the matrix, and the other parts of the matrix are zero.
func NewBasisScaledBy(scale Vector3) Basis { return xy.NewBasisScaledBy(scale) }

// NewBasisRotatedAround constructs a pure rotation basis matrix, rotated around the given axis by angle (in radians).
// The axis must be a normalized vector.
func NewBasisRotatedAround(axis Vector3, angle Radians) Basis {
	return xy.NewBasisRotatedAround(axis, angle)
}

type (
	EulerAngles = xy.EulerAngles
	Radians     = xy.Radians
	Degrees     = xy.Degrees
)

type (
	Side       = xy.Side
	EulerOrder = xy.EulerOrder
)

type (
	Bool   = internal.Bool
	Int    = internal.Int
	Float  = internal.Float
	String = internal.String

	Vector2     = internal.Vector2
	Vector2i    = internal.Vector2i
	Rect2       = internal.Rect2
	Rect2i      = internal.Rect2i
	Vector3     = internal.Vector3
	Vector3i    = internal.Vector3i
	Transform2D = internal.Transform2D
	Vector4     = internal.Vector4
	Vector4i    = internal.Vector4i
	Plane       = internal.Plane
	Quaternion  = internal.Quaternion
	AABB        = internal.AABB
	Basis       = internal.Basis
	Transform3D = internal.Transform3D
	Projection  = internal.Projection

	Color      = internal.Color
	StringName = internal.StringName
	NodePath   = internal.NodePath
	RID        = internal.RID
	Object     = internal.Object
	Callable   = internal.Callable

	// Signal's T must be a function type.
	Signal[T any] struct {
		internal.Signal

		Emit T
	}

	Dictionary = internal.Dictionary
	Array      = internal.Array

	PackedByteArray    = internal.PackedByteArray
	PackedInt32Array   = internal.PackedInt32Array
	PackedInt64Array   = internal.PackedInt64Array
	PackedFloat32Array = internal.PackedFloat32Array
	PackedFloat64Array = internal.PackedFloat64Array
	PackedStringArray  = internal.PackedStringArray
	PackedVector2Array = internal.PackedVector2Array
	PackedVector3Array = internal.PackedVector3Array
	PackedColorArray   = internal.PackedColorArray
)
