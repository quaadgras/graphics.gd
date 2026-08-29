package main

import (
	"math"

	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/CharacterBody3D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/Vector3"
)

type Mob struct {
	CharacterBody3D.Extension[Mob]

	// Squashed is emitted when the player jumped on the mob.
	Squashed Signal.Void `gd:"squashed"`

	MinSpeed Float.X // Minimum speed of the mob in meters per second.
	MaxSpeed Float.X // Maximum speed of the mob in meters per second.

}

func NewMob() *Mob {
	return &Mob{MinSpeed: 10, MaxSpeed: 18}
}

func (m *Mob) PhysicsProcess(_ Float.X) {
	m.AsCharacterBody3D().MoveAndSlide()
}

func (m *Mob) Initialize(startPosition, playerPosition Vector3.XYZ) {
	node := m.AsNode3D()
	// Ignore the player's height, so that the mob's orientation is not slightly
	// shifted if the mob spawns while the player is jumping.
	target := Vector3.XYZ{playerPosition.X, startPosition.Y, playerPosition.Z}
	node.LookAtFromPosition(startPosition, target)

	// Rotate this mob randomly within range of -45 and +45 degrees,
	// so that it doesn't move directly towards the player.
	node.RotateY(Angle.Radians(Float.RandomBetween(-math.Pi/4, math.Pi/4)))

	randomSpeed := Float.RandomBetween(m.MinSpeed, m.MaxSpeed)
	// We calculate a forward velocity first, which represents the speed.
	velocity := Vector3.MulX(Vector3.Forward, randomSpeed)
	// We then rotate the vector based on the mob's Y rotation to move in the direction it's looking.
	velocity = Vector3.Rotated(velocity, Vector3.Up, node.Rotation().Y)
	m.AsCharacterBody3D().SetVelocity(velocity)

	// Initialize runs before the mob enters the tree, so the declarative
	// AnimationPlayer field hasn't been resolved yet; look the child up directly.
	animation := Object.To[AnimationPlayer.Instance](m.AsNode().GetNode("AnimationPlayer"))
	animation.SetSpeedScale(randomSpeed / m.MinSpeed)
}

func (m *Mob) Squash() {
	m.Squashed.Emit()
	m.AsNode().QueueFree()
}

func (m *Mob) OnVisibleOnScreenNotifierScreenExited() {
	m.AsNode().QueueFree()
}
