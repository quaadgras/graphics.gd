package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimationNodeOneShot"
	"graphics.gd/classdb/AnimationTree"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

// Keep these in sync with the AnimationTree's state names.
const (
	StateIdle = "idle"
	StateWalk = "walk"
	StateRun  = "run"
	StateFly  = "fly"
	StateFall = "fall"
)

const (
	WalkSpeed         = 200.0
	AccelerationSpeed = WalkSpeed * 6.0
	JumpVelocity      = -400.0
	// TerminalVelocity is the maximum speed at which the player can fall.
	TerminalVelocity = 400
)

type Player struct {
	CharacterBody2D.Extension[Player]

	Sprite        Node2D.Instance        `gd:"Sprite2D"`
	AnimationTree AnimationTree.Instance `gd:"AnimationTree"`

	fallingSlow          bool
	fallingFast          bool
	noMoveHorizontalTime Float.X

	gravity     Float.X
	spriteScale Float.X
}

func (p *Player) Ready() {
	switch g := ProjectSettings.GetSetting("physics/2d/default_gravity", 0).(type) {
	case int:
		p.gravity = Float.X(g)
	case int64:
		p.gravity = Float.X(g)
	case float64:
		p.gravity = Float.X(g)
	}
	p.spriteScale = p.Sprite.Scale().X
	p.AnimationTree.AsAnimationMixer().SetActive(true)
}

func (p *Player) PhysicsProcess(delta Float.X) {
	body := p.AsCharacterBody2D()
	velocity := body.Velocity()

	var isJumping bool
	if Input.IsActionJustPressed("jump", false) {
		isJumping = p.tryJump(&velocity)
	} else if Input.IsActionJustReleased("jump", false) && velocity.Y < 0.0 {
		// The player let go of jump early, reduce vertical momentum.
		velocity.Y *= 0.6
	}
	// Fall.
	velocity.Y = min(TerminalVelocity, velocity.Y+p.gravity*delta)

	var direction = Input.GetAxis("move_left", "move_right") * WalkSpeed
	velocity.X = Float.MoveToward(velocity.X, direction, AccelerationSpeed*delta)

	if p.noMoveHorizontalTime > 0.0 {
		// After doing a hard fall, don't move for a short time.
		velocity.X = 0.0
		p.noMoveHorizontalTime -= delta
	}

	if !Float.IsApproximatelyZero(velocity.X) {
		scale := p.Sprite.Scale()
		if velocity.X > 0.0 {
			scale.X = 1.0 * p.spriteScale
		} else {
			scale.X = -1.0 * p.spriteScale
		}
		p.Sprite.SetScale(scale)
	}

	body.SetVelocity(velocity)
	body.MoveAndSlide()
	velocity = body.Velocity()

	// After applying our motion, update our animation to match.

	// Calculate falling speed for animation purposes.
	if velocity.Y >= TerminalVelocity {
		p.fallingFast = true
		p.fallingSlow = false
	} else if velocity.Y > 300 {
		p.fallingSlow = true
	}

	if isJumping {
		Object.Set(p.AnimationTree, "parameters/jump/request", AnimationNodeOneShot.OneShotRequestFire)
	}

	if body.IsOnFloor() {
		// Most animations change when we run, land, or take off.
		if p.fallingFast {
			Object.Set(p.AnimationTree, "parameters/land_hard/request", AnimationNodeOneShot.OneShotRequestFire)
			p.noMoveHorizontalTime = 0.4
		} else if p.fallingSlow {
			Object.Set(p.AnimationTree, "parameters/land/request", AnimationNodeOneShot.OneShotRequestFire)
		}

		if Float.Abs(velocity.X) > 50 {
			Object.Set(p.AnimationTree, "parameters/state/transition_request", StateRun)
			Object.Set(p.AnimationTree, "parameters/run_timescale/scale", Float.Abs(velocity.X)/60)
		} else if velocity.X != 0 {
			Object.Set(p.AnimationTree, "parameters/state/transition_request", StateWalk)
			Object.Set(p.AnimationTree, "parameters/walk_timescale/scale", Float.Abs(velocity.X)/12)
		} else {
			Object.Set(p.AnimationTree, "parameters/state/transition_request", StateIdle)
		}

		p.fallingFast = false
		p.fallingSlow = false
	} else {
		if velocity.Y > 0 {
			Object.Set(p.AnimationTree, "parameters/state/transition_request", StateFall)
		} else {
			Object.Set(p.AnimationTree, "parameters/state/transition_request", StateFly)
		}
	}
}

func (p *Player) tryJump(velocity *Vector2.XY) bool {
	if p.AsCharacterBody2D().IsOnFloor() {
		velocity.Y = JumpVelocity
		return true
	}
	return false
}

func main() {
	classdb.Register[Level]()
	classdb.Register[Player]()
	startup.Scene()
}
