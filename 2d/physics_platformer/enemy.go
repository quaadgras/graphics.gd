package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/AudioStreamPlayer2D"
	"graphics.gd/classdb/CollisionShape2D"
	"graphics.gd/classdb/PhysicsDirectBodyState2D"
	"graphics.gd/classdb/RayCast2D"
	"graphics.gd/classdb/RigidBody2D"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

const EnemyWalkSpeed = 50

type EnemyState int

const (
	EnemyWalking EnemyState = iota
	EnemyDying
)

type Enemy struct {
	RigidBody2D.Extension[Enemy]

	RaycastLeft     RayCast2D.Instance
	RaycastRight    RayCast2D.Instance
	Sprite          Sprite2D.Instance `gd:"Sprite2D"`
	AnimationPlayer AnimationPlayer.Instance
	SoundExplode    AudioStreamPlayer2D.Instance
	SoundHit        AudioStreamPlayer2D.Instance
	Shape1          CollisionShape2D.Instance
	Shape2          CollisionShape2D.Instance
	Shape3          CollisionShape2D.Instance

	state     EnemyState
	direction Float.X
	anim      string
}

func NewEnemy() *Enemy {
	return &Enemy{direction: -1}
}

func (e *Enemy) flip() {
	e.direction = -e.direction
	scale := e.Sprite.AsNode2D().Scale()
	scale.X = -e.direction
	e.Sprite.AsNode2D().SetScale(scale)
}

func (e *Enemy) IntegrateForces(state PhysicsDirectBodyState2D.Instance) {
	velocity := state.LinearVelocity()
	newAnim := e.anim

	if e.state == EnemyDying {
		newAnim = "explode"
	} else if e.state == EnemyWalking {
		newAnim = "walk"

		var wallSide Float.X

		for i := range state.GetContactCount() {
			collider := state.GetContactColliderObject(i)
			collisionNormal := state.GetContactLocalNormal(i)

			if bullet, ok := Object.As[*Bullet](collider); ok && !bullet.disabled {
				Callable.Defer(Callable.New(func() { e.bulletCollider(bullet, state, collisionNormal) }))
				break
			}

			if collisionNormal.X > 0.9 {
				wallSide = 1
			} else if collisionNormal.X < -0.9 {
				wallSide = -1
			}
		}

		if wallSide != 0 && wallSide != e.direction {
			e.flip()
		}
		if e.direction < 0 && !e.RaycastLeft.IsColliding() && e.RaycastRight.IsColliding() {
			e.flip()
		} else if e.direction > 0 && !e.RaycastRight.IsColliding() && e.RaycastLeft.IsColliding() {
			e.flip()
		}

		velocity.X = e.direction * EnemyWalkSpeed
	}

	if e.anim != newAnim {
		e.anim = newAnim
		e.AnimationPlayer.MoreArgs().Play(e.anim, -1, 1, false)
	}

	state.SetLinearVelocity(velocity)
}

// Die is called from the "explode" animation's method track.
func (e *Enemy) Die() {
	e.AsNode().QueueFree()
}

// PreExplode is called from the "explode" animation's method track.
func (e *Enemy) PreExplode() {
	// Make sure nothing collides against this.
	e.Shape1.AsNode().QueueFree()
	e.Shape2.AsNode().QueueFree()
	e.Shape3.AsNode().QueueFree()

	e.SoundExplode.Play()
}

func (e *Enemy) bulletCollider(bullet *Bullet, state PhysicsDirectBodyState2D.Instance, collisionNormal Vector2.XY) {
	e.state = EnemyDying

	state.SetAngularVelocity(Float.Sign(collisionNormal.X) * 33.0)
	e.AsRigidBody2D().PhysicsMaterialOverride().SetFriction(1)
	bullet.Disable()
	e.SoundHit.Play()
}
