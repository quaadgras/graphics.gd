package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Area2D"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node2D"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
)

type Player struct {
	CharacterBody2D.Extension[Player]

	gravity Float.X
}

const (
	WalkForce    = 600
	WalkMaxSpeed = 200
	StopForce    = 1300
	JumpSpeed    = 200
)

func (p *Player) Ready() {
	p.gravity = Float.X(ProjectSettings.GetSetting("physics/2d/default_gravity", 0).(int64))
}

func (p *Player) PhysicsProcess(delta Float.X) {
	var velocity = p.AsCharacterBody2D().Velocity()
	var walk = WalkForce * Input.GetAxis("move_left", "move_right")
	if Float.Abs(walk) < WalkForce*0.2 {
		velocity.X = Float.MoveToward(velocity.X, 0, StopForce*delta)
	} else {
		velocity.X += walk * delta
	}
	velocity.X = Float.Clamp(velocity.X, -WalkMaxSpeed, WalkMaxSpeed)
	velocity.Y += p.gravity * delta
	p.AsCharacterBody2D().SetVelocity(velocity)
	p.AsCharacterBody2D().MoveAndSlide()
	velocity = p.AsCharacterBody2D().Velocity()
	if p.AsCharacterBody2D().IsOnFloor() && Input.IsActionJustPressed("jump", false) {
		velocity.Y = -JumpSpeed
	}
	p.AsCharacterBody2D().SetVelocity(velocity)
}

type Princess struct {
	Area2D.Extension[Princess]

	WinText Label.Instance `gd:"../WinText"`
}

func (p *Princess) OnBodyEntered(body Node2D.Instance) {
	if body.AsNode().Name() == "Player" {
		p.WinText.AsCanvasItem().Show()
	}
}

func main() {
	classdb.Register[Player]()
	classdb.Register[Princess]()
	startup.Scene()
}
