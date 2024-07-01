//go:generate go run ./internal/tool/generate
//go:generate go fmt ./...
package gd

import (
	"reflect"

	gd "grow.graphics/gd/internal"
	"grow.graphics/xy"
	"runtime.link/mmm"
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

	Signal = gd.Signal

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

func NewArrayOf[T any](godot Context) ArrayOf[T] {
	array := gd.TypedArray[T](godot.Array())
	rtype := reflect.TypeOf([0]T{}).Elem()
	tmp := gd.NewContext(godot.API)
	defer tmp.End()
	godot.API.Array.SetTyped(Array(array), variantTypeOf(rtype), tmp.StringName(classNameOf(rtype)), Object{})
	return array
}

type ArrayOf[T any] interface {
	gd.IsArray
	gd.IsArrayType[T]
	mmm.ManagedPointer[uintptr]

	All(method Callable) bool
	Any(method Callable) bool
	Append(value T)
	AppendArray(array gd.ArrayOf[T])
	Assign(array gd.ArrayOf[T])
	Back(ctx Context) T
	Bsearch(value T, before bool) int64
	BsearchCustom(value T, fn Callable, before bool) int64
	Clear()
	Count(value T) int64
	Duplicate(ctx Context, deep bool) gd.ArrayOf[T]
	Erase(value T)
	Fill(value T)
	Filter(ctx Context, method Callable) gd.ArrayOf[T]
	Find(what T, from int64) int64
	Free()
	Front(ctx Context) T
	GetTypedBuiltin() int64
	GetTypedClassName(ctx Context) StringName
	GetTypedScript(ctx Context) Variant
	Has(value T) bool
	Hash() int64
	Index(ctx Context, index int64) T
	Insert(position int64, value T) int64
	IsEmpty() bool
	IsReadOnly() bool
	IsSameTyped(array Array) bool
	IsTyped() bool
	MakeReadOnly()
	Map(ctx Context, method Callable) gd.ArrayOf[T]
	Max(ctx Context) T
	Min(ctx Context) T
	PickRandom(ctx Context) T
	PopAt(ctx Context, position int64) T
	PopBack(ctx Context) T
	PopFront(ctx Context) T
	PushBack(value T)
	PushFront(value T)
	Reduce(ctx Context, method Callable, accum T) T
	RemoveAt(position int64)
	Resize(size int64) int64
	Reverse()
	Rfind(what T, from int64) int64
	SetIndex(index int64, value T)
	Shuffle()
	Size() int64
	Slice(ctx Context, begin int64, end int64, step int64, deep bool) gd.ArrayOf[T]
	Sort()
	SortCustom(fn Callable)

	UnmarshalInto(any) error
}
