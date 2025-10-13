package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/AnimatedSprite2D"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/startup"
	"graphics.gd/variant/Angle"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Vector2"
)

const (
	MotionSpeed = 160
)

type AnimDirection struct {
	Name  string
	FlipH bool
}

var AnimDirections = map[string][8]AnimDirection{
	"idle": {
		{"side_right_idle", false},
		{"45front_right_idle", false},
		{"front_idle", false},
		{"45front_left_idle", false},
		{"side_left_idle", false},
		{"45back_left_idle", false},
		{"back_idle", false},
		{"45back_right_idle", false},
	},
	"walk": {
		{"side_right_walk", false},
		{"45front_right_walk", false},
		{"front_walk", false},
		{"45front_left_walk", false},
		{"side_left_walk", false},
		{"45back_left_walk", false},
		{"back_walk", false},
		{"45back_right_walk", false},
	},
}

type Goblin struct {
	CharacterBody2D.Extension[Goblin]

	Sprite2D AnimatedSprite2D.Instance

	LastDirection Vector2.XY
}

func (g *Goblin) PhysicsProcess(delta Float.X) {
	var motion Vector2.XY
	motion.X = Input.GetActionStrength("move_right", false) - Input.GetActionStrength("move_left", false)
	motion.Y = Input.GetActionStrength("move_down", false) - Input.GetActionStrength("move_up", false)
	motion.Y /= 2
	motion = Vector2.MulX(Vector2.Normalized(motion), MotionSpeed)
	g.AsCharacterBody2D().SetVelocity(motion)
	g.AsCharacterBody2D().MoveAndSlide()
	var dir = g.AsCharacterBody2D().Velocity()
	if Vector2.Length(dir) > 0 {
		g.LastDirection = Vector2.Normalized(dir)
		g.UpdateAnimation("walk")
	} else {
		g.UpdateAnimation("idle")
	}
}

func (g *Goblin) UpdateAnimation(state string) {
	var angle = Angle.InDegrees(Vector2.AngleRadians(g.LastDirection)) + 22.5
	var dir = int(Float.Floor(angle/45)) % 8
	if dir < 0 {
		dir = 8 + dir
	}
	g.Sprite2D.PlayNamed(AnimDirections[state][dir].Name)
	g.Sprite2D.SetFlipH(AnimDirections[state][dir].FlipH)
}

func main() {
	classdb.Register[Goblin]()
	startup.Scene()
}
