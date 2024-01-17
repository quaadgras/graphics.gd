package gd

import gd "grow.graphics/gd/internal"

type AudioFrame struct {
	Left, Right float32
}

type PhysicsServer2DExtensionMotionResult struct {
	Travel, Remainder, CollisionPoint, CollisionNormal, ColliderVelocity gd.Vector2
	CollisionDepth, CollisionSafeFraction, CollisionUnsafeFraction       float32
	CollisionLocalShape                                                  int32
	ColliderID                                                           uint64
	ColliderRID                                                          gd.RID
	ColliderShape                                                        int32
}

type PhysicsServer2DExtensionRayResult struct {
	Position, Normal gd.Vector2
	RID              gd.RID
	ColliderID       uint64
	Collider         *gd.Object
	Shape            int32
}

type PhysicsServer2DExtensionShapeRestInfo struct {
	Point, Normal, LinearVelocity gd.Vector2
	RID                           gd.RID
	ColliderID                    uint64
	Shape                         int32
}

type PhysicsServer2DExtensionShapeResult struct {
	RID        gd.RID
	ColliderID uint64
	Collider   *gd.Object
	Shape      int32
}

type PhysicsServer3DExtensionRayResult struct {
	Position, Normal gd.Vector3
	RID              gd.RID
	ColliderID       uint64
	Collider         uintptr
	Shape            int32
}
type PhysicsServer3DExtensionShapeResult struct {
	RID        gd.RID
	ColliderID uint64
	Collider   uintptr
	Shape      int32
}
type PhysicsServer3DExtensionShapeRestInfo struct {
	Point, Normal  gd.Vector3
	RID            gd.RID
	ColliderID     uint64
	Shape          int32
	LinearVelocity gd.Vector3
}
type PhysicsServer3DExtensionMotionCollision struct {
	Position, Normal gd.Vector3
	ColliderVelocity gd.Vector3
	Depth            float32
	LocalShape       int32
	ColliderID       uint64
	Collider         gd.RID
	ColliderShape    int32
}
type PhysicsServer3DExtensionMotionResult struct {
	Travel                  gd.Vector3
	Remainder               gd.Vector3
	CollisionSafeFraction   float32
	CollisionUnsafeFraction float32
	Collisions              [32]PhysicsServer3DExtensionMotionCollision
	CollisionCount          int32
}
type ScriptLanguageExtensionProfilingInfo struct {
	Signature  gd.StringName
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
	FontRID  gd.RID
	FontSize int32
	Index    int32
}
type CaretInfo struct {
	LeadingCaret, TrailingCaret gd.Rect2

	LeadingDirection, TrailingDirection int64
}
