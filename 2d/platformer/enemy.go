package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/RayCast2D"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

type EnemyState int

const (
	EnemyWalking EnemyState = iota
	EnemyDead
)

const EnemyWalkSpeed = 22.0

type Enemy struct {
	CharacterBody2D.Extension[Enemy]

	PlatformDetector   RayCast2D.Instance
	FloorDetectorLeft  RayCast2D.Instance
	FloorDetectorRight RayCast2D.Instance
	Sprite             Sprite2D.Instance `gd:"Sprite2D"`
	AnimationPlayer    AnimationPlayer.Instance

	state   EnemyState
	gravity Float.X
}

func (e *Enemy) Ready() {
	e.gravity = defaultGravity()
}

func (e *Enemy) PhysicsProcess(delta Float.X) {
	body := e.AsCharacterBody2D()
	velocity := body.Velocity()
	if e.state == EnemyWalking && Vector2.IsApproximatelyZero(velocity) {
		velocity.X = EnemyWalkSpeed
	}
	velocity.Y += e.gravity * delta
	if !e.FloorDetectorLeft.IsColliding() {
		velocity.X = EnemyWalkSpeed
	} else if !e.FloorDetectorRight.IsColliding() {
		velocity.X = -EnemyWalkSpeed
	}

	if body.IsOnWall() {
		velocity.X = -velocity.X
	}

	body.SetVelocity(velocity)
	body.MoveAndSlide()
	velocity = body.Velocity()

	scale := e.Sprite.AsNode2D().Scale()
	if velocity.X > 0 {
		scale.X = 0.8
	} else if velocity.X < 0 {
		scale.X = -0.8
	}
	e.Sprite.AsNode2D().SetScale(scale)

	animation := e.getNewAnimation()
	if animation != e.AnimationPlayer.CurrentAnimation() {
		e.AnimationPlayer.MoreArgs().Play(animation, -1, 1, false)
	}
}

func (e *Enemy) Destroy() {
	e.state = EnemyDead
	e.AsCharacterBody2D().SetVelocity(Vector2.Zero)
}

func (e *Enemy) getNewAnimation() string {
	if e.state == EnemyWalking {
		if e.AsCharacterBody2D().Velocity().X == 0 {
			return "idle"
		}
		return "walk"
	}
	return "destroy"
}
