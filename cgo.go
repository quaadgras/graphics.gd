//go:build !generate

package gd

import "unsafe"

type Float = float32
type Int = int64

type String struct {
	Class[String, Pointer]
}

func (s String) String() string {
	if s.super.Val == [2]uintptr{} {
		return ""
	}
	var buf = make([]byte, s.super.API.String_length(&s))
	s.super.API.Extension.Strings.Get(&s, buf)
	return string(buf)
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
	Class[String, Pointer]
}

func (s StringName) String() string {
	if s.super.Val == [2]uintptr{} {
		return ""
	}
	var tmp String
	s.super.API.String_NewFromStringName(s)
	defer s.super.API.Extension.Variants.Destructor(TypeString)(unsafe.Pointer(&tmp))
	return tmp.String()
}

type NodePath struct {
	Class[String, Pointer]
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
	Class[Dictionary, Pointer]
}

type Array struct {
	Class[Array, Pointer]
}

type ArrayOf[T any] struct {
	Class[ArrayOf[T], Pointer]
}

type PackedByteArray struct {
	Class[PackedByteArray, Pointer]
}

type PackedInt32Array struct {
	Class[PackedInt32Array, Pointer]
}

type PackedInt64Array struct {
	Class[PackedInt64Array, Pointer]
}

type PackedFloat32Array struct {
	Class[PackedFloat32Array, Pointer]
}

type PackedFloat64Array struct {
	Class[PackedFloat64Array, Pointer]
}

type PackedStringArray struct {
	Class[PackedStringArray, Pointer]
}

func (array PackedStringArray) Slice() []string {
	var slice = make([]string, array.super.API.PackedStringArray_size(&array))
	for i := 0; i < len(slice); i++ {
		elem := array.super.API.Extension.PackedStringArray.Index(&array, int64(i))
		var tmp String
		tmp.super.Val[0] = *elem
		tmp.super.API = array.super.API
		slice[i] = tmp.String()
	}
	return slice
}

type PackedVector2Array struct {
	Class[PackedVector2Array, Pointer]
}

type PackedVector3Array struct {
	Class[PackedVector3Array, Pointer]
}

type PackedColorArray struct {
	Class[PackedColorArray, Pointer]
}

type Variant struct {
	buf [24]byte
}

type AudioFrame struct {
	Left, Right float32
}

type PhysicsServer2DExtensionMotionResult struct {
	Travel, Remainder, CollisionPoint, CollisionNormal, ColliderVelocity Vector2
	CollisionDepth, CollisionSafeFraction, CollisionUnsafeFraction       float32
	CollisionLocalShape                                                  int32
	ColliderID                                                           uint64
	ColliderRID                                                          RID
	ColliderShape                                                        int32
}

type PhysicsServer2DExtensionRayResult struct {
	Position, Normal Vector2
	RID              RID
	ColliderID       uint64
	Collider         *Object
	Shape            int32
}

type PhysicsServer2DExtensionShapeRestInfo struct {
	Point, Normal, LinearVelocity Vector2
	RID                           RID
	ColliderID                    uint64
	Shape                         int32
}

type PhysicsServer2DExtensionShapeResult struct {
	RID        RID
	ColliderID uint64
	Collider   *Object
	Shape      int32
}

type PhysicsServer3DExtensionRayResult struct {
	Position, Normal Vector3
	RID              RID
	ColliderID       uint64
	Collider         uintptr
	Shape            int32
}
type PhysicsServer3DExtensionShapeResult struct {
	RID        RID
	ColliderID uint64
	Collider   uintptr
	Shape      int32
}
type PhysicsServer3DExtensionShapeRestInfo struct {
	Point, Normal  Vector3
	RID            RID
	ColliderID     uint64
	Shape          int32
	LinearVelocity Vector3
}
type PhysicsServer3DExtensionMotionCollision struct {
	Position, Normal Vector3
	ColliderVelocity Vector3
	Depth            float32
	LocalShape       int32
	ColliderID       uint64
	Collider         RID
	ColliderShape    int32
}
type PhysicsServer3DExtensionMotionResult struct {
	Travel                  Vector3
	Remainder               Vector3
	CollisionSafeFraction   float32
	CollisionUnsafeFraction float32
	Collisions              [32]PhysicsServer3DExtensionMotionCollision
	CollisionCount          int32
}
type ScriptLanguageExtensionProfilingInfo struct {
	Signature  StringName
	Call_count uint64
	TotalTime  uint64
	SelfTime   uint64
}
type Glyph struct {
	Start    int32
	End      int32
	Count    uint8
	Repeat   uint8
	Flags    uint16
	Xoffset  float32
	Yoffset  float32
	Advance  float32
	FontRID  RID
	FontSize int32
	Index    int32
}
type CaretInfo struct {
	LeadingCaret, TrailingCaret Rect2

	LeadingDirection, TrailingDirection int64
}
