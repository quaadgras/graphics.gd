package classdb

import "grow.graphics/gd/internal"

type AudioFrame struct {
	Left, Right float32
}

type PhysicsServer2DExtensionMotionResult struct {
	Travel, Remainder, CollisionPoint, CollisionNormal, ColliderVelocity internal.Vector2
	CollisionDepth, CollisionSafeFraction, CollisionUnsafeFraction       float32
	CollisionLocalShape                                                  int32
	ColliderID                                                           uint64
	ColliderRID                                                          internal.RID
	ColliderShape                                                        int32
}

type PhysicsServer2DExtensionRayResult struct {
	Position, Normal internal.Vector2
	RID              internal.RID
	ColliderID       uint64
	Collider         *internal.Object
	Shape            int32
}

type PhysicsServer2DExtensionShapeRestInfo struct {
	Point, Normal, LinearVelocity internal.Vector2
	RID                           internal.RID
	ColliderID                    uint64
	Shape                         int32
}

type PhysicsServer2DExtensionShapeResult struct {
	RID        internal.RID
	ColliderID uint64
	Collider   *internal.Object
	Shape      int32
}

type PhysicsServer3DExtensionRayResult struct {
	Position, Normal internal.Vector3
	RID              internal.RID
	ColliderID       uint64
	Collider         uintptr
	Shape            int32
}
type PhysicsServer3DExtensionShapeResult struct {
	RID        internal.RID
	ColliderID uint64
	Collider   uintptr
	Shape      int32
}
type PhysicsServer3DExtensionShapeRestInfo struct {
	Point, Normal  internal.Vector3
	RID            internal.RID
	ColliderID     uint64
	Shape          int32
	LinearVelocity internal.Vector3
}
type PhysicsServer3DExtensionMotionCollision struct {
	Position, Normal internal.Vector3
	ColliderVelocity internal.Vector3
	Depth            float32
	LocalShape       int32
	ColliderID       uint64
	Collider         internal.RID
	ColliderShape    int32
}
type PhysicsServer3DExtensionMotionResult struct {
	Travel                  internal.Vector3
	Remainder               internal.Vector3
	CollisionSafeFraction   float32
	CollisionUnsafeFraction float32
	Collisions              [32]PhysicsServer3DExtensionMotionCollision
	CollisionCount          int32
}
type ScriptLanguageExtensionProfilingInfo struct {
	Signature  internal.StringName
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
	FontRID  internal.RID
	FontSize int32
	Index    int32
}
type CaretInfo struct {
	LeadingCaret, TrailingCaret internal.Rect2

	LeadingDirection, TrailingDirection int64
}
