package main

import (
	"graphics.gd/classdb/AnimationPlayer"
	"graphics.gd/classdb/AudioStreamPlayer2D"
	"graphics.gd/classdb/Camera2D"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/RayCast2D"
	"graphics.gd/classdb/Sprite2D"
	"graphics.gd/classdb/Timer"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/Vector2"
)

const (
	WalkSpeed         = 300.0
	AccelerationSpeed = WalkSpeed * 6.0
	JumpVelocity      = -725.0
	TerminalVelocity  = 700 // Maximum speed at which the player can fall.
)

type Player struct {
	CharacterBody2D.Extension[Player]

	CoinCollected Signal.Void `gd:"coin_collected"`

	// ActionSuffix is appended to the input actions the player listens for.
	// Used to separate controls for multiple players in splitscreen.
	ActionSuffix string

	PlatformDetector RayCast2D.Instance
	AnimationPlayer  AnimationPlayer.Instance
	ShootTimer       Timer.Instance               `gd:"ShootAnimation"`
	Sprite           Sprite2D.Instance            `gd:"Sprite2D"`
	JumpSound        AudioStreamPlayer2D.Instance `gd:"Jump"`
	Gun              *Gun                         `gd:"Sprite2D/Gun"`
	Camera           Camera2D.Instance

	gravity           Float.X
	doubleJumpCharged bool
}

func (p *Player) Ready() {
	p.gravity = defaultGravity()
}

func (p *Player) PhysicsProcess(delta Float.X) {
	body := p.AsCharacterBody2D()
	velocity := body.Velocity()
	if body.IsOnFloor() {
		p.doubleJumpCharged = true
	}
	if Input.IsActionJustPressed("jump"+p.ActionSuffix, false) {
		velocity = p.tryJump(velocity)
	} else if Input.IsActionJustReleased("jump"+p.ActionSuffix, false) && velocity.Y < 0 {
		// The player let go of jump early, reduce vertical momentum.
		velocity.Y *= 0.6
	}
	// Fall.
	velocity.Y = min(TerminalVelocity, velocity.Y+p.gravity*delta)

	direction := Input.GetAxis("move_left"+p.ActionSuffix, "move_right"+p.ActionSuffix) * WalkSpeed
	velocity.X = Float.MoveToward(velocity.X, direction, AccelerationSpeed*delta)

	if !Float.IsApproximatelyZero(velocity.X) {
		scale := p.Sprite.AsNode2D().Scale()
		if velocity.X > 0 {
			scale.X = 1
		} else {
			scale.X = -1
		}
		p.Sprite.AsNode2D().SetScale(scale)
	}

	body.SetFloorStopOnSlope(!p.PlatformDetector.IsColliding())
	body.SetVelocity(velocity)
	body.MoveAndSlide()

	isShooting := false
	if Input.IsActionJustPressed("shoot"+p.ActionSuffix, false) {
		isShooting = p.Gun.Shoot(p.Sprite.AsNode2D().Scale().X)
	}

	animation := p.getNewAnimation(isShooting)
	if animation != p.AnimationPlayer.CurrentAnimation() && p.ShootTimer.IsStopped() {
		if isShooting {
			p.ShootTimer.Start()
		}
		p.AnimationPlayer.MoreArgs().Play(animation, -1, 1, false)
	}
}

func (p *Player) getNewAnimation(isShooting bool) string {
	body := p.AsCharacterBody2D()
	var animation string
	if body.IsOnFloor() {
		if Float.Abs(body.Velocity().X) > 0.1 {
			animation = "run"
		} else {
			animation = "idle"
		}
	} else {
		if body.Velocity().Y > 0 {
			animation = "falling"
		} else {
			animation = "jumping"
		}
	}
	if isShooting {
		animation += "_weapon"
	}
	return animation
}

func (p *Player) tryJump(velocity Vector2.XY) Vector2.XY {
	if p.AsCharacterBody2D().IsOnFloor() {
		p.JumpSound.SetPitchScale(1.0)
	} else if p.doubleJumpCharged {
		p.doubleJumpCharged = false
		velocity.X *= 2.5
		p.JumpSound.SetPitchScale(1.5)
	} else {
		return velocity
	}
	velocity.Y = JumpVelocity
	p.JumpSound.Play()
	return velocity
}

// defaultGravity reads the project's 2D gravity setting.
func defaultGravity() Float.X {
	switch g := ProjectSettings.GetSetting("physics/2d/default_gravity", 980).(type) {
	case int:
		return Float.X(g)
	case int64:
		return Float.X(g)
	case float64:
		return Float.X(g)
	}
	return 980
}
