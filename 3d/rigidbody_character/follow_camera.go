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

type FollowCamera struct {
	Camera3D.Extension[FollowCamera]

	MinDistance  Float.X
	MaxDistance  Float.X
	AngleVAdjust Float.X

	collisionException []RID.Body3D
	maxHeight          Float.X
	minHeight          Float.X

	targetNode Node3D.Instance
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
	f.targetNode, _ = Object.As[Node3D.Instance](f.AsNode().GetParent())
	if body, ok := Object.As[CollisionObject3D.Instance](f.targetNode.AsNode().GetParent()); ok {
		f.collisionException = append(f.collisionException, body.GetRid())
	}
	// Detach the camera transform from the parent spatial node.
	f.AsNode3D().SetTopLevel(true)
}

func (f *FollowCamera) PhysicsProcess(_ Float.X) {
	node3d := f.AsNode3D()
	var target_pos = f.targetNode.GlobalTransform().Origin
	var camera_pos = node3d.GlobalTransform().Origin

	var delta_pos = Vector3.Sub(camera_pos, target_pos)

	// Regular delta follow.

	// Check ranges.
	if Vector3.Length(delta_pos) < f.MinDistance {
		delta_pos = Vector3.MulX(Vector3.Normalized(delta_pos), f.MinDistance)
	} else if Vector3.Length(delta_pos) > f.MaxDistance {
		delta_pos = Vector3.MulX(Vector3.Normalized(delta_pos), f.MaxDistance)
	}

	// Check upper and lower height.
	delta_pos.Y = Float.Clamp(delta_pos.Y, f.minHeight, f.maxHeight)
	camera_pos = Vector3.Add(target_pos, delta_pos)

	node3d.MoreArgs().LookAtFromPosition(camera_pos, target_pos, Vector3.Up, false)

	// Turn a little up or down.
	var t = node3d.Transform()
	t.Basis = Basis.Mul(Basis.RotatesAxisAngle(t.Basis.X, Angle.InRadians(Angle.Degrees(f.AngleVAdjust))), t.Basis)
	node3d.SetTransform(t)
}
