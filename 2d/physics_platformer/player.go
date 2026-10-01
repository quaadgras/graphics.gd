package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/AudioStreamPlayer2D"
	"graphics.gd/classdb/CPUParticles2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Marker2D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/PhysicsDirectBodyState2D"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/RigidBody2D"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

const (
	WalkAccel            = 1000.0
	WalkDeaccel          = 1000.0
	WalkMaxVelocity      = 200.0
	AirAccel             = 250.0
	AirDeaccel           = 250.0
	JumpVelocity         = 380.0
	StopJumpForce        = 450.0
	MaxShootPoseTime     = 0.3
	MaxFloorAirborneTime = 0.15
)

var (
	BulletScene = Resource.Load[PackedScene.Is[*Bullet]]("res://player/bullet.tscn")
	EnemyScene  = Resource.Load[PackedScene.Is[*Enemy]]("res://enemy/enemy.tscn")
)

type Player struct {
	RigidBody2D.Extension[Player]

	SoundJump       AudioStreamPlayer2D.Instance
	SoundShoot      AudioStreamPlayer2D.Instance
	Sprite          Sprite2D.Instance       `gd:"Sprite2D"`
	SpriteSmoke     CPUParticles2D.Instance `gd:"Sprite2D/Smoke"`
	AnimationPlayer AnimationPlayer.Instance
	BulletShoot     Marker2D.Instance

	anim         string
	sidingLeft   bool
	jumping      bool
	stoppingJump bool
	shooting     bool

	floorHVelocity Float.X

	airborneTime Float.X
	shootTime    Float.X
}

func NewPlayer() *Player {
	return &Player{airborneTime: 1e20, shootTime: 1e20}
}

func (p *Player) IntegrateForces(state PhysicsDirectBodyState2D.Instance) {
	velocity := state.LinearVelocity()
	step := state.Step()

	newAnim := p.anim
	newSidingLeft := p.sidingLeft

	// Get player input.
	moveLeft := Input.IsActionPressed("move_left", false)
	moveRight := Input.IsActionPressed("move_right", false)
	jump := Input.IsActionPressed("jump", false)
	shoot := Input.IsActionPressed("shoot", false)
	spawn := Input.IsActionJustPressed("spawn", false)

	if spawn {
		Callable.Defer(Callable.New(p.spawnEnemyAbove))
	}

	// Deapply previous floor velocity.
	velocity.X -= p.floorHVelocity
	p.floorHVelocity = 0

	// Find the floor (a contact with upwards facing collision normal).
	foundFloor := false
	floorIndex := -1

	for i := range state.GetContactCount() {
		collisionNormal := state.GetContactLocalNormal(i)
		if Vector2.Dot(collisionNormal, Vector2.XY{X: 0, Y: -1}) > 0.6 {
			foundFloor = true
			floorIndex = i
		}
	}

	// A good idea when implementing characters of all kinds,
	// compensates for physics imprecision, as well as human reaction delay.
	if shoot && !p.shooting {
		Callable.Defer(Callable.New(p.shotBullet))
	} else {
		p.shootTime += step
	}

	if foundFloor {
		p.airborneTime = 0
	} else {
		p.airborneTime += step // Time it spent in the air.
	}

	onFloor := p.airborneTime < MaxFloorAirborneTime

	// Process jump.
	if p.jumping {
		if velocity.Y > 0 {
			// Set off the jumping flag if going down.
			p.jumping = false
		} else if !jump {
			p.stoppingJump = true
		}
		if p.stoppingJump {
			velocity.Y += StopJumpForce * step
		}
	}

	if onFloor {
		// Process logic when character is on floor.
		if moveLeft && !moveRight {
			if velocity.X > -WalkMaxVelocity {
				velocity.X -= WalkAccel * step
			}
		} else if moveRight && !moveLeft {
			if velocity.X < WalkMaxVelocity {
				velocity.X += WalkAccel * step
			}
		} else {
			xv := Float.Abs(velocity.X)
			xv -= WalkDeaccel * step
			if xv < 0 {
				xv = 0
			}
			velocity.X = Float.Sign(velocity.X) * xv
		}

		// Check jump.
		if !p.jumping && jump {
			velocity.Y = -JumpVelocity
			p.jumping = true
			p.stoppingJump = false
			p.SoundJump.Play()
		}

		// Check siding.
		if velocity.X < 0 && moveLeft {
			newSidingLeft = true
		} else if velocity.X > 0 && moveRight {
			newSidingLeft = false
		}
		if p.jumping {
			newAnim = "jumping"
		} else if Float.Abs(velocity.X) < 0.1 {
			if p.shootTime < MaxShootPoseTime {
				newAnim = "idle_weapon"
			} else {
				newAnim = "idle"
			}
		} else {
			if p.shootTime < MaxShootPoseTime {
				newAnim = "run_weapon"
			} else {
				newAnim = "run"
			}
		}
	} else {
		// Process logic when the character is in the air.
		if moveLeft && !moveRight {
			if velocity.X > -WalkMaxVelocity {
				velocity.X -= AirAccel * step
			}
		} else if moveRight && !moveLeft {
			if velocity.X < WalkMaxVelocity {
				velocity.X += AirAccel * step
			}
		} else {
			xv := Float.Abs(velocity.X)
			xv -= AirDeaccel * step
			if xv < 0 {
				xv = 0
			}
			velocity.X = Float.Sign(velocity.X) * xv
		}

		if velocity.Y < 0 {
			if p.shootTime < MaxShootPoseTime {
				newAnim = "jumping_weapon"
			} else {
				newAnim = "jumping"
			}
		} else {
			if p.shootTime < MaxShootPoseTime {
				newAnim = "falling_weapon"
			} else {
				newAnim = "falling"
			}
		}
	}

	// Update siding.
	if newSidingLeft != p.sidingLeft {
		scale := p.Sprite.AsNode2D().Scale()
		if newSidingLeft {
			scale.X = -1
		} else {
			scale.X = 1
		}
		p.Sprite.AsNode2D().SetScale(scale)
		p.sidingLeft = newSidingLeft
	}

	// Change animation.
	if newAnim != p.anim {
		p.anim = newAnim
		p.AnimationPlayer.MoreArgs().Play(p.anim, -1, 1, false)
	}

	p.shooting = shoot

	// Apply floor velocity.
	if foundFloor {
		p.floorHVelocity = state.GetContactColliderVelocityAtPosition(floorIndex).X
		velocity.X += p.floorHVelocity
	}

	// Finally, apply gravity and set back the linear velocity.
	velocity = Vector2.Add(velocity, Vector2.MulX(state.TotalGravity(), step))
	state.SetLinearVelocity(velocity)
}

func (p *Player) shotBullet() {
	p.shootTime = 0
	bullet := PackedScene.Instantiate[*Bullet](PackedScene.Instance(BulletScene))
	var speedScale Float.X = 1
	if p.sidingLeft {
		speedScale = -1
	}

	node := p.AsNode2D()
	bullet.AsNode2D().SetPosition(Vector2.Add(node.Position(), Vector2.Mul(p.BulletShoot.AsNode2D().Position(), Vector2.XY{X: speedScale, Y: 1})))
	node.AsNode().GetParent().AddChild(bullet.AsNode())

	bullet.AsRigidBody2D().SetLinearVelocity(Vector2.XY{X: 400 * speedScale, Y: -40})

	p.SpriteSmoke.Restart()
	p.SoundShoot.Play()

	p.AsPhysicsBody2D().AddCollisionExceptionWith(bullet.AsNode()) // Make bullet and this not collide.
}

func (p *Player) spawnEnemyAbove() {
	enemy := PackedScene.Instantiate[*Enemy](PackedScene.Instance(EnemyScene))
	enemy.AsNode2D().SetPosition(Vector2.Add(p.AsNode2D().Position(), Vector2.MulX(Vector2.Up, 50)))
	p.AsNode().GetParent().AddChild(enemy.AsNode())
}
