package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/TileData"
	"graphics.gd/classdb/TileMapLayer"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2i"
)

const (
	WalkForce    = 600
	WalkMaxSpeed = 200
	StopForce    = 1300
	JumpSpeed    = 200
)

type Player struct {
	CharacterBody2D.Extension[Player] `gd:"DynamicTilemapLayersPlayer"`

	gravity Float.X
}

func (p *Player) Ready() {
	switch g := ProjectSettings.GetSetting("physics/2d/default_gravity", 500).(type) {
	case int:
		p.gravity = Float.X(g)
	case int64:
		p.gravity = Float.X(g)
	case float64:
		p.gravity = Float.X(g)
	}
}

func (p *Player) PhysicsProcess(delta Float.X) {
	var (
		walk = WalkForce * (Input.GetAxis("move_left", "move_right")) // Horizontal movement code. First, get the player's input.
	)
	velocity := p.AsCharacterBody2D().Velocity()
	if Float.Abs(walk) < WalkForce*0.2 { //  Slow down the player if they're not trying to move.
		velocity.X = Float.MoveToward(velocity.X, 0, StopForce*delta) // The velocity, slowed down a bit, and then reassigned.
	} else {
		velocity.X += walk * delta
	}
	velocity.X = Float.Clamp(velocity.X, -WalkMaxSpeed, WalkMaxSpeed) // Clamp to the maximum horizontal movement speed.
	velocity.Y += p.gravity * delta                                   // Vertical movement code. Apply gravity.
	// Move based on the velocity and snap to the ground.
	// TODO: This information should be set to the CharacterBody properties instead of arguments: snap, Vector2.DOWN, Vector2.UP
	// TODO: Rename velocity to linear_velocity in the rest of the script.
	p.AsCharacterBody2D().MoveAndSlide()
	// Check for jumping. is_on_floor() must be called after movement code.
	if p.AsCharacterBody2D().IsOnFloor() && Input.IsActionJustPressed("jump", false) {
		velocity.Y = -JumpSpeed
	}
	p.AsCharacterBody2D().SetVelocity(velocity)
}

type Level struct {
	TileMapLayer.Extension[Level] `gd:"DynamicTilemapLayersLevel"`

	player_in_secret bool
	layer_alpha      Float.X
}

func NewLevel() *Level {
	return &Level{
		layer_alpha: 1,
	}
}

func (level *Level) Ready() {
	level.AsNode().SetProcess(false)
}

func (level *Level) Process(delta Float.X) {
	if level.player_in_secret {
		if level.layer_alpha > 0.3 {
			level.layer_alpha = Float.MoveToward(level.layer_alpha, 0.3, delta)
			level.AsCanvasItem().SetSelfModulate(Color.RGBA{1, 1, 1, level.layer_alpha})
		} else {
			level.AsNode().SetProcess(false)
		}
	} else {
		if level.layer_alpha < 1 {
			level.layer_alpha = Float.MoveToward(level.layer_alpha, 1, delta)
			level.AsCanvasItem().SetSelfModulate(Color.RGBA{1, 1, 1, level.layer_alpha})
		} else {
			level.AsNode().SetProcess(false)
		}
	}
}

func (level *Level) UseTileDataRuntimeUpdate(Vector2i.XY) bool { return true }

func (level *Level) TileDataRuntimeUpdate(_ Vector2i.XY, data TileData.Instance) {
	data.SetCollisionPolygonsCount(0, 0)
}

func (level *Level) OnSecretDetectorBodyEntered(body Node2D.Instance) {
	if _, ok := Object.As[CharacterBody2D.Instance](body); !ok {
		return
	}
	level.player_in_secret = true
	level.AsNode().SetProcess(true)
}

func (level *Level) OnSecretDetectorBodyExited(body Node2D.Instance) {
	if _, ok := Object.As[CharacterBody2D.Instance](body); !ok {
		return
	}
	level.player_in_secret = false
	level.AsNode().SetProcess(true)
}

func main() {
	classdb.Register[Player]()
	classdb.Register[Level](NewLevel)
	startup.Scene()
}
