//go:build !generate

package gd

import (
	"structs"

	"graphics.gd/internal/callframe"
)

// cache is responsible for keeping a local copy for the various
// function pointers that are looked up at runtime and used for
// calling methods on classes.
type cache struct {
	utility utility
	builtin builtin
	typeset typeset
	variant variant
	Methods methods
}

type variant struct {
	FromType [TypeMax]func(ret callframe.Ptr[[3]uint64], arg callframe.Addr)
	IntoType [TypeMax]func(ret callframe.Addr, arg callframe.Ptr[[3]uint64])
}

type ObjectID uint64

type AudioFrame struct {
	_ structs.HostLayout

	Left, Right float32
}

type PhysicsServer2DExtensionMotionResult struct {
	_ structs.HostLayout

	Travel, Remainder, CollisionPoint, CollisionNormal, ColliderVelocity Vector2
	CollisionDepth, CollisionSafeFraction, CollisionUnsafeFraction       float32
	CollisionLocalShape                                                  int32
	ColliderID                                                           uint64
	ColliderRID                                                          RID
	ColliderShape                                                        int32
}

type PhysicsServer2DExtensionRayResult struct {
	_ structs.HostLayout

	Position, Normal Vector2
	RID              RID
	ColliderID       uint64
	Collider         gdptr
	Shape            int32
}

type PhysicsServer2DExtensionShapeRestInfo struct {
	_ structs.HostLayout

	Point, Normal, LinearVelocity Vector2
	RID                           RID
	ColliderID                    uint64
	Shape                         int32
}

type PhysicsServer2DExtensionShapeResult struct {
	_ structs.HostLayout

	RID        RID
	ColliderID uint64
	Collider   gdptr
	Shape      int32
}

type PhysicsServer3DExtensionRayResult struct {
	_ structs.HostLayout

	Position, Normal Vector3
	RID              RID
	ColliderID       uint64
	Collider         gdptr
	Shape            int32
	FaceIndex        int32
}
type PhysicsServer3DExtensionShapeResult struct {
	_ structs.HostLayout

	RID        RID
	ColliderID uint64
	Collider   gdptr
	Shape      int32
}
type PhysicsServer3DExtensionShapeRestInfo struct {
	_ structs.HostLayout

	Point, Normal  Vector3
	RID            RID
	ColliderID     uint64
	Shape          int32
	LinearVelocity Vector3
}
type PhysicsServer3DExtensionMotionCollision struct {
	_ structs.HostLayout

	Position, Normal        Vector3
	ColliderVelocity        Vector3
	ColliderAngularVelocity Vector3
	Depth                   float32
	LocalShape              int32
	ColliderID              ObjectID
	Collider                RID
	ColliderShape           int32
}
type PhysicsServer3DExtensionMotionResult struct {
	_ structs.HostLayout

	Travel                  Vector3
	Remainder               Vector3
	CollisionDepth          float32
	CollisionSafeFraction   float32
	CollisionUnsafeFraction float32
	Collisions              [32]PhysicsServer3DExtensionMotionCollision
	CollisionCount          int32
}
type ScriptLanguageExtensionProfilingInfo struct {
	_ structs.HostLayout

	Signature    gdptr // StringName.Pointer()
	CallCount    uint64
	TotalTime    uint64
	SelfTime     uint64
	InternalTime uint64
}
type Glyph struct {
	_ structs.HostLayout

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
	_ structs.HostLayout

	LeadingCaret, TrailingCaret Rect2

	LeadingDirection, TrailingDirection uint32
}
