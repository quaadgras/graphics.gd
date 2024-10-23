//go:generate go run ./internal/tool/generate
//go:generate go fmt ./...
package gd

import (
	"iter"
	"reflect"

	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/internal/mmm"
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

func NewArrayOf[T any](godot Lifetime) ArrayOf[T] {
	array := gd.TypedArray[T](godot.Array())
	rtype := reflect.TypeOf([0]T{}).Elem()
	tmp := NewLifetime(godot)
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
	Back(ctx Lifetime) T
	Bsearch(value T, before bool) int64
	BsearchCustom(value T, fn Callable, before bool) int64
	Clear()
	Count(value T) int64
	Duplicate(ctx Lifetime, deep bool) gd.ArrayOf[T]
	Erase(value T)
	Fill(value T)
	Filter(ctx Lifetime, method Callable) gd.ArrayOf[T]
	Find(what T, from int64) int64
	Free()
	Front(ctx Lifetime) T
	GetTypedBuiltin() int64
	GetTypedClassName(ctx Lifetime) StringName
	GetTypedScript(ctx Lifetime) Variant
	Has(value T) bool
	Hash() int64
	Index(ctx Lifetime, index int64) T
	Insert(position int64, value T) int64
	IsEmpty() bool
	IsReadOnly() bool
	IsSameTyped(array Array) bool
	IsTyped() bool
	MakeReadOnly()
	Map(ctx Lifetime, method Callable) gd.ArrayOf[T]
	Max(ctx Lifetime) T
	Min(ctx Lifetime) T
	PickRandom(ctx Lifetime) T
	PopAt(ctx Lifetime, position int64) T
	PopBack(ctx Lifetime) T
	PopFront(ctx Lifetime) T
	PushBack(value T)
	PushFront(value T)
	Reduce(ctx Lifetime, method Callable, accum T) T
	RemoveAt(position int64)
	Resize(size int64) int64
	Reverse()
	Rfind(what T, from int64) int64
	SetIndex(index int64, value T)
	Shuffle()
	Size() int64
	Slice(ctx Lifetime, begin int64, end int64, step int64, deep bool) gd.ArrayOf[T]
	Sort()
	SortCustom(fn Callable)

	UnmarshalInto(any) error
	Iter() iter.Seq2[Int, T]
}
