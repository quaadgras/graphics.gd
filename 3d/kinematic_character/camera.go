package main

import (
	"graphics.gd/classdb/Camera3D"
	"graphics.gd/classdb/CollisionObject3D"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Vector3"
)

// FollowCamera keeps a bounded distance from its parent node.
type FollowCamera struct {
	Camera3D.Extension[FollowCamera] `gd:"FollowCamera"`

	MinDistance  Float.X
	MaxDistance  Float.X
	AngleVAdjust Float.X

	collisionException []RID.Body3D
	maxHeight          Float.X
	minHeight          Float.X

	target Node3D.Instance
}

func NewFollowCamera() *FollowCamera {
	return &FollowCamera{
		MinDistance: 0.5,
		MaxDistance: 3.0,
		maxHeight:   2.0,
		minHeight:   0,
	}
}

func (f *FollowCamera) Ready() {
	f.target = Object.To[Node3D.Instance](f.AsNode().GetParent())
	if body, ok := Object.As[CollisionObject3D.Instance](f.target.AsNode().GetParent()); ok {
		f.collisionException = append(f.collisionException, body.GetRid())
	}
	// Detaches the camera transform from the parent spatial node.
	f.AsNode3D().SetTopLevel(true)
}

func (f *FollowCamera) PhysicsProcess(_ Float.X) {
	node := f.AsNode3D()
	targetPos := f.target.GlobalTransform().Origin
	cameraPos := node.GlobalTransform().Origin

	deltaPos := Vector3.Sub(cameraPos, targetPos)

	// Regular delta follow.

	// Check ranges.
	if Vector3.Length(deltaPos) < f.MinDistance {
		deltaPos = Vector3.MulX(Vector3.Normalized(deltaPos), f.MinDistance)
	} else if Vector3.Length(deltaPos) > f.MaxDistance {
		deltaPos = Vector3.MulX(Vector3.Normalized(deltaPos), f.MaxDistance)
	}

	// Check upper and lower height.
	if deltaPos.Y > f.maxHeight {
		deltaPos.Y = f.maxHeight
	}
	if deltaPos.Y < f.minHeight {
		deltaPos.Y = f.minHeight
	}

	cameraPos = Vector3.Add(targetPos, deltaPos)

	node.LookAtFromPosition(cameraPos, targetPos)

	// Turn a little up or down.
	t := node.Transform()
	t.Basis = Basis.Mul(Basis.RotatesAxisAngle(t.Basis.X, Angle.InRadians(Angle.Degrees(f.AngleVAdjust))), t.Basis)
	node.SetTransform(t)
}
