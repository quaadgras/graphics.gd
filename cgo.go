//go:build !generate

package gd

import "unsafe"

// class defined here to avoid go.dev/issue/50729
type class struct {
	Pointer unsafe.Pointer
	Runtime *API
}

type Float = float32

type String struct {
	ptr uintptr
}

type Vector2 struct {
	X, Y Float
}

type Vector2i struct {
	X, Y int32
}

type Rect2 struct {
	Position Vector2
	Size     Vector2
}

type Rect2i struct {
	Position Vector2i
	Size     Vector2i
}

type Vector3 struct {
	X, Y, Z Float
}

type Vector3i struct {
	X, Y, Z int32
}

type Transform2D struct {
	X Vector2
	Y Vector2
	O Vector2
}

type Vector4 struct {
	X, Y, Z, W Float
}

type Vector4i struct {
	X, Y, Z, W int32
}

type Plane struct {
	Matrix Vector4
}

type Quaternion struct {
	X, Y, Z, W Float
}

type AABB struct {
	Position Vector3
	Size     Vector3
}

type Basis struct {
	Rows [3]Vector3
}

type Transform3D struct {
	Basis  Basis
	Origin Vector3
}

type Projection struct {
	X, Y, Z, W Vector4
}

type Color struct {
	R, G, B, A Float
}

type StringName struct {
	ptr uintptr
}

type NodePath struct {
	ptr uintptr
}

type RID int64

type Callable struct {
	obj uintptr
	ptr uintptr
}

type Signal struct {
	obj uintptr
	ptr uintptr
}

type Dictionary struct {
	ptr uintptr
}

type Array struct {
	ptr uintptr
}

type ArrayOf[T any] struct {
	ptr uintptr
}

type PackedByteArray struct {
	ptr uintptr
}

type PackedInt32Array struct {
	ptr uintptr
}

type PackedInt64Array struct {
	ptr uintptr
}

type PackedFloat32Array struct {
	ptr uintptr
}

type PackedFloat64Array struct {
	ptr uintptr
}

type PackedStringArray struct {
	ptr uintptr
}

type PackedVector2Array struct {
	ptr uintptr
}

type PackedVector3Array struct {
	ptr uintptr
}

type PackedColorArray struct {
	ptr uintptr
}

type Variant struct {
	buf [24]byte
}
