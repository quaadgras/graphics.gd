package main

import (
	"math"

	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/CharacterBody3D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/Vector3"
)

type Player struct {
	CharacterBody3D.Extension[Player]

	Hit Signal.Void `gd:"hit"`

	Speed            Float.X // How fast the player moves in meters per second.
	JumpImpulse      Float.X // Vertical impulse applied to the character upon jumping in meters per second.
	BounceImpulse    Float.X // Vertical impulse applied to the character upon bouncing over a mob in meters per second.
	FallAcceleration Float.X // The downward acceleration when in the air, in meters per second.

	AnimationPlayer AnimationPlayer.Instance
}

func NewPlayer() *Player {
	return &Player{Speed: 14, JumpImpulse: 20, BounceImpulse: 16, FallAcceleration: 75}
}

func (p *Player) PhysicsProcess(delta Float.X) {
	body := p.AsCharacterBody3D()
	node := p.AsNode3D()

	var direction Vector3.XYZ
	if Input.IsActionPressed("move_right", false) {
		direction.X += 1
	}
	if Input.IsActionPressed("move_left", false) {
		direction.X -= 1
	}
	if Input.IsActionPressed("move_back", false) {
		direction.Z += 1
	}
	if Input.IsActionPressed("move_forward", false) {
		direction.Z -= 1
	}

	if direction != Vector3.Zero {
		// In the lines below, we turn the character when moving and make the animation play faster.
		direction = Vector3.Normalized(direction)
		// Setting the basis property will affect the rotation of the node.
		node.SetBasis(Basis.LookingAt(direction, Vector3.Up))
		p.AnimationPlayer.SetSpeedScale(4)
	} else {
		p.AnimationPlayer.SetSpeedScale(1)
	}

	velocity := body.Velocity()
	velocity.X = direction.X * p.Speed
	velocity.Z = direction.Z * p.Speed

	// Jumping.
	if body.IsOnFloor() && Input.IsActionJustPressed("jump", false) {
		velocity.Y += p.JumpImpulse
	}

	// We apply gravity every frame so the character always collides with the ground when moving.
	// This is necessary for the is_on_floor() function to work as a body can always detect
	// the floor, walls, etc. when a collision happens the same frame.
	velocity.Y -= p.FallAcceleration * delta
	body.SetVelocity(velocity)
	body.MoveAndSlide()
	velocity = body.Velocity()

	// Here, we check if we landed on top of a mob and if so, we kill it and bounce.
	// With move_and_slide(), Godot makes the body move sometimes multiple times in a row to
	// smooth out the character's motion. So we have to loop over all collisions that may have
	// happened.
	// If there are no "slides" this frame, the loop below won't run.
	for index := range body.GetSlideCollisionCount() {
		collision := body.GetSlideCollision(index)
		if mob, ok := Object.As[*Mob](collision.GetCollider()); ok {
			if Vector3.Dot(Vector3.Up, collision.GetNormal()) > 0.1 {
				mob.Squash()
				velocity.Y = p.BounceImpulse
				body.SetVelocity(velocity)
				// Prevent this block from running more than once,
				// which would award the player more than 1 point for squashing a single mob.
				break
			}
		}
	}

	// This makes the character follow a nice arc when jumping
	rotation := node.Rotation()
	rotation.X = Angle.Radians(math.Pi / 6 * velocity.Y / p.JumpImpulse)
	node.SetRotation(rotation)
}

func (p *Player) Die() {
	p.Hit.Emit()
	p.AsNode().QueueFree()
}

func (p *Player) OnMobDetectorBodyEntered(_ Node3D.Instance) {
	p.Die()
}
