//go:generate go run ./internal/tool/generate
//go:generate go fmt ./...
package gd

import (
	gd "grow.graphics/gd/internal"
	"grow.graphics/xy"
)

type Axis int

const (
	X Axis = iota
	Y
	Z
	W
)

type (
	EulerAngles = xy.EulerAngles
	Radians     = xy.Radians
	Degrees     = xy.Degrees
)

type (
	Side       = xy.Side
	EulerOrder = xy.EulerOrder
)

type privateSignal = gd.Signal

type (
	Bool   = gd.Bool
	Int    = gd.Int
	Float  = gd.Float
	String = gd.String

	Vector2     = gd.Vector2
	Vector2i    = gd.Vector2i
	Rect2       = gd.Rect2
	Rect2i      = gd.Rect2i
	Vector3     = gd.Vector3
	Vector3i    = gd.Vector3i
	Transform2D = gd.Transform2D
	Vector4     = gd.Vector4
	Vector4i    = gd.Vector4i
	Plane       = gd.Plane
	Quaternion  = gd.Quaternion
	AABB        = gd.AABB
	Basis       = gd.Basis
	Transform3D = gd.Transform3D
	Projection  = gd.Projection

	Color      = gd.Color
	StringName = gd.StringName
	NodePath   = gd.NodePath
	RID        = gd.RID
	Callable   = gd.Callable

	Signal     = gd.Signal
	RefCounted = gd.RefCounted

	Dictionary = gd.Dictionary
	Array      = gd.Array
	Variant    = gd.Variant

	PackedByteArray    = gd.PackedByteArray
	PackedInt32Array   = gd.PackedInt32Array
	PackedInt64Array   = gd.PackedInt64Array
	PackedFloat32Array = gd.PackedFloat32Array
	PackedFloat64Array = gd.PackedFloat64Array
	PackedStringArray  = gd.PackedStringArray
	PackedVector2Array = gd.PackedVector2Array
	PackedVector3Array = gd.PackedVector3Array
	PackedColorArray   = gd.PackedColorArray
)

// SignalAs's T must be a function type that represents the arguments that are required
// to be passed to the signal.
type SignalAs[T any] struct {
	Signal

	Emit T
}

type isSignal interface {
	gd.IsSignal
	setSignal(gd.Signal)
}

func (s *SignalAs[T]) setSignal(signal gd.Signal) { s.Signal = signal }
